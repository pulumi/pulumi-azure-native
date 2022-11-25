


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryEnvironmentVersion(ctx *pulumi.Context, args *LookupRegistryEnvironmentVersionArgs, opts ...pulumi.InvokeOption) (*LookupRegistryEnvironmentVersionResult, error) {
	var rv LookupRegistryEnvironmentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistryEnvironmentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryEnvironmentVersionArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
}


type LookupRegistryEnvironmentVersionResult struct {
	EnvironmentVersionProperties EnvironmentVersionResponse `pulumi:"environmentVersionProperties"`
	Id                           string                     `pulumi:"id"`
	Name                         string                     `pulumi:"name"`
	SystemData                   SystemDataResponse         `pulumi:"systemData"`
	Type                         string                     `pulumi:"type"`
}


func (val *LookupRegistryEnvironmentVersionResult) Defaults() *LookupRegistryEnvironmentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentVersionProperties = *tmp.EnvironmentVersionProperties.Defaults()

	return &tmp
}

func LookupRegistryEnvironmentVersionOutput(ctx *pulumi.Context, args LookupRegistryEnvironmentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryEnvironmentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryEnvironmentVersionResult, error) {
			args := v.(LookupRegistryEnvironmentVersionArgs)
			r, err := LookupRegistryEnvironmentVersion(ctx, &args, opts...)
			var s LookupRegistryEnvironmentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryEnvironmentVersionResultOutput)
}

type LookupRegistryEnvironmentVersionOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
}

func (LookupRegistryEnvironmentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryEnvironmentVersionArgs)(nil)).Elem()
}


type LookupRegistryEnvironmentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryEnvironmentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryEnvironmentVersionResult)(nil)).Elem()
}

func (o LookupRegistryEnvironmentVersionResultOutput) ToLookupRegistryEnvironmentVersionResultOutput() LookupRegistryEnvironmentVersionResultOutput {
	return o
}

func (o LookupRegistryEnvironmentVersionResultOutput) ToLookupRegistryEnvironmentVersionResultOutputWithContext(ctx context.Context) LookupRegistryEnvironmentVersionResultOutput {
	return o
}

func (o LookupRegistryEnvironmentVersionResultOutput) EnvironmentVersionProperties() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) EnvironmentVersionResponse {
		return v.EnvironmentVersionProperties
	}).(EnvironmentVersionResponseOutput)
}

func (o LookupRegistryEnvironmentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryEnvironmentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryEnvironmentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryEnvironmentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryEnvironmentVersionResultOutput{})
}
