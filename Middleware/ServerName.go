// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package Middleware

import (
	"gopkg.in/macaron.v1"
	"github.com/rodkranz/fakeApi/module/settings"
)

func ServerName(ctx *macaron.Context) {
	ctx.Header().Add("Server", settings.Title)
	ctx.Next()
}