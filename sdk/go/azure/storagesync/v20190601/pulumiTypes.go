


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerEndpointCloudTieringStatusResponse struct {
	Health                 string `pulumi:"health"`
	LastCloudTieringResult int    `pulumi:"lastCloudTieringResult"`
	LastSuccessTimestamp   string `pulumi:"lastSuccessTimestamp"`
	LastUpdatedTimestamp   string `pulumi:"lastUpdatedTimestamp"`
}

type ServerEndpointCloudTieringStatusResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointCloudTieringStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointCloudTieringStatusResponse)(nil)).Elem()
}

func (o ServerEndpointCloudTieringStatusResponseOutput) ToServerEndpointCloudTieringStatusResponseOutput() ServerEndpointCloudTieringStatusResponseOutput {
	return o
}

func (o ServerEndpointCloudTieringStatusResponseOutput) ToServerEndpointCloudTieringStatusResponseOutputWithContext(ctx context.Context) ServerEndpointCloudTieringStatusResponseOutput {
	return o
}

func (o ServerEndpointCloudTieringStatusResponseOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) string { return v.Health }).(pulumi.StringOutput)
}

func (o ServerEndpointCloudTieringStatusResponseOutput) LastCloudTieringResult() pulumi.IntOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) int { return v.LastCloudTieringResult }).(pulumi.IntOutput)
}

func (o ServerEndpointCloudTieringStatusResponseOutput) LastSuccessTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) string { return v.LastSuccessTimestamp }).(pulumi.StringOutput)
}

func (o ServerEndpointCloudTieringStatusResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

type ServerEndpointFilesNotSyncingErrorResponse struct {
	ErrorCode       int     `pulumi:"errorCode"`
	PersistentCount float64 `pulumi:"persistentCount"`
	TransientCount  float64 `pulumi:"transientCount"`
}

type ServerEndpointFilesNotSyncingErrorResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointFilesNotSyncingErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointFilesNotSyncingErrorResponse)(nil)).Elem()
}

func (o ServerEndpointFilesNotSyncingErrorResponseOutput) ToServerEndpointFilesNotSyncingErrorResponseOutput() ServerEndpointFilesNotSyncingErrorResponseOutput {
	return o
}

func (o ServerEndpointFilesNotSyncingErrorResponseOutput) ToServerEndpointFilesNotSyncingErrorResponseOutputWithContext(ctx context.Context) ServerEndpointFilesNotSyncingErrorResponseOutput {
	return o
}

func (o ServerEndpointFilesNotSyncingErrorResponseOutput) ErrorCode() pulumi.IntOutput {
	return o.ApplyT(func(v ServerEndpointFilesNotSyncingErrorResponse) int { return v.ErrorCode }).(pulumi.IntOutput)
}

func (o ServerEndpointFilesNotSyncingErrorResponseOutput) PersistentCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointFilesNotSyncingErrorResponse) float64 { return v.PersistentCount }).(pulumi.Float64Output)
}

func (o ServerEndpointFilesNotSyncingErrorResponseOutput) TransientCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointFilesNotSyncingErrorResponse) float64 { return v.TransientCount }).(pulumi.Float64Output)
}

type ServerEndpointFilesNotSyncingErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerEndpointFilesNotSyncingErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerEndpointFilesNotSyncingErrorResponse)(nil)).Elem()
}

func (o ServerEndpointFilesNotSyncingErrorResponseArrayOutput) ToServerEndpointFilesNotSyncingErrorResponseArrayOutput() ServerEndpointFilesNotSyncingErrorResponseArrayOutput {
	return o
}

func (o ServerEndpointFilesNotSyncingErrorResponseArrayOutput) ToServerEndpointFilesNotSyncingErrorResponseArrayOutputWithContext(ctx context.Context) ServerEndpointFilesNotSyncingErrorResponseArrayOutput {
	return o
}

func (o ServerEndpointFilesNotSyncingErrorResponseArrayOutput) Index(i pulumi.IntInput) ServerEndpointFilesNotSyncingErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerEndpointFilesNotSyncingErrorResponse {
		return vs[0].([]ServerEndpointFilesNotSyncingErrorResponse)[vs[1].(int)]
	}).(ServerEndpointFilesNotSyncingErrorResponseOutput)
}

type ServerEndpointRecallErrorResponse struct {
	Count     float64 `pulumi:"count"`
	ErrorCode int     `pulumi:"errorCode"`
}

type ServerEndpointRecallErrorResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointRecallErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointRecallErrorResponse)(nil)).Elem()
}

func (o ServerEndpointRecallErrorResponseOutput) ToServerEndpointRecallErrorResponseOutput() ServerEndpointRecallErrorResponseOutput {
	return o
}

func (o ServerEndpointRecallErrorResponseOutput) ToServerEndpointRecallErrorResponseOutputWithContext(ctx context.Context) ServerEndpointRecallErrorResponseOutput {
	return o
}

func (o ServerEndpointRecallErrorResponseOutput) Count() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointRecallErrorResponse) float64 { return v.Count }).(pulumi.Float64Output)
}

