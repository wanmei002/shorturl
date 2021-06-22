package logic

import (
	"context"

	"github.com/wanmei002/shorturl/rpc/internal/svc"
	"github.com/wanmei002/shorturl/rpc/transform"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	ret := &transform.ExpandResp{}
	
	switch in.Shorten {
	case "aaaaa":
		ret.Url = "www.zzh.com"
	default:
		ret.Url = "www.default.com"
	}

	return ret, nil
}
