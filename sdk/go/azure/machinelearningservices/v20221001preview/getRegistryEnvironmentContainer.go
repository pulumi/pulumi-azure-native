


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryEnvironmentContainer(ctx *pulumi.Context, args *LookupRegistryEnvironmentContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryEnvironmentContainerResult, error) {
	var rv LookupRegistryEnvironmentContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistryEnvironmentContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryEnvironmentContainerArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegistryEnvironmentContainerResult struct {
	EnvironmentContainerProperties EnvironmentContainerResponse `pulumi:"environmentContainerProperties"`
	Id                             string                       `pulumi:"id"`
	Name                           string                       `pulumi:"name"`
	SystemData                     SystemDataResponse           `pulumi:"systemData"`
	Type                           string                       `pulumi:"type"`
}


func (val *LookupRegistryEnvironmentContainerResult) Defaults() *LookupRegistryEnvironmentContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentContainerProperties = *tmp.EnvironmentContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryEnvironmentContainerOutput(ctx *pulumi.Context, args LookupRegistryEnvironmentContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryEnvironmentContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryEnvironmentContainerResult, error) {
			args := v.(LookupRegistryEnvironmentContainerArgs)
			r, err := LookupRegistryEnvironmentContainer(ctx, &args, opts...)
			var s LookupRegistryEnvironmentContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryEnvironmentContainerResultOutput)
}

type LookupRegistryEnvironmentContainerOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryEnvironmentContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryEnvironmentContainerArgs)(nil)).Elem()
}


type LookupRegistryEnvironmentContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryEnvironmentContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryEnvironmentContainerResult)(nil)).Elem()
}

func (o LookupRegistryEnvironmentContainerResultOutput) ToLookupRegistryEnvironmentContainerResultOutput() LookupRegistryEnvironmentContainerResultOutput {
	return o
}

func (o LookupRegistryEnvironmentContainerResultOutput) ToLookupRegistryEnvironmentContainerResultOutputWithContext(ctx context.Context) LookupRegistryEnvironmentContainerResultOutput {
	return o
}

func (o LookupRegistryEnvironmentContainerResultOutput) EnvironmentContainerProperties() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) EnvironmentContainerResponse {
		return v.EnvironmentContainerProperties
	}).(EnvironmentContainerResponseOutput)
}

func (o LookupRegistryEnvironmentContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryEnvironmentContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryEnvironmentContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryEnvironmentContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryEnvironmentContainerResultOutput{})
}
