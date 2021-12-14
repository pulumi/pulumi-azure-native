


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:storage/v20210401:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupStorageAccountArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupStorageAccountResult struct {
	AccessTier                            string                                         `pulumi:"accessTier"`
	AllowBlobPublicAccess                 *bool                                          `pulumi:"allowBlobPublicAccess"`
	AllowCrossTenantReplication           *bool                                          `pulumi:"allowCrossTenantReplication"`
	AllowSharedKeyAccess                  *bool                                          `pulumi:"allowSharedKeyAccess"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthenticationResponse `pulumi:"azureFilesIdentityBasedAuthentication"`
	BlobRestoreStatus                     BlobRestoreStatusResponse                      `pulumi:"blobRestoreStatus"`
	CreationTime                          string                                         `pulumi:"creationTime"`
	CustomDomain                          CustomDomainResponse                           `pulumi:"customDomain"`
	EnableHttpsTrafficOnly                *bool                                          `pulumi:"enableHttpsTrafficOnly"`
	EnableNfsV3                           *bool                                          `pulumi:"enableNfsV3"`
	Encryption                            EncryptionResponse                             `pulumi:"encryption"`
	ExtendedLocation                      *ExtendedLocationResponse                      `pulumi:"extendedLocation"`
	FailoverInProgress                    bool                                           `pulumi:"failoverInProgress"`
	GeoReplicationStats                   GeoReplicationStatsResponse                    `pulumi:"geoReplicationStats"`
	Id                                    string                                         `pulumi:"id"`
	Identity                              *IdentityResponse                              `pulumi:"identity"`
	IsHnsEnabled                          *bool                                          `pulumi:"isHnsEnabled"`
	KeyCreationTime                       KeyCreationTimeResponse                        `pulumi:"keyCreationTime"`
	KeyPolicy                             KeyPolicyResponse                              `pulumi:"keyPolicy"`
	Kind                                  string                                         `pulumi:"kind"`
	LargeFileSharesState                  *string                                        `pulumi:"largeFileSharesState"`
	LastGeoFailoverTime                   string                                         `pulumi:"lastGeoFailoverTime"`
	Location                              string                                         `pulumi:"location"`
	MinimumTlsVersion                     *string                                        `pulumi:"minimumTlsVersion"`
	Name                                  string                                         `pulumi:"name"`
	NetworkRuleSet                        NetworkRuleSetResponse                         `pulumi:"networkRuleSet"`
	PrimaryEndpoints                      EndpointsResponse                              `pulumi:"primaryEndpoints"`
	PrimaryLocation                       string                                         `pulumi:"primaryLocation"`
	PrivateEndpointConnections            []PrivateEndpointConnectionResponse            `pulumi:"privateEndpointConnections"`
	ProvisioningState                     string                                         `pulumi:"provisioningState"`
	RoutingPreference                     *RoutingPreferenceResponse                     `pulumi:"routingPreference"`
	SasPolicy                             SasPolicyResponse                              `pulumi:"sasPolicy"`
	SecondaryEndpoints                    EndpointsResponse                              `pulumi:"secondaryEndpoints"`
	SecondaryLocation                     string                                         `pulumi:"secondaryLocation"`
	Sku                                   SkuResponse                                    `pulumi:"sku"`
	StatusOfPrimary                       string                                         `pulumi:"statusOfPrimary"`
	StatusOfSecondary                     string                                         `pulumi:"statusOfSecondary"`
	Tags                                  map[string]string                              `pulumi:"tags"`
	Type                                  string                                         `pulumi:"type"`
}


func (val *LookupStorageAccountResult) Defaults() *LookupStorageAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = *tmp.Encryption.Defaults()

	tmp.NetworkRuleSet = *tmp.NetworkRuleSet.Defaults()

	tmp.SasPolicy = *tmp.SasPolicy.Defaults()

	return &tmp
}
