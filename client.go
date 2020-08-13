package invoice

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	AppId          string
	ApiKey         string
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
	// fmt.Println(string(data), dataMap)

	isExist, ok := dataMap["isExist"].(string)
	if !ok {
		return false, fmt.Errorf("isExist field missing")
	}

	return isExist == "Y", nil

}

func (c *Client) CheckBusinessAdministrationNumberExist(ban string) (bool, error) {

	timestamp := fmt.Sprintf("%d", time.Now().Unix()+1000)

	params := url.Values{}
	params.Add("version", "1.0")
	params.Add("serial", "0000000001")
	params.Add("action", "qryBanUnitTp")
	params.Add("ban", ban)
	params.Add("timeStamp", timestamp)
	params.Add("appId", c.AppId)
	requestBody := params.Encode()

	// fmt.Println("appId", c.AppId)
	// fmt.Println("param", requestBody)

	// TODO: make sure the order is alphabet order
	params.Add("signature", c.hmac(requestBody))
	requestBody = params.Encode()

	host := HOST_VC_EINVOICE
	data, err := c.doRquest(host, requestBody)
	if err != nil {
		return false, err
	}

	dataMap := map[string]interface{}{}
	json.Unmarshal(data, &dataMap)
	// fmt.Println(string(data), dataMap)

	isExist, ok := dataMap["banUnitTpStatus"].(string)
	if !ok {
		return false, fmt.Errorf("banUnitTpStatus field missing")
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

func (c *Client) hmac(message string) string {
	key := c.ApiKey
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	// fmt.Println(expectedMAC)
	return base64.StdEncoding.EncodeToString(expectedMAC)
}
