


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPoolMode string

const (
	// System agent pools are primarily for hosting critical system pods such as CoreDNS and metrics-server. System agent pools osType must be Linux. System agent pools VM SKU must have at least 2vCPUs and 4GB of memory.
	AgentPoolModeSystem = AgentPoolMode("System")
	// User agent pools are primarily for hosting your application pods.
	AgentPoolModeUser = AgentPoolMode("User")
)

func (AgentPoolMode) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolMode)(nil)).Elem()
}

func (e AgentPoolMode) ToAgentPoolModeOutput() AgentPoolModeOutput {
	return pulumi.ToOutput(e).(AgentPoolModeOutput)
}

func (e AgentPoolMode) ToAgentPoolModeOutputWithContext(ctx context.Context) AgentPoolModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AgentPoolModeOutput)
}

func (e AgentPoolMode) ToAgentPoolModePtrOutput() AgentPoolModePtrOutput {
	return e.ToAgentPoolModePtrOutputWithContext(context.Background())
}

func (e AgentPoolMode) ToAgentPoolModePtrOutputWithContext(ctx context.Context) AgentPoolModePtrOutput {
	return AgentPoolMode(e).ToAgentPoolModeOutputWithContext(ctx).ToAgentPoolModePtrOutputWithContext(ctx)
}

func (e AgentPoolMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgentPoolMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgentPoolMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AgentPoolMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AgentPoolModeOutput struct{ *pulumi.OutputState }

func (AgentPoolModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolMode)(nil)).Elem()
}

func (o AgentPoolModeOutput) ToAgentPoolModeOutput() AgentPoolModeOutput {
	return o
}

func (o AgentPoolModeOutput) ToAgentPoolModeOutputWithContext(ctx context.Context) AgentPoolModeOutput {
	return o
}

func (o AgentPoolModeOutput) ToAgentPoolModePtrOutput() AgentPoolModePtrOutput {
	return o.ToAgentPoolModePtrOutputWithContext(context.Background())
}

func (o AgentPoolModeOutput) ToAgentPoolModePtrOutputWithContext(ctx context.Context) AgentPoolModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPoolMode) *AgentPoolMode {
		return &v
	}).(AgentPoolModePtrOutput)
}

func (o AgentPoolModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AgentPoolModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgentPoolMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AgentPoolModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgentPoolModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgentPoolMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AgentPoolModePtrOutput struct{ *pulumi.OutputState }

func (AgentPoolModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolMode)(nil)).Elem()
}

func (o AgentPoolModePtrOutput) ToAgentPoolModePtrOutput() AgentPoolModePtrOutput {
	return o
}

func (o AgentPoolModePtrOutput) ToAgentPoolModePtrOutputWithContext(ctx context.Context) AgentPoolModePtrOutput {
	return o
}

func (o AgentPoolModePtrOutput) Elem() AgentPoolModeOutput {
	return o.ApplyT(func(v *AgentPoolMode) AgentPoolMode {
		if v != nil {
			return *v
		}
		var ret AgentPoolMode
		return ret
	}).(AgentPoolModeOutput)
}

func (o AgentPoolModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgentPoolModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AgentPoolMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AgentPoolModeInput interface {
	pulumi.Input

	ToAgentPoolModeOutput() AgentPoolModeOutput
	ToAgentPoolModeOutputWithContext(context.Context) AgentPoolModeOutput
}

var agentPoolModePtrType = reflect.TypeOf((**AgentPoolMode)(nil)).Elem()

type AgentPoolModePtrInput interface {
	pulumi.Input

	ToAgentPoolModePtrOutput() AgentPoolModePtrOutput
	ToAgentPoolModePtrOutputWithContext(context.Context) AgentPoolModePtrOutput
}

type agentPoolModePtr string

func AgentPoolModePtr(v string) AgentPoolModePtrInput {
	return (*agentPoolModePtr)(&v)
}

func (*agentPoolModePtr) ElementType() reflect.Type {
	return agentPoolModePtrType
}

func (in *agentPoolModePtr) ToAgentPoolModePtrOutput() AgentPoolModePtrOutput {
	return pulumi.ToOutput(in).(AgentPoolModePtrOutput)
}

func (in *agentPoolModePtr) ToAgentPoolModePtrOutputWithContext(ctx context.Context) AgentPoolModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AgentPoolModePtrOutput)
}

type AgentPoolType string

const (
	// Create an Agent Pool backed by a Virtual Machine Scale Set.
	AgentPoolTypeVirtualMachineScaleSets = AgentPoolType("VirtualMachineScaleSets")
	// Use of this is strongly discouraged.
	AgentPoolTypeAvailabilitySet = AgentPoolType("AvailabilitySet")
)

func (AgentPoolType) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolType)(nil)).Elem()
}

func (e AgentPoolType) ToAgentPoolTypeOutput() AgentPoolTypeOutput {
	return pulumi.ToOutput(e).(AgentPoolTypeOutput)
}

func (e AgentPoolType) ToAgentPoolTypeOutputWithContext(ctx context.Context) AgentPoolTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AgentPoolTypeOutput)
}

func (e AgentPoolType) ToAgentPoolTypePtrOutput() AgentPoolTypePtrOutput {
	return e.ToAgentPoolTypePtrOutputWithContext(context.Background())
}

func (e AgentPoolType) ToAgentPoolTypePtrOutputWithContext(ctx context.Context) AgentPoolTypePtrOutput {
	return AgentPoolType(e).ToAgentPoolTypeOutputWithContext(ctx).ToAgentPoolTypePtrOutputWithContext(ctx)
}

func (e AgentPoolType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgentPoolType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgentPoolType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AgentPoolType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AgentPoolTypeOutput struct{ *pulumi.OutputState }

func (AgentPoolTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPoolType)(nil)).Elem()
}

func (o AgentPoolTypeOutput) ToAgentPoolTypeOutput() AgentPoolTypeOutput {
	return o
}

func (o AgentPoolTypeOutput) ToAgentPoolTypeOutputWithContext(ctx context.Context) AgentPoolTypeOutput {
	return o
}

func (o AgentPoolTypeOutput) ToAgentPoolTypePtrOutput() AgentPoolTypePtrOutput {
	return o.ToAgentPoolTypePtrOutputWithContext(context.Background())
}

func (o AgentPoolTypeOutput) ToAgentPoolTypePtrOutputWithContext(ctx context.Context) AgentPoolTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentPoolType) *AgentPoolType {
		return &v
	}).(AgentPoolTypePtrOutput)
}

func (o AgentPoolTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AgentPoolTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgentPoolType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AgentPoolTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgentPoolTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgentPoolType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AgentPoolTypePtrOutput struct{ *pulumi.OutputState }

func (AgentPoolTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPoolType)(nil)).Elem()
}

func (o AgentPoolTypePtrOutput) ToAgentPoolTypePtrOutput() AgentPoolTypePtrOutput {
	return o
}

func (o AgentPoolTypePtrOutput) ToAgentPoolTypePtrOutputWithContext(ctx context.Context) AgentPoolTypePtrOutput {
	return o
}

func (o AgentPoolTypePtrOutput) Elem() AgentPoolTypeOutput {
	return o.ApplyT(func(v *AgentPoolType) AgentPoolType {
		if v != nil {
			return *v
		}
		var ret AgentPoolType
		return ret
	}).(AgentPoolTypeOutput)
}

func (o AgentPoolTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgentPoolTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AgentPoolType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AgentPoolTypeInput interface {
	pulumi.Input

	ToAgentPoolTypeOutput() AgentPoolTypeOutput
	ToAgentPoolTypeOutputWithContext(context.Context) AgentPoolTypeOutput
}

var agentPoolTypePtrType = reflect.TypeOf((**AgentPoolType)(nil)).Elem()

type AgentPoolTypePtrInput interface {
	pulumi.Input

	ToAgentPoolTypePtrOutput() AgentPoolTypePtrOutput
	ToAgentPoolTypePtrOutputWithContext(context.Context) AgentPoolTypePtrOutput
}

type agentPoolTypePtr string

func AgentPoolTypePtr(v string) AgentPoolTypePtrInput {
	return (*agentPoolTypePtr)(&v)
}

func (*agentPoolTypePtr) ElementType() reflect.Type {
	return agentPoolTypePtrType
}

func (in *agentPoolTypePtr) ToAgentPoolTypePtrOutput() AgentPoolTypePtrOutput {
	return pulumi.ToOutput(in).(AgentPoolTypePtrOutput)
}

func (in *agentPoolTypePtr) ToAgentPoolTypePtrOutputWithContext(ctx context.Context) AgentPoolTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AgentPoolTypePtrOutput)
}

type ConnectionStatus string

const (
	ConnectionStatusPending      = ConnectionStatus("Pending")
	ConnectionStatusApproved     = ConnectionStatus("Approved")
	ConnectionStatusRejected     = ConnectionStatus("Rejected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
)

func (ConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatus)(nil)).Elem()
}

func (e ConnectionStatus) ToConnectionStatusOutput() ConnectionStatusOutput {
	return pulumi.ToOutput(e).(ConnectionStatusOutput)
}

func (e ConnectionStatus) ToConnectionStatusOutputWithContext(ctx context.Context) ConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionStatusOutput)
}

func (e ConnectionStatus) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return e.ToConnectionStatusPtrOutputWithContext(context.Background())
}

func (e ConnectionStatus) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return ConnectionStatus(e).ToConnectionStatusOutputWithContext(ctx).ToConnectionStatusPtrOutputWithContext(ctx)
}

func (e ConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionStatusOutput struct{ *pulumi.OutputState }

func (ConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatus)(nil)).Elem()
}

func (o ConnectionStatusOutput) ToConnectionStatusOutput() ConnectionStatusOutput {
	return o
}

func (o ConnectionStatusOutput) ToConnectionStatusOutputWithContext(ctx context.Context) ConnectionStatusOutput {
	return o
}

func (o ConnectionStatusOutput) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return o.ToConnectionStatusPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionStatus) *ConnectionStatus {
		return &v
	}).(ConnectionStatusPtrOutput)
}

func (o ConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStatus)(nil)).Elem()
}

func (o ConnectionStatusPtrOutput) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return o
}

func (o ConnectionStatusPtrOutput) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return o
}

func (o ConnectionStatusPtrOutput) Elem() ConnectionStatusOutput {
	return o.ApplyT(func(v *ConnectionStatus) ConnectionStatus {
		if v != nil {
			return *v
		}
		var ret ConnectionStatus
		return ret
	}).(ConnectionStatusOutput)
}

func (o ConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionStatusInput interface {
	pulumi.Input

	ToConnectionStatusOutput() ConnectionStatusOutput
	ToConnectionStatusOutputWithContext(context.Context) ConnectionStatusOutput
}

var connectionStatusPtrType = reflect.TypeOf((**ConnectionStatus)(nil)).Elem()

type ConnectionStatusPtrInput interface {
	pulumi.Input

	ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput
	ToConnectionStatusPtrOutputWithContext(context.Context) ConnectionStatusPtrOutput
}

type connectionStatusPtr string

func ConnectionStatusPtr(v string) ConnectionStatusPtrInput {
	return (*connectionStatusPtr)(&v)
}

func (*connectionStatusPtr) ElementType() reflect.Type {
	return connectionStatusPtrType
}

func (in *connectionStatusPtr) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(ConnectionStatusPtrOutput)
}

func (in *connectionStatusPtr) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionStatusPtrOutput)
}

type Expander string

const (
	// Selects the node group that will have the least idle CPU (if tied, unused memory) after scale-up. This is useful when you have different classes of nodes, for example, high CPU or high memory nodes, and only want to expand those when there are pending pods that need a lot of those resources.
	Expander_Least_waste = Expander("least-waste")
	// Selects the node group that would be able to schedule the most pods when scaling up. This is useful when you are using nodeSelector to make sure certain pods land on certain nodes. Note that this won't cause the autoscaler to select bigger nodes vs. smaller, as it can add multiple smaller nodes at once.
	Expander_Most_pods = Expander("most-pods")
	// Selects the node group that has the highest priority assigned by the user. It's configuration is described in more details [here](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/expander/priority/readme.md).
	ExpanderPriority = Expander("priority")
	// Used when you don't have a particular need for the node groups to scale differently.
	ExpanderRandom = Expander("random")
)

