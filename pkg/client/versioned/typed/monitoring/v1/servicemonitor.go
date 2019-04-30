package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	scheme "github.com/coreos/prometheus-operator/pkg/client/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type ServiceMonitorsGetter interface {
	ServiceMonitors(namespace string) ServiceMonitorInterface
}
type ServiceMonitorInterface interface {
	Create(*v1.ServiceMonitor) (*v1.ServiceMonitor, error)
	Update(*v1.ServiceMonitor) (*v1.ServiceMonitor, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.ServiceMonitor, error)
	List(opts meta_v1.ListOptions) (*v1.ServiceMonitorList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ServiceMonitor, err error)
	ServiceMonitorExpansion
}
type serviceMonitors struct {
	client	rest.Interface
	ns	string
}

func newServiceMonitors(c *MonitoringV1Client, namespace string) *serviceMonitors {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceMonitors{client: c.RESTClient(), ns: namespace}
}
func (c *serviceMonitors) Get(name string, options meta_v1.GetOptions) (result *v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ServiceMonitor{}
	err = c.client.Get().Namespace(c.ns).Resource("servicemonitors").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceMonitors) List(opts meta_v1.ListOptions) (result *v1.ServiceMonitorList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ServiceMonitorList{}
	err = c.client.Get().Namespace(c.ns).Resource("servicemonitors").VersionedParams(&opts, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceMonitors) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("servicemonitors").VersionedParams(&opts, scheme.ParameterCodec).Watch()
}
func (c *serviceMonitors) Create(serviceMonitor *v1.ServiceMonitor) (result *v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ServiceMonitor{}
	err = c.client.Post().Namespace(c.ns).Resource("servicemonitors").Body(serviceMonitor).Do().Into(result)
	return
}
func (c *serviceMonitors) Update(serviceMonitor *v1.ServiceMonitor) (result *v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ServiceMonitor{}
	err = c.client.Put().Namespace(c.ns).Resource("servicemonitors").Name(serviceMonitor.Name).Body(serviceMonitor).Do().Into(result)
	return
}
func (c *serviceMonitors) Delete(name string, options *meta_v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("servicemonitors").Name(name).Body(options).Do().Error()
}
func (c *serviceMonitors) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("servicemonitors").VersionedParams(&listOptions, scheme.ParameterCodec).Body(options).Do().Error()
}
func (c *serviceMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ServiceMonitor{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("servicemonitors").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
