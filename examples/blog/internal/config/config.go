package config

import (
	"encoding/json"

	"github.com/go-ashe/gkit/app"
	"github.com/go-ashe/gkit/log"
	"github.com/go-ashe/gkit/monitor"

	"github.com/go-ashe/gkit/config"
	"github.com/go-ashe/gkit/config/file"

	_ "github.com/go-ashe/gkit/x/codec/yaml"
)

type Config struct {
	App   *App              `json:"app" yaml:"app"`
	Redis map[string]*Redis `json:"redis" yaml:"redis"`
}

var conf Config

func InitConfig(path string) error {
	log.Info("init config...")
	conf = Config{
		App:   &App{},
		Redis: make(map[string]*Redis, 0),
	}
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	if err := c.Load(); err != nil {
		return err
	}

	b, _ := c.Source()
	log.Infof("raw config data:%s", string(b))

	// Unmarshal the config to struct
	if err := c.Scan(&conf); err != nil {
		return err
	}

	b, _ = json.Marshal(conf)
	log.Infof("marshaled config data:%s", string(b))

	return nil
}

func GetHttpConfig() *app.HTTPServer {
	if conf.App.HTTP == nil {
		return &app.HTTPServer{
			Port: 8080,
		}
	}
	return conf.App.HTTP
}

func GetGRPCConfig() *app.GRPCServer {
	if conf.App.GRPC == nil {
		return &app.GRPCServer{
			Port: 9090,
		}
	}
	return conf.App.GRPC
}

func GetMonitorConfig() *monitor.Config {
	return conf.App.Monitor
}

func GetRedisConfig() map[string]*Redis {
	return conf.Redis
}

func GetLogLevel() log.Level {
	switch conf.App.LogLevel {
	case "debug":
		return log.LevelDebug
	case "release":
		return log.LevelInfo
	}
	return log.LevelDebug
}

type RunMode string

const (
	RunModeLocal RunMode = "local"
	RunModeTest  RunMode = "test"
	RunModeProd  RunMode = "prod"
)

func IsProd() bool {
	return conf.App.RunMode == RunModeProd
}
