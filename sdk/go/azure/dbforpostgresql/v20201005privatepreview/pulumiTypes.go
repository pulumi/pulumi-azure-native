


package v20201005privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

type ServerGroupPropertiesDelegatedSubnetArguments struct {
	SubnetArmResourceId *string `pulumi:"subnetArmResourceId"`
}





type ServerGroupPropertiesDelegatedSubnetArgumentsInput interface {
	pulumi.Input

	ToServerGroupPropertiesDelegatedSubnetArgumentsOutput() ServerGroupPropertiesDelegatedSubnetArgumentsOutput
	ToServerGroupPropertiesDelegatedSubnetArgumentsOutputWithContext(context.Context) ServerGroupPropertiesDelegatedSubnetArgumentsOutput
}

type ServerGroupPropertiesDelegatedSubnetArgumentsArgs struct {
	SubnetArmResourceId pulumi.StringPtrInput `pulumi:"subnetArmResourceId"`
}

func (ServerGroupPropertiesDelegatedSubnetArgumentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupPropertiesDelegatedSubnetArguments)(nil)).Elem()
}

func (i ServerGroupPropertiesDelegatedSubnetArgumentsArgs) ToServerGroupPropertiesDelegatedSubnetArgumentsOutput() ServerGroupPropertiesDelegatedSubnetArgumentsOutput {
	return i.ToServerGroupPropertiesDelegatedSubnetArgumentsOutputWithContext(context.Background())
}

func (i ServerGroupPropertiesDelegatedSubnetArgumentsArgs) ToServerGroupPropertiesDelegatedSubnetArgumentsOutputWithContext(ctx context.Context) ServerGroupPropertiesDelegatedSubnetArgumentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupPropertiesDelegatedSubnetArgumentsOutput)
}

func (i ServerGroupPropertiesDelegatedSubnetArgumentsArgs) ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput() ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput {
	return i.ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (i ServerGroupPropertiesDelegatedSubnetArgumentsArgs) ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupPropertiesDelegatedSubnetArgumentsOutput).ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx)
}









type ServerGroupPropertiesDelegatedSubnetArgumentsPtrInput interface {
	pulumi.Input

	ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput() ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput
	ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(context.Context) ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput
}

type serverGroupPropertiesDelegatedSubnetArgumentsPtrType ServerGroupPropertiesDelegatedSubnetArgumentsArgs

func ServerGroupPropertiesDelegatedSubnetArgumentsPtr(v *ServerGroupPropertiesDelegatedSubnetArgumentsArgs) ServerGroupPropertiesDelegatedSubnetArgumentsPtrInput {
	return (*serverGroupPropertiesDelegatedSubnetArgumentsPtrType)(v)
}

func (*serverGroupPropertiesDelegatedSubnetArgumentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupPropertiesDelegatedSubnetArguments)(nil)).Elem()
}

func (i *serverGroupPropertiesDelegatedSubnetArgumentsPtrType) ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput() ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput {
	return i.ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (i *serverGroupPropertiesDelegatedSubnetArgumentsPtrType) ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput)
}

type ServerGroupPropertiesDelegatedSubnetArgumentsOutput struct{ *pulumi.OutputState }

func (ServerGroupPropertiesDelegatedSubnetArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupPropertiesDelegatedSubnetArguments)(nil)).Elem()
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsOutput) ToServerGroupPropertiesDelegatedSubnetArgumentsOutput() ServerGroupPropertiesDelegatedSubnetArgumentsOutput {
	return o
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsOutput) ToServerGroupPropertiesDelegatedSubnetArgumentsOutputWithContext(ctx context.Context) ServerGroupPropertiesDelegatedSubnetArgumentsOutput {
	return o
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsOutput) ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput() ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput {
	return o.ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsOutput) ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerGroupPropertiesDelegatedSubnetArguments) *ServerGroupPropertiesDelegatedSubnetArguments {
		return &v
	}).(ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput)
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupPropertiesDelegatedSubnetArguments) *string { return v.SubnetArmResourceId }).(pulumi.StringPtrOutput)
}

type ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput struct{ *pulumi.OutputState }

func (ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupPropertiesDelegatedSubnetArguments)(nil)).Elem()
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput) ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput() ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput) ToServerGroupPropertiesDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput) Elem() ServerGroupPropertiesDelegatedSubnetArgumentsOutput {
	return o.ApplyT(func(v *ServerGroupPropertiesDelegatedSubnetArguments) ServerGroupPropertiesDelegatedSubnetArguments {
		if v != nil {
			return *v
		}
		var ret ServerGroupPropertiesDelegatedSubnetArguments
		return ret
	}).(ServerGroupPropertiesDelegatedSubnetArgumentsOutput)
}

