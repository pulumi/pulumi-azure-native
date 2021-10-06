


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AKS struct {
	ComputeLocation  *string        `pulumi:"computeLocation"`
	ComputeType      string         `pulumi:"computeType"`
	Description      *string        `pulumi:"description"`
	DisableLocalAuth *bool          `pulumi:"disableLocalAuth"`
	Properties       *AKSProperties `pulumi:"properties"`
	ResourceId       *string        `pulumi:"resourceId"`
}





type AKSInput interface {
	pulumi.Input

	ToAKSOutput() AKSOutput
	ToAKSOutputWithContext(context.Context) AKSOutput
}

type AKSArgs struct {
	ComputeLocation  pulumi.StringPtrInput `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput    `pulumi:"computeType"`
	Description      pulumi.StringPtrInput `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput   `pulumi:"disableLocalAuth"`
	Properties       AKSPropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput `pulumi:"resourceId"`
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

func (o AKSOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKS) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o AKSOutput) Properties() AKSPropertiesPtrOutput {
	return o.ApplyT(func(v AKS) *AKSProperties { return v.Properties }).(AKSPropertiesPtrOutput)
}

func (o AKSOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKS) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AKSProperties struct {
	AgentCount                 *int                        `pulumi:"agentCount"`
	AgentVmSize                *string                     `pulumi:"agentVmSize"`
	AksNetworkingConfiguration *AksNetworkingConfiguration `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                     `pulumi:"clusterFqdn"`
	ClusterPurpose             *string                     `pulumi:"clusterPurpose"`
	LoadBalancerSubnet         *string                     `pulumi:"loadBalancerSubnet"`
	LoadBalancerType           *string                     `pulumi:"loadBalancerType"`
	SslConfiguration           *SslConfiguration           `pulumi:"sslConfiguration"`
}





type AKSPropertiesInput interface {
	pulumi.Input

	ToAKSPropertiesOutput() AKSPropertiesOutput
	ToAKSPropertiesOutputWithContext(context.Context) AKSPropertiesOutput
}

type AKSPropertiesArgs struct {
	AgentCount                 pulumi.IntPtrInput                 `pulumi:"agentCount"`
	AgentVmSize                pulumi.StringPtrInput              `pulumi:"agentVmSize"`
	AksNetworkingConfiguration AksNetworkingConfigurationPtrInput `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                pulumi.StringPtrInput              `pulumi:"clusterFqdn"`
	ClusterPurpose             pulumi.StringPtrInput              `pulumi:"clusterPurpose"`
	LoadBalancerSubnet         pulumi.StringPtrInput              `pulumi:"loadBalancerSubnet"`
	LoadBalancerType           pulumi.StringPtrInput              `pulumi:"loadBalancerType"`
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

func (o AKSPropertiesOutput) AgentVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.AgentVmSize }).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesOutput) AksNetworkingConfiguration() AksNetworkingConfigurationPtrOutput {
	return o.ApplyT(func(v AKSProperties) *AksNetworkingConfiguration { return v.AksNetworkingConfiguration }).(AksNetworkingConfigurationPtrOutput)
}

func (o AKSPropertiesOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.ClusterFqdn }).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesOutput) ClusterPurpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.ClusterPurpose }).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesOutput) LoadBalancerSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.LoadBalancerSubnet }).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesOutput) LoadBalancerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSProperties) *string { return v.LoadBalancerType }).(pulumi.StringPtrOutput)
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

func (o AKSPropertiesPtrOutput) AgentVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *string {
		if v == nil {
			return nil
		}
		return v.AgentVmSize
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

func (o AKSPropertiesPtrOutput) ClusterPurpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterPurpose
	}).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesPtrOutput) LoadBalancerSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSubnet
	}).(pulumi.StringPtrOutput)
}

func (o AKSPropertiesPtrOutput) LoadBalancerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSProperties) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerType
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
	ComputeLocation    *string                 `pulumi:"computeLocation"`
	ComputeType        string                  `pulumi:"computeType"`
	CreatedOn          string                  `pulumi:"createdOn"`
	Description        *string                 `pulumi:"description"`
	DisableLocalAuth   *bool                   `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                    `pulumi:"isAttachedCompute"`
	ModifiedOn         string                  `pulumi:"modifiedOn"`
	Properties         *AKSResponseProperties  `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                  `pulumi:"provisioningState"`
	ResourceId         *string                 `pulumi:"resourceId"`
}





type AKSResponseInput interface {
	pulumi.Input

	ToAKSResponseOutput() AKSResponseOutput
	ToAKSResponseOutputWithContext(context.Context) AKSResponseOutput
}

type AKSResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput           `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput              `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput              `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput           `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput             `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput              `pulumi:"modifiedOn"`
	Properties         AKSResponsePropertiesPtrInput   `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput              `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput           `pulumi:"resourceId"`
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

func (o AKSResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AKSResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
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

func (o AKSResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v AKSResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o AKSResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AKSResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AKSResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AKSResponseProperties struct {
	AgentCount                 *int                                `pulumi:"agentCount"`
	AgentVmSize                *string                             `pulumi:"agentVmSize"`
	AksNetworkingConfiguration *AksNetworkingConfigurationResponse `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                *string                             `pulumi:"clusterFqdn"`
	ClusterPurpose             *string                             `pulumi:"clusterPurpose"`
	LoadBalancerSubnet         *string                             `pulumi:"loadBalancerSubnet"`
	LoadBalancerType           *string                             `pulumi:"loadBalancerType"`
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
	AgentVmSize                pulumi.StringPtrInput                      `pulumi:"agentVmSize"`
	AksNetworkingConfiguration AksNetworkingConfigurationResponsePtrInput `pulumi:"aksNetworkingConfiguration"`
	ClusterFqdn                pulumi.StringPtrInput                      `pulumi:"clusterFqdn"`
	ClusterPurpose             pulumi.StringPtrInput                      `pulumi:"clusterPurpose"`
	LoadBalancerSubnet         pulumi.StringPtrInput                      `pulumi:"loadBalancerSubnet"`
	LoadBalancerType           pulumi.StringPtrInput                      `pulumi:"loadBalancerType"`
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

func (o AKSResponsePropertiesOutput) AgentVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.AgentVmSize }).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesOutput) AksNetworkingConfiguration() AksNetworkingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *AksNetworkingConfigurationResponse { return v.AksNetworkingConfiguration }).(AksNetworkingConfigurationResponsePtrOutput)
}

func (o AKSResponsePropertiesOutput) ClusterFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.ClusterFqdn }).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesOutput) ClusterPurpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.ClusterPurpose }).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesOutput) LoadBalancerSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.LoadBalancerSubnet }).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesOutput) LoadBalancerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AKSResponseProperties) *string { return v.LoadBalancerType }).(pulumi.StringPtrOutput)
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

func (o AKSResponsePropertiesPtrOutput) AgentVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.AgentVmSize
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

func (o AKSResponsePropertiesPtrOutput) ClusterPurpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterPurpose
	}).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) LoadBalancerSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSubnet
	}).(pulumi.StringPtrOutput)
}

func (o AKSResponsePropertiesPtrOutput) LoadBalancerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AKSResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerType
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
	ComputeLocation  *string               `pulumi:"computeLocation"`
	ComputeType      string                `pulumi:"computeType"`
	Description      *string               `pulumi:"description"`
	DisableLocalAuth *bool                 `pulumi:"disableLocalAuth"`
	Properties       *AmlComputeProperties `pulumi:"properties"`
	ResourceId       *string               `pulumi:"resourceId"`
}





type AmlComputeInput interface {
	pulumi.Input

	ToAmlComputeOutput() AmlComputeOutput
	ToAmlComputeOutputWithContext(context.Context) AmlComputeOutput
}

type AmlComputeArgs struct {
	ComputeLocation  pulumi.StringPtrInput        `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput           `pulumi:"computeType"`
	Description      pulumi.StringPtrInput        `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput          `pulumi:"disableLocalAuth"`
	Properties       AmlComputePropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput        `pulumi:"resourceId"`
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

func (o AmlComputeOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlCompute) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
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
	EnableNodePublicIp          *bool                   `pulumi:"enableNodePublicIp"`
	IsolatedNetwork             *bool                   `pulumi:"isolatedNetwork"`
	OsType                      *string                 `pulumi:"osType"`
	RemoteLoginPortPublicAccess *string                 `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings               *ScaleSettings          `pulumi:"scaleSettings"`
	Subnet                      *ResourceId             `pulumi:"subnet"`
	UserAccountCredentials      *UserAccountCredentials `pulumi:"userAccountCredentials"`
	VirtualMachineImage         *VirtualMachineImage    `pulumi:"virtualMachineImage"`
	VmPriority                  *string                 `pulumi:"vmPriority"`
	VmSize                      *string                 `pulumi:"vmSize"`
}





type AmlComputePropertiesInput interface {
	pulumi.Input

	ToAmlComputePropertiesOutput() AmlComputePropertiesOutput
	ToAmlComputePropertiesOutputWithContext(context.Context) AmlComputePropertiesOutput
}

type AmlComputePropertiesArgs struct {
	EnableNodePublicIp          pulumi.BoolPtrInput            `pulumi:"enableNodePublicIp"`
	IsolatedNetwork             pulumi.BoolPtrInput            `pulumi:"isolatedNetwork"`
	OsType                      pulumi.StringPtrInput          `pulumi:"osType"`
	RemoteLoginPortPublicAccess pulumi.StringPtrInput          `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings               ScaleSettingsPtrInput          `pulumi:"scaleSettings"`
	Subnet                      ResourceIdPtrInput             `pulumi:"subnet"`
	UserAccountCredentials      UserAccountCredentialsPtrInput `pulumi:"userAccountCredentials"`
	VirtualMachineImage         VirtualMachineImagePtrInput    `pulumi:"virtualMachineImage"`
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

func (o AmlComputePropertiesOutput) EnableNodePublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *bool { return v.EnableNodePublicIp }).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesOutput) IsolatedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *bool { return v.IsolatedNetwork }).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *string { return v.OsType }).(pulumi.StringPtrOutput)
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

func (o AmlComputePropertiesOutput) VirtualMachineImage() VirtualMachineImagePtrOutput {
	return o.ApplyT(func(v AmlComputeProperties) *VirtualMachineImage { return v.VirtualMachineImage }).(VirtualMachineImagePtrOutput)
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

func (o AmlComputePropertiesPtrOutput) EnableNodePublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNodePublicIp
	}).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) IsolatedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsolatedNetwork
	}).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
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

