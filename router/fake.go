// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package router

import (
	"net/http"
	"log"
	"github.com/rodkranz/fakeApi/module/tools"
	"gopkg.in/macaron.v1"
	"github.com/rodkranz/fakeApi/module/files"
	"github.com/rodkranz/fakeApi/module/entity"
	"fmt"
)

func FakeApi(ctx *macaron.Context) {
	file := tools.UrlToPath(ctx.Req.URL.Path[1:])
	if files.IsNotExist(file) {
		ctx.WriteHeader(http.StatusNotFound)
		ctx.Write([]byte("{\"error\": \"Seed file is missing.\", \"file\": \"" + file + "\"}"))
		return
	}

	var endpoint map[string]interface{} = entity.Endpoint{}
	if err := files.Load(file, endpoint); err != nil {
		log.Printf("Error to load file %v: %v", file, err.Error())
		ctx.WriteHeader(http.StatusInternalServerError)
		ctx.Write([]byte("{\"error\": \"Load seed file\", \"detail\": \"" + err.Error() + "\"}"))
		return
	}

	status, has := tools.HeaderExtract(ctx.Req.Header, "X-Response-Code")
	if has {
		status = fmt.Sprintf("%v_%v", ctx.Req.Method, status)
	} else {
		status, has = tools.RandMapString(endpoint, ctx.Req.Method)
	}

	if !has {
		log.Printf("Error to find rule: %v", status)
		ctx.WriteHeader(http.StatusNotFound)
		ctx.Write([]byte("{\"error\": \"Status not found\", \"detail\": \"" + status + "\"}"))
		return
	}

	resp, has := endpoint[status]
	if !has {
		log.Printf("Status not found: %v", status)
		ctx.WriteHeader(http.StatusNotFound)
		ctx.Write([]byte("{\"error\": \"Url/Status not found\", \"detail\": \"" + status + "\"}"))
		return
	}

	data, err := tools.StructToJson(resp)
	if err != nil {
		log.Printf("Error to convert struct to json: %v", err)
		ctx.WriteHeader(http.StatusInternalServerError)
		ctx.Write([]byte("{\"error\": \"Seed file has error\", \"detail\": \"" + err.Error() + "\"}"))
		return
	}

	_, statusCode := tools.SplitMethodAndStatus(status)

	ctx.WriteHeader(statusCode)
	ctx.Write(data)
}
