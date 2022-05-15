package main

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func main() {
	// uses the current context in kubeconfig
	// config, _ := clientcmd.BuildConfigFromFlags("", "./kubeconfig")
	// fmt.Println(config.String())

	config := rest.Config{
		Host:        "https://172.20.0.1/k8s/clusters/c-f756c",
		APIPath:     "",
		BearerToken: "kubeconfig-u-gz9nckvsh4:mh8bldcbxs6cbbcx9bcmxgqn2f2rjc85xddbkv6bk67zhrq7dpdwpf",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	// creates the clientset
	clientset, _ := kubernetes.NewForConfig(&config)
	metricsclientset, _ := metricsv.NewForConfig(&config)
	// access the API to list pods
	pods, _ := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	ff, _ := json.Marshal(pods)
	fmt.Println(string(ff))
	podMetrics, _ := metricsclientset.MetricsV1beta1().NodeMetricses().List(context.TODO(), v1.ListOptions{})
	f, _ := json.Marshal(podMetrics)
	fmt.Println(string(f))

}
