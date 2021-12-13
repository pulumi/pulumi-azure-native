


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InputLinuxParameters struct {
	ClassificationsToInclude  []string `pulumi:"classificationsToInclude"`
	PackageNameMasksToExclude []string `pulumi:"packageNameMasksToExclude"`
	PackageNameMasksToInclude []string `pulumi:"packageNameMasksToInclude"`
}





type InputLinuxParametersInput interface {
	pulumi.Input

	ToInputLinuxParametersOutput() InputLinuxParametersOutput
	ToInputLinuxParametersOutputWithContext(context.Context) InputLinuxParametersOutput
}

type InputLinuxParametersArgs struct {
	ClassificationsToInclude  pulumi.StringArrayInput `pulumi:"classificationsToInclude"`
	PackageNameMasksToExclude pulumi.StringArrayInput `pulumi:"packageNameMasksToExclude"`
	PackageNameMasksToInclude pulumi.StringArrayInput `pulumi:"packageNameMasksToInclude"`
}

func (InputLinuxParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputLinuxParameters)(nil)).Elem()
}

func (i InputLinuxParametersArgs) ToInputLinuxParametersOutput() InputLinuxParametersOutput {
	return i.ToInputLinuxParametersOutputWithContext(context.Background())
}

func (i InputLinuxParametersArgs) ToInputLinuxParametersOutputWithContext(ctx context.Context) InputLinuxParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputLinuxParametersOutput)
}

func (i InputLinuxParametersArgs) ToInputLinuxParametersPtrOutput() InputLinuxParametersPtrOutput {
	return i.ToInputLinuxParametersPtrOutputWithContext(context.Background())
}

func (i InputLinuxParametersArgs) ToInputLinuxParametersPtrOutputWithContext(ctx context.Context) InputLinuxParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputLinuxParametersOutput).ToInputLinuxParametersPtrOutputWithContext(ctx)
}









type InputLinuxParametersPtrInput interface {
	pulumi.Input

	ToInputLinuxParametersPtrOutput() InputLinuxParametersPtrOutput
	ToInputLinuxParametersPtrOutputWithContext(context.Context) InputLinuxParametersPtrOutput
}

type inputLinuxParametersPtrType InputLinuxParametersArgs

func InputLinuxParametersPtr(v *InputLinuxParametersArgs) InputLinuxParametersPtrInput {
	return (*inputLinuxParametersPtrType)(v)
}

func (*inputLinuxParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputLinuxParameters)(nil)).Elem()
}

func (i *inputLinuxParametersPtrType) ToInputLinuxParametersPtrOutput() InputLinuxParametersPtrOutput {
	return i.ToInputLinuxParametersPtrOutputWithContext(context.Background())
}

func (i *inputLinuxParametersPtrType) ToInputLinuxParametersPtrOutputWithContext(ctx context.Context) InputLinuxParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputLinuxParametersPtrOutput)
}

type InputLinuxParametersOutput struct{ *pulumi.OutputState }

func (InputLinuxParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputLinuxParameters)(nil)).Elem()
}

func (o InputLinuxParametersOutput) ToInputLinuxParametersOutput() InputLinuxParametersOutput {
	return o
}

func (o InputLinuxParametersOutput) ToInputLinuxParametersOutputWithContext(ctx context.Context) InputLinuxParametersOutput {
	return o
}

func (o InputLinuxParametersOutput) ToInputLinuxParametersPtrOutput() InputLinuxParametersPtrOutput {
	return o.ToInputLinuxParametersPtrOutputWithContext(context.Background())
}

func (o InputLinuxParametersOutput) ToInputLinuxParametersPtrOutputWithContext(ctx context.Context) InputLinuxParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputLinuxParameters) *InputLinuxParameters {
		return &v
	}).(InputLinuxParametersPtrOutput)
}

func (o InputLinuxParametersOutput) ClassificationsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputLinuxParameters) []string { return v.ClassificationsToInclude }).(pulumi.StringArrayOutput)
}

func (o InputLinuxParametersOutput) PackageNameMasksToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputLinuxParameters) []string { return v.PackageNameMasksToExclude }).(pulumi.StringArrayOutput)
}

func (o InputLinuxParametersOutput) PackageNameMasksToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputLinuxParameters) []string { return v.PackageNameMasksToInclude }).(pulumi.StringArrayOutput)
}

type InputLinuxParametersPtrOutput struct{ *pulumi.OutputState }

func (InputLinuxParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputLinuxParameters)(nil)).Elem()
}

func (o InputLinuxParametersPtrOutput) ToInputLinuxParametersPtrOutput() InputLinuxParametersPtrOutput {
	return o
}

