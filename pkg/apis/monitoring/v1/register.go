package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"github.com/coreos/prometheus-operator/pkg/apis/monitoring"
)

var SchemeGroupVersion = schema.GroupVersion{Group: monitoring.GroupName, Version: Version}

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder		runtime.SchemeBuilder
	localSchemeBuilder	= &SchemeBuilder
	AddToScheme		= localSchemeBuilder.AddToScheme
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(addKnownTypes)
}
func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(SchemeGroupVersion, &Prometheus{}, &PrometheusList{}, &ServiceMonitor{}, &ServiceMonitorList{}, &Alertmanager{}, &AlertmanagerList{}, &PrometheusRule{}, &PrometheusRuleList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
