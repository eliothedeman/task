package task

import (
	"sync"
)

// Manager keeps track of tasks
type Manager struct {
	wg sync.WaitGroup
}

// Run a function in the manager
func (m *Manager) Run(f func(), then ...func()) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		f()
		for _, x := range then {
			x()
		}
	}()
}

// RunN runs `n` copies of a functiona
func (m *Manager) RunN(f func(), n int, then ...func()) {
	for i := 0; i < n; i++ {
		if i == 0 {
			m.Run(f, then...)
		}
		m.Run(f)
	}
}

// Wait until all tasks have completed
func (m *Manager) Wait() {
	m.wg.Wait()
}
