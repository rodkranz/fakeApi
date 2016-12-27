// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package web

import (
	"net/http"

	"github.com/rodkranz/fakeApi/module/context"
	"github.com/rodkranz/fakeApi/module/base"
	"github.com/rodkranz/fakeApi/module/fakeApi"
	"github.com/rodkranz/fakeApi/router/common"
)

const (
	DOCS_TEMPLATE base.TplName = "docs"
)

func Docs(ctx *context.Context, fakeApi *fakeApi.ApiFake)  {
	docs := &common.Docs{
		Domain: fakeApi.Domain,
		Path:   fakeApi.Folder,
	}
	docs.LoadSeeds()

	ctx.Data["Title"] = fakeApi.Domain + " Welcome [Doc FakeApi]"
	ctx.Data["Docs"]  = docs

	ctx.HTML(http.StatusOK, DOCS_TEMPLATE)
}
