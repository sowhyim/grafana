package handler

import (
	"bytes"
	"fmt"
	"grafana/grafanastruct"
	"net/http"
	"strings"
)

var token = "https://oapi.dingtalk.com/robot/send?access_token=8385c8b14e809f8fc0aa2dfe5351a87ae5dfa32e7b4f06de0b56bb7b18da65ec"

func ListenGrafana(in *grafanastruct.FromGrafana) {
	index := strings.Index(in.RuleURL[8:], string(47))
	url := "http://120.76.198.104:3080" + in.RuleURL[index+8:]
	msg := fmt.Sprintf(dingdingSay, in.Message, in.Title, in.ImageURL, url)
	fmt.Println(msg)
	_,err:=http.Post(token, "application/json", bytes.NewBuffer([]byte(msg)))
	if err!=nil{
		fmt.Println(err)
	}
}

func SendMSG(in string) bool {
	ch := strings.Split(in, " ")
	if len(ch) != 3 {
		return false
	}

	return true
}

var dingdingSay = `{
    "msgtype": "link", 
    "link": {
        "text": "%s", 
        "title": "%s", 
        "picUrl": "%s", 
        "messageUrl": "%s"
    }
}`
