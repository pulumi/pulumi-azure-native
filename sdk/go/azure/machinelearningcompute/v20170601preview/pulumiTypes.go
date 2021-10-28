


package v20170601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AcsClusterProperties struct {
	AgentCount             *int                        `pulumi:"agentCount"`
	AgentVmSize            *string                     `pulumi:"agentVmSize"`
	OrchestratorProperties KubernetesClusterProperties `pulumi:"orchestratorProperties"`
	OrchestratorType       string                      `pulumi:"orchestratorType"`
	SystemServices         []string                    `pulumi:"systemServices"`
}





type AcsClusterPropertiesInput interface {
	pulumi.Input

	ToAcsClusterPropertiesOutput() AcsClusterPropertiesOutput
	ToAcsClusterPropertiesOutputWithContext(context.Context) AcsClusterPropertiesOutput
}

type AcsClusterPropertiesArgs struct {
	AgentCount             pulumi.IntPtrInput               `pulumi:"agentCount"`
	AgentVmSize            pulumi.StringPtrInput            `pulumi:"agentVmSize"`
	OrchestratorProperties KubernetesClusterPropertiesInput `pulumi:"orchestratorProperties"`
	OrchestratorType       pulumi.StringInput               `pulumi:"orchestratorType"`
	SystemServices         pulumi.StringArrayInput          `pulumi:"systemServices"`
}

func (AcsClusterPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AcsClusterProperties)(nil)).Elem()
}

func (i AcsClusterPropertiesArgs) ToAcsClusterPropertiesOutput() AcsClusterPropertiesOutput {
	return i.ToAcsClusterPropertiesOutputWithContext(context.Background())
}

func (i AcsClusterPropertiesArgs) ToAcsClusterPropertiesOutputWithContext(ctx context.Context) AcsClusterPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcsClusterPropertiesOutput)
}

func (i AcsClusterPropertiesArgs) ToAcsClusterPropertiesPtrOutput() AcsClusterPropertiesPtrOutput {
	return i.ToAcsClusterPropertiesPtrOutputWithContext(context.Background())
}

func (i AcsClusterPropertiesArgs) ToAcsClusterPropertiesPtrOutputWithContext(ctx context.Context) AcsClusterPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcsClusterPropertiesOutput).ToAcsClusterPropertiesPtrOutputWithContext(ctx)
}









type AcsClusterPropertiesPtrInput interface {
	pulumi.Input

	ToAcsClusterPropertiesPtrOutput() AcsClusterPropertiesPtrOutput
	ToAcsClusterPropertiesPtrOutputWithContext(context.Context) AcsClusterPropertiesPtrOutput
}

type acsClusterPropertiesPtrType AcsClusterPropertiesArgs

func AcsClusterPropertiesPtr(v *AcsClusterPropertiesArgs) AcsClusterPropertiesPtrInput {
	return (*acsClusterPropertiesPtrType)(v)
}

func (*acsClusterPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AcsClusterProperties)(nil)).Elem()
}

func (i *acsClusterPropertiesPtrType) ToAcsClusterPropertiesPtrOutput() AcsClusterPropertiesPtrOutput {
	return i.ToAcsClusterPropertiesPtrOutputWithContext(context.Background())
}

func (i *acsClusterPropertiesPtrType) ToAcsClusterPropertiesPtrOutputWithContext(ctx context.Context) AcsClusterPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcsClusterPropertiesPtrOutput)
}

type AcsClusterPropertiesOutput struct{ *pulumi.OutputState }

func (AcsClusterPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcsClusterProperties)(nil)).Elem()
}

func (o AcsClusterPropertiesOutput) ToAcsClusterPropertiesOutput() AcsClusterPropertiesOutput {
	return o
}

func (o AcsClusterPropertiesOutput) ToAcsClusterPropertiesOutputWithContext(ctx context.Context) AcsClusterPropertiesOutput {
	return o
}

func (o AcsClusterPropertiesOutput) ToAcsClusterPropertiesPtrOutput() AcsClusterPropertiesPtrOutput {
	return o.ToAcsClusterPropertiesPtrOutputWithContext(context.Background())
}

func (o AcsClusterPropertiesOutput) ToAcsClusterPropertiesPtrOutputWithContext(ctx context.Context) AcsClusterPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AcsClusterProperties) *AcsClusterProperties {
		return &v
	}).(AcsClusterPropertiesPtrOutput)
}

func (o AcsClusterPropertiesOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AcsClusterProperties) *int { return v.AgentCount }).(pulumi.IntPtrOutput)
}

func (o AcsClusterPropertiesOutput) AgentVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AcsClusterProperties) *string { return v.AgentVmSize }).(pulumi.StringPtrOutput)
}

func (o AcsClusterPropertiesOutput) OrchestratorProperties() KubernetesClusterPropertiesOutput {
	return o.ApplyT(func(v AcsClusterProperties) KubernetesClusterProperties { return v.OrchestratorProperties }).(KubernetesClusterPropertiesOutput)
}

func (o AcsClusterPropertiesOutput) OrchestratorType() pulumi.StringOutput {
	return o.ApplyT(func(v AcsClusterProperties) string { return v.OrchestratorType }).(pulumi.StringOutput)
}

func (o AcsClusterPropertiesOutput) SystemServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AcsClusterProperties) []string { return v.SystemServices }).(pulumi.StringArrayOutput)
}

type AcsClusterPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AcsClusterPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AcsClusterProperties)(nil)).Elem()
}

func (o AcsClusterPropertiesPtrOutput) ToAcsClusterPropertiesPtrOutput() AcsClusterPropertiesPtrOutput {
	return o
}

func (o AcsClusterPropertiesPtrOutput) ToAcsClusterPropertiesPtrOutputWithContext(ctx context.Context) AcsClusterPropertiesPtrOutput {
	return o
}

func (o AcsClusterPropertiesPtrOutput) Elem() AcsClusterPropertiesOutput {
	return o.ApplyT(func(v *AcsClusterProperties) AcsClusterProperties {
		if v != nil {
			return *v
		}
		var ret AcsClusterProperties
		return ret
	}).(AcsClusterPropertiesOutput)
}

func (o AcsClusterPropertiesPtrOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AcsClusterProperties) *int {
		if v == nil {
			return nil
		}
		return v.AgentCount
	}).(pulumi.IntPtrOutput)
}

func (o AcsClusterPropertiesPtrOutput) AgentVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AcsClusterProperties) *string {
		if v == nil {
			return nil
		}
		return v.AgentVmSize
	}).(pulumi.StringPtrOutput)
}

func (o AcsClusterPropertiesPtrOutput) OrchestratorProperties() KubernetesClusterPropertiesPtrOutput {
	return o.ApplyT(func(v *AcsClusterProperties) *KubernetesClusterProperties {
		if v == nil {
			return nil
		}
		return &v.OrchestratorProperties
	}).(KubernetesClusterPropertiesPtrOutput)
}

func (o AcsClusterPropertiesPtrOutput) OrchestratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AcsClusterProperties) *string {
		if v == nil {
			return nil
		}
		return &v.OrchestratorType
	}).(pulumi.StringPtrOutput)
}

func (o AcsClusterPropertiesPtrOutput) SystemServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AcsClusterProperties) []string {
		if v == nil {
			return nil
		}
		return v.SystemServices
	}).(pulumi.StringArrayOutput)
}

type AcsClusterPropertiesResponse struct {
	AgentCount             *int                                `pulumi:"agentCount"`
	AgentVmSize            *string                             `pulumi:"agentVmSize"`
	ClusterFqdn            string                              `pulumi:"clusterFqdn"`
	OrchestratorProperties KubernetesClusterPropertiesResponse `pulumi:"orchestratorProperties"`
	OrchestratorType       string                              `pulumi:"orchestratorType"`
	SystemServices         []string                            `pulumi:"systemServices"`
}





type AcsClusterPropertiesResponseInput interface {
	pulumi.Input

	ToAcsClusterPropertiesResponseOutput() AcsClusterPropertiesResponseOutput
	ToAcsClusterPropertiesResponseOutputWithContext(context.Context) AcsClusterPropertiesResponseOutput
}

type AcsClusterPropertiesResponseArgs struct {
	AgentCount             pulumi.IntPtrInput                       `pulumi:"agentCount"`
	AgentVmSize            pulumi.StringPtrInput                    `pulumi:"agentVmSize"`
	ClusterFqdn            pulumi.StringInput                       `pulumi:"clusterFqdn"`
	OrchestratorProperties KubernetesClusterPropertiesResponseInput `pulumi:"orchestratorProperties"`
	OrchestratorType       pulumi.StringInput                       `pulumi:"orchestratorType"`
	SystemServices         pulumi.StringArrayInput                  `pulumi:"systemServices"`
}

func (AcsClusterPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AcsClusterPropertiesResponse)(nil)).Elem()
}

func (i AcsClusterPropertiesResponseArgs) ToAcsClusterPropertiesResponseOutput() AcsClusterPropertiesResponseOutput {
	return i.ToAcsClusterPropertiesResponseOutputWithContext(context.Background())
}

func (i AcsClusterPropertiesResponseArgs) ToAcsClusterPropertiesResponseOutputWithContext(ctx context.Context) AcsClusterPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcsClusterPropertiesResponseOutput)
}

func (i AcsClusterPropertiesResponseArgs) ToAcsClusterPropertiesResponsePtrOutput() AcsClusterPropertiesResponsePtrOutput {
	return i.ToAcsClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AcsClusterPropertiesResponseArgs) ToAcsClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) AcsClusterPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcsClusterPropertiesResponseOutput).ToAcsClusterPropertiesResponsePtrOutputWithContext(ctx)
}









type AcsClusterPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAcsClusterPropertiesResponsePtrOutput() AcsClusterPropertiesResponsePtrOutput
	ToAcsClusterPropertiesResponsePtrOutputWithContext(context.Context) AcsClusterPropertiesResponsePtrOutput
}

type acsClusterPropertiesResponsePtrType AcsClusterPropertiesResponseArgs

func AcsClusterPropertiesResponsePtr(v *AcsClusterPropertiesResponseArgs) AcsClusterPropertiesResponsePtrInput {
	return (*acsClusterPropertiesResponsePtrType)(v)
}

func (*acsClusterPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AcsClusterPropertiesResponse)(nil)).Elem()
}

func (i *acsClusterPropertiesResponsePtrType) ToAcsClusterPropertiesResponsePtrOutput() AcsClusterPropertiesResponsePtrOutput {
	return i.ToAcsClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *acsClusterPropertiesResponsePtrType) ToAcsClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) AcsClusterPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcsClusterPropertiesResponsePtrOutput)
}

type AcsClusterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AcsClusterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcsClusterPropertiesResponse)(nil)).Elem()
}

func (o AcsClusterPropertiesResponseOutput) ToAcsClusterPropertiesResponseOutput() AcsClusterPropertiesResponseOutput {
	return o
}

func (o AcsClusterPropertiesResponseOutput) ToAcsClusterPropertiesResponseOutputWithContext(ctx context.Context) AcsClusterPropertiesResponseOutput {
	return o
}

func (o AcsClusterPropertiesResponseOutput) ToAcsClusterPropertiesResponsePtrOutput() AcsClusterPropertiesResponsePtrOutput {
	return o.ToAcsClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AcsClusterPropertiesResponseOutput) ToAcsClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) AcsClusterPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AcsClusterPropertiesResponse) *AcsClusterPropertiesResponse {
		return &v
	}).(AcsClusterPropertiesResponsePtrOutput)
}

func (o AcsClusterPropertiesResponseOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AcsClusterPropertiesResponse) *int { return v.AgentCount }).(pulumi.IntPtrOutput)
}

func (o AcsClusterPropertiesResponseOutput) AgentVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AcsClusterPropertiesResponse) *string { return v.AgentVmSize }).(pulumi.StringPtrOutput)
}

func (o AcsClusterPropertiesResponseOutput) ClusterFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v AcsClusterPropertiesResponse) string { return v.ClusterFqdn }).(pulumi.StringOutput)
}

func (o AcsClusterPropertiesResponseOutput) OrchestratorProperties() KubernetesClusterPropertiesResponseOutput {
	return o.ApplyT(func(v AcsClusterPropertiesResponse) KubernetesClusterPropertiesResponse {
		return v.OrchestratorProperties
	}).(KubernetesClusterPropertiesResponseOutput)
}

func (o AcsClusterPropertiesResponseOutput) OrchestratorType() pulumi.StringOutput {
	return o.ApplyT(func(v AcsClusterPropertiesResponse) string { return v.OrchestratorType }).(pulumi.StringOutput)
}

func (o AcsClusterPropertiesResponseOutput) SystemServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AcsClusterPropertiesResponse) []string { return v.SystemServices }).(pulumi.StringArrayOutput)
}

type AcsClusterPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AcsClusterPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AcsClusterPropertiesResponse)(nil)).Elem()
}

func (o AcsClusterPropertiesResponsePtrOutput) ToAcsClusterPropertiesResponsePtrOutput() AcsClusterPropertiesResponsePtrOutput {
	return o
}

func (o AcsClusterPropertiesResponsePtrOutput) ToAcsClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) AcsClusterPropertiesResponsePtrOutput {
	return o
}

func (o AcsClusterPropertiesResponsePtrOutput) Elem() AcsClusterPropertiesResponseOutput {
	return o.ApplyT(func(v *AcsClusterPropertiesResponse) AcsClusterPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AcsClusterPropertiesResponse
		return ret
	}).(AcsClusterPropertiesResponseOutput)
}

func (o AcsClusterPropertiesResponsePtrOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AcsClusterPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.AgentCount
	}).(pulumi.IntPtrOutput)
}

func (o AcsClusterPropertiesResponsePtrOutput) AgentVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AcsClusterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentVmSize
	}).(pulumi.StringPtrOutput)
}

func (o AcsClusterPropertiesResponsePtrOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AcsClusterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClusterFqdn
	}).(pulumi.StringPtrOutput)
}

