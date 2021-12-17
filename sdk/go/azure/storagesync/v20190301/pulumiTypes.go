


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

type SyncActivityStatusResponse struct {
	AppliedBytes      float64 `pulumi:"appliedBytes"`
	AppliedItemCount  float64 `pulumi:"appliedItemCount"`
	PerItemErrorCount float64 `pulumi:"perItemErrorCount"`
	Timestamp         string  `pulumi:"timestamp"`
	TotalBytes        float64 `pulumi:"totalBytes"`
	TotalItemCount    float64 `pulumi:"totalItemCount"`
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

type SyncSessionStatusResponse struct {
	FilesNotSyncingErrors          []FilesNotSyncingErrorResponse `pulumi:"filesNotSyncingErrors"`
	LastSyncPerItemErrorCount      float64                        `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult                 int                            `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp       string                         `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp              string                         `pulumi:"lastSyncTimestamp"`
	PersistentFilesNotSyncingCount float64                        `pulumi:"persistentFilesNotSyncingCount"`
	TransientFilesNotSyncingCount  float64                        `pulumi:"transientFilesNotSyncingCount"`
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

func init() {
	pulumi.RegisterOutputType(FilesNotSyncingErrorResponseOutput{})
	pulumi.RegisterOutputType(FilesNotSyncingErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncActivityStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponseOutput{})
}
