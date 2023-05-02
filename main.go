package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/appscodepc/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create a Kubernetes clientset using the kubernetes.NewForConfig() function and use it to interact with the cluster.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Pods from default namespace")

	for _, pod := range pods.Items {
		fmt.Printf("%s ", pod.Name)
	}
	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("\nDeployment from default namespace")

	for _, d := range deployments.Items {
		fmt.Printf("%s ", d.Name)
	}
}
