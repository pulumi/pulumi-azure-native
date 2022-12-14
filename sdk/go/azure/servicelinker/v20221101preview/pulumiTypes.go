


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessKeyInfoBase struct {
	AuthType    string   `pulumi:"authType"`
	Permissions []string `pulumi:"permissions"`
}

type AccessKeyInfoBaseResponse struct {
	AuthType    string   `pulumi:"authType"`
	Permissions []string `pulumi:"permissions"`
}

type AzureKeyVaultProperties struct {
	ConnectAsKubernetesCsiDriver *bool  `pulumi:"connectAsKubernetesCsiDriver"`
	Type                         string `pulumi:"type"`
}

type AzureKeyVaultPropertiesResponse struct {
	ConnectAsKubernetesCsiDriver *bool  `pulumi:"connectAsKubernetesCsiDriver"`
	Type                         string `pulumi:"type"`
}

type AzureResource struct {
	Id                 *string                  `pulumi:"id"`
	ResourceProperties *AzureKeyVaultProperties `pulumi:"resourceProperties"`
	Type               string                   `pulumi:"type"`
}

type AzureResourceResponse struct {
	Id                 *string                          `pulumi:"id"`
	ResourceProperties *AzureKeyVaultPropertiesResponse `pulumi:"resourceProperties"`
	Type               string                           `pulumi:"type"`
}

type BasicErrorDryrunPrerequisiteResultResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
	Type    string  `pulumi:"type"`
}

type ConfigurationInfo struct {
	Action                   *string           `pulumi:"action"`
	AdditionalConfigurations map[string]string `pulumi:"additionalConfigurations"`
	CustomizedKeys           map[string]string `pulumi:"customizedKeys"`
	DeleteOrUpdateBehavior   *string           `pulumi:"deleteOrUpdateBehavior"`
}





type ConfigurationInfoInput interface {
	pulumi.Input

	ToConfigurationInfoOutput() ConfigurationInfoOutput
	ToConfigurationInfoOutputWithContext(context.Context) ConfigurationInfoOutput
}

type ConfigurationInfoArgs struct {
	Action                   pulumi.StringPtrInput `pulumi:"action"`
	AdditionalConfigurations pulumi.StringMapInput `pulumi:"additionalConfigurations"`
	CustomizedKeys           pulumi.StringMapInput `pulumi:"customizedKeys"`
	DeleteOrUpdateBehavior   pulumi.StringPtrInput `pulumi:"deleteOrUpdateBehavior"`
}

func (ConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationInfo)(nil)).Elem()
}

func (i ConfigurationInfoArgs) ToConfigurationInfoOutput() ConfigurationInfoOutput {
	return i.ToConfigurationInfoOutputWithContext(context.Background())
}

func (i ConfigurationInfoArgs) ToConfigurationInfoOutputWithContext(ctx context.Context) ConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationInfoOutput)
}

func (i ConfigurationInfoArgs) ToConfigurationInfoPtrOutput() ConfigurationInfoPtrOutput {
	return i.ToConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ConfigurationInfoArgs) ToConfigurationInfoPtrOutputWithContext(ctx context.Context) ConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationInfoOutput).ToConfigurationInfoPtrOutputWithContext(ctx)
}









type ConfigurationInfoPtrInput interface {
	pulumi.Input

	ToConfigurationInfoPtrOutput() ConfigurationInfoPtrOutput
	ToConfigurationInfoPtrOutputWithContext(context.Context) ConfigurationInfoPtrOutput
}

type configurationInfoPtrType ConfigurationInfoArgs

func ConfigurationInfoPtr(v *ConfigurationInfoArgs) ConfigurationInfoPtrInput {
	return (*configurationInfoPtrType)(v)
}

func (*configurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationInfo)(nil)).Elem()
}

func (i *configurationInfoPtrType) ToConfigurationInfoPtrOutput() ConfigurationInfoPtrOutput {
	return i.ToConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *configurationInfoPtrType) ToConfigurationInfoPtrOutputWithContext(ctx context.Context) ConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationInfoPtrOutput)
}

type ConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationInfo)(nil)).Elem()
}

func (o ConfigurationInfoOutput) ToConfigurationInfoOutput() ConfigurationInfoOutput {
	return o
}

func (o ConfigurationInfoOutput) ToConfigurationInfoOutputWithContext(ctx context.Context) ConfigurationInfoOutput {
	return o
}

func (o ConfigurationInfoOutput) ToConfigurationInfoPtrOutput() ConfigurationInfoPtrOutput {
	return o.ToConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ConfigurationInfoOutput) ToConfigurationInfoPtrOutputWithContext(ctx context.Context) ConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationInfo) *ConfigurationInfo {
		return &v
	}).(ConfigurationInfoPtrOutput)
}

func (o ConfigurationInfoOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationInfo) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o ConfigurationInfoOutput) AdditionalConfigurations() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConfigurationInfo) map[string]string { return v.AdditionalConfigurations }).(pulumi.StringMapOutput)
}

func (o ConfigurationInfoOutput) CustomizedKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConfigurationInfo) map[string]string { return v.CustomizedKeys }).(pulumi.StringMapOutput)
}

func (o ConfigurationInfoOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationInfo) *string { return v.DeleteOrUpdateBehavior }).(pulumi.StringPtrOutput)
}

type ConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationInfo)(nil)).Elem()
}

func (o ConfigurationInfoPtrOutput) ToConfigurationInfoPtrOutput() ConfigurationInfoPtrOutput {
	return o
}

func (o ConfigurationInfoPtrOutput) ToConfigurationInfoPtrOutputWithContext(ctx context.Context) ConfigurationInfoPtrOutput {
	return o
}

func (o ConfigurationInfoPtrOutput) Elem() ConfigurationInfoOutput {
	return o.ApplyT(func(v *ConfigurationInfo) ConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ConfigurationInfo
		return ret
	}).(ConfigurationInfoOutput)
}

func (o ConfigurationInfoPtrOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Action
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationInfoPtrOutput) AdditionalConfigurations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationInfo) map[string]string {
		if v == nil {
			return nil
		}
		return v.AdditionalConfigurations
	}).(pulumi.StringMapOutput)
}

func (o ConfigurationInfoPtrOutput) CustomizedKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationInfo) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomizedKeys
	}).(pulumi.StringMapOutput)
}

func (o ConfigurationInfoPtrOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOrUpdateBehavior
	}).(pulumi.StringPtrOutput)
}

type ConfigurationInfoResponse struct {
	Action                   *string           `pulumi:"action"`
	AdditionalConfigurations map[string]string `pulumi:"additionalConfigurations"`
	CustomizedKeys           map[string]string `pulumi:"customizedKeys"`
	DeleteOrUpdateBehavior   *string           `pulumi:"deleteOrUpdateBehavior"`
}

type ConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationInfoResponse)(nil)).Elem()
}

func (o ConfigurationInfoResponseOutput) ToConfigurationInfoResponseOutput() ConfigurationInfoResponseOutput {
	return o
}

func (o ConfigurationInfoResponseOutput) ToConfigurationInfoResponseOutputWithContext(ctx context.Context) ConfigurationInfoResponseOutput {
	return o
}