func (Expander) ElementType() reflect.Type {
	return reflect.TypeOf((*Expander)(nil)).Elem()
}

func (e Expander) ToExpanderOutput() ExpanderOutput {
	return pulumi.ToOutput(e).(ExpanderOutput)
}

func (e Expander) ToExpanderOutputWithContext(ctx context.Context) ExpanderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpanderOutput)
}

func (e Expander) ToExpanderPtrOutput() ExpanderPtrOutput {
	return e.ToExpanderPtrOutputWithContext(context.Background())
}

func (e Expander) ToExpanderPtrOutputWithContext(ctx context.Context) ExpanderPtrOutput {
	return Expander(e).ToExpanderOutputWithContext(ctx).ToExpanderPtrOutputWithContext(ctx)
}

func (e Expander) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Expander) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Expander) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Expander) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpanderOutput struct{ *pulumi.OutputState }

func (ExpanderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Expander)(nil)).Elem()
}

func (o ExpanderOutput) ToExpanderOutput() ExpanderOutput {
	return o
}

func (o ExpanderOutput) ToExpanderOutputWithContext(ctx context.Context) ExpanderOutput {
	return o
}

func (o ExpanderOutput) ToExpanderPtrOutput() ExpanderPtrOutput {
	return o.ToExpanderPtrOutputWithContext(context.Background())
}

func (o ExpanderOutput) ToExpanderPtrOutputWithContext(ctx context.Context) ExpanderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Expander) *Expander {
		return &v
	}).(ExpanderPtrOutput)
}

func (o ExpanderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpanderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Expander) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpanderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpanderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Expander) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpanderPtrOutput struct{ *pulumi.OutputState }

func (ExpanderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Expander)(nil)).Elem()
}

func (o ExpanderPtrOutput) ToExpanderPtrOutput() ExpanderPtrOutput {
	return o
}

func (o ExpanderPtrOutput) ToExpanderPtrOutputWithContext(ctx context.Context) ExpanderPtrOutput {
	return o
}

func (o ExpanderPtrOutput) Elem() ExpanderOutput {
	return o.ApplyT(func(v *Expander) Expander {
		if v != nil {
			return *v
		}
		var ret Expander
		return ret
	}).(ExpanderOutput)
}

func (o ExpanderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpanderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Expander) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpanderInput interface {
	pulumi.Input

	ToExpanderOutput() ExpanderOutput
	ToExpanderOutputWithContext(context.Context) ExpanderOutput
}

var expanderPtrType = reflect.TypeOf((**Expander)(nil)).Elem()

type ExpanderPtrInput interface {
	pulumi.Input

	ToExpanderPtrOutput() ExpanderPtrOutput
	ToExpanderPtrOutputWithContext(context.Context) ExpanderPtrOutput
}

type expanderPtr string

func ExpanderPtr(v string) ExpanderPtrInput {
	return (*expanderPtr)(&v)
}

func (*expanderPtr) ElementType() reflect.Type {
	return expanderPtrType
}

func (in *expanderPtr) ToExpanderPtrOutput() ExpanderPtrOutput {
	return pulumi.ToOutput(in).(ExpanderPtrOutput)
}

func (in *expanderPtr) ToExpanderPtrOutputWithContext(ctx context.Context) ExpanderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpanderPtrOutput)
}

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone = ExtendedLocationTypes("EdgeZone")
)

func (ExtendedLocationTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationTypes)(nil)).Elem()
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput {
	return pulumi.ToOutput(e).(ExtendedLocationTypesOutput)
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesOutputWithContext(ctx context.Context) ExtendedLocationTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExtendedLocationTypesOutput)
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return e.ToExtendedLocationTypesPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return ExtendedLocationTypes(e).ToExtendedLocationTypesOutputWithContext(ctx).ToExtendedLocationTypesPtrOutputWithContext(ctx)
}

func (e ExtendedLocationTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExtendedLocationTypesOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationTypes)(nil)).Elem()
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput {
	return o
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesOutputWithContext(ctx context.Context) ExtendedLocationTypesOutput {
	return o
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return o.ToExtendedLocationTypesPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationTypes) *ExtendedLocationTypes {
		return &v
	}).(ExtendedLocationTypesPtrOutput)
}

func (o ExtendedLocationTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExtendedLocationTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationTypesPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationTypes)(nil)).Elem()
}

func (o ExtendedLocationTypesPtrOutput) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return o
}

func (o ExtendedLocationTypesPtrOutput) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return o
}

func (o ExtendedLocationTypesPtrOutput) Elem() ExtendedLocationTypesOutput {
	return o.ApplyT(func(v *ExtendedLocationTypes) ExtendedLocationTypes {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationTypes
		return ret
	}).(ExtendedLocationTypesOutput)
}

func (o ExtendedLocationTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExtendedLocationTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExtendedLocationTypesInput interface {
	pulumi.Input

	ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput
	ToExtendedLocationTypesOutputWithContext(context.Context) ExtendedLocationTypesOutput
}

var extendedLocationTypesPtrType = reflect.TypeOf((**ExtendedLocationTypes)(nil)).Elem()

type ExtendedLocationTypesPtrInput interface {
	pulumi.Input

	ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput
	ToExtendedLocationTypesPtrOutputWithContext(context.Context) ExtendedLocationTypesPtrOutput
}

type extendedLocationTypesPtr string

func ExtendedLocationTypesPtr(v string) ExtendedLocationTypesPtrInput {
	return (*extendedLocationTypesPtr)(&v)
}

func (*extendedLocationTypesPtr) ElementType() reflect.Type {
	return extendedLocationTypesPtrType
}

func (in *extendedLocationTypesPtr) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return pulumi.ToOutput(in).(ExtendedLocationTypesPtrOutput)
}

func (in *extendedLocationTypesPtr) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExtendedLocationTypesPtrOutput)
}

type GPUInstanceProfile string

const (
	GPUInstanceProfileMIG1g = GPUInstanceProfile("MIG1g")
	GPUInstanceProfileMIG2g = GPUInstanceProfile("MIG2g")
	GPUInstanceProfileMIG3g = GPUInstanceProfile("MIG3g")
	GPUInstanceProfileMIG4g = GPUInstanceProfile("MIG4g")
	GPUInstanceProfileMIG7g = GPUInstanceProfile("MIG7g")
)

func (GPUInstanceProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*GPUInstanceProfile)(nil)).Elem()
}

func (e GPUInstanceProfile) ToGPUInstanceProfileOutput() GPUInstanceProfileOutput {
	return pulumi.ToOutput(e).(GPUInstanceProfileOutput)
}

func (e GPUInstanceProfile) ToGPUInstanceProfileOutputWithContext(ctx context.Context) GPUInstanceProfileOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GPUInstanceProfileOutput)
}

func (e GPUInstanceProfile) ToGPUInstanceProfilePtrOutput() GPUInstanceProfilePtrOutput {
	return e.ToGPUInstanceProfilePtrOutputWithContext(context.Background())
}

func (e GPUInstanceProfile) ToGPUInstanceProfilePtrOutputWithContext(ctx context.Context) GPUInstanceProfilePtrOutput {
	return GPUInstanceProfile(e).ToGPUInstanceProfileOutputWithContext(ctx).ToGPUInstanceProfilePtrOutputWithContext(ctx)
}

func (e GPUInstanceProfile) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GPUInstanceProfile) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GPUInstanceProfile) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GPUInstanceProfile) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GPUInstanceProfileOutput struct{ *pulumi.OutputState }

func (GPUInstanceProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GPUInstanceProfile)(nil)).Elem()
}

func (o GPUInstanceProfileOutput) ToGPUInstanceProfileOutput() GPUInstanceProfileOutput {
	return o
}

func (o GPUInstanceProfileOutput) ToGPUInstanceProfileOutputWithContext(ctx context.Context) GPUInstanceProfileOutput {
	return o
}

func (o GPUInstanceProfileOutput) ToGPUInstanceProfilePtrOutput() GPUInstanceProfilePtrOutput {
	return o.ToGPUInstanceProfilePtrOutputWithContext(context.Background())
}

func (o GPUInstanceProfileOutput) ToGPUInstanceProfilePtrOutputWithContext(ctx context.Context) GPUInstanceProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GPUInstanceProfile) *GPUInstanceProfile {
		return &v
	}).(GPUInstanceProfilePtrOutput)
}

func (o GPUInstanceProfileOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GPUInstanceProfileOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GPUInstanceProfile) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GPUInstanceProfileOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GPUInstanceProfileOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GPUInstanceProfile) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GPUInstanceProfilePtrOutput struct{ *pulumi.OutputState }

func (GPUInstanceProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GPUInstanceProfile)(nil)).Elem()
}

func (o GPUInstanceProfilePtrOutput) ToGPUInstanceProfilePtrOutput() GPUInstanceProfilePtrOutput {
	return o
}

func (o GPUInstanceProfilePtrOutput) ToGPUInstanceProfilePtrOutputWithContext(ctx context.Context) GPUInstanceProfilePtrOutput {
	return o
}

func (o GPUInstanceProfilePtrOutput) Elem() GPUInstanceProfileOutput {
	return o.ApplyT(func(v *GPUInstanceProfile) GPUInstanceProfile {
		if v != nil {
			return *v
		}
		var ret GPUInstanceProfile
		return ret
	}).(GPUInstanceProfileOutput)
}

func (o GPUInstanceProfilePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GPUInstanceProfilePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GPUInstanceProfile) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GPUInstanceProfileInput interface {
	pulumi.Input

	ToGPUInstanceProfileOutput() GPUInstanceProfileOutput
	ToGPUInstanceProfileOutputWithContext(context.Context) GPUInstanceProfileOutput
}

var gpuinstanceProfilePtrType = reflect.TypeOf((**GPUInstanceProfile)(nil)).Elem()

type GPUInstanceProfilePtrInput interface {
	pulumi.Input

	ToGPUInstanceProfilePtrOutput() GPUInstanceProfilePtrOutput
	ToGPUInstanceProfilePtrOutputWithContext(context.Context) GPUInstanceProfilePtrOutput
}

type gpuinstanceProfilePtr string

func GPUInstanceProfilePtr(v string) GPUInstanceProfilePtrInput {
	return (*gpuinstanceProfilePtr)(&v)
}

func (*gpuinstanceProfilePtr) ElementType() reflect.Type {
	return gpuinstanceProfilePtrType
}

func (in *gpuinstanceProfilePtr) ToGPUInstanceProfilePtrOutput() GPUInstanceProfilePtrOutput {
	return pulumi.ToOutput(in).(GPUInstanceProfilePtrOutput)
}

func (in *gpuinstanceProfilePtr) ToGPUInstanceProfilePtrOutputWithContext(ctx context.Context) GPUInstanceProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GPUInstanceProfilePtrOutput)
}

type KubeletDiskType string

const (
	// Kubelet will use the OS disk for its data.
	KubeletDiskTypeOS = KubeletDiskType("OS")
	// Kubelet will use the temporary disk for its data.
	KubeletDiskTypeTemporary = KubeletDiskType("Temporary")
)

func (KubeletDiskType) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeletDiskType)(nil)).Elem()
}

func (e KubeletDiskType) ToKubeletDiskTypeOutput() KubeletDiskTypeOutput {
	return pulumi.ToOutput(e).(KubeletDiskTypeOutput)
}

func (e KubeletDiskType) ToKubeletDiskTypeOutputWithContext(ctx context.Context) KubeletDiskTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KubeletDiskTypeOutput)
}

func (e KubeletDiskType) ToKubeletDiskTypePtrOutput() KubeletDiskTypePtrOutput {
	return e.ToKubeletDiskTypePtrOutputWithContext(context.Background())
}

func (e KubeletDiskType) ToKubeletDiskTypePtrOutputWithContext(ctx context.Context) KubeletDiskTypePtrOutput {
	return KubeletDiskType(e).ToKubeletDiskTypeOutputWithContext(ctx).ToKubeletDiskTypePtrOutputWithContext(ctx)
}

func (e KubeletDiskType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KubeletDiskType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KubeletDiskType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KubeletDiskType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KubeletDiskTypeOutput struct{ *pulumi.OutputState }

func (KubeletDiskTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeletDiskType)(nil)).Elem()
}

func (o KubeletDiskTypeOutput) ToKubeletDiskTypeOutput() KubeletDiskTypeOutput {
	return o
}

