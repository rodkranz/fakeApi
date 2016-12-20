// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tools

import (
	"fmt"
	"path"
	"strings"
	"encoding/json"

	"github.com/rodkranz/fakeApi/module/settings"
)

func PathToUrl(p string) (string) {
	file := strings.Replace(p, "_", "/", -1)
	file = strings.Replace(p, path.Ext(p), "", -1)
	return fmt.Sprintf("%v/%v", settings.Folder, file)
}

func UrlToPath(url string) (string) {
	file := strings.Replace(url, "/", "_", -1)
	return fmt.Sprintf("%v/%v.json", settings.Folder, file)
}

func StructToJson(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}