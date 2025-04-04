// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	perpetualstypes "github.com/bitoro-network/chain/protocol/x/perpetuals/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PreparePerpetualsKeeper is an autogenerated mock type for the PreparePerpetualsKeeper type
type PreparePerpetualsKeeper struct {
	mock.Mock
}

// GetAddPremiumVotes provides a mock function with given fields: ctx
func (_m *PreparePerpetualsKeeper) GetAddPremiumVotes(ctx types.Context) *perpetualstypes.MsgAddPremiumVotes {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAddPremiumVotes")
	}

	var r0 *perpetualstypes.MsgAddPremiumVotes
	if rf, ok := ret.Get(0).(func(types.Context) *perpetualstypes.MsgAddPremiumVotes); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*perpetualstypes.MsgAddPremiumVotes)
		}
	}

	return r0
}

// NewPreparePerpetualsKeeper creates a new instance of PreparePerpetualsKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPreparePerpetualsKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *PreparePerpetualsKeeper {
	mock := &PreparePerpetualsKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
