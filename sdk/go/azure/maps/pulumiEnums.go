


package maps

type SigningKey string

const (
	SigningKeyPrimaryKey   = SigningKey("primaryKey")
	SigningKeySecondaryKey = SigningKey("secondaryKey")
)

func init() {
}
