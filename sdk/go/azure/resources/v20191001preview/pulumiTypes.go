


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerConfiguration struct {
	ContainerGroupName *string `pulumi:"containerGroupName"`
}





type ContainerConfigurationInput interface {
	pulumi.Input

	ToContainerConfigurationOutput() ContainerConfigurationOutput
	ToContainerConfigurationOutputWithContext(context.Context) ContainerConfigurationOutput
}

type ContainerConfigurationArgs struct {
	ContainerGroupName pulumi.StringPtrInput `pulumi:"containerGroupName"`
}

func (ContainerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerConfiguration)(nil)).Elem()
}

func (i ContainerConfigurationArgs) ToContainerConfigurationOutput() ContainerConfigurationOutput {
	return i.ToContainerConfigurationOutputWithContext(context.Background())
}

func (i ContainerConfigurationArgs) ToContainerConfigurationOutputWithContext(ctx context.Context) ContainerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerConfigurationOutput)
}

func (i ContainerConfigurationArgs) ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput {
	return i.ToContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i ContainerConfigurationArgs) ToContainerConfigurationPtrOutputWithContext(ctx context.Context) ContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerConfigurationOutput).ToContainerConfigurationPtrOutputWithContext(ctx)
}









type ContainerConfigurationPtrInput interface {
	pulumi.Input

	ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput
	ToContainerConfigurationPtrOutputWithContext(context.Context) ContainerConfigurationPtrOutput
}

type containerConfigurationPtrType ContainerConfigurationArgs

func ContainerConfigurationPtr(v *ContainerConfigurationArgs) ContainerConfigurationPtrInput {
	return (*containerConfigurationPtrType)(v)
}

func (*containerConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerConfiguration)(nil)).Elem()
}

func (i *containerConfigurationPtrType) ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput {
	return i.ToContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i *containerConfigurationPtrType) ToContainerConfigurationPtrOutputWithContext(ctx context.Context) ContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerConfigurationPtrOutput)
}

type ContainerConfigurationOutput struct{ *pulumi.OutputState }

func (ContainerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerConfiguration)(nil)).Elem()
}

func (o ContainerConfigurationOutput) ToContainerConfigurationOutput() ContainerConfigurationOutput {
	return o
}

func (o ContainerConfigurationOutput) ToContainerConfigurationOutputWithContext(ctx context.Context) ContainerConfigurationOutput {
	return o
}

func (o ContainerConfigurationOutput) ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput {
	return o.ToContainerConfigurationPtrOutputWithContext(context.Background())
}

func (o ContainerConfigurationOutput) ToContainerConfigurationPtrOutputWithContext(ctx context.Context) ContainerConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerConfiguration) *ContainerConfiguration {
		return &v
	}).(ContainerConfigurationPtrOutput)
}

func (o ContainerConfigurationOutput) ContainerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerConfiguration) *string { return v.ContainerGroupName }).(pulumi.StringPtrOutput)
}

type ContainerConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ContainerConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerConfiguration)(nil)).Elem()
}

func (o ContainerConfigurationPtrOutput) ToContainerConfigurationPtrOutput() ContainerConfigurationPtrOutput {
	return o
}

func (o ContainerConfigurationPtrOutput) ToContainerConfigurationPtrOutputWithContext(ctx context.Context) ContainerConfigurationPtrOutput {
	return o
}

func (o ContainerConfigurationPtrOutput) Elem() ContainerConfigurationOutput {
	return o.ApplyT(func(v *ContainerConfiguration) ContainerConfiguration {
		if v != nil {
			return *v
		}
		var ret ContainerConfiguration
		return ret
	}).(ContainerConfigurationOutput)
}

func (o ContainerConfigurationPtrOutput) ContainerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ContainerGroupName
	}).(pulumi.StringPtrOutput)
}

type ContainerConfigurationResponse struct {
	ContainerGroupName *string `pulumi:"containerGroupName"`
}

type ContainerConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContainerConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerConfigurationResponse)(nil)).Elem()
}

func (o ContainerConfigurationResponseOutput) ToContainerConfigurationResponseOutput() ContainerConfigurationResponseOutput {
	return o
}

func (o ContainerConfigurationResponseOutput) ToContainerConfigurationResponseOutputWithContext(ctx context.Context) ContainerConfigurationResponseOutput {
	return o
}

func (o ContainerConfigurationResponseOutput) ContainerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerConfigurationResponse) *string { return v.ContainerGroupName }).(pulumi.StringPtrOutput)
}

type ContainerConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerConfigurationResponse)(nil)).Elem()
}

func (o ContainerConfigurationResponsePtrOutput) ToContainerConfigurationResponsePtrOutput() ContainerConfigurationResponsePtrOutput {
	return o
}

func (o ContainerConfigurationResponsePtrOutput) ToContainerConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerConfigurationResponsePtrOutput {
	return o
}

func (o ContainerConfigurationResponsePtrOutput) Elem() ContainerConfigurationResponseOutput {
	return o.ApplyT(func(v *ContainerConfigurationResponse) ContainerConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ContainerConfigurationResponse
		return ret
	}).(ContainerConfigurationResponseOutput)
}

