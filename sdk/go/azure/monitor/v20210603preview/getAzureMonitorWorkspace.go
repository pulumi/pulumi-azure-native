


package v20210603preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzureMonitorWorkspace(ctx *pulumi.Context, args *LookupAzureMonitorWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupAzureMonitorWorkspaceResult, error) {
	var rv LookupAzureMonitorWorkspaceResult
	err := ctx.Invoke("azure-native:monitor/v20210603preview:getAzureMonitorWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureMonitorWorkspaceArgs struct {
	MonitoringAccountName string `pulumi:"monitoringAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupAzureMonitorWorkspaceResult struct {
	AccountId                string                                            `pulumi:"accountId"`
	DefaultIngestionSettings MonitoringAccountResponseDefaultIngestionSettings `pulumi:"defaultIngestionSettings"`
	Etag                     string                                            `pulumi:"etag"`
	Id                       string                                            `pulumi:"id"`
	Location                 string                                            `pulumi:"location"`
	Metrics                  MonitoringAccountResponseMetrics                  `pulumi:"metrics"`
	Name                     string                                            `pulumi:"name"`
	ProvisioningState        string                                            `pulumi:"provisioningState"`
	SystemData               SystemDataResponse                                `pulumi:"systemData"`
	Tags                     map[string]string                                 `pulumi:"tags"`
	Type                     string                                            `pulumi:"type"`
}

func LookupAzureMonitorWorkspaceOutput(ctx *pulumi.Context, args LookupAzureMonitorWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupAzureMonitorWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureMonitorWorkspaceResult, error) {
			args := v.(LookupAzureMonitorWorkspaceArgs)
			r, err := LookupAzureMonitorWorkspace(ctx, &args, opts...)
			var s LookupAzureMonitorWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureMonitorWorkspaceResultOutput)
}

type LookupAzureMonitorWorkspaceOutputArgs struct {
	MonitoringAccountName pulumi.StringInput `pulumi:"monitoringAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAzureMonitorWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureMonitorWorkspaceArgs)(nil)).Elem()
}


type LookupAzureMonitorWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupAzureMonitorWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureMonitorWorkspaceResult)(nil)).Elem()
}

func (o LookupAzureMonitorWorkspaceResultOutput) ToLookupAzureMonitorWorkspaceResultOutput() LookupAzureMonitorWorkspaceResultOutput {
	return o
}

func (o LookupAzureMonitorWorkspaceResultOutput) ToLookupAzureMonitorWorkspaceResultOutputWithContext(ctx context.Context) LookupAzureMonitorWorkspaceResultOutput {
	return o
}

func (o LookupAzureMonitorWorkspaceResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) DefaultIngestionSettings() MonitoringAccountResponseDefaultIngestionSettingsOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) MonitoringAccountResponseDefaultIngestionSettings {
		return v.DefaultIngestionSettings
	}).(MonitoringAccountResponseDefaultIngestionSettingsOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) Metrics() MonitoringAccountResponseMetricsOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) MonitoringAccountResponseMetrics { return v.Metrics }).(MonitoringAccountResponseMetricsOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAzureMonitorWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureMonitorWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureMonitorWorkspaceResultOutput{})
}
