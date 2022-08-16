


package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGlobalUserOperationStatus(ctx *pulumi.Context, args *GetGlobalUserOperationStatusArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserOperationStatusResult, error) {
	var rv GetGlobalUserOperationStatusResult
	err := ctx.Invoke("azure-native:labservices:getGlobalUserOperationStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserOperationStatusArgs struct {
	OperationUrl string `pulumi:"operationUrl"`
	UserName     string `pulumi:"userName"`
}


type GetGlobalUserOperationStatusResult struct {
	Status string `pulumi:"status"`
}

func GetGlobalUserOperationStatusOutput(ctx *pulumi.Context, args GetGlobalUserOperationStatusOutputArgs, opts ...pulumi.InvokeOption) GetGlobalUserOperationStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGlobalUserOperationStatusResult, error) {
			args := v.(GetGlobalUserOperationStatusArgs)
			r, err := GetGlobalUserOperationStatus(ctx, &args, opts...)
			var s GetGlobalUserOperationStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGlobalUserOperationStatusResultOutput)
}

type GetGlobalUserOperationStatusOutputArgs struct {
	OperationUrl pulumi.StringInput `pulumi:"operationUrl"`
	UserName     pulumi.StringInput `pulumi:"userName"`
}

func (GetGlobalUserOperationStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserOperationStatusArgs)(nil)).Elem()
}


type GetGlobalUserOperationStatusResultOutput struct{ *pulumi.OutputState }

func (GetGlobalUserOperationStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserOperationStatusResult)(nil)).Elem()
}

func (o GetGlobalUserOperationStatusResultOutput) ToGetGlobalUserOperationStatusResultOutput() GetGlobalUserOperationStatusResultOutput {
	return o
}

func (o GetGlobalUserOperationStatusResultOutput) ToGetGlobalUserOperationStatusResultOutputWithContext(ctx context.Context) GetGlobalUserOperationStatusResultOutput {
	return o
}

func (o GetGlobalUserOperationStatusResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetGlobalUserOperationStatusResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGlobalUserOperationStatusResultOutput{})
}
