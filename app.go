package main

import (
	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/mount"
	"github.com/gohttp/response"
	"github.com/gohttp/serve"
	"net/http"
	"fmt"
)

func main() {
	response.Pretty = false

	app := app.New()
	app.Use(logger.New())
	app.Use(mount.New("/public", serve.New("public")))

	app.Get("/foo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "foo")
	}))
	app.Get("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "bar")
	})

	app.Listen(":3000")
}