


package v20200214privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type MaintenanceWindow struct {
	CustomWindow *string `pulumi:"customWindow"`
	DayOfWeek    *int    `pulumi:"dayOfWeek"`
	StartHour    *int    `pulumi:"startHour"`
	StartMinute  *int    `pulumi:"startMinute"`
}





type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(context.Context) MaintenanceWindowOutput
}

type MaintenanceWindowArgs struct {
	CustomWindow pulumi.StringPtrInput `pulumi:"customWindow"`
	DayOfWeek    pulumi.IntPtrInput    `pulumi:"dayOfWeek"`
	StartHour    pulumi.IntPtrInput    `pulumi:"startHour"`
	StartMinute  pulumi.IntPtrInput    `pulumi:"startMinute"`
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return i.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput).ToMaintenanceWindowPtrOutputWithContext(ctx)
}









type MaintenanceWindowPtrInput interface {
	pulumi.Input

	ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput
	ToMaintenanceWindowPtrOutputWithContext(context.Context) MaintenanceWindowPtrOutput
}

type maintenanceWindowPtrType MaintenanceWindowArgs

func MaintenanceWindowPtr(v *MaintenanceWindowArgs) MaintenanceWindowPtrInput {
	return (*maintenanceWindowPtrType)(v)
}

func (*maintenanceWindowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (i *maintenanceWindowPtrType) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return i.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (i *maintenanceWindowPtrType) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowPtrOutput)
}

type MaintenanceWindowOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return o.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MaintenanceWindow) *MaintenanceWindow {
		return &v
	}).(MaintenanceWindowPtrOutput)
}

func (o MaintenanceWindowOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *string { return v.CustomWindow }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.DayOfWeek }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.StartHour }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.StartMinute }).(pulumi.IntPtrOutput)
}

type MaintenanceWindowPtrOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowPtrOutput) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return o
}

func (o MaintenanceWindowPtrOutput) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return o
}

func (o MaintenanceWindowPtrOutput) Elem() MaintenanceWindowOutput {
	return o.ApplyT(func(v *MaintenanceWindow) MaintenanceWindow {
		if v != nil {
			return *v
		}
		var ret MaintenanceWindow
		return ret
	}).(MaintenanceWindowOutput)
}

func (o MaintenanceWindowPtrOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *string {
		if v == nil {
			return nil
		}
		return v.CustomWindow
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowPtrOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowPtrOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.StartHour
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowPtrOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.StartMinute
	}).(pulumi.IntPtrOutput)
}

type MaintenanceWindowResponse struct {
	CustomWindow *string `pulumi:"customWindow"`
	DayOfWeek    *int    `pulumi:"dayOfWeek"`
	StartHour    *int    `pulumi:"startHour"`
	StartMinute  *int    `pulumi:"startMinute"`
}

type MaintenanceWindowResponseOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowResponse)(nil)).Elem()
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponseOutput() MaintenanceWindowResponseOutput {
	return o
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponseOutputWithContext(ctx context.Context) MaintenanceWindowResponseOutput {
	return o
}

func (o MaintenanceWindowResponseOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *string { return v.CustomWindow }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowResponseOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.DayOfWeek }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponseOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.StartHour }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponseOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.StartMinute }).(pulumi.IntPtrOutput)
}

type MaintenanceWindowResponsePtrOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowResponse)(nil)).Elem()
}

func (o MaintenanceWindowResponsePtrOutput) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return o
}

func (o MaintenanceWindowResponsePtrOutput) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return o
}

func (o MaintenanceWindowResponsePtrOutput) Elem() MaintenanceWindowResponseOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) MaintenanceWindowResponse {
		if v != nil {
			return *v
		}
		var ret MaintenanceWindowResponse
		return ret
	}).(MaintenanceWindowResponseOutput)
}

func (o MaintenanceWindowResponsePtrOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomWindow
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.StartHour
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.StartMinute
	}).(pulumi.IntPtrOutput)
}

type ServerPropertiesDelegatedSubnetArguments struct {
	SubnetArmResourceId *string `pulumi:"subnetArmResourceId"`
}





