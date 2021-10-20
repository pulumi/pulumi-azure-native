


package v20210901

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
	LivenessProbe        *ContainerProbe       `pulumi:"livenessProbe"`
	Name                 string                `pulumi:"name"`
	Ports                []ContainerPort       `pulumi:"ports"`
	ReadinessProbe       *ContainerProbe       `pulumi:"readinessProbe"`
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
	LivenessProbe        ContainerProbePtrInput        `pulumi:"livenessProbe"`
	Name                 pulumi.StringInput            `pulumi:"name"`
	Ports                ContainerPortArrayInput       `pulumi:"ports"`
	ReadinessProbe       ContainerProbePtrInput        `pulumi:"readinessProbe"`
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

func (o ContainerOutput) LivenessProbe() ContainerProbePtrOutput {
	return o.ApplyT(func(v Container) *ContainerProbe { return v.LivenessProbe }).(ContainerProbePtrOutput)
}

func (o ContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Container) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerOutput) Ports() ContainerPortArrayOutput {
	return o.ApplyT(func(v Container) []ContainerPort { return v.Ports }).(ContainerPortArrayOutput)
}

func (o ContainerOutput) ReadinessProbe() ContainerProbePtrOutput {
	return o.ApplyT(func(v Container) *ContainerProbe { return v.ReadinessProbe }).(ContainerProbePtrOutput)
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

type ContainerExec struct {
	Command []string `pulumi:"command"`
}





type ContainerExecInput interface {
	pulumi.Input

	ToContainerExecOutput() ContainerExecOutput
	ToContainerExecOutputWithContext(context.Context) ContainerExecOutput
}

type ContainerExecArgs struct {
	Command pulumi.StringArrayInput `pulumi:"command"`
}

func (ContainerExecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerExec)(nil)).Elem()
}

func (i ContainerExecArgs) ToContainerExecOutput() ContainerExecOutput {
	return i.ToContainerExecOutputWithContext(context.Background())
}

func (i ContainerExecArgs) ToContainerExecOutputWithContext(ctx context.Context) ContainerExecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerExecOutput)
}

func (i ContainerExecArgs) ToContainerExecPtrOutput() ContainerExecPtrOutput {
	return i.ToContainerExecPtrOutputWithContext(context.Background())
}

func (i ContainerExecArgs) ToContainerExecPtrOutputWithContext(ctx context.Context) ContainerExecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerExecOutput).ToContainerExecPtrOutputWithContext(ctx)
}









type ContainerExecPtrInput interface {
	pulumi.Input

	ToContainerExecPtrOutput() ContainerExecPtrOutput
	ToContainerExecPtrOutputWithContext(context.Context) ContainerExecPtrOutput
}

type containerExecPtrType ContainerExecArgs

func ContainerExecPtr(v *ContainerExecArgs) ContainerExecPtrInput {
	return (*containerExecPtrType)(v)
}

func (*containerExecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerExec)(nil)).Elem()
}

func (i *containerExecPtrType) ToContainerExecPtrOutput() ContainerExecPtrOutput {
	return i.ToContainerExecPtrOutputWithContext(context.Background())
}

func (i *containerExecPtrType) ToContainerExecPtrOutputWithContext(ctx context.Context) ContainerExecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerExecPtrOutput)
}

type ContainerExecOutput struct{ *pulumi.OutputState }

func (ContainerExecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerExec)(nil)).Elem()
}

func (o ContainerExecOutput) ToContainerExecOutput() ContainerExecOutput {
	return o
}

func (o ContainerExecOutput) ToContainerExecOutputWithContext(ctx context.Context) ContainerExecOutput {
	return o
}

func (o ContainerExecOutput) ToContainerExecPtrOutput() ContainerExecPtrOutput {
	return o.ToContainerExecPtrOutputWithContext(context.Background())
}

func (o ContainerExecOutput) ToContainerExecPtrOutputWithContext(ctx context.Context) ContainerExecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerExec) *ContainerExec {
		return &v
	}).(ContainerExecPtrOutput)
}

func (o ContainerExecOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerExec) []string { return v.Command }).(pulumi.StringArrayOutput)
}

type ContainerExecPtrOutput struct{ *pulumi.OutputState }

func (ContainerExecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerExec)(nil)).Elem()
}

func (o ContainerExecPtrOutput) ToContainerExecPtrOutput() ContainerExecPtrOutput {
	return o
}

func (o ContainerExecPtrOutput) ToContainerExecPtrOutputWithContext(ctx context.Context) ContainerExecPtrOutput {
	return o
}

func (o ContainerExecPtrOutput) Elem() ContainerExecOutput {
	return o.ApplyT(func(v *ContainerExec) ContainerExec {
		if v != nil {
			return *v
		}
		var ret ContainerExec
		return ret
	}).(ContainerExecOutput)
}

func (o ContainerExecPtrOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerExec) []string {
		if v == nil {
			return nil
		}
		return v.Command
	}).(pulumi.StringArrayOutput)
}

type ContainerExecResponse struct {
	Command []string `pulumi:"command"`
}





type ContainerExecResponseInput interface {
	pulumi.Input

	ToContainerExecResponseOutput() ContainerExecResponseOutput
	ToContainerExecResponseOutputWithContext(context.Context) ContainerExecResponseOutput
}

type ContainerExecResponseArgs struct {
	Command pulumi.StringArrayInput `pulumi:"command"`
}

func (ContainerExecResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerExecResponse)(nil)).Elem()
}

func (i ContainerExecResponseArgs) ToContainerExecResponseOutput() ContainerExecResponseOutput {
	return i.ToContainerExecResponseOutputWithContext(context.Background())
}

func (i ContainerExecResponseArgs) ToContainerExecResponseOutputWithContext(ctx context.Context) ContainerExecResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerExecResponseOutput)
}

func (i ContainerExecResponseArgs) ToContainerExecResponsePtrOutput() ContainerExecResponsePtrOutput {
	return i.ToContainerExecResponsePtrOutputWithContext(context.Background())
}

func (i ContainerExecResponseArgs) ToContainerExecResponsePtrOutputWithContext(ctx context.Context) ContainerExecResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerExecResponseOutput).ToContainerExecResponsePtrOutputWithContext(ctx)
}









type ContainerExecResponsePtrInput interface {
	pulumi.Input

	ToContainerExecResponsePtrOutput() ContainerExecResponsePtrOutput
	ToContainerExecResponsePtrOutputWithContext(context.Context) ContainerExecResponsePtrOutput
}

type containerExecResponsePtrType ContainerExecResponseArgs

func ContainerExecResponsePtr(v *ContainerExecResponseArgs) ContainerExecResponsePtrInput {
	return (*containerExecResponsePtrType)(v)
}

func (*containerExecResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerExecResponse)(nil)).Elem()
}

func (i *containerExecResponsePtrType) ToContainerExecResponsePtrOutput() ContainerExecResponsePtrOutput {
	return i.ToContainerExecResponsePtrOutputWithContext(context.Background())
}

func (i *containerExecResponsePtrType) ToContainerExecResponsePtrOutputWithContext(ctx context.Context) ContainerExecResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerExecResponsePtrOutput)
}

type ContainerExecResponseOutput struct{ *pulumi.OutputState }

func (ContainerExecResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerExecResponse)(nil)).Elem()
}

func (o ContainerExecResponseOutput) ToContainerExecResponseOutput() ContainerExecResponseOutput {
	return o
}

func (o ContainerExecResponseOutput) ToContainerExecResponseOutputWithContext(ctx context.Context) ContainerExecResponseOutput {
	return o
}

func (o ContainerExecResponseOutput) ToContainerExecResponsePtrOutput() ContainerExecResponsePtrOutput {
	return o.ToContainerExecResponsePtrOutputWithContext(context.Background())
}

func (o ContainerExecResponseOutput) ToContainerExecResponsePtrOutputWithContext(ctx context.Context) ContainerExecResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerExecResponse) *ContainerExecResponse {
		return &v
	}).(ContainerExecResponsePtrOutput)
}

func (o ContainerExecResponseOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerExecResponse) []string { return v.Command }).(pulumi.StringArrayOutput)
}

type ContainerExecResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerExecResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerExecResponse)(nil)).Elem()
}

func (o ContainerExecResponsePtrOutput) ToContainerExecResponsePtrOutput() ContainerExecResponsePtrOutput {
	return o
}

func (o ContainerExecResponsePtrOutput) ToContainerExecResponsePtrOutputWithContext(ctx context.Context) ContainerExecResponsePtrOutput {
	return o
}

func (o ContainerExecResponsePtrOutput) Elem() ContainerExecResponseOutput {
	return o.ApplyT(func(v *ContainerExecResponse) ContainerExecResponse {
		if v != nil {
			return *v
		}
		var ret ContainerExecResponse
		return ret
	}).(ContainerExecResponseOutput)
}

func (o ContainerExecResponsePtrOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerExecResponse) []string {
		if v == nil {
			return nil
		}
		return v.Command
	}).(pulumi.StringArrayOutput)
}

type ContainerGroupDiagnostics struct {
	LogAnalytics *LogAnalytics `pulumi:"logAnalytics"`
}





type ContainerGroupDiagnosticsInput interface {
	pulumi.Input

	ToContainerGroupDiagnosticsOutput() ContainerGroupDiagnosticsOutput
	ToContainerGroupDiagnosticsOutputWithContext(context.Context) ContainerGroupDiagnosticsOutput
}

type ContainerGroupDiagnosticsArgs struct {
	LogAnalytics LogAnalyticsPtrInput `pulumi:"logAnalytics"`
}

func (ContainerGroupDiagnosticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupDiagnostics)(nil)).Elem()
}

func (i ContainerGroupDiagnosticsArgs) ToContainerGroupDiagnosticsOutput() ContainerGroupDiagnosticsOutput {
	return i.ToContainerGroupDiagnosticsOutputWithContext(context.Background())
}

func (i ContainerGroupDiagnosticsArgs) ToContainerGroupDiagnosticsOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupDiagnosticsOutput)
}

func (i ContainerGroupDiagnosticsArgs) ToContainerGroupDiagnosticsPtrOutput() ContainerGroupDiagnosticsPtrOutput {
	return i.ToContainerGroupDiagnosticsPtrOutputWithContext(context.Background())
}

func (i ContainerGroupDiagnosticsArgs) ToContainerGroupDiagnosticsPtrOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupDiagnosticsOutput).ToContainerGroupDiagnosticsPtrOutputWithContext(ctx)
}









type ContainerGroupDiagnosticsPtrInput interface {
	pulumi.Input

	ToContainerGroupDiagnosticsPtrOutput() ContainerGroupDiagnosticsPtrOutput
	ToContainerGroupDiagnosticsPtrOutputWithContext(context.Context) ContainerGroupDiagnosticsPtrOutput
}

type containerGroupDiagnosticsPtrType ContainerGroupDiagnosticsArgs

func ContainerGroupDiagnosticsPtr(v *ContainerGroupDiagnosticsArgs) ContainerGroupDiagnosticsPtrInput {
	return (*containerGroupDiagnosticsPtrType)(v)
}

func (*containerGroupDiagnosticsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupDiagnostics)(nil)).Elem()
}

func (i *containerGroupDiagnosticsPtrType) ToContainerGroupDiagnosticsPtrOutput() ContainerGroupDiagnosticsPtrOutput {
	return i.ToContainerGroupDiagnosticsPtrOutputWithContext(context.Background())
}

func (i *containerGroupDiagnosticsPtrType) ToContainerGroupDiagnosticsPtrOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupDiagnosticsPtrOutput)
}

type ContainerGroupDiagnosticsOutput struct{ *pulumi.OutputState }

func (ContainerGroupDiagnosticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupDiagnostics)(nil)).Elem()
}

func (o ContainerGroupDiagnosticsOutput) ToContainerGroupDiagnosticsOutput() ContainerGroupDiagnosticsOutput {
	return o
}

func (o ContainerGroupDiagnosticsOutput) ToContainerGroupDiagnosticsOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsOutput {
	return o
}

func (o ContainerGroupDiagnosticsOutput) ToContainerGroupDiagnosticsPtrOutput() ContainerGroupDiagnosticsPtrOutput {
	return o.ToContainerGroupDiagnosticsPtrOutputWithContext(context.Background())
}

func (o ContainerGroupDiagnosticsOutput) ToContainerGroupDiagnosticsPtrOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroupDiagnostics) *ContainerGroupDiagnostics {
		return &v
	}).(ContainerGroupDiagnosticsPtrOutput)
}

func (o ContainerGroupDiagnosticsOutput) LogAnalytics() LogAnalyticsPtrOutput {
	return o.ApplyT(func(v ContainerGroupDiagnostics) *LogAnalytics { return v.LogAnalytics }).(LogAnalyticsPtrOutput)
}

type ContainerGroupDiagnosticsPtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupDiagnosticsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupDiagnostics)(nil)).Elem()
}

func (o ContainerGroupDiagnosticsPtrOutput) ToContainerGroupDiagnosticsPtrOutput() ContainerGroupDiagnosticsPtrOutput {
	return o
}

func (o ContainerGroupDiagnosticsPtrOutput) ToContainerGroupDiagnosticsPtrOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsPtrOutput {
	return o
}

func (o ContainerGroupDiagnosticsPtrOutput) Elem() ContainerGroupDiagnosticsOutput {
	return o.ApplyT(func(v *ContainerGroupDiagnostics) ContainerGroupDiagnostics {
		if v != nil {
			return *v
		}
		var ret ContainerGroupDiagnostics
		return ret
	}).(ContainerGroupDiagnosticsOutput)
}

func (o ContainerGroupDiagnosticsPtrOutput) LogAnalytics() LogAnalyticsPtrOutput {
	return o.ApplyT(func(v *ContainerGroupDiagnostics) *LogAnalytics {
		if v == nil {
			return nil
		}
		return v.LogAnalytics
	}).(LogAnalyticsPtrOutput)
}

type ContainerGroupDiagnosticsResponse struct {
	LogAnalytics *LogAnalyticsResponse `pulumi:"logAnalytics"`
}





type ContainerGroupDiagnosticsResponseInput interface {
	pulumi.Input

	ToContainerGroupDiagnosticsResponseOutput() ContainerGroupDiagnosticsResponseOutput
	ToContainerGroupDiagnosticsResponseOutputWithContext(context.Context) ContainerGroupDiagnosticsResponseOutput
}

type ContainerGroupDiagnosticsResponseArgs struct {
	LogAnalytics LogAnalyticsResponsePtrInput `pulumi:"logAnalytics"`
}

func (ContainerGroupDiagnosticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupDiagnosticsResponse)(nil)).Elem()
}

func (i ContainerGroupDiagnosticsResponseArgs) ToContainerGroupDiagnosticsResponseOutput() ContainerGroupDiagnosticsResponseOutput {
	return i.ToContainerGroupDiagnosticsResponseOutputWithContext(context.Background())
}

func (i ContainerGroupDiagnosticsResponseArgs) ToContainerGroupDiagnosticsResponseOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupDiagnosticsResponseOutput)
}

func (i ContainerGroupDiagnosticsResponseArgs) ToContainerGroupDiagnosticsResponsePtrOutput() ContainerGroupDiagnosticsResponsePtrOutput {
	return i.ToContainerGroupDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (i ContainerGroupDiagnosticsResponseArgs) ToContainerGroupDiagnosticsResponsePtrOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupDiagnosticsResponseOutput).ToContainerGroupDiagnosticsResponsePtrOutputWithContext(ctx)
}









type ContainerGroupDiagnosticsResponsePtrInput interface {
	pulumi.Input

	ToContainerGroupDiagnosticsResponsePtrOutput() ContainerGroupDiagnosticsResponsePtrOutput
	ToContainerGroupDiagnosticsResponsePtrOutputWithContext(context.Context) ContainerGroupDiagnosticsResponsePtrOutput
}

type containerGroupDiagnosticsResponsePtrType ContainerGroupDiagnosticsResponseArgs

func ContainerGroupDiagnosticsResponsePtr(v *ContainerGroupDiagnosticsResponseArgs) ContainerGroupDiagnosticsResponsePtrInput {
	return (*containerGroupDiagnosticsResponsePtrType)(v)
}

func (*containerGroupDiagnosticsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupDiagnosticsResponse)(nil)).Elem()
}

func (i *containerGroupDiagnosticsResponsePtrType) ToContainerGroupDiagnosticsResponsePtrOutput() ContainerGroupDiagnosticsResponsePtrOutput {
	return i.ToContainerGroupDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (i *containerGroupDiagnosticsResponsePtrType) ToContainerGroupDiagnosticsResponsePtrOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupDiagnosticsResponsePtrOutput)
}

type ContainerGroupDiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (ContainerGroupDiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupDiagnosticsResponse)(nil)).Elem()
}

func (o ContainerGroupDiagnosticsResponseOutput) ToContainerGroupDiagnosticsResponseOutput() ContainerGroupDiagnosticsResponseOutput {
	return o
}

func (o ContainerGroupDiagnosticsResponseOutput) ToContainerGroupDiagnosticsResponseOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsResponseOutput {
	return o
}

func (o ContainerGroupDiagnosticsResponseOutput) ToContainerGroupDiagnosticsResponsePtrOutput() ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ToContainerGroupDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (o ContainerGroupDiagnosticsResponseOutput) ToContainerGroupDiagnosticsResponsePtrOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroupDiagnosticsResponse) *ContainerGroupDiagnosticsResponse {
		return &v
	}).(ContainerGroupDiagnosticsResponsePtrOutput)
}

func (o ContainerGroupDiagnosticsResponseOutput) LogAnalytics() LogAnalyticsResponsePtrOutput {
	return o.ApplyT(func(v ContainerGroupDiagnosticsResponse) *LogAnalyticsResponse { return v.LogAnalytics }).(LogAnalyticsResponsePtrOutput)
}

type ContainerGroupDiagnosticsResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupDiagnosticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupDiagnosticsResponse)(nil)).Elem()
}

func (o ContainerGroupDiagnosticsResponsePtrOutput) ToContainerGroupDiagnosticsResponsePtrOutput() ContainerGroupDiagnosticsResponsePtrOutput {
	return o
}

func (o ContainerGroupDiagnosticsResponsePtrOutput) ToContainerGroupDiagnosticsResponsePtrOutputWithContext(ctx context.Context) ContainerGroupDiagnosticsResponsePtrOutput {
	return o
}

func (o ContainerGroupDiagnosticsResponsePtrOutput) Elem() ContainerGroupDiagnosticsResponseOutput {
	return o.ApplyT(func(v *ContainerGroupDiagnosticsResponse) ContainerGroupDiagnosticsResponse {
		if v != nil {
			return *v
		}
		var ret ContainerGroupDiagnosticsResponse
		return ret
	}).(ContainerGroupDiagnosticsResponseOutput)
}

func (o ContainerGroupDiagnosticsResponsePtrOutput) LogAnalytics() LogAnalyticsResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroupDiagnosticsResponse) *LogAnalyticsResponse {
		if v == nil {
			return nil
		}
		return v.LogAnalytics
	}).(LogAnalyticsResponsePtrOutput)
}

type ContainerGroupIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ContainerGroupIdentityInput interface {
	pulumi.Input

	ToContainerGroupIdentityOutput() ContainerGroupIdentityOutput
	ToContainerGroupIdentityOutputWithContext(context.Context) ContainerGroupIdentityOutput
}

type ContainerGroupIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (ContainerGroupIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupIdentity)(nil)).Elem()
}

func (i ContainerGroupIdentityArgs) ToContainerGroupIdentityOutput() ContainerGroupIdentityOutput {
	return i.ToContainerGroupIdentityOutputWithContext(context.Background())
}

func (i ContainerGroupIdentityArgs) ToContainerGroupIdentityOutputWithContext(ctx context.Context) ContainerGroupIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupIdentityOutput)
}

func (i ContainerGroupIdentityArgs) ToContainerGroupIdentityPtrOutput() ContainerGroupIdentityPtrOutput {
	return i.ToContainerGroupIdentityPtrOutputWithContext(context.Background())
}

func (i ContainerGroupIdentityArgs) ToContainerGroupIdentityPtrOutputWithContext(ctx context.Context) ContainerGroupIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupIdentityOutput).ToContainerGroupIdentityPtrOutputWithContext(ctx)
}









type ContainerGroupIdentityPtrInput interface {
	pulumi.Input

	ToContainerGroupIdentityPtrOutput() ContainerGroupIdentityPtrOutput
	ToContainerGroupIdentityPtrOutputWithContext(context.Context) ContainerGroupIdentityPtrOutput
}

type containerGroupIdentityPtrType ContainerGroupIdentityArgs

func ContainerGroupIdentityPtr(v *ContainerGroupIdentityArgs) ContainerGroupIdentityPtrInput {
	return (*containerGroupIdentityPtrType)(v)
}

func (*containerGroupIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupIdentity)(nil)).Elem()
}

func (i *containerGroupIdentityPtrType) ToContainerGroupIdentityPtrOutput() ContainerGroupIdentityPtrOutput {
	return i.ToContainerGroupIdentityPtrOutputWithContext(context.Background())
}

func (i *containerGroupIdentityPtrType) ToContainerGroupIdentityPtrOutputWithContext(ctx context.Context) ContainerGroupIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupIdentityPtrOutput)
}

type ContainerGroupIdentityOutput struct{ *pulumi.OutputState }

func (ContainerGroupIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupIdentity)(nil)).Elem()
}

func (o ContainerGroupIdentityOutput) ToContainerGroupIdentityOutput() ContainerGroupIdentityOutput {
	return o
}

func (o ContainerGroupIdentityOutput) ToContainerGroupIdentityOutputWithContext(ctx context.Context) ContainerGroupIdentityOutput {
	return o
}

func (o ContainerGroupIdentityOutput) ToContainerGroupIdentityPtrOutput() ContainerGroupIdentityPtrOutput {
	return o.ToContainerGroupIdentityPtrOutputWithContext(context.Background())
}

func (o ContainerGroupIdentityOutput) ToContainerGroupIdentityPtrOutputWithContext(ctx context.Context) ContainerGroupIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroupIdentity) *ContainerGroupIdentity {
		return &v
	}).(ContainerGroupIdentityPtrOutput)
}

func (o ContainerGroupIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ContainerGroupIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o ContainerGroupIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ContainerGroupIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ContainerGroupIdentityPtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupIdentity)(nil)).Elem()
}

func (o ContainerGroupIdentityPtrOutput) ToContainerGroupIdentityPtrOutput() ContainerGroupIdentityPtrOutput {
	return o
}

func (o ContainerGroupIdentityPtrOutput) ToContainerGroupIdentityPtrOutputWithContext(ctx context.Context) ContainerGroupIdentityPtrOutput {
	return o
}

func (o ContainerGroupIdentityPtrOutput) Elem() ContainerGroupIdentityOutput {
	return o.ApplyT(func(v *ContainerGroupIdentity) ContainerGroupIdentity {
		if v != nil {
			return *v
		}
		var ret ContainerGroupIdentity
		return ret
	}).(ContainerGroupIdentityOutput)
}

func (o ContainerGroupIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ContainerGroupIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o ContainerGroupIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ContainerGroupIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ContainerGroupIdentityResponse struct {
	PrincipalId            string                                                          `pulumi:"principalId"`
	TenantId               string                                                          `pulumi:"tenantId"`
	Type                   *string                                                         `pulumi:"type"`
	UserAssignedIdentities map[string]ContainerGroupIdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}





type ContainerGroupIdentityResponseInput interface {
	pulumi.Input

	ToContainerGroupIdentityResponseOutput() ContainerGroupIdentityResponseOutput
	ToContainerGroupIdentityResponseOutputWithContext(context.Context) ContainerGroupIdentityResponseOutput
}

type ContainerGroupIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                                           `pulumi:"principalId"`
	TenantId               pulumi.StringInput                                           `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                                        `pulumi:"type"`
	UserAssignedIdentities ContainerGroupIdentityResponseUserAssignedIdentitiesMapInput `pulumi:"userAssignedIdentities"`
}

func (ContainerGroupIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupIdentityResponse)(nil)).Elem()
}

func (i ContainerGroupIdentityResponseArgs) ToContainerGroupIdentityResponseOutput() ContainerGroupIdentityResponseOutput {
	return i.ToContainerGroupIdentityResponseOutputWithContext(context.Background())
}

func (i ContainerGroupIdentityResponseArgs) ToContainerGroupIdentityResponseOutputWithContext(ctx context.Context) ContainerGroupIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupIdentityResponseOutput)
}

func (i ContainerGroupIdentityResponseArgs) ToContainerGroupIdentityResponsePtrOutput() ContainerGroupIdentityResponsePtrOutput {
	return i.ToContainerGroupIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ContainerGroupIdentityResponseArgs) ToContainerGroupIdentityResponsePtrOutputWithContext(ctx context.Context) ContainerGroupIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupIdentityResponseOutput).ToContainerGroupIdentityResponsePtrOutputWithContext(ctx)
}









type ContainerGroupIdentityResponsePtrInput interface {
	pulumi.Input

	ToContainerGroupIdentityResponsePtrOutput() ContainerGroupIdentityResponsePtrOutput
	ToContainerGroupIdentityResponsePtrOutputWithContext(context.Context) ContainerGroupIdentityResponsePtrOutput
}

type containerGroupIdentityResponsePtrType ContainerGroupIdentityResponseArgs

func ContainerGroupIdentityResponsePtr(v *ContainerGroupIdentityResponseArgs) ContainerGroupIdentityResponsePtrInput {
	return (*containerGroupIdentityResponsePtrType)(v)
}

func (*containerGroupIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupIdentityResponse)(nil)).Elem()
}

func (i *containerGroupIdentityResponsePtrType) ToContainerGroupIdentityResponsePtrOutput() ContainerGroupIdentityResponsePtrOutput {
	return i.ToContainerGroupIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *containerGroupIdentityResponsePtrType) ToContainerGroupIdentityResponsePtrOutputWithContext(ctx context.Context) ContainerGroupIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupIdentityResponsePtrOutput)
}

type ContainerGroupIdentityResponseOutput struct{ *pulumi.OutputState }

func (ContainerGroupIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupIdentityResponse)(nil)).Elem()
}

func (o ContainerGroupIdentityResponseOutput) ToContainerGroupIdentityResponseOutput() ContainerGroupIdentityResponseOutput {
	return o
}

func (o ContainerGroupIdentityResponseOutput) ToContainerGroupIdentityResponseOutputWithContext(ctx context.Context) ContainerGroupIdentityResponseOutput {
	return o
}

func (o ContainerGroupIdentityResponseOutput) ToContainerGroupIdentityResponsePtrOutput() ContainerGroupIdentityResponsePtrOutput {
	return o.ToContainerGroupIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ContainerGroupIdentityResponseOutput) ToContainerGroupIdentityResponsePtrOutputWithContext(ctx context.Context) ContainerGroupIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroupIdentityResponse) *ContainerGroupIdentityResponse {
		return &v
	}).(ContainerGroupIdentityResponsePtrOutput)
}

func (o ContainerGroupIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerGroupIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ContainerGroupIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerGroupIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ContainerGroupIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerGroupIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ContainerGroupIdentityResponseOutput) UserAssignedIdentities() ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v ContainerGroupIdentityResponse) map[string]ContainerGroupIdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ContainerGroupIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupIdentityResponse)(nil)).Elem()
}

func (o ContainerGroupIdentityResponsePtrOutput) ToContainerGroupIdentityResponsePtrOutput() ContainerGroupIdentityResponsePtrOutput {
	return o
}

func (o ContainerGroupIdentityResponsePtrOutput) ToContainerGroupIdentityResponsePtrOutputWithContext(ctx context.Context) ContainerGroupIdentityResponsePtrOutput {
	return o
}

func (o ContainerGroupIdentityResponsePtrOutput) Elem() ContainerGroupIdentityResponseOutput {
	return o.ApplyT(func(v *ContainerGroupIdentityResponse) ContainerGroupIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ContainerGroupIdentityResponse
		return ret
	}).(ContainerGroupIdentityResponseOutput)
}

func (o ContainerGroupIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroupIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerGroupIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroupIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerGroupIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroupIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ContainerGroupIdentityResponsePtrOutput) UserAssignedIdentities() ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *ContainerGroupIdentityResponse) map[string]ContainerGroupIdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ContainerGroupIdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type ContainerGroupIdentityResponseUserAssignedIdentitiesInput interface {
	pulumi.Input

	ToContainerGroupIdentityResponseUserAssignedIdentitiesOutput() ContainerGroupIdentityResponseUserAssignedIdentitiesOutput
	ToContainerGroupIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Context) ContainerGroupIdentityResponseUserAssignedIdentitiesOutput
}

type ContainerGroupIdentityResponseUserAssignedIdentitiesArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ContainerGroupIdentityResponseUserAssignedIdentitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ContainerGroupIdentityResponseUserAssignedIdentitiesArgs) ToContainerGroupIdentityResponseUserAssignedIdentitiesOutput() ContainerGroupIdentityResponseUserAssignedIdentitiesOutput {
	return i.ToContainerGroupIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Background())
}

func (i ContainerGroupIdentityResponseUserAssignedIdentitiesArgs) ToContainerGroupIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ContainerGroupIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupIdentityResponseUserAssignedIdentitiesOutput)
}





type ContainerGroupIdentityResponseUserAssignedIdentitiesMapInput interface {
	pulumi.Input

	ToContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput() ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput
	ToContainerGroupIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Context) ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput
}

type ContainerGroupIdentityResponseUserAssignedIdentitiesMap map[string]ContainerGroupIdentityResponseUserAssignedIdentitiesInput

func (ContainerGroupIdentityResponseUserAssignedIdentitiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ContainerGroupIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ContainerGroupIdentityResponseUserAssignedIdentitiesMap) ToContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput() ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput {
	return i.ToContainerGroupIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Background())
}

func (i ContainerGroupIdentityResponseUserAssignedIdentitiesMap) ToContainerGroupIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ContainerGroupIdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (ContainerGroupIdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ContainerGroupIdentityResponseUserAssignedIdentitiesOutput) ToContainerGroupIdentityResponseUserAssignedIdentitiesOutput() ContainerGroupIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ContainerGroupIdentityResponseUserAssignedIdentitiesOutput) ToContainerGroupIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ContainerGroupIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ContainerGroupIdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerGroupIdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ContainerGroupIdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerGroupIdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ContainerGroupIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput) ToContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput() ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput) ToContainerGroupIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) ContainerGroupIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ContainerGroupIdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]ContainerGroupIdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(ContainerGroupIdentityResponseUserAssignedIdentitiesOutput)
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

type ContainerGroupSubnetId struct {
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
}





type ContainerGroupSubnetIdInput interface {
	pulumi.Input

	ToContainerGroupSubnetIdOutput() ContainerGroupSubnetIdOutput
	ToContainerGroupSubnetIdOutputWithContext(context.Context) ContainerGroupSubnetIdOutput
}

type ContainerGroupSubnetIdArgs struct {
	Id   pulumi.StringInput    `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ContainerGroupSubnetIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupSubnetId)(nil)).Elem()
}

func (i ContainerGroupSubnetIdArgs) ToContainerGroupSubnetIdOutput() ContainerGroupSubnetIdOutput {
	return i.ToContainerGroupSubnetIdOutputWithContext(context.Background())
}

func (i ContainerGroupSubnetIdArgs) ToContainerGroupSubnetIdOutputWithContext(ctx context.Context) ContainerGroupSubnetIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupSubnetIdOutput)
}





type ContainerGroupSubnetIdArrayInput interface {
	pulumi.Input

	ToContainerGroupSubnetIdArrayOutput() ContainerGroupSubnetIdArrayOutput
	ToContainerGroupSubnetIdArrayOutputWithContext(context.Context) ContainerGroupSubnetIdArrayOutput
}

type ContainerGroupSubnetIdArray []ContainerGroupSubnetIdInput

func (ContainerGroupSubnetIdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerGroupSubnetId)(nil)).Elem()
}

func (i ContainerGroupSubnetIdArray) ToContainerGroupSubnetIdArrayOutput() ContainerGroupSubnetIdArrayOutput {
	return i.ToContainerGroupSubnetIdArrayOutputWithContext(context.Background())
}

func (i ContainerGroupSubnetIdArray) ToContainerGroupSubnetIdArrayOutputWithContext(ctx context.Context) ContainerGroupSubnetIdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupSubnetIdArrayOutput)
}

type ContainerGroupSubnetIdOutput struct{ *pulumi.OutputState }

func (ContainerGroupSubnetIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupSubnetId)(nil)).Elem()
}

func (o ContainerGroupSubnetIdOutput) ToContainerGroupSubnetIdOutput() ContainerGroupSubnetIdOutput {
	return o
}

func (o ContainerGroupSubnetIdOutput) ToContainerGroupSubnetIdOutputWithContext(ctx context.Context) ContainerGroupSubnetIdOutput {
	return o
}

func (o ContainerGroupSubnetIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerGroupSubnetId) string { return v.Id }).(pulumi.StringOutput)
}

func (o ContainerGroupSubnetIdOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerGroupSubnetId) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ContainerGroupSubnetIdArrayOutput struct{ *pulumi.OutputState }

func (ContainerGroupSubnetIdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerGroupSubnetId)(nil)).Elem()
}

func (o ContainerGroupSubnetIdArrayOutput) ToContainerGroupSubnetIdArrayOutput() ContainerGroupSubnetIdArrayOutput {
	return o
}

func (o ContainerGroupSubnetIdArrayOutput) ToContainerGroupSubnetIdArrayOutputWithContext(ctx context.Context) ContainerGroupSubnetIdArrayOutput {
	return o
}

func (o ContainerGroupSubnetIdArrayOutput) Index(i pulumi.IntInput) ContainerGroupSubnetIdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerGroupSubnetId {
		return vs[0].([]ContainerGroupSubnetId)[vs[1].(int)]
	}).(ContainerGroupSubnetIdOutput)
}

type ContainerGroupSubnetIdResponse struct {
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
}





type ContainerGroupSubnetIdResponseInput interface {
	pulumi.Input

	ToContainerGroupSubnetIdResponseOutput() ContainerGroupSubnetIdResponseOutput
	ToContainerGroupSubnetIdResponseOutputWithContext(context.Context) ContainerGroupSubnetIdResponseOutput
}

