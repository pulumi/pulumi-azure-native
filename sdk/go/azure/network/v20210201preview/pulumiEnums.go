


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddressPrefixType string

const (
	AddressPrefixTypeIPPrefix   = AddressPrefixType("IPPrefix")
	AddressPrefixTypeServiceTag = AddressPrefixType("ServiceTag")
)

func (AddressPrefixType) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressPrefixType)(nil)).Elem()
}

func (e AddressPrefixType) ToAddressPrefixTypeOutput() AddressPrefixTypeOutput {
	return pulumi.ToOutput(e).(AddressPrefixTypeOutput)
}

func (e AddressPrefixType) ToAddressPrefixTypeOutputWithContext(ctx context.Context) AddressPrefixTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AddressPrefixTypeOutput)
}

func (e AddressPrefixType) ToAddressPrefixTypePtrOutput() AddressPrefixTypePtrOutput {
	return e.ToAddressPrefixTypePtrOutputWithContext(context.Background())
}

func (e AddressPrefixType) ToAddressPrefixTypePtrOutputWithContext(ctx context.Context) AddressPrefixTypePtrOutput {
	return AddressPrefixType(e).ToAddressPrefixTypeOutputWithContext(ctx).ToAddressPrefixTypePtrOutputWithContext(ctx)
}

func (e AddressPrefixType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddressPrefixType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddressPrefixType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AddressPrefixType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AddressPrefixTypeOutput struct{ *pulumi.OutputState }

func (AddressPrefixTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressPrefixType)(nil)).Elem()
}

func (o AddressPrefixTypeOutput) ToAddressPrefixTypeOutput() AddressPrefixTypeOutput {
	return o
}

func (o AddressPrefixTypeOutput) ToAddressPrefixTypeOutputWithContext(ctx context.Context) AddressPrefixTypeOutput {
	return o
}

func (o AddressPrefixTypeOutput) ToAddressPrefixTypePtrOutput() AddressPrefixTypePtrOutput {
	return o.ToAddressPrefixTypePtrOutputWithContext(context.Background())
}

func (o AddressPrefixTypeOutput) ToAddressPrefixTypePtrOutputWithContext(ctx context.Context) AddressPrefixTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressPrefixType) *AddressPrefixType {
		return &v
	}).(AddressPrefixTypePtrOutput)
}

func (o AddressPrefixTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AddressPrefixTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddressPrefixType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AddressPrefixTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddressPrefixTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddressPrefixType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AddressPrefixTypePtrOutput struct{ *pulumi.OutputState }

func (AddressPrefixTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressPrefixType)(nil)).Elem()
}

func (o AddressPrefixTypePtrOutput) ToAddressPrefixTypePtrOutput() AddressPrefixTypePtrOutput {
	return o
}

func (o AddressPrefixTypePtrOutput) ToAddressPrefixTypePtrOutputWithContext(ctx context.Context) AddressPrefixTypePtrOutput {
	return o
}

func (o AddressPrefixTypePtrOutput) Elem() AddressPrefixTypeOutput {
	return o.ApplyT(func(v *AddressPrefixType) AddressPrefixType {
		if v != nil {
			return *v
		}
		var ret AddressPrefixType
		return ret
	}).(AddressPrefixTypeOutput)
}

func (o AddressPrefixTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddressPrefixTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AddressPrefixType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AddressPrefixTypeInput interface {
	pulumi.Input

	ToAddressPrefixTypeOutput() AddressPrefixTypeOutput
	ToAddressPrefixTypeOutputWithContext(context.Context) AddressPrefixTypeOutput
}

var addressPrefixTypePtrType = reflect.TypeOf((**AddressPrefixType)(nil)).Elem()

type AddressPrefixTypePtrInput interface {
	pulumi.Input

	ToAddressPrefixTypePtrOutput() AddressPrefixTypePtrOutput
	ToAddressPrefixTypePtrOutputWithContext(context.Context) AddressPrefixTypePtrOutput
}

type addressPrefixTypePtr string

func AddressPrefixTypePtr(v string) AddressPrefixTypePtrInput {
	return (*addressPrefixTypePtr)(&v)
}

func (*addressPrefixTypePtr) ElementType() reflect.Type {
	return addressPrefixTypePtrType
}

func (in *addressPrefixTypePtr) ToAddressPrefixTypePtrOutput() AddressPrefixTypePtrOutput {
	return pulumi.ToOutput(in).(AddressPrefixTypePtrOutput)
}

func (in *addressPrefixTypePtr) ToAddressPrefixTypePtrOutputWithContext(ctx context.Context) AddressPrefixTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AddressPrefixTypePtrOutput)
}

type AdminRuleKind string

const (
	AdminRuleKindCustom  = AdminRuleKind("Custom")
	AdminRuleKindDefault = AdminRuleKind("Default")
)

func (AdminRuleKind) ElementType() reflect.Type {
	return reflect.TypeOf((*AdminRuleKind)(nil)).Elem()
}

func (e AdminRuleKind) ToAdminRuleKindOutput() AdminRuleKindOutput {
	return pulumi.ToOutput(e).(AdminRuleKindOutput)
}

func (e AdminRuleKind) ToAdminRuleKindOutputWithContext(ctx context.Context) AdminRuleKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AdminRuleKindOutput)
}

func (e AdminRuleKind) ToAdminRuleKindPtrOutput() AdminRuleKindPtrOutput {
	return e.ToAdminRuleKindPtrOutputWithContext(context.Background())
}

func (e AdminRuleKind) ToAdminRuleKindPtrOutputWithContext(ctx context.Context) AdminRuleKindPtrOutput {
	return AdminRuleKind(e).ToAdminRuleKindOutputWithContext(ctx).ToAdminRuleKindPtrOutputWithContext(ctx)
}

func (e AdminRuleKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdminRuleKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdminRuleKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AdminRuleKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AdminRuleKindOutput struct{ *pulumi.OutputState }

func (AdminRuleKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdminRuleKind)(nil)).Elem()
}

func (o AdminRuleKindOutput) ToAdminRuleKindOutput() AdminRuleKindOutput {
	return o
}

func (o AdminRuleKindOutput) ToAdminRuleKindOutputWithContext(ctx context.Context) AdminRuleKindOutput {
	return o
}

func (o AdminRuleKindOutput) ToAdminRuleKindPtrOutput() AdminRuleKindPtrOutput {
	return o.ToAdminRuleKindPtrOutputWithContext(context.Background())
}

func (o AdminRuleKindOutput) ToAdminRuleKindPtrOutputWithContext(ctx context.Context) AdminRuleKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdminRuleKind) *AdminRuleKind {
		return &v
	}).(AdminRuleKindPtrOutput)
}

func (o AdminRuleKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AdminRuleKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdminRuleKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AdminRuleKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdminRuleKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdminRuleKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AdminRuleKindPtrOutput struct{ *pulumi.OutputState }

func (AdminRuleKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminRuleKind)(nil)).Elem()
}

func (o AdminRuleKindPtrOutput) ToAdminRuleKindPtrOutput() AdminRuleKindPtrOutput {
	return o
}

func (o AdminRuleKindPtrOutput) ToAdminRuleKindPtrOutputWithContext(ctx context.Context) AdminRuleKindPtrOutput {
	return o
}

func (o AdminRuleKindPtrOutput) Elem() AdminRuleKindOutput {
	return o.ApplyT(func(v *AdminRuleKind) AdminRuleKind {
		if v != nil {
			return *v
		}
		var ret AdminRuleKind
		return ret
	}).(AdminRuleKindOutput)
}

func (o AdminRuleKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdminRuleKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AdminRuleKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AdminRuleKindInput interface {
	pulumi.Input

	ToAdminRuleKindOutput() AdminRuleKindOutput
	ToAdminRuleKindOutputWithContext(context.Context) AdminRuleKindOutput
}

var adminRuleKindPtrType = reflect.TypeOf((**AdminRuleKind)(nil)).Elem()

type AdminRuleKindPtrInput interface {
	pulumi.Input

	ToAdminRuleKindPtrOutput() AdminRuleKindPtrOutput
	ToAdminRuleKindPtrOutputWithContext(context.Context) AdminRuleKindPtrOutput
}

type adminRuleKindPtr string

func AdminRuleKindPtr(v string) AdminRuleKindPtrInput {
	return (*adminRuleKindPtr)(&v)
}

func (*adminRuleKindPtr) ElementType() reflect.Type {
	return adminRuleKindPtrType
}

func (in *adminRuleKindPtr) ToAdminRuleKindPtrOutput() AdminRuleKindPtrOutput {
	return pulumi.ToOutput(in).(AdminRuleKindPtrOutput)
}

func (in *adminRuleKindPtr) ToAdminRuleKindPtrOutputWithContext(ctx context.Context) AdminRuleKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AdminRuleKindPtrOutput)
}

