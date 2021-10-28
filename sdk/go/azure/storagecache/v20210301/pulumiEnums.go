


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CacheIdentityType string

const (
	CacheIdentityTypeSystemAssigned = CacheIdentityType("SystemAssigned")
	CacheIdentityTypeNone           = CacheIdentityType("None")
)

func (CacheIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheIdentityType)(nil)).Elem()
}

func (e CacheIdentityType) ToCacheIdentityTypeOutput() CacheIdentityTypeOutput {
	return pulumi.ToOutput(e).(CacheIdentityTypeOutput)
}

func (e CacheIdentityType) ToCacheIdentityTypeOutputWithContext(ctx context.Context) CacheIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CacheIdentityTypeOutput)
}

func (e CacheIdentityType) ToCacheIdentityTypePtrOutput() CacheIdentityTypePtrOutput {
	return e.ToCacheIdentityTypePtrOutputWithContext(context.Background())
}

func (e CacheIdentityType) ToCacheIdentityTypePtrOutputWithContext(ctx context.Context) CacheIdentityTypePtrOutput {
	return CacheIdentityType(e).ToCacheIdentityTypeOutputWithContext(ctx).ToCacheIdentityTypePtrOutputWithContext(ctx)
}

func (e CacheIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CacheIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CacheIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CacheIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CacheIdentityTypeOutput struct{ *pulumi.OutputState }

func (CacheIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheIdentityType)(nil)).Elem()
}

func (o CacheIdentityTypeOutput) ToCacheIdentityTypeOutput() CacheIdentityTypeOutput {
	return o
}

func (o CacheIdentityTypeOutput) ToCacheIdentityTypeOutputWithContext(ctx context.Context) CacheIdentityTypeOutput {
	return o
}

func (o CacheIdentityTypeOutput) ToCacheIdentityTypePtrOutput() CacheIdentityTypePtrOutput {
	return o.ToCacheIdentityTypePtrOutputWithContext(context.Background())
}

func (o CacheIdentityTypeOutput) ToCacheIdentityTypePtrOutputWithContext(ctx context.Context) CacheIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CacheIdentityType) *CacheIdentityType {
		return &v
	}).(CacheIdentityTypePtrOutput)
}

func (o CacheIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CacheIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CacheIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CacheIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CacheIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CacheIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CacheIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (CacheIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheIdentityType)(nil)).Elem()
}

func (o CacheIdentityTypePtrOutput) ToCacheIdentityTypePtrOutput() CacheIdentityTypePtrOutput {
	return o
}

func (o CacheIdentityTypePtrOutput) ToCacheIdentityTypePtrOutputWithContext(ctx context.Context) CacheIdentityTypePtrOutput {
	return o
}

func (o CacheIdentityTypePtrOutput) Elem() CacheIdentityTypeOutput {
	return o.ApplyT(func(v *CacheIdentityType) CacheIdentityType {
		if v != nil {
			return *v
		}
		var ret CacheIdentityType
		return ret
	}).(CacheIdentityTypeOutput)
}

func (o CacheIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CacheIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CacheIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CacheIdentityTypeInput interface {
	pulumi.Input

	ToCacheIdentityTypeOutput() CacheIdentityTypeOutput
	ToCacheIdentityTypeOutputWithContext(context.Context) CacheIdentityTypeOutput
}

var cacheIdentityTypePtrType = reflect.TypeOf((**CacheIdentityType)(nil)).Elem()

type CacheIdentityTypePtrInput interface {
	pulumi.Input

	ToCacheIdentityTypePtrOutput() CacheIdentityTypePtrOutput
	ToCacheIdentityTypePtrOutputWithContext(context.Context) CacheIdentityTypePtrOutput
}

type cacheIdentityTypePtr string

func CacheIdentityTypePtr(v string) CacheIdentityTypePtrInput {
	return (*cacheIdentityTypePtr)(&v)
}

func (*cacheIdentityTypePtr) ElementType() reflect.Type {
	return cacheIdentityTypePtrType
}

func (in *cacheIdentityTypePtr) ToCacheIdentityTypePtrOutput() CacheIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(CacheIdentityTypePtrOutput)
}

