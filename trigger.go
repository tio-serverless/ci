package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

const travisUrl = "https://api.travis-ci.com/repo/%s/requests"
const message = `{
	"request": {
		"branch": "%s",
		"message": "Trigger by api v3"
	}
}`

func trigger(branch, name, token string) error {
	name = url.PathEscape(name)
	travisUrl := fmt.Sprintf(travisUrl, name)

	msg := fmt.Sprintf(message, branch)

	log.Debugf("url: %s", travisUrl)
	log.Debugf("body: %s", msg)

	payload := strings.NewReader(msg)

	req, _ := http.NewRequest("POST", travisUrl, payload)

	req.Header.Add("Travis-API-Version", "3")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	log.Debugf("Travis Returns [%s]", string(body))

	if res.StatusCode < 300 && res.StatusCode >= http.StatusOK {
		return nil
	}

	return fmt.Errorf("Travis Returns %d", res.StatusCode)
}
