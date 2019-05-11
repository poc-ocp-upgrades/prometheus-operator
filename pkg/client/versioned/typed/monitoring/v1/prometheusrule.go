package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	scheme "github.com/coreos/prometheus-operator/pkg/client/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type PrometheusRulesGetter interface {
	PrometheusRules(namespace string) PrometheusRuleInterface
}
type PrometheusRuleInterface interface {
	Create(*v1.PrometheusRule) (*v1.PrometheusRule, error)
	Update(*v1.PrometheusRule) (*v1.PrometheusRule, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.PrometheusRule, error)
	List(opts meta_v1.ListOptions) (*v1.PrometheusRuleList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PrometheusRule, err error)
	PrometheusRuleExpansion
}
type prometheusRules struct {
	client	rest.Interface
	ns		string
}

func newPrometheusRules(c *MonitoringV1Client, namespace string) *prometheusRules {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &prometheusRules{client: c.RESTClient(), ns: namespace}
}
func (c *prometheusRules) Get(name string, options meta_v1.GetOptions) (result *v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.PrometheusRule{}
	err = c.client.Get().Namespace(c.ns).Resource("prometheusrules").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *prometheusRules) List(opts meta_v1.ListOptions) (result *v1.PrometheusRuleList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.PrometheusRuleList{}
	err = c.client.Get().Namespace(c.ns).Resource("prometheusrules").VersionedParams(&opts, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *prometheusRules) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("prometheusrules").VersionedParams(&opts, scheme.ParameterCodec).Watch()
}
func (c *prometheusRules) Create(prometheusRule *v1.PrometheusRule) (result *v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.PrometheusRule{}
	err = c.client.Post().Namespace(c.ns).Resource("prometheusrules").Body(prometheusRule).Do().Into(result)
	return
}
func (c *prometheusRules) Update(prometheusRule *v1.PrometheusRule) (result *v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.PrometheusRule{}
	err = c.client.Put().Namespace(c.ns).Resource("prometheusrules").Name(prometheusRule.Name).Body(prometheusRule).Do().Into(result)
	return
}
func (c *prometheusRules) Delete(name string, options *meta_v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("prometheusrules").Name(name).Body(options).Do().Error()
}
func (c *prometheusRules) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("prometheusrules").VersionedParams(&listOptions, scheme.ParameterCodec).Body(options).Do().Error()
}
func (c *prometheusRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.PrometheusRule{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("prometheusrules").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
