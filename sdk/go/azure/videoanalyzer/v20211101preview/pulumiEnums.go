


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessPolicyEccAlgo string

const (
	// ES265
	AccessPolicyEccAlgoES256 = AccessPolicyEccAlgo("ES256")
	// ES384
	AccessPolicyEccAlgoES384 = AccessPolicyEccAlgo("ES384")
	// ES512
	AccessPolicyEccAlgoES512 = AccessPolicyEccAlgo("ES512")
)

func (AccessPolicyEccAlgo) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyEccAlgo)(nil)).Elem()
}

func (e AccessPolicyEccAlgo) ToAccessPolicyEccAlgoOutput() AccessPolicyEccAlgoOutput {
	return pulumi.ToOutput(e).(AccessPolicyEccAlgoOutput)
}

func (e AccessPolicyEccAlgo) ToAccessPolicyEccAlgoOutputWithContext(ctx context.Context) AccessPolicyEccAlgoOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessPolicyEccAlgoOutput)
}

func (e AccessPolicyEccAlgo) ToAccessPolicyEccAlgoPtrOutput() AccessPolicyEccAlgoPtrOutput {
	return e.ToAccessPolicyEccAlgoPtrOutputWithContext(context.Background())
}

func (e AccessPolicyEccAlgo) ToAccessPolicyEccAlgoPtrOutputWithContext(ctx context.Context) AccessPolicyEccAlgoPtrOutput {
	return AccessPolicyEccAlgo(e).ToAccessPolicyEccAlgoOutputWithContext(ctx).ToAccessPolicyEccAlgoPtrOutputWithContext(ctx)
}

func (e AccessPolicyEccAlgo) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessPolicyEccAlgo) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessPolicyEccAlgo) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessPolicyEccAlgo) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessPolicyEccAlgoOutput struct{ *pulumi.OutputState }

func (AccessPolicyEccAlgoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyEccAlgo)(nil)).Elem()
}

func (o AccessPolicyEccAlgoOutput) ToAccessPolicyEccAlgoOutput() AccessPolicyEccAlgoOutput {
	return o
}

func (o AccessPolicyEccAlgoOutput) ToAccessPolicyEccAlgoOutputWithContext(ctx context.Context) AccessPolicyEccAlgoOutput {
	return o
}

func (o AccessPolicyEccAlgoOutput) ToAccessPolicyEccAlgoPtrOutput() AccessPolicyEccAlgoPtrOutput {
	return o.ToAccessPolicyEccAlgoPtrOutputWithContext(context.Background())
}

func (o AccessPolicyEccAlgoOutput) ToAccessPolicyEccAlgoPtrOutputWithContext(ctx context.Context) AccessPolicyEccAlgoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessPolicyEccAlgo) *AccessPolicyEccAlgo {
		return &v
	}).(AccessPolicyEccAlgoPtrOutput)
}

func (o AccessPolicyEccAlgoOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessPolicyEccAlgoOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessPolicyEccAlgo) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessPolicyEccAlgoOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPolicyEccAlgoOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessPolicyEccAlgo) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessPolicyEccAlgoPtrOutput struct{ *pulumi.OutputState }

func (AccessPolicyEccAlgoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicyEccAlgo)(nil)).Elem()
}

func (o AccessPolicyEccAlgoPtrOutput) ToAccessPolicyEccAlgoPtrOutput() AccessPolicyEccAlgoPtrOutput {
	return o
}

func (o AccessPolicyEccAlgoPtrOutput) ToAccessPolicyEccAlgoPtrOutputWithContext(ctx context.Context) AccessPolicyEccAlgoPtrOutput {
	return o
}

func (o AccessPolicyEccAlgoPtrOutput) Elem() AccessPolicyEccAlgoOutput {
	return o.ApplyT(func(v *AccessPolicyEccAlgo) AccessPolicyEccAlgo {
		if v != nil {
			return *v
		}
		var ret AccessPolicyEccAlgo
		return ret
	}).(AccessPolicyEccAlgoOutput)
}

func (o AccessPolicyEccAlgoPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPolicyEccAlgoPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessPolicyEccAlgo) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessPolicyEccAlgoInput interface {
	pulumi.Input

	ToAccessPolicyEccAlgoOutput() AccessPolicyEccAlgoOutput
	ToAccessPolicyEccAlgoOutputWithContext(context.Context) AccessPolicyEccAlgoOutput
}

var accessPolicyEccAlgoPtrType = reflect.TypeOf((**AccessPolicyEccAlgo)(nil)).Elem()

type AccessPolicyEccAlgoPtrInput interface {
	pulumi.Input

	ToAccessPolicyEccAlgoPtrOutput() AccessPolicyEccAlgoPtrOutput
	ToAccessPolicyEccAlgoPtrOutputWithContext(context.Context) AccessPolicyEccAlgoPtrOutput
}

type accessPolicyEccAlgoPtr string

func AccessPolicyEccAlgoPtr(v string) AccessPolicyEccAlgoPtrInput {
	return (*accessPolicyEccAlgoPtr)(&v)
}

func (*accessPolicyEccAlgoPtr) ElementType() reflect.Type {
	return accessPolicyEccAlgoPtrType
}

func (in *accessPolicyEccAlgoPtr) ToAccessPolicyEccAlgoPtrOutput() AccessPolicyEccAlgoPtrOutput {
	return pulumi.ToOutput(in).(AccessPolicyEccAlgoPtrOutput)
}

func (in *accessPolicyEccAlgoPtr) ToAccessPolicyEccAlgoPtrOutputWithContext(ctx context.Context) AccessPolicyEccAlgoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessPolicyEccAlgoPtrOutput)
}

type AccessPolicyRole string

