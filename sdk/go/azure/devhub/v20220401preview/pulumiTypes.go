


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ACR struct {
	AcrRegistryName   *string `pulumi:"acrRegistryName"`
	AcrRepositoryName *string `pulumi:"acrRepositoryName"`
	AcrResourceGroup  *string `pulumi:"acrResourceGroup"`
	AcrSubscriptionId *string `pulumi:"acrSubscriptionId"`
}





type ACRInput interface {
	pulumi.Input

	ToACROutput() ACROutput
	ToACROutputWithContext(context.Context) ACROutput
}

type ACRArgs struct {
	AcrRegistryName   pulumi.StringPtrInput `pulumi:"acrRegistryName"`
	AcrRepositoryName pulumi.StringPtrInput `pulumi:"acrRepositoryName"`
	AcrResourceGroup  pulumi.StringPtrInput `pulumi:"acrResourceGroup"`
	AcrSubscriptionId pulumi.StringPtrInput `pulumi:"acrSubscriptionId"`
}

func (ACRArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACR)(nil)).Elem()
}

func (i ACRArgs) ToACROutput() ACROutput {
	return i.ToACROutputWithContext(context.Background())
}

func (i ACRArgs) ToACROutputWithContext(ctx context.Context) ACROutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACROutput)
}

func (i ACRArgs) ToACRPtrOutput() ACRPtrOutput {
	return i.ToACRPtrOutputWithContext(context.Background())
}

func (i ACRArgs) ToACRPtrOutputWithContext(ctx context.Context) ACRPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACROutput).ToACRPtrOutputWithContext(ctx)
}









type ACRPtrInput interface {
	pulumi.Input

	ToACRPtrOutput() ACRPtrOutput
	ToACRPtrOutputWithContext(context.Context) ACRPtrOutput
}

type acrPtrType ACRArgs

func ACRPtr(v *ACRArgs) ACRPtrInput {
	return (*acrPtrType)(v)
}

func (*acrPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ACR)(nil)).Elem()
}

func (i *acrPtrType) ToACRPtrOutput() ACRPtrOutput {
	return i.ToACRPtrOutputWithContext(context.Background())
}

func (i *acrPtrType) ToACRPtrOutputWithContext(ctx context.Context) ACRPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACRPtrOutput)
}

type ACROutput struct{ *pulumi.OutputState }

func (ACROutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACR)(nil)).Elem()
}

func (o ACROutput) ToACROutput() ACROutput {
	return o
}

func (o ACROutput) ToACROutputWithContext(ctx context.Context) ACROutput {
	return o
}

func (o ACROutput) ToACRPtrOutput() ACRPtrOutput {
	return o.ToACRPtrOutputWithContext(context.Background())
}

func (o ACROutput) ToACRPtrOutputWithContext(ctx context.Context) ACRPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACR) *ACR {
		return &v
	}).(ACRPtrOutput)
}

func (o ACROutput) AcrRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACR) *string { return v.AcrRegistryName }).(pulumi.StringPtrOutput)
}

func (o ACROutput) AcrRepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACR) *string { return v.AcrRepositoryName }).(pulumi.StringPtrOutput)
}

func (o ACROutput) AcrResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACR) *string { return v.AcrResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ACROutput) AcrSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACR) *string { return v.AcrSubscriptionId }).(pulumi.StringPtrOutput)
}

type ACRPtrOutput struct{ *pulumi.OutputState }

func (ACRPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACR)(nil)).Elem()
}

func (o ACRPtrOutput) ToACRPtrOutput() ACRPtrOutput {
	return o
}

func (o ACRPtrOutput) ToACRPtrOutputWithContext(ctx context.Context) ACRPtrOutput {
	return o
}

func (o ACRPtrOutput) Elem() ACROutput {
	return o.ApplyT(func(v *ACR) ACR {
		if v != nil {
			return *v
		}
		var ret ACR
		return ret
	}).(ACROutput)
}

func (o ACRPtrOutput) AcrRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACR) *string {
		if v == nil {
			return nil
		}
		return v.AcrRegistryName
	}).(pulumi.StringPtrOutput)
}

func (o ACRPtrOutput) AcrRepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACR) *string {
		if v == nil {
			return nil
		}
		return v.AcrRepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o ACRPtrOutput) AcrResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACR) *string {
		if v == nil {
			return nil
		}
		return v.AcrResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o ACRPtrOutput) AcrSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACR) *string {
		if v == nil {
			return nil
		}
		return v.AcrSubscriptionId
	}).(pulumi.StringPtrOutput)
}

