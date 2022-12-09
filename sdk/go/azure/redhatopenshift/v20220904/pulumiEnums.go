


package v20220904

type EncryptionAtHost string

const (
	EncryptionAtHostDisabled = EncryptionAtHost("Disabled")
	EncryptionAtHostEnabled  = EncryptionAtHost("Enabled")
)

type FipsValidatedModules string

const (
	FipsValidatedModulesDisabled = FipsValidatedModules("Disabled")
	FipsValidatedModulesEnabled  = FipsValidatedModules("Enabled")
)

type Visibility string

const (
	VisibilityPrivate = Visibility("Private")
	VisibilityPublic  = Visibility("Public")
)

func init() {
}
