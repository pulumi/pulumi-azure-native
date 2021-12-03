


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentAutoUpdateStatus string

const (
	AgentAutoUpdateStatusDisabled = AgentAutoUpdateStatus("Disabled")
	AgentAutoUpdateStatusEnabled  = AgentAutoUpdateStatus("Enabled")
)

func (AgentAutoUpdateStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentAutoUpdateStatus)(nil)).Elem()
}

func (e AgentAutoUpdateStatus) ToAgentAutoUpdateStatusOutput() AgentAutoUpdateStatusOutput {
	return pulumi.ToOutput(e).(AgentAutoUpdateStatusOutput)
}

func (e AgentAutoUpdateStatus) ToAgentAutoUpdateStatusOutputWithContext(ctx context.Context) AgentAutoUpdateStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AgentAutoUpdateStatusOutput)
}

func (e AgentAutoUpdateStatus) ToAgentAutoUpdateStatusPtrOutput() AgentAutoUpdateStatusPtrOutput {
	return e.ToAgentAutoUpdateStatusPtrOutputWithContext(context.Background())
}

func (e AgentAutoUpdateStatus) ToAgentAutoUpdateStatusPtrOutputWithContext(ctx context.Context) AgentAutoUpdateStatusPtrOutput {
	return AgentAutoUpdateStatus(e).ToAgentAutoUpdateStatusOutputWithContext(ctx).ToAgentAutoUpdateStatusPtrOutputWithContext(ctx)
}

func (e AgentAutoUpdateStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgentAutoUpdateStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgentAutoUpdateStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AgentAutoUpdateStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AgentAutoUpdateStatusOutput struct{ *pulumi.OutputState }

func (AgentAutoUpdateStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentAutoUpdateStatus)(nil)).Elem()
}

func (o AgentAutoUpdateStatusOutput) ToAgentAutoUpdateStatusOutput() AgentAutoUpdateStatusOutput {
	return o
}

func (o AgentAutoUpdateStatusOutput) ToAgentAutoUpdateStatusOutputWithContext(ctx context.Context) AgentAutoUpdateStatusOutput {
	return o
}

func (o AgentAutoUpdateStatusOutput) ToAgentAutoUpdateStatusPtrOutput() AgentAutoUpdateStatusPtrOutput {
	return o.ToAgentAutoUpdateStatusPtrOutputWithContext(context.Background())
}

func (o AgentAutoUpdateStatusOutput) ToAgentAutoUpdateStatusPtrOutputWithContext(ctx context.Context) AgentAutoUpdateStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentAutoUpdateStatus) *AgentAutoUpdateStatus {
		return &v
	}).(AgentAutoUpdateStatusPtrOutput)
}

func (o AgentAutoUpdateStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AgentAutoUpdateStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgentAutoUpdateStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AgentAutoUpdateStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgentAutoUpdateStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgentAutoUpdateStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AgentAutoUpdateStatusPtrOutput struct{ *pulumi.OutputState }

func (AgentAutoUpdateStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentAutoUpdateStatus)(nil)).Elem()
}

func (o AgentAutoUpdateStatusPtrOutput) ToAgentAutoUpdateStatusPtrOutput() AgentAutoUpdateStatusPtrOutput {
	return o
}

func (o AgentAutoUpdateStatusPtrOutput) ToAgentAutoUpdateStatusPtrOutputWithContext(ctx context.Context) AgentAutoUpdateStatusPtrOutput {
	return o
}

func (o AgentAutoUpdateStatusPtrOutput) Elem() AgentAutoUpdateStatusOutput {
	return o.ApplyT(func(v *AgentAutoUpdateStatus) AgentAutoUpdateStatus {
		if v != nil {
			return *v
		}
		var ret AgentAutoUpdateStatus
		return ret
	}).(AgentAutoUpdateStatusOutput)
}

func (o AgentAutoUpdateStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgentAutoUpdateStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AgentAutoUpdateStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AgentAutoUpdateStatusInput interface {
	pulumi.Input

	ToAgentAutoUpdateStatusOutput() AgentAutoUpdateStatusOutput
	ToAgentAutoUpdateStatusOutputWithContext(context.Context) AgentAutoUpdateStatusOutput
}

var agentAutoUpdateStatusPtrType = reflect.TypeOf((**AgentAutoUpdateStatus)(nil)).Elem()

type AgentAutoUpdateStatusPtrInput interface {
	pulumi.Input

	ToAgentAutoUpdateStatusPtrOutput() AgentAutoUpdateStatusPtrOutput
	ToAgentAutoUpdateStatusPtrOutputWithContext(context.Context) AgentAutoUpdateStatusPtrOutput
}

type agentAutoUpdateStatusPtr string

func AgentAutoUpdateStatusPtr(v string) AgentAutoUpdateStatusPtrInput {
	return (*agentAutoUpdateStatusPtr)(&v)
}

func (*agentAutoUpdateStatusPtr) ElementType() reflect.Type {
	return agentAutoUpdateStatusPtrType
}

func (in *agentAutoUpdateStatusPtr) ToAgentAutoUpdateStatusPtrOutput() AgentAutoUpdateStatusPtrOutput {
	return pulumi.ToOutput(in).(AgentAutoUpdateStatusPtrOutput)
}

func (in *agentAutoUpdateStatusPtr) ToAgentAutoUpdateStatusPtrOutputWithContext(ctx context.Context) AgentAutoUpdateStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AgentAutoUpdateStatusPtrOutput)
}

type AutomationAccountAuthenticationType string

const (
	AutomationAccountAuthenticationTypeRunAsAccount           = AutomationAccountAuthenticationType("RunAsAccount")
	AutomationAccountAuthenticationTypeSystemAssignedIdentity = AutomationAccountAuthenticationType("SystemAssignedIdentity")
)

func (AutomationAccountAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationAccountAuthenticationType)(nil)).Elem()
}

func (e AutomationAccountAuthenticationType) ToAutomationAccountAuthenticationTypeOutput() AutomationAccountAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(AutomationAccountAuthenticationTypeOutput)
}

func (e AutomationAccountAuthenticationType) ToAutomationAccountAuthenticationTypeOutputWithContext(ctx context.Context) AutomationAccountAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutomationAccountAuthenticationTypeOutput)
}

func (e AutomationAccountAuthenticationType) ToAutomationAccountAuthenticationTypePtrOutput() AutomationAccountAuthenticationTypePtrOutput {
	return e.ToAutomationAccountAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e AutomationAccountAuthenticationType) ToAutomationAccountAuthenticationTypePtrOutputWithContext(ctx context.Context) AutomationAccountAuthenticationTypePtrOutput {
	return AutomationAccountAuthenticationType(e).ToAutomationAccountAuthenticationTypeOutputWithContext(ctx).ToAutomationAccountAuthenticationTypePtrOutputWithContext(ctx)
}

func (e AutomationAccountAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationAccountAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationAccountAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationAccountAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutomationAccountAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (AutomationAccountAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationAccountAuthenticationType)(nil)).Elem()
}

func (o AutomationAccountAuthenticationTypeOutput) ToAutomationAccountAuthenticationTypeOutput() AutomationAccountAuthenticationTypeOutput {
	return o
}

