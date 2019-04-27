package main

import (
	"fmt"
	"github.com/coreos/prometheus-operator/pkg/prometheus"
)

func printCompatMatrixDocs() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Println(`<br>
<div class="alert alert-info" role="alert">
    <i class="fa fa-exclamation-triangle"></i><b> Note:</b> Starting with v0.12.0, Prometheus Operator requires use of Kubernetes v1.7.x and up.
</div>

# Compatibility

The Prometheus Operator supports a number of Kubernetes and Prometheus releases.

## Kubernetes

The Prometheus Operator uses client-go to communicate with Kubernetes clusters. The supported Kubernetes cluster version is determined by client-go. The compatibility matrix for client-go and Kubernetes clusters can be found [here](https://github.com/kubernetes/client-go#compatibility-matrix). All additional compatibility is only best effort, or happens to still/already be supported. The currently used client-go version is "v4.0.0-beta.0".

Due to the use of CustomResourceDefinitions Kubernetes >= v1.7.0 is required.

## Prometheus

The versions of Prometheus compatible to be run with the Prometheus Operator are:`)
	fmt.Println("")
	for _, v := range prometheus.CompatibilityMatrix {
		fmt.Printf("* %s\n", v)
	}
}
