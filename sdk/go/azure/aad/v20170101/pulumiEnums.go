


package v20170101

type ExternalAccess string

const (
	ExternalAccessEnabled  = ExternalAccess("Enabled")
	ExternalAccessDisabled = ExternalAccess("Disabled")
)

type FilteredSync string

const (
	FilteredSyncEnabled  = FilteredSync("Enabled")
	FilteredSyncDisabled = FilteredSync("Disabled")
)

type Ldaps string

const (
	LdapsEnabled  = Ldaps("Enabled")
	LdapsDisabled = Ldaps("Disabled")
)

type NotifyDcAdmins string

const (
	NotifyDcAdminsEnabled  = NotifyDcAdmins("Enabled")
	NotifyDcAdminsDisabled = NotifyDcAdmins("Disabled")
)

type NotifyGlobalAdmins string

const (
	NotifyGlobalAdminsEnabled  = NotifyGlobalAdmins("Enabled")
	NotifyGlobalAdminsDisabled = NotifyGlobalAdmins("Disabled")
)

type NtlmV1 string

const (
	NtlmV1Enabled  = NtlmV1("Enabled")
	NtlmV1Disabled = NtlmV1("Disabled")
)

type SyncNtlmPasswords string

const (
	SyncNtlmPasswordsEnabled  = SyncNtlmPasswords("Enabled")
	SyncNtlmPasswordsDisabled = SyncNtlmPasswords("Disabled")
)

type TlsV1 string

const (
	TlsV1Enabled  = TlsV1("Enabled")
	TlsV1Disabled = TlsV1("Disabled")
)

func init() {
}