func (o KubeletDiskTypeOutput) ToKubeletDiskTypeOutputWithContext(ctx context.Context) KubeletDiskTypeOutput {
	return o
}

func (o KubeletDiskTypeOutput) ToKubeletDiskTypePtrOutput() KubeletDiskTypePtrOutput {
	return o.ToKubeletDiskTypePtrOutputWithContext(context.Background())
}

func (o KubeletDiskTypeOutput) ToKubeletDiskTypePtrOutputWithContext(ctx context.Context) KubeletDiskTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubeletDiskType) *KubeletDiskType {
		return &v
	}).(KubeletDiskTypePtrOutput)
}

func (o KubeletDiskTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KubeletDiskTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KubeletDiskType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KubeletDiskTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KubeletDiskTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KubeletDiskType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KubeletDiskTypePtrOutput struct{ *pulumi.OutputState }

func (KubeletDiskTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeletDiskType)(nil)).Elem()
}

func (o KubeletDiskTypePtrOutput) ToKubeletDiskTypePtrOutput() KubeletDiskTypePtrOutput {
	return o
}

func (o KubeletDiskTypePtrOutput) ToKubeletDiskTypePtrOutputWithContext(ctx context.Context) KubeletDiskTypePtrOutput {
	return o
}

func (o KubeletDiskTypePtrOutput) Elem() KubeletDiskTypeOutput {
	return o.ApplyT(func(v *KubeletDiskType) KubeletDiskType {
		if v != nil {
			return *v
		}
		var ret KubeletDiskType
		return ret
	}).(KubeletDiskTypeOutput)
}

func (o KubeletDiskTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KubeletDiskTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KubeletDiskType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KubeletDiskTypeInput interface {
	pulumi.Input

	ToKubeletDiskTypeOutput() KubeletDiskTypeOutput
	ToKubeletDiskTypeOutputWithContext(context.Context) KubeletDiskTypeOutput
}

var kubeletDiskTypePtrType = reflect.TypeOf((**KubeletDiskType)(nil)).Elem()

type KubeletDiskTypePtrInput interface {
	pulumi.Input

	ToKubeletDiskTypePtrOutput() KubeletDiskTypePtrOutput
	ToKubeletDiskTypePtrOutputWithContext(context.Context) KubeletDiskTypePtrOutput
}

type kubeletDiskTypePtr string

func KubeletDiskTypePtr(v string) KubeletDiskTypePtrInput {
	return (*kubeletDiskTypePtr)(&v)
}

func (*kubeletDiskTypePtr) ElementType() reflect.Type {
	return kubeletDiskTypePtrType
}

func (in *kubeletDiskTypePtr) ToKubeletDiskTypePtrOutput() KubeletDiskTypePtrOutput {
	return pulumi.ToOutput(in).(KubeletDiskTypePtrOutput)
}

func (in *kubeletDiskTypePtr) ToKubeletDiskTypePtrOutputWithContext(ctx context.Context) KubeletDiskTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KubeletDiskTypePtrOutput)
}

type LicenseType string

const (
	// No additional licensing is applied.
	LicenseTypeNone = LicenseType("None")
	// Enables Azure Hybrid User Benefits for Windows VMs.
	LicenseType_Windows_Server = LicenseType("Windows_Server")
)

func (LicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*LicenseType)(nil)).Elem()
}

func (e LicenseType) ToLicenseTypeOutput() LicenseTypeOutput {
	return pulumi.ToOutput(e).(LicenseTypeOutput)
}

func (e LicenseType) ToLicenseTypeOutputWithContext(ctx context.Context) LicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LicenseTypeOutput)
}

func (e LicenseType) ToLicenseTypePtrOutput() LicenseTypePtrOutput {
	return e.ToLicenseTypePtrOutputWithContext(context.Background())
}

func (e LicenseType) ToLicenseTypePtrOutputWithContext(ctx context.Context) LicenseTypePtrOutput {
	return LicenseType(e).ToLicenseTypeOutputWithContext(ctx).ToLicenseTypePtrOutputWithContext(ctx)
}

func (e LicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LicenseTypeOutput struct{ *pulumi.OutputState }

func (LicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LicenseType)(nil)).Elem()
}

func (o LicenseTypeOutput) ToLicenseTypeOutput() LicenseTypeOutput {
	return o
}

func (o LicenseTypeOutput) ToLicenseTypeOutputWithContext(ctx context.Context) LicenseTypeOutput {
	return o
}

func (o LicenseTypeOutput) ToLicenseTypePtrOutput() LicenseTypePtrOutput {
	return o.ToLicenseTypePtrOutputWithContext(context.Background())
}

func (o LicenseTypeOutput) ToLicenseTypePtrOutputWithContext(ctx context.Context) LicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LicenseType) *LicenseType {
		return &v
	}).(LicenseTypePtrOutput)
}

func (o LicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LicenseTypePtrOutput struct{ *pulumi.OutputState }

func (LicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LicenseType)(nil)).Elem()
}

func (o LicenseTypePtrOutput) ToLicenseTypePtrOutput() LicenseTypePtrOutput {
	return o
}

func (o LicenseTypePtrOutput) ToLicenseTypePtrOutputWithContext(ctx context.Context) LicenseTypePtrOutput {
	return o
}

func (o LicenseTypePtrOutput) Elem() LicenseTypeOutput {
	return o.ApplyT(func(v *LicenseType) LicenseType {
		if v != nil {
			return *v
		}
		var ret LicenseType
		return ret
	}).(LicenseTypeOutput)
}

func (o LicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LicenseTypeInput interface {
	pulumi.Input

	ToLicenseTypeOutput() LicenseTypeOutput
	ToLicenseTypeOutputWithContext(context.Context) LicenseTypeOutput
}

var licenseTypePtrType = reflect.TypeOf((**LicenseType)(nil)).Elem()

type LicenseTypePtrInput interface {
	pulumi.Input

	ToLicenseTypePtrOutput() LicenseTypePtrOutput
	ToLicenseTypePtrOutputWithContext(context.Context) LicenseTypePtrOutput
}

type licenseTypePtr string

func LicenseTypePtr(v string) LicenseTypePtrInput {
	return (*licenseTypePtr)(&v)
}

func (*licenseTypePtr) ElementType() reflect.Type {
	return licenseTypePtrType
}

func (in *licenseTypePtr) ToLicenseTypePtrOutput() LicenseTypePtrOutput {
	return pulumi.ToOutput(in).(LicenseTypePtrOutput)
}

func (in *licenseTypePtr) ToLicenseTypePtrOutputWithContext(ctx context.Context) LicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LicenseTypePtrOutput)
}

type LoadBalancerSku string

const (
	// Use a a standard Load Balancer. This is the recommended Load Balancer SKU. For more information about on working with the load balancer in the managed cluster, see the [standard Load Balancer](https://docs.microsoft.com/azure/aks/load-balancer-standard) article.
	LoadBalancerSkuStandard = LoadBalancerSku("standard")
	// Use a basic Load Balancer with limited functionality.
	LoadBalancerSkuBasic = LoadBalancerSku("basic")
)

func (LoadBalancerSku) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSku)(nil)).Elem()
}

func (e LoadBalancerSku) ToLoadBalancerSkuOutput() LoadBalancerSkuOutput {
	return pulumi.ToOutput(e).(LoadBalancerSkuOutput)
}

func (e LoadBalancerSku) ToLoadBalancerSkuOutputWithContext(ctx context.Context) LoadBalancerSkuOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoadBalancerSkuOutput)
}

func (e LoadBalancerSku) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return e.ToLoadBalancerSkuPtrOutputWithContext(context.Background())
}

func (e LoadBalancerSku) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return LoadBalancerSku(e).ToLoadBalancerSkuOutputWithContext(ctx).ToLoadBalancerSkuPtrOutputWithContext(ctx)
}

func (e LoadBalancerSku) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerSku) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerSku) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoadBalancerSku) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoadBalancerSkuOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSku)(nil)).Elem()
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuOutput() LoadBalancerSkuOutput {
	return o
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuOutputWithContext(ctx context.Context) LoadBalancerSkuOutput {
	return o
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return o.ToLoadBalancerSkuPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuOutput) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerSku) *LoadBalancerSku {
		return &v
	}).(LoadBalancerSkuPtrOutput)
}

func (o LoadBalancerSkuOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoadBalancerSkuOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerSku) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoadBalancerSkuOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerSku) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerSkuPtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerSku)(nil)).Elem()
}

func (o LoadBalancerSkuPtrOutput) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return o
}

func (o LoadBalancerSkuPtrOutput) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return o
}

func (o LoadBalancerSkuPtrOutput) Elem() LoadBalancerSkuOutput {
	return o.ApplyT(func(v *LoadBalancerSku) LoadBalancerSku {
		if v != nil {
			return *v
		}
		var ret LoadBalancerSku
		return ret
	}).(LoadBalancerSkuOutput)
}

func (o LoadBalancerSkuPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoadBalancerSku) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LoadBalancerSkuInput interface {
	pulumi.Input

	ToLoadBalancerSkuOutput() LoadBalancerSkuOutput
	ToLoadBalancerSkuOutputWithContext(context.Context) LoadBalancerSkuOutput
}

var loadBalancerSkuPtrType = reflect.TypeOf((**LoadBalancerSku)(nil)).Elem()

type LoadBalancerSkuPtrInput interface {
	pulumi.Input

	ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput
	ToLoadBalancerSkuPtrOutputWithContext(context.Context) LoadBalancerSkuPtrOutput
}

type loadBalancerSkuPtr string

func LoadBalancerSkuPtr(v string) LoadBalancerSkuPtrInput {
	return (*loadBalancerSkuPtr)(&v)
}

func (*loadBalancerSkuPtr) ElementType() reflect.Type {
	return loadBalancerSkuPtrType
}

func (in *loadBalancerSkuPtr) ToLoadBalancerSkuPtrOutput() LoadBalancerSkuPtrOutput {
	return pulumi.ToOutput(in).(LoadBalancerSkuPtrOutput)
}

func (in *loadBalancerSkuPtr) ToLoadBalancerSkuPtrOutputWithContext(ctx context.Context) LoadBalancerSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoadBalancerSkuPtrOutput)
}

type ManagedClusterSKUName string

const (
	ManagedClusterSKUNameBasic = ManagedClusterSKUName("Basic")
)

func (ManagedClusterSKUName) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterSKUName)(nil)).Elem()
}

func (e ManagedClusterSKUName) ToManagedClusterSKUNameOutput() ManagedClusterSKUNameOutput {
	return pulumi.ToOutput(e).(ManagedClusterSKUNameOutput)
}

func (e ManagedClusterSKUName) ToManagedClusterSKUNameOutputWithContext(ctx context.Context) ManagedClusterSKUNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedClusterSKUNameOutput)
}

func (e ManagedClusterSKUName) ToManagedClusterSKUNamePtrOutput() ManagedClusterSKUNamePtrOutput {
	return e.ToManagedClusterSKUNamePtrOutputWithContext(context.Background())
}

func (e ManagedClusterSKUName) ToManagedClusterSKUNamePtrOutputWithContext(ctx context.Context) ManagedClusterSKUNamePtrOutput {
	return ManagedClusterSKUName(e).ToManagedClusterSKUNameOutputWithContext(ctx).ToManagedClusterSKUNamePtrOutputWithContext(ctx)
}

func (e ManagedClusterSKUName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedClusterSKUName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedClusterSKUName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedClusterSKUName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedClusterSKUNameOutput struct{ *pulumi.OutputState }

func (ManagedClusterSKUNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterSKUName)(nil)).Elem()
}

func (o ManagedClusterSKUNameOutput) ToManagedClusterSKUNameOutput() ManagedClusterSKUNameOutput {
	return o
}

func (o ManagedClusterSKUNameOutput) ToManagedClusterSKUNameOutputWithContext(ctx context.Context) ManagedClusterSKUNameOutput {
	return o
}

func (o ManagedClusterSKUNameOutput) ToManagedClusterSKUNamePtrOutput() ManagedClusterSKUNamePtrOutput {
	return o.ToManagedClusterSKUNamePtrOutputWithContext(context.Background())
}

func (o ManagedClusterSKUNameOutput) ToManagedClusterSKUNamePtrOutputWithContext(ctx context.Context) ManagedClusterSKUNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterSKUName) *ManagedClusterSKUName {
		return &v
	}).(ManagedClusterSKUNamePtrOutput)
}

