


package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiIssue(ctx *pulumi.Context, args *LookupApiIssueArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueResult, error) {
	var rv LookupApiIssueResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:getApiIssue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiIssueArgs struct {
	ApiId                     string `pulumi:"apiId"`
	ExpandCommentsAttachments *bool  `pulumi:"expandCommentsAttachments"`
	IssueId                   string `pulumi:"issueId"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ServiceName               string `pulumi:"serviceName"`
}


type LookupApiIssueResult struct {
	ApiId       *string `pulumi:"apiId"`
	CreatedDate *string `pulumi:"createdDate"`
	Description string  `pulumi:"description"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	State       *string `pulumi:"state"`
	Title       string  `pulumi:"title"`
	Type        string  `pulumi:"type"`
	UserId      string  `pulumi:"userId"`
}
