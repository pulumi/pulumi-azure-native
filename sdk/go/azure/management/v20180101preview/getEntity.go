


package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEntity(ctx *pulumi.Context, args *GetEntityArgs, opts ...pulumi.InvokeOption) (*GetEntityResult, error) {
	var rv GetEntityResult
	err := ctx.Invoke("azure-native:management/v20180101preview:getEntity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEntityArgs struct {
	GroupName *string `pulumi:"groupName"`
	Skiptoken *string `pulumi:"skiptoken"`
}


type GetEntityResult struct {
	NextLink string               `pulumi:"nextLink"`
	Value    []EntityInfoResponse `pulumi:"value"`
}

func GetEntityOutput(ctx *pulumi.Context, args GetEntityOutputArgs, opts ...pulumi.InvokeOption) GetEntityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEntityResult, error) {
			args := v.(GetEntityArgs)
			r, err := GetEntity(ctx, &args, opts...)
			var s GetEntityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEntityResultOutput)
}

type GetEntityOutputArgs struct {
	GroupName pulumi.StringPtrInput `pulumi:"groupName"`
	Skiptoken pulumi.StringPtrInput `pulumi:"skiptoken"`
}

func (GetEntityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityArgs)(nil)).Elem()
}


type GetEntityResultOutput struct{ *pulumi.OutputState }

func (GetEntityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityResult)(nil)).Elem()
}

func (o GetEntityResultOutput) ToGetEntityResultOutput() GetEntityResultOutput {
	return o
}

func (o GetEntityResultOutput) ToGetEntityResultOutputWithContext(ctx context.Context) GetEntityResultOutput {
	return o
}

func (o GetEntityResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o GetEntityResultOutput) Value() EntityInfoResponseArrayOutput {
	return o.ApplyT(func(v GetEntityResult) []EntityInfoResponse { return v.Value }).(EntityInfoResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEntityResultOutput{})
}