func (o InputLinuxParametersPtrOutput) ToInputLinuxParametersPtrOutputWithContext(ctx context.Context) InputLinuxParametersPtrOutput {
	return o
}

func (o InputLinuxParametersPtrOutput) Elem() InputLinuxParametersOutput {
	return o.ApplyT(func(v *InputLinuxParameters) InputLinuxParameters {
		if v != nil {
			return *v
		}
		var ret InputLinuxParameters
		return ret
	}).(InputLinuxParametersOutput)
}

func (o InputLinuxParametersPtrOutput) ClassificationsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputLinuxParameters) []string {
		if v == nil {
			return nil
		}
		return v.ClassificationsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o InputLinuxParametersPtrOutput) PackageNameMasksToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputLinuxParameters) []string {
		if v == nil {
			return nil
		}
		return v.PackageNameMasksToExclude
	}).(pulumi.StringArrayOutput)
}

func (o InputLinuxParametersPtrOutput) PackageNameMasksToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputLinuxParameters) []string {
		if v == nil {
			return nil
		}
		return v.PackageNameMasksToInclude
	}).(pulumi.StringArrayOutput)
}

type InputLinuxParametersResponse struct {
	ClassificationsToInclude  []string `pulumi:"classificationsToInclude"`
	PackageNameMasksToExclude []string `pulumi:"packageNameMasksToExclude"`
	PackageNameMasksToInclude []string `pulumi:"packageNameMasksToInclude"`
}

type InputLinuxParametersResponseOutput struct{ *pulumi.OutputState }

func (InputLinuxParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputLinuxParametersResponse)(nil)).Elem()
}

func (o InputLinuxParametersResponseOutput) ToInputLinuxParametersResponseOutput() InputLinuxParametersResponseOutput {
	return o
}

func (o InputLinuxParametersResponseOutput) ToInputLinuxParametersResponseOutputWithContext(ctx context.Context) InputLinuxParametersResponseOutput {
	return o
}

func (o InputLinuxParametersResponseOutput) ClassificationsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputLinuxParametersResponse) []string { return v.ClassificationsToInclude }).(pulumi.StringArrayOutput)
}

func (o InputLinuxParametersResponseOutput) PackageNameMasksToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputLinuxParametersResponse) []string { return v.PackageNameMasksToExclude }).(pulumi.StringArrayOutput)
}

func (o InputLinuxParametersResponseOutput) PackageNameMasksToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputLinuxParametersResponse) []string { return v.PackageNameMasksToInclude }).(pulumi.StringArrayOutput)
}

type InputLinuxParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (InputLinuxParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputLinuxParametersResponse)(nil)).Elem()
}

func (o InputLinuxParametersResponsePtrOutput) ToInputLinuxParametersResponsePtrOutput() InputLinuxParametersResponsePtrOutput {
	return o
}

func (o InputLinuxParametersResponsePtrOutput) ToInputLinuxParametersResponsePtrOutputWithContext(ctx context.Context) InputLinuxParametersResponsePtrOutput {
	return o
}

func (o InputLinuxParametersResponsePtrOutput) Elem() InputLinuxParametersResponseOutput {
	return o.ApplyT(func(v *InputLinuxParametersResponse) InputLinuxParametersResponse {
		if v != nil {
			return *v
		}
		var ret InputLinuxParametersResponse
		return ret
	}).(InputLinuxParametersResponseOutput)
}

func (o InputLinuxParametersResponsePtrOutput) ClassificationsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputLinuxParametersResponse) []string {
		if v == nil {
			return nil
		}
		return v.ClassificationsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o InputLinuxParametersResponsePtrOutput) PackageNameMasksToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputLinuxParametersResponse) []string {
		if v == nil {
			return nil
		}
		return v.PackageNameMasksToExclude
	}).(pulumi.StringArrayOutput)
}

func (o InputLinuxParametersResponsePtrOutput) PackageNameMasksToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputLinuxParametersResponse) []string {
		if v == nil {
			return nil
		}
		return v.PackageNameMasksToInclude
	}).(pulumi.StringArrayOutput)
}

type InputPatchConfiguration struct {
	LinuxParameters   *InputLinuxParameters   `pulumi:"linuxParameters"`
	PostTasks         []TaskProperties        `pulumi:"postTasks"`
	PreTasks          []TaskProperties        `pulumi:"preTasks"`
	RebootSetting     *string                 `pulumi:"rebootSetting"`
	WindowsParameters *InputWindowsParameters `pulumi:"windowsParameters"`
}


