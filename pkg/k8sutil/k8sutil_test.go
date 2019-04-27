package k8sutil

import (
	"strings"
	"testing"
	"k8s.io/apimachinery/pkg/util/validation"
)

func Test_SanitizeVolumeName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		expected	string
	}{{name: "@$!!@$%!#$%!#$%!#$!#$%%$#@!#", expected: ""}, {name: "NAME", expected: "name"}, {name: "foo--", expected: "foo"}, {name: "foo^%#$bar", expected: "foo-bar"}, {name: "fOo^%#$bar", expected: "foo-bar"}, {name: strings.Repeat("a", validation.DNS1123LabelMaxLength*2), expected: strings.Repeat("a", validation.DNS1123LabelMaxLength)}}
	for i, c := range cases {
		out := SanitizeVolumeName(c.name)
		if c.expected != out {
			t.Errorf("expected test case %d to be %q but got %q", i, c.expected, out)
		}
	}
}