func (o AcsClusterPropertiesResponsePtrOutput) OrchestratorProperties() KubernetesClusterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AcsClusterPropertiesResponse) *KubernetesClusterPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.OrchestratorProperties
	}).(KubernetesClusterPropertiesResponsePtrOutput)
}

func (o AcsClusterPropertiesResponsePtrOutput) OrchestratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AcsClusterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OrchestratorType
	}).(pulumi.StringPtrOutput)
}

func (o AcsClusterPropertiesResponsePtrOutput) SystemServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AcsClusterPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.SystemServices
	}).(pulumi.StringArrayOutput)
}

type AppInsightsCredentials struct {
	ApiKey *string `pulumi:"apiKey"`
	AppId  *string `pulumi:"appId"`
}





type AppInsightsCredentialsInput interface {
	pulumi.Input

	ToAppInsightsCredentialsOutput() AppInsightsCredentialsOutput
	ToAppInsightsCredentialsOutputWithContext(context.Context) AppInsightsCredentialsOutput
}

type AppInsightsCredentialsArgs struct {
	ApiKey pulumi.StringPtrInput `pulumi:"apiKey"`
	AppId  pulumi.StringPtrInput `pulumi:"appId"`
}

func (AppInsightsCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppInsightsCredentials)(nil)).Elem()
}

func (i AppInsightsCredentialsArgs) ToAppInsightsCredentialsOutput() AppInsightsCredentialsOutput {
	return i.ToAppInsightsCredentialsOutputWithContext(context.Background())
}

func (i AppInsightsCredentialsArgs) ToAppInsightsCredentialsOutputWithContext(ctx context.Context) AppInsightsCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInsightsCredentialsOutput)
}

func (i AppInsightsCredentialsArgs) ToAppInsightsCredentialsPtrOutput() AppInsightsCredentialsPtrOutput {
	return i.ToAppInsightsCredentialsPtrOutputWithContext(context.Background())
}

func (i AppInsightsCredentialsArgs) ToAppInsightsCredentialsPtrOutputWithContext(ctx context.Context) AppInsightsCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInsightsCredentialsOutput).ToAppInsightsCredentialsPtrOutputWithContext(ctx)
}









type AppInsightsCredentialsPtrInput interface {
	pulumi.Input

	ToAppInsightsCredentialsPtrOutput() AppInsightsCredentialsPtrOutput
	ToAppInsightsCredentialsPtrOutputWithContext(context.Context) AppInsightsCredentialsPtrOutput
}

type appInsightsCredentialsPtrType AppInsightsCredentialsArgs

func AppInsightsCredentialsPtr(v *AppInsightsCredentialsArgs) AppInsightsCredentialsPtrInput {
	return (*appInsightsCredentialsPtrType)(v)
}

func (*appInsightsCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppInsightsCredentials)(nil)).Elem()
}

func (i *appInsightsCredentialsPtrType) ToAppInsightsCredentialsPtrOutput() AppInsightsCredentialsPtrOutput {
	return i.ToAppInsightsCredentialsPtrOutputWithContext(context.Background())
}

func (i *appInsightsCredentialsPtrType) ToAppInsightsCredentialsPtrOutputWithContext(ctx context.Context) AppInsightsCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInsightsCredentialsPtrOutput)
}

type AppInsightsCredentialsOutput struct{ *pulumi.OutputState }

func (AppInsightsCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppInsightsCredentials)(nil)).Elem()
}

func (o AppInsightsCredentialsOutput) ToAppInsightsCredentialsOutput() AppInsightsCredentialsOutput {
	return o
}

func (o AppInsightsCredentialsOutput) ToAppInsightsCredentialsOutputWithContext(ctx context.Context) AppInsightsCredentialsOutput {
	return o
}

func (o AppInsightsCredentialsOutput) ToAppInsightsCredentialsPtrOutput() AppInsightsCredentialsPtrOutput {
	return o.ToAppInsightsCredentialsPtrOutputWithContext(context.Background())
}

func (o AppInsightsCredentialsOutput) ToAppInsightsCredentialsPtrOutputWithContext(ctx context.Context) AppInsightsCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppInsightsCredentials) *AppInsightsCredentials {
		return &v
	}).(AppInsightsCredentialsPtrOutput)
}

func (o AppInsightsCredentialsOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppInsightsCredentials) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o AppInsightsCredentialsOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppInsightsCredentials) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

type AppInsightsCredentialsPtrOutput struct{ *pulumi.OutputState }

func (AppInsightsCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppInsightsCredentials)(nil)).Elem()
}

func (o AppInsightsCredentialsPtrOutput) ToAppInsightsCredentialsPtrOutput() AppInsightsCredentialsPtrOutput {
	return o
}

func (o AppInsightsCredentialsPtrOutput) ToAppInsightsCredentialsPtrOutputWithContext(ctx context.Context) AppInsightsCredentialsPtrOutput {
	return o
}

func (o AppInsightsCredentialsPtrOutput) Elem() AppInsightsCredentialsOutput {
	return o.ApplyT(func(v *AppInsightsCredentials) AppInsightsCredentials {
		if v != nil {
			return *v
		}
		var ret AppInsightsCredentials
		return ret
	}).(AppInsightsCredentialsOutput)
}

func (o AppInsightsCredentialsPtrOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppInsightsCredentials) *string {
		if v == nil {
			return nil
		}
		return v.ApiKey
	}).(pulumi.StringPtrOutput)
}

func (o AppInsightsCredentialsPtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppInsightsCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

type AppInsightsCredentialsResponse struct {
	ApiKey *string `pulumi:"apiKey"`
	AppId  *string `pulumi:"appId"`
}





type AppInsightsCredentialsResponseInput interface {
	pulumi.Input

	ToAppInsightsCredentialsResponseOutput() AppInsightsCredentialsResponseOutput
	ToAppInsightsCredentialsResponseOutputWithContext(context.Context) AppInsightsCredentialsResponseOutput
}

type AppInsightsCredentialsResponseArgs struct {
	ApiKey pulumi.StringPtrInput `pulumi:"apiKey"`
	AppId  pulumi.StringPtrInput `pulumi:"appId"`
}

func (AppInsightsCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppInsightsCredentialsResponse)(nil)).Elem()
}

func (i AppInsightsCredentialsResponseArgs) ToAppInsightsCredentialsResponseOutput() AppInsightsCredentialsResponseOutput {
	return i.ToAppInsightsCredentialsResponseOutputWithContext(context.Background())
}

func (i AppInsightsCredentialsResponseArgs) ToAppInsightsCredentialsResponseOutputWithContext(ctx context.Context) AppInsightsCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInsightsCredentialsResponseOutput)
}

func (i AppInsightsCredentialsResponseArgs) ToAppInsightsCredentialsResponsePtrOutput() AppInsightsCredentialsResponsePtrOutput {
	return i.ToAppInsightsCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i AppInsightsCredentialsResponseArgs) ToAppInsightsCredentialsResponsePtrOutputWithContext(ctx context.Context) AppInsightsCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInsightsCredentialsResponseOutput).ToAppInsightsCredentialsResponsePtrOutputWithContext(ctx)
}









type AppInsightsCredentialsResponsePtrInput interface {
	pulumi.Input

	ToAppInsightsCredentialsResponsePtrOutput() AppInsightsCredentialsResponsePtrOutput
	ToAppInsightsCredentialsResponsePtrOutputWithContext(context.Context) AppInsightsCredentialsResponsePtrOutput
}

type appInsightsCredentialsResponsePtrType AppInsightsCredentialsResponseArgs

func AppInsightsCredentialsResponsePtr(v *AppInsightsCredentialsResponseArgs) AppInsightsCredentialsResponsePtrInput {
	return (*appInsightsCredentialsResponsePtrType)(v)
}

func (*appInsightsCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppInsightsCredentialsResponse)(nil)).Elem()
}

func (i *appInsightsCredentialsResponsePtrType) ToAppInsightsCredentialsResponsePtrOutput() AppInsightsCredentialsResponsePtrOutput {
	return i.ToAppInsightsCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *appInsightsCredentialsResponsePtrType) ToAppInsightsCredentialsResponsePtrOutputWithContext(ctx context.Context) AppInsightsCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppInsightsCredentialsResponsePtrOutput)
}

type AppInsightsCredentialsResponseOutput struct{ *pulumi.OutputState }

func (AppInsightsCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppInsightsCredentialsResponse)(nil)).Elem()
}

func (o AppInsightsCredentialsResponseOutput) ToAppInsightsCredentialsResponseOutput() AppInsightsCredentialsResponseOutput {
	return o
}

func (o AppInsightsCredentialsResponseOutput) ToAppInsightsCredentialsResponseOutputWithContext(ctx context.Context) AppInsightsCredentialsResponseOutput {
	return o
}

func (o AppInsightsCredentialsResponseOutput) ToAppInsightsCredentialsResponsePtrOutput() AppInsightsCredentialsResponsePtrOutput {
	return o.ToAppInsightsCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o AppInsightsCredentialsResponseOutput) ToAppInsightsCredentialsResponsePtrOutputWithContext(ctx context.Context) AppInsightsCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppInsightsCredentialsResponse) *AppInsightsCredentialsResponse {
		return &v
	}).(AppInsightsCredentialsResponsePtrOutput)
}

func (o AppInsightsCredentialsResponseOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppInsightsCredentialsResponse) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o AppInsightsCredentialsResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppInsightsCredentialsResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

type AppInsightsCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (AppInsightsCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppInsightsCredentialsResponse)(nil)).Elem()
}

func (o AppInsightsCredentialsResponsePtrOutput) ToAppInsightsCredentialsResponsePtrOutput() AppInsightsCredentialsResponsePtrOutput {
	return o
}

func (o AppInsightsCredentialsResponsePtrOutput) ToAppInsightsCredentialsResponsePtrOutputWithContext(ctx context.Context) AppInsightsCredentialsResponsePtrOutput {
	return o
}

func (o AppInsightsCredentialsResponsePtrOutput) Elem() AppInsightsCredentialsResponseOutput {
	return o.ApplyT(func(v *AppInsightsCredentialsResponse) AppInsightsCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret AppInsightsCredentialsResponse
		return ret
	}).(AppInsightsCredentialsResponseOutput)
}

func (o AppInsightsCredentialsResponsePtrOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppInsightsCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiKey
	}).(pulumi.StringPtrOutput)
}

func (o AppInsightsCredentialsResponsePtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppInsightsCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

type AutoScaleConfiguration struct {
	MaxReplicas            *int     `pulumi:"maxReplicas"`
	MinReplicas            *int     `pulumi:"minReplicas"`
	RefreshPeriodInSeconds *int     `pulumi:"refreshPeriodInSeconds"`
	Status                 *string  `pulumi:"status"`
	TargetUtilization      *float64 `pulumi:"targetUtilization"`
}





type AutoScaleConfigurationInput interface {
	pulumi.Input

	ToAutoScaleConfigurationOutput() AutoScaleConfigurationOutput
	ToAutoScaleConfigurationOutputWithContext(context.Context) AutoScaleConfigurationOutput
}

type AutoScaleConfigurationArgs struct {
	MaxReplicas            pulumi.IntPtrInput     `pulumi:"maxReplicas"`
	MinReplicas            pulumi.IntPtrInput     `pulumi:"minReplicas"`
	RefreshPeriodInSeconds pulumi.IntPtrInput     `pulumi:"refreshPeriodInSeconds"`
	Status                 pulumi.StringPtrInput  `pulumi:"status"`
	TargetUtilization      pulumi.Float64PtrInput `pulumi:"targetUtilization"`
}

func (AutoScaleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleConfiguration)(nil)).Elem()
}

func (i AutoScaleConfigurationArgs) ToAutoScaleConfigurationOutput() AutoScaleConfigurationOutput {
	return i.ToAutoScaleConfigurationOutputWithContext(context.Background())
}

func (i AutoScaleConfigurationArgs) ToAutoScaleConfigurationOutputWithContext(ctx context.Context) AutoScaleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleConfigurationOutput)
}

func (i AutoScaleConfigurationArgs) ToAutoScaleConfigurationPtrOutput() AutoScaleConfigurationPtrOutput {
	return i.ToAutoScaleConfigurationPtrOutputWithContext(context.Background())
}

func (i AutoScaleConfigurationArgs) ToAutoScaleConfigurationPtrOutputWithContext(ctx context.Context) AutoScaleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleConfigurationOutput).ToAutoScaleConfigurationPtrOutputWithContext(ctx)
}









type AutoScaleConfigurationPtrInput interface {
	pulumi.Input

	ToAutoScaleConfigurationPtrOutput() AutoScaleConfigurationPtrOutput
	ToAutoScaleConfigurationPtrOutputWithContext(context.Context) AutoScaleConfigurationPtrOutput
}

type autoScaleConfigurationPtrType AutoScaleConfigurationArgs

func AutoScaleConfigurationPtr(v *AutoScaleConfigurationArgs) AutoScaleConfigurationPtrInput {
	return (*autoScaleConfigurationPtrType)(v)
}

func (*autoScaleConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleConfiguration)(nil)).Elem()
}

func (i *autoScaleConfigurationPtrType) ToAutoScaleConfigurationPtrOutput() AutoScaleConfigurationPtrOutput {
	return i.ToAutoScaleConfigurationPtrOutputWithContext(context.Background())
}

func (i *autoScaleConfigurationPtrType) ToAutoScaleConfigurationPtrOutputWithContext(ctx context.Context) AutoScaleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleConfigurationPtrOutput)
}

type AutoScaleConfigurationOutput struct{ *pulumi.OutputState }

func (AutoScaleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleConfiguration)(nil)).Elem()
}

func (o AutoScaleConfigurationOutput) ToAutoScaleConfigurationOutput() AutoScaleConfigurationOutput {
	return o
}

func (o AutoScaleConfigurationOutput) ToAutoScaleConfigurationOutputWithContext(ctx context.Context) AutoScaleConfigurationOutput {
	return o
}

func (o AutoScaleConfigurationOutput) ToAutoScaleConfigurationPtrOutput() AutoScaleConfigurationPtrOutput {
	return o.ToAutoScaleConfigurationPtrOutputWithContext(context.Background())
}

