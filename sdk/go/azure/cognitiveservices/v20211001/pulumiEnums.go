


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeploymentScaleType string

const (
	DeploymentScaleTypeManual = DeploymentScaleType("Manual")
)

func (DeploymentScaleType) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentScaleType)(nil)).Elem()
}

func (e DeploymentScaleType) ToDeploymentScaleTypeOutput() DeploymentScaleTypeOutput {
	return pulumi.ToOutput(e).(DeploymentScaleTypeOutput)
}

func (e DeploymentScaleType) ToDeploymentScaleTypeOutputWithContext(ctx context.Context) DeploymentScaleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeploymentScaleTypeOutput)
}

func (e DeploymentScaleType) ToDeploymentScaleTypePtrOutput() DeploymentScaleTypePtrOutput {
	return e.ToDeploymentScaleTypePtrOutputWithContext(context.Background())
}

func (e DeploymentScaleType) ToDeploymentScaleTypePtrOutputWithContext(ctx context.Context) DeploymentScaleTypePtrOutput {
	return DeploymentScaleType(e).ToDeploymentScaleTypeOutputWithContext(ctx).ToDeploymentScaleTypePtrOutputWithContext(ctx)
}

func (e DeploymentScaleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeploymentScaleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeploymentScaleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeploymentScaleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeploymentScaleTypeOutput struct{ *pulumi.OutputState }

func (DeploymentScaleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentScaleType)(nil)).Elem()
}

func (o DeploymentScaleTypeOutput) ToDeploymentScaleTypeOutput() DeploymentScaleTypeOutput {
	return o
}

func (o DeploymentScaleTypeOutput) ToDeploymentScaleTypeOutputWithContext(ctx context.Context) DeploymentScaleTypeOutput {
	return o
}

func (o DeploymentScaleTypeOutput) ToDeploymentScaleTypePtrOutput() DeploymentScaleTypePtrOutput {
	return o.ToDeploymentScaleTypePtrOutputWithContext(context.Background())
}

func (o DeploymentScaleTypeOutput) ToDeploymentScaleTypePtrOutputWithContext(ctx context.Context) DeploymentScaleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentScaleType) *DeploymentScaleType {
		return &v
	}).(DeploymentScaleTypePtrOutput)
}

func (o DeploymentScaleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeploymentScaleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeploymentScaleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeploymentScaleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeploymentScaleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeploymentScaleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeploymentScaleTypePtrOutput struct{ *pulumi.OutputState }

func (DeploymentScaleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentScaleType)(nil)).Elem()
}

func (o DeploymentScaleTypePtrOutput) ToDeploymentScaleTypePtrOutput() DeploymentScaleTypePtrOutput {
	return o
}

func (o DeploymentScaleTypePtrOutput) ToDeploymentScaleTypePtrOutputWithContext(ctx context.Context) DeploymentScaleTypePtrOutput {
	return o
}

func (o DeploymentScaleTypePtrOutput) Elem() DeploymentScaleTypeOutput {
	return o.ApplyT(func(v *DeploymentScaleType) DeploymentScaleType {
		if v != nil {
			return *v
		}
		var ret DeploymentScaleType
		return ret
	}).(DeploymentScaleTypeOutput)
}

func (o DeploymentScaleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeploymentScaleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeploymentScaleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeploymentScaleTypeInput interface {
	pulumi.Input

	ToDeploymentScaleTypeOutput() DeploymentScaleTypeOutput
	ToDeploymentScaleTypeOutputWithContext(context.Context) DeploymentScaleTypeOutput
}

var deploymentScaleTypePtrType = reflect.TypeOf((**DeploymentScaleType)(nil)).Elem()

type DeploymentScaleTypePtrInput interface {
	pulumi.Input

	ToDeploymentScaleTypePtrOutput() DeploymentScaleTypePtrOutput
	ToDeploymentScaleTypePtrOutputWithContext(context.Context) DeploymentScaleTypePtrOutput
}

type deploymentScaleTypePtr string

func DeploymentScaleTypePtr(v string) DeploymentScaleTypePtrInput {
	return (*deploymentScaleTypePtr)(&v)
}

func (*deploymentScaleTypePtr) ElementType() reflect.Type {
	return deploymentScaleTypePtrType
}

func (in *deploymentScaleTypePtr) ToDeploymentScaleTypePtrOutput() DeploymentScaleTypePtrOutput {
	return pulumi.ToOutput(in).(DeploymentScaleTypePtrOutput)
}

func (in *deploymentScaleTypePtr) ToDeploymentScaleTypePtrOutputWithContext(ctx context.Context) DeploymentScaleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeploymentScaleTypePtrOutput)
}

