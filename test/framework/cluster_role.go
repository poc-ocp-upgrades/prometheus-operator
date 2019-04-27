package framework

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

func CreateClusterRole(kubeClient kubernetes.Interface, relativePath string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterRole, err := parseClusterRoleYaml(relativePath)
	if err != nil {
		return err
	}
	_, err = kubeClient.RbacV1().ClusterRoles().Get(clusterRole.Name, metav1.GetOptions{})
	if err == nil {
		_, err = kubeClient.RbacV1().ClusterRoles().Update(clusterRole)
		if err != nil {
			return err
		}
	} else {
		_, err = kubeClient.RbacV1().ClusterRoles().Create(clusterRole)
		if err != nil {
			return err
		}
	}
	return nil
}
func DeleteClusterRole(kubeClient kubernetes.Interface, relativePath string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterRole, err := parseClusterRoleYaml(relativePath)
	if err != nil {
		return err
	}
	return kubeClient.RbacV1().ClusterRoles().Delete(clusterRole.Name, &metav1.DeleteOptions{})
}
func parseClusterRoleYaml(relativePath string) (*rbacv1.ClusterRole, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	manifest, err := PathToOSFile(relativePath)
	if err != nil {
		return nil, err
	}
	clusterRole := rbacv1.ClusterRole{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&clusterRole); err != nil {
		return nil, err
	}
	return &clusterRole, nil
}
