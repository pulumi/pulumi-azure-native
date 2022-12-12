


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBastionHost(ctx *pulumi.Context, args *LookupBastionHostArgs, opts ...pulumi.InvokeOption) (*LookupBastionHostResult, error) {
	var rv LookupBastionHostResult
	err := ctx.Invoke("azure-native:network/v20210301:getBastionHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBastionHostArgs struct {
	BastionHostName   string `pulumi:"bastionHostName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBastionHostResult struct {
	DisableCopyPaste    *bool                                `pulumi:"disableCopyPaste"`
	DnsName             *string                              `pulumi:"dnsName"`
	EnableFileCopy      *bool                                `pulumi:"enableFileCopy"`
	EnableIpConnect     *bool                                `pulumi:"enableIpConnect"`
	EnableShareableLink *bool                                `pulumi:"enableShareableLink"`
	EnableTunneling     *bool                                `pulumi:"enableTunneling"`
	Etag                string                               `pulumi:"etag"`
	Id                  *string                              `pulumi:"id"`
	IpConfigurations    []BastionHostIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location            *string                              `pulumi:"location"`
	Name                string                               `pulumi:"name"`
	ProvisioningState   string                               `pulumi:"provisioningState"`
	ScaleUnits          *int                                 `pulumi:"scaleUnits"`
	Sku                 *SkuResponse                         `pulumi:"sku"`
	Tags                map[string]string                    `pulumi:"tags"`
	Type                string                               `pulumi:"type"`
}


func (val *LookupBastionHostResult) Defaults() *LookupBastionHostResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableCopyPaste) {
		disableCopyPaste_ := false
		tmp.DisableCopyPaste = &disableCopyPaste_
	}
	if isZero(tmp.EnableFileCopy) {
		enableFileCopy_ := false
		tmp.EnableFileCopy = &enableFileCopy_
	}
	if isZero(tmp.EnableIpConnect) {
		enableIpConnect_ := false
		tmp.EnableIpConnect = &enableIpConnect_
	}
	if isZero(tmp.EnableShareableLink) {
		enableShareableLink_ := false
		tmp.EnableShareableLink = &enableShareableLink_
	}
	if isZero(tmp.EnableTunneling) {
		enableTunneling_ := false
		tmp.EnableTunneling = &enableTunneling_
	}
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupBastionHostOutput(ctx *pulumi.Context, args LookupBastionHostOutputArgs, opts ...pulumi.InvokeOption) LookupBastionHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBastionHostResult, error) {
			args := v.(LookupBastionHostArgs)
			r, err := LookupBastionHost(ctx, &args, opts...)
			var s LookupBastionHostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBastionHostResultOutput)
}

type LookupBastionHostOutputArgs struct {
	BastionHostName   pulumi.StringInput `pulumi:"bastionHostName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBastionHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBastionHostArgs)(nil)).Elem()
}


type LookupBastionHostResultOutput struct{ *pulumi.OutputState }

func (LookupBastionHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBastionHostResult)(nil)).Elem()
}

func (o LookupBastionHostResultOutput) ToLookupBastionHostResultOutput() LookupBastionHostResultOutput {
	return o
}

func (o LookupBastionHostResultOutput) ToLookupBastionHostResultOutputWithContext(ctx context.Context) LookupBastionHostResultOutput {
	return o
}

func (o LookupBastionHostResultOutput) DisableCopyPaste() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *bool { return v.DisableCopyPaste }).(pulumi.BoolPtrOutput)
}

func (o LookupBastionHostResultOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

func (o LookupBastionHostResultOutput) EnableFileCopy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *bool { return v.EnableFileCopy }).(pulumi.BoolPtrOutput)
}

func (o LookupBastionHostResultOutput) EnableIpConnect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *bool { return v.EnableIpConnect }).(pulumi.BoolPtrOutput)
}

func (o LookupBastionHostResultOutput) EnableShareableLink() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *bool { return v.EnableShareableLink }).(pulumi.BoolPtrOutput)
}

func (o LookupBastionHostResultOutput) EnableTunneling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *bool { return v.EnableTunneling }).(pulumi.BoolPtrOutput)
}

func (o LookupBastionHostResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBastionHostResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupBastionHostResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupBastionHostResultOutput) IpConfigurations() BastionHostIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupBastionHostResult) []BastionHostIPConfigurationResponse { return v.IpConfigurations }).(BastionHostIPConfigurationResponseArrayOutput)
}

func (o LookupBastionHostResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupBastionHostResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBastionHostResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBastionHostResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBastionHostResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupBastionHostResultOutput) ScaleUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *int { return v.ScaleUnits }).(pulumi.IntPtrOutput)
}

func (o LookupBastionHostResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBastionHostResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupBastionHostResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBastionHostResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupBastionHostResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBastionHostResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBastionHostResultOutput{})
}
