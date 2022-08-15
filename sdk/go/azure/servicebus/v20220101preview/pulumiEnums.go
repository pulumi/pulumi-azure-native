


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessRights string

const (
	AccessRightsManage = AccessRights("Manage")
	AccessRightsSend   = AccessRights("Send")
	AccessRightsListen = AccessRights("Listen")
)

type DefaultAction string

const (
	DefaultActionAllow = DefaultAction("Allow")
	DefaultActionDeny  = DefaultAction("Deny")
)

type EndPointProvisioningState string

const (
	EndPointProvisioningStateCreating  = EndPointProvisioningState("Creating")
	EndPointProvisioningStateUpdating  = EndPointProvisioningState("Updating")
	EndPointProvisioningStateDeleting  = EndPointProvisioningState("Deleting")
	EndPointProvisioningStateSucceeded = EndPointProvisioningState("Succeeded")
	EndPointProvisioningStateCanceled  = EndPointProvisioningState("Canceled")
	EndPointProvisioningStateFailed    = EndPointProvisioningState("Failed")
)

type EntityStatus string

const (
	EntityStatusActive          = EntityStatus("Active")
	EntityStatusDisabled        = EntityStatus("Disabled")
	EntityStatusRestoring       = EntityStatus("Restoring")
	EntityStatusSendDisabled    = EntityStatus("SendDisabled")
	EntityStatusReceiveDisabled = EntityStatus("ReceiveDisabled")
	EntityStatusCreating        = EntityStatus("Creating")
	EntityStatusDeleting        = EntityStatus("Deleting")
	EntityStatusRenaming        = EntityStatus("Renaming")
	EntityStatusUnknown         = EntityStatus("Unknown")
)

func (EntityStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityStatus)(nil)).Elem()
}

func (e EntityStatus) ToEntityStatusOutput() EntityStatusOutput {
	return pulumi.ToOutput(e).(EntityStatusOutput)
}

func (e EntityStatus) ToEntityStatusOutputWithContext(ctx context.Context) EntityStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EntityStatusOutput)
}

func (e EntityStatus) ToEntityStatusPtrOutput() EntityStatusPtrOutput {
	return e.ToEntityStatusPtrOutputWithContext(context.Background())
}

func (e EntityStatus) ToEntityStatusPtrOutputWithContext(ctx context.Context) EntityStatusPtrOutput {
	return EntityStatus(e).ToEntityStatusOutputWithContext(ctx).ToEntityStatusPtrOutputWithContext(ctx)
}

func (e EntityStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EntityStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EntityStatusOutput struct{ *pulumi.OutputState }

func (EntityStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityStatus)(nil)).Elem()
}

func (o EntityStatusOutput) ToEntityStatusOutput() EntityStatusOutput {
	return o
}

func (o EntityStatusOutput) ToEntityStatusOutputWithContext(ctx context.Context) EntityStatusOutput {
	return o
}

func (o EntityStatusOutput) ToEntityStatusPtrOutput() EntityStatusPtrOutput {
	return o.ToEntityStatusPtrOutputWithContext(context.Background())
}

func (o EntityStatusOutput) ToEntityStatusPtrOutputWithContext(ctx context.Context) EntityStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntityStatus) *EntityStatus {
		return &v
	}).(EntityStatusPtrOutput)
}

func (o EntityStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EntityStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EntityStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EntityStatusPtrOutput struct{ *pulumi.OutputState }

func (EntityStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityStatus)(nil)).Elem()
}

func (o EntityStatusPtrOutput) ToEntityStatusPtrOutput() EntityStatusPtrOutput {
	return o
}

func (o EntityStatusPtrOutput) ToEntityStatusPtrOutputWithContext(ctx context.Context) EntityStatusPtrOutput {
	return o
}

func (o EntityStatusPtrOutput) Elem() EntityStatusOutput {
	return o.ApplyT(func(v *EntityStatus) EntityStatus {
		if v != nil {
			return *v
		}
		var ret EntityStatus
		return ret
	}).(EntityStatusOutput)
}

func (o EntityStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EntityStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EntityStatusInput interface {
	pulumi.Input

	ToEntityStatusOutput() EntityStatusOutput
	ToEntityStatusOutputWithContext(context.Context) EntityStatusOutput
}

var entityStatusPtrType = reflect.TypeOf((**EntityStatus)(nil)).Elem()

type EntityStatusPtrInput interface {
	pulumi.Input

	ToEntityStatusPtrOutput() EntityStatusPtrOutput
	ToEntityStatusPtrOutputWithContext(context.Context) EntityStatusPtrOutput
}

type entityStatusPtr string

func EntityStatusPtr(v string) EntityStatusPtrInput {
	return (*entityStatusPtr)(&v)
}

func (*entityStatusPtr) ElementType() reflect.Type {
	return entityStatusPtrType
}

func (in *entityStatusPtr) ToEntityStatusPtrOutput() EntityStatusPtrOutput {
	return pulumi.ToOutput(in).(EntityStatusPtrOutput)
}

func (in *entityStatusPtr) ToEntityStatusPtrOutputWithContext(ctx context.Context) EntityStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EntityStatusPtrOutput)
}

type FilterType string

const (
	FilterTypeSqlFilter         = FilterType("SqlFilter")
	FilterTypeCorrelationFilter = FilterType("CorrelationFilter")
)

type KeySource string

const (
	KeySource_Microsoft_KeyVault = KeySource("Microsoft.KeyVault")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
)

type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow = NetworkRuleIPAction("Allow")
)

type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusPending      = PrivateLinkConnectionStatus("Pending")
	PrivateLinkConnectionStatusApproved     = PrivateLinkConnectionStatus("Approved")
	PrivateLinkConnectionStatusRejected     = PrivateLinkConnectionStatus("Rejected")
	PrivateLinkConnectionStatusDisconnected = PrivateLinkConnectionStatus("Disconnected")
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled            = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled           = PublicNetworkAccess("Disabled")
	PublicNetworkAccessSecuredByPerimeter = PublicNetworkAccess("SecuredByPerimeter")
)

type PublicNetworkAccessFlag string

const (
	PublicNetworkAccessFlagEnabled  = PublicNetworkAccessFlag("Enabled")
	PublicNetworkAccessFlagDisabled = PublicNetworkAccessFlag("Disabled")
)

type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

type SkuTier string

const (
	SkuTierBasic    = SkuTier("Basic")
	SkuTierStandard = SkuTier("Standard")
	SkuTierPremium  = SkuTier("Premium")
)

type TlsVersion string

const (
	TlsVersion_1_0 = TlsVersion("1.0")
	TlsVersion_1_1 = TlsVersion("1.1")
	TlsVersion_1_2 = TlsVersion("1.2")
)

func init() {
	pulumi.RegisterOutputType(EntityStatusOutput{})
	pulumi.RegisterOutputType(EntityStatusPtrOutput{})
}
