package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/robfig/cron"
)

func createRequest() *http.Request {
	var payload = []byte(`{key: value}`)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func CollectMetrics() {
	req, err := http.NewRequest("GET", "http://127.0.0.1:9090/metrics", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, _ := client.Do(req)

	result, _ := ioutil.ReadAll(resp.Body)
	formated := fmt.Sprintf()
	log.Printf(string(result))
}

func PushConfigs() {
	req := createRequest()
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	respBody, _ := ioutil.ReadAll(response.Body)
	log.Printf(string(respBody))
}

func main() {
	cronJob := cron.New()
	cronJob.Start()
	cronJob.AddFunc("* * * * * ?", CollectMetrics)
	for {
	}
}
