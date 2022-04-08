package accountapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL string = "http://localhost:8080"

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Add all the correct http error codes for different cases
func Create(newAccount *AccountData) (int, error) {
	if newAccount == nil {
		return 400, errors.New("cannot create empty account")
	}
	newAcc, err := json.Marshal(newAccount)
	if err != nil {
		return 400, err
	}
	path := "/v1/organisation/accounts"
	res, err := http.Post(baseURL+path, "application/json", bytes.NewBuffer(newAcc))

	if res.StatusCode != 200 {
		return res.StatusCode, err
	}

	return res.StatusCode, nil
}

func Fetch(id string) (*fullResponse, error) {

	if id == "" {
		return nil, errors.New("empty id")
	}

	path := "/v1/organisation/accounts/" + id
	res, err := http.Get(baseURL + path)
	if err != nil {
		log.Fatal("err -", err)
	}
	defer res.Body.Close()
	bodyBytes, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var response fullResponse

	jsonErr := json.Unmarshal(bodyBytes, &response)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Printf("res %+v\n", &response)
	return &response, nil
}