type ConfigurationType string

const (
	ConfigurationTypeSecurityAdmin = ConfigurationType("SecurityAdmin")
	ConfigurationTypeSecurityUser  = ConfigurationType("SecurityUser")
	ConfigurationTypeConnectivity  = ConfigurationType("Connectivity")
)

func (ConfigurationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationType)(nil)).Elem()
}

func (e ConfigurationType) ToConfigurationTypeOutput() ConfigurationTypeOutput {
	return pulumi.ToOutput(e).(ConfigurationTypeOutput)
}

func (e ConfigurationType) ToConfigurationTypeOutputWithContext(ctx context.Context) ConfigurationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfigurationTypeOutput)
}

func (e ConfigurationType) ToConfigurationTypePtrOutput() ConfigurationTypePtrOutput {
	return e.ToConfigurationTypePtrOutputWithContext(context.Background())
}

func (e ConfigurationType) ToConfigurationTypePtrOutputWithContext(ctx context.Context) ConfigurationTypePtrOutput {
	return ConfigurationType(e).ToConfigurationTypeOutputWithContext(ctx).ToConfigurationTypePtrOutputWithContext(ctx)
}

func (e ConfigurationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigurationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigurationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConfigurationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfigurationTypeOutput struct{ *pulumi.OutputState }

func (ConfigurationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationType)(nil)).Elem()
}

func (o ConfigurationTypeOutput) ToConfigurationTypeOutput() ConfigurationTypeOutput {
	return o
}

func (o ConfigurationTypeOutput) ToConfigurationTypeOutputWithContext(ctx context.Context) ConfigurationTypeOutput {
	return o
}

func (o ConfigurationTypeOutput) ToConfigurationTypePtrOutput() ConfigurationTypePtrOutput {
	return o.ToConfigurationTypePtrOutputWithContext(context.Background())
}

func (o ConfigurationTypeOutput) ToConfigurationTypePtrOutputWithContext(ctx context.Context) ConfigurationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationType) *ConfigurationType {
		return &v
	}).(ConfigurationTypePtrOutput)
}

func (o ConfigurationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfigurationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigurationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfigurationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigurationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigurationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfigurationTypePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationType)(nil)).Elem()
}

func (o ConfigurationTypePtrOutput) ToConfigurationTypePtrOutput() ConfigurationTypePtrOutput {
	return o
}

func (o ConfigurationTypePtrOutput) ToConfigurationTypePtrOutputWithContext(ctx context.Context) ConfigurationTypePtrOutput {
	return o
}

func (o ConfigurationTypePtrOutput) Elem() ConfigurationTypeOutput {
	return o.ApplyT(func(v *ConfigurationType) ConfigurationType {
		if v != nil {
			return *v
		}
		var ret ConfigurationType
		return ret
	}).(ConfigurationTypeOutput)
}

func (o ConfigurationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigurationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConfigurationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConfigurationTypeInput interface {
	pulumi.Input

	ToConfigurationTypeOutput() ConfigurationTypeOutput
	ToConfigurationTypeOutputWithContext(context.Context) ConfigurationTypeOutput
}

var configurationTypePtrType = reflect.TypeOf((**ConfigurationType)(nil)).Elem()

type ConfigurationTypePtrInput interface {
	pulumi.Input

	ToConfigurationTypePtrOutput() ConfigurationTypePtrOutput
	ToConfigurationTypePtrOutputWithContext(context.Context) ConfigurationTypePtrOutput
}

type configurationTypePtr string

func ConfigurationTypePtr(v string) ConfigurationTypePtrInput {
	return (*configurationTypePtr)(&v)
}

func (*configurationTypePtr) ElementType() reflect.Type {
	return configurationTypePtrType
}

func (in *configurationTypePtr) ToConfigurationTypePtrOutput() ConfigurationTypePtrOutput {
	return pulumi.ToOutput(in).(ConfigurationTypePtrOutput)
}

func (in *configurationTypePtr) ToConfigurationTypePtrOutputWithContext(ctx context.Context) ConfigurationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfigurationTypePtrOutput)
}

type ConnectivityTopology string

const (
	ConnectivityTopologyHubAndSpoke = ConnectivityTopology("HubAndSpoke")
	ConnectivityTopologyMesh        = ConnectivityTopology("Mesh")
)

func (ConnectivityTopology) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityTopology)(nil)).Elem()
}

func (e ConnectivityTopology) ToConnectivityTopologyOutput() ConnectivityTopologyOutput {
	return pulumi.ToOutput(e).(ConnectivityTopologyOutput)
}

func (e ConnectivityTopology) ToConnectivityTopologyOutputWithContext(ctx context.Context) ConnectivityTopologyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectivityTopologyOutput)
}

func (e ConnectivityTopology) ToConnectivityTopologyPtrOutput() ConnectivityTopologyPtrOutput {
	return e.ToConnectivityTopologyPtrOutputWithContext(context.Background())
}

func (e ConnectivityTopology) ToConnectivityTopologyPtrOutputWithContext(ctx context.Context) ConnectivityTopologyPtrOutput {
	return ConnectivityTopology(e).ToConnectivityTopologyOutputWithContext(ctx).ToConnectivityTopologyPtrOutputWithContext(ctx)
}

func (e ConnectivityTopology) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectivityTopology) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectivityTopology) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectivityTopology) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectivityTopologyOutput struct{ *pulumi.OutputState }

func (ConnectivityTopologyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityTopology)(nil)).Elem()
}

func (o ConnectivityTopologyOutput) ToConnectivityTopologyOutput() ConnectivityTopologyOutput {
	return o
}

func (o ConnectivityTopologyOutput) ToConnectivityTopologyOutputWithContext(ctx context.Context) ConnectivityTopologyOutput {
	return o
}

func (o ConnectivityTopologyOutput) ToConnectivityTopologyPtrOutput() ConnectivityTopologyPtrOutput {
	return o.ToConnectivityTopologyPtrOutputWithContext(context.Background())
}

func (o ConnectivityTopologyOutput) ToConnectivityTopologyPtrOutputWithContext(ctx context.Context) ConnectivityTopologyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectivityTopology) *ConnectivityTopology {
		return &v
	}).(ConnectivityTopologyPtrOutput)
}

func (o ConnectivityTopologyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectivityTopologyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectivityTopology) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectivityTopologyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectivityTopologyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectivityTopology) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectivityTopologyPtrOutput struct{ *pulumi.OutputState }

func (ConnectivityTopologyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectivityTopology)(nil)).Elem()
}

func (o ConnectivityTopologyPtrOutput) ToConnectivityTopologyPtrOutput() ConnectivityTopologyPtrOutput {
	return o
}

func (o ConnectivityTopologyPtrOutput) ToConnectivityTopologyPtrOutputWithContext(ctx context.Context) ConnectivityTopologyPtrOutput {
	return o
}

func (o ConnectivityTopologyPtrOutput) Elem() ConnectivityTopologyOutput {
	return o.ApplyT(func(v *ConnectivityTopology) ConnectivityTopology {
		if v != nil {
			return *v
		}
		var ret ConnectivityTopology
		return ret
	}).(ConnectivityTopologyOutput)
}

func (o ConnectivityTopologyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectivityTopologyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectivityTopology) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectivityTopologyInput interface {
	pulumi.Input

	ToConnectivityTopologyOutput() ConnectivityTopologyOutput
	ToConnectivityTopologyOutputWithContext(context.Context) ConnectivityTopologyOutput
}

var connectivityTopologyPtrType = reflect.TypeOf((**ConnectivityTopology)(nil)).Elem()

type ConnectivityTopologyPtrInput interface {
	pulumi.Input

	ToConnectivityTopologyPtrOutput() ConnectivityTopologyPtrOutput
	ToConnectivityTopologyPtrOutputWithContext(context.Context) ConnectivityTopologyPtrOutput
}

type connectivityTopologyPtr string

func ConnectivityTopologyPtr(v string) ConnectivityTopologyPtrInput {
	return (*connectivityTopologyPtr)(&v)
}

func (*connectivityTopologyPtr) ElementType() reflect.Type {
	return connectivityTopologyPtrType
}

func (in *connectivityTopologyPtr) ToConnectivityTopologyPtrOutput() ConnectivityTopologyPtrOutput {
	return pulumi.ToOutput(in).(ConnectivityTopologyPtrOutput)
}

func (in *connectivityTopologyPtr) ToConnectivityTopologyPtrOutputWithContext(ctx context.Context) ConnectivityTopologyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectivityTopologyPtrOutput)
}

type DeleteExistingNSGs string

const (
	DeleteExistingNSGsFalse = DeleteExistingNSGs("False")
	DeleteExistingNSGsTrue  = DeleteExistingNSGs("True")
)

func (DeleteExistingNSGs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteExistingNSGs)(nil)).Elem()
}

func (e DeleteExistingNSGs) ToDeleteExistingNSGsOutput() DeleteExistingNSGsOutput {
	return pulumi.ToOutput(e).(DeleteExistingNSGsOutput)
}

func (e DeleteExistingNSGs) ToDeleteExistingNSGsOutputWithContext(ctx context.Context) DeleteExistingNSGsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeleteExistingNSGsOutput)
}

func (e DeleteExistingNSGs) ToDeleteExistingNSGsPtrOutput() DeleteExistingNSGsPtrOutput {
	return e.ToDeleteExistingNSGsPtrOutputWithContext(context.Background())
}

func (e DeleteExistingNSGs) ToDeleteExistingNSGsPtrOutputWithContext(ctx context.Context) DeleteExistingNSGsPtrOutput {
	return DeleteExistingNSGs(e).ToDeleteExistingNSGsOutputWithContext(ctx).ToDeleteExistingNSGsPtrOutputWithContext(ctx)
}

func (e DeleteExistingNSGs) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeleteExistingNSGs) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeleteExistingNSGs) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeleteExistingNSGs) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeleteExistingNSGsOutput struct{ *pulumi.OutputState }

func (DeleteExistingNSGsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteExistingNSGs)(nil)).Elem()
}

func (o DeleteExistingNSGsOutput) ToDeleteExistingNSGsOutput() DeleteExistingNSGsOutput {
	return o
}

func (o DeleteExistingNSGsOutput) ToDeleteExistingNSGsOutputWithContext(ctx context.Context) DeleteExistingNSGsOutput {
	return o
}

func (o DeleteExistingNSGsOutput) ToDeleteExistingNSGsPtrOutput() DeleteExistingNSGsPtrOutput {
	return o.ToDeleteExistingNSGsPtrOutputWithContext(context.Background())
}

func (o DeleteExistingNSGsOutput) ToDeleteExistingNSGsPtrOutputWithContext(ctx context.Context) DeleteExistingNSGsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeleteExistingNSGs) *DeleteExistingNSGs {
		return &v
	}).(DeleteExistingNSGsPtrOutput)
}

func (o DeleteExistingNSGsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeleteExistingNSGsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeleteExistingNSGs) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeleteExistingNSGsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeleteExistingNSGsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeleteExistingNSGs) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeleteExistingNSGsPtrOutput struct{ *pulumi.OutputState }

func (DeleteExistingNSGsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteExistingNSGs)(nil)).Elem()
}

func (o DeleteExistingNSGsPtrOutput) ToDeleteExistingNSGsPtrOutput() DeleteExistingNSGsPtrOutput {
	return o
}

func (o DeleteExistingNSGsPtrOutput) ToDeleteExistingNSGsPtrOutputWithContext(ctx context.Context) DeleteExistingNSGsPtrOutput {
	return o
}

func (o DeleteExistingNSGsPtrOutput) Elem() DeleteExistingNSGsOutput {
	return o.ApplyT(func(v *DeleteExistingNSGs) DeleteExistingNSGs {
		if v != nil {
			return *v
		}
		var ret DeleteExistingNSGs
		return ret
	}).(DeleteExistingNSGsOutput)
}

func (o DeleteExistingNSGsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeleteExistingNSGsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeleteExistingNSGs) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeleteExistingNSGsInput interface {
	pulumi.Input

	ToDeleteExistingNSGsOutput() DeleteExistingNSGsOutput
	ToDeleteExistingNSGsOutputWithContext(context.Context) DeleteExistingNSGsOutput
}

var deleteExistingNSGsPtrType = reflect.TypeOf((**DeleteExistingNSGs)(nil)).Elem()

type DeleteExistingNSGsPtrInput interface {
	pulumi.Input

	ToDeleteExistingNSGsPtrOutput() DeleteExistingNSGsPtrOutput
	ToDeleteExistingNSGsPtrOutputWithContext(context.Context) DeleteExistingNSGsPtrOutput
}

type deleteExistingNSGsPtr string

func DeleteExistingNSGsPtr(v string) DeleteExistingNSGsPtrInput {
	return (*deleteExistingNSGsPtr)(&v)
}

func (*deleteExistingNSGsPtr) ElementType() reflect.Type {
	return deleteExistingNSGsPtrType
}

func (in *deleteExistingNSGsPtr) ToDeleteExistingNSGsPtrOutput() DeleteExistingNSGsPtrOutput {
	return pulumi.ToOutput(in).(DeleteExistingNSGsPtrOutput)
}

func (in *deleteExistingNSGsPtr) ToDeleteExistingNSGsPtrOutputWithContext(ctx context.Context) DeleteExistingNSGsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeleteExistingNSGsPtrOutput)
}

type DeleteExistingPeering string

const (
	DeleteExistingPeeringFalse = DeleteExistingPeering("False")
	DeleteExistingPeeringTrue  = DeleteExistingPeering("True")
)

func (DeleteExistingPeering) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteExistingPeering)(nil)).Elem()
}

func (e DeleteExistingPeering) ToDeleteExistingPeeringOutput() DeleteExistingPeeringOutput {
	return pulumi.ToOutput(e).(DeleteExistingPeeringOutput)
}

func (e DeleteExistingPeering) ToDeleteExistingPeeringOutputWithContext(ctx context.Context) DeleteExistingPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeleteExistingPeeringOutput)
}

func (e DeleteExistingPeering) ToDeleteExistingPeeringPtrOutput() DeleteExistingPeeringPtrOutput {
	return e.ToDeleteExistingPeeringPtrOutputWithContext(context.Background())
}

func (e DeleteExistingPeering) ToDeleteExistingPeeringPtrOutputWithContext(ctx context.Context) DeleteExistingPeeringPtrOutput {
	return DeleteExistingPeering(e).ToDeleteExistingPeeringOutputWithContext(ctx).ToDeleteExistingPeeringPtrOutputWithContext(ctx)
}

func (e DeleteExistingPeering) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeleteExistingPeering) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeleteExistingPeering) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeleteExistingPeering) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeleteExistingPeeringOutput struct{ *pulumi.OutputState }

func (DeleteExistingPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteExistingPeering)(nil)).Elem()
}

func (o DeleteExistingPeeringOutput) ToDeleteExistingPeeringOutput() DeleteExistingPeeringOutput {
	return o
}

func (o DeleteExistingPeeringOutput) ToDeleteExistingPeeringOutputWithContext(ctx context.Context) DeleteExistingPeeringOutput {
	return o
}

func (o DeleteExistingPeeringOutput) ToDeleteExistingPeeringPtrOutput() DeleteExistingPeeringPtrOutput {
	return o.ToDeleteExistingPeeringPtrOutputWithContext(context.Background())
}

func (o DeleteExistingPeeringOutput) ToDeleteExistingPeeringPtrOutputWithContext(ctx context.Context) DeleteExistingPeeringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeleteExistingPeering) *DeleteExistingPeering {
		return &v
	}).(DeleteExistingPeeringPtrOutput)
}

func (o DeleteExistingPeeringOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeleteExistingPeeringOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeleteExistingPeering) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeleteExistingPeeringOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeleteExistingPeeringOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeleteExistingPeering) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeleteExistingPeeringPtrOutput struct{ *pulumi.OutputState }

func (DeleteExistingPeeringPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteExistingPeering)(nil)).Elem()
}

func (o DeleteExistingPeeringPtrOutput) ToDeleteExistingPeeringPtrOutput() DeleteExistingPeeringPtrOutput {
	return o
}

func (o DeleteExistingPeeringPtrOutput) ToDeleteExistingPeeringPtrOutputWithContext(ctx context.Context) DeleteExistingPeeringPtrOutput {
	return o
}

func (o DeleteExistingPeeringPtrOutput) Elem() DeleteExistingPeeringOutput {
	return o.ApplyT(func(v *DeleteExistingPeering) DeleteExistingPeering {
		if v != nil {
			return *v
		}
		var ret DeleteExistingPeering
		return ret
	}).(DeleteExistingPeeringOutput)
}

func (o DeleteExistingPeeringPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeleteExistingPeeringPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeleteExistingPeering) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeleteExistingPeeringInput interface {
	pulumi.Input

	ToDeleteExistingPeeringOutput() DeleteExistingPeeringOutput
	ToDeleteExistingPeeringOutputWithContext(context.Context) DeleteExistingPeeringOutput
}

var deleteExistingPeeringPtrType = reflect.TypeOf((**DeleteExistingPeering)(nil)).Elem()

type DeleteExistingPeeringPtrInput interface {
	pulumi.Input

	ToDeleteExistingPeeringPtrOutput() DeleteExistingPeeringPtrOutput
	ToDeleteExistingPeeringPtrOutputWithContext(context.Context) DeleteExistingPeeringPtrOutput
}

type deleteExistingPeeringPtr string

func DeleteExistingPeeringPtr(v string) DeleteExistingPeeringPtrInput {
	return (*deleteExistingPeeringPtr)(&v)
}

func (*deleteExistingPeeringPtr) ElementType() reflect.Type {
	return deleteExistingPeeringPtrType
}

func (in *deleteExistingPeeringPtr) ToDeleteExistingPeeringPtrOutput() DeleteExistingPeeringPtrOutput {
	return pulumi.ToOutput(in).(DeleteExistingPeeringPtrOutput)
}

func (in *deleteExistingPeeringPtr) ToDeleteExistingPeeringPtrOutputWithContext(ctx context.Context) DeleteExistingPeeringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeleteExistingPeeringPtrOutput)
}

type GroupConnectivity string

const (
	GroupConnectivityNone              = GroupConnectivity("None")
	GroupConnectivityDirectlyConnected = GroupConnectivity("DirectlyConnected")
)

func (GroupConnectivity) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConnectivity)(nil)).Elem()
}

func (e GroupConnectivity) ToGroupConnectivityOutput() GroupConnectivityOutput {
	return pulumi.ToOutput(e).(GroupConnectivityOutput)
}

func (e GroupConnectivity) ToGroupConnectivityOutputWithContext(ctx context.Context) GroupConnectivityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GroupConnectivityOutput)
}

func (e GroupConnectivity) ToGroupConnectivityPtrOutput() GroupConnectivityPtrOutput {
	return e.ToGroupConnectivityPtrOutputWithContext(context.Background())
}

func (e GroupConnectivity) ToGroupConnectivityPtrOutputWithContext(ctx context.Context) GroupConnectivityPtrOutput {
	return GroupConnectivity(e).ToGroupConnectivityOutputWithContext(ctx).ToGroupConnectivityPtrOutputWithContext(ctx)
}

func (e GroupConnectivity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GroupConnectivity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GroupConnectivity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GroupConnectivity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GroupConnectivityOutput struct{ *pulumi.OutputState }

func (GroupConnectivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupConnectivity)(nil)).Elem()
}

func (o GroupConnectivityOutput) ToGroupConnectivityOutput() GroupConnectivityOutput {
	return o
}

func (o GroupConnectivityOutput) ToGroupConnectivityOutputWithContext(ctx context.Context) GroupConnectivityOutput {
	return o
}

func (o GroupConnectivityOutput) ToGroupConnectivityPtrOutput() GroupConnectivityPtrOutput {
	return o.ToGroupConnectivityPtrOutputWithContext(context.Background())
}

func (o GroupConnectivityOutput) ToGroupConnectivityPtrOutputWithContext(ctx context.Context) GroupConnectivityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupConnectivity) *GroupConnectivity {
		return &v
	}).(GroupConnectivityPtrOutput)
}

func (o GroupConnectivityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GroupConnectivityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GroupConnectivity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GroupConnectivityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GroupConnectivityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GroupConnectivity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GroupConnectivityPtrOutput struct{ *pulumi.OutputState }

func (GroupConnectivityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupConnectivity)(nil)).Elem()
}

func (o GroupConnectivityPtrOutput) ToGroupConnectivityPtrOutput() GroupConnectivityPtrOutput {
	return o
}

func (o GroupConnectivityPtrOutput) ToGroupConnectivityPtrOutputWithContext(ctx context.Context) GroupConnectivityPtrOutput {
	return o
}

func (o GroupConnectivityPtrOutput) Elem() GroupConnectivityOutput {
	return o.ApplyT(func(v *GroupConnectivity) GroupConnectivity {
		if v != nil {
			return *v
		}
		var ret GroupConnectivity
		return ret
	}).(GroupConnectivityOutput)
}

func (o GroupConnectivityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GroupConnectivityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GroupConnectivity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GroupConnectivityInput interface {
	pulumi.Input

	ToGroupConnectivityOutput() GroupConnectivityOutput
	ToGroupConnectivityOutputWithContext(context.Context) GroupConnectivityOutput
}

var groupConnectivityPtrType = reflect.TypeOf((**GroupConnectivity)(nil)).Elem()

type GroupConnectivityPtrInput interface {
	pulumi.Input

	ToGroupConnectivityPtrOutput() GroupConnectivityPtrOutput
	ToGroupConnectivityPtrOutputWithContext(context.Context) GroupConnectivityPtrOutput
}

type groupConnectivityPtr string

func GroupConnectivityPtr(v string) GroupConnectivityPtrInput {
	return (*groupConnectivityPtr)(&v)
}

func (*groupConnectivityPtr) ElementType() reflect.Type {
	return groupConnectivityPtrType
}

func (in *groupConnectivityPtr) ToGroupConnectivityPtrOutput() GroupConnectivityPtrOutput {
	return pulumi.ToOutput(in).(GroupConnectivityPtrOutput)
}

func (in *groupConnectivityPtr) ToGroupConnectivityPtrOutputWithContext(ctx context.Context) GroupConnectivityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GroupConnectivityPtrOutput)
}

type IsGlobal string

const (
	IsGlobalFalse = IsGlobal("False")
	IsGlobalTrue  = IsGlobal("True")
)

func (IsGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((*IsGlobal)(nil)).Elem()
}

func (e IsGlobal) ToIsGlobalOutput() IsGlobalOutput {
	return pulumi.ToOutput(e).(IsGlobalOutput)
}

func (e IsGlobal) ToIsGlobalOutputWithContext(ctx context.Context) IsGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IsGlobalOutput)
}

func (e IsGlobal) ToIsGlobalPtrOutput() IsGlobalPtrOutput {
	return e.ToIsGlobalPtrOutputWithContext(context.Background())
}

func (e IsGlobal) ToIsGlobalPtrOutputWithContext(ctx context.Context) IsGlobalPtrOutput {
	return IsGlobal(e).ToIsGlobalOutputWithContext(ctx).ToIsGlobalPtrOutputWithContext(ctx)
}

func (e IsGlobal) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IsGlobal) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IsGlobal) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IsGlobal) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IsGlobalOutput struct{ *pulumi.OutputState }

func (IsGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IsGlobal)(nil)).Elem()
}

func (o IsGlobalOutput) ToIsGlobalOutput() IsGlobalOutput {
	return o
}

func (o IsGlobalOutput) ToIsGlobalOutputWithContext(ctx context.Context) IsGlobalOutput {
	return o
}

func (o IsGlobalOutput) ToIsGlobalPtrOutput() IsGlobalPtrOutput {
	return o.ToIsGlobalPtrOutputWithContext(context.Background())
}

func (o IsGlobalOutput) ToIsGlobalPtrOutputWithContext(ctx context.Context) IsGlobalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IsGlobal) *IsGlobal {
		return &v
	}).(IsGlobalPtrOutput)
}

func (o IsGlobalOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IsGlobalOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IsGlobal) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IsGlobalOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IsGlobalOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IsGlobal) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IsGlobalPtrOutput struct{ *pulumi.OutputState }

func (IsGlobalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IsGlobal)(nil)).Elem()
}

func (o IsGlobalPtrOutput) ToIsGlobalPtrOutput() IsGlobalPtrOutput {
	return o
}

func (o IsGlobalPtrOutput) ToIsGlobalPtrOutputWithContext(ctx context.Context) IsGlobalPtrOutput {
	return o
}

func (o IsGlobalPtrOutput) Elem() IsGlobalOutput {
	return o.ApplyT(func(v *IsGlobal) IsGlobal {
		if v != nil {
			return *v
		}
		var ret IsGlobal
		return ret
	}).(IsGlobalOutput)
}

func (o IsGlobalPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IsGlobalPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IsGlobal) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IsGlobalInput interface {
	pulumi.Input

	ToIsGlobalOutput() IsGlobalOutput
	ToIsGlobalOutputWithContext(context.Context) IsGlobalOutput
}

var isGlobalPtrType = reflect.TypeOf((**IsGlobal)(nil)).Elem()

