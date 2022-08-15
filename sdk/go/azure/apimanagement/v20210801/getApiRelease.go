


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiRelease(ctx *pulumi.Context, args *LookupApiReleaseArgs, opts ...pulumi.InvokeOption) (*LookupApiReleaseResult, error) {
	var rv LookupApiReleaseResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:getApiRelease", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiReleaseArgs struct {
	ApiId             string `pulumi:"apiId"`
	ReleaseId         string `pulumi:"releaseId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiReleaseResult struct {
	ApiId           *string `pulumi:"apiId"`
	CreatedDateTime string  `pulumi:"createdDateTime"`
	Id              string  `pulumi:"id"`
	Name            string  `pulumi:"name"`
	Notes           *string `pulumi:"notes"`
	Type            string  `pulumi:"type"`
	UpdatedDateTime string  `pulumi:"updatedDateTime"`
}

func LookupApiReleaseOutput(ctx *pulumi.Context, args LookupApiReleaseOutputArgs, opts ...pulumi.InvokeOption) LookupApiReleaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiReleaseResult, error) {
			args := v.(LookupApiReleaseArgs)
			r, err := LookupApiRelease(ctx, &args, opts...)
			var s LookupApiReleaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiReleaseResultOutput)
}

type LookupApiReleaseOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	ReleaseId         pulumi.StringInput `pulumi:"releaseId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiReleaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiReleaseArgs)(nil)).Elem()
}


type LookupApiReleaseResultOutput struct{ *pulumi.OutputState }

func (LookupApiReleaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiReleaseResult)(nil)).Elem()
}

func (o LookupApiReleaseResultOutput) ToLookupApiReleaseResultOutput() LookupApiReleaseResultOutput {
	return o
}

func (o LookupApiReleaseResultOutput) ToLookupApiReleaseResultOutputWithContext(ctx context.Context) LookupApiReleaseResultOutput {
	return o
}

func (o LookupApiReleaseResultOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiReleaseResult) *string { return v.ApiId }).(pulumi.StringPtrOutput)
}

func (o LookupApiReleaseResultOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiReleaseResult) string { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o LookupApiReleaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiReleaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiReleaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiReleaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiReleaseResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiReleaseResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupApiReleaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiReleaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApiReleaseResultOutput) UpdatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiReleaseResult) string { return v.UpdatedDateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiReleaseResultOutput{})
}
