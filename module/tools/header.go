// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tools

import (
	"net/http"
)

func HeaderExtract(h http.Header, name string) (string, bool) {
	header, has := h[name]
	if has {
		return header[0], true
	}
	return "", false
}