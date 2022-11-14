


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAnomalySecurityMLAnalyticsSettings(ctx *pulumi.Context, args *LookupAnomalySecurityMLAnalyticsSettingsArgs, opts ...pulumi.InvokeOption) (*LookupAnomalySecurityMLAnalyticsSettingsResult, error) {
	var rv LookupAnomalySecurityMLAnalyticsSettingsResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101preview:getAnomalySecurityMLAnalyticsSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnomalySecurityMLAnalyticsSettingsArgs struct {
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	SettingsResourceName string `pulumi:"settingsResourceName"`
	WorkspaceName        string `pulumi:"workspaceName"`
}


type LookupAnomalySecurityMLAnalyticsSettingsResult struct {
	AnomalySettingsVersion   *int                                            `pulumi:"anomalySettingsVersion"`
	AnomalyVersion           string                                          `pulumi:"anomalyVersion"`
	CustomizableObservations interface{}                                     `pulumi:"customizableObservations"`
	Description              *string                                         `pulumi:"description"`
	DisplayName              string                                          `pulumi:"displayName"`
	Enabled                  bool                                            `pulumi:"enabled"`
	Etag                     *string                                         `pulumi:"etag"`
	Frequency                string                                          `pulumi:"frequency"`
	Id                       string                                          `pulumi:"id"`
	IsDefaultSettings        bool                                            `pulumi:"isDefaultSettings"`
	Kind                     string                                          `pulumi:"kind"`
	LastModifiedUtc          string                                          `pulumi:"lastModifiedUtc"`
	Name                     string                                          `pulumi:"name"`
	RequiredDataConnectors   []SecurityMLAnalyticsSettingsDataSourceResponse `pulumi:"requiredDataConnectors"`
	SettingsDefinitionId     *string                                         `pulumi:"settingsDefinitionId"`
	SettingsStatus           string                                          `pulumi:"settingsStatus"`
	SystemData               SystemDataResponse                              `pulumi:"systemData"`
	Tactics                  []string                                        `pulumi:"tactics"`
	Techniques               []string                                        `pulumi:"techniques"`
	Type                     string                                          `pulumi:"type"`
}

func LookupAnomalySecurityMLAnalyticsSettingsOutput(ctx *pulumi.Context, args LookupAnomalySecurityMLAnalyticsSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupAnomalySecurityMLAnalyticsSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnomalySecurityMLAnalyticsSettingsResult, error) {
			args := v.(LookupAnomalySecurityMLAnalyticsSettingsArgs)
			r, err := LookupAnomalySecurityMLAnalyticsSettings(ctx, &args, opts...)
			var s LookupAnomalySecurityMLAnalyticsSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnomalySecurityMLAnalyticsSettingsResultOutput)
}

type LookupAnomalySecurityMLAnalyticsSettingsOutputArgs struct {
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	SettingsResourceName pulumi.StringInput `pulumi:"settingsResourceName"`
	WorkspaceName        pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAnomalySecurityMLAnalyticsSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomalySecurityMLAnalyticsSettingsArgs)(nil)).Elem()
}


type LookupAnomalySecurityMLAnalyticsSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupAnomalySecurityMLAnalyticsSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomalySecurityMLAnalyticsSettingsResult)(nil)).Elem()
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) ToLookupAnomalySecurityMLAnalyticsSettingsResultOutput() LookupAnomalySecurityMLAnalyticsSettingsResultOutput {
	return o
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) ToLookupAnomalySecurityMLAnalyticsSettingsResultOutputWithContext(ctx context.Context) LookupAnomalySecurityMLAnalyticsSettingsResultOutput {
	return o
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) AnomalySettingsVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) *int { return v.AnomalySettingsVersion }).(pulumi.IntPtrOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) AnomalyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.AnomalyVersion }).(pulumi.StringOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) CustomizableObservations() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) interface{} { return v.CustomizableObservations }).(pulumi.AnyOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) IsDefaultSettings() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) bool { return v.IsDefaultSettings }).(pulumi.BoolOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) RequiredDataConnectors() SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) []SecurityMLAnalyticsSettingsDataSourceResponse {
		return v.RequiredDataConnectors
	}).(SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) SettingsDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) *string { return v.SettingsDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) SettingsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.SettingsStatus }).(pulumi.StringOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnomalySecurityMLAnalyticsSettingsResultOutput{})
}