func (o AutomationAccountAuthenticationTypeOutput) ToAutomationAccountAuthenticationTypeOutputWithContext(ctx context.Context) AutomationAccountAuthenticationTypeOutput {
	return o
}

func (o AutomationAccountAuthenticationTypeOutput) ToAutomationAccountAuthenticationTypePtrOutput() AutomationAccountAuthenticationTypePtrOutput {
	return o.ToAutomationAccountAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o AutomationAccountAuthenticationTypeOutput) ToAutomationAccountAuthenticationTypePtrOutputWithContext(ctx context.Context) AutomationAccountAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomationAccountAuthenticationType) *AutomationAccountAuthenticationType {
		return &v
	}).(AutomationAccountAuthenticationTypePtrOutput)
}

func (o AutomationAccountAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutomationAccountAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationAccountAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutomationAccountAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationAccountAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationAccountAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutomationAccountAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (AutomationAccountAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationAccountAuthenticationType)(nil)).Elem()
}

func (o AutomationAccountAuthenticationTypePtrOutput) ToAutomationAccountAuthenticationTypePtrOutput() AutomationAccountAuthenticationTypePtrOutput {
	return o
}

func (o AutomationAccountAuthenticationTypePtrOutput) ToAutomationAccountAuthenticationTypePtrOutputWithContext(ctx context.Context) AutomationAccountAuthenticationTypePtrOutput {
	return o
}

func (o AutomationAccountAuthenticationTypePtrOutput) Elem() AutomationAccountAuthenticationTypeOutput {
	return o.ApplyT(func(v *AutomationAccountAuthenticationType) AutomationAccountAuthenticationType {
		if v != nil {
			return *v
		}
		var ret AutomationAccountAuthenticationType
		return ret
	}).(AutomationAccountAuthenticationTypeOutput)
}

func (o AutomationAccountAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationAccountAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutomationAccountAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutomationAccountAuthenticationTypeInput interface {
	pulumi.Input

	ToAutomationAccountAuthenticationTypeOutput() AutomationAccountAuthenticationTypeOutput
	ToAutomationAccountAuthenticationTypeOutputWithContext(context.Context) AutomationAccountAuthenticationTypeOutput
}

var automationAccountAuthenticationTypePtrType = reflect.TypeOf((**AutomationAccountAuthenticationType)(nil)).Elem()

type AutomationAccountAuthenticationTypePtrInput interface {
	pulumi.Input

	ToAutomationAccountAuthenticationTypePtrOutput() AutomationAccountAuthenticationTypePtrOutput
	ToAutomationAccountAuthenticationTypePtrOutputWithContext(context.Context) AutomationAccountAuthenticationTypePtrOutput
}

type automationAccountAuthenticationTypePtr string

func AutomationAccountAuthenticationTypePtr(v string) AutomationAccountAuthenticationTypePtrInput {
	return (*automationAccountAuthenticationTypePtr)(&v)
}

func (*automationAccountAuthenticationTypePtr) ElementType() reflect.Type {
	return automationAccountAuthenticationTypePtrType
}

func (in *automationAccountAuthenticationTypePtr) ToAutomationAccountAuthenticationTypePtrOutput() AutomationAccountAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(AutomationAccountAuthenticationTypePtrOutput)
}

func (in *automationAccountAuthenticationTypePtr) ToAutomationAccountAuthenticationTypePtrOutputWithContext(ctx context.Context) AutomationAccountAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutomationAccountAuthenticationTypePtrOutput)
}

type DiskAccountType string

const (
	DiskAccountType_Standard_LRS    = DiskAccountType("Standard_LRS")
	DiskAccountType_Premium_LRS     = DiskAccountType("Premium_LRS")
	DiskAccountType_StandardSSD_LRS = DiskAccountType("StandardSSD_LRS")
)

func (DiskAccountType) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskAccountType)(nil)).Elem()
}

func (e DiskAccountType) ToDiskAccountTypeOutput() DiskAccountTypeOutput {
	return pulumi.ToOutput(e).(DiskAccountTypeOutput)
}

func (e DiskAccountType) ToDiskAccountTypeOutputWithContext(ctx context.Context) DiskAccountTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskAccountTypeOutput)
}

func (e DiskAccountType) ToDiskAccountTypePtrOutput() DiskAccountTypePtrOutput {
	return e.ToDiskAccountTypePtrOutputWithContext(context.Background())
}

func (e DiskAccountType) ToDiskAccountTypePtrOutputWithContext(ctx context.Context) DiskAccountTypePtrOutput {
	return DiskAccountType(e).ToDiskAccountTypeOutputWithContext(ctx).ToDiskAccountTypePtrOutputWithContext(ctx)
}

func (e DiskAccountType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskAccountType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskAccountType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskAccountType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskAccountTypeOutput struct{ *pulumi.OutputState }

func (DiskAccountTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskAccountType)(nil)).Elem()
}

func (o DiskAccountTypeOutput) ToDiskAccountTypeOutput() DiskAccountTypeOutput {
	return o
}

func (o DiskAccountTypeOutput) ToDiskAccountTypeOutputWithContext(ctx context.Context) DiskAccountTypeOutput {
	return o
}

func (o DiskAccountTypeOutput) ToDiskAccountTypePtrOutput() DiskAccountTypePtrOutput {
	return o.ToDiskAccountTypePtrOutputWithContext(context.Background())
}

func (o DiskAccountTypeOutput) ToDiskAccountTypePtrOutputWithContext(ctx context.Context) DiskAccountTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskAccountType) *DiskAccountType {
		return &v
	}).(DiskAccountTypePtrOutput)
}

func (o DiskAccountTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskAccountTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskAccountType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskAccountTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskAccountTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskAccountType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskAccountTypePtrOutput struct{ *pulumi.OutputState }

func (DiskAccountTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAccountType)(nil)).Elem()
}

func (o DiskAccountTypePtrOutput) ToDiskAccountTypePtrOutput() DiskAccountTypePtrOutput {
	return o
}

func (o DiskAccountTypePtrOutput) ToDiskAccountTypePtrOutputWithContext(ctx context.Context) DiskAccountTypePtrOutput {
	return o
}

func (o DiskAccountTypePtrOutput) Elem() DiskAccountTypeOutput {
	return o.ApplyT(func(v *DiskAccountType) DiskAccountType {
		if v != nil {
			return *v
		}
		var ret DiskAccountType
		return ret
	}).(DiskAccountTypeOutput)
}

func (o DiskAccountTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskAccountTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskAccountType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskAccountTypeInput interface {
	pulumi.Input

	ToDiskAccountTypeOutput() DiskAccountTypeOutput
	ToDiskAccountTypeOutputWithContext(context.Context) DiskAccountTypeOutput
}

var diskAccountTypePtrType = reflect.TypeOf((**DiskAccountType)(nil)).Elem()

type DiskAccountTypePtrInput interface {
	pulumi.Input

	ToDiskAccountTypePtrOutput() DiskAccountTypePtrOutput
	ToDiskAccountTypePtrOutputWithContext(context.Context) DiskAccountTypePtrOutput
}

type diskAccountTypePtr string

func DiskAccountTypePtr(v string) DiskAccountTypePtrInput {
	return (*diskAccountTypePtr)(&v)
}

func (*diskAccountTypePtr) ElementType() reflect.Type {
	return diskAccountTypePtrType
}

func (in *diskAccountTypePtr) ToDiskAccountTypePtrOutput() DiskAccountTypePtrOutput {
	return pulumi.ToOutput(in).(DiskAccountTypePtrOutput)
}

func (in *diskAccountTypePtr) ToDiskAccountTypePtrOutputWithContext(ctx context.Context) DiskAccountTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskAccountTypePtrOutput)
}

