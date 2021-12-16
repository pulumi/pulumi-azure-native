


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

func init() {
	pulumi.RegisterOutputType(BatchAccountOutput{})
}
