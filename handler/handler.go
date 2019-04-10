package handler

import (
	"bytes"
	"fmt"
	"grafana/grafanastruct"
	"net/http"
	"strings"
)

var token = "https://oapi.dingtalk.com/robot/send?access_token=7a1924682a85549535a258a9132b55270e116466bd744a12c3ad67df0cff9bec"

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
