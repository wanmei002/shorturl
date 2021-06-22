package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/wanmei002/shorturl/api/internal/config"
	"github.com/wanmei002/shorturl/rpc/transformer"
)

type ServiceContext struct {
	Config config.Config
	Transformer transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
