


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistryArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupRegistryResult struct {
	Id                 string                          `pulumi:"id"`
	Identity           *ManagedServiceIdentityResponse `pulumi:"identity"`
	Kind               *string                         `pulumi:"kind"`
	Location           string                          `pulumi:"location"`
	Name               string                          `pulumi:"name"`
	RegistryProperties RegistryResponse                `pulumi:"registryProperties"`
	Sku                *SkuResponse                    `pulumi:"sku"`
	SystemData         SystemDataResponse              `pulumi:"systemData"`
	Tags               map[string]string               `pulumi:"tags"`
	Type               string                          `pulumi:"type"`
}

func LookupRegistryOutput(ctx *pulumi.Context, args LookupRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryResult, error) {
			args := v.(LookupRegistryArgs)
			r, err := LookupRegistry(ctx, &args, opts...)
			var s LookupRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryResultOutput)
}

type LookupRegistryOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryArgs)(nil)).Elem()
}

type LookupRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryResult)(nil)).Elem()
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutput() LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutputWithContext(ctx context.Context) LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupRegistryResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupRegistryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryResultOutput) RegistryProperties() RegistryResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) RegistryResponse { return v.RegistryProperties }).(RegistryResponseOutput)
}

func (o LookupRegistryResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupRegistryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegistryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryResultOutput{})
}
