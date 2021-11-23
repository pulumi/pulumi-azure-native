


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





type ServerEndpointCloudTieringStatusResponseInput interface {
	pulumi.Input

	ToServerEndpointCloudTieringStatusResponseOutput() ServerEndpointCloudTieringStatusResponseOutput
	ToServerEndpointCloudTieringStatusResponseOutputWithContext(context.Context) ServerEndpointCloudTieringStatusResponseOutput
}

type ServerEndpointCloudTieringStatusResponseArgs struct {
	Health                 pulumi.StringInput `pulumi:"health"`
	LastCloudTieringResult pulumi.IntInput    `pulumi:"lastCloudTieringResult"`
	LastSuccessTimestamp   pulumi.StringInput `pulumi:"lastSuccessTimestamp"`
	LastUpdatedTimestamp   pulumi.StringInput `pulumi:"lastUpdatedTimestamp"`
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

func (o ServerEndpointCloudTieringStatusResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointCloudTieringStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Health
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

type ServerEndpointSyncActivityStatusResponse struct {
	AppliedBytes      float64 `pulumi:"appliedBytes"`
	AppliedItemCount  float64 `pulumi:"appliedItemCount"`
	PerItemErrorCount float64 `pulumi:"perItemErrorCount"`
	Timestamp         string  `pulumi:"timestamp"`
	TotalBytes        float64 `pulumi:"totalBytes"`
	TotalItemCount    float64 `pulumi:"totalItemCount"`
}





type ServerEndpointSyncActivityStatusResponseInput interface {
	pulumi.Input

	ToServerEndpointSyncActivityStatusResponseOutput() ServerEndpointSyncActivityStatusResponseOutput
	ToServerEndpointSyncActivityStatusResponseOutputWithContext(context.Context) ServerEndpointSyncActivityStatusResponseOutput
}

type ServerEndpointSyncActivityStatusResponseArgs struct {
	AppliedBytes      pulumi.Float64Input `pulumi:"appliedBytes"`
	AppliedItemCount  pulumi.Float64Input `pulumi:"appliedItemCount"`
	PerItemErrorCount pulumi.Float64Input `pulumi:"perItemErrorCount"`
	Timestamp         pulumi.StringInput  `pulumi:"timestamp"`
	TotalBytes        pulumi.Float64Input `pulumi:"totalBytes"`
	TotalItemCount    pulumi.Float64Input `pulumi:"totalItemCount"`
}

func (ServerEndpointSyncActivityStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointSyncActivityStatusResponse)(nil)).Elem()
}

func (i ServerEndpointSyncActivityStatusResponseArgs) ToServerEndpointSyncActivityStatusResponseOutput() ServerEndpointSyncActivityStatusResponseOutput {
	return i.ToServerEndpointSyncActivityStatusResponseOutputWithContext(context.Background())
}

func (i ServerEndpointSyncActivityStatusResponseArgs) ToServerEndpointSyncActivityStatusResponseOutputWithContext(ctx context.Context) ServerEndpointSyncActivityStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncActivityStatusResponseOutput)
}

func (i ServerEndpointSyncActivityStatusResponseArgs) ToServerEndpointSyncActivityStatusResponsePtrOutput() ServerEndpointSyncActivityStatusResponsePtrOutput {
	return i.ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(context.Background())
}

func (i ServerEndpointSyncActivityStatusResponseArgs) ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncActivityStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncActivityStatusResponseOutput).ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(ctx)
}









type ServerEndpointSyncActivityStatusResponsePtrInput interface {
	pulumi.Input

	ToServerEndpointSyncActivityStatusResponsePtrOutput() ServerEndpointSyncActivityStatusResponsePtrOutput
	ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(context.Context) ServerEndpointSyncActivityStatusResponsePtrOutput
}

type serverEndpointSyncActivityStatusResponsePtrType ServerEndpointSyncActivityStatusResponseArgs

func ServerEndpointSyncActivityStatusResponsePtr(v *ServerEndpointSyncActivityStatusResponseArgs) ServerEndpointSyncActivityStatusResponsePtrInput {
	return (*serverEndpointSyncActivityStatusResponsePtrType)(v)
}

func (*serverEndpointSyncActivityStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointSyncActivityStatusResponse)(nil)).Elem()
}

func (i *serverEndpointSyncActivityStatusResponsePtrType) ToServerEndpointSyncActivityStatusResponsePtrOutput() ServerEndpointSyncActivityStatusResponsePtrOutput {
	return i.ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(context.Background())
}

func (i *serverEndpointSyncActivityStatusResponsePtrType) ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncActivityStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncActivityStatusResponsePtrOutput)
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

func (o ServerEndpointSyncActivityStatusResponseOutput) ToServerEndpointSyncActivityStatusResponsePtrOutput() ServerEndpointSyncActivityStatusResponsePtrOutput {
	return o.ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(context.Background())
}

func (o ServerEndpointSyncActivityStatusResponseOutput) ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncActivityStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerEndpointSyncActivityStatusResponse) *ServerEndpointSyncActivityStatusResponse {
		return &v
	}).(ServerEndpointSyncActivityStatusResponsePtrOutput)
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

type ServerEndpointSyncActivityStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerEndpointSyncActivityStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointSyncActivityStatusResponse)(nil)).Elem()
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) ToServerEndpointSyncActivityStatusResponsePtrOutput() ServerEndpointSyncActivityStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) ToServerEndpointSyncActivityStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncActivityStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) Elem() ServerEndpointSyncActivityStatusResponseOutput {
	return o.ApplyT(func(v *ServerEndpointSyncActivityStatusResponse) ServerEndpointSyncActivityStatusResponse {
		if v != nil {
			return *v
		}
		var ret ServerEndpointSyncActivityStatusResponse
		return ret
	}).(ServerEndpointSyncActivityStatusResponseOutput)
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) AppliedBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.AppliedBytes
	}).(pulumi.Float64PtrOutput)
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) AppliedItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.AppliedItemCount
	}).(pulumi.Float64PtrOutput)
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) PerItemErrorCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.PerItemErrorCount
	}).(pulumi.Float64PtrOutput)
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncActivityStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Timestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) TotalBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TotalBytes
	}).(pulumi.Float64PtrOutput)
}

func (o ServerEndpointSyncActivityStatusResponsePtrOutput) TotalItemCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncActivityStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TotalItemCount
	}).(pulumi.Float64PtrOutput)
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





type ServerEndpointSyncSessionStatusResponseInput interface {
	pulumi.Input

	ToServerEndpointSyncSessionStatusResponseOutput() ServerEndpointSyncSessionStatusResponseOutput
	ToServerEndpointSyncSessionStatusResponseOutputWithContext(context.Context) ServerEndpointSyncSessionStatusResponseOutput
}

type ServerEndpointSyncSessionStatusResponseArgs struct {
	FilesNotSyncingErrors          ServerEndpointFilesNotSyncingErrorResponseArrayInput `pulumi:"filesNotSyncingErrors"`
	LastSyncPerItemErrorCount      pulumi.Float64Input                                  `pulumi:"lastSyncPerItemErrorCount"`
	LastSyncResult                 pulumi.IntInput                                      `pulumi:"lastSyncResult"`
	LastSyncSuccessTimestamp       pulumi.StringInput                                   `pulumi:"lastSyncSuccessTimestamp"`
	LastSyncTimestamp              pulumi.StringInput                                   `pulumi:"lastSyncTimestamp"`
	PersistentFilesNotSyncingCount pulumi.Float64Input                                  `pulumi:"persistentFilesNotSyncingCount"`
	TransientFilesNotSyncingCount  pulumi.Float64Input                                  `pulumi:"transientFilesNotSyncingCount"`
}

func (ServerEndpointSyncSessionStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointSyncSessionStatusResponse)(nil)).Elem()
}

