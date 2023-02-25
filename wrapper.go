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
func (c *WrapperClient) GetPromptStyles() (PromptStyles, error) {
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

/*
Get Realesrgan Models
*/
func (c *WrapperClient) GetRealesrganModels() (RealesrganModels, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/realesrgan-models")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	models := RealesrganModels{}
	err = json.Unmarshal(body, &models)
	return models, err
}

/*
Get Face Restorers
*/
func (c *WrapperClient) GetFaceRestorers() (FaceRestorers, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/face-restorers")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	restorers := FaceRestorers{}
	err = json.Unmarshal(body, &restorers)
	return restorers, err
}

/*
Get Embeddings
!Not stable
*/
func (c *WrapperClient) GetEmbeddings() (Embeddings, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/embeddings")
	if err != nil {
		return Embeddings{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	embeddings := Embeddings{}
	err = json.Unmarshal(body, &embeddings)
	return embeddings, err
}

/*
Get Hypernetworks
*/
func (c *WrapperClient) GetHypernetworks() (Hypernetworks, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/hypernetworks")
	if err != nil {
		return Hypernetworks{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	hypernetworks := Hypernetworks{}
	err = json.Unmarshal(body, &hypernetworks)
	return hypernetworks, err
}

/*
Get Upscalers
*/
func (c *WrapperClient) GetUpscalers() (Upscalers, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/upscalers")
	if err != nil {
		return Upscalers{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	upscalers := Upscalers{}
	err = json.Unmarshal(body, &upscalers)
	return upscalers, err
}

/*
Get Cmd Flags
*/
func (c *WrapperClient) GetCmdFlags() (CmdFlags, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/sdapi/v1/cmd-flags")
	if err != nil {
		return CmdFlags{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	cmdFlags := CmdFlags{}
	err = json.Unmarshal(body, &cmdFlags)
	return cmdFlags, err
}

// Get Current User
func (c *WrapperClient) GetCurrentUser() (string, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/user")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// Login Check
func (c *WrapperClient) LoginCheck() (string, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/login_check")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// Get Token
func (c *WrapperClient) GetToken() (Token, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/token")
	if err != nil {
		return Token{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	token := Token{}
	err = json.Unmarshal(body, &token)
	return token, err
}

// App Id
func (c *WrapperClient) AppId() (AppId, error) {
	resp, err := c.Client.Get(c.ApiUrl + "/app_id")
	if err != nil {
		return AppId{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	appId := AppId{}
	err = json.Unmarshal(body, &appId)
	return appId, err
}

// Reset Iterator
func (c *WrapperClient) Reset(req ResetReq)(ResetResp,error){
	b, err := json.Marshal(req)
	if err != nil {
		return ResetResp{}, err
	}
	resp, err := c.Client.Post(c.ApiUrl+"/reset", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return ResetResp{}, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := ResetResp{}
	err = json.Unmarshal(body, &res)
	return res, err
}

//!!! Not implemented yet.

func (c *WrapperClient) FetchFile() {
	panic("not implemented")
}

func (c *WrapperClient) TrainHypernetwork() {
	panic("not implemented")
}

func (c *WrapperClient) TrainEmbedding() {
	panic("not implemented")
}

func (c *WrapperClient) Preprocess() {
	panic("not implemented")
}

func (c *WrapperClient) CreateHypernetwork() {
	panic("not implemented")
}

func (c *WrapperClient) CreateEmbedding() {
	panic("not implemented")
}

func (c *WrapperClient) RefreshCheckpoints() {
	panic("not implemented")
}

func (c *WrapperClient) GetSamples() {
	panic("not implemented")
}

func (c *WrapperClient) SetConfig() {
	panic("not implemented")
}

func (c *WrapperClient) GetConifg() {
	panic("not implemented")
}

func (c *WrapperClient) Skip() {
	panic("not implemented")
}

func (c *WrapperClient) Interrupt() {
	panic("not implemented")
}

func (c *WrapperClient) Interrogate() {
	panic("not implemented")
}

func (c *WrapperClient) Progress() {
	panic("not implemented")
}

func (c *WrapperClient) PngInfo() {
	panic("not implemented")
}

func (c *WrapperClient) ExtrasBatchImages() {
	panic("not implemented")
}

func (c *WrapperClient) ExtrasSingleImage() {
	panic("not implemented")
}

func (c *WrapperClient) Img2Img() {
	panic("not implemented")
}

func (c *WrapperClient) RobotsTxt() {
	panic("not implemented")
}

func (c *WrapperClient) StartupEvents() {
	panic("not implemented")
}

func (c *WrapperClient) GetQueueStatus() {
	panic("not implemented")
}
