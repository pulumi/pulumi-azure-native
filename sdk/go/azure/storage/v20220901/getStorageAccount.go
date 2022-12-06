


package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:storage/v20220901:getStorageAccount", args, &rv, opts...)
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
	AllowedCopyScope                      *string                                        `pulumi:"allowedCopyScope"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthenticationResponse `pulumi:"azureFilesIdentityBasedAuthentication"`
	BlobRestoreStatus                     BlobRestoreStatusResponse                      `pulumi:"blobRestoreStatus"`
	CreationTime                          string                                         `pulumi:"creationTime"`
	CustomDomain                          CustomDomainResponse                           `pulumi:"customDomain"`
	DefaultToOAuthAuthentication          *bool                                          `pulumi:"defaultToOAuthAuthentication"`
	DnsEndpointType                       *string                                        `pulumi:"dnsEndpointType"`
	EnableHttpsTrafficOnly                *bool                                          `pulumi:"enableHttpsTrafficOnly"`
	EnableNfsV3                           *bool                                          `pulumi:"enableNfsV3"`
	Encryption                            EncryptionResponse                             `pulumi:"encryption"`
	ExtendedLocation                      *ExtendedLocationResponse                      `pulumi:"extendedLocation"`
	FailoverInProgress                    bool                                           `pulumi:"failoverInProgress"`
	GeoReplicationStats                   GeoReplicationStatsResponse                    `pulumi:"geoReplicationStats"`
	Id                                    string                                         `pulumi:"id"`
	Identity                              *IdentityResponse                              `pulumi:"identity"`
	ImmutableStorageWithVersioning        *ImmutableStorageAccountResponse               `pulumi:"immutableStorageWithVersioning"`
	IsHnsEnabled                          *bool                                          `pulumi:"isHnsEnabled"`
	IsLocalUserEnabled                    *bool                                          `pulumi:"isLocalUserEnabled"`
	IsSftpEnabled                         *bool                                          `pulumi:"isSftpEnabled"`
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
	PublicNetworkAccess                   *string                                        `pulumi:"publicNetworkAccess"`
	RoutingPreference                     *RoutingPreferenceResponse                     `pulumi:"routingPreference"`
	SasPolicy                             SasPolicyResponse                              `pulumi:"sasPolicy"`
	SecondaryEndpoints                    EndpointsResponse                              `pulumi:"secondaryEndpoints"`
	SecondaryLocation                     string                                         `pulumi:"secondaryLocation"`
	Sku                                   SkuResponse                                    `pulumi:"sku"`
	StatusOfPrimary                       string                                         `pulumi:"statusOfPrimary"`
	StatusOfSecondary                     string                                         `pulumi:"statusOfSecondary"`
	StorageAccountSkuConversionStatus     *StorageAccountSkuConversionStatusResponse     `pulumi:"storageAccountSkuConversionStatus"`
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

func LookupStorageAccountOutput(ctx *pulumi.Context, args LookupStorageAccountOutputArgs, opts ...pulumi.InvokeOption) LookupStorageAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageAccountResult, error) {
			args := v.(LookupStorageAccountArgs)
			r, err := LookupStorageAccount(ctx, &args, opts...)
			var s LookupStorageAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageAccountResultOutput)
}

type LookupStorageAccountOutputArgs struct {
	AccountName       pulumi.StringInput    `pulumi:"accountName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupStorageAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountArgs)(nil)).Elem()
}


type LookupStorageAccountResultOutput struct{ *pulumi.OutputState }

func (LookupStorageAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountResult)(nil)).Elem()
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutput() LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutputWithContext(ctx context.Context) LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) AccessTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.AccessTier }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) AllowBlobPublicAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.AllowBlobPublicAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) AllowCrossTenantReplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.AllowCrossTenantReplication }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) AllowSharedKeyAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.AllowSharedKeyAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) AllowedCopyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.AllowedCopyScope }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) AzureFilesIdentityBasedAuthentication() AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *AzureFilesIdentityBasedAuthenticationResponse {
		return v.AzureFilesIdentityBasedAuthentication
	}).(AzureFilesIdentityBasedAuthenticationResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) BlobRestoreStatus() BlobRestoreStatusResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) BlobRestoreStatusResponse { return v.BlobRestoreStatus }).(BlobRestoreStatusResponseOutput)
}

func (o LookupStorageAccountResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) CustomDomain() CustomDomainResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) CustomDomainResponse { return v.CustomDomain }).(CustomDomainResponseOutput)
}

func (o LookupStorageAccountResultOutput) DefaultToOAuthAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.DefaultToOAuthAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) DnsEndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.DnsEndpointType }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) EnableHttpsTrafficOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.EnableHttpsTrafficOnly }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) EnableNfsV3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.EnableNfsV3 }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) Encryption() EncryptionResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) EncryptionResponse { return v.Encryption }).(EncryptionResponseOutput)
}

func (o LookupStorageAccountResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) FailoverInProgress() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) bool { return v.FailoverInProgress }).(pulumi.BoolOutput)
}

func (o LookupStorageAccountResultOutput) GeoReplicationStats() GeoReplicationStatsResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) GeoReplicationStatsResponse { return v.GeoReplicationStats }).(GeoReplicationStatsResponseOutput)
}

func (o LookupStorageAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) ImmutableStorageWithVersioning() ImmutableStorageAccountResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *ImmutableStorageAccountResponse {
		return v.ImmutableStorageWithVersioning
	}).(ImmutableStorageAccountResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) IsHnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.IsHnsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) IsLocalUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.IsLocalUserEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) IsSftpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.IsSftpEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) KeyCreationTime() KeyCreationTimeResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) KeyCreationTimeResponse { return v.KeyCreationTime }).(KeyCreationTimeResponseOutput)
}

func (o LookupStorageAccountResultOutput) KeyPolicy() KeyPolicyResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) KeyPolicyResponse { return v.KeyPolicy }).(KeyPolicyResponseOutput)
}

func (o LookupStorageAccountResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) LargeFileSharesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.LargeFileSharesState }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) LastGeoFailoverTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.LastGeoFailoverTime }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) NetworkRuleSet() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) NetworkRuleSetResponse { return v.NetworkRuleSet }).(NetworkRuleSetResponseOutput)
}

func (o LookupStorageAccountResultOutput) PrimaryEndpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) EndpointsResponse { return v.PrimaryEndpoints }).(EndpointsResponseOutput)
}

func (o LookupStorageAccountResultOutput) PrimaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.PrimaryLocation }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupStorageAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) RoutingPreference() RoutingPreferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *RoutingPreferenceResponse { return v.RoutingPreference }).(RoutingPreferenceResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) SasPolicy() SasPolicyResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) SasPolicyResponse { return v.SasPolicy }).(SasPolicyResponseOutput)
}

func (o LookupStorageAccountResultOutput) SecondaryEndpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) EndpointsResponse { return v.SecondaryEndpoints }).(EndpointsResponseOutput)
}

func (o LookupStorageAccountResultOutput) SecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.SecondaryLocation }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupStorageAccountResultOutput) StatusOfPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.StatusOfPrimary }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) StatusOfSecondary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.StatusOfSecondary }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) StorageAccountSkuConversionStatus() StorageAccountSkuConversionStatusResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *StorageAccountSkuConversionStatusResponse {
		return v.StorageAccountSkuConversionStatus
	}).(StorageAccountSkuConversionStatusResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStorageAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountResultOutput{})
}
