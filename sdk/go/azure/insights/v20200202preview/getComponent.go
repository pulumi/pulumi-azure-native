


package v20200202preview

import (
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
