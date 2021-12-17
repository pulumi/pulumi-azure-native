


package v20200801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupElasticPool(ctx *pulumi.Context, args *LookupElasticPoolArgs, opts ...pulumi.InvokeOption) (*LookupElasticPoolResult, error) {
	var rv LookupElasticPoolResult
	err := ctx.Invoke("azure-native:sql/v20200801preview:getElasticPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupElasticPoolArgs struct {
	ElasticPoolName   string `pulumi:"elasticPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupElasticPoolResult struct {
	CreationDate               string                                  `pulumi:"creationDate"`
	Id                         string                                  `pulumi:"id"`
	Kind                       string                                  `pulumi:"kind"`
	LicenseType                *string                                 `pulumi:"licenseType"`
	Location                   string                                  `pulumi:"location"`
	MaintenanceConfigurationId *string                                 `pulumi:"maintenanceConfigurationId"`
	MaxSizeBytes               *float64                                `pulumi:"maxSizeBytes"`
	Name                       string                                  `pulumi:"name"`
	PerDatabaseSettings        *ElasticPoolPerDatabaseSettingsResponse `pulumi:"perDatabaseSettings"`
	Sku                        *SkuResponse                            `pulumi:"sku"`
	State                      string                                  `pulumi:"state"`
	Tags                       map[string]string                       `pulumi:"tags"`
	Type                       string                                  `pulumi:"type"`
	ZoneRedundant              *bool                                   `pulumi:"zoneRedundant"`
}
