


package v20200101

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





type AKSInput interface {
	pulumi.Input

	ToAKSOutput() AKSOutput
	ToAKSOutputWithContext(context.Context) AKSOutput
}

type AKSArgs struct {
	ComputeLocation pulumi.StringPtrInput `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput    `pulumi:"computeType"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Properties      AKSPropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (AKSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKS)(nil)).Elem()
}

func (i AKSArgs) ToAKSOutput() AKSOutput {
	return i.ToAKSOutputWithContext(context.Background())
}

func (i AKSArgs) ToAKSOutputWithContext(ctx context.Context) AKSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSOutput)
}

type AKSOutput struct{ *pulumi.OutputState }

func (AKSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKS)(nil)).Elem()
}

func (o AKSOutput) ToAKSOutput() AKSOutput {
	return o
}

func (o AKSOutput) ToAKSOutputWithContext(ctx context.Context) AKSOutput {
	return o
}

func (o AKSOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKS) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o AKSOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AKS) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AKSOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKS) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AKSOutput) Properties() AKSPropertiesPtrOutput {
	return o.ApplyT(func(v AKS) *AKSProperties { return v.Properties }).(AKSPropertiesPtrOutput)
}

func (o AKSOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKS) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AKSProperties struct {
	AgentCount                 *int                        `pulumi:"agentCount"`
	AgentVMSize                *string                     `pulumi:"agentVMSize"`
	AksNetworkingConfiguration *AksNetworkingConfiguration `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                     `pulumi:"clusterFqdn"`
	SslConfiguration           *SslConfiguration           `pulumi:"sslConfiguration"`
}





type AKSPropertiesInput interface {
	pulumi.Input

	ToAKSPropertiesOutput() AKSPropertiesOutput
	ToAKSPropertiesOutputWithContext(context.Context) AKSPropertiesOutput
}

type AKSPropertiesArgs struct {
	AgentCount                 pulumi.IntPtrInput                 `pulumi:"agentCount"`
	AgentVMSize                pulumi.StringPtrInput              `pulumi:"agentVMSize"`
	AksNetworkingConfiguration AksNetworkingConfigurationPtrInput `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                pulumi.StringPtrInput              `pulumi:"clusterFqdn"`
	SslConfiguration           SslConfigurationPtrInput           `pulumi:"sslConfiguration"`
}

func (AKSPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSProperties)(nil)).Elem()
}

func (i AKSPropertiesArgs) ToAKSPropertiesOutput() AKSPropertiesOutput {
	return i.ToAKSPropertiesOutputWithContext(context.Background())
}

func (i AKSPropertiesArgs) ToAKSPropertiesOutputWithContext(ctx context.Context) AKSPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSPropertiesOutput)
}

func (i AKSPropertiesArgs) ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput {
	return i.ToAKSPropertiesPtrOutputWithContext(context.Background())
}

func (i AKSPropertiesArgs) ToAKSPropertiesPtrOutputWithContext(ctx context.Context) AKSPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSPropertiesOutput).ToAKSPropertiesPtrOutputWithContext(ctx)
}









type AKSPropertiesPtrInput interface {
	pulumi.Input

	ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput
	ToAKSPropertiesPtrOutputWithContext(context.Context) AKSPropertiesPtrOutput
}

type akspropertiesPtrType AKSPropertiesArgs

func AKSPropertiesPtr(v *AKSPropertiesArgs) AKSPropertiesPtrInput {
	return (*akspropertiesPtrType)(v)
}

func (*akspropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSProperties)(nil)).Elem()
}

func (i *akspropertiesPtrType) ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput {
	return i.ToAKSPropertiesPtrOutputWithContext(context.Background())
}

func (i *akspropertiesPtrType) ToAKSPropertiesPtrOutputWithContext(ctx context.Context) AKSPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSPropertiesPtrOutput)
}

type AKSPropertiesOutput struct{ *pulumi.OutputState }

func (AKSPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSProperties)(nil)).Elem()
}

func (o AKSPropertiesOutput) ToAKSPropertiesOutput() AKSPropertiesOutput {
	return o
}

func (o AKSPropertiesOutput) ToAKSPropertiesOutputWithContext(ctx context.Context) AKSPropertiesOutput {
	return o
}

func (o AKSPropertiesOutput) ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput {
	return o.ToAKSPropertiesPtrOutputWithContext(context.Background())
}

func (o AKSPropertiesOutput) ToAKSPropertiesPtrOutputWithContext(ctx context.Context) AKSPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSProperties) *AKSProperties {
		return &v
	}).(AKSPropertiesPtrOutput)
}

func (o AKSPropertiesOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSProperties) *int { return v.AgentCount }).(pulumi.IntPtrOutput)
}

func (o AKSPropertiesOutput) AgentVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.AgentVMSize }).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesOutput) AksNetworkingConfiguration() AksNetworkingConfigurationPtrOutput {
	return o.ApplyT(func(v AKSProperties) *AksNetworkingConfiguration { return v.AksNetworkingConfiguration }).(AksNetworkingConfigurationPtrOutput)
}

func (o AKSPropertiesOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.ClusterFqdn }).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesOutput) SslConfiguration() SslConfigurationPtrOutput {
	return o.ApplyT(func(v AKSProperties) *SslConfiguration { return v.SslConfiguration }).(SslConfigurationPtrOutput)
}

type AKSPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AKSPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSProperties)(nil)).Elem()
}

func (o AKSPropertiesPtrOutput) ToAKSPropertiesPtrOutput() AKSPropertiesPtrOutput {
	return o
}

func (o AKSPropertiesPtrOutput) ToAKSPropertiesPtrOutputWithContext(ctx context.Context) AKSPropertiesPtrOutput {
	return o
}

func (o AKSPropertiesPtrOutput) Elem() AKSPropertiesOutput {
	return o.ApplyT(func(v *AKSProperties) AKSProperties {
		if v != nil {
			return *v
		}
		var ret AKSProperties
		return ret
	}).(AKSPropertiesOutput)
}

func (o AKSPropertiesPtrOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *int {
		if v == nil {
			return nil
		}
		return v.AgentCount
	}).(pulumi.IntPtrOutput)
}

func (o AKSPropertiesPtrOutput) AgentVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *string {
		if v == nil {
			return nil
		}
		return v.AgentVMSize
	}).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesPtrOutput) AksNetworkingConfiguration() AksNetworkingConfigurationPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *AksNetworkingConfiguration {
		if v == nil {
			return nil
		}
		return v.AksNetworkingConfiguration
	}).(AksNetworkingConfigurationPtrOutput)
}

func (o AKSPropertiesPtrOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterFqdn
	}).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesPtrOutput) SslConfiguration() SslConfigurationPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *SslConfiguration {
		if v == nil {
			return nil
		}
		return v.SslConfiguration
	}).(SslConfigurationPtrOutput)
}

type AKSResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *AKSResponseProperties                `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}





type AKSResponseInput interface {
	pulumi.Input

	ToAKSResponseOutput() AKSResponseOutput
	ToAKSResponseOutputWithContext(context.Context) AKSResponseOutput
}

type AKSResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         AKSResponsePropertiesPtrInput                 `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (AKSResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSResponse)(nil)).Elem()
}

func (i AKSResponseArgs) ToAKSResponseOutput() AKSResponseOutput {
	return i.ToAKSResponseOutputWithContext(context.Background())
}

func (i AKSResponseArgs) ToAKSResponseOutputWithContext(ctx context.Context) AKSResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSResponseOutput)
}

type AKSResponseOutput struct{ *pulumi.OutputState }

func (AKSResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSResponse)(nil)).Elem()
}

func (o AKSResponseOutput) ToAKSResponseOutput() AKSResponseOutput {
	return o
}

func (o AKSResponseOutput) ToAKSResponseOutputWithContext(ctx context.Context) AKSResponseOutput {
	return o
}

func (o AKSResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o AKSResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AKSResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v AKSResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o AKSResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) Properties() AKSResponsePropertiesPtrOutput {
	return o.ApplyT(func(v AKSResponse) *AKSResponseProperties { return v.Properties }).(AKSResponsePropertiesPtrOutput)
}

func (o AKSResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v AKSResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o AKSResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AKSResponseProperties struct {
	AgentCount                 *int                                `pulumi:"agentCount"`
	AgentVMSize                *string                             `pulumi:"agentVMSize"`
	AksNetworkingConfiguration *AksNetworkingConfigurationResponse `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                             `pulumi:"clusterFqdn"`
	SslConfiguration           *SslConfigurationResponse           `pulumi:"sslConfiguration"`
	SystemServices             []SystemServiceResponse             `pulumi:"systemServices"`
}





type AKSResponsePropertiesInput interface {
	pulumi.Input

	ToAKSResponsePropertiesOutput() AKSResponsePropertiesOutput
	ToAKSResponsePropertiesOutputWithContext(context.Context) AKSResponsePropertiesOutput
}

type AKSResponsePropertiesArgs struct {
	AgentCount                 pulumi.IntPtrInput                         `pulumi:"agentCount"`
	AgentVMSize                pulumi.StringPtrInput                      `pulumi:"agentVMSize"`
	AksNetworkingConfiguration AksNetworkingConfigurationResponsePtrInput `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                pulumi.StringPtrInput                      `pulumi:"clusterFqdn"`
	SslConfiguration           SslConfigurationResponsePtrInput           `pulumi:"sslConfiguration"`
	SystemServices             SystemServiceResponseArrayInput            `pulumi:"systemServices"`
}

func (AKSResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSResponseProperties)(nil)).Elem()
}

func (i AKSResponsePropertiesArgs) ToAKSResponsePropertiesOutput() AKSResponsePropertiesOutput {
	return i.ToAKSResponsePropertiesOutputWithContext(context.Background())
}

func (i AKSResponsePropertiesArgs) ToAKSResponsePropertiesOutputWithContext(ctx context.Context) AKSResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSResponsePropertiesOutput)
}

func (i AKSResponsePropertiesArgs) ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput {
	return i.ToAKSResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i AKSResponsePropertiesArgs) ToAKSResponsePropertiesPtrOutputWithContext(ctx context.Context) AKSResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSResponsePropertiesOutput).ToAKSResponsePropertiesPtrOutputWithContext(ctx)
}









type AKSResponsePropertiesPtrInput interface {
	pulumi.Input

	ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput
	ToAKSResponsePropertiesPtrOutputWithContext(context.Context) AKSResponsePropertiesPtrOutput
}

type aksresponsePropertiesPtrType AKSResponsePropertiesArgs

func AKSResponsePropertiesPtr(v *AKSResponsePropertiesArgs) AKSResponsePropertiesPtrInput {
	return (*aksresponsePropertiesPtrType)(v)
}

func (*aksresponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSResponseProperties)(nil)).Elem()
}

func (i *aksresponsePropertiesPtrType) ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput {
	return i.ToAKSResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *aksresponsePropertiesPtrType) ToAKSResponsePropertiesPtrOutputWithContext(ctx context.Context) AKSResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AKSResponsePropertiesPtrOutput)
}

type AKSResponsePropertiesOutput struct{ *pulumi.OutputState }

func (AKSResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSResponseProperties)(nil)).Elem()
}

func (o AKSResponsePropertiesOutput) ToAKSResponsePropertiesOutput() AKSResponsePropertiesOutput {
	return o
}

func (o AKSResponsePropertiesOutput) ToAKSResponsePropertiesOutputWithContext(ctx context.Context) AKSResponsePropertiesOutput {
	return o
}

func (o AKSResponsePropertiesOutput) ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput {
	return o.ToAKSResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o AKSResponsePropertiesOutput) ToAKSResponsePropertiesPtrOutputWithContext(ctx context.Context) AKSResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSResponseProperties) *AKSResponseProperties {
		return &v
	}).(AKSResponsePropertiesPtrOutput)
}

func (o AKSResponsePropertiesOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *int { return v.AgentCount }).(pulumi.IntPtrOutput)
}

func (o AKSResponsePropertiesOutput) AgentVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.AgentVMSize }).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesOutput) AksNetworkingConfiguration() AksNetworkingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *AksNetworkingConfigurationResponse { return v.AksNetworkingConfiguration }).(AksNetworkingConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.ClusterFqdn }).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesOutput) SslConfiguration() SslConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *SslConfigurationResponse { return v.SslConfiguration }).(SslConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesOutput) SystemServices() SystemServiceResponseArrayOutput {
	return o.ApplyT(func(v AKSResponseProperties) []SystemServiceResponse { return v.SystemServices }).(SystemServiceResponseArrayOutput)
}

type AKSResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AKSResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSResponseProperties)(nil)).Elem()
}

func (o AKSResponsePropertiesPtrOutput) ToAKSResponsePropertiesPtrOutput() AKSResponsePropertiesPtrOutput {
	return o
}

func (o AKSResponsePropertiesPtrOutput) ToAKSResponsePropertiesPtrOutputWithContext(ctx context.Context) AKSResponsePropertiesPtrOutput {
	return o
}

func (o AKSResponsePropertiesPtrOutput) Elem() AKSResponsePropertiesOutput {
	return o.ApplyT(func(v *AKSResponseProperties) AKSResponseProperties {
		if v != nil {
			return *v
		}
		var ret AKSResponseProperties
		return ret
	}).(AKSResponsePropertiesOutput)
}

func (o AKSResponsePropertiesPtrOutput) AgentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.AgentCount
	}).(pulumi.IntPtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) AgentVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.AgentVMSize
	}).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) AksNetworkingConfiguration() AksNetworkingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *AksNetworkingConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.AksNetworkingConfiguration
	}).(AksNetworkingConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterFqdn
	}).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) SslConfiguration() SslConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *SslConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.SslConfiguration
	}).(SslConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) SystemServices() SystemServiceResponseArrayOutput {
	return o.ApplyT(func(v *AKSResponseProperties) []SystemServiceResponse {
		if v == nil {
			return nil
		}
		return v.SystemServices
	}).(SystemServiceResponseArrayOutput)
}

type AksNetworkingConfiguration struct {
	DnsServiceIP     *string `pulumi:"dnsServiceIP"`
	DockerBridgeCidr *string `pulumi:"dockerBridgeCidr"`
	ServiceCidr      *string `pulumi:"serviceCidr"`
	SubnetId         *string `pulumi:"subnetId"`
}





type AksNetworkingConfigurationInput interface {
	pulumi.Input

	ToAksNetworkingConfigurationOutput() AksNetworkingConfigurationOutput
	ToAksNetworkingConfigurationOutputWithContext(context.Context) AksNetworkingConfigurationOutput
}

type AksNetworkingConfigurationArgs struct {
	DnsServiceIP     pulumi.StringPtrInput `pulumi:"dnsServiceIP"`
	DockerBridgeCidr pulumi.StringPtrInput `pulumi:"dockerBridgeCidr"`
	ServiceCidr      pulumi.StringPtrInput `pulumi:"serviceCidr"`
	SubnetId         pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (AksNetworkingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AksNetworkingConfiguration)(nil)).Elem()
}

func (i AksNetworkingConfigurationArgs) ToAksNetworkingConfigurationOutput() AksNetworkingConfigurationOutput {
	return i.ToAksNetworkingConfigurationOutputWithContext(context.Background())
}

func (i AksNetworkingConfigurationArgs) ToAksNetworkingConfigurationOutputWithContext(ctx context.Context) AksNetworkingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationOutput)
}

func (i AksNetworkingConfigurationArgs) ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput {
	return i.ToAksNetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (i AksNetworkingConfigurationArgs) ToAksNetworkingConfigurationPtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationOutput).ToAksNetworkingConfigurationPtrOutputWithContext(ctx)
}









type AksNetworkingConfigurationPtrInput interface {
	pulumi.Input

	ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput
	ToAksNetworkingConfigurationPtrOutputWithContext(context.Context) AksNetworkingConfigurationPtrOutput
}

type aksNetworkingConfigurationPtrType AksNetworkingConfigurationArgs

func AksNetworkingConfigurationPtr(v *AksNetworkingConfigurationArgs) AksNetworkingConfigurationPtrInput {
	return (*aksNetworkingConfigurationPtrType)(v)
}

func (*aksNetworkingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AksNetworkingConfiguration)(nil)).Elem()
}

func (i *aksNetworkingConfigurationPtrType) ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput {
	return i.ToAksNetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (i *aksNetworkingConfigurationPtrType) ToAksNetworkingConfigurationPtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationPtrOutput)
}

type AksNetworkingConfigurationOutput struct{ *pulumi.OutputState }

func (AksNetworkingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AksNetworkingConfiguration)(nil)).Elem()
}

func (o AksNetworkingConfigurationOutput) ToAksNetworkingConfigurationOutput() AksNetworkingConfigurationOutput {
	return o
}

func (o AksNetworkingConfigurationOutput) ToAksNetworkingConfigurationOutputWithContext(ctx context.Context) AksNetworkingConfigurationOutput {
	return o
}

func (o AksNetworkingConfigurationOutput) ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput {
	return o.ToAksNetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (o AksNetworkingConfigurationOutput) ToAksNetworkingConfigurationPtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AksNetworkingConfiguration) *AksNetworkingConfiguration {
		return &v
	}).(AksNetworkingConfigurationPtrOutput)
}

func (o AksNetworkingConfigurationOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfiguration) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfiguration) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfiguration) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfiguration) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type AksNetworkingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AksNetworkingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AksNetworkingConfiguration)(nil)).Elem()
}

func (o AksNetworkingConfigurationPtrOutput) ToAksNetworkingConfigurationPtrOutput() AksNetworkingConfigurationPtrOutput {
	return o
}

func (o AksNetworkingConfigurationPtrOutput) ToAksNetworkingConfigurationPtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationPtrOutput {
	return o
}

func (o AksNetworkingConfigurationPtrOutput) Elem() AksNetworkingConfigurationOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) AksNetworkingConfiguration {
		if v != nil {
			return *v
		}
		var ret AksNetworkingConfiguration
		return ret
	}).(AksNetworkingConfigurationOutput)
}

func (o AksNetworkingConfigurationPtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationPtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationPtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type AksNetworkingConfigurationResponse struct {
	DnsServiceIP     *string `pulumi:"dnsServiceIP"`
	DockerBridgeCidr *string `pulumi:"dockerBridgeCidr"`
	ServiceCidr      *string `pulumi:"serviceCidr"`
	SubnetId         *string `pulumi:"subnetId"`
}





type AksNetworkingConfigurationResponseInput interface {
	pulumi.Input

	ToAksNetworkingConfigurationResponseOutput() AksNetworkingConfigurationResponseOutput
	ToAksNetworkingConfigurationResponseOutputWithContext(context.Context) AksNetworkingConfigurationResponseOutput
}

type AksNetworkingConfigurationResponseArgs struct {
	DnsServiceIP     pulumi.StringPtrInput `pulumi:"dnsServiceIP"`
	DockerBridgeCidr pulumi.StringPtrInput `pulumi:"dockerBridgeCidr"`
	ServiceCidr      pulumi.StringPtrInput `pulumi:"serviceCidr"`
	SubnetId         pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (AksNetworkingConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AksNetworkingConfigurationResponse)(nil)).Elem()
}

func (i AksNetworkingConfigurationResponseArgs) ToAksNetworkingConfigurationResponseOutput() AksNetworkingConfigurationResponseOutput {
	return i.ToAksNetworkingConfigurationResponseOutputWithContext(context.Background())
}

func (i AksNetworkingConfigurationResponseArgs) ToAksNetworkingConfigurationResponseOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationResponseOutput)
}

func (i AksNetworkingConfigurationResponseArgs) ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput {
	return i.ToAksNetworkingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i AksNetworkingConfigurationResponseArgs) ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationResponseOutput).ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx)
}









type AksNetworkingConfigurationResponsePtrInput interface {
	pulumi.Input

	ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput
	ToAksNetworkingConfigurationResponsePtrOutputWithContext(context.Context) AksNetworkingConfigurationResponsePtrOutput
}

type aksNetworkingConfigurationResponsePtrType AksNetworkingConfigurationResponseArgs

func AksNetworkingConfigurationResponsePtr(v *AksNetworkingConfigurationResponseArgs) AksNetworkingConfigurationResponsePtrInput {
	return (*aksNetworkingConfigurationResponsePtrType)(v)
}

func (*aksNetworkingConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AksNetworkingConfigurationResponse)(nil)).Elem()
}

func (i *aksNetworkingConfigurationResponsePtrType) ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput {
	return i.ToAksNetworkingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *aksNetworkingConfigurationResponsePtrType) ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AksNetworkingConfigurationResponsePtrOutput)
}

type AksNetworkingConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AksNetworkingConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AksNetworkingConfigurationResponse)(nil)).Elem()
}

func (o AksNetworkingConfigurationResponseOutput) ToAksNetworkingConfigurationResponseOutput() AksNetworkingConfigurationResponseOutput {
	return o
}

func (o AksNetworkingConfigurationResponseOutput) ToAksNetworkingConfigurationResponseOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponseOutput {
	return o
}

func (o AksNetworkingConfigurationResponseOutput) ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput {
	return o.ToAksNetworkingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o AksNetworkingConfigurationResponseOutput) ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AksNetworkingConfigurationResponse) *AksNetworkingConfigurationResponse {
		return &v
	}).(AksNetworkingConfigurationResponsePtrOutput)
}

func (o AksNetworkingConfigurationResponseOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfigurationResponse) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponseOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfigurationResponse) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfigurationResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AksNetworkingConfigurationResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type AksNetworkingConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AksNetworkingConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AksNetworkingConfigurationResponse)(nil)).Elem()
}

func (o AksNetworkingConfigurationResponsePtrOutput) ToAksNetworkingConfigurationResponsePtrOutput() AksNetworkingConfigurationResponsePtrOutput {
	return o
}

func (o AksNetworkingConfigurationResponsePtrOutput) ToAksNetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AksNetworkingConfigurationResponsePtrOutput {
	return o
}

func (o AksNetworkingConfigurationResponsePtrOutput) Elem() AksNetworkingConfigurationResponseOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) AksNetworkingConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AksNetworkingConfigurationResponse
		return ret
	}).(AksNetworkingConfigurationResponseOutput)
}

func (o AksNetworkingConfigurationResponsePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponsePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o AksNetworkingConfigurationResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AksNetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type AmlCompute struct {
	ComputeLocation *string               `pulumi:"computeLocation"`
	ComputeType     string                `pulumi:"computeType"`
	Description     *string               `pulumi:"description"`
	Properties      *AmlComputeProperties `pulumi:"properties"`
	ResourceId      *string               `pulumi:"resourceId"`
}





type AmlComputeInput interface {
	pulumi.Input

	ToAmlComputeOutput() AmlComputeOutput
	ToAmlComputeOutputWithContext(context.Context) AmlComputeOutput
}

type AmlComputeArgs struct {
	ComputeLocation pulumi.StringPtrInput        `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput           `pulumi:"computeType"`
	Description     pulumi.StringPtrInput        `pulumi:"description"`
	Properties      AmlComputePropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput        `pulumi:"resourceId"`
}

func (AmlComputeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlCompute)(nil)).Elem()
}

func (i AmlComputeArgs) ToAmlComputeOutput() AmlComputeOutput {
	return i.ToAmlComputeOutputWithContext(context.Background())
}

func (i AmlComputeArgs) ToAmlComputeOutputWithContext(ctx context.Context) AmlComputeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeOutput)
}

type AmlComputeOutput struct{ *pulumi.OutputState }

func (AmlComputeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlCompute)(nil)).Elem()
}

func (o AmlComputeOutput) ToAmlComputeOutput() AmlComputeOutput {
	return o
}

func (o AmlComputeOutput) ToAmlComputeOutputWithContext(ctx context.Context) AmlComputeOutput {
	return o
}

func (o AmlComputeOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlCompute) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o AmlComputeOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AmlCompute) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AmlComputeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlCompute) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AmlComputeOutput) Properties() AmlComputePropertiesPtrOutput {
	return o.ApplyT(func(v AmlCompute) *AmlComputeProperties { return v.Properties }).(AmlComputePropertiesPtrOutput)
}

func (o AmlComputeOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlCompute) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AmlComputeNodeInformationResponse struct {
	NodeId           string  `pulumi:"nodeId"`
	NodeState        string  `pulumi:"nodeState"`
	Port             float64 `pulumi:"port"`
	PrivateIpAddress string  `pulumi:"privateIpAddress"`
	PublicIpAddress  string  `pulumi:"publicIpAddress"`
	RunId            string  `pulumi:"runId"`
}





type AmlComputeNodeInformationResponseInput interface {
	pulumi.Input

	ToAmlComputeNodeInformationResponseOutput() AmlComputeNodeInformationResponseOutput
	ToAmlComputeNodeInformationResponseOutputWithContext(context.Context) AmlComputeNodeInformationResponseOutput
}

type AmlComputeNodeInformationResponseArgs struct {
	NodeId           pulumi.StringInput  `pulumi:"nodeId"`
	NodeState        pulumi.StringInput  `pulumi:"nodeState"`
	Port             pulumi.Float64Input `pulumi:"port"`
	PrivateIpAddress pulumi.StringInput  `pulumi:"privateIpAddress"`
	PublicIpAddress  pulumi.StringInput  `pulumi:"publicIpAddress"`
	RunId            pulumi.StringInput  `pulumi:"runId"`
}

func (AmlComputeNodeInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeNodeInformationResponse)(nil)).Elem()
}

func (i AmlComputeNodeInformationResponseArgs) ToAmlComputeNodeInformationResponseOutput() AmlComputeNodeInformationResponseOutput {
	return i.ToAmlComputeNodeInformationResponseOutputWithContext(context.Background())
}

func (i AmlComputeNodeInformationResponseArgs) ToAmlComputeNodeInformationResponseOutputWithContext(ctx context.Context) AmlComputeNodeInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeNodeInformationResponseOutput)
}





type AmlComputeNodeInformationResponseArrayInput interface {
	pulumi.Input

	ToAmlComputeNodeInformationResponseArrayOutput() AmlComputeNodeInformationResponseArrayOutput
	ToAmlComputeNodeInformationResponseArrayOutputWithContext(context.Context) AmlComputeNodeInformationResponseArrayOutput
}

type AmlComputeNodeInformationResponseArray []AmlComputeNodeInformationResponseInput

func (AmlComputeNodeInformationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AmlComputeNodeInformationResponse)(nil)).Elem()
}

func (i AmlComputeNodeInformationResponseArray) ToAmlComputeNodeInformationResponseArrayOutput() AmlComputeNodeInformationResponseArrayOutput {
	return i.ToAmlComputeNodeInformationResponseArrayOutputWithContext(context.Background())
}

func (i AmlComputeNodeInformationResponseArray) ToAmlComputeNodeInformationResponseArrayOutputWithContext(ctx context.Context) AmlComputeNodeInformationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeNodeInformationResponseArrayOutput)
}

type AmlComputeNodeInformationResponseOutput struct{ *pulumi.OutputState }

func (AmlComputeNodeInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeNodeInformationResponse)(nil)).Elem()
}

func (o AmlComputeNodeInformationResponseOutput) ToAmlComputeNodeInformationResponseOutput() AmlComputeNodeInformationResponseOutput {
	return o
}

func (o AmlComputeNodeInformationResponseOutput) ToAmlComputeNodeInformationResponseOutputWithContext(ctx context.Context) AmlComputeNodeInformationResponseOutput {
	return o
}

func (o AmlComputeNodeInformationResponseOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.NodeId }).(pulumi.StringOutput)
}

func (o AmlComputeNodeInformationResponseOutput) NodeState() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.NodeState }).(pulumi.StringOutput)
}

func (o AmlComputeNodeInformationResponseOutput) Port() pulumi.Float64Output {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) float64 { return v.Port }).(pulumi.Float64Output)
}

func (o AmlComputeNodeInformationResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o AmlComputeNodeInformationResponseOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

func (o AmlComputeNodeInformationResponseOutput) RunId() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeNodeInformationResponse) string { return v.RunId }).(pulumi.StringOutput)
}

type AmlComputeNodeInformationResponseArrayOutput struct{ *pulumi.OutputState }

func (AmlComputeNodeInformationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AmlComputeNodeInformationResponse)(nil)).Elem()
}

func (o AmlComputeNodeInformationResponseArrayOutput) ToAmlComputeNodeInformationResponseArrayOutput() AmlComputeNodeInformationResponseArrayOutput {
	return o
}

func (o AmlComputeNodeInformationResponseArrayOutput) ToAmlComputeNodeInformationResponseArrayOutputWithContext(ctx context.Context) AmlComputeNodeInformationResponseArrayOutput {
	return o
}

func (o AmlComputeNodeInformationResponseArrayOutput) Index(i pulumi.IntInput) AmlComputeNodeInformationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AmlComputeNodeInformationResponse {
		return vs[0].([]AmlComputeNodeInformationResponse)[vs[1].(int)]
	}).(AmlComputeNodeInformationResponseOutput)
}

type AmlComputeProperties struct {
	RemoteLoginPortPublicAccess *string                 `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings               *ScaleSettings          `pulumi:"scaleSettings"`
	Subnet                      *ResourceId             `pulumi:"subnet"`
	UserAccountCredentials      *UserAccountCredentials `pulumi:"userAccountCredentials"`
	VmPriority                  *string                 `pulumi:"vmPriority"`
	VmSize                      *string                 `pulumi:"vmSize"`
}





type AmlComputePropertiesInput interface {
	pulumi.Input

	ToAmlComputePropertiesOutput() AmlComputePropertiesOutput
	ToAmlComputePropertiesOutputWithContext(context.Context) AmlComputePropertiesOutput
}

type AmlComputePropertiesArgs struct {
	RemoteLoginPortPublicAccess pulumi.StringPtrInput          `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings               ScaleSettingsPtrInput          `pulumi:"scaleSettings"`
	Subnet                      ResourceIdPtrInput             `pulumi:"subnet"`
	UserAccountCredentials      UserAccountCredentialsPtrInput `pulumi:"userAccountCredentials"`
	VmPriority                  pulumi.StringPtrInput          `pulumi:"vmPriority"`
	VmSize                      pulumi.StringPtrInput          `pulumi:"vmSize"`
}

func (AmlComputePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeProperties)(nil)).Elem()
}

func (i AmlComputePropertiesArgs) ToAmlComputePropertiesOutput() AmlComputePropertiesOutput {
	return i.ToAmlComputePropertiesOutputWithContext(context.Background())
}

func (i AmlComputePropertiesArgs) ToAmlComputePropertiesOutputWithContext(ctx context.Context) AmlComputePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesOutput)
}

func (i AmlComputePropertiesArgs) ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput {
	return i.ToAmlComputePropertiesPtrOutputWithContext(context.Background())
}

func (i AmlComputePropertiesArgs) ToAmlComputePropertiesPtrOutputWithContext(ctx context.Context) AmlComputePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesOutput).ToAmlComputePropertiesPtrOutputWithContext(ctx)
}









type AmlComputePropertiesPtrInput interface {
	pulumi.Input

	ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput
	ToAmlComputePropertiesPtrOutputWithContext(context.Context) AmlComputePropertiesPtrOutput
}

type amlComputePropertiesPtrType AmlComputePropertiesArgs

func AmlComputePropertiesPtr(v *AmlComputePropertiesArgs) AmlComputePropertiesPtrInput {
	return (*amlComputePropertiesPtrType)(v)
}

func (*amlComputePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputeProperties)(nil)).Elem()
}

func (i *amlComputePropertiesPtrType) ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput {
	return i.ToAmlComputePropertiesPtrOutputWithContext(context.Background())
}

func (i *amlComputePropertiesPtrType) ToAmlComputePropertiesPtrOutputWithContext(ctx context.Context) AmlComputePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesPtrOutput)
}

type AmlComputePropertiesOutput struct{ *pulumi.OutputState }

func (AmlComputePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeProperties)(nil)).Elem()
}

func (o AmlComputePropertiesOutput) ToAmlComputePropertiesOutput() AmlComputePropertiesOutput {
	return o
}

func (o AmlComputePropertiesOutput) ToAmlComputePropertiesOutputWithContext(ctx context.Context) AmlComputePropertiesOutput {
	return o
}

func (o AmlComputePropertiesOutput) ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput {
	return o.ToAmlComputePropertiesPtrOutputWithContext(context.Background())
}

func (o AmlComputePropertiesOutput) ToAmlComputePropertiesPtrOutputWithContext(ctx context.Context) AmlComputePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AmlComputeProperties) *AmlComputeProperties {
		return &v
	}).(AmlComputePropertiesPtrOutput)
}

func (o AmlComputePropertiesOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *string { return v.RemoteLoginPortPublicAccess }).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesOutput) ScaleSettings() ScaleSettingsPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *ScaleSettings { return v.ScaleSettings }).(ScaleSettingsPtrOutput)
}

func (o AmlComputePropertiesOutput) Subnet() ResourceIdPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *ResourceId { return v.Subnet }).(ResourceIdPtrOutput)
}

func (o AmlComputePropertiesOutput) UserAccountCredentials() UserAccountCredentialsPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *UserAccountCredentials { return v.UserAccountCredentials }).(UserAccountCredentialsPtrOutput)
}

func (o AmlComputePropertiesOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *string { return v.VmPriority }).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type AmlComputePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AmlComputePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputeProperties)(nil)).Elem()
}

func (o AmlComputePropertiesPtrOutput) ToAmlComputePropertiesPtrOutput() AmlComputePropertiesPtrOutput {
	return o
}

func (o AmlComputePropertiesPtrOutput) ToAmlComputePropertiesPtrOutputWithContext(ctx context.Context) AmlComputePropertiesPtrOutput {
	return o
}

func (o AmlComputePropertiesPtrOutput) Elem() AmlComputePropertiesOutput {
	return o.ApplyT(func(v *AmlComputeProperties) AmlComputeProperties {
		if v != nil {
			return *v
		}
		var ret AmlComputeProperties
		return ret
	}).(AmlComputePropertiesOutput)
}

func (o AmlComputePropertiesPtrOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *string {
		if v == nil {
			return nil
		}
		return v.RemoteLoginPortPublicAccess
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) ScaleSettings() ScaleSettingsPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *ScaleSettings {
		if v == nil {
			return nil
		}
		return v.ScaleSettings
	}).(ScaleSettingsPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) Subnet() ResourceIdPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *ResourceId {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ResourceIdPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) UserAccountCredentials() UserAccountCredentialsPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *UserAccountCredentials {
		if v == nil {
			return nil
		}
		return v.UserAccountCredentials
	}).(UserAccountCredentialsPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmPriority
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type AmlComputeResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *AmlComputeResponseProperties         `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}





type AmlComputeResponseInput interface {
	pulumi.Input

	ToAmlComputeResponseOutput() AmlComputeResponseOutput
	ToAmlComputeResponseOutputWithContext(context.Context) AmlComputeResponseOutput
}

type AmlComputeResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         AmlComputeResponsePropertiesPtrInput          `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (AmlComputeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeResponse)(nil)).Elem()
}

func (i AmlComputeResponseArgs) ToAmlComputeResponseOutput() AmlComputeResponseOutput {
	return i.ToAmlComputeResponseOutputWithContext(context.Background())
}

func (i AmlComputeResponseArgs) ToAmlComputeResponseOutputWithContext(ctx context.Context) AmlComputeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeResponseOutput)
}

type AmlComputeResponseOutput struct{ *pulumi.OutputState }

func (AmlComputeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeResponse)(nil)).Elem()
}

func (o AmlComputeResponseOutput) ToAmlComputeResponseOutput() AmlComputeResponseOutput {
	return o
}

func (o AmlComputeResponseOutput) ToAmlComputeResponseOutputWithContext(ctx context.Context) AmlComputeResponseOutput {
	return o
}

func (o AmlComputeResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v AmlComputeResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o AmlComputeResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) Properties() AmlComputeResponsePropertiesPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *AmlComputeResponseProperties { return v.Properties }).(AmlComputeResponsePropertiesPtrOutput)
}

func (o AmlComputeResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v AmlComputeResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o AmlComputeResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AmlComputeResponseProperties struct {
	AllocationState               string                                `pulumi:"allocationState"`
	AllocationStateTransitionTime string                                `pulumi:"allocationStateTransitionTime"`
	CurrentNodeCount              int                                   `pulumi:"currentNodeCount"`
	Errors                        []MachineLearningServiceErrorResponse `pulumi:"errors"`
	NodeStateCounts               NodeStateCountsResponse               `pulumi:"nodeStateCounts"`
	RemoteLoginPortPublicAccess   *string                               `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 *ScaleSettingsResponse                `pulumi:"scaleSettings"`
	Subnet                        *ResourceIdResponse                   `pulumi:"subnet"`
	TargetNodeCount               int                                   `pulumi:"targetNodeCount"`
	UserAccountCredentials        *UserAccountCredentialsResponse       `pulumi:"userAccountCredentials"`
	VmPriority                    *string                               `pulumi:"vmPriority"`
	VmSize                        *string                               `pulumi:"vmSize"`
}





type AmlComputeResponsePropertiesInput interface {
	pulumi.Input

	ToAmlComputeResponsePropertiesOutput() AmlComputeResponsePropertiesOutput
	ToAmlComputeResponsePropertiesOutputWithContext(context.Context) AmlComputeResponsePropertiesOutput
}

type AmlComputeResponsePropertiesArgs struct {
	AllocationState               pulumi.StringInput                            `pulumi:"allocationState"`
	AllocationStateTransitionTime pulumi.StringInput                            `pulumi:"allocationStateTransitionTime"`
	CurrentNodeCount              pulumi.IntInput                               `pulumi:"currentNodeCount"`
	Errors                        MachineLearningServiceErrorResponseArrayInput `pulumi:"errors"`
	NodeStateCounts               NodeStateCountsResponseInput                  `pulumi:"nodeStateCounts"`
	RemoteLoginPortPublicAccess   pulumi.StringPtrInput                         `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 ScaleSettingsResponsePtrInput                 `pulumi:"scaleSettings"`
	Subnet                        ResourceIdResponsePtrInput                    `pulumi:"subnet"`
	TargetNodeCount               pulumi.IntInput                               `pulumi:"targetNodeCount"`
	UserAccountCredentials        UserAccountCredentialsResponsePtrInput        `pulumi:"userAccountCredentials"`
	VmPriority                    pulumi.StringPtrInput                         `pulumi:"vmPriority"`
	VmSize                        pulumi.StringPtrInput                         `pulumi:"vmSize"`
}

func (AmlComputeResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeResponseProperties)(nil)).Elem()
}

func (i AmlComputeResponsePropertiesArgs) ToAmlComputeResponsePropertiesOutput() AmlComputeResponsePropertiesOutput {
	return i.ToAmlComputeResponsePropertiesOutputWithContext(context.Background())
}

func (i AmlComputeResponsePropertiesArgs) ToAmlComputeResponsePropertiesOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeResponsePropertiesOutput)
}

func (i AmlComputeResponsePropertiesArgs) ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput {
	return i.ToAmlComputeResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i AmlComputeResponsePropertiesArgs) ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeResponsePropertiesOutput).ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx)
}









type AmlComputeResponsePropertiesPtrInput interface {
	pulumi.Input

	ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput
	ToAmlComputeResponsePropertiesPtrOutputWithContext(context.Context) AmlComputeResponsePropertiesPtrOutput
}

type amlComputeResponsePropertiesPtrType AmlComputeResponsePropertiesArgs

func AmlComputeResponsePropertiesPtr(v *AmlComputeResponsePropertiesArgs) AmlComputeResponsePropertiesPtrInput {
	return (*amlComputeResponsePropertiesPtrType)(v)
}

func (*amlComputeResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputeResponseProperties)(nil)).Elem()
}

func (i *amlComputeResponsePropertiesPtrType) ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput {
	return i.ToAmlComputeResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *amlComputeResponsePropertiesPtrType) ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputeResponsePropertiesPtrOutput)
}

type AmlComputeResponsePropertiesOutput struct{ *pulumi.OutputState }

func (AmlComputeResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputeResponseProperties)(nil)).Elem()
}

func (o AmlComputeResponsePropertiesOutput) ToAmlComputeResponsePropertiesOutput() AmlComputeResponsePropertiesOutput {
	return o
}

func (o AmlComputeResponsePropertiesOutput) ToAmlComputeResponsePropertiesOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesOutput {
	return o
}

func (o AmlComputeResponsePropertiesOutput) ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput {
	return o.ToAmlComputeResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o AmlComputeResponsePropertiesOutput) ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AmlComputeResponseProperties) *AmlComputeResponseProperties {
		return &v
	}).(AmlComputeResponsePropertiesPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) AllocationState() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) string { return v.AllocationState }).(pulumi.StringOutput)
}

func (o AmlComputeResponsePropertiesOutput) AllocationStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) string { return v.AllocationStateTransitionTime }).(pulumi.StringOutput)
}

func (o AmlComputeResponsePropertiesOutput) CurrentNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) int { return v.CurrentNodeCount }).(pulumi.IntOutput)
}

func (o AmlComputeResponsePropertiesOutput) Errors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) []MachineLearningServiceErrorResponse { return v.Errors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o AmlComputeResponsePropertiesOutput) NodeStateCounts() NodeStateCountsResponseOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) NodeStateCountsResponse { return v.NodeStateCounts }).(NodeStateCountsResponseOutput)
}

func (o AmlComputeResponsePropertiesOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *string { return v.RemoteLoginPortPublicAccess }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) ScaleSettings() ScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *ScaleSettingsResponse { return v.ScaleSettings }).(ScaleSettingsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *ResourceIdResponse { return v.Subnet }).(ResourceIdResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) TargetNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) int { return v.TargetNodeCount }).(pulumi.IntOutput)
}

func (o AmlComputeResponsePropertiesOutput) UserAccountCredentials() UserAccountCredentialsResponsePtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *UserAccountCredentialsResponse { return v.UserAccountCredentials }).(UserAccountCredentialsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *string { return v.VmPriority }).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponseProperties) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type AmlComputeResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AmlComputeResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputeResponseProperties)(nil)).Elem()
}

func (o AmlComputeResponsePropertiesPtrOutput) ToAmlComputeResponsePropertiesPtrOutput() AmlComputeResponsePropertiesPtrOutput {
	return o
}

func (o AmlComputeResponsePropertiesPtrOutput) ToAmlComputeResponsePropertiesPtrOutputWithContext(ctx context.Context) AmlComputeResponsePropertiesPtrOutput {
	return o
}

func (o AmlComputeResponsePropertiesPtrOutput) Elem() AmlComputeResponsePropertiesOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) AmlComputeResponseProperties {
		if v != nil {
			return *v
		}
		var ret AmlComputeResponseProperties
		return ret
	}).(AmlComputeResponsePropertiesOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) AllocationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AllocationState
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) AllocationStateTransitionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AllocationStateTransitionTime
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) CurrentNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *int {
		if v == nil {
			return nil
		}
		return &v.CurrentNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) Errors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) []MachineLearningServiceErrorResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) NodeStateCounts() NodeStateCountsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *NodeStateCountsResponse {
		if v == nil {
			return nil
		}
		return &v.NodeStateCounts
	}).(NodeStateCountsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.RemoteLoginPortPublicAccess
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) ScaleSettings() ScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *ScaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ScaleSettings
	}).(ScaleSettingsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ResourceIdResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) TargetNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *int {
		if v == nil {
			return nil
		}
		return &v.TargetNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) UserAccountCredentials() UserAccountCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *UserAccountCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.UserAccountCredentials
	}).(UserAccountCredentialsResponsePtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmPriority
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputeResponsePropertiesPtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type DataFactory struct {
	ComputeLocation *string `pulumi:"computeLocation"`
	ComputeType     string  `pulumi:"computeType"`
	Description     *string `pulumi:"description"`
	ResourceId      *string `pulumi:"resourceId"`
}





type DataFactoryInput interface {
	pulumi.Input

	ToDataFactoryOutput() DataFactoryOutput
	ToDataFactoryOutputWithContext(context.Context) DataFactoryOutput
}

type DataFactoryArgs struct {
	ComputeLocation pulumi.StringPtrInput `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput    `pulumi:"computeType"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	ResourceId      pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (DataFactoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFactory)(nil)).Elem()
}

func (i DataFactoryArgs) ToDataFactoryOutput() DataFactoryOutput {
	return i.ToDataFactoryOutputWithContext(context.Background())
}

func (i DataFactoryArgs) ToDataFactoryOutputWithContext(ctx context.Context) DataFactoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFactoryOutput)
}

type DataFactoryOutput struct{ *pulumi.OutputState }

func (DataFactoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFactory)(nil)).Elem()
}

func (o DataFactoryOutput) ToDataFactoryOutput() DataFactoryOutput {
	return o
}

func (o DataFactoryOutput) ToDataFactoryOutputWithContext(ctx context.Context) DataFactoryOutput {
	return o
}

func (o DataFactoryOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactory) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DataFactoryOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactory) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DataFactoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactory) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataFactoryOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactory) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataFactoryResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}





type DataFactoryResponseInput interface {
	pulumi.Input

	ToDataFactoryResponseOutput() DataFactoryResponseOutput
	ToDataFactoryResponseOutputWithContext(context.Context) DataFactoryResponseOutput
}

type DataFactoryResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (DataFactoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFactoryResponse)(nil)).Elem()
}

func (i DataFactoryResponseArgs) ToDataFactoryResponseOutput() DataFactoryResponseOutput {
	return i.ToDataFactoryResponseOutputWithContext(context.Background())
}

func (i DataFactoryResponseArgs) ToDataFactoryResponseOutputWithContext(ctx context.Context) DataFactoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFactoryResponseOutput)
}

type DataFactoryResponseOutput struct{ *pulumi.OutputState }

func (DataFactoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFactoryResponse)(nil)).Elem()
}

func (o DataFactoryResponseOutput) ToDataFactoryResponseOutput() DataFactoryResponseOutput {
	return o
}

func (o DataFactoryResponseOutput) ToDataFactoryResponseOutputWithContext(ctx context.Context) DataFactoryResponseOutput {
	return o
}

func (o DataFactoryResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactoryResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DataFactoryResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactoryResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataFactoryResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v DataFactoryResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o DataFactoryResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v DataFactoryResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o DataFactoryResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactoryResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeAnalytics struct {
	ComputeLocation *string                      `pulumi:"computeLocation"`
	ComputeType     string                       `pulumi:"computeType"`
	Description     *string                      `pulumi:"description"`
	Properties      *DataLakeAnalyticsProperties `pulumi:"properties"`
	ResourceId      *string                      `pulumi:"resourceId"`
}





type DataLakeAnalyticsInput interface {
	pulumi.Input

	ToDataLakeAnalyticsOutput() DataLakeAnalyticsOutput
	ToDataLakeAnalyticsOutputWithContext(context.Context) DataLakeAnalyticsOutput
}

type DataLakeAnalyticsArgs struct {
	ComputeLocation pulumi.StringPtrInput               `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput                  `pulumi:"computeType"`
	Description     pulumi.StringPtrInput               `pulumi:"description"`
	Properties      DataLakeAnalyticsPropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput               `pulumi:"resourceId"`
}

func (DataLakeAnalyticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalytics)(nil)).Elem()
}

func (i DataLakeAnalyticsArgs) ToDataLakeAnalyticsOutput() DataLakeAnalyticsOutput {
	return i.ToDataLakeAnalyticsOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsArgs) ToDataLakeAnalyticsOutputWithContext(ctx context.Context) DataLakeAnalyticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsOutput)
}

type DataLakeAnalyticsOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalytics)(nil)).Elem()
}

func (o DataLakeAnalyticsOutput) ToDataLakeAnalyticsOutput() DataLakeAnalyticsOutput {
	return o
}

func (o DataLakeAnalyticsOutput) ToDataLakeAnalyticsOutputWithContext(ctx context.Context) DataLakeAnalyticsOutput {
	return o
}

func (o DataLakeAnalyticsOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DataLakeAnalyticsOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalytics) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataLakeAnalyticsOutput) Properties() DataLakeAnalyticsPropertiesPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *DataLakeAnalyticsProperties { return v.Properties }).(DataLakeAnalyticsPropertiesPtrOutput)
}

func (o DataLakeAnalyticsOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}





type DataLakeAnalyticsPropertiesInput interface {
	pulumi.Input

	ToDataLakeAnalyticsPropertiesOutput() DataLakeAnalyticsPropertiesOutput
	ToDataLakeAnalyticsPropertiesOutputWithContext(context.Context) DataLakeAnalyticsPropertiesOutput
}

type DataLakeAnalyticsPropertiesArgs struct {
	DataLakeStoreAccountName pulumi.StringPtrInput `pulumi:"dataLakeStoreAccountName"`
}

func (DataLakeAnalyticsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsProperties)(nil)).Elem()
}

func (i DataLakeAnalyticsPropertiesArgs) ToDataLakeAnalyticsPropertiesOutput() DataLakeAnalyticsPropertiesOutput {
	return i.ToDataLakeAnalyticsPropertiesOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsPropertiesArgs) ToDataLakeAnalyticsPropertiesOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsPropertiesOutput)
}

func (i DataLakeAnalyticsPropertiesArgs) ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput {
	return i.ToDataLakeAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsPropertiesArgs) ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsPropertiesOutput).ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx)
}









type DataLakeAnalyticsPropertiesPtrInput interface {
	pulumi.Input

	ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput
	ToDataLakeAnalyticsPropertiesPtrOutputWithContext(context.Context) DataLakeAnalyticsPropertiesPtrOutput
}

type dataLakeAnalyticsPropertiesPtrType DataLakeAnalyticsPropertiesArgs

func DataLakeAnalyticsPropertiesPtr(v *DataLakeAnalyticsPropertiesArgs) DataLakeAnalyticsPropertiesPtrInput {
	return (*dataLakeAnalyticsPropertiesPtrType)(v)
}

func (*dataLakeAnalyticsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeAnalyticsProperties)(nil)).Elem()
}

func (i *dataLakeAnalyticsPropertiesPtrType) ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput {
	return i.ToDataLakeAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (i *dataLakeAnalyticsPropertiesPtrType) ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsPropertiesPtrOutput)
}

type DataLakeAnalyticsPropertiesOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsProperties)(nil)).Elem()
}

func (o DataLakeAnalyticsPropertiesOutput) ToDataLakeAnalyticsPropertiesOutput() DataLakeAnalyticsPropertiesOutput {
	return o
}

func (o DataLakeAnalyticsPropertiesOutput) ToDataLakeAnalyticsPropertiesOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesOutput {
	return o
}

func (o DataLakeAnalyticsPropertiesOutput) ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput {
	return o.ToDataLakeAnalyticsPropertiesPtrOutputWithContext(context.Background())
}

func (o DataLakeAnalyticsPropertiesOutput) ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataLakeAnalyticsProperties) *DataLakeAnalyticsProperties {
		return &v
	}).(DataLakeAnalyticsPropertiesPtrOutput)
}

func (o DataLakeAnalyticsPropertiesOutput) DataLakeStoreAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsProperties) *string { return v.DataLakeStoreAccountName }).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeAnalyticsProperties)(nil)).Elem()
}

func (o DataLakeAnalyticsPropertiesPtrOutput) ToDataLakeAnalyticsPropertiesPtrOutput() DataLakeAnalyticsPropertiesPtrOutput {
	return o
}

func (o DataLakeAnalyticsPropertiesPtrOutput) ToDataLakeAnalyticsPropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsPropertiesPtrOutput {
	return o
}

func (o DataLakeAnalyticsPropertiesPtrOutput) Elem() DataLakeAnalyticsPropertiesOutput {
	return o.ApplyT(func(v *DataLakeAnalyticsProperties) DataLakeAnalyticsProperties {
		if v != nil {
			return *v
		}
		var ret DataLakeAnalyticsProperties
		return ret
	}).(DataLakeAnalyticsPropertiesOutput)
}

func (o DataLakeAnalyticsPropertiesPtrOutput) DataLakeStoreAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeAnalyticsProperties) *string {
		if v == nil {
			return nil
		}
		return v.DataLakeStoreAccountName
	}).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *DataLakeAnalyticsResponseProperties  `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}





type DataLakeAnalyticsResponseInput interface {
	pulumi.Input

	ToDataLakeAnalyticsResponseOutput() DataLakeAnalyticsResponseOutput
	ToDataLakeAnalyticsResponseOutputWithContext(context.Context) DataLakeAnalyticsResponseOutput
}

type DataLakeAnalyticsResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         DataLakeAnalyticsResponsePropertiesPtrInput   `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (DataLakeAnalyticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsResponse)(nil)).Elem()
}

func (i DataLakeAnalyticsResponseArgs) ToDataLakeAnalyticsResponseOutput() DataLakeAnalyticsResponseOutput {
	return i.ToDataLakeAnalyticsResponseOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsResponseArgs) ToDataLakeAnalyticsResponseOutputWithContext(ctx context.Context) DataLakeAnalyticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsResponseOutput)
}

type DataLakeAnalyticsResponseOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsResponse)(nil)).Elem()
}

func (o DataLakeAnalyticsResponseOutput) ToDataLakeAnalyticsResponseOutput() DataLakeAnalyticsResponseOutput {
	return o
}

func (o DataLakeAnalyticsResponseOutput) ToDataLakeAnalyticsResponseOutputWithContext(ctx context.Context) DataLakeAnalyticsResponseOutput {
	return o
}

func (o DataLakeAnalyticsResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DataLakeAnalyticsResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataLakeAnalyticsResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o DataLakeAnalyticsResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsResponseOutput) Properties() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *DataLakeAnalyticsResponseProperties { return v.Properties }).(DataLakeAnalyticsResponsePropertiesPtrOutput)
}

func (o DataLakeAnalyticsResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o DataLakeAnalyticsResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataLakeAnalyticsResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsResponseProperties struct {
	DataLakeStoreAccountName *string `pulumi:"dataLakeStoreAccountName"`
}





type DataLakeAnalyticsResponsePropertiesInput interface {
	pulumi.Input

	ToDataLakeAnalyticsResponsePropertiesOutput() DataLakeAnalyticsResponsePropertiesOutput
	ToDataLakeAnalyticsResponsePropertiesOutputWithContext(context.Context) DataLakeAnalyticsResponsePropertiesOutput
}

type DataLakeAnalyticsResponsePropertiesArgs struct {
	DataLakeStoreAccountName pulumi.StringPtrInput `pulumi:"dataLakeStoreAccountName"`
}

func (DataLakeAnalyticsResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsResponseProperties)(nil)).Elem()
}

func (i DataLakeAnalyticsResponsePropertiesArgs) ToDataLakeAnalyticsResponsePropertiesOutput() DataLakeAnalyticsResponsePropertiesOutput {
	return i.ToDataLakeAnalyticsResponsePropertiesOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsResponsePropertiesArgs) ToDataLakeAnalyticsResponsePropertiesOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsResponsePropertiesOutput)
}

func (i DataLakeAnalyticsResponsePropertiesArgs) ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return i.ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i DataLakeAnalyticsResponsePropertiesArgs) ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsResponsePropertiesOutput).ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx)
}









type DataLakeAnalyticsResponsePropertiesPtrInput interface {
	pulumi.Input

	ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput
	ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput
}

type dataLakeAnalyticsResponsePropertiesPtrType DataLakeAnalyticsResponsePropertiesArgs

func DataLakeAnalyticsResponsePropertiesPtr(v *DataLakeAnalyticsResponsePropertiesArgs) DataLakeAnalyticsResponsePropertiesPtrInput {
	return (*dataLakeAnalyticsResponsePropertiesPtrType)(v)
}

func (*dataLakeAnalyticsResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeAnalyticsResponseProperties)(nil)).Elem()
}

func (i *dataLakeAnalyticsResponsePropertiesPtrType) ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return i.ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *dataLakeAnalyticsResponsePropertiesPtrType) ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeAnalyticsResponsePropertiesPtrOutput)
}

type DataLakeAnalyticsResponsePropertiesOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLakeAnalyticsResponseProperties)(nil)).Elem()
}

func (o DataLakeAnalyticsResponsePropertiesOutput) ToDataLakeAnalyticsResponsePropertiesOutput() DataLakeAnalyticsResponsePropertiesOutput {
	return o
}

func (o DataLakeAnalyticsResponsePropertiesOutput) ToDataLakeAnalyticsResponsePropertiesOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesOutput {
	return o
}

func (o DataLakeAnalyticsResponsePropertiesOutput) ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o.ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o DataLakeAnalyticsResponsePropertiesOutput) ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataLakeAnalyticsResponseProperties) *DataLakeAnalyticsResponseProperties {
		return &v
	}).(DataLakeAnalyticsResponsePropertiesPtrOutput)
}

func (o DataLakeAnalyticsResponsePropertiesOutput) DataLakeStoreAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponseProperties) *string { return v.DataLakeStoreAccountName }).(pulumi.StringPtrOutput)
}

type DataLakeAnalyticsResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DataLakeAnalyticsResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeAnalyticsResponseProperties)(nil)).Elem()
}

func (o DataLakeAnalyticsResponsePropertiesPtrOutput) ToDataLakeAnalyticsResponsePropertiesPtrOutput() DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o
}

func (o DataLakeAnalyticsResponsePropertiesPtrOutput) ToDataLakeAnalyticsResponsePropertiesPtrOutputWithContext(ctx context.Context) DataLakeAnalyticsResponsePropertiesPtrOutput {
	return o
}

