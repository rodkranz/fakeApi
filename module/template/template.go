// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package template

import (
	"html/template"
	"github.com/rodkranz/fakeApi/module/base"
)

func NewFuncMap() []template.FuncMap {
	return []template.FuncMap{map[string]interface{}{
		"SHA1":            base.EncodeSha1,
		"Marshal":         base.Marshal,
	}}
}
