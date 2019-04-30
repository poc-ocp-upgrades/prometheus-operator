package v1

import (
	time "time"
	monitoring_v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	internalinterfaces "github.com/coreos/prometheus-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/coreos/prometheus-operator/pkg/client/listers/monitoring/v1"
	versioned "github.com/coreos/prometheus-operator/pkg/client/versioned"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

type PrometheusRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PrometheusRuleLister
}
type prometheusRuleInformer struct {
	factory			internalinterfaces.SharedInformerFactory
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
	namespace		string
}

func NewPrometheusRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredPrometheusRuleInformer(client, namespace, resyncPeriod, indexers, nil)
}
func NewFilteredPrometheusRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewSharedIndexInformer(&cache.ListWatch{ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.MonitoringV1().PrometheusRules(namespace).List(options)
	}, WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.MonitoringV1().PrometheusRules(namespace).Watch(options)
	}}, &monitoring_v1.PrometheusRule{}, resyncPeriod, indexers)
}
func (f *prometheusRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredPrometheusRuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}
func (f *prometheusRuleInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.factory.InformerFor(&monitoring_v1.PrometheusRule{}, f.defaultInformer)
}
func (f *prometheusRuleInformer) Lister() v1.PrometheusRuleLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v1.NewPrometheusRuleLister(f.Informer().GetIndexer())
}
