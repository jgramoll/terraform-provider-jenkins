package client

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ErrInvalidDecodeResponseParameter invalid parameter for decodeResponse
var ErrInvalidDecodeResponseParameter = errors.New("nil interface provided to decodeResponse")

// Config for Client
type Config struct {
	Address  string
	Username string
	Token    string
}

// Client to talk to Jenkins
type Client struct {
	Config Config
	client *http.Client
}

// NewClient Return a new client with loaded configuration
func NewClient(config Config) *Client {
	httpClient := &http.Client{}

	return &Client{
		Config: config,
		client: httpClient,
	}
}

// NewRequest create http request
func (client *Client) NewRequest(method string, path string) (*http.Request, error) {
	return client.NewRequestWithBody(method, path, nil)
}

// NewRequestWithBody create http request with data as body
func (client *Client) NewRequestWithBody(method string, path string, data interface{}) (*http.Request, error) {
	reqURL, urlErr := url.Parse(client.Config.Address + path)
	if urlErr != nil {
		return nil, urlErr
	}

	xmlValue, xmlErr := xml.Marshal(data)
	if xmlErr != nil {
		return nil, xmlErr
	}

	log.Printf("[INFO] Sending %s %s with body %s\n", method, reqURL, xmlValue)
	req, err := http.NewRequest(method, reqURL.String(), bytes.NewBuffer(xmlValue))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/xml;charset=UTF-8")
	return req, nil
}

// Do send http request
func (client *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := client.do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// DoWithResponse send http request and parse response body
func (client *Client) DoWithResponse(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := client.do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	err = decodeResponse(resp, v)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// do internal function used by Do and DoWithResponse to validate response
func (client *Client) do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(client.Config.Username, client.Config.Token)

	resp, err := client.client.Do(req)
	if err != nil {
		return resp, err
	}

	err = validateResponse(resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func decodeResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return ErrInvalidDecodeResponseParameter
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	log.Println("[DEBUG] Got response body", bodyString)

	// TODO hack around xml1.1 and xm1l.0
	// https://github.com/golang/go/issues/25755
	bodyString = strings.Replace(bodyString, "version='1.1'", "version='1.0'", 1)

	return xml.Unmarshal([]byte(bodyString), &v)
}

func validateResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	log.Println("[INFO] Error response body", bodyString)

	jenkinsError := JenkinsError{}
	err := xml.Unmarshal([]byte(bodyString), &jenkinsError)
	if err != nil {
		return err
	}

	return &jenkinsError
}