type ExtendedLocationType string

const (
	ExtendedLocationTypeEdgeZone = ExtendedLocationType("EdgeZone")
)

func (ExtendedLocationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationType)(nil)).Elem()
}

func (e ExtendedLocationType) ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput {
	return pulumi.ToOutput(e).(ExtendedLocationTypeOutput)
}

func (e ExtendedLocationType) ToExtendedLocationTypeOutputWithContext(ctx context.Context) ExtendedLocationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExtendedLocationTypeOutput)
}

func (e ExtendedLocationType) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return e.ToExtendedLocationTypePtrOutputWithContext(context.Background())
}

func (e ExtendedLocationType) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return ExtendedLocationType(e).ToExtendedLocationTypeOutputWithContext(ctx).ToExtendedLocationTypePtrOutputWithContext(ctx)
}

func (e ExtendedLocationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExtendedLocationTypeOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationType)(nil)).Elem()
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput {
	return o
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypeOutputWithContext(ctx context.Context) ExtendedLocationTypeOutput {
	return o
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return o.ToExtendedLocationTypePtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationType) *ExtendedLocationType {
		return &v
	}).(ExtendedLocationTypePtrOutput)
}

func (o ExtendedLocationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExtendedLocationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationTypePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationType)(nil)).Elem()
}

func (o ExtendedLocationTypePtrOutput) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return o
}

func (o ExtendedLocationTypePtrOutput) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return o
}

func (o ExtendedLocationTypePtrOutput) Elem() ExtendedLocationTypeOutput {
	return o.ApplyT(func(v *ExtendedLocationType) ExtendedLocationType {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationType
		return ret
	}).(ExtendedLocationTypeOutput)
}

func (o ExtendedLocationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExtendedLocationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExtendedLocationTypeInput interface {
	pulumi.Input

	ToExtendedLocationTypeOutput() ExtendedLocationTypeOutput
	ToExtendedLocationTypeOutputWithContext(context.Context) ExtendedLocationTypeOutput
}

var extendedLocationTypePtrType = reflect.TypeOf((**ExtendedLocationType)(nil)).Elem()

type ExtendedLocationTypePtrInput interface {
	pulumi.Input

	ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput
	ToExtendedLocationTypePtrOutputWithContext(context.Context) ExtendedLocationTypePtrOutput
}

type extendedLocationTypePtr string

func ExtendedLocationTypePtr(v string) ExtendedLocationTypePtrInput {
	return (*extendedLocationTypePtr)(&v)
}

func (*extendedLocationTypePtr) ElementType() reflect.Type {
	return extendedLocationTypePtrType
}

func (in *extendedLocationTypePtr) ToExtendedLocationTypePtrOutput() ExtendedLocationTypePtrOutput {
	return pulumi.ToOutput(in).(ExtendedLocationTypePtrOutput)
}

func (in *extendedLocationTypePtr) ToExtendedLocationTypePtrOutputWithContext(ctx context.Context) ExtendedLocationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExtendedLocationTypePtrOutput)
}

type FailoverDeploymentModel string

const (
	FailoverDeploymentModelNotApplicable   = FailoverDeploymentModel("NotApplicable")
	FailoverDeploymentModelClassic         = FailoverDeploymentModel("Classic")
	FailoverDeploymentModelResourceManager = FailoverDeploymentModel("ResourceManager")
)

func (FailoverDeploymentModel) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverDeploymentModel)(nil)).Elem()
}

func (e FailoverDeploymentModel) ToFailoverDeploymentModelOutput() FailoverDeploymentModelOutput {
	return pulumi.ToOutput(e).(FailoverDeploymentModelOutput)
}

func (e FailoverDeploymentModel) ToFailoverDeploymentModelOutputWithContext(ctx context.Context) FailoverDeploymentModelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FailoverDeploymentModelOutput)
}

func (e FailoverDeploymentModel) ToFailoverDeploymentModelPtrOutput() FailoverDeploymentModelPtrOutput {
	return e.ToFailoverDeploymentModelPtrOutputWithContext(context.Background())
}

func (e FailoverDeploymentModel) ToFailoverDeploymentModelPtrOutputWithContext(ctx context.Context) FailoverDeploymentModelPtrOutput {
	return FailoverDeploymentModel(e).ToFailoverDeploymentModelOutputWithContext(ctx).ToFailoverDeploymentModelPtrOutputWithContext(ctx)
}

func (e FailoverDeploymentModel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FailoverDeploymentModel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FailoverDeploymentModel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FailoverDeploymentModel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FailoverDeploymentModelOutput struct{ *pulumi.OutputState }

func (FailoverDeploymentModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverDeploymentModel)(nil)).Elem()
}

func (o FailoverDeploymentModelOutput) ToFailoverDeploymentModelOutput() FailoverDeploymentModelOutput {
	return o
}

func (o FailoverDeploymentModelOutput) ToFailoverDeploymentModelOutputWithContext(ctx context.Context) FailoverDeploymentModelOutput {
	return o
}

func (o FailoverDeploymentModelOutput) ToFailoverDeploymentModelPtrOutput() FailoverDeploymentModelPtrOutput {
	return o.ToFailoverDeploymentModelPtrOutputWithContext(context.Background())
}

func (o FailoverDeploymentModelOutput) ToFailoverDeploymentModelPtrOutputWithContext(ctx context.Context) FailoverDeploymentModelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverDeploymentModel) *FailoverDeploymentModel {
		return &v
	}).(FailoverDeploymentModelPtrOutput)
}

func (o FailoverDeploymentModelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FailoverDeploymentModelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FailoverDeploymentModel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FailoverDeploymentModelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FailoverDeploymentModelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FailoverDeploymentModel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FailoverDeploymentModelPtrOutput struct{ *pulumi.OutputState }

func (FailoverDeploymentModelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverDeploymentModel)(nil)).Elem()
}

func (o FailoverDeploymentModelPtrOutput) ToFailoverDeploymentModelPtrOutput() FailoverDeploymentModelPtrOutput {
	return o
}

func (o FailoverDeploymentModelPtrOutput) ToFailoverDeploymentModelPtrOutputWithContext(ctx context.Context) FailoverDeploymentModelPtrOutput {
	return o
}

func (o FailoverDeploymentModelPtrOutput) Elem() FailoverDeploymentModelOutput {
	return o.ApplyT(func(v *FailoverDeploymentModel) FailoverDeploymentModel {
		if v != nil {
			return *v
		}
		var ret FailoverDeploymentModel
		return ret
	}).(FailoverDeploymentModelOutput)
}

