


package v20181015

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGlobalUserEnvironment(ctx *pulumi.Context, args *GetGlobalUserEnvironmentArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserEnvironmentResult, error) {
	var rv GetGlobalUserEnvironmentResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getGlobalUserEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserEnvironmentArgs struct {
	EnvironmentId string  `pulumi:"environmentId"`
	Expand        *string `pulumi:"expand"`
	UserName      string  `pulumi:"userName"`
}


type GetGlobalUserEnvironmentResult struct {
	Environment EnvironmentDetailsResponse `pulumi:"environment"`
}
