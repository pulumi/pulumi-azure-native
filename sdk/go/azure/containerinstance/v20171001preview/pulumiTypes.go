


package v20171001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureFileVolume struct {
	ReadOnly           *bool   `pulumi:"readOnly"`
	ShareName          string  `pulumi:"shareName"`
	StorageAccountKey  *string `pulumi:"storageAccountKey"`
	StorageAccountName string  `pulumi:"storageAccountName"`
}





type AzureFileVolumeInput interface {
	pulumi.Input

	ToAzureFileVolumeOutput() AzureFileVolumeOutput
	ToAzureFileVolumeOutputWithContext(context.Context) AzureFileVolumeOutput
}

type AzureFileVolumeArgs struct {
	ReadOnly           pulumi.BoolPtrInput   `pulumi:"readOnly"`
	ShareName          pulumi.StringInput    `pulumi:"shareName"`
	StorageAccountKey  pulumi.StringPtrInput `pulumi:"storageAccountKey"`
	StorageAccountName pulumi.StringInput    `pulumi:"storageAccountName"`
}

func (AzureFileVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileVolume)(nil)).Elem()
}

func (i AzureFileVolumeArgs) ToAzureFileVolumeOutput() AzureFileVolumeOutput {
	return i.ToAzureFileVolumeOutputWithContext(context.Background())
}

func (i AzureFileVolumeArgs) ToAzureFileVolumeOutputWithContext(ctx context.Context) AzureFileVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileVolumeOutput)
}

func (i AzureFileVolumeArgs) ToAzureFileVolumePtrOutput() AzureFileVolumePtrOutput {
	return i.ToAzureFileVolumePtrOutputWithContext(context.Background())
}

func (i AzureFileVolumeArgs) ToAzureFileVolumePtrOutputWithContext(ctx context.Context) AzureFileVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileVolumeOutput).ToAzureFileVolumePtrOutputWithContext(ctx)
}









type AzureFileVolumePtrInput interface {
	pulumi.Input

	ToAzureFileVolumePtrOutput() AzureFileVolumePtrOutput
	ToAzureFileVolumePtrOutputWithContext(context.Context) AzureFileVolumePtrOutput
}

type azureFileVolumePtrType AzureFileVolumeArgs

func AzureFileVolumePtr(v *AzureFileVolumeArgs) AzureFileVolumePtrInput {
	return (*azureFileVolumePtrType)(v)
}

func (*azureFileVolumePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileVolume)(nil)).Elem()
}

func (i *azureFileVolumePtrType) ToAzureFileVolumePtrOutput() AzureFileVolumePtrOutput {
	return i.ToAzureFileVolumePtrOutputWithContext(context.Background())
}

func (i *azureFileVolumePtrType) ToAzureFileVolumePtrOutputWithContext(ctx context.Context) AzureFileVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileVolumePtrOutput)
}

type AzureFileVolumeOutput struct{ *pulumi.OutputState }

func (AzureFileVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileVolume)(nil)).Elem()
}

func (o AzureFileVolumeOutput) ToAzureFileVolumeOutput() AzureFileVolumeOutput {
	return o
}

func (o AzureFileVolumeOutput) ToAzureFileVolumeOutputWithContext(ctx context.Context) AzureFileVolumeOutput {
	return o
}

func (o AzureFileVolumeOutput) ToAzureFileVolumePtrOutput() AzureFileVolumePtrOutput {
	return o.ToAzureFileVolumePtrOutputWithContext(context.Background())
}

func (o AzureFileVolumeOutput) ToAzureFileVolumePtrOutputWithContext(ctx context.Context) AzureFileVolumePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFileVolume) *AzureFileVolume {
		return &v
	}).(AzureFileVolumePtrOutput)
}

func (o AzureFileVolumeOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFileVolume) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

func (o AzureFileVolumeOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileVolume) string { return v.ShareName }).(pulumi.StringOutput)
}

func (o AzureFileVolumeOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileVolume) *string { return v.StorageAccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureFileVolumeOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileVolume) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

type AzureFileVolumePtrOutput struct{ *pulumi.OutputState }

func (AzureFileVolumePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileVolume)(nil)).Elem()
}

func (o AzureFileVolumePtrOutput) ToAzureFileVolumePtrOutput() AzureFileVolumePtrOutput {
	return o
}

func (o AzureFileVolumePtrOutput) ToAzureFileVolumePtrOutputWithContext(ctx context.Context) AzureFileVolumePtrOutput {
	return o
}

func (o AzureFileVolumePtrOutput) Elem() AzureFileVolumeOutput {
	return o.ApplyT(func(v *AzureFileVolume) AzureFileVolume {
		if v != nil {
			return *v
		}
		var ret AzureFileVolume
		return ret
	}).(AzureFileVolumeOutput)
}

func (o AzureFileVolumePtrOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureFileVolume) *bool {
		if v == nil {
			return nil
		}
		return v.ReadOnly
	}).(pulumi.BoolPtrOutput)
}

func (o AzureFileVolumePtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileVolume) *string {
		if v == nil {
			return nil
		}
		return &v.ShareName
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileVolumePtrOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileVolume) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountKey
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileVolumePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileVolume) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type AzureFileVolumeResponse struct {
	ReadOnly           *bool   `pulumi:"readOnly"`
	ShareName          string  `pulumi:"shareName"`
	StorageAccountKey  *string `pulumi:"storageAccountKey"`
	StorageAccountName string  `pulumi:"storageAccountName"`
}





type AzureFileVolumeResponseInput interface {
	pulumi.Input

	ToAzureFileVolumeResponseOutput() AzureFileVolumeResponseOutput
	ToAzureFileVolumeResponseOutputWithContext(context.Context) AzureFileVolumeResponseOutput
}

type AzureFileVolumeResponseArgs struct {
	ReadOnly           pulumi.BoolPtrInput   `pulumi:"readOnly"`
	ShareName          pulumi.StringInput    `pulumi:"shareName"`
	StorageAccountKey  pulumi.StringPtrInput `pulumi:"storageAccountKey"`
	StorageAccountName pulumi.StringInput    `pulumi:"storageAccountName"`
}

func (AzureFileVolumeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileVolumeResponse)(nil)).Elem()
}

func (i AzureFileVolumeResponseArgs) ToAzureFileVolumeResponseOutput() AzureFileVolumeResponseOutput {
	return i.ToAzureFileVolumeResponseOutputWithContext(context.Background())
}

func (i AzureFileVolumeResponseArgs) ToAzureFileVolumeResponseOutputWithContext(ctx context.Context) AzureFileVolumeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileVolumeResponseOutput)
}

func (i AzureFileVolumeResponseArgs) ToAzureFileVolumeResponsePtrOutput() AzureFileVolumeResponsePtrOutput {
	return i.ToAzureFileVolumeResponsePtrOutputWithContext(context.Background())
}

func (i AzureFileVolumeResponseArgs) ToAzureFileVolumeResponsePtrOutputWithContext(ctx context.Context) AzureFileVolumeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileVolumeResponseOutput).ToAzureFileVolumeResponsePtrOutputWithContext(ctx)
}









type AzureFileVolumeResponsePtrInput interface {
	pulumi.Input

	ToAzureFileVolumeResponsePtrOutput() AzureFileVolumeResponsePtrOutput
	ToAzureFileVolumeResponsePtrOutputWithContext(context.Context) AzureFileVolumeResponsePtrOutput
}

type azureFileVolumeResponsePtrType AzureFileVolumeResponseArgs

func AzureFileVolumeResponsePtr(v *AzureFileVolumeResponseArgs) AzureFileVolumeResponsePtrInput {
	return (*azureFileVolumeResponsePtrType)(v)
}

func (*azureFileVolumeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileVolumeResponse)(nil)).Elem()
}

func (i *azureFileVolumeResponsePtrType) ToAzureFileVolumeResponsePtrOutput() AzureFileVolumeResponsePtrOutput {
	return i.ToAzureFileVolumeResponsePtrOutputWithContext(context.Background())
}

func (i *azureFileVolumeResponsePtrType) ToAzureFileVolumeResponsePtrOutputWithContext(ctx context.Context) AzureFileVolumeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFileVolumeResponsePtrOutput)
}

type AzureFileVolumeResponseOutput struct{ *pulumi.OutputState }

func (AzureFileVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFileVolumeResponse)(nil)).Elem()
}

func (o AzureFileVolumeResponseOutput) ToAzureFileVolumeResponseOutput() AzureFileVolumeResponseOutput {
	return o
}

func (o AzureFileVolumeResponseOutput) ToAzureFileVolumeResponseOutputWithContext(ctx context.Context) AzureFileVolumeResponseOutput {
	return o
}

func (o AzureFileVolumeResponseOutput) ToAzureFileVolumeResponsePtrOutput() AzureFileVolumeResponsePtrOutput {
	return o.ToAzureFileVolumeResponsePtrOutputWithContext(context.Background())
}

func (o AzureFileVolumeResponseOutput) ToAzureFileVolumeResponsePtrOutputWithContext(ctx context.Context) AzureFileVolumeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFileVolumeResponse) *AzureFileVolumeResponse {
		return &v
	}).(AzureFileVolumeResponsePtrOutput)
}

func (o AzureFileVolumeResponseOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFileVolumeResponse) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

func (o AzureFileVolumeResponseOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileVolumeResponse) string { return v.ShareName }).(pulumi.StringOutput)
}

func (o AzureFileVolumeResponseOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFileVolumeResponse) *string { return v.StorageAccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureFileVolumeResponseOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFileVolumeResponse) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

type AzureFileVolumeResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureFileVolumeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFileVolumeResponse)(nil)).Elem()
}

func (o AzureFileVolumeResponsePtrOutput) ToAzureFileVolumeResponsePtrOutput() AzureFileVolumeResponsePtrOutput {
	return o
}

func (o AzureFileVolumeResponsePtrOutput) ToAzureFileVolumeResponsePtrOutputWithContext(ctx context.Context) AzureFileVolumeResponsePtrOutput {
	return o
}

func (o AzureFileVolumeResponsePtrOutput) Elem() AzureFileVolumeResponseOutput {
	return o.ApplyT(func(v *AzureFileVolumeResponse) AzureFileVolumeResponse {
		if v != nil {
			return *v
		}
		var ret AzureFileVolumeResponse
		return ret
	}).(AzureFileVolumeResponseOutput)
}

func (o AzureFileVolumeResponsePtrOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureFileVolumeResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ReadOnly
	}).(pulumi.BoolPtrOutput)
}

func (o AzureFileVolumeResponsePtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ShareName
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileVolumeResponsePtrOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountKey
	}).(pulumi.StringPtrOutput)
}

func (o AzureFileVolumeResponsePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureFileVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type Container struct {
	Command              []string              `pulumi:"command"`
	EnvironmentVariables []EnvironmentVariable `pulumi:"environmentVariables"`
	Image                string                `pulumi:"image"`
	Name                 string                `pulumi:"name"`
	Ports                []ContainerPort       `pulumi:"ports"`
	Resources            ResourceRequirements  `pulumi:"resources"`
	VolumeMounts         []VolumeMount         `pulumi:"volumeMounts"`
}





type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(context.Context) ContainerOutput
}