func (o AmlComputePropertiesPtrOutput) VirtualMachineImage() VirtualMachineImagePtrOutput {
	return o.ApplyT(func(v *AmlComputeProperties) *VirtualMachineImage {
		if v == nil {
			return nil
		}
		return v.VirtualMachineImage
	}).(VirtualMachineImagePtrOutput)
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

type AmlComputePropertiesResponse struct {
	AllocationState               string                          `pulumi:"allocationState"`
	AllocationStateTransitionTime string                          `pulumi:"allocationStateTransitionTime"`
	CurrentNodeCount              int                             `pulumi:"currentNodeCount"`
	EnableNodePublicIp            *bool                           `pulumi:"enableNodePublicIp"`
	Errors                        []ErrorResponseResponse         `pulumi:"errors"`
	IsolatedNetwork               *bool                           `pulumi:"isolatedNetwork"`
	NodeStateCounts               NodeStateCountsResponse         `pulumi:"nodeStateCounts"`
	OsType                        *string                         `pulumi:"osType"`
	RemoteLoginPortPublicAccess   *string                         `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 *ScaleSettingsResponse          `pulumi:"scaleSettings"`
	Subnet                        *ResourceIdResponse             `pulumi:"subnet"`
	TargetNodeCount               int                             `pulumi:"targetNodeCount"`
	UserAccountCredentials        *UserAccountCredentialsResponse `pulumi:"userAccountCredentials"`
	VirtualMachineImage           *VirtualMachineImageResponse    `pulumi:"virtualMachineImage"`
	VmPriority                    *string                         `pulumi:"vmPriority"`
	VmSize                        *string                         `pulumi:"vmSize"`
}





type AmlComputePropertiesResponseInput interface {
	pulumi.Input

	ToAmlComputePropertiesResponseOutput() AmlComputePropertiesResponseOutput
	ToAmlComputePropertiesResponseOutputWithContext(context.Context) AmlComputePropertiesResponseOutput
}

type AmlComputePropertiesResponseArgs struct {
	AllocationState               pulumi.StringInput                     `pulumi:"allocationState"`
	AllocationStateTransitionTime pulumi.StringInput                     `pulumi:"allocationStateTransitionTime"`
	CurrentNodeCount              pulumi.IntInput                        `pulumi:"currentNodeCount"`
	EnableNodePublicIp            pulumi.BoolPtrInput                    `pulumi:"enableNodePublicIp"`
	Errors                        ErrorResponseResponseArrayInput        `pulumi:"errors"`
	IsolatedNetwork               pulumi.BoolPtrInput                    `pulumi:"isolatedNetwork"`
	NodeStateCounts               NodeStateCountsResponseInput           `pulumi:"nodeStateCounts"`
	OsType                        pulumi.StringPtrInput                  `pulumi:"osType"`
	RemoteLoginPortPublicAccess   pulumi.StringPtrInput                  `pulumi:"remoteLoginPortPublicAccess"`
	ScaleSettings                 ScaleSettingsResponsePtrInput          `pulumi:"scaleSettings"`
	Subnet                        ResourceIdResponsePtrInput             `pulumi:"subnet"`
	TargetNodeCount               pulumi.IntInput                        `pulumi:"targetNodeCount"`
	UserAccountCredentials        UserAccountCredentialsResponsePtrInput `pulumi:"userAccountCredentials"`
	VirtualMachineImage           VirtualMachineImageResponsePtrInput    `pulumi:"virtualMachineImage"`
	VmPriority                    pulumi.StringPtrInput                  `pulumi:"vmPriority"`
	VmSize                        pulumi.StringPtrInput                  `pulumi:"vmSize"`
}

func (AmlComputePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputePropertiesResponse)(nil)).Elem()
}

func (i AmlComputePropertiesResponseArgs) ToAmlComputePropertiesResponseOutput() AmlComputePropertiesResponseOutput {
	return i.ToAmlComputePropertiesResponseOutputWithContext(context.Background())
}

func (i AmlComputePropertiesResponseArgs) ToAmlComputePropertiesResponseOutputWithContext(ctx context.Context) AmlComputePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesResponseOutput)
}

func (i AmlComputePropertiesResponseArgs) ToAmlComputePropertiesResponsePtrOutput() AmlComputePropertiesResponsePtrOutput {
	return i.ToAmlComputePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AmlComputePropertiesResponseArgs) ToAmlComputePropertiesResponsePtrOutputWithContext(ctx context.Context) AmlComputePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesResponseOutput).ToAmlComputePropertiesResponsePtrOutputWithContext(ctx)
}









type AmlComputePropertiesResponsePtrInput interface {
	pulumi.Input

	ToAmlComputePropertiesResponsePtrOutput() AmlComputePropertiesResponsePtrOutput
	ToAmlComputePropertiesResponsePtrOutputWithContext(context.Context) AmlComputePropertiesResponsePtrOutput
}

type amlComputePropertiesResponsePtrType AmlComputePropertiesResponseArgs

func AmlComputePropertiesResponsePtr(v *AmlComputePropertiesResponseArgs) AmlComputePropertiesResponsePtrInput {
	return (*amlComputePropertiesResponsePtrType)(v)
}

func (*amlComputePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputePropertiesResponse)(nil)).Elem()
}

func (i *amlComputePropertiesResponsePtrType) ToAmlComputePropertiesResponsePtrOutput() AmlComputePropertiesResponsePtrOutput {
	return i.ToAmlComputePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *amlComputePropertiesResponsePtrType) ToAmlComputePropertiesResponsePtrOutputWithContext(ctx context.Context) AmlComputePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmlComputePropertiesResponsePtrOutput)
}

type AmlComputePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AmlComputePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmlComputePropertiesResponse)(nil)).Elem()
}

func (o AmlComputePropertiesResponseOutput) ToAmlComputePropertiesResponseOutput() AmlComputePropertiesResponseOutput {
	return o
}

func (o AmlComputePropertiesResponseOutput) ToAmlComputePropertiesResponseOutputWithContext(ctx context.Context) AmlComputePropertiesResponseOutput {
	return o
}

func (o AmlComputePropertiesResponseOutput) ToAmlComputePropertiesResponsePtrOutput() AmlComputePropertiesResponsePtrOutput {
	return o.ToAmlComputePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AmlComputePropertiesResponseOutput) ToAmlComputePropertiesResponsePtrOutputWithContext(ctx context.Context) AmlComputePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AmlComputePropertiesResponse) *AmlComputePropertiesResponse {
		return &v
	}).(AmlComputePropertiesResponsePtrOutput)
}

func (o AmlComputePropertiesResponseOutput) AllocationState() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) string { return v.AllocationState }).(pulumi.StringOutput)
}

func (o AmlComputePropertiesResponseOutput) AllocationStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) string { return v.AllocationStateTransitionTime }).(pulumi.StringOutput)
}

func (o AmlComputePropertiesResponseOutput) CurrentNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) int { return v.CurrentNodeCount }).(pulumi.IntOutput)
}

func (o AmlComputePropertiesResponseOutput) EnableNodePublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *bool { return v.EnableNodePublicIp }).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesResponseOutput) Errors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) []ErrorResponseResponse { return v.Errors }).(ErrorResponseResponseArrayOutput)
}

func (o AmlComputePropertiesResponseOutput) IsolatedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *bool { return v.IsolatedNetwork }).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesResponseOutput) NodeStateCounts() NodeStateCountsResponseOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) NodeStateCountsResponse { return v.NodeStateCounts }).(NodeStateCountsResponseOutput)
}

func (o AmlComputePropertiesResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesResponseOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *string { return v.RemoteLoginPortPublicAccess }).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesResponseOutput) ScaleSettings() ScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *ScaleSettingsResponse { return v.ScaleSettings }).(ScaleSettingsResponsePtrOutput)
}

func (o AmlComputePropertiesResponseOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *ResourceIdResponse { return v.Subnet }).(ResourceIdResponsePtrOutput)
}

func (o AmlComputePropertiesResponseOutput) TargetNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) int { return v.TargetNodeCount }).(pulumi.IntOutput)
}

func (o AmlComputePropertiesResponseOutput) UserAccountCredentials() UserAccountCredentialsResponsePtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *UserAccountCredentialsResponse { return v.UserAccountCredentials }).(UserAccountCredentialsResponsePtrOutput)
}

func (o AmlComputePropertiesResponseOutput) VirtualMachineImage() VirtualMachineImageResponsePtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *VirtualMachineImageResponse { return v.VirtualMachineImage }).(VirtualMachineImageResponsePtrOutput)
}

func (o AmlComputePropertiesResponseOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *string { return v.VmPriority }).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputePropertiesResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type AmlComputePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AmlComputePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmlComputePropertiesResponse)(nil)).Elem()
}

func (o AmlComputePropertiesResponsePtrOutput) ToAmlComputePropertiesResponsePtrOutput() AmlComputePropertiesResponsePtrOutput {
	return o
}

func (o AmlComputePropertiesResponsePtrOutput) ToAmlComputePropertiesResponsePtrOutputWithContext(ctx context.Context) AmlComputePropertiesResponsePtrOutput {
	return o
}

func (o AmlComputePropertiesResponsePtrOutput) Elem() AmlComputePropertiesResponseOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) AmlComputePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AmlComputePropertiesResponse
		return ret
	}).(AmlComputePropertiesResponseOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) AllocationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AllocationState
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) AllocationStateTransitionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AllocationStateTransitionTime
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) CurrentNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.CurrentNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) EnableNodePublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableNodePublicIp
	}).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) Errors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) []ErrorResponseResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(ErrorResponseResponseArrayOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) IsolatedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsolatedNetwork
	}).(pulumi.BoolPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) NodeStateCounts() NodeStateCountsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *NodeStateCountsResponse {
		if v == nil {
			return nil
		}
		return &v.NodeStateCounts
	}).(NodeStateCountsResponsePtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) RemoteLoginPortPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemoteLoginPortPublicAccess
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) ScaleSettings() ScaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *ScaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ScaleSettings
	}).(ScaleSettingsResponsePtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ResourceIdResponsePtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) TargetNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TargetNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) UserAccountCredentials() UserAccountCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *UserAccountCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.UserAccountCredentials
	}).(UserAccountCredentialsResponsePtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) VirtualMachineImage() VirtualMachineImageResponsePtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *VirtualMachineImageResponse {
		if v == nil {
			return nil
		}
		return v.VirtualMachineImage
	}).(VirtualMachineImageResponsePtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) VmPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmPriority
	}).(pulumi.StringPtrOutput)
}

func (o AmlComputePropertiesResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AmlComputePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type AmlComputeResponse struct {
	ComputeLocation    *string                       `pulumi:"computeLocation"`
	ComputeType        string                        `pulumi:"computeType"`
	CreatedOn          string                        `pulumi:"createdOn"`
	Description        *string                       `pulumi:"description"`
	DisableLocalAuth   *bool                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                        `pulumi:"modifiedOn"`
	Properties         *AmlComputePropertiesResponse `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse       `pulumi:"provisioningErrors"`
	ProvisioningState  string                        `pulumi:"provisioningState"`
	ResourceId         *string                       `pulumi:"resourceId"`
}





type AmlComputeResponseInput interface {
	pulumi.Input

	ToAmlComputeResponseOutput() AmlComputeResponseOutput
	ToAmlComputeResponseOutputWithContext(context.Context) AmlComputeResponseOutput
}

type AmlComputeResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                   `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                   `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput                  `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                     `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                   `pulumi:"modifiedOn"`
	Properties         AmlComputePropertiesResponsePtrInput `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput      `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                   `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                `pulumi:"resourceId"`
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

func (o AmlComputeResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o AmlComputeResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v AmlComputeResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o AmlComputeResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) Properties() AmlComputePropertiesResponsePtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *AmlComputePropertiesResponse { return v.Properties }).(AmlComputePropertiesResponsePtrOutput)
}

func (o AmlComputeResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v AmlComputeResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o AmlComputeResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AmlComputeResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AmlComputeResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmlComputeResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type AssignedUser struct {
	ObjectId string `pulumi:"objectId"`
	TenantId string `pulumi:"tenantId"`
}





type AssignedUserInput interface {
	pulumi.Input

	ToAssignedUserOutput() AssignedUserOutput
	ToAssignedUserOutputWithContext(context.Context) AssignedUserOutput
}

type AssignedUserArgs struct {
	ObjectId pulumi.StringInput `pulumi:"objectId"`
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (AssignedUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedUser)(nil)).Elem()
}

func (i AssignedUserArgs) ToAssignedUserOutput() AssignedUserOutput {
	return i.ToAssignedUserOutputWithContext(context.Background())
}

func (i AssignedUserArgs) ToAssignedUserOutputWithContext(ctx context.Context) AssignedUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedUserOutput)
}

func (i AssignedUserArgs) ToAssignedUserPtrOutput() AssignedUserPtrOutput {
	return i.ToAssignedUserPtrOutputWithContext(context.Background())
}

func (i AssignedUserArgs) ToAssignedUserPtrOutputWithContext(ctx context.Context) AssignedUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedUserOutput).ToAssignedUserPtrOutputWithContext(ctx)
}









type AssignedUserPtrInput interface {
	pulumi.Input

	ToAssignedUserPtrOutput() AssignedUserPtrOutput
	ToAssignedUserPtrOutputWithContext(context.Context) AssignedUserPtrOutput
}

type assignedUserPtrType AssignedUserArgs

func AssignedUserPtr(v *AssignedUserArgs) AssignedUserPtrInput {
	return (*assignedUserPtrType)(v)
}

func (*assignedUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedUser)(nil)).Elem()
}

func (i *assignedUserPtrType) ToAssignedUserPtrOutput() AssignedUserPtrOutput {
	return i.ToAssignedUserPtrOutputWithContext(context.Background())
}

func (i *assignedUserPtrType) ToAssignedUserPtrOutputWithContext(ctx context.Context) AssignedUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedUserPtrOutput)
}

type AssignedUserOutput struct{ *pulumi.OutputState }

func (AssignedUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedUser)(nil)).Elem()
}

func (o AssignedUserOutput) ToAssignedUserOutput() AssignedUserOutput {
	return o
}

func (o AssignedUserOutput) ToAssignedUserOutputWithContext(ctx context.Context) AssignedUserOutput {
	return o
}

func (o AssignedUserOutput) ToAssignedUserPtrOutput() AssignedUserPtrOutput {
	return o.ToAssignedUserPtrOutputWithContext(context.Background())
}

func (o AssignedUserOutput) ToAssignedUserPtrOutputWithContext(ctx context.Context) AssignedUserPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignedUser) *AssignedUser {
		return &v
	}).(AssignedUserPtrOutput)
}

func (o AssignedUserOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v AssignedUser) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o AssignedUserOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v AssignedUser) string { return v.TenantId }).(pulumi.StringOutput)
}

type AssignedUserPtrOutput struct{ *pulumi.OutputState }

func (AssignedUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedUser)(nil)).Elem()
}

func (o AssignedUserPtrOutput) ToAssignedUserPtrOutput() AssignedUserPtrOutput {
	return o
}

func (o AssignedUserPtrOutput) ToAssignedUserPtrOutputWithContext(ctx context.Context) AssignedUserPtrOutput {
	return o
}

func (o AssignedUserPtrOutput) Elem() AssignedUserOutput {
	return o.ApplyT(func(v *AssignedUser) AssignedUser {
		if v != nil {
			return *v
		}
		var ret AssignedUser
		return ret
	}).(AssignedUserOutput)
}

func (o AssignedUserPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedUser) *string {
		if v == nil {
			return nil
		}
		return &v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o AssignedUserPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedUser) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AssignedUserResponse struct {
	ObjectId string `pulumi:"objectId"`
	TenantId string `pulumi:"tenantId"`
}





type AssignedUserResponseInput interface {
	pulumi.Input

	ToAssignedUserResponseOutput() AssignedUserResponseOutput
	ToAssignedUserResponseOutputWithContext(context.Context) AssignedUserResponseOutput
}

type AssignedUserResponseArgs struct {
	ObjectId pulumi.StringInput `pulumi:"objectId"`
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (AssignedUserResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedUserResponse)(nil)).Elem()
}

func (i AssignedUserResponseArgs) ToAssignedUserResponseOutput() AssignedUserResponseOutput {
	return i.ToAssignedUserResponseOutputWithContext(context.Background())
}

func (i AssignedUserResponseArgs) ToAssignedUserResponseOutputWithContext(ctx context.Context) AssignedUserResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedUserResponseOutput)
}

func (i AssignedUserResponseArgs) ToAssignedUserResponsePtrOutput() AssignedUserResponsePtrOutput {
	return i.ToAssignedUserResponsePtrOutputWithContext(context.Background())
}

func (i AssignedUserResponseArgs) ToAssignedUserResponsePtrOutputWithContext(ctx context.Context) AssignedUserResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedUserResponseOutput).ToAssignedUserResponsePtrOutputWithContext(ctx)
}









type AssignedUserResponsePtrInput interface {
	pulumi.Input

	ToAssignedUserResponsePtrOutput() AssignedUserResponsePtrOutput
	ToAssignedUserResponsePtrOutputWithContext(context.Context) AssignedUserResponsePtrOutput
}

type assignedUserResponsePtrType AssignedUserResponseArgs

func AssignedUserResponsePtr(v *AssignedUserResponseArgs) AssignedUserResponsePtrInput {
	return (*assignedUserResponsePtrType)(v)
}

func (*assignedUserResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedUserResponse)(nil)).Elem()
}

func (i *assignedUserResponsePtrType) ToAssignedUserResponsePtrOutput() AssignedUserResponsePtrOutput {
	return i.ToAssignedUserResponsePtrOutputWithContext(context.Background())
}

func (i *assignedUserResponsePtrType) ToAssignedUserResponsePtrOutputWithContext(ctx context.Context) AssignedUserResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedUserResponsePtrOutput)
}

type AssignedUserResponseOutput struct{ *pulumi.OutputState }

func (AssignedUserResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedUserResponse)(nil)).Elem()
}

func (o AssignedUserResponseOutput) ToAssignedUserResponseOutput() AssignedUserResponseOutput {
	return o
}

func (o AssignedUserResponseOutput) ToAssignedUserResponseOutputWithContext(ctx context.Context) AssignedUserResponseOutput {
	return o
}

func (o AssignedUserResponseOutput) ToAssignedUserResponsePtrOutput() AssignedUserResponsePtrOutput {
	return o.ToAssignedUserResponsePtrOutputWithContext(context.Background())
}

func (o AssignedUserResponseOutput) ToAssignedUserResponsePtrOutputWithContext(ctx context.Context) AssignedUserResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignedUserResponse) *AssignedUserResponse {
		return &v
	}).(AssignedUserResponsePtrOutput)
}

func (o AssignedUserResponseOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v AssignedUserResponse) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o AssignedUserResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v AssignedUserResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type AssignedUserResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignedUserResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedUserResponse)(nil)).Elem()
}

func (o AssignedUserResponsePtrOutput) ToAssignedUserResponsePtrOutput() AssignedUserResponsePtrOutput {
	return o
}

func (o AssignedUserResponsePtrOutput) ToAssignedUserResponsePtrOutputWithContext(ctx context.Context) AssignedUserResponsePtrOutput {
	return o
}

func (o AssignedUserResponsePtrOutput) Elem() AssignedUserResponseOutput {
	return o.ApplyT(func(v *AssignedUserResponse) AssignedUserResponse {
		if v != nil {
			return *v
		}
		var ret AssignedUserResponse
		return ret
	}).(AssignedUserResponseOutput)
}

func (o AssignedUserResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedUserResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o AssignedUserResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedUserResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AutoPauseProperties struct {
	DelayInMinutes *int  `pulumi:"delayInMinutes"`
	Enabled        *bool `pulumi:"enabled"`
}





type AutoPausePropertiesInput interface {
	pulumi.Input

	ToAutoPausePropertiesOutput() AutoPausePropertiesOutput
	ToAutoPausePropertiesOutputWithContext(context.Context) AutoPausePropertiesOutput
}

type AutoPausePropertiesArgs struct {
	DelayInMinutes pulumi.IntPtrInput  `pulumi:"delayInMinutes"`
	Enabled        pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (AutoPausePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPauseProperties)(nil)).Elem()
}

func (i AutoPausePropertiesArgs) ToAutoPausePropertiesOutput() AutoPausePropertiesOutput {
	return i.ToAutoPausePropertiesOutputWithContext(context.Background())
}

func (i AutoPausePropertiesArgs) ToAutoPausePropertiesOutputWithContext(ctx context.Context) AutoPausePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesOutput)
}

func (i AutoPausePropertiesArgs) ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput {
	return i.ToAutoPausePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoPausePropertiesArgs) ToAutoPausePropertiesPtrOutputWithContext(ctx context.Context) AutoPausePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesOutput).ToAutoPausePropertiesPtrOutputWithContext(ctx)
}









type AutoPausePropertiesPtrInput interface {
	pulumi.Input

	ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput
	ToAutoPausePropertiesPtrOutputWithContext(context.Context) AutoPausePropertiesPtrOutput
}

type autoPausePropertiesPtrType AutoPausePropertiesArgs

func AutoPausePropertiesPtr(v *AutoPausePropertiesArgs) AutoPausePropertiesPtrInput {
	return (*autoPausePropertiesPtrType)(v)
}

func (*autoPausePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPauseProperties)(nil)).Elem()
}

func (i *autoPausePropertiesPtrType) ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput {
	return i.ToAutoPausePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoPausePropertiesPtrType) ToAutoPausePropertiesPtrOutputWithContext(ctx context.Context) AutoPausePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesPtrOutput)
}

type AutoPausePropertiesOutput struct{ *pulumi.OutputState }

func (AutoPausePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPauseProperties)(nil)).Elem()
}

func (o AutoPausePropertiesOutput) ToAutoPausePropertiesOutput() AutoPausePropertiesOutput {
	return o
}

func (o AutoPausePropertiesOutput) ToAutoPausePropertiesOutputWithContext(ctx context.Context) AutoPausePropertiesOutput {
	return o
}

func (o AutoPausePropertiesOutput) ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput {
	return o.ToAutoPausePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoPausePropertiesOutput) ToAutoPausePropertiesPtrOutputWithContext(ctx context.Context) AutoPausePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoPauseProperties) *AutoPauseProperties {
		return &v
	}).(AutoPausePropertiesPtrOutput)
}

func (o AutoPausePropertiesOutput) DelayInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoPauseProperties) *int { return v.DelayInMinutes }).(pulumi.IntPtrOutput)
}

func (o AutoPausePropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoPauseProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type AutoPausePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoPausePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPauseProperties)(nil)).Elem()
}

func (o AutoPausePropertiesPtrOutput) ToAutoPausePropertiesPtrOutput() AutoPausePropertiesPtrOutput {
	return o
}

func (o AutoPausePropertiesPtrOutput) ToAutoPausePropertiesPtrOutputWithContext(ctx context.Context) AutoPausePropertiesPtrOutput {
	return o
}

func (o AutoPausePropertiesPtrOutput) Elem() AutoPausePropertiesOutput {
	return o.ApplyT(func(v *AutoPauseProperties) AutoPauseProperties {
		if v != nil {
			return *v
		}
		var ret AutoPauseProperties
		return ret
	}).(AutoPausePropertiesOutput)
}

func (o AutoPausePropertiesPtrOutput) DelayInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoPauseProperties) *int {
		if v == nil {
			return nil
		}
		return v.DelayInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o AutoPausePropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoPauseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type AutoPausePropertiesResponse struct {
	DelayInMinutes *int  `pulumi:"delayInMinutes"`
	Enabled        *bool `pulumi:"enabled"`
}





type AutoPausePropertiesResponseInput interface {
	pulumi.Input

	ToAutoPausePropertiesResponseOutput() AutoPausePropertiesResponseOutput
	ToAutoPausePropertiesResponseOutputWithContext(context.Context) AutoPausePropertiesResponseOutput
}

type AutoPausePropertiesResponseArgs struct {
	DelayInMinutes pulumi.IntPtrInput  `pulumi:"delayInMinutes"`
	Enabled        pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (AutoPausePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPausePropertiesResponse)(nil)).Elem()
}

func (i AutoPausePropertiesResponseArgs) ToAutoPausePropertiesResponseOutput() AutoPausePropertiesResponseOutput {
	return i.ToAutoPausePropertiesResponseOutputWithContext(context.Background())
}

func (i AutoPausePropertiesResponseArgs) ToAutoPausePropertiesResponseOutputWithContext(ctx context.Context) AutoPausePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesResponseOutput)
}

func (i AutoPausePropertiesResponseArgs) ToAutoPausePropertiesResponsePtrOutput() AutoPausePropertiesResponsePtrOutput {
	return i.ToAutoPausePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AutoPausePropertiesResponseArgs) ToAutoPausePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoPausePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesResponseOutput).ToAutoPausePropertiesResponsePtrOutputWithContext(ctx)
}









type AutoPausePropertiesResponsePtrInput interface {
	pulumi.Input

	ToAutoPausePropertiesResponsePtrOutput() AutoPausePropertiesResponsePtrOutput
	ToAutoPausePropertiesResponsePtrOutputWithContext(context.Context) AutoPausePropertiesResponsePtrOutput
}

type autoPausePropertiesResponsePtrType AutoPausePropertiesResponseArgs

func AutoPausePropertiesResponsePtr(v *AutoPausePropertiesResponseArgs) AutoPausePropertiesResponsePtrInput {
	return (*autoPausePropertiesResponsePtrType)(v)
}

func (*autoPausePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPausePropertiesResponse)(nil)).Elem()
}

func (i *autoPausePropertiesResponsePtrType) ToAutoPausePropertiesResponsePtrOutput() AutoPausePropertiesResponsePtrOutput {
	return i.ToAutoPausePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *autoPausePropertiesResponsePtrType) ToAutoPausePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoPausePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPausePropertiesResponsePtrOutput)
}

type AutoPausePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoPausePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPausePropertiesResponse)(nil)).Elem()
}

func (o AutoPausePropertiesResponseOutput) ToAutoPausePropertiesResponseOutput() AutoPausePropertiesResponseOutput {
	return o
}

func (o AutoPausePropertiesResponseOutput) ToAutoPausePropertiesResponseOutputWithContext(ctx context.Context) AutoPausePropertiesResponseOutput {
	return o
}

func (o AutoPausePropertiesResponseOutput) ToAutoPausePropertiesResponsePtrOutput() AutoPausePropertiesResponsePtrOutput {
	return o.ToAutoPausePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AutoPausePropertiesResponseOutput) ToAutoPausePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoPausePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoPausePropertiesResponse) *AutoPausePropertiesResponse {
		return &v
	}).(AutoPausePropertiesResponsePtrOutput)
}

func (o AutoPausePropertiesResponseOutput) DelayInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoPausePropertiesResponse) *int { return v.DelayInMinutes }).(pulumi.IntPtrOutput)
}

func (o AutoPausePropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoPausePropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type AutoPausePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoPausePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPausePropertiesResponse)(nil)).Elem()
}

func (o AutoPausePropertiesResponsePtrOutput) ToAutoPausePropertiesResponsePtrOutput() AutoPausePropertiesResponsePtrOutput {
	return o
}

func (o AutoPausePropertiesResponsePtrOutput) ToAutoPausePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoPausePropertiesResponsePtrOutput {
	return o
}

func (o AutoPausePropertiesResponsePtrOutput) Elem() AutoPausePropertiesResponseOutput {
	return o.ApplyT(func(v *AutoPausePropertiesResponse) AutoPausePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutoPausePropertiesResponse
		return ret
	}).(AutoPausePropertiesResponseOutput)
}

func (o AutoPausePropertiesResponsePtrOutput) DelayInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoPausePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.DelayInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o AutoPausePropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoPausePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type AutoScaleProperties struct {
	Enabled      *bool `pulumi:"enabled"`
	MaxNodeCount *int  `pulumi:"maxNodeCount"`
	MinNodeCount *int  `pulumi:"minNodeCount"`
}





type AutoScalePropertiesInput interface {
	pulumi.Input

	ToAutoScalePropertiesOutput() AutoScalePropertiesOutput
	ToAutoScalePropertiesOutputWithContext(context.Context) AutoScalePropertiesOutput
}

type AutoScalePropertiesArgs struct {
	Enabled      pulumi.BoolPtrInput `pulumi:"enabled"`
	MaxNodeCount pulumi.IntPtrInput  `pulumi:"maxNodeCount"`
	MinNodeCount pulumi.IntPtrInput  `pulumi:"minNodeCount"`
}

func (AutoScalePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleProperties)(nil)).Elem()
}

func (i AutoScalePropertiesArgs) ToAutoScalePropertiesOutput() AutoScalePropertiesOutput {
	return i.ToAutoScalePropertiesOutputWithContext(context.Background())
}

func (i AutoScalePropertiesArgs) ToAutoScalePropertiesOutputWithContext(ctx context.Context) AutoScalePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesOutput)
}

func (i AutoScalePropertiesArgs) ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput {
	return i.ToAutoScalePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoScalePropertiesArgs) ToAutoScalePropertiesPtrOutputWithContext(ctx context.Context) AutoScalePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesOutput).ToAutoScalePropertiesPtrOutputWithContext(ctx)
}









type AutoScalePropertiesPtrInput interface {
	pulumi.Input

	ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput
	ToAutoScalePropertiesPtrOutputWithContext(context.Context) AutoScalePropertiesPtrOutput
}

type autoScalePropertiesPtrType AutoScalePropertiesArgs

func AutoScalePropertiesPtr(v *AutoScalePropertiesArgs) AutoScalePropertiesPtrInput {
	return (*autoScalePropertiesPtrType)(v)
}

func (*autoScalePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleProperties)(nil)).Elem()
}

func (i *autoScalePropertiesPtrType) ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput {
	return i.ToAutoScalePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoScalePropertiesPtrType) ToAutoScalePropertiesPtrOutputWithContext(ctx context.Context) AutoScalePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesPtrOutput)
}

type AutoScalePropertiesOutput struct{ *pulumi.OutputState }

func (AutoScalePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScaleProperties)(nil)).Elem()
}

func (o AutoScalePropertiesOutput) ToAutoScalePropertiesOutput() AutoScalePropertiesOutput {
	return o
}

func (o AutoScalePropertiesOutput) ToAutoScalePropertiesOutputWithContext(ctx context.Context) AutoScalePropertiesOutput {
	return o
}

func (o AutoScalePropertiesOutput) ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput {
	return o.ToAutoScalePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoScalePropertiesOutput) ToAutoScalePropertiesPtrOutputWithContext(ctx context.Context) AutoScalePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScaleProperties) *AutoScaleProperties {
		return &v
	}).(AutoScalePropertiesPtrOutput)
}

func (o AutoScalePropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoScaleProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AutoScalePropertiesOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleProperties) *int { return v.MaxNodeCount }).(pulumi.IntPtrOutput)
}

func (o AutoScalePropertiesOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScaleProperties) *int { return v.MinNodeCount }).(pulumi.IntPtrOutput)
}

type AutoScalePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoScalePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScaleProperties)(nil)).Elem()
}

func (o AutoScalePropertiesPtrOutput) ToAutoScalePropertiesPtrOutput() AutoScalePropertiesPtrOutput {
	return o
}

func (o AutoScalePropertiesPtrOutput) ToAutoScalePropertiesPtrOutputWithContext(ctx context.Context) AutoScalePropertiesPtrOutput {
	return o
}

func (o AutoScalePropertiesPtrOutput) Elem() AutoScalePropertiesOutput {
	return o.ApplyT(func(v *AutoScaleProperties) AutoScaleProperties {
		if v != nil {
			return *v
		}
		var ret AutoScaleProperties
		return ret
	}).(AutoScalePropertiesOutput)
}

func (o AutoScalePropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoScaleProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AutoScalePropertiesPtrOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleProperties) *int {
		if v == nil {
			return nil
		}
		return v.MaxNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoScalePropertiesPtrOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScaleProperties) *int {
		if v == nil {
			return nil
		}
		return v.MinNodeCount
	}).(pulumi.IntPtrOutput)
}

type AutoScalePropertiesResponse struct {
	Enabled      *bool `pulumi:"enabled"`
	MaxNodeCount *int  `pulumi:"maxNodeCount"`
	MinNodeCount *int  `pulumi:"minNodeCount"`
}





type AutoScalePropertiesResponseInput interface {
	pulumi.Input

	ToAutoScalePropertiesResponseOutput() AutoScalePropertiesResponseOutput
	ToAutoScalePropertiesResponseOutputWithContext(context.Context) AutoScalePropertiesResponseOutput
}

type AutoScalePropertiesResponseArgs struct {
	Enabled      pulumi.BoolPtrInput `pulumi:"enabled"`
	MaxNodeCount pulumi.IntPtrInput  `pulumi:"maxNodeCount"`
	MinNodeCount pulumi.IntPtrInput  `pulumi:"minNodeCount"`
}

func (AutoScalePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalePropertiesResponse)(nil)).Elem()
}

func (i AutoScalePropertiesResponseArgs) ToAutoScalePropertiesResponseOutput() AutoScalePropertiesResponseOutput {
	return i.ToAutoScalePropertiesResponseOutputWithContext(context.Background())
}

func (i AutoScalePropertiesResponseArgs) ToAutoScalePropertiesResponseOutputWithContext(ctx context.Context) AutoScalePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesResponseOutput)
}

func (i AutoScalePropertiesResponseArgs) ToAutoScalePropertiesResponsePtrOutput() AutoScalePropertiesResponsePtrOutput {
	return i.ToAutoScalePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AutoScalePropertiesResponseArgs) ToAutoScalePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoScalePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesResponseOutput).ToAutoScalePropertiesResponsePtrOutputWithContext(ctx)
}









type AutoScalePropertiesResponsePtrInput interface {
	pulumi.Input

	ToAutoScalePropertiesResponsePtrOutput() AutoScalePropertiesResponsePtrOutput
	ToAutoScalePropertiesResponsePtrOutputWithContext(context.Context) AutoScalePropertiesResponsePtrOutput
}

type autoScalePropertiesResponsePtrType AutoScalePropertiesResponseArgs

func AutoScalePropertiesResponsePtr(v *AutoScalePropertiesResponseArgs) AutoScalePropertiesResponsePtrInput {
	return (*autoScalePropertiesResponsePtrType)(v)
}

func (*autoScalePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalePropertiesResponse)(nil)).Elem()
}

func (i *autoScalePropertiesResponsePtrType) ToAutoScalePropertiesResponsePtrOutput() AutoScalePropertiesResponsePtrOutput {
	return i.ToAutoScalePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *autoScalePropertiesResponsePtrType) ToAutoScalePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoScalePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalePropertiesResponsePtrOutput)
}

type AutoScalePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoScalePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalePropertiesResponse)(nil)).Elem()
}

func (o AutoScalePropertiesResponseOutput) ToAutoScalePropertiesResponseOutput() AutoScalePropertiesResponseOutput {
	return o
}

func (o AutoScalePropertiesResponseOutput) ToAutoScalePropertiesResponseOutputWithContext(ctx context.Context) AutoScalePropertiesResponseOutput {
	return o
}

func (o AutoScalePropertiesResponseOutput) ToAutoScalePropertiesResponsePtrOutput() AutoScalePropertiesResponsePtrOutput {
	return o.ToAutoScalePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AutoScalePropertiesResponseOutput) ToAutoScalePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoScalePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScalePropertiesResponse) *AutoScalePropertiesResponse {
		return &v
	}).(AutoScalePropertiesResponsePtrOutput)
}

func (o AutoScalePropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoScalePropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AutoScalePropertiesResponseOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScalePropertiesResponse) *int { return v.MaxNodeCount }).(pulumi.IntPtrOutput)
}

func (o AutoScalePropertiesResponseOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoScalePropertiesResponse) *int { return v.MinNodeCount }).(pulumi.IntPtrOutput)
}

type AutoScalePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoScalePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalePropertiesResponse)(nil)).Elem()
}

func (o AutoScalePropertiesResponsePtrOutput) ToAutoScalePropertiesResponsePtrOutput() AutoScalePropertiesResponsePtrOutput {
	return o
}

func (o AutoScalePropertiesResponsePtrOutput) ToAutoScalePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoScalePropertiesResponsePtrOutput {
	return o
}

func (o AutoScalePropertiesResponsePtrOutput) Elem() AutoScalePropertiesResponseOutput {
	return o.ApplyT(func(v *AutoScalePropertiesResponse) AutoScalePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutoScalePropertiesResponse
		return ret
	}).(AutoScalePropertiesResponseOutput)
}

func (o AutoScalePropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoScalePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AutoScalePropertiesResponsePtrOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxNodeCount
	}).(pulumi.IntPtrOutput)
}

func (o AutoScalePropertiesResponsePtrOutput) MinNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoScalePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinNodeCount
	}).(pulumi.IntPtrOutput)
}

type ComputeInstance struct {
	ComputeLocation  *string                    `pulumi:"computeLocation"`
	ComputeType      string                     `pulumi:"computeType"`
	Description      *string                    `pulumi:"description"`
	DisableLocalAuth *bool                      `pulumi:"disableLocalAuth"`
	Properties       *ComputeInstanceProperties `pulumi:"properties"`
	ResourceId       *string                    `pulumi:"resourceId"`
}





type ComputeInstanceInput interface {
	pulumi.Input

	ToComputeInstanceOutput() ComputeInstanceOutput
	ToComputeInstanceOutputWithContext(context.Context) ComputeInstanceOutput
}

type ComputeInstanceArgs struct {
	ComputeLocation  pulumi.StringPtrInput             `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput                `pulumi:"computeType"`
	Description      pulumi.StringPtrInput             `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput               `pulumi:"disableLocalAuth"`
	Properties       ComputeInstancePropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput             `pulumi:"resourceId"`
}

func (ComputeInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstance)(nil)).Elem()
}

func (i ComputeInstanceArgs) ToComputeInstanceOutput() ComputeInstanceOutput {
	return i.ToComputeInstanceOutputWithContext(context.Background())
}

func (i ComputeInstanceArgs) ToComputeInstanceOutputWithContext(ctx context.Context) ComputeInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceOutput)
}

type ComputeInstanceOutput struct{ *pulumi.OutputState }

func (ComputeInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstance)(nil)).Elem()
}

func (o ComputeInstanceOutput) ToComputeInstanceOutput() ComputeInstanceOutput {
	return o
}

func (o ComputeInstanceOutput) ToComputeInstanceOutputWithContext(ctx context.Context) ComputeInstanceOutput {
	return o
}

func (o ComputeInstanceOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstance) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstance) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o ComputeInstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstance) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComputeInstance) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o ComputeInstanceOutput) Properties() ComputeInstancePropertiesPtrOutput {
	return o.ApplyT(func(v ComputeInstance) *ComputeInstanceProperties { return v.Properties }).(ComputeInstancePropertiesPtrOutput)
}

func (o ComputeInstanceOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstance) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ComputeInstanceApplicationResponse struct {
	DisplayName *string `pulumi:"displayName"`
	EndpointUri *string `pulumi:"endpointUri"`
}





type ComputeInstanceApplicationResponseInput interface {
	pulumi.Input

	ToComputeInstanceApplicationResponseOutput() ComputeInstanceApplicationResponseOutput
	ToComputeInstanceApplicationResponseOutputWithContext(context.Context) ComputeInstanceApplicationResponseOutput
}

type ComputeInstanceApplicationResponseArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	EndpointUri pulumi.StringPtrInput `pulumi:"endpointUri"`
}

func (ComputeInstanceApplicationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceApplicationResponse)(nil)).Elem()
}

func (i ComputeInstanceApplicationResponseArgs) ToComputeInstanceApplicationResponseOutput() ComputeInstanceApplicationResponseOutput {
	return i.ToComputeInstanceApplicationResponseOutputWithContext(context.Background())
}

func (i ComputeInstanceApplicationResponseArgs) ToComputeInstanceApplicationResponseOutputWithContext(ctx context.Context) ComputeInstanceApplicationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceApplicationResponseOutput)
}





type ComputeInstanceApplicationResponseArrayInput interface {
	pulumi.Input

	ToComputeInstanceApplicationResponseArrayOutput() ComputeInstanceApplicationResponseArrayOutput
	ToComputeInstanceApplicationResponseArrayOutputWithContext(context.Context) ComputeInstanceApplicationResponseArrayOutput
}

type ComputeInstanceApplicationResponseArray []ComputeInstanceApplicationResponseInput

func (ComputeInstanceApplicationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputeInstanceApplicationResponse)(nil)).Elem()
}

func (i ComputeInstanceApplicationResponseArray) ToComputeInstanceApplicationResponseArrayOutput() ComputeInstanceApplicationResponseArrayOutput {
	return i.ToComputeInstanceApplicationResponseArrayOutputWithContext(context.Background())
}

func (i ComputeInstanceApplicationResponseArray) ToComputeInstanceApplicationResponseArrayOutputWithContext(ctx context.Context) ComputeInstanceApplicationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceApplicationResponseArrayOutput)
}

type ComputeInstanceApplicationResponseOutput struct{ *pulumi.OutputState }

func (ComputeInstanceApplicationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceApplicationResponse)(nil)).Elem()
}

func (o ComputeInstanceApplicationResponseOutput) ToComputeInstanceApplicationResponseOutput() ComputeInstanceApplicationResponseOutput {
	return o
}

func (o ComputeInstanceApplicationResponseOutput) ToComputeInstanceApplicationResponseOutputWithContext(ctx context.Context) ComputeInstanceApplicationResponseOutput {
	return o
}

func (o ComputeInstanceApplicationResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceApplicationResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceApplicationResponseOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceApplicationResponse) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

type ComputeInstanceApplicationResponseArrayOutput struct{ *pulumi.OutputState }

func (ComputeInstanceApplicationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputeInstanceApplicationResponse)(nil)).Elem()
}

func (o ComputeInstanceApplicationResponseArrayOutput) ToComputeInstanceApplicationResponseArrayOutput() ComputeInstanceApplicationResponseArrayOutput {
	return o
}

func (o ComputeInstanceApplicationResponseArrayOutput) ToComputeInstanceApplicationResponseArrayOutputWithContext(ctx context.Context) ComputeInstanceApplicationResponseArrayOutput {
	return o
}

func (o ComputeInstanceApplicationResponseArrayOutput) Index(i pulumi.IntInput) ComputeInstanceApplicationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ComputeInstanceApplicationResponse {
		return vs[0].([]ComputeInstanceApplicationResponse)[vs[1].(int)]
	}).(ComputeInstanceApplicationResponseOutput)
}

type ComputeInstanceConnectivityEndpointsResponse struct {
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	PublicIpAddress  string `pulumi:"publicIpAddress"`
}





type ComputeInstanceConnectivityEndpointsResponseInput interface {
	pulumi.Input

	ToComputeInstanceConnectivityEndpointsResponseOutput() ComputeInstanceConnectivityEndpointsResponseOutput
	ToComputeInstanceConnectivityEndpointsResponseOutputWithContext(context.Context) ComputeInstanceConnectivityEndpointsResponseOutput
}

type ComputeInstanceConnectivityEndpointsResponseArgs struct {
	PrivateIpAddress pulumi.StringInput `pulumi:"privateIpAddress"`
	PublicIpAddress  pulumi.StringInput `pulumi:"publicIpAddress"`
}

func (ComputeInstanceConnectivityEndpointsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceConnectivityEndpointsResponse)(nil)).Elem()
}

func (i ComputeInstanceConnectivityEndpointsResponseArgs) ToComputeInstanceConnectivityEndpointsResponseOutput() ComputeInstanceConnectivityEndpointsResponseOutput {
	return i.ToComputeInstanceConnectivityEndpointsResponseOutputWithContext(context.Background())
}

func (i ComputeInstanceConnectivityEndpointsResponseArgs) ToComputeInstanceConnectivityEndpointsResponseOutputWithContext(ctx context.Context) ComputeInstanceConnectivityEndpointsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceConnectivityEndpointsResponseOutput)
}

func (i ComputeInstanceConnectivityEndpointsResponseArgs) ToComputeInstanceConnectivityEndpointsResponsePtrOutput() ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return i.ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i ComputeInstanceConnectivityEndpointsResponseArgs) ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceConnectivityEndpointsResponseOutput).ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(ctx)
}









type ComputeInstanceConnectivityEndpointsResponsePtrInput interface {
	pulumi.Input

	ToComputeInstanceConnectivityEndpointsResponsePtrOutput() ComputeInstanceConnectivityEndpointsResponsePtrOutput
	ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(context.Context) ComputeInstanceConnectivityEndpointsResponsePtrOutput
}

type computeInstanceConnectivityEndpointsResponsePtrType ComputeInstanceConnectivityEndpointsResponseArgs

func ComputeInstanceConnectivityEndpointsResponsePtr(v *ComputeInstanceConnectivityEndpointsResponseArgs) ComputeInstanceConnectivityEndpointsResponsePtrInput {
	return (*computeInstanceConnectivityEndpointsResponsePtrType)(v)
}

func (*computeInstanceConnectivityEndpointsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceConnectivityEndpointsResponse)(nil)).Elem()
}

func (i *computeInstanceConnectivityEndpointsResponsePtrType) ToComputeInstanceConnectivityEndpointsResponsePtrOutput() ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return i.ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i *computeInstanceConnectivityEndpointsResponsePtrType) ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceConnectivityEndpointsResponsePtrOutput)
}

type ComputeInstanceConnectivityEndpointsResponseOutput struct{ *pulumi.OutputState }

func (ComputeInstanceConnectivityEndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceConnectivityEndpointsResponse)(nil)).Elem()
}

func (o ComputeInstanceConnectivityEndpointsResponseOutput) ToComputeInstanceConnectivityEndpointsResponseOutput() ComputeInstanceConnectivityEndpointsResponseOutput {
	return o
}

func (o ComputeInstanceConnectivityEndpointsResponseOutput) ToComputeInstanceConnectivityEndpointsResponseOutputWithContext(ctx context.Context) ComputeInstanceConnectivityEndpointsResponseOutput {
	return o
}

func (o ComputeInstanceConnectivityEndpointsResponseOutput) ToComputeInstanceConnectivityEndpointsResponsePtrOutput() ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return o.ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o ComputeInstanceConnectivityEndpointsResponseOutput) ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeInstanceConnectivityEndpointsResponse) *ComputeInstanceConnectivityEndpointsResponse {
		return &v
	}).(ComputeInstanceConnectivityEndpointsResponsePtrOutput)
}

func (o ComputeInstanceConnectivityEndpointsResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceConnectivityEndpointsResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o ComputeInstanceConnectivityEndpointsResponseOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceConnectivityEndpointsResponse) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

type ComputeInstanceConnectivityEndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeInstanceConnectivityEndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceConnectivityEndpointsResponse)(nil)).Elem()
}

func (o ComputeInstanceConnectivityEndpointsResponsePtrOutput) ToComputeInstanceConnectivityEndpointsResponsePtrOutput() ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return o
}

func (o ComputeInstanceConnectivityEndpointsResponsePtrOutput) ToComputeInstanceConnectivityEndpointsResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return o
}

func (o ComputeInstanceConnectivityEndpointsResponsePtrOutput) Elem() ComputeInstanceConnectivityEndpointsResponseOutput {
	return o.ApplyT(func(v *ComputeInstanceConnectivityEndpointsResponse) ComputeInstanceConnectivityEndpointsResponse {
		if v != nil {
			return *v
		}
		var ret ComputeInstanceConnectivityEndpointsResponse
		return ret
	}).(ComputeInstanceConnectivityEndpointsResponseOutput)
}

func (o ComputeInstanceConnectivityEndpointsResponsePtrOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceConnectivityEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrivateIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceConnectivityEndpointsResponsePtrOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceConnectivityEndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublicIpAddress
	}).(pulumi.StringPtrOutput)
}

type ComputeInstanceCreatedByResponse struct {
	UserId    string `pulumi:"userId"`
	UserName  string `pulumi:"userName"`
	UserOrgId string `pulumi:"userOrgId"`
}





type ComputeInstanceCreatedByResponseInput interface {
	pulumi.Input

	ToComputeInstanceCreatedByResponseOutput() ComputeInstanceCreatedByResponseOutput
	ToComputeInstanceCreatedByResponseOutputWithContext(context.Context) ComputeInstanceCreatedByResponseOutput
}

type ComputeInstanceCreatedByResponseArgs struct {
	UserId    pulumi.StringInput `pulumi:"userId"`
	UserName  pulumi.StringInput `pulumi:"userName"`
	UserOrgId pulumi.StringInput `pulumi:"userOrgId"`
}

func (ComputeInstanceCreatedByResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceCreatedByResponse)(nil)).Elem()
}

func (i ComputeInstanceCreatedByResponseArgs) ToComputeInstanceCreatedByResponseOutput() ComputeInstanceCreatedByResponseOutput {
	return i.ToComputeInstanceCreatedByResponseOutputWithContext(context.Background())
}

func (i ComputeInstanceCreatedByResponseArgs) ToComputeInstanceCreatedByResponseOutputWithContext(ctx context.Context) ComputeInstanceCreatedByResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceCreatedByResponseOutput)
}

func (i ComputeInstanceCreatedByResponseArgs) ToComputeInstanceCreatedByResponsePtrOutput() ComputeInstanceCreatedByResponsePtrOutput {
	return i.ToComputeInstanceCreatedByResponsePtrOutputWithContext(context.Background())
}

func (i ComputeInstanceCreatedByResponseArgs) ToComputeInstanceCreatedByResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceCreatedByResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceCreatedByResponseOutput).ToComputeInstanceCreatedByResponsePtrOutputWithContext(ctx)
}









type ComputeInstanceCreatedByResponsePtrInput interface {
	pulumi.Input

	ToComputeInstanceCreatedByResponsePtrOutput() ComputeInstanceCreatedByResponsePtrOutput
	ToComputeInstanceCreatedByResponsePtrOutputWithContext(context.Context) ComputeInstanceCreatedByResponsePtrOutput
}

type computeInstanceCreatedByResponsePtrType ComputeInstanceCreatedByResponseArgs

func ComputeInstanceCreatedByResponsePtr(v *ComputeInstanceCreatedByResponseArgs) ComputeInstanceCreatedByResponsePtrInput {
	return (*computeInstanceCreatedByResponsePtrType)(v)
}

func (*computeInstanceCreatedByResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceCreatedByResponse)(nil)).Elem()
}

func (i *computeInstanceCreatedByResponsePtrType) ToComputeInstanceCreatedByResponsePtrOutput() ComputeInstanceCreatedByResponsePtrOutput {
	return i.ToComputeInstanceCreatedByResponsePtrOutputWithContext(context.Background())
}

func (i *computeInstanceCreatedByResponsePtrType) ToComputeInstanceCreatedByResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceCreatedByResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceCreatedByResponsePtrOutput)
}

type ComputeInstanceCreatedByResponseOutput struct{ *pulumi.OutputState }

func (ComputeInstanceCreatedByResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceCreatedByResponse)(nil)).Elem()
}

func (o ComputeInstanceCreatedByResponseOutput) ToComputeInstanceCreatedByResponseOutput() ComputeInstanceCreatedByResponseOutput {
	return o
}

func (o ComputeInstanceCreatedByResponseOutput) ToComputeInstanceCreatedByResponseOutputWithContext(ctx context.Context) ComputeInstanceCreatedByResponseOutput {
	return o
}

func (o ComputeInstanceCreatedByResponseOutput) ToComputeInstanceCreatedByResponsePtrOutput() ComputeInstanceCreatedByResponsePtrOutput {
	return o.ToComputeInstanceCreatedByResponsePtrOutputWithContext(context.Background())
}

func (o ComputeInstanceCreatedByResponseOutput) ToComputeInstanceCreatedByResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceCreatedByResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeInstanceCreatedByResponse) *ComputeInstanceCreatedByResponse {
		return &v
	}).(ComputeInstanceCreatedByResponsePtrOutput)
}

func (o ComputeInstanceCreatedByResponseOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceCreatedByResponse) string { return v.UserId }).(pulumi.StringOutput)
}

func (o ComputeInstanceCreatedByResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceCreatedByResponse) string { return v.UserName }).(pulumi.StringOutput)
}

func (o ComputeInstanceCreatedByResponseOutput) UserOrgId() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceCreatedByResponse) string { return v.UserOrgId }).(pulumi.StringOutput)
}

type ComputeInstanceCreatedByResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeInstanceCreatedByResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceCreatedByResponse)(nil)).Elem()
}

func (o ComputeInstanceCreatedByResponsePtrOutput) ToComputeInstanceCreatedByResponsePtrOutput() ComputeInstanceCreatedByResponsePtrOutput {
	return o
}

func (o ComputeInstanceCreatedByResponsePtrOutput) ToComputeInstanceCreatedByResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceCreatedByResponsePtrOutput {
	return o
}

func (o ComputeInstanceCreatedByResponsePtrOutput) Elem() ComputeInstanceCreatedByResponseOutput {
	return o.ApplyT(func(v *ComputeInstanceCreatedByResponse) ComputeInstanceCreatedByResponse {
		if v != nil {
			return *v
		}
		var ret ComputeInstanceCreatedByResponse
		return ret
	}).(ComputeInstanceCreatedByResponseOutput)
}

func (o ComputeInstanceCreatedByResponsePtrOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceCreatedByResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserId
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceCreatedByResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceCreatedByResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserName
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceCreatedByResponsePtrOutput) UserOrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceCreatedByResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserOrgId
	}).(pulumi.StringPtrOutput)
}

type ComputeInstanceLastOperationResponse struct {
	OperationName   *string `pulumi:"operationName"`
	OperationStatus *string `pulumi:"operationStatus"`
	OperationTime   *string `pulumi:"operationTime"`
}





type ComputeInstanceLastOperationResponseInput interface {
	pulumi.Input

	ToComputeInstanceLastOperationResponseOutput() ComputeInstanceLastOperationResponseOutput
	ToComputeInstanceLastOperationResponseOutputWithContext(context.Context) ComputeInstanceLastOperationResponseOutput
}

type ComputeInstanceLastOperationResponseArgs struct {
	OperationName   pulumi.StringPtrInput `pulumi:"operationName"`
	OperationStatus pulumi.StringPtrInput `pulumi:"operationStatus"`
	OperationTime   pulumi.StringPtrInput `pulumi:"operationTime"`
}

func (ComputeInstanceLastOperationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceLastOperationResponse)(nil)).Elem()
}

func (i ComputeInstanceLastOperationResponseArgs) ToComputeInstanceLastOperationResponseOutput() ComputeInstanceLastOperationResponseOutput {
	return i.ToComputeInstanceLastOperationResponseOutputWithContext(context.Background())
}

func (i ComputeInstanceLastOperationResponseArgs) ToComputeInstanceLastOperationResponseOutputWithContext(ctx context.Context) ComputeInstanceLastOperationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceLastOperationResponseOutput)
}

func (i ComputeInstanceLastOperationResponseArgs) ToComputeInstanceLastOperationResponsePtrOutput() ComputeInstanceLastOperationResponsePtrOutput {
	return i.ToComputeInstanceLastOperationResponsePtrOutputWithContext(context.Background())
}

func (i ComputeInstanceLastOperationResponseArgs) ToComputeInstanceLastOperationResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceLastOperationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceLastOperationResponseOutput).ToComputeInstanceLastOperationResponsePtrOutputWithContext(ctx)
}









type ComputeInstanceLastOperationResponsePtrInput interface {
	pulumi.Input

	ToComputeInstanceLastOperationResponsePtrOutput() ComputeInstanceLastOperationResponsePtrOutput
	ToComputeInstanceLastOperationResponsePtrOutputWithContext(context.Context) ComputeInstanceLastOperationResponsePtrOutput
}

type computeInstanceLastOperationResponsePtrType ComputeInstanceLastOperationResponseArgs

func ComputeInstanceLastOperationResponsePtr(v *ComputeInstanceLastOperationResponseArgs) ComputeInstanceLastOperationResponsePtrInput {
	return (*computeInstanceLastOperationResponsePtrType)(v)
}

func (*computeInstanceLastOperationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceLastOperationResponse)(nil)).Elem()
}

func (i *computeInstanceLastOperationResponsePtrType) ToComputeInstanceLastOperationResponsePtrOutput() ComputeInstanceLastOperationResponsePtrOutput {
	return i.ToComputeInstanceLastOperationResponsePtrOutputWithContext(context.Background())
}

func (i *computeInstanceLastOperationResponsePtrType) ToComputeInstanceLastOperationResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceLastOperationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceLastOperationResponsePtrOutput)
}

type ComputeInstanceLastOperationResponseOutput struct{ *pulumi.OutputState }

func (ComputeInstanceLastOperationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceLastOperationResponse)(nil)).Elem()
}

func (o ComputeInstanceLastOperationResponseOutput) ToComputeInstanceLastOperationResponseOutput() ComputeInstanceLastOperationResponseOutput {
	return o
}

func (o ComputeInstanceLastOperationResponseOutput) ToComputeInstanceLastOperationResponseOutputWithContext(ctx context.Context) ComputeInstanceLastOperationResponseOutput {
	return o
}

func (o ComputeInstanceLastOperationResponseOutput) ToComputeInstanceLastOperationResponsePtrOutput() ComputeInstanceLastOperationResponsePtrOutput {
	return o.ToComputeInstanceLastOperationResponsePtrOutputWithContext(context.Background())
}

func (o ComputeInstanceLastOperationResponseOutput) ToComputeInstanceLastOperationResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceLastOperationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeInstanceLastOperationResponse) *ComputeInstanceLastOperationResponse {
		return &v
	}).(ComputeInstanceLastOperationResponsePtrOutput)
}

func (o ComputeInstanceLastOperationResponseOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceLastOperationResponse) *string { return v.OperationName }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceLastOperationResponseOutput) OperationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceLastOperationResponse) *string { return v.OperationStatus }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceLastOperationResponseOutput) OperationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceLastOperationResponse) *string { return v.OperationTime }).(pulumi.StringPtrOutput)
}

type ComputeInstanceLastOperationResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeInstanceLastOperationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceLastOperationResponse)(nil)).Elem()
}

func (o ComputeInstanceLastOperationResponsePtrOutput) ToComputeInstanceLastOperationResponsePtrOutput() ComputeInstanceLastOperationResponsePtrOutput {
	return o
}

func (o ComputeInstanceLastOperationResponsePtrOutput) ToComputeInstanceLastOperationResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceLastOperationResponsePtrOutput {
	return o
}

func (o ComputeInstanceLastOperationResponsePtrOutput) Elem() ComputeInstanceLastOperationResponseOutput {
	return o.ApplyT(func(v *ComputeInstanceLastOperationResponse) ComputeInstanceLastOperationResponse {
		if v != nil {
			return *v
		}
		var ret ComputeInstanceLastOperationResponse
		return ret
	}).(ComputeInstanceLastOperationResponseOutput)
}

func (o ComputeInstanceLastOperationResponsePtrOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceLastOperationResponse) *string {
		if v == nil {
			return nil
		}
		return v.OperationName
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceLastOperationResponsePtrOutput) OperationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceLastOperationResponse) *string {
		if v == nil {
			return nil
		}
		return v.OperationStatus
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceLastOperationResponsePtrOutput) OperationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceLastOperationResponse) *string {
		if v == nil {
			return nil
		}
		return v.OperationTime
	}).(pulumi.StringPtrOutput)
}

type ComputeInstanceProperties struct {
	ApplicationSharingPolicy         *string                          `pulumi:"applicationSharingPolicy"`
	ComputeInstanceAuthorizationType *string                          `pulumi:"computeInstanceAuthorizationType"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettings `pulumi:"personalComputeInstanceSettings"`
	SetupScripts                     *SetupScripts                    `pulumi:"setupScripts"`
	SshSettings                      *ComputeInstanceSshSettings      `pulumi:"sshSettings"`
	Subnet                           *ResourceId                      `pulumi:"subnet"`
	VmSize                           *string                          `pulumi:"vmSize"`
}





type ComputeInstancePropertiesInput interface {
	pulumi.Input

	ToComputeInstancePropertiesOutput() ComputeInstancePropertiesOutput
	ToComputeInstancePropertiesOutputWithContext(context.Context) ComputeInstancePropertiesOutput
}

type ComputeInstancePropertiesArgs struct {
	ApplicationSharingPolicy         pulumi.StringPtrInput                   `pulumi:"applicationSharingPolicy"`
	ComputeInstanceAuthorizationType pulumi.StringPtrInput                   `pulumi:"computeInstanceAuthorizationType"`
	PersonalComputeInstanceSettings  PersonalComputeInstanceSettingsPtrInput `pulumi:"personalComputeInstanceSettings"`
	SetupScripts                     SetupScriptsPtrInput                    `pulumi:"setupScripts"`
	SshSettings                      ComputeInstanceSshSettingsPtrInput      `pulumi:"sshSettings"`
	Subnet                           ResourceIdPtrInput                      `pulumi:"subnet"`
	VmSize                           pulumi.StringPtrInput                   `pulumi:"vmSize"`
}

func (ComputeInstancePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceProperties)(nil)).Elem()
}

func (i ComputeInstancePropertiesArgs) ToComputeInstancePropertiesOutput() ComputeInstancePropertiesOutput {
	return i.ToComputeInstancePropertiesOutputWithContext(context.Background())
}

func (i ComputeInstancePropertiesArgs) ToComputeInstancePropertiesOutputWithContext(ctx context.Context) ComputeInstancePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstancePropertiesOutput)
}

func (i ComputeInstancePropertiesArgs) ToComputeInstancePropertiesPtrOutput() ComputeInstancePropertiesPtrOutput {
	return i.ToComputeInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i ComputeInstancePropertiesArgs) ToComputeInstancePropertiesPtrOutputWithContext(ctx context.Context) ComputeInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstancePropertiesOutput).ToComputeInstancePropertiesPtrOutputWithContext(ctx)
}









type ComputeInstancePropertiesPtrInput interface {
	pulumi.Input

	ToComputeInstancePropertiesPtrOutput() ComputeInstancePropertiesPtrOutput
	ToComputeInstancePropertiesPtrOutputWithContext(context.Context) ComputeInstancePropertiesPtrOutput
}

type computeInstancePropertiesPtrType ComputeInstancePropertiesArgs

func ComputeInstancePropertiesPtr(v *ComputeInstancePropertiesArgs) ComputeInstancePropertiesPtrInput {
	return (*computeInstancePropertiesPtrType)(v)
}

func (*computeInstancePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceProperties)(nil)).Elem()
}

func (i *computeInstancePropertiesPtrType) ToComputeInstancePropertiesPtrOutput() ComputeInstancePropertiesPtrOutput {
	return i.ToComputeInstancePropertiesPtrOutputWithContext(context.Background())
}

func (i *computeInstancePropertiesPtrType) ToComputeInstancePropertiesPtrOutputWithContext(ctx context.Context) ComputeInstancePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstancePropertiesPtrOutput)
}

type ComputeInstancePropertiesOutput struct{ *pulumi.OutputState }

func (ComputeInstancePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceProperties)(nil)).Elem()
}

func (o ComputeInstancePropertiesOutput) ToComputeInstancePropertiesOutput() ComputeInstancePropertiesOutput {
	return o
}

func (o ComputeInstancePropertiesOutput) ToComputeInstancePropertiesOutputWithContext(ctx context.Context) ComputeInstancePropertiesOutput {
	return o
}

