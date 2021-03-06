// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"

// ZoneAPI is an autogenerated mock type for the ZoneAPI type
type ZoneAPI struct {
	mock.Mock
}

// ListZones provides a mock function with given fields: project
func (_m *ZoneAPI) ListZones(project string) ([]string, error) {
	ret := _m.Called(project)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
