package tesla

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

//
type Vehicle struct {
	Id *int `json:"id,omitempty"`
}

//
type VehicleService struct {
	client *Client
}

func (service *VehicleService) AccessToken() string {
	return *service.client.token.AccessToken
}

func (service *VehicleService) HonkHonk(id int) string {

	uri := fmt.Sprintf("/api/1/vehicles/%d/command/honk_horn", id)
	data := url.Values{}
	data.Set("vehicle_id", string(id))

	req, _ := service.client.NewRequest("POST", uri, data)

	access_token := service.AccessToken()
	req.Header.Add("Authorization", "Bearer "+access_token)

	//var tok Vehicle

	resp, _ := service.client.client.Do(req)
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)

	// Return
	return string(respData)
}

func (service *VehicleService) All() string {
	data := url.Values{}
	req, _ := service.client.NewRequest("GET", "api/1/vehicles", data)

	access_token := service.AccessToken()
	req.Header.Add("Authorization", "Bearer "+access_token)

	resp, _ := service.client.client.Do(req)
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)

	// Return
	return string(respData)
}
