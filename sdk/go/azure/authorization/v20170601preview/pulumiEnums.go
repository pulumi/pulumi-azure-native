


package v20170601preview

type PolicyType string

const (
	PolicyTypeNotSpecified = PolicyType("NotSpecified")
	PolicyTypeBuiltIn      = PolicyType("BuiltIn")
	PolicyTypeCustom       = PolicyType("Custom")
)

func init() {
}
