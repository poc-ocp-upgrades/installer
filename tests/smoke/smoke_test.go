package smoke

import (
	"errors"
	"flag"
	"os"
	"testing"
	"time"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	kubeconfigEnv		= "SMOKE_KUBECONFIG"
	apiServerSelector	= "k8s-app=kube-apiserver"
	kubeSystemNamespace	= "kube-system"
	tectonicSystemNamespace	= "tectonic-system"
)

var (
	runClusterTests bool
)

func TestMain(m *testing.M) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.BoolVar(&runClusterTests, "cluster", false, "run cluster tests (default false)")
	flag.Parse()
	os.Exit(m.Run())
}
func Test(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Run("Common", testCommon)
	if runClusterTests {
		t.Run("Cluster", testCluster)
	}
}
func newClient(t *testing.T) (*kubernetes.Clientset, clientcmd.ClientConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	kcfgPath := os.Getenv(kubeconfigEnv)
	if len(kcfgPath) == 0 {
		t.Fatalf("no kubeconfig path in environment variable %s", kubeconfigEnv)
	}
	rules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kcfgPath}
	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	restConfig, err := cfg.ClientConfig()
	if err != nil {
		t.Fatalf("could not create client config: %v", err)
	}
	cs, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		t.Fatalf("could not create clientset: %v", err)
	}
	return cs, cfg
}
func stopped(done <-chan struct{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func sleepOrDone(sleep time.Duration, done <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case <-time.After(sleep):
		return
	case <-done:
		return
	}
}

type retriable func(t *testing.T) error

func timeout(d time.Duration) <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := time.NewTimer(d)
	c := make(chan struct{})
	go func() {
		<-t.C
		close(c)
	}()
	return c
}
func retry(r retriable, t *testing.T, interval, max time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	done := timeout(max)
	for !stopped(done) {
		err := r(t)
		if err == nil {
			return nil
		}
		t.Logf("failed with error: %v", err)
		t.Logf("retrying in %v", interval)
		sleepOrDone(interval, done)
	}
	return errors.New("timed out while retrying")
}
