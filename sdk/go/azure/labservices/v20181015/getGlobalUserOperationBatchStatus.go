


package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGlobalUserOperationBatchStatus(ctx *pulumi.Context, args *GetGlobalUserOperationBatchStatusArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserOperationBatchStatusResult, error) {
	var rv GetGlobalUserOperationBatchStatusResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getGlobalUserOperationBatchStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserOperationBatchStatusArgs struct {
	Urls     []string `pulumi:"urls"`
	UserName string   `pulumi:"userName"`
}


type GetGlobalUserOperationBatchStatusResult struct {
	Items []OperationBatchStatusResponseItemResponse `pulumi:"items"`
}

func GetGlobalUserOperationBatchStatusOutput(ctx *pulumi.Context, args GetGlobalUserOperationBatchStatusOutputArgs, opts ...pulumi.InvokeOption) GetGlobalUserOperationBatchStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGlobalUserOperationBatchStatusResult, error) {
			args := v.(GetGlobalUserOperationBatchStatusArgs)
			r, err := GetGlobalUserOperationBatchStatus(ctx, &args, opts...)
			var s GetGlobalUserOperationBatchStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGlobalUserOperationBatchStatusResultOutput)
}

type GetGlobalUserOperationBatchStatusOutputArgs struct {
	Urls     pulumi.StringArrayInput `pulumi:"urls"`
	UserName pulumi.StringInput      `pulumi:"userName"`
}

func (GetGlobalUserOperationBatchStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserOperationBatchStatusArgs)(nil)).Elem()
}


type GetGlobalUserOperationBatchStatusResultOutput struct{ *pulumi.OutputState }

func (GetGlobalUserOperationBatchStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserOperationBatchStatusResult)(nil)).Elem()
}

func (o GetGlobalUserOperationBatchStatusResultOutput) ToGetGlobalUserOperationBatchStatusResultOutput() GetGlobalUserOperationBatchStatusResultOutput {
	return o
}

func (o GetGlobalUserOperationBatchStatusResultOutput) ToGetGlobalUserOperationBatchStatusResultOutputWithContext(ctx context.Context) GetGlobalUserOperationBatchStatusResultOutput {
	return o
}

func (o GetGlobalUserOperationBatchStatusResultOutput) Items() OperationBatchStatusResponseItemResponseArrayOutput {
	return o.ApplyT(func(v GetGlobalUserOperationBatchStatusResult) []OperationBatchStatusResponseItemResponse {
		return v.Items
	}).(OperationBatchStatusResponseItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGlobalUserOperationBatchStatusResultOutput{})
}
