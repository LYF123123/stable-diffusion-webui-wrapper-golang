package wrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type WrapperClient struct {
	Client *http.Client
	ApiUrl string
}

func NewWrapperClient() *WrapperClient {
	return &WrapperClient{
		Client: &http.Client{},
	}
}

func (c *WrapperClient) SetAPIUrl(apiUrl string) error {
	_, err := url.ParseRequestURI(apiUrl)
	if err != nil { // the api url is wrong
		return err
	}
	c.ApiUrl = apiUrl
	return nil
}

func (c *WrapperClient) SetProxy(proxyUrl string) error {
	proxy, err := url.ParseRequestURI(proxyUrl)
	if err != nil {
		return err
	}
	c.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	return nil
}

func (c *WrapperClient) Text2Imgapi(prompt string) (string, error) {
	defaultPayload := getDefaultDataTXT2IMGReq()
	defaultPayload.Prompt = defaultPayload.Prompt+prompt
	log.Println(defaultPayload)
	b, err := json.Marshal(defaultPayload)
	if err != nil {
		return "", err
	}
	log.Println(b)
	resp, err := c.Client.Post(c.ApiUrl+"/sdapi/v1/txt2img", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res:=tXT2IMGResp{}
	err=json.Unmarshal(body,&res)
	return resp.Proto,err
}
