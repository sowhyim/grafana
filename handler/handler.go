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
	url := strings.Replace(in.RuleURL, "localhost", "10.0.0.30", 1)
	msg := fmt.Sprintf(dingdingSay, in.Message, in.Title, in.ImageURL, url)
	http.Post(token, "application/json", bytes.NewBuffer([]byte(msg)))
}

func handlerUrl(url string) {

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
