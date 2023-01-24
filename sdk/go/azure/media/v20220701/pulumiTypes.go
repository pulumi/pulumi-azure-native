


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AacAudio struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	Profile      *string `pulumi:"profile"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type AacAudioResponse struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	Profile      *string `pulumi:"profile"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type AbsoluteClipTime struct {
	OdataType string `pulumi:"odataType"`
	Time      string `pulumi:"time"`
}

type AbsoluteClipTimeResponse struct {
	OdataType string `pulumi:"odataType"`
	Time      string `pulumi:"time"`
}

type Audio struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type AudioAnalyzerPreset struct {
	AudioLanguage       *string           `pulumi:"audioLanguage"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
}

type AudioAnalyzerPresetResponse struct {
	AudioLanguage       *string           `pulumi:"audioLanguage"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
}

type AudioOverlay struct {
	AudioGainLevel  *float64 `pulumi:"audioGainLevel"`
	End             *string  `pulumi:"end"`
	FadeInDuration  *string  `pulumi:"fadeInDuration"`
	FadeOutDuration *string  `pulumi:"fadeOutDuration"`
	InputLabel      string   `pulumi:"inputLabel"`
	OdataType       string   `pulumi:"odataType"`
	Start           *string  `pulumi:"start"`
}

type AudioOverlayResponse struct {
	AudioGainLevel  *float64 `pulumi:"audioGainLevel"`
	End             *string  `pulumi:"end"`
	FadeInDuration  *string  `pulumi:"fadeInDuration"`
	FadeOutDuration *string  `pulumi:"fadeOutDuration"`
	InputLabel      string   `pulumi:"inputLabel"`
	OdataType       string   `pulumi:"odataType"`
	Start           *string  `pulumi:"start"`
}

type AudioResponse struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type AudioTrackDescriptor struct {
	ChannelMapping *string `pulumi:"channelMapping"`
	OdataType      string  `pulumi:"odataType"`
}

type AudioTrackDescriptorResponse struct {
	ChannelMapping *string `pulumi:"channelMapping"`
	OdataType      string  `pulumi:"odataType"`
}

type BuiltInStandardEncoderPreset struct {
	Configurations *PresetConfigurations `pulumi:"configurations"`
	OdataType      string                `pulumi:"odataType"`
	PresetName     string                `pulumi:"presetName"`
}

type BuiltInStandardEncoderPresetResponse struct {
	Configurations *PresetConfigurationsResponse `pulumi:"configurations"`
	OdataType      string                        `pulumi:"odataType"`
	PresetName     string                        `pulumi:"presetName"`
}

type CopyAudio struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}

type CopyAudioResponse struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}

type CopyVideo struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}

type CopyVideoResponse struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}

type DDAudio struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type DDAudioResponse struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type Deinterlace struct {
	Mode   *string `pulumi:"mode"`
	Parity *string `pulumi:"parity"`
}

type DeinterlaceResponse struct {
	Mode   *string `pulumi:"mode"`
	Parity *string `pulumi:"parity"`
}

type FaceDetectorPreset struct {
	BlurType            *string           `pulumi:"blurType"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
	Resolution          *string           `pulumi:"resolution"`
}

type FaceDetectorPresetResponse struct {
	BlurType            *string           `pulumi:"blurType"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
	Resolution          *string           `pulumi:"resolution"`
}

type Fade struct {
	Duration  string  `pulumi:"duration"`
	FadeColor string  `pulumi:"fadeColor"`
	Start     *string `pulumi:"start"`
}

type FadeResponse struct {
	Duration  string  `pulumi:"duration"`
	FadeColor string  `pulumi:"fadeColor"`
	Start     *string `pulumi:"start"`
}

type Filters struct {
	Crop        *Rectangle    `pulumi:"crop"`
	Deinterlace *Deinterlace  `pulumi:"deinterlace"`
	FadeIn      *Fade         `pulumi:"fadeIn"`
	FadeOut     *Fade         `pulumi:"fadeOut"`
	Overlays    []interface{} `pulumi:"overlays"`
	Rotation    *string       `pulumi:"rotation"`
}

type FiltersResponse struct {
	Crop        *RectangleResponse   `pulumi:"crop"`
	Deinterlace *DeinterlaceResponse `pulumi:"deinterlace"`
	FadeIn      *FadeResponse        `pulumi:"fadeIn"`
	FadeOut     *FadeResponse        `pulumi:"fadeOut"`
	Overlays    []interface{}        `pulumi:"overlays"`
	Rotation    *string              `pulumi:"rotation"`
}

