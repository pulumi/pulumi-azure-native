


package v20180701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:storage/v20180701:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupStorageAccountResult struct {
	AccessTier                     string                      `pulumi:"accessTier"`
	CreationTime                   string                      `pulumi:"creationTime"`
	CustomDomain                   CustomDomainResponse        `pulumi:"customDomain"`
	EnableAzureFilesAadIntegration *bool                       `pulumi:"enableAzureFilesAadIntegration"`
	EnableHttpsTrafficOnly         *bool                       `pulumi:"enableHttpsTrafficOnly"`
	Encryption                     EncryptionResponse          `pulumi:"encryption"`
	FailoverInProgress             bool                        `pulumi:"failoverInProgress"`
	GeoReplicationStats            GeoReplicationStatsResponse `pulumi:"geoReplicationStats"`
	Id                             string                      `pulumi:"id"`
	Identity                       *IdentityResponse           `pulumi:"identity"`
	IsHnsEnabled                   *bool                       `pulumi:"isHnsEnabled"`
	Kind                           string                      `pulumi:"kind"`
	LastGeoFailoverTime            string                      `pulumi:"lastGeoFailoverTime"`
	Location                       string                      `pulumi:"location"`
	Name                           string                      `pulumi:"name"`
	NetworkRuleSet                 NetworkRuleSetResponse      `pulumi:"networkRuleSet"`
	PrimaryEndpoints               EndpointsResponse           `pulumi:"primaryEndpoints"`
	PrimaryLocation                string                      `pulumi:"primaryLocation"`
	ProvisioningState              string                      `pulumi:"provisioningState"`
	SecondaryEndpoints             EndpointsResponse           `pulumi:"secondaryEndpoints"`
	SecondaryLocation              string                      `pulumi:"secondaryLocation"`
	Sku                            SkuResponse                 `pulumi:"sku"`
	StatusOfPrimary                string                      `pulumi:"statusOfPrimary"`
	StatusOfSecondary              string                      `pulumi:"statusOfSecondary"`
	Tags                           map[string]string           `pulumi:"tags"`
	Type                           string                      `pulumi:"type"`
}
