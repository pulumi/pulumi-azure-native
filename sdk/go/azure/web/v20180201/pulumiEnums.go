


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessControlEntryAction string

const (
	AccessControlEntryActionPermit = AccessControlEntryAction("Permit")
	AccessControlEntryActionDeny   = AccessControlEntryAction("Deny")
)

func (AccessControlEntryAction) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControlEntryAction)(nil)).Elem()
}

func (e AccessControlEntryAction) ToAccessControlEntryActionOutput() AccessControlEntryActionOutput {
	return pulumi.ToOutput(e).(AccessControlEntryActionOutput)
}

func (e AccessControlEntryAction) ToAccessControlEntryActionOutputWithContext(ctx context.Context) AccessControlEntryActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessControlEntryActionOutput)
}

func (e AccessControlEntryAction) ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput {
	return e.ToAccessControlEntryActionPtrOutputWithContext(context.Background())
}

func (e AccessControlEntryAction) ToAccessControlEntryActionPtrOutputWithContext(ctx context.Context) AccessControlEntryActionPtrOutput {
	return AccessControlEntryAction(e).ToAccessControlEntryActionOutputWithContext(ctx).ToAccessControlEntryActionPtrOutputWithContext(ctx)
}

func (e AccessControlEntryAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessControlEntryAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessControlEntryAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessControlEntryAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessControlEntryActionOutput struct{ *pulumi.OutputState }

func (AccessControlEntryActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControlEntryAction)(nil)).Elem()
}

func (o AccessControlEntryActionOutput) ToAccessControlEntryActionOutput() AccessControlEntryActionOutput {
	return o
}

func (o AccessControlEntryActionOutput) ToAccessControlEntryActionOutputWithContext(ctx context.Context) AccessControlEntryActionOutput {
	return o
}

func (o AccessControlEntryActionOutput) ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput {
	return o.ToAccessControlEntryActionPtrOutputWithContext(context.Background())
}

func (o AccessControlEntryActionOutput) ToAccessControlEntryActionPtrOutputWithContext(ctx context.Context) AccessControlEntryActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessControlEntryAction) *AccessControlEntryAction {
		return &v
	}).(AccessControlEntryActionPtrOutput)
}

func (o AccessControlEntryActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessControlEntryActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessControlEntryAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessControlEntryActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessControlEntryActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessControlEntryAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessControlEntryActionPtrOutput struct{ *pulumi.OutputState }

func (AccessControlEntryActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessControlEntryAction)(nil)).Elem()
}

func (o AccessControlEntryActionPtrOutput) ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput {
	return o
}

func (o AccessControlEntryActionPtrOutput) ToAccessControlEntryActionPtrOutputWithContext(ctx context.Context) AccessControlEntryActionPtrOutput {
	return o
}

func (o AccessControlEntryActionPtrOutput) Elem() AccessControlEntryActionOutput {
	return o.ApplyT(func(v *AccessControlEntryAction) AccessControlEntryAction {
		if v != nil {
			return *v
		}
		var ret AccessControlEntryAction
		return ret
	}).(AccessControlEntryActionOutput)
}

func (o AccessControlEntryActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessControlEntryActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessControlEntryAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessControlEntryActionInput interface {
	pulumi.Input

	ToAccessControlEntryActionOutput() AccessControlEntryActionOutput
	ToAccessControlEntryActionOutputWithContext(context.Context) AccessControlEntryActionOutput
}

var accessControlEntryActionPtrType = reflect.TypeOf((**AccessControlEntryAction)(nil)).Elem()

type AccessControlEntryActionPtrInput interface {
	pulumi.Input

	ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput
	ToAccessControlEntryActionPtrOutputWithContext(context.Context) AccessControlEntryActionPtrOutput
}

type accessControlEntryActionPtr string

func AccessControlEntryActionPtr(v string) AccessControlEntryActionPtrInput {
	return (*accessControlEntryActionPtr)(&v)
}

func (*accessControlEntryActionPtr) ElementType() reflect.Type {
	return accessControlEntryActionPtrType
}

func (in *accessControlEntryActionPtr) ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput {
	return pulumi.ToOutput(in).(AccessControlEntryActionPtrOutput)
}

func (in *accessControlEntryActionPtr) ToAccessControlEntryActionPtrOutputWithContext(ctx context.Context) AccessControlEntryActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessControlEntryActionPtrOutput)
}

type AutoHealActionType string

const (
	AutoHealActionTypeRecycle      = AutoHealActionType("Recycle")
	AutoHealActionTypeLogEvent     = AutoHealActionType("LogEvent")
	AutoHealActionTypeCustomAction = AutoHealActionType("CustomAction")
)

func (AutoHealActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionType)(nil)).Elem()
}

func (e AutoHealActionType) ToAutoHealActionTypeOutput() AutoHealActionTypeOutput {
	return pulumi.ToOutput(e).(AutoHealActionTypeOutput)
}

func (e AutoHealActionType) ToAutoHealActionTypeOutputWithContext(ctx context.Context) AutoHealActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoHealActionTypeOutput)
}

func (e AutoHealActionType) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return e.ToAutoHealActionTypePtrOutputWithContext(context.Background())
}

func (e AutoHealActionType) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return AutoHealActionType(e).ToAutoHealActionTypeOutputWithContext(ctx).ToAutoHealActionTypePtrOutputWithContext(ctx)
}

func (e AutoHealActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoHealActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoHealActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoHealActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoHealActionTypeOutput struct{ *pulumi.OutputState }

func (AutoHealActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionType)(nil)).Elem()
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypeOutput() AutoHealActionTypeOutput {
	return o
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypeOutputWithContext(ctx context.Context) AutoHealActionTypeOutput {
	return o
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return o.ToAutoHealActionTypePtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealActionType) *AutoHealActionType {
		return &v
	}).(AutoHealActionTypePtrOutput)
}

func (o AutoHealActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoHealActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoHealActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoHealActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoHealActionTypePtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActionType)(nil)).Elem()
}

func (o AutoHealActionTypePtrOutput) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return o
}

func (o AutoHealActionTypePtrOutput) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return o
}

func (o AutoHealActionTypePtrOutput) Elem() AutoHealActionTypeOutput {
	return o.ApplyT(func(v *AutoHealActionType) AutoHealActionType {
		if v != nil {
			return *v
		}
		var ret AutoHealActionType
		return ret
	}).(AutoHealActionTypeOutput)
}

func (o AutoHealActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoHealActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoHealActionTypeInput interface {
	pulumi.Input

	ToAutoHealActionTypeOutput() AutoHealActionTypeOutput
	ToAutoHealActionTypeOutputWithContext(context.Context) AutoHealActionTypeOutput
}

var autoHealActionTypePtrType = reflect.TypeOf((**AutoHealActionType)(nil)).Elem()

type AutoHealActionTypePtrInput interface {
	pulumi.Input

	ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput
	ToAutoHealActionTypePtrOutputWithContext(context.Context) AutoHealActionTypePtrOutput
}

type autoHealActionTypePtr string

func AutoHealActionTypePtr(v string) AutoHealActionTypePtrInput {
	return (*autoHealActionTypePtr)(&v)
}

func (*autoHealActionTypePtr) ElementType() reflect.Type {
	return autoHealActionTypePtrType
}

func (in *autoHealActionTypePtr) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return pulumi.ToOutput(in).(AutoHealActionTypePtrOutput)
}

func (in *autoHealActionTypePtr) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoHealActionTypePtrOutput)
}

type AzureResourceType string

const (
	AzureResourceTypeWebsite        = AzureResourceType("Website")
	AzureResourceTypeTrafficManager = AzureResourceType("TrafficManager")
)

func (AzureResourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceType)(nil)).Elem()
}

func (e AzureResourceType) ToAzureResourceTypeOutput() AzureResourceTypeOutput {
	return pulumi.ToOutput(e).(AzureResourceTypeOutput)
}

func (e AzureResourceType) ToAzureResourceTypeOutputWithContext(ctx context.Context) AzureResourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureResourceTypeOutput)
}

func (e AzureResourceType) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return e.ToAzureResourceTypePtrOutputWithContext(context.Background())
}

func (e AzureResourceType) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return AzureResourceType(e).ToAzureResourceTypeOutputWithContext(ctx).ToAzureResourceTypePtrOutputWithContext(ctx)
}

func (e AzureResourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureResourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureResourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureResourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureResourceTypeOutput struct{ *pulumi.OutputState }

func (AzureResourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceType)(nil)).Elem()
}

func (o AzureResourceTypeOutput) ToAzureResourceTypeOutput() AzureResourceTypeOutput {
	return o
}

func (o AzureResourceTypeOutput) ToAzureResourceTypeOutputWithContext(ctx context.Context) AzureResourceTypeOutput {
	return o
}

func (o AzureResourceTypeOutput) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return o.ToAzureResourceTypePtrOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureResourceType) *AzureResourceType {
		return &v
	}).(AzureResourceTypePtrOutput)
}

func (o AzureResourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureResourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureResourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureResourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureResourceTypePtrOutput struct{ *pulumi.OutputState }

func (AzureResourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureResourceType)(nil)).Elem()
}

func (o AzureResourceTypePtrOutput) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return o
}

func (o AzureResourceTypePtrOutput) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return o
}

func (o AzureResourceTypePtrOutput) Elem() AzureResourceTypeOutput {
	return o.ApplyT(func(v *AzureResourceType) AzureResourceType {
		if v != nil {
			return *v
		}
		var ret AzureResourceType
		return ret
	}).(AzureResourceTypeOutput)
}

func (o AzureResourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureResourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureResourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureResourceTypeInput interface {
	pulumi.Input

	ToAzureResourceTypeOutput() AzureResourceTypeOutput
	ToAzureResourceTypeOutputWithContext(context.Context) AzureResourceTypeOutput
}

var azureResourceTypePtrType = reflect.TypeOf((**AzureResourceType)(nil)).Elem()

type AzureResourceTypePtrInput interface {
	pulumi.Input

	ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput
	ToAzureResourceTypePtrOutputWithContext(context.Context) AzureResourceTypePtrOutput
}

type azureResourceTypePtr string

func AzureResourceTypePtr(v string) AzureResourceTypePtrInput {
	return (*azureResourceTypePtr)(&v)
}

func (*azureResourceTypePtr) ElementType() reflect.Type {
	return azureResourceTypePtrType
}

func (in *azureResourceTypePtr) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return pulumi.ToOutput(in).(AzureResourceTypePtrOutput)
}

func (in *azureResourceTypePtr) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureResourceTypePtrOutput)
}

type AzureStorageType string

const (
	AzureStorageTypeAzureFiles = AzureStorageType("AzureFiles")
	AzureStorageTypeAzureBlob  = AzureStorageType("AzureBlob")
)

func (AzureStorageType) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageType)(nil)).Elem()
}

func (e AzureStorageType) ToAzureStorageTypeOutput() AzureStorageTypeOutput {
	return pulumi.ToOutput(e).(AzureStorageTypeOutput)
}

func (e AzureStorageType) ToAzureStorageTypeOutputWithContext(ctx context.Context) AzureStorageTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureStorageTypeOutput)
}

func (e AzureStorageType) ToAzureStorageTypePtrOutput() AzureStorageTypePtrOutput {
	return e.ToAzureStorageTypePtrOutputWithContext(context.Background())
}

func (e AzureStorageType) ToAzureStorageTypePtrOutputWithContext(ctx context.Context) AzureStorageTypePtrOutput {
	return AzureStorageType(e).ToAzureStorageTypeOutputWithContext(ctx).ToAzureStorageTypePtrOutputWithContext(ctx)
}

func (e AzureStorageType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureStorageType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureStorageType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureStorageType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureStorageTypeOutput struct{ *pulumi.OutputState }

func (AzureStorageTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageType)(nil)).Elem()
}

func (o AzureStorageTypeOutput) ToAzureStorageTypeOutput() AzureStorageTypeOutput {
	return o
}

func (o AzureStorageTypeOutput) ToAzureStorageTypeOutputWithContext(ctx context.Context) AzureStorageTypeOutput {
	return o
}

func (o AzureStorageTypeOutput) ToAzureStorageTypePtrOutput() AzureStorageTypePtrOutput {
	return o.ToAzureStorageTypePtrOutputWithContext(context.Background())
}

func (o AzureStorageTypeOutput) ToAzureStorageTypePtrOutputWithContext(ctx context.Context) AzureStorageTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureStorageType) *AzureStorageType {
		return &v
	}).(AzureStorageTypePtrOutput)
}

func (o AzureStorageTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureStorageTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureStorageType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureStorageTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureStorageTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureStorageType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureStorageTypePtrOutput struct{ *pulumi.OutputState }

func (AzureStorageTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStorageType)(nil)).Elem()
}

