


package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BatchAccount struct {
	pulumi.CustomResourceState

	AccountEndpoint                       pulumi.StringOutput                              `pulumi:"accountEndpoint"`
	ActiveJobAndJobScheduleQuota          pulumi.IntOutput                                 `pulumi:"activeJobAndJobScheduleQuota"`
	AutoStorage                           AutoStoragePropertiesResponseOutput              `pulumi:"autoStorage"`
	DedicatedCoreQuota                    pulumi.IntOutput                                 `pulumi:"dedicatedCoreQuota"`
	DedicatedCoreQuotaPerVMFamily         VirtualMachineFamilyCoreQuotaResponseArrayOutput `pulumi:"dedicatedCoreQuotaPerVMFamily"`
	DedicatedCoreQuotaPerVMFamilyEnforced pulumi.BoolOutput                                `pulumi:"dedicatedCoreQuotaPerVMFamilyEnforced"`
	Encryption                            EncryptionPropertiesResponseOutput               `pulumi:"encryption"`
	Identity                              BatchAccountIdentityResponsePtrOutput            `pulumi:"identity"`
	KeyVaultReference                     KeyVaultReferenceResponseOutput                  `pulumi:"keyVaultReference"`
	Location                              pulumi.StringOutput                              `pulumi:"location"`
	LowPriorityCoreQuota                  pulumi.IntOutput                                 `pulumi:"lowPriorityCoreQuota"`
	Name                                  pulumi.StringOutput                              `pulumi:"name"`
	PoolAllocationMode                    pulumi.StringOutput                              `pulumi:"poolAllocationMode"`
	PoolQuota                             pulumi.IntOutput                                 `pulumi:"poolQuota"`
	PrivateEndpointConnections            PrivateEndpointConnectionResponseArrayOutput     `pulumi:"privateEndpointConnections"`
	ProvisioningState                     pulumi.StringOutput                              `pulumi:"provisioningState"`
	PublicNetworkAccess                   pulumi.StringOutput                              `pulumi:"publicNetworkAccess"`
	Tags                                  pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                                  pulumi.StringOutput                              `pulumi:"type"`
}


func NewBatchAccount(ctx *pulumi.Context,
	name string, args *BatchAccountArgs, opts ...pulumi.ResourceOption) (*BatchAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:batch/v20151201:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170101:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170501:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220101:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220601:BatchAccount"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20221001:BatchAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource BatchAccount
	err := ctx.RegisterResource("azure-native:batch:BatchAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBatchAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BatchAccountState, opts ...pulumi.ResourceOption) (*BatchAccount, error) {
	var resource BatchAccount
	err := ctx.ReadResource("azure-native:batch:BatchAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type batchAccountState struct {
}

type BatchAccountState struct {
}

func (BatchAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*batchAccountState)(nil)).Elem()
}

type batchAccountArgs struct {
	AccountName         *string                    `pulumi:"accountName"`
	AutoStorage         *AutoStorageBaseProperties `pulumi:"autoStorage"`
	Encryption          *EncryptionProperties      `pulumi:"encryption"`
	Identity            *BatchAccountIdentity      `pulumi:"identity"`
	KeyVaultReference   *KeyVaultReference         `pulumi:"keyVaultReference"`
	Location            *string                    `pulumi:"location"`
	PoolAllocationMode  *PoolAllocationMode        `pulumi:"poolAllocationMode"`
	PublicNetworkAccess *PublicNetworkAccessType   `pulumi:"publicNetworkAccess"`
	ResourceGroupName   string                     `pulumi:"resourceGroupName"`
	Tags                map[string]string          `pulumi:"tags"`
}


type BatchAccountArgs struct {
	AccountName         pulumi.StringPtrInput
	AutoStorage         AutoStorageBasePropertiesPtrInput
	Encryption          EncryptionPropertiesPtrInput
	Identity            BatchAccountIdentityPtrInput
	KeyVaultReference   KeyVaultReferencePtrInput
	Location            pulumi.StringPtrInput
	PoolAllocationMode  PoolAllocationModePtrInput
	PublicNetworkAccess PublicNetworkAccessTypePtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
}

func (BatchAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*batchAccountArgs)(nil)).Elem()
}

type BatchAccountInput interface {
	pulumi.Input

	ToBatchAccountOutput() BatchAccountOutput
	ToBatchAccountOutputWithContext(ctx context.Context) BatchAccountOutput
}

func (*BatchAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchAccount)(nil)).Elem()
}

func (i *BatchAccount) ToBatchAccountOutput() BatchAccountOutput {
	return i.ToBatchAccountOutputWithContext(context.Background())
}

func (i *BatchAccount) ToBatchAccountOutputWithContext(ctx context.Context) BatchAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchAccountOutput)
}

type BatchAccountOutput struct{ *pulumi.OutputState }

func (BatchAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchAccount)(nil)).Elem()
}

func (o BatchAccountOutput) ToBatchAccountOutput() BatchAccountOutput {
	return o
}

func (o BatchAccountOutput) ToBatchAccountOutputWithContext(ctx context.Context) BatchAccountOutput {
	return o
}

func (o BatchAccountOutput) AccountEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.AccountEndpoint }).(pulumi.StringOutput)
}

func (o BatchAccountOutput) ActiveJobAndJobScheduleQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.IntOutput { return v.ActiveJobAndJobScheduleQuota }).(pulumi.IntOutput)
}

func (o BatchAccountOutput) AutoStorage() AutoStoragePropertiesResponseOutput {
	return o.ApplyT(func(v *BatchAccount) AutoStoragePropertiesResponseOutput { return v.AutoStorage }).(AutoStoragePropertiesResponseOutput)
}

func (o BatchAccountOutput) DedicatedCoreQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.IntOutput { return v.DedicatedCoreQuota }).(pulumi.IntOutput)
}

func (o BatchAccountOutput) DedicatedCoreQuotaPerVMFamily() VirtualMachineFamilyCoreQuotaResponseArrayOutput {
	return o.ApplyT(func(v *BatchAccount) VirtualMachineFamilyCoreQuotaResponseArrayOutput {
		return v.DedicatedCoreQuotaPerVMFamily
	}).(VirtualMachineFamilyCoreQuotaResponseArrayOutput)
}

func (o BatchAccountOutput) DedicatedCoreQuotaPerVMFamilyEnforced() pulumi.BoolOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.BoolOutput { return v.DedicatedCoreQuotaPerVMFamilyEnforced }).(pulumi.BoolOutput)
}

func (o BatchAccountOutput) Encryption() EncryptionPropertiesResponseOutput {
	return o.ApplyT(func(v *BatchAccount) EncryptionPropertiesResponseOutput { return v.Encryption }).(EncryptionPropertiesResponseOutput)
}

func (o BatchAccountOutput) Identity() BatchAccountIdentityResponsePtrOutput {
	return o.ApplyT(func(v *BatchAccount) BatchAccountIdentityResponsePtrOutput { return v.Identity }).(BatchAccountIdentityResponsePtrOutput)
}

func (o BatchAccountOutput) KeyVaultReference() KeyVaultReferenceResponseOutput {
	return o.ApplyT(func(v *BatchAccount) KeyVaultReferenceResponseOutput { return v.KeyVaultReference }).(KeyVaultReferenceResponseOutput)
}

func (o BatchAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BatchAccountOutput) LowPriorityCoreQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.IntOutput { return v.LowPriorityCoreQuota }).(pulumi.IntOutput)
}

func (o BatchAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BatchAccountOutput) PoolAllocationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.PoolAllocationMode }).(pulumi.StringOutput)
}

func (o BatchAccountOutput) PoolQuota() pulumi.IntOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.IntOutput { return v.PoolQuota }).(pulumi.IntOutput)
}

func (o BatchAccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *BatchAccount) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o BatchAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BatchAccountOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func (o BatchAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o BatchAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BatchAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BatchAccountOutput{})
}
