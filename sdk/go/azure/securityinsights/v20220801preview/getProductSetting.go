


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupProductSetting(ctx *pulumi.Context, args *LookupProductSettingArgs, opts ...pulumi.InvokeOption) (*LookupProductSettingResult, error) {
	var rv LookupProductSettingResult
	err := ctx.Invoke("azure-native:securityinsights/v20220801preview:getProductSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductSettingArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SettingsName      string `pulumi:"settingsName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupProductSettingResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupProductSettingOutput(ctx *pulumi.Context, args LookupProductSettingOutputArgs, opts ...pulumi.InvokeOption) LookupProductSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProductSettingResult, error) {
			args := v.(LookupProductSettingArgs)
			r, err := LookupProductSetting(ctx, &args, opts...)
			var s LookupProductSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProductSettingResultOutput)
}

type LookupProductSettingOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SettingsName      pulumi.StringInput `pulumi:"settingsName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupProductSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductSettingArgs)(nil)).Elem()
}


type LookupProductSettingResultOutput struct{ *pulumi.OutputState }

func (LookupProductSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductSettingResult)(nil)).Elem()
}

func (o LookupProductSettingResultOutput) ToLookupProductSettingResultOutput() LookupProductSettingResultOutput {
	return o
}

func (o LookupProductSettingResultOutput) ToLookupProductSettingResultOutputWithContext(ctx context.Context) LookupProductSettingResultOutput {
	return o
}

func (o LookupProductSettingResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProductSettingResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupProductSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProductSettingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductSettingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupProductSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProductSettingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupProductSettingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupProductSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProductSettingResultOutput{})
}
