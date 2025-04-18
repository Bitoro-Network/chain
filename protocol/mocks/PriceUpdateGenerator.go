// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	pricestypes "github.com/bitoro-network/chain/protocol/x/prices/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PriceUpdateGenerator is an autogenerated mock type for the PriceUpdateGenerator type
type PriceUpdateGenerator struct {
	mock.Mock
}

// GetValidMarketPriceUpdates provides a mock function with given fields: ctx, extCommitBz
func (_m *PriceUpdateGenerator) GetValidMarketPriceUpdates(ctx types.Context, extCommitBz []byte) (*pricestypes.MsgUpdateMarketPrices, error) {
	ret := _m.Called(ctx, extCommitBz)

	if len(ret) == 0 {
		panic("no return value specified for GetValidMarketPriceUpdates")
	}

	var r0 *pricestypes.MsgUpdateMarketPrices
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, []byte) (*pricestypes.MsgUpdateMarketPrices, error)); ok {
		return rf(ctx, extCommitBz)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []byte) *pricestypes.MsgUpdateMarketPrices); ok {
		r0 = rf(ctx, extCommitBz)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pricestypes.MsgUpdateMarketPrices)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, []byte) error); ok {
		r1 = rf(ctx, extCommitBz)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPriceUpdateGenerator creates a new instance of PriceUpdateGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceUpdateGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceUpdateGenerator {
	mock := &PriceUpdateGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
