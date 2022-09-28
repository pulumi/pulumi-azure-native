


package v20211001

type Workload string

const (
	WorkloadProduction = Workload("Production")
	WorkloadDevTest    = Workload("DevTest")
)

func init() {
}
