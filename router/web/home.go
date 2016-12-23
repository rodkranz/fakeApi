// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package web

import (
	"fmt"
	"io/ioutil"
	"path"
	"net/http"

	"github.com/rodkranz/fakeApi/module/settings"
	"github.com/rodkranz/fakeApi/module/context"
	"github.com/rodkranz/fakeApi/module/fakeApi"
	"github.com/rodkranz/fakeApi/module/tools"
	"github.com/rodkranz/fakeApi/module/base"
)

/**
 * TODO: Use the template system.
 */
const homeTemplate base.TplName = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link href="https://fonts.googleapis.com/css?family=Open+Sans+Condensed:300" rel="stylesheet">
        <title>Fake Api</title>
        <style>
            body {
                font-family: 'Open Sans Condensed', sans-serif;
                font-size: 20px;
            }
            .container {
                margin-top: 10%;
                padding-left: 20%;padding-right: 20%; text-align: center
            }
            .page-header small {
                color: #696666;
            }
            .page-header {
                border-bottom: #ccc 1px solid;
            }
        </style>
    </head>
    <body>
    <div class="container">
        <div class="page-header">
            <h1>fake api
            <small>is working</small></h1>
        </div>
        <p class="lead">Right now you don't need to wait for api any more.</p>
    </div>
</html>`

func listFiles(p string) []string {
	list := make([]string, 0)
	files, _ := ioutil.ReadDir(p)
	for _, f := range files {
		if !f.IsDir() && path.Ext(f.Name()) == settings.Ext {
			list = append(list, f.Name())
		}
	}
	return list
}

func Home(ctx *context.Context) {
	ctx.WriteHeader(http.StatusOK)
	ctx.Write([]byte(homeTemplate))
}

func Docs(ctx *context.APIContext, fakeApi *fakeApi.ApiFake) {
	list := listFiles(fakeApi.Folder)
	data := make([]interface{}, len(list))
	for idx, n := range list {
		data[idx] = fmt.Sprintf("/%v", tools.PathToUrl(n))
	}

	status := http.StatusOK

	ctx.JSON(status, map[string]interface{}{
		"message":  "List of endpoints available for this domain",
		"status":   status,
		"domain":   fakeApi.Domain,
		"resource": data,
	})
}