type HostingModel string

const (
	HostingModelWeb                   = HostingModel("Web")
	HostingModelConnectedContainer    = HostingModel("ConnectedContainer")
	HostingModelDisconnectedContainer = HostingModel("DisconnectedContainer")
)

func (HostingModel) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingModel)(nil)).Elem()
}

func (e HostingModel) ToHostingModelOutput() HostingModelOutput {
	return pulumi.ToOutput(e).(HostingModelOutput)
}

func (e HostingModel) ToHostingModelOutputWithContext(ctx context.Context) HostingModelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostingModelOutput)
}

func (e HostingModel) ToHostingModelPtrOutput() HostingModelPtrOutput {
	return e.ToHostingModelPtrOutputWithContext(context.Background())
}

func (e HostingModel) ToHostingModelPtrOutputWithContext(ctx context.Context) HostingModelPtrOutput {
	return HostingModel(e).ToHostingModelOutputWithContext(ctx).ToHostingModelPtrOutputWithContext(ctx)
}

func (e HostingModel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingModel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingModel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostingModel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostingModelOutput struct{ *pulumi.OutputState }

func (HostingModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingModel)(nil)).Elem()
}

func (o HostingModelOutput) ToHostingModelOutput() HostingModelOutput {
	return o
}

func (o HostingModelOutput) ToHostingModelOutputWithContext(ctx context.Context) HostingModelOutput {
	return o
}

func (o HostingModelOutput) ToHostingModelPtrOutput() HostingModelPtrOutput {
	return o.ToHostingModelPtrOutputWithContext(context.Background())
}

func (o HostingModelOutput) ToHostingModelPtrOutputWithContext(ctx context.Context) HostingModelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingModel) *HostingModel {
		return &v
	}).(HostingModelPtrOutput)
}

func (o HostingModelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostingModelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostingModel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostingModelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostingModelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostingModel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostingModelPtrOutput struct{ *pulumi.OutputState }

func (HostingModelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingModel)(nil)).Elem()
}

func (o HostingModelPtrOutput) ToHostingModelPtrOutput() HostingModelPtrOutput {
	return o
}

func (o HostingModelPtrOutput) ToHostingModelPtrOutputWithContext(ctx context.Context) HostingModelPtrOutput {
	return o
}

func (o HostingModelPtrOutput) Elem() HostingModelOutput {
	return o.ApplyT(func(v *HostingModel) HostingModel {
		if v != nil {
			return *v
		}
		var ret HostingModel
		return ret
	}).(HostingModelOutput)
}

func (o HostingModelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostingModelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostingModel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostingModelInput interface {
	pulumi.Input

	ToHostingModelOutput() HostingModelOutput
	ToHostingModelOutputWithContext(context.Context) HostingModelOutput
}

var hostingModelPtrType = reflect.TypeOf((**HostingModel)(nil)).Elem()

type HostingModelPtrInput interface {
	pulumi.Input

	ToHostingModelPtrOutput() HostingModelPtrOutput
	ToHostingModelPtrOutputWithContext(context.Context) HostingModelPtrOutput
}

type hostingModelPtr string

func HostingModelPtr(v string) HostingModelPtrInput {
	return (*hostingModelPtr)(&v)
}

func (*hostingModelPtr) ElementType() reflect.Type {
	return hostingModelPtrType
}

func (in *hostingModelPtr) ToHostingModelPtrOutput() HostingModelPtrOutput {
	return pulumi.ToOutput(in).(HostingModelPtrOutput)
}

func (in *hostingModelPtr) ToHostingModelPtrOutputWithContext(ctx context.Context) HostingModelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostingModelPtrOutput)
}

type KeySource string

const (
	KeySource_Microsoft_CognitiveServices = KeySource("Microsoft.CognitiveServices")
	KeySource_Microsoft_KeyVault          = KeySource("Microsoft.KeyVault")
)

func (KeySource) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySource)(nil)).Elem()
}

func (e KeySource) ToKeySourceOutput() KeySourceOutput {
	return pulumi.ToOutput(e).(KeySourceOutput)
}

func (e KeySource) ToKeySourceOutputWithContext(ctx context.Context) KeySourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeySourceOutput)
}

func (e KeySource) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return e.ToKeySourcePtrOutputWithContext(context.Background())
}

