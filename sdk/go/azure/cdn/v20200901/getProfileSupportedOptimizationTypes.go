


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProfileSupportedOptimizationTypes(ctx *pulumi.Context, args *GetProfileSupportedOptimizationTypesArgs, opts ...pulumi.InvokeOption) (*GetProfileSupportedOptimizationTypesResult, error) {
	var rv GetProfileSupportedOptimizationTypesResult
	err := ctx.Invoke("azure-native:cdn/v20200901:getProfileSupportedOptimizationTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProfileSupportedOptimizationTypesArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetProfileSupportedOptimizationTypesResult struct {
	SupportedOptimizationTypes []string `pulumi:"supportedOptimizationTypes"`
}
