// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"reflect"

	"github.com/rodkranz/fakeApi/modules/base"
	"github.com/rodkranz/fakeApi/modules/context"
	"github.com/rodkranz/fakeApi/modules/entity"
	"github.com/rodkranz/fakeApi/modules/fakeApi"
	"github.com/rodkranz/fakeApi/modules/files"
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

	ctx.Data["seedFile"] = file
	return file
}

// load seed file depending on route
func loadSeedFile(ctx *context.APIContext) (endpoint map[string]interface{}) {
	file := ctx.Data["seedFile"].(string)
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
	methodStatusCode := ctx.Data["methodStatusCode"].(string)

	// if find response return and finish function
	if data, has := endpoint[methodStatusCode]; has {
		return data
	}

	// return 404 if data doesn't exist
	ctx.Error(
		http.StatusNotFound,
		"Method in seed file not found.",
		map[string]interface{}{
			"status_code": ctx.Data["statusCode"],
			"method":      ctx.Data["method"],
			"domain":      fake.Domain,
			"file_name":   path.Base(ctx.Data["seedFile"].(string)),
		})

	return nil
}

func loadContextBody(ctx *context.APIContext) {
	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.Error(
			http.StatusBadRequest,
			err.Error(), nil,
		)
		return
	}
	defer ctx.Req.Body().ReadCloser()

	if len(body) == 0 {
		ctx.Data["Body"] = nil
		return
	}

	entityBody := make(map[string]interface{})
	contentType := ctx.Req.Header.Get("Content-Type")

	// try to parse json format
	if strings.Index(contentType, "application/json") != -1 {
		if err := json.Unmarshal(body, &entityBody); err != nil {
			ctx.Error(
				http.StatusBadRequest,
				err.Error(), nil,
			)
			return
		}
		ctx.Data["Body"] = entityBody
		return
	}

	// try to parse form-data format
	if strings.Index(contentType, "application/x-www-form-urlencoded") != -1 {
		values, err := url.ParseQuery(string(body))
		if err != nil {
			ctx.Error(
				http.StatusBadRequest,
				err.Error(), nil,
			)
			return
		}

		for key, val := range values {
			entityBody[key] = val[0]
		}

		ctx.Data["Body"] = entityBody
		return
	}

	if strings.Index(contentType, "multipart/form-data") != -1 {
		lines := strings.Split(string(body), string("\n"))

		var token, name, value string
		for i := 0; i < len(lines); i++ {
			line := strings.Trim(lines[i], "\n\r")
			if len(token) == 0 && strings.HasPrefix(line, strings.Repeat("-", 26)) {
				token = strings.TrimSpace(line[26:])
				continue
			}

			if strings.HasPrefix(line, strings.Repeat("-", 26)+token) {
				entityBody[name] = value
				name, value = "", ""
				continue
			}

			if strings.Index(line, "Content-Disposition: form-data;") != -1 {
				start := strings.Index(line, "name=") + 6
				name = line[start : len(line)-1]
				value = ""
				i++
				continue
			}

			if strings.HasPrefix(line, strings.Repeat("-", 26)+token+"--") {
				i = len(lines)
				continue
			}

			value += line
		}

		ctx.Data["Body"] = entityBody
		return
	}

	ctx.Data["Body"] = nil
}

func loadContextParam(ctx *context.APIContext) {
	params := make(map[string]interface{}, len(ctx.Req.URL.Query()))
	for k, v := range ctx.Req.URL.Query() {
		if len(v) == 1 {
			params[k] = v[0]
		}
	}
	ctx.Data["Params"] = params
}

// checkInputData check if has "input" at seed and match if format is correct
func checkInputData(ctx *context.APIContext) {
	endpoint := ctx.Data["endpoints"].(map[string]interface{})
	// check if have format for input.
	entityExpected, has := endpoint["INPUT"]
	if !has {
		return
	}

	if ctx.Data["Body"] == nil {
		return
	}

	// body
	entityBody := ctx.Data["Body"].(map[string]interface{})

	// Validate if struct that I received is equal of documentation
	if base.EqualFormatMap(entityBody, entityExpected) {
		return
	}

	// write error of struct.
	ctx.Error(
		http.StatusBadRequest,
		"Input format is invalid with in documantation.",
		map[string]interface{}{
			"file_name": path.Base(ctx.Data["seedFile"].(string)),
			"exected":   entityExpected,
			"received":  entityBody,
		})
}

func checkMethodAndStatus(ctx *context.APIContext, fake *fakeApi.ApiFake) {
	endpoint := ctx.Data["endpoints"].(map[string]interface{})

	method, statusCode, has := fake.GetMethodAndStatusCode()
	if !has {
		methodStatusCode, _ := base.RandMapString(endpoint, method)
		method, statusCode = base.SplitMethodAndStatus(methodStatusCode)
	}

	// set in context to share with application
	ctx.Data["method"] = method
	ctx.Data["statusCode"] = statusCode
	ctx.Data["methodStatusCode"] = fmt.Sprintf("%v_%v", method, statusCode)

}

// checkCondition if has condition in seed file
func checkCondition(ctx *context.APIContext) {
	if ctx.Data["endpoints"] == nil {
		return
	}

	endpoints := ctx.Data["endpoints"].(map[string]interface{})
	conditions, has := endpoints["CONDITIONS"].([]interface{})
	if !has {
		return
	}

	for _, v := range conditions {
		condition := v.(map[string]interface{})

		if ctx.Data["Body"] != nil {
			if reflect.DeepEqual(ctx.Data["Body"], condition["DATA"]) {
				ctx.Data["methodStatusCode"] = condition["ACTION"]

				_, statusCode := base.SplitMethodAndStatus(condition["ACTION"].(string))
				ctx.Data["statusCode"] = statusCode
				return
			}
		}

		if ctx.Data["Params"] != nil {
			fmt.Printf("OK \n")
			if reflect.DeepEqual(ctx.Data["Params"], condition["DATA"]) {
				ctx.Data["methodStatusCode"] = condition["ACTION"]
				_, statusCode := base.SplitMethodAndStatus(condition["ACTION"].(string))
				ctx.Data["statusCode"] = statusCode
				return
			}
		}

	}
}
