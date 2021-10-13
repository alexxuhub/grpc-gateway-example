package route

import (
	"alex.github.com/go-gateway/intern/controller"
	"alex.github.com/go-gateway/pkg/alex/home"
	"google.golang.org/grpc"
)

type rpcHandler struct {
	server *grpc.Server
}

func Register(s *grpc.Server) {
	h := rpcHandler{server: s}
	h.addGroup()
}

func (r *rpcHandler) addGroup() {
	home.RegisterCropServiceServer(r.server, &controller.CorpHandler{})
}
