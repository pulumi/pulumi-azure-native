


package v20170801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListServerGatewayStatus(ctx *pulumi.Context, args *ListServerGatewayStatusArgs, opts ...pulumi.InvokeOption) (*ListServerGatewayStatusResult, error) {
	var rv ListServerGatewayStatusResult
	err := ctx.Invoke("azure-native:analysisservices/v20170801:listServerGatewayStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServerGatewayStatusArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type ListServerGatewayStatusResult struct {
	Status *int `pulumi:"status"`
}

func ListServerGatewayStatusOutput(ctx *pulumi.Context, args ListServerGatewayStatusOutputArgs, opts ...pulumi.InvokeOption) ListServerGatewayStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServerGatewayStatusResult, error) {
			args := v.(ListServerGatewayStatusArgs)
			r, err := ListServerGatewayStatus(ctx, &args, opts...)
			var s ListServerGatewayStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServerGatewayStatusResultOutput)
}

type ListServerGatewayStatusOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (ListServerGatewayStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServerGatewayStatusArgs)(nil)).Elem()
}


type ListServerGatewayStatusResultOutput struct{ *pulumi.OutputState }

func (ListServerGatewayStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServerGatewayStatusResult)(nil)).Elem()
}

func (o ListServerGatewayStatusResultOutput) ToListServerGatewayStatusResultOutput() ListServerGatewayStatusResultOutput {
	return o
}

func (o ListServerGatewayStatusResultOutput) ToListServerGatewayStatusResultOutputWithContext(ctx context.Context) ListServerGatewayStatusResultOutput {
	return o
}

func (o ListServerGatewayStatusResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ListServerGatewayStatusResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServerGatewayStatusResultOutput{})
}