type FromAllInputFile struct {
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}

type FromAllInputFileResponse struct {
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}

type FromEachInputFile struct {
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}

type FromEachInputFileResponse struct {
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}

type H264Layer struct {
	AdaptiveBFrame  *bool    `pulumi:"adaptiveBFrame"`
	BFrames         *int     `pulumi:"bFrames"`
	Bitrate         int      `pulumi:"bitrate"`
	BufferWindow    *string  `pulumi:"bufferWindow"`
	Crf             *float64 `pulumi:"crf"`
	EntropyMode     *string  `pulumi:"entropyMode"`
	FrameRate       *string  `pulumi:"frameRate"`
	Height          *string  `pulumi:"height"`
	Label           *string  `pulumi:"label"`
	Level           *string  `pulumi:"level"`
	MaxBitrate      *int     `pulumi:"maxBitrate"`
	Profile         *string  `pulumi:"profile"`
	ReferenceFrames *int     `pulumi:"referenceFrames"`
	Slices          *int     `pulumi:"slices"`
	Width           *string  `pulumi:"width"`
}

type H264LayerResponse struct {
	AdaptiveBFrame  *bool    `pulumi:"adaptiveBFrame"`
	BFrames         *int     `pulumi:"bFrames"`
	Bitrate         int      `pulumi:"bitrate"`
	BufferWindow    *string  `pulumi:"bufferWindow"`
	Crf             *float64 `pulumi:"crf"`
	EntropyMode     *string  `pulumi:"entropyMode"`
	FrameRate       *string  `pulumi:"frameRate"`
	Height          *string  `pulumi:"height"`
	Label           *string  `pulumi:"label"`
	Level           *string  `pulumi:"level"`
	MaxBitrate      *int     `pulumi:"maxBitrate"`
	Profile         *string  `pulumi:"profile"`
	ReferenceFrames *int     `pulumi:"referenceFrames"`
	Slices          *int     `pulumi:"slices"`
	Width           *string  `pulumi:"width"`
}

type H264Video struct {
	Complexity           *string     `pulumi:"complexity"`
	KeyFrameInterval     *string     `pulumi:"keyFrameInterval"`
	Label                *string     `pulumi:"label"`
	Layers               []H264Layer `pulumi:"layers"`
	OdataType            string      `pulumi:"odataType"`
	RateControlMode      *string     `pulumi:"rateControlMode"`
	SceneChangeDetection *bool       `pulumi:"sceneChangeDetection"`
	StretchMode          *string     `pulumi:"stretchMode"`
	SyncMode             *string     `pulumi:"syncMode"`
}

type H264VideoResponse struct {
	Complexity           *string             `pulumi:"complexity"`
	KeyFrameInterval     *string             `pulumi:"keyFrameInterval"`
	Label                *string             `pulumi:"label"`
	Layers               []H264LayerResponse `pulumi:"layers"`
	OdataType            string              `pulumi:"odataType"`
	RateControlMode      *string             `pulumi:"rateControlMode"`
	SceneChangeDetection *bool               `pulumi:"sceneChangeDetection"`
	StretchMode          *string             `pulumi:"stretchMode"`
	SyncMode             *string             `pulumi:"syncMode"`
}

type H265Layer struct {
	AdaptiveBFrame  *bool    `pulumi:"adaptiveBFrame"`
	BFrames         *int     `pulumi:"bFrames"`
	Bitrate         int      `pulumi:"bitrate"`
	BufferWindow    *string  `pulumi:"bufferWindow"`
	Crf             *float64 `pulumi:"crf"`
	FrameRate       *string  `pulumi:"frameRate"`
	Height          *string  `pulumi:"height"`
	Label           *string  `pulumi:"label"`
	Level           *string  `pulumi:"level"`
	MaxBitrate      *int     `pulumi:"maxBitrate"`
	Profile         *string  `pulumi:"profile"`
	ReferenceFrames *int     `pulumi:"referenceFrames"`
	Slices          *int     `pulumi:"slices"`
	Width           *string  `pulumi:"width"`
}

type H265LayerResponse struct {
	AdaptiveBFrame  *bool    `pulumi:"adaptiveBFrame"`
	BFrames         *int     `pulumi:"bFrames"`
	Bitrate         int      `pulumi:"bitrate"`
	BufferWindow    *string  `pulumi:"bufferWindow"`
	Crf             *float64 `pulumi:"crf"`
	FrameRate       *string  `pulumi:"frameRate"`
	Height          *string  `pulumi:"height"`
	Label           *string  `pulumi:"label"`
	Level           *string  `pulumi:"level"`
	MaxBitrate      *int     `pulumi:"maxBitrate"`
	Profile         *string  `pulumi:"profile"`
	ReferenceFrames *int     `pulumi:"referenceFrames"`
	Slices          *int     `pulumi:"slices"`
	Width           *string  `pulumi:"width"`
}

