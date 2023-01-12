


package v20220701

type AacAudioProfile string

const (
	// Specifies that the output audio is to be encoded into AAC Low Complexity profile (AAC-LC).
	AacAudioProfileAacLc = AacAudioProfile("AacLc")
	// Specifies that the output audio is to be encoded into HE-AAC v1 profile.
	AacAudioProfileHeAacV1 = AacAudioProfile("HeAacV1")
	// Specifies that the output audio is to be encoded into HE-AAC v2 profile.
	AacAudioProfileHeAacV2 = AacAudioProfile("HeAacV2")
)

type AnalysisResolution string

const (
	AnalysisResolutionSourceResolution   = AnalysisResolution("SourceResolution")
	AnalysisResolutionStandardDefinition = AnalysisResolution("StandardDefinition")
)

type AttributeFilter string

const (
	// All tracks will be included.
	AttributeFilterAll = AttributeFilter("All")
	// The first track will be included when the attribute is sorted in descending order.  Generally used to select the largest bitrate.
	AttributeFilterTop = AttributeFilter("Top")
	// The first track will be included when the attribute is sorted in ascending order.  Generally used to select the smallest bitrate.
	AttributeFilterBottom = AttributeFilter("Bottom")
	// Any tracks that have an attribute equal to the value given will be included.
	AttributeFilterValueEquals = AttributeFilter("ValueEquals")
)

type AudioAnalysisMode string

const (
	// Performs all operations included in the Basic mode, additionally performing language detection and speaker diarization.
	AudioAnalysisModeStandard = AudioAnalysisMode("Standard")
	// This mode performs speech-to-text transcription and generation of a VTT subtitle/caption file. The output of this mode includes an Insights JSON file including only the keywords, transcription,and timing information. Automatic language detection and speaker diarization are not included in this mode.
	AudioAnalysisModeBasic = AudioAnalysisMode("Basic")
)

type BlurType string

const (
	// Box: debug filter, bounding box only
	BlurTypeBox = BlurType("Box")
	// Low: box-car blur filter
	BlurTypeLow = BlurType("Low")
	// Med: Gaussian blur filter
	BlurTypeMed = BlurType("Med")
	// High: Confuse blur filter
	BlurTypeHigh = BlurType("High")
	// Black: Black out filter
	BlurTypeBlack = BlurType("Black")
)

type ChannelMapping string

const (
	// The Front Left Channel.
	ChannelMappingFrontLeft = ChannelMapping("FrontLeft")
	// The Front Right Channel.
	ChannelMappingFrontRight = ChannelMapping("FrontRight")
	// The Center Channel.
	ChannelMappingCenter = ChannelMapping("Center")
	// Low Frequency Effects Channel.  Sometimes referred to as the subwoofer.
	ChannelMappingLowFrequencyEffects = ChannelMapping("LowFrequencyEffects")
	// The Back Left Channel.  Sometimes referred to as the Left Surround Channel.
	ChannelMappingBackLeft = ChannelMapping("BackLeft")
	// The Back Right Channel.  Sometimes referred to as the Right Surround Channel.
	ChannelMappingBackRight = ChannelMapping("BackRight")
	// The Left Stereo channel.  Sometimes referred to as Down Mix Left.
	ChannelMappingStereoLeft = ChannelMapping("StereoLeft")
	// The Right Stereo channel.  Sometimes referred to as Down Mix Right.
	ChannelMappingStereoRight = ChannelMapping("StereoRight")
)

type Complexity string

const (
	// Configures the encoder to use settings optimized for faster encoding. Quality is sacrificed to decrease encoding time.
	ComplexitySpeed = Complexity("Speed")
	// Configures the encoder to use settings that achieve a balance between speed and quality.
	ComplexityBalanced = Complexity("Balanced")
	// Configures the encoder to use settings optimized to produce higher quality output at the expense of slower overall encode time.
	ComplexityQuality = Complexity("Quality")
)

type DeinterlaceMode string

const (
	// Disables de-interlacing of the source video.
	DeinterlaceModeOff = DeinterlaceMode("Off")
	// Apply automatic pixel adaptive de-interlacing on each frame in the input video.
	DeinterlaceModeAutoPixelAdaptive = DeinterlaceMode("AutoPixelAdaptive")
)

type DeinterlaceParity string

const (
	// Automatically detect the order of fields
	DeinterlaceParityAuto = DeinterlaceParity("Auto")
	// Apply top field first processing of input video.
	DeinterlaceParityTopFieldFirst = DeinterlaceParity("TopFieldFirst")
	// Apply bottom field first processing of input video.
	DeinterlaceParityBottomFieldFirst = DeinterlaceParity("BottomFieldFirst")
)

