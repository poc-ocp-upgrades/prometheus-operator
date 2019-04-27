package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type PrometheusLister interface {
	List(selector labels.Selector) (ret []*v1.Prometheus, err error)
	Prometheuses(namespace string) PrometheusNamespaceLister
	PrometheusListerExpansion
}
type prometheusLister struct{ indexer cache.Indexer }

func NewPrometheusLister(indexer cache.Indexer) PrometheusLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &prometheusLister{indexer: indexer}
}
func (s *prometheusLister) List(selector labels.Selector) (ret []*v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Prometheus))
	})
	return ret, err
}
func (s *prometheusLister) Prometheuses(namespace string) PrometheusNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return prometheusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type PrometheusNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1.Prometheus, err error)
	Get(name string) (*v1.Prometheus, error)
	PrometheusNamespaceListerExpansion
}
type prometheusNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s prometheusNamespaceLister) List(selector labels.Selector) (ret []*v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Prometheus))
	})
	return ret, err
}
func (s prometheusNamespaceLister) Get(name string) (*v1.Prometheus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("prometheus"), name)
	}
	return obj.(*v1.Prometheus), nil
}
