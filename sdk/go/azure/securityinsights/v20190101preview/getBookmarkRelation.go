


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBookmarkRelation(ctx *pulumi.Context, args *LookupBookmarkRelationArgs, opts ...pulumi.InvokeOption) (*LookupBookmarkRelationResult, error) {
	var rv LookupBookmarkRelationResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getBookmarkRelation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBookmarkRelationArgs struct {
	BookmarkId                          string `pulumi:"bookmarkId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	RelationName                        string `pulumi:"relationName"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupBookmarkRelationResult struct {
	Etag                *string `pulumi:"etag"`
	Id                  string  `pulumi:"id"`
	Name                string  `pulumi:"name"`
	RelatedResourceId   string  `pulumi:"relatedResourceId"`
	RelatedResourceKind string  `pulumi:"relatedResourceKind"`
	RelatedResourceName string  `pulumi:"relatedResourceName"`
	RelatedResourceType string  `pulumi:"relatedResourceType"`
	Type                string  `pulumi:"type"`
}
