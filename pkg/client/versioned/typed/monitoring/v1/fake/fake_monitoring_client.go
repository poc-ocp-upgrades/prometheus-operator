package fake

import (
	v1 "github.com/coreos/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMonitoringV1 struct{ *testing.Fake }

func (c *FakeMonitoringV1) Alertmanagers(namespace string) v1.AlertmanagerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeAlertmanagers{c, namespace}
}
func (c *FakeMonitoringV1) Prometheuses(namespace string) v1.PrometheusInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakePrometheuses{c, namespace}
}
func (c *FakeMonitoringV1) PrometheusRules(namespace string) v1.PrometheusRuleInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakePrometheusRules{c, namespace}
}
func (c *FakeMonitoringV1) ServiceMonitors(namespace string) v1.ServiceMonitorInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceMonitors{c, namespace}
}
func (c *FakeMonitoringV1) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}
