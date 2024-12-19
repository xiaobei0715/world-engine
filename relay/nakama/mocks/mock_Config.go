// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/heroiclabs/nakama-common/runtime"
	mock "github.com/stretchr/testify/mock"
)

// MockConfig is an autogenerated mock type for the Config type
type MockConfig struct {
	mock.Mock
}

type MockConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConfig) EXPECT() *MockConfig_Expecter {
	return &MockConfig_Expecter{mock: &_m.Mock}
}

// GetGoogleAuth provides a mock function with no fields
func (_m *MockConfig) GetGoogleAuth() runtime.GoogleAuthConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGoogleAuth")
	}

	var r0 runtime.GoogleAuthConfig
	if rf, ok := ret.Get(0).(func() runtime.GoogleAuthConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.GoogleAuthConfig)
		}
	}

	return r0
}

// MockConfig_GetGoogleAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGoogleAuth'
type MockConfig_GetGoogleAuth_Call struct {
	*mock.Call
}

// GetGoogleAuth is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetGoogleAuth() *MockConfig_GetGoogleAuth_Call {
	return &MockConfig_GetGoogleAuth_Call{Call: _e.mock.On("GetGoogleAuth")}
}

func (_c *MockConfig_GetGoogleAuth_Call) Run(run func()) *MockConfig_GetGoogleAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetGoogleAuth_Call) Return(_a0 runtime.GoogleAuthConfig) *MockConfig_GetGoogleAuth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetGoogleAuth_Call) RunAndReturn(run func() runtime.GoogleAuthConfig) *MockConfig_GetGoogleAuth_Call {
	_c.Call.Return(run)
	return _c
}

// GetIAP provides a mock function with no fields
func (_m *MockConfig) GetIAP() runtime.IAPConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetIAP")
	}

	var r0 runtime.IAPConfig
	if rf, ok := ret.Get(0).(func() runtime.IAPConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.IAPConfig)
		}
	}

	return r0
}

// MockConfig_GetIAP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIAP'
type MockConfig_GetIAP_Call struct {
	*mock.Call
}

// GetIAP is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetIAP() *MockConfig_GetIAP_Call {
	return &MockConfig_GetIAP_Call{Call: _e.mock.On("GetIAP")}
}

func (_c *MockConfig_GetIAP_Call) Run(run func()) *MockConfig_GetIAP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetIAP_Call) Return(_a0 runtime.IAPConfig) *MockConfig_GetIAP_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetIAP_Call) RunAndReturn(run func() runtime.IAPConfig) *MockConfig_GetIAP_Call {
	_c.Call.Return(run)
	return _c
}

// GetLogger provides a mock function with no fields
func (_m *MockConfig) GetLogger() runtime.LoggerConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLogger")
	}

	var r0 runtime.LoggerConfig
	if rf, ok := ret.Get(0).(func() runtime.LoggerConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.LoggerConfig)
		}
	}

	return r0
}

// MockConfig_GetLogger_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLogger'
type MockConfig_GetLogger_Call struct {
	*mock.Call
}

// GetLogger is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetLogger() *MockConfig_GetLogger_Call {
	return &MockConfig_GetLogger_Call{Call: _e.mock.On("GetLogger")}
}

func (_c *MockConfig_GetLogger_Call) Run(run func()) *MockConfig_GetLogger_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetLogger_Call) Return(_a0 runtime.LoggerConfig) *MockConfig_GetLogger_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetLogger_Call) RunAndReturn(run func() runtime.LoggerConfig) *MockConfig_GetLogger_Call {
	_c.Call.Return(run)
	return _c
}

// GetName provides a mock function with no fields
func (_m *MockConfig) GetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockConfig_GetName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetName'
type MockConfig_GetName_Call struct {
	*mock.Call
}

// GetName is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetName() *MockConfig_GetName_Call {
	return &MockConfig_GetName_Call{Call: _e.mock.On("GetName")}
}

func (_c *MockConfig_GetName_Call) Run(run func()) *MockConfig_GetName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetName_Call) Return(_a0 string) *MockConfig_GetName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetName_Call) RunAndReturn(run func() string) *MockConfig_GetName_Call {
	_c.Call.Return(run)
	return _c
}

// GetRuntime provides a mock function with no fields
func (_m *MockConfig) GetRuntime() runtime.RuntimeConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRuntime")
	}

	var r0 runtime.RuntimeConfig
	if rf, ok := ret.Get(0).(func() runtime.RuntimeConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.RuntimeConfig)
		}
	}

	return r0
}

// MockConfig_GetRuntime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRuntime'
type MockConfig_GetRuntime_Call struct {
	*mock.Call
}

// GetRuntime is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetRuntime() *MockConfig_GetRuntime_Call {
	return &MockConfig_GetRuntime_Call{Call: _e.mock.On("GetRuntime")}
}

func (_c *MockConfig_GetRuntime_Call) Run(run func()) *MockConfig_GetRuntime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetRuntime_Call) Return(_a0 runtime.RuntimeConfig) *MockConfig_GetRuntime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetRuntime_Call) RunAndReturn(run func() runtime.RuntimeConfig) *MockConfig_GetRuntime_Call {
	_c.Call.Return(run)
	return _c
}