func (o ServerEndpointRecallErrorResponseOutput) ErrorCode() pulumi.IntOutput {
	return o.ApplyT(func(v ServerEndpointRecallErrorResponse) int { return v.ErrorCode }).(pulumi.IntOutput)
}

type ServerEndpointRecallErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerEndpointRecallErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerEndpointRecallErrorResponse)(nil)).Elem()
}

func (o ServerEndpointRecallErrorResponseArrayOutput) ToServerEndpointRecallErrorResponseArrayOutput() ServerEndpointRecallErrorResponseArrayOutput {
	return o
}

func (o ServerEndpointRecallErrorResponseArrayOutput) ToServerEndpointRecallErrorResponseArrayOutputWithContext(ctx context.Context) ServerEndpointRecallErrorResponseArrayOutput {
	return o
}

func (o ServerEndpointRecallErrorResponseArrayOutput) Index(i pulumi.IntInput) ServerEndpointRecallErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerEndpointRecallErrorResponse {
		return vs[0].([]ServerEndpointRecallErrorResponse)[vs[1].(int)]
	}).(ServerEndpointRecallErrorResponseOutput)
}

type ServerEndpointRecallStatusResponse struct {
	LastUpdatedTimestamp   string                              `pulumi:"lastUpdatedTimestamp"`
	RecallErrors           []ServerEndpointRecallErrorResponse `pulumi:"recallErrors"`
	TotalRecallErrorsCount float64                             `pulumi:"totalRecallErrorsCount"`
}

type ServerEndpointRecallStatusResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointRecallStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointRecallStatusResponse)(nil)).Elem()
}

func (o ServerEndpointRecallStatusResponseOutput) ToServerEndpointRecallStatusResponseOutput() ServerEndpointRecallStatusResponseOutput {
	return o
}

func (o ServerEndpointRecallStatusResponseOutput) ToServerEndpointRecallStatusResponseOutputWithContext(ctx context.Context) ServerEndpointRecallStatusResponseOutput {
	return o
}

func (o ServerEndpointRecallStatusResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointRecallStatusResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

func (o ServerEndpointRecallStatusResponseOutput) RecallErrors() ServerEndpointRecallErrorResponseArrayOutput {
	return o.ApplyT(func(v ServerEndpointRecallStatusResponse) []ServerEndpointRecallErrorResponse { return v.RecallErrors }).(ServerEndpointRecallErrorResponseArrayOutput)
}

func (o ServerEndpointRecallStatusResponseOutput) TotalRecallErrorsCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointRecallStatusResponse) float64 { return v.TotalRecallErrorsCount }).(pulumi.Float64Output)
}

type ServerEndpointSyncActivityStatusResponse struct {
	AppliedBytes      float64 `pulumi:"appliedBytes"`
	AppliedItemCount  float64 `pulumi:"appliedItemCount"`
	PerItemErrorCount float64 `pulumi:"perItemErrorCount"`
	Timestamp         string  `pulumi:"timestamp"`
	TotalBytes        float64 `pulumi:"totalBytes"`
	TotalItemCount    float64 `pulumi:"totalItemCount"`
}

type ServerEndpointSyncActivityStatusResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointSyncActivityStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointSyncActivityStatusResponse)(nil)).Elem()
}

func (o ServerEndpointSyncActivityStatusResponseOutput) ToServerEndpointSyncActivityStatusResponseOutput() ServerEndpointSyncActivityStatusResponseOutput {
	return o
}

func (o ServerEndpointSyncActivityStatusResponseOutput) ToServerEndpointSyncActivityStatusResponseOutputWithContext(ctx context.Context) ServerEndpointSyncActivityStatusResponseOutput {
	return o
}

func (o ServerEndpointSyncActivityStatusResponseOutput) AppliedBytes() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncActivityStatusResponse) float64 { return v.AppliedBytes }).(pulumi.Float64Output)
}

func (o ServerEndpointSyncActivityStatusResponseOutput) AppliedItemCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncActivityStatusResponse) float64 { return v.AppliedItemCount }).(pulumi.Float64Output)
}

func (o ServerEndpointSyncActivityStatusResponseOutput) PerItemErrorCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncActivityStatusResponse) float64 { return v.PerItemErrorCount }).(pulumi.Float64Output)
}

func (o ServerEndpointSyncActivityStatusResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncActivityStatusResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncActivityStatusResponseOutput) TotalBytes() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncActivityStatusResponse) float64 { return v.TotalBytes }).(pulumi.Float64Output)
}

func (o ServerEndpointSyncActivityStatusResponseOutput) TotalItemCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncActivityStatusResponse) float64 { return v.TotalItemCount }).(pulumi.Float64Output)
}

