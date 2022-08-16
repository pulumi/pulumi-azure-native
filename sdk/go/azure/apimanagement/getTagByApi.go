


package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagByApi(ctx *pulumi.Context, args *LookupTagByApiArgs, opts ...pulumi.InvokeOption) (*LookupTagByApiResult, error) {
	var rv LookupTagByApiResult
	err := ctx.Invoke("azure-native:apimanagement:getTagByApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagByApiArgs struct {
	ApiId             string `pulumi:"apiId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagId             string `pulumi:"tagId"`
}


type LookupTagByApiResult struct {
	DisplayName string `pulumi:"displayName"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
}

func LookupTagByApiOutput(ctx *pulumi.Context, args LookupTagByApiOutputArgs, opts ...pulumi.InvokeOption) LookupTagByApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagByApiResult, error) {
			args := v.(LookupTagByApiArgs)
			r, err := LookupTagByApi(ctx, &args, opts...)
			var s LookupTagByApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagByApiResultOutput)
}

type LookupTagByApiOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	TagId             pulumi.StringInput `pulumi:"tagId"`
}

func (LookupTagByApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagByApiArgs)(nil)).Elem()
}


type LookupTagByApiResultOutput struct{ *pulumi.OutputState }

func (LookupTagByApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagByApiResult)(nil)).Elem()
}

func (o LookupTagByApiResultOutput) ToLookupTagByApiResultOutput() LookupTagByApiResultOutput {
	return o
}

func (o LookupTagByApiResultOutput) ToLookupTagByApiResultOutputWithContext(ctx context.Context) LookupTagByApiResultOutput {
	return o
}

func (o LookupTagByApiResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByApiResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupTagByApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByApiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagByApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByApiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagByApiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByApiResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagByApiResultOutput{})
}