func (o DataLakeAnalyticsResponsePropertiesPtrOutput) Elem() DataLakeAnalyticsResponsePropertiesOutput {
	return o.ApplyT(func(v *DataLakeAnalyticsResponseProperties) DataLakeAnalyticsResponseProperties {
		if v != nil {
			return *v
		}
		var ret DataLakeAnalyticsResponseProperties
		return ret
	}).(DataLakeAnalyticsResponsePropertiesOutput)
}

func (o DataLakeAnalyticsResponsePropertiesPtrOutput) DataLakeStoreAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeAnalyticsResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DataLakeStoreAccountName
	}).(pulumi.StringPtrOutput)
}

type Databricks struct {
	ComputeLocation *string               `pulumi:"computeLocation"`
	ComputeType     string                `pulumi:"computeType"`
	Description     *string               `pulumi:"description"`
	Properties      *DatabricksProperties `pulumi:"properties"`
	ResourceId      *string               `pulumi:"resourceId"`
}





type DatabricksInput interface {
	pulumi.Input

	ToDatabricksOutput() DatabricksOutput
	ToDatabricksOutputWithContext(context.Context) DatabricksOutput
}

type DatabricksArgs struct {
	ComputeLocation pulumi.StringPtrInput        `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput           `pulumi:"computeType"`
	Description     pulumi.StringPtrInput        `pulumi:"description"`
	Properties      DatabricksPropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput        `pulumi:"resourceId"`
}

func (DatabricksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Databricks)(nil)).Elem()
}

func (i DatabricksArgs) ToDatabricksOutput() DatabricksOutput {
	return i.ToDatabricksOutputWithContext(context.Background())
}

func (i DatabricksArgs) ToDatabricksOutputWithContext(ctx context.Context) DatabricksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksOutput)
}

type DatabricksOutput struct{ *pulumi.OutputState }

func (DatabricksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Databricks)(nil)).Elem()
}

func (o DatabricksOutput) ToDatabricksOutput() DatabricksOutput {
	return o
}

func (o DatabricksOutput) ToDatabricksOutputWithContext(ctx context.Context) DatabricksOutput {
	return o
}

func (o DatabricksOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Databricks) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DatabricksOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v Databricks) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DatabricksOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Databricks) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatabricksOutput) Properties() DatabricksPropertiesPtrOutput {
	return o.ApplyT(func(v Databricks) *DatabricksProperties { return v.Properties }).(DatabricksPropertiesPtrOutput)
}

func (o DatabricksOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Databricks) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DatabricksProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
}





type DatabricksPropertiesInput interface {
	pulumi.Input

	ToDatabricksPropertiesOutput() DatabricksPropertiesOutput
	ToDatabricksPropertiesOutputWithContext(context.Context) DatabricksPropertiesOutput
}

type DatabricksPropertiesArgs struct {
	DatabricksAccessToken pulumi.StringPtrInput `pulumi:"databricksAccessToken"`
}

func (DatabricksPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksProperties)(nil)).Elem()
}

func (i DatabricksPropertiesArgs) ToDatabricksPropertiesOutput() DatabricksPropertiesOutput {
	return i.ToDatabricksPropertiesOutputWithContext(context.Background())
}

func (i DatabricksPropertiesArgs) ToDatabricksPropertiesOutputWithContext(ctx context.Context) DatabricksPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesOutput)
}

func (i DatabricksPropertiesArgs) ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput {
	return i.ToDatabricksPropertiesPtrOutputWithContext(context.Background())
}

func (i DatabricksPropertiesArgs) ToDatabricksPropertiesPtrOutputWithContext(ctx context.Context) DatabricksPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesOutput).ToDatabricksPropertiesPtrOutputWithContext(ctx)
}









type DatabricksPropertiesPtrInput interface {
	pulumi.Input

	ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput
	ToDatabricksPropertiesPtrOutputWithContext(context.Context) DatabricksPropertiesPtrOutput
}

type databricksPropertiesPtrType DatabricksPropertiesArgs

func DatabricksPropertiesPtr(v *DatabricksPropertiesArgs) DatabricksPropertiesPtrInput {
	return (*databricksPropertiesPtrType)(v)
}

func (*databricksPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksProperties)(nil)).Elem()
}

func (i *databricksPropertiesPtrType) ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput {
	return i.ToDatabricksPropertiesPtrOutputWithContext(context.Background())
}

func (i *databricksPropertiesPtrType) ToDatabricksPropertiesPtrOutputWithContext(ctx context.Context) DatabricksPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesPtrOutput)
}

type DatabricksPropertiesOutput struct{ *pulumi.OutputState }

func (DatabricksPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksProperties)(nil)).Elem()
}

func (o DatabricksPropertiesOutput) ToDatabricksPropertiesOutput() DatabricksPropertiesOutput {
	return o
}

func (o DatabricksPropertiesOutput) ToDatabricksPropertiesOutputWithContext(ctx context.Context) DatabricksPropertiesOutput {
	return o
}

func (o DatabricksPropertiesOutput) ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput {
	return o.ToDatabricksPropertiesPtrOutputWithContext(context.Background())
}

func (o DatabricksPropertiesOutput) ToDatabricksPropertiesPtrOutputWithContext(ctx context.Context) DatabricksPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabricksProperties) *DatabricksProperties {
		return &v
	}).(DatabricksPropertiesPtrOutput)
}

func (o DatabricksPropertiesOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksProperties) *string { return v.DatabricksAccessToken }).(pulumi.StringPtrOutput)
}

type DatabricksPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DatabricksPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksProperties)(nil)).Elem()
}

func (o DatabricksPropertiesPtrOutput) ToDatabricksPropertiesPtrOutput() DatabricksPropertiesPtrOutput {
	return o
}

func (o DatabricksPropertiesPtrOutput) ToDatabricksPropertiesPtrOutputWithContext(ctx context.Context) DatabricksPropertiesPtrOutput {
	return o
}

func (o DatabricksPropertiesPtrOutput) Elem() DatabricksPropertiesOutput {
	return o.ApplyT(func(v *DatabricksProperties) DatabricksProperties {
		if v != nil {
			return *v
		}
		var ret DatabricksProperties
		return ret
	}).(DatabricksPropertiesOutput)
}

func (o DatabricksPropertiesPtrOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksProperties) *string {
		if v == nil {
			return nil
		}
		return v.DatabricksAccessToken
	}).(pulumi.StringPtrOutput)
}

type DatabricksResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *DatabricksResponseProperties         `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}





type DatabricksResponseInput interface {
	pulumi.Input

	ToDatabricksResponseOutput() DatabricksResponseOutput
	ToDatabricksResponseOutputWithContext(context.Context) DatabricksResponseOutput
}

type DatabricksResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         DatabricksResponsePropertiesPtrInput          `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (DatabricksResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksResponse)(nil)).Elem()
}

func (i DatabricksResponseArgs) ToDatabricksResponseOutput() DatabricksResponseOutput {
	return i.ToDatabricksResponseOutputWithContext(context.Background())
}

func (i DatabricksResponseArgs) ToDatabricksResponseOutputWithContext(ctx context.Context) DatabricksResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksResponseOutput)
}

type DatabricksResponseOutput struct{ *pulumi.OutputState }

func (DatabricksResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksResponse)(nil)).Elem()
}

func (o DatabricksResponseOutput) ToDatabricksResponseOutput() DatabricksResponseOutput {
	return o
}

func (o DatabricksResponseOutput) ToDatabricksResponseOutputWithContext(ctx context.Context) DatabricksResponseOutput {
	return o
}

func (o DatabricksResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o DatabricksResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DatabricksResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v DatabricksResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o DatabricksResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) Properties() DatabricksResponsePropertiesPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *DatabricksResponseProperties { return v.Properties }).(DatabricksResponsePropertiesPtrOutput)
}

func (o DatabricksResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v DatabricksResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o DatabricksResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DatabricksResponseProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
}





type DatabricksResponsePropertiesInput interface {
	pulumi.Input

	ToDatabricksResponsePropertiesOutput() DatabricksResponsePropertiesOutput
	ToDatabricksResponsePropertiesOutputWithContext(context.Context) DatabricksResponsePropertiesOutput
}

type DatabricksResponsePropertiesArgs struct {
	DatabricksAccessToken pulumi.StringPtrInput `pulumi:"databricksAccessToken"`
}

func (DatabricksResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksResponseProperties)(nil)).Elem()
}

func (i DatabricksResponsePropertiesArgs) ToDatabricksResponsePropertiesOutput() DatabricksResponsePropertiesOutput {
	return i.ToDatabricksResponsePropertiesOutputWithContext(context.Background())
}

func (i DatabricksResponsePropertiesArgs) ToDatabricksResponsePropertiesOutputWithContext(ctx context.Context) DatabricksResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksResponsePropertiesOutput)
}

func (i DatabricksResponsePropertiesArgs) ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput {
	return i.ToDatabricksResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i DatabricksResponsePropertiesArgs) ToDatabricksResponsePropertiesPtrOutputWithContext(ctx context.Context) DatabricksResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksResponsePropertiesOutput).ToDatabricksResponsePropertiesPtrOutputWithContext(ctx)
}









type DatabricksResponsePropertiesPtrInput interface {
	pulumi.Input

	ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput
	ToDatabricksResponsePropertiesPtrOutputWithContext(context.Context) DatabricksResponsePropertiesPtrOutput
}

type databricksResponsePropertiesPtrType DatabricksResponsePropertiesArgs

func DatabricksResponsePropertiesPtr(v *DatabricksResponsePropertiesArgs) DatabricksResponsePropertiesPtrInput {
	return (*databricksResponsePropertiesPtrType)(v)
}

func (*databricksResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksResponseProperties)(nil)).Elem()
}

func (i *databricksResponsePropertiesPtrType) ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput {
	return i.ToDatabricksResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *databricksResponsePropertiesPtrType) ToDatabricksResponsePropertiesPtrOutputWithContext(ctx context.Context) DatabricksResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksResponsePropertiesPtrOutput)
}

type DatabricksResponsePropertiesOutput struct{ *pulumi.OutputState }

func (DatabricksResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksResponseProperties)(nil)).Elem()
}

func (o DatabricksResponsePropertiesOutput) ToDatabricksResponsePropertiesOutput() DatabricksResponsePropertiesOutput {
	return o
}

func (o DatabricksResponsePropertiesOutput) ToDatabricksResponsePropertiesOutputWithContext(ctx context.Context) DatabricksResponsePropertiesOutput {
	return o
}

func (o DatabricksResponsePropertiesOutput) ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput {
	return o.ToDatabricksResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o DatabricksResponsePropertiesOutput) ToDatabricksResponsePropertiesPtrOutputWithContext(ctx context.Context) DatabricksResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabricksResponseProperties) *DatabricksResponseProperties {
		return &v
	}).(DatabricksResponsePropertiesPtrOutput)
}

func (o DatabricksResponsePropertiesOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponseProperties) *string { return v.DatabricksAccessToken }).(pulumi.StringPtrOutput)
}

type DatabricksResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DatabricksResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksResponseProperties)(nil)).Elem()
}

func (o DatabricksResponsePropertiesPtrOutput) ToDatabricksResponsePropertiesPtrOutput() DatabricksResponsePropertiesPtrOutput {
	return o
}

func (o DatabricksResponsePropertiesPtrOutput) ToDatabricksResponsePropertiesPtrOutputWithContext(ctx context.Context) DatabricksResponsePropertiesPtrOutput {
	return o
}

func (o DatabricksResponsePropertiesPtrOutput) Elem() DatabricksResponsePropertiesOutput {
	return o.ApplyT(func(v *DatabricksResponseProperties) DatabricksResponseProperties {
		if v != nil {
			return *v
		}
		var ret DatabricksResponseProperties
		return ret
	}).(DatabricksResponsePropertiesOutput)
}

func (o DatabricksResponsePropertiesPtrOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DatabricksAccessToken
	}).(pulumi.StringPtrOutput)
}

type EncryptionProperty struct {
	KeyVaultProperties KeyVaultProperties `pulumi:"keyVaultProperties"`
	Status             string             `pulumi:"status"`
}





type EncryptionPropertyInput interface {
	pulumi.Input

	ToEncryptionPropertyOutput() EncryptionPropertyOutput
	ToEncryptionPropertyOutputWithContext(context.Context) EncryptionPropertyOutput
}

type EncryptionPropertyArgs struct {
	KeyVaultProperties KeyVaultPropertiesInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringInput      `pulumi:"status"`
}

func (EncryptionPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperty)(nil)).Elem()
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyOutput() EncryptionPropertyOutput {
	return i.ToEncryptionPropertyOutputWithContext(context.Background())
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyOutputWithContext(ctx context.Context) EncryptionPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyOutput)
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return i.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyOutput).ToEncryptionPropertyPtrOutputWithContext(ctx)
}









type EncryptionPropertyPtrInput interface {
	pulumi.Input

	ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput
	ToEncryptionPropertyPtrOutputWithContext(context.Context) EncryptionPropertyPtrOutput
}

type encryptionPropertyPtrType EncryptionPropertyArgs

func EncryptionPropertyPtr(v *EncryptionPropertyArgs) EncryptionPropertyPtrInput {
	return (*encryptionPropertyPtrType)(v)
}

func (*encryptionPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperty)(nil)).Elem()
}

func (i *encryptionPropertyPtrType) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return i.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertyPtrType) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyPtrOutput)
}

type EncryptionPropertyOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperty)(nil)).Elem()
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyOutput() EncryptionPropertyOutput {
	return o
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyOutputWithContext(ctx context.Context) EncryptionPropertyOutput {
	return o
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return o.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionProperty) *EncryptionProperty {
		return &v
	}).(EncryptionPropertyPtrOutput)
}

func (o EncryptionPropertyOutput) KeyVaultProperties() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v EncryptionProperty) KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesOutput)
}

func (o EncryptionPropertyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionProperty) string { return v.Status }).(pulumi.StringOutput)
}

type EncryptionPropertyPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperty)(nil)).Elem()
}

func (o EncryptionPropertyPtrOutput) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return o
}

func (o EncryptionPropertyPtrOutput) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return o
}

func (o EncryptionPropertyPtrOutput) Elem() EncryptionPropertyOutput {
	return o.ApplyT(func(v *EncryptionProperty) EncryptionProperty {
		if v != nil {
			return *v
		}
		var ret EncryptionProperty
		return ret
	}).(EncryptionPropertyOutput)
}

func (o EncryptionPropertyPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return &v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionPropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type EncryptionPropertyResponse struct {
	KeyVaultProperties KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             string                     `pulumi:"status"`
}





type EncryptionPropertyResponseInput interface {
	pulumi.Input

	ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput
	ToEncryptionPropertyResponseOutputWithContext(context.Context) EncryptionPropertyResponseOutput
}

type EncryptionPropertyResponseArgs struct {
	KeyVaultProperties KeyVaultPropertiesResponseInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringInput              `pulumi:"status"`
}

func (EncryptionPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertyResponse)(nil)).Elem()
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput {
	return i.ToEncryptionPropertyResponseOutputWithContext(context.Background())
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponseOutputWithContext(ctx context.Context) EncryptionPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponseOutput)
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return i.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponseOutput).ToEncryptionPropertyResponsePtrOutputWithContext(ctx)
}









type EncryptionPropertyResponsePtrInput interface {
	pulumi.Input

	ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput
	ToEncryptionPropertyResponsePtrOutputWithContext(context.Context) EncryptionPropertyResponsePtrOutput
}

type encryptionPropertyResponsePtrType EncryptionPropertyResponseArgs

func EncryptionPropertyResponsePtr(v *EncryptionPropertyResponseArgs) EncryptionPropertyResponsePtrInput {
	return (*encryptionPropertyResponsePtrType)(v)
}

func (*encryptionPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertyResponse)(nil)).Elem()
}

func (i *encryptionPropertyResponsePtrType) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return i.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionPropertyResponsePtrType) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponsePtrOutput)
}

type EncryptionPropertyResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertyResponse)(nil)).Elem()
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput {
	return o
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponseOutputWithContext(ctx context.Context) EncryptionPropertyResponseOutput {
	return o
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return o.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPropertyResponse) *EncryptionPropertyResponse {
		return &v
	}).(EncryptionPropertyResponsePtrOutput)
}

func (o EncryptionPropertyResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponseOutput)
}

func (o EncryptionPropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type EncryptionPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertyResponse)(nil)).Elem()
}

func (o EncryptionPropertyResponsePtrOutput) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return o
}

func (o EncryptionPropertyResponsePtrOutput) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return o
}

func (o EncryptionPropertyResponsePtrOutput) Elem() EncryptionPropertyResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) EncryptionPropertyResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertyResponse
		return ret
	}).(EncryptionPropertyResponseOutput)
}

func (o EncryptionPropertyResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionPropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ErrorDetailResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	Code    pulumi.StringInput `pulumi:"code"`
	Message pulumi.StringInput `pulumi:"message"`
}

func (ErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return i.ToErrorDetailResponseOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput)
}





type ErrorDetailResponseArrayInput interface {
	pulumi.Input

	ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput
	ToErrorDetailResponseArrayOutputWithContext(context.Context) ErrorDetailResponseArrayOutput
}

type ErrorDetailResponseArray []ErrorDetailResponseInput

func (ErrorDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return i.ToErrorDetailResponseArrayOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseArrayOutput)
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type ErrorResponseResponse struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
}





type ErrorResponseResponseInput interface {
	pulumi.Input

	ToErrorResponseResponseOutput() ErrorResponseResponseOutput
	ToErrorResponseResponseOutputWithContext(context.Context) ErrorResponseResponseOutput
}

type ErrorResponseResponseArgs struct {
	Code    pulumi.StringInput            `pulumi:"code"`
	Details ErrorDetailResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput            `pulumi:"message"`
}

func (ErrorResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (i ErrorResponseResponseArgs) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return i.ToErrorResponseResponseOutputWithContext(context.Background())
}

func (i ErrorResponseResponseArgs) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseResponseOutput)
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

func (o ErrorResponseResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorResponseResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponseResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorResponseResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Message }).(pulumi.StringOutput)
}

type HDInsight struct {
	ComputeLocation *string              `pulumi:"computeLocation"`
	ComputeType     string               `pulumi:"computeType"`
	Description     *string              `pulumi:"description"`
	Properties      *HDInsightProperties `pulumi:"properties"`
	ResourceId      *string              `pulumi:"resourceId"`
}





type HDInsightInput interface {
	pulumi.Input

	ToHDInsightOutput() HDInsightOutput
	ToHDInsightOutputWithContext(context.Context) HDInsightOutput
}

type HDInsightArgs struct {
	ComputeLocation pulumi.StringPtrInput       `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput          `pulumi:"computeType"`
	Description     pulumi.StringPtrInput       `pulumi:"description"`
	Properties      HDInsightPropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput       `pulumi:"resourceId"`
}

func (HDInsightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsight)(nil)).Elem()
}

func (i HDInsightArgs) ToHDInsightOutput() HDInsightOutput {
	return i.ToHDInsightOutputWithContext(context.Background())
}