type ServerEndpointSyncSessionStatusResponse struct {
	FilesNotSyncingErrors          []ServerEndpointFilesNotSyncingErrorResponse `pulumi:"filesNotSyncingErrors"`
	LastSyncPerItemErrorCount      float64                                      `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult                 int                                          `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp       string                                       `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp              string                                       `pulumi:"lastSyncTimestamp"`
	PersistentFilesNotSyncingCount float64                                      `pulumi:"persistentFilesNotSyncingCount"`
	TransientFilesNotSyncingCount  float64                                      `pulumi:"transientFilesNotSyncingCount"`
}

type ServerEndpointSyncSessionStatusResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointSyncSessionStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointSyncSessionStatusResponse)(nil)).Elem()
}

func (o ServerEndpointSyncSessionStatusResponseOutput) ToServerEndpointSyncSessionStatusResponseOutput() ServerEndpointSyncSessionStatusResponseOutput {
	return o
}

func (o ServerEndpointSyncSessionStatusResponseOutput) ToServerEndpointSyncSessionStatusResponseOutputWithContext(ctx context.Context) ServerEndpointSyncSessionStatusResponseOutput {
	return o
}

func (o ServerEndpointSyncSessionStatusResponseOutput) FilesNotSyncingErrors() ServerEndpointFilesNotSyncingErrorResponseArrayOutput {
	return o.ApplyT(func(v ServerEndpointSyncSessionStatusResponse) []ServerEndpointFilesNotSyncingErrorResponse {
		return v.FilesNotSyncingErrors
	}).(ServerEndpointFilesNotSyncingErrorResponseArrayOutput)
}

func (o ServerEndpointSyncSessionStatusResponseOutput) LastSyncPerItemErrorCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncSessionStatusResponse) float64 { return v.LastSyncPerItemErrorCount }).(pulumi.Float64Output)
}

func (o ServerEndpointSyncSessionStatusResponseOutput) LastSyncResult() pulumi.IntOutput {
	return o.ApplyT(func(v ServerEndpointSyncSessionStatusResponse) int { return v.LastSyncResult }).(pulumi.IntOutput)
}

func (o ServerEndpointSyncSessionStatusResponseOutput) LastSyncSuccessTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncSessionStatusResponse) string { return v.LastSyncSuccessTimestamp }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncSessionStatusResponseOutput) LastSyncTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncSessionStatusResponse) string { return v.LastSyncTimestamp }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncSessionStatusResponseOutput) PersistentFilesNotSyncingCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncSessionStatusResponse) float64 { return v.PersistentFilesNotSyncingCount }).(pulumi.Float64Output)
}

func (o ServerEndpointSyncSessionStatusResponseOutput) TransientFilesNotSyncingCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncSessionStatusResponse) float64 { return v.TransientFilesNotSyncingCount }).(pulumi.Float64Output)
}

type ServerEndpointSyncStatusResponse struct {
	CombinedHealth                      string                                   `pulumi:"combinedHealth"`
	DownloadActivity                    ServerEndpointSyncActivityStatusResponse `pulumi:"downloadActivity"`
	DownloadHealth                      string                                   `pulumi:"downloadHealth"`
	DownloadStatus                      ServerEndpointSyncSessionStatusResponse  `pulumi:"downloadStatus"`
	LastUpdatedTimestamp                string                                   `pulumi:"lastUpdatedTimestamp"`
	OfflineDataTransferStatus           string                                   `pulumi:"offlineDataTransferStatus"`
	SyncActivity                        string                                   `pulumi:"syncActivity"`
	TotalPersistentFilesNotSyncingCount float64                                  `pulumi:"totalPersistentFilesNotSyncingCount"`
	UploadActivity                      ServerEndpointSyncActivityStatusResponse `pulumi:"uploadActivity"`
	UploadHealth                        string                                   `pulumi:"uploadHealth"`
	UploadStatus                        ServerEndpointSyncSessionStatusResponse  `pulumi:"uploadStatus"`
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

func (o ServerEndpointSyncStatusResponseOutput) DownloadActivity() ServerEndpointSyncActivityStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) ServerEndpointSyncActivityStatusResponse {
		return v.DownloadActivity
	}).(ServerEndpointSyncActivityStatusResponseOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) DownloadHealth() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.DownloadHealth }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) DownloadStatus() ServerEndpointSyncSessionStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) ServerEndpointSyncSessionStatusResponse {
		return v.DownloadStatus
	}).(ServerEndpointSyncSessionStatusResponseOutput)
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

func (o ServerEndpointSyncStatusResponseOutput) UploadActivity() ServerEndpointSyncActivityStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) ServerEndpointSyncActivityStatusResponse {
		return v.UploadActivity
	}).(ServerEndpointSyncActivityStatusResponseOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) UploadHealth() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.UploadHealth }).(pulumi.StringOutput)
}

func (o ServerEndpointSyncStatusResponseOutput) UploadStatus() ServerEndpointSyncSessionStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) ServerEndpointSyncSessionStatusResponse {
		return v.UploadStatus
	}).(ServerEndpointSyncSessionStatusResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerEndpointCloudTieringStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointFilesNotSyncingErrorResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointFilesNotSyncingErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallErrorResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncActivityStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncSessionStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponseOutput{})
}
