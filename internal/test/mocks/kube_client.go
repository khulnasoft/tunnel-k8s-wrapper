package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/khulnasoft/tunnel-k8s-wrapper/kube"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
)

// KubeOperations is a mock for kube.Operations interface
type KubeOperations struct {
	mock.Mock
	kube.Operations
	Cms []*corev1.ConfigMap
}

// GetNamespace is a mock for kube.Operations.GetNamespace
func (m *KubeOperations) GetNamespace(ctx context.Context, ns string) (*v1.Namespace, error) {
	a := m.Called(ctx, ns)
	return a.Get(0).(*v1.Namespace), a.Error(1)
}

// DeployConfigmap is a mock for kube.Operations.DeployConfigmap
// We capture the configmaps in an array so that we can check the values that we need
func (m *KubeOperations) DeployConfigmap(_ context.Context, _ string, cm *corev1.ConfigMap) error {
	m.Cms = append(m.Cms, cm)
	return nil
}
