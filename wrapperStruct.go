package wrapper

// Token Resp
type Token struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

// App Id Resp
type AppId struct {
	AppId uint64 `json:"app_id"`
}

// Reset Iterator Req
type ResetReq struct {
	SessionHash string `json:"session_hash"`
	FnIndex     uint64 `json:"fn_index"`
}

// Reset Iterator Resp

type ResetResp struct {
	Success bool `json:"success"`
}

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

// TXT2IMG Resp
type TXT2IMGResp struct {
	Images []string `json:"images"`
	Info   string   `json:"info"`
}

// IMG2IMG Req
type IMG2IMGReq struct {
	Prompt string `json:"prompt"`
	// Steps  int64  `json:"steps"`	// Not Work
	InitImages []string `json:"init_images"` // should be base64 encoded
}

// IMG2IMG Resp
type IMG2IMGResp struct {
	Images []string `json:"images"`
	Info   string   `json:"info"`
}

// Extras Single Image Req
type ExtrasSingleImageReq struct {
	ResizeMode                 int64  `json:"resize_mode"`
	ShowExtrasResults          bool   `json:"show_extras_results"`
	GfpganVisibility           int64  `json:"gfpgan_visibility"`
	CodeformerVisibility       int64  `json:"codeformer_visibility"`
	CodeformerWeight           int64  `json:"codeformer_weight"`
	UpscalingResize            int64  `json:"upscaling_resize"`
	UpscalingResizeW           int64  `json:"upscaling_resize_w"`
	UpscalingResizeH           int64  `json:"upscaling_resize_h"`
	UpscalingCrop              bool   `json:"upscaling_crop"`
	Upscaler1                  string `json:"upscaler_1"`
	Upscaler2                  string `json:"upscaler_2"`
	ExtrasUpscaler2_Visibility int64  `json:"extras_upscaler_2_visibility"`
	UpscaleFirst               bool   `json:"upscale_first"`
	Image                      string `json:"image"`
}

// Extras Single Image Resp
type ExtrasSingleImageResp struct {
	HtmlInfo string `json:"html_info"`
	Image    string `json:"image"`
}

