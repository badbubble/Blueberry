package k8s

import (
	"Blueberry/pkg/setting"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

var Client *kubernetes.Clientset

func Init() (err error) {
	config, err := clientcmd.BuildConfigFromFlags("", setting.Conf.KubeConfig)
	if err != nil {
		fmt.Printf("clientcmd.BuildConfigFromFlags error: %v", err)
		return
	}

	Client, err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Unable to create k8s client: %v", err)
		return
	}
	return nil
}
