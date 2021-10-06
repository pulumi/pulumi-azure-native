


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerFarm(ctx *pulumi.Context, args *LookupServerFarmArgs, opts ...pulumi.InvokeOption) (*LookupServerFarmResult, error) {
	var rv LookupServerFarmResult
	err := ctx.Invoke("azure-native:web/v20150801:getServerFarm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerFarmArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupServerFarmResult struct {
	AdminSiteName             *string                            `pulumi:"adminSiteName"`
	GeoRegion                 string                             `pulumi:"geoRegion"`
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	Id                        *string                            `pulumi:"id"`
	Kind                      *string                            `pulumi:"kind"`
	Location                  string                             `pulumi:"location"`
	MaximumNumberOfWorkers    *int                               `pulumi:"maximumNumberOfWorkers"`
	Name                      *string                            `pulumi:"name"`
	NumberOfSites             int                                `pulumi:"numberOfSites"`
	PerSiteScaling            *bool                              `pulumi:"perSiteScaling"`
	Reserved                  *bool                              `pulumi:"reserved"`
	ResourceGroup             string                             `pulumi:"resourceGroup"`
	Sku                       *SkuDescriptionResponse            `pulumi:"sku"`
	Status                    string                             `pulumi:"status"`
	Subscription              string                             `pulumi:"subscription"`
	Tags                      map[string]string                  `pulumi:"tags"`
	Type                      *string                            `pulumi:"type"`
	WorkerTierName            *string                            `pulumi:"workerTierName"`
}