const (
	// Reader role allows for read-only operations to be performed through the client APIs.
	AccessPolicyRoleReader = AccessPolicyRole("Reader")
)

func (AccessPolicyRole) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRole)(nil)).Elem()
}

func (e AccessPolicyRole) ToAccessPolicyRoleOutput() AccessPolicyRoleOutput {
	return pulumi.ToOutput(e).(AccessPolicyRoleOutput)
}

func (e AccessPolicyRole) ToAccessPolicyRoleOutputWithContext(ctx context.Context) AccessPolicyRoleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessPolicyRoleOutput)
}

func (e AccessPolicyRole) ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput {
	return e.ToAccessPolicyRolePtrOutputWithContext(context.Background())
}

func (e AccessPolicyRole) ToAccessPolicyRolePtrOutputWithContext(ctx context.Context) AccessPolicyRolePtrOutput {
	return AccessPolicyRole(e).ToAccessPolicyRoleOutputWithContext(ctx).ToAccessPolicyRolePtrOutputWithContext(ctx)
}

func (e AccessPolicyRole) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessPolicyRole) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessPolicyRole) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessPolicyRole) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessPolicyRoleOutput struct{ *pulumi.OutputState }

func (AccessPolicyRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRole)(nil)).Elem()
}

func (o AccessPolicyRoleOutput) ToAccessPolicyRoleOutput() AccessPolicyRoleOutput {
	return o
}

func (o AccessPolicyRoleOutput) ToAccessPolicyRoleOutputWithContext(ctx context.Context) AccessPolicyRoleOutput {
	return o
}

func (o AccessPolicyRoleOutput) ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput {
	return o.ToAccessPolicyRolePtrOutputWithContext(context.Background())
}

func (o AccessPolicyRoleOutput) ToAccessPolicyRolePtrOutputWithContext(ctx context.Context) AccessPolicyRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessPolicyRole) *AccessPolicyRole {
		return &v
	}).(AccessPolicyRolePtrOutput)
}

func (o AccessPolicyRoleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessPolicyRoleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessPolicyRole) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessPolicyRoleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPolicyRoleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessPolicyRole) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessPolicyRolePtrOutput struct{ *pulumi.OutputState }

func (AccessPolicyRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicyRole)(nil)).Elem()
}

func (o AccessPolicyRolePtrOutput) ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput {
	return o
}

func (o AccessPolicyRolePtrOutput) ToAccessPolicyRolePtrOutputWithContext(ctx context.Context) AccessPolicyRolePtrOutput {
	return o
}

func (o AccessPolicyRolePtrOutput) Elem() AccessPolicyRoleOutput {
	return o.ApplyT(func(v *AccessPolicyRole) AccessPolicyRole {
		if v != nil {
			return *v
		}
		var ret AccessPolicyRole
		return ret
	}).(AccessPolicyRoleOutput)
}

func (o AccessPolicyRolePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPolicyRolePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessPolicyRole) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessPolicyRoleInput interface {
	pulumi.Input

	ToAccessPolicyRoleOutput() AccessPolicyRoleOutput
	ToAccessPolicyRoleOutputWithContext(context.Context) AccessPolicyRoleOutput
}

var accessPolicyRolePtrType = reflect.TypeOf((**AccessPolicyRole)(nil)).Elem()

type AccessPolicyRolePtrInput interface {
	pulumi.Input

	ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput
	ToAccessPolicyRolePtrOutputWithContext(context.Context) AccessPolicyRolePtrOutput
}

type accessPolicyRolePtr string

func AccessPolicyRolePtr(v string) AccessPolicyRolePtrInput {
	return (*accessPolicyRolePtr)(&v)
}

func (*accessPolicyRolePtr) ElementType() reflect.Type {
	return accessPolicyRolePtrType
}

func (in *accessPolicyRolePtr) ToAccessPolicyRolePtrOutput() AccessPolicyRolePtrOutput {
	return pulumi.ToOutput(in).(AccessPolicyRolePtrOutput)
}

func (in *accessPolicyRolePtr) ToAccessPolicyRolePtrOutputWithContext(ctx context.Context) AccessPolicyRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessPolicyRolePtrOutput)
}

type AccessPolicyRsaAlgo string

const (
	// RS256
	AccessPolicyRsaAlgoRS256 = AccessPolicyRsaAlgo("RS256")
	// RS384
	AccessPolicyRsaAlgoRS384 = AccessPolicyRsaAlgo("RS384")
	// RS512
	AccessPolicyRsaAlgoRS512 = AccessPolicyRsaAlgo("RS512")
)

func (AccessPolicyRsaAlgo) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRsaAlgo)(nil)).Elem()
}

func (e AccessPolicyRsaAlgo) ToAccessPolicyRsaAlgoOutput() AccessPolicyRsaAlgoOutput {
	return pulumi.ToOutput(e).(AccessPolicyRsaAlgoOutput)
}

func (e AccessPolicyRsaAlgo) ToAccessPolicyRsaAlgoOutputWithContext(ctx context.Context) AccessPolicyRsaAlgoOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessPolicyRsaAlgoOutput)
}

func (e AccessPolicyRsaAlgo) ToAccessPolicyRsaAlgoPtrOutput() AccessPolicyRsaAlgoPtrOutput {
	return e.ToAccessPolicyRsaAlgoPtrOutputWithContext(context.Background())
}

func (e AccessPolicyRsaAlgo) ToAccessPolicyRsaAlgoPtrOutputWithContext(ctx context.Context) AccessPolicyRsaAlgoPtrOutput {
	return AccessPolicyRsaAlgo(e).ToAccessPolicyRsaAlgoOutputWithContext(ctx).ToAccessPolicyRsaAlgoPtrOutputWithContext(ctx)
}

func (e AccessPolicyRsaAlgo) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessPolicyRsaAlgo) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessPolicyRsaAlgo) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessPolicyRsaAlgo) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessPolicyRsaAlgoOutput struct{ *pulumi.OutputState }

func (AccessPolicyRsaAlgoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyRsaAlgo)(nil)).Elem()
}

func (o AccessPolicyRsaAlgoOutput) ToAccessPolicyRsaAlgoOutput() AccessPolicyRsaAlgoOutput {
	return o
}

func (o AccessPolicyRsaAlgoOutput) ToAccessPolicyRsaAlgoOutputWithContext(ctx context.Context) AccessPolicyRsaAlgoOutput {
	return o
}

func (o AccessPolicyRsaAlgoOutput) ToAccessPolicyRsaAlgoPtrOutput() AccessPolicyRsaAlgoPtrOutput {
	return o.ToAccessPolicyRsaAlgoPtrOutputWithContext(context.Background())
}

func (o AccessPolicyRsaAlgoOutput) ToAccessPolicyRsaAlgoPtrOutputWithContext(ctx context.Context) AccessPolicyRsaAlgoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessPolicyRsaAlgo) *AccessPolicyRsaAlgo {
		return &v
	}).(AccessPolicyRsaAlgoPtrOutput)
}

func (o AccessPolicyRsaAlgoOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessPolicyRsaAlgoOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessPolicyRsaAlgo) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessPolicyRsaAlgoOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPolicyRsaAlgoOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessPolicyRsaAlgo) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessPolicyRsaAlgoPtrOutput struct{ *pulumi.OutputState }

func (AccessPolicyRsaAlgoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPolicyRsaAlgo)(nil)).Elem()
}

func (o AccessPolicyRsaAlgoPtrOutput) ToAccessPolicyRsaAlgoPtrOutput() AccessPolicyRsaAlgoPtrOutput {
	return o
}

func (o AccessPolicyRsaAlgoPtrOutput) ToAccessPolicyRsaAlgoPtrOutputWithContext(ctx context.Context) AccessPolicyRsaAlgoPtrOutput {
	return o
}

func (o AccessPolicyRsaAlgoPtrOutput) Elem() AccessPolicyRsaAlgoOutput {
	return o.ApplyT(func(v *AccessPolicyRsaAlgo) AccessPolicyRsaAlgo {
		if v != nil {
			return *v
		}
		var ret AccessPolicyRsaAlgo
		return ret
	}).(AccessPolicyRsaAlgoOutput)
}

func (o AccessPolicyRsaAlgoPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPolicyRsaAlgoPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessPolicyRsaAlgo) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessPolicyRsaAlgoInput interface {
	pulumi.Input

	ToAccessPolicyRsaAlgoOutput() AccessPolicyRsaAlgoOutput
	ToAccessPolicyRsaAlgoOutputWithContext(context.Context) AccessPolicyRsaAlgoOutput
}

var accessPolicyRsaAlgoPtrType = reflect.TypeOf((**AccessPolicyRsaAlgo)(nil)).Elem()

type AccessPolicyRsaAlgoPtrInput interface {
	pulumi.Input

	ToAccessPolicyRsaAlgoPtrOutput() AccessPolicyRsaAlgoPtrOutput
	ToAccessPolicyRsaAlgoPtrOutputWithContext(context.Context) AccessPolicyRsaAlgoPtrOutput
}

type accessPolicyRsaAlgoPtr string

func AccessPolicyRsaAlgoPtr(v string) AccessPolicyRsaAlgoPtrInput {
	return (*accessPolicyRsaAlgoPtr)(&v)
}

func (*accessPolicyRsaAlgoPtr) ElementType() reflect.Type {
	return accessPolicyRsaAlgoPtrType
}

func (in *accessPolicyRsaAlgoPtr) ToAccessPolicyRsaAlgoPtrOutput() AccessPolicyRsaAlgoPtrOutput {
	return pulumi.ToOutput(in).(AccessPolicyRsaAlgoPtrOutput)
}

func (in *accessPolicyRsaAlgoPtr) ToAccessPolicyRsaAlgoPtrOutputWithContext(ctx context.Context) AccessPolicyRsaAlgoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessPolicyRsaAlgoPtrOutput)
}

type AccountEncryptionKeyType string

const (
	// The Account Key is encrypted with a System Key.
	AccountEncryptionKeyTypeSystemKey = AccountEncryptionKeyType("SystemKey")
	// The Account Key is encrypted with a Customer Key.
	AccountEncryptionKeyTypeCustomerKey = AccountEncryptionKeyType("CustomerKey")
)

func (AccountEncryptionKeyType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryptionKeyType)(nil)).Elem()
}

func (e AccountEncryptionKeyType) ToAccountEncryptionKeyTypeOutput() AccountEncryptionKeyTypeOutput {
	return pulumi.ToOutput(e).(AccountEncryptionKeyTypeOutput)
}

func (e AccountEncryptionKeyType) ToAccountEncryptionKeyTypeOutputWithContext(ctx context.Context) AccountEncryptionKeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccountEncryptionKeyTypeOutput)
}

func (e AccountEncryptionKeyType) ToAccountEncryptionKeyTypePtrOutput() AccountEncryptionKeyTypePtrOutput {
	return e.ToAccountEncryptionKeyTypePtrOutputWithContext(context.Background())
}

func (e AccountEncryptionKeyType) ToAccountEncryptionKeyTypePtrOutputWithContext(ctx context.Context) AccountEncryptionKeyTypePtrOutput {
	return AccountEncryptionKeyType(e).ToAccountEncryptionKeyTypeOutputWithContext(ctx).ToAccountEncryptionKeyTypePtrOutputWithContext(ctx)
}