func (in *cacheIdentityTypePtr) ToCacheIdentityTypePtrOutputWithContext(ctx context.Context) CacheIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CacheIdentityTypePtrOutput)
}

type NfsAccessRuleAccess string

const (
	NfsAccessRuleAccessNo = NfsAccessRuleAccess("no")
	NfsAccessRuleAccessRo = NfsAccessRuleAccess("ro")
	NfsAccessRuleAccessRw = NfsAccessRuleAccess("rw")
)

func (NfsAccessRuleAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessRuleAccess)(nil)).Elem()
}

func (e NfsAccessRuleAccess) ToNfsAccessRuleAccessOutput() NfsAccessRuleAccessOutput {
	return pulumi.ToOutput(e).(NfsAccessRuleAccessOutput)
}

func (e NfsAccessRuleAccess) ToNfsAccessRuleAccessOutputWithContext(ctx context.Context) NfsAccessRuleAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NfsAccessRuleAccessOutput)
}

func (e NfsAccessRuleAccess) ToNfsAccessRuleAccessPtrOutput() NfsAccessRuleAccessPtrOutput {
	return e.ToNfsAccessRuleAccessPtrOutputWithContext(context.Background())
}

func (e NfsAccessRuleAccess) ToNfsAccessRuleAccessPtrOutputWithContext(ctx context.Context) NfsAccessRuleAccessPtrOutput {
	return NfsAccessRuleAccess(e).ToNfsAccessRuleAccessOutputWithContext(ctx).ToNfsAccessRuleAccessPtrOutputWithContext(ctx)
}

func (e NfsAccessRuleAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NfsAccessRuleAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NfsAccessRuleAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NfsAccessRuleAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NfsAccessRuleAccessOutput struct{ *pulumi.OutputState }

func (NfsAccessRuleAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessRuleAccess)(nil)).Elem()
}

func (o NfsAccessRuleAccessOutput) ToNfsAccessRuleAccessOutput() NfsAccessRuleAccessOutput {
	return o
}

func (o NfsAccessRuleAccessOutput) ToNfsAccessRuleAccessOutputWithContext(ctx context.Context) NfsAccessRuleAccessOutput {
	return o
}

func (o NfsAccessRuleAccessOutput) ToNfsAccessRuleAccessPtrOutput() NfsAccessRuleAccessPtrOutput {
	return o.ToNfsAccessRuleAccessPtrOutputWithContext(context.Background())
}

func (o NfsAccessRuleAccessOutput) ToNfsAccessRuleAccessPtrOutputWithContext(ctx context.Context) NfsAccessRuleAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NfsAccessRuleAccess) *NfsAccessRuleAccess {
		return &v
	}).(NfsAccessRuleAccessPtrOutput)
}

func (o NfsAccessRuleAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NfsAccessRuleAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NfsAccessRuleAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NfsAccessRuleAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NfsAccessRuleAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NfsAccessRuleAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NfsAccessRuleAccessPtrOutput struct{ *pulumi.OutputState }

func (NfsAccessRuleAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsAccessRuleAccess)(nil)).Elem()
}

func (o NfsAccessRuleAccessPtrOutput) ToNfsAccessRuleAccessPtrOutput() NfsAccessRuleAccessPtrOutput {
	return o
}

func (o NfsAccessRuleAccessPtrOutput) ToNfsAccessRuleAccessPtrOutputWithContext(ctx context.Context) NfsAccessRuleAccessPtrOutput {
	return o
}

func (o NfsAccessRuleAccessPtrOutput) Elem() NfsAccessRuleAccessOutput {
	return o.ApplyT(func(v *NfsAccessRuleAccess) NfsAccessRuleAccess {
		if v != nil {
			return *v
		}
		var ret NfsAccessRuleAccess
		return ret
	}).(NfsAccessRuleAccessOutput)
}

