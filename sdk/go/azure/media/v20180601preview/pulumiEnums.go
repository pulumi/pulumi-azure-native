


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AacAudioProfile string

const (
	// Specifies that the output audio is to be encoded into AAC Low Complexity profile (AAC-LC).
	AacAudioProfileAacLc = AacAudioProfile("AacLc")
	// Specifies that the output audio is to be encoded into HE-AAC v1 profile.
	AacAudioProfileHeAacV1 = AacAudioProfile("HeAacV1")
	// Specifies that the output audio is to be encoded into HE-AAC v2 profile.
	AacAudioProfileHeAacV2 = AacAudioProfile("HeAacV2")
)

func (AacAudioProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*AacAudioProfile)(nil)).Elem()
}

func (e AacAudioProfile) ToAacAudioProfileOutput() AacAudioProfileOutput {
	return pulumi.ToOutput(e).(AacAudioProfileOutput)
}

func (e AacAudioProfile) ToAacAudioProfileOutputWithContext(ctx context.Context) AacAudioProfileOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AacAudioProfileOutput)
}

func (e AacAudioProfile) ToAacAudioProfilePtrOutput() AacAudioProfilePtrOutput {
	return e.ToAacAudioProfilePtrOutputWithContext(context.Background())
}

func (e AacAudioProfile) ToAacAudioProfilePtrOutputWithContext(ctx context.Context) AacAudioProfilePtrOutput {
	return AacAudioProfile(e).ToAacAudioProfileOutputWithContext(ctx).ToAacAudioProfilePtrOutputWithContext(ctx)
}

func (e AacAudioProfile) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AacAudioProfile) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AacAudioProfile) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AacAudioProfile) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AacAudioProfileOutput struct{ *pulumi.OutputState }

func (AacAudioProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AacAudioProfile)(nil)).Elem()
}

func (o AacAudioProfileOutput) ToAacAudioProfileOutput() AacAudioProfileOutput {
	return o
}

func (o AacAudioProfileOutput) ToAacAudioProfileOutputWithContext(ctx context.Context) AacAudioProfileOutput {
	return o
}

func (o AacAudioProfileOutput) ToAacAudioProfilePtrOutput() AacAudioProfilePtrOutput {
	return o.ToAacAudioProfilePtrOutputWithContext(context.Background())
}

func (o AacAudioProfileOutput) ToAacAudioProfilePtrOutputWithContext(ctx context.Context) AacAudioProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AacAudioProfile) *AacAudioProfile {
		return &v
	}).(AacAudioProfilePtrOutput)
}

func (o AacAudioProfileOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AacAudioProfileOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AacAudioProfile) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AacAudioProfileOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AacAudioProfileOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AacAudioProfile) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AacAudioProfilePtrOutput struct{ *pulumi.OutputState }

func (AacAudioProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AacAudioProfile)(nil)).Elem()
}

func (o AacAudioProfilePtrOutput) ToAacAudioProfilePtrOutput() AacAudioProfilePtrOutput {
	return o
}

func (o AacAudioProfilePtrOutput) ToAacAudioProfilePtrOutputWithContext(ctx context.Context) AacAudioProfilePtrOutput {
	return o
}

func (o AacAudioProfilePtrOutput) Elem() AacAudioProfileOutput {
	return o.ApplyT(func(v *AacAudioProfile) AacAudioProfile {
		if v != nil {
			return *v
		}
		var ret AacAudioProfile
		return ret
	}).(AacAudioProfileOutput)
}

func (o AacAudioProfilePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AacAudioProfilePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AacAudioProfile) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AacAudioProfileInput interface {
	pulumi.Input

	ToAacAudioProfileOutput() AacAudioProfileOutput
	ToAacAudioProfileOutputWithContext(context.Context) AacAudioProfileOutput
}

var aacAudioProfilePtrType = reflect.TypeOf((**AacAudioProfile)(nil)).Elem()

type AacAudioProfilePtrInput interface {
	pulumi.Input

	ToAacAudioProfilePtrOutput() AacAudioProfilePtrOutput
	ToAacAudioProfilePtrOutputWithContext(context.Context) AacAudioProfilePtrOutput
}

type aacAudioProfilePtr string

func AacAudioProfilePtr(v string) AacAudioProfilePtrInput {
	return (*aacAudioProfilePtr)(&v)
}

func (*aacAudioProfilePtr) ElementType() reflect.Type {
	return aacAudioProfilePtrType
}

func (in *aacAudioProfilePtr) ToAacAudioProfilePtrOutput() AacAudioProfilePtrOutput {
	return pulumi.ToOutput(in).(AacAudioProfilePtrOutput)
}

func (in *aacAudioProfilePtr) ToAacAudioProfilePtrOutputWithContext(ctx context.Context) AacAudioProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AacAudioProfilePtrOutput)
}

type AssetContainerPermission string

const (
	// The SAS URL will allow read access to the container.
	AssetContainerPermissionRead = AssetContainerPermission("Read")
	// The SAS URL will allow read and write access to the container.
	AssetContainerPermissionReadWrite = AssetContainerPermission("ReadWrite")
	// The SAS URL will allow read, write and delete access to the container.
	AssetContainerPermissionReadWriteDelete = AssetContainerPermission("ReadWriteDelete")
)

func (AssetContainerPermission) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetContainerPermission)(nil)).Elem()
}

func (e AssetContainerPermission) ToAssetContainerPermissionOutput() AssetContainerPermissionOutput {
	return pulumi.ToOutput(e).(AssetContainerPermissionOutput)
}

func (e AssetContainerPermission) ToAssetContainerPermissionOutputWithContext(ctx context.Context) AssetContainerPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssetContainerPermissionOutput)
}

func (e AssetContainerPermission) ToAssetContainerPermissionPtrOutput() AssetContainerPermissionPtrOutput {
	return e.ToAssetContainerPermissionPtrOutputWithContext(context.Background())
}

func (e AssetContainerPermission) ToAssetContainerPermissionPtrOutputWithContext(ctx context.Context) AssetContainerPermissionPtrOutput {
	return AssetContainerPermission(e).ToAssetContainerPermissionOutputWithContext(ctx).ToAssetContainerPermissionPtrOutputWithContext(ctx)
}

func (e AssetContainerPermission) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetContainerPermission) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssetContainerPermission) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssetContainerPermission) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssetContainerPermissionOutput struct{ *pulumi.OutputState }

func (AssetContainerPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssetContainerPermission)(nil)).Elem()
}

func (o AssetContainerPermissionOutput) ToAssetContainerPermissionOutput() AssetContainerPermissionOutput {
	return o
}

func (o AssetContainerPermissionOutput) ToAssetContainerPermissionOutputWithContext(ctx context.Context) AssetContainerPermissionOutput {
	return o
}

func (o AssetContainerPermissionOutput) ToAssetContainerPermissionPtrOutput() AssetContainerPermissionPtrOutput {
	return o.ToAssetContainerPermissionPtrOutputWithContext(context.Background())
}

func (o AssetContainerPermissionOutput) ToAssetContainerPermissionPtrOutputWithContext(ctx context.Context) AssetContainerPermissionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssetContainerPermission) *AssetContainerPermission {
		return &v
	}).(AssetContainerPermissionPtrOutput)
}

func (o AssetContainerPermissionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssetContainerPermissionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetContainerPermission) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssetContainerPermissionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetContainerPermissionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssetContainerPermission) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssetContainerPermissionPtrOutput struct{ *pulumi.OutputState }

func (AssetContainerPermissionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetContainerPermission)(nil)).Elem()
}

func (o AssetContainerPermissionPtrOutput) ToAssetContainerPermissionPtrOutput() AssetContainerPermissionPtrOutput {
	return o
}

func (o AssetContainerPermissionPtrOutput) ToAssetContainerPermissionPtrOutputWithContext(ctx context.Context) AssetContainerPermissionPtrOutput {
	return o
}

func (o AssetContainerPermissionPtrOutput) Elem() AssetContainerPermissionOutput {
	return o.ApplyT(func(v *AssetContainerPermission) AssetContainerPermission {
		if v != nil {
			return *v
		}
		var ret AssetContainerPermission
		return ret
	}).(AssetContainerPermissionOutput)
}

func (o AssetContainerPermissionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssetContainerPermissionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssetContainerPermission) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AssetContainerPermissionInput interface {
	pulumi.Input

	ToAssetContainerPermissionOutput() AssetContainerPermissionOutput
	ToAssetContainerPermissionOutputWithContext(context.Context) AssetContainerPermissionOutput
}

var assetContainerPermissionPtrType = reflect.TypeOf((**AssetContainerPermission)(nil)).Elem()

type AssetContainerPermissionPtrInput interface {
	pulumi.Input

	ToAssetContainerPermissionPtrOutput() AssetContainerPermissionPtrOutput
	ToAssetContainerPermissionPtrOutputWithContext(context.Context) AssetContainerPermissionPtrOutput
}

type assetContainerPermissionPtr string

func AssetContainerPermissionPtr(v string) AssetContainerPermissionPtrInput {
	return (*assetContainerPermissionPtr)(&v)
}

func (*assetContainerPermissionPtr) ElementType() reflect.Type {
	return assetContainerPermissionPtrType
}

func (in *assetContainerPermissionPtr) ToAssetContainerPermissionPtrOutput() AssetContainerPermissionPtrOutput {
	return pulumi.ToOutput(in).(AssetContainerPermissionPtrOutput)
}

func (in *assetContainerPermissionPtr) ToAssetContainerPermissionPtrOutputWithContext(ctx context.Context) AssetContainerPermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssetContainerPermissionPtrOutput)
}

type ContentKeyPolicyFairPlayRentalAndLeaseKeyType string

const (
	// Represents a ContentKeyPolicyFairPlayRentalAndLeaseKeyType that is unavailable in current API version.
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUnknown = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("Unknown")
	// Key duration is not specified.
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUndefined = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("Undefined")
	// Content key can be persisted with an unlimited duration
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentUnlimited = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("PersistentUnlimited")
	// Content key can be persisted and the valid duration is limited by the Rental Duration value
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentLimited = ContentKeyPolicyFairPlayRentalAndLeaseKeyType("PersistentLimited")
)

func (ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayRentalAndLeaseKeyType)(nil)).Elem()
}

func (e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput {
	return pulumi.ToOutput(e).(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput)
}

func (e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput)
}

func (e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput {
	return e.ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput {
	return ContentKeyPolicyFairPlayRentalAndLeaseKeyType(e).ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutputWithContext(ctx).ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutputWithContext(ctx)
}

func (e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyFairPlayRentalAndLeaseKeyType)(nil)).Elem()
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput {
	return o
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput {
	return o
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput {
	return o.ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyFairPlayRentalAndLeaseKeyType) *ContentKeyPolicyFairPlayRentalAndLeaseKeyType {
		return &v
	}).(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput)
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyFairPlayRentalAndLeaseKeyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyFairPlayRentalAndLeaseKeyType)(nil)).Elem()
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput {
	return o
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput {
	return o
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput) Elem() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput {
	return o.ApplyT(func(v *ContentKeyPolicyFairPlayRentalAndLeaseKeyType) ContentKeyPolicyFairPlayRentalAndLeaseKeyType {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyFairPlayRentalAndLeaseKeyType
		return ret
	}).(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput)
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentKeyPolicyFairPlayRentalAndLeaseKeyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeInput interface {
	pulumi.Input

	ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput
	ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutputWithContext(context.Context) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput
}

var contentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrType = reflect.TypeOf((**ContentKeyPolicyFairPlayRentalAndLeaseKeyType)(nil)).Elem()

type ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrInput interface {
	pulumi.Input

	ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput
	ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutputWithContext(context.Context) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput
}

type contentKeyPolicyFairPlayRentalAndLeaseKeyTypePtr string

func ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtr(v string) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrInput {
	return (*contentKeyPolicyFairPlayRentalAndLeaseKeyTypePtr)(&v)
}

func (*contentKeyPolicyFairPlayRentalAndLeaseKeyTypePtr) ElementType() reflect.Type {
	return contentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrType
}

func (in *contentKeyPolicyFairPlayRentalAndLeaseKeyTypePtr) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput() ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput {
	return pulumi.ToOutput(in).(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput)
}

func (in *contentKeyPolicyFairPlayRentalAndLeaseKeyTypePtr) ToContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput)
}

type ContentKeyPolicyPlayReadyContentType string

const (
	// Represents a ContentKeyPolicyPlayReadyContentType that is unavailable in current API version.
	ContentKeyPolicyPlayReadyContentTypeUnknown = ContentKeyPolicyPlayReadyContentType("Unknown")
	// Unspecified content type.
	ContentKeyPolicyPlayReadyContentTypeUnspecified = ContentKeyPolicyPlayReadyContentType("Unspecified")
	// Ultraviolet download content type.
	ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload = ContentKeyPolicyPlayReadyContentType("UltraVioletDownload")
	// Ultraviolet streaming content type.
	ContentKeyPolicyPlayReadyContentTypeUltraVioletStreaming = ContentKeyPolicyPlayReadyContentType("UltraVioletStreaming")
)

func (ContentKeyPolicyPlayReadyContentType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentType)(nil)).Elem()
}

func (e ContentKeyPolicyPlayReadyContentType) ToContentKeyPolicyPlayReadyContentTypeOutput() ContentKeyPolicyPlayReadyContentTypeOutput {
	return pulumi.ToOutput(e).(ContentKeyPolicyPlayReadyContentTypeOutput)
}

func (e ContentKeyPolicyPlayReadyContentType) ToContentKeyPolicyPlayReadyContentTypeOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentKeyPolicyPlayReadyContentTypeOutput)
}

func (e ContentKeyPolicyPlayReadyContentType) ToContentKeyPolicyPlayReadyContentTypePtrOutput() ContentKeyPolicyPlayReadyContentTypePtrOutput {
	return e.ToContentKeyPolicyPlayReadyContentTypePtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyPlayReadyContentType) ToContentKeyPolicyPlayReadyContentTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentTypePtrOutput {
	return ContentKeyPolicyPlayReadyContentType(e).ToContentKeyPolicyPlayReadyContentTypeOutputWithContext(ctx).ToContentKeyPolicyPlayReadyContentTypePtrOutputWithContext(ctx)
}

func (e ContentKeyPolicyPlayReadyContentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyPlayReadyContentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyPlayReadyContentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyPlayReadyContentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentKeyPolicyPlayReadyContentTypeOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyContentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyContentType)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyContentTypeOutput) ToContentKeyPolicyPlayReadyContentTypeOutput() ContentKeyPolicyPlayReadyContentTypeOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentTypeOutput) ToContentKeyPolicyPlayReadyContentTypeOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentTypeOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentTypeOutput) ToContentKeyPolicyPlayReadyContentTypePtrOutput() ContentKeyPolicyPlayReadyContentTypePtrOutput {
	return o.ToContentKeyPolicyPlayReadyContentTypePtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyContentTypeOutput) ToContentKeyPolicyPlayReadyContentTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyPlayReadyContentType) *ContentKeyPolicyPlayReadyContentType {
		return &v
	}).(ContentKeyPolicyPlayReadyContentTypePtrOutput)
}

func (o ContentKeyPolicyPlayReadyContentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyContentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyPlayReadyContentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyContentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyContentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyPlayReadyContentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyPlayReadyContentTypePtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyContentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyContentType)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyContentTypePtrOutput) ToContentKeyPolicyPlayReadyContentTypePtrOutput() ContentKeyPolicyPlayReadyContentTypePtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentTypePtrOutput) ToContentKeyPolicyPlayReadyContentTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentTypePtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyContentTypePtrOutput) Elem() ContentKeyPolicyPlayReadyContentTypeOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyContentType) ContentKeyPolicyPlayReadyContentType {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyPlayReadyContentType
		return ret
	}).(ContentKeyPolicyPlayReadyContentTypeOutput)
}

func (o ContentKeyPolicyPlayReadyContentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyContentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentKeyPolicyPlayReadyContentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentKeyPolicyPlayReadyContentTypeInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyContentTypeOutput() ContentKeyPolicyPlayReadyContentTypeOutput
	ToContentKeyPolicyPlayReadyContentTypeOutputWithContext(context.Context) ContentKeyPolicyPlayReadyContentTypeOutput
}

var contentKeyPolicyPlayReadyContentTypePtrType = reflect.TypeOf((**ContentKeyPolicyPlayReadyContentType)(nil)).Elem()

type ContentKeyPolicyPlayReadyContentTypePtrInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyContentTypePtrOutput() ContentKeyPolicyPlayReadyContentTypePtrOutput
	ToContentKeyPolicyPlayReadyContentTypePtrOutputWithContext(context.Context) ContentKeyPolicyPlayReadyContentTypePtrOutput
}

type contentKeyPolicyPlayReadyContentTypePtr string

func ContentKeyPolicyPlayReadyContentTypePtr(v string) ContentKeyPolicyPlayReadyContentTypePtrInput {
	return (*contentKeyPolicyPlayReadyContentTypePtr)(&v)
}

func (*contentKeyPolicyPlayReadyContentTypePtr) ElementType() reflect.Type {
	return contentKeyPolicyPlayReadyContentTypePtrType
}

func (in *contentKeyPolicyPlayReadyContentTypePtr) ToContentKeyPolicyPlayReadyContentTypePtrOutput() ContentKeyPolicyPlayReadyContentTypePtrOutput {
	return pulumi.ToOutput(in).(ContentKeyPolicyPlayReadyContentTypePtrOutput)
}

func (in *contentKeyPolicyPlayReadyContentTypePtr) ToContentKeyPolicyPlayReadyContentTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyContentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentKeyPolicyPlayReadyContentTypePtrOutput)
}

type ContentKeyPolicyPlayReadyLicenseType string

const (
	// Represents a ContentKeyPolicyPlayReadyLicenseType that is unavailable in current API version.
	ContentKeyPolicyPlayReadyLicenseTypeUnknown = ContentKeyPolicyPlayReadyLicenseType("Unknown")
	// Non persistent license.
	ContentKeyPolicyPlayReadyLicenseTypeNonPersistent = ContentKeyPolicyPlayReadyLicenseType("NonPersistent")
	// Persistent license. Allows offline playback.
	ContentKeyPolicyPlayReadyLicenseTypePersistent = ContentKeyPolicyPlayReadyLicenseType("Persistent")
)

func (ContentKeyPolicyPlayReadyLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseType)(nil)).Elem()
}

func (e ContentKeyPolicyPlayReadyLicenseType) ToContentKeyPolicyPlayReadyLicenseTypeOutput() ContentKeyPolicyPlayReadyLicenseTypeOutput {
	return pulumi.ToOutput(e).(ContentKeyPolicyPlayReadyLicenseTypeOutput)
}

func (e ContentKeyPolicyPlayReadyLicenseType) ToContentKeyPolicyPlayReadyLicenseTypeOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentKeyPolicyPlayReadyLicenseTypeOutput)
}

func (e ContentKeyPolicyPlayReadyLicenseType) ToContentKeyPolicyPlayReadyLicenseTypePtrOutput() ContentKeyPolicyPlayReadyLicenseTypePtrOutput {
	return e.ToContentKeyPolicyPlayReadyLicenseTypePtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyPlayReadyLicenseType) ToContentKeyPolicyPlayReadyLicenseTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseTypePtrOutput {
	return ContentKeyPolicyPlayReadyLicenseType(e).ToContentKeyPolicyPlayReadyLicenseTypeOutputWithContext(ctx).ToContentKeyPolicyPlayReadyLicenseTypePtrOutputWithContext(ctx)
}

func (e ContentKeyPolicyPlayReadyLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyPlayReadyLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyPlayReadyLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyPlayReadyLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentKeyPolicyPlayReadyLicenseTypeOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseType)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyLicenseTypeOutput) ToContentKeyPolicyPlayReadyLicenseTypeOutput() ContentKeyPolicyPlayReadyLicenseTypeOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseTypeOutput) ToContentKeyPolicyPlayReadyLicenseTypeOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseTypeOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseTypeOutput) ToContentKeyPolicyPlayReadyLicenseTypePtrOutput() ContentKeyPolicyPlayReadyLicenseTypePtrOutput {
	return o.ToContentKeyPolicyPlayReadyLicenseTypePtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyLicenseTypeOutput) ToContentKeyPolicyPlayReadyLicenseTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyPlayReadyLicenseType) *ContentKeyPolicyPlayReadyLicenseType {
		return &v
	}).(ContentKeyPolicyPlayReadyLicenseTypePtrOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyPlayReadyLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyPlayReadyLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyPlayReadyLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyLicenseType)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyLicenseTypePtrOutput) ToContentKeyPolicyPlayReadyLicenseTypePtrOutput() ContentKeyPolicyPlayReadyLicenseTypePtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseTypePtrOutput) ToContentKeyPolicyPlayReadyLicenseTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseTypePtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyLicenseTypePtrOutput) Elem() ContentKeyPolicyPlayReadyLicenseTypeOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyLicenseType) ContentKeyPolicyPlayReadyLicenseType {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyPlayReadyLicenseType
		return ret
	}).(ContentKeyPolicyPlayReadyLicenseTypeOutput)
}

func (o ContentKeyPolicyPlayReadyLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentKeyPolicyPlayReadyLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentKeyPolicyPlayReadyLicenseTypeInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyLicenseTypeOutput() ContentKeyPolicyPlayReadyLicenseTypeOutput
	ToContentKeyPolicyPlayReadyLicenseTypeOutputWithContext(context.Context) ContentKeyPolicyPlayReadyLicenseTypeOutput
}

var contentKeyPolicyPlayReadyLicenseTypePtrType = reflect.TypeOf((**ContentKeyPolicyPlayReadyLicenseType)(nil)).Elem()

type ContentKeyPolicyPlayReadyLicenseTypePtrInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyLicenseTypePtrOutput() ContentKeyPolicyPlayReadyLicenseTypePtrOutput
	ToContentKeyPolicyPlayReadyLicenseTypePtrOutputWithContext(context.Context) ContentKeyPolicyPlayReadyLicenseTypePtrOutput
}

type contentKeyPolicyPlayReadyLicenseTypePtr string

func ContentKeyPolicyPlayReadyLicenseTypePtr(v string) ContentKeyPolicyPlayReadyLicenseTypePtrInput {
	return (*contentKeyPolicyPlayReadyLicenseTypePtr)(&v)
}

func (*contentKeyPolicyPlayReadyLicenseTypePtr) ElementType() reflect.Type {
	return contentKeyPolicyPlayReadyLicenseTypePtrType
}

func (in *contentKeyPolicyPlayReadyLicenseTypePtr) ToContentKeyPolicyPlayReadyLicenseTypePtrOutput() ContentKeyPolicyPlayReadyLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(ContentKeyPolicyPlayReadyLicenseTypePtrOutput)
}

func (in *contentKeyPolicyPlayReadyLicenseTypePtr) ToContentKeyPolicyPlayReadyLicenseTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentKeyPolicyPlayReadyLicenseTypePtrOutput)
}

type ContentKeyPolicyPlayReadyUnknownOutputPassingOption string

const (
	// Represents a ContentKeyPolicyPlayReadyUnknownOutputPassingOption that is unavailable in current API version.
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionUnknown = ContentKeyPolicyPlayReadyUnknownOutputPassingOption("Unknown")
	// Passing the video portion of protected content to an Unknown Output is not allowed.
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed = ContentKeyPolicyPlayReadyUnknownOutputPassingOption("NotAllowed")
	// Passing the video portion of protected content to an Unknown Output is allowed.
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowed = ContentKeyPolicyPlayReadyUnknownOutputPassingOption("Allowed")
	// Passing the video portion of protected content to an Unknown Output is allowed but with constrained resolution.
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowedWithVideoConstriction = ContentKeyPolicyPlayReadyUnknownOutputPassingOption("AllowedWithVideoConstriction")
)

func (ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyUnknownOutputPassingOption)(nil)).Elem()
}

func (e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput {
	return pulumi.ToOutput(e).(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput)
}

func (e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput)
}

func (e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput {
	return e.ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput {
	return ContentKeyPolicyPlayReadyUnknownOutputPassingOption(e).ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutputWithContext(ctx).ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutputWithContext(ctx)
}

func (e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyPlayReadyUnknownOutputPassingOption)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput {
	return o.ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyPlayReadyUnknownOutputPassingOption) *ContentKeyPolicyPlayReadyUnknownOutputPassingOption {
		return &v
	}).(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput)
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyPlayReadyUnknownOutputPassingOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyPlayReadyUnknownOutputPassingOption)(nil)).Elem()
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput {
	return o
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput) Elem() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput {
	return o.ApplyT(func(v *ContentKeyPolicyPlayReadyUnknownOutputPassingOption) ContentKeyPolicyPlayReadyUnknownOutputPassingOption {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyPlayReadyUnknownOutputPassingOption
		return ret
	}).(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput)
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentKeyPolicyPlayReadyUnknownOutputPassingOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentKeyPolicyPlayReadyUnknownOutputPassingOptionInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput
	ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutputWithContext(context.Context) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput
}

var contentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrType = reflect.TypeOf((**ContentKeyPolicyPlayReadyUnknownOutputPassingOption)(nil)).Elem()

type ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrInput interface {
	pulumi.Input

	ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput
	ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutputWithContext(context.Context) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput
}

type contentKeyPolicyPlayReadyUnknownOutputPassingOptionPtr string

func ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtr(v string) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrInput {
	return (*contentKeyPolicyPlayReadyUnknownOutputPassingOptionPtr)(&v)
}

func (*contentKeyPolicyPlayReadyUnknownOutputPassingOptionPtr) ElementType() reflect.Type {
	return contentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrType
}

func (in *contentKeyPolicyPlayReadyUnknownOutputPassingOptionPtr) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput() ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput {
	return pulumi.ToOutput(in).(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput)
}

func (in *contentKeyPolicyPlayReadyUnknownOutputPassingOptionPtr) ToContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutputWithContext(ctx context.Context) ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput)
}

type ContentKeyPolicyRestrictionTokenType string

const (
	// Represents a ContentKeyPolicyRestrictionTokenType that is unavailable in current API version.
	ContentKeyPolicyRestrictionTokenTypeUnknown = ContentKeyPolicyRestrictionTokenType("Unknown")
	// Simple Web Token.
	ContentKeyPolicyRestrictionTokenTypeSwt = ContentKeyPolicyRestrictionTokenType("Swt")
	// JSON Web Token.
	ContentKeyPolicyRestrictionTokenTypeJwt = ContentKeyPolicyRestrictionTokenType("Jwt")
)

func (ContentKeyPolicyRestrictionTokenType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyRestrictionTokenType)(nil)).Elem()
}

func (e ContentKeyPolicyRestrictionTokenType) ToContentKeyPolicyRestrictionTokenTypeOutput() ContentKeyPolicyRestrictionTokenTypeOutput {
	return pulumi.ToOutput(e).(ContentKeyPolicyRestrictionTokenTypeOutput)
}

func (e ContentKeyPolicyRestrictionTokenType) ToContentKeyPolicyRestrictionTokenTypeOutputWithContext(ctx context.Context) ContentKeyPolicyRestrictionTokenTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentKeyPolicyRestrictionTokenTypeOutput)
}

func (e ContentKeyPolicyRestrictionTokenType) ToContentKeyPolicyRestrictionTokenTypePtrOutput() ContentKeyPolicyRestrictionTokenTypePtrOutput {
	return e.ToContentKeyPolicyRestrictionTokenTypePtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyRestrictionTokenType) ToContentKeyPolicyRestrictionTokenTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyRestrictionTokenTypePtrOutput {
	return ContentKeyPolicyRestrictionTokenType(e).ToContentKeyPolicyRestrictionTokenTypeOutputWithContext(ctx).ToContentKeyPolicyRestrictionTokenTypePtrOutputWithContext(ctx)
}

func (e ContentKeyPolicyRestrictionTokenType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyRestrictionTokenType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentKeyPolicyRestrictionTokenType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentKeyPolicyRestrictionTokenType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentKeyPolicyRestrictionTokenTypeOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyRestrictionTokenTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyRestrictionTokenType)(nil)).Elem()
}

func (o ContentKeyPolicyRestrictionTokenTypeOutput) ToContentKeyPolicyRestrictionTokenTypeOutput() ContentKeyPolicyRestrictionTokenTypeOutput {
	return o
}

func (o ContentKeyPolicyRestrictionTokenTypeOutput) ToContentKeyPolicyRestrictionTokenTypeOutputWithContext(ctx context.Context) ContentKeyPolicyRestrictionTokenTypeOutput {
	return o
}

func (o ContentKeyPolicyRestrictionTokenTypeOutput) ToContentKeyPolicyRestrictionTokenTypePtrOutput() ContentKeyPolicyRestrictionTokenTypePtrOutput {
	return o.ToContentKeyPolicyRestrictionTokenTypePtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyRestrictionTokenTypeOutput) ToContentKeyPolicyRestrictionTokenTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyRestrictionTokenTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentKeyPolicyRestrictionTokenType) *ContentKeyPolicyRestrictionTokenType {
		return &v
	}).(ContentKeyPolicyRestrictionTokenTypePtrOutput)
}

func (o ContentKeyPolicyRestrictionTokenTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentKeyPolicyRestrictionTokenTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyRestrictionTokenType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentKeyPolicyRestrictionTokenTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyRestrictionTokenTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentKeyPolicyRestrictionTokenType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentKeyPolicyRestrictionTokenTypePtrOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyRestrictionTokenTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicyRestrictionTokenType)(nil)).Elem()
}

func (o ContentKeyPolicyRestrictionTokenTypePtrOutput) ToContentKeyPolicyRestrictionTokenTypePtrOutput() ContentKeyPolicyRestrictionTokenTypePtrOutput {
	return o
}

func (o ContentKeyPolicyRestrictionTokenTypePtrOutput) ToContentKeyPolicyRestrictionTokenTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyRestrictionTokenTypePtrOutput {
	return o
}

func (o ContentKeyPolicyRestrictionTokenTypePtrOutput) Elem() ContentKeyPolicyRestrictionTokenTypeOutput {
	return o.ApplyT(func(v *ContentKeyPolicyRestrictionTokenType) ContentKeyPolicyRestrictionTokenType {
		if v != nil {
			return *v
		}
		var ret ContentKeyPolicyRestrictionTokenType
		return ret
	}).(ContentKeyPolicyRestrictionTokenTypeOutput)
}

func (o ContentKeyPolicyRestrictionTokenTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentKeyPolicyRestrictionTokenTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentKeyPolicyRestrictionTokenType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentKeyPolicyRestrictionTokenTypeInput interface {
	pulumi.Input

	ToContentKeyPolicyRestrictionTokenTypeOutput() ContentKeyPolicyRestrictionTokenTypeOutput
	ToContentKeyPolicyRestrictionTokenTypeOutputWithContext(context.Context) ContentKeyPolicyRestrictionTokenTypeOutput
}

var contentKeyPolicyRestrictionTokenTypePtrType = reflect.TypeOf((**ContentKeyPolicyRestrictionTokenType)(nil)).Elem()

type ContentKeyPolicyRestrictionTokenTypePtrInput interface {
	pulumi.Input

	ToContentKeyPolicyRestrictionTokenTypePtrOutput() ContentKeyPolicyRestrictionTokenTypePtrOutput
	ToContentKeyPolicyRestrictionTokenTypePtrOutputWithContext(context.Context) ContentKeyPolicyRestrictionTokenTypePtrOutput
}

type contentKeyPolicyRestrictionTokenTypePtr string

func ContentKeyPolicyRestrictionTokenTypePtr(v string) ContentKeyPolicyRestrictionTokenTypePtrInput {
	return (*contentKeyPolicyRestrictionTokenTypePtr)(&v)
}

func (*contentKeyPolicyRestrictionTokenTypePtr) ElementType() reflect.Type {
	return contentKeyPolicyRestrictionTokenTypePtrType
}

func (in *contentKeyPolicyRestrictionTokenTypePtr) ToContentKeyPolicyRestrictionTokenTypePtrOutput() ContentKeyPolicyRestrictionTokenTypePtrOutput {
	return pulumi.ToOutput(in).(ContentKeyPolicyRestrictionTokenTypePtrOutput)
}

func (in *contentKeyPolicyRestrictionTokenTypePtr) ToContentKeyPolicyRestrictionTokenTypePtrOutputWithContext(ctx context.Context) ContentKeyPolicyRestrictionTokenTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentKeyPolicyRestrictionTokenTypePtrOutput)
}

type DeinterlaceMode string

const (
	// Disables de-interlacing of the source video.
	DeinterlaceModeOff = DeinterlaceMode("Off")
	// Apply automatic pixel adaptive de-interlacing on each frame in the input video.
	DeinterlaceModeAutoPixelAdaptive = DeinterlaceMode("AutoPixelAdaptive")
)

func (DeinterlaceMode) ElementType() reflect.Type {
	return reflect.TypeOf((*DeinterlaceMode)(nil)).Elem()
}

func (e DeinterlaceMode) ToDeinterlaceModeOutput() DeinterlaceModeOutput {
	return pulumi.ToOutput(e).(DeinterlaceModeOutput)
}

func (e DeinterlaceMode) ToDeinterlaceModeOutputWithContext(ctx context.Context) DeinterlaceModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeinterlaceModeOutput)
}

func (e DeinterlaceMode) ToDeinterlaceModePtrOutput() DeinterlaceModePtrOutput {
	return e.ToDeinterlaceModePtrOutputWithContext(context.Background())
}

func (e DeinterlaceMode) ToDeinterlaceModePtrOutputWithContext(ctx context.Context) DeinterlaceModePtrOutput {
	return DeinterlaceMode(e).ToDeinterlaceModeOutputWithContext(ctx).ToDeinterlaceModePtrOutputWithContext(ctx)
}

func (e DeinterlaceMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeinterlaceMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeinterlaceMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeinterlaceMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeinterlaceModeOutput struct{ *pulumi.OutputState }

func (DeinterlaceModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeinterlaceMode)(nil)).Elem()
}

func (o DeinterlaceModeOutput) ToDeinterlaceModeOutput() DeinterlaceModeOutput {
	return o
}

func (o DeinterlaceModeOutput) ToDeinterlaceModeOutputWithContext(ctx context.Context) DeinterlaceModeOutput {
	return o
}

func (o DeinterlaceModeOutput) ToDeinterlaceModePtrOutput() DeinterlaceModePtrOutput {
	return o.ToDeinterlaceModePtrOutputWithContext(context.Background())
}

func (o DeinterlaceModeOutput) ToDeinterlaceModePtrOutputWithContext(ctx context.Context) DeinterlaceModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeinterlaceMode) *DeinterlaceMode {
		return &v
	}).(DeinterlaceModePtrOutput)
}

func (o DeinterlaceModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeinterlaceModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeinterlaceMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeinterlaceModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeinterlaceModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeinterlaceMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeinterlaceModePtrOutput struct{ *pulumi.OutputState }

func (DeinterlaceModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeinterlaceMode)(nil)).Elem()
}

func (o DeinterlaceModePtrOutput) ToDeinterlaceModePtrOutput() DeinterlaceModePtrOutput {
	return o
}

func (o DeinterlaceModePtrOutput) ToDeinterlaceModePtrOutputWithContext(ctx context.Context) DeinterlaceModePtrOutput {
	return o
}

func (o DeinterlaceModePtrOutput) Elem() DeinterlaceModeOutput {
	return o.ApplyT(func(v *DeinterlaceMode) DeinterlaceMode {
		if v != nil {
			return *v
		}
		var ret DeinterlaceMode
		return ret
	}).(DeinterlaceModeOutput)
}

func (o DeinterlaceModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeinterlaceModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeinterlaceMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeinterlaceModeInput interface {
	pulumi.Input

	ToDeinterlaceModeOutput() DeinterlaceModeOutput
	ToDeinterlaceModeOutputWithContext(context.Context) DeinterlaceModeOutput
}

var deinterlaceModePtrType = reflect.TypeOf((**DeinterlaceMode)(nil)).Elem()

type DeinterlaceModePtrInput interface {
	pulumi.Input

	ToDeinterlaceModePtrOutput() DeinterlaceModePtrOutput
	ToDeinterlaceModePtrOutputWithContext(context.Context) DeinterlaceModePtrOutput
}

type deinterlaceModePtr string

func DeinterlaceModePtr(v string) DeinterlaceModePtrInput {
	return (*deinterlaceModePtr)(&v)
}

func (*deinterlaceModePtr) ElementType() reflect.Type {
	return deinterlaceModePtrType
}

func (in *deinterlaceModePtr) ToDeinterlaceModePtrOutput() DeinterlaceModePtrOutput {
	return pulumi.ToOutput(in).(DeinterlaceModePtrOutput)
}

func (in *deinterlaceModePtr) ToDeinterlaceModePtrOutputWithContext(ctx context.Context) DeinterlaceModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeinterlaceModePtrOutput)
}

type DeinterlaceParity string

const (
	// Automatically detect the order of fields
	DeinterlaceParityAuto = DeinterlaceParity("Auto")
	// Apply top field first processing of input video.
	DeinterlaceParityTopFieldFirst = DeinterlaceParity("TopFieldFirst")
	// Apply bottom field first processing of input video.
	DeinterlaceParityBottomFieldFirst = DeinterlaceParity("BottomFieldFirst")
)

func (DeinterlaceParity) ElementType() reflect.Type {
	return reflect.TypeOf((*DeinterlaceParity)(nil)).Elem()
}

func (e DeinterlaceParity) ToDeinterlaceParityOutput() DeinterlaceParityOutput {
	return pulumi.ToOutput(e).(DeinterlaceParityOutput)
}

func (e DeinterlaceParity) ToDeinterlaceParityOutputWithContext(ctx context.Context) DeinterlaceParityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeinterlaceParityOutput)
}

func (e DeinterlaceParity) ToDeinterlaceParityPtrOutput() DeinterlaceParityPtrOutput {
	return e.ToDeinterlaceParityPtrOutputWithContext(context.Background())
}

func (e DeinterlaceParity) ToDeinterlaceParityPtrOutputWithContext(ctx context.Context) DeinterlaceParityPtrOutput {
	return DeinterlaceParity(e).ToDeinterlaceParityOutputWithContext(ctx).ToDeinterlaceParityPtrOutputWithContext(ctx)
}

func (e DeinterlaceParity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeinterlaceParity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeinterlaceParity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeinterlaceParity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeinterlaceParityOutput struct{ *pulumi.OutputState }

func (DeinterlaceParityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeinterlaceParity)(nil)).Elem()
}

func (o DeinterlaceParityOutput) ToDeinterlaceParityOutput() DeinterlaceParityOutput {
	return o
}

func (o DeinterlaceParityOutput) ToDeinterlaceParityOutputWithContext(ctx context.Context) DeinterlaceParityOutput {
	return o
}

func (o DeinterlaceParityOutput) ToDeinterlaceParityPtrOutput() DeinterlaceParityPtrOutput {
	return o.ToDeinterlaceParityPtrOutputWithContext(context.Background())
}

func (o DeinterlaceParityOutput) ToDeinterlaceParityPtrOutputWithContext(ctx context.Context) DeinterlaceParityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeinterlaceParity) *DeinterlaceParity {
		return &v
	}).(DeinterlaceParityPtrOutput)
}

func (o DeinterlaceParityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeinterlaceParityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeinterlaceParity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeinterlaceParityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeinterlaceParityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeinterlaceParity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeinterlaceParityPtrOutput struct{ *pulumi.OutputState }

func (DeinterlaceParityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeinterlaceParity)(nil)).Elem()
}

func (o DeinterlaceParityPtrOutput) ToDeinterlaceParityPtrOutput() DeinterlaceParityPtrOutput {
	return o
}

func (o DeinterlaceParityPtrOutput) ToDeinterlaceParityPtrOutputWithContext(ctx context.Context) DeinterlaceParityPtrOutput {
	return o
}

func (o DeinterlaceParityPtrOutput) Elem() DeinterlaceParityOutput {
	return o.ApplyT(func(v *DeinterlaceParity) DeinterlaceParity {
		if v != nil {
			return *v
		}
		var ret DeinterlaceParity
		return ret
	}).(DeinterlaceParityOutput)
}

func (o DeinterlaceParityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeinterlaceParityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeinterlaceParity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeinterlaceParityInput interface {
	pulumi.Input

	ToDeinterlaceParityOutput() DeinterlaceParityOutput
	ToDeinterlaceParityOutputWithContext(context.Context) DeinterlaceParityOutput
}

var deinterlaceParityPtrType = reflect.TypeOf((**DeinterlaceParity)(nil)).Elem()

type DeinterlaceParityPtrInput interface {
	pulumi.Input

	ToDeinterlaceParityPtrOutput() DeinterlaceParityPtrOutput
	ToDeinterlaceParityPtrOutputWithContext(context.Context) DeinterlaceParityPtrOutput
}

type deinterlaceParityPtr string

func DeinterlaceParityPtr(v string) DeinterlaceParityPtrInput {
	return (*deinterlaceParityPtr)(&v)
}

func (*deinterlaceParityPtr) ElementType() reflect.Type {
	return deinterlaceParityPtrType
}

func (in *deinterlaceParityPtr) ToDeinterlaceParityPtrOutput() DeinterlaceParityPtrOutput {
	return pulumi.ToOutput(in).(DeinterlaceParityPtrOutput)
}

func (in *deinterlaceParityPtr) ToDeinterlaceParityPtrOutputWithContext(ctx context.Context) DeinterlaceParityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeinterlaceParityPtrOutput)
}

type EncoderNamedPreset string

const (
	// Produces a set of GOP aligned MP4 files with H.264 video and stereo AAC audio. Auto-generates a bitrate ladder based on the input resolution and bitrate. The auto-generated preset will never exceed the input resolution and bitrate. For example, if the input is 720p at 3 Mbps, output will remain 720p at best, and will start at rates lower than 3 Mbps. The output will have video and audio in separate MP4 files, which is optimal for adaptive streaming.
	EncoderNamedPresetAdaptiveStreaming = EncoderNamedPreset("AdaptiveStreaming")
	// Produces a single MP4 file containing only stereo audio encoded at 192 kbps.
	EncoderNamedPresetAACGoodQualityAudio = EncoderNamedPreset("AACGoodQualityAudio")
	// Produces a set of 8 GOP-aligned MP4 files, ranging from 6000 kbps to 400 kbps, and stereo AAC audio. Resolution starts at 1080p and goes down to 360p.
	EncoderNamedPresetH264MultipleBitrate1080p = EncoderNamedPreset("H264MultipleBitrate1080p")
	// Produces a set of 6 GOP-aligned MP4 files, ranging from 3400 kbps to 400 kbps, and stereo AAC audio. Resolution starts at 720p and goes down to 360p.
	EncoderNamedPresetH264MultipleBitrate720p = EncoderNamedPreset("H264MultipleBitrate720p")
	// Produces a set of 5 GOP-aligned MP4 files, ranging from 1600kbps to 400 kbps, and stereo AAC audio. Resolution starts at 480p and goes down to 360p.
	EncoderNamedPresetH264MultipleBitrateSD = EncoderNamedPreset("H264MultipleBitrateSD")
)

func (EncoderNamedPreset) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderNamedPreset)(nil)).Elem()
}

func (e EncoderNamedPreset) ToEncoderNamedPresetOutput() EncoderNamedPresetOutput {
	return pulumi.ToOutput(e).(EncoderNamedPresetOutput)
}

func (e EncoderNamedPreset) ToEncoderNamedPresetOutputWithContext(ctx context.Context) EncoderNamedPresetOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncoderNamedPresetOutput)
}

func (e EncoderNamedPreset) ToEncoderNamedPresetPtrOutput() EncoderNamedPresetPtrOutput {
	return e.ToEncoderNamedPresetPtrOutputWithContext(context.Background())
}

func (e EncoderNamedPreset) ToEncoderNamedPresetPtrOutputWithContext(ctx context.Context) EncoderNamedPresetPtrOutput {
	return EncoderNamedPreset(e).ToEncoderNamedPresetOutputWithContext(ctx).ToEncoderNamedPresetPtrOutputWithContext(ctx)
}

func (e EncoderNamedPreset) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncoderNamedPreset) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncoderNamedPreset) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncoderNamedPreset) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncoderNamedPresetOutput struct{ *pulumi.OutputState }

func (EncoderNamedPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderNamedPreset)(nil)).Elem()
}

func (o EncoderNamedPresetOutput) ToEncoderNamedPresetOutput() EncoderNamedPresetOutput {
	return o
}

func (o EncoderNamedPresetOutput) ToEncoderNamedPresetOutputWithContext(ctx context.Context) EncoderNamedPresetOutput {
	return o
}

func (o EncoderNamedPresetOutput) ToEncoderNamedPresetPtrOutput() EncoderNamedPresetPtrOutput {
	return o.ToEncoderNamedPresetPtrOutputWithContext(context.Background())
}

func (o EncoderNamedPresetOutput) ToEncoderNamedPresetPtrOutputWithContext(ctx context.Context) EncoderNamedPresetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncoderNamedPreset) *EncoderNamedPreset {
		return &v
	}).(EncoderNamedPresetPtrOutput)
}

func (o EncoderNamedPresetOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncoderNamedPresetOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncoderNamedPreset) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncoderNamedPresetOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncoderNamedPresetOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncoderNamedPreset) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncoderNamedPresetPtrOutput struct{ *pulumi.OutputState }

func (EncoderNamedPresetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncoderNamedPreset)(nil)).Elem()
}

func (o EncoderNamedPresetPtrOutput) ToEncoderNamedPresetPtrOutput() EncoderNamedPresetPtrOutput {
	return o
}

func (o EncoderNamedPresetPtrOutput) ToEncoderNamedPresetPtrOutputWithContext(ctx context.Context) EncoderNamedPresetPtrOutput {
	return o
}

func (o EncoderNamedPresetPtrOutput) Elem() EncoderNamedPresetOutput {
	return o.ApplyT(func(v *EncoderNamedPreset) EncoderNamedPreset {
		if v != nil {
			return *v
		}
		var ret EncoderNamedPreset
		return ret
	}).(EncoderNamedPresetOutput)
}

func (o EncoderNamedPresetPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncoderNamedPresetPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncoderNamedPreset) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncoderNamedPresetInput interface {
	pulumi.Input

	ToEncoderNamedPresetOutput() EncoderNamedPresetOutput
	ToEncoderNamedPresetOutputWithContext(context.Context) EncoderNamedPresetOutput
}

var encoderNamedPresetPtrType = reflect.TypeOf((**EncoderNamedPreset)(nil)).Elem()

type EncoderNamedPresetPtrInput interface {
	pulumi.Input

	ToEncoderNamedPresetPtrOutput() EncoderNamedPresetPtrOutput
	ToEncoderNamedPresetPtrOutputWithContext(context.Context) EncoderNamedPresetPtrOutput
}

type encoderNamedPresetPtr string

func EncoderNamedPresetPtr(v string) EncoderNamedPresetPtrInput {
	return (*encoderNamedPresetPtr)(&v)
}

func (*encoderNamedPresetPtr) ElementType() reflect.Type {
	return encoderNamedPresetPtrType
}

func (in *encoderNamedPresetPtr) ToEncoderNamedPresetPtrOutput() EncoderNamedPresetPtrOutput {
	return pulumi.ToOutput(in).(EncoderNamedPresetPtrOutput)
}

func (in *encoderNamedPresetPtr) ToEncoderNamedPresetPtrOutputWithContext(ctx context.Context) EncoderNamedPresetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncoderNamedPresetPtrOutput)
}

type EntropyMode string

const (
	// Context Adaptive Binary Arithmetic Coder (CABAC) entropy encoding.
	EntropyModeCabac = EntropyMode("Cabac")
	// Context Adaptive Variable Length Coder (CAVLC) entropy encoding.
	EntropyModeCavlc = EntropyMode("Cavlc")
)

func (EntropyMode) ElementType() reflect.Type {
	return reflect.TypeOf((*EntropyMode)(nil)).Elem()
}

func (e EntropyMode) ToEntropyModeOutput() EntropyModeOutput {
	return pulumi.ToOutput(e).(EntropyModeOutput)
}

func (e EntropyMode) ToEntropyModeOutputWithContext(ctx context.Context) EntropyModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EntropyModeOutput)
}

func (e EntropyMode) ToEntropyModePtrOutput() EntropyModePtrOutput {
	return e.ToEntropyModePtrOutputWithContext(context.Background())
}

func (e EntropyMode) ToEntropyModePtrOutputWithContext(ctx context.Context) EntropyModePtrOutput {
	return EntropyMode(e).ToEntropyModeOutputWithContext(ctx).ToEntropyModePtrOutputWithContext(ctx)
}

func (e EntropyMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntropyMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntropyMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EntropyMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EntropyModeOutput struct{ *pulumi.OutputState }

func (EntropyModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntropyMode)(nil)).Elem()
}

func (o EntropyModeOutput) ToEntropyModeOutput() EntropyModeOutput {
	return o
}

func (o EntropyModeOutput) ToEntropyModeOutputWithContext(ctx context.Context) EntropyModeOutput {
	return o
}

func (o EntropyModeOutput) ToEntropyModePtrOutput() EntropyModePtrOutput {
	return o.ToEntropyModePtrOutputWithContext(context.Background())
}

func (o EntropyModeOutput) ToEntropyModePtrOutputWithContext(ctx context.Context) EntropyModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntropyMode) *EntropyMode {
		return &v
	}).(EntropyModePtrOutput)
}

func (o EntropyModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EntropyModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntropyMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EntropyModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntropyModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntropyMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EntropyModePtrOutput struct{ *pulumi.OutputState }

func (EntropyModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntropyMode)(nil)).Elem()
}

func (o EntropyModePtrOutput) ToEntropyModePtrOutput() EntropyModePtrOutput {
	return o
}

func (o EntropyModePtrOutput) ToEntropyModePtrOutputWithContext(ctx context.Context) EntropyModePtrOutput {
	return o
}

func (o EntropyModePtrOutput) Elem() EntropyModeOutput {
	return o.ApplyT(func(v *EntropyMode) EntropyMode {
		if v != nil {
			return *v
		}
		var ret EntropyMode
		return ret
	}).(EntropyModeOutput)
}

func (o EntropyModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntropyModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EntropyMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EntropyModeInput interface {
	pulumi.Input

	ToEntropyModeOutput() EntropyModeOutput
	ToEntropyModeOutputWithContext(context.Context) EntropyModeOutput
}

var entropyModePtrType = reflect.TypeOf((**EntropyMode)(nil)).Elem()

type EntropyModePtrInput interface {
	pulumi.Input

	ToEntropyModePtrOutput() EntropyModePtrOutput
	ToEntropyModePtrOutputWithContext(context.Context) EntropyModePtrOutput
}

type entropyModePtr string

func EntropyModePtr(v string) EntropyModePtrInput {
	return (*entropyModePtr)(&v)
}

func (*entropyModePtr) ElementType() reflect.Type {
	return entropyModePtrType
}

func (in *entropyModePtr) ToEntropyModePtrOutput() EntropyModePtrOutput {
	return pulumi.ToOutput(in).(EntropyModePtrOutput)
}

func (in *entropyModePtr) ToEntropyModePtrOutputWithContext(ctx context.Context) EntropyModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EntropyModePtrOutput)
}

type H264Complexity string

const (
	// Tells the encoder to use settings that are optimized for faster encoding. Quality is sacrificed to decrease encoding time.
	H264ComplexitySpeed = H264Complexity("Speed")
	// Tells the encoder to use settings that achieve a balance between speed and quality.
	H264ComplexityBalanced = H264Complexity("Balanced")
	// Tells the encoder to use settings that are optimized to produce higher quality output at the expense of slower overall encode time.
	H264ComplexityQuality = H264Complexity("Quality")
)

func (H264Complexity) ElementType() reflect.Type {
	return reflect.TypeOf((*H264Complexity)(nil)).Elem()
}

func (e H264Complexity) ToH264ComplexityOutput() H264ComplexityOutput {
	return pulumi.ToOutput(e).(H264ComplexityOutput)
}

func (e H264Complexity) ToH264ComplexityOutputWithContext(ctx context.Context) H264ComplexityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(H264ComplexityOutput)
}

func (e H264Complexity) ToH264ComplexityPtrOutput() H264ComplexityPtrOutput {
	return e.ToH264ComplexityPtrOutputWithContext(context.Background())
}

func (e H264Complexity) ToH264ComplexityPtrOutputWithContext(ctx context.Context) H264ComplexityPtrOutput {
	return H264Complexity(e).ToH264ComplexityOutputWithContext(ctx).ToH264ComplexityPtrOutputWithContext(ctx)
}

func (e H264Complexity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e H264Complexity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e H264Complexity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e H264Complexity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type H264ComplexityOutput struct{ *pulumi.OutputState }

func (H264ComplexityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H264Complexity)(nil)).Elem()
}

func (o H264ComplexityOutput) ToH264ComplexityOutput() H264ComplexityOutput {
	return o
}

func (o H264ComplexityOutput) ToH264ComplexityOutputWithContext(ctx context.Context) H264ComplexityOutput {
	return o
}

func (o H264ComplexityOutput) ToH264ComplexityPtrOutput() H264ComplexityPtrOutput {
	return o.ToH264ComplexityPtrOutputWithContext(context.Background())
}

func (o H264ComplexityOutput) ToH264ComplexityPtrOutputWithContext(ctx context.Context) H264ComplexityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v H264Complexity) *H264Complexity {
		return &v
	}).(H264ComplexityPtrOutput)
}

func (o H264ComplexityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o H264ComplexityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e H264Complexity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o H264ComplexityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o H264ComplexityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e H264Complexity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type H264ComplexityPtrOutput struct{ *pulumi.OutputState }

func (H264ComplexityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**H264Complexity)(nil)).Elem()
}

func (o H264ComplexityPtrOutput) ToH264ComplexityPtrOutput() H264ComplexityPtrOutput {
	return o
}

func (o H264ComplexityPtrOutput) ToH264ComplexityPtrOutputWithContext(ctx context.Context) H264ComplexityPtrOutput {
	return o
}

func (o H264ComplexityPtrOutput) Elem() H264ComplexityOutput {
	return o.ApplyT(func(v *H264Complexity) H264Complexity {
		if v != nil {
			return *v
		}
		var ret H264Complexity
		return ret
	}).(H264ComplexityOutput)
}

func (o H264ComplexityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o H264ComplexityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *H264Complexity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type H264ComplexityInput interface {
	pulumi.Input

	ToH264ComplexityOutput() H264ComplexityOutput
	ToH264ComplexityOutputWithContext(context.Context) H264ComplexityOutput
}

var h264complexityPtrType = reflect.TypeOf((**H264Complexity)(nil)).Elem()

type H264ComplexityPtrInput interface {
	pulumi.Input

	ToH264ComplexityPtrOutput() H264ComplexityPtrOutput
	ToH264ComplexityPtrOutputWithContext(context.Context) H264ComplexityPtrOutput
}

type h264complexityPtr string

func H264ComplexityPtr(v string) H264ComplexityPtrInput {
	return (*h264complexityPtr)(&v)
}

func (*h264complexityPtr) ElementType() reflect.Type {
	return h264complexityPtrType
}

func (in *h264complexityPtr) ToH264ComplexityPtrOutput() H264ComplexityPtrOutput {
	return pulumi.ToOutput(in).(H264ComplexityPtrOutput)
}

func (in *h264complexityPtr) ToH264ComplexityPtrOutputWithContext(ctx context.Context) H264ComplexityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(H264ComplexityPtrOutput)
}

type H264VideoProfile string

const (
	// Tells the encoder to automatically determine the appropriate H.264 profile.
	H264VideoProfileAuto = H264VideoProfile("Auto")
	// Baseline profile
	H264VideoProfileBaseline = H264VideoProfile("Baseline")
	// Main profile
	H264VideoProfileMain = H264VideoProfile("Main")
	// High profile.
	H264VideoProfileHigh = H264VideoProfile("High")
	// High 4:2:2 profile.
	H264VideoProfileHigh422 = H264VideoProfile("High422")
	// High 4:4:4 predictive profile.
	H264VideoProfileHigh444 = H264VideoProfile("High444")
)

func (H264VideoProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*H264VideoProfile)(nil)).Elem()
}

func (e H264VideoProfile) ToH264VideoProfileOutput() H264VideoProfileOutput {
	return pulumi.ToOutput(e).(H264VideoProfileOutput)
}

func (e H264VideoProfile) ToH264VideoProfileOutputWithContext(ctx context.Context) H264VideoProfileOutput {
	return pulumi.ToOutputWithContext(ctx, e).(H264VideoProfileOutput)
}

func (e H264VideoProfile) ToH264VideoProfilePtrOutput() H264VideoProfilePtrOutput {
	return e.ToH264VideoProfilePtrOutputWithContext(context.Background())
}

func (e H264VideoProfile) ToH264VideoProfilePtrOutputWithContext(ctx context.Context) H264VideoProfilePtrOutput {
	return H264VideoProfile(e).ToH264VideoProfileOutputWithContext(ctx).ToH264VideoProfilePtrOutputWithContext(ctx)
}

func (e H264VideoProfile) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e H264VideoProfile) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e H264VideoProfile) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e H264VideoProfile) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type H264VideoProfileOutput struct{ *pulumi.OutputState }

func (H264VideoProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*H264VideoProfile)(nil)).Elem()
}

func (o H264VideoProfileOutput) ToH264VideoProfileOutput() H264VideoProfileOutput {
	return o
}

func (o H264VideoProfileOutput) ToH264VideoProfileOutputWithContext(ctx context.Context) H264VideoProfileOutput {
	return o
}

func (o H264VideoProfileOutput) ToH264VideoProfilePtrOutput() H264VideoProfilePtrOutput {
	return o.ToH264VideoProfilePtrOutputWithContext(context.Background())
}

func (o H264VideoProfileOutput) ToH264VideoProfilePtrOutputWithContext(ctx context.Context) H264VideoProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v H264VideoProfile) *H264VideoProfile {
		return &v
	}).(H264VideoProfilePtrOutput)
}

func (o H264VideoProfileOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o H264VideoProfileOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e H264VideoProfile) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o H264VideoProfileOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o H264VideoProfileOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e H264VideoProfile) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type H264VideoProfilePtrOutput struct{ *pulumi.OutputState }

func (H264VideoProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**H264VideoProfile)(nil)).Elem()
}

func (o H264VideoProfilePtrOutput) ToH264VideoProfilePtrOutput() H264VideoProfilePtrOutput {
	return o
}

func (o H264VideoProfilePtrOutput) ToH264VideoProfilePtrOutputWithContext(ctx context.Context) H264VideoProfilePtrOutput {
	return o
}

func (o H264VideoProfilePtrOutput) Elem() H264VideoProfileOutput {
	return o.ApplyT(func(v *H264VideoProfile) H264VideoProfile {
		if v != nil {
			return *v
		}
		var ret H264VideoProfile
		return ret
	}).(H264VideoProfileOutput)
}

func (o H264VideoProfilePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o H264VideoProfilePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *H264VideoProfile) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type H264VideoProfileInput interface {
	pulumi.Input

	ToH264VideoProfileOutput() H264VideoProfileOutput
	ToH264VideoProfileOutputWithContext(context.Context) H264VideoProfileOutput
}

var h264videoProfilePtrType = reflect.TypeOf((**H264VideoProfile)(nil)).Elem()

type H264VideoProfilePtrInput interface {
	pulumi.Input

	ToH264VideoProfilePtrOutput() H264VideoProfilePtrOutput
	ToH264VideoProfilePtrOutputWithContext(context.Context) H264VideoProfilePtrOutput
}

type h264videoProfilePtr string

func H264VideoProfilePtr(v string) H264VideoProfilePtrInput {
	return (*h264videoProfilePtr)(&v)
}

func (*h264videoProfilePtr) ElementType() reflect.Type {
	return h264videoProfilePtrType
}

func (in *h264videoProfilePtr) ToH264VideoProfilePtrOutput() H264VideoProfilePtrOutput {
	return pulumi.ToOutput(in).(H264VideoProfilePtrOutput)
}

func (in *h264videoProfilePtr) ToH264VideoProfilePtrOutputWithContext(ctx context.Context) H264VideoProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(H264VideoProfilePtrOutput)
}

type LiveEventEncodingType string

const (
	LiveEventEncodingTypeNone  = LiveEventEncodingType("None")
	LiveEventEncodingTypeBasic = LiveEventEncodingType("Basic")
)

func (LiveEventEncodingType) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingType)(nil)).Elem()
}

func (e LiveEventEncodingType) ToLiveEventEncodingTypeOutput() LiveEventEncodingTypeOutput {
	return pulumi.ToOutput(e).(LiveEventEncodingTypeOutput)
}

func (e LiveEventEncodingType) ToLiveEventEncodingTypeOutputWithContext(ctx context.Context) LiveEventEncodingTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LiveEventEncodingTypeOutput)
}

func (e LiveEventEncodingType) ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput {
	return e.ToLiveEventEncodingTypePtrOutputWithContext(context.Background())
}

func (e LiveEventEncodingType) ToLiveEventEncodingTypePtrOutputWithContext(ctx context.Context) LiveEventEncodingTypePtrOutput {
	return LiveEventEncodingType(e).ToLiveEventEncodingTypeOutputWithContext(ctx).ToLiveEventEncodingTypePtrOutputWithContext(ctx)
}

func (e LiveEventEncodingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LiveEventEncodingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LiveEventEncodingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LiveEventEncodingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LiveEventEncodingTypeOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingType)(nil)).Elem()
}

func (o LiveEventEncodingTypeOutput) ToLiveEventEncodingTypeOutput() LiveEventEncodingTypeOutput {
	return o
}

func (o LiveEventEncodingTypeOutput) ToLiveEventEncodingTypeOutputWithContext(ctx context.Context) LiveEventEncodingTypeOutput {
	return o
}

func (o LiveEventEncodingTypeOutput) ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput {
	return o.ToLiveEventEncodingTypePtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingTypeOutput) ToLiveEventEncodingTypePtrOutputWithContext(ctx context.Context) LiveEventEncodingTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventEncodingType) *LiveEventEncodingType {
		return &v
	}).(LiveEventEncodingTypePtrOutput)
}

func (o LiveEventEncodingTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LiveEventEncodingTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LiveEventEncodingType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LiveEventEncodingTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LiveEventEncodingType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LiveEventEncodingTypePtrOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncodingType)(nil)).Elem()
}

func (o LiveEventEncodingTypePtrOutput) ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput {
	return o
}

func (o LiveEventEncodingTypePtrOutput) ToLiveEventEncodingTypePtrOutputWithContext(ctx context.Context) LiveEventEncodingTypePtrOutput {
	return o
}

func (o LiveEventEncodingTypePtrOutput) Elem() LiveEventEncodingTypeOutput {
	return o.ApplyT(func(v *LiveEventEncodingType) LiveEventEncodingType {
		if v != nil {
			return *v
		}
		var ret LiveEventEncodingType
		return ret
	}).(LiveEventEncodingTypeOutput)
}

func (o LiveEventEncodingTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LiveEventEncodingType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LiveEventEncodingTypeInput interface {
	pulumi.Input

	ToLiveEventEncodingTypeOutput() LiveEventEncodingTypeOutput
	ToLiveEventEncodingTypeOutputWithContext(context.Context) LiveEventEncodingTypeOutput
}

var liveEventEncodingTypePtrType = reflect.TypeOf((**LiveEventEncodingType)(nil)).Elem()

type LiveEventEncodingTypePtrInput interface {
	pulumi.Input

	ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput
	ToLiveEventEncodingTypePtrOutputWithContext(context.Context) LiveEventEncodingTypePtrOutput
}

type liveEventEncodingTypePtr string

func LiveEventEncodingTypePtr(v string) LiveEventEncodingTypePtrInput {
	return (*liveEventEncodingTypePtr)(&v)
}

func (*liveEventEncodingTypePtr) ElementType() reflect.Type {
	return liveEventEncodingTypePtrType
}

func (in *liveEventEncodingTypePtr) ToLiveEventEncodingTypePtrOutput() LiveEventEncodingTypePtrOutput {
	return pulumi.ToOutput(in).(LiveEventEncodingTypePtrOutput)
}

func (in *liveEventEncodingTypePtr) ToLiveEventEncodingTypePtrOutputWithContext(ctx context.Context) LiveEventEncodingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LiveEventEncodingTypePtrOutput)
}

type LiveEventInputProtocol string

const (
	LiveEventInputProtocolFragmentedMP4 = LiveEventInputProtocol("FragmentedMP4")
	LiveEventInputProtocolRTMP          = LiveEventInputProtocol("RTMP")
)

func (LiveEventInputProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputProtocol)(nil)).Elem()
}

func (e LiveEventInputProtocol) ToLiveEventInputProtocolOutput() LiveEventInputProtocolOutput {
	return pulumi.ToOutput(e).(LiveEventInputProtocolOutput)
}

func (e LiveEventInputProtocol) ToLiveEventInputProtocolOutputWithContext(ctx context.Context) LiveEventInputProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LiveEventInputProtocolOutput)
}

func (e LiveEventInputProtocol) ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput {
	return e.ToLiveEventInputProtocolPtrOutputWithContext(context.Background())
}

func (e LiveEventInputProtocol) ToLiveEventInputProtocolPtrOutputWithContext(ctx context.Context) LiveEventInputProtocolPtrOutput {
	return LiveEventInputProtocol(e).ToLiveEventInputProtocolOutputWithContext(ctx).ToLiveEventInputProtocolPtrOutputWithContext(ctx)
}

func (e LiveEventInputProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LiveEventInputProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LiveEventInputProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LiveEventInputProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LiveEventInputProtocolOutput struct{ *pulumi.OutputState }

func (LiveEventInputProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputProtocol)(nil)).Elem()
}

func (o LiveEventInputProtocolOutput) ToLiveEventInputProtocolOutput() LiveEventInputProtocolOutput {
	return o
}

func (o LiveEventInputProtocolOutput) ToLiveEventInputProtocolOutputWithContext(ctx context.Context) LiveEventInputProtocolOutput {
	return o
}

func (o LiveEventInputProtocolOutput) ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput {
	return o.ToLiveEventInputProtocolPtrOutputWithContext(context.Background())
}

func (o LiveEventInputProtocolOutput) ToLiveEventInputProtocolPtrOutputWithContext(ctx context.Context) LiveEventInputProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventInputProtocol) *LiveEventInputProtocol {
		return &v
	}).(LiveEventInputProtocolPtrOutput)
}

func (o LiveEventInputProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LiveEventInputProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LiveEventInputProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LiveEventInputProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LiveEventInputProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LiveEventInputProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LiveEventInputProtocolPtrOutput struct{ *pulumi.OutputState }

func (LiveEventInputProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventInputProtocol)(nil)).Elem()
}

func (o LiveEventInputProtocolPtrOutput) ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput {
	return o
}

func (o LiveEventInputProtocolPtrOutput) ToLiveEventInputProtocolPtrOutputWithContext(ctx context.Context) LiveEventInputProtocolPtrOutput {
	return o
}

func (o LiveEventInputProtocolPtrOutput) Elem() LiveEventInputProtocolOutput {
	return o.ApplyT(func(v *LiveEventInputProtocol) LiveEventInputProtocol {
		if v != nil {
			return *v
		}
		var ret LiveEventInputProtocol
		return ret
	}).(LiveEventInputProtocolOutput)
}

func (o LiveEventInputProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LiveEventInputProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LiveEventInputProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LiveEventInputProtocolInput interface {
	pulumi.Input

	ToLiveEventInputProtocolOutput() LiveEventInputProtocolOutput
	ToLiveEventInputProtocolOutputWithContext(context.Context) LiveEventInputProtocolOutput
}

var liveEventInputProtocolPtrType = reflect.TypeOf((**LiveEventInputProtocol)(nil)).Elem()

type LiveEventInputProtocolPtrInput interface {
	pulumi.Input

	ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput
	ToLiveEventInputProtocolPtrOutputWithContext(context.Context) LiveEventInputProtocolPtrOutput
}

type liveEventInputProtocolPtr string

func LiveEventInputProtocolPtr(v string) LiveEventInputProtocolPtrInput {
	return (*liveEventInputProtocolPtr)(&v)
}

func (*liveEventInputProtocolPtr) ElementType() reflect.Type {
	return liveEventInputProtocolPtrType
}

func (in *liveEventInputProtocolPtr) ToLiveEventInputProtocolPtrOutput() LiveEventInputProtocolPtrOutput {
	return pulumi.ToOutput(in).(LiveEventInputProtocolPtrOutput)
}

func (in *liveEventInputProtocolPtr) ToLiveEventInputProtocolPtrOutputWithContext(ctx context.Context) LiveEventInputProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LiveEventInputProtocolPtrOutput)
}

type OnErrorType string

const (
	// Tells the service that if this TransformOutput fails, then any other incomplete TransformOutputs can be stopped.
	OnErrorTypeStopProcessingJob = OnErrorType("StopProcessingJob")
	// Tells the service that if this TransformOutput fails, then allow any other TransformOutput to continue.
	OnErrorTypeContinueJob = OnErrorType("ContinueJob")
)

func (OnErrorType) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorType)(nil)).Elem()
}

func (e OnErrorType) ToOnErrorTypeOutput() OnErrorTypeOutput {
	return pulumi.ToOutput(e).(OnErrorTypeOutput)
}

func (e OnErrorType) ToOnErrorTypeOutputWithContext(ctx context.Context) OnErrorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OnErrorTypeOutput)
}

func (e OnErrorType) ToOnErrorTypePtrOutput() OnErrorTypePtrOutput {
	return e.ToOnErrorTypePtrOutputWithContext(context.Background())
}

func (e OnErrorType) ToOnErrorTypePtrOutputWithContext(ctx context.Context) OnErrorTypePtrOutput {
	return OnErrorType(e).ToOnErrorTypeOutputWithContext(ctx).ToOnErrorTypePtrOutputWithContext(ctx)
}

func (e OnErrorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OnErrorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OnErrorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OnErrorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OnErrorTypeOutput struct{ *pulumi.OutputState }

func (OnErrorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorType)(nil)).Elem()
}

func (o OnErrorTypeOutput) ToOnErrorTypeOutput() OnErrorTypeOutput {
	return o
}

func (o OnErrorTypeOutput) ToOnErrorTypeOutputWithContext(ctx context.Context) OnErrorTypeOutput {
	return o
}

func (o OnErrorTypeOutput) ToOnErrorTypePtrOutput() OnErrorTypePtrOutput {
	return o.ToOnErrorTypePtrOutputWithContext(context.Background())
}

func (o OnErrorTypeOutput) ToOnErrorTypePtrOutputWithContext(ctx context.Context) OnErrorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OnErrorType) *OnErrorType {
		return &v
	}).(OnErrorTypePtrOutput)
}

func (o OnErrorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OnErrorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OnErrorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OnErrorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OnErrorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OnErrorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OnErrorTypePtrOutput struct{ *pulumi.OutputState }

func (OnErrorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnErrorType)(nil)).Elem()
}

func (o OnErrorTypePtrOutput) ToOnErrorTypePtrOutput() OnErrorTypePtrOutput {
	return o
}

func (o OnErrorTypePtrOutput) ToOnErrorTypePtrOutputWithContext(ctx context.Context) OnErrorTypePtrOutput {
	return o
}

func (o OnErrorTypePtrOutput) Elem() OnErrorTypeOutput {
	return o.ApplyT(func(v *OnErrorType) OnErrorType {
		if v != nil {
			return *v
		}
		var ret OnErrorType
		return ret
	}).(OnErrorTypeOutput)
}

func (o OnErrorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OnErrorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OnErrorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OnErrorTypeInput interface {
	pulumi.Input

	ToOnErrorTypeOutput() OnErrorTypeOutput
	ToOnErrorTypeOutputWithContext(context.Context) OnErrorTypeOutput
}

var onErrorTypePtrType = reflect.TypeOf((**OnErrorType)(nil)).Elem()

type OnErrorTypePtrInput interface {
	pulumi.Input

	ToOnErrorTypePtrOutput() OnErrorTypePtrOutput
	ToOnErrorTypePtrOutputWithContext(context.Context) OnErrorTypePtrOutput
}

type onErrorTypePtr string

func OnErrorTypePtr(v string) OnErrorTypePtrInput {
	return (*onErrorTypePtr)(&v)
}

func (*onErrorTypePtr) ElementType() reflect.Type {
	return onErrorTypePtrType
}

func (in *onErrorTypePtr) ToOnErrorTypePtrOutput() OnErrorTypePtrOutput {
	return pulumi.ToOutput(in).(OnErrorTypePtrOutput)
}

func (in *onErrorTypePtr) ToOnErrorTypePtrOutputWithContext(ctx context.Context) OnErrorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OnErrorTypePtrOutput)
}

type Priority string

const (
	// Used for TransformOutputs that can be generated after Normal and High priority TransformOutputs.
	PriorityLow = Priority("Low")
	// Used for TransformOutputs that can be generated at Normal priority.
	PriorityNormal = Priority("Normal")
	// Used for TransformOutputs that should take precedence over others.
	PriorityHigh = Priority("High")
)

func (Priority) ElementType() reflect.Type {
	return reflect.TypeOf((*Priority)(nil)).Elem()
}

func (e Priority) ToPriorityOutput() PriorityOutput {
	return pulumi.ToOutput(e).(PriorityOutput)
}

func (e Priority) ToPriorityOutputWithContext(ctx context.Context) PriorityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PriorityOutput)
}

func (e Priority) ToPriorityPtrOutput() PriorityPtrOutput {
	return e.ToPriorityPtrOutputWithContext(context.Background())
}

func (e Priority) ToPriorityPtrOutputWithContext(ctx context.Context) PriorityPtrOutput {
	return Priority(e).ToPriorityOutputWithContext(ctx).ToPriorityPtrOutputWithContext(ctx)
}

func (e Priority) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Priority) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Priority) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Priority) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PriorityOutput struct{ *pulumi.OutputState }

func (PriorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Priority)(nil)).Elem()
}

func (o PriorityOutput) ToPriorityOutput() PriorityOutput {
	return o
}

func (o PriorityOutput) ToPriorityOutputWithContext(ctx context.Context) PriorityOutput {
	return o
}

func (o PriorityOutput) ToPriorityPtrOutput() PriorityPtrOutput {
	return o.ToPriorityPtrOutputWithContext(context.Background())
}

func (o PriorityOutput) ToPriorityPtrOutputWithContext(ctx context.Context) PriorityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Priority) *Priority {
		return &v
	}).(PriorityPtrOutput)
}

func (o PriorityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PriorityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Priority) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PriorityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PriorityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Priority) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PriorityPtrOutput struct{ *pulumi.OutputState }

func (PriorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Priority)(nil)).Elem()
}

func (o PriorityPtrOutput) ToPriorityPtrOutput() PriorityPtrOutput {
	return o
}

func (o PriorityPtrOutput) ToPriorityPtrOutputWithContext(ctx context.Context) PriorityPtrOutput {
	return o
}

func (o PriorityPtrOutput) Elem() PriorityOutput {
	return o.ApplyT(func(v *Priority) Priority {
		if v != nil {
			return *v
		}
		var ret Priority
		return ret
	}).(PriorityOutput)
}

func (o PriorityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PriorityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Priority) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PriorityInput interface {
	pulumi.Input

	ToPriorityOutput() PriorityOutput
	ToPriorityOutputWithContext(context.Context) PriorityOutput
}

var priorityPtrType = reflect.TypeOf((**Priority)(nil)).Elem()

type PriorityPtrInput interface {
	pulumi.Input

	ToPriorityPtrOutput() PriorityPtrOutput
	ToPriorityPtrOutputWithContext(context.Context) PriorityPtrOutput
}

type priorityPtr string

func PriorityPtr(v string) PriorityPtrInput {
	return (*priorityPtr)(&v)
}

func (*priorityPtr) ElementType() reflect.Type {
	return priorityPtrType
}

func (in *priorityPtr) ToPriorityPtrOutput() PriorityPtrOutput {
	return pulumi.ToOutput(in).(PriorityPtrOutput)
}

func (in *priorityPtr) ToPriorityPtrOutputWithContext(ctx context.Context) PriorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PriorityPtrOutput)
}

type Rotation string

const (
	// Automatically detect and rotate as needed.
	RotationAuto = Rotation("Auto")
	// Do not rotate the video.  If the output format supports it, any metadata about rotation is kept intact.
	RotationNone = Rotation("None")
	// Do not rotate the video but remove any metadata about the rotation.
	RotationRotate0 = Rotation("Rotate0")
	// Rotate 90 degrees clockwise.
	RotationRotate90 = Rotation("Rotate90")
	// Rotate 180 degrees clockwise.
	RotationRotate180 = Rotation("Rotate180")
	// Rotate 270 degrees clockwise.
	RotationRotate270 = Rotation("Rotate270")
)

func (Rotation) ElementType() reflect.Type {
	return reflect.TypeOf((*Rotation)(nil)).Elem()
}

func (e Rotation) ToRotationOutput() RotationOutput {
	return pulumi.ToOutput(e).(RotationOutput)
}

func (e Rotation) ToRotationOutputWithContext(ctx context.Context) RotationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RotationOutput)
}

func (e Rotation) ToRotationPtrOutput() RotationPtrOutput {
	return e.ToRotationPtrOutputWithContext(context.Background())
}

func (e Rotation) ToRotationPtrOutputWithContext(ctx context.Context) RotationPtrOutput {
	return Rotation(e).ToRotationOutputWithContext(ctx).ToRotationPtrOutputWithContext(ctx)
}

func (e Rotation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Rotation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Rotation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Rotation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RotationOutput struct{ *pulumi.OutputState }

func (RotationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rotation)(nil)).Elem()
}

func (o RotationOutput) ToRotationOutput() RotationOutput {
	return o
}

func (o RotationOutput) ToRotationOutputWithContext(ctx context.Context) RotationOutput {
	return o
}

func (o RotationOutput) ToRotationPtrOutput() RotationPtrOutput {
	return o.ToRotationPtrOutputWithContext(context.Background())
}

func (o RotationOutput) ToRotationPtrOutputWithContext(ctx context.Context) RotationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Rotation) *Rotation {
		return &v
	}).(RotationPtrOutput)
}

func (o RotationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RotationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Rotation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RotationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RotationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Rotation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RotationPtrOutput struct{ *pulumi.OutputState }

func (RotationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rotation)(nil)).Elem()
}

func (o RotationPtrOutput) ToRotationPtrOutput() RotationPtrOutput {
	return o
}

func (o RotationPtrOutput) ToRotationPtrOutputWithContext(ctx context.Context) RotationPtrOutput {
	return o
}

func (o RotationPtrOutput) Elem() RotationOutput {
	return o.ApplyT(func(v *Rotation) Rotation {
		if v != nil {
			return *v
		}
		var ret Rotation
		return ret
	}).(RotationOutput)
}

func (o RotationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RotationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Rotation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RotationInput interface {
	pulumi.Input

	ToRotationOutput() RotationOutput
	ToRotationOutputWithContext(context.Context) RotationOutput
}

var rotationPtrType = reflect.TypeOf((**Rotation)(nil)).Elem()

type RotationPtrInput interface {
	pulumi.Input

	ToRotationPtrOutput() RotationPtrOutput
	ToRotationPtrOutputWithContext(context.Context) RotationPtrOutput
}

type rotationPtr string

func RotationPtr(v string) RotationPtrInput {
	return (*rotationPtr)(&v)
}

func (*rotationPtr) ElementType() reflect.Type {
	return rotationPtrType
}

func (in *rotationPtr) ToRotationPtrOutput() RotationPtrOutput {
	return pulumi.ToOutput(in).(RotationPtrOutput)
}

func (in *rotationPtr) ToRotationPtrOutputWithContext(ctx context.Context) RotationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RotationPtrOutput)
}

type StorageAccountType string

const (
	// The primary storage account for the Media Services account.
	StorageAccountTypePrimary = StorageAccountType("Primary")
	// A secondary storage account for the Media Services account.
	StorageAccountTypeSecondary = StorageAccountType("Secondary")
)

func (StorageAccountType) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountType)(nil)).Elem()
}

func (e StorageAccountType) ToStorageAccountTypeOutput() StorageAccountTypeOutput {
	return pulumi.ToOutput(e).(StorageAccountTypeOutput)
}

func (e StorageAccountType) ToStorageAccountTypeOutputWithContext(ctx context.Context) StorageAccountTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageAccountTypeOutput)
}

func (e StorageAccountType) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return e.ToStorageAccountTypePtrOutputWithContext(context.Background())
}

func (e StorageAccountType) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return StorageAccountType(e).ToStorageAccountTypeOutputWithContext(ctx).ToStorageAccountTypePtrOutputWithContext(ctx)
}

func (e StorageAccountType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAccountType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageAccountTypeOutput struct{ *pulumi.OutputState }

func (StorageAccountTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountType)(nil)).Elem()
}

func (o StorageAccountTypeOutput) ToStorageAccountTypeOutput() StorageAccountTypeOutput {
	return o
}

func (o StorageAccountTypeOutput) ToStorageAccountTypeOutputWithContext(ctx context.Context) StorageAccountTypeOutput {
	return o
}

func (o StorageAccountTypeOutput) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return o.ToStorageAccountTypePtrOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountType) *StorageAccountType {
		return &v
	}).(StorageAccountTypePtrOutput)
}

func (o StorageAccountTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageAccountTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageAccountTypePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountType)(nil)).Elem()
}

func (o StorageAccountTypePtrOutput) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return o
}

func (o StorageAccountTypePtrOutput) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return o
}

func (o StorageAccountTypePtrOutput) Elem() StorageAccountTypeOutput {
	return o.ApplyT(func(v *StorageAccountType) StorageAccountType {
		if v != nil {
			return *v
		}
		var ret StorageAccountType
		return ret
	}).(StorageAccountTypeOutput)
}

func (o StorageAccountTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageAccountType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageAccountTypeInput interface {
	pulumi.Input

	ToStorageAccountTypeOutput() StorageAccountTypeOutput
	ToStorageAccountTypeOutputWithContext(context.Context) StorageAccountTypeOutput
}

var storageAccountTypePtrType = reflect.TypeOf((**StorageAccountType)(nil)).Elem()

type StorageAccountTypePtrInput interface {
	pulumi.Input

	ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput
	ToStorageAccountTypePtrOutputWithContext(context.Context) StorageAccountTypePtrOutput
}

type storageAccountTypePtr string

func StorageAccountTypePtr(v string) StorageAccountTypePtrInput {
	return (*storageAccountTypePtr)(&v)
}

func (*storageAccountTypePtr) ElementType() reflect.Type {
	return storageAccountTypePtrType
}

func (in *storageAccountTypePtr) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return pulumi.ToOutput(in).(StorageAccountTypePtrOutput)
}

func (in *storageAccountTypePtr) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageAccountTypePtrOutput)
}

type StreamOptionsFlag string

