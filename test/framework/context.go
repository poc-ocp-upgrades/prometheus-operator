package framework

import (
	"strconv"
	"strings"
	"testing"
	"time"
	"golang.org/x/sync/errgroup"
)

type TestCtx struct {
	ID		string
	cleanUpFns	[]finalizerFn
}
type finalizerFn func() error

func (f *Framework) NewTestCtx(t *testing.T) TestCtx {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := strings.TrimPrefix(strings.Replace(strings.ToLower(t.Name()), "/", "-", -1), "test")
	id := prefix + "-" + strconv.FormatInt(time.Now().Unix(), 36)
	return TestCtx{ID: id}
}
func (ctx *TestCtx) GetObjID() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ctx.ID + "-" + strconv.Itoa(len(ctx.cleanUpFns))
}
func (ctx *TestCtx) Cleanup(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var eg errgroup.Group
	for i := len(ctx.cleanUpFns) - 1; i >= 0; i-- {
		eg.Go(ctx.cleanUpFns[i])
	}
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
func (ctx *TestCtx) AddFinalizerFn(fn finalizerFn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx.cleanUpFns = append(ctx.cleanUpFns, fn)
}
