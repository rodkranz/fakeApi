package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"gopkg.in/macaron.v1"

	"github.com/rodkranz/fakeApi/module/tools"
	"github.com/rodkranz/fakeApi/module/settings"
	"log"
	"strings"
	"strconv"
)

func init() {
	settings.AutoLoad("json")
}

func handler(ctx *macaron.Context, eps *settings.Endpoints) {
	var lookingFor string
	header := ctx.Req.Header["X-Requested-Code"]
	if len(header) == 0 {
		var e map[string]interface{}
		e = eps.Endpoints
		lookingFor = tools.RandMapString(e, ctx.Req.Method)
	} else {
		lookingFor = fmt.Sprintf("%v_%v", ctx.Req.Method, ctx.Req.Header["X-Response-Code"][0])
	}

	val, has := eps.Endpoints[lookingFor]
	if !has {
		ctx.Next()
		return
	}

	data, err := json.Marshal(val)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	status, err := strconv.ParseInt(strings.Split(lookingFor, "_")[1], 10, 64)
	if err != nil {
		status = 200
	}

	ctx.WriteHeader(int(status))
	ctx.Write(data)
}

func findMatch(ctx *macaron.Context) {
	for _, ep := range settings.Api.Eps {
		if ctx.Req.URL.Path == ep.Url {
			handler(ctx, ep)
			return
		}
	}

	ctx.Write([]byte("Not Found"))
}

func Middleware(ctx *macaron.Context) {
	ctx.Header().Add("Server", settings.Api.Title)
	ctx.Header().Add("Content-Type", "application/json")
	if settings.Api.CrossDomain {
		ctx.Header().Add("Access-Control-Allow-Origin", "*")
		ctx.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, X-Requested-Code")
		ctx.Header().Add("Access-Control-Max-Age", "86400")
	}
	ctx.Next()
}

func handleOptions(ctx *macaron.Context) string {
	ctx.Resp.WriteHeader(http.StatusOK)
	return ""
}

func main() {
	m := macaron.Classic()
	m.Use(Middleware)

	m.Options("/*", handleOptions)
	m.Any("/*", findMatch)

	log.Println("Server is running...")
	log.Println("Access from http://0.0.0.0:4000/")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}
