


package v20180815preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStandardEnvironment(ctx *pulumi.Context, args *LookupStandardEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupStandardEnvironmentResult, error) {
	var rv LookupStandardEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20180815preview:getStandardEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStandardEnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupStandardEnvironmentResult struct {
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
	Tags                         map[string]string              `pulumi:"tags"`
	Type                         string                         `pulumi:"type"`
}
