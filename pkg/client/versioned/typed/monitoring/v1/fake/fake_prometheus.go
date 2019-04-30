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

type FakePrometheuses struct {
	Fake	*FakeMonitoringV1
	ns	string
}

var prometheusesResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "prometheuses"}
var prometheusesKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "Prometheus"}

func (c *FakePrometheuses) Get(name string, options v1.GetOptions) (result *monitoring_v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(prometheusesResource, c.ns, name), &monitoring_v1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Prometheus), err
}
func (c *FakePrometheuses) List(opts v1.ListOptions) (result *monitoring_v1.PrometheusList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(prometheusesResource, prometheusesKind, c.ns, opts), &monitoring_v1.PrometheusList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoring_v1.PrometheusList{ListMeta: obj.(*monitoring_v1.PrometheusList).ListMeta}
	for _, item := range obj.(*monitoring_v1.PrometheusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakePrometheuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(prometheusesResource, c.ns, opts))
}
func (c *FakePrometheuses) Create(prometheus *monitoring_v1.Prometheus) (result *monitoring_v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(prometheusesResource, c.ns, prometheus), &monitoring_v1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Prometheus), err
}
func (c *FakePrometheuses) Update(prometheus *monitoring_v1.Prometheus) (result *monitoring_v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(prometheusesResource, c.ns, prometheus), &monitoring_v1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Prometheus), err
}
func (c *FakePrometheuses) UpdateStatus(prometheus *monitoring_v1.Prometheus) (*monitoring_v1.Prometheus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(prometheusesResource, "status", c.ns, prometheus), &monitoring_v1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Prometheus), err
}
func (c *FakePrometheuses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(prometheusesResource, c.ns, name), &monitoring_v1.Prometheus{})
	return err
}
func (c *FakePrometheuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(prometheusesResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &monitoring_v1.PrometheusList{})
	return err
}
func (c *FakePrometheuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *monitoring_v1.Prometheus, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(prometheusesResource, c.ns, name, data, subresources...), &monitoring_v1.Prometheus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Prometheus), err
}
