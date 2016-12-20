// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package settings

import "github.com/rodkranz/fakeApi/module/entity"

var (
	CrossDomain bool
	Title       string
	Folder      string
	Urls        map[string]entity.Endpoint
)

func init() {
	CrossDomain = true
	Title = "GoLang"
	Folder = "json"
	Urls = make(map[string]entity.Endpoint, 0)
}