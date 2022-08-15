


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSetting(ctx *pulumi.Context, args *LookupSettingArgs, opts ...pulumi.InvokeOption) (*LookupSettingResult, error) {
	var rv LookupSettingResult
	err := ctx.Invoke("azure-native:costmanagement/v20191101:getSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSettingArgs struct {
	SettingName string `pulumi:"settingName"`
}


type LookupSettingResult struct {
	Cache   []SettingsPropertiesResponseCache `pulumi:"cache"`
	Id      string                            `pulumi:"id"`
	Kind    string                            `pulumi:"kind"`
	Name    string                            `pulumi:"name"`
	Scope   string                            `pulumi:"scope"`
	StartOn *string                           `pulumi:"startOn"`
	Type    string                            `pulumi:"type"`
}

func LookupSettingOutput(ctx *pulumi.Context, args LookupSettingOutputArgs, opts ...pulumi.InvokeOption) LookupSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSettingResult, error) {
			args := v.(LookupSettingArgs)
			r, err := LookupSetting(ctx, &args, opts...)
			var s LookupSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSettingResultOutput)
}

type LookupSettingOutputArgs struct {
	SettingName pulumi.StringInput `pulumi:"settingName"`
}

func (LookupSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSettingArgs)(nil)).Elem()
}


type LookupSettingResultOutput struct{ *pulumi.OutputState }

func (LookupSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSettingResult)(nil)).Elem()
}

func (o LookupSettingResultOutput) ToLookupSettingResultOutput() LookupSettingResultOutput {
	return o
}

func (o LookupSettingResultOutput) ToLookupSettingResultOutputWithContext(ctx context.Context) LookupSettingResultOutput {
	return o
}

func (o LookupSettingResultOutput) Cache() SettingsPropertiesResponseCacheArrayOutput {
	return o.ApplyT(func(v LookupSettingResult) []SettingsPropertiesResponseCache { return v.Cache }).(SettingsPropertiesResponseCacheArrayOutput)
}

func (o LookupSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSettingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSettingResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupSettingResultOutput) StartOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSettingResult) *string { return v.StartOn }).(pulumi.StringPtrOutput)
}

func (o LookupSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSettingResultOutput{})
}
