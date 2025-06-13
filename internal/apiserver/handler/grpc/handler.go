package grpc

import (
	"github.com/xiahuaxiahua0616/ifonly/internal/apiserver/biz"
	apiv1 "github.com/xiahuaxiahua0616/ifonly/pkg/api/apiserver/v1"
)

type Handler struct {
	apiv1.UnimplementedIfOnlyServer

	biz biz.IBiz
}

func NewHandler(biz biz.IBiz) *Handler {
	return &Handler{
		biz: biz,
	}
}
