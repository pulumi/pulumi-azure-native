


package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConsoleCreateProperties struct {
	OsType            string  `pulumi:"osType"`
	ProvisioningState *string `pulumi:"provisioningState"`
	Uri               *string `pulumi:"uri"`
}





type ConsoleCreatePropertiesInput interface {
	pulumi.Input

	ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput
	ToConsoleCreatePropertiesOutputWithContext(context.Context) ConsoleCreatePropertiesOutput
}

type ConsoleCreatePropertiesArgs struct {
	OsType            pulumi.StringInput    `pulumi:"osType"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	Uri               pulumi.StringPtrInput `pulumi:"uri"`
}

func (ConsoleCreatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleCreateProperties)(nil)).Elem()
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput {
	return i.ToConsoleCreatePropertiesOutputWithContext(context.Background())
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesOutputWithContext(ctx context.Context) ConsoleCreatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleCreatePropertiesOutput)
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput {
	return i.ToConsoleCreatePropertiesPtrOutputWithContext(context.Background())
}

func (i ConsoleCreatePropertiesArgs) ToConsoleCreatePropertiesPtrOutputWithContext(ctx context.Context) ConsoleCreatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleCreatePropertiesOutput).ToConsoleCreatePropertiesPtrOutputWithContext(ctx)
}









type ConsoleCreatePropertiesPtrInput interface {
	pulumi.Input

	ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput
	ToConsoleCreatePropertiesPtrOutputWithContext(context.Context) ConsoleCreatePropertiesPtrOutput
}

type consoleCreatePropertiesPtrType ConsoleCreatePropertiesArgs

func ConsoleCreatePropertiesPtr(v *ConsoleCreatePropertiesArgs) ConsoleCreatePropertiesPtrInput {
	return (*consoleCreatePropertiesPtrType)(v)
}

func (*consoleCreatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsoleCreateProperties)(nil)).Elem()
}

func (i *consoleCreatePropertiesPtrType) ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput {
	return i.ToConsoleCreatePropertiesPtrOutputWithContext(context.Background())
}

func (i *consoleCreatePropertiesPtrType) ToConsoleCreatePropertiesPtrOutputWithContext(ctx context.Context) ConsoleCreatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleCreatePropertiesPtrOutput)
}

type ConsoleCreatePropertiesOutput struct{ *pulumi.OutputState }

func (ConsoleCreatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsoleCreateProperties)(nil)).Elem()
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesOutput() ConsoleCreatePropertiesOutput {
	return o
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesOutputWithContext(ctx context.Context) ConsoleCreatePropertiesOutput {
	return o
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput {
	return o.ToConsoleCreatePropertiesPtrOutputWithContext(context.Background())
}

func (o ConsoleCreatePropertiesOutput) ToConsoleCreatePropertiesPtrOutputWithContext(ctx context.Context) ConsoleCreatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsoleCreateProperties) *ConsoleCreateProperties {
		return &v
	}).(ConsoleCreatePropertiesPtrOutput)
}

func (o ConsoleCreatePropertiesOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ConsoleCreatePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ConsoleCreatePropertiesOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ConsoleCreatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConsoleCreatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsoleCreateProperties)(nil)).Elem()
}

func (o ConsoleCreatePropertiesPtrOutput) ToConsoleCreatePropertiesPtrOutput() ConsoleCreatePropertiesPtrOutput {
	return o
}

func (o ConsoleCreatePropertiesPtrOutput) ToConsoleCreatePropertiesPtrOutputWithContext(ctx context.Context) ConsoleCreatePropertiesPtrOutput {
	return o
}

func (o ConsoleCreatePropertiesPtrOutput) Elem() ConsoleCreatePropertiesOutput {
	return o.ApplyT(func(v *ConsoleCreateProperties) ConsoleCreateProperties {
		if v != nil {
			return *v
		}
		var ret ConsoleCreateProperties
		return ret
	}).(ConsoleCreatePropertiesOutput)
}

func (o ConsoleCreatePropertiesPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsoleCreateProperties) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ConsoleCreatePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsoleCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ConsoleCreatePropertiesPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsoleCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type ConsolePropertiesResponse struct {
	OsType            string `pulumi:"osType"`
	ProvisioningState string `pulumi:"provisioningState"`
	Uri               string `pulumi:"uri"`
}





type ConsolePropertiesResponseInput interface {
	pulumi.Input

	ToConsolePropertiesResponseOutput() ConsolePropertiesResponseOutput
	ToConsolePropertiesResponseOutputWithContext(context.Context) ConsolePropertiesResponseOutput
}

type ConsolePropertiesResponseArgs struct {
	OsType            pulumi.StringInput `pulumi:"osType"`
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
	Uri               pulumi.StringInput `pulumi:"uri"`
}

func (ConsolePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsolePropertiesResponse)(nil)).Elem()
}

func (i ConsolePropertiesResponseArgs) ToConsolePropertiesResponseOutput() ConsolePropertiesResponseOutput {
	return i.ToConsolePropertiesResponseOutputWithContext(context.Background())
}

func (i ConsolePropertiesResponseArgs) ToConsolePropertiesResponseOutputWithContext(ctx context.Context) ConsolePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsolePropertiesResponseOutput)
}

func (i ConsolePropertiesResponseArgs) ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput {
	return i.ToConsolePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConsolePropertiesResponseArgs) ToConsolePropertiesResponsePtrOutputWithContext(ctx context.Context) ConsolePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsolePropertiesResponseOutput).ToConsolePropertiesResponsePtrOutputWithContext(ctx)
}









type ConsolePropertiesResponsePtrInput interface {
	pulumi.Input

	ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput
	ToConsolePropertiesResponsePtrOutputWithContext(context.Context) ConsolePropertiesResponsePtrOutput
}

type consolePropertiesResponsePtrType ConsolePropertiesResponseArgs

func ConsolePropertiesResponsePtr(v *ConsolePropertiesResponseArgs) ConsolePropertiesResponsePtrInput {
	return (*consolePropertiesResponsePtrType)(v)
}

func (*consolePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsolePropertiesResponse)(nil)).Elem()
}

func (i *consolePropertiesResponsePtrType) ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput {
	return i.ToConsolePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *consolePropertiesResponsePtrType) ToConsolePropertiesResponsePtrOutputWithContext(ctx context.Context) ConsolePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsolePropertiesResponsePtrOutput)
}

type ConsolePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConsolePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsolePropertiesResponse)(nil)).Elem()
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponseOutput() ConsolePropertiesResponseOutput {
	return o
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponseOutputWithContext(ctx context.Context) ConsolePropertiesResponseOutput {
	return o
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput {
	return o.ToConsolePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConsolePropertiesResponseOutput) ToConsolePropertiesResponsePtrOutputWithContext(ctx context.Context) ConsolePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsolePropertiesResponse) *ConsolePropertiesResponse {
		return &v
	}).(ConsolePropertiesResponsePtrOutput)
}

func (o ConsolePropertiesResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ConsolePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConsolePropertiesResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ConsolePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConsolePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsolePropertiesResponse)(nil)).Elem()
}

func (o ConsolePropertiesResponsePtrOutput) ToConsolePropertiesResponsePtrOutput() ConsolePropertiesResponsePtrOutput {
	return o
}

func (o ConsolePropertiesResponsePtrOutput) ToConsolePropertiesResponsePtrOutputWithContext(ctx context.Context) ConsolePropertiesResponsePtrOutput {
	return o
}

func (o ConsolePropertiesResponsePtrOutput) Elem() ConsolePropertiesResponseOutput {
	return o.ApplyT(func(v *ConsolePropertiesResponse) ConsolePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConsolePropertiesResponse
		return ret
	}).(ConsolePropertiesResponseOutput)
}

func (o ConsolePropertiesResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsolePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ConsolePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsolePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ConsolePropertiesResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsolePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type StorageProfile struct {
	DiskSizeInGB             *int    `pulumi:"diskSizeInGB"`
	FileShareName            *string `pulumi:"fileShareName"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	DiskSizeInGB             pulumi.IntPtrInput    `pulumi:"diskSizeInGB"`
	FileShareName            pulumi.StringPtrInput `pulumi:"fileShareName"`
	StorageAccountResourceId pulumi.StringPtrInput `pulumi:"storageAccountResourceId"`
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

func (o StorageProfileOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.DiskSizeInGB }).(pulumi.IntPtrOutput)
}

