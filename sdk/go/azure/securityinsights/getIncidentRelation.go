


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIncidentRelation(ctx *pulumi.Context, args *LookupIncidentRelationArgs, opts ...pulumi.InvokeOption) (*LookupIncidentRelationResult, error) {
	var rv LookupIncidentRelationResult
	err := ctx.Invoke("azure-native:securityinsights:getIncidentRelation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentRelationArgs struct {
	IncidentId                          string `pulumi:"incidentId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	RelationName                        string `pulumi:"relationName"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupIncidentRelationResult struct {
	Etag                *string            `pulumi:"etag"`
	Id                  string             `pulumi:"id"`
	Name                string             `pulumi:"name"`
	RelatedResourceId   string             `pulumi:"relatedResourceId"`
	RelatedResourceKind string             `pulumi:"relatedResourceKind"`
	RelatedResourceName string             `pulumi:"relatedResourceName"`
	RelatedResourceType string             `pulumi:"relatedResourceType"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Type                string             `pulumi:"type"`
}
