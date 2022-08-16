


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiIssue(ctx *pulumi.Context, args *LookupApiIssueArgs, opts ...pulumi.InvokeOption) (*LookupApiIssueResult, error) {
	var rv LookupApiIssueResult
	err := ctx.Invoke("azure-native:apimanagement/v20210101preview:getApiIssue", args, &rv, opts...)
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

func LookupApiIssueOutput(ctx *pulumi.Context, args LookupApiIssueOutputArgs, opts ...pulumi.InvokeOption) LookupApiIssueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiIssueResult, error) {
			args := v.(LookupApiIssueArgs)
			r, err := LookupApiIssue(ctx, &args, opts...)
			var s LookupApiIssueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiIssueResultOutput)
}

type LookupApiIssueOutputArgs struct {
	ApiId                     pulumi.StringInput  `pulumi:"apiId"`
	ExpandCommentsAttachments pulumi.BoolPtrInput `pulumi:"expandCommentsAttachments"`
	IssueId                   pulumi.StringInput  `pulumi:"issueId"`
	ResourceGroupName         pulumi.StringInput  `pulumi:"resourceGroupName"`
	ServiceName               pulumi.StringInput  `pulumi:"serviceName"`
}

func (LookupApiIssueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueArgs)(nil)).Elem()
}


type LookupApiIssueResultOutput struct{ *pulumi.OutputState }

func (LookupApiIssueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiIssueResult)(nil)).Elem()
}

func (o LookupApiIssueResultOutput) ToLookupApiIssueResultOutput() LookupApiIssueResultOutput {
	return o
}

func (o LookupApiIssueResultOutput) ToLookupApiIssueResultOutputWithContext(ctx context.Context) LookupApiIssueResultOutput {
	return o
}

func (o LookupApiIssueResultOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiIssueResult) *string { return v.ApiId }).(pulumi.StringPtrOutput)
}

func (o LookupApiIssueResultOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiIssueResult) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o LookupApiIssueResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupApiIssueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiIssueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiIssueResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiIssueResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupApiIssueResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o LookupApiIssueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApiIssueResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiIssueResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiIssueResultOutput{})
}
