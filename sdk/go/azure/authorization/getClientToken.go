


package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetClientToken(ctx *pulumi.Context, args *GetClientTokenArgs, opts ...pulumi.InvokeOption) (*GetClientTokenResult, error) {
	var rv GetClientTokenResult
	err := ctx.Invoke("azure-native:authorization:getClientToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetClientTokenArgs struct {
	Endpoint *string `pulumi:"endpoint"`
}


type GetClientTokenResult struct {
	Token string `pulumi:"token"`
}
