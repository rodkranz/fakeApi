// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package router

import (
	"github.com/rodkranz/fakeApi/modules/log"
	"github.com/rodkranz/fakeApi/modules/setting"
)

func GlobalInit() {
	setting.NewContext()
	log.Trace("Custom path: %s", setting.CustomPath)
	log.Trace("Log path: %s", setting.LogRootPath)

	setting.NewServices()
}