type ContainerGroupSubnetIdResponseArgs struct {
	Id   pulumi.StringInput    `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ContainerGroupSubnetIdResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupSubnetIdResponse)(nil)).Elem()
}

func (i ContainerGroupSubnetIdResponseArgs) ToContainerGroupSubnetIdResponseOutput() ContainerGroupSubnetIdResponseOutput {
	return i.ToContainerGroupSubnetIdResponseOutputWithContext(context.Background())
}

func (i ContainerGroupSubnetIdResponseArgs) ToContainerGroupSubnetIdResponseOutputWithContext(ctx context.Context) ContainerGroupSubnetIdResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupSubnetIdResponseOutput)
}





type ContainerGroupSubnetIdResponseArrayInput interface {
	pulumi.Input

	ToContainerGroupSubnetIdResponseArrayOutput() ContainerGroupSubnetIdResponseArrayOutput
	ToContainerGroupSubnetIdResponseArrayOutputWithContext(context.Context) ContainerGroupSubnetIdResponseArrayOutput
}

type ContainerGroupSubnetIdResponseArray []ContainerGroupSubnetIdResponseInput

func (ContainerGroupSubnetIdResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerGroupSubnetIdResponse)(nil)).Elem()
}

func (i ContainerGroupSubnetIdResponseArray) ToContainerGroupSubnetIdResponseArrayOutput() ContainerGroupSubnetIdResponseArrayOutput {
	return i.ToContainerGroupSubnetIdResponseArrayOutputWithContext(context.Background())
}

func (i ContainerGroupSubnetIdResponseArray) ToContainerGroupSubnetIdResponseArrayOutputWithContext(ctx context.Context) ContainerGroupSubnetIdResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupSubnetIdResponseArrayOutput)
}

type ContainerGroupSubnetIdResponseOutput struct{ *pulumi.OutputState }

func (ContainerGroupSubnetIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupSubnetIdResponse)(nil)).Elem()
}

func (o ContainerGroupSubnetIdResponseOutput) ToContainerGroupSubnetIdResponseOutput() ContainerGroupSubnetIdResponseOutput {
	return o
}

func (o ContainerGroupSubnetIdResponseOutput) ToContainerGroupSubnetIdResponseOutputWithContext(ctx context.Context) ContainerGroupSubnetIdResponseOutput {
	return o
}

func (o ContainerGroupSubnetIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerGroupSubnetIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ContainerGroupSubnetIdResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerGroupSubnetIdResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ContainerGroupSubnetIdResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerGroupSubnetIdResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerGroupSubnetIdResponse)(nil)).Elem()
}

func (o ContainerGroupSubnetIdResponseArrayOutput) ToContainerGroupSubnetIdResponseArrayOutput() ContainerGroupSubnetIdResponseArrayOutput {
	return o
}

func (o ContainerGroupSubnetIdResponseArrayOutput) ToContainerGroupSubnetIdResponseArrayOutputWithContext(ctx context.Context) ContainerGroupSubnetIdResponseArrayOutput {
	return o
}

func (o ContainerGroupSubnetIdResponseArrayOutput) Index(i pulumi.IntInput) ContainerGroupSubnetIdResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerGroupSubnetIdResponse {
		return vs[0].([]ContainerGroupSubnetIdResponse)[vs[1].(int)]
	}).(ContainerGroupSubnetIdResponseOutput)
}

type ContainerHttpGet struct {
	HttpHeaders []HttpHeader `pulumi:"httpHeaders"`
	Path        *string      `pulumi:"path"`
	Port        int          `pulumi:"port"`
	Scheme      *string      `pulumi:"scheme"`
}





type ContainerHttpGetInput interface {
	pulumi.Input

	ToContainerHttpGetOutput() ContainerHttpGetOutput
	ToContainerHttpGetOutputWithContext(context.Context) ContainerHttpGetOutput
}

type ContainerHttpGetArgs struct {
	HttpHeaders HttpHeaderArrayInput  `pulumi:"httpHeaders"`
	Path        pulumi.StringPtrInput `pulumi:"path"`
	Port        pulumi.IntInput       `pulumi:"port"`
	Scheme      pulumi.StringPtrInput `pulumi:"scheme"`
}

func (ContainerHttpGetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerHttpGet)(nil)).Elem()
}

func (i ContainerHttpGetArgs) ToContainerHttpGetOutput() ContainerHttpGetOutput {
	return i.ToContainerHttpGetOutputWithContext(context.Background())
}

func (i ContainerHttpGetArgs) ToContainerHttpGetOutputWithContext(ctx context.Context) ContainerHttpGetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerHttpGetOutput)
}

func (i ContainerHttpGetArgs) ToContainerHttpGetPtrOutput() ContainerHttpGetPtrOutput {
	return i.ToContainerHttpGetPtrOutputWithContext(context.Background())
}

func (i ContainerHttpGetArgs) ToContainerHttpGetPtrOutputWithContext(ctx context.Context) ContainerHttpGetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerHttpGetOutput).ToContainerHttpGetPtrOutputWithContext(ctx)
}









type ContainerHttpGetPtrInput interface {
	pulumi.Input

	ToContainerHttpGetPtrOutput() ContainerHttpGetPtrOutput
	ToContainerHttpGetPtrOutputWithContext(context.Context) ContainerHttpGetPtrOutput
}

type containerHttpGetPtrType ContainerHttpGetArgs

func ContainerHttpGetPtr(v *ContainerHttpGetArgs) ContainerHttpGetPtrInput {
	return (*containerHttpGetPtrType)(v)
}

func (*containerHttpGetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerHttpGet)(nil)).Elem()
}

func (i *containerHttpGetPtrType) ToContainerHttpGetPtrOutput() ContainerHttpGetPtrOutput {
	return i.ToContainerHttpGetPtrOutputWithContext(context.Background())
}

func (i *containerHttpGetPtrType) ToContainerHttpGetPtrOutputWithContext(ctx context.Context) ContainerHttpGetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerHttpGetPtrOutput)
}

type ContainerHttpGetOutput struct{ *pulumi.OutputState }

func (ContainerHttpGetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerHttpGet)(nil)).Elem()
}

func (o ContainerHttpGetOutput) ToContainerHttpGetOutput() ContainerHttpGetOutput {
	return o
}

func (o ContainerHttpGetOutput) ToContainerHttpGetOutputWithContext(ctx context.Context) ContainerHttpGetOutput {
	return o
}

func (o ContainerHttpGetOutput) ToContainerHttpGetPtrOutput() ContainerHttpGetPtrOutput {
	return o.ToContainerHttpGetPtrOutputWithContext(context.Background())
}

func (o ContainerHttpGetOutput) ToContainerHttpGetPtrOutputWithContext(ctx context.Context) ContainerHttpGetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerHttpGet) *ContainerHttpGet {
		return &v
	}).(ContainerHttpGetPtrOutput)
}

func (o ContainerHttpGetOutput) HttpHeaders() HttpHeaderArrayOutput {
	return o.ApplyT(func(v ContainerHttpGet) []HttpHeader { return v.HttpHeaders }).(HttpHeaderArrayOutput)
}

func (o ContainerHttpGetOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerHttpGet) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ContainerHttpGetOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerHttpGet) int { return v.Port }).(pulumi.IntOutput)
}

func (o ContainerHttpGetOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerHttpGet) *string { return v.Scheme }).(pulumi.StringPtrOutput)
}

type ContainerHttpGetPtrOutput struct{ *pulumi.OutputState }

func (ContainerHttpGetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerHttpGet)(nil)).Elem()
}

func (o ContainerHttpGetPtrOutput) ToContainerHttpGetPtrOutput() ContainerHttpGetPtrOutput {
	return o
}

func (o ContainerHttpGetPtrOutput) ToContainerHttpGetPtrOutputWithContext(ctx context.Context) ContainerHttpGetPtrOutput {
	return o
}

func (o ContainerHttpGetPtrOutput) Elem() ContainerHttpGetOutput {
	return o.ApplyT(func(v *ContainerHttpGet) ContainerHttpGet {
		if v != nil {
			return *v
		}
		var ret ContainerHttpGet
		return ret
	}).(ContainerHttpGetOutput)
}

func (o ContainerHttpGetPtrOutput) HttpHeaders() HttpHeaderArrayOutput {
	return o.ApplyT(func(v *ContainerHttpGet) []HttpHeader {
		if v == nil {
			return nil
		}
		return v.HttpHeaders
	}).(HttpHeaderArrayOutput)
}

func (o ContainerHttpGetPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerHttpGet) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ContainerHttpGetPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerHttpGet) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

func (o ContainerHttpGetPtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerHttpGet) *string {
		if v == nil {
			return nil
		}
		return v.Scheme
	}).(pulumi.StringPtrOutput)
}

type ContainerHttpGetResponse struct {
	HttpHeaders []HttpHeaderResponse `pulumi:"httpHeaders"`
	Path        *string              `pulumi:"path"`
	Port        int                  `pulumi:"port"`
	Scheme      *string              `pulumi:"scheme"`
}





type ContainerHttpGetResponseInput interface {
	pulumi.Input

	ToContainerHttpGetResponseOutput() ContainerHttpGetResponseOutput
	ToContainerHttpGetResponseOutputWithContext(context.Context) ContainerHttpGetResponseOutput
}

type ContainerHttpGetResponseArgs struct {
	HttpHeaders HttpHeaderResponseArrayInput `pulumi:"httpHeaders"`
	Path        pulumi.StringPtrInput        `pulumi:"path"`
	Port        pulumi.IntInput              `pulumi:"port"`
	Scheme      pulumi.StringPtrInput        `pulumi:"scheme"`
}

func (ContainerHttpGetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerHttpGetResponse)(nil)).Elem()
}

func (i ContainerHttpGetResponseArgs) ToContainerHttpGetResponseOutput() ContainerHttpGetResponseOutput {
	return i.ToContainerHttpGetResponseOutputWithContext(context.Background())
}

func (i ContainerHttpGetResponseArgs) ToContainerHttpGetResponseOutputWithContext(ctx context.Context) ContainerHttpGetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerHttpGetResponseOutput)
}

func (i ContainerHttpGetResponseArgs) ToContainerHttpGetResponsePtrOutput() ContainerHttpGetResponsePtrOutput {
	return i.ToContainerHttpGetResponsePtrOutputWithContext(context.Background())
}

func (i ContainerHttpGetResponseArgs) ToContainerHttpGetResponsePtrOutputWithContext(ctx context.Context) ContainerHttpGetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerHttpGetResponseOutput).ToContainerHttpGetResponsePtrOutputWithContext(ctx)
}









type ContainerHttpGetResponsePtrInput interface {
	pulumi.Input

	ToContainerHttpGetResponsePtrOutput() ContainerHttpGetResponsePtrOutput
	ToContainerHttpGetResponsePtrOutputWithContext(context.Context) ContainerHttpGetResponsePtrOutput
}

type containerHttpGetResponsePtrType ContainerHttpGetResponseArgs

func ContainerHttpGetResponsePtr(v *ContainerHttpGetResponseArgs) ContainerHttpGetResponsePtrInput {
	return (*containerHttpGetResponsePtrType)(v)
}

func (*containerHttpGetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerHttpGetResponse)(nil)).Elem()
}

func (i *containerHttpGetResponsePtrType) ToContainerHttpGetResponsePtrOutput() ContainerHttpGetResponsePtrOutput {
	return i.ToContainerHttpGetResponsePtrOutputWithContext(context.Background())
}

func (i *containerHttpGetResponsePtrType) ToContainerHttpGetResponsePtrOutputWithContext(ctx context.Context) ContainerHttpGetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerHttpGetResponsePtrOutput)
}

type ContainerHttpGetResponseOutput struct{ *pulumi.OutputState }

func (ContainerHttpGetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerHttpGetResponse)(nil)).Elem()
}

func (o ContainerHttpGetResponseOutput) ToContainerHttpGetResponseOutput() ContainerHttpGetResponseOutput {
	return o
}

func (o ContainerHttpGetResponseOutput) ToContainerHttpGetResponseOutputWithContext(ctx context.Context) ContainerHttpGetResponseOutput {
	return o
}

func (o ContainerHttpGetResponseOutput) ToContainerHttpGetResponsePtrOutput() ContainerHttpGetResponsePtrOutput {
	return o.ToContainerHttpGetResponsePtrOutputWithContext(context.Background())
}

func (o ContainerHttpGetResponseOutput) ToContainerHttpGetResponsePtrOutputWithContext(ctx context.Context) ContainerHttpGetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerHttpGetResponse) *ContainerHttpGetResponse {
		return &v
	}).(ContainerHttpGetResponsePtrOutput)
}

func (o ContainerHttpGetResponseOutput) HttpHeaders() HttpHeaderResponseArrayOutput {
	return o.ApplyT(func(v ContainerHttpGetResponse) []HttpHeaderResponse { return v.HttpHeaders }).(HttpHeaderResponseArrayOutput)
}

func (o ContainerHttpGetResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerHttpGetResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ContainerHttpGetResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerHttpGetResponse) int { return v.Port }).(pulumi.IntOutput)
}

func (o ContainerHttpGetResponseOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerHttpGetResponse) *string { return v.Scheme }).(pulumi.StringPtrOutput)
}

type ContainerHttpGetResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerHttpGetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerHttpGetResponse)(nil)).Elem()
}

func (o ContainerHttpGetResponsePtrOutput) ToContainerHttpGetResponsePtrOutput() ContainerHttpGetResponsePtrOutput {
	return o
}

func (o ContainerHttpGetResponsePtrOutput) ToContainerHttpGetResponsePtrOutputWithContext(ctx context.Context) ContainerHttpGetResponsePtrOutput {
	return o
}

func (o ContainerHttpGetResponsePtrOutput) Elem() ContainerHttpGetResponseOutput {
	return o.ApplyT(func(v *ContainerHttpGetResponse) ContainerHttpGetResponse {
		if v != nil {
			return *v
		}
		var ret ContainerHttpGetResponse
		return ret
	}).(ContainerHttpGetResponseOutput)
}

func (o ContainerHttpGetResponsePtrOutput) HttpHeaders() HttpHeaderResponseArrayOutput {
	return o.ApplyT(func(v *ContainerHttpGetResponse) []HttpHeaderResponse {
		if v == nil {
			return nil
		}
		return v.HttpHeaders
	}).(HttpHeaderResponseArrayOutput)
}

func (o ContainerHttpGetResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerHttpGetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ContainerHttpGetResponsePtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerHttpGetResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Port
	}).(pulumi.IntPtrOutput)
}

func (o ContainerHttpGetResponsePtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerHttpGetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Scheme
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

type ContainerProbe struct {
	Exec                *ContainerExec    `pulumi:"exec"`
	FailureThreshold    *int              `pulumi:"failureThreshold"`
	HttpGet             *ContainerHttpGet `pulumi:"httpGet"`
	InitialDelaySeconds *int              `pulumi:"initialDelaySeconds"`
	PeriodSeconds       *int              `pulumi:"periodSeconds"`
	SuccessThreshold    *int              `pulumi:"successThreshold"`
	TimeoutSeconds      *int              `pulumi:"timeoutSeconds"`
}





type ContainerProbeInput interface {
	pulumi.Input

	ToContainerProbeOutput() ContainerProbeOutput
	ToContainerProbeOutputWithContext(context.Context) ContainerProbeOutput
}

type ContainerProbeArgs struct {
	Exec                ContainerExecPtrInput    `pulumi:"exec"`
	FailureThreshold    pulumi.IntPtrInput       `pulumi:"failureThreshold"`
	HttpGet             ContainerHttpGetPtrInput `pulumi:"httpGet"`
	InitialDelaySeconds pulumi.IntPtrInput       `pulumi:"initialDelaySeconds"`
	PeriodSeconds       pulumi.IntPtrInput       `pulumi:"periodSeconds"`
	SuccessThreshold    pulumi.IntPtrInput       `pulumi:"successThreshold"`
	TimeoutSeconds      pulumi.IntPtrInput       `pulumi:"timeoutSeconds"`
}

func (ContainerProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerProbe)(nil)).Elem()
}

func (i ContainerProbeArgs) ToContainerProbeOutput() ContainerProbeOutput {
	return i.ToContainerProbeOutputWithContext(context.Background())
}

func (i ContainerProbeArgs) ToContainerProbeOutputWithContext(ctx context.Context) ContainerProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerProbeOutput)
}

func (i ContainerProbeArgs) ToContainerProbePtrOutput() ContainerProbePtrOutput {
	return i.ToContainerProbePtrOutputWithContext(context.Background())
}

func (i ContainerProbeArgs) ToContainerProbePtrOutputWithContext(ctx context.Context) ContainerProbePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerProbeOutput).ToContainerProbePtrOutputWithContext(ctx)
}









type ContainerProbePtrInput interface {
	pulumi.Input

	ToContainerProbePtrOutput() ContainerProbePtrOutput
	ToContainerProbePtrOutputWithContext(context.Context) ContainerProbePtrOutput
}

type containerProbePtrType ContainerProbeArgs

func ContainerProbePtr(v *ContainerProbeArgs) ContainerProbePtrInput {
	return (*containerProbePtrType)(v)
}

func (*containerProbePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerProbe)(nil)).Elem()
}

func (i *containerProbePtrType) ToContainerProbePtrOutput() ContainerProbePtrOutput {
	return i.ToContainerProbePtrOutputWithContext(context.Background())
}

func (i *containerProbePtrType) ToContainerProbePtrOutputWithContext(ctx context.Context) ContainerProbePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerProbePtrOutput)
}

type ContainerProbeOutput struct{ *pulumi.OutputState }

func (ContainerProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerProbe)(nil)).Elem()
}

func (o ContainerProbeOutput) ToContainerProbeOutput() ContainerProbeOutput {
	return o
}

func (o ContainerProbeOutput) ToContainerProbeOutputWithContext(ctx context.Context) ContainerProbeOutput {
	return o
}

func (o ContainerProbeOutput) ToContainerProbePtrOutput() ContainerProbePtrOutput {
	return o.ToContainerProbePtrOutputWithContext(context.Background())
}

func (o ContainerProbeOutput) ToContainerProbePtrOutputWithContext(ctx context.Context) ContainerProbePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerProbe) *ContainerProbe {
		return &v
	}).(ContainerProbePtrOutput)
}

func (o ContainerProbeOutput) Exec() ContainerExecPtrOutput {
	return o.ApplyT(func(v ContainerProbe) *ContainerExec { return v.Exec }).(ContainerExecPtrOutput)
}

func (o ContainerProbeOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbe) *int { return v.FailureThreshold }).(pulumi.IntPtrOutput)
}

func (o ContainerProbeOutput) HttpGet() ContainerHttpGetPtrOutput {
	return o.ApplyT(func(v ContainerProbe) *ContainerHttpGet { return v.HttpGet }).(ContainerHttpGetPtrOutput)
}

func (o ContainerProbeOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbe) *int { return v.InitialDelaySeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerProbeOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbe) *int { return v.PeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerProbeOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbe) *int { return v.SuccessThreshold }).(pulumi.IntPtrOutput)
}

func (o ContainerProbeOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbe) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

type ContainerProbePtrOutput struct{ *pulumi.OutputState }

func (ContainerProbePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerProbe)(nil)).Elem()
}

func (o ContainerProbePtrOutput) ToContainerProbePtrOutput() ContainerProbePtrOutput {
	return o
}

func (o ContainerProbePtrOutput) ToContainerProbePtrOutputWithContext(ctx context.Context) ContainerProbePtrOutput {
	return o
}

func (o ContainerProbePtrOutput) Elem() ContainerProbeOutput {
	return o.ApplyT(func(v *ContainerProbe) ContainerProbe {
		if v != nil {
			return *v
		}
		var ret ContainerProbe
		return ret
	}).(ContainerProbeOutput)
}

func (o ContainerProbePtrOutput) Exec() ContainerExecPtrOutput {
	return o.ApplyT(func(v *ContainerProbe) *ContainerExec {
		if v == nil {
			return nil
		}
		return v.Exec
	}).(ContainerExecPtrOutput)
}

func (o ContainerProbePtrOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbe) *int {
		if v == nil {
			return nil
		}
		return v.FailureThreshold
	}).(pulumi.IntPtrOutput)
}

func (o ContainerProbePtrOutput) HttpGet() ContainerHttpGetPtrOutput {
	return o.ApplyT(func(v *ContainerProbe) *ContainerHttpGet {
		if v == nil {
			return nil
		}
		return v.HttpGet
	}).(ContainerHttpGetPtrOutput)
}

func (o ContainerProbePtrOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbe) *int {
		if v == nil {
			return nil
		}
		return v.InitialDelaySeconds
	}).(pulumi.IntPtrOutput)
}

func (o ContainerProbePtrOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbe) *int {
		if v == nil {
			return nil
		}
		return v.PeriodSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ContainerProbePtrOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbe) *int {
		if v == nil {
			return nil
		}
		return v.SuccessThreshold
	}).(pulumi.IntPtrOutput)
}

func (o ContainerProbePtrOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbe) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type ContainerProbeResponse struct {
	Exec                *ContainerExecResponse    `pulumi:"exec"`
	FailureThreshold    *int                      `pulumi:"failureThreshold"`
	HttpGet             *ContainerHttpGetResponse `pulumi:"httpGet"`
	InitialDelaySeconds *int                      `pulumi:"initialDelaySeconds"`
	PeriodSeconds       *int                      `pulumi:"periodSeconds"`
	SuccessThreshold    *int                      `pulumi:"successThreshold"`
	TimeoutSeconds      *int                      `pulumi:"timeoutSeconds"`
}





type ContainerProbeResponseInput interface {
	pulumi.Input

	ToContainerProbeResponseOutput() ContainerProbeResponseOutput
	ToContainerProbeResponseOutputWithContext(context.Context) ContainerProbeResponseOutput
}

type ContainerProbeResponseArgs struct {
	Exec                ContainerExecResponsePtrInput    `pulumi:"exec"`
	FailureThreshold    pulumi.IntPtrInput               `pulumi:"failureThreshold"`
	HttpGet             ContainerHttpGetResponsePtrInput `pulumi:"httpGet"`
	InitialDelaySeconds pulumi.IntPtrInput               `pulumi:"initialDelaySeconds"`
	PeriodSeconds       pulumi.IntPtrInput               `pulumi:"periodSeconds"`
	SuccessThreshold    pulumi.IntPtrInput               `pulumi:"successThreshold"`
	TimeoutSeconds      pulumi.IntPtrInput               `pulumi:"timeoutSeconds"`
}

func (ContainerProbeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerProbeResponse)(nil)).Elem()
}

func (i ContainerProbeResponseArgs) ToContainerProbeResponseOutput() ContainerProbeResponseOutput {
	return i.ToContainerProbeResponseOutputWithContext(context.Background())
}

func (i ContainerProbeResponseArgs) ToContainerProbeResponseOutputWithContext(ctx context.Context) ContainerProbeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerProbeResponseOutput)
}

func (i ContainerProbeResponseArgs) ToContainerProbeResponsePtrOutput() ContainerProbeResponsePtrOutput {
	return i.ToContainerProbeResponsePtrOutputWithContext(context.Background())
}

func (i ContainerProbeResponseArgs) ToContainerProbeResponsePtrOutputWithContext(ctx context.Context) ContainerProbeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerProbeResponseOutput).ToContainerProbeResponsePtrOutputWithContext(ctx)
}









type ContainerProbeResponsePtrInput interface {
	pulumi.Input

	ToContainerProbeResponsePtrOutput() ContainerProbeResponsePtrOutput
	ToContainerProbeResponsePtrOutputWithContext(context.Context) ContainerProbeResponsePtrOutput
}

type containerProbeResponsePtrType ContainerProbeResponseArgs

func ContainerProbeResponsePtr(v *ContainerProbeResponseArgs) ContainerProbeResponsePtrInput {
	return (*containerProbeResponsePtrType)(v)
}

func (*containerProbeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerProbeResponse)(nil)).Elem()
}

func (i *containerProbeResponsePtrType) ToContainerProbeResponsePtrOutput() ContainerProbeResponsePtrOutput {
	return i.ToContainerProbeResponsePtrOutputWithContext(context.Background())
}

func (i *containerProbeResponsePtrType) ToContainerProbeResponsePtrOutputWithContext(ctx context.Context) ContainerProbeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerProbeResponsePtrOutput)
}

type ContainerProbeResponseOutput struct{ *pulumi.OutputState }

func (ContainerProbeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerProbeResponse)(nil)).Elem()
}

func (o ContainerProbeResponseOutput) ToContainerProbeResponseOutput() ContainerProbeResponseOutput {
	return o
}

func (o ContainerProbeResponseOutput) ToContainerProbeResponseOutputWithContext(ctx context.Context) ContainerProbeResponseOutput {
	return o
}

func (o ContainerProbeResponseOutput) ToContainerProbeResponsePtrOutput() ContainerProbeResponsePtrOutput {
	return o.ToContainerProbeResponsePtrOutputWithContext(context.Background())
}

func (o ContainerProbeResponseOutput) ToContainerProbeResponsePtrOutputWithContext(ctx context.Context) ContainerProbeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerProbeResponse) *ContainerProbeResponse {
		return &v
	}).(ContainerProbeResponsePtrOutput)
}

func (o ContainerProbeResponseOutput) Exec() ContainerExecResponsePtrOutput {
	return o.ApplyT(func(v ContainerProbeResponse) *ContainerExecResponse { return v.Exec }).(ContainerExecResponsePtrOutput)
}

func (o ContainerProbeResponseOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbeResponse) *int { return v.FailureThreshold }).(pulumi.IntPtrOutput)
}

func (o ContainerProbeResponseOutput) HttpGet() ContainerHttpGetResponsePtrOutput {
	return o.ApplyT(func(v ContainerProbeResponse) *ContainerHttpGetResponse { return v.HttpGet }).(ContainerHttpGetResponsePtrOutput)
}

func (o ContainerProbeResponseOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbeResponse) *int { return v.InitialDelaySeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerProbeResponseOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbeResponse) *int { return v.PeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o ContainerProbeResponseOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbeResponse) *int { return v.SuccessThreshold }).(pulumi.IntPtrOutput)
}

func (o ContainerProbeResponseOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerProbeResponse) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

type ContainerProbeResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerProbeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerProbeResponse)(nil)).Elem()
}

func (o ContainerProbeResponsePtrOutput) ToContainerProbeResponsePtrOutput() ContainerProbeResponsePtrOutput {
	return o
}

func (o ContainerProbeResponsePtrOutput) ToContainerProbeResponsePtrOutputWithContext(ctx context.Context) ContainerProbeResponsePtrOutput {
	return o
}

func (o ContainerProbeResponsePtrOutput) Elem() ContainerProbeResponseOutput {
	return o.ApplyT(func(v *ContainerProbeResponse) ContainerProbeResponse {
		if v != nil {
			return *v
		}
		var ret ContainerProbeResponse
		return ret
	}).(ContainerProbeResponseOutput)
}

func (o ContainerProbeResponsePtrOutput) Exec() ContainerExecResponsePtrOutput {
	return o.ApplyT(func(v *ContainerProbeResponse) *ContainerExecResponse {
		if v == nil {
			return nil
		}
		return v.Exec
	}).(ContainerExecResponsePtrOutput)
}

func (o ContainerProbeResponsePtrOutput) FailureThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbeResponse) *int {
		if v == nil {
			return nil
		}
		return v.FailureThreshold
	}).(pulumi.IntPtrOutput)
}

func (o ContainerProbeResponsePtrOutput) HttpGet() ContainerHttpGetResponsePtrOutput {
	return o.ApplyT(func(v *ContainerProbeResponse) *ContainerHttpGetResponse {
		if v == nil {
			return nil
		}
		return v.HttpGet
	}).(ContainerHttpGetResponsePtrOutput)
}

func (o ContainerProbeResponsePtrOutput) InitialDelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbeResponse) *int {
		if v == nil {
			return nil
		}
		return v.InitialDelaySeconds
	}).(pulumi.IntPtrOutput)
}

func (o ContainerProbeResponsePtrOutput) PeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbeResponse) *int {
		if v == nil {
			return nil
		}
		return v.PeriodSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ContainerProbeResponsePtrOutput) SuccessThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbeResponse) *int {
		if v == nil {
			return nil
		}
		return v.SuccessThreshold
	}).(pulumi.IntPtrOutput)
}

func (o ContainerProbeResponsePtrOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerProbeResponse) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutSeconds
	}).(pulumi.IntPtrOutput)
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
	LivenessProbe        *ContainerProbeResponse                 `pulumi:"livenessProbe"`
	Name                 string                                  `pulumi:"name"`
	Ports                []ContainerPortResponse                 `pulumi:"ports"`
	ReadinessProbe       *ContainerProbeResponse                 `pulumi:"readinessProbe"`
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
	LivenessProbe        ContainerProbeResponsePtrInput               `pulumi:"livenessProbe"`
	Name                 pulumi.StringInput                           `pulumi:"name"`
	Ports                ContainerPortResponseArrayInput              `pulumi:"ports"`
	ReadinessProbe       ContainerProbeResponsePtrInput               `pulumi:"readinessProbe"`
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

func (o ContainerResponseOutput) LivenessProbe() ContainerProbeResponsePtrOutput {
	return o.ApplyT(func(v ContainerResponse) *ContainerProbeResponse { return v.LivenessProbe }).(ContainerProbeResponsePtrOutput)
}

func (o ContainerResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerResponseOutput) Ports() ContainerPortResponseArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []ContainerPortResponse { return v.Ports }).(ContainerPortResponseArrayOutput)
}

func (o ContainerResponseOutput) ReadinessProbe() ContainerProbeResponsePtrOutput {
	return o.ApplyT(func(v ContainerResponse) *ContainerProbeResponse { return v.ReadinessProbe }).(ContainerProbeResponsePtrOutput)
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
	DetailStatus string `pulumi:"detailStatus"`
	ExitCode     int    `pulumi:"exitCode"`
	FinishTime   string `pulumi:"finishTime"`
	StartTime    string `pulumi:"startTime"`
	State        string `pulumi:"state"`
}





type ContainerStateResponseInput interface {
	pulumi.Input

	ToContainerStateResponseOutput() ContainerStateResponseOutput
	ToContainerStateResponseOutputWithContext(context.Context) ContainerStateResponseOutput
}

type ContainerStateResponseArgs struct {
	DetailStatus pulumi.StringInput `pulumi:"detailStatus"`
	ExitCode     pulumi.IntInput    `pulumi:"exitCode"`
	FinishTime   pulumi.StringInput `pulumi:"finishTime"`
	StartTime    pulumi.StringInput `pulumi:"startTime"`
	State        pulumi.StringInput `pulumi:"state"`
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

func (o ContainerStateResponseOutput) DetailStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerStateResponse) string { return v.DetailStatus }).(pulumi.StringOutput)
}

func (o ContainerStateResponseOutput) ExitCode() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerStateResponse) int { return v.ExitCode }).(pulumi.IntOutput)
}

func (o ContainerStateResponseOutput) FinishTime() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerStateResponse) string { return v.FinishTime }).(pulumi.StringOutput)
}

func (o ContainerStateResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerStateResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o ContainerStateResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerStateResponse) string { return v.State }).(pulumi.StringOutput)
}

type DnsConfiguration struct {
	NameServers   []string `pulumi:"nameServers"`
	Options       *string  `pulumi:"options"`
	SearchDomains *string  `pulumi:"searchDomains"`
}





type DnsConfigurationInput interface {
	pulumi.Input

	ToDnsConfigurationOutput() DnsConfigurationOutput
	ToDnsConfigurationOutputWithContext(context.Context) DnsConfigurationOutput
}

type DnsConfigurationArgs struct {
	NameServers   pulumi.StringArrayInput `pulumi:"nameServers"`
	Options       pulumi.StringPtrInput   `pulumi:"options"`
	SearchDomains pulumi.StringPtrInput   `pulumi:"searchDomains"`
}

func (DnsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfiguration)(nil)).Elem()
}

func (i DnsConfigurationArgs) ToDnsConfigurationOutput() DnsConfigurationOutput {
	return i.ToDnsConfigurationOutputWithContext(context.Background())
}

func (i DnsConfigurationArgs) ToDnsConfigurationOutputWithContext(ctx context.Context) DnsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigurationOutput)
}

func (i DnsConfigurationArgs) ToDnsConfigurationPtrOutput() DnsConfigurationPtrOutput {
	return i.ToDnsConfigurationPtrOutputWithContext(context.Background())
}

func (i DnsConfigurationArgs) ToDnsConfigurationPtrOutputWithContext(ctx context.Context) DnsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigurationOutput).ToDnsConfigurationPtrOutputWithContext(ctx)
}









type DnsConfigurationPtrInput interface {
	pulumi.Input

	ToDnsConfigurationPtrOutput() DnsConfigurationPtrOutput
	ToDnsConfigurationPtrOutputWithContext(context.Context) DnsConfigurationPtrOutput
}

type dnsConfigurationPtrType DnsConfigurationArgs

func DnsConfigurationPtr(v *DnsConfigurationArgs) DnsConfigurationPtrInput {
	return (*dnsConfigurationPtrType)(v)
}

func (*dnsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfiguration)(nil)).Elem()
}

func (i *dnsConfigurationPtrType) ToDnsConfigurationPtrOutput() DnsConfigurationPtrOutput {
	return i.ToDnsConfigurationPtrOutputWithContext(context.Background())
}

func (i *dnsConfigurationPtrType) ToDnsConfigurationPtrOutputWithContext(ctx context.Context) DnsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigurationPtrOutput)
}

type DnsConfigurationOutput struct{ *pulumi.OutputState }

func (DnsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfiguration)(nil)).Elem()
}

func (o DnsConfigurationOutput) ToDnsConfigurationOutput() DnsConfigurationOutput {
	return o
}

func (o DnsConfigurationOutput) ToDnsConfigurationOutputWithContext(ctx context.Context) DnsConfigurationOutput {
	return o
}

func (o DnsConfigurationOutput) ToDnsConfigurationPtrOutput() DnsConfigurationPtrOutput {
	return o.ToDnsConfigurationPtrOutputWithContext(context.Background())
}

func (o DnsConfigurationOutput) ToDnsConfigurationPtrOutputWithContext(ctx context.Context) DnsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsConfiguration) *DnsConfiguration {
		return &v
	}).(DnsConfigurationPtrOutput)
}

func (o DnsConfigurationOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DnsConfiguration) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

func (o DnsConfigurationOutput) Options() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfiguration) *string { return v.Options }).(pulumi.StringPtrOutput)
}

func (o DnsConfigurationOutput) SearchDomains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfiguration) *string { return v.SearchDomains }).(pulumi.StringPtrOutput)
}

type DnsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (DnsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfiguration)(nil)).Elem()
}

func (o DnsConfigurationPtrOutput) ToDnsConfigurationPtrOutput() DnsConfigurationPtrOutput {
	return o
}

func (o DnsConfigurationPtrOutput) ToDnsConfigurationPtrOutputWithContext(ctx context.Context) DnsConfigurationPtrOutput {
	return o
}

func (o DnsConfigurationPtrOutput) Elem() DnsConfigurationOutput {
	return o.ApplyT(func(v *DnsConfiguration) DnsConfiguration {
		if v != nil {
			return *v
		}
		var ret DnsConfiguration
		return ret
	}).(DnsConfigurationOutput)
}

func (o DnsConfigurationPtrOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DnsConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.NameServers
	}).(pulumi.StringArrayOutput)
}

func (o DnsConfigurationPtrOutput) Options() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Options
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigurationPtrOutput) SearchDomains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SearchDomains
	}).(pulumi.StringPtrOutput)
}

type DnsConfigurationResponse struct {
	NameServers   []string `pulumi:"nameServers"`
	Options       *string  `pulumi:"options"`
	SearchDomains *string  `pulumi:"searchDomains"`
}





type DnsConfigurationResponseInput interface {
	pulumi.Input

	ToDnsConfigurationResponseOutput() DnsConfigurationResponseOutput
	ToDnsConfigurationResponseOutputWithContext(context.Context) DnsConfigurationResponseOutput
}

type DnsConfigurationResponseArgs struct {
	NameServers   pulumi.StringArrayInput `pulumi:"nameServers"`
	Options       pulumi.StringPtrInput   `pulumi:"options"`
	SearchDomains pulumi.StringPtrInput   `pulumi:"searchDomains"`
}

func (DnsConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfigurationResponse)(nil)).Elem()
}

func (i DnsConfigurationResponseArgs) ToDnsConfigurationResponseOutput() DnsConfigurationResponseOutput {
	return i.ToDnsConfigurationResponseOutputWithContext(context.Background())
}

func (i DnsConfigurationResponseArgs) ToDnsConfigurationResponseOutputWithContext(ctx context.Context) DnsConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigurationResponseOutput)
}

func (i DnsConfigurationResponseArgs) ToDnsConfigurationResponsePtrOutput() DnsConfigurationResponsePtrOutput {
	return i.ToDnsConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i DnsConfigurationResponseArgs) ToDnsConfigurationResponsePtrOutputWithContext(ctx context.Context) DnsConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigurationResponseOutput).ToDnsConfigurationResponsePtrOutputWithContext(ctx)
}









type DnsConfigurationResponsePtrInput interface {
	pulumi.Input

	ToDnsConfigurationResponsePtrOutput() DnsConfigurationResponsePtrOutput
	ToDnsConfigurationResponsePtrOutputWithContext(context.Context) DnsConfigurationResponsePtrOutput
}

type dnsConfigurationResponsePtrType DnsConfigurationResponseArgs

func DnsConfigurationResponsePtr(v *DnsConfigurationResponseArgs) DnsConfigurationResponsePtrInput {
	return (*dnsConfigurationResponsePtrType)(v)
}

func (*dnsConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfigurationResponse)(nil)).Elem()
}

func (i *dnsConfigurationResponsePtrType) ToDnsConfigurationResponsePtrOutput() DnsConfigurationResponsePtrOutput {
	return i.ToDnsConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *dnsConfigurationResponsePtrType) ToDnsConfigurationResponsePtrOutputWithContext(ctx context.Context) DnsConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigurationResponsePtrOutput)
}

type DnsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (DnsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfigurationResponse)(nil)).Elem()
}

func (o DnsConfigurationResponseOutput) ToDnsConfigurationResponseOutput() DnsConfigurationResponseOutput {
	return o
}

func (o DnsConfigurationResponseOutput) ToDnsConfigurationResponseOutputWithContext(ctx context.Context) DnsConfigurationResponseOutput {
	return o
}

func (o DnsConfigurationResponseOutput) ToDnsConfigurationResponsePtrOutput() DnsConfigurationResponsePtrOutput {
	return o.ToDnsConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o DnsConfigurationResponseOutput) ToDnsConfigurationResponsePtrOutputWithContext(ctx context.Context) DnsConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsConfigurationResponse) *DnsConfigurationResponse {
		return &v
	}).(DnsConfigurationResponsePtrOutput)
}

func (o DnsConfigurationResponseOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DnsConfigurationResponse) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

func (o DnsConfigurationResponseOutput) Options() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfigurationResponse) *string { return v.Options }).(pulumi.StringPtrOutput)
}

func (o DnsConfigurationResponseOutput) SearchDomains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfigurationResponse) *string { return v.SearchDomains }).(pulumi.StringPtrOutput)
}

type DnsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (DnsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfigurationResponse)(nil)).Elem()
}

func (o DnsConfigurationResponsePtrOutput) ToDnsConfigurationResponsePtrOutput() DnsConfigurationResponsePtrOutput {
	return o
}

func (o DnsConfigurationResponsePtrOutput) ToDnsConfigurationResponsePtrOutputWithContext(ctx context.Context) DnsConfigurationResponsePtrOutput {
	return o
}

func (o DnsConfigurationResponsePtrOutput) Elem() DnsConfigurationResponseOutput {
	return o.ApplyT(func(v *DnsConfigurationResponse) DnsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret DnsConfigurationResponse
		return ret
	}).(DnsConfigurationResponseOutput)
}

func (o DnsConfigurationResponsePtrOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DnsConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.NameServers
	}).(pulumi.StringArrayOutput)
}

func (o DnsConfigurationResponsePtrOutput) Options() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Options
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigurationResponsePtrOutput) SearchDomains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SearchDomains
	}).(pulumi.StringPtrOutput)
}

type EncryptionProperties struct {
	KeyName      string `pulumi:"keyName"`
	KeyVersion   string `pulumi:"keyVersion"`
	VaultBaseUrl string `pulumi:"vaultBaseUrl"`
}





type EncryptionPropertiesInput interface {
	pulumi.Input

	ToEncryptionPropertiesOutput() EncryptionPropertiesOutput
	ToEncryptionPropertiesOutputWithContext(context.Context) EncryptionPropertiesOutput
}

type EncryptionPropertiesArgs struct {
	KeyName      pulumi.StringInput `pulumi:"keyName"`
	KeyVersion   pulumi.StringInput `pulumi:"keyVersion"`
	VaultBaseUrl pulumi.StringInput `pulumi:"vaultBaseUrl"`
}

func (EncryptionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperties)(nil)).Elem()
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesOutput() EncryptionPropertiesOutput {
	return i.ToEncryptionPropertiesOutputWithContext(context.Background())
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesOutputWithContext(ctx context.Context) EncryptionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesOutput)
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return i.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesOutput).ToEncryptionPropertiesPtrOutputWithContext(ctx)
}









type EncryptionPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput
	ToEncryptionPropertiesPtrOutputWithContext(context.Context) EncryptionPropertiesPtrOutput
}

type encryptionPropertiesPtrType EncryptionPropertiesArgs

func EncryptionPropertiesPtr(v *EncryptionPropertiesArgs) EncryptionPropertiesPtrInput {
	return (*encryptionPropertiesPtrType)(v)
}

func (*encryptionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperties)(nil)).Elem()
}

func (i *encryptionPropertiesPtrType) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return i.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesPtrType) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesPtrOutput)
}

type EncryptionPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperties)(nil)).Elem()
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesOutput() EncryptionPropertiesOutput {
	return o
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesOutputWithContext(ctx context.Context) EncryptionPropertiesOutput {
	return o
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return o.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionProperties) *EncryptionProperties {
		return &v
	}).(EncryptionPropertiesPtrOutput)
}

func (o EncryptionPropertiesOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionProperties) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o EncryptionPropertiesOutput) KeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionProperties) string { return v.KeyVersion }).(pulumi.StringOutput)
}

func (o EncryptionPropertiesOutput) VaultBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionProperties) string { return v.VaultBaseUrl }).(pulumi.StringOutput)
}

type EncryptionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperties)(nil)).Elem()
}

func (o EncryptionPropertiesPtrOutput) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return o
}

func (o EncryptionPropertiesPtrOutput) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return o
}

func (o EncryptionPropertiesPtrOutput) Elem() EncryptionPropertiesOutput {
	return o.ApplyT(func(v *EncryptionProperties) EncryptionProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionProperties
		return ret
	}).(EncryptionPropertiesOutput)
}

func (o EncryptionPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesPtrOutput) VaultBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.VaultBaseUrl
	}).(pulumi.StringPtrOutput)
}

type EncryptionPropertiesResponse struct {
	KeyName      string `pulumi:"keyName"`
	KeyVersion   string `pulumi:"keyVersion"`
	VaultBaseUrl string `pulumi:"vaultBaseUrl"`
}





type EncryptionPropertiesResponseInput interface {
	pulumi.Input

	ToEncryptionPropertiesResponseOutput() EncryptionPropertiesResponseOutput
	ToEncryptionPropertiesResponseOutputWithContext(context.Context) EncryptionPropertiesResponseOutput
}

type EncryptionPropertiesResponseArgs struct {
	KeyName      pulumi.StringInput `pulumi:"keyName"`
	KeyVersion   pulumi.StringInput `pulumi:"keyVersion"`
	VaultBaseUrl pulumi.StringInput `pulumi:"vaultBaseUrl"`
}

func (EncryptionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesResponse)(nil)).Elem()
}

func (i EncryptionPropertiesResponseArgs) ToEncryptionPropertiesResponseOutput() EncryptionPropertiesResponseOutput {
	return i.ToEncryptionPropertiesResponseOutputWithContext(context.Background())
}

func (i EncryptionPropertiesResponseArgs) ToEncryptionPropertiesResponseOutputWithContext(ctx context.Context) EncryptionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesResponseOutput)
}

func (i EncryptionPropertiesResponseArgs) ToEncryptionPropertiesResponsePtrOutput() EncryptionPropertiesResponsePtrOutput {
	return i.ToEncryptionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesResponseArgs) ToEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesResponseOutput).ToEncryptionPropertiesResponsePtrOutputWithContext(ctx)
}









type EncryptionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesResponsePtrOutput() EncryptionPropertiesResponsePtrOutput
	ToEncryptionPropertiesResponsePtrOutputWithContext(context.Context) EncryptionPropertiesResponsePtrOutput
}

type encryptionPropertiesResponsePtrType EncryptionPropertiesResponseArgs

func EncryptionPropertiesResponsePtr(v *EncryptionPropertiesResponseArgs) EncryptionPropertiesResponsePtrInput {
	return (*encryptionPropertiesResponsePtrType)(v)
}

func (*encryptionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesResponse)(nil)).Elem()
}

func (i *encryptionPropertiesResponsePtrType) ToEncryptionPropertiesResponsePtrOutput() EncryptionPropertiesResponsePtrOutput {
	return i.ToEncryptionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesResponsePtrType) ToEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesResponsePtrOutput)
}

type EncryptionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesResponse)(nil)).Elem()
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponseOutput() EncryptionPropertiesResponseOutput {
	return o
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponseOutputWithContext(ctx context.Context) EncryptionPropertiesResponseOutput {
	return o
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponsePtrOutput() EncryptionPropertiesResponsePtrOutput {
	return o.ToEncryptionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPropertiesResponse) *EncryptionPropertiesResponse {
		return &v
	}).(EncryptionPropertiesResponsePtrOutput)
}

func (o EncryptionPropertiesResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o EncryptionPropertiesResponseOutput) KeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) string { return v.KeyVersion }).(pulumi.StringOutput)
}

func (o EncryptionPropertiesResponseOutput) VaultBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) string { return v.VaultBaseUrl }).(pulumi.StringOutput)
}

type EncryptionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesResponse)(nil)).Elem()
}

func (o EncryptionPropertiesResponsePtrOutput) ToEncryptionPropertiesResponsePtrOutput() EncryptionPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesResponsePtrOutput) ToEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesResponsePtrOutput) Elem() EncryptionPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) EncryptionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesResponse
		return ret
	}).(EncryptionPropertiesResponseOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) VaultBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VaultBaseUrl
	}).(pulumi.StringPtrOutput)
}

type EnvironmentVariable struct {
	Name        string  `pulumi:"name"`
	SecureValue *string `pulumi:"secureValue"`
	Value       *string `pulumi:"value"`
}





type EnvironmentVariableInput interface {
	pulumi.Input

	ToEnvironmentVariableOutput() EnvironmentVariableOutput
	ToEnvironmentVariableOutputWithContext(context.Context) EnvironmentVariableOutput
}

type EnvironmentVariableArgs struct {
	Name        pulumi.StringInput    `pulumi:"name"`
	SecureValue pulumi.StringPtrInput `pulumi:"secureValue"`
	Value       pulumi.StringPtrInput `pulumi:"value"`
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

func (o EnvironmentVariableOutput) SecureValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariable) *string { return v.SecureValue }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariable) *string { return v.Value }).(pulumi.StringPtrOutput)
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
	Name        string  `pulumi:"name"`
	SecureValue *string `pulumi:"secureValue"`
	Value       *string `pulumi:"value"`
}





type EnvironmentVariableResponseInput interface {
	pulumi.Input

	ToEnvironmentVariableResponseOutput() EnvironmentVariableResponseOutput
	ToEnvironmentVariableResponseOutputWithContext(context.Context) EnvironmentVariableResponseOutput
}

type EnvironmentVariableResponseArgs struct {
	Name        pulumi.StringInput    `pulumi:"name"`
	SecureValue pulumi.StringPtrInput `pulumi:"secureValue"`
	Value       pulumi.StringPtrInput `pulumi:"value"`
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

func (o EnvironmentVariableResponseOutput) SecureValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariableResponse) *string { return v.SecureValue }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVariableResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariableResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
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
	Count          int    `pulumi:"count"`
	FirstTimestamp string `pulumi:"firstTimestamp"`
	LastTimestamp  string `pulumi:"lastTimestamp"`
	Message        string `pulumi:"message"`
	Name           string `pulumi:"name"`
	Type           string `pulumi:"type"`
}





type EventResponseInput interface {
	pulumi.Input

	ToEventResponseOutput() EventResponseOutput
	ToEventResponseOutputWithContext(context.Context) EventResponseOutput
}

type EventResponseArgs struct {
	Count          pulumi.IntInput    `pulumi:"count"`
	FirstTimestamp pulumi.StringInput `pulumi:"firstTimestamp"`
	LastTimestamp  pulumi.StringInput `pulumi:"lastTimestamp"`
	Message        pulumi.StringInput `pulumi:"message"`
	Name           pulumi.StringInput `pulumi:"name"`
	Type           pulumi.StringInput `pulumi:"type"`
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

func (o EventResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v EventResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o EventResponseOutput) FirstTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v EventResponse) string { return v.FirstTimestamp }).(pulumi.StringOutput)
}

func (o EventResponseOutput) LastTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v EventResponse) string { return v.LastTimestamp }).(pulumi.StringOutput)
}

func (o EventResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v EventResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o EventResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EventResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EventResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventResponse) string { return v.Type }).(pulumi.StringOutput)
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

type GitRepoVolume struct {
	Directory  *string `pulumi:"directory"`
	Repository string  `pulumi:"repository"`
	Revision   *string `pulumi:"revision"`
}





type GitRepoVolumeInput interface {
	pulumi.Input

	ToGitRepoVolumeOutput() GitRepoVolumeOutput
	ToGitRepoVolumeOutputWithContext(context.Context) GitRepoVolumeOutput
}

type GitRepoVolumeArgs struct {
	Directory  pulumi.StringPtrInput `pulumi:"directory"`
	Repository pulumi.StringInput    `pulumi:"repository"`
	Revision   pulumi.StringPtrInput `pulumi:"revision"`
}

func (GitRepoVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepoVolume)(nil)).Elem()
}

func (i GitRepoVolumeArgs) ToGitRepoVolumeOutput() GitRepoVolumeOutput {
	return i.ToGitRepoVolumeOutputWithContext(context.Background())
}

func (i GitRepoVolumeArgs) ToGitRepoVolumeOutputWithContext(ctx context.Context) GitRepoVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepoVolumeOutput)
}

func (i GitRepoVolumeArgs) ToGitRepoVolumePtrOutput() GitRepoVolumePtrOutput {
	return i.ToGitRepoVolumePtrOutputWithContext(context.Background())
}

func (i GitRepoVolumeArgs) ToGitRepoVolumePtrOutputWithContext(ctx context.Context) GitRepoVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepoVolumeOutput).ToGitRepoVolumePtrOutputWithContext(ctx)
}









type GitRepoVolumePtrInput interface {
	pulumi.Input

	ToGitRepoVolumePtrOutput() GitRepoVolumePtrOutput
	ToGitRepoVolumePtrOutputWithContext(context.Context) GitRepoVolumePtrOutput
}

type gitRepoVolumePtrType GitRepoVolumeArgs

func GitRepoVolumePtr(v *GitRepoVolumeArgs) GitRepoVolumePtrInput {
	return (*gitRepoVolumePtrType)(v)
}

func (*gitRepoVolumePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepoVolume)(nil)).Elem()
}

func (i *gitRepoVolumePtrType) ToGitRepoVolumePtrOutput() GitRepoVolumePtrOutput {
	return i.ToGitRepoVolumePtrOutputWithContext(context.Background())
}

func (i *gitRepoVolumePtrType) ToGitRepoVolumePtrOutputWithContext(ctx context.Context) GitRepoVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepoVolumePtrOutput)
}

type GitRepoVolumeOutput struct{ *pulumi.OutputState }

func (GitRepoVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepoVolume)(nil)).Elem()
}

func (o GitRepoVolumeOutput) ToGitRepoVolumeOutput() GitRepoVolumeOutput {
	return o
}

func (o GitRepoVolumeOutput) ToGitRepoVolumeOutputWithContext(ctx context.Context) GitRepoVolumeOutput {
	return o
}

func (o GitRepoVolumeOutput) ToGitRepoVolumePtrOutput() GitRepoVolumePtrOutput {
	return o.ToGitRepoVolumePtrOutputWithContext(context.Background())
}

func (o GitRepoVolumeOutput) ToGitRepoVolumePtrOutputWithContext(ctx context.Context) GitRepoVolumePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitRepoVolume) *GitRepoVolume {
		return &v
	}).(GitRepoVolumePtrOutput)
}

func (o GitRepoVolumeOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepoVolume) *string { return v.Directory }).(pulumi.StringPtrOutput)
}

func (o GitRepoVolumeOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GitRepoVolume) string { return v.Repository }).(pulumi.StringOutput)
}

func (o GitRepoVolumeOutput) Revision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepoVolume) *string { return v.Revision }).(pulumi.StringPtrOutput)
}

type GitRepoVolumePtrOutput struct{ *pulumi.OutputState }

func (GitRepoVolumePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepoVolume)(nil)).Elem()
}

func (o GitRepoVolumePtrOutput) ToGitRepoVolumePtrOutput() GitRepoVolumePtrOutput {
	return o
}

func (o GitRepoVolumePtrOutput) ToGitRepoVolumePtrOutputWithContext(ctx context.Context) GitRepoVolumePtrOutput {
	return o
}

func (o GitRepoVolumePtrOutput) Elem() GitRepoVolumeOutput {
	return o.ApplyT(func(v *GitRepoVolume) GitRepoVolume {
		if v != nil {
			return *v
		}
		var ret GitRepoVolume
		return ret
	}).(GitRepoVolumeOutput)
}

func (o GitRepoVolumePtrOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepoVolume) *string {
		if v == nil {
			return nil
		}
		return v.Directory
	}).(pulumi.StringPtrOutput)
}

func (o GitRepoVolumePtrOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepoVolume) *string {
		if v == nil {
			return nil
		}
		return &v.Repository
	}).(pulumi.StringPtrOutput)
}

func (o GitRepoVolumePtrOutput) Revision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepoVolume) *string {
		if v == nil {
			return nil
		}
		return v.Revision
	}).(pulumi.StringPtrOutput)
}

type GitRepoVolumeResponse struct {
	Directory  *string `pulumi:"directory"`
	Repository string  `pulumi:"repository"`
	Revision   *string `pulumi:"revision"`
}





type GitRepoVolumeResponseInput interface {
	pulumi.Input

	ToGitRepoVolumeResponseOutput() GitRepoVolumeResponseOutput
	ToGitRepoVolumeResponseOutputWithContext(context.Context) GitRepoVolumeResponseOutput
}

type GitRepoVolumeResponseArgs struct {
	Directory  pulumi.StringPtrInput `pulumi:"directory"`
	Repository pulumi.StringInput    `pulumi:"repository"`
	Revision   pulumi.StringPtrInput `pulumi:"revision"`
}

func (GitRepoVolumeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepoVolumeResponse)(nil)).Elem()
}

func (i GitRepoVolumeResponseArgs) ToGitRepoVolumeResponseOutput() GitRepoVolumeResponseOutput {
	return i.ToGitRepoVolumeResponseOutputWithContext(context.Background())
}

func (i GitRepoVolumeResponseArgs) ToGitRepoVolumeResponseOutputWithContext(ctx context.Context) GitRepoVolumeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepoVolumeResponseOutput)
}

func (i GitRepoVolumeResponseArgs) ToGitRepoVolumeResponsePtrOutput() GitRepoVolumeResponsePtrOutput {
	return i.ToGitRepoVolumeResponsePtrOutputWithContext(context.Background())
}

func (i GitRepoVolumeResponseArgs) ToGitRepoVolumeResponsePtrOutputWithContext(ctx context.Context) GitRepoVolumeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepoVolumeResponseOutput).ToGitRepoVolumeResponsePtrOutputWithContext(ctx)
}









type GitRepoVolumeResponsePtrInput interface {
	pulumi.Input

	ToGitRepoVolumeResponsePtrOutput() GitRepoVolumeResponsePtrOutput
	ToGitRepoVolumeResponsePtrOutputWithContext(context.Context) GitRepoVolumeResponsePtrOutput
}

type gitRepoVolumeResponsePtrType GitRepoVolumeResponseArgs

func GitRepoVolumeResponsePtr(v *GitRepoVolumeResponseArgs) GitRepoVolumeResponsePtrInput {
	return (*gitRepoVolumeResponsePtrType)(v)
}

func (*gitRepoVolumeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepoVolumeResponse)(nil)).Elem()
}

func (i *gitRepoVolumeResponsePtrType) ToGitRepoVolumeResponsePtrOutput() GitRepoVolumeResponsePtrOutput {
	return i.ToGitRepoVolumeResponsePtrOutputWithContext(context.Background())
}

func (i *gitRepoVolumeResponsePtrType) ToGitRepoVolumeResponsePtrOutputWithContext(ctx context.Context) GitRepoVolumeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepoVolumeResponsePtrOutput)
}

type GitRepoVolumeResponseOutput struct{ *pulumi.OutputState }

func (GitRepoVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepoVolumeResponse)(nil)).Elem()
}

func (o GitRepoVolumeResponseOutput) ToGitRepoVolumeResponseOutput() GitRepoVolumeResponseOutput {
	return o
}

func (o GitRepoVolumeResponseOutput) ToGitRepoVolumeResponseOutputWithContext(ctx context.Context) GitRepoVolumeResponseOutput {
	return o
}

func (o GitRepoVolumeResponseOutput) ToGitRepoVolumeResponsePtrOutput() GitRepoVolumeResponsePtrOutput {
	return o.ToGitRepoVolumeResponsePtrOutputWithContext(context.Background())
}

func (o GitRepoVolumeResponseOutput) ToGitRepoVolumeResponsePtrOutputWithContext(ctx context.Context) GitRepoVolumeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitRepoVolumeResponse) *GitRepoVolumeResponse {
		return &v
	}).(GitRepoVolumeResponsePtrOutput)
}

func (o GitRepoVolumeResponseOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepoVolumeResponse) *string { return v.Directory }).(pulumi.StringPtrOutput)
}

func (o GitRepoVolumeResponseOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v GitRepoVolumeResponse) string { return v.Repository }).(pulumi.StringOutput)
}

func (o GitRepoVolumeResponseOutput) Revision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepoVolumeResponse) *string { return v.Revision }).(pulumi.StringPtrOutput)
}

type GitRepoVolumeResponsePtrOutput struct{ *pulumi.OutputState }

func (GitRepoVolumeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepoVolumeResponse)(nil)).Elem()
}

func (o GitRepoVolumeResponsePtrOutput) ToGitRepoVolumeResponsePtrOutput() GitRepoVolumeResponsePtrOutput {
	return o
}

func (o GitRepoVolumeResponsePtrOutput) ToGitRepoVolumeResponsePtrOutputWithContext(ctx context.Context) GitRepoVolumeResponsePtrOutput {
	return o
}

func (o GitRepoVolumeResponsePtrOutput) Elem() GitRepoVolumeResponseOutput {
	return o.ApplyT(func(v *GitRepoVolumeResponse) GitRepoVolumeResponse {
		if v != nil {
			return *v
		}
		var ret GitRepoVolumeResponse
		return ret
	}).(GitRepoVolumeResponseOutput)
}

func (o GitRepoVolumeResponsePtrOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepoVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Directory
	}).(pulumi.StringPtrOutput)
}

func (o GitRepoVolumeResponsePtrOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepoVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Repository
	}).(pulumi.StringPtrOutput)
}

func (o GitRepoVolumeResponsePtrOutput) Revision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepoVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Revision
	}).(pulumi.StringPtrOutput)
}

type GpuResource struct {
	Count int    `pulumi:"count"`
	Sku   string `pulumi:"sku"`
}





type GpuResourceInput interface {
	pulumi.Input

	ToGpuResourceOutput() GpuResourceOutput
	ToGpuResourceOutputWithContext(context.Context) GpuResourceOutput
}

type GpuResourceArgs struct {
	Count pulumi.IntInput    `pulumi:"count"`
	Sku   pulumi.StringInput `pulumi:"sku"`
}

func (GpuResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GpuResource)(nil)).Elem()
}

func (i GpuResourceArgs) ToGpuResourceOutput() GpuResourceOutput {
	return i.ToGpuResourceOutputWithContext(context.Background())
}

func (i GpuResourceArgs) ToGpuResourceOutputWithContext(ctx context.Context) GpuResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GpuResourceOutput)
}

func (i GpuResourceArgs) ToGpuResourcePtrOutput() GpuResourcePtrOutput {
	return i.ToGpuResourcePtrOutputWithContext(context.Background())
}

func (i GpuResourceArgs) ToGpuResourcePtrOutputWithContext(ctx context.Context) GpuResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GpuResourceOutput).ToGpuResourcePtrOutputWithContext(ctx)
}









type GpuResourcePtrInput interface {
	pulumi.Input

	ToGpuResourcePtrOutput() GpuResourcePtrOutput
	ToGpuResourcePtrOutputWithContext(context.Context) GpuResourcePtrOutput
}

type gpuResourcePtrType GpuResourceArgs

func GpuResourcePtr(v *GpuResourceArgs) GpuResourcePtrInput {
	return (*gpuResourcePtrType)(v)
}

func (*gpuResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GpuResource)(nil)).Elem()
}

func (i *gpuResourcePtrType) ToGpuResourcePtrOutput() GpuResourcePtrOutput {
	return i.ToGpuResourcePtrOutputWithContext(context.Background())
}

func (i *gpuResourcePtrType) ToGpuResourcePtrOutputWithContext(ctx context.Context) GpuResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GpuResourcePtrOutput)
}

type GpuResourceOutput struct{ *pulumi.OutputState }

func (GpuResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GpuResource)(nil)).Elem()
}

func (o GpuResourceOutput) ToGpuResourceOutput() GpuResourceOutput {
	return o
}

func (o GpuResourceOutput) ToGpuResourceOutputWithContext(ctx context.Context) GpuResourceOutput {
	return o
}

func (o GpuResourceOutput) ToGpuResourcePtrOutput() GpuResourcePtrOutput {
	return o.ToGpuResourcePtrOutputWithContext(context.Background())
}

func (o GpuResourceOutput) ToGpuResourcePtrOutputWithContext(ctx context.Context) GpuResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GpuResource) *GpuResource {
		return &v
	}).(GpuResourcePtrOutput)
}

func (o GpuResourceOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v GpuResource) int { return v.Count }).(pulumi.IntOutput)
}

func (o GpuResourceOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v GpuResource) string { return v.Sku }).(pulumi.StringOutput)
}

type GpuResourcePtrOutput struct{ *pulumi.OutputState }

func (GpuResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GpuResource)(nil)).Elem()
}

func (o GpuResourcePtrOutput) ToGpuResourcePtrOutput() GpuResourcePtrOutput {
	return o
}

func (o GpuResourcePtrOutput) ToGpuResourcePtrOutputWithContext(ctx context.Context) GpuResourcePtrOutput {
	return o
}

func (o GpuResourcePtrOutput) Elem() GpuResourceOutput {
	return o.ApplyT(func(v *GpuResource) GpuResource {
		if v != nil {
			return *v
		}
		var ret GpuResource
		return ret
	}).(GpuResourceOutput)
}

func (o GpuResourcePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GpuResource) *int {
		if v == nil {
			return nil
		}
		return &v.Count
	}).(pulumi.IntPtrOutput)
}

func (o GpuResourcePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GpuResource) *string {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(pulumi.StringPtrOutput)
}

type GpuResourceResponse struct {
	Count int    `pulumi:"count"`
	Sku   string `pulumi:"sku"`
}





type GpuResourceResponseInput interface {
	pulumi.Input

	ToGpuResourceResponseOutput() GpuResourceResponseOutput
	ToGpuResourceResponseOutputWithContext(context.Context) GpuResourceResponseOutput
}

type GpuResourceResponseArgs struct {
	Count pulumi.IntInput    `pulumi:"count"`
	Sku   pulumi.StringInput `pulumi:"sku"`
}

func (GpuResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GpuResourceResponse)(nil)).Elem()
}

func (i GpuResourceResponseArgs) ToGpuResourceResponseOutput() GpuResourceResponseOutput {
	return i.ToGpuResourceResponseOutputWithContext(context.Background())
}

func (i GpuResourceResponseArgs) ToGpuResourceResponseOutputWithContext(ctx context.Context) GpuResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GpuResourceResponseOutput)
}

func (i GpuResourceResponseArgs) ToGpuResourceResponsePtrOutput() GpuResourceResponsePtrOutput {
	return i.ToGpuResourceResponsePtrOutputWithContext(context.Background())
}

func (i GpuResourceResponseArgs) ToGpuResourceResponsePtrOutputWithContext(ctx context.Context) GpuResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GpuResourceResponseOutput).ToGpuResourceResponsePtrOutputWithContext(ctx)
}









type GpuResourceResponsePtrInput interface {
	pulumi.Input

	ToGpuResourceResponsePtrOutput() GpuResourceResponsePtrOutput
	ToGpuResourceResponsePtrOutputWithContext(context.Context) GpuResourceResponsePtrOutput
}

type gpuResourceResponsePtrType GpuResourceResponseArgs

func GpuResourceResponsePtr(v *GpuResourceResponseArgs) GpuResourceResponsePtrInput {
	return (*gpuResourceResponsePtrType)(v)
}

func (*gpuResourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GpuResourceResponse)(nil)).Elem()
}

func (i *gpuResourceResponsePtrType) ToGpuResourceResponsePtrOutput() GpuResourceResponsePtrOutput {
	return i.ToGpuResourceResponsePtrOutputWithContext(context.Background())
}

func (i *gpuResourceResponsePtrType) ToGpuResourceResponsePtrOutputWithContext(ctx context.Context) GpuResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GpuResourceResponsePtrOutput)
}

type GpuResourceResponseOutput struct{ *pulumi.OutputState }

func (GpuResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GpuResourceResponse)(nil)).Elem()
}

func (o GpuResourceResponseOutput) ToGpuResourceResponseOutput() GpuResourceResponseOutput {
	return o
}

func (o GpuResourceResponseOutput) ToGpuResourceResponseOutputWithContext(ctx context.Context) GpuResourceResponseOutput {
	return o
}

func (o GpuResourceResponseOutput) ToGpuResourceResponsePtrOutput() GpuResourceResponsePtrOutput {
	return o.ToGpuResourceResponsePtrOutputWithContext(context.Background())
}

func (o GpuResourceResponseOutput) ToGpuResourceResponsePtrOutputWithContext(ctx context.Context) GpuResourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GpuResourceResponse) *GpuResourceResponse {
		return &v
	}).(GpuResourceResponsePtrOutput)
}

func (o GpuResourceResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v GpuResourceResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o GpuResourceResponseOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v GpuResourceResponse) string { return v.Sku }).(pulumi.StringOutput)
}

type GpuResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (GpuResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GpuResourceResponse)(nil)).Elem()
}

func (o GpuResourceResponsePtrOutput) ToGpuResourceResponsePtrOutput() GpuResourceResponsePtrOutput {
	return o
}

func (o GpuResourceResponsePtrOutput) ToGpuResourceResponsePtrOutputWithContext(ctx context.Context) GpuResourceResponsePtrOutput {
	return o
}

func (o GpuResourceResponsePtrOutput) Elem() GpuResourceResponseOutput {
	return o.ApplyT(func(v *GpuResourceResponse) GpuResourceResponse {
		if v != nil {
			return *v
		}
		var ret GpuResourceResponse
		return ret
	}).(GpuResourceResponseOutput)
}

func (o GpuResourceResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GpuResourceResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Count
	}).(pulumi.IntPtrOutput)
}

func (o GpuResourceResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GpuResourceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(pulumi.StringPtrOutput)
}

type HttpHeader struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type HttpHeaderInput interface {
	pulumi.Input

	ToHttpHeaderOutput() HttpHeaderOutput
	ToHttpHeaderOutputWithContext(context.Context) HttpHeaderOutput
}

type HttpHeaderArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (HttpHeaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpHeader)(nil)).Elem()
}

func (i HttpHeaderArgs) ToHttpHeaderOutput() HttpHeaderOutput {
	return i.ToHttpHeaderOutputWithContext(context.Background())
}

func (i HttpHeaderArgs) ToHttpHeaderOutputWithContext(ctx context.Context) HttpHeaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpHeaderOutput)
}





type HttpHeaderArrayInput interface {
	pulumi.Input

	ToHttpHeaderArrayOutput() HttpHeaderArrayOutput
	ToHttpHeaderArrayOutputWithContext(context.Context) HttpHeaderArrayOutput
}

type HttpHeaderArray []HttpHeaderInput

func (HttpHeaderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpHeader)(nil)).Elem()
}

func (i HttpHeaderArray) ToHttpHeaderArrayOutput() HttpHeaderArrayOutput {
	return i.ToHttpHeaderArrayOutputWithContext(context.Background())
}

func (i HttpHeaderArray) ToHttpHeaderArrayOutputWithContext(ctx context.Context) HttpHeaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpHeaderArrayOutput)
}

type HttpHeaderOutput struct{ *pulumi.OutputState }

func (HttpHeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpHeader)(nil)).Elem()
}

func (o HttpHeaderOutput) ToHttpHeaderOutput() HttpHeaderOutput {
	return o
}

func (o HttpHeaderOutput) ToHttpHeaderOutputWithContext(ctx context.Context) HttpHeaderOutput {
	return o
}

func (o HttpHeaderOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpHeader) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HttpHeaderOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpHeader) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HttpHeaderArrayOutput struct{ *pulumi.OutputState }

func (HttpHeaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpHeader)(nil)).Elem()
}

func (o HttpHeaderArrayOutput) ToHttpHeaderArrayOutput() HttpHeaderArrayOutput {
	return o
}

func (o HttpHeaderArrayOutput) ToHttpHeaderArrayOutputWithContext(ctx context.Context) HttpHeaderArrayOutput {
	return o
}

func (o HttpHeaderArrayOutput) Index(i pulumi.IntInput) HttpHeaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpHeader {
		return vs[0].([]HttpHeader)[vs[1].(int)]
	}).(HttpHeaderOutput)
}

type HttpHeaderResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type HttpHeaderResponseInput interface {
	pulumi.Input

	ToHttpHeaderResponseOutput() HttpHeaderResponseOutput
	ToHttpHeaderResponseOutputWithContext(context.Context) HttpHeaderResponseOutput
}

type HttpHeaderResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (HttpHeaderResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpHeaderResponse)(nil)).Elem()
}

func (i HttpHeaderResponseArgs) ToHttpHeaderResponseOutput() HttpHeaderResponseOutput {
	return i.ToHttpHeaderResponseOutputWithContext(context.Background())
}

func (i HttpHeaderResponseArgs) ToHttpHeaderResponseOutputWithContext(ctx context.Context) HttpHeaderResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpHeaderResponseOutput)
}





type HttpHeaderResponseArrayInput interface {
	pulumi.Input

	ToHttpHeaderResponseArrayOutput() HttpHeaderResponseArrayOutput
	ToHttpHeaderResponseArrayOutputWithContext(context.Context) HttpHeaderResponseArrayOutput
}

type HttpHeaderResponseArray []HttpHeaderResponseInput

func (HttpHeaderResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpHeaderResponse)(nil)).Elem()
}

func (i HttpHeaderResponseArray) ToHttpHeaderResponseArrayOutput() HttpHeaderResponseArrayOutput {
	return i.ToHttpHeaderResponseArrayOutputWithContext(context.Background())
}

func (i HttpHeaderResponseArray) ToHttpHeaderResponseArrayOutputWithContext(ctx context.Context) HttpHeaderResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpHeaderResponseArrayOutput)
}

type HttpHeaderResponseOutput struct{ *pulumi.OutputState }

func (HttpHeaderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpHeaderResponse)(nil)).Elem()
}

func (o HttpHeaderResponseOutput) ToHttpHeaderResponseOutput() HttpHeaderResponseOutput {
	return o
}

func (o HttpHeaderResponseOutput) ToHttpHeaderResponseOutputWithContext(ctx context.Context) HttpHeaderResponseOutput {
	return o
}

func (o HttpHeaderResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpHeaderResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HttpHeaderResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpHeaderResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HttpHeaderResponseArrayOutput struct{ *pulumi.OutputState }

func (HttpHeaderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpHeaderResponse)(nil)).Elem()
}

func (o HttpHeaderResponseArrayOutput) ToHttpHeaderResponseArrayOutput() HttpHeaderResponseArrayOutput {
	return o
}

func (o HttpHeaderResponseArrayOutput) ToHttpHeaderResponseArrayOutputWithContext(ctx context.Context) HttpHeaderResponseArrayOutput {
	return o
}

func (o HttpHeaderResponseArrayOutput) Index(i pulumi.IntInput) HttpHeaderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpHeaderResponse {
		return vs[0].([]HttpHeaderResponse)[vs[1].(int)]
	}).(HttpHeaderResponseOutput)
}

type ImageRegistryCredential struct {
	Identity    *string `pulumi:"identity"`
	IdentityUrl *string `pulumi:"identityUrl"`
	Password    *string `pulumi:"password"`
	Server      string  `pulumi:"server"`
	Username    string  `pulumi:"username"`
}





type ImageRegistryCredentialInput interface {
	pulumi.Input

	ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput
	ToImageRegistryCredentialOutputWithContext(context.Context) ImageRegistryCredentialOutput
}

type ImageRegistryCredentialArgs struct {
	Identity    pulumi.StringPtrInput `pulumi:"identity"`
	IdentityUrl pulumi.StringPtrInput `pulumi:"identityUrl"`
	Password    pulumi.StringPtrInput `pulumi:"password"`
	Server      pulumi.StringInput    `pulumi:"server"`
	Username    pulumi.StringInput    `pulumi:"username"`
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

func (o ImageRegistryCredentialOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredential) *string { return v.Identity }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialOutput) IdentityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredential) *string { return v.IdentityUrl }).(pulumi.StringPtrOutput)
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
	Identity    *string `pulumi:"identity"`
	IdentityUrl *string `pulumi:"identityUrl"`
	Password    *string `pulumi:"password"`
	Server      string  `pulumi:"server"`
	Username    string  `pulumi:"username"`
}





type ImageRegistryCredentialResponseInput interface {
	pulumi.Input

	ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput
	ToImageRegistryCredentialResponseOutputWithContext(context.Context) ImageRegistryCredentialResponseOutput
}

type ImageRegistryCredentialResponseArgs struct {
	Identity    pulumi.StringPtrInput `pulumi:"identity"`
	IdentityUrl pulumi.StringPtrInput `pulumi:"identityUrl"`
	Password    pulumi.StringPtrInput `pulumi:"password"`
	Server      pulumi.StringInput    `pulumi:"server"`
	Username    pulumi.StringInput    `pulumi:"username"`
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

func (o ImageRegistryCredentialResponseOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) *string { return v.Identity }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponseOutput) IdentityUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) *string { return v.IdentityUrl }).(pulumi.StringPtrOutput)
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

type InitContainerDefinition struct {
	Command              []string              `pulumi:"command"`
	EnvironmentVariables []EnvironmentVariable `pulumi:"environmentVariables"`
	Image                *string               `pulumi:"image"`
	Name                 string                `pulumi:"name"`
	VolumeMounts         []VolumeMount         `pulumi:"volumeMounts"`
}





type InitContainerDefinitionInput interface {
	pulumi.Input

	ToInitContainerDefinitionOutput() InitContainerDefinitionOutput
	ToInitContainerDefinitionOutputWithContext(context.Context) InitContainerDefinitionOutput
}

type InitContainerDefinitionArgs struct {
	Command              pulumi.StringArrayInput       `pulumi:"command"`
	EnvironmentVariables EnvironmentVariableArrayInput `pulumi:"environmentVariables"`
	Image                pulumi.StringPtrInput         `pulumi:"image"`
	Name                 pulumi.StringInput            `pulumi:"name"`
	VolumeMounts         VolumeMountArrayInput         `pulumi:"volumeMounts"`
}

func (InitContainerDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InitContainerDefinition)(nil)).Elem()
}

func (i InitContainerDefinitionArgs) ToInitContainerDefinitionOutput() InitContainerDefinitionOutput {
	return i.ToInitContainerDefinitionOutputWithContext(context.Background())
}

func (i InitContainerDefinitionArgs) ToInitContainerDefinitionOutputWithContext(ctx context.Context) InitContainerDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InitContainerDefinitionOutput)
}





type InitContainerDefinitionArrayInput interface {
	pulumi.Input

	ToInitContainerDefinitionArrayOutput() InitContainerDefinitionArrayOutput
	ToInitContainerDefinitionArrayOutputWithContext(context.Context) InitContainerDefinitionArrayOutput
}

type InitContainerDefinitionArray []InitContainerDefinitionInput

func (InitContainerDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InitContainerDefinition)(nil)).Elem()
}

func (i InitContainerDefinitionArray) ToInitContainerDefinitionArrayOutput() InitContainerDefinitionArrayOutput {
	return i.ToInitContainerDefinitionArrayOutputWithContext(context.Background())
}

func (i InitContainerDefinitionArray) ToInitContainerDefinitionArrayOutputWithContext(ctx context.Context) InitContainerDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InitContainerDefinitionArrayOutput)
}

type InitContainerDefinitionOutput struct{ *pulumi.OutputState }

func (InitContainerDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InitContainerDefinition)(nil)).Elem()
}

func (o InitContainerDefinitionOutput) ToInitContainerDefinitionOutput() InitContainerDefinitionOutput {
	return o
}

func (o InitContainerDefinitionOutput) ToInitContainerDefinitionOutputWithContext(ctx context.Context) InitContainerDefinitionOutput {
	return o
}

func (o InitContainerDefinitionOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InitContainerDefinition) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o InitContainerDefinitionOutput) EnvironmentVariables() EnvironmentVariableArrayOutput {
	return o.ApplyT(func(v InitContainerDefinition) []EnvironmentVariable { return v.EnvironmentVariables }).(EnvironmentVariableArrayOutput)
}

func (o InitContainerDefinitionOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InitContainerDefinition) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o InitContainerDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InitContainerDefinition) string { return v.Name }).(pulumi.StringOutput)
}

func (o InitContainerDefinitionOutput) VolumeMounts() VolumeMountArrayOutput {
	return o.ApplyT(func(v InitContainerDefinition) []VolumeMount { return v.VolumeMounts }).(VolumeMountArrayOutput)
}

type InitContainerDefinitionArrayOutput struct{ *pulumi.OutputState }

func (InitContainerDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InitContainerDefinition)(nil)).Elem()
}

func (o InitContainerDefinitionArrayOutput) ToInitContainerDefinitionArrayOutput() InitContainerDefinitionArrayOutput {
	return o
}

func (o InitContainerDefinitionArrayOutput) ToInitContainerDefinitionArrayOutputWithContext(ctx context.Context) InitContainerDefinitionArrayOutput {
	return o
}

func (o InitContainerDefinitionArrayOutput) Index(i pulumi.IntInput) InitContainerDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InitContainerDefinition {
		return vs[0].([]InitContainerDefinition)[vs[1].(int)]
	}).(InitContainerDefinitionOutput)
}

type InitContainerDefinitionResponse struct {
	Command              []string                                              `pulumi:"command"`
	EnvironmentVariables []EnvironmentVariableResponse                         `pulumi:"environmentVariables"`
	Image                *string                                               `pulumi:"image"`
	InstanceView         InitContainerPropertiesDefinitionResponseInstanceView `pulumi:"instanceView"`
	Name                 string                                                `pulumi:"name"`
	VolumeMounts         []VolumeMountResponse                                 `pulumi:"volumeMounts"`
}





type InitContainerDefinitionResponseInput interface {
	pulumi.Input

	ToInitContainerDefinitionResponseOutput() InitContainerDefinitionResponseOutput
	ToInitContainerDefinitionResponseOutputWithContext(context.Context) InitContainerDefinitionResponseOutput
}

type InitContainerDefinitionResponseArgs struct {
	Command              pulumi.StringArrayInput                                    `pulumi:"command"`
	EnvironmentVariables EnvironmentVariableResponseArrayInput                      `pulumi:"environmentVariables"`
	Image                pulumi.StringPtrInput                                      `pulumi:"image"`
	InstanceView         InitContainerPropertiesDefinitionResponseInstanceViewInput `pulumi:"instanceView"`
	Name                 pulumi.StringInput                                         `pulumi:"name"`
	VolumeMounts         VolumeMountResponseArrayInput                              `pulumi:"volumeMounts"`
}

func (InitContainerDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InitContainerDefinitionResponse)(nil)).Elem()
}

func (i InitContainerDefinitionResponseArgs) ToInitContainerDefinitionResponseOutput() InitContainerDefinitionResponseOutput {
	return i.ToInitContainerDefinitionResponseOutputWithContext(context.Background())
}

func (i InitContainerDefinitionResponseArgs) ToInitContainerDefinitionResponseOutputWithContext(ctx context.Context) InitContainerDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InitContainerDefinitionResponseOutput)
}





type InitContainerDefinitionResponseArrayInput interface {
	pulumi.Input

	ToInitContainerDefinitionResponseArrayOutput() InitContainerDefinitionResponseArrayOutput
	ToInitContainerDefinitionResponseArrayOutputWithContext(context.Context) InitContainerDefinitionResponseArrayOutput
}

type InitContainerDefinitionResponseArray []InitContainerDefinitionResponseInput

func (InitContainerDefinitionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InitContainerDefinitionResponse)(nil)).Elem()
}

func (i InitContainerDefinitionResponseArray) ToInitContainerDefinitionResponseArrayOutput() InitContainerDefinitionResponseArrayOutput {
	return i.ToInitContainerDefinitionResponseArrayOutputWithContext(context.Background())
}

func (i InitContainerDefinitionResponseArray) ToInitContainerDefinitionResponseArrayOutputWithContext(ctx context.Context) InitContainerDefinitionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InitContainerDefinitionResponseArrayOutput)
}

type InitContainerDefinitionResponseOutput struct{ *pulumi.OutputState }

func (InitContainerDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InitContainerDefinitionResponse)(nil)).Elem()
}

func (o InitContainerDefinitionResponseOutput) ToInitContainerDefinitionResponseOutput() InitContainerDefinitionResponseOutput {
	return o
}

func (o InitContainerDefinitionResponseOutput) ToInitContainerDefinitionResponseOutputWithContext(ctx context.Context) InitContainerDefinitionResponseOutput {
	return o
}

func (o InitContainerDefinitionResponseOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InitContainerDefinitionResponse) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o InitContainerDefinitionResponseOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v InitContainerDefinitionResponse) []EnvironmentVariableResponse { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

func (o InitContainerDefinitionResponseOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InitContainerDefinitionResponse) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o InitContainerDefinitionResponseOutput) InstanceView() InitContainerPropertiesDefinitionResponseInstanceViewOutput {
	return o.ApplyT(func(v InitContainerDefinitionResponse) InitContainerPropertiesDefinitionResponseInstanceView {
		return v.InstanceView
	}).(InitContainerPropertiesDefinitionResponseInstanceViewOutput)
}

func (o InitContainerDefinitionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InitContainerDefinitionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o InitContainerDefinitionResponseOutput) VolumeMounts() VolumeMountResponseArrayOutput {
	return o.ApplyT(func(v InitContainerDefinitionResponse) []VolumeMountResponse { return v.VolumeMounts }).(VolumeMountResponseArrayOutput)
}

type InitContainerDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (InitContainerDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InitContainerDefinitionResponse)(nil)).Elem()
}

func (o InitContainerDefinitionResponseArrayOutput) ToInitContainerDefinitionResponseArrayOutput() InitContainerDefinitionResponseArrayOutput {
	return o
}

func (o InitContainerDefinitionResponseArrayOutput) ToInitContainerDefinitionResponseArrayOutputWithContext(ctx context.Context) InitContainerDefinitionResponseArrayOutput {
	return o
}

func (o InitContainerDefinitionResponseArrayOutput) Index(i pulumi.IntInput) InitContainerDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InitContainerDefinitionResponse {
		return vs[0].([]InitContainerDefinitionResponse)[vs[1].(int)]
	}).(InitContainerDefinitionResponseOutput)
}

type InitContainerPropertiesDefinitionResponseInstanceView struct {
	CurrentState  ContainerStateResponse `pulumi:"currentState"`
	Events        []EventResponse        `pulumi:"events"`
	PreviousState ContainerStateResponse `pulumi:"previousState"`
	RestartCount  int                    `pulumi:"restartCount"`
}





type InitContainerPropertiesDefinitionResponseInstanceViewInput interface {
	pulumi.Input

	ToInitContainerPropertiesDefinitionResponseInstanceViewOutput() InitContainerPropertiesDefinitionResponseInstanceViewOutput
	ToInitContainerPropertiesDefinitionResponseInstanceViewOutputWithContext(context.Context) InitContainerPropertiesDefinitionResponseInstanceViewOutput
}

type InitContainerPropertiesDefinitionResponseInstanceViewArgs struct {
	CurrentState  ContainerStateResponseInput `pulumi:"currentState"`
	Events        EventResponseArrayInput     `pulumi:"events"`
	PreviousState ContainerStateResponseInput `pulumi:"previousState"`
	RestartCount  pulumi.IntInput             `pulumi:"restartCount"`
}

func (InitContainerPropertiesDefinitionResponseInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InitContainerPropertiesDefinitionResponseInstanceView)(nil)).Elem()
}

func (i InitContainerPropertiesDefinitionResponseInstanceViewArgs) ToInitContainerPropertiesDefinitionResponseInstanceViewOutput() InitContainerPropertiesDefinitionResponseInstanceViewOutput {
	return i.ToInitContainerPropertiesDefinitionResponseInstanceViewOutputWithContext(context.Background())
}

func (i InitContainerPropertiesDefinitionResponseInstanceViewArgs) ToInitContainerPropertiesDefinitionResponseInstanceViewOutputWithContext(ctx context.Context) InitContainerPropertiesDefinitionResponseInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InitContainerPropertiesDefinitionResponseInstanceViewOutput)
}

type InitContainerPropertiesDefinitionResponseInstanceViewOutput struct{ *pulumi.OutputState }

func (InitContainerPropertiesDefinitionResponseInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InitContainerPropertiesDefinitionResponseInstanceView)(nil)).Elem()
}

func (o InitContainerPropertiesDefinitionResponseInstanceViewOutput) ToInitContainerPropertiesDefinitionResponseInstanceViewOutput() InitContainerPropertiesDefinitionResponseInstanceViewOutput {
	return o
}

func (o InitContainerPropertiesDefinitionResponseInstanceViewOutput) ToInitContainerPropertiesDefinitionResponseInstanceViewOutputWithContext(ctx context.Context) InitContainerPropertiesDefinitionResponseInstanceViewOutput {
	return o
}

func (o InitContainerPropertiesDefinitionResponseInstanceViewOutput) CurrentState() ContainerStateResponseOutput {
	return o.ApplyT(func(v InitContainerPropertiesDefinitionResponseInstanceView) ContainerStateResponse {
		return v.CurrentState
	}).(ContainerStateResponseOutput)
}

func (o InitContainerPropertiesDefinitionResponseInstanceViewOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v InitContainerPropertiesDefinitionResponseInstanceView) []EventResponse { return v.Events }).(EventResponseArrayOutput)
}

func (o InitContainerPropertiesDefinitionResponseInstanceViewOutput) PreviousState() ContainerStateResponseOutput {
	return o.ApplyT(func(v InitContainerPropertiesDefinitionResponseInstanceView) ContainerStateResponse {
		return v.PreviousState
	}).(ContainerStateResponseOutput)
}

func (o InitContainerPropertiesDefinitionResponseInstanceViewOutput) RestartCount() pulumi.IntOutput {
	return o.ApplyT(func(v InitContainerPropertiesDefinitionResponseInstanceView) int { return v.RestartCount }).(pulumi.IntOutput)
}

type IpAddress struct {
	DnsNameLabel *string `pulumi:"dnsNameLabel"`
	Ip           *string `pulumi:"ip"`
	Ports        []Port  `pulumi:"ports"`
	Type         string  `pulumi:"type"`
}





type IpAddressInput interface {
	pulumi.Input

	ToIpAddressOutput() IpAddressOutput
	ToIpAddressOutputWithContext(context.Context) IpAddressOutput
}

type IpAddressArgs struct {
	DnsNameLabel pulumi.StringPtrInput `pulumi:"dnsNameLabel"`
	Ip           pulumi.StringPtrInput `pulumi:"ip"`
	Ports        PortArrayInput        `pulumi:"ports"`
	Type         pulumi.StringInput    `pulumi:"type"`
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

func (o IpAddressOutput) DnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddress) *string { return v.DnsNameLabel }).(pulumi.StringPtrOutput)
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

func (o IpAddressPtrOutput) DnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAddress) *string {
		if v == nil {
			return nil
		}
		return v.DnsNameLabel
	}).(pulumi.StringPtrOutput)
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
	DnsNameLabel *string        `pulumi:"dnsNameLabel"`
	Fqdn         string         `pulumi:"fqdn"`
	Ip           *string        `pulumi:"ip"`
	Ports        []PortResponse `pulumi:"ports"`
	Type         string         `pulumi:"type"`
}





type IpAddressResponseInput interface {
	pulumi.Input

	ToIpAddressResponseOutput() IpAddressResponseOutput
	ToIpAddressResponseOutputWithContext(context.Context) IpAddressResponseOutput
}

type IpAddressResponseArgs struct {
	DnsNameLabel pulumi.StringPtrInput  `pulumi:"dnsNameLabel"`
	Fqdn         pulumi.StringInput     `pulumi:"fqdn"`
	Ip           pulumi.StringPtrInput  `pulumi:"ip"`
	Ports        PortResponseArrayInput `pulumi:"ports"`
	Type         pulumi.StringInput     `pulumi:"type"`
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

func (o IpAddressResponseOutput) DnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressResponse) *string { return v.DnsNameLabel }).(pulumi.StringPtrOutput)
}

func (o IpAddressResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v IpAddressResponse) string { return v.Fqdn }).(pulumi.StringOutput)
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

func (o IpAddressResponsePtrOutput) DnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o IpAddressResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Fqdn
	}).(pulumi.StringPtrOutput)
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

type LogAnalytics struct {
	LogType             *string           `pulumi:"logType"`
	Metadata            map[string]string `pulumi:"metadata"`
	WorkspaceId         string            `pulumi:"workspaceId"`
	WorkspaceKey        string            `pulumi:"workspaceKey"`
	WorkspaceResourceId *string           `pulumi:"workspaceResourceId"`
}





type LogAnalyticsInput interface {
	pulumi.Input

	ToLogAnalyticsOutput() LogAnalyticsOutput
	ToLogAnalyticsOutputWithContext(context.Context) LogAnalyticsOutput
}

type LogAnalyticsArgs struct {
	LogType             pulumi.StringPtrInput `pulumi:"logType"`
	Metadata            pulumi.StringMapInput `pulumi:"metadata"`
	WorkspaceId         pulumi.StringInput    `pulumi:"workspaceId"`
	WorkspaceKey        pulumi.StringInput    `pulumi:"workspaceKey"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (LogAnalyticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalytics)(nil)).Elem()
}

func (i LogAnalyticsArgs) ToLogAnalyticsOutput() LogAnalyticsOutput {
	return i.ToLogAnalyticsOutputWithContext(context.Background())
}

func (i LogAnalyticsArgs) ToLogAnalyticsOutputWithContext(ctx context.Context) LogAnalyticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsOutput)
}

func (i LogAnalyticsArgs) ToLogAnalyticsPtrOutput() LogAnalyticsPtrOutput {
	return i.ToLogAnalyticsPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsArgs) ToLogAnalyticsPtrOutputWithContext(ctx context.Context) LogAnalyticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsOutput).ToLogAnalyticsPtrOutputWithContext(ctx)
}









type LogAnalyticsPtrInput interface {
	pulumi.Input

	ToLogAnalyticsPtrOutput() LogAnalyticsPtrOutput
	ToLogAnalyticsPtrOutputWithContext(context.Context) LogAnalyticsPtrOutput
}

type logAnalyticsPtrType LogAnalyticsArgs

func LogAnalyticsPtr(v *LogAnalyticsArgs) LogAnalyticsPtrInput {
	return (*logAnalyticsPtrType)(v)
}

func (*logAnalyticsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalytics)(nil)).Elem()
}

func (i *logAnalyticsPtrType) ToLogAnalyticsPtrOutput() LogAnalyticsPtrOutput {
	return i.ToLogAnalyticsPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsPtrType) ToLogAnalyticsPtrOutputWithContext(ctx context.Context) LogAnalyticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsPtrOutput)
}

type LogAnalyticsOutput struct{ *pulumi.OutputState }

func (LogAnalyticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalytics)(nil)).Elem()
}

func (o LogAnalyticsOutput) ToLogAnalyticsOutput() LogAnalyticsOutput {
	return o
}

func (o LogAnalyticsOutput) ToLogAnalyticsOutputWithContext(ctx context.Context) LogAnalyticsOutput {
	return o
}

func (o LogAnalyticsOutput) ToLogAnalyticsPtrOutput() LogAnalyticsPtrOutput {
	return o.ToLogAnalyticsPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsOutput) ToLogAnalyticsPtrOutputWithContext(ctx context.Context) LogAnalyticsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalytics) *LogAnalytics {
		return &v
	}).(LogAnalyticsPtrOutput)
}

func (o LogAnalyticsOutput) LogType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalytics) *string { return v.LogType }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LogAnalytics) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LogAnalyticsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalytics) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o LogAnalyticsOutput) WorkspaceKey() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalytics) string { return v.WorkspaceKey }).(pulumi.StringOutput)
}

func (o LogAnalyticsOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalytics) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalytics)(nil)).Elem()
}

func (o LogAnalyticsPtrOutput) ToLogAnalyticsPtrOutput() LogAnalyticsPtrOutput {
	return o
}