type ACRResponse struct {
	AcrRegistryName   *string `pulumi:"acrRegistryName"`
	AcrRepositoryName *string `pulumi:"acrRepositoryName"`
	AcrResourceGroup  *string `pulumi:"acrResourceGroup"`
	AcrSubscriptionId *string `pulumi:"acrSubscriptionId"`
}

type ACRResponseOutput struct{ *pulumi.OutputState }

func (ACRResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACRResponse)(nil)).Elem()
}

func (o ACRResponseOutput) ToACRResponseOutput() ACRResponseOutput {
	return o
}

func (o ACRResponseOutput) ToACRResponseOutputWithContext(ctx context.Context) ACRResponseOutput {
	return o
}

func (o ACRResponseOutput) AcrRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACRResponse) *string { return v.AcrRegistryName }).(pulumi.StringPtrOutput)
}

func (o ACRResponseOutput) AcrRepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACRResponse) *string { return v.AcrRepositoryName }).(pulumi.StringPtrOutput)
}

func (o ACRResponseOutput) AcrResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACRResponse) *string { return v.AcrResourceGroup }).(pulumi.StringPtrOutput)
}

func (o ACRResponseOutput) AcrSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACRResponse) *string { return v.AcrSubscriptionId }).(pulumi.StringPtrOutput)
}

type ACRResponsePtrOutput struct{ *pulumi.OutputState }

func (ACRResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACRResponse)(nil)).Elem()
}

func (o ACRResponsePtrOutput) ToACRResponsePtrOutput() ACRResponsePtrOutput {
	return o
}

func (o ACRResponsePtrOutput) ToACRResponsePtrOutputWithContext(ctx context.Context) ACRResponsePtrOutput {
	return o
}

func (o ACRResponsePtrOutput) Elem() ACRResponseOutput {
	return o.ApplyT(func(v *ACRResponse) ACRResponse {
		if v != nil {
			return *v
		}
		var ret ACRResponse
		return ret
	}).(ACRResponseOutput)
}

func (o ACRResponsePtrOutput) AcrRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACRResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcrRegistryName
	}).(pulumi.StringPtrOutput)
}

func (o ACRResponsePtrOutput) AcrRepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACRResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcrRepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o ACRResponsePtrOutput) AcrResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACRResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcrResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o ACRResponsePtrOutput) AcrSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ACRResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcrSubscriptionId
	}).(pulumi.StringPtrOutput)
}

type DeploymentProperties struct {
	HelmChartPath         *string           `pulumi:"helmChartPath"`
	HelmValues            *string           `pulumi:"helmValues"`
	KubeManifestLocations []string          `pulumi:"kubeManifestLocations"`
	ManifestType          *string           `pulumi:"manifestType"`
	Overrides             map[string]string `pulumi:"overrides"`
}





type DeploymentPropertiesInput interface {
	pulumi.Input

	ToDeploymentPropertiesOutput() DeploymentPropertiesOutput
	ToDeploymentPropertiesOutputWithContext(context.Context) DeploymentPropertiesOutput
}

type DeploymentPropertiesArgs struct {
	HelmChartPath         pulumi.StringPtrInput   `pulumi:"helmChartPath"`
	HelmValues            pulumi.StringPtrInput   `pulumi:"helmValues"`
	KubeManifestLocations pulumi.StringArrayInput `pulumi:"kubeManifestLocations"`
	ManifestType          pulumi.StringPtrInput   `pulumi:"manifestType"`
	Overrides             pulumi.StringMapInput   `pulumi:"overrides"`
}

func (DeploymentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentProperties)(nil)).Elem()
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesOutput() DeploymentPropertiesOutput {
	return i.ToDeploymentPropertiesOutputWithContext(context.Background())
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesOutputWithContext(ctx context.Context) DeploymentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesOutput)
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return i.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesOutput).ToDeploymentPropertiesPtrOutputWithContext(ctx)
}









type DeploymentPropertiesPtrInput interface {
	pulumi.Input

	ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput
	ToDeploymentPropertiesPtrOutputWithContext(context.Context) DeploymentPropertiesPtrOutput
}

type deploymentPropertiesPtrType DeploymentPropertiesArgs

func DeploymentPropertiesPtr(v *DeploymentPropertiesArgs) DeploymentPropertiesPtrInput {
	return (*deploymentPropertiesPtrType)(v)
}