type EncoderNamedPreset string

const (
	// Produces an MP4 file where the video is encoded with H.264 codec at 2200 kbps and a picture height of 480 pixels, and the stereo audio is encoded with AAC-LC codec at 128 kbps.
	EncoderNamedPresetH264SingleBitrateSD = EncoderNamedPreset("H264SingleBitrateSD")
	// Produces an MP4 file where the video is encoded with H.264 codec at 4500 kbps and a picture height of 720 pixels, and the stereo audio is encoded with AAC-LC codec at 128 kbps.
	EncoderNamedPresetH264SingleBitrate720p = EncoderNamedPreset("H264SingleBitrate720p")
	// Produces an MP4 file where the video is encoded with H.264 codec at 6750 kbps and a picture height of 1080 pixels, and the stereo audio is encoded with AAC-LC codec at 128 kbps.
	EncoderNamedPresetH264SingleBitrate1080p = EncoderNamedPreset("H264SingleBitrate1080p")
	// Produces a set of GOP aligned MP4 files with H.264 video and stereo AAC audio. Auto-generates a bitrate ladder based on the input resolution, bitrate and frame rate. The auto-generated preset will never exceed the input resolution. For example, if the input is 720p, output will remain 720p at best.
	EncoderNamedPresetAdaptiveStreaming = EncoderNamedPreset("AdaptiveStreaming")
	// Produces a single MP4 file containing only AAC stereo audio encoded at 192 kbps.
	EncoderNamedPresetAACGoodQualityAudio = EncoderNamedPreset("AACGoodQualityAudio")
	// Produces a single MP4 file containing only DD(Digital Dolby) stereo audio encoded at 192 kbps.
	EncoderNamedPresetDDGoodQualityAudio = EncoderNamedPreset("DDGoodQualityAudio")
	// Exposes an experimental preset for content-aware encoding. Given any input content, the service attempts to automatically determine the optimal number of layers, appropriate bitrate and resolution settings for delivery by adaptive streaming. The underlying algorithms will continue to evolve over time. The output will contain MP4 files with video and audio interleaved.
	EncoderNamedPresetContentAwareEncodingExperimental = EncoderNamedPreset("ContentAwareEncodingExperimental")
	// Produces a set of GOP-aligned MP4s by using content-aware encoding. Given any input content, the service performs an initial lightweight analysis of the input content, and uses the results to determine the optimal number of layers, appropriate bitrate and resolution settings for delivery by adaptive streaming. This preset is particularly effective for low and medium complexity videos, where the output files will be at lower bitrates but at a quality that still delivers a good experience to viewers. The output will contain MP4 files with video and audio interleaved.
	EncoderNamedPresetContentAwareEncoding = EncoderNamedPreset("ContentAwareEncoding")
	// Copy all video and audio streams from the input asset as non-interleaved video and audio output files. This preset can be used to clip an existing asset or convert a group of key frame (GOP) aligned MP4 files as an asset that can be streamed.
	EncoderNamedPresetCopyAllBitrateNonInterleaved = EncoderNamedPreset("CopyAllBitrateNonInterleaved")
	// Produces a set of 8 GOP-aligned MP4 files, ranging from 6000 kbps to 400 kbps, and stereo AAC audio. Resolution starts at 1080p and goes down to 180p.
	EncoderNamedPresetH264MultipleBitrate1080p = EncoderNamedPreset("H264MultipleBitrate1080p")
	// Produces a set of 6 GOP-aligned MP4 files, ranging from 3400 kbps to 400 kbps, and stereo AAC audio. Resolution starts at 720p and goes down to 180p.
	EncoderNamedPresetH264MultipleBitrate720p = EncoderNamedPreset("H264MultipleBitrate720p")
	// Produces a set of 5 GOP-aligned MP4 files, ranging from 1900kbps to 400 kbps, and stereo AAC audio. Resolution starts at 480p and goes down to 240p.
	EncoderNamedPresetH264MultipleBitrateSD = EncoderNamedPreset("H264MultipleBitrateSD")
	// Produces a set of GOP-aligned MP4s by using content-aware encoding. Given any input content, the service performs an initial lightweight analysis of the input content, and uses the results to determine the optimal number of layers, appropriate bitrate and resolution settings for delivery by adaptive streaming. This preset is particularly effective for low and medium complexity videos, where the output files will be at lower bitrates but at a quality that still delivers a good experience to viewers. The output will contain MP4 files with video and audio interleaved.
	EncoderNamedPresetH265ContentAwareEncoding = EncoderNamedPreset("H265ContentAwareEncoding")
	// Produces a set of GOP aligned MP4 files with H.265 video and stereo AAC audio. Auto-generates a bitrate ladder based on the input resolution, bitrate and frame rate. The auto-generated preset will never exceed the input resolution. For example, if the input is 720p, output will remain 720p at best.
	EncoderNamedPresetH265AdaptiveStreaming = EncoderNamedPreset("H265AdaptiveStreaming")
	// Produces an MP4 file where the video is encoded with H.265 codec at 1800 kbps and a picture height of 720 pixels, and the stereo audio is encoded with AAC-LC codec at 128 kbps.
	EncoderNamedPresetH265SingleBitrate720p = EncoderNamedPreset("H265SingleBitrate720p")
	// Produces an MP4 file where the video is encoded with H.265 codec at 3500 kbps and a picture height of 1080 pixels, and the stereo audio is encoded with AAC-LC codec at 128 kbps.
	EncoderNamedPresetH265SingleBitrate1080p = EncoderNamedPreset("H265SingleBitrate1080p")
	// Produces an MP4 file where the video is encoded with H.265 codec at 9500 kbps and a picture height of 2160 pixels, and the stereo audio is encoded with AAC-LC codec at 128 kbps.
	EncoderNamedPresetH265SingleBitrate4K = EncoderNamedPreset("H265SingleBitrate4K")
)

