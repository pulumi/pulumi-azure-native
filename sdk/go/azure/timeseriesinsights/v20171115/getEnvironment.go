


package v20171115

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	var rv LookupEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20171115:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupEnvironmentResult struct {
	CreationTime                 string                         `pulumi:"creationTime"`
	DataAccessFqdn               string                         `pulumi:"dataAccessFqdn"`
	DataAccessId                 string                         `pulumi:"dataAccessId"`
	DataRetentionTime            string                         `pulumi:"dataRetentionTime"`
	Id                           string                         `pulumi:"id"`
	Location                     string                         `pulumi:"location"`
	Name                         string                         `pulumi:"name"`
	PartitionKeyProperties       []PartitionKeyPropertyResponse `pulumi:"partitionKeyProperties"`
	ProvisioningState            string                         `pulumi:"provisioningState"`
	Sku                          *SkuResponse                   `pulumi:"sku"`
	Status                       EnvironmentStatusResponse      `pulumi:"status"`
	StorageLimitExceededBehavior *string                        `pulumi:"storageLimitExceededBehavior"`
	Tags                         map[string]string              `pulumi:"tags"`
	Type                         string                         `pulumi:"type"`
}