type H265Video struct {
	Complexity           *string     `pulumi:"complexity"`
	KeyFrameInterval     *string     `pulumi:"keyFrameInterval"`
	Label                *string     `pulumi:"label"`
	Layers               []H265Layer `pulumi:"layers"`
	OdataType            string      `pulumi:"odataType"`
	SceneChangeDetection *bool       `pulumi:"sceneChangeDetection"`
	StretchMode          *string     `pulumi:"stretchMode"`
	SyncMode             *string     `pulumi:"syncMode"`
}

type H265VideoResponse struct {
	Complexity           *string             `pulumi:"complexity"`
	KeyFrameInterval     *string             `pulumi:"keyFrameInterval"`
	Label                *string             `pulumi:"label"`
	Layers               []H265LayerResponse `pulumi:"layers"`
	OdataType            string              `pulumi:"odataType"`
	SceneChangeDetection *bool               `pulumi:"sceneChangeDetection"`
	StretchMode          *string             `pulumi:"stretchMode"`
	SyncMode             *string             `pulumi:"syncMode"`
}

type Image struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	Range            *string `pulumi:"range"`
	Start            string  `pulumi:"start"`
	Step             *string `pulumi:"step"`
	StretchMode      *string `pulumi:"stretchMode"`
	SyncMode         *string `pulumi:"syncMode"`
}

type ImageFormat struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}

type ImageFormatResponse struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}

type ImageResponse struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	Range            *string `pulumi:"range"`
	Start            string  `pulumi:"start"`
	Step             *string `pulumi:"step"`
	StretchMode      *string `pulumi:"stretchMode"`
	SyncMode         *string `pulumi:"syncMode"`
}

type InputFile struct {
	Filename       *string       `pulumi:"filename"`
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}

type InputFileResponse struct {
	Filename       *string       `pulumi:"filename"`
	IncludedTracks []interface{} `pulumi:"includedTracks"`
	OdataType      string        `pulumi:"odataType"`
}

type JobErrorDetailResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type JobErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (JobErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorDetailResponse)(nil)).Elem()
}

func (o JobErrorDetailResponseOutput) ToJobErrorDetailResponseOutput() JobErrorDetailResponseOutput {
	return o
}

func (o JobErrorDetailResponseOutput) ToJobErrorDetailResponseOutputWithContext(ctx context.Context) JobErrorDetailResponseOutput {
	return o
}

func (o JobErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o JobErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

type JobErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (JobErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobErrorDetailResponse)(nil)).Elem()
}

func (o JobErrorDetailResponseArrayOutput) ToJobErrorDetailResponseArrayOutput() JobErrorDetailResponseArrayOutput {
	return o
}

func (o JobErrorDetailResponseArrayOutput) ToJobErrorDetailResponseArrayOutputWithContext(ctx context.Context) JobErrorDetailResponseArrayOutput {
	return o
}

func (o JobErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) JobErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobErrorDetailResponse {
		return vs[0].([]JobErrorDetailResponse)[vs[1].(int)]
	}).(JobErrorDetailResponseOutput)
}

type JobErrorResponse struct {
	Category string                   `pulumi:"category"`
	Code     string                   `pulumi:"code"`
	Details  []JobErrorDetailResponse `pulumi:"details"`
	Message  string                   `pulumi:"message"`
	Retry    string                   `pulumi:"retry"`
}

type JobErrorResponseOutput struct{ *pulumi.OutputState }

func (JobErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorResponse)(nil)).Elem()
}

func (o JobErrorResponseOutput) ToJobErrorResponseOutput() JobErrorResponseOutput {
	return o
}

func (o JobErrorResponseOutput) ToJobErrorResponseOutputWithContext(ctx context.Context) JobErrorResponseOutput {
	return o
}

func (o JobErrorResponseOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Category }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Details() JobErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v JobErrorResponse) []JobErrorDetailResponse { return v.Details }).(JobErrorDetailResponseArrayOutput)
}

func (o JobErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Retry() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Retry }).(pulumi.StringOutput)
}

type JobInputAsset struct {
	AssetName        string        `pulumi:"assetName"`
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}

type JobInputAssetResponse struct {
	AssetName        string        `pulumi:"assetName"`
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}

type JobInputClip struct {
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}

