// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	powercontrol "go.anx.io/go-anxcloud/pkg/vsphere/powercontrol"
)

// PowerControl is an autogenerated mock type for the API type
type PowerControl struct {
	mock.Mock
}

// AwaitCompletion provides a mock function with given fields: ctx, vmID, taskID
func (_m *PowerControl) AwaitCompletion(ctx context.Context, vmID string, taskID string) error {
	ret := _m.Called(ctx, vmID, taskID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, vmID, taskID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, vmIdentifier
func (_m *PowerControl) Get(ctx context.Context, vmIdentifier string) (powercontrol.State, error) {
	ret := _m.Called(ctx, vmIdentifier)

	var r0 powercontrol.State
	if rf, ok := ret.Get(0).(func(context.Context, string) powercontrol.State); ok {
		r0 = rf(ctx, vmIdentifier)
	} else {
		r0 = ret.Get(0).(powercontrol.State)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, vmIdentifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, vmIdentifier, request
func (_m *PowerControl) Set(ctx context.Context, vmIdentifier string, request powercontrol.Request) (powercontrol.Task, error) {
	ret := _m.Called(ctx, vmIdentifier, request)

	var r0 powercontrol.Task
	if rf, ok := ret.Get(0).(func(context.Context, string, powercontrol.Request) powercontrol.Task); ok {
		r0 = rf(ctx, vmIdentifier, request)
	} else {
		r0 = ret.Get(0).(powercontrol.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, powercontrol.Request) error); ok {
		r1 = rf(ctx, vmIdentifier, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