func (o LogAnalyticsPtrOutput) ToLogAnalyticsPtrOutputWithContext(ctx context.Context) LogAnalyticsPtrOutput {
	return o
}

func (o LogAnalyticsPtrOutput) Elem() LogAnalyticsOutput {
	return o.ApplyT(func(v *LogAnalytics) LogAnalytics {
		if v != nil {
			return *v
		}
		var ret LogAnalytics
		return ret
	}).(LogAnalyticsOutput)
}

func (o LogAnalyticsPtrOutput) LogType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalytics) *string {
		if v == nil {
			return nil
		}
		return v.LogType
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsPtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LogAnalytics) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

func (o LogAnalyticsPtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalytics) *string {
		if v == nil {
			return nil
		}
		return &v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsPtrOutput) WorkspaceKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalytics) *string {
		if v == nil {
			return nil
		}
		return &v.WorkspaceKey
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsPtrOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalytics) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceResourceId
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsResponse struct {
	LogType             *string           `pulumi:"logType"`
	Metadata            map[string]string `pulumi:"metadata"`
	WorkspaceId         string            `pulumi:"workspaceId"`
	WorkspaceKey        string            `pulumi:"workspaceKey"`
	WorkspaceResourceId *string           `pulumi:"workspaceResourceId"`
}





