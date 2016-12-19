// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tools

import (
	"time"
	"math/rand"
	"strings"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max - min)
}

func RandString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandSliceString(list []string) string {
	return list[RandInt(0, len(list))]
}

func RandSliceInt(list []int) int {
	return list[RandInt(0, len(list))]
}

func RandSliceInterface(list []interface{}) interface{} {
	return list[RandInt(0, len(list))]
}

func RandMapString(m map[string]interface{}, prefix string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
	}
	return keys[RandInt(0, len(keys))]
}