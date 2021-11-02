


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentVMSizeTypes string

const (
	AgentVMSizeTypes_Standard_A0     = AgentVMSizeTypes("Standard_A0")
	AgentVMSizeTypes_Standard_A1     = AgentVMSizeTypes("Standard_A1")
	AgentVMSizeTypes_Standard_A2     = AgentVMSizeTypes("Standard_A2")
	AgentVMSizeTypes_Standard_A3     = AgentVMSizeTypes("Standard_A3")
	AgentVMSizeTypes_Standard_A4     = AgentVMSizeTypes("Standard_A4")
	AgentVMSizeTypes_Standard_A5     = AgentVMSizeTypes("Standard_A5")
	AgentVMSizeTypes_Standard_A6     = AgentVMSizeTypes("Standard_A6")
	AgentVMSizeTypes_Standard_A7     = AgentVMSizeTypes("Standard_A7")
	AgentVMSizeTypes_Standard_A8     = AgentVMSizeTypes("Standard_A8")
	AgentVMSizeTypes_Standard_A9     = AgentVMSizeTypes("Standard_A9")
	AgentVMSizeTypes_Standard_A10    = AgentVMSizeTypes("Standard_A10")
	AgentVMSizeTypes_Standard_A11    = AgentVMSizeTypes("Standard_A11")
	AgentVMSizeTypes_Standard_D1     = AgentVMSizeTypes("Standard_D1")
	AgentVMSizeTypes_Standard_D2     = AgentVMSizeTypes("Standard_D2")
	AgentVMSizeTypes_Standard_D3     = AgentVMSizeTypes("Standard_D3")
	AgentVMSizeTypes_Standard_D4     = AgentVMSizeTypes("Standard_D4")
	AgentVMSizeTypes_Standard_D11    = AgentVMSizeTypes("Standard_D11")
	AgentVMSizeTypes_Standard_D12    = AgentVMSizeTypes("Standard_D12")
	AgentVMSizeTypes_Standard_D13    = AgentVMSizeTypes("Standard_D13")
	AgentVMSizeTypes_Standard_D14    = AgentVMSizeTypes("Standard_D14")
	AgentVMSizeTypes_Standard_D1_v2  = AgentVMSizeTypes("Standard_D1_v2")
	AgentVMSizeTypes_Standard_D2_v2  = AgentVMSizeTypes("Standard_D2_v2")
	AgentVMSizeTypes_Standard_D3_v2  = AgentVMSizeTypes("Standard_D3_v2")
	AgentVMSizeTypes_Standard_D4_v2  = AgentVMSizeTypes("Standard_D4_v2")
	AgentVMSizeTypes_Standard_D5_v2  = AgentVMSizeTypes("Standard_D5_v2")
	AgentVMSizeTypes_Standard_D11_v2 = AgentVMSizeTypes("Standard_D11_v2")
	AgentVMSizeTypes_Standard_D12_v2 = AgentVMSizeTypes("Standard_D12_v2")
	AgentVMSizeTypes_Standard_D13_v2 = AgentVMSizeTypes("Standard_D13_v2")
	AgentVMSizeTypes_Standard_D14_v2 = AgentVMSizeTypes("Standard_D14_v2")
	AgentVMSizeTypes_Standard_G1     = AgentVMSizeTypes("Standard_G1")
	AgentVMSizeTypes_Standard_G2     = AgentVMSizeTypes("Standard_G2")
	AgentVMSizeTypes_Standard_G3     = AgentVMSizeTypes("Standard_G3")
	AgentVMSizeTypes_Standard_G4     = AgentVMSizeTypes("Standard_G4")
	AgentVMSizeTypes_Standard_G5     = AgentVMSizeTypes("Standard_G5")
	AgentVMSizeTypes_Standard_DS1    = AgentVMSizeTypes("Standard_DS1")
	AgentVMSizeTypes_Standard_DS2    = AgentVMSizeTypes("Standard_DS2")
	AgentVMSizeTypes_Standard_DS3    = AgentVMSizeTypes("Standard_DS3")
	AgentVMSizeTypes_Standard_DS4    = AgentVMSizeTypes("Standard_DS4")
	AgentVMSizeTypes_Standard_DS11   = AgentVMSizeTypes("Standard_DS11")
	AgentVMSizeTypes_Standard_DS12   = AgentVMSizeTypes("Standard_DS12")
	AgentVMSizeTypes_Standard_DS13   = AgentVMSizeTypes("Standard_DS13")
	AgentVMSizeTypes_Standard_DS14   = AgentVMSizeTypes("Standard_DS14")
	AgentVMSizeTypes_Standard_GS1    = AgentVMSizeTypes("Standard_GS1")
	AgentVMSizeTypes_Standard_GS2    = AgentVMSizeTypes("Standard_GS2")
	AgentVMSizeTypes_Standard_GS3    = AgentVMSizeTypes("Standard_GS3")
	AgentVMSizeTypes_Standard_GS4    = AgentVMSizeTypes("Standard_GS4")
	AgentVMSizeTypes_Standard_GS5    = AgentVMSizeTypes("Standard_GS5")
)

