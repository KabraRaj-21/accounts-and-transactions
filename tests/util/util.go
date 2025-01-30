package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"transaction/internal/server/http/types"
)

func MakeHTTPApiCall(method, url, requestBody string) *http.Response {
	req, err := getHttpRequest(method, url, requestBody)
	if err != nil {
		fmt.Println("error occured when creating request body, err: " + err.Error())
		return nil
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error occured when making http request, err: " + err.Error())
		return nil
	}
	return resp
}

func getHttpRequest(method, url string, requestBody string) (*http.Request, error) {
	var body io.Reader
	if requestBody != "" {
		body = strings.NewReader(requestBody)
	}
	return http.NewRequest(method, url, body)
}

func ExtractResponseBody(res *http.Response, dest interface{}) {
	byteArray, _ := io.ReadAll(res.Body)
	var responseBody types.APIResponse
	json.Unmarshal(byteArray, &responseBody)

	marshalledData, _ := json.Marshal(responseBody.Data)
	json.Unmarshal(marshalledData, dest)
}