func (e KeySource) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return KeySource(e).ToKeySourceOutputWithContext(ctx).ToKeySourcePtrOutputWithContext(ctx)
}

func (e KeySource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeySource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeySourceOutput struct{ *pulumi.OutputState }

func (KeySourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySource)(nil)).Elem()
}

func (o KeySourceOutput) ToKeySourceOutput() KeySourceOutput {
	return o
}

func (o KeySourceOutput) ToKeySourceOutputWithContext(ctx context.Context) KeySourceOutput {
	return o
}

func (o KeySourceOutput) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return o.ToKeySourcePtrOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeySource) *KeySource {
		return &v
	}).(KeySourcePtrOutput)
}

func (o KeySourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeySource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeySourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeySource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeySourcePtrOutput struct{ *pulumi.OutputState }

func (KeySourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeySource)(nil)).Elem()
}

func (o KeySourcePtrOutput) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return o
}

func (o KeySourcePtrOutput) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return o
}

func (o KeySourcePtrOutput) Elem() KeySourceOutput {
	return o.ApplyT(func(v *KeySource) KeySource {
		if v != nil {
			return *v
		}
		var ret KeySource
		return ret
	}).(KeySourceOutput)
}

func (o KeySourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeySourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeySource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KeySourceInput interface {
	pulumi.Input

	ToKeySourceOutput() KeySourceOutput
	ToKeySourceOutputWithContext(context.Context) KeySourceOutput
}

var keySourcePtrType = reflect.TypeOf((**KeySource)(nil)).Elem()

type KeySourcePtrInput interface {
	pulumi.Input

	ToKeySourcePtrOutput() KeySourcePtrOutput
	ToKeySourcePtrOutputWithContext(context.Context) KeySourcePtrOutput
}

type keySourcePtr string

func KeySourcePtr(v string) KeySourcePtrInput {
	return (*keySourcePtr)(&v)
}

func (*keySourcePtr) ElementType() reflect.Type {
	return keySourcePtrType
}

func (in *keySourcePtr) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return pulumi.ToOutput(in).(KeySourcePtrOutput)
}

func (in *keySourcePtr) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeySourcePtrOutput)
}

type NetworkRuleAction string

const (
	NetworkRuleActionAllow = NetworkRuleAction("Allow")
	NetworkRuleActionDeny  = NetworkRuleAction("Deny")
)

func (NetworkRuleAction) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleAction)(nil)).Elem()
}

func (e NetworkRuleAction) ToNetworkRuleActionOutput() NetworkRuleActionOutput {
	return pulumi.ToOutput(e).(NetworkRuleActionOutput)
}

func (e NetworkRuleAction) ToNetworkRuleActionOutputWithContext(ctx context.Context) NetworkRuleActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkRuleActionOutput)
}

func (e NetworkRuleAction) ToNetworkRuleActionPtrOutput() NetworkRuleActionPtrOutput {
	return e.ToNetworkRuleActionPtrOutputWithContext(context.Background())
}

func (e NetworkRuleAction) ToNetworkRuleActionPtrOutputWithContext(ctx context.Context) NetworkRuleActionPtrOutput {
	return NetworkRuleAction(e).ToNetworkRuleActionOutputWithContext(ctx).ToNetworkRuleActionPtrOutputWithContext(ctx)
}

func (e NetworkRuleAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkRuleAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkRuleAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkRuleAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkRuleActionOutput struct{ *pulumi.OutputState }

func (NetworkRuleActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleAction)(nil)).Elem()
}

func (o NetworkRuleActionOutput) ToNetworkRuleActionOutput() NetworkRuleActionOutput {
	return o
}

func (o NetworkRuleActionOutput) ToNetworkRuleActionOutputWithContext(ctx context.Context) NetworkRuleActionOutput {
	return o
}

func (o NetworkRuleActionOutput) ToNetworkRuleActionPtrOutput() NetworkRuleActionPtrOutput {
	return o.ToNetworkRuleActionPtrOutputWithContext(context.Background())
}

func (o NetworkRuleActionOutput) ToNetworkRuleActionPtrOutputWithContext(ctx context.Context) NetworkRuleActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleAction) *NetworkRuleAction {
		return &v
	}).(NetworkRuleActionPtrOutput)
}

func (o NetworkRuleActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkRuleActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkRuleAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkRuleActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkRuleActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkRuleAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkRuleActionPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleAction)(nil)).Elem()
}

func (o NetworkRuleActionPtrOutput) ToNetworkRuleActionPtrOutput() NetworkRuleActionPtrOutput {
	return o
}

