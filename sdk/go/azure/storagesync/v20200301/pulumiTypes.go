


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudTieringCachePerformanceResponse struct {
	CacheHitBytes        float64 `pulumi:"cacheHitBytes"`
	CacheHitBytesPercent int     `pulumi:"cacheHitBytesPercent"`
	CacheMissBytes       float64 `pulumi:"cacheMissBytes"`
	LastUpdatedTimestamp string  `pulumi:"lastUpdatedTimestamp"`
}

type CloudTieringCachePerformanceResponseOutput struct{ *pulumi.OutputState }

func (CloudTieringCachePerformanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringCachePerformanceResponse)(nil)).Elem()
}

func (o CloudTieringCachePerformanceResponseOutput) ToCloudTieringCachePerformanceResponseOutput() CloudTieringCachePerformanceResponseOutput {
	return o
}

func (o CloudTieringCachePerformanceResponseOutput) ToCloudTieringCachePerformanceResponseOutputWithContext(ctx context.Context) CloudTieringCachePerformanceResponseOutput {
	return o
}

func (o CloudTieringCachePerformanceResponseOutput) CacheHitBytes() pulumi.Float64Output {
	return o.ApplyT(func(v CloudTieringCachePerformanceResponse) float64 { return v.CacheHitBytes }).(pulumi.Float64Output)
}

func (o CloudTieringCachePerformanceResponseOutput) CacheHitBytesPercent() pulumi.IntOutput {
	return o.ApplyT(func(v CloudTieringCachePerformanceResponse) int { return v.CacheHitBytesPercent }).(pulumi.IntOutput)
}

func (o CloudTieringCachePerformanceResponseOutput) CacheMissBytes() pulumi.Float64Output {
	return o.ApplyT(func(v CloudTieringCachePerformanceResponse) float64 { return v.CacheMissBytes }).(pulumi.Float64Output)
}

func (o CloudTieringCachePerformanceResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CloudTieringCachePerformanceResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

type CloudTieringDatePolicyStatusResponse struct {
	LastUpdatedTimestamp                 string `pulumi:"lastUpdatedTimestamp"`
	TieredFilesMostRecentAccessTimestamp string `pulumi:"tieredFilesMostRecentAccessTimestamp"`
}

type CloudTieringDatePolicyStatusResponseOutput struct{ *pulumi.OutputState }

func (CloudTieringDatePolicyStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringDatePolicyStatusResponse)(nil)).Elem()
}

func (o CloudTieringDatePolicyStatusResponseOutput) ToCloudTieringDatePolicyStatusResponseOutput() CloudTieringDatePolicyStatusResponseOutput {
	return o
}

func (o CloudTieringDatePolicyStatusResponseOutput) ToCloudTieringDatePolicyStatusResponseOutputWithContext(ctx context.Context) CloudTieringDatePolicyStatusResponseOutput {
	return o
}

func (o CloudTieringDatePolicyStatusResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CloudTieringDatePolicyStatusResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

func (o CloudTieringDatePolicyStatusResponseOutput) TieredFilesMostRecentAccessTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CloudTieringDatePolicyStatusResponse) string { return v.TieredFilesMostRecentAccessTimestamp }).(pulumi.StringOutput)
}

type CloudTieringFilesNotTieringResponse struct {
	Errors               []FilesNotTieringErrorResponse `pulumi:"errors"`
	LastUpdatedTimestamp string                         `pulumi:"lastUpdatedTimestamp"`
	TotalFileCount       float64                        `pulumi:"totalFileCount"`
}

type CloudTieringFilesNotTieringResponseOutput struct{ *pulumi.OutputState }

func (CloudTieringFilesNotTieringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringFilesNotTieringResponse)(nil)).Elem()
}

func (o CloudTieringFilesNotTieringResponseOutput) ToCloudTieringFilesNotTieringResponseOutput() CloudTieringFilesNotTieringResponseOutput {
	return o
}

func (o CloudTieringFilesNotTieringResponseOutput) ToCloudTieringFilesNotTieringResponseOutputWithContext(ctx context.Context) CloudTieringFilesNotTieringResponseOutput {
	return o
}

func (o CloudTieringFilesNotTieringResponseOutput) Errors() FilesNotTieringErrorResponseArrayOutput {
	return o.ApplyT(func(v CloudTieringFilesNotTieringResponse) []FilesNotTieringErrorResponse { return v.Errors }).(FilesNotTieringErrorResponseArrayOutput)
}

func (o CloudTieringFilesNotTieringResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CloudTieringFilesNotTieringResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

func (o CloudTieringFilesNotTieringResponseOutput) TotalFileCount() pulumi.Float64Output {
	return o.ApplyT(func(v CloudTieringFilesNotTieringResponse) float64 { return v.TotalFileCount }).(pulumi.Float64Output)
}

type CloudTieringSpaceSavingsResponse struct {
	CachedSizeBytes      float64 `pulumi:"cachedSizeBytes"`
	LastUpdatedTimestamp string  `pulumi:"lastUpdatedTimestamp"`
	SpaceSavingsBytes    float64 `pulumi:"spaceSavingsBytes"`
	SpaceSavingsPercent  int     `pulumi:"spaceSavingsPercent"`
	TotalSizeCloudBytes  float64 `pulumi:"totalSizeCloudBytes"`
	VolumeSizeBytes      float64 `pulumi:"volumeSizeBytes"`
}

type CloudTieringSpaceSavingsResponseOutput struct{ *pulumi.OutputState }

func (CloudTieringSpaceSavingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringSpaceSavingsResponse)(nil)).Elem()
}

func (o CloudTieringSpaceSavingsResponseOutput) ToCloudTieringSpaceSavingsResponseOutput() CloudTieringSpaceSavingsResponseOutput {
	return o
}

func (o CloudTieringSpaceSavingsResponseOutput) ToCloudTieringSpaceSavingsResponseOutputWithContext(ctx context.Context) CloudTieringSpaceSavingsResponseOutput {
	return o
}

func (o CloudTieringSpaceSavingsResponseOutput) CachedSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v CloudTieringSpaceSavingsResponse) float64 { return v.CachedSizeBytes }).(pulumi.Float64Output)
}

func (o CloudTieringSpaceSavingsResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CloudTieringSpaceSavingsResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

func (o CloudTieringSpaceSavingsResponseOutput) SpaceSavingsBytes() pulumi.Float64Output {
	return o.ApplyT(func(v CloudTieringSpaceSavingsResponse) float64 { return v.SpaceSavingsBytes }).(pulumi.Float64Output)
}

func (o CloudTieringSpaceSavingsResponseOutput) SpaceSavingsPercent() pulumi.IntOutput {
	return o.ApplyT(func(v CloudTieringSpaceSavingsResponse) int { return v.SpaceSavingsPercent }).(pulumi.IntOutput)
}

func (o CloudTieringSpaceSavingsResponseOutput) TotalSizeCloudBytes() pulumi.Float64Output {
	return o.ApplyT(func(v CloudTieringSpaceSavingsResponse) float64 { return v.TotalSizeCloudBytes }).(pulumi.Float64Output)
}

func (o CloudTieringSpaceSavingsResponseOutput) VolumeSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v CloudTieringSpaceSavingsResponse) float64 { return v.VolumeSizeBytes }).(pulumi.Float64Output)
}

type CloudTieringVolumeFreeSpacePolicyStatusResponse struct {
	CurrentVolumeFreeSpacePercent  int    `pulumi:"currentVolumeFreeSpacePercent"`
	EffectiveVolumeFreeSpacePolicy int    `pulumi:"effectiveVolumeFreeSpacePolicy"`
	LastUpdatedTimestamp           string `pulumi:"lastUpdatedTimestamp"`
}

type CloudTieringVolumeFreeSpacePolicyStatusResponseOutput struct{ *pulumi.OutputState }

func (CloudTieringVolumeFreeSpacePolicyStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringVolumeFreeSpacePolicyStatusResponse)(nil)).Elem()
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponseOutput) ToCloudTieringVolumeFreeSpacePolicyStatusResponseOutput() CloudTieringVolumeFreeSpacePolicyStatusResponseOutput {
	return o
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponseOutput) ToCloudTieringVolumeFreeSpacePolicyStatusResponseOutputWithContext(ctx context.Context) CloudTieringVolumeFreeSpacePolicyStatusResponseOutput {
	return o
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponseOutput) CurrentVolumeFreeSpacePercent() pulumi.IntOutput {
	return o.ApplyT(func(v CloudTieringVolumeFreeSpacePolicyStatusResponse) int { return v.CurrentVolumeFreeSpacePercent }).(pulumi.IntOutput)
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponseOutput) EffectiveVolumeFreeSpacePolicy() pulumi.IntOutput {
	return o.ApplyT(func(v CloudTieringVolumeFreeSpacePolicyStatusResponse) int { return v.EffectiveVolumeFreeSpacePolicy }).(pulumi.IntOutput)
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CloudTieringVolumeFreeSpacePolicyStatusResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

type FilesNotTieringErrorResponse struct {
	ErrorCode int     `pulumi:"errorCode"`
	FileCount float64 `pulumi:"fileCount"`
}

type FilesNotTieringErrorResponseOutput struct{ *pulumi.OutputState }

func (FilesNotTieringErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilesNotTieringErrorResponse)(nil)).Elem()
}

func (o FilesNotTieringErrorResponseOutput) ToFilesNotTieringErrorResponseOutput() FilesNotTieringErrorResponseOutput {
	return o
}

func (o FilesNotTieringErrorResponseOutput) ToFilesNotTieringErrorResponseOutputWithContext(ctx context.Context) FilesNotTieringErrorResponseOutput {
	return o
}

func (o FilesNotTieringErrorResponseOutput) ErrorCode() pulumi.IntOutput {
	return o.ApplyT(func(v FilesNotTieringErrorResponse) int { return v.ErrorCode }).(pulumi.IntOutput)
}