func (o FailoverDeploymentModelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FailoverDeploymentModelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FailoverDeploymentModel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FailoverDeploymentModelInput interface {
	pulumi.Input

	ToFailoverDeploymentModelOutput() FailoverDeploymentModelOutput
	ToFailoverDeploymentModelOutputWithContext(context.Context) FailoverDeploymentModelOutput
}

var failoverDeploymentModelPtrType = reflect.TypeOf((**FailoverDeploymentModel)(nil)).Elem()

type FailoverDeploymentModelPtrInput interface {
	pulumi.Input

	ToFailoverDeploymentModelPtrOutput() FailoverDeploymentModelPtrOutput
	ToFailoverDeploymentModelPtrOutputWithContext(context.Context) FailoverDeploymentModelPtrOutput
}

type failoverDeploymentModelPtr string

func FailoverDeploymentModelPtr(v string) FailoverDeploymentModelPtrInput {
	return (*failoverDeploymentModelPtr)(&v)
}

func (*failoverDeploymentModelPtr) ElementType() reflect.Type {
	return failoverDeploymentModelPtrType
}

func (in *failoverDeploymentModelPtr) ToFailoverDeploymentModelPtrOutput() FailoverDeploymentModelPtrOutput {
	return pulumi.ToOutput(in).(FailoverDeploymentModelPtrOutput)
}

func (in *failoverDeploymentModelPtr) ToFailoverDeploymentModelPtrOutputWithContext(ctx context.Context) FailoverDeploymentModelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FailoverDeploymentModelPtrOutput)
}

type LicenseType string

const (
	LicenseTypeNotSpecified  = LicenseType("NotSpecified")
	LicenseTypeNoLicenseType = LicenseType("NoLicenseType")
	LicenseTypeWindowsServer = LicenseType("WindowsServer")
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

type PossibleOperationsDirections string

const (
	PossibleOperationsDirectionsPrimaryToRecovery = PossibleOperationsDirections("PrimaryToRecovery")
	PossibleOperationsDirectionsRecoveryToPrimary = PossibleOperationsDirections("RecoveryToPrimary")
)

func (PossibleOperationsDirections) ElementType() reflect.Type {
	return reflect.TypeOf((*PossibleOperationsDirections)(nil)).Elem()
}

func (e PossibleOperationsDirections) ToPossibleOperationsDirectionsOutput() PossibleOperationsDirectionsOutput {
	return pulumi.ToOutput(e).(PossibleOperationsDirectionsOutput)
}

func (e PossibleOperationsDirections) ToPossibleOperationsDirectionsOutputWithContext(ctx context.Context) PossibleOperationsDirectionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PossibleOperationsDirectionsOutput)
}

func (e PossibleOperationsDirections) ToPossibleOperationsDirectionsPtrOutput() PossibleOperationsDirectionsPtrOutput {
	return e.ToPossibleOperationsDirectionsPtrOutputWithContext(context.Background())
}

func (e PossibleOperationsDirections) ToPossibleOperationsDirectionsPtrOutputWithContext(ctx context.Context) PossibleOperationsDirectionsPtrOutput {
	return PossibleOperationsDirections(e).ToPossibleOperationsDirectionsOutputWithContext(ctx).ToPossibleOperationsDirectionsPtrOutputWithContext(ctx)
}

func (e PossibleOperationsDirections) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PossibleOperationsDirections) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PossibleOperationsDirections) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PossibleOperationsDirections) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PossibleOperationsDirectionsOutput struct{ *pulumi.OutputState }

func (PossibleOperationsDirectionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PossibleOperationsDirections)(nil)).Elem()
}

func (o PossibleOperationsDirectionsOutput) ToPossibleOperationsDirectionsOutput() PossibleOperationsDirectionsOutput {
	return o
}

func (o PossibleOperationsDirectionsOutput) ToPossibleOperationsDirectionsOutputWithContext(ctx context.Context) PossibleOperationsDirectionsOutput {
	return o
}

func (o PossibleOperationsDirectionsOutput) ToPossibleOperationsDirectionsPtrOutput() PossibleOperationsDirectionsPtrOutput {
	return o.ToPossibleOperationsDirectionsPtrOutputWithContext(context.Background())
}

func (o PossibleOperationsDirectionsOutput) ToPossibleOperationsDirectionsPtrOutputWithContext(ctx context.Context) PossibleOperationsDirectionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PossibleOperationsDirections) *PossibleOperationsDirections {
		return &v
	}).(PossibleOperationsDirectionsPtrOutput)
}

func (o PossibleOperationsDirectionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PossibleOperationsDirectionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PossibleOperationsDirections) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PossibleOperationsDirectionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PossibleOperationsDirectionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PossibleOperationsDirections) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PossibleOperationsDirectionsPtrOutput struct{ *pulumi.OutputState }

func (PossibleOperationsDirectionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PossibleOperationsDirections)(nil)).Elem()
}

func (o PossibleOperationsDirectionsPtrOutput) ToPossibleOperationsDirectionsPtrOutput() PossibleOperationsDirectionsPtrOutput {
	return o
}

func (o PossibleOperationsDirectionsPtrOutput) ToPossibleOperationsDirectionsPtrOutputWithContext(ctx context.Context) PossibleOperationsDirectionsPtrOutput {
	return o
}

func (o PossibleOperationsDirectionsPtrOutput) Elem() PossibleOperationsDirectionsOutput {
	return o.ApplyT(func(v *PossibleOperationsDirections) PossibleOperationsDirections {
		if v != nil {
			return *v
		}
		var ret PossibleOperationsDirections
		return ret
	}).(PossibleOperationsDirectionsOutput)
}

func (o PossibleOperationsDirectionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PossibleOperationsDirectionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PossibleOperationsDirections) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PossibleOperationsDirectionsInput interface {
	pulumi.Input

	ToPossibleOperationsDirectionsOutput() PossibleOperationsDirectionsOutput
	ToPossibleOperationsDirectionsOutputWithContext(context.Context) PossibleOperationsDirectionsOutput
}

var possibleOperationsDirectionsPtrType = reflect.TypeOf((**PossibleOperationsDirections)(nil)).Elem()

type PossibleOperationsDirectionsPtrInput interface {
	pulumi.Input

	ToPossibleOperationsDirectionsPtrOutput() PossibleOperationsDirectionsPtrOutput
	ToPossibleOperationsDirectionsPtrOutputWithContext(context.Context) PossibleOperationsDirectionsPtrOutput
}

type possibleOperationsDirectionsPtr string

func PossibleOperationsDirectionsPtr(v string) PossibleOperationsDirectionsPtrInput {
	return (*possibleOperationsDirectionsPtr)(&v)
}

func (*possibleOperationsDirectionsPtr) ElementType() reflect.Type {
	return possibleOperationsDirectionsPtrType
}

func (in *possibleOperationsDirectionsPtr) ToPossibleOperationsDirectionsPtrOutput() PossibleOperationsDirectionsPtrOutput {
	return pulumi.ToOutput(in).(PossibleOperationsDirectionsPtrOutput)
}

func (in *possibleOperationsDirectionsPtr) ToPossibleOperationsDirectionsPtrOutputWithContext(ctx context.Context) PossibleOperationsDirectionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PossibleOperationsDirectionsPtrOutput)
}

type RecoveryPlanActionLocation string

const (
	RecoveryPlanActionLocationPrimary  = RecoveryPlanActionLocation("Primary")
	RecoveryPlanActionLocationRecovery = RecoveryPlanActionLocation("Recovery")
)

func (RecoveryPlanActionLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanActionLocation)(nil)).Elem()
}

