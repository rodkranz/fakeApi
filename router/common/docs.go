
// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/rodkranz/fakeApi/modules/base"
	"github.com/rodkranz/fakeApi/modules/files"
	"github.com/rodkranz/fakeApi/modules/setting"
)

type Docs struct {
	Domain string
	Path   string
	Info   *Info
	Groups map[int]*Group
}

type Group struct {
	Error   error
	Docs    map[string]*Doc
	Indices []string
}

func (d *Docs) LoadSeeds() {
	d.Groups = make(map[int]*Group)
	fs, _ := ioutil.ReadDir(d.Path)
	for _, f := range fs {
		if !f.IsDir() && strings.ToLower(f.Name()) == "docs.json" {
			info := &Info{
				Path: fmt.Sprintf("%v/%v", d.Path, f.Name()),
			}
			info.LoadInfo()
			d.Info = info
			continue
		}

		if !f.IsDir() && path.Ext(f.Name()) == setting.SeedExtension {
			doc := &Doc{
				Path: fmt.Sprintf("%v/%v", d.Path, f.Name()),
				Url:  fmt.Sprintf("/%v", base.PathToUrl(f.Name())),
			}
			doc.LoadInfo()

			str, has := d.Groups[doc.Level]
			if !has {
				str = &Group{
					Docs: make(map[string]*Doc),
				}
				d.Groups[doc.Level] = str
			}

			str.Docs[doc.Url] = doc
			d.Groups[doc.Level] = str
		}
	}

	if d.Info == nil {
		d.Info = &Info{Domain: d.Domain, Path: d.Path, Error: errors.New("Doc file not found.")}
	}

	for lvl, grp := range d.Groups {
		grp.Indices = make([]string, len(grp.Docs))
		i := 0
		for url := range grp.Docs {
			grp.Indices[i] = url
			i++
		}
		sort.Strings(grp.Indices)
		d.Groups[lvl] = grp
	}
}

type Info struct {
	Title       string
	SubTitle    string
	Description string
	Domain      string
	Group       map[int]string
	Path        string
	Error       error
}

func (i *Info) LoadInfo() {
	data, err := ioutil.ReadFile(i.Path)
	if err != nil {
		i.Error = err
		return
	}

	if err = json.Unmarshal(data, i); err != nil {
		i.Error = fmt.Errorf("Something is worng with file %s error %s", path.Base(i.Path), err.Error())
		return
	}
}

type Doc struct {
	Level     int
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
		d.Level = -1
		return
	}

	d.Endpoints = make([]*Endpoint, 0, len(eps))
	for methodAndStatus, data := range eps {
		ep := &Endpoint{}

		// add info about endpoint (optional)
		if strings.ToLower(methodAndStatus) == "conditions" {
			continue
		}

		if strings.ToLower(methodAndStatus) == "doc" {
			docInfo := data.(map[string]interface{})
			if title, has := docInfo["title"]; has {
				d.Title = title.(string)
			}
			if desc, has := docInfo["description"]; has {
				d.Desc = desc.(string)
			}
			if lvl, has := docInfo["level"]; has {
				iLvl, err := strconv.ParseInt(fmt.Sprintf("%v", lvl), 10, 64)
				if err == nil {
					d.Level = int(iLvl)
				}
			}
			continue
		}

		// if field name is input render the model of request
		if strings.ToLower(methodAndStatus) == "input" {
			//* if methodandstatus is input
			ep.Method = "INPUT"
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
