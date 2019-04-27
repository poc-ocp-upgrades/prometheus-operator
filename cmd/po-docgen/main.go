package main

import (
	"os"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch os.Args[1] {
	case "api":
		printAPIDocs(os.Args[2])
	case "compatibility":
		printCompatMatrixDocs()
	}
}