func (o ContainerConfigurationResponsePtrOutput) ContainerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContainerGroupName
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

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type ErrorAdditionalInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o ErrorAdditionalInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) Index(i pulumi.IntInput) ErrorAdditionalInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorAdditionalInfoResponse {
		return vs[0].([]ErrorAdditionalInfoResponse)[vs[1].(int)]
	}).(ErrorAdditionalInfoResponseOutput)
}

type ErrorResponseResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorResponseResponse       `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}

type ErrorResponseResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponseResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorResponseResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorResponseResponseOutput) Details() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponseResponse) []ErrorResponseResponse { return v.Details }).(ErrorResponseResponseArrayOutput)
}

func (o ErrorResponseResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorResponseResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorResponseResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponsePtrOutput) ToErrorResponseResponsePtrOutput() ErrorResponseResponsePtrOutput {
	return o
}

func (o ErrorResponseResponsePtrOutput) ToErrorResponseResponsePtrOutputWithContext(ctx context.Context) ErrorResponseResponsePtrOutput {
	return o
}

func (o ErrorResponseResponsePtrOutput) Elem() ErrorResponseResponseOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) ErrorResponseResponse {
		if v != nil {
			return *v
		}
		var ret ErrorResponseResponse
		return ret
	}).(ErrorResponseResponseOutput)
}

func (o ErrorResponseResponsePtrOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) []ErrorAdditionalInfoResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalInfo
	}).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorResponseResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponseResponsePtrOutput) Details() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) []ErrorResponseResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorResponseResponseArrayOutput)
}

func (o ErrorResponseResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponseResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
}

type ErrorResponseResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseArrayOutput) ToErrorResponseResponseArrayOutput() ErrorResponseResponseArrayOutput {
	return o
}

func (o ErrorResponseResponseArrayOutput) ToErrorResponseResponseArrayOutputWithContext(ctx context.Context) ErrorResponseResponseArrayOutput {
	return o
}

func (o ErrorResponseResponseArrayOutput) Index(i pulumi.IntInput) ErrorResponseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorResponseResponse {
		return vs[0].([]ErrorResponseResponse)[vs[1].(int)]
	}).(ErrorResponseResponseOutput)
}

type ManagedServiceIdentity struct {
	Type                   *string                `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ScriptStatusResponse struct {
	ContainerInstanceId string                 `pulumi:"containerInstanceId"`
	EndTime             string                 `pulumi:"endTime"`
	Error               *ErrorResponseResponse `pulumi:"error"`
	ExpirationTime      string                 `pulumi:"expirationTime"`
	StartTime           string                 `pulumi:"startTime"`
	StorageAccountId    string                 `pulumi:"storageAccountId"`
}

type ScriptStatusResponseOutput struct{ *pulumi.OutputState }

func (ScriptStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptStatusResponse)(nil)).Elem()
}

func (o ScriptStatusResponseOutput) ToScriptStatusResponseOutput() ScriptStatusResponseOutput {
	return o
}

func (o ScriptStatusResponseOutput) ToScriptStatusResponseOutputWithContext(ctx context.Context) ScriptStatusResponseOutput {
	return o
}

func (o ScriptStatusResponseOutput) ContainerInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStatusResponse) string { return v.ContainerInstanceId }).(pulumi.StringOutput)
}

func (o ScriptStatusResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStatusResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o ScriptStatusResponseOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v ScriptStatusResponse) *ErrorResponseResponse { return v.Error }).(ErrorResponseResponsePtrOutput)
}

func (o ScriptStatusResponseOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStatusResponse) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

func (o ScriptStatusResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStatusResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o ScriptStatusResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v ScriptStatusResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type StorageAccountConfiguration struct {
	StorageAccountKey  *string `pulumi:"storageAccountKey"`
	StorageAccountName *string `pulumi:"storageAccountName"`
}





type StorageAccountConfigurationInput interface {
	pulumi.Input

	ToStorageAccountConfigurationOutput() StorageAccountConfigurationOutput
	ToStorageAccountConfigurationOutputWithContext(context.Context) StorageAccountConfigurationOutput
}

type StorageAccountConfigurationArgs struct {
	StorageAccountKey  pulumi.StringPtrInput `pulumi:"storageAccountKey"`
	StorageAccountName pulumi.StringPtrInput `pulumi:"storageAccountName"`
}

func (StorageAccountConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountConfiguration)(nil)).Elem()
}

func (i StorageAccountConfigurationArgs) ToStorageAccountConfigurationOutput() StorageAccountConfigurationOutput {
	return i.ToStorageAccountConfigurationOutputWithContext(context.Background())
}

func (i StorageAccountConfigurationArgs) ToStorageAccountConfigurationOutputWithContext(ctx context.Context) StorageAccountConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountConfigurationOutput)
}

func (i StorageAccountConfigurationArgs) ToStorageAccountConfigurationPtrOutput() StorageAccountConfigurationPtrOutput {
	return i.ToStorageAccountConfigurationPtrOutputWithContext(context.Background())
}

func (i StorageAccountConfigurationArgs) ToStorageAccountConfigurationPtrOutputWithContext(ctx context.Context) StorageAccountConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountConfigurationOutput).ToStorageAccountConfigurationPtrOutputWithContext(ctx)
}









