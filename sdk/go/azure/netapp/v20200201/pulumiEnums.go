


package v20200201

type EndpointType string

const (
	EndpointTypeSrc = EndpointType("src")
	EndpointTypeDst = EndpointType("dst")
)

type PoolServiceLevel string

const (
	// Standard service level
	PoolServiceLevelStandard = PoolServiceLevel("Standard")
	// Premium service level
	PoolServiceLevelPremium = PoolServiceLevel("Premium")
	// Ultra service level
	PoolServiceLevelUltra = PoolServiceLevel("Ultra")
)

type ReplicationSchedule string

const (
	ReplicationSchedule_10minutely = ReplicationSchedule("_10minutely")
	ReplicationScheduleHourly      = ReplicationSchedule("hourly")
	ReplicationScheduleDaily       = ReplicationSchedule("daily")
)

type VolumeServiceLevel string

const (
	// Standard service level
	VolumeServiceLevelStandard = VolumeServiceLevel("Standard")
	// Premium service level
	VolumeServiceLevelPremium = VolumeServiceLevel("Premium")
	// Ultra service level
	VolumeServiceLevelUltra = VolumeServiceLevel("Ultra")
)

func init() {
}
