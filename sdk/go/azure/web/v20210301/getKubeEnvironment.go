


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKubeEnvironment(ctx *pulumi.Context, args *LookupKubeEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupKubeEnvironmentResult, error) {
	var rv LookupKubeEnvironmentResult
	err := ctx.Invoke("azure-native:web/v20210301:getKubeEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKubeEnvironmentArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupKubeEnvironmentResult struct {
	AksResourceID               *string                             `pulumi:"aksResourceID"`
	AppLogsConfiguration        *AppLogsConfigurationResponse       `pulumi:"appLogsConfiguration"`
	ArcConfiguration            *ArcConfigurationResponse           `pulumi:"arcConfiguration"`
	ContainerAppsConfiguration  *ContainerAppsConfigurationResponse `pulumi:"containerAppsConfiguration"`
	DefaultDomain               string                              `pulumi:"defaultDomain"`
	DeploymentErrors            string                              `pulumi:"deploymentErrors"`
	EnvironmentType             *string                             `pulumi:"environmentType"`
	ExtendedLocation            *ExtendedLocationResponse           `pulumi:"extendedLocation"`
	Id                          string                              `pulumi:"id"`
	InternalLoadBalancerEnabled *bool                               `pulumi:"internalLoadBalancerEnabled"`
	Kind                        *string                             `pulumi:"kind"`
	Location                    string                              `pulumi:"location"`
	Name                        string                              `pulumi:"name"`
	ProvisioningState           string                              `pulumi:"provisioningState"`
	StaticIp                    *string                             `pulumi:"staticIp"`
	Tags                        map[string]string                   `pulumi:"tags"`
	Type                        string                              `pulumi:"type"`
}

func LookupKubeEnvironmentOutput(ctx *pulumi.Context, args LookupKubeEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupKubeEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKubeEnvironmentResult, error) {
			args := v.(LookupKubeEnvironmentArgs)
			r, err := LookupKubeEnvironment(ctx, &args, opts...)
			var s LookupKubeEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKubeEnvironmentResultOutput)
}

type LookupKubeEnvironmentOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKubeEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubeEnvironmentArgs)(nil)).Elem()
}


type LookupKubeEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupKubeEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubeEnvironmentResult)(nil)).Elem()
}

func (o LookupKubeEnvironmentResultOutput) ToLookupKubeEnvironmentResultOutput() LookupKubeEnvironmentResultOutput {
	return o
}

func (o LookupKubeEnvironmentResultOutput) ToLookupKubeEnvironmentResultOutputWithContext(ctx context.Context) LookupKubeEnvironmentResultOutput {
	return o
}

func (o LookupKubeEnvironmentResultOutput) AksResourceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *string { return v.AksResourceID }).(pulumi.StringPtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) AppLogsConfiguration() AppLogsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *AppLogsConfigurationResponse { return v.AppLogsConfiguration }).(AppLogsConfigurationResponsePtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) ArcConfiguration() ArcConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *ArcConfigurationResponse { return v.ArcConfiguration }).(ArcConfigurationResponsePtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) ContainerAppsConfiguration() ContainerAppsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *ContainerAppsConfigurationResponse {
		return v.ContainerAppsConfiguration
	}).(ContainerAppsConfigurationResponsePtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) DefaultDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) string { return v.DefaultDomain }).(pulumi.StringOutput)
}

func (o LookupKubeEnvironmentResultOutput) DeploymentErrors() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) string { return v.DeploymentErrors }).(pulumi.StringOutput)
}

func (o LookupKubeEnvironmentResultOutput) EnvironmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *string { return v.EnvironmentType }).(pulumi.StringPtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKubeEnvironmentResultOutput) InternalLoadBalancerEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *bool { return v.InternalLoadBalancerEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKubeEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKubeEnvironmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKubeEnvironmentResultOutput) StaticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) *string { return v.StaticIp }).(pulumi.StringPtrOutput)
}

func (o LookupKubeEnvironmentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupKubeEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubeEnvironmentResultOutput{})
}
