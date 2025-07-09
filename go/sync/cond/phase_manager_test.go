package cond

import (
	"sync"
	"testing"
	"time"
)

// 测试基本的阶段推进功能
func TestPhaseManager_BasicPhaseAdvancement(t *testing.T) {
	pm := NewPhaseManager()

	// 初始阶段应该是0
	if phase := pm.GetPhase(); phase != 0 {
		t.Errorf("Expected initial phase to be 0, got %d", phase)
	}

	// 推进到阶段1
	pm.NextPhase()
	if phase := pm.GetPhase(); phase != 1 {
		t.Errorf("Expected phase to be 1, got %d", phase)
	}

	// 推进到阶段2
	pm.NextPhase()
	if phase := pm.GetPhase(); phase != 2 {
		t.Errorf("Expected phase to be 2, got %d", phase)
	}
}

// 测试单个goroutine等待阶段
func TestPhaseManager_SingleWorkerWait(t *testing.T) {
	pm := NewPhaseManager()
	done := make(chan bool)
	started := make(chan bool) // 用来确认goroutine已经开始等待

	// 启动一个goroutine等待阶段2
	go func() {
		started <- true // 通知已经启动
		pm.WaitForPhase(2)
		done <- true
	}()

	<-started // 等待goroutine确实启动

	// 推进到阶段1，worker应该继续等待
	pm.NextPhase()
	select {
	case <-done:
		t.Error("Worker should still be waiting for phase 2")
	case <-time.After(50 * time.Millisecond):
		// 正常，worker仍在等待
	}

	// 推进到阶段2，worker应该被唤醒
	pm.NextPhase()
	select {
	case <-done:
		// 正常，worker被唤醒
	case <-time.After(100 * time.Millisecond):
		t.Error("Worker should have been woken up after reaching phase 2")
	}
}

// 测试多个goroutine等待同一阶段
func TestPhaseManager_MultipleWorkersWaitSamePhase(t *testing.T) {
	pm := NewPhaseManager()
	workerCount := 5
	done := make(chan bool, workerCount)
	var startedWG sync.WaitGroup

	startedWG.Add(workerCount)

	// 启动多个goroutine等待阶段3
	for i := 0; i < workerCount; i++ {
		go func(id int) {
			startedWG.Done() // 通知已经启动
			pm.WaitForPhase(3)
			done <- true
		}(i)
	}

	startedWG.Wait() // 等待所有goroutine启动

	// 推进到阶段3
	pm.NextPhase() // 阶段1
	pm.NextPhase() // 阶段2
	pm.NextPhase() // 阶段3

	// 所有worker应该被唤醒
	for i := 0; i < workerCount; i++ {
		select {
		case <-done:
			// 正常
		case <-time.After(100 * time.Millisecond):
			t.Errorf("Worker %d should have been woken up", i)
		}
	}
}

// 测试多个goroutine等待不同阶段
func TestPhaseManager_MultipleWorkersWaitDifferentPhases(t *testing.T) {
	pm := NewPhaseManager()
	done1 := make(chan bool)
	done2 := make(chan bool)
	done3 := make(chan bool)
	var startedWG sync.WaitGroup

	startedWG.Add(3)

	// 启动goroutine等待不同阶段
	go func() {
		startedWG.Done()
		pm.WaitForPhase(1)
		done1 <- true
	}()

	go func() {
		startedWG.Done()
		pm.WaitForPhase(2)
		done2 <- true
	}()

	go func() {
		startedWG.Done()
		pm.WaitForPhase(3)
		done3 <- true
	}()

	startedWG.Wait() // 等待所有goroutine启动

	// 推进到阶段1，只有等待阶段1的worker应该被唤醒
	pm.NextPhase()
	select {
	case <-done1:
		// 正常
	case <-time.After(100 * time.Millisecond):
		t.Error("Worker waiting for phase 1 should have been woken up")
	}

	// 其他worker应该仍在等待
	select {
	case <-done2:
		t.Error("Worker waiting for phase 2 should still be waiting")
	case <-done3:
		t.Error("Worker waiting for phase 3 should still be waiting")
	case <-time.After(50 * time.Millisecond):
		// 正常
	}

	// 推进到阶段2
	pm.NextPhase()
	select {
	case <-done2:
		// 正常
	case <-time.After(100 * time.Millisecond):
		t.Error("Worker waiting for phase 2 should have been woken up")
	}

	// 推进到阶段3
	pm.NextPhase()
	select {
	case <-done3:
		// 正常
	case <-time.After(100 * time.Millisecond):
		t.Error("Worker waiting for phase 3 should have been woken up")
	}
}

// 测试等待已经到达的阶段
func TestPhaseManager_WaitForReachedPhase(t *testing.T) {
	pm := NewPhaseManager()

	// 推进到阶段3
	pm.NextPhase() // 阶段1
	pm.NextPhase() // 阶段2
	pm.NextPhase() // 阶段3

	// 使用超时保护的功能测试
	testCases := []int{0, 1, 2, 3} // 等待已到达的阶段（包括初始阶段0）

	for _, targetPhase := range testCases {
		// 使用goroutine + timeout保护，防止测试卡住
		done := make(chan bool, 1)

		go func(phase int) {
			pm.WaitForPhase(phase)
			done <- true
		}(targetPhase)

		select {
		case <-done:
			// 正常完成
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("WaitForPhase(%d) timed out - possible deadlock or bug", targetPhase)
		}
	}

	// 验证仍在正确的阶段
	if currentPhase := pm.GetPhase(); currentPhase != 3 {
		t.Errorf("Expected phase to remain 3, got %d", currentPhase)
	}
}

// 并发安全性压力测试
func TestPhaseManager_ConcurrentStress(t *testing.T) {
	pm := NewPhaseManager()
	workerCount := 100
	maxPhase := 10
	var wg sync.WaitGroup
	var startedWG sync.WaitGroup

	startedWG.Add(workerCount)

	// 启动多个worker并发等待不同阶段
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			startedWG.Done()
			targetPhase := (workerID % maxPhase) + 1
			pm.WaitForPhase(targetPhase)
		}(i)
	}

	// 等待所有worker启动后再开始推进阶段
	startedWG.Wait()

	// 逐步推进阶段
	go func() {
		for i := 0; i < maxPhase; i++ {
			pm.NextPhase()
		}
	}()

	// 等待所有worker完成
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		// 正常完成
	case <-time.After(5 * time.Second):
		t.Error("Stress test timed out")
	}

	// 验证最终阶段
	if phase := pm.GetPhase(); phase != maxPhase {
		t.Errorf("Expected final phase to be %d, got %d", maxPhase, phase)
	}
}

// 性能基准测试
func BenchmarkPhaseManager_WaitForPhase(b *testing.B) {
	pm := NewPhaseManager()
	pm.NextPhase() // 推进到阶段1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pm.WaitForPhase(1) // 等待已到达的阶段，应该立即返回
	}
}

func BenchmarkPhaseManager_NextPhase(b *testing.B) {
	pm := NewPhaseManager()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pm.NextPhase()
	}
}