func (o ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroupPropertiesDelegatedSubnetArguments) *string {
		if v == nil {
			return nil
		}
		return v.SubnetArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ServerGroupPropertiesPrivateDnsZoneArguments struct {
	PrivateDnsZoneArmResourceId *string `pulumi:"privateDnsZoneArmResourceId"`
}





type ServerGroupPropertiesPrivateDnsZoneArgumentsInput interface {
	pulumi.Input

	ToServerGroupPropertiesPrivateDnsZoneArgumentsOutput() ServerGroupPropertiesPrivateDnsZoneArgumentsOutput
	ToServerGroupPropertiesPrivateDnsZoneArgumentsOutputWithContext(context.Context) ServerGroupPropertiesPrivateDnsZoneArgumentsOutput
}

type ServerGroupPropertiesPrivateDnsZoneArgumentsArgs struct {
	PrivateDnsZoneArmResourceId pulumi.StringPtrInput `pulumi:"privateDnsZoneArmResourceId"`
}

func (ServerGroupPropertiesPrivateDnsZoneArgumentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupPropertiesPrivateDnsZoneArguments)(nil)).Elem()
}

func (i ServerGroupPropertiesPrivateDnsZoneArgumentsArgs) ToServerGroupPropertiesPrivateDnsZoneArgumentsOutput() ServerGroupPropertiesPrivateDnsZoneArgumentsOutput {
	return i.ToServerGroupPropertiesPrivateDnsZoneArgumentsOutputWithContext(context.Background())
}

func (i ServerGroupPropertiesPrivateDnsZoneArgumentsArgs) ToServerGroupPropertiesPrivateDnsZoneArgumentsOutputWithContext(ctx context.Context) ServerGroupPropertiesPrivateDnsZoneArgumentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupPropertiesPrivateDnsZoneArgumentsOutput)
}

func (i ServerGroupPropertiesPrivateDnsZoneArgumentsArgs) ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return i.ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(context.Background())
}

func (i ServerGroupPropertiesPrivateDnsZoneArgumentsArgs) ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupPropertiesPrivateDnsZoneArgumentsOutput).ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx)
}









type ServerGroupPropertiesPrivateDnsZoneArgumentsPtrInput interface {
	pulumi.Input

	ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput
	ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(context.Context) ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput
}

type serverGroupPropertiesPrivateDnsZoneArgumentsPtrType ServerGroupPropertiesPrivateDnsZoneArgumentsArgs

func ServerGroupPropertiesPrivateDnsZoneArgumentsPtr(v *ServerGroupPropertiesPrivateDnsZoneArgumentsArgs) ServerGroupPropertiesPrivateDnsZoneArgumentsPtrInput {
	return (*serverGroupPropertiesPrivateDnsZoneArgumentsPtrType)(v)
}

func (*serverGroupPropertiesPrivateDnsZoneArgumentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupPropertiesPrivateDnsZoneArguments)(nil)).Elem()
}

func (i *serverGroupPropertiesPrivateDnsZoneArgumentsPtrType) ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return i.ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(context.Background())
}

func (i *serverGroupPropertiesPrivateDnsZoneArgumentsPtrType) ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput)
}

type ServerGroupPropertiesPrivateDnsZoneArgumentsOutput struct{ *pulumi.OutputState }

func (ServerGroupPropertiesPrivateDnsZoneArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupPropertiesPrivateDnsZoneArguments)(nil)).Elem()
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsOutput) ToServerGroupPropertiesPrivateDnsZoneArgumentsOutput() ServerGroupPropertiesPrivateDnsZoneArgumentsOutput {
	return o
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsOutput) ToServerGroupPropertiesPrivateDnsZoneArgumentsOutputWithContext(ctx context.Context) ServerGroupPropertiesPrivateDnsZoneArgumentsOutput {
	return o
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsOutput) ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return o.ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(context.Background())
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsOutput) ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerGroupPropertiesPrivateDnsZoneArguments) *ServerGroupPropertiesPrivateDnsZoneArguments {
		return &v
	}).(ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput)
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupPropertiesPrivateDnsZoneArguments) *string { return v.PrivateDnsZoneArmResourceId }).(pulumi.StringPtrOutput)
}

type ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput struct{ *pulumi.OutputState }

