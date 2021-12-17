


package managednetwork

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
