


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentUpdateProperties struct {
	MaintenanceWindowTimeZone *string                       `pulumi:"maintenanceWindowTimeZone"`
	MaintenanceWindows        []MaintenanceWindowProperties `pulumi:"maintenanceWindows"`
	Type                      *string                       `pulumi:"type"`
	UseSessionHostLocalTime   *bool                         `pulumi:"useSessionHostLocalTime"`
}





type AgentUpdatePropertiesInput interface {
	pulumi.Input

	ToAgentUpdatePropertiesOutput() AgentUpdatePropertiesOutput
	ToAgentUpdatePropertiesOutputWithContext(context.Context) AgentUpdatePropertiesOutput
}

type AgentUpdatePropertiesArgs struct {
	MaintenanceWindowTimeZone pulumi.StringPtrInput                 `pulumi:"maintenanceWindowTimeZone"`
	MaintenanceWindows        MaintenanceWindowPropertiesArrayInput `pulumi:"maintenanceWindows"`
	Type                      pulumi.StringPtrInput                 `pulumi:"type"`
	UseSessionHostLocalTime   pulumi.BoolPtrInput                   `pulumi:"useSessionHostLocalTime"`
}

func (AgentUpdatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentUpdateProperties)(nil)).Elem()
}

func (i AgentUpdatePropertiesArgs) ToAgentUpdatePropertiesOutput() AgentUpdatePropertiesOutput {
	return i.ToAgentUpdatePropertiesOutputWithContext(context.Background())
}

func (i AgentUpdatePropertiesArgs) ToAgentUpdatePropertiesOutputWithContext(ctx context.Context) AgentUpdatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentUpdatePropertiesOutput)
}

func (i AgentUpdatePropertiesArgs) ToAgentUpdatePropertiesPtrOutput() AgentUpdatePropertiesPtrOutput {
	return i.ToAgentUpdatePropertiesPtrOutputWithContext(context.Background())
}

func (i AgentUpdatePropertiesArgs) ToAgentUpdatePropertiesPtrOutputWithContext(ctx context.Context) AgentUpdatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentUpdatePropertiesOutput).ToAgentUpdatePropertiesPtrOutputWithContext(ctx)
}









type AgentUpdatePropertiesPtrInput interface {
	pulumi.Input

	ToAgentUpdatePropertiesPtrOutput() AgentUpdatePropertiesPtrOutput
	ToAgentUpdatePropertiesPtrOutputWithContext(context.Context) AgentUpdatePropertiesPtrOutput
}

type agentUpdatePropertiesPtrType AgentUpdatePropertiesArgs

func AgentUpdatePropertiesPtr(v *AgentUpdatePropertiesArgs) AgentUpdatePropertiesPtrInput {
	return (*agentUpdatePropertiesPtrType)(v)
}

func (*agentUpdatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentUpdateProperties)(nil)).Elem()
}

func (i *agentUpdatePropertiesPtrType) ToAgentUpdatePropertiesPtrOutput() AgentUpdatePropertiesPtrOutput {
	return i.ToAgentUpdatePropertiesPtrOutputWithContext(context.Background())
}

func (i *agentUpdatePropertiesPtrType) ToAgentUpdatePropertiesPtrOutputWithContext(ctx context.Context) AgentUpdatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentUpdatePropertiesPtrOutput)
}

type AgentUpdatePropertiesOutput struct{ *pulumi.OutputState }

func (AgentUpdatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentUpdateProperties)(nil)).Elem()
}

func (o AgentUpdatePropertiesOutput) ToAgentUpdatePropertiesOutput() AgentUpdatePropertiesOutput {
	return o
}

func (o AgentUpdatePropertiesOutput) ToAgentUpdatePropertiesOutputWithContext(ctx context.Context) AgentUpdatePropertiesOutput {
	return o
}

func (o AgentUpdatePropertiesOutput) ToAgentUpdatePropertiesPtrOutput() AgentUpdatePropertiesPtrOutput {
	return o.ToAgentUpdatePropertiesPtrOutputWithContext(context.Background())
}

func (o AgentUpdatePropertiesOutput) ToAgentUpdatePropertiesPtrOutputWithContext(ctx context.Context) AgentUpdatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgentUpdateProperties) *AgentUpdateProperties {
		return &v
	}).(AgentUpdatePropertiesPtrOutput)
}

func (o AgentUpdatePropertiesOutput) MaintenanceWindowTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentUpdateProperties) *string { return v.MaintenanceWindowTimeZone }).(pulumi.StringPtrOutput)
}

func (o AgentUpdatePropertiesOutput) MaintenanceWindows() MaintenanceWindowPropertiesArrayOutput {
	return o.ApplyT(func(v AgentUpdateProperties) []MaintenanceWindowProperties { return v.MaintenanceWindows }).(MaintenanceWindowPropertiesArrayOutput)
}

func (o AgentUpdatePropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentUpdateProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o AgentUpdatePropertiesOutput) UseSessionHostLocalTime() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AgentUpdateProperties) *bool { return v.UseSessionHostLocalTime }).(pulumi.BoolPtrOutput)
}

type AgentUpdatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AgentUpdatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentUpdateProperties)(nil)).Elem()
}

func (o AgentUpdatePropertiesPtrOutput) ToAgentUpdatePropertiesPtrOutput() AgentUpdatePropertiesPtrOutput {
	return o
}

func (o AgentUpdatePropertiesPtrOutput) ToAgentUpdatePropertiesPtrOutputWithContext(ctx context.Context) AgentUpdatePropertiesPtrOutput {
	return o
}

func (o AgentUpdatePropertiesPtrOutput) Elem() AgentUpdatePropertiesOutput {
	return o.ApplyT(func(v *AgentUpdateProperties) AgentUpdateProperties {
		if v != nil {
			return *v
		}
		var ret AgentUpdateProperties
		return ret
	}).(AgentUpdatePropertiesOutput)
}

func (o AgentUpdatePropertiesPtrOutput) MaintenanceWindowTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentUpdateProperties) *string {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowTimeZone
	}).(pulumi.StringPtrOutput)
}

func (o AgentUpdatePropertiesPtrOutput) MaintenanceWindows() MaintenanceWindowPropertiesArrayOutput {
	return o.ApplyT(func(v *AgentUpdateProperties) []MaintenanceWindowProperties {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindows
	}).(MaintenanceWindowPropertiesArrayOutput)
}

func (o AgentUpdatePropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentUpdateProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o AgentUpdatePropertiesPtrOutput) UseSessionHostLocalTime() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentUpdateProperties) *bool {
		if v == nil {
			return nil
		}
		return v.UseSessionHostLocalTime
	}).(pulumi.BoolPtrOutput)
}

type AgentUpdatePropertiesResponse struct {
	MaintenanceWindowTimeZone *string                               `pulumi:"maintenanceWindowTimeZone"`
	MaintenanceWindows        []MaintenanceWindowPropertiesResponse `pulumi:"maintenanceWindows"`
	Type                      *string                               `pulumi:"type"`
	UseSessionHostLocalTime   *bool                                 `pulumi:"useSessionHostLocalTime"`
}

type AgentUpdatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AgentUpdatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentUpdatePropertiesResponse)(nil)).Elem()
}

func (o AgentUpdatePropertiesResponseOutput) ToAgentUpdatePropertiesResponseOutput() AgentUpdatePropertiesResponseOutput {
	return o
}

func (o AgentUpdatePropertiesResponseOutput) ToAgentUpdatePropertiesResponseOutputWithContext(ctx context.Context) AgentUpdatePropertiesResponseOutput {
	return o
}

func (o AgentUpdatePropertiesResponseOutput) MaintenanceWindowTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentUpdatePropertiesResponse) *string { return v.MaintenanceWindowTimeZone }).(pulumi.StringPtrOutput)
}

func (o AgentUpdatePropertiesResponseOutput) MaintenanceWindows() MaintenanceWindowPropertiesResponseArrayOutput {
	return o.ApplyT(func(v AgentUpdatePropertiesResponse) []MaintenanceWindowPropertiesResponse {
		return v.MaintenanceWindows
	}).(MaintenanceWindowPropertiesResponseArrayOutput)
}

func (o AgentUpdatePropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentUpdatePropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o AgentUpdatePropertiesResponseOutput) UseSessionHostLocalTime() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AgentUpdatePropertiesResponse) *bool { return v.UseSessionHostLocalTime }).(pulumi.BoolPtrOutput)
}

type AgentUpdatePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AgentUpdatePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentUpdatePropertiesResponse)(nil)).Elem()
}

func (o AgentUpdatePropertiesResponsePtrOutput) ToAgentUpdatePropertiesResponsePtrOutput() AgentUpdatePropertiesResponsePtrOutput {
	return o
}

func (o AgentUpdatePropertiesResponsePtrOutput) ToAgentUpdatePropertiesResponsePtrOutputWithContext(ctx context.Context) AgentUpdatePropertiesResponsePtrOutput {
	return o
}

func (o AgentUpdatePropertiesResponsePtrOutput) Elem() AgentUpdatePropertiesResponseOutput {
	return o.ApplyT(func(v *AgentUpdatePropertiesResponse) AgentUpdatePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AgentUpdatePropertiesResponse
		return ret
	}).(AgentUpdatePropertiesResponseOutput)
}

