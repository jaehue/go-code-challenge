package models

import (
	"context"
	"os"
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ctx context.Context
)

func TestMain(m *testing.M) {
	xormEngine, err := xorm.NewEngine("sqlite3", "prismpop.db")
	if err != nil {
		panic(err)
	}

	if err := Destroy(xormEngine); err != nil {
		panic(err)
	}
	if err := Init(xormEngine); err != nil {
		panic(err)
	}
	db = xormEngine

	os.Exit(m.Run())
}