type IsGlobalPtrInput interface {
	pulumi.Input

	ToIsGlobalPtrOutput() IsGlobalPtrOutput
	ToIsGlobalPtrOutputWithContext(context.Context) IsGlobalPtrOutput
}

type isGlobalPtr string

func IsGlobalPtr(v string) IsGlobalPtrInput {
	return (*isGlobalPtr)(&v)
}

func (*isGlobalPtr) ElementType() reflect.Type {
	return isGlobalPtrType
}

func (in *isGlobalPtr) ToIsGlobalPtrOutput() IsGlobalPtrOutput {
	return pulumi.ToOutput(in).(IsGlobalPtrOutput)
}

func (in *isGlobalPtr) ToIsGlobalPtrOutputWithContext(ctx context.Context) IsGlobalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IsGlobalPtrOutput)
}

type SecurityConfigurationRuleAccess string

const (
	SecurityConfigurationRuleAccessAllow       = SecurityConfigurationRuleAccess("Allow")
	SecurityConfigurationRuleAccessDeny        = SecurityConfigurationRuleAccess("Deny")
	SecurityConfigurationRuleAccessAlwaysAllow = SecurityConfigurationRuleAccess("AlwaysAllow")
)

func (SecurityConfigurationRuleAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConfigurationRuleAccess)(nil)).Elem()
}

func (e SecurityConfigurationRuleAccess) ToSecurityConfigurationRuleAccessOutput() SecurityConfigurationRuleAccessOutput {
	return pulumi.ToOutput(e).(SecurityConfigurationRuleAccessOutput)
}

func (e SecurityConfigurationRuleAccess) ToSecurityConfigurationRuleAccessOutputWithContext(ctx context.Context) SecurityConfigurationRuleAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityConfigurationRuleAccessOutput)
}

func (e SecurityConfigurationRuleAccess) ToSecurityConfigurationRuleAccessPtrOutput() SecurityConfigurationRuleAccessPtrOutput {
	return e.ToSecurityConfigurationRuleAccessPtrOutputWithContext(context.Background())
}

func (e SecurityConfigurationRuleAccess) ToSecurityConfigurationRuleAccessPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleAccessPtrOutput {
	return SecurityConfigurationRuleAccess(e).ToSecurityConfigurationRuleAccessOutputWithContext(ctx).ToSecurityConfigurationRuleAccessPtrOutputWithContext(ctx)
}

func (e SecurityConfigurationRuleAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityConfigurationRuleAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityConfigurationRuleAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityConfigurationRuleAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityConfigurationRuleAccessOutput struct{ *pulumi.OutputState }

func (SecurityConfigurationRuleAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConfigurationRuleAccess)(nil)).Elem()
}

func (o SecurityConfigurationRuleAccessOutput) ToSecurityConfigurationRuleAccessOutput() SecurityConfigurationRuleAccessOutput {
	return o
}

func (o SecurityConfigurationRuleAccessOutput) ToSecurityConfigurationRuleAccessOutputWithContext(ctx context.Context) SecurityConfigurationRuleAccessOutput {
	return o
}

func (o SecurityConfigurationRuleAccessOutput) ToSecurityConfigurationRuleAccessPtrOutput() SecurityConfigurationRuleAccessPtrOutput {
	return o.ToSecurityConfigurationRuleAccessPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleAccessOutput) ToSecurityConfigurationRuleAccessPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityConfigurationRuleAccess) *SecurityConfigurationRuleAccess {
		return &v
	}).(SecurityConfigurationRuleAccessPtrOutput)
}

func (o SecurityConfigurationRuleAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityConfigurationRuleAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityConfigurationRuleAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityConfigurationRuleAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityConfigurationRuleAccessPtrOutput struct{ *pulumi.OutputState }

func (SecurityConfigurationRuleAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConfigurationRuleAccess)(nil)).Elem()
}

func (o SecurityConfigurationRuleAccessPtrOutput) ToSecurityConfigurationRuleAccessPtrOutput() SecurityConfigurationRuleAccessPtrOutput {
	return o
}

func (o SecurityConfigurationRuleAccessPtrOutput) ToSecurityConfigurationRuleAccessPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleAccessPtrOutput {
	return o
}

func (o SecurityConfigurationRuleAccessPtrOutput) Elem() SecurityConfigurationRuleAccessOutput {
	return o.ApplyT(func(v *SecurityConfigurationRuleAccess) SecurityConfigurationRuleAccess {
		if v != nil {
			return *v
		}
		var ret SecurityConfigurationRuleAccess
		return ret
	}).(SecurityConfigurationRuleAccessOutput)
}

func (o SecurityConfigurationRuleAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityConfigurationRuleAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityConfigurationRuleAccessInput interface {
	pulumi.Input

	ToSecurityConfigurationRuleAccessOutput() SecurityConfigurationRuleAccessOutput
	ToSecurityConfigurationRuleAccessOutputWithContext(context.Context) SecurityConfigurationRuleAccessOutput
}

var securityConfigurationRuleAccessPtrType = reflect.TypeOf((**SecurityConfigurationRuleAccess)(nil)).Elem()

type SecurityConfigurationRuleAccessPtrInput interface {
	pulumi.Input

	ToSecurityConfigurationRuleAccessPtrOutput() SecurityConfigurationRuleAccessPtrOutput
	ToSecurityConfigurationRuleAccessPtrOutputWithContext(context.Context) SecurityConfigurationRuleAccessPtrOutput
}

type securityConfigurationRuleAccessPtr string

func SecurityConfigurationRuleAccessPtr(v string) SecurityConfigurationRuleAccessPtrInput {
	return (*securityConfigurationRuleAccessPtr)(&v)
}

func (*securityConfigurationRuleAccessPtr) ElementType() reflect.Type {
	return securityConfigurationRuleAccessPtrType
}

func (in *securityConfigurationRuleAccessPtr) ToSecurityConfigurationRuleAccessPtrOutput() SecurityConfigurationRuleAccessPtrOutput {
	return pulumi.ToOutput(in).(SecurityConfigurationRuleAccessPtrOutput)
}

func (in *securityConfigurationRuleAccessPtr) ToSecurityConfigurationRuleAccessPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityConfigurationRuleAccessPtrOutput)
}

type SecurityConfigurationRuleDirection string

const (
	SecurityConfigurationRuleDirectionInbound  = SecurityConfigurationRuleDirection("Inbound")
	SecurityConfigurationRuleDirectionOutbound = SecurityConfigurationRuleDirection("Outbound")
)

func (SecurityConfigurationRuleDirection) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConfigurationRuleDirection)(nil)).Elem()
}

func (e SecurityConfigurationRuleDirection) ToSecurityConfigurationRuleDirectionOutput() SecurityConfigurationRuleDirectionOutput {
	return pulumi.ToOutput(e).(SecurityConfigurationRuleDirectionOutput)
}

func (e SecurityConfigurationRuleDirection) ToSecurityConfigurationRuleDirectionOutputWithContext(ctx context.Context) SecurityConfigurationRuleDirectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityConfigurationRuleDirectionOutput)
}

func (e SecurityConfigurationRuleDirection) ToSecurityConfigurationRuleDirectionPtrOutput() SecurityConfigurationRuleDirectionPtrOutput {
	return e.ToSecurityConfigurationRuleDirectionPtrOutputWithContext(context.Background())
}

func (e SecurityConfigurationRuleDirection) ToSecurityConfigurationRuleDirectionPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleDirectionPtrOutput {
	return SecurityConfigurationRuleDirection(e).ToSecurityConfigurationRuleDirectionOutputWithContext(ctx).ToSecurityConfigurationRuleDirectionPtrOutputWithContext(ctx)
}

func (e SecurityConfigurationRuleDirection) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityConfigurationRuleDirection) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityConfigurationRuleDirection) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityConfigurationRuleDirection) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityConfigurationRuleDirectionOutput struct{ *pulumi.OutputState }

func (SecurityConfigurationRuleDirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConfigurationRuleDirection)(nil)).Elem()
}

func (o SecurityConfigurationRuleDirectionOutput) ToSecurityConfigurationRuleDirectionOutput() SecurityConfigurationRuleDirectionOutput {
	return o
}

func (o SecurityConfigurationRuleDirectionOutput) ToSecurityConfigurationRuleDirectionOutputWithContext(ctx context.Context) SecurityConfigurationRuleDirectionOutput {
	return o
}

func (o SecurityConfigurationRuleDirectionOutput) ToSecurityConfigurationRuleDirectionPtrOutput() SecurityConfigurationRuleDirectionPtrOutput {
	return o.ToSecurityConfigurationRuleDirectionPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleDirectionOutput) ToSecurityConfigurationRuleDirectionPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleDirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityConfigurationRuleDirection) *SecurityConfigurationRuleDirection {
		return &v
	}).(SecurityConfigurationRuleDirectionPtrOutput)
}

func (o SecurityConfigurationRuleDirectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleDirectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityConfigurationRuleDirection) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityConfigurationRuleDirectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleDirectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityConfigurationRuleDirection) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityConfigurationRuleDirectionPtrOutput struct{ *pulumi.OutputState }

func (SecurityConfigurationRuleDirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConfigurationRuleDirection)(nil)).Elem()
}

func (o SecurityConfigurationRuleDirectionPtrOutput) ToSecurityConfigurationRuleDirectionPtrOutput() SecurityConfigurationRuleDirectionPtrOutput {
	return o
}

func (o SecurityConfigurationRuleDirectionPtrOutput) ToSecurityConfigurationRuleDirectionPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleDirectionPtrOutput {
	return o
}

func (o SecurityConfigurationRuleDirectionPtrOutput) Elem() SecurityConfigurationRuleDirectionOutput {
	return o.ApplyT(func(v *SecurityConfigurationRuleDirection) SecurityConfigurationRuleDirection {
		if v != nil {
			return *v
		}
		var ret SecurityConfigurationRuleDirection
		return ret
	}).(SecurityConfigurationRuleDirectionOutput)
}

func (o SecurityConfigurationRuleDirectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleDirectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityConfigurationRuleDirection) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityConfigurationRuleDirectionInput interface {
	pulumi.Input

	ToSecurityConfigurationRuleDirectionOutput() SecurityConfigurationRuleDirectionOutput
	ToSecurityConfigurationRuleDirectionOutputWithContext(context.Context) SecurityConfigurationRuleDirectionOutput
}

var securityConfigurationRuleDirectionPtrType = reflect.TypeOf((**SecurityConfigurationRuleDirection)(nil)).Elem()

type SecurityConfigurationRuleDirectionPtrInput interface {
	pulumi.Input

	ToSecurityConfigurationRuleDirectionPtrOutput() SecurityConfigurationRuleDirectionPtrOutput
	ToSecurityConfigurationRuleDirectionPtrOutputWithContext(context.Context) SecurityConfigurationRuleDirectionPtrOutput
}

type securityConfigurationRuleDirectionPtr string

func SecurityConfigurationRuleDirectionPtr(v string) SecurityConfigurationRuleDirectionPtrInput {
	return (*securityConfigurationRuleDirectionPtr)(&v)
}

func (*securityConfigurationRuleDirectionPtr) ElementType() reflect.Type {
	return securityConfigurationRuleDirectionPtrType
}

func (in *securityConfigurationRuleDirectionPtr) ToSecurityConfigurationRuleDirectionPtrOutput() SecurityConfigurationRuleDirectionPtrOutput {
	return pulumi.ToOutput(in).(SecurityConfigurationRuleDirectionPtrOutput)
}

func (in *securityConfigurationRuleDirectionPtr) ToSecurityConfigurationRuleDirectionPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleDirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityConfigurationRuleDirectionPtrOutput)
}

type SecurityConfigurationRuleProtocol string

const (
	SecurityConfigurationRuleProtocolTcp  = SecurityConfigurationRuleProtocol("Tcp")
	SecurityConfigurationRuleProtocolUdp  = SecurityConfigurationRuleProtocol("Udp")
	SecurityConfigurationRuleProtocolIcmp = SecurityConfigurationRuleProtocol("Icmp")
	SecurityConfigurationRuleProtocolEsp  = SecurityConfigurationRuleProtocol("Esp")
	SecurityConfigurationRuleProtocolAny  = SecurityConfigurationRuleProtocol("Any")
	SecurityConfigurationRuleProtocolAh   = SecurityConfigurationRuleProtocol("Ah")
)

func (SecurityConfigurationRuleProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConfigurationRuleProtocol)(nil)).Elem()
}

func (e SecurityConfigurationRuleProtocol) ToSecurityConfigurationRuleProtocolOutput() SecurityConfigurationRuleProtocolOutput {
	return pulumi.ToOutput(e).(SecurityConfigurationRuleProtocolOutput)
}

func (e SecurityConfigurationRuleProtocol) ToSecurityConfigurationRuleProtocolOutputWithContext(ctx context.Context) SecurityConfigurationRuleProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityConfigurationRuleProtocolOutput)
}

func (e SecurityConfigurationRuleProtocol) ToSecurityConfigurationRuleProtocolPtrOutput() SecurityConfigurationRuleProtocolPtrOutput {
	return e.ToSecurityConfigurationRuleProtocolPtrOutputWithContext(context.Background())
}

func (e SecurityConfigurationRuleProtocol) ToSecurityConfigurationRuleProtocolPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleProtocolPtrOutput {
	return SecurityConfigurationRuleProtocol(e).ToSecurityConfigurationRuleProtocolOutputWithContext(ctx).ToSecurityConfigurationRuleProtocolPtrOutputWithContext(ctx)
}

func (e SecurityConfigurationRuleProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityConfigurationRuleProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityConfigurationRuleProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityConfigurationRuleProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityConfigurationRuleProtocolOutput struct{ *pulumi.OutputState }

func (SecurityConfigurationRuleProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConfigurationRuleProtocol)(nil)).Elem()
}

func (o SecurityConfigurationRuleProtocolOutput) ToSecurityConfigurationRuleProtocolOutput() SecurityConfigurationRuleProtocolOutput {
	return o
}

func (o SecurityConfigurationRuleProtocolOutput) ToSecurityConfigurationRuleProtocolOutputWithContext(ctx context.Context) SecurityConfigurationRuleProtocolOutput {
	return o
}

func (o SecurityConfigurationRuleProtocolOutput) ToSecurityConfigurationRuleProtocolPtrOutput() SecurityConfigurationRuleProtocolPtrOutput {
	return o.ToSecurityConfigurationRuleProtocolPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleProtocolOutput) ToSecurityConfigurationRuleProtocolPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityConfigurationRuleProtocol) *SecurityConfigurationRuleProtocol {
		return &v
	}).(SecurityConfigurationRuleProtocolPtrOutput)
}

func (o SecurityConfigurationRuleProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityConfigurationRuleProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityConfigurationRuleProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityConfigurationRuleProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityConfigurationRuleProtocolPtrOutput struct{ *pulumi.OutputState }

func (SecurityConfigurationRuleProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConfigurationRuleProtocol)(nil)).Elem()
}

func (o SecurityConfigurationRuleProtocolPtrOutput) ToSecurityConfigurationRuleProtocolPtrOutput() SecurityConfigurationRuleProtocolPtrOutput {
	return o
}

func (o SecurityConfigurationRuleProtocolPtrOutput) ToSecurityConfigurationRuleProtocolPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleProtocolPtrOutput {
	return o
}

func (o SecurityConfigurationRuleProtocolPtrOutput) Elem() SecurityConfigurationRuleProtocolOutput {
	return o.ApplyT(func(v *SecurityConfigurationRuleProtocol) SecurityConfigurationRuleProtocol {
		if v != nil {
			return *v
		}
		var ret SecurityConfigurationRuleProtocol
		return ret
	}).(SecurityConfigurationRuleProtocolOutput)
}

func (o SecurityConfigurationRuleProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationRuleProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityConfigurationRuleProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityConfigurationRuleProtocolInput interface {
	pulumi.Input

	ToSecurityConfigurationRuleProtocolOutput() SecurityConfigurationRuleProtocolOutput
	ToSecurityConfigurationRuleProtocolOutputWithContext(context.Context) SecurityConfigurationRuleProtocolOutput
}

var securityConfigurationRuleProtocolPtrType = reflect.TypeOf((**SecurityConfigurationRuleProtocol)(nil)).Elem()

type SecurityConfigurationRuleProtocolPtrInput interface {
	pulumi.Input

	ToSecurityConfigurationRuleProtocolPtrOutput() SecurityConfigurationRuleProtocolPtrOutput
	ToSecurityConfigurationRuleProtocolPtrOutputWithContext(context.Context) SecurityConfigurationRuleProtocolPtrOutput
}

type securityConfigurationRuleProtocolPtr string

func SecurityConfigurationRuleProtocolPtr(v string) SecurityConfigurationRuleProtocolPtrInput {
	return (*securityConfigurationRuleProtocolPtr)(&v)
}

func (*securityConfigurationRuleProtocolPtr) ElementType() reflect.Type {
	return securityConfigurationRuleProtocolPtrType
}

func (in *securityConfigurationRuleProtocolPtr) ToSecurityConfigurationRuleProtocolPtrOutput() SecurityConfigurationRuleProtocolPtrOutput {
	return pulumi.ToOutput(in).(SecurityConfigurationRuleProtocolPtrOutput)
}

func (in *securityConfigurationRuleProtocolPtr) ToSecurityConfigurationRuleProtocolPtrOutputWithContext(ctx context.Context) SecurityConfigurationRuleProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityConfigurationRuleProtocolPtrOutput)
}

