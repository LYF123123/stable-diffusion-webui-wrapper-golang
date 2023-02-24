package wrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

/*
set the api endpoint url
eg. http://127.0.0.1:7860
*/
func (c *WrapperClient) SetAPIUrl(apiUrl string) error {
	_, err := url.ParseRequestURI(apiUrl)
	if err != nil { // the api url is wrong
		return err
	}
	c.ApiUrl = apiUrl
	return nil
}

/*
	Set Proxy
	eg. http://127.0.0.1:8080
*/

func (c *WrapperClient) SetProxy(proxyUrl string) error {
	proxy, err := url.ParseRequestURI(proxyUrl)
	if err != nil {
		return err
	}
	c.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	return nil
}

/*
	Set prompt, split with ",".
	return with base64encode image
*/

func (c *WrapperClient) Text2Img(prompt string) (string, error) {
	defaultPayload := getDefaultDataTXT2IMGReq()
	defaultPayload.Prompt = defaultPayload.Prompt + prompt
	b, err := json.Marshal(defaultPayload)
	if err != nil {
		return "", err
	}
	resp, err := c.Client.Post(c.ApiUrl+"/sdapi/v1/txt2img", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := TXT2IMGResp{}
	err = json.Unmarshal(body, &res)
	return res.Images[0], err
}

/*
	Get Memory Status
*/

func (c *WrapperClient) GetMemory() (MemStatus, error) {

	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/memory")
	if err != nil {
		return MemStatus{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	status := MemStatus{}
	err = json.Unmarshal(body, &status)
	return status, err
}

/*
	Get Sd Models
*/

func (c *WrapperClient) GetSdModels() (SDModels, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/sd-models")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	models := SDModels{}
	err = json.Unmarshal(body, &models)
	return models, err
}

/*
	Get Prompt Styles
*/
func (c *WrapperClient)GetPromptStyles()(PromptStyles,error){
	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/prompt-styles")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	styles := PromptStyles{}
	err = json.Unmarshal(body, &styles)
	return styles, err
}