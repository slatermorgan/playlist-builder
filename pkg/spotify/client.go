package spotify

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Config struct {
	BaseURL string
}

type Client struct {
	config *Config
}

const contentType = "application/json"

// Needs to be a pointer
func serializeToStruct(data io.Reader, responseModel interface{}) error {
	bodyBytes, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(bodyBytes), responseModel)
	if err != nil {
		return err
	}

	return nil
}

func createReader(requestModel interface{}) (io.Reader, error) {
	requestByte, err := json.Marshal(requestModel)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(requestByte), nil
}

func (c *Client) Get(responseModel interface{}, url string) error {
	resp, err := http.Get(c.config.BaseURL + url)
	if err != nil {
		return err
	}

	err = serializeToStruct(resp.Body, responseModel)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Post(
	requestModel interface{},
	responseModel interface{},
	url string,
) error {
	body, err := createReader(requestModel)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.config.BaseURL+url, contentType, body)
	if err != nil {
		return err
	}

	err = serializeToStruct(resp.Body, responseModel)
	if err != nil {
		return err
	}

	return nil
}
