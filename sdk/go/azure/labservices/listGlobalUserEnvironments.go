


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListGlobalUserEnvironments(ctx *pulumi.Context, args *ListGlobalUserEnvironmentsArgs, opts ...pulumi.InvokeOption) (*ListGlobalUserEnvironmentsResult, error) {
	var rv ListGlobalUserEnvironmentsResult
	err := ctx.Invoke("azure-native:labservices:listGlobalUserEnvironments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGlobalUserEnvironmentsArgs struct {
	LabId    *string `pulumi:"labId"`
	UserName string  `pulumi:"userName"`
}


type ListGlobalUserEnvironmentsResult struct {
	Environments []EnvironmentDetailsResponse `pulumi:"environments"`
}