type ContainerArgs struct {
	Command              pulumi.StringArrayInput       `pulumi:"command"`
	EnvironmentVariables EnvironmentVariableArrayInput `pulumi:"environmentVariables"`
	Image                pulumi.StringInput            `pulumi:"image"`
	Name                 pulumi.StringInput            `pulumi:"name"`
	Ports                ContainerPortArrayInput       `pulumi:"ports"`
	Resources            ResourceRequirementsInput     `pulumi:"resources"`
	VolumeMounts         VolumeMountArrayInput         `pulumi:"volumeMounts"`
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (i ContainerArgs) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i ContainerArgs) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}





type ContainerArrayInput interface {
	pulumi.Input

	ToContainerArrayOutput() ContainerArrayOutput
	ToContainerArrayOutputWithContext(context.Context) ContainerArrayOutput
}

type ContainerArray []ContainerInput

func (ContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Container)(nil)).Elem()
}

func (i ContainerArray) ToContainerArrayOutput() ContainerArrayOutput {
	return i.ToContainerArrayOutputWithContext(context.Background())
}

func (i ContainerArray) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerArrayOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func (o ContainerOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Container) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o ContainerOutput) EnvironmentVariables() EnvironmentVariableArrayOutput {
	return o.ApplyT(func(v Container) []EnvironmentVariable { return v.EnvironmentVariables }).(EnvironmentVariableArrayOutput)
}

func (o ContainerOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v Container) string { return v.Image }).(pulumi.StringOutput)
}

func (o ContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Container) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerOutput) Ports() ContainerPortArrayOutput {
	return o.ApplyT(func(v Container) []ContainerPort { return v.Ports }).(ContainerPortArrayOutput)
}

func (o ContainerOutput) Resources() ResourceRequirementsOutput {
	return o.ApplyT(func(v Container) ResourceRequirements { return v.Resources }).(ResourceRequirementsOutput)
}

func (o ContainerOutput) VolumeMounts() VolumeMountArrayOutput {
	return o.ApplyT(func(v Container) []VolumeMount { return v.VolumeMounts }).(VolumeMountArrayOutput)
}

type ContainerArrayOutput struct{ *pulumi.OutputState }

func (ContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Container)(nil)).Elem()
}

func (o ContainerArrayOutput) ToContainerArrayOutput() ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) Index(i pulumi.IntInput) ContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Container {
		return vs[0].([]Container)[vs[1].(int)]
	}).(ContainerOutput)
}

type ContainerGroupResponseInstanceView struct {
	Events []EventResponse `pulumi:"events"`
	State  string          `pulumi:"state"`
}





type ContainerGroupResponseInstanceViewInput interface {
	pulumi.Input

	ToContainerGroupResponseInstanceViewOutput() ContainerGroupResponseInstanceViewOutput
	ToContainerGroupResponseInstanceViewOutputWithContext(context.Context) ContainerGroupResponseInstanceViewOutput
}

type ContainerGroupResponseInstanceViewArgs struct {
	Events EventResponseArrayInput `pulumi:"events"`
	State  pulumi.StringInput      `pulumi:"state"`
}

func (ContainerGroupResponseInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupResponseInstanceView)(nil)).Elem()
}

func (i ContainerGroupResponseInstanceViewArgs) ToContainerGroupResponseInstanceViewOutput() ContainerGroupResponseInstanceViewOutput {
	return i.ToContainerGroupResponseInstanceViewOutputWithContext(context.Background())
}

func (i ContainerGroupResponseInstanceViewArgs) ToContainerGroupResponseInstanceViewOutputWithContext(ctx context.Context) ContainerGroupResponseInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupResponseInstanceViewOutput)
}

func (i ContainerGroupResponseInstanceViewArgs) ToContainerGroupResponseInstanceViewPtrOutput() ContainerGroupResponseInstanceViewPtrOutput {
	return i.ToContainerGroupResponseInstanceViewPtrOutputWithContext(context.Background())
}

func (i ContainerGroupResponseInstanceViewArgs) ToContainerGroupResponseInstanceViewPtrOutputWithContext(ctx context.Context) ContainerGroupResponseInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupResponseInstanceViewOutput).ToContainerGroupResponseInstanceViewPtrOutputWithContext(ctx)
}









type ContainerGroupResponseInstanceViewPtrInput interface {
	pulumi.Input

	ToContainerGroupResponseInstanceViewPtrOutput() ContainerGroupResponseInstanceViewPtrOutput
	ToContainerGroupResponseInstanceViewPtrOutputWithContext(context.Context) ContainerGroupResponseInstanceViewPtrOutput
}

type containerGroupResponseInstanceViewPtrType ContainerGroupResponseInstanceViewArgs

func ContainerGroupResponseInstanceViewPtr(v *ContainerGroupResponseInstanceViewArgs) ContainerGroupResponseInstanceViewPtrInput {
	return (*containerGroupResponseInstanceViewPtrType)(v)
}

func (*containerGroupResponseInstanceViewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupResponseInstanceView)(nil)).Elem()
}

func (i *containerGroupResponseInstanceViewPtrType) ToContainerGroupResponseInstanceViewPtrOutput() ContainerGroupResponseInstanceViewPtrOutput {
	return i.ToContainerGroupResponseInstanceViewPtrOutputWithContext(context.Background())
}

func (i *containerGroupResponseInstanceViewPtrType) ToContainerGroupResponseInstanceViewPtrOutputWithContext(ctx context.Context) ContainerGroupResponseInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupResponseInstanceViewPtrOutput)
}

type ContainerGroupResponseInstanceViewOutput struct{ *pulumi.OutputState }

func (ContainerGroupResponseInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupResponseInstanceView)(nil)).Elem()
}

func (o ContainerGroupResponseInstanceViewOutput) ToContainerGroupResponseInstanceViewOutput() ContainerGroupResponseInstanceViewOutput {
	return o
}

func (o ContainerGroupResponseInstanceViewOutput) ToContainerGroupResponseInstanceViewOutputWithContext(ctx context.Context) ContainerGroupResponseInstanceViewOutput {
	return o
}

func (o ContainerGroupResponseInstanceViewOutput) ToContainerGroupResponseInstanceViewPtrOutput() ContainerGroupResponseInstanceViewPtrOutput {
	return o.ToContainerGroupResponseInstanceViewPtrOutputWithContext(context.Background())
}

func (o ContainerGroupResponseInstanceViewOutput) ToContainerGroupResponseInstanceViewPtrOutputWithContext(ctx context.Context) ContainerGroupResponseInstanceViewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroupResponseInstanceView) *ContainerGroupResponseInstanceView {
		return &v
	}).(ContainerGroupResponseInstanceViewPtrOutput)
}

func (o ContainerGroupResponseInstanceViewOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v ContainerGroupResponseInstanceView) []EventResponse { return v.Events }).(EventResponseArrayOutput)
}

func (o ContainerGroupResponseInstanceViewOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerGroupResponseInstanceView) string { return v.State }).(pulumi.StringOutput)
}

type ContainerGroupResponseInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupResponseInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupResponseInstanceView)(nil)).Elem()
}

func (o ContainerGroupResponseInstanceViewPtrOutput) ToContainerGroupResponseInstanceViewPtrOutput() ContainerGroupResponseInstanceViewPtrOutput {
	return o
}

func (o ContainerGroupResponseInstanceViewPtrOutput) ToContainerGroupResponseInstanceViewPtrOutputWithContext(ctx context.Context) ContainerGroupResponseInstanceViewPtrOutput {
	return o
}

func (o ContainerGroupResponseInstanceViewPtrOutput) Elem() ContainerGroupResponseInstanceViewOutput {
	return o.ApplyT(func(v *ContainerGroupResponseInstanceView) ContainerGroupResponseInstanceView {
		if v != nil {
			return *v
		}
		var ret ContainerGroupResponseInstanceView
		return ret
	}).(ContainerGroupResponseInstanceViewOutput)
}

func (o ContainerGroupResponseInstanceViewPtrOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroupResponseInstanceView) []EventResponse {
		if v == nil {
			return nil
		}
		return v.Events
	}).(EventResponseArrayOutput)
}

func (o ContainerGroupResponseInstanceViewPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroupResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type ContainerPort struct {
	Port     int     `pulumi:"port"`
	Protocol *string `pulumi:"protocol"`
}





type ContainerPortInput interface {
	pulumi.Input

	ToContainerPortOutput() ContainerPortOutput
	ToContainerPortOutputWithContext(context.Context) ContainerPortOutput
}

type ContainerPortArgs struct {
	Port     pulumi.IntInput       `pulumi:"port"`
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
}

func (ContainerPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPort)(nil)).Elem()
}

func (i ContainerPortArgs) ToContainerPortOutput() ContainerPortOutput {
	return i.ToContainerPortOutputWithContext(context.Background())
}

func (i ContainerPortArgs) ToContainerPortOutputWithContext(ctx context.Context) ContainerPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPortOutput)
}





type ContainerPortArrayInput interface {
	pulumi.Input

	ToContainerPortArrayOutput() ContainerPortArrayOutput
	ToContainerPortArrayOutputWithContext(context.Context) ContainerPortArrayOutput
}

type ContainerPortArray []ContainerPortInput

func (ContainerPortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerPort)(nil)).Elem()
}

func (i ContainerPortArray) ToContainerPortArrayOutput() ContainerPortArrayOutput {
	return i.ToContainerPortArrayOutputWithContext(context.Background())
}

func (i ContainerPortArray) ToContainerPortArrayOutputWithContext(ctx context.Context) ContainerPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPortArrayOutput)
}

type ContainerPortOutput struct{ *pulumi.OutputState }

func (ContainerPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPort)(nil)).Elem()
}

func (o ContainerPortOutput) ToContainerPortOutput() ContainerPortOutput {
	return o
}

func (o ContainerPortOutput) ToContainerPortOutputWithContext(ctx context.Context) ContainerPortOutput {
	return o
}

func (o ContainerPortOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerPort) int { return v.Port }).(pulumi.IntOutput)
}

func (o ContainerPortOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerPort) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type ContainerPortArrayOutput struct{ *pulumi.OutputState }

func (ContainerPortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerPort)(nil)).Elem()
}

func (o ContainerPortArrayOutput) ToContainerPortArrayOutput() ContainerPortArrayOutput {
	return o
}

func (o ContainerPortArrayOutput) ToContainerPortArrayOutputWithContext(ctx context.Context) ContainerPortArrayOutput {
	return o
}

func (o ContainerPortArrayOutput) Index(i pulumi.IntInput) ContainerPortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerPort {
		return vs[0].([]ContainerPort)[vs[1].(int)]
	}).(ContainerPortOutput)
}

type ContainerPortResponse struct {
	Port     int     `pulumi:"port"`
	Protocol *string `pulumi:"protocol"`
}





type ContainerPortResponseInput interface {
	pulumi.Input

	ToContainerPortResponseOutput() ContainerPortResponseOutput
	ToContainerPortResponseOutputWithContext(context.Context) ContainerPortResponseOutput
}

type ContainerPortResponseArgs struct {
	Port     pulumi.IntInput       `pulumi:"port"`
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
}

func (ContainerPortResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPortResponse)(nil)).Elem()
}

func (i ContainerPortResponseArgs) ToContainerPortResponseOutput() ContainerPortResponseOutput {
	return i.ToContainerPortResponseOutputWithContext(context.Background())
}

func (i ContainerPortResponseArgs) ToContainerPortResponseOutputWithContext(ctx context.Context) ContainerPortResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPortResponseOutput)
}





type ContainerPortResponseArrayInput interface {
	pulumi.Input

	ToContainerPortResponseArrayOutput() ContainerPortResponseArrayOutput
	ToContainerPortResponseArrayOutputWithContext(context.Context) ContainerPortResponseArrayOutput
}

type ContainerPortResponseArray []ContainerPortResponseInput

func (ContainerPortResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerPortResponse)(nil)).Elem()
}

func (i ContainerPortResponseArray) ToContainerPortResponseArrayOutput() ContainerPortResponseArrayOutput {
	return i.ToContainerPortResponseArrayOutputWithContext(context.Background())
}

func (i ContainerPortResponseArray) ToContainerPortResponseArrayOutputWithContext(ctx context.Context) ContainerPortResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPortResponseArrayOutput)
}

type ContainerPortResponseOutput struct{ *pulumi.OutputState }

func (ContainerPortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPortResponse)(nil)).Elem()
}

func (o ContainerPortResponseOutput) ToContainerPortResponseOutput() ContainerPortResponseOutput {
	return o
}

func (o ContainerPortResponseOutput) ToContainerPortResponseOutputWithContext(ctx context.Context) ContainerPortResponseOutput {
	return o
}

func (o ContainerPortResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerPortResponse) int { return v.Port }).(pulumi.IntOutput)
}

func (o ContainerPortResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerPortResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type ContainerPortResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerPortResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerPortResponse)(nil)).Elem()
}

func (o ContainerPortResponseArrayOutput) ToContainerPortResponseArrayOutput() ContainerPortResponseArrayOutput {
	return o
}

func (o ContainerPortResponseArrayOutput) ToContainerPortResponseArrayOutputWithContext(ctx context.Context) ContainerPortResponseArrayOutput {
	return o
}

func (o ContainerPortResponseArrayOutput) Index(i pulumi.IntInput) ContainerPortResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerPortResponse {
		return vs[0].([]ContainerPortResponse)[vs[1].(int)]
	}).(ContainerPortResponseOutput)
}

type ContainerPropertiesResponseInstanceView struct {
	CurrentState  ContainerStateResponse `pulumi:"currentState"`
	Events        []EventResponse        `pulumi:"events"`
	PreviousState ContainerStateResponse `pulumi:"previousState"`
	RestartCount  int                    `pulumi:"restartCount"`
}





type ContainerPropertiesResponseInstanceViewInput interface {
	pulumi.Input

	ToContainerPropertiesResponseInstanceViewOutput() ContainerPropertiesResponseInstanceViewOutput
	ToContainerPropertiesResponseInstanceViewOutputWithContext(context.Context) ContainerPropertiesResponseInstanceViewOutput
}

type ContainerPropertiesResponseInstanceViewArgs struct {
	CurrentState  ContainerStateResponseInput `pulumi:"currentState"`
	Events        EventResponseArrayInput     `pulumi:"events"`
	PreviousState ContainerStateResponseInput `pulumi:"previousState"`
	RestartCount  pulumi.IntInput             `pulumi:"restartCount"`
}

func (ContainerPropertiesResponseInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPropertiesResponseInstanceView)(nil)).Elem()
}

func (i ContainerPropertiesResponseInstanceViewArgs) ToContainerPropertiesResponseInstanceViewOutput() ContainerPropertiesResponseInstanceViewOutput {
	return i.ToContainerPropertiesResponseInstanceViewOutputWithContext(context.Background())
}

func (i ContainerPropertiesResponseInstanceViewArgs) ToContainerPropertiesResponseInstanceViewOutputWithContext(ctx context.Context) ContainerPropertiesResponseInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPropertiesResponseInstanceViewOutput)
}

type ContainerPropertiesResponseInstanceViewOutput struct{ *pulumi.OutputState }

func (ContainerPropertiesResponseInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPropertiesResponseInstanceView)(nil)).Elem()
}

func (o ContainerPropertiesResponseInstanceViewOutput) ToContainerPropertiesResponseInstanceViewOutput() ContainerPropertiesResponseInstanceViewOutput {
	return o
}

func (o ContainerPropertiesResponseInstanceViewOutput) ToContainerPropertiesResponseInstanceViewOutputWithContext(ctx context.Context) ContainerPropertiesResponseInstanceViewOutput {
	return o
}

func (o ContainerPropertiesResponseInstanceViewOutput) CurrentState() ContainerStateResponseOutput {
	return o.ApplyT(func(v ContainerPropertiesResponseInstanceView) ContainerStateResponse { return v.CurrentState }).(ContainerStateResponseOutput)
}

func (o ContainerPropertiesResponseInstanceViewOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v ContainerPropertiesResponseInstanceView) []EventResponse { return v.Events }).(EventResponseArrayOutput)
}

func (o ContainerPropertiesResponseInstanceViewOutput) PreviousState() ContainerStateResponseOutput {
	return o.ApplyT(func(v ContainerPropertiesResponseInstanceView) ContainerStateResponse { return v.PreviousState }).(ContainerStateResponseOutput)
}

func (o ContainerPropertiesResponseInstanceViewOutput) RestartCount() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerPropertiesResponseInstanceView) int { return v.RestartCount }).(pulumi.IntOutput)
}

type ContainerResponse struct {
	Command              []string                                `pulumi:"command"`
	EnvironmentVariables []EnvironmentVariableResponse           `pulumi:"environmentVariables"`
	Image                string                                  `pulumi:"image"`
	InstanceView         ContainerPropertiesResponseInstanceView `pulumi:"instanceView"`
	Name                 string                                  `pulumi:"name"`
	Ports                []ContainerPortResponse                 `pulumi:"ports"`
	Resources            ResourceRequirementsResponse            `pulumi:"resources"`
	VolumeMounts         []VolumeMountResponse                   `pulumi:"volumeMounts"`
}





type ContainerResponseInput interface {
	pulumi.Input

	ToContainerResponseOutput() ContainerResponseOutput
	ToContainerResponseOutputWithContext(context.Context) ContainerResponseOutput
}

type ContainerResponseArgs struct {
	Command              pulumi.StringArrayInput                      `pulumi:"command"`
	EnvironmentVariables EnvironmentVariableResponseArrayInput        `pulumi:"environmentVariables"`
	Image                pulumi.StringInput                           `pulumi:"image"`
	InstanceView         ContainerPropertiesResponseInstanceViewInput `pulumi:"instanceView"`
	Name                 pulumi.StringInput                           `pulumi:"name"`
	Ports                ContainerPortResponseArrayInput              `pulumi:"ports"`
	Resources            ResourceRequirementsResponseInput            `pulumi:"resources"`
	VolumeMounts         VolumeMountResponseArrayInput                `pulumi:"volumeMounts"`
}

func (ContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResponse)(nil)).Elem()
}

func (i ContainerResponseArgs) ToContainerResponseOutput() ContainerResponseOutput {
	return i.ToContainerResponseOutputWithContext(context.Background())
}

func (i ContainerResponseArgs) ToContainerResponseOutputWithContext(ctx context.Context) ContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResponseOutput)
}





type ContainerResponseArrayInput interface {
	pulumi.Input

	ToContainerResponseArrayOutput() ContainerResponseArrayOutput
	ToContainerResponseArrayOutputWithContext(context.Context) ContainerResponseArrayOutput
}

type ContainerResponseArray []ContainerResponseInput

func (ContainerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerResponse)(nil)).Elem()
}

func (i ContainerResponseArray) ToContainerResponseArrayOutput() ContainerResponseArrayOutput {
	return i.ToContainerResponseArrayOutputWithContext(context.Background())
}

func (i ContainerResponseArray) ToContainerResponseArrayOutputWithContext(ctx context.Context) ContainerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResponseArrayOutput)
}

type ContainerResponseOutput struct{ *pulumi.OutputState }

func (ContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResponse)(nil)).Elem()
}

func (o ContainerResponseOutput) ToContainerResponseOutput() ContainerResponseOutput {
	return o
}

func (o ContainerResponseOutput) ToContainerResponseOutputWithContext(ctx context.Context) ContainerResponseOutput {
	return o
}

func (o ContainerResponseOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o ContainerResponseOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []EnvironmentVariableResponse { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

func (o ContainerResponseOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerResponse) string { return v.Image }).(pulumi.StringOutput)
}

func (o ContainerResponseOutput) InstanceView() ContainerPropertiesResponseInstanceViewOutput {
	return o.ApplyT(func(v ContainerResponse) ContainerPropertiesResponseInstanceView { return v.InstanceView }).(ContainerPropertiesResponseInstanceViewOutput)
}

func (o ContainerResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerResponseOutput) Ports() ContainerPortResponseArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []ContainerPortResponse { return v.Ports }).(ContainerPortResponseArrayOutput)
}

func (o ContainerResponseOutput) Resources() ResourceRequirementsResponseOutput {
	return o.ApplyT(func(v ContainerResponse) ResourceRequirementsResponse { return v.Resources }).(ResourceRequirementsResponseOutput)
}

func (o ContainerResponseOutput) VolumeMounts() VolumeMountResponseArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []VolumeMountResponse { return v.VolumeMounts }).(VolumeMountResponseArrayOutput)
}

type ContainerResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerResponse)(nil)).Elem()
}

func (o ContainerResponseArrayOutput) ToContainerResponseArrayOutput() ContainerResponseArrayOutput {
	return o
}

func (o ContainerResponseArrayOutput) ToContainerResponseArrayOutputWithContext(ctx context.Context) ContainerResponseArrayOutput {
	return o
}

func (o ContainerResponseArrayOutput) Index(i pulumi.IntInput) ContainerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerResponse {
		return vs[0].([]ContainerResponse)[vs[1].(int)]
	}).(ContainerResponseOutput)
}

type ContainerStateResponse struct {
	DetailStatus *string `pulumi:"detailStatus"`
	ExitCode     *int    `pulumi:"exitCode"`
	FinishTime   *string `pulumi:"finishTime"`
	StartTime    *string `pulumi:"startTime"`
	State        *string `pulumi:"state"`
}





type ContainerStateResponseInput interface {
	pulumi.Input

	ToContainerStateResponseOutput() ContainerStateResponseOutput
	ToContainerStateResponseOutputWithContext(context.Context) ContainerStateResponseOutput
}

type ContainerStateResponseArgs struct {
	DetailStatus pulumi.StringPtrInput `pulumi:"detailStatus"`
	ExitCode     pulumi.IntPtrInput    `pulumi:"exitCode"`
	FinishTime   pulumi.StringPtrInput `pulumi:"finishTime"`
	StartTime    pulumi.StringPtrInput `pulumi:"startTime"`
	State        pulumi.StringPtrInput `pulumi:"state"`
}