func (o ConfigurationInfoResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationInfoResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o ConfigurationInfoResponseOutput) AdditionalConfigurations() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConfigurationInfoResponse) map[string]string { return v.AdditionalConfigurations }).(pulumi.StringMapOutput)
}

func (o ConfigurationInfoResponseOutput) CustomizedKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ConfigurationInfoResponse) map[string]string { return v.CustomizedKeys }).(pulumi.StringMapOutput)
}

func (o ConfigurationInfoResponseOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationInfoResponse) *string { return v.DeleteOrUpdateBehavior }).(pulumi.StringPtrOutput)
}

type ConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationInfoResponse)(nil)).Elem()
}

func (o ConfigurationInfoResponsePtrOutput) ToConfigurationInfoResponsePtrOutput() ConfigurationInfoResponsePtrOutput {
	return o
}

func (o ConfigurationInfoResponsePtrOutput) ToConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ConfigurationInfoResponsePtrOutput {
	return o
}

func (o ConfigurationInfoResponsePtrOutput) Elem() ConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ConfigurationInfoResponse) ConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationInfoResponse
		return ret
	}).(ConfigurationInfoResponseOutput)
}

func (o ConfigurationInfoResponsePtrOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Action
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationInfoResponsePtrOutput) AdditionalConfigurations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationInfoResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.AdditionalConfigurations
	}).(pulumi.StringMapOutput)
}

func (o ConfigurationInfoResponsePtrOutput) CustomizedKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationInfoResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomizedKeys
	}).(pulumi.StringMapOutput)
}

func (o ConfigurationInfoResponsePtrOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOrUpdateBehavior
	}).(pulumi.StringPtrOutput)
}

type ConfluentBootstrapServer struct {
	Endpoint *string `pulumi:"endpoint"`
	Type     string  `pulumi:"type"`
}

type ConfluentBootstrapServerResponse struct {
	Endpoint *string `pulumi:"endpoint"`
	Type     string  `pulumi:"type"`
}

type ConfluentSchemaRegistry struct {
	Endpoint *string `pulumi:"endpoint"`
	Type     string  `pulumi:"type"`
}

type ConfluentSchemaRegistryResponse struct {
	Endpoint *string `pulumi:"endpoint"`
	Type     string  `pulumi:"type"`
}

type CreateOrUpdateDryrunParameters struct {
	ActionName            string                 `pulumi:"actionName"`
	AuthInfo              interface{}            `pulumi:"authInfo"`
	ClientType            *string                `pulumi:"clientType"`
	ConfigurationInfo     *ConfigurationInfo     `pulumi:"configurationInfo"`
	PublicNetworkSolution *PublicNetworkSolution `pulumi:"publicNetworkSolution"`
	Scope                 *string                `pulumi:"scope"`
	SecretStore           *SecretStore           `pulumi:"secretStore"`
	TargetService         interface{}            `pulumi:"targetService"`
	VNetSolution          *VNetSolution          `pulumi:"vNetSolution"`
}





type CreateOrUpdateDryrunParametersInput interface {
	pulumi.Input

	ToCreateOrUpdateDryrunParametersOutput() CreateOrUpdateDryrunParametersOutput
	ToCreateOrUpdateDryrunParametersOutputWithContext(context.Context) CreateOrUpdateDryrunParametersOutput
}

type CreateOrUpdateDryrunParametersArgs struct {
	ActionName            pulumi.StringInput            `pulumi:"actionName"`
	AuthInfo              pulumi.Input                  `pulumi:"authInfo"`
	ClientType            pulumi.StringPtrInput         `pulumi:"clientType"`
	ConfigurationInfo     ConfigurationInfoPtrInput     `pulumi:"configurationInfo"`
	PublicNetworkSolution PublicNetworkSolutionPtrInput `pulumi:"publicNetworkSolution"`
	Scope                 pulumi.StringPtrInput         `pulumi:"scope"`
	SecretStore           SecretStorePtrInput           `pulumi:"secretStore"`
	TargetService         pulumi.Input                  `pulumi:"targetService"`
	VNetSolution          VNetSolutionPtrInput          `pulumi:"vNetSolution"`
}

func (CreateOrUpdateDryrunParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateOrUpdateDryrunParameters)(nil)).Elem()
}

func (i CreateOrUpdateDryrunParametersArgs) ToCreateOrUpdateDryrunParametersOutput() CreateOrUpdateDryrunParametersOutput {
	return i.ToCreateOrUpdateDryrunParametersOutputWithContext(context.Background())
}

func (i CreateOrUpdateDryrunParametersArgs) ToCreateOrUpdateDryrunParametersOutputWithContext(ctx context.Context) CreateOrUpdateDryrunParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateOrUpdateDryrunParametersOutput)
}

func (i CreateOrUpdateDryrunParametersArgs) ToCreateOrUpdateDryrunParametersPtrOutput() CreateOrUpdateDryrunParametersPtrOutput {
	return i.ToCreateOrUpdateDryrunParametersPtrOutputWithContext(context.Background())
}

func (i CreateOrUpdateDryrunParametersArgs) ToCreateOrUpdateDryrunParametersPtrOutputWithContext(ctx context.Context) CreateOrUpdateDryrunParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateOrUpdateDryrunParametersOutput).ToCreateOrUpdateDryrunParametersPtrOutputWithContext(ctx)
}









type CreateOrUpdateDryrunParametersPtrInput interface {
	pulumi.Input

	ToCreateOrUpdateDryrunParametersPtrOutput() CreateOrUpdateDryrunParametersPtrOutput
	ToCreateOrUpdateDryrunParametersPtrOutputWithContext(context.Context) CreateOrUpdateDryrunParametersPtrOutput
}

type createOrUpdateDryrunParametersPtrType CreateOrUpdateDryrunParametersArgs

func CreateOrUpdateDryrunParametersPtr(v *CreateOrUpdateDryrunParametersArgs) CreateOrUpdateDryrunParametersPtrInput {
	return (*createOrUpdateDryrunParametersPtrType)(v)
}

func (*createOrUpdateDryrunParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateOrUpdateDryrunParameters)(nil)).Elem()
}

func (i *createOrUpdateDryrunParametersPtrType) ToCreateOrUpdateDryrunParametersPtrOutput() CreateOrUpdateDryrunParametersPtrOutput {
	return i.ToCreateOrUpdateDryrunParametersPtrOutputWithContext(context.Background())
}

func (i *createOrUpdateDryrunParametersPtrType) ToCreateOrUpdateDryrunParametersPtrOutputWithContext(ctx context.Context) CreateOrUpdateDryrunParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateOrUpdateDryrunParametersPtrOutput)
}

type CreateOrUpdateDryrunParametersOutput struct{ *pulumi.OutputState }

func (CreateOrUpdateDryrunParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateOrUpdateDryrunParameters)(nil)).Elem()
}

func (o CreateOrUpdateDryrunParametersOutput) ToCreateOrUpdateDryrunParametersOutput() CreateOrUpdateDryrunParametersOutput {
	return o
}

func (o CreateOrUpdateDryrunParametersOutput) ToCreateOrUpdateDryrunParametersOutputWithContext(ctx context.Context) CreateOrUpdateDryrunParametersOutput {
	return o
}

