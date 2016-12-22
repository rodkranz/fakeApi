// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package context

import (
	"log"
	"gopkg.in/macaron.v1"
	"github.com/rodkranz/fakeApi/module/settings"
	"runtime"
)

type APIContext struct {
	*Context
}

// Error render error for API
func (ctx *APIContext) Error(status int, title string, obj interface{}) {
	var message string
	if err, ok := obj.(error); ok {
		obj = err.Error()
	}

	if status == 500 {
		log.Printf("%s: %s", title, message)
	}

	ctx.JSON(status, map[string]interface{}{
		"message":  title,
		"status":   status,
		"resource": obj,
	})
}

// Render render response of api
func (ctx *APIContext) Render(status int, title string, obj interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"message":  title,
		"status":   status,
		"resource": obj,
	})
}

// APIContexter return context of macaron for API
func APIContexter() macaron.Handler {
	return func(c *Context) {
		ctx := &APIContext{
			Context: c,
		}

		ctx.Resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		ctx.Resp.Header().Set("Server", "GoLang " + runtime.Version())
		ctx.Resp.Header().Set("Developer", "Rodrigo Lopes")

		if settings.CrossDomain {
			//ctx.Resp.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Resp.Header().Set("Access-Control-Allow-Origin", ctx.Req.Host)
			ctx.Resp.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			ctx.Resp.Header().Set("Access-Control-Max-Age", "1000")
			ctx.Resp.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-Response-Code, X-Response-Daley, X-Fake-Domain")
		}

		c.Map(ctx)
	}
}