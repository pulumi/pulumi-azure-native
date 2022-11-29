


package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiIssueComment(ctx *pulumi.Context, args *LookupApiIssueCommentArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueCommentResult, error) {
	var rv LookupApiIssueCommentResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getApiIssueComment", args, &rv, opts...)
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

func LookupApiIssueCommentOutput(ctx *pulumi.Context, args LookupApiIssueCommentOutputArgs, opts ...pulumi.InvokeOption) LookupApiIssueCommentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiIssueCommentResult, error) {
			args := v.(LookupApiIssueCommentArgs)
			r, err := LookupApiIssueComment(ctx, &args, opts...)
			var s LookupApiIssueCommentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiIssueCommentResultOutput)
}

type LookupApiIssueCommentOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	CommentId         pulumi.StringInput `pulumi:"commentId"`
	IssueId           pulumi.StringInput `pulumi:"issueId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiIssueCommentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueCommentArgs)(nil)).Elem()
}


type LookupApiIssueCommentResultOutput struct{ *pulumi.OutputState }

func (LookupApiIssueCommentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueCommentResult)(nil)).Elem()
}

func (o LookupApiIssueCommentResultOutput) ToLookupApiIssueCommentResultOutput() LookupApiIssueCommentResultOutput {
	return o
}

func (o LookupApiIssueCommentResultOutput) ToLookupApiIssueCommentResultOutputWithContext(ctx context.Context) LookupApiIssueCommentResultOutput {
	return o
}

func (o LookupApiIssueCommentResultOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o LookupApiIssueCommentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiIssueCommentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiIssueCommentResultOutput) Text() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.Text }).(pulumi.StringOutput)
}

func (o LookupApiIssueCommentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApiIssueCommentResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueCommentResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiIssueCommentResultOutput{})
}
