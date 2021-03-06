package main

import (
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	kubeClient kubernetes.Interface
	kubeConfig clientcmd.ClientConfig
	namespace  string
}

func main() {

	var client Client

	// initialize client-go clients
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	client.kubeConfig = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := client.kubeConfig.ClientConfig()
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	namespace, _, err := client.kubeConfig.Namespace()
	if err != nil {
		log.Fatal(err)
	}
	client.namespace = namespace

	deployApi := clientset.AppsV1()

	serviceApi := clientset.CoreV1()

	deployment, err := deployApi.Deployments(client.namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	service, err := serviceApi.Services(client.namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("In namespace %s on server %s\n\n", client.namespace, config.Host)
	for _, b := range deployment.Items {
		fmt.Printf("\tdeploy/%s deploys %s\n", b.Name, b.Spec.Template.Spec.Containers[0].Image)
	}
	fmt.Println("")

	for _, b := range service.Items {
		if b.Name == "kubernetes" {
			continue
		}
		fmt.Printf("\tsvc/%s is of type %s\n", b.Name, b.Spec.Type)
	}
	// fmt.Println("View details with 'kubectl describe <resource>/<name>' or list everything with 'kubectl get all'.")

}