func (AgentVMSizeTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentVMSizeTypes)(nil)).Elem()
}

func (e AgentVMSizeTypes) ToAgentVMSizeTypesOutput() AgentVMSizeTypesOutput {
	return pulumi.ToOutput(e).(AgentVMSizeTypesOutput)
}

func (e AgentVMSizeTypes) ToAgentVMSizeTypesOutputWithContext(ctx context.Context) AgentVMSizeTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AgentVMSizeTypesOutput)
}

func (e AgentVMSizeTypes) ToAgentVMSizeTypesPtrOutput() AgentVMSizeTypesPtrOutput {
	return e.ToAgentVMSizeTypesPtrOutputWithContext(context.Background())
}

func (e AgentVMSizeTypes) ToAgentVMSizeTypesPtrOutputWithContext(ctx context.Context) AgentVMSizeTypesPtrOutput {
	return AgentVMSizeTypes(e).ToAgentVMSizeTypesOutputWithContext(ctx).ToAgentVMSizeTypesPtrOutputWithContext(ctx)
}

func (e AgentVMSizeTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgentVMSizeTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgentVMSizeTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AgentVMSizeTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AgentVMSizeTypesOutput struct{ *pulumi.OutputState }

func (AgentVMSizeTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentVMSizeTypes)(nil)).Elem()
}

func (o AgentVMSizeTypesOutput) ToAgentVMSizeTypesOutput() AgentVMSizeTypesOutput {
	return o
}

func (o AgentVMSizeTypesOutput) ToAgentVMSizeTypesOutputWithContext(ctx context.Context) AgentVMSizeTypesOutput {
	return o
}

func (o AgentVMSizeTypesOutput) ToAgentVMSizeTypesPtrOutput() AgentVMSizeTypesPtrOutput {
	return o.ToAgentVMSizeTypesPtrOutputWithContext(context.Background())
}

func (o AgentVMSizeTypesOutput) ToAgentVMSizeTypesPtrOutputWithContext(ctx context.Context) AgentVMSizeTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentVMSizeTypes) *AgentVMSizeTypes {
		return &v
	}).(AgentVMSizeTypesPtrOutput)
}

func (o AgentVMSizeTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AgentVMSizeTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgentVMSizeTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AgentVMSizeTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgentVMSizeTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgentVMSizeTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AgentVMSizeTypesPtrOutput struct{ *pulumi.OutputState }

func (AgentVMSizeTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentVMSizeTypes)(nil)).Elem()
}

func (o AgentVMSizeTypesPtrOutput) ToAgentVMSizeTypesPtrOutput() AgentVMSizeTypesPtrOutput {
	return o
}

func (o AgentVMSizeTypesPtrOutput) ToAgentVMSizeTypesPtrOutputWithContext(ctx context.Context) AgentVMSizeTypesPtrOutput {
	return o
}

func (o AgentVMSizeTypesPtrOutput) Elem() AgentVMSizeTypesOutput {
	return o.ApplyT(func(v *AgentVMSizeTypes) AgentVMSizeTypes {
		if v != nil {
			return *v
		}
		var ret AgentVMSizeTypes
		return ret
	}).(AgentVMSizeTypesOutput)
}

