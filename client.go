package invoice

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	AppId          string
	ConnectionHost string
}

func NewClient(appId string) *Client {
	return &Client{
		AppId: appId,
	}

}

func (c *Client) CheckMobilePhoneExist(code string) (bool, error) {

	params := url.Values{}
	params.Add("version", "1.0")
	params.Add("action", "bcv")
	params.Add("barCode", code)
	params.Add("TxID", "1.0")
	params.Add("appId", c.AppId)
	requestBody := params.Encode()

	host := HOST_VC_EINVOICE
	data, err := c.doRquest(host, requestBody)
	if err != nil {
		return false, err
	}

	dataMap := map[string]interface{}{}
	json.Unmarshal(data, &dataMap)
	// fmt.Println(string(data), dataMap)

	isExist, ok := dataMap["isExist"].(string)
	if !ok {
		return false, fmt.Errorf("isExist field missing")
	}

	return isExist == "Y", nil

}

func (c *Client) CheckLoveCodeExist(code string) (bool, error) {

	params := url.Values{}
	params.Add("version", "1.0")
	params.Add("action", "preserveCodeCheck")
	params.Add("pCode", code)
	params.Add("TxID", "1.0")
	params.Add("appId", c.AppId)
	requestBody := params.Encode()

	host := HOST_VC_EINVOICE
	data, err := c.doRquest(host, requestBody)
	if err != nil {
		return false, err
	}

	dataMap := map[string]interface{}{}
	json.Unmarshal(data, &dataMap)
	fmt.Println(string(data), dataMap)

	isExist, ok := dataMap["isExist"].(string)
	if !ok {
		return false, fmt.Errorf("isExist field missing")
	}

	return isExist == "Y", nil

}

func (c *Client) doRquest(host string, requestBody string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			// InsecureSkipVerify: true
			ServerName: host,
		},
	}
	client := &http.Client{Transport: tr}

	body := strings.NewReader(requestBody)

	hostname := host

	if c.ConnectionHost != "" {
		hostname = c.ConnectionHost

	}

	req, err := http.NewRequest("POST", "https://"+hostname+"/BIZAPIVAN/biz", body)
	if err != nil {
		// handle err
		return nil, fmt.Errorf("new request: %q", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client connection: %q", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, err

}
