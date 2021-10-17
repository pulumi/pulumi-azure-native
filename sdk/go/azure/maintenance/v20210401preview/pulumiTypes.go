


package v20210401preview

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





type InputLinuxParametersResponseInput interface {
	pulumi.Input

	ToInputLinuxParametersResponseOutput() InputLinuxParametersResponseOutput
	ToInputLinuxParametersResponseOutputWithContext(context.Context) InputLinuxParametersResponseOutput
}

type InputLinuxParametersResponseArgs struct {
	ClassificationsToInclude  pulumi.StringArrayInput `pulumi:"classificationsToInclude"`
	PackageNameMasksToExclude pulumi.StringArrayInput `pulumi:"packageNameMasksToExclude"`
	PackageNameMasksToInclude pulumi.StringArrayInput `pulumi:"packageNameMasksToInclude"`
}

func (InputLinuxParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputLinuxParametersResponse)(nil)).Elem()
}

func (i InputLinuxParametersResponseArgs) ToInputLinuxParametersResponseOutput() InputLinuxParametersResponseOutput {
	return i.ToInputLinuxParametersResponseOutputWithContext(context.Background())
}

func (i InputLinuxParametersResponseArgs) ToInputLinuxParametersResponseOutputWithContext(ctx context.Context) InputLinuxParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputLinuxParametersResponseOutput)
}

func (i InputLinuxParametersResponseArgs) ToInputLinuxParametersResponsePtrOutput() InputLinuxParametersResponsePtrOutput {
	return i.ToInputLinuxParametersResponsePtrOutputWithContext(context.Background())
}

func (i InputLinuxParametersResponseArgs) ToInputLinuxParametersResponsePtrOutputWithContext(ctx context.Context) InputLinuxParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputLinuxParametersResponseOutput).ToInputLinuxParametersResponsePtrOutputWithContext(ctx)
}









type InputLinuxParametersResponsePtrInput interface {
	pulumi.Input

	ToInputLinuxParametersResponsePtrOutput() InputLinuxParametersResponsePtrOutput
	ToInputLinuxParametersResponsePtrOutputWithContext(context.Context) InputLinuxParametersResponsePtrOutput
}

type inputLinuxParametersResponsePtrType InputLinuxParametersResponseArgs

func InputLinuxParametersResponsePtr(v *InputLinuxParametersResponseArgs) InputLinuxParametersResponsePtrInput {
	return (*inputLinuxParametersResponsePtrType)(v)
}

func (*inputLinuxParametersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputLinuxParametersResponse)(nil)).Elem()
}

func (i *inputLinuxParametersResponsePtrType) ToInputLinuxParametersResponsePtrOutput() InputLinuxParametersResponsePtrOutput {
	return i.ToInputLinuxParametersResponsePtrOutputWithContext(context.Background())
}

func (i *inputLinuxParametersResponsePtrType) ToInputLinuxParametersResponsePtrOutputWithContext(ctx context.Context) InputLinuxParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputLinuxParametersResponsePtrOutput)
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

func (o InputLinuxParametersResponseOutput) ToInputLinuxParametersResponsePtrOutput() InputLinuxParametersResponsePtrOutput {
	return o.ToInputLinuxParametersResponsePtrOutputWithContext(context.Background())
}

func (o InputLinuxParametersResponseOutput) ToInputLinuxParametersResponsePtrOutputWithContext(ctx context.Context) InputLinuxParametersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputLinuxParametersResponse) *InputLinuxParametersResponse {
		return &v
	}).(InputLinuxParametersResponsePtrOutput)
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





type InputPatchConfigurationResponseInput interface {
	pulumi.Input

	ToInputPatchConfigurationResponseOutput() InputPatchConfigurationResponseOutput
	ToInputPatchConfigurationResponseOutputWithContext(context.Context) InputPatchConfigurationResponseOutput
}

type InputPatchConfigurationResponseArgs struct {
	LinuxParameters   InputLinuxParametersResponsePtrInput   `pulumi:"linuxParameters"`
	PostTasks         TaskPropertiesResponseArrayInput       `pulumi:"postTasks"`
	PreTasks          TaskPropertiesResponseArrayInput       `pulumi:"preTasks"`
	RebootSetting     pulumi.StringPtrInput                  `pulumi:"rebootSetting"`
	WindowsParameters InputWindowsParametersResponsePtrInput `pulumi:"windowsParameters"`
}

func (InputPatchConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputPatchConfigurationResponse)(nil)).Elem()
}

func (i InputPatchConfigurationResponseArgs) ToInputPatchConfigurationResponseOutput() InputPatchConfigurationResponseOutput {
	return i.ToInputPatchConfigurationResponseOutputWithContext(context.Background())
}