func (o ManagedClusterSKUNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedClusterSKUNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedClusterSKUName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedClusterSKUNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedClusterSKUNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedClusterSKUName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterSKUNamePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterSKUNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSKUName)(nil)).Elem()
}

func (o ManagedClusterSKUNamePtrOutput) ToManagedClusterSKUNamePtrOutput() ManagedClusterSKUNamePtrOutput {
	return o
}

func (o ManagedClusterSKUNamePtrOutput) ToManagedClusterSKUNamePtrOutputWithContext(ctx context.Context) ManagedClusterSKUNamePtrOutput {
	return o
}

func (o ManagedClusterSKUNamePtrOutput) Elem() ManagedClusterSKUNameOutput {
	return o.ApplyT(func(v *ManagedClusterSKUName) ManagedClusterSKUName {
		if v != nil {
			return *v
		}
		var ret ManagedClusterSKUName
		return ret
	}).(ManagedClusterSKUNameOutput)
}

func (o ManagedClusterSKUNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedClusterSKUNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedClusterSKUName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedClusterSKUNameInput interface {
	pulumi.Input

	ToManagedClusterSKUNameOutput() ManagedClusterSKUNameOutput
	ToManagedClusterSKUNameOutputWithContext(context.Context) ManagedClusterSKUNameOutput
}

var managedClusterSKUNamePtrType = reflect.TypeOf((**ManagedClusterSKUName)(nil)).Elem()

type ManagedClusterSKUNamePtrInput interface {
	pulumi.Input

	ToManagedClusterSKUNamePtrOutput() ManagedClusterSKUNamePtrOutput
	ToManagedClusterSKUNamePtrOutputWithContext(context.Context) ManagedClusterSKUNamePtrOutput
}

type managedClusterSKUNamePtr string

func ManagedClusterSKUNamePtr(v string) ManagedClusterSKUNamePtrInput {
	return (*managedClusterSKUNamePtr)(&v)
}

func (*managedClusterSKUNamePtr) ElementType() reflect.Type {
	return managedClusterSKUNamePtrType
}

func (in *managedClusterSKUNamePtr) ToManagedClusterSKUNamePtrOutput() ManagedClusterSKUNamePtrOutput {
	return pulumi.ToOutput(in).(ManagedClusterSKUNamePtrOutput)
}

func (in *managedClusterSKUNamePtr) ToManagedClusterSKUNamePtrOutputWithContext(ctx context.Context) ManagedClusterSKUNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedClusterSKUNamePtrOutput)
}

type ManagedClusterSKUTier string

const (
	// Guarantees 99.95% availability of the Kubernetes API server endpoint for clusters that use Availability Zones and 99.9% of availability for clusters that don't use Availability Zones.
	ManagedClusterSKUTierPaid = ManagedClusterSKUTier("Paid")
	// No guaranteed SLA, no additional charges. Free tier clusters have an SLO of 99.5%.
	ManagedClusterSKUTierFree = ManagedClusterSKUTier("Free")
)

func (ManagedClusterSKUTier) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterSKUTier)(nil)).Elem()
}

func (e ManagedClusterSKUTier) ToManagedClusterSKUTierOutput() ManagedClusterSKUTierOutput {
	return pulumi.ToOutput(e).(ManagedClusterSKUTierOutput)
}

func (e ManagedClusterSKUTier) ToManagedClusterSKUTierOutputWithContext(ctx context.Context) ManagedClusterSKUTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedClusterSKUTierOutput)
}

func (e ManagedClusterSKUTier) ToManagedClusterSKUTierPtrOutput() ManagedClusterSKUTierPtrOutput {
	return e.ToManagedClusterSKUTierPtrOutputWithContext(context.Background())
}

func (e ManagedClusterSKUTier) ToManagedClusterSKUTierPtrOutputWithContext(ctx context.Context) ManagedClusterSKUTierPtrOutput {
	return ManagedClusterSKUTier(e).ToManagedClusterSKUTierOutputWithContext(ctx).ToManagedClusterSKUTierPtrOutputWithContext(ctx)
}

func (e ManagedClusterSKUTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedClusterSKUTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedClusterSKUTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedClusterSKUTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedClusterSKUTierOutput struct{ *pulumi.OutputState }

func (ManagedClusterSKUTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterSKUTier)(nil)).Elem()
}

func (o ManagedClusterSKUTierOutput) ToManagedClusterSKUTierOutput() ManagedClusterSKUTierOutput {
	return o
}

func (o ManagedClusterSKUTierOutput) ToManagedClusterSKUTierOutputWithContext(ctx context.Context) ManagedClusterSKUTierOutput {
	return o
}

func (o ManagedClusterSKUTierOutput) ToManagedClusterSKUTierPtrOutput() ManagedClusterSKUTierPtrOutput {
	return o.ToManagedClusterSKUTierPtrOutputWithContext(context.Background())
}

func (o ManagedClusterSKUTierOutput) ToManagedClusterSKUTierPtrOutputWithContext(ctx context.Context) ManagedClusterSKUTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterSKUTier) *ManagedClusterSKUTier {
		return &v
	}).(ManagedClusterSKUTierPtrOutput)
}

func (o ManagedClusterSKUTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedClusterSKUTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedClusterSKUTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedClusterSKUTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedClusterSKUTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedClusterSKUTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterSKUTierPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterSKUTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterSKUTier)(nil)).Elem()
}

func (o ManagedClusterSKUTierPtrOutput) ToManagedClusterSKUTierPtrOutput() ManagedClusterSKUTierPtrOutput {
	return o
}

func (o ManagedClusterSKUTierPtrOutput) ToManagedClusterSKUTierPtrOutputWithContext(ctx context.Context) ManagedClusterSKUTierPtrOutput {
	return o
}

func (o ManagedClusterSKUTierPtrOutput) Elem() ManagedClusterSKUTierOutput {
	return o.ApplyT(func(v *ManagedClusterSKUTier) ManagedClusterSKUTier {
		if v != nil {
			return *v
		}
		var ret ManagedClusterSKUTier
		return ret
	}).(ManagedClusterSKUTierOutput)
}

func (o ManagedClusterSKUTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedClusterSKUTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedClusterSKUTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedClusterSKUTierInput interface {
	pulumi.Input

	ToManagedClusterSKUTierOutput() ManagedClusterSKUTierOutput
	ToManagedClusterSKUTierOutputWithContext(context.Context) ManagedClusterSKUTierOutput
}

var managedClusterSKUTierPtrType = reflect.TypeOf((**ManagedClusterSKUTier)(nil)).Elem()

type ManagedClusterSKUTierPtrInput interface {
	pulumi.Input

	ToManagedClusterSKUTierPtrOutput() ManagedClusterSKUTierPtrOutput
	ToManagedClusterSKUTierPtrOutputWithContext(context.Context) ManagedClusterSKUTierPtrOutput
}

type managedClusterSKUTierPtr string

func ManagedClusterSKUTierPtr(v string) ManagedClusterSKUTierPtrInput {
	return (*managedClusterSKUTierPtr)(&v)
}

func (*managedClusterSKUTierPtr) ElementType() reflect.Type {
	return managedClusterSKUTierPtrType
}

func (in *managedClusterSKUTierPtr) ToManagedClusterSKUTierPtrOutput() ManagedClusterSKUTierPtrOutput {
	return pulumi.ToOutput(in).(ManagedClusterSKUTierPtrOutput)
}

func (in *managedClusterSKUTierPtr) ToManagedClusterSKUTierPtrOutputWithContext(ctx context.Context) ManagedClusterSKUTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedClusterSKUTierPtrOutput)
}

type NetworkMode string

const (
	// No bridge is created. Intra-VM Pod to Pod communication is through IP routes created by Azure CNI. See [Transparent Mode](https://docs.microsoft.com/azure/aks/faq#transparent-mode) for more information.
	NetworkModeTransparent = NetworkMode("transparent")
	// This is no longer supported
	NetworkModeBridge = NetworkMode("bridge")
)

func (NetworkMode) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkMode)(nil)).Elem()
}

func (e NetworkMode) ToNetworkModeOutput() NetworkModeOutput {
	return pulumi.ToOutput(e).(NetworkModeOutput)
}

func (e NetworkMode) ToNetworkModeOutputWithContext(ctx context.Context) NetworkModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkModeOutput)
}

func (e NetworkMode) ToNetworkModePtrOutput() NetworkModePtrOutput {
	return e.ToNetworkModePtrOutputWithContext(context.Background())
}

func (e NetworkMode) ToNetworkModePtrOutputWithContext(ctx context.Context) NetworkModePtrOutput {
	return NetworkMode(e).ToNetworkModeOutputWithContext(ctx).ToNetworkModePtrOutputWithContext(ctx)
}

func (e NetworkMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkModeOutput struct{ *pulumi.OutputState }

func (NetworkModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkMode)(nil)).Elem()
}

func (o NetworkModeOutput) ToNetworkModeOutput() NetworkModeOutput {
	return o
}

func (o NetworkModeOutput) ToNetworkModeOutputWithContext(ctx context.Context) NetworkModeOutput {
	return o
}

func (o NetworkModeOutput) ToNetworkModePtrOutput() NetworkModePtrOutput {
	return o.ToNetworkModePtrOutputWithContext(context.Background())
}

func (o NetworkModeOutput) ToNetworkModePtrOutputWithContext(ctx context.Context) NetworkModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkMode) *NetworkMode {
		return &v
	}).(NetworkModePtrOutput)
}

func (o NetworkModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkModePtrOutput struct{ *pulumi.OutputState }

func (NetworkModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkMode)(nil)).Elem()
}

func (o NetworkModePtrOutput) ToNetworkModePtrOutput() NetworkModePtrOutput {
	return o
}

func (o NetworkModePtrOutput) ToNetworkModePtrOutputWithContext(ctx context.Context) NetworkModePtrOutput {
	return o
}

func (o NetworkModePtrOutput) Elem() NetworkModeOutput {
	return o.ApplyT(func(v *NetworkMode) NetworkMode {
		if v != nil {
			return *v
		}
		var ret NetworkMode
		return ret
	}).(NetworkModeOutput)
}

func (o NetworkModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkModeInput interface {
	pulumi.Input

	ToNetworkModeOutput() NetworkModeOutput
	ToNetworkModeOutputWithContext(context.Context) NetworkModeOutput
}

var networkModePtrType = reflect.TypeOf((**NetworkMode)(nil)).Elem()

type NetworkModePtrInput interface {
	pulumi.Input

	ToNetworkModePtrOutput() NetworkModePtrOutput
	ToNetworkModePtrOutputWithContext(context.Context) NetworkModePtrOutput
}

type networkModePtr string

func NetworkModePtr(v string) NetworkModePtrInput {
	return (*networkModePtr)(&v)
}

func (*networkModePtr) ElementType() reflect.Type {
	return networkModePtrType
}

func (in *networkModePtr) ToNetworkModePtrOutput() NetworkModePtrOutput {
	return pulumi.ToOutput(in).(NetworkModePtrOutput)
}

func (in *networkModePtr) ToNetworkModePtrOutputWithContext(ctx context.Context) NetworkModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkModePtrOutput)
}

type NetworkPlugin string

const (
	// Use the Azure CNI network plugin. See [Azure CNI (advanced) networking](https://docs.microsoft.com/azure/aks/concepts-network#azure-cni-advanced-networking) for more information.
	NetworkPluginAzure = NetworkPlugin("azure")
	// Use the Kubenet network plugin. See [Kubenet (basic) networking](https://docs.microsoft.com/azure/aks/concepts-network#kubenet-basic-networking) for more information.
	NetworkPluginKubenet = NetworkPlugin("kubenet")
)

func (NetworkPlugin) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPlugin)(nil)).Elem()
}

func (e NetworkPlugin) ToNetworkPluginOutput() NetworkPluginOutput {
	return pulumi.ToOutput(e).(NetworkPluginOutput)
}

func (e NetworkPlugin) ToNetworkPluginOutputWithContext(ctx context.Context) NetworkPluginOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkPluginOutput)
}

func (e NetworkPlugin) ToNetworkPluginPtrOutput() NetworkPluginPtrOutput {
	return e.ToNetworkPluginPtrOutputWithContext(context.Background())
}

func (e NetworkPlugin) ToNetworkPluginPtrOutputWithContext(ctx context.Context) NetworkPluginPtrOutput {
	return NetworkPlugin(e).ToNetworkPluginOutputWithContext(ctx).ToNetworkPluginPtrOutputWithContext(ctx)
}

func (e NetworkPlugin) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkPlugin) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkPlugin) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkPlugin) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkPluginOutput struct{ *pulumi.OutputState }

func (NetworkPluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPlugin)(nil)).Elem()
}

func (o NetworkPluginOutput) ToNetworkPluginOutput() NetworkPluginOutput {
	return o
}

func (o NetworkPluginOutput) ToNetworkPluginOutputWithContext(ctx context.Context) NetworkPluginOutput {
	return o
}

func (o NetworkPluginOutput) ToNetworkPluginPtrOutput() NetworkPluginPtrOutput {
	return o.ToNetworkPluginPtrOutputWithContext(context.Background())
}

func (o NetworkPluginOutput) ToNetworkPluginPtrOutputWithContext(ctx context.Context) NetworkPluginPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkPlugin) *NetworkPlugin {
		return &v
	}).(NetworkPluginPtrOutput)
}

func (o NetworkPluginOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkPluginOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkPlugin) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkPluginOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkPluginOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkPlugin) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkPluginPtrOutput struct{ *pulumi.OutputState }

func (NetworkPluginPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPlugin)(nil)).Elem()
}

func (o NetworkPluginPtrOutput) ToNetworkPluginPtrOutput() NetworkPluginPtrOutput {
	return o
}

func (o NetworkPluginPtrOutput) ToNetworkPluginPtrOutputWithContext(ctx context.Context) NetworkPluginPtrOutput {
	return o
}

func (o NetworkPluginPtrOutput) Elem() NetworkPluginOutput {
	return o.ApplyT(func(v *NetworkPlugin) NetworkPlugin {
		if v != nil {
			return *v
		}
		var ret NetworkPlugin
		return ret
	}).(NetworkPluginOutput)
}

func (o NetworkPluginPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkPluginPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkPlugin) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkPluginInput interface {
	pulumi.Input

	ToNetworkPluginOutput() NetworkPluginOutput
	ToNetworkPluginOutputWithContext(context.Context) NetworkPluginOutput
}

var networkPluginPtrType = reflect.TypeOf((**NetworkPlugin)(nil)).Elem()

type NetworkPluginPtrInput interface {
	pulumi.Input

	ToNetworkPluginPtrOutput() NetworkPluginPtrOutput
	ToNetworkPluginPtrOutputWithContext(context.Context) NetworkPluginPtrOutput
}

type networkPluginPtr string

func NetworkPluginPtr(v string) NetworkPluginPtrInput {
	return (*networkPluginPtr)(&v)
}

func (*networkPluginPtr) ElementType() reflect.Type {
	return networkPluginPtrType
}

func (in *networkPluginPtr) ToNetworkPluginPtrOutput() NetworkPluginPtrOutput {
	return pulumi.ToOutput(in).(NetworkPluginPtrOutput)
}

func (in *networkPluginPtr) ToNetworkPluginPtrOutputWithContext(ctx context.Context) NetworkPluginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkPluginPtrOutput)
}

type NetworkPolicy string

const (
	// Use Calico network policies. See [differences between Azure and Calico policies](https://docs.microsoft.com/azure/aks/use-network-policies#differences-between-azure-and-calico-policies-and-their-capabilities) for more information.
	NetworkPolicyCalico = NetworkPolicy("calico")
	// Use Azure network policies. See [differences between Azure and Calico policies](https://docs.microsoft.com/azure/aks/use-network-policies#differences-between-azure-and-calico-policies-and-their-capabilities) for more information.
	NetworkPolicyAzure = NetworkPolicy("azure")
)

func (NetworkPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPolicy)(nil)).Elem()
}

func (e NetworkPolicy) ToNetworkPolicyOutput() NetworkPolicyOutput {
	return pulumi.ToOutput(e).(NetworkPolicyOutput)
}

func (e NetworkPolicy) ToNetworkPolicyOutputWithContext(ctx context.Context) NetworkPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkPolicyOutput)
}

func (e NetworkPolicy) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return e.ToNetworkPolicyPtrOutputWithContext(context.Background())
}

func (e NetworkPolicy) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return NetworkPolicy(e).ToNetworkPolicyOutputWithContext(ctx).ToNetworkPolicyPtrOutputWithContext(ctx)
}

func (e NetworkPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkPolicyOutput struct{ *pulumi.OutputState }

func (NetworkPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkPolicy)(nil)).Elem()
}

func (o NetworkPolicyOutput) ToNetworkPolicyOutput() NetworkPolicyOutput {
	return o
}

func (o NetworkPolicyOutput) ToNetworkPolicyOutputWithContext(ctx context.Context) NetworkPolicyOutput {
	return o
}

func (o NetworkPolicyOutput) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return o.ToNetworkPolicyPtrOutputWithContext(context.Background())
}

func (o NetworkPolicyOutput) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkPolicy) *NetworkPolicy {
		return &v
	}).(NetworkPolicyPtrOutput)
}

func (o NetworkPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkPolicyPtrOutput struct{ *pulumi.OutputState }

func (NetworkPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPolicy)(nil)).Elem()
}

func (o NetworkPolicyPtrOutput) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return o
}

func (o NetworkPolicyPtrOutput) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return o
}

func (o NetworkPolicyPtrOutput) Elem() NetworkPolicyOutput {
	return o.ApplyT(func(v *NetworkPolicy) NetworkPolicy {
		if v != nil {
			return *v
		}
		var ret NetworkPolicy
		return ret
	}).(NetworkPolicyOutput)
}

func (o NetworkPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkPolicyInput interface {
	pulumi.Input

	ToNetworkPolicyOutput() NetworkPolicyOutput
	ToNetworkPolicyOutputWithContext(context.Context) NetworkPolicyOutput
}

var networkPolicyPtrType = reflect.TypeOf((**NetworkPolicy)(nil)).Elem()

type NetworkPolicyPtrInput interface {
	pulumi.Input

	ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput
	ToNetworkPolicyPtrOutputWithContext(context.Context) NetworkPolicyPtrOutput
}

type networkPolicyPtr string

func NetworkPolicyPtr(v string) NetworkPolicyPtrInput {
	return (*networkPolicyPtr)(&v)
}

func (*networkPolicyPtr) ElementType() reflect.Type {
	return networkPolicyPtrType
}

func (in *networkPolicyPtr) ToNetworkPolicyPtrOutput() NetworkPolicyPtrOutput {
	return pulumi.ToOutput(in).(NetworkPolicyPtrOutput)
}

func (in *networkPolicyPtr) ToNetworkPolicyPtrOutputWithContext(ctx context.Context) NetworkPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkPolicyPtrOutput)
}

type OSDiskType string

const (
	// Azure replicates the operating system disk for a virtual machine to Azure storage to avoid data loss should the VM need to be relocated to another host. Since containers aren't designed to have local state persisted, this behavior offers limited value while providing some drawbacks, including slower node provisioning and higher read/write latency.
	OSDiskTypeManaged = OSDiskType("Managed")
	// Ephemeral OS disks are stored only on the host machine, just like a temporary disk. This provides lower read/write latency, along with faster node scaling and cluster upgrades.
	OSDiskTypeEphemeral = OSDiskType("Ephemeral")
)

func (OSDiskType) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskType)(nil)).Elem()
}

func (e OSDiskType) ToOSDiskTypeOutput() OSDiskTypeOutput {
	return pulumi.ToOutput(e).(OSDiskTypeOutput)
}

func (e OSDiskType) ToOSDiskTypeOutputWithContext(ctx context.Context) OSDiskTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OSDiskTypeOutput)
}

func (e OSDiskType) ToOSDiskTypePtrOutput() OSDiskTypePtrOutput {
	return e.ToOSDiskTypePtrOutputWithContext(context.Background())
}

func (e OSDiskType) ToOSDiskTypePtrOutputWithContext(ctx context.Context) OSDiskTypePtrOutput {
	return OSDiskType(e).ToOSDiskTypeOutputWithContext(ctx).ToOSDiskTypePtrOutputWithContext(ctx)
}

func (e OSDiskType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSDiskType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSDiskType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OSDiskType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OSDiskTypeOutput struct{ *pulumi.OutputState }

func (OSDiskTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskType)(nil)).Elem()
}

func (o OSDiskTypeOutput) ToOSDiskTypeOutput() OSDiskTypeOutput {
	return o
}

func (o OSDiskTypeOutput) ToOSDiskTypeOutputWithContext(ctx context.Context) OSDiskTypeOutput {
	return o
}

func (o OSDiskTypeOutput) ToOSDiskTypePtrOutput() OSDiskTypePtrOutput {
	return o.ToOSDiskTypePtrOutputWithContext(context.Background())
}

func (o OSDiskTypeOutput) ToOSDiskTypePtrOutputWithContext(ctx context.Context) OSDiskTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSDiskType) *OSDiskType {
		return &v
	}).(OSDiskTypePtrOutput)
}

func (o OSDiskTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OSDiskTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSDiskType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OSDiskTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSDiskTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSDiskType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OSDiskTypePtrOutput struct{ *pulumi.OutputState }

func (OSDiskTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskType)(nil)).Elem()
}

func (o OSDiskTypePtrOutput) ToOSDiskTypePtrOutput() OSDiskTypePtrOutput {
	return o
}

func (o OSDiskTypePtrOutput) ToOSDiskTypePtrOutputWithContext(ctx context.Context) OSDiskTypePtrOutput {
	return o
}

func (o OSDiskTypePtrOutput) Elem() OSDiskTypeOutput {
	return o.ApplyT(func(v *OSDiskType) OSDiskType {
		if v != nil {
			return *v
		}
		var ret OSDiskType
		return ret
	}).(OSDiskTypeOutput)
}

func (o OSDiskTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSDiskTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OSDiskType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OSDiskTypeInput interface {
	pulumi.Input

	ToOSDiskTypeOutput() OSDiskTypeOutput
	ToOSDiskTypeOutputWithContext(context.Context) OSDiskTypeOutput
}

var osdiskTypePtrType = reflect.TypeOf((**OSDiskType)(nil)).Elem()

type OSDiskTypePtrInput interface {
	pulumi.Input

	ToOSDiskTypePtrOutput() OSDiskTypePtrOutput
	ToOSDiskTypePtrOutputWithContext(context.Context) OSDiskTypePtrOutput
}

type osdiskTypePtr string

func OSDiskTypePtr(v string) OSDiskTypePtrInput {
	return (*osdiskTypePtr)(&v)
}

func (*osdiskTypePtr) ElementType() reflect.Type {
	return osdiskTypePtrType
}

func (in *osdiskTypePtr) ToOSDiskTypePtrOutput() OSDiskTypePtrOutput {
	return pulumi.ToOutput(in).(OSDiskTypePtrOutput)
}

func (in *osdiskTypePtr) ToOSDiskTypePtrOutputWithContext(ctx context.Context) OSDiskTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OSDiskTypePtrOutput)
}

type OSSKU string

const (
	OSSKUUbuntu     = OSSKU("Ubuntu")
	OSSKUCBLMariner = OSSKU("CBLMariner")
)

func (OSSKU) ElementType() reflect.Type {
	return reflect.TypeOf((*OSSKU)(nil)).Elem()
}

func (e OSSKU) ToOSSKUOutput() OSSKUOutput {
	return pulumi.ToOutput(e).(OSSKUOutput)
}

func (e OSSKU) ToOSSKUOutputWithContext(ctx context.Context) OSSKUOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OSSKUOutput)
}

func (e OSSKU) ToOSSKUPtrOutput() OSSKUPtrOutput {
	return e.ToOSSKUPtrOutputWithContext(context.Background())
}

func (e OSSKU) ToOSSKUPtrOutputWithContext(ctx context.Context) OSSKUPtrOutput {
	return OSSKU(e).ToOSSKUOutputWithContext(ctx).ToOSSKUPtrOutputWithContext(ctx)
}

func (e OSSKU) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSSKU) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSSKU) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OSSKU) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OSSKUOutput struct{ *pulumi.OutputState }

func (OSSKUOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSSKU)(nil)).Elem()
}

func (o OSSKUOutput) ToOSSKUOutput() OSSKUOutput {
	return o
}

