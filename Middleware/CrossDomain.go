// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package Middleware

import (
	"gopkg.in/macaron.v1"
	"github.com/rodkranz/fakeApi/module/settings"
)

func CrossDomain(ctx *macaron.Context) {
	if settings.CrossDomain {
		ctx.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-Response-Code, Cache-Control")
		ctx.Header().Add("Access-Control-Max-Age", "86400")
	}
	ctx.Next()
}
