


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBookmark(ctx *pulumi.Context, args *LookupBookmarkArgs, opts ...pulumi.InvokeOption) (*LookupBookmarkResult, error) {
	var rv LookupBookmarkResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getBookmark", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBookmarkArgs struct {
	BookmarkId        string `pulumi:"bookmarkId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupBookmarkResult struct {
	Created        *string               `pulumi:"created"`
	CreatedBy      *UserInfoResponse     `pulumi:"createdBy"`
	DisplayName    string                `pulumi:"displayName"`
	Etag           *string               `pulumi:"etag"`
	EventTime      *string               `pulumi:"eventTime"`
	Id             string                `pulumi:"id"`
	IncidentInfo   *IncidentInfoResponse `pulumi:"incidentInfo"`
	Labels         []string              `pulumi:"labels"`
	Name           string                `pulumi:"name"`
	Notes          *string               `pulumi:"notes"`
	Query          string                `pulumi:"query"`
	QueryEndTime   *string               `pulumi:"queryEndTime"`
	QueryResult    *string               `pulumi:"queryResult"`
	QueryStartTime *string               `pulumi:"queryStartTime"`
	Type           string                `pulumi:"type"`
	Updated        *string               `pulumi:"updated"`
	UpdatedBy      *UserInfoResponse     `pulumi:"updatedBy"`
}
