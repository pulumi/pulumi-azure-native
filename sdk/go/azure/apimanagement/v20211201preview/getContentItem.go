


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContentItem(ctx *pulumi.Context, args *LookupContentItemArgs, opts ...pulumi.InvokeOption) (*LookupContentItemResult, error) {
	var rv LookupContentItemResult
	err := ctx.Invoke("azure-native:apimanagement/v20211201preview:getContentItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentItemArgs struct {
	ContentItemId     string `pulumi:"contentItemId"`
	ContentTypeId     string `pulumi:"contentTypeId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupContentItemResult struct {
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

func LookupContentItemOutput(ctx *pulumi.Context, args LookupContentItemOutputArgs, opts ...pulumi.InvokeOption) LookupContentItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentItemResult, error) {
			args := v.(LookupContentItemArgs)
			r, err := LookupContentItem(ctx, &args, opts...)
			var s LookupContentItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentItemResultOutput)
}

type LookupContentItemOutputArgs struct {
	ContentItemId     pulumi.StringInput `pulumi:"contentItemId"`
	ContentTypeId     pulumi.StringInput `pulumi:"contentTypeId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupContentItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentItemArgs)(nil)).Elem()
}


type LookupContentItemResultOutput struct{ *pulumi.OutputState }

func (LookupContentItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentItemResult)(nil)).Elem()
}

func (o LookupContentItemResultOutput) ToLookupContentItemResultOutput() LookupContentItemResultOutput {
	return o
}

func (o LookupContentItemResultOutput) ToLookupContentItemResultOutputWithContext(ctx context.Context) LookupContentItemResultOutput {
	return o
}

func (o LookupContentItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentItemResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContentItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentItemResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContentItemResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupContentItemResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupContentItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentItemResultOutput{})
}
