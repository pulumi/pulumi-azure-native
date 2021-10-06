


package v20191001

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





type CloudTieringCachePerformanceResponseInput interface {
	pulumi.Input

	ToCloudTieringCachePerformanceResponseOutput() CloudTieringCachePerformanceResponseOutput
	ToCloudTieringCachePerformanceResponseOutputWithContext(context.Context) CloudTieringCachePerformanceResponseOutput
}

type CloudTieringCachePerformanceResponseArgs struct {
	CacheHitBytes        pulumi.Float64Input `pulumi:"cacheHitBytes"`
	CacheHitBytesPercent pulumi.IntInput     `pulumi:"cacheHitBytesPercent"`
	CacheMissBytes       pulumi.Float64Input `pulumi:"cacheMissBytes"`
	LastUpdatedTimestamp pulumi.StringInput  `pulumi:"lastUpdatedTimestamp"`
}

func (CloudTieringCachePerformanceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringCachePerformanceResponse)(nil)).Elem()
}

func (i CloudTieringCachePerformanceResponseArgs) ToCloudTieringCachePerformanceResponseOutput() CloudTieringCachePerformanceResponseOutput {
	return i.ToCloudTieringCachePerformanceResponseOutputWithContext(context.Background())
}

func (i CloudTieringCachePerformanceResponseArgs) ToCloudTieringCachePerformanceResponseOutputWithContext(ctx context.Context) CloudTieringCachePerformanceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringCachePerformanceResponseOutput)
}

func (i CloudTieringCachePerformanceResponseArgs) ToCloudTieringCachePerformanceResponsePtrOutput() CloudTieringCachePerformanceResponsePtrOutput {
	return i.ToCloudTieringCachePerformanceResponsePtrOutputWithContext(context.Background())
}

func (i CloudTieringCachePerformanceResponseArgs) ToCloudTieringCachePerformanceResponsePtrOutputWithContext(ctx context.Context) CloudTieringCachePerformanceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringCachePerformanceResponseOutput).ToCloudTieringCachePerformanceResponsePtrOutputWithContext(ctx)
}









type CloudTieringCachePerformanceResponsePtrInput interface {
	pulumi.Input

	ToCloudTieringCachePerformanceResponsePtrOutput() CloudTieringCachePerformanceResponsePtrOutput
	ToCloudTieringCachePerformanceResponsePtrOutputWithContext(context.Context) CloudTieringCachePerformanceResponsePtrOutput
}

type cloudTieringCachePerformanceResponsePtrType CloudTieringCachePerformanceResponseArgs

func CloudTieringCachePerformanceResponsePtr(v *CloudTieringCachePerformanceResponseArgs) CloudTieringCachePerformanceResponsePtrInput {
	return (*cloudTieringCachePerformanceResponsePtrType)(v)
}

func (*cloudTieringCachePerformanceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringCachePerformanceResponse)(nil)).Elem()
}

func (i *cloudTieringCachePerformanceResponsePtrType) ToCloudTieringCachePerformanceResponsePtrOutput() CloudTieringCachePerformanceResponsePtrOutput {
	return i.ToCloudTieringCachePerformanceResponsePtrOutputWithContext(context.Background())
}

func (i *cloudTieringCachePerformanceResponsePtrType) ToCloudTieringCachePerformanceResponsePtrOutputWithContext(ctx context.Context) CloudTieringCachePerformanceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringCachePerformanceResponsePtrOutput)
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

func (o CloudTieringCachePerformanceResponseOutput) ToCloudTieringCachePerformanceResponsePtrOutput() CloudTieringCachePerformanceResponsePtrOutput {
	return o.ToCloudTieringCachePerformanceResponsePtrOutputWithContext(context.Background())
}

func (o CloudTieringCachePerformanceResponseOutput) ToCloudTieringCachePerformanceResponsePtrOutputWithContext(ctx context.Context) CloudTieringCachePerformanceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudTieringCachePerformanceResponse) *CloudTieringCachePerformanceResponse {
		return &v
	}).(CloudTieringCachePerformanceResponsePtrOutput)
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

type CloudTieringCachePerformanceResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudTieringCachePerformanceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringCachePerformanceResponse)(nil)).Elem()
}

func (o CloudTieringCachePerformanceResponsePtrOutput) ToCloudTieringCachePerformanceResponsePtrOutput() CloudTieringCachePerformanceResponsePtrOutput {
	return o
}

func (o CloudTieringCachePerformanceResponsePtrOutput) ToCloudTieringCachePerformanceResponsePtrOutputWithContext(ctx context.Context) CloudTieringCachePerformanceResponsePtrOutput {
	return o
}

func (o CloudTieringCachePerformanceResponsePtrOutput) Elem() CloudTieringCachePerformanceResponseOutput {
	return o.ApplyT(func(v *CloudTieringCachePerformanceResponse) CloudTieringCachePerformanceResponse {
		if v != nil {
			return *v
		}
		var ret CloudTieringCachePerformanceResponse
		return ret
	}).(CloudTieringCachePerformanceResponseOutput)
}

func (o CloudTieringCachePerformanceResponsePtrOutput) CacheHitBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudTieringCachePerformanceResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.CacheHitBytes
	}).(pulumi.Float64PtrOutput)
}

func (o CloudTieringCachePerformanceResponsePtrOutput) CacheHitBytesPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudTieringCachePerformanceResponse) *int {
		if v == nil {
			return nil
		}
		return &v.CacheHitBytesPercent
	}).(pulumi.IntPtrOutput)
}

func (o CloudTieringCachePerformanceResponsePtrOutput) CacheMissBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudTieringCachePerformanceResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.CacheMissBytes
	}).(pulumi.Float64PtrOutput)
}

func (o CloudTieringCachePerformanceResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudTieringCachePerformanceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

type CloudTieringDatePolicyStatusResponse struct {
	LastUpdatedTimestamp                 string `pulumi:"lastUpdatedTimestamp"`
	TieredFilesMostRecentAccessTimestamp string `pulumi:"tieredFilesMostRecentAccessTimestamp"`
}





type CloudTieringDatePolicyStatusResponseInput interface {
	pulumi.Input

	ToCloudTieringDatePolicyStatusResponseOutput() CloudTieringDatePolicyStatusResponseOutput
	ToCloudTieringDatePolicyStatusResponseOutputWithContext(context.Context) CloudTieringDatePolicyStatusResponseOutput
}

type CloudTieringDatePolicyStatusResponseArgs struct {
	LastUpdatedTimestamp                 pulumi.StringInput `pulumi:"lastUpdatedTimestamp"`
	TieredFilesMostRecentAccessTimestamp pulumi.StringInput `pulumi:"tieredFilesMostRecentAccessTimestamp"`
}

func (CloudTieringDatePolicyStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringDatePolicyStatusResponse)(nil)).Elem()
}

func (i CloudTieringDatePolicyStatusResponseArgs) ToCloudTieringDatePolicyStatusResponseOutput() CloudTieringDatePolicyStatusResponseOutput {
	return i.ToCloudTieringDatePolicyStatusResponseOutputWithContext(context.Background())
}

func (i CloudTieringDatePolicyStatusResponseArgs) ToCloudTieringDatePolicyStatusResponseOutputWithContext(ctx context.Context) CloudTieringDatePolicyStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringDatePolicyStatusResponseOutput)
}

func (i CloudTieringDatePolicyStatusResponseArgs) ToCloudTieringDatePolicyStatusResponsePtrOutput() CloudTieringDatePolicyStatusResponsePtrOutput {
	return i.ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(context.Background())
}

func (i CloudTieringDatePolicyStatusResponseArgs) ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(ctx context.Context) CloudTieringDatePolicyStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringDatePolicyStatusResponseOutput).ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(ctx)
}









type CloudTieringDatePolicyStatusResponsePtrInput interface {
	pulumi.Input

	ToCloudTieringDatePolicyStatusResponsePtrOutput() CloudTieringDatePolicyStatusResponsePtrOutput
	ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(context.Context) CloudTieringDatePolicyStatusResponsePtrOutput
}

type cloudTieringDatePolicyStatusResponsePtrType CloudTieringDatePolicyStatusResponseArgs

func CloudTieringDatePolicyStatusResponsePtr(v *CloudTieringDatePolicyStatusResponseArgs) CloudTieringDatePolicyStatusResponsePtrInput {
	return (*cloudTieringDatePolicyStatusResponsePtrType)(v)
}

func (*cloudTieringDatePolicyStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringDatePolicyStatusResponse)(nil)).Elem()
}

func (i *cloudTieringDatePolicyStatusResponsePtrType) ToCloudTieringDatePolicyStatusResponsePtrOutput() CloudTieringDatePolicyStatusResponsePtrOutput {
	return i.ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(context.Background())
}

func (i *cloudTieringDatePolicyStatusResponsePtrType) ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(ctx context.Context) CloudTieringDatePolicyStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringDatePolicyStatusResponsePtrOutput)
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

func (o CloudTieringDatePolicyStatusResponseOutput) ToCloudTieringDatePolicyStatusResponsePtrOutput() CloudTieringDatePolicyStatusResponsePtrOutput {
	return o.ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(context.Background())
}

func (o CloudTieringDatePolicyStatusResponseOutput) ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(ctx context.Context) CloudTieringDatePolicyStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudTieringDatePolicyStatusResponse) *CloudTieringDatePolicyStatusResponse {
		return &v
	}).(CloudTieringDatePolicyStatusResponsePtrOutput)
}

func (o CloudTieringDatePolicyStatusResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CloudTieringDatePolicyStatusResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

func (o CloudTieringDatePolicyStatusResponseOutput) TieredFilesMostRecentAccessTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CloudTieringDatePolicyStatusResponse) string { return v.TieredFilesMostRecentAccessTimestamp }).(pulumi.StringOutput)
}

type CloudTieringDatePolicyStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudTieringDatePolicyStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringDatePolicyStatusResponse)(nil)).Elem()
}

func (o CloudTieringDatePolicyStatusResponsePtrOutput) ToCloudTieringDatePolicyStatusResponsePtrOutput() CloudTieringDatePolicyStatusResponsePtrOutput {
	return o
}

func (o CloudTieringDatePolicyStatusResponsePtrOutput) ToCloudTieringDatePolicyStatusResponsePtrOutputWithContext(ctx context.Context) CloudTieringDatePolicyStatusResponsePtrOutput {
	return o
}

func (o CloudTieringDatePolicyStatusResponsePtrOutput) Elem() CloudTieringDatePolicyStatusResponseOutput {
	return o.ApplyT(func(v *CloudTieringDatePolicyStatusResponse) CloudTieringDatePolicyStatusResponse {
		if v != nil {
			return *v
		}
		var ret CloudTieringDatePolicyStatusResponse
		return ret
	}).(CloudTieringDatePolicyStatusResponseOutput)
}

func (o CloudTieringDatePolicyStatusResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudTieringDatePolicyStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o CloudTieringDatePolicyStatusResponsePtrOutput) TieredFilesMostRecentAccessTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudTieringDatePolicyStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TieredFilesMostRecentAccessTimestamp
	}).(pulumi.StringPtrOutput)
}

type CloudTieringFilesNotTieringResponse struct {
	Errors               []FilesNotTieringErrorResponse `pulumi:"errors"`
	LastUpdatedTimestamp string                         `pulumi:"lastUpdatedTimestamp"`
	TotalFileCount       float64                        `pulumi:"totalFileCount"`
}





type CloudTieringFilesNotTieringResponseInput interface {
	pulumi.Input

	ToCloudTieringFilesNotTieringResponseOutput() CloudTieringFilesNotTieringResponseOutput
	ToCloudTieringFilesNotTieringResponseOutputWithContext(context.Context) CloudTieringFilesNotTieringResponseOutput
}

type CloudTieringFilesNotTieringResponseArgs struct {
	Errors               FilesNotTieringErrorResponseArrayInput `pulumi:"errors"`
	LastUpdatedTimestamp pulumi.StringInput                     `pulumi:"lastUpdatedTimestamp"`
	TotalFileCount       pulumi.Float64Input                    `pulumi:"totalFileCount"`
}

func (CloudTieringFilesNotTieringResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringFilesNotTieringResponse)(nil)).Elem()
}

func (i CloudTieringFilesNotTieringResponseArgs) ToCloudTieringFilesNotTieringResponseOutput() CloudTieringFilesNotTieringResponseOutput {
	return i.ToCloudTieringFilesNotTieringResponseOutputWithContext(context.Background())
}

func (i CloudTieringFilesNotTieringResponseArgs) ToCloudTieringFilesNotTieringResponseOutputWithContext(ctx context.Context) CloudTieringFilesNotTieringResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringFilesNotTieringResponseOutput)
}

func (i CloudTieringFilesNotTieringResponseArgs) ToCloudTieringFilesNotTieringResponsePtrOutput() CloudTieringFilesNotTieringResponsePtrOutput {
	return i.ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(context.Background())
}

func (i CloudTieringFilesNotTieringResponseArgs) ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(ctx context.Context) CloudTieringFilesNotTieringResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringFilesNotTieringResponseOutput).ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(ctx)
}









type CloudTieringFilesNotTieringResponsePtrInput interface {
	pulumi.Input

	ToCloudTieringFilesNotTieringResponsePtrOutput() CloudTieringFilesNotTieringResponsePtrOutput
	ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(context.Context) CloudTieringFilesNotTieringResponsePtrOutput
}

type cloudTieringFilesNotTieringResponsePtrType CloudTieringFilesNotTieringResponseArgs

func CloudTieringFilesNotTieringResponsePtr(v *CloudTieringFilesNotTieringResponseArgs) CloudTieringFilesNotTieringResponsePtrInput {
	return (*cloudTieringFilesNotTieringResponsePtrType)(v)
}

func (*cloudTieringFilesNotTieringResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringFilesNotTieringResponse)(nil)).Elem()
}

func (i *cloudTieringFilesNotTieringResponsePtrType) ToCloudTieringFilesNotTieringResponsePtrOutput() CloudTieringFilesNotTieringResponsePtrOutput {
	return i.ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(context.Background())
}

func (i *cloudTieringFilesNotTieringResponsePtrType) ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(ctx context.Context) CloudTieringFilesNotTieringResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringFilesNotTieringResponsePtrOutput)
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

func (o CloudTieringFilesNotTieringResponseOutput) ToCloudTieringFilesNotTieringResponsePtrOutput() CloudTieringFilesNotTieringResponsePtrOutput {
	return o.ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(context.Background())
}

func (o CloudTieringFilesNotTieringResponseOutput) ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(ctx context.Context) CloudTieringFilesNotTieringResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudTieringFilesNotTieringResponse) *CloudTieringFilesNotTieringResponse {
		return &v
	}).(CloudTieringFilesNotTieringResponsePtrOutput)
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

type CloudTieringFilesNotTieringResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudTieringFilesNotTieringResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringFilesNotTieringResponse)(nil)).Elem()
}

func (o CloudTieringFilesNotTieringResponsePtrOutput) ToCloudTieringFilesNotTieringResponsePtrOutput() CloudTieringFilesNotTieringResponsePtrOutput {
	return o
}

func (o CloudTieringFilesNotTieringResponsePtrOutput) ToCloudTieringFilesNotTieringResponsePtrOutputWithContext(ctx context.Context) CloudTieringFilesNotTieringResponsePtrOutput {
	return o
}

func (o CloudTieringFilesNotTieringResponsePtrOutput) Elem() CloudTieringFilesNotTieringResponseOutput {
	return o.ApplyT(func(v *CloudTieringFilesNotTieringResponse) CloudTieringFilesNotTieringResponse {
		if v != nil {
			return *v
		}
		var ret CloudTieringFilesNotTieringResponse
		return ret
	}).(CloudTieringFilesNotTieringResponseOutput)
}

func (o CloudTieringFilesNotTieringResponsePtrOutput) Errors() FilesNotTieringErrorResponseArrayOutput {
	return o.ApplyT(func(v *CloudTieringFilesNotTieringResponse) []FilesNotTieringErrorResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(FilesNotTieringErrorResponseArrayOutput)
}

func (o CloudTieringFilesNotTieringResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudTieringFilesNotTieringResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o CloudTieringFilesNotTieringResponsePtrOutput) TotalFileCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudTieringFilesNotTieringResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TotalFileCount
	}).(pulumi.Float64PtrOutput)
}

type CloudTieringSpaceSavingsResponse struct {
	CachedSizeBytes      float64 `pulumi:"cachedSizeBytes"`
	LastUpdatedTimestamp string  `pulumi:"lastUpdatedTimestamp"`
	SpaceSavingsBytes    float64 `pulumi:"spaceSavingsBytes"`
	SpaceSavingsPercent  int     `pulumi:"spaceSavingsPercent"`
	TotalSizeCloudBytes  float64 `pulumi:"totalSizeCloudBytes"`
	VolumeSizeBytes      float64 `pulumi:"volumeSizeBytes"`
}





type CloudTieringSpaceSavingsResponseInput interface {
	pulumi.Input

	ToCloudTieringSpaceSavingsResponseOutput() CloudTieringSpaceSavingsResponseOutput
	ToCloudTieringSpaceSavingsResponseOutputWithContext(context.Context) CloudTieringSpaceSavingsResponseOutput
}

type CloudTieringSpaceSavingsResponseArgs struct {
	CachedSizeBytes      pulumi.Float64Input `pulumi:"cachedSizeBytes"`
	LastUpdatedTimestamp pulumi.StringInput  `pulumi:"lastUpdatedTimestamp"`
	SpaceSavingsBytes    pulumi.Float64Input `pulumi:"spaceSavingsBytes"`
	SpaceSavingsPercent  pulumi.IntInput     `pulumi:"spaceSavingsPercent"`
	TotalSizeCloudBytes  pulumi.Float64Input `pulumi:"totalSizeCloudBytes"`
	VolumeSizeBytes      pulumi.Float64Input `pulumi:"volumeSizeBytes"`
}

func (CloudTieringSpaceSavingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringSpaceSavingsResponse)(nil)).Elem()
}

func (i CloudTieringSpaceSavingsResponseArgs) ToCloudTieringSpaceSavingsResponseOutput() CloudTieringSpaceSavingsResponseOutput {
	return i.ToCloudTieringSpaceSavingsResponseOutputWithContext(context.Background())
}

func (i CloudTieringSpaceSavingsResponseArgs) ToCloudTieringSpaceSavingsResponseOutputWithContext(ctx context.Context) CloudTieringSpaceSavingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringSpaceSavingsResponseOutput)
}

func (i CloudTieringSpaceSavingsResponseArgs) ToCloudTieringSpaceSavingsResponsePtrOutput() CloudTieringSpaceSavingsResponsePtrOutput {
	return i.ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(context.Background())
}

func (i CloudTieringSpaceSavingsResponseArgs) ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(ctx context.Context) CloudTieringSpaceSavingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringSpaceSavingsResponseOutput).ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(ctx)
}









type CloudTieringSpaceSavingsResponsePtrInput interface {
	pulumi.Input

	ToCloudTieringSpaceSavingsResponsePtrOutput() CloudTieringSpaceSavingsResponsePtrOutput
	ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(context.Context) CloudTieringSpaceSavingsResponsePtrOutput
}

type cloudTieringSpaceSavingsResponsePtrType CloudTieringSpaceSavingsResponseArgs

func CloudTieringSpaceSavingsResponsePtr(v *CloudTieringSpaceSavingsResponseArgs) CloudTieringSpaceSavingsResponsePtrInput {
	return (*cloudTieringSpaceSavingsResponsePtrType)(v)
}

func (*cloudTieringSpaceSavingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringSpaceSavingsResponse)(nil)).Elem()
}

func (i *cloudTieringSpaceSavingsResponsePtrType) ToCloudTieringSpaceSavingsResponsePtrOutput() CloudTieringSpaceSavingsResponsePtrOutput {
	return i.ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(context.Background())
}

func (i *cloudTieringSpaceSavingsResponsePtrType) ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(ctx context.Context) CloudTieringSpaceSavingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringSpaceSavingsResponsePtrOutput)
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

func (o CloudTieringSpaceSavingsResponseOutput) ToCloudTieringSpaceSavingsResponsePtrOutput() CloudTieringSpaceSavingsResponsePtrOutput {
	return o.ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(context.Background())
}

func (o CloudTieringSpaceSavingsResponseOutput) ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(ctx context.Context) CloudTieringSpaceSavingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudTieringSpaceSavingsResponse) *CloudTieringSpaceSavingsResponse {
		return &v
	}).(CloudTieringSpaceSavingsResponsePtrOutput)
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

type CloudTieringSpaceSavingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudTieringSpaceSavingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringSpaceSavingsResponse)(nil)).Elem()
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) ToCloudTieringSpaceSavingsResponsePtrOutput() CloudTieringSpaceSavingsResponsePtrOutput {
	return o
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) ToCloudTieringSpaceSavingsResponsePtrOutputWithContext(ctx context.Context) CloudTieringSpaceSavingsResponsePtrOutput {
	return o
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) Elem() CloudTieringSpaceSavingsResponseOutput {
	return o.ApplyT(func(v *CloudTieringSpaceSavingsResponse) CloudTieringSpaceSavingsResponse {
		if v != nil {
			return *v
		}
		var ret CloudTieringSpaceSavingsResponse
		return ret
	}).(CloudTieringSpaceSavingsResponseOutput)
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) CachedSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudTieringSpaceSavingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.CachedSizeBytes
	}).(pulumi.Float64PtrOutput)
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudTieringSpaceSavingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) SpaceSavingsBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudTieringSpaceSavingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.SpaceSavingsBytes
	}).(pulumi.Float64PtrOutput)
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) SpaceSavingsPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudTieringSpaceSavingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.SpaceSavingsPercent
	}).(pulumi.IntPtrOutput)
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) TotalSizeCloudBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudTieringSpaceSavingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TotalSizeCloudBytes
	}).(pulumi.Float64PtrOutput)
}

func (o CloudTieringSpaceSavingsResponsePtrOutput) VolumeSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CloudTieringSpaceSavingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.VolumeSizeBytes
	}).(pulumi.Float64PtrOutput)
}

type CloudTieringVolumeFreeSpacePolicyStatusResponse struct {
	CurrentVolumeFreeSpacePercent  int    `pulumi:"currentVolumeFreeSpacePercent"`
	EffectiveVolumeFreeSpacePolicy int    `pulumi:"effectiveVolumeFreeSpacePolicy"`
	LastUpdatedTimestamp           string `pulumi:"lastUpdatedTimestamp"`
}





type CloudTieringVolumeFreeSpacePolicyStatusResponseInput interface {
	pulumi.Input

	ToCloudTieringVolumeFreeSpacePolicyStatusResponseOutput() CloudTieringVolumeFreeSpacePolicyStatusResponseOutput
	ToCloudTieringVolumeFreeSpacePolicyStatusResponseOutputWithContext(context.Context) CloudTieringVolumeFreeSpacePolicyStatusResponseOutput
}

type CloudTieringVolumeFreeSpacePolicyStatusResponseArgs struct {
	CurrentVolumeFreeSpacePercent  pulumi.IntInput    `pulumi:"currentVolumeFreeSpacePercent"`
	EffectiveVolumeFreeSpacePolicy pulumi.IntInput    `pulumi:"effectiveVolumeFreeSpacePolicy"`
	LastUpdatedTimestamp           pulumi.StringInput `pulumi:"lastUpdatedTimestamp"`
}

func (CloudTieringVolumeFreeSpacePolicyStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudTieringVolumeFreeSpacePolicyStatusResponse)(nil)).Elem()
}

func (i CloudTieringVolumeFreeSpacePolicyStatusResponseArgs) ToCloudTieringVolumeFreeSpacePolicyStatusResponseOutput() CloudTieringVolumeFreeSpacePolicyStatusResponseOutput {
	return i.ToCloudTieringVolumeFreeSpacePolicyStatusResponseOutputWithContext(context.Background())
}

func (i CloudTieringVolumeFreeSpacePolicyStatusResponseArgs) ToCloudTieringVolumeFreeSpacePolicyStatusResponseOutputWithContext(ctx context.Context) CloudTieringVolumeFreeSpacePolicyStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringVolumeFreeSpacePolicyStatusResponseOutput)
}

func (i CloudTieringVolumeFreeSpacePolicyStatusResponseArgs) ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput() CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return i.ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(context.Background())
}

func (i CloudTieringVolumeFreeSpacePolicyStatusResponseArgs) ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(ctx context.Context) CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringVolumeFreeSpacePolicyStatusResponseOutput).ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(ctx)
}









type CloudTieringVolumeFreeSpacePolicyStatusResponsePtrInput interface {
	pulumi.Input

	ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput() CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput
	ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(context.Context) CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput
}

type cloudTieringVolumeFreeSpacePolicyStatusResponsePtrType CloudTieringVolumeFreeSpacePolicyStatusResponseArgs

func CloudTieringVolumeFreeSpacePolicyStatusResponsePtr(v *CloudTieringVolumeFreeSpacePolicyStatusResponseArgs) CloudTieringVolumeFreeSpacePolicyStatusResponsePtrInput {
	return (*cloudTieringVolumeFreeSpacePolicyStatusResponsePtrType)(v)
}

func (*cloudTieringVolumeFreeSpacePolicyStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringVolumeFreeSpacePolicyStatusResponse)(nil)).Elem()
}

func (i *cloudTieringVolumeFreeSpacePolicyStatusResponsePtrType) ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput() CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return i.ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(context.Background())
}

func (i *cloudTieringVolumeFreeSpacePolicyStatusResponsePtrType) ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(ctx context.Context) CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput)
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

func (o CloudTieringVolumeFreeSpacePolicyStatusResponseOutput) ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput() CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return o.ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(context.Background())
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponseOutput) ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(ctx context.Context) CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudTieringVolumeFreeSpacePolicyStatusResponse) *CloudTieringVolumeFreeSpacePolicyStatusResponse {
		return &v
	}).(CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput)
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

type CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudTieringVolumeFreeSpacePolicyStatusResponse)(nil)).Elem()
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput) ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput() CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return o
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput) ToCloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutputWithContext(ctx context.Context) CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return o
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput) Elem() CloudTieringVolumeFreeSpacePolicyStatusResponseOutput {
	return o.ApplyT(func(v *CloudTieringVolumeFreeSpacePolicyStatusResponse) CloudTieringVolumeFreeSpacePolicyStatusResponse {
		if v != nil {
			return *v
		}
		var ret CloudTieringVolumeFreeSpacePolicyStatusResponse
		return ret
	}).(CloudTieringVolumeFreeSpacePolicyStatusResponseOutput)
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput) CurrentVolumeFreeSpacePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudTieringVolumeFreeSpacePolicyStatusResponse) *int {
		if v == nil {
			return nil
		}
		return &v.CurrentVolumeFreeSpacePercent
	}).(pulumi.IntPtrOutput)
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput) EffectiveVolumeFreeSpacePolicy() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudTieringVolumeFreeSpacePolicyStatusResponse) *int {
		if v == nil {
			return nil
		}
		return &v.EffectiveVolumeFreeSpacePolicy
	}).(pulumi.IntPtrOutput)
}

func (o CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudTieringVolumeFreeSpacePolicyStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

type FilesNotTieringErrorResponse struct {
	ErrorCode int     `pulumi:"errorCode"`
	FileCount float64 `pulumi:"fileCount"`
}





type FilesNotTieringErrorResponseInput interface {
	pulumi.Input

	ToFilesNotTieringErrorResponseOutput() FilesNotTieringErrorResponseOutput
	ToFilesNotTieringErrorResponseOutputWithContext(context.Context) FilesNotTieringErrorResponseOutput
}

type FilesNotTieringErrorResponseArgs struct {
	ErrorCode pulumi.IntInput     `pulumi:"errorCode"`
	FileCount pulumi.Float64Input `pulumi:"fileCount"`
}

func (FilesNotTieringErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilesNotTieringErrorResponse)(nil)).Elem()
}

func (i FilesNotTieringErrorResponseArgs) ToFilesNotTieringErrorResponseOutput() FilesNotTieringErrorResponseOutput {
	return i.ToFilesNotTieringErrorResponseOutputWithContext(context.Background())
}

func (i FilesNotTieringErrorResponseArgs) ToFilesNotTieringErrorResponseOutputWithContext(ctx context.Context) FilesNotTieringErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilesNotTieringErrorResponseOutput)
}





type FilesNotTieringErrorResponseArrayInput interface {
	pulumi.Input

	ToFilesNotTieringErrorResponseArrayOutput() FilesNotTieringErrorResponseArrayOutput
	ToFilesNotTieringErrorResponseArrayOutputWithContext(context.Context) FilesNotTieringErrorResponseArrayOutput
}

type FilesNotTieringErrorResponseArray []FilesNotTieringErrorResponseInput

func (FilesNotTieringErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilesNotTieringErrorResponse)(nil)).Elem()
}

func (i FilesNotTieringErrorResponseArray) ToFilesNotTieringErrorResponseArrayOutput() FilesNotTieringErrorResponseArrayOutput {
	return i.ToFilesNotTieringErrorResponseArrayOutputWithContext(context.Background())
}

func (i FilesNotTieringErrorResponseArray) ToFilesNotTieringErrorResponseArrayOutputWithContext(ctx context.Context) FilesNotTieringErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilesNotTieringErrorResponseArrayOutput)
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





type ServerEndpointCloudTieringStatusResponseInput interface {
	pulumi.Input

	ToServerEndpointCloudTieringStatusResponseOutput() ServerEndpointCloudTieringStatusResponseOutput
	ToServerEndpointCloudTieringStatusResponseOutputWithContext(context.Context) ServerEndpointCloudTieringStatusResponseOutput
}

type ServerEndpointCloudTieringStatusResponseArgs struct {
	CachePerformance            CloudTieringCachePerformanceResponseInput            `pulumi:"cachePerformance"`
	DatePolicyStatus            CloudTieringDatePolicyStatusResponseInput            `pulumi:"datePolicyStatus"`
	FilesNotTiering             CloudTieringFilesNotTieringResponseInput             `pulumi:"filesNotTiering"`
	Health                      pulumi.StringInput                                   `pulumi:"health"`
	HealthLastUpdatedTimestamp  pulumi.StringInput                                   `pulumi:"healthLastUpdatedTimestamp"`
	LastCloudTieringResult      pulumi.IntInput                                      `pulumi:"lastCloudTieringResult"`
	LastSuccessTimestamp        pulumi.StringInput                                   `pulumi:"lastSuccessTimestamp"`
	LastUpdatedTimestamp        pulumi.StringInput                                   `pulumi:"lastUpdatedTimestamp"`
	SpaceSavings                CloudTieringSpaceSavingsResponseInput                `pulumi:"spaceSavings"`
	VolumeFreeSpacePolicyStatus CloudTieringVolumeFreeSpacePolicyStatusResponseInput `pulumi:"volumeFreeSpacePolicyStatus"`
}

func (ServerEndpointCloudTieringStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointCloudTieringStatusResponse)(nil)).Elem()
}

func (i ServerEndpointCloudTieringStatusResponseArgs) ToServerEndpointCloudTieringStatusResponseOutput() ServerEndpointCloudTieringStatusResponseOutput {
	return i.ToServerEndpointCloudTieringStatusResponseOutputWithContext(context.Background())
}

func (i ServerEndpointCloudTieringStatusResponseArgs) ToServerEndpointCloudTieringStatusResponseOutputWithContext(ctx context.Context) ServerEndpointCloudTieringStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointCloudTieringStatusResponseOutput)
}

func (i ServerEndpointCloudTieringStatusResponseArgs) ToServerEndpointCloudTieringStatusResponsePtrOutput() ServerEndpointCloudTieringStatusResponsePtrOutput {
	return i.ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(context.Background())
}

func (i ServerEndpointCloudTieringStatusResponseArgs) ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointCloudTieringStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointCloudTieringStatusResponseOutput).ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(ctx)
}









type ServerEndpointCloudTieringStatusResponsePtrInput interface {
	pulumi.Input

	ToServerEndpointCloudTieringStatusResponsePtrOutput() ServerEndpointCloudTieringStatusResponsePtrOutput
	ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(context.Context) ServerEndpointCloudTieringStatusResponsePtrOutput
}

type serverEndpointCloudTieringStatusResponsePtrType ServerEndpointCloudTieringStatusResponseArgs

func ServerEndpointCloudTieringStatusResponsePtr(v *ServerEndpointCloudTieringStatusResponseArgs) ServerEndpointCloudTieringStatusResponsePtrInput {
	return (*serverEndpointCloudTieringStatusResponsePtrType)(v)
}

func (*serverEndpointCloudTieringStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointCloudTieringStatusResponse)(nil)).Elem()
}

func (i *serverEndpointCloudTieringStatusResponsePtrType) ToServerEndpointCloudTieringStatusResponsePtrOutput() ServerEndpointCloudTieringStatusResponsePtrOutput {
	return i.ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(context.Background())
}

func (i *serverEndpointCloudTieringStatusResponsePtrType) ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointCloudTieringStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointCloudTieringStatusResponsePtrOutput)
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

func (o ServerEndpointCloudTieringStatusResponseOutput) ToServerEndpointCloudTieringStatusResponsePtrOutput() ServerEndpointCloudTieringStatusResponsePtrOutput {
	return o.ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(context.Background())
}

func (o ServerEndpointCloudTieringStatusResponseOutput) ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointCloudTieringStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerEndpointCloudTieringStatusResponse) *ServerEndpointCloudTieringStatusResponse {
		return &v
	}).(ServerEndpointCloudTieringStatusResponsePtrOutput)
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

type ServerEndpointCloudTieringStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerEndpointCloudTieringStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointCloudTieringStatusResponse)(nil)).Elem()
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) ToServerEndpointCloudTieringStatusResponsePtrOutput() ServerEndpointCloudTieringStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) ToServerEndpointCloudTieringStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointCloudTieringStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) Elem() ServerEndpointCloudTieringStatusResponseOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) ServerEndpointCloudTieringStatusResponse {
		if v != nil {
			return *v
		}
		var ret ServerEndpointCloudTieringStatusResponse
		return ret
	}).(ServerEndpointCloudTieringStatusResponseOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) CachePerformance() CloudTieringCachePerformanceResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *CloudTieringCachePerformanceResponse {
		if v == nil {
			return nil
		}
		return &v.CachePerformance
	}).(CloudTieringCachePerformanceResponsePtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) DatePolicyStatus() CloudTieringDatePolicyStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *CloudTieringDatePolicyStatusResponse {
		if v == nil {
			return nil
		}
		return &v.DatePolicyStatus
	}).(CloudTieringDatePolicyStatusResponsePtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) FilesNotTiering() CloudTieringFilesNotTieringResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *CloudTieringFilesNotTieringResponse {
		if v == nil {
			return nil
		}
		return &v.FilesNotTiering
	}).(CloudTieringFilesNotTieringResponsePtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Health
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) HealthLastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthLastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) LastCloudTieringResult() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *int {
		if v == nil {
			return nil
		}
		return &v.LastCloudTieringResult
	}).(pulumi.IntPtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) LastSuccessTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSuccessTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) SpaceSavings() CloudTieringSpaceSavingsResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *CloudTieringSpaceSavingsResponse {
		if v == nil {
			return nil
		}
		return &v.SpaceSavings
	}).(CloudTieringSpaceSavingsResponsePtrOutput)
}

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) VolumeFreeSpacePolicyStatus() CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *CloudTieringVolumeFreeSpacePolicyStatusResponse {
		if v == nil {
			return nil
		}
		return &v.VolumeFreeSpacePolicyStatus
	}).(CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput)
}

type ServerEndpointFilesNotSyncingErrorResponse struct {
	ErrorCode       int     `pulumi:"errorCode"`
	PersistentCount float64 `pulumi:"persistentCount"`
	TransientCount  float64 `pulumi:"transientCount"`
}





type ServerEndpointFilesNotSyncingErrorResponseInput interface {
	pulumi.Input

	ToServerEndpointFilesNotSyncingErrorResponseOutput() ServerEndpointFilesNotSyncingErrorResponseOutput
	ToServerEndpointFilesNotSyncingErrorResponseOutputWithContext(context.Context) ServerEndpointFilesNotSyncingErrorResponseOutput
}

type ServerEndpointFilesNotSyncingErrorResponseArgs struct {
	ErrorCode       pulumi.IntInput     `pulumi:"errorCode"`
	PersistentCount pulumi.Float64Input `pulumi:"persistentCount"`
	TransientCount  pulumi.Float64Input `pulumi:"transientCount"`
}

func (ServerEndpointFilesNotSyncingErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointFilesNotSyncingErrorResponse)(nil)).Elem()
}

func (i ServerEndpointFilesNotSyncingErrorResponseArgs) ToServerEndpointFilesNotSyncingErrorResponseOutput() ServerEndpointFilesNotSyncingErrorResponseOutput {
	return i.ToServerEndpointFilesNotSyncingErrorResponseOutputWithContext(context.Background())
}

func (i ServerEndpointFilesNotSyncingErrorResponseArgs) ToServerEndpointFilesNotSyncingErrorResponseOutputWithContext(ctx context.Context) ServerEndpointFilesNotSyncingErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointFilesNotSyncingErrorResponseOutput)
}





type ServerEndpointFilesNotSyncingErrorResponseArrayInput interface {
	pulumi.Input

	ToServerEndpointFilesNotSyncingErrorResponseArrayOutput() ServerEndpointFilesNotSyncingErrorResponseArrayOutput
	ToServerEndpointFilesNotSyncingErrorResponseArrayOutputWithContext(context.Context) ServerEndpointFilesNotSyncingErrorResponseArrayOutput
}

type ServerEndpointFilesNotSyncingErrorResponseArray []ServerEndpointFilesNotSyncingErrorResponseInput

func (ServerEndpointFilesNotSyncingErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerEndpointFilesNotSyncingErrorResponse)(nil)).Elem()
}

func (i ServerEndpointFilesNotSyncingErrorResponseArray) ToServerEndpointFilesNotSyncingErrorResponseArrayOutput() ServerEndpointFilesNotSyncingErrorResponseArrayOutput {
	return i.ToServerEndpointFilesNotSyncingErrorResponseArrayOutputWithContext(context.Background())
}

func (i ServerEndpointFilesNotSyncingErrorResponseArray) ToServerEndpointFilesNotSyncingErrorResponseArrayOutputWithContext(ctx context.Context) ServerEndpointFilesNotSyncingErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointFilesNotSyncingErrorResponseArrayOutput)
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





type ServerEndpointRecallErrorResponseInput interface {
	pulumi.Input

	ToServerEndpointRecallErrorResponseOutput() ServerEndpointRecallErrorResponseOutput
	ToServerEndpointRecallErrorResponseOutputWithContext(context.Context) ServerEndpointRecallErrorResponseOutput
}

type ServerEndpointRecallErrorResponseArgs struct {
	Count     pulumi.Float64Input `pulumi:"count"`
	ErrorCode pulumi.IntInput     `pulumi:"errorCode"`
}

func (ServerEndpointRecallErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointRecallErrorResponse)(nil)).Elem()
}

func (i ServerEndpointRecallErrorResponseArgs) ToServerEndpointRecallErrorResponseOutput() ServerEndpointRecallErrorResponseOutput {
	return i.ToServerEndpointRecallErrorResponseOutputWithContext(context.Background())
}

func (i ServerEndpointRecallErrorResponseArgs) ToServerEndpointRecallErrorResponseOutputWithContext(ctx context.Context) ServerEndpointRecallErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointRecallErrorResponseOutput)
}





type ServerEndpointRecallErrorResponseArrayInput interface {
	pulumi.Input

	ToServerEndpointRecallErrorResponseArrayOutput() ServerEndpointRecallErrorResponseArrayOutput
	ToServerEndpointRecallErrorResponseArrayOutputWithContext(context.Context) ServerEndpointRecallErrorResponseArrayOutput
}

type ServerEndpointRecallErrorResponseArray []ServerEndpointRecallErrorResponseInput

func (ServerEndpointRecallErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerEndpointRecallErrorResponse)(nil)).Elem()
}

func (i ServerEndpointRecallErrorResponseArray) ToServerEndpointRecallErrorResponseArrayOutput() ServerEndpointRecallErrorResponseArrayOutput {
	return i.ToServerEndpointRecallErrorResponseArrayOutputWithContext(context.Background())
}

func (i ServerEndpointRecallErrorResponseArray) ToServerEndpointRecallErrorResponseArrayOutputWithContext(ctx context.Context) ServerEndpointRecallErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointRecallErrorResponseArrayOutput)
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





type ServerEndpointRecallStatusResponseInput interface {
	pulumi.Input

	ToServerEndpointRecallStatusResponseOutput() ServerEndpointRecallStatusResponseOutput
	ToServerEndpointRecallStatusResponseOutputWithContext(context.Context) ServerEndpointRecallStatusResponseOutput
}

type ServerEndpointRecallStatusResponseArgs struct {
	LastUpdatedTimestamp   pulumi.StringInput                          `pulumi:"lastUpdatedTimestamp"`
	RecallErrors           ServerEndpointRecallErrorResponseArrayInput `pulumi:"recallErrors"`
	TotalRecallErrorsCount pulumi.Float64Input                         `pulumi:"totalRecallErrorsCount"`
}

func (ServerEndpointRecallStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointRecallStatusResponse)(nil)).Elem()
}

func (i ServerEndpointRecallStatusResponseArgs) ToServerEndpointRecallStatusResponseOutput() ServerEndpointRecallStatusResponseOutput {
	return i.ToServerEndpointRecallStatusResponseOutputWithContext(context.Background())
}

func (i ServerEndpointRecallStatusResponseArgs) ToServerEndpointRecallStatusResponseOutputWithContext(ctx context.Context) ServerEndpointRecallStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointRecallStatusResponseOutput)
}

func (i ServerEndpointRecallStatusResponseArgs) ToServerEndpointRecallStatusResponsePtrOutput() ServerEndpointRecallStatusResponsePtrOutput {
	return i.ToServerEndpointRecallStatusResponsePtrOutputWithContext(context.Background())
}

func (i ServerEndpointRecallStatusResponseArgs) ToServerEndpointRecallStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointRecallStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointRecallStatusResponseOutput).ToServerEndpointRecallStatusResponsePtrOutputWithContext(ctx)
}









type ServerEndpointRecallStatusResponsePtrInput interface {
	pulumi.Input

	ToServerEndpointRecallStatusResponsePtrOutput() ServerEndpointRecallStatusResponsePtrOutput
	ToServerEndpointRecallStatusResponsePtrOutputWithContext(context.Context) ServerEndpointRecallStatusResponsePtrOutput
}

type serverEndpointRecallStatusResponsePtrType ServerEndpointRecallStatusResponseArgs

func ServerEndpointRecallStatusResponsePtr(v *ServerEndpointRecallStatusResponseArgs) ServerEndpointRecallStatusResponsePtrInput {
	return (*serverEndpointRecallStatusResponsePtrType)(v)
}

func (*serverEndpointRecallStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointRecallStatusResponse)(nil)).Elem()
}

func (i *serverEndpointRecallStatusResponsePtrType) ToServerEndpointRecallStatusResponsePtrOutput() ServerEndpointRecallStatusResponsePtrOutput {
	return i.ToServerEndpointRecallStatusResponsePtrOutputWithContext(context.Background())
}

func (i *serverEndpointRecallStatusResponsePtrType) ToServerEndpointRecallStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointRecallStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointRecallStatusResponsePtrOutput)
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

func (o ServerEndpointRecallStatusResponseOutput) ToServerEndpointRecallStatusResponsePtrOutput() ServerEndpointRecallStatusResponsePtrOutput {
	return o.ToServerEndpointRecallStatusResponsePtrOutputWithContext(context.Background())
}

func (o ServerEndpointRecallStatusResponseOutput) ToServerEndpointRecallStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointRecallStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerEndpointRecallStatusResponse) *ServerEndpointRecallStatusResponse {
		return &v
	}).(ServerEndpointRecallStatusResponsePtrOutput)
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

type ServerEndpointRecallStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerEndpointRecallStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointRecallStatusResponse)(nil)).Elem()
}

func (o ServerEndpointRecallStatusResponsePtrOutput) ToServerEndpointRecallStatusResponsePtrOutput() ServerEndpointRecallStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointRecallStatusResponsePtrOutput) ToServerEndpointRecallStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointRecallStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointRecallStatusResponsePtrOutput) Elem() ServerEndpointRecallStatusResponseOutput {
	return o.ApplyT(func(v *ServerEndpointRecallStatusResponse) ServerEndpointRecallStatusResponse {
		if v != nil {
			return *v
		}
		var ret ServerEndpointRecallStatusResponse
		return ret
	}).(ServerEndpointRecallStatusResponseOutput)
}

func (o ServerEndpointRecallStatusResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointRecallStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointRecallStatusResponsePtrOutput) RecallErrors() ServerEndpointRecallErrorResponseArrayOutput {
	return o.ApplyT(func(v *ServerEndpointRecallStatusResponse) []ServerEndpointRecallErrorResponse {
		if v == nil {
			return nil
		}
		return v.RecallErrors
	}).(ServerEndpointRecallErrorResponseArrayOutput)
}

func (o ServerEndpointRecallStatusResponsePtrOutput) TotalRecallErrorsCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointRecallStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TotalRecallErrorsCount
	}).(pulumi.Float64PtrOutput)
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
	FilesNotSyncingErrors          []ServerEndpointFilesNotSyncingErrorResponse `pulumi:"filesNotSyncingErrors"`
	LastSyncPerItemErrorCount      float64                                      `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult                 int                                          `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp       string                                       `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp              string                                       `pulumi:"lastSyncTimestamp"`
	PersistentFilesNotSyncingCount float64                                      `pulumi:"persistentFilesNotSyncingCount"`
	TransientFilesNotSyncingCount  float64                                      `pulumi:"transientFilesNotSyncingCount"`
}





type SyncSessionStatusResponseInput interface {
	pulumi.Input

	ToSyncSessionStatusResponseOutput() SyncSessionStatusResponseOutput
	ToSyncSessionStatusResponseOutputWithContext(context.Context) SyncSessionStatusResponseOutput
}

type SyncSessionStatusResponseArgs struct {
	FilesNotSyncingErrors          ServerEndpointFilesNotSyncingErrorResponseArrayInput `pulumi:"filesNotSyncingErrors"`
	LastSyncPerItemErrorCount      pulumi.Float64Input                                  `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult                 pulumi.IntInput                                      `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp       pulumi.StringInput                                   `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp              pulumi.StringInput                                   `pulumi:"lastSyncTimestamp"`
	PersistentFilesNotSyncingCount pulumi.Float64Input                                  `pulumi:"persistentFilesNotSyncingCount"`
	TransientFilesNotSyncingCount  pulumi.Float64Input                                  `pulumi:"transientFilesNotSyncingCount"`
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

func (o SyncSessionStatusResponseOutput) FilesNotSyncingErrors() ServerEndpointFilesNotSyncingErrorResponseArrayOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) []ServerEndpointFilesNotSyncingErrorResponse {
		return v.FilesNotSyncingErrors
	}).(ServerEndpointFilesNotSyncingErrorResponseArrayOutput)
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

func (o SyncSessionStatusResponsePtrOutput) FilesNotSyncingErrors() ServerEndpointFilesNotSyncingErrorResponseArrayOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) []ServerEndpointFilesNotSyncingErrorResponse {
		if v == nil {
			return nil
		}
		return v.FilesNotSyncingErrors
	}).(ServerEndpointFilesNotSyncingErrorResponseArrayOutput)
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
	pulumi.RegisterOutputType(CloudTieringCachePerformanceResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringCachePerformanceResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudTieringDatePolicyStatusResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringDatePolicyStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudTieringFilesNotTieringResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringFilesNotTieringResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudTieringSpaceSavingsResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringSpaceSavingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudTieringVolumeFreeSpacePolicyStatusResponseOutput{})
	pulumi.RegisterOutputType(CloudTieringVolumeFreeSpacePolicyStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(FilesNotTieringErrorResponseOutput{})
	pulumi.RegisterOutputType(FilesNotTieringErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointCloudTieringStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointCloudTieringStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerEndpointFilesNotSyncingErrorResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointFilesNotSyncingErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallErrorResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncActivityStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncActivityStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponsePtrOutput{})
}