func (o AzureStorageTypePtrOutput) ToAzureStorageTypePtrOutput() AzureStorageTypePtrOutput {
	return o
}

func (o AzureStorageTypePtrOutput) ToAzureStorageTypePtrOutputWithContext(ctx context.Context) AzureStorageTypePtrOutput {
	return o
}

func (o AzureStorageTypePtrOutput) Elem() AzureStorageTypeOutput {
	return o.ApplyT(func(v *AzureStorageType) AzureStorageType {
		if v != nil {
			return *v
		}
		var ret AzureStorageType
		return ret
	}).(AzureStorageTypeOutput)
}

func (o AzureStorageTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureStorageTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureStorageType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureStorageTypeInput interface {
	pulumi.Input

	ToAzureStorageTypeOutput() AzureStorageTypeOutput
	ToAzureStorageTypeOutputWithContext(context.Context) AzureStorageTypeOutput
}

var azureStorageTypePtrType = reflect.TypeOf((**AzureStorageType)(nil)).Elem()

type AzureStorageTypePtrInput interface {
	pulumi.Input

	ToAzureStorageTypePtrOutput() AzureStorageTypePtrOutput
	ToAzureStorageTypePtrOutputWithContext(context.Context) AzureStorageTypePtrOutput
}

type azureStorageTypePtr string

func AzureStorageTypePtr(v string) AzureStorageTypePtrInput {
	return (*azureStorageTypePtr)(&v)
}

func (*azureStorageTypePtr) ElementType() reflect.Type {
	return azureStorageTypePtrType
}

func (in *azureStorageTypePtr) ToAzureStorageTypePtrOutput() AzureStorageTypePtrOutput {
	return pulumi.ToOutput(in).(AzureStorageTypePtrOutput)
}

func (in *azureStorageTypePtr) ToAzureStorageTypePtrOutputWithContext(ctx context.Context) AzureStorageTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureStorageTypePtrOutput)
}

type BuiltInAuthenticationProvider string

const (
	BuiltInAuthenticationProviderAzureActiveDirectory = BuiltInAuthenticationProvider("AzureActiveDirectory")
	BuiltInAuthenticationProviderFacebook             = BuiltInAuthenticationProvider("Facebook")
	BuiltInAuthenticationProviderGoogle               = BuiltInAuthenticationProvider("Google")
	BuiltInAuthenticationProviderMicrosoftAccount     = BuiltInAuthenticationProvider("MicrosoftAccount")
	BuiltInAuthenticationProviderTwitter              = BuiltInAuthenticationProvider("Twitter")
)

func (BuiltInAuthenticationProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInAuthenticationProvider)(nil)).Elem()
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput {
	return pulumi.ToOutput(e).(BuiltInAuthenticationProviderOutput)
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BuiltInAuthenticationProviderOutput)
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return e.ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Background())
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return BuiltInAuthenticationProvider(e).ToBuiltInAuthenticationProviderOutputWithContext(ctx).ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx)
}

func (e BuiltInAuthenticationProvider) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BuiltInAuthenticationProvider) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BuiltInAuthenticationProvider) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BuiltInAuthenticationProvider) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BuiltInAuthenticationProviderOutput struct{ *pulumi.OutputState }

func (BuiltInAuthenticationProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInAuthenticationProvider)(nil)).Elem()
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput {
	return o
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderOutput {
	return o
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return o.ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuiltInAuthenticationProvider) *BuiltInAuthenticationProvider {
		return &v
	}).(BuiltInAuthenticationProviderPtrOutput)
}

func (o BuiltInAuthenticationProviderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BuiltInAuthenticationProvider) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BuiltInAuthenticationProviderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BuiltInAuthenticationProvider) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BuiltInAuthenticationProviderPtrOutput struct{ *pulumi.OutputState }

func (BuiltInAuthenticationProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuiltInAuthenticationProvider)(nil)).Elem()
}

func (o BuiltInAuthenticationProviderPtrOutput) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return o
}

func (o BuiltInAuthenticationProviderPtrOutput) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return o
}

func (o BuiltInAuthenticationProviderPtrOutput) Elem() BuiltInAuthenticationProviderOutput {
	return o.ApplyT(func(v *BuiltInAuthenticationProvider) BuiltInAuthenticationProvider {
		if v != nil {
			return *v
		}
		var ret BuiltInAuthenticationProvider
		return ret
	}).(BuiltInAuthenticationProviderOutput)
}

func (o BuiltInAuthenticationProviderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BuiltInAuthenticationProvider) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BuiltInAuthenticationProviderInput interface {
	pulumi.Input

	ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput
	ToBuiltInAuthenticationProviderOutputWithContext(context.Context) BuiltInAuthenticationProviderOutput
}

var builtInAuthenticationProviderPtrType = reflect.TypeOf((**BuiltInAuthenticationProvider)(nil)).Elem()

type BuiltInAuthenticationProviderPtrInput interface {
	pulumi.Input

	ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput
	ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Context) BuiltInAuthenticationProviderPtrOutput
}

type builtInAuthenticationProviderPtr string

func BuiltInAuthenticationProviderPtr(v string) BuiltInAuthenticationProviderPtrInput {
	return (*builtInAuthenticationProviderPtr)(&v)
}

func (*builtInAuthenticationProviderPtr) ElementType() reflect.Type {
	return builtInAuthenticationProviderPtrType
}

func (in *builtInAuthenticationProviderPtr) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return pulumi.ToOutput(in).(BuiltInAuthenticationProviderPtrOutput)
}

func (in *builtInAuthenticationProviderPtr) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BuiltInAuthenticationProviderPtrOutput)
}

type ComputeModeOptions string

const (
	ComputeModeOptionsShared    = ComputeModeOptions("Shared")
	ComputeModeOptionsDedicated = ComputeModeOptions("Dedicated")
	ComputeModeOptionsDynamic   = ComputeModeOptions("Dynamic")
)

func (ComputeModeOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeModeOptions)(nil)).Elem()
}

func (e ComputeModeOptions) ToComputeModeOptionsOutput() ComputeModeOptionsOutput {
	return pulumi.ToOutput(e).(ComputeModeOptionsOutput)
}

func (e ComputeModeOptions) ToComputeModeOptionsOutputWithContext(ctx context.Context) ComputeModeOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeModeOptionsOutput)
}

func (e ComputeModeOptions) ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput {
	return e.ToComputeModeOptionsPtrOutputWithContext(context.Background())
}

func (e ComputeModeOptions) ToComputeModeOptionsPtrOutputWithContext(ctx context.Context) ComputeModeOptionsPtrOutput {
	return ComputeModeOptions(e).ToComputeModeOptionsOutputWithContext(ctx).ToComputeModeOptionsPtrOutputWithContext(ctx)
}

func (e ComputeModeOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeModeOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeModeOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeModeOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeModeOptionsOutput struct{ *pulumi.OutputState }

func (ComputeModeOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeModeOptions)(nil)).Elem()
}

func (o ComputeModeOptionsOutput) ToComputeModeOptionsOutput() ComputeModeOptionsOutput {
	return o
}

func (o ComputeModeOptionsOutput) ToComputeModeOptionsOutputWithContext(ctx context.Context) ComputeModeOptionsOutput {
	return o
}

func (o ComputeModeOptionsOutput) ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput {
	return o.ToComputeModeOptionsPtrOutputWithContext(context.Background())
}

func (o ComputeModeOptionsOutput) ToComputeModeOptionsPtrOutputWithContext(ctx context.Context) ComputeModeOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeModeOptions) *ComputeModeOptions {
		return &v
	}).(ComputeModeOptionsPtrOutput)
}

func (o ComputeModeOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeModeOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeModeOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeModeOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeModeOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeModeOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeModeOptionsPtrOutput struct{ *pulumi.OutputState }

func (ComputeModeOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeModeOptions)(nil)).Elem()
}

func (o ComputeModeOptionsPtrOutput) ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput {
	return o
}

func (o ComputeModeOptionsPtrOutput) ToComputeModeOptionsPtrOutputWithContext(ctx context.Context) ComputeModeOptionsPtrOutput {
	return o
}

func (o ComputeModeOptionsPtrOutput) Elem() ComputeModeOptionsOutput {
	return o.ApplyT(func(v *ComputeModeOptions) ComputeModeOptions {
		if v != nil {
			return *v
		}
		var ret ComputeModeOptions
		return ret
	}).(ComputeModeOptionsOutput)
}

func (o ComputeModeOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeModeOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeModeOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeModeOptionsInput interface {
	pulumi.Input

	ToComputeModeOptionsOutput() ComputeModeOptionsOutput
	ToComputeModeOptionsOutputWithContext(context.Context) ComputeModeOptionsOutput
}

var computeModeOptionsPtrType = reflect.TypeOf((**ComputeModeOptions)(nil)).Elem()

type ComputeModeOptionsPtrInput interface {
	pulumi.Input

	ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput
	ToComputeModeOptionsPtrOutputWithContext(context.Context) ComputeModeOptionsPtrOutput
}

type computeModeOptionsPtr string

func ComputeModeOptionsPtr(v string) ComputeModeOptionsPtrInput {
	return (*computeModeOptionsPtr)(&v)
}

func (*computeModeOptionsPtr) ElementType() reflect.Type {
	return computeModeOptionsPtrType
}

func (in *computeModeOptionsPtr) ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput {
	return pulumi.ToOutput(in).(ComputeModeOptionsPtrOutput)
}

func (in *computeModeOptionsPtr) ToComputeModeOptionsPtrOutputWithContext(ctx context.Context) ComputeModeOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeModeOptionsPtrOutput)
}

type ConnectionStringType string

const (
	ConnectionStringTypeMySql           = ConnectionStringType("MySql")
	ConnectionStringTypeSQLServer       = ConnectionStringType("SQLServer")
	ConnectionStringTypeSQLAzure        = ConnectionStringType("SQLAzure")
	ConnectionStringTypeCustom          = ConnectionStringType("Custom")
	ConnectionStringTypeNotificationHub = ConnectionStringType("NotificationHub")
	ConnectionStringTypeServiceBus      = ConnectionStringType("ServiceBus")
	ConnectionStringTypeEventHub        = ConnectionStringType("EventHub")
	ConnectionStringTypeApiHub          = ConnectionStringType("ApiHub")
	ConnectionStringTypeDocDb           = ConnectionStringType("DocDb")
	ConnectionStringTypeRedisCache      = ConnectionStringType("RedisCache")
	ConnectionStringTypePostgreSQL      = ConnectionStringType("PostgreSQL")
)

func (ConnectionStringType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStringType)(nil)).Elem()
}

func (e ConnectionStringType) ToConnectionStringTypeOutput() ConnectionStringTypeOutput {
	return pulumi.ToOutput(e).(ConnectionStringTypeOutput)
}

func (e ConnectionStringType) ToConnectionStringTypeOutputWithContext(ctx context.Context) ConnectionStringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionStringTypeOutput)
}

func (e ConnectionStringType) ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput {
	return e.ToConnectionStringTypePtrOutputWithContext(context.Background())
}

func (e ConnectionStringType) ToConnectionStringTypePtrOutputWithContext(ctx context.Context) ConnectionStringTypePtrOutput {
	return ConnectionStringType(e).ToConnectionStringTypeOutputWithContext(ctx).ToConnectionStringTypePtrOutputWithContext(ctx)
}

func (e ConnectionStringType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStringType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStringType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionStringType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionStringTypeOutput struct{ *pulumi.OutputState }

func (ConnectionStringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStringType)(nil)).Elem()
}

func (o ConnectionStringTypeOutput) ToConnectionStringTypeOutput() ConnectionStringTypeOutput {
	return o
}

func (o ConnectionStringTypeOutput) ToConnectionStringTypeOutputWithContext(ctx context.Context) ConnectionStringTypeOutput {
	return o
}

func (o ConnectionStringTypeOutput) ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput {
	return o.ToConnectionStringTypePtrOutputWithContext(context.Background())
}

func (o ConnectionStringTypeOutput) ToConnectionStringTypePtrOutputWithContext(ctx context.Context) ConnectionStringTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionStringType) *ConnectionStringType {
		return &v
	}).(ConnectionStringTypePtrOutput)
}

func (o ConnectionStringTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionStringTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStringType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionStringTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStringTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStringType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionStringTypePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStringTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStringType)(nil)).Elem()
}

func (o ConnectionStringTypePtrOutput) ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput {
	return o
}

func (o ConnectionStringTypePtrOutput) ToConnectionStringTypePtrOutputWithContext(ctx context.Context) ConnectionStringTypePtrOutput {
	return o
}

func (o ConnectionStringTypePtrOutput) Elem() ConnectionStringTypeOutput {
	return o.ApplyT(func(v *ConnectionStringType) ConnectionStringType {
		if v != nil {
			return *v
		}
		var ret ConnectionStringType
		return ret
	}).(ConnectionStringTypeOutput)
}

