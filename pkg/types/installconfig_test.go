package types

import (
	"sort"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPlatformNamesSorted(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sorted := make([]string, len(PlatformNames))
	copy(sorted, PlatformNames)
	sort.Strings(sorted)
	assert.Equal(t, sorted, PlatformNames)
}