type JobInputClipResponse struct {
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}

type JobInputHttp struct {
	BaseUri          *string       `pulumi:"baseUri"`
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}

type JobInputHttpResponse struct {
	BaseUri          *string       `pulumi:"baseUri"`
	End              interface{}   `pulumi:"end"`
	Files            []string      `pulumi:"files"`
	InputDefinitions []interface{} `pulumi:"inputDefinitions"`
	Label            *string       `pulumi:"label"`
	OdataType        string        `pulumi:"odataType"`
	Start            interface{}   `pulumi:"start"`
}

type JobInputSequence struct {
	Inputs    []JobInputClip `pulumi:"inputs"`
	OdataType string         `pulumi:"odataType"`
}

type JobInputSequenceResponse struct {
	Inputs    []JobInputClipResponse `pulumi:"inputs"`
	OdataType string                 `pulumi:"odataType"`
}

type JobInputs struct {
	Inputs    []interface{} `pulumi:"inputs"`
	OdataType string        `pulumi:"odataType"`
}

type JobInputsResponse struct {
	Inputs    []interface{} `pulumi:"inputs"`
	OdataType string        `pulumi:"odataType"`
}

type JobOutputAsset struct {
	AssetName      string      `pulumi:"assetName"`
	Label          *string     `pulumi:"label"`
	OdataType      string      `pulumi:"odataType"`
	PresetOverride interface{} `pulumi:"presetOverride"`
}





type JobOutputAssetInput interface {
	pulumi.Input

	ToJobOutputAssetOutput() JobOutputAssetOutput
	ToJobOutputAssetOutputWithContext(context.Context) JobOutputAssetOutput
}

type JobOutputAssetArgs struct {
	AssetName      pulumi.StringInput    `pulumi:"assetName"`
	Label          pulumi.StringPtrInput `pulumi:"label"`
	OdataType      pulumi.StringInput    `pulumi:"odataType"`
	PresetOverride pulumi.Input          `pulumi:"presetOverride"`
}

func (JobOutputAssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAsset)(nil)).Elem()
}

func (i JobOutputAssetArgs) ToJobOutputAssetOutput() JobOutputAssetOutput {
	return i.ToJobOutputAssetOutputWithContext(context.Background())
}

func (i JobOutputAssetArgs) ToJobOutputAssetOutputWithContext(ctx context.Context) JobOutputAssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutputAssetOutput)
}





type JobOutputAssetArrayInput interface {
	pulumi.Input

	ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput
	ToJobOutputAssetArrayOutputWithContext(context.Context) JobOutputAssetArrayOutput
}

type JobOutputAssetArray []JobOutputAssetInput

func (JobOutputAssetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAsset)(nil)).Elem()
}

func (i JobOutputAssetArray) ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput {
	return i.ToJobOutputAssetArrayOutputWithContext(context.Background())
}

func (i JobOutputAssetArray) ToJobOutputAssetArrayOutputWithContext(ctx context.Context) JobOutputAssetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutputAssetArrayOutput)
}

type JobOutputAssetOutput struct{ *pulumi.OutputState }

func (JobOutputAssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAsset)(nil)).Elem()
}

func (o JobOutputAssetOutput) ToJobOutputAssetOutput() JobOutputAssetOutput {
	return o
}

func (o JobOutputAssetOutput) ToJobOutputAssetOutputWithContext(ctx context.Context) JobOutputAssetOutput {
	return o
}

func (o JobOutputAssetOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAsset) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o JobOutputAssetOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobOutputAsset) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobOutputAssetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAsset) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobOutputAssetOutput) PresetOverride() pulumi.AnyOutput {
	return o.ApplyT(func(v JobOutputAsset) interface{} { return v.PresetOverride }).(pulumi.AnyOutput)
}

type JobOutputAssetArrayOutput struct{ *pulumi.OutputState }

func (JobOutputAssetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAsset)(nil)).Elem()
}

func (o JobOutputAssetArrayOutput) ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput {
	return o
}

func (o JobOutputAssetArrayOutput) ToJobOutputAssetArrayOutputWithContext(ctx context.Context) JobOutputAssetArrayOutput {
	return o
}

func (o JobOutputAssetArrayOutput) Index(i pulumi.IntInput) JobOutputAssetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobOutputAsset {
		return vs[0].([]JobOutputAsset)[vs[1].(int)]
	}).(JobOutputAssetOutput)
}

