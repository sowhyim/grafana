package handler

import (
	"bytes"
	"fmt"
	"grafana/grafanastruct"
	"net/http"
	"strings"
)

var token = "https://oapi.dingtalk.com/robot/send?access_token=48f97d91f4fa1a31b7720e0ae30d76217151fc1676c9e5efd8ab034c0b86a6a1"

func ListenGrafana(in *grafanastruct.FromGrafana) {
	index := strings.Index(in.RuleURL[8:], string(47))
	fmt.Println(index,string(47),)
	url := "http://120.76.198.104:3080" + in.RuleURL[index+8:]
	fmt.Println(in.RuleURL)
	fmt.Println(url)
	msg := fmt.Sprintf(dingdingSay, in.Message, in.Title, in.ImageURL, url)
	rsp,err:=http.Post(token, "application/json", bytes.NewBuffer([]byte(msg)))
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(rsp)
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
