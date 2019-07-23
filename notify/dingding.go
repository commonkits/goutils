package notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://oapi.dingtalk.com/robot/send?access_token="
)

type At struct {
	IsAtAll   bool     `json:"isAtAll"`
	AtMobiles []string `json:"atMobiles"`
}

type Text struct {
	Content string `json:"content"`
}

type TextMsg struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
	At      At     `json:"at"`
}

type DingResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func SendText(token string, content string, atMobiles ...string) error {
	textMsg := TextMsg{
		Msgtype: "text",
		Text: Text{
			Content: content,
		},
		At: At{
			AtMobiles: atMobiles,
		},
	}
	dingRequest, err := json.Marshal(textMsg)
	if err != nil {
		return err
	}

	resp, err := http.Post(url+token, "application/json", bytes.NewBuffer(dingRequest))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dingResp DingResp
	err = json.Unmarshal(body, &dingResp)
	if err != nil {
		return err
	}

	if dingResp.Errcode != 0 {
		return errors.New(dingResp.Errmsg)
	}
	return nil
}

func SendTextAtAll(token string, content string) error {
	textMsg := TextMsg{
		Msgtype: "text",
		Text: Text{
			Content: content,
		},
		At: At{
			IsAtAll: true,
		},
	}
	dingRequest, err := json.Marshal(textMsg)
	if err != nil {
		return err
	}

	resp, err := http.Post(url+token, "application/json", bytes.NewBuffer(dingRequest))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dingResp DingResp
	err = json.Unmarshal(body, &dingResp)
	if err != nil {
		return err
	}

	if dingResp.Errcode != 0 {
		return errors.New(dingResp.Errmsg)
	}
	return nil
}
