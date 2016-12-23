// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tools

import (
	"strconv"
	"strings"
	"path"
)

func PathToUrl(p string) string {
	p = strings.Replace(p, path.Ext(p), "", 1)

	p = strings.Replace(p, "__", "#", -1)
	p = strings.Replace(p, "_", "/", -1)
	p = strings.Replace(p, "#", "_", -1)

	return p
}

func UrLToPath(p string) string {
	p = strings.Replace(p, "__", "#", -1)
	p = strings.Replace(p, "/", "_", -1)
	p = strings.Replace(p, "#", "_", -1)

	return p
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
