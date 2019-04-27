package framework

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	monitoringclient "github.com/coreos/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	"github.com/coreos/prometheus-operator/pkg/k8sutil"
	"github.com/pkg/errors"
)

type Framework struct {
	KubeClient	kubernetes.Interface
	MonClientV1	monitoringclient.MonitoringV1Interface
	HTTPClient	*http.Client
	MasterHost	string
	DefaultTimeout	time.Duration
}

func New(kubeconfig, opImage string) (*Framework, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "build config from flags failed")
	}
	cli, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "creating new kube-client failed")
	}
	httpc := cli.CoreV1().RESTClient().(*rest.RESTClient).Client
	if err != nil {
		return nil, errors.Wrap(err, "creating http-client failed")
	}
	mClientV1, err := monitoringclient.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "creating v1 monitoring client failed")
	}
	f := &Framework{MasterHost: config.Host, KubeClient: cli, MonClientV1: mClientV1, HTTPClient: httpc, DefaultTimeout: time.Minute}
	return f, nil
}
func (f *Framework) CreatePrometheusOperator(ns, opImage string, namespacesToWatch []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := CreateServiceAccount(f.KubeClient, ns, "../../example/rbac/prometheus-operator/prometheus-operator-service-account.yaml")
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return errors.Wrap(err, "failed to create prometheus operator service account")
	}
	if err := CreateClusterRole(f.KubeClient, "../../example/rbac/prometheus-operator/prometheus-operator-cluster-role.yaml"); err != nil && !apierrors.IsAlreadyExists(err) {
		return errors.Wrap(err, "failed to create prometheus cluster role")
	}
	if _, err := CreateClusterRoleBinding(f.KubeClient, ns, "../../example/rbac/prometheus-operator/prometheus-operator-cluster-role-binding.yaml"); err != nil && !apierrors.IsAlreadyExists(err) {
		return errors.Wrap(err, "failed to create prometheus cluster role binding")
	}
	deploy, err := MakeDeployment("../../example/rbac/prometheus-operator/prometheus-operator-deployment.yaml")
	if err != nil {
		return err
	}
	if opImage != "" {
		deploy.Spec.Template.Spec.Containers[0].Image = opImage
		repoAndTag := strings.Split(opImage, ":")
		if len(repoAndTag) != 2 {
			return errors.Errorf("expected operator image '%v' split by colon to result in two substrings but got '%v'", opImage, repoAndTag)
		}
		for i, arg := range deploy.Spec.Template.Spec.Containers[0].Args {
			if strings.Contains(arg, "--prometheus-config-reloader=") {
				deploy.Spec.Template.Spec.Containers[0].Args[i] = "--prometheus-config-reloader=" + "quay.io/coreos/prometheus-config-reloader:" + repoAndTag[1]
			}
		}
	}
	deploy.Spec.Template.Spec.Containers[0].Args = append(deploy.Spec.Template.Spec.Containers[0].Args, "--log-level=all")
	for _, ns := range namespacesToWatch {
		deploy.Spec.Template.Spec.Containers[0].Args = append(deploy.Spec.Template.Spec.Containers[0].Args, fmt.Sprintf("--namespaces=%v", ns))
	}
	err = CreateDeployment(f.KubeClient, ns, deploy)
	if err != nil {
		return err
	}
	opts := metav1.ListOptions{LabelSelector: fields.SelectorFromSet(fields.Set(deploy.Spec.Template.ObjectMeta.Labels)).String()}
	err = WaitForPodsReady(f.KubeClient, ns, f.DefaultTimeout, 1, opts)
	if err != nil {
		return errors.Wrap(err, "failed to wait for prometheus operator to become ready")
	}
	err = k8sutil.WaitForCRDReady(func(opts metav1.ListOptions) (runtime.Object, error) {
		return f.MonClientV1.Prometheuses(v1.NamespaceAll).List(opts)
	})
	if err != nil {
		return errors.Wrap(err, "Prometheus CRD not ready: %v\n")
	}
	err = k8sutil.WaitForCRDReady(func(opts metav1.ListOptions) (runtime.Object, error) {
		return f.MonClientV1.ServiceMonitors(v1.NamespaceAll).List(opts)
	})
	if err != nil {
		return errors.Wrap(err, "ServiceMonitor CRD not ready: %v\n")
	}
	err = k8sutil.WaitForCRDReady(func(opts metav1.ListOptions) (runtime.Object, error) {
		return f.MonClientV1.PrometheusRules(v1.NamespaceAll).List(opts)
	})
	if err != nil {
		return errors.Wrap(err, "PrometheusRule CRD not ready: %v\n")
	}
	err = k8sutil.WaitForCRDReady(func(opts metav1.ListOptions) (runtime.Object, error) {
		return f.MonClientV1.Alertmanagers(v1.NamespaceAll).List(opts)
	})
	if err != nil {
		return errors.Wrap(err, "Alertmanager CRD not ready: %v\n")
	}
	return nil
}
func (ctx *TestCtx) SetupPrometheusRBAC(t *testing.T, ns string, kubeClient kubernetes.Interface) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := CreateClusterRole(kubeClient, "../../example/rbac/prometheus/prometheus-cluster-role.yaml"); err != nil && !apierrors.IsAlreadyExists(err) {
		t.Fatalf("failed to create prometheus cluster role: %v", err)
	}
	if finalizerFn, err := CreateServiceAccount(kubeClient, ns, "../../example/rbac/prometheus/prometheus-service-account.yaml"); err != nil {
		t.Fatal(errors.Wrap(err, "failed to create prometheus service account"))
	} else {
		ctx.AddFinalizerFn(finalizerFn)
	}
	if finalizerFn, err := CreateRoleBinding(kubeClient, ns, "../framework/ressources/prometheus-role-binding.yml"); err != nil {
		t.Fatal(errors.Wrap(err, "failed to create prometheus role binding"))
	} else {
		ctx.AddFinalizerFn(finalizerFn)
	}
}
func (ctx *TestCtx) SetupPrometheusRBACGlobal(t *testing.T, ns string, kubeClient kubernetes.Interface) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := CreateClusterRole(kubeClient, "../../example/rbac/prometheus/prometheus-cluster-role.yaml"); err != nil && !apierrors.IsAlreadyExists(err) {
		t.Fatalf("failed to create prometheus cluster role: %v", err)
	}
	if finalizerFn, err := CreateServiceAccount(kubeClient, ns, "../../example/rbac/prometheus/prometheus-service-account.yaml"); err != nil {
		t.Fatal(errors.Wrap(err, "failed to create prometheus service account"))
	} else {
		ctx.AddFinalizerFn(finalizerFn)
	}
	if finalizerFn, err := CreateClusterRoleBinding(kubeClient, ns, "../../example/rbac/prometheus/prometheus-cluster-role-binding.yaml"); err != nil {
		t.Fatal(errors.Wrap(err, "failed to create prometheus cluster role binding"))
	} else {
		ctx.AddFinalizerFn(finalizerFn)
	}
}
