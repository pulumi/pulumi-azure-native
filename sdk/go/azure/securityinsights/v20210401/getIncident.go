


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIncident(ctx *pulumi.Context, args *LookupIncidentArgs, opts ...pulumi.InvokeOption) (*LookupIncidentResult, error) {
	var rv LookupIncidentResult
	err := ctx.Invoke("azure-native:securityinsights/v20210401:getIncident", args, &rv, opts...)
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
	RelatedAnalyticRuleIds []string                       `pulumi:"relatedAnalyticRuleIds"`
	Severity               string                         `pulumi:"severity"`
	Status                 string                         `pulumi:"status"`
	SystemData             SystemDataResponse             `pulumi:"systemData"`
	Title                  string                         `pulumi:"title"`
	Type                   string                         `pulumi:"type"`
}
