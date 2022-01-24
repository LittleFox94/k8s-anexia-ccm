// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	acl "go.anx.io/go-anxcloud/pkg/lbaas/acl"
	backend "go.anx.io/go-anxcloud/pkg/lbaas/backend"

	bind "go.anx.io/go-anxcloud/pkg/lbaas/bind"

	frontend "go.anx.io/go-anxcloud/pkg/lbaas/frontend"

	loadbalancer "go.anx.io/go-anxcloud/pkg/lbaas/loadbalancer"

	mock "github.com/stretchr/testify/mock"

	server "go.anx.io/go-anxcloud/pkg/lbaas/server"
)

// LBaaS is an autogenerated mock type for the API type
type LBaaS struct {
	mock.Mock
}

// ACL provides a mock function with given fields:
func (_m *LBaaS) ACL() acl.API {
	ret := _m.Called()

	var r0 acl.API
	if rf, ok := ret.Get(0).(func() acl.API); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(acl.API)
		}
	}

	return r0
}

// Backend provides a mock function with given fields:
func (_m *LBaaS) Backend() backend.API {
	ret := _m.Called()

	var r0 backend.API
	if rf, ok := ret.Get(0).(func() backend.API); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(backend.API)
		}
	}

	return r0
}

// Bind provides a mock function with given fields:
func (_m *LBaaS) Bind() bind.API {
	ret := _m.Called()

	var r0 bind.API
	if rf, ok := ret.Get(0).(func() bind.API); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bind.API)
		}
	}

	return r0
}

// Frontend provides a mock function with given fields:
func (_m *LBaaS) Frontend() frontend.API {
	ret := _m.Called()

	var r0 frontend.API
	if rf, ok := ret.Get(0).(func() frontend.API); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(frontend.API)
		}
	}

	return r0
}

// LoadBalancer provides a mock function with given fields:
func (_m *LBaaS) LoadBalancer() loadbalancer.API {
	ret := _m.Called()

	var r0 loadbalancer.API
	if rf, ok := ret.Get(0).(func() loadbalancer.API); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(loadbalancer.API)
		}
	}

	return r0
}

// Server provides a mock function with given fields:
func (_m *LBaaS) Server() server.API {
	ret := _m.Called()

	var r0 server.API
	if rf, ok := ret.Get(0).(func() server.API); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(server.API)
		}
	}

	return r0
}
