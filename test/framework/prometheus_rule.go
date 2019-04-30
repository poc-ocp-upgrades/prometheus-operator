package framework

import (
	"fmt"
	"time"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/wait"
)

func (f *Framework) MakeBasicRule(ns, name string, groups []monitoringv1.RuleGroup) *monitoringv1.PrometheusRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &monitoringv1.PrometheusRule{ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns, Labels: map[string]string{"role": "rulefile"}}, Spec: monitoringv1.PrometheusRuleSpec{Groups: groups}}
}
func (f *Framework) CreateRule(ns string, ar *monitoringv1.PrometheusRule) (*monitoringv1.PrometheusRule, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result, err := f.MonClientV1.PrometheusRules(ns).Create(ar)
	if err != nil {
		return nil, fmt.Errorf("creating %v RuleFile failed: %v", ar.Name, err)
	}
	return result, nil
}
func (f *Framework) MakeAndCreateFiringRule(ns, name, alertName string) (*monitoringv1.PrometheusRule, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	groups := []monitoringv1.RuleGroup{{Name: alertName, Rules: []monitoringv1.Rule{{Alert: alertName, Expr: intstr.FromString("vector(1)")}}}}
	file := f.MakeBasicRule(ns, name, groups)
	result, err := f.CreateRule(ns, file)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (f *Framework) WaitForRule(ns, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.Poll(time.Second, f.DefaultTimeout, func() (bool, error) {
		_, err := f.MonClientV1.PrometheusRules(ns).Get(name, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return false, nil
		} else if err != nil {
			return false, err
		}
		return true, nil
	})
}
func (f *Framework) UpdateRule(ns string, ar *monitoringv1.PrometheusRule) (*monitoringv1.PrometheusRule, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result, err := f.MonClientV1.PrometheusRules(ns).Update(ar)
	if err != nil {
		return nil, fmt.Errorf("updating %v RuleFile failed: %v", ar.Name, err)
	}
	return result, nil
}
func (f *Framework) DeleteRule(ns string, r string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := f.MonClientV1.PrometheusRules(ns).Delete(r, &metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("deleteing %v Prometheus rule in namespace %v failed: %v", r, ns, err.Error())
	}
	return nil
}
