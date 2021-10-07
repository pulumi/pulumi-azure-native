


package deviceupdate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiagnosticStorageProperties struct {
	AuthenticationType string  `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	ResourceId         string  `pulumi:"resourceId"`
}





type DiagnosticStoragePropertiesInput interface {
	pulumi.Input

	ToDiagnosticStoragePropertiesOutput() DiagnosticStoragePropertiesOutput
	ToDiagnosticStoragePropertiesOutputWithContext(context.Context) DiagnosticStoragePropertiesOutput
}

type DiagnosticStoragePropertiesArgs struct {
	AuthenticationType pulumi.StringInput    `pulumi:"authenticationType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	ResourceId         pulumi.StringInput    `pulumi:"resourceId"`
}

func (DiagnosticStoragePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticStorageProperties)(nil)).Elem()
}

func (i DiagnosticStoragePropertiesArgs) ToDiagnosticStoragePropertiesOutput() DiagnosticStoragePropertiesOutput {
	return i.ToDiagnosticStoragePropertiesOutputWithContext(context.Background())
}

func (i DiagnosticStoragePropertiesArgs) ToDiagnosticStoragePropertiesOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesOutput)
}

func (i DiagnosticStoragePropertiesArgs) ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput {
	return i.ToDiagnosticStoragePropertiesPtrOutputWithContext(context.Background())
}

func (i DiagnosticStoragePropertiesArgs) ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesOutput).ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx)
}









type DiagnosticStoragePropertiesPtrInput interface {
	pulumi.Input

	ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput
	ToDiagnosticStoragePropertiesPtrOutputWithContext(context.Context) DiagnosticStoragePropertiesPtrOutput
}

type diagnosticStoragePropertiesPtrType DiagnosticStoragePropertiesArgs

func DiagnosticStoragePropertiesPtr(v *DiagnosticStoragePropertiesArgs) DiagnosticStoragePropertiesPtrInput {
	return (*diagnosticStoragePropertiesPtrType)(v)
}

func (*diagnosticStoragePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticStorageProperties)(nil)).Elem()
}

func (i *diagnosticStoragePropertiesPtrType) ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput {
	return i.ToDiagnosticStoragePropertiesPtrOutputWithContext(context.Background())
}

func (i *diagnosticStoragePropertiesPtrType) ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesPtrOutput)
}

type DiagnosticStoragePropertiesOutput struct{ *pulumi.OutputState }

func (DiagnosticStoragePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticStorageProperties)(nil)).Elem()
}

func (o DiagnosticStoragePropertiesOutput) ToDiagnosticStoragePropertiesOutput() DiagnosticStoragePropertiesOutput {
	return o
}

func (o DiagnosticStoragePropertiesOutput) ToDiagnosticStoragePropertiesOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesOutput {
	return o
}

func (o DiagnosticStoragePropertiesOutput) ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput {
	return o.ToDiagnosticStoragePropertiesPtrOutputWithContext(context.Background())
}

func (o DiagnosticStoragePropertiesOutput) ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticStorageProperties) *DiagnosticStorageProperties {
		return &v
	}).(DiagnosticStoragePropertiesPtrOutput)
}

func (o DiagnosticStoragePropertiesOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticStorageProperties) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o DiagnosticStoragePropertiesOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticStorageProperties) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticStorageProperties) string { return v.ResourceId }).(pulumi.StringOutput)
}

type DiagnosticStoragePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticStoragePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticStorageProperties)(nil)).Elem()
}

func (o DiagnosticStoragePropertiesPtrOutput) ToDiagnosticStoragePropertiesPtrOutput() DiagnosticStoragePropertiesPtrOutput {
	return o
}

func (o DiagnosticStoragePropertiesPtrOutput) ToDiagnosticStoragePropertiesPtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesPtrOutput {
	return o
}

func (o DiagnosticStoragePropertiesPtrOutput) Elem() DiagnosticStoragePropertiesOutput {
	return o.ApplyT(func(v *DiagnosticStorageProperties) DiagnosticStorageProperties {
		if v != nil {
			return *v
		}
		var ret DiagnosticStorageProperties
		return ret
	}).(DiagnosticStoragePropertiesOutput)
}

func (o DiagnosticStoragePropertiesPtrOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStorageProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AuthenticationType
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesPtrOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStorageProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStorageProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type DiagnosticStoragePropertiesResponse struct {
	AuthenticationType string  `pulumi:"authenticationType"`
	ConnectionString   *string `pulumi:"connectionString"`
	ResourceId         string  `pulumi:"resourceId"`
}





type DiagnosticStoragePropertiesResponseInput interface {
	pulumi.Input

	ToDiagnosticStoragePropertiesResponseOutput() DiagnosticStoragePropertiesResponseOutput
	ToDiagnosticStoragePropertiesResponseOutputWithContext(context.Context) DiagnosticStoragePropertiesResponseOutput
}

type DiagnosticStoragePropertiesResponseArgs struct {
	AuthenticationType pulumi.StringInput    `pulumi:"authenticationType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	ResourceId         pulumi.StringInput    `pulumi:"resourceId"`
}

func (DiagnosticStoragePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticStoragePropertiesResponse)(nil)).Elem()
}

func (i DiagnosticStoragePropertiesResponseArgs) ToDiagnosticStoragePropertiesResponseOutput() DiagnosticStoragePropertiesResponseOutput {
	return i.ToDiagnosticStoragePropertiesResponseOutputWithContext(context.Background())
}

func (i DiagnosticStoragePropertiesResponseArgs) ToDiagnosticStoragePropertiesResponseOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesResponseOutput)
}

func (i DiagnosticStoragePropertiesResponseArgs) ToDiagnosticStoragePropertiesResponsePtrOutput() DiagnosticStoragePropertiesResponsePtrOutput {
	return i.ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DiagnosticStoragePropertiesResponseArgs) ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesResponseOutput).ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(ctx)
}









type DiagnosticStoragePropertiesResponsePtrInput interface {
	pulumi.Input

	ToDiagnosticStoragePropertiesResponsePtrOutput() DiagnosticStoragePropertiesResponsePtrOutput
	ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(context.Context) DiagnosticStoragePropertiesResponsePtrOutput
}

type diagnosticStoragePropertiesResponsePtrType DiagnosticStoragePropertiesResponseArgs

func DiagnosticStoragePropertiesResponsePtr(v *DiagnosticStoragePropertiesResponseArgs) DiagnosticStoragePropertiesResponsePtrInput {
	return (*diagnosticStoragePropertiesResponsePtrType)(v)
}

func (*diagnosticStoragePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticStoragePropertiesResponse)(nil)).Elem()
}

func (i *diagnosticStoragePropertiesResponsePtrType) ToDiagnosticStoragePropertiesResponsePtrOutput() DiagnosticStoragePropertiesResponsePtrOutput {
	return i.ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *diagnosticStoragePropertiesResponsePtrType) ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticStoragePropertiesResponsePtrOutput)
}

type DiagnosticStoragePropertiesResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticStoragePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticStoragePropertiesResponse)(nil)).Elem()
}

func (o DiagnosticStoragePropertiesResponseOutput) ToDiagnosticStoragePropertiesResponseOutput() DiagnosticStoragePropertiesResponseOutput {
	return o
}

func (o DiagnosticStoragePropertiesResponseOutput) ToDiagnosticStoragePropertiesResponseOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesResponseOutput {
	return o
}

func (o DiagnosticStoragePropertiesResponseOutput) ToDiagnosticStoragePropertiesResponsePtrOutput() DiagnosticStoragePropertiesResponsePtrOutput {
	return o.ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DiagnosticStoragePropertiesResponseOutput) ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticStoragePropertiesResponse) *DiagnosticStoragePropertiesResponse {
		return &v
	}).(DiagnosticStoragePropertiesResponsePtrOutput)
}

func (o DiagnosticStoragePropertiesResponseOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticStoragePropertiesResponse) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o DiagnosticStoragePropertiesResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticStoragePropertiesResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticStoragePropertiesResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

type DiagnosticStoragePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticStoragePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticStoragePropertiesResponse)(nil)).Elem()
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) ToDiagnosticStoragePropertiesResponsePtrOutput() DiagnosticStoragePropertiesResponsePtrOutput {
	return o
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) ToDiagnosticStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) DiagnosticStoragePropertiesResponsePtrOutput {
	return o
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) Elem() DiagnosticStoragePropertiesResponseOutput {
	return o.ApplyT(func(v *DiagnosticStoragePropertiesResponse) DiagnosticStoragePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticStoragePropertiesResponse
		return ret
	}).(DiagnosticStoragePropertiesResponseOutput)
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AuthenticationType
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o DiagnosticStoragePropertiesResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type IotHubSettings struct {
	EventHubConnectionString *string `pulumi:"eventHubConnectionString"`
	IoTHubConnectionString   *string `pulumi:"ioTHubConnectionString"`
	ResourceId               string  `pulumi:"resourceId"`
}





type IotHubSettingsInput interface {
	pulumi.Input

	ToIotHubSettingsOutput() IotHubSettingsOutput
	ToIotHubSettingsOutputWithContext(context.Context) IotHubSettingsOutput
}

type IotHubSettingsArgs struct {
	EventHubConnectionString pulumi.StringPtrInput `pulumi:"eventHubConnectionString"`
	IoTHubConnectionString   pulumi.StringPtrInput `pulumi:"ioTHubConnectionString"`
	ResourceId               pulumi.StringInput    `pulumi:"resourceId"`
}

func (IotHubSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSettings)(nil)).Elem()
}

func (i IotHubSettingsArgs) ToIotHubSettingsOutput() IotHubSettingsOutput {
	return i.ToIotHubSettingsOutputWithContext(context.Background())
}

func (i IotHubSettingsArgs) ToIotHubSettingsOutputWithContext(ctx context.Context) IotHubSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSettingsOutput)
}





type IotHubSettingsArrayInput interface {
	pulumi.Input

	ToIotHubSettingsArrayOutput() IotHubSettingsArrayOutput
	ToIotHubSettingsArrayOutputWithContext(context.Context) IotHubSettingsArrayOutput
}

type IotHubSettingsArray []IotHubSettingsInput

func (IotHubSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubSettings)(nil)).Elem()
}

func (i IotHubSettingsArray) ToIotHubSettingsArrayOutput() IotHubSettingsArrayOutput {
	return i.ToIotHubSettingsArrayOutputWithContext(context.Background())
}

func (i IotHubSettingsArray) ToIotHubSettingsArrayOutputWithContext(ctx context.Context) IotHubSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSettingsArrayOutput)
}

type IotHubSettingsOutput struct{ *pulumi.OutputState }

func (IotHubSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSettings)(nil)).Elem()
}

func (o IotHubSettingsOutput) ToIotHubSettingsOutput() IotHubSettingsOutput {
	return o
}

func (o IotHubSettingsOutput) ToIotHubSettingsOutputWithContext(ctx context.Context) IotHubSettingsOutput {
	return o
}

func (o IotHubSettingsOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubSettings) *string { return v.EventHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o IotHubSettingsOutput) IoTHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubSettings) *string { return v.IoTHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o IotHubSettingsOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSettings) string { return v.ResourceId }).(pulumi.StringOutput)
}

type IotHubSettingsArrayOutput struct{ *pulumi.OutputState }

func (IotHubSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubSettings)(nil)).Elem()
}

func (o IotHubSettingsArrayOutput) ToIotHubSettingsArrayOutput() IotHubSettingsArrayOutput {
	return o
}

func (o IotHubSettingsArrayOutput) ToIotHubSettingsArrayOutputWithContext(ctx context.Context) IotHubSettingsArrayOutput {
	return o
}

func (o IotHubSettingsArrayOutput) Index(i pulumi.IntInput) IotHubSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubSettings {
		return vs[0].([]IotHubSettings)[vs[1].(int)]
	}).(IotHubSettingsOutput)
}

type IotHubSettingsResponse struct {
	EventHubConnectionString *string `pulumi:"eventHubConnectionString"`
	IoTHubConnectionString   *string `pulumi:"ioTHubConnectionString"`
	ResourceId               string  `pulumi:"resourceId"`
}





type IotHubSettingsResponseInput interface {
	pulumi.Input

	ToIotHubSettingsResponseOutput() IotHubSettingsResponseOutput
	ToIotHubSettingsResponseOutputWithContext(context.Context) IotHubSettingsResponseOutput
}

type IotHubSettingsResponseArgs struct {
	EventHubConnectionString pulumi.StringPtrInput `pulumi:"eventHubConnectionString"`
	IoTHubConnectionString   pulumi.StringPtrInput `pulumi:"ioTHubConnectionString"`
	ResourceId               pulumi.StringInput    `pulumi:"resourceId"`
}

func (IotHubSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSettingsResponse)(nil)).Elem()
}

func (i IotHubSettingsResponseArgs) ToIotHubSettingsResponseOutput() IotHubSettingsResponseOutput {
	return i.ToIotHubSettingsResponseOutputWithContext(context.Background())
}

func (i IotHubSettingsResponseArgs) ToIotHubSettingsResponseOutputWithContext(ctx context.Context) IotHubSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSettingsResponseOutput)
}





type IotHubSettingsResponseArrayInput interface {
	pulumi.Input

	ToIotHubSettingsResponseArrayOutput() IotHubSettingsResponseArrayOutput
	ToIotHubSettingsResponseArrayOutputWithContext(context.Context) IotHubSettingsResponseArrayOutput
}

type IotHubSettingsResponseArray []IotHubSettingsResponseInput

func (IotHubSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubSettingsResponse)(nil)).Elem()
}

func (i IotHubSettingsResponseArray) ToIotHubSettingsResponseArrayOutput() IotHubSettingsResponseArrayOutput {
	return i.ToIotHubSettingsResponseArrayOutputWithContext(context.Background())
}

func (i IotHubSettingsResponseArray) ToIotHubSettingsResponseArrayOutputWithContext(ctx context.Context) IotHubSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSettingsResponseArrayOutput)
}

type IotHubSettingsResponseOutput struct{ *pulumi.OutputState }

func (IotHubSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSettingsResponse)(nil)).Elem()
}

func (o IotHubSettingsResponseOutput) ToIotHubSettingsResponseOutput() IotHubSettingsResponseOutput {
	return o
}

func (o IotHubSettingsResponseOutput) ToIotHubSettingsResponseOutputWithContext(ctx context.Context) IotHubSettingsResponseOutput {
	return o
}

func (o IotHubSettingsResponseOutput) EventHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubSettingsResponse) *string { return v.EventHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o IotHubSettingsResponseOutput) IoTHubConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubSettingsResponse) *string { return v.IoTHubConnectionString }).(pulumi.StringPtrOutput)
}

func (o IotHubSettingsResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSettingsResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

type IotHubSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (IotHubSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubSettingsResponse)(nil)).Elem()
}

func (o IotHubSettingsResponseArrayOutput) ToIotHubSettingsResponseArrayOutput() IotHubSettingsResponseArrayOutput {
	return o
}

func (o IotHubSettingsResponseArrayOutput) ToIotHubSettingsResponseArrayOutputWithContext(ctx context.Context) IotHubSettingsResponseArrayOutput {
	return o
}

func (o IotHubSettingsResponseArrayOutput) Index(i pulumi.IntInput) IotHubSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubSettingsResponse {
		return vs[0].([]IotHubSettingsResponse)[vs[1].(int)]
	}).(IotHubSettingsResponseOutput)
}

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityResponseInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput
	ToManagedServiceIdentityResponseOutputWithContext(context.Context) ManagedServiceIdentityResponseOutput
}

type ManagedServiceIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                   `pulumi:"principalId"`
	TenantId               pulumi.StringInput                   `pulumi:"tenantId"`
	Type                   pulumi.StringInput                   `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return i.ToManagedServiceIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseOutput)
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return i.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseOutput).ToManagedServiceIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedServiceIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput
	ToManagedServiceIdentityResponsePtrOutputWithContext(context.Context) ManagedServiceIdentityResponsePtrOutput
}

