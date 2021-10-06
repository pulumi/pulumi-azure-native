


package timeseriesinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGen2Environment(ctx *pulumi.Context, args *LookupGen2EnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupGen2EnvironmentResult, error) {
	var rv LookupGen2EnvironmentResult
	err := ctx.Invoke("azure-native:timeseriesinsights:getGen2Environment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGen2EnvironmentArgs struct {
	EnvironmentName   string  `pulumi:"environmentName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupGen2EnvironmentResult struct {
	CreationTime           string                                    `pulumi:"creationTime"`
	DataAccessFqdn         string                                    `pulumi:"dataAccessFqdn"`
	DataAccessId           string                                    `pulumi:"dataAccessId"`
	Id                     string                                    `pulumi:"id"`
	Kind                   string                                    `pulumi:"kind"`
	Location               string                                    `pulumi:"location"`
	Name                   string                                    `pulumi:"name"`
	ProvisioningState      string                                    `pulumi:"provisioningState"`
	Sku                    SkuResponse                               `pulumi:"sku"`
	Status                 EnvironmentStatusResponse                 `pulumi:"status"`
	StorageConfiguration   Gen2StorageConfigurationOutputResponse    `pulumi:"storageConfiguration"`
	Tags                   map[string]string                         `pulumi:"tags"`
	TimeSeriesIdProperties []TimeSeriesIdPropertyResponse            `pulumi:"timeSeriesIdProperties"`
	Type                   string                                    `pulumi:"type"`
	WarmStoreConfiguration *WarmStoreConfigurationPropertiesResponse `pulumi:"warmStoreConfiguration"`
}
