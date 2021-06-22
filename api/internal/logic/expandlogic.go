package logic

import (
	"context"
	"github.com/wanmei002/shorturl/rpc/transformer"
	
	"github.com/wanmei002/shorturl/api/internal/svc"
	"github.com/wanmei002/shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExpandLogic {
	return ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req types.ExpandReq) (types.ExpandResp, error) {
	
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandReq{
		Shorten: req.Shorten,
	})
	
	if err != nil {
		return types.ExpandResp{}, err
	}
	
	

	return types.ExpandResp{
		Url: resp.Url,
	}, nil
}
