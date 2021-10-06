


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerEndpointHealthResponse struct {
	CombinedHealth            *string                     `pulumi:"combinedHealth"`
	CurrentProgress           *SyncProgressStatusResponse `pulumi:"currentProgress"`
	DownloadHealth            *string                     `pulumi:"downloadHealth"`
	DownloadStatus            *SyncSessionStatusResponse  `pulumi:"downloadStatus"`
	LastUpdatedTimestamp      *string                     `pulumi:"lastUpdatedTimestamp"`
	OfflineDataTransferStatus *string                     `pulumi:"offlineDataTransferStatus"`
	UploadHealth              *string                     `pulumi:"uploadHealth"`
	UploadStatus              *SyncSessionStatusResponse  `pulumi:"uploadStatus"`
}





type ServerEndpointHealthResponseInput interface {
	pulumi.Input

	ToServerEndpointHealthResponseOutput() ServerEndpointHealthResponseOutput
	ToServerEndpointHealthResponseOutputWithContext(context.Context) ServerEndpointHealthResponseOutput
}

type ServerEndpointHealthResponseArgs struct {
	CombinedHealth            pulumi.StringPtrInput              `pulumi:"combinedHealth"`
	CurrentProgress           SyncProgressStatusResponsePtrInput `pulumi:"currentProgress"`
	DownloadHealth            pulumi.StringPtrInput              `pulumi:"downloadHealth"`
	DownloadStatus            SyncSessionStatusResponsePtrInput  `pulumi:"downloadStatus"`
	LastUpdatedTimestamp      pulumi.StringPtrInput              `pulumi:"lastUpdatedTimestamp"`
	OfflineDataTransferStatus pulumi.StringPtrInput              `pulumi:"offlineDataTransferStatus"`
	UploadHealth              pulumi.StringPtrInput              `pulumi:"uploadHealth"`
	UploadStatus              SyncSessionStatusResponsePtrInput  `pulumi:"uploadStatus"`
}

func (ServerEndpointHealthResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointHealthResponse)(nil)).Elem()
}

func (i ServerEndpointHealthResponseArgs) ToServerEndpointHealthResponseOutput() ServerEndpointHealthResponseOutput {
	return i.ToServerEndpointHealthResponseOutputWithContext(context.Background())
}

func (i ServerEndpointHealthResponseArgs) ToServerEndpointHealthResponseOutputWithContext(ctx context.Context) ServerEndpointHealthResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointHealthResponseOutput)
}

func (i ServerEndpointHealthResponseArgs) ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput {
	return i.ToServerEndpointHealthResponsePtrOutputWithContext(context.Background())
}

func (i ServerEndpointHealthResponseArgs) ToServerEndpointHealthResponsePtrOutputWithContext(ctx context.Context) ServerEndpointHealthResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointHealthResponseOutput).ToServerEndpointHealthResponsePtrOutputWithContext(ctx)
}









type ServerEndpointHealthResponsePtrInput interface {
	pulumi.Input

	ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput
	ToServerEndpointHealthResponsePtrOutputWithContext(context.Context) ServerEndpointHealthResponsePtrOutput
}

type serverEndpointHealthResponsePtrType ServerEndpointHealthResponseArgs

func ServerEndpointHealthResponsePtr(v *ServerEndpointHealthResponseArgs) ServerEndpointHealthResponsePtrInput {
	return (*serverEndpointHealthResponsePtrType)(v)
}

func (*serverEndpointHealthResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointHealthResponse)(nil)).Elem()
}

func (i *serverEndpointHealthResponsePtrType) ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput {
	return i.ToServerEndpointHealthResponsePtrOutputWithContext(context.Background())
}

func (i *serverEndpointHealthResponsePtrType) ToServerEndpointHealthResponsePtrOutputWithContext(ctx context.Context) ServerEndpointHealthResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointHealthResponsePtrOutput)
}

type ServerEndpointHealthResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointHealthResponse)(nil)).Elem()
}

func (o ServerEndpointHealthResponseOutput) ToServerEndpointHealthResponseOutput() ServerEndpointHealthResponseOutput {
	return o
}

func (o ServerEndpointHealthResponseOutput) ToServerEndpointHealthResponseOutputWithContext(ctx context.Context) ServerEndpointHealthResponseOutput {
	return o
}

func (o ServerEndpointHealthResponseOutput) ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput {
	return o.ToServerEndpointHealthResponsePtrOutputWithContext(context.Background())
}

func (o ServerEndpointHealthResponseOutput) ToServerEndpointHealthResponsePtrOutputWithContext(ctx context.Context) ServerEndpointHealthResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerEndpointHealthResponse) *ServerEndpointHealthResponse {
		return &v
	}).(ServerEndpointHealthResponsePtrOutput)
}

func (o ServerEndpointHealthResponseOutput) CombinedHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.CombinedHealth }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponseOutput) CurrentProgress() SyncProgressStatusResponsePtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *SyncProgressStatusResponse { return v.CurrentProgress }).(SyncProgressStatusResponsePtrOutput)
}

func (o ServerEndpointHealthResponseOutput) DownloadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.DownloadHealth }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponseOutput) DownloadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *SyncSessionStatusResponse { return v.DownloadStatus }).(SyncSessionStatusResponsePtrOutput)
}

func (o ServerEndpointHealthResponseOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.LastUpdatedTimestamp }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponseOutput) OfflineDataTransferStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.OfflineDataTransferStatus }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponseOutput) UploadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.UploadHealth }).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponseOutput) UploadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *SyncSessionStatusResponse { return v.UploadStatus }).(SyncSessionStatusResponsePtrOutput)
}

type ServerEndpointHealthResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerEndpointHealthResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointHealthResponse)(nil)).Elem()
}

func (o ServerEndpointHealthResponsePtrOutput) ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput {
	return o
}

func (o ServerEndpointHealthResponsePtrOutput) ToServerEndpointHealthResponsePtrOutputWithContext(ctx context.Context) ServerEndpointHealthResponsePtrOutput {
	return o
}

func (o ServerEndpointHealthResponsePtrOutput) Elem() ServerEndpointHealthResponseOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) ServerEndpointHealthResponse {
		if v != nil {
			return *v
		}
		var ret ServerEndpointHealthResponse
		return ret
	}).(ServerEndpointHealthResponseOutput)
}

func (o ServerEndpointHealthResponsePtrOutput) CombinedHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.CombinedHealth
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponsePtrOutput) CurrentProgress() SyncProgressStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *SyncProgressStatusResponse {
		if v == nil {
			return nil
		}
		return v.CurrentProgress
	}).(SyncProgressStatusResponsePtrOutput)
}

func (o ServerEndpointHealthResponsePtrOutput) DownloadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.DownloadHealth
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponsePtrOutput) DownloadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *SyncSessionStatusResponse {
		if v == nil {
			return nil
		}
		return v.DownloadStatus
	}).(SyncSessionStatusResponsePtrOutput)
}

func (o ServerEndpointHealthResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponsePtrOutput) OfflineDataTransferStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.OfflineDataTransferStatus
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponsePtrOutput) UploadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.UploadHealth
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointHealthResponsePtrOutput) UploadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *SyncSessionStatusResponse {
		if v == nil {
			return nil
		}
		return v.UploadStatus
	}).(SyncSessionStatusResponsePtrOutput)
}

type SyncProgressStatusResponse struct {
	AppliedBytes      *int    `pulumi:"appliedBytes"`
	AppliedItemCount  *int    `pulumi:"appliedItemCount"`
	PerItemErrorCount *int    `pulumi:"perItemErrorCount"`
	ProgressTimestamp *string `pulumi:"progressTimestamp"`
	SyncDirection     *string `pulumi:"syncDirection"`
	TotalBytes        *int    `pulumi:"totalBytes"`
	TotalItemCount    *int    `pulumi:"totalItemCount"`
}





type SyncProgressStatusResponseInput interface {
	pulumi.Input

	ToSyncProgressStatusResponseOutput() SyncProgressStatusResponseOutput
	ToSyncProgressStatusResponseOutputWithContext(context.Context) SyncProgressStatusResponseOutput
}

type SyncProgressStatusResponseArgs struct {
	AppliedBytes      pulumi.IntPtrInput    `pulumi:"appliedBytes"`
	AppliedItemCount  pulumi.IntPtrInput    `pulumi:"appliedItemCount"`
	PerItemErrorCount pulumi.IntPtrInput    `pulumi:"perItemErrorCount"`
	ProgressTimestamp pulumi.StringPtrInput `pulumi:"progressTimestamp"`
	SyncDirection     pulumi.StringPtrInput `pulumi:"syncDirection"`
	TotalBytes        pulumi.IntPtrInput    `pulumi:"totalBytes"`
	TotalItemCount    pulumi.IntPtrInput    `pulumi:"totalItemCount"`
}

func (SyncProgressStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncProgressStatusResponse)(nil)).Elem()
}

func (i SyncProgressStatusResponseArgs) ToSyncProgressStatusResponseOutput() SyncProgressStatusResponseOutput {
	return i.ToSyncProgressStatusResponseOutputWithContext(context.Background())
}

func (i SyncProgressStatusResponseArgs) ToSyncProgressStatusResponseOutputWithContext(ctx context.Context) SyncProgressStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncProgressStatusResponseOutput)
}

func (i SyncProgressStatusResponseArgs) ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput {
	return i.ToSyncProgressStatusResponsePtrOutputWithContext(context.Background())
}

func (i SyncProgressStatusResponseArgs) ToSyncProgressStatusResponsePtrOutputWithContext(ctx context.Context) SyncProgressStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncProgressStatusResponseOutput).ToSyncProgressStatusResponsePtrOutputWithContext(ctx)
}









type SyncProgressStatusResponsePtrInput interface {
	pulumi.Input

	ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput
	ToSyncProgressStatusResponsePtrOutputWithContext(context.Context) SyncProgressStatusResponsePtrOutput
}

type syncProgressStatusResponsePtrType SyncProgressStatusResponseArgs

func SyncProgressStatusResponsePtr(v *SyncProgressStatusResponseArgs) SyncProgressStatusResponsePtrInput {
	return (*syncProgressStatusResponsePtrType)(v)
}

func (*syncProgressStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncProgressStatusResponse)(nil)).Elem()
}

func (i *syncProgressStatusResponsePtrType) ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput {
	return i.ToSyncProgressStatusResponsePtrOutputWithContext(context.Background())
}

func (i *syncProgressStatusResponsePtrType) ToSyncProgressStatusResponsePtrOutputWithContext(ctx context.Context) SyncProgressStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncProgressStatusResponsePtrOutput)
}

type SyncProgressStatusResponseOutput struct{ *pulumi.OutputState }

func (SyncProgressStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncProgressStatusResponse)(nil)).Elem()
}

func (o SyncProgressStatusResponseOutput) ToSyncProgressStatusResponseOutput() SyncProgressStatusResponseOutput {
	return o
}

func (o SyncProgressStatusResponseOutput) ToSyncProgressStatusResponseOutputWithContext(ctx context.Context) SyncProgressStatusResponseOutput {
	return o
}

func (o SyncProgressStatusResponseOutput) ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput {
	return o.ToSyncProgressStatusResponsePtrOutputWithContext(context.Background())
}

func (o SyncProgressStatusResponseOutput) ToSyncProgressStatusResponsePtrOutputWithContext(ctx context.Context) SyncProgressStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncProgressStatusResponse) *SyncProgressStatusResponse {
		return &v
	}).(SyncProgressStatusResponsePtrOutput)
}

func (o SyncProgressStatusResponseOutput) AppliedBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.AppliedBytes }).(pulumi.IntPtrOutput)
}

func (o SyncProgressStatusResponseOutput) AppliedItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.AppliedItemCount }).(pulumi.IntPtrOutput)
}

func (o SyncProgressStatusResponseOutput) PerItemErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.PerItemErrorCount }).(pulumi.IntPtrOutput)
}

func (o SyncProgressStatusResponseOutput) ProgressTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *string { return v.ProgressTimestamp }).(pulumi.StringPtrOutput)
}

func (o SyncProgressStatusResponseOutput) SyncDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *string { return v.SyncDirection }).(pulumi.StringPtrOutput)
}

func (o SyncProgressStatusResponseOutput) TotalBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.TotalBytes }).(pulumi.IntPtrOutput)
}

func (o SyncProgressStatusResponseOutput) TotalItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.TotalItemCount }).(pulumi.IntPtrOutput)
}

type SyncProgressStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncProgressStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncProgressStatusResponse)(nil)).Elem()
}

func (o SyncProgressStatusResponsePtrOutput) ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput {
	return o
}

func (o SyncProgressStatusResponsePtrOutput) ToSyncProgressStatusResponsePtrOutputWithContext(ctx context.Context) SyncProgressStatusResponsePtrOutput {
	return o
}

func (o SyncProgressStatusResponsePtrOutput) Elem() SyncProgressStatusResponseOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) SyncProgressStatusResponse {
		if v != nil {
			return *v
		}
		var ret SyncProgressStatusResponse
		return ret
	}).(SyncProgressStatusResponseOutput)
}

func (o SyncProgressStatusResponsePtrOutput) AppliedBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.AppliedBytes
	}).(pulumi.IntPtrOutput)
}

func (o SyncProgressStatusResponsePtrOutput) AppliedItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.AppliedItemCount
	}).(pulumi.IntPtrOutput)
}

func (o SyncProgressStatusResponsePtrOutput) PerItemErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.PerItemErrorCount
	}).(pulumi.IntPtrOutput)
}

func (o SyncProgressStatusResponsePtrOutput) ProgressTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProgressTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o SyncProgressStatusResponsePtrOutput) SyncDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.SyncDirection
	}).(pulumi.StringPtrOutput)
}

func (o SyncProgressStatusResponsePtrOutput) TotalBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.TotalBytes
	}).(pulumi.IntPtrOutput)
}

func (o SyncProgressStatusResponsePtrOutput) TotalItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.TotalItemCount
	}).(pulumi.IntPtrOutput)
}

type SyncSessionStatusResponse struct {
	LastSyncPerItemErrorCount *int    `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult            *int    `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp  *string `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp         *string `pulumi:"lastSyncTimestamp"`
}





type SyncSessionStatusResponseInput interface {
	pulumi.Input

	ToSyncSessionStatusResponseOutput() SyncSessionStatusResponseOutput
	ToSyncSessionStatusResponseOutputWithContext(context.Context) SyncSessionStatusResponseOutput
}

type SyncSessionStatusResponseArgs struct {
	LastSyncPerItemErrorCount pulumi.IntPtrInput    `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult            pulumi.IntPtrInput    `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp  pulumi.StringPtrInput `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp         pulumi.StringPtrInput `pulumi:"lastSyncTimestamp"`
}

func (SyncSessionStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncSessionStatusResponse)(nil)).Elem()
}

func (i SyncSessionStatusResponseArgs) ToSyncSessionStatusResponseOutput() SyncSessionStatusResponseOutput {
	return i.ToSyncSessionStatusResponseOutputWithContext(context.Background())
}

func (i SyncSessionStatusResponseArgs) ToSyncSessionStatusResponseOutputWithContext(ctx context.Context) SyncSessionStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncSessionStatusResponseOutput)
}

