package validator

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/khulnasoft/tunnel-k8s-wrapper/kube"
)

// Flags is a flag validator that makes sure the provided values are correct
type Flags struct {
	GithubAgentNamespace string
	GithubAgentID        string
	Workloads            string
	Namespace            string
	KubeClient           kube.Operations
}

// Check makes sure that all flags are valid
func (f *Flags) Check(ctx context.Context) error {
	if err := f.validateGithubNamespace(ctx); err != nil {
		return err
	}
	if err := f.validateGithubAgentID(); err != nil {
		return err
	}
	if err := f.validateWorkloads(); err != nil {
		return err
	}
	return f.validateNamespace()
}

func (f *Flags) validateGithubNamespace(ctx context.Context) error {
	if _, err := f.KubeClient.GetNamespace(ctx, f.GithubAgentNamespace); err != nil {
		return fmt.Errorf("github-agent namespace %v is not correct: %w", f.GithubAgentNamespace, err)
	}
	return nil
}

func (f *Flags) validateGithubAgentID() error {
	if f.GithubAgentID == "" {
		return fmt.Errorf("github-agent-id cannot be empty")
	}
	return nil
}

func (f *Flags) validateWorkloads() error {
	values := strings.Split(f.Workloads, ",")
	allowedValues := map[string]bool{
		"pod":                   true,
		"replicaset":            true,
		"replicationcontroller": true,
		"statefulset":           true,
		"daemonset":             true,
		"cronjob":               true,
		"job":                   true,
		"deployment":            true,
	}
	// write code to make sure that values are inside allowedValues
	for _, v := range values {
		val := strings.ToLower(v)
		_, ok := allowedValues[val]
		if !ok {
			return fmt.Errorf("invalid workload type: %s", val)
		}
	}

	return nil
}

func (f *Flags) validateNamespace() error {
	if f.Namespace == "" {
		return errors.New("namespace cannot be empty")
	}
	return nil
}
