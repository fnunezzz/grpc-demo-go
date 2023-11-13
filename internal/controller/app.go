package controller

import (
	"context"
	"time"

	"github.com/fnunezzz/grpc-demo-go/internal/middleware"
	proto "github.com/fnunezzz/grpc-demo-go/internal/proto/gen"
	"github.com/fnunezzz/grpc-demo-go/internal/service"
	"google.golang.org/grpc"
)

type appServerContract struct {
	healthService service.AppService
	proto.UnimplementedAppServer
}



func NewAppController(server *grpc.Server, healthService service.AppService) {
	contract := &appServerContract{healthService: healthService}
	proto.RegisterAppServer(server, contract)
}

func (s appServerContract) HealthCheck(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	err := middleware.AuthMiddleware(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.healthService.HealthCheck()
	
	if err != nil {
		return &proto.HealthCheckResponse{Time: time.Now().Local().String(), Status: false}, nil
	}

	return &proto.HealthCheckResponse{Time: time.Now().Local().String(), Status: true}, nil
}
