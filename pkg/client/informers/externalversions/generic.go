package externalversions

import (
	"fmt"
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}
type genericInformer struct {
	informer	cache.SharedIndexInformer
	resource	schema.GroupResource
}

func (f *genericInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.informer
}
func (f *genericInformer) Lister() cache.GenericLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch resource {
	case v1.SchemeGroupVersion.WithResource("alertmanagers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().Alertmanagers().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("prometheuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().Prometheuses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("prometheusrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().PrometheusRules().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("servicemonitors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Monitoring().V1().ServiceMonitors().Informer()}, nil
	}
	return nil, fmt.Errorf("no informer found for %v", resource)
}
