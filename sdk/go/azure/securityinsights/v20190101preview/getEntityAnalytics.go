


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEntityAnalytics(ctx *pulumi.Context, args *LookupEntityAnalyticsArgs, opts ...pulumi.InvokeOption) (*LookupEntityAnalyticsResult, error) {
	var rv LookupEntityAnalyticsResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getEntityAnalytics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntityAnalyticsArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SettingsName                        string `pulumi:"settingsName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupEntityAnalyticsResult struct {
	Etag      *string `pulumi:"etag"`
	Id        string  `pulumi:"id"`
	IsEnabled bool    `pulumi:"isEnabled"`
	Kind      string  `pulumi:"kind"`
	Name      string  `pulumi:"name"`
	Type      string  `pulumi:"type"`
}

func LookupEntityAnalyticsOutput(ctx *pulumi.Context, args LookupEntityAnalyticsOutputArgs, opts ...pulumi.InvokeOption) LookupEntityAnalyticsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEntityAnalyticsResult, error) {
			args := v.(LookupEntityAnalyticsArgs)
			r, err := LookupEntityAnalytics(ctx, &args, opts...)
			var s LookupEntityAnalyticsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEntityAnalyticsResultOutput)
}

type LookupEntityAnalyticsOutputArgs struct {
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   pulumi.StringInput `pulumi:"resourceGroupName"`
	SettingsName                        pulumi.StringInput `pulumi:"settingsName"`
	WorkspaceName                       pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEntityAnalyticsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityAnalyticsArgs)(nil)).Elem()
}


type LookupEntityAnalyticsResultOutput struct{ *pulumi.OutputState }

func (LookupEntityAnalyticsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityAnalyticsResult)(nil)).Elem()
}

func (o LookupEntityAnalyticsResultOutput) ToLookupEntityAnalyticsResultOutput() LookupEntityAnalyticsResultOutput {
	return o
}

func (o LookupEntityAnalyticsResultOutput) ToLookupEntityAnalyticsResultOutputWithContext(ctx context.Context) LookupEntityAnalyticsResultOutput {
	return o
}

func (o LookupEntityAnalyticsResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupEntityAnalyticsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEntityAnalyticsResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o LookupEntityAnalyticsResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupEntityAnalyticsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEntityAnalyticsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityAnalyticsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEntityAnalyticsResultOutput{})
}