func (o CreateOrUpdateDryrunParametersOutput) ToCreateOrUpdateDryrunParametersPtrOutput() CreateOrUpdateDryrunParametersPtrOutput {
	return o.ToCreateOrUpdateDryrunParametersPtrOutputWithContext(context.Background())
}

func (o CreateOrUpdateDryrunParametersOutput) ToCreateOrUpdateDryrunParametersPtrOutputWithContext(ctx context.Context) CreateOrUpdateDryrunParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateOrUpdateDryrunParameters) *CreateOrUpdateDryrunParameters {
		return &v
	}).(CreateOrUpdateDryrunParametersPtrOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) ActionName() pulumi.StringOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) string { return v.ActionName }).(pulumi.StringOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) interface{} { return v.AuthInfo }).(pulumi.AnyOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) *string { return v.ClientType }).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) ConfigurationInfo() ConfigurationInfoPtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) *ConfigurationInfo { return v.ConfigurationInfo }).(ConfigurationInfoPtrOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) PublicNetworkSolution() PublicNetworkSolutionPtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) *PublicNetworkSolution { return v.PublicNetworkSolution }).(PublicNetworkSolutionPtrOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) SecretStore() SecretStorePtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) *SecretStore { return v.SecretStore }).(SecretStorePtrOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) interface{} { return v.TargetService }).(pulumi.AnyOutput)
}

func (o CreateOrUpdateDryrunParametersOutput) VNetSolution() VNetSolutionPtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParameters) *VNetSolution { return v.VNetSolution }).(VNetSolutionPtrOutput)
}

type CreateOrUpdateDryrunParametersPtrOutput struct{ *pulumi.OutputState }

func (CreateOrUpdateDryrunParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateOrUpdateDryrunParameters)(nil)).Elem()
}

func (o CreateOrUpdateDryrunParametersPtrOutput) ToCreateOrUpdateDryrunParametersPtrOutput() CreateOrUpdateDryrunParametersPtrOutput {
	return o
}

func (o CreateOrUpdateDryrunParametersPtrOutput) ToCreateOrUpdateDryrunParametersPtrOutputWithContext(ctx context.Context) CreateOrUpdateDryrunParametersPtrOutput {
	return o
}

func (o CreateOrUpdateDryrunParametersPtrOutput) Elem() CreateOrUpdateDryrunParametersOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) CreateOrUpdateDryrunParameters {
		if v != nil {
			return *v
		}
		var ret CreateOrUpdateDryrunParameters
		return ret
	}).(CreateOrUpdateDryrunParametersOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) ActionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) *string {
		if v == nil {
			return nil
		}
		return &v.ActionName
	}).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) interface{} {
		if v == nil {
			return nil
		}
		return v.AuthInfo
	}).(pulumi.AnyOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) *string {
		if v == nil {
			return nil
		}
		return v.ClientType
	}).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) ConfigurationInfo() ConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) *ConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.ConfigurationInfo
	}).(ConfigurationInfoPtrOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) PublicNetworkSolution() PublicNetworkSolutionPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) *PublicNetworkSolution {
		if v == nil {
			return nil
		}
		return v.PublicNetworkSolution
	}).(PublicNetworkSolutionPtrOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) SecretStore() SecretStorePtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) *SecretStore {
		if v == nil {
			return nil
		}
		return v.SecretStore
	}).(SecretStorePtrOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) interface{} {
		if v == nil {
			return nil
		}
		return v.TargetService
	}).(pulumi.AnyOutput)
}

func (o CreateOrUpdateDryrunParametersPtrOutput) VNetSolution() VNetSolutionPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParameters) *VNetSolution {
		if v == nil {
			return nil
		}
		return v.VNetSolution
	}).(VNetSolutionPtrOutput)
}

