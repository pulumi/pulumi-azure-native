


package v20200501

type LockLevel string

const (
	LockLevelNotSpecified = LockLevel("NotSpecified")
	LockLevelCanNotDelete = LockLevel("CanNotDelete")
	LockLevelReadOnly     = LockLevel("ReadOnly")
)

type PublicNetworkAccessOptions string

const (
	PublicNetworkAccessOptionsEnabled  = PublicNetworkAccessOptions("Enabled")
	PublicNetworkAccessOptionsDisabled = PublicNetworkAccessOptions("Disabled")
)

func init() {
}
