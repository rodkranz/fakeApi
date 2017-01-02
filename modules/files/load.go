// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package files

import (
	"errors"
	"fmt"
	"github.com/rodkranz/fakeApi/modules/entity"
	"io/ioutil"
	"path"
)

func Load(p string, i entity.Endpoint) error {
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	if err = i.Unmarshal(data); err != nil {
		return errors.New(fmt.Sprintf("Something is worng with file %s error %s", path.Base(p), err.Error()))
	}

	return nil
}