func (*deploymentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProperties)(nil)).Elem()
}

func (i *deploymentPropertiesPtrType) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return i.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i *deploymentPropertiesPtrType) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesPtrOutput)
}

type DeploymentPropertiesOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentProperties)(nil)).Elem()
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesOutput() DeploymentPropertiesOutput {
	return o
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesOutputWithContext(ctx context.Context) DeploymentPropertiesOutput {
	return o
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return o.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentProperties) *DeploymentProperties {
		return &v
	}).(DeploymentPropertiesPtrOutput)
}

func (o DeploymentPropertiesOutput) HelmChartPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *string { return v.HelmChartPath }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesOutput) HelmValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *string { return v.HelmValues }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesOutput) KubeManifestLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DeploymentProperties) []string { return v.KubeManifestLocations }).(pulumi.StringArrayOutput)
}

func (o DeploymentPropertiesOutput) ManifestType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *string { return v.ManifestType }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesOutput) Overrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v DeploymentProperties) map[string]string { return v.Overrides }).(pulumi.StringMapOutput)
}

type DeploymentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProperties)(nil)).Elem()
}

func (o DeploymentPropertiesPtrOutput) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return o
}

func (o DeploymentPropertiesPtrOutput) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return o
}

func (o DeploymentPropertiesPtrOutput) Elem() DeploymentPropertiesOutput {
	return o.ApplyT(func(v *DeploymentProperties) DeploymentProperties {
		if v != nil {
			return *v
		}
		var ret DeploymentProperties
		return ret
	}).(DeploymentPropertiesOutput)
}

func (o DeploymentPropertiesPtrOutput) HelmChartPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *string {
		if v == nil {
			return nil
		}
		return v.HelmChartPath
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesPtrOutput) HelmValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *string {
		if v == nil {
			return nil
		}
		return v.HelmValues
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesPtrOutput) KubeManifestLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DeploymentProperties) []string {
		if v == nil {
			return nil
		}
		return v.KubeManifestLocations
	}).(pulumi.StringArrayOutput)
}

func (o DeploymentPropertiesPtrOutput) ManifestType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManifestType
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesPtrOutput) Overrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Overrides
	}).(pulumi.StringMapOutput)
}

type DeploymentPropertiesResponse struct {
	HelmChartPath         *string           `pulumi:"helmChartPath"`
	HelmValues            *string           `pulumi:"helmValues"`
	KubeManifestLocations []string          `pulumi:"kubeManifestLocations"`
	ManifestType          *string           `pulumi:"manifestType"`
	Overrides             map[string]string `pulumi:"overrides"`
}

type DeploymentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentPropertiesResponse)(nil)).Elem()
}

func (o DeploymentPropertiesResponseOutput) ToDeploymentPropertiesResponseOutput() DeploymentPropertiesResponseOutput {
	return o
}

func (o DeploymentPropertiesResponseOutput) ToDeploymentPropertiesResponseOutputWithContext(ctx context.Context) DeploymentPropertiesResponseOutput {
	return o
}

func (o DeploymentPropertiesResponseOutput) HelmChartPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesResponse) *string { return v.HelmChartPath }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesResponseOutput) HelmValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesResponse) *string { return v.HelmValues }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesResponseOutput) KubeManifestLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesResponse) []string { return v.KubeManifestLocations }).(pulumi.StringArrayOutput)
}

func (o DeploymentPropertiesResponseOutput) ManifestType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesResponse) *string { return v.ManifestType }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesResponseOutput) Overrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v DeploymentPropertiesResponse) map[string]string { return v.Overrides }).(pulumi.StringMapOutput)
}

type DeploymentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentPropertiesResponse)(nil)).Elem()
}

func (o DeploymentPropertiesResponsePtrOutput) ToDeploymentPropertiesResponsePtrOutput() DeploymentPropertiesResponsePtrOutput {
	return o
}

func (o DeploymentPropertiesResponsePtrOutput) ToDeploymentPropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentPropertiesResponsePtrOutput {
	return o
}

func (o DeploymentPropertiesResponsePtrOutput) Elem() DeploymentPropertiesResponseOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) DeploymentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentPropertiesResponse
		return ret
	}).(DeploymentPropertiesResponseOutput)
}

func (o DeploymentPropertiesResponsePtrOutput) HelmChartPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.HelmChartPath
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesResponsePtrOutput) HelmValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.HelmValues
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesResponsePtrOutput) KubeManifestLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.KubeManifestLocations
	}).(pulumi.StringArrayOutput)
}