func (ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupPropertiesPrivateDnsZoneArguments)(nil)).Elem()
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput) ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput() ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return o
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput) ToServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput {
	return o
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput) Elem() ServerGroupPropertiesPrivateDnsZoneArgumentsOutput {
	return o.ApplyT(func(v *ServerGroupPropertiesPrivateDnsZoneArguments) ServerGroupPropertiesPrivateDnsZoneArguments {
		if v != nil {
			return *v
		}
		var ret ServerGroupPropertiesPrivateDnsZoneArguments
		return ret
	}).(ServerGroupPropertiesPrivateDnsZoneArgumentsOutput)
}

func (o ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroupPropertiesPrivateDnsZoneArguments) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ServerGroupPropertiesResponseDelegatedSubnetArguments struct {
	SubnetArmResourceId *string `pulumi:"subnetArmResourceId"`
}

type ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput struct{ *pulumi.OutputState }

func (ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupPropertiesResponseDelegatedSubnetArguments)(nil)).Elem()
}

func (o ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput) ToServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput() ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput {
	return o
}

func (o ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput) ToServerGroupPropertiesResponseDelegatedSubnetArgumentsOutputWithContext(ctx context.Context) ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput {
	return o
}

func (o ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupPropertiesResponseDelegatedSubnetArguments) *string { return v.SubnetArmResourceId }).(pulumi.StringPtrOutput)
}

type ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput struct{ *pulumi.OutputState }

func (ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupPropertiesResponseDelegatedSubnetArguments)(nil)).Elem()
}

func (o ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput) ToServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput() ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput) ToServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput) Elem() ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput {
	return o.ApplyT(func(v *ServerGroupPropertiesResponseDelegatedSubnetArguments) ServerGroupPropertiesResponseDelegatedSubnetArguments {
		if v != nil {
			return *v
		}
		var ret ServerGroupPropertiesResponseDelegatedSubnetArguments
		return ret
	}).(ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput)
}

func (o ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroupPropertiesResponseDelegatedSubnetArguments) *string {
		if v == nil {
			return nil
		}
		return v.SubnetArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ServerGroupPropertiesResponsePrivateDnsZoneArguments struct {
	PrivateDnsZoneArmResourceId *string `pulumi:"privateDnsZoneArmResourceId"`
}

type ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput struct{ *pulumi.OutputState }

func (ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupPropertiesResponsePrivateDnsZoneArguments)(nil)).Elem()
}

func (o ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput) ToServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput() ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput {
	return o
}

func (o ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput) ToServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutputWithContext(ctx context.Context) ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput {
	return o
}

func (o ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupPropertiesResponsePrivateDnsZoneArguments) *string {
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput struct{ *pulumi.OutputState }

func (ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupPropertiesResponsePrivateDnsZoneArguments)(nil)).Elem()
}

func (o ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) ToServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput() ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o
}

func (o ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) ToServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutputWithContext(ctx context.Context) ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput {
	return o
}

func (o ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) Elem() ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput {
	return o.ApplyT(func(v *ServerGroupPropertiesResponsePrivateDnsZoneArguments) ServerGroupPropertiesResponsePrivateDnsZoneArguments {
		if v != nil {
			return *v
		}
		var ret ServerGroupPropertiesResponsePrivateDnsZoneArguments
		return ret
	}).(ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput)
}

func (o ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroupPropertiesResponsePrivateDnsZoneArguments) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ServerNameItemResponse struct {
	FullyQualifiedDomainName string  `pulumi:"fullyQualifiedDomainName"`
	Name                     *string `pulumi:"name"`
}

type ServerNameItemResponseOutput struct{ *pulumi.OutputState }

func (ServerNameItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerNameItemResponse)(nil)).Elem()
}

func (o ServerNameItemResponseOutput) ToServerNameItemResponseOutput() ServerNameItemResponseOutput {
	return o
}

func (o ServerNameItemResponseOutput) ToServerNameItemResponseOutputWithContext(ctx context.Context) ServerNameItemResponseOutput {
	return o
}

func (o ServerNameItemResponseOutput) FullyQualifiedDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServerNameItemResponse) string { return v.FullyQualifiedDomainName }).(pulumi.StringOutput)
}

func (o ServerNameItemResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerNameItemResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ServerNameItemResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerNameItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerNameItemResponse)(nil)).Elem()
}

func (o ServerNameItemResponseArrayOutput) ToServerNameItemResponseArrayOutput() ServerNameItemResponseArrayOutput {
	return o
}

func (o ServerNameItemResponseArrayOutput) ToServerNameItemResponseArrayOutputWithContext(ctx context.Context) ServerNameItemResponseArrayOutput {
	return o
}

func (o ServerNameItemResponseArrayOutput) Index(i pulumi.IntInput) ServerNameItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerNameItemResponse {
		return vs[0].([]ServerNameItemResponse)[vs[1].(int)]
	}).(ServerNameItemResponseOutput)
}