func (o NfsAccessRuleAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NfsAccessRuleAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NfsAccessRuleAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NfsAccessRuleAccessInput interface {
	pulumi.Input

	ToNfsAccessRuleAccessOutput() NfsAccessRuleAccessOutput
	ToNfsAccessRuleAccessOutputWithContext(context.Context) NfsAccessRuleAccessOutput
}

var nfsAccessRuleAccessPtrType = reflect.TypeOf((**NfsAccessRuleAccess)(nil)).Elem()

type NfsAccessRuleAccessPtrInput interface {
	pulumi.Input

	ToNfsAccessRuleAccessPtrOutput() NfsAccessRuleAccessPtrOutput
	ToNfsAccessRuleAccessPtrOutputWithContext(context.Context) NfsAccessRuleAccessPtrOutput
}

type nfsAccessRuleAccessPtr string

func NfsAccessRuleAccessPtr(v string) NfsAccessRuleAccessPtrInput {
	return (*nfsAccessRuleAccessPtr)(&v)
}

func (*nfsAccessRuleAccessPtr) ElementType() reflect.Type {
	return nfsAccessRuleAccessPtrType
}

func (in *nfsAccessRuleAccessPtr) ToNfsAccessRuleAccessPtrOutput() NfsAccessRuleAccessPtrOutput {
	return pulumi.ToOutput(in).(NfsAccessRuleAccessPtrOutput)
}

func (in *nfsAccessRuleAccessPtr) ToNfsAccessRuleAccessPtrOutputWithContext(ctx context.Context) NfsAccessRuleAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NfsAccessRuleAccessPtrOutput)
}

type NfsAccessRuleScope string

const (
	NfsAccessRuleScopeDefault = NfsAccessRuleScope("default")
	NfsAccessRuleScopeNetwork = NfsAccessRuleScope("network")
	NfsAccessRuleScopeHost    = NfsAccessRuleScope("host")
)

func (NfsAccessRuleScope) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessRuleScope)(nil)).Elem()
}

func (e NfsAccessRuleScope) ToNfsAccessRuleScopeOutput() NfsAccessRuleScopeOutput {
	return pulumi.ToOutput(e).(NfsAccessRuleScopeOutput)
}

func (e NfsAccessRuleScope) ToNfsAccessRuleScopeOutputWithContext(ctx context.Context) NfsAccessRuleScopeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NfsAccessRuleScopeOutput)
}

func (e NfsAccessRuleScope) ToNfsAccessRuleScopePtrOutput() NfsAccessRuleScopePtrOutput {
	return e.ToNfsAccessRuleScopePtrOutputWithContext(context.Background())
}

func (e NfsAccessRuleScope) ToNfsAccessRuleScopePtrOutputWithContext(ctx context.Context) NfsAccessRuleScopePtrOutput {
	return NfsAccessRuleScope(e).ToNfsAccessRuleScopeOutputWithContext(ctx).ToNfsAccessRuleScopePtrOutputWithContext(ctx)
}

func (e NfsAccessRuleScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NfsAccessRuleScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NfsAccessRuleScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NfsAccessRuleScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NfsAccessRuleScopeOutput struct{ *pulumi.OutputState }

func (NfsAccessRuleScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsAccessRuleScope)(nil)).Elem()
}

func (o NfsAccessRuleScopeOutput) ToNfsAccessRuleScopeOutput() NfsAccessRuleScopeOutput {
	return o
}

func (o NfsAccessRuleScopeOutput) ToNfsAccessRuleScopeOutputWithContext(ctx context.Context) NfsAccessRuleScopeOutput {
	return o
}

func (o NfsAccessRuleScopeOutput) ToNfsAccessRuleScopePtrOutput() NfsAccessRuleScopePtrOutput {
	return o.ToNfsAccessRuleScopePtrOutputWithContext(context.Background())
}

func (o NfsAccessRuleScopeOutput) ToNfsAccessRuleScopePtrOutputWithContext(ctx context.Context) NfsAccessRuleScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NfsAccessRuleScope) *NfsAccessRuleScope {
		return &v
	}).(NfsAccessRuleScopePtrOutput)
}

func (o NfsAccessRuleScopeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NfsAccessRuleScopeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NfsAccessRuleScope) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NfsAccessRuleScopeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NfsAccessRuleScopeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NfsAccessRuleScope) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NfsAccessRuleScopePtrOutput struct{ *pulumi.OutputState }

