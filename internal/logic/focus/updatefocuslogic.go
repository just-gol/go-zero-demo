// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"zerorequest/internal/svc"
	"zerorequest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFocusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFocusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFocusLogic {
	return &UpdateFocusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFocusLogic) UpdateFocus(req *types.UpdateFocusRequest) (resp *types.CommonResponse, err error) {
	logx.Info("UpdateFocus", req.Id, req.Image, req.Link, req.Title)
	return &types.CommonResponse{
		Code:    200,
		Message: "success",
		Data:    nil,
	}, nil
}