const (
	StreamOptionsFlagDefault    = StreamOptionsFlag("Default")
	StreamOptionsFlagLowLatency = StreamOptionsFlag("LowLatency")
)

func (StreamOptionsFlag) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamOptionsFlag)(nil)).Elem()
}

func (e StreamOptionsFlag) ToStreamOptionsFlagOutput() StreamOptionsFlagOutput {
	return pulumi.ToOutput(e).(StreamOptionsFlagOutput)
}

func (e StreamOptionsFlag) ToStreamOptionsFlagOutputWithContext(ctx context.Context) StreamOptionsFlagOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StreamOptionsFlagOutput)
}

func (e StreamOptionsFlag) ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput {
	return e.ToStreamOptionsFlagPtrOutputWithContext(context.Background())
}

func (e StreamOptionsFlag) ToStreamOptionsFlagPtrOutputWithContext(ctx context.Context) StreamOptionsFlagPtrOutput {
	return StreamOptionsFlag(e).ToStreamOptionsFlagOutputWithContext(ctx).ToStreamOptionsFlagPtrOutputWithContext(ctx)
}

func (e StreamOptionsFlag) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StreamOptionsFlag) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StreamOptionsFlag) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StreamOptionsFlag) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StreamOptionsFlagOutput struct{ *pulumi.OutputState }

func (StreamOptionsFlagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamOptionsFlag)(nil)).Elem()
}

func (o StreamOptionsFlagOutput) ToStreamOptionsFlagOutput() StreamOptionsFlagOutput {
	return o
}

func (o StreamOptionsFlagOutput) ToStreamOptionsFlagOutputWithContext(ctx context.Context) StreamOptionsFlagOutput {
	return o
}

func (o StreamOptionsFlagOutput) ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput {
	return o.ToStreamOptionsFlagPtrOutputWithContext(context.Background())
}

func (o StreamOptionsFlagOutput) ToStreamOptionsFlagPtrOutputWithContext(ctx context.Context) StreamOptionsFlagPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamOptionsFlag) *StreamOptionsFlag {
		return &v
	}).(StreamOptionsFlagPtrOutput)
}

func (o StreamOptionsFlagOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StreamOptionsFlagOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StreamOptionsFlag) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StreamOptionsFlagOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StreamOptionsFlagOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StreamOptionsFlag) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StreamOptionsFlagPtrOutput struct{ *pulumi.OutputState }

func (StreamOptionsFlagPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamOptionsFlag)(nil)).Elem()
}

func (o StreamOptionsFlagPtrOutput) ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput {
	return o
}

func (o StreamOptionsFlagPtrOutput) ToStreamOptionsFlagPtrOutputWithContext(ctx context.Context) StreamOptionsFlagPtrOutput {
	return o
}

func (o StreamOptionsFlagPtrOutput) Elem() StreamOptionsFlagOutput {
	return o.ApplyT(func(v *StreamOptionsFlag) StreamOptionsFlag {
		if v != nil {
			return *v
		}
		var ret StreamOptionsFlag
		return ret
	}).(StreamOptionsFlagOutput)
}

func (o StreamOptionsFlagPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StreamOptionsFlagPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StreamOptionsFlag) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StreamOptionsFlagInput interface {
	pulumi.Input

	ToStreamOptionsFlagOutput() StreamOptionsFlagOutput
	ToStreamOptionsFlagOutputWithContext(context.Context) StreamOptionsFlagOutput
}

var streamOptionsFlagPtrType = reflect.TypeOf((**StreamOptionsFlag)(nil)).Elem()

type StreamOptionsFlagPtrInput interface {
	pulumi.Input

	ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput
	ToStreamOptionsFlagPtrOutputWithContext(context.Context) StreamOptionsFlagPtrOutput
}

type streamOptionsFlagPtr string

func StreamOptionsFlagPtr(v string) StreamOptionsFlagPtrInput {
	return (*streamOptionsFlagPtr)(&v)
}

func (*streamOptionsFlagPtr) ElementType() reflect.Type {
	return streamOptionsFlagPtrType
}

func (in *streamOptionsFlagPtr) ToStreamOptionsFlagPtrOutput() StreamOptionsFlagPtrOutput {
	return pulumi.ToOutput(in).(StreamOptionsFlagPtrOutput)
}

func (in *streamOptionsFlagPtr) ToStreamOptionsFlagPtrOutputWithContext(ctx context.Context) StreamOptionsFlagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StreamOptionsFlagPtrOutput)
}





type StreamOptionsFlagArrayInput interface {
	pulumi.Input

	ToStreamOptionsFlagArrayOutput() StreamOptionsFlagArrayOutput
	ToStreamOptionsFlagArrayOutputWithContext(context.Context) StreamOptionsFlagArrayOutput
}

type StreamOptionsFlagArray []StreamOptionsFlag

func (StreamOptionsFlagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamOptionsFlag)(nil)).Elem()
}

func (i StreamOptionsFlagArray) ToStreamOptionsFlagArrayOutput() StreamOptionsFlagArrayOutput {
	return i.ToStreamOptionsFlagArrayOutputWithContext(context.Background())
}

func (i StreamOptionsFlagArray) ToStreamOptionsFlagArrayOutputWithContext(ctx context.Context) StreamOptionsFlagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamOptionsFlagArrayOutput)
}

type StreamOptionsFlagArrayOutput struct{ *pulumi.OutputState }

func (StreamOptionsFlagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamOptionsFlag)(nil)).Elem()
}

func (o StreamOptionsFlagArrayOutput) ToStreamOptionsFlagArrayOutput() StreamOptionsFlagArrayOutput {
	return o
}

func (o StreamOptionsFlagArrayOutput) ToStreamOptionsFlagArrayOutputWithContext(ctx context.Context) StreamOptionsFlagArrayOutput {
	return o
}

func (o StreamOptionsFlagArrayOutput) Index(i pulumi.IntInput) StreamOptionsFlagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamOptionsFlag {
		return vs[0].([]StreamOptionsFlag)[vs[1].(int)]
	}).(StreamOptionsFlagOutput)
}

type StretchMode string

const (
	// Strictly respect the output resolution without considering the pixel aspect ratio or display aspect ratio of the input video.
	StretchModeNone = StretchMode("None")
	// Override the output resolution, and change it to match the display aspect ratio of the input, without padding. For example, if the input is 1920x1080 and the encoding preset asks for 1280x1280, then the value in the preset is overridden, and the output will be at 1280x720, which maintains the input aspect ratio of 16:9.
	StretchModeAutoSize = StretchMode("AutoSize")
	// Pad the output (with either letterbox or pillar box) to honor the output resolution, while ensuring that the active video region in the output has the same aspect ratio as the input. For example, if the input is 1920x1080 and the encoding preset asks for 1280x1280, then the output will be at 1280x1280, which contains an inner rectangle of 1280x720 at aspect ratio of 16:9, and pillar box regions 280 pixels wide at the left and right.
	StretchModeAutoFit = StretchMode("AutoFit")
)

func (StretchMode) ElementType() reflect.Type {
	return reflect.TypeOf((*StretchMode)(nil)).Elem()
}

func (e StretchMode) ToStretchModeOutput() StretchModeOutput {
	return pulumi.ToOutput(e).(StretchModeOutput)
}

func (e StretchMode) ToStretchModeOutputWithContext(ctx context.Context) StretchModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StretchModeOutput)
}

func (e StretchMode) ToStretchModePtrOutput() StretchModePtrOutput {
	return e.ToStretchModePtrOutputWithContext(context.Background())
}

func (e StretchMode) ToStretchModePtrOutputWithContext(ctx context.Context) StretchModePtrOutput {
	return StretchMode(e).ToStretchModeOutputWithContext(ctx).ToStretchModePtrOutputWithContext(ctx)
}

func (e StretchMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StretchMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StretchMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StretchMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StretchModeOutput struct{ *pulumi.OutputState }

func (StretchModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StretchMode)(nil)).Elem()
}

func (o StretchModeOutput) ToStretchModeOutput() StretchModeOutput {
	return o
}

func (o StretchModeOutput) ToStretchModeOutputWithContext(ctx context.Context) StretchModeOutput {
	return o
}

func (o StretchModeOutput) ToStretchModePtrOutput() StretchModePtrOutput {
	return o.ToStretchModePtrOutputWithContext(context.Background())
}

func (o StretchModeOutput) ToStretchModePtrOutputWithContext(ctx context.Context) StretchModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StretchMode) *StretchMode {
		return &v
	}).(StretchModePtrOutput)
}

func (o StretchModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StretchModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StretchMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StretchModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StretchModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StretchMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StretchModePtrOutput struct{ *pulumi.OutputState }

func (StretchModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StretchMode)(nil)).Elem()
}

func (o StretchModePtrOutput) ToStretchModePtrOutput() StretchModePtrOutput {
	return o
}

func (o StretchModePtrOutput) ToStretchModePtrOutputWithContext(ctx context.Context) StretchModePtrOutput {
	return o
}

func (o StretchModePtrOutput) Elem() StretchModeOutput {
	return o.ApplyT(func(v *StretchMode) StretchMode {
		if v != nil {
			return *v
		}
		var ret StretchMode
		return ret
	}).(StretchModeOutput)
}

func (o StretchModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StretchModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StretchMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StretchModeInput interface {
	pulumi.Input

	ToStretchModeOutput() StretchModeOutput
	ToStretchModeOutputWithContext(context.Context) StretchModeOutput
}

var stretchModePtrType = reflect.TypeOf((**StretchMode)(nil)).Elem()

type StretchModePtrInput interface {
	pulumi.Input

	ToStretchModePtrOutput() StretchModePtrOutput
	ToStretchModePtrOutputWithContext(context.Context) StretchModePtrOutput
}

type stretchModePtr string

func StretchModePtr(v string) StretchModePtrInput {
	return (*stretchModePtr)(&v)
}

func (*stretchModePtr) ElementType() reflect.Type {
	return stretchModePtrType
}

func (in *stretchModePtr) ToStretchModePtrOutput() StretchModePtrOutput {
	return pulumi.ToOutput(in).(StretchModePtrOutput)
}

func (in *stretchModePtr) ToStretchModePtrOutputWithContext(ctx context.Context) StretchModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StretchModePtrOutput)
}

type TrackPropertyCompareOperation string

const (
	// Unknown track property compare operation
	TrackPropertyCompareOperationUnknown = TrackPropertyCompareOperation("Unknown")
	// Equal operation
	TrackPropertyCompareOperationEqual = TrackPropertyCompareOperation("Equal")
)

func (TrackPropertyCompareOperation) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyCompareOperation)(nil)).Elem()
}

func (e TrackPropertyCompareOperation) ToTrackPropertyCompareOperationOutput() TrackPropertyCompareOperationOutput {
	return pulumi.ToOutput(e).(TrackPropertyCompareOperationOutput)
}

func (e TrackPropertyCompareOperation) ToTrackPropertyCompareOperationOutputWithContext(ctx context.Context) TrackPropertyCompareOperationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TrackPropertyCompareOperationOutput)
}

func (e TrackPropertyCompareOperation) ToTrackPropertyCompareOperationPtrOutput() TrackPropertyCompareOperationPtrOutput {
	return e.ToTrackPropertyCompareOperationPtrOutputWithContext(context.Background())
}

func (e TrackPropertyCompareOperation) ToTrackPropertyCompareOperationPtrOutputWithContext(ctx context.Context) TrackPropertyCompareOperationPtrOutput {
	return TrackPropertyCompareOperation(e).ToTrackPropertyCompareOperationOutputWithContext(ctx).ToTrackPropertyCompareOperationPtrOutputWithContext(ctx)
}

func (e TrackPropertyCompareOperation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrackPropertyCompareOperation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrackPropertyCompareOperation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TrackPropertyCompareOperation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TrackPropertyCompareOperationOutput struct{ *pulumi.OutputState }

func (TrackPropertyCompareOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyCompareOperation)(nil)).Elem()
}

func (o TrackPropertyCompareOperationOutput) ToTrackPropertyCompareOperationOutput() TrackPropertyCompareOperationOutput {
	return o
}

func (o TrackPropertyCompareOperationOutput) ToTrackPropertyCompareOperationOutputWithContext(ctx context.Context) TrackPropertyCompareOperationOutput {
	return o
}

func (o TrackPropertyCompareOperationOutput) ToTrackPropertyCompareOperationPtrOutput() TrackPropertyCompareOperationPtrOutput {
	return o.ToTrackPropertyCompareOperationPtrOutputWithContext(context.Background())
}

func (o TrackPropertyCompareOperationOutput) ToTrackPropertyCompareOperationPtrOutputWithContext(ctx context.Context) TrackPropertyCompareOperationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrackPropertyCompareOperation) *TrackPropertyCompareOperation {
		return &v
	}).(TrackPropertyCompareOperationPtrOutput)
}

func (o TrackPropertyCompareOperationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TrackPropertyCompareOperationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrackPropertyCompareOperation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TrackPropertyCompareOperationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrackPropertyCompareOperationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrackPropertyCompareOperation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TrackPropertyCompareOperationPtrOutput struct{ *pulumi.OutputState }

func (TrackPropertyCompareOperationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrackPropertyCompareOperation)(nil)).Elem()
}

func (o TrackPropertyCompareOperationPtrOutput) ToTrackPropertyCompareOperationPtrOutput() TrackPropertyCompareOperationPtrOutput {
	return o
}

func (o TrackPropertyCompareOperationPtrOutput) ToTrackPropertyCompareOperationPtrOutputWithContext(ctx context.Context) TrackPropertyCompareOperationPtrOutput {
	return o
}

func (o TrackPropertyCompareOperationPtrOutput) Elem() TrackPropertyCompareOperationOutput {
	return o.ApplyT(func(v *TrackPropertyCompareOperation) TrackPropertyCompareOperation {
		if v != nil {
			return *v
		}
		var ret TrackPropertyCompareOperation
		return ret
	}).(TrackPropertyCompareOperationOutput)
}

func (o TrackPropertyCompareOperationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrackPropertyCompareOperationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TrackPropertyCompareOperation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TrackPropertyCompareOperationInput interface {
	pulumi.Input

	ToTrackPropertyCompareOperationOutput() TrackPropertyCompareOperationOutput
	ToTrackPropertyCompareOperationOutputWithContext(context.Context) TrackPropertyCompareOperationOutput
}

var trackPropertyCompareOperationPtrType = reflect.TypeOf((**TrackPropertyCompareOperation)(nil)).Elem()

type TrackPropertyCompareOperationPtrInput interface {
	pulumi.Input

	ToTrackPropertyCompareOperationPtrOutput() TrackPropertyCompareOperationPtrOutput
	ToTrackPropertyCompareOperationPtrOutputWithContext(context.Context) TrackPropertyCompareOperationPtrOutput
}

type trackPropertyCompareOperationPtr string

func TrackPropertyCompareOperationPtr(v string) TrackPropertyCompareOperationPtrInput {
	return (*trackPropertyCompareOperationPtr)(&v)
}

func (*trackPropertyCompareOperationPtr) ElementType() reflect.Type {
	return trackPropertyCompareOperationPtrType
}

func (in *trackPropertyCompareOperationPtr) ToTrackPropertyCompareOperationPtrOutput() TrackPropertyCompareOperationPtrOutput {
	return pulumi.ToOutput(in).(TrackPropertyCompareOperationPtrOutput)
}

func (in *trackPropertyCompareOperationPtr) ToTrackPropertyCompareOperationPtrOutputWithContext(ctx context.Context) TrackPropertyCompareOperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TrackPropertyCompareOperationPtrOutput)
}

type TrackPropertyType string

const (
	// Unknown track property
	TrackPropertyTypeUnknown = TrackPropertyType("Unknown")
	// Track FourCC
	TrackPropertyTypeFourCC = TrackPropertyType("FourCC")
)

func (TrackPropertyType) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyType)(nil)).Elem()
}

func (e TrackPropertyType) ToTrackPropertyTypeOutput() TrackPropertyTypeOutput {
	return pulumi.ToOutput(e).(TrackPropertyTypeOutput)
}

func (e TrackPropertyType) ToTrackPropertyTypeOutputWithContext(ctx context.Context) TrackPropertyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TrackPropertyTypeOutput)
}

func (e TrackPropertyType) ToTrackPropertyTypePtrOutput() TrackPropertyTypePtrOutput {
	return e.ToTrackPropertyTypePtrOutputWithContext(context.Background())
}

func (e TrackPropertyType) ToTrackPropertyTypePtrOutputWithContext(ctx context.Context) TrackPropertyTypePtrOutput {
	return TrackPropertyType(e).ToTrackPropertyTypeOutputWithContext(ctx).ToTrackPropertyTypePtrOutputWithContext(ctx)
}

func (e TrackPropertyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrackPropertyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrackPropertyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TrackPropertyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TrackPropertyTypeOutput struct{ *pulumi.OutputState }

func (TrackPropertyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyType)(nil)).Elem()
}

func (o TrackPropertyTypeOutput) ToTrackPropertyTypeOutput() TrackPropertyTypeOutput {
	return o
}

func (o TrackPropertyTypeOutput) ToTrackPropertyTypeOutputWithContext(ctx context.Context) TrackPropertyTypeOutput {
	return o
}

func (o TrackPropertyTypeOutput) ToTrackPropertyTypePtrOutput() TrackPropertyTypePtrOutput {
	return o.ToTrackPropertyTypePtrOutputWithContext(context.Background())
}

func (o TrackPropertyTypeOutput) ToTrackPropertyTypePtrOutputWithContext(ctx context.Context) TrackPropertyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrackPropertyType) *TrackPropertyType {
		return &v
	}).(TrackPropertyTypePtrOutput)
}

func (o TrackPropertyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TrackPropertyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrackPropertyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TrackPropertyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrackPropertyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrackPropertyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TrackPropertyTypePtrOutput struct{ *pulumi.OutputState }

func (TrackPropertyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrackPropertyType)(nil)).Elem()
}

func (o TrackPropertyTypePtrOutput) ToTrackPropertyTypePtrOutput() TrackPropertyTypePtrOutput {
	return o
}

func (o TrackPropertyTypePtrOutput) ToTrackPropertyTypePtrOutputWithContext(ctx context.Context) TrackPropertyTypePtrOutput {
	return o
}

func (o TrackPropertyTypePtrOutput) Elem() TrackPropertyTypeOutput {
	return o.ApplyT(func(v *TrackPropertyType) TrackPropertyType {
		if v != nil {
			return *v
		}
		var ret TrackPropertyType
		return ret
	}).(TrackPropertyTypeOutput)
}

func (o TrackPropertyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrackPropertyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TrackPropertyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TrackPropertyTypeInput interface {
	pulumi.Input

	ToTrackPropertyTypeOutput() TrackPropertyTypeOutput
	ToTrackPropertyTypeOutputWithContext(context.Context) TrackPropertyTypeOutput
}

var trackPropertyTypePtrType = reflect.TypeOf((**TrackPropertyType)(nil)).Elem()

type TrackPropertyTypePtrInput interface {
	pulumi.Input

	ToTrackPropertyTypePtrOutput() TrackPropertyTypePtrOutput
	ToTrackPropertyTypePtrOutputWithContext(context.Context) TrackPropertyTypePtrOutput
}

type trackPropertyTypePtr string

func TrackPropertyTypePtr(v string) TrackPropertyTypePtrInput {
	return (*trackPropertyTypePtr)(&v)
}

func (*trackPropertyTypePtr) ElementType() reflect.Type {
	return trackPropertyTypePtrType
}

func (in *trackPropertyTypePtr) ToTrackPropertyTypePtrOutput() TrackPropertyTypePtrOutput {
	return pulumi.ToOutput(in).(TrackPropertyTypePtrOutput)
}

func (in *trackPropertyTypePtr) ToTrackPropertyTypePtrOutputWithContext(ctx context.Context) TrackPropertyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TrackPropertyTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AacAudioProfileInput)(nil)).Elem(), AacAudioProfile("AacLc"))
	pulumi.RegisterInputType(reflect.TypeOf((*AacAudioProfilePtrInput)(nil)).Elem(), AacAudioProfile("AacLc"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetContainerPermissionInput)(nil)).Elem(), AssetContainerPermission("Read"))
	pulumi.RegisterInputType(reflect.TypeOf((*AssetContainerPermissionPtrInput)(nil)).Elem(), AssetContainerPermission("Read"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeInput)(nil)).Elem(), ContentKeyPolicyFairPlayRentalAndLeaseKeyType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrInput)(nil)).Elem(), ContentKeyPolicyFairPlayRentalAndLeaseKeyType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyContentTypeInput)(nil)).Elem(), ContentKeyPolicyPlayReadyContentType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyContentTypePtrInput)(nil)).Elem(), ContentKeyPolicyPlayReadyContentType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseTypeInput)(nil)).Elem(), ContentKeyPolicyPlayReadyLicenseType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyLicenseTypePtrInput)(nil)).Elem(), ContentKeyPolicyPlayReadyLicenseType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyUnknownOutputPassingOptionInput)(nil)).Elem(), ContentKeyPolicyPlayReadyUnknownOutputPassingOption("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrInput)(nil)).Elem(), ContentKeyPolicyPlayReadyUnknownOutputPassingOption("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyRestrictionTokenTypeInput)(nil)).Elem(), ContentKeyPolicyRestrictionTokenType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContentKeyPolicyRestrictionTokenTypePtrInput)(nil)).Elem(), ContentKeyPolicyRestrictionTokenType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeinterlaceModeInput)(nil)).Elem(), DeinterlaceMode("Off"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeinterlaceModePtrInput)(nil)).Elem(), DeinterlaceMode("Off"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeinterlaceParityInput)(nil)).Elem(), DeinterlaceParity("Auto"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeinterlaceParityPtrInput)(nil)).Elem(), DeinterlaceParity("Auto"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderNamedPresetInput)(nil)).Elem(), EncoderNamedPreset("AdaptiveStreaming"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncoderNamedPresetPtrInput)(nil)).Elem(), EncoderNamedPreset("AdaptiveStreaming"))
	pulumi.RegisterInputType(reflect.TypeOf((*EntropyModeInput)(nil)).Elem(), EntropyMode("Cabac"))
	pulumi.RegisterInputType(reflect.TypeOf((*EntropyModePtrInput)(nil)).Elem(), EntropyMode("Cabac"))
	pulumi.RegisterInputType(reflect.TypeOf((*H264ComplexityInput)(nil)).Elem(), H264Complexity("Speed"))
	pulumi.RegisterInputType(reflect.TypeOf((*H264ComplexityPtrInput)(nil)).Elem(), H264Complexity("Speed"))
	pulumi.RegisterInputType(reflect.TypeOf((*H264VideoProfileInput)(nil)).Elem(), H264VideoProfile("Auto"))
	pulumi.RegisterInputType(reflect.TypeOf((*H264VideoProfilePtrInput)(nil)).Elem(), H264VideoProfile("Auto"))
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingTypeInput)(nil)).Elem(), LiveEventEncodingType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventEncodingTypePtrInput)(nil)).Elem(), LiveEventEncodingType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputProtocolInput)(nil)).Elem(), LiveEventInputProtocol("FragmentedMP4"))
	pulumi.RegisterInputType(reflect.TypeOf((*LiveEventInputProtocolPtrInput)(nil)).Elem(), LiveEventInputProtocol("FragmentedMP4"))
	pulumi.RegisterInputType(reflect.TypeOf((*OnErrorTypeInput)(nil)).Elem(), OnErrorType("StopProcessingJob"))
	pulumi.RegisterInputType(reflect.TypeOf((*OnErrorTypePtrInput)(nil)).Elem(), OnErrorType("StopProcessingJob"))
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityInput)(nil)).Elem(), Priority("Low"))
	pulumi.RegisterInputType(reflect.TypeOf((*PriorityPtrInput)(nil)).Elem(), Priority("Low"))
	pulumi.RegisterInputType(reflect.TypeOf((*RotationInput)(nil)).Elem(), Rotation("Auto"))
	pulumi.RegisterInputType(reflect.TypeOf((*RotationPtrInput)(nil)).Elem(), Rotation("Auto"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountTypeInput)(nil)).Elem(), StorageAccountType("Primary"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountTypePtrInput)(nil)).Elem(), StorageAccountType("Primary"))
	pulumi.RegisterInputType(reflect.TypeOf((*StreamOptionsFlagInput)(nil)).Elem(), StreamOptionsFlag("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*StreamOptionsFlagPtrInput)(nil)).Elem(), StreamOptionsFlag("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*StreamOptionsFlagArrayInput)(nil)).Elem(), StreamOptionsFlagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StretchModeInput)(nil)).Elem(), StretchMode("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*StretchModePtrInput)(nil)).Elem(), StretchMode("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrackPropertyCompareOperationInput)(nil)).Elem(), TrackPropertyCompareOperation("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrackPropertyCompareOperationPtrInput)(nil)).Elem(), TrackPropertyCompareOperation("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrackPropertyTypeInput)(nil)).Elem(), TrackPropertyType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrackPropertyTypePtrInput)(nil)).Elem(), TrackPropertyType("Unknown"))
	pulumi.RegisterOutputType(AacAudioProfileOutput{})
	pulumi.RegisterOutputType(AacAudioProfilePtrOutput{})
	pulumi.RegisterOutputType(AssetContainerPermissionOutput{})
	pulumi.RegisterOutputType(AssetContainerPermissionPtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyContentTypeOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyContentTypePtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyLicenseTypeOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionPtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyRestrictionTokenTypeOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyRestrictionTokenTypePtrOutput{})
	pulumi.RegisterOutputType(DeinterlaceModeOutput{})
	pulumi.RegisterOutputType(DeinterlaceModePtrOutput{})
	pulumi.RegisterOutputType(DeinterlaceParityOutput{})
	pulumi.RegisterOutputType(DeinterlaceParityPtrOutput{})
	pulumi.RegisterOutputType(EncoderNamedPresetOutput{})
	pulumi.RegisterOutputType(EncoderNamedPresetPtrOutput{})
	pulumi.RegisterOutputType(EntropyModeOutput{})
	pulumi.RegisterOutputType(EntropyModePtrOutput{})
	pulumi.RegisterOutputType(H264ComplexityOutput{})
	pulumi.RegisterOutputType(H264ComplexityPtrOutput{})
	pulumi.RegisterOutputType(H264VideoProfileOutput{})
	pulumi.RegisterOutputType(H264VideoProfilePtrOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingTypeOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingTypePtrOutput{})
	pulumi.RegisterOutputType(LiveEventInputProtocolOutput{})
	pulumi.RegisterOutputType(LiveEventInputProtocolPtrOutput{})
	pulumi.RegisterOutputType(OnErrorTypeOutput{})
	pulumi.RegisterOutputType(OnErrorTypePtrOutput{})
	pulumi.RegisterOutputType(PriorityOutput{})
	pulumi.RegisterOutputType(PriorityPtrOutput{})
	pulumi.RegisterOutputType(RotationOutput{})
	pulumi.RegisterOutputType(RotationPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountTypeOutput{})
	pulumi.RegisterOutputType(StorageAccountTypePtrOutput{})
	pulumi.RegisterOutputType(StreamOptionsFlagOutput{})
	pulumi.RegisterOutputType(StreamOptionsFlagPtrOutput{})
	pulumi.RegisterOutputType(StreamOptionsFlagArrayOutput{})
	pulumi.RegisterOutputType(StretchModeOutput{})
	pulumi.RegisterOutputType(StretchModePtrOutput{})
	pulumi.RegisterOutputType(TrackPropertyCompareOperationOutput{})
	pulumi.RegisterOutputType(TrackPropertyCompareOperationPtrOutput{})
	pulumi.RegisterOutputType(TrackPropertyTypeOutput{})
	pulumi.RegisterOutputType(TrackPropertyTypePtrOutput{})
}