func (o DeploymentPropertiesResponsePtrOutput) ManifestType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ManifestType
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesResponsePtrOutput) Overrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentPropertiesResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Overrides
	}).(pulumi.StringMapOutput)
}

type GitHubWorkflowProfileOidcCredentials struct {
	AzureClientId *string `pulumi:"azureClientId"`
	AzureTenantId *string `pulumi:"azureTenantId"`
}





type GitHubWorkflowProfileOidcCredentialsInput interface {
	pulumi.Input

	ToGitHubWorkflowProfileOidcCredentialsOutput() GitHubWorkflowProfileOidcCredentialsOutput
	ToGitHubWorkflowProfileOidcCredentialsOutputWithContext(context.Context) GitHubWorkflowProfileOidcCredentialsOutput
}

type GitHubWorkflowProfileOidcCredentialsArgs struct {
	AzureClientId pulumi.StringPtrInput `pulumi:"azureClientId"`
	AzureTenantId pulumi.StringPtrInput `pulumi:"azureTenantId"`
}

func (GitHubWorkflowProfileOidcCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubWorkflowProfileOidcCredentials)(nil)).Elem()
}

func (i GitHubWorkflowProfileOidcCredentialsArgs) ToGitHubWorkflowProfileOidcCredentialsOutput() GitHubWorkflowProfileOidcCredentialsOutput {
	return i.ToGitHubWorkflowProfileOidcCredentialsOutputWithContext(context.Background())
}

func (i GitHubWorkflowProfileOidcCredentialsArgs) ToGitHubWorkflowProfileOidcCredentialsOutputWithContext(ctx context.Context) GitHubWorkflowProfileOidcCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubWorkflowProfileOidcCredentialsOutput)
}

func (i GitHubWorkflowProfileOidcCredentialsArgs) ToGitHubWorkflowProfileOidcCredentialsPtrOutput() GitHubWorkflowProfileOidcCredentialsPtrOutput {
	return i.ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(context.Background())
}

func (i GitHubWorkflowProfileOidcCredentialsArgs) ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(ctx context.Context) GitHubWorkflowProfileOidcCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubWorkflowProfileOidcCredentialsOutput).ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(ctx)
}









type GitHubWorkflowProfileOidcCredentialsPtrInput interface {
	pulumi.Input

	ToGitHubWorkflowProfileOidcCredentialsPtrOutput() GitHubWorkflowProfileOidcCredentialsPtrOutput
	ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(context.Context) GitHubWorkflowProfileOidcCredentialsPtrOutput
}

type gitHubWorkflowProfileOidcCredentialsPtrType GitHubWorkflowProfileOidcCredentialsArgs

func GitHubWorkflowProfileOidcCredentialsPtr(v *GitHubWorkflowProfileOidcCredentialsArgs) GitHubWorkflowProfileOidcCredentialsPtrInput {
	return (*gitHubWorkflowProfileOidcCredentialsPtrType)(v)
}

func (*gitHubWorkflowProfileOidcCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubWorkflowProfileOidcCredentials)(nil)).Elem()
}

func (i *gitHubWorkflowProfileOidcCredentialsPtrType) ToGitHubWorkflowProfileOidcCredentialsPtrOutput() GitHubWorkflowProfileOidcCredentialsPtrOutput {
	return i.ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(context.Background())
}

func (i *gitHubWorkflowProfileOidcCredentialsPtrType) ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(ctx context.Context) GitHubWorkflowProfileOidcCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubWorkflowProfileOidcCredentialsPtrOutput)
}

type GitHubWorkflowProfileOidcCredentialsOutput struct{ *pulumi.OutputState }

func (GitHubWorkflowProfileOidcCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubWorkflowProfileOidcCredentials)(nil)).Elem()
}

func (o GitHubWorkflowProfileOidcCredentialsOutput) ToGitHubWorkflowProfileOidcCredentialsOutput() GitHubWorkflowProfileOidcCredentialsOutput {
	return o
}

func (o GitHubWorkflowProfileOidcCredentialsOutput) ToGitHubWorkflowProfileOidcCredentialsOutputWithContext(ctx context.Context) GitHubWorkflowProfileOidcCredentialsOutput {
	return o
}

