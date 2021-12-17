


package v20190501preview

type LiveEventEncodingType string

const (
	LiveEventEncodingTypeNone         = LiveEventEncodingType("None")
	LiveEventEncodingTypeBasic        = LiveEventEncodingType("Basic")
	LiveEventEncodingTypeStandard     = LiveEventEncodingType("Standard")
	LiveEventEncodingTypePremium1080p = LiveEventEncodingType("Premium1080p")
)

type LiveEventInputProtocol string

const (
	LiveEventInputProtocolFragmentedMP4 = LiveEventInputProtocol("FragmentedMP4")
	LiveEventInputProtocolRTMP          = LiveEventInputProtocol("RTMP")
)

type StreamOptionsFlag string

const (
	StreamOptionsFlagDefault    = StreamOptionsFlag("Default")
	StreamOptionsFlagLowLatency = StreamOptionsFlag("LowLatency")
)

func init() {
}
