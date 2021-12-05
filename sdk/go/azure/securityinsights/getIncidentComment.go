


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIncidentComment(ctx *pulumi.Context, args *LookupIncidentCommentArgs, opts ...pulumi.InvokeOption) (*LookupIncidentCommentResult, error) {
	var rv LookupIncidentCommentResult
	err := ctx.Invoke("azure-native:securityinsights:getIncidentComment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentCommentArgs struct {
	IncidentCommentId                   string `pulumi:"incidentCommentId"`
	IncidentId                          string `pulumi:"incidentId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupIncidentCommentResult struct {
	Author              ClientInfoResponse `pulumi:"author"`
	CreatedTimeUtc      string             `pulumi:"createdTimeUtc"`
	Etag                *string            `pulumi:"etag"`
	Id                  string             `pulumi:"id"`
	LastModifiedTimeUtc string             `pulumi:"lastModifiedTimeUtc"`
	Message             string             `pulumi:"message"`
	Name                string             `pulumi:"name"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Type                string             `pulumi:"type"`
}
