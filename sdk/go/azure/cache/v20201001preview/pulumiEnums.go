


package v20201001preview

type ClusteringPolicy string

const (
	ClusteringPolicyEnterpriseCluster = ClusteringPolicy("EnterpriseCluster")
	ClusteringPolicyOSSCluster        = ClusteringPolicy("OSSCluster")
)

type EvictionPolicy string

const (
	EvictionPolicyAllKeysLFU     = EvictionPolicy("AllKeysLFU")
	EvictionPolicyAllKeysLRU     = EvictionPolicy("AllKeysLRU")
	EvictionPolicyAllKeysRandom  = EvictionPolicy("AllKeysRandom")
	EvictionPolicyVolatileLRU    = EvictionPolicy("VolatileLRU")
	EvictionPolicyVolatileLFU    = EvictionPolicy("VolatileLFU")
	EvictionPolicyVolatileTTL    = EvictionPolicy("VolatileTTL")
	EvictionPolicyVolatileRandom = EvictionPolicy("VolatileRandom")
	EvictionPolicyNoEviction     = EvictionPolicy("NoEviction")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

type Protocol string

const (
	ProtocolEncrypted = Protocol("Encrypted")
	ProtocolPlaintext = Protocol("Plaintext")
)

type SkuName string

const (
	SkuName_Enterprise_E10        = SkuName("Enterprise_E10")
	SkuName_Enterprise_E20        = SkuName("Enterprise_E20")
	SkuName_Enterprise_E50        = SkuName("Enterprise_E50")
	SkuName_Enterprise_E100       = SkuName("Enterprise_E100")
	SkuName_EnterpriseFlash_F300  = SkuName("EnterpriseFlash_F300")
	SkuName_EnterpriseFlash_F700  = SkuName("EnterpriseFlash_F700")
	SkuName_EnterpriseFlash_F1500 = SkuName("EnterpriseFlash_F1500")
)

func init() {
}
