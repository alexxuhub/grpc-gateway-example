package server

import (
	"alex.github.com/go-gateway/intern/route"
	"alex.github.com/go-gateway/pkg/alex/home"
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	grpcPort    = ":10001"
	grpcAddress = "localhost"
)

func RunServer() error {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("listen tcp address error")
	}
	s := grpc.NewServer()
	route.Register(s)
	go func() {
		err := s.Serve(lis)
		if err != nil {
			glog.Error(" start grpc server 失败")
		}
	}()
	defer glog.Flush()
	if err = runGateWayServer(); err != nil {
		return err
	}
	return nil
}

func runGateWayServer() error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := home.RegisterCropServiceHandlerFromEndpoint(ctx, mux, grpcAddress+grpcPort, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)

}
