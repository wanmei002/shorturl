package logic

import (
	"context"

	"github.com/wanmei002/shorturl/rpc/internal/svc"
	"github.com/wanmei002/shorturl/rpc/transform"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	
	ret := &transform.ShortenResp{}
	switch in.Url {
	case "www.zzh.com":
		ret.Shorten = "aaaaa"
	default:
		ret.Shorten = "bbbbb"
	}
	
	return ret, nil
}
