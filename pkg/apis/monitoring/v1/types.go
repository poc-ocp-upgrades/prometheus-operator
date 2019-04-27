package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	Version			= "v1"
	PrometheusesKind	= "Prometheus"
	PrometheusName		= "prometheuses"
	PrometheusKindKey	= "prometheus"
	AlertmanagersKind	= "Alertmanager"
	AlertmanagerName	= "alertmanagers"
	AlertManagerKindKey	= "alertmanager"
	ServiceMonitorsKind	= "ServiceMonitor"
	ServiceMonitorName	= "servicemonitors"
	ServiceMonitorKindKey	= "servicemonitor"
	PrometheusRuleKind	= "PrometheusRule"
	PrometheusRuleName	= "prometheusrules"
	PrometheusRuleKindKey	= "prometheusrule"
)

type Prometheus struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			PrometheusSpec		`json:"spec"`
	Status			*PrometheusStatus	`json:"status,omitempty"`
}
type PrometheusList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]*Prometheus	`json:"items"`
}
type PrometheusSpec struct {
	PodMetadata			*metav1.ObjectMeta		`json:"podMetadata,omitempty"`
	ServiceMonitorSelector		*metav1.LabelSelector		`json:"serviceMonitorSelector,omitempty"`
	ServiceMonitorNamespaceSelector	*metav1.LabelSelector		`json:"serviceMonitorNamespaceSelector,omitempty"`
	Version				string				`json:"version,omitempty"`
	Tag				string				`json:"tag,omitempty"`
	SHA				string				`json:"sha,omitempty"`
	Paused				bool				`json:"paused,omitempty"`
	Image				*string				`json:"image,omitempty"`
	BaseImage			string				`json:"baseImage,omitempty"`
	ImagePullSecrets		[]v1.LocalObjectReference	`json:"imagePullSecrets,omitempty"`
	Replicas			*int32				`json:"replicas,omitempty"`
	Retention			string				`json:"retention,omitempty"`
	LogLevel			string				`json:"logLevel,omitempty"`
	ScrapeInterval			string				`json:"scrapeInterval,omitempty"`
	EvaluationInterval		string				`json:"evaluationInterval,omitempty"`
	ExternalLabels			map[string]string		`json:"externalLabels,omitempty"`
	ExternalURL			string				`json:"externalUrl,omitempty"`
	RoutePrefix			string				`json:"routePrefix,omitempty"`
	Query				*QuerySpec			`json:"query,omitempty"`
	Storage				*StorageSpec			`json:"storage,omitempty"`
	RuleSelector			*metav1.LabelSelector		`json:"ruleSelector,omitempty"`
	RuleNamespaceSelector		*metav1.LabelSelector		`json:"ruleNamespaceSelector,omitempty"`
	Alerting			*AlertingSpec			`json:"alerting,omitempty"`
	Resources			v1.ResourceRequirements		`json:"resources,omitempty"`
	NodeSelector			map[string]string		`json:"nodeSelector,omitempty"`
	ServiceAccountName		string				`json:"serviceAccountName,omitempty"`
	Secrets				[]string			`json:"secrets,omitempty"`
	ConfigMaps			[]string			`json:"configMaps,omitempty"`
	Affinity			*v1.Affinity			`json:"affinity,omitempty"`
	Tolerations			[]v1.Toleration			`json:"tolerations,omitempty"`
	RemoteWrite			[]RemoteWriteSpec		`json:"remoteWrite,omitempty"`
	RemoteRead			[]RemoteReadSpec		`json:"remoteRead,omitempty"`
	SecurityContext			*v1.PodSecurityContext		`json:"securityContext,omitempty"`
	ListenLocal			bool				`json:"listenLocal,omitempty"`
	Containers			[]v1.Container			`json:"containers,omitempty"`
	AdditionalScrapeConfigs		*v1.SecretKeySelector		`json:"additionalScrapeConfigs,omitempty"`
	AdditionalAlertRelabelConfigs	*v1.SecretKeySelector		`json:"additionalAlertRelabelConfigs,omitempty"`
	AdditionalAlertManagerConfigs	*v1.SecretKeySelector		`json:"additionalAlertManagerConfigs,omitempty"`
	APIServerConfig			*APIServerConfig		`json:"apiserverConfig,omitempty"`
	Thanos				*ThanosSpec			`json:"thanos,omitempty"`
	PriorityClassName		string				`json:"priorityClassName,omitempty"`
}
type PrometheusStatus struct {
	Paused			bool	`json:"paused"`
	Replicas		int32	`json:"replicas"`
	UpdatedReplicas		int32	`json:"updatedReplicas"`
	AvailableReplicas	int32	`json:"availableReplicas"`
	UnavailableReplicas	int32	`json:"unavailableReplicas"`
}
type AlertingSpec struct {
	Alertmanagers []AlertmanagerEndpoints `json:"alertmanagers"`
}
type StorageSpec struct {
	EmptyDir		*v1.EmptyDirVolumeSource	`json:"emptyDir,omitempty"`
	VolumeClaimTemplate	v1.PersistentVolumeClaim	`json:"volumeClaimTemplate,omitempty"`
}
type QuerySpec struct {
	LookbackDelta	*string	`json:"lookbackDelta,omitempty"`
	MaxConcurrency	*int32	`json:"maxConcurrency,omitempty"`
	Timeout		*string	`json:"timeout,omitempty"`
}
type ThanosSpec struct {
	Peers		*string			`json:"peers,omitempty"`
	Image		*string			`json:"image,omitempty"`
	Version		*string			`json:"version,omitempty"`
	Tag		*string			`json:"tag,omitempty"`
	SHA		*string			`json:"sha,omitempty"`
	BaseImage	*string			`json:"baseImage,omitempty"`
	Resources	v1.ResourceRequirements	`json:"resources,omitempty"`
	GCS		*ThanosGCSSpec		`json:"gcs,omitempty"`
	S3		*ThanosS3Spec		`json:"s3,omitempty"`
}
type ThanosGCSSpec struct {
	Bucket		*string			`json:"bucket,omitempty"`
	SecretKey	*v1.SecretKeySelector	`json:"credentials,omitempty"`
}
type ThanosS3Spec struct {
	Bucket			*string			`json:"bucket,omitempty"`
	Endpoint		*string			`json:"endpoint,omitempty"`
	AccessKey		*v1.SecretKeySelector	`json:"accessKey,omitempty"`
	SecretKey		*v1.SecretKeySelector	`json:"secretKey,omitempty"`
	Insecure		*bool			`json:"insecure,omitempty"`
	SignatureVersion2	*bool			`json:"signatureVersion2,omitempty"`
	EncryptSSE		*bool			`json:"encryptsse,omitempty"`
}
type RemoteWriteSpec struct {
	URL			string		`json:"url"`
	RemoteTimeout		string		`json:"remoteTimeout,omitempty"`
	WriteRelabelConfigs	[]RelabelConfig	`json:"writeRelabelConfigs,omitempty"`
	BasicAuth		*BasicAuth	`json:"basicAuth,omitempty"`
	BearerToken		string		`json:"bearerToken,omitempty"`
	BearerTokenFile		string		`json:"bearerTokenFile,omitempty"`
	TLSConfig		*TLSConfig	`json:"tlsConfig,omitempty"`
	ProxyURL		string		`json:"proxyUrl,omitempty"`
	QueueConfig		*QueueConfig	`json:"queueConfig,omitempty"`
}
type QueueConfig struct {
	Capacity		int	`json:"capacity,omitempty"`
	MaxShards		int	`json:"maxShards,omitempty"`
	MaxSamplesPerSend	int	`json:"maxSamplesPerSend,omitempty"`
	BatchSendDeadline	string	`json:"batchSendDeadline,omitempty"`
	MaxRetries		int	`json:"maxRetries,omitempty"`
	MinBackoff		string	`json:"minBackoff,omitempty"`
	MaxBackoff		string	`json:"maxBackoff,omitempty"`
}
type RemoteReadSpec struct {
	URL			string			`json:"url"`
	RequiredMatchers	map[string]string	`json:"requiredMatchers,omitempty"`
	RemoteTimeout		string			`json:"remoteTimeout,omitempty"`
	ReadRecent		bool			`json:"readRecent,omitempty"`
	BasicAuth		*BasicAuth		`json:"basicAuth,omitempty"`
	BearerToken		string			`json:"bearerToken,omitempty"`
	BearerTokenFile		string			`json:"bearerTokenFile,omitempty"`
	TLSConfig		*TLSConfig		`json:"tlsConfig,omitempty"`
	ProxyURL		string			`json:"proxyUrl,omitempty"`
}
type RelabelConfig struct {
	SourceLabels	[]string	`json:"sourceLabels,omitempty"`
	Separator	string		`json:"separator,omitempty"`
	TargetLabel	string		`json:"targetLabel,omitempty"`
	Regex		string		`json:"regex,omitempty"`
	Modulus		uint64		`json:"modulus,omitempty"`
	Replacement	string		`json:"replacement,omitempty"`
	Action		string		`json:"action,omitempty"`
}
type APIServerConfig struct {
	Host		string		`json:"host"`
	BasicAuth	*BasicAuth	`json:"basicAuth,omitempty"`
	BearerToken	string		`json:"bearerToken,omitempty"`
	BearerTokenFile	string		`json:"bearerTokenFile,omitempty"`
	TLSConfig	*TLSConfig	`json:"tlsConfig,omitempty"`
}
type AlertmanagerEndpoints struct {
	Namespace	string			`json:"namespace"`
	Name		string			`json:"name"`
	Port		intstr.IntOrString	`json:"port"`
	Scheme		string			`json:"scheme,omitempty"`
	PathPrefix	string			`json:"pathPrefix,omitempty"`
	TLSConfig	*TLSConfig		`json:"tlsConfig,omitempty"`
	BearerTokenFile	string			`json:"bearerTokenFile,omitempty"`
}
type ServiceMonitor struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ServiceMonitorSpec	`json:"spec"`
}
type ServiceMonitorSpec struct {
	JobLabel		string			`json:"jobLabel,omitempty"`
	TargetLabels		[]string		`json:"targetLabels,omitempty"`
	PodTargetLabels		[]string		`json:"podTargetLabels,omitempty"`
	Endpoints		[]Endpoint		`json:"endpoints"`
	Selector		metav1.LabelSelector	`json:"selector"`
	NamespaceSelector	NamespaceSelector	`json:"namespaceSelector,omitempty"`
	SampleLimit		uint64			`json:"sampleLimit,omitempty"`
}
type Endpoint struct {
	Port			string			`json:"port,omitempty"`
	TargetPort		*intstr.IntOrString	`json:"targetPort,omitempty"`
	Path			string			`json:"path,omitempty"`
	Scheme			string			`json:"scheme,omitempty"`
	Params			map[string][]string	`json:"params,omitempty"`
	Interval		string			`json:"interval,omitempty"`
	ScrapeTimeout		string			`json:"scrapeTimeout,omitempty"`
	TLSConfig		*TLSConfig		`json:"tlsConfig,omitempty"`
	BearerTokenFile		string			`json:"bearerTokenFile,omitempty"`
	HonorLabels		bool			`json:"honorLabels,omitempty"`
	BasicAuth		*BasicAuth		`json:"basicAuth,omitempty"`
	MetricRelabelConfigs	[]*RelabelConfig	`json:"metricRelabelings,omitempty"`
	RelabelConfigs		[]*RelabelConfig	`json:"relabelings,omitempty"`
	ProxyURL		*string			`json:"proxyUrl,omitempty"`
}
type BasicAuth struct {
	Username	v1.SecretKeySelector	`json:"username,omitempty"`
	Password	v1.SecretKeySelector	`json:"password,omitempty"`
}
type TLSConfig struct {
	CAFile			string	`json:"caFile,omitempty"`
	CertFile		string	`json:"certFile,omitempty"`
	KeyFile			string	`json:"keyFile,omitempty"`
	ServerName		string	`json:"serverName,omitempty"`
	InsecureSkipVerify	bool	`json:"insecureSkipVerify,omitempty"`
}
type ServiceMonitorList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]*ServiceMonitor	`json:"items"`
}
type PrometheusRuleList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]*PrometheusRule	`json:"items"`
}
type PrometheusRule struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			PrometheusRuleSpec	`json:"spec"`
}
type PrometheusRuleSpec struct {
	Groups []RuleGroup `json:"groups,omitempty"`
}
type RuleGroup struct {
	Name		string	`json:"name"`
	Interval	string	`json:"interval,omitempty"`
	Rules		[]Rule	`json:"rules"`
}
type Rule struct {
	Record		string			`json:"record,omitempty"`
	Alert		string			`json:"alert,omitempty"`
	Expr		intstr.IntOrString	`json:"expr"`
	For		string			`json:"for,omitempty"`
	Labels		map[string]string	`json:"labels,omitempty"`
	Annotations	map[string]string	`json:"annotations,omitempty"`
}
type Alertmanager struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			AlertmanagerSpec	`json:"spec"`
	Status			*AlertmanagerStatus	`json:"status,omitempty"`
}
type AlertmanagerSpec struct {
	PodMetadata		*metav1.ObjectMeta		`json:"podMetadata,omitempty"`
	Image			*string				`json:"image,omitempty"`
	Version			string				`json:"version,omitempty"`
	Tag			string				`json:"tag,omitempty"`
	SHA			string				`json:"sha,omitempty"`
	BaseImage		string				`json:"baseImage,omitempty"`
	ImagePullSecrets	[]v1.LocalObjectReference	`json:"imagePullSecrets,omitempty"`
	Secrets			[]string			`json:"secrets,omitempty"`
	ConfigMaps		[]string			`json:"configMaps,omitempty"`
	LogLevel		string				`json:"logLevel,omitempty"`
	Replicas		*int32				`json:"replicas,omitempty"`
	Retention		string				`json:"retention,omitempty"`
	Storage			*StorageSpec			`json:"storage,omitempty"`
	ExternalURL		string				`json:"externalUrl,omitempty"`
	RoutePrefix		string				`json:"routePrefix,omitempty"`
	Paused			bool				`json:"paused,omitempty"`
	NodeSelector		map[string]string		`json:"nodeSelector,omitempty"`
	Resources		v1.ResourceRequirements		`json:"resources,omitempty"`
	Affinity		*v1.Affinity			`json:"affinity,omitempty"`
	Tolerations		[]v1.Toleration			`json:"tolerations,omitempty"`
	SecurityContext		*v1.PodSecurityContext		`json:"securityContext,omitempty"`
	ServiceAccountName	string				`json:"serviceAccountName,omitempty"`
	ListenLocal		bool				`json:"listenLocal,omitempty"`
	Containers		[]v1.Container			`json:"containers,omitempty"`
	PriorityClassName	string				`json:"priorityClassName,omitempty"`
	AdditionalPeers		[]string			`json:"additionalPeers,omitempty"`
}
type AlertmanagerList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]Alertmanager	`json:"items"`
}
type AlertmanagerStatus struct {
	Paused			bool	`json:"paused"`
	Replicas		int32	`json:"replicas"`
	UpdatedReplicas		int32	`json:"updatedReplicas"`
	AvailableReplicas	int32	`json:"availableReplicas"`
	UnavailableReplicas	int32	`json:"unavailableReplicas"`
}
type NamespaceSelector struct {
	Any		bool		`json:"any,omitempty"`
	MatchNames	[]string	`json:"matchNames,omitempty"`
}

func (l *Alertmanager) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.DeepCopy()
}
func (l *AlertmanagerList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.DeepCopy()
}
func (l *Prometheus) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.DeepCopy()
}
func (l *PrometheusList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.DeepCopy()
}
func (l *ServiceMonitor) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.DeepCopy()
}
func (l *ServiceMonitorList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.DeepCopy()
}
func (f *PrometheusRule) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.DeepCopy()
}
func (l *PrometheusRuleList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.DeepCopy()
}
