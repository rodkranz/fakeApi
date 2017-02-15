// Package slack contain send actions
package slack

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"fmt"
	"github.com/rodkranz/fakeApi/modules/setting"
)

// Send payload to slack for publish
func Notify(s *Payload) (string, error) {
	if !setting.Slack.Active {
		return "", &ErrSlackIsNotActivated{}
	}

	data, err := s.JSON()
	if err != nil {
		return "", err
	}

	fmt.Print(string(data))

	v := url.Values{}
	v.Set("payload", string(data))
	res, err := http.PostForm(setting.Slack.API, v)
	if err != nil {
		return "", &ErrSlackPost{Err: err}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", &ErrSlackReadBody{Err: err}
	}

	return string(body), nil
}
