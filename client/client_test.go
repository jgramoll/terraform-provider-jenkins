package client

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var client *Client
var testPath string

func init() {
	testPath = "/api/xml"
	client = newTestClient()
}

func TestClientNewRequest(t *testing.T) {
	req, err := client.NewRequest("get", testPath)
	if err != nil {
		t.Fatal(err)
	}
	expectedURL := client.Config.Address + testPath
	if req.URL.String() != expectedURL {
		t.Fatalf("request url should be %#v, not %#v", expectedURL, req.URL.String())
	}
}

type TestBody struct {
	Field string `xml:"field"`
}

func TestClientNewRequestWithBody(t *testing.T) {
	body := TestBody{Field: "#value"}
	req, err := client.NewRequestWithBody("get", testPath, body)
	if err != nil {
		t.Fatal(err)
	}
	byteBody, bodyErr := ioutil.ReadAll(req.Body)
	if bodyErr != nil {
		t.Fatal(bodyErr)
	}

	actualBody := string(byteBody)
	expectedBody := "<TestBody><field>#value</field></TestBody>"
	if actualBody != expectedBody {
		t.Fatalf("request body should be %#v, not %#v", expectedBody, actualBody)
	}
}

func newTestClient() *Client {
	address := os.Getenv("JENKINS_ADDRESS")
	if address == "" {
		log.Println("[Error] JENKINS_ADDRESS not defined")
	}
	username := os.Getenv("JENKINS_USERNAME")
	if username == "" {
		log.Println("[Error] JENKINS_USERNAME not defined")
	}
	token := os.Getenv("JENKINS_TOKEN")
	if token == "" {
		log.Println("[Error] JENKINS_TOKEN not defined")
	}

	c := Config{
		Address:  address,
		Username: username,
		Token:    token,
	}
	return NewClient(c)
}