type JobOutputAssetResponse struct {
	AssetName      string           `pulumi:"assetName"`
	EndTime        string           `pulumi:"endTime"`
	Error          JobErrorResponse `pulumi:"error"`
	Label          *string          `pulumi:"label"`
	OdataType      string           `pulumi:"odataType"`
	PresetOverride interface{}      `pulumi:"presetOverride"`
	Progress       int              `pulumi:"progress"`
	StartTime      string           `pulumi:"startTime"`
	State          string           `pulumi:"state"`
}

type JobOutputAssetResponseOutput struct{ *pulumi.OutputState }

func (JobOutputAssetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAssetResponse)(nil)).Elem()
}

func (o JobOutputAssetResponseOutput) ToJobOutputAssetResponseOutput() JobOutputAssetResponseOutput {
	return o
}

func (o JobOutputAssetResponseOutput) ToJobOutputAssetResponseOutputWithContext(ctx context.Context) JobOutputAssetResponseOutput {
	return o
}

func (o JobOutputAssetResponseOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) Error() JobErrorResponseOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) JobErrorResponse { return v.Error }).(JobErrorResponseOutput)
}

func (o JobOutputAssetResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o JobOutputAssetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) PresetOverride() pulumi.AnyOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) interface{} { return v.PresetOverride }).(pulumi.AnyOutput)
}

func (o JobOutputAssetResponseOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) int { return v.Progress }).(pulumi.IntOutput)
}

func (o JobOutputAssetResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.State }).(pulumi.StringOutput)
}

type JobOutputAssetResponseArrayOutput struct{ *pulumi.OutputState }

func (JobOutputAssetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAssetResponse)(nil)).Elem()
}

func (o JobOutputAssetResponseArrayOutput) ToJobOutputAssetResponseArrayOutput() JobOutputAssetResponseArrayOutput {
	return o
}

func (o JobOutputAssetResponseArrayOutput) ToJobOutputAssetResponseArrayOutputWithContext(ctx context.Context) JobOutputAssetResponseArrayOutput {
	return o
}

func (o JobOutputAssetResponseArrayOutput) Index(i pulumi.IntInput) JobOutputAssetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobOutputAssetResponse {
		return vs[0].([]JobOutputAssetResponse)[vs[1].(int)]
	}).(JobOutputAssetResponseOutput)
}

type JpgFormat struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}

type JpgFormatResponse struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}

type JpgImage struct {
	KeyFrameInterval *string    `pulumi:"keyFrameInterval"`
	Label            *string    `pulumi:"label"`
	Layers           []JpgLayer `pulumi:"layers"`
	OdataType        string     `pulumi:"odataType"`
	Range            *string    `pulumi:"range"`
	SpriteColumn     *int       `pulumi:"spriteColumn"`
	Start            string     `pulumi:"start"`
	Step             *string    `pulumi:"step"`
	StretchMode      *string    `pulumi:"stretchMode"`
	SyncMode         *string    `pulumi:"syncMode"`
}

type JpgImageResponse struct {
	KeyFrameInterval *string            `pulumi:"keyFrameInterval"`
	Label            *string            `pulumi:"label"`
	Layers           []JpgLayerResponse `pulumi:"layers"`
	OdataType        string             `pulumi:"odataType"`
	Range            *string            `pulumi:"range"`
	SpriteColumn     *int               `pulumi:"spriteColumn"`
	Start            string             `pulumi:"start"`
	Step             *string            `pulumi:"step"`
	StretchMode      *string            `pulumi:"stretchMode"`
	SyncMode         *string            `pulumi:"syncMode"`
}

type JpgLayer struct {
	Height  *string `pulumi:"height"`
	Label   *string `pulumi:"label"`
	Quality *int    `pulumi:"quality"`
	Width   *string `pulumi:"width"`
}

type JpgLayerResponse struct {
	Height  *string `pulumi:"height"`
	Label   *string `pulumi:"label"`
	Quality *int    `pulumi:"quality"`
	Width   *string `pulumi:"width"`
}

type Mp4Format struct {
	FilenamePattern string       `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}

type Mp4FormatResponse struct {
	FilenamePattern string               `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}

type MultiBitrateFormat struct {
	FilenamePattern string       `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}

type MultiBitrateFormatResponse struct {
	FilenamePattern string               `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}

type OutputFile struct {
	Labels []string `pulumi:"labels"`
}

type OutputFileResponse struct {
	Labels []string `pulumi:"labels"`
}

type PngFormat struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}

type PngFormatResponse struct {
	FilenamePattern string `pulumi:"filenamePattern"`
	OdataType       string `pulumi:"odataType"`
}

