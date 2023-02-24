package wrapper


type tXT2IMG struct {
	Prompt                            string           `json:"prompt"`
	Steps                             int64            `json:"steps"`
	Width                             int64            `json:"width"`
	Height                            int64            `json:"height"`
}

func getDefaultDataTXT2IMGReq()tXT2IMG{
	t:=tXT2IMG{
		Prompt: "masterpiece, best quality,Amazing,finely detail,Depth of field,extremely detailed CG unity 8k wallpaper,",
		Steps:50,
		Width:512,
		Height: 512,
	}
	return t
}

// TXT2IMG  Resp
type tXT2IMGResp struct {
	Images     []string   `json:"images"`
	Info       string     `json:"info"`
}
