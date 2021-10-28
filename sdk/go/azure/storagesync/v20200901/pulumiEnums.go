


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FeatureStatus string

const (
	FeatureStatusOn  = FeatureStatus("on")
	FeatureStatusOff = FeatureStatus("off")
)

func (FeatureStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*FeatureStatus)(nil)).Elem()
}

func (e FeatureStatus) ToFeatureStatusOutput() FeatureStatusOutput {
	return pulumi.ToOutput(e).(FeatureStatusOutput)
}

func (e FeatureStatus) ToFeatureStatusOutputWithContext(ctx context.Context) FeatureStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FeatureStatusOutput)
}

func (e FeatureStatus) ToFeatureStatusPtrOutput() FeatureStatusPtrOutput {
	return e.ToFeatureStatusPtrOutputWithContext(context.Background())
}

func (e FeatureStatus) ToFeatureStatusPtrOutputWithContext(ctx context.Context) FeatureStatusPtrOutput {
	return FeatureStatus(e).ToFeatureStatusOutputWithContext(ctx).ToFeatureStatusPtrOutputWithContext(ctx)
}

func (e FeatureStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FeatureStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FeatureStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FeatureStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FeatureStatusOutput struct{ *pulumi.OutputState }

func (FeatureStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FeatureStatus)(nil)).Elem()
}

func (o FeatureStatusOutput) ToFeatureStatusOutput() FeatureStatusOutput {
	return o
}

func (o FeatureStatusOutput) ToFeatureStatusOutputWithContext(ctx context.Context) FeatureStatusOutput {
	return o
}

func (o FeatureStatusOutput) ToFeatureStatusPtrOutput() FeatureStatusPtrOutput {
	return o.ToFeatureStatusPtrOutputWithContext(context.Background())
}

func (o FeatureStatusOutput) ToFeatureStatusPtrOutputWithContext(ctx context.Context) FeatureStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FeatureStatus) *FeatureStatus {
		return &v
	}).(FeatureStatusPtrOutput)
}

func (o FeatureStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FeatureStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FeatureStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FeatureStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FeatureStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FeatureStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FeatureStatusPtrOutput struct{ *pulumi.OutputState }

func (FeatureStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureStatus)(nil)).Elem()
}

func (o FeatureStatusPtrOutput) ToFeatureStatusPtrOutput() FeatureStatusPtrOutput {
	return o
}

func (o FeatureStatusPtrOutput) ToFeatureStatusPtrOutputWithContext(ctx context.Context) FeatureStatusPtrOutput {
	return o
}

func (o FeatureStatusPtrOutput) Elem() FeatureStatusOutput {
	return o.ApplyT(func(v *FeatureStatus) FeatureStatus {
		if v != nil {
			return *v
		}
		var ret FeatureStatus
		return ret
	}).(FeatureStatusOutput)
}

func (o FeatureStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FeatureStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FeatureStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FeatureStatusInput interface {
	pulumi.Input

	ToFeatureStatusOutput() FeatureStatusOutput
	ToFeatureStatusOutputWithContext(context.Context) FeatureStatusOutput
}

var featureStatusPtrType = reflect.TypeOf((**FeatureStatus)(nil)).Elem()

type FeatureStatusPtrInput interface {
	pulumi.Input

	ToFeatureStatusPtrOutput() FeatureStatusPtrOutput
	ToFeatureStatusPtrOutputWithContext(context.Context) FeatureStatusPtrOutput
}

type featureStatusPtr string

func FeatureStatusPtr(v string) FeatureStatusPtrInput {
	return (*featureStatusPtr)(&v)
}

func (*featureStatusPtr) ElementType() reflect.Type {
	return featureStatusPtrType
}

func (in *featureStatusPtr) ToFeatureStatusPtrOutput() FeatureStatusPtrOutput {
	return pulumi.ToOutput(in).(FeatureStatusPtrOutput)
}

func (in *featureStatusPtr) ToFeatureStatusPtrOutputWithContext(ctx context.Context) FeatureStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FeatureStatusPtrOutput)
}

type IncomingTrafficPolicy string

const (
	IncomingTrafficPolicyAllowAllTraffic          = IncomingTrafficPolicy("AllowAllTraffic")
	IncomingTrafficPolicyAllowVirtualNetworksOnly = IncomingTrafficPolicy("AllowVirtualNetworksOnly")
)

func (IncomingTrafficPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*IncomingTrafficPolicy)(nil)).Elem()
}

func (e IncomingTrafficPolicy) ToIncomingTrafficPolicyOutput() IncomingTrafficPolicyOutput {
	return pulumi.ToOutput(e).(IncomingTrafficPolicyOutput)
}

func (e IncomingTrafficPolicy) ToIncomingTrafficPolicyOutputWithContext(ctx context.Context) IncomingTrafficPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncomingTrafficPolicyOutput)
}

func (e IncomingTrafficPolicy) ToIncomingTrafficPolicyPtrOutput() IncomingTrafficPolicyPtrOutput {
	return e.ToIncomingTrafficPolicyPtrOutputWithContext(context.Background())
}

func (e IncomingTrafficPolicy) ToIncomingTrafficPolicyPtrOutputWithContext(ctx context.Context) IncomingTrafficPolicyPtrOutput {
	return IncomingTrafficPolicy(e).ToIncomingTrafficPolicyOutputWithContext(ctx).ToIncomingTrafficPolicyPtrOutputWithContext(ctx)
}

func (e IncomingTrafficPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncomingTrafficPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncomingTrafficPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncomingTrafficPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncomingTrafficPolicyOutput struct{ *pulumi.OutputState }

func (IncomingTrafficPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncomingTrafficPolicy)(nil)).Elem()
}

func (o IncomingTrafficPolicyOutput) ToIncomingTrafficPolicyOutput() IncomingTrafficPolicyOutput {
	return o
}

func (o IncomingTrafficPolicyOutput) ToIncomingTrafficPolicyOutputWithContext(ctx context.Context) IncomingTrafficPolicyOutput {
	return o
}

func (o IncomingTrafficPolicyOutput) ToIncomingTrafficPolicyPtrOutput() IncomingTrafficPolicyPtrOutput {
	return o.ToIncomingTrafficPolicyPtrOutputWithContext(context.Background())
}

func (o IncomingTrafficPolicyOutput) ToIncomingTrafficPolicyPtrOutputWithContext(ctx context.Context) IncomingTrafficPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncomingTrafficPolicy) *IncomingTrafficPolicy {
		return &v
	}).(IncomingTrafficPolicyPtrOutput)
}

func (o IncomingTrafficPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncomingTrafficPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncomingTrafficPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncomingTrafficPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncomingTrafficPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncomingTrafficPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncomingTrafficPolicyPtrOutput struct{ *pulumi.OutputState }

func (IncomingTrafficPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncomingTrafficPolicy)(nil)).Elem()
}

func (o IncomingTrafficPolicyPtrOutput) ToIncomingTrafficPolicyPtrOutput() IncomingTrafficPolicyPtrOutput {
	return o
}

func (o IncomingTrafficPolicyPtrOutput) ToIncomingTrafficPolicyPtrOutputWithContext(ctx context.Context) IncomingTrafficPolicyPtrOutput {
	return o
}

func (o IncomingTrafficPolicyPtrOutput) Elem() IncomingTrafficPolicyOutput {
	return o.ApplyT(func(v *IncomingTrafficPolicy) IncomingTrafficPolicy {
		if v != nil {
			return *v
		}
		var ret IncomingTrafficPolicy
		return ret
	}).(IncomingTrafficPolicyOutput)
}

func (o IncomingTrafficPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncomingTrafficPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncomingTrafficPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IncomingTrafficPolicyInput interface {
	pulumi.Input

	ToIncomingTrafficPolicyOutput() IncomingTrafficPolicyOutput
	ToIncomingTrafficPolicyOutputWithContext(context.Context) IncomingTrafficPolicyOutput
}

var incomingTrafficPolicyPtrType = reflect.TypeOf((**IncomingTrafficPolicy)(nil)).Elem()

type IncomingTrafficPolicyPtrInput interface {
	pulumi.Input

	ToIncomingTrafficPolicyPtrOutput() IncomingTrafficPolicyPtrOutput
	ToIncomingTrafficPolicyPtrOutputWithContext(context.Context) IncomingTrafficPolicyPtrOutput
}

type incomingTrafficPolicyPtr string

func IncomingTrafficPolicyPtr(v string) IncomingTrafficPolicyPtrInput {
	return (*incomingTrafficPolicyPtr)(&v)
}

func (*incomingTrafficPolicyPtr) ElementType() reflect.Type {
	return incomingTrafficPolicyPtrType
}

func (in *incomingTrafficPolicyPtr) ToIncomingTrafficPolicyPtrOutput() IncomingTrafficPolicyPtrOutput {
	return pulumi.ToOutput(in).(IncomingTrafficPolicyPtrOutput)
}

func (in *incomingTrafficPolicyPtr) ToIncomingTrafficPolicyPtrOutputWithContext(ctx context.Context) IncomingTrafficPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncomingTrafficPolicyPtrOutput)
}

type InitialDownloadPolicy string

const (
	InitialDownloadPolicyNamespaceOnly              = InitialDownloadPolicy("NamespaceOnly")
	InitialDownloadPolicyNamespaceThenModifiedFiles = InitialDownloadPolicy("NamespaceThenModifiedFiles")
	InitialDownloadPolicyAvoidTieredFiles           = InitialDownloadPolicy("AvoidTieredFiles")
)

func (InitialDownloadPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*InitialDownloadPolicy)(nil)).Elem()
}

func (e InitialDownloadPolicy) ToInitialDownloadPolicyOutput() InitialDownloadPolicyOutput {
	return pulumi.ToOutput(e).(InitialDownloadPolicyOutput)
}

func (e InitialDownloadPolicy) ToInitialDownloadPolicyOutputWithContext(ctx context.Context) InitialDownloadPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InitialDownloadPolicyOutput)
}

func (e InitialDownloadPolicy) ToInitialDownloadPolicyPtrOutput() InitialDownloadPolicyPtrOutput {
	return e.ToInitialDownloadPolicyPtrOutputWithContext(context.Background())
}

func (e InitialDownloadPolicy) ToInitialDownloadPolicyPtrOutputWithContext(ctx context.Context) InitialDownloadPolicyPtrOutput {
	return InitialDownloadPolicy(e).ToInitialDownloadPolicyOutputWithContext(ctx).ToInitialDownloadPolicyPtrOutputWithContext(ctx)
}

func (e InitialDownloadPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InitialDownloadPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InitialDownloadPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InitialDownloadPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InitialDownloadPolicyOutput struct{ *pulumi.OutputState }

func (InitialDownloadPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InitialDownloadPolicy)(nil)).Elem()
}

func (o InitialDownloadPolicyOutput) ToInitialDownloadPolicyOutput() InitialDownloadPolicyOutput {
	return o
}

func (o InitialDownloadPolicyOutput) ToInitialDownloadPolicyOutputWithContext(ctx context.Context) InitialDownloadPolicyOutput {
	return o
}

func (o InitialDownloadPolicyOutput) ToInitialDownloadPolicyPtrOutput() InitialDownloadPolicyPtrOutput {
	return o.ToInitialDownloadPolicyPtrOutputWithContext(context.Background())
}

func (o InitialDownloadPolicyOutput) ToInitialDownloadPolicyPtrOutputWithContext(ctx context.Context) InitialDownloadPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InitialDownloadPolicy) *InitialDownloadPolicy {
		return &v
	}).(InitialDownloadPolicyPtrOutput)
}

func (o InitialDownloadPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InitialDownloadPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InitialDownloadPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InitialDownloadPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InitialDownloadPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InitialDownloadPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InitialDownloadPolicyPtrOutput struct{ *pulumi.OutputState }

func (InitialDownloadPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InitialDownloadPolicy)(nil)).Elem()
}

func (o InitialDownloadPolicyPtrOutput) ToInitialDownloadPolicyPtrOutput() InitialDownloadPolicyPtrOutput {
	return o
}

func (o InitialDownloadPolicyPtrOutput) ToInitialDownloadPolicyPtrOutputWithContext(ctx context.Context) InitialDownloadPolicyPtrOutput {
	return o
}

func (o InitialDownloadPolicyPtrOutput) Elem() InitialDownloadPolicyOutput {
	return o.ApplyT(func(v *InitialDownloadPolicy) InitialDownloadPolicy {
		if v != nil {
			return *v
		}
		var ret InitialDownloadPolicy
		return ret
	}).(InitialDownloadPolicyOutput)
}

func (o InitialDownloadPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InitialDownloadPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InitialDownloadPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InitialDownloadPolicyInput interface {
	pulumi.Input

	ToInitialDownloadPolicyOutput() InitialDownloadPolicyOutput
	ToInitialDownloadPolicyOutputWithContext(context.Context) InitialDownloadPolicyOutput
}

var initialDownloadPolicyPtrType = reflect.TypeOf((**InitialDownloadPolicy)(nil)).Elem()

type InitialDownloadPolicyPtrInput interface {
	pulumi.Input

	ToInitialDownloadPolicyPtrOutput() InitialDownloadPolicyPtrOutput
	ToInitialDownloadPolicyPtrOutputWithContext(context.Context) InitialDownloadPolicyPtrOutput
}

type initialDownloadPolicyPtr string

func InitialDownloadPolicyPtr(v string) InitialDownloadPolicyPtrInput {
	return (*initialDownloadPolicyPtr)(&v)
}

func (*initialDownloadPolicyPtr) ElementType() reflect.Type {
	return initialDownloadPolicyPtrType
}

func (in *initialDownloadPolicyPtr) ToInitialDownloadPolicyPtrOutput() InitialDownloadPolicyPtrOutput {
	return pulumi.ToOutput(in).(InitialDownloadPolicyPtrOutput)
}

func (in *initialDownloadPolicyPtr) ToInitialDownloadPolicyPtrOutputWithContext(ctx context.Context) InitialDownloadPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InitialDownloadPolicyPtrOutput)
}

type InitialUploadPolicy string

const (
	InitialUploadPolicyServerAuthoritative = InitialUploadPolicy("ServerAuthoritative")
	InitialUploadPolicyMerge               = InitialUploadPolicy("Merge")
)

func (InitialUploadPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*InitialUploadPolicy)(nil)).Elem()
}

func (e InitialUploadPolicy) ToInitialUploadPolicyOutput() InitialUploadPolicyOutput {
	return pulumi.ToOutput(e).(InitialUploadPolicyOutput)
}

func (e InitialUploadPolicy) ToInitialUploadPolicyOutputWithContext(ctx context.Context) InitialUploadPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InitialUploadPolicyOutput)
}

func (e InitialUploadPolicy) ToInitialUploadPolicyPtrOutput() InitialUploadPolicyPtrOutput {
	return e.ToInitialUploadPolicyPtrOutputWithContext(context.Background())
}

func (e InitialUploadPolicy) ToInitialUploadPolicyPtrOutputWithContext(ctx context.Context) InitialUploadPolicyPtrOutput {
	return InitialUploadPolicy(e).ToInitialUploadPolicyOutputWithContext(ctx).ToInitialUploadPolicyPtrOutputWithContext(ctx)
}

func (e InitialUploadPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InitialUploadPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InitialUploadPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InitialUploadPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InitialUploadPolicyOutput struct{ *pulumi.OutputState }

func (InitialUploadPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InitialUploadPolicy)(nil)).Elem()
}

func (o InitialUploadPolicyOutput) ToInitialUploadPolicyOutput() InitialUploadPolicyOutput {
	return o
}

func (o InitialUploadPolicyOutput) ToInitialUploadPolicyOutputWithContext(ctx context.Context) InitialUploadPolicyOutput {
	return o
}

func (o InitialUploadPolicyOutput) ToInitialUploadPolicyPtrOutput() InitialUploadPolicyPtrOutput {
	return o.ToInitialUploadPolicyPtrOutputWithContext(context.Background())
}

func (o InitialUploadPolicyOutput) ToInitialUploadPolicyPtrOutputWithContext(ctx context.Context) InitialUploadPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InitialUploadPolicy) *InitialUploadPolicy {
		return &v
	}).(InitialUploadPolicyPtrOutput)
}

func (o InitialUploadPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InitialUploadPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InitialUploadPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InitialUploadPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InitialUploadPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InitialUploadPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InitialUploadPolicyPtrOutput struct{ *pulumi.OutputState }

func (InitialUploadPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InitialUploadPolicy)(nil)).Elem()
}

func (o InitialUploadPolicyPtrOutput) ToInitialUploadPolicyPtrOutput() InitialUploadPolicyPtrOutput {
	return o
}

func (o InitialUploadPolicyPtrOutput) ToInitialUploadPolicyPtrOutputWithContext(ctx context.Context) InitialUploadPolicyPtrOutput {
	return o
}

func (o InitialUploadPolicyPtrOutput) Elem() InitialUploadPolicyOutput {
	return o.ApplyT(func(v *InitialUploadPolicy) InitialUploadPolicy {
		if v != nil {
			return *v
		}
		var ret InitialUploadPolicy
		return ret
	}).(InitialUploadPolicyOutput)
}

