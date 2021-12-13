


package subscription

type Workload string

const (
	WorkloadProduction = Workload("Production")
	WorkloadDevTest    = Workload("DevTest")
)

func init() {
}