func (o GitHubWorkflowProfileOidcCredentialsOutput) ToGitHubWorkflowProfileOidcCredentialsPtrOutput() GitHubWorkflowProfileOidcCredentialsPtrOutput {
	return o.ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(context.Background())
}

func (o GitHubWorkflowProfileOidcCredentialsOutput) ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(ctx context.Context) GitHubWorkflowProfileOidcCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubWorkflowProfileOidcCredentials) *GitHubWorkflowProfileOidcCredentials {
		return &v
	}).(GitHubWorkflowProfileOidcCredentialsPtrOutput)
}

func (o GitHubWorkflowProfileOidcCredentialsOutput) AzureClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubWorkflowProfileOidcCredentials) *string { return v.AzureClientId }).(pulumi.StringPtrOutput)
}

func (o GitHubWorkflowProfileOidcCredentialsOutput) AzureTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubWorkflowProfileOidcCredentials) *string { return v.AzureTenantId }).(pulumi.StringPtrOutput)
}

type GitHubWorkflowProfileOidcCredentialsPtrOutput struct{ *pulumi.OutputState }

func (GitHubWorkflowProfileOidcCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubWorkflowProfileOidcCredentials)(nil)).Elem()
}

func (o GitHubWorkflowProfileOidcCredentialsPtrOutput) ToGitHubWorkflowProfileOidcCredentialsPtrOutput() GitHubWorkflowProfileOidcCredentialsPtrOutput {
	return o
}

func (o GitHubWorkflowProfileOidcCredentialsPtrOutput) ToGitHubWorkflowProfileOidcCredentialsPtrOutputWithContext(ctx context.Context) GitHubWorkflowProfileOidcCredentialsPtrOutput {
	return o
}

func (o GitHubWorkflowProfileOidcCredentialsPtrOutput) Elem() GitHubWorkflowProfileOidcCredentialsOutput {
	return o.ApplyT(func(v *GitHubWorkflowProfileOidcCredentials) GitHubWorkflowProfileOidcCredentials {
		if v != nil {
			return *v
		}
		var ret GitHubWorkflowProfileOidcCredentials
		return ret
	}).(GitHubWorkflowProfileOidcCredentialsOutput)
}

func (o GitHubWorkflowProfileOidcCredentialsPtrOutput) AzureClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubWorkflowProfileOidcCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AzureClientId
	}).(pulumi.StringPtrOutput)
}

func (o GitHubWorkflowProfileOidcCredentialsPtrOutput) AzureTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubWorkflowProfileOidcCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AzureTenantId
	}).(pulumi.StringPtrOutput)
}

type GitHubWorkflowProfileResponseOidcCredentials struct {
	AzureClientId *string `pulumi:"azureClientId"`
	AzureTenantId *string `pulumi:"azureTenantId"`
}

type GitHubWorkflowProfileResponseOidcCredentialsOutput struct{ *pulumi.OutputState }

func (GitHubWorkflowProfileResponseOidcCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubWorkflowProfileResponseOidcCredentials)(nil)).Elem()
}

func (o GitHubWorkflowProfileResponseOidcCredentialsOutput) ToGitHubWorkflowProfileResponseOidcCredentialsOutput() GitHubWorkflowProfileResponseOidcCredentialsOutput {
	return o
}

func (o GitHubWorkflowProfileResponseOidcCredentialsOutput) ToGitHubWorkflowProfileResponseOidcCredentialsOutputWithContext(ctx context.Context) GitHubWorkflowProfileResponseOidcCredentialsOutput {
	return o
}

func (o GitHubWorkflowProfileResponseOidcCredentialsOutput) AzureClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubWorkflowProfileResponseOidcCredentials) *string { return v.AzureClientId }).(pulumi.StringPtrOutput)
}

func (o GitHubWorkflowProfileResponseOidcCredentialsOutput) AzureTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubWorkflowProfileResponseOidcCredentials) *string { return v.AzureTenantId }).(pulumi.StringPtrOutput)
}

type GitHubWorkflowProfileResponseOidcCredentialsPtrOutput struct{ *pulumi.OutputState }

func (GitHubWorkflowProfileResponseOidcCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubWorkflowProfileResponseOidcCredentials)(nil)).Elem()
}

func (o GitHubWorkflowProfileResponseOidcCredentialsPtrOutput) ToGitHubWorkflowProfileResponseOidcCredentialsPtrOutput() GitHubWorkflowProfileResponseOidcCredentialsPtrOutput {
	return o
}

