


package v20180701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:storage/v20180701:getStorageAccount", args, &rv, opts...)
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


func (val *LookupStorageAccountResult) Defaults() *LookupStorageAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = *tmp.Encryption.Defaults()

	tmp.NetworkRuleSet = *tmp.NetworkRuleSet.Defaults()

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

func (o LookupStorageAccountResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) CustomDomain() CustomDomainResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) CustomDomainResponse { return v.CustomDomain }).(CustomDomainResponseOutput)
}

func (o LookupStorageAccountResultOutput) EnableAzureFilesAadIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.EnableAzureFilesAadIntegration }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) EnableHttpsTrafficOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.EnableHttpsTrafficOnly }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) Encryption() EncryptionResponseOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) EncryptionResponse { return v.Encryption }).(EncryptionResponseOutput)
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

func (o LookupStorageAccountResultOutput) IsHnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *bool { return v.IsHnsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupStorageAccountResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) LastGeoFailoverTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.LastGeoFailoverTime }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Location }).(pulumi.StringOutput)
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

func (o LookupStorageAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
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

func (o LookupStorageAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStorageAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountResultOutput{})
}
