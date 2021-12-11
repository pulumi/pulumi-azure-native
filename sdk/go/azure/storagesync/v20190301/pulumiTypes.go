


package v20190301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FilesNotSyncingErrorResponse struct {
	ErrorCode       int     `pulumi:"errorCode"`
	PersistentCount float64 `pulumi:"persistentCount"`
	TransientCount  float64 `pulumi:"transientCount"`
}





type FilesNotSyncingErrorResponseInput interface {
	pulumi.Input

	ToFilesNotSyncingErrorResponseOutput() FilesNotSyncingErrorResponseOutput
	ToFilesNotSyncingErrorResponseOutputWithContext(context.Context) FilesNotSyncingErrorResponseOutput
}

type FilesNotSyncingErrorResponseArgs struct {
	ErrorCode       pulumi.IntInput     `pulumi:"errorCode"`
	PersistentCount pulumi.Float64Input `pulumi:"persistentCount"`
	TransientCount  pulumi.Float64Input `pulumi:"transientCount"`
}

func (FilesNotSyncingErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilesNotSyncingErrorResponse)(nil)).Elem()
}

func (i FilesNotSyncingErrorResponseArgs) ToFilesNotSyncingErrorResponseOutput() FilesNotSyncingErrorResponseOutput {
	return i.ToFilesNotSyncingErrorResponseOutputWithContext(context.Background())
}

func (i FilesNotSyncingErrorResponseArgs) ToFilesNotSyncingErrorResponseOutputWithContext(ctx context.Context) FilesNotSyncingErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilesNotSyncingErrorResponseOutput)
}





type FilesNotSyncingErrorResponseArrayInput interface {
	pulumi.Input

	ToFilesNotSyncingErrorResponseArrayOutput() FilesNotSyncingErrorResponseArrayOutput
	ToFilesNotSyncingErrorResponseArrayOutputWithContext(context.Context) FilesNotSyncingErrorResponseArrayOutput
}

type FilesNotSyncingErrorResponseArray []FilesNotSyncingErrorResponseInput

func (FilesNotSyncingErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilesNotSyncingErrorResponse)(nil)).Elem()
}

func (i FilesNotSyncingErrorResponseArray) ToFilesNotSyncingErrorResponseArrayOutput() FilesNotSyncingErrorResponseArrayOutput {
	return i.ToFilesNotSyncingErrorResponseArrayOutputWithContext(context.Background())
}

func (i FilesNotSyncingErrorResponseArray) ToFilesNotSyncingErrorResponseArrayOutputWithContext(ctx context.Context) FilesNotSyncingErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilesNotSyncingErrorResponseArrayOutput)
}

type FilesNotSyncingErrorResponseOutput struct{ *pulumi.OutputState }

func (FilesNotSyncingErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilesNotSyncingErrorResponse)(nil)).Elem()
}

func (o FilesNotSyncingErrorResponseOutput) ToFilesNotSyncingErrorResponseOutput() FilesNotSyncingErrorResponseOutput {
	return o
}

func (o FilesNotSyncingErrorResponseOutput) ToFilesNotSyncingErrorResponseOutputWithContext(ctx context.Context) FilesNotSyncingErrorResponseOutput {
	return o
}

func (o FilesNotSyncingErrorResponseOutput) ErrorCode() pulumi.IntOutput {
	return o.ApplyT(func(v FilesNotSyncingErrorResponse) int { return v.ErrorCode }).(pulumi.IntOutput)
}

func (o FilesNotSyncingErrorResponseOutput) PersistentCount() pulumi.Float64Output {
	return o.ApplyT(func(v FilesNotSyncingErrorResponse) float64 { return v.PersistentCount }).(pulumi.Float64Output)
}

func (o FilesNotSyncingErrorResponseOutput) TransientCount() pulumi.Float64Output {
	return o.ApplyT(func(v FilesNotSyncingErrorResponse) float64 { return v.TransientCount }).(pulumi.Float64Output)
}

type FilesNotSyncingErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (FilesNotSyncingErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilesNotSyncingErrorResponse)(nil)).Elem()
}

func (o FilesNotSyncingErrorResponseArrayOutput) ToFilesNotSyncingErrorResponseArrayOutput() FilesNotSyncingErrorResponseArrayOutput {
	return o
}

func (o FilesNotSyncingErrorResponseArrayOutput) ToFilesNotSyncingErrorResponseArrayOutputWithContext(ctx context.Context) FilesNotSyncingErrorResponseArrayOutput {
	return o
}

func (o FilesNotSyncingErrorResponseArrayOutput) Index(i pulumi.IntInput) FilesNotSyncingErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilesNotSyncingErrorResponse {
		return vs[0].([]FilesNotSyncingErrorResponse)[vs[1].(int)]
	}).(FilesNotSyncingErrorResponseOutput)
}

type ServerEndpointSyncStatusResponse struct {
	CombinedHealth                      string                     `pulumi:"combinedHealth"`
	DownloadActivity                    SyncActivityStatusResponse `pulumi:"downloadActivity"`
	DownloadHealth                      string                     `pulumi:"downloadHealth"`
	DownloadStatus                      SyncSessionStatusResponse  `pulumi:"downloadStatus"`
	LastUpdatedTimestamp                string                     `pulumi:"lastUpdatedTimestamp"`
	OfflineDataTransferStatus           string                     `pulumi:"offlineDataTransferStatus"`
	SyncActivity                        string                     `pulumi:"syncActivity"`
	TotalPersistentFilesNotSyncingCount float64                    `pulumi:"totalPersistentFilesNotSyncingCount"`
	UploadActivity                      SyncActivityStatusResponse `pulumi:"uploadActivity"`
	UploadHealth                        string                     `pulumi:"uploadHealth"`
	UploadStatus                        SyncSessionStatusResponse  `pulumi:"uploadStatus"`
}





type ServerEndpointSyncStatusResponseInput interface {
	pulumi.Input

	ToServerEndpointSyncStatusResponseOutput() ServerEndpointSyncStatusResponseOutput
	ToServerEndpointSyncStatusResponseOutputWithContext(context.Context) ServerEndpointSyncStatusResponseOutput
}

type ServerEndpointSyncStatusResponseArgs struct {
	CombinedHealth                      pulumi.StringInput              `pulumi:"combinedHealth"`
	DownloadActivity                    SyncActivityStatusResponseInput `pulumi:"downloadActivity"`
	DownloadHealth                      pulumi.StringInput              `pulumi:"downloadHealth"`
	DownloadStatus                      SyncSessionStatusResponseInput  `pulumi:"downloadStatus"`
	LastUpdatedTimestamp                pulumi.StringInput              `pulumi:"lastUpdatedTimestamp"`
	OfflineDataTransferStatus           pulumi.StringInput              `pulumi:"offlineDataTransferStatus"`
	SyncActivity                        pulumi.StringInput              `pulumi:"syncActivity"`
	TotalPersistentFilesNotSyncingCount pulumi.Float64Input             `pulumi:"totalPersistentFilesNotSyncingCount"`
	UploadActivity                      SyncActivityStatusResponseInput `pulumi:"uploadActivity"`
	UploadHealth                        pulumi.StringInput              `pulumi:"uploadHealth"`
	UploadStatus                        SyncSessionStatusResponseInput  `pulumi:"uploadStatus"`
}

func (ServerEndpointSyncStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointSyncStatusResponse)(nil)).Elem()
}

func (i ServerEndpointSyncStatusResponseArgs) ToServerEndpointSyncStatusResponseOutput() ServerEndpointSyncStatusResponseOutput {
	return i.ToServerEndpointSyncStatusResponseOutputWithContext(context.Background())
}

func (i ServerEndpointSyncStatusResponseArgs) ToServerEndpointSyncStatusResponseOutputWithContext(ctx context.Context) ServerEndpointSyncStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncStatusResponseOutput)
}

func (i ServerEndpointSyncStatusResponseArgs) ToServerEndpointSyncStatusResponsePtrOutput() ServerEndpointSyncStatusResponsePtrOutput {
	return i.ToServerEndpointSyncStatusResponsePtrOutputWithContext(context.Background())
}

func (i ServerEndpointSyncStatusResponseArgs) ToServerEndpointSyncStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncStatusResponseOutput).ToServerEndpointSyncStatusResponsePtrOutputWithContext(ctx)
}









type ServerEndpointSyncStatusResponsePtrInput interface {
	pulumi.Input

	ToServerEndpointSyncStatusResponsePtrOutput() ServerEndpointSyncStatusResponsePtrOutput
	ToServerEndpointSyncStatusResponsePtrOutputWithContext(context.Context) ServerEndpointSyncStatusResponsePtrOutput
}

type serverEndpointSyncStatusResponsePtrType ServerEndpointSyncStatusResponseArgs

func ServerEndpointSyncStatusResponsePtr(v *ServerEndpointSyncStatusResponseArgs) ServerEndpointSyncStatusResponsePtrInput {
	return (*serverEndpointSyncStatusResponsePtrType)(v)
}

func (*serverEndpointSyncStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointSyncStatusResponse)(nil)).Elem()
}

func (i *serverEndpointSyncStatusResponsePtrType) ToServerEndpointSyncStatusResponsePtrOutput() ServerEndpointSyncStatusResponsePtrOutput {
	return i.ToServerEndpointSyncStatusResponsePtrOutputWithContext(context.Background())
}

func (i *serverEndpointSyncStatusResponsePtrType) ToServerEndpointSyncStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncStatusResponsePtrOutput)
}

type ServerEndpointSyncStatusResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointSyncStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointSyncStatusResponse)(nil)).Elem()
}

func (o ServerEndpointSyncStatusResponseOutput) ToServerEndpointSyncStatusResponseOutput() ServerEndpointSyncStatusResponseOutput {
	return o
}

func (o ServerEndpointSyncStatusResponseOutput) ToServerEndpointSyncStatusResponseOutputWithContext(ctx context.Context) ServerEndpointSyncStatusResponseOutput {
	return o
}

func (o ServerEndpointSyncStatusResponseOutput) ToServerEndpointSyncStatusResponsePtrOutput() ServerEndpointSyncStatusResponsePtrOutput {
	return o.ToServerEndpointSyncStatusResponsePtrOutputWithContext(context.Background())
}

func (o ServerEndpointSyncStatusResponseOutput) ToServerEndpointSyncStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerEndpointSyncStatusResponse) *ServerEndpointSyncStatusResponse {
		return &v
	}).(ServerEndpointSyncStatusResponsePtrOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) CombinedHealth() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.CombinedHealth }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) DownloadActivity() SyncActivityStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) SyncActivityStatusResponse { return v.DownloadActivity }).(SyncActivityStatusResponseOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) DownloadHealth() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.DownloadHealth }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) DownloadStatus() SyncSessionStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) SyncSessionStatusResponse { return v.DownloadStatus }).(SyncSessionStatusResponseOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) OfflineDataTransferStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.OfflineDataTransferStatus }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) SyncActivity() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.SyncActivity }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) TotalPersistentFilesNotSyncingCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) float64 { return v.TotalPersistentFilesNotSyncingCount }).(pulumi.Float64Output)
}

func (o ServerEndpointSyncStatusResponseOutput) UploadActivity() SyncActivityStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) SyncActivityStatusResponse { return v.UploadActivity }).(SyncActivityStatusResponseOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) UploadHealth() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.UploadHealth }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) UploadStatus() SyncSessionStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) SyncSessionStatusResponse { return v.UploadStatus }).(SyncSessionStatusResponseOutput)
}

type ServerEndpointSyncStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerEndpointSyncStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointSyncStatusResponse)(nil)).Elem()
}

func (o ServerEndpointSyncStatusResponsePtrOutput) ToServerEndpointSyncStatusResponsePtrOutput() ServerEndpointSyncStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointSyncStatusResponsePtrOutput) ToServerEndpointSyncStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointSyncStatusResponsePtrOutput) Elem() ServerEndpointSyncStatusResponseOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) ServerEndpointSyncStatusResponse {
		if v != nil {
			return *v
		}
		var ret ServerEndpointSyncStatusResponse
		return ret
	}).(ServerEndpointSyncStatusResponseOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) CombinedHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CombinedHealth
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) DownloadActivity() SyncActivityStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *SyncActivityStatusResponse {
		if v == nil {
			return nil
		}
		return &v.DownloadActivity
	}).(SyncActivityStatusResponsePtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) DownloadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DownloadHealth
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) DownloadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *SyncSessionStatusResponse {
		if v == nil {
			return nil
		}
		return &v.DownloadStatus
	}).(SyncSessionStatusResponsePtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) OfflineDataTransferStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OfflineDataTransferStatus
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) SyncActivity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SyncActivity
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) TotalPersistentFilesNotSyncingCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TotalPersistentFilesNotSyncingCount
	}).(pulumi.Float64PtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) UploadActivity() SyncActivityStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *SyncActivityStatusResponse {
		if v == nil {
			return nil
		}
		return &v.UploadActivity
	}).(SyncActivityStatusResponsePtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) UploadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UploadHealth
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) UploadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *SyncSessionStatusResponse {
		if v == nil {
			return nil
		}
		return &v.UploadStatus
	}).(SyncSessionStatusResponsePtrOutput)
}

type SyncActivityStatusResponse struct {
	AppliedBytes      float64 `pulumi:"appliedBytes"`
	AppliedItemCount  float64 `pulumi:"appliedItemCount"`
	PerItemErrorCount float64 `pulumi:"perItemErrorCount"`
	Timestamp         string  `pulumi:"timestamp"`
	TotalBytes        float64 `pulumi:"totalBytes"`
	TotalItemCount    float64 `pulumi:"totalItemCount"`
}





type SyncActivityStatusResponseInput interface {
	pulumi.Input

	ToSyncActivityStatusResponseOutput() SyncActivityStatusResponseOutput
	ToSyncActivityStatusResponseOutputWithContext(context.Context) SyncActivityStatusResponseOutput
}

type SyncActivityStatusResponseArgs struct {
	AppliedBytes      pulumi.Float64Input `pulumi:"appliedBytes"`
	AppliedItemCount  pulumi.Float64Input `pulumi:"appliedItemCount"`
	PerItemErrorCount pulumi.Float64Input `pulumi:"perItemErrorCount"`
	Timestamp         pulumi.StringInput  `pulumi:"timestamp"`
	TotalBytes        pulumi.Float64Input `pulumi:"totalBytes"`
	TotalItemCount    pulumi.Float64Input `pulumi:"totalItemCount"`
}

func (SyncActivityStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncActivityStatusResponse)(nil)).Elem()
}

func (i SyncActivityStatusResponseArgs) ToSyncActivityStatusResponseOutput() SyncActivityStatusResponseOutput {
	return i.ToSyncActivityStatusResponseOutputWithContext(context.Background())
}

func (i SyncActivityStatusResponseArgs) ToSyncActivityStatusResponseOutputWithContext(ctx context.Context) SyncActivityStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncActivityStatusResponseOutput)
}

func (i SyncActivityStatusResponseArgs) ToSyncActivityStatusResponsePtrOutput() SyncActivityStatusResponsePtrOutput {
	return i.ToSyncActivityStatusResponsePtrOutputWithContext(context.Background())
}

func (i SyncActivityStatusResponseArgs) ToSyncActivityStatusResponsePtrOutputWithContext(ctx context.Context) SyncActivityStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncActivityStatusResponseOutput).ToSyncActivityStatusResponsePtrOutputWithContext(ctx)
}









type SyncActivityStatusResponsePtrInput interface {
	pulumi.Input

	ToSyncActivityStatusResponsePtrOutput() SyncActivityStatusResponsePtrOutput
	ToSyncActivityStatusResponsePtrOutputWithContext(context.Context) SyncActivityStatusResponsePtrOutput
}

type syncActivityStatusResponsePtrType SyncActivityStatusResponseArgs

func SyncActivityStatusResponsePtr(v *SyncActivityStatusResponseArgs) SyncActivityStatusResponsePtrInput {
	return (*syncActivityStatusResponsePtrType)(v)
}

func (*syncActivityStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncActivityStatusResponse)(nil)).Elem()
}

func (i *syncActivityStatusResponsePtrType) ToSyncActivityStatusResponsePtrOutput() SyncActivityStatusResponsePtrOutput {
	return i.ToSyncActivityStatusResponsePtrOutputWithContext(context.Background())
}

func (i *syncActivityStatusResponsePtrType) ToSyncActivityStatusResponsePtrOutputWithContext(ctx context.Context) SyncActivityStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncActivityStatusResponsePtrOutput)
}

type SyncActivityStatusResponseOutput struct{ *pulumi.OutputState }

func (SyncActivityStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncActivityStatusResponse)(nil)).Elem()
}

func (o SyncActivityStatusResponseOutput) ToSyncActivityStatusResponseOutput() SyncActivityStatusResponseOutput {
	return o
}

func (o SyncActivityStatusResponseOutput) ToSyncActivityStatusResponseOutputWithContext(ctx context.Context) SyncActivityStatusResponseOutput {
	return o
}

func (o SyncActivityStatusResponseOutput) ToSyncActivityStatusResponsePtrOutput() SyncActivityStatusResponsePtrOutput {
	return o.ToSyncActivityStatusResponsePtrOutputWithContext(context.Background())
}

func (o SyncActivityStatusResponseOutput) ToSyncActivityStatusResponsePtrOutputWithContext(ctx context.Context) SyncActivityStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncActivityStatusResponse) *SyncActivityStatusResponse {
		return &v
	}).(SyncActivityStatusResponsePtrOutput)
}

func (o SyncActivityStatusResponseOutput) AppliedBytes() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.AppliedBytes }).(pulumi.Float64Output)
}

func (o SyncActivityStatusResponseOutput) AppliedItemCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.AppliedItemCount }).(pulumi.Float64Output)
}

func (o SyncActivityStatusResponseOutput) PerItemErrorCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.PerItemErrorCount }).(pulumi.Float64Output)
}

func (o SyncActivityStatusResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v SyncActivityStatusResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o SyncActivityStatusResponseOutput) TotalBytes() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.TotalBytes }).(pulumi.Float64Output)
}

func (o SyncActivityStatusResponseOutput) TotalItemCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.TotalItemCount }).(pulumi.Float64Output)
}

type SyncActivityStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncActivityStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncActivityStatusResponse)(nil)).Elem()
}

func (o SyncActivityStatusResponsePtrOutput) ToSyncActivityStatusResponsePtrOutput() SyncActivityStatusResponsePtrOutput {
	return o
}

func (o SyncActivityStatusResponsePtrOutput) ToSyncActivityStatusResponsePtrOutputWithContext(ctx context.Context) SyncActivityStatusResponsePtrOutput {
	return o
}

func (o SyncActivityStatusResponsePtrOutput) Elem() SyncActivityStatusResponseOutput {
	return o.ApplyT(func(v *SyncActivityStatusResponse) SyncActivityStatusResponse {
		if v != nil {
			return *v
		}
		var ret SyncActivityStatusResponse
		return ret
	}).(SyncActivityStatusResponseOutput)
}

func (o SyncActivityStatusResponsePtrOutput) AppliedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.AppliedBytes
	}).(pulumi.Float64PtrOutput)
}

func (o SyncActivityStatusResponsePtrOutput) AppliedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.AppliedItemCount
	}).(pulumi.Float64PtrOutput)
}

func (o SyncActivityStatusResponsePtrOutput) PerItemErrorCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.PerItemErrorCount
	}).(pulumi.Float64PtrOutput)
}

func (o SyncActivityStatusResponsePtrOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncActivityStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Timestamp
	}).(pulumi.StringPtrOutput)
}

func (o SyncActivityStatusResponsePtrOutput) TotalBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TotalBytes
	}).(pulumi.Float64PtrOutput)
}

func (o SyncActivityStatusResponsePtrOutput) TotalItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TotalItemCount
	}).(pulumi.Float64PtrOutput)
}

type SyncSessionStatusResponse struct {
	FilesNotSyncingErrors          []FilesNotSyncingErrorResponse `pulumi:"filesNotSyncingErrors"`
	LastSyncPerItemErrorCount      float64                        `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult                 int                            `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp       string                         `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp              string                         `pulumi:"lastSyncTimestamp"`
	PersistentFilesNotSyncingCount float64                        `pulumi:"persistentFilesNotSyncingCount"`
	TransientFilesNotSyncingCount  float64                        `pulumi:"transientFilesNotSyncingCount"`
}





type SyncSessionStatusResponseInput interface {
	pulumi.Input

	ToSyncSessionStatusResponseOutput() SyncSessionStatusResponseOutput
	ToSyncSessionStatusResponseOutputWithContext(context.Context) SyncSessionStatusResponseOutput
}

type SyncSessionStatusResponseArgs struct {
	FilesNotSyncingErrors          FilesNotSyncingErrorResponseArrayInput `pulumi:"filesNotSyncingErrors"`
	LastSyncPerItemErrorCount      pulumi.Float64Input                    `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult                 pulumi.IntInput                        `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp       pulumi.StringInput                     `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp              pulumi.StringInput                     `pulumi:"lastSyncTimestamp"`
	PersistentFilesNotSyncingCount pulumi.Float64Input                    `pulumi:"persistentFilesNotSyncingCount"`
	TransientFilesNotSyncingCount  pulumi.Float64Input                    `pulumi:"transientFilesNotSyncingCount"`
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

func (o SyncSessionStatusResponseOutput) FilesNotSyncingErrors() FilesNotSyncingErrorResponseArrayOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) []FilesNotSyncingErrorResponse { return v.FilesNotSyncingErrors }).(FilesNotSyncingErrorResponseArrayOutput)
}

func (o SyncSessionStatusResponseOutput) LastSyncPerItemErrorCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncSessionStatusResponse) float64 { return v.LastSyncPerItemErrorCount }).(pulumi.Float64Output)
}

func (o SyncSessionStatusResponseOutput) LastSyncResult() pulumi.IntOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) int { return v.LastSyncResult }).(pulumi.IntOutput)
}

func (o SyncSessionStatusResponseOutput) LastSyncSuccessTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) string { return v.LastSyncSuccessTimestamp }).(pulumi.StringOutput)
}

func (o SyncSessionStatusResponseOutput) LastSyncTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) string { return v.LastSyncTimestamp }).(pulumi.StringOutput)
}

func (o SyncSessionStatusResponseOutput) PersistentFilesNotSyncingCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncSessionStatusResponse) float64 { return v.PersistentFilesNotSyncingCount }).(pulumi.Float64Output)
}

func (o SyncSessionStatusResponseOutput) TransientFilesNotSyncingCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncSessionStatusResponse) float64 { return v.TransientFilesNotSyncingCount }).(pulumi.Float64Output)
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

func (o SyncSessionStatusResponsePtrOutput) FilesNotSyncingErrors() FilesNotSyncingErrorResponseArrayOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) []FilesNotSyncingErrorResponse {
		if v == nil {
			return nil
		}
		return v.FilesNotSyncingErrors
	}).(FilesNotSyncingErrorResponseArrayOutput)
}

func (o SyncSessionStatusResponsePtrOutput) LastSyncPerItemErrorCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.LastSyncPerItemErrorCount
	}).(pulumi.Float64PtrOutput)
}

func (o SyncSessionStatusResponsePtrOutput) LastSyncResult() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *int {
		if v == nil {
			return nil
		}
		return &v.LastSyncResult
	}).(pulumi.IntPtrOutput)
}

func (o SyncSessionStatusResponsePtrOutput) LastSyncSuccessTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSyncSuccessTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o SyncSessionStatusResponsePtrOutput) LastSyncTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSyncTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o SyncSessionStatusResponsePtrOutput) PersistentFilesNotSyncingCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.PersistentFilesNotSyncingCount
	}).(pulumi.Float64PtrOutput)
}

func (o SyncSessionStatusResponsePtrOutput) TransientFilesNotSyncingCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransientFilesNotSyncingCount
	}).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FilesNotSyncingErrorResponseOutput{})
	pulumi.RegisterOutputType(FilesNotSyncingErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncActivityStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncActivityStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponsePtrOutput{})
}
