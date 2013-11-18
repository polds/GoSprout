package GoSprout

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	endpoint = "https://api.sproutvideo.com/v1"
)

func do(method, uri, api, values string) ([]byte, error) {
	method = strings.ToUpper(method)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}

	client := &http.Client{
		Transport: tr,
	}

	var req *http.Request
	var err error
	url := fmt.Sprintf("%s/%s", endpoint, uri)
	if method == "POST" {
		req, err = postRequest(url, values)
	} else {
		url = fmt.Sprintf("%s?%s", url, values)
		req, err = getRequest(url)
	}
	if err != nil {
		return nil, err
	}

	req.Header.Add("SproutVideo-Api-Key", api)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func postRequest(url string, body string) (*http.Request, error) {
	return http.NewRequest("POST", url, strings.NewReader(body))
}

func getRequest(url string) (*http.Request, error) {
	return http.NewRequest("GET", url, nil)
}