type managedServiceIdentityResponsePtrType ManagedServiceIdentityResponseArgs

func ManagedServiceIdentityResponsePtr(v *ManagedServiceIdentityResponseArgs) ManagedServiceIdentityResponsePtrInput {
	return (*managedServiceIdentityResponsePtrType)(v)
}

func (*managedServiceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (i *managedServiceIdentityResponsePtrType) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return i.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityResponsePtrType) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponsePtrOutput)
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityResponse) *ManagedServiceIdentityResponse {
		return &v
	}).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return i.ToUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseOutput)
}





type UserAssignedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput
	ToUserAssignedIdentityResponseMapOutputWithContext(context.Context) UserAssignedIdentityResponseMapOutput
}

type UserAssignedIdentityResponseMap map[string]UserAssignedIdentityResponseInput

func (UserAssignedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return i.ToUserAssignedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseMapOutput)
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticStoragePropertiesOutput{})
	pulumi.RegisterOutputType(DiagnosticStoragePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticStoragePropertiesResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticStoragePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IotHubSettingsOutput{})
	pulumi.RegisterOutputType(IotHubSettingsArrayOutput{})
	pulumi.RegisterOutputType(IotHubSettingsResponseOutput{})
	pulumi.RegisterOutputType(IotHubSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
