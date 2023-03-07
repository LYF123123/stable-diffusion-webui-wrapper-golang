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
	NegativePrompt string `json:"negative_prompt"`
	Steps  int64  `json:"steps"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	SamplerName string `json:"sampler_name"`
}

func getDefaultDataTXT2IMGReq() TXT2IMGReq {
	t := TXT2IMGReq{
		Prompt: "masterpiece, best quality,Amazing,finely detail,Depth of field,extremely detailed CG unity 8k wallpaper,",
		NegativePrompt: "(worst quality:1.25), (low quality:1.25), (lowres:1.1), (monochrome:1.1), (greyscale), multiple views, comic, sketch, (blurry:1.05),",
		Steps:  50,
		Width:  512,
		Height: 512,
		SamplerName: "DPM++ SDE Karras",
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

// Set Stable Diffusion Checkpoint
type CheckpointReq struct{
	SdModelCheckpoint string `json:"sd_model_checkpoint"`
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

// Config Resp
type ConfigResp struct {
    SamplesSave                        bool          `json:"samples_save"`                         
    SamplesFormat                      string        `json:"samples_format"`                       
    SamplesFilenamePattern             string        `json:"samples_filename_pattern"`             
    SaveImagesAddNumber                bool          `json:"save_images_add_number"`               
    GridSave                           bool          `json:"grid_save"`                            
    GridFormat                         string        `json:"grid_format"`                          
    GridExtendedFilename               bool          `json:"grid_extended_filename"`               
    GridOnlyIfMultiple                 bool          `json:"grid_only_if_multiple"`                
    GridPreventEmptySpots              bool          `json:"grid_prevent_empty_spots"`             
    NRows                              float64         `json:"n_rows"`                               
    EnablePnginfo                      bool          `json:"enable_pnginfo"`                       
    SaveTxt                            bool          `json:"save_txt"`                             
    SaveImagesBeforeFaceRestoration    bool          `json:"save_images_before_face_restoration"`  
    SaveImagesBeforeHighresFix         bool          `json:"save_images_before_highres_fix"`       
    SaveImagesBeforeColorCorrection    bool          `json:"save_images_before_color_correction"`  
    JPEGQuality                        float64         `json:"jpeg_quality"`                         
    ExportFor4Chan                     bool          `json:"export_for_4chan"`                     
    ImgDownscaleThreshold              float64         `json:"img_downscale_threshold"`              
    TargetSideLength                   float64         `json:"target_side_length"`                   
    UseOriginalNameBatch               bool          `json:"use_original_name_batch"`              
    UseUpscalerNameAsSuffix            bool          `json:"use_upscaler_name_as_suffix"`          
    SaveSelectedOnly                   bool          `json:"save_selected_only"`                   
    DoNotAddWatermark                  bool          `json:"do_not_add_watermark"`                 
    TempDir                            string        `json:"temp_dir"`                             
    CleanTempDirAtStart                bool          `json:"clean_temp_dir_at_start"`              
    OutdirSamples                      string        `json:"outdir_samples"`                       
    OutdirTxt2ImgSamples               string        `json:"outdir_txt2img_samples"`               
    OutdirImg2ImgSamples               string        `json:"outdir_img2img_samples"`               
    OutdirExtrasSamples                string        `json:"outdir_extras_samples"`                
    OutdirGrids                        string        `json:"outdir_grids"`                         
    OutdirTxt2ImgGrids                 string        `json:"outdir_txt2img_grids"`                 
    OutdirImg2ImgGrids                 string        `json:"outdir_img2img_grids"`                 
    OutdirSave                         string        `json:"outdir_save"`                          
    SaveToDirs                         bool          `json:"save_to_dirs"`                         
    GridSaveToDirs                     bool          `json:"grid_save_to_dirs"`                    
    UseSaveToDirsForUI                 bool          `json:"use_save_to_dirs_for_ui"`              
    DirectoriesFilenamePattern         string        `json:"directories_filename_pattern"`         
    DirectoriesMaxPromptWords          float64         `json:"directories_max_prompt_words"`         
    ESRGANTile                         float64         `json:"ESRGAN_tile"`                          
    ESRGANTileOverlap                  float64         `json:"ESRGAN_tile_overlap"`                  
    RealesrganEnabledModels            []string      `json:"realesrgan_enabled_models"`            
    UpscalerForImg2Img                 interface{}   `json:"upscaler_for_img2img"`                 
    LdsrSteps                          float64         `json:"ldsr_steps"`                           
    LdsrCached                         bool          `json:"ldsr_cached"`                          
    SWINTile                           float64         `json:"SWIN_tile"`                            
    SWINTileOverlap                    float64         `json:"SWIN_tile_overlap"`                    
    FaceRestorationModel               string        `json:"face_restoration_model"`               
    CodeFormerWeight                   float64       `json:"code_former_weight"`                   
    FaceRestorationUnload              bool          `json:"face_restoration_unload"`              
    ShowWarnings                       bool          `json:"show_warnings"`                        
    MemmonPollRate                     float64         `json:"memmon_poll_rate"`                     
    SamplesLogStdout                   bool          `json:"samples_log_stdout"`                   
    MultipleTqdm                       bool          `json:"multiple_tqdm"`                        
    PrintHypernetExtra                 bool          `json:"print_hypernet_extra"`                 
    UnloadModelsWhenTraining           bool          `json:"unload_models_when_training"`          
    PinMemory                          bool          `json:"pin_memory"`                           
    SaveOptimizerState                 bool          `json:"save_optimizer_state"`                 
    SaveTrainingSettingsToTxt          bool          `json:"save_training_settings_to_txt"`        
    DatasetFilenameWordRegex           string        `json:"dataset_filename_word_regex"`          
    DatasetFilenameJoinString          string        `json:"dataset_filename_join_string"`         
    TrainingImageRepeatsPerEpoch       float64         `json:"training_image_repeats_per_epoch"`     
    TrainingWriteCSVEvery              float64         `json:"training_write_csv_every"`             
    TrainingXattentionOptimizations    bool          `json:"training_xattention_optimizations"`    
    TrainingEnableTensorboard          bool          `json:"training_enable_tensorboard"`          
    TrainingTensorboardSaveImages      bool          `json:"training_tensorboard_save_images"`     
    TrainingTensorboardFlushEvery      float64         `json:"training_tensorboard_flush_every"`     
    SDModelCheckpoint                  string        `json:"sd_model_checkpoint"`                  
    SDCheckpointCache                  float64         `json:"sd_checkpoint_cache"`                  
    SDVaeCheckpointCache               float64         `json:"sd_vae_checkpoint_cache"`              
    SDVae                              string        `json:"sd_vae"`                               
    SDVaeAsDefault                     bool          `json:"sd_vae_as_default"`                    
    InpaintingMaskWeight               float64         `json:"inpainting_mask_weight"`               
    InitialNoiseMultiplier             float64         `json:"initial_noise_multiplier"`             
    Img2ImgColorCorrection             bool          `json:"img2img_color_correction"`             
    Img2ImgFixSteps                    bool          `json:"img2img_fix_steps"`                    
    Img2ImgBackgroundColor             string        `json:"img2img_background_color"`             
    EnableQuantization                 bool          `json:"enable_quantization"`                  
    EnableEmphasis                     bool          `json:"enable_emphasis"`                      
    EnableBatchSeeds                   bool          `json:"enable_batch_seeds"`                   
    CommaPaddingBacktrack              float64         `json:"comma_padding_backtrack"`              
    CLIPStopAtLastLayers               float64         `json:"CLIP_stop_at_last_layers"`             
    UpcastAttn                         bool          `json:"upcast_attn"`                          
    UseOldEmphasisImplementation       bool          `json:"use_old_emphasis_implementation"`      
    UseOldKarrasSchedulerSigmas        bool          `json:"use_old_karras_scheduler_sigmas"`      
    NoDpmppSdeBatchDeterminism         bool          `json:"no_dpmpp_sde_batch_determinism"`       
    UseOldHiresFixWidthHeight          bool          `json:"use_old_hires_fix_width_height"`       
    InterrogateKeepModelsInMemory      bool          `json:"interrogate_keep_models_in_memory"`    
    InterrogateReturnRanks             bool          `json:"interrogate_return_ranks"`             
    InterrogateClipNumBeams            float64         `json:"interrogate_clip_num_beams"`           
    InterrogateClipMinLength           float64         `json:"interrogate_clip_min_length"`          
    InterrogateClipMaxLength           float64         `json:"interrogate_clip_max_length"`          
    InterrogateClipDictLimit           float64         `json:"interrogate_clip_dict_limit"`          
    InterrogateClipSkipCategories      []interface{} `json:"interrogate_clip_skip_categories"`     
    InterrogateDeepbooruScoreThreshold float64       `json:"interrogate_deepbooru_score_threshold"`
    DeepbooruSortAlpha                 bool          `json:"deepbooru_sort_alpha"`                 
    DeepbooruUseSpaces                 bool          `json:"deepbooru_use_spaces"`                 
    DeepbooruEscape                    bool          `json:"deepbooru_escape"`                     
    DeepbooruFilterTags                string        `json:"deepbooru_filter_tags"`                
    ExtraNetworksDefaultView           string        `json:"extra_networks_default_view"`          
    ExtraNetworksDefaultMultiplier     float64         `json:"extra_networks_default_multiplier"`    
    SDHypernetwork                     string        `json:"sd_hypernetwork"`                      
    SDLora                             string        `json:"sd_lora"`                              
    LoraApplyToOutputs                 bool          `json:"lora_apply_to_outputs"`                
    ReturnGrid                         bool          `json:"return_grid"`                          
    DoNotShowImages                    bool          `json:"do_not_show_images"`                   
    AddModelHashToInfo                 bool          `json:"add_model_hash_to_info"`               
    AddModelNameToInfo                 bool          `json:"add_model_name_to_info"`               
    DisableWeightsAutoSwap             bool          `json:"disable_weights_auto_swap"`            
    SendSeed                           bool          `json:"send_seed"`                            
    SendSize                           bool          `json:"send_size"`                            
    Font                               string        `json:"font"`                                 
    JSModalLightbox                    bool          `json:"js_modal_lightbox"`                    
    JSModalLightboxInitiallyZoomed     bool          `json:"js_modal_lightbox_initially_zoomed"`   
    ShowProgressInTitle                bool          `json:"show_progress_in_title"`               
    SamplersInDropdown                 bool          `json:"samplers_in_dropdown"`                 
    DimensionsAndBatchTogether         bool          `json:"dimensions_and_batch_together"`        
    KeyeditPrecisionAttention          float64       `json:"keyedit_precision_attention"`          
    KeyeditPrecisionExtra              float64       `json:"keyedit_precision_extra"`              
    Quicksettings                      string        `json:"quicksettings"`                        
    UIReorder                          string        `json:"ui_reorder"`                           
    UIExtraNetworksTabReorder          string        `json:"ui_extra_networks_tab_reorder"`        
    Localization                       string        `json:"localization"`                         
    ShowProgressbar                    bool          `json:"show_progressbar"`                     
    LivePreviewsEnable                 bool          `json:"live_previews_enable"`                 
    ShowProgressGrid                   bool          `json:"show_progress_grid"`                   
    ShowProgressEveryNSteps            float64         `json:"show_progress_every_n_steps"`          
    ShowProgressType                   string        `json:"show_progress_type"`                   
    LivePreviewContent                 string        `json:"live_preview_content"`                 
    LivePreviewRefreshPeriod           float64         `json:"live_preview_refresh_period"`          
    HideSamplers                       []interface{} `json:"hide_samplers"`                        
    EtaDdim                            float64         `json:"eta_ddim"`                             
    EtaAncestral                       float64         `json:"eta_ancestral"`                        
    DdimDiscretize                     string        `json:"ddim_discretize"`                      
    SChurn                             float64         `json:"s_churn"`                              
    STmin                              float64         `json:"s_tmin"`                               
    SNoise                             float64         `json:"s_noise"`                              
    EtaNoiseSeedDelta                  float64         `json:"eta_noise_seed_delta"`                 
    AlwaysDiscardNextToLastSigma       bool          `json:"always_discard_next_to_last_sigma"`    
    PostprocessingEnableInMainUI       []interface{} `json:"postprocessing_enable_in_main_ui"`     
    PostprocessingOperationOrder       []interface{} `json:"postprocessing_operation_order"`       
    UpscalingMaxImagesInCache          float64         `json:"upscaling_max_images_in_cache"`        
    DisabledExtensions                 []interface{} `json:"disabled_extensions"`                  
    SDCheckpointHash                   string        `json:"sd_checkpoint_hash"`                   
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
