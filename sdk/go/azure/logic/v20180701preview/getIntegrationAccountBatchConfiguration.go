


package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountBatchConfiguration(ctx *pulumi.Context, args *LookupIntegrationAccountBatchConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountBatchConfigurationResult, error) {
	var rv LookupIntegrationAccountBatchConfigurationResult
	err := ctx.Invoke("azure-native:logic/v20180701preview:getIntegrationAccountBatchConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountBatchConfigurationArgs struct {
	BatchConfigurationName string `pulumi:"batchConfigurationName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountBatchConfigurationResult struct {
	Id         string                               `pulumi:"id"`
	Location   *string                              `pulumi:"location"`
	Name       string                               `pulumi:"name"`
	Properties BatchConfigurationPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                    `pulumi:"tags"`
	Type       string                               `pulumi:"type"`
}
