package configmap

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/khulnasoft/tunnel-k8s-wrapper/data/config"
	"github.com/khulnasoft/tunnel-k8s-wrapper/kube"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ChainCreator can create chained configmaps
type ChainCreator struct {
	logger             *zap.Logger
	KubeClient         kube.Operations
	targetNamespace    string
	configmapNamespace string
	agentID            string
	tunnelVersion       string
	maxConfigSizeBytes int
}

// NewChainCreator returns a new ChainCreator
func NewChainCreator(c *config.Configuration, maxBytesPerConfigMap int) *ChainCreator {
	if maxBytesPerConfigMap == 0 {
		panic(fmt.Errorf("cannot handle 0Bytes per configmap"))
	}
	return &ChainCreator{
		logger:             c.Logger,
		KubeClient:         c.KubeClient,
		targetNamespace:    c.TargetNamespace,
		configmapNamespace: c.AgentNamespace,
		agentID:            c.AgentID,
		tunnelVersion:       c.TunnelScannerVersion,
		maxConfigSizeBytes: maxBytesPerConfigMap,
	}
}

// CreateChainedConfigMaps creates the chained configmaps by splitting the payload
func (c *ChainCreator) CreateChainedConfigMaps(ctx context.Context, payload []byte) error {
	i := 0
	configmapSequence := 0
	configmapsSpecs := []*corev1.ConfigMap{}
	for {
		var bytes []byte
		if i >= len(payload)-1 {
			// we have fetched all bytes from payload
			// we are done exit
			break
		} else if i < len(payload)-1 && (i+c.maxConfigSizeBytes) > len(payload)-1 {
			// fetch the remaining bytes
			bytes = payload[i:]
			i = len(payload)
		} else {
			// fetch maxConfigSizeBytes
			bytes = payload[i:(i + c.maxConfigSizeBytes)]
			i += c.maxConfigSizeBytes
		}
		// write the configmap
		configmapSequence++
		configmapsSpecs = append(configmapsSpecs, c.getConfigMapSpecs(bytes, configmapSequence))
	}

	// Remove the last configmap ocs_next label
	cm := configmapsSpecs[len(configmapsSpecs)-1]
	delete(cm.ObjectMeta.Labels, "agent.github.com/ocs-next")
	configmapsSpecs[len(configmapsSpecs)-1] = cm

	// deploy configmaps
	for _, cm := range configmapsSpecs {
		c.logger.Info("Creating configmap",
			zap.String("name", cm.ObjectMeta.GetName()),
			zap.String("configmap namespace", cm.GetObjectMeta().GetNamespace()))
		if err := c.KubeClient.DeployConfigmap(ctx, c.configmapNamespace, cm); err != nil {
			// We do not need to delete configmaps that we might have created at this point.
			// The github agent will do that for us.
			return fmt.Errorf("could not deploy configmap: %w", err)
		}
	}

	return nil
}

func (c *ChainCreator) getConfigMapSpecs(bytes []byte, seq int) *corev1.ConfigMap {
	name := fmt.Sprintf("ocs-%v-%v-%d", c.targetNamespace, c.agentID, seq)
	labels := map[string]string{
		"agent.github.com/scan":          "ocs",
		"agent.github.com/tunnel-version": c.tunnelVersion,
		"agent.github.com/ocs-ns":        c.targetNamespace,
		"agent.github.com/ocs-next":      fmt.Sprintf("ocs-%v-%v-%d", c.targetNamespace, c.agentID, seq+1),
		"agent.github.com/agent-id":      c.agentID,
	}
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: c.configmapNamespace,
			Labels:    labels,
		},
		BinaryData: map[string][]byte{
			"data": bytes,
		},
	}
}
