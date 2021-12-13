


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AKS struct {
	ComputeLocation *string        `pulumi:"computeLocation"`
	ComputeType     string         `pulumi:"computeType"`
	Description     *string        `pulumi:"description"`
	Properties      *AKSProperties `pulumi:"properties"`
	ResourceId      *string        `pulumi:"resourceId"`
}

type AKSProperties struct {
	AgentCount       *int              `pulumi:"agentCount"`
	AgentVMSize      *string           `pulumi:"agentVMSize"`
	ClusterFqdn      *string           `pulumi:"clusterFqdn"`
	SslConfiguration *SslConfiguration `pulumi:"sslConfiguration"`
}

type AKSResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *AKSResponseProperties                `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type AKSResponseProperties struct {
	AgentCount       *int                      `pulumi:"agentCount"`
	AgentVMSize      *string                   `pulumi:"agentVMSize"`
	ClusterFqdn      *string                   `pulumi:"clusterFqdn"`
	SslConfiguration *SslConfigurationResponse `pulumi:"sslConfiguration"`
	SystemServices   []SystemServiceResponse   `pulumi:"systemServices"`
}

type BatchAI struct {
	ComputeLocation *string            `pulumi:"computeLocation"`
	ComputeType     string             `pulumi:"computeType"`
	Description     *string            `pulumi:"description"`
	Properties      *BatchAIProperties `pulumi:"properties"`
	ResourceId      *string            `pulumi:"resourceId"`
}

type BatchAIProperties struct {
	ScaleSettings *ScaleSettings `pulumi:"scaleSettings"`
	VmPriority    *string        `pulumi:"vmPriority"`
	VmSize        *string        `pulumi:"vmSize"`
}

type BatchAIResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *BatchAIResponseProperties            `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type BatchAIResponseProperties struct {
	ScaleSettings *ScaleSettingsResponse `pulumi:"scaleSettings"`
	VmPriority    *string                `pulumi:"vmPriority"`
	VmSize        *string                `pulumi:"vmSize"`
}

type DataFactory struct {
	ComputeLocation *string `pulumi:"computeLocation"`
	ComputeType     string  `pulumi:"computeType"`
	Description     *string `pulumi:"description"`
	ResourceId      *string `pulumi:"resourceId"`
}

type DataFactoryResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type ErrorDetailResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type ErrorResponseResponse struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
}

type HDInsight struct {
	ComputeLocation *string              `pulumi:"computeLocation"`
	ComputeType     string               `pulumi:"computeType"`
	Description     *string              `pulumi:"description"`
	Properties      *HDInsightProperties `pulumi:"properties"`
	ResourceId      *string              `pulumi:"resourceId"`
}

type HDInsightProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
}

type HDInsightResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *HDInsightResponseProperties          `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type HDInsightResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
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

type MachineLearningServiceErrorResponse struct {
	Error *ErrorResponseResponse `pulumi:"error"`
}

type PasswordResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type RegistryListCredentialsResultResponse struct {
	Location  string             `pulumi:"location"`
	Passwords []PasswordResponse `pulumi:"passwords"`
	Username  string             `pulumi:"username"`
}

type ScaleSettings struct {
	AutoScaleEnabled *bool `pulumi:"autoScaleEnabled"`
	MaxNodeCount     *int  `pulumi:"maxNodeCount"`
	MinNodeCount     *int  `pulumi:"minNodeCount"`
}

type ScaleSettingsResponse struct {
	AutoScaleEnabled *bool `pulumi:"autoScaleEnabled"`
	MaxNodeCount     *int  `pulumi:"maxNodeCount"`
	MinNodeCount     *int  `pulumi:"minNodeCount"`
}

type SslConfiguration struct {
	Cert   *string `pulumi:"cert"`
	Cname  *string `pulumi:"cname"`
	Key    *string `pulumi:"key"`
	Status *string `pulumi:"status"`
}

type SslConfigurationResponse struct {
	Cert   *string `pulumi:"cert"`
	Cname  *string `pulumi:"cname"`
	Key    *string `pulumi:"key"`
	Status *string `pulumi:"status"`
}

type SystemServiceResponse struct {
	PublicIpAddress   string `pulumi:"publicIpAddress"`
	SystemServiceType string `pulumi:"systemServiceType"`
	Version           string `pulumi:"version"`
}

type VirtualMachine struct {
	ComputeLocation *string                   `pulumi:"computeLocation"`
	ComputeType     string                    `pulumi:"computeType"`
	Description     *string                   `pulumi:"description"`
	Properties      *VirtualMachineProperties `pulumi:"properties"`
	ResourceId      *string                   `pulumi:"resourceId"`
}

type VirtualMachineProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
	VirtualMachineSize   *string                       `pulumi:"virtualMachineSize"`
}

type VirtualMachineResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *VirtualMachineResponseProperties     `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}

type VirtualMachineResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
	VirtualMachineSize   *string                               `pulumi:"virtualMachineSize"`
}

type VirtualMachineSshCredentials struct {
	Password       *string `pulumi:"password"`
	PrivateKeyData *string `pulumi:"privateKeyData"`
	PublicKeyData  *string `pulumi:"publicKeyData"`
	Username       *string `pulumi:"username"`
}

type VirtualMachineSshCredentialsResponse struct {
	Password       *string `pulumi:"password"`
	PrivateKeyData *string `pulumi:"privateKeyData"`
	PublicKeyData  *string `pulumi:"publicKeyData"`
	Username       *string `pulumi:"username"`
}

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
}
