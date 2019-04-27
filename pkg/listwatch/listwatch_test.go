package listwatch

import (
	"sync"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var _ watch.Interface = &multiWatch{}

func setupMultiWatch(n int, t *testing.T, rvs ...string) ([]*watch.FakeWatcher, *multiWatch) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(rvs) == 0 {
		rvs = make([]string, n)
	}
	ws := make([]*watch.FakeWatcher, n)
	lws := make([]cache.ListerWatcher, n)
	for i := range ws {
		w := watch.NewFake()
		ws[i] = w
		lws[i] = &cache.ListWatch{WatchFunc: func(_ metav1.ListOptions) (watch.Interface, error) {
			return w, nil
		}}
	}
	m, err := newMultiWatch(lws, rvs, metav1.ListOptions{})
	if err != nil {
		t.Fatalf("failed to create new multiWatch: %v", err)
	}
	return ws, m
}
func TestNewMultiWatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected newMultiWatch to panic when number of resource versions is less than ListerWatchers")
			}
		}()
		_, _ = setupMultiWatch(2, t, "1")
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("newMultiWatch should not panic when number of resource versions matches ListerWatchers; got: %v", r)
			}
		}()
		_, _ = setupMultiWatch(2, t, "1", "2")
	}()
}
func TestMultiWatchResultChan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ws, m := setupMultiWatch(10, t)
	defer m.Stop()
	var events []watch.Event
	var wg sync.WaitGroup
	for _, w := range ws {
		w := w
		wg.Add(1)
		go func() {
			w.Add(&runtime.Unknown{})
		}()
	}
	go func() {
		for {
			event, ok := <-m.ResultChan()
			if !ok {
				break
			}
			events = append(events, event)
			wg.Done()
		}
	}()
	wg.Wait()
	if len(events) != len(ws) {
		t.Errorf("expected %d events but got %d", len(ws), len(events))
	}
}
func TestMultiWatchStop(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ws, m := setupMultiWatch(10, t)
	m.Stop()
	var stopped int
	for _, w := range ws {
		_, running := <-w.ResultChan()
		if !running && w.IsStopped() {
			stopped++
		}
	}
	if stopped != len(ws) {
		t.Errorf("expected %d watchers to be stopped but got %d", len(ws), stopped)
	}
	select {
	case <-m.stopped:
	default:
		t.Error("expected multiWatch to be stopped")
	}
	_, running := <-m.ResultChan()
	if running {
		t.Errorf("expected multiWatch chan to be closed")
	}
}

type mockListerWatcher struct {
	evCh	chan watch.Event
	stopped	bool
}

func (m *mockListerWatcher) List(options metav1.ListOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (m *mockListerWatcher) Watch(options metav1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m, nil
}
func (m *mockListerWatcher) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.stopped = true
}
func (m *mockListerWatcher) ResultChan() <-chan watch.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.evCh
}
func TestRacyMultiWatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	evCh := make(chan watch.Event)
	lw := &mockListerWatcher{evCh: evCh}
	mw, err := newMultiWatch([]cache.ListerWatcher{lw}, []string{"foo"}, metav1.ListOptions{})
	if err != nil {
		t.Error(err)
		return
	}
	evCh <- watch.Event{Type: "foo"}
	if got := <-mw.ResultChan(); got.Type != "foo" {
		t.Errorf("expected foo, got %s", got.Type)
		return
	}
	evCh <- watch.Event{Type: "bar"}
	mw.Stop()
	if got := lw.stopped; got != true {
		t.Errorf("expected watcher to be closed true, got %t", got)
	}
	mw.Stop()
	mw.Stop()
}
