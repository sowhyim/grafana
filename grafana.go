package main

import (
	"encoding/json"
	"fmt"
	"grafana/grafanastruct"
	"grafana/handler"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/alter", GrafanaHandler)
	http.ListenAndServe(":10000", nil)
}

func GrafanaHandler(w http.ResponseWriter, r *http.Request) {
	by, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("read byte from Body error :", err)
	}
	var msg grafanastruct.FromGrafana
	err = json.Unmarshal(by, &msg)
	if err != nil {
		log.Println("json.Unmarshal faild :", err)
	}
	fmt.Println(msg)
	handler.ListenGrafana(&msg)
}
