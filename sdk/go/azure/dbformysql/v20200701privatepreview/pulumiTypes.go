


package v20200701privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DelegatedSubnetArguments struct {
	SubnetArmResourceId *string `pulumi:"subnetArmResourceId"`
}





type DelegatedSubnetArgumentsInput interface {
	pulumi.Input

	ToDelegatedSubnetArgumentsOutput() DelegatedSubnetArgumentsOutput
	ToDelegatedSubnetArgumentsOutputWithContext(context.Context) DelegatedSubnetArgumentsOutput
}

type DelegatedSubnetArgumentsArgs struct {
	SubnetArmResourceId pulumi.StringPtrInput `pulumi:"subnetArmResourceId"`
}

func (DelegatedSubnetArgumentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DelegatedSubnetArguments)(nil)).Elem()
}

func (i DelegatedSubnetArgumentsArgs) ToDelegatedSubnetArgumentsOutput() DelegatedSubnetArgumentsOutput {
	return i.ToDelegatedSubnetArgumentsOutputWithContext(context.Background())
}

func (i DelegatedSubnetArgumentsArgs) ToDelegatedSubnetArgumentsOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegatedSubnetArgumentsOutput)
}

func (i DelegatedSubnetArgumentsArgs) ToDelegatedSubnetArgumentsPtrOutput() DelegatedSubnetArgumentsPtrOutput {
	return i.ToDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (i DelegatedSubnetArgumentsArgs) ToDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegatedSubnetArgumentsOutput).ToDelegatedSubnetArgumentsPtrOutputWithContext(ctx)
}









type DelegatedSubnetArgumentsPtrInput interface {
	pulumi.Input

	ToDelegatedSubnetArgumentsPtrOutput() DelegatedSubnetArgumentsPtrOutput
	ToDelegatedSubnetArgumentsPtrOutputWithContext(context.Context) DelegatedSubnetArgumentsPtrOutput
}

type delegatedSubnetArgumentsPtrType DelegatedSubnetArgumentsArgs

func DelegatedSubnetArgumentsPtr(v *DelegatedSubnetArgumentsArgs) DelegatedSubnetArgumentsPtrInput {
	return (*delegatedSubnetArgumentsPtrType)(v)
}

func (*delegatedSubnetArgumentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DelegatedSubnetArguments)(nil)).Elem()
}

func (i *delegatedSubnetArgumentsPtrType) ToDelegatedSubnetArgumentsPtrOutput() DelegatedSubnetArgumentsPtrOutput {
	return i.ToDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (i *delegatedSubnetArgumentsPtrType) ToDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegatedSubnetArgumentsPtrOutput)
}

type DelegatedSubnetArgumentsOutput struct{ *pulumi.OutputState }

func (DelegatedSubnetArgumentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DelegatedSubnetArguments)(nil)).Elem()
}

func (o DelegatedSubnetArgumentsOutput) ToDelegatedSubnetArgumentsOutput() DelegatedSubnetArgumentsOutput {
	return o
}

func (o DelegatedSubnetArgumentsOutput) ToDelegatedSubnetArgumentsOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsOutput {
	return o
}

func (o DelegatedSubnetArgumentsOutput) ToDelegatedSubnetArgumentsPtrOutput() DelegatedSubnetArgumentsPtrOutput {
	return o.ToDelegatedSubnetArgumentsPtrOutputWithContext(context.Background())
}

func (o DelegatedSubnetArgumentsOutput) ToDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DelegatedSubnetArguments) *DelegatedSubnetArguments {
		return &v
	}).(DelegatedSubnetArgumentsPtrOutput)
}

func (o DelegatedSubnetArgumentsOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DelegatedSubnetArguments) *string { return v.SubnetArmResourceId }).(pulumi.StringPtrOutput)
}

type DelegatedSubnetArgumentsPtrOutput struct{ *pulumi.OutputState }

func (DelegatedSubnetArgumentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DelegatedSubnetArguments)(nil)).Elem()
}

func (o DelegatedSubnetArgumentsPtrOutput) ToDelegatedSubnetArgumentsPtrOutput() DelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o DelegatedSubnetArgumentsPtrOutput) ToDelegatedSubnetArgumentsPtrOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsPtrOutput {
	return o
}

