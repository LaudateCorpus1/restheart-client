package restheart

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var configuration Configuration

func init() {
	configuration = Configuration{}
	configuration.Load()
}

// Client creates a simple struct with data
// needed when making a Call to RESTHeart API
type Client struct {
	ObjectType      string
	ObjectName      string
	RequestMethod   string
	RequestPayload  string
	ResponsePayload string
}

// Call makes a http request to the RESTHeart API
func (client *Client) Call() error {
	// build request URL
	requestURL := strings.Join(
		[]string{configuration.Endpoint, client.ObjectType, client.ObjectName}, "/")
	httpClient := &http.Client{}

	request, err := http.NewRequest(
		client.RequestMethod,
		requestURL,
		strings.NewReader(client.RequestPayload))

	if err != nil {
		return fmt.Errorf("Unable to instantiate new http request: %s", err)
	}

	request.Header.Set("Content-Type", "application/json")

	// config objects are readable without username and pass
	// this may change in the future to implement token auth,
	// so we'll leave this a bit free
	if client.ObjectType != "config" {
		request.SetBasicAuth(configuration.Username, configuration.Password)
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("Unable to perform http request: %s", err)
	}

	switch response.StatusCode {
	case 401:
		return errors.New("401 not authorized")
	case 404:
		return errors.New("404 not found")
	case 406:
		return errors.New("406 invalid request")
	case 500:
		return errors.New("500 internal server error")
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	response.Body.Close()
	client.ResponsePayload = buffer.String()
	return nil
}
