package FlagStudio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResponseToken struct {
	Code int `json:"code"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// parseToken parses the token from the response body
func parseToken(buf []byte) (string, error) {
	var data ResponseToken
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return "", err
	}
	if data.Code != 200 || data.Data.Token == "" {
		return "", fmt.Errorf("FlagStudio API, Get Token error, Code %d\n, Token: %s", data.Code, data.Data.Token)
	}
	return data.Data.Token, nil
}

// get token from the API
func GetToken(apikey string) (string, error) {
	req, err := http.NewRequest("GET", API_GET_TOKEN, nil)
	if err != nil {
		return "", fmt.Errorf("FlagStudio API, Error initializing network components, err: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("apikey", apikey)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("FlagStudio API, Error sending request, err: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("FlagStudio API, Error reading response, err: %v", err)
	}

	token, err := parseToken(body)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("FlagStudio API, Error parsing response, err: %v", err)
	}
	return token, nil
}