func (i InputPatchConfigurationResponseArgs) ToInputPatchConfigurationResponseOutputWithContext(ctx context.Context) InputPatchConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPatchConfigurationResponseOutput)
}

func (i InputPatchConfigurationResponseArgs) ToInputPatchConfigurationResponsePtrOutput() InputPatchConfigurationResponsePtrOutput {
	return i.ToInputPatchConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i InputPatchConfigurationResponseArgs) ToInputPatchConfigurationResponsePtrOutputWithContext(ctx context.Context) InputPatchConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPatchConfigurationResponseOutput).ToInputPatchConfigurationResponsePtrOutputWithContext(ctx)
}









type InputPatchConfigurationResponsePtrInput interface {
	pulumi.Input

	ToInputPatchConfigurationResponsePtrOutput() InputPatchConfigurationResponsePtrOutput
	ToInputPatchConfigurationResponsePtrOutputWithContext(context.Context) InputPatchConfigurationResponsePtrOutput
}

type inputPatchConfigurationResponsePtrType InputPatchConfigurationResponseArgs

func InputPatchConfigurationResponsePtr(v *InputPatchConfigurationResponseArgs) InputPatchConfigurationResponsePtrInput {
	return (*inputPatchConfigurationResponsePtrType)(v)
}

func (*inputPatchConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputPatchConfigurationResponse)(nil)).Elem()
}

func (i *inputPatchConfigurationResponsePtrType) ToInputPatchConfigurationResponsePtrOutput() InputPatchConfigurationResponsePtrOutput {
	return i.ToInputPatchConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *inputPatchConfigurationResponsePtrType) ToInputPatchConfigurationResponsePtrOutputWithContext(ctx context.Context) InputPatchConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputPatchConfigurationResponsePtrOutput)
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

func (o InputPatchConfigurationResponseOutput) ToInputPatchConfigurationResponsePtrOutput() InputPatchConfigurationResponsePtrOutput {
	return o.ToInputPatchConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o InputPatchConfigurationResponseOutput) ToInputPatchConfigurationResponsePtrOutputWithContext(ctx context.Context) InputPatchConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputPatchConfigurationResponse) *InputPatchConfigurationResponse {
		return &v
	}).(InputPatchConfigurationResponsePtrOutput)
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





type InputWindowsParametersResponseInput interface {
	pulumi.Input

	ToInputWindowsParametersResponseOutput() InputWindowsParametersResponseOutput
	ToInputWindowsParametersResponseOutputWithContext(context.Context) InputWindowsParametersResponseOutput
}

type InputWindowsParametersResponseArgs struct {
	ClassificationsToInclude  pulumi.StringArrayInput `pulumi:"classificationsToInclude"`
	ExcludeKbsRequiringReboot pulumi.BoolPtrInput     `pulumi:"excludeKbsRequiringReboot"`
	KbNumbersToExclude        pulumi.StringArrayInput `pulumi:"kbNumbersToExclude"`
	KbNumbersToInclude        pulumi.StringArrayInput `pulumi:"kbNumbersToInclude"`
}

func (InputWindowsParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputWindowsParametersResponse)(nil)).Elem()
}

func (i InputWindowsParametersResponseArgs) ToInputWindowsParametersResponseOutput() InputWindowsParametersResponseOutput {
	return i.ToInputWindowsParametersResponseOutputWithContext(context.Background())
}

func (i InputWindowsParametersResponseArgs) ToInputWindowsParametersResponseOutputWithContext(ctx context.Context) InputWindowsParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputWindowsParametersResponseOutput)
}

func (i InputWindowsParametersResponseArgs) ToInputWindowsParametersResponsePtrOutput() InputWindowsParametersResponsePtrOutput {
	return i.ToInputWindowsParametersResponsePtrOutputWithContext(context.Background())
}

func (i InputWindowsParametersResponseArgs) ToInputWindowsParametersResponsePtrOutputWithContext(ctx context.Context) InputWindowsParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputWindowsParametersResponseOutput).ToInputWindowsParametersResponsePtrOutputWithContext(ctx)
}









type InputWindowsParametersResponsePtrInput interface {
	pulumi.Input

	ToInputWindowsParametersResponsePtrOutput() InputWindowsParametersResponsePtrOutput
	ToInputWindowsParametersResponsePtrOutputWithContext(context.Context) InputWindowsParametersResponsePtrOutput
}

type inputWindowsParametersResponsePtrType InputWindowsParametersResponseArgs

func InputWindowsParametersResponsePtr(v *InputWindowsParametersResponseArgs) InputWindowsParametersResponsePtrInput {
	return (*inputWindowsParametersResponsePtrType)(v)
}

