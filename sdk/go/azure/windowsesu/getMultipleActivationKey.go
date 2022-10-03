


package windowsesu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMultipleActivationKey(ctx *pulumi.Context, args *LookupMultipleActivationKeyArgs, opts ...pulumi.InvokeOption) (*LookupMultipleActivationKeyResult, error) {
	var rv LookupMultipleActivationKeyResult
	err := ctx.Invoke("azure-native:windowsesu:getMultipleActivationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupMultipleActivationKeyArgs struct {
	MultipleActivationKeyName string `pulumi:"multipleActivationKeyName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupMultipleActivationKeyResult struct {
	AgreementNumber       *string           `pulumi:"agreementNumber"`
	ExpirationDate        string            `pulumi:"expirationDate"`
	Id                    string            `pulumi:"id"`
	InstalledServerNumber *int              `pulumi:"installedServerNumber"`
	IsEligible            *bool             `pulumi:"isEligible"`
	Location              string            `pulumi:"location"`
	MultipleActivationKey string            `pulumi:"multipleActivationKey"`
	Name                  string            `pulumi:"name"`
	OsType                *string           `pulumi:"osType"`
	ProvisioningState     string            `pulumi:"provisioningState"`
	SupportType           *string           `pulumi:"supportType"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  string            `pulumi:"type"`
}


func (val *LookupMultipleActivationKeyResult) Defaults() *LookupMultipleActivationKeyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SupportType) {
		supportType_ := "SupplementalServicing"
		tmp.SupportType = &supportType_
	}
	return &tmp
}

func LookupMultipleActivationKeyOutput(ctx *pulumi.Context, args LookupMultipleActivationKeyOutputArgs, opts ...pulumi.InvokeOption) LookupMultipleActivationKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMultipleActivationKeyResult, error) {
			args := v.(LookupMultipleActivationKeyArgs)
			r, err := LookupMultipleActivationKey(ctx, &args, opts...)
			var s LookupMultipleActivationKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMultipleActivationKeyResultOutput)
}

type LookupMultipleActivationKeyOutputArgs struct {
	MultipleActivationKeyName pulumi.StringInput `pulumi:"multipleActivationKeyName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMultipleActivationKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMultipleActivationKeyArgs)(nil)).Elem()
}


type LookupMultipleActivationKeyResultOutput struct{ *pulumi.OutputState }

func (LookupMultipleActivationKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMultipleActivationKeyResult)(nil)).Elem()
}

func (o LookupMultipleActivationKeyResultOutput) ToLookupMultipleActivationKeyResultOutput() LookupMultipleActivationKeyResultOutput {
	return o
}

func (o LookupMultipleActivationKeyResultOutput) ToLookupMultipleActivationKeyResultOutputWithContext(ctx context.Context) LookupMultipleActivationKeyResultOutput {
	return o
}

func (o LookupMultipleActivationKeyResultOutput) AgreementNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) *string { return v.AgreementNumber }).(pulumi.StringPtrOutput)
}

func (o LookupMultipleActivationKeyResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o LookupMultipleActivationKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMultipleActivationKeyResultOutput) InstalledServerNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) *int { return v.InstalledServerNumber }).(pulumi.IntPtrOutput)
}

func (o LookupMultipleActivationKeyResultOutput) IsEligible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) *bool { return v.IsEligible }).(pulumi.BoolPtrOutput)
}

func (o LookupMultipleActivationKeyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMultipleActivationKeyResultOutput) MultipleActivationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) string { return v.MultipleActivationKey }).(pulumi.StringOutput)
}

func (o LookupMultipleActivationKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMultipleActivationKeyResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupMultipleActivationKeyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupMultipleActivationKeyResultOutput) SupportType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) *string { return v.SupportType }).(pulumi.StringPtrOutput)
}

func (o LookupMultipleActivationKeyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMultipleActivationKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMultipleActivationKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMultipleActivationKeyResultOutput{})
}