func (i ServerEndpointSyncSessionStatusResponseArgs) ToServerEndpointSyncSessionStatusResponseOutput() ServerEndpointSyncSessionStatusResponseOutput {
	return i.ToServerEndpointSyncSessionStatusResponseOutputWithContext(context.Background())
}

func (i ServerEndpointSyncSessionStatusResponseArgs) ToServerEndpointSyncSessionStatusResponseOutputWithContext(ctx context.Context) ServerEndpointSyncSessionStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncSessionStatusResponseOutput)
}

func (i ServerEndpointSyncSessionStatusResponseArgs) ToServerEndpointSyncSessionStatusResponsePtrOutput() ServerEndpointSyncSessionStatusResponsePtrOutput {
	return i.ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (i ServerEndpointSyncSessionStatusResponseArgs) ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncSessionStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncSessionStatusResponseOutput).ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(ctx)
}









type ServerEndpointSyncSessionStatusResponsePtrInput interface {
	pulumi.Input

	ToServerEndpointSyncSessionStatusResponsePtrOutput() ServerEndpointSyncSessionStatusResponsePtrOutput
	ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(context.Context) ServerEndpointSyncSessionStatusResponsePtrOutput
}

type serverEndpointSyncSessionStatusResponsePtrType ServerEndpointSyncSessionStatusResponseArgs

func ServerEndpointSyncSessionStatusResponsePtr(v *ServerEndpointSyncSessionStatusResponseArgs) ServerEndpointSyncSessionStatusResponsePtrInput {
	return (*serverEndpointSyncSessionStatusResponsePtrType)(v)
}

func (*serverEndpointSyncSessionStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointSyncSessionStatusResponse)(nil)).Elem()
}

func (i *serverEndpointSyncSessionStatusResponsePtrType) ToServerEndpointSyncSessionStatusResponsePtrOutput() ServerEndpointSyncSessionStatusResponsePtrOutput {
	return i.ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (i *serverEndpointSyncSessionStatusResponsePtrType) ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncSessionStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointSyncSessionStatusResponsePtrOutput)
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

func (o ServerEndpointSyncSessionStatusResponseOutput) ToServerEndpointSyncSessionStatusResponsePtrOutput() ServerEndpointSyncSessionStatusResponsePtrOutput {
	return o.ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (o ServerEndpointSyncSessionStatusResponseOutput) ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncSessionStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerEndpointSyncSessionStatusResponse) *ServerEndpointSyncSessionStatusResponse {
		return &v
	}).(ServerEndpointSyncSessionStatusResponsePtrOutput)
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

type ServerEndpointSyncSessionStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerEndpointSyncSessionStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointSyncSessionStatusResponse)(nil)).Elem()
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) ToServerEndpointSyncSessionStatusResponsePtrOutput() ServerEndpointSyncSessionStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) ToServerEndpointSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) ServerEndpointSyncSessionStatusResponsePtrOutput {
	return o
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) Elem() ServerEndpointSyncSessionStatusResponseOutput {
	return o.ApplyT(func(v *ServerEndpointSyncSessionStatusResponse) ServerEndpointSyncSessionStatusResponse {
		if v != nil {
			return *v
		}
		var ret ServerEndpointSyncSessionStatusResponse
		return ret
	}).(ServerEndpointSyncSessionStatusResponseOutput)
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) FilesNotSyncingErrors() ServerEndpointFilesNotSyncingErrorResponseArrayOutput {
	return o.ApplyT(func(v *ServerEndpointSyncSessionStatusResponse) []ServerEndpointFilesNotSyncingErrorResponse {
		if v == nil {
			return nil
		}
		return v.FilesNotSyncingErrors
	}).(ServerEndpointFilesNotSyncingErrorResponseArrayOutput)
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) LastSyncPerItemErrorCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncSessionStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.LastSyncPerItemErrorCount
	}).(pulumi.Float64PtrOutput)
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) LastSyncResult() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncSessionStatusResponse) *int {
		if v == nil {
			return nil
		}
		return &v.LastSyncResult
	}).(pulumi.IntPtrOutput)
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) LastSyncSuccessTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncSessionStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSyncSuccessTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) LastSyncTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncSessionStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSyncTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) PersistentFilesNotSyncingCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncSessionStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.PersistentFilesNotSyncingCount
	}).(pulumi.Float64PtrOutput)
}