type PngImage struct {
	KeyFrameInterval *string    `pulumi:"keyFrameInterval"`
	Label            *string    `pulumi:"label"`
	Layers           []PngLayer `pulumi:"layers"`
	OdataType        string     `pulumi:"odataType"`
	Range            *string    `pulumi:"range"`
	Start            string     `pulumi:"start"`
	Step             *string    `pulumi:"step"`
	StretchMode      *string    `pulumi:"stretchMode"`
	SyncMode         *string    `pulumi:"syncMode"`
}

type PngImageResponse struct {
	KeyFrameInterval *string            `pulumi:"keyFrameInterval"`
	Label            *string            `pulumi:"label"`
	Layers           []PngLayerResponse `pulumi:"layers"`
	OdataType        string             `pulumi:"odataType"`
	Range            *string            `pulumi:"range"`
	Start            string             `pulumi:"start"`
	Step             *string            `pulumi:"step"`
	StretchMode      *string            `pulumi:"stretchMode"`
	SyncMode         *string            `pulumi:"syncMode"`
}

type PngLayer struct {
	Height *string `pulumi:"height"`
	Label  *string `pulumi:"label"`
	Width  *string `pulumi:"width"`
}

type PngLayerResponse struct {
	Height *string `pulumi:"height"`
	Label  *string `pulumi:"label"`
	Width  *string `pulumi:"width"`
}

type PresetConfigurations struct {
	Complexity                *string  `pulumi:"complexity"`
	InterleaveOutput          *string  `pulumi:"interleaveOutput"`
	KeyFrameIntervalInSeconds *float64 `pulumi:"keyFrameIntervalInSeconds"`
	MaxBitrateBps             *int     `pulumi:"maxBitrateBps"`
	MaxHeight                 *int     `pulumi:"maxHeight"`
	MaxLayers                 *int     `pulumi:"maxLayers"`
	MinBitrateBps             *int     `pulumi:"minBitrateBps"`
	MinHeight                 *int     `pulumi:"minHeight"`
}

type PresetConfigurationsResponse struct {
	Complexity                *string  `pulumi:"complexity"`
	InterleaveOutput          *string  `pulumi:"interleaveOutput"`
	KeyFrameIntervalInSeconds *float64 `pulumi:"keyFrameIntervalInSeconds"`
	MaxBitrateBps             *int     `pulumi:"maxBitrateBps"`
	MaxHeight                 *int     `pulumi:"maxHeight"`
	MaxLayers                 *int     `pulumi:"maxLayers"`
	MinBitrateBps             *int     `pulumi:"minBitrateBps"`
	MinHeight                 *int     `pulumi:"minHeight"`
}

type Rectangle struct {
	Height *string `pulumi:"height"`
	Left   *string `pulumi:"left"`
	Top    *string `pulumi:"top"`
	Width  *string `pulumi:"width"`
}

type RectangleResponse struct {
	Height *string `pulumi:"height"`
	Left   *string `pulumi:"left"`
	Top    *string `pulumi:"top"`
	Width  *string `pulumi:"width"`
}

type SelectAudioTrackByAttribute struct {
	Attribute      string  `pulumi:"attribute"`
	ChannelMapping *string `pulumi:"channelMapping"`
	Filter         string  `pulumi:"filter"`
	FilterValue    *string `pulumi:"filterValue"`
	OdataType      string  `pulumi:"odataType"`
}

type SelectAudioTrackByAttributeResponse struct {
	Attribute      string  `pulumi:"attribute"`
	ChannelMapping *string `pulumi:"channelMapping"`
	Filter         string  `pulumi:"filter"`
	FilterValue    *string `pulumi:"filterValue"`
	OdataType      string  `pulumi:"odataType"`
}

type SelectAudioTrackById struct {
	ChannelMapping *string `pulumi:"channelMapping"`
	OdataType      string  `pulumi:"odataType"`
	TrackId        float64 `pulumi:"trackId"`
}

type SelectAudioTrackByIdResponse struct {
	ChannelMapping *string `pulumi:"channelMapping"`
	OdataType      string  `pulumi:"odataType"`
	TrackId        float64 `pulumi:"trackId"`
}

type SelectVideoTrackByAttribute struct {
	Attribute   string  `pulumi:"attribute"`
	Filter      string  `pulumi:"filter"`
	FilterValue *string `pulumi:"filterValue"`
	OdataType   string  `pulumi:"odataType"`
}

type SelectVideoTrackByAttributeResponse struct {
	Attribute   string  `pulumi:"attribute"`
	Filter      string  `pulumi:"filter"`
	FilterValue *string `pulumi:"filterValue"`
	OdataType   string  `pulumi:"odataType"`
}