func (e AccountEncryptionKeyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccountEncryptionKeyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccountEncryptionKeyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccountEncryptionKeyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccountEncryptionKeyTypeOutput struct{ *pulumi.OutputState }

func (AccountEncryptionKeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountEncryptionKeyType)(nil)).Elem()
}

func (o AccountEncryptionKeyTypeOutput) ToAccountEncryptionKeyTypeOutput() AccountEncryptionKeyTypeOutput {
	return o
}

func (o AccountEncryptionKeyTypeOutput) ToAccountEncryptionKeyTypeOutputWithContext(ctx context.Context) AccountEncryptionKeyTypeOutput {
	return o
}

func (o AccountEncryptionKeyTypeOutput) ToAccountEncryptionKeyTypePtrOutput() AccountEncryptionKeyTypePtrOutput {
	return o.ToAccountEncryptionKeyTypePtrOutputWithContext(context.Background())
}

func (o AccountEncryptionKeyTypeOutput) ToAccountEncryptionKeyTypePtrOutputWithContext(ctx context.Context) AccountEncryptionKeyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccountEncryptionKeyType) *AccountEncryptionKeyType {
		return &v
	}).(AccountEncryptionKeyTypePtrOutput)
}

func (o AccountEncryptionKeyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccountEncryptionKeyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccountEncryptionKeyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccountEncryptionKeyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccountEncryptionKeyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccountEncryptionKeyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccountEncryptionKeyTypePtrOutput struct{ *pulumi.OutputState }

func (AccountEncryptionKeyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountEncryptionKeyType)(nil)).Elem()
}

func (o AccountEncryptionKeyTypePtrOutput) ToAccountEncryptionKeyTypePtrOutput() AccountEncryptionKeyTypePtrOutput {
	return o
}

func (o AccountEncryptionKeyTypePtrOutput) ToAccountEncryptionKeyTypePtrOutputWithContext(ctx context.Context) AccountEncryptionKeyTypePtrOutput {
	return o
}

func (o AccountEncryptionKeyTypePtrOutput) Elem() AccountEncryptionKeyTypeOutput {
	return o.ApplyT(func(v *AccountEncryptionKeyType) AccountEncryptionKeyType {
		if v != nil {
			return *v
		}
		var ret AccountEncryptionKeyType
		return ret
	}).(AccountEncryptionKeyTypeOutput)
}

func (o AccountEncryptionKeyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccountEncryptionKeyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccountEncryptionKeyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccountEncryptionKeyTypeInput interface {
	pulumi.Input

	ToAccountEncryptionKeyTypeOutput() AccountEncryptionKeyTypeOutput
	ToAccountEncryptionKeyTypeOutputWithContext(context.Context) AccountEncryptionKeyTypeOutput
}

var accountEncryptionKeyTypePtrType = reflect.TypeOf((**AccountEncryptionKeyType)(nil)).Elem()

type AccountEncryptionKeyTypePtrInput interface {
	pulumi.Input

	ToAccountEncryptionKeyTypePtrOutput() AccountEncryptionKeyTypePtrOutput
	ToAccountEncryptionKeyTypePtrOutputWithContext(context.Context) AccountEncryptionKeyTypePtrOutput
}

type accountEncryptionKeyTypePtr string

func AccountEncryptionKeyTypePtr(v string) AccountEncryptionKeyTypePtrInput {
	return (*accountEncryptionKeyTypePtr)(&v)
}

func (*accountEncryptionKeyTypePtr) ElementType() reflect.Type {
	return accountEncryptionKeyTypePtrType
}

func (in *accountEncryptionKeyTypePtr) ToAccountEncryptionKeyTypePtrOutput() AccountEncryptionKeyTypePtrOutput {
	return pulumi.ToOutput(in).(AccountEncryptionKeyTypePtrOutput)
}

func (in *accountEncryptionKeyTypePtr) ToAccountEncryptionKeyTypePtrOutputWithContext(ctx context.Context) AccountEncryptionKeyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccountEncryptionKeyTypePtrOutput)
}

type EncoderSystemPresetType string

const (
	// Produces an MP4 file where the video is encoded with H.264 codec at a picture height of 540 pixels, and at a maximum bitrate of 2000 Kbps. Encoded video has the same average frame rate as the input. The aspect ratio of the input is preserved. If the input content has audio, then it is encoded with AAC-LC codec at 96 Kbps
	EncoderSystemPresetType_SingleLayer_540p_H264_AAC = EncoderSystemPresetType("SingleLayer_540p_H264_AAC")
	// Produces an MP4 file where the video is encoded with H.264 codec at a picture height of 720 pixels, and at a maximum bitrate of 3500 Kbps. Encoded video has the same average frame rate as the input. The aspect ratio of the input is preserved. If the input content has audio, then it is encoded with AAC-LC codec at 96 Kbps
	EncoderSystemPresetType_SingleLayer_720p_H264_AAC = EncoderSystemPresetType("SingleLayer_720p_H264_AAC")
	// Produces an MP4 file where the video is encoded with H.264 codec at a picture height of 1080 pixels, and at a maximum bitrate of 6000 Kbps. Encoded video has the same average frame rate as the input. The aspect ratio of the input is preserved. If the input content has audio, then it is encoded with AAC-LC codec at 128 Kbps
	EncoderSystemPresetType_SingleLayer_1080p_H264_AAC = EncoderSystemPresetType("SingleLayer_1080p_H264_AAC")
	// Produces an MP4 file where the video is encoded with H.264 codec at a picture height of 2160 pixels, and at a maximum bitrate of 16000 Kbps. Encoded video has the same average frame rate as the input. The aspect ratio of the input is preserved. If the input content has audio, then it is encoded with AAC-LC codec at 128 Kbps
	EncoderSystemPresetType_SingleLayer_2160p_H264_AAC = EncoderSystemPresetType("SingleLayer_2160p_H264_AAC")
)

