


package v20211201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	var rv LookupDiskResult
	err := ctx.Invoke("azure-native:compute/v20211201:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskArgs struct {
	DiskName          string `pulumi:"diskName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDiskResult struct {
	BurstingEnabled              *bool                                 `pulumi:"burstingEnabled"`
	CompletionPercent            *float64                              `pulumi:"completionPercent"`
	CreationData                 CreationDataResponse                  `pulumi:"creationData"`
	DataAccessAuthMode           *string                               `pulumi:"dataAccessAuthMode"`
	DiskAccessId                 *string                               `pulumi:"diskAccessId"`
	DiskIOPSReadOnly             *float64                              `pulumi:"diskIOPSReadOnly"`
	DiskIOPSReadWrite            *float64                              `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadOnly             *float64                              `pulumi:"diskMBpsReadOnly"`
	DiskMBpsReadWrite            *float64                              `pulumi:"diskMBpsReadWrite"`
	DiskSizeBytes                float64                               `pulumi:"diskSizeBytes"`
	DiskSizeGB                   *int                                  `pulumi:"diskSizeGB"`
	DiskState                    string                                `pulumi:"diskState"`
	Encryption                   *EncryptionResponse                   `pulumi:"encryption"`
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	ExtendedLocation             *ExtendedLocationResponse             `pulumi:"extendedLocation"`
	HyperVGeneration             *string                               `pulumi:"hyperVGeneration"`
	Id                           string                                `pulumi:"id"`
	Location                     string                                `pulumi:"location"`
	ManagedBy                    string                                `pulumi:"managedBy"`
	ManagedByExtended            []string                              `pulumi:"managedByExtended"`
	MaxShares                    *int                                  `pulumi:"maxShares"`
	Name                         string                                `pulumi:"name"`
	NetworkAccessPolicy          *string                               `pulumi:"networkAccessPolicy"`
	OsType                       *string                               `pulumi:"osType"`
	PropertyUpdatesInProgress    PropertyUpdatesInProgressResponse     `pulumi:"propertyUpdatesInProgress"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	PublicNetworkAccess          *string                               `pulumi:"publicNetworkAccess"`
	PurchasePlan                 *PurchasePlanResponse                 `pulumi:"purchasePlan"`
	SecurityProfile              *DiskSecurityProfileResponse          `pulumi:"securityProfile"`
	ShareInfo                    []ShareInfoElementResponse            `pulumi:"shareInfo"`
	Sku                          *DiskSkuResponse                      `pulumi:"sku"`
	SupportedCapabilities        *SupportedCapabilitiesResponse        `pulumi:"supportedCapabilities"`
	SupportsHibernation          *bool                                 `pulumi:"supportsHibernation"`
	Tags                         map[string]string                     `pulumi:"tags"`
	Tier                         *string                               `pulumi:"tier"`
	TimeCreated                  string                                `pulumi:"timeCreated"`
	Type                         string                                `pulumi:"type"`
	UniqueId                     string                                `pulumi:"uniqueId"`
	Zones                        []string                              `pulumi:"zones"`
}

func LookupDiskOutput(ctx *pulumi.Context, args LookupDiskOutputArgs, opts ...pulumi.InvokeOption) LookupDiskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskResult, error) {
			args := v.(LookupDiskArgs)
			r, err := LookupDisk(ctx, &args, opts...)
			var s LookupDiskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskResultOutput)
}

type LookupDiskOutputArgs struct {
	DiskName          pulumi.StringInput `pulumi:"diskName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskArgs)(nil)).Elem()
}


type LookupDiskResultOutput struct{ *pulumi.OutputState }

func (LookupDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskResult)(nil)).Elem()
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutput() LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutputWithContext(ctx context.Context) LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) BurstingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *bool { return v.BurstingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupDiskResultOutput) CompletionPercent() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.CompletionPercent }).(pulumi.Float64PtrOutput)
}

func (o LookupDiskResultOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) CreationDataResponse { return v.CreationData }).(CreationDataResponseOutput)
}

func (o LookupDiskResultOutput) DataAccessAuthMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.DataAccessAuthMode }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) DiskAccessId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.DiskAccessId }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) DiskIOPSReadOnly() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.DiskIOPSReadOnly }).(pulumi.Float64PtrOutput)
}

func (o LookupDiskResultOutput) DiskIOPSReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.DiskIOPSReadWrite }).(pulumi.Float64PtrOutput)
}

func (o LookupDiskResultOutput) DiskMBpsReadOnly() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.DiskMBpsReadOnly }).(pulumi.Float64PtrOutput)
}

func (o LookupDiskResultOutput) DiskMBpsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.DiskMBpsReadWrite }).(pulumi.Float64PtrOutput)
}

func (o LookupDiskResultOutput) DiskSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDiskResult) float64 { return v.DiskSizeBytes }).(pulumi.Float64Output)
}

func (o LookupDiskResultOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o LookupDiskResultOutput) DiskState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.DiskState }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

func (o LookupDiskResultOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *EncryptionSettingsCollectionResponse { return v.EncryptionSettingsCollection }).(EncryptionSettingsCollectionResponsePtrOutput)
}

func (o LookupDiskResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupDiskResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) ManagedByExtended() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.ManagedByExtended }).(pulumi.StringArrayOutput)
}

func (o LookupDiskResultOutput) MaxShares() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *int { return v.MaxShares }).(pulumi.IntPtrOutput)
}

func (o LookupDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) NetworkAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.NetworkAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) PropertyUpdatesInProgress() PropertyUpdatesInProgressResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) PropertyUpdatesInProgressResponse { return v.PropertyUpdatesInProgress }).(PropertyUpdatesInProgressResponseOutput)
}

func (o LookupDiskResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) PurchasePlan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *PurchasePlanResponse { return v.PurchasePlan }).(PurchasePlanResponsePtrOutput)
}

func (o LookupDiskResultOutput) SecurityProfile() DiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *DiskSecurityProfileResponse { return v.SecurityProfile }).(DiskSecurityProfileResponsePtrOutput)
}

func (o LookupDiskResultOutput) ShareInfo() ShareInfoElementResponseArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []ShareInfoElementResponse { return v.ShareInfo }).(ShareInfoElementResponseArrayOutput)
}

func (o LookupDiskResultOutput) Sku() DiskSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *DiskSkuResponse { return v.Sku }).(DiskSkuResponsePtrOutput)
}

func (o LookupDiskResultOutput) SupportedCapabilities() SupportedCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *SupportedCapabilitiesResponse { return v.SupportedCapabilities }).(SupportedCapabilitiesResponsePtrOutput)
}

func (o LookupDiskResultOutput) SupportsHibernation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *bool { return v.SupportsHibernation }).(pulumi.BoolPtrOutput)
}

func (o LookupDiskResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDiskResultOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func (o LookupDiskResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func (o LookupDiskResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskResultOutput{})
}
