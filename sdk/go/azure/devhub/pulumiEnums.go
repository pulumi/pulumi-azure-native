


package devhub

type ManifestType string

const (
	// Repositories using helm
	ManifestTypeHelm = ManifestType("helm")
	// Repositories using kubernetes manifests
	ManifestTypeKube = ManifestType("kube")
)

func init() {
}