func (o ConnectionStringTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStringTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionStringType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionStringTypeInput interface {
	pulumi.Input

	ToConnectionStringTypeOutput() ConnectionStringTypeOutput
	ToConnectionStringTypeOutputWithContext(context.Context) ConnectionStringTypeOutput
}

var connectionStringTypePtrType = reflect.TypeOf((**ConnectionStringType)(nil)).Elem()

type ConnectionStringTypePtrInput interface {
	pulumi.Input

	ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput
	ToConnectionStringTypePtrOutputWithContext(context.Context) ConnectionStringTypePtrOutput
}

type connectionStringTypePtr string

func ConnectionStringTypePtr(v string) ConnectionStringTypePtrInput {
	return (*connectionStringTypePtr)(&v)
}

func (*connectionStringTypePtr) ElementType() reflect.Type {
	return connectionStringTypePtrType
}

func (in *connectionStringTypePtr) ToConnectionStringTypePtrOutput() ConnectionStringTypePtrOutput {
	return pulumi.ToOutput(in).(ConnectionStringTypePtrOutput)
}

func (in *connectionStringTypePtr) ToConnectionStringTypePtrOutputWithContext(ctx context.Context) ConnectionStringTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionStringTypePtrOutput)
}

type CustomHostNameDnsRecordType string

const (
	CustomHostNameDnsRecordTypeCName = CustomHostNameDnsRecordType("CName")
	CustomHostNameDnsRecordTypeA     = CustomHostNameDnsRecordType("A")
)

func (CustomHostNameDnsRecordType) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostNameDnsRecordType)(nil)).Elem()
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput {
	return pulumi.ToOutput(e).(CustomHostNameDnsRecordTypeOutput)
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypeOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CustomHostNameDnsRecordTypeOutput)
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return e.ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Background())
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return CustomHostNameDnsRecordType(e).ToCustomHostNameDnsRecordTypeOutputWithContext(ctx).ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx)
}

func (e CustomHostNameDnsRecordType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomHostNameDnsRecordType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomHostNameDnsRecordType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CustomHostNameDnsRecordType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CustomHostNameDnsRecordTypeOutput struct{ *pulumi.OutputState }

func (CustomHostNameDnsRecordTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostNameDnsRecordType)(nil)).Elem()
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput {
	return o
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypeOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypeOutput {
	return o
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return o.ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomHostNameDnsRecordType) *CustomHostNameDnsRecordType {
		return &v
	}).(CustomHostNameDnsRecordTypePtrOutput)
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomHostNameDnsRecordType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomHostNameDnsRecordType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CustomHostNameDnsRecordTypePtrOutput struct{ *pulumi.OutputState }

func (CustomHostNameDnsRecordTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomHostNameDnsRecordType)(nil)).Elem()
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return o
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return o
}

func (o CustomHostNameDnsRecordTypePtrOutput) Elem() CustomHostNameDnsRecordTypeOutput {
	return o.ApplyT(func(v *CustomHostNameDnsRecordType) CustomHostNameDnsRecordType {
		if v != nil {
			return *v
		}
		var ret CustomHostNameDnsRecordType
		return ret
	}).(CustomHostNameDnsRecordTypeOutput)
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CustomHostNameDnsRecordType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CustomHostNameDnsRecordTypeInput interface {
	pulumi.Input

	ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput
	ToCustomHostNameDnsRecordTypeOutputWithContext(context.Context) CustomHostNameDnsRecordTypeOutput
}

var customHostNameDnsRecordTypePtrType = reflect.TypeOf((**CustomHostNameDnsRecordType)(nil)).Elem()

type CustomHostNameDnsRecordTypePtrInput interface {
	pulumi.Input

	ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput
	ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Context) CustomHostNameDnsRecordTypePtrOutput
}

type customHostNameDnsRecordTypePtr string

func CustomHostNameDnsRecordTypePtr(v string) CustomHostNameDnsRecordTypePtrInput {
	return (*customHostNameDnsRecordTypePtr)(&v)
}

func (*customHostNameDnsRecordTypePtr) ElementType() reflect.Type {
	return customHostNameDnsRecordTypePtrType
}

func (in *customHostNameDnsRecordTypePtr) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return pulumi.ToOutput(in).(CustomHostNameDnsRecordTypePtrOutput)
}

func (in *customHostNameDnsRecordTypePtr) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CustomHostNameDnsRecordTypePtrOutput)
}

type DatabaseType string

const (
	DatabaseTypeSqlAzure   = DatabaseType("SqlAzure")
	DatabaseTypeMySql      = DatabaseType("MySql")
	DatabaseTypeLocalMySql = DatabaseType("LocalMySql")
	DatabaseTypePostgreSql = DatabaseType("PostgreSql")
)

func (DatabaseType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseType)(nil)).Elem()
}

func (e DatabaseType) ToDatabaseTypeOutput() DatabaseTypeOutput {
	return pulumi.ToOutput(e).(DatabaseTypeOutput)
}

func (e DatabaseType) ToDatabaseTypeOutputWithContext(ctx context.Context) DatabaseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseTypeOutput)
}

func (e DatabaseType) ToDatabaseTypePtrOutput() DatabaseTypePtrOutput {
	return e.ToDatabaseTypePtrOutputWithContext(context.Background())
}

func (e DatabaseType) ToDatabaseTypePtrOutputWithContext(ctx context.Context) DatabaseTypePtrOutput {
	return DatabaseType(e).ToDatabaseTypeOutputWithContext(ctx).ToDatabaseTypePtrOutputWithContext(ctx)
}

func (e DatabaseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseTypeOutput struct{ *pulumi.OutputState }

func (DatabaseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseType)(nil)).Elem()
}

func (o DatabaseTypeOutput) ToDatabaseTypeOutput() DatabaseTypeOutput {
	return o
}

func (o DatabaseTypeOutput) ToDatabaseTypeOutputWithContext(ctx context.Context) DatabaseTypeOutput {
	return o
}

func (o DatabaseTypeOutput) ToDatabaseTypePtrOutput() DatabaseTypePtrOutput {
	return o.ToDatabaseTypePtrOutputWithContext(context.Background())
}

func (o DatabaseTypeOutput) ToDatabaseTypePtrOutputWithContext(ctx context.Context) DatabaseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseType) *DatabaseType {
		return &v
	}).(DatabaseTypePtrOutput)
}

func (o DatabaseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseTypePtrOutput struct{ *pulumi.OutputState }

func (DatabaseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseType)(nil)).Elem()
}

func (o DatabaseTypePtrOutput) ToDatabaseTypePtrOutput() DatabaseTypePtrOutput {
	return o
}

func (o DatabaseTypePtrOutput) ToDatabaseTypePtrOutputWithContext(ctx context.Context) DatabaseTypePtrOutput {
	return o
}

func (o DatabaseTypePtrOutput) Elem() DatabaseTypeOutput {
	return o.ApplyT(func(v *DatabaseType) DatabaseType {
		if v != nil {
			return *v
		}
		var ret DatabaseType
		return ret
	}).(DatabaseTypeOutput)
}

func (o DatabaseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabaseTypeInput interface {
	pulumi.Input

	ToDatabaseTypeOutput() DatabaseTypeOutput
	ToDatabaseTypeOutputWithContext(context.Context) DatabaseTypeOutput
}

var databaseTypePtrType = reflect.TypeOf((**DatabaseType)(nil)).Elem()

type DatabaseTypePtrInput interface {
	pulumi.Input

	ToDatabaseTypePtrOutput() DatabaseTypePtrOutput
	ToDatabaseTypePtrOutputWithContext(context.Context) DatabaseTypePtrOutput
}

type databaseTypePtr string

func DatabaseTypePtr(v string) DatabaseTypePtrInput {
	return (*databaseTypePtr)(&v)
}

func (*databaseTypePtr) ElementType() reflect.Type {
	return databaseTypePtrType
}

func (in *databaseTypePtr) ToDatabaseTypePtrOutput() DatabaseTypePtrOutput {
	return pulumi.ToOutput(in).(DatabaseTypePtrOutput)
}

func (in *databaseTypePtr) ToDatabaseTypePtrOutputWithContext(ctx context.Context) DatabaseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseTypePtrOutput)
}

type FrequencyUnit string

const (
	FrequencyUnitDay  = FrequencyUnit("Day")
	FrequencyUnitHour = FrequencyUnit("Hour")
)

func (FrequencyUnit) ElementType() reflect.Type {
	return reflect.TypeOf((*FrequencyUnit)(nil)).Elem()
}

func (e FrequencyUnit) ToFrequencyUnitOutput() FrequencyUnitOutput {
	return pulumi.ToOutput(e).(FrequencyUnitOutput)
}

func (e FrequencyUnit) ToFrequencyUnitOutputWithContext(ctx context.Context) FrequencyUnitOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FrequencyUnitOutput)
}

func (e FrequencyUnit) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return e.ToFrequencyUnitPtrOutputWithContext(context.Background())
}

func (e FrequencyUnit) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return FrequencyUnit(e).ToFrequencyUnitOutputWithContext(ctx).ToFrequencyUnitPtrOutputWithContext(ctx)
}

func (e FrequencyUnit) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrequencyUnit) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrequencyUnit) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FrequencyUnit) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FrequencyUnitOutput struct{ *pulumi.OutputState }

func (FrequencyUnitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrequencyUnit)(nil)).Elem()
}

func (o FrequencyUnitOutput) ToFrequencyUnitOutput() FrequencyUnitOutput {
	return o
}

func (o FrequencyUnitOutput) ToFrequencyUnitOutputWithContext(ctx context.Context) FrequencyUnitOutput {
	return o
}

func (o FrequencyUnitOutput) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return o.ToFrequencyUnitPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrequencyUnit) *FrequencyUnit {
		return &v
	}).(FrequencyUnitPtrOutput)
}

func (o FrequencyUnitOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrequencyUnit) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FrequencyUnitOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrequencyUnit) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FrequencyUnitPtrOutput struct{ *pulumi.OutputState }

func (FrequencyUnitPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrequencyUnit)(nil)).Elem()
}

func (o FrequencyUnitPtrOutput) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return o
}

func (o FrequencyUnitPtrOutput) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return o
}

func (o FrequencyUnitPtrOutput) Elem() FrequencyUnitOutput {
	return o.ApplyT(func(v *FrequencyUnit) FrequencyUnit {
		if v != nil {
			return *v
		}
		var ret FrequencyUnit
		return ret
	}).(FrequencyUnitOutput)
}

func (o FrequencyUnitPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FrequencyUnit) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FrequencyUnitInput interface {
	pulumi.Input

	ToFrequencyUnitOutput() FrequencyUnitOutput
	ToFrequencyUnitOutputWithContext(context.Context) FrequencyUnitOutput
}

var frequencyUnitPtrType = reflect.TypeOf((**FrequencyUnit)(nil)).Elem()

type FrequencyUnitPtrInput interface {
	pulumi.Input

	ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput
	ToFrequencyUnitPtrOutputWithContext(context.Context) FrequencyUnitPtrOutput
}

type frequencyUnitPtr string

func FrequencyUnitPtr(v string) FrequencyUnitPtrInput {
	return (*frequencyUnitPtr)(&v)
}

func (*frequencyUnitPtr) ElementType() reflect.Type {
	return frequencyUnitPtrType
}

func (in *frequencyUnitPtr) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return pulumi.ToOutput(in).(FrequencyUnitPtrOutput)
}

func (in *frequencyUnitPtr) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FrequencyUnitPtrOutput)
}

type FtpsState string

const (
	FtpsStateAllAllowed = FtpsState("AllAllowed")
	FtpsStateFtpsOnly   = FtpsState("FtpsOnly")
	FtpsStateDisabled   = FtpsState("Disabled")
)

func (FtpsState) ElementType() reflect.Type {
	return reflect.TypeOf((*FtpsState)(nil)).Elem()
}

func (e FtpsState) ToFtpsStateOutput() FtpsStateOutput {
	return pulumi.ToOutput(e).(FtpsStateOutput)
}

func (e FtpsState) ToFtpsStateOutputWithContext(ctx context.Context) FtpsStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FtpsStateOutput)
}

func (e FtpsState) ToFtpsStatePtrOutput() FtpsStatePtrOutput {
	return e.ToFtpsStatePtrOutputWithContext(context.Background())
}

func (e FtpsState) ToFtpsStatePtrOutputWithContext(ctx context.Context) FtpsStatePtrOutput {
	return FtpsState(e).ToFtpsStateOutputWithContext(ctx).ToFtpsStatePtrOutputWithContext(ctx)
}

func (e FtpsState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FtpsState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FtpsState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FtpsState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FtpsStateOutput struct{ *pulumi.OutputState }

func (FtpsStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FtpsState)(nil)).Elem()
}

func (o FtpsStateOutput) ToFtpsStateOutput() FtpsStateOutput {
	return o
}

func (o FtpsStateOutput) ToFtpsStateOutputWithContext(ctx context.Context) FtpsStateOutput {
	return o
}