func (o DelegatedSubnetArgumentsPtrOutput) Elem() DelegatedSubnetArgumentsOutput {
	return o.ApplyT(func(v *DelegatedSubnetArguments) DelegatedSubnetArguments {
		if v != nil {
			return *v
		}
		var ret DelegatedSubnetArguments
		return ret
	}).(DelegatedSubnetArgumentsOutput)
}

func (o DelegatedSubnetArgumentsPtrOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetArguments) *string {
		if v == nil {
			return nil
		}
		return v.SubnetArmResourceId
	}).(pulumi.StringPtrOutput)
}

type DelegatedSubnetArgumentsResponse struct {
	SubnetArmResourceId *string `pulumi:"subnetArmResourceId"`
}





type DelegatedSubnetArgumentsResponseInput interface {
	pulumi.Input

	ToDelegatedSubnetArgumentsResponseOutput() DelegatedSubnetArgumentsResponseOutput
	ToDelegatedSubnetArgumentsResponseOutputWithContext(context.Context) DelegatedSubnetArgumentsResponseOutput
}

type DelegatedSubnetArgumentsResponseArgs struct {
	SubnetArmResourceId pulumi.StringPtrInput `pulumi:"subnetArmResourceId"`
}

func (DelegatedSubnetArgumentsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DelegatedSubnetArgumentsResponse)(nil)).Elem()
}

func (i DelegatedSubnetArgumentsResponseArgs) ToDelegatedSubnetArgumentsResponseOutput() DelegatedSubnetArgumentsResponseOutput {
	return i.ToDelegatedSubnetArgumentsResponseOutputWithContext(context.Background())
}

func (i DelegatedSubnetArgumentsResponseArgs) ToDelegatedSubnetArgumentsResponseOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegatedSubnetArgumentsResponseOutput)
}

func (i DelegatedSubnetArgumentsResponseArgs) ToDelegatedSubnetArgumentsResponsePtrOutput() DelegatedSubnetArgumentsResponsePtrOutput {
	return i.ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(context.Background())
}

func (i DelegatedSubnetArgumentsResponseArgs) ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegatedSubnetArgumentsResponseOutput).ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(ctx)
}









type DelegatedSubnetArgumentsResponsePtrInput interface {
	pulumi.Input

	ToDelegatedSubnetArgumentsResponsePtrOutput() DelegatedSubnetArgumentsResponsePtrOutput
	ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(context.Context) DelegatedSubnetArgumentsResponsePtrOutput
}

type delegatedSubnetArgumentsResponsePtrType DelegatedSubnetArgumentsResponseArgs

func DelegatedSubnetArgumentsResponsePtr(v *DelegatedSubnetArgumentsResponseArgs) DelegatedSubnetArgumentsResponsePtrInput {
	return (*delegatedSubnetArgumentsResponsePtrType)(v)
}

func (*delegatedSubnetArgumentsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DelegatedSubnetArgumentsResponse)(nil)).Elem()
}

func (i *delegatedSubnetArgumentsResponsePtrType) ToDelegatedSubnetArgumentsResponsePtrOutput() DelegatedSubnetArgumentsResponsePtrOutput {
	return i.ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(context.Background())
}

func (i *delegatedSubnetArgumentsResponsePtrType) ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegatedSubnetArgumentsResponsePtrOutput)
}

type DelegatedSubnetArgumentsResponseOutput struct{ *pulumi.OutputState }

func (DelegatedSubnetArgumentsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DelegatedSubnetArgumentsResponse)(nil)).Elem()
}

func (o DelegatedSubnetArgumentsResponseOutput) ToDelegatedSubnetArgumentsResponseOutput() DelegatedSubnetArgumentsResponseOutput {
	return o
}

func (o DelegatedSubnetArgumentsResponseOutput) ToDelegatedSubnetArgumentsResponseOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsResponseOutput {
	return o
}

func (o DelegatedSubnetArgumentsResponseOutput) ToDelegatedSubnetArgumentsResponsePtrOutput() DelegatedSubnetArgumentsResponsePtrOutput {
	return o.ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(context.Background())
}

