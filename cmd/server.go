// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path"
	"strings"

	"gopkg.in/macaron.v1"
	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/fakeApi/modules/context"
	"github.com/rodkranz/fakeApi/modules/fakeApi"
	"github.com/rodkranz/fakeApi/modules/log"
	"github.com/rodkranz/fakeApi/modules/setting"
	"github.com/rodkranz/fakeApi/modules/template"
	"github.com/rodkranz/fakeApi/router"

	routeApi "github.com/rodkranz/fakeApi/router/api"
	routeWeb "github.com/rodkranz/fakeApi/router/web"
	"github.com/rodkranz/fakeApi/modules/versions"
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
		BaseFolder: setting.SeedFolder,
	}))

	// Static folder
	m.Use(macaron.Static(path.Join("public")))

	m.Use(context.Contexter())
	return m
}

func runServer(ctx *cli.Context) error {
	if ctx.IsSet("config") {
		setting.CustomConf = ctx.String("config")
	}

	versions.CheckTemplateVersion()
	router.GlobalInit()

	m := newMacaron()

	// Web
	m.Get("/", routeWeb.Home)
	m.Get("/docs", routeWeb.Docs)
	m.Options("*", routeWeb.HandleOptions)

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

	// Not found handler.
	m.NotFound(routeWeb.NotFound)

	// Flag for port number in case first time run conflict.
	if ctx.IsSet("port") {
		setting.AppUrl = strings.Replace(setting.AppUrl, setting.HTTPPort, ctx.String("port"), 1)
		setting.HTTPPort = ctx.String("port")
	}

	listenAddr := fmt.Sprintf("%s:%s", setting.HTTPAddr, setting.HTTPPort)
	log.Info("Listen: %v://%s%s", setting.Protocol, listenAddr, setting.AppSubUrl)

	var err error
	switch setting.Protocol {
	case setting.HTTP:
		err = http.ListenAndServe(listenAddr, m)
	case setting.HTTPS:
		server := &http.Server{Addr: listenAddr, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS10}, Handler: m}
		err = server.ListenAndServeTLS(setting.CertFile, setting.KeyFile)
	default:
		log.Fatal(4, "Invalid protocol: %s", setting.Protocol)
	}

	if err != nil {
		log.Fatal(4, "Fail to start server: %v", err)
	}
	return nil
}
