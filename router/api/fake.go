// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package api

import (
	"github.com/rodkranz/fakeApi/modules/context"
	"github.com/rodkranz/fakeApi/modules/fakeApi"
)

func FakeApi(ctx *context.APIContext, fakeApi *fakeApi.ApiFake) {

	// Validate if file exists if not render error 404.
	isFileExists(ctx, fakeApi)
	if ctx.Written() {
		return
	}

	// Load Seed file if has any error inside seed will render the error.
	loadSeedFile(ctx)
	if ctx.Written() {
		return
	}

	// Find data and retrieve
	data := getDataByHeaderResponseCode(ctx, fakeApi)
	if ctx.Written() {
		return
	}

	statusCode := ctx.Data["status_code"].(int)
	ctx.JSON(statusCode, data)
}
