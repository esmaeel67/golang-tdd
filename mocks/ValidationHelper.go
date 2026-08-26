package mocks

import "github.com/stretchr/testify/mock"

type ValidationHelper struct {
	mock.Mock
}

// CheckInput provides a mock function with given fields: operator, operands
func (_m *ValidationHelper) CheckInput(operator string, operands []float64) error {
	ret := _m.Called(operator, operands)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []float64) error); ok {
		r0 = rf(operator, operands)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

type mockConstructorTestingTNewValidationHelper interface {
	mock.TestingT
	Cleanup(func())
}

func NewValidationHelper(t mockConstructorTestingTNewValidationHelper) *ValidationHelper {
	mock := &ValidationHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() {
		mock.AssertExpectations(t)
	})

	return mock
}