type EntropyMode string

const (
	// Context Adaptive Binary Arithmetic Coder (CABAC) entropy encoding.
	EntropyModeCabac = EntropyMode("Cabac")
	// Context Adaptive Variable Length Coder (CAVLC) entropy encoding.
	EntropyModeCavlc = EntropyMode("Cavlc")
)

type FaceRedactorMode string

const (
	// Analyze mode detects faces and outputs a metadata file with the results. Allows editing of the metadata file before faces are blurred with Redact mode.
	FaceRedactorModeAnalyze = FaceRedactorMode("Analyze")
	// Redact mode consumes the metadata file from Analyze mode and redacts the faces found.
	FaceRedactorModeRedact = FaceRedactorMode("Redact")
	// Combined mode does the Analyze and Redact steps in one pass when editing the analyzed faces is not desired.
	FaceRedactorModeCombined = FaceRedactorMode("Combined")
)

type H264Complexity string

const (
	// Tells the encoder to use settings that are optimized for faster encoding. Quality is sacrificed to decrease encoding time.
	H264ComplexitySpeed = H264Complexity("Speed")
	// Tells the encoder to use settings that achieve a balance between speed and quality.
	H264ComplexityBalanced = H264Complexity("Balanced")
	// Tells the encoder to use settings that are optimized to produce higher quality output at the expense of slower overall encode time.
	H264ComplexityQuality = H264Complexity("Quality")
)

type H264RateControlMode string

const (
	// Average Bitrate (ABR) mode that hits the target bitrate: Default mode.
	H264RateControlModeABR = H264RateControlMode("ABR")
	// Constant Bitrate (CBR) mode that tightens bitrate variations around target bitrate.
	H264RateControlModeCBR = H264RateControlMode("CBR")
	// Constant Rate Factor (CRF) mode that targets at constant subjective quality.
	H264RateControlModeCRF = H264RateControlMode("CRF")
)

type H264VideoProfile string

const (
	// Tells the encoder to automatically determine the appropriate H.264 profile.
	H264VideoProfileAuto = H264VideoProfile("Auto")
	// Baseline profile
	H264VideoProfileBaseline = H264VideoProfile("Baseline")
	// Main profile
	H264VideoProfileMain = H264VideoProfile("Main")
	// High profile.
	H264VideoProfileHigh = H264VideoProfile("High")
	// High 4:2:2 profile.
	H264VideoProfileHigh422 = H264VideoProfile("High422")
	// High 4:4:4 predictive profile.
	H264VideoProfileHigh444 = H264VideoProfile("High444")
)

type H265Complexity string

const (
	// Tells the encoder to use settings that are optimized for faster encoding. Quality is sacrificed to decrease encoding time.
	H265ComplexitySpeed = H265Complexity("Speed")
	// Tells the encoder to use settings that achieve a balance between speed and quality.
	H265ComplexityBalanced = H265Complexity("Balanced")
	// Tells the encoder to use settings that are optimized to produce higher quality output at the expense of slower overall encode time.
	H265ComplexityQuality = H265Complexity("Quality")
)

type H265VideoProfile string

const (
	// Tells the encoder to automatically determine the appropriate H.265 profile.
	H265VideoProfileAuto = H265VideoProfile("Auto")
	// Main profile (https://x265.readthedocs.io/en/default/cli.html?highlight=profile#profile-level-tier)
	H265VideoProfileMain = H265VideoProfile("Main")
	// Main 10 profile (https://en.wikipedia.org/wiki/High_Efficiency_Video_Coding#Main_10)
	H265VideoProfileMain10 = H265VideoProfile("Main10")
)