func (NfsAccessRuleScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsAccessRuleScope)(nil)).Elem()
}

func (o NfsAccessRuleScopePtrOutput) ToNfsAccessRuleScopePtrOutput() NfsAccessRuleScopePtrOutput {
	return o
}

func (o NfsAccessRuleScopePtrOutput) ToNfsAccessRuleScopePtrOutputWithContext(ctx context.Context) NfsAccessRuleScopePtrOutput {
	return o
}

func (o NfsAccessRuleScopePtrOutput) Elem() NfsAccessRuleScopeOutput {
	return o.ApplyT(func(v *NfsAccessRuleScope) NfsAccessRuleScope {
		if v != nil {
			return *v
		}
		var ret NfsAccessRuleScope
		return ret
	}).(NfsAccessRuleScopeOutput)
}

func (o NfsAccessRuleScopePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NfsAccessRuleScopePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NfsAccessRuleScope) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NfsAccessRuleScopeInput interface {
	pulumi.Input

	ToNfsAccessRuleScopeOutput() NfsAccessRuleScopeOutput
	ToNfsAccessRuleScopeOutputWithContext(context.Context) NfsAccessRuleScopeOutput
}

var nfsAccessRuleScopePtrType = reflect.TypeOf((**NfsAccessRuleScope)(nil)).Elem()

type NfsAccessRuleScopePtrInput interface {
	pulumi.Input

	ToNfsAccessRuleScopePtrOutput() NfsAccessRuleScopePtrOutput
	ToNfsAccessRuleScopePtrOutputWithContext(context.Context) NfsAccessRuleScopePtrOutput
}

type nfsAccessRuleScopePtr string

func NfsAccessRuleScopePtr(v string) NfsAccessRuleScopePtrInput {
	return (*nfsAccessRuleScopePtr)(&v)
}

func (*nfsAccessRuleScopePtr) ElementType() reflect.Type {
	return nfsAccessRuleScopePtrType
}

func (in *nfsAccessRuleScopePtr) ToNfsAccessRuleScopePtrOutput() NfsAccessRuleScopePtrOutput {
	return pulumi.ToOutput(in).(NfsAccessRuleScopePtrOutput)
}

func (in *nfsAccessRuleScopePtr) ToNfsAccessRuleScopePtrOutputWithContext(ctx context.Context) NfsAccessRuleScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NfsAccessRuleScopePtrOutput)
}

type ProvisioningStateType string

const (
	ProvisioningStateTypeSucceeded = ProvisioningStateType("Succeeded")
	ProvisioningStateTypeFailed    = ProvisioningStateType("Failed")
	ProvisioningStateTypeCancelled = ProvisioningStateType("Cancelled")
	ProvisioningStateTypeCreating  = ProvisioningStateType("Creating")
	ProvisioningStateTypeDeleting  = ProvisioningStateType("Deleting")
	ProvisioningStateTypeUpdating  = ProvisioningStateType("Updating")
)

func (ProvisioningStateType) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningStateType)(nil)).Elem()
}

func (e ProvisioningStateType) ToProvisioningStateTypeOutput() ProvisioningStateTypeOutput {
	return pulumi.ToOutput(e).(ProvisioningStateTypeOutput)
}

func (e ProvisioningStateType) ToProvisioningStateTypeOutputWithContext(ctx context.Context) ProvisioningStateTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateTypeOutput)
}

func (e ProvisioningStateType) ToProvisioningStateTypePtrOutput() ProvisioningStateTypePtrOutput {
	return e.ToProvisioningStateTypePtrOutputWithContext(context.Background())
}

func (e ProvisioningStateType) ToProvisioningStateTypePtrOutputWithContext(ctx context.Context) ProvisioningStateTypePtrOutput {
	return ProvisioningStateType(e).ToProvisioningStateTypeOutputWithContext(ctx).ToProvisioningStateTypePtrOutputWithContext(ctx)
}

