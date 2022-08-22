


package v20170605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type CloudEndpoint struct {
	pulumi.CustomResourceState

	BackupEnabled            pulumi.BoolOutput      `pulumi:"backupEnabled"`
	FriendlyName             pulumi.StringPtrOutput `pulumi:"friendlyName"`
	LastWorkflowId           pulumi.StringPtrOutput `pulumi:"lastWorkflowId"`
	Name                     pulumi.StringOutput    `pulumi:"name"`
	PartnershipId            pulumi.StringPtrOutput `pulumi:"partnershipId"`
	ProvisioningState        pulumi.StringPtrOutput `pulumi:"provisioningState"`
	StorageAccount           pulumi.StringPtrOutput `pulumi:"storageAccount"`
	StorageAccountKey        pulumi.StringPtrOutput `pulumi:"storageAccountKey"`
	StorageAccountResourceId pulumi.StringPtrOutput `pulumi:"storageAccountResourceId"`
	StorageAccountShareName  pulumi.StringPtrOutput `pulumi:"storageAccountShareName"`
	StorageAccountTenantId   pulumi.StringPtrOutput `pulumi:"storageAccountTenantId"`
	Type                     pulumi.StringOutput    `pulumi:"type"`
}


func NewCloudEndpoint(ctx *pulumi.Context,
	name string, args *CloudEndpointArgs, opts ...pulumi.ResourceOption) (*CloudEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageSyncServiceName == nil {
		return nil, errors.New("invalid value for required argument 'StorageSyncServiceName'")
	}
	if args.SyncGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SyncGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagesync:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180402:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:CloudEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200901:CloudEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource CloudEndpoint
	err := ctx.RegisterResource("azure-native:storagesync/v20170605preview:CloudEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCloudEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudEndpointState, opts ...pulumi.ResourceOption) (*CloudEndpoint, error) {
	var resource CloudEndpoint
	err := ctx.ReadResource("azure-native:storagesync/v20170605preview:CloudEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cloudEndpointState struct {
}

type CloudEndpointState struct {
}

func (CloudEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudEndpointState)(nil)).Elem()
}

type cloudEndpointArgs struct {
	CloudEndpointName        *string `pulumi:"cloudEndpointName"`
	FriendlyName             *string `pulumi:"friendlyName"`
	LastWorkflowId           *string `pulumi:"lastWorkflowId"`
	PartnershipId            *string `pulumi:"partnershipId"`
	ProvisioningState        *string `pulumi:"provisioningState"`
	ResourceGroupName        string  `pulumi:"resourceGroupName"`
	StorageAccount           *string `pulumi:"storageAccount"`
	StorageAccountKey        *string `pulumi:"storageAccountKey"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
	StorageAccountShareName  *string `pulumi:"storageAccountShareName"`
	StorageAccountTenantId   *string `pulumi:"storageAccountTenantId"`
	StorageSyncServiceName   string  `pulumi:"storageSyncServiceName"`
	SyncGroupName            string  `pulumi:"syncGroupName"`
}


type CloudEndpointArgs struct {
	CloudEndpointName        pulumi.StringPtrInput
	FriendlyName             pulumi.StringPtrInput
	LastWorkflowId           pulumi.StringPtrInput
	PartnershipId            pulumi.StringPtrInput
	ProvisioningState        pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	StorageAccount           pulumi.StringPtrInput
	StorageAccountKey        pulumi.StringPtrInput
	StorageAccountResourceId pulumi.StringPtrInput
	StorageAccountShareName  pulumi.StringPtrInput
	StorageAccountTenantId   pulumi.StringPtrInput
	StorageSyncServiceName   pulumi.StringInput
	SyncGroupName            pulumi.StringInput
}

func (CloudEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudEndpointArgs)(nil)).Elem()
}

type CloudEndpointInput interface {
	pulumi.Input

	ToCloudEndpointOutput() CloudEndpointOutput
	ToCloudEndpointOutputWithContext(ctx context.Context) CloudEndpointOutput
}

func (*CloudEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudEndpoint)(nil)).Elem()
}

func (i *CloudEndpoint) ToCloudEndpointOutput() CloudEndpointOutput {
	return i.ToCloudEndpointOutputWithContext(context.Background())
}

func (i *CloudEndpoint) ToCloudEndpointOutputWithContext(ctx context.Context) CloudEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudEndpointOutput)
}

type CloudEndpointOutput struct{ *pulumi.OutputState }

func (CloudEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudEndpoint)(nil)).Elem()
}

func (o CloudEndpointOutput) ToCloudEndpointOutput() CloudEndpointOutput {
	return o
}

func (o CloudEndpointOutput) ToCloudEndpointOutputWithContext(ctx context.Context) CloudEndpointOutput {
	return o
}

func (o CloudEndpointOutput) BackupEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.BoolOutput { return v.BackupEnabled }).(pulumi.BoolOutput)
}

func (o CloudEndpointOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) LastWorkflowId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.LastWorkflowId }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudEndpointOutput) PartnershipId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.PartnershipId }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.StorageAccountKey }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) StorageAccountShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.StorageAccountShareName }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) StorageAccountTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringPtrOutput { return v.StorageAccountTenantId }).(pulumi.StringPtrOutput)
}

func (o CloudEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudEndpointOutput{})
}
