package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServiceMonitorLister interface {
	List(selector labels.Selector) (ret []*v1.ServiceMonitor, err error)
	ServiceMonitors(namespace string) ServiceMonitorNamespaceLister
	ServiceMonitorListerExpansion
}
type serviceMonitorLister struct{ indexer cache.Indexer }

func NewServiceMonitorLister(indexer cache.Indexer) ServiceMonitorLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceMonitorLister{indexer: indexer}
}
func (s *serviceMonitorLister) List(selector labels.Selector) (ret []*v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ServiceMonitor))
	})
	return ret, err
}
func (s *serviceMonitorLister) ServiceMonitors(namespace string) ServiceMonitorNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return serviceMonitorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type ServiceMonitorNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1.ServiceMonitor, err error)
	Get(name string) (*v1.ServiceMonitor, error)
	ServiceMonitorNamespaceListerExpansion
}
type serviceMonitorNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s serviceMonitorNamespaceLister) List(selector labels.Selector) (ret []*v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ServiceMonitor))
	})
	return ret, err
}
func (s serviceMonitorNamespaceLister) Get(name string) (*v1.ServiceMonitor, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("servicemonitor"), name)
	}
	return obj.(*v1.ServiceMonitor), nil
}
