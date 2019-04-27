package v1

import (
	internalinterfaces "github.com/coreos/prometheus-operator/pkg/client/informers/externalversions/internalinterfaces"
)

type Interface interface {
	Alertmanagers() AlertmanagerInformer
	Prometheuses() PrometheusInformer
	PrometheusRules() PrometheusRuleInformer
	ServiceMonitors() ServiceMonitorInformer
}
type version struct {
	factory			internalinterfaces.SharedInformerFactory
	namespace		string
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
}

func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}
func (v *version) Alertmanagers() AlertmanagerInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &alertmanagerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
func (v *version) Prometheuses() PrometheusInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &prometheusInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
func (v *version) PrometheusRules() PrometheusRuleInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &prometheusRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
func (v *version) ServiceMonitors() ServiceMonitorInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceMonitorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
