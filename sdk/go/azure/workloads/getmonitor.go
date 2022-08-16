


package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Getmonitor(ctx *pulumi.Context, args *GetmonitorArgs, opts ...pulumi.InvokeOption) (*GetmonitorResult, error) {
	var rv GetmonitorResult
	err := ctx.Invoke("azure-native:workloads:getmonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetmonitorArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetmonitorResult struct {
	AppLocation                       *string                              `pulumi:"appLocation"`
	Errors                            MonitorPropertiesResponseErrors      `pulumi:"errors"`
	Id                                string                               `pulumi:"id"`
	Identity                          *UserAssignedServiceIdentityResponse `pulumi:"identity"`
	Location                          string                               `pulumi:"location"`
	LogAnalyticsWorkspaceArmId        *string                              `pulumi:"logAnalyticsWorkspaceArmId"`
	ManagedResourceGroupConfiguration *ManagedRGConfigurationResponse      `pulumi:"managedResourceGroupConfiguration"`
	MonitorSubnet                     *string                              `pulumi:"monitorSubnet"`
	MsiArmId                          string                               `pulumi:"msiArmId"`
	Name                              string                               `pulumi:"name"`
	ProvisioningState                 string                               `pulumi:"provisioningState"`
	RoutingPreference                 *string                              `pulumi:"routingPreference"`
	SystemData                        SystemDataResponse                   `pulumi:"systemData"`
	Tags                              map[string]string                    `pulumi:"tags"`
	Type                              string                               `pulumi:"type"`
	ZoneRedundancyPreference          *string                              `pulumi:"zoneRedundancyPreference"`
}

func GetmonitorOutput(ctx *pulumi.Context, args GetmonitorOutputArgs, opts ...pulumi.InvokeOption) GetmonitorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetmonitorResult, error) {
			args := v.(GetmonitorArgs)
			r, err := Getmonitor(ctx, &args, opts...)
			var s GetmonitorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetmonitorResultOutput)
}

type GetmonitorOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetmonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmonitorArgs)(nil)).Elem()
}


type GetmonitorResultOutput struct{ *pulumi.OutputState }

func (GetmonitorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmonitorResult)(nil)).Elem()
}

func (o GetmonitorResultOutput) ToGetmonitorResultOutput() GetmonitorResultOutput {
	return o
}

func (o GetmonitorResultOutput) ToGetmonitorResultOutputWithContext(ctx context.Context) GetmonitorResultOutput {
	return o
}

func (o GetmonitorResultOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmonitorResult) *string { return v.AppLocation }).(pulumi.StringPtrOutput)
}

func (o GetmonitorResultOutput) Errors() MonitorPropertiesResponseErrorsOutput {
	return o.ApplyT(func(v GetmonitorResult) MonitorPropertiesResponseErrors { return v.Errors }).(MonitorPropertiesResponseErrorsOutput)
}

func (o GetmonitorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetmonitorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetmonitorResultOutput) Identity() UserAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v GetmonitorResult) *UserAssignedServiceIdentityResponse { return v.Identity }).(UserAssignedServiceIdentityResponsePtrOutput)
}

func (o GetmonitorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetmonitorResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetmonitorResultOutput) LogAnalyticsWorkspaceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmonitorResult) *string { return v.LogAnalyticsWorkspaceArmId }).(pulumi.StringPtrOutput)
}

func (o GetmonitorResultOutput) ManagedResourceGroupConfiguration() ManagedRGConfigurationResponsePtrOutput {
	return o.ApplyT(func(v GetmonitorResult) *ManagedRGConfigurationResponse { return v.ManagedResourceGroupConfiguration }).(ManagedRGConfigurationResponsePtrOutput)
}

func (o GetmonitorResultOutput) MonitorSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmonitorResult) *string { return v.MonitorSubnet }).(pulumi.StringPtrOutput)
}

func (o GetmonitorResultOutput) MsiArmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetmonitorResult) string { return v.MsiArmId }).(pulumi.StringOutput)
}

func (o GetmonitorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetmonitorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetmonitorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetmonitorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetmonitorResultOutput) RoutingPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmonitorResult) *string { return v.RoutingPreference }).(pulumi.StringPtrOutput)
}

func (o GetmonitorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetmonitorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetmonitorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetmonitorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetmonitorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetmonitorResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetmonitorResultOutput) ZoneRedundancyPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmonitorResult) *string { return v.ZoneRedundancyPreference }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetmonitorResultOutput{})
}
