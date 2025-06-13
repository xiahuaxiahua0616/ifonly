package biz

type IBiz interface {
	// 获取用户业务
	UserV1()
	PostV1()
}

// type biz struct {
// 	// store store.IStore
// 	// authz *a
// 	store store.IStore
// 	authz *auth.Authz
// }

// // 确保 biz 实现了 IBiz 接口.
// var _ IBiz = (*biz)(nil)

// // NewBiz 创建一个 IBiz 类型的实例.
// func NewBiz(store store.IStore, authz *auth.Authz) *biz {
// 	return &biz{store: store, authz: authz}
// }

// // UserV1 返回一个实现了 UserBiz 接口的实例.
// func (b *biz) UserV1() userv1.UserBiz {
// 	return userv1.New(b.store, b.authz)
// }

// // PostV1 返回一个实现了 PostBiz 接口的实例.
// func (b *biz) PostV1() postv1.PostBiz {
// 	return postv1.New(b.store)
// }