func (o AgentUpdatePropertiesResponsePtrOutput) MaintenanceWindowTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentUpdatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowTimeZone
	}).(pulumi.StringPtrOutput)
}

func (o AgentUpdatePropertiesResponsePtrOutput) MaintenanceWindows() MaintenanceWindowPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *AgentUpdatePropertiesResponse) []MaintenanceWindowPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindows
	}).(MaintenanceWindowPropertiesResponseArrayOutput)
}

func (o AgentUpdatePropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentUpdatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o AgentUpdatePropertiesResponsePtrOutput) UseSessionHostLocalTime() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentUpdatePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseSessionHostLocalTime
	}).(pulumi.BoolPtrOutput)
}

type MaintenanceWindowProperties struct {
	DayOfWeek *DayOfWeek `pulumi:"dayOfWeek"`
	Hour      *int       `pulumi:"hour"`
}





type MaintenanceWindowPropertiesInput interface {
	pulumi.Input

	ToMaintenanceWindowPropertiesOutput() MaintenanceWindowPropertiesOutput
	ToMaintenanceWindowPropertiesOutputWithContext(context.Context) MaintenanceWindowPropertiesOutput
}

type MaintenanceWindowPropertiesArgs struct {
	DayOfWeek DayOfWeekPtrInput  `pulumi:"dayOfWeek"`
	Hour      pulumi.IntPtrInput `pulumi:"hour"`
}

func (MaintenanceWindowPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowProperties)(nil)).Elem()
}

func (i MaintenanceWindowPropertiesArgs) ToMaintenanceWindowPropertiesOutput() MaintenanceWindowPropertiesOutput {
	return i.ToMaintenanceWindowPropertiesOutputWithContext(context.Background())
}

func (i MaintenanceWindowPropertiesArgs) ToMaintenanceWindowPropertiesOutputWithContext(ctx context.Context) MaintenanceWindowPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowPropertiesOutput)
}





type MaintenanceWindowPropertiesArrayInput interface {
	pulumi.Input

	ToMaintenanceWindowPropertiesArrayOutput() MaintenanceWindowPropertiesArrayOutput
	ToMaintenanceWindowPropertiesArrayOutputWithContext(context.Context) MaintenanceWindowPropertiesArrayOutput
}

type MaintenanceWindowPropertiesArray []MaintenanceWindowPropertiesInput

func (MaintenanceWindowPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MaintenanceWindowProperties)(nil)).Elem()
}

func (i MaintenanceWindowPropertiesArray) ToMaintenanceWindowPropertiesArrayOutput() MaintenanceWindowPropertiesArrayOutput {
	return i.ToMaintenanceWindowPropertiesArrayOutputWithContext(context.Background())
}

func (i MaintenanceWindowPropertiesArray) ToMaintenanceWindowPropertiesArrayOutputWithContext(ctx context.Context) MaintenanceWindowPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowPropertiesArrayOutput)
}

type MaintenanceWindowPropertiesOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowProperties)(nil)).Elem()
}

func (o MaintenanceWindowPropertiesOutput) ToMaintenanceWindowPropertiesOutput() MaintenanceWindowPropertiesOutput {
	return o
}

func (o MaintenanceWindowPropertiesOutput) ToMaintenanceWindowPropertiesOutputWithContext(ctx context.Context) MaintenanceWindowPropertiesOutput {
	return o
}

func (o MaintenanceWindowPropertiesOutput) DayOfWeek() DayOfWeekPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowProperties) *DayOfWeek { return v.DayOfWeek }).(DayOfWeekPtrOutput)
}

func (o MaintenanceWindowPropertiesOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowProperties) *int { return v.Hour }).(pulumi.IntPtrOutput)
}

type MaintenanceWindowPropertiesArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MaintenanceWindowProperties)(nil)).Elem()
}

func (o MaintenanceWindowPropertiesArrayOutput) ToMaintenanceWindowPropertiesArrayOutput() MaintenanceWindowPropertiesArrayOutput {
	return o
}

func (o MaintenanceWindowPropertiesArrayOutput) ToMaintenanceWindowPropertiesArrayOutputWithContext(ctx context.Context) MaintenanceWindowPropertiesArrayOutput {
	return o
}

func (o MaintenanceWindowPropertiesArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MaintenanceWindowProperties {
		return vs[0].([]MaintenanceWindowProperties)[vs[1].(int)]
	}).(MaintenanceWindowPropertiesOutput)
}

type MaintenanceWindowPropertiesResponse struct {
	DayOfWeek *string `pulumi:"dayOfWeek"`
	Hour      *int    `pulumi:"hour"`
}

type MaintenanceWindowPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowPropertiesResponse)(nil)).Elem()
}

func (o MaintenanceWindowPropertiesResponseOutput) ToMaintenanceWindowPropertiesResponseOutput() MaintenanceWindowPropertiesResponseOutput {
	return o
}

func (o MaintenanceWindowPropertiesResponseOutput) ToMaintenanceWindowPropertiesResponseOutputWithContext(ctx context.Context) MaintenanceWindowPropertiesResponseOutput {
	return o
}

func (o MaintenanceWindowPropertiesResponseOutput) DayOfWeek() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowPropertiesResponse) *string { return v.DayOfWeek }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowPropertiesResponseOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowPropertiesResponse) *int { return v.Hour }).(pulumi.IntPtrOutput)
}

type MaintenanceWindowPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MaintenanceWindowPropertiesResponse)(nil)).Elem()
}

func (o MaintenanceWindowPropertiesResponseArrayOutput) ToMaintenanceWindowPropertiesResponseArrayOutput() MaintenanceWindowPropertiesResponseArrayOutput {
	return o
}

func (o MaintenanceWindowPropertiesResponseArrayOutput) ToMaintenanceWindowPropertiesResponseArrayOutputWithContext(ctx context.Context) MaintenanceWindowPropertiesResponseArrayOutput {
	return o
}

func (o MaintenanceWindowPropertiesResponseArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MaintenanceWindowPropertiesResponse {
		return vs[0].([]MaintenanceWindowPropertiesResponse)[vs[1].(int)]
	}).(MaintenanceWindowPropertiesResponseOutput)
}

type MigrationRequestProperties struct {
	MigrationPath *string `pulumi:"migrationPath"`
	Operation     *string `pulumi:"operation"`
}





type MigrationRequestPropertiesInput interface {
	pulumi.Input

	ToMigrationRequestPropertiesOutput() MigrationRequestPropertiesOutput
	ToMigrationRequestPropertiesOutputWithContext(context.Context) MigrationRequestPropertiesOutput
}

type MigrationRequestPropertiesArgs struct {
	MigrationPath pulumi.StringPtrInput `pulumi:"migrationPath"`
	Operation     pulumi.StringPtrInput `pulumi:"operation"`
}

func (MigrationRequestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationRequestProperties)(nil)).Elem()
}

func (i MigrationRequestPropertiesArgs) ToMigrationRequestPropertiesOutput() MigrationRequestPropertiesOutput {
	return i.ToMigrationRequestPropertiesOutputWithContext(context.Background())
}

func (i MigrationRequestPropertiesArgs) ToMigrationRequestPropertiesOutputWithContext(ctx context.Context) MigrationRequestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationRequestPropertiesOutput)
}

func (i MigrationRequestPropertiesArgs) ToMigrationRequestPropertiesPtrOutput() MigrationRequestPropertiesPtrOutput {
	return i.ToMigrationRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i MigrationRequestPropertiesArgs) ToMigrationRequestPropertiesPtrOutputWithContext(ctx context.Context) MigrationRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationRequestPropertiesOutput).ToMigrationRequestPropertiesPtrOutputWithContext(ctx)
}









type MigrationRequestPropertiesPtrInput interface {
	pulumi.Input

	ToMigrationRequestPropertiesPtrOutput() MigrationRequestPropertiesPtrOutput
	ToMigrationRequestPropertiesPtrOutputWithContext(context.Context) MigrationRequestPropertiesPtrOutput
}

type migrationRequestPropertiesPtrType MigrationRequestPropertiesArgs

func MigrationRequestPropertiesPtr(v *MigrationRequestPropertiesArgs) MigrationRequestPropertiesPtrInput {
	return (*migrationRequestPropertiesPtrType)(v)
}

func (*migrationRequestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationRequestProperties)(nil)).Elem()
}

func (i *migrationRequestPropertiesPtrType) ToMigrationRequestPropertiesPtrOutput() MigrationRequestPropertiesPtrOutput {
	return i.ToMigrationRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i *migrationRequestPropertiesPtrType) ToMigrationRequestPropertiesPtrOutputWithContext(ctx context.Context) MigrationRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationRequestPropertiesPtrOutput)
}

type MigrationRequestPropertiesOutput struct{ *pulumi.OutputState }

func (MigrationRequestPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationRequestProperties)(nil)).Elem()
}

func (o MigrationRequestPropertiesOutput) ToMigrationRequestPropertiesOutput() MigrationRequestPropertiesOutput {
	return o
}

func (o MigrationRequestPropertiesOutput) ToMigrationRequestPropertiesOutputWithContext(ctx context.Context) MigrationRequestPropertiesOutput {
	return o
}