func (e RecoveryPlanActionLocation) ToRecoveryPlanActionLocationOutput() RecoveryPlanActionLocationOutput {
	return pulumi.ToOutput(e).(RecoveryPlanActionLocationOutput)
}

func (e RecoveryPlanActionLocation) ToRecoveryPlanActionLocationOutputWithContext(ctx context.Context) RecoveryPlanActionLocationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecoveryPlanActionLocationOutput)
}

func (e RecoveryPlanActionLocation) ToRecoveryPlanActionLocationPtrOutput() RecoveryPlanActionLocationPtrOutput {
	return e.ToRecoveryPlanActionLocationPtrOutputWithContext(context.Background())
}

func (e RecoveryPlanActionLocation) ToRecoveryPlanActionLocationPtrOutputWithContext(ctx context.Context) RecoveryPlanActionLocationPtrOutput {
	return RecoveryPlanActionLocation(e).ToRecoveryPlanActionLocationOutputWithContext(ctx).ToRecoveryPlanActionLocationPtrOutputWithContext(ctx)
}

func (e RecoveryPlanActionLocation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecoveryPlanActionLocation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecoveryPlanActionLocation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecoveryPlanActionLocation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecoveryPlanActionLocationOutput struct{ *pulumi.OutputState }

func (RecoveryPlanActionLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanActionLocation)(nil)).Elem()
}

func (o RecoveryPlanActionLocationOutput) ToRecoveryPlanActionLocationOutput() RecoveryPlanActionLocationOutput {
	return o
}

func (o RecoveryPlanActionLocationOutput) ToRecoveryPlanActionLocationOutputWithContext(ctx context.Context) RecoveryPlanActionLocationOutput {
	return o
}

func (o RecoveryPlanActionLocationOutput) ToRecoveryPlanActionLocationPtrOutput() RecoveryPlanActionLocationPtrOutput {
	return o.ToRecoveryPlanActionLocationPtrOutputWithContext(context.Background())
}

func (o RecoveryPlanActionLocationOutput) ToRecoveryPlanActionLocationPtrOutputWithContext(ctx context.Context) RecoveryPlanActionLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecoveryPlanActionLocation) *RecoveryPlanActionLocation {
		return &v
	}).(RecoveryPlanActionLocationPtrOutput)
}

func (o RecoveryPlanActionLocationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecoveryPlanActionLocationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecoveryPlanActionLocation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecoveryPlanActionLocationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecoveryPlanActionLocationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecoveryPlanActionLocation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecoveryPlanActionLocationPtrOutput struct{ *pulumi.OutputState }

func (RecoveryPlanActionLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecoveryPlanActionLocation)(nil)).Elem()
}

func (o RecoveryPlanActionLocationPtrOutput) ToRecoveryPlanActionLocationPtrOutput() RecoveryPlanActionLocationPtrOutput {
	return o
}

func (o RecoveryPlanActionLocationPtrOutput) ToRecoveryPlanActionLocationPtrOutputWithContext(ctx context.Context) RecoveryPlanActionLocationPtrOutput {
	return o
}

func (o RecoveryPlanActionLocationPtrOutput) Elem() RecoveryPlanActionLocationOutput {
	return o.ApplyT(func(v *RecoveryPlanActionLocation) RecoveryPlanActionLocation {
		if v != nil {
			return *v
		}
		var ret RecoveryPlanActionLocation
		return ret
	}).(RecoveryPlanActionLocationOutput)
}

func (o RecoveryPlanActionLocationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecoveryPlanActionLocationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecoveryPlanActionLocation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RecoveryPlanActionLocationInput interface {
	pulumi.Input

	ToRecoveryPlanActionLocationOutput() RecoveryPlanActionLocationOutput
	ToRecoveryPlanActionLocationOutputWithContext(context.Context) RecoveryPlanActionLocationOutput
}

var recoveryPlanActionLocationPtrType = reflect.TypeOf((**RecoveryPlanActionLocation)(nil)).Elem()

type RecoveryPlanActionLocationPtrInput interface {
	pulumi.Input

	ToRecoveryPlanActionLocationPtrOutput() RecoveryPlanActionLocationPtrOutput
	ToRecoveryPlanActionLocationPtrOutputWithContext(context.Context) RecoveryPlanActionLocationPtrOutput
}

type recoveryPlanActionLocationPtr string

func RecoveryPlanActionLocationPtr(v string) RecoveryPlanActionLocationPtrInput {
	return (*recoveryPlanActionLocationPtr)(&v)
}

func (*recoveryPlanActionLocationPtr) ElementType() reflect.Type {
	return recoveryPlanActionLocationPtrType
}

func (in *recoveryPlanActionLocationPtr) ToRecoveryPlanActionLocationPtrOutput() RecoveryPlanActionLocationPtrOutput {
	return pulumi.ToOutput(in).(RecoveryPlanActionLocationPtrOutput)
}

func (in *recoveryPlanActionLocationPtr) ToRecoveryPlanActionLocationPtrOutputWithContext(ctx context.Context) RecoveryPlanActionLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecoveryPlanActionLocationPtrOutput)
}

type RecoveryPlanGroupType string

const (
	RecoveryPlanGroupTypeShutdown = RecoveryPlanGroupType("Shutdown")
	RecoveryPlanGroupTypeBoot     = RecoveryPlanGroupType("Boot")
	RecoveryPlanGroupTypeFailover = RecoveryPlanGroupType("Failover")
)

func (RecoveryPlanGroupType) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanGroupType)(nil)).Elem()
}

func (e RecoveryPlanGroupType) ToRecoveryPlanGroupTypeOutput() RecoveryPlanGroupTypeOutput {
	return pulumi.ToOutput(e).(RecoveryPlanGroupTypeOutput)
}

func (e RecoveryPlanGroupType) ToRecoveryPlanGroupTypeOutputWithContext(ctx context.Context) RecoveryPlanGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecoveryPlanGroupTypeOutput)
}

func (e RecoveryPlanGroupType) ToRecoveryPlanGroupTypePtrOutput() RecoveryPlanGroupTypePtrOutput {
	return e.ToRecoveryPlanGroupTypePtrOutputWithContext(context.Background())
}

func (e RecoveryPlanGroupType) ToRecoveryPlanGroupTypePtrOutputWithContext(ctx context.Context) RecoveryPlanGroupTypePtrOutput {
	return RecoveryPlanGroupType(e).ToRecoveryPlanGroupTypeOutputWithContext(ctx).ToRecoveryPlanGroupTypePtrOutputWithContext(ctx)
}

func (e RecoveryPlanGroupType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecoveryPlanGroupType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecoveryPlanGroupType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecoveryPlanGroupType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecoveryPlanGroupTypeOutput struct{ *pulumi.OutputState }

func (RecoveryPlanGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanGroupType)(nil)).Elem()
}

func (o RecoveryPlanGroupTypeOutput) ToRecoveryPlanGroupTypeOutput() RecoveryPlanGroupTypeOutput {
	return o
}

func (o RecoveryPlanGroupTypeOutput) ToRecoveryPlanGroupTypeOutputWithContext(ctx context.Context) RecoveryPlanGroupTypeOutput {
	return o
}

func (o RecoveryPlanGroupTypeOutput) ToRecoveryPlanGroupTypePtrOutput() RecoveryPlanGroupTypePtrOutput {
	return o.ToRecoveryPlanGroupTypePtrOutputWithContext(context.Background())
}