func (o NetworkRuleActionPtrOutput) ToNetworkRuleActionPtrOutputWithContext(ctx context.Context) NetworkRuleActionPtrOutput {
	return o
}

func (o NetworkRuleActionPtrOutput) Elem() NetworkRuleActionOutput {
	return o.ApplyT(func(v *NetworkRuleAction) NetworkRuleAction {
		if v != nil {
			return *v
		}
		var ret NetworkRuleAction
		return ret
	}).(NetworkRuleActionOutput)
}

func (o NetworkRuleActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkRuleActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkRuleAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkRuleActionInput interface {
	pulumi.Input

	ToNetworkRuleActionOutput() NetworkRuleActionOutput
	ToNetworkRuleActionOutputWithContext(context.Context) NetworkRuleActionOutput
}

var networkRuleActionPtrType = reflect.TypeOf((**NetworkRuleAction)(nil)).Elem()

type NetworkRuleActionPtrInput interface {
	pulumi.Input

	ToNetworkRuleActionPtrOutput() NetworkRuleActionPtrOutput
	ToNetworkRuleActionPtrOutputWithContext(context.Context) NetworkRuleActionPtrOutput
}

type networkRuleActionPtr string

func NetworkRuleActionPtr(v string) NetworkRuleActionPtrInput {
	return (*networkRuleActionPtr)(&v)
}

func (*networkRuleActionPtr) ElementType() reflect.Type {
	return networkRuleActionPtrType
}

func (in *networkRuleActionPtr) ToNetworkRuleActionPtrOutput() NetworkRuleActionPtrOutput {
	return pulumi.ToOutput(in).(NetworkRuleActionPtrOutput)
}

func (in *networkRuleActionPtr) ToNetworkRuleActionPtrOutputWithContext(ctx context.Context) NetworkRuleActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkRuleActionPtrOutput)
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
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
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

type ResourceIdentityType string

const (
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type SkuTier string

const (
	SkuTierFree       = SkuTier("Free")
	SkuTierBasic      = SkuTier("Basic")
	SkuTierStandard   = SkuTier("Standard")
	SkuTierPremium    = SkuTier("Premium")
	SkuTierEnterprise = SkuTier("Enterprise")
)

func (SkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (e SkuTier) ToSkuTierOutput() SkuTierOutput {
	return pulumi.ToOutput(e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return e.ToSkuTierPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return SkuTier(e).ToSkuTierOutputWithContext(ctx).ToSkuTierPtrOutputWithContext(ctx)
}

func (e SkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTierOutput struct{ *pulumi.OutputState }

func (SkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (o SkuTierOutput) ToSkuTierOutput() SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o.ToSkuTierPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuTier) *SkuTier {
		return &v
	}).(SkuTierPtrOutput)
}

func (o SkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTierPtrOutput struct{ *pulumi.OutputState }

func (SkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuTier)(nil)).Elem()
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) Elem() SkuTierOutput {
	return o.ApplyT(func(v *SkuTier) SkuTier {
		if v != nil {
			return *v
		}
		var ret SkuTier
		return ret
	}).(SkuTierOutput)
}

func (o SkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuTierInput interface {
	pulumi.Input

	ToSkuTierOutput() SkuTierOutput
	ToSkuTierOutputWithContext(context.Context) SkuTierOutput
}

var skuTierPtrType = reflect.TypeOf((**SkuTier)(nil)).Elem()

type SkuTierPtrInput interface {
	pulumi.Input

	ToSkuTierPtrOutput() SkuTierPtrOutput
	ToSkuTierPtrOutputWithContext(context.Context) SkuTierPtrOutput
}

type skuTierPtr string

func SkuTierPtr(v string) SkuTierPtrInput {
	return (*skuTierPtr)(&v)
}

func (*skuTierPtr) ElementType() reflect.Type {
	return skuTierPtrType
}

func (in *skuTierPtr) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return pulumi.ToOutput(in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTierPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentScaleTypeOutput{})
	pulumi.RegisterOutputType(DeploymentScaleTypePtrOutput{})
	pulumi.RegisterOutputType(HostingModelOutput{})
	pulumi.RegisterOutputType(HostingModelPtrOutput{})
	pulumi.RegisterOutputType(KeySourceOutput{})
	pulumi.RegisterOutputType(KeySourcePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleActionOutput{})
	pulumi.RegisterOutputType(NetworkRuleActionPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
}
