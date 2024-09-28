package kubernetes

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// ClientInterface, Kubernetes istemcisi için bir interface tanımlar
type ClientInterface interface {
	kubernetes.Interface
}

// NewClient, yeni bir Kubernetes istemcisi oluşturur
func NewClient(kubeconfigPath string) (ClientInterface, error) {
	config, err := getConfig(kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get Kubernetes config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	return clientset, nil
}

// getConfig, Kubernetes yapılandırmasını alır
func getConfig(kubeconfigPath string) (*rest.Config, error) {
	if kubeconfigPath == "" {
		kubeconfigPath = getDefaultKubeConfigPath()
	}

	// Küme içinde çalışıyorsa, in-cluster yapılandırmayı kullan
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	// Aksi takdirde, kubeconfig dosyasını kullan
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to build config from flags: %w", err)
	}

	return config, nil
}

// getDefaultKubeConfigPath, varsayılan kubeconfig dosya yolunu döndürür
func getDefaultKubeConfigPath() string {
	home := homedir.HomeDir()
	if home == "" {
		return ""
	}
	return filepath.Join(home, ".kube", "config")
}