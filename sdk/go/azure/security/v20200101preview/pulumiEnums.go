


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthenticationType string

const (
	// AWS cloud account connector user credentials authentication
	AuthenticationTypeAwsCreds = AuthenticationType("awsCreds")
	// AWS account connector assume role authentication
	AuthenticationTypeAwsAssumeRole = AuthenticationType("awsAssumeRole")
	// GCP account connector service to service authentication
	AuthenticationTypeGcpCredentials = AuthenticationType("gcpCredentials")
)

func (AuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationType)(nil)).Elem()
}

func (e AuthenticationType) ToAuthenticationTypeOutput() AuthenticationTypeOutput {
	return pulumi.ToOutput(e).(AuthenticationTypeOutput)
}

func (e AuthenticationType) ToAuthenticationTypeOutputWithContext(ctx context.Context) AuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthenticationTypeOutput)
}

func (e AuthenticationType) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return e.ToAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e AuthenticationType) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return AuthenticationType(e).ToAuthenticationTypeOutputWithContext(ctx).ToAuthenticationTypePtrOutputWithContext(ctx)
}

func (e AuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthenticationTypeOutput struct{ *pulumi.OutputState }

func (AuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationType)(nil)).Elem()
}

func (o AuthenticationTypeOutput) ToAuthenticationTypeOutput() AuthenticationTypeOutput {
	return o
}

func (o AuthenticationTypeOutput) ToAuthenticationTypeOutputWithContext(ctx context.Context) AuthenticationTypeOutput {
	return o
}

func (o AuthenticationTypeOutput) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return o.ToAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationType) *AuthenticationType {
		return &v
	}).(AuthenticationTypePtrOutput)
}

func (o AuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (AuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationType)(nil)).Elem()
}

func (o AuthenticationTypePtrOutput) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return o
}

func (o AuthenticationTypePtrOutput) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return o
}

func (o AuthenticationTypePtrOutput) Elem() AuthenticationTypeOutput {
	return o.ApplyT(func(v *AuthenticationType) AuthenticationType {
		if v != nil {
			return *v
		}
		var ret AuthenticationType
		return ret
	}).(AuthenticationTypeOutput)
}

func (o AuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuthenticationTypeInput interface {
	pulumi.Input

	ToAuthenticationTypeOutput() AuthenticationTypeOutput
	ToAuthenticationTypeOutputWithContext(context.Context) AuthenticationTypeOutput
}

var authenticationTypePtrType = reflect.TypeOf((**AuthenticationType)(nil)).Elem()

type AuthenticationTypePtrInput interface {
	pulumi.Input

	ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput
	ToAuthenticationTypePtrOutputWithContext(context.Context) AuthenticationTypePtrOutput
}

type authenticationTypePtr string

func AuthenticationTypePtr(v string) AuthenticationTypePtrInput {
	return (*authenticationTypePtr)(&v)
}

func (*authenticationTypePtr) ElementType() reflect.Type {
	return authenticationTypePtrType
}

func (in *authenticationTypePtr) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(AuthenticationTypePtrOutput)
}

func (in *authenticationTypePtr) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthenticationTypePtrOutput)
}

type AutoProvision string

const (
	// Install missing Azure Arc agents on machines automatically
	AutoProvisionOn = AutoProvision("On")
	// Do not install Azure Arc agent on the machines automatically
	AutoProvisionOff = AutoProvision("Off")
)

func (AutoProvision) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoProvision)(nil)).Elem()
}

func (e AutoProvision) ToAutoProvisionOutput() AutoProvisionOutput {
	return pulumi.ToOutput(e).(AutoProvisionOutput)
}

func (e AutoProvision) ToAutoProvisionOutputWithContext(ctx context.Context) AutoProvisionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoProvisionOutput)
}

func (e AutoProvision) ToAutoProvisionPtrOutput() AutoProvisionPtrOutput {
	return e.ToAutoProvisionPtrOutputWithContext(context.Background())
}

func (e AutoProvision) ToAutoProvisionPtrOutputWithContext(ctx context.Context) AutoProvisionPtrOutput {
	return AutoProvision(e).ToAutoProvisionOutputWithContext(ctx).ToAutoProvisionPtrOutputWithContext(ctx)
}

func (e AutoProvision) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoProvision) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoProvision) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoProvision) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoProvisionOutput struct{ *pulumi.OutputState }

func (AutoProvisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoProvision)(nil)).Elem()
}

func (o AutoProvisionOutput) ToAutoProvisionOutput() AutoProvisionOutput {
	return o
}

func (o AutoProvisionOutput) ToAutoProvisionOutputWithContext(ctx context.Context) AutoProvisionOutput {
	return o
}

func (o AutoProvisionOutput) ToAutoProvisionPtrOutput() AutoProvisionPtrOutput {
	return o.ToAutoProvisionPtrOutputWithContext(context.Background())
}

func (o AutoProvisionOutput) ToAutoProvisionPtrOutputWithContext(ctx context.Context) AutoProvisionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoProvision) *AutoProvision {
		return &v
	}).(AutoProvisionPtrOutput)
}

func (o AutoProvisionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoProvisionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoProvision) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoProvisionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoProvisionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoProvision) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoProvisionPtrOutput struct{ *pulumi.OutputState }

func (AutoProvisionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoProvision)(nil)).Elem()
}

func (o AutoProvisionPtrOutput) ToAutoProvisionPtrOutput() AutoProvisionPtrOutput {
	return o
}

func (o AutoProvisionPtrOutput) ToAutoProvisionPtrOutputWithContext(ctx context.Context) AutoProvisionPtrOutput {
	return o
}

func (o AutoProvisionPtrOutput) Elem() AutoProvisionOutput {
	return o.ApplyT(func(v *AutoProvision) AutoProvision {
		if v != nil {
			return *v
		}
		var ret AutoProvision
		return ret
	}).(AutoProvisionOutput)
}

func (o AutoProvisionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoProvisionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoProvision) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoProvisionInput interface {
	pulumi.Input

	ToAutoProvisionOutput() AutoProvisionOutput
	ToAutoProvisionOutputWithContext(context.Context) AutoProvisionOutput
}

var autoProvisionPtrType = reflect.TypeOf((**AutoProvision)(nil)).Elem()

type AutoProvisionPtrInput interface {
	pulumi.Input

	ToAutoProvisionPtrOutput() AutoProvisionPtrOutput
	ToAutoProvisionPtrOutputWithContext(context.Context) AutoProvisionPtrOutput
}

type autoProvisionPtr string

func AutoProvisionPtr(v string) AutoProvisionPtrInput {
	return (*autoProvisionPtr)(&v)
}

func (*autoProvisionPtr) ElementType() reflect.Type {
	return autoProvisionPtrType
}

func (in *autoProvisionPtr) ToAutoProvisionPtrOutput() AutoProvisionPtrOutput {
	return pulumi.ToOutput(in).(AutoProvisionPtrOutput)
}

func (in *autoProvisionPtr) ToAutoProvisionPtrOutputWithContext(ctx context.Context) AutoProvisionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoProvisionPtrOutput)
}

type MinimalSeverity string

const (
	// Get notifications on new alerts with High severity
	MinimalSeverityHigh = MinimalSeverity("High")
	// Get notifications on new alerts with medium or high severity
	MinimalSeverityMedium = MinimalSeverity("Medium")
	// Don't get notifications on new alerts with low, medium or high severity
	MinimalSeverityLow = MinimalSeverity("Low")
)

func (MinimalSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimalSeverity)(nil)).Elem()
}

func (e MinimalSeverity) ToMinimalSeverityOutput() MinimalSeverityOutput {
	return pulumi.ToOutput(e).(MinimalSeverityOutput)
}

func (e MinimalSeverity) ToMinimalSeverityOutputWithContext(ctx context.Context) MinimalSeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MinimalSeverityOutput)
}

func (e MinimalSeverity) ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput {
	return e.ToMinimalSeverityPtrOutputWithContext(context.Background())
}

func (e MinimalSeverity) ToMinimalSeverityPtrOutputWithContext(ctx context.Context) MinimalSeverityPtrOutput {
	return MinimalSeverity(e).ToMinimalSeverityOutputWithContext(ctx).ToMinimalSeverityPtrOutputWithContext(ctx)
}

func (e MinimalSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimalSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimalSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MinimalSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MinimalSeverityOutput struct{ *pulumi.OutputState }

func (MinimalSeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimalSeverity)(nil)).Elem()
}

func (o MinimalSeverityOutput) ToMinimalSeverityOutput() MinimalSeverityOutput {
	return o
}

func (o MinimalSeverityOutput) ToMinimalSeverityOutputWithContext(ctx context.Context) MinimalSeverityOutput {
	return o
}

func (o MinimalSeverityOutput) ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput {
	return o.ToMinimalSeverityPtrOutputWithContext(context.Background())
}

func (o MinimalSeverityOutput) ToMinimalSeverityPtrOutputWithContext(ctx context.Context) MinimalSeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MinimalSeverity) *MinimalSeverity {
		return &v
	}).(MinimalSeverityPtrOutput)
}

func (o MinimalSeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MinimalSeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimalSeverity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MinimalSeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimalSeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimalSeverity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MinimalSeverityPtrOutput struct{ *pulumi.OutputState }

func (MinimalSeverityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MinimalSeverity)(nil)).Elem()
}

func (o MinimalSeverityPtrOutput) ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput {
	return o
}

func (o MinimalSeverityPtrOutput) ToMinimalSeverityPtrOutputWithContext(ctx context.Context) MinimalSeverityPtrOutput {
	return o
}

func (o MinimalSeverityPtrOutput) Elem() MinimalSeverityOutput {
	return o.ApplyT(func(v *MinimalSeverity) MinimalSeverity {
		if v != nil {
			return *v
		}
		var ret MinimalSeverity
		return ret
	}).(MinimalSeverityOutput)
}

func (o MinimalSeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimalSeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MinimalSeverity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MinimalSeverityInput interface {
	pulumi.Input

	ToMinimalSeverityOutput() MinimalSeverityOutput
	ToMinimalSeverityOutputWithContext(context.Context) MinimalSeverityOutput
}

var minimalSeverityPtrType = reflect.TypeOf((**MinimalSeverity)(nil)).Elem()

type MinimalSeverityPtrInput interface {
	pulumi.Input

	ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput
	ToMinimalSeverityPtrOutputWithContext(context.Context) MinimalSeverityPtrOutput
}

type minimalSeverityPtr string

func MinimalSeverityPtr(v string) MinimalSeverityPtrInput {
	return (*minimalSeverityPtr)(&v)
}

func (*minimalSeverityPtr) ElementType() reflect.Type {
	return minimalSeverityPtrType
}

func (in *minimalSeverityPtr) ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput {
	return pulumi.ToOutput(in).(MinimalSeverityPtrOutput)
}

func (in *minimalSeverityPtr) ToMinimalSeverityPtrOutputWithContext(ctx context.Context) MinimalSeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MinimalSeverityPtrOutput)
}

type Roles string

const (
	// If enabled, send notification on new alerts to the account admins
	RolesAccountAdmin = Roles("AccountAdmin")
	// If enabled, send notification on new alerts to the service admins
	RolesServiceAdmin = Roles("ServiceAdmin")
	// If enabled, send notification on new alerts to the subscription owners
	RolesOwner = Roles("Owner")
	// If enabled, send notification on new alerts to the subscription contributors
	RolesContributor = Roles("Contributor")
)

func (Roles) ElementType() reflect.Type {
	return reflect.TypeOf((*Roles)(nil)).Elem()
}

func (e Roles) ToRolesOutput() RolesOutput {
	return pulumi.ToOutput(e).(RolesOutput)
}

func (e Roles) ToRolesOutputWithContext(ctx context.Context) RolesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RolesOutput)
}

func (e Roles) ToRolesPtrOutput() RolesPtrOutput {
	return e.ToRolesPtrOutputWithContext(context.Background())
}

func (e Roles) ToRolesPtrOutputWithContext(ctx context.Context) RolesPtrOutput {
	return Roles(e).ToRolesOutputWithContext(ctx).ToRolesPtrOutputWithContext(ctx)
}

func (e Roles) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Roles) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Roles) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Roles) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RolesOutput struct{ *pulumi.OutputState }

func (RolesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Roles)(nil)).Elem()
}

func (o RolesOutput) ToRolesOutput() RolesOutput {
	return o
}

func (o RolesOutput) ToRolesOutputWithContext(ctx context.Context) RolesOutput {
	return o
}

func (o RolesOutput) ToRolesPtrOutput() RolesPtrOutput {
	return o.ToRolesPtrOutputWithContext(context.Background())
}

func (o RolesOutput) ToRolesPtrOutputWithContext(ctx context.Context) RolesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Roles) *Roles {
		return &v
	}).(RolesPtrOutput)
}

func (o RolesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RolesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Roles) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RolesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RolesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Roles) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RolesPtrOutput struct{ *pulumi.OutputState }

func (RolesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Roles)(nil)).Elem()
}

func (o RolesPtrOutput) ToRolesPtrOutput() RolesPtrOutput {
	return o
}

func (o RolesPtrOutput) ToRolesPtrOutputWithContext(ctx context.Context) RolesPtrOutput {
	return o
}

func (o RolesPtrOutput) Elem() RolesOutput {
	return o.ApplyT(func(v *Roles) Roles {
		if v != nil {
			return *v
		}
		var ret Roles
		return ret
	}).(RolesOutput)
}

func (o RolesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RolesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Roles) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RolesInput interface {
	pulumi.Input

	ToRolesOutput() RolesOutput
	ToRolesOutputWithContext(context.Context) RolesOutput
}

var rolesPtrType = reflect.TypeOf((**Roles)(nil)).Elem()

type RolesPtrInput interface {
	pulumi.Input

	ToRolesPtrOutput() RolesPtrOutput
	ToRolesPtrOutputWithContext(context.Context) RolesPtrOutput
}

type rolesPtr string

func RolesPtr(v string) RolesPtrInput {
	return (*rolesPtr)(&v)
}

func (*rolesPtr) ElementType() reflect.Type {
	return rolesPtrType
}

func (in *rolesPtr) ToRolesPtrOutput() RolesPtrOutput {
	return pulumi.ToOutput(in).(RolesPtrOutput)
}

func (in *rolesPtr) ToRolesPtrOutputWithContext(ctx context.Context) RolesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RolesPtrOutput)
}

type State string

const (
	// Send notification on new alerts to the subscription's admins
	StateOn = State("On")
	// Don't send notification on new alerts to the subscription's admins
	StateOff = State("Off")
)

func (State) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (e State) ToStateOutput() StateOutput {
	return pulumi.ToOutput(e).(StateOutput)
}

func (e State) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StateOutput)
}

func (e State) ToStatePtrOutput() StatePtrOutput {
	return e.ToStatePtrOutputWithContext(context.Background())
}

func (e State) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return State(e).ToStateOutputWithContext(ctx).ToStatePtrOutputWithContext(ctx)
}

func (e State) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e State) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StateOutput struct{ *pulumi.OutputState }

func (StateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (o StateOutput) ToStateOutput() StateOutput {
	return o
}

func (o StateOutput) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return o
}

func (o StateOutput) ToStatePtrOutput() StatePtrOutput {
	return o.ToStatePtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v State) *State {
		return &v
	}).(StatePtrOutput)
}

func (o StateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatePtrOutput struct{ *pulumi.OutputState }

func (StatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**State)(nil)).Elem()
}

func (o StatePtrOutput) ToStatePtrOutput() StatePtrOutput {
	return o
}

func (o StatePtrOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o
}

func (o StatePtrOutput) Elem() StateOutput {
	return o.ApplyT(func(v *State) State {
		if v != nil {
			return *v
		}
		var ret State
		return ret
	}).(StateOutput)
}

func (o StatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *State) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StateInput interface {
	pulumi.Input

	ToStateOutput() StateOutput
	ToStateOutputWithContext(context.Context) StateOutput
}

var statePtrType = reflect.TypeOf((**State)(nil)).Elem()

type StatePtrInput interface {
	pulumi.Input

	ToStatePtrOutput() StatePtrOutput
	ToStatePtrOutputWithContext(context.Context) StatePtrOutput
}

type statePtr string

func StatePtr(v string) StatePtrInput {
	return (*statePtr)(&v)
}

func (*statePtr) ElementType() reflect.Type {
	return statePtrType
}

func (in *statePtr) ToStatePtrOutput() StatePtrOutput {
	return pulumi.ToOutput(in).(StatePtrOutput)
}

func (in *statePtr) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthenticationTypeOutput{})
	pulumi.RegisterOutputType(AuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(AutoProvisionOutput{})
	pulumi.RegisterOutputType(AutoProvisionPtrOutput{})
	pulumi.RegisterOutputType(MinimalSeverityOutput{})
	pulumi.RegisterOutputType(MinimalSeverityPtrOutput{})
	pulumi.RegisterOutputType(RolesOutput{})
	pulumi.RegisterOutputType(RolesPtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
}
