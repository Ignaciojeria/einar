package kubernetes

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func pod() {
	// The namespace is available under the service account secrets
	namespace, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err.Error())
	}

	// The pod name is available under the HOSTNAME environment variable
	podName := os.Getenv("HOSTNAME")
	if podName == "" {
		panic("HOSTNAME is empty")
	}

	// Creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// Creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Fetch the Pod object
	pod, err := clientset.CoreV1().Pods(string(namespace)).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Print the Pod UID
	fmt.Printf("Pod UID: %s\n", pod.UID)
}