type CreateOrUpdateDryrunParametersResponse struct {
	ActionName            string                         `pulumi:"actionName"`
	AuthInfo              interface{}                    `pulumi:"authInfo"`
	ClientType            *string                        `pulumi:"clientType"`
	ConfigurationInfo     *ConfigurationInfoResponse     `pulumi:"configurationInfo"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
	PublicNetworkSolution *PublicNetworkSolutionResponse `pulumi:"publicNetworkSolution"`
	Scope                 *string                        `pulumi:"scope"`
	SecretStore           *SecretStoreResponse           `pulumi:"secretStore"`
	TargetService         interface{}                    `pulumi:"targetService"`
	VNetSolution          *VNetSolutionResponse          `pulumi:"vNetSolution"`
}

type CreateOrUpdateDryrunParametersResponseOutput struct{ *pulumi.OutputState }

func (CreateOrUpdateDryrunParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateOrUpdateDryrunParametersResponse)(nil)).Elem()
}

func (o CreateOrUpdateDryrunParametersResponseOutput) ToCreateOrUpdateDryrunParametersResponseOutput() CreateOrUpdateDryrunParametersResponseOutput {
	return o
}

func (o CreateOrUpdateDryrunParametersResponseOutput) ToCreateOrUpdateDryrunParametersResponseOutputWithContext(ctx context.Context) CreateOrUpdateDryrunParametersResponseOutput {
	return o
}

func (o CreateOrUpdateDryrunParametersResponseOutput) ActionName() pulumi.StringOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) string { return v.ActionName }).(pulumi.StringOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) interface{} { return v.AuthInfo }).(pulumi.AnyOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) *string { return v.ClientType }).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) ConfigurationInfo() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) *ConfigurationInfoResponse { return v.ConfigurationInfo }).(ConfigurationInfoResponsePtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) PublicNetworkSolution() PublicNetworkSolutionResponsePtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) *PublicNetworkSolutionResponse {
		return v.PublicNetworkSolution
	}).(PublicNetworkSolutionResponsePtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) *SecretStoreResponse { return v.SecretStore }).(SecretStoreResponsePtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) interface{} { return v.TargetService }).(pulumi.AnyOutput)
}

func (o CreateOrUpdateDryrunParametersResponseOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v CreateOrUpdateDryrunParametersResponse) *VNetSolutionResponse { return v.VNetSolution }).(VNetSolutionResponsePtrOutput)
}

type CreateOrUpdateDryrunParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (CreateOrUpdateDryrunParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateOrUpdateDryrunParametersResponse)(nil)).Elem()
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) ToCreateOrUpdateDryrunParametersResponsePtrOutput() CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) ToCreateOrUpdateDryrunParametersResponsePtrOutputWithContext(ctx context.Context) CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) Elem() CreateOrUpdateDryrunParametersResponseOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) CreateOrUpdateDryrunParametersResponse {
		if v != nil {
			return *v
		}
		var ret CreateOrUpdateDryrunParametersResponse
		return ret
	}).(CreateOrUpdateDryrunParametersResponseOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) ActionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionName
	}).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.AuthInfo
	}).(pulumi.AnyOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientType
	}).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) ConfigurationInfo() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) *ConfigurationInfoResponse {
		if v == nil {
			return nil
		}
		return v.ConfigurationInfo
	}).(ConfigurationInfoResponsePtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) PublicNetworkSolution() PublicNetworkSolutionResponsePtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) *PublicNetworkSolutionResponse {
		if v == nil {
			return nil
		}
		return v.PublicNetworkSolution
	}).(PublicNetworkSolutionResponsePtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) *SecretStoreResponse {
		if v == nil {
			return nil
		}
		return v.SecretStore
	}).(SecretStoreResponsePtrOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.TargetService
	}).(pulumi.AnyOutput)
}

func (o CreateOrUpdateDryrunParametersResponsePtrOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v *CreateOrUpdateDryrunParametersResponse) *VNetSolutionResponse {
		if v == nil {
			return nil
		}
		return v.VNetSolution
	}).(VNetSolutionResponsePtrOutput)
}

type DryrunOperationPreviewResponse struct {
	Action        *string `pulumi:"action"`
	Description   *string `pulumi:"description"`
	Name          *string `pulumi:"name"`
	OperationType *string `pulumi:"operationType"`
	Scope         *string `pulumi:"scope"`
}

type DryrunOperationPreviewResponseOutput struct{ *pulumi.OutputState }

func (DryrunOperationPreviewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DryrunOperationPreviewResponse)(nil)).Elem()
}

func (o DryrunOperationPreviewResponseOutput) ToDryrunOperationPreviewResponseOutput() DryrunOperationPreviewResponseOutput {
	return o
}

func (o DryrunOperationPreviewResponseOutput) ToDryrunOperationPreviewResponseOutputWithContext(ctx context.Context) DryrunOperationPreviewResponseOutput {
	return o
}

func (o DryrunOperationPreviewResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DryrunOperationPreviewResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o DryrunOperationPreviewResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DryrunOperationPreviewResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DryrunOperationPreviewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DryrunOperationPreviewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DryrunOperationPreviewResponseOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DryrunOperationPreviewResponse) *string { return v.OperationType }).(pulumi.StringPtrOutput)
}

func (o DryrunOperationPreviewResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DryrunOperationPreviewResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type DryrunOperationPreviewResponseArrayOutput struct{ *pulumi.OutputState }

func (DryrunOperationPreviewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DryrunOperationPreviewResponse)(nil)).Elem()
}

func (o DryrunOperationPreviewResponseArrayOutput) ToDryrunOperationPreviewResponseArrayOutput() DryrunOperationPreviewResponseArrayOutput {
	return o
}

func (o DryrunOperationPreviewResponseArrayOutput) ToDryrunOperationPreviewResponseArrayOutputWithContext(ctx context.Context) DryrunOperationPreviewResponseArrayOutput {
	return o
}

func (o DryrunOperationPreviewResponseArrayOutput) Index(i pulumi.IntInput) DryrunOperationPreviewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DryrunOperationPreviewResponse {
		return vs[0].([]DryrunOperationPreviewResponse)[vs[1].(int)]
	}).(DryrunOperationPreviewResponseOutput)
}

type FirewallRules struct {
	AzureServices  *string  `pulumi:"azureServices"`
	CallerClientIP *string  `pulumi:"callerClientIP"`
	IpRanges       []string `pulumi:"ipRanges"`
}





type FirewallRulesInput interface {
	pulumi.Input

	ToFirewallRulesOutput() FirewallRulesOutput
	ToFirewallRulesOutputWithContext(context.Context) FirewallRulesOutput
}

type FirewallRulesArgs struct {
	AzureServices  pulumi.StringPtrInput   `pulumi:"azureServices"`
	CallerClientIP pulumi.StringPtrInput   `pulumi:"callerClientIP"`
	IpRanges       pulumi.StringArrayInput `pulumi:"ipRanges"`
}

func (FirewallRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRules)(nil)).Elem()
}

func (i FirewallRulesArgs) ToFirewallRulesOutput() FirewallRulesOutput {
	return i.ToFirewallRulesOutputWithContext(context.Background())
}

func (i FirewallRulesArgs) ToFirewallRulesOutputWithContext(ctx context.Context) FirewallRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesOutput)
}

func (i FirewallRulesArgs) ToFirewallRulesPtrOutput() FirewallRulesPtrOutput {
	return i.ToFirewallRulesPtrOutputWithContext(context.Background())
}

func (i FirewallRulesArgs) ToFirewallRulesPtrOutputWithContext(ctx context.Context) FirewallRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesOutput).ToFirewallRulesPtrOutputWithContext(ctx)
}









type FirewallRulesPtrInput interface {
	pulumi.Input

	ToFirewallRulesPtrOutput() FirewallRulesPtrOutput
	ToFirewallRulesPtrOutputWithContext(context.Context) FirewallRulesPtrOutput
}

type firewallRulesPtrType FirewallRulesArgs

func FirewallRulesPtr(v *FirewallRulesArgs) FirewallRulesPtrInput {
	return (*firewallRulesPtrType)(v)
}

func (*firewallRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRules)(nil)).Elem()
}

func (i *firewallRulesPtrType) ToFirewallRulesPtrOutput() FirewallRulesPtrOutput {
	return i.ToFirewallRulesPtrOutputWithContext(context.Background())
}

func (i *firewallRulesPtrType) ToFirewallRulesPtrOutputWithContext(ctx context.Context) FirewallRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRulesPtrOutput)
}

type FirewallRulesOutput struct{ *pulumi.OutputState }

func (FirewallRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRules)(nil)).Elem()
}

func (o FirewallRulesOutput) ToFirewallRulesOutput() FirewallRulesOutput {
	return o
}

func (o FirewallRulesOutput) ToFirewallRulesOutputWithContext(ctx context.Context) FirewallRulesOutput {
	return o
}

func (o FirewallRulesOutput) ToFirewallRulesPtrOutput() FirewallRulesPtrOutput {
	return o.ToFirewallRulesPtrOutputWithContext(context.Background())
}

func (o FirewallRulesOutput) ToFirewallRulesPtrOutputWithContext(ctx context.Context) FirewallRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallRules) *FirewallRules {
		return &v
	}).(FirewallRulesPtrOutput)
}

func (o FirewallRulesOutput) AzureServices() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRules) *string { return v.AzureServices }).(pulumi.StringPtrOutput)
}

func (o FirewallRulesOutput) CallerClientIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRules) *string { return v.CallerClientIP }).(pulumi.StringPtrOutput)
}

func (o FirewallRulesOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FirewallRules) []string { return v.IpRanges }).(pulumi.StringArrayOutput)
}

type FirewallRulesPtrOutput struct{ *pulumi.OutputState }

func (FirewallRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRules)(nil)).Elem()
}

func (o FirewallRulesPtrOutput) ToFirewallRulesPtrOutput() FirewallRulesPtrOutput {
	return o
}

func (o FirewallRulesPtrOutput) ToFirewallRulesPtrOutputWithContext(ctx context.Context) FirewallRulesPtrOutput {
	return o
}

func (o FirewallRulesPtrOutput) Elem() FirewallRulesOutput {
	return o.ApplyT(func(v *FirewallRules) FirewallRules {
		if v != nil {
			return *v
		}
		var ret FirewallRules
		return ret
	}).(FirewallRulesOutput)
}

func (o FirewallRulesPtrOutput) AzureServices() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRules) *string {
		if v == nil {
			return nil
		}
		return v.AzureServices
	}).(pulumi.StringPtrOutput)
}

func (o FirewallRulesPtrOutput) CallerClientIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRules) *string {
		if v == nil {
			return nil
		}
		return v.CallerClientIP
	}).(pulumi.StringPtrOutput)
}

func (o FirewallRulesPtrOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallRules) []string {
		if v == nil {
			return nil
		}
		return v.IpRanges
	}).(pulumi.StringArrayOutput)
}

type FirewallRulesResponse struct {
	AzureServices  *string  `pulumi:"azureServices"`
	CallerClientIP *string  `pulumi:"callerClientIP"`
	IpRanges       []string `pulumi:"ipRanges"`
}

type FirewallRulesResponseOutput struct{ *pulumi.OutputState }

func (FirewallRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRulesResponse)(nil)).Elem()
}

func (o FirewallRulesResponseOutput) ToFirewallRulesResponseOutput() FirewallRulesResponseOutput {
	return o
}

func (o FirewallRulesResponseOutput) ToFirewallRulesResponseOutputWithContext(ctx context.Context) FirewallRulesResponseOutput {
	return o
}

func (o FirewallRulesResponseOutput) AzureServices() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRulesResponse) *string { return v.AzureServices }).(pulumi.StringPtrOutput)
}

func (o FirewallRulesResponseOutput) CallerClientIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FirewallRulesResponse) *string { return v.CallerClientIP }).(pulumi.StringPtrOutput)
}

func (o FirewallRulesResponseOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FirewallRulesResponse) []string { return v.IpRanges }).(pulumi.StringArrayOutput)
}

type FirewallRulesResponsePtrOutput struct{ *pulumi.OutputState }

func (FirewallRulesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRulesResponse)(nil)).Elem()
}

func (o FirewallRulesResponsePtrOutput) ToFirewallRulesResponsePtrOutput() FirewallRulesResponsePtrOutput {
	return o
}

func (o FirewallRulesResponsePtrOutput) ToFirewallRulesResponsePtrOutputWithContext(ctx context.Context) FirewallRulesResponsePtrOutput {
	return o
}

func (o FirewallRulesResponsePtrOutput) Elem() FirewallRulesResponseOutput {
	return o.ApplyT(func(v *FirewallRulesResponse) FirewallRulesResponse {
		if v != nil {
			return *v
		}
		var ret FirewallRulesResponse
		return ret
	}).(FirewallRulesResponseOutput)
}

func (o FirewallRulesResponsePtrOutput) AzureServices() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRulesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureServices
	}).(pulumi.StringPtrOutput)
}

func (o FirewallRulesResponsePtrOutput) CallerClientIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRulesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CallerClientIP
	}).(pulumi.StringPtrOutput)
}

func (o FirewallRulesResponsePtrOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallRulesResponse) []string {
		if v == nil {
			return nil
		}
		return v.IpRanges
	}).(pulumi.StringArrayOutput)
}

type KeyVaultSecretReferenceSecretInfo struct {
	Name       *string `pulumi:"name"`
	SecretType string  `pulumi:"secretType"`
	Version    *string `pulumi:"version"`
}

type KeyVaultSecretReferenceSecretInfoResponse struct {
	Name       *string `pulumi:"name"`
	SecretType string  `pulumi:"secretType"`
	Version    *string `pulumi:"version"`
}

type KeyVaultSecretUriSecretInfo struct {
	SecretType string  `pulumi:"secretType"`
	Value      *string `pulumi:"value"`
}

type KeyVaultSecretUriSecretInfoResponse struct {
	SecretType string  `pulumi:"secretType"`
	Value      *string `pulumi:"value"`
}

type PermissionsMissingDryrunPrerequisiteResultResponse struct {
	Permissions     []string `pulumi:"permissions"`
	RecommendedRole *string  `pulumi:"recommendedRole"`
	Scope           *string  `pulumi:"scope"`
	Type            string   `pulumi:"type"`
}

type PublicNetworkSolution struct {
	Action                 *string        `pulumi:"action"`
	DeleteOrUpdateBehavior *string        `pulumi:"deleteOrUpdateBehavior"`
	FirewallRules          *FirewallRules `pulumi:"firewallRules"`
}





type PublicNetworkSolutionInput interface {
	pulumi.Input

	ToPublicNetworkSolutionOutput() PublicNetworkSolutionOutput
	ToPublicNetworkSolutionOutputWithContext(context.Context) PublicNetworkSolutionOutput
}

type PublicNetworkSolutionArgs struct {
	Action                 pulumi.StringPtrInput `pulumi:"action"`
	DeleteOrUpdateBehavior pulumi.StringPtrInput `pulumi:"deleteOrUpdateBehavior"`
	FirewallRules          FirewallRulesPtrInput `pulumi:"firewallRules"`
}

func (PublicNetworkSolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkSolution)(nil)).Elem()
}

func (i PublicNetworkSolutionArgs) ToPublicNetworkSolutionOutput() PublicNetworkSolutionOutput {
	return i.ToPublicNetworkSolutionOutputWithContext(context.Background())
}

func (i PublicNetworkSolutionArgs) ToPublicNetworkSolutionOutputWithContext(ctx context.Context) PublicNetworkSolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicNetworkSolutionOutput)
}

func (i PublicNetworkSolutionArgs) ToPublicNetworkSolutionPtrOutput() PublicNetworkSolutionPtrOutput {
	return i.ToPublicNetworkSolutionPtrOutputWithContext(context.Background())
}

func (i PublicNetworkSolutionArgs) ToPublicNetworkSolutionPtrOutputWithContext(ctx context.Context) PublicNetworkSolutionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicNetworkSolutionOutput).ToPublicNetworkSolutionPtrOutputWithContext(ctx)
}









type PublicNetworkSolutionPtrInput interface {
	pulumi.Input

	ToPublicNetworkSolutionPtrOutput() PublicNetworkSolutionPtrOutput
	ToPublicNetworkSolutionPtrOutputWithContext(context.Context) PublicNetworkSolutionPtrOutput
}

type publicNetworkSolutionPtrType PublicNetworkSolutionArgs

func PublicNetworkSolutionPtr(v *PublicNetworkSolutionArgs) PublicNetworkSolutionPtrInput {
	return (*publicNetworkSolutionPtrType)(v)
}

func (*publicNetworkSolutionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkSolution)(nil)).Elem()
}

func (i *publicNetworkSolutionPtrType) ToPublicNetworkSolutionPtrOutput() PublicNetworkSolutionPtrOutput {
	return i.ToPublicNetworkSolutionPtrOutputWithContext(context.Background())
}

func (i *publicNetworkSolutionPtrType) ToPublicNetworkSolutionPtrOutputWithContext(ctx context.Context) PublicNetworkSolutionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicNetworkSolutionPtrOutput)
}

type PublicNetworkSolutionOutput struct{ *pulumi.OutputState }

func (PublicNetworkSolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkSolution)(nil)).Elem()
}

func (o PublicNetworkSolutionOutput) ToPublicNetworkSolutionOutput() PublicNetworkSolutionOutput {
	return o
}

func (o PublicNetworkSolutionOutput) ToPublicNetworkSolutionOutputWithContext(ctx context.Context) PublicNetworkSolutionOutput {
	return o
}

func (o PublicNetworkSolutionOutput) ToPublicNetworkSolutionPtrOutput() PublicNetworkSolutionPtrOutput {
	return o.ToPublicNetworkSolutionPtrOutputWithContext(context.Background())
}

func (o PublicNetworkSolutionOutput) ToPublicNetworkSolutionPtrOutputWithContext(ctx context.Context) PublicNetworkSolutionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkSolution) *PublicNetworkSolution {
		return &v
	}).(PublicNetworkSolutionPtrOutput)
}

func (o PublicNetworkSolutionOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicNetworkSolution) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o PublicNetworkSolutionOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicNetworkSolution) *string { return v.DeleteOrUpdateBehavior }).(pulumi.StringPtrOutput)
}

func (o PublicNetworkSolutionOutput) FirewallRules() FirewallRulesPtrOutput {
	return o.ApplyT(func(v PublicNetworkSolution) *FirewallRules { return v.FirewallRules }).(FirewallRulesPtrOutput)
}

type PublicNetworkSolutionPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkSolutionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkSolution)(nil)).Elem()
}

func (o PublicNetworkSolutionPtrOutput) ToPublicNetworkSolutionPtrOutput() PublicNetworkSolutionPtrOutput {
	return o
}

func (o PublicNetworkSolutionPtrOutput) ToPublicNetworkSolutionPtrOutputWithContext(ctx context.Context) PublicNetworkSolutionPtrOutput {
	return o
}

func (o PublicNetworkSolutionPtrOutput) Elem() PublicNetworkSolutionOutput {
	return o.ApplyT(func(v *PublicNetworkSolution) PublicNetworkSolution {
		if v != nil {
			return *v
		}
		var ret PublicNetworkSolution
		return ret
	}).(PublicNetworkSolutionOutput)
}

func (o PublicNetworkSolutionPtrOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicNetworkSolution) *string {
		if v == nil {
			return nil
		}
		return v.Action
	}).(pulumi.StringPtrOutput)
}

func (o PublicNetworkSolutionPtrOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicNetworkSolution) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOrUpdateBehavior
	}).(pulumi.StringPtrOutput)
}

func (o PublicNetworkSolutionPtrOutput) FirewallRules() FirewallRulesPtrOutput {
	return o.ApplyT(func(v *PublicNetworkSolution) *FirewallRules {
		if v == nil {
			return nil
		}
		return v.FirewallRules
	}).(FirewallRulesPtrOutput)
}

type PublicNetworkSolutionResponse struct {
	Action                 *string                `pulumi:"action"`
	DeleteOrUpdateBehavior *string                `pulumi:"deleteOrUpdateBehavior"`
	FirewallRules          *FirewallRulesResponse `pulumi:"firewallRules"`
}

type PublicNetworkSolutionResponseOutput struct{ *pulumi.OutputState }

func (PublicNetworkSolutionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkSolutionResponse)(nil)).Elem()
}

func (o PublicNetworkSolutionResponseOutput) ToPublicNetworkSolutionResponseOutput() PublicNetworkSolutionResponseOutput {
	return o
}

func (o PublicNetworkSolutionResponseOutput) ToPublicNetworkSolutionResponseOutputWithContext(ctx context.Context) PublicNetworkSolutionResponseOutput {
	return o
}

func (o PublicNetworkSolutionResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicNetworkSolutionResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o PublicNetworkSolutionResponseOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicNetworkSolutionResponse) *string { return v.DeleteOrUpdateBehavior }).(pulumi.StringPtrOutput)
}

func (o PublicNetworkSolutionResponseOutput) FirewallRules() FirewallRulesResponsePtrOutput {
	return o.ApplyT(func(v PublicNetworkSolutionResponse) *FirewallRulesResponse { return v.FirewallRules }).(FirewallRulesResponsePtrOutput)
}

type PublicNetworkSolutionResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkSolutionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkSolutionResponse)(nil)).Elem()
}

func (o PublicNetworkSolutionResponsePtrOutput) ToPublicNetworkSolutionResponsePtrOutput() PublicNetworkSolutionResponsePtrOutput {
	return o
}

func (o PublicNetworkSolutionResponsePtrOutput) ToPublicNetworkSolutionResponsePtrOutputWithContext(ctx context.Context) PublicNetworkSolutionResponsePtrOutput {
	return o
}

func (o PublicNetworkSolutionResponsePtrOutput) Elem() PublicNetworkSolutionResponseOutput {
	return o.ApplyT(func(v *PublicNetworkSolutionResponse) PublicNetworkSolutionResponse {
		if v != nil {
			return *v
		}
		var ret PublicNetworkSolutionResponse
		return ret
	}).(PublicNetworkSolutionResponseOutput)
}

func (o PublicNetworkSolutionResponsePtrOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicNetworkSolutionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Action
	}).(pulumi.StringPtrOutput)
}

func (o PublicNetworkSolutionResponsePtrOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicNetworkSolutionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOrUpdateBehavior
	}).(pulumi.StringPtrOutput)
}

func (o PublicNetworkSolutionResponsePtrOutput) FirewallRules() FirewallRulesResponsePtrOutput {
	return o.ApplyT(func(v *PublicNetworkSolutionResponse) *FirewallRulesResponse {
		if v == nil {
			return nil
		}
		return v.FirewallRules
	}).(FirewallRulesResponsePtrOutput)
}

type SecretAuthInfo struct {
	AuthType   string      `pulumi:"authType"`
	Name       *string     `pulumi:"name"`
	SecretInfo interface{} `pulumi:"secretInfo"`
}

type SecretAuthInfoResponse struct {
	AuthType   string      `pulumi:"authType"`
	Name       *string     `pulumi:"name"`
	SecretInfo interface{} `pulumi:"secretInfo"`
}

type SecretStore struct {
	KeyVaultId         *string `pulumi:"keyVaultId"`
	KeyVaultSecretName *string `pulumi:"keyVaultSecretName"`
}





type SecretStoreInput interface {
	pulumi.Input

	ToSecretStoreOutput() SecretStoreOutput
	ToSecretStoreOutputWithContext(context.Context) SecretStoreOutput
}

type SecretStoreArgs struct {
	KeyVaultId         pulumi.StringPtrInput `pulumi:"keyVaultId"`
	KeyVaultSecretName pulumi.StringPtrInput `pulumi:"keyVaultSecretName"`
}

func (SecretStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStore)(nil)).Elem()
}

func (i SecretStoreArgs) ToSecretStoreOutput() SecretStoreOutput {
	return i.ToSecretStoreOutputWithContext(context.Background())
}

func (i SecretStoreArgs) ToSecretStoreOutputWithContext(ctx context.Context) SecretStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStoreOutput)
}

func (i SecretStoreArgs) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return i.ToSecretStorePtrOutputWithContext(context.Background())
}

func (i SecretStoreArgs) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStoreOutput).ToSecretStorePtrOutputWithContext(ctx)
}









type SecretStorePtrInput interface {
	pulumi.Input

	ToSecretStorePtrOutput() SecretStorePtrOutput
	ToSecretStorePtrOutputWithContext(context.Context) SecretStorePtrOutput
}

type secretStorePtrType SecretStoreArgs

func SecretStorePtr(v *SecretStoreArgs) SecretStorePtrInput {
	return (*secretStorePtrType)(v)
}

func (*secretStorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStore)(nil)).Elem()
}

func (i *secretStorePtrType) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return i.ToSecretStorePtrOutputWithContext(context.Background())
}

func (i *secretStorePtrType) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStorePtrOutput)
}

type SecretStoreOutput struct{ *pulumi.OutputState }

func (SecretStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStore)(nil)).Elem()
}

func (o SecretStoreOutput) ToSecretStoreOutput() SecretStoreOutput {
	return o
}

func (o SecretStoreOutput) ToSecretStoreOutputWithContext(ctx context.Context) SecretStoreOutput {
	return o
}

func (o SecretStoreOutput) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return o.ToSecretStorePtrOutputWithContext(context.Background())
}

func (o SecretStoreOutput) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretStore) *SecretStore {
		return &v
	}).(SecretStorePtrOutput)
}

func (o SecretStoreOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretStore) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o SecretStoreOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretStore) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

type SecretStorePtrOutput struct{ *pulumi.OutputState }

func (SecretStorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStore)(nil)).Elem()
}

func (o SecretStorePtrOutput) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return o
}

func (o SecretStorePtrOutput) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return o
}

func (o SecretStorePtrOutput) Elem() SecretStoreOutput {
	return o.ApplyT(func(v *SecretStore) SecretStore {
		if v != nil {
			return *v
		}
		var ret SecretStore
		return ret
	}).(SecretStoreOutput)
}

func (o SecretStorePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStore) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

func (o SecretStorePtrOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStore) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultSecretName
	}).(pulumi.StringPtrOutput)
}

type SecretStoreResponse struct {
	KeyVaultId         *string `pulumi:"keyVaultId"`
	KeyVaultSecretName *string `pulumi:"keyVaultSecretName"`
}

type SecretStoreResponseOutput struct{ *pulumi.OutputState }

func (SecretStoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStoreResponse)(nil)).Elem()
}

func (o SecretStoreResponseOutput) ToSecretStoreResponseOutput() SecretStoreResponseOutput {
	return o
}

func (o SecretStoreResponseOutput) ToSecretStoreResponseOutputWithContext(ctx context.Context) SecretStoreResponseOutput {
	return o
}

func (o SecretStoreResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretStoreResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o SecretStoreResponseOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretStoreResponse) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

type SecretStoreResponsePtrOutput struct{ *pulumi.OutputState }

func (SecretStoreResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStoreResponse)(nil)).Elem()
}

func (o SecretStoreResponsePtrOutput) ToSecretStoreResponsePtrOutput() SecretStoreResponsePtrOutput {
	return o
}

func (o SecretStoreResponsePtrOutput) ToSecretStoreResponsePtrOutputWithContext(ctx context.Context) SecretStoreResponsePtrOutput {
	return o
}

func (o SecretStoreResponsePtrOutput) Elem() SecretStoreResponseOutput {
	return o.ApplyT(func(v *SecretStoreResponse) SecretStoreResponse {
		if v != nil {
			return *v
		}
		var ret SecretStoreResponse
		return ret
	}).(SecretStoreResponseOutput)
}

func (o SecretStoreResponsePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStoreResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

func (o SecretStoreResponsePtrOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStoreResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultSecretName
	}).(pulumi.StringPtrOutput)
}

type SelfHostedServer struct {
	Endpoint *string `pulumi:"endpoint"`
	Type     string  `pulumi:"type"`
}

type SelfHostedServerResponse struct {
	Endpoint *string `pulumi:"endpoint"`
	Type     string  `pulumi:"type"`
}

type ServicePrincipalCertificateAuthInfo struct {
	AuthType               string   `pulumi:"authType"`
	Certificate            string   `pulumi:"certificate"`
	ClientId               string   `pulumi:"clientId"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	PrincipalId            string   `pulumi:"principalId"`
	Roles                  []string `pulumi:"roles"`
}

