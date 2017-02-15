// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package web

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/rodkranz/fakeApi/modules/context"
	"github.com/rodkranz/fakeApi/modules/gitlab"
	"github.com/rodkranz/fakeApi/modules/setting"
	"github.com/rodkranz/fakeApi/modules/slack"
)

func Hook(ctx *context.Context) {
	event := ctx.Req.Header.Get("X-Gitlab-Event")
	token := ctx.Req.Header.Get("X-Gitlab-Token")

	if len(event) == 0 || len(token) == 0 {
		ctx.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"event":   event,
				"token":   token,
				"message": "Token or event invalid!",
			})
		return
	}

	confWH, has := setting.WebHooks[token]
	if !has {
		ctx.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": "Secret not found!",
				"info": map[string]interface{}{
					"secret":    token,
					"available": setting.WebHooks,
				},
			})
		return
	}

	if len(confWH.Event) != 0 && confWH.Event != event {
		ctx.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": "Not configured or this event!",
				"info": map[string]interface{}{
					"event":     event,
					"available": confWH.Event,
				},
			})
		return
	}

	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": "Error to try to read body",
				"error":   err,
			})
		return
	}

	payloadGitLab, err := gitlab.NewPayload(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": "Error to try to parse body",
				"error":   err,
			})
		return
	}

	if len(confWH.Ref) != 0 && confWH.Ref != payloadGitLab.Ref {
		ctx.JSON(http.StatusOK,
			map[string]interface{}{
				"message": "Ignore for this ref.",
			})
		return
	}

	//cmdArgs := []string{"-C", confWH.Folder, "pull"}
	cmdArgs := []string{"--work-tree=" + confWH.Folder, "--git-dir=" + confWH.Folder + "/.git", "pull"} // Git 1.7
	cmdOut, err := exec.Command("git", cmdArgs...).Output()
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			map[string]interface{}{
				"message": "There was an error running git pull command",
				"error":   err.Error(),
			})
		return
	}

	res := make(map[string]interface{}, 0)

	if setting.Slack.Active {
		slackPayload := slack.NewPayload()
		slackPayload.HydrateFromGitLab(payloadGitLab, confWH)
		output, err := slack.Notify(slackPayload)
		if err != nil {
			res["stdSlackErr"] = err
		}
		res["stdSlackOut"] = output
	}

	res["message"] = fmt.Sprintf("Fakes folder %v updated", confWH.Name)
	res["output"] = cmdOut

	ctx.JSON(http.StatusOK, res)
}
