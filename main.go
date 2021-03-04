package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	projectId   = 131206
	summary     = "Sample task"
	issueTypeId = 737474
	priorityId  = 3
	Logger      *log.Logger
)

func main() {
	lambda.Start(slackToBacklog)
	Logger.Printf("finish.")
}

func slackToBacklog() (string, error) {
	if os.Getenv("HOST") == "" {
		return "FAILED!! Check To Log", fmt.Errorf("環境変数%sがセットされていません", "HOST")
	}
	host := os.Getenv("HOST")

	if os.Getenv("API_KEY") == "" {
		return "FAILED!! Check To Log", fmt.Errorf("環境変数%sがセットされていません", "API_KEY")
	}
	apiKey := os.Getenv("API_KEY")

	values := url.Values{
		"projectId":   {strconv.Itoa(projectId)},
		"summary":     {summary},
		"issueTypeId": {strconv.Itoa(issueTypeId)},
		"priorityId":  {strconv.Itoa(priorityId)},
	}
	client := http.DefaultClient
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v2/issues?apiKey=%s", host, apiKey), strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	println(string(body))

	return "SUCCESS!", nil
}
