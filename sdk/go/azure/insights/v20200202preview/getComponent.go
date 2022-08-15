


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComponent(ctx *pulumi.Context, args *LookupComponentArgs, opts ...pulumi.InvokeOption) (*LookupComponentResult, error) {
	var rv LookupComponentResult
	err := ctx.Invoke("azure-native:insights/v20200202preview:getComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupComponentArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupComponentResult struct {
	AppId                           string                              `pulumi:"appId"`
	ApplicationId                   string                              `pulumi:"applicationId"`
	ApplicationType                 string                              `pulumi:"applicationType"`
	ConnectionString                string                              `pulumi:"connectionString"`
	CreationDate                    string                              `pulumi:"creationDate"`
	DisableIpMasking                *bool                               `pulumi:"disableIpMasking"`
	DisableLocalAuth                *bool                               `pulumi:"disableLocalAuth"`
	Etag                            *string                             `pulumi:"etag"`
	FlowType                        *string                             `pulumi:"flowType"`
	ForceCustomerStorageForProfiler *bool                               `pulumi:"forceCustomerStorageForProfiler"`
	HockeyAppId                     *string                             `pulumi:"hockeyAppId"`
	HockeyAppToken                  string                              `pulumi:"hockeyAppToken"`
	Id                              string                              `pulumi:"id"`
	ImmediatePurgeDataOn30Days      *bool                               `pulumi:"immediatePurgeDataOn30Days"`
	IngestionMode                   *string                             `pulumi:"ingestionMode"`
	InstrumentationKey              string                              `pulumi:"instrumentationKey"`
	Kind                            string                              `pulumi:"kind"`
	LaMigrationDate                 string                              `pulumi:"laMigrationDate"`
	Location                        string                              `pulumi:"location"`
	Name                            string                              `pulumi:"name"`
	PrivateLinkScopedResources      []PrivateLinkScopedResourceResponse `pulumi:"privateLinkScopedResources"`
	ProvisioningState               string                              `pulumi:"provisioningState"`
	PublicNetworkAccessForIngestion *string                             `pulumi:"publicNetworkAccessForIngestion"`
	PublicNetworkAccessForQuery     *string                             `pulumi:"publicNetworkAccessForQuery"`
	RequestSource                   *string                             `pulumi:"requestSource"`
	RetentionInDays                 int                                 `pulumi:"retentionInDays"`
	SamplingPercentage              *float64                            `pulumi:"samplingPercentage"`
	Tags                            map[string]string                   `pulumi:"tags"`
	TenantId                        string                              `pulumi:"tenantId"`
	Type                            string                              `pulumi:"type"`
	WorkspaceResourceId             *string                             `pulumi:"workspaceResourceId"`
}


func (val *LookupComponentResult) Defaults() *LookupComponentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ApplicationType) {
		tmp.ApplicationType = "web"
	}
	if isZero(tmp.FlowType) {
		flowType_ := "Bluefield"
		tmp.FlowType = &flowType_
	}
	if isZero(tmp.IngestionMode) {
		ingestionMode_ := "LogAnalytics"
		tmp.IngestionMode = &ingestionMode_
	}
	if isZero(tmp.RequestSource) {
		requestSource_ := "rest"
		tmp.RequestSource = &requestSource_
	}
	return &tmp
}

func LookupComponentOutput(ctx *pulumi.Context, args LookupComponentOutputArgs, opts ...pulumi.InvokeOption) LookupComponentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentResult, error) {
			args := v.(LookupComponentArgs)
			r, err := LookupComponent(ctx, &args, opts...)
			var s LookupComponentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentResultOutput)
}

type LookupComponentOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentArgs)(nil)).Elem()
}


type LookupComponentResultOutput struct{ *pulumi.OutputState }

func (LookupComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentResult)(nil)).Elem()
}

func (o LookupComponentResultOutput) ToLookupComponentResultOutput() LookupComponentResultOutput {
	return o
}

func (o LookupComponentResultOutput) ToLookupComponentResultOutputWithContext(ctx context.Context) LookupComponentResultOutput {
	return o
}

func (o LookupComponentResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.AppId }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) ApplicationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.ApplicationType }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) DisableIpMasking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *bool { return v.DisableIpMasking }).(pulumi.BoolPtrOutput)
}

func (o LookupComponentResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupComponentResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) FlowType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.FlowType }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) ForceCustomerStorageForProfiler() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *bool { return v.ForceCustomerStorageForProfiler }).(pulumi.BoolPtrOutput)
}

func (o LookupComponentResultOutput) HockeyAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.HockeyAppId }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) HockeyAppToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.HockeyAppToken }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *bool { return v.ImmediatePurgeDataOn30Days }).(pulumi.BoolPtrOutput)
}

func (o LookupComponentResultOutput) IngestionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.IngestionMode }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) InstrumentationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.InstrumentationKey }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) LaMigrationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.LaMigrationDate }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) PrivateLinkScopedResources() PrivateLinkScopedResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupComponentResult) []PrivateLinkScopedResourceResponse { return v.PrivateLinkScopedResources }).(PrivateLinkScopedResourceResponseArrayOutput)
}

func (o LookupComponentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) PublicNetworkAccessForIngestion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.PublicNetworkAccessForIngestion }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) PublicNetworkAccessForQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.PublicNetworkAccessForQuery }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) RequestSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.RequestSource }).(pulumi.StringPtrOutput)
}

func (o LookupComponentResultOutput) RetentionInDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComponentResult) int { return v.RetentionInDays }).(pulumi.IntOutput)
}

func (o LookupComponentResultOutput) SamplingPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *float64 { return v.SamplingPercentage }).(pulumi.Float64PtrOutput)
}

func (o LookupComponentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComponentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupComponentResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupComponentResultOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentResult) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentResultOutput{})
}