func (o MigrationRequestPropertiesOutput) ToMigrationRequestPropertiesPtrOutput() MigrationRequestPropertiesPtrOutput {
	return o.ToMigrationRequestPropertiesPtrOutputWithContext(context.Background())
}

func (o MigrationRequestPropertiesOutput) ToMigrationRequestPropertiesPtrOutputWithContext(ctx context.Context) MigrationRequestPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationRequestProperties) *MigrationRequestProperties {
		return &v
	}).(MigrationRequestPropertiesPtrOutput)
}

func (o MigrationRequestPropertiesOutput) MigrationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationRequestProperties) *string { return v.MigrationPath }).(pulumi.StringPtrOutput)
}

func (o MigrationRequestPropertiesOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationRequestProperties) *string { return v.Operation }).(pulumi.StringPtrOutput)
}

type MigrationRequestPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MigrationRequestPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationRequestProperties)(nil)).Elem()
}

func (o MigrationRequestPropertiesPtrOutput) ToMigrationRequestPropertiesPtrOutput() MigrationRequestPropertiesPtrOutput {
	return o
}

func (o MigrationRequestPropertiesPtrOutput) ToMigrationRequestPropertiesPtrOutputWithContext(ctx context.Context) MigrationRequestPropertiesPtrOutput {
	return o
}

func (o MigrationRequestPropertiesPtrOutput) Elem() MigrationRequestPropertiesOutput {
	return o.ApplyT(func(v *MigrationRequestProperties) MigrationRequestProperties {
		if v != nil {
			return *v
		}
		var ret MigrationRequestProperties
		return ret
	}).(MigrationRequestPropertiesOutput)
}

func (o MigrationRequestPropertiesPtrOutput) MigrationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.MigrationPath
	}).(pulumi.StringPtrOutput)
}

func (o MigrationRequestPropertiesPtrOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.Operation
	}).(pulumi.StringPtrOutput)
}

type MigrationRequestPropertiesResponse struct {
	MigrationPath *string `pulumi:"migrationPath"`
	Operation     *string `pulumi:"operation"`
}

type MigrationRequestPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MigrationRequestPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationRequestPropertiesResponse)(nil)).Elem()
}

func (o MigrationRequestPropertiesResponseOutput) ToMigrationRequestPropertiesResponseOutput() MigrationRequestPropertiesResponseOutput {
	return o
}

func (o MigrationRequestPropertiesResponseOutput) ToMigrationRequestPropertiesResponseOutputWithContext(ctx context.Context) MigrationRequestPropertiesResponseOutput {
	return o
}

func (o MigrationRequestPropertiesResponseOutput) MigrationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationRequestPropertiesResponse) *string { return v.MigrationPath }).(pulumi.StringPtrOutput)
}

func (o MigrationRequestPropertiesResponseOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrationRequestPropertiesResponse) *string { return v.Operation }).(pulumi.StringPtrOutput)
}

type MigrationRequestPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrationRequestPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationRequestPropertiesResponse)(nil)).Elem()
}

func (o MigrationRequestPropertiesResponsePtrOutput) ToMigrationRequestPropertiesResponsePtrOutput() MigrationRequestPropertiesResponsePtrOutput {
	return o
}

func (o MigrationRequestPropertiesResponsePtrOutput) ToMigrationRequestPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationRequestPropertiesResponsePtrOutput {
	return o
}

func (o MigrationRequestPropertiesResponsePtrOutput) Elem() MigrationRequestPropertiesResponseOutput {
	return o.ApplyT(func(v *MigrationRequestPropertiesResponse) MigrationRequestPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MigrationRequestPropertiesResponse
		return ret
	}).(MigrationRequestPropertiesResponseOutput)
}

func (o MigrationRequestPropertiesResponsePtrOutput) MigrationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationRequestPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MigrationPath
	}).(pulumi.StringPtrOutput)
}

func (o MigrationRequestPropertiesResponsePtrOutput) Operation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationRequestPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Operation
	}).(pulumi.StringPtrOutput)
}

type MsixPackageApplications struct {
	AppId          *string `pulumi:"appId"`
	AppUserModelID *string `pulumi:"appUserModelID"`
	Description    *string `pulumi:"description"`
	FriendlyName   *string `pulumi:"friendlyName"`
	IconImageName  *string `pulumi:"iconImageName"`
	RawIcon        *string `pulumi:"rawIcon"`
	RawPng         *string `pulumi:"rawPng"`
}





type MsixPackageApplicationsInput interface {
	pulumi.Input

	ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput
	ToMsixPackageApplicationsOutputWithContext(context.Context) MsixPackageApplicationsOutput
}

