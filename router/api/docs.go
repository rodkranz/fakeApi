// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package web

import (
	"fmt"
	"path"
	"net/http"
	"io/ioutil"

	"github.com/rodkranz/fakeApi/module/fakeApi"
	"github.com/rodkranz/fakeApi/module/tools"
	"github.com/rodkranz/fakeApi/module/settings"
	"github.com/rodkranz/fakeApi/module/context"
	"github.com/rodkranz/fakeApi/module/files"
)

type Docs struct {
	Domain string
	Path   string
	Docs   []*Doc
}

func (d *Docs) LoadSeeds() {
	files, _ := ioutil.ReadDir(d.Path)
	for _, f := range files {
		if !f.IsDir() && path.Ext(f.Name()) == settings.Ext {
			doc := &Doc{
				Path: fmt.Sprintf("%v/%v", d.Path, f.Name()),
				Url:  fmt.Sprintf("/%v", tools.PathToUrl(f.Name())),
			}

			doc.LoadInfo()
			d.Docs = append(d.Docs, doc)
		}
	}
}

type Doc struct {
	Path      string
	Url       string
	Error     error
	Endpoints []*Endpoint
}

func (d *Doc) LoadInfo() {
	eps := make(map[string]interface{})
	if d.Error = files.Load(d.Path, eps); d.Error != nil {
		return
	}

	d.Endpoints = make([]*Endpoint, 0, len(eps))
	for methodAndStatus, data := range eps {
		ep := &Endpoint{}
		ep.Method, ep.StatusCode = tools.SplitMethodAndStatus(methodAndStatus)
		ep.StatusText = http.StatusText(ep.StatusCode)
		ep.Data = data

		d.Endpoints = append(d.Endpoints, ep)
	}
}

type Endpoint struct {
	Method     string
	StatusCode int
	StatusText string
	Data       interface{}
}

func ApiDocs(ctx *context.APIContext, fakeApi *fakeApi.ApiFake) {
	docs := &Docs{
		Domain: fakeApi.Domain,
		Path:   fakeApi.Folder,
	}
	docs.LoadSeeds()

	status := http.StatusOK

	ctx.JSON(status, map[string]interface{}{
		"message":  "List of endpoints available for this domain",
		"status":   status,
		"domain":   fakeApi.Domain,
		"resource": docs,
	})
}
