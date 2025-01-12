package config

import (
	"go.uber.org/zap"

	"github.com/khulnasoft/tunnel-k8s-wrapper/kube"
)

// Configuration holds data that are required to initialize various structs
type Configuration struct {
	Logger              *zap.Logger
	KubeClient          kube.Operations
	TargetNamespace     string
	AgentID             string
	AgentNamespace      string
	TunnelScannerVersion string
}
