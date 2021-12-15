


package v20181101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDiagnosticLogsConfiguration(ctx *pulumi.Context, args *LookupWebAppDiagnosticLogsConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDiagnosticLogsConfigurationResult, error) {
	var rv LookupWebAppDiagnosticLogsConfigurationResult
	err := ctx.Invoke("azure-native:web/v20181101:getWebAppDiagnosticLogsConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebAppDiagnosticLogsConfigurationArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppDiagnosticLogsConfigurationResult struct {
	ApplicationLogs       *ApplicationLogsConfigResponse `pulumi:"applicationLogs"`
	DetailedErrorMessages *EnabledConfigResponse         `pulumi:"detailedErrorMessages"`
	FailedRequestsTracing *EnabledConfigResponse         `pulumi:"failedRequestsTracing"`
	HttpLogs              *HttpLogsConfigResponse        `pulumi:"httpLogs"`
	Id                    string                         `pulumi:"id"`
	Kind                  *string                        `pulumi:"kind"`
	Name                  string                         `pulumi:"name"`
	Type                  string                         `pulumi:"type"`
}


func (val *LookupWebAppDiagnosticLogsConfigurationResult) Defaults() *LookupWebAppDiagnosticLogsConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ApplicationLogs = tmp.ApplicationLogs.Defaults()

	return &tmp
}