func (o ComputeInstancePropertiesOutput) ToComputeInstancePropertiesPtrOutput() ComputeInstancePropertiesPtrOutput {
	return o.ToComputeInstancePropertiesPtrOutputWithContext(context.Background())
}

func (o ComputeInstancePropertiesOutput) ToComputeInstancePropertiesPtrOutputWithContext(ctx context.Context) ComputeInstancePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeInstanceProperties) *ComputeInstanceProperties {
		return &v
	}).(ComputeInstancePropertiesPtrOutput)
}

func (o ComputeInstancePropertiesOutput) ApplicationSharingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceProperties) *string { return v.ApplicationSharingPolicy }).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesOutput) ComputeInstanceAuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceProperties) *string { return v.ComputeInstanceAuthorizationType }).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesOutput) PersonalComputeInstanceSettings() PersonalComputeInstanceSettingsPtrOutput {
	return o.ApplyT(func(v ComputeInstanceProperties) *PersonalComputeInstanceSettings {
		return v.PersonalComputeInstanceSettings
	}).(PersonalComputeInstanceSettingsPtrOutput)
}

func (o ComputeInstancePropertiesOutput) SetupScripts() SetupScriptsPtrOutput {
	return o.ApplyT(func(v ComputeInstanceProperties) *SetupScripts { return v.SetupScripts }).(SetupScriptsPtrOutput)
}

func (o ComputeInstancePropertiesOutput) SshSettings() ComputeInstanceSshSettingsPtrOutput {
	return o.ApplyT(func(v ComputeInstanceProperties) *ComputeInstanceSshSettings { return v.SshSettings }).(ComputeInstanceSshSettingsPtrOutput)
}

func (o ComputeInstancePropertiesOutput) Subnet() ResourceIdPtrOutput {
	return o.ApplyT(func(v ComputeInstanceProperties) *ResourceId { return v.Subnet }).(ResourceIdPtrOutput)
}

func (o ComputeInstancePropertiesOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceProperties) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type ComputeInstancePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ComputeInstancePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceProperties)(nil)).Elem()
}

func (o ComputeInstancePropertiesPtrOutput) ToComputeInstancePropertiesPtrOutput() ComputeInstancePropertiesPtrOutput {
	return o
}

func (o ComputeInstancePropertiesPtrOutput) ToComputeInstancePropertiesPtrOutputWithContext(ctx context.Context) ComputeInstancePropertiesPtrOutput {
	return o
}

func (o ComputeInstancePropertiesPtrOutput) Elem() ComputeInstancePropertiesOutput {
	return o.ApplyT(func(v *ComputeInstanceProperties) ComputeInstanceProperties {
		if v != nil {
			return *v
		}
		var ret ComputeInstanceProperties
		return ret
	}).(ComputeInstancePropertiesOutput)
}

func (o ComputeInstancePropertiesPtrOutput) ApplicationSharingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationSharingPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesPtrOutput) ComputeInstanceAuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ComputeInstanceAuthorizationType
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesPtrOutput) PersonalComputeInstanceSettings() PersonalComputeInstanceSettingsPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceProperties) *PersonalComputeInstanceSettings {
		if v == nil {
			return nil
		}
		return v.PersonalComputeInstanceSettings
	}).(PersonalComputeInstanceSettingsPtrOutput)
}

func (o ComputeInstancePropertiesPtrOutput) SetupScripts() SetupScriptsPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceProperties) *SetupScripts {
		if v == nil {
			return nil
		}
		return v.SetupScripts
	}).(SetupScriptsPtrOutput)
}

func (o ComputeInstancePropertiesPtrOutput) SshSettings() ComputeInstanceSshSettingsPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceProperties) *ComputeInstanceSshSettings {
		if v == nil {
			return nil
		}
		return v.SshSettings
	}).(ComputeInstanceSshSettingsPtrOutput)
}

func (o ComputeInstancePropertiesPtrOutput) Subnet() ResourceIdPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceProperties) *ResourceId {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ResourceIdPtrOutput)
}

func (o ComputeInstancePropertiesPtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type ComputeInstancePropertiesResponse struct {
	ApplicationSharingPolicy         *string                                      `pulumi:"applicationSharingPolicy"`
	Applications                     []ComputeInstanceApplicationResponse         `pulumi:"applications"`
	ComputeInstanceAuthorizationType *string                                      `pulumi:"computeInstanceAuthorizationType"`
	ConnectivityEndpoints            ComputeInstanceConnectivityEndpointsResponse `pulumi:"connectivityEndpoints"`
	CreatedBy                        ComputeInstanceCreatedByResponse             `pulumi:"createdBy"`
	Errors                           []ErrorResponseResponse                      `pulumi:"errors"`
	LastOperation                    ComputeInstanceLastOperationResponse         `pulumi:"lastOperation"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettingsResponse     `pulumi:"personalComputeInstanceSettings"`
	SetupScripts                     *SetupScriptsResponse                        `pulumi:"setupScripts"`
	SshSettings                      *ComputeInstanceSshSettingsResponse          `pulumi:"sshSettings"`
	State                            string                                       `pulumi:"state"`
	Subnet                           *ResourceIdResponse                          `pulumi:"subnet"`
	VmSize                           *string                                      `pulumi:"vmSize"`
}





type ComputeInstancePropertiesResponseInput interface {
	pulumi.Input

	ToComputeInstancePropertiesResponseOutput() ComputeInstancePropertiesResponseOutput
	ToComputeInstancePropertiesResponseOutputWithContext(context.Context) ComputeInstancePropertiesResponseOutput
}

type ComputeInstancePropertiesResponseArgs struct {
	ApplicationSharingPolicy         pulumi.StringPtrInput                             `pulumi:"applicationSharingPolicy"`
	Applications                     ComputeInstanceApplicationResponseArrayInput      `pulumi:"applications"`
	ComputeInstanceAuthorizationType pulumi.StringPtrInput                             `pulumi:"computeInstanceAuthorizationType"`
	ConnectivityEndpoints            ComputeInstanceConnectivityEndpointsResponseInput `pulumi:"connectivityEndpoints"`
	CreatedBy                        ComputeInstanceCreatedByResponseInput             `pulumi:"createdBy"`
	Errors                           ErrorResponseResponseArrayInput                   `pulumi:"errors"`
	LastOperation                    ComputeInstanceLastOperationResponseInput         `pulumi:"lastOperation"`
	PersonalComputeInstanceSettings  PersonalComputeInstanceSettingsResponsePtrInput   `pulumi:"personalComputeInstanceSettings"`
	SetupScripts                     SetupScriptsResponsePtrInput                      `pulumi:"setupScripts"`
	SshSettings                      ComputeInstanceSshSettingsResponsePtrInput        `pulumi:"sshSettings"`
	State                            pulumi.StringInput                                `pulumi:"state"`
	Subnet                           ResourceIdResponsePtrInput                        `pulumi:"subnet"`
	VmSize                           pulumi.StringPtrInput                             `pulumi:"vmSize"`
}

func (ComputeInstancePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstancePropertiesResponse)(nil)).Elem()
}

func (i ComputeInstancePropertiesResponseArgs) ToComputeInstancePropertiesResponseOutput() ComputeInstancePropertiesResponseOutput {
	return i.ToComputeInstancePropertiesResponseOutputWithContext(context.Background())
}

func (i ComputeInstancePropertiesResponseArgs) ToComputeInstancePropertiesResponseOutputWithContext(ctx context.Context) ComputeInstancePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstancePropertiesResponseOutput)
}

func (i ComputeInstancePropertiesResponseArgs) ToComputeInstancePropertiesResponsePtrOutput() ComputeInstancePropertiesResponsePtrOutput {
	return i.ToComputeInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ComputeInstancePropertiesResponseArgs) ToComputeInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) ComputeInstancePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstancePropertiesResponseOutput).ToComputeInstancePropertiesResponsePtrOutputWithContext(ctx)
}









type ComputeInstancePropertiesResponsePtrInput interface {
	pulumi.Input

	ToComputeInstancePropertiesResponsePtrOutput() ComputeInstancePropertiesResponsePtrOutput
	ToComputeInstancePropertiesResponsePtrOutputWithContext(context.Context) ComputeInstancePropertiesResponsePtrOutput
}

type computeInstancePropertiesResponsePtrType ComputeInstancePropertiesResponseArgs

func ComputeInstancePropertiesResponsePtr(v *ComputeInstancePropertiesResponseArgs) ComputeInstancePropertiesResponsePtrInput {
	return (*computeInstancePropertiesResponsePtrType)(v)
}

func (*computeInstancePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstancePropertiesResponse)(nil)).Elem()
}

func (i *computeInstancePropertiesResponsePtrType) ToComputeInstancePropertiesResponsePtrOutput() ComputeInstancePropertiesResponsePtrOutput {
	return i.ToComputeInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *computeInstancePropertiesResponsePtrType) ToComputeInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) ComputeInstancePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstancePropertiesResponsePtrOutput)
}

type ComputeInstancePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ComputeInstancePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstancePropertiesResponse)(nil)).Elem()
}

func (o ComputeInstancePropertiesResponseOutput) ToComputeInstancePropertiesResponseOutput() ComputeInstancePropertiesResponseOutput {
	return o
}

func (o ComputeInstancePropertiesResponseOutput) ToComputeInstancePropertiesResponseOutputWithContext(ctx context.Context) ComputeInstancePropertiesResponseOutput {
	return o
}

func (o ComputeInstancePropertiesResponseOutput) ToComputeInstancePropertiesResponsePtrOutput() ComputeInstancePropertiesResponsePtrOutput {
	return o.ToComputeInstancePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ComputeInstancePropertiesResponseOutput) ToComputeInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) ComputeInstancePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeInstancePropertiesResponse) *ComputeInstancePropertiesResponse {
		return &v
	}).(ComputeInstancePropertiesResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponseOutput) ApplicationSharingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) *string { return v.ApplicationSharingPolicy }).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesResponseOutput) Applications() ComputeInstanceApplicationResponseArrayOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) []ComputeInstanceApplicationResponse { return v.Applications }).(ComputeInstanceApplicationResponseArrayOutput)
}

func (o ComputeInstancePropertiesResponseOutput) ComputeInstanceAuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) *string { return v.ComputeInstanceAuthorizationType }).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesResponseOutput) ConnectivityEndpoints() ComputeInstanceConnectivityEndpointsResponseOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) ComputeInstanceConnectivityEndpointsResponse {
		return v.ConnectivityEndpoints
	}).(ComputeInstanceConnectivityEndpointsResponseOutput)
}

func (o ComputeInstancePropertiesResponseOutput) CreatedBy() ComputeInstanceCreatedByResponseOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) ComputeInstanceCreatedByResponse { return v.CreatedBy }).(ComputeInstanceCreatedByResponseOutput)
}

func (o ComputeInstancePropertiesResponseOutput) Errors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) []ErrorResponseResponse { return v.Errors }).(ErrorResponseResponseArrayOutput)
}

func (o ComputeInstancePropertiesResponseOutput) LastOperation() ComputeInstanceLastOperationResponseOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) ComputeInstanceLastOperationResponse { return v.LastOperation }).(ComputeInstanceLastOperationResponseOutput)
}

func (o ComputeInstancePropertiesResponseOutput) PersonalComputeInstanceSettings() PersonalComputeInstanceSettingsResponsePtrOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) *PersonalComputeInstanceSettingsResponse {
		return v.PersonalComputeInstanceSettings
	}).(PersonalComputeInstanceSettingsResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponseOutput) SetupScripts() SetupScriptsResponsePtrOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) *SetupScriptsResponse { return v.SetupScripts }).(SetupScriptsResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponseOutput) SshSettings() ComputeInstanceSshSettingsResponsePtrOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) *ComputeInstanceSshSettingsResponse { return v.SshSettings }).(ComputeInstanceSshSettingsResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o ComputeInstancePropertiesResponseOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) *ResourceIdResponse { return v.Subnet }).(ResourceIdResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstancePropertiesResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type ComputeInstancePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeInstancePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstancePropertiesResponse)(nil)).Elem()
}

func (o ComputeInstancePropertiesResponsePtrOutput) ToComputeInstancePropertiesResponsePtrOutput() ComputeInstancePropertiesResponsePtrOutput {
	return o
}

func (o ComputeInstancePropertiesResponsePtrOutput) ToComputeInstancePropertiesResponsePtrOutputWithContext(ctx context.Context) ComputeInstancePropertiesResponsePtrOutput {
	return o
}

func (o ComputeInstancePropertiesResponsePtrOutput) Elem() ComputeInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) ComputeInstancePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ComputeInstancePropertiesResponse
		return ret
	}).(ComputeInstancePropertiesResponseOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) ApplicationSharingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationSharingPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) Applications() ComputeInstanceApplicationResponseArrayOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) []ComputeInstanceApplicationResponse {
		if v == nil {
			return nil
		}
		return v.Applications
	}).(ComputeInstanceApplicationResponseArrayOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) ComputeInstanceAuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputeInstanceAuthorizationType
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) ConnectivityEndpoints() ComputeInstanceConnectivityEndpointsResponsePtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *ComputeInstanceConnectivityEndpointsResponse {
		if v == nil {
			return nil
		}
		return &v.ConnectivityEndpoints
	}).(ComputeInstanceConnectivityEndpointsResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) CreatedBy() ComputeInstanceCreatedByResponsePtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *ComputeInstanceCreatedByResponse {
		if v == nil {
			return nil
		}
		return &v.CreatedBy
	}).(ComputeInstanceCreatedByResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) Errors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) []ErrorResponseResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(ErrorResponseResponseArrayOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) LastOperation() ComputeInstanceLastOperationResponsePtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *ComputeInstanceLastOperationResponse {
		if v == nil {
			return nil
		}
		return &v.LastOperation
	}).(ComputeInstanceLastOperationResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) PersonalComputeInstanceSettings() PersonalComputeInstanceSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *PersonalComputeInstanceSettingsResponse {
		if v == nil {
			return nil
		}
		return v.PersonalComputeInstanceSettings
	}).(PersonalComputeInstanceSettingsResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) SetupScripts() SetupScriptsResponsePtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *SetupScriptsResponse {
		if v == nil {
			return nil
		}
		return v.SetupScripts
	}).(SetupScriptsResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) SshSettings() ComputeInstanceSshSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *ComputeInstanceSshSettingsResponse {
		if v == nil {
			return nil
		}
		return v.SshSettings
	}).(ComputeInstanceSshSettingsResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) Subnet() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *ResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ResourceIdResponsePtrOutput)
}

func (o ComputeInstancePropertiesResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstancePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type ComputeInstanceResponse struct {
	ComputeLocation    *string                            `pulumi:"computeLocation"`
	ComputeType        string                             `pulumi:"computeType"`
	CreatedOn          string                             `pulumi:"createdOn"`
	Description        *string                            `pulumi:"description"`
	DisableLocalAuth   *bool                              `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                               `pulumi:"isAttachedCompute"`
	ModifiedOn         string                             `pulumi:"modifiedOn"`
	Properties         *ComputeInstancePropertiesResponse `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse            `pulumi:"provisioningErrors"`
	ProvisioningState  string                             `pulumi:"provisioningState"`
	ResourceId         *string                            `pulumi:"resourceId"`
}





type ComputeInstanceResponseInput interface {
	pulumi.Input

	ToComputeInstanceResponseOutput() ComputeInstanceResponseOutput
	ToComputeInstanceResponseOutputWithContext(context.Context) ComputeInstanceResponseOutput
}

type ComputeInstanceResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                     `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                        `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                        `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                     `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput                       `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                          `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                        `pulumi:"modifiedOn"`
	Properties         ComputeInstancePropertiesResponsePtrInput `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput           `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                        `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                     `pulumi:"resourceId"`
}

func (ComputeInstanceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceResponse)(nil)).Elem()
}

func (i ComputeInstanceResponseArgs) ToComputeInstanceResponseOutput() ComputeInstanceResponseOutput {
	return i.ToComputeInstanceResponseOutputWithContext(context.Background())
}

func (i ComputeInstanceResponseArgs) ToComputeInstanceResponseOutputWithContext(ctx context.Context) ComputeInstanceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceResponseOutput)
}

type ComputeInstanceResponseOutput struct{ *pulumi.OutputState }

func (ComputeInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceResponse)(nil)).Elem()
}

func (o ComputeInstanceResponseOutput) ToComputeInstanceResponseOutput() ComputeInstanceResponseOutput {
	return o
}

func (o ComputeInstanceResponseOutput) ToComputeInstanceResponseOutputWithContext(ctx context.Context) ComputeInstanceResponseOutput {
	return o
}

func (o ComputeInstanceResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o ComputeInstanceResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o ComputeInstanceResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o ComputeInstanceResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o ComputeInstanceResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o ComputeInstanceResponseOutput) Properties() ComputeInstancePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) *ComputeInstancePropertiesResponse { return v.Properties }).(ComputeInstancePropertiesResponsePtrOutput)
}

func (o ComputeInstanceResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o ComputeInstanceResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ComputeInstanceResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ComputeInstanceSshSettings struct {
	AdminPublicKey  *string `pulumi:"adminPublicKey"`
	SshPublicAccess *string `pulumi:"sshPublicAccess"`
}





type ComputeInstanceSshSettingsInput interface {
	pulumi.Input

	ToComputeInstanceSshSettingsOutput() ComputeInstanceSshSettingsOutput
	ToComputeInstanceSshSettingsOutputWithContext(context.Context) ComputeInstanceSshSettingsOutput
}

type ComputeInstanceSshSettingsArgs struct {
	AdminPublicKey  pulumi.StringPtrInput `pulumi:"adminPublicKey"`
	SshPublicAccess pulumi.StringPtrInput `pulumi:"sshPublicAccess"`
}

func (ComputeInstanceSshSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceSshSettings)(nil)).Elem()
}

func (i ComputeInstanceSshSettingsArgs) ToComputeInstanceSshSettingsOutput() ComputeInstanceSshSettingsOutput {
	return i.ToComputeInstanceSshSettingsOutputWithContext(context.Background())
}

func (i ComputeInstanceSshSettingsArgs) ToComputeInstanceSshSettingsOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceSshSettingsOutput)
}

func (i ComputeInstanceSshSettingsArgs) ToComputeInstanceSshSettingsPtrOutput() ComputeInstanceSshSettingsPtrOutput {
	return i.ToComputeInstanceSshSettingsPtrOutputWithContext(context.Background())
}

func (i ComputeInstanceSshSettingsArgs) ToComputeInstanceSshSettingsPtrOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceSshSettingsOutput).ToComputeInstanceSshSettingsPtrOutputWithContext(ctx)
}









type ComputeInstanceSshSettingsPtrInput interface {
	pulumi.Input

	ToComputeInstanceSshSettingsPtrOutput() ComputeInstanceSshSettingsPtrOutput
	ToComputeInstanceSshSettingsPtrOutputWithContext(context.Context) ComputeInstanceSshSettingsPtrOutput
}

type computeInstanceSshSettingsPtrType ComputeInstanceSshSettingsArgs

func ComputeInstanceSshSettingsPtr(v *ComputeInstanceSshSettingsArgs) ComputeInstanceSshSettingsPtrInput {
	return (*computeInstanceSshSettingsPtrType)(v)
}

func (*computeInstanceSshSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceSshSettings)(nil)).Elem()
}

func (i *computeInstanceSshSettingsPtrType) ToComputeInstanceSshSettingsPtrOutput() ComputeInstanceSshSettingsPtrOutput {
	return i.ToComputeInstanceSshSettingsPtrOutputWithContext(context.Background())
}

func (i *computeInstanceSshSettingsPtrType) ToComputeInstanceSshSettingsPtrOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceSshSettingsPtrOutput)
}

type ComputeInstanceSshSettingsOutput struct{ *pulumi.OutputState }

func (ComputeInstanceSshSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceSshSettings)(nil)).Elem()
}

func (o ComputeInstanceSshSettingsOutput) ToComputeInstanceSshSettingsOutput() ComputeInstanceSshSettingsOutput {
	return o
}

func (o ComputeInstanceSshSettingsOutput) ToComputeInstanceSshSettingsOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsOutput {
	return o
}

func (o ComputeInstanceSshSettingsOutput) ToComputeInstanceSshSettingsPtrOutput() ComputeInstanceSshSettingsPtrOutput {
	return o.ToComputeInstanceSshSettingsPtrOutputWithContext(context.Background())
}

func (o ComputeInstanceSshSettingsOutput) ToComputeInstanceSshSettingsPtrOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeInstanceSshSettings) *ComputeInstanceSshSettings {
		return &v
	}).(ComputeInstanceSshSettingsPtrOutput)
}

func (o ComputeInstanceSshSettingsOutput) AdminPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceSshSettings) *string { return v.AdminPublicKey }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceSshSettingsOutput) SshPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceSshSettings) *string { return v.SshPublicAccess }).(pulumi.StringPtrOutput)
}

type ComputeInstanceSshSettingsPtrOutput struct{ *pulumi.OutputState }

func (ComputeInstanceSshSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceSshSettings)(nil)).Elem()
}

func (o ComputeInstanceSshSettingsPtrOutput) ToComputeInstanceSshSettingsPtrOutput() ComputeInstanceSshSettingsPtrOutput {
	return o
}

func (o ComputeInstanceSshSettingsPtrOutput) ToComputeInstanceSshSettingsPtrOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsPtrOutput {
	return o
}

func (o ComputeInstanceSshSettingsPtrOutput) Elem() ComputeInstanceSshSettingsOutput {
	return o.ApplyT(func(v *ComputeInstanceSshSettings) ComputeInstanceSshSettings {
		if v != nil {
			return *v
		}
		var ret ComputeInstanceSshSettings
		return ret
	}).(ComputeInstanceSshSettingsOutput)
}

func (o ComputeInstanceSshSettingsPtrOutput) AdminPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceSshSettings) *string {
		if v == nil {
			return nil
		}
		return v.AdminPublicKey
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceSshSettingsPtrOutput) SshPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceSshSettings) *string {
		if v == nil {
			return nil
		}
		return v.SshPublicAccess
	}).(pulumi.StringPtrOutput)
}

type ComputeInstanceSshSettingsResponse struct {
	AdminPublicKey  *string `pulumi:"adminPublicKey"`
	AdminUserName   string  `pulumi:"adminUserName"`
	SshPort         int     `pulumi:"sshPort"`
	SshPublicAccess *string `pulumi:"sshPublicAccess"`
}





type ComputeInstanceSshSettingsResponseInput interface {
	pulumi.Input

	ToComputeInstanceSshSettingsResponseOutput() ComputeInstanceSshSettingsResponseOutput
	ToComputeInstanceSshSettingsResponseOutputWithContext(context.Context) ComputeInstanceSshSettingsResponseOutput
}

type ComputeInstanceSshSettingsResponseArgs struct {
	AdminPublicKey  pulumi.StringPtrInput `pulumi:"adminPublicKey"`
	AdminUserName   pulumi.StringInput    `pulumi:"adminUserName"`
	SshPort         pulumi.IntInput       `pulumi:"sshPort"`
	SshPublicAccess pulumi.StringPtrInput `pulumi:"sshPublicAccess"`
}

func (ComputeInstanceSshSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceSshSettingsResponse)(nil)).Elem()
}

func (i ComputeInstanceSshSettingsResponseArgs) ToComputeInstanceSshSettingsResponseOutput() ComputeInstanceSshSettingsResponseOutput {
	return i.ToComputeInstanceSshSettingsResponseOutputWithContext(context.Background())
}

func (i ComputeInstanceSshSettingsResponseArgs) ToComputeInstanceSshSettingsResponseOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceSshSettingsResponseOutput)
}

func (i ComputeInstanceSshSettingsResponseArgs) ToComputeInstanceSshSettingsResponsePtrOutput() ComputeInstanceSshSettingsResponsePtrOutput {
	return i.ToComputeInstanceSshSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ComputeInstanceSshSettingsResponseArgs) ToComputeInstanceSshSettingsResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceSshSettingsResponseOutput).ToComputeInstanceSshSettingsResponsePtrOutputWithContext(ctx)
}









type ComputeInstanceSshSettingsResponsePtrInput interface {
	pulumi.Input

	ToComputeInstanceSshSettingsResponsePtrOutput() ComputeInstanceSshSettingsResponsePtrOutput
	ToComputeInstanceSshSettingsResponsePtrOutputWithContext(context.Context) ComputeInstanceSshSettingsResponsePtrOutput
}

type computeInstanceSshSettingsResponsePtrType ComputeInstanceSshSettingsResponseArgs

func ComputeInstanceSshSettingsResponsePtr(v *ComputeInstanceSshSettingsResponseArgs) ComputeInstanceSshSettingsResponsePtrInput {
	return (*computeInstanceSshSettingsResponsePtrType)(v)
}

func (*computeInstanceSshSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceSshSettingsResponse)(nil)).Elem()
}

func (i *computeInstanceSshSettingsResponsePtrType) ToComputeInstanceSshSettingsResponsePtrOutput() ComputeInstanceSshSettingsResponsePtrOutput {
	return i.ToComputeInstanceSshSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *computeInstanceSshSettingsResponsePtrType) ToComputeInstanceSshSettingsResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceSshSettingsResponsePtrOutput)
}

type ComputeInstanceSshSettingsResponseOutput struct{ *pulumi.OutputState }

func (ComputeInstanceSshSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceSshSettingsResponse)(nil)).Elem()
}

func (o ComputeInstanceSshSettingsResponseOutput) ToComputeInstanceSshSettingsResponseOutput() ComputeInstanceSshSettingsResponseOutput {
	return o
}

func (o ComputeInstanceSshSettingsResponseOutput) ToComputeInstanceSshSettingsResponseOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsResponseOutput {
	return o
}

func (o ComputeInstanceSshSettingsResponseOutput) ToComputeInstanceSshSettingsResponsePtrOutput() ComputeInstanceSshSettingsResponsePtrOutput {
	return o.ToComputeInstanceSshSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ComputeInstanceSshSettingsResponseOutput) ToComputeInstanceSshSettingsResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeInstanceSshSettingsResponse) *ComputeInstanceSshSettingsResponse {
		return &v
	}).(ComputeInstanceSshSettingsResponsePtrOutput)
}

func (o ComputeInstanceSshSettingsResponseOutput) AdminPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceSshSettingsResponse) *string { return v.AdminPublicKey }).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceSshSettingsResponseOutput) AdminUserName() pulumi.StringOutput {
	return o.ApplyT(func(v ComputeInstanceSshSettingsResponse) string { return v.AdminUserName }).(pulumi.StringOutput)
}

func (o ComputeInstanceSshSettingsResponseOutput) SshPort() pulumi.IntOutput {
	return o.ApplyT(func(v ComputeInstanceSshSettingsResponse) int { return v.SshPort }).(pulumi.IntOutput)
}

func (o ComputeInstanceSshSettingsResponseOutput) SshPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeInstanceSshSettingsResponse) *string { return v.SshPublicAccess }).(pulumi.StringPtrOutput)
}

type ComputeInstanceSshSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ComputeInstanceSshSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceSshSettingsResponse)(nil)).Elem()
}

func (o ComputeInstanceSshSettingsResponsePtrOutput) ToComputeInstanceSshSettingsResponsePtrOutput() ComputeInstanceSshSettingsResponsePtrOutput {
	return o
}

func (o ComputeInstanceSshSettingsResponsePtrOutput) ToComputeInstanceSshSettingsResponsePtrOutputWithContext(ctx context.Context) ComputeInstanceSshSettingsResponsePtrOutput {
	return o
}

func (o ComputeInstanceSshSettingsResponsePtrOutput) Elem() ComputeInstanceSshSettingsResponseOutput {
	return o.ApplyT(func(v *ComputeInstanceSshSettingsResponse) ComputeInstanceSshSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ComputeInstanceSshSettingsResponse
		return ret
	}).(ComputeInstanceSshSettingsResponseOutput)
}

func (o ComputeInstanceSshSettingsResponsePtrOutput) AdminPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceSshSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPublicKey
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceSshSettingsResponsePtrOutput) AdminUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceSshSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUserName
	}).(pulumi.StringPtrOutput)
}

func (o ComputeInstanceSshSettingsResponsePtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceSshSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.SshPort
	}).(pulumi.IntPtrOutput)
}

func (o ComputeInstanceSshSettingsResponsePtrOutput) SshPublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceSshSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshPublicAccess
	}).(pulumi.StringPtrOutput)
}

type CosmosDbSettings struct {
	CollectionsThroughput *int `pulumi:"collectionsThroughput"`
}





type CosmosDbSettingsInput interface {
	pulumi.Input

	ToCosmosDbSettingsOutput() CosmosDbSettingsOutput
	ToCosmosDbSettingsOutputWithContext(context.Context) CosmosDbSettingsOutput
}

type CosmosDbSettingsArgs struct {
	CollectionsThroughput pulumi.IntPtrInput `pulumi:"collectionsThroughput"`
}

func (CosmosDbSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbSettings)(nil)).Elem()
}

func (i CosmosDbSettingsArgs) ToCosmosDbSettingsOutput() CosmosDbSettingsOutput {
	return i.ToCosmosDbSettingsOutputWithContext(context.Background())
}

func (i CosmosDbSettingsArgs) ToCosmosDbSettingsOutputWithContext(ctx context.Context) CosmosDbSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsOutput)
}

func (i CosmosDbSettingsArgs) ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput {
	return i.ToCosmosDbSettingsPtrOutputWithContext(context.Background())
}

func (i CosmosDbSettingsArgs) ToCosmosDbSettingsPtrOutputWithContext(ctx context.Context) CosmosDbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsOutput).ToCosmosDbSettingsPtrOutputWithContext(ctx)
}









type CosmosDbSettingsPtrInput interface {
	pulumi.Input

	ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput
	ToCosmosDbSettingsPtrOutputWithContext(context.Context) CosmosDbSettingsPtrOutput
}

type cosmosDbSettingsPtrType CosmosDbSettingsArgs

func CosmosDbSettingsPtr(v *CosmosDbSettingsArgs) CosmosDbSettingsPtrInput {
	return (*cosmosDbSettingsPtrType)(v)
}

func (*cosmosDbSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbSettings)(nil)).Elem()
}

func (i *cosmosDbSettingsPtrType) ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput {
	return i.ToCosmosDbSettingsPtrOutputWithContext(context.Background())
}

func (i *cosmosDbSettingsPtrType) ToCosmosDbSettingsPtrOutputWithContext(ctx context.Context) CosmosDbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsPtrOutput)
}

type CosmosDbSettingsOutput struct{ *pulumi.OutputState }

func (CosmosDbSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbSettings)(nil)).Elem()
}

func (o CosmosDbSettingsOutput) ToCosmosDbSettingsOutput() CosmosDbSettingsOutput {
	return o
}

func (o CosmosDbSettingsOutput) ToCosmosDbSettingsOutputWithContext(ctx context.Context) CosmosDbSettingsOutput {
	return o
}

func (o CosmosDbSettingsOutput) ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput {
	return o.ToCosmosDbSettingsPtrOutputWithContext(context.Background())
}

func (o CosmosDbSettingsOutput) ToCosmosDbSettingsPtrOutputWithContext(ctx context.Context) CosmosDbSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CosmosDbSettings) *CosmosDbSettings {
		return &v
	}).(CosmosDbSettingsPtrOutput)
}

func (o CosmosDbSettingsOutput) CollectionsThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CosmosDbSettings) *int { return v.CollectionsThroughput }).(pulumi.IntPtrOutput)
}

type CosmosDbSettingsPtrOutput struct{ *pulumi.OutputState }

func (CosmosDbSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbSettings)(nil)).Elem()
}

func (o CosmosDbSettingsPtrOutput) ToCosmosDbSettingsPtrOutput() CosmosDbSettingsPtrOutput {
	return o
}

func (o CosmosDbSettingsPtrOutput) ToCosmosDbSettingsPtrOutputWithContext(ctx context.Context) CosmosDbSettingsPtrOutput {
	return o
}

func (o CosmosDbSettingsPtrOutput) Elem() CosmosDbSettingsOutput {
	return o.ApplyT(func(v *CosmosDbSettings) CosmosDbSettings {
		if v != nil {
			return *v
		}
		var ret CosmosDbSettings
		return ret
	}).(CosmosDbSettingsOutput)
}

func (o CosmosDbSettingsPtrOutput) CollectionsThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CosmosDbSettings) *int {
		if v == nil {
			return nil
		}
		return v.CollectionsThroughput
	}).(pulumi.IntPtrOutput)
}

type CosmosDbSettingsResponse struct {
	CollectionsThroughput *int `pulumi:"collectionsThroughput"`
}





type CosmosDbSettingsResponseInput interface {
	pulumi.Input

	ToCosmosDbSettingsResponseOutput() CosmosDbSettingsResponseOutput
	ToCosmosDbSettingsResponseOutputWithContext(context.Context) CosmosDbSettingsResponseOutput
}

type CosmosDbSettingsResponseArgs struct {
	CollectionsThroughput pulumi.IntPtrInput `pulumi:"collectionsThroughput"`
}

func (CosmosDbSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbSettingsResponse)(nil)).Elem()
}

func (i CosmosDbSettingsResponseArgs) ToCosmosDbSettingsResponseOutput() CosmosDbSettingsResponseOutput {
	return i.ToCosmosDbSettingsResponseOutputWithContext(context.Background())
}

func (i CosmosDbSettingsResponseArgs) ToCosmosDbSettingsResponseOutputWithContext(ctx context.Context) CosmosDbSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsResponseOutput)
}

func (i CosmosDbSettingsResponseArgs) ToCosmosDbSettingsResponsePtrOutput() CosmosDbSettingsResponsePtrOutput {
	return i.ToCosmosDbSettingsResponsePtrOutputWithContext(context.Background())
}

func (i CosmosDbSettingsResponseArgs) ToCosmosDbSettingsResponsePtrOutputWithContext(ctx context.Context) CosmosDbSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsResponseOutput).ToCosmosDbSettingsResponsePtrOutputWithContext(ctx)
}









type CosmosDbSettingsResponsePtrInput interface {
	pulumi.Input

	ToCosmosDbSettingsResponsePtrOutput() CosmosDbSettingsResponsePtrOutput
	ToCosmosDbSettingsResponsePtrOutputWithContext(context.Context) CosmosDbSettingsResponsePtrOutput
}

type cosmosDbSettingsResponsePtrType CosmosDbSettingsResponseArgs

func CosmosDbSettingsResponsePtr(v *CosmosDbSettingsResponseArgs) CosmosDbSettingsResponsePtrInput {
	return (*cosmosDbSettingsResponsePtrType)(v)
}

func (*cosmosDbSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbSettingsResponse)(nil)).Elem()
}

func (i *cosmosDbSettingsResponsePtrType) ToCosmosDbSettingsResponsePtrOutput() CosmosDbSettingsResponsePtrOutput {
	return i.ToCosmosDbSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *cosmosDbSettingsResponsePtrType) ToCosmosDbSettingsResponsePtrOutputWithContext(ctx context.Context) CosmosDbSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosmosDbSettingsResponsePtrOutput)
}

type CosmosDbSettingsResponseOutput struct{ *pulumi.OutputState }

func (CosmosDbSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbSettingsResponse)(nil)).Elem()
}

func (o CosmosDbSettingsResponseOutput) ToCosmosDbSettingsResponseOutput() CosmosDbSettingsResponseOutput {
	return o
}

func (o CosmosDbSettingsResponseOutput) ToCosmosDbSettingsResponseOutputWithContext(ctx context.Context) CosmosDbSettingsResponseOutput {
	return o
}

func (o CosmosDbSettingsResponseOutput) ToCosmosDbSettingsResponsePtrOutput() CosmosDbSettingsResponsePtrOutput {
	return o.ToCosmosDbSettingsResponsePtrOutputWithContext(context.Background())
}

func (o CosmosDbSettingsResponseOutput) ToCosmosDbSettingsResponsePtrOutputWithContext(ctx context.Context) CosmosDbSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CosmosDbSettingsResponse) *CosmosDbSettingsResponse {
		return &v
	}).(CosmosDbSettingsResponsePtrOutput)
}

func (o CosmosDbSettingsResponseOutput) CollectionsThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CosmosDbSettingsResponse) *int { return v.CollectionsThroughput }).(pulumi.IntPtrOutput)
}

type CosmosDbSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CosmosDbSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbSettingsResponse)(nil)).Elem()
}

func (o CosmosDbSettingsResponsePtrOutput) ToCosmosDbSettingsResponsePtrOutput() CosmosDbSettingsResponsePtrOutput {
	return o
}

func (o CosmosDbSettingsResponsePtrOutput) ToCosmosDbSettingsResponsePtrOutputWithContext(ctx context.Context) CosmosDbSettingsResponsePtrOutput {
	return o
}

func (o CosmosDbSettingsResponsePtrOutput) Elem() CosmosDbSettingsResponseOutput {
	return o.ApplyT(func(v *CosmosDbSettingsResponse) CosmosDbSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CosmosDbSettingsResponse
		return ret
	}).(CosmosDbSettingsResponseOutput)
}

