


package v20200901

type EndpointType string

const (
	EndpointTypeSrc = EndpointType("src")
	EndpointTypeDst = EndpointType("dst")
)

type QosType string

const (
	// qos type Auto
	QosTypeAuto = QosType("Auto")
	// qos type Manual
	QosTypeManual = QosType("Manual")
)

type ReplicationSchedule string

const (
	ReplicationSchedule_10minutely = ReplicationSchedule("_10minutely")
	ReplicationScheduleHourly      = ReplicationSchedule("hourly")
	ReplicationScheduleDaily       = ReplicationSchedule("daily")
)

type SecurityStyle string

const (
	SecurityStyleNtfs = SecurityStyle("ntfs")
	SecurityStyleUnix = SecurityStyle("unix")
)

type ServiceLevel string

const (
	// Standard service level
	ServiceLevelStandard = ServiceLevel("Standard")
	// Premium service level
	ServiceLevelPremium = ServiceLevel("Premium")
	// Ultra service level
	ServiceLevelUltra = ServiceLevel("Ultra")
)

func init() {
}
