


package v20200901

type Workload string

const (
	WorkloadProduction = Workload("Production")
	WorkloadDevTest    = Workload("DevTest")
)

func init() {
}
