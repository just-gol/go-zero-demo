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

type AddFocusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFocusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFocusLogic {
	return &AddFocusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFocusLogic) AddFocus(req *types.AddFocusRequest) (resp *types.CommonResponse, err error) {
	focus := mysql.Focus{
		Title:    sql.NullString{String: req.Title, Valid: true},
		Pic:      sql.NullString{String: req.Pic, Valid: true},
		Link:     sql.NullString{String: req.Link, Valid: true},
		Position: int64(req.Position),
	}
	_, err = l.svcCtx.FocusModel.Insert(l.ctx, &focus)
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
		Data:    "",
		Success: true,
	}, nil
}
