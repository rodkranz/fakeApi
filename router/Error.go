// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package router

import (
	"gopkg.in/macaron.v1"
	"net/http"
)

func HandleOptions(ctx *macaron.Context) {
	ctx.Resp.WriteHeader(http.StatusOK)
}

func NotFound(ctx *macaron.Context) string {
	ctx.Resp.WriteHeader(http.StatusNotFound)
	return "Not found"
}
func InternalServerError(ctx *macaron.Context, err error) string {
	ctx.Resp.WriteHeader(http.StatusInternalServerError)
	return "Internal server Error: " + err.Error()
}
