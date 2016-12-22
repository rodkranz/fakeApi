// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package router

import (
	"fmt"
	"net/http"

	"github.com/rodkranz/fakeApi/module/context"
	"github.com/rodkranz/fakeApi/module/tools"
	"github.com/rodkranz/fakeApi/module/files"
	"github.com/rodkranz/fakeApi/module/entity"
)

// isFileExists get url and check if file exists in seed folder
// if not exist set 404 error.
func isFileExists(ctx *context.APIContext) string {
	file := tools.UrlToPath(ctx.Req.URL.Path[1:])
	if files.IsNotExist(file) {
		ctx.Error(
			http.StatusNotFound,
			"Seed file is missing",
			map[string]interface{}{
				"file_name": file,
			})
	}

	ctx.Data["seed_file"] = file
	return file
}

func loadSeedFile(ctx *context.APIContext) (endpoint map[string]interface{}) {
	file := ctx.Data["seed_file"].(string)
	endpoint = entity.Endpoint{}
	err := files.Load(file, endpoint)
	ctx.Data["endpoints"] = endpoint

	// if there is not error finish the func.
	if err == nil {
		return
	}

	// if has any error set error.
	ctx.Error(
		http.StatusInternalServerError,
		"Error to read file seed.",
		map[string]interface{}{
			"file_name": file,
			"exception": err.Error(),
		})

	return
}


// getDataByHeaderResponseCode returns data belongs url + method + status
func getDataByHeaderResponseCode(ctx *context.APIContext) interface{} {
	endpoint := ctx.Data["endpoints"].(map[string]interface{})

	// Get status code and method if it doesn't exist get random
	status, has := tools.HeaderExtract(ctx.Req.Header, "X-Response-Code")
	if has {
		status = fmt.Sprintf("%v_%v", ctx.Req.Method, status)
	} else {
		status, has = tools.RandMapString(endpoint, ctx.Req.Method)
	}


	// split method and status
	method, statusCode := tools.SplitMethodAndStatus(status)

	// set in context to share with application
	ctx.Data["method"] = method
	ctx.Data["status_code"] = statusCode

	// if find response return and finish function
	if data, has := endpoint[status]; has {
		return data
	}

	// return 404 if data doesn't exist
	ctx.Error(
		http.StatusNotFound,
		"Method in seed file not found.",
		map[string]interface{}{
			"status_code":  statusCode,
			"method":       method,
			"file_name":    ctx.Data["seed_file"],
		})

	return nil
}
