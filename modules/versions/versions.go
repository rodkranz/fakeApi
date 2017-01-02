// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package versions

import (
	"io/ioutil"

	"github.com/rodkranz/fakeApi/modules/log"
	"github.com/rodkranz/tmp/modules/setting"
)

// checkVersion checks if binary matches the version of templates files.
func CheckTemplateVersion() {
	// Templates.
	data, err := ioutil.ReadFile(setting.StaticRootPath + "/templates/.VERSION")
	if err != nil {
		log.Fatal(4, "Fail to read 'templates/.VERSION': %v", err)
	}

	if string(data) != setting.AppVer {
		log.Fatal(4, "Binary and template file version does not match, did you forget to recompile?")
	}
}
