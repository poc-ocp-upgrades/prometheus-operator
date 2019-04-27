package v1

import (
	"encoding/json"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestMarshallServiceMonitor(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sm := &ServiceMonitor{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default", Labels: map[string]string{"group": "group1"}}, Spec: ServiceMonitorSpec{NamespaceSelector: NamespaceSelector{MatchNames: []string{"test"}}, Endpoints: []Endpoint{{Port: "metric"}}}}
	expected := `{"metadata":{"name":"test","namespace":"default","creationTimestamp":null,"labels":{"group":"group1"}},"spec":{"endpoints":[{"port":"metric"}],"selector":{},"namespaceSelector":{"matchNames":["test"]}}}`
	r, err := json.Marshal(sm)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	rs := string(r)
	if rs != expected {
		t.Fatalf("Got %s expected: %s ", rs, expected)
	}
}
