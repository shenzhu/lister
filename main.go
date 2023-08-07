package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/shenzhu/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// There're some default values for the config, such as config.QPS = 5, we are able to customize them
	config.Timeout = 120 * time.Second
	if err != nil {
		// handle error
		fmt.Printf("error building config from flags: %s\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			// handle error
			fmt.Printf("error getting incluster config: %s\n", err.Error())
		}
	}

	// A set of clients(for a bunch of groups and apiversions, but not clients for CRD and api services) to be used for interacting with the cluster
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// handle error
		fmt.Printf("error building kubernetes clientset: %s\n", err.Error())
	}

	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		// handle error
		fmt.Printf("error getting pods: %s\n", err.Error())
	}

	fmt.Println("Pods from default namespace")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		// handle error
		fmt.Printf("error getting deployments: %s\n", err.Error())
	}

	fmt.Println("Deployments from default namespace")
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}

	// Informers
	informerFactory := informers.NewSharedInformerFactory(clientset, 10*time.Minute)
	podinformer := informerFactory.Core().V1().Pods()
	podinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(new interface{}) {
			fmt.Println("add was called")
		},
		UpdateFunc: func(old, new interface{}) {
			fmt.Println("update was called")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete was called")
		},
	})

	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)
	pod, err := podinformer.Lister().Pods("default").Get("default")
	fmt.Println(pod)
}