func (o FtpsStateOutput) ToFtpsStatePtrOutput() FtpsStatePtrOutput {
	return o.ToFtpsStatePtrOutputWithContext(context.Background())
}

func (o FtpsStateOutput) ToFtpsStatePtrOutputWithContext(ctx context.Context) FtpsStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FtpsState) *FtpsState {
		return &v
	}).(FtpsStatePtrOutput)
}

func (o FtpsStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FtpsStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FtpsState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FtpsStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FtpsStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FtpsState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FtpsStatePtrOutput struct{ *pulumi.OutputState }

func (FtpsStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FtpsState)(nil)).Elem()
}

func (o FtpsStatePtrOutput) ToFtpsStatePtrOutput() FtpsStatePtrOutput {
	return o
}

func (o FtpsStatePtrOutput) ToFtpsStatePtrOutputWithContext(ctx context.Context) FtpsStatePtrOutput {
	return o
}

func (o FtpsStatePtrOutput) Elem() FtpsStateOutput {
	return o.ApplyT(func(v *FtpsState) FtpsState {
		if v != nil {
			return *v
		}
		var ret FtpsState
		return ret
	}).(FtpsStateOutput)
}

func (o FtpsStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FtpsStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FtpsState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FtpsStateInput interface {
	pulumi.Input

	ToFtpsStateOutput() FtpsStateOutput
	ToFtpsStateOutputWithContext(context.Context) FtpsStateOutput
}

var ftpsStatePtrType = reflect.TypeOf((**FtpsState)(nil)).Elem()

type FtpsStatePtrInput interface {
	pulumi.Input

	ToFtpsStatePtrOutput() FtpsStatePtrOutput
	ToFtpsStatePtrOutputWithContext(context.Context) FtpsStatePtrOutput
}

type ftpsStatePtr string

func FtpsStatePtr(v string) FtpsStatePtrInput {
	return (*ftpsStatePtr)(&v)
}

func (*ftpsStatePtr) ElementType() reflect.Type {
	return ftpsStatePtrType
}

func (in *ftpsStatePtr) ToFtpsStatePtrOutput() FtpsStatePtrOutput {
	return pulumi.ToOutput(in).(FtpsStatePtrOutput)
}

func (in *ftpsStatePtr) ToFtpsStatePtrOutputWithContext(ctx context.Context) FtpsStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FtpsStatePtrOutput)
}

type HostNameType string

const (
	HostNameTypeVerified = HostNameType("Verified")
	HostNameTypeManaged  = HostNameType("Managed")
)

func (HostNameType) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameType)(nil)).Elem()
}

func (e HostNameType) ToHostNameTypeOutput() HostNameTypeOutput {
	return pulumi.ToOutput(e).(HostNameTypeOutput)
}

func (e HostNameType) ToHostNameTypeOutputWithContext(ctx context.Context) HostNameTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostNameTypeOutput)
}

func (e HostNameType) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return e.ToHostNameTypePtrOutputWithContext(context.Background())
}

func (e HostNameType) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return HostNameType(e).ToHostNameTypeOutputWithContext(ctx).ToHostNameTypePtrOutputWithContext(ctx)
}

func (e HostNameType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostNameType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostNameType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostNameType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostNameTypeOutput struct{ *pulumi.OutputState }

func (HostNameTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameType)(nil)).Elem()
}

func (o HostNameTypeOutput) ToHostNameTypeOutput() HostNameTypeOutput {
	return o
}

func (o HostNameTypeOutput) ToHostNameTypeOutputWithContext(ctx context.Context) HostNameTypeOutput {
	return o
}

func (o HostNameTypeOutput) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return o.ToHostNameTypePtrOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostNameType) *HostNameType {
		return &v
	}).(HostNameTypePtrOutput)
}

func (o HostNameTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostNameType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostNameTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostNameType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostNameTypePtrOutput struct{ *pulumi.OutputState }

func (HostNameTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostNameType)(nil)).Elem()
}

func (o HostNameTypePtrOutput) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return o
}

func (o HostNameTypePtrOutput) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return o
}

func (o HostNameTypePtrOutput) Elem() HostNameTypeOutput {
	return o.ApplyT(func(v *HostNameType) HostNameType {
		if v != nil {
			return *v
		}
		var ret HostNameType
		return ret
	}).(HostNameTypeOutput)
}

func (o HostNameTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostNameTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostNameType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostNameTypeInput interface {
	pulumi.Input

	ToHostNameTypeOutput() HostNameTypeOutput
	ToHostNameTypeOutputWithContext(context.Context) HostNameTypeOutput
}

var hostNameTypePtrType = reflect.TypeOf((**HostNameType)(nil)).Elem()

type HostNameTypePtrInput interface {
	pulumi.Input

	ToHostNameTypePtrOutput() HostNameTypePtrOutput
	ToHostNameTypePtrOutputWithContext(context.Context) HostNameTypePtrOutput
}

type hostNameTypePtr string

func HostNameTypePtr(v string) HostNameTypePtrInput {
	return (*hostNameTypePtr)(&v)
}

func (*hostNameTypePtr) ElementType() reflect.Type {
	return hostNameTypePtrType
}

func (in *hostNameTypePtr) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return pulumi.ToOutput(in).(HostNameTypePtrOutput)
}

func (in *hostNameTypePtr) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostNameTypePtrOutput)
}

type HostType string

const (
	HostTypeStandard   = HostType("Standard")
	HostTypeRepository = HostType("Repository")
)

func (HostType) ElementType() reflect.Type {
	return reflect.TypeOf((*HostType)(nil)).Elem()
}

func (e HostType) ToHostTypeOutput() HostTypeOutput {
	return pulumi.ToOutput(e).(HostTypeOutput)
}

func (e HostType) ToHostTypeOutputWithContext(ctx context.Context) HostTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostTypeOutput)
}

func (e HostType) ToHostTypePtrOutput() HostTypePtrOutput {
	return e.ToHostTypePtrOutputWithContext(context.Background())
}

func (e HostType) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return HostType(e).ToHostTypeOutputWithContext(ctx).ToHostTypePtrOutputWithContext(ctx)
}

func (e HostType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostTypeOutput struct{ *pulumi.OutputState }

func (HostTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostType)(nil)).Elem()
}

func (o HostTypeOutput) ToHostTypeOutput() HostTypeOutput {
	return o
}

func (o HostTypeOutput) ToHostTypeOutputWithContext(ctx context.Context) HostTypeOutput {
	return o
}

func (o HostTypeOutput) ToHostTypePtrOutput() HostTypePtrOutput {
	return o.ToHostTypePtrOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostType) *HostType {
		return &v
	}).(HostTypePtrOutput)
}

func (o HostTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostTypePtrOutput struct{ *pulumi.OutputState }

func (HostTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostType)(nil)).Elem()
}

func (o HostTypePtrOutput) ToHostTypePtrOutput() HostTypePtrOutput {
	return o
}

func (o HostTypePtrOutput) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return o
}

func (o HostTypePtrOutput) Elem() HostTypeOutput {
	return o.ApplyT(func(v *HostType) HostType {
		if v != nil {
			return *v
		}
		var ret HostType
		return ret
	}).(HostTypeOutput)
}

func (o HostTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostTypeInput interface {
	pulumi.Input

	ToHostTypeOutput() HostTypeOutput
	ToHostTypeOutputWithContext(context.Context) HostTypeOutput
}

var hostTypePtrType = reflect.TypeOf((**HostType)(nil)).Elem()

type HostTypePtrInput interface {
	pulumi.Input

	ToHostTypePtrOutput() HostTypePtrOutput
	ToHostTypePtrOutputWithContext(context.Context) HostTypePtrOutput
}

type hostTypePtr string

func HostTypePtr(v string) HostTypePtrInput {
	return (*hostTypePtr)(&v)
}

func (*hostTypePtr) ElementType() reflect.Type {
	return hostTypePtrType
}

func (in *hostTypePtr) ToHostTypePtrOutput() HostTypePtrOutput {
	return pulumi.ToOutput(in).(HostTypePtrOutput)
}

func (in *hostTypePtr) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostTypePtrOutput)
}

type InternalLoadBalancingMode string

const (
	InternalLoadBalancingModeNone       = InternalLoadBalancingMode("None")
	InternalLoadBalancingModeWeb        = InternalLoadBalancingMode("Web")
	InternalLoadBalancingModePublishing = InternalLoadBalancingMode("Publishing")
)

func (InternalLoadBalancingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*InternalLoadBalancingMode)(nil)).Elem()
}

func (e InternalLoadBalancingMode) ToInternalLoadBalancingModeOutput() InternalLoadBalancingModeOutput {
	return pulumi.ToOutput(e).(InternalLoadBalancingModeOutput)
}

func (e InternalLoadBalancingMode) ToInternalLoadBalancingModeOutputWithContext(ctx context.Context) InternalLoadBalancingModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InternalLoadBalancingModeOutput)
}

func (e InternalLoadBalancingMode) ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput {
	return e.ToInternalLoadBalancingModePtrOutputWithContext(context.Background())
}

func (e InternalLoadBalancingMode) ToInternalLoadBalancingModePtrOutputWithContext(ctx context.Context) InternalLoadBalancingModePtrOutput {
	return InternalLoadBalancingMode(e).ToInternalLoadBalancingModeOutputWithContext(ctx).ToInternalLoadBalancingModePtrOutputWithContext(ctx)
}

func (e InternalLoadBalancingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InternalLoadBalancingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InternalLoadBalancingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InternalLoadBalancingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InternalLoadBalancingModeOutput struct{ *pulumi.OutputState }

func (InternalLoadBalancingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InternalLoadBalancingMode)(nil)).Elem()
}

func (o InternalLoadBalancingModeOutput) ToInternalLoadBalancingModeOutput() InternalLoadBalancingModeOutput {
	return o
}

func (o InternalLoadBalancingModeOutput) ToInternalLoadBalancingModeOutputWithContext(ctx context.Context) InternalLoadBalancingModeOutput {
	return o
}

func (o InternalLoadBalancingModeOutput) ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput {
	return o.ToInternalLoadBalancingModePtrOutputWithContext(context.Background())
}

func (o InternalLoadBalancingModeOutput) ToInternalLoadBalancingModePtrOutputWithContext(ctx context.Context) InternalLoadBalancingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InternalLoadBalancingMode) *InternalLoadBalancingMode {
		return &v
	}).(InternalLoadBalancingModePtrOutput)
}

func (o InternalLoadBalancingModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InternalLoadBalancingModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InternalLoadBalancingMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InternalLoadBalancingModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InternalLoadBalancingModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InternalLoadBalancingMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InternalLoadBalancingModePtrOutput struct{ *pulumi.OutputState }

func (InternalLoadBalancingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternalLoadBalancingMode)(nil)).Elem()
}

func (o InternalLoadBalancingModePtrOutput) ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput {
	return o
}

func (o InternalLoadBalancingModePtrOutput) ToInternalLoadBalancingModePtrOutputWithContext(ctx context.Context) InternalLoadBalancingModePtrOutput {
	return o
}

func (o InternalLoadBalancingModePtrOutput) Elem() InternalLoadBalancingModeOutput {
	return o.ApplyT(func(v *InternalLoadBalancingMode) InternalLoadBalancingMode {
		if v != nil {
			return *v
		}
		var ret InternalLoadBalancingMode
		return ret
	}).(InternalLoadBalancingModeOutput)
}

func (o InternalLoadBalancingModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InternalLoadBalancingModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InternalLoadBalancingMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InternalLoadBalancingModeInput interface {
	pulumi.Input

	ToInternalLoadBalancingModeOutput() InternalLoadBalancingModeOutput
	ToInternalLoadBalancingModeOutputWithContext(context.Context) InternalLoadBalancingModeOutput
}

var internalLoadBalancingModePtrType = reflect.TypeOf((**InternalLoadBalancingMode)(nil)).Elem()

type InternalLoadBalancingModePtrInput interface {
	pulumi.Input

	ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput
	ToInternalLoadBalancingModePtrOutputWithContext(context.Context) InternalLoadBalancingModePtrOutput
}

type internalLoadBalancingModePtr string

func InternalLoadBalancingModePtr(v string) InternalLoadBalancingModePtrInput {
	return (*internalLoadBalancingModePtr)(&v)
}

func (*internalLoadBalancingModePtr) ElementType() reflect.Type {
	return internalLoadBalancingModePtrType
}

func (in *internalLoadBalancingModePtr) ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput {
	return pulumi.ToOutput(in).(InternalLoadBalancingModePtrOutput)
}

func (in *internalLoadBalancingModePtr) ToInternalLoadBalancingModePtrOutputWithContext(ctx context.Context) InternalLoadBalancingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InternalLoadBalancingModePtrOutput)
}

type IpFilterTag string

const (
	IpFilterTagDefault  = IpFilterTag("Default")
	IpFilterTagXffProxy = IpFilterTag("XffProxy")
)

