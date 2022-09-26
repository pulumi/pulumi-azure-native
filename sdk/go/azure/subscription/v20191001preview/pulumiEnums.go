


package v20191001preview

type Workload string

const (
	WorkloadProduction = Workload("Production")
	WorkloadDevTest    = Workload("DevTest")
)

func init() {
}
