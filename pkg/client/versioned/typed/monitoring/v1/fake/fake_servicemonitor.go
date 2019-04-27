package fake

import (
	monitoring_v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeServiceMonitors struct {
	Fake	*FakeMonitoringV1
	ns	string
}

var servicemonitorsResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "servicemonitors"}
var servicemonitorsKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "ServiceMonitor"}

func (c *FakeServiceMonitors) Get(name string, options v1.GetOptions) (result *monitoring_v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(servicemonitorsResource, c.ns, name), &monitoring_v1.ServiceMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitor), err
}
func (c *FakeServiceMonitors) List(opts v1.ListOptions) (result *monitoring_v1.ServiceMonitorList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(servicemonitorsResource, servicemonitorsKind, c.ns, opts), &monitoring_v1.ServiceMonitorList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoring_v1.ServiceMonitorList{ListMeta: obj.(*monitoring_v1.ServiceMonitorList).ListMeta}
	for _, item := range obj.(*monitoring_v1.ServiceMonitorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeServiceMonitors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(servicemonitorsResource, c.ns, opts))
}
func (c *FakeServiceMonitors) Create(serviceMonitor *monitoring_v1.ServiceMonitor) (result *monitoring_v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(servicemonitorsResource, c.ns, serviceMonitor), &monitoring_v1.ServiceMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitor), err
}
func (c *FakeServiceMonitors) Update(serviceMonitor *monitoring_v1.ServiceMonitor) (result *monitoring_v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(servicemonitorsResource, c.ns, serviceMonitor), &monitoring_v1.ServiceMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitor), err
}
func (c *FakeServiceMonitors) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(servicemonitorsResource, c.ns, name), &monitoring_v1.ServiceMonitor{})
	return err
}
func (c *FakeServiceMonitors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(servicemonitorsResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &monitoring_v1.ServiceMonitorList{})
	return err
}
func (c *FakeServiceMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *monitoring_v1.ServiceMonitor, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(servicemonitorsResource, c.ns, name, data, subresources...), &monitoring_v1.ServiceMonitor{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitor), err
}
