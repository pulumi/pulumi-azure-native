


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagByOperation(ctx *pulumi.Context, args *LookupTagByOperationArgs, opts ...pulumi.InvokeOption) (*LookupTagByOperationResult, error) {
	var rv LookupTagByOperationResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:getTagByOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagByOperationArgs struct {
	ApiId             string `pulumi:"apiId"`
	OperationId       string `pulumi:"operationId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagId             string `pulumi:"tagId"`
}


type LookupTagByOperationResult struct {
	DisplayName string `pulumi:"displayName"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
}

func LookupTagByOperationOutput(ctx *pulumi.Context, args LookupTagByOperationOutputArgs, opts ...pulumi.InvokeOption) LookupTagByOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagByOperationResult, error) {
			args := v.(LookupTagByOperationArgs)
			r, err := LookupTagByOperation(ctx, &args, opts...)
			var s LookupTagByOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagByOperationResultOutput)
}

type LookupTagByOperationOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	OperationId       pulumi.StringInput `pulumi:"operationId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	TagId             pulumi.StringInput `pulumi:"tagId"`
}

func (LookupTagByOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagByOperationArgs)(nil)).Elem()
}


type LookupTagByOperationResultOutput struct{ *pulumi.OutputState }

func (LookupTagByOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagByOperationResult)(nil)).Elem()
}

func (o LookupTagByOperationResultOutput) ToLookupTagByOperationResultOutput() LookupTagByOperationResultOutput {
	return o
}

func (o LookupTagByOperationResultOutput) ToLookupTagByOperationResultOutputWithContext(ctx context.Context) LookupTagByOperationResultOutput {
	return o
}

func (o LookupTagByOperationResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByOperationResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupTagByOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagByOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagByOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagByOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagByOperationResultOutput{})
}
