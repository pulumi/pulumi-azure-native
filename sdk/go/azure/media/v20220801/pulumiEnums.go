


package v20220801

type LiveEventEncodingType string

const (
	// This is the same as PassthroughStandard, please see description below. This enumeration value is being deprecated.
	LiveEventEncodingTypeNone = LiveEventEncodingType("None")
	// A contribution live encoder sends a single bitrate stream to the live event and Media Services creates multiple bitrate streams. The output cannot exceed 720p in resolution.
	LiveEventEncodingTypeStandard = LiveEventEncodingType("Standard")
	// A contribution live encoder sends a single bitrate stream to the live event and Media Services creates multiple bitrate streams. The output cannot exceed 1080p in resolution.
	LiveEventEncodingTypePremium1080p = LiveEventEncodingType("Premium1080p")
	// The ingested stream passes through the live event from the contribution encoder without any further processing. In the PassthroughBasic mode, ingestion is limited to up to 5Mbps and only 1 concurrent live output is allowed. Live transcription is not available.
	LiveEventEncodingTypePassthroughBasic = LiveEventEncodingType("PassthroughBasic")
	// The ingested stream passes through the live event from the contribution encoder without any further processing. Live transcription is available. Ingestion bitrate limits are much higher and up to 3 concurrent live outputs are allowed.
	LiveEventEncodingTypePassthroughStandard = LiveEventEncodingType("PassthroughStandard")
)

type LiveEventInputProtocol string

const (
	// Smooth Streaming input will be sent by the contribution encoder to the live event.
	LiveEventInputProtocolFragmentedMP4 = LiveEventInputProtocol("FragmentedMP4")
	// RTMP input will be sent by the contribution encoder to the live event.
	LiveEventInputProtocolRTMP = LiveEventInputProtocol("RTMP")
)

type StreamOptionsFlag string

const (
	// Live streaming with no special latency optimizations.
	StreamOptionsFlagDefault = StreamOptionsFlag("Default")
	// The live event provides lower end to end latency by reducing its internal buffers.
	StreamOptionsFlagLowLatency = StreamOptionsFlag("LowLatency")
	// The live event is optimized for end to end latency. This option is only available for encoding live events with RTMP input. The outputs can be streamed using HLS or DASH formats. The outputs' archive or DVR rewind length is limited to 6 hours. Use "LowLatency" stream option for all other scenarios.
	StreamOptionsFlagLowLatencyV2 = StreamOptionsFlag("LowLatencyV2")
)

type StretchMode string

const (
	// Strictly respects the output resolution specified in the encoding preset without considering the pixel aspect ratio or display aspect ratio of the input video.
	StretchModeNone = StretchMode("None")
	// Override the output resolution, and change it to match the display aspect ratio of the input, without padding. For example, if the input is 1920x1080 and the encoding preset asks for 1280x1280, then the value in the preset is overridden, and the output will be at 1280x720, which maintains the input aspect ratio of 16:9.
	StretchModeAutoSize = StretchMode("AutoSize")
	// Pad the output (with either letterbox or pillar box) to honor the output resolution, while ensuring that the active video region in the output has the same aspect ratio as the input. For example, if the input is 1920x1080 and the encoding preset asks for 1280x1280, then the output will be at 1280x1280, which contains an inner rectangle of 1280x720 at aspect ratio of 16:9, and pillar box regions 280 pixels wide at the left and right.
	StretchModeAutoFit = StretchMode("AutoFit")
)

func init() {
}