type LogAnalyticsResponseInput interface {
	pulumi.Input

	ToLogAnalyticsResponseOutput() LogAnalyticsResponseOutput
	ToLogAnalyticsResponseOutputWithContext(context.Context) LogAnalyticsResponseOutput
}

type LogAnalyticsResponseArgs struct {
	LogType             pulumi.StringPtrInput `pulumi:"logType"`
	Metadata            pulumi.StringMapInput `pulumi:"metadata"`
	WorkspaceId         pulumi.StringInput    `pulumi:"workspaceId"`
	WorkspaceKey        pulumi.StringInput    `pulumi:"workspaceKey"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (LogAnalyticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsResponse)(nil)).Elem()
}

func (i LogAnalyticsResponseArgs) ToLogAnalyticsResponseOutput() LogAnalyticsResponseOutput {
	return i.ToLogAnalyticsResponseOutputWithContext(context.Background())
}

func (i LogAnalyticsResponseArgs) ToLogAnalyticsResponseOutputWithContext(ctx context.Context) LogAnalyticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsResponseOutput)
}

func (i LogAnalyticsResponseArgs) ToLogAnalyticsResponsePtrOutput() LogAnalyticsResponsePtrOutput {
	return i.ToLogAnalyticsResponsePtrOutputWithContext(context.Background())
}

func (i LogAnalyticsResponseArgs) ToLogAnalyticsResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsResponseOutput).ToLogAnalyticsResponsePtrOutputWithContext(ctx)
}









type LogAnalyticsResponsePtrInput interface {
	pulumi.Input

	ToLogAnalyticsResponsePtrOutput() LogAnalyticsResponsePtrOutput
	ToLogAnalyticsResponsePtrOutputWithContext(context.Context) LogAnalyticsResponsePtrOutput
}

type logAnalyticsResponsePtrType LogAnalyticsResponseArgs

func LogAnalyticsResponsePtr(v *LogAnalyticsResponseArgs) LogAnalyticsResponsePtrInput {
	return (*logAnalyticsResponsePtrType)(v)
}

func (*logAnalyticsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsResponse)(nil)).Elem()
}

func (i *logAnalyticsResponsePtrType) ToLogAnalyticsResponsePtrOutput() LogAnalyticsResponsePtrOutput {
	return i.ToLogAnalyticsResponsePtrOutputWithContext(context.Background())
}

func (i *logAnalyticsResponsePtrType) ToLogAnalyticsResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsResponsePtrOutput)
}

type LogAnalyticsResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsResponse)(nil)).Elem()
}

func (o LogAnalyticsResponseOutput) ToLogAnalyticsResponseOutput() LogAnalyticsResponseOutput {
	return o
}

func (o LogAnalyticsResponseOutput) ToLogAnalyticsResponseOutputWithContext(ctx context.Context) LogAnalyticsResponseOutput {
	return o
}

func (o LogAnalyticsResponseOutput) ToLogAnalyticsResponsePtrOutput() LogAnalyticsResponsePtrOutput {
	return o.ToLogAnalyticsResponsePtrOutputWithContext(context.Background())
}

func (o LogAnalyticsResponseOutput) ToLogAnalyticsResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsResponse) *LogAnalyticsResponse {
		return &v
	}).(LogAnalyticsResponsePtrOutput)
}

func (o LogAnalyticsResponseOutput) LogType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsResponse) *string { return v.LogType }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsResponseOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LogAnalyticsResponse) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LogAnalyticsResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o LogAnalyticsResponseOutput) WorkspaceKey() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsResponse) string { return v.WorkspaceKey }).(pulumi.StringOutput)
}

func (o LogAnalyticsResponseOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsResponse) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsResponsePtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsResponse)(nil)).Elem()
}

func (o LogAnalyticsResponsePtrOutput) ToLogAnalyticsResponsePtrOutput() LogAnalyticsResponsePtrOutput {
	return o
}

func (o LogAnalyticsResponsePtrOutput) ToLogAnalyticsResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsResponsePtrOutput {
	return o
}

func (o LogAnalyticsResponsePtrOutput) Elem() LogAnalyticsResponseOutput {
	return o.ApplyT(func(v *LogAnalyticsResponse) LogAnalyticsResponse {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsResponse
		return ret
	}).(LogAnalyticsResponseOutput)
}

func (o LogAnalyticsResponsePtrOutput) LogType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LogType
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsResponsePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LogAnalyticsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

func (o LogAnalyticsResponsePtrOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.WorkspaceId
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsResponsePtrOutput) WorkspaceKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.WorkspaceKey
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsResponsePtrOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceResourceId
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
	Cpu        *float64     `pulumi:"cpu"`
	Gpu        *GpuResource `pulumi:"gpu"`
	MemoryInGB *float64     `pulumi:"memoryInGB"`
}





type ResourceLimitsInput interface {
	pulumi.Input

	ToResourceLimitsOutput() ResourceLimitsOutput
	ToResourceLimitsOutputWithContext(context.Context) ResourceLimitsOutput
}

type ResourceLimitsArgs struct {
	Cpu        pulumi.Float64PtrInput `pulumi:"cpu"`
	Gpu        GpuResourcePtrInput    `pulumi:"gpu"`
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

func (o ResourceLimitsOutput) Gpu() GpuResourcePtrOutput {
	return o.ApplyT(func(v ResourceLimits) *GpuResource { return v.Gpu }).(GpuResourcePtrOutput)
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

func (o ResourceLimitsPtrOutput) Gpu() GpuResourcePtrOutput {
	return o.ApplyT(func(v *ResourceLimits) *GpuResource {
		if v == nil {
			return nil
		}
		return v.Gpu
	}).(GpuResourcePtrOutput)
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
	Cpu        *float64             `pulumi:"cpu"`
	Gpu        *GpuResourceResponse `pulumi:"gpu"`
	MemoryInGB *float64             `pulumi:"memoryInGB"`
}





type ResourceLimitsResponseInput interface {
	pulumi.Input

	ToResourceLimitsResponseOutput() ResourceLimitsResponseOutput
	ToResourceLimitsResponseOutputWithContext(context.Context) ResourceLimitsResponseOutput
}

type ResourceLimitsResponseArgs struct {
	Cpu        pulumi.Float64PtrInput      `pulumi:"cpu"`
	Gpu        GpuResourceResponsePtrInput `pulumi:"gpu"`
	MemoryInGB pulumi.Float64PtrInput      `pulumi:"memoryInGB"`
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

func (o ResourceLimitsResponseOutput) Gpu() GpuResourceResponsePtrOutput {
	return o.ApplyT(func(v ResourceLimitsResponse) *GpuResourceResponse { return v.Gpu }).(GpuResourceResponsePtrOutput)
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

func (o ResourceLimitsResponsePtrOutput) Gpu() GpuResourceResponsePtrOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) *GpuResourceResponse {
		if v == nil {
			return nil
		}
		return v.Gpu
	}).(GpuResourceResponsePtrOutput)
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
	Cpu        float64      `pulumi:"cpu"`
	Gpu        *GpuResource `pulumi:"gpu"`
	MemoryInGB float64      `pulumi:"memoryInGB"`
}





type ResourceRequestsInput interface {
	pulumi.Input

	ToResourceRequestsOutput() ResourceRequestsOutput
	ToResourceRequestsOutputWithContext(context.Context) ResourceRequestsOutput
}

type ResourceRequestsArgs struct {
	Cpu        pulumi.Float64Input `pulumi:"cpu"`
	Gpu        GpuResourcePtrInput `pulumi:"gpu"`
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

func (o ResourceRequestsOutput) Gpu() GpuResourcePtrOutput {
	return o.ApplyT(func(v ResourceRequests) *GpuResource { return v.Gpu }).(GpuResourcePtrOutput)
}

func (o ResourceRequestsOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequests) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

type ResourceRequestsResponse struct {
	Cpu        float64              `pulumi:"cpu"`
	Gpu        *GpuResourceResponse `pulumi:"gpu"`
	MemoryInGB float64              `pulumi:"memoryInGB"`
}





type ResourceRequestsResponseInput interface {
	pulumi.Input

	ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput
	ToResourceRequestsResponseOutputWithContext(context.Context) ResourceRequestsResponseOutput
}

type ResourceRequestsResponseArgs struct {
	Cpu        pulumi.Float64Input         `pulumi:"cpu"`
	Gpu        GpuResourceResponsePtrInput `pulumi:"gpu"`
	MemoryInGB pulumi.Float64Input         `pulumi:"memoryInGB"`
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

func (o ResourceRequestsResponseOutput) Gpu() GpuResourceResponsePtrOutput {
	return o.ApplyT(func(v ResourceRequestsResponse) *GpuResourceResponse { return v.Gpu }).(GpuResourceResponsePtrOutput)
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
	AzureFile *AzureFileVolume  `pulumi:"azureFile"`
	EmptyDir  interface{}       `pulumi:"emptyDir"`
	GitRepo   *GitRepoVolume    `pulumi:"gitRepo"`
	Name      string            `pulumi:"name"`
	Secret    map[string]string `pulumi:"secret"`
}





type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(context.Context) VolumeOutput
}

type VolumeArgs struct {
	AzureFile AzureFileVolumePtrInput `pulumi:"azureFile"`
	EmptyDir  pulumi.Input            `pulumi:"emptyDir"`
	GitRepo   GitRepoVolumePtrInput   `pulumi:"gitRepo"`
	Name      pulumi.StringInput      `pulumi:"name"`
	Secret    pulumi.StringMapInput   `pulumi:"secret"`
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

func (o VolumeOutput) GitRepo() GitRepoVolumePtrOutput {
	return o.ApplyT(func(v Volume) *GitRepoVolume { return v.GitRepo }).(GitRepoVolumePtrOutput)
}

func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Volume) string { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeOutput) Secret() pulumi.StringMapOutput {
	return o.ApplyT(func(v Volume) map[string]string { return v.Secret }).(pulumi.StringMapOutput)
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
	GitRepo   *GitRepoVolumeResponse   `pulumi:"gitRepo"`
	Name      string                   `pulumi:"name"`
	Secret    map[string]string        `pulumi:"secret"`
}





type VolumeResponseInput interface {
	pulumi.Input

	ToVolumeResponseOutput() VolumeResponseOutput
	ToVolumeResponseOutputWithContext(context.Context) VolumeResponseOutput
}

type VolumeResponseArgs struct {
	AzureFile AzureFileVolumeResponsePtrInput `pulumi:"azureFile"`
	EmptyDir  pulumi.Input                    `pulumi:"emptyDir"`
	GitRepo   GitRepoVolumeResponsePtrInput   `pulumi:"gitRepo"`
	Name      pulumi.StringInput              `pulumi:"name"`
	Secret    pulumi.StringMapInput           `pulumi:"secret"`
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

func (o VolumeResponseOutput) GitRepo() GitRepoVolumeResponsePtrOutput {
	return o.ApplyT(func(v VolumeResponse) *GitRepoVolumeResponse { return v.GitRepo }).(GitRepoVolumeResponsePtrOutput)
}

func (o VolumeResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeResponseOutput) Secret() pulumi.StringMapOutput {
	return o.ApplyT(func(v VolumeResponse) map[string]string { return v.Secret }).(pulumi.StringMapOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFileVolumeInput)(nil)).Elem(), AzureFileVolumeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFileVolumePtrInput)(nil)).Elem(), AzureFileVolumeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFileVolumeResponseInput)(nil)).Elem(), AzureFileVolumeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFileVolumeResponsePtrInput)(nil)).Elem(), AzureFileVolumeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerInput)(nil)).Elem(), ContainerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerArrayInput)(nil)).Elem(), ContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerExecInput)(nil)).Elem(), ContainerExecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerExecPtrInput)(nil)).Elem(), ContainerExecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerExecResponseInput)(nil)).Elem(), ContainerExecResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerExecResponsePtrInput)(nil)).Elem(), ContainerExecResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupDiagnosticsInput)(nil)).Elem(), ContainerGroupDiagnosticsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupDiagnosticsPtrInput)(nil)).Elem(), ContainerGroupDiagnosticsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupDiagnosticsResponseInput)(nil)).Elem(), ContainerGroupDiagnosticsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupDiagnosticsResponsePtrInput)(nil)).Elem(), ContainerGroupDiagnosticsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupIdentityInput)(nil)).Elem(), ContainerGroupIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupIdentityPtrInput)(nil)).Elem(), ContainerGroupIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupIdentityResponseInput)(nil)).Elem(), ContainerGroupIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupIdentityResponsePtrInput)(nil)).Elem(), ContainerGroupIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupIdentityResponseUserAssignedIdentitiesInput)(nil)).Elem(), ContainerGroupIdentityResponseUserAssignedIdentitiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupIdentityResponseUserAssignedIdentitiesMapInput)(nil)).Elem(), ContainerGroupIdentityResponseUserAssignedIdentitiesMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupResponseInstanceViewInput)(nil)).Elem(), ContainerGroupResponseInstanceViewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupResponseInstanceViewPtrInput)(nil)).Elem(), ContainerGroupResponseInstanceViewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupSubnetIdInput)(nil)).Elem(), ContainerGroupSubnetIdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupSubnetIdArrayInput)(nil)).Elem(), ContainerGroupSubnetIdArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupSubnetIdResponseInput)(nil)).Elem(), ContainerGroupSubnetIdResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupSubnetIdResponseArrayInput)(nil)).Elem(), ContainerGroupSubnetIdResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerHttpGetInput)(nil)).Elem(), ContainerHttpGetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerHttpGetPtrInput)(nil)).Elem(), ContainerHttpGetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerHttpGetResponseInput)(nil)).Elem(), ContainerHttpGetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerHttpGetResponsePtrInput)(nil)).Elem(), ContainerHttpGetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerPortInput)(nil)).Elem(), ContainerPortArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerPortArrayInput)(nil)).Elem(), ContainerPortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerPortResponseInput)(nil)).Elem(), ContainerPortResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerPortResponseArrayInput)(nil)).Elem(), ContainerPortResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerProbeInput)(nil)).Elem(), ContainerProbeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerProbePtrInput)(nil)).Elem(), ContainerProbeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerProbeResponseInput)(nil)).Elem(), ContainerProbeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerProbeResponsePtrInput)(nil)).Elem(), ContainerProbeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerPropertiesResponseInstanceViewInput)(nil)).Elem(), ContainerPropertiesResponseInstanceViewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerResponseInput)(nil)).Elem(), ContainerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerResponseArrayInput)(nil)).Elem(), ContainerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerStateResponseInput)(nil)).Elem(), ContainerStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsConfigurationInput)(nil)).Elem(), DnsConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsConfigurationPtrInput)(nil)).Elem(), DnsConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsConfigurationResponseInput)(nil)).Elem(), DnsConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsConfigurationResponsePtrInput)(nil)).Elem(), DnsConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPropertiesInput)(nil)).Elem(), EncryptionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPropertiesPtrInput)(nil)).Elem(), EncryptionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPropertiesResponseInput)(nil)).Elem(), EncryptionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPropertiesResponsePtrInput)(nil)).Elem(), EncryptionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentVariableInput)(nil)).Elem(), EnvironmentVariableArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentVariableArrayInput)(nil)).Elem(), EnvironmentVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentVariableResponseInput)(nil)).Elem(), EnvironmentVariableResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentVariableResponseArrayInput)(nil)).Elem(), EnvironmentVariableResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventResponseInput)(nil)).Elem(), EventResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventResponseArrayInput)(nil)).Elem(), EventResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepoVolumeInput)(nil)).Elem(), GitRepoVolumeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepoVolumePtrInput)(nil)).Elem(), GitRepoVolumeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepoVolumeResponseInput)(nil)).Elem(), GitRepoVolumeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GitRepoVolumeResponsePtrInput)(nil)).Elem(), GitRepoVolumeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GpuResourceInput)(nil)).Elem(), GpuResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GpuResourcePtrInput)(nil)).Elem(), GpuResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GpuResourceResponseInput)(nil)).Elem(), GpuResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GpuResourceResponsePtrInput)(nil)).Elem(), GpuResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpHeaderInput)(nil)).Elem(), HttpHeaderArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpHeaderArrayInput)(nil)).Elem(), HttpHeaderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpHeaderResponseInput)(nil)).Elem(), HttpHeaderResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpHeaderResponseArrayInput)(nil)).Elem(), HttpHeaderResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRegistryCredentialInput)(nil)).Elem(), ImageRegistryCredentialArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRegistryCredentialArrayInput)(nil)).Elem(), ImageRegistryCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRegistryCredentialResponseInput)(nil)).Elem(), ImageRegistryCredentialResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRegistryCredentialResponseArrayInput)(nil)).Elem(), ImageRegistryCredentialResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InitContainerDefinitionInput)(nil)).Elem(), InitContainerDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InitContainerDefinitionArrayInput)(nil)).Elem(), InitContainerDefinitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InitContainerDefinitionResponseInput)(nil)).Elem(), InitContainerDefinitionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InitContainerDefinitionResponseArrayInput)(nil)).Elem(), InitContainerDefinitionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InitContainerPropertiesDefinitionResponseInstanceViewInput)(nil)).Elem(), InitContainerPropertiesDefinitionResponseInstanceViewArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressInput)(nil)).Elem(), IpAddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressPtrInput)(nil)).Elem(), IpAddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressResponseInput)(nil)).Elem(), IpAddressResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressResponsePtrInput)(nil)).Elem(), IpAddressResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnalyticsInput)(nil)).Elem(), LogAnalyticsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnalyticsPtrInput)(nil)).Elem(), LogAnalyticsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnalyticsResponseInput)(nil)).Elem(), LogAnalyticsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnalyticsResponsePtrInput)(nil)).Elem(), LogAnalyticsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortInput)(nil)).Elem(), PortArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortArrayInput)(nil)).Elem(), PortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortResponseInput)(nil)).Elem(), PortResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortResponseArrayInput)(nil)).Elem(), PortResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceLimitsInput)(nil)).Elem(), ResourceLimitsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceLimitsPtrInput)(nil)).Elem(), ResourceLimitsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceLimitsResponseInput)(nil)).Elem(), ResourceLimitsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceLimitsResponsePtrInput)(nil)).Elem(), ResourceLimitsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceRequestsInput)(nil)).Elem(), ResourceRequestsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceRequestsResponseInput)(nil)).Elem(), ResourceRequestsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceRequirementsInput)(nil)).Elem(), ResourceRequirementsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceRequirementsResponseInput)(nil)).Elem(), ResourceRequirementsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeInput)(nil)).Elem(), VolumeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeArrayInput)(nil)).Elem(), VolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMountInput)(nil)).Elem(), VolumeMountArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMountArrayInput)(nil)).Elem(), VolumeMountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMountResponseInput)(nil)).Elem(), VolumeMountResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMountResponseArrayInput)(nil)).Elem(), VolumeMountResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeResponseInput)(nil)).Elem(), VolumeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeResponseArrayInput)(nil)).Elem(), VolumeResponseArray{})
	pulumi.RegisterOutputType(AzureFileVolumeOutput{})
	pulumi.RegisterOutputType(AzureFileVolumePtrOutput{})
	pulumi.RegisterOutputType(AzureFileVolumeResponseOutput{})
	pulumi.RegisterOutputType(AzureFileVolumeResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerArrayOutput{})
	pulumi.RegisterOutputType(ContainerExecOutput{})
	pulumi.RegisterOutputType(ContainerExecPtrOutput{})
	pulumi.RegisterOutputType(ContainerExecResponseOutput{})
	pulumi.RegisterOutputType(ContainerExecResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupDiagnosticsOutput{})
	pulumi.RegisterOutputType(ContainerGroupDiagnosticsPtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupDiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(ContainerGroupDiagnosticsResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupIdentityOutput{})
	pulumi.RegisterOutputType(ContainerGroupIdentityPtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupIdentityResponseOutput{})
	pulumi.RegisterOutputType(ContainerGroupIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupIdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(ContainerGroupIdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(ContainerGroupResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(ContainerGroupResponseInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupSubnetIdOutput{})
	pulumi.RegisterOutputType(ContainerGroupSubnetIdArrayOutput{})
	pulumi.RegisterOutputType(ContainerGroupSubnetIdResponseOutput{})
	pulumi.RegisterOutputType(ContainerGroupSubnetIdResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerHttpGetOutput{})
	pulumi.RegisterOutputType(ContainerHttpGetPtrOutput{})
	pulumi.RegisterOutputType(ContainerHttpGetResponseOutput{})
	pulumi.RegisterOutputType(ContainerHttpGetResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerPortOutput{})
	pulumi.RegisterOutputType(ContainerPortArrayOutput{})
	pulumi.RegisterOutputType(ContainerPortResponseOutput{})
	pulumi.RegisterOutputType(ContainerPortResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerProbeOutput{})
	pulumi.RegisterOutputType(ContainerProbePtrOutput{})
	pulumi.RegisterOutputType(ContainerProbeResponseOutput{})
	pulumi.RegisterOutputType(ContainerProbeResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerPropertiesResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(ContainerResponseOutput{})
	pulumi.RegisterOutputType(ContainerResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerStateResponseOutput{})
	pulumi.RegisterOutputType(DnsConfigurationOutput{})
	pulumi.RegisterOutputType(DnsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(DnsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(DnsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseArrayOutput{})
	pulumi.RegisterOutputType(EventResponseOutput{})
	pulumi.RegisterOutputType(EventResponseArrayOutput{})
	pulumi.RegisterOutputType(GitRepoVolumeOutput{})
	pulumi.RegisterOutputType(GitRepoVolumePtrOutput{})
	pulumi.RegisterOutputType(GitRepoVolumeResponseOutput{})
	pulumi.RegisterOutputType(GitRepoVolumeResponsePtrOutput{})
	pulumi.RegisterOutputType(GpuResourceOutput{})
	pulumi.RegisterOutputType(GpuResourcePtrOutput{})
	pulumi.RegisterOutputType(GpuResourceResponseOutput{})
	pulumi.RegisterOutputType(GpuResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpHeaderOutput{})
	pulumi.RegisterOutputType(HttpHeaderArrayOutput{})
	pulumi.RegisterOutputType(HttpHeaderResponseOutput{})
	pulumi.RegisterOutputType(HttpHeaderResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialArrayOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseArrayOutput{})
	pulumi.RegisterOutputType(InitContainerDefinitionOutput{})
	pulumi.RegisterOutputType(InitContainerDefinitionArrayOutput{})
	pulumi.RegisterOutputType(InitContainerDefinitionResponseOutput{})
	pulumi.RegisterOutputType(InitContainerDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(InitContainerPropertiesDefinitionResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(IpAddressOutput{})
	pulumi.RegisterOutputType(IpAddressPtrOutput{})
	pulumi.RegisterOutputType(IpAddressResponseOutput{})
	pulumi.RegisterOutputType(IpAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsOutput{})
	pulumi.RegisterOutputType(LogAnalyticsPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsResponsePtrOutput{})
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
