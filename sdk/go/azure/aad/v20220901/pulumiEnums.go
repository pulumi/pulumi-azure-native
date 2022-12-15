


package v20220901

type ChannelBinding string

const (
	ChannelBindingEnabled  = ChannelBinding("Enabled")
	ChannelBindingDisabled = ChannelBinding("Disabled")
)

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

type KerberosArmoring string

const (
	KerberosArmoringEnabled  = KerberosArmoring("Enabled")
	KerberosArmoringDisabled = KerberosArmoring("Disabled")
)

type KerberosRc4Encryption string

const (
	KerberosRc4EncryptionEnabled  = KerberosRc4Encryption("Enabled")
	KerberosRc4EncryptionDisabled = KerberosRc4Encryption("Disabled")
)

type LdapSigning string

const (
	LdapSigningEnabled  = LdapSigning("Enabled")
	LdapSigningDisabled = LdapSigning("Disabled")
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

type Status string

const (
	StatusNone    = Status("None")
	StatusRunning = Status("Running")
	StatusOK      = Status("OK")
	StatusFailure = Status("Failure")
	StatusWarning = Status("Warning")
	StatusSkipped = Status("Skipped")
)

type SyncKerberosPasswords string

const (
	SyncKerberosPasswordsEnabled  = SyncKerberosPasswords("Enabled")
	SyncKerberosPasswordsDisabled = SyncKerberosPasswords("Disabled")
)

type SyncNtlmPasswords string

const (
	SyncNtlmPasswordsEnabled  = SyncNtlmPasswords("Enabled")
	SyncNtlmPasswordsDisabled = SyncNtlmPasswords("Disabled")
)

type SyncOnPremPasswords string

const (
	SyncOnPremPasswordsEnabled  = SyncOnPremPasswords("Enabled")
	SyncOnPremPasswordsDisabled = SyncOnPremPasswords("Disabled")
)

type SyncScope string

const (
	SyncScopeAll       = SyncScope("All")
	SyncScopeCloudOnly = SyncScope("CloudOnly")
)

type TlsV1 string

const (
	TlsV1Enabled  = TlsV1("Enabled")
	TlsV1Disabled = TlsV1("Disabled")
)

func init() {
}
