// Code generated by mockery v1.0.0. DO NOT EDIT.

package render

import mock "github.com/stretchr/testify/mock"
import model "github.com/slok/meterm/internal/model"
import render "github.com/slok/meterm/internal/view/render"

// GraphWidget is an autogenerated mock type for the GraphWidget type
type GraphWidget struct {
	mock.Mock
}

// GetGraphPointQuantity provides a mock function with given fields:
func (_m *GraphWidget) GetGraphPointQuantity() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetWidgetCfg provides a mock function with given fields:
func (_m *GraphWidget) GetWidgetCfg() model.Widget {
	ret := _m.Called()

	var r0 model.Widget
	if rf, ok := ret.Get(0).(func() model.Widget); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.Widget)
	}

	return r0
}

// Sync provides a mock function with given fields: series
func (_m *GraphWidget) Sync(series []render.Series) error {
	ret := _m.Called(series)

	var r0 error
	if rf, ok := ret.Get(0).(func([]render.Series) error); ok {
		r0 = rf(series)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}