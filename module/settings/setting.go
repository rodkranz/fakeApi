// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package settings

var (
	APP_VER     string
	CrossDomain bool
	Title       string
	Folder      string
)

func init() {
	CrossDomain = true
	Title = "GoLang"
	Folder = "json"
}