


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerEndpoint(ctx *pulumi.Context, args *LookupServerEndpointArgs, opts ...pulumi.InvokeOption) (*LookupServerEndpointResult, error) {
	var rv LookupServerEndpointResult
	err := ctx.Invoke("azure-native:storagesync/v20200301:getServerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerEndpointArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerEndpointName     string `pulumi:"serverEndpointName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
	SyncGroupName          string `pulumi:"syncGroupName"`
}


type LookupServerEndpointResult struct {
	CloudTiering                                *string                                  `pulumi:"cloudTiering"`
	CloudTieringStatus                          ServerEndpointCloudTieringStatusResponse `pulumi:"cloudTieringStatus"`
	FriendlyName                                *string                                  `pulumi:"friendlyName"`
	Id                                          string                                   `pulumi:"id"`
	InitialDownloadPolicy                       *string                                  `pulumi:"initialDownloadPolicy"`
	LastOperationName                           string                                   `pulumi:"lastOperationName"`
	LastWorkflowId                              string                                   `pulumi:"lastWorkflowId"`
	LocalCacheMode                              *string                                  `pulumi:"localCacheMode"`
	Name                                        string                                   `pulumi:"name"`
	OfflineDataTransfer                         *string                                  `pulumi:"offlineDataTransfer"`
	OfflineDataTransferShareName                *string                                  `pulumi:"offlineDataTransferShareName"`
	OfflineDataTransferStorageAccountResourceId string                                   `pulumi:"offlineDataTransferStorageAccountResourceId"`
	OfflineDataTransferStorageAccountTenantId   string                                   `pulumi:"offlineDataTransferStorageAccountTenantId"`
	ProvisioningState                           string                                   `pulumi:"provisioningState"`
	RecallStatus                                ServerEndpointRecallStatusResponse       `pulumi:"recallStatus"`
	ServerLocalPath                             *string                                  `pulumi:"serverLocalPath"`
	ServerResourceId                            *string                                  `pulumi:"serverResourceId"`
	SyncStatus                                  ServerEndpointSyncStatusResponse         `pulumi:"syncStatus"`
	TierFilesOlderThanDays                      *int                                     `pulumi:"tierFilesOlderThanDays"`
	Type                                        string                                   `pulumi:"type"`
	VolumeFreeSpacePercent                      *int                                     `pulumi:"volumeFreeSpacePercent"`
}

func LookupServerEndpointOutput(ctx *pulumi.Context, args LookupServerEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupServerEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerEndpointResult, error) {
			args := v.(LookupServerEndpointArgs)
			r, err := LookupServerEndpoint(ctx, &args, opts...)
			var s LookupServerEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerEndpointResultOutput)
}

type LookupServerEndpointOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerEndpointName     pulumi.StringInput `pulumi:"serverEndpointName"`
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
	SyncGroupName          pulumi.StringInput `pulumi:"syncGroupName"`
}

func (LookupServerEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerEndpointArgs)(nil)).Elem()
}


type LookupServerEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupServerEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerEndpointResult)(nil)).Elem()
}

func (o LookupServerEndpointResultOutput) ToLookupServerEndpointResultOutput() LookupServerEndpointResultOutput {
	return o
}

func (o LookupServerEndpointResultOutput) ToLookupServerEndpointResultOutputWithContext(ctx context.Context) LookupServerEndpointResultOutput {
	return o
}

func (o LookupServerEndpointResultOutput) CloudTiering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.CloudTiering }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) CloudTieringStatus() ServerEndpointCloudTieringStatusResponseOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) ServerEndpointCloudTieringStatusResponse {
		return v.CloudTieringStatus
	}).(ServerEndpointCloudTieringStatusResponseOutput)
}

func (o LookupServerEndpointResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) InitialDownloadPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.InitialDownloadPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) LastOperationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.LastOperationName }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) LastWorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.LastWorkflowId }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) LocalCacheMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.LocalCacheMode }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) OfflineDataTransfer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.OfflineDataTransfer }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) OfflineDataTransferShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.OfflineDataTransferShareName }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) OfflineDataTransferStorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.OfflineDataTransferStorageAccountResourceId }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) OfflineDataTransferStorageAccountTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.OfflineDataTransferStorageAccountTenantId }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) RecallStatus() ServerEndpointRecallStatusResponseOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) ServerEndpointRecallStatusResponse { return v.RecallStatus }).(ServerEndpointRecallStatusResponseOutput)
}

func (o LookupServerEndpointResultOutput) ServerLocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.ServerLocalPath }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) ServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *string { return v.ServerResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupServerEndpointResultOutput) SyncStatus() ServerEndpointSyncStatusResponseOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) ServerEndpointSyncStatusResponse { return v.SyncStatus }).(ServerEndpointSyncStatusResponseOutput)
}

func (o LookupServerEndpointResultOutput) TierFilesOlderThanDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.TierFilesOlderThanDays }).(pulumi.IntPtrOutput)
}

func (o LookupServerEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServerEndpointResultOutput) VolumeFreeSpacePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerEndpointResult) *int { return v.VolumeFreeSpacePercent }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerEndpointResultOutput{})
}
