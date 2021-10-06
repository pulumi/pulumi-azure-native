


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetguestDiagnosticsSetting(ctx *pulumi.Context, args *GetguestDiagnosticsSettingArgs, opts ...pulumi.InvokeOption) (*GetguestDiagnosticsSettingResult, error) {
	var rv GetguestDiagnosticsSettingResult
	err := ctx.Invoke("azure-native:insights/v20180601preview:getguestDiagnosticsSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetguestDiagnosticsSettingArgs struct {
	DiagnosticSettingsName string `pulumi:"diagnosticSettingsName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type GetguestDiagnosticsSettingResult struct {
	DataSources  []DataSourceResponse `pulumi:"dataSources"`
	Id           string               `pulumi:"id"`
	Location     string               `pulumi:"location"`
	Name         string               `pulumi:"name"`
	OsType       *string              `pulumi:"osType"`
	ProxySetting *string              `pulumi:"proxySetting"`
	Tags         map[string]string    `pulumi:"tags"`
	Type         string               `pulumi:"type"`
}