type MsixPackageApplicationsArgs struct {
	AppId          pulumi.StringPtrInput `pulumi:"appId"`
	AppUserModelID pulumi.StringPtrInput `pulumi:"appUserModelID"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	FriendlyName   pulumi.StringPtrInput `pulumi:"friendlyName"`
	IconImageName  pulumi.StringPtrInput `pulumi:"iconImageName"`
	RawIcon        pulumi.StringPtrInput `pulumi:"rawIcon"`
	RawPng         pulumi.StringPtrInput `pulumi:"rawPng"`
}

func (MsixPackageApplicationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplications)(nil)).Elem()
}

func (i MsixPackageApplicationsArgs) ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput {
	return i.ToMsixPackageApplicationsOutputWithContext(context.Background())
}

func (i MsixPackageApplicationsArgs) ToMsixPackageApplicationsOutputWithContext(ctx context.Context) MsixPackageApplicationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageApplicationsOutput)
}





type MsixPackageApplicationsArrayInput interface {
	pulumi.Input

	ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput
	ToMsixPackageApplicationsArrayOutputWithContext(context.Context) MsixPackageApplicationsArrayOutput
}

type MsixPackageApplicationsArray []MsixPackageApplicationsInput

func (MsixPackageApplicationsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplications)(nil)).Elem()
}

func (i MsixPackageApplicationsArray) ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput {
	return i.ToMsixPackageApplicationsArrayOutputWithContext(context.Background())
}

func (i MsixPackageApplicationsArray) ToMsixPackageApplicationsArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageApplicationsArrayOutput)
}

type MsixPackageApplicationsOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplications)(nil)).Elem()
}

func (o MsixPackageApplicationsOutput) ToMsixPackageApplicationsOutput() MsixPackageApplicationsOutput {
	return o
}

func (o MsixPackageApplicationsOutput) ToMsixPackageApplicationsOutputWithContext(ctx context.Context) MsixPackageApplicationsOutput {
	return o
}

func (o MsixPackageApplicationsOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) AppUserModelID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.AppUserModelID }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) IconImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.IconImageName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) RawIcon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.RawIcon }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsOutput) RawPng() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplications) *string { return v.RawPng }).(pulumi.StringPtrOutput)
}

type MsixPackageApplicationsArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplications)(nil)).Elem()
}

func (o MsixPackageApplicationsArrayOutput) ToMsixPackageApplicationsArrayOutput() MsixPackageApplicationsArrayOutput {
	return o
}

func (o MsixPackageApplicationsArrayOutput) ToMsixPackageApplicationsArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsArrayOutput {
	return o
}

func (o MsixPackageApplicationsArrayOutput) Index(i pulumi.IntInput) MsixPackageApplicationsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageApplications {
		return vs[0].([]MsixPackageApplications)[vs[1].(int)]
	}).(MsixPackageApplicationsOutput)
}

type MsixPackageApplicationsResponse struct {
	AppId          *string `pulumi:"appId"`
	AppUserModelID *string `pulumi:"appUserModelID"`
	Description    *string `pulumi:"description"`
	FriendlyName   *string `pulumi:"friendlyName"`
	IconImageName  *string `pulumi:"iconImageName"`
	RawIcon        *string `pulumi:"rawIcon"`
	RawPng         *string `pulumi:"rawPng"`
}

type MsixPackageApplicationsResponseOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageApplicationsResponse)(nil)).Elem()
}

func (o MsixPackageApplicationsResponseOutput) ToMsixPackageApplicationsResponseOutput() MsixPackageApplicationsResponseOutput {
	return o
}

func (o MsixPackageApplicationsResponseOutput) ToMsixPackageApplicationsResponseOutputWithContext(ctx context.Context) MsixPackageApplicationsResponseOutput {
	return o
}

func (o MsixPackageApplicationsResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) AppUserModelID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.AppUserModelID }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) IconImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.IconImageName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) RawIcon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.RawIcon }).(pulumi.StringPtrOutput)
}

func (o MsixPackageApplicationsResponseOutput) RawPng() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageApplicationsResponse) *string { return v.RawPng }).(pulumi.StringPtrOutput)
}

type MsixPackageApplicationsResponseArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageApplicationsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageApplicationsResponse)(nil)).Elem()
}

func (o MsixPackageApplicationsResponseArrayOutput) ToMsixPackageApplicationsResponseArrayOutput() MsixPackageApplicationsResponseArrayOutput {
	return o
}

func (o MsixPackageApplicationsResponseArrayOutput) ToMsixPackageApplicationsResponseArrayOutputWithContext(ctx context.Context) MsixPackageApplicationsResponseArrayOutput {
	return o
}

func (o MsixPackageApplicationsResponseArrayOutput) Index(i pulumi.IntInput) MsixPackageApplicationsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageApplicationsResponse {
		return vs[0].([]MsixPackageApplicationsResponse)[vs[1].(int)]
	}).(MsixPackageApplicationsResponseOutput)
}

type MsixPackageDependencies struct {
	DependencyName *string `pulumi:"dependencyName"`
	MinVersion     *string `pulumi:"minVersion"`
	Publisher      *string `pulumi:"publisher"`
}





type MsixPackageDependenciesInput interface {
	pulumi.Input

	ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput
	ToMsixPackageDependenciesOutputWithContext(context.Context) MsixPackageDependenciesOutput
}

type MsixPackageDependenciesArgs struct {
	DependencyName pulumi.StringPtrInput `pulumi:"dependencyName"`
	MinVersion     pulumi.StringPtrInput `pulumi:"minVersion"`
	Publisher      pulumi.StringPtrInput `pulumi:"publisher"`
}

func (MsixPackageDependenciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependencies)(nil)).Elem()
}

func (i MsixPackageDependenciesArgs) ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput {
	return i.ToMsixPackageDependenciesOutputWithContext(context.Background())
}

func (i MsixPackageDependenciesArgs) ToMsixPackageDependenciesOutputWithContext(ctx context.Context) MsixPackageDependenciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageDependenciesOutput)
}





type MsixPackageDependenciesArrayInput interface {
	pulumi.Input

	ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput
	ToMsixPackageDependenciesArrayOutputWithContext(context.Context) MsixPackageDependenciesArrayOutput
}

type MsixPackageDependenciesArray []MsixPackageDependenciesInput

func (MsixPackageDependenciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependencies)(nil)).Elem()
}

func (i MsixPackageDependenciesArray) ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput {
	return i.ToMsixPackageDependenciesArrayOutputWithContext(context.Background())
}

func (i MsixPackageDependenciesArray) ToMsixPackageDependenciesArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsixPackageDependenciesArrayOutput)
}

type MsixPackageDependenciesOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependencies)(nil)).Elem()
}

func (o MsixPackageDependenciesOutput) ToMsixPackageDependenciesOutput() MsixPackageDependenciesOutput {
	return o
}

func (o MsixPackageDependenciesOutput) ToMsixPackageDependenciesOutputWithContext(ctx context.Context) MsixPackageDependenciesOutput {
	return o
}

func (o MsixPackageDependenciesOutput) DependencyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.DependencyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesOutput) MinVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.MinVersion }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependencies) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type MsixPackageDependenciesArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependencies)(nil)).Elem()
}

func (o MsixPackageDependenciesArrayOutput) ToMsixPackageDependenciesArrayOutput() MsixPackageDependenciesArrayOutput {
	return o
}

func (o MsixPackageDependenciesArrayOutput) ToMsixPackageDependenciesArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesArrayOutput {
	return o
}

func (o MsixPackageDependenciesArrayOutput) Index(i pulumi.IntInput) MsixPackageDependenciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageDependencies {
		return vs[0].([]MsixPackageDependencies)[vs[1].(int)]
	}).(MsixPackageDependenciesOutput)
}

type MsixPackageDependenciesResponse struct {
	DependencyName *string `pulumi:"dependencyName"`
	MinVersion     *string `pulumi:"minVersion"`
	Publisher      *string `pulumi:"publisher"`
}

type MsixPackageDependenciesResponseOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsixPackageDependenciesResponse)(nil)).Elem()
}

func (o MsixPackageDependenciesResponseOutput) ToMsixPackageDependenciesResponseOutput() MsixPackageDependenciesResponseOutput {
	return o
}

func (o MsixPackageDependenciesResponseOutput) ToMsixPackageDependenciesResponseOutputWithContext(ctx context.Context) MsixPackageDependenciesResponseOutput {
	return o
}

func (o MsixPackageDependenciesResponseOutput) DependencyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.DependencyName }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesResponseOutput) MinVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.MinVersion }).(pulumi.StringPtrOutput)
}

func (o MsixPackageDependenciesResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsixPackageDependenciesResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type MsixPackageDependenciesResponseArrayOutput struct{ *pulumi.OutputState }

func (MsixPackageDependenciesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MsixPackageDependenciesResponse)(nil)).Elem()
}

func (o MsixPackageDependenciesResponseArrayOutput) ToMsixPackageDependenciesResponseArrayOutput() MsixPackageDependenciesResponseArrayOutput {
	return o
}

func (o MsixPackageDependenciesResponseArrayOutput) ToMsixPackageDependenciesResponseArrayOutputWithContext(ctx context.Context) MsixPackageDependenciesResponseArrayOutput {
	return o
}

func (o MsixPackageDependenciesResponseArrayOutput) Index(i pulumi.IntInput) MsixPackageDependenciesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MsixPackageDependenciesResponse {
		return vs[0].([]MsixPackageDependenciesResponse)[vs[1].(int)]
	}).(MsixPackageDependenciesResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RegistrationInfo struct {
	ExpirationTime             *string `pulumi:"expirationTime"`
	RegistrationTokenOperation *string `pulumi:"registrationTokenOperation"`
	Token                      *string `pulumi:"token"`
}





type RegistrationInfoInput interface {
	pulumi.Input

	ToRegistrationInfoOutput() RegistrationInfoOutput
	ToRegistrationInfoOutputWithContext(context.Context) RegistrationInfoOutput
}

type RegistrationInfoArgs struct {
	ExpirationTime             pulumi.StringPtrInput `pulumi:"expirationTime"`
	RegistrationTokenOperation pulumi.StringPtrInput `pulumi:"registrationTokenOperation"`
	Token                      pulumi.StringPtrInput `pulumi:"token"`
}

func (RegistrationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfo)(nil)).Elem()
}

func (i RegistrationInfoArgs) ToRegistrationInfoOutput() RegistrationInfoOutput {
	return i.ToRegistrationInfoOutputWithContext(context.Background())
}

func (i RegistrationInfoArgs) ToRegistrationInfoOutputWithContext(ctx context.Context) RegistrationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoOutput)
}

func (i RegistrationInfoArgs) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return i.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (i RegistrationInfoArgs) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoOutput).ToRegistrationInfoPtrOutputWithContext(ctx)
}









type RegistrationInfoPtrInput interface {
	pulumi.Input

	ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput
	ToRegistrationInfoPtrOutputWithContext(context.Context) RegistrationInfoPtrOutput
}

type registrationInfoPtrType RegistrationInfoArgs

func RegistrationInfoPtr(v *RegistrationInfoArgs) RegistrationInfoPtrInput {
	return (*registrationInfoPtrType)(v)
}

func (*registrationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfo)(nil)).Elem()
}

func (i *registrationInfoPtrType) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return i.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (i *registrationInfoPtrType) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationInfoPtrOutput)
}

type RegistrationInfoOutput struct{ *pulumi.OutputState }

func (RegistrationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfo)(nil)).Elem()
}

func (o RegistrationInfoOutput) ToRegistrationInfoOutput() RegistrationInfoOutput {
	return o
}

func (o RegistrationInfoOutput) ToRegistrationInfoOutputWithContext(ctx context.Context) RegistrationInfoOutput {
	return o
}

func (o RegistrationInfoOutput) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return o.ToRegistrationInfoPtrOutputWithContext(context.Background())
}

func (o RegistrationInfoOutput) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationInfo) *RegistrationInfo {
		return &v
	}).(RegistrationInfoPtrOutput)
}

func (o RegistrationInfoOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.ExpirationTime }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.RegistrationTokenOperation }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfo) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type RegistrationInfoPtrOutput struct{ *pulumi.OutputState }

func (RegistrationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfo)(nil)).Elem()
}

func (o RegistrationInfoPtrOutput) ToRegistrationInfoPtrOutput() RegistrationInfoPtrOutput {
	return o
}

func (o RegistrationInfoPtrOutput) ToRegistrationInfoPtrOutputWithContext(ctx context.Context) RegistrationInfoPtrOutput {
	return o
}

func (o RegistrationInfoPtrOutput) Elem() RegistrationInfoOutput {
	return o.ApplyT(func(v *RegistrationInfo) RegistrationInfo {
		if v != nil {
			return *v
		}
		var ret RegistrationInfo
		return ret
	}).(RegistrationInfoOutput)
}

func (o RegistrationInfoPtrOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoPtrOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationTokenOperation
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoPtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type RegistrationInfoResponse struct {
	ExpirationTime             *string `pulumi:"expirationTime"`
	RegistrationTokenOperation *string `pulumi:"registrationTokenOperation"`
	Token                      *string `pulumi:"token"`
}

type RegistrationInfoResponseOutput struct{ *pulumi.OutputState }

func (RegistrationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationInfoResponse)(nil)).Elem()
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponseOutput() RegistrationInfoResponseOutput {
	return o
}

func (o RegistrationInfoResponseOutput) ToRegistrationInfoResponseOutputWithContext(ctx context.Context) RegistrationInfoResponseOutput {
	return o
}

func (o RegistrationInfoResponseOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.ExpirationTime }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponseOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.RegistrationTokenOperation }).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponseOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationInfoResponse) *string { return v.Token }).(pulumi.StringPtrOutput)
}

type RegistrationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (RegistrationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationInfoResponse)(nil)).Elem()
}

func (o RegistrationInfoResponsePtrOutput) ToRegistrationInfoResponsePtrOutput() RegistrationInfoResponsePtrOutput {
	return o
}

func (o RegistrationInfoResponsePtrOutput) ToRegistrationInfoResponsePtrOutputWithContext(ctx context.Context) RegistrationInfoResponsePtrOutput {
	return o
}

func (o RegistrationInfoResponsePtrOutput) Elem() RegistrationInfoResponseOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) RegistrationInfoResponse {
		if v != nil {
			return *v
		}
		var ret RegistrationInfoResponse
		return ret
	}).(RegistrationInfoResponseOutput)
}

func (o RegistrationInfoResponsePtrOutput) ExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponsePtrOutput) RegistrationTokenOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationTokenOperation
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationInfoResponsePtrOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Token
	}).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type ResourceModelWithAllowedPropertySetIdentityInput interface {
	pulumi.Input

	ToResourceModelWithAllowedPropertySetIdentityOutput() ResourceModelWithAllowedPropertySetIdentityOutput
	ToResourceModelWithAllowedPropertySetIdentityOutputWithContext(context.Context) ResourceModelWithAllowedPropertySetIdentityOutput
}

type ResourceModelWithAllowedPropertySetIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (ResourceModelWithAllowedPropertySetIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetIdentity)(nil)).Elem()
}

func (i ResourceModelWithAllowedPropertySetIdentityArgs) ToResourceModelWithAllowedPropertySetIdentityOutput() ResourceModelWithAllowedPropertySetIdentityOutput {
	return i.ToResourceModelWithAllowedPropertySetIdentityOutputWithContext(context.Background())
}

func (i ResourceModelWithAllowedPropertySetIdentityArgs) ToResourceModelWithAllowedPropertySetIdentityOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetIdentityOutput)
}

func (i ResourceModelWithAllowedPropertySetIdentityArgs) ToResourceModelWithAllowedPropertySetIdentityPtrOutput() ResourceModelWithAllowedPropertySetIdentityPtrOutput {
	return i.ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceModelWithAllowedPropertySetIdentityArgs) ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetIdentityOutput).ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(ctx)
}









type ResourceModelWithAllowedPropertySetIdentityPtrInput interface {
	pulumi.Input

	ToResourceModelWithAllowedPropertySetIdentityPtrOutput() ResourceModelWithAllowedPropertySetIdentityPtrOutput
	ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(context.Context) ResourceModelWithAllowedPropertySetIdentityPtrOutput
}

type resourceModelWithAllowedPropertySetIdentityPtrType ResourceModelWithAllowedPropertySetIdentityArgs

func ResourceModelWithAllowedPropertySetIdentityPtr(v *ResourceModelWithAllowedPropertySetIdentityArgs) ResourceModelWithAllowedPropertySetIdentityPtrInput {
	return (*resourceModelWithAllowedPropertySetIdentityPtrType)(v)
}

func (*resourceModelWithAllowedPropertySetIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetIdentity)(nil)).Elem()
}

func (i *resourceModelWithAllowedPropertySetIdentityPtrType) ToResourceModelWithAllowedPropertySetIdentityPtrOutput() ResourceModelWithAllowedPropertySetIdentityPtrOutput {
	return i.ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceModelWithAllowedPropertySetIdentityPtrType) ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetIdentityPtrOutput)
}

type ResourceModelWithAllowedPropertySetIdentityOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetIdentity)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetIdentityOutput) ToResourceModelWithAllowedPropertySetIdentityOutput() ResourceModelWithAllowedPropertySetIdentityOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetIdentityOutput) ToResourceModelWithAllowedPropertySetIdentityOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetIdentityOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetIdentityOutput) ToResourceModelWithAllowedPropertySetIdentityPtrOutput() ResourceModelWithAllowedPropertySetIdentityPtrOutput {
	return o.ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceModelWithAllowedPropertySetIdentityOutput) ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceModelWithAllowedPropertySetIdentity) *ResourceModelWithAllowedPropertySetIdentity {
		return &v
	}).(ResourceModelWithAllowedPropertySetIdentityPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type ResourceModelWithAllowedPropertySetIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetIdentity)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetIdentityPtrOutput) ToResourceModelWithAllowedPropertySetIdentityPtrOutput() ResourceModelWithAllowedPropertySetIdentityPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetIdentityPtrOutput) ToResourceModelWithAllowedPropertySetIdentityPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetIdentityPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetIdentityPtrOutput) Elem() ResourceModelWithAllowedPropertySetIdentityOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetIdentity) ResourceModelWithAllowedPropertySetIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceModelWithAllowedPropertySetIdentity
		return ret
	}).(ResourceModelWithAllowedPropertySetIdentityOutput)
}

func (o ResourceModelWithAllowedPropertySetIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ResourceModelWithAllowedPropertySetPlan struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}





type ResourceModelWithAllowedPropertySetPlanInput interface {
	pulumi.Input

	ToResourceModelWithAllowedPropertySetPlanOutput() ResourceModelWithAllowedPropertySetPlanOutput
	ToResourceModelWithAllowedPropertySetPlanOutputWithContext(context.Context) ResourceModelWithAllowedPropertySetPlanOutput
}

type ResourceModelWithAllowedPropertySetPlanArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	Product       pulumi.StringInput    `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringInput    `pulumi:"publisher"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (ResourceModelWithAllowedPropertySetPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetPlan)(nil)).Elem()
}

func (i ResourceModelWithAllowedPropertySetPlanArgs) ToResourceModelWithAllowedPropertySetPlanOutput() ResourceModelWithAllowedPropertySetPlanOutput {
	return i.ToResourceModelWithAllowedPropertySetPlanOutputWithContext(context.Background())
}

func (i ResourceModelWithAllowedPropertySetPlanArgs) ToResourceModelWithAllowedPropertySetPlanOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetPlanOutput)
}

func (i ResourceModelWithAllowedPropertySetPlanArgs) ToResourceModelWithAllowedPropertySetPlanPtrOutput() ResourceModelWithAllowedPropertySetPlanPtrOutput {
	return i.ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(context.Background())
}

func (i ResourceModelWithAllowedPropertySetPlanArgs) ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetPlanOutput).ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(ctx)
}









type ResourceModelWithAllowedPropertySetPlanPtrInput interface {
	pulumi.Input

	ToResourceModelWithAllowedPropertySetPlanPtrOutput() ResourceModelWithAllowedPropertySetPlanPtrOutput
	ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(context.Context) ResourceModelWithAllowedPropertySetPlanPtrOutput
}

type resourceModelWithAllowedPropertySetPlanPtrType ResourceModelWithAllowedPropertySetPlanArgs

func ResourceModelWithAllowedPropertySetPlanPtr(v *ResourceModelWithAllowedPropertySetPlanArgs) ResourceModelWithAllowedPropertySetPlanPtrInput {
	return (*resourceModelWithAllowedPropertySetPlanPtrType)(v)
}

func (*resourceModelWithAllowedPropertySetPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetPlan)(nil)).Elem()
}

func (i *resourceModelWithAllowedPropertySetPlanPtrType) ToResourceModelWithAllowedPropertySetPlanPtrOutput() ResourceModelWithAllowedPropertySetPlanPtrOutput {
	return i.ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(context.Background())
}

func (i *resourceModelWithAllowedPropertySetPlanPtrType) ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetPlanPtrOutput)
}

type ResourceModelWithAllowedPropertySetPlanOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetPlan)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) ToResourceModelWithAllowedPropertySetPlanOutput() ResourceModelWithAllowedPropertySetPlanOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) ToResourceModelWithAllowedPropertySetPlanOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetPlanOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) ToResourceModelWithAllowedPropertySetPlanPtrOutput() ResourceModelWithAllowedPropertySetPlanPtrOutput {
	return o.ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(context.Background())
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetPlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceModelWithAllowedPropertySetPlan) *ResourceModelWithAllowedPropertySetPlan {
		return &v
	}).(ResourceModelWithAllowedPropertySetPlanPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetPlan) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetPlan) string { return v.Product }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetPlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetPlan) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetPlan) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetPlanPtrOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetPlan)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetPlanPtrOutput) ToResourceModelWithAllowedPropertySetPlanPtrOutput() ResourceModelWithAllowedPropertySetPlanPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetPlanPtrOutput) ToResourceModelWithAllowedPropertySetPlanPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetPlanPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetPlanPtrOutput) Elem() ResourceModelWithAllowedPropertySetPlanOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetPlan) ResourceModelWithAllowedPropertySetPlan {
		if v != nil {
			return *v
		}
		var ret ResourceModelWithAllowedPropertySetPlan
		return ret
	}).(ResourceModelWithAllowedPropertySetPlanOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetPlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetPlan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetPlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetPlan) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetResponseIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ResourceModelWithAllowedPropertySetResponseIdentityOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetResponseIdentity)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityOutput) ToResourceModelWithAllowedPropertySetResponseIdentityOutput() ResourceModelWithAllowedPropertySetResponseIdentityOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityOutput) ToResourceModelWithAllowedPropertySetResponseIdentityOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetResponseIdentityOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetResponseIdentity)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput) ToResourceModelWithAllowedPropertySetResponseIdentityPtrOutput() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput) ToResourceModelWithAllowedPropertySetResponseIdentityPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput) Elem() ResourceModelWithAllowedPropertySetResponseIdentityOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseIdentity) ResourceModelWithAllowedPropertySetResponseIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceModelWithAllowedPropertySetResponseIdentity
		return ret
	}).(ResourceModelWithAllowedPropertySetResponseIdentityOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetResponsePlan struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}

type ResourceModelWithAllowedPropertySetResponsePlanOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetResponsePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetResponsePlan)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetResponsePlanOutput) ToResourceModelWithAllowedPropertySetResponsePlanOutput() ResourceModelWithAllowedPropertySetResponsePlanOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponsePlanOutput) ToResourceModelWithAllowedPropertySetResponsePlanOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetResponsePlanOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponsePlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponsePlan) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponsePlan) string { return v.Product }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponsePlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponsePlan) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponsePlan) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetResponsePlanPtrOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetResponsePlan)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) ToResourceModelWithAllowedPropertySetResponsePlanPtrOutput() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) ToResourceModelWithAllowedPropertySetResponsePlanPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) Elem() ResourceModelWithAllowedPropertySetResponsePlanOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponsePlan) ResourceModelWithAllowedPropertySetResponsePlan {
		if v != nil {
			return *v
		}
		var ret ResourceModelWithAllowedPropertySetResponsePlan
		return ret
	}).(ResourceModelWithAllowedPropertySetResponsePlanOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponsePlan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponsePlan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponsePlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponsePlan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponsePlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponsePlan) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetResponseSku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}

type ResourceModelWithAllowedPropertySetResponseSkuOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetResponseSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetResponseSku)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetResponseSkuOutput) ToResourceModelWithAllowedPropertySetResponseSkuOutput() ResourceModelWithAllowedPropertySetResponseSkuOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponseSkuOutput) ToResourceModelWithAllowedPropertySetResponseSkuOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetResponseSkuOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponseSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponseSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponseSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponseSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponseSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetResponseSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetResponseSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetResponseSku)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) ToResourceModelWithAllowedPropertySetResponseSkuPtrOutput() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) ToResourceModelWithAllowedPropertySetResponseSkuPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) Elem() ResourceModelWithAllowedPropertySetResponseSkuOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseSku) ResourceModelWithAllowedPropertySetResponseSku {
		if v != nil {
			return *v
		}
		var ret ResourceModelWithAllowedPropertySetResponseSku
		return ret
	}).(ResourceModelWithAllowedPropertySetResponseSkuOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetResponseSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetResponseSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ResourceModelWithAllowedPropertySetSku struct {
	Capacity *int     `pulumi:"capacity"`
	Family   *string  `pulumi:"family"`
	Name     string   `pulumi:"name"`
	Size     *string  `pulumi:"size"`
	Tier     *SkuTier `pulumi:"tier"`
}





type ResourceModelWithAllowedPropertySetSkuInput interface {
	pulumi.Input

	ToResourceModelWithAllowedPropertySetSkuOutput() ResourceModelWithAllowedPropertySetSkuOutput
	ToResourceModelWithAllowedPropertySetSkuOutputWithContext(context.Context) ResourceModelWithAllowedPropertySetSkuOutput
}

type ResourceModelWithAllowedPropertySetSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     SkuTierPtrInput       `pulumi:"tier"`
}

func (ResourceModelWithAllowedPropertySetSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetSku)(nil)).Elem()
}

func (i ResourceModelWithAllowedPropertySetSkuArgs) ToResourceModelWithAllowedPropertySetSkuOutput() ResourceModelWithAllowedPropertySetSkuOutput {
	return i.ToResourceModelWithAllowedPropertySetSkuOutputWithContext(context.Background())
}

func (i ResourceModelWithAllowedPropertySetSkuArgs) ToResourceModelWithAllowedPropertySetSkuOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetSkuOutput)
}

func (i ResourceModelWithAllowedPropertySetSkuArgs) ToResourceModelWithAllowedPropertySetSkuPtrOutput() ResourceModelWithAllowedPropertySetSkuPtrOutput {
	return i.ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(context.Background())
}

func (i ResourceModelWithAllowedPropertySetSkuArgs) ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetSkuOutput).ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(ctx)
}









type ResourceModelWithAllowedPropertySetSkuPtrInput interface {
	pulumi.Input

	ToResourceModelWithAllowedPropertySetSkuPtrOutput() ResourceModelWithAllowedPropertySetSkuPtrOutput
	ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(context.Context) ResourceModelWithAllowedPropertySetSkuPtrOutput
}

type resourceModelWithAllowedPropertySetSkuPtrType ResourceModelWithAllowedPropertySetSkuArgs

func ResourceModelWithAllowedPropertySetSkuPtr(v *ResourceModelWithAllowedPropertySetSkuArgs) ResourceModelWithAllowedPropertySetSkuPtrInput {
	return (*resourceModelWithAllowedPropertySetSkuPtrType)(v)
}

func (*resourceModelWithAllowedPropertySetSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetSku)(nil)).Elem()
}

func (i *resourceModelWithAllowedPropertySetSkuPtrType) ToResourceModelWithAllowedPropertySetSkuPtrOutput() ResourceModelWithAllowedPropertySetSkuPtrOutput {
	return i.ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(context.Background())
}

func (i *resourceModelWithAllowedPropertySetSkuPtrType) ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceModelWithAllowedPropertySetSkuPtrOutput)
}

type ResourceModelWithAllowedPropertySetSkuOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceModelWithAllowedPropertySetSku)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) ToResourceModelWithAllowedPropertySetSkuOutput() ResourceModelWithAllowedPropertySetSkuOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) ToResourceModelWithAllowedPropertySetSkuOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetSkuOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) ToResourceModelWithAllowedPropertySetSkuPtrOutput() ResourceModelWithAllowedPropertySetSkuPtrOutput {
	return o.ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(context.Background())
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceModelWithAllowedPropertySetSku) *ResourceModelWithAllowedPropertySetSku {
		return &v
	}).(ResourceModelWithAllowedPropertySetSkuPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetSku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v ResourceModelWithAllowedPropertySetSku) *SkuTier { return v.Tier }).(SkuTierPtrOutput)
}

type ResourceModelWithAllowedPropertySetSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceModelWithAllowedPropertySetSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceModelWithAllowedPropertySetSku)(nil)).Elem()
}

func (o ResourceModelWithAllowedPropertySetSkuPtrOutput) ToResourceModelWithAllowedPropertySetSkuPtrOutput() ResourceModelWithAllowedPropertySetSkuPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetSkuPtrOutput) ToResourceModelWithAllowedPropertySetSkuPtrOutputWithContext(ctx context.Context) ResourceModelWithAllowedPropertySetSkuPtrOutput {
	return o
}

func (o ResourceModelWithAllowedPropertySetSkuPtrOutput) Elem() ResourceModelWithAllowedPropertySetSkuOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetSku) ResourceModelWithAllowedPropertySetSku {
		if v != nil {
			return *v
		}
		var ret ResourceModelWithAllowedPropertySetSku
		return ret
	}).(ResourceModelWithAllowedPropertySetSkuOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetSku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ResourceModelWithAllowedPropertySetSkuPtrOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v *ResourceModelWithAllowedPropertySetSku) *SkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SkuTierPtrOutput)
}

type ScalingHostPoolReference struct {
	HostPoolArmPath    *string `pulumi:"hostPoolArmPath"`
	ScalingPlanEnabled *bool   `pulumi:"scalingPlanEnabled"`
}





type ScalingHostPoolReferenceInput interface {
	pulumi.Input

	ToScalingHostPoolReferenceOutput() ScalingHostPoolReferenceOutput
	ToScalingHostPoolReferenceOutputWithContext(context.Context) ScalingHostPoolReferenceOutput
}

type ScalingHostPoolReferenceArgs struct {
	HostPoolArmPath    pulumi.StringPtrInput `pulumi:"hostPoolArmPath"`
	ScalingPlanEnabled pulumi.BoolPtrInput   `pulumi:"scalingPlanEnabled"`
}

func (ScalingHostPoolReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingHostPoolReference)(nil)).Elem()
}

func (i ScalingHostPoolReferenceArgs) ToScalingHostPoolReferenceOutput() ScalingHostPoolReferenceOutput {
	return i.ToScalingHostPoolReferenceOutputWithContext(context.Background())
}

func (i ScalingHostPoolReferenceArgs) ToScalingHostPoolReferenceOutputWithContext(ctx context.Context) ScalingHostPoolReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingHostPoolReferenceOutput)
}





type ScalingHostPoolReferenceArrayInput interface {
	pulumi.Input

	ToScalingHostPoolReferenceArrayOutput() ScalingHostPoolReferenceArrayOutput
	ToScalingHostPoolReferenceArrayOutputWithContext(context.Context) ScalingHostPoolReferenceArrayOutput
}

type ScalingHostPoolReferenceArray []ScalingHostPoolReferenceInput

func (ScalingHostPoolReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingHostPoolReference)(nil)).Elem()
}

func (i ScalingHostPoolReferenceArray) ToScalingHostPoolReferenceArrayOutput() ScalingHostPoolReferenceArrayOutput {
	return i.ToScalingHostPoolReferenceArrayOutputWithContext(context.Background())
}

func (i ScalingHostPoolReferenceArray) ToScalingHostPoolReferenceArrayOutputWithContext(ctx context.Context) ScalingHostPoolReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingHostPoolReferenceArrayOutput)
}

type ScalingHostPoolReferenceOutput struct{ *pulumi.OutputState }

func (ScalingHostPoolReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingHostPoolReference)(nil)).Elem()
}

func (o ScalingHostPoolReferenceOutput) ToScalingHostPoolReferenceOutput() ScalingHostPoolReferenceOutput {
	return o
}

func (o ScalingHostPoolReferenceOutput) ToScalingHostPoolReferenceOutputWithContext(ctx context.Context) ScalingHostPoolReferenceOutput {
	return o
}

func (o ScalingHostPoolReferenceOutput) HostPoolArmPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingHostPoolReference) *string { return v.HostPoolArmPath }).(pulumi.StringPtrOutput)
}

func (o ScalingHostPoolReferenceOutput) ScalingPlanEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalingHostPoolReference) *bool { return v.ScalingPlanEnabled }).(pulumi.BoolPtrOutput)
}

type ScalingHostPoolReferenceArrayOutput struct{ *pulumi.OutputState }

func (ScalingHostPoolReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingHostPoolReference)(nil)).Elem()
}

func (o ScalingHostPoolReferenceArrayOutput) ToScalingHostPoolReferenceArrayOutput() ScalingHostPoolReferenceArrayOutput {
	return o
}

func (o ScalingHostPoolReferenceArrayOutput) ToScalingHostPoolReferenceArrayOutputWithContext(ctx context.Context) ScalingHostPoolReferenceArrayOutput {
	return o
}

func (o ScalingHostPoolReferenceArrayOutput) Index(i pulumi.IntInput) ScalingHostPoolReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingHostPoolReference {
		return vs[0].([]ScalingHostPoolReference)[vs[1].(int)]
	}).(ScalingHostPoolReferenceOutput)
}

type ScalingHostPoolReferenceResponse struct {
	HostPoolArmPath    *string `pulumi:"hostPoolArmPath"`
	ScalingPlanEnabled *bool   `pulumi:"scalingPlanEnabled"`
}

type ScalingHostPoolReferenceResponseOutput struct{ *pulumi.OutputState }

func (ScalingHostPoolReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingHostPoolReferenceResponse)(nil)).Elem()
}

func (o ScalingHostPoolReferenceResponseOutput) ToScalingHostPoolReferenceResponseOutput() ScalingHostPoolReferenceResponseOutput {
	return o
}

func (o ScalingHostPoolReferenceResponseOutput) ToScalingHostPoolReferenceResponseOutputWithContext(ctx context.Context) ScalingHostPoolReferenceResponseOutput {
	return o
}

func (o ScalingHostPoolReferenceResponseOutput) HostPoolArmPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingHostPoolReferenceResponse) *string { return v.HostPoolArmPath }).(pulumi.StringPtrOutput)
}

func (o ScalingHostPoolReferenceResponseOutput) ScalingPlanEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalingHostPoolReferenceResponse) *bool { return v.ScalingPlanEnabled }).(pulumi.BoolPtrOutput)
}

type ScalingHostPoolReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ScalingHostPoolReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingHostPoolReferenceResponse)(nil)).Elem()
}

func (o ScalingHostPoolReferenceResponseArrayOutput) ToScalingHostPoolReferenceResponseArrayOutput() ScalingHostPoolReferenceResponseArrayOutput {
	return o
}

func (o ScalingHostPoolReferenceResponseArrayOutput) ToScalingHostPoolReferenceResponseArrayOutputWithContext(ctx context.Context) ScalingHostPoolReferenceResponseArrayOutput {
	return o
}

func (o ScalingHostPoolReferenceResponseArrayOutput) Index(i pulumi.IntInput) ScalingHostPoolReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingHostPoolReferenceResponse {
		return vs[0].([]ScalingHostPoolReferenceResponse)[vs[1].(int)]
	}).(ScalingHostPoolReferenceResponseOutput)
}

type ScalingSchedule struct {
	DaysOfWeek                     []string `pulumi:"daysOfWeek"`
	Name                           *string  `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  *string  `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               *Time    `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     *string  `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  *Time    `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   *int     `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       *bool    `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm *string  `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        *int     `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    *string  `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              *Time    `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          *string  `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        *int     `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     *int     `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   *string  `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          *int     `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                *Time    `pulumi:"rampUpStartTime"`
}





type ScalingScheduleInput interface {
	pulumi.Input

	ToScalingScheduleOutput() ScalingScheduleOutput
	ToScalingScheduleOutputWithContext(context.Context) ScalingScheduleOutput
}

type ScalingScheduleArgs struct {
	DaysOfWeek                     pulumi.StringArrayInput `pulumi:"daysOfWeek"`
	Name                           pulumi.StringPtrInput   `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  pulumi.StringPtrInput   `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               TimePtrInput            `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     pulumi.StringPtrInput   `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  TimePtrInput            `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   pulumi.IntPtrInput      `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       pulumi.BoolPtrInput     `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm pulumi.StringPtrInput   `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        pulumi.IntPtrInput      `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    pulumi.StringPtrInput   `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              TimePtrInput            `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          pulumi.StringPtrInput   `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        pulumi.IntPtrInput      `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     pulumi.IntPtrInput      `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   pulumi.StringPtrInput   `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          pulumi.IntPtrInput      `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                TimePtrInput            `pulumi:"rampUpStartTime"`
}

func (ScalingScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingSchedule)(nil)).Elem()
}

func (i ScalingScheduleArgs) ToScalingScheduleOutput() ScalingScheduleOutput {
	return i.ToScalingScheduleOutputWithContext(context.Background())
}

func (i ScalingScheduleArgs) ToScalingScheduleOutputWithContext(ctx context.Context) ScalingScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingScheduleOutput)
}





type ScalingScheduleArrayInput interface {
	pulumi.Input

	ToScalingScheduleArrayOutput() ScalingScheduleArrayOutput
	ToScalingScheduleArrayOutputWithContext(context.Context) ScalingScheduleArrayOutput
}

type ScalingScheduleArray []ScalingScheduleInput

func (ScalingScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingSchedule)(nil)).Elem()
}

func (i ScalingScheduleArray) ToScalingScheduleArrayOutput() ScalingScheduleArrayOutput {
	return i.ToScalingScheduleArrayOutputWithContext(context.Background())
}

func (i ScalingScheduleArray) ToScalingScheduleArrayOutputWithContext(ctx context.Context) ScalingScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingScheduleArrayOutput)
}

type ScalingScheduleOutput struct{ *pulumi.OutputState }

func (ScalingScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingSchedule)(nil)).Elem()
}

func (o ScalingScheduleOutput) ToScalingScheduleOutput() ScalingScheduleOutput {
	return o
}

func (o ScalingScheduleOutput) ToScalingScheduleOutputWithContext(ctx context.Context) ScalingScheduleOutput {
	return o
}

func (o ScalingScheduleOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalingSchedule) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o ScalingScheduleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) OffPeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.OffPeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) OffPeakStartTime() TimePtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *Time { return v.OffPeakStartTime }).(TimePtrOutput)
}

func (o ScalingScheduleOutput) PeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.PeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) PeakStartTime() TimePtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *Time { return v.PeakStartTime }).(TimePtrOutput)
}

func (o ScalingScheduleOutput) RampDownCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampDownCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampDownForceLogoffUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *bool { return v.RampDownForceLogoffUsers }).(pulumi.BoolPtrOutput)
}

func (o ScalingScheduleOutput) RampDownLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampDownLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampDownMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampDownMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampDownNotificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampDownNotificationMessage }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampDownStartTime() TimePtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *Time { return v.RampDownStartTime }).(TimePtrOutput)
}

func (o ScalingScheduleOutput) RampDownStopHostsWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampDownStopHostsWhen }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampDownWaitTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampDownWaitTimeMinutes }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampUpCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampUpCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampUpLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *string { return v.RampUpLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleOutput) RampUpMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *int { return v.RampUpMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleOutput) RampUpStartTime() TimePtrOutput {
	return o.ApplyT(func(v ScalingSchedule) *Time { return v.RampUpStartTime }).(TimePtrOutput)
}

type ScalingScheduleArrayOutput struct{ *pulumi.OutputState }

func (ScalingScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingSchedule)(nil)).Elem()
}

func (o ScalingScheduleArrayOutput) ToScalingScheduleArrayOutput() ScalingScheduleArrayOutput {
	return o
}

func (o ScalingScheduleArrayOutput) ToScalingScheduleArrayOutputWithContext(ctx context.Context) ScalingScheduleArrayOutput {
	return o
}

func (o ScalingScheduleArrayOutput) Index(i pulumi.IntInput) ScalingScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingSchedule {
		return vs[0].([]ScalingSchedule)[vs[1].(int)]
	}).(ScalingScheduleOutput)
}

type ScalingScheduleResponse struct {
	DaysOfWeek                     []string      `pulumi:"daysOfWeek"`
	Name                           *string       `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  *string       `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               *TimeResponse `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     *string       `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  *TimeResponse `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   *int          `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       *bool         `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm *string       `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        *int          `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    *string       `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              *TimeResponse `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          *string       `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        *int          `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     *int          `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   *string       `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          *int          `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                *TimeResponse `pulumi:"rampUpStartTime"`
}

type ScalingScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScalingScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingScheduleResponse)(nil)).Elem()
}

func (o ScalingScheduleResponseOutput) ToScalingScheduleResponseOutput() ScalingScheduleResponseOutput {
	return o
}

func (o ScalingScheduleResponseOutput) ToScalingScheduleResponseOutputWithContext(ctx context.Context) ScalingScheduleResponseOutput {
	return o
}

func (o ScalingScheduleResponseOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o ScalingScheduleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) OffPeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.OffPeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) OffPeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *TimeResponse { return v.OffPeakStartTime }).(TimeResponsePtrOutput)
}

func (o ScalingScheduleResponseOutput) PeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.PeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) PeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *TimeResponse { return v.PeakStartTime }).(TimeResponsePtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampDownCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownForceLogoffUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *bool { return v.RampDownForceLogoffUsers }).(pulumi.BoolPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampDownLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampDownMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownNotificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampDownNotificationMessage }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *TimeResponse { return v.RampDownStartTime }).(TimeResponsePtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownStopHostsWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampDownStopHostsWhen }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampDownWaitTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampDownWaitTimeMinutes }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampUpCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampUpCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampUpLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *string { return v.RampUpLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampUpMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *int { return v.RampUpMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingScheduleResponseOutput) RampUpStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v ScalingScheduleResponse) *TimeResponse { return v.RampUpStartTime }).(TimeResponsePtrOutput)
}

type ScalingScheduleResponseArrayOutput struct{ *pulumi.OutputState }

func (ScalingScheduleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingScheduleResponse)(nil)).Elem()
}

func (o ScalingScheduleResponseArrayOutput) ToScalingScheduleResponseArrayOutput() ScalingScheduleResponseArrayOutput {
	return o
}

func (o ScalingScheduleResponseArrayOutput) ToScalingScheduleResponseArrayOutputWithContext(ctx context.Context) ScalingScheduleResponseArrayOutput {
	return o
}

func (o ScalingScheduleResponseArrayOutput) Index(i pulumi.IntInput) ScalingScheduleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingScheduleResponse {
		return vs[0].([]ScalingScheduleResponse)[vs[1].(int)]
	}).(ScalingScheduleResponseOutput)
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

type Time struct {
	Hour   int `pulumi:"hour"`
	Minute int `pulumi:"minute"`
}





type TimeInput interface {
	pulumi.Input

	ToTimeOutput() TimeOutput
	ToTimeOutputWithContext(context.Context) TimeOutput
}

type TimeArgs struct {
	Hour   pulumi.IntInput `pulumi:"hour"`
	Minute pulumi.IntInput `pulumi:"minute"`
}

func (TimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Time)(nil)).Elem()
}

func (i TimeArgs) ToTimeOutput() TimeOutput {
	return i.ToTimeOutputWithContext(context.Background())
}

func (i TimeArgs) ToTimeOutputWithContext(ctx context.Context) TimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeOutput)
}

func (i TimeArgs) ToTimePtrOutput() TimePtrOutput {
	return i.ToTimePtrOutputWithContext(context.Background())
}

func (i TimeArgs) ToTimePtrOutputWithContext(ctx context.Context) TimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeOutput).ToTimePtrOutputWithContext(ctx)
}









type TimePtrInput interface {
	pulumi.Input

	ToTimePtrOutput() TimePtrOutput
	ToTimePtrOutputWithContext(context.Context) TimePtrOutput
}

type timePtrType TimeArgs

func TimePtr(v *TimeArgs) TimePtrInput {
	return (*timePtrType)(v)
}

func (*timePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Time)(nil)).Elem()
}

func (i *timePtrType) ToTimePtrOutput() TimePtrOutput {
	return i.ToTimePtrOutputWithContext(context.Background())
}

func (i *timePtrType) ToTimePtrOutputWithContext(ctx context.Context) TimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimePtrOutput)
}

type TimeOutput struct{ *pulumi.OutputState }

func (TimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Time)(nil)).Elem()
}

func (o TimeOutput) ToTimeOutput() TimeOutput {
	return o
}

func (o TimeOutput) ToTimeOutputWithContext(ctx context.Context) TimeOutput {
	return o
}

func (o TimeOutput) ToTimePtrOutput() TimePtrOutput {
	return o.ToTimePtrOutputWithContext(context.Background())
}

func (o TimeOutput) ToTimePtrOutputWithContext(ctx context.Context) TimePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Time) *Time {
		return &v
	}).(TimePtrOutput)
}

func (o TimeOutput) Hour() pulumi.IntOutput {
	return o.ApplyT(func(v Time) int { return v.Hour }).(pulumi.IntOutput)
}

func (o TimeOutput) Minute() pulumi.IntOutput {
	return o.ApplyT(func(v Time) int { return v.Minute }).(pulumi.IntOutput)
}

type TimePtrOutput struct{ *pulumi.OutputState }

func (TimePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Time)(nil)).Elem()
}

func (o TimePtrOutput) ToTimePtrOutput() TimePtrOutput {
	return o
}

func (o TimePtrOutput) ToTimePtrOutputWithContext(ctx context.Context) TimePtrOutput {
	return o
}

func (o TimePtrOutput) Elem() TimeOutput {
	return o.ApplyT(func(v *Time) Time {
		if v != nil {
			return *v
		}
		var ret Time
		return ret
	}).(TimeOutput)
}

func (o TimePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Time) *int {
		if v == nil {
			return nil
		}
		return &v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o TimePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Time) *int {
		if v == nil {
			return nil
		}
		return &v.Minute
	}).(pulumi.IntPtrOutput)
}

type TimeResponse struct {
	Hour   int `pulumi:"hour"`
	Minute int `pulumi:"minute"`
}

type TimeResponseOutput struct{ *pulumi.OutputState }

func (TimeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeResponse)(nil)).Elem()
}

func (o TimeResponseOutput) ToTimeResponseOutput() TimeResponseOutput {
	return o
}

func (o TimeResponseOutput) ToTimeResponseOutputWithContext(ctx context.Context) TimeResponseOutput {
	return o
}

func (o TimeResponseOutput) Hour() pulumi.IntOutput {
	return o.ApplyT(func(v TimeResponse) int { return v.Hour }).(pulumi.IntOutput)
}

func (o TimeResponseOutput) Minute() pulumi.IntOutput {
	return o.ApplyT(func(v TimeResponse) int { return v.Minute }).(pulumi.IntOutput)
}

type TimeResponsePtrOutput struct{ *pulumi.OutputState }

func (TimeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeResponse)(nil)).Elem()
}

func (o TimeResponsePtrOutput) ToTimeResponsePtrOutput() TimeResponsePtrOutput {
	return o
}

func (o TimeResponsePtrOutput) ToTimeResponsePtrOutputWithContext(ctx context.Context) TimeResponsePtrOutput {
	return o
}

func (o TimeResponsePtrOutput) Elem() TimeResponseOutput {
	return o.ApplyT(func(v *TimeResponse) TimeResponse {
		if v != nil {
			return *v
		}
		var ret TimeResponse
		return ret
	}).(TimeResponseOutput)
}

func (o TimeResponsePtrOutput) Hour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Hour
	}).(pulumi.IntPtrOutput)
}

func (o TimeResponsePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Minute
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentUpdatePropertiesOutput{})
	pulumi.RegisterOutputType(AgentUpdatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AgentUpdatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AgentUpdatePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPropertiesOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPropertiesArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(MigrationRequestPropertiesOutput{})
	pulumi.RegisterOutputType(MigrationRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MigrationRequestPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MigrationRequestPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsResponseOutput{})
	pulumi.RegisterOutputType(MsixPackageApplicationsResponseArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesArrayOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesResponseOutput{})
	pulumi.RegisterOutputType(MsixPackageDependenciesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(RegistrationInfoOutput{})
	pulumi.RegisterOutputType(RegistrationInfoPtrOutput{})
	pulumi.RegisterOutputType(RegistrationInfoResponseOutput{})
	pulumi.RegisterOutputType(RegistrationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetIdentityOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetPlanOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetPlanPtrOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetResponseIdentityOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetResponsePlanOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetResponseSkuOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetSkuOutput{})
	pulumi.RegisterOutputType(ResourceModelWithAllowedPropertySetSkuPtrOutput{})
	pulumi.RegisterOutputType(ScalingHostPoolReferenceOutput{})
	pulumi.RegisterOutputType(ScalingHostPoolReferenceArrayOutput{})
	pulumi.RegisterOutputType(ScalingHostPoolReferenceResponseOutput{})
	pulumi.RegisterOutputType(ScalingHostPoolReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(ScalingScheduleOutput{})
	pulumi.RegisterOutputType(ScalingScheduleArrayOutput{})
	pulumi.RegisterOutputType(ScalingScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScalingScheduleResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TimeOutput{})
	pulumi.RegisterOutputType(TimePtrOutput{})
	pulumi.RegisterOutputType(TimeResponseOutput{})
	pulumi.RegisterOutputType(TimeResponsePtrOutput{})
}
