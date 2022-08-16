


package v20140401preview

type AccountResourceRequestOperationType string

const (
	AccountResourceRequestOperationTypeUnknown = AccountResourceRequestOperationType("unknown")
	AccountResourceRequestOperationTypeCreate  = AccountResourceRequestOperationType("create")
	AccountResourceRequestOperationTypeUpdate  = AccountResourceRequestOperationType("update")
	AccountResourceRequestOperationTypeLink    = AccountResourceRequestOperationType("link")
)

func init() {
}
