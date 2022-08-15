


package v20220513

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
