// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package template

import (
	"github.com/rodkranz/fakeApi/module/base"
	"html/template"
)

func NewFuncMap() []template.FuncMap {
	return []template.FuncMap{map[string]interface{}{
		"SHA1":             base.EncodeSha1,
		"Marshal":          base.Marshal,
		"RandString":       base.RandString,
		"RandStringPrefix": base.RandStringPrefix,
	}}
}