// Cmd Flags Resp
type CmdFlags struct {
	DataDir                       string        `json:"data_dir"`
	Config                        string        `json:"config"`
	Ckpt                          string        `json:"ckpt"`
	CkptDir                       string        `json:"ckpt_dir"`
	VaeDir                        string        `json:"vae_dir"`
	GfpganDir                     string        `json:"gfpgan_dir"`
	GfpganModel                   string        `json:"gfpgan_model"`
	NoHalf                        bool          `json:"no_half"`
	NoHalfVae                     bool          `json:"no_half_vae"`
	NoProgressbarHiding           bool          `json:"no_progressbar_hiding"`
	MaxBatchCount                 int64         `json:"max_batch_count"`
	EmbeddingsDir                 string        `json:"embeddings_dir"`
	TextualInversionTemplatesDir  string        `json:"textual_inversion_templates_dir"`
	HypernetworkDir               string        `json:"hypernetwork_dir"`
	LocalizationsDir              string        `json:"localizations_dir"`
	AllowCode                     bool          `json:"allow_code"`
	Medvram                       bool          `json:"medvram"`
	Lowvram                       bool          `json:"lowvram"`
	Lowram                        bool          `json:"lowram"`
	AlwaysBatchCondUncond         bool          `json:"always_batch_cond_uncond"`
	UnloadGfpgan                  bool          `json:"unload_gfpgan"`
	Precision                     string        `json:"precision"`
	UpcastSampling                bool          `json:"upcast_sampling"`
	Share                         bool          `json:"share"`
	Ngrok                         string        `json:"ngrok"`
	NgrokRegion                   string        `json:"ngrok_region"`
	EnableInsecureExtensionAccess bool          `json:"enable_insecure_extension_access"`
	CodeformerModelsPath          string        `json:"codeformer_models_path"`
	GfpganModelsPath              string        `json:"gfpgan_models_path"`
	EsrganModelsPath              string        `json:"esrgan_models_path"`
	BsrganModelsPath              string        `json:"bsrgan_models_path"`
	RealesrganModelsPath          string        `json:"realesrgan_models_path"`
	ClipModelsPath                string        `json:"clip_models_path"`
	Xformers                      bool          `json:"xformers"`
	ForceEnableXformers           bool          `json:"force_enable_xformers"`
	XformersFlashAttention        bool          `json:"xformers_flash_attention"`
	Deepdanbooru                  bool          `json:"deepdanbooru"`
	OptSplitAttention             bool          `json:"opt_split_attention"`
	OptSubQuadAttention           bool          `json:"opt_sub_quad_attention"`
	SubQuadQChunkSize             int64         `json:"sub_quad_q_chunk_size"`
	SubQuadKvChunkSize            string        `json:"sub_quad_kv_chunk_size"`
	SubQuadChunkThreshold         string        `json:"sub_quad_chunk_threshold"`
	OptSplitAttentionInvokeai     bool          `json:"opt_split_attention_invokeai"`
	OptSplitAttentionV1           bool          `json:"opt_split_attention_v1"`
	DisableOptSplitAttention      bool          `json:"disable_opt_split_attention"`
	DisableNanCheck               bool          `json:"disable_nan_check"`
	UseCPU                        []interface{} `json:"use_cpu"`
	Listen                        bool          `json:"listen"`
	Port                          string        `json:"port"`
	ShowNegativePrompt            bool          `json:"show_negative_prompt"`
	UIConfigFile                  string        `json:"ui_config_file"`
	HideUIDirConfig               bool          `json:"hide_ui_dir_config"`
	FreezeSettings                bool          `json:"freeze_settings"`
	UISettingsFile                string        `json:"ui_settings_file"`
	GradioDebug                   bool          `json:"gradio_debug"`
	GradioAuth                    string        `json:"gradio_auth"`
	GradioAuthPath                string        `json:"gradio_auth_path"`
	GradioImg2ImgTool             string        `json:"gradio_img2img_tool"`
	GradioInpaintTool             string        `json:"gradio_inpaint_tool"`
	OptChannelslast               bool          `json:"opt_channelslast"`
	StylesFile                    string        `json:"styles_file"`
	Autolaunch                    bool          `json:"autolaunch"`
	Theme                         string        `json:"theme"`
	UseTextboxSeed                bool          `json:"use_textbox_seed"`
	DisableConsoleProgressbars    bool          `json:"disable_console_progressbars"`
	EnableConsolePrompts          bool          `json:"enable_console_prompts"`
	VaePath                       string        `json:"vae_path"`
	DisableSafeUnpickle           bool          `json:"disable_safe_unpickle"`
	API                           bool          `json:"api"`
	APIAuth                       string        `json:"api_auth"`
	APILog                        bool          `json:"api_log"`
	Nowebui                       bool          `json:"nowebui"`
	UIDebugMode                   bool          `json:"ui_debug_mode"`
	DeviceID                      string        `json:"device_id"`
	Administrator                 bool          `json:"administrator"`
	CorsAllowOrigins              string        `json:"cors_allow_origins"`
	CorsAllowOriginsRegex         string        `json:"cors_allow_origins_regex"`
	TLSKeyfile                    string        `json:"tls_keyfile"`
	TLSCertfile                   string        `json:"tls_certfile"`
	ServerName                    string        `json:"server_name"`
	GradioQueue                   bool          `json:"gradio_queue"`
	SkipVersionCheck              bool          `json:"skip_version_check"`
	NoHashing                     bool          `json:"no_hashing"`
	NoDownloadSDModel             bool          `json:"no_download_sd_model"`
	LdsrModelsPath                string        `json:"ldsr_models_path"`
	LoraDir                       string        `json:"lora_dir"`
	ScunetModelsPath              string        `json:"scunet_models_path"`
	SwinirModelsPath              string        `json:"swinir_models_path"`
}

// Upscalers Resp
type Upscalers []Upscaler
type Upscaler struct {
	Name      string  `json:"name"`
	ModelName string  `json:"model_name"`
	ModelPath string  `json:"model_path"`
	ModelUrl  string  `json:"model_url"`
	Scale     float64 `json:"scale"`
}

// Sd Models resp
type SDModels []SDModel

type SDModel struct {
	Title     string `json:"title"`
	ModelName string `json:"model_name"`
	Hash      string `json:"hash"`
	Sha256    string `json:"sha256"`
	Filename  string `json:"filename"`
	Config    string `json:"config"`
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
	Name   string `json:"name"`
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
