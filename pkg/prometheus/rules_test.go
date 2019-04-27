package prometheus

import (
	"strings"
	"testing"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/api/core/v1"
)

func TestMakeRulesConfigMaps(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Run("ShouldReturnAtLeastOneConfigMap", shouldReturnAtLeastOneConfigMap)
	t.Run("ShouldErrorOnTooLargeRuleFile", shouldErrorOnTooLargeRuleFile)
	t.Run("ShouldSplitUpLargeSmallIntoTwo", shouldSplitUpLargeSmallIntoTwo)
}
func shouldReturnAtLeastOneConfigMap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &monitoringv1.Prometheus{}
	ruleFiles := map[string]string{}
	configMaps, err := makeRulesConfigMaps(p, ruleFiles)
	if err != nil {
		t.Fatalf("expected no error but got: %v", err.Error())
	}
	if len(configMaps) != 1 {
		t.Fatalf("expected one ConfigMaps but got %v", len(configMaps))
	}
}
func shouldErrorOnTooLargeRuleFile(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	expectedError := "rule file 'my-rule-file' is too large for a single Kubernetes ConfigMap"
	p := &monitoringv1.Prometheus{}
	ruleFiles := map[string]string{}
	ruleFiles["my-rule-file"] = strings.Repeat("a", v1.MaxSecretSize+1)
	_, err := makeRulesConfigMaps(p, ruleFiles)
	if err == nil || err.Error() != expectedError {
		t.Fatalf("expected makeRulesConfigMaps to return error '%v' but got '%v'", expectedError, err)
	}
}
func shouldSplitUpLargeSmallIntoTwo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &monitoringv1.Prometheus{}
	ruleFiles := map[string]string{}
	ruleFiles["first"] = strings.Repeat("a", maxConfigMapDataSize)
	ruleFiles["second"] = "a"
	configMaps, err := makeRulesConfigMaps(p, ruleFiles)
	if err != nil {
		t.Fatalf("expected no error but got: %v", err)
	}
	if len(configMaps) != 2 {
		t.Fatalf("expected rule files to be split up into two ConfigMaps, but got '%v' instead", len(configMaps))
	}
	if configMaps[0].Data["first"] != ruleFiles["first"] || configMaps[1].Data["second"] != ruleFiles["second"] {
		t.Fatal("expected ConfigMap data to match rule file content")
	}
}
