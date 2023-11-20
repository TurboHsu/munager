package network

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func DoHTTPGETWithHeaders(url string, headers [][]string) (body []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, header := range headers {
		req.Header.Set(header[0], header[1])
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DoHTTPPostWithHeaders(url string, data url.Values, headers [][]string) (body []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	for _, header := range headers {
		req.Header.Set(header[0], header[1])
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DoHTTPGET(url string) (body []byte, err error) {
	return DoHTTPGETWithHeaders(url, nil)
}

func DoHTTPPost(url string, data url.Values) (body []byte, err error) {
	return DoHTTPPostWithHeaders(url, data, nil)
}
