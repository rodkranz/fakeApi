// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package api

import (
	"net/http"

	"github.com/rodkranz/fakeApi/module/context"
)

func HandleOptions(ctx *context.APIContext) {
	ctx.Resp.WriteHeader(http.StatusOK)
}

func NotFound(ctx *context.APIContext) {
	status := http.StatusNotFound

	ctx.JSON(status, map[string]interface{}{
		"message":  http.StatusText(status),
		"status":   status,
		"resource": nil,
	})
}

func InternalServerError(ctx *context.APIContext) {
	status := http.StatusInternalServerError

	ctx.JSON(status, map[string]interface{}{
		"message":  http.StatusText(status),
		"status":   status,
		"resource": nil,
	})
}
