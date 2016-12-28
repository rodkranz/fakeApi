// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"log"
	"net/http"
	"path"

	"gopkg.in/urfave/cli.v2"
	"gopkg.in/macaron.v1"

	"github.com/rodkranz/fakeApi/module/context"
	"github.com/rodkranz/fakeApi/module/fakeApi"
	"github.com/rodkranz/fakeApi/module/settings"
	"github.com/rodkranz/fakeApi/module/template"

	routeApi "github.com/rodkranz/fakeApi/router/api"
	routeWeb "github.com/rodkranz/fakeApi/router/web"

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
		Directory:         path.Join("templates"),
		AppendDirectories: []string{path.Join("templates")},
		Funcs:             template.NewFuncMap(),
		IndentJSON:        macaron.Env != macaron.PROD,
	}))

	m.Use(fakeApi.Register(fakeApi.ApiFakeOptions{
		DefaultApi: "default",
		BaseFolder: settings.Folder,
	}))

	// Static folder
	m.Use(macaron.Static(path.Join("public")))

	m.Use(context.Contexter())
	return m
}

func runServer(ctx *cli.Context) error {
	m := newMacaron()

	// Web
	m.Get("/", routeWeb.Home)
	m.Get("/docs", routeWeb.Docs)

	// Api
	m.Group("/api", func() {
		// Any Request with options will return 200.
		m.Options("*", routeApi.HandleOptions)

		m.Get("", routeApi.ApiDocs)

		// Fake Api Dynamic Routers
		m.Group("/", func() {
			m.Any("*", routeApi.FakeApi)
		}, context.APIContexter())
	}, context.APIContexter())


	log.Println("Server is running...")
	log.Println("Access from http://0.0.0.0:9090/")
	return http.ListenAndServe("0.0.0.0:9090", m)
}
