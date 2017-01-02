// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/rodkranz/fakeApi/modules/base"
	"github.com/rodkranz/fakeApi/modules/files"
	"github.com/rodkranz/fakeApi/modules/setting"
)

type Docs struct {
	Domain string
	Path   string
	Docs   []*Doc
}

func (d *Docs) LoadSeeds() {
	fs, _ := ioutil.ReadDir(d.Path)
	for _, f := range fs {
		if !f.IsDir() && path.Ext(f.Name()) == setting.SeedExtension {
			doc := &Doc{
				Path: fmt.Sprintf("%v/%v", d.Path, f.Name()),
				Url:  fmt.Sprintf("/%v", base.PathToUrl(f.Name())),
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
	Title     string
	Desc      string
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

		// add info about endpoint (optional)
		if strings.ToLower(methodAndStatus) == "doc" {
			docInfo := data.(map[string]interface{})
			if title, has := docInfo["title"]; has {
				d.Title = title.(string)
			}
			if desc, has := docInfo["description"]; has {
				d.Desc = desc.(string)
			}
			continue
		}

		// if field name is input render the model of request
		if strings.ToLower(methodAndStatus) == "input" {
			//* if methodandstatus is input
			ep.Method = "IMPUT"
			ep.StatusCode = 0
			ep.Data = data
		} else {
			// if find dynamic values
			ep.Method, ep.StatusCode = base.SplitMethodAndStatus(methodAndStatus)
			ep.StatusText = http.StatusText(ep.StatusCode)
			ep.Data = data

		}

		d.Endpoints = append(d.Endpoints, ep)
	}
}

type Endpoint struct {
	Method     string
	StatusCode int
	StatusText string
	Data       interface{}
}
