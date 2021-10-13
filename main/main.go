package main

import (
	"alex.github.com/go-gateway/intern/cmd"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func main() {
	app := &cli.App{
		Name:     "grpc gateway system ",
		Usage:    "简单系统支持 http协议和grpc 协议",
		Compiled: time.Now(),
		Commands: []*cli.Command{
			cmd.Server,
		},
	}

	_ = app.Run(os.Args)
}
