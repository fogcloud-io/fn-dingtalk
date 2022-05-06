package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func Handler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	reqBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %s", err)
		http.Error(w, "ioutil.ReadAll: "+err.Error(), http.StatusInternalServerError)
		return
	}
	msg := DingtalkMsg{
		MsgType: "text",
		Text: struct {
			Content string "json:\"content\""
		}{
			Content: string(reqBytes),
		},
	}

	respData, err := sendDingtalkMsg(accessToken, &msg)
	if err != nil {
		log.Printf("sendDingtalkMsg: %s", err)
		http.Error(w, "sendDingtalkMsg: "+err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(respData)
	}
}

func sendDingtalkMsg(accessToken string, msg *DingtalkMsg) (respData []byte, err error) {
	rawData, _ := json.Marshal(msg)
	buf := bytes.NewBuffer(rawData)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?access_token=%s", dingtalkURL, accessToken), buf)
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("sendDingtalkMsg: %s", err)
	}
	return
}
