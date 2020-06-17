package models

import (
	"context"
	"github.com/go-xorm/xorm"
)

var (
	db   *xorm.Engine
	beans = []interface{}{
		new(Post),
		new(PostProduct),
		new(Brand),
		new(Product),
}
)

func DB(ctx context.Context) xorm.Interface {
	return db.NewSession()
}



func Init(xormEngine *xorm.Engine) error {
	return xormEngine.Sync(beans...)
}
func Destroy(xormEngine *xorm.Engine) error {
	return xormEngine.DropTables(beans...)
}
