


package storagemover

type CopyMode string

const (
	CopyModeAdditive = CopyMode("Additive")
	CopyModeMirror   = CopyMode("Mirror")
)

type EndpointType string

const (
	EndpointTypeAzureStorageBlobContainer = EndpointType("AzureStorageBlobContainer")
	EndpointTypeNfsMount                  = EndpointType("NfsMount")
)

type NfsVersion string

const (
	NfsVersionNFSauto = NfsVersion("NFSauto")
	NfsVersionNFSv3   = NfsVersion("NFSv3")
	NfsVersionNFSv4   = NfsVersion("NFSv4")
)

func init() {
}