func (i HDInsightArgs) ToHDInsightOutputWithContext(ctx context.Context) HDInsightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightOutput)
}

type HDInsightOutput struct{ *pulumi.OutputState }

func (HDInsightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsight)(nil)).Elem()
}

func (o HDInsightOutput) ToHDInsightOutput() HDInsightOutput {
	return o
}

func (o HDInsightOutput) ToHDInsightOutputWithContext(ctx context.Context) HDInsightOutput {
	return o
}

func (o HDInsightOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsight) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o HDInsightOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsight) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o HDInsightOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsight) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o HDInsightOutput) Properties() HDInsightPropertiesPtrOutput {
	return o.ApplyT(func(v HDInsight) *HDInsightProperties { return v.Properties }).(HDInsightPropertiesPtrOutput)
}

func (o HDInsightOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsight) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type HDInsightProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
}





type HDInsightPropertiesInput interface {
	pulumi.Input

	ToHDInsightPropertiesOutput() HDInsightPropertiesOutput
	ToHDInsightPropertiesOutputWithContext(context.Context) HDInsightPropertiesOutput
}

type HDInsightPropertiesArgs struct {
	Address              pulumi.StringPtrInput                `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsPtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                   `pulumi:"sshPort"`
}

func (HDInsightPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightProperties)(nil)).Elem()
}

func (i HDInsightPropertiesArgs) ToHDInsightPropertiesOutput() HDInsightPropertiesOutput {
	return i.ToHDInsightPropertiesOutputWithContext(context.Background())
}

func (i HDInsightPropertiesArgs) ToHDInsightPropertiesOutputWithContext(ctx context.Context) HDInsightPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesOutput)
}

func (i HDInsightPropertiesArgs) ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput {
	return i.ToHDInsightPropertiesPtrOutputWithContext(context.Background())
}

func (i HDInsightPropertiesArgs) ToHDInsightPropertiesPtrOutputWithContext(ctx context.Context) HDInsightPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesOutput).ToHDInsightPropertiesPtrOutputWithContext(ctx)
}









type HDInsightPropertiesPtrInput interface {
	pulumi.Input

	ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput
	ToHDInsightPropertiesPtrOutputWithContext(context.Context) HDInsightPropertiesPtrOutput
}

type hdinsightPropertiesPtrType HDInsightPropertiesArgs

func HDInsightPropertiesPtr(v *HDInsightPropertiesArgs) HDInsightPropertiesPtrInput {
	return (*hdinsightPropertiesPtrType)(v)
}

func (*hdinsightPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightProperties)(nil)).Elem()
}

func (i *hdinsightPropertiesPtrType) ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput {
	return i.ToHDInsightPropertiesPtrOutputWithContext(context.Background())
}

func (i *hdinsightPropertiesPtrType) ToHDInsightPropertiesPtrOutputWithContext(ctx context.Context) HDInsightPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesPtrOutput)
}

type HDInsightPropertiesOutput struct{ *pulumi.OutputState }

func (HDInsightPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightProperties)(nil)).Elem()
}

func (o HDInsightPropertiesOutput) ToHDInsightPropertiesOutput() HDInsightPropertiesOutput {
	return o
}

func (o HDInsightPropertiesOutput) ToHDInsightPropertiesOutputWithContext(ctx context.Context) HDInsightPropertiesOutput {
	return o
}

func (o HDInsightPropertiesOutput) ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput {
	return o.ToHDInsightPropertiesPtrOutputWithContext(context.Background())
}

func (o HDInsightPropertiesOutput) ToHDInsightPropertiesPtrOutputWithContext(ctx context.Context) HDInsightPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HDInsightProperties) *HDInsightProperties {
		return &v
	}).(HDInsightPropertiesPtrOutput)
}

func (o HDInsightPropertiesOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightProperties) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o HDInsightPropertiesOutput) AdministratorAccount() VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyT(func(v HDInsightProperties) *VirtualMachineSshCredentials { return v.AdministratorAccount }).(VirtualMachineSshCredentialsPtrOutput)
}

func (o HDInsightPropertiesOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HDInsightProperties) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

type HDInsightPropertiesPtrOutput struct{ *pulumi.OutputState }

func (HDInsightPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightProperties)(nil)).Elem()
}

func (o HDInsightPropertiesPtrOutput) ToHDInsightPropertiesPtrOutput() HDInsightPropertiesPtrOutput {
	return o
}

func (o HDInsightPropertiesPtrOutput) ToHDInsightPropertiesPtrOutputWithContext(ctx context.Context) HDInsightPropertiesPtrOutput {
	return o
}

func (o HDInsightPropertiesPtrOutput) Elem() HDInsightPropertiesOutput {
	return o.ApplyT(func(v *HDInsightProperties) HDInsightProperties {
		if v != nil {
			return *v
		}
		var ret HDInsightProperties
		return ret
	}).(HDInsightPropertiesOutput)
}

func (o HDInsightPropertiesPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HDInsightProperties) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o HDInsightPropertiesPtrOutput) AdministratorAccount() VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyT(func(v *HDInsightProperties) *VirtualMachineSshCredentials {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsPtrOutput)
}

func (o HDInsightPropertiesPtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HDInsightProperties) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
}

type HDInsightResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *HDInsightResponseProperties          `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}





type HDInsightResponseInput interface {
	pulumi.Input

	ToHDInsightResponseOutput() HDInsightResponseOutput
	ToHDInsightResponseOutputWithContext(context.Context) HDInsightResponseOutput
}

type HDInsightResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         HDInsightResponsePropertiesPtrInput           `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (HDInsightResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightResponse)(nil)).Elem()
}

func (i HDInsightResponseArgs) ToHDInsightResponseOutput() HDInsightResponseOutput {
	return i.ToHDInsightResponseOutputWithContext(context.Background())
}

func (i HDInsightResponseArgs) ToHDInsightResponseOutputWithContext(ctx context.Context) HDInsightResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightResponseOutput)
}

type HDInsightResponseOutput struct{ *pulumi.OutputState }

func (HDInsightResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightResponse)(nil)).Elem()
}

func (o HDInsightResponseOutput) ToHDInsightResponseOutput() HDInsightResponseOutput {
	return o
}

func (o HDInsightResponseOutput) ToHDInsightResponseOutputWithContext(ctx context.Context) HDInsightResponseOutput {
	return o
}

func (o HDInsightResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o HDInsightResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o HDInsightResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v HDInsightResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o HDInsightResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) Properties() HDInsightResponsePropertiesPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *HDInsightResponseProperties { return v.Properties }).(HDInsightResponsePropertiesPtrOutput)
}

func (o HDInsightResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v HDInsightResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o HDInsightResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type HDInsightResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
}





type HDInsightResponsePropertiesInput interface {
	pulumi.Input

	ToHDInsightResponsePropertiesOutput() HDInsightResponsePropertiesOutput
	ToHDInsightResponsePropertiesOutputWithContext(context.Context) HDInsightResponsePropertiesOutput
}

type HDInsightResponsePropertiesArgs struct {
	Address              pulumi.StringPtrInput                        `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsResponsePtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                           `pulumi:"sshPort"`
}

func (HDInsightResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightResponseProperties)(nil)).Elem()
}

func (i HDInsightResponsePropertiesArgs) ToHDInsightResponsePropertiesOutput() HDInsightResponsePropertiesOutput {
	return i.ToHDInsightResponsePropertiesOutputWithContext(context.Background())
}

func (i HDInsightResponsePropertiesArgs) ToHDInsightResponsePropertiesOutputWithContext(ctx context.Context) HDInsightResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightResponsePropertiesOutput)
}

func (i HDInsightResponsePropertiesArgs) ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput {
	return i.ToHDInsightResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i HDInsightResponsePropertiesArgs) ToHDInsightResponsePropertiesPtrOutputWithContext(ctx context.Context) HDInsightResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightResponsePropertiesOutput).ToHDInsightResponsePropertiesPtrOutputWithContext(ctx)
}









type HDInsightResponsePropertiesPtrInput interface {
	pulumi.Input

	ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput
	ToHDInsightResponsePropertiesPtrOutputWithContext(context.Context) HDInsightResponsePropertiesPtrOutput
}

type hdinsightResponsePropertiesPtrType HDInsightResponsePropertiesArgs

func HDInsightResponsePropertiesPtr(v *HDInsightResponsePropertiesArgs) HDInsightResponsePropertiesPtrInput {
	return (*hdinsightResponsePropertiesPtrType)(v)
}

func (*hdinsightResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightResponseProperties)(nil)).Elem()
}

func (i *hdinsightResponsePropertiesPtrType) ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput {
	return i.ToHDInsightResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *hdinsightResponsePropertiesPtrType) ToHDInsightResponsePropertiesPtrOutputWithContext(ctx context.Context) HDInsightResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightResponsePropertiesPtrOutput)
}

type HDInsightResponsePropertiesOutput struct{ *pulumi.OutputState }

func (HDInsightResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightResponseProperties)(nil)).Elem()
}

func (o HDInsightResponsePropertiesOutput) ToHDInsightResponsePropertiesOutput() HDInsightResponsePropertiesOutput {
	return o
}

func (o HDInsightResponsePropertiesOutput) ToHDInsightResponsePropertiesOutputWithContext(ctx context.Context) HDInsightResponsePropertiesOutput {
	return o
}

func (o HDInsightResponsePropertiesOutput) ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput {
	return o.ToHDInsightResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o HDInsightResponsePropertiesOutput) ToHDInsightResponsePropertiesPtrOutputWithContext(ctx context.Context) HDInsightResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HDInsightResponseProperties) *HDInsightResponseProperties {
		return &v
	}).(HDInsightResponsePropertiesPtrOutput)
}

func (o HDInsightResponsePropertiesOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponseProperties) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o HDInsightResponsePropertiesOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v HDInsightResponseProperties) *VirtualMachineSshCredentialsResponse {
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o HDInsightResponsePropertiesOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HDInsightResponseProperties) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

type HDInsightResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (HDInsightResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightResponseProperties)(nil)).Elem()
}

func (o HDInsightResponsePropertiesPtrOutput) ToHDInsightResponsePropertiesPtrOutput() HDInsightResponsePropertiesPtrOutput {
	return o
}

func (o HDInsightResponsePropertiesPtrOutput) ToHDInsightResponsePropertiesPtrOutputWithContext(ctx context.Context) HDInsightResponsePropertiesPtrOutput {
	return o
}

func (o HDInsightResponsePropertiesPtrOutput) Elem() HDInsightResponsePropertiesOutput {
	return o.ApplyT(func(v *HDInsightResponseProperties) HDInsightResponseProperties {
		if v != nil {
			return *v
		}
		var ret HDInsightResponseProperties
		return ret
	}).(HDInsightResponsePropertiesOutput)
}

func (o HDInsightResponsePropertiesPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HDInsightResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o HDInsightResponsePropertiesPtrOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *HDInsightResponseProperties) *VirtualMachineSshCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o HDInsightResponsePropertiesPtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HDInsightResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
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





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
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

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
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

