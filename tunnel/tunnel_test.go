package tunnel_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/tunnel-k8s-wrapper/internal/test/mocks"
	"github.com/khulnasoft/tunnel-k8s-wrapper/logging"
	"github.com/khulnasoft/tunnel-k8s-wrapper/tunnel"
)

func TestVersion(t *testing.T) {
	versionData := "Version: 0.45.1\nVulnerability DB:\n  Version: 2\n  UpdatedAt: 2023-12-13 04:23:34.686016965 +0000 UTC\n  NextUpdate: 2023-12-13 10:23:34.686016425 +0000 UTC\n  DownloadedAt: 2023-12-13 13:36:33.466336 +0000 UTC\n"
	l := logging.NewLogger("warn")
	mockExec := &mocks.Executor{}
	mockExec.On("Exec", []string{"-v"}).Times(1).Return([]byte(versionData), "something", nil)
	s := tunnel.Scanner{
		Logger:   l,
		Executor: mockExec,
	}

	version, stdErr, err := s.Version()
	assert.Equal(t, "0.45.1", version)
	assert.Equal(t, "something", stdErr)
	assert.Nil(t, err)
	mockExec.AssertExpectations(t)
}

func TestScan(t *testing.T) {
	workloads := "pod,replicaset"
	namespace := "default"
	timeout := time.Minute * 10
	tunnelJavaDB := "registry.github.com/github-org/security-products/dependencies/tunnel-java-db:1,ghcr.io/khulnasoft/tunnel-java-db:1"
	l := logging.NewLogger("warn")
	mockExec := &mocks.Executor{}
	mockExec.On("Exec", []string{
		"k8s",
		"--include-kinds",
		workloads,
		"--report=all",
		"--scanners=vuln",
		"--disable-node-collector",
		"--db-repository",
		"registry.github.com/github-org/security-products/dependencies/tunnel-db-glad",
		"--include-namespaces",
		namespace,
		"--timeout",
		timeout.String(),
		"--java-db-repository", tunnelJavaDB,
		"--format",
		"json",
		"--output",
		"result.json",
	}).
		Times(1).Return([]byte("something"), "something", nil)
	s := tunnel.Scanner{
		Logger:   l,
		Executor: mockExec,
	}

	stdErr, err := s.Scan(workloads, namespace, timeout, tunnelJavaDB)
	assert.Equal(t, "something", stdErr)
	assert.Nil(t, err)
	mockExec.AssertExpectations(t)
}