func (IpFilterTag) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterTag)(nil)).Elem()
}

func (e IpFilterTag) ToIpFilterTagOutput() IpFilterTagOutput {
	return pulumi.ToOutput(e).(IpFilterTagOutput)
}

func (e IpFilterTag) ToIpFilterTagOutputWithContext(ctx context.Context) IpFilterTagOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpFilterTagOutput)
}

func (e IpFilterTag) ToIpFilterTagPtrOutput() IpFilterTagPtrOutput {
	return e.ToIpFilterTagPtrOutputWithContext(context.Background())
}

func (e IpFilterTag) ToIpFilterTagPtrOutputWithContext(ctx context.Context) IpFilterTagPtrOutput {
	return IpFilterTag(e).ToIpFilterTagOutputWithContext(ctx).ToIpFilterTagPtrOutputWithContext(ctx)
}

func (e IpFilterTag) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpFilterTag) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpFilterTag) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpFilterTag) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpFilterTagOutput struct{ *pulumi.OutputState }

func (IpFilterTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterTag)(nil)).Elem()
}

func (o IpFilterTagOutput) ToIpFilterTagOutput() IpFilterTagOutput {
	return o
}

func (o IpFilterTagOutput) ToIpFilterTagOutputWithContext(ctx context.Context) IpFilterTagOutput {
	return o
}

func (o IpFilterTagOutput) ToIpFilterTagPtrOutput() IpFilterTagPtrOutput {
	return o.ToIpFilterTagPtrOutputWithContext(context.Background())
}

func (o IpFilterTagOutput) ToIpFilterTagPtrOutputWithContext(ctx context.Context) IpFilterTagPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpFilterTag) *IpFilterTag {
		return &v
	}).(IpFilterTagPtrOutput)
}

func (o IpFilterTagOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpFilterTagOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpFilterTag) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpFilterTagOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpFilterTagOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpFilterTag) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpFilterTagPtrOutput struct{ *pulumi.OutputState }

func (IpFilterTagPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpFilterTag)(nil)).Elem()
}

func (o IpFilterTagPtrOutput) ToIpFilterTagPtrOutput() IpFilterTagPtrOutput {
	return o
}

func (o IpFilterTagPtrOutput) ToIpFilterTagPtrOutputWithContext(ctx context.Context) IpFilterTagPtrOutput {
	return o
}

func (o IpFilterTagPtrOutput) Elem() IpFilterTagOutput {
	return o.ApplyT(func(v *IpFilterTag) IpFilterTag {
		if v != nil {
			return *v
		}
		var ret IpFilterTag
		return ret
	}).(IpFilterTagOutput)
}

func (o IpFilterTagPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpFilterTagPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpFilterTag) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IpFilterTagInput interface {
	pulumi.Input

	ToIpFilterTagOutput() IpFilterTagOutput
	ToIpFilterTagOutputWithContext(context.Context) IpFilterTagOutput
}

var ipFilterTagPtrType = reflect.TypeOf((**IpFilterTag)(nil)).Elem()

type IpFilterTagPtrInput interface {
	pulumi.Input

	ToIpFilterTagPtrOutput() IpFilterTagPtrOutput
	ToIpFilterTagPtrOutputWithContext(context.Context) IpFilterTagPtrOutput
}

type ipFilterTagPtr string

func IpFilterTagPtr(v string) IpFilterTagPtrInput {
	return (*ipFilterTagPtr)(&v)
}

func (*ipFilterTagPtr) ElementType() reflect.Type {
	return ipFilterTagPtrType
}

func (in *ipFilterTagPtr) ToIpFilterTagPtrOutput() IpFilterTagPtrOutput {
	return pulumi.ToOutput(in).(IpFilterTagPtrOutput)
}

func (in *ipFilterTagPtr) ToIpFilterTagPtrOutputWithContext(ctx context.Context) IpFilterTagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpFilterTagPtrOutput)
}

type LogLevel string

const (
	LogLevelOff         = LogLevel("Off")
	LogLevelVerbose     = LogLevel("Verbose")
	LogLevelInformation = LogLevel("Information")
	LogLevelWarning     = LogLevel("Warning")
	LogLevelError       = LogLevel("Error")
)

func (LogLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*LogLevel)(nil)).Elem()
}

func (e LogLevel) ToLogLevelOutput() LogLevelOutput {
	return pulumi.ToOutput(e).(LogLevelOutput)
}

func (e LogLevel) ToLogLevelOutputWithContext(ctx context.Context) LogLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LogLevelOutput)
}

func (e LogLevel) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return e.ToLogLevelPtrOutputWithContext(context.Background())
}

func (e LogLevel) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return LogLevel(e).ToLogLevelOutputWithContext(ctx).ToLogLevelPtrOutputWithContext(ctx)
}

func (e LogLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LogLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LogLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LogLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LogLevelOutput struct{ *pulumi.OutputState }

func (LogLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogLevel)(nil)).Elem()
}

func (o LogLevelOutput) ToLogLevelOutput() LogLevelOutput {
	return o
}

func (o LogLevelOutput) ToLogLevelOutputWithContext(ctx context.Context) LogLevelOutput {
	return o
}

func (o LogLevelOutput) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return o.ToLogLevelPtrOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogLevel) *LogLevel {
		return &v
	}).(LogLevelPtrOutput)
}

func (o LogLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LogLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LogLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LogLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LogLevelPtrOutput struct{ *pulumi.OutputState }

func (LogLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogLevel)(nil)).Elem()
}

func (o LogLevelPtrOutput) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return o
}

func (o LogLevelPtrOutput) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return o
}

func (o LogLevelPtrOutput) Elem() LogLevelOutput {
	return o.ApplyT(func(v *LogLevel) LogLevel {
		if v != nil {
			return *v
		}
		var ret LogLevel
		return ret
	}).(LogLevelOutput)
}

func (o LogLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LogLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LogLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LogLevelInput interface {
	pulumi.Input

	ToLogLevelOutput() LogLevelOutput
	ToLogLevelOutputWithContext(context.Context) LogLevelOutput
}

var logLevelPtrType = reflect.TypeOf((**LogLevel)(nil)).Elem()

type LogLevelPtrInput interface {
	pulumi.Input

	ToLogLevelPtrOutput() LogLevelPtrOutput
	ToLogLevelPtrOutputWithContext(context.Context) LogLevelPtrOutput
}

type logLevelPtr string

func LogLevelPtr(v string) LogLevelPtrInput {
	return (*logLevelPtr)(&v)
}

func (*logLevelPtr) ElementType() reflect.Type {
	return logLevelPtrType
}

func (in *logLevelPtr) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return pulumi.ToOutput(in).(LogLevelPtrOutput)
}

func (in *logLevelPtr) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LogLevelPtrOutput)
}

type ManagedPipelineMode string

const (
	ManagedPipelineModeIntegrated = ManagedPipelineMode("Integrated")
	ManagedPipelineModeClassic    = ManagedPipelineMode("Classic")
)

func (ManagedPipelineMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPipelineMode)(nil)).Elem()
}

func (e ManagedPipelineMode) ToManagedPipelineModeOutput() ManagedPipelineModeOutput {
	return pulumi.ToOutput(e).(ManagedPipelineModeOutput)
}

func (e ManagedPipelineMode) ToManagedPipelineModeOutputWithContext(ctx context.Context) ManagedPipelineModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedPipelineModeOutput)
}

func (e ManagedPipelineMode) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return e.ToManagedPipelineModePtrOutputWithContext(context.Background())
}

func (e ManagedPipelineMode) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return ManagedPipelineMode(e).ToManagedPipelineModeOutputWithContext(ctx).ToManagedPipelineModePtrOutputWithContext(ctx)
}

func (e ManagedPipelineMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedPipelineMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedPipelineMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedPipelineMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedPipelineModeOutput struct{ *pulumi.OutputState }

func (ManagedPipelineModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPipelineMode)(nil)).Elem()
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModeOutput() ManagedPipelineModeOutput {
	return o
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModeOutputWithContext(ctx context.Context) ManagedPipelineModeOutput {
	return o
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return o.ToManagedPipelineModePtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedPipelineMode) *ManagedPipelineMode {
		return &v
	}).(ManagedPipelineModePtrOutput)
}

func (o ManagedPipelineModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedPipelineMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedPipelineModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedPipelineMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedPipelineModePtrOutput struct{ *pulumi.OutputState }

func (ManagedPipelineModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPipelineMode)(nil)).Elem()
}

func (o ManagedPipelineModePtrOutput) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return o
}

func (o ManagedPipelineModePtrOutput) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return o
}

func (o ManagedPipelineModePtrOutput) Elem() ManagedPipelineModeOutput {
	return o.ApplyT(func(v *ManagedPipelineMode) ManagedPipelineMode {
		if v != nil {
			return *v
		}
		var ret ManagedPipelineMode
		return ret
	}).(ManagedPipelineModeOutput)
}

func (o ManagedPipelineModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedPipelineMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedPipelineModeInput interface {
	pulumi.Input

	ToManagedPipelineModeOutput() ManagedPipelineModeOutput
	ToManagedPipelineModeOutputWithContext(context.Context) ManagedPipelineModeOutput
}

var managedPipelineModePtrType = reflect.TypeOf((**ManagedPipelineMode)(nil)).Elem()

type ManagedPipelineModePtrInput interface {
	pulumi.Input

	ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput
	ToManagedPipelineModePtrOutputWithContext(context.Context) ManagedPipelineModePtrOutput
}

type managedPipelineModePtr string

func ManagedPipelineModePtr(v string) ManagedPipelineModePtrInput {
	return (*managedPipelineModePtr)(&v)
}

func (*managedPipelineModePtr) ElementType() reflect.Type {
	return managedPipelineModePtrType
}

func (in *managedPipelineModePtr) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return pulumi.ToOutput(in).(ManagedPipelineModePtrOutput)
}

func (in *managedPipelineModePtr) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedPipelineModePtrOutput)
}

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
)

func (ManagedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return e.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return ManagedServiceIdentityType(e).ToManagedServiceIdentityTypeOutputWithContext(ctx).ToManagedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityType) *ManagedServiceIdentityType {
		return &v
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) Elem() ManagedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityType) ManagedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityType
		return ret
	}).(ManagedServiceIdentityTypeOutput)
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedServiceIdentityTypeInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput
	ToManagedServiceIdentityTypeOutputWithContext(context.Context) ManagedServiceIdentityTypeOutput
}

var managedServiceIdentityTypePtrType = reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()

type ManagedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput
	ToManagedServiceIdentityTypePtrOutputWithContext(context.Context) ManagedServiceIdentityTypePtrOutput
}

type managedServiceIdentityTypePtr string

func ManagedServiceIdentityTypePtr(v string) ManagedServiceIdentityTypePtrInput {
	return (*managedServiceIdentityTypePtr)(&v)
}

func (*managedServiceIdentityTypePtr) ElementType() reflect.Type {
	return managedServiceIdentityTypePtrType
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedServiceIdentityTypePtrOutput)
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedServiceIdentityTypePtrOutput)
}

type PublicCertificateLocation string

const (
	PublicCertificateLocationCurrentUserMy  = PublicCertificateLocation("CurrentUserMy")
	PublicCertificateLocationLocalMachineMy = PublicCertificateLocation("LocalMachineMy")
	PublicCertificateLocationUnknown        = PublicCertificateLocation("Unknown")
)

func (PublicCertificateLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicCertificateLocation)(nil)).Elem()
}

func (e PublicCertificateLocation) ToPublicCertificateLocationOutput() PublicCertificateLocationOutput {
	return pulumi.ToOutput(e).(PublicCertificateLocationOutput)
}

func (e PublicCertificateLocation) ToPublicCertificateLocationOutputWithContext(ctx context.Context) PublicCertificateLocationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicCertificateLocationOutput)
}

func (e PublicCertificateLocation) ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput {
	return e.ToPublicCertificateLocationPtrOutputWithContext(context.Background())
}

func (e PublicCertificateLocation) ToPublicCertificateLocationPtrOutputWithContext(ctx context.Context) PublicCertificateLocationPtrOutput {
	return PublicCertificateLocation(e).ToPublicCertificateLocationOutputWithContext(ctx).ToPublicCertificateLocationPtrOutputWithContext(ctx)
}

func (e PublicCertificateLocation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicCertificateLocation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicCertificateLocation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicCertificateLocation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicCertificateLocationOutput struct{ *pulumi.OutputState }

func (PublicCertificateLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicCertificateLocation)(nil)).Elem()
}

func (o PublicCertificateLocationOutput) ToPublicCertificateLocationOutput() PublicCertificateLocationOutput {
	return o
}

func (o PublicCertificateLocationOutput) ToPublicCertificateLocationOutputWithContext(ctx context.Context) PublicCertificateLocationOutput {
	return o
}

func (o PublicCertificateLocationOutput) ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput {
	return o.ToPublicCertificateLocationPtrOutputWithContext(context.Background())
}