type SecurityType string

const (
	SecurityTypeAdminPolicy = SecurityType("AdminPolicy")
	SecurityTypeUserPolicy  = SecurityType("UserPolicy")
)

func (SecurityType) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityType)(nil)).Elem()
}

func (e SecurityType) ToSecurityTypeOutput() SecurityTypeOutput {
	return pulumi.ToOutput(e).(SecurityTypeOutput)
}

func (e SecurityType) ToSecurityTypeOutputWithContext(ctx context.Context) SecurityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityTypeOutput)
}

func (e SecurityType) ToSecurityTypePtrOutput() SecurityTypePtrOutput {
	return e.ToSecurityTypePtrOutputWithContext(context.Background())
}

func (e SecurityType) ToSecurityTypePtrOutputWithContext(ctx context.Context) SecurityTypePtrOutput {
	return SecurityType(e).ToSecurityTypeOutputWithContext(ctx).ToSecurityTypePtrOutputWithContext(ctx)
}

func (e SecurityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityTypeOutput struct{ *pulumi.OutputState }

func (SecurityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityType)(nil)).Elem()
}

func (o SecurityTypeOutput) ToSecurityTypeOutput() SecurityTypeOutput {
	return o
}

func (o SecurityTypeOutput) ToSecurityTypeOutputWithContext(ctx context.Context) SecurityTypeOutput {
	return o
}

func (o SecurityTypeOutput) ToSecurityTypePtrOutput() SecurityTypePtrOutput {
	return o.ToSecurityTypePtrOutputWithContext(context.Background())
}

func (o SecurityTypeOutput) ToSecurityTypePtrOutputWithContext(ctx context.Context) SecurityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityType) *SecurityType {
		return &v
	}).(SecurityTypePtrOutput)
}

func (o SecurityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityTypePtrOutput struct{ *pulumi.OutputState }

func (SecurityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityType)(nil)).Elem()
}

func (o SecurityTypePtrOutput) ToSecurityTypePtrOutput() SecurityTypePtrOutput {
	return o
}

func (o SecurityTypePtrOutput) ToSecurityTypePtrOutputWithContext(ctx context.Context) SecurityTypePtrOutput {
	return o
}

func (o SecurityTypePtrOutput) Elem() SecurityTypeOutput {
	return o.ApplyT(func(v *SecurityType) SecurityType {
		if v != nil {
			return *v
		}
		var ret SecurityType
		return ret
	}).(SecurityTypeOutput)
}

func (o SecurityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityTypeInput interface {
	pulumi.Input

	ToSecurityTypeOutput() SecurityTypeOutput
	ToSecurityTypeOutputWithContext(context.Context) SecurityTypeOutput
}

var securityTypePtrType = reflect.TypeOf((**SecurityType)(nil)).Elem()

type SecurityTypePtrInput interface {
	pulumi.Input

	ToSecurityTypePtrOutput() SecurityTypePtrOutput
	ToSecurityTypePtrOutputWithContext(context.Context) SecurityTypePtrOutput
}

type securityTypePtr string

func SecurityTypePtr(v string) SecurityTypePtrInput {
	return (*securityTypePtr)(&v)
}

func (*securityTypePtr) ElementType() reflect.Type {
	return securityTypePtrType
}

func (in *securityTypePtr) ToSecurityTypePtrOutput() SecurityTypePtrOutput {
	return pulumi.ToOutput(in).(SecurityTypePtrOutput)
}

func (in *securityTypePtr) ToSecurityTypePtrOutputWithContext(ctx context.Context) SecurityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityTypePtrOutput)
}

type UseHubGateway string

const (
	UseHubGatewayFalse = UseHubGateway("False")
	UseHubGatewayTrue  = UseHubGateway("True")
)

func (UseHubGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*UseHubGateway)(nil)).Elem()
}

func (e UseHubGateway) ToUseHubGatewayOutput() UseHubGatewayOutput {
	return pulumi.ToOutput(e).(UseHubGatewayOutput)
}

func (e UseHubGateway) ToUseHubGatewayOutputWithContext(ctx context.Context) UseHubGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UseHubGatewayOutput)
}

func (e UseHubGateway) ToUseHubGatewayPtrOutput() UseHubGatewayPtrOutput {
	return e.ToUseHubGatewayPtrOutputWithContext(context.Background())
}

func (e UseHubGateway) ToUseHubGatewayPtrOutputWithContext(ctx context.Context) UseHubGatewayPtrOutput {
	return UseHubGateway(e).ToUseHubGatewayOutputWithContext(ctx).ToUseHubGatewayPtrOutputWithContext(ctx)
}

func (e UseHubGateway) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UseHubGateway) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UseHubGateway) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UseHubGateway) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UseHubGatewayOutput struct{ *pulumi.OutputState }

func (UseHubGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UseHubGateway)(nil)).Elem()
}

func (o UseHubGatewayOutput) ToUseHubGatewayOutput() UseHubGatewayOutput {
	return o
}

func (o UseHubGatewayOutput) ToUseHubGatewayOutputWithContext(ctx context.Context) UseHubGatewayOutput {
	return o
}

func (o UseHubGatewayOutput) ToUseHubGatewayPtrOutput() UseHubGatewayPtrOutput {
	return o.ToUseHubGatewayPtrOutputWithContext(context.Background())
}

func (o UseHubGatewayOutput) ToUseHubGatewayPtrOutputWithContext(ctx context.Context) UseHubGatewayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UseHubGateway) *UseHubGateway {
		return &v
	}).(UseHubGatewayPtrOutput)
}

func (o UseHubGatewayOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UseHubGatewayOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UseHubGateway) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UseHubGatewayOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UseHubGatewayOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UseHubGateway) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UseHubGatewayPtrOutput struct{ *pulumi.OutputState }

func (UseHubGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UseHubGateway)(nil)).Elem()
}

func (o UseHubGatewayPtrOutput) ToUseHubGatewayPtrOutput() UseHubGatewayPtrOutput {
	return o
}

func (o UseHubGatewayPtrOutput) ToUseHubGatewayPtrOutputWithContext(ctx context.Context) UseHubGatewayPtrOutput {
	return o
}

func (o UseHubGatewayPtrOutput) Elem() UseHubGatewayOutput {
	return o.ApplyT(func(v *UseHubGateway) UseHubGateway {
		if v != nil {
			return *v
		}
		var ret UseHubGateway
		return ret
	}).(UseHubGatewayOutput)
}

func (o UseHubGatewayPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UseHubGatewayPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UseHubGateway) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UseHubGatewayInput interface {
	pulumi.Input

	ToUseHubGatewayOutput() UseHubGatewayOutput
	ToUseHubGatewayOutputWithContext(context.Context) UseHubGatewayOutput
}

var useHubGatewayPtrType = reflect.TypeOf((**UseHubGateway)(nil)).Elem()

type UseHubGatewayPtrInput interface {
	pulumi.Input

	ToUseHubGatewayPtrOutput() UseHubGatewayPtrOutput
	ToUseHubGatewayPtrOutputWithContext(context.Context) UseHubGatewayPtrOutput
}

type useHubGatewayPtr string

func UseHubGatewayPtr(v string) UseHubGatewayPtrInput {
	return (*useHubGatewayPtr)(&v)
}

func (*useHubGatewayPtr) ElementType() reflect.Type {
	return useHubGatewayPtrType
}

func (in *useHubGatewayPtr) ToUseHubGatewayPtrOutput() UseHubGatewayPtrOutput {
	return pulumi.ToOutput(in).(UseHubGatewayPtrOutput)
}

func (in *useHubGatewayPtr) ToUseHubGatewayPtrOutputWithContext(ctx context.Context) UseHubGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UseHubGatewayPtrOutput)
}

type UserRuleKind string

const (
	UserRuleKindCustom  = UserRuleKind("Custom")
	UserRuleKindDefault = UserRuleKind("Default")
)

func (UserRuleKind) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRuleKind)(nil)).Elem()
}

func (e UserRuleKind) ToUserRuleKindOutput() UserRuleKindOutput {
	return pulumi.ToOutput(e).(UserRuleKindOutput)
}

