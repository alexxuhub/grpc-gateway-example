package cmd

import (
	G "alex.github.com/go-gateway/global"
	"alex.github.com/go-gateway/intern/server"
	"github.com/urfave/cli/v2"
)

var (
	Server = &cli.Command{
		Name:        "server",
		Description: "run the server for http and grpc",
		Action:      runServer,
	}
)

func init() {
	G.InitResource()
}

func runServer(ctx *cli.Context) error {
	G.SugarLog.Info("now to run grpc gateway server")
	err := server.RunServer()
	if err != nil {
		G.SugarLog.Error(err.Error())
		return err
	}
	return nil
}
