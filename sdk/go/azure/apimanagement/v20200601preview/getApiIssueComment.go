


package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiIssueComment(ctx *pulumi.Context, args *LookupApiIssueCommentArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueCommentResult, error) {
	var rv LookupApiIssueCommentResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getApiIssueComment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiIssueCommentArgs struct {
	ApiId             string `pulumi:"apiId"`
	CommentId         string `pulumi:"commentId"`
	IssueId           string `pulumi:"issueId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiIssueCommentResult struct {
	CreatedDate *string `pulumi:"createdDate"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	Text        string  `pulumi:"text"`
	Type        string  `pulumi:"type"`
	UserId      string  `pulumi:"userId"`
}
