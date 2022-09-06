


package v20150101

type LockLevel string

const (
	LockLevelNotSpecified = LockLevel("NotSpecified")
	LockLevelCanNotDelete = LockLevel("CanNotDelete")
	LockLevelReadOnly     = LockLevel("ReadOnly")
)

func init() {
}
