


package v20210301

type OperatorScopeType string

const (
	OperatorScopeTypeCluster   = OperatorScopeType("cluster")
	OperatorScopeTypeNamespace = OperatorScopeType("namespace")
)

type OperatorType string

const (
	OperatorTypeFlux = OperatorType("Flux")
)

func init() {
}