type SelectVideoTrackById struct {
	OdataType string  `pulumi:"odataType"`
	TrackId   float64 `pulumi:"trackId"`
}

type SelectVideoTrackByIdResponse struct {
	OdataType string  `pulumi:"odataType"`
	TrackId   float64 `pulumi:"trackId"`
}

type StandardEncoderPreset struct {
	Codecs              []interface{}     `pulumi:"codecs"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Filters             *Filters          `pulumi:"filters"`
	Formats             []interface{}     `pulumi:"formats"`
	OdataType           string            `pulumi:"odataType"`
}

type StandardEncoderPresetResponse struct {
	Codecs              []interface{}     `pulumi:"codecs"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	Filters             *FiltersResponse  `pulumi:"filters"`
	Formats             []interface{}     `pulumi:"formats"`
	OdataType           string            `pulumi:"odataType"`
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type TransformOutputType struct {
	OnError          *string     `pulumi:"onError"`
	Preset           interface{} `pulumi:"preset"`
	RelativePriority *string     `pulumi:"relativePriority"`
}





type TransformOutputTypeInput interface {
	pulumi.Input

	ToTransformOutputTypeOutput() TransformOutputTypeOutput
	ToTransformOutputTypeOutputWithContext(context.Context) TransformOutputTypeOutput
}

type TransformOutputTypeArgs struct {
	OnError          pulumi.StringPtrInput `pulumi:"onError"`
	Preset           pulumi.Input          `pulumi:"preset"`
	RelativePriority pulumi.StringPtrInput `pulumi:"relativePriority"`
}

func (TransformOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputType)(nil)).Elem()
}

func (i TransformOutputTypeArgs) ToTransformOutputTypeOutput() TransformOutputTypeOutput {
	return i.ToTransformOutputTypeOutputWithContext(context.Background())
}

func (i TransformOutputTypeArgs) ToTransformOutputTypeOutputWithContext(ctx context.Context) TransformOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutputTypeOutput)
}





type TransformOutputTypeArrayInput interface {
	pulumi.Input

	ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput
	ToTransformOutputTypeArrayOutputWithContext(context.Context) TransformOutputTypeArrayOutput
}

type TransformOutputTypeArray []TransformOutputTypeInput

func (TransformOutputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputType)(nil)).Elem()
}

func (i TransformOutputTypeArray) ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput {
	return i.ToTransformOutputTypeArrayOutputWithContext(context.Background())
}

func (i TransformOutputTypeArray) ToTransformOutputTypeArrayOutputWithContext(ctx context.Context) TransformOutputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutputTypeArrayOutput)
}

type TransformOutputTypeOutput struct{ *pulumi.OutputState }

func (TransformOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputType)(nil)).Elem()
}

func (o TransformOutputTypeOutput) ToTransformOutputTypeOutput() TransformOutputTypeOutput {
	return o
}

func (o TransformOutputTypeOutput) ToTransformOutputTypeOutputWithContext(ctx context.Context) TransformOutputTypeOutput {
	return o
}

func (o TransformOutputTypeOutput) OnError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputType) *string { return v.OnError }).(pulumi.StringPtrOutput)
}

func (o TransformOutputTypeOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v TransformOutputType) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o TransformOutputTypeOutput) RelativePriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputType) *string { return v.RelativePriority }).(pulumi.StringPtrOutput)
}

type TransformOutputTypeArrayOutput struct{ *pulumi.OutputState }

func (TransformOutputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputType)(nil)).Elem()
}

func (o TransformOutputTypeArrayOutput) ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput {
	return o
}

func (o TransformOutputTypeArrayOutput) ToTransformOutputTypeArrayOutputWithContext(ctx context.Context) TransformOutputTypeArrayOutput {
	return o
}

func (o TransformOutputTypeArrayOutput) Index(i pulumi.IntInput) TransformOutputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TransformOutputType {
		return vs[0].([]TransformOutputType)[vs[1].(int)]
	}).(TransformOutputTypeOutput)
}

type TransformOutputResponse struct {
	OnError          *string     `pulumi:"onError"`
	Preset           interface{} `pulumi:"preset"`
	RelativePriority *string     `pulumi:"relativePriority"`
}

type TransformOutputResponseOutput struct{ *pulumi.OutputState }

func (TransformOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputResponse)(nil)).Elem()
}

func (o TransformOutputResponseOutput) ToTransformOutputResponseOutput() TransformOutputResponseOutput {
	return o
}

func (o TransformOutputResponseOutput) ToTransformOutputResponseOutputWithContext(ctx context.Context) TransformOutputResponseOutput {
	return o
}

