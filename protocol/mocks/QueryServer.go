// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/bitoro-network/chain/protocol/x/prices/types"
)

// QueryServer is an autogenerated mock type for the QueryServer type
type QueryServer struct {
	mock.Mock
}

// AllMarketParams provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) AllMarketParams(_a0 context.Context, _a1 *types.QueryAllMarketParamsRequest) (*types.QueryAllMarketParamsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AllMarketParams")
	}

	var r0 *types.QueryAllMarketParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMarketParamsRequest) (*types.QueryAllMarketParamsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMarketParamsRequest) *types.QueryAllMarketParamsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMarketParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMarketParamsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllMarketPrices provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) AllMarketPrices(_a0 context.Context, _a1 *types.QueryAllMarketPricesRequest) (*types.QueryAllMarketPricesResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AllMarketPrices")
	}

	var r0 *types.QueryAllMarketPricesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMarketPricesRequest) (*types.QueryAllMarketPricesResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMarketPricesRequest) *types.QueryAllMarketPricesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMarketPricesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMarketPricesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarketParam provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) MarketParam(_a0 context.Context, _a1 *types.QueryMarketParamRequest) (*types.QueryMarketParamResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for MarketParam")
	}

	var r0 *types.QueryMarketParamResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMarketParamRequest) (*types.QueryMarketParamResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMarketParamRequest) *types.QueryMarketParamResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryMarketParamResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryMarketParamRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarketPrice provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) MarketPrice(_a0 context.Context, _a1 *types.QueryMarketPriceRequest) (*types.QueryMarketPriceResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for MarketPrice")
	}

	var r0 *types.QueryMarketPriceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMarketPriceRequest) (*types.QueryMarketPriceResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMarketPriceRequest) *types.QueryMarketPriceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryMarketPriceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryMarketPriceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NextMarketId provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) NextMarketId(_a0 context.Context, _a1 *types.QueryNextMarketIdRequest) (*types.QueryNextMarketIdResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for NextMarketId")
	}

	var r0 *types.QueryNextMarketIdResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryNextMarketIdRequest) (*types.QueryNextMarketIdResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryNextMarketIdRequest) *types.QueryNextMarketIdResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryNextMarketIdResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryNextMarketIdRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryServer creates a new instance of QueryServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryServer {
	mock := &QueryServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
