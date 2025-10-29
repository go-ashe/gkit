package grpc

import (
	"fmt"

	"github.com/go-ashe/gkit/app"
	"github.com/go-ashe/gkit/examples/blog/api"
	tg "github.com/go-ashe/gkit/transport/grpc"
	"google.golang.org/grpc"
)

func NewServer(c *app.GRPCServer, s api.BlogServer) *tg.Server {
	opts := []tg.ServerOption{
		tg.UnaryInterceptor([]grpc.UnaryServerInterceptor{
			//grpc_prometheus.UnaryServerInterceptor,
			//requestid.GRPCServerMiddleware(),
		}...),
	}
	if c.Port > 0 {
		opts = append(opts, tg.Address(fmt.Sprintf(":%d", c.Port)))
	}

	srv := tg.NewServer(opts...)
	api.RegisterBlogServer(srv, s)
	return srv
}
