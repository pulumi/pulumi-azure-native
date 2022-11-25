


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryModelContainer(ctx *pulumi.Context, args *LookupRegistryModelContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryModelContainerResult, error) {
	var rv LookupRegistryModelContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistryModelContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryModelContainerArgs struct {
	ModelName         string `pulumi:"modelName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegistryModelContainerResult struct {
	Id                       string                 `pulumi:"id"`
	ModelContainerProperties ModelContainerResponse `pulumi:"modelContainerProperties"`
	Name                     string                 `pulumi:"name"`
	SystemData               SystemDataResponse     `pulumi:"systemData"`
	Type                     string                 `pulumi:"type"`
}


func (val *LookupRegistryModelContainerResult) Defaults() *LookupRegistryModelContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ModelContainerProperties = *tmp.ModelContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryModelContainerOutput(ctx *pulumi.Context, args LookupRegistryModelContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryModelContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryModelContainerResult, error) {
			args := v.(LookupRegistryModelContainerArgs)
			r, err := LookupRegistryModelContainer(ctx, &args, opts...)
			var s LookupRegistryModelContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryModelContainerResultOutput)
}

type LookupRegistryModelContainerOutputArgs struct {
	ModelName         pulumi.StringInput `pulumi:"modelName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryModelContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryModelContainerArgs)(nil)).Elem()
}


type LookupRegistryModelContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryModelContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryModelContainerResult)(nil)).Elem()
}

func (o LookupRegistryModelContainerResultOutput) ToLookupRegistryModelContainerResultOutput() LookupRegistryModelContainerResultOutput {
	return o
}

func (o LookupRegistryModelContainerResultOutput) ToLookupRegistryModelContainerResultOutputWithContext(ctx context.Context) LookupRegistryModelContainerResultOutput {
	return o
}

func (o LookupRegistryModelContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryModelContainerResultOutput) ModelContainerProperties() ModelContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) ModelContainerResponse { return v.ModelContainerProperties }).(ModelContainerResponseOutput)
}

func (o LookupRegistryModelContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryModelContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryModelContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryModelContainerResultOutput{})
}
