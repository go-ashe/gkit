package config

import (
	"github.com/go-ashe/gkit/app"
	"github.com/go-ashe/gkit/monitor"
)

type App struct {
	// GRPC 服务端配置项
	GRPC *app.GRPCServer `json:"grpc" yaml:"grpc"`
	// HTTP 服务端配置项
	HTTP *app.HTTPServer `json:"http" yaml:"http"`
	// HTTP 服务端配置项
	Monitor *monitor.Config `json:"monitor" yaml:"monitor"`

	LogLevel string  `json:"log_level" yaml:"log_level"`
	RunMode  RunMode `json:"run_mode" yaml:"run_mode"`
}