func (val *InputPatchConfiguration) Defaults() *InputPatchConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RebootSetting) {
		rebootSetting_ := "IfRequired"
		tmp.RebootSetting = &rebootSetting_
	}
	return &tmp
}





type InputPatchConfigurationInput interface {
	pulumi.Input

	ToInputPatchConfigurationOutput() InputPatchConfigurationOutput
	ToInputPatchConfigurationOutputWithContext(context.Context) InputPatchConfigurationOutput
}

type InputPatchConfigurationArgs struct {
	LinuxParameters   InputLinuxParametersPtrInput   `pulumi:"linuxParameters"`
	PostTasks         TaskPropertiesArrayInput       `pulumi:"postTasks"`
	PreTasks          TaskPropertiesArrayInput       `pulumi:"preTasks"`
	RebootSetting     pulumi.StringPtrInput          `pulumi:"rebootSetting"`
	WindowsParameters InputWindowsParametersPtrInput `pulumi:"windowsParameters"`
}

func (InputPatchConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputPatchConfiguration)(nil)).Elem()
}

func (i InputPatchConfigurationArgs) ToInputPatchConfigurationOutput() InputPatchConfigurationOutput {
	return i.ToInputPatchConfigurationOutputWithContext(context.Background())
}

func (i InputPatchConfigurationArgs) ToInputPatchConfigurationOutputWithContext(ctx context.Context) InputPatchConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPatchConfigurationOutput)
}

func (i InputPatchConfigurationArgs) ToInputPatchConfigurationPtrOutput() InputPatchConfigurationPtrOutput {
	return i.ToInputPatchConfigurationPtrOutputWithContext(context.Background())
}

func (i InputPatchConfigurationArgs) ToInputPatchConfigurationPtrOutputWithContext(ctx context.Context) InputPatchConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPatchConfigurationOutput).ToInputPatchConfigurationPtrOutputWithContext(ctx)
}









type InputPatchConfigurationPtrInput interface {
	pulumi.Input

	ToInputPatchConfigurationPtrOutput() InputPatchConfigurationPtrOutput
	ToInputPatchConfigurationPtrOutputWithContext(context.Context) InputPatchConfigurationPtrOutput
}

type inputPatchConfigurationPtrType InputPatchConfigurationArgs

func InputPatchConfigurationPtr(v *InputPatchConfigurationArgs) InputPatchConfigurationPtrInput {
	return (*inputPatchConfigurationPtrType)(v)
}

func (*inputPatchConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputPatchConfiguration)(nil)).Elem()
}

func (i *inputPatchConfigurationPtrType) ToInputPatchConfigurationPtrOutput() InputPatchConfigurationPtrOutput {
	return i.ToInputPatchConfigurationPtrOutputWithContext(context.Background())
}

func (i *inputPatchConfigurationPtrType) ToInputPatchConfigurationPtrOutputWithContext(ctx context.Context) InputPatchConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPatchConfigurationPtrOutput)
}

type InputPatchConfigurationOutput struct{ *pulumi.OutputState }

func (InputPatchConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputPatchConfiguration)(nil)).Elem()
}

func (o InputPatchConfigurationOutput) ToInputPatchConfigurationOutput() InputPatchConfigurationOutput {
	return o
}

func (o InputPatchConfigurationOutput) ToInputPatchConfigurationOutputWithContext(ctx context.Context) InputPatchConfigurationOutput {
	return o
}

func (o InputPatchConfigurationOutput) ToInputPatchConfigurationPtrOutput() InputPatchConfigurationPtrOutput {
	return o.ToInputPatchConfigurationPtrOutputWithContext(context.Background())
}

func (o InputPatchConfigurationOutput) ToInputPatchConfigurationPtrOutputWithContext(ctx context.Context) InputPatchConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputPatchConfiguration) *InputPatchConfiguration {
		return &v
	}).(InputPatchConfigurationPtrOutput)
}

func (o InputPatchConfigurationOutput) LinuxParameters() InputLinuxParametersPtrOutput {
	return o.ApplyT(func(v InputPatchConfiguration) *InputLinuxParameters { return v.LinuxParameters }).(InputLinuxParametersPtrOutput)
}

func (o InputPatchConfigurationOutput) PostTasks() TaskPropertiesArrayOutput {
	return o.ApplyT(func(v InputPatchConfiguration) []TaskProperties { return v.PostTasks }).(TaskPropertiesArrayOutput)
}

func (o InputPatchConfigurationOutput) PreTasks() TaskPropertiesArrayOutput {
	return o.ApplyT(func(v InputPatchConfiguration) []TaskProperties { return v.PreTasks }).(TaskPropertiesArrayOutput)
}