type StorageAccountConfigurationPtrInput interface {
	pulumi.Input

	ToStorageAccountConfigurationPtrOutput() StorageAccountConfigurationPtrOutput
	ToStorageAccountConfigurationPtrOutputWithContext(context.Context) StorageAccountConfigurationPtrOutput
}

type storageAccountConfigurationPtrType StorageAccountConfigurationArgs

func StorageAccountConfigurationPtr(v *StorageAccountConfigurationArgs) StorageAccountConfigurationPtrInput {
	return (*storageAccountConfigurationPtrType)(v)
}

func (*storageAccountConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountConfiguration)(nil)).Elem()
}

func (i *storageAccountConfigurationPtrType) ToStorageAccountConfigurationPtrOutput() StorageAccountConfigurationPtrOutput {
	return i.ToStorageAccountConfigurationPtrOutputWithContext(context.Background())
}

func (i *storageAccountConfigurationPtrType) ToStorageAccountConfigurationPtrOutputWithContext(ctx context.Context) StorageAccountConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountConfigurationPtrOutput)
}

type StorageAccountConfigurationOutput struct{ *pulumi.OutputState }

func (StorageAccountConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountConfiguration)(nil)).Elem()
}

func (o StorageAccountConfigurationOutput) ToStorageAccountConfigurationOutput() StorageAccountConfigurationOutput {
	return o
}

func (o StorageAccountConfigurationOutput) ToStorageAccountConfigurationOutputWithContext(ctx context.Context) StorageAccountConfigurationOutput {
	return o
}

func (o StorageAccountConfigurationOutput) ToStorageAccountConfigurationPtrOutput() StorageAccountConfigurationPtrOutput {
	return o.ToStorageAccountConfigurationPtrOutputWithContext(context.Background())
}

func (o StorageAccountConfigurationOutput) ToStorageAccountConfigurationPtrOutputWithContext(ctx context.Context) StorageAccountConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountConfiguration) *StorageAccountConfiguration {
		return &v
	}).(StorageAccountConfigurationPtrOutput)
}

func (o StorageAccountConfigurationOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountConfiguration) *string { return v.StorageAccountKey }).(pulumi.StringPtrOutput)
}

func (o StorageAccountConfigurationOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountConfiguration) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type StorageAccountConfigurationPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountConfiguration)(nil)).Elem()
}

func (o StorageAccountConfigurationPtrOutput) ToStorageAccountConfigurationPtrOutput() StorageAccountConfigurationPtrOutput {
	return o
}

func (o StorageAccountConfigurationPtrOutput) ToStorageAccountConfigurationPtrOutputWithContext(ctx context.Context) StorageAccountConfigurationPtrOutput {
	return o
}

func (o StorageAccountConfigurationPtrOutput) Elem() StorageAccountConfigurationOutput {
	return o.ApplyT(func(v *StorageAccountConfiguration) StorageAccountConfiguration {
		if v != nil {
			return *v
		}
		var ret StorageAccountConfiguration
		return ret
	}).(StorageAccountConfigurationOutput)
}

func (o StorageAccountConfigurationPtrOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountConfigurationPtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type StorageAccountConfigurationResponse struct {
	StorageAccountKey  *string `pulumi:"storageAccountKey"`
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type StorageAccountConfigurationResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountConfigurationResponse)(nil)).Elem()
}

func (o StorageAccountConfigurationResponseOutput) ToStorageAccountConfigurationResponseOutput() StorageAccountConfigurationResponseOutput {
	return o
}

func (o StorageAccountConfigurationResponseOutput) ToStorageAccountConfigurationResponseOutputWithContext(ctx context.Context) StorageAccountConfigurationResponseOutput {
	return o
}

func (o StorageAccountConfigurationResponseOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountConfigurationResponse) *string { return v.StorageAccountKey }).(pulumi.StringPtrOutput)
}

func (o StorageAccountConfigurationResponseOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountConfigurationResponse) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type StorageAccountConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountConfigurationResponse)(nil)).Elem()
}

func (o StorageAccountConfigurationResponsePtrOutput) ToStorageAccountConfigurationResponsePtrOutput() StorageAccountConfigurationResponsePtrOutput {
	return o
}

func (o StorageAccountConfigurationResponsePtrOutput) ToStorageAccountConfigurationResponsePtrOutputWithContext(ctx context.Context) StorageAccountConfigurationResponsePtrOutput {
	return o
}

func (o StorageAccountConfigurationResponsePtrOutput) Elem() StorageAccountConfigurationResponseOutput {
	return o.ApplyT(func(v *StorageAccountConfigurationResponse) StorageAccountConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountConfigurationResponse
		return ret
	}).(StorageAccountConfigurationResponseOutput)
}

func (o StorageAccountConfigurationResponsePtrOutput) StorageAccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountConfigurationResponsePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ContainerConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContainerConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ScriptStatusResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountConfigurationOutput{})
	pulumi.RegisterOutputType(StorageAccountConfigurationPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountConfigurationResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
