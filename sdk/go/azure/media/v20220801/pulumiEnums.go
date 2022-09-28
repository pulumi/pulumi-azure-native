


package v20220801

type AssetContainerPermission string

const (
	// The SAS URL will allow read access to the container.
	AssetContainerPermissionRead = AssetContainerPermission("Read")
	// The SAS URL will allow read and write access to the container.
	AssetContainerPermissionReadWrite = AssetContainerPermission("ReadWrite")
	// The SAS URL will allow read, write and delete access to the container.
	AssetContainerPermissionReadWriteDelete = AssetContainerPermission("ReadWriteDelete")
)

type ContentKeyPolicyFairPlayRentalAndLeaseKeyType string

const (
	// Represents a ContentKeyPolicyFairPlayRentalAndLeaseKeyType that is unavailable in current API version.
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUnknown = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("Unknown")
	// Key duration is not specified.
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUndefined = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("Undefined")
	// Dual expiry for offline rental.
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeDualExpiry = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("DualExpiry")
	// Content key can be persisted with an unlimited duration
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentUnlimited = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("PersistentUnlimited")
	// Content key can be persisted and the valid duration is limited by the Rental Duration value
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentLimited = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("PersistentLimited")
)

type ContentKeyPolicyPlayReadyContentType string

const (
	// Represents a ContentKeyPolicyPlayReadyContentType that is unavailable in current API version.
	ContentKeyPolicyPlayReadyContentTypeUnknown = ContentKeyPolicyPlayReadyContentType("Unknown")
	// Unspecified content type.
	ContentKeyPolicyPlayReadyContentTypeUnspecified = ContentKeyPolicyPlayReadyContentType("Unspecified")
	// Ultraviolet download content type.
	ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload = ContentKeyPolicyPlayReadyContentType("UltraVioletDownload")
	// Ultraviolet streaming content type.
	ContentKeyPolicyPlayReadyContentTypeUltraVioletStreaming = ContentKeyPolicyPlayReadyContentType("UltraVioletStreaming")
)

type ContentKeyPolicyPlayReadyLicenseType string

const (
	// Represents a ContentKeyPolicyPlayReadyLicenseType that is unavailable in current API version.
	ContentKeyPolicyPlayReadyLicenseTypeUnknown = ContentKeyPolicyPlayReadyLicenseType("Unknown")
	// Non persistent license.
	ContentKeyPolicyPlayReadyLicenseTypeNonPersistent = ContentKeyPolicyPlayReadyLicenseType("NonPersistent")
	// Persistent license. Allows offline playback.
	ContentKeyPolicyPlayReadyLicenseTypePersistent = ContentKeyPolicyPlayReadyLicenseType("Persistent")
)

type ContentKeyPolicyPlayReadyUnknownOutputPassingOption string

const (
	// Represents a ContentKeyPolicyPlayReadyUnknownOutputPassingOption that is unavailable in current API version.
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionUnknown = ContentKeyPolicyPlayReadyUnknownOutputPassingOption("Unknown")
	// Passing the video portion of protected content to an Unknown Output is not allowed.
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed = ContentKeyPolicyPlayReadyUnknownOutputPassingOption("NotAllowed")
	// Passing the video portion of protected content to an Unknown Output is allowed.
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowed = ContentKeyPolicyPlayReadyUnknownOutputPassingOption("Allowed")
	// Passing the video portion of protected content to an Unknown Output is allowed but with constrained resolution.
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowedWithVideoConstriction = ContentKeyPolicyPlayReadyUnknownOutputPassingOption("AllowedWithVideoConstriction")
)

type ContentKeyPolicyRestrictionTokenType string

const (
	// Represents a ContentKeyPolicyRestrictionTokenType that is unavailable in current API version.
	ContentKeyPolicyRestrictionTokenTypeUnknown = ContentKeyPolicyRestrictionTokenType("Unknown")
	// Simple Web Token.
	ContentKeyPolicyRestrictionTokenTypeSwt = ContentKeyPolicyRestrictionTokenType("Swt")
	// JSON Web Token.
	ContentKeyPolicyRestrictionTokenTypeJwt = ContentKeyPolicyRestrictionTokenType("Jwt")
)

type FilterTrackPropertyCompareOperation string

const (
	// The equal operation.
	FilterTrackPropertyCompareOperationEqual = FilterTrackPropertyCompareOperation("Equal")
	// The not equal operation.
	FilterTrackPropertyCompareOperationNotEqual = FilterTrackPropertyCompareOperation("NotEqual")
)

type FilterTrackPropertyType string

const (
	// The unknown track property type.
	FilterTrackPropertyTypeUnknown = FilterTrackPropertyType("Unknown")
	// The type.
	FilterTrackPropertyTypeType = FilterTrackPropertyType("Type")
	// The name.
	FilterTrackPropertyTypeName = FilterTrackPropertyType("Name")
	// The language.
	FilterTrackPropertyTypeLanguage = FilterTrackPropertyType("Language")
	// The fourCC.
	FilterTrackPropertyTypeFourCC = FilterTrackPropertyType("FourCC")
	// The bitrate.
	FilterTrackPropertyTypeBitrate = FilterTrackPropertyType("Bitrate")
)

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

type SecurityLevel string

const (
	// Represents a SecurityLevel that is unavailable in current API version.
	SecurityLevelUnknown = SecurityLevel("Unknown")
	// For clients under development or test. No protection against unauthorized use.
	SecurityLevelSL150 = SecurityLevel("SL150")
	// For hardened devices and applications consuming commercial content. Software or hardware protection.
	SecurityLevelSL2000 = SecurityLevel("SL2000")
	// For hardened devices only. Hardware protection.
	SecurityLevelSL3000 = SecurityLevel("SL3000")
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

type TrackPropertyCompareOperation string

const (
	// Unknown track property compare operation
	TrackPropertyCompareOperationUnknown = TrackPropertyCompareOperation("Unknown")
	// Equal operation
	TrackPropertyCompareOperationEqual = TrackPropertyCompareOperation("Equal")
)

type TrackPropertyType string

const (
	// Unknown track property
	TrackPropertyTypeUnknown = TrackPropertyType("Unknown")
	// Track FourCC
	TrackPropertyTypeFourCC = TrackPropertyType("FourCC")
)

type Visibility string

const (
	// The track is hidden to video player.
	VisibilityHidden = Visibility("Hidden")
	// The track is visible to video player.
	VisibilityVisible = Visibility("Visible")
)

func init() {
}
