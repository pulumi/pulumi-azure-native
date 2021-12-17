


package hybridconnectivity

type CreatedByType string

const (
	CreatedByTypeUser            = CreatedByType("User")
	CreatedByTypeApplication     = CreatedByType("Application")
	CreatedByTypeManagedIdentity = CreatedByType("ManagedIdentity")
	CreatedByTypeKey             = CreatedByType("Key")
)

type Type string

const (
	TypeDefault = Type("default")
	TypeCustom  = Type("custom")
)

func init() {
}