type ServicePrincipalCertificateAuthInfoResponse struct {
	AuthType               string   `pulumi:"authType"`
	Certificate            string   `pulumi:"certificate"`
	ClientId               string   `pulumi:"clientId"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	PrincipalId            string   `pulumi:"principalId"`
	Roles                  []string `pulumi:"roles"`
}

type ServicePrincipalSecretAuthInfo struct {
	AuthType               string   `pulumi:"authType"`
	ClientId               string   `pulumi:"clientId"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	PrincipalId            string   `pulumi:"principalId"`
	Roles                  []string `pulumi:"roles"`
	Secret                 string   `pulumi:"secret"`
	UserName               *string  `pulumi:"userName"`
}

type ServicePrincipalSecretAuthInfoResponse struct {
	AuthType               string   `pulumi:"authType"`
	ClientId               string   `pulumi:"clientId"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	PrincipalId            string   `pulumi:"principalId"`
	Roles                  []string `pulumi:"roles"`
	Secret                 string   `pulumi:"secret"`
	UserName               *string  `pulumi:"userName"`
}

type SourceConfigurationResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type SourceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SourceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceConfigurationResponse)(nil)).Elem()
}

func (o SourceConfigurationResponseOutput) ToSourceConfigurationResponseOutput() SourceConfigurationResponseOutput {
	return o
}

func (o SourceConfigurationResponseOutput) ToSourceConfigurationResponseOutputWithContext(ctx context.Context) SourceConfigurationResponseOutput {
	return o
}

func (o SourceConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SourceConfigurationResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceConfigurationResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SourceConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (SourceConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceConfigurationResponse)(nil)).Elem()
}

func (o SourceConfigurationResponseArrayOutput) ToSourceConfigurationResponseArrayOutput() SourceConfigurationResponseArrayOutput {
	return o
}

func (o SourceConfigurationResponseArrayOutput) ToSourceConfigurationResponseArrayOutputWithContext(ctx context.Context) SourceConfigurationResponseArrayOutput {
	return o
}

func (o SourceConfigurationResponseArrayOutput) Index(i pulumi.IntInput) SourceConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceConfigurationResponse {
		return vs[0].([]SourceConfigurationResponse)[vs[1].(int)]
	}).(SourceConfigurationResponseOutput)
}

type SystemAssignedIdentityAuthInfo struct {
	AuthType               string   `pulumi:"authType"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	Roles                  []string `pulumi:"roles"`
	UserName               *string  `pulumi:"userName"`
}