func (o GitHubWorkflowProfileResponseOidcCredentialsPtrOutput) ToGitHubWorkflowProfileResponseOidcCredentialsPtrOutputWithContext(ctx context.Context) GitHubWorkflowProfileResponseOidcCredentialsPtrOutput {
	return o
}

func (o GitHubWorkflowProfileResponseOidcCredentialsPtrOutput) Elem() GitHubWorkflowProfileResponseOidcCredentialsOutput {
	return o.ApplyT(func(v *GitHubWorkflowProfileResponseOidcCredentials) GitHubWorkflowProfileResponseOidcCredentials {
		if v != nil {
			return *v
		}
		var ret GitHubWorkflowProfileResponseOidcCredentials
		return ret
	}).(GitHubWorkflowProfileResponseOidcCredentialsOutput)
}

func (o GitHubWorkflowProfileResponseOidcCredentialsPtrOutput) AzureClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubWorkflowProfileResponseOidcCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AzureClientId
	}).(pulumi.StringPtrOutput)
}

func (o GitHubWorkflowProfileResponseOidcCredentialsPtrOutput) AzureTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubWorkflowProfileResponseOidcCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AzureTenantId
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

type WorkflowRunResponse struct {
	LastRunAt      string `pulumi:"lastRunAt"`
	Succeeded      bool   `pulumi:"succeeded"`
	WorkflowRunURL string `pulumi:"workflowRunURL"`
}

type WorkflowRunResponseOutput struct{ *pulumi.OutputState }

func (WorkflowRunResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowRunResponse)(nil)).Elem()
}

func (o WorkflowRunResponseOutput) ToWorkflowRunResponseOutput() WorkflowRunResponseOutput {
	return o
}

func (o WorkflowRunResponseOutput) ToWorkflowRunResponseOutputWithContext(ctx context.Context) WorkflowRunResponseOutput {
	return o
}

func (o WorkflowRunResponseOutput) LastRunAt() pulumi.StringOutput {
	return o.ApplyT(func(v WorkflowRunResponse) string { return v.LastRunAt }).(pulumi.StringOutput)
}

func (o WorkflowRunResponseOutput) Succeeded() pulumi.BoolOutput {
	return o.ApplyT(func(v WorkflowRunResponse) bool { return v.Succeeded }).(pulumi.BoolOutput)
}

func (o WorkflowRunResponseOutput) WorkflowRunURL() pulumi.StringOutput {
	return o.ApplyT(func(v WorkflowRunResponse) string { return v.WorkflowRunURL }).(pulumi.StringOutput)
}

type WorkflowRunResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkflowRunResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowRunResponse)(nil)).Elem()
}

func (o WorkflowRunResponsePtrOutput) ToWorkflowRunResponsePtrOutput() WorkflowRunResponsePtrOutput {
	return o
}

func (o WorkflowRunResponsePtrOutput) ToWorkflowRunResponsePtrOutputWithContext(ctx context.Context) WorkflowRunResponsePtrOutput {
	return o
}

func (o WorkflowRunResponsePtrOutput) Elem() WorkflowRunResponseOutput {
	return o.ApplyT(func(v *WorkflowRunResponse) WorkflowRunResponse {
		if v != nil {
			return *v
		}
		var ret WorkflowRunResponse
		return ret
	}).(WorkflowRunResponseOutput)
}

func (o WorkflowRunResponsePtrOutput) LastRunAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowRunResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastRunAt
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowRunResponsePtrOutput) Succeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkflowRunResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Succeeded
	}).(pulumi.BoolPtrOutput)
}

func (o WorkflowRunResponsePtrOutput) WorkflowRunURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowRunResponse) *string {
		if v == nil {
			return nil
		}
		return &v.WorkflowRunURL
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ACROutput{})
	pulumi.RegisterOutputType(ACRPtrOutput{})
	pulumi.RegisterOutputType(ACRResponseOutput{})
	pulumi.RegisterOutputType(ACRResponsePtrOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubWorkflowProfileOidcCredentialsOutput{})
	pulumi.RegisterOutputType(GitHubWorkflowProfileOidcCredentialsPtrOutput{})
	pulumi.RegisterOutputType(GitHubWorkflowProfileResponseOidcCredentialsOutput{})
	pulumi.RegisterOutputType(GitHubWorkflowProfileResponseOidcCredentialsPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(WorkflowRunResponseOutput{})
	pulumi.RegisterOutputType(WorkflowRunResponsePtrOutput{})
}