func (o PublicCertificateLocationOutput) ToPublicCertificateLocationPtrOutputWithContext(ctx context.Context) PublicCertificateLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicCertificateLocation) *PublicCertificateLocation {
		return &v
	}).(PublicCertificateLocationPtrOutput)
}

func (o PublicCertificateLocationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicCertificateLocationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicCertificateLocation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicCertificateLocationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicCertificateLocationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicCertificateLocation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicCertificateLocationPtrOutput struct{ *pulumi.OutputState }

func (PublicCertificateLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicCertificateLocation)(nil)).Elem()
}

func (o PublicCertificateLocationPtrOutput) ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput {
	return o
}

func (o PublicCertificateLocationPtrOutput) ToPublicCertificateLocationPtrOutputWithContext(ctx context.Context) PublicCertificateLocationPtrOutput {
	return o
}

func (o PublicCertificateLocationPtrOutput) Elem() PublicCertificateLocationOutput {
	return o.ApplyT(func(v *PublicCertificateLocation) PublicCertificateLocation {
		if v != nil {
			return *v
		}
		var ret PublicCertificateLocation
		return ret
	}).(PublicCertificateLocationOutput)
}

func (o PublicCertificateLocationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicCertificateLocationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicCertificateLocation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicCertificateLocationInput interface {
	pulumi.Input

	ToPublicCertificateLocationOutput() PublicCertificateLocationOutput
	ToPublicCertificateLocationOutputWithContext(context.Context) PublicCertificateLocationOutput
}

var publicCertificateLocationPtrType = reflect.TypeOf((**PublicCertificateLocation)(nil)).Elem()

type PublicCertificateLocationPtrInput interface {
	pulumi.Input

	ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput
	ToPublicCertificateLocationPtrOutputWithContext(context.Context) PublicCertificateLocationPtrOutput
}

type publicCertificateLocationPtr string

func PublicCertificateLocationPtr(v string) PublicCertificateLocationPtrInput {
	return (*publicCertificateLocationPtr)(&v)
}

func (*publicCertificateLocationPtr) ElementType() reflect.Type {
	return publicCertificateLocationPtrType
}

func (in *publicCertificateLocationPtr) ToPublicCertificateLocationPtrOutput() PublicCertificateLocationPtrOutput {
	return pulumi.ToOutput(in).(PublicCertificateLocationPtrOutput)
}

func (in *publicCertificateLocationPtr) ToPublicCertificateLocationPtrOutputWithContext(ctx context.Context) PublicCertificateLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicCertificateLocationPtrOutput)
}

type RedundancyMode string

const (
	RedundancyModeNone         = RedundancyMode("None")
	RedundancyModeManual       = RedundancyMode("Manual")
	RedundancyModeFailover     = RedundancyMode("Failover")
	RedundancyModeActiveActive = RedundancyMode("ActiveActive")
	RedundancyModeGeoRedundant = RedundancyMode("GeoRedundant")
)

func (RedundancyMode) ElementType() reflect.Type {
	return reflect.TypeOf((*RedundancyMode)(nil)).Elem()
}

func (e RedundancyMode) ToRedundancyModeOutput() RedundancyModeOutput {
	return pulumi.ToOutput(e).(RedundancyModeOutput)
}

func (e RedundancyMode) ToRedundancyModeOutputWithContext(ctx context.Context) RedundancyModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RedundancyModeOutput)
}

func (e RedundancyMode) ToRedundancyModePtrOutput() RedundancyModePtrOutput {
	return e.ToRedundancyModePtrOutputWithContext(context.Background())
}

func (e RedundancyMode) ToRedundancyModePtrOutputWithContext(ctx context.Context) RedundancyModePtrOutput {
	return RedundancyMode(e).ToRedundancyModeOutputWithContext(ctx).ToRedundancyModePtrOutputWithContext(ctx)
}

func (e RedundancyMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RedundancyMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RedundancyMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RedundancyMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RedundancyModeOutput struct{ *pulumi.OutputState }

func (RedundancyModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedundancyMode)(nil)).Elem()
}

func (o RedundancyModeOutput) ToRedundancyModeOutput() RedundancyModeOutput {
	return o
}

func (o RedundancyModeOutput) ToRedundancyModeOutputWithContext(ctx context.Context) RedundancyModeOutput {
	return o
}

func (o RedundancyModeOutput) ToRedundancyModePtrOutput() RedundancyModePtrOutput {
	return o.ToRedundancyModePtrOutputWithContext(context.Background())
}

func (o RedundancyModeOutput) ToRedundancyModePtrOutputWithContext(ctx context.Context) RedundancyModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RedundancyMode) *RedundancyMode {
		return &v
	}).(RedundancyModePtrOutput)
}

func (o RedundancyModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RedundancyModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RedundancyMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RedundancyModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RedundancyModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RedundancyMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RedundancyModePtrOutput struct{ *pulumi.OutputState }

func (RedundancyModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedundancyMode)(nil)).Elem()
}

func (o RedundancyModePtrOutput) ToRedundancyModePtrOutput() RedundancyModePtrOutput {
	return o
}

func (o RedundancyModePtrOutput) ToRedundancyModePtrOutputWithContext(ctx context.Context) RedundancyModePtrOutput {
	return o
}

func (o RedundancyModePtrOutput) Elem() RedundancyModeOutput {
	return o.ApplyT(func(v *RedundancyMode) RedundancyMode {
		if v != nil {
			return *v
		}
		var ret RedundancyMode
		return ret
	}).(RedundancyModeOutput)
}

func (o RedundancyModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RedundancyModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RedundancyMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RedundancyModeInput interface {
	pulumi.Input

	ToRedundancyModeOutput() RedundancyModeOutput
	ToRedundancyModeOutputWithContext(context.Context) RedundancyModeOutput
}

var redundancyModePtrType = reflect.TypeOf((**RedundancyMode)(nil)).Elem()

type RedundancyModePtrInput interface {
	pulumi.Input

	ToRedundancyModePtrOutput() RedundancyModePtrOutput
	ToRedundancyModePtrOutputWithContext(context.Context) RedundancyModePtrOutput
}

type redundancyModePtr string

func RedundancyModePtr(v string) RedundancyModePtrInput {
	return (*redundancyModePtr)(&v)
}

func (*redundancyModePtr) ElementType() reflect.Type {
	return redundancyModePtrType
}

func (in *redundancyModePtr) ToRedundancyModePtrOutput() RedundancyModePtrOutput {
	return pulumi.ToOutput(in).(RedundancyModePtrOutput)
}

func (in *redundancyModePtr) ToRedundancyModePtrOutputWithContext(ctx context.Context) RedundancyModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RedundancyModePtrOutput)
}

type RouteType string

const (
	RouteTypeDEFAULT   = RouteType("DEFAULT")
	RouteTypeINHERITED = RouteType("INHERITED")
	RouteTypeSTATIC    = RouteType("STATIC")
)

func (RouteType) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteType)(nil)).Elem()
}

func (e RouteType) ToRouteTypeOutput() RouteTypeOutput {
	return pulumi.ToOutput(e).(RouteTypeOutput)
}

func (e RouteType) ToRouteTypeOutputWithContext(ctx context.Context) RouteTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RouteTypeOutput)
}

func (e RouteType) ToRouteTypePtrOutput() RouteTypePtrOutput {
	return e.ToRouteTypePtrOutputWithContext(context.Background())
}

func (e RouteType) ToRouteTypePtrOutputWithContext(ctx context.Context) RouteTypePtrOutput {
	return RouteType(e).ToRouteTypeOutputWithContext(ctx).ToRouteTypePtrOutputWithContext(ctx)
}

func (e RouteType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RouteType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RouteType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RouteType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RouteTypeOutput struct{ *pulumi.OutputState }

func (RouteTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteType)(nil)).Elem()
}

func (o RouteTypeOutput) ToRouteTypeOutput() RouteTypeOutput {
	return o
}

func (o RouteTypeOutput) ToRouteTypeOutputWithContext(ctx context.Context) RouteTypeOutput {
	return o
}

func (o RouteTypeOutput) ToRouteTypePtrOutput() RouteTypePtrOutput {
	return o.ToRouteTypePtrOutputWithContext(context.Background())
}

func (o RouteTypeOutput) ToRouteTypePtrOutputWithContext(ctx context.Context) RouteTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteType) *RouteType {
		return &v
	}).(RouteTypePtrOutput)
}

func (o RouteTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RouteTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RouteType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RouteTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RouteTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RouteType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RouteTypePtrOutput struct{ *pulumi.OutputState }

func (RouteTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteType)(nil)).Elem()
}

func (o RouteTypePtrOutput) ToRouteTypePtrOutput() RouteTypePtrOutput {
	return o
}

func (o RouteTypePtrOutput) ToRouteTypePtrOutputWithContext(ctx context.Context) RouteTypePtrOutput {
	return o
}

func (o RouteTypePtrOutput) Elem() RouteTypeOutput {
	return o.ApplyT(func(v *RouteType) RouteType {
		if v != nil {
			return *v
		}
		var ret RouteType
		return ret
	}).(RouteTypeOutput)
}

func (o RouteTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RouteTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RouteType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RouteTypeInput interface {
	pulumi.Input

	ToRouteTypeOutput() RouteTypeOutput
	ToRouteTypeOutputWithContext(context.Context) RouteTypeOutput
}

var routeTypePtrType = reflect.TypeOf((**RouteType)(nil)).Elem()

type RouteTypePtrInput interface {
	pulumi.Input

	ToRouteTypePtrOutput() RouteTypePtrOutput
	ToRouteTypePtrOutputWithContext(context.Context) RouteTypePtrOutput
}

type routeTypePtr string

func RouteTypePtr(v string) RouteTypePtrInput {
	return (*routeTypePtr)(&v)
}

func (*routeTypePtr) ElementType() reflect.Type {
	return routeTypePtrType
}

func (in *routeTypePtr) ToRouteTypePtrOutput() RouteTypePtrOutput {
	return pulumi.ToOutput(in).(RouteTypePtrOutput)
}

func (in *routeTypePtr) ToRouteTypePtrOutputWithContext(ctx context.Context) RouteTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RouteTypePtrOutput)
}

type ScmType string

const (
	ScmTypeNone         = ScmType("None")
	ScmTypeDropbox      = ScmType("Dropbox")
	ScmTypeTfs          = ScmType("Tfs")
	ScmTypeLocalGit     = ScmType("LocalGit")
	ScmTypeGitHub       = ScmType("GitHub")
	ScmTypeCodePlexGit  = ScmType("CodePlexGit")
	ScmTypeCodePlexHg   = ScmType("CodePlexHg")
	ScmTypeBitbucketGit = ScmType("BitbucketGit")
	ScmTypeBitbucketHg  = ScmType("BitbucketHg")
	ScmTypeExternalGit  = ScmType("ExternalGit")
	ScmTypeExternalHg   = ScmType("ExternalHg")
	ScmTypeOneDrive     = ScmType("OneDrive")
	ScmTypeVSO          = ScmType("VSO")
)

func (ScmType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScmType)(nil)).Elem()
}

func (e ScmType) ToScmTypeOutput() ScmTypeOutput {
	return pulumi.ToOutput(e).(ScmTypeOutput)
}

func (e ScmType) ToScmTypeOutputWithContext(ctx context.Context) ScmTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScmTypeOutput)
}

func (e ScmType) ToScmTypePtrOutput() ScmTypePtrOutput {
	return e.ToScmTypePtrOutputWithContext(context.Background())
}

func (e ScmType) ToScmTypePtrOutputWithContext(ctx context.Context) ScmTypePtrOutput {
	return ScmType(e).ToScmTypeOutputWithContext(ctx).ToScmTypePtrOutputWithContext(ctx)
}

func (e ScmType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScmType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScmType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScmType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScmTypeOutput struct{ *pulumi.OutputState }

func (ScmTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScmType)(nil)).Elem()
}

func (o ScmTypeOutput) ToScmTypeOutput() ScmTypeOutput {
	return o
}

func (o ScmTypeOutput) ToScmTypeOutputWithContext(ctx context.Context) ScmTypeOutput {
	return o
}

func (o ScmTypeOutput) ToScmTypePtrOutput() ScmTypePtrOutput {
	return o.ToScmTypePtrOutputWithContext(context.Background())
}

func (o ScmTypeOutput) ToScmTypePtrOutputWithContext(ctx context.Context) ScmTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScmType) *ScmType {
		return &v
	}).(ScmTypePtrOutput)
}

func (o ScmTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScmTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScmType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScmTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScmTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScmType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScmTypePtrOutput struct{ *pulumi.OutputState }

func (ScmTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScmType)(nil)).Elem()
}

func (o ScmTypePtrOutput) ToScmTypePtrOutput() ScmTypePtrOutput {
	return o
}

