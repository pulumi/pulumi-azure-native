


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContentType(ctx *pulumi.Context, args *LookupContentTypeArgs, opts ...pulumi.InvokeOption) (*LookupContentTypeResult, error) {
	var rv LookupContentTypeResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:getContentType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentTypeArgs struct {
	ContentTypeId     string `pulumi:"contentTypeId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupContentTypeResult struct {
	Description *string     `pulumi:"description"`
	Id          string      `pulumi:"id"`
	Name        string      `pulumi:"name"`
	Schema      interface{} `pulumi:"schema"`
	Type        string      `pulumi:"type"`
	Version     *string     `pulumi:"version"`
}

func LookupContentTypeOutput(ctx *pulumi.Context, args LookupContentTypeOutputArgs, opts ...pulumi.InvokeOption) LookupContentTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentTypeResult, error) {
			args := v.(LookupContentTypeArgs)
			r, err := LookupContentType(ctx, &args, opts...)
			var s LookupContentTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentTypeResultOutput)
}

type LookupContentTypeOutputArgs struct {
	ContentTypeId     pulumi.StringInput `pulumi:"contentTypeId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupContentTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentTypeArgs)(nil)).Elem()
}


type LookupContentTypeResultOutput struct{ *pulumi.OutputState }

func (LookupContentTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentTypeResult)(nil)).Elem()
}

func (o LookupContentTypeResultOutput) ToLookupContentTypeResultOutput() LookupContentTypeResultOutput {
	return o
}

func (o LookupContentTypeResultOutput) ToLookupContentTypeResultOutputWithContext(ctx context.Context) LookupContentTypeResultOutput {
	return o
}

func (o LookupContentTypeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTypeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupContentTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContentTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContentTypeResultOutput) Schema() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupContentTypeResult) interface{} { return v.Schema }).(pulumi.AnyOutput)
}

func (o LookupContentTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupContentTypeResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContentTypeResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentTypeResultOutput{})
}