func (o InputPatchConfigurationOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputPatchConfiguration) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

func (o InputPatchConfigurationOutput) WindowsParameters() InputWindowsParametersPtrOutput {
	return o.ApplyT(func(v InputPatchConfiguration) *InputWindowsParameters { return v.WindowsParameters }).(InputWindowsParametersPtrOutput)
}

type InputPatchConfigurationPtrOutput struct{ *pulumi.OutputState }

func (InputPatchConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputPatchConfiguration)(nil)).Elem()
}

func (o InputPatchConfigurationPtrOutput) ToInputPatchConfigurationPtrOutput() InputPatchConfigurationPtrOutput {
	return o
}

func (o InputPatchConfigurationPtrOutput) ToInputPatchConfigurationPtrOutputWithContext(ctx context.Context) InputPatchConfigurationPtrOutput {
	return o
}

func (o InputPatchConfigurationPtrOutput) Elem() InputPatchConfigurationOutput {
	return o.ApplyT(func(v *InputPatchConfiguration) InputPatchConfiguration {
		if v != nil {
			return *v
		}
		var ret InputPatchConfiguration
		return ret
	}).(InputPatchConfigurationOutput)
}

func (o InputPatchConfigurationPtrOutput) LinuxParameters() InputLinuxParametersPtrOutput {
	return o.ApplyT(func(v *InputPatchConfiguration) *InputLinuxParameters {
		if v == nil {
			return nil
		}
		return v.LinuxParameters
	}).(InputLinuxParametersPtrOutput)
}

func (o InputPatchConfigurationPtrOutput) PostTasks() TaskPropertiesArrayOutput {
	return o.ApplyT(func(v *InputPatchConfiguration) []TaskProperties {
		if v == nil {
			return nil
		}
		return v.PostTasks
	}).(TaskPropertiesArrayOutput)
}

func (o InputPatchConfigurationPtrOutput) PreTasks() TaskPropertiesArrayOutput {
	return o.ApplyT(func(v *InputPatchConfiguration) []TaskProperties {
		if v == nil {
			return nil
		}
		return v.PreTasks
	}).(TaskPropertiesArrayOutput)
}

func (o InputPatchConfigurationPtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InputPatchConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

func (o InputPatchConfigurationPtrOutput) WindowsParameters() InputWindowsParametersPtrOutput {
	return o.ApplyT(func(v *InputPatchConfiguration) *InputWindowsParameters {
		if v == nil {
			return nil
		}
		return v.WindowsParameters
	}).(InputWindowsParametersPtrOutput)
}

type InputPatchConfigurationResponse struct {
	LinuxParameters   *InputLinuxParametersResponse   `pulumi:"linuxParameters"`
	PostTasks         []TaskPropertiesResponse        `pulumi:"postTasks"`
	PreTasks          []TaskPropertiesResponse        `pulumi:"preTasks"`
	RebootSetting     *string                         `pulumi:"rebootSetting"`
	WindowsParameters *InputWindowsParametersResponse `pulumi:"windowsParameters"`
}


func (val *InputPatchConfigurationResponse) Defaults() *InputPatchConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RebootSetting) {
		rebootSetting_ := "IfRequired"
		tmp.RebootSetting = &rebootSetting_
	}
	return &tmp
}

type InputPatchConfigurationResponseOutput struct{ *pulumi.OutputState }

func (InputPatchConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputPatchConfigurationResponse)(nil)).Elem()
}

func (o InputPatchConfigurationResponseOutput) ToInputPatchConfigurationResponseOutput() InputPatchConfigurationResponseOutput {
	return o
}

func (o InputPatchConfigurationResponseOutput) ToInputPatchConfigurationResponseOutputWithContext(ctx context.Context) InputPatchConfigurationResponseOutput {
	return o
}

func (o InputPatchConfigurationResponseOutput) LinuxParameters() InputLinuxParametersResponsePtrOutput {
	return o.ApplyT(func(v InputPatchConfigurationResponse) *InputLinuxParametersResponse { return v.LinuxParameters }).(InputLinuxParametersResponsePtrOutput)
}

func (o InputPatchConfigurationResponseOutput) PostTasks() TaskPropertiesResponseArrayOutput {
	return o.ApplyT(func(v InputPatchConfigurationResponse) []TaskPropertiesResponse { return v.PostTasks }).(TaskPropertiesResponseArrayOutput)
}

func (o InputPatchConfigurationResponseOutput) PreTasks() TaskPropertiesResponseArrayOutput {
	return o.ApplyT(func(v InputPatchConfigurationResponse) []TaskPropertiesResponse { return v.PreTasks }).(TaskPropertiesResponseArrayOutput)
}

func (o InputPatchConfigurationResponseOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputPatchConfigurationResponse) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

func (o InputPatchConfigurationResponseOutput) WindowsParameters() InputWindowsParametersResponsePtrOutput {
	return o.ApplyT(func(v InputPatchConfigurationResponse) *InputWindowsParametersResponse { return v.WindowsParameters }).(InputWindowsParametersResponsePtrOutput)
}

type InputPatchConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (InputPatchConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputPatchConfigurationResponse)(nil)).Elem()
}

func (o InputPatchConfigurationResponsePtrOutput) ToInputPatchConfigurationResponsePtrOutput() InputPatchConfigurationResponsePtrOutput {
	return o
}

func (o InputPatchConfigurationResponsePtrOutput) ToInputPatchConfigurationResponsePtrOutputWithContext(ctx context.Context) InputPatchConfigurationResponsePtrOutput {
	return o
}

func (o InputPatchConfigurationResponsePtrOutput) Elem() InputPatchConfigurationResponseOutput {
	return o.ApplyT(func(v *InputPatchConfigurationResponse) InputPatchConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret InputPatchConfigurationResponse
		return ret
	}).(InputPatchConfigurationResponseOutput)
}

func (o InputPatchConfigurationResponsePtrOutput) LinuxParameters() InputLinuxParametersResponsePtrOutput {
	return o.ApplyT(func(v *InputPatchConfigurationResponse) *InputLinuxParametersResponse {
		if v == nil {
			return nil
		}
		return v.LinuxParameters
	}).(InputLinuxParametersResponsePtrOutput)
}

func (o InputPatchConfigurationResponsePtrOutput) PostTasks() TaskPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *InputPatchConfigurationResponse) []TaskPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.PostTasks
	}).(TaskPropertiesResponseArrayOutput)
}

func (o InputPatchConfigurationResponsePtrOutput) PreTasks() TaskPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *InputPatchConfigurationResponse) []TaskPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.PreTasks
	}).(TaskPropertiesResponseArrayOutput)
}

func (o InputPatchConfigurationResponsePtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InputPatchConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

func (o InputPatchConfigurationResponsePtrOutput) WindowsParameters() InputWindowsParametersResponsePtrOutput {
	return o.ApplyT(func(v *InputPatchConfigurationResponse) *InputWindowsParametersResponse {
		if v == nil {
			return nil
		}
		return v.WindowsParameters
	}).(InputWindowsParametersResponsePtrOutput)
}

type InputWindowsParameters struct {
	ClassificationsToInclude  []string `pulumi:"classificationsToInclude"`
	ExcludeKbsRequiringReboot *bool    `pulumi:"excludeKbsRequiringReboot"`
	KbNumbersToExclude        []string `pulumi:"kbNumbersToExclude"`
	KbNumbersToInclude        []string `pulumi:"kbNumbersToInclude"`
}





type InputWindowsParametersInput interface {
	pulumi.Input

	ToInputWindowsParametersOutput() InputWindowsParametersOutput
	ToInputWindowsParametersOutputWithContext(context.Context) InputWindowsParametersOutput
}

type InputWindowsParametersArgs struct {
	ClassificationsToInclude  pulumi.StringArrayInput `pulumi:"classificationsToInclude"`
	ExcludeKbsRequiringReboot pulumi.BoolPtrInput     `pulumi:"excludeKbsRequiringReboot"`
	KbNumbersToExclude        pulumi.StringArrayInput `pulumi:"kbNumbersToExclude"`
	KbNumbersToInclude        pulumi.StringArrayInput `pulumi:"kbNumbersToInclude"`
}

func (InputWindowsParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputWindowsParameters)(nil)).Elem()
}

func (i InputWindowsParametersArgs) ToInputWindowsParametersOutput() InputWindowsParametersOutput {
	return i.ToInputWindowsParametersOutputWithContext(context.Background())
}

func (i InputWindowsParametersArgs) ToInputWindowsParametersOutputWithContext(ctx context.Context) InputWindowsParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputWindowsParametersOutput)
}

func (i InputWindowsParametersArgs) ToInputWindowsParametersPtrOutput() InputWindowsParametersPtrOutput {
	return i.ToInputWindowsParametersPtrOutputWithContext(context.Background())
}

func (i InputWindowsParametersArgs) ToInputWindowsParametersPtrOutputWithContext(ctx context.Context) InputWindowsParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputWindowsParametersOutput).ToInputWindowsParametersPtrOutputWithContext(ctx)
}









