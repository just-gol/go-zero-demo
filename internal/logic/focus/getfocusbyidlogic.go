// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"

	"zerorequest/internal/svc"
	"zerorequest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFocusByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFocusByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFocusByIdLogic {
	return &GetFocusByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFocusByIdLogic) GetFocusById(req *types.FocusRequest) (resp *types.CommonResponse, err error) {
	logx.Info("GetFocusById", req.Id)
	id := req.Id
	focus, err := l.svcCtx.FocusModel.FindOne(l.ctx, int64(id))
	if err != nil {
		return &types.CommonResponse{
			Code:    500,
			Message: "error",
			Data:    "",
			Success: false,
		}, err
	}
	return &types.CommonResponse{
		Code:    200,
		Message: "success",
		Data: types.Focus{
			Id:    int(focus.Id),
			Pic:   focus.Pic.String,
			Link:  focus.Link.String,
			Title: focus.Title.String,
		},
		Success: true,
	}, nil
}
