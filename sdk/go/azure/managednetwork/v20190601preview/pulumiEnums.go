


package v20190601preview

type Kind string

const (
	KindConnectivity = Kind("Connectivity")
)

type Type string

const (
	TypeHubAndSpokeTopology = Type("HubAndSpokeTopology")
	TypeMeshTopology        = Type("MeshTopology")
)

func init() {
}
