


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryComponentContainer(ctx *pulumi.Context, args *LookupRegistryComponentContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryComponentContainerResult, error) {
	var rv LookupRegistryComponentContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistryComponentContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryComponentContainerArgs struct {
	ComponentName     string `pulumi:"componentName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegistryComponentContainerResult struct {
	ComponentContainerProperties ComponentContainerResponse `pulumi:"componentContainerProperties"`
	Id                           string                     `pulumi:"id"`
	Name                         string                     `pulumi:"name"`
	SystemData                   SystemDataResponse         `pulumi:"systemData"`
	Type                         string                     `pulumi:"type"`
}


func (val *LookupRegistryComponentContainerResult) Defaults() *LookupRegistryComponentContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComponentContainerProperties = *tmp.ComponentContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryComponentContainerOutput(ctx *pulumi.Context, args LookupRegistryComponentContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryComponentContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryComponentContainerResult, error) {
			args := v.(LookupRegistryComponentContainerArgs)
			r, err := LookupRegistryComponentContainer(ctx, &args, opts...)
			var s LookupRegistryComponentContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryComponentContainerResultOutput)
}

type LookupRegistryComponentContainerOutputArgs struct {
	ComponentName     pulumi.StringInput `pulumi:"componentName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryComponentContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryComponentContainerArgs)(nil)).Elem()
}


type LookupRegistryComponentContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryComponentContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryComponentContainerResult)(nil)).Elem()
}

func (o LookupRegistryComponentContainerResultOutput) ToLookupRegistryComponentContainerResultOutput() LookupRegistryComponentContainerResultOutput {
	return o
}

func (o LookupRegistryComponentContainerResultOutput) ToLookupRegistryComponentContainerResultOutputWithContext(ctx context.Context) LookupRegistryComponentContainerResultOutput {
	return o
}

func (o LookupRegistryComponentContainerResultOutput) ComponentContainerProperties() ComponentContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) ComponentContainerResponse {
		return v.ComponentContainerProperties
	}).(ComponentContainerResponseOutput)
}

func (o LookupRegistryComponentContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryComponentContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryComponentContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryComponentContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryComponentContainerResultOutput{})
}