func (e ProvisioningStateType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningStateType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningStateType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningStateType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateTypeOutput struct{ *pulumi.OutputState }

func (ProvisioningStateTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningStateType)(nil)).Elem()
}

func (o ProvisioningStateTypeOutput) ToProvisioningStateTypeOutput() ProvisioningStateTypeOutput {
	return o
}

func (o ProvisioningStateTypeOutput) ToProvisioningStateTypeOutputWithContext(ctx context.Context) ProvisioningStateTypeOutput {
	return o
}

func (o ProvisioningStateTypeOutput) ToProvisioningStateTypePtrOutput() ProvisioningStateTypePtrOutput {
	return o.ToProvisioningStateTypePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateTypeOutput) ToProvisioningStateTypePtrOutputWithContext(ctx context.Context) ProvisioningStateTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningStateType) *ProvisioningStateType {
		return &v
	}).(ProvisioningStateTypePtrOutput)
}

func (o ProvisioningStateTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningStateType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningStateType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStateTypePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStateTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningStateType)(nil)).Elem()
}

func (o ProvisioningStateTypePtrOutput) ToProvisioningStateTypePtrOutput() ProvisioningStateTypePtrOutput {
	return o
}

func (o ProvisioningStateTypePtrOutput) ToProvisioningStateTypePtrOutputWithContext(ctx context.Context) ProvisioningStateTypePtrOutput {
	return o
}

func (o ProvisioningStateTypePtrOutput) Elem() ProvisioningStateTypeOutput {
	return o.ApplyT(func(v *ProvisioningStateType) ProvisioningStateType {
		if v != nil {
			return *v
		}
		var ret ProvisioningStateType
		return ret
	}).(ProvisioningStateTypeOutput)
}

func (o ProvisioningStateTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningStateType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProvisioningStateTypeInput interface {
	pulumi.Input

	ToProvisioningStateTypeOutput() ProvisioningStateTypeOutput
	ToProvisioningStateTypeOutputWithContext(context.Context) ProvisioningStateTypeOutput
}

var provisioningStateTypePtrType = reflect.TypeOf((**ProvisioningStateType)(nil)).Elem()

type ProvisioningStateTypePtrInput interface {
	pulumi.Input

	ToProvisioningStateTypePtrOutput() ProvisioningStateTypePtrOutput
	ToProvisioningStateTypePtrOutputWithContext(context.Context) ProvisioningStateTypePtrOutput
}

type provisioningStateTypePtr string

func ProvisioningStateTypePtr(v string) ProvisioningStateTypePtrInput {
	return (*provisioningStateTypePtr)(&v)
}

func (*provisioningStateTypePtr) ElementType() reflect.Type {
	return provisioningStateTypePtrType
}

func (in *provisioningStateTypePtr) ToProvisioningStateTypePtrOutput() ProvisioningStateTypePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStateTypePtrOutput)
}

func (in *provisioningStateTypePtr) ToProvisioningStateTypePtrOutputWithContext(ctx context.Context) ProvisioningStateTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStateTypePtrOutput)
}

type StorageTargetType string

const (
	StorageTargetTypeNfs3    = StorageTargetType("nfs3")
	StorageTargetTypeClfs    = StorageTargetType("clfs")
	StorageTargetTypeUnknown = StorageTargetType("unknown")
	StorageTargetTypeBlobNfs = StorageTargetType("blobNfs")
)

func (StorageTargetType) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageTargetType)(nil)).Elem()
}

func (e StorageTargetType) ToStorageTargetTypeOutput() StorageTargetTypeOutput {
	return pulumi.ToOutput(e).(StorageTargetTypeOutput)
}

func (e StorageTargetType) ToStorageTargetTypeOutputWithContext(ctx context.Context) StorageTargetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageTargetTypeOutput)
}

func (e StorageTargetType) ToStorageTargetTypePtrOutput() StorageTargetTypePtrOutput {
	return e.ToStorageTargetTypePtrOutputWithContext(context.Background())
}

