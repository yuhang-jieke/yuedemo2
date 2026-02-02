package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestUpload(t *testing.T) {
	url := "http://127.0.0.1:8080/upload"

	payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"address\"\r\n\r\nC:\\Users\\ZhuanZ\\Pictures\\Screenshots\\屏幕截图 2025-09-11 094201.png\r\n-----011000010111000001101001--\r\n\r\n")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Njk5MjY4MDgsImlhdCI6MTc2OTkyMzIwOCwidXNlcklkIjoiMTAifQ.O2mBBqchfs8w3Gid7Wd_As_F4C-HMays4rsMIvEIrE0")
	req.Header.Add("content-type", "multipart/form-data; boundary=---011000010111000001101001")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
