// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"log"
	"net/http"

	"gopkg.in/urfave/cli.v2"
	"gopkg.in/macaron.v1"

	routeApi "github.com/rodkranz/fakeApi/router/api"
	"github.com/rodkranz/fakeApi/module/context"
	"github.com/rodkranz/fakeApi/module/fakeApi"
	"github.com/rodkranz/fakeApi/module/settings"
)

var Server = &cli.Command{
	Name:        "server",
	Usage:       "Run Fake API Server",
	Description: `Start server fake.`,
	Action:      runServer,
	Flags:       []cli.Flag{},
}

func newMacaron() *macaron.Macaron {
	m := macaron.New()

	m.Use(macaron.Renderer(macaron.RenderOptions{
		IndentJSON:        macaron.Env != macaron.PROD,
	}))

	m.Use(fakeApi.Register(fakeApi.ApiFakeOptions{
		DefaultApi: "default",
		BaseFolder: settings.Folder,
	}))

	m.Use(context.Contexter())
	return m
}

func runServer(ctx *cli.Context) error {
	m := newMacaron()

	m.Group("/api", func() {
		// Any Request with options will return 200.
		m.Options("/*", routeApi.HandleOptions)

		// Fake Api Dynamic Routers
		m.Group("/", func() {

			m.Any("/*", routeApi.FakeApi)

		}, context.APIContexter())
	})


	log.Println("Server is running...")
	log.Println("Access from http://0.0.0.0:9090/")
	return http.ListenAndServe("0.0.0.0:9090", m)
}
