package asset

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type parentsAsset struct{ x int }

func (a *parentsAsset) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "parents-asset"
}
func (a *parentsAsset) Dependencies() []Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []Asset{}
}
func (a *parentsAsset) Generate(Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func TestParentsGetPointer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	origAsset := &parentsAsset{x: 1}
	parents := Parents{}
	parents.Add(origAsset)
	retrievedAsset := &parentsAsset{}
	parents.Get(retrievedAsset)
	assert.Equal(t, 1, retrievedAsset.x)
}
