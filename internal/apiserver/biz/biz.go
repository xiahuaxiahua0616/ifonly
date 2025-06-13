package biz

import "github.com/xiahuaxiahua0616/ifonly/internal/apiserver/store"

type IBiz interface {
	// 获取用户业务
	UserV1()
	PostV1()
}

// biz 是 IBiz 的一个具体实现.
type biz struct {
	store store.IStore
	// authz *auth.Authz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(store store.IStore) *biz {
	return &biz{store: store}
}

// UserV1 返回一个实现了 UserBiz 接口的实例.
func (b *biz) UserV1() {

}

// PostV1 返回一个实现了 PostBiz 接口的实例.
func (b *biz) PostV1() {

}