type ServerPropertiesDelegatedSubnetArgumentsInput interface {
	pulumi.Input

	ToServerPropertiesDelegatedSubnetArgumentsOutput() ServerPropertiesDelegatedSubnetArgumentsOutput
	ToServerPropertiesDelegatedSubnetArgumentsOutputWithContext(context.Context) ServerPropertiesDelegatedSubnetArgumentsOutput
}

type ServerPropertiesDelegatedSubnetArgumentsArgs struct {
	SubnetArmResourceId pulumi.StringPtrInput `pulumi:"subnetArmResourceId"`
}

func (ServerPropertiesDelegatedSubnetArgumentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesDelegatedSubnetArguments)(nil)).Elem()
}

func (i ServerPropertiesDelegatedSubnetArgumentsArgs) ToServerPropertiesDelegatedSubnetArgumentsOutput() ServerPropertiesDelegatedSubnetArgumentsOutput {
	return i.ToServerPropertiesDelegatedSubnetArgumentsOutputWithContext(context.Background())
}

func (i ServerPropertiesDelegatedSubnetArgumentsArgs) ToServerPropertiesDelegatedSubnetArgumentsOutputWithContext(ctx context.Context) ServerPropertiesDelegatedSubnetArgumentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesDelegatedSubnetArgumentsOutput)
}

func (i ServerPropertiesDelegatedSubnetArgumentsArgs) ToServerPropertiesDelegatedSubnetArgumentsPtrOutput() ServerPropertiesDelegatedSubnetArgumentsPtrOutput {
	return i.ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (i ServerPropertiesDelegatedSubnetArgumentsArgs) ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesDelegatedSubnetArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesDelegatedSubnetArgumentsOutput).ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx)
}









type ServerPropertiesDelegatedSubnetArgumentsPtrInput interface {
	pulumi.Input

	ToServerPropertiesDelegatedSubnetArgumentsPtrOutput() ServerPropertiesDelegatedSubnetArgumentsPtrOutput
	ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(context.Context) ServerPropertiesDelegatedSubnetArgumentsPtrOutput
}

type serverPropertiesDelegatedSubnetArgumentsPtrType ServerPropertiesDelegatedSubnetArgumentsArgs

func ServerPropertiesDelegatedSubnetArgumentsPtr(v *ServerPropertiesDelegatedSubnetArgumentsArgs) ServerPropertiesDelegatedSubnetArgumentsPtrInput {
	return (*serverPropertiesDelegatedSubnetArgumentsPtrType)(v)
}

func (*serverPropertiesDelegatedSubnetArgumentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPropertiesDelegatedSubnetArguments)(nil)).Elem()
}

func (i *serverPropertiesDelegatedSubnetArgumentsPtrType) ToServerPropertiesDelegatedSubnetArgumentsPtrOutput() ServerPropertiesDelegatedSubnetArgumentsPtrOutput {
	return i.ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (i *serverPropertiesDelegatedSubnetArgumentsPtrType) ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesDelegatedSubnetArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesDelegatedSubnetArgumentsPtrOutput)
}

type ServerPropertiesDelegatedSubnetArgumentsOutput struct{ *pulumi.OutputState }

func (ServerPropertiesDelegatedSubnetArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesDelegatedSubnetArguments)(nil)).Elem()
}

func (o ServerPropertiesDelegatedSubnetArgumentsOutput) ToServerPropertiesDelegatedSubnetArgumentsOutput() ServerPropertiesDelegatedSubnetArgumentsOutput {
	return o
}

func (o ServerPropertiesDelegatedSubnetArgumentsOutput) ToServerPropertiesDelegatedSubnetArgumentsOutputWithContext(ctx context.Context) ServerPropertiesDelegatedSubnetArgumentsOutput {
	return o
}

func (o ServerPropertiesDelegatedSubnetArgumentsOutput) ToServerPropertiesDelegatedSubnetArgumentsPtrOutput() ServerPropertiesDelegatedSubnetArgumentsPtrOutput {
	return o.ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (o ServerPropertiesDelegatedSubnetArgumentsOutput) ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesDelegatedSubnetArgumentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerPropertiesDelegatedSubnetArguments) *ServerPropertiesDelegatedSubnetArguments {
		return &v
	}).(ServerPropertiesDelegatedSubnetArgumentsPtrOutput)
}