func (o ScmTypePtrOutput) ToScmTypePtrOutputWithContext(ctx context.Context) ScmTypePtrOutput {
	return o
}

func (o ScmTypePtrOutput) Elem() ScmTypeOutput {
	return o.ApplyT(func(v *ScmType) ScmType {
		if v != nil {
			return *v
		}
		var ret ScmType
		return ret
	}).(ScmTypeOutput)
}

func (o ScmTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScmTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScmType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScmTypeInput interface {
	pulumi.Input

	ToScmTypeOutput() ScmTypeOutput
	ToScmTypeOutputWithContext(context.Context) ScmTypeOutput
}

var scmTypePtrType = reflect.TypeOf((**ScmType)(nil)).Elem()

type ScmTypePtrInput interface {
	pulumi.Input

	ToScmTypePtrOutput() ScmTypePtrOutput
	ToScmTypePtrOutputWithContext(context.Context) ScmTypePtrOutput
}

type scmTypePtr string

func ScmTypePtr(v string) ScmTypePtrInput {
	return (*scmTypePtr)(&v)
}

func (*scmTypePtr) ElementType() reflect.Type {
	return scmTypePtrType
}

func (in *scmTypePtr) ToScmTypePtrOutput() ScmTypePtrOutput {
	return pulumi.ToOutput(in).(ScmTypePtrOutput)
}

func (in *scmTypePtr) ToScmTypePtrOutputWithContext(ctx context.Context) ScmTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScmTypePtrOutput)
}

type SiteLoadBalancing string

const (
	SiteLoadBalancingWeightedRoundRobin   = SiteLoadBalancing("WeightedRoundRobin")
	SiteLoadBalancingLeastRequests        = SiteLoadBalancing("LeastRequests")
	SiteLoadBalancingLeastResponseTime    = SiteLoadBalancing("LeastResponseTime")
	SiteLoadBalancingWeightedTotalTraffic = SiteLoadBalancing("WeightedTotalTraffic")
	SiteLoadBalancingRequestHash          = SiteLoadBalancing("RequestHash")
)

func (SiteLoadBalancing) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLoadBalancing)(nil)).Elem()
}

func (e SiteLoadBalancing) ToSiteLoadBalancingOutput() SiteLoadBalancingOutput {
	return pulumi.ToOutput(e).(SiteLoadBalancingOutput)
}

func (e SiteLoadBalancing) ToSiteLoadBalancingOutputWithContext(ctx context.Context) SiteLoadBalancingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SiteLoadBalancingOutput)
}

func (e SiteLoadBalancing) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return e.ToSiteLoadBalancingPtrOutputWithContext(context.Background())
}

func (e SiteLoadBalancing) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return SiteLoadBalancing(e).ToSiteLoadBalancingOutputWithContext(ctx).ToSiteLoadBalancingPtrOutputWithContext(ctx)
}

func (e SiteLoadBalancing) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SiteLoadBalancing) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SiteLoadBalancing) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SiteLoadBalancing) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SiteLoadBalancingOutput struct{ *pulumi.OutputState }

func (SiteLoadBalancingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLoadBalancing)(nil)).Elem()
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingOutput() SiteLoadBalancingOutput {
	return o
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingOutputWithContext(ctx context.Context) SiteLoadBalancingOutput {
	return o
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return o.ToSiteLoadBalancingPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteLoadBalancing) *SiteLoadBalancing {
		return &v
	}).(SiteLoadBalancingPtrOutput)
}

func (o SiteLoadBalancingOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SiteLoadBalancing) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SiteLoadBalancingOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SiteLoadBalancing) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SiteLoadBalancingPtrOutput struct{ *pulumi.OutputState }

func (SiteLoadBalancingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLoadBalancing)(nil)).Elem()
}

func (o SiteLoadBalancingPtrOutput) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return o
}

func (o SiteLoadBalancingPtrOutput) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return o
}

func (o SiteLoadBalancingPtrOutput) Elem() SiteLoadBalancingOutput {
	return o.ApplyT(func(v *SiteLoadBalancing) SiteLoadBalancing {
		if v != nil {
			return *v
		}
		var ret SiteLoadBalancing
		return ret
	}).(SiteLoadBalancingOutput)
}

func (o SiteLoadBalancingPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SiteLoadBalancing) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SiteLoadBalancingInput interface {
	pulumi.Input

	ToSiteLoadBalancingOutput() SiteLoadBalancingOutput
	ToSiteLoadBalancingOutputWithContext(context.Context) SiteLoadBalancingOutput
}

var siteLoadBalancingPtrType = reflect.TypeOf((**SiteLoadBalancing)(nil)).Elem()

type SiteLoadBalancingPtrInput interface {
	pulumi.Input

	ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput
	ToSiteLoadBalancingPtrOutputWithContext(context.Context) SiteLoadBalancingPtrOutput
}

type siteLoadBalancingPtr string

func SiteLoadBalancingPtr(v string) SiteLoadBalancingPtrInput {
	return (*siteLoadBalancingPtr)(&v)
}

func (*siteLoadBalancingPtr) ElementType() reflect.Type {
	return siteLoadBalancingPtrType
}

func (in *siteLoadBalancingPtr) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return pulumi.ToOutput(in).(SiteLoadBalancingPtrOutput)
}

func (in *siteLoadBalancingPtr) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SiteLoadBalancingPtrOutput)
}

type SslState string

const (
	SslStateDisabled       = SslState("Disabled")
	SslStateSniEnabled     = SslState("SniEnabled")
	SslStateIpBasedEnabled = SslState("IpBasedEnabled")
)

func (SslState) ElementType() reflect.Type {
	return reflect.TypeOf((*SslState)(nil)).Elem()
}

func (e SslState) ToSslStateOutput() SslStateOutput {
	return pulumi.ToOutput(e).(SslStateOutput)
}

func (e SslState) ToSslStateOutputWithContext(ctx context.Context) SslStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslStateOutput)
}

func (e SslState) ToSslStatePtrOutput() SslStatePtrOutput {
	return e.ToSslStatePtrOutputWithContext(context.Background())
}

func (e SslState) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return SslState(e).ToSslStateOutputWithContext(ctx).ToSslStatePtrOutputWithContext(ctx)
}

func (e SslState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslStateOutput struct{ *pulumi.OutputState }

func (SslStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslState)(nil)).Elem()
}

func (o SslStateOutput) ToSslStateOutput() SslStateOutput {
	return o
}

func (o SslStateOutput) ToSslStateOutputWithContext(ctx context.Context) SslStateOutput {
	return o
}

func (o SslStateOutput) ToSslStatePtrOutput() SslStatePtrOutput {
	return o.ToSslStatePtrOutputWithContext(context.Background())
}

func (o SslStateOutput) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslState) *SslState {
		return &v
	}).(SslStatePtrOutput)
}

func (o SslStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslStatePtrOutput struct{ *pulumi.OutputState }

func (SslStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslState)(nil)).Elem()
}

func (o SslStatePtrOutput) ToSslStatePtrOutput() SslStatePtrOutput {
	return o
}

func (o SslStatePtrOutput) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return o
}

func (o SslStatePtrOutput) Elem() SslStateOutput {
	return o.ApplyT(func(v *SslState) SslState {
		if v != nil {
			return *v
		}
		var ret SslState
		return ret
	}).(SslStateOutput)
}

func (o SslStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SslStateInput interface {
	pulumi.Input

	ToSslStateOutput() SslStateOutput
	ToSslStateOutputWithContext(context.Context) SslStateOutput
}

var sslStatePtrType = reflect.TypeOf((**SslState)(nil)).Elem()

type SslStatePtrInput interface {
	pulumi.Input

	ToSslStatePtrOutput() SslStatePtrOutput
	ToSslStatePtrOutputWithContext(context.Context) SslStatePtrOutput
}

type sslStatePtr string

func SslStatePtr(v string) SslStatePtrInput {
	return (*sslStatePtr)(&v)
}

func (*sslStatePtr) ElementType() reflect.Type {
	return sslStatePtrType
}

func (in *sslStatePtr) ToSslStatePtrOutput() SslStatePtrOutput {
	return pulumi.ToOutput(in).(SslStatePtrOutput)
}

func (in *sslStatePtr) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslStatePtrOutput)
}

type SupportedTlsVersions string

const (
	SupportedTlsVersions_1_0 = SupportedTlsVersions("1.0")
	SupportedTlsVersions_1_1 = SupportedTlsVersions("1.1")
	SupportedTlsVersions_1_2 = SupportedTlsVersions("1.2")
)

func (SupportedTlsVersions) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedTlsVersions)(nil)).Elem()
}

func (e SupportedTlsVersions) ToSupportedTlsVersionsOutput() SupportedTlsVersionsOutput {
	return pulumi.ToOutput(e).(SupportedTlsVersionsOutput)
}

func (e SupportedTlsVersions) ToSupportedTlsVersionsOutputWithContext(ctx context.Context) SupportedTlsVersionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SupportedTlsVersionsOutput)
}

func (e SupportedTlsVersions) ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput {
	return e.ToSupportedTlsVersionsPtrOutputWithContext(context.Background())
}

func (e SupportedTlsVersions) ToSupportedTlsVersionsPtrOutputWithContext(ctx context.Context) SupportedTlsVersionsPtrOutput {
	return SupportedTlsVersions(e).ToSupportedTlsVersionsOutputWithContext(ctx).ToSupportedTlsVersionsPtrOutputWithContext(ctx)
}

func (e SupportedTlsVersions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedTlsVersions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedTlsVersions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SupportedTlsVersions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SupportedTlsVersionsOutput struct{ *pulumi.OutputState }

func (SupportedTlsVersionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedTlsVersions)(nil)).Elem()
}

func (o SupportedTlsVersionsOutput) ToSupportedTlsVersionsOutput() SupportedTlsVersionsOutput {
	return o
}

func (o SupportedTlsVersionsOutput) ToSupportedTlsVersionsOutputWithContext(ctx context.Context) SupportedTlsVersionsOutput {
	return o
}

func (o SupportedTlsVersionsOutput) ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput {
	return o.ToSupportedTlsVersionsPtrOutputWithContext(context.Background())
}

func (o SupportedTlsVersionsOutput) ToSupportedTlsVersionsPtrOutputWithContext(ctx context.Context) SupportedTlsVersionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportedTlsVersions) *SupportedTlsVersions {
		return &v
	}).(SupportedTlsVersionsPtrOutput)
}

func (o SupportedTlsVersionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SupportedTlsVersionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedTlsVersions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SupportedTlsVersionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedTlsVersionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedTlsVersions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SupportedTlsVersionsPtrOutput struct{ *pulumi.OutputState }

func (SupportedTlsVersionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedTlsVersions)(nil)).Elem()
}

func (o SupportedTlsVersionsPtrOutput) ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput {
	return o
}

func (o SupportedTlsVersionsPtrOutput) ToSupportedTlsVersionsPtrOutputWithContext(ctx context.Context) SupportedTlsVersionsPtrOutput {
	return o
}

func (o SupportedTlsVersionsPtrOutput) Elem() SupportedTlsVersionsOutput {
	return o.ApplyT(func(v *SupportedTlsVersions) SupportedTlsVersions {
		if v != nil {
			return *v
		}
		var ret SupportedTlsVersions
		return ret
	}).(SupportedTlsVersionsOutput)
}

func (o SupportedTlsVersionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedTlsVersionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SupportedTlsVersions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SupportedTlsVersionsInput interface {
	pulumi.Input

	ToSupportedTlsVersionsOutput() SupportedTlsVersionsOutput
	ToSupportedTlsVersionsOutputWithContext(context.Context) SupportedTlsVersionsOutput
}

var supportedTlsVersionsPtrType = reflect.TypeOf((**SupportedTlsVersions)(nil)).Elem()

type SupportedTlsVersionsPtrInput interface {
	pulumi.Input

	ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput
	ToSupportedTlsVersionsPtrOutputWithContext(context.Context) SupportedTlsVersionsPtrOutput
}

type supportedTlsVersionsPtr string

func SupportedTlsVersionsPtr(v string) SupportedTlsVersionsPtrInput {
	return (*supportedTlsVersionsPtr)(&v)
}

func (*supportedTlsVersionsPtr) ElementType() reflect.Type {
	return supportedTlsVersionsPtrType
}

func (in *supportedTlsVersionsPtr) ToSupportedTlsVersionsPtrOutput() SupportedTlsVersionsPtrOutput {
	return pulumi.ToOutput(in).(SupportedTlsVersionsPtrOutput)
}

func (in *supportedTlsVersionsPtr) ToSupportedTlsVersionsPtrOutputWithContext(ctx context.Context) SupportedTlsVersionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SupportedTlsVersionsPtrOutput)
}

type UnauthenticatedClientAction string

const (
	UnauthenticatedClientActionRedirectToLoginPage = UnauthenticatedClientAction("RedirectToLoginPage")
	UnauthenticatedClientActionAllowAnonymous      = UnauthenticatedClientAction("AllowAnonymous")
)

