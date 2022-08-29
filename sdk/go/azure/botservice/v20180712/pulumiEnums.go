


package v20180712

type EnterpriseChannelNodeState string

const (
	EnterpriseChannelNodeStateCreating     = EnterpriseChannelNodeState("Creating")
	EnterpriseChannelNodeStateCreateFailed = EnterpriseChannelNodeState("CreateFailed")
	EnterpriseChannelNodeStateStarted      = EnterpriseChannelNodeState("Started")
	EnterpriseChannelNodeStateStarting     = EnterpriseChannelNodeState("Starting")
	EnterpriseChannelNodeStateStartFailed  = EnterpriseChannelNodeState("StartFailed")
	EnterpriseChannelNodeStateStopped      = EnterpriseChannelNodeState("Stopped")
	EnterpriseChannelNodeStateStopping     = EnterpriseChannelNodeState("Stopping")
	EnterpriseChannelNodeStateStopFailed   = EnterpriseChannelNodeState("StopFailed")
	EnterpriseChannelNodeStateDeleting     = EnterpriseChannelNodeState("Deleting")
	EnterpriseChannelNodeStateDeleteFailed = EnterpriseChannelNodeState("DeleteFailed")
)

type EnterpriseChannelStateEnum string

const (
	EnterpriseChannelStateEnumCreating     = EnterpriseChannelStateEnum("Creating")
	EnterpriseChannelStateEnumCreateFailed = EnterpriseChannelStateEnum("CreateFailed")
	EnterpriseChannelStateEnumStarted      = EnterpriseChannelStateEnum("Started")
	EnterpriseChannelStateEnumStarting     = EnterpriseChannelStateEnum("Starting")
	EnterpriseChannelStateEnumStartFailed  = EnterpriseChannelStateEnum("StartFailed")
	EnterpriseChannelStateEnumStopped      = EnterpriseChannelStateEnum("Stopped")
	EnterpriseChannelStateEnumStopping     = EnterpriseChannelStateEnum("Stopping")
	EnterpriseChannelStateEnumStopFailed   = EnterpriseChannelStateEnum("StopFailed")
	EnterpriseChannelStateEnumDeleting     = EnterpriseChannelStateEnum("Deleting")
	EnterpriseChannelStateEnumDeleteFailed = EnterpriseChannelStateEnum("DeleteFailed")
)

type Kind string

const (
	KindSdk      = Kind("sdk")
	KindDesigner = Kind("designer")
	KindBot      = Kind("bot")
	KindFunction = Kind("function")
)

type SkuName string

const (
	SkuNameF0 = SkuName("F0")
	SkuNameS1 = SkuName("S1")
)

func init() {
}