func (o AutoScaleConfigurationOutput) ToAutoScaleConfigurationPtrOutputWithContext(ctx context.Context) AutoScaleConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScaleConfiguration) *AutoScaleConfiguration {
		return &v
	}).(AutoScaleConfigurationPtrOutput)
}

func (o AutoScaleConfigurationOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleConfiguration) *int { return v.MaxReplicas }).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleConfiguration) *int { return v.MinReplicas }).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationOutput) RefreshPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleConfiguration) *int { return v.RefreshPeriodInSeconds }).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleConfiguration) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o AutoScaleConfigurationOutput) TargetUtilization() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AutoScaleConfiguration) *float64 { return v.TargetUtilization }).(pulumi.Float64PtrOutput)
}

type AutoScaleConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AutoScaleConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleConfiguration)(nil)).Elem()
}

func (o AutoScaleConfigurationPtrOutput) ToAutoScaleConfigurationPtrOutput() AutoScaleConfigurationPtrOutput {
	return o
}

func (o AutoScaleConfigurationPtrOutput) ToAutoScaleConfigurationPtrOutputWithContext(ctx context.Context) AutoScaleConfigurationPtrOutput {
	return o
}

func (o AutoScaleConfigurationPtrOutput) Elem() AutoScaleConfigurationOutput {
	return o.ApplyT(func(v *AutoScaleConfiguration) AutoScaleConfiguration {
		if v != nil {
			return *v
		}
		var ret AutoScaleConfiguration
		return ret
	}).(AutoScaleConfigurationOutput)
}

func (o AutoScaleConfigurationPtrOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.MaxReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationPtrOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.MinReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationPtrOutput) RefreshPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.RefreshPeriodInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o AutoScaleConfigurationPtrOutput) TargetUtilization() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AutoScaleConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return v.TargetUtilization
	}).(pulumi.Float64PtrOutput)
}

type AutoScaleConfigurationResponse struct {
	MaxReplicas            *int     `pulumi:"maxReplicas"`
	MinReplicas            *int     `pulumi:"minReplicas"`
	RefreshPeriodInSeconds *int     `pulumi:"refreshPeriodInSeconds"`
	Status                 *string  `pulumi:"status"`
	TargetUtilization      *float64 `pulumi:"targetUtilization"`
}





type AutoScaleConfigurationResponseInput interface {
	pulumi.Input

	ToAutoScaleConfigurationResponseOutput() AutoScaleConfigurationResponseOutput
	ToAutoScaleConfigurationResponseOutputWithContext(context.Context) AutoScaleConfigurationResponseOutput
}

type AutoScaleConfigurationResponseArgs struct {
	MaxReplicas            pulumi.IntPtrInput     `pulumi:"maxReplicas"`
	MinReplicas            pulumi.IntPtrInput     `pulumi:"minReplicas"`
	RefreshPeriodInSeconds pulumi.IntPtrInput     `pulumi:"refreshPeriodInSeconds"`
	Status                 pulumi.StringPtrInput  `pulumi:"status"`
	TargetUtilization      pulumi.Float64PtrInput `pulumi:"targetUtilization"`
}

func (AutoScaleConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleConfigurationResponse)(nil)).Elem()
}

func (i AutoScaleConfigurationResponseArgs) ToAutoScaleConfigurationResponseOutput() AutoScaleConfigurationResponseOutput {
	return i.ToAutoScaleConfigurationResponseOutputWithContext(context.Background())
}

func (i AutoScaleConfigurationResponseArgs) ToAutoScaleConfigurationResponseOutputWithContext(ctx context.Context) AutoScaleConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleConfigurationResponseOutput)
}

func (i AutoScaleConfigurationResponseArgs) ToAutoScaleConfigurationResponsePtrOutput() AutoScaleConfigurationResponsePtrOutput {
	return i.ToAutoScaleConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i AutoScaleConfigurationResponseArgs) ToAutoScaleConfigurationResponsePtrOutputWithContext(ctx context.Context) AutoScaleConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleConfigurationResponseOutput).ToAutoScaleConfigurationResponsePtrOutputWithContext(ctx)
}









type AutoScaleConfigurationResponsePtrInput interface {
	pulumi.Input

	ToAutoScaleConfigurationResponsePtrOutput() AutoScaleConfigurationResponsePtrOutput
	ToAutoScaleConfigurationResponsePtrOutputWithContext(context.Context) AutoScaleConfigurationResponsePtrOutput
}

type autoScaleConfigurationResponsePtrType AutoScaleConfigurationResponseArgs

func AutoScaleConfigurationResponsePtr(v *AutoScaleConfigurationResponseArgs) AutoScaleConfigurationResponsePtrInput {
	return (*autoScaleConfigurationResponsePtrType)(v)
}

func (*autoScaleConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleConfigurationResponse)(nil)).Elem()
}

func (i *autoScaleConfigurationResponsePtrType) ToAutoScaleConfigurationResponsePtrOutput() AutoScaleConfigurationResponsePtrOutput {
	return i.ToAutoScaleConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *autoScaleConfigurationResponsePtrType) ToAutoScaleConfigurationResponsePtrOutputWithContext(ctx context.Context) AutoScaleConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScaleConfigurationResponsePtrOutput)
}

type AutoScaleConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AutoScaleConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleConfigurationResponse)(nil)).Elem()
}

func (o AutoScaleConfigurationResponseOutput) ToAutoScaleConfigurationResponseOutput() AutoScaleConfigurationResponseOutput {
	return o
}

func (o AutoScaleConfigurationResponseOutput) ToAutoScaleConfigurationResponseOutputWithContext(ctx context.Context) AutoScaleConfigurationResponseOutput {
	return o
}

func (o AutoScaleConfigurationResponseOutput) ToAutoScaleConfigurationResponsePtrOutput() AutoScaleConfigurationResponsePtrOutput {
	return o.ToAutoScaleConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o AutoScaleConfigurationResponseOutput) ToAutoScaleConfigurationResponsePtrOutputWithContext(ctx context.Context) AutoScaleConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScaleConfigurationResponse) *AutoScaleConfigurationResponse {
		return &v
	}).(AutoScaleConfigurationResponsePtrOutput)
}

func (o AutoScaleConfigurationResponseOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleConfigurationResponse) *int { return v.MaxReplicas }).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationResponseOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleConfigurationResponse) *int { return v.MinReplicas }).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationResponseOutput) RefreshPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleConfigurationResponse) *int { return v.RefreshPeriodInSeconds }).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoScaleConfigurationResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o AutoScaleConfigurationResponseOutput) TargetUtilization() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AutoScaleConfigurationResponse) *float64 { return v.TargetUtilization }).(pulumi.Float64PtrOutput)
}

type AutoScaleConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoScaleConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleConfigurationResponse)(nil)).Elem()
}

func (o AutoScaleConfigurationResponsePtrOutput) ToAutoScaleConfigurationResponsePtrOutput() AutoScaleConfigurationResponsePtrOutput {
	return o
}

func (o AutoScaleConfigurationResponsePtrOutput) ToAutoScaleConfigurationResponsePtrOutputWithContext(ctx context.Context) AutoScaleConfigurationResponsePtrOutput {
	return o
}

func (o AutoScaleConfigurationResponsePtrOutput) Elem() AutoScaleConfigurationResponseOutput {
	return o.ApplyT(func(v *AutoScaleConfigurationResponse) AutoScaleConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AutoScaleConfigurationResponse
		return ret
	}).(AutoScaleConfigurationResponseOutput)
}

func (o AutoScaleConfigurationResponsePtrOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationResponsePtrOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinReplicas
	}).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationResponsePtrOutput) RefreshPeriodInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.RefreshPeriodInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o AutoScaleConfigurationResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoScaleConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o AutoScaleConfigurationResponsePtrOutput) TargetUtilization() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AutoScaleConfigurationResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TargetUtilization
	}).(pulumi.Float64PtrOutput)
}

type ContainerRegistryCredentialsResponse struct {
	LoginServer string `pulumi:"loginServer"`
	Password    string `pulumi:"password"`
	Password2   string `pulumi:"password2"`
}





type ContainerRegistryCredentialsResponseInput interface {
	pulumi.Input

	ToContainerRegistryCredentialsResponseOutput() ContainerRegistryCredentialsResponseOutput
	ToContainerRegistryCredentialsResponseOutputWithContext(context.Context) ContainerRegistryCredentialsResponseOutput
}

type ContainerRegistryCredentialsResponseArgs struct {
	LoginServer pulumi.StringInput `pulumi:"loginServer"`
	Password    pulumi.StringInput `pulumi:"password"`
	Password2   pulumi.StringInput `pulumi:"password2"`
}

func (ContainerRegistryCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryCredentialsResponse)(nil)).Elem()
}

func (i ContainerRegistryCredentialsResponseArgs) ToContainerRegistryCredentialsResponseOutput() ContainerRegistryCredentialsResponseOutput {
	return i.ToContainerRegistryCredentialsResponseOutputWithContext(context.Background())
}

func (i ContainerRegistryCredentialsResponseArgs) ToContainerRegistryCredentialsResponseOutputWithContext(ctx context.Context) ContainerRegistryCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryCredentialsResponseOutput)
}

func (i ContainerRegistryCredentialsResponseArgs) ToContainerRegistryCredentialsResponsePtrOutput() ContainerRegistryCredentialsResponsePtrOutput {
	return i.ToContainerRegistryCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i ContainerRegistryCredentialsResponseArgs) ToContainerRegistryCredentialsResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryCredentialsResponseOutput).ToContainerRegistryCredentialsResponsePtrOutputWithContext(ctx)
}









type ContainerRegistryCredentialsResponsePtrInput interface {
	pulumi.Input

	ToContainerRegistryCredentialsResponsePtrOutput() ContainerRegistryCredentialsResponsePtrOutput
	ToContainerRegistryCredentialsResponsePtrOutputWithContext(context.Context) ContainerRegistryCredentialsResponsePtrOutput
}

type containerRegistryCredentialsResponsePtrType ContainerRegistryCredentialsResponseArgs

func ContainerRegistryCredentialsResponsePtr(v *ContainerRegistryCredentialsResponseArgs) ContainerRegistryCredentialsResponsePtrInput {
	return (*containerRegistryCredentialsResponsePtrType)(v)
}

func (*containerRegistryCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryCredentialsResponse)(nil)).Elem()
}

func (i *containerRegistryCredentialsResponsePtrType) ToContainerRegistryCredentialsResponsePtrOutput() ContainerRegistryCredentialsResponsePtrOutput {
	return i.ToContainerRegistryCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *containerRegistryCredentialsResponsePtrType) ToContainerRegistryCredentialsResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryCredentialsResponsePtrOutput)
}

type ContainerRegistryCredentialsResponseOutput struct{ *pulumi.OutputState }

func (ContainerRegistryCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryCredentialsResponse)(nil)).Elem()
}

func (o ContainerRegistryCredentialsResponseOutput) ToContainerRegistryCredentialsResponseOutput() ContainerRegistryCredentialsResponseOutput {
	return o
}

func (o ContainerRegistryCredentialsResponseOutput) ToContainerRegistryCredentialsResponseOutputWithContext(ctx context.Context) ContainerRegistryCredentialsResponseOutput {
	return o
}

func (o ContainerRegistryCredentialsResponseOutput) ToContainerRegistryCredentialsResponsePtrOutput() ContainerRegistryCredentialsResponsePtrOutput {
	return o.ToContainerRegistryCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o ContainerRegistryCredentialsResponseOutput) ToContainerRegistryCredentialsResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerRegistryCredentialsResponse) *ContainerRegistryCredentialsResponse {
		return &v
	}).(ContainerRegistryCredentialsResponsePtrOutput)
}

func (o ContainerRegistryCredentialsResponseOutput) LoginServer() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerRegistryCredentialsResponse) string { return v.LoginServer }).(pulumi.StringOutput)
}

func (o ContainerRegistryCredentialsResponseOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerRegistryCredentialsResponse) string { return v.Password }).(pulumi.StringOutput)
}

func (o ContainerRegistryCredentialsResponseOutput) Password2() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerRegistryCredentialsResponse) string { return v.Password2 }).(pulumi.StringOutput)
}

type ContainerRegistryCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerRegistryCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryCredentialsResponse)(nil)).Elem()
}

func (o ContainerRegistryCredentialsResponsePtrOutput) ToContainerRegistryCredentialsResponsePtrOutput() ContainerRegistryCredentialsResponsePtrOutput {
	return o
}

func (o ContainerRegistryCredentialsResponsePtrOutput) ToContainerRegistryCredentialsResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryCredentialsResponsePtrOutput {
	return o
}

func (o ContainerRegistryCredentialsResponsePtrOutput) Elem() ContainerRegistryCredentialsResponseOutput {
	return o.ApplyT(func(v *ContainerRegistryCredentialsResponse) ContainerRegistryCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret ContainerRegistryCredentialsResponse
		return ret
	}).(ContainerRegistryCredentialsResponseOutput)
}

func (o ContainerRegistryCredentialsResponsePtrOutput) LoginServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LoginServer
	}).(pulumi.StringPtrOutput)
}

func (o ContainerRegistryCredentialsResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ContainerRegistryCredentialsResponsePtrOutput) Password2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Password2
	}).(pulumi.StringPtrOutput)
}

type ContainerRegistryProperties struct {
	ResourceId *string `pulumi:"resourceId"`
}





type ContainerRegistryPropertiesInput interface {
	pulumi.Input

	ToContainerRegistryPropertiesOutput() ContainerRegistryPropertiesOutput
	ToContainerRegistryPropertiesOutputWithContext(context.Context) ContainerRegistryPropertiesOutput
}

type ContainerRegistryPropertiesArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ContainerRegistryPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryProperties)(nil)).Elem()
}

func (i ContainerRegistryPropertiesArgs) ToContainerRegistryPropertiesOutput() ContainerRegistryPropertiesOutput {
	return i.ToContainerRegistryPropertiesOutputWithContext(context.Background())
}

func (i ContainerRegistryPropertiesArgs) ToContainerRegistryPropertiesOutputWithContext(ctx context.Context) ContainerRegistryPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryPropertiesOutput)
}

func (i ContainerRegistryPropertiesArgs) ToContainerRegistryPropertiesPtrOutput() ContainerRegistryPropertiesPtrOutput {
	return i.ToContainerRegistryPropertiesPtrOutputWithContext(context.Background())
}

func (i ContainerRegistryPropertiesArgs) ToContainerRegistryPropertiesPtrOutputWithContext(ctx context.Context) ContainerRegistryPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryPropertiesOutput).ToContainerRegistryPropertiesPtrOutputWithContext(ctx)
}









type ContainerRegistryPropertiesPtrInput interface {
	pulumi.Input

	ToContainerRegistryPropertiesPtrOutput() ContainerRegistryPropertiesPtrOutput
	ToContainerRegistryPropertiesPtrOutputWithContext(context.Context) ContainerRegistryPropertiesPtrOutput
}

type containerRegistryPropertiesPtrType ContainerRegistryPropertiesArgs

func ContainerRegistryPropertiesPtr(v *ContainerRegistryPropertiesArgs) ContainerRegistryPropertiesPtrInput {
	return (*containerRegistryPropertiesPtrType)(v)
}

func (*containerRegistryPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryProperties)(nil)).Elem()
}

func (i *containerRegistryPropertiesPtrType) ToContainerRegistryPropertiesPtrOutput() ContainerRegistryPropertiesPtrOutput {
	return i.ToContainerRegistryPropertiesPtrOutputWithContext(context.Background())
}

func (i *containerRegistryPropertiesPtrType) ToContainerRegistryPropertiesPtrOutputWithContext(ctx context.Context) ContainerRegistryPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryPropertiesPtrOutput)
}

type ContainerRegistryPropertiesOutput struct{ *pulumi.OutputState }

func (ContainerRegistryPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryProperties)(nil)).Elem()
}

func (o ContainerRegistryPropertiesOutput) ToContainerRegistryPropertiesOutput() ContainerRegistryPropertiesOutput {
	return o
}

func (o ContainerRegistryPropertiesOutput) ToContainerRegistryPropertiesOutputWithContext(ctx context.Context) ContainerRegistryPropertiesOutput {
	return o
}

func (o ContainerRegistryPropertiesOutput) ToContainerRegistryPropertiesPtrOutput() ContainerRegistryPropertiesPtrOutput {
	return o.ToContainerRegistryPropertiesPtrOutputWithContext(context.Background())
}

func (o ContainerRegistryPropertiesOutput) ToContainerRegistryPropertiesPtrOutputWithContext(ctx context.Context) ContainerRegistryPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerRegistryProperties) *ContainerRegistryProperties {
		return &v
	}).(ContainerRegistryPropertiesPtrOutput)
}

func (o ContainerRegistryPropertiesOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerRegistryProperties) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ContainerRegistryPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ContainerRegistryPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryProperties)(nil)).Elem()
}

func (o ContainerRegistryPropertiesPtrOutput) ToContainerRegistryPropertiesPtrOutput() ContainerRegistryPropertiesPtrOutput {
	return o
}

func (o ContainerRegistryPropertiesPtrOutput) ToContainerRegistryPropertiesPtrOutputWithContext(ctx context.Context) ContainerRegistryPropertiesPtrOutput {
	return o
}

func (o ContainerRegistryPropertiesPtrOutput) Elem() ContainerRegistryPropertiesOutput {
	return o.ApplyT(func(v *ContainerRegistryProperties) ContainerRegistryProperties {
		if v != nil {
			return *v
		}
		var ret ContainerRegistryProperties
		return ret
	}).(ContainerRegistryPropertiesOutput)
}

func (o ContainerRegistryPropertiesPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type ContainerRegistryPropertiesResponse struct {
	ResourceId *string `pulumi:"resourceId"`
}





type ContainerRegistryPropertiesResponseInput interface {
	pulumi.Input

	ToContainerRegistryPropertiesResponseOutput() ContainerRegistryPropertiesResponseOutput
	ToContainerRegistryPropertiesResponseOutputWithContext(context.Context) ContainerRegistryPropertiesResponseOutput
}

type ContainerRegistryPropertiesResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ContainerRegistryPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryPropertiesResponse)(nil)).Elem()
}

func (i ContainerRegistryPropertiesResponseArgs) ToContainerRegistryPropertiesResponseOutput() ContainerRegistryPropertiesResponseOutput {
	return i.ToContainerRegistryPropertiesResponseOutputWithContext(context.Background())
}

func (i ContainerRegistryPropertiesResponseArgs) ToContainerRegistryPropertiesResponseOutputWithContext(ctx context.Context) ContainerRegistryPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryPropertiesResponseOutput)
}

func (i ContainerRegistryPropertiesResponseArgs) ToContainerRegistryPropertiesResponsePtrOutput() ContainerRegistryPropertiesResponsePtrOutput {
	return i.ToContainerRegistryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ContainerRegistryPropertiesResponseArgs) ToContainerRegistryPropertiesResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryPropertiesResponseOutput).ToContainerRegistryPropertiesResponsePtrOutputWithContext(ctx)
}









type ContainerRegistryPropertiesResponsePtrInput interface {
	pulumi.Input

	ToContainerRegistryPropertiesResponsePtrOutput() ContainerRegistryPropertiesResponsePtrOutput
	ToContainerRegistryPropertiesResponsePtrOutputWithContext(context.Context) ContainerRegistryPropertiesResponsePtrOutput
}

type containerRegistryPropertiesResponsePtrType ContainerRegistryPropertiesResponseArgs

func ContainerRegistryPropertiesResponsePtr(v *ContainerRegistryPropertiesResponseArgs) ContainerRegistryPropertiesResponsePtrInput {
	return (*containerRegistryPropertiesResponsePtrType)(v)
}

func (*containerRegistryPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryPropertiesResponse)(nil)).Elem()
}

func (i *containerRegistryPropertiesResponsePtrType) ToContainerRegistryPropertiesResponsePtrOutput() ContainerRegistryPropertiesResponsePtrOutput {
	return i.ToContainerRegistryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *containerRegistryPropertiesResponsePtrType) ToContainerRegistryPropertiesResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryPropertiesResponsePtrOutput)
}

type ContainerRegistryPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ContainerRegistryPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryPropertiesResponse)(nil)).Elem()
}

func (o ContainerRegistryPropertiesResponseOutput) ToContainerRegistryPropertiesResponseOutput() ContainerRegistryPropertiesResponseOutput {
	return o
}

func (o ContainerRegistryPropertiesResponseOutput) ToContainerRegistryPropertiesResponseOutputWithContext(ctx context.Context) ContainerRegistryPropertiesResponseOutput {
	return o
}

func (o ContainerRegistryPropertiesResponseOutput) ToContainerRegistryPropertiesResponsePtrOutput() ContainerRegistryPropertiesResponsePtrOutput {
	return o.ToContainerRegistryPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ContainerRegistryPropertiesResponseOutput) ToContainerRegistryPropertiesResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerRegistryPropertiesResponse) *ContainerRegistryPropertiesResponse {
		return &v
	}).(ContainerRegistryPropertiesResponsePtrOutput)
}

func (o ContainerRegistryPropertiesResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerRegistryPropertiesResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ContainerRegistryPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerRegistryPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryPropertiesResponse)(nil)).Elem()
}

func (o ContainerRegistryPropertiesResponsePtrOutput) ToContainerRegistryPropertiesResponsePtrOutput() ContainerRegistryPropertiesResponsePtrOutput {
	return o
}

func (o ContainerRegistryPropertiesResponsePtrOutput) ToContainerRegistryPropertiesResponsePtrOutputWithContext(ctx context.Context) ContainerRegistryPropertiesResponsePtrOutput {
	return o
}

func (o ContainerRegistryPropertiesResponsePtrOutput) Elem() ContainerRegistryPropertiesResponseOutput {
	return o.ApplyT(func(v *ContainerRegistryPropertiesResponse) ContainerRegistryPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ContainerRegistryPropertiesResponse
		return ret
	}).(ContainerRegistryPropertiesResponseOutput)
}

func (o ContainerRegistryPropertiesResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceCredentialsResponse struct {
	AcsKubeConfig                 string                             `pulumi:"acsKubeConfig"`
	ImagePullSecretName           string                             `pulumi:"imagePullSecretName"`
	ServicePrincipalConfiguration ServicePrincipalPropertiesResponse `pulumi:"servicePrincipalConfiguration"`
}





type ContainerServiceCredentialsResponseInput interface {
	pulumi.Input

	ToContainerServiceCredentialsResponseOutput() ContainerServiceCredentialsResponseOutput
	ToContainerServiceCredentialsResponseOutputWithContext(context.Context) ContainerServiceCredentialsResponseOutput
}

type ContainerServiceCredentialsResponseArgs struct {
	AcsKubeConfig                 pulumi.StringInput                      `pulumi:"acsKubeConfig"`
	ImagePullSecretName           pulumi.StringInput                      `pulumi:"imagePullSecretName"`
	ServicePrincipalConfiguration ServicePrincipalPropertiesResponseInput `pulumi:"servicePrincipalConfiguration"`
}

func (ContainerServiceCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceCredentialsResponse)(nil)).Elem()
}

func (i ContainerServiceCredentialsResponseArgs) ToContainerServiceCredentialsResponseOutput() ContainerServiceCredentialsResponseOutput {
	return i.ToContainerServiceCredentialsResponseOutputWithContext(context.Background())
}

func (i ContainerServiceCredentialsResponseArgs) ToContainerServiceCredentialsResponseOutputWithContext(ctx context.Context) ContainerServiceCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCredentialsResponseOutput)
}

func (i ContainerServiceCredentialsResponseArgs) ToContainerServiceCredentialsResponsePtrOutput() ContainerServiceCredentialsResponsePtrOutput {
	return i.ToContainerServiceCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceCredentialsResponseArgs) ToContainerServiceCredentialsResponsePtrOutputWithContext(ctx context.Context) ContainerServiceCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCredentialsResponseOutput).ToContainerServiceCredentialsResponsePtrOutputWithContext(ctx)
}









type ContainerServiceCredentialsResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceCredentialsResponsePtrOutput() ContainerServiceCredentialsResponsePtrOutput
	ToContainerServiceCredentialsResponsePtrOutputWithContext(context.Context) ContainerServiceCredentialsResponsePtrOutput
}

type containerServiceCredentialsResponsePtrType ContainerServiceCredentialsResponseArgs

func ContainerServiceCredentialsResponsePtr(v *ContainerServiceCredentialsResponseArgs) ContainerServiceCredentialsResponsePtrInput {
	return (*containerServiceCredentialsResponsePtrType)(v)
}

func (*containerServiceCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceCredentialsResponse)(nil)).Elem()
}

func (i *containerServiceCredentialsResponsePtrType) ToContainerServiceCredentialsResponsePtrOutput() ContainerServiceCredentialsResponsePtrOutput {
	return i.ToContainerServiceCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceCredentialsResponsePtrType) ToContainerServiceCredentialsResponsePtrOutputWithContext(ctx context.Context) ContainerServiceCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCredentialsResponsePtrOutput)
}

type ContainerServiceCredentialsResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceCredentialsResponse)(nil)).Elem()
}

func (o ContainerServiceCredentialsResponseOutput) ToContainerServiceCredentialsResponseOutput() ContainerServiceCredentialsResponseOutput {
	return o
}

func (o ContainerServiceCredentialsResponseOutput) ToContainerServiceCredentialsResponseOutputWithContext(ctx context.Context) ContainerServiceCredentialsResponseOutput {
	return o
}

func (o ContainerServiceCredentialsResponseOutput) ToContainerServiceCredentialsResponsePtrOutput() ContainerServiceCredentialsResponsePtrOutput {
	return o.ToContainerServiceCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceCredentialsResponseOutput) ToContainerServiceCredentialsResponsePtrOutputWithContext(ctx context.Context) ContainerServiceCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceCredentialsResponse) *ContainerServiceCredentialsResponse {
		return &v
	}).(ContainerServiceCredentialsResponsePtrOutput)
}

func (o ContainerServiceCredentialsResponseOutput) AcsKubeConfig() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceCredentialsResponse) string { return v.AcsKubeConfig }).(pulumi.StringOutput)
}

func (o ContainerServiceCredentialsResponseOutput) ImagePullSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceCredentialsResponse) string { return v.ImagePullSecretName }).(pulumi.StringOutput)
}

func (o ContainerServiceCredentialsResponseOutput) ServicePrincipalConfiguration() ServicePrincipalPropertiesResponseOutput {
	return o.ApplyT(func(v ContainerServiceCredentialsResponse) ServicePrincipalPropertiesResponse {
		return v.ServicePrincipalConfiguration
	}).(ServicePrincipalPropertiesResponseOutput)
}

type ContainerServiceCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceCredentialsResponse)(nil)).Elem()
}

func (o ContainerServiceCredentialsResponsePtrOutput) ToContainerServiceCredentialsResponsePtrOutput() ContainerServiceCredentialsResponsePtrOutput {
	return o
}

func (o ContainerServiceCredentialsResponsePtrOutput) ToContainerServiceCredentialsResponsePtrOutputWithContext(ctx context.Context) ContainerServiceCredentialsResponsePtrOutput {
	return o
}

func (o ContainerServiceCredentialsResponsePtrOutput) Elem() ContainerServiceCredentialsResponseOutput {
	return o.ApplyT(func(v *ContainerServiceCredentialsResponse) ContainerServiceCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceCredentialsResponse
		return ret
	}).(ContainerServiceCredentialsResponseOutput)
}

func (o ContainerServiceCredentialsResponsePtrOutput) AcsKubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AcsKubeConfig
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceCredentialsResponsePtrOutput) ImagePullSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ImagePullSecretName
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceCredentialsResponsePtrOutput) ServicePrincipalConfiguration() ServicePrincipalPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ContainerServiceCredentialsResponse) *ServicePrincipalPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.ServicePrincipalConfiguration
	}).(ServicePrincipalPropertiesResponsePtrOutput)
}

type GlobalServiceConfiguration struct {
	AutoScale   *AutoScaleConfiguration   `pulumi:"autoScale"`
	Etag        *string                   `pulumi:"etag"`
	ServiceAuth *ServiceAuthConfiguration `pulumi:"serviceAuth"`
	Ssl         *SslConfiguration         `pulumi:"ssl"`
}





type GlobalServiceConfigurationInput interface {
	pulumi.Input

	ToGlobalServiceConfigurationOutput() GlobalServiceConfigurationOutput
	ToGlobalServiceConfigurationOutputWithContext(context.Context) GlobalServiceConfigurationOutput
}

type GlobalServiceConfigurationArgs struct {
	AutoScale   AutoScaleConfigurationPtrInput   `pulumi:"autoScale"`
	Etag        pulumi.StringPtrInput            `pulumi:"etag"`
	ServiceAuth ServiceAuthConfigurationPtrInput `pulumi:"serviceAuth"`
	Ssl         SslConfigurationPtrInput         `pulumi:"ssl"`
}

func (GlobalServiceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalServiceConfiguration)(nil)).Elem()
}

func (i GlobalServiceConfigurationArgs) ToGlobalServiceConfigurationOutput() GlobalServiceConfigurationOutput {
	return i.ToGlobalServiceConfigurationOutputWithContext(context.Background())
}

func (i GlobalServiceConfigurationArgs) ToGlobalServiceConfigurationOutputWithContext(ctx context.Context) GlobalServiceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalServiceConfigurationOutput)
}

func (i GlobalServiceConfigurationArgs) ToGlobalServiceConfigurationPtrOutput() GlobalServiceConfigurationPtrOutput {
	return i.ToGlobalServiceConfigurationPtrOutputWithContext(context.Background())
}

func (i GlobalServiceConfigurationArgs) ToGlobalServiceConfigurationPtrOutputWithContext(ctx context.Context) GlobalServiceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalServiceConfigurationOutput).ToGlobalServiceConfigurationPtrOutputWithContext(ctx)
}









type GlobalServiceConfigurationPtrInput interface {
	pulumi.Input

	ToGlobalServiceConfigurationPtrOutput() GlobalServiceConfigurationPtrOutput
	ToGlobalServiceConfigurationPtrOutputWithContext(context.Context) GlobalServiceConfigurationPtrOutput
}

type globalServiceConfigurationPtrType GlobalServiceConfigurationArgs

func GlobalServiceConfigurationPtr(v *GlobalServiceConfigurationArgs) GlobalServiceConfigurationPtrInput {
	return (*globalServiceConfigurationPtrType)(v)
}

func (*globalServiceConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalServiceConfiguration)(nil)).Elem()
}

func (i *globalServiceConfigurationPtrType) ToGlobalServiceConfigurationPtrOutput() GlobalServiceConfigurationPtrOutput {
	return i.ToGlobalServiceConfigurationPtrOutputWithContext(context.Background())
}

func (i *globalServiceConfigurationPtrType) ToGlobalServiceConfigurationPtrOutputWithContext(ctx context.Context) GlobalServiceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalServiceConfigurationPtrOutput)
}

type GlobalServiceConfigurationOutput struct{ *pulumi.OutputState }

func (GlobalServiceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalServiceConfiguration)(nil)).Elem()
}

func (o GlobalServiceConfigurationOutput) ToGlobalServiceConfigurationOutput() GlobalServiceConfigurationOutput {
	return o
}

func (o GlobalServiceConfigurationOutput) ToGlobalServiceConfigurationOutputWithContext(ctx context.Context) GlobalServiceConfigurationOutput {
	return o
}

func (o GlobalServiceConfigurationOutput) ToGlobalServiceConfigurationPtrOutput() GlobalServiceConfigurationPtrOutput {
	return o.ToGlobalServiceConfigurationPtrOutputWithContext(context.Background())
}

func (o GlobalServiceConfigurationOutput) ToGlobalServiceConfigurationPtrOutputWithContext(ctx context.Context) GlobalServiceConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GlobalServiceConfiguration) *GlobalServiceConfiguration {
		return &v
	}).(GlobalServiceConfigurationPtrOutput)
}

func (o GlobalServiceConfigurationOutput) AutoScale() AutoScaleConfigurationPtrOutput {
	return o.ApplyT(func(v GlobalServiceConfiguration) *AutoScaleConfiguration { return v.AutoScale }).(AutoScaleConfigurationPtrOutput)
}

func (o GlobalServiceConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalServiceConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GlobalServiceConfigurationOutput) ServiceAuth() ServiceAuthConfigurationPtrOutput {
	return o.ApplyT(func(v GlobalServiceConfiguration) *ServiceAuthConfiguration { return v.ServiceAuth }).(ServiceAuthConfigurationPtrOutput)
}

func (o GlobalServiceConfigurationOutput) Ssl() SslConfigurationPtrOutput {
	return o.ApplyT(func(v GlobalServiceConfiguration) *SslConfiguration { return v.Ssl }).(SslConfigurationPtrOutput)
}

type GlobalServiceConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GlobalServiceConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalServiceConfiguration)(nil)).Elem()
}

func (o GlobalServiceConfigurationPtrOutput) ToGlobalServiceConfigurationPtrOutput() GlobalServiceConfigurationPtrOutput {
	return o
}

func (o GlobalServiceConfigurationPtrOutput) ToGlobalServiceConfigurationPtrOutputWithContext(ctx context.Context) GlobalServiceConfigurationPtrOutput {
	return o
}

func (o GlobalServiceConfigurationPtrOutput) Elem() GlobalServiceConfigurationOutput {
	return o.ApplyT(func(v *GlobalServiceConfiguration) GlobalServiceConfiguration {
		if v != nil {
			return *v
		}
		var ret GlobalServiceConfiguration
		return ret
	}).(GlobalServiceConfigurationOutput)
}

func (o GlobalServiceConfigurationPtrOutput) AutoScale() AutoScaleConfigurationPtrOutput {
	return o.ApplyT(func(v *GlobalServiceConfiguration) *AutoScaleConfiguration {
		if v == nil {
			return nil
		}
		return v.AutoScale
	}).(AutoScaleConfigurationPtrOutput)
}

func (o GlobalServiceConfigurationPtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalServiceConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o GlobalServiceConfigurationPtrOutput) ServiceAuth() ServiceAuthConfigurationPtrOutput {
	return o.ApplyT(func(v *GlobalServiceConfiguration) *ServiceAuthConfiguration {
		if v == nil {
			return nil
		}
		return v.ServiceAuth
	}).(ServiceAuthConfigurationPtrOutput)
}

func (o GlobalServiceConfigurationPtrOutput) Ssl() SslConfigurationPtrOutput {
	return o.ApplyT(func(v *GlobalServiceConfiguration) *SslConfiguration {
		if v == nil {
			return nil
		}
		return v.Ssl
	}).(SslConfigurationPtrOutput)
}

type GlobalServiceConfigurationResponse struct {
	AutoScale   *AutoScaleConfigurationResponse   `pulumi:"autoScale"`
	Etag        *string                           `pulumi:"etag"`
	ServiceAuth *ServiceAuthConfigurationResponse `pulumi:"serviceAuth"`
	Ssl         *SslConfigurationResponse         `pulumi:"ssl"`
}





type GlobalServiceConfigurationResponseInput interface {
	pulumi.Input

	ToGlobalServiceConfigurationResponseOutput() GlobalServiceConfigurationResponseOutput
	ToGlobalServiceConfigurationResponseOutputWithContext(context.Context) GlobalServiceConfigurationResponseOutput
}

type GlobalServiceConfigurationResponseArgs struct {
	AutoScale   AutoScaleConfigurationResponsePtrInput   `pulumi:"autoScale"`
	Etag        pulumi.StringPtrInput                    `pulumi:"etag"`
	ServiceAuth ServiceAuthConfigurationResponsePtrInput `pulumi:"serviceAuth"`
	Ssl         SslConfigurationResponsePtrInput         `pulumi:"ssl"`
}

func (GlobalServiceConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalServiceConfigurationResponse)(nil)).Elem()
}

func (i GlobalServiceConfigurationResponseArgs) ToGlobalServiceConfigurationResponseOutput() GlobalServiceConfigurationResponseOutput {
	return i.ToGlobalServiceConfigurationResponseOutputWithContext(context.Background())
}

func (i GlobalServiceConfigurationResponseArgs) ToGlobalServiceConfigurationResponseOutputWithContext(ctx context.Context) GlobalServiceConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalServiceConfigurationResponseOutput)
}

func (i GlobalServiceConfigurationResponseArgs) ToGlobalServiceConfigurationResponsePtrOutput() GlobalServiceConfigurationResponsePtrOutput {
	return i.ToGlobalServiceConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i GlobalServiceConfigurationResponseArgs) ToGlobalServiceConfigurationResponsePtrOutputWithContext(ctx context.Context) GlobalServiceConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalServiceConfigurationResponseOutput).ToGlobalServiceConfigurationResponsePtrOutputWithContext(ctx)
}









type GlobalServiceConfigurationResponsePtrInput interface {
	pulumi.Input

	ToGlobalServiceConfigurationResponsePtrOutput() GlobalServiceConfigurationResponsePtrOutput
	ToGlobalServiceConfigurationResponsePtrOutputWithContext(context.Context) GlobalServiceConfigurationResponsePtrOutput
}

type globalServiceConfigurationResponsePtrType GlobalServiceConfigurationResponseArgs

func GlobalServiceConfigurationResponsePtr(v *GlobalServiceConfigurationResponseArgs) GlobalServiceConfigurationResponsePtrInput {
	return (*globalServiceConfigurationResponsePtrType)(v)
}

func (*globalServiceConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalServiceConfigurationResponse)(nil)).Elem()
}

func (i *globalServiceConfigurationResponsePtrType) ToGlobalServiceConfigurationResponsePtrOutput() GlobalServiceConfigurationResponsePtrOutput {
	return i.ToGlobalServiceConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *globalServiceConfigurationResponsePtrType) ToGlobalServiceConfigurationResponsePtrOutputWithContext(ctx context.Context) GlobalServiceConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalServiceConfigurationResponsePtrOutput)
}

type GlobalServiceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GlobalServiceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalServiceConfigurationResponse)(nil)).Elem()
}

func (o GlobalServiceConfigurationResponseOutput) ToGlobalServiceConfigurationResponseOutput() GlobalServiceConfigurationResponseOutput {
	return o
}

func (o GlobalServiceConfigurationResponseOutput) ToGlobalServiceConfigurationResponseOutputWithContext(ctx context.Context) GlobalServiceConfigurationResponseOutput {
	return o
}

func (o GlobalServiceConfigurationResponseOutput) ToGlobalServiceConfigurationResponsePtrOutput() GlobalServiceConfigurationResponsePtrOutput {
	return o.ToGlobalServiceConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o GlobalServiceConfigurationResponseOutput) ToGlobalServiceConfigurationResponsePtrOutputWithContext(ctx context.Context) GlobalServiceConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GlobalServiceConfigurationResponse) *GlobalServiceConfigurationResponse {
		return &v
	}).(GlobalServiceConfigurationResponsePtrOutput)
}

func (o GlobalServiceConfigurationResponseOutput) AutoScale() AutoScaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v GlobalServiceConfigurationResponse) *AutoScaleConfigurationResponse { return v.AutoScale }).(AutoScaleConfigurationResponsePtrOutput)
}

func (o GlobalServiceConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalServiceConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GlobalServiceConfigurationResponseOutput) ServiceAuth() ServiceAuthConfigurationResponsePtrOutput {
	return o.ApplyT(func(v GlobalServiceConfigurationResponse) *ServiceAuthConfigurationResponse { return v.ServiceAuth }).(ServiceAuthConfigurationResponsePtrOutput)
}

func (o GlobalServiceConfigurationResponseOutput) Ssl() SslConfigurationResponsePtrOutput {
	return o.ApplyT(func(v GlobalServiceConfigurationResponse) *SslConfigurationResponse { return v.Ssl }).(SslConfigurationResponsePtrOutput)
}

type GlobalServiceConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GlobalServiceConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalServiceConfigurationResponse)(nil)).Elem()
}

func (o GlobalServiceConfigurationResponsePtrOutput) ToGlobalServiceConfigurationResponsePtrOutput() GlobalServiceConfigurationResponsePtrOutput {
	return o
}

func (o GlobalServiceConfigurationResponsePtrOutput) ToGlobalServiceConfigurationResponsePtrOutputWithContext(ctx context.Context) GlobalServiceConfigurationResponsePtrOutput {
	return o
}

func (o GlobalServiceConfigurationResponsePtrOutput) Elem() GlobalServiceConfigurationResponseOutput {
	return o.ApplyT(func(v *GlobalServiceConfigurationResponse) GlobalServiceConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GlobalServiceConfigurationResponse
		return ret
	}).(GlobalServiceConfigurationResponseOutput)
}

func (o GlobalServiceConfigurationResponsePtrOutput) AutoScale() AutoScaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GlobalServiceConfigurationResponse) *AutoScaleConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.AutoScale
	}).(AutoScaleConfigurationResponsePtrOutput)
}

func (o GlobalServiceConfigurationResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalServiceConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o GlobalServiceConfigurationResponsePtrOutput) ServiceAuth() ServiceAuthConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GlobalServiceConfigurationResponse) *ServiceAuthConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.ServiceAuth
	}).(ServiceAuthConfigurationResponsePtrOutput)
}

func (o GlobalServiceConfigurationResponsePtrOutput) Ssl() SslConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GlobalServiceConfigurationResponse) *SslConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Ssl
	}).(SslConfigurationResponsePtrOutput)
}

type KubernetesClusterProperties struct {
	ServicePrincipal ServicePrincipalProperties `pulumi:"servicePrincipal"`
}





