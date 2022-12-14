


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azure-native:cdn/v20221101preview:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupProfileResult struct {
	ExtendedProperties           map[string]string               `pulumi:"extendedProperties"`
	FrontDoorId                  string                          `pulumi:"frontDoorId"`
	Id                           string                          `pulumi:"id"`
	Identity                     *ManagedServiceIdentityResponse `pulumi:"identity"`
	Kind                         string                          `pulumi:"kind"`
	Location                     string                          `pulumi:"location"`
	Name                         string                          `pulumi:"name"`
	OriginResponseTimeoutSeconds *int                            `pulumi:"originResponseTimeoutSeconds"`
	ProvisioningState            string                          `pulumi:"provisioningState"`
	ResourceState                string                          `pulumi:"resourceState"`
	Sku                          SkuResponse                     `pulumi:"sku"`
	SystemData                   SystemDataResponse              `pulumi:"systemData"`
	Tags                         map[string]string               `pulumi:"tags"`
	Type                         string                          `pulumi:"type"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResult, error) {
			args := v.(LookupProfileArgs)
			r, err := LookupProfile(ctx, &args, opts...)
			var s LookupProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfileResultOutput)
}

type LookupProfileOutputArgs struct {
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}


type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ExtendedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string]string { return v.ExtendedProperties }).(pulumi.StringMapOutput)
}

func (o LookupProfileResultOutput) FrontDoorId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.FrontDoorId }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupProfileResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) OriginResponseTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *int { return v.OriginResponseTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupProfileResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupProfileResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupProfileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProfileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}
