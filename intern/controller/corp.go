package controller

import (
	G "alex.github.com/go-gateway/global"
	"alex.github.com/go-gateway/pkg/alex/home"
	"context"
	"time"
)

type CorpHandler struct {
	home.UnimplementedCropServiceServer
}

func (c *CorpHandler) GetCorpTicket(ctx context.Context, in *home.GetCorpTicketRequest) (*home.GetCorpTicketResponse, error) {
	url,name := in.GetUrl(),in.GetName()
	G.SugarLog.Info("get url from request :",url,",name:",name)

	return &home.GetCorpTicketResponse{
		Timestamp: int32(time.Now().Unix()),
		Noncestr:  "dsajkdfas",
		CorpId:    "10012",
		Sign:      "sdklas8217820asd0asd",
		Agentid:   800,
	},nil
}

func (c *CorpHandler) PostCorpTicket(ctx context.Context, in *home.PostCorpTicketRequest) (*home.PostCorpTicketResponse, error) {
	name,age :=  in.GetName(),in.GetAge()
	G.SugarLog.Info("get body from request , name:",name,",age:",age)
	return &home.PostCorpTicketResponse{
		Code:    0,
		Message: "success",
		Data:    nil,
	},nil
}
