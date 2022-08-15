


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
