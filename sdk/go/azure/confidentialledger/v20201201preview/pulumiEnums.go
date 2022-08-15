


package v20201201preview

type LedgerRoleName string

const (
	LedgerRoleNameReader        = LedgerRoleName("Reader")
	LedgerRoleNameContributor   = LedgerRoleName("Contributor")
	LedgerRoleNameAdministrator = LedgerRoleName("Administrator")
)

type LedgerType string

const (
	LedgerTypeUnknown = LedgerType("Unknown")
	LedgerTypePublic  = LedgerType("Public")
	LedgerTypePrivate = LedgerType("Private")
)

func init() {
}
