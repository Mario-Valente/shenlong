package k8s

import (
	"errors"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func CreateClientK8s(pathKubeconfig string) (*kubernetes.Clientset, error) {
	var kubeconfig string

	home := homedir.HomeDir()
	if home != "" {
		kubeconfig = filepath.Join(home, pathKubeconfig, ".kube", "config")
	} else {
		return nil, errors.New("home directory not found")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error building kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating clientset: %v", err)
	}

	return clientset, nil
}
