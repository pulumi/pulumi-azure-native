


package v20200207preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSapMonitor(ctx *pulumi.Context, args *LookupSapMonitorArgs, opts ...pulumi.InvokeOption) (*LookupSapMonitorResult, error) {
	var rv LookupSapMonitorResult
	err := ctx.Invoke("azure-native:hanaonazure/v20200207preview:getSapMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSapMonitorArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SapMonitorName    string `pulumi:"sapMonitorName"`
}


type LookupSapMonitorResult struct {
	EnableCustomerAnalytics        *bool             `pulumi:"enableCustomerAnalytics"`
	Id                             string            `pulumi:"id"`
	Location                       string            `pulumi:"location"`
	LogAnalyticsWorkspaceArmId     *string           `pulumi:"logAnalyticsWorkspaceArmId"`
	LogAnalyticsWorkspaceId        *string           `pulumi:"logAnalyticsWorkspaceId"`
	LogAnalyticsWorkspaceSharedKey *string           `pulumi:"logAnalyticsWorkspaceSharedKey"`
	ManagedResourceGroupName       string            `pulumi:"managedResourceGroupName"`
	MonitorSubnet                  *string           `pulumi:"monitorSubnet"`
	Name                           string            `pulumi:"name"`
	ProvisioningState              string            `pulumi:"provisioningState"`
	SapMonitorCollectorVersion     string            `pulumi:"sapMonitorCollectorVersion"`
	Tags                           map[string]string `pulumi:"tags"`
	Type                           string            `pulumi:"type"`
}

func LookupSapMonitorOutput(ctx *pulumi.Context, args LookupSapMonitorOutputArgs, opts ...pulumi.InvokeOption) LookupSapMonitorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSapMonitorResult, error) {
			args := v.(LookupSapMonitorArgs)
			r, err := LookupSapMonitor(ctx, &args, opts...)
			var s LookupSapMonitorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSapMonitorResultOutput)
}

type LookupSapMonitorOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SapMonitorName    pulumi.StringInput `pulumi:"sapMonitorName"`
}

func (LookupSapMonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSapMonitorArgs)(nil)).Elem()
}


type LookupSapMonitorResultOutput struct{ *pulumi.OutputState }

func (LookupSapMonitorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSapMonitorResult)(nil)).Elem()
}

func (o LookupSapMonitorResultOutput) ToLookupSapMonitorResultOutput() LookupSapMonitorResultOutput {
	return o
}

func (o LookupSapMonitorResultOutput) ToLookupSapMonitorResultOutputWithContext(ctx context.Context) LookupSapMonitorResultOutput {
	return o
}

func (o LookupSapMonitorResultOutput) EnableCustomerAnalytics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *bool { return v.EnableCustomerAnalytics }).(pulumi.BoolPtrOutput)
}

func (o LookupSapMonitorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSapMonitorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSapMonitorResultOutput) LogAnalyticsWorkspaceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *string { return v.LogAnalyticsWorkspaceArmId }).(pulumi.StringPtrOutput)
}

func (o LookupSapMonitorResultOutput) LogAnalyticsWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *string { return v.LogAnalyticsWorkspaceId }).(pulumi.StringPtrOutput)
}

func (o LookupSapMonitorResultOutput) LogAnalyticsWorkspaceSharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *string { return v.LogAnalyticsWorkspaceSharedKey }).(pulumi.StringPtrOutput)
}

func (o LookupSapMonitorResultOutput) ManagedResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.ManagedResourceGroupName }).(pulumi.StringOutput)
}

func (o LookupSapMonitorResultOutput) MonitorSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) *string { return v.MonitorSubnet }).(pulumi.StringPtrOutput)
}

func (o LookupSapMonitorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSapMonitorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSapMonitorResultOutput) SapMonitorCollectorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.SapMonitorCollectorVersion }).(pulumi.StringOutput)
}

func (o LookupSapMonitorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSapMonitorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapMonitorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSapMonitorResultOutput{})
}
