// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package article

import (
	"context"

	"zerorequest/internal/svc"
	"zerorequest/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByIdLogic) GetArticleById(req *types.ArticleRequest) (resp *types.CommonResponse, err error) {
	article, _ := l.svcCtx.ArticleModel.FindOne(l.ctx, int64(req.Id))
	return &types.CommonResponse{
		Code:    200,
		Message: "success",
		Data: types.Article{
			Id:          int(article.Id),
			Title:       article.Title.String,
			Content:     article.Content.String,
			Description: article.Description.String,
			Picture:     article.Picture.String,
		},
		Success: true,
	}, nil
}