func (EncoderSystemPresetType) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderSystemPresetType)(nil)).Elem()
}

func (e EncoderSystemPresetType) ToEncoderSystemPresetTypeOutput() EncoderSystemPresetTypeOutput {
	return pulumi.ToOutput(e).(EncoderSystemPresetTypeOutput)
}

func (e EncoderSystemPresetType) ToEncoderSystemPresetTypeOutputWithContext(ctx context.Context) EncoderSystemPresetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncoderSystemPresetTypeOutput)
}

func (e EncoderSystemPresetType) ToEncoderSystemPresetTypePtrOutput() EncoderSystemPresetTypePtrOutput {
	return e.ToEncoderSystemPresetTypePtrOutputWithContext(context.Background())
}

func (e EncoderSystemPresetType) ToEncoderSystemPresetTypePtrOutputWithContext(ctx context.Context) EncoderSystemPresetTypePtrOutput {
	return EncoderSystemPresetType(e).ToEncoderSystemPresetTypeOutputWithContext(ctx).ToEncoderSystemPresetTypePtrOutputWithContext(ctx)
}

func (e EncoderSystemPresetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncoderSystemPresetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncoderSystemPresetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncoderSystemPresetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncoderSystemPresetTypeOutput struct{ *pulumi.OutputState }

func (EncoderSystemPresetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncoderSystemPresetType)(nil)).Elem()
}

func (o EncoderSystemPresetTypeOutput) ToEncoderSystemPresetTypeOutput() EncoderSystemPresetTypeOutput {
	return o
}

func (o EncoderSystemPresetTypeOutput) ToEncoderSystemPresetTypeOutputWithContext(ctx context.Context) EncoderSystemPresetTypeOutput {
	return o
}

func (o EncoderSystemPresetTypeOutput) ToEncoderSystemPresetTypePtrOutput() EncoderSystemPresetTypePtrOutput {
	return o.ToEncoderSystemPresetTypePtrOutputWithContext(context.Background())
}

func (o EncoderSystemPresetTypeOutput) ToEncoderSystemPresetTypePtrOutputWithContext(ctx context.Context) EncoderSystemPresetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncoderSystemPresetType) *EncoderSystemPresetType {
		return &v
	}).(EncoderSystemPresetTypePtrOutput)
}

func (o EncoderSystemPresetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncoderSystemPresetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncoderSystemPresetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncoderSystemPresetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncoderSystemPresetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncoderSystemPresetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncoderSystemPresetTypePtrOutput struct{ *pulumi.OutputState }

func (EncoderSystemPresetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncoderSystemPresetType)(nil)).Elem()
}

func (o EncoderSystemPresetTypePtrOutput) ToEncoderSystemPresetTypePtrOutput() EncoderSystemPresetTypePtrOutput {
	return o
}

func (o EncoderSystemPresetTypePtrOutput) ToEncoderSystemPresetTypePtrOutputWithContext(ctx context.Context) EncoderSystemPresetTypePtrOutput {
	return o
}

func (o EncoderSystemPresetTypePtrOutput) Elem() EncoderSystemPresetTypeOutput {
	return o.ApplyT(func(v *EncoderSystemPresetType) EncoderSystemPresetType {
		if v != nil {
			return *v
		}
		var ret EncoderSystemPresetType
		return ret
	}).(EncoderSystemPresetTypeOutput)
}

func (o EncoderSystemPresetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncoderSystemPresetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncoderSystemPresetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncoderSystemPresetTypeInput interface {
	pulumi.Input

	ToEncoderSystemPresetTypeOutput() EncoderSystemPresetTypeOutput
	ToEncoderSystemPresetTypeOutputWithContext(context.Context) EncoderSystemPresetTypeOutput
}

var encoderSystemPresetTypePtrType = reflect.TypeOf((**EncoderSystemPresetType)(nil)).Elem()

type EncoderSystemPresetTypePtrInput interface {
	pulumi.Input

	ToEncoderSystemPresetTypePtrOutput() EncoderSystemPresetTypePtrOutput
	ToEncoderSystemPresetTypePtrOutputWithContext(context.Context) EncoderSystemPresetTypePtrOutput
}

type encoderSystemPresetTypePtr string

func EncoderSystemPresetTypePtr(v string) EncoderSystemPresetTypePtrInput {
	return (*encoderSystemPresetTypePtr)(&v)
}

func (*encoderSystemPresetTypePtr) ElementType() reflect.Type {
	return encoderSystemPresetTypePtrType
}

func (in *encoderSystemPresetTypePtr) ToEncoderSystemPresetTypePtrOutput() EncoderSystemPresetTypePtrOutput {
	return pulumi.ToOutput(in).(EncoderSystemPresetTypePtrOutput)
}

func (in *encoderSystemPresetTypePtr) ToEncoderSystemPresetTypePtrOutputWithContext(ctx context.Context) EncoderSystemPresetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncoderSystemPresetTypePtrOutput)
}

type Kind string

const (
	// Live pipeline topology resource.
	KindLive = Kind("Live")
	// Batch pipeline topology resource.
	KindBatch = Kind("Batch")
)

func (Kind) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (e Kind) ToKindOutput() KindOutput {
	return pulumi.ToOutput(e).(KindOutput)
}

func (e Kind) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KindOutput)
}

