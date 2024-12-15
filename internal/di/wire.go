//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Tanakaryuki/go-restapi/internal/app/handler"
	th "github.com/Tanakaryuki/go-restapi/internal/app/handler/task"
	uh "github.com/Tanakaryuki/go-restapi/internal/app/handler/user"
	tr "github.com/Tanakaryuki/go-restapi/internal/app/repository/task"
	ur "github.com/Tanakaryuki/go-restapi/internal/app/repository/user"
	ts "github.com/Tanakaryuki/go-restapi/internal/app/service/task"
	us "github.com/Tanakaryuki/go-restapi/internal/app/service/user"
	"github.com/Tanakaryuki/go-restapi/pkg/db"
	"github.com/google/wire"
)

func InitHandler() *handler.Root {
	wire.Build(
		db.New,
		th.New,
		tr.New,
		ts.New,
		uh.New,
		ur.New,
		us.New,
		handler.New,
	)
	return &handler.Root{}
}
