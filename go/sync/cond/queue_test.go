package cond

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 测试基本的单生产者单消费者
func TestQueue_SingleProducerConsumer(t *testing.T) {
	q := NewQueue(5)

	// 生产者
	go func() {
		for i := 0; i < 10; i++ {
			q.Put(i)
		}
	}()

	// 消费者
	var results []int
	for i := 0; i < 10; i++ {
		item := q.Get()
		results = append(results, item)
	}

	// 验证结果
	for i, v := range results {
		if v != i {
			t.Errorf("Expected %d, got %d", i, v)
		}
	}
}

// 测试多生产者多消费者
func TestQueue_MultipleProducersConsumers(t *testing.T) {
	q := NewQueue(10)
	numProducers := 3
	numConsumers := 2
	itemsPerProducer := 100
	totalItems := numProducers * itemsPerProducer

	var wg sync.WaitGroup

	// 记录所有生产的数据
	produced := make(map[int]bool)
	var producedMu sync.Mutex

	// 记录所有消费的数据
	consumed := make(map[int]bool)
	var consumedMu sync.Mutex

	// 启动生产者
	for p := 0; p < numProducers; p++ {
		wg.Add(1)
		go func(producerID int) {
			defer wg.Done()
			start := producerID * itemsPerProducer
			for i := start; i < start+itemsPerProducer; i++ {
				q.Put(i)

				producedMu.Lock()
				produced[i] = true
				producedMu.Unlock()
			}
		}(p)
	}

	// 启动消费者
	for c := 0; c < numConsumers; c++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			itemsConsumed := 0
			for itemsConsumed < totalItems/numConsumers {
				item := q.Get()

				consumedMu.Lock()
				if consumed[item] {
					t.Errorf("Item %d consumed multiple times", item)
				}
				consumed[item] = true
				consumedMu.Unlock()

				itemsConsumed++
			}
		}()
	}

	// 等待所有生产者完成
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	// 设置超时
	select {
	case <-done:
		// 验证所有生产的数据都被消费了
		if len(produced) != len(consumed) {
			t.Errorf("Produced %d items, consumed %d items", len(produced), len(consumed))
		}

		for item := range produced {
			if !consumed[item] {
				t.Errorf("Item %d was produced but not consumed", item)
			}
		}
	case <-time.After(10 * time.Second):
		t.Fatal("Test timed out")
	}
}

// 测试队列容量限制
func TestQueue_CapacityLimit(t *testing.T) {
	maxSize := 3
	q := NewQueue(maxSize)

	// 填满队列
	for i := 0; i < maxSize; i++ {
		q.Put(i)
	}

	// 测试阻塞行为
	blocked := make(chan bool, 1)
	go func() {
		q.Put(999) // 这应该阻塞
		blocked <- true
	}()

	// 确保生产者被阻塞
	select {
	case <-blocked:
		t.Error("Producer should be blocked when queue is full")
	case <-time.After(100 * time.Millisecond):
		// 预期行为：生产者被阻塞
	}

	// 消费一个元素，解除阻塞
	item := q.Get()
	if item != 0 {
		t.Errorf("Expected 0, got %d", item)
	}

	// 现在生产者应该能够继续
	select {
	case <-blocked:
		// 预期行为：生产者解除阻塞
	case <-time.After(1 * time.Second):
		t.Error("Producer should be unblocked after consuming")
	}
}

// 测试空队列阻塞行为
func TestQueue_EmptyQueueBlocking(t *testing.T) {
	q := NewQueue(5)

	// 测试从空队列获取元素会阻塞
	blocked := make(chan bool, 1)
	go func() {
		q.Get() // 这应该阻塞
		blocked <- true
	}()

	// 确保消费者被阻塞
	select {
	case <-blocked:
		t.Error("Consumer should be blocked when queue is empty")
	case <-time.After(100 * time.Millisecond):
		// 预期行为：消费者被阻塞
	}

	// 生产一个元素，解除阻塞
	q.Put(42)

	// 现在消费者应该能够继续
	select {
	case <-blocked:
		// 预期行为：消费者解除阻塞
	case <-time.After(1 * time.Second):
		t.Error("Consumer should be unblocked after producing")
	}
}

// 压力测试：大量并发操作
func TestQueue_StressTest(t *testing.T) {
	q := NewQueue(50)
	numGoroutines := 20
	operationsPerGoroutine := 1000

	var wg sync.WaitGroup

	// 统计操作次数
	var putCount, getCount atomic.Int64

	// 启动生产者和消费者
	for i := 0; i < numGoroutines; i++ {
		// 生产者
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				q.Put(id*operationsPerGoroutine + j)
				putCount.Add(1)
			}
		}(i)

		// 消费者
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				q.Get()
				getCount.Add(1)
			}
		}()
	}

	// 等待完成
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		expectedOps := int64(numGoroutines * operationsPerGoroutine)
		if putCount.Load() != expectedOps {
			t.Errorf("Expected %d put operations, got %d", expectedOps, putCount.Load())
		}
		if getCount.Load() != expectedOps {
			t.Errorf("Expected %d get operations, got %d", expectedOps, getCount.Load())
		}
	case <-time.After(30 * time.Second):
		t.Fatal("Stress test timed out")
	}
}

// 基准测试
func BenchmarkQueue_PutGet(b *testing.B) {
	q := NewQueue(100)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// 每次先放入再取出，避免阻塞
			q.Put(1)
			q.Get()
		}
	})
}

func BenchmarkQueue_ProducerConsumer(b *testing.B) {
	q := NewQueue(10)

	// 使用两个单独的基准测试函数
	b.Run("Put", func(b *testing.B) {
		// 启动一个后台消费者确保队列不会满
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					q.Get()
				}
			}
		}()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			q.Put(i)
		}
		done <- true
	})

	b.Run("Get", func(b *testing.B) {
		// 启动一个后台生产者确保队列不会空
		done := make(chan bool)
		go func() {
			i := 0
			for {
				select {
				case <-done:
					return
				default:
					q.Put(i)
					i++
				}
			}
		}()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			q.Get()
		}
		done <- true
	})
}

// 测试并发时的数据一致性
func TestQueue_DataConsistency(t *testing.T) {
	q := NewQueue(10)

	// 使用通道来收集消费的数据
	results := make(chan int, 100)
	var wg sync.WaitGroup

	// 单个生产者生产有序数据
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			q.Put(i)
		}
	}()

	// 单个消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			item := q.Get()
			results <- item
		}
		close(results)
	}()

	wg.Wait()

	// 验证数据的完整性（不验证顺序，因为可能被重排）
	received := make(map[int]bool)
	for item := range results {
		if received[item] {
			t.Errorf("Received duplicate item: %d", item)
		}
		received[item] = true
	}

	// 验证所有数据都收到了
	if len(received) != 100 {
		t.Errorf("Expected 100 unique items, got %d", len(received))
	}

	for i := 0; i < 100; i++ {
		if !received[i] {
			t.Errorf("Item %d not received", i)
		}
	}
}
