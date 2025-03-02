// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/github/client (interfaces: RepositoriesService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v48/github"
)

// MockRepositoriesService is a mock of RepositoriesService interface.
type MockRepositoriesService struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoriesServiceMockRecorder
}

// MockRepositoriesServiceMockRecorder is the mock recorder for MockRepositoriesService.
type MockRepositoriesServiceMockRecorder struct {
	mock *MockRepositoriesService
}

// NewMockRepositoriesService creates a new mock instance.
func NewMockRepositoriesService(ctrl *gomock.Controller) *MockRepositoriesService {
	mock := &MockRepositoriesService{ctrl: ctrl}
	mock.recorder = &MockRepositoriesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoriesService) EXPECT() *MockRepositoriesServiceMockRecorder {
	return m.recorder
}

// GetContents mocks base method.
func (m *MockRepositoriesService) GetContents(arg0 context.Context, arg1, arg2, arg3 string, arg4 *github.RepositoryContentGetOptions) (*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContents", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*github.RepositoryContent)
	ret1, _ := ret[1].([]*github.RepositoryContent)
	ret2, _ := ret[2].(*github.Response)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetContents indicates an expected call of GetContents.
func (mr *MockRepositoriesServiceMockRecorder) GetContents(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContents", reflect.TypeOf((*MockRepositoriesService)(nil).GetContents), arg0, arg1, arg2, arg3, arg4)
}

// ListByOrg mocks base method.
func (m *MockRepositoriesService) ListByOrg(arg0 context.Context, arg1 string, arg2 *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByOrg", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Repository)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByOrg indicates an expected call of ListByOrg.
func (mr *MockRepositoriesServiceMockRecorder) ListByOrg(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByOrg", reflect.TypeOf((*MockRepositoriesService)(nil).ListByOrg), arg0, arg1, arg2)
}

// ListReleaseAssets mocks base method.
func (m *MockRepositoriesService) ListReleaseAssets(arg0 context.Context, arg1, arg2 string, arg3 int64, arg4 *github.ListOptions) ([]*github.ReleaseAsset, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleaseAssets", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*github.ReleaseAsset)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListReleaseAssets indicates an expected call of ListReleaseAssets.
func (mr *MockRepositoriesServiceMockRecorder) ListReleaseAssets(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleaseAssets", reflect.TypeOf((*MockRepositoriesService)(nil).ListReleaseAssets), arg0, arg1, arg2, arg3, arg4)
}

// ListReleases mocks base method.
func (m *MockRepositoriesService) ListReleases(arg0 context.Context, arg1, arg2 string, arg3 *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleases", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.RepositoryRelease)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListReleases indicates an expected call of ListReleases.
func (mr *MockRepositoriesServiceMockRecorder) ListReleases(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleases", reflect.TypeOf((*MockRepositoriesService)(nil).ListReleases), arg0, arg1, arg2, arg3)
}

// ListTrafficClones mocks base method.
func (m *MockRepositoriesService) ListTrafficClones(arg0 context.Context, arg1, arg2 string, arg3 *github.TrafficBreakdownOptions) (*github.TrafficClones, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrafficClones", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.TrafficClones)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTrafficClones indicates an expected call of ListTrafficClones.
func (mr *MockRepositoriesServiceMockRecorder) ListTrafficClones(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficClones", reflect.TypeOf((*MockRepositoriesService)(nil).ListTrafficClones), arg0, arg1, arg2, arg3)
}

// ListTrafficPaths mocks base method.
func (m *MockRepositoriesService) ListTrafficPaths(arg0 context.Context, arg1, arg2 string) ([]*github.TrafficPath, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrafficPaths", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.TrafficPath)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTrafficPaths indicates an expected call of ListTrafficPaths.
func (mr *MockRepositoriesServiceMockRecorder) ListTrafficPaths(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficPaths", reflect.TypeOf((*MockRepositoriesService)(nil).ListTrafficPaths), arg0, arg1, arg2)
}

// ListTrafficReferrers mocks base method.
func (m *MockRepositoriesService) ListTrafficReferrers(arg0 context.Context, arg1, arg2 string) ([]*github.TrafficReferrer, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrafficReferrers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.TrafficReferrer)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTrafficReferrers indicates an expected call of ListTrafficReferrers.
func (mr *MockRepositoriesServiceMockRecorder) ListTrafficReferrers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficReferrers", reflect.TypeOf((*MockRepositoriesService)(nil).ListTrafficReferrers), arg0, arg1, arg2)
}

// ListTrafficViews mocks base method.
func (m *MockRepositoriesService) ListTrafficViews(arg0 context.Context, arg1, arg2 string, arg3 *github.TrafficBreakdownOptions) (*github.TrafficViews, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrafficViews", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.TrafficViews)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTrafficViews indicates an expected call of ListTrafficViews.
func (mr *MockRepositoriesServiceMockRecorder) ListTrafficViews(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficViews", reflect.TypeOf((*MockRepositoriesService)(nil).ListTrafficViews), arg0, arg1, arg2, arg3)
}