func (o OSSKUOutput) ToOSSKUOutputWithContext(ctx context.Context) OSSKUOutput {
	return o
}

func (o OSSKUOutput) ToOSSKUPtrOutput() OSSKUPtrOutput {
	return o.ToOSSKUPtrOutputWithContext(context.Background())
}

func (o OSSKUOutput) ToOSSKUPtrOutputWithContext(ctx context.Context) OSSKUPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSSKU) *OSSKU {
		return &v
	}).(OSSKUPtrOutput)
}

func (o OSSKUOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OSSKUOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSSKU) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OSSKUOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSSKUOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSSKU) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OSSKUPtrOutput struct{ *pulumi.OutputState }

func (OSSKUPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSSKU)(nil)).Elem()
}

func (o OSSKUPtrOutput) ToOSSKUPtrOutput() OSSKUPtrOutput {
	return o
}

func (o OSSKUPtrOutput) ToOSSKUPtrOutputWithContext(ctx context.Context) OSSKUPtrOutput {
	return o
}

func (o OSSKUPtrOutput) Elem() OSSKUOutput {
	return o.ApplyT(func(v *OSSKU) OSSKU {
		if v != nil {
			return *v
		}
		var ret OSSKU
		return ret
	}).(OSSKUOutput)
}

func (o OSSKUPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSSKUPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OSSKU) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OSSKUInput interface {
	pulumi.Input

	ToOSSKUOutput() OSSKUOutput
	ToOSSKUOutputWithContext(context.Context) OSSKUOutput
}

var osskuPtrType = reflect.TypeOf((**OSSKU)(nil)).Elem()

type OSSKUPtrInput interface {
	pulumi.Input

	ToOSSKUPtrOutput() OSSKUPtrOutput
	ToOSSKUPtrOutputWithContext(context.Context) OSSKUPtrOutput
}

type osskuPtr string

func OSSKUPtr(v string) OSSKUPtrInput {
	return (*osskuPtr)(&v)
}

func (*osskuPtr) ElementType() reflect.Type {
	return osskuPtrType
}

func (in *osskuPtr) ToOSSKUPtrOutput() OSSKUPtrOutput {
	return pulumi.ToOutput(in).(OSSKUPtrOutput)
}

func (in *osskuPtr) ToOSSKUPtrOutputWithContext(ctx context.Context) OSSKUPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OSSKUPtrOutput)
}

type OSType string

const (
	// Use Linux.
	OSTypeLinux = OSType("Linux")
	// Use Windows.
	OSTypeWindows = OSType("Windows")
)

func (OSType) ElementType() reflect.Type {
	return reflect.TypeOf((*OSType)(nil)).Elem()
}

func (e OSType) ToOSTypeOutput() OSTypeOutput {
	return pulumi.ToOutput(e).(OSTypeOutput)
}

func (e OSType) ToOSTypeOutputWithContext(ctx context.Context) OSTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OSTypeOutput)
}

func (e OSType) ToOSTypePtrOutput() OSTypePtrOutput {
	return e.ToOSTypePtrOutputWithContext(context.Background())
}

func (e OSType) ToOSTypePtrOutputWithContext(ctx context.Context) OSTypePtrOutput {
	return OSType(e).ToOSTypeOutputWithContext(ctx).ToOSTypePtrOutputWithContext(ctx)
}

func (e OSType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OSType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OSType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OSTypeOutput struct{ *pulumi.OutputState }

func (OSTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSType)(nil)).Elem()
}

func (o OSTypeOutput) ToOSTypeOutput() OSTypeOutput {
	return o
}

func (o OSTypeOutput) ToOSTypeOutputWithContext(ctx context.Context) OSTypeOutput {
	return o
}

func (o OSTypeOutput) ToOSTypePtrOutput() OSTypePtrOutput {
	return o.ToOSTypePtrOutputWithContext(context.Background())
}

func (o OSTypeOutput) ToOSTypePtrOutputWithContext(ctx context.Context) OSTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSType) *OSType {
		return &v
	}).(OSTypePtrOutput)
}

func (o OSTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OSTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OSTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OSType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OSTypePtrOutput struct{ *pulumi.OutputState }

func (OSTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSType)(nil)).Elem()
}

func (o OSTypePtrOutput) ToOSTypePtrOutput() OSTypePtrOutput {
	return o
}

func (o OSTypePtrOutput) ToOSTypePtrOutputWithContext(ctx context.Context) OSTypePtrOutput {
	return o
}

func (o OSTypePtrOutput) Elem() OSTypeOutput {
	return o.ApplyT(func(v *OSType) OSType {
		if v != nil {
			return *v
		}
		var ret OSType
		return ret
	}).(OSTypeOutput)
}

func (o OSTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OSType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OSTypeInput interface {
	pulumi.Input

	ToOSTypeOutput() OSTypeOutput
	ToOSTypeOutputWithContext(context.Context) OSTypeOutput
}

var ostypePtrType = reflect.TypeOf((**OSType)(nil)).Elem()

type OSTypePtrInput interface {
	pulumi.Input

	ToOSTypePtrOutput() OSTypePtrOutput
	ToOSTypePtrOutputWithContext(context.Context) OSTypePtrOutput
}

type ostypePtr string

func OSTypePtr(v string) OSTypePtrInput {
	return (*ostypePtr)(&v)
}

func (*ostypePtr) ElementType() reflect.Type {
	return ostypePtrType
}

func (in *ostypePtr) ToOSTypePtrOutput() OSTypePtrOutput {
	return pulumi.ToOutput(in).(OSTypePtrOutput)
}

func (in *ostypePtr) ToOSTypePtrOutputWithContext(ctx context.Context) OSTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OSTypePtrOutput)
}

type OutboundType string

const (
	// The load balancer is used for egress through an AKS assigned public IP. This supports Kubernetes services of type 'loadBalancer'. For more information see [outbound type loadbalancer](https://docs.microsoft.com/azure/aks/egress-outboundtype#outbound-type-of-loadbalancer).
	OutboundTypeLoadBalancer = OutboundType("loadBalancer")
	// Egress paths must be defined by the user. This is an advanced scenario and requires proper network configuration. For more information see [outbound type userDefinedRouting](https://docs.microsoft.com/azure/aks/egress-outboundtype#outbound-type-of-userdefinedrouting).
	OutboundTypeUserDefinedRouting = OutboundType("userDefinedRouting")
)

func (OutboundType) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundType)(nil)).Elem()
}

func (e OutboundType) ToOutboundTypeOutput() OutboundTypeOutput {
	return pulumi.ToOutput(e).(OutboundTypeOutput)
}

func (e OutboundType) ToOutboundTypeOutputWithContext(ctx context.Context) OutboundTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OutboundTypeOutput)
}

func (e OutboundType) ToOutboundTypePtrOutput() OutboundTypePtrOutput {
	return e.ToOutboundTypePtrOutputWithContext(context.Background())
}

func (e OutboundType) ToOutboundTypePtrOutputWithContext(ctx context.Context) OutboundTypePtrOutput {
	return OutboundType(e).ToOutboundTypeOutputWithContext(ctx).ToOutboundTypePtrOutputWithContext(ctx)
}

func (e OutboundType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutboundType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutboundType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OutboundType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OutboundTypeOutput struct{ *pulumi.OutputState }

func (OutboundTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundType)(nil)).Elem()
}

func (o OutboundTypeOutput) ToOutboundTypeOutput() OutboundTypeOutput {
	return o
}

func (o OutboundTypeOutput) ToOutboundTypeOutputWithContext(ctx context.Context) OutboundTypeOutput {
	return o
}

func (o OutboundTypeOutput) ToOutboundTypePtrOutput() OutboundTypePtrOutput {
	return o.ToOutboundTypePtrOutputWithContext(context.Background())
}

func (o OutboundTypeOutput) ToOutboundTypePtrOutputWithContext(ctx context.Context) OutboundTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutboundType) *OutboundType {
		return &v
	}).(OutboundTypePtrOutput)
}

func (o OutboundTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OutboundTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutboundType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OutboundTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutboundTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutboundType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OutboundTypePtrOutput struct{ *pulumi.OutputState }

func (OutboundTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundType)(nil)).Elem()
}

func (o OutboundTypePtrOutput) ToOutboundTypePtrOutput() OutboundTypePtrOutput {
	return o
}

func (o OutboundTypePtrOutput) ToOutboundTypePtrOutputWithContext(ctx context.Context) OutboundTypePtrOutput {
	return o
}

func (o OutboundTypePtrOutput) Elem() OutboundTypeOutput {
	return o.ApplyT(func(v *OutboundType) OutboundType {
		if v != nil {
			return *v
		}
		var ret OutboundType
		return ret
	}).(OutboundTypeOutput)
}

func (o OutboundTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutboundTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OutboundType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OutboundTypeInput interface {
	pulumi.Input

	ToOutboundTypeOutput() OutboundTypeOutput
	ToOutboundTypeOutputWithContext(context.Context) OutboundTypeOutput
}

var outboundTypePtrType = reflect.TypeOf((**OutboundType)(nil)).Elem()

type OutboundTypePtrInput interface {
	pulumi.Input

	ToOutboundTypePtrOutput() OutboundTypePtrOutput
	ToOutboundTypePtrOutputWithContext(context.Context) OutboundTypePtrOutput
}

type outboundTypePtr string

func OutboundTypePtr(v string) OutboundTypePtrInput {
	return (*outboundTypePtr)(&v)
}

func (*outboundTypePtr) ElementType() reflect.Type {
	return outboundTypePtrType
}

func (in *outboundTypePtr) ToOutboundTypePtrOutput() OutboundTypePtrOutput {
	return pulumi.ToOutput(in).(OutboundTypePtrOutput)
}

func (in *outboundTypePtr) ToOutboundTypePtrOutputWithContext(ctx context.Context) OutboundTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OutboundTypePtrOutput)
}

type ResourceIdentityType string

const (
	// Use an implicitly created system assigned managed identity to manage cluster resources. Master components in the control plane such as kube-controller-manager will use the system assigned managed identity to manipulate Azure resources.
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	// Use a user-specified identity to manage cluster resources. Master components in the control plane such as kube-controller-manager will use the specified user assigned managed identity to manipulate Azure resources.
	ResourceIdentityTypeUserAssigned = ResourceIdentityType("UserAssigned")
	// Do not use a managed identity for the Managed Cluster, service principal will be used instead.
	ResourceIdentityTypeNone = ResourceIdentityType("None")
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

type ScaleSetEvictionPolicy string

const (
	// Nodes in the underlying Scale Set of the node pool are deleted when they're evicted.
	ScaleSetEvictionPolicyDelete = ScaleSetEvictionPolicy("Delete")
	// Nodes in the underlying Scale Set of the node pool are set to the stopped-deallocated state upon eviction. Nodes in the stopped-deallocated state count against your compute quota and can cause issues with cluster scaling or upgrading.
	ScaleSetEvictionPolicyDeallocate = ScaleSetEvictionPolicy("Deallocate")
)

func (ScaleSetEvictionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSetEvictionPolicy)(nil)).Elem()
}

func (e ScaleSetEvictionPolicy) ToScaleSetEvictionPolicyOutput() ScaleSetEvictionPolicyOutput {
	return pulumi.ToOutput(e).(ScaleSetEvictionPolicyOutput)
}

func (e ScaleSetEvictionPolicy) ToScaleSetEvictionPolicyOutputWithContext(ctx context.Context) ScaleSetEvictionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScaleSetEvictionPolicyOutput)
}

func (e ScaleSetEvictionPolicy) ToScaleSetEvictionPolicyPtrOutput() ScaleSetEvictionPolicyPtrOutput {
	return e.ToScaleSetEvictionPolicyPtrOutputWithContext(context.Background())
}

func (e ScaleSetEvictionPolicy) ToScaleSetEvictionPolicyPtrOutputWithContext(ctx context.Context) ScaleSetEvictionPolicyPtrOutput {
	return ScaleSetEvictionPolicy(e).ToScaleSetEvictionPolicyOutputWithContext(ctx).ToScaleSetEvictionPolicyPtrOutputWithContext(ctx)
}

func (e ScaleSetEvictionPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleSetEvictionPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleSetEvictionPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScaleSetEvictionPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScaleSetEvictionPolicyOutput struct{ *pulumi.OutputState }

func (ScaleSetEvictionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSetEvictionPolicy)(nil)).Elem()
}

