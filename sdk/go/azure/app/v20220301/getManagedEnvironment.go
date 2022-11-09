


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedEnvironment(ctx *pulumi.Context, args *LookupManagedEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupManagedEnvironmentResult, error) {
	var rv LookupManagedEnvironmentResult
	err := ctx.Invoke("azure-native:app/v20220301:getManagedEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedEnvironmentArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagedEnvironmentResult struct {
	AppLogsConfiguration     *AppLogsConfigurationResponse `pulumi:"appLogsConfiguration"`
	DaprAIConnectionString   *string                       `pulumi:"daprAIConnectionString"`
	DaprAIInstrumentationKey *string                       `pulumi:"daprAIInstrumentationKey"`
	DefaultDomain            string                        `pulumi:"defaultDomain"`
	DeploymentErrors         string                        `pulumi:"deploymentErrors"`
	Id                       string                        `pulumi:"id"`
	Location                 string                        `pulumi:"location"`
	Name                     string                        `pulumi:"name"`
	ProvisioningState        string                        `pulumi:"provisioningState"`
	StaticIp                 string                        `pulumi:"staticIp"`
	SystemData               SystemDataResponse            `pulumi:"systemData"`
	Tags                     map[string]string             `pulumi:"tags"`
	Type                     string                        `pulumi:"type"`
	VnetConfiguration        *VnetConfigurationResponse    `pulumi:"vnetConfiguration"`
	ZoneRedundant            *bool                         `pulumi:"zoneRedundant"`
}

func LookupManagedEnvironmentOutput(ctx *pulumi.Context, args LookupManagedEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupManagedEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedEnvironmentResult, error) {
			args := v.(LookupManagedEnvironmentArgs)
			r, err := LookupManagedEnvironment(ctx, &args, opts...)
			var s LookupManagedEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedEnvironmentResultOutput)
}

type LookupManagedEnvironmentOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedEnvironmentArgs)(nil)).Elem()
}


type LookupManagedEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupManagedEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedEnvironmentResult)(nil)).Elem()
}

func (o LookupManagedEnvironmentResultOutput) ToLookupManagedEnvironmentResultOutput() LookupManagedEnvironmentResultOutput {
	return o
}

func (o LookupManagedEnvironmentResultOutput) ToLookupManagedEnvironmentResultOutputWithContext(ctx context.Context) LookupManagedEnvironmentResultOutput {
	return o
}

func (o LookupManagedEnvironmentResultOutput) AppLogsConfiguration() AppLogsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *AppLogsConfigurationResponse { return v.AppLogsConfiguration }).(AppLogsConfigurationResponsePtrOutput)
}

func (o LookupManagedEnvironmentResultOutput) DaprAIConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *string { return v.DaprAIConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupManagedEnvironmentResultOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *string { return v.DaprAIInstrumentationKey }).(pulumi.StringPtrOutput)
}

func (o LookupManagedEnvironmentResultOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.DefaultDomain }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentResultOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.DeploymentErrors }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentResultOutput) StaticIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.StaticIp }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupManagedEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupManagedEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupManagedEnvironmentResultOutput) VnetConfiguration() VnetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *VnetConfigurationResponse { return v.VnetConfiguration }).(VnetConfigurationResponsePtrOutput)
}

func (o LookupManagedEnvironmentResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagedEnvironmentResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedEnvironmentResultOutput{})
}
