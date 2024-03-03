package rest

import (
	"encoding/json"
	"fmt"
	"github.com/leryn1122/terraform-provider-vmware-fusion/pkg/support"
	"gopkg.in/resty.v1"
	"net/http"
	"net/url"
)

const (
	DefaultHost = "127.0.0.1"
	DefaultPort = 8697
)

const (
	VMWareRestContentType      string = "application/vnd.vmware.vmw.rest-v1+json"
	ApplicationJSONContentType string = "application/json"
)

// Client to invoke VMWare Fusion RESTful API.
type Client struct {
	Client   *resty.Client
	BaseURL  *url.URL
	Username string
	Password string
	Insecure bool
}

func NewClient(username, password, baseUrl string) (*Client, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Client:   resty.New(),
		BaseURL:  u,
		Username: username,
		Password: password,
		Insecure: true,
	}
	return client, nil
}

func (c *Client) request(url string, method string, payload []byte, params ...map[string]interface{}) ([]byte, VmwareFusionError) {
	if support.Contains([]string{http.MethodGet, http.MethodPost, http.MethodPut}, method) {
		c.Client.Header.Add("Content-Type", VMWareRestContentType)
	} else if http.MethodDelete == method {
		// Do NOT set `Content-Type`
	} else {
		c.Client.Header.Add("Content-Type", ApplicationJSONContentType)
	}

	c.Client.SetBasicAuth(c.Username, c.Password)

	request := c.Client.NewRequest()
	request.SetBody(payload)
	url = support.FormatString(url, params)
	response, err := request.Execute(method, fmt.Sprintf("%s/api%s", c.BaseURL, url))
	if err != nil {
		return nil, FromRawError(err)
	}

	var vmErr VmwareFusionError
	status := response.StatusCode()
	if support.Contains([]int{http.StatusOK, http.StatusCreated, http.StatusNoContent}, status) {

	}
	return response.Body(), vmErr
}

type Params = map[string]interface{}

func SendRequest[T any, R any](client *Client, method string, url string, request T, params Params) (*R, *VmwareFusionError) {
	var vmErr VmwareFusionError

	payload, err := json.Marshal(request)
	if err != nil {
		vmErr = FromRawError(err)
		return nil, &vmErr
	}

	response, vmErr := client.request(url, method, payload, params)
	if response == nil {
		return nil, &vmErr
	}

	var resp R
	err = json.Unmarshal(response, resp)
	if err != nil {
		vmErr = FromRawError(err)
		return nil, &vmErr
	}

	return &resp, nil
}