func (o CosmosDbSettingsResponsePtrOutput) CollectionsThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CosmosDbSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.CollectionsThroughput
	}).(pulumi.IntPtrOutput)
}

type DataFactory struct {
	ComputeLocation  *string `pulumi:"computeLocation"`
	ComputeType      string  `pulumi:"computeType"`
	Description      *string `pulumi:"description"`
	DisableLocalAuth *bool   `pulumi:"disableLocalAuth"`
	ResourceId       *string `pulumi:"resourceId"`
}





type DataFactoryInput interface {
	pulumi.Input

	ToDataFactoryOutput() DataFactoryOutput
	ToDataFactoryOutputWithContext(context.Context) DataFactoryOutput
}

type DataFactoryArgs struct {
	ComputeLocation  pulumi.StringPtrInput `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput    `pulumi:"computeType"`
	Description      pulumi.StringPtrInput `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput   `pulumi:"disableLocalAuth"`
	ResourceId       pulumi.StringPtrInput `pulumi:"resourceId"`
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

func (o DataFactoryOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataFactory) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o DataFactoryOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactory) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataFactoryResponse struct {
	ComputeLocation    *string                 `pulumi:"computeLocation"`
	ComputeType        string                  `pulumi:"computeType"`
	CreatedOn          string                  `pulumi:"createdOn"`
	Description        *string                 `pulumi:"description"`
	DisableLocalAuth   *bool                   `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                    `pulumi:"isAttachedCompute"`
	ModifiedOn         string                  `pulumi:"modifiedOn"`
	ProvisioningErrors []ErrorResponseResponse `pulumi:"provisioningErrors"`
	ProvisioningState  string                  `pulumi:"provisioningState"`
	ResourceId         *string                 `pulumi:"resourceId"`
}





type DataFactoryResponseInput interface {
	pulumi.Input

	ToDataFactoryResponseOutput() DataFactoryResponseOutput
	ToDataFactoryResponseOutputWithContext(context.Context) DataFactoryResponseOutput
}

type DataFactoryResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput           `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput              `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput              `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput           `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput             `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput              `pulumi:"modifiedOn"`
	ProvisioningErrors ErrorResponseResponseArrayInput `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput              `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput           `pulumi:"resourceId"`
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

func (o DataFactoryResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataFactoryResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o DataFactoryResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v DataFactoryResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o DataFactoryResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v DataFactoryResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o DataFactoryResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DataFactoryResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataFactoryResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFactoryResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DataLakeAnalytics struct {
	ComputeLocation  *string                      `pulumi:"computeLocation"`
	ComputeType      string                       `pulumi:"computeType"`
	Description      *string                      `pulumi:"description"`
	DisableLocalAuth *bool                        `pulumi:"disableLocalAuth"`
	Properties       *DataLakeAnalyticsProperties `pulumi:"properties"`
	ResourceId       *string                      `pulumi:"resourceId"`
}





type DataLakeAnalyticsInput interface {
	pulumi.Input

	ToDataLakeAnalyticsOutput() DataLakeAnalyticsOutput
	ToDataLakeAnalyticsOutputWithContext(context.Context) DataLakeAnalyticsOutput
}

type DataLakeAnalyticsArgs struct {
	ComputeLocation  pulumi.StringPtrInput               `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput                  `pulumi:"computeType"`
	Description      pulumi.StringPtrInput               `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput                 `pulumi:"disableLocalAuth"`
	Properties       DataLakeAnalyticsPropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput               `pulumi:"resourceId"`
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

func (o DataLakeAnalyticsOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataLakeAnalytics) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
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
	ComputeLocation    *string                              `pulumi:"computeLocation"`
	ComputeType        string                               `pulumi:"computeType"`
	CreatedOn          string                               `pulumi:"createdOn"`
	Description        *string                              `pulumi:"description"`
	DisableLocalAuth   *bool                                `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                                 `pulumi:"isAttachedCompute"`
	ModifiedOn         string                               `pulumi:"modifiedOn"`
	Properties         *DataLakeAnalyticsResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse              `pulumi:"provisioningErrors"`
	ProvisioningState  string                               `pulumi:"provisioningState"`
	ResourceId         *string                              `pulumi:"resourceId"`
}





type DataLakeAnalyticsResponseInput interface {
	pulumi.Input

	ToDataLakeAnalyticsResponseOutput() DataLakeAnalyticsResponseOutput
	ToDataLakeAnalyticsResponseOutputWithContext(context.Context) DataLakeAnalyticsResponseOutput
}

type DataLakeAnalyticsResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                       `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                          `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                          `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                       `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                            `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                          `pulumi:"modifiedOn"`
	Properties         DataLakeAnalyticsResponsePropertiesPtrInput `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput             `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                          `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                       `pulumi:"resourceId"`
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

func (o DataLakeAnalyticsResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
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

func (o DataLakeAnalyticsResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v DataLakeAnalyticsResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
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
	ComputeLocation  *string               `pulumi:"computeLocation"`
	ComputeType      string                `pulumi:"computeType"`
	Description      *string               `pulumi:"description"`
	DisableLocalAuth *bool                 `pulumi:"disableLocalAuth"`
	Properties       *DatabricksProperties `pulumi:"properties"`
	ResourceId       *string               `pulumi:"resourceId"`
}





type DatabricksInput interface {
	pulumi.Input

	ToDatabricksOutput() DatabricksOutput
	ToDatabricksOutputWithContext(context.Context) DatabricksOutput
}

type DatabricksArgs struct {
	ComputeLocation  pulumi.StringPtrInput        `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput           `pulumi:"computeType"`
	Description      pulumi.StringPtrInput        `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput          `pulumi:"disableLocalAuth"`
	Properties       DatabricksPropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput        `pulumi:"resourceId"`
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

func (o DatabricksOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Databricks) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o DatabricksOutput) Properties() DatabricksPropertiesPtrOutput {
	return o.ApplyT(func(v Databricks) *DatabricksProperties { return v.Properties }).(DatabricksPropertiesPtrOutput)
}

func (o DatabricksOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Databricks) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type DatabricksProperties struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
	WorkspaceUrl          *string `pulumi:"workspaceUrl"`
}





type DatabricksPropertiesInput interface {
	pulumi.Input

	ToDatabricksPropertiesOutput() DatabricksPropertiesOutput
	ToDatabricksPropertiesOutputWithContext(context.Context) DatabricksPropertiesOutput
}

type DatabricksPropertiesArgs struct {
	DatabricksAccessToken pulumi.StringPtrInput `pulumi:"databricksAccessToken"`
	WorkspaceUrl          pulumi.StringPtrInput `pulumi:"workspaceUrl"`
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

func (o DatabricksPropertiesOutput) WorkspaceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksProperties) *string { return v.WorkspaceUrl }).(pulumi.StringPtrOutput)
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

func (o DatabricksPropertiesPtrOutput) WorkspaceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceUrl
	}).(pulumi.StringPtrOutput)
}

type DatabricksPropertiesResponse struct {
	DatabricksAccessToken *string `pulumi:"databricksAccessToken"`
	WorkspaceUrl          *string `pulumi:"workspaceUrl"`
}





type DatabricksPropertiesResponseInput interface {
	pulumi.Input

	ToDatabricksPropertiesResponseOutput() DatabricksPropertiesResponseOutput
	ToDatabricksPropertiesResponseOutputWithContext(context.Context) DatabricksPropertiesResponseOutput
}

type DatabricksPropertiesResponseArgs struct {
	DatabricksAccessToken pulumi.StringPtrInput `pulumi:"databricksAccessToken"`
	WorkspaceUrl          pulumi.StringPtrInput `pulumi:"workspaceUrl"`
}

func (DatabricksPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksPropertiesResponse)(nil)).Elem()
}

func (i DatabricksPropertiesResponseArgs) ToDatabricksPropertiesResponseOutput() DatabricksPropertiesResponseOutput {
	return i.ToDatabricksPropertiesResponseOutputWithContext(context.Background())
}

func (i DatabricksPropertiesResponseArgs) ToDatabricksPropertiesResponseOutputWithContext(ctx context.Context) DatabricksPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesResponseOutput)
}

func (i DatabricksPropertiesResponseArgs) ToDatabricksPropertiesResponsePtrOutput() DatabricksPropertiesResponsePtrOutput {
	return i.ToDatabricksPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DatabricksPropertiesResponseArgs) ToDatabricksPropertiesResponsePtrOutputWithContext(ctx context.Context) DatabricksPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesResponseOutput).ToDatabricksPropertiesResponsePtrOutputWithContext(ctx)
}









type DatabricksPropertiesResponsePtrInput interface {
	pulumi.Input

	ToDatabricksPropertiesResponsePtrOutput() DatabricksPropertiesResponsePtrOutput
	ToDatabricksPropertiesResponsePtrOutputWithContext(context.Context) DatabricksPropertiesResponsePtrOutput
}

type databricksPropertiesResponsePtrType DatabricksPropertiesResponseArgs

func DatabricksPropertiesResponsePtr(v *DatabricksPropertiesResponseArgs) DatabricksPropertiesResponsePtrInput {
	return (*databricksPropertiesResponsePtrType)(v)
}

func (*databricksPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksPropertiesResponse)(nil)).Elem()
}

func (i *databricksPropertiesResponsePtrType) ToDatabricksPropertiesResponsePtrOutput() DatabricksPropertiesResponsePtrOutput {
	return i.ToDatabricksPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *databricksPropertiesResponsePtrType) ToDatabricksPropertiesResponsePtrOutputWithContext(ctx context.Context) DatabricksPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksPropertiesResponsePtrOutput)
}

type DatabricksPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DatabricksPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabricksPropertiesResponse)(nil)).Elem()
}

func (o DatabricksPropertiesResponseOutput) ToDatabricksPropertiesResponseOutput() DatabricksPropertiesResponseOutput {
	return o
}

func (o DatabricksPropertiesResponseOutput) ToDatabricksPropertiesResponseOutputWithContext(ctx context.Context) DatabricksPropertiesResponseOutput {
	return o
}

func (o DatabricksPropertiesResponseOutput) ToDatabricksPropertiesResponsePtrOutput() DatabricksPropertiesResponsePtrOutput {
	return o.ToDatabricksPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DatabricksPropertiesResponseOutput) ToDatabricksPropertiesResponsePtrOutputWithContext(ctx context.Context) DatabricksPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabricksPropertiesResponse) *DatabricksPropertiesResponse {
		return &v
	}).(DatabricksPropertiesResponsePtrOutput)
}

func (o DatabricksPropertiesResponseOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksPropertiesResponse) *string { return v.DatabricksAccessToken }).(pulumi.StringPtrOutput)
}

func (o DatabricksPropertiesResponseOutput) WorkspaceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksPropertiesResponse) *string { return v.WorkspaceUrl }).(pulumi.StringPtrOutput)
}

type DatabricksPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DatabricksPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksPropertiesResponse)(nil)).Elem()
}

func (o DatabricksPropertiesResponsePtrOutput) ToDatabricksPropertiesResponsePtrOutput() DatabricksPropertiesResponsePtrOutput {
	return o
}

func (o DatabricksPropertiesResponsePtrOutput) ToDatabricksPropertiesResponsePtrOutputWithContext(ctx context.Context) DatabricksPropertiesResponsePtrOutput {
	return o
}

func (o DatabricksPropertiesResponsePtrOutput) Elem() DatabricksPropertiesResponseOutput {
	return o.ApplyT(func(v *DatabricksPropertiesResponse) DatabricksPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DatabricksPropertiesResponse
		return ret
	}).(DatabricksPropertiesResponseOutput)
}

func (o DatabricksPropertiesResponsePtrOutput) DatabricksAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DatabricksAccessToken
	}).(pulumi.StringPtrOutput)
}

func (o DatabricksPropertiesResponsePtrOutput) WorkspaceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceUrl
	}).(pulumi.StringPtrOutput)
}

type DatabricksResponse struct {
	ComputeLocation    *string                       `pulumi:"computeLocation"`
	ComputeType        string                        `pulumi:"computeType"`
	CreatedOn          string                        `pulumi:"createdOn"`
	Description        *string                       `pulumi:"description"`
	DisableLocalAuth   *bool                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                        `pulumi:"modifiedOn"`
	Properties         *DatabricksPropertiesResponse `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse       `pulumi:"provisioningErrors"`
	ProvisioningState  string                        `pulumi:"provisioningState"`
	ResourceId         *string                       `pulumi:"resourceId"`
}





type DatabricksResponseInput interface {
	pulumi.Input

	ToDatabricksResponseOutput() DatabricksResponseOutput
	ToDatabricksResponseOutputWithContext(context.Context) DatabricksResponseOutput
}

type DatabricksResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                   `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                   `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput                  `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                     `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                   `pulumi:"modifiedOn"`
	Properties         DatabricksPropertiesResponsePtrInput `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput      `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                   `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                `pulumi:"resourceId"`
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

func (o DatabricksResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o DatabricksResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v DatabricksResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o DatabricksResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) Properties() DatabricksPropertiesResponsePtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *DatabricksPropertiesResponse { return v.Properties }).(DatabricksPropertiesResponsePtrOutput)
}

func (o DatabricksResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v DatabricksResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o DatabricksResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DatabricksResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatabricksResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabricksResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type EncryptionProperty struct {
	Identity           *IdentityForCmk    `pulumi:"identity"`
	KeyVaultProperties KeyVaultProperties `pulumi:"keyVaultProperties"`
	Status             string             `pulumi:"status"`
}





type EncryptionPropertyInput interface {
	pulumi.Input

	ToEncryptionPropertyOutput() EncryptionPropertyOutput
	ToEncryptionPropertyOutputWithContext(context.Context) EncryptionPropertyOutput
}

type EncryptionPropertyArgs struct {
	Identity           IdentityForCmkPtrInput  `pulumi:"identity"`
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

func (o EncryptionPropertyOutput) Identity() IdentityForCmkPtrOutput {
	return o.ApplyT(func(v EncryptionProperty) *IdentityForCmk { return v.Identity }).(IdentityForCmkPtrOutput)
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

func (o EncryptionPropertyPtrOutput) Identity() IdentityForCmkPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *IdentityForCmk {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(IdentityForCmkPtrOutput)
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
	Identity           *IdentityForCmkResponse    `pulumi:"identity"`
	KeyVaultProperties KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             string                     `pulumi:"status"`
}





type EncryptionPropertyResponseInput interface {
	pulumi.Input

	ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput
	ToEncryptionPropertyResponseOutputWithContext(context.Context) EncryptionPropertyResponseOutput
}

type EncryptionPropertyResponseArgs struct {
	Identity           IdentityForCmkResponsePtrInput  `pulumi:"identity"`
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

func (o EncryptionPropertyResponseOutput) Identity() IdentityForCmkResponsePtrOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) *IdentityForCmkResponse { return v.Identity }).(IdentityForCmkResponsePtrOutput)
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

func (o EncryptionPropertyResponsePtrOutput) Identity() IdentityForCmkResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *IdentityForCmkResponse {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(IdentityForCmkResponsePtrOutput)
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

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}





type ErrorAdditionalInfoResponseInput interface {
	pulumi.Input

	ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput
	ToErrorAdditionalInfoResponseOutputWithContext(context.Context) ErrorAdditionalInfoResponseOutput
}

type ErrorAdditionalInfoResponseArgs struct {
	Info pulumi.Input       `pulumi:"info"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ErrorAdditionalInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (i ErrorAdditionalInfoResponseArgs) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return i.ToErrorAdditionalInfoResponseOutputWithContext(context.Background())
}

func (i ErrorAdditionalInfoResponseArgs) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorAdditionalInfoResponseOutput)
}





type ErrorAdditionalInfoResponseArrayInput interface {
	pulumi.Input

	ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput
	ToErrorAdditionalInfoResponseArrayOutputWithContext(context.Context) ErrorAdditionalInfoResponseArrayOutput
}

type ErrorAdditionalInfoResponseArray []ErrorAdditionalInfoResponseInput

func (ErrorAdditionalInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (i ErrorAdditionalInfoResponseArray) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return i.ToErrorAdditionalInfoResponseArrayOutputWithContext(context.Background())
}

func (i ErrorAdditionalInfoResponseArray) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorAdditionalInfoResponseArrayOutput)
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

type ErrorDetailResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorDetailResponse         `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	AdditionalInfo ErrorAdditionalInfoResponseArrayInput `pulumi:"additionalInfo"`
	Code           pulumi.StringInput                    `pulumi:"code"`
	Details        ErrorDetailResponseArrayInput         `pulumi:"details"`
	Message        pulumi.StringInput                    `pulumi:"message"`
	Target         pulumi.StringInput                    `pulumi:"target"`
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

func (i ErrorDetailResponseArgs) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return i.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput).ToErrorDetailResponsePtrOutputWithContext(ctx)
}









type ErrorDetailResponsePtrInput interface {
	pulumi.Input

	ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput
	ToErrorDetailResponsePtrOutputWithContext(context.Context) ErrorDetailResponsePtrOutput
}

type errorDetailResponsePtrType ErrorDetailResponseArgs

func ErrorDetailResponsePtr(v *ErrorDetailResponseArgs) ErrorDetailResponsePtrInput {
	return (*errorDetailResponsePtrType)(v)
}

func (*errorDetailResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (i *errorDetailResponsePtrType) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return i.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (i *errorDetailResponsePtrType) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponsePtrOutput)
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

func (o ErrorDetailResponseOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o.ToErrorDetailResponsePtrOutputWithContext(context.Background())
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorDetailResponse) *ErrorDetailResponse {
		return &v
	}).(ErrorDetailResponsePtrOutput)
}

func (o ErrorDetailResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDetailResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutput() ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) ToErrorDetailResponsePtrOutputWithContext(ctx context.Context) ErrorDetailResponsePtrOutput {
	return o
}

func (o ErrorDetailResponsePtrOutput) Elem() ErrorDetailResponseOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) ErrorDetailResponse {
		if v != nil {
			return *v
		}
		var ret ErrorDetailResponse
		return ret
	}).(ErrorDetailResponseOutput)
}

func (o ErrorDetailResponsePtrOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []ErrorAdditionalInfoResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalInfo
	}).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) []ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
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
	Error *ErrorDetailResponse `pulumi:"error"`
}





type ErrorResponseResponseInput interface {
	pulumi.Input

	ToErrorResponseResponseOutput() ErrorResponseResponseOutput
	ToErrorResponseResponseOutputWithContext(context.Context) ErrorResponseResponseOutput
}

type ErrorResponseResponseArgs struct {
	Error ErrorDetailResponsePtrInput `pulumi:"error"`
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





type ErrorResponseResponseArrayInput interface {
	pulumi.Input

	ToErrorResponseResponseArrayOutput() ErrorResponseResponseArrayOutput
	ToErrorResponseResponseArrayOutputWithContext(context.Context) ErrorResponseResponseArrayOutput
}

type ErrorResponseResponseArray []ErrorResponseResponseInput

func (ErrorResponseResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorResponseResponse)(nil)).Elem()
}

func (i ErrorResponseResponseArray) ToErrorResponseResponseArrayOutput() ErrorResponseResponseArrayOutput {
	return i.ToErrorResponseResponseArrayOutputWithContext(context.Background())
}

func (i ErrorResponseResponseArray) ToErrorResponseResponseArrayOutputWithContext(ctx context.Context) ErrorResponseResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseResponseArrayOutput)
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

func (o ErrorResponseResponseOutput) Error() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v ErrorResponseResponse) *ErrorDetailResponse { return v.Error }).(ErrorDetailResponsePtrOutput)
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

type HDInsight struct {
	ComputeLocation  *string              `pulumi:"computeLocation"`
	ComputeType      string               `pulumi:"computeType"`
	Description      *string              `pulumi:"description"`
	DisableLocalAuth *bool                `pulumi:"disableLocalAuth"`
	Properties       *HDInsightProperties `pulumi:"properties"`
	ResourceId       *string              `pulumi:"resourceId"`
}





type HDInsightInput interface {
	pulumi.Input

	ToHDInsightOutput() HDInsightOutput
	ToHDInsightOutputWithContext(context.Context) HDInsightOutput
}

type HDInsightArgs struct {
	ComputeLocation  pulumi.StringPtrInput       `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput          `pulumi:"computeType"`
	Description      pulumi.StringPtrInput       `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput         `pulumi:"disableLocalAuth"`
	Properties       HDInsightPropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput       `pulumi:"resourceId"`
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

func (o HDInsightOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HDInsight) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
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

type HDInsightPropertiesResponse struct {
	Address              *string                               `pulumi:"address"`
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	SshPort              *int                                  `pulumi:"sshPort"`
}





type HDInsightPropertiesResponseInput interface {
	pulumi.Input

	ToHDInsightPropertiesResponseOutput() HDInsightPropertiesResponseOutput
	ToHDInsightPropertiesResponseOutputWithContext(context.Context) HDInsightPropertiesResponseOutput
}

type HDInsightPropertiesResponseArgs struct {
	Address              pulumi.StringPtrInput                        `pulumi:"address"`
	AdministratorAccount VirtualMachineSshCredentialsResponsePtrInput `pulumi:"administratorAccount"`
	SshPort              pulumi.IntPtrInput                           `pulumi:"sshPort"`
}

func (HDInsightPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightPropertiesResponse)(nil)).Elem()
}

func (i HDInsightPropertiesResponseArgs) ToHDInsightPropertiesResponseOutput() HDInsightPropertiesResponseOutput {
	return i.ToHDInsightPropertiesResponseOutputWithContext(context.Background())
}

func (i HDInsightPropertiesResponseArgs) ToHDInsightPropertiesResponseOutputWithContext(ctx context.Context) HDInsightPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesResponseOutput)
}

func (i HDInsightPropertiesResponseArgs) ToHDInsightPropertiesResponsePtrOutput() HDInsightPropertiesResponsePtrOutput {
	return i.ToHDInsightPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i HDInsightPropertiesResponseArgs) ToHDInsightPropertiesResponsePtrOutputWithContext(ctx context.Context) HDInsightPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesResponseOutput).ToHDInsightPropertiesResponsePtrOutputWithContext(ctx)
}









type HDInsightPropertiesResponsePtrInput interface {
	pulumi.Input

	ToHDInsightPropertiesResponsePtrOutput() HDInsightPropertiesResponsePtrOutput
	ToHDInsightPropertiesResponsePtrOutputWithContext(context.Context) HDInsightPropertiesResponsePtrOutput
}

type hdinsightPropertiesResponsePtrType HDInsightPropertiesResponseArgs

func HDInsightPropertiesResponsePtr(v *HDInsightPropertiesResponseArgs) HDInsightPropertiesResponsePtrInput {
	return (*hdinsightPropertiesResponsePtrType)(v)
}

func (*hdinsightPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightPropertiesResponse)(nil)).Elem()
}

func (i *hdinsightPropertiesResponsePtrType) ToHDInsightPropertiesResponsePtrOutput() HDInsightPropertiesResponsePtrOutput {
	return i.ToHDInsightPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *hdinsightPropertiesResponsePtrType) ToHDInsightPropertiesResponsePtrOutputWithContext(ctx context.Context) HDInsightPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HDInsightPropertiesResponsePtrOutput)
}

type HDInsightPropertiesResponseOutput struct{ *pulumi.OutputState }

func (HDInsightPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightPropertiesResponse)(nil)).Elem()
}

func (o HDInsightPropertiesResponseOutput) ToHDInsightPropertiesResponseOutput() HDInsightPropertiesResponseOutput {
	return o
}

func (o HDInsightPropertiesResponseOutput) ToHDInsightPropertiesResponseOutputWithContext(ctx context.Context) HDInsightPropertiesResponseOutput {
	return o
}

func (o HDInsightPropertiesResponseOutput) ToHDInsightPropertiesResponsePtrOutput() HDInsightPropertiesResponsePtrOutput {
	return o.ToHDInsightPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o HDInsightPropertiesResponseOutput) ToHDInsightPropertiesResponsePtrOutputWithContext(ctx context.Context) HDInsightPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HDInsightPropertiesResponse) *HDInsightPropertiesResponse {
		return &v
	}).(HDInsightPropertiesResponsePtrOutput)
}

func (o HDInsightPropertiesResponseOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightPropertiesResponse) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o HDInsightPropertiesResponseOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v HDInsightPropertiesResponse) *VirtualMachineSshCredentialsResponse {
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o HDInsightPropertiesResponseOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HDInsightPropertiesResponse) *int { return v.SshPort }).(pulumi.IntPtrOutput)
}

type HDInsightPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HDInsightPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightPropertiesResponse)(nil)).Elem()
}

func (o HDInsightPropertiesResponsePtrOutput) ToHDInsightPropertiesResponsePtrOutput() HDInsightPropertiesResponsePtrOutput {
	return o
}

func (o HDInsightPropertiesResponsePtrOutput) ToHDInsightPropertiesResponsePtrOutputWithContext(ctx context.Context) HDInsightPropertiesResponsePtrOutput {
	return o
}

func (o HDInsightPropertiesResponsePtrOutput) Elem() HDInsightPropertiesResponseOutput {
	return o.ApplyT(func(v *HDInsightPropertiesResponse) HDInsightPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HDInsightPropertiesResponse
		return ret
	}).(HDInsightPropertiesResponseOutput)
}

func (o HDInsightPropertiesResponsePtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HDInsightPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

func (o HDInsightPropertiesResponsePtrOutput) AdministratorAccount() VirtualMachineSshCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *HDInsightPropertiesResponse) *VirtualMachineSshCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.AdministratorAccount
	}).(VirtualMachineSshCredentialsResponsePtrOutput)
}

func (o HDInsightPropertiesResponsePtrOutput) SshPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HDInsightPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.SshPort
	}).(pulumi.IntPtrOutput)
}

type HDInsightResponse struct {
	ComputeLocation    *string                      `pulumi:"computeLocation"`
	ComputeType        string                       `pulumi:"computeType"`
	CreatedOn          string                       `pulumi:"createdOn"`
	Description        *string                      `pulumi:"description"`
	DisableLocalAuth   *bool                        `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                         `pulumi:"isAttachedCompute"`
	ModifiedOn         string                       `pulumi:"modifiedOn"`
	Properties         *HDInsightPropertiesResponse `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse      `pulumi:"provisioningErrors"`
	ProvisioningState  string                       `pulumi:"provisioningState"`
	ResourceId         *string                      `pulumi:"resourceId"`
}





type HDInsightResponseInput interface {
	pulumi.Input

	ToHDInsightResponseOutput() HDInsightResponseOutput
	ToHDInsightResponseOutputWithContext(context.Context) HDInsightResponseOutput
}

type HDInsightResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput               `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                  `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                  `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput               `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput                 `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                    `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                  `pulumi:"modifiedOn"`
	Properties         HDInsightPropertiesResponsePtrInput `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput     `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                  `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput               `pulumi:"resourceId"`
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

func (o HDInsightResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o HDInsightResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v HDInsightResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o HDInsightResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) Properties() HDInsightPropertiesResponsePtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *HDInsightPropertiesResponse { return v.Properties }).(HDInsightPropertiesResponsePtrOutput)
}

func (o HDInsightResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v HDInsightResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o HDInsightResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v HDInsightResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HDInsightResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HDInsightResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
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

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
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

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityForCmk struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type IdentityForCmkInput interface {
	pulumi.Input

	ToIdentityForCmkOutput() IdentityForCmkOutput
	ToIdentityForCmkOutputWithContext(context.Context) IdentityForCmkOutput
}

type IdentityForCmkArgs struct {
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (IdentityForCmkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityForCmk)(nil)).Elem()
}

func (i IdentityForCmkArgs) ToIdentityForCmkOutput() IdentityForCmkOutput {
	return i.ToIdentityForCmkOutputWithContext(context.Background())
}

func (i IdentityForCmkArgs) ToIdentityForCmkOutputWithContext(ctx context.Context) IdentityForCmkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkOutput)
}

func (i IdentityForCmkArgs) ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput {
	return i.ToIdentityForCmkPtrOutputWithContext(context.Background())
}

func (i IdentityForCmkArgs) ToIdentityForCmkPtrOutputWithContext(ctx context.Context) IdentityForCmkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkOutput).ToIdentityForCmkPtrOutputWithContext(ctx)
}









type IdentityForCmkPtrInput interface {
	pulumi.Input

	ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput
	ToIdentityForCmkPtrOutputWithContext(context.Context) IdentityForCmkPtrOutput
}

type identityForCmkPtrType IdentityForCmkArgs

func IdentityForCmkPtr(v *IdentityForCmkArgs) IdentityForCmkPtrInput {
	return (*identityForCmkPtrType)(v)
}

func (*identityForCmkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityForCmk)(nil)).Elem()
}

func (i *identityForCmkPtrType) ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput {
	return i.ToIdentityForCmkPtrOutputWithContext(context.Background())
}

func (i *identityForCmkPtrType) ToIdentityForCmkPtrOutputWithContext(ctx context.Context) IdentityForCmkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkPtrOutput)
}

type IdentityForCmkOutput struct{ *pulumi.OutputState }

func (IdentityForCmkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityForCmk)(nil)).Elem()
}

func (o IdentityForCmkOutput) ToIdentityForCmkOutput() IdentityForCmkOutput {
	return o
}

func (o IdentityForCmkOutput) ToIdentityForCmkOutputWithContext(ctx context.Context) IdentityForCmkOutput {
	return o
}

func (o IdentityForCmkOutput) ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput {
	return o.ToIdentityForCmkPtrOutputWithContext(context.Background())
}

func (o IdentityForCmkOutput) ToIdentityForCmkPtrOutputWithContext(ctx context.Context) IdentityForCmkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityForCmk) *IdentityForCmk {
		return &v
	}).(IdentityForCmkPtrOutput)
}

func (o IdentityForCmkOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityForCmk) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type IdentityForCmkPtrOutput struct{ *pulumi.OutputState }

func (IdentityForCmkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityForCmk)(nil)).Elem()
}

func (o IdentityForCmkPtrOutput) ToIdentityForCmkPtrOutput() IdentityForCmkPtrOutput {
	return o
}

func (o IdentityForCmkPtrOutput) ToIdentityForCmkPtrOutputWithContext(ctx context.Context) IdentityForCmkPtrOutput {
	return o
}

func (o IdentityForCmkPtrOutput) Elem() IdentityForCmkOutput {
	return o.ApplyT(func(v *IdentityForCmk) IdentityForCmk {
		if v != nil {
			return *v
		}
		var ret IdentityForCmk
		return ret
	}).(IdentityForCmkOutput)
}

func (o IdentityForCmkPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityForCmk) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type IdentityForCmkResponse struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type IdentityForCmkResponseInput interface {
	pulumi.Input

	ToIdentityForCmkResponseOutput() IdentityForCmkResponseOutput
	ToIdentityForCmkResponseOutputWithContext(context.Context) IdentityForCmkResponseOutput
}

type IdentityForCmkResponseArgs struct {
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (IdentityForCmkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityForCmkResponse)(nil)).Elem()
}

func (i IdentityForCmkResponseArgs) ToIdentityForCmkResponseOutput() IdentityForCmkResponseOutput {
	return i.ToIdentityForCmkResponseOutputWithContext(context.Background())
}

func (i IdentityForCmkResponseArgs) ToIdentityForCmkResponseOutputWithContext(ctx context.Context) IdentityForCmkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkResponseOutput)
}

func (i IdentityForCmkResponseArgs) ToIdentityForCmkResponsePtrOutput() IdentityForCmkResponsePtrOutput {
	return i.ToIdentityForCmkResponsePtrOutputWithContext(context.Background())
}

func (i IdentityForCmkResponseArgs) ToIdentityForCmkResponsePtrOutputWithContext(ctx context.Context) IdentityForCmkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkResponseOutput).ToIdentityForCmkResponsePtrOutputWithContext(ctx)
}









type IdentityForCmkResponsePtrInput interface {
	pulumi.Input

	ToIdentityForCmkResponsePtrOutput() IdentityForCmkResponsePtrOutput
	ToIdentityForCmkResponsePtrOutputWithContext(context.Context) IdentityForCmkResponsePtrOutput
}

type identityForCmkResponsePtrType IdentityForCmkResponseArgs

func IdentityForCmkResponsePtr(v *IdentityForCmkResponseArgs) IdentityForCmkResponsePtrInput {
	return (*identityForCmkResponsePtrType)(v)
}

func (*identityForCmkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityForCmkResponse)(nil)).Elem()
}

func (i *identityForCmkResponsePtrType) ToIdentityForCmkResponsePtrOutput() IdentityForCmkResponsePtrOutput {
	return i.ToIdentityForCmkResponsePtrOutputWithContext(context.Background())
}

func (i *identityForCmkResponsePtrType) ToIdentityForCmkResponsePtrOutputWithContext(ctx context.Context) IdentityForCmkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityForCmkResponsePtrOutput)
}

type IdentityForCmkResponseOutput struct{ *pulumi.OutputState }

func (IdentityForCmkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityForCmkResponse)(nil)).Elem()
}

func (o IdentityForCmkResponseOutput) ToIdentityForCmkResponseOutput() IdentityForCmkResponseOutput {
	return o
}

func (o IdentityForCmkResponseOutput) ToIdentityForCmkResponseOutputWithContext(ctx context.Context) IdentityForCmkResponseOutput {
	return o
}

func (o IdentityForCmkResponseOutput) ToIdentityForCmkResponsePtrOutput() IdentityForCmkResponsePtrOutput {
	return o.ToIdentityForCmkResponsePtrOutputWithContext(context.Background())
}

func (o IdentityForCmkResponseOutput) ToIdentityForCmkResponsePtrOutputWithContext(ctx context.Context) IdentityForCmkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityForCmkResponse) *IdentityForCmkResponse {
		return &v
	}).(IdentityForCmkResponsePtrOutput)
}

func (o IdentityForCmkResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityForCmkResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type IdentityForCmkResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityForCmkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityForCmkResponse)(nil)).Elem()
}

func (o IdentityForCmkResponsePtrOutput) ToIdentityForCmkResponsePtrOutput() IdentityForCmkResponsePtrOutput {
	return o
}

func (o IdentityForCmkResponsePtrOutput) ToIdentityForCmkResponsePtrOutputWithContext(ctx context.Context) IdentityForCmkResponsePtrOutput {
	return o
}

func (o IdentityForCmkResponsePtrOutput) Elem() IdentityForCmkResponseOutput {
	return o.ApplyT(func(v *IdentityForCmkResponse) IdentityForCmkResponse {
		if v != nil {
			return *v
		}
		var ret IdentityForCmkResponse
		return ret
	}).(IdentityForCmkResponseOutput)
}