type InputWindowsParametersPtrInput interface {
	pulumi.Input

	ToInputWindowsParametersPtrOutput() InputWindowsParametersPtrOutput
	ToInputWindowsParametersPtrOutputWithContext(context.Context) InputWindowsParametersPtrOutput
}

type inputWindowsParametersPtrType InputWindowsParametersArgs

func InputWindowsParametersPtr(v *InputWindowsParametersArgs) InputWindowsParametersPtrInput {
	return (*inputWindowsParametersPtrType)(v)
}

func (*inputWindowsParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputWindowsParameters)(nil)).Elem()
}

func (i *inputWindowsParametersPtrType) ToInputWindowsParametersPtrOutput() InputWindowsParametersPtrOutput {
	return i.ToInputWindowsParametersPtrOutputWithContext(context.Background())
}

func (i *inputWindowsParametersPtrType) ToInputWindowsParametersPtrOutputWithContext(ctx context.Context) InputWindowsParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputWindowsParametersPtrOutput)
}

type InputWindowsParametersOutput struct{ *pulumi.OutputState }

func (InputWindowsParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputWindowsParameters)(nil)).Elem()
}

func (o InputWindowsParametersOutput) ToInputWindowsParametersOutput() InputWindowsParametersOutput {
	return o
}

func (o InputWindowsParametersOutput) ToInputWindowsParametersOutputWithContext(ctx context.Context) InputWindowsParametersOutput {
	return o
}

func (o InputWindowsParametersOutput) ToInputWindowsParametersPtrOutput() InputWindowsParametersPtrOutput {
	return o.ToInputWindowsParametersPtrOutputWithContext(context.Background())
}

func (o InputWindowsParametersOutput) ToInputWindowsParametersPtrOutputWithContext(ctx context.Context) InputWindowsParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputWindowsParameters) *InputWindowsParameters {
		return &v
	}).(InputWindowsParametersPtrOutput)
}

func (o InputWindowsParametersOutput) ClassificationsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputWindowsParameters) []string { return v.ClassificationsToInclude }).(pulumi.StringArrayOutput)
}

func (o InputWindowsParametersOutput) ExcludeKbsRequiringReboot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InputWindowsParameters) *bool { return v.ExcludeKbsRequiringReboot }).(pulumi.BoolPtrOutput)
}

func (o InputWindowsParametersOutput) KbNumbersToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputWindowsParameters) []string { return v.KbNumbersToExclude }).(pulumi.StringArrayOutput)
}

func (o InputWindowsParametersOutput) KbNumbersToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputWindowsParameters) []string { return v.KbNumbersToInclude }).(pulumi.StringArrayOutput)
}

type InputWindowsParametersPtrOutput struct{ *pulumi.OutputState }

func (InputWindowsParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputWindowsParameters)(nil)).Elem()
}

func (o InputWindowsParametersPtrOutput) ToInputWindowsParametersPtrOutput() InputWindowsParametersPtrOutput {
	return o
}

func (o InputWindowsParametersPtrOutput) ToInputWindowsParametersPtrOutputWithContext(ctx context.Context) InputWindowsParametersPtrOutput {
	return o
}

func (o InputWindowsParametersPtrOutput) Elem() InputWindowsParametersOutput {
	return o.ApplyT(func(v *InputWindowsParameters) InputWindowsParameters {
		if v != nil {
			return *v
		}
		var ret InputWindowsParameters
		return ret
	}).(InputWindowsParametersOutput)
}

func (o InputWindowsParametersPtrOutput) ClassificationsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputWindowsParameters) []string {
		if v == nil {
			return nil
		}
		return v.ClassificationsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o InputWindowsParametersPtrOutput) ExcludeKbsRequiringReboot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InputWindowsParameters) *bool {
		if v == nil {
			return nil
		}
		return v.ExcludeKbsRequiringReboot
	}).(pulumi.BoolPtrOutput)
}

func (o InputWindowsParametersPtrOutput) KbNumbersToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputWindowsParameters) []string {
		if v == nil {
			return nil
		}
		return v.KbNumbersToExclude
	}).(pulumi.StringArrayOutput)
}

func (o InputWindowsParametersPtrOutput) KbNumbersToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputWindowsParameters) []string {
		if v == nil {
			return nil
		}
		return v.KbNumbersToInclude
	}).(pulumi.StringArrayOutput)
}

type InputWindowsParametersResponse struct {
	ClassificationsToInclude  []string `pulumi:"classificationsToInclude"`
	ExcludeKbsRequiringReboot *bool    `pulumi:"excludeKbsRequiringReboot"`
	KbNumbersToExclude        []string `pulumi:"kbNumbersToExclude"`
	KbNumbersToInclude        []string `pulumi:"kbNumbersToInclude"`
}