func (ContainerStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerStateResponse)(nil)).Elem()
}

func (i ContainerStateResponseArgs) ToContainerStateResponseOutput() ContainerStateResponseOutput {
	return i.ToContainerStateResponseOutputWithContext(context.Background())
}

func (i ContainerStateResponseArgs) ToContainerStateResponseOutputWithContext(ctx context.Context) ContainerStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerStateResponseOutput)
}

type ContainerStateResponseOutput struct{ *pulumi.OutputState }

func (ContainerStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerStateResponse)(nil)).Elem()
}

func (o ContainerStateResponseOutput) ToContainerStateResponseOutput() ContainerStateResponseOutput {
	return o
}

func (o ContainerStateResponseOutput) ToContainerStateResponseOutputWithContext(ctx context.Context) ContainerStateResponseOutput {
	return o
}

func (o ContainerStateResponseOutput) DetailStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.DetailStatus }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) ExitCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *int { return v.ExitCode }).(pulumi.IntPtrOutput)
}

func (o ContainerStateResponseOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.FinishTime }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type EnvironmentVariable struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type EnvironmentVariableInput interface {
	pulumi.Input

	ToEnvironmentVariableOutput() EnvironmentVariableOutput
	ToEnvironmentVariableOutputWithContext(context.Context) EnvironmentVariableOutput
}

type EnvironmentVariableArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (EnvironmentVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariable)(nil)).Elem()
}

func (i EnvironmentVariableArgs) ToEnvironmentVariableOutput() EnvironmentVariableOutput {
	return i.ToEnvironmentVariableOutputWithContext(context.Background())
}

func (i EnvironmentVariableArgs) ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableOutput)
}





type EnvironmentVariableArrayInput interface {
	pulumi.Input

	ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput
	ToEnvironmentVariableArrayOutputWithContext(context.Context) EnvironmentVariableArrayOutput
}

type EnvironmentVariableArray []EnvironmentVariableInput

func (EnvironmentVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariable)(nil)).Elem()
}

func (i EnvironmentVariableArray) ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput {
	return i.ToEnvironmentVariableArrayOutputWithContext(context.Background())
}

func (i EnvironmentVariableArray) ToEnvironmentVariableArrayOutputWithContext(ctx context.Context) EnvironmentVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableArrayOutput)
}

type EnvironmentVariableOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableOutput) ToEnvironmentVariableOutput() EnvironmentVariableOutput {
	return o
}

func (o EnvironmentVariableOutput) ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput {
	return o
}

func (o EnvironmentVariableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentVariable) string { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentVariableOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentVariable) string { return v.Value }).(pulumi.StringOutput)
}

type EnvironmentVariableArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableArrayOutput) ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput {
	return o
}

func (o EnvironmentVariableArrayOutput) ToEnvironmentVariableArrayOutputWithContext(ctx context.Context) EnvironmentVariableArrayOutput {
	return o
}

func (o EnvironmentVariableArrayOutput) Index(i pulumi.IntInput) EnvironmentVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVariable {
		return vs[0].([]EnvironmentVariable)[vs[1].(int)]
	}).(EnvironmentVariableOutput)
}

type EnvironmentVariableResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type EnvironmentVariableResponseInput interface {
	pulumi.Input

	ToEnvironmentVariableResponseOutput() EnvironmentVariableResponseOutput
	ToEnvironmentVariableResponseOutputWithContext(context.Context) EnvironmentVariableResponseOutput
}

type EnvironmentVariableResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (EnvironmentVariableResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariableResponse)(nil)).Elem()
}

func (i EnvironmentVariableResponseArgs) ToEnvironmentVariableResponseOutput() EnvironmentVariableResponseOutput {
	return i.ToEnvironmentVariableResponseOutputWithContext(context.Background())
}

func (i EnvironmentVariableResponseArgs) ToEnvironmentVariableResponseOutputWithContext(ctx context.Context) EnvironmentVariableResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableResponseOutput)
}





type EnvironmentVariableResponseArrayInput interface {
	pulumi.Input

	ToEnvironmentVariableResponseArrayOutput() EnvironmentVariableResponseArrayOutput
	ToEnvironmentVariableResponseArrayOutputWithContext(context.Context) EnvironmentVariableResponseArrayOutput
}

type EnvironmentVariableResponseArray []EnvironmentVariableResponseInput

func (EnvironmentVariableResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariableResponse)(nil)).Elem()
}

func (i EnvironmentVariableResponseArray) ToEnvironmentVariableResponseArrayOutput() EnvironmentVariableResponseArrayOutput {
	return i.ToEnvironmentVariableResponseArrayOutputWithContext(context.Background())
}

func (i EnvironmentVariableResponseArray) ToEnvironmentVariableResponseArrayOutputWithContext(ctx context.Context) EnvironmentVariableResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableResponseArrayOutput)
}

type EnvironmentVariableResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariableResponse)(nil)).Elem()
}

func (o EnvironmentVariableResponseOutput) ToEnvironmentVariableResponseOutput() EnvironmentVariableResponseOutput {
	return o
}

func (o EnvironmentVariableResponseOutput) ToEnvironmentVariableResponseOutputWithContext(ctx context.Context) EnvironmentVariableResponseOutput {
	return o
}

func (o EnvironmentVariableResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentVariableResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentVariableResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v EnvironmentVariableResponse) string { return v.Value }).(pulumi.StringOutput)
}

type EnvironmentVariableResponseArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariableResponse)(nil)).Elem()
}

func (o EnvironmentVariableResponseArrayOutput) ToEnvironmentVariableResponseArrayOutput() EnvironmentVariableResponseArrayOutput {
	return o
}

func (o EnvironmentVariableResponseArrayOutput) ToEnvironmentVariableResponseArrayOutputWithContext(ctx context.Context) EnvironmentVariableResponseArrayOutput {
	return o
}

func (o EnvironmentVariableResponseArrayOutput) Index(i pulumi.IntInput) EnvironmentVariableResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVariableResponse {
		return vs[0].([]EnvironmentVariableResponse)[vs[1].(int)]
	}).(EnvironmentVariableResponseOutput)
}

type EventResponse struct {
	Count          *int    `pulumi:"count"`
	FirstTimestamp *string `pulumi:"firstTimestamp"`
	LastTimestamp  *string `pulumi:"lastTimestamp"`
	Message        *string `pulumi:"message"`
	Name           *string `pulumi:"name"`
	Type           *string `pulumi:"type"`
}





type EventResponseInput interface {
	pulumi.Input

	ToEventResponseOutput() EventResponseOutput
	ToEventResponseOutputWithContext(context.Context) EventResponseOutput
}