type KubernetesClusterPropertiesInput interface {
	pulumi.Input

	ToKubernetesClusterPropertiesOutput() KubernetesClusterPropertiesOutput
	ToKubernetesClusterPropertiesOutputWithContext(context.Context) KubernetesClusterPropertiesOutput
}

type KubernetesClusterPropertiesArgs struct {
	ServicePrincipal ServicePrincipalPropertiesInput `pulumi:"servicePrincipal"`
}

func (KubernetesClusterPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterProperties)(nil)).Elem()
}

func (i KubernetesClusterPropertiesArgs) ToKubernetesClusterPropertiesOutput() KubernetesClusterPropertiesOutput {
	return i.ToKubernetesClusterPropertiesOutputWithContext(context.Background())
}

func (i KubernetesClusterPropertiesArgs) ToKubernetesClusterPropertiesOutputWithContext(ctx context.Context) KubernetesClusterPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPropertiesOutput)
}

func (i KubernetesClusterPropertiesArgs) ToKubernetesClusterPropertiesPtrOutput() KubernetesClusterPropertiesPtrOutput {
	return i.ToKubernetesClusterPropertiesPtrOutputWithContext(context.Background())
}

func (i KubernetesClusterPropertiesArgs) ToKubernetesClusterPropertiesPtrOutputWithContext(ctx context.Context) KubernetesClusterPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPropertiesOutput).ToKubernetesClusterPropertiesPtrOutputWithContext(ctx)
}









type KubernetesClusterPropertiesPtrInput interface {
	pulumi.Input

	ToKubernetesClusterPropertiesPtrOutput() KubernetesClusterPropertiesPtrOutput
	ToKubernetesClusterPropertiesPtrOutputWithContext(context.Context) KubernetesClusterPropertiesPtrOutput
}

type kubernetesClusterPropertiesPtrType KubernetesClusterPropertiesArgs

func KubernetesClusterPropertiesPtr(v *KubernetesClusterPropertiesArgs) KubernetesClusterPropertiesPtrInput {
	return (*kubernetesClusterPropertiesPtrType)(v)
}

func (*kubernetesClusterPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterProperties)(nil)).Elem()
}

func (i *kubernetesClusterPropertiesPtrType) ToKubernetesClusterPropertiesPtrOutput() KubernetesClusterPropertiesPtrOutput {
	return i.ToKubernetesClusterPropertiesPtrOutputWithContext(context.Background())
}

func (i *kubernetesClusterPropertiesPtrType) ToKubernetesClusterPropertiesPtrOutputWithContext(ctx context.Context) KubernetesClusterPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPropertiesPtrOutput)
}

type KubernetesClusterPropertiesOutput struct{ *pulumi.OutputState }

func (KubernetesClusterPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterProperties)(nil)).Elem()
}

func (o KubernetesClusterPropertiesOutput) ToKubernetesClusterPropertiesOutput() KubernetesClusterPropertiesOutput {
	return o
}

func (o KubernetesClusterPropertiesOutput) ToKubernetesClusterPropertiesOutputWithContext(ctx context.Context) KubernetesClusterPropertiesOutput {
	return o
}

func (o KubernetesClusterPropertiesOutput) ToKubernetesClusterPropertiesPtrOutput() KubernetesClusterPropertiesPtrOutput {
	return o.ToKubernetesClusterPropertiesPtrOutputWithContext(context.Background())
}

func (o KubernetesClusterPropertiesOutput) ToKubernetesClusterPropertiesPtrOutputWithContext(ctx context.Context) KubernetesClusterPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesClusterProperties) *KubernetesClusterProperties {
		return &v
	}).(KubernetesClusterPropertiesPtrOutput)
}

func (o KubernetesClusterPropertiesOutput) ServicePrincipal() ServicePrincipalPropertiesOutput {
	return o.ApplyT(func(v KubernetesClusterProperties) ServicePrincipalProperties { return v.ServicePrincipal }).(ServicePrincipalPropertiesOutput)
}

type KubernetesClusterPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KubernetesClusterPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterProperties)(nil)).Elem()
}

func (o KubernetesClusterPropertiesPtrOutput) ToKubernetesClusterPropertiesPtrOutput() KubernetesClusterPropertiesPtrOutput {
	return o
}

func (o KubernetesClusterPropertiesPtrOutput) ToKubernetesClusterPropertiesPtrOutputWithContext(ctx context.Context) KubernetesClusterPropertiesPtrOutput {
	return o
}

func (o KubernetesClusterPropertiesPtrOutput) Elem() KubernetesClusterPropertiesOutput {
	return o.ApplyT(func(v *KubernetesClusterProperties) KubernetesClusterProperties {
		if v != nil {
			return *v
		}
		var ret KubernetesClusterProperties
		return ret
	}).(KubernetesClusterPropertiesOutput)
}

func (o KubernetesClusterPropertiesPtrOutput) ServicePrincipal() ServicePrincipalPropertiesPtrOutput {
	return o.ApplyT(func(v *KubernetesClusterProperties) *ServicePrincipalProperties {
		if v == nil {
			return nil
		}
		return &v.ServicePrincipal
	}).(ServicePrincipalPropertiesPtrOutput)
}

type KubernetesClusterPropertiesResponse struct {
	ServicePrincipal ServicePrincipalPropertiesResponse `pulumi:"servicePrincipal"`
}





type KubernetesClusterPropertiesResponseInput interface {
	pulumi.Input

	ToKubernetesClusterPropertiesResponseOutput() KubernetesClusterPropertiesResponseOutput
	ToKubernetesClusterPropertiesResponseOutputWithContext(context.Context) KubernetesClusterPropertiesResponseOutput
}

type KubernetesClusterPropertiesResponseArgs struct {
	ServicePrincipal ServicePrincipalPropertiesResponseInput `pulumi:"servicePrincipal"`
}

func (KubernetesClusterPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterPropertiesResponse)(nil)).Elem()
}

func (i KubernetesClusterPropertiesResponseArgs) ToKubernetesClusterPropertiesResponseOutput() KubernetesClusterPropertiesResponseOutput {
	return i.ToKubernetesClusterPropertiesResponseOutputWithContext(context.Background())
}

func (i KubernetesClusterPropertiesResponseArgs) ToKubernetesClusterPropertiesResponseOutputWithContext(ctx context.Context) KubernetesClusterPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPropertiesResponseOutput)
}

func (i KubernetesClusterPropertiesResponseArgs) ToKubernetesClusterPropertiesResponsePtrOutput() KubernetesClusterPropertiesResponsePtrOutput {
	return i.ToKubernetesClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KubernetesClusterPropertiesResponseArgs) ToKubernetesClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) KubernetesClusterPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPropertiesResponseOutput).ToKubernetesClusterPropertiesResponsePtrOutputWithContext(ctx)
}









type KubernetesClusterPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKubernetesClusterPropertiesResponsePtrOutput() KubernetesClusterPropertiesResponsePtrOutput
	ToKubernetesClusterPropertiesResponsePtrOutputWithContext(context.Context) KubernetesClusterPropertiesResponsePtrOutput
}

type kubernetesClusterPropertiesResponsePtrType KubernetesClusterPropertiesResponseArgs

func KubernetesClusterPropertiesResponsePtr(v *KubernetesClusterPropertiesResponseArgs) KubernetesClusterPropertiesResponsePtrInput {
	return (*kubernetesClusterPropertiesResponsePtrType)(v)
}

func (*kubernetesClusterPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterPropertiesResponse)(nil)).Elem()
}

func (i *kubernetesClusterPropertiesResponsePtrType) ToKubernetesClusterPropertiesResponsePtrOutput() KubernetesClusterPropertiesResponsePtrOutput {
	return i.ToKubernetesClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *kubernetesClusterPropertiesResponsePtrType) ToKubernetesClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) KubernetesClusterPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterPropertiesResponsePtrOutput)
}

type KubernetesClusterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KubernetesClusterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterPropertiesResponse)(nil)).Elem()
}

func (o KubernetesClusterPropertiesResponseOutput) ToKubernetesClusterPropertiesResponseOutput() KubernetesClusterPropertiesResponseOutput {
	return o
}

func (o KubernetesClusterPropertiesResponseOutput) ToKubernetesClusterPropertiesResponseOutputWithContext(ctx context.Context) KubernetesClusterPropertiesResponseOutput {
	return o
}

func (o KubernetesClusterPropertiesResponseOutput) ToKubernetesClusterPropertiesResponsePtrOutput() KubernetesClusterPropertiesResponsePtrOutput {
	return o.ToKubernetesClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KubernetesClusterPropertiesResponseOutput) ToKubernetesClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) KubernetesClusterPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesClusterPropertiesResponse) *KubernetesClusterPropertiesResponse {
		return &v
	}).(KubernetesClusterPropertiesResponsePtrOutput)
}

func (o KubernetesClusterPropertiesResponseOutput) ServicePrincipal() ServicePrincipalPropertiesResponseOutput {
	return o.ApplyT(func(v KubernetesClusterPropertiesResponse) ServicePrincipalPropertiesResponse {
		return v.ServicePrincipal
	}).(ServicePrincipalPropertiesResponseOutput)
}

type KubernetesClusterPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesClusterPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesClusterPropertiesResponse)(nil)).Elem()
}

func (o KubernetesClusterPropertiesResponsePtrOutput) ToKubernetesClusterPropertiesResponsePtrOutput() KubernetesClusterPropertiesResponsePtrOutput {
	return o
}

func (o KubernetesClusterPropertiesResponsePtrOutput) ToKubernetesClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) KubernetesClusterPropertiesResponsePtrOutput {
	return o
}

func (o KubernetesClusterPropertiesResponsePtrOutput) Elem() KubernetesClusterPropertiesResponseOutput {
	return o.ApplyT(func(v *KubernetesClusterPropertiesResponse) KubernetesClusterPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesClusterPropertiesResponse
		return ret
	}).(KubernetesClusterPropertiesResponseOutput)
}

func (o KubernetesClusterPropertiesResponsePtrOutput) ServicePrincipal() ServicePrincipalPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *KubernetesClusterPropertiesResponse) *ServicePrincipalPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.ServicePrincipal
	}).(ServicePrincipalPropertiesResponsePtrOutput)
}

type ServiceAuthConfiguration struct {
	PrimaryAuthKeyHash   string `pulumi:"primaryAuthKeyHash"`
	SecondaryAuthKeyHash string `pulumi:"secondaryAuthKeyHash"`
}





type ServiceAuthConfigurationInput interface {
	pulumi.Input

	ToServiceAuthConfigurationOutput() ServiceAuthConfigurationOutput
	ToServiceAuthConfigurationOutputWithContext(context.Context) ServiceAuthConfigurationOutput
}

type ServiceAuthConfigurationArgs struct {
	PrimaryAuthKeyHash   pulumi.StringInput `pulumi:"primaryAuthKeyHash"`
	SecondaryAuthKeyHash pulumi.StringInput `pulumi:"secondaryAuthKeyHash"`
}

func (ServiceAuthConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthConfiguration)(nil)).Elem()
}

func (i ServiceAuthConfigurationArgs) ToServiceAuthConfigurationOutput() ServiceAuthConfigurationOutput {
	return i.ToServiceAuthConfigurationOutputWithContext(context.Background())
}

func (i ServiceAuthConfigurationArgs) ToServiceAuthConfigurationOutputWithContext(ctx context.Context) ServiceAuthConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthConfigurationOutput)
}

func (i ServiceAuthConfigurationArgs) ToServiceAuthConfigurationPtrOutput() ServiceAuthConfigurationPtrOutput {
	return i.ToServiceAuthConfigurationPtrOutputWithContext(context.Background())
}

func (i ServiceAuthConfigurationArgs) ToServiceAuthConfigurationPtrOutputWithContext(ctx context.Context) ServiceAuthConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthConfigurationOutput).ToServiceAuthConfigurationPtrOutputWithContext(ctx)
}









type ServiceAuthConfigurationPtrInput interface {
	pulumi.Input

	ToServiceAuthConfigurationPtrOutput() ServiceAuthConfigurationPtrOutput
	ToServiceAuthConfigurationPtrOutputWithContext(context.Context) ServiceAuthConfigurationPtrOutput
}

type serviceAuthConfigurationPtrType ServiceAuthConfigurationArgs

func ServiceAuthConfigurationPtr(v *ServiceAuthConfigurationArgs) ServiceAuthConfigurationPtrInput {
	return (*serviceAuthConfigurationPtrType)(v)
}

func (*serviceAuthConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthConfiguration)(nil)).Elem()
}

func (i *serviceAuthConfigurationPtrType) ToServiceAuthConfigurationPtrOutput() ServiceAuthConfigurationPtrOutput {
	return i.ToServiceAuthConfigurationPtrOutputWithContext(context.Background())
}

func (i *serviceAuthConfigurationPtrType) ToServiceAuthConfigurationPtrOutputWithContext(ctx context.Context) ServiceAuthConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthConfigurationPtrOutput)
}

type ServiceAuthConfigurationOutput struct{ *pulumi.OutputState }

func (ServiceAuthConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthConfiguration)(nil)).Elem()
}

func (o ServiceAuthConfigurationOutput) ToServiceAuthConfigurationOutput() ServiceAuthConfigurationOutput {
	return o
}

func (o ServiceAuthConfigurationOutput) ToServiceAuthConfigurationOutputWithContext(ctx context.Context) ServiceAuthConfigurationOutput {
	return o
}

func (o ServiceAuthConfigurationOutput) ToServiceAuthConfigurationPtrOutput() ServiceAuthConfigurationPtrOutput {
	return o.ToServiceAuthConfigurationPtrOutputWithContext(context.Background())
}

func (o ServiceAuthConfigurationOutput) ToServiceAuthConfigurationPtrOutputWithContext(ctx context.Context) ServiceAuthConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceAuthConfiguration) *ServiceAuthConfiguration {
		return &v
	}).(ServiceAuthConfigurationPtrOutput)
}