type KeyVaultProperties struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	KeyIdentifier    string  `pulumi:"keyIdentifier"`
	KeyVaultArmId    string  `pulumi:"keyVaultArmId"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	KeyIdentifier    pulumi.StringInput    `pulumi:"keyIdentifier"`
	KeyVaultArmId    pulumi.StringInput    `pulumi:"keyVaultArmId"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultProperties) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultArmId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultProperties) string { return v.KeyVaultArmId }).(pulumi.StringOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultArmId
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	KeyIdentifier    string  `pulumi:"keyIdentifier"`
	KeyVaultArmId    string  `pulumi:"keyVaultArmId"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	KeyIdentifier    pulumi.StringInput    `pulumi:"keyIdentifier"`
	KeyVaultArmId    pulumi.StringInput    `pulumi:"keyVaultArmId"`
}

func (KeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return i.ToKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput)
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput).ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type KeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput
	ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) KeyVaultPropertiesResponsePtrOutput
}

type keyVaultPropertiesResponsePtrType KeyVaultPropertiesResponseArgs

func KeyVaultPropertiesResponsePtr(v *KeyVaultPropertiesResponseArgs) KeyVaultPropertiesResponsePtrInput {
	return (*keyVaultPropertiesResponsePtrType)(v)
}

func (*keyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponsePtrOutput)
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultPropertiesResponse) *KeyVaultPropertiesResponse {
		return &v
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultArmId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.KeyVaultArmId }).(pulumi.StringOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultArmId
	}).(pulumi.StringPtrOutput)
}

type MachineLearningServiceErrorResponse struct {
	Error ErrorResponseResponse `pulumi:"error"`
}





type MachineLearningServiceErrorResponseInput interface {
	pulumi.Input

	ToMachineLearningServiceErrorResponseOutput() MachineLearningServiceErrorResponseOutput
	ToMachineLearningServiceErrorResponseOutputWithContext(context.Context) MachineLearningServiceErrorResponseOutput
}

type MachineLearningServiceErrorResponseArgs struct {
	Error ErrorResponseResponseInput `pulumi:"error"`
}

func (MachineLearningServiceErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningServiceErrorResponse)(nil)).Elem()
}

func (i MachineLearningServiceErrorResponseArgs) ToMachineLearningServiceErrorResponseOutput() MachineLearningServiceErrorResponseOutput {
	return i.ToMachineLearningServiceErrorResponseOutputWithContext(context.Background())
}

func (i MachineLearningServiceErrorResponseArgs) ToMachineLearningServiceErrorResponseOutputWithContext(ctx context.Context) MachineLearningServiceErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningServiceErrorResponseOutput)
}





type MachineLearningServiceErrorResponseArrayInput interface {
	pulumi.Input

	ToMachineLearningServiceErrorResponseArrayOutput() MachineLearningServiceErrorResponseArrayOutput
	ToMachineLearningServiceErrorResponseArrayOutputWithContext(context.Context) MachineLearningServiceErrorResponseArrayOutput
}

type MachineLearningServiceErrorResponseArray []MachineLearningServiceErrorResponseInput

func (MachineLearningServiceErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineLearningServiceErrorResponse)(nil)).Elem()
}

func (i MachineLearningServiceErrorResponseArray) ToMachineLearningServiceErrorResponseArrayOutput() MachineLearningServiceErrorResponseArrayOutput {
	return i.ToMachineLearningServiceErrorResponseArrayOutputWithContext(context.Background())
}

func (i MachineLearningServiceErrorResponseArray) ToMachineLearningServiceErrorResponseArrayOutputWithContext(ctx context.Context) MachineLearningServiceErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningServiceErrorResponseArrayOutput)
}

type MachineLearningServiceErrorResponseOutput struct{ *pulumi.OutputState }

func (MachineLearningServiceErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningServiceErrorResponse)(nil)).Elem()
}

func (o MachineLearningServiceErrorResponseOutput) ToMachineLearningServiceErrorResponseOutput() MachineLearningServiceErrorResponseOutput {
	return o
}

func (o MachineLearningServiceErrorResponseOutput) ToMachineLearningServiceErrorResponseOutputWithContext(ctx context.Context) MachineLearningServiceErrorResponseOutput {
	return o
}

func (o MachineLearningServiceErrorResponseOutput) Error() ErrorResponseResponseOutput {
	return o.ApplyT(func(v MachineLearningServiceErrorResponse) ErrorResponseResponse { return v.Error }).(ErrorResponseResponseOutput)
}

type MachineLearningServiceErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (MachineLearningServiceErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineLearningServiceErrorResponse)(nil)).Elem()
}

func (o MachineLearningServiceErrorResponseArrayOutput) ToMachineLearningServiceErrorResponseArrayOutput() MachineLearningServiceErrorResponseArrayOutput {
	return o
}

func (o MachineLearningServiceErrorResponseArrayOutput) ToMachineLearningServiceErrorResponseArrayOutputWithContext(ctx context.Context) MachineLearningServiceErrorResponseArrayOutput {
	return o
}

func (o MachineLearningServiceErrorResponseArrayOutput) Index(i pulumi.IntInput) MachineLearningServiceErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineLearningServiceErrorResponse {
		return vs[0].([]MachineLearningServiceErrorResponse)[vs[1].(int)]
	}).(MachineLearningServiceErrorResponseOutput)
}

type NodeStateCountsResponse struct {
	IdleNodeCount      int `pulumi:"idleNodeCount"`
	LeavingNodeCount   int `pulumi:"leavingNodeCount"`
	PreemptedNodeCount int `pulumi:"preemptedNodeCount"`
	PreparingNodeCount int `pulumi:"preparingNodeCount"`
	RunningNodeCount   int `pulumi:"runningNodeCount"`
	UnusableNodeCount  int `pulumi:"unusableNodeCount"`
}





type NodeStateCountsResponseInput interface {
	pulumi.Input

	ToNodeStateCountsResponseOutput() NodeStateCountsResponseOutput
	ToNodeStateCountsResponseOutputWithContext(context.Context) NodeStateCountsResponseOutput
}

type NodeStateCountsResponseArgs struct {
	IdleNodeCount      pulumi.IntInput `pulumi:"idleNodeCount"`
	LeavingNodeCount   pulumi.IntInput `pulumi:"leavingNodeCount"`
	PreemptedNodeCount pulumi.IntInput `pulumi:"preemptedNodeCount"`
	PreparingNodeCount pulumi.IntInput `pulumi:"preparingNodeCount"`
	RunningNodeCount   pulumi.IntInput `pulumi:"runningNodeCount"`
	UnusableNodeCount  pulumi.IntInput `pulumi:"unusableNodeCount"`
}

func (NodeStateCountsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeStateCountsResponse)(nil)).Elem()
}

func (i NodeStateCountsResponseArgs) ToNodeStateCountsResponseOutput() NodeStateCountsResponseOutput {
	return i.ToNodeStateCountsResponseOutputWithContext(context.Background())
}

func (i NodeStateCountsResponseArgs) ToNodeStateCountsResponseOutputWithContext(ctx context.Context) NodeStateCountsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeStateCountsResponseOutput)
}

func (i NodeStateCountsResponseArgs) ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput {
	return i.ToNodeStateCountsResponsePtrOutputWithContext(context.Background())
}

func (i NodeStateCountsResponseArgs) ToNodeStateCountsResponsePtrOutputWithContext(ctx context.Context) NodeStateCountsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeStateCountsResponseOutput).ToNodeStateCountsResponsePtrOutputWithContext(ctx)
}









type NodeStateCountsResponsePtrInput interface {
	pulumi.Input

	ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput
	ToNodeStateCountsResponsePtrOutputWithContext(context.Context) NodeStateCountsResponsePtrOutput
}

type nodeStateCountsResponsePtrType NodeStateCountsResponseArgs

func NodeStateCountsResponsePtr(v *NodeStateCountsResponseArgs) NodeStateCountsResponsePtrInput {
	return (*nodeStateCountsResponsePtrType)(v)
}

func (*nodeStateCountsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeStateCountsResponse)(nil)).Elem()
}

func (i *nodeStateCountsResponsePtrType) ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput {
	return i.ToNodeStateCountsResponsePtrOutputWithContext(context.Background())
}

func (i *nodeStateCountsResponsePtrType) ToNodeStateCountsResponsePtrOutputWithContext(ctx context.Context) NodeStateCountsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeStateCountsResponsePtrOutput)
}

type NodeStateCountsResponseOutput struct{ *pulumi.OutputState }

func (NodeStateCountsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeStateCountsResponse)(nil)).Elem()
}

func (o NodeStateCountsResponseOutput) ToNodeStateCountsResponseOutput() NodeStateCountsResponseOutput {
	return o
}

func (o NodeStateCountsResponseOutput) ToNodeStateCountsResponseOutputWithContext(ctx context.Context) NodeStateCountsResponseOutput {
	return o
}

func (o NodeStateCountsResponseOutput) ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput {
	return o.ToNodeStateCountsResponsePtrOutputWithContext(context.Background())
}

func (o NodeStateCountsResponseOutput) ToNodeStateCountsResponsePtrOutputWithContext(ctx context.Context) NodeStateCountsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NodeStateCountsResponse) *NodeStateCountsResponse {
		return &v
	}).(NodeStateCountsResponsePtrOutput)
}

func (o NodeStateCountsResponseOutput) IdleNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.IdleNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) LeavingNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.LeavingNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) PreemptedNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.PreemptedNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) PreparingNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.PreparingNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) RunningNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.RunningNodeCount }).(pulumi.IntOutput)
}

func (o NodeStateCountsResponseOutput) UnusableNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodeStateCountsResponse) int { return v.UnusableNodeCount }).(pulumi.IntOutput)
}

type NodeStateCountsResponsePtrOutput struct{ *pulumi.OutputState }

func (NodeStateCountsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeStateCountsResponse)(nil)).Elem()
}

func (o NodeStateCountsResponsePtrOutput) ToNodeStateCountsResponsePtrOutput() NodeStateCountsResponsePtrOutput {
	return o
}

func (o NodeStateCountsResponsePtrOutput) ToNodeStateCountsResponsePtrOutputWithContext(ctx context.Context) NodeStateCountsResponsePtrOutput {
	return o
}

func (o NodeStateCountsResponsePtrOutput) Elem() NodeStateCountsResponseOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) NodeStateCountsResponse {
		if v != nil {
			return *v
		}
		var ret NodeStateCountsResponse
		return ret
	}).(NodeStateCountsResponseOutput)
}

func (o NodeStateCountsResponsePtrOutput) IdleNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.IdleNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) LeavingNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.LeavingNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) PreemptedNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.PreemptedNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) PreparingNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.PreparingNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) RunningNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.RunningNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o NodeStateCountsResponsePtrOutput) UnusableNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeStateCountsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.UnusableNodeCount
	}).(pulumi.IntPtrOutput)
}

type PasswordResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type PasswordResponseInput interface {
	pulumi.Input

	ToPasswordResponseOutput() PasswordResponseOutput
	ToPasswordResponseOutputWithContext(context.Context) PasswordResponseOutput
}

type PasswordResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (PasswordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PasswordResponse)(nil)).Elem()
}

func (i PasswordResponseArgs) ToPasswordResponseOutput() PasswordResponseOutput {
	return i.ToPasswordResponseOutputWithContext(context.Background())
}

func (i PasswordResponseArgs) ToPasswordResponseOutputWithContext(ctx context.Context) PasswordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordResponseOutput)
}





type PasswordResponseArrayInput interface {
	pulumi.Input

	ToPasswordResponseArrayOutput() PasswordResponseArrayOutput
	ToPasswordResponseArrayOutputWithContext(context.Context) PasswordResponseArrayOutput
}

type PasswordResponseArray []PasswordResponseInput

func (PasswordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PasswordResponse)(nil)).Elem()
}

func (i PasswordResponseArray) ToPasswordResponseArrayOutput() PasswordResponseArrayOutput {
	return i.ToPasswordResponseArrayOutputWithContext(context.Background())
}

func (i PasswordResponseArray) ToPasswordResponseArrayOutputWithContext(ctx context.Context) PasswordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordResponseArrayOutput)
}

type PasswordResponseOutput struct{ *pulumi.OutputState }

func (PasswordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PasswordResponse)(nil)).Elem()
}

func (o PasswordResponseOutput) ToPasswordResponseOutput() PasswordResponseOutput {
	return o
}

func (o PasswordResponseOutput) ToPasswordResponseOutputWithContext(ctx context.Context) PasswordResponseOutput {
	return o
}

func (o PasswordResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PasswordResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PasswordResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v PasswordResponse) string { return v.Value }).(pulumi.StringOutput)
}

type PasswordResponseArrayOutput struct{ *pulumi.OutputState }

func (PasswordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PasswordResponse)(nil)).Elem()
}

func (o PasswordResponseArrayOutput) ToPasswordResponseArrayOutput() PasswordResponseArrayOutput {
	return o
}

func (o PasswordResponseArrayOutput) ToPasswordResponseArrayOutputWithContext(ctx context.Context) PasswordResponseArrayOutput {
	return o
}

func (o PasswordResponseArrayOutput) Index(i pulumi.IntInput) PasswordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PasswordResponse {
		return vs[0].([]PasswordResponse)[vs[1].(int)]
	}).(PasswordResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RegistryListCredentialsResultResponse struct {
	Location  string             `pulumi:"location"`
	Passwords []PasswordResponse `pulumi:"passwords"`
	Username  string             `pulumi:"username"`
}





type RegistryListCredentialsResultResponseInput interface {
	pulumi.Input

	ToRegistryListCredentialsResultResponseOutput() RegistryListCredentialsResultResponseOutput
	ToRegistryListCredentialsResultResponseOutputWithContext(context.Context) RegistryListCredentialsResultResponseOutput
}

type RegistryListCredentialsResultResponseArgs struct {
	Location  pulumi.StringInput         `pulumi:"location"`
	Passwords PasswordResponseArrayInput `pulumi:"passwords"`
	Username  pulumi.StringInput         `pulumi:"username"`
}

func (RegistryListCredentialsResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryListCredentialsResultResponse)(nil)).Elem()
}

func (i RegistryListCredentialsResultResponseArgs) ToRegistryListCredentialsResultResponseOutput() RegistryListCredentialsResultResponseOutput {
	return i.ToRegistryListCredentialsResultResponseOutputWithContext(context.Background())
}

func (i RegistryListCredentialsResultResponseArgs) ToRegistryListCredentialsResultResponseOutputWithContext(ctx context.Context) RegistryListCredentialsResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryListCredentialsResultResponseOutput)
}

type RegistryListCredentialsResultResponseOutput struct{ *pulumi.OutputState }

func (RegistryListCredentialsResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryListCredentialsResultResponse)(nil)).Elem()
}

func (o RegistryListCredentialsResultResponseOutput) ToRegistryListCredentialsResultResponseOutput() RegistryListCredentialsResultResponseOutput {
	return o
}

func (o RegistryListCredentialsResultResponseOutput) ToRegistryListCredentialsResultResponseOutputWithContext(ctx context.Context) RegistryListCredentialsResultResponseOutput {
	return o
}

func (o RegistryListCredentialsResultResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RegistryListCredentialsResultResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o RegistryListCredentialsResultResponseOutput) Passwords() PasswordResponseArrayOutput {
	return o.ApplyT(func(v RegistryListCredentialsResultResponse) []PasswordResponse { return v.Passwords }).(PasswordResponseArrayOutput)
}

func (o RegistryListCredentialsResultResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v RegistryListCredentialsResultResponse) string { return v.Username }).(pulumi.StringOutput)
}

type ResourceId struct {
	Id string `pulumi:"id"`
}





type ResourceIdInput interface {
	pulumi.Input

	ToResourceIdOutput() ResourceIdOutput
	ToResourceIdOutputWithContext(context.Context) ResourceIdOutput
}

type ResourceIdArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceId)(nil)).Elem()
}

func (i ResourceIdArgs) ToResourceIdOutput() ResourceIdOutput {
	return i.ToResourceIdOutputWithContext(context.Background())
}

func (i ResourceIdArgs) ToResourceIdOutputWithContext(ctx context.Context) ResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdOutput)
}

func (i ResourceIdArgs) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return i.ToResourceIdPtrOutputWithContext(context.Background())
}

func (i ResourceIdArgs) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdOutput).ToResourceIdPtrOutputWithContext(ctx)
}









type ResourceIdPtrInput interface {
	pulumi.Input

	ToResourceIdPtrOutput() ResourceIdPtrOutput
	ToResourceIdPtrOutputWithContext(context.Context) ResourceIdPtrOutput
}

type resourceIdPtrType ResourceIdArgs

func ResourceIdPtr(v *ResourceIdArgs) ResourceIdPtrInput {
	return (*resourceIdPtrType)(v)
}

func (*resourceIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceId)(nil)).Elem()
}

func (i *resourceIdPtrType) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return i.ToResourceIdPtrOutputWithContext(context.Background())
}

func (i *resourceIdPtrType) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdPtrOutput)
}

type ResourceIdOutput struct{ *pulumi.OutputState }

func (ResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceId)(nil)).Elem()
}

func (o ResourceIdOutput) ToResourceIdOutput() ResourceIdOutput {
	return o
}

func (o ResourceIdOutput) ToResourceIdOutputWithContext(ctx context.Context) ResourceIdOutput {
	return o
}

func (o ResourceIdOutput) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return o.ToResourceIdPtrOutputWithContext(context.Background())
}

func (o ResourceIdOutput) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceId) *ResourceId {
		return &v
	}).(ResourceIdPtrOutput)
}

func (o ResourceIdOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceId) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceIdPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceId)(nil)).Elem()
}

func (o ResourceIdPtrOutput) ToResourceIdPtrOutput() ResourceIdPtrOutput {
	return o
}

func (o ResourceIdPtrOutput) ToResourceIdPtrOutputWithContext(ctx context.Context) ResourceIdPtrOutput {
	return o
}

func (o ResourceIdPtrOutput) Elem() ResourceIdOutput {
	return o.ApplyT(func(v *ResourceId) ResourceId {
		if v != nil {
			return *v
		}
		var ret ResourceId
		return ret
	}).(ResourceIdOutput)
}

func (o ResourceIdPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceId) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceIdResponse struct {
	Id string `pulumi:"id"`
}





type ResourceIdResponseInput interface {
	pulumi.Input

	ToResourceIdResponseOutput() ResourceIdResponseOutput
	ToResourceIdResponseOutputWithContext(context.Context) ResourceIdResponseOutput
}

type ResourceIdResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ResourceIdResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdResponse)(nil)).Elem()
}

func (i ResourceIdResponseArgs) ToResourceIdResponseOutput() ResourceIdResponseOutput {
	return i.ToResourceIdResponseOutputWithContext(context.Background())
}

func (i ResourceIdResponseArgs) ToResourceIdResponseOutputWithContext(ctx context.Context) ResourceIdResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponseOutput)
}

func (i ResourceIdResponseArgs) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return i.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdResponseArgs) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponseOutput).ToResourceIdResponsePtrOutputWithContext(ctx)
}









type ResourceIdResponsePtrInput interface {
	pulumi.Input

	ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput
	ToResourceIdResponsePtrOutputWithContext(context.Context) ResourceIdResponsePtrOutput
}

type resourceIdResponsePtrType ResourceIdResponseArgs

func ResourceIdResponsePtr(v *ResourceIdResponseArgs) ResourceIdResponsePtrInput {
	return (*resourceIdResponsePtrType)(v)
}

func (*resourceIdResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdResponse)(nil)).Elem()
}

func (i *resourceIdResponsePtrType) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return i.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdResponsePtrType) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponsePtrOutput)
}

type ResourceIdResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponseOutput) ToResourceIdResponseOutput() ResourceIdResponseOutput {
	return o
}

func (o ResourceIdResponseOutput) ToResourceIdResponseOutputWithContext(ctx context.Context) ResourceIdResponseOutput {
	return o
}

func (o ResourceIdResponseOutput) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return o.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdResponseOutput) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdResponse) *ResourceIdResponse {
		return &v
	}).(ResourceIdResponsePtrOutput)
}

func (o ResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponsePtrOutput) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return o
}

func (o ResourceIdResponsePtrOutput) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return o
}

func (o ResourceIdResponsePtrOutput) Elem() ResourceIdResponseOutput {
	return o.ApplyT(func(v *ResourceIdResponse) ResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdResponse
		return ret
	}).(ResourceIdResponseOutput)
}

func (o ResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ScaleSettings struct {
	MaxNodeCount                int     `pulumi:"maxNodeCount"`
	MinNodeCount                *int    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown *string `pulumi:"nodeIdleTimeBeforeScaleDown"`
}





type ScaleSettingsInput interface {
	pulumi.Input

	ToScaleSettingsOutput() ScaleSettingsOutput
	ToScaleSettingsOutputWithContext(context.Context) ScaleSettingsOutput
}

type ScaleSettingsArgs struct {
	MaxNodeCount                pulumi.IntInput       `pulumi:"maxNodeCount"`
	MinNodeCount                pulumi.IntPtrInput    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown pulumi.StringPtrInput `pulumi:"nodeIdleTimeBeforeScaleDown"`
}

func (ScaleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettings)(nil)).Elem()
}

func (i ScaleSettingsArgs) ToScaleSettingsOutput() ScaleSettingsOutput {
	return i.ToScaleSettingsOutputWithContext(context.Background())
}

func (i ScaleSettingsArgs) ToScaleSettingsOutputWithContext(ctx context.Context) ScaleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsOutput)
}

func (i ScaleSettingsArgs) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return i.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (i ScaleSettingsArgs) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsOutput).ToScaleSettingsPtrOutputWithContext(ctx)
}









type ScaleSettingsPtrInput interface {
	pulumi.Input

	ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput
	ToScaleSettingsPtrOutputWithContext(context.Context) ScaleSettingsPtrOutput
}

type scaleSettingsPtrType ScaleSettingsArgs

func ScaleSettingsPtr(v *ScaleSettingsArgs) ScaleSettingsPtrInput {
	return (*scaleSettingsPtrType)(v)
}

func (*scaleSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettings)(nil)).Elem()
}

func (i *scaleSettingsPtrType) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return i.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (i *scaleSettingsPtrType) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsPtrOutput)
}

type ScaleSettingsOutput struct{ *pulumi.OutputState }

func (ScaleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettings)(nil)).Elem()
}

func (o ScaleSettingsOutput) ToScaleSettingsOutput() ScaleSettingsOutput {
	return o
}

func (o ScaleSettingsOutput) ToScaleSettingsOutputWithContext(ctx context.Context) ScaleSettingsOutput {
	return o
}

func (o ScaleSettingsOutput) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return o.ToScaleSettingsPtrOutputWithContext(context.Background())
}

func (o ScaleSettingsOutput) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleSettings) *ScaleSettings {
		return &v
	}).(ScaleSettingsPtrOutput)
}

func (o ScaleSettingsOutput) MaxNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v ScaleSettings) int { return v.MaxNodeCount }).(pulumi.IntOutput)
}

func (o ScaleSettingsOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScaleSettings) *int { return v.MinNodeCount }).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsOutput) NodeIdleTimeBeforeScaleDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleSettings) *string { return v.NodeIdleTimeBeforeScaleDown }).(pulumi.StringPtrOutput)
}

type ScaleSettingsPtrOutput struct{ *pulumi.OutputState }

func (ScaleSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettings)(nil)).Elem()
}

func (o ScaleSettingsPtrOutput) ToScaleSettingsPtrOutput() ScaleSettingsPtrOutput {
	return o
}

func (o ScaleSettingsPtrOutput) ToScaleSettingsPtrOutputWithContext(ctx context.Context) ScaleSettingsPtrOutput {
	return o
}

func (o ScaleSettingsPtrOutput) Elem() ScaleSettingsOutput {
	return o.ApplyT(func(v *ScaleSettings) ScaleSettings {
		if v != nil {
			return *v
		}
		var ret ScaleSettings
		return ret
	}).(ScaleSettingsOutput)
}

func (o ScaleSettingsPtrOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleSettings) *int {
		if v == nil {
			return nil
		}
		return &v.MaxNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsPtrOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleSettings) *int {
		if v == nil {
			return nil
		}
		return v.MinNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsPtrOutput) NodeIdleTimeBeforeScaleDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScaleSettings) *string {
		if v == nil {
			return nil
		}
		return v.NodeIdleTimeBeforeScaleDown
	}).(pulumi.StringPtrOutput)
}

type ScaleSettingsResponse struct {
	MaxNodeCount                int     `pulumi:"maxNodeCount"`
	MinNodeCount                *int    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown *string `pulumi:"nodeIdleTimeBeforeScaleDown"`
}





type ScaleSettingsResponseInput interface {
	pulumi.Input

	ToScaleSettingsResponseOutput() ScaleSettingsResponseOutput
	ToScaleSettingsResponseOutputWithContext(context.Context) ScaleSettingsResponseOutput
}

type ScaleSettingsResponseArgs struct {
	MaxNodeCount                pulumi.IntInput       `pulumi:"maxNodeCount"`
	MinNodeCount                pulumi.IntPtrInput    `pulumi:"minNodeCount"`
	NodeIdleTimeBeforeScaleDown pulumi.StringPtrInput `pulumi:"nodeIdleTimeBeforeScaleDown"`
}

func (ScaleSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettingsResponse)(nil)).Elem()
}

func (i ScaleSettingsResponseArgs) ToScaleSettingsResponseOutput() ScaleSettingsResponseOutput {
	return i.ToScaleSettingsResponseOutputWithContext(context.Background())
}

func (i ScaleSettingsResponseArgs) ToScaleSettingsResponseOutputWithContext(ctx context.Context) ScaleSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsResponseOutput)
}

func (i ScaleSettingsResponseArgs) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return i.ToScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ScaleSettingsResponseArgs) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsResponseOutput).ToScaleSettingsResponsePtrOutputWithContext(ctx)
}









type ScaleSettingsResponsePtrInput interface {
	pulumi.Input

	ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput
	ToScaleSettingsResponsePtrOutputWithContext(context.Context) ScaleSettingsResponsePtrOutput
}

type scaleSettingsResponsePtrType ScaleSettingsResponseArgs

func ScaleSettingsResponsePtr(v *ScaleSettingsResponseArgs) ScaleSettingsResponsePtrInput {
	return (*scaleSettingsResponsePtrType)(v)
}

func (*scaleSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettingsResponse)(nil)).Elem()
}

func (i *scaleSettingsResponsePtrType) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return i.ToScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *scaleSettingsResponsePtrType) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleSettingsResponsePtrOutput)
}

type ScaleSettingsResponseOutput struct{ *pulumi.OutputState }

func (ScaleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleSettingsResponse)(nil)).Elem()
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponseOutput() ScaleSettingsResponseOutput {
	return o
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponseOutputWithContext(ctx context.Context) ScaleSettingsResponseOutput {
	return o
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return o.ToScaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ScaleSettingsResponseOutput) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleSettingsResponse) *ScaleSettingsResponse {
		return &v
	}).(ScaleSettingsResponsePtrOutput)
}

func (o ScaleSettingsResponseOutput) MaxNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v ScaleSettingsResponse) int { return v.MaxNodeCount }).(pulumi.IntOutput)
}

func (o ScaleSettingsResponseOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScaleSettingsResponse) *int { return v.MinNodeCount }).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsResponseOutput) NodeIdleTimeBeforeScaleDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleSettingsResponse) *string { return v.NodeIdleTimeBeforeScaleDown }).(pulumi.StringPtrOutput)
}

type ScaleSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ScaleSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleSettingsResponse)(nil)).Elem()
}

func (o ScaleSettingsResponsePtrOutput) ToScaleSettingsResponsePtrOutput() ScaleSettingsResponsePtrOutput {
	return o
}

func (o ScaleSettingsResponsePtrOutput) ToScaleSettingsResponsePtrOutputWithContext(ctx context.Context) ScaleSettingsResponsePtrOutput {
	return o
}

func (o ScaleSettingsResponsePtrOutput) Elem() ScaleSettingsResponseOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) ScaleSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ScaleSettingsResponse
		return ret
	}).(ScaleSettingsResponseOutput)
}

func (o ScaleSettingsResponsePtrOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsResponsePtrOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o ScaleSettingsResponsePtrOutput) NodeIdleTimeBeforeScaleDown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScaleSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NodeIdleTimeBeforeScaleDown
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SslConfiguration struct {
	Cert   *string `pulumi:"cert"`
	Cname  *string `pulumi:"cname"`
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
	Cname  pulumi.StringPtrInput `pulumi:"cname"`
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

func (o SslConfigurationOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.Cname }).(pulumi.StringPtrOutput)
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

func (o SslConfigurationPtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Cname
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
	Cname  *string `pulumi:"cname"`
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
	Cname  pulumi.StringPtrInput `pulumi:"cname"`
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

func (o SslConfigurationResponseOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.Cname }).(pulumi.StringPtrOutput)
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

func (o SslConfigurationResponsePtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cname
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

type SystemServiceResponse struct {
	PublicIpAddress   string `pulumi:"publicIpAddress"`
	SystemServiceType string `pulumi:"systemServiceType"`
	Version           string `pulumi:"version"`
}





type SystemServiceResponseInput interface {
	pulumi.Input

	ToSystemServiceResponseOutput() SystemServiceResponseOutput
	ToSystemServiceResponseOutputWithContext(context.Context) SystemServiceResponseOutput
}

type SystemServiceResponseArgs struct {
	PublicIpAddress   pulumi.StringInput `pulumi:"publicIpAddress"`
	SystemServiceType pulumi.StringInput `pulumi:"systemServiceType"`
	Version           pulumi.StringInput `pulumi:"version"`
}

func (SystemServiceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemServiceResponse)(nil)).Elem()
}

func (i SystemServiceResponseArgs) ToSystemServiceResponseOutput() SystemServiceResponseOutput {
	return i.ToSystemServiceResponseOutputWithContext(context.Background())
}

func (i SystemServiceResponseArgs) ToSystemServiceResponseOutputWithContext(ctx context.Context) SystemServiceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemServiceResponseOutput)
}





type SystemServiceResponseArrayInput interface {
	pulumi.Input

	ToSystemServiceResponseArrayOutput() SystemServiceResponseArrayOutput
	ToSystemServiceResponseArrayOutputWithContext(context.Context) SystemServiceResponseArrayOutput
}

type SystemServiceResponseArray []SystemServiceResponseInput

func (SystemServiceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemServiceResponse)(nil)).Elem()
}

func (i SystemServiceResponseArray) ToSystemServiceResponseArrayOutput() SystemServiceResponseArrayOutput {
	return i.ToSystemServiceResponseArrayOutputWithContext(context.Background())
}

func (i SystemServiceResponseArray) ToSystemServiceResponseArrayOutputWithContext(ctx context.Context) SystemServiceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemServiceResponseArrayOutput)
}

type SystemServiceResponseOutput struct{ *pulumi.OutputState }

func (SystemServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemServiceResponse)(nil)).Elem()
}

func (o SystemServiceResponseOutput) ToSystemServiceResponseOutput() SystemServiceResponseOutput {
	return o
}

func (o SystemServiceResponseOutput) ToSystemServiceResponseOutputWithContext(ctx context.Context) SystemServiceResponseOutput {
	return o
}

func (o SystemServiceResponseOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v SystemServiceResponse) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

func (o SystemServiceResponseOutput) SystemServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v SystemServiceResponse) string { return v.SystemServiceType }).(pulumi.StringOutput)
}

func (o SystemServiceResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v SystemServiceResponse) string { return v.Version }).(pulumi.StringOutput)
}

type SystemServiceResponseArrayOutput struct{ *pulumi.OutputState }

func (SystemServiceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemServiceResponse)(nil)).Elem()
}

func (o SystemServiceResponseArrayOutput) ToSystemServiceResponseArrayOutput() SystemServiceResponseArrayOutput {
	return o
}

func (o SystemServiceResponseArrayOutput) ToSystemServiceResponseArrayOutputWithContext(ctx context.Context) SystemServiceResponseArrayOutput {
	return o
}

func (o SystemServiceResponseArrayOutput) Index(i pulumi.IntInput) SystemServiceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemServiceResponse {
		return vs[0].([]SystemServiceResponse)[vs[1].(int)]
	}).(SystemServiceResponseOutput)
}

type UserAccountCredentials struct {
	AdminUserName         string  `pulumi:"adminUserName"`
	AdminUserPassword     *string `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey *string `pulumi:"adminUserSshPublicKey"`
}





