// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package settings

import (
	"os"
	"path/filepath"
	"path"
	"io/ioutil"
	"log"
	"strings"
	"fmt"
)

var Api *Config

func init() {
	Api = &Config{
		CrossDomain: true,
		Title:       "GoLang",
		Eps:         make([]*Endpoints, 0),
	}
}

func AutoLoad(folder string) {
	filepath.Walk(folder, func(p string, i os.FileInfo, err error) error {
		if path.Ext(p) != ".json" {
			return nil
		}

		return LoadInfo(p)
	})
}

func LoadInfo(p string) error {
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	ep := &Endpoint{}
	if err = ep.Unmarshal(data); err != nil {
		log.Fatalf("Something is worng with settings: %v", err.Error())
		os.Exit(1)
	}

	name := path.Base(p)
	name = strings.Replace(name, path.Ext(p), "", -1)
	name = strings.Replace(name, "_", "/", -1)

	eps := &Endpoints{
		Url: fmt.Sprintf("/%v", name),
		Endpoints: *ep,
	}

	Api.RegisterEndpoints(eps)
	return nil
}
