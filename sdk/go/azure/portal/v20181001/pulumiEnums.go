


package v20181001

type FontSize string

const (
	FontSizeNotSpecified = FontSize("NotSpecified")
	FontSizeSmall        = FontSize("Small")
	FontSizeMedium       = FontSize("Medium")
	FontSizeLarge        = FontSize("Large")
)

type FontStyle string

const (
	FontStyleNotSpecified = FontStyle("NotSpecified")
	FontStyleMonospace    = FontStyle("Monospace")
	FontStyleCourier      = FontStyle("Courier")
)

type OsType string

const (
	OsTypeWindows = OsType("Windows")
	OsTypeLinux   = OsType("Linux")
)

type ProvisioningState string

const (
	ProvisioningStateNotSpecified = ProvisioningState("NotSpecified")
	ProvisioningStateAccepted     = ProvisioningState("Accepted")
	ProvisioningStatePending      = ProvisioningState("Pending")
	ProvisioningStateUpdating     = ProvisioningState("Updating")
	ProvisioningStateCreating     = ProvisioningState("Creating")
	ProvisioningStateRepairing    = ProvisioningState("Repairing")
	ProvisioningStateFailed       = ProvisioningState("Failed")
	ProvisioningStateCanceled     = ProvisioningState("Canceled")
	ProvisioningStateSucceeded    = ProvisioningState("Succeeded")
)

type ShellType string

const (
	ShellTypeBash       = ShellType("bash")
	ShellTypePwsh       = ShellType("pwsh")
	ShellTypePowershell = ShellType("powershell")
)

func init() {
}