type ServerRoleGroup struct {
	EnableHa         *bool    `pulumi:"enableHa"`
	Name             *string  `pulumi:"name"`
	Role             *string  `pulumi:"role"`
	ServerCount      *int     `pulumi:"serverCount"`
	ServerEdition    *string  `pulumi:"serverEdition"`
	StorageQuotaInMb *float64 `pulumi:"storageQuotaInMb"`
	VCores           *float64 `pulumi:"vCores"`
}





type ServerRoleGroupInput interface {
	pulumi.Input

	ToServerRoleGroupOutput() ServerRoleGroupOutput
	ToServerRoleGroupOutputWithContext(context.Context) ServerRoleGroupOutput
}

type ServerRoleGroupArgs struct {
	EnableHa         pulumi.BoolPtrInput    `pulumi:"enableHa"`
	Name             pulumi.StringPtrInput  `pulumi:"name"`
	Role             pulumi.StringPtrInput  `pulumi:"role"`
	ServerCount      pulumi.IntPtrInput     `pulumi:"serverCount"`
	ServerEdition    pulumi.StringPtrInput  `pulumi:"serverEdition"`
	StorageQuotaInMb pulumi.Float64PtrInput `pulumi:"storageQuotaInMb"`
	VCores           pulumi.Float64PtrInput `pulumi:"vCores"`
}

func (ServerRoleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerRoleGroup)(nil)).Elem()
}

func (i ServerRoleGroupArgs) ToServerRoleGroupOutput() ServerRoleGroupOutput {
	return i.ToServerRoleGroupOutputWithContext(context.Background())
}

func (i ServerRoleGroupArgs) ToServerRoleGroupOutputWithContext(ctx context.Context) ServerRoleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRoleGroupOutput)
}





type ServerRoleGroupArrayInput interface {
	pulumi.Input

	ToServerRoleGroupArrayOutput() ServerRoleGroupArrayOutput
	ToServerRoleGroupArrayOutputWithContext(context.Context) ServerRoleGroupArrayOutput
}

type ServerRoleGroupArray []ServerRoleGroupInput

func (ServerRoleGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerRoleGroup)(nil)).Elem()
}

func (i ServerRoleGroupArray) ToServerRoleGroupArrayOutput() ServerRoleGroupArrayOutput {
	return i.ToServerRoleGroupArrayOutputWithContext(context.Background())
}

func (i ServerRoleGroupArray) ToServerRoleGroupArrayOutputWithContext(ctx context.Context) ServerRoleGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRoleGroupArrayOutput)
}

type ServerRoleGroupOutput struct{ *pulumi.OutputState }

func (ServerRoleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerRoleGroup)(nil)).Elem()
}

func (o ServerRoleGroupOutput) ToServerRoleGroupOutput() ServerRoleGroupOutput {
	return o
}

func (o ServerRoleGroupOutput) ToServerRoleGroupOutputWithContext(ctx context.Context) ServerRoleGroupOutput {
	return o
}

func (o ServerRoleGroupOutput) EnableHa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServerRoleGroup) *bool { return v.EnableHa }).(pulumi.BoolPtrOutput)
}

func (o ServerRoleGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerRoleGroup) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServerRoleGroupOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerRoleGroup) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o ServerRoleGroupOutput) ServerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServerRoleGroup) *int { return v.ServerCount }).(pulumi.IntPtrOutput)
}

func (o ServerRoleGroupOutput) ServerEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerRoleGroup) *string { return v.ServerEdition }).(pulumi.StringPtrOutput)
}

func (o ServerRoleGroupOutput) StorageQuotaInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ServerRoleGroup) *float64 { return v.StorageQuotaInMb }).(pulumi.Float64PtrOutput)
}

func (o ServerRoleGroupOutput) VCores() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ServerRoleGroup) *float64 { return v.VCores }).(pulumi.Float64PtrOutput)
}

type ServerRoleGroupArrayOutput struct{ *pulumi.OutputState }

func (ServerRoleGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerRoleGroup)(nil)).Elem()
}

func (o ServerRoleGroupArrayOutput) ToServerRoleGroupArrayOutput() ServerRoleGroupArrayOutput {
	return o
}

func (o ServerRoleGroupArrayOutput) ToServerRoleGroupArrayOutputWithContext(ctx context.Context) ServerRoleGroupArrayOutput {
	return o
}

func (o ServerRoleGroupArrayOutput) Index(i pulumi.IntInput) ServerRoleGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerRoleGroup {
		return vs[0].([]ServerRoleGroup)[vs[1].(int)]
	}).(ServerRoleGroupOutput)
}

