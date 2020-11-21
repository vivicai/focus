package service

import (
	"context"
	"focus/app/model"
	"github.com/gogf/gf/net/ghttp"
)

// 上下文管理服务
var Context = new(contextService)

type contextService struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextService) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextUserKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextUserKey)
	if value == nil {
		panic("找不到上下文对象，代码执行流程可能出现了问题，请检查下中间件注册顺序")
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *contextService) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