func (o IdentityForCmkResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityForCmkResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                   `pulumi:"principalId"`
	TenantId               pulumi.StringInput                   `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
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

func (o IdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserAssignedIdentityResponse { return v.UserAssignedIdentities }).(UserAssignedIdentityResponseMapOutput)
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

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type InstanceTypeSchema struct {
	NodeSelector map[string]string            `pulumi:"nodeSelector"`
	Resources    *InstanceTypeSchemaResources `pulumi:"resources"`
}





type InstanceTypeSchemaInput interface {
	pulumi.Input

	ToInstanceTypeSchemaOutput() InstanceTypeSchemaOutput
	ToInstanceTypeSchemaOutputWithContext(context.Context) InstanceTypeSchemaOutput
}

type InstanceTypeSchemaArgs struct {
	NodeSelector pulumi.StringMapInput               `pulumi:"nodeSelector"`
	Resources    InstanceTypeSchemaResourcesPtrInput `pulumi:"resources"`
}

func (InstanceTypeSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTypeSchema)(nil)).Elem()
}

func (i InstanceTypeSchemaArgs) ToInstanceTypeSchemaOutput() InstanceTypeSchemaOutput {
	return i.ToInstanceTypeSchemaOutputWithContext(context.Background())
}

func (i InstanceTypeSchemaArgs) ToInstanceTypeSchemaOutputWithContext(ctx context.Context) InstanceTypeSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaOutput)
}





type InstanceTypeSchemaMapInput interface {
	pulumi.Input

	ToInstanceTypeSchemaMapOutput() InstanceTypeSchemaMapOutput
	ToInstanceTypeSchemaMapOutputWithContext(context.Context) InstanceTypeSchemaMapOutput
}

type InstanceTypeSchemaMap map[string]InstanceTypeSchemaInput

func (InstanceTypeSchemaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceTypeSchema)(nil)).Elem()
}

func (i InstanceTypeSchemaMap) ToInstanceTypeSchemaMapOutput() InstanceTypeSchemaMapOutput {
	return i.ToInstanceTypeSchemaMapOutputWithContext(context.Background())
}

func (i InstanceTypeSchemaMap) ToInstanceTypeSchemaMapOutputWithContext(ctx context.Context) InstanceTypeSchemaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaMapOutput)
}

type InstanceTypeSchemaOutput struct{ *pulumi.OutputState }

func (InstanceTypeSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTypeSchema)(nil)).Elem()
}

func (o InstanceTypeSchemaOutput) ToInstanceTypeSchemaOutput() InstanceTypeSchemaOutput {
	return o
}

func (o InstanceTypeSchemaOutput) ToInstanceTypeSchemaOutputWithContext(ctx context.Context) InstanceTypeSchemaOutput {
	return o
}

func (o InstanceTypeSchemaOutput) NodeSelector() pulumi.StringMapOutput {
	return o.ApplyT(func(v InstanceTypeSchema) map[string]string { return v.NodeSelector }).(pulumi.StringMapOutput)
}

func (o InstanceTypeSchemaOutput) Resources() InstanceTypeSchemaResourcesPtrOutput {
	return o.ApplyT(func(v InstanceTypeSchema) *InstanceTypeSchemaResources { return v.Resources }).(InstanceTypeSchemaResourcesPtrOutput)
}

type InstanceTypeSchemaMapOutput struct{ *pulumi.OutputState }

func (InstanceTypeSchemaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceTypeSchema)(nil)).Elem()
}

func (o InstanceTypeSchemaMapOutput) ToInstanceTypeSchemaMapOutput() InstanceTypeSchemaMapOutput {
	return o
}

func (o InstanceTypeSchemaMapOutput) ToInstanceTypeSchemaMapOutputWithContext(ctx context.Context) InstanceTypeSchemaMapOutput {
	return o
}

func (o InstanceTypeSchemaMapOutput) MapIndex(k pulumi.StringInput) InstanceTypeSchemaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceTypeSchema {
		return vs[0].(map[string]InstanceTypeSchema)[vs[1].(string)]
	}).(InstanceTypeSchemaOutput)
}

type InstanceTypeSchemaResources struct {
	Limits   map[string]string `pulumi:"limits"`
	Requests map[string]string `pulumi:"requests"`
}





type InstanceTypeSchemaResourcesInput interface {
	pulumi.Input

	ToInstanceTypeSchemaResourcesOutput() InstanceTypeSchemaResourcesOutput
	ToInstanceTypeSchemaResourcesOutputWithContext(context.Context) InstanceTypeSchemaResourcesOutput
}

type InstanceTypeSchemaResourcesArgs struct {
	Limits   pulumi.StringMapInput `pulumi:"limits"`
	Requests pulumi.StringMapInput `pulumi:"requests"`
}

func (InstanceTypeSchemaResourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTypeSchemaResources)(nil)).Elem()
}

func (i InstanceTypeSchemaResourcesArgs) ToInstanceTypeSchemaResourcesOutput() InstanceTypeSchemaResourcesOutput {
	return i.ToInstanceTypeSchemaResourcesOutputWithContext(context.Background())
}

func (i InstanceTypeSchemaResourcesArgs) ToInstanceTypeSchemaResourcesOutputWithContext(ctx context.Context) InstanceTypeSchemaResourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaResourcesOutput)
}

func (i InstanceTypeSchemaResourcesArgs) ToInstanceTypeSchemaResourcesPtrOutput() InstanceTypeSchemaResourcesPtrOutput {
	return i.ToInstanceTypeSchemaResourcesPtrOutputWithContext(context.Background())
}

func (i InstanceTypeSchemaResourcesArgs) ToInstanceTypeSchemaResourcesPtrOutputWithContext(ctx context.Context) InstanceTypeSchemaResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaResourcesOutput).ToInstanceTypeSchemaResourcesPtrOutputWithContext(ctx)
}









type InstanceTypeSchemaResourcesPtrInput interface {
	pulumi.Input

	ToInstanceTypeSchemaResourcesPtrOutput() InstanceTypeSchemaResourcesPtrOutput
	ToInstanceTypeSchemaResourcesPtrOutputWithContext(context.Context) InstanceTypeSchemaResourcesPtrOutput
}

type instanceTypeSchemaResourcesPtrType InstanceTypeSchemaResourcesArgs

func InstanceTypeSchemaResourcesPtr(v *InstanceTypeSchemaResourcesArgs) InstanceTypeSchemaResourcesPtrInput {
	return (*instanceTypeSchemaResourcesPtrType)(v)
}

func (*instanceTypeSchemaResourcesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTypeSchemaResources)(nil)).Elem()
}

func (i *instanceTypeSchemaResourcesPtrType) ToInstanceTypeSchemaResourcesPtrOutput() InstanceTypeSchemaResourcesPtrOutput {
	return i.ToInstanceTypeSchemaResourcesPtrOutputWithContext(context.Background())
}

func (i *instanceTypeSchemaResourcesPtrType) ToInstanceTypeSchemaResourcesPtrOutputWithContext(ctx context.Context) InstanceTypeSchemaResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaResourcesPtrOutput)
}

type InstanceTypeSchemaResourcesOutput struct{ *pulumi.OutputState }

func (InstanceTypeSchemaResourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTypeSchemaResources)(nil)).Elem()
}

func (o InstanceTypeSchemaResourcesOutput) ToInstanceTypeSchemaResourcesOutput() InstanceTypeSchemaResourcesOutput {
	return o
}

func (o InstanceTypeSchemaResourcesOutput) ToInstanceTypeSchemaResourcesOutputWithContext(ctx context.Context) InstanceTypeSchemaResourcesOutput {
	return o
}

func (o InstanceTypeSchemaResourcesOutput) ToInstanceTypeSchemaResourcesPtrOutput() InstanceTypeSchemaResourcesPtrOutput {
	return o.ToInstanceTypeSchemaResourcesPtrOutputWithContext(context.Background())
}

func (o InstanceTypeSchemaResourcesOutput) ToInstanceTypeSchemaResourcesPtrOutputWithContext(ctx context.Context) InstanceTypeSchemaResourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceTypeSchemaResources) *InstanceTypeSchemaResources {
		return &v
	}).(InstanceTypeSchemaResourcesPtrOutput)
}

func (o InstanceTypeSchemaResourcesOutput) Limits() pulumi.StringMapOutput {
	return o.ApplyT(func(v InstanceTypeSchemaResources) map[string]string { return v.Limits }).(pulumi.StringMapOutput)
}

func (o InstanceTypeSchemaResourcesOutput) Requests() pulumi.StringMapOutput {
	return o.ApplyT(func(v InstanceTypeSchemaResources) map[string]string { return v.Requests }).(pulumi.StringMapOutput)
}

type InstanceTypeSchemaResourcesPtrOutput struct{ *pulumi.OutputState }

func (InstanceTypeSchemaResourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTypeSchemaResources)(nil)).Elem()
}

func (o InstanceTypeSchemaResourcesPtrOutput) ToInstanceTypeSchemaResourcesPtrOutput() InstanceTypeSchemaResourcesPtrOutput {
	return o
}

func (o InstanceTypeSchemaResourcesPtrOutput) ToInstanceTypeSchemaResourcesPtrOutputWithContext(ctx context.Context) InstanceTypeSchemaResourcesPtrOutput {
	return o
}

func (o InstanceTypeSchemaResourcesPtrOutput) Elem() InstanceTypeSchemaResourcesOutput {
	return o.ApplyT(func(v *InstanceTypeSchemaResources) InstanceTypeSchemaResources {
		if v != nil {
			return *v
		}
		var ret InstanceTypeSchemaResources
		return ret
	}).(InstanceTypeSchemaResourcesOutput)
}

func (o InstanceTypeSchemaResourcesPtrOutput) Limits() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstanceTypeSchemaResources) map[string]string {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(pulumi.StringMapOutput)
}

func (o InstanceTypeSchemaResourcesPtrOutput) Requests() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstanceTypeSchemaResources) map[string]string {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(pulumi.StringMapOutput)
}

type InstanceTypeSchemaResponse struct {
	NodeSelector map[string]string                    `pulumi:"nodeSelector"`
	Resources    *InstanceTypeSchemaResponseResources `pulumi:"resources"`
}





type InstanceTypeSchemaResponseInput interface {
	pulumi.Input

	ToInstanceTypeSchemaResponseOutput() InstanceTypeSchemaResponseOutput
	ToInstanceTypeSchemaResponseOutputWithContext(context.Context) InstanceTypeSchemaResponseOutput
}

type InstanceTypeSchemaResponseArgs struct {
	NodeSelector pulumi.StringMapInput                       `pulumi:"nodeSelector"`
	Resources    InstanceTypeSchemaResponseResourcesPtrInput `pulumi:"resources"`
}

func (InstanceTypeSchemaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTypeSchemaResponse)(nil)).Elem()
}

func (i InstanceTypeSchemaResponseArgs) ToInstanceTypeSchemaResponseOutput() InstanceTypeSchemaResponseOutput {
	return i.ToInstanceTypeSchemaResponseOutputWithContext(context.Background())
}

func (i InstanceTypeSchemaResponseArgs) ToInstanceTypeSchemaResponseOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaResponseOutput)
}





type InstanceTypeSchemaResponseMapInput interface {
	pulumi.Input

	ToInstanceTypeSchemaResponseMapOutput() InstanceTypeSchemaResponseMapOutput
	ToInstanceTypeSchemaResponseMapOutputWithContext(context.Context) InstanceTypeSchemaResponseMapOutput
}

type InstanceTypeSchemaResponseMap map[string]InstanceTypeSchemaResponseInput

func (InstanceTypeSchemaResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceTypeSchemaResponse)(nil)).Elem()
}

func (i InstanceTypeSchemaResponseMap) ToInstanceTypeSchemaResponseMapOutput() InstanceTypeSchemaResponseMapOutput {
	return i.ToInstanceTypeSchemaResponseMapOutputWithContext(context.Background())
}

func (i InstanceTypeSchemaResponseMap) ToInstanceTypeSchemaResponseMapOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaResponseMapOutput)
}

type InstanceTypeSchemaResponseOutput struct{ *pulumi.OutputState }

func (InstanceTypeSchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTypeSchemaResponse)(nil)).Elem()
}

func (o InstanceTypeSchemaResponseOutput) ToInstanceTypeSchemaResponseOutput() InstanceTypeSchemaResponseOutput {
	return o
}

func (o InstanceTypeSchemaResponseOutput) ToInstanceTypeSchemaResponseOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseOutput {
	return o
}

func (o InstanceTypeSchemaResponseOutput) NodeSelector() pulumi.StringMapOutput {
	return o.ApplyT(func(v InstanceTypeSchemaResponse) map[string]string { return v.NodeSelector }).(pulumi.StringMapOutput)
}

func (o InstanceTypeSchemaResponseOutput) Resources() InstanceTypeSchemaResponseResourcesPtrOutput {
	return o.ApplyT(func(v InstanceTypeSchemaResponse) *InstanceTypeSchemaResponseResources { return v.Resources }).(InstanceTypeSchemaResponseResourcesPtrOutput)
}

type InstanceTypeSchemaResponseMapOutput struct{ *pulumi.OutputState }

func (InstanceTypeSchemaResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceTypeSchemaResponse)(nil)).Elem()
}

func (o InstanceTypeSchemaResponseMapOutput) ToInstanceTypeSchemaResponseMapOutput() InstanceTypeSchemaResponseMapOutput {
	return o
}

func (o InstanceTypeSchemaResponseMapOutput) ToInstanceTypeSchemaResponseMapOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseMapOutput {
	return o
}

func (o InstanceTypeSchemaResponseMapOutput) MapIndex(k pulumi.StringInput) InstanceTypeSchemaResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceTypeSchemaResponse {
		return vs[0].(map[string]InstanceTypeSchemaResponse)[vs[1].(string)]
	}).(InstanceTypeSchemaResponseOutput)
}

type InstanceTypeSchemaResponseResources struct {
	Limits   map[string]string `pulumi:"limits"`
	Requests map[string]string `pulumi:"requests"`
}





type InstanceTypeSchemaResponseResourcesInput interface {
	pulumi.Input

	ToInstanceTypeSchemaResponseResourcesOutput() InstanceTypeSchemaResponseResourcesOutput
	ToInstanceTypeSchemaResponseResourcesOutputWithContext(context.Context) InstanceTypeSchemaResponseResourcesOutput
}

type InstanceTypeSchemaResponseResourcesArgs struct {
	Limits   pulumi.StringMapInput `pulumi:"limits"`
	Requests pulumi.StringMapInput `pulumi:"requests"`
}

func (InstanceTypeSchemaResponseResourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTypeSchemaResponseResources)(nil)).Elem()
}

func (i InstanceTypeSchemaResponseResourcesArgs) ToInstanceTypeSchemaResponseResourcesOutput() InstanceTypeSchemaResponseResourcesOutput {
	return i.ToInstanceTypeSchemaResponseResourcesOutputWithContext(context.Background())
}

func (i InstanceTypeSchemaResponseResourcesArgs) ToInstanceTypeSchemaResponseResourcesOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseResourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaResponseResourcesOutput)
}

func (i InstanceTypeSchemaResponseResourcesArgs) ToInstanceTypeSchemaResponseResourcesPtrOutput() InstanceTypeSchemaResponseResourcesPtrOutput {
	return i.ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(context.Background())
}

func (i InstanceTypeSchemaResponseResourcesArgs) ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaResponseResourcesOutput).ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(ctx)
}









type InstanceTypeSchemaResponseResourcesPtrInput interface {
	pulumi.Input

	ToInstanceTypeSchemaResponseResourcesPtrOutput() InstanceTypeSchemaResponseResourcesPtrOutput
	ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(context.Context) InstanceTypeSchemaResponseResourcesPtrOutput
}

type instanceTypeSchemaResponseResourcesPtrType InstanceTypeSchemaResponseResourcesArgs

func InstanceTypeSchemaResponseResourcesPtr(v *InstanceTypeSchemaResponseResourcesArgs) InstanceTypeSchemaResponseResourcesPtrInput {
	return (*instanceTypeSchemaResponseResourcesPtrType)(v)
}

func (*instanceTypeSchemaResponseResourcesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTypeSchemaResponseResources)(nil)).Elem()
}

func (i *instanceTypeSchemaResponseResourcesPtrType) ToInstanceTypeSchemaResponseResourcesPtrOutput() InstanceTypeSchemaResponseResourcesPtrOutput {
	return i.ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(context.Background())
}

func (i *instanceTypeSchemaResponseResourcesPtrType) ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTypeSchemaResponseResourcesPtrOutput)
}

type InstanceTypeSchemaResponseResourcesOutput struct{ *pulumi.OutputState }

func (InstanceTypeSchemaResponseResourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTypeSchemaResponseResources)(nil)).Elem()
}

func (o InstanceTypeSchemaResponseResourcesOutput) ToInstanceTypeSchemaResponseResourcesOutput() InstanceTypeSchemaResponseResourcesOutput {
	return o
}

func (o InstanceTypeSchemaResponseResourcesOutput) ToInstanceTypeSchemaResponseResourcesOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseResourcesOutput {
	return o
}

func (o InstanceTypeSchemaResponseResourcesOutput) ToInstanceTypeSchemaResponseResourcesPtrOutput() InstanceTypeSchemaResponseResourcesPtrOutput {
	return o.ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(context.Background())
}

func (o InstanceTypeSchemaResponseResourcesOutput) ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseResourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceTypeSchemaResponseResources) *InstanceTypeSchemaResponseResources {
		return &v
	}).(InstanceTypeSchemaResponseResourcesPtrOutput)
}

func (o InstanceTypeSchemaResponseResourcesOutput) Limits() pulumi.StringMapOutput {
	return o.ApplyT(func(v InstanceTypeSchemaResponseResources) map[string]string { return v.Limits }).(pulumi.StringMapOutput)
}

func (o InstanceTypeSchemaResponseResourcesOutput) Requests() pulumi.StringMapOutput {
	return o.ApplyT(func(v InstanceTypeSchemaResponseResources) map[string]string { return v.Requests }).(pulumi.StringMapOutput)
}

type InstanceTypeSchemaResponseResourcesPtrOutput struct{ *pulumi.OutputState }

func (InstanceTypeSchemaResponseResourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTypeSchemaResponseResources)(nil)).Elem()
}

func (o InstanceTypeSchemaResponseResourcesPtrOutput) ToInstanceTypeSchemaResponseResourcesPtrOutput() InstanceTypeSchemaResponseResourcesPtrOutput {
	return o
}

func (o InstanceTypeSchemaResponseResourcesPtrOutput) ToInstanceTypeSchemaResponseResourcesPtrOutputWithContext(ctx context.Context) InstanceTypeSchemaResponseResourcesPtrOutput {
	return o
}

func (o InstanceTypeSchemaResponseResourcesPtrOutput) Elem() InstanceTypeSchemaResponseResourcesOutput {
	return o.ApplyT(func(v *InstanceTypeSchemaResponseResources) InstanceTypeSchemaResponseResources {
		if v != nil {
			return *v
		}
		var ret InstanceTypeSchemaResponseResources
		return ret
	}).(InstanceTypeSchemaResponseResourcesOutput)
}

func (o InstanceTypeSchemaResponseResourcesPtrOutput) Limits() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstanceTypeSchemaResponseResources) map[string]string {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(pulumi.StringMapOutput)
}

func (o InstanceTypeSchemaResponseResourcesPtrOutput) Requests() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstanceTypeSchemaResponseResources) map[string]string {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(pulumi.StringMapOutput)
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

type Kubernetes struct {
	ComputeLocation  *string               `pulumi:"computeLocation"`
	ComputeType      string                `pulumi:"computeType"`
	Description      *string               `pulumi:"description"`
	DisableLocalAuth *bool                 `pulumi:"disableLocalAuth"`
	Properties       *KubernetesProperties `pulumi:"properties"`
	ResourceId       *string               `pulumi:"resourceId"`
}





type KubernetesInput interface {
	pulumi.Input

	ToKubernetesOutput() KubernetesOutput
	ToKubernetesOutputWithContext(context.Context) KubernetesOutput
}

type KubernetesArgs struct {
	ComputeLocation  pulumi.StringPtrInput        `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput           `pulumi:"computeType"`
	Description      pulumi.StringPtrInput        `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput          `pulumi:"disableLocalAuth"`
	Properties       KubernetesPropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput        `pulumi:"resourceId"`
}

func (KubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Kubernetes)(nil)).Elem()
}

func (i KubernetesArgs) ToKubernetesOutput() KubernetesOutput {
	return i.ToKubernetesOutputWithContext(context.Background())
}

func (i KubernetesArgs) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesOutput)
}

type KubernetesOutput struct{ *pulumi.OutputState }

func (KubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Kubernetes)(nil)).Elem()
}

func (o KubernetesOutput) ToKubernetesOutput() KubernetesOutput {
	return o
}

func (o KubernetesOutput) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return o
}

func (o KubernetesOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Kubernetes) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o KubernetesOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v Kubernetes) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o KubernetesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Kubernetes) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o KubernetesOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Kubernetes) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o KubernetesOutput) Properties() KubernetesPropertiesPtrOutput {
	return o.ApplyT(func(v Kubernetes) *KubernetesProperties { return v.Properties }).(KubernetesPropertiesPtrOutput)
}

func (o KubernetesOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Kubernetes) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type KubernetesProperties struct {
	DefaultInstanceType           *string                       `pulumi:"defaultInstanceType"`
	ExtensionInstanceReleaseTrain *string                       `pulumi:"extensionInstanceReleaseTrain"`
	ExtensionPrincipalId          *string                       `pulumi:"extensionPrincipalId"`
	InstanceTypes                 map[string]InstanceTypeSchema `pulumi:"instanceTypes"`
	Namespace                     *string                       `pulumi:"namespace"`
	RelayConnectionString         *string                       `pulumi:"relayConnectionString"`
	ServiceBusConnectionString    *string                       `pulumi:"serviceBusConnectionString"`
	VcName                        *string                       `pulumi:"vcName"`
}





type KubernetesPropertiesInput interface {
	pulumi.Input

	ToKubernetesPropertiesOutput() KubernetesPropertiesOutput
	ToKubernetesPropertiesOutputWithContext(context.Context) KubernetesPropertiesOutput
}

type KubernetesPropertiesArgs struct {
	DefaultInstanceType           pulumi.StringPtrInput      `pulumi:"defaultInstanceType"`
	ExtensionInstanceReleaseTrain pulumi.StringPtrInput      `pulumi:"extensionInstanceReleaseTrain"`
	ExtensionPrincipalId          pulumi.StringPtrInput      `pulumi:"extensionPrincipalId"`
	InstanceTypes                 InstanceTypeSchemaMapInput `pulumi:"instanceTypes"`
	Namespace                     pulumi.StringPtrInput      `pulumi:"namespace"`
	RelayConnectionString         pulumi.StringPtrInput      `pulumi:"relayConnectionString"`
	ServiceBusConnectionString    pulumi.StringPtrInput      `pulumi:"serviceBusConnectionString"`
	VcName                        pulumi.StringPtrInput      `pulumi:"vcName"`
}

func (KubernetesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesProperties)(nil)).Elem()
}

func (i KubernetesPropertiesArgs) ToKubernetesPropertiesOutput() KubernetesPropertiesOutput {
	return i.ToKubernetesPropertiesOutputWithContext(context.Background())
}

func (i KubernetesPropertiesArgs) ToKubernetesPropertiesOutputWithContext(ctx context.Context) KubernetesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesPropertiesOutput)
}

func (i KubernetesPropertiesArgs) ToKubernetesPropertiesPtrOutput() KubernetesPropertiesPtrOutput {
	return i.ToKubernetesPropertiesPtrOutputWithContext(context.Background())
}

func (i KubernetesPropertiesArgs) ToKubernetesPropertiesPtrOutputWithContext(ctx context.Context) KubernetesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesPropertiesOutput).ToKubernetesPropertiesPtrOutputWithContext(ctx)
}









type KubernetesPropertiesPtrInput interface {
	pulumi.Input

	ToKubernetesPropertiesPtrOutput() KubernetesPropertiesPtrOutput
	ToKubernetesPropertiesPtrOutputWithContext(context.Context) KubernetesPropertiesPtrOutput
}

type kubernetesPropertiesPtrType KubernetesPropertiesArgs

func KubernetesPropertiesPtr(v *KubernetesPropertiesArgs) KubernetesPropertiesPtrInput {
	return (*kubernetesPropertiesPtrType)(v)
}

func (*kubernetesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesProperties)(nil)).Elem()
}

func (i *kubernetesPropertiesPtrType) ToKubernetesPropertiesPtrOutput() KubernetesPropertiesPtrOutput {
	return i.ToKubernetesPropertiesPtrOutputWithContext(context.Background())
}

func (i *kubernetesPropertiesPtrType) ToKubernetesPropertiesPtrOutputWithContext(ctx context.Context) KubernetesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesPropertiesPtrOutput)
}

type KubernetesPropertiesOutput struct{ *pulumi.OutputState }

func (KubernetesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesProperties)(nil)).Elem()
}

func (o KubernetesPropertiesOutput) ToKubernetesPropertiesOutput() KubernetesPropertiesOutput {
	return o
}

func (o KubernetesPropertiesOutput) ToKubernetesPropertiesOutputWithContext(ctx context.Context) KubernetesPropertiesOutput {
	return o
}

func (o KubernetesPropertiesOutput) ToKubernetesPropertiesPtrOutput() KubernetesPropertiesPtrOutput {
	return o.ToKubernetesPropertiesPtrOutputWithContext(context.Background())
}

func (o KubernetesPropertiesOutput) ToKubernetesPropertiesPtrOutputWithContext(ctx context.Context) KubernetesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesProperties) *KubernetesProperties {
		return &v
	}).(KubernetesPropertiesPtrOutput)
}

func (o KubernetesPropertiesOutput) DefaultInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesProperties) *string { return v.DefaultInstanceType }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesOutput) ExtensionInstanceReleaseTrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesProperties) *string { return v.ExtensionInstanceReleaseTrain }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesOutput) ExtensionPrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesProperties) *string { return v.ExtensionPrincipalId }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesOutput) InstanceTypes() InstanceTypeSchemaMapOutput {
	return o.ApplyT(func(v KubernetesProperties) map[string]InstanceTypeSchema { return v.InstanceTypes }).(InstanceTypeSchemaMapOutput)
}

func (o KubernetesPropertiesOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesProperties) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesOutput) RelayConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesProperties) *string { return v.RelayConnectionString }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesOutput) ServiceBusConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesProperties) *string { return v.ServiceBusConnectionString }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesOutput) VcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesProperties) *string { return v.VcName }).(pulumi.StringPtrOutput)
}

type KubernetesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KubernetesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesProperties)(nil)).Elem()
}

func (o KubernetesPropertiesPtrOutput) ToKubernetesPropertiesPtrOutput() KubernetesPropertiesPtrOutput {
	return o
}

func (o KubernetesPropertiesPtrOutput) ToKubernetesPropertiesPtrOutputWithContext(ctx context.Context) KubernetesPropertiesPtrOutput {
	return o
}

func (o KubernetesPropertiesPtrOutput) Elem() KubernetesPropertiesOutput {
	return o.ApplyT(func(v *KubernetesProperties) KubernetesProperties {
		if v != nil {
			return *v
		}
		var ret KubernetesProperties
		return ret
	}).(KubernetesPropertiesOutput)
}

func (o KubernetesPropertiesPtrOutput) DefaultInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesProperties) *string {
		if v == nil {
			return nil
		}
		return v.DefaultInstanceType
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesPtrOutput) ExtensionInstanceReleaseTrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesProperties) *string {
		if v == nil {
			return nil
		}
		return v.ExtensionInstanceReleaseTrain
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesPtrOutput) ExtensionPrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesProperties) *string {
		if v == nil {
			return nil
		}
		return v.ExtensionPrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesPtrOutput) InstanceTypes() InstanceTypeSchemaMapOutput {
	return o.ApplyT(func(v *KubernetesProperties) map[string]InstanceTypeSchema {
		if v == nil {
			return nil
		}
		return v.InstanceTypes
	}).(InstanceTypeSchemaMapOutput)
}

func (o KubernetesPropertiesPtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesProperties) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesPtrOutput) RelayConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesProperties) *string {
		if v == nil {
			return nil
		}
		return v.RelayConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesPtrOutput) ServiceBusConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServiceBusConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesPtrOutput) VcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesProperties) *string {
		if v == nil {
			return nil
		}
		return v.VcName
	}).(pulumi.StringPtrOutput)
}

type KubernetesPropertiesResponse struct {
	DefaultInstanceType           *string                               `pulumi:"defaultInstanceType"`
	ExtensionInstanceReleaseTrain *string                               `pulumi:"extensionInstanceReleaseTrain"`
	ExtensionPrincipalId          *string                               `pulumi:"extensionPrincipalId"`
	InstanceTypes                 map[string]InstanceTypeSchemaResponse `pulumi:"instanceTypes"`
	Namespace                     *string                               `pulumi:"namespace"`
	RelayConnectionString         *string                               `pulumi:"relayConnectionString"`
	ServiceBusConnectionString    *string                               `pulumi:"serviceBusConnectionString"`
	VcName                        *string                               `pulumi:"vcName"`
}





type KubernetesPropertiesResponseInput interface {
	pulumi.Input

	ToKubernetesPropertiesResponseOutput() KubernetesPropertiesResponseOutput
	ToKubernetesPropertiesResponseOutputWithContext(context.Context) KubernetesPropertiesResponseOutput
}

type KubernetesPropertiesResponseArgs struct {
	DefaultInstanceType           pulumi.StringPtrInput              `pulumi:"defaultInstanceType"`
	ExtensionInstanceReleaseTrain pulumi.StringPtrInput              `pulumi:"extensionInstanceReleaseTrain"`
	ExtensionPrincipalId          pulumi.StringPtrInput              `pulumi:"extensionPrincipalId"`
	InstanceTypes                 InstanceTypeSchemaResponseMapInput `pulumi:"instanceTypes"`
	Namespace                     pulumi.StringPtrInput              `pulumi:"namespace"`
	RelayConnectionString         pulumi.StringPtrInput              `pulumi:"relayConnectionString"`
	ServiceBusConnectionString    pulumi.StringPtrInput              `pulumi:"serviceBusConnectionString"`
	VcName                        pulumi.StringPtrInput              `pulumi:"vcName"`
}

func (KubernetesPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesPropertiesResponse)(nil)).Elem()
}

func (i KubernetesPropertiesResponseArgs) ToKubernetesPropertiesResponseOutput() KubernetesPropertiesResponseOutput {
	return i.ToKubernetesPropertiesResponseOutputWithContext(context.Background())
}

func (i KubernetesPropertiesResponseArgs) ToKubernetesPropertiesResponseOutputWithContext(ctx context.Context) KubernetesPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesPropertiesResponseOutput)
}

func (i KubernetesPropertiesResponseArgs) ToKubernetesPropertiesResponsePtrOutput() KubernetesPropertiesResponsePtrOutput {
	return i.ToKubernetesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KubernetesPropertiesResponseArgs) ToKubernetesPropertiesResponsePtrOutputWithContext(ctx context.Context) KubernetesPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesPropertiesResponseOutput).ToKubernetesPropertiesResponsePtrOutputWithContext(ctx)
}









type KubernetesPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKubernetesPropertiesResponsePtrOutput() KubernetesPropertiesResponsePtrOutput
	ToKubernetesPropertiesResponsePtrOutputWithContext(context.Context) KubernetesPropertiesResponsePtrOutput
}

type kubernetesPropertiesResponsePtrType KubernetesPropertiesResponseArgs

func KubernetesPropertiesResponsePtr(v *KubernetesPropertiesResponseArgs) KubernetesPropertiesResponsePtrInput {
	return (*kubernetesPropertiesResponsePtrType)(v)
}

func (*kubernetesPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesPropertiesResponse)(nil)).Elem()
}

func (i *kubernetesPropertiesResponsePtrType) ToKubernetesPropertiesResponsePtrOutput() KubernetesPropertiesResponsePtrOutput {
	return i.ToKubernetesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *kubernetesPropertiesResponsePtrType) ToKubernetesPropertiesResponsePtrOutputWithContext(ctx context.Context) KubernetesPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesPropertiesResponsePtrOutput)
}

type KubernetesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KubernetesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesPropertiesResponse)(nil)).Elem()
}

func (o KubernetesPropertiesResponseOutput) ToKubernetesPropertiesResponseOutput() KubernetesPropertiesResponseOutput {
	return o
}

func (o KubernetesPropertiesResponseOutput) ToKubernetesPropertiesResponseOutputWithContext(ctx context.Context) KubernetesPropertiesResponseOutput {
	return o
}

func (o KubernetesPropertiesResponseOutput) ToKubernetesPropertiesResponsePtrOutput() KubernetesPropertiesResponsePtrOutput {
	return o.ToKubernetesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KubernetesPropertiesResponseOutput) ToKubernetesPropertiesResponsePtrOutputWithContext(ctx context.Context) KubernetesPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubernetesPropertiesResponse) *KubernetesPropertiesResponse {
		return &v
	}).(KubernetesPropertiesResponsePtrOutput)
}

func (o KubernetesPropertiesResponseOutput) DefaultInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesPropertiesResponse) *string { return v.DefaultInstanceType }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponseOutput) ExtensionInstanceReleaseTrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesPropertiesResponse) *string { return v.ExtensionInstanceReleaseTrain }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponseOutput) ExtensionPrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesPropertiesResponse) *string { return v.ExtensionPrincipalId }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponseOutput) InstanceTypes() InstanceTypeSchemaResponseMapOutput {
	return o.ApplyT(func(v KubernetesPropertiesResponse) map[string]InstanceTypeSchemaResponse { return v.InstanceTypes }).(InstanceTypeSchemaResponseMapOutput)
}

func (o KubernetesPropertiesResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesPropertiesResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponseOutput) RelayConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesPropertiesResponse) *string { return v.RelayConnectionString }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponseOutput) ServiceBusConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesPropertiesResponse) *string { return v.ServiceBusConnectionString }).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponseOutput) VcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesPropertiesResponse) *string { return v.VcName }).(pulumi.StringPtrOutput)
}

type KubernetesPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesPropertiesResponse)(nil)).Elem()
}

func (o KubernetesPropertiesResponsePtrOutput) ToKubernetesPropertiesResponsePtrOutput() KubernetesPropertiesResponsePtrOutput {
	return o
}

func (o KubernetesPropertiesResponsePtrOutput) ToKubernetesPropertiesResponsePtrOutputWithContext(ctx context.Context) KubernetesPropertiesResponsePtrOutput {
	return o
}

func (o KubernetesPropertiesResponsePtrOutput) Elem() KubernetesPropertiesResponseOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) KubernetesPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesPropertiesResponse
		return ret
	}).(KubernetesPropertiesResponseOutput)
}

func (o KubernetesPropertiesResponsePtrOutput) DefaultInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultInstanceType
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponsePtrOutput) ExtensionInstanceReleaseTrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExtensionInstanceReleaseTrain
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponsePtrOutput) ExtensionPrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExtensionPrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponsePtrOutput) InstanceTypes() InstanceTypeSchemaResponseMapOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) map[string]InstanceTypeSchemaResponse {
		if v == nil {
			return nil
		}
		return v.InstanceTypes
	}).(InstanceTypeSchemaResponseMapOutput)
}

func (o KubernetesPropertiesResponsePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponsePtrOutput) RelayConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RelayConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponsePtrOutput) ServiceBusConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceBusConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o KubernetesPropertiesResponsePtrOutput) VcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VcName
	}).(pulumi.StringPtrOutput)
}

type KubernetesResponse struct {
	ComputeLocation    *string                       `pulumi:"computeLocation"`
	ComputeType        string                        `pulumi:"computeType"`
	CreatedOn          string                        `pulumi:"createdOn"`
	Description        *string                       `pulumi:"description"`
	DisableLocalAuth   *bool                         `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                          `pulumi:"isAttachedCompute"`
	ModifiedOn         string                        `pulumi:"modifiedOn"`
	Properties         *KubernetesPropertiesResponse `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse       `pulumi:"provisioningErrors"`
	ProvisioningState  string                        `pulumi:"provisioningState"`
	ResourceId         *string                       `pulumi:"resourceId"`
}





type KubernetesResponseInput interface {
	pulumi.Input

	ToKubernetesResponseOutput() KubernetesResponseOutput
	ToKubernetesResponseOutputWithContext(context.Context) KubernetesResponseOutput
}

type KubernetesResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                   `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                   `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput                  `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                     `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                   `pulumi:"modifiedOn"`
	Properties         KubernetesPropertiesResponsePtrInput `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput      `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                   `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                `pulumi:"resourceId"`
}

func (KubernetesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesResponse)(nil)).Elem()
}

func (i KubernetesResponseArgs) ToKubernetesResponseOutput() KubernetesResponseOutput {
	return i.ToKubernetesResponseOutputWithContext(context.Background())
}

func (i KubernetesResponseArgs) ToKubernetesResponseOutputWithContext(ctx context.Context) KubernetesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesResponseOutput)
}

type KubernetesResponseOutput struct{ *pulumi.OutputState }

func (KubernetesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesResponse)(nil)).Elem()
}

func (o KubernetesResponseOutput) ToKubernetesResponseOutput() KubernetesResponseOutput {
	return o
}

func (o KubernetesResponseOutput) ToKubernetesResponseOutputWithContext(ctx context.Context) KubernetesResponseOutput {
	return o
}

func (o KubernetesResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o KubernetesResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o KubernetesResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o KubernetesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o KubernetesResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KubernetesResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o KubernetesResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v KubernetesResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o KubernetesResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o KubernetesResponseOutput) Properties() KubernetesPropertiesResponsePtrOutput {
	return o.ApplyT(func(v KubernetesResponse) *KubernetesPropertiesResponse { return v.Properties }).(KubernetesPropertiesResponsePtrOutput)
}

func (o KubernetesResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v KubernetesResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o KubernetesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KubernetesResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ListNotebookKeysResultResponse struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}





type ListNotebookKeysResultResponseInput interface {
	pulumi.Input

	ToListNotebookKeysResultResponseOutput() ListNotebookKeysResultResponseOutput
	ToListNotebookKeysResultResponseOutputWithContext(context.Context) ListNotebookKeysResultResponseOutput
}

type ListNotebookKeysResultResponseArgs struct {
	PrimaryAccessKey   pulumi.StringInput `pulumi:"primaryAccessKey"`
	SecondaryAccessKey pulumi.StringInput `pulumi:"secondaryAccessKey"`
}

func (ListNotebookKeysResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookKeysResultResponse)(nil)).Elem()
}

func (i ListNotebookKeysResultResponseArgs) ToListNotebookKeysResultResponseOutput() ListNotebookKeysResultResponseOutput {
	return i.ToListNotebookKeysResultResponseOutputWithContext(context.Background())
}

func (i ListNotebookKeysResultResponseArgs) ToListNotebookKeysResultResponseOutputWithContext(ctx context.Context) ListNotebookKeysResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListNotebookKeysResultResponseOutput)
}

type ListNotebookKeysResultResponseOutput struct{ *pulumi.OutputState }

func (ListNotebookKeysResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookKeysResultResponse)(nil)).Elem()
}

func (o ListNotebookKeysResultResponseOutput) ToListNotebookKeysResultResponseOutput() ListNotebookKeysResultResponseOutput {
	return o
}

func (o ListNotebookKeysResultResponseOutput) ToListNotebookKeysResultResponseOutputWithContext(ctx context.Context) ListNotebookKeysResultResponseOutput {
	return o
}

