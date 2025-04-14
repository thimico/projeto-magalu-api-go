package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"projeto-magalu-api-go/pkg/api/model"
)

type ProductService struct {
	mock.Mock
}

func (_m *ProductService) GetByID(ctx context.Context, id string) (*model.ProductOut, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.ProductOut
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.ProductOut); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProductOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *ProductService) Create(ctx context.Context, in *model.ProductIn) (*model.ProductOut, error) {
	ret := _m.Called(ctx, in)
	var r0 *model.ProductOut
	if rf, ok := ret.Get(0).(func(context.Context, *model.ProductIn) *model.ProductOut); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Get(0).(*model.ProductOut)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.ProductIn) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}
