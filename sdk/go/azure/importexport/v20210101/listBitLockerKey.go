


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBitLockerKey(ctx *pulumi.Context, args *ListBitLockerKeyArgs, opts ...pulumi.InvokeOption) (*ListBitLockerKeyResult, error) {
	var rv ListBitLockerKeyResult
	err := ctx.Invoke("azure-native:importexport/v20210101:listBitLockerKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBitLockerKeyArgs struct {
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListBitLockerKeyResult struct {
	Value []DriveBitLockerKeyResponse `pulumi:"value"`
}

func ListBitLockerKeyOutput(ctx *pulumi.Context, args ListBitLockerKeyOutputArgs, opts ...pulumi.InvokeOption) ListBitLockerKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBitLockerKeyResult, error) {
			args := v.(ListBitLockerKeyArgs)
			r, err := ListBitLockerKey(ctx, &args, opts...)
			var s ListBitLockerKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBitLockerKeyResultOutput)
}

type ListBitLockerKeyOutputArgs struct {
	JobName           pulumi.StringInput `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListBitLockerKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBitLockerKeyArgs)(nil)).Elem()
}


type ListBitLockerKeyResultOutput struct{ *pulumi.OutputState }

func (ListBitLockerKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBitLockerKeyResult)(nil)).Elem()
}

func (o ListBitLockerKeyResultOutput) ToListBitLockerKeyResultOutput() ListBitLockerKeyResultOutput {
	return o
}

func (o ListBitLockerKeyResultOutput) ToListBitLockerKeyResultOutputWithContext(ctx context.Context) ListBitLockerKeyResultOutput {
	return o
}

func (o ListBitLockerKeyResultOutput) Value() DriveBitLockerKeyResponseArrayOutput {
	return o.ApplyT(func(v ListBitLockerKeyResult) []DriveBitLockerKeyResponse { return v.Value }).(DriveBitLockerKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBitLockerKeyResultOutput{})
}
