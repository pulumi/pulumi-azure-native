


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerEndpoint struct {
	pulumi.CustomResourceState

	CloudTiering                                pulumi.StringPtrOutput                         `pulumi:"cloudTiering"`
	CloudTieringStatus                          ServerEndpointCloudTieringStatusResponseOutput `pulumi:"cloudTieringStatus"`
	FriendlyName                                pulumi.StringPtrOutput                         `pulumi:"friendlyName"`
	InitialDownloadPolicy                       pulumi.StringPtrOutput                         `pulumi:"initialDownloadPolicy"`
	InitialUploadPolicy                         pulumi.StringPtrOutput                         `pulumi:"initialUploadPolicy"`
	LastOperationName                           pulumi.StringOutput                            `pulumi:"lastOperationName"`
	LastWorkflowId                              pulumi.StringOutput                            `pulumi:"lastWorkflowId"`
	LocalCacheMode                              pulumi.StringPtrOutput                         `pulumi:"localCacheMode"`
	Name                                        pulumi.StringOutput                            `pulumi:"name"`
	OfflineDataTransfer                         pulumi.StringPtrOutput                         `pulumi:"offlineDataTransfer"`
	OfflineDataTransferShareName                pulumi.StringPtrOutput                         `pulumi:"offlineDataTransferShareName"`
	OfflineDataTransferStorageAccountResourceId pulumi.StringOutput                            `pulumi:"offlineDataTransferStorageAccountResourceId"`
	OfflineDataTransferStorageAccountTenantId   pulumi.StringOutput                            `pulumi:"offlineDataTransferStorageAccountTenantId"`
	ProvisioningState                           pulumi.StringOutput                            `pulumi:"provisioningState"`
	RecallStatus                                ServerEndpointRecallStatusResponseOutput       `pulumi:"recallStatus"`
	ServerLocalPath                             pulumi.StringPtrOutput                         `pulumi:"serverLocalPath"`
	ServerName                                  pulumi.StringOutput                            `pulumi:"serverName"`
	ServerResourceId                            pulumi.StringPtrOutput                         `pulumi:"serverResourceId"`
	SyncStatus                                  ServerEndpointSyncStatusResponseOutput         `pulumi:"syncStatus"`
	SystemData                                  SystemDataResponseOutput                       `pulumi:"systemData"`
	TierFilesOlderThanDays                      pulumi.IntPtrOutput                            `pulumi:"tierFilesOlderThanDays"`
	Type                                        pulumi.StringOutput                            `pulumi:"type"`
	VolumeFreeSpacePercent                      pulumi.IntPtrOutput                            `pulumi:"volumeFreeSpacePercent"`
}


func NewServerEndpoint(ctx *pulumi.Context,
	name string, args *ServerEndpointArgs, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
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
	if isZero(args.InitialDownloadPolicy) {
		args.InitialDownloadPolicy = pulumi.StringPtr("NamespaceThenModifiedFiles")
	}
	if isZero(args.InitialUploadPolicy) {
		args.InitialUploadPolicy = pulumi.StringPtr("Merge")
	}
	if isZero(args.LocalCacheMode) {
		args.LocalCacheMode = pulumi.StringPtr("UpdateLocallyCachedFiles")
	}
	if isZero(args.TierFilesOlderThanDays) {
		args.TierFilesOlderThanDays = pulumi.IntPtr(0)
	}
	if isZero(args.VolumeFreeSpacePercent) {
		args.VolumeFreeSpacePercent = pulumi.IntPtr(20)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagesync:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20170605preview:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180402:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:ServerEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20220601:ServerEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerEndpoint
	err := ctx.RegisterResource("azure-native:storagesync/v20200901:ServerEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerEndpointState, opts ...pulumi.ResourceOption) (*ServerEndpoint, error) {
	var resource ServerEndpoint
	err := ctx.ReadResource("azure-native:storagesync/v20200901:ServerEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverEndpointState struct {
}

type ServerEndpointState struct {
}

func (ServerEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverEndpointState)(nil)).Elem()
}

type serverEndpointArgs struct {
	CloudTiering                 *string `pulumi:"cloudTiering"`
	FriendlyName                 *string `pulumi:"friendlyName"`
	InitialDownloadPolicy        *string `pulumi:"initialDownloadPolicy"`
	InitialUploadPolicy          *string `pulumi:"initialUploadPolicy"`
	LocalCacheMode               *string `pulumi:"localCacheMode"`
	OfflineDataTransfer          *string `pulumi:"offlineDataTransfer"`
	OfflineDataTransferShareName *string `pulumi:"offlineDataTransferShareName"`
	ResourceGroupName            string  `pulumi:"resourceGroupName"`
	ServerEndpointName           *string `pulumi:"serverEndpointName"`
	ServerLocalPath              *string `pulumi:"serverLocalPath"`
	ServerResourceId             *string `pulumi:"serverResourceId"`
	StorageSyncServiceName       string  `pulumi:"storageSyncServiceName"`
	SyncGroupName                string  `pulumi:"syncGroupName"`
	TierFilesOlderThanDays       *int    `pulumi:"tierFilesOlderThanDays"`
	VolumeFreeSpacePercent       *int    `pulumi:"volumeFreeSpacePercent"`
}


type ServerEndpointArgs struct {
	CloudTiering                 pulumi.StringPtrInput
	FriendlyName                 pulumi.StringPtrInput
	InitialDownloadPolicy        pulumi.StringPtrInput
	InitialUploadPolicy          pulumi.StringPtrInput
	LocalCacheMode               pulumi.StringPtrInput
	OfflineDataTransfer          pulumi.StringPtrInput
	OfflineDataTransferShareName pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	ServerEndpointName           pulumi.StringPtrInput
	ServerLocalPath              pulumi.StringPtrInput
	ServerResourceId             pulumi.StringPtrInput
	StorageSyncServiceName       pulumi.StringInput
	SyncGroupName                pulumi.StringInput
	TierFilesOlderThanDays       pulumi.IntPtrInput
	VolumeFreeSpacePercent       pulumi.IntPtrInput
}

func (ServerEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverEndpointArgs)(nil)).Elem()
}