func (e StorageTargetType) ToStorageTargetTypePtrOutputWithContext(ctx context.Context) StorageTargetTypePtrOutput {
	return StorageTargetType(e).ToStorageTargetTypeOutputWithContext(ctx).ToStorageTargetTypePtrOutputWithContext(ctx)
}

func (e StorageTargetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageTargetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageTargetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageTargetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageTargetTypeOutput struct{ *pulumi.OutputState }

func (StorageTargetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageTargetType)(nil)).Elem()
}

func (o StorageTargetTypeOutput) ToStorageTargetTypeOutput() StorageTargetTypeOutput {
	return o
}

func (o StorageTargetTypeOutput) ToStorageTargetTypeOutputWithContext(ctx context.Context) StorageTargetTypeOutput {
	return o
}

func (o StorageTargetTypeOutput) ToStorageTargetTypePtrOutput() StorageTargetTypePtrOutput {
	return o.ToStorageTargetTypePtrOutputWithContext(context.Background())
}

func (o StorageTargetTypeOutput) ToStorageTargetTypePtrOutputWithContext(ctx context.Context) StorageTargetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageTargetType) *StorageTargetType {
		return &v
	}).(StorageTargetTypePtrOutput)
}

func (o StorageTargetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageTargetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageTargetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageTargetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageTargetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageTargetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageTargetTypePtrOutput struct{ *pulumi.OutputState }

func (StorageTargetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageTargetType)(nil)).Elem()
}

func (o StorageTargetTypePtrOutput) ToStorageTargetTypePtrOutput() StorageTargetTypePtrOutput {
	return o
}

func (o StorageTargetTypePtrOutput) ToStorageTargetTypePtrOutputWithContext(ctx context.Context) StorageTargetTypePtrOutput {
	return o
}

func (o StorageTargetTypePtrOutput) Elem() StorageTargetTypeOutput {
	return o.ApplyT(func(v *StorageTargetType) StorageTargetType {
		if v != nil {
			return *v
		}
		var ret StorageTargetType
		return ret
	}).(StorageTargetTypeOutput)
}

func (o StorageTargetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageTargetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageTargetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageTargetTypeInput interface {
	pulumi.Input

	ToStorageTargetTypeOutput() StorageTargetTypeOutput
	ToStorageTargetTypeOutputWithContext(context.Context) StorageTargetTypeOutput
}

var storageTargetTypePtrType = reflect.TypeOf((**StorageTargetType)(nil)).Elem()

type StorageTargetTypePtrInput interface {
	pulumi.Input

	ToStorageTargetTypePtrOutput() StorageTargetTypePtrOutput
	ToStorageTargetTypePtrOutputWithContext(context.Context) StorageTargetTypePtrOutput
}

type storageTargetTypePtr string

func StorageTargetTypePtr(v string) StorageTargetTypePtrInput {
	return (*storageTargetTypePtr)(&v)
}

func (*storageTargetTypePtr) ElementType() reflect.Type {
	return storageTargetTypePtrType
}

func (in *storageTargetTypePtr) ToStorageTargetTypePtrOutput() StorageTargetTypePtrOutput {
	return pulumi.ToOutput(in).(StorageTargetTypePtrOutput)
}

func (in *storageTargetTypePtr) ToStorageTargetTypePtrOutputWithContext(ctx context.Context) StorageTargetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageTargetTypePtrOutput)
}

type UsernameSource string

const (
	UsernameSourceAD   = UsernameSource("AD")
	UsernameSourceLDAP = UsernameSource("LDAP")
	UsernameSourceFile = UsernameSource("File")
	UsernameSourceNone = UsernameSource("None")
)

func (UsernameSource) ElementType() reflect.Type {
	return reflect.TypeOf((*UsernameSource)(nil)).Elem()
}

func (e UsernameSource) ToUsernameSourceOutput() UsernameSourceOutput {
	return pulumi.ToOutput(e).(UsernameSourceOutput)
}

func (e UsernameSource) ToUsernameSourceOutputWithContext(ctx context.Context) UsernameSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UsernameSourceOutput)
}

