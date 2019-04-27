package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	scheme "github.com/coreos/prometheus-operator/pkg/client/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type PrometheusesGetter interface {
	Prometheuses(namespace string) PrometheusInterface
}
type PrometheusInterface interface {
	Create(*v1.Prometheus) (*v1.Prometheus, error)
	Update(*v1.Prometheus) (*v1.Prometheus, error)
	UpdateStatus(*v1.Prometheus) (*v1.Prometheus, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Prometheus, error)
	List(opts meta_v1.ListOptions) (*v1.PrometheusList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Prometheus, err error)
	PrometheusExpansion
}
type prometheuses struct {
	client	rest.Interface
	ns	string
}

func newPrometheuses(c *MonitoringV1Client, namespace string) *prometheuses {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &prometheuses{client: c.RESTClient(), ns: namespace}
}
func (c *prometheuses) Get(name string, options meta_v1.GetOptions) (result *v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.Prometheus{}
	err = c.client.Get().Namespace(c.ns).Resource("prometheuses").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *prometheuses) List(opts meta_v1.ListOptions) (result *v1.PrometheusList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.PrometheusList{}
	err = c.client.Get().Namespace(c.ns).Resource("prometheuses").VersionedParams(&opts, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *prometheuses) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("prometheuses").VersionedParams(&opts, scheme.ParameterCodec).Watch()
}
func (c *prometheuses) Create(prometheus *v1.Prometheus) (result *v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.Prometheus{}
	err = c.client.Post().Namespace(c.ns).Resource("prometheuses").Body(prometheus).Do().Into(result)
	return
}
func (c *prometheuses) Update(prometheus *v1.Prometheus) (result *v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.Prometheus{}
	err = c.client.Put().Namespace(c.ns).Resource("prometheuses").Name(prometheus.Name).Body(prometheus).Do().Into(result)
	return
}
func (c *prometheuses) UpdateStatus(prometheus *v1.Prometheus) (result *v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.Prometheus{}
	err = c.client.Put().Namespace(c.ns).Resource("prometheuses").Name(prometheus.Name).SubResource("status").Body(prometheus).Do().Into(result)
	return
}
func (c *prometheuses) Delete(name string, options *meta_v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("prometheuses").Name(name).Body(options).Do().Error()
}
func (c *prometheuses) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("prometheuses").VersionedParams(&listOptions, scheme.ParameterCodec).Body(options).Do().Error()
}
func (c *prometheuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.Prometheus{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("prometheuses").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