func (e Kind) ToKindPtrOutput() KindPtrOutput {
	return e.ToKindPtrOutputWithContext(context.Background())
}

func (e Kind) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return Kind(e).ToKindOutputWithContext(ctx).ToKindPtrOutputWithContext(ctx)
}

func (e Kind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Kind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KindOutput struct{ *pulumi.OutputState }

func (KindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (o KindOutput) ToKindOutput() KindOutput {
	return o
}

func (o KindOutput) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return o
}

func (o KindOutput) ToKindPtrOutput() KindPtrOutput {
	return o.ToKindPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Kind) *Kind {
		return &v
	}).(KindPtrOutput)
}

func (o KindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KindPtrOutput struct{ *pulumi.OutputState }

func (KindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind)(nil)).Elem()
}

func (o KindPtrOutput) ToKindPtrOutput() KindPtrOutput {
	return o
}

func (o KindPtrOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o
}

func (o KindPtrOutput) Elem() KindOutput {
	return o.ApplyT(func(v *Kind) Kind {
		if v != nil {
			return *v
		}
		var ret Kind
		return ret
	}).(KindOutput)
}

func (o KindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Kind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KindInput interface {
	pulumi.Input

	ToKindOutput() KindOutput
	ToKindOutputWithContext(context.Context) KindOutput
}

var kindPtrType = reflect.TypeOf((**Kind)(nil)).Elem()

type KindPtrInput interface {
	pulumi.Input

	ToKindPtrOutput() KindPtrOutput
	ToKindPtrOutputWithContext(context.Context) KindPtrOutput
}

type kindPtr string

func KindPtr(v string) KindPtrInput {
	return (*kindPtr)(&v)
}

func (*kindPtr) ElementType() reflect.Type {
	return kindPtrType
}

func (in *kindPtr) ToKindPtrOutput() KindPtrOutput {
	return pulumi.ToOutput(in).(KindPtrOutput)
}

func (in *kindPtr) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KindPtrOutput)
}

type ParameterType string

const (
	// The parameter's value is a string.
	ParameterTypeString = ParameterType("String")
	// The parameter's value is a string that holds sensitive information.
	ParameterTypeSecretString = ParameterType("SecretString")
	// The parameter's value is a 32-bit signed integer.
	ParameterTypeInt = ParameterType("Int")
	// The parameter's value is a 64-bit double-precision floating point.
	ParameterTypeDouble = ParameterType("Double")
	// The parameter's value is a boolean value that is either true or false.
	ParameterTypeBool = ParameterType("Bool")
)

func (ParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterType)(nil)).Elem()
}

func (e ParameterType) ToParameterTypeOutput() ParameterTypeOutput {
	return pulumi.ToOutput(e).(ParameterTypeOutput)
}

func (e ParameterType) ToParameterTypeOutputWithContext(ctx context.Context) ParameterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ParameterTypeOutput)
}

func (e ParameterType) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return e.ToParameterTypePtrOutputWithContext(context.Background())
}

func (e ParameterType) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return ParameterType(e).ToParameterTypeOutputWithContext(ctx).ToParameterTypePtrOutputWithContext(ctx)
}

func (e ParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ParameterTypeOutput struct{ *pulumi.OutputState }

func (ParameterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterType)(nil)).Elem()
}

func (o ParameterTypeOutput) ToParameterTypeOutput() ParameterTypeOutput {
	return o
}

func (o ParameterTypeOutput) ToParameterTypeOutputWithContext(ctx context.Context) ParameterTypeOutput {
	return o
}

func (o ParameterTypeOutput) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return o.ToParameterTypePtrOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParameterType) *ParameterType {
		return &v
	}).(ParameterTypePtrOutput)
}

func (o ParameterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ParameterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ParameterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ParameterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ParameterTypePtrOutput struct{ *pulumi.OutputState }

func (ParameterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterType)(nil)).Elem()
}

func (o ParameterTypePtrOutput) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return o
}

func (o ParameterTypePtrOutput) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return o
}

func (o ParameterTypePtrOutput) Elem() ParameterTypeOutput {
	return o.ApplyT(func(v *ParameterType) ParameterType {
		if v != nil {
			return *v
		}
		var ret ParameterType
		return ret
	}).(ParameterTypeOutput)
}

func (o ParameterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ParameterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ParameterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ParameterTypeInput interface {
	pulumi.Input

	ToParameterTypeOutput() ParameterTypeOutput
	ToParameterTypeOutputWithContext(context.Context) ParameterTypeOutput
}

var parameterTypePtrType = reflect.TypeOf((**ParameterType)(nil)).Elem()

type ParameterTypePtrInput interface {
	pulumi.Input

	ToParameterTypePtrOutput() ParameterTypePtrOutput
	ToParameterTypePtrOutputWithContext(context.Context) ParameterTypePtrOutput
}

type parameterTypePtr string

func ParameterTypePtr(v string) ParameterTypePtrInput {
	return (*parameterTypePtr)(&v)
}

func (*parameterTypePtr) ElementType() reflect.Type {
	return parameterTypePtrType
}

func (in *parameterTypePtr) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return pulumi.ToOutput(in).(ParameterTypePtrOutput)
}

func (in *parameterTypePtr) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ParameterTypePtrOutput)
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

type PublicNetworkAccess string

const (
	// Public network access is enabled.
	PublicNetworkAccessEnabled = PublicNetworkAccess("Enabled")
	// Public network access is disabled.
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return e.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return PublicNetworkAccess(e).ToPublicNetworkAccessOutputWithContext(ctx).ToPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccess) *PublicNetworkAccess {
		return &v
	}).(PublicNetworkAccessPtrOutput)
}