func (e UsernameSource) ToUsernameSourcePtrOutput() UsernameSourcePtrOutput {
	return e.ToUsernameSourcePtrOutputWithContext(context.Background())
}

func (e UsernameSource) ToUsernameSourcePtrOutputWithContext(ctx context.Context) UsernameSourcePtrOutput {
	return UsernameSource(e).ToUsernameSourceOutputWithContext(ctx).ToUsernameSourcePtrOutputWithContext(ctx)
}

func (e UsernameSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UsernameSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UsernameSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UsernameSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UsernameSourceOutput struct{ *pulumi.OutputState }

func (UsernameSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsernameSource)(nil)).Elem()
}

func (o UsernameSourceOutput) ToUsernameSourceOutput() UsernameSourceOutput {
	return o
}

func (o UsernameSourceOutput) ToUsernameSourceOutputWithContext(ctx context.Context) UsernameSourceOutput {
	return o
}

func (o UsernameSourceOutput) ToUsernameSourcePtrOutput() UsernameSourcePtrOutput {
	return o.ToUsernameSourcePtrOutputWithContext(context.Background())
}

func (o UsernameSourceOutput) ToUsernameSourcePtrOutputWithContext(ctx context.Context) UsernameSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UsernameSource) *UsernameSource {
		return &v
	}).(UsernameSourcePtrOutput)
}

func (o UsernameSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UsernameSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UsernameSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UsernameSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UsernameSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UsernameSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UsernameSourcePtrOutput struct{ *pulumi.OutputState }

func (UsernameSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsernameSource)(nil)).Elem()
}

func (o UsernameSourcePtrOutput) ToUsernameSourcePtrOutput() UsernameSourcePtrOutput {
	return o
}

func (o UsernameSourcePtrOutput) ToUsernameSourcePtrOutputWithContext(ctx context.Context) UsernameSourcePtrOutput {
	return o
}

func (o UsernameSourcePtrOutput) Elem() UsernameSourceOutput {
	return o.ApplyT(func(v *UsernameSource) UsernameSource {
		if v != nil {
			return *v
		}
		var ret UsernameSource
		return ret
	}).(UsernameSourceOutput)
}

func (o UsernameSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UsernameSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UsernameSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UsernameSourceInput interface {
	pulumi.Input

	ToUsernameSourceOutput() UsernameSourceOutput
	ToUsernameSourceOutputWithContext(context.Context) UsernameSourceOutput
}

var usernameSourcePtrType = reflect.TypeOf((**UsernameSource)(nil)).Elem()

type UsernameSourcePtrInput interface {
	pulumi.Input

	ToUsernameSourcePtrOutput() UsernameSourcePtrOutput
	ToUsernameSourcePtrOutputWithContext(context.Context) UsernameSourcePtrOutput
}

type usernameSourcePtr string

func UsernameSourcePtr(v string) UsernameSourcePtrInput {
	return (*usernameSourcePtr)(&v)
}

func (*usernameSourcePtr) ElementType() reflect.Type {
	return usernameSourcePtrType
}

func (in *usernameSourcePtr) ToUsernameSourcePtrOutput() UsernameSourcePtrOutput {
	return pulumi.ToOutput(in).(UsernameSourcePtrOutput)
}

func (in *usernameSourcePtr) ToUsernameSourcePtrOutputWithContext(ctx context.Context) UsernameSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UsernameSourcePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheIdentityTypeOutput{})
	pulumi.RegisterOutputType(CacheIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(NfsAccessRuleAccessOutput{})
	pulumi.RegisterOutputType(NfsAccessRuleAccessPtrOutput{})
	pulumi.RegisterOutputType(NfsAccessRuleScopeOutput{})
	pulumi.RegisterOutputType(NfsAccessRuleScopePtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateTypeOutput{})
	pulumi.RegisterOutputType(ProvisioningStateTypePtrOutput{})
	pulumi.RegisterOutputType(StorageTargetTypeOutput{})
	pulumi.RegisterOutputType(StorageTargetTypePtrOutput{})
	pulumi.RegisterOutputType(UsernameSourceOutput{})
	pulumi.RegisterOutputType(UsernameSourcePtrOutput{})
}