func (o TransformOutputResponseOutput) OnError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputResponse) *string { return v.OnError }).(pulumi.StringPtrOutput)
}

func (o TransformOutputResponseOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v TransformOutputResponse) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o TransformOutputResponseOutput) RelativePriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputResponse) *string { return v.RelativePriority }).(pulumi.StringPtrOutput)
}

type TransformOutputResponseArrayOutput struct{ *pulumi.OutputState }

func (TransformOutputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputResponse)(nil)).Elem()
}

func (o TransformOutputResponseArrayOutput) ToTransformOutputResponseArrayOutput() TransformOutputResponseArrayOutput {
	return o
}

func (o TransformOutputResponseArrayOutput) ToTransformOutputResponseArrayOutputWithContext(ctx context.Context) TransformOutputResponseArrayOutput {
	return o
}

func (o TransformOutputResponseArrayOutput) Index(i pulumi.IntInput) TransformOutputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TransformOutputResponse {
		return vs[0].([]TransformOutputResponse)[vs[1].(int)]
	}).(TransformOutputResponseOutput)
}

type TransportStreamFormat struct {
	FilenamePattern string       `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}

type TransportStreamFormatResponse struct {
	FilenamePattern string               `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}

type UtcClipTime struct {
	OdataType string `pulumi:"odataType"`
	Time      string `pulumi:"time"`
}

type UtcClipTimeResponse struct {
	OdataType string `pulumi:"odataType"`
	Time      string `pulumi:"time"`
}

type Video struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	StretchMode      *string `pulumi:"stretchMode"`
	SyncMode         *string `pulumi:"syncMode"`
}

type VideoAnalyzerPreset struct {
	AudioLanguage       *string           `pulumi:"audioLanguage"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	InsightsToExtract   *string           `pulumi:"insightsToExtract"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
}

type VideoAnalyzerPresetResponse struct {
	AudioLanguage       *string           `pulumi:"audioLanguage"`
	ExperimentalOptions map[string]string `pulumi:"experimentalOptions"`
	InsightsToExtract   *string           `pulumi:"insightsToExtract"`
	Mode                *string           `pulumi:"mode"`
	OdataType           string            `pulumi:"odataType"`
}

type VideoOverlay struct {
	AudioGainLevel  *float64   `pulumi:"audioGainLevel"`
	CropRectangle   *Rectangle `pulumi:"cropRectangle"`
	End             *string    `pulumi:"end"`
	FadeInDuration  *string    `pulumi:"fadeInDuration"`
	FadeOutDuration *string    `pulumi:"fadeOutDuration"`
	InputLabel      string     `pulumi:"inputLabel"`
	OdataType       string     `pulumi:"odataType"`
	Opacity         *float64   `pulumi:"opacity"`
	Position        *Rectangle `pulumi:"position"`
	Start           *string    `pulumi:"start"`
}

type VideoOverlayResponse struct {
	AudioGainLevel  *float64           `pulumi:"audioGainLevel"`
	CropRectangle   *RectangleResponse `pulumi:"cropRectangle"`
	End             *string            `pulumi:"end"`
	FadeInDuration  *string            `pulumi:"fadeInDuration"`
	FadeOutDuration *string            `pulumi:"fadeOutDuration"`
	InputLabel      string             `pulumi:"inputLabel"`
	OdataType       string             `pulumi:"odataType"`
	Opacity         *float64           `pulumi:"opacity"`
	Position        *RectangleResponse `pulumi:"position"`
	Start           *string            `pulumi:"start"`
}

type VideoResponse struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	StretchMode      *string `pulumi:"stretchMode"`
	SyncMode         *string `pulumi:"syncMode"`
}

type VideoTrackDescriptor struct {
	OdataType string `pulumi:"odataType"`
}

type VideoTrackDescriptorResponse struct {
	OdataType string `pulumi:"odataType"`
}

func init() {
	pulumi.RegisterOutputType(JobErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(JobErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(JobErrorResponseOutput{})
	pulumi.RegisterOutputType(JobOutputAssetOutput{})
	pulumi.RegisterOutputType(JobOutputAssetArrayOutput{})
	pulumi.RegisterOutputType(JobOutputAssetResponseOutput{})
	pulumi.RegisterOutputType(JobOutputAssetResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TransformOutputTypeOutput{})
	pulumi.RegisterOutputType(TransformOutputTypeArrayOutput{})
	pulumi.RegisterOutputType(TransformOutputResponseOutput{})
	pulumi.RegisterOutputType(TransformOutputResponseArrayOutput{})
}