func (e UserRuleKind) ToUserRuleKindOutputWithContext(ctx context.Context) UserRuleKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UserRuleKindOutput)
}

func (e UserRuleKind) ToUserRuleKindPtrOutput() UserRuleKindPtrOutput {
	return e.ToUserRuleKindPtrOutputWithContext(context.Background())
}

func (e UserRuleKind) ToUserRuleKindPtrOutputWithContext(ctx context.Context) UserRuleKindPtrOutput {
	return UserRuleKind(e).ToUserRuleKindOutputWithContext(ctx).ToUserRuleKindPtrOutputWithContext(ctx)
}

func (e UserRuleKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserRuleKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserRuleKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UserRuleKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UserRuleKindOutput struct{ *pulumi.OutputState }

func (UserRuleKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRuleKind)(nil)).Elem()
}

func (o UserRuleKindOutput) ToUserRuleKindOutput() UserRuleKindOutput {
	return o
}

func (o UserRuleKindOutput) ToUserRuleKindOutputWithContext(ctx context.Context) UserRuleKindOutput {
	return o
}

func (o UserRuleKindOutput) ToUserRuleKindPtrOutput() UserRuleKindPtrOutput {
	return o.ToUserRuleKindPtrOutputWithContext(context.Background())
}

func (o UserRuleKindOutput) ToUserRuleKindPtrOutputWithContext(ctx context.Context) UserRuleKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserRuleKind) *UserRuleKind {
		return &v
	}).(UserRuleKindPtrOutput)
}

func (o UserRuleKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UserRuleKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserRuleKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UserRuleKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserRuleKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserRuleKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UserRuleKindPtrOutput struct{ *pulumi.OutputState }

func (UserRuleKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRuleKind)(nil)).Elem()
}

func (o UserRuleKindPtrOutput) ToUserRuleKindPtrOutput() UserRuleKindPtrOutput {
	return o
}

func (o UserRuleKindPtrOutput) ToUserRuleKindPtrOutputWithContext(ctx context.Context) UserRuleKindPtrOutput {
	return o
}

func (o UserRuleKindPtrOutput) Elem() UserRuleKindOutput {
	return o.ApplyT(func(v *UserRuleKind) UserRuleKind {
		if v != nil {
			return *v
		}
		var ret UserRuleKind
		return ret
	}).(UserRuleKindOutput)
}

func (o UserRuleKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserRuleKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UserRuleKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UserRuleKindInput interface {
	pulumi.Input

	ToUserRuleKindOutput() UserRuleKindOutput
	ToUserRuleKindOutputWithContext(context.Context) UserRuleKindOutput
}

var userRuleKindPtrType = reflect.TypeOf((**UserRuleKind)(nil)).Elem()

type UserRuleKindPtrInput interface {
	pulumi.Input

	ToUserRuleKindPtrOutput() UserRuleKindPtrOutput
	ToUserRuleKindPtrOutputWithContext(context.Context) UserRuleKindPtrOutput
}

type userRuleKindPtr string

func UserRuleKindPtr(v string) UserRuleKindPtrInput {
	return (*userRuleKindPtr)(&v)
}

func (*userRuleKindPtr) ElementType() reflect.Type {
	return userRuleKindPtrType
}

func (in *userRuleKindPtr) ToUserRuleKindPtrOutput() UserRuleKindPtrOutput {
	return pulumi.ToOutput(in).(UserRuleKindPtrOutput)
}

func (in *userRuleKindPtr) ToUserRuleKindPtrOutputWithContext(ctx context.Context) UserRuleKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UserRuleKindPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPrefixTypeInput)(nil)).Elem(), AddressPrefixType("IPPrefix"))
	pulumi.RegisterInputType(reflect.TypeOf((*AddressPrefixTypePtrInput)(nil)).Elem(), AddressPrefixType("IPPrefix"))
	pulumi.RegisterInputType(reflect.TypeOf((*AdminRuleKindInput)(nil)).Elem(), AdminRuleKind("Custom"))
	pulumi.RegisterInputType(reflect.TypeOf((*AdminRuleKindPtrInput)(nil)).Elem(), AdminRuleKind("Custom"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationTypeInput)(nil)).Elem(), ConfigurationType("SecurityAdmin"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationTypePtrInput)(nil)).Elem(), ConfigurationType("SecurityAdmin"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectivityTopologyInput)(nil)).Elem(), ConnectivityTopology("HubAndSpoke"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectivityTopologyPtrInput)(nil)).Elem(), ConnectivityTopology("HubAndSpoke"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeleteExistingNSGsInput)(nil)).Elem(), DeleteExistingNSGs("False"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeleteExistingNSGsPtrInput)(nil)).Elem(), DeleteExistingNSGs("False"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeleteExistingPeeringInput)(nil)).Elem(), DeleteExistingPeering("False"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeleteExistingPeeringPtrInput)(nil)).Elem(), DeleteExistingPeering("False"))
	pulumi.RegisterInputType(reflect.TypeOf((*GroupConnectivityInput)(nil)).Elem(), GroupConnectivity("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*GroupConnectivityPtrInput)(nil)).Elem(), GroupConnectivity("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*IsGlobalInput)(nil)).Elem(), IsGlobal("False"))
	pulumi.RegisterInputType(reflect.TypeOf((*IsGlobalPtrInput)(nil)).Elem(), IsGlobal("False"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConfigurationRuleAccessInput)(nil)).Elem(), SecurityConfigurationRuleAccess("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConfigurationRuleAccessPtrInput)(nil)).Elem(), SecurityConfigurationRuleAccess("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConfigurationRuleDirectionInput)(nil)).Elem(), SecurityConfigurationRuleDirection("Inbound"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConfigurationRuleDirectionPtrInput)(nil)).Elem(), SecurityConfigurationRuleDirection("Inbound"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConfigurationRuleProtocolInput)(nil)).Elem(), SecurityConfigurationRuleProtocol("Tcp"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityConfigurationRuleProtocolPtrInput)(nil)).Elem(), SecurityConfigurationRuleProtocol("Tcp"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityTypeInput)(nil)).Elem(), SecurityType("AdminPolicy"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityTypePtrInput)(nil)).Elem(), SecurityType("AdminPolicy"))
	pulumi.RegisterInputType(reflect.TypeOf((*UseHubGatewayInput)(nil)).Elem(), UseHubGateway("False"))
	pulumi.RegisterInputType(reflect.TypeOf((*UseHubGatewayPtrInput)(nil)).Elem(), UseHubGateway("False"))
	pulumi.RegisterInputType(reflect.TypeOf((*UserRuleKindInput)(nil)).Elem(), UserRuleKind("Custom"))
	pulumi.RegisterInputType(reflect.TypeOf((*UserRuleKindPtrInput)(nil)).Elem(), UserRuleKind("Custom"))
	pulumi.RegisterOutputType(AddressPrefixTypeOutput{})
	pulumi.RegisterOutputType(AddressPrefixTypePtrOutput{})
	pulumi.RegisterOutputType(AdminRuleKindOutput{})
	pulumi.RegisterOutputType(AdminRuleKindPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationTypeOutput{})
	pulumi.RegisterOutputType(ConfigurationTypePtrOutput{})
	pulumi.RegisterOutputType(ConnectivityTopologyOutput{})
	pulumi.RegisterOutputType(ConnectivityTopologyPtrOutput{})
	pulumi.RegisterOutputType(DeleteExistingNSGsOutput{})
	pulumi.RegisterOutputType(DeleteExistingNSGsPtrOutput{})
	pulumi.RegisterOutputType(DeleteExistingPeeringOutput{})
	pulumi.RegisterOutputType(DeleteExistingPeeringPtrOutput{})
	pulumi.RegisterOutputType(GroupConnectivityOutput{})
	pulumi.RegisterOutputType(GroupConnectivityPtrOutput{})
	pulumi.RegisterOutputType(IsGlobalOutput{})
	pulumi.RegisterOutputType(IsGlobalPtrOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationRuleAccessOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationRuleAccessPtrOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationRuleDirectionOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationRuleDirectionPtrOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationRuleProtocolOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationRuleProtocolPtrOutput{})
	pulumi.RegisterOutputType(SecurityTypeOutput{})
	pulumi.RegisterOutputType(SecurityTypePtrOutput{})
	pulumi.RegisterOutputType(UseHubGatewayOutput{})
	pulumi.RegisterOutputType(UseHubGatewayPtrOutput{})
	pulumi.RegisterOutputType(UserRuleKindOutput{})
	pulumi.RegisterOutputType(UserRuleKindPtrOutput{})
}
