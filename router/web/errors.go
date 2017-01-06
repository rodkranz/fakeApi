// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package web

import (
	"fmt"
	"net/http"

	"github.com/rodkranz/fakeApi/modules/base"
	"github.com/rodkranz/fakeApi/modules/context"
)

const (
	Error404 base.TplName = "status/notFound"
)

func NotFound(ctx *context.Context) {
	ctx.Data["Title"] = "Page Not Found"

	ctx.Data["ErrorTitle"] = http.StatusNotFound
	ctx.Data["ErrorSmall"] = http.StatusText(http.StatusNotFound)
	ctx.Data["ErrorDescription"] = fmt.Sprintf("Page [%s] not found.", ctx.Req.URL.Path)

	ctx.HTML(404, Error404)
}
