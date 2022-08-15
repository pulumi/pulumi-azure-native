


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiIssueAttachment(ctx *pulumi.Context, args *LookupApiIssueAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueAttachmentResult, error) {
	var rv LookupApiIssueAttachmentResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201:getApiIssueAttachment", args, &rv, opts...)
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

func LookupApiIssueAttachmentOutput(ctx *pulumi.Context, args LookupApiIssueAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupApiIssueAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiIssueAttachmentResult, error) {
			args := v.(LookupApiIssueAttachmentArgs)
			r, err := LookupApiIssueAttachment(ctx, &args, opts...)
			var s LookupApiIssueAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiIssueAttachmentResultOutput)
}

type LookupApiIssueAttachmentOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	AttachmentId      pulumi.StringInput `pulumi:"attachmentId"`
	IssueId           pulumi.StringInput `pulumi:"issueId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiIssueAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueAttachmentArgs)(nil)).Elem()
}


type LookupApiIssueAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupApiIssueAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueAttachmentResult)(nil)).Elem()
}

func (o LookupApiIssueAttachmentResultOutput) ToLookupApiIssueAttachmentResultOutput() LookupApiIssueAttachmentResultOutput {
	return o
}

func (o LookupApiIssueAttachmentResultOutput) ToLookupApiIssueAttachmentResultOutputWithContext(ctx context.Context) LookupApiIssueAttachmentResultOutput {
	return o
}

func (o LookupApiIssueAttachmentResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueAttachmentResult) string { return v.Content }).(pulumi.StringOutput)
}

func (o LookupApiIssueAttachmentResultOutput) ContentFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueAttachmentResult) string { return v.ContentFormat }).(pulumi.StringOutput)
}

func (o LookupApiIssueAttachmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueAttachmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiIssueAttachmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueAttachmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiIssueAttachmentResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueAttachmentResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o LookupApiIssueAttachmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueAttachmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiIssueAttachmentResultOutput{})
}
