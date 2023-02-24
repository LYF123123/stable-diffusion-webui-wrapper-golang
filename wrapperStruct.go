package wrapper

// TXT2IMG Req

type TXT2IMGReq struct {
	Prompt string `json:"prompt"`
	Steps  int64  `json:"steps"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

func getDefaultDataTXT2IMGReq() TXT2IMGReq {
	t := TXT2IMGReq{
		Prompt: "masterpiece, best quality,Amazing,finely detail,Depth of field,extremely detailed CG unity 8k wallpaper,",
		Steps:  50,
		Width:  512,
		Height: 512,
	}
	return t
}

// TXT2IMG  Resp
type TXT2IMGResp struct {
	Images []string `json:"images"`
	Info   string   `json:"info"`
}

// Sd Models resp

type SDModels []SDModel

type SDModel struct {
	Title     string  `json:"title"`
	ModelName string  `json:"model_name"`
	Hash      *string `json:"hash"`
	Sha256    *string `json:"sha256"`
	Filename  string  `json:"filename"`
	Config    string  `json:"config"`
}

// Prompt Styles resp

type PromptStyles []PromptStyle

type PromptStyle struct {
	Name string `json:"name"`
	Prompt string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt"`
}

// Memory Resp

type MemStatus struct {
	RAM  RAM  `json:"ram"`
	Cuda Cuda `json:"cuda"`
}

type Cuda struct {
	System    RAM    `json:"system"`
	Active    Active `json:"active"`
	Allocated Active `json:"allocated"`
	Reserved  Active `json:"reserved"`
	Inactive  Active `json:"inactive"`
	Events    Events `json:"events"`
}

type Active struct {
	Current float64 `json:"current"`
	Peak    float64 `json:"peak"`
}

type Events struct {
	Retries float64 `json:"retries"`
	OOM     float64 `json:"oom"`
}

type RAM struct {
	Free  float64 `json:"free"`
	Used  float64 `json:"used"`
	Total float64 `json:"total"`
}
