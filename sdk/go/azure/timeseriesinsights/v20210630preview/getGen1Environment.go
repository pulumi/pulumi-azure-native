


package v20210630preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGen1Environment(ctx *pulumi.Context, args *LookupGen1EnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupGen1EnvironmentResult, error) {
	var rv LookupGen1EnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20210630preview:getGen1Environment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGen1EnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupGen1EnvironmentResult struct {
	CreationTime                 string                         `pulumi:"creationTime"`
	DataAccessFqdn               string                         `pulumi:"dataAccessFqdn"`
	DataAccessId                 string                         `pulumi:"dataAccessId"`
	DataRetentionTime            string                         `pulumi:"dataRetentionTime"`
	Id                           string                         `pulumi:"id"`
	Kind                         string                         `pulumi:"kind"`
	Location                     string                         `pulumi:"location"`
	Name                         string                         `pulumi:"name"`
	PartitionKeyProperties       []TimeSeriesIdPropertyResponse `pulumi:"partitionKeyProperties"`
	ProvisioningState            string                         `pulumi:"provisioningState"`
	Sku                          SkuResponse                    `pulumi:"sku"`
	Status                       EnvironmentStatusResponse      `pulumi:"status"`
	StorageLimitExceededBehavior *string                        `pulumi:"storageLimitExceededBehavior"`
	SupportsCustomerManagedKey   bool                           `pulumi:"supportsCustomerManagedKey"`
	Tags                         map[string]string              `pulumi:"tags"`
	Type                         string                         `pulumi:"type"`
}
