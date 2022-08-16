


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAdvancedThreatProtection(ctx *pulumi.Context, args *LookupAdvancedThreatProtectionArgs, opts ...pulumi.InvokeOption) (*LookupAdvancedThreatProtectionResult, error) {
	var rv LookupAdvancedThreatProtectionResult
	err := ctx.Invoke("azure-native:security/v20170801preview:getAdvancedThreatProtection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdvancedThreatProtectionArgs struct {
	ResourceId  string `pulumi:"resourceId"`
	SettingName string `pulumi:"settingName"`
}


type LookupAdvancedThreatProtectionResult struct {
	Id        string `pulumi:"id"`
	IsEnabled *bool  `pulumi:"isEnabled"`
	Name      string `pulumi:"name"`
	Type      string `pulumi:"type"`
}

func LookupAdvancedThreatProtectionOutput(ctx *pulumi.Context, args LookupAdvancedThreatProtectionOutputArgs, opts ...pulumi.InvokeOption) LookupAdvancedThreatProtectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdvancedThreatProtectionResult, error) {
			args := v.(LookupAdvancedThreatProtectionArgs)
			r, err := LookupAdvancedThreatProtection(ctx, &args, opts...)
			var s LookupAdvancedThreatProtectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdvancedThreatProtectionResultOutput)
}

type LookupAdvancedThreatProtectionOutputArgs struct {
	ResourceId  pulumi.StringInput `pulumi:"resourceId"`
	SettingName pulumi.StringInput `pulumi:"settingName"`
}

func (LookupAdvancedThreatProtectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdvancedThreatProtectionArgs)(nil)).Elem()
}


type LookupAdvancedThreatProtectionResultOutput struct{ *pulumi.OutputState }

func (LookupAdvancedThreatProtectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdvancedThreatProtectionResult)(nil)).Elem()
}

func (o LookupAdvancedThreatProtectionResultOutput) ToLookupAdvancedThreatProtectionResultOutput() LookupAdvancedThreatProtectionResultOutput {
	return o
}

func (o LookupAdvancedThreatProtectionResultOutput) ToLookupAdvancedThreatProtectionResultOutputWithContext(ctx context.Context) LookupAdvancedThreatProtectionResultOutput {
	return o
}

func (o LookupAdvancedThreatProtectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdvancedThreatProtectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAdvancedThreatProtectionResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAdvancedThreatProtectionResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAdvancedThreatProtectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdvancedThreatProtectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAdvancedThreatProtectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdvancedThreatProtectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdvancedThreatProtectionResultOutput{})
}