func (o ServiceAuthConfigurationOutput) PrimaryAuthKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAuthConfiguration) string { return v.PrimaryAuthKeyHash }).(pulumi.StringOutput)
}

func (o ServiceAuthConfigurationOutput) SecondaryAuthKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAuthConfiguration) string { return v.SecondaryAuthKeyHash }).(pulumi.StringOutput)
}

type ServiceAuthConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthConfiguration)(nil)).Elem()
}

func (o ServiceAuthConfigurationPtrOutput) ToServiceAuthConfigurationPtrOutput() ServiceAuthConfigurationPtrOutput {
	return o
}

func (o ServiceAuthConfigurationPtrOutput) ToServiceAuthConfigurationPtrOutputWithContext(ctx context.Context) ServiceAuthConfigurationPtrOutput {
	return o
}

func (o ServiceAuthConfigurationPtrOutput) Elem() ServiceAuthConfigurationOutput {
	return o.ApplyT(func(v *ServiceAuthConfiguration) ServiceAuthConfiguration {
		if v != nil {
			return *v
		}
		var ret ServiceAuthConfiguration
		return ret
	}).(ServiceAuthConfigurationOutput)
}

func (o ServiceAuthConfigurationPtrOutput) PrimaryAuthKeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryAuthKeyHash
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthConfigurationPtrOutput) SecondaryAuthKeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.SecondaryAuthKeyHash
	}).(pulumi.StringPtrOutput)
}

type ServiceAuthConfigurationResponse struct {
	PrimaryAuthKeyHash   string `pulumi:"primaryAuthKeyHash"`
	SecondaryAuthKeyHash string `pulumi:"secondaryAuthKeyHash"`
}





type ServiceAuthConfigurationResponseInput interface {
	pulumi.Input

	ToServiceAuthConfigurationResponseOutput() ServiceAuthConfigurationResponseOutput
	ToServiceAuthConfigurationResponseOutputWithContext(context.Context) ServiceAuthConfigurationResponseOutput
}

type ServiceAuthConfigurationResponseArgs struct {
	PrimaryAuthKeyHash   pulumi.StringInput `pulumi:"primaryAuthKeyHash"`
	SecondaryAuthKeyHash pulumi.StringInput `pulumi:"secondaryAuthKeyHash"`
}

func (ServiceAuthConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthConfigurationResponse)(nil)).Elem()
}

func (i ServiceAuthConfigurationResponseArgs) ToServiceAuthConfigurationResponseOutput() ServiceAuthConfigurationResponseOutput {
	return i.ToServiceAuthConfigurationResponseOutputWithContext(context.Background())
}

func (i ServiceAuthConfigurationResponseArgs) ToServiceAuthConfigurationResponseOutputWithContext(ctx context.Context) ServiceAuthConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthConfigurationResponseOutput)
}

func (i ServiceAuthConfigurationResponseArgs) ToServiceAuthConfigurationResponsePtrOutput() ServiceAuthConfigurationResponsePtrOutput {
	return i.ToServiceAuthConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i ServiceAuthConfigurationResponseArgs) ToServiceAuthConfigurationResponsePtrOutputWithContext(ctx context.Context) ServiceAuthConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthConfigurationResponseOutput).ToServiceAuthConfigurationResponsePtrOutputWithContext(ctx)
}









type ServiceAuthConfigurationResponsePtrInput interface {
	pulumi.Input

	ToServiceAuthConfigurationResponsePtrOutput() ServiceAuthConfigurationResponsePtrOutput
	ToServiceAuthConfigurationResponsePtrOutputWithContext(context.Context) ServiceAuthConfigurationResponsePtrOutput
}

type serviceAuthConfigurationResponsePtrType ServiceAuthConfigurationResponseArgs

func ServiceAuthConfigurationResponsePtr(v *ServiceAuthConfigurationResponseArgs) ServiceAuthConfigurationResponsePtrInput {
	return (*serviceAuthConfigurationResponsePtrType)(v)
}

func (*serviceAuthConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthConfigurationResponse)(nil)).Elem()
}

func (i *serviceAuthConfigurationResponsePtrType) ToServiceAuthConfigurationResponsePtrOutput() ServiceAuthConfigurationResponsePtrOutput {
	return i.ToServiceAuthConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *serviceAuthConfigurationResponsePtrType) ToServiceAuthConfigurationResponsePtrOutputWithContext(ctx context.Context) ServiceAuthConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthConfigurationResponsePtrOutput)
}

type ServiceAuthConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ServiceAuthConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthConfigurationResponse)(nil)).Elem()
}

func (o ServiceAuthConfigurationResponseOutput) ToServiceAuthConfigurationResponseOutput() ServiceAuthConfigurationResponseOutput {
	return o
}

func (o ServiceAuthConfigurationResponseOutput) ToServiceAuthConfigurationResponseOutputWithContext(ctx context.Context) ServiceAuthConfigurationResponseOutput {
	return o
}

func (o ServiceAuthConfigurationResponseOutput) ToServiceAuthConfigurationResponsePtrOutput() ServiceAuthConfigurationResponsePtrOutput {
	return o.ToServiceAuthConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o ServiceAuthConfigurationResponseOutput) ToServiceAuthConfigurationResponsePtrOutputWithContext(ctx context.Context) ServiceAuthConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceAuthConfigurationResponse) *ServiceAuthConfigurationResponse {
		return &v
	}).(ServiceAuthConfigurationResponsePtrOutput)
}

func (o ServiceAuthConfigurationResponseOutput) PrimaryAuthKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAuthConfigurationResponse) string { return v.PrimaryAuthKeyHash }).(pulumi.StringOutput)
}

func (o ServiceAuthConfigurationResponseOutput) SecondaryAuthKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAuthConfigurationResponse) string { return v.SecondaryAuthKeyHash }).(pulumi.StringOutput)
}

type ServiceAuthConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthConfigurationResponse)(nil)).Elem()
}

func (o ServiceAuthConfigurationResponsePtrOutput) ToServiceAuthConfigurationResponsePtrOutput() ServiceAuthConfigurationResponsePtrOutput {
	return o
}

func (o ServiceAuthConfigurationResponsePtrOutput) ToServiceAuthConfigurationResponsePtrOutputWithContext(ctx context.Context) ServiceAuthConfigurationResponsePtrOutput {
	return o
}

func (o ServiceAuthConfigurationResponsePtrOutput) Elem() ServiceAuthConfigurationResponseOutput {
	return o.ApplyT(func(v *ServiceAuthConfigurationResponse) ServiceAuthConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ServiceAuthConfigurationResponse
		return ret
	}).(ServiceAuthConfigurationResponseOutput)
}

func (o ServiceAuthConfigurationResponsePtrOutput) PrimaryAuthKeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryAuthKeyHash
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthConfigurationResponsePtrOutput) SecondaryAuthKeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondaryAuthKeyHash
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalProperties struct {
	ClientId string `pulumi:"clientId"`
	Secret   string `pulumi:"secret"`
}





type ServicePrincipalPropertiesInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput
	ToServicePrincipalPropertiesOutputWithContext(context.Context) ServicePrincipalPropertiesOutput
}

type ServicePrincipalPropertiesArgs struct {
	ClientId pulumi.StringInput `pulumi:"clientId"`
	Secret   pulumi.StringInput `pulumi:"secret"`
}

func (ServicePrincipalPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProperties)(nil)).Elem()
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput {
	return i.ToServicePrincipalPropertiesOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesOutputWithContext(ctx context.Context) ServicePrincipalPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesOutput)
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return i.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesOutput).ToServicePrincipalPropertiesPtrOutputWithContext(ctx)
}









type ServicePrincipalPropertiesPtrInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput
	ToServicePrincipalPropertiesPtrOutputWithContext(context.Context) ServicePrincipalPropertiesPtrOutput
}

type servicePrincipalPropertiesPtrType ServicePrincipalPropertiesArgs

func ServicePrincipalPropertiesPtr(v *ServicePrincipalPropertiesArgs) ServicePrincipalPropertiesPtrInput {
	return (*servicePrincipalPropertiesPtrType)(v)
}

func (*servicePrincipalPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProperties)(nil)).Elem()
}

func (i *servicePrincipalPropertiesPtrType) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return i.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (i *servicePrincipalPropertiesPtrType) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesPtrOutput)
}

type ServicePrincipalPropertiesOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProperties)(nil)).Elem()
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput {
	return o
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesOutputWithContext(ctx context.Context) ServicePrincipalPropertiesOutput {
	return o
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return o.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipalProperties) *ServicePrincipalProperties {
		return &v
	}).(ServicePrincipalPropertiesPtrOutput)
}

func (o ServicePrincipalPropertiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalProperties) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ServicePrincipalPropertiesOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalProperties) string { return v.Secret }).(pulumi.StringOutput)
}

type ServicePrincipalPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProperties)(nil)).Elem()
}

func (o ServicePrincipalPropertiesPtrOutput) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return o
}

func (o ServicePrincipalPropertiesPtrOutput) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return o
}

func (o ServicePrincipalPropertiesPtrOutput) Elem() ServicePrincipalPropertiesOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) ServicePrincipalProperties {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalProperties
		return ret
	}).(ServicePrincipalPropertiesOutput)
}

func (o ServicePrincipalPropertiesPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalPropertiesResponse struct {
	ClientId string `pulumi:"clientId"`
	Secret   string `pulumi:"secret"`
}





type ServicePrincipalPropertiesResponseInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesResponseOutput() ServicePrincipalPropertiesResponseOutput
	ToServicePrincipalPropertiesResponseOutputWithContext(context.Context) ServicePrincipalPropertiesResponseOutput
}

type ServicePrincipalPropertiesResponseArgs struct {
	ClientId pulumi.StringInput `pulumi:"clientId"`
	Secret   pulumi.StringInput `pulumi:"secret"`
}

func (ServicePrincipalPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (i ServicePrincipalPropertiesResponseArgs) ToServicePrincipalPropertiesResponseOutput() ServicePrincipalPropertiesResponseOutput {
	return i.ToServicePrincipalPropertiesResponseOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesResponseArgs) ToServicePrincipalPropertiesResponseOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesResponseOutput)
}

func (i ServicePrincipalPropertiesResponseArgs) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return i.ToServicePrincipalPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesResponseArgs) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesResponseOutput).ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx)
}









type ServicePrincipalPropertiesResponsePtrInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput
	ToServicePrincipalPropertiesResponsePtrOutputWithContext(context.Context) ServicePrincipalPropertiesResponsePtrOutput
}

type servicePrincipalPropertiesResponsePtrType ServicePrincipalPropertiesResponseArgs

func ServicePrincipalPropertiesResponsePtr(v *ServicePrincipalPropertiesResponseArgs) ServicePrincipalPropertiesResponsePtrInput {
	return (*servicePrincipalPropertiesResponsePtrType)(v)
}

func (*servicePrincipalPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (i *servicePrincipalPropertiesResponsePtrType) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return i.ToServicePrincipalPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *servicePrincipalPropertiesResponsePtrType) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesResponsePtrOutput)
}

type ServicePrincipalPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponseOutput() ServicePrincipalPropertiesResponseOutput {
	return o
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponseOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponseOutput {
	return o
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return o.ToServicePrincipalPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipalPropertiesResponse) *ServicePrincipalPropertiesResponse {
		return &v
	}).(ServicePrincipalPropertiesResponsePtrOutput)
}

func (o ServicePrincipalPropertiesResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalPropertiesResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ServicePrincipalPropertiesResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePrincipalPropertiesResponse) string { return v.Secret }).(pulumi.StringOutput)
}

type ServicePrincipalPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return o
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return o
}

func (o ServicePrincipalPropertiesResponsePtrOutput) Elem() ServicePrincipalPropertiesResponseOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) ServicePrincipalPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalPropertiesResponse
		return ret
	}).(ServicePrincipalPropertiesResponseOutput)
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SslConfiguration struct {
	Cert   *string `pulumi:"cert"`
	Key    *string `pulumi:"key"`
	Status *string `pulumi:"status"`
}





type SslConfigurationInput interface {
	pulumi.Input

	ToSslConfigurationOutput() SslConfigurationOutput
	ToSslConfigurationOutputWithContext(context.Context) SslConfigurationOutput
}