func (o RecoveryPlanGroupTypeOutput) ToRecoveryPlanGroupTypePtrOutputWithContext(ctx context.Context) RecoveryPlanGroupTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecoveryPlanGroupType) *RecoveryPlanGroupType {
		return &v
	}).(RecoveryPlanGroupTypePtrOutput)
}

func (o RecoveryPlanGroupTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecoveryPlanGroupTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecoveryPlanGroupType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecoveryPlanGroupTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecoveryPlanGroupTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecoveryPlanGroupType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecoveryPlanGroupTypePtrOutput struct{ *pulumi.OutputState }

func (RecoveryPlanGroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecoveryPlanGroupType)(nil)).Elem()
}

func (o RecoveryPlanGroupTypePtrOutput) ToRecoveryPlanGroupTypePtrOutput() RecoveryPlanGroupTypePtrOutput {
	return o
}

func (o RecoveryPlanGroupTypePtrOutput) ToRecoveryPlanGroupTypePtrOutputWithContext(ctx context.Context) RecoveryPlanGroupTypePtrOutput {
	return o
}

func (o RecoveryPlanGroupTypePtrOutput) Elem() RecoveryPlanGroupTypeOutput {
	return o.ApplyT(func(v *RecoveryPlanGroupType) RecoveryPlanGroupType {
		if v != nil {
			return *v
		}
		var ret RecoveryPlanGroupType
		return ret
	}).(RecoveryPlanGroupTypeOutput)
}

func (o RecoveryPlanGroupTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecoveryPlanGroupTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecoveryPlanGroupType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RecoveryPlanGroupTypeInput interface {
	pulumi.Input

	ToRecoveryPlanGroupTypeOutput() RecoveryPlanGroupTypeOutput
	ToRecoveryPlanGroupTypeOutputWithContext(context.Context) RecoveryPlanGroupTypeOutput
}

var recoveryPlanGroupTypePtrType = reflect.TypeOf((**RecoveryPlanGroupType)(nil)).Elem()

type RecoveryPlanGroupTypePtrInput interface {
	pulumi.Input

	ToRecoveryPlanGroupTypePtrOutput() RecoveryPlanGroupTypePtrOutput
	ToRecoveryPlanGroupTypePtrOutputWithContext(context.Context) RecoveryPlanGroupTypePtrOutput
}

type recoveryPlanGroupTypePtr string

func RecoveryPlanGroupTypePtr(v string) RecoveryPlanGroupTypePtrInput {
	return (*recoveryPlanGroupTypePtr)(&v)
}

func (*recoveryPlanGroupTypePtr) ElementType() reflect.Type {
	return recoveryPlanGroupTypePtrType
}

func (in *recoveryPlanGroupTypePtr) ToRecoveryPlanGroupTypePtrOutput() RecoveryPlanGroupTypePtrOutput {
	return pulumi.ToOutput(in).(RecoveryPlanGroupTypePtrOutput)
}

func (in *recoveryPlanGroupTypePtr) ToRecoveryPlanGroupTypePtrOutputWithContext(ctx context.Context) RecoveryPlanGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecoveryPlanGroupTypePtrOutput)
}

type ReplicationProtectedItemOperation string

const (
	ReplicationProtectedItemOperationReverseReplicate    = ReplicationProtectedItemOperation("ReverseReplicate")
	ReplicationProtectedItemOperationCommit              = ReplicationProtectedItemOperation("Commit")
	ReplicationProtectedItemOperationPlannedFailover     = ReplicationProtectedItemOperation("PlannedFailover")
	ReplicationProtectedItemOperationUnplannedFailover   = ReplicationProtectedItemOperation("UnplannedFailover")
	ReplicationProtectedItemOperationDisableProtection   = ReplicationProtectedItemOperation("DisableProtection")
	ReplicationProtectedItemOperationTestFailover        = ReplicationProtectedItemOperation("TestFailover")
	ReplicationProtectedItemOperationTestFailoverCleanup = ReplicationProtectedItemOperation("TestFailoverCleanup")
	ReplicationProtectedItemOperationFailback            = ReplicationProtectedItemOperation("Failback")
	ReplicationProtectedItemOperationFinalizeFailback    = ReplicationProtectedItemOperation("FinalizeFailback")
	ReplicationProtectedItemOperationCancelFailover      = ReplicationProtectedItemOperation("CancelFailover")
	ReplicationProtectedItemOperationChangePit           = ReplicationProtectedItemOperation("ChangePit")
	ReplicationProtectedItemOperationRepairReplication   = ReplicationProtectedItemOperation("RepairReplication")
	ReplicationProtectedItemOperationSwitchProtection    = ReplicationProtectedItemOperation("SwitchProtection")
	ReplicationProtectedItemOperationCompleteMigration   = ReplicationProtectedItemOperation("CompleteMigration")
)

func (ReplicationProtectedItemOperation) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProtectedItemOperation)(nil)).Elem()
}

func (e ReplicationProtectedItemOperation) ToReplicationProtectedItemOperationOutput() ReplicationProtectedItemOperationOutput {
	return pulumi.ToOutput(e).(ReplicationProtectedItemOperationOutput)
}

func (e ReplicationProtectedItemOperation) ToReplicationProtectedItemOperationOutputWithContext(ctx context.Context) ReplicationProtectedItemOperationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReplicationProtectedItemOperationOutput)
}

func (e ReplicationProtectedItemOperation) ToReplicationProtectedItemOperationPtrOutput() ReplicationProtectedItemOperationPtrOutput {
	return e.ToReplicationProtectedItemOperationPtrOutputWithContext(context.Background())
}

func (e ReplicationProtectedItemOperation) ToReplicationProtectedItemOperationPtrOutputWithContext(ctx context.Context) ReplicationProtectedItemOperationPtrOutput {
	return ReplicationProtectedItemOperation(e).ToReplicationProtectedItemOperationOutputWithContext(ctx).ToReplicationProtectedItemOperationPtrOutputWithContext(ctx)
}

func (e ReplicationProtectedItemOperation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationProtectedItemOperation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationProtectedItemOperation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReplicationProtectedItemOperation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReplicationProtectedItemOperationOutput struct{ *pulumi.OutputState }

func (ReplicationProtectedItemOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProtectedItemOperation)(nil)).Elem()
}

func (o ReplicationProtectedItemOperationOutput) ToReplicationProtectedItemOperationOutput() ReplicationProtectedItemOperationOutput {
	return o
}

func (o ReplicationProtectedItemOperationOutput) ToReplicationProtectedItemOperationOutputWithContext(ctx context.Context) ReplicationProtectedItemOperationOutput {
	return o
}

func (o ReplicationProtectedItemOperationOutput) ToReplicationProtectedItemOperationPtrOutput() ReplicationProtectedItemOperationPtrOutput {
	return o.ToReplicationProtectedItemOperationPtrOutputWithContext(context.Background())
}

func (o ReplicationProtectedItemOperationOutput) ToReplicationProtectedItemOperationPtrOutputWithContext(ctx context.Context) ReplicationProtectedItemOperationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationProtectedItemOperation) *ReplicationProtectedItemOperation {
		return &v
	}).(ReplicationProtectedItemOperationPtrOutput)
}

func (o ReplicationProtectedItemOperationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReplicationProtectedItemOperationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationProtectedItemOperation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReplicationProtectedItemOperationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationProtectedItemOperationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationProtectedItemOperation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReplicationProtectedItemOperationPtrOutput struct{ *pulumi.OutputState }

func (ReplicationProtectedItemOperationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectedItemOperation)(nil)).Elem()
}

func (o ReplicationProtectedItemOperationPtrOutput) ToReplicationProtectedItemOperationPtrOutput() ReplicationProtectedItemOperationPtrOutput {
	return o
}

func (o ReplicationProtectedItemOperationPtrOutput) ToReplicationProtectedItemOperationPtrOutputWithContext(ctx context.Context) ReplicationProtectedItemOperationPtrOutput {
	return o
}

func (o ReplicationProtectedItemOperationPtrOutput) Elem() ReplicationProtectedItemOperationOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemOperation) ReplicationProtectedItemOperation {
		if v != nil {
			return *v
		}
		var ret ReplicationProtectedItemOperation
		return ret
	}).(ReplicationProtectedItemOperationOutput)
}

func (o ReplicationProtectedItemOperationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationProtectedItemOperationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReplicationProtectedItemOperation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReplicationProtectedItemOperationInput interface {
	pulumi.Input

	ToReplicationProtectedItemOperationOutput() ReplicationProtectedItemOperationOutput
	ToReplicationProtectedItemOperationOutputWithContext(context.Context) ReplicationProtectedItemOperationOutput
}

var replicationProtectedItemOperationPtrType = reflect.TypeOf((**ReplicationProtectedItemOperation)(nil)).Elem()

type ReplicationProtectedItemOperationPtrInput interface {
	pulumi.Input

	ToReplicationProtectedItemOperationPtrOutput() ReplicationProtectedItemOperationPtrOutput
	ToReplicationProtectedItemOperationPtrOutputWithContext(context.Context) ReplicationProtectedItemOperationPtrOutput
}

type replicationProtectedItemOperationPtr string

func ReplicationProtectedItemOperationPtr(v string) ReplicationProtectedItemOperationPtrInput {
	return (*replicationProtectedItemOperationPtr)(&v)
}

func (*replicationProtectedItemOperationPtr) ElementType() reflect.Type {
	return replicationProtectedItemOperationPtrType
}

func (in *replicationProtectedItemOperationPtr) ToReplicationProtectedItemOperationPtrOutput() ReplicationProtectedItemOperationPtrOutput {
	return pulumi.ToOutput(in).(ReplicationProtectedItemOperationPtrOutput)
}

func (in *replicationProtectedItemOperationPtr) ToReplicationProtectedItemOperationPtrOutputWithContext(ctx context.Context) ReplicationProtectedItemOperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReplicationProtectedItemOperationPtrOutput)
}

type SetMultiVmSyncStatus string

const (
	SetMultiVmSyncStatusEnable  = SetMultiVmSyncStatus("Enable")
	SetMultiVmSyncStatusDisable = SetMultiVmSyncStatus("Disable")
)

func (SetMultiVmSyncStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*SetMultiVmSyncStatus)(nil)).Elem()
}

func (e SetMultiVmSyncStatus) ToSetMultiVmSyncStatusOutput() SetMultiVmSyncStatusOutput {
	return pulumi.ToOutput(e).(SetMultiVmSyncStatusOutput)
}

func (e SetMultiVmSyncStatus) ToSetMultiVmSyncStatusOutputWithContext(ctx context.Context) SetMultiVmSyncStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SetMultiVmSyncStatusOutput)
}

func (e SetMultiVmSyncStatus) ToSetMultiVmSyncStatusPtrOutput() SetMultiVmSyncStatusPtrOutput {
	return e.ToSetMultiVmSyncStatusPtrOutputWithContext(context.Background())
}

func (e SetMultiVmSyncStatus) ToSetMultiVmSyncStatusPtrOutputWithContext(ctx context.Context) SetMultiVmSyncStatusPtrOutput {
	return SetMultiVmSyncStatus(e).ToSetMultiVmSyncStatusOutputWithContext(ctx).ToSetMultiVmSyncStatusPtrOutputWithContext(ctx)
}

func (e SetMultiVmSyncStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SetMultiVmSyncStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SetMultiVmSyncStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SetMultiVmSyncStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SetMultiVmSyncStatusOutput struct{ *pulumi.OutputState }

func (SetMultiVmSyncStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SetMultiVmSyncStatus)(nil)).Elem()
}

func (o SetMultiVmSyncStatusOutput) ToSetMultiVmSyncStatusOutput() SetMultiVmSyncStatusOutput {
	return o
}

func (o SetMultiVmSyncStatusOutput) ToSetMultiVmSyncStatusOutputWithContext(ctx context.Context) SetMultiVmSyncStatusOutput {
	return o
}

func (o SetMultiVmSyncStatusOutput) ToSetMultiVmSyncStatusPtrOutput() SetMultiVmSyncStatusPtrOutput {
	return o.ToSetMultiVmSyncStatusPtrOutputWithContext(context.Background())
}

func (o SetMultiVmSyncStatusOutput) ToSetMultiVmSyncStatusPtrOutputWithContext(ctx context.Context) SetMultiVmSyncStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SetMultiVmSyncStatus) *SetMultiVmSyncStatus {
		return &v
	}).(SetMultiVmSyncStatusPtrOutput)
}

func (o SetMultiVmSyncStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SetMultiVmSyncStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SetMultiVmSyncStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SetMultiVmSyncStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SetMultiVmSyncStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SetMultiVmSyncStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SetMultiVmSyncStatusPtrOutput struct{ *pulumi.OutputState }

func (SetMultiVmSyncStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SetMultiVmSyncStatus)(nil)).Elem()
}

func (o SetMultiVmSyncStatusPtrOutput) ToSetMultiVmSyncStatusPtrOutput() SetMultiVmSyncStatusPtrOutput {
	return o
}

func (o SetMultiVmSyncStatusPtrOutput) ToSetMultiVmSyncStatusPtrOutputWithContext(ctx context.Context) SetMultiVmSyncStatusPtrOutput {
	return o
}

func (o SetMultiVmSyncStatusPtrOutput) Elem() SetMultiVmSyncStatusOutput {
	return o.ApplyT(func(v *SetMultiVmSyncStatus) SetMultiVmSyncStatus {
		if v != nil {
			return *v
		}
		var ret SetMultiVmSyncStatus
		return ret
	}).(SetMultiVmSyncStatusOutput)
}

