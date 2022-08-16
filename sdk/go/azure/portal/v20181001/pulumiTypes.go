


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

func (o ConsoleCreatePropertiesOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ConsoleCreatePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ConsoleCreatePropertiesOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsoleCreateProperties) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ConsolePropertiesResponse struct {
	OsType            string `pulumi:"osType"`
	ProvisioningState string `pulumi:"provisioningState"`
	Uri               string `pulumi:"uri"`
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

func (o ConsolePropertiesResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ConsolePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ConsolePropertiesResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ConsolePropertiesResponse) string { return v.Uri }).(pulumi.StringOutput)
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

func (o StorageProfileOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.DiskSizeInGB }).(pulumi.IntPtrOutput)
}

func (o StorageProfileOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.FileShareName }).(pulumi.StringPtrOutput)
}

func (o StorageProfileOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
}

type StorageProfileResponse struct {
	DiskSizeInGB             *int    `pulumi:"diskSizeInGB"`
	FileShareName            *string `pulumi:"fileShareName"`
	StorageAccountResourceId *string `pulumi:"storageAccountResourceId"`
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

func (o StorageProfileResponseOutput) DiskSizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.DiskSizeInGB }).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponseOutput) FileShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.FileShareName }).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponseOutput) StorageAccountResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.StorageAccountResourceId }).(pulumi.StringPtrOutput)
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

func (o TerminalSettingsOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettings) *string { return v.FontSize }).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettings) *string { return v.FontStyle }).(pulumi.StringPtrOutput)
}

type TerminalSettingsResponse struct {
	FontSize  *string `pulumi:"fontSize"`
	FontStyle *string `pulumi:"fontStyle"`
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

func (o TerminalSettingsResponseOutput) FontSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettingsResponse) *string { return v.FontSize }).(pulumi.StringPtrOutput)
}

func (o TerminalSettingsResponseOutput) FontStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminalSettingsResponse) *string { return v.FontStyle }).(pulumi.StringPtrOutput)
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

type UserPropertiesResponse struct {
	PreferredLocation  string                   `pulumi:"preferredLocation"`
	PreferredOsType    string                   `pulumi:"preferredOsType"`
	PreferredShellType string                   `pulumi:"preferredShellType"`
	StorageProfile     StorageProfileResponse   `pulumi:"storageProfile"`
	TerminalSettings   TerminalSettingsResponse `pulumi:"terminalSettings"`
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

func init() {
	pulumi.RegisterOutputType(ConsoleCreatePropertiesOutput{})
	pulumi.RegisterOutputType(ConsolePropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(TerminalSettingsOutput{})
	pulumi.RegisterOutputType(TerminalSettingsResponseOutput{})
	pulumi.RegisterOutputType(UserPropertiesOutput{})
	pulumi.RegisterOutputType(UserPropertiesResponseOutput{})
}
