package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	neturl "net/url"
	"os"
	"strings"
	"time"
)

func DoGet(url string, header map[string][]string, paramsMap map[string][]string, expectContentType string) []byte {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	params := neturl.Values{}
	for k, v := range paramsMap {
		params[k] = v
	}
	parseUrl, _ := neturl.Parse(url)
	parseUrl.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", parseUrl.String(), nil)
	for k, v := range header {
		for hv := range v {
			req.Header.Add(k, v[hv])
		}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Print("request failed:", err)
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, expectContentType) {
		print(string(body))
		return nil
	}
	return body
}

func DoPost(url string, header map[string][]string, bodyMap map[string][]string, expectContentType string) []byte {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	params := neturl.Values{}
	for k, v := range bodyMap {
		for pv := range v {
			params.Add(k, v[pv])
		}
	}
	req, _ := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	for k, v := range header {
		for hv := range v {
			req.Header.Add(k, v[hv])
		}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Print("request failed:", err)
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, expectContentType) {
		print(string(body))
		return nil
	}
	return body
}

func DoPostWithJson(url string, header map[string][]string, requestParams []byte, expectContentType string) []byte {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestParams))
	for k, v := range header {
		for hv := range v {
			req.Header.Add(k, v[hv])
		}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Print("request failed:", err)
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, expectContentType) {
		print(string(body))
		return nil
	}
	return body
}

func DoPostWithFile(url string, header map[string][]string, bodyMap map[string][]string, fileName string, filePath string, expectContentType string) []byte {
	if filePath == "" {
		DoPost(url, header, bodyMap, expectContentType)
	}
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()
	part, err := writer.CreateFormFile(fileName, file.Name())
	if err != nil {
		return nil
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil
	}
	for k, v := range bodyMap {
		for hv := range v {
			if err := writer.WriteField(k, v[hv]); err != nil {
				return nil
			}
		}
	}
	if err := writer.Close(); err != nil {
		return nil
	}
	httpRequest, _ := http.NewRequest("POST", url, requestBody)
	for k, v := range header {
		for hv := range v {
			httpRequest.Header.Add(k, v[hv])
		}
	}
	httpRequest.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	res, err := client.Do(httpRequest)
	if err != nil {
		fmt.Print("request failed:", err)
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, expectContentType) {
		print(string(body))
		return nil
	}
	return body
}
