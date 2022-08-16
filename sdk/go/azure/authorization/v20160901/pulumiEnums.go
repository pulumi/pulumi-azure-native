


package v20160901

type LockLevel string

const (
	LockLevelNotSpecified = LockLevel("NotSpecified")
	LockLevelCanNotDelete = LockLevel("CanNotDelete")
	LockLevelReadOnly     = LockLevel("ReadOnly")
)

func init() {
}
