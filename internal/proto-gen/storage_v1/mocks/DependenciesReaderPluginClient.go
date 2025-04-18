// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	storage_v1 "github.com/jaegertracing/jaeger/internal/proto-gen/storage_v1"
)

// DependenciesReaderPluginClient is an autogenerated mock type for the DependenciesReaderPluginClient type
type DependenciesReaderPluginClient struct {
	mock.Mock
}

// GetDependencies provides a mock function with given fields: ctx, in, opts
func (_m *DependenciesReaderPluginClient) GetDependencies(ctx context.Context, in *storage_v1.GetDependenciesRequest, opts ...grpc.CallOption) (*storage_v1.GetDependenciesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDependencies")
	}

	var r0 *storage_v1.GetDependenciesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storage_v1.GetDependenciesRequest, ...grpc.CallOption) (*storage_v1.GetDependenciesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storage_v1.GetDependenciesRequest, ...grpc.CallOption) *storage_v1.GetDependenciesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage_v1.GetDependenciesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storage_v1.GetDependenciesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDependenciesReaderPluginClient creates a new instance of DependenciesReaderPluginClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDependenciesReaderPluginClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DependenciesReaderPluginClient {
	mock := &DependenciesReaderPluginClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
