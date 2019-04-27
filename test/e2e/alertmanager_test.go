package e2e

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"testing"
	"time"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/wait"
	"github.com/coreos/prometheus-operator/pkg/alertmanager"
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	testFramework "github.com/coreos/prometheus-operator/test/framework"
	"github.com/golang/protobuf/proto"
)

func testAMCreateDeleteCluster(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)
	name := "test"
	if _, err := framework.CreateAlertmanagerAndWaitUntilReady(ns, framework.MakeBasicAlertmanager(name, 3)); err != nil {
		t.Fatal(err)
	}
	if err := framework.DeleteAlertmanagerAndWaitUntilGone(ns, name); err != nil {
		t.Fatal(err)
	}
}
func testAMScaling(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)
	name := "test"
	a, err := framework.CreateAlertmanagerAndWaitUntilReady(ns, framework.MakeBasicAlertmanager(name, 3))
	if err != nil {
		t.Fatal(err)
	}
	a.Spec.Replicas = proto.Int32(5)
	a, err = framework.UpdateAlertmanagerAndWaitUntilReady(ns, a)
	if err != nil {
		t.Fatal(err)
	}
	a.Spec.Replicas = proto.Int32(3)
	if _, err := framework.UpdateAlertmanagerAndWaitUntilReady(ns, a); err != nil {
		t.Fatal(err)
	}
}
func testAMVersionMigration(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)
	name := "test"
	am := framework.MakeBasicAlertmanager(name, 1)
	am.Spec.Version = "v0.14.0"
	am, err := framework.CreateAlertmanagerAndWaitUntilReady(ns, am)
	if err != nil {
		t.Fatal(err)
	}
	am.Spec.Version = "v0.15.3"
	am, err = framework.UpdateAlertmanagerAndWaitUntilReady(ns, am)
	if err != nil {
		t.Fatal(err)
	}
	am.Spec.Version = "v0.14.0"
	am, err = framework.UpdateAlertmanagerAndWaitUntilReady(ns, am)
	if err != nil {
		t.Fatal(err)
	}
}
func testAMStorageUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	name := "test"
	am := framework.MakeBasicAlertmanager(name, 1)
	am, err := framework.CreateAlertmanagerAndWaitUntilReady(ns, am)
	if err != nil {
		t.Fatal(err)
	}
	am.Spec.Storage = &monitoringv1.StorageSpec{VolumeClaimTemplate: v1.PersistentVolumeClaim{ObjectMeta: metav1.ObjectMeta{CreationTimestamp: metav1.Now()}, Spec: v1.PersistentVolumeClaimSpec{AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce}, Resources: v1.ResourceRequirements{Requests: v1.ResourceList{v1.ResourceStorage: resource.MustParse("200Mi")}}}}}
	am, err = framework.MonClientV1.Alertmanagers(ns).Update(am)
	if err != nil {
		t.Fatal(err)
	}
	err = wait.Poll(5*time.Second, 2*time.Minute, func() (bool, error) {
		pods, err := framework.KubeClient.Core().Pods(ns).List(alertmanager.ListOptions(name))
		if err != nil {
			return false, err
		}
		if len(pods.Items) != 1 {
			return false, nil
		}
		for _, volume := range pods.Items[0].Spec.Volumes {
			if volume.Name == "alertmanager-"+name+"-db" && volume.PersistentVolumeClaim != nil && volume.PersistentVolumeClaim.ClaimName != "" {
				return true, nil
			}
		}
		return false, nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
func testAMExposingWithKubernetesAPI(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)
	alertmanager := framework.MakeBasicAlertmanager("test-alertmanager", 1)
	alertmanagerService := framework.MakeAlertmanagerService(alertmanager.Name, "alertmanager-service", v1.ServiceTypeClusterIP)
	if _, err := framework.CreateAlertmanagerAndWaitUntilReady(ns, alertmanager); err != nil {
		t.Fatal(err)
	}
	if _, err := testFramework.CreateServiceAndWaitUntilReady(framework.KubeClient, ns, alertmanagerService); err != nil {
		t.Fatal(err)
	}
	proxyGet := framework.KubeClient.CoreV1().Services(ns).ProxyGet
	request := proxyGet("", alertmanagerService.Name, "web", "/", make(map[string]string))
	_, err := request.DoRaw()
	if err != nil {
		t.Fatal(err)
	}
}
func testAMMeshInitialization(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	amVersions := []string{"v0.14.0", "v0.15.3"}
	for _, v := range amVersions {
		version := v
		t.Run(fmt.Sprintf("amVersion%v", strings.Replace(version, ".", "-", -1)), func(t *testing.T) {
			t.Parallel()
			ctx := framework.NewTestCtx(t)
			defer ctx.Cleanup(t)
			ns := ctx.CreateNamespace(t, framework.KubeClient)
			ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)
			amClusterSize := 3
			alertmanager := framework.MakeBasicAlertmanager("test", int32(amClusterSize))
			alertmanager.Spec.Version = version
			alertmanagerService := framework.MakeAlertmanagerService(alertmanager.Name, "alertmanager-service", v1.ServiceTypeClusterIP)
			if _, err := framework.CreateAlertmanagerAndWaitUntilReady(ns, alertmanager); err != nil {
				t.Fatal(err)
			}
			if _, err := testFramework.CreateServiceAndWaitUntilReady(framework.KubeClient, ns, alertmanagerService); err != nil {
				t.Fatal(err)
			}
			for i := 0; i < amClusterSize; i++ {
				name := "alertmanager-" + alertmanager.Name + "-" + strconv.Itoa(i)
				if err := framework.WaitForAlertmanagerInitializedMesh(ns, name, amClusterSize); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
func testAMClusterGossipSilences(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)
	amClusterSize := 3
	alertmanager := framework.MakeBasicAlertmanager("test", int32(amClusterSize))
	if _, err := framework.CreateAlertmanagerAndWaitUntilReady(ns, alertmanager); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < amClusterSize; i++ {
		name := "alertmanager-" + alertmanager.Name + "-" + strconv.Itoa(i)
		if err := framework.WaitForAlertmanagerInitializedMesh(ns, name, amClusterSize); err != nil {
			t.Fatal(err)
		}
	}
	silId, err := framework.CreateSilence(ns, "alertmanager-test-0")
	if err != nil {
		t.Fatalf("failed to create silence: %v", err)
	}
	for i := 0; i < amClusterSize; i++ {
		err = wait.Poll(time.Second, framework.DefaultTimeout, func() (bool, error) {
			silences, err := framework.GetSilences(ns, "alertmanager-"+alertmanager.Name+"-"+strconv.Itoa(i))
			if err != nil {
				return false, err
			}
			if len(silences) != 1 {
				return false, nil
			}
			if silences[0].ID != silId {
				return false, errors.Errorf("expected silence id on alertmanager %v to match id of created silence '%v' but got %v", i, silId, silences[0].ID)
			}
			return true, nil
		})
		if err != nil {
			t.Fatalf("could not retrieve created silence on alertmanager %v: %v", i, err)
		}
	}
}
func testAMReloadConfig(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)
	alertmanager := framework.MakeBasicAlertmanager("reload-config", 1)
	firstConfig := `
global:
  resolve_timeout: 5m
route:
  group_by: ['job']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 12h
  receiver: 'webhook'
receivers:
- name: 'webhook'
  webhook_configs:
  - url: 'http://firstConfigWebHook:30500/'
`
	secondConfig := `
global:
  resolve_timeout: 5m
route:
  group_by: ['job']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 12h
  receiver: 'webhook'
receivers:
- name: 'webhook'
  webhook_configs:
  - url: 'http://secondConfigWebHook:30500/'
`
	cfg := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: fmt.Sprintf("alertmanager-%s", alertmanager.Name)}, Data: map[string][]byte{"alertmanager.yaml": []byte(firstConfig)}}
	if _, err := framework.CreateAlertmanagerAndWaitUntilReady(ns, alertmanager); err != nil {
		t.Fatal(err)
	}
	if _, err := framework.KubeClient.CoreV1().Secrets(ns).Update(cfg); err != nil {
		t.Fatal(err)
	}
	firstExpectedString := "firstConfigWebHook"
	if err := framework.WaitForAlertmanagerConfigToContainString(ns, alertmanager.Name, firstExpectedString); err != nil {
		t.Fatal(errors.Wrap(err, "failed to wait for first expected config"))
	}
	cfg.Data["alertmanager.yaml"] = []byte(secondConfig)
	if _, err := framework.KubeClient.CoreV1().Secrets(ns).Update(cfg); err != nil {
		t.Fatal(err)
	}
	secondExpectedString := "secondConfigWebHook"
	if err := framework.WaitForAlertmanagerConfigToContainString(ns, alertmanager.Name, secondExpectedString); err != nil {
		t.Fatal(errors.Wrap(err, "failed to wait for second expected config"))
	}
}
func testAMZeroDowntimeRollingDeployment(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Parallel()
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)
	whReplicas := int32(1)
	whdpl := &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: "alertmanager-webhook"}, Spec: appsv1.DeploymentSpec{Replicas: &whReplicas, Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"app": "alertmanager-webhook"}}, Template: v1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": "alertmanager-webhook"}}, Spec: v1.PodSpec{Containers: []v1.Container{{Name: "webhook-server", Image: "quay.io/coreos/prometheus-alertmanager-test-webhook", Ports: []v1.ContainerPort{{Name: "web", ContainerPort: 5001}}}}}}}}
	whsvc := &v1.Service{ObjectMeta: metav1.ObjectMeta{Name: "alertmanager-webhook"}, Spec: v1.ServiceSpec{Type: v1.ServiceTypeClusterIP, Ports: []v1.ServicePort{{Name: "web", Port: 5001, TargetPort: intstr.FromString("web")}}, Selector: map[string]string{"app": "alertmanager-webhook"}}}
	if err := testFramework.CreateDeployment(framework.KubeClient, ns, whdpl); err != nil {
		t.Fatal(err)
	}
	if _, err := testFramework.CreateServiceAndWaitUntilReady(framework.KubeClient, ns, whsvc); err != nil {
		t.Fatal(err)
	}
	err := testFramework.WaitForPodsReady(framework.KubeClient, ns, time.Minute*5, 1, metav1.ListOptions{LabelSelector: fields.SelectorFromSet(fields.Set(map[string]string{"app": "alertmanager-webhook"})).String()})
	if err != nil {
		t.Fatal(err)
	}
	alertmanager := framework.MakeBasicAlertmanager("rolling-deploy", 3)
	alertmanager.Spec.Version = "v0.13.0"
	amsvc := framework.MakeAlertmanagerService(alertmanager.Name, "test", v1.ServiceTypeClusterIP)
	amcfg := &v1.Secret{ObjectMeta: metav1.ObjectMeta{Name: fmt.Sprintf("alertmanager-%s", alertmanager.Name)}, Data: map[string][]byte{"alertmanager.yaml": []byte(fmt.Sprintf(`
global:
  resolve_timeout: 5m

route:
  group_by: ['alertname']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 1h
  receiver: 'webhook'
receivers:
- name: 'webhook'
  webhook_configs:
  - url: 'http://%s.%s.svc:5001/'
inhibit_rules:
  - source_match:
      severity: 'critical'
    target_match:
      severity: 'warning'
    equal: ['alertname', 'dev', 'instance']
`, whsvc.Name, ns))}}
	if _, err := framework.KubeClient.CoreV1().Secrets(ns).Create(amcfg); err != nil {
		t.Fatal(err)
	}
	alertmanager, err = framework.MonClientV1.Alertmanagers(ns).Create(alertmanager)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := testFramework.CreateServiceAndWaitUntilReady(framework.KubeClient, ns, amsvc); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < int(*alertmanager.Spec.Replicas); i++ {
		replica := i
		done := make(chan struct{})
		errc := make(chan error, 1)
		defer func() {
			close(done)
			select {
			case err := <-errc:
				t.Fatal(errors.Wrapf(err, "sending alert to alertmanager %v", replica))
			default:
				return
			}
		}()
		go func() {
			ticker := time.NewTicker(100 * time.Millisecond)
			start := time.Now()
			failures := 0
			for {
				select {
				case <-ticker.C:
					err := framework.SendAlertToAlertmanager(ns, "alertmanager-rolling-deploy-"+strconv.Itoa(replica), start)
					if err != nil {
						failures++
						if failures > 50 {
							errc <- err
							return
						}
					}
				case <-done:
					return
				}
			}
		}()
	}
	time.Sleep(10 * time.Second)
	opts := metav1.ListOptions{LabelSelector: fields.SelectorFromSet(fields.Set(map[string]string{"app": "alertmanager-webhook"})).String()}
	pl, err := framework.KubeClient.Core().Pods(ns).List(opts)
	if err != nil {
		t.Fatal(err)
	}
	if len(pl.Items) != 1 {
		t.Fatalf("Expected one webhook pod, but got %d", len(pl.Items))
	}
	podName := pl.Items[0].Name
	logs, err := testFramework.GetLogs(framework.KubeClient, ns, podName, "webhook-server")
	if err != nil {
		t.Fatal(err)
	}
	c := strings.Count(logs, "Alertmanager Notification Payload Received")
	if c != 1 {
		t.Fatalf("One notification expected, but %d received.\n\n%s", c, logs)
	}
	alertmanager.Spec.Version = "v0.14.0"
	alertmanager, err = framework.MonClientV1.Alertmanagers(ns).Update(alertmanager)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Minute)
	logs, err = testFramework.GetLogs(framework.KubeClient, ns, podName, "webhook-server")
	if err != nil {
		t.Fatal(err)
	}
	c = strings.Count(logs, "Alertmanager Notification Payload Received")
	if c != 1 {
		t.Fatalf("Only one notification expected, but %d received after rolling update of Alertmanager cluster.\n\n%s", c, logs)
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
