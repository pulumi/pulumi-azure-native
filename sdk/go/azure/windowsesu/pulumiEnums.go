


package windowsesu

type OsType string

const (
	OsTypeWindows7            = OsType("Windows7")
	OsTypeWindowsServer2008   = OsType("WindowsServer2008")
	OsTypeWindowsServer2008R2 = OsType("WindowsServer2008R2")
)

type SupportType string

const (
	SupportTypeSupplementalServicing = SupportType("SupplementalServicing")
	SupportTypePremiumAssurance      = SupportType("PremiumAssurance")
)

func init() {
}
