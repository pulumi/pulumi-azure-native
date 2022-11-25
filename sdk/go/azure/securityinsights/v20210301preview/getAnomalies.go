


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAnomalies(ctx *pulumi.Context, args *LookupAnomaliesArgs, opts ...pulumi.InvokeOption) (*LookupAnomaliesResult, error) {
	var rv LookupAnomaliesResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getAnomalies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnomaliesArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	SettingsName                        string `pulumi:"settingsName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupAnomaliesResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	IsEnabled  bool               `pulumi:"isEnabled"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupAnomaliesOutput(ctx *pulumi.Context, args LookupAnomaliesOutputArgs, opts ...pulumi.InvokeOption) LookupAnomaliesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnomaliesResult, error) {
			args := v.(LookupAnomaliesArgs)
			r, err := LookupAnomalies(ctx, &args, opts...)
			var s LookupAnomaliesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnomaliesResultOutput)
}

type LookupAnomaliesOutputArgs struct {
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   pulumi.StringInput `pulumi:"resourceGroupName"`
	SettingsName                        pulumi.StringInput `pulumi:"settingsName"`
	WorkspaceName                       pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAnomaliesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomaliesArgs)(nil)).Elem()
}


type LookupAnomaliesResultOutput struct{ *pulumi.OutputState }

func (LookupAnomaliesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomaliesResult)(nil)).Elem()
}

func (o LookupAnomaliesResultOutput) ToLookupAnomaliesResultOutput() LookupAnomaliesResultOutput {
	return o
}

func (o LookupAnomaliesResultOutput) ToLookupAnomaliesResultOutputWithContext(ctx context.Context) LookupAnomaliesResultOutput {
	return o
}

func (o LookupAnomaliesResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAnomaliesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAnomaliesResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o LookupAnomaliesResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAnomaliesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAnomaliesResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAnomaliesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomaliesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnomaliesResultOutput{})
}
