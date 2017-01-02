// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package base

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

// Encode string to sha1 hex value.
func EncodeSha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Marshal(i interface{}) string {
	data, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return ""
	}

	return string(data)
}
