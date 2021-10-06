


package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagerActivationKey(ctx *pulumi.Context, args *ListManagerActivationKeyArgs, opts ...pulumi.InvokeOption) (*ListManagerActivationKeyResult, error) {
	var rv ListManagerActivationKeyResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:listManagerActivationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagerActivationKeyArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListManagerActivationKeyResult struct {
	ActivationKey string `pulumi:"activationKey"`
}