type InsightsType string

const (
	// Generate audio only insights. Ignore video even if present. Fails if no audio is present.
	InsightsTypeAudioInsightsOnly = InsightsType("AudioInsightsOnly")
	// Generate video only insights. Ignore audio if present. Fails if no video is present.
	InsightsTypeVideoInsightsOnly = InsightsType("VideoInsightsOnly")
	// Generate both audio and video insights. Fails if either audio or video Insights fail.
	InsightsTypeAllInsights = InsightsType("AllInsights")
)

type InterleaveOutput string

const (
	// The output is video-only or audio-only.
	InterleaveOutputNonInterleavedOutput = InterleaveOutput("NonInterleavedOutput")
	// The output includes both audio and video.
	InterleaveOutputInterleavedOutput = InterleaveOutput("InterleavedOutput")
)

type OnErrorType string

const (
	// Tells the service that if this TransformOutput fails, then any other incomplete TransformOutputs can be stopped.
	OnErrorTypeStopProcessingJob = OnErrorType("StopProcessingJob")
	// Tells the service that if this TransformOutput fails, then allow any other TransformOutput to continue.
	OnErrorTypeContinueJob = OnErrorType("ContinueJob")
)

type Priority string

const (
	// Used for TransformOutputs that can be generated after Normal and High priority TransformOutputs.
	PriorityLow = Priority("Low")
	// Used for TransformOutputs that can be generated at Normal priority.
	PriorityNormal = Priority("Normal")
	// Used for TransformOutputs that should take precedence over others.
	PriorityHigh = Priority("High")
)

type Rotation string

const (
	// Automatically detect and rotate as needed.
	RotationAuto = Rotation("Auto")
	// Do not rotate the video.  If the output format supports it, any metadata about rotation is kept intact.
	RotationNone = Rotation("None")
	// Do not rotate the video but remove any metadata about the rotation.
	RotationRotate0 = Rotation("Rotate0")
	// Rotate 90 degrees clockwise.
	RotationRotate90 = Rotation("Rotate90")
	// Rotate 180 degrees clockwise.
	RotationRotate180 = Rotation("Rotate180")
	// Rotate 270 degrees clockwise.
	RotationRotate270 = Rotation("Rotate270")
)

type StretchMode string

const (
	// Strictly respect the output resolution without considering the pixel aspect ratio or display aspect ratio of the input video.
	StretchModeNone = StretchMode("None")
	// Override the output resolution, and change it to match the display aspect ratio of the input, without padding. For example, if the input is 1920x1080 and the encoding preset asks for 1280x1280, then the value in the preset is overridden, and the output will be at 1280x720, which maintains the input aspect ratio of 16:9.
	StretchModeAutoSize = StretchMode("AutoSize")
	// Pad the output (with either letterbox or pillar box) to honor the output resolution, while ensuring that the active video region in the output has the same aspect ratio as the input. For example, if the input is 1920x1080 and the encoding preset asks for 1280x1280, then the output will be at 1280x1280, which contains an inner rectangle of 1280x720 at aspect ratio of 16:9, and pillar box regions 280 pixels wide at the left and right.
	StretchModeAutoFit = StretchMode("AutoFit")
)

type TrackAttribute string

const (
	// The bitrate of the track.
	TrackAttributeBitrate = TrackAttribute("Bitrate")
	// The language of the track.
	TrackAttributeLanguage = TrackAttribute("Language")
)

type VideoSyncMode string

const (
	// This is the default method. Chooses between Cfr and Vfr depending on muxer capabilities. For output format MP4, the default mode is Cfr.
	VideoSyncModeAuto = VideoSyncMode("Auto")
	// The presentation timestamps on frames are passed through from the input file to the output file writer. Recommended when the input source has variable frame rate, and are attempting to produce multiple layers for adaptive streaming in the output which have aligned GOP boundaries. Note: if two or more frames in the input have duplicate timestamps, then the output will also have the same behavior
	VideoSyncModePassthrough = VideoSyncMode("Passthrough")
	// Input frames will be repeated and/or dropped as needed to achieve exactly the requested constant frame rate. Recommended when the output frame rate is explicitly set at a specified value
	VideoSyncModeCfr = VideoSyncMode("Cfr")
	// Similar to the Passthrough mode, but if the input has frames that have duplicate timestamps, then only one frame is passed through to the output, and others are dropped. Recommended when the number of output frames is expected to be equal to the number of input frames. For example, the output is used to calculate a quality metric like PSNR against the input
	VideoSyncModeVfr = VideoSyncMode("Vfr")
)

func init() {
}