func (o DelegatedSubnetArgumentsResponseOutput) ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DelegatedSubnetArgumentsResponse) *DelegatedSubnetArgumentsResponse {
		return &v
	}).(DelegatedSubnetArgumentsResponsePtrOutput)
}

func (o DelegatedSubnetArgumentsResponseOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DelegatedSubnetArgumentsResponse) *string { return v.SubnetArmResourceId }).(pulumi.StringPtrOutput)
}

type DelegatedSubnetArgumentsResponsePtrOutput struct{ *pulumi.OutputState }

func (DelegatedSubnetArgumentsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DelegatedSubnetArgumentsResponse)(nil)).Elem()
}

func (o DelegatedSubnetArgumentsResponsePtrOutput) ToDelegatedSubnetArgumentsResponsePtrOutput() DelegatedSubnetArgumentsResponsePtrOutput {
	return o
}

func (o DelegatedSubnetArgumentsResponsePtrOutput) ToDelegatedSubnetArgumentsResponsePtrOutputWithContext(ctx context.Context) DelegatedSubnetArgumentsResponsePtrOutput {
	return o
}

func (o DelegatedSubnetArgumentsResponsePtrOutput) Elem() DelegatedSubnetArgumentsResponseOutput {
	return o.ApplyT(func(v *DelegatedSubnetArgumentsResponse) DelegatedSubnetArgumentsResponse {
		if v != nil {
			return *v
		}
		var ret DelegatedSubnetArgumentsResponse
		return ret
	}).(DelegatedSubnetArgumentsResponseOutput)
}

func (o DelegatedSubnetArgumentsResponsePtrOutput) SubnetArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetArgumentsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetArmResourceId
	}).(pulumi.StringPtrOutput)
}

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





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
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

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
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





type MaintenanceWindowResponseInput interface {
	pulumi.Input

	ToMaintenanceWindowResponseOutput() MaintenanceWindowResponseOutput
	ToMaintenanceWindowResponseOutputWithContext(context.Context) MaintenanceWindowResponseOutput
}

type MaintenanceWindowResponseArgs struct {
	CustomWindow pulumi.StringPtrInput `pulumi:"customWindow"`
	DayOfWeek    pulumi.IntPtrInput    `pulumi:"dayOfWeek"`
	StartHour    pulumi.IntPtrInput    `pulumi:"startHour"`
	StartMinute  pulumi.IntPtrInput    `pulumi:"startMinute"`
}

func (MaintenanceWindowResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowResponse)(nil)).Elem()
}

func (i MaintenanceWindowResponseArgs) ToMaintenanceWindowResponseOutput() MaintenanceWindowResponseOutput {
	return i.ToMaintenanceWindowResponseOutputWithContext(context.Background())
}

func (i MaintenanceWindowResponseArgs) ToMaintenanceWindowResponseOutputWithContext(ctx context.Context) MaintenanceWindowResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowResponseOutput)
}

func (i MaintenanceWindowResponseArgs) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return i.ToMaintenanceWindowResponsePtrOutputWithContext(context.Background())
}

func (i MaintenanceWindowResponseArgs) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowResponseOutput).ToMaintenanceWindowResponsePtrOutputWithContext(ctx)
}









type MaintenanceWindowResponsePtrInput interface {
	pulumi.Input

	ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput
	ToMaintenanceWindowResponsePtrOutputWithContext(context.Context) MaintenanceWindowResponsePtrOutput
}

type maintenanceWindowResponsePtrType MaintenanceWindowResponseArgs

func MaintenanceWindowResponsePtr(v *MaintenanceWindowResponseArgs) MaintenanceWindowResponsePtrInput {
	return (*maintenanceWindowResponsePtrType)(v)
}

func (*maintenanceWindowResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowResponse)(nil)).Elem()
}

func (i *maintenanceWindowResponsePtrType) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return i.ToMaintenanceWindowResponsePtrOutputWithContext(context.Background())
}

