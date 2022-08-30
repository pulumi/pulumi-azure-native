


package v20180301

type PolicyMode string

const (
	PolicyModeNotSpecified = PolicyMode("NotSpecified")
	PolicyModeIndexed      = PolicyMode("Indexed")
	PolicyModeAll          = PolicyMode("All")
)

type PolicyType string

const (
	PolicyTypeNotSpecified = PolicyType("NotSpecified")
	PolicyTypeBuiltIn      = PolicyType("BuiltIn")
	PolicyTypeCustom       = PolicyType("Custom")
)

func init() {
}
