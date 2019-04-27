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

type FakePrometheusRules struct {
	Fake	*FakeMonitoringV1
	ns	string
}

var prometheusrulesResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "prometheusrules"}
var prometheusrulesKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "PrometheusRule"}

func (c *FakePrometheusRules) Get(name string, options v1.GetOptions) (result *monitoring_v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(prometheusrulesResource, c.ns, name), &monitoring_v1.PrometheusRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.PrometheusRule), err
}
func (c *FakePrometheusRules) List(opts v1.ListOptions) (result *monitoring_v1.PrometheusRuleList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(prometheusrulesResource, prometheusrulesKind, c.ns, opts), &monitoring_v1.PrometheusRuleList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoring_v1.PrometheusRuleList{ListMeta: obj.(*monitoring_v1.PrometheusRuleList).ListMeta}
	for _, item := range obj.(*monitoring_v1.PrometheusRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakePrometheusRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(prometheusrulesResource, c.ns, opts))
}
func (c *FakePrometheusRules) Create(prometheusRule *monitoring_v1.PrometheusRule) (result *monitoring_v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(prometheusrulesResource, c.ns, prometheusRule), &monitoring_v1.PrometheusRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.PrometheusRule), err
}
func (c *FakePrometheusRules) Update(prometheusRule *monitoring_v1.PrometheusRule) (result *monitoring_v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(prometheusrulesResource, c.ns, prometheusRule), &monitoring_v1.PrometheusRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.PrometheusRule), err
}
func (c *FakePrometheusRules) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(prometheusrulesResource, c.ns, name), &monitoring_v1.PrometheusRule{})
	return err
}
func (c *FakePrometheusRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(prometheusrulesResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &monitoring_v1.PrometheusRuleList{})
	return err
}
func (c *FakePrometheusRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *monitoring_v1.PrometheusRule, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(prometheusrulesResource, c.ns, name, data, subresources...), &monitoring_v1.PrometheusRule{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.PrometheusRule), err
}
