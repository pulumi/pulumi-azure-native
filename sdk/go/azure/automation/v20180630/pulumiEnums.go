


package v20180630

type RunbookTypeEnum string

const (
	RunbookTypeEnumScript                  = RunbookTypeEnum("Script")
	RunbookTypeEnumGraph                   = RunbookTypeEnum("Graph")
	RunbookTypeEnumPowerShellWorkflow      = RunbookTypeEnum("PowerShellWorkflow")
	RunbookTypeEnumPowerShell              = RunbookTypeEnum("PowerShell")
	RunbookTypeEnumGraphPowerShellWorkflow = RunbookTypeEnum("GraphPowerShellWorkflow")
	RunbookTypeEnumGraphPowerShell         = RunbookTypeEnum("GraphPowerShell")
	RunbookTypeEnumPython2                 = RunbookTypeEnum("Python2")
	RunbookTypeEnumPython3                 = RunbookTypeEnum("Python3")
)

func init() {
}