func (o AgentVMSizeTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgentVMSizeTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AgentVMSizeTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AgentVMSizeTypesInput interface {
	pulumi.Input

	ToAgentVMSizeTypesOutput() AgentVMSizeTypesOutput
	ToAgentVMSizeTypesOutputWithContext(context.Context) AgentVMSizeTypesOutput
}

var agentVMSizeTypesPtrType = reflect.TypeOf((**AgentVMSizeTypes)(nil)).Elem()

type AgentVMSizeTypesPtrInput interface {
	pulumi.Input

	ToAgentVMSizeTypesPtrOutput() AgentVMSizeTypesPtrOutput
	ToAgentVMSizeTypesPtrOutputWithContext(context.Context) AgentVMSizeTypesPtrOutput
}

type agentVMSizeTypesPtr string

func AgentVMSizeTypesPtr(v string) AgentVMSizeTypesPtrInput {
	return (*agentVMSizeTypesPtr)(&v)
}

func (*agentVMSizeTypesPtr) ElementType() reflect.Type {
	return agentVMSizeTypesPtrType
}

func (in *agentVMSizeTypesPtr) ToAgentVMSizeTypesPtrOutput() AgentVMSizeTypesPtrOutput {
	return pulumi.ToOutput(in).(AgentVMSizeTypesPtrOutput)
}

func (in *agentVMSizeTypesPtr) ToAgentVMSizeTypesPtrOutputWithContext(ctx context.Context) AgentVMSizeTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AgentVMSizeTypesPtrOutput)
}

type ClusterType string

const (
	ClusterTypeACS   = ClusterType("ACS")
	ClusterTypeLocal = ClusterType("Local")
)

func (ClusterType) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterType)(nil)).Elem()
}

func (e ClusterType) ToClusterTypeOutput() ClusterTypeOutput {
	return pulumi.ToOutput(e).(ClusterTypeOutput)
}

func (e ClusterType) ToClusterTypeOutputWithContext(ctx context.Context) ClusterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterTypeOutput)
}

func (e ClusterType) ToClusterTypePtrOutput() ClusterTypePtrOutput {
	return e.ToClusterTypePtrOutputWithContext(context.Background())
}

func (e ClusterType) ToClusterTypePtrOutputWithContext(ctx context.Context) ClusterTypePtrOutput {
	return ClusterType(e).ToClusterTypeOutputWithContext(ctx).ToClusterTypePtrOutputWithContext(ctx)
}

func (e ClusterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterTypeOutput struct{ *pulumi.OutputState }

func (ClusterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterType)(nil)).Elem()
}

func (o ClusterTypeOutput) ToClusterTypeOutput() ClusterTypeOutput {
	return o
}

func (o ClusterTypeOutput) ToClusterTypeOutputWithContext(ctx context.Context) ClusterTypeOutput {
	return o
}

func (o ClusterTypeOutput) ToClusterTypePtrOutput() ClusterTypePtrOutput {
	return o.ToClusterTypePtrOutputWithContext(context.Background())
}

func (o ClusterTypeOutput) ToClusterTypePtrOutputWithContext(ctx context.Context) ClusterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterType) *ClusterType {
		return &v
	}).(ClusterTypePtrOutput)
}

func (o ClusterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterTypePtrOutput struct{ *pulumi.OutputState }

func (ClusterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterType)(nil)).Elem()
}

func (o ClusterTypePtrOutput) ToClusterTypePtrOutput() ClusterTypePtrOutput {
	return o
}

func (o ClusterTypePtrOutput) ToClusterTypePtrOutputWithContext(ctx context.Context) ClusterTypePtrOutput {
	return o
}

func (o ClusterTypePtrOutput) Elem() ClusterTypeOutput {
	return o.ApplyT(func(v *ClusterType) ClusterType {
		if v != nil {
			return *v
		}
		var ret ClusterType
		return ret
	}).(ClusterTypeOutput)
}

