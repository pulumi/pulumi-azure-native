


package v20220110preview

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
)

type Status string

const (
	StatusSucceeded = Status("Succeeded")
	StatusLaunching = Status("Launching")
	StatusUpdating  = Status("Updating")
	StatusDeleting  = Status("Deleting")
	StatusDeleted   = Status("Deleted")
	StatusFailed    = Status("Failed")
)

func init() {
}