func (o SetMultiVmSyncStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SetMultiVmSyncStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SetMultiVmSyncStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SetMultiVmSyncStatusInput interface {
	pulumi.Input

	ToSetMultiVmSyncStatusOutput() SetMultiVmSyncStatusOutput
	ToSetMultiVmSyncStatusOutputWithContext(context.Context) SetMultiVmSyncStatusOutput
}

var setMultiVmSyncStatusPtrType = reflect.TypeOf((**SetMultiVmSyncStatus)(nil)).Elem()

type SetMultiVmSyncStatusPtrInput interface {
	pulumi.Input

	ToSetMultiVmSyncStatusPtrOutput() SetMultiVmSyncStatusPtrOutput
	ToSetMultiVmSyncStatusPtrOutputWithContext(context.Context) SetMultiVmSyncStatusPtrOutput
}

type setMultiVmSyncStatusPtr string

func SetMultiVmSyncStatusPtr(v string) SetMultiVmSyncStatusPtrInput {
	return (*setMultiVmSyncStatusPtr)(&v)
}

func (*setMultiVmSyncStatusPtr) ElementType() reflect.Type {
	return setMultiVmSyncStatusPtrType
}

func (in *setMultiVmSyncStatusPtr) ToSetMultiVmSyncStatusPtrOutput() SetMultiVmSyncStatusPtrOutput {
	return pulumi.ToOutput(in).(SetMultiVmSyncStatusPtrOutput)
}

func (in *setMultiVmSyncStatusPtr) ToSetMultiVmSyncStatusPtrOutputWithContext(ctx context.Context) SetMultiVmSyncStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SetMultiVmSyncStatusPtrOutput)
}

type SqlServerLicenseType string

const (
	SqlServerLicenseTypeNotSpecified  = SqlServerLicenseType("NotSpecified")
	SqlServerLicenseTypeNoLicenseType = SqlServerLicenseType("NoLicenseType")
	SqlServerLicenseTypePAYG          = SqlServerLicenseType("PAYG")
	SqlServerLicenseTypeAHUB          = SqlServerLicenseType("AHUB")
)

func (SqlServerLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerLicenseType)(nil)).Elem()
}

func (e SqlServerLicenseType) ToSqlServerLicenseTypeOutput() SqlServerLicenseTypeOutput {
	return pulumi.ToOutput(e).(SqlServerLicenseTypeOutput)
}

func (e SqlServerLicenseType) ToSqlServerLicenseTypeOutputWithContext(ctx context.Context) SqlServerLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SqlServerLicenseTypeOutput)
}

func (e SqlServerLicenseType) ToSqlServerLicenseTypePtrOutput() SqlServerLicenseTypePtrOutput {
	return e.ToSqlServerLicenseTypePtrOutputWithContext(context.Background())
}

func (e SqlServerLicenseType) ToSqlServerLicenseTypePtrOutputWithContext(ctx context.Context) SqlServerLicenseTypePtrOutput {
	return SqlServerLicenseType(e).ToSqlServerLicenseTypeOutputWithContext(ctx).ToSqlServerLicenseTypePtrOutputWithContext(ctx)
}

func (e SqlServerLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlServerLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlServerLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SqlServerLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SqlServerLicenseTypeOutput struct{ *pulumi.OutputState }

func (SqlServerLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerLicenseType)(nil)).Elem()
}

func (o SqlServerLicenseTypeOutput) ToSqlServerLicenseTypeOutput() SqlServerLicenseTypeOutput {
	return o
}

func (o SqlServerLicenseTypeOutput) ToSqlServerLicenseTypeOutputWithContext(ctx context.Context) SqlServerLicenseTypeOutput {
	return o
}

func (o SqlServerLicenseTypeOutput) ToSqlServerLicenseTypePtrOutput() SqlServerLicenseTypePtrOutput {
	return o.ToSqlServerLicenseTypePtrOutputWithContext(context.Background())
}

func (o SqlServerLicenseTypeOutput) ToSqlServerLicenseTypePtrOutputWithContext(ctx context.Context) SqlServerLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlServerLicenseType) *SqlServerLicenseType {
		return &v
	}).(SqlServerLicenseTypePtrOutput)
}

func (o SqlServerLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SqlServerLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlServerLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SqlServerLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlServerLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlServerLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SqlServerLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (SqlServerLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerLicenseType)(nil)).Elem()
}

func (o SqlServerLicenseTypePtrOutput) ToSqlServerLicenseTypePtrOutput() SqlServerLicenseTypePtrOutput {
	return o
}

func (o SqlServerLicenseTypePtrOutput) ToSqlServerLicenseTypePtrOutputWithContext(ctx context.Context) SqlServerLicenseTypePtrOutput {
	return o
}

func (o SqlServerLicenseTypePtrOutput) Elem() SqlServerLicenseTypeOutput {
	return o.ApplyT(func(v *SqlServerLicenseType) SqlServerLicenseType {
		if v != nil {
			return *v
		}
		var ret SqlServerLicenseType
		return ret
	}).(SqlServerLicenseTypeOutput)
}

func (o SqlServerLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlServerLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SqlServerLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SqlServerLicenseTypeInput interface {
	pulumi.Input

	ToSqlServerLicenseTypeOutput() SqlServerLicenseTypeOutput
	ToSqlServerLicenseTypeOutputWithContext(context.Context) SqlServerLicenseTypeOutput
}

var sqlServerLicenseTypePtrType = reflect.TypeOf((**SqlServerLicenseType)(nil)).Elem()

type SqlServerLicenseTypePtrInput interface {
	pulumi.Input

	ToSqlServerLicenseTypePtrOutput() SqlServerLicenseTypePtrOutput
	ToSqlServerLicenseTypePtrOutputWithContext(context.Context) SqlServerLicenseTypePtrOutput
}

type sqlServerLicenseTypePtr string

func SqlServerLicenseTypePtr(v string) SqlServerLicenseTypePtrInput {
	return (*sqlServerLicenseTypePtr)(&v)
}

func (*sqlServerLicenseTypePtr) ElementType() reflect.Type {
	return sqlServerLicenseTypePtrType
}

func (in *sqlServerLicenseTypePtr) ToSqlServerLicenseTypePtrOutput() SqlServerLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(SqlServerLicenseTypePtrOutput)
}

func (in *sqlServerLicenseTypePtr) ToSqlServerLicenseTypePtrOutputWithContext(ctx context.Context) SqlServerLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SqlServerLicenseTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentAutoUpdateStatusOutput{})
	pulumi.RegisterOutputType(AgentAutoUpdateStatusPtrOutput{})
	pulumi.RegisterOutputType(AutomationAccountAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(AutomationAccountAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(DiskAccountTypeOutput{})
	pulumi.RegisterOutputType(DiskAccountTypePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypeOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypePtrOutput{})
	pulumi.RegisterOutputType(FailoverDeploymentModelOutput{})
	pulumi.RegisterOutputType(FailoverDeploymentModelPtrOutput{})
	pulumi.RegisterOutputType(LicenseTypeOutput{})
	pulumi.RegisterOutputType(LicenseTypePtrOutput{})
	pulumi.RegisterOutputType(PossibleOperationsDirectionsOutput{})
	pulumi.RegisterOutputType(PossibleOperationsDirectionsPtrOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionLocationOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionLocationPtrOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupTypeOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupTypePtrOutput{})
	pulumi.RegisterOutputType(ReplicationProtectedItemOperationOutput{})
	pulumi.RegisterOutputType(ReplicationProtectedItemOperationPtrOutput{})
	pulumi.RegisterOutputType(SetMultiVmSyncStatusOutput{})
	pulumi.RegisterOutputType(SetMultiVmSyncStatusPtrOutput{})
	pulumi.RegisterOutputType(SqlServerLicenseTypeOutput{})
	pulumi.RegisterOutputType(SqlServerLicenseTypePtrOutput{})
}