func (o InitialUploadPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InitialUploadPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InitialUploadPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InitialUploadPolicyInput interface {
	pulumi.Input

	ToInitialUploadPolicyOutput() InitialUploadPolicyOutput
	ToInitialUploadPolicyOutputWithContext(context.Context) InitialUploadPolicyOutput
}

var initialUploadPolicyPtrType = reflect.TypeOf((**InitialUploadPolicy)(nil)).Elem()

type InitialUploadPolicyPtrInput interface {
	pulumi.Input

	ToInitialUploadPolicyPtrOutput() InitialUploadPolicyPtrOutput
	ToInitialUploadPolicyPtrOutputWithContext(context.Context) InitialUploadPolicyPtrOutput
}

type initialUploadPolicyPtr string

func InitialUploadPolicyPtr(v string) InitialUploadPolicyPtrInput {
	return (*initialUploadPolicyPtr)(&v)
}

func (*initialUploadPolicyPtr) ElementType() reflect.Type {
	return initialUploadPolicyPtrType
}

func (in *initialUploadPolicyPtr) ToInitialUploadPolicyPtrOutput() InitialUploadPolicyPtrOutput {
	return pulumi.ToOutput(in).(InitialUploadPolicyPtrOutput)
}

func (in *initialUploadPolicyPtr) ToInitialUploadPolicyPtrOutputWithContext(ctx context.Context) InitialUploadPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InitialUploadPolicyPtrOutput)
}

type LocalCacheMode string

const (
	LocalCacheModeDownloadNewAndModifiedFiles = LocalCacheMode("DownloadNewAndModifiedFiles")
	LocalCacheModeUpdateLocallyCachedFiles    = LocalCacheMode("UpdateLocallyCachedFiles")
)

func (LocalCacheMode) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalCacheMode)(nil)).Elem()
}

func (e LocalCacheMode) ToLocalCacheModeOutput() LocalCacheModeOutput {
	return pulumi.ToOutput(e).(LocalCacheModeOutput)
}

func (e LocalCacheMode) ToLocalCacheModeOutputWithContext(ctx context.Context) LocalCacheModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LocalCacheModeOutput)
}

func (e LocalCacheMode) ToLocalCacheModePtrOutput() LocalCacheModePtrOutput {
	return e.ToLocalCacheModePtrOutputWithContext(context.Background())
}

func (e LocalCacheMode) ToLocalCacheModePtrOutputWithContext(ctx context.Context) LocalCacheModePtrOutput {
	return LocalCacheMode(e).ToLocalCacheModeOutputWithContext(ctx).ToLocalCacheModePtrOutputWithContext(ctx)
}

func (e LocalCacheMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LocalCacheMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LocalCacheMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LocalCacheMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LocalCacheModeOutput struct{ *pulumi.OutputState }

func (LocalCacheModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalCacheMode)(nil)).Elem()
}

func (o LocalCacheModeOutput) ToLocalCacheModeOutput() LocalCacheModeOutput {
	return o
}

func (o LocalCacheModeOutput) ToLocalCacheModeOutputWithContext(ctx context.Context) LocalCacheModeOutput {
	return o
}

func (o LocalCacheModeOutput) ToLocalCacheModePtrOutput() LocalCacheModePtrOutput {
	return o.ToLocalCacheModePtrOutputWithContext(context.Background())
}

func (o LocalCacheModeOutput) ToLocalCacheModePtrOutputWithContext(ctx context.Context) LocalCacheModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalCacheMode) *LocalCacheMode {
		return &v
	}).(LocalCacheModePtrOutput)
}

func (o LocalCacheModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LocalCacheModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LocalCacheMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LocalCacheModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LocalCacheModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LocalCacheMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LocalCacheModePtrOutput struct{ *pulumi.OutputState }

func (LocalCacheModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalCacheMode)(nil)).Elem()
}

func (o LocalCacheModePtrOutput) ToLocalCacheModePtrOutput() LocalCacheModePtrOutput {
	return o
}

func (o LocalCacheModePtrOutput) ToLocalCacheModePtrOutputWithContext(ctx context.Context) LocalCacheModePtrOutput {
	return o
}

func (o LocalCacheModePtrOutput) Elem() LocalCacheModeOutput {
	return o.ApplyT(func(v *LocalCacheMode) LocalCacheMode {
		if v != nil {
			return *v
		}
		var ret LocalCacheMode
		return ret
	}).(LocalCacheModeOutput)
}

