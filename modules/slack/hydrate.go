// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package slack

import (
	"fmt"
	"time"

	"github.com/rodkranz/fakeApi/modules/gitlab"
	"github.com/rodkranz/fakeApi/modules/setting"
)

func (p *Payload) HydrateFromGitLab(g *gitlab.Payload, w *setting.Webhook) {
	p.Username = "GoBot"
	p.Channel = fmt.Sprintf("#%v", w.Channel)
	p.Username = setting.Slack.Name
	p.Emotion = setting.Slack.Icon

	p.Text = fmt.Sprintf("[<%v|%v>] The %v has been updated!", g.Project.WebURL, g.Repository.Name, w.Name)

	for _, commit := range g.Commits {
		a := &Attachment{}
		a.AuthorName = commit.Author.GetName()

		a.TitleLink = g.Project.WebURL + "/commit/" + commit.ID
		a.Title = commit.ID[0:7]
		a.Text = commit.Message
		a.Color = "#36a64f"

		p.AppendAttachment(a)
	}

	p.AppendAttachment(&Attachment{
		Footer:     setting.Slack.Name,
		FooterIcon: setting.Slack.Avatar,
		Ts:         int32(time.Now().Local().Unix()),
	})
}
