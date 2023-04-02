package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "greeter/proto"
)

type Greeter struct{}

func (e *Greeter) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Greeter.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