func (o ServerPropertiesDelegatedSubnetArgumentsOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesDelegatedSubnetArguments) *string { return v.SubnetArmResourceId }).(pulumi.StringPtrOutput)
}

type ServerPropertiesDelegatedSubnetArgumentsPtrOutput struct{ *pulumi.OutputState }

func (ServerPropertiesDelegatedSubnetArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPropertiesDelegatedSubnetArguments)(nil)).Elem()
}

func (o ServerPropertiesDelegatedSubnetArgumentsPtrOutput) ToServerPropertiesDelegatedSubnetArgumentsPtrOutput() ServerPropertiesDelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o ServerPropertiesDelegatedSubnetArgumentsPtrOutput) ToServerPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesDelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o ServerPropertiesDelegatedSubnetArgumentsPtrOutput) Elem() ServerPropertiesDelegatedSubnetArgumentsOutput {
	return o.ApplyT(func(v *ServerPropertiesDelegatedSubnetArguments) ServerPropertiesDelegatedSubnetArguments {
		if v != nil {
			return *v
		}
		var ret ServerPropertiesDelegatedSubnetArguments
		return ret
	}).(ServerPropertiesDelegatedSubnetArgumentsOutput)
}

func (o ServerPropertiesDelegatedSubnetArgumentsPtrOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerPropertiesDelegatedSubnetArguments) *string {
		if v == nil {
			return nil
		}
		return v.SubnetArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ServerPropertiesPrivateDnsZoneArguments struct {
	PrivateDnsZoneArmResourceId *string `pulumi:"privateDnsZoneArmResourceId"`
}





type ServerPropertiesPrivateDnsZoneArgumentsInput interface {
	pulumi.Input

	ToServerPropertiesPrivateDnsZoneArgumentsOutput() ServerPropertiesPrivateDnsZoneArgumentsOutput
	ToServerPropertiesPrivateDnsZoneArgumentsOutputWithContext(context.Context) ServerPropertiesPrivateDnsZoneArgumentsOutput
}

type ServerPropertiesPrivateDnsZoneArgumentsArgs struct {
	PrivateDnsZoneArmResourceId pulumi.StringPtrInput `pulumi:"privateDnsZoneArmResourceId"`
}

func (ServerPropertiesPrivateDnsZoneArgumentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesPrivateDnsZoneArguments)(nil)).Elem()
}

func (i ServerPropertiesPrivateDnsZoneArgumentsArgs) ToServerPropertiesPrivateDnsZoneArgumentsOutput() ServerPropertiesPrivateDnsZoneArgumentsOutput {
	return i.ToServerPropertiesPrivateDnsZoneArgumentsOutputWithContext(context.Background())
}

func (i ServerPropertiesPrivateDnsZoneArgumentsArgs) ToServerPropertiesPrivateDnsZoneArgumentsOutputWithContext(ctx context.Context) ServerPropertiesPrivateDnsZoneArgumentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesPrivateDnsZoneArgumentsOutput)
}

func (i ServerPropertiesPrivateDnsZoneArgumentsArgs) ToServerPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return i.ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(context.Background())
}

func (i ServerPropertiesPrivateDnsZoneArgumentsArgs) ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesPrivateDnsZoneArgumentsOutput).ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx)
}









type ServerPropertiesPrivateDnsZoneArgumentsPtrInput interface {
	pulumi.Input

	ToServerPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerPropertiesPrivateDnsZoneArgumentsPtrOutput
	ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(context.Context) ServerPropertiesPrivateDnsZoneArgumentsPtrOutput
}

type serverPropertiesPrivateDnsZoneArgumentsPtrType ServerPropertiesPrivateDnsZoneArgumentsArgs

func ServerPropertiesPrivateDnsZoneArgumentsPtr(v *ServerPropertiesPrivateDnsZoneArgumentsArgs) ServerPropertiesPrivateDnsZoneArgumentsPtrInput {
	return (*serverPropertiesPrivateDnsZoneArgumentsPtrType)(v)
}

func (*serverPropertiesPrivateDnsZoneArgumentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPropertiesPrivateDnsZoneArguments)(nil)).Elem()
}

func (i *serverPropertiesPrivateDnsZoneArgumentsPtrType) ToServerPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return i.ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(context.Background())
}

