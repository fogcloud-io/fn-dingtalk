package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	accessToken = os.Getenv("ACCESS_TOKEN")
	dingtalkURL = "https://oapi.dingtalk.com/robot/send"
)

type DingtalkMsg struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// Handle a serverless request
func Handle(req []byte) string {
	msg := DingtalkMsg{
		MsgType: "text",
		Text: struct {
			Content string "json:\"content\""
		}{
			Content: string(req),
		},
	}

	respData, err := sendDingtalkMsg(accessToken, &msg)
	if err != nil {
		return fmt.Sprintf("sendDingtalkMsg: %s", err)
	} else {
		return string(respData)
	}
}

func sendDingtalkMsg(accessToken string, msg *DingtalkMsg) (respData []byte, err error) {
	rawData, _ := json.Marshal(msg)
	buf := bytes.NewBuffer(rawData)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?access_token=%s", dingtalkURL, accessToken), buf)
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	return
}