func (o ClusterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClusterTypeInput interface {
	pulumi.Input

	ToClusterTypeOutput() ClusterTypeOutput
	ToClusterTypeOutputWithContext(context.Context) ClusterTypeOutput
}

var clusterTypePtrType = reflect.TypeOf((**ClusterType)(nil)).Elem()

type ClusterTypePtrInput interface {
	pulumi.Input

	ToClusterTypePtrOutput() ClusterTypePtrOutput
	ToClusterTypePtrOutputWithContext(context.Context) ClusterTypePtrOutput
}

type clusterTypePtr string

func ClusterTypePtr(v string) ClusterTypePtrInput {
	return (*clusterTypePtr)(&v)
}

func (*clusterTypePtr) ElementType() reflect.Type {
	return clusterTypePtrType
}

func (in *clusterTypePtr) ToClusterTypePtrOutput() ClusterTypePtrOutput {
	return pulumi.ToOutput(in).(ClusterTypePtrOutput)
}

func (in *clusterTypePtr) ToClusterTypePtrOutputWithContext(ctx context.Context) ClusterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterTypePtrOutput)
}

type OrchestratorType string

const (
	OrchestratorTypeKubernetes = OrchestratorType("Kubernetes")
	OrchestratorTypeNone       = OrchestratorType("None")
)

func (OrchestratorType) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorType)(nil)).Elem()
}

func (e OrchestratorType) ToOrchestratorTypeOutput() OrchestratorTypeOutput {
	return pulumi.ToOutput(e).(OrchestratorTypeOutput)
}

func (e OrchestratorType) ToOrchestratorTypeOutputWithContext(ctx context.Context) OrchestratorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OrchestratorTypeOutput)
}

func (e OrchestratorType) ToOrchestratorTypePtrOutput() OrchestratorTypePtrOutput {
	return e.ToOrchestratorTypePtrOutputWithContext(context.Background())
}

func (e OrchestratorType) ToOrchestratorTypePtrOutputWithContext(ctx context.Context) OrchestratorTypePtrOutput {
	return OrchestratorType(e).ToOrchestratorTypeOutputWithContext(ctx).ToOrchestratorTypePtrOutputWithContext(ctx)
}

func (e OrchestratorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrchestratorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrchestratorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OrchestratorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OrchestratorTypeOutput struct{ *pulumi.OutputState }

func (OrchestratorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorType)(nil)).Elem()
}

func (o OrchestratorTypeOutput) ToOrchestratorTypeOutput() OrchestratorTypeOutput {
	return o
}

func (o OrchestratorTypeOutput) ToOrchestratorTypeOutputWithContext(ctx context.Context) OrchestratorTypeOutput {
	return o
}

func (o OrchestratorTypeOutput) ToOrchestratorTypePtrOutput() OrchestratorTypePtrOutput {
	return o.ToOrchestratorTypePtrOutputWithContext(context.Background())
}

func (o OrchestratorTypeOutput) ToOrchestratorTypePtrOutputWithContext(ctx context.Context) OrchestratorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrchestratorType) *OrchestratorType {
		return &v
	}).(OrchestratorTypePtrOutput)
}

func (o OrchestratorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OrchestratorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OrchestratorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OrchestratorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OrchestratorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OrchestratorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OrchestratorTypePtrOutput struct{ *pulumi.OutputState }

func (OrchestratorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratorType)(nil)).Elem()
}

func (o OrchestratorTypePtrOutput) ToOrchestratorTypePtrOutput() OrchestratorTypePtrOutput {
	return o
}

func (o OrchestratorTypePtrOutput) ToOrchestratorTypePtrOutputWithContext(ctx context.Context) OrchestratorTypePtrOutput {
	return o
}

func (o OrchestratorTypePtrOutput) Elem() OrchestratorTypeOutput {
	return o.ApplyT(func(v *OrchestratorType) OrchestratorType {
		if v != nil {
			return *v
		}
		var ret OrchestratorType
		return ret
	}).(OrchestratorTypeOutput)
}

