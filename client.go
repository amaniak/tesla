package tesla

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	mockUri    = "https://private-836c6-timdorr.apiary-mock.com"
	prodUri    = "https://owner-api.teslamotors.com"
	httpClient *http.Client
	userAgent  = "Model S 2.1.79 (Nexus 5; Android REL 4.4.4; en_US)"
)

//
// Client for connecting to Tesla Customer Portal
//

type Client struct {

	// configuration
	client  *http.Client
	BaseURL *url.URL
	token   *Token

	// credentials
	credentials *Credentials

	// services
	Vehicle *VehicleService
}

// Request builder
func (c *Client) NewRequest(method, urlStr string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(params.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	if userAgent != "" {
		req.Header.Add("User-Agent", userAgent)
	}
	return req, nil
}

// Authorize request
func (c *Client) Authorize(cred *Credentials) *Token {

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Add("client_id", cred.ClientId)
	data.Add("client_secret", cred.ClientSecret)
	data.Add("email", cred.Username)
	data.Add("password", cred.Password)

	// jsonBody := []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, _ := c.NewRequest("POST", "/oauth/token", data)

	tok := new(Token)

	resp, _ := c.client.Do(req)
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(respData, &tok)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	// store for later
	Ok("[OK] Tesla Portal token recieved")
	c.token = tok
	return tok
}

// Connect to Tesla Customer Portal
func Connect() *Client {
	httpClient = http.DefaultClient
	baseURL, _ := url.Parse(prodUri)
	c := &Client{httpClient, baseURL, nil, nil, nil}
	c.credentials = new(Credentials)
	c.Vehicle = &VehicleService{client: c}
	return c
}
