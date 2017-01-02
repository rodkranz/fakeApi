// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package web

import (
	"net/http"

	"github.com/rodkranz/fakeApi/modules/base"
	"github.com/rodkranz/fakeApi/modules/context"
	"github.com/rodkranz/fakeApi/modules/fakeApi"
)

const (
	HOME_TEMPLATE base.TplName = "home"
)

func Home(ctx *context.Context, fakeApi *fakeApi.ApiFake) {
	ctx.Data["Title"] = "[FakeApi] " + fakeApi.Domain

	ctx.HTML(http.StatusOK, HOME_TEMPLATE)
}

func HandleOptions(ctx *context.Context) {
	ctx.Resp.WriteHeader(http.StatusOK)
}