func (o OrchestratorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OrchestratorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OrchestratorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OrchestratorTypeInput interface {
	pulumi.Input

	ToOrchestratorTypeOutput() OrchestratorTypeOutput
	ToOrchestratorTypeOutputWithContext(context.Context) OrchestratorTypeOutput
}

var orchestratorTypePtrType = reflect.TypeOf((**OrchestratorType)(nil)).Elem()

type OrchestratorTypePtrInput interface {
	pulumi.Input

	ToOrchestratorTypePtrOutput() OrchestratorTypePtrOutput
	ToOrchestratorTypePtrOutputWithContext(context.Context) OrchestratorTypePtrOutput
}

type orchestratorTypePtr string

func OrchestratorTypePtr(v string) OrchestratorTypePtrInput {
	return (*orchestratorTypePtr)(&v)
}

func (*orchestratorTypePtr) ElementType() reflect.Type {
	return orchestratorTypePtrType
}

func (in *orchestratorTypePtr) ToOrchestratorTypePtrOutput() OrchestratorTypePtrOutput {
	return pulumi.ToOutput(in).(OrchestratorTypePtrOutput)
}

func (in *orchestratorTypePtr) ToOrchestratorTypePtrOutputWithContext(ctx context.Context) OrchestratorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OrchestratorTypePtrOutput)
}

type Status string

const (
	StatusEnabled  = Status("Enabled")
	StatusDisabled = Status("Disabled")
)

func (Status) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (e Status) ToStatusOutput() StatusOutput {
	return pulumi.ToOutput(e).(StatusOutput)
}

func (e Status) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusOutput)
}

func (e Status) ToStatusPtrOutput() StatusPtrOutput {
	return e.ToStatusPtrOutputWithContext(context.Background())
}

func (e Status) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return Status(e).ToStatusOutputWithContext(ctx).ToStatusPtrOutputWithContext(ctx)
}

func (e Status) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Status) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusOutput struct{ *pulumi.OutputState }

func (StatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (o StatusOutput) ToStatusOutput() StatusOutput {
	return o
}

func (o StatusOutput) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return o
}

func (o StatusOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o.ToStatusPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Status) *Status {
		return &v
	}).(StatusPtrOutput)
}

func (o StatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusPtrOutput struct{ *pulumi.OutputState }

func (StatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Status)(nil)).Elem()
}

func (o StatusPtrOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) Elem() StatusOutput {
	return o.ApplyT(func(v *Status) Status {
		if v != nil {
			return *v
		}
		var ret Status
		return ret
	}).(StatusOutput)
}

func (o StatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Status) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StatusInput interface {
	pulumi.Input

	ToStatusOutput() StatusOutput
	ToStatusOutputWithContext(context.Context) StatusOutput
}

var statusPtrType = reflect.TypeOf((**Status)(nil)).Elem()

type StatusPtrInput interface {
	pulumi.Input

	ToStatusPtrOutput() StatusPtrOutput
	ToStatusPtrOutputWithContext(context.Context) StatusPtrOutput
}

type statusPtr string

func StatusPtr(v string) StatusPtrInput {
	return (*statusPtr)(&v)
}

func (*statusPtr) ElementType() reflect.Type {
	return statusPtrType
}

func (in *statusPtr) ToStatusPtrOutput() StatusPtrOutput {
	return pulumi.ToOutput(in).(StatusPtrOutput)
}

func (in *statusPtr) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusPtrOutput)
}

type SystemServiceType string

const (
	SystemServiceTypeNone            = SystemServiceType("None")
	SystemServiceTypeScoringFrontEnd = SystemServiceType("ScoringFrontEnd")
	SystemServiceTypeBatchFrontEnd   = SystemServiceType("BatchFrontEnd")
)

func (SystemServiceType) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemServiceType)(nil)).Elem()
}

func (e SystemServiceType) ToSystemServiceTypeOutput() SystemServiceTypeOutput {
	return pulumi.ToOutput(e).(SystemServiceTypeOutput)
}

func (e SystemServiceType) ToSystemServiceTypeOutputWithContext(ctx context.Context) SystemServiceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SystemServiceTypeOutput)
}

func (e SystemServiceType) ToSystemServiceTypePtrOutput() SystemServiceTypePtrOutput {
	return e.ToSystemServiceTypePtrOutputWithContext(context.Background())
}

