


package v20180630preview

type ActionAfterReboot string

const (
	ActionAfterRebootContinueConfiguration = ActionAfterReboot("ContinueConfiguration")
	ActionAfterRebootStopConfiguration     = ActionAfterReboot("StopConfiguration")
)

type ConfigurationMode string

const (
	ConfigurationModeApplyOnly           = ConfigurationMode("ApplyOnly")
	ConfigurationModeApplyAndMonitor     = ConfigurationMode("ApplyAndMonitor")
	ConfigurationModeApplyAndAutoCorrect = ConfigurationMode("ApplyAndAutoCorrect")
)

type Kind string

const (
	KindDSC = Kind("DSC")
)

func init() {
}