func (o ScaleSetEvictionPolicyOutput) ToScaleSetEvictionPolicyOutput() ScaleSetEvictionPolicyOutput {
	return o
}

func (o ScaleSetEvictionPolicyOutput) ToScaleSetEvictionPolicyOutputWithContext(ctx context.Context) ScaleSetEvictionPolicyOutput {
	return o
}

func (o ScaleSetEvictionPolicyOutput) ToScaleSetEvictionPolicyPtrOutput() ScaleSetEvictionPolicyPtrOutput {
	return o.ToScaleSetEvictionPolicyPtrOutputWithContext(context.Background())
}

func (o ScaleSetEvictionPolicyOutput) ToScaleSetEvictionPolicyPtrOutputWithContext(ctx context.Context) ScaleSetEvictionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleSetEvictionPolicy) *ScaleSetEvictionPolicy {
		return &v
	}).(ScaleSetEvictionPolicyPtrOutput)
}

func (o ScaleSetEvictionPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScaleSetEvictionPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleSetEvictionPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScaleSetEvictionPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleSetEvictionPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleSetEvictionPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScaleSetEvictionPolicyPtrOutput struct{ *pulumi.OutputState }

func (ScaleSetEvictionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSetEvictionPolicy)(nil)).Elem()
}

func (o ScaleSetEvictionPolicyPtrOutput) ToScaleSetEvictionPolicyPtrOutput() ScaleSetEvictionPolicyPtrOutput {
	return o
}

func (o ScaleSetEvictionPolicyPtrOutput) ToScaleSetEvictionPolicyPtrOutputWithContext(ctx context.Context) ScaleSetEvictionPolicyPtrOutput {
	return o
}

func (o ScaleSetEvictionPolicyPtrOutput) Elem() ScaleSetEvictionPolicyOutput {
	return o.ApplyT(func(v *ScaleSetEvictionPolicy) ScaleSetEvictionPolicy {
		if v != nil {
			return *v
		}
		var ret ScaleSetEvictionPolicy
		return ret
	}).(ScaleSetEvictionPolicyOutput)
}

func (o ScaleSetEvictionPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleSetEvictionPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScaleSetEvictionPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScaleSetEvictionPolicyInput interface {
	pulumi.Input

	ToScaleSetEvictionPolicyOutput() ScaleSetEvictionPolicyOutput
	ToScaleSetEvictionPolicyOutputWithContext(context.Context) ScaleSetEvictionPolicyOutput
}

var scaleSetEvictionPolicyPtrType = reflect.TypeOf((**ScaleSetEvictionPolicy)(nil)).Elem()

type ScaleSetEvictionPolicyPtrInput interface {
	pulumi.Input

	ToScaleSetEvictionPolicyPtrOutput() ScaleSetEvictionPolicyPtrOutput
	ToScaleSetEvictionPolicyPtrOutputWithContext(context.Context) ScaleSetEvictionPolicyPtrOutput
}

type scaleSetEvictionPolicyPtr string

func ScaleSetEvictionPolicyPtr(v string) ScaleSetEvictionPolicyPtrInput {
	return (*scaleSetEvictionPolicyPtr)(&v)
}

func (*scaleSetEvictionPolicyPtr) ElementType() reflect.Type {
	return scaleSetEvictionPolicyPtrType
}

func (in *scaleSetEvictionPolicyPtr) ToScaleSetEvictionPolicyPtrOutput() ScaleSetEvictionPolicyPtrOutput {
	return pulumi.ToOutput(in).(ScaleSetEvictionPolicyPtrOutput)
}

func (in *scaleSetEvictionPolicyPtr) ToScaleSetEvictionPolicyPtrOutputWithContext(ctx context.Context) ScaleSetEvictionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScaleSetEvictionPolicyPtrOutput)
}

type ScaleSetPriority string

const (
	// Spot priority VMs will be used. There is no SLA for spot nodes. See [spot on AKS](https://docs.microsoft.com/azure/aks/spot-node-pool) for more information.
	ScaleSetPrioritySpot = ScaleSetPriority("Spot")
	// Regular VMs will be used.
	ScaleSetPriorityRegular = ScaleSetPriority("Regular")
)

func (ScaleSetPriority) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSetPriority)(nil)).Elem()
}

func (e ScaleSetPriority) ToScaleSetPriorityOutput() ScaleSetPriorityOutput {
	return pulumi.ToOutput(e).(ScaleSetPriorityOutput)
}

func (e ScaleSetPriority) ToScaleSetPriorityOutputWithContext(ctx context.Context) ScaleSetPriorityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScaleSetPriorityOutput)
}

func (e ScaleSetPriority) ToScaleSetPriorityPtrOutput() ScaleSetPriorityPtrOutput {
	return e.ToScaleSetPriorityPtrOutputWithContext(context.Background())
}

func (e ScaleSetPriority) ToScaleSetPriorityPtrOutputWithContext(ctx context.Context) ScaleSetPriorityPtrOutput {
	return ScaleSetPriority(e).ToScaleSetPriorityOutputWithContext(ctx).ToScaleSetPriorityPtrOutputWithContext(ctx)
}

func (e ScaleSetPriority) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleSetPriority) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleSetPriority) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScaleSetPriority) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScaleSetPriorityOutput struct{ *pulumi.OutputState }

func (ScaleSetPriorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSetPriority)(nil)).Elem()
}

func (o ScaleSetPriorityOutput) ToScaleSetPriorityOutput() ScaleSetPriorityOutput {
	return o
}

func (o ScaleSetPriorityOutput) ToScaleSetPriorityOutputWithContext(ctx context.Context) ScaleSetPriorityOutput {
	return o
}

func (o ScaleSetPriorityOutput) ToScaleSetPriorityPtrOutput() ScaleSetPriorityPtrOutput {
	return o.ToScaleSetPriorityPtrOutputWithContext(context.Background())
}

func (o ScaleSetPriorityOutput) ToScaleSetPriorityPtrOutputWithContext(ctx context.Context) ScaleSetPriorityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleSetPriority) *ScaleSetPriority {
		return &v
	}).(ScaleSetPriorityPtrOutput)
}

func (o ScaleSetPriorityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScaleSetPriorityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleSetPriority) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScaleSetPriorityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleSetPriorityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleSetPriority) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScaleSetPriorityPtrOutput struct{ *pulumi.OutputState }

func (ScaleSetPriorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSetPriority)(nil)).Elem()
}

func (o ScaleSetPriorityPtrOutput) ToScaleSetPriorityPtrOutput() ScaleSetPriorityPtrOutput {
	return o
}

func (o ScaleSetPriorityPtrOutput) ToScaleSetPriorityPtrOutputWithContext(ctx context.Context) ScaleSetPriorityPtrOutput {
	return o
}

func (o ScaleSetPriorityPtrOutput) Elem() ScaleSetPriorityOutput {
	return o.ApplyT(func(v *ScaleSetPriority) ScaleSetPriority {
		if v != nil {
			return *v
		}
		var ret ScaleSetPriority
		return ret
	}).(ScaleSetPriorityOutput)
}

func (o ScaleSetPriorityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleSetPriorityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScaleSetPriority) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScaleSetPriorityInput interface {
	pulumi.Input

	ToScaleSetPriorityOutput() ScaleSetPriorityOutput
	ToScaleSetPriorityOutputWithContext(context.Context) ScaleSetPriorityOutput
}

var scaleSetPriorityPtrType = reflect.TypeOf((**ScaleSetPriority)(nil)).Elem()

type ScaleSetPriorityPtrInput interface {
	pulumi.Input

	ToScaleSetPriorityPtrOutput() ScaleSetPriorityPtrOutput
	ToScaleSetPriorityPtrOutputWithContext(context.Context) ScaleSetPriorityPtrOutput
}

type scaleSetPriorityPtr string

func ScaleSetPriorityPtr(v string) ScaleSetPriorityPtrInput {
	return (*scaleSetPriorityPtr)(&v)
}

func (*scaleSetPriorityPtr) ElementType() reflect.Type {
	return scaleSetPriorityPtrType
}

func (in *scaleSetPriorityPtr) ToScaleSetPriorityPtrOutput() ScaleSetPriorityPtrOutput {
	return pulumi.ToOutput(in).(ScaleSetPriorityPtrOutput)
}

func (in *scaleSetPriorityPtr) ToScaleSetPriorityPtrOutputWithContext(ctx context.Context) ScaleSetPriorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScaleSetPriorityPtrOutput)
}

type UpgradeChannel string

const (
	// Automatically upgrade the cluster to the latest supported patch release on the latest supported minor version. In cases where the cluster is at a version of Kubernetes that is at an N-2 minor version where N is the latest supported minor version, the cluster first upgrades to the latest supported patch version on N-1 minor version. For example, if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4, 1.18.6, and 1.19.1 are available, your cluster first is upgraded to 1.18.6, then is upgraded to 1.19.1.
	UpgradeChannelRapid = UpgradeChannel("rapid")
	// Automatically upgrade the cluster to the latest supported patch release on minor version N-1, where N is the latest supported minor version. For example, if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4, 1.18.6, and 1.19.1 are available, your cluster is upgraded to 1.18.6.
	UpgradeChannelStable = UpgradeChannel("stable")
	// Automatically upgrade the cluster to the latest supported patch version when it becomes available while keeping the minor version the same. For example, if a cluster is running version 1.17.7 and versions 1.17.9, 1.18.4, 1.18.6, and 1.19.1 are available, your cluster is upgraded to 1.17.9.
	UpgradeChannelPatch = UpgradeChannel("patch")
	// Automatically upgrade the node image to the latest version available. Microsoft provides patches and new images for image nodes frequently (usually weekly), but your running nodes won't get the new images unless you do a node image upgrade. Turning on the node-image channel will automatically update your node images whenever a new version is available.
	UpgradeChannel_Node_image = UpgradeChannel("node-image")
	// Disables auto-upgrades and keeps the cluster at its current version of Kubernetes.
	UpgradeChannelNone = UpgradeChannel("none")
)

func (UpgradeChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradeChannel)(nil)).Elem()
}

func (e UpgradeChannel) ToUpgradeChannelOutput() UpgradeChannelOutput {
	return pulumi.ToOutput(e).(UpgradeChannelOutput)
}

func (e UpgradeChannel) ToUpgradeChannelOutputWithContext(ctx context.Context) UpgradeChannelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UpgradeChannelOutput)
}

func (e UpgradeChannel) ToUpgradeChannelPtrOutput() UpgradeChannelPtrOutput {
	return e.ToUpgradeChannelPtrOutputWithContext(context.Background())
}

func (e UpgradeChannel) ToUpgradeChannelPtrOutputWithContext(ctx context.Context) UpgradeChannelPtrOutput {
	return UpgradeChannel(e).ToUpgradeChannelOutputWithContext(ctx).ToUpgradeChannelPtrOutputWithContext(ctx)
}

func (e UpgradeChannel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpgradeChannel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpgradeChannel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UpgradeChannel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UpgradeChannelOutput struct{ *pulumi.OutputState }

func (UpgradeChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradeChannel)(nil)).Elem()
}

func (o UpgradeChannelOutput) ToUpgradeChannelOutput() UpgradeChannelOutput {
	return o
}

func (o UpgradeChannelOutput) ToUpgradeChannelOutputWithContext(ctx context.Context) UpgradeChannelOutput {
	return o
}

func (o UpgradeChannelOutput) ToUpgradeChannelPtrOutput() UpgradeChannelPtrOutput {
	return o.ToUpgradeChannelPtrOutputWithContext(context.Background())
}

func (o UpgradeChannelOutput) ToUpgradeChannelPtrOutputWithContext(ctx context.Context) UpgradeChannelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpgradeChannel) *UpgradeChannel {
		return &v
	}).(UpgradeChannelPtrOutput)
}

func (o UpgradeChannelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UpgradeChannelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpgradeChannel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UpgradeChannelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpgradeChannelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpgradeChannel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UpgradeChannelPtrOutput struct{ *pulumi.OutputState }

func (UpgradeChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeChannel)(nil)).Elem()
}

func (o UpgradeChannelPtrOutput) ToUpgradeChannelPtrOutput() UpgradeChannelPtrOutput {
	return o
}

func (o UpgradeChannelPtrOutput) ToUpgradeChannelPtrOutputWithContext(ctx context.Context) UpgradeChannelPtrOutput {
	return o
}

