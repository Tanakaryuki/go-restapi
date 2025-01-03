package main

import (
	"fmt"
	"net/http"

	"github.com/Tanakaryuki/go-restapi/internal/di"
	"github.com/Tanakaryuki/go-restapi/pkg/config"
	"github.com/Tanakaryuki/go-restapi/pkg/handler"
	"github.com/Tanakaryuki/go-restapi/pkg/log"
	"github.com/Tanakaryuki/go-restapi/pkg/middleware"
)

func main() {
	config.LoadEnv()

	h := di.InitHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.Handle("POST /signup", handler.AppHandler(h.UserHander.CreateUser()))
	mux.Handle("POST /login", handler.AppHandler(h.UserHander.Login()))
	mux.Handle("GET /me", middleware.Auth(handler.AppHandler(h.UserHander.GetMe())))

	mux.Handle("GET /tasks/{id}", middleware.Auth(handler.AppHandler(h.TaskHander.GetTask())))
	mux.Handle("POST /tasks", middleware.Auth(handler.AppHandler(h.TaskHander.CreateTask())))

	loggedMux := log.LoggingMiddleware(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	server.ListenAndServe()
}