func (i *maintenanceWindowResponsePtrType) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowResponsePtrOutput)
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

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return o.ToMaintenanceWindowResponsePtrOutputWithContext(context.Background())
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MaintenanceWindowResponse) *MaintenanceWindowResponse {
		return &v
	}).(MaintenanceWindowResponsePtrOutput)
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





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
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
	BackupRetentionDays *int    `pulumi:"backupRetentionDays"`
	StorageAutogrow     *string `pulumi:"storageAutogrow"`
	StorageIops         *int    `pulumi:"storageIops"`
	StorageMB           *int    `pulumi:"storageMB"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	BackupRetentionDays pulumi.IntPtrInput    `pulumi:"backupRetentionDays"`
	StorageAutogrow     pulumi.StringPtrInput `pulumi:"storageAutogrow"`
	StorageIops         pulumi.IntPtrInput    `pulumi:"storageIops"`
	StorageMB           pulumi.IntPtrInput    `pulumi:"storageMB"`
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

func (o StorageProfileOutput) StorageAutogrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.StorageAutogrow }).(pulumi.StringPtrOutput)
}

func (o StorageProfileOutput) StorageIops() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.StorageIops }).(pulumi.IntPtrOutput)
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

func (o StorageProfilePtrOutput) StorageAutogrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.StorageAutogrow
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfilePtrOutput) StorageIops() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *int {
		if v == nil {
			return nil
		}
		return v.StorageIops
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
	BackupRetentionDays *int    `pulumi:"backupRetentionDays"`
	FileStorageSkuName  string  `pulumi:"fileStorageSkuName"`
	StorageAutogrow     *string `pulumi:"storageAutogrow"`
	StorageIops         *int    `pulumi:"storageIops"`
	StorageMB           *int    `pulumi:"storageMB"`
}





type StorageProfileResponseInput interface {
	pulumi.Input

	ToStorageProfileResponseOutput() StorageProfileResponseOutput
	ToStorageProfileResponseOutputWithContext(context.Context) StorageProfileResponseOutput
}

type StorageProfileResponseArgs struct {
	BackupRetentionDays pulumi.IntPtrInput    `pulumi:"backupRetentionDays"`
	FileStorageSkuName  pulumi.StringInput    `pulumi:"fileStorageSkuName"`
	StorageAutogrow     pulumi.StringPtrInput `pulumi:"storageAutogrow"`
	StorageIops         pulumi.IntPtrInput    `pulumi:"storageIops"`
	StorageMB           pulumi.IntPtrInput    `pulumi:"storageMB"`
}

func (StorageProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (i StorageProfileResponseArgs) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return i.ToStorageProfileResponseOutputWithContext(context.Background())
}

func (i StorageProfileResponseArgs) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponseOutput)
}

func (i StorageProfileResponseArgs) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return i.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (i StorageProfileResponseArgs) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponseOutput).ToStorageProfileResponsePtrOutputWithContext(ctx)
}









type StorageProfileResponsePtrInput interface {
	pulumi.Input

	ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput
	ToStorageProfileResponsePtrOutputWithContext(context.Context) StorageProfileResponsePtrOutput
}

type storageProfileResponsePtrType StorageProfileResponseArgs

func StorageProfileResponsePtr(v *StorageProfileResponseArgs) StorageProfileResponsePtrInput {
	return (*storageProfileResponsePtrType)(v)
}

func (*storageProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (i *storageProfileResponsePtrType) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return i.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (i *storageProfileResponsePtrType) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponsePtrOutput)
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

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfileResponse) *StorageProfileResponse {
		return &v
	}).(StorageProfileResponsePtrOutput)
}

func (o StorageProfileResponseOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponseOutput) FileStorageSkuName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageProfileResponse) string { return v.FileStorageSkuName }).(pulumi.StringOutput)
}

func (o StorageProfileResponseOutput) StorageAutogrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.StorageAutogrow }).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponseOutput) StorageIops() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.StorageIops }).(pulumi.IntPtrOutput)
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

func (o StorageProfileResponsePtrOutput) FileStorageSkuName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FileStorageSkuName
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponsePtrOutput) StorageAutogrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAutogrow
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponsePtrOutput) StorageIops() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.StorageIops
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
	pulumi.RegisterOutputType(DelegatedSubnetArgumentsOutput{})
	pulumi.RegisterOutputType(DelegatedSubnetArgumentsPtrOutput{})
	pulumi.RegisterOutputType(DelegatedSubnetArgumentsResponseOutput{})
	pulumi.RegisterOutputType(DelegatedSubnetArgumentsResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
}
