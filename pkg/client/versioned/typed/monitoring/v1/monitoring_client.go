package v1

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/coreos/prometheus-operator/pkg/client/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type MonitoringV1Interface interface {
	RESTClient() rest.Interface
	AlertmanagersGetter
	PrometheusesGetter
	PrometheusRulesGetter
	ServiceMonitorsGetter
}
type MonitoringV1Client struct{ restClient rest.Interface }

func (c *MonitoringV1Client) Alertmanagers(namespace string) AlertmanagerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newAlertmanagers(c, namespace)
}
func (c *MonitoringV1Client) Prometheuses(namespace string) PrometheusInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newPrometheuses(c, namespace)
}
func (c *MonitoringV1Client) PrometheusRules(namespace string) PrometheusRuleInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newPrometheusRules(c, namespace)
}
func (c *MonitoringV1Client) ServiceMonitors(namespace string) ServiceMonitorInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceMonitors(c, namespace)
}
func NewForConfig(c *rest.Config) (*MonitoringV1Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MonitoringV1Client{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *MonitoringV1Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
func New(c rest.Interface) *MonitoringV1Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MonitoringV1Client{c}
}
func setConfigDefaults(config *rest.Config) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	return nil
}
func (c *MonitoringV1Client) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil
	}
	return c.restClient
}
