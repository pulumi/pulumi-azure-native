


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectedEnvironment(ctx *pulumi.Context, args *LookupConnectedEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupConnectedEnvironmentResult, error) {
	var rv LookupConnectedEnvironmentResult
	err := ctx.Invoke("azure-native:app/v20220601preview:getConnectedEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectedEnvironmentArgs struct {
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type LookupConnectedEnvironmentResult struct {
	CustomDomainConfiguration *CustomDomainConfigurationResponse `pulumi:"customDomainConfiguration"`
	DaprAIConnectionString    *string                            `pulumi:"daprAIConnectionString"`
	DefaultDomain             string                             `pulumi:"defaultDomain"`
	DeploymentErrors          string                             `pulumi:"deploymentErrors"`
	ExtendedLocation          *ExtendedLocationResponse          `pulumi:"extendedLocation"`
	Id                        string                             `pulumi:"id"`
	Location                  string                             `pulumi:"location"`
	Name                      string                             `pulumi:"name"`
	ProvisioningState         string                             `pulumi:"provisioningState"`
	StaticIp                  *string                            `pulumi:"staticIp"`
	SystemData                SystemDataResponse                 `pulumi:"systemData"`
	Tags                      map[string]string                  `pulumi:"tags"`
	Type                      string                             `pulumi:"type"`
}

func LookupConnectedEnvironmentOutput(ctx *pulumi.Context, args LookupConnectedEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedEnvironmentResult, error) {
			args := v.(LookupConnectedEnvironmentArgs)
			r, err := LookupConnectedEnvironment(ctx, &args, opts...)
			var s LookupConnectedEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedEnvironmentResultOutput)
}

type LookupConnectedEnvironmentOutputArgs struct {
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentArgs)(nil)).Elem()
}


type LookupConnectedEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentResult)(nil)).Elem()
}

func (o LookupConnectedEnvironmentResultOutput) ToLookupConnectedEnvironmentResultOutput() LookupConnectedEnvironmentResultOutput {
	return o
}

func (o LookupConnectedEnvironmentResultOutput) ToLookupConnectedEnvironmentResultOutputWithContext(ctx context.Context) LookupConnectedEnvironmentResultOutput {
	return o
}

func (o LookupConnectedEnvironmentResultOutput) CustomDomainConfiguration() CustomDomainConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) *CustomDomainConfigurationResponse {
		return v.CustomDomainConfiguration
	}).(CustomDomainConfigurationResponsePtrOutput)
}

func (o LookupConnectedEnvironmentResultOutput) DaprAIConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) *string { return v.DaprAIConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupConnectedEnvironmentResultOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) string { return v.DefaultDomain }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentResultOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) string { return v.DeploymentErrors }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupConnectedEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupConnectedEnvironmentResultOutput) StaticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) *string { return v.StaticIp }).(pulumi.StringPtrOutput)
}

func (o LookupConnectedEnvironmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectedEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupConnectedEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedEnvironmentResultOutput{})
}
