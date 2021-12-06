


package operationsmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementConfiguration(ctx *pulumi.Context, args *LookupManagementConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupManagementConfigurationResult, error) {
	var rv LookupManagementConfigurationResult
	err := ctx.Invoke("azure-native:operationsmanagement:getManagementConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementConfigurationArgs struct {
	ManagementConfigurationName string `pulumi:"managementConfigurationName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
}


type LookupManagementConfigurationResult struct {
	Id         string                                    `pulumi:"id"`
	Location   *string                                   `pulumi:"location"`
	Name       string                                    `pulumi:"name"`
	Properties ManagementConfigurationPropertiesResponse `pulumi:"properties"`
	Type       string                                    `pulumi:"type"`
}