func (o UpgradeChannelPtrOutput) Elem() UpgradeChannelOutput {
	return o.ApplyT(func(v *UpgradeChannel) UpgradeChannel {
		if v != nil {
			return *v
		}
		var ret UpgradeChannel
		return ret
	}).(UpgradeChannelOutput)
}

func (o UpgradeChannelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpgradeChannelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UpgradeChannel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UpgradeChannelInput interface {
	pulumi.Input

	ToUpgradeChannelOutput() UpgradeChannelOutput
	ToUpgradeChannelOutputWithContext(context.Context) UpgradeChannelOutput
}

var upgradeChannelPtrType = reflect.TypeOf((**UpgradeChannel)(nil)).Elem()

type UpgradeChannelPtrInput interface {
	pulumi.Input

	ToUpgradeChannelPtrOutput() UpgradeChannelPtrOutput
	ToUpgradeChannelPtrOutputWithContext(context.Context) UpgradeChannelPtrOutput
}

type upgradeChannelPtr string

func UpgradeChannelPtr(v string) UpgradeChannelPtrInput {
	return (*upgradeChannelPtr)(&v)
}

func (*upgradeChannelPtr) ElementType() reflect.Type {
	return upgradeChannelPtrType
}

func (in *upgradeChannelPtr) ToUpgradeChannelPtrOutput() UpgradeChannelPtrOutput {
	return pulumi.ToOutput(in).(UpgradeChannelPtrOutput)
}

func (in *upgradeChannelPtr) ToUpgradeChannelPtrOutputWithContext(ctx context.Context) UpgradeChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UpgradeChannelPtrOutput)
}

type WeekDay string

const (
	WeekDaySunday    = WeekDay("Sunday")
	WeekDayMonday    = WeekDay("Monday")
	WeekDayTuesday   = WeekDay("Tuesday")
	WeekDayWednesday = WeekDay("Wednesday")
	WeekDayThursday  = WeekDay("Thursday")
	WeekDayFriday    = WeekDay("Friday")
	WeekDaySaturday  = WeekDay("Saturday")
)

func (WeekDay) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDay)(nil)).Elem()
}

func (e WeekDay) ToWeekDayOutput() WeekDayOutput {
	return pulumi.ToOutput(e).(WeekDayOutput)
}

func (e WeekDay) ToWeekDayOutputWithContext(ctx context.Context) WeekDayOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WeekDayOutput)
}

func (e WeekDay) ToWeekDayPtrOutput() WeekDayPtrOutput {
	return e.ToWeekDayPtrOutputWithContext(context.Background())
}

func (e WeekDay) ToWeekDayPtrOutputWithContext(ctx context.Context) WeekDayPtrOutput {
	return WeekDay(e).ToWeekDayOutputWithContext(ctx).ToWeekDayPtrOutputWithContext(ctx)
}

func (e WeekDay) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekDay) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekDay) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WeekDay) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WeekDayOutput struct{ *pulumi.OutputState }

func (WeekDayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDay)(nil)).Elem()
}

func (o WeekDayOutput) ToWeekDayOutput() WeekDayOutput {
	return o
}

func (o WeekDayOutput) ToWeekDayOutputWithContext(ctx context.Context) WeekDayOutput {
	return o
}

func (o WeekDayOutput) ToWeekDayPtrOutput() WeekDayPtrOutput {
	return o.ToWeekDayPtrOutputWithContext(context.Background())
}

func (o WeekDayOutput) ToWeekDayPtrOutputWithContext(ctx context.Context) WeekDayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeekDay) *WeekDay {
		return &v
	}).(WeekDayPtrOutput)
}

func (o WeekDayOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WeekDayOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WeekDay) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WeekDayOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WeekDayOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WeekDay) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WeekDayPtrOutput struct{ *pulumi.OutputState }

func (WeekDayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekDay)(nil)).Elem()
}

func (o WeekDayPtrOutput) ToWeekDayPtrOutput() WeekDayPtrOutput {
	return o
}

func (o WeekDayPtrOutput) ToWeekDayPtrOutputWithContext(ctx context.Context) WeekDayPtrOutput {
	return o
}

func (o WeekDayPtrOutput) Elem() WeekDayOutput {
	return o.ApplyT(func(v *WeekDay) WeekDay {
		if v != nil {
			return *v
		}
		var ret WeekDay
		return ret
	}).(WeekDayOutput)
}

func (o WeekDayPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WeekDayPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WeekDay) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WeekDayInput interface {
	pulumi.Input

	ToWeekDayOutput() WeekDayOutput
	ToWeekDayOutputWithContext(context.Context) WeekDayOutput
}

var weekDayPtrType = reflect.TypeOf((**WeekDay)(nil)).Elem()

type WeekDayPtrInput interface {
	pulumi.Input

	ToWeekDayPtrOutput() WeekDayPtrOutput
	ToWeekDayPtrOutputWithContext(context.Context) WeekDayPtrOutput
}

type weekDayPtr string

func WeekDayPtr(v string) WeekDayPtrInput {
	return (*weekDayPtr)(&v)
}

func (*weekDayPtr) ElementType() reflect.Type {
	return weekDayPtrType
}

func (in *weekDayPtr) ToWeekDayPtrOutput() WeekDayPtrOutput {
	return pulumi.ToOutput(in).(WeekDayPtrOutput)
}

func (in *weekDayPtr) ToWeekDayPtrOutputWithContext(ctx context.Context) WeekDayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WeekDayPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPoolModeInput)(nil)).Elem(), AgentPoolMode("System"))
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPoolModePtrInput)(nil)).Elem(), AgentPoolMode("System"))
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPoolTypeInput)(nil)).Elem(), AgentPoolType("VirtualMachineScaleSets"))
	pulumi.RegisterInputType(reflect.TypeOf((*AgentPoolTypePtrInput)(nil)).Elem(), AgentPoolType("VirtualMachineScaleSets"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStatusInput)(nil)).Elem(), ConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStatusPtrInput)(nil)).Elem(), ConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExpanderInput)(nil)).Elem(), Expander("least-waste"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExpanderPtrInput)(nil)).Elem(), Expander("least-waste"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationTypesInput)(nil)).Elem(), ExtendedLocationTypes("EdgeZone"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExtendedLocationTypesPtrInput)(nil)).Elem(), ExtendedLocationTypes("EdgeZone"))
	pulumi.RegisterInputType(reflect.TypeOf((*GPUInstanceProfileInput)(nil)).Elem(), GPUInstanceProfile("MIG1g"))
	pulumi.RegisterInputType(reflect.TypeOf((*GPUInstanceProfilePtrInput)(nil)).Elem(), GPUInstanceProfile("MIG1g"))
	pulumi.RegisterInputType(reflect.TypeOf((*KubeletDiskTypeInput)(nil)).Elem(), KubeletDiskType("OS"))
	pulumi.RegisterInputType(reflect.TypeOf((*KubeletDiskTypePtrInput)(nil)).Elem(), KubeletDiskType("OS"))
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseTypeInput)(nil)).Elem(), LicenseType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseTypePtrInput)(nil)).Elem(), LicenseType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerSkuInput)(nil)).Elem(), LoadBalancerSku("standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerSkuPtrInput)(nil)).Elem(), LoadBalancerSku("standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedClusterSKUNameInput)(nil)).Elem(), ManagedClusterSKUName("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedClusterSKUNamePtrInput)(nil)).Elem(), ManagedClusterSKUName("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedClusterSKUTierInput)(nil)).Elem(), ManagedClusterSKUTier("Paid"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedClusterSKUTierPtrInput)(nil)).Elem(), ManagedClusterSKUTier("Paid"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkModeInput)(nil)).Elem(), NetworkMode("transparent"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkModePtrInput)(nil)).Elem(), NetworkMode("transparent"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPluginInput)(nil)).Elem(), NetworkPlugin("azure"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPluginPtrInput)(nil)).Elem(), NetworkPlugin("azure"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyInput)(nil)).Elem(), NetworkPolicy("calico"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPolicyPtrInput)(nil)).Elem(), NetworkPolicy("calico"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSDiskTypeInput)(nil)).Elem(), OSDiskType("Managed"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSDiskTypePtrInput)(nil)).Elem(), OSDiskType("Managed"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSSKUInput)(nil)).Elem(), OSSKU("Ubuntu"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSSKUPtrInput)(nil)).Elem(), OSSKU("Ubuntu"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSTypeInput)(nil)).Elem(), OSType("Linux"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSTypePtrInput)(nil)).Elem(), OSType("Linux"))
	pulumi.RegisterInputType(reflect.TypeOf((*OutboundTypeInput)(nil)).Elem(), OutboundType("loadBalancer"))
	pulumi.RegisterInputType(reflect.TypeOf((*OutboundTypePtrInput)(nil)).Elem(), OutboundType("loadBalancer"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleSetEvictionPolicyInput)(nil)).Elem(), ScaleSetEvictionPolicy("Delete"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleSetEvictionPolicyPtrInput)(nil)).Elem(), ScaleSetEvictionPolicy("Delete"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleSetPriorityInput)(nil)).Elem(), ScaleSetPriority("Spot"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleSetPriorityPtrInput)(nil)).Elem(), ScaleSetPriority("Spot"))
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeChannelInput)(nil)).Elem(), UpgradeChannel("rapid"))
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeChannelPtrInput)(nil)).Elem(), UpgradeChannel("rapid"))
	pulumi.RegisterInputType(reflect.TypeOf((*WeekDayInput)(nil)).Elem(), WeekDay("Sunday"))
	pulumi.RegisterInputType(reflect.TypeOf((*WeekDayPtrInput)(nil)).Elem(), WeekDay("Sunday"))
	pulumi.RegisterOutputType(AgentPoolModeOutput{})
	pulumi.RegisterOutputType(AgentPoolModePtrOutput{})
	pulumi.RegisterOutputType(AgentPoolTypeOutput{})
	pulumi.RegisterOutputType(AgentPoolTypePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStatusOutput{})
	pulumi.RegisterOutputType(ConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ExpanderOutput{})
	pulumi.RegisterOutputType(ExpanderPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesPtrOutput{})
	pulumi.RegisterOutputType(GPUInstanceProfileOutput{})
	pulumi.RegisterOutputType(GPUInstanceProfilePtrOutput{})
	pulumi.RegisterOutputType(KubeletDiskTypeOutput{})
	pulumi.RegisterOutputType(KubeletDiskTypePtrOutput{})
	pulumi.RegisterOutputType(LicenseTypeOutput{})
	pulumi.RegisterOutputType(LicenseTypePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterSKUNameOutput{})
	pulumi.RegisterOutputType(ManagedClusterSKUNamePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterSKUTierOutput{})
	pulumi.RegisterOutputType(ManagedClusterSKUTierPtrOutput{})
	pulumi.RegisterOutputType(NetworkModeOutput{})
	pulumi.RegisterOutputType(NetworkModePtrOutput{})
	pulumi.RegisterOutputType(NetworkPluginOutput{})
	pulumi.RegisterOutputType(NetworkPluginPtrOutput{})
	pulumi.RegisterOutputType(NetworkPolicyOutput{})
	pulumi.RegisterOutputType(NetworkPolicyPtrOutput{})
	pulumi.RegisterOutputType(OSDiskTypeOutput{})
	pulumi.RegisterOutputType(OSDiskTypePtrOutput{})
	pulumi.RegisterOutputType(OSSKUOutput{})
	pulumi.RegisterOutputType(OSSKUPtrOutput{})
	pulumi.RegisterOutputType(OSTypeOutput{})
	pulumi.RegisterOutputType(OSTypePtrOutput{})
	pulumi.RegisterOutputType(OutboundTypeOutput{})
	pulumi.RegisterOutputType(OutboundTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(ScaleSetEvictionPolicyOutput{})
	pulumi.RegisterOutputType(ScaleSetEvictionPolicyPtrOutput{})
	pulumi.RegisterOutputType(ScaleSetPriorityOutput{})
	pulumi.RegisterOutputType(ScaleSetPriorityPtrOutput{})
	pulumi.RegisterOutputType(UpgradeChannelOutput{})
	pulumi.RegisterOutputType(UpgradeChannelPtrOutput{})
	pulumi.RegisterOutputType(WeekDayOutput{})
	pulumi.RegisterOutputType(WeekDayPtrOutput{})
}
