// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"log"
	"net/http"

	"gopkg.in/urfave/cli.v2"
	"gopkg.in/macaron.v1"

	"github.com/rodkranz/fakeApi/router"
	"github.com/rodkranz/fakeApi/Middleware"
)

var CmdServer = &cli.Command{
	Name:        "server",
	Usage:       "Run Fake API Server",
	Description: `Start server fake.`,
	Action:      runServer,
	Flags:       []cli.Flag{},
}

func newMacaron() *macaron.Macaron {
	m := macaron.New()

	// Server name
	m.Use(Middleware.ServerName)
	// Cross domain
	m.Use(Middleware.CrossDomain)

	return m
}


func runServer(c *cli.Context) error {
	m := newMacaron()

	m.Options("/*", router.HandleOptions)
	m.Any("*", router.FakeApi)

	m.NotFound(router.NotFound)
	m.InternalServerError(router.InternalServerError)

	log.Println("Server is running...")
	log.Println("Access from http://0.0.0.0:9090/")
	return http.ListenAndServe("0.0.0.0:9090", m)
}
