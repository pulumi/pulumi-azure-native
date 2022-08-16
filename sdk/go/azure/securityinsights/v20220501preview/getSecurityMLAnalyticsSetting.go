


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSecurityMLAnalyticsSetting(ctx *pulumi.Context, args *LookupSecurityMLAnalyticsSettingArgs, opts ...pulumi.InvokeOption) (*LookupSecurityMLAnalyticsSettingResult, error) {
	var rv LookupSecurityMLAnalyticsSettingResult
	err := ctx.Invoke("azure-native:securityinsights/v20220501preview:getSecurityMLAnalyticsSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityMLAnalyticsSettingArgs struct {
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	SettingsResourceName string `pulumi:"settingsResourceName"`
	WorkspaceName        string `pulumi:"workspaceName"`
}


type LookupSecurityMLAnalyticsSettingResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupSecurityMLAnalyticsSettingOutput(ctx *pulumi.Context, args LookupSecurityMLAnalyticsSettingOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityMLAnalyticsSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityMLAnalyticsSettingResult, error) {
			args := v.(LookupSecurityMLAnalyticsSettingArgs)
			r, err := LookupSecurityMLAnalyticsSetting(ctx, &args, opts...)
			var s LookupSecurityMLAnalyticsSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityMLAnalyticsSettingResultOutput)
}

type LookupSecurityMLAnalyticsSettingOutputArgs struct {
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	SettingsResourceName pulumi.StringInput `pulumi:"settingsResourceName"`
	WorkspaceName        pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSecurityMLAnalyticsSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityMLAnalyticsSettingArgs)(nil)).Elem()
}


type LookupSecurityMLAnalyticsSettingResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityMLAnalyticsSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityMLAnalyticsSettingResult)(nil)).Elem()
}

func (o LookupSecurityMLAnalyticsSettingResultOutput) ToLookupSecurityMLAnalyticsSettingResultOutput() LookupSecurityMLAnalyticsSettingResultOutput {
	return o
}

func (o LookupSecurityMLAnalyticsSettingResultOutput) ToLookupSecurityMLAnalyticsSettingResultOutputWithContext(ctx context.Context) LookupSecurityMLAnalyticsSettingResultOutput {
	return o
}

func (o LookupSecurityMLAnalyticsSettingResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityMLAnalyticsSettingResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityMLAnalyticsSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityMLAnalyticsSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityMLAnalyticsSettingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityMLAnalyticsSettingResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupSecurityMLAnalyticsSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityMLAnalyticsSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityMLAnalyticsSettingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityMLAnalyticsSettingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSecurityMLAnalyticsSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityMLAnalyticsSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityMLAnalyticsSettingResultOutput{})
}