func (o FilesNotTieringErrorResponseOutput) FileCount() pulumi.Float64Output {
	return o.ApplyT(func(v FilesNotTieringErrorResponse) float64 { return v.FileCount }).(pulumi.Float64Output)
}

type FilesNotTieringErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (FilesNotTieringErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilesNotTieringErrorResponse)(nil)).Elem()
}

func (o FilesNotTieringErrorResponseArrayOutput) ToFilesNotTieringErrorResponseArrayOutput() FilesNotTieringErrorResponseArrayOutput {
	return o
}

func (o FilesNotTieringErrorResponseArrayOutput) ToFilesNotTieringErrorResponseArrayOutputWithContext(ctx context.Context) FilesNotTieringErrorResponseArrayOutput {
	return o
}

func (o FilesNotTieringErrorResponseArrayOutput) Index(i pulumi.IntInput) FilesNotTieringErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilesNotTieringErrorResponse {
		return vs[0].([]FilesNotTieringErrorResponse)[vs[1].(int)]
	}).(FilesNotTieringErrorResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ServerEndpointCloudTieringStatusResponse struct {
	CachePerformance            CloudTieringCachePerformanceResponse            `pulumi:"cachePerformance"`
	DatePolicyStatus            CloudTieringDatePolicyStatusResponse            `pulumi:"datePolicyStatus"`
	FilesNotTiering             CloudTieringFilesNotTieringResponse             `pulumi:"filesNotTiering"`
	Health                      string                                          `pulumi:"health"`
	HealthLastUpdatedTimestamp  string                                          `pulumi:"healthLastUpdatedTimestamp"`
	LastCloudTieringResult      int                                             `pulumi:"lastCloudTieringResult"`
	LastSuccessTimestamp        string                                          `pulumi:"lastSuccessTimestamp"`
	LastUpdatedTimestamp        string                                          `pulumi:"lastUpdatedTimestamp"`
	SpaceSavings                CloudTieringSpaceSavingsResponse                `pulumi:"spaceSavings"`
	VolumeFreeSpacePolicyStatus CloudTieringVolumeFreeSpacePolicyStatusResponse `pulumi:"volumeFreeSpacePolicyStatus"`
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

func (o ServerEndpointCloudTieringStatusResponseOutput) CachePerformance() CloudTieringCachePerformanceResponseOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) CloudTieringCachePerformanceResponse {
		return v.CachePerformance
	}).(CloudTieringCachePerformanceResponseOutput)
}

func (o ServerEndpointCloudTieringStatusResponseOutput) DatePolicyStatus() CloudTieringDatePolicyStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) CloudTieringDatePolicyStatusResponse {
		return v.DatePolicyStatus
	}).(CloudTieringDatePolicyStatusResponseOutput)
}

func (o ServerEndpointCloudTieringStatusResponseOutput) FilesNotTiering() CloudTieringFilesNotTieringResponseOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) CloudTieringFilesNotTieringResponse {
		return v.FilesNotTiering
	}).(CloudTieringFilesNotTieringResponseOutput)
}

func (o ServerEndpointCloudTieringStatusResponseOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) string { return v.Health }).(pulumi.StringOutput)
}

func (o ServerEndpointCloudTieringStatusResponseOutput) HealthLastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) string { return v.HealthLastUpdatedTimestamp }).(pulumi.StringOutput)
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

func (o ServerEndpointCloudTieringStatusResponseOutput) SpaceSavings() CloudTieringSpaceSavingsResponseOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) CloudTieringSpaceSavingsResponse {
		return v.SpaceSavings
	}).(CloudTieringSpaceSavingsResponseOutput)
}

func (o ServerEndpointCloudTieringStatusResponseOutput) VolumeFreeSpacePolicyStatus() CloudTieringVolumeFreeSpacePolicyStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointCloudTieringStatusResponse) CloudTieringVolumeFreeSpacePolicyStatusResponse {
		return v.VolumeFreeSpacePolicyStatus
	}).(CloudTieringVolumeFreeSpacePolicyStatusResponseOutput)
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
	SyncMode          string  `pulumi:"syncMode"`
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

func (o ServerEndpointSyncActivityStatusResponseOutput) SyncMode() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncActivityStatusResponse) string { return v.SyncMode }).(pulumi.StringOutput)
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
	LastSyncMode                   string                                       `pulumi:"lastSyncMode"`
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

func (o ServerEndpointSyncSessionStatusResponseOutput) LastSyncMode() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncSessionStatusResponse) string { return v.LastSyncMode }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(CloudTieringCachePerformanceResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringDatePolicyStatusResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringFilesNotTieringResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringSpaceSavingsResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringVolumeFreeSpacePolicyStatusResponseOutput{})
	pulumi.RegisterOutputType(FilesNotTieringErrorResponseOutput{})
	pulumi.RegisterOutputType(FilesNotTieringErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
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
