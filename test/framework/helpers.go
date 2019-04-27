package framework

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"github.com/coreos/prometheus-operator/pkg/k8sutil"
	"github.com/pkg/errors"
)

func PathToOSFile(relativPath string) (*os.File, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	path, err := filepath.Abs(relativPath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed generate absolut file path of %s", relativPath))
	}
	manifest, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to open file %s", path))
	}
	return manifest, nil
}
func WaitForPodsReady(kubeClient kubernetes.Interface, namespace string, timeout time.Duration, expectedReplicas int, opts metav1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.Poll(time.Second, timeout, func() (bool, error) {
		pl, err := kubeClient.Core().Pods(namespace).List(opts)
		if err != nil {
			return false, err
		}
		runningAndReady := 0
		for _, p := range pl.Items {
			isRunningAndReady, err := k8sutil.PodRunningAndReady(p)
			if err != nil {
				return false, err
			}
			if isRunningAndReady {
				runningAndReady++
			}
		}
		if runningAndReady == expectedReplicas {
			return true, nil
		}
		return false, nil
	})
}
func WaitForPodsRunImage(kubeClient kubernetes.Interface, namespace string, expectedReplicas int, image string, opts metav1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.Poll(time.Second, time.Minute*5, func() (bool, error) {
		pl, err := kubeClient.Core().Pods(namespace).List(opts)
		if err != nil {
			return false, err
		}
		runningImage := 0
		for _, p := range pl.Items {
			if podRunsImage(p, image) {
				runningImage++
			}
		}
		if runningImage == expectedReplicas {
			return true, nil
		}
		return false, nil
	})
}
func WaitForHTTPSuccessStatusCode(timeout time.Duration, url string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var resp *http.Response
	err := wait.Poll(time.Second, timeout, func() (bool, error) {
		var err error
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return true, nil
		}
		return false, nil
	})
	return errors.Wrap(err, fmt.Sprintf("waiting for %v to return a successful status code timed out. Last response from server was: %v", url, resp))
}
func podRunsImage(p v1.Pod, image string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, c := range p.Spec.Containers {
		if image == c.Image {
			return true
		}
	}
	return false
}
func GetLogs(kubeClient kubernetes.Interface, namespace string, podName, containerName string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logs, err := kubeClient.Core().RESTClient().Get().Resource("pods").Namespace(namespace).Name(podName).SubResource("log").Param("container", containerName).Do().Raw()
	if err != nil {
		return "", err
	}
	return string(logs), err
}
func (f *Framework) Poll(timeout, pollInterval time.Duration, pollFunc func() (bool, error)) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := time.After(timeout)
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-t:
			return fmt.Errorf("timed out")
		case <-ticker.C:
			b, err := pollFunc()
			if err != nil {
				return err
			}
			if b {
				return nil
			}
		}
	}
}
func ProxyGetPod(kubeClient kubernetes.Interface, namespace, podName, port, path string) *rest.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kubeClient.CoreV1().RESTClient().Get().Namespace(namespace).Resource("pods").SubResource("proxy").Name(podName).Suffix(path)
}
func ProxyPostPod(kubeClient kubernetes.Interface, namespace, podName, port, path, body string) *rest.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kubeClient.CoreV1().RESTClient().Post().Namespace(namespace).Resource("pods").SubResource("proxy").Name(podName).Suffix(path).Body([]byte(body)).SetHeader("Content-Type", "application/json")
}
