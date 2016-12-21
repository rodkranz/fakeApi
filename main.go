// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/macaron.v1"

	"github.com/rodkranz/fakeApi/module/settings"
	"github.com/rodkranz/fakeApi/module/files"
	"github.com/rodkranz/fakeApi/module/tools"
	"github.com/rodkranz/fakeApi/module/entity"
)

func findMatch(ctx *macaron.Context) {
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

func Middleware(ctx *macaron.Context) {
	ctx.Header().Add("Server", settings.Title)
	ctx.Header().Add("Content-Type", "application/json")
	if settings.CrossDomain {
		ctx.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-Response-Code")
		ctx.Header().Add("Access-Control-Max-Age", "86400")
	}
	ctx.Next()
}

func handleOptions(ctx *macaron.Context) {
	ctx.Resp.WriteHeader(http.StatusOK)
}

func notFound(ctx *macaron.Context) string {
	ctx.Resp.WriteHeader(http.StatusNotFound)
	return "Not found"
}
func internalServerError(ctx *macaron.Context, err error) string {
	ctx.Resp.WriteHeader(http.StatusInternalServerError)
	return "Internal server Error: " + err.Error()
}

func main() {
	m := macaron.Classic()
	m.Use(Middleware)

	m.Options("/*", handleOptions)
	m.Any("*", findMatch)

	m.NotFound(notFound)
	m.InternalServerError(internalServerError)

	log.Println("Server is running...")
	log.Println("Access from http://0.0.0.0:9090/")
	log.Println(http.ListenAndServe("0.0.0.0:9090", m))
}