type ServerRoleGroupResponse struct {
	EnableHa         *bool                    `pulumi:"enableHa"`
	EnablePublicIp   bool                     `pulumi:"enablePublicIp"`
	Name             *string                  `pulumi:"name"`
	Role             *string                  `pulumi:"role"`
	ServerCount      *int                     `pulumi:"serverCount"`
	ServerEdition    *string                  `pulumi:"serverEdition"`
	ServerNames      []ServerNameItemResponse `pulumi:"serverNames"`
	StorageQuotaInMb *float64                 `pulumi:"storageQuotaInMb"`
	VCores           *float64                 `pulumi:"vCores"`
}

type ServerRoleGroupResponseOutput struct{ *pulumi.OutputState }

func (ServerRoleGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerRoleGroupResponse)(nil)).Elem()
}

func (o ServerRoleGroupResponseOutput) ToServerRoleGroupResponseOutput() ServerRoleGroupResponseOutput {
	return o
}

func (o ServerRoleGroupResponseOutput) ToServerRoleGroupResponseOutputWithContext(ctx context.Context) ServerRoleGroupResponseOutput {
	return o
}

func (o ServerRoleGroupResponseOutput) EnableHa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) *bool { return v.EnableHa }).(pulumi.BoolPtrOutput)
}

func (o ServerRoleGroupResponseOutput) EnablePublicIp() pulumi.BoolOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) bool { return v.EnablePublicIp }).(pulumi.BoolOutput)
}

func (o ServerRoleGroupResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServerRoleGroupResponseOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o ServerRoleGroupResponseOutput) ServerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) *int { return v.ServerCount }).(pulumi.IntPtrOutput)
}

func (o ServerRoleGroupResponseOutput) ServerEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) *string { return v.ServerEdition }).(pulumi.StringPtrOutput)
}

func (o ServerRoleGroupResponseOutput) ServerNames() ServerNameItemResponseArrayOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) []ServerNameItemResponse { return v.ServerNames }).(ServerNameItemResponseArrayOutput)
}

func (o ServerRoleGroupResponseOutput) StorageQuotaInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) *float64 { return v.StorageQuotaInMb }).(pulumi.Float64PtrOutput)
}

func (o ServerRoleGroupResponseOutput) VCores() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ServerRoleGroupResponse) *float64 { return v.VCores }).(pulumi.Float64PtrOutput)
}

type ServerRoleGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerRoleGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerRoleGroupResponse)(nil)).Elem()
}

func (o ServerRoleGroupResponseArrayOutput) ToServerRoleGroupResponseArrayOutput() ServerRoleGroupResponseArrayOutput {
	return o
}

func (o ServerRoleGroupResponseArrayOutput) ToServerRoleGroupResponseArrayOutputWithContext(ctx context.Context) ServerRoleGroupResponseArrayOutput {
	return o
}

func (o ServerRoleGroupResponseArrayOutput) Index(i pulumi.IntInput) ServerRoleGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerRoleGroupResponse {
		return vs[0].([]ServerRoleGroupResponse)[vs[1].(int)]
	}).(ServerRoleGroupResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

func init() {
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerGroupPropertiesDelegatedSubnetArgumentsOutput{})
	pulumi.RegisterOutputType(ServerGroupPropertiesDelegatedSubnetArgumentsPtrOutput{})
	pulumi.RegisterOutputType(ServerGroupPropertiesPrivateDnsZoneArgumentsOutput{})
	pulumi.RegisterOutputType(ServerGroupPropertiesPrivateDnsZoneArgumentsPtrOutput{})
	pulumi.RegisterOutputType(ServerGroupPropertiesResponseDelegatedSubnetArgumentsOutput{})
	pulumi.RegisterOutputType(ServerGroupPropertiesResponseDelegatedSubnetArgumentsPtrOutput{})
	pulumi.RegisterOutputType(ServerGroupPropertiesResponsePrivateDnsZoneArgumentsOutput{})
	pulumi.RegisterOutputType(ServerGroupPropertiesResponsePrivateDnsZoneArgumentsPtrOutput{})
	pulumi.RegisterOutputType(ServerNameItemResponseOutput{})
	pulumi.RegisterOutputType(ServerNameItemResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerRoleGroupOutput{})
	pulumi.RegisterOutputType(ServerRoleGroupArrayOutput{})
	pulumi.RegisterOutputType(ServerRoleGroupResponseOutput{})
	pulumi.RegisterOutputType(ServerRoleGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
