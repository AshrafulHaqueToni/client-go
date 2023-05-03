package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("error %s getting inclusterconfig\n", err.Error())
	}

	// create a Kubernetes clientset using the kubernetes.NewForConfig() function and use it to interact with the cluster.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, creating clientset\n", err.Error())
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s, while listing all the pods from default namespace\n", err.Error())
	}
	fmt.Println("Pods from default namespace")

	for _, pod := range pods.Items {
		fmt.Printf("%s ", pod.Name)
	}
	_, err = clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("listing deployments %s\n", err.Error())
	}
	fmt.Println("\nDeployment from default namespace")

	for {
	}
}
