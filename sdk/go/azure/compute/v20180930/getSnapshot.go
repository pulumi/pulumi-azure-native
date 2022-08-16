


package v20180930

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("azure-native:compute/v20180930:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSnapshotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SnapshotName      string `pulumi:"snapshotName"`
}


type LookupSnapshotResult struct {
	CreationData                 CreationDataResponse                  `pulumi:"creationData"`
	DiskSizeGB                   *int                                  `pulumi:"diskSizeGB"`
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             *string                               `pulumi:"hyperVGeneration"`
	Id                           string                                `pulumi:"id"`
	Location                     string                                `pulumi:"location"`
	ManagedBy                    string                                `pulumi:"managedBy"`
	Name                         string                                `pulumi:"name"`
	OsType                       *string                               `pulumi:"osType"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	Sku                          *SnapshotSkuResponse                  `pulumi:"sku"`
	Tags                         map[string]string                     `pulumi:"tags"`
	TimeCreated                  string                                `pulumi:"timeCreated"`
	Type                         string                                `pulumi:"type"`
}


func (val *LookupSnapshotResult) Defaults() *LookupSnapshotResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
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

func (o LookupSnapshotResultOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o LookupSnapshotResultOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *EncryptionSettingsCollectionResponse {
		return v.EncryptionSettingsCollection
	}).(EncryptionSettingsCollectionResponsePtrOutput)
}

func (o LookupSnapshotResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
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

func (o LookupSnapshotResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupSnapshotResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSnapshotResultOutput) Sku() SnapshotSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *SnapshotSkuResponse { return v.Sku }).(SnapshotSkuResponsePtrOutput)
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

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