type InputWindowsParametersResponseOutput struct{ *pulumi.OutputState }

func (InputWindowsParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputWindowsParametersResponse)(nil)).Elem()
}

func (o InputWindowsParametersResponseOutput) ToInputWindowsParametersResponseOutput() InputWindowsParametersResponseOutput {
	return o
}

func (o InputWindowsParametersResponseOutput) ToInputWindowsParametersResponseOutputWithContext(ctx context.Context) InputWindowsParametersResponseOutput {
	return o
}

func (o InputWindowsParametersResponseOutput) ClassificationsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputWindowsParametersResponse) []string { return v.ClassificationsToInclude }).(pulumi.StringArrayOutput)
}

func (o InputWindowsParametersResponseOutput) ExcludeKbsRequiringReboot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InputWindowsParametersResponse) *bool { return v.ExcludeKbsRequiringReboot }).(pulumi.BoolPtrOutput)
}

func (o InputWindowsParametersResponseOutput) KbNumbersToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputWindowsParametersResponse) []string { return v.KbNumbersToExclude }).(pulumi.StringArrayOutput)
}

func (o InputWindowsParametersResponseOutput) KbNumbersToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InputWindowsParametersResponse) []string { return v.KbNumbersToInclude }).(pulumi.StringArrayOutput)
}

type InputWindowsParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (InputWindowsParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputWindowsParametersResponse)(nil)).Elem()
}

func (o InputWindowsParametersResponsePtrOutput) ToInputWindowsParametersResponsePtrOutput() InputWindowsParametersResponsePtrOutput {
	return o
}

func (o InputWindowsParametersResponsePtrOutput) ToInputWindowsParametersResponsePtrOutputWithContext(ctx context.Context) InputWindowsParametersResponsePtrOutput {
	return o
}

func (o InputWindowsParametersResponsePtrOutput) Elem() InputWindowsParametersResponseOutput {
	return o.ApplyT(func(v *InputWindowsParametersResponse) InputWindowsParametersResponse {
		if v != nil {
			return *v
		}
		var ret InputWindowsParametersResponse
		return ret
	}).(InputWindowsParametersResponseOutput)
}

func (o InputWindowsParametersResponsePtrOutput) ClassificationsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputWindowsParametersResponse) []string {
		if v == nil {
			return nil
		}
		return v.ClassificationsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o InputWindowsParametersResponsePtrOutput) ExcludeKbsRequiringReboot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InputWindowsParametersResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ExcludeKbsRequiringReboot
	}).(pulumi.BoolPtrOutput)
}

func (o InputWindowsParametersResponsePtrOutput) KbNumbersToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputWindowsParametersResponse) []string {
		if v == nil {
			return nil
		}
		return v.KbNumbersToExclude
	}).(pulumi.StringArrayOutput)
}

func (o InputWindowsParametersResponsePtrOutput) KbNumbersToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InputWindowsParametersResponse) []string {
		if v == nil {
			return nil
		}
		return v.KbNumbersToInclude
	}).(pulumi.StringArrayOutput)
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

type TaskProperties struct {
	Parameters map[string]string `pulumi:"parameters"`
	Source     *string           `pulumi:"source"`
	TaskScope  *string           `pulumi:"taskScope"`
}


func (val *TaskProperties) Defaults() *TaskProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.TaskScope) {
		taskScope_ := "Global"
		tmp.TaskScope = &taskScope_
	}
	return &tmp
}





type TaskPropertiesInput interface {
	pulumi.Input

	ToTaskPropertiesOutput() TaskPropertiesOutput
	ToTaskPropertiesOutputWithContext(context.Context) TaskPropertiesOutput
}

type TaskPropertiesArgs struct {
	Parameters pulumi.StringMapInput `pulumi:"parameters"`
	Source     pulumi.StringPtrInput `pulumi:"source"`
	TaskScope  pulumi.StringPtrInput `pulumi:"taskScope"`
}

func (TaskPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskProperties)(nil)).Elem()
}

func (i TaskPropertiesArgs) ToTaskPropertiesOutput() TaskPropertiesOutput {
	return i.ToTaskPropertiesOutputWithContext(context.Background())
}

func (i TaskPropertiesArgs) ToTaskPropertiesOutputWithContext(ctx context.Context) TaskPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskPropertiesOutput)
}





type TaskPropertiesArrayInput interface {
	pulumi.Input

	ToTaskPropertiesArrayOutput() TaskPropertiesArrayOutput
	ToTaskPropertiesArrayOutputWithContext(context.Context) TaskPropertiesArrayOutput
}

