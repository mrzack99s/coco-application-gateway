package utils

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

func HttpJSONRequestWithBytesResponse(method, fullURL, forwardFor string, header http.Header, requestData io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, fullURL, requestData)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	needRemove := []string{"Accept-Encoding", "Origin", "Referer"}
	oldHeader := header.Clone()

	for _, v := range needRemove {
		header.Del(v)
	}

	header.Add("X-Client-IP", forwardFor)
	header.Add("X-Forward-For", forwardFor)

	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	newRespHeader := resp.Header.Clone()
	for _, v := range needRemove {
		newRespHeader.Add(v, oldHeader.Get(v))
	}
	resp.Header = newRespHeader

	return resp, err
}

func HttpGETRequest(fullURL string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
