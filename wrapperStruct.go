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
	Hash      string `json:"hash"`
	Sha256    string `json:"sha256"`
	Filename  string  `json:"filename"`
	Config    string  `json:"config"`
}

//Hypernetworks Resp
type Hypernetworks []Hypernetwork
type Hypernetwork struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

//Face Restorers resp
type FaceRestorers []FaceRestorer

type FaceRestorer struct {
	Name   string  `json:"name"`
	CmdDir string `json:"cmd_dir"`
}

// Realesrgan Models resp
type RealesrganModels []RealesrganModel

type RealesrganModel struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Scale int64  `json:"scale"`
}

// Prompt Styles resp

type PromptStyles []PromptStyle

type PromptStyle struct {
	Name           string `json:"name"`
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt"`
}

// Embeddings Resp
type Embeddings struct {
	Loaded  Loaded `json:"loaded"`
	Skipped Loaded `json:"skipped"`
}

type Loaded struct {
	AdditionalProp1 AdditionalProp `json:"additionalProp1"`
	AdditionalProp2 AdditionalProp `json:"additionalProp2"`
	AdditionalProp3 AdditionalProp `json:"additionalProp3"`
}

type AdditionalProp struct {
	Step             int64  `json:"step"`
	SDCheckpoint     string `json:"sd_checkpoint"`
	SDCheckpointName string `json:"sd_checkpoint_name"`
	Shape            int64  `json:"shape"`
	Vectors          int64  `json:"vectors"`
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
