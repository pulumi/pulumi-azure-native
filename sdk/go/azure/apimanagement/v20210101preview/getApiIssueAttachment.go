


package v20210101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiIssueAttachment(ctx *pulumi.Context, args *LookupApiIssueAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueAttachmentResult, error) {
	var rv LookupApiIssueAttachmentResult
	err := ctx.Invoke("azure-native:apimanagement/v20210101preview:getApiIssueAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiIssueAttachmentArgs struct {
	ApiId             string `pulumi:"apiId"`
	AttachmentId      string `pulumi:"attachmentId"`
	IssueId           string `pulumi:"issueId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiIssueAttachmentResult struct {
	Content       string `pulumi:"content"`
	ContentFormat string `pulumi:"contentFormat"`
	Id            string `pulumi:"id"`
	Name          string `pulumi:"name"`
	Title         string `pulumi:"title"`
	Type          string `pulumi:"type"`
}