type SystemAssignedIdentityAuthInfoResponse struct {
	AuthType               string   `pulumi:"authType"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	Roles                  []string `pulumi:"roles"`
	UserName               *string  `pulumi:"userName"`
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

type UserAccountAuthInfo struct {
	AuthType               string   `pulumi:"authType"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	PrincipalId            *string  `pulumi:"principalId"`
	Roles                  []string `pulumi:"roles"`
	UserName               *string  `pulumi:"userName"`
}

type UserAccountAuthInfoResponse struct {
	AuthType               string   `pulumi:"authType"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	PrincipalId            *string  `pulumi:"principalId"`
	Roles                  []string `pulumi:"roles"`
	UserName               *string  `pulumi:"userName"`
}

type UserAssignedIdentityAuthInfo struct {
	AuthType               string   `pulumi:"authType"`
	ClientId               *string  `pulumi:"clientId"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	Roles                  []string `pulumi:"roles"`
	SubscriptionId         *string  `pulumi:"subscriptionId"`
	UserName               *string  `pulumi:"userName"`
}

type UserAssignedIdentityAuthInfoResponse struct {
	AuthType               string   `pulumi:"authType"`
	ClientId               *string  `pulumi:"clientId"`
	DeleteOrUpdateBehavior *string  `pulumi:"deleteOrUpdateBehavior"`
	Roles                  []string `pulumi:"roles"`
	SubscriptionId         *string  `pulumi:"subscriptionId"`
	UserName               *string  `pulumi:"userName"`
}

