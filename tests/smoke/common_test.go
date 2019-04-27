package smoke

import (
	"fmt"
	"testing"
	"time"
)

func testCommon(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Run("APIAvailable", testAPIAvailable)
}
func testAPIAvailable(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wait := 15 * time.Minute
	err := retry(apiAvailable, t, 10*time.Second, wait)
	if err != nil {
		t.Fatalf("Failed to connect to API server in %v.", wait)
	}
	t.Log("API server is available.")
}
func apiAvailable(t *testing.T) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, _ := newClient(t)
	_, err := client.ServerVersion()
	if err != nil {
		return fmt.Errorf("failed to connect to API server: %v", err)
	}
	return nil
}
