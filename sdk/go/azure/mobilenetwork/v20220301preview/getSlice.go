


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSlice(ctx *pulumi.Context, args *LookupSliceArgs, opts ...pulumi.InvokeOption) (*LookupSliceResult, error) {
	var rv LookupSliceResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20220301preview:getSlice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSliceArgs struct {
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SliceName         string `pulumi:"sliceName"`
}


type LookupSliceResult struct {
	CreatedAt          *string            `pulumi:"createdAt"`
	CreatedBy          *string            `pulumi:"createdBy"`
	CreatedByType      *string            `pulumi:"createdByType"`
	Description        *string            `pulumi:"description"`
	Id                 string             `pulumi:"id"`
	LastModifiedAt     *string            `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string            `pulumi:"lastModifiedBy"`
	LastModifiedByType *string            `pulumi:"lastModifiedByType"`
	Location           string             `pulumi:"location"`
	Name               string             `pulumi:"name"`
	ProvisioningState  string             `pulumi:"provisioningState"`
	Snssai             SnssaiResponse     `pulumi:"snssai"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	Tags               map[string]string  `pulumi:"tags"`
	Type               string             `pulumi:"type"`
}

func LookupSliceOutput(ctx *pulumi.Context, args LookupSliceOutputArgs, opts ...pulumi.InvokeOption) LookupSliceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSliceResult, error) {
			args := v.(LookupSliceArgs)
			r, err := LookupSlice(ctx, &args, opts...)
			var s LookupSliceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSliceResultOutput)
}

type LookupSliceOutputArgs struct {
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SliceName         pulumi.StringInput `pulumi:"sliceName"`
}

func (LookupSliceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSliceArgs)(nil)).Elem()
}


type LookupSliceResultOutput struct{ *pulumi.OutputState }

func (LookupSliceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSliceResult)(nil)).Elem()
}

func (o LookupSliceResultOutput) ToLookupSliceResultOutput() LookupSliceResultOutput {
	return o
}

func (o LookupSliceResultOutput) ToLookupSliceResultOutputWithContext(ctx context.Context) LookupSliceResultOutput {
	return o
}

func (o LookupSliceResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSliceResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSliceResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSliceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSliceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSliceResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSliceResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSliceResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSliceResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSliceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSliceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSliceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSliceResultOutput) Snssai() SnssaiResponseOutput {
	return o.ApplyT(func(v LookupSliceResult) SnssaiResponse { return v.Snssai }).(SnssaiResponseOutput)
}

func (o LookupSliceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSliceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSliceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSliceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSliceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSliceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSliceResultOutput{})
}
