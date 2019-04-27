package v1

import (
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

func (in *APIServerConfig) DeepCopyInto(out *APIServerConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	return
}
func (in *APIServerConfig) DeepCopy() *APIServerConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(APIServerConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *AlertingSpec) DeepCopyInto(out *AlertingSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Alertmanagers != nil {
		in, out := &in.Alertmanagers, &out.Alertmanagers
		*out = make([]AlertmanagerEndpoints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *AlertingSpec) DeepCopy() *AlertingSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AlertingSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *Alertmanager) DeepCopyInto(out *Alertmanager) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(AlertmanagerStatus)
			**out = **in
		}
	}
	return
}
func (in *Alertmanager) DeepCopy() *Alertmanager {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Alertmanager)
	in.DeepCopyInto(out)
	return out
}
func (in *AlertmanagerEndpoints) DeepCopyInto(out *AlertmanagerEndpoints) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Port = in.Port
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	return
}
func (in *AlertmanagerEndpoints) DeepCopy() *AlertmanagerEndpoints {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AlertmanagerEndpoints)
	in.DeepCopyInto(out)
	return out
}
func (in *AlertmanagerList) DeepCopyInto(out *AlertmanagerList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alertmanager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *AlertmanagerList) DeepCopy() *AlertmanagerList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AlertmanagerList)
	in.DeepCopyInto(out)
	return out
}
func (in *AlertmanagerSpec) DeepCopyInto(out *AlertmanagerSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.PodMetadata != nil {
		in, out := &in.PodMetadata, &out.PodMetadata
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.ObjectMeta)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]core_v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PodSecurityContext)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]core_v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalPeers != nil {
		in, out := &in.AdditionalPeers, &out.AdditionalPeers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *AlertmanagerSpec) DeepCopy() *AlertmanagerSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AlertmanagerSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *AlertmanagerStatus) DeepCopyInto(out *AlertmanagerStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *AlertmanagerStatus) DeepCopy() *AlertmanagerStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AlertmanagerStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *BasicAuth) DeepCopyInto(out *BasicAuth) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Password.DeepCopyInto(&out.Password)
	return
}
func (in *BasicAuth) DeepCopy() *BasicAuth {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BasicAuth)
	in.DeepCopyInto(out)
	return out
}
func (in *CrdKind) DeepCopyInto(out *CrdKind) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *CrdKind) DeepCopy() *CrdKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CrdKind)
	in.DeepCopyInto(out)
	return out
}
func (in *CrdKinds) DeepCopyInto(out *CrdKinds) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Prometheus = in.Prometheus
	out.Alertmanager = in.Alertmanager
	out.ServiceMonitor = in.ServiceMonitor
	out.PrometheusRule = in.PrometheusRule
	return
}
func (in *CrdKinds) DeepCopy() *CrdKinds {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CrdKinds)
	in.DeepCopyInto(out)
	return out
}
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		if *in == nil {
			*out = nil
		} else {
			*out = new(intstr.IntOrString)
			**out = **in
		}
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.MetricRelabelConfigs != nil {
		in, out := &in.MetricRelabelConfigs, &out.MetricRelabelConfigs
		*out = make([]*RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(RelabelConfig)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.RelabelConfigs != nil {
		in, out := &in.RelabelConfigs, &out.RelabelConfigs
		*out = make([]*RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(RelabelConfig)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.ProxyURL != nil {
		in, out := &in.ProxyURL, &out.ProxyURL
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}
func (in *Endpoint) DeepCopy() *Endpoint {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}
func (in *NamespaceSelector) DeepCopyInto(out *NamespaceSelector) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.MatchNames != nil {
		in, out := &in.MatchNames, &out.MatchNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *NamespaceSelector) DeepCopy() *NamespaceSelector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(NamespaceSelector)
	in.DeepCopyInto(out)
	return out
}
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(PrometheusStatus)
			**out = **in
		}
	}
	return
}
func (in *Prometheus) DeepCopy() *Prometheus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}
func (in *PrometheusList) DeepCopyInto(out *PrometheusList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*Prometheus, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Prometheus)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}
func (in *PrometheusList) DeepCopy() *PrometheusList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PrometheusList)
	in.DeepCopyInto(out)
	return out
}
func (in *PrometheusRule) DeepCopyInto(out *PrometheusRule) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}
func (in *PrometheusRule) DeepCopy() *PrometheusRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PrometheusRule)
	in.DeepCopyInto(out)
	return out
}
func (in *PrometheusRuleList) DeepCopyInto(out *PrometheusRuleList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*PrometheusRule, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(PrometheusRule)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}
func (in *PrometheusRuleList) DeepCopy() *PrometheusRuleList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleList)
	in.DeepCopyInto(out)
	return out
}
func (in *PrometheusRuleSpec) DeepCopyInto(out *PrometheusRuleSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]RuleGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *PrometheusRuleSpec) DeepCopy() *PrometheusRuleSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *PrometheusSpec) DeepCopyInto(out *PrometheusSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.PodMetadata != nil {
		in, out := &in.PodMetadata, &out.PodMetadata
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.ObjectMeta)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ServiceMonitorSelector != nil {
		in, out := &in.ServiceMonitorSelector, &out.ServiceMonitorSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ServiceMonitorNamespaceSelector != nil {
		in, out := &in.ServiceMonitorNamespaceSelector, &out.ServiceMonitorNamespaceSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]core_v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.ExternalLabels != nil {
		in, out := &in.ExternalLabels, &out.ExternalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		if *in == nil {
			*out = nil
		} else {
			*out = new(QuerySpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		if *in == nil {
			*out = nil
		} else {
			*out = new(StorageSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RuleSelector != nil {
		in, out := &in.RuleSelector, &out.RuleSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RuleNamespaceSelector != nil {
		in, out := &in.RuleNamespaceSelector, &out.RuleNamespaceSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Alerting != nil {
		in, out := &in.Alerting, &out.Alerting
		if *in == nil {
			*out = nil
		} else {
			*out = new(AlertingSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]core_v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemoteWrite != nil {
		in, out := &in.RemoteWrite, &out.RemoteWrite
		*out = make([]RemoteWriteSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemoteRead != nil {
		in, out := &in.RemoteRead, &out.RemoteRead
		*out = make([]RemoteReadSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PodSecurityContext)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]core_v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalScrapeConfigs != nil {
		in, out := &in.AdditionalScrapeConfigs, &out.AdditionalScrapeConfigs
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AdditionalAlertRelabelConfigs != nil {
		in, out := &in.AdditionalAlertRelabelConfigs, &out.AdditionalAlertRelabelConfigs
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AdditionalAlertManagerConfigs != nil {
		in, out := &in.AdditionalAlertManagerConfigs, &out.AdditionalAlertManagerConfigs
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.APIServerConfig != nil {
		in, out := &in.APIServerConfig, &out.APIServerConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(APIServerConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Thanos != nil {
		in, out := &in.Thanos, &out.Thanos
		if *in == nil {
			*out = nil
		} else {
			*out = new(ThanosSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}
func (in *PrometheusSpec) DeepCopy() *PrometheusSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PrometheusSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *PrometheusStatus) DeepCopyInto(out *PrometheusStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *PrometheusStatus) DeepCopy() *PrometheusStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PrometheusStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *QuerySpec) DeepCopyInto(out *QuerySpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.LookbackDelta != nil {
		in, out := &in.LookbackDelta, &out.LookbackDelta
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.MaxConcurrency != nil {
		in, out := &in.MaxConcurrency, &out.MaxConcurrency
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}
func (in *QuerySpec) DeepCopy() *QuerySpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(QuerySpec)
	in.DeepCopyInto(out)
	return out
}
func (in *QueueConfig) DeepCopyInto(out *QueueConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *QueueConfig) DeepCopy() *QueueConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(QueueConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *RelabelConfig) DeepCopyInto(out *RelabelConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SourceLabels != nil {
		in, out := &in.SourceLabels, &out.SourceLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *RelabelConfig) DeepCopy() *RelabelConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RelabelConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *RemoteReadSpec) DeepCopyInto(out *RemoteReadSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.RequiredMatchers != nil {
		in, out := &in.RequiredMatchers, &out.RequiredMatchers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	return
}
func (in *RemoteReadSpec) DeepCopy() *RemoteReadSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RemoteReadSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *RemoteWriteSpec) DeepCopyInto(out *RemoteWriteSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.WriteRelabelConfigs != nil {
		in, out := &in.WriteRelabelConfigs, &out.WriteRelabelConfigs
		*out = make([]RelabelConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BasicAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(TLSConfig)
			**out = **in
		}
	}
	if in.QueueConfig != nil {
		in, out := &in.QueueConfig, &out.QueueConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(QueueConfig)
			**out = **in
		}
	}
	return
}
func (in *RemoteWriteSpec) DeepCopy() *RemoteWriteSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RemoteWriteSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *Rule) DeepCopyInto(out *Rule) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Expr = in.Expr
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}
func (in *Rule) DeepCopy() *Rule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}
func (in *RuleGroup) DeepCopyInto(out *RuleGroup) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *RuleGroup) DeepCopy() *RuleGroup {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RuleGroup)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceMonitor) DeepCopyInto(out *ServiceMonitor) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}
func (in *ServiceMonitor) DeepCopy() *ServiceMonitor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceMonitor)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceMonitorList) DeepCopyInto(out *ServiceMonitorList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*ServiceMonitor, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(ServiceMonitor)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}
func (in *ServiceMonitorList) DeepCopy() *ServiceMonitorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorList)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceMonitorSpec) DeepCopyInto(out *ServiceMonitorSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.TargetLabels != nil {
		in, out := &in.TargetLabels, &out.TargetLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodTargetLabels != nil {
		in, out := &in.PodTargetLabels, &out.PodTargetLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	return
}
func (in *ServiceMonitorSpec) DeepCopy() *ServiceMonitorSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.EmptyDirVolumeSource)
			(*in).DeepCopyInto(*out)
		}
	}
	in.VolumeClaimTemplate.DeepCopyInto(&out.VolumeClaimTemplate)
	return
}
func (in *StorageSpec) DeepCopy() *StorageSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *TLSConfig) DeepCopy() *TLSConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *ThanosGCSSpec) DeepCopyInto(out *ThanosGCSSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}
func (in *ThanosGCSSpec) DeepCopy() *ThanosGCSSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ThanosGCSSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ThanosS3Spec) DeepCopyInto(out *ThanosS3Spec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.SecretKeySelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.SignatureVersion2 != nil {
		in, out := &in.SignatureVersion2, &out.SignatureVersion2
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.EncryptSSE != nil {
		in, out := &in.EncryptSSE, &out.EncryptSSE
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}
func (in *ThanosS3Spec) DeepCopy() *ThanosS3Spec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ThanosS3Spec)
	in.DeepCopyInto(out)
	return out
}
func (in *ThanosSpec) DeepCopyInto(out *ThanosSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.SHA != nil {
		in, out := &in.SHA, &out.SHA
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.BaseImage != nil {
		in, out := &in.BaseImage, &out.BaseImage
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		if *in == nil {
			*out = nil
		} else {
			*out = new(ThanosGCSSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		if *in == nil {
			*out = nil
		} else {
			*out = new(ThanosS3Spec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}
func (in *ThanosSpec) DeepCopy() *ThanosSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ThanosSpec)
	in.DeepCopyInto(out)
	return out
}
