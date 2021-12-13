


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServicePlan(ctx *pulumi.Context, args *LookupAppServicePlanArgs, opts ...pulumi.InvokeOption) (*LookupAppServicePlanResult, error) {
	var rv LookupAppServicePlanResult
	err := ctx.Invoke("azure-native:web/v20210201:getAppServicePlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppServicePlanArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAppServicePlanResult struct {
	ElasticScaleEnabled       *bool                              `pulumi:"elasticScaleEnabled"`
	ExtendedLocation          *ExtendedLocationResponse          `pulumi:"extendedLocation"`
	FreeOfferExpirationTime   *string                            `pulumi:"freeOfferExpirationTime"`
	GeoRegion                 string                             `pulumi:"geoRegion"`
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	HyperV                    *bool                              `pulumi:"hyperV"`
	Id                        string                             `pulumi:"id"`
	IsSpot                    *bool                              `pulumi:"isSpot"`
	IsXenon                   *bool                              `pulumi:"isXenon"`
	Kind                      *string                            `pulumi:"kind"`
	KubeEnvironmentProfile    *KubeEnvironmentProfileResponse    `pulumi:"kubeEnvironmentProfile"`
	Location                  string                             `pulumi:"location"`
	MaximumElasticWorkerCount *int                               `pulumi:"maximumElasticWorkerCount"`
	MaximumNumberOfWorkers    int                                `pulumi:"maximumNumberOfWorkers"`
	Name                      string                             `pulumi:"name"`
	NumberOfSites             int                                `pulumi:"numberOfSites"`
	PerSiteScaling            *bool                              `pulumi:"perSiteScaling"`
	ProvisioningState         string                             `pulumi:"provisioningState"`
	Reserved                  *bool                              `pulumi:"reserved"`
	ResourceGroup             string                             `pulumi:"resourceGroup"`
	Sku                       *SkuDescriptionResponse            `pulumi:"sku"`
	SpotExpirationTime        *string                            `pulumi:"spotExpirationTime"`
	Status                    string                             `pulumi:"status"`
	Subscription              string                             `pulumi:"subscription"`
	Tags                      map[string]string                  `pulumi:"tags"`
	TargetWorkerCount         *int                               `pulumi:"targetWorkerCount"`
	TargetWorkerSizeId        *int                               `pulumi:"targetWorkerSizeId"`
	Type                      string                             `pulumi:"type"`
	WorkerTierName            *string                            `pulumi:"workerTierName"`
	ZoneRedundant             *bool                              `pulumi:"zoneRedundant"`
}


func (val *LookupAppServicePlanResult) Defaults() *LookupAppServicePlanResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HyperV) {
		hyperV_ := false
		tmp.HyperV = &hyperV_
	}
	if isZero(tmp.IsXenon) {
		isXenon_ := false
		tmp.IsXenon = &isXenon_
	}
	if isZero(tmp.PerSiteScaling) {
		perSiteScaling_ := false
		tmp.PerSiteScaling = &perSiteScaling_
	}
	if isZero(tmp.Reserved) {
		reserved_ := false
		tmp.Reserved = &reserved_
	}
	if isZero(tmp.ZoneRedundant) {
		zoneRedundant_ := false
		tmp.ZoneRedundant = &zoneRedundant_
	}
	return &tmp
}
