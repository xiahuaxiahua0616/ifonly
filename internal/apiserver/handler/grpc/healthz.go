package grpc

import (
	"context"
	"time"

	"github.com/xiahuaxiahua0616/ifonly/internal/pkg/log"
	apiv1 "github.com/xiahuaxiahua0616/ifonly/pkg/api/apiserver/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) Healthz(ctx context.Context, rq *emptypb.Empty) (*apiv1.HealthzResponse, error) {
	log.W(ctx).Infow("Healthz handler is called", "method", "Healthz", "status", "healthy")
	return &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}