func (*inputWindowsParametersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InputWindowsParametersResponse)(nil)).Elem()
}

func (i *inputWindowsParametersResponsePtrType) ToInputWindowsParametersResponsePtrOutput() InputWindowsParametersResponsePtrOutput {
	return i.ToInputWindowsParametersResponsePtrOutputWithContext(context.Background())
}

func (i *inputWindowsParametersResponsePtrType) ToInputWindowsParametersResponsePtrOutputWithContext(ctx context.Context) InputWindowsParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputWindowsParametersResponsePtrOutput)
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

func (o InputWindowsParametersResponseOutput) ToInputWindowsParametersResponsePtrOutput() InputWindowsParametersResponsePtrOutput {
	return o.ToInputWindowsParametersResponsePtrOutputWithContext(context.Background())
}

func (o InputWindowsParametersResponseOutput) ToInputWindowsParametersResponsePtrOutputWithContext(ctx context.Context) InputWindowsParametersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputWindowsParametersResponse) *InputWindowsParametersResponse {
		return &v
	}).(InputWindowsParametersResponsePtrOutput)
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

type TaskProperties struct {
	Parameters map[string]string `pulumi:"parameters"`
	Source     *string           `pulumi:"source"`
	TaskScope  *string           `pulumi:"taskScope"`
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





type TaskPropertiesResponseInput interface {
	pulumi.Input

	ToTaskPropertiesResponseOutput() TaskPropertiesResponseOutput
	ToTaskPropertiesResponseOutputWithContext(context.Context) TaskPropertiesResponseOutput
}

type TaskPropertiesResponseArgs struct {
	Parameters pulumi.StringMapInput `pulumi:"parameters"`
	Source     pulumi.StringPtrInput `pulumi:"source"`
	TaskScope  pulumi.StringPtrInput `pulumi:"taskScope"`
}

func (TaskPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskPropertiesResponse)(nil)).Elem()
}

func (i TaskPropertiesResponseArgs) ToTaskPropertiesResponseOutput() TaskPropertiesResponseOutput {
	return i.ToTaskPropertiesResponseOutputWithContext(context.Background())
}

func (i TaskPropertiesResponseArgs) ToTaskPropertiesResponseOutputWithContext(ctx context.Context) TaskPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskPropertiesResponseOutput)
}





type TaskPropertiesResponseArrayInput interface {
	pulumi.Input

	ToTaskPropertiesResponseArrayOutput() TaskPropertiesResponseArrayOutput
	ToTaskPropertiesResponseArrayOutputWithContext(context.Context) TaskPropertiesResponseArrayOutput
}

type TaskPropertiesResponseArray []TaskPropertiesResponseInput

func (TaskPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TaskPropertiesResponse)(nil)).Elem()
}

func (i TaskPropertiesResponseArray) ToTaskPropertiesResponseArrayOutput() TaskPropertiesResponseArrayOutput {
	return i.ToTaskPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i TaskPropertiesResponseArray) ToTaskPropertiesResponseArrayOutputWithContext(ctx context.Context) TaskPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskPropertiesResponseArrayOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*InputLinuxParametersInput)(nil)).Elem(), InputLinuxParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputLinuxParametersPtrInput)(nil)).Elem(), InputLinuxParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputLinuxParametersResponseInput)(nil)).Elem(), InputLinuxParametersResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputLinuxParametersResponsePtrInput)(nil)).Elem(), InputLinuxParametersResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputPatchConfigurationInput)(nil)).Elem(), InputPatchConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputPatchConfigurationPtrInput)(nil)).Elem(), InputPatchConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputPatchConfigurationResponseInput)(nil)).Elem(), InputPatchConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputPatchConfigurationResponsePtrInput)(nil)).Elem(), InputPatchConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputWindowsParametersInput)(nil)).Elem(), InputWindowsParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputWindowsParametersPtrInput)(nil)).Elem(), InputWindowsParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputWindowsParametersResponseInput)(nil)).Elem(), InputWindowsParametersResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputWindowsParametersResponsePtrInput)(nil)).Elem(), InputWindowsParametersResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskPropertiesInput)(nil)).Elem(), TaskPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskPropertiesArrayInput)(nil)).Elem(), TaskPropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskPropertiesResponseInput)(nil)).Elem(), TaskPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskPropertiesResponseArrayInput)(nil)).Elem(), TaskPropertiesResponseArray{})
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
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TaskPropertiesOutput{})
	pulumi.RegisterOutputType(TaskPropertiesArrayOutput{})
	pulumi.RegisterOutputType(TaskPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TaskPropertiesResponseArrayOutput{})
}
