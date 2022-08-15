


package v20191101preview

type EnableHelmOperator string

const (
	EnableHelmOperatorTrue  = EnableHelmOperator("true")
	EnableHelmOperatorFalse = EnableHelmOperator("false")
)

type OperatorScope string

const (
	OperatorScopeCluster   = OperatorScope("cluster")
	OperatorScopeNamespace = OperatorScope("namespace")
)

type OperatorType string

const (
	OperatorTypeFlux = OperatorType("Flux")
)

func init() {
}