type TaskPropertiesArray []TaskPropertiesInput

func (TaskPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TaskProperties)(nil)).Elem()
}

func (i TaskPropertiesArray) ToTaskPropertiesArrayOutput() TaskPropertiesArrayOutput {
	return i.ToTaskPropertiesArrayOutputWithContext(context.Background())
}

func (i TaskPropertiesArray) ToTaskPropertiesArrayOutputWithContext(ctx context.Context) TaskPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskPropertiesArrayOutput)
}

type TaskPropertiesOutput struct{ *pulumi.OutputState }

func (TaskPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskProperties)(nil)).Elem()
}

func (o TaskPropertiesOutput) ToTaskPropertiesOutput() TaskPropertiesOutput {
	return o
}

func (o TaskPropertiesOutput) ToTaskPropertiesOutputWithContext(ctx context.Context) TaskPropertiesOutput {
	return o
}

func (o TaskPropertiesOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v TaskProperties) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o TaskPropertiesOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskProperties) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o TaskPropertiesOutput) TaskScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskProperties) *string { return v.TaskScope }).(pulumi.StringPtrOutput)
}

type TaskPropertiesArrayOutput struct{ *pulumi.OutputState }

func (TaskPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TaskProperties)(nil)).Elem()
}

func (o TaskPropertiesArrayOutput) ToTaskPropertiesArrayOutput() TaskPropertiesArrayOutput {
	return o
}

func (o TaskPropertiesArrayOutput) ToTaskPropertiesArrayOutputWithContext(ctx context.Context) TaskPropertiesArrayOutput {
	return o
}

func (o TaskPropertiesArrayOutput) Index(i pulumi.IntInput) TaskPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TaskProperties {
		return vs[0].([]TaskProperties)[vs[1].(int)]
	}).(TaskPropertiesOutput)
}

type TaskPropertiesResponse struct {
	Parameters map[string]string `pulumi:"parameters"`
	Source     *string           `pulumi:"source"`
	TaskScope  *string           `pulumi:"taskScope"`
}


func (val *TaskPropertiesResponse) Defaults() *TaskPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.TaskScope) {
		taskScope_ := "Global"
		tmp.TaskScope = &taskScope_
	}
	return &tmp
}

type TaskPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TaskPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskPropertiesResponse)(nil)).Elem()
}

func (o TaskPropertiesResponseOutput) ToTaskPropertiesResponseOutput() TaskPropertiesResponseOutput {
	return o
}

func (o TaskPropertiesResponseOutput) ToTaskPropertiesResponseOutputWithContext(ctx context.Context) TaskPropertiesResponseOutput {
	return o
}

func (o TaskPropertiesResponseOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v TaskPropertiesResponse) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o TaskPropertiesResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskPropertiesResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o TaskPropertiesResponseOutput) TaskScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskPropertiesResponse) *string { return v.TaskScope }).(pulumi.StringPtrOutput)
}

type TaskPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (TaskPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TaskPropertiesResponse)(nil)).Elem()
}

func (o TaskPropertiesResponseArrayOutput) ToTaskPropertiesResponseArrayOutput() TaskPropertiesResponseArrayOutput {
	return o
}

func (o TaskPropertiesResponseArrayOutput) ToTaskPropertiesResponseArrayOutputWithContext(ctx context.Context) TaskPropertiesResponseArrayOutput {
	return o
}

func (o TaskPropertiesResponseArrayOutput) Index(i pulumi.IntInput) TaskPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TaskPropertiesResponse {
		return vs[0].([]TaskPropertiesResponse)[vs[1].(int)]
	}).(TaskPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(InputLinuxParametersOutput{})
	pulumi.RegisterOutputType(InputLinuxParametersPtrOutput{})
	pulumi.RegisterOutputType(InputLinuxParametersResponseOutput{})
	pulumi.RegisterOutputType(InputLinuxParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(InputPatchConfigurationOutput{})
	pulumi.RegisterOutputType(InputPatchConfigurationPtrOutput{})
	pulumi.RegisterOutputType(InputPatchConfigurationResponseOutput{})
	pulumi.RegisterOutputType(InputPatchConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(InputWindowsParametersOutput{})
	pulumi.RegisterOutputType(InputWindowsParametersPtrOutput{})
	pulumi.RegisterOutputType(InputWindowsParametersResponseOutput{})
	pulumi.RegisterOutputType(InputWindowsParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TaskPropertiesOutput{})
	pulumi.RegisterOutputType(TaskPropertiesArrayOutput{})
	pulumi.RegisterOutputType(TaskPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TaskPropertiesResponseArrayOutput{})
}
