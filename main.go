// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"runtime"

	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/fakeApi/cmd"
	"github.com/rodkranz/fakeApi/module/settings"
	"os"
)

const VER = "1.0.0"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	settings.APP_VER = VER
}

func main() {
	app := cli.App{
		Name: "FakeApi",
		Usage: "Build a api server",
		Version: VER,
		Commands: []*cli.Command{
			cmd.CmdServer,
		},
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