type VNetSolution struct {
	DeleteOrUpdateBehavior *string `pulumi:"deleteOrUpdateBehavior"`
	Type                   *string `pulumi:"type"`
}





type VNetSolutionInput interface {
	pulumi.Input

	ToVNetSolutionOutput() VNetSolutionOutput
	ToVNetSolutionOutputWithContext(context.Context) VNetSolutionOutput
}

type VNetSolutionArgs struct {
	DeleteOrUpdateBehavior pulumi.StringPtrInput `pulumi:"deleteOrUpdateBehavior"`
	Type                   pulumi.StringPtrInput `pulumi:"type"`
}

func (VNetSolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolution)(nil)).Elem()
}

func (i VNetSolutionArgs) ToVNetSolutionOutput() VNetSolutionOutput {
	return i.ToVNetSolutionOutputWithContext(context.Background())
}

func (i VNetSolutionArgs) ToVNetSolutionOutputWithContext(ctx context.Context) VNetSolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionOutput)
}

func (i VNetSolutionArgs) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return i.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (i VNetSolutionArgs) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionOutput).ToVNetSolutionPtrOutputWithContext(ctx)
}









type VNetSolutionPtrInput interface {
	pulumi.Input

	ToVNetSolutionPtrOutput() VNetSolutionPtrOutput
	ToVNetSolutionPtrOutputWithContext(context.Context) VNetSolutionPtrOutput
}