func (o ListNotebookKeysResultResponseOutput) PrimaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookKeysResultResponse) string { return v.PrimaryAccessKey }).(pulumi.StringOutput)
}

func (o ListNotebookKeysResultResponseOutput) SecondaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookKeysResultResponse) string { return v.SecondaryAccessKey }).(pulumi.StringOutput)
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

type NotebookPreparationErrorResponse struct {
	ErrorMessage *string `pulumi:"errorMessage"`
	StatusCode   *int    `pulumi:"statusCode"`
}





type NotebookPreparationErrorResponseInput interface {
	pulumi.Input

	ToNotebookPreparationErrorResponseOutput() NotebookPreparationErrorResponseOutput
	ToNotebookPreparationErrorResponseOutputWithContext(context.Context) NotebookPreparationErrorResponseOutput
}

type NotebookPreparationErrorResponseArgs struct {
	ErrorMessage pulumi.StringPtrInput `pulumi:"errorMessage"`
	StatusCode   pulumi.IntPtrInput    `pulumi:"statusCode"`
}

func (NotebookPreparationErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookPreparationErrorResponse)(nil)).Elem()
}

func (i NotebookPreparationErrorResponseArgs) ToNotebookPreparationErrorResponseOutput() NotebookPreparationErrorResponseOutput {
	return i.ToNotebookPreparationErrorResponseOutputWithContext(context.Background())
}

func (i NotebookPreparationErrorResponseArgs) ToNotebookPreparationErrorResponseOutputWithContext(ctx context.Context) NotebookPreparationErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookPreparationErrorResponseOutput)
}

func (i NotebookPreparationErrorResponseArgs) ToNotebookPreparationErrorResponsePtrOutput() NotebookPreparationErrorResponsePtrOutput {
	return i.ToNotebookPreparationErrorResponsePtrOutputWithContext(context.Background())
}

func (i NotebookPreparationErrorResponseArgs) ToNotebookPreparationErrorResponsePtrOutputWithContext(ctx context.Context) NotebookPreparationErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookPreparationErrorResponseOutput).ToNotebookPreparationErrorResponsePtrOutputWithContext(ctx)
}









type NotebookPreparationErrorResponsePtrInput interface {
	pulumi.Input

	ToNotebookPreparationErrorResponsePtrOutput() NotebookPreparationErrorResponsePtrOutput
	ToNotebookPreparationErrorResponsePtrOutputWithContext(context.Context) NotebookPreparationErrorResponsePtrOutput
}

type notebookPreparationErrorResponsePtrType NotebookPreparationErrorResponseArgs

func NotebookPreparationErrorResponsePtr(v *NotebookPreparationErrorResponseArgs) NotebookPreparationErrorResponsePtrInput {
	return (*notebookPreparationErrorResponsePtrType)(v)
}

func (*notebookPreparationErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookPreparationErrorResponse)(nil)).Elem()
}

func (i *notebookPreparationErrorResponsePtrType) ToNotebookPreparationErrorResponsePtrOutput() NotebookPreparationErrorResponsePtrOutput {
	return i.ToNotebookPreparationErrorResponsePtrOutputWithContext(context.Background())
}

func (i *notebookPreparationErrorResponsePtrType) ToNotebookPreparationErrorResponsePtrOutputWithContext(ctx context.Context) NotebookPreparationErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookPreparationErrorResponsePtrOutput)
}

type NotebookPreparationErrorResponseOutput struct{ *pulumi.OutputState }

func (NotebookPreparationErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookPreparationErrorResponse)(nil)).Elem()
}

func (o NotebookPreparationErrorResponseOutput) ToNotebookPreparationErrorResponseOutput() NotebookPreparationErrorResponseOutput {
	return o
}

func (o NotebookPreparationErrorResponseOutput) ToNotebookPreparationErrorResponseOutputWithContext(ctx context.Context) NotebookPreparationErrorResponseOutput {
	return o
}

func (o NotebookPreparationErrorResponseOutput) ToNotebookPreparationErrorResponsePtrOutput() NotebookPreparationErrorResponsePtrOutput {
	return o.ToNotebookPreparationErrorResponsePtrOutputWithContext(context.Background())
}

func (o NotebookPreparationErrorResponseOutput) ToNotebookPreparationErrorResponsePtrOutputWithContext(ctx context.Context) NotebookPreparationErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotebookPreparationErrorResponse) *NotebookPreparationErrorResponse {
		return &v
	}).(NotebookPreparationErrorResponsePtrOutput)
}

func (o NotebookPreparationErrorResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookPreparationErrorResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o NotebookPreparationErrorResponseOutput) StatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NotebookPreparationErrorResponse) *int { return v.StatusCode }).(pulumi.IntPtrOutput)
}

type NotebookPreparationErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (NotebookPreparationErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookPreparationErrorResponse)(nil)).Elem()
}

func (o NotebookPreparationErrorResponsePtrOutput) ToNotebookPreparationErrorResponsePtrOutput() NotebookPreparationErrorResponsePtrOutput {
	return o
}

func (o NotebookPreparationErrorResponsePtrOutput) ToNotebookPreparationErrorResponsePtrOutputWithContext(ctx context.Context) NotebookPreparationErrorResponsePtrOutput {
	return o
}

func (o NotebookPreparationErrorResponsePtrOutput) Elem() NotebookPreparationErrorResponseOutput {
	return o.ApplyT(func(v *NotebookPreparationErrorResponse) NotebookPreparationErrorResponse {
		if v != nil {
			return *v
		}
		var ret NotebookPreparationErrorResponse
		return ret
	}).(NotebookPreparationErrorResponseOutput)
}

func (o NotebookPreparationErrorResponsePtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookPreparationErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o NotebookPreparationErrorResponsePtrOutput) StatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NotebookPreparationErrorResponse) *int {
		if v == nil {
			return nil
		}
		return v.StatusCode
	}).(pulumi.IntPtrOutput)
}

type NotebookResourceInfoResponse struct {
	Fqdn                     *string                           `pulumi:"fqdn"`
	NotebookPreparationError *NotebookPreparationErrorResponse `pulumi:"notebookPreparationError"`
	ResourceId               *string                           `pulumi:"resourceId"`
}





type NotebookResourceInfoResponseInput interface {
	pulumi.Input

	ToNotebookResourceInfoResponseOutput() NotebookResourceInfoResponseOutput
	ToNotebookResourceInfoResponseOutputWithContext(context.Context) NotebookResourceInfoResponseOutput
}

type NotebookResourceInfoResponseArgs struct {
	Fqdn                     pulumi.StringPtrInput                    `pulumi:"fqdn"`
	NotebookPreparationError NotebookPreparationErrorResponsePtrInput `pulumi:"notebookPreparationError"`
	ResourceId               pulumi.StringPtrInput                    `pulumi:"resourceId"`
}

func (NotebookResourceInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookResourceInfoResponse)(nil)).Elem()
}

func (i NotebookResourceInfoResponseArgs) ToNotebookResourceInfoResponseOutput() NotebookResourceInfoResponseOutput {
	return i.ToNotebookResourceInfoResponseOutputWithContext(context.Background())
}

func (i NotebookResourceInfoResponseArgs) ToNotebookResourceInfoResponseOutputWithContext(ctx context.Context) NotebookResourceInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceInfoResponseOutput)
}

func (i NotebookResourceInfoResponseArgs) ToNotebookResourceInfoResponsePtrOutput() NotebookResourceInfoResponsePtrOutput {
	return i.ToNotebookResourceInfoResponsePtrOutputWithContext(context.Background())
}

func (i NotebookResourceInfoResponseArgs) ToNotebookResourceInfoResponsePtrOutputWithContext(ctx context.Context) NotebookResourceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceInfoResponseOutput).ToNotebookResourceInfoResponsePtrOutputWithContext(ctx)
}









type NotebookResourceInfoResponsePtrInput interface {
	pulumi.Input

	ToNotebookResourceInfoResponsePtrOutput() NotebookResourceInfoResponsePtrOutput
	ToNotebookResourceInfoResponsePtrOutputWithContext(context.Context) NotebookResourceInfoResponsePtrOutput
}

type notebookResourceInfoResponsePtrType NotebookResourceInfoResponseArgs

func NotebookResourceInfoResponsePtr(v *NotebookResourceInfoResponseArgs) NotebookResourceInfoResponsePtrInput {
	return (*notebookResourceInfoResponsePtrType)(v)
}

func (*notebookResourceInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookResourceInfoResponse)(nil)).Elem()
}

func (i *notebookResourceInfoResponsePtrType) ToNotebookResourceInfoResponsePtrOutput() NotebookResourceInfoResponsePtrOutput {
	return i.ToNotebookResourceInfoResponsePtrOutputWithContext(context.Background())
}

func (i *notebookResourceInfoResponsePtrType) ToNotebookResourceInfoResponsePtrOutputWithContext(ctx context.Context) NotebookResourceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookResourceInfoResponsePtrOutput)
}

type NotebookResourceInfoResponseOutput struct{ *pulumi.OutputState }

func (NotebookResourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookResourceInfoResponse)(nil)).Elem()
}

func (o NotebookResourceInfoResponseOutput) ToNotebookResourceInfoResponseOutput() NotebookResourceInfoResponseOutput {
	return o
}

func (o NotebookResourceInfoResponseOutput) ToNotebookResourceInfoResponseOutputWithContext(ctx context.Context) NotebookResourceInfoResponseOutput {
	return o
}

func (o NotebookResourceInfoResponseOutput) ToNotebookResourceInfoResponsePtrOutput() NotebookResourceInfoResponsePtrOutput {
	return o.ToNotebookResourceInfoResponsePtrOutputWithContext(context.Background())
}

func (o NotebookResourceInfoResponseOutput) ToNotebookResourceInfoResponsePtrOutputWithContext(ctx context.Context) NotebookResourceInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotebookResourceInfoResponse) *NotebookResourceInfoResponse {
		return &v
	}).(NotebookResourceInfoResponsePtrOutput)
}

func (o NotebookResourceInfoResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceInfoResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o NotebookResourceInfoResponseOutput) NotebookPreparationError() NotebookPreparationErrorResponsePtrOutput {
	return o.ApplyT(func(v NotebookResourceInfoResponse) *NotebookPreparationErrorResponse {
		return v.NotebookPreparationError
	}).(NotebookPreparationErrorResponsePtrOutput)
}

func (o NotebookResourceInfoResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotebookResourceInfoResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type NotebookResourceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (NotebookResourceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookResourceInfoResponse)(nil)).Elem()
}

func (o NotebookResourceInfoResponsePtrOutput) ToNotebookResourceInfoResponsePtrOutput() NotebookResourceInfoResponsePtrOutput {
	return o
}

func (o NotebookResourceInfoResponsePtrOutput) ToNotebookResourceInfoResponsePtrOutputWithContext(ctx context.Context) NotebookResourceInfoResponsePtrOutput {
	return o
}

func (o NotebookResourceInfoResponsePtrOutput) Elem() NotebookResourceInfoResponseOutput {
	return o.ApplyT(func(v *NotebookResourceInfoResponse) NotebookResourceInfoResponse {
		if v != nil {
			return *v
		}
		var ret NotebookResourceInfoResponse
		return ret
	}).(NotebookResourceInfoResponseOutput)
}

func (o NotebookResourceInfoResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o NotebookResourceInfoResponsePtrOutput) NotebookPreparationError() NotebookPreparationErrorResponsePtrOutput {
	return o.ApplyT(func(v *NotebookResourceInfoResponse) *NotebookPreparationErrorResponse {
		if v == nil {
			return nil
		}
		return v.NotebookPreparationError
	}).(NotebookPreparationErrorResponsePtrOutput)
}

func (o NotebookResourceInfoResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotebookResourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
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

type PersonalComputeInstanceSettings struct {
	AssignedUser *AssignedUser `pulumi:"assignedUser"`
}





type PersonalComputeInstanceSettingsInput interface {
	pulumi.Input

	ToPersonalComputeInstanceSettingsOutput() PersonalComputeInstanceSettingsOutput
	ToPersonalComputeInstanceSettingsOutputWithContext(context.Context) PersonalComputeInstanceSettingsOutput
}

type PersonalComputeInstanceSettingsArgs struct {
	AssignedUser AssignedUserPtrInput `pulumi:"assignedUser"`
}

func (PersonalComputeInstanceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PersonalComputeInstanceSettings)(nil)).Elem()
}

func (i PersonalComputeInstanceSettingsArgs) ToPersonalComputeInstanceSettingsOutput() PersonalComputeInstanceSettingsOutput {
	return i.ToPersonalComputeInstanceSettingsOutputWithContext(context.Background())
}

func (i PersonalComputeInstanceSettingsArgs) ToPersonalComputeInstanceSettingsOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalComputeInstanceSettingsOutput)
}

func (i PersonalComputeInstanceSettingsArgs) ToPersonalComputeInstanceSettingsPtrOutput() PersonalComputeInstanceSettingsPtrOutput {
	return i.ToPersonalComputeInstanceSettingsPtrOutputWithContext(context.Background())
}

func (i PersonalComputeInstanceSettingsArgs) ToPersonalComputeInstanceSettingsPtrOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalComputeInstanceSettingsOutput).ToPersonalComputeInstanceSettingsPtrOutputWithContext(ctx)
}









type PersonalComputeInstanceSettingsPtrInput interface {
	pulumi.Input

	ToPersonalComputeInstanceSettingsPtrOutput() PersonalComputeInstanceSettingsPtrOutput
	ToPersonalComputeInstanceSettingsPtrOutputWithContext(context.Context) PersonalComputeInstanceSettingsPtrOutput
}

type personalComputeInstanceSettingsPtrType PersonalComputeInstanceSettingsArgs

func PersonalComputeInstanceSettingsPtr(v *PersonalComputeInstanceSettingsArgs) PersonalComputeInstanceSettingsPtrInput {
	return (*personalComputeInstanceSettingsPtrType)(v)
}

func (*personalComputeInstanceSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalComputeInstanceSettings)(nil)).Elem()
}

func (i *personalComputeInstanceSettingsPtrType) ToPersonalComputeInstanceSettingsPtrOutput() PersonalComputeInstanceSettingsPtrOutput {
	return i.ToPersonalComputeInstanceSettingsPtrOutputWithContext(context.Background())
}

func (i *personalComputeInstanceSettingsPtrType) ToPersonalComputeInstanceSettingsPtrOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalComputeInstanceSettingsPtrOutput)
}

type PersonalComputeInstanceSettingsOutput struct{ *pulumi.OutputState }

func (PersonalComputeInstanceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersonalComputeInstanceSettings)(nil)).Elem()
}

func (o PersonalComputeInstanceSettingsOutput) ToPersonalComputeInstanceSettingsOutput() PersonalComputeInstanceSettingsOutput {
	return o
}

func (o PersonalComputeInstanceSettingsOutput) ToPersonalComputeInstanceSettingsOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsOutput {
	return o
}

func (o PersonalComputeInstanceSettingsOutput) ToPersonalComputeInstanceSettingsPtrOutput() PersonalComputeInstanceSettingsPtrOutput {
	return o.ToPersonalComputeInstanceSettingsPtrOutputWithContext(context.Background())
}

func (o PersonalComputeInstanceSettingsOutput) ToPersonalComputeInstanceSettingsPtrOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersonalComputeInstanceSettings) *PersonalComputeInstanceSettings {
		return &v
	}).(PersonalComputeInstanceSettingsPtrOutput)
}

func (o PersonalComputeInstanceSettingsOutput) AssignedUser() AssignedUserPtrOutput {
	return o.ApplyT(func(v PersonalComputeInstanceSettings) *AssignedUser { return v.AssignedUser }).(AssignedUserPtrOutput)
}

type PersonalComputeInstanceSettingsPtrOutput struct{ *pulumi.OutputState }

func (PersonalComputeInstanceSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalComputeInstanceSettings)(nil)).Elem()
}

func (o PersonalComputeInstanceSettingsPtrOutput) ToPersonalComputeInstanceSettingsPtrOutput() PersonalComputeInstanceSettingsPtrOutput {
	return o
}

func (o PersonalComputeInstanceSettingsPtrOutput) ToPersonalComputeInstanceSettingsPtrOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsPtrOutput {
	return o
}

func (o PersonalComputeInstanceSettingsPtrOutput) Elem() PersonalComputeInstanceSettingsOutput {
	return o.ApplyT(func(v *PersonalComputeInstanceSettings) PersonalComputeInstanceSettings {
		if v != nil {
			return *v
		}
		var ret PersonalComputeInstanceSettings
		return ret
	}).(PersonalComputeInstanceSettingsOutput)
}

func (o PersonalComputeInstanceSettingsPtrOutput) AssignedUser() AssignedUserPtrOutput {
	return o.ApplyT(func(v *PersonalComputeInstanceSettings) *AssignedUser {
		if v == nil {
			return nil
		}
		return v.AssignedUser
	}).(AssignedUserPtrOutput)
}

type PersonalComputeInstanceSettingsResponse struct {
	AssignedUser *AssignedUserResponse `pulumi:"assignedUser"`
}





type PersonalComputeInstanceSettingsResponseInput interface {
	pulumi.Input

	ToPersonalComputeInstanceSettingsResponseOutput() PersonalComputeInstanceSettingsResponseOutput
	ToPersonalComputeInstanceSettingsResponseOutputWithContext(context.Context) PersonalComputeInstanceSettingsResponseOutput
}

type PersonalComputeInstanceSettingsResponseArgs struct {
	AssignedUser AssignedUserResponsePtrInput `pulumi:"assignedUser"`
}

func (PersonalComputeInstanceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PersonalComputeInstanceSettingsResponse)(nil)).Elem()
}

func (i PersonalComputeInstanceSettingsResponseArgs) ToPersonalComputeInstanceSettingsResponseOutput() PersonalComputeInstanceSettingsResponseOutput {
	return i.ToPersonalComputeInstanceSettingsResponseOutputWithContext(context.Background())
}

func (i PersonalComputeInstanceSettingsResponseArgs) ToPersonalComputeInstanceSettingsResponseOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalComputeInstanceSettingsResponseOutput)
}

func (i PersonalComputeInstanceSettingsResponseArgs) ToPersonalComputeInstanceSettingsResponsePtrOutput() PersonalComputeInstanceSettingsResponsePtrOutput {
	return i.ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(context.Background())
}

func (i PersonalComputeInstanceSettingsResponseArgs) ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalComputeInstanceSettingsResponseOutput).ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(ctx)
}









type PersonalComputeInstanceSettingsResponsePtrInput interface {
	pulumi.Input

	ToPersonalComputeInstanceSettingsResponsePtrOutput() PersonalComputeInstanceSettingsResponsePtrOutput
	ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(context.Context) PersonalComputeInstanceSettingsResponsePtrOutput
}

type personalComputeInstanceSettingsResponsePtrType PersonalComputeInstanceSettingsResponseArgs

func PersonalComputeInstanceSettingsResponsePtr(v *PersonalComputeInstanceSettingsResponseArgs) PersonalComputeInstanceSettingsResponsePtrInput {
	return (*personalComputeInstanceSettingsResponsePtrType)(v)
}

func (*personalComputeInstanceSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalComputeInstanceSettingsResponse)(nil)).Elem()
}

func (i *personalComputeInstanceSettingsResponsePtrType) ToPersonalComputeInstanceSettingsResponsePtrOutput() PersonalComputeInstanceSettingsResponsePtrOutput {
	return i.ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *personalComputeInstanceSettingsResponsePtrType) ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersonalComputeInstanceSettingsResponsePtrOutput)
}

type PersonalComputeInstanceSettingsResponseOutput struct{ *pulumi.OutputState }

func (PersonalComputeInstanceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersonalComputeInstanceSettingsResponse)(nil)).Elem()
}

func (o PersonalComputeInstanceSettingsResponseOutput) ToPersonalComputeInstanceSettingsResponseOutput() PersonalComputeInstanceSettingsResponseOutput {
	return o
}

func (o PersonalComputeInstanceSettingsResponseOutput) ToPersonalComputeInstanceSettingsResponseOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsResponseOutput {
	return o
}

func (o PersonalComputeInstanceSettingsResponseOutput) ToPersonalComputeInstanceSettingsResponsePtrOutput() PersonalComputeInstanceSettingsResponsePtrOutput {
	return o.ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(context.Background())
}

func (o PersonalComputeInstanceSettingsResponseOutput) ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersonalComputeInstanceSettingsResponse) *PersonalComputeInstanceSettingsResponse {
		return &v
	}).(PersonalComputeInstanceSettingsResponsePtrOutput)
}

func (o PersonalComputeInstanceSettingsResponseOutput) AssignedUser() AssignedUserResponsePtrOutput {
	return o.ApplyT(func(v PersonalComputeInstanceSettingsResponse) *AssignedUserResponse { return v.AssignedUser }).(AssignedUserResponsePtrOutput)
}

type PersonalComputeInstanceSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PersonalComputeInstanceSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalComputeInstanceSettingsResponse)(nil)).Elem()
}

func (o PersonalComputeInstanceSettingsResponsePtrOutput) ToPersonalComputeInstanceSettingsResponsePtrOutput() PersonalComputeInstanceSettingsResponsePtrOutput {
	return o
}

func (o PersonalComputeInstanceSettingsResponsePtrOutput) ToPersonalComputeInstanceSettingsResponsePtrOutputWithContext(ctx context.Context) PersonalComputeInstanceSettingsResponsePtrOutput {
	return o
}

func (o PersonalComputeInstanceSettingsResponsePtrOutput) Elem() PersonalComputeInstanceSettingsResponseOutput {
	return o.ApplyT(func(v *PersonalComputeInstanceSettingsResponse) PersonalComputeInstanceSettingsResponse {
		if v != nil {
			return *v
		}
		var ret PersonalComputeInstanceSettingsResponse
		return ret
	}).(PersonalComputeInstanceSettingsResponseOutput)
}

