package validator_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"

	"github.com/khulnasoft/tunnel-k8s-wrapper/internal/test/mocks"
	"github.com/khulnasoft/tunnel-k8s-wrapper/validator"
)

func TestCheck(t *testing.T) {
	mockKubeClient := &mocks.KubeOperations{}
	var tcs = []struct {
		name     string
		expError error
		before   func()
		flags    validator.Flags
		after    func()
	}{
		{
			name:     "Fails to validate github-namespace",
			expError: errors.New("github-agent namespace github-agent is not correct: mock"),
			before: func() {
				var nilNs *v1.Namespace
				mockKubeClient.On("GetNamespace", mock.Anything, "github-agent").
					Times(1).Return(nilNs, errors.New("mock"))
			},
			flags: validator.Flags{
				KubeClient:           mockKubeClient,
				GithubAgentNamespace: "github-agent",
			},
			after: func() {
				mockKubeClient.AssertExpectations(t)
			},
		},
		{
			name:     "Fails to validate github agent ID",
			expError: errors.New("github-agent-id cannot be empty"),
			before: func() {
				mockKubeClient.On("GetNamespace", mock.Anything, "github-agent").
					Times(1).Return(&v1.Namespace{}, nil)
			},
			flags: validator.Flags{
				KubeClient:           mockKubeClient,
				GithubAgentNamespace: "github-agent",
			},
			after: func() {
				mockKubeClient.AssertExpectations(t)
			},
		},
		{
			name:     "Fails to validate workloads",
			expError: errors.New("invalid workload type: something wrong"),
			before: func() {
				mockKubeClient.On("GetNamespace", mock.Anything, "github-agent").
					Times(1).Return(&v1.Namespace{}, nil)
			},
			flags: validator.Flags{
				KubeClient:           mockKubeClient,
				GithubAgentID:        "1",
				GithubAgentNamespace: "github-agent",
				Workloads:            "Pod,something wrong",
			},
			after: func() {
				mockKubeClient.AssertExpectations(t)
			},
		},
		{
			name:     "Fails to validate namespace",
			expError: errors.New("namespace cannot be empty"),
			before: func() {
				mockKubeClient.On("GetNamespace", mock.Anything, "github-agent").
					Times(1).Return(&v1.Namespace{}, nil)
			},
			flags: validator.Flags{
				KubeClient:           mockKubeClient,
				GithubAgentNamespace: "github-agent",
				GithubAgentID:        "1",
				Workloads:            "Pod,ReplicaSet",
			},
			after: func() {
				mockKubeClient.AssertExpectations(t)
			},
		},
		{
			name:     "Validates correctly",
			expError: nil,
			before: func() {
				mockKubeClient.On("GetNamespace", mock.Anything, "github-agent").
					Times(1).Return(&v1.Namespace{}, nil)
			},
			flags: validator.Flags{
				KubeClient:           mockKubeClient,
				GithubAgentNamespace: "github-agent",
				GithubAgentID:        "1",
				Workloads:            "Pod,ReplicaSet",
				Namespace:            "default",
			},
			after: func() {
				mockKubeClient.AssertExpectations(t)
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tc.before()
			if tc.expError != nil {
				assert.Equal(t, tc.expError.Error(), tc.flags.Check(context.Background()).Error())
			} else {
				assert.Nil(t, tc.flags.Check(context.Background()))
			}
			tc.after()
		})
	}
}