type ServerEndpointInput interface {
	pulumi.Input

	ToServerEndpointOutput() ServerEndpointOutput
	ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput
}

func (*ServerEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpoint)(nil)).Elem()
}

func (i *ServerEndpoint) ToServerEndpointOutput() ServerEndpointOutput {
	return i.ToServerEndpointOutputWithContext(context.Background())
}

func (i *ServerEndpoint) ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointOutput)
}

type ServerEndpointOutput struct{ *pulumi.OutputState }

func (ServerEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpoint)(nil)).Elem()
}

func (o ServerEndpointOutput) ToServerEndpointOutput() ServerEndpointOutput {
	return o
}

func (o ServerEndpointOutput) ToServerEndpointOutputWithContext(ctx context.Context) ServerEndpointOutput {
	return o
}

func (o ServerEndpointOutput) CloudTiering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.CloudTiering }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) CloudTieringStatus() ServerEndpointCloudTieringStatusResponseOutput {
	return o.ApplyT(func(v *ServerEndpoint) ServerEndpointCloudTieringStatusResponseOutput { return v.CloudTieringStatus }).(ServerEndpointCloudTieringStatusResponseOutput)
}

func (o ServerEndpointOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) InitialDownloadPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.InitialDownloadPolicy }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) InitialUploadPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.InitialUploadPolicy }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) LastOperationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.LastOperationName }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) LastWorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.LastWorkflowId }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) LocalCacheMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.LocalCacheMode }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) OfflineDataTransfer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.OfflineDataTransfer }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) OfflineDataTransferShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.OfflineDataTransferShareName }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) OfflineDataTransferStorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.OfflineDataTransferStorageAccountResourceId }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) OfflineDataTransferStorageAccountTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.OfflineDataTransferStorageAccountTenantId }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) RecallStatus() ServerEndpointRecallStatusResponseOutput {
	return o.ApplyT(func(v *ServerEndpoint) ServerEndpointRecallStatusResponseOutput { return v.RecallStatus }).(ServerEndpointRecallStatusResponseOutput)
}

func (o ServerEndpointOutput) ServerLocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.ServerLocalPath }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) ServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringPtrOutput { return v.ServerResourceId }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointOutput) SyncStatus() ServerEndpointSyncStatusResponseOutput {
	return o.ApplyT(func(v *ServerEndpoint) ServerEndpointSyncStatusResponseOutput { return v.SyncStatus }).(ServerEndpointSyncStatusResponseOutput)
}

func (o ServerEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServerEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ServerEndpointOutput) TierFilesOlderThanDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.TierFilesOlderThanDays }).(pulumi.IntPtrOutput)
}

func (o ServerEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ServerEndpointOutput) VolumeFreeSpacePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpoint) pulumi.IntPtrOutput { return v.VolumeFreeSpacePercent }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerEndpointOutput{})
}
