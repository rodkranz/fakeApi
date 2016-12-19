// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tools

import "strings"

func IsSliceContainsStr(sl []string, str string) (bool, string) {
	str = strings.ToLower(str)
	for _, s := range sl {
		if strings.ToLower(s) == str {
			return true, s
		}
	}
	return false, ""
}