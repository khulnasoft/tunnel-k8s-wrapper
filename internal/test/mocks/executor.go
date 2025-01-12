package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/khulnasoft/tunnel-k8s-wrapper/tunnel"
)

// Executor is a mock for the tunnel Executor interface
type Executor struct {
	mock.Mock
	tunnel.Executor
}

// Exec is a mock for tunnel.Executor.Exec
func (m *Executor) Exec(args []string) ([]byte, string, error) {
	a := m.Called(args)
	return a.Get(0).([]byte), a.String(1), a.Error(2)
}