func (UnauthenticatedClientAction) ElementType() reflect.Type {
	return reflect.TypeOf((*UnauthenticatedClientAction)(nil)).Elem()
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput {
	return pulumi.ToOutput(e).(UnauthenticatedClientActionOutput)
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionOutputWithContext(ctx context.Context) UnauthenticatedClientActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UnauthenticatedClientActionOutput)
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return e.ToUnauthenticatedClientActionPtrOutputWithContext(context.Background())
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return UnauthenticatedClientAction(e).ToUnauthenticatedClientActionOutputWithContext(ctx).ToUnauthenticatedClientActionPtrOutputWithContext(ctx)
}

func (e UnauthenticatedClientAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnauthenticatedClientAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnauthenticatedClientAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UnauthenticatedClientAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UnauthenticatedClientActionOutput struct{ *pulumi.OutputState }

func (UnauthenticatedClientActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnauthenticatedClientAction)(nil)).Elem()
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput {
	return o
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionOutputWithContext(ctx context.Context) UnauthenticatedClientActionOutput {
	return o
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return o.ToUnauthenticatedClientActionPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnauthenticatedClientAction) *UnauthenticatedClientAction {
		return &v
	}).(UnauthenticatedClientActionPtrOutput)
}

func (o UnauthenticatedClientActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnauthenticatedClientAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UnauthenticatedClientActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnauthenticatedClientAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UnauthenticatedClientActionPtrOutput struct{ *pulumi.OutputState }

func (UnauthenticatedClientActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnauthenticatedClientAction)(nil)).Elem()
}

func (o UnauthenticatedClientActionPtrOutput) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return o
}

func (o UnauthenticatedClientActionPtrOutput) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return o
}

func (o UnauthenticatedClientActionPtrOutput) Elem() UnauthenticatedClientActionOutput {
	return o.ApplyT(func(v *UnauthenticatedClientAction) UnauthenticatedClientAction {
		if v != nil {
			return *v
		}
		var ret UnauthenticatedClientAction
		return ret
	}).(UnauthenticatedClientActionOutput)
}

func (o UnauthenticatedClientActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UnauthenticatedClientAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UnauthenticatedClientActionInput interface {
	pulumi.Input

	ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput
	ToUnauthenticatedClientActionOutputWithContext(context.Context) UnauthenticatedClientActionOutput
}

var unauthenticatedClientActionPtrType = reflect.TypeOf((**UnauthenticatedClientAction)(nil)).Elem()

type UnauthenticatedClientActionPtrInput interface {
	pulumi.Input

	ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput
	ToUnauthenticatedClientActionPtrOutputWithContext(context.Context) UnauthenticatedClientActionPtrOutput
}

type unauthenticatedClientActionPtr string

func UnauthenticatedClientActionPtr(v string) UnauthenticatedClientActionPtrInput {
	return (*unauthenticatedClientActionPtr)(&v)
}

func (*unauthenticatedClientActionPtr) ElementType() reflect.Type {
	return unauthenticatedClientActionPtrType
}

func (in *unauthenticatedClientActionPtr) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return pulumi.ToOutput(in).(UnauthenticatedClientActionPtrOutput)
}

func (in *unauthenticatedClientActionPtr) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UnauthenticatedClientActionPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessControlEntryActionInput)(nil)).Elem(), AccessControlEntryAction("Permit"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessControlEntryActionPtrInput)(nil)).Elem(), AccessControlEntryAction("Permit"))
	pulumi.RegisterInputType(reflect.TypeOf((*AutoHealActionTypeInput)(nil)).Elem(), AutoHealActionType("Recycle"))
	pulumi.RegisterInputType(reflect.TypeOf((*AutoHealActionTypePtrInput)(nil)).Elem(), AutoHealActionType("Recycle"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureResourceTypeInput)(nil)).Elem(), AzureResourceType("Website"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureResourceTypePtrInput)(nil)).Elem(), AzureResourceType("Website"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureStorageTypeInput)(nil)).Elem(), AzureStorageType("AzureFiles"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureStorageTypePtrInput)(nil)).Elem(), AzureStorageType("AzureFiles"))
	pulumi.RegisterInputType(reflect.TypeOf((*BuiltInAuthenticationProviderInput)(nil)).Elem(), BuiltInAuthenticationProvider("AzureActiveDirectory"))
	pulumi.RegisterInputType(reflect.TypeOf((*BuiltInAuthenticationProviderPtrInput)(nil)).Elem(), BuiltInAuthenticationProvider("AzureActiveDirectory"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeModeOptionsInput)(nil)).Elem(), ComputeModeOptions("Shared"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeModeOptionsPtrInput)(nil)).Elem(), ComputeModeOptions("Shared"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStringTypeInput)(nil)).Elem(), ConnectionStringType("MySql"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStringTypePtrInput)(nil)).Elem(), ConnectionStringType("MySql"))
	pulumi.RegisterInputType(reflect.TypeOf((*CustomHostNameDnsRecordTypeInput)(nil)).Elem(), CustomHostNameDnsRecordType("CName"))
	pulumi.RegisterInputType(reflect.TypeOf((*CustomHostNameDnsRecordTypePtrInput)(nil)).Elem(), CustomHostNameDnsRecordType("CName"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseTypeInput)(nil)).Elem(), DatabaseType("SqlAzure"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseTypePtrInput)(nil)).Elem(), DatabaseType("SqlAzure"))
	pulumi.RegisterInputType(reflect.TypeOf((*FrequencyUnitInput)(nil)).Elem(), FrequencyUnit("Day"))
	pulumi.RegisterInputType(reflect.TypeOf((*FrequencyUnitPtrInput)(nil)).Elem(), FrequencyUnit("Day"))
	pulumi.RegisterInputType(reflect.TypeOf((*FtpsStateInput)(nil)).Elem(), FtpsState("AllAllowed"))
	pulumi.RegisterInputType(reflect.TypeOf((*FtpsStatePtrInput)(nil)).Elem(), FtpsState("AllAllowed"))
	pulumi.RegisterInputType(reflect.TypeOf((*HostNameTypeInput)(nil)).Elem(), HostNameType("Verified"))
	pulumi.RegisterInputType(reflect.TypeOf((*HostNameTypePtrInput)(nil)).Elem(), HostNameType("Verified"))
	pulumi.RegisterInputType(reflect.TypeOf((*HostTypeInput)(nil)).Elem(), HostType("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*HostTypePtrInput)(nil)).Elem(), HostType("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*InternalLoadBalancingModeInput)(nil)).Elem(), InternalLoadBalancingMode("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*InternalLoadBalancingModePtrInput)(nil)).Elem(), InternalLoadBalancingMode("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterTagInput)(nil)).Elem(), IpFilterTag("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterTagPtrInput)(nil)).Elem(), IpFilterTag("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*LogLevelInput)(nil)).Elem(), LogLevel("Off"))
	pulumi.RegisterInputType(reflect.TypeOf((*LogLevelPtrInput)(nil)).Elem(), LogLevel("Off"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPipelineModeInput)(nil)).Elem(), ManagedPipelineMode("Integrated"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPipelineModePtrInput)(nil)).Elem(), ManagedPipelineMode("Integrated"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedServiceIdentityTypeInput)(nil)).Elem(), ManagedServiceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedServiceIdentityTypePtrInput)(nil)).Elem(), ManagedServiceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicCertificateLocationInput)(nil)).Elem(), PublicCertificateLocation("CurrentUserMy"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicCertificateLocationPtrInput)(nil)).Elem(), PublicCertificateLocation("CurrentUserMy"))
	pulumi.RegisterInputType(reflect.TypeOf((*RedundancyModeInput)(nil)).Elem(), RedundancyMode("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*RedundancyModePtrInput)(nil)).Elem(), RedundancyMode("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTypeInput)(nil)).Elem(), RouteType("DEFAULT"))
	pulumi.RegisterInputType(reflect.TypeOf((*RouteTypePtrInput)(nil)).Elem(), RouteType("DEFAULT"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScmTypeInput)(nil)).Elem(), ScmType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScmTypePtrInput)(nil)).Elem(), ScmType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*SiteLoadBalancingInput)(nil)).Elem(), SiteLoadBalancing("WeightedRoundRobin"))
	pulumi.RegisterInputType(reflect.TypeOf((*SiteLoadBalancingPtrInput)(nil)).Elem(), SiteLoadBalancing("WeightedRoundRobin"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslStateInput)(nil)).Elem(), SslState("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslStatePtrInput)(nil)).Elem(), SslState("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*SupportedTlsVersionsInput)(nil)).Elem(), SupportedTlsVersions("1.0"))
	pulumi.RegisterInputType(reflect.TypeOf((*SupportedTlsVersionsPtrInput)(nil)).Elem(), SupportedTlsVersions("1.0"))
	pulumi.RegisterInputType(reflect.TypeOf((*UnauthenticatedClientActionInput)(nil)).Elem(), UnauthenticatedClientAction("RedirectToLoginPage"))
	pulumi.RegisterInputType(reflect.TypeOf((*UnauthenticatedClientActionPtrInput)(nil)).Elem(), UnauthenticatedClientAction("RedirectToLoginPage"))
	pulumi.RegisterOutputType(AccessControlEntryActionOutput{})
	pulumi.RegisterOutputType(AccessControlEntryActionPtrOutput{})
	pulumi.RegisterOutputType(AutoHealActionTypeOutput{})
	pulumi.RegisterOutputType(AutoHealActionTypePtrOutput{})
	pulumi.RegisterOutputType(AzureResourceTypeOutput{})
	pulumi.RegisterOutputType(AzureResourceTypePtrOutput{})
	pulumi.RegisterOutputType(AzureStorageTypeOutput{})
	pulumi.RegisterOutputType(AzureStorageTypePtrOutput{})
	pulumi.RegisterOutputType(BuiltInAuthenticationProviderOutput{})
	pulumi.RegisterOutputType(BuiltInAuthenticationProviderPtrOutput{})
	pulumi.RegisterOutputType(ComputeModeOptionsOutput{})
	pulumi.RegisterOutputType(ComputeModeOptionsPtrOutput{})
	pulumi.RegisterOutputType(ConnectionStringTypeOutput{})
	pulumi.RegisterOutputType(ConnectionStringTypePtrOutput{})
	pulumi.RegisterOutputType(CustomHostNameDnsRecordTypeOutput{})
	pulumi.RegisterOutputType(CustomHostNameDnsRecordTypePtrOutput{})
	pulumi.RegisterOutputType(DatabaseTypeOutput{})
	pulumi.RegisterOutputType(DatabaseTypePtrOutput{})
	pulumi.RegisterOutputType(FrequencyUnitOutput{})
	pulumi.RegisterOutputType(FrequencyUnitPtrOutput{})
	pulumi.RegisterOutputType(FtpsStateOutput{})
	pulumi.RegisterOutputType(FtpsStatePtrOutput{})
	pulumi.RegisterOutputType(HostNameTypeOutput{})
	pulumi.RegisterOutputType(HostNameTypePtrOutput{})
	pulumi.RegisterOutputType(HostTypeOutput{})
	pulumi.RegisterOutputType(HostTypePtrOutput{})
	pulumi.RegisterOutputType(InternalLoadBalancingModeOutput{})
	pulumi.RegisterOutputType(InternalLoadBalancingModePtrOutput{})
	pulumi.RegisterOutputType(IpFilterTagOutput{})
	pulumi.RegisterOutputType(IpFilterTagPtrOutput{})
	pulumi.RegisterOutputType(LogLevelOutput{})
	pulumi.RegisterOutputType(LogLevelPtrOutput{})
	pulumi.RegisterOutputType(ManagedPipelineModeOutput{})
	pulumi.RegisterOutputType(ManagedPipelineModePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(PublicCertificateLocationOutput{})
	pulumi.RegisterOutputType(PublicCertificateLocationPtrOutput{})
	pulumi.RegisterOutputType(RedundancyModeOutput{})
	pulumi.RegisterOutputType(RedundancyModePtrOutput{})
	pulumi.RegisterOutputType(RouteTypeOutput{})
	pulumi.RegisterOutputType(RouteTypePtrOutput{})
	pulumi.RegisterOutputType(ScmTypeOutput{})
	pulumi.RegisterOutputType(ScmTypePtrOutput{})
	pulumi.RegisterOutputType(SiteLoadBalancingOutput{})
	pulumi.RegisterOutputType(SiteLoadBalancingPtrOutput{})
	pulumi.RegisterOutputType(SslStateOutput{})
	pulumi.RegisterOutputType(SslStatePtrOutput{})
	pulumi.RegisterOutputType(SupportedTlsVersionsOutput{})
	pulumi.RegisterOutputType(SupportedTlsVersionsPtrOutput{})
	pulumi.RegisterOutputType(UnauthenticatedClientActionOutput{})
	pulumi.RegisterOutputType(UnauthenticatedClientActionPtrOutput{})
}
