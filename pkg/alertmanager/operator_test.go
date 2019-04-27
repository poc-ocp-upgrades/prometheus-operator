package alertmanager

import (
	"testing"
)

func TestListOptions(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < 1000; i++ {
		o := ListOptions("test")
		if o.LabelSelector != "app=alertmanager,alertmanager=test" && o.LabelSelector != "alertmanager=test,app=alertmanager" {
			t.Fatalf("LabelSelector not computed correctly\n\nExpected: \"app=alertmanager,alertmanager=test\"\n\nGot:      %#+v", o.LabelSelector)
		}
	}
}
