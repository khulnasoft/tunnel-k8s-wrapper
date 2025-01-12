package kube

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/cmd/util"
)

// Client enable us to interact with a Kubernetes cluster
type Client struct {
	k kubernetes.Interface
}

// Operations gathers all the operations we can perform on a Kubernetes cluster
type Operations interface {
	GetNamespace(ctx context.Context, ns string) (*v1.Namespace, error)
	DeployConfigmap(ctx context.Context, namespace string, cm *corev1.ConfigMap) error
}

// NewClient returns a new Operations interface
func NewClient() (Operations, error) {
	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	k8sFactory := util.NewFactory(kubeConfigFlags)
	kubeClient, err := k8sFactory.KubernetesClientSet()
	if err != nil {
		return nil, err
	}

	return &Client{
		k: kubeClient,
	}, nil
}

// GetNamespace returns a namespace object based on a given name
func (c *Client) GetNamespace(ctx context.Context, ns string) (*v1.Namespace, error) {
	return c.k.CoreV1().Namespaces().Get(ctx, ns, metav1.GetOptions{})
}

// DeployConfigmap deploys a configmap on a given namespace. Be aware that when this method returns the
// configmap might be still under construction
func (c *Client) DeployConfigmap(ctx context.Context, namespace string, cm *corev1.ConfigMap) error {
	_, err := c.k.CoreV1().ConfigMaps(namespace).Create(ctx, cm, metav1.CreateOptions{})
	return err
}
