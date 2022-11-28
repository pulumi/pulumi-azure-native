


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIncident(ctx *pulumi.Context, args *LookupIncidentArgs, opts ...pulumi.InvokeOption) (*LookupIncidentResult, error) {
	var rv LookupIncidentResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101preview:getIncident", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentArgs struct {
	IncidentId        string `pulumi:"incidentId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupIncidentResult struct {
	AdditionalData         IncidentAdditionalDataResponse `pulumi:"additionalData"`
	Classification         *string                        `pulumi:"classification"`
	ClassificationComment  *string                        `pulumi:"classificationComment"`
	ClassificationReason   *string                        `pulumi:"classificationReason"`
	CreatedTimeUtc         string                         `pulumi:"createdTimeUtc"`
	Description            *string                        `pulumi:"description"`
	Etag                   *string                        `pulumi:"etag"`
	FirstActivityTimeUtc   *string                        `pulumi:"firstActivityTimeUtc"`
	Id                     string                         `pulumi:"id"`
	IncidentNumber         int                            `pulumi:"incidentNumber"`
	IncidentUrl            string                         `pulumi:"incidentUrl"`
	Labels                 []IncidentLabelResponse        `pulumi:"labels"`
	LastActivityTimeUtc    *string                        `pulumi:"lastActivityTimeUtc"`
	LastModifiedTimeUtc    string                         `pulumi:"lastModifiedTimeUtc"`
	Name                   string                         `pulumi:"name"`
	Owner                  *IncidentOwnerInfoResponse     `pulumi:"owner"`
	ProviderIncidentId     *string                        `pulumi:"providerIncidentId"`
	ProviderName           *string                        `pulumi:"providerName"`
	RelatedAnalyticRuleIds []string                       `pulumi:"relatedAnalyticRuleIds"`
	Severity               string                         `pulumi:"severity"`
	Status                 string                         `pulumi:"status"`
	SystemData             SystemDataResponse             `pulumi:"systemData"`
	TeamInformation        *TeamInformationResponse       `pulumi:"teamInformation"`
	Title                  string                         `pulumi:"title"`
	Type                   string                         `pulumi:"type"`
}

func LookupIncidentOutput(ctx *pulumi.Context, args LookupIncidentOutputArgs, opts ...pulumi.InvokeOption) LookupIncidentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIncidentResult, error) {
			args := v.(LookupIncidentArgs)
			r, err := LookupIncident(ctx, &args, opts...)
			var s LookupIncidentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIncidentResultOutput)
}

type LookupIncidentOutputArgs struct {
	IncidentId        pulumi.StringInput `pulumi:"incidentId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIncidentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentArgs)(nil)).Elem()
}


type LookupIncidentResultOutput struct{ *pulumi.OutputState }

func (LookupIncidentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentResult)(nil)).Elem()
}

func (o LookupIncidentResultOutput) ToLookupIncidentResultOutput() LookupIncidentResultOutput {
	return o
}

func (o LookupIncidentResultOutput) ToLookupIncidentResultOutputWithContext(ctx context.Context) LookupIncidentResultOutput {
	return o
}

func (o LookupIncidentResultOutput) AdditionalData() IncidentAdditionalDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentResult) IncidentAdditionalDataResponse { return v.AdditionalData }).(IncidentAdditionalDataResponseOutput)
}

func (o LookupIncidentResultOutput) Classification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.Classification }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) ClassificationComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.ClassificationComment }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) ClassificationReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.ClassificationReason }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupIncidentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) FirstActivityTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.FirstActivityTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIncidentResultOutput) IncidentNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIncidentResult) int { return v.IncidentNumber }).(pulumi.IntOutput)
}

func (o LookupIncidentResultOutput) IncidentUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.IncidentUrl }).(pulumi.StringOutput)
}

func (o LookupIncidentResultOutput) Labels() IncidentLabelResponseArrayOutput {
	return o.ApplyT(func(v LookupIncidentResult) []IncidentLabelResponse { return v.Labels }).(IncidentLabelResponseArrayOutput)
}

func (o LookupIncidentResultOutput) LastActivityTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.LastActivityTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

func (o LookupIncidentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIncidentResultOutput) Owner() IncidentOwnerInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *IncidentOwnerInfoResponse { return v.Owner }).(IncidentOwnerInfoResponsePtrOutput)
}

func (o LookupIncidentResultOutput) ProviderIncidentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.ProviderIncidentId }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) ProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *string { return v.ProviderName }).(pulumi.StringPtrOutput)
}

func (o LookupIncidentResultOutput) RelatedAnalyticRuleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIncidentResult) []string { return v.RelatedAnalyticRuleIds }).(pulumi.StringArrayOutput)
}

func (o LookupIncidentResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupIncidentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupIncidentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupIncidentResultOutput) TeamInformation() TeamInformationResponsePtrOutput {
	return o.ApplyT(func(v LookupIncidentResult) *TeamInformationResponse { return v.TeamInformation }).(TeamInformationResponsePtrOutput)
}

func (o LookupIncidentResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o LookupIncidentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIncidentResultOutput{})
}
