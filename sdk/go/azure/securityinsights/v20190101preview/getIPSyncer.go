


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIPSyncer(ctx *pulumi.Context, args *LookupIPSyncerArgs, opts ...pulumi.InvokeOption) (*LookupIPSyncerResult, error) {
	var rv LookupIPSyncerResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getIPSyncer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPSyncerArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SettingsName                        string `pulumi:"settingsName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupIPSyncerResult struct {
	Etag      *string `pulumi:"etag"`
	Id        string  `pulumi:"id"`
	IsEnabled bool    `pulumi:"isEnabled"`
	Kind      string  `pulumi:"kind"`
	Name      string  `pulumi:"name"`
	Type      string  `pulumi:"type"`
}

func LookupIPSyncerOutput(ctx *pulumi.Context, args LookupIPSyncerOutputArgs, opts ...pulumi.InvokeOption) LookupIPSyncerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPSyncerResult, error) {
			args := v.(LookupIPSyncerArgs)
			r, err := LookupIPSyncer(ctx, &args, opts...)
			var s LookupIPSyncerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIPSyncerResultOutput)
}

type LookupIPSyncerOutputArgs struct {
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   pulumi.StringInput `pulumi:"resourceGroupName"`
	SettingsName                        pulumi.StringInput `pulumi:"settingsName"`
	WorkspaceName                       pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIPSyncerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPSyncerArgs)(nil)).Elem()
}


type LookupIPSyncerResultOutput struct{ *pulumi.OutputState }

func (LookupIPSyncerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPSyncerResult)(nil)).Elem()
}

func (o LookupIPSyncerResultOutput) ToLookupIPSyncerResultOutput() LookupIPSyncerResultOutput {
	return o
}

func (o LookupIPSyncerResultOutput) ToLookupIPSyncerResultOutputWithContext(ctx context.Context) LookupIPSyncerResultOutput {
	return o
}

func (o LookupIPSyncerResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPSyncerResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIPSyncerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPSyncerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIPSyncerResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIPSyncerResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o LookupIPSyncerResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPSyncerResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupIPSyncerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPSyncerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIPSyncerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPSyncerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPSyncerResultOutput{})
}
