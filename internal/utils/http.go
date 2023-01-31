package utils

import (
	"crypto/tls"
	"io"
	"net/http"
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

	header.Del("Accept-Encoding")
	header.Add("X-Client-IP", forwardFor)
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
