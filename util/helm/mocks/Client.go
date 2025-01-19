// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import (
	helm "github.com/argoproj/argo-cd/v3/util/helm"
	io "github.com/argoproj/argo-cd/v3/util/io"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CleanChartCache provides a mock function with given fields: chart, version, project
func (_m *Client) CleanChartCache(chart string, version string, project string) error {
	ret := _m.Called(chart, version, project)

	if len(ret) == 0 {
		panic("no return value specified for CleanChartCache")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(chart, version, project)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtractChart provides a mock function with given fields: chart, version, project, passCredentials, manifestMaxExtractedSize, disableManifestMaxExtractedSize
func (_m *Client) ExtractChart(chart string, version string, project string, passCredentials bool, manifestMaxExtractedSize int64, disableManifestMaxExtractedSize bool) (string, io.Closer, error) {
	ret := _m.Called(chart, version, project, passCredentials, manifestMaxExtractedSize, disableManifestMaxExtractedSize)

	if len(ret) == 0 {
		panic("no return value specified for ExtractChart")
	}

	var r0 string
	var r1 io.Closer
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, string, bool, int64, bool) (string, io.Closer, error)); ok {
		return rf(chart, version, project, passCredentials, manifestMaxExtractedSize, disableManifestMaxExtractedSize)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, bool, int64, bool) string); ok {
		r0 = rf(chart, version, project, passCredentials, manifestMaxExtractedSize, disableManifestMaxExtractedSize)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string, bool, int64, bool) io.Closer); ok {
		r1 = rf(chart, version, project, passCredentials, manifestMaxExtractedSize, disableManifestMaxExtractedSize)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(io.Closer)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string, string, bool, int64, bool) error); ok {
		r2 = rf(chart, version, project, passCredentials, manifestMaxExtractedSize, disableManifestMaxExtractedSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetIndex provides a mock function with given fields: noCache, maxIndexSize
func (_m *Client) GetIndex(noCache bool, maxIndexSize int64) (*helm.Index, error) {
	ret := _m.Called(noCache, maxIndexSize)

	if len(ret) == 0 {
		panic("no return value specified for GetIndex")
	}

	var r0 *helm.Index
	var r1 error
	if rf, ok := ret.Get(0).(func(bool, int64) (*helm.Index, error)); ok {
		return rf(noCache, maxIndexSize)
	}
	if rf, ok := ret.Get(0).(func(bool, int64) *helm.Index); ok {
		r0 = rf(noCache, maxIndexSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*helm.Index)
		}
	}

	if rf, ok := ret.Get(1).(func(bool, int64) error); ok {
		r1 = rf(noCache, maxIndexSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields: chart, noCache
func (_m *Client) GetTags(chart string, noCache bool) (*helm.TagsList, error) {
	ret := _m.Called(chart, noCache)

	if len(ret) == 0 {
		panic("no return value specified for GetTags")
	}

	var r0 *helm.TagsList
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (*helm.TagsList, error)); ok {
		return rf(chart, noCache)
	}
	if rf, ok := ret.Get(0).(func(string, bool) *helm.TagsList); ok {
		r0 = rf(chart, noCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*helm.TagsList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(chart, noCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestHelmOCI provides a mock function with no fields
func (_m *Client) TestHelmOCI() (bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TestHelmOCI")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
