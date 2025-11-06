// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package focus

import (
	"context"
	"database/sql"

	"zerorequest/internal/svc"
	"zerorequest/internal/types"
	"zerorequest/model/mysql"

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
	focus := mysql.Focus{
		Id:       int64(req.Id),
		Title:    sql.NullString{String: req.Title, Valid: true},
		Pic:      sql.NullString{String: req.Pic, Valid: true},
		Link:     sql.NullString{String: req.Link, Valid: true},
		Position: int64(req.Position),
	}
	err = l.svcCtx.FocusModel.Update(l.ctx, &focus)
	if err != nil {
		return &types.CommonResponse{
			Code:    500,
			Message: "error",
			Data:    "",
			Success: false,
		}, nil

	}
	return &types.CommonResponse{
		Code:    200,
		Message: "success",
		Data:    "",
		Success: true,
	}, nil
}
