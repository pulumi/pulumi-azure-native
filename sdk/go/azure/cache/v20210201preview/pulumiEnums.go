


package v20210201preview

type AofFrequency string

const (
	AofFrequency_1s    = AofFrequency("1s")
	AofFrequencyAlways = AofFrequency("always")
)

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

type RdbFrequency string

const (
	RdbFrequency_1h  = RdbFrequency("1h")
	RdbFrequency_6h  = RdbFrequency("6h")
	RdbFrequency_12h = RdbFrequency("12h")
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

type TlsVersion string

const (
	TlsVersion_1_0 = TlsVersion("1.0")
	TlsVersion_1_1 = TlsVersion("1.1")
	TlsVersion_1_2 = TlsVersion("1.2")
)

func init() {
}
