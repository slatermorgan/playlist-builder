package spotify

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Config struct {
	BaseURL string
}

type Client struct {
	config *Config
	token  string
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

func (c *Client) setBearer(req *http.Request) {
	bearer := "Bearer " + c.token
	req.Header.Set("Authorization", bearer)
}

func (c *Client) Get(responseModel interface{}, url string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	c.setBearer(req)

	resp, err := client.Do(req)
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

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	c.setBearer(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	err = serializeToStruct(resp.Body, responseModel)
	if err != nil {
		return err
	}

	return nil
}
