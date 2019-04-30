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

type FakeAlertmanagers struct {
	Fake	*FakeMonitoringV1
	ns	string
}

var alertmanagersResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "alertmanagers"}
var alertmanagersKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "Alertmanager"}

func (c *FakeAlertmanagers) Get(name string, options v1.GetOptions) (result *monitoring_v1.Alertmanager, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(alertmanagersResource, c.ns, name), &monitoring_v1.Alertmanager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Alertmanager), err
}
func (c *FakeAlertmanagers) List(opts v1.ListOptions) (result *monitoring_v1.AlertmanagerList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(alertmanagersResource, alertmanagersKind, c.ns, opts), &monitoring_v1.AlertmanagerList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoring_v1.AlertmanagerList{ListMeta: obj.(*monitoring_v1.AlertmanagerList).ListMeta}
	for _, item := range obj.(*monitoring_v1.AlertmanagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeAlertmanagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(alertmanagersResource, c.ns, opts))
}
func (c *FakeAlertmanagers) Create(alertmanager *monitoring_v1.Alertmanager) (result *monitoring_v1.Alertmanager, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(alertmanagersResource, c.ns, alertmanager), &monitoring_v1.Alertmanager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Alertmanager), err
}
func (c *FakeAlertmanagers) Update(alertmanager *monitoring_v1.Alertmanager) (result *monitoring_v1.Alertmanager, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(alertmanagersResource, c.ns, alertmanager), &monitoring_v1.Alertmanager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Alertmanager), err
}
func (c *FakeAlertmanagers) UpdateStatus(alertmanager *monitoring_v1.Alertmanager) (*monitoring_v1.Alertmanager, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(alertmanagersResource, "status", c.ns, alertmanager), &monitoring_v1.Alertmanager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Alertmanager), err
}
func (c *FakeAlertmanagers) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(alertmanagersResource, c.ns, name), &monitoring_v1.Alertmanager{})
	return err
}
func (c *FakeAlertmanagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(alertmanagersResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &monitoring_v1.AlertmanagerList{})
	return err
}
func (c *FakeAlertmanagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *monitoring_v1.Alertmanager, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(alertmanagersResource, c.ns, name, data, subresources...), &monitoring_v1.Alertmanager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.Alertmanager), err
}
