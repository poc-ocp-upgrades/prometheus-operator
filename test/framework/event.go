package framework

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) PrintEvents() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	events, err := f.KubeClient.CoreV1().Events("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	if events != nil {
		fmt.Println("=== Kubernetes events:")
		for _, e := range events.Items {
			fmt.Printf("FirstTimestamp: '%v', Reason: '%v', Message: '%v'\n", e.FirstTimestamp, e.Reason, e.Message)
		}
	}
	return nil
}
