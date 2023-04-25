package template

// HandlerFNC is the handler template used for new function projects.
var HandlerFNC = `package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "{{.Vendor}}{{.Service}}/proto"
)

type {{title .Service}} struct{}

func (e *{{title .Service}}) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received {{title .Service}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
`

// HandlerSRV is the handler template used for new service projects.
var HandlerSRV = `package handler

import (
	"context"
	"go-micro.dev/v4/logger"

	pb "{{.Vendor}}{{.Service}}/proto"
	"github.com/kioku-project/kioku/store"
)

type {{title .Service}} struct{store store.Store}

func New(s store.Store) *{{title .Service}} { return &{{title .Service}}{store: s} }

func (e *{{title .Service}}) Call(ctx context.Context, req *pb.{{title .Service}}Request, rsp *pb.{{title .Service}}Response) error {
	logger.Infof("Received {{title .Service}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
`

var HealthSRV = `package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "{{.Vendor}}{{.Service}}/proto"
)

type Health struct{}

func (h *Health) Check(ctx context.Context, req *pb.HealthCheckRequest, rsp *pb.HealthCheckResponse) error {
	rsp.Status = pb.HealthCheckResponse_SERVING
	return nil
}

func (h *Health) Watch(ctx context.Context, req *pb.HealthCheckRequest, stream pb.Health_WatchStream) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
`
