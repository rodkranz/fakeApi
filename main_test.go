// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main_test

import (
	"testing"
	"os"
	"flag"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
