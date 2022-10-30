package storage

import (
	"sync"

	"github.com/google/uuid"
)

type queue chan Messager

// NewMemory 内存模式
func NewMemoryQueue(poolNum uint) AdapterQueue {
	return &MemoryQueue{
		queue:   new(sync.Map),
		PoolNum: poolNum,
	}
}

type MemoryQueue struct {
	queue   *sync.Map
	wait    sync.WaitGroup
	mutex   sync.RWMutex
	PoolNum uint
}

func (*MemoryQueue) String() string {
	return "memory"
}

func (m *MemoryQueue) makeQueue() queue {
	if m.PoolNum <= 0 {
		return make(queue)
	}
	return make(queue, m.PoolNum)
}

func (m *MemoryQueue) Append(message Messager) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	memoryMessage := new(Message)
	memoryMessage.SetID(message.GetID())
	memoryMessage.SetStream(message.GetStream())
	memoryMessage.SetValues(message.GetValues())
	v, ok := m.queue.Load(message.GetStream())
	if !ok {
		v = m.makeQueue()
		m.queue.Store(message.GetStream(), v)
	}
	var q queue
	switch v.(type) {
	case queue:
		q = v.(queue)
	default:
		q = m.makeQueue()
		m.queue.Store(message.GetStream(), q)
	}
	go func(gm Messager, gq queue) {
		gm.SetID(uuid.New().String())
		gq <- gm
	}(memoryMessage, q)
	return nil
}

func (m *MemoryQueue) Register(name string, f ConsumerFunc) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	v, ok := m.queue.Load(name)
	if !ok {
		v = m.makeQueue()
		m.queue.Store(name, v)
	}
	var q queue
	switch v.(type) {
	case queue:
		q = v.(queue)
	default:
		q = m.makeQueue()
		m.queue.Store(name, q)
	}
	go func(out queue, gf ConsumerFunc) {
		var err error
		for message := range q {
			err = gf(message)
			if err != nil {
				out <- message
				err = nil
			}
		}
	}(q, f)
}

func (m *MemoryQueue) Run() {
	m.wait.Add(1)
	m.wait.Wait()
}

func (m *MemoryQueue) Shutdown() {
	m.wait.Done()
}
