package myip

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var ipAddressServiceHost = "https://api.myip.com/"

type result struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	CC      string `json:"cc"`
}

func GetIPAddress() (string, error) {
	response, err := http.Get(ipAddressServiceHost)

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	res := result{}
	json.Unmarshal(data, &res)

	return res.IP, nil
}
