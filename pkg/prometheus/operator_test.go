package prometheus

import (
	"reflect"
	"testing"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/kylelemons/godebug/pretty"
)

func TestListOptions(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < 1000; i++ {
		o := ListOptions("test")
		if o.LabelSelector != "app=prometheus,prometheus=test" && o.LabelSelector != "prometheus=test,app=prometheus" {
			t.Fatalf("LabelSelector not computed correctly\n\nExpected: \"app=prometheus,prometheus=test\"\n\nGot:      %#+v", o.LabelSelector)
		}
	}
}
func TestCreateStatefulSetInputHash(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p1 := monitoringv1.Prometheus{}
	p1.Spec.Version = "v1.7.0"
	p2 := monitoringv1.Prometheus{}
	p2.Spec.Version = "v1.7.2"
	c := Config{}
	p1Hash, err := createSSetInputHash(p1, c, []string{})
	if err != nil {
		t.Fatal(err)
	}
	p2Hash, err := createSSetInputHash(p2, c, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if p1Hash == p2Hash {
		t.Fatal("expected two different Prometheus CRDs to result in two different hash but got equal hash")
	}
}
func TestGetNodeAddresses(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name			string
		nodes			*v1.NodeList
		expectedAddresses	[]string
		expectedErrors		int
	}{{name: "simple", nodes: &v1.NodeList{Items: []v1.Node{v1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node-0"}, Status: v1.NodeStatus{Addresses: []v1.NodeAddress{v1.NodeAddress{Address: "10.0.0.1", Type: v1.NodeInternalIP}}}}}}, expectedAddresses: []string{"10.0.0.1"}, expectedErrors: 0}, {name: "missing ip on one node", nodes: &v1.NodeList{Items: []v1.Node{v1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node-0"}, Status: v1.NodeStatus{Addresses: []v1.NodeAddress{v1.NodeAddress{Address: "node-0", Type: v1.NodeHostName}}}}, v1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node-1"}, Status: v1.NodeStatus{Addresses: []v1.NodeAddress{v1.NodeAddress{Address: "10.0.0.1", Type: v1.NodeInternalIP}}}}}}, expectedAddresses: []string{"10.0.0.1"}, expectedErrors: 1}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			addrs, errs := getNodeAddresses(c.nodes)
			if len(errs) != c.expectedErrors {
				t.Errorf("Expected %d errors, got %d. Errors: %v", c.expectedErrors, len(errs), errs)
			}
			ips := make([]string, 0)
			for _, addr := range addrs {
				ips = append(ips, addr.IP)
			}
			if !reflect.DeepEqual(ips, c.expectedAddresses) {
				t.Error(pretty.Compare(ips, c.expectedAddresses))
			}
		})
	}
}
