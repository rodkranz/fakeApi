// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package api

import (
	"fmt"
	"net/http"
	"path"
	"encoding/json"

	"github.com/rodkranz/fakeApi/modules/context"
	"github.com/rodkranz/fakeApi/modules/entity"
	"github.com/rodkranz/fakeApi/modules/fakeApi"
	"github.com/rodkranz/fakeApi/modules/files"
	"github.com/rodkranz/fakeApi/modules/base"
)

// isFileExists get url and check if file exists in seed folder
// if not exist set 404 error.
func isFileExists(ctx *context.APIContext, fake *fakeApi.ApiFake) string {
	file, err := fake.GetSeedPath(ctx.Context.Req.URL.Path[1:])

	if err != nil {
		ctx.Error(
			http.StatusNotFound,
			err.Error(),
			map[string]interface{}{
				"domain":    fake.Domain,
				"file_name": path.Base(file),
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
			"file_name": path.Base(file),
			"exception": err.Error(),
		})

	return
}

// getDataByHeaderResponseCode returns data belongs url + method + status
func getDataByHeaderResponseCode(ctx *context.APIContext, fake *fakeApi.ApiFake) interface{} {
	endpoint := ctx.Data["endpoints"].(map[string]interface{})

	method, status, has := fake.GetMethodAndStatusCode()
	if !has {
		method_status, _ := base.RandMapString(endpoint, method)
		method, status = base.SplitMethodAndStatus(method_status)
	}

	// set in context to share with application
	ctx.Data["method"] = method
	ctx.Data["status_code"] = status

	method_status := fmt.Sprintf("%v_%v", method, status)

	// if find response return and finish function
	if data, has := endpoint[method_status]; has {
		return data
	}

	// return 404 if data doesn't exist
	ctx.Error(
		http.StatusNotFound,
		"Method in seed file not found.",
		map[string]interface{}{
			"status_code": status,
			"method":      method,
			"domain":      fake.Domain,
			"file_name":   path.Base(ctx.Data["seed_file"].(string)),
		})

	return nil
}

func checkInputData(ctx *context.APIContext) {
	endpoint := ctx.Data["endpoints"].(map[string]interface{})
	// check if have format for input.
	entityExpected, has := endpoint["INPUT"]
	if !has {
		return
	}

	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.Error(
			http.StatusBadRequest,
			err.Error(), nil,
		)
		return
	}
	defer ctx.Req.Body().ReadCloser()

	entityBody := make(map[string]interface{})
	if err := json.Unmarshal(body, &entityBody); err != nil {
		ctx.Error(
			http.StatusBadRequest,
			err.Error(), nil,
		)
		return
	}

	// Validate if struct that I received is equal of documentation
	if base.EqualFormatMap(entityBody, entityExpected) {
		return
	}

	// write error of struct.
	ctx.Error(
		http.StatusBadRequest,
		"Input format is invalid with in documantation.",
		map[string]interface{}{
			"file_name":    path.Base(ctx.Data["seed_file"].(string)),
			"exected":      entityExpected,
			"received":     entityBody,
		})
}
