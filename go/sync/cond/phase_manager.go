package cond

import "sync"

// 多阶段同步示例

type PhaseManager struct {
	mu    sync.Mutex
	cond  *sync.Cond
	phase int
}

func NewPhaseManager() *PhaseManager {
	pm := &PhaseManager{
		phase: 0,
	}
	pm.cond = sync.NewCond(&pm.mu)
	return pm
}

// 所有 worker 等待进入下一阶段
func (pm *PhaseManager) WaitForPhase(targetPhase int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for pm.phase < targetPhase {
		pm.cond.Wait()
	}
}

// 推进到下一阶段，唤醒所有等待者
func (pm *PhaseManager) NextPhase() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.phase++
	pm.cond.Broadcast() // 通知所有 worker 进入新阶段
}

// 获取当前阶段
func (pm *PhaseManager) GetPhase() int {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	return pm.phase
}