func (i SyncSessionStatusResponseArgs) ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput {
	return i.ToSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (i SyncSessionStatusResponseArgs) ToSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) SyncSessionStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncSessionStatusResponseOutput).ToSyncSessionStatusResponsePtrOutputWithContext(ctx)
}









type SyncSessionStatusResponsePtrInput interface {
	pulumi.Input

	ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput
	ToSyncSessionStatusResponsePtrOutputWithContext(context.Context) SyncSessionStatusResponsePtrOutput
}

type syncSessionStatusResponsePtrType SyncSessionStatusResponseArgs

func SyncSessionStatusResponsePtr(v *SyncSessionStatusResponseArgs) SyncSessionStatusResponsePtrInput {
	return (*syncSessionStatusResponsePtrType)(v)
}

func (*syncSessionStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncSessionStatusResponse)(nil)).Elem()
}

func (i *syncSessionStatusResponsePtrType) ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput {
	return i.ToSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (i *syncSessionStatusResponsePtrType) ToSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) SyncSessionStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncSessionStatusResponsePtrOutput)
}

type SyncSessionStatusResponseOutput struct{ *pulumi.OutputState }

func (SyncSessionStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncSessionStatusResponse)(nil)).Elem()
}

func (o SyncSessionStatusResponseOutput) ToSyncSessionStatusResponseOutput() SyncSessionStatusResponseOutput {
	return o
}

func (o SyncSessionStatusResponseOutput) ToSyncSessionStatusResponseOutputWithContext(ctx context.Context) SyncSessionStatusResponseOutput {
	return o
}

func (o SyncSessionStatusResponseOutput) ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput {
	return o.ToSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (o SyncSessionStatusResponseOutput) ToSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) SyncSessionStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncSessionStatusResponse) *SyncSessionStatusResponse {
		return &v
	}).(SyncSessionStatusResponsePtrOutput)
}

func (o SyncSessionStatusResponseOutput) LastSyncPerItemErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *int { return v.LastSyncPerItemErrorCount }).(pulumi.IntPtrOutput)
}

func (o SyncSessionStatusResponseOutput) LastSyncResult() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *int { return v.LastSyncResult }).(pulumi.IntPtrOutput)
}

func (o SyncSessionStatusResponseOutput) LastSyncSuccessTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *string { return v.LastSyncSuccessTimestamp }).(pulumi.StringPtrOutput)
}

func (o SyncSessionStatusResponseOutput) LastSyncTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *string { return v.LastSyncTimestamp }).(pulumi.StringPtrOutput)
}

type SyncSessionStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncSessionStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncSessionStatusResponse)(nil)).Elem()
}

func (o SyncSessionStatusResponsePtrOutput) ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput {
	return o
}

func (o SyncSessionStatusResponsePtrOutput) ToSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) SyncSessionStatusResponsePtrOutput {
	return o
}

func (o SyncSessionStatusResponsePtrOutput) Elem() SyncSessionStatusResponseOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) SyncSessionStatusResponse {
		if v != nil {
			return *v
		}
		var ret SyncSessionStatusResponse
		return ret
	}).(SyncSessionStatusResponseOutput)
}

func (o SyncSessionStatusResponsePtrOutput) LastSyncPerItemErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.LastSyncPerItemErrorCount
	}).(pulumi.IntPtrOutput)
}

func (o SyncSessionStatusResponsePtrOutput) LastSyncResult() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.LastSyncResult
	}).(pulumi.IntPtrOutput)
}

func (o SyncSessionStatusResponsePtrOutput) LastSyncSuccessTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastSyncSuccessTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o SyncSessionStatusResponsePtrOutput) LastSyncTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastSyncTimestamp
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerEndpointHealthResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointHealthResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncProgressStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncProgressStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponsePtrOutput{})
}