func (e SystemServiceType) ToSystemServiceTypePtrOutputWithContext(ctx context.Context) SystemServiceTypePtrOutput {
	return SystemServiceType(e).ToSystemServiceTypeOutputWithContext(ctx).ToSystemServiceTypePtrOutputWithContext(ctx)
}

func (e SystemServiceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemServiceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemServiceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SystemServiceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SystemServiceTypeOutput struct{ *pulumi.OutputState }

func (SystemServiceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemServiceType)(nil)).Elem()
}

func (o SystemServiceTypeOutput) ToSystemServiceTypeOutput() SystemServiceTypeOutput {
	return o
}

func (o SystemServiceTypeOutput) ToSystemServiceTypeOutputWithContext(ctx context.Context) SystemServiceTypeOutput {
	return o
}

func (o SystemServiceTypeOutput) ToSystemServiceTypePtrOutput() SystemServiceTypePtrOutput {
	return o.ToSystemServiceTypePtrOutputWithContext(context.Background())
}

func (o SystemServiceTypeOutput) ToSystemServiceTypePtrOutputWithContext(ctx context.Context) SystemServiceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemServiceType) *SystemServiceType {
		return &v
	}).(SystemServiceTypePtrOutput)
}

func (o SystemServiceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SystemServiceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemServiceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SystemServiceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemServiceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemServiceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SystemServiceTypePtrOutput struct{ *pulumi.OutputState }

func (SystemServiceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemServiceType)(nil)).Elem()
}

func (o SystemServiceTypePtrOutput) ToSystemServiceTypePtrOutput() SystemServiceTypePtrOutput {
	return o
}

func (o SystemServiceTypePtrOutput) ToSystemServiceTypePtrOutputWithContext(ctx context.Context) SystemServiceTypePtrOutput {
	return o
}

func (o SystemServiceTypePtrOutput) Elem() SystemServiceTypeOutput {
	return o.ApplyT(func(v *SystemServiceType) SystemServiceType {
		if v != nil {
			return *v
		}
		var ret SystemServiceType
		return ret
	}).(SystemServiceTypeOutput)
}

func (o SystemServiceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemServiceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SystemServiceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SystemServiceTypeInput interface {
	pulumi.Input

	ToSystemServiceTypeOutput() SystemServiceTypeOutput
	ToSystemServiceTypeOutputWithContext(context.Context) SystemServiceTypeOutput
}

var systemServiceTypePtrType = reflect.TypeOf((**SystemServiceType)(nil)).Elem()

type SystemServiceTypePtrInput interface {
	pulumi.Input

	ToSystemServiceTypePtrOutput() SystemServiceTypePtrOutput
	ToSystemServiceTypePtrOutputWithContext(context.Context) SystemServiceTypePtrOutput
}

type systemServiceTypePtr string

func SystemServiceTypePtr(v string) SystemServiceTypePtrInput {
	return (*systemServiceTypePtr)(&v)
}

func (*systemServiceTypePtr) ElementType() reflect.Type {
	return systemServiceTypePtrType
}

func (in *systemServiceTypePtr) ToSystemServiceTypePtrOutput() SystemServiceTypePtrOutput {
	return pulumi.ToOutput(in).(SystemServiceTypePtrOutput)
}

func (in *systemServiceTypePtr) ToSystemServiceTypePtrOutputWithContext(ctx context.Context) SystemServiceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SystemServiceTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentVMSizeTypesOutput{})
	pulumi.RegisterOutputType(AgentVMSizeTypesPtrOutput{})
	pulumi.RegisterOutputType(ClusterTypeOutput{})
	pulumi.RegisterOutputType(ClusterTypePtrOutput{})
	pulumi.RegisterOutputType(OrchestratorTypeOutput{})
	pulumi.RegisterOutputType(OrchestratorTypePtrOutput{})
	pulumi.RegisterOutputType(StatusOutput{})
	pulumi.RegisterOutputType(StatusPtrOutput{})
	pulumi.RegisterOutputType(SystemServiceTypeOutput{})
	pulumi.RegisterOutputType(SystemServiceTypePtrOutput{})
}
