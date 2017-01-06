// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package base

import (
	"reflect"
)

func EqualFormatMap(x, y interface{}) bool {
	if x == nil || y == nil {
		return false
	}

	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)

	if v1.Len() != v2.Len() {
		return false
	}

	if v2.Type().String() != "map[string]interface {}" || v1.Type().String() != "map[string]interface {}" {
		return false
	}

	ym := y.(map[string]interface{})
	xm := x.(map[string]interface{})
	for k := range xm {
		if _, has := ym[k]; !has {
			return false
		}
	}

	return true
}
