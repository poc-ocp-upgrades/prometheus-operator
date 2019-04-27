package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type PrometheusRuleLister interface {
	List(selector labels.Selector) (ret []*v1.PrometheusRule, err error)
	PrometheusRules(namespace string) PrometheusRuleNamespaceLister
	PrometheusRuleListerExpansion
}
type prometheusRuleLister struct{ indexer cache.Indexer }

func NewPrometheusRuleLister(indexer cache.Indexer) PrometheusRuleLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &prometheusRuleLister{indexer: indexer}
}
func (s *prometheusRuleLister) List(selector labels.Selector) (ret []*v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PrometheusRule))
	})
	return ret, err
}
func (s *prometheusRuleLister) PrometheusRules(namespace string) PrometheusRuleNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return prometheusRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type PrometheusRuleNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1.PrometheusRule, err error)
	Get(name string) (*v1.PrometheusRule, error)
	PrometheusRuleNamespaceListerExpansion
}
type prometheusRuleNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s prometheusRuleNamespaceLister) List(selector labels.Selector) (ret []*v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PrometheusRule))
	})
	return ret, err
}
func (s prometheusRuleNamespaceLister) Get(name string) (*v1.PrometheusRule, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("prometheusrule"), name)
	}
	return obj.(*v1.PrometheusRule), nil
}
