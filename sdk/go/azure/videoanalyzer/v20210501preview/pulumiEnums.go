


package v20210501preview

type AccessPolicyEccAlgo string

const (
	// ES265
	AccessPolicyEccAlgoES256 = AccessPolicyEccAlgo("ES256")
	// ES384
	AccessPolicyEccAlgoES384 = AccessPolicyEccAlgo("ES384")
	// ES512
	AccessPolicyEccAlgoES512 = AccessPolicyEccAlgo("ES512")
)

type AccessPolicyRole string

const (
	// Reader role allows for read-only operations to be performed through the client APIs.
	AccessPolicyRoleReader = AccessPolicyRole("Reader")
)

type AccessPolicyRsaAlgo string

const (
	// RS256
	AccessPolicyRsaAlgoRS256 = AccessPolicyRsaAlgo("RS256")
	// RS384
	AccessPolicyRsaAlgoRS384 = AccessPolicyRsaAlgo("RS384")
	// RS512
	AccessPolicyRsaAlgoRS512 = AccessPolicyRsaAlgo("RS512")
)

type AccountEncryptionKeyType string

const (
	// The Account Key is encrypted with a System Key.
	AccountEncryptionKeyTypeSystemKey = AccountEncryptionKeyType("SystemKey")
	// The Account Key is encrypted with a Customer Key.
	AccountEncryptionKeyTypeCustomerKey = AccountEncryptionKeyType("CustomerKey")
)

func init() {
}
