


package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("azure-native:compute:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SnapshotName      string `pulumi:"snapshotName"`
}


type LookupSnapshotResult struct {
	CreationData                 CreationDataResponse                  `pulumi:"creationData"`
	DiskAccessId                 *string                               `pulumi:"diskAccessId"`
	DiskSizeBytes                float64                               `pulumi:"diskSizeBytes"`
	DiskSizeGB                   *int                                  `pulumi:"diskSizeGB"`
	DiskState                    string                                `pulumi:"diskState"`
	Encryption                   *EncryptionResponse                   `pulumi:"encryption"`
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	ExtendedLocation             *ExtendedLocationResponse             `pulumi:"extendedLocation"`
	HyperVGeneration             *string                               `pulumi:"hyperVGeneration"`
	Id                           string                                `pulumi:"id"`
	Incremental                  *bool                                 `pulumi:"incremental"`
	Location                     string                                `pulumi:"location"`
	ManagedBy                    string                                `pulumi:"managedBy"`
	Name                         string                                `pulumi:"name"`
	NetworkAccessPolicy          *string                               `pulumi:"networkAccessPolicy"`
	OsType                       *string                               `pulumi:"osType"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	PurchasePlan                 *PurchasePlanResponse                 `pulumi:"purchasePlan"`
	Sku                          *SnapshotSkuResponse                  `pulumi:"sku"`
	SupportsHibernation          *bool                                 `pulumi:"supportsHibernation"`
	Tags                         map[string]string                     `pulumi:"tags"`
	TimeCreated                  string                                `pulumi:"timeCreated"`
	Type                         string                                `pulumi:"type"`
	UniqueId                     string                                `pulumi:"uniqueId"`
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResult, error) {
			args := v.(LookupSnapshotArgs)
			r, err := LookupSnapshot(ctx, &args, opts...)
			var s LookupSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotResultOutput)
}

type LookupSnapshotOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SnapshotName      pulumi.StringInput `pulumi:"snapshotName"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}


type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v LookupSnapshotResult) CreationDataResponse { return v.CreationData }).(CreationDataResponseOutput)
}

func (o LookupSnapshotResultOutput) DiskAccessId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.DiskAccessId }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) DiskSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSnapshotResult) float64 { return v.DiskSizeBytes }).(pulumi.Float64Output)
}

func (o LookupSnapshotResultOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o LookupSnapshotResultOutput) DiskState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.DiskState }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

func (o LookupSnapshotResultOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *EncryptionSettingsCollectionResponse {
		return v.EncryptionSettingsCollection
	}).(EncryptionSettingsCollectionResponsePtrOutput)
}

func (o LookupSnapshotResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupSnapshotResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Incremental() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *bool { return v.Incremental }).(pulumi.BoolPtrOutput)
}

func (o LookupSnapshotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) NetworkAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.NetworkAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) PurchasePlan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *PurchasePlanResponse { return v.PurchasePlan }).(PurchasePlanResponsePtrOutput)
}

func (o LookupSnapshotResultOutput) Sku() SnapshotSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *SnapshotSkuResponse { return v.Sku }).(SnapshotSkuResponsePtrOutput)
}

func (o LookupSnapshotResultOutput) SupportsHibernation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *bool { return v.SupportsHibernation }).(pulumi.BoolPtrOutput)
}

func (o LookupSnapshotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSnapshotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSnapshotResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
