


package databoxedge

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOrderDCAccessCode(ctx *pulumi.Context, args *ListOrderDCAccessCodeArgs, opts ...pulumi.InvokeOption) (*ListOrderDCAccessCodeResult, error) {
	var rv ListOrderDCAccessCodeResult
	err := ctx.Invoke("azure-native:databoxedge:listOrderDCAccessCode", args, &rv, opts...)
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
