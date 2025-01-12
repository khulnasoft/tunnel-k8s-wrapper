package converter_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/khulnasoft/tunnel-k8s-wrapper/converter"
	"github.com/khulnasoft/tunnel-k8s-wrapper/data/config"
	"github.com/khulnasoft/tunnel-k8s-wrapper/data/kas"
	"github.com/khulnasoft/tunnel-k8s-wrapper/internal/test/mocks"
	"github.com/khulnasoft/tunnel-k8s-wrapper/logging"
)

func TestPrepareData(t *testing.T) {
	mockKubeClient := &mocks.KubeOperations{}
	tunnelRepBytes, err := os.ReadFile("fixtures/tunnel_report.json")
	assert.Nil(t, err)
	var report kas.Report
	require.NoError(t, json.Unmarshal(tunnelRepBytes, &report))
	expectedBase64Bytes, err := os.ReadFile("expect/base64.bytes")
	assert.Nil(t, err)

	config := config.Configuration{
		TargetNamespace:     "kube-system",
		AgentID:             "1234",
		TunnelScannerVersion: "v0.455",
		Logger:              logging.NewLogger("warn"),
		KubeClient:          mockKubeClient,
	}
	dt := converter.NewReportConverter(&config, &report)

	actualBytes, err := dt.PrepareData()
	assert.Nil(t, err)
	assert.Equal(t, expectedBase64Bytes, actualBytes)
	mockKubeClient.AssertExpectations(t)

}
