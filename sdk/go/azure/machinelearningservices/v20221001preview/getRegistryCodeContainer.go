


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryCodeContainer(ctx *pulumi.Context, args *LookupRegistryCodeContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryCodeContainerResult, error) {
	var rv LookupRegistryCodeContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistryCodeContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryCodeContainerArgs struct {
	CodeName          string `pulumi:"codeName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegistryCodeContainerResult struct {
	CodeContainerProperties CodeContainerResponse `pulumi:"codeContainerProperties"`
	Id                      string                `pulumi:"id"`
	Name                    string                `pulumi:"name"`
	SystemData              SystemDataResponse    `pulumi:"systemData"`
	Type                    string                `pulumi:"type"`
}


func (val *LookupRegistryCodeContainerResult) Defaults() *LookupRegistryCodeContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.CodeContainerProperties = *tmp.CodeContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryCodeContainerOutput(ctx *pulumi.Context, args LookupRegistryCodeContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryCodeContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryCodeContainerResult, error) {
			args := v.(LookupRegistryCodeContainerArgs)
			r, err := LookupRegistryCodeContainer(ctx, &args, opts...)
			var s LookupRegistryCodeContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryCodeContainerResultOutput)
}

type LookupRegistryCodeContainerOutputArgs struct {
	CodeName          pulumi.StringInput `pulumi:"codeName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryCodeContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCodeContainerArgs)(nil)).Elem()
}


type LookupRegistryCodeContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryCodeContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCodeContainerResult)(nil)).Elem()
}

func (o LookupRegistryCodeContainerResultOutput) ToLookupRegistryCodeContainerResultOutput() LookupRegistryCodeContainerResultOutput {
	return o
}

func (o LookupRegistryCodeContainerResultOutput) ToLookupRegistryCodeContainerResultOutputWithContext(ctx context.Context) LookupRegistryCodeContainerResultOutput {
	return o
}

func (o LookupRegistryCodeContainerResultOutput) CodeContainerProperties() CodeContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) CodeContainerResponse { return v.CodeContainerProperties }).(CodeContainerResponseOutput)
}

func (o LookupRegistryCodeContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryCodeContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryCodeContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryCodeContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryCodeContainerResultOutput{})
}
