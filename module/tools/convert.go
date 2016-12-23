// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tools

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/rodkranz/fakeApi/module/settings"
)

func PathToUrl(p string) string {
	file := strings.Replace(p, "_", "/", -1)
	file = strings.Replace(p, path.Ext(p), "", -1)
	return fmt.Sprintf("%v/%v", settings.Folder, file)
}

func StructToJson(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func SplitMethodAndStatus(s string) (method string, code int) {
	if !strings.Contains(s, "_") {
		return method, 200
	}

	splitted := strings.Split(s, "_")
	method = splitted[0]
	if i, err := strconv.ParseInt(splitted[1], 10, 64); err != nil {
		code = 200
	} else {
		code = int(i)
	}
	return
}
