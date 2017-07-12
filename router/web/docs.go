// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package web

import (
	"net/http"

	"github.com/rodkranz/fakeApi/modules/base"
	"github.com/rodkranz/fakeApi/modules/context"
	"github.com/rodkranz/fakeApi/modules/fakeApi"
	"github.com/rodkranz/fakeApi/router/common"
)

const (
	DOCS_TEMPLATE base.TplName = "docs"
)

func Docs(ctx *context.Context, fakeApi *fakeApi.ApiFake) {
	docs := &common.Docs{
		Domain: fakeApi.Domain,
		Path:   fakeApi.Folder,
	}
	docs.LoadSeeds()

	ctx.Data["Title"] = "[Docs FakeApi] " + fakeApi.Domain
	ctx.Data["Doc"] = docs
	ctx.Data["Domain"] = fakeApi.Domain
	ctx.Data["IsSingle"] = len(docs.Groups) == 1

	ctx.HTML(http.StatusOK, DOCS_TEMPLATE)
}