// GetSatori provides a mock function with no fields
func (_m *MockConfig) GetSatori() runtime.SatoriConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSatori")
	}

	var r0 runtime.SatoriConfig
	if rf, ok := ret.Get(0).(func() runtime.SatoriConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.SatoriConfig)
		}
	}

	return r0
}

// MockConfig_GetSatori_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSatori'
type MockConfig_GetSatori_Call struct {
	*mock.Call
}

// GetSatori is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetSatori() *MockConfig_GetSatori_Call {
	return &MockConfig_GetSatori_Call{Call: _e.mock.On("GetSatori")}
}

func (_c *MockConfig_GetSatori_Call) Run(run func()) *MockConfig_GetSatori_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetSatori_Call) Return(_a0 runtime.SatoriConfig) *MockConfig_GetSatori_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetSatori_Call) RunAndReturn(run func() runtime.SatoriConfig) *MockConfig_GetSatori_Call {
	_c.Call.Return(run)
	return _c
}

// GetSession provides a mock function with no fields
func (_m *MockConfig) GetSession() runtime.SessionConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSession")
	}

	var r0 runtime.SessionConfig
	if rf, ok := ret.Get(0).(func() runtime.SessionConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.SessionConfig)
		}
	}

	return r0
}

// MockConfig_GetSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSession'
type MockConfig_GetSession_Call struct {
	*mock.Call
}

// GetSession is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetSession() *MockConfig_GetSession_Call {
	return &MockConfig_GetSession_Call{Call: _e.mock.On("GetSession")}
}

func (_c *MockConfig_GetSession_Call) Run(run func()) *MockConfig_GetSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetSession_Call) Return(_a0 runtime.SessionConfig) *MockConfig_GetSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetSession_Call) RunAndReturn(run func() runtime.SessionConfig) *MockConfig_GetSession_Call {
	_c.Call.Return(run)
	return _c
}

// GetShutdownGraceSec provides a mock function with no fields
func (_m *MockConfig) GetShutdownGraceSec() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetShutdownGraceSec")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockConfig_GetShutdownGraceSec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShutdownGraceSec'
type MockConfig_GetShutdownGraceSec_Call struct {
	*mock.Call
}

// GetShutdownGraceSec is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetShutdownGraceSec() *MockConfig_GetShutdownGraceSec_Call {
	return &MockConfig_GetShutdownGraceSec_Call{Call: _e.mock.On("GetShutdownGraceSec")}
}

func (_c *MockConfig_GetShutdownGraceSec_Call) Run(run func()) *MockConfig_GetShutdownGraceSec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetShutdownGraceSec_Call) Return(_a0 int) *MockConfig_GetShutdownGraceSec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetShutdownGraceSec_Call) RunAndReturn(run func() int) *MockConfig_GetShutdownGraceSec_Call {
	_c.Call.Return(run)
	return _c
}

// GetSocial provides a mock function with no fields
func (_m *MockConfig) GetSocial() runtime.SocialConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSocial")
	}

	var r0 runtime.SocialConfig
	if rf, ok := ret.Get(0).(func() runtime.SocialConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.SocialConfig)
		}
	}

	return r0
}

// MockConfig_GetSocial_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSocial'
type MockConfig_GetSocial_Call struct {
	*mock.Call
}

// GetSocial is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetSocial() *MockConfig_GetSocial_Call {
	return &MockConfig_GetSocial_Call{Call: _e.mock.On("GetSocial")}
}

func (_c *MockConfig_GetSocial_Call) Run(run func()) *MockConfig_GetSocial_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetSocial_Call) Return(_a0 runtime.SocialConfig) *MockConfig_GetSocial_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetSocial_Call) RunAndReturn(run func() runtime.SocialConfig) *MockConfig_GetSocial_Call {
	_c.Call.Return(run)
	return _c
}

// GetSocket provides a mock function with no fields
func (_m *MockConfig) GetSocket() runtime.SocketConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSocket")
	}

	var r0 runtime.SocketConfig
	if rf, ok := ret.Get(0).(func() runtime.SocketConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.SocketConfig)
		}
	}

	return r0
}

// MockConfig_GetSocket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSocket'
type MockConfig_GetSocket_Call struct {
	*mock.Call
}

// GetSocket is a helper method to define mock.On call
func (_e *MockConfig_Expecter) GetSocket() *MockConfig_GetSocket_Call {
	return &MockConfig_GetSocket_Call{Call: _e.mock.On("GetSocket")}
}

func (_c *MockConfig_GetSocket_Call) Run(run func()) *MockConfig_GetSocket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfig_GetSocket_Call) Return(_a0 runtime.SocketConfig) *MockConfig_GetSocket_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConfig_GetSocket_Call) RunAndReturn(run func() runtime.SocketConfig) *MockConfig_GetSocket_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConfig creates a new instance of MockConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConfig {
	mock := &MockConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