func (o PublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) Elem() PublicNetworkAccessOutput {
	return o.ApplyT(func(v *PublicNetworkAccess) PublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccess
		return ret
	}).(PublicNetworkAccessOutput)
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessInput interface {
	pulumi.Input

	ToPublicNetworkAccessOutput() PublicNetworkAccessOutput
	ToPublicNetworkAccessOutputWithContext(context.Context) PublicNetworkAccessOutput
}

var publicNetworkAccessPtrType = reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()

type PublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput
	ToPublicNetworkAccessPtrOutputWithContext(context.Context) PublicNetworkAccessPtrOutput
}

type publicNetworkAccessPtr string

func PublicNetworkAccessPtr(v string) PublicNetworkAccessPtrInput {
	return (*publicNetworkAccessPtr)(&v)
}

func (*publicNetworkAccessPtr) ElementType() reflect.Type {
	return publicNetworkAccessPtrType
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessPtrOutput)
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessPtrOutput)
}

type RtspTransport string

const (
	// HTTP transport. RTSP messages are exchanged over long running HTTP requests and RTP packets are interleaved within the HTTP channel.
	RtspTransportHttp = RtspTransport("Http")
	// TCP transport. RTSP is used directly over TCP and RTP packets are interleaved within the TCP channel.
	RtspTransportTcp = RtspTransport("Tcp")
)

func (RtspTransport) ElementType() reflect.Type {
	return reflect.TypeOf((*RtspTransport)(nil)).Elem()
}

func (e RtspTransport) ToRtspTransportOutput() RtspTransportOutput {
	return pulumi.ToOutput(e).(RtspTransportOutput)
}

func (e RtspTransport) ToRtspTransportOutputWithContext(ctx context.Context) RtspTransportOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RtspTransportOutput)
}

func (e RtspTransport) ToRtspTransportPtrOutput() RtspTransportPtrOutput {
	return e.ToRtspTransportPtrOutputWithContext(context.Background())
}

func (e RtspTransport) ToRtspTransportPtrOutputWithContext(ctx context.Context) RtspTransportPtrOutput {
	return RtspTransport(e).ToRtspTransportOutputWithContext(ctx).ToRtspTransportPtrOutputWithContext(ctx)
}

func (e RtspTransport) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RtspTransport) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RtspTransport) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RtspTransport) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RtspTransportOutput struct{ *pulumi.OutputState }

func (RtspTransportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RtspTransport)(nil)).Elem()
}

func (o RtspTransportOutput) ToRtspTransportOutput() RtspTransportOutput {
	return o
}

func (o RtspTransportOutput) ToRtspTransportOutputWithContext(ctx context.Context) RtspTransportOutput {
	return o
}

func (o RtspTransportOutput) ToRtspTransportPtrOutput() RtspTransportPtrOutput {
	return o.ToRtspTransportPtrOutputWithContext(context.Background())
}

func (o RtspTransportOutput) ToRtspTransportPtrOutputWithContext(ctx context.Context) RtspTransportPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RtspTransport) *RtspTransport {
		return &v
	}).(RtspTransportPtrOutput)
}

func (o RtspTransportOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RtspTransportOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RtspTransport) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RtspTransportOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RtspTransportOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RtspTransport) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RtspTransportPtrOutput struct{ *pulumi.OutputState }

func (RtspTransportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RtspTransport)(nil)).Elem()
}

func (o RtspTransportPtrOutput) ToRtspTransportPtrOutput() RtspTransportPtrOutput {
	return o
}

func (o RtspTransportPtrOutput) ToRtspTransportPtrOutputWithContext(ctx context.Context) RtspTransportPtrOutput {
	return o
}

func (o RtspTransportPtrOutput) Elem() RtspTransportOutput {
	return o.ApplyT(func(v *RtspTransport) RtspTransport {
		if v != nil {
			return *v
		}
		var ret RtspTransport
		return ret
	}).(RtspTransportOutput)
}

func (o RtspTransportPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RtspTransportPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RtspTransport) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RtspTransportInput interface {
	pulumi.Input

	ToRtspTransportOutput() RtspTransportOutput
	ToRtspTransportOutputWithContext(context.Context) RtspTransportOutput
}

var rtspTransportPtrType = reflect.TypeOf((**RtspTransport)(nil)).Elem()

type RtspTransportPtrInput interface {
	pulumi.Input

	ToRtspTransportPtrOutput() RtspTransportPtrOutput
	ToRtspTransportPtrOutputWithContext(context.Context) RtspTransportPtrOutput
}

type rtspTransportPtr string

func RtspTransportPtr(v string) RtspTransportPtrInput {
	return (*rtspTransportPtr)(&v)
}

func (*rtspTransportPtr) ElementType() reflect.Type {
	return rtspTransportPtrType
}

func (in *rtspTransportPtr) ToRtspTransportPtrOutput() RtspTransportPtrOutput {
	return pulumi.ToOutput(in).(RtspTransportPtrOutput)
}

func (in *rtspTransportPtr) ToRtspTransportPtrOutputWithContext(ctx context.Context) RtspTransportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RtspTransportPtrOutput)
}

type SkuName string

