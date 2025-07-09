package cond

import "sync"

// 生产者-消费者模式

type Queue struct {
	mu      sync.Mutex
	cond    *sync.Cond
	items   []int
	maxSize int
}

func NewQueue(maxSize int) *Queue {
	q := &Queue{maxSize: maxSize}
	q.cond = sync.NewCond(&q.mu)
	return q
}

// 生产者
func (q *Queue) Put(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// 等待队列有空间
	for len(q.items) >= q.maxSize {
		q.cond.Wait() // 释放锁并等待
	}

	q.items = append(q.items, item)
	q.cond.Signal() // 通知消费者
}

// 消费者
func (q *Queue) Get() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	// 等待队列有数据
	for len(q.items) == 0 {
		q.cond.Wait() // 释放锁并等待
	}

	item := q.items[0]
	q.items = q.items[1:]
	q.cond.Signal() // 通知生产者
	return item
}
