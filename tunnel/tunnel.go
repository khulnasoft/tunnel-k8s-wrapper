package tunnel

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	// ReportFileName is the file name containing the vulnerability report
	ReportFileName = "result.json"
)

var errTunnelVersion = errors.New("tunnel version string is malformed")

// Scanner is a struct that is able to execute Tunnel related commands
type Scanner struct {
	Logger   *zap.Logger
	Executor Executor
}

// New returns a new instance of a tunnel scanner
func New(l *zap.Logger) *Scanner {
	return &Scanner{
		Logger:   l,
		Executor: &command{Logger: l},
	}
}

// Version returns the tunnel version
func (s Scanner) Version() (string, string, error) {
	args := []string{"-v"}
	out, stdErr, err := s.Executor.Exec(args)
	if err != nil {
		return "", stdErr, err
	}
	versionData := string(out)

	lines := strings.Split(versionData, "\n")
	if len(lines) < 1 {
		return "", stdErr, errTunnelVersion
	}

	match := regexp.MustCompile("Version: (.+)").FindStringSubmatch(lines[0])
	if len(match) < 2 {
		return "", stdErr, errTunnelVersion
	}

	return match[1], stdErr, nil
}

// Scan performs a tunnel K8S scan
func (s Scanner) Scan(workloads string, namespace string, timeout time.Duration, tunnelJavaDB string) (string, error) {
	args := []string{
		"k8s",
		"--include-kinds", workloads,
		"--report=all",
		"--scanners=vuln",
		// Node-collector is a scan job that collects node configuration parameters and permission information
		// This is not needed since we are only interested in scanning the workloads in the specified namespaces
		// More details https://khulnasoft.github.io/tunnel/v0.52/docs/target/kubernetes/
		"--disable-node-collector",
		"--db-repository", "registry.github.com/github-org/security-products/dependencies/tunnel-db-glad",
		"--include-namespaces", namespace,
		"--timeout", timeout.String(),
		"--java-db-repository", tunnelJavaDB,
		"--format", "json",
		"--output", ReportFileName,
	}
	_, stdErr, err := s.Executor.Exec(args)
	if err == nil {
		// Tunnel redirects stdOut and stdErr in stdErr
		s.Logger.Info("tunnel k8s command", zap.String("output", stdErr))
	}
	return stdErr, err
}

type command struct {
	Logger *zap.Logger
}

// Executor is an interface to interact with the command line
type Executor interface {
	Exec(args []string) ([]byte, string, error)
}

// Exec executes a tunnel command with the given arguments
func (c command) Exec(args []string) ([]byte, string, error) {
	var stdErr bytes.Buffer
	cmd := exec.Command("tunnel", args...)
	cmd.Stderr = &stdErr
	out, err := cmd.Output()
	c.Logger.Info("Executing", zap.String("cmd", cmd.String()))
	return out, stdErr.String(), err
}
