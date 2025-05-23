// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockSessionConfig is an autogenerated mock type for the SessionConfig type
type MockSessionConfig struct {
	mock.Mock
}

type MockSessionConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSessionConfig) EXPECT() *MockSessionConfig_Expecter {
	return &MockSessionConfig_Expecter{mock: &_m.Mock}
}

// GetEncryptionKey provides a mock function with no fields
func (_m *MockSessionConfig) GetEncryptionKey() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEncryptionKey")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSessionConfig_GetEncryptionKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEncryptionKey'
type MockSessionConfig_GetEncryptionKey_Call struct {
	*mock.Call
}

// GetEncryptionKey is a helper method to define mock.On call
func (_e *MockSessionConfig_Expecter) GetEncryptionKey() *MockSessionConfig_GetEncryptionKey_Call {
	return &MockSessionConfig_GetEncryptionKey_Call{Call: _e.mock.On("GetEncryptionKey")}
}

func (_c *MockSessionConfig_GetEncryptionKey_Call) Run(run func()) *MockSessionConfig_GetEncryptionKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionConfig_GetEncryptionKey_Call) Return(_a0 string) *MockSessionConfig_GetEncryptionKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionConfig_GetEncryptionKey_Call) RunAndReturn(run func() string) *MockSessionConfig_GetEncryptionKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetRefreshEncryptionKey provides a mock function with no fields
func (_m *MockSessionConfig) GetRefreshEncryptionKey() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRefreshEncryptionKey")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSessionConfig_GetRefreshEncryptionKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRefreshEncryptionKey'
type MockSessionConfig_GetRefreshEncryptionKey_Call struct {
	*mock.Call
}

// GetRefreshEncryptionKey is a helper method to define mock.On call
func (_e *MockSessionConfig_Expecter) GetRefreshEncryptionKey() *MockSessionConfig_GetRefreshEncryptionKey_Call {
	return &MockSessionConfig_GetRefreshEncryptionKey_Call{Call: _e.mock.On("GetRefreshEncryptionKey")}
}

func (_c *MockSessionConfig_GetRefreshEncryptionKey_Call) Run(run func()) *MockSessionConfig_GetRefreshEncryptionKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionConfig_GetRefreshEncryptionKey_Call) Return(_a0 string) *MockSessionConfig_GetRefreshEncryptionKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionConfig_GetRefreshEncryptionKey_Call) RunAndReturn(run func() string) *MockSessionConfig_GetRefreshEncryptionKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetRefreshTokenExpirySec provides a mock function with no fields
func (_m *MockSessionConfig) GetRefreshTokenExpirySec() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRefreshTokenExpirySec")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockSessionConfig_GetRefreshTokenExpirySec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRefreshTokenExpirySec'
type MockSessionConfig_GetRefreshTokenExpirySec_Call struct {
	*mock.Call
}

// GetRefreshTokenExpirySec is a helper method to define mock.On call
func (_e *MockSessionConfig_Expecter) GetRefreshTokenExpirySec() *MockSessionConfig_GetRefreshTokenExpirySec_Call {
	return &MockSessionConfig_GetRefreshTokenExpirySec_Call{Call: _e.mock.On("GetRefreshTokenExpirySec")}
}

func (_c *MockSessionConfig_GetRefreshTokenExpirySec_Call) Run(run func()) *MockSessionConfig_GetRefreshTokenExpirySec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionConfig_GetRefreshTokenExpirySec_Call) Return(_a0 int64) *MockSessionConfig_GetRefreshTokenExpirySec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionConfig_GetRefreshTokenExpirySec_Call) RunAndReturn(run func() int64) *MockSessionConfig_GetRefreshTokenExpirySec_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingleMatch provides a mock function with no fields
func (_m *MockSessionConfig) GetSingleMatch() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSingleMatch")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSessionConfig_GetSingleMatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingleMatch'
type MockSessionConfig_GetSingleMatch_Call struct {
	*mock.Call
}

// GetSingleMatch is a helper method to define mock.On call
func (_e *MockSessionConfig_Expecter) GetSingleMatch() *MockSessionConfig_GetSingleMatch_Call {
	return &MockSessionConfig_GetSingleMatch_Call{Call: _e.mock.On("GetSingleMatch")}
}

func (_c *MockSessionConfig_GetSingleMatch_Call) Run(run func()) *MockSessionConfig_GetSingleMatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionConfig_GetSingleMatch_Call) Return(_a0 bool) *MockSessionConfig_GetSingleMatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionConfig_GetSingleMatch_Call) RunAndReturn(run func() bool) *MockSessionConfig_GetSingleMatch_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingleParty provides a mock function with no fields
func (_m *MockSessionConfig) GetSingleParty() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSingleParty")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSessionConfig_GetSingleParty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingleParty'
type MockSessionConfig_GetSingleParty_Call struct {
	*mock.Call
}

// GetSingleParty is a helper method to define mock.On call
func (_e *MockSessionConfig_Expecter) GetSingleParty() *MockSessionConfig_GetSingleParty_Call {
	return &MockSessionConfig_GetSingleParty_Call{Call: _e.mock.On("GetSingleParty")}
}

func (_c *MockSessionConfig_GetSingleParty_Call) Run(run func()) *MockSessionConfig_GetSingleParty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionConfig_GetSingleParty_Call) Return(_a0 bool) *MockSessionConfig_GetSingleParty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionConfig_GetSingleParty_Call) RunAndReturn(run func() bool) *MockSessionConfig_GetSingleParty_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingleSession provides a mock function with no fields
func (_m *MockSessionConfig) GetSingleSession() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSingleSession")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSessionConfig_GetSingleSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingleSession'
type MockSessionConfig_GetSingleSession_Call struct {
	*mock.Call
}

// GetSingleSession is a helper method to define mock.On call
func (_e *MockSessionConfig_Expecter) GetSingleSession() *MockSessionConfig_GetSingleSession_Call {
	return &MockSessionConfig_GetSingleSession_Call{Call: _e.mock.On("GetSingleSession")}
}

func (_c *MockSessionConfig_GetSingleSession_Call) Run(run func()) *MockSessionConfig_GetSingleSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionConfig_GetSingleSession_Call) Return(_a0 bool) *MockSessionConfig_GetSingleSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionConfig_GetSingleSession_Call) RunAndReturn(run func() bool) *MockSessionConfig_GetSingleSession_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingleSocket provides a mock function with no fields
func (_m *MockSessionConfig) GetSingleSocket() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSingleSocket")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSessionConfig_GetSingleSocket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingleSocket'
type MockSessionConfig_GetSingleSocket_Call struct {
	*mock.Call
}

// GetSingleSocket is a helper method to define mock.On call
func (_e *MockSessionConfig_Expecter) GetSingleSocket() *MockSessionConfig_GetSingleSocket_Call {
	return &MockSessionConfig_GetSingleSocket_Call{Call: _e.mock.On("GetSingleSocket")}
}

func (_c *MockSessionConfig_GetSingleSocket_Call) Run(run func()) *MockSessionConfig_GetSingleSocket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionConfig_GetSingleSocket_Call) Return(_a0 bool) *MockSessionConfig_GetSingleSocket_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionConfig_GetSingleSocket_Call) RunAndReturn(run func() bool) *MockSessionConfig_GetSingleSocket_Call {
	_c.Call.Return(run)
	return _c
}

// GetTokenExpirySec provides a mock function with no fields
func (_m *MockSessionConfig) GetTokenExpirySec() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTokenExpirySec")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockSessionConfig_GetTokenExpirySec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenExpirySec'
type MockSessionConfig_GetTokenExpirySec_Call struct {
	*mock.Call
}

// GetTokenExpirySec is a helper method to define mock.On call
func (_e *MockSessionConfig_Expecter) GetTokenExpirySec() *MockSessionConfig_GetTokenExpirySec_Call {
	return &MockSessionConfig_GetTokenExpirySec_Call{Call: _e.mock.On("GetTokenExpirySec")}
}

func (_c *MockSessionConfig_GetTokenExpirySec_Call) Run(run func()) *MockSessionConfig_GetTokenExpirySec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSessionConfig_GetTokenExpirySec_Call) Return(_a0 int64) *MockSessionConfig_GetTokenExpirySec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSessionConfig_GetTokenExpirySec_Call) RunAndReturn(run func() int64) *MockSessionConfig_GetTokenExpirySec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSessionConfig creates a new instance of MockSessionConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSessionConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSessionConfig {
	mock := &MockSessionConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