func (o StorageProfileOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.FileShareName }).(pulumi.StringPtrOutput)
}

func (o StorageProfileOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
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

func (o StorageProfilePtrOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeInGB
	}).(pulumi.IntPtrOutput)
}

func (o StorageProfilePtrOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.FileShareName
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfilePtrOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountResourceId
	}).(pulumi.StringPtrOutput)
}

type StorageProfileResponse struct {
	DiskSizeInGB             *int    `pulumi:"diskSizeInGB"`
	FileShareName            *string `pulumi:"fileShareName"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
}





type StorageProfileResponseInput interface {
	pulumi.Input

	ToStorageProfileResponseOutput() StorageProfileResponseOutput
	ToStorageProfileResponseOutputWithContext(context.Context) StorageProfileResponseOutput
}

type StorageProfileResponseArgs struct {
	DiskSizeInGB             pulumi.IntPtrInput    `pulumi:"diskSizeInGB"`
	FileShareName            pulumi.StringPtrInput `pulumi:"fileShareName"`
	StorageAccountResourceId pulumi.StringPtrInput `pulumi:"storageAccountResourceId"`
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

func (o StorageProfileResponseOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.DiskSizeInGB }).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponseOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.FileShareName }).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponseOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
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

func (o StorageProfileResponsePtrOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeInGB
	}).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponsePtrOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.FileShareName
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponsePtrOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountResourceId
	}).(pulumi.StringPtrOutput)
}

type TerminalSettings struct {
	FontSize  *string `pulumi:"fontSize"`
	FontStyle *string `pulumi:"fontStyle"`
}





type TerminalSettingsInput interface {
	pulumi.Input

	ToTerminalSettingsOutput() TerminalSettingsOutput
	ToTerminalSettingsOutputWithContext(context.Context) TerminalSettingsOutput
}

type TerminalSettingsArgs struct {
	FontSize  pulumi.StringPtrInput `pulumi:"fontSize"`
	FontStyle pulumi.StringPtrInput `pulumi:"fontStyle"`
}

func (TerminalSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettings)(nil)).Elem()
}

func (i TerminalSettingsArgs) ToTerminalSettingsOutput() TerminalSettingsOutput {
	return i.ToTerminalSettingsOutputWithContext(context.Background())
}

func (i TerminalSettingsArgs) ToTerminalSettingsOutputWithContext(ctx context.Context) TerminalSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsOutput)
}

func (i TerminalSettingsArgs) ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput {
	return i.ToTerminalSettingsPtrOutputWithContext(context.Background())
}

func (i TerminalSettingsArgs) ToTerminalSettingsPtrOutputWithContext(ctx context.Context) TerminalSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsOutput).ToTerminalSettingsPtrOutputWithContext(ctx)
}









type TerminalSettingsPtrInput interface {
	pulumi.Input

	ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput
	ToTerminalSettingsPtrOutputWithContext(context.Context) TerminalSettingsPtrOutput
}

type terminalSettingsPtrType TerminalSettingsArgs

func TerminalSettingsPtr(v *TerminalSettingsArgs) TerminalSettingsPtrInput {
	return (*terminalSettingsPtrType)(v)
}

func (*terminalSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminalSettings)(nil)).Elem()
}

func (i *terminalSettingsPtrType) ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput {
	return i.ToTerminalSettingsPtrOutputWithContext(context.Background())
}

func (i *terminalSettingsPtrType) ToTerminalSettingsPtrOutputWithContext(ctx context.Context) TerminalSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsPtrOutput)
}

type TerminalSettingsOutput struct{ *pulumi.OutputState }

func (TerminalSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettings)(nil)).Elem()
}

func (o TerminalSettingsOutput) ToTerminalSettingsOutput() TerminalSettingsOutput {
	return o
}

func (o TerminalSettingsOutput) ToTerminalSettingsOutputWithContext(ctx context.Context) TerminalSettingsOutput {
	return o
}

func (o TerminalSettingsOutput) ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput {
	return o.ToTerminalSettingsPtrOutputWithContext(context.Background())
}

func (o TerminalSettingsOutput) ToTerminalSettingsPtrOutputWithContext(ctx context.Context) TerminalSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TerminalSettings) *TerminalSettings {
		return &v
	}).(TerminalSettingsPtrOutput)
}

func (o TerminalSettingsOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettings) *string { return v.FontSize }).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettings) *string { return v.FontStyle }).(pulumi.StringPtrOutput)
}

type TerminalSettingsPtrOutput struct{ *pulumi.OutputState }

func (TerminalSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminalSettings)(nil)).Elem()
}

func (o TerminalSettingsPtrOutput) ToTerminalSettingsPtrOutput() TerminalSettingsPtrOutput {
	return o
}

func (o TerminalSettingsPtrOutput) ToTerminalSettingsPtrOutputWithContext(ctx context.Context) TerminalSettingsPtrOutput {
	return o
}

func (o TerminalSettingsPtrOutput) Elem() TerminalSettingsOutput {
	return o.ApplyT(func(v *TerminalSettings) TerminalSettings {
		if v != nil {
			return *v
		}
		var ret TerminalSettings
		return ret
	}).(TerminalSettingsOutput)
}

func (o TerminalSettingsPtrOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminalSettings) *string {
		if v == nil {
			return nil
		}
		return v.FontSize
	}).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsPtrOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminalSettings) *string {
		if v == nil {
			return nil
		}
		return v.FontStyle
	}).(pulumi.StringPtrOutput)
}

type TerminalSettingsResponse struct {
	FontSize  *string `pulumi:"fontSize"`
	FontStyle *string `pulumi:"fontStyle"`
}





type TerminalSettingsResponseInput interface {
	pulumi.Input

	ToTerminalSettingsResponseOutput() TerminalSettingsResponseOutput
	ToTerminalSettingsResponseOutputWithContext(context.Context) TerminalSettingsResponseOutput
}

type TerminalSettingsResponseArgs struct {
	FontSize  pulumi.StringPtrInput `pulumi:"fontSize"`
	FontStyle pulumi.StringPtrInput `pulumi:"fontStyle"`
}

func (TerminalSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettingsResponse)(nil)).Elem()
}

func (i TerminalSettingsResponseArgs) ToTerminalSettingsResponseOutput() TerminalSettingsResponseOutput {
	return i.ToTerminalSettingsResponseOutputWithContext(context.Background())
}

func (i TerminalSettingsResponseArgs) ToTerminalSettingsResponseOutputWithContext(ctx context.Context) TerminalSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsResponseOutput)
}

func (i TerminalSettingsResponseArgs) ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput {
	return i.ToTerminalSettingsResponsePtrOutputWithContext(context.Background())
}

func (i TerminalSettingsResponseArgs) ToTerminalSettingsResponsePtrOutputWithContext(ctx context.Context) TerminalSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsResponseOutput).ToTerminalSettingsResponsePtrOutputWithContext(ctx)
}









type TerminalSettingsResponsePtrInput interface {
	pulumi.Input

	ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput
	ToTerminalSettingsResponsePtrOutputWithContext(context.Context) TerminalSettingsResponsePtrOutput
}

type terminalSettingsResponsePtrType TerminalSettingsResponseArgs

func TerminalSettingsResponsePtr(v *TerminalSettingsResponseArgs) TerminalSettingsResponsePtrInput {
	return (*terminalSettingsResponsePtrType)(v)
}

func (*terminalSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminalSettingsResponse)(nil)).Elem()
}

func (i *terminalSettingsResponsePtrType) ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput {
	return i.ToTerminalSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *terminalSettingsResponsePtrType) ToTerminalSettingsResponsePtrOutputWithContext(ctx context.Context) TerminalSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminalSettingsResponsePtrOutput)
}

type TerminalSettingsResponseOutput struct{ *pulumi.OutputState }

func (TerminalSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminalSettingsResponse)(nil)).Elem()
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponseOutput() TerminalSettingsResponseOutput {
	return o
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponseOutputWithContext(ctx context.Context) TerminalSettingsResponseOutput {
	return o
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput {
	return o.ToTerminalSettingsResponsePtrOutputWithContext(context.Background())
}

func (o TerminalSettingsResponseOutput) ToTerminalSettingsResponsePtrOutputWithContext(ctx context.Context) TerminalSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TerminalSettingsResponse) *TerminalSettingsResponse {
		return &v
	}).(TerminalSettingsResponsePtrOutput)
}

func (o TerminalSettingsResponseOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettingsResponse) *string { return v.FontSize }).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsResponseOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettingsResponse) *string { return v.FontStyle }).(pulumi.StringPtrOutput)
}

type TerminalSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (TerminalSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminalSettingsResponse)(nil)).Elem()
}

func (o TerminalSettingsResponsePtrOutput) ToTerminalSettingsResponsePtrOutput() TerminalSettingsResponsePtrOutput {
	return o
}

func (o TerminalSettingsResponsePtrOutput) ToTerminalSettingsResponsePtrOutputWithContext(ctx context.Context) TerminalSettingsResponsePtrOutput {
	return o
}

func (o TerminalSettingsResponsePtrOutput) Elem() TerminalSettingsResponseOutput {
	return o.ApplyT(func(v *TerminalSettingsResponse) TerminalSettingsResponse {
		if v != nil {
			return *v
		}
		var ret TerminalSettingsResponse
		return ret
	}).(TerminalSettingsResponseOutput)
}

func (o TerminalSettingsResponsePtrOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminalSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FontSize
	}).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsResponsePtrOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminalSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FontStyle
	}).(pulumi.StringPtrOutput)
}

type UserProperties struct {
	PreferredLocation  string           `pulumi:"preferredLocation"`
	PreferredOsType    string           `pulumi:"preferredOsType"`
	PreferredShellType string           `pulumi:"preferredShellType"`
	StorageProfile     StorageProfile   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettings `pulumi:"terminalSettings"`
}





type UserPropertiesInput interface {
	pulumi.Input

	ToUserPropertiesOutput() UserPropertiesOutput
	ToUserPropertiesOutputWithContext(context.Context) UserPropertiesOutput
}

type UserPropertiesArgs struct {
	PreferredLocation  pulumi.StringInput    `pulumi:"preferredLocation"`
	PreferredOsType    pulumi.StringInput    `pulumi:"preferredOsType"`
	PreferredShellType pulumi.StringInput    `pulumi:"preferredShellType"`
	StorageProfile     StorageProfileInput   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettingsInput `pulumi:"terminalSettings"`
}

func (UserPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProperties)(nil)).Elem()
}

func (i UserPropertiesArgs) ToUserPropertiesOutput() UserPropertiesOutput {
	return i.ToUserPropertiesOutputWithContext(context.Background())
}

func (i UserPropertiesArgs) ToUserPropertiesOutputWithContext(ctx context.Context) UserPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesOutput)
}

func (i UserPropertiesArgs) ToUserPropertiesPtrOutput() UserPropertiesPtrOutput {
	return i.ToUserPropertiesPtrOutputWithContext(context.Background())
}

func (i UserPropertiesArgs) ToUserPropertiesPtrOutputWithContext(ctx context.Context) UserPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesOutput).ToUserPropertiesPtrOutputWithContext(ctx)
}









type UserPropertiesPtrInput interface {
	pulumi.Input

	ToUserPropertiesPtrOutput() UserPropertiesPtrOutput
	ToUserPropertiesPtrOutputWithContext(context.Context) UserPropertiesPtrOutput
}

type userPropertiesPtrType UserPropertiesArgs

func UserPropertiesPtr(v *UserPropertiesArgs) UserPropertiesPtrInput {
	return (*userPropertiesPtrType)(v)
}

func (*userPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserProperties)(nil)).Elem()
}

func (i *userPropertiesPtrType) ToUserPropertiesPtrOutput() UserPropertiesPtrOutput {
	return i.ToUserPropertiesPtrOutputWithContext(context.Background())
}

func (i *userPropertiesPtrType) ToUserPropertiesPtrOutputWithContext(ctx context.Context) UserPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesPtrOutput)
}

type UserPropertiesOutput struct{ *pulumi.OutputState }

func (UserPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProperties)(nil)).Elem()
}

func (o UserPropertiesOutput) ToUserPropertiesOutput() UserPropertiesOutput {
	return o
}

func (o UserPropertiesOutput) ToUserPropertiesOutputWithContext(ctx context.Context) UserPropertiesOutput {
	return o
}

func (o UserPropertiesOutput) ToUserPropertiesPtrOutput() UserPropertiesPtrOutput {
	return o.ToUserPropertiesPtrOutputWithContext(context.Background())
}

func (o UserPropertiesOutput) ToUserPropertiesPtrOutputWithContext(ctx context.Context) UserPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserProperties) *UserProperties {
		return &v
	}).(UserPropertiesPtrOutput)
}

func (o UserPropertiesOutput) PreferredLocation() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredLocation }).(pulumi.StringOutput)
}

func (o UserPropertiesOutput) PreferredOsType() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredOsType }).(pulumi.StringOutput)
}

func (o UserPropertiesOutput) PreferredShellType() pulumi.StringOutput {
	return o.ApplyT(func(v UserProperties) string { return v.PreferredShellType }).(pulumi.StringOutput)
}

func (o UserPropertiesOutput) StorageProfile() StorageProfileOutput {
	return o.ApplyT(func(v UserProperties) StorageProfile { return v.StorageProfile }).(StorageProfileOutput)
}

func (o UserPropertiesOutput) TerminalSettings() TerminalSettingsOutput {
	return o.ApplyT(func(v UserProperties) TerminalSettings { return v.TerminalSettings }).(TerminalSettingsOutput)
}

type UserPropertiesPtrOutput struct{ *pulumi.OutputState }

func (UserPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserProperties)(nil)).Elem()
}

func (o UserPropertiesPtrOutput) ToUserPropertiesPtrOutput() UserPropertiesPtrOutput {
	return o
}

func (o UserPropertiesPtrOutput) ToUserPropertiesPtrOutputWithContext(ctx context.Context) UserPropertiesPtrOutput {
	return o
}

func (o UserPropertiesPtrOutput) Elem() UserPropertiesOutput {
	return o.ApplyT(func(v *UserProperties) UserProperties {
		if v != nil {
			return *v
		}
		var ret UserProperties
		return ret
	}).(UserPropertiesOutput)
}

func (o UserPropertiesPtrOutput) PreferredLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredLocation
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesPtrOutput) PreferredOsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredOsType
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesPtrOutput) PreferredShellType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredShellType
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesPtrOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v *UserProperties) *StorageProfile {
		if v == nil {
			return nil
		}
		return &v.StorageProfile
	}).(StorageProfilePtrOutput)
}

func (o UserPropertiesPtrOutput) TerminalSettings() TerminalSettingsPtrOutput {
	return o.ApplyT(func(v *UserProperties) *TerminalSettings {
		if v == nil {
			return nil
		}
		return &v.TerminalSettings
	}).(TerminalSettingsPtrOutput)
}

type UserPropertiesResponse struct {
	PreferredLocation  string                   `pulumi:"preferredLocation"`
	PreferredOsType    string                   `pulumi:"preferredOsType"`
	PreferredShellType string                   `pulumi:"preferredShellType"`
	StorageProfile     StorageProfileResponse   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettingsResponse `pulumi:"terminalSettings"`
}





type UserPropertiesResponseInput interface {
	pulumi.Input

	ToUserPropertiesResponseOutput() UserPropertiesResponseOutput
	ToUserPropertiesResponseOutputWithContext(context.Context) UserPropertiesResponseOutput
}

type UserPropertiesResponseArgs struct {
	PreferredLocation  pulumi.StringInput            `pulumi:"preferredLocation"`
	PreferredOsType    pulumi.StringInput            `pulumi:"preferredOsType"`
	PreferredShellType pulumi.StringInput            `pulumi:"preferredShellType"`
	StorageProfile     StorageProfileResponseInput   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettingsResponseInput `pulumi:"terminalSettings"`
}

func (UserPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPropertiesResponse)(nil)).Elem()
}

func (i UserPropertiesResponseArgs) ToUserPropertiesResponseOutput() UserPropertiesResponseOutput {
	return i.ToUserPropertiesResponseOutputWithContext(context.Background())
}

func (i UserPropertiesResponseArgs) ToUserPropertiesResponseOutputWithContext(ctx context.Context) UserPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesResponseOutput)
}

func (i UserPropertiesResponseArgs) ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput {
	return i.ToUserPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i UserPropertiesResponseArgs) ToUserPropertiesResponsePtrOutputWithContext(ctx context.Context) UserPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesResponseOutput).ToUserPropertiesResponsePtrOutputWithContext(ctx)
}









type UserPropertiesResponsePtrInput interface {
	pulumi.Input

	ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput
	ToUserPropertiesResponsePtrOutputWithContext(context.Context) UserPropertiesResponsePtrOutput
}

type userPropertiesResponsePtrType UserPropertiesResponseArgs

func UserPropertiesResponsePtr(v *UserPropertiesResponseArgs) UserPropertiesResponsePtrInput {
	return (*userPropertiesResponsePtrType)(v)
}

func (*userPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPropertiesResponse)(nil)).Elem()
}

func (i *userPropertiesResponsePtrType) ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput {
	return i.ToUserPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *userPropertiesResponsePtrType) ToUserPropertiesResponsePtrOutputWithContext(ctx context.Context) UserPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPropertiesResponsePtrOutput)
}

type UserPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPropertiesResponse)(nil)).Elem()
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponseOutput() UserPropertiesResponseOutput {
	return o
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponseOutputWithContext(ctx context.Context) UserPropertiesResponseOutput {
	return o
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput {
	return o.ToUserPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o UserPropertiesResponseOutput) ToUserPropertiesResponsePtrOutputWithContext(ctx context.Context) UserPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserPropertiesResponse) *UserPropertiesResponse {
		return &v
	}).(UserPropertiesResponsePtrOutput)
}

func (o UserPropertiesResponseOutput) PreferredLocation() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredLocation }).(pulumi.StringOutput)
}

func (o UserPropertiesResponseOutput) PreferredOsType() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredOsType }).(pulumi.StringOutput)
}

func (o UserPropertiesResponseOutput) PreferredShellType() pulumi.StringOutput {
	return o.ApplyT(func(v UserPropertiesResponse) string { return v.PreferredShellType }).(pulumi.StringOutput)
}

func (o UserPropertiesResponseOutput) StorageProfile() StorageProfileResponseOutput {
	return o.ApplyT(func(v UserPropertiesResponse) StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponseOutput)
}

func (o UserPropertiesResponseOutput) TerminalSettings() TerminalSettingsResponseOutput {
	return o.ApplyT(func(v UserPropertiesResponse) TerminalSettingsResponse { return v.TerminalSettings }).(TerminalSettingsResponseOutput)
}

type UserPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (UserPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPropertiesResponse)(nil)).Elem()
}

func (o UserPropertiesResponsePtrOutput) ToUserPropertiesResponsePtrOutput() UserPropertiesResponsePtrOutput {
	return o
}

func (o UserPropertiesResponsePtrOutput) ToUserPropertiesResponsePtrOutputWithContext(ctx context.Context) UserPropertiesResponsePtrOutput {
	return o
}

func (o UserPropertiesResponsePtrOutput) Elem() UserPropertiesResponseOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) UserPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret UserPropertiesResponse
		return ret
	}).(UserPropertiesResponseOutput)
}

func (o UserPropertiesResponsePtrOutput) PreferredLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredLocation
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesResponsePtrOutput) PreferredOsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredOsType
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesResponsePtrOutput) PreferredShellType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreferredShellType
	}).(pulumi.StringPtrOutput)
}

func (o UserPropertiesResponsePtrOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *StorageProfileResponse {
		if v == nil {
			return nil
		}
		return &v.StorageProfile
	}).(StorageProfileResponsePtrOutput)
}

func (o UserPropertiesResponsePtrOutput) TerminalSettings() TerminalSettingsResponsePtrOutput {
	return o.ApplyT(func(v *UserPropertiesResponse) *TerminalSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.TerminalSettings
	}).(TerminalSettingsResponsePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsoleCreatePropertiesInput)(nil)).Elem(), ConsoleCreatePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsoleCreatePropertiesPtrInput)(nil)).Elem(), ConsoleCreatePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsolePropertiesResponseInput)(nil)).Elem(), ConsolePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsolePropertiesResponsePtrInput)(nil)).Elem(), ConsolePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileInput)(nil)).Elem(), StorageProfileArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfilePtrInput)(nil)).Elem(), StorageProfileArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileResponseInput)(nil)).Elem(), StorageProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageProfileResponsePtrInput)(nil)).Elem(), StorageProfileResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerminalSettingsInput)(nil)).Elem(), TerminalSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerminalSettingsPtrInput)(nil)).Elem(), TerminalSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerminalSettingsResponseInput)(nil)).Elem(), TerminalSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TerminalSettingsResponsePtrInput)(nil)).Elem(), TerminalSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertiesInput)(nil)).Elem(), UserPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertiesPtrInput)(nil)).Elem(), UserPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertiesResponseInput)(nil)).Elem(), UserPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPropertiesResponsePtrInput)(nil)).Elem(), UserPropertiesResponseArgs{})
	pulumi.RegisterOutputType(ConsoleCreatePropertiesOutput{})
	pulumi.RegisterOutputType(ConsoleCreatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConsolePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConsolePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(TerminalSettingsOutput{})
	pulumi.RegisterOutputType(TerminalSettingsPtrOutput{})
	pulumi.RegisterOutputType(TerminalSettingsResponseOutput{})
	pulumi.RegisterOutputType(TerminalSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserPropertiesOutput{})
	pulumi.RegisterOutputType(UserPropertiesPtrOutput{})
	pulumi.RegisterOutputType(UserPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserPropertiesResponsePtrOutput{})
}