func (o PersonalComputeInstanceSettingsResponsePtrOutput) AssignedUser() AssignedUserResponsePtrOutput {
	return o.ApplyT(func(v *PersonalComputeInstanceSettingsResponse) *AssignedUserResponse {
		if v == nil {
			return nil
		}
		return v.AssignedUser
	}).(AssignedUserResponsePtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Identity                          *IdentityResponse                         `pulumi:"identity"`
	Location                          *string                                   `pulumi:"location"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Sku                               *SkuResponse                              `pulumi:"sku"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Tags                              map[string]string                         `pulumi:"tags"`
	Type                              string                                    `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                             `pulumi:"id"`
	Identity                          IdentityResponsePtrInput                       `pulumi:"identity"`
	Location                          pulumi.StringPtrInput                          `pulumi:"location"`
	Name                              pulumi.StringInput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
	Sku                               SkuResponsePtrInput                            `pulumi:"sku"`
	SystemData                        SystemDataResponseInput                        `pulumi:"systemData"`
	Tags                              pulumi.StringMapInput                          `pulumi:"tags"`
	Type                              pulumi.StringInput                             `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id          string `pulumi:"id"`
	SubnetArmId string `pulumi:"subnetArmId"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id          pulumi.StringInput `pulumi:"id"`
	SubnetArmId pulumi.StringInput `pulumi:"subnetArmId"`
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

func (o PrivateEndpointResponseOutput) SubnetArmId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.SubnetArmId }).(pulumi.StringOutput)
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

func (o PrivateEndpointResponsePtrOutput) SubnetArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SubnetArmId
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
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
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
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

type ScriptReference struct {
	ScriptArguments *string `pulumi:"scriptArguments"`
	ScriptData      *string `pulumi:"scriptData"`
	ScriptSource    *string `pulumi:"scriptSource"`
	Timeout         *string `pulumi:"timeout"`
}





type ScriptReferenceInput interface {
	pulumi.Input

	ToScriptReferenceOutput() ScriptReferenceOutput
	ToScriptReferenceOutputWithContext(context.Context) ScriptReferenceOutput
}

type ScriptReferenceArgs struct {
	ScriptArguments pulumi.StringPtrInput `pulumi:"scriptArguments"`
	ScriptData      pulumi.StringPtrInput `pulumi:"scriptData"`
	ScriptSource    pulumi.StringPtrInput `pulumi:"scriptSource"`
	Timeout         pulumi.StringPtrInput `pulumi:"timeout"`
}

func (ScriptReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptReference)(nil)).Elem()
}

func (i ScriptReferenceArgs) ToScriptReferenceOutput() ScriptReferenceOutput {
	return i.ToScriptReferenceOutputWithContext(context.Background())
}

func (i ScriptReferenceArgs) ToScriptReferenceOutputWithContext(ctx context.Context) ScriptReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptReferenceOutput)
}

func (i ScriptReferenceArgs) ToScriptReferencePtrOutput() ScriptReferencePtrOutput {
	return i.ToScriptReferencePtrOutputWithContext(context.Background())
}

func (i ScriptReferenceArgs) ToScriptReferencePtrOutputWithContext(ctx context.Context) ScriptReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptReferenceOutput).ToScriptReferencePtrOutputWithContext(ctx)
}









type ScriptReferencePtrInput interface {
	pulumi.Input

	ToScriptReferencePtrOutput() ScriptReferencePtrOutput
	ToScriptReferencePtrOutputWithContext(context.Context) ScriptReferencePtrOutput
}

type scriptReferencePtrType ScriptReferenceArgs

func ScriptReferencePtr(v *ScriptReferenceArgs) ScriptReferencePtrInput {
	return (*scriptReferencePtrType)(v)
}

func (*scriptReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptReference)(nil)).Elem()
}

func (i *scriptReferencePtrType) ToScriptReferencePtrOutput() ScriptReferencePtrOutput {
	return i.ToScriptReferencePtrOutputWithContext(context.Background())
}

func (i *scriptReferencePtrType) ToScriptReferencePtrOutputWithContext(ctx context.Context) ScriptReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptReferencePtrOutput)
}

type ScriptReferenceOutput struct{ *pulumi.OutputState }

func (ScriptReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptReference)(nil)).Elem()
}

func (o ScriptReferenceOutput) ToScriptReferenceOutput() ScriptReferenceOutput {
	return o
}

func (o ScriptReferenceOutput) ToScriptReferenceOutputWithContext(ctx context.Context) ScriptReferenceOutput {
	return o
}

func (o ScriptReferenceOutput) ToScriptReferencePtrOutput() ScriptReferencePtrOutput {
	return o.ToScriptReferencePtrOutputWithContext(context.Background())
}

func (o ScriptReferenceOutput) ToScriptReferencePtrOutputWithContext(ctx context.Context) ScriptReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScriptReference) *ScriptReference {
		return &v
	}).(ScriptReferencePtrOutput)
}

func (o ScriptReferenceOutput) ScriptArguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptReference) *string { return v.ScriptArguments }).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceOutput) ScriptData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptReference) *string { return v.ScriptData }).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceOutput) ScriptSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptReference) *string { return v.ScriptSource }).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptReference) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

type ScriptReferencePtrOutput struct{ *pulumi.OutputState }

func (ScriptReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptReference)(nil)).Elem()
}

func (o ScriptReferencePtrOutput) ToScriptReferencePtrOutput() ScriptReferencePtrOutput {
	return o
}

func (o ScriptReferencePtrOutput) ToScriptReferencePtrOutputWithContext(ctx context.Context) ScriptReferencePtrOutput {
	return o
}

func (o ScriptReferencePtrOutput) Elem() ScriptReferenceOutput {
	return o.ApplyT(func(v *ScriptReference) ScriptReference {
		if v != nil {
			return *v
		}
		var ret ScriptReference
		return ret
	}).(ScriptReferenceOutput)
}

func (o ScriptReferencePtrOutput) ScriptArguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptReference) *string {
		if v == nil {
			return nil
		}
		return v.ScriptArguments
	}).(pulumi.StringPtrOutput)
}

func (o ScriptReferencePtrOutput) ScriptData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptReference) *string {
		if v == nil {
			return nil
		}
		return v.ScriptData
	}).(pulumi.StringPtrOutput)
}

func (o ScriptReferencePtrOutput) ScriptSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptReference) *string {
		if v == nil {
			return nil
		}
		return v.ScriptSource
	}).(pulumi.StringPtrOutput)
}

func (o ScriptReferencePtrOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptReference) *string {
		if v == nil {
			return nil
		}
		return v.Timeout
	}).(pulumi.StringPtrOutput)
}

type ScriptReferenceResponse struct {
	ScriptArguments *string `pulumi:"scriptArguments"`
	ScriptData      *string `pulumi:"scriptData"`
	ScriptSource    *string `pulumi:"scriptSource"`
	Timeout         *string `pulumi:"timeout"`
}





type ScriptReferenceResponseInput interface {
	pulumi.Input

	ToScriptReferenceResponseOutput() ScriptReferenceResponseOutput
	ToScriptReferenceResponseOutputWithContext(context.Context) ScriptReferenceResponseOutput
}

type ScriptReferenceResponseArgs struct {
	ScriptArguments pulumi.StringPtrInput `pulumi:"scriptArguments"`
	ScriptData      pulumi.StringPtrInput `pulumi:"scriptData"`
	ScriptSource    pulumi.StringPtrInput `pulumi:"scriptSource"`
	Timeout         pulumi.StringPtrInput `pulumi:"timeout"`
}

func (ScriptReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptReferenceResponse)(nil)).Elem()
}

func (i ScriptReferenceResponseArgs) ToScriptReferenceResponseOutput() ScriptReferenceResponseOutput {
	return i.ToScriptReferenceResponseOutputWithContext(context.Background())
}

func (i ScriptReferenceResponseArgs) ToScriptReferenceResponseOutputWithContext(ctx context.Context) ScriptReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptReferenceResponseOutput)
}

func (i ScriptReferenceResponseArgs) ToScriptReferenceResponsePtrOutput() ScriptReferenceResponsePtrOutput {
	return i.ToScriptReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ScriptReferenceResponseArgs) ToScriptReferenceResponsePtrOutputWithContext(ctx context.Context) ScriptReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptReferenceResponseOutput).ToScriptReferenceResponsePtrOutputWithContext(ctx)
}









type ScriptReferenceResponsePtrInput interface {
	pulumi.Input

	ToScriptReferenceResponsePtrOutput() ScriptReferenceResponsePtrOutput
	ToScriptReferenceResponsePtrOutputWithContext(context.Context) ScriptReferenceResponsePtrOutput
}

type scriptReferenceResponsePtrType ScriptReferenceResponseArgs

func ScriptReferenceResponsePtr(v *ScriptReferenceResponseArgs) ScriptReferenceResponsePtrInput {
	return (*scriptReferenceResponsePtrType)(v)
}

func (*scriptReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptReferenceResponse)(nil)).Elem()
}

func (i *scriptReferenceResponsePtrType) ToScriptReferenceResponsePtrOutput() ScriptReferenceResponsePtrOutput {
	return i.ToScriptReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *scriptReferenceResponsePtrType) ToScriptReferenceResponsePtrOutputWithContext(ctx context.Context) ScriptReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptReferenceResponsePtrOutput)
}

type ScriptReferenceResponseOutput struct{ *pulumi.OutputState }

func (ScriptReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptReferenceResponse)(nil)).Elem()
}

func (o ScriptReferenceResponseOutput) ToScriptReferenceResponseOutput() ScriptReferenceResponseOutput {
	return o
}

func (o ScriptReferenceResponseOutput) ToScriptReferenceResponseOutputWithContext(ctx context.Context) ScriptReferenceResponseOutput {
	return o
}

func (o ScriptReferenceResponseOutput) ToScriptReferenceResponsePtrOutput() ScriptReferenceResponsePtrOutput {
	return o.ToScriptReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ScriptReferenceResponseOutput) ToScriptReferenceResponsePtrOutputWithContext(ctx context.Context) ScriptReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScriptReferenceResponse) *ScriptReferenceResponse {
		return &v
	}).(ScriptReferenceResponsePtrOutput)
}

func (o ScriptReferenceResponseOutput) ScriptArguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptReferenceResponse) *string { return v.ScriptArguments }).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceResponseOutput) ScriptData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptReferenceResponse) *string { return v.ScriptData }).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceResponseOutput) ScriptSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptReferenceResponse) *string { return v.ScriptSource }).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceResponseOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScriptReferenceResponse) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

type ScriptReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ScriptReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptReferenceResponse)(nil)).Elem()
}

func (o ScriptReferenceResponsePtrOutput) ToScriptReferenceResponsePtrOutput() ScriptReferenceResponsePtrOutput {
	return o
}

func (o ScriptReferenceResponsePtrOutput) ToScriptReferenceResponsePtrOutputWithContext(ctx context.Context) ScriptReferenceResponsePtrOutput {
	return o
}

func (o ScriptReferenceResponsePtrOutput) Elem() ScriptReferenceResponseOutput {
	return o.ApplyT(func(v *ScriptReferenceResponse) ScriptReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ScriptReferenceResponse
		return ret
	}).(ScriptReferenceResponseOutput)
}

func (o ScriptReferenceResponsePtrOutput) ScriptArguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScriptArguments
	}).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceResponsePtrOutput) ScriptData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScriptData
	}).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceResponsePtrOutput) ScriptSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScriptSource
	}).(pulumi.StringPtrOutput)
}

func (o ScriptReferenceResponsePtrOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Timeout
	}).(pulumi.StringPtrOutput)
}

type ScriptsToExecute struct {
	CreationScript *ScriptReference `pulumi:"creationScript"`
	StartupScript  *ScriptReference `pulumi:"startupScript"`
}





type ScriptsToExecuteInput interface {
	pulumi.Input

	ToScriptsToExecuteOutput() ScriptsToExecuteOutput
	ToScriptsToExecuteOutputWithContext(context.Context) ScriptsToExecuteOutput
}

type ScriptsToExecuteArgs struct {
	CreationScript ScriptReferencePtrInput `pulumi:"creationScript"`
	StartupScript  ScriptReferencePtrInput `pulumi:"startupScript"`
}

func (ScriptsToExecuteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptsToExecute)(nil)).Elem()
}

func (i ScriptsToExecuteArgs) ToScriptsToExecuteOutput() ScriptsToExecuteOutput {
	return i.ToScriptsToExecuteOutputWithContext(context.Background())
}

func (i ScriptsToExecuteArgs) ToScriptsToExecuteOutputWithContext(ctx context.Context) ScriptsToExecuteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptsToExecuteOutput)
}

func (i ScriptsToExecuteArgs) ToScriptsToExecutePtrOutput() ScriptsToExecutePtrOutput {
	return i.ToScriptsToExecutePtrOutputWithContext(context.Background())
}

func (i ScriptsToExecuteArgs) ToScriptsToExecutePtrOutputWithContext(ctx context.Context) ScriptsToExecutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptsToExecuteOutput).ToScriptsToExecutePtrOutputWithContext(ctx)
}









type ScriptsToExecutePtrInput interface {
	pulumi.Input

	ToScriptsToExecutePtrOutput() ScriptsToExecutePtrOutput
	ToScriptsToExecutePtrOutputWithContext(context.Context) ScriptsToExecutePtrOutput
}

type scriptsToExecutePtrType ScriptsToExecuteArgs

func ScriptsToExecutePtr(v *ScriptsToExecuteArgs) ScriptsToExecutePtrInput {
	return (*scriptsToExecutePtrType)(v)
}

func (*scriptsToExecutePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptsToExecute)(nil)).Elem()
}

func (i *scriptsToExecutePtrType) ToScriptsToExecutePtrOutput() ScriptsToExecutePtrOutput {
	return i.ToScriptsToExecutePtrOutputWithContext(context.Background())
}

func (i *scriptsToExecutePtrType) ToScriptsToExecutePtrOutputWithContext(ctx context.Context) ScriptsToExecutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptsToExecutePtrOutput)
}

type ScriptsToExecuteOutput struct{ *pulumi.OutputState }

func (ScriptsToExecuteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptsToExecute)(nil)).Elem()
}

func (o ScriptsToExecuteOutput) ToScriptsToExecuteOutput() ScriptsToExecuteOutput {
	return o
}

func (o ScriptsToExecuteOutput) ToScriptsToExecuteOutputWithContext(ctx context.Context) ScriptsToExecuteOutput {
	return o
}

func (o ScriptsToExecuteOutput) ToScriptsToExecutePtrOutput() ScriptsToExecutePtrOutput {
	return o.ToScriptsToExecutePtrOutputWithContext(context.Background())
}

func (o ScriptsToExecuteOutput) ToScriptsToExecutePtrOutputWithContext(ctx context.Context) ScriptsToExecutePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScriptsToExecute) *ScriptsToExecute {
		return &v
	}).(ScriptsToExecutePtrOutput)
}

func (o ScriptsToExecuteOutput) CreationScript() ScriptReferencePtrOutput {
	return o.ApplyT(func(v ScriptsToExecute) *ScriptReference { return v.CreationScript }).(ScriptReferencePtrOutput)
}

func (o ScriptsToExecuteOutput) StartupScript() ScriptReferencePtrOutput {
	return o.ApplyT(func(v ScriptsToExecute) *ScriptReference { return v.StartupScript }).(ScriptReferencePtrOutput)
}

type ScriptsToExecutePtrOutput struct{ *pulumi.OutputState }

func (ScriptsToExecutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptsToExecute)(nil)).Elem()
}

func (o ScriptsToExecutePtrOutput) ToScriptsToExecutePtrOutput() ScriptsToExecutePtrOutput {
	return o
}

func (o ScriptsToExecutePtrOutput) ToScriptsToExecutePtrOutputWithContext(ctx context.Context) ScriptsToExecutePtrOutput {
	return o
}

func (o ScriptsToExecutePtrOutput) Elem() ScriptsToExecuteOutput {
	return o.ApplyT(func(v *ScriptsToExecute) ScriptsToExecute {
		if v != nil {
			return *v
		}
		var ret ScriptsToExecute
		return ret
	}).(ScriptsToExecuteOutput)
}

func (o ScriptsToExecutePtrOutput) CreationScript() ScriptReferencePtrOutput {
	return o.ApplyT(func(v *ScriptsToExecute) *ScriptReference {
		if v == nil {
			return nil
		}
		return v.CreationScript
	}).(ScriptReferencePtrOutput)
}

func (o ScriptsToExecutePtrOutput) StartupScript() ScriptReferencePtrOutput {
	return o.ApplyT(func(v *ScriptsToExecute) *ScriptReference {
		if v == nil {
			return nil
		}
		return v.StartupScript
	}).(ScriptReferencePtrOutput)
}

type ScriptsToExecuteResponse struct {
	CreationScript *ScriptReferenceResponse `pulumi:"creationScript"`
	StartupScript  *ScriptReferenceResponse `pulumi:"startupScript"`
}





type ScriptsToExecuteResponseInput interface {
	pulumi.Input

	ToScriptsToExecuteResponseOutput() ScriptsToExecuteResponseOutput
	ToScriptsToExecuteResponseOutputWithContext(context.Context) ScriptsToExecuteResponseOutput
}

type ScriptsToExecuteResponseArgs struct {
	CreationScript ScriptReferenceResponsePtrInput `pulumi:"creationScript"`
	StartupScript  ScriptReferenceResponsePtrInput `pulumi:"startupScript"`
}

func (ScriptsToExecuteResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptsToExecuteResponse)(nil)).Elem()
}

func (i ScriptsToExecuteResponseArgs) ToScriptsToExecuteResponseOutput() ScriptsToExecuteResponseOutput {
	return i.ToScriptsToExecuteResponseOutputWithContext(context.Background())
}

func (i ScriptsToExecuteResponseArgs) ToScriptsToExecuteResponseOutputWithContext(ctx context.Context) ScriptsToExecuteResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptsToExecuteResponseOutput)
}

func (i ScriptsToExecuteResponseArgs) ToScriptsToExecuteResponsePtrOutput() ScriptsToExecuteResponsePtrOutput {
	return i.ToScriptsToExecuteResponsePtrOutputWithContext(context.Background())
}

func (i ScriptsToExecuteResponseArgs) ToScriptsToExecuteResponsePtrOutputWithContext(ctx context.Context) ScriptsToExecuteResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptsToExecuteResponseOutput).ToScriptsToExecuteResponsePtrOutputWithContext(ctx)
}









type ScriptsToExecuteResponsePtrInput interface {
	pulumi.Input

	ToScriptsToExecuteResponsePtrOutput() ScriptsToExecuteResponsePtrOutput
	ToScriptsToExecuteResponsePtrOutputWithContext(context.Context) ScriptsToExecuteResponsePtrOutput
}

type scriptsToExecuteResponsePtrType ScriptsToExecuteResponseArgs

func ScriptsToExecuteResponsePtr(v *ScriptsToExecuteResponseArgs) ScriptsToExecuteResponsePtrInput {
	return (*scriptsToExecuteResponsePtrType)(v)
}

func (*scriptsToExecuteResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptsToExecuteResponse)(nil)).Elem()
}

func (i *scriptsToExecuteResponsePtrType) ToScriptsToExecuteResponsePtrOutput() ScriptsToExecuteResponsePtrOutput {
	return i.ToScriptsToExecuteResponsePtrOutputWithContext(context.Background())
}

func (i *scriptsToExecuteResponsePtrType) ToScriptsToExecuteResponsePtrOutputWithContext(ctx context.Context) ScriptsToExecuteResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptsToExecuteResponsePtrOutput)
}

type ScriptsToExecuteResponseOutput struct{ *pulumi.OutputState }

func (ScriptsToExecuteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptsToExecuteResponse)(nil)).Elem()
}

func (o ScriptsToExecuteResponseOutput) ToScriptsToExecuteResponseOutput() ScriptsToExecuteResponseOutput {
	return o
}

func (o ScriptsToExecuteResponseOutput) ToScriptsToExecuteResponseOutputWithContext(ctx context.Context) ScriptsToExecuteResponseOutput {
	return o
}

func (o ScriptsToExecuteResponseOutput) ToScriptsToExecuteResponsePtrOutput() ScriptsToExecuteResponsePtrOutput {
	return o.ToScriptsToExecuteResponsePtrOutputWithContext(context.Background())
}

func (o ScriptsToExecuteResponseOutput) ToScriptsToExecuteResponsePtrOutputWithContext(ctx context.Context) ScriptsToExecuteResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScriptsToExecuteResponse) *ScriptsToExecuteResponse {
		return &v
	}).(ScriptsToExecuteResponsePtrOutput)
}

func (o ScriptsToExecuteResponseOutput) CreationScript() ScriptReferenceResponsePtrOutput {
	return o.ApplyT(func(v ScriptsToExecuteResponse) *ScriptReferenceResponse { return v.CreationScript }).(ScriptReferenceResponsePtrOutput)
}

func (o ScriptsToExecuteResponseOutput) StartupScript() ScriptReferenceResponsePtrOutput {
	return o.ApplyT(func(v ScriptsToExecuteResponse) *ScriptReferenceResponse { return v.StartupScript }).(ScriptReferenceResponsePtrOutput)
}

type ScriptsToExecuteResponsePtrOutput struct{ *pulumi.OutputState }

func (ScriptsToExecuteResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptsToExecuteResponse)(nil)).Elem()
}

func (o ScriptsToExecuteResponsePtrOutput) ToScriptsToExecuteResponsePtrOutput() ScriptsToExecuteResponsePtrOutput {
	return o
}

func (o ScriptsToExecuteResponsePtrOutput) ToScriptsToExecuteResponsePtrOutputWithContext(ctx context.Context) ScriptsToExecuteResponsePtrOutput {
	return o
}

func (o ScriptsToExecuteResponsePtrOutput) Elem() ScriptsToExecuteResponseOutput {
	return o.ApplyT(func(v *ScriptsToExecuteResponse) ScriptsToExecuteResponse {
		if v != nil {
			return *v
		}
		var ret ScriptsToExecuteResponse
		return ret
	}).(ScriptsToExecuteResponseOutput)
}

func (o ScriptsToExecuteResponsePtrOutput) CreationScript() ScriptReferenceResponsePtrOutput {
	return o.ApplyT(func(v *ScriptsToExecuteResponse) *ScriptReferenceResponse {
		if v == nil {
			return nil
		}
		return v.CreationScript
	}).(ScriptReferenceResponsePtrOutput)
}

func (o ScriptsToExecuteResponsePtrOutput) StartupScript() ScriptReferenceResponsePtrOutput {
	return o.ApplyT(func(v *ScriptsToExecuteResponse) *ScriptReferenceResponse {
		if v == nil {
			return nil
		}
		return v.StartupScript
	}).(ScriptReferenceResponsePtrOutput)
}

type ServiceManagedResourcesSettings struct {
	CosmosDb *CosmosDbSettings `pulumi:"cosmosDb"`
}





type ServiceManagedResourcesSettingsInput interface {
	pulumi.Input

	ToServiceManagedResourcesSettingsOutput() ServiceManagedResourcesSettingsOutput
	ToServiceManagedResourcesSettingsOutputWithContext(context.Context) ServiceManagedResourcesSettingsOutput
}

type ServiceManagedResourcesSettingsArgs struct {
	CosmosDb CosmosDbSettingsPtrInput `pulumi:"cosmosDb"`
}

func (ServiceManagedResourcesSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedResourcesSettings)(nil)).Elem()
}

func (i ServiceManagedResourcesSettingsArgs) ToServiceManagedResourcesSettingsOutput() ServiceManagedResourcesSettingsOutput {
	return i.ToServiceManagedResourcesSettingsOutputWithContext(context.Background())
}

func (i ServiceManagedResourcesSettingsArgs) ToServiceManagedResourcesSettingsOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsOutput)
}

func (i ServiceManagedResourcesSettingsArgs) ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput {
	return i.ToServiceManagedResourcesSettingsPtrOutputWithContext(context.Background())
}

func (i ServiceManagedResourcesSettingsArgs) ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsOutput).ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx)
}









type ServiceManagedResourcesSettingsPtrInput interface {
	pulumi.Input

	ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput
	ToServiceManagedResourcesSettingsPtrOutputWithContext(context.Context) ServiceManagedResourcesSettingsPtrOutput
}

type serviceManagedResourcesSettingsPtrType ServiceManagedResourcesSettingsArgs

func ServiceManagedResourcesSettingsPtr(v *ServiceManagedResourcesSettingsArgs) ServiceManagedResourcesSettingsPtrInput {
	return (*serviceManagedResourcesSettingsPtrType)(v)
}

func (*serviceManagedResourcesSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedResourcesSettings)(nil)).Elem()
}

func (i *serviceManagedResourcesSettingsPtrType) ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput {
	return i.ToServiceManagedResourcesSettingsPtrOutputWithContext(context.Background())
}

func (i *serviceManagedResourcesSettingsPtrType) ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsPtrOutput)
}

type ServiceManagedResourcesSettingsOutput struct{ *pulumi.OutputState }

func (ServiceManagedResourcesSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedResourcesSettings)(nil)).Elem()
}

func (o ServiceManagedResourcesSettingsOutput) ToServiceManagedResourcesSettingsOutput() ServiceManagedResourcesSettingsOutput {
	return o
}

func (o ServiceManagedResourcesSettingsOutput) ToServiceManagedResourcesSettingsOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsOutput {
	return o
}

func (o ServiceManagedResourcesSettingsOutput) ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput {
	return o.ToServiceManagedResourcesSettingsPtrOutputWithContext(context.Background())
}

func (o ServiceManagedResourcesSettingsOutput) ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceManagedResourcesSettings) *ServiceManagedResourcesSettings {
		return &v
	}).(ServiceManagedResourcesSettingsPtrOutput)
}

func (o ServiceManagedResourcesSettingsOutput) CosmosDb() CosmosDbSettingsPtrOutput {
	return o.ApplyT(func(v ServiceManagedResourcesSettings) *CosmosDbSettings { return v.CosmosDb }).(CosmosDbSettingsPtrOutput)
}

type ServiceManagedResourcesSettingsPtrOutput struct{ *pulumi.OutputState }

func (ServiceManagedResourcesSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedResourcesSettings)(nil)).Elem()
}

func (o ServiceManagedResourcesSettingsPtrOutput) ToServiceManagedResourcesSettingsPtrOutput() ServiceManagedResourcesSettingsPtrOutput {
	return o
}

func (o ServiceManagedResourcesSettingsPtrOutput) ToServiceManagedResourcesSettingsPtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsPtrOutput {
	return o
}

func (o ServiceManagedResourcesSettingsPtrOutput) Elem() ServiceManagedResourcesSettingsOutput {
	return o.ApplyT(func(v *ServiceManagedResourcesSettings) ServiceManagedResourcesSettings {
		if v != nil {
			return *v
		}
		var ret ServiceManagedResourcesSettings
		return ret
	}).(ServiceManagedResourcesSettingsOutput)
}

func (o ServiceManagedResourcesSettingsPtrOutput) CosmosDb() CosmosDbSettingsPtrOutput {
	return o.ApplyT(func(v *ServiceManagedResourcesSettings) *CosmosDbSettings {
		if v == nil {
			return nil
		}
		return v.CosmosDb
	}).(CosmosDbSettingsPtrOutput)
}

type ServiceManagedResourcesSettingsResponse struct {
	CosmosDb *CosmosDbSettingsResponse `pulumi:"cosmosDb"`
}





type ServiceManagedResourcesSettingsResponseInput interface {
	pulumi.Input

	ToServiceManagedResourcesSettingsResponseOutput() ServiceManagedResourcesSettingsResponseOutput
	ToServiceManagedResourcesSettingsResponseOutputWithContext(context.Context) ServiceManagedResourcesSettingsResponseOutput
}

type ServiceManagedResourcesSettingsResponseArgs struct {
	CosmosDb CosmosDbSettingsResponsePtrInput `pulumi:"cosmosDb"`
}

func (ServiceManagedResourcesSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedResourcesSettingsResponse)(nil)).Elem()
}

func (i ServiceManagedResourcesSettingsResponseArgs) ToServiceManagedResourcesSettingsResponseOutput() ServiceManagedResourcesSettingsResponseOutput {
	return i.ToServiceManagedResourcesSettingsResponseOutputWithContext(context.Background())
}

func (i ServiceManagedResourcesSettingsResponseArgs) ToServiceManagedResourcesSettingsResponseOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsResponseOutput)
}

func (i ServiceManagedResourcesSettingsResponseArgs) ToServiceManagedResourcesSettingsResponsePtrOutput() ServiceManagedResourcesSettingsResponsePtrOutput {
	return i.ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ServiceManagedResourcesSettingsResponseArgs) ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsResponseOutput).ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(ctx)
}









type ServiceManagedResourcesSettingsResponsePtrInput interface {
	pulumi.Input

	ToServiceManagedResourcesSettingsResponsePtrOutput() ServiceManagedResourcesSettingsResponsePtrOutput
	ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(context.Context) ServiceManagedResourcesSettingsResponsePtrOutput
}

type serviceManagedResourcesSettingsResponsePtrType ServiceManagedResourcesSettingsResponseArgs

func ServiceManagedResourcesSettingsResponsePtr(v *ServiceManagedResourcesSettingsResponseArgs) ServiceManagedResourcesSettingsResponsePtrInput {
	return (*serviceManagedResourcesSettingsResponsePtrType)(v)
}

func (*serviceManagedResourcesSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedResourcesSettingsResponse)(nil)).Elem()
}

func (i *serviceManagedResourcesSettingsResponsePtrType) ToServiceManagedResourcesSettingsResponsePtrOutput() ServiceManagedResourcesSettingsResponsePtrOutput {
	return i.ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *serviceManagedResourcesSettingsResponsePtrType) ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceManagedResourcesSettingsResponsePtrOutput)
}

type ServiceManagedResourcesSettingsResponseOutput struct{ *pulumi.OutputState }

func (ServiceManagedResourcesSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedResourcesSettingsResponse)(nil)).Elem()
}

func (o ServiceManagedResourcesSettingsResponseOutput) ToServiceManagedResourcesSettingsResponseOutput() ServiceManagedResourcesSettingsResponseOutput {
	return o
}

func (o ServiceManagedResourcesSettingsResponseOutput) ToServiceManagedResourcesSettingsResponseOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsResponseOutput {
	return o
}

func (o ServiceManagedResourcesSettingsResponseOutput) ToServiceManagedResourcesSettingsResponsePtrOutput() ServiceManagedResourcesSettingsResponsePtrOutput {
	return o.ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ServiceManagedResourcesSettingsResponseOutput) ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceManagedResourcesSettingsResponse) *ServiceManagedResourcesSettingsResponse {
		return &v
	}).(ServiceManagedResourcesSettingsResponsePtrOutput)
}

func (o ServiceManagedResourcesSettingsResponseOutput) CosmosDb() CosmosDbSettingsResponsePtrOutput {
	return o.ApplyT(func(v ServiceManagedResourcesSettingsResponse) *CosmosDbSettingsResponse { return v.CosmosDb }).(CosmosDbSettingsResponsePtrOutput)
}

type ServiceManagedResourcesSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceManagedResourcesSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedResourcesSettingsResponse)(nil)).Elem()
}

func (o ServiceManagedResourcesSettingsResponsePtrOutput) ToServiceManagedResourcesSettingsResponsePtrOutput() ServiceManagedResourcesSettingsResponsePtrOutput {
	return o
}

func (o ServiceManagedResourcesSettingsResponsePtrOutput) ToServiceManagedResourcesSettingsResponsePtrOutputWithContext(ctx context.Context) ServiceManagedResourcesSettingsResponsePtrOutput {
	return o
}

func (o ServiceManagedResourcesSettingsResponsePtrOutput) Elem() ServiceManagedResourcesSettingsResponseOutput {
	return o.ApplyT(func(v *ServiceManagedResourcesSettingsResponse) ServiceManagedResourcesSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ServiceManagedResourcesSettingsResponse
		return ret
	}).(ServiceManagedResourcesSettingsResponseOutput)
}

func (o ServiceManagedResourcesSettingsResponsePtrOutput) CosmosDb() CosmosDbSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServiceManagedResourcesSettingsResponse) *CosmosDbSettingsResponse {
		if v == nil {
			return nil
		}
		return v.CosmosDb
	}).(CosmosDbSettingsResponsePtrOutput)
}

type SetupScripts struct {
	Scripts *ScriptsToExecute `pulumi:"scripts"`
}





type SetupScriptsInput interface {
	pulumi.Input

	ToSetupScriptsOutput() SetupScriptsOutput
	ToSetupScriptsOutputWithContext(context.Context) SetupScriptsOutput
}

type SetupScriptsArgs struct {
	Scripts ScriptsToExecutePtrInput `pulumi:"scripts"`
}

func (SetupScriptsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SetupScripts)(nil)).Elem()
}

func (i SetupScriptsArgs) ToSetupScriptsOutput() SetupScriptsOutput {
	return i.ToSetupScriptsOutputWithContext(context.Background())
}

func (i SetupScriptsArgs) ToSetupScriptsOutputWithContext(ctx context.Context) SetupScriptsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetupScriptsOutput)
}

func (i SetupScriptsArgs) ToSetupScriptsPtrOutput() SetupScriptsPtrOutput {
	return i.ToSetupScriptsPtrOutputWithContext(context.Background())
}

func (i SetupScriptsArgs) ToSetupScriptsPtrOutputWithContext(ctx context.Context) SetupScriptsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetupScriptsOutput).ToSetupScriptsPtrOutputWithContext(ctx)
}









type SetupScriptsPtrInput interface {
	pulumi.Input

	ToSetupScriptsPtrOutput() SetupScriptsPtrOutput
	ToSetupScriptsPtrOutputWithContext(context.Context) SetupScriptsPtrOutput
}

type setupScriptsPtrType SetupScriptsArgs

func SetupScriptsPtr(v *SetupScriptsArgs) SetupScriptsPtrInput {
	return (*setupScriptsPtrType)(v)
}

func (*setupScriptsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SetupScripts)(nil)).Elem()
}

func (i *setupScriptsPtrType) ToSetupScriptsPtrOutput() SetupScriptsPtrOutput {
	return i.ToSetupScriptsPtrOutputWithContext(context.Background())
}

func (i *setupScriptsPtrType) ToSetupScriptsPtrOutputWithContext(ctx context.Context) SetupScriptsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetupScriptsPtrOutput)
}

type SetupScriptsOutput struct{ *pulumi.OutputState }

func (SetupScriptsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SetupScripts)(nil)).Elem()
}

func (o SetupScriptsOutput) ToSetupScriptsOutput() SetupScriptsOutput {
	return o
}

func (o SetupScriptsOutput) ToSetupScriptsOutputWithContext(ctx context.Context) SetupScriptsOutput {
	return o
}

func (o SetupScriptsOutput) ToSetupScriptsPtrOutput() SetupScriptsPtrOutput {
	return o.ToSetupScriptsPtrOutputWithContext(context.Background())
}

func (o SetupScriptsOutput) ToSetupScriptsPtrOutputWithContext(ctx context.Context) SetupScriptsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SetupScripts) *SetupScripts {
		return &v
	}).(SetupScriptsPtrOutput)
}

func (o SetupScriptsOutput) Scripts() ScriptsToExecutePtrOutput {
	return o.ApplyT(func(v SetupScripts) *ScriptsToExecute { return v.Scripts }).(ScriptsToExecutePtrOutput)
}

type SetupScriptsPtrOutput struct{ *pulumi.OutputState }

func (SetupScriptsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SetupScripts)(nil)).Elem()
}

func (o SetupScriptsPtrOutput) ToSetupScriptsPtrOutput() SetupScriptsPtrOutput {
	return o
}

func (o SetupScriptsPtrOutput) ToSetupScriptsPtrOutputWithContext(ctx context.Context) SetupScriptsPtrOutput {
	return o
}

func (o SetupScriptsPtrOutput) Elem() SetupScriptsOutput {
	return o.ApplyT(func(v *SetupScripts) SetupScripts {
		if v != nil {
			return *v
		}
		var ret SetupScripts
		return ret
	}).(SetupScriptsOutput)
}

func (o SetupScriptsPtrOutput) Scripts() ScriptsToExecutePtrOutput {
	return o.ApplyT(func(v *SetupScripts) *ScriptsToExecute {
		if v == nil {
			return nil
		}
		return v.Scripts
	}).(ScriptsToExecutePtrOutput)
}

type SetupScriptsResponse struct {
	Scripts *ScriptsToExecuteResponse `pulumi:"scripts"`
}





type SetupScriptsResponseInput interface {
	pulumi.Input

	ToSetupScriptsResponseOutput() SetupScriptsResponseOutput
	ToSetupScriptsResponseOutputWithContext(context.Context) SetupScriptsResponseOutput
}

type SetupScriptsResponseArgs struct {
	Scripts ScriptsToExecuteResponsePtrInput `pulumi:"scripts"`
}

func (SetupScriptsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SetupScriptsResponse)(nil)).Elem()
}

func (i SetupScriptsResponseArgs) ToSetupScriptsResponseOutput() SetupScriptsResponseOutput {
	return i.ToSetupScriptsResponseOutputWithContext(context.Background())
}

func (i SetupScriptsResponseArgs) ToSetupScriptsResponseOutputWithContext(ctx context.Context) SetupScriptsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetupScriptsResponseOutput)
}

func (i SetupScriptsResponseArgs) ToSetupScriptsResponsePtrOutput() SetupScriptsResponsePtrOutput {
	return i.ToSetupScriptsResponsePtrOutputWithContext(context.Background())
}

func (i SetupScriptsResponseArgs) ToSetupScriptsResponsePtrOutputWithContext(ctx context.Context) SetupScriptsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetupScriptsResponseOutput).ToSetupScriptsResponsePtrOutputWithContext(ctx)
}









type SetupScriptsResponsePtrInput interface {
	pulumi.Input

	ToSetupScriptsResponsePtrOutput() SetupScriptsResponsePtrOutput
	ToSetupScriptsResponsePtrOutputWithContext(context.Context) SetupScriptsResponsePtrOutput
}

type setupScriptsResponsePtrType SetupScriptsResponseArgs

func SetupScriptsResponsePtr(v *SetupScriptsResponseArgs) SetupScriptsResponsePtrInput {
	return (*setupScriptsResponsePtrType)(v)
}

func (*setupScriptsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SetupScriptsResponse)(nil)).Elem()
}

func (i *setupScriptsResponsePtrType) ToSetupScriptsResponsePtrOutput() SetupScriptsResponsePtrOutput {
	return i.ToSetupScriptsResponsePtrOutputWithContext(context.Background())
}

func (i *setupScriptsResponsePtrType) ToSetupScriptsResponsePtrOutputWithContext(ctx context.Context) SetupScriptsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SetupScriptsResponsePtrOutput)
}

type SetupScriptsResponseOutput struct{ *pulumi.OutputState }

func (SetupScriptsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SetupScriptsResponse)(nil)).Elem()
}

func (o SetupScriptsResponseOutput) ToSetupScriptsResponseOutput() SetupScriptsResponseOutput {
	return o
}

func (o SetupScriptsResponseOutput) ToSetupScriptsResponseOutputWithContext(ctx context.Context) SetupScriptsResponseOutput {
	return o
}

func (o SetupScriptsResponseOutput) ToSetupScriptsResponsePtrOutput() SetupScriptsResponsePtrOutput {
	return o.ToSetupScriptsResponsePtrOutputWithContext(context.Background())
}

func (o SetupScriptsResponseOutput) ToSetupScriptsResponsePtrOutputWithContext(ctx context.Context) SetupScriptsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SetupScriptsResponse) *SetupScriptsResponse {
		return &v
	}).(SetupScriptsResponsePtrOutput)
}

func (o SetupScriptsResponseOutput) Scripts() ScriptsToExecuteResponsePtrOutput {
	return o.ApplyT(func(v SetupScriptsResponse) *ScriptsToExecuteResponse { return v.Scripts }).(ScriptsToExecuteResponsePtrOutput)
}

type SetupScriptsResponsePtrOutput struct{ *pulumi.OutputState }

func (SetupScriptsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SetupScriptsResponse)(nil)).Elem()
}

func (o SetupScriptsResponsePtrOutput) ToSetupScriptsResponsePtrOutput() SetupScriptsResponsePtrOutput {
	return o
}

func (o SetupScriptsResponsePtrOutput) ToSetupScriptsResponsePtrOutputWithContext(ctx context.Context) SetupScriptsResponsePtrOutput {
	return o
}

func (o SetupScriptsResponsePtrOutput) Elem() SetupScriptsResponseOutput {
	return o.ApplyT(func(v *SetupScriptsResponse) SetupScriptsResponse {
		if v != nil {
			return *v
		}
		var ret SetupScriptsResponse
		return ret
	}).(SetupScriptsResponseOutput)
}

func (o SetupScriptsResponsePtrOutput) Scripts() ScriptsToExecuteResponsePtrOutput {
	return o.ApplyT(func(v *SetupScriptsResponse) *ScriptsToExecuteResponse {
		if v == nil {
			return nil
		}
		return v.Scripts
	}).(ScriptsToExecuteResponsePtrOutput)
}

type SharedPrivateLinkResource struct {
	GroupId               *string `pulumi:"groupId"`
	Name                  *string `pulumi:"name"`
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	RequestMessage        *string `pulumi:"requestMessage"`
	Status                *string `pulumi:"status"`
}





type SharedPrivateLinkResourceInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput
	ToSharedPrivateLinkResourceOutputWithContext(context.Context) SharedPrivateLinkResourceOutput
}

type SharedPrivateLinkResourceArgs struct {
	GroupId               pulumi.StringPtrInput `pulumi:"groupId"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	PrivateLinkResourceId pulumi.StringPtrInput `pulumi:"privateLinkResourceId"`
	RequestMessage        pulumi.StringPtrInput `pulumi:"requestMessage"`
	Status                pulumi.StringPtrInput `pulumi:"status"`
}

func (SharedPrivateLinkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResource)(nil)).Elem()
}

func (i SharedPrivateLinkResourceArgs) ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput {
	return i.ToSharedPrivateLinkResourceOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceArgs) ToSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SharedPrivateLinkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceOutput)
}





type SharedPrivateLinkResourceArrayInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceArrayOutput() SharedPrivateLinkResourceArrayOutput
	ToSharedPrivateLinkResourceArrayOutputWithContext(context.Context) SharedPrivateLinkResourceArrayOutput
}

type SharedPrivateLinkResourceArray []SharedPrivateLinkResourceInput

func (SharedPrivateLinkResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResource)(nil)).Elem()
}

func (i SharedPrivateLinkResourceArray) ToSharedPrivateLinkResourceArrayOutput() SharedPrivateLinkResourceArrayOutput {
	return i.ToSharedPrivateLinkResourceArrayOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceArray) ToSharedPrivateLinkResourceArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceArrayOutput)
}

type SharedPrivateLinkResourceOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResource)(nil)).Elem()
}

func (o SharedPrivateLinkResourceOutput) ToSharedPrivateLinkResourceOutput() SharedPrivateLinkResourceOutput {
	return o
}

func (o SharedPrivateLinkResourceOutput) ToSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SharedPrivateLinkResourceOutput {
	return o
}

func (o SharedPrivateLinkResourceOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResource) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceArrayOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResource)(nil)).Elem()
}

func (o SharedPrivateLinkResourceArrayOutput) ToSharedPrivateLinkResourceArrayOutput() SharedPrivateLinkResourceArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceArrayOutput) ToSharedPrivateLinkResourceArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceArrayOutput) Index(i pulumi.IntInput) SharedPrivateLinkResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedPrivateLinkResource {
		return vs[0].([]SharedPrivateLinkResource)[vs[1].(int)]
	}).(SharedPrivateLinkResourceOutput)
}

type SharedPrivateLinkResourceResponse struct {
	GroupId               *string `pulumi:"groupId"`
	Name                  *string `pulumi:"name"`
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	RequestMessage        *string `pulumi:"requestMessage"`
	Status                *string `pulumi:"status"`
}





type SharedPrivateLinkResourceResponseInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput
	ToSharedPrivateLinkResourceResponseOutputWithContext(context.Context) SharedPrivateLinkResourceResponseOutput
}

type SharedPrivateLinkResourceResponseArgs struct {
	GroupId               pulumi.StringPtrInput `pulumi:"groupId"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	PrivateLinkResourceId pulumi.StringPtrInput `pulumi:"privateLinkResourceId"`
	RequestMessage        pulumi.StringPtrInput `pulumi:"requestMessage"`
	Status                pulumi.StringPtrInput `pulumi:"status"`
}

func (SharedPrivateLinkResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourceResponseArgs) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return i.ToSharedPrivateLinkResourceResponseOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceResponseArgs) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceResponseOutput)
}





type SharedPrivateLinkResourceResponseArrayInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput
	ToSharedPrivateLinkResourceResponseArrayOutputWithContext(context.Context) SharedPrivateLinkResourceResponseArrayOutput
}

type SharedPrivateLinkResourceResponseArray []SharedPrivateLinkResourceResponseInput

func (SharedPrivateLinkResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (i SharedPrivateLinkResourceResponseArray) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return i.ToSharedPrivateLinkResourceResponseArrayOutputWithContext(context.Background())
}

func (i SharedPrivateLinkResourceResponseArray) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPrivateLinkResourceResponseArrayOutput)
}

type SharedPrivateLinkResourceResponseOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) Index(i pulumi.IntInput) SharedPrivateLinkResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedPrivateLinkResourceResponse {
		return vs[0].([]SharedPrivateLinkResourceResponse)[vs[1].(int)]
	}).(SharedPrivateLinkResourceResponseOutput)
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
	Cert                    *string `pulumi:"cert"`
	Cname                   *string `pulumi:"cname"`
	Key                     *string `pulumi:"key"`
	LeafDomainLabel         *string `pulumi:"leafDomainLabel"`
	OverwriteExistingDomain *bool   `pulumi:"overwriteExistingDomain"`
	Status                  *string `pulumi:"status"`
}





type SslConfigurationInput interface {
	pulumi.Input

	ToSslConfigurationOutput() SslConfigurationOutput
	ToSslConfigurationOutputWithContext(context.Context) SslConfigurationOutput
}

type SslConfigurationArgs struct {
	Cert                    pulumi.StringPtrInput `pulumi:"cert"`
	Cname                   pulumi.StringPtrInput `pulumi:"cname"`
	Key                     pulumi.StringPtrInput `pulumi:"key"`
	LeafDomainLabel         pulumi.StringPtrInput `pulumi:"leafDomainLabel"`
	OverwriteExistingDomain pulumi.BoolPtrInput   `pulumi:"overwriteExistingDomain"`
	Status                  pulumi.StringPtrInput `pulumi:"status"`
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

func (o SslConfigurationOutput) LeafDomainLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *string { return v.LeafDomainLabel }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationOutput) OverwriteExistingDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SslConfiguration) *bool { return v.OverwriteExistingDomain }).(pulumi.BoolPtrOutput)
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

func (o SslConfigurationPtrOutput) LeafDomainLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.LeafDomainLabel
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationPtrOutput) OverwriteExistingDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SslConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.OverwriteExistingDomain
	}).(pulumi.BoolPtrOutput)
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
	Cert                    *string `pulumi:"cert"`
	Cname                   *string `pulumi:"cname"`
	Key                     *string `pulumi:"key"`
	LeafDomainLabel         *string `pulumi:"leafDomainLabel"`
	OverwriteExistingDomain *bool   `pulumi:"overwriteExistingDomain"`
	Status                  *string `pulumi:"status"`
}





type SslConfigurationResponseInput interface {
	pulumi.Input

	ToSslConfigurationResponseOutput() SslConfigurationResponseOutput
	ToSslConfigurationResponseOutputWithContext(context.Context) SslConfigurationResponseOutput
}

type SslConfigurationResponseArgs struct {
	Cert                    pulumi.StringPtrInput `pulumi:"cert"`
	Cname                   pulumi.StringPtrInput `pulumi:"cname"`
	Key                     pulumi.StringPtrInput `pulumi:"key"`
	LeafDomainLabel         pulumi.StringPtrInput `pulumi:"leafDomainLabel"`
	OverwriteExistingDomain pulumi.BoolPtrInput   `pulumi:"overwriteExistingDomain"`
	Status                  pulumi.StringPtrInput `pulumi:"status"`
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

func (o SslConfigurationResponseOutput) LeafDomainLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *string { return v.LeafDomainLabel }).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponseOutput) OverwriteExistingDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SslConfigurationResponse) *bool { return v.OverwriteExistingDomain }).(pulumi.BoolPtrOutput)
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

func (o SslConfigurationResponsePtrOutput) LeafDomainLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.LeafDomainLabel
	}).(pulumi.StringPtrOutput)
}

func (o SslConfigurationResponsePtrOutput) OverwriteExistingDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.OverwriteExistingDomain
	}).(pulumi.BoolPtrOutput)
}

func (o SslConfigurationResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SslConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type SynapseSpark struct {
	ComputeLocation  *string                 `pulumi:"computeLocation"`
	ComputeType      string                  `pulumi:"computeType"`
	Description      *string                 `pulumi:"description"`
	DisableLocalAuth *bool                   `pulumi:"disableLocalAuth"`
	Properties       *SynapseSparkProperties `pulumi:"properties"`
	ResourceId       *string                 `pulumi:"resourceId"`
}





type SynapseSparkInput interface {
	pulumi.Input

	ToSynapseSparkOutput() SynapseSparkOutput
	ToSynapseSparkOutputWithContext(context.Context) SynapseSparkOutput
}

type SynapseSparkArgs struct {
	ComputeLocation  pulumi.StringPtrInput          `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput             `pulumi:"computeType"`
	Description      pulumi.StringPtrInput          `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput            `pulumi:"disableLocalAuth"`
	Properties       SynapseSparkPropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput          `pulumi:"resourceId"`
}

func (SynapseSparkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSpark)(nil)).Elem()
}

func (i SynapseSparkArgs) ToSynapseSparkOutput() SynapseSparkOutput {
	return i.ToSynapseSparkOutputWithContext(context.Background())
}

func (i SynapseSparkArgs) ToSynapseSparkOutputWithContext(ctx context.Context) SynapseSparkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkOutput)
}

type SynapseSparkOutput struct{ *pulumi.OutputState }

func (SynapseSparkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSpark)(nil)).Elem()
}

func (o SynapseSparkOutput) ToSynapseSparkOutput() SynapseSparkOutput {
	return o
}

func (o SynapseSparkOutput) ToSynapseSparkOutputWithContext(ctx context.Context) SynapseSparkOutput {
	return o
}

func (o SynapseSparkOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSpark) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v SynapseSpark) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o SynapseSparkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSpark) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SynapseSpark) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o SynapseSparkOutput) Properties() SynapseSparkPropertiesPtrOutput {
	return o.ApplyT(func(v SynapseSpark) *SynapseSparkProperties { return v.Properties }).(SynapseSparkPropertiesPtrOutput)
}

func (o SynapseSparkOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSpark) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type SynapseSparkProperties struct {
	AutoPauseProperties *AutoPauseProperties `pulumi:"autoPauseProperties"`
	AutoScaleProperties *AutoScaleProperties `pulumi:"autoScaleProperties"`
	NodeCount           *int                 `pulumi:"nodeCount"`
	NodeSize            *string              `pulumi:"nodeSize"`
	NodeSizeFamily      *string              `pulumi:"nodeSizeFamily"`
	PoolName            *string              `pulumi:"poolName"`
	ResourceGroup       *string              `pulumi:"resourceGroup"`
	SparkVersion        *string              `pulumi:"sparkVersion"`
	SubscriptionId      *string              `pulumi:"subscriptionId"`
	WorkspaceName       *string              `pulumi:"workspaceName"`
}





type SynapseSparkPropertiesInput interface {
	pulumi.Input

	ToSynapseSparkPropertiesOutput() SynapseSparkPropertiesOutput
	ToSynapseSparkPropertiesOutputWithContext(context.Context) SynapseSparkPropertiesOutput
}

type SynapseSparkPropertiesArgs struct {
	AutoPauseProperties AutoPausePropertiesPtrInput `pulumi:"autoPauseProperties"`
	AutoScaleProperties AutoScalePropertiesPtrInput `pulumi:"autoScaleProperties"`
	NodeCount           pulumi.IntPtrInput          `pulumi:"nodeCount"`
	NodeSize            pulumi.StringPtrInput       `pulumi:"nodeSize"`
	NodeSizeFamily      pulumi.StringPtrInput       `pulumi:"nodeSizeFamily"`
	PoolName            pulumi.StringPtrInput       `pulumi:"poolName"`
	ResourceGroup       pulumi.StringPtrInput       `pulumi:"resourceGroup"`
	SparkVersion        pulumi.StringPtrInput       `pulumi:"sparkVersion"`
	SubscriptionId      pulumi.StringPtrInput       `pulumi:"subscriptionId"`
	WorkspaceName       pulumi.StringPtrInput       `pulumi:"workspaceName"`
}

func (SynapseSparkPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSparkProperties)(nil)).Elem()
}

func (i SynapseSparkPropertiesArgs) ToSynapseSparkPropertiesOutput() SynapseSparkPropertiesOutput {
	return i.ToSynapseSparkPropertiesOutputWithContext(context.Background())
}

func (i SynapseSparkPropertiesArgs) ToSynapseSparkPropertiesOutputWithContext(ctx context.Context) SynapseSparkPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkPropertiesOutput)
}

func (i SynapseSparkPropertiesArgs) ToSynapseSparkPropertiesPtrOutput() SynapseSparkPropertiesPtrOutput {
	return i.ToSynapseSparkPropertiesPtrOutputWithContext(context.Background())
}

func (i SynapseSparkPropertiesArgs) ToSynapseSparkPropertiesPtrOutputWithContext(ctx context.Context) SynapseSparkPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkPropertiesOutput).ToSynapseSparkPropertiesPtrOutputWithContext(ctx)
}









type SynapseSparkPropertiesPtrInput interface {
	pulumi.Input

	ToSynapseSparkPropertiesPtrOutput() SynapseSparkPropertiesPtrOutput
	ToSynapseSparkPropertiesPtrOutputWithContext(context.Context) SynapseSparkPropertiesPtrOutput
}

type synapseSparkPropertiesPtrType SynapseSparkPropertiesArgs

func SynapseSparkPropertiesPtr(v *SynapseSparkPropertiesArgs) SynapseSparkPropertiesPtrInput {
	return (*synapseSparkPropertiesPtrType)(v)
}

func (*synapseSparkPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseSparkProperties)(nil)).Elem()
}

func (i *synapseSparkPropertiesPtrType) ToSynapseSparkPropertiesPtrOutput() SynapseSparkPropertiesPtrOutput {
	return i.ToSynapseSparkPropertiesPtrOutputWithContext(context.Background())
}

func (i *synapseSparkPropertiesPtrType) ToSynapseSparkPropertiesPtrOutputWithContext(ctx context.Context) SynapseSparkPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkPropertiesPtrOutput)
}

type SynapseSparkPropertiesOutput struct{ *pulumi.OutputState }

func (SynapseSparkPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSparkProperties)(nil)).Elem()
}

func (o SynapseSparkPropertiesOutput) ToSynapseSparkPropertiesOutput() SynapseSparkPropertiesOutput {
	return o
}

func (o SynapseSparkPropertiesOutput) ToSynapseSparkPropertiesOutputWithContext(ctx context.Context) SynapseSparkPropertiesOutput {
	return o
}

func (o SynapseSparkPropertiesOutput) ToSynapseSparkPropertiesPtrOutput() SynapseSparkPropertiesPtrOutput {
	return o.ToSynapseSparkPropertiesPtrOutputWithContext(context.Background())
}

func (o SynapseSparkPropertiesOutput) ToSynapseSparkPropertiesPtrOutputWithContext(ctx context.Context) SynapseSparkPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SynapseSparkProperties) *SynapseSparkProperties {
		return &v
	}).(SynapseSparkPropertiesPtrOutput)
}

func (o SynapseSparkPropertiesOutput) AutoPauseProperties() AutoPausePropertiesPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *AutoPauseProperties { return v.AutoPauseProperties }).(AutoPausePropertiesPtrOutput)
}

func (o SynapseSparkPropertiesOutput) AutoScaleProperties() AutoScalePropertiesPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *AutoScaleProperties { return v.AutoScaleProperties }).(AutoScalePropertiesPtrOutput)
}

func (o SynapseSparkPropertiesOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

func (o SynapseSparkPropertiesOutput) NodeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *string { return v.NodeSize }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesOutput) NodeSizeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *string { return v.NodeSizeFamily }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesOutput) PoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *string { return v.PoolName }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesOutput) SparkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *string { return v.SparkVersion }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesOutput) WorkspaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkProperties) *string { return v.WorkspaceName }).(pulumi.StringPtrOutput)
}

type SynapseSparkPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SynapseSparkPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseSparkProperties)(nil)).Elem()
}

func (o SynapseSparkPropertiesPtrOutput) ToSynapseSparkPropertiesPtrOutput() SynapseSparkPropertiesPtrOutput {
	return o
}

func (o SynapseSparkPropertiesPtrOutput) ToSynapseSparkPropertiesPtrOutputWithContext(ctx context.Context) SynapseSparkPropertiesPtrOutput {
	return o
}

func (o SynapseSparkPropertiesPtrOutput) Elem() SynapseSparkPropertiesOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) SynapseSparkProperties {
		if v != nil {
			return *v
		}
		var ret SynapseSparkProperties
		return ret
	}).(SynapseSparkPropertiesOutput)
}

func (o SynapseSparkPropertiesPtrOutput) AutoPauseProperties() AutoPausePropertiesPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *AutoPauseProperties {
		if v == nil {
			return nil
		}
		return v.AutoPauseProperties
	}).(AutoPausePropertiesPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) AutoScaleProperties() AutoScalePropertiesPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *AutoScaleProperties {
		if v == nil {
			return nil
		}
		return v.AutoScaleProperties
	}).(AutoScalePropertiesPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *int {
		if v == nil {
			return nil
		}
		return v.NodeCount
	}).(pulumi.IntPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) NodeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *string {
		if v == nil {
			return nil
		}
		return v.NodeSize
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) NodeSizeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *string {
		if v == nil {
			return nil
		}
		return v.NodeSizeFamily
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) PoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *string {
		if v == nil {
			return nil
		}
		return v.PoolName
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) SparkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *string {
		if v == nil {
			return nil
		}
		return v.SparkVersion
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkPropertiesPtrOutput) WorkspaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceName
	}).(pulumi.StringPtrOutput)
}

type SynapseSparkResponse struct {
	ComputeLocation    *string                         `pulumi:"computeLocation"`
	ComputeType        string                          `pulumi:"computeType"`
	CreatedOn          string                          `pulumi:"createdOn"`
	Description        *string                         `pulumi:"description"`
	DisableLocalAuth   *bool                           `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                            `pulumi:"isAttachedCompute"`
	ModifiedOn         string                          `pulumi:"modifiedOn"`
	Properties         *SynapseSparkResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse         `pulumi:"provisioningErrors"`
	ProvisioningState  string                          `pulumi:"provisioningState"`
	ResourceId         *string                         `pulumi:"resourceId"`
}





