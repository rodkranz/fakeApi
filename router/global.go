// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package router

import (
	"strings"

	"gopkg.in/macaron.v1"

	"github.com/rodkranz/fakeApi/modules/log"
	"github.com/rodkranz/fakeApi/modules/setting"
)

func checkRunMode() {
	switch setting.Cfg.Section("").Key("RUN_MODE").String() {
	case "prod":
		macaron.Env = macaron.PROD
		macaron.ColorLog = false
		setting.ProdMode = true
	}
	log.Info("Run Mode: %s", strings.Title(macaron.Env))
}

func GlobalInit() {
	setting.NewContext()
	log.Trace("Custom path: %s", setting.CustomPath)
	log.Trace("Log path: %s", setting.LogRootPath)

	setting.NewServices()
	checkRunMode()
}
