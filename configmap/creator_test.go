package configmap_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/tunnel-k8s-wrapper/configmap"
	"github.com/khulnasoft/tunnel-k8s-wrapper/data/config"
	"github.com/khulnasoft/tunnel-k8s-wrapper/internal/test/mocks"
	"github.com/khulnasoft/tunnel-k8s-wrapper/logging"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateChainedConfigMaps(t *testing.T) {
	mockKubeClient := &mocks.KubeOperations{}

	// Our payload is a byte array with values from 1 to 30
	bytes := make([]byte, 30)
	for i := 1; i < len(bytes)+1; i++ {
		bytes[i-1] = byte(i)
	}

	agentID := "1234"
	tunnelVersion := "v0.455"
	targetNamespace := "kube-system"
	agentNamespace := "agent-namespace"
	// We expect to divide the payload in 3 parts
	expectedCms := []*corev1.ConfigMap{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ocs-kube-system-1234-1",
				Namespace: agentNamespace,
				Labels: map[string]string{
					"agent.github.com/scan":          "ocs",
					"agent.github.com/tunnel-version": tunnelVersion,
					"agent.github.com/ocs-ns":        targetNamespace,
					"agent.github.com/ocs-next":      "ocs-kube-system-1234-2",
					"agent.github.com/agent-id":      agentID,
				},
			},
			BinaryData: map[string][]byte{
				"data": bytes[0:10],
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ocs-kube-system-1234-2",
				Namespace: agentNamespace,
				Labels: map[string]string{
					"agent.github.com/scan":          "ocs",
					"agent.github.com/tunnel-version": tunnelVersion,
					"agent.github.com/ocs-ns":        targetNamespace,
					"agent.github.com/ocs-next":      "ocs-kube-system-1234-3",
					"agent.github.com/agent-id":      agentID,
				},
			},
			BinaryData: map[string][]byte{
				"data": bytes[10:20],
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ocs-kube-system-1234-3",
				Namespace: agentNamespace,
				Labels: map[string]string{
					"agent.github.com/scan":          "ocs",
					"agent.github.com/tunnel-version": tunnelVersion,
					"agent.github.com/ocs-ns":        targetNamespace,
					"agent.github.com/agent-id":      agentID,
				},
			},
			BinaryData: map[string][]byte{
				"data": bytes[20:30],
			},
		},
	}

	// We pass as maxBytesPerConfigMap 10 for a total payload of 30 bytes
	cr := configmap.NewChainCreator(&config.Configuration{
		TargetNamespace:     targetNamespace,
		AgentID:             agentID,
		AgentNamespace:      agentNamespace,
		TunnelScannerVersion: tunnelVersion,
		Logger:              logging.NewLogger("warn"),
		KubeClient:          mockKubeClient,
	}, 10)

	assert.Nil(t, cr.CreateChainedConfigMaps(context.Background(), bytes))
	assert.Equal(t, expectedCms, mockKubeClient.Cms)
}