type EventResponseArgs struct {
	Count          pulumi.IntPtrInput    `pulumi:"count"`
	FirstTimestamp pulumi.StringPtrInput `pulumi:"firstTimestamp"`
	LastTimestamp  pulumi.StringPtrInput `pulumi:"lastTimestamp"`
	Message        pulumi.StringPtrInput `pulumi:"message"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	Type           pulumi.StringPtrInput `pulumi:"type"`
}

func (EventResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventResponse)(nil)).Elem()
}

func (i EventResponseArgs) ToEventResponseOutput() EventResponseOutput {
	return i.ToEventResponseOutputWithContext(context.Background())
}

func (i EventResponseArgs) ToEventResponseOutputWithContext(ctx context.Context) EventResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventResponseOutput)
}





type EventResponseArrayInput interface {
	pulumi.Input

	ToEventResponseArrayOutput() EventResponseArrayOutput
	ToEventResponseArrayOutputWithContext(context.Context) EventResponseArrayOutput
}

type EventResponseArray []EventResponseInput

func (EventResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventResponse)(nil)).Elem()
}

func (i EventResponseArray) ToEventResponseArrayOutput() EventResponseArrayOutput {
	return i.ToEventResponseArrayOutputWithContext(context.Background())
}

func (i EventResponseArray) ToEventResponseArrayOutputWithContext(ctx context.Context) EventResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventResponseArrayOutput)
}

type EventResponseOutput struct{ *pulumi.OutputState }

func (EventResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventResponse)(nil)).Elem()
}

func (o EventResponseOutput) ToEventResponseOutput() EventResponseOutput {
	return o
}

func (o EventResponseOutput) ToEventResponseOutputWithContext(ctx context.Context) EventResponseOutput {
	return o
}

func (o EventResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EventResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o EventResponseOutput) FirstTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponse) *string { return v.FirstTimestamp }).(pulumi.StringPtrOutput)
}

func (o EventResponseOutput) LastTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponse) *string { return v.LastTimestamp }).(pulumi.StringPtrOutput)
}

func (o EventResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o EventResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EventResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type EventResponseArrayOutput struct{ *pulumi.OutputState }

func (EventResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventResponse)(nil)).Elem()
}

func (o EventResponseArrayOutput) ToEventResponseArrayOutput() EventResponseArrayOutput {
	return o
}

func (o EventResponseArrayOutput) ToEventResponseArrayOutputWithContext(ctx context.Context) EventResponseArrayOutput {
	return o
}

func (o EventResponseArrayOutput) Index(i pulumi.IntInput) EventResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventResponse {
		return vs[0].([]EventResponse)[vs[1].(int)]
	}).(EventResponseOutput)
}

type ImageRegistryCredential struct {
	Password *string `pulumi:"password"`
	Server   string  `pulumi:"server"`
	Username string  `pulumi:"username"`
}





type ImageRegistryCredentialInput interface {
	pulumi.Input

	ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput
	ToImageRegistryCredentialOutputWithContext(context.Context) ImageRegistryCredentialOutput
}

type ImageRegistryCredentialArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Server   pulumi.StringInput    `pulumi:"server"`
	Username pulumi.StringInput    `pulumi:"username"`
}

func (ImageRegistryCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredential)(nil)).Elem()
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput {
	return i.ToImageRegistryCredentialOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialOutputWithContext(ctx context.Context) ImageRegistryCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialOutput)
}





type ImageRegistryCredentialArrayInput interface {
	pulumi.Input

	ToImageRegistryCredentialArrayOutput() ImageRegistryCredentialArrayOutput
	ToImageRegistryCredentialArrayOutputWithContext(context.Context) ImageRegistryCredentialArrayOutput
}

type ImageRegistryCredentialArray []ImageRegistryCredentialInput

func (ImageRegistryCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageRegistryCredential)(nil)).Elem()
}

func (i ImageRegistryCredentialArray) ToImageRegistryCredentialArrayOutput() ImageRegistryCredentialArrayOutput {
	return i.ToImageRegistryCredentialArrayOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialArray) ToImageRegistryCredentialArrayOutputWithContext(ctx context.Context) ImageRegistryCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialArrayOutput)
}

type ImageRegistryCredentialOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredential)(nil)).Elem()
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput {
	return o
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialOutputWithContext(ctx context.Context) ImageRegistryCredentialOutput {
	return o
}

func (o ImageRegistryCredentialOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredential) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredential) string { return v.Server }).(pulumi.StringOutput)
}

func (o ImageRegistryCredentialOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredential) string { return v.Username }).(pulumi.StringOutput)
}

type ImageRegistryCredentialArrayOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageRegistryCredential)(nil)).Elem()
}

func (o ImageRegistryCredentialArrayOutput) ToImageRegistryCredentialArrayOutput() ImageRegistryCredentialArrayOutput {
	return o
}

func (o ImageRegistryCredentialArrayOutput) ToImageRegistryCredentialArrayOutputWithContext(ctx context.Context) ImageRegistryCredentialArrayOutput {
	return o
}

func (o ImageRegistryCredentialArrayOutput) Index(i pulumi.IntInput) ImageRegistryCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageRegistryCredential {
		return vs[0].([]ImageRegistryCredential)[vs[1].(int)]
	}).(ImageRegistryCredentialOutput)
}

type ImageRegistryCredentialResponse struct {
	Password *string `pulumi:"password"`
	Server   string  `pulumi:"server"`
	Username string  `pulumi:"username"`
}





type ImageRegistryCredentialResponseInput interface {
	pulumi.Input

	ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput
	ToImageRegistryCredentialResponseOutputWithContext(context.Context) ImageRegistryCredentialResponseOutput
}

type ImageRegistryCredentialResponseArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Server   pulumi.StringInput    `pulumi:"server"`
	Username pulumi.StringInput    `pulumi:"username"`
}

func (ImageRegistryCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredentialResponse)(nil)).Elem()
}

func (i ImageRegistryCredentialResponseArgs) ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput {
	return i.ToImageRegistryCredentialResponseOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialResponseArgs) ToImageRegistryCredentialResponseOutputWithContext(ctx context.Context) ImageRegistryCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialResponseOutput)
}





type ImageRegistryCredentialResponseArrayInput interface {
	pulumi.Input

	ToImageRegistryCredentialResponseArrayOutput() ImageRegistryCredentialResponseArrayOutput
	ToImageRegistryCredentialResponseArrayOutputWithContext(context.Context) ImageRegistryCredentialResponseArrayOutput
}

type ImageRegistryCredentialResponseArray []ImageRegistryCredentialResponseInput

func (ImageRegistryCredentialResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageRegistryCredentialResponse)(nil)).Elem()
}

func (i ImageRegistryCredentialResponseArray) ToImageRegistryCredentialResponseArrayOutput() ImageRegistryCredentialResponseArrayOutput {
	return i.ToImageRegistryCredentialResponseArrayOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialResponseArray) ToImageRegistryCredentialResponseArrayOutputWithContext(ctx context.Context) ImageRegistryCredentialResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialResponseArrayOutput)
}

type ImageRegistryCredentialResponseOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredentialResponse)(nil)).Elem()
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput {
	return o
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponseOutputWithContext(ctx context.Context) ImageRegistryCredentialResponseOutput {
	return o
}

func (o ImageRegistryCredentialResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponseOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) string { return v.Server }).(pulumi.StringOutput)
}

func (o ImageRegistryCredentialResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) string { return v.Username }).(pulumi.StringOutput)
}

type ImageRegistryCredentialResponseArrayOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageRegistryCredentialResponse)(nil)).Elem()
}

func (o ImageRegistryCredentialResponseArrayOutput) ToImageRegistryCredentialResponseArrayOutput() ImageRegistryCredentialResponseArrayOutput {
	return o
}

func (o ImageRegistryCredentialResponseArrayOutput) ToImageRegistryCredentialResponseArrayOutputWithContext(ctx context.Context) ImageRegistryCredentialResponseArrayOutput {
	return o
}

func (o ImageRegistryCredentialResponseArrayOutput) Index(i pulumi.IntInput) ImageRegistryCredentialResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageRegistryCredentialResponse {
		return vs[0].([]ImageRegistryCredentialResponse)[vs[1].(int)]
	}).(ImageRegistryCredentialResponseOutput)
}

type IpAddress struct {
	Ip    *string `pulumi:"ip"`
	Ports []Port  `pulumi:"ports"`
	Type  string  `pulumi:"type"`
}





type IpAddressInput interface {
	pulumi.Input

	ToIpAddressOutput() IpAddressOutput
	ToIpAddressOutputWithContext(context.Context) IpAddressOutput
}

type IpAddressArgs struct {
	Ip    pulumi.StringPtrInput `pulumi:"ip"`
	Ports PortArrayInput        `pulumi:"ports"`
	Type  pulumi.StringInput    `pulumi:"type"`
}

func (IpAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddress)(nil)).Elem()
}

func (i IpAddressArgs) ToIpAddressOutput() IpAddressOutput {
	return i.ToIpAddressOutputWithContext(context.Background())
}

func (i IpAddressArgs) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOutput)
}

func (i IpAddressArgs) ToIpAddressPtrOutput() IpAddressPtrOutput {
	return i.ToIpAddressPtrOutputWithContext(context.Background())
}

func (i IpAddressArgs) ToIpAddressPtrOutputWithContext(ctx context.Context) IpAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOutput).ToIpAddressPtrOutputWithContext(ctx)
}









type IpAddressPtrInput interface {
	pulumi.Input

	ToIpAddressPtrOutput() IpAddressPtrOutput
	ToIpAddressPtrOutputWithContext(context.Context) IpAddressPtrOutput
}

type ipAddressPtrType IpAddressArgs

func IpAddressPtr(v *IpAddressArgs) IpAddressPtrInput {
	return (*ipAddressPtrType)(v)
}

func (*ipAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAddress)(nil)).Elem()
}

func (i *ipAddressPtrType) ToIpAddressPtrOutput() IpAddressPtrOutput {
	return i.ToIpAddressPtrOutputWithContext(context.Background())
}

func (i *ipAddressPtrType) ToIpAddressPtrOutputWithContext(ctx context.Context) IpAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressPtrOutput)
}

type IpAddressOutput struct{ *pulumi.OutputState }

func (IpAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddress)(nil)).Elem()
}

func (o IpAddressOutput) ToIpAddressOutput() IpAddressOutput {
	return o
}

func (o IpAddressOutput) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return o
}

func (o IpAddressOutput) ToIpAddressPtrOutput() IpAddressPtrOutput {
	return o.ToIpAddressPtrOutputWithContext(context.Background())
}

func (o IpAddressOutput) ToIpAddressPtrOutputWithContext(ctx context.Context) IpAddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpAddress) *IpAddress {
		return &v
	}).(IpAddressPtrOutput)
}

func (o IpAddressOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddress) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o IpAddressOutput) Ports() PortArrayOutput {
	return o.ApplyT(func(v IpAddress) []Port { return v.Ports }).(PortArrayOutput)
}

func (o IpAddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IpAddress) string { return v.Type }).(pulumi.StringOutput)
}

type IpAddressPtrOutput struct{ *pulumi.OutputState }

func (IpAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAddress)(nil)).Elem()
}

func (o IpAddressPtrOutput) ToIpAddressPtrOutput() IpAddressPtrOutput {
	return o
}

func (o IpAddressPtrOutput) ToIpAddressPtrOutputWithContext(ctx context.Context) IpAddressPtrOutput {
	return o
}

func (o IpAddressPtrOutput) Elem() IpAddressOutput {
	return o.ApplyT(func(v *IpAddress) IpAddress {
		if v != nil {
			return *v
		}
		var ret IpAddress
		return ret
	}).(IpAddressOutput)
}

func (o IpAddressPtrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAddress) *string {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(pulumi.StringPtrOutput)
}

func (o IpAddressPtrOutput) Ports() PortArrayOutput {
	return o.ApplyT(func(v *IpAddress) []Port {
		if v == nil {
			return nil
		}
		return v.Ports
	}).(PortArrayOutput)
}

func (o IpAddressPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAddress) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type IpAddressResponse struct {
	Ip    *string        `pulumi:"ip"`
	Ports []PortResponse `pulumi:"ports"`
	Type  string         `pulumi:"type"`
}





type IpAddressResponseInput interface {
	pulumi.Input

	ToIpAddressResponseOutput() IpAddressResponseOutput
	ToIpAddressResponseOutputWithContext(context.Context) IpAddressResponseOutput
}

type IpAddressResponseArgs struct {
	Ip    pulumi.StringPtrInput  `pulumi:"ip"`
	Ports PortResponseArrayInput `pulumi:"ports"`
	Type  pulumi.StringInput     `pulumi:"type"`
}

func (IpAddressResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressResponse)(nil)).Elem()
}

func (i IpAddressResponseArgs) ToIpAddressResponseOutput() IpAddressResponseOutput {
	return i.ToIpAddressResponseOutputWithContext(context.Background())
}

func (i IpAddressResponseArgs) ToIpAddressResponseOutputWithContext(ctx context.Context) IpAddressResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressResponseOutput)
}

func (i IpAddressResponseArgs) ToIpAddressResponsePtrOutput() IpAddressResponsePtrOutput {
	return i.ToIpAddressResponsePtrOutputWithContext(context.Background())
}

func (i IpAddressResponseArgs) ToIpAddressResponsePtrOutputWithContext(ctx context.Context) IpAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressResponseOutput).ToIpAddressResponsePtrOutputWithContext(ctx)
}









type IpAddressResponsePtrInput interface {
	pulumi.Input

	ToIpAddressResponsePtrOutput() IpAddressResponsePtrOutput
	ToIpAddressResponsePtrOutputWithContext(context.Context) IpAddressResponsePtrOutput
}

type ipAddressResponsePtrType IpAddressResponseArgs

func IpAddressResponsePtr(v *IpAddressResponseArgs) IpAddressResponsePtrInput {
	return (*ipAddressResponsePtrType)(v)
}

func (*ipAddressResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAddressResponse)(nil)).Elem()
}

func (i *ipAddressResponsePtrType) ToIpAddressResponsePtrOutput() IpAddressResponsePtrOutput {
	return i.ToIpAddressResponsePtrOutputWithContext(context.Background())
}

func (i *ipAddressResponsePtrType) ToIpAddressResponsePtrOutputWithContext(ctx context.Context) IpAddressResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressResponsePtrOutput)
}

type IpAddressResponseOutput struct{ *pulumi.OutputState }

func (IpAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressResponse)(nil)).Elem()
}

func (o IpAddressResponseOutput) ToIpAddressResponseOutput() IpAddressResponseOutput {
	return o
}

func (o IpAddressResponseOutput) ToIpAddressResponseOutputWithContext(ctx context.Context) IpAddressResponseOutput {
	return o
}

func (o IpAddressResponseOutput) ToIpAddressResponsePtrOutput() IpAddressResponsePtrOutput {
	return o.ToIpAddressResponsePtrOutputWithContext(context.Background())
}

func (o IpAddressResponseOutput) ToIpAddressResponsePtrOutputWithContext(ctx context.Context) IpAddressResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpAddressResponse) *IpAddressResponse {
		return &v
	}).(IpAddressResponsePtrOutput)
}

func (o IpAddressResponseOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressResponse) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o IpAddressResponseOutput) Ports() PortResponseArrayOutput {
	return o.ApplyT(func(v IpAddressResponse) []PortResponse { return v.Ports }).(PortResponseArrayOutput)
}

func (o IpAddressResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IpAddressResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IpAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (IpAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAddressResponse)(nil)).Elem()
}

func (o IpAddressResponsePtrOutput) ToIpAddressResponsePtrOutput() IpAddressResponsePtrOutput {
	return o
}

func (o IpAddressResponsePtrOutput) ToIpAddressResponsePtrOutputWithContext(ctx context.Context) IpAddressResponsePtrOutput {
	return o
}

func (o IpAddressResponsePtrOutput) Elem() IpAddressResponseOutput {
	return o.ApplyT(func(v *IpAddressResponse) IpAddressResponse {
		if v != nil {
			return *v
		}
		var ret IpAddressResponse
		return ret
	}).(IpAddressResponseOutput)
}

func (o IpAddressResponsePtrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(pulumi.StringPtrOutput)
}

func (o IpAddressResponsePtrOutput) Ports() PortResponseArrayOutput {
	return o.ApplyT(func(v *IpAddressResponse) []PortResponse {
		if v == nil {
			return nil
		}
		return v.Ports
	}).(PortResponseArrayOutput)
}

func (o IpAddressResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Port struct {
	Port     int     `pulumi:"port"`
	Protocol *string `pulumi:"protocol"`
}





type PortInput interface {
	pulumi.Input

	ToPortOutput() PortOutput
	ToPortOutputWithContext(context.Context) PortOutput
}

type PortArgs struct {
	Port     pulumi.IntInput       `pulumi:"port"`
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
}

func (PortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Port)(nil)).Elem()
}

func (i PortArgs) ToPortOutput() PortOutput {
	return i.ToPortOutputWithContext(context.Background())
}

func (i PortArgs) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortOutput)
}





type PortArrayInput interface {
	pulumi.Input

	ToPortArrayOutput() PortArrayOutput
	ToPortArrayOutputWithContext(context.Context) PortArrayOutput
}

type PortArray []PortInput

func (PortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Port)(nil)).Elem()
}

func (i PortArray) ToPortArrayOutput() PortArrayOutput {
	return i.ToPortArrayOutputWithContext(context.Background())
}

func (i PortArray) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortArrayOutput)
}

type PortOutput struct{ *pulumi.OutputState }

func (PortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Port)(nil)).Elem()
}

func (o PortOutput) ToPortOutput() PortOutput {
	return o
}

func (o PortOutput) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return o
}

func (o PortOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v Port) int { return v.Port }).(pulumi.IntOutput)
}

func (o PortOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Port) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type PortArrayOutput struct{ *pulumi.OutputState }

func (PortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Port)(nil)).Elem()
}

func (o PortArrayOutput) ToPortArrayOutput() PortArrayOutput {
	return o
}

func (o PortArrayOutput) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return o
}

func (o PortArrayOutput) Index(i pulumi.IntInput) PortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Port {
		return vs[0].([]Port)[vs[1].(int)]
	}).(PortOutput)
}

type PortResponse struct {
	Port     int     `pulumi:"port"`
	Protocol *string `pulumi:"protocol"`
}





type PortResponseInput interface {
	pulumi.Input

	ToPortResponseOutput() PortResponseOutput
	ToPortResponseOutputWithContext(context.Context) PortResponseOutput
}

type PortResponseArgs struct {
	Port     pulumi.IntInput       `pulumi:"port"`
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
}

func (PortResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PortResponse)(nil)).Elem()
}

func (i PortResponseArgs) ToPortResponseOutput() PortResponseOutput {
	return i.ToPortResponseOutputWithContext(context.Background())
}

func (i PortResponseArgs) ToPortResponseOutputWithContext(ctx context.Context) PortResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortResponseOutput)
}





type PortResponseArrayInput interface {
	pulumi.Input

	ToPortResponseArrayOutput() PortResponseArrayOutput
	ToPortResponseArrayOutputWithContext(context.Context) PortResponseArrayOutput
}

type PortResponseArray []PortResponseInput

func (PortResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PortResponse)(nil)).Elem()
}

func (i PortResponseArray) ToPortResponseArrayOutput() PortResponseArrayOutput {
	return i.ToPortResponseArrayOutputWithContext(context.Background())
}

func (i PortResponseArray) ToPortResponseArrayOutputWithContext(ctx context.Context) PortResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortResponseArrayOutput)
}

type PortResponseOutput struct{ *pulumi.OutputState }

func (PortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortResponse)(nil)).Elem()
}

func (o PortResponseOutput) ToPortResponseOutput() PortResponseOutput {
	return o
}

func (o PortResponseOutput) ToPortResponseOutputWithContext(ctx context.Context) PortResponseOutput {
	return o
}

func (o PortResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v PortResponse) int { return v.Port }).(pulumi.IntOutput)
}

func (o PortResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PortResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type PortResponseArrayOutput struct{ *pulumi.OutputState }

func (PortResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PortResponse)(nil)).Elem()
}

func (o PortResponseArrayOutput) ToPortResponseArrayOutput() PortResponseArrayOutput {
	return o
}

func (o PortResponseArrayOutput) ToPortResponseArrayOutputWithContext(ctx context.Context) PortResponseArrayOutput {
	return o
}

func (o PortResponseArrayOutput) Index(i pulumi.IntInput) PortResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PortResponse {
		return vs[0].([]PortResponse)[vs[1].(int)]
	}).(PortResponseOutput)
}

type ResourceLimits struct {
	Cpu        *float64 `pulumi:"cpu"`
	MemoryInGB *float64 `pulumi:"memoryInGB"`
}





type ResourceLimitsInput interface {
	pulumi.Input

	ToResourceLimitsOutput() ResourceLimitsOutput
	ToResourceLimitsOutputWithContext(context.Context) ResourceLimitsOutput
}

type ResourceLimitsArgs struct {
	Cpu        pulumi.Float64PtrInput `pulumi:"cpu"`
	MemoryInGB pulumi.Float64PtrInput `pulumi:"memoryInGB"`
}

func (ResourceLimitsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimits)(nil)).Elem()
}

func (i ResourceLimitsArgs) ToResourceLimitsOutput() ResourceLimitsOutput {
	return i.ToResourceLimitsOutputWithContext(context.Background())
}

func (i ResourceLimitsArgs) ToResourceLimitsOutputWithContext(ctx context.Context) ResourceLimitsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsOutput)
}

func (i ResourceLimitsArgs) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return i.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (i ResourceLimitsArgs) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsOutput).ToResourceLimitsPtrOutputWithContext(ctx)
}









type ResourceLimitsPtrInput interface {
	pulumi.Input

	ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput
	ToResourceLimitsPtrOutputWithContext(context.Context) ResourceLimitsPtrOutput
}

type resourceLimitsPtrType ResourceLimitsArgs

func ResourceLimitsPtr(v *ResourceLimitsArgs) ResourceLimitsPtrInput {
	return (*resourceLimitsPtrType)(v)
}

func (*resourceLimitsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimits)(nil)).Elem()
}

func (i *resourceLimitsPtrType) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return i.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (i *resourceLimitsPtrType) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsPtrOutput)
}

type ResourceLimitsOutput struct{ *pulumi.OutputState }

func (ResourceLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimits)(nil)).Elem()
}

func (o ResourceLimitsOutput) ToResourceLimitsOutput() ResourceLimitsOutput {
	return o
}

func (o ResourceLimitsOutput) ToResourceLimitsOutputWithContext(ctx context.Context) ResourceLimitsOutput {
	return o
}

func (o ResourceLimitsOutput) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return o.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (o ResourceLimitsOutput) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceLimits) *ResourceLimits {
		return &v
	}).(ResourceLimitsPtrOutput)
}

func (o ResourceLimitsOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimits) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimits) *float64 { return v.MemoryInGB }).(pulumi.Float64PtrOutput)
}

type ResourceLimitsPtrOutput struct{ *pulumi.OutputState }

func (ResourceLimitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimits)(nil)).Elem()
}

func (o ResourceLimitsPtrOutput) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return o
}

func (o ResourceLimitsPtrOutput) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return o
}

func (o ResourceLimitsPtrOutput) Elem() ResourceLimitsOutput {
	return o.ApplyT(func(v *ResourceLimits) ResourceLimits {
		if v != nil {
			return *v
		}
		var ret ResourceLimits
		return ret
	}).(ResourceLimitsOutput)
}

func (o ResourceLimitsPtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsPtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

type ResourceLimitsResponse struct {
	Cpu        *float64 `pulumi:"cpu"`
	MemoryInGB *float64 `pulumi:"memoryInGB"`
}





type ResourceLimitsResponseInput interface {
	pulumi.Input

	ToResourceLimitsResponseOutput() ResourceLimitsResponseOutput
	ToResourceLimitsResponseOutputWithContext(context.Context) ResourceLimitsResponseOutput
}

type ResourceLimitsResponseArgs struct {
	Cpu        pulumi.Float64PtrInput `pulumi:"cpu"`
	MemoryInGB pulumi.Float64PtrInput `pulumi:"memoryInGB"`
}

func (ResourceLimitsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimitsResponse)(nil)).Elem()
}

func (i ResourceLimitsResponseArgs) ToResourceLimitsResponseOutput() ResourceLimitsResponseOutput {
	return i.ToResourceLimitsResponseOutputWithContext(context.Background())
}

func (i ResourceLimitsResponseArgs) ToResourceLimitsResponseOutputWithContext(ctx context.Context) ResourceLimitsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsResponseOutput)
}

func (i ResourceLimitsResponseArgs) ToResourceLimitsResponsePtrOutput() ResourceLimitsResponsePtrOutput {
	return i.ToResourceLimitsResponsePtrOutputWithContext(context.Background())
}

func (i ResourceLimitsResponseArgs) ToResourceLimitsResponsePtrOutputWithContext(ctx context.Context) ResourceLimitsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsResponseOutput).ToResourceLimitsResponsePtrOutputWithContext(ctx)
}









type ResourceLimitsResponsePtrInput interface {
	pulumi.Input

	ToResourceLimitsResponsePtrOutput() ResourceLimitsResponsePtrOutput
	ToResourceLimitsResponsePtrOutputWithContext(context.Context) ResourceLimitsResponsePtrOutput
}

type resourceLimitsResponsePtrType ResourceLimitsResponseArgs

func ResourceLimitsResponsePtr(v *ResourceLimitsResponseArgs) ResourceLimitsResponsePtrInput {
	return (*resourceLimitsResponsePtrType)(v)
}

func (*resourceLimitsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimitsResponse)(nil)).Elem()
}

func (i *resourceLimitsResponsePtrType) ToResourceLimitsResponsePtrOutput() ResourceLimitsResponsePtrOutput {
	return i.ToResourceLimitsResponsePtrOutputWithContext(context.Background())
}

func (i *resourceLimitsResponsePtrType) ToResourceLimitsResponsePtrOutputWithContext(ctx context.Context) ResourceLimitsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsResponsePtrOutput)
}

type ResourceLimitsResponseOutput struct{ *pulumi.OutputState }

func (ResourceLimitsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimitsResponse)(nil)).Elem()
}

func (o ResourceLimitsResponseOutput) ToResourceLimitsResponseOutput() ResourceLimitsResponseOutput {
	return o
}

func (o ResourceLimitsResponseOutput) ToResourceLimitsResponseOutputWithContext(ctx context.Context) ResourceLimitsResponseOutput {
	return o
}

func (o ResourceLimitsResponseOutput) ToResourceLimitsResponsePtrOutput() ResourceLimitsResponsePtrOutput {
	return o.ToResourceLimitsResponsePtrOutputWithContext(context.Background())
}

func (o ResourceLimitsResponseOutput) ToResourceLimitsResponsePtrOutputWithContext(ctx context.Context) ResourceLimitsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceLimitsResponse) *ResourceLimitsResponse {
		return &v
	}).(ResourceLimitsResponsePtrOutput)
}

func (o ResourceLimitsResponseOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimitsResponse) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsResponseOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimitsResponse) *float64 { return v.MemoryInGB }).(pulumi.Float64PtrOutput)
}

type ResourceLimitsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceLimitsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimitsResponse)(nil)).Elem()
}

func (o ResourceLimitsResponsePtrOutput) ToResourceLimitsResponsePtrOutput() ResourceLimitsResponsePtrOutput {
	return o
}

func (o ResourceLimitsResponsePtrOutput) ToResourceLimitsResponsePtrOutputWithContext(ctx context.Context) ResourceLimitsResponsePtrOutput {
	return o
}

func (o ResourceLimitsResponsePtrOutput) Elem() ResourceLimitsResponseOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) ResourceLimitsResponse {
		if v != nil {
			return *v
		}
		var ret ResourceLimitsResponse
		return ret
	}).(ResourceLimitsResponseOutput)
}

func (o ResourceLimitsResponsePtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsResponsePtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

type ResourceRequests struct {
	Cpu        float64 `pulumi:"cpu"`
	MemoryInGB float64 `pulumi:"memoryInGB"`
}





type ResourceRequestsInput interface {
	pulumi.Input

	ToResourceRequestsOutput() ResourceRequestsOutput
	ToResourceRequestsOutputWithContext(context.Context) ResourceRequestsOutput
}

type ResourceRequestsArgs struct {
	Cpu        pulumi.Float64Input `pulumi:"cpu"`
	MemoryInGB pulumi.Float64Input `pulumi:"memoryInGB"`
}

func (ResourceRequestsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequests)(nil)).Elem()
}

func (i ResourceRequestsArgs) ToResourceRequestsOutput() ResourceRequestsOutput {
	return i.ToResourceRequestsOutputWithContext(context.Background())
}

func (i ResourceRequestsArgs) ToResourceRequestsOutputWithContext(ctx context.Context) ResourceRequestsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsOutput)
}

type ResourceRequestsOutput struct{ *pulumi.OutputState }

func (ResourceRequestsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequests)(nil)).Elem()
}

func (o ResourceRequestsOutput) ToResourceRequestsOutput() ResourceRequestsOutput {
	return o
}

func (o ResourceRequestsOutput) ToResourceRequestsOutputWithContext(ctx context.Context) ResourceRequestsOutput {
	return o
}

func (o ResourceRequestsOutput) Cpu() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequests) float64 { return v.Cpu }).(pulumi.Float64Output)
}

func (o ResourceRequestsOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequests) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

type ResourceRequestsResponse struct {
	Cpu        float64 `pulumi:"cpu"`
	MemoryInGB float64 `pulumi:"memoryInGB"`
}





type ResourceRequestsResponseInput interface {
	pulumi.Input

	ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput
	ToResourceRequestsResponseOutputWithContext(context.Context) ResourceRequestsResponseOutput
}

type ResourceRequestsResponseArgs struct {
	Cpu        pulumi.Float64Input `pulumi:"cpu"`
	MemoryInGB pulumi.Float64Input `pulumi:"memoryInGB"`
}

func (ResourceRequestsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequestsResponse)(nil)).Elem()
}

func (i ResourceRequestsResponseArgs) ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput {
	return i.ToResourceRequestsResponseOutputWithContext(context.Background())
}

func (i ResourceRequestsResponseArgs) ToResourceRequestsResponseOutputWithContext(ctx context.Context) ResourceRequestsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsResponseOutput)
}

type ResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (ResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequestsResponse)(nil)).Elem()
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput {
	return o
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponseOutputWithContext(ctx context.Context) ResourceRequestsResponseOutput {
	return o
}

func (o ResourceRequestsResponseOutput) Cpu() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequestsResponse) float64 { return v.Cpu }).(pulumi.Float64Output)
}

func (o ResourceRequestsResponseOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequestsResponse) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

type ResourceRequirements struct {
	Limits   *ResourceLimits  `pulumi:"limits"`
	Requests ResourceRequests `pulumi:"requests"`
}





type ResourceRequirementsInput interface {
	pulumi.Input

	ToResourceRequirementsOutput() ResourceRequirementsOutput
	ToResourceRequirementsOutputWithContext(context.Context) ResourceRequirementsOutput
}

type ResourceRequirementsArgs struct {
	Limits   ResourceLimitsPtrInput `pulumi:"limits"`
	Requests ResourceRequestsInput  `pulumi:"requests"`
}

func (ResourceRequirementsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirements)(nil)).Elem()
}

func (i ResourceRequirementsArgs) ToResourceRequirementsOutput() ResourceRequirementsOutput {
	return i.ToResourceRequirementsOutputWithContext(context.Background())
}

func (i ResourceRequirementsArgs) ToResourceRequirementsOutputWithContext(ctx context.Context) ResourceRequirementsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequirementsOutput)
}

type ResourceRequirementsOutput struct{ *pulumi.OutputState }

func (ResourceRequirementsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirements)(nil)).Elem()
}

func (o ResourceRequirementsOutput) ToResourceRequirementsOutput() ResourceRequirementsOutput {
	return o
}

func (o ResourceRequirementsOutput) ToResourceRequirementsOutputWithContext(ctx context.Context) ResourceRequirementsOutput {
	return o
}

func (o ResourceRequirementsOutput) Limits() ResourceLimitsPtrOutput {
	return o.ApplyT(func(v ResourceRequirements) *ResourceLimits { return v.Limits }).(ResourceLimitsPtrOutput)
}

func (o ResourceRequirementsOutput) Requests() ResourceRequestsOutput {
	return o.ApplyT(func(v ResourceRequirements) ResourceRequests { return v.Requests }).(ResourceRequestsOutput)
}

type ResourceRequirementsResponse struct {
	Limits   *ResourceLimitsResponse  `pulumi:"limits"`
	Requests ResourceRequestsResponse `pulumi:"requests"`
}





type ResourceRequirementsResponseInput interface {
	pulumi.Input

	ToResourceRequirementsResponseOutput() ResourceRequirementsResponseOutput
	ToResourceRequirementsResponseOutputWithContext(context.Context) ResourceRequirementsResponseOutput
}

type ResourceRequirementsResponseArgs struct {
	Limits   ResourceLimitsResponsePtrInput `pulumi:"limits"`
	Requests ResourceRequestsResponseInput  `pulumi:"requests"`
}

func (ResourceRequirementsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirementsResponse)(nil)).Elem()
}

func (i ResourceRequirementsResponseArgs) ToResourceRequirementsResponseOutput() ResourceRequirementsResponseOutput {
	return i.ToResourceRequirementsResponseOutputWithContext(context.Background())
}

func (i ResourceRequirementsResponseArgs) ToResourceRequirementsResponseOutputWithContext(ctx context.Context) ResourceRequirementsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequirementsResponseOutput)
}

type ResourceRequirementsResponseOutput struct{ *pulumi.OutputState }

func (ResourceRequirementsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirementsResponse)(nil)).Elem()
}

func (o ResourceRequirementsResponseOutput) ToResourceRequirementsResponseOutput() ResourceRequirementsResponseOutput {
	return o
}

func (o ResourceRequirementsResponseOutput) ToResourceRequirementsResponseOutputWithContext(ctx context.Context) ResourceRequirementsResponseOutput {
	return o
}

func (o ResourceRequirementsResponseOutput) Limits() ResourceLimitsResponsePtrOutput {
	return o.ApplyT(func(v ResourceRequirementsResponse) *ResourceLimitsResponse { return v.Limits }).(ResourceLimitsResponsePtrOutput)
}

func (o ResourceRequirementsResponseOutput) Requests() ResourceRequestsResponseOutput {
	return o.ApplyT(func(v ResourceRequirementsResponse) ResourceRequestsResponse { return v.Requests }).(ResourceRequestsResponseOutput)
}

type Volume struct {
	AzureFile *AzureFileVolume `pulumi:"azureFile"`
	EmptyDir  interface{}      `pulumi:"emptyDir"`
	Name      string           `pulumi:"name"`
}





type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(context.Context) VolumeOutput
}

type VolumeArgs struct {
	AzureFile AzureFileVolumePtrInput `pulumi:"azureFile"`
	EmptyDir  pulumi.Input            `pulumi:"emptyDir"`
	Name      pulumi.StringInput      `pulumi:"name"`
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil)).Elem()
}

func (i VolumeArgs) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i VolumeArgs) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}





type VolumeArrayInput interface {
	pulumi.Input

	ToVolumeArrayOutput() VolumeArrayOutput
	ToVolumeArrayOutputWithContext(context.Context) VolumeArrayOutput
}

type VolumeArray []VolumeInput

func (VolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Volume)(nil)).Elem()
}

func (i VolumeArray) ToVolumeArrayOutput() VolumeArrayOutput {
	return i.ToVolumeArrayOutputWithContext(context.Background())
}

func (i VolumeArray) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArrayOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

func (o VolumeOutput) AzureFile() AzureFileVolumePtrOutput {
	return o.ApplyT(func(v Volume) *AzureFileVolume { return v.AzureFile }).(AzureFileVolumePtrOutput)
}

func (o VolumeOutput) EmptyDir() pulumi.AnyOutput {
	return o.ApplyT(func(v Volume) interface{} { return v.EmptyDir }).(pulumi.AnyOutput)
}

func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Volume) string { return v.Name }).(pulumi.StringOutput)
}

type VolumeArrayOutput struct{ *pulumi.OutputState }

func (VolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Volume)(nil)).Elem()
}

func (o VolumeArrayOutput) ToVolumeArrayOutput() VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) Index(i pulumi.IntInput) VolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Volume {
		return vs[0].([]Volume)[vs[1].(int)]
	}).(VolumeOutput)
}

type VolumeMount struct {
	MountPath string `pulumi:"mountPath"`
	Name      string `pulumi:"name"`
	ReadOnly  *bool  `pulumi:"readOnly"`
}





type VolumeMountInput interface {
	pulumi.Input

	ToVolumeMountOutput() VolumeMountOutput
	ToVolumeMountOutputWithContext(context.Context) VolumeMountOutput
}

type VolumeMountArgs struct {
	MountPath pulumi.StringInput  `pulumi:"mountPath"`
	Name      pulumi.StringInput  `pulumi:"name"`
	ReadOnly  pulumi.BoolPtrInput `pulumi:"readOnly"`
}

func (VolumeMountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeMount)(nil)).Elem()
}

func (i VolumeMountArgs) ToVolumeMountOutput() VolumeMountOutput {
	return i.ToVolumeMountOutputWithContext(context.Background())
}

func (i VolumeMountArgs) ToVolumeMountOutputWithContext(ctx context.Context) VolumeMountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMountOutput)
}





type VolumeMountArrayInput interface {
	pulumi.Input

	ToVolumeMountArrayOutput() VolumeMountArrayOutput
	ToVolumeMountArrayOutputWithContext(context.Context) VolumeMountArrayOutput
}

type VolumeMountArray []VolumeMountInput

func (VolumeMountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeMount)(nil)).Elem()
}

func (i VolumeMountArray) ToVolumeMountArrayOutput() VolumeMountArrayOutput {
	return i.ToVolumeMountArrayOutputWithContext(context.Background())
}

func (i VolumeMountArray) ToVolumeMountArrayOutputWithContext(ctx context.Context) VolumeMountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMountArrayOutput)
}

type VolumeMountOutput struct{ *pulumi.OutputState }

func (VolumeMountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeMount)(nil)).Elem()
}

func (o VolumeMountOutput) ToVolumeMountOutput() VolumeMountOutput {
	return o
}

func (o VolumeMountOutput) ToVolumeMountOutputWithContext(ctx context.Context) VolumeMountOutput {
	return o
}

func (o VolumeMountOutput) MountPath() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeMount) string { return v.MountPath }).(pulumi.StringOutput)
}

func (o VolumeMountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeMount) string { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeMountOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeMount) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

type VolumeMountArrayOutput struct{ *pulumi.OutputState }

func (VolumeMountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeMount)(nil)).Elem()
}

func (o VolumeMountArrayOutput) ToVolumeMountArrayOutput() VolumeMountArrayOutput {
	return o
}

func (o VolumeMountArrayOutput) ToVolumeMountArrayOutputWithContext(ctx context.Context) VolumeMountArrayOutput {
	return o
}

func (o VolumeMountArrayOutput) Index(i pulumi.IntInput) VolumeMountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeMount {
		return vs[0].([]VolumeMount)[vs[1].(int)]
	}).(VolumeMountOutput)
}

type VolumeMountResponse struct {
	MountPath string `pulumi:"mountPath"`
	Name      string `pulumi:"name"`
	ReadOnly  *bool  `pulumi:"readOnly"`
}





type VolumeMountResponseInput interface {
	pulumi.Input

	ToVolumeMountResponseOutput() VolumeMountResponseOutput
	ToVolumeMountResponseOutputWithContext(context.Context) VolumeMountResponseOutput
}

type VolumeMountResponseArgs struct {
	MountPath pulumi.StringInput  `pulumi:"mountPath"`
	Name      pulumi.StringInput  `pulumi:"name"`
	ReadOnly  pulumi.BoolPtrInput `pulumi:"readOnly"`
}

func (VolumeMountResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeMountResponse)(nil)).Elem()
}

func (i VolumeMountResponseArgs) ToVolumeMountResponseOutput() VolumeMountResponseOutput {
	return i.ToVolumeMountResponseOutputWithContext(context.Background())
}

func (i VolumeMountResponseArgs) ToVolumeMountResponseOutputWithContext(ctx context.Context) VolumeMountResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMountResponseOutput)
}





type VolumeMountResponseArrayInput interface {
	pulumi.Input

	ToVolumeMountResponseArrayOutput() VolumeMountResponseArrayOutput
	ToVolumeMountResponseArrayOutputWithContext(context.Context) VolumeMountResponseArrayOutput
}

type VolumeMountResponseArray []VolumeMountResponseInput

func (VolumeMountResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeMountResponse)(nil)).Elem()
}

func (i VolumeMountResponseArray) ToVolumeMountResponseArrayOutput() VolumeMountResponseArrayOutput {
	return i.ToVolumeMountResponseArrayOutputWithContext(context.Background())
}

func (i VolumeMountResponseArray) ToVolumeMountResponseArrayOutputWithContext(ctx context.Context) VolumeMountResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMountResponseArrayOutput)
}

type VolumeMountResponseOutput struct{ *pulumi.OutputState }

func (VolumeMountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeMountResponse)(nil)).Elem()
}

func (o VolumeMountResponseOutput) ToVolumeMountResponseOutput() VolumeMountResponseOutput {
	return o
}

func (o VolumeMountResponseOutput) ToVolumeMountResponseOutputWithContext(ctx context.Context) VolumeMountResponseOutput {
	return o
}

func (o VolumeMountResponseOutput) MountPath() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeMountResponse) string { return v.MountPath }).(pulumi.StringOutput)
}

func (o VolumeMountResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeMountResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeMountResponseOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeMountResponse) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

type VolumeMountResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeMountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeMountResponse)(nil)).Elem()
}

func (o VolumeMountResponseArrayOutput) ToVolumeMountResponseArrayOutput() VolumeMountResponseArrayOutput {
	return o
}

func (o VolumeMountResponseArrayOutput) ToVolumeMountResponseArrayOutputWithContext(ctx context.Context) VolumeMountResponseArrayOutput {
	return o
}

func (o VolumeMountResponseArrayOutput) Index(i pulumi.IntInput) VolumeMountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeMountResponse {
		return vs[0].([]VolumeMountResponse)[vs[1].(int)]
	}).(VolumeMountResponseOutput)
}

type VolumeResponse struct {
	AzureFile *AzureFileVolumeResponse `pulumi:"azureFile"`
	EmptyDir  interface{}              `pulumi:"emptyDir"`
	Name      string                   `pulumi:"name"`
}





type VolumeResponseInput interface {
	pulumi.Input

	ToVolumeResponseOutput() VolumeResponseOutput
	ToVolumeResponseOutputWithContext(context.Context) VolumeResponseOutput
}

type VolumeResponseArgs struct {
	AzureFile AzureFileVolumeResponsePtrInput `pulumi:"azureFile"`
	EmptyDir  pulumi.Input                    `pulumi:"emptyDir"`
	Name      pulumi.StringInput              `pulumi:"name"`
}

func (VolumeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeResponse)(nil)).Elem()
}

func (i VolumeResponseArgs) ToVolumeResponseOutput() VolumeResponseOutput {
	return i.ToVolumeResponseOutputWithContext(context.Background())
}

func (i VolumeResponseArgs) ToVolumeResponseOutputWithContext(ctx context.Context) VolumeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeResponseOutput)
}





type VolumeResponseArrayInput interface {
	pulumi.Input

	ToVolumeResponseArrayOutput() VolumeResponseArrayOutput
	ToVolumeResponseArrayOutputWithContext(context.Context) VolumeResponseArrayOutput
}

type VolumeResponseArray []VolumeResponseInput

func (VolumeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeResponse)(nil)).Elem()
}

func (i VolumeResponseArray) ToVolumeResponseArrayOutput() VolumeResponseArrayOutput {
	return i.ToVolumeResponseArrayOutputWithContext(context.Background())
}

func (i VolumeResponseArray) ToVolumeResponseArrayOutputWithContext(ctx context.Context) VolumeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeResponseArrayOutput)
}

type VolumeResponseOutput struct{ *pulumi.OutputState }

func (VolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeResponse)(nil)).Elem()
}

func (o VolumeResponseOutput) ToVolumeResponseOutput() VolumeResponseOutput {
	return o
}

func (o VolumeResponseOutput) ToVolumeResponseOutputWithContext(ctx context.Context) VolumeResponseOutput {
	return o
}

func (o VolumeResponseOutput) AzureFile() AzureFileVolumeResponsePtrOutput {
	return o.ApplyT(func(v VolumeResponse) *AzureFileVolumeResponse { return v.AzureFile }).(AzureFileVolumeResponsePtrOutput)
}

func (o VolumeResponseOutput) EmptyDir() pulumi.AnyOutput {
	return o.ApplyT(func(v VolumeResponse) interface{} { return v.EmptyDir }).(pulumi.AnyOutput)
}

func (o VolumeResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeResponse) string { return v.Name }).(pulumi.StringOutput)
}

type VolumeResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeResponse)(nil)).Elem()
}

func (o VolumeResponseArrayOutput) ToVolumeResponseArrayOutput() VolumeResponseArrayOutput {
	return o
}

func (o VolumeResponseArrayOutput) ToVolumeResponseArrayOutputWithContext(ctx context.Context) VolumeResponseArrayOutput {
	return o
}

func (o VolumeResponseArrayOutput) Index(i pulumi.IntInput) VolumeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeResponse {
		return vs[0].([]VolumeResponse)[vs[1].(int)]
	}).(VolumeResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureFileVolumeOutput{})
	pulumi.RegisterOutputType(AzureFileVolumePtrOutput{})
	pulumi.RegisterOutputType(AzureFileVolumeResponseOutput{})
	pulumi.RegisterOutputType(AzureFileVolumeResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerArrayOutput{})
	pulumi.RegisterOutputType(ContainerGroupResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(ContainerGroupResponseInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(ContainerPortOutput{})
	pulumi.RegisterOutputType(ContainerPortArrayOutput{})
	pulumi.RegisterOutputType(ContainerPortResponseOutput{})
	pulumi.RegisterOutputType(ContainerPortResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerPropertiesResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(ContainerResponseOutput{})
	pulumi.RegisterOutputType(ContainerResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerStateResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseArrayOutput{})
	pulumi.RegisterOutputType(EventResponseOutput{})
	pulumi.RegisterOutputType(EventResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialArrayOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseArrayOutput{})
	pulumi.RegisterOutputType(IpAddressOutput{})
	pulumi.RegisterOutputType(IpAddressPtrOutput{})
	pulumi.RegisterOutputType(IpAddressResponseOutput{})
	pulumi.RegisterOutputType(IpAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(PortOutput{})
	pulumi.RegisterOutputType(PortArrayOutput{})
	pulumi.RegisterOutputType(PortResponseOutput{})
	pulumi.RegisterOutputType(PortResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceLimitsOutput{})
	pulumi.RegisterOutputType(ResourceLimitsPtrOutput{})
	pulumi.RegisterOutputType(ResourceLimitsResponseOutput{})
	pulumi.RegisterOutputType(ResourceLimitsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceRequestsOutput{})
	pulumi.RegisterOutputType(ResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(ResourceRequirementsOutput{})
	pulumi.RegisterOutputType(ResourceRequirementsResponseOutput{})
	pulumi.RegisterOutputType(VolumeOutput{})
	pulumi.RegisterOutputType(VolumeArrayOutput{})
	pulumi.RegisterOutputType(VolumeMountOutput{})
	pulumi.RegisterOutputType(VolumeMountArrayOutput{})
	pulumi.RegisterOutputType(VolumeMountResponseOutput{})
	pulumi.RegisterOutputType(VolumeMountResponseArrayOutput{})
	pulumi.RegisterOutputType(VolumeResponseOutput{})
	pulumi.RegisterOutputType(VolumeResponseArrayOutput{})
}
