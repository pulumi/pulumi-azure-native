


package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNamedValue(ctx *pulumi.Context, args *ListNamedValueArgs, opts ...pulumi.InvokeOption) (*ListNamedValueResult, error) {
	var rv ListNamedValueResult
	err := ctx.Invoke("azure-native:apimanagement:listNamedValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamedValueArgs struct {
	NamedValueId      string `pulumi:"namedValueId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListNamedValueResult struct {
	Value *string `pulumi:"value"`
}
