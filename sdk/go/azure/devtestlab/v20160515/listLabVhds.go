


package v20160515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListLabVhds(ctx *pulumi.Context, args *ListLabVhdsArgs, opts ...pulumi.InvokeOption) (*ListLabVhdsResult, error) {
	var rv ListLabVhdsResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:listLabVhds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLabVhdsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListLabVhdsResult struct {
	NextLink *string          `pulumi:"nextLink"`
	Value    []LabVhdResponse `pulumi:"value"`
}