type SynapseSparkResponseInput interface {
	pulumi.Input

	ToSynapseSparkResponseOutput() SynapseSparkResponseOutput
	ToSynapseSparkResponseOutputWithContext(context.Context) SynapseSparkResponseOutput
}

type SynapseSparkResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                  `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                     `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                     `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                  `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput                    `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                       `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                     `pulumi:"modifiedOn"`
	Properties         SynapseSparkResponsePropertiesPtrInput `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput        `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                     `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                  `pulumi:"resourceId"`
}

func (SynapseSparkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSparkResponse)(nil)).Elem()
}

func (i SynapseSparkResponseArgs) ToSynapseSparkResponseOutput() SynapseSparkResponseOutput {
	return i.ToSynapseSparkResponseOutputWithContext(context.Background())
}

func (i SynapseSparkResponseArgs) ToSynapseSparkResponseOutputWithContext(ctx context.Context) SynapseSparkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkResponseOutput)
}

type SynapseSparkResponseOutput struct{ *pulumi.OutputState }

func (SynapseSparkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSparkResponse)(nil)).Elem()
}

func (o SynapseSparkResponseOutput) ToSynapseSparkResponseOutput() SynapseSparkResponseOutput {
	return o
}

func (o SynapseSparkResponseOutput) ToSynapseSparkResponseOutputWithContext(ctx context.Context) SynapseSparkResponseOutput {
	return o
}

func (o SynapseSparkResponseOutput) ComputeLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponse) *string { return v.ComputeLocation }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponseOutput) ComputeType() pulumi.StringOutput {
	return o.ApplyT(func(v SynapseSparkResponse) string { return v.ComputeType }).(pulumi.StringOutput)
}

func (o SynapseSparkResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v SynapseSparkResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o SynapseSparkResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o SynapseSparkResponseOutput) IsAttachedCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v SynapseSparkResponse) bool { return v.IsAttachedCompute }).(pulumi.BoolOutput)
}

func (o SynapseSparkResponseOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v SynapseSparkResponse) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o SynapseSparkResponseOutput) Properties() SynapseSparkResponsePropertiesPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponse) *SynapseSparkResponseProperties { return v.Properties }).(SynapseSparkResponsePropertiesPtrOutput)
}

func (o SynapseSparkResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v SynapseSparkResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o SynapseSparkResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SynapseSparkResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SynapseSparkResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type SynapseSparkResponseProperties struct {
	AutoPauseProperties *AutoPausePropertiesResponse `pulumi:"autoPauseProperties"`
	AutoScaleProperties *AutoScalePropertiesResponse `pulumi:"autoScaleProperties"`
	NodeCount           *int                         `pulumi:"nodeCount"`
	NodeSize            *string                      `pulumi:"nodeSize"`
	NodeSizeFamily      *string                      `pulumi:"nodeSizeFamily"`
	PoolName            *string                      `pulumi:"poolName"`
	ResourceGroup       *string                      `pulumi:"resourceGroup"`
	SparkVersion        *string                      `pulumi:"sparkVersion"`
	SubscriptionId      *string                      `pulumi:"subscriptionId"`
	WorkspaceName       *string                      `pulumi:"workspaceName"`
}





type SynapseSparkResponsePropertiesInput interface {
	pulumi.Input

	ToSynapseSparkResponsePropertiesOutput() SynapseSparkResponsePropertiesOutput
	ToSynapseSparkResponsePropertiesOutputWithContext(context.Context) SynapseSparkResponsePropertiesOutput
}

type SynapseSparkResponsePropertiesArgs struct {
	AutoPauseProperties AutoPausePropertiesResponsePtrInput `pulumi:"autoPauseProperties"`
	AutoScaleProperties AutoScalePropertiesResponsePtrInput `pulumi:"autoScaleProperties"`
	NodeCount           pulumi.IntPtrInput                  `pulumi:"nodeCount"`
	NodeSize            pulumi.StringPtrInput               `pulumi:"nodeSize"`
	NodeSizeFamily      pulumi.StringPtrInput               `pulumi:"nodeSizeFamily"`
	PoolName            pulumi.StringPtrInput               `pulumi:"poolName"`
	ResourceGroup       pulumi.StringPtrInput               `pulumi:"resourceGroup"`
	SparkVersion        pulumi.StringPtrInput               `pulumi:"sparkVersion"`
	SubscriptionId      pulumi.StringPtrInput               `pulumi:"subscriptionId"`
	WorkspaceName       pulumi.StringPtrInput               `pulumi:"workspaceName"`
}

func (SynapseSparkResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSparkResponseProperties)(nil)).Elem()
}

func (i SynapseSparkResponsePropertiesArgs) ToSynapseSparkResponsePropertiesOutput() SynapseSparkResponsePropertiesOutput {
	return i.ToSynapseSparkResponsePropertiesOutputWithContext(context.Background())
}

func (i SynapseSparkResponsePropertiesArgs) ToSynapseSparkResponsePropertiesOutputWithContext(ctx context.Context) SynapseSparkResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkResponsePropertiesOutput)
}

func (i SynapseSparkResponsePropertiesArgs) ToSynapseSparkResponsePropertiesPtrOutput() SynapseSparkResponsePropertiesPtrOutput {
	return i.ToSynapseSparkResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i SynapseSparkResponsePropertiesArgs) ToSynapseSparkResponsePropertiesPtrOutputWithContext(ctx context.Context) SynapseSparkResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkResponsePropertiesOutput).ToSynapseSparkResponsePropertiesPtrOutputWithContext(ctx)
}









type SynapseSparkResponsePropertiesPtrInput interface {
	pulumi.Input

	ToSynapseSparkResponsePropertiesPtrOutput() SynapseSparkResponsePropertiesPtrOutput
	ToSynapseSparkResponsePropertiesPtrOutputWithContext(context.Context) SynapseSparkResponsePropertiesPtrOutput
}

type synapseSparkResponsePropertiesPtrType SynapseSparkResponsePropertiesArgs

func SynapseSparkResponsePropertiesPtr(v *SynapseSparkResponsePropertiesArgs) SynapseSparkResponsePropertiesPtrInput {
	return (*synapseSparkResponsePropertiesPtrType)(v)
}

func (*synapseSparkResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseSparkResponseProperties)(nil)).Elem()
}

func (i *synapseSparkResponsePropertiesPtrType) ToSynapseSparkResponsePropertiesPtrOutput() SynapseSparkResponsePropertiesPtrOutput {
	return i.ToSynapseSparkResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *synapseSparkResponsePropertiesPtrType) ToSynapseSparkResponsePropertiesPtrOutputWithContext(ctx context.Context) SynapseSparkResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseSparkResponsePropertiesPtrOutput)
}

type SynapseSparkResponsePropertiesOutput struct{ *pulumi.OutputState }

func (SynapseSparkResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SynapseSparkResponseProperties)(nil)).Elem()
}

func (o SynapseSparkResponsePropertiesOutput) ToSynapseSparkResponsePropertiesOutput() SynapseSparkResponsePropertiesOutput {
	return o
}

func (o SynapseSparkResponsePropertiesOutput) ToSynapseSparkResponsePropertiesOutputWithContext(ctx context.Context) SynapseSparkResponsePropertiesOutput {
	return o
}

func (o SynapseSparkResponsePropertiesOutput) ToSynapseSparkResponsePropertiesPtrOutput() SynapseSparkResponsePropertiesPtrOutput {
	return o.ToSynapseSparkResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o SynapseSparkResponsePropertiesOutput) ToSynapseSparkResponsePropertiesPtrOutputWithContext(ctx context.Context) SynapseSparkResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SynapseSparkResponseProperties) *SynapseSparkResponseProperties {
		return &v
	}).(SynapseSparkResponsePropertiesPtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) AutoPauseProperties() AutoPausePropertiesResponsePtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *AutoPausePropertiesResponse { return v.AutoPauseProperties }).(AutoPausePropertiesResponsePtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) AutoScaleProperties() AutoScalePropertiesResponsePtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *AutoScalePropertiesResponse { return v.AutoScaleProperties }).(AutoScalePropertiesResponsePtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) NodeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *string { return v.NodeSize }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) NodeSizeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *string { return v.NodeSizeFamily }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) PoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *string { return v.PoolName }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) SparkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *string { return v.SparkVersion }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesOutput) WorkspaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SynapseSparkResponseProperties) *string { return v.WorkspaceName }).(pulumi.StringPtrOutput)
}

type SynapseSparkResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SynapseSparkResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseSparkResponseProperties)(nil)).Elem()
}

func (o SynapseSparkResponsePropertiesPtrOutput) ToSynapseSparkResponsePropertiesPtrOutput() SynapseSparkResponsePropertiesPtrOutput {
	return o
}

func (o SynapseSparkResponsePropertiesPtrOutput) ToSynapseSparkResponsePropertiesPtrOutputWithContext(ctx context.Context) SynapseSparkResponsePropertiesPtrOutput {
	return o
}

func (o SynapseSparkResponsePropertiesPtrOutput) Elem() SynapseSparkResponsePropertiesOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) SynapseSparkResponseProperties {
		if v != nil {
			return *v
		}
		var ret SynapseSparkResponseProperties
		return ret
	}).(SynapseSparkResponsePropertiesOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) AutoPauseProperties() AutoPausePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *AutoPausePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.AutoPauseProperties
	}).(AutoPausePropertiesResponsePtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) AutoScaleProperties() AutoScalePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *AutoScalePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.AutoScaleProperties
	}).(AutoScalePropertiesResponsePtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.NodeCount
	}).(pulumi.IntPtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) NodeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.NodeSize
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) NodeSizeFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.NodeSizeFamily
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) PoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.PoolName
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) SparkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.SparkVersion
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o SynapseSparkResponsePropertiesPtrOutput) WorkspaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SynapseSparkResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.WorkspaceName
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (UserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return i.ToUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseOutput)
}





type UserAssignedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput
	ToUserAssignedIdentityResponseMapOutputWithContext(context.Context) UserAssignedIdentityResponseMapOutput
}

type UserAssignedIdentityResponseMap map[string]UserAssignedIdentityResponseInput

func (UserAssignedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return i.ToUserAssignedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseMapOutput)
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

func (o UserAssignedIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
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

type VirtualMachine struct {
	ComputeLocation  *string                   `pulumi:"computeLocation"`
	ComputeType      string                    `pulumi:"computeType"`
	Description      *string                   `pulumi:"description"`
	DisableLocalAuth *bool                     `pulumi:"disableLocalAuth"`
	Properties       *VirtualMachineProperties `pulumi:"properties"`
	ResourceId       *string                   `pulumi:"resourceId"`
}





type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(context.Context) VirtualMachineOutput
}

type VirtualMachineArgs struct {
	ComputeLocation  pulumi.StringPtrInput            `pulumi:"computeLocation"`
	ComputeType      pulumi.StringInput               `pulumi:"computeType"`
	Description      pulumi.StringPtrInput            `pulumi:"description"`
	DisableLocalAuth pulumi.BoolPtrInput              `pulumi:"disableLocalAuth"`
	Properties       VirtualMachinePropertiesPtrInput `pulumi:"properties"`
	ResourceId       pulumi.StringPtrInput            `pulumi:"resourceId"`
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

func (o VirtualMachineOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineOutput) Properties() VirtualMachinePropertiesPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *VirtualMachineProperties { return v.Properties }).(VirtualMachinePropertiesPtrOutput)
}

func (o VirtualMachineOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachine) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VirtualMachineImage struct {
	Id string `pulumi:"id"`
}





type VirtualMachineImageInput interface {
	pulumi.Input

	ToVirtualMachineImageOutput() VirtualMachineImageOutput
	ToVirtualMachineImageOutputWithContext(context.Context) VirtualMachineImageOutput
}

type VirtualMachineImageArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (VirtualMachineImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImage)(nil)).Elem()
}

func (i VirtualMachineImageArgs) ToVirtualMachineImageOutput() VirtualMachineImageOutput {
	return i.ToVirtualMachineImageOutputWithContext(context.Background())
}

func (i VirtualMachineImageArgs) ToVirtualMachineImageOutputWithContext(ctx context.Context) VirtualMachineImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageOutput)
}

func (i VirtualMachineImageArgs) ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput {
	return i.ToVirtualMachineImagePtrOutputWithContext(context.Background())
}

func (i VirtualMachineImageArgs) ToVirtualMachineImagePtrOutputWithContext(ctx context.Context) VirtualMachineImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageOutput).ToVirtualMachineImagePtrOutputWithContext(ctx)
}









type VirtualMachineImagePtrInput interface {
	pulumi.Input

	ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput
	ToVirtualMachineImagePtrOutputWithContext(context.Context) VirtualMachineImagePtrOutput
}

type virtualMachineImagePtrType VirtualMachineImageArgs

func VirtualMachineImagePtr(v *VirtualMachineImageArgs) VirtualMachineImagePtrInput {
	return (*virtualMachineImagePtrType)(v)
}

func (*virtualMachineImagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImage)(nil)).Elem()
}

func (i *virtualMachineImagePtrType) ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput {
	return i.ToVirtualMachineImagePtrOutputWithContext(context.Background())
}

func (i *virtualMachineImagePtrType) ToVirtualMachineImagePtrOutputWithContext(ctx context.Context) VirtualMachineImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImagePtrOutput)
}

type VirtualMachineImageOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImage)(nil)).Elem()
}

func (o VirtualMachineImageOutput) ToVirtualMachineImageOutput() VirtualMachineImageOutput {
	return o
}

func (o VirtualMachineImageOutput) ToVirtualMachineImageOutputWithContext(ctx context.Context) VirtualMachineImageOutput {
	return o
}

func (o VirtualMachineImageOutput) ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput {
	return o.ToVirtualMachineImagePtrOutputWithContext(context.Background())
}

func (o VirtualMachineImageOutput) ToVirtualMachineImagePtrOutputWithContext(ctx context.Context) VirtualMachineImagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineImage) *VirtualMachineImage {
		return &v
	}).(VirtualMachineImagePtrOutput)
}

func (o VirtualMachineImageOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineImage) string { return v.Id }).(pulumi.StringOutput)
}

type VirtualMachineImagePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineImagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImage)(nil)).Elem()
}

func (o VirtualMachineImagePtrOutput) ToVirtualMachineImagePtrOutput() VirtualMachineImagePtrOutput {
	return o
}

func (o VirtualMachineImagePtrOutput) ToVirtualMachineImagePtrOutputWithContext(ctx context.Context) VirtualMachineImagePtrOutput {
	return o
}

func (o VirtualMachineImagePtrOutput) Elem() VirtualMachineImageOutput {
	return o.ApplyT(func(v *VirtualMachineImage) VirtualMachineImage {
		if v != nil {
			return *v
		}
		var ret VirtualMachineImage
		return ret
	}).(VirtualMachineImageOutput)
}

func (o VirtualMachineImagePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineImage) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineImageResponse struct {
	Id string `pulumi:"id"`
}





type VirtualMachineImageResponseInput interface {
	pulumi.Input

	ToVirtualMachineImageResponseOutput() VirtualMachineImageResponseOutput
	ToVirtualMachineImageResponseOutputWithContext(context.Context) VirtualMachineImageResponseOutput
}

type VirtualMachineImageResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (VirtualMachineImageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImageResponse)(nil)).Elem()
}

func (i VirtualMachineImageResponseArgs) ToVirtualMachineImageResponseOutput() VirtualMachineImageResponseOutput {
	return i.ToVirtualMachineImageResponseOutputWithContext(context.Background())
}

func (i VirtualMachineImageResponseArgs) ToVirtualMachineImageResponseOutputWithContext(ctx context.Context) VirtualMachineImageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageResponseOutput)
}

func (i VirtualMachineImageResponseArgs) ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput {
	return i.ToVirtualMachineImageResponsePtrOutputWithContext(context.Background())
}

func (i VirtualMachineImageResponseArgs) ToVirtualMachineImageResponsePtrOutputWithContext(ctx context.Context) VirtualMachineImageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageResponseOutput).ToVirtualMachineImageResponsePtrOutputWithContext(ctx)
}









type VirtualMachineImageResponsePtrInput interface {
	pulumi.Input

	ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput
	ToVirtualMachineImageResponsePtrOutputWithContext(context.Context) VirtualMachineImageResponsePtrOutput
}

type virtualMachineImageResponsePtrType VirtualMachineImageResponseArgs

func VirtualMachineImageResponsePtr(v *VirtualMachineImageResponseArgs) VirtualMachineImageResponsePtrInput {
	return (*virtualMachineImageResponsePtrType)(v)
}

func (*virtualMachineImageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImageResponse)(nil)).Elem()
}

func (i *virtualMachineImageResponsePtrType) ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput {
	return i.ToVirtualMachineImageResponsePtrOutputWithContext(context.Background())
}

func (i *virtualMachineImageResponsePtrType) ToVirtualMachineImageResponsePtrOutputWithContext(ctx context.Context) VirtualMachineImageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageResponsePtrOutput)
}

type VirtualMachineImageResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineImageResponse)(nil)).Elem()
}

func (o VirtualMachineImageResponseOutput) ToVirtualMachineImageResponseOutput() VirtualMachineImageResponseOutput {
	return o
}

func (o VirtualMachineImageResponseOutput) ToVirtualMachineImageResponseOutputWithContext(ctx context.Context) VirtualMachineImageResponseOutput {
	return o
}

func (o VirtualMachineImageResponseOutput) ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput {
	return o.ToVirtualMachineImageResponsePtrOutputWithContext(context.Background())
}

func (o VirtualMachineImageResponseOutput) ToVirtualMachineImageResponsePtrOutputWithContext(ctx context.Context) VirtualMachineImageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineImageResponse) *VirtualMachineImageResponse {
		return &v
	}).(VirtualMachineImageResponsePtrOutput)
}

func (o VirtualMachineImageResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineImageResponse) string { return v.Id }).(pulumi.StringOutput)
}

type VirtualMachineImageResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImageResponse)(nil)).Elem()
}

func (o VirtualMachineImageResponsePtrOutput) ToVirtualMachineImageResponsePtrOutput() VirtualMachineImageResponsePtrOutput {
	return o
}

func (o VirtualMachineImageResponsePtrOutput) ToVirtualMachineImageResponsePtrOutputWithContext(ctx context.Context) VirtualMachineImageResponsePtrOutput {
	return o
}

func (o VirtualMachineImageResponsePtrOutput) Elem() VirtualMachineImageResponseOutput {
	return o.ApplyT(func(v *VirtualMachineImageResponse) VirtualMachineImageResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineImageResponse
		return ret
	}).(VirtualMachineImageResponseOutput)
}

func (o VirtualMachineImageResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineImageResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineProperties struct {
	Address                   *string                       `pulumi:"address"`
	AdministratorAccount      *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	IsNotebookInstanceCompute *bool                         `pulumi:"isNotebookInstanceCompute"`
	SshPort                   *int                          `pulumi:"sshPort"`
	VirtualMachineSize        *string                       `pulumi:"virtualMachineSize"`
}





type VirtualMachinePropertiesInput interface {
	pulumi.Input

	ToVirtualMachinePropertiesOutput() VirtualMachinePropertiesOutput
	ToVirtualMachinePropertiesOutputWithContext(context.Context) VirtualMachinePropertiesOutput
}

type VirtualMachinePropertiesArgs struct {
	Address                   pulumi.StringPtrInput                `pulumi:"address"`
	AdministratorAccount      VirtualMachineSshCredentialsPtrInput `pulumi:"administratorAccount"`
	IsNotebookInstanceCompute pulumi.BoolPtrInput                  `pulumi:"isNotebookInstanceCompute"`
	SshPort                   pulumi.IntPtrInput                   `pulumi:"sshPort"`
	VirtualMachineSize        pulumi.StringPtrInput                `pulumi:"virtualMachineSize"`
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

func (o VirtualMachinePropertiesOutput) IsNotebookInstanceCompute() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineProperties) *bool { return v.IsNotebookInstanceCompute }).(pulumi.BoolPtrOutput)
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

func (o VirtualMachinePropertiesPtrOutput) IsNotebookInstanceCompute() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsNotebookInstanceCompute
	}).(pulumi.BoolPtrOutput)
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
	ComputeLocation    *string                           `pulumi:"computeLocation"`
	ComputeType        string                            `pulumi:"computeType"`
	CreatedOn          string                            `pulumi:"createdOn"`
	Description        *string                           `pulumi:"description"`
	DisableLocalAuth   *bool                             `pulumi:"disableLocalAuth"`
	IsAttachedCompute  bool                              `pulumi:"isAttachedCompute"`
	ModifiedOn         string                            `pulumi:"modifiedOn"`
	Properties         *VirtualMachineResponseProperties `pulumi:"properties"`
	ProvisioningErrors []ErrorResponseResponse           `pulumi:"provisioningErrors"`
	ProvisioningState  string                            `pulumi:"provisioningState"`
	ResourceId         *string                           `pulumi:"resourceId"`
}





type VirtualMachineResponseInput interface {
	pulumi.Input

	ToVirtualMachineResponseOutput() VirtualMachineResponseOutput
	ToVirtualMachineResponseOutputWithContext(context.Context) VirtualMachineResponseOutput
}

type VirtualMachineResponseArgs struct {
	ComputeLocation    pulumi.StringPtrInput                    `pulumi:"computeLocation"`
	ComputeType        pulumi.StringInput                       `pulumi:"computeType"`
	CreatedOn          pulumi.StringInput                       `pulumi:"createdOn"`
	Description        pulumi.StringPtrInput                    `pulumi:"description"`
	DisableLocalAuth   pulumi.BoolPtrInput                      `pulumi:"disableLocalAuth"`
	IsAttachedCompute  pulumi.BoolInput                         `pulumi:"isAttachedCompute"`
	ModifiedOn         pulumi.StringInput                       `pulumi:"modifiedOn"`
	Properties         VirtualMachineResponsePropertiesPtrInput `pulumi:"properties"`
	ProvisioningErrors ErrorResponseResponseArrayInput          `pulumi:"provisioningErrors"`
	ProvisioningState  pulumi.StringInput                       `pulumi:"provisioningState"`
	ResourceId         pulumi.StringPtrInput                    `pulumi:"resourceId"`
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

func (o VirtualMachineResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
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

func (o VirtualMachineResponseOutput) ProvisioningErrors() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineResponse) []ErrorResponseResponse { return v.ProvisioningErrors }).(ErrorResponseResponseArrayOutput)
}

func (o VirtualMachineResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VirtualMachineResponseProperties struct {
	Address                   *string                               `pulumi:"address"`
	AdministratorAccount      *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	IsNotebookInstanceCompute *bool                                 `pulumi:"isNotebookInstanceCompute"`
	SshPort                   *int                                  `pulumi:"sshPort"`
	VirtualMachineSize        *string                               `pulumi:"virtualMachineSize"`
}





type VirtualMachineResponsePropertiesInput interface {
	pulumi.Input

	ToVirtualMachineResponsePropertiesOutput() VirtualMachineResponsePropertiesOutput
	ToVirtualMachineResponsePropertiesOutputWithContext(context.Context) VirtualMachineResponsePropertiesOutput
}

type VirtualMachineResponsePropertiesArgs struct {
	Address                   pulumi.StringPtrInput                        `pulumi:"address"`
	AdministratorAccount      VirtualMachineSshCredentialsResponsePtrInput `pulumi:"administratorAccount"`
	IsNotebookInstanceCompute pulumi.BoolPtrInput                          `pulumi:"isNotebookInstanceCompute"`
	SshPort                   pulumi.IntPtrInput                           `pulumi:"sshPort"`
	VirtualMachineSize        pulumi.StringPtrInput                        `pulumi:"virtualMachineSize"`
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

func (o VirtualMachineResponsePropertiesOutput) IsNotebookInstanceCompute() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineResponseProperties) *bool { return v.IsNotebookInstanceCompute }).(pulumi.BoolPtrOutput)
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

func (o VirtualMachineResponsePropertiesPtrOutput) IsNotebookInstanceCompute() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsNotebookInstanceCompute
	}).(pulumi.BoolPtrOutput)
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
	pulumi.RegisterOutputType(AmlComputePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AmlComputePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AmlComputeResponseOutput{})
	pulumi.RegisterOutputType(AssignedUserOutput{})
	pulumi.RegisterOutputType(AssignedUserPtrOutput{})
	pulumi.RegisterOutputType(AssignedUserResponseOutput{})
	pulumi.RegisterOutputType(AssignedUserResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoPausePropertiesOutput{})
	pulumi.RegisterOutputType(AutoPausePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoPausePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoPausePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoScalePropertiesOutput{})
	pulumi.RegisterOutputType(AutoScalePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoScalePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoScalePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeInstanceOutput{})
	pulumi.RegisterOutputType(ComputeInstanceApplicationResponseOutput{})
	pulumi.RegisterOutputType(ComputeInstanceApplicationResponseArrayOutput{})
	pulumi.RegisterOutputType(ComputeInstanceConnectivityEndpointsResponseOutput{})
	pulumi.RegisterOutputType(ComputeInstanceConnectivityEndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeInstanceCreatedByResponseOutput{})
	pulumi.RegisterOutputType(ComputeInstanceCreatedByResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeInstanceLastOperationResponseOutput{})
	pulumi.RegisterOutputType(ComputeInstanceLastOperationResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeInstancePropertiesOutput{})
	pulumi.RegisterOutputType(ComputeInstancePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ComputeInstancePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ComputeInstancePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeInstanceResponseOutput{})
	pulumi.RegisterOutputType(ComputeInstanceSshSettingsOutput{})
	pulumi.RegisterOutputType(ComputeInstanceSshSettingsPtrOutput{})
	pulumi.RegisterOutputType(ComputeInstanceSshSettingsResponseOutput{})
	pulumi.RegisterOutputType(ComputeInstanceSshSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsPtrOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsResponseOutput{})
	pulumi.RegisterOutputType(CosmosDbSettingsResponsePtrOutput{})
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
	pulumi.RegisterOutputType(DatabricksPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DatabricksPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DatabricksResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseArrayOutput{})
	pulumi.RegisterOutputType(HDInsightOutput{})
	pulumi.RegisterOutputType(HDInsightPropertiesOutput{})
	pulumi.RegisterOutputType(HDInsightPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HDInsightPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HDInsightPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(HDInsightResponseOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityForCmkOutput{})
	pulumi.RegisterOutputType(IdentityForCmkPtrOutput{})
	pulumi.RegisterOutputType(IdentityForCmkResponseOutput{})
	pulumi.RegisterOutputType(IdentityForCmkResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceTypeSchemaOutput{})
	pulumi.RegisterOutputType(InstanceTypeSchemaMapOutput{})
	pulumi.RegisterOutputType(InstanceTypeSchemaResourcesOutput{})
	pulumi.RegisterOutputType(InstanceTypeSchemaResourcesPtrOutput{})
	pulumi.RegisterOutputType(InstanceTypeSchemaResponseOutput{})
	pulumi.RegisterOutputType(InstanceTypeSchemaResponseMapOutput{})
	pulumi.RegisterOutputType(InstanceTypeSchemaResponseResourcesOutput{})
	pulumi.RegisterOutputType(InstanceTypeSchemaResponseResourcesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesOutput{})
	pulumi.RegisterOutputType(KubernetesPropertiesOutput{})
	pulumi.RegisterOutputType(KubernetesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KubernetesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KubernetesPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KubernetesResponseOutput{})
	pulumi.RegisterOutputType(ListNotebookKeysResultResponseOutput{})
	pulumi.RegisterOutputType(NodeStateCountsResponseOutput{})
	pulumi.RegisterOutputType(NodeStateCountsResponsePtrOutput{})
	pulumi.RegisterOutputType(NotebookPreparationErrorResponseOutput{})
	pulumi.RegisterOutputType(NotebookPreparationErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(NotebookResourceInfoResponseOutput{})
	pulumi.RegisterOutputType(NotebookResourceInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(PasswordResponseOutput{})
	pulumi.RegisterOutputType(PasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(PersonalComputeInstanceSettingsOutput{})
	pulumi.RegisterOutputType(PersonalComputeInstanceSettingsPtrOutput{})
	pulumi.RegisterOutputType(PersonalComputeInstanceSettingsResponseOutput{})
	pulumi.RegisterOutputType(PersonalComputeInstanceSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ScriptReferenceOutput{})
	pulumi.RegisterOutputType(ScriptReferencePtrOutput{})
	pulumi.RegisterOutputType(ScriptReferenceResponseOutput{})
	pulumi.RegisterOutputType(ScriptReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ScriptsToExecuteOutput{})
	pulumi.RegisterOutputType(ScriptsToExecutePtrOutput{})
	pulumi.RegisterOutputType(ScriptsToExecuteResponseOutput{})
	pulumi.RegisterOutputType(ScriptsToExecuteResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceManagedResourcesSettingsOutput{})
	pulumi.RegisterOutputType(ServiceManagedResourcesSettingsPtrOutput{})
	pulumi.RegisterOutputType(ServiceManagedResourcesSettingsResponseOutput{})
	pulumi.RegisterOutputType(ServiceManagedResourcesSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SetupScriptsOutput{})
	pulumi.RegisterOutputType(SetupScriptsPtrOutput{})
	pulumi.RegisterOutputType(SetupScriptsResponseOutput{})
	pulumi.RegisterOutputType(SetupScriptsResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceArrayOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SslConfigurationOutput{})
	pulumi.RegisterOutputType(SslConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SslConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SslConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SynapseSparkOutput{})
	pulumi.RegisterOutputType(SynapseSparkPropertiesOutput{})
	pulumi.RegisterOutputType(SynapseSparkPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SynapseSparkResponseOutput{})
	pulumi.RegisterOutputType(SynapseSparkResponsePropertiesOutput{})
	pulumi.RegisterOutputType(SynapseSparkResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemServiceResponseOutput{})
	pulumi.RegisterOutputType(SystemServiceResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsPtrOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsResponseOutput{})
	pulumi.RegisterOutputType(UserAccountCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualMachineOutput{})
	pulumi.RegisterOutputType(VirtualMachineImageOutput{})
	pulumi.RegisterOutputType(VirtualMachineImagePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineImageResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineImageResponsePtrOutput{})
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
