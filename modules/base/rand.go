// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package base

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandStringPrefix(prefix string, l int) string {
	return fmt.Sprintf("%s-%s", prefix, RandString(l))
}

func RandString(l int) string {
	bytes := []byte{}
	for len(bytes) < l {
		if RandInt(0, 1)%2 == 0 {
			bytes = append(bytes, byte(RandInt(65, 90)))
		}
		if RandInt(0, 1)%2 == 0 {
			bytes = append(bytes, byte(RandInt(97, 122)))
		}
		if RandInt(0, 1)%2 == 0 {
			bytes = append(bytes, byte(RandInt(48, 57)))
		}
	}
	return string(bytes[:l])
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

func RandMapString(m map[string]interface{}, prefix string) (string, bool) {
	keys := make([]string, 0, len(m))
	for k := range m {
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
	}
	if len(keys) == 0 {
		return "", false
	}
	return keys[RandInt(0, len(keys))], true
}
