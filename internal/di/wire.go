//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Tanakaryuki/go-restapi/internal/app/handler"
	th "github.com/Tanakaryuki/go-restapi/internal/app/handler/task"
	tr "github.com/Tanakaryuki/go-restapi/internal/app/repository/task"
	ts "github.com/Tanakaryuki/go-restapi/internal/app/service/task"
	"github.com/Tanakaryuki/go-restapi/pkg/db"
	"github.com/google/wire"
)

func InitHandler() *handler.Root {
	wire.Build(
		db.New,
		th.New,
		tr.New,
		ts.New,
		handler.New,
	)
	return &handler.Root{}
}
