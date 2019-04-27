package asset

import (
	"reflect"
)

type Parents map[reflect.Type]Asset

func (p Parents) Add(assets ...Asset) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, a := range assets {
		p[reflect.TypeOf(a)] = a
	}
}
func (p Parents) Get(assets ...Asset) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, a := range assets {
		reflect.ValueOf(a).Elem().Set(reflect.ValueOf(p[reflect.TypeOf(a)]).Elem())
	}
}