func (o ServerEndpointSyncSessionStatusResponsePtrOutput) TransientFilesNotSyncingCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncSessionStatusResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransientFilesNotSyncingCount
	}).(pulumi.Float64PtrOutput)
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





type ServerEndpointSyncStatusResponseInput interface {
	pulumi.Input

	ToServerEndpointSyncStatusResponseOutput() ServerEndpointSyncStatusResponseOutput
	ToServerEndpointSyncStatusResponseOutputWithContext(context.Context) ServerEndpointSyncStatusResponseOutput
}

type ServerEndpointSyncStatusResponseArgs struct {
	CombinedHealth                      pulumi.StringInput                            `pulumi:"combinedHealth"`
	DownloadActivity                    ServerEndpointSyncActivityStatusResponseInput `pulumi:"downloadActivity"`
	DownloadHealth                      pulumi.StringInput                            `pulumi:"downloadHealth"`
	DownloadStatus                      ServerEndpointSyncSessionStatusResponseInput  `pulumi:"downloadStatus"`
	LastUpdatedTimestamp                pulumi.StringInput                            `pulumi:"lastUpdatedTimestamp"`
	OfflineDataTransferStatus           pulumi.StringInput                            `pulumi:"offlineDataTransferStatus"`
	SyncActivity                        pulumi.StringInput                            `pulumi:"syncActivity"`
	TotalPersistentFilesNotSyncingCount pulumi.Float64Input                           `pulumi:"totalPersistentFilesNotSyncingCount"`
	UploadActivity                      ServerEndpointSyncActivityStatusResponseInput `pulumi:"uploadActivity"`
	UploadHealth                        pulumi.StringInput                            `pulumi:"uploadHealth"`
	UploadStatus                        ServerEndpointSyncSessionStatusResponseInput  `pulumi:"uploadStatus"`
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

func (o ServerEndpointSyncStatusResponsePtrOutput) DownloadActivity() ServerEndpointSyncActivityStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *ServerEndpointSyncActivityStatusResponse {
		if v == nil {
			return nil
		}
		return &v.DownloadActivity
	}).(ServerEndpointSyncActivityStatusResponsePtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) DownloadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DownloadHealth
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) DownloadStatus() ServerEndpointSyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *ServerEndpointSyncSessionStatusResponse {
		if v == nil {
			return nil
		}
		return &v.DownloadStatus
	}).(ServerEndpointSyncSessionStatusResponsePtrOutput)
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

func (o ServerEndpointSyncStatusResponsePtrOutput) UploadActivity() ServerEndpointSyncActivityStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *ServerEndpointSyncActivityStatusResponse {
		if v == nil {
			return nil
		}
		return &v.UploadActivity
	}).(ServerEndpointSyncActivityStatusResponsePtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) UploadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UploadHealth
	}).(pulumi.StringPtrOutput)
}

func (o ServerEndpointSyncStatusResponsePtrOutput) UploadStatus() ServerEndpointSyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointSyncStatusResponse) *ServerEndpointSyncSessionStatusResponse {
		if v == nil {
			return nil
		}
		return &v.UploadStatus
	}).(ServerEndpointSyncSessionStatusResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerEndpointCloudTieringStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointCloudTieringStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerEndpointFilesNotSyncingErrorResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointFilesNotSyncingErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallErrorResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointRecallStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncActivityStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncActivityStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncSessionStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncSessionStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponsePtrOutput{})
}