type UserAccountCredentialsInput interface {
	pulumi.Input

	ToUserAccountCredentialsOutput() UserAccountCredentialsOutput
	ToUserAccountCredentialsOutputWithContext(context.Context) UserAccountCredentialsOutput
}

type UserAccountCredentialsArgs struct {
	AdminUserName         pulumi.StringInput    `pulumi:"adminUserName"`
	AdminUserPassword     pulumi.StringPtrInput `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey pulumi.StringPtrInput `pulumi:"adminUserSshPublicKey"`
}

func (UserAccountCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountCredentials)(nil)).Elem()
}

func (i UserAccountCredentialsArgs) ToUserAccountCredentialsOutput() UserAccountCredentialsOutput {
	return i.ToUserAccountCredentialsOutputWithContext(context.Background())
}

func (i UserAccountCredentialsArgs) ToUserAccountCredentialsOutputWithContext(ctx context.Context) UserAccountCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsOutput)
}

func (i UserAccountCredentialsArgs) ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput {
	return i.ToUserAccountCredentialsPtrOutputWithContext(context.Background())
}

func (i UserAccountCredentialsArgs) ToUserAccountCredentialsPtrOutputWithContext(ctx context.Context) UserAccountCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsOutput).ToUserAccountCredentialsPtrOutputWithContext(ctx)
}









type UserAccountCredentialsPtrInput interface {
	pulumi.Input

	ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput
	ToUserAccountCredentialsPtrOutputWithContext(context.Context) UserAccountCredentialsPtrOutput
}

type userAccountCredentialsPtrType UserAccountCredentialsArgs

func UserAccountCredentialsPtr(v *UserAccountCredentialsArgs) UserAccountCredentialsPtrInput {
	return (*userAccountCredentialsPtrType)(v)
}

func (*userAccountCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccountCredentials)(nil)).Elem()
}

func (i *userAccountCredentialsPtrType) ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput {
	return i.ToUserAccountCredentialsPtrOutputWithContext(context.Background())
}

func (i *userAccountCredentialsPtrType) ToUserAccountCredentialsPtrOutputWithContext(ctx context.Context) UserAccountCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsPtrOutput)
}

type UserAccountCredentialsOutput struct{ *pulumi.OutputState }

func (UserAccountCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountCredentials)(nil)).Elem()
}

func (o UserAccountCredentialsOutput) ToUserAccountCredentialsOutput() UserAccountCredentialsOutput {
	return o
}

func (o UserAccountCredentialsOutput) ToUserAccountCredentialsOutputWithContext(ctx context.Context) UserAccountCredentialsOutput {
	return o
}

func (o UserAccountCredentialsOutput) ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput {
	return o.ToUserAccountCredentialsPtrOutputWithContext(context.Background())
}

func (o UserAccountCredentialsOutput) ToUserAccountCredentialsPtrOutputWithContext(ctx context.Context) UserAccountCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAccountCredentials) *UserAccountCredentials {
		return &v
	}).(UserAccountCredentialsPtrOutput)
}

func (o UserAccountCredentialsOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccountCredentials) string { return v.AdminUserName }).(pulumi.StringOutput)
}

func (o UserAccountCredentialsOutput) AdminUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountCredentials) *string { return v.AdminUserPassword }).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsOutput) AdminUserSshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountCredentials) *string { return v.AdminUserSshPublicKey }).(pulumi.StringPtrOutput)
}

type UserAccountCredentialsPtrOutput struct{ *pulumi.OutputState }

func (UserAccountCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccountCredentials)(nil)).Elem()
}

func (o UserAccountCredentialsPtrOutput) ToUserAccountCredentialsPtrOutput() UserAccountCredentialsPtrOutput {
	return o
}

func (o UserAccountCredentialsPtrOutput) ToUserAccountCredentialsPtrOutputWithContext(ctx context.Context) UserAccountCredentialsPtrOutput {
	return o
}

func (o UserAccountCredentialsPtrOutput) Elem() UserAccountCredentialsOutput {
	return o.ApplyT(func(v *UserAccountCredentials) UserAccountCredentials {
		if v != nil {
			return *v
		}
		var ret UserAccountCredentials
		return ret
	}).(UserAccountCredentialsOutput)
}

func (o UserAccountCredentialsPtrOutput) AdminUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUserName
	}).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsPtrOutput) AdminUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AdminUserPassword
	}).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsPtrOutput) AdminUserSshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentials) *string {
		if v == nil {
			return nil
		}
		return v.AdminUserSshPublicKey
	}).(pulumi.StringPtrOutput)
}

type UserAccountCredentialsResponse struct {
	AdminUserName         string  `pulumi:"adminUserName"`
	AdminUserPassword     *string `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey *string `pulumi:"adminUserSshPublicKey"`
}





type UserAccountCredentialsResponseInput interface {
	pulumi.Input

	ToUserAccountCredentialsResponseOutput() UserAccountCredentialsResponseOutput
	ToUserAccountCredentialsResponseOutputWithContext(context.Context) UserAccountCredentialsResponseOutput
}

type UserAccountCredentialsResponseArgs struct {
	AdminUserName         pulumi.StringInput    `pulumi:"adminUserName"`
	AdminUserPassword     pulumi.StringPtrInput `pulumi:"adminUserPassword"`
	AdminUserSshPublicKey pulumi.StringPtrInput `pulumi:"adminUserSshPublicKey"`
}

func (UserAccountCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountCredentialsResponse)(nil)).Elem()
}

func (i UserAccountCredentialsResponseArgs) ToUserAccountCredentialsResponseOutput() UserAccountCredentialsResponseOutput {
	return i.ToUserAccountCredentialsResponseOutputWithContext(context.Background())
}

func (i UserAccountCredentialsResponseArgs) ToUserAccountCredentialsResponseOutputWithContext(ctx context.Context) UserAccountCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsResponseOutput)
}

func (i UserAccountCredentialsResponseArgs) ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput {
	return i.ToUserAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i UserAccountCredentialsResponseArgs) ToUserAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) UserAccountCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsResponseOutput).ToUserAccountCredentialsResponsePtrOutputWithContext(ctx)
}









type UserAccountCredentialsResponsePtrInput interface {
	pulumi.Input

	ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput
	ToUserAccountCredentialsResponsePtrOutputWithContext(context.Context) UserAccountCredentialsResponsePtrOutput
}

type userAccountCredentialsResponsePtrType UserAccountCredentialsResponseArgs

func UserAccountCredentialsResponsePtr(v *UserAccountCredentialsResponseArgs) UserAccountCredentialsResponsePtrInput {
	return (*userAccountCredentialsResponsePtrType)(v)
}

func (*userAccountCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccountCredentialsResponse)(nil)).Elem()
}

func (i *userAccountCredentialsResponsePtrType) ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput {
	return i.ToUserAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *userAccountCredentialsResponsePtrType) ToUserAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) UserAccountCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAccountCredentialsResponsePtrOutput)
}

type UserAccountCredentialsResponseOutput struct{ *pulumi.OutputState }

func (UserAccountCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAccountCredentialsResponse)(nil)).Elem()
}

func (o UserAccountCredentialsResponseOutput) ToUserAccountCredentialsResponseOutput() UserAccountCredentialsResponseOutput {
	return o
}

func (o UserAccountCredentialsResponseOutput) ToUserAccountCredentialsResponseOutputWithContext(ctx context.Context) UserAccountCredentialsResponseOutput {
	return o
}

func (o UserAccountCredentialsResponseOutput) ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput {
	return o.ToUserAccountCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o UserAccountCredentialsResponseOutput) ToUserAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) UserAccountCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAccountCredentialsResponse) *UserAccountCredentialsResponse {
		return &v
	}).(UserAccountCredentialsResponsePtrOutput)
}

func (o UserAccountCredentialsResponseOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v UserAccountCredentialsResponse) string { return v.AdminUserName }).(pulumi.StringOutput)
}

func (o UserAccountCredentialsResponseOutput) AdminUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountCredentialsResponse) *string { return v.AdminUserPassword }).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsResponseOutput) AdminUserSshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAccountCredentialsResponse) *string { return v.AdminUserSshPublicKey }).(pulumi.StringPtrOutput)
}

type UserAccountCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (UserAccountCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAccountCredentialsResponse)(nil)).Elem()
}

func (o UserAccountCredentialsResponsePtrOutput) ToUserAccountCredentialsResponsePtrOutput() UserAccountCredentialsResponsePtrOutput {
	return o
}

func (o UserAccountCredentialsResponsePtrOutput) ToUserAccountCredentialsResponsePtrOutputWithContext(ctx context.Context) UserAccountCredentialsResponsePtrOutput {
	return o
}

func (o UserAccountCredentialsResponsePtrOutput) Elem() UserAccountCredentialsResponseOutput {
	return o.ApplyT(func(v *UserAccountCredentialsResponse) UserAccountCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret UserAccountCredentialsResponse
		return ret
	}).(UserAccountCredentialsResponseOutput)
}

func (o UserAccountCredentialsResponsePtrOutput) AdminUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUserName
	}).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsResponsePtrOutput) AdminUserPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUserPassword
	}).(pulumi.StringPtrOutput)
}

func (o UserAccountCredentialsResponsePtrOutput) AdminUserSshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAccountCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUserSshPublicKey
	}).(pulumi.StringPtrOutput)
}

type VirtualMachine struct {
	ComputeLocation *string                   `pulumi:"computeLocation"`
	ComputeType     string                    `pulumi:"computeType"`
	Description     *string                   `pulumi:"description"`
	Properties      *VirtualMachineProperties `pulumi:"properties"`
	ResourceId      *string                   `pulumi:"resourceId"`
}





type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(context.Context) VirtualMachineOutput
}

type VirtualMachineArgs struct {
	ComputeLocation pulumi.StringPtrInput            `pulumi:"computeLocation"`
	ComputeType     pulumi.StringInput               `pulumi:"computeType"`
	Description     pulumi.StringPtrInput            `pulumi:"description"`
	Properties      VirtualMachinePropertiesPtrInput `pulumi:"properties"`
	ResourceId      pulumi.StringPtrInput            `pulumi:"resourceId"`
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachine)(nil)).Elem()
}

func (i VirtualMachineArgs) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i VirtualMachineArgs) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachine) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Properties() VirtualMachinePropertiesPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *VirtualMachineProperties { return v.Properties }).(VirtualMachinePropertiesPtrOutput)
}

func (o VirtualMachineOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VirtualMachineProperties struct {
	Address              *string                       `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	SshPort              *int                          `pulumi:"sshPort"`
	VirtualMachineSize   *string                       `pulumi:"virtualMachineSize"`
}





type VirtualMachinePropertiesInput interface {
	pulumi.Input

	ToVirtualMachinePropertiesOutput() VirtualMachinePropertiesOutput
	ToVirtualMachinePropertiesOutputWithContext(context.Context) VirtualMachinePropertiesOutput
}

type VirtualMachinePropertiesArgs struct {
	Address              pulumi.StringPtrInput                `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsPtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                   `pulumi:"sshPort"`
	VirtualMachineSize   pulumi.StringPtrInput                `pulumi:"virtualMachineSize"`
}

func (VirtualMachinePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineProperties)(nil)).Elem()
}

func (i VirtualMachinePropertiesArgs) ToVirtualMachinePropertiesOutput() VirtualMachinePropertiesOutput {
	return i.ToVirtualMachinePropertiesOutputWithContext(context.Background())
}

func (i VirtualMachinePropertiesArgs) ToVirtualMachinePropertiesOutputWithContext(ctx context.Context) VirtualMachinePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePropertiesOutput)
}

func (i VirtualMachinePropertiesArgs) ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput {
	return i.ToVirtualMachinePropertiesPtrOutputWithContext(context.Background())
}

func (i VirtualMachinePropertiesArgs) ToVirtualMachinePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachinePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePropertiesOutput).ToVirtualMachinePropertiesPtrOutputWithContext(ctx)
}









type VirtualMachinePropertiesPtrInput interface {
	pulumi.Input

	ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput
	ToVirtualMachinePropertiesPtrOutputWithContext(context.Context) VirtualMachinePropertiesPtrOutput
}

type virtualMachinePropertiesPtrType VirtualMachinePropertiesArgs

func VirtualMachinePropertiesPtr(v *VirtualMachinePropertiesArgs) VirtualMachinePropertiesPtrInput {
	return (*virtualMachinePropertiesPtrType)(v)
}

func (*virtualMachinePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineProperties)(nil)).Elem()
}

func (i *virtualMachinePropertiesPtrType) ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput {
	return i.ToVirtualMachinePropertiesPtrOutputWithContext(context.Background())
}

func (i *virtualMachinePropertiesPtrType) ToVirtualMachinePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachinePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePropertiesPtrOutput)
}

type VirtualMachinePropertiesOutput struct{ *pulumi.OutputState }

func (VirtualMachinePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineProperties)(nil)).Elem()
}

func (o VirtualMachinePropertiesOutput) ToVirtualMachinePropertiesOutput() VirtualMachinePropertiesOutput {
	return o
}

func (o VirtualMachinePropertiesOutput) ToVirtualMachinePropertiesOutputWithContext(ctx context.Context) VirtualMachinePropertiesOutput {
	return o
}

func (o VirtualMachinePropertiesOutput) ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput {
	return o.ToVirtualMachinePropertiesPtrOutputWithContext(context.Background())
}

func (o VirtualMachinePropertiesOutput) ToVirtualMachinePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachinePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineProperties) *VirtualMachineProperties {
		return &v
	}).(VirtualMachinePropertiesPtrOutput)
}

func (o VirtualMachinePropertiesOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePropertiesOutput) AdministratorAccount() VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *VirtualMachineSshCredentials { return v.AdministratorAccount }).(VirtualMachineSshCredentialsPtrOutput)
}

func (o VirtualMachinePropertiesOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

func (o VirtualMachinePropertiesOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *string { return v.VirtualMachineSize }).(pulumi.StringPtrOutput)
}

type VirtualMachinePropertiesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachinePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineProperties)(nil)).Elem()
}

func (o VirtualMachinePropertiesPtrOutput) ToVirtualMachinePropertiesPtrOutput() VirtualMachinePropertiesPtrOutput {
	return o
}

