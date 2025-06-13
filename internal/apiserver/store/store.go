package store

import (
	"context"
	"sync"

	"github.com/onexstack/onexstack/pkg/store/where"
	"gorm.io/gorm"
)

var (
	once sync.Once

	S *datastore
)

type IStore interface {
	DB(ctx context.Context, wheres ...where.Where) *gorm.DB
	TX(ctx context.Context, fn func(ctx context.Context) error) error

	User()
	Post()
}

// transactionKey 用于在 context.Context 中存储事务上下文的键.
type transactionKey struct{}

type datastore struct {
	core *gorm.DB
}

// var _ IStore = (*datastore)(nil)

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{db}
	})

	return S
}

func (store *datastore) DB(ctx context.Context, wheres ...where.Where) *gorm.DB {
	db := store.core
	// 从上下文提取事务实例
	if tx, ok := ctx.Value(transactionKey{}).(*gorm.DB); ok {
		db = tx
	}

	// 遍历所有传入的条件并逐一叠加到数据库查询对象上
	for _, whr := range wheres {
		db = whr.Where(db)
	}

	return db
}

func (store *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return store.core.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}

// func (store *datastore) User() UserStore {
// 	return newUserStore(store)
// }