func (i *serverPropertiesPrivateDnsZoneArgumentsPtrType) ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesPrivateDnsZoneArgumentsPtrOutput)
}

type ServerPropertiesPrivateDnsZoneArgumentsOutput struct{ *pulumi.OutputState }

func (ServerPropertiesPrivateDnsZoneArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesPrivateDnsZoneArguments)(nil)).Elem()
}

func (o ServerPropertiesPrivateDnsZoneArgumentsOutput) ToServerPropertiesPrivateDnsZoneArgumentsOutput() ServerPropertiesPrivateDnsZoneArgumentsOutput {
	return o
}

func (o ServerPropertiesPrivateDnsZoneArgumentsOutput) ToServerPropertiesPrivateDnsZoneArgumentsOutputWithContext(ctx context.Context) ServerPropertiesPrivateDnsZoneArgumentsOutput {
	return o
}

func (o ServerPropertiesPrivateDnsZoneArgumentsOutput) ToServerPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return o.ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(context.Background())
}

func (o ServerPropertiesPrivateDnsZoneArgumentsOutput) ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerPropertiesPrivateDnsZoneArguments) *ServerPropertiesPrivateDnsZoneArguments {
		return &v
	}).(ServerPropertiesPrivateDnsZoneArgumentsPtrOutput)
}

func (o ServerPropertiesPrivateDnsZoneArgumentsOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesPrivateDnsZoneArguments) *string { return v.PrivateDnsZoneArmResourceId }).(pulumi.StringPtrOutput)
}

type ServerPropertiesPrivateDnsZoneArgumentsPtrOutput struct{ *pulumi.OutputState }

func (ServerPropertiesPrivateDnsZoneArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPropertiesPrivateDnsZoneArguments)(nil)).Elem()
}

func (o ServerPropertiesPrivateDnsZoneArgumentsPtrOutput) ToServerPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return o
}

func (o ServerPropertiesPrivateDnsZoneArgumentsPtrOutput) ToServerPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return o
}

func (o ServerPropertiesPrivateDnsZoneArgumentsPtrOutput) Elem() ServerPropertiesPrivateDnsZoneArgumentsOutput {
	return o.ApplyT(func(v *ServerPropertiesPrivateDnsZoneArguments) ServerPropertiesPrivateDnsZoneArguments {
		if v != nil {
			return *v
		}
		var ret ServerPropertiesPrivateDnsZoneArguments
		return ret
	}).(ServerPropertiesPrivateDnsZoneArgumentsOutput)
}

func (o ServerPropertiesPrivateDnsZoneArgumentsPtrOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerPropertiesPrivateDnsZoneArguments) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ServerPropertiesResponseDelegatedSubnetArguments struct {
	SubnetArmResourceId *string `pulumi:"subnetArmResourceId"`
}

type ServerPropertiesResponseDelegatedSubnetArgumentsOutput struct{ *pulumi.OutputState }

func (ServerPropertiesResponseDelegatedSubnetArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesResponseDelegatedSubnetArguments)(nil)).Elem()
}

func (o ServerPropertiesResponseDelegatedSubnetArgumentsOutput) ToServerPropertiesResponseDelegatedSubnetArgumentsOutput() ServerPropertiesResponseDelegatedSubnetArgumentsOutput {
	return o
}

func (o ServerPropertiesResponseDelegatedSubnetArgumentsOutput) ToServerPropertiesResponseDelegatedSubnetArgumentsOutputWithContext(ctx context.Context) ServerPropertiesResponseDelegatedSubnetArgumentsOutput {
	return o
}

func (o ServerPropertiesResponseDelegatedSubnetArgumentsOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesResponseDelegatedSubnetArguments) *string { return v.SubnetArmResourceId }).(pulumi.StringPtrOutput)
}

type ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput struct{ *pulumi.OutputState }

func (ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPropertiesResponseDelegatedSubnetArguments)(nil)).Elem()
}

