


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOrderDCAccessCode(ctx *pulumi.Context, args *ListOrderDCAccessCodeArgs, opts ...pulumi.InvokeOption) (*ListOrderDCAccessCodeResult, error) {
	var rv ListOrderDCAccessCodeResult
	err := ctx.Invoke("azure-native:databoxedge/v20221201preview:listOrderDCAccessCode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOrderDCAccessCodeArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListOrderDCAccessCodeResult struct {
	AuthCode *string `pulumi:"authCode"`
}

func ListOrderDCAccessCodeOutput(ctx *pulumi.Context, args ListOrderDCAccessCodeOutputArgs, opts ...pulumi.InvokeOption) ListOrderDCAccessCodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOrderDCAccessCodeResult, error) {
			args := v.(ListOrderDCAccessCodeArgs)
			r, err := ListOrderDCAccessCode(ctx, &args, opts...)
			var s ListOrderDCAccessCodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOrderDCAccessCodeResultOutput)
}

type ListOrderDCAccessCodeOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListOrderDCAccessCodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOrderDCAccessCodeArgs)(nil)).Elem()
}


type ListOrderDCAccessCodeResultOutput struct{ *pulumi.OutputState }

func (ListOrderDCAccessCodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOrderDCAccessCodeResult)(nil)).Elem()
}

func (o ListOrderDCAccessCodeResultOutput) ToListOrderDCAccessCodeResultOutput() ListOrderDCAccessCodeResultOutput {
	return o
}

func (o ListOrderDCAccessCodeResultOutput) ToListOrderDCAccessCodeResultOutputWithContext(ctx context.Context) ListOrderDCAccessCodeResultOutput {
	return o
}

func (o ListOrderDCAccessCodeResultOutput) AuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOrderDCAccessCodeResult) *string { return v.AuthCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOrderDCAccessCodeResultOutput{})
}
