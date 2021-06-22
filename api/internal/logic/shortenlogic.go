package logic

import (
	"context"
	"github.com/wanmei002/shorturl/rpc/transformer"
	
	"github.com/wanmei002/shorturl/api/internal/svc"
	"github.com/wanmei002/shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortenLogic {
	return ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenReq) (types.ShortenResp, error) {
	
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{
		Url: req.Url,
	})
	
	if err != nil {
		return types.ShortenResp{}, err
	}
	
	return types.ShortenResp{
		Shorten: resp.Shorten,
	}, nil
}