const (
	// Represents the Live S1 SKU name. Using this SKU you can create live pipelines to capture, record, and stream live video from RTSP-capable cameras at bitrate settings from 0.5 Kbps to 3000 Kbps.
	SkuName_Live_S1 = SkuName("Live_S1")
	// Represents the Batch S1 SKU name. Using this SKU you can create pipeline jobs to process recorded content.
	SkuName_Batch_S1 = SkuName("Batch_S1")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

type VideoScaleMode string

const (
	// Pads the video with black horizontal stripes (letterbox) or black vertical stripes (pillar-box) so the video is resized to the specified dimensions while not altering the content aspect ratio.
	VideoScaleModePad = VideoScaleMode("Pad")
	// Preserves the same aspect ratio as the input video. If only one video dimension is provided, the second dimension is calculated based on the input video aspect ratio. When 2 dimensions are provided, the video is resized to fit the most constraining dimension, considering the input video size and aspect ratio.
	VideoScaleModePreserveAspectRatio = VideoScaleMode("PreserveAspectRatio")
	// Stretches the original video so it resized to the specified dimensions.
	VideoScaleModeStretch = VideoScaleMode("Stretch")
)

func (VideoScaleMode) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoScaleMode)(nil)).Elem()
}

func (e VideoScaleMode) ToVideoScaleModeOutput() VideoScaleModeOutput {
	return pulumi.ToOutput(e).(VideoScaleModeOutput)
}

func (e VideoScaleMode) ToVideoScaleModeOutputWithContext(ctx context.Context) VideoScaleModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VideoScaleModeOutput)
}

func (e VideoScaleMode) ToVideoScaleModePtrOutput() VideoScaleModePtrOutput {
	return e.ToVideoScaleModePtrOutputWithContext(context.Background())
}

func (e VideoScaleMode) ToVideoScaleModePtrOutputWithContext(ctx context.Context) VideoScaleModePtrOutput {
	return VideoScaleMode(e).ToVideoScaleModeOutputWithContext(ctx).ToVideoScaleModePtrOutputWithContext(ctx)
}

func (e VideoScaleMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VideoScaleMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VideoScaleMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VideoScaleMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VideoScaleModeOutput struct{ *pulumi.OutputState }

func (VideoScaleModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VideoScaleMode)(nil)).Elem()
}

func (o VideoScaleModeOutput) ToVideoScaleModeOutput() VideoScaleModeOutput {
	return o
}

func (o VideoScaleModeOutput) ToVideoScaleModeOutputWithContext(ctx context.Context) VideoScaleModeOutput {
	return o
}

func (o VideoScaleModeOutput) ToVideoScaleModePtrOutput() VideoScaleModePtrOutput {
	return o.ToVideoScaleModePtrOutputWithContext(context.Background())
}

func (o VideoScaleModeOutput) ToVideoScaleModePtrOutputWithContext(ctx context.Context) VideoScaleModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VideoScaleMode) *VideoScaleMode {
		return &v
	}).(VideoScaleModePtrOutput)
}

func (o VideoScaleModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VideoScaleModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VideoScaleMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VideoScaleModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VideoScaleModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VideoScaleMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VideoScaleModePtrOutput struct{ *pulumi.OutputState }

func (VideoScaleModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoScaleMode)(nil)).Elem()
}

func (o VideoScaleModePtrOutput) ToVideoScaleModePtrOutput() VideoScaleModePtrOutput {
	return o
}

func (o VideoScaleModePtrOutput) ToVideoScaleModePtrOutputWithContext(ctx context.Context) VideoScaleModePtrOutput {
	return o
}

func (o VideoScaleModePtrOutput) Elem() VideoScaleModeOutput {
	return o.ApplyT(func(v *VideoScaleMode) VideoScaleMode {
		if v != nil {
			return *v
		}
		var ret VideoScaleMode
		return ret
	}).(VideoScaleModeOutput)
}

func (o VideoScaleModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VideoScaleModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VideoScaleMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VideoScaleModeInput interface {
	pulumi.Input

	ToVideoScaleModeOutput() VideoScaleModeOutput
	ToVideoScaleModeOutputWithContext(context.Context) VideoScaleModeOutput
}

var videoScaleModePtrType = reflect.TypeOf((**VideoScaleMode)(nil)).Elem()

type VideoScaleModePtrInput interface {
	pulumi.Input

	ToVideoScaleModePtrOutput() VideoScaleModePtrOutput
	ToVideoScaleModePtrOutputWithContext(context.Context) VideoScaleModePtrOutput
}

type videoScaleModePtr string

func VideoScaleModePtr(v string) VideoScaleModePtrInput {
	return (*videoScaleModePtr)(&v)
}

func (*videoScaleModePtr) ElementType() reflect.Type {
	return videoScaleModePtrType
}

func (in *videoScaleModePtr) ToVideoScaleModePtrOutput() VideoScaleModePtrOutput {
	return pulumi.ToOutput(in).(VideoScaleModePtrOutput)
}

func (in *videoScaleModePtr) ToVideoScaleModePtrOutputWithContext(ctx context.Context) VideoScaleModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VideoScaleModePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPolicyEccAlgoOutput{})
	pulumi.RegisterOutputType(AccessPolicyEccAlgoPtrOutput{})
	pulumi.RegisterOutputType(AccessPolicyRoleOutput{})
	pulumi.RegisterOutputType(AccessPolicyRolePtrOutput{})
	pulumi.RegisterOutputType(AccessPolicyRsaAlgoOutput{})
	pulumi.RegisterOutputType(AccessPolicyRsaAlgoPtrOutput{})
	pulumi.RegisterOutputType(AccountEncryptionKeyTypeOutput{})
	pulumi.RegisterOutputType(AccountEncryptionKeyTypePtrOutput{})
	pulumi.RegisterOutputType(EncoderSystemPresetTypeOutput{})
	pulumi.RegisterOutputType(EncoderSystemPresetTypePtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(ParameterTypeOutput{})
	pulumi.RegisterOutputType(ParameterTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(RtspTransportOutput{})
	pulumi.RegisterOutputType(RtspTransportPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(VideoScaleModeOutput{})
	pulumi.RegisterOutputType(VideoScaleModePtrOutput{})
}
