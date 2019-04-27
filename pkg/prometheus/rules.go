package prometheus

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"github.com/ghodss/yaml"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const labelPrometheusName = "prometheus-name"

var maxConfigMapDataSize = int(float64(v1.MaxSecretSize) * 0.5)

func (c *Operator) createOrUpdateRuleConfigMaps(p *monitoringv1.Prometheus) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cClient := c.kclient.CoreV1().ConfigMaps(p.Namespace)
	namespaces, err := c.selectRuleNamespaces(p)
	if err != nil {
		return nil, err
	}
	newRules, err := c.selectRules(p, namespaces)
	if err != nil {
		return nil, err
	}
	currentConfigMapList, err := cClient.List(prometheusRulesConfigMapSelector(p.Name))
	if err != nil {
		return nil, err
	}
	currentConfigMaps := currentConfigMapList.Items
	currentRules := map[string]string{}
	for _, cm := range currentConfigMaps {
		for ruleFileName, ruleFile := range cm.Data {
			currentRules[ruleFileName] = ruleFile
		}
	}
	equal := reflect.DeepEqual(newRules, currentRules)
	if equal && len(currentConfigMaps) != 0 {
		level.Debug(c.logger).Log("msg", "no PrometheusRule changes", "namespace", p.Namespace, "prometheus", p.Name)
		currentConfigMapNames := []string{}
		for _, cm := range currentConfigMaps {
			currentConfigMapNames = append(currentConfigMapNames, cm.Name)
		}
		return currentConfigMapNames, nil
	}
	newConfigMaps, err := makeRulesConfigMaps(p, newRules)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make rules ConfigMaps")
	}
	newConfigMapNames := []string{}
	for _, cm := range newConfigMaps {
		newConfigMapNames = append(newConfigMapNames, cm.Name)
	}
	if len(currentConfigMaps) == 0 {
		level.Debug(c.logger).Log("msg", "no PrometheusRule configmap found, creating new one", "namespace", p.Namespace, "prometheus", p.Name)
		for _, cm := range newConfigMaps {
			_, err = cClient.Create(&cm)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to create ConfigMap '%v'", cm.Name)
			}
		}
		return newConfigMapNames, nil
	}
	for _, cm := range currentConfigMaps {
		err := cClient.Delete(cm.Name, &metav1.DeleteOptions{})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to delete current ConfigMap '%v'", cm.Name)
		}
	}
	level.Debug(c.logger).Log("msg", "updating PrometheusRule", "namespace", p.Namespace, "prometheus", p.Name)
	for _, cm := range newConfigMaps {
		_, err = cClient.Create(&cm)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create new ConfigMap '%v'", cm.Name)
		}
	}
	return newConfigMapNames, nil
}
func prometheusRulesConfigMapSelector(prometheusName string) metav1.ListOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return metav1.ListOptions{LabelSelector: fmt.Sprintf("%v=%v", labelPrometheusName, prometheusName)}
}
func (c *Operator) selectRuleNamespaces(p *monitoringv1.Prometheus) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespaces := []string{}
	if p.Spec.RuleNamespaceSelector == nil {
		namespaces = append(namespaces, p.Namespace)
	} else {
		ruleNamespaceSelector, err := metav1.LabelSelectorAsSelector(p.Spec.RuleNamespaceSelector)
		if err != nil {
			return namespaces, errors.Wrap(err, "convert rule namespace label selector to selector")
		}
		namespaces, err = c.listMatchingNamespaces(ruleNamespaceSelector)
		if err != nil {
			return nil, err
		}
	}
	level.Debug(c.logger).Log("msg", "selected RuleNamespaces", "namespaces", strings.Join(namespaces, ","), "namespace", p.Namespace, "prometheus", p.Name)
	return namespaces, nil
}
func (c *Operator) selectRules(p *monitoringv1.Prometheus, namespaces []string) (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rules := map[string]string{}
	ruleSelector, err := metav1.LabelSelectorAsSelector(p.Spec.RuleSelector)
	if err != nil {
		return rules, errors.Wrap(err, "convert rule label selector to selector")
	}
	for _, ns := range namespaces {
		var marshalErr error
		err := cache.ListAllByNamespace(c.ruleInf.GetIndexer(), ns, ruleSelector, func(obj interface{}) {
			rule := obj.(*monitoringv1.PrometheusRule)
			content, err := yaml.Marshal(rule.Spec)
			if err != nil {
				marshalErr = err
				return
			}
			rules[fmt.Sprintf("%v-%v.yaml", rule.Namespace, rule.Name)] = string(content)
		})
		if err != nil {
			return nil, err
		}
		if marshalErr != nil {
			return nil, marshalErr
		}
	}
	ruleNames := []string{}
	for name := range rules {
		ruleNames = append(ruleNames, name)
	}
	level.Debug(c.logger).Log("msg", "selected Rules", "rules", strings.Join(ruleNames, ","), "namespace", p.Namespace, "prometheus", p.Name)
	return rules, nil
}
func makeRulesConfigMaps(p *monitoringv1.Prometheus, ruleFiles map[string]string) ([]v1.ConfigMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for filename, file := range ruleFiles {
		if len(file) > maxConfigMapDataSize {
			return nil, errors.Errorf("rule file '%v' is too large for a single Kubernetes ConfigMap", filename)
		}
	}
	buckets := []map[string]string{{}}
	currBucketIndex := 0
	fileNames := []string{}
	for n := range ruleFiles {
		fileNames = append(fileNames, n)
	}
	sort.Strings(fileNames)
	for _, filename := range fileNames {
		if bucketSize(buckets[currBucketIndex])+len(ruleFiles[filename]) > maxConfigMapDataSize {
			buckets = append(buckets, map[string]string{})
			currBucketIndex++
		}
		buckets[currBucketIndex][filename] = ruleFiles[filename]
	}
	ruleFileConfigMaps := []v1.ConfigMap{}
	for i, bucket := range buckets {
		cm := makeRulesConfigMap(p, bucket)
		cm.Name = cm.Name + "-" + strconv.Itoa(i)
		ruleFileConfigMaps = append(ruleFileConfigMaps, cm)
	}
	return ruleFileConfigMaps, nil
}
func bucketSize(bucket map[string]string) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	totalSize := 0
	for _, v := range bucket {
		totalSize += len(v)
	}
	return totalSize
}
func makeRulesConfigMap(p *monitoringv1.Prometheus, ruleFiles map[string]string) v1.ConfigMap {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	boolTrue := true
	labels := map[string]string{labelPrometheusName: p.Name}
	for k, v := range managedByOperatorLabels {
		labels[k] = v
	}
	return v1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: prometheusRuleConfigMapName(p.Name), Labels: labels, OwnerReferences: []metav1.OwnerReference{{APIVersion: p.APIVersion, BlockOwnerDeletion: &boolTrue, Controller: &boolTrue, Kind: p.Kind, Name: p.Name, UID: p.UID}}}, Data: ruleFiles}
}
func prometheusRuleConfigMapName(prometheusName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "prometheus-" + prometheusName + "-rulefiles"
}