func (o VirtualMachinePropertiesPtrOutput) ToVirtualMachinePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachinePropertiesPtrOutput {
	return o
}

func (o VirtualMachinePropertiesPtrOutput) Elem() VirtualMachinePropertiesOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) VirtualMachineProperties {
		if v != nil {
			return *v
		}
		var ret VirtualMachineProperties
		return ret
	}).(VirtualMachinePropertiesOutput)
}

func (o VirtualMachinePropertiesPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePropertiesPtrOutput) AdministratorAccount() VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *VirtualMachineSshCredentials {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsPtrOutput)
}

func (o VirtualMachinePropertiesPtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachinePropertiesPtrOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *string {
		if v == nil {
			return nil
		}
		return v.VirtualMachineSize
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineResponse struct {
	ComputeLocation    *string                               `pulumi:"computeLocation"`
	ComputeType        string                                `pulumi:"computeType"`
	CreatedOn          string                                `pulumi:"createdOn"`
	Description        *string                               `pulumi:"description"`
	IsAttachedCompute  bool                                  `pulumi:"isAttachedCompute"`
	ModifiedOn         string                                `pulumi:"modifiedOn"`
	Properties         *VirtualMachineResponseProperties     `pulumi:"properties"`
	ProvisioningErrors []MachineLearningServiceErrorResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                                `pulumi:"provisioningState"`
	ResourceId         *string                               `pulumi:"resourceId"`
}





type VirtualMachineResponseInput interface {
	pulumi.Input

	ToVirtualMachineResponseOutput() VirtualMachineResponseOutput
	ToVirtualMachineResponseOutputWithContext(context.Context) VirtualMachineResponseOutput
}

type VirtualMachineResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                         `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                            `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                            `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                         `pulumi:"description"`
	IsAttachedCompute  pulumi.BoolInput                              `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                            `pulumi:"modifiedOn"`
	Properties         VirtualMachineResponsePropertiesPtrInput      `pulumi:"properties"`
	ProvisioningErrors MachineLearningServiceErrorResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                            `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                         `pulumi:"resourceId"`
}

func (VirtualMachineResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResponse)(nil)).Elem()
}

func (i VirtualMachineResponseArgs) ToVirtualMachineResponseOutput() VirtualMachineResponseOutput {
	return i.ToVirtualMachineResponseOutputWithContext(context.Background())
}

func (i VirtualMachineResponseArgs) ToVirtualMachineResponseOutputWithContext(ctx context.Context) VirtualMachineResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResponseOutput)
}

type VirtualMachineResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResponse)(nil)).Elem()
}

func (o VirtualMachineResponseOutput) ToVirtualMachineResponseOutput() VirtualMachineResponseOutput {
	return o
}

func (o VirtualMachineResponseOutput) ToVirtualMachineResponseOutputWithContext(ctx context.Context) VirtualMachineResponseOutput {
	return o
}

func (o VirtualMachineResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v VirtualMachineResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o VirtualMachineResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) Properties() VirtualMachineResponsePropertiesPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *VirtualMachineResponseProperties { return v.Properties }).(VirtualMachineResponsePropertiesPtrOutput)
}

func (o VirtualMachineResponseOutput) ProvisioningErrors() MachineLearningServiceErrorResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineResponse) []MachineLearningServiceErrorResponse { return v.ProvisioningErrors }).(MachineLearningServiceErrorResponseArrayOutput)
}

func (o VirtualMachineResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VirtualMachineResponseProperties struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
	VirtualMachineSize   *string                               `pulumi:"virtualMachineSize"`
}





type VirtualMachineResponsePropertiesInput interface {
	pulumi.Input

	ToVirtualMachineResponsePropertiesOutput() VirtualMachineResponsePropertiesOutput
	ToVirtualMachineResponsePropertiesOutputWithContext(context.Context) VirtualMachineResponsePropertiesOutput
}

type VirtualMachineResponsePropertiesArgs struct {
	Address              pulumi.StringPtrInput                        `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsResponsePtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                           `pulumi:"sshPort"`
	VirtualMachineSize   pulumi.StringPtrInput                        `pulumi:"virtualMachineSize"`
}

func (VirtualMachineResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResponseProperties)(nil)).Elem()
}

func (i VirtualMachineResponsePropertiesArgs) ToVirtualMachineResponsePropertiesOutput() VirtualMachineResponsePropertiesOutput {
	return i.ToVirtualMachineResponsePropertiesOutputWithContext(context.Background())
}

func (i VirtualMachineResponsePropertiesArgs) ToVirtualMachineResponsePropertiesOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResponsePropertiesOutput)
}

func (i VirtualMachineResponsePropertiesArgs) ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput {
	return i.ToVirtualMachineResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i VirtualMachineResponsePropertiesArgs) ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResponsePropertiesOutput).ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx)
}









type VirtualMachineResponsePropertiesPtrInput interface {
	pulumi.Input

	ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput
	ToVirtualMachineResponsePropertiesPtrOutputWithContext(context.Context) VirtualMachineResponsePropertiesPtrOutput
}

type virtualMachineResponsePropertiesPtrType VirtualMachineResponsePropertiesArgs

func VirtualMachineResponsePropertiesPtr(v *VirtualMachineResponsePropertiesArgs) VirtualMachineResponsePropertiesPtrInput {
	return (*virtualMachineResponsePropertiesPtrType)(v)
}

func (*virtualMachineResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineResponseProperties)(nil)).Elem()
}

func (i *virtualMachineResponsePropertiesPtrType) ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput {
	return i.ToVirtualMachineResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *virtualMachineResponsePropertiesPtrType) ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResponsePropertiesPtrOutput)
}

type VirtualMachineResponsePropertiesOutput struct{ *pulumi.OutputState }

func (VirtualMachineResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResponseProperties)(nil)).Elem()
}

func (o VirtualMachineResponsePropertiesOutput) ToVirtualMachineResponsePropertiesOutput() VirtualMachineResponsePropertiesOutput {
	return o
}

func (o VirtualMachineResponsePropertiesOutput) ToVirtualMachineResponsePropertiesOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesOutput {
	return o
}

func (o VirtualMachineResponsePropertiesOutput) ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput {
	return o.ToVirtualMachineResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o VirtualMachineResponsePropertiesOutput) ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineResponseProperties) *VirtualMachineResponseProperties {
		return &v
	}).(VirtualMachineResponsePropertiesPtrOutput)
}

func (o VirtualMachineResponsePropertiesOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResponsePropertiesOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *VirtualMachineSshCredentialsResponse {
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o VirtualMachineResponsePropertiesOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineResponsePropertiesOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *string { return v.VirtualMachineSize }).(pulumi.StringPtrOutput)
}

type VirtualMachineResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineResponseProperties)(nil)).Elem()
}

func (o VirtualMachineResponsePropertiesPtrOutput) ToVirtualMachineResponsePropertiesPtrOutput() VirtualMachineResponsePropertiesPtrOutput {
	return o
}

func (o VirtualMachineResponsePropertiesPtrOutput) ToVirtualMachineResponsePropertiesPtrOutputWithContext(ctx context.Context) VirtualMachineResponsePropertiesPtrOutput {
	return o
}

func (o VirtualMachineResponsePropertiesPtrOutput) Elem() VirtualMachineResponsePropertiesOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) VirtualMachineResponseProperties {
		if v != nil {
			return *v
		}
		var ret VirtualMachineResponseProperties
		return ret
	}).(VirtualMachineResponsePropertiesOutput)
}

func (o VirtualMachineResponsePropertiesPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResponsePropertiesPtrOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *VirtualMachineSshCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o VirtualMachineResponsePropertiesPtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineResponsePropertiesPtrOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.VirtualMachineSize
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineSshCredentials struct {
	Password       *string `pulumi:"password"`
	PrivateKeyData *string `pulumi:"privateKeyData"`
	PublicKeyData  *string `pulumi:"publicKeyData"`
	Username       *string `pulumi:"username"`
}





type VirtualMachineSshCredentialsInput interface {
	pulumi.Input

	ToVirtualMachineSshCredentialsOutput() VirtualMachineSshCredentialsOutput
	ToVirtualMachineSshCredentialsOutputWithContext(context.Context) VirtualMachineSshCredentialsOutput
}

type VirtualMachineSshCredentialsArgs struct {
	Password       pulumi.StringPtrInput `pulumi:"password"`
	PrivateKeyData pulumi.StringPtrInput `pulumi:"privateKeyData"`
	PublicKeyData  pulumi.StringPtrInput `pulumi:"publicKeyData"`
	Username       pulumi.StringPtrInput `pulumi:"username"`
}

func (VirtualMachineSshCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSshCredentials)(nil)).Elem()
}

func (i VirtualMachineSshCredentialsArgs) ToVirtualMachineSshCredentialsOutput() VirtualMachineSshCredentialsOutput {
	return i.ToVirtualMachineSshCredentialsOutputWithContext(context.Background())
}

func (i VirtualMachineSshCredentialsArgs) ToVirtualMachineSshCredentialsOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsOutput)
}

func (i VirtualMachineSshCredentialsArgs) ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput {
	return i.ToVirtualMachineSshCredentialsPtrOutputWithContext(context.Background())
}

func (i VirtualMachineSshCredentialsArgs) ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsOutput).ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx)
}









type VirtualMachineSshCredentialsPtrInput interface {
	pulumi.Input

	ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput
	ToVirtualMachineSshCredentialsPtrOutputWithContext(context.Context) VirtualMachineSshCredentialsPtrOutput
}

type virtualMachineSshCredentialsPtrType VirtualMachineSshCredentialsArgs

func VirtualMachineSshCredentialsPtr(v *VirtualMachineSshCredentialsArgs) VirtualMachineSshCredentialsPtrInput {
	return (*virtualMachineSshCredentialsPtrType)(v)
}

func (*virtualMachineSshCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSshCredentials)(nil)).Elem()
}

func (i *virtualMachineSshCredentialsPtrType) ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput {
	return i.ToVirtualMachineSshCredentialsPtrOutputWithContext(context.Background())
}

func (i *virtualMachineSshCredentialsPtrType) ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsPtrOutput)
}

type VirtualMachineSshCredentialsOutput struct{ *pulumi.OutputState }

func (VirtualMachineSshCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSshCredentials)(nil)).Elem()
}

func (o VirtualMachineSshCredentialsOutput) ToVirtualMachineSshCredentialsOutput() VirtualMachineSshCredentialsOutput {
	return o
}

func (o VirtualMachineSshCredentialsOutput) ToVirtualMachineSshCredentialsOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsOutput {
	return o
}

func (o VirtualMachineSshCredentialsOutput) ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput {
	return o.ToVirtualMachineSshCredentialsPtrOutputWithContext(context.Background())
}

func (o VirtualMachineSshCredentialsOutput) ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineSshCredentials) *VirtualMachineSshCredentials {
		return &v
	}).(VirtualMachineSshCredentialsPtrOutput)
}

func (o VirtualMachineSshCredentialsOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentials) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsOutput) PrivateKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentials) *string { return v.PrivateKeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentials) *string { return v.PublicKeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentials) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type VirtualMachineSshCredentialsPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineSshCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSshCredentials)(nil)).Elem()
}

func (o VirtualMachineSshCredentialsPtrOutput) ToVirtualMachineSshCredentialsPtrOutput() VirtualMachineSshCredentialsPtrOutput {
	return o
}

func (o VirtualMachineSshCredentialsPtrOutput) ToVirtualMachineSshCredentialsPtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsPtrOutput {
	return o
}

func (o VirtualMachineSshCredentialsPtrOutput) Elem() VirtualMachineSshCredentialsOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) VirtualMachineSshCredentials {
		if v != nil {
			return *v
		}
		var ret VirtualMachineSshCredentials
		return ret
	}).(VirtualMachineSshCredentialsOutput)
}

func (o VirtualMachineSshCredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsPtrOutput) PrivateKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKeyData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsPtrOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) *string {
		if v == nil {
			return nil
		}
		return v.PublicKeyData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentials) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineSshCredentialsResponse struct {
	Password       *string `pulumi:"password"`
	PrivateKeyData *string `pulumi:"privateKeyData"`
	PublicKeyData  *string `pulumi:"publicKeyData"`
	Username       *string `pulumi:"username"`
}





type VirtualMachineSshCredentialsResponseInput interface {
	pulumi.Input

	ToVirtualMachineSshCredentialsResponseOutput() VirtualMachineSshCredentialsResponseOutput
	ToVirtualMachineSshCredentialsResponseOutputWithContext(context.Context) VirtualMachineSshCredentialsResponseOutput
}

type VirtualMachineSshCredentialsResponseArgs struct {
	Password       pulumi.StringPtrInput `pulumi:"password"`
	PrivateKeyData pulumi.StringPtrInput `pulumi:"privateKeyData"`
	PublicKeyData  pulumi.StringPtrInput `pulumi:"publicKeyData"`
	Username       pulumi.StringPtrInput `pulumi:"username"`
}

func (VirtualMachineSshCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSshCredentialsResponse)(nil)).Elem()
}

func (i VirtualMachineSshCredentialsResponseArgs) ToVirtualMachineSshCredentialsResponseOutput() VirtualMachineSshCredentialsResponseOutput {
	return i.ToVirtualMachineSshCredentialsResponseOutputWithContext(context.Background())
}

func (i VirtualMachineSshCredentialsResponseArgs) ToVirtualMachineSshCredentialsResponseOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsResponseOutput)
}

func (i VirtualMachineSshCredentialsResponseArgs) ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput {
	return i.ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i VirtualMachineSshCredentialsResponseArgs) ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsResponseOutput).ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx)
}









type VirtualMachineSshCredentialsResponsePtrInput interface {
	pulumi.Input

	ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput
	ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(context.Context) VirtualMachineSshCredentialsResponsePtrOutput
}

type virtualMachineSshCredentialsResponsePtrType VirtualMachineSshCredentialsResponseArgs

func VirtualMachineSshCredentialsResponsePtr(v *VirtualMachineSshCredentialsResponseArgs) VirtualMachineSshCredentialsResponsePtrInput {
	return (*virtualMachineSshCredentialsResponsePtrType)(v)
}

func (*virtualMachineSshCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSshCredentialsResponse)(nil)).Elem()
}

func (i *virtualMachineSshCredentialsResponsePtrType) ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput {
	return i.ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *virtualMachineSshCredentialsResponsePtrType) ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSshCredentialsResponsePtrOutput)
}

type VirtualMachineSshCredentialsResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineSshCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSshCredentialsResponse)(nil)).Elem()
}

func (o VirtualMachineSshCredentialsResponseOutput) ToVirtualMachineSshCredentialsResponseOutput() VirtualMachineSshCredentialsResponseOutput {
	return o
}

func (o VirtualMachineSshCredentialsResponseOutput) ToVirtualMachineSshCredentialsResponseOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponseOutput {
	return o
}

func (o VirtualMachineSshCredentialsResponseOutput) ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o VirtualMachineSshCredentialsResponseOutput) ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineSshCredentialsResponse) *VirtualMachineSshCredentialsResponse {
		return &v
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o VirtualMachineSshCredentialsResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentialsResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponseOutput) PrivateKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentialsResponse) *string { return v.PrivateKeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponseOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentialsResponse) *string { return v.PublicKeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineSshCredentialsResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type VirtualMachineSshCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineSshCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSshCredentialsResponse)(nil)).Elem()
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) ToVirtualMachineSshCredentialsResponsePtrOutput() VirtualMachineSshCredentialsResponsePtrOutput {
	return o
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) ToVirtualMachineSshCredentialsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineSshCredentialsResponsePtrOutput {
	return o
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) Elem() VirtualMachineSshCredentialsResponseOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) VirtualMachineSshCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineSshCredentialsResponse
		return ret
	}).(VirtualMachineSshCredentialsResponseOutput)
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) PrivateKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKeyData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicKeyData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineSshCredentialsResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSshCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AKSOutput{})
	pulumi.RegisterOutputType(AKSPropertiesOutput{})
	pulumi.RegisterOutputType(AKSPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AKSResponseOutput{})
	pulumi.RegisterOutputType(AKSResponsePropertiesOutput{})
	pulumi.RegisterOutputType(AKSResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AksNetworkingConfigurationOutput{})
	pulumi.RegisterOutputType(AksNetworkingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AksNetworkingConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AksNetworkingConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(AmlComputeOutput{})
	pulumi.RegisterOutputType(AmlComputeNodeInformationResponseOutput{})
	pulumi.RegisterOutputType(AmlComputeNodeInformationResponseArrayOutput{})
	pulumi.RegisterOutputType(AmlComputePropertiesOutput{})
	pulumi.RegisterOutputType(AmlComputePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AmlComputeResponseOutput{})
	pulumi.RegisterOutputType(AmlComputeResponsePropertiesOutput{})
	pulumi.RegisterOutputType(AmlComputeResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DataFactoryOutput{})
	pulumi.RegisterOutputType(DataFactoryResponseOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsPropertiesOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsResponseOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsResponsePropertiesOutput{})
	pulumi.RegisterOutputType(DataLakeAnalyticsResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DatabricksOutput{})
	pulumi.RegisterOutputType(DatabricksPropertiesOutput{})
	pulumi.RegisterOutputType(DatabricksPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DatabricksResponseOutput{})
	pulumi.RegisterOutputType(DatabricksResponsePropertiesOutput{})
	pulumi.RegisterOutputType(DatabricksResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(HDInsightOutput{})
	pulumi.RegisterOutputType(HDInsightPropertiesOutput{})
	pulumi.RegisterOutputType(HDInsightPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HDInsightResponseOutput{})
	pulumi.RegisterOutputType(HDInsightResponsePropertiesOutput{})
	pulumi.RegisterOutputType(HDInsightResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MachineLearningServiceErrorResponseOutput{})
	pulumi.RegisterOutputType(MachineLearningServiceErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(NodeStateCountsResponseOutput{})
	pulumi.RegisterOutputType(NodeStateCountsResponsePtrOutput{})
	pulumi.RegisterOutputType(PasswordResponseOutput{})
	pulumi.RegisterOutputType(PasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(RegistryListCredentialsResultResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdOutput{})
	pulumi.RegisterOutputType(ResourceIdPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(ScaleSettingsOutput{})
	pulumi.RegisterOutputType(ScaleSettingsPtrOutput{})
	pulumi.RegisterOutputType(ScaleSettingsResponseOutput{})
	pulumi.RegisterOutputType(ScaleSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SslConfigurationOutput{})
	pulumi.RegisterOutputType(SslConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SslConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SslConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemServiceResponseOutput{})
	pulumi.RegisterOutputType(SystemServiceResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsPtrOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsResponseOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineOutput{})
	pulumi.RegisterOutputType(VirtualMachinePropertiesOutput{})
	pulumi.RegisterOutputType(VirtualMachinePropertiesPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineResponsePropertiesOutput{})
	pulumi.RegisterOutputType(VirtualMachineResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineSshCredentialsOutput{})
	pulumi.RegisterOutputType(VirtualMachineSshCredentialsPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineSshCredentialsResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineSshCredentialsResponsePtrOutput{})
}