func (o LocalCacheModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LocalCacheModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LocalCacheMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LocalCacheModeInput interface {
	pulumi.Input

	ToLocalCacheModeOutput() LocalCacheModeOutput
	ToLocalCacheModeOutputWithContext(context.Context) LocalCacheModeOutput
}

var localCacheModePtrType = reflect.TypeOf((**LocalCacheMode)(nil)).Elem()

type LocalCacheModePtrInput interface {
	pulumi.Input

	ToLocalCacheModePtrOutput() LocalCacheModePtrOutput
	ToLocalCacheModePtrOutputWithContext(context.Context) LocalCacheModePtrOutput
}

type localCacheModePtr string

func LocalCacheModePtr(v string) LocalCacheModePtrInput {
	return (*localCacheModePtr)(&v)
}

func (*localCacheModePtr) ElementType() reflect.Type {
	return localCacheModePtrType
}

func (in *localCacheModePtr) ToLocalCacheModePtrOutput() LocalCacheModePtrOutput {
	return pulumi.ToOutput(in).(LocalCacheModePtrOutput)
}

func (in *localCacheModePtr) ToLocalCacheModePtrOutputWithContext(ctx context.Context) LocalCacheModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LocalCacheModePtrOutput)
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

func (PrivateEndpointServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return e.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return PrivateEndpointServiceConnectionStatus(e).ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx).ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointServiceConnectionStatus) *PrivateEndpointServiceConnectionStatus {
		return &v
	}).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) Elem() PrivateEndpointServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointServiceConnectionStatus) PrivateEndpointServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointServiceConnectionStatus
		return ret
	}).(PrivateEndpointServiceConnectionStatusOutput)
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput
	ToPrivateEndpointServiceConnectionStatusOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusOutput
}

var privateEndpointServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()

type PrivateEndpointServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput
	ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusPtrOutput
}

type privateEndpointServiceConnectionStatusPtr string

func PrivateEndpointServiceConnectionStatusPtr(v string) PrivateEndpointServiceConnectionStatusPtrInput {
	return (*privateEndpointServiceConnectionStatusPtr)(&v)
}

func (*privateEndpointServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointServiceConnectionStatusPtrType
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureStatusInput)(nil)).Elem(), FeatureStatus("on"))
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureStatusPtrInput)(nil)).Elem(), FeatureStatus("on"))
	pulumi.RegisterInputType(reflect.TypeOf((*IncomingTrafficPolicyInput)(nil)).Elem(), IncomingTrafficPolicy("AllowAllTraffic"))
	pulumi.RegisterInputType(reflect.TypeOf((*IncomingTrafficPolicyPtrInput)(nil)).Elem(), IncomingTrafficPolicy("AllowAllTraffic"))
	pulumi.RegisterInputType(reflect.TypeOf((*InitialDownloadPolicyInput)(nil)).Elem(), InitialDownloadPolicy("NamespaceOnly"))
	pulumi.RegisterInputType(reflect.TypeOf((*InitialDownloadPolicyPtrInput)(nil)).Elem(), InitialDownloadPolicy("NamespaceOnly"))
	pulumi.RegisterInputType(reflect.TypeOf((*InitialUploadPolicyInput)(nil)).Elem(), InitialUploadPolicy("ServerAuthoritative"))
	pulumi.RegisterInputType(reflect.TypeOf((*InitialUploadPolicyPtrInput)(nil)).Elem(), InitialUploadPolicy("ServerAuthoritative"))
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCacheModeInput)(nil)).Elem(), LocalCacheMode("DownloadNewAndModifiedFiles"))
	pulumi.RegisterInputType(reflect.TypeOf((*LocalCacheModePtrInput)(nil)).Elem(), LocalCacheMode("DownloadNewAndModifiedFiles"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterOutputType(FeatureStatusOutput{})
	pulumi.RegisterOutputType(FeatureStatusPtrOutput{})
	pulumi.RegisterOutputType(IncomingTrafficPolicyOutput{})
	pulumi.RegisterOutputType(IncomingTrafficPolicyPtrOutput{})
	pulumi.RegisterOutputType(InitialDownloadPolicyOutput{})
	pulumi.RegisterOutputType(InitialDownloadPolicyPtrOutput{})
	pulumi.RegisterOutputType(InitialUploadPolicyOutput{})
	pulumi.RegisterOutputType(InitialUploadPolicyPtrOutput{})
	pulumi.RegisterOutputType(LocalCacheModeOutput{})
	pulumi.RegisterOutputType(LocalCacheModePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
}