type SslConfigurationArgs struct {
	Cert   pulumi.StringPtrInput `pulumi:"cert"`
	Key    pulumi.StringPtrInput `pulumi:"key"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (SslConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfiguration)(nil)).Elem()
}

func (i SslConfigurationArgs) ToSslConfigurationOutput() SslConfigurationOutput {
	return i.ToSslConfigurationOutputWithContext(context.Background())
}

func (i SslConfigurationArgs) ToSslConfigurationOutputWithContext(ctx context.Context) SslConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationOutput)
}

func (i SslConfigurationArgs) ToSslConfigurationPtrOutput() SslConfigurationPtrOutput {
	return i.ToSslConfigurationPtrOutputWithContext(context.Background())
}

func (i SslConfigurationArgs) ToSslConfigurationPtrOutputWithContext(ctx context.Context) SslConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationOutput).ToSslConfigurationPtrOutputWithContext(ctx)
}









type SslConfigurationPtrInput interface {
	pulumi.Input

	ToSslConfigurationPtrOutput() SslConfigurationPtrOutput
	ToSslConfigurationPtrOutputWithContext(context.Context) SslConfigurationPtrOutput
}

type sslConfigurationPtrType SslConfigurationArgs

func SslConfigurationPtr(v *SslConfigurationArgs) SslConfigurationPtrInput {
	return (*sslConfigurationPtrType)(v)
}

func (*sslConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfiguration)(nil)).Elem()
}

func (i *sslConfigurationPtrType) ToSslConfigurationPtrOutput() SslConfigurationPtrOutput {
	return i.ToSslConfigurationPtrOutputWithContext(context.Background())
}

func (i *sslConfigurationPtrType) ToSslConfigurationPtrOutputWithContext(ctx context.Context) SslConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationPtrOutput)
}

type SslConfigurationOutput struct{ *pulumi.OutputState }

func (SslConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfiguration)(nil)).Elem()
}

func (o SslConfigurationOutput) ToSslConfigurationOutput() SslConfigurationOutput {
	return o
}

func (o SslConfigurationOutput) ToSslConfigurationOutputWithContext(ctx context.Context) SslConfigurationOutput {
	return o
}

func (o SslConfigurationOutput) ToSslConfigurationPtrOutput() SslConfigurationPtrOutput {
	return o.ToSslConfigurationPtrOutputWithContext(context.Background())
}

func (o SslConfigurationOutput) ToSslConfigurationPtrOutputWithContext(ctx context.Context) SslConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslConfiguration) *SslConfiguration {
		return &v
	}).(SslConfigurationPtrOutput)
}

func (o SslConfigurationOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.Cert }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SslConfigurationPtrOutput struct{ *pulumi.OutputState }

func (SslConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfiguration)(nil)).Elem()
}

func (o SslConfigurationPtrOutput) ToSslConfigurationPtrOutput() SslConfigurationPtrOutput {
	return o
}

func (o SslConfigurationPtrOutput) ToSslConfigurationPtrOutputWithContext(ctx context.Context) SslConfigurationPtrOutput {
	return o
}

func (o SslConfigurationPtrOutput) Elem() SslConfigurationOutput {
	return o.ApplyT(func(v *SslConfiguration) SslConfiguration {
		if v != nil {
			return *v
		}
		var ret SslConfiguration
		return ret
	}).(SslConfigurationOutput)
}

func (o SslConfigurationPtrOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Cert
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type SslConfigurationResponse struct {
	Cert   *string `pulumi:"cert"`
	Key    *string `pulumi:"key"`
	Status *string `pulumi:"status"`
}





type SslConfigurationResponseInput interface {
	pulumi.Input

	ToSslConfigurationResponseOutput() SslConfigurationResponseOutput
	ToSslConfigurationResponseOutputWithContext(context.Context) SslConfigurationResponseOutput
}

type SslConfigurationResponseArgs struct {
	Cert   pulumi.StringPtrInput `pulumi:"cert"`
	Key    pulumi.StringPtrInput `pulumi:"key"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (SslConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigurationResponse)(nil)).Elem()
}

func (i SslConfigurationResponseArgs) ToSslConfigurationResponseOutput() SslConfigurationResponseOutput {
	return i.ToSslConfigurationResponseOutputWithContext(context.Background())
}

func (i SslConfigurationResponseArgs) ToSslConfigurationResponseOutputWithContext(ctx context.Context) SslConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationResponseOutput)
}

func (i SslConfigurationResponseArgs) ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput {
	return i.ToSslConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i SslConfigurationResponseArgs) ToSslConfigurationResponsePtrOutputWithContext(ctx context.Context) SslConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationResponseOutput).ToSslConfigurationResponsePtrOutputWithContext(ctx)
}









type SslConfigurationResponsePtrInput interface {
	pulumi.Input

	ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput
	ToSslConfigurationResponsePtrOutputWithContext(context.Context) SslConfigurationResponsePtrOutput
}

type sslConfigurationResponsePtrType SslConfigurationResponseArgs

func SslConfigurationResponsePtr(v *SslConfigurationResponseArgs) SslConfigurationResponsePtrInput {
	return (*sslConfigurationResponsePtrType)(v)
}

func (*sslConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfigurationResponse)(nil)).Elem()
}

func (i *sslConfigurationResponsePtrType) ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput {
	return i.ToSslConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *sslConfigurationResponsePtrType) ToSslConfigurationResponsePtrOutputWithContext(ctx context.Context) SslConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslConfigurationResponsePtrOutput)
}

type SslConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SslConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigurationResponse)(nil)).Elem()
}

func (o SslConfigurationResponseOutput) ToSslConfigurationResponseOutput() SslConfigurationResponseOutput {
	return o
}

func (o SslConfigurationResponseOutput) ToSslConfigurationResponseOutputWithContext(ctx context.Context) SslConfigurationResponseOutput {
	return o
}

func (o SslConfigurationResponseOutput) ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput {
	return o.ToSslConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o SslConfigurationResponseOutput) ToSslConfigurationResponsePtrOutputWithContext(ctx context.Context) SslConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslConfigurationResponse) *SslConfigurationResponse {
		return &v
	}).(SslConfigurationResponsePtrOutput)
}

func (o SslConfigurationResponseOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.Cert }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SslConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (SslConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfigurationResponse)(nil)).Elem()
}

func (o SslConfigurationResponsePtrOutput) ToSslConfigurationResponsePtrOutput() SslConfigurationResponsePtrOutput {
	return o
}

func (o SslConfigurationResponsePtrOutput) ToSslConfigurationResponsePtrOutputWithContext(ctx context.Context) SslConfigurationResponsePtrOutput {
	return o
}

func (o SslConfigurationResponsePtrOutput) Elem() SslConfigurationResponseOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) SslConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret SslConfigurationResponse
		return ret
	}).(SslConfigurationResponseOutput)
}

func (o SslConfigurationResponsePtrOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cert
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponsePtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type StorageAccountCredentialsResponse struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	ResourceId   string `pulumi:"resourceId"`
	SecondaryKey string `pulumi:"secondaryKey"`
}





type StorageAccountCredentialsResponseInput interface {
	pulumi.Input

	ToStorageAccountCredentialsResponseOutput() StorageAccountCredentialsResponseOutput
	ToStorageAccountCredentialsResponseOutputWithContext(context.Context) StorageAccountCredentialsResponseOutput
}

type StorageAccountCredentialsResponseArgs struct {
	PrimaryKey   pulumi.StringInput `pulumi:"primaryKey"`
	ResourceId   pulumi.StringInput `pulumi:"resourceId"`
	SecondaryKey pulumi.StringInput `pulumi:"secondaryKey"`
}

func (StorageAccountCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountCredentialsResponse)(nil)).Elem()
}

func (i StorageAccountCredentialsResponseArgs) ToStorageAccountCredentialsResponseOutput() StorageAccountCredentialsResponseOutput {
	return i.ToStorageAccountCredentialsResponseOutputWithContext(context.Background())
}

func (i StorageAccountCredentialsResponseArgs) ToStorageAccountCredentialsResponseOutputWithContext(ctx context.Context) StorageAccountCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountCredentialsResponseOutput)
}

func (i StorageAccountCredentialsResponseArgs) ToStorageAccountCredentialsResponsePtrOutput() StorageAccountCredentialsResponsePtrOutput {
	return i.ToStorageAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i StorageAccountCredentialsResponseArgs) ToStorageAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) StorageAccountCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountCredentialsResponseOutput).ToStorageAccountCredentialsResponsePtrOutputWithContext(ctx)
}









type StorageAccountCredentialsResponsePtrInput interface {
	pulumi.Input

	ToStorageAccountCredentialsResponsePtrOutput() StorageAccountCredentialsResponsePtrOutput
	ToStorageAccountCredentialsResponsePtrOutputWithContext(context.Context) StorageAccountCredentialsResponsePtrOutput
}

type storageAccountCredentialsResponsePtrType StorageAccountCredentialsResponseArgs

func StorageAccountCredentialsResponsePtr(v *StorageAccountCredentialsResponseArgs) StorageAccountCredentialsResponsePtrInput {
	return (*storageAccountCredentialsResponsePtrType)(v)
}

func (*storageAccountCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountCredentialsResponse)(nil)).Elem()
}

func (i *storageAccountCredentialsResponsePtrType) ToStorageAccountCredentialsResponsePtrOutput() StorageAccountCredentialsResponsePtrOutput {
	return i.ToStorageAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *storageAccountCredentialsResponsePtrType) ToStorageAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) StorageAccountCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountCredentialsResponsePtrOutput)
}

type StorageAccountCredentialsResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountCredentialsResponse)(nil)).Elem()
}

func (o StorageAccountCredentialsResponseOutput) ToStorageAccountCredentialsResponseOutput() StorageAccountCredentialsResponseOutput {
	return o
}

func (o StorageAccountCredentialsResponseOutput) ToStorageAccountCredentialsResponseOutputWithContext(ctx context.Context) StorageAccountCredentialsResponseOutput {
	return o
}

func (o StorageAccountCredentialsResponseOutput) ToStorageAccountCredentialsResponsePtrOutput() StorageAccountCredentialsResponsePtrOutput {
	return o.ToStorageAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountCredentialsResponseOutput) ToStorageAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) StorageAccountCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountCredentialsResponse) *StorageAccountCredentialsResponse {
		return &v
	}).(StorageAccountCredentialsResponsePtrOutput)
}

func (o StorageAccountCredentialsResponseOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountCredentialsResponse) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o StorageAccountCredentialsResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountCredentialsResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o StorageAccountCredentialsResponseOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountCredentialsResponse) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

type StorageAccountCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountCredentialsResponse)(nil)).Elem()
}

func (o StorageAccountCredentialsResponsePtrOutput) ToStorageAccountCredentialsResponsePtrOutput() StorageAccountCredentialsResponsePtrOutput {
	return o
}

func (o StorageAccountCredentialsResponsePtrOutput) ToStorageAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) StorageAccountCredentialsResponsePtrOutput {
	return o
}

func (o StorageAccountCredentialsResponsePtrOutput) Elem() StorageAccountCredentialsResponseOutput {
	return o.ApplyT(func(v *StorageAccountCredentialsResponse) StorageAccountCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountCredentialsResponse
		return ret
	}).(StorageAccountCredentialsResponseOutput)
}

func (o StorageAccountCredentialsResponsePtrOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountCredentialsResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountCredentialsResponsePtrOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

type StorageAccountProperties struct {
	ResourceId *string `pulumi:"resourceId"`
}





type StorageAccountPropertiesInput interface {
	pulumi.Input

	ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput
	ToStorageAccountPropertiesOutputWithContext(context.Context) StorageAccountPropertiesOutput
}

type StorageAccountPropertiesArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (StorageAccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return i.ToStorageAccountPropertiesOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput)
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput).ToStorageAccountPropertiesPtrOutputWithContext(ctx)
}









type StorageAccountPropertiesPtrInput interface {
	pulumi.Input

	ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput
	ToStorageAccountPropertiesPtrOutputWithContext(context.Context) StorageAccountPropertiesPtrOutput
}

type storageAccountPropertiesPtrType StorageAccountPropertiesArgs

func StorageAccountPropertiesPtr(v *StorageAccountPropertiesArgs) StorageAccountPropertiesPtrInput {
	return (*storageAccountPropertiesPtrType)(v)
}

func (*storageAccountPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesPtrOutput)
}

type StorageAccountPropertiesOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountProperties) *StorageAccountProperties {
		return &v
	}).(StorageAccountPropertiesPtrOutput)
}

func (o StorageAccountPropertiesOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountProperties) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) Elem() StorageAccountPropertiesOutput {
	return o.ApplyT(func(v *StorageAccountProperties) StorageAccountProperties {
		if v != nil {
			return *v
		}
		var ret StorageAccountProperties
		return ret
	}).(StorageAccountPropertiesOutput)
}

func (o StorageAccountPropertiesPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponse struct {
	ResourceId *string `pulumi:"resourceId"`
}





type StorageAccountPropertiesResponseInput interface {
	pulumi.Input

	ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput
	ToStorageAccountPropertiesResponseOutputWithContext(context.Context) StorageAccountPropertiesResponseOutput
}

type StorageAccountPropertiesResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (StorageAccountPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return i.ToStorageAccountPropertiesResponseOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponseOutput)
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return i.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponseOutput).ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx)
}









type StorageAccountPropertiesResponsePtrInput interface {
	pulumi.Input

	ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput
	ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Context) StorageAccountPropertiesResponsePtrOutput
}

type storageAccountPropertiesResponsePtrType StorageAccountPropertiesResponseArgs

func StorageAccountPropertiesResponsePtr(v *StorageAccountPropertiesResponseArgs) StorageAccountPropertiesResponsePtrInput {
	return (*storageAccountPropertiesResponsePtrType)(v)
}

func (*storageAccountPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (i *storageAccountPropertiesResponsePtrType) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return i.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *storageAccountPropertiesResponsePtrType) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponsePtrOutput)
}

type StorageAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountPropertiesResponse) *StorageAccountPropertiesResponse {
		return &v
	}).(StorageAccountPropertiesResponsePtrOutput)
}

func (o StorageAccountPropertiesResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) Elem() StorageAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) StorageAccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountPropertiesResponse
		return ret
	}).(StorageAccountPropertiesResponseOutput)
}

func (o StorageAccountPropertiesResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AcsClusterPropertiesOutput{})
	pulumi.RegisterOutputType(AcsClusterPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AcsClusterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AcsClusterPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AppInsightsCredentialsOutput{})
	pulumi.RegisterOutputType(AppInsightsCredentialsPtrOutput{})
	pulumi.RegisterOutputType(AppInsightsCredentialsResponseOutput{})
	pulumi.RegisterOutputType(AppInsightsCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoScaleConfigurationOutput{})
	pulumi.RegisterOutputType(AutoScaleConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AutoScaleConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AutoScaleConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerRegistryCredentialsResponseOutput{})
	pulumi.RegisterOutputType(ContainerRegistryCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerRegistryPropertiesOutput{})
	pulumi.RegisterOutputType(ContainerRegistryPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ContainerRegistryPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ContainerRegistryPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceCredentialsResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(GlobalServiceConfigurationOutput{})
	pulumi.RegisterOutputType(GlobalServiceConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GlobalServiceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GlobalServiceConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesClusterPropertiesOutput{})
	pulumi.RegisterOutputType(KubernetesClusterPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KubernetesClusterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KubernetesClusterPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceAuthConfigurationOutput{})
	pulumi.RegisterOutputType(ServiceAuthConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ServiceAuthConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ServiceAuthConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SslConfigurationOutput{})
	pulumi.RegisterOutputType(SslConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SslConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SslConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountCredentialsResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponsePtrOutput{})
}
