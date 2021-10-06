


package v20180815preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLongTermEnvironment(ctx *pulumi.Context, args *LookupLongTermEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupLongTermEnvironmentResult, error) {
	var rv LookupLongTermEnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights/v20180815preview:getLongTermEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLongTermEnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLongTermEnvironmentResult struct {
	CreationTime           string                                     `pulumi:"creationTime"`
	DataAccessFqdn         string                                     `pulumi:"dataAccessFqdn"`
	DataAccessId           string                                     `pulumi:"dataAccessId"`
	Id                     string                                     `pulumi:"id"`
	Kind                   string                                     `pulumi:"kind"`
	Location               string                                     `pulumi:"location"`
	Name                   string                                     `pulumi:"name"`
	ProvisioningState      string                                     `pulumi:"provisioningState"`
	Sku                    SkuResponse                                `pulumi:"sku"`
	Status                 EnvironmentStatusResponse                  `pulumi:"status"`
	StorageConfiguration   LongTermStorageConfigurationOutputResponse `pulumi:"storageConfiguration"`
	Tags                   map[string]string                          `pulumi:"tags"`
	TimeSeriesIdProperties []TimeSeriesIdPropertyResponse             `pulumi:"timeSeriesIdProperties"`
	Type                   string                                     `pulumi:"type"`
	WarmStoreConfiguration *WarmStoreConfigurationPropertiesResponse  `pulumi:"warmStoreConfiguration"`
}