type vnetSolutionPtrType VNetSolutionArgs

func VNetSolutionPtr(v *VNetSolutionArgs) VNetSolutionPtrInput {
	return (*vnetSolutionPtrType)(v)
}

func (*vnetSolutionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolution)(nil)).Elem()
}

func (i *vnetSolutionPtrType) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return i.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (i *vnetSolutionPtrType) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionPtrOutput)
}

type VNetSolutionOutput struct{ *pulumi.OutputState }

func (VNetSolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolution)(nil)).Elem()
}

func (o VNetSolutionOutput) ToVNetSolutionOutput() VNetSolutionOutput {
	return o
}

func (o VNetSolutionOutput) ToVNetSolutionOutputWithContext(ctx context.Context) VNetSolutionOutput {
	return o
}

func (o VNetSolutionOutput) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return o.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (o VNetSolutionOutput) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VNetSolution) *VNetSolution {
		return &v
	}).(VNetSolutionPtrOutput)
}

func (o VNetSolutionOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VNetSolution) *string { return v.DeleteOrUpdateBehavior }).(pulumi.StringPtrOutput)
}

func (o VNetSolutionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VNetSolution) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VNetSolutionPtrOutput struct{ *pulumi.OutputState }

func (VNetSolutionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolution)(nil)).Elem()
}

func (o VNetSolutionPtrOutput) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return o
}

func (o VNetSolutionPtrOutput) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return o
}

func (o VNetSolutionPtrOutput) Elem() VNetSolutionOutput {
	return o.ApplyT(func(v *VNetSolution) VNetSolution {
		if v != nil {
			return *v
		}
		var ret VNetSolution
		return ret
	}).(VNetSolutionOutput)
}

func (o VNetSolutionPtrOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VNetSolution) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOrUpdateBehavior
	}).(pulumi.StringPtrOutput)
}

func (o VNetSolutionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VNetSolution) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VNetSolutionResponse struct {
	DeleteOrUpdateBehavior *string `pulumi:"deleteOrUpdateBehavior"`
	Type                   *string `pulumi:"type"`
}

type VNetSolutionResponseOutput struct{ *pulumi.OutputState }

func (VNetSolutionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolutionResponse)(nil)).Elem()
}

func (o VNetSolutionResponseOutput) ToVNetSolutionResponseOutput() VNetSolutionResponseOutput {
	return o
}

func (o VNetSolutionResponseOutput) ToVNetSolutionResponseOutputWithContext(ctx context.Context) VNetSolutionResponseOutput {
	return o
}

func (o VNetSolutionResponseOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VNetSolutionResponse) *string { return v.DeleteOrUpdateBehavior }).(pulumi.StringPtrOutput)
}

func (o VNetSolutionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VNetSolutionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VNetSolutionResponsePtrOutput struct{ *pulumi.OutputState }

func (VNetSolutionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolutionResponse)(nil)).Elem()
}

func (o VNetSolutionResponsePtrOutput) ToVNetSolutionResponsePtrOutput() VNetSolutionResponsePtrOutput {
	return o
}

func (o VNetSolutionResponsePtrOutput) ToVNetSolutionResponsePtrOutputWithContext(ctx context.Context) VNetSolutionResponsePtrOutput {
	return o
}

func (o VNetSolutionResponsePtrOutput) Elem() VNetSolutionResponseOutput {
	return o.ApplyT(func(v *VNetSolutionResponse) VNetSolutionResponse {
		if v != nil {
			return *v
		}
		var ret VNetSolutionResponse
		return ret
	}).(VNetSolutionResponseOutput)
}

func (o VNetSolutionResponsePtrOutput) DeleteOrUpdateBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VNetSolutionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOrUpdateBehavior
	}).(pulumi.StringPtrOutput)
}

func (o VNetSolutionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VNetSolutionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ValueSecretInfo struct {
	SecretType string  `pulumi:"secretType"`
	Value      *string `pulumi:"value"`
}

type ValueSecretInfoResponse struct {
	SecretType string  `pulumi:"secretType"`
	Value      *string `pulumi:"value"`
}

func init() {
	pulumi.RegisterOutputType(ConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(CreateOrUpdateDryrunParametersOutput{})
	pulumi.RegisterOutputType(CreateOrUpdateDryrunParametersPtrOutput{})
	pulumi.RegisterOutputType(CreateOrUpdateDryrunParametersResponseOutput{})
	pulumi.RegisterOutputType(CreateOrUpdateDryrunParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(DryrunOperationPreviewResponseOutput{})
	pulumi.RegisterOutputType(DryrunOperationPreviewResponseArrayOutput{})
	pulumi.RegisterOutputType(FirewallRulesOutput{})
	pulumi.RegisterOutputType(FirewallRulesPtrOutput{})
	pulumi.RegisterOutputType(FirewallRulesResponseOutput{})
	pulumi.RegisterOutputType(FirewallRulesResponsePtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkSolutionOutput{})
	pulumi.RegisterOutputType(PublicNetworkSolutionPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkSolutionResponseOutput{})
	pulumi.RegisterOutputType(PublicNetworkSolutionResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretStoreOutput{})
	pulumi.RegisterOutputType(SecretStorePtrOutput{})
	pulumi.RegisterOutputType(SecretStoreResponseOutput{})
	pulumi.RegisterOutputType(SecretStoreResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SourceConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VNetSolutionOutput{})
	pulumi.RegisterOutputType(VNetSolutionPtrOutput{})
	pulumi.RegisterOutputType(VNetSolutionResponseOutput{})
	pulumi.RegisterOutputType(VNetSolutionResponsePtrOutput{})
}
