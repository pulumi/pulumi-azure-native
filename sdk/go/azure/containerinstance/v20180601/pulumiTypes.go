


package v20180601

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

type ContainerGroupResponseInstanceView struct {
	Events []EventResponse `pulumi:"events"`
	State  string          `pulumi:"state"`
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

func (o ContainerGroupResponseInstanceViewOutput) Events() EventResponseArrayOutput {
	return o.ApplyT(func(v ContainerGroupResponseInstanceView) []EventResponse { return v.Events }).(EventResponseArrayOutput)
}

func (o ContainerGroupResponseInstanceViewOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerGroupResponseInstanceView) string { return v.State }).(pulumi.StringOutput)
}

type ContainerHttpGet struct {
	Path   *string `pulumi:"path"`
	Port   int     `pulumi:"port"`
	Scheme *string `pulumi:"scheme"`
}





type ContainerHttpGetInput interface {
	pulumi.Input

	ToContainerHttpGetOutput() ContainerHttpGetOutput
	ToContainerHttpGetOutputWithContext(context.Context) ContainerHttpGetOutput
}

type ContainerHttpGetArgs struct {
	Path   pulumi.StringPtrInput `pulumi:"path"`
	Port   pulumi.IntInput       `pulumi:"port"`
	Scheme pulumi.StringPtrInput `pulumi:"scheme"`
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
	Path   *string `pulumi:"path"`
	Port   int     `pulumi:"port"`
	Scheme *string `pulumi:"scheme"`
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
	DetailStatus *string `pulumi:"detailStatus"`
	ExitCode     *int    `pulumi:"exitCode"`
	FinishTime   *string `pulumi:"finishTime"`
	StartTime    *string `pulumi:"startTime"`
	State        *string `pulumi:"state"`
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
	Count          *int    `pulumi:"count"`
	FirstTimestamp *string `pulumi:"firstTimestamp"`
	LastTimestamp  *string `pulumi:"lastTimestamp"`
	Message        *string `pulumi:"message"`
	Name           *string `pulumi:"name"`
	Type           *string `pulumi:"type"`
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
	WorkspaceId  string `pulumi:"workspaceId"`
	WorkspaceKey string `pulumi:"workspaceKey"`
}





type LogAnalyticsInput interface {
	pulumi.Input

	ToLogAnalyticsOutput() LogAnalyticsOutput
	ToLogAnalyticsOutputWithContext(context.Context) LogAnalyticsOutput
}

type LogAnalyticsArgs struct {
	WorkspaceId  pulumi.StringInput `pulumi:"workspaceId"`
	WorkspaceKey pulumi.StringInput `pulumi:"workspaceKey"`
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

func (o LogAnalyticsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalytics) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o LogAnalyticsOutput) WorkspaceKey() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalytics) string { return v.WorkspaceKey }).(pulumi.StringOutput)
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

type LogAnalyticsResponse struct {
	WorkspaceId  string `pulumi:"workspaceId"`
	WorkspaceKey string `pulumi:"workspaceKey"`
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

func (o LogAnalyticsResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o LogAnalyticsResponseOutput) WorkspaceKey() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsResponse) string { return v.WorkspaceKey }).(pulumi.StringOutput)
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
	pulumi.RegisterOutputType(ContainerGroupResponseInstanceViewOutput{})
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
	pulumi.RegisterOutputType(ImageRegistryCredentialOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialArrayOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseArrayOutput{})
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
