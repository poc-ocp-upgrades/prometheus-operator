package framework

import (
	"time"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"github.com/pkg/errors"
)

func (f *Framework) WaitForConfigMapExist(ns, name string) (*v1.ConfigMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var configMap *v1.ConfigMap
	err := wait.Poll(2*time.Second, f.DefaultTimeout, func() (bool, error) {
		var err error
		configMap, err = f.KubeClient.CoreV1().ConfigMaps(ns).Get(name, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		return true, nil
	})
	return configMap, errors.Wrapf(err, "waiting for ConfigMap '%v' in namespace '%v'", name, ns)
}
func (f *Framework) WaitForConfigMapNotExist(ns, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := wait.Poll(2*time.Second, f.DefaultTimeout, func() (bool, error) {
		var err error
		_, err = f.KubeClient.CoreV1().ConfigMaps(ns).Get(name, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return true, nil
		}
		if err != nil {
			return false, err
		}
		return false, nil
	})
	return errors.Wrapf(err, "waiting for ConfigMap '%v' in namespace '%v' to not exist", name, ns)
}
