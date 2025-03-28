package service

import (
	"context"

	"github.com/rongpengju/gin-template/logic/types"
)

// 校验是否实现的所有的方法
var _ DemoInterface = (*demoService)(nil)

type DemoInterface interface {
	// Demo 示例方法
	Demo(req *types.DemoRequest) (*types.DemoResponse, error)
}

type demoService struct {
	ctx context.Context
}

func NewDemoService(ctx context.Context) DemoInterface {
	return &demoService{ctx: ctx}
}

func (d *demoService) Demo(req *types.DemoRequest) (*types.DemoResponse, error) {
	return &types.DemoResponse{}, nil
}