func (o ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput) ToServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput() ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput) ToServerPropertiesResponseDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput) Elem() ServerPropertiesResponseDelegatedSubnetArgumentsOutput {
	return o.ApplyT(func(v *ServerPropertiesResponseDelegatedSubnetArguments) ServerPropertiesResponseDelegatedSubnetArguments {
		if v != nil {
			return *v
		}
		var ret ServerPropertiesResponseDelegatedSubnetArguments
		return ret
	}).(ServerPropertiesResponseDelegatedSubnetArgumentsOutput)
}

func (o ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerPropertiesResponseDelegatedSubnetArguments) *string {
		if v == nil {
			return nil
		}
		return v.SubnetArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ServerPropertiesResponsePrivateDnsZoneArguments struct {
	PrivateDnsZoneArmResourceId *string `pulumi:"privateDnsZoneArmResourceId"`
}

type ServerPropertiesResponsePrivateDnsZoneArgumentsOutput struct{ *pulumi.OutputState }

func (ServerPropertiesResponsePrivateDnsZoneArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesResponsePrivateDnsZoneArguments)(nil)).Elem()
}

func (o ServerPropertiesResponsePrivateDnsZoneArgumentsOutput) ToServerPropertiesResponsePrivateDnsZoneArgumentsOutput() ServerPropertiesResponsePrivateDnsZoneArgumentsOutput {
	return o
}

func (o ServerPropertiesResponsePrivateDnsZoneArgumentsOutput) ToServerPropertiesResponsePrivateDnsZoneArgumentsOutputWithContext(ctx context.Context) ServerPropertiesResponsePrivateDnsZoneArgumentsOutput {
	return o
}

func (o ServerPropertiesResponsePrivateDnsZoneArgumentsOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesResponsePrivateDnsZoneArguments) *string { return v.PrivateDnsZoneArmResourceId }).(pulumi.StringPtrOutput)
}

type ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput struct{ *pulumi.OutputState }

func (ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPropertiesResponsePrivateDnsZoneArguments)(nil)).Elem()
}

func (o ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) ToServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput() ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o
}

func (o ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) ToServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o
}

func (o ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) Elem() ServerPropertiesResponsePrivateDnsZoneArgumentsOutput {
	return o.ApplyT(func(v *ServerPropertiesResponsePrivateDnsZoneArguments) ServerPropertiesResponsePrivateDnsZoneArguments {
		if v != nil {
			return *v
		}
		var ret ServerPropertiesResponsePrivateDnsZoneArguments
		return ret
	}).(ServerPropertiesResponsePrivateDnsZoneArgumentsOutput)
}

func (o ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerPropertiesResponsePrivateDnsZoneArguments) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type StorageProfile struct {
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	StorageMB           *int `pulumi:"storageMB"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	BackupRetentionDays pulumi.IntPtrInput `pulumi:"backupRetentionDays"`
	StorageMB           pulumi.IntPtrInput `pulumi:"storageMB"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o StorageProfileOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.StorageMB }).(pulumi.IntPtrOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionDays
	}).(pulumi.IntPtrOutput)
}

func (o StorageProfilePtrOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *int {
		if v == nil {
			return nil
		}
		return v.StorageMB
	}).(pulumi.IntPtrOutput)
}

type StorageProfileResponse struct {
	BackupRetentionDays *int `pulumi:"backupRetentionDays"`
	StorageMB           *int `pulumi:"storageMB"`
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponseOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.StorageMB }).(pulumi.IntPtrOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionDays
	}).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponsePtrOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.StorageMB
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerPropertiesDelegatedSubnetArgumentsOutput{})
	pulumi.RegisterOutputType(ServerPropertiesDelegatedSubnetArgumentsPtrOutput{})
	pulumi.RegisterOutputType(ServerPropertiesPrivateDnsZoneArgumentsOutput{})
	pulumi.RegisterOutputType(ServerPropertiesPrivateDnsZoneArgumentsPtrOutput{})
	pulumi.RegisterOutputType(ServerPropertiesResponseDelegatedSubnetArgumentsOutput{})
	pulumi.RegisterOutputType(ServerPropertiesResponseDelegatedSubnetArgumentsPtrOutput{})
	pulumi.RegisterOutputType(ServerPropertiesResponsePrivateDnsZoneArgumentsOutput{})
	pulumi.RegisterOutputType(ServerPropertiesResponsePrivateDnsZoneArgumentsPtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
}
