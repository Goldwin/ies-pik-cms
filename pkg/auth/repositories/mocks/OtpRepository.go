// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	entities "github.com/Goldwin/ies-pik-cms/pkg/auth/entities"
	mock "github.com/stretchr/testify/mock"
)

// OtpRepository is an autogenerated mock type for the OtpRepository type
type OtpRepository struct {
	mock.Mock
}

type OtpRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *OtpRepository) EXPECT() *OtpRepository_Expecter {
	return &OtpRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: _a0
func (_m *OtpRepository) Delete(_a0 *entities.Otp) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Otp) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OtpRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type OtpRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 *entities.Otp
func (_e *OtpRepository_Expecter) Delete(_a0 interface{}) *OtpRepository_Delete_Call {
	return &OtpRepository_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *OtpRepository_Delete_Call) Run(run func(_a0 *entities.Otp)) *OtpRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.Otp))
	})
	return _c
}

func (_c *OtpRepository_Delete_Call) Return(_a0 error) *OtpRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OtpRepository_Delete_Call) RunAndReturn(run func(*entities.Otp) error) *OtpRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0
func (_m *OtpRepository) Get(_a0 string) (*entities.Otp, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *entities.Otp
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Otp, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Otp); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Otp)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OtpRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type OtpRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 string
func (_e *OtpRepository_Expecter) Get(_a0 interface{}) *OtpRepository_Get_Call {
	return &OtpRepository_Get_Call{Call: _e.mock.On("Get", _a0)}
}

func (_c *OtpRepository_Get_Call) Run(run func(_a0 string)) *OtpRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OtpRepository_Get_Call) Return(_a0 *entities.Otp, _a1 error) *OtpRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OtpRepository_Get_Call) RunAndReturn(run func(string) (*entities.Otp, error)) *OtpRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetOtp provides a mock function with given fields: _a0
func (_m *OtpRepository) GetOtp(_a0 entities.EmailAddress) (*entities.Otp, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetOtp")
	}

	var r0 *entities.Otp
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.EmailAddress) (*entities.Otp, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(entities.EmailAddress) *entities.Otp); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Otp)
		}
	}

	if rf, ok := ret.Get(1).(func(entities.EmailAddress) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OtpRepository_GetOtp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOtp'
type OtpRepository_GetOtp_Call struct {
	*mock.Call
}

// GetOtp is a helper method to define mock.On call
//   - _a0 entities.EmailAddress
func (_e *OtpRepository_Expecter) GetOtp(_a0 interface{}) *OtpRepository_GetOtp_Call {
	return &OtpRepository_GetOtp_Call{Call: _e.mock.On("GetOtp", _a0)}
}

func (_c *OtpRepository_GetOtp_Call) Run(run func(_a0 entities.EmailAddress)) *OtpRepository_GetOtp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.EmailAddress))
	})
	return _c
}

func (_c *OtpRepository_GetOtp_Call) Return(_a0 *entities.Otp, _a1 error) *OtpRepository_GetOtp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OtpRepository_GetOtp_Call) RunAndReturn(run func(entities.EmailAddress) (*entities.Otp, error)) *OtpRepository_GetOtp_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: _a0
func (_m *OtpRepository) List(_a0 []string) ([]*entities.Otp, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*entities.Otp
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*entities.Otp, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]string) []*entities.Otp); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Otp)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OtpRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type OtpRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 []string
func (_e *OtpRepository_Expecter) List(_a0 interface{}) *OtpRepository_List_Call {
	return &OtpRepository_List_Call{Call: _e.mock.On("List", _a0)}
}

func (_c *OtpRepository_List_Call) Run(run func(_a0 []string)) *OtpRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *OtpRepository_List_Call) Return(_a0 []*entities.Otp, _a1 error) *OtpRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OtpRepository_List_Call) RunAndReturn(run func([]string) ([]*entities.Otp, error)) *OtpRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: _a0
func (_m *OtpRepository) Save(_a0 *entities.Otp) (*entities.Otp, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *entities.Otp
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.Otp) (*entities.Otp, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*entities.Otp) *entities.Otp); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Otp)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.Otp) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OtpRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type OtpRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - _a0 *entities.Otp
func (_e *OtpRepository_Expecter) Save(_a0 interface{}) *OtpRepository_Save_Call {
	return &OtpRepository_Save_Call{Call: _e.mock.On("Save", _a0)}
}

func (_c *OtpRepository_Save_Call) Run(run func(_a0 *entities.Otp)) *OtpRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.Otp))
	})
	return _c
}

func (_c *OtpRepository_Save_Call) Return(_a0 *entities.Otp, _a1 error) *OtpRepository_Save_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OtpRepository_Save_Call) RunAndReturn(run func(*entities.Otp) (*entities.Otp, error)) *OtpRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewOtpRepository creates a new instance of OtpRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOtpRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OtpRepository {
	mock := &OtpRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
