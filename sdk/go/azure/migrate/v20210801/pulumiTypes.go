


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutomaticResolutionPropertiesResponse struct {
	MoveResourceId *string `pulumi:"moveResourceId"`
}

type AutomaticResolutionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutomaticResolutionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticResolutionPropertiesResponse)(nil)).Elem()
}

func (o AutomaticResolutionPropertiesResponseOutput) ToAutomaticResolutionPropertiesResponseOutput() AutomaticResolutionPropertiesResponseOutput {
	return o
}

func (o AutomaticResolutionPropertiesResponseOutput) ToAutomaticResolutionPropertiesResponseOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponseOutput {
	return o
}

func (o AutomaticResolutionPropertiesResponseOutput) MoveResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomaticResolutionPropertiesResponse) *string { return v.MoveResourceId }).(pulumi.StringPtrOutput)
}

type AutomaticResolutionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutomaticResolutionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticResolutionPropertiesResponse)(nil)).Elem()
}

func (o AutomaticResolutionPropertiesResponsePtrOutput) ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput {
	return o
}

func (o AutomaticResolutionPropertiesResponsePtrOutput) ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponsePtrOutput {
	return o
}

func (o AutomaticResolutionPropertiesResponsePtrOutput) Elem() AutomaticResolutionPropertiesResponseOutput {
	return o.ApplyT(func(v *AutomaticResolutionPropertiesResponse) AutomaticResolutionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutomaticResolutionPropertiesResponse
		return ret
	}).(AutomaticResolutionPropertiesResponseOutput)
}

func (o AutomaticResolutionPropertiesResponsePtrOutput) MoveResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomaticResolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MoveResourceId
	}).(pulumi.StringPtrOutput)
}

type AvailabilitySetResourceSettings struct {
	FaultDomain        *int              `pulumi:"faultDomain"`
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	UpdateDomain       *int              `pulumi:"updateDomain"`
}

type AvailabilitySetResourceSettingsResponse struct {
	FaultDomain        *int              `pulumi:"faultDomain"`
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	UpdateDomain       *int              `pulumi:"updateDomain"`
}

type DiskEncryptionSetResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}

type DiskEncryptionSetResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}

type Identity struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
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

func (o IdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
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

func (o IdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
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

func (o IdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
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
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
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

type JobStatusResponse struct {
	JobName     string `pulumi:"jobName"`
	JobProgress string `pulumi:"jobProgress"`
}

type JobStatusResponseOutput struct{ *pulumi.OutputState }

func (JobStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStatusResponse)(nil)).Elem()
}

func (o JobStatusResponseOutput) ToJobStatusResponseOutput() JobStatusResponseOutput {
	return o
}

func (o JobStatusResponseOutput) ToJobStatusResponseOutputWithContext(ctx context.Context) JobStatusResponseOutput {
	return o
}

func (o JobStatusResponseOutput) JobName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStatusResponse) string { return v.JobName }).(pulumi.StringOutput)
}

func (o JobStatusResponseOutput) JobProgress() pulumi.StringOutput {
	return o.ApplyT(func(v JobStatusResponse) string { return v.JobProgress }).(pulumi.StringOutput)
}

type JobStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStatusResponse)(nil)).Elem()
}

func (o JobStatusResponsePtrOutput) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return o
}

func (o JobStatusResponsePtrOutput) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return o
}

func (o JobStatusResponsePtrOutput) Elem() JobStatusResponseOutput {
	return o.ApplyT(func(v *JobStatusResponse) JobStatusResponse {
		if v != nil {
			return *v
		}
		var ret JobStatusResponse
		return ret
	}).(JobStatusResponseOutput)
}

func (o JobStatusResponsePtrOutput) JobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.JobName
	}).(pulumi.StringPtrOutput)
}

func (o JobStatusResponsePtrOutput) JobProgress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.JobProgress
	}).(pulumi.StringPtrOutput)
}

type KeyVaultResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}

type KeyVaultResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}

type LBBackendAddressPoolResourceSettings struct {
	Name *string `pulumi:"name"`
}

type LBBackendAddressPoolResourceSettingsResponse struct {
	Name *string `pulumi:"name"`
}

type LBFrontendIPConfigurationResourceSettings struct {
	Name                      *string          `pulumi:"name"`
	PrivateIpAddress          *string          `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod *string          `pulumi:"privateIpAllocationMethod"`
	Subnet                    *SubnetReference `pulumi:"subnet"`
	Zones                     *string          `pulumi:"zones"`
}

type LBFrontendIPConfigurationResourceSettingsResponse struct {
	Name                      *string                  `pulumi:"name"`
	PrivateIpAddress          *string                  `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod *string                  `pulumi:"privateIpAllocationMethod"`
	Subnet                    *SubnetReferenceResponse `pulumi:"subnet"`
	Zones                     *string                  `pulumi:"zones"`
}

type LoadBalancerBackendAddressPoolReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}

type LoadBalancerBackendAddressPoolReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}

type LoadBalancerNatRuleReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}

type LoadBalancerNatRuleReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}

type LoadBalancerResourceSettings struct {
	BackendAddressPools      []LBBackendAddressPoolResourceSettings      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations []LBFrontendIPConfigurationResourceSettings `pulumi:"frontendIPConfigurations"`
	ResourceType             string                                      `pulumi:"resourceType"`
	Sku                      *string                                     `pulumi:"sku"`
	Tags                     map[string]string                           `pulumi:"tags"`
	TargetResourceName       string                                      `pulumi:"targetResourceName"`
	Zones                    *string                                     `pulumi:"zones"`
}

type LoadBalancerResourceSettingsResponse struct {
	BackendAddressPools      []LBBackendAddressPoolResourceSettingsResponse      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations []LBFrontendIPConfigurationResourceSettingsResponse `pulumi:"frontendIPConfigurations"`
	ResourceType             string                                              `pulumi:"resourceType"`
	Sku                      *string                                             `pulumi:"sku"`
	Tags                     map[string]string                                   `pulumi:"tags"`
	TargetResourceName       string                                              `pulumi:"targetResourceName"`
	Zones                    *string                                             `pulumi:"zones"`
}

type ManualResolutionPropertiesResponse struct {
	TargetId *string `pulumi:"targetId"`
}

type ManualResolutionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManualResolutionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManualResolutionPropertiesResponse)(nil)).Elem()
}

func (o ManualResolutionPropertiesResponseOutput) ToManualResolutionPropertiesResponseOutput() ManualResolutionPropertiesResponseOutput {
	return o
}

func (o ManualResolutionPropertiesResponseOutput) ToManualResolutionPropertiesResponseOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponseOutput {
	return o
}

func (o ManualResolutionPropertiesResponseOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManualResolutionPropertiesResponse) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type ManualResolutionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ManualResolutionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManualResolutionPropertiesResponse)(nil)).Elem()
}

func (o ManualResolutionPropertiesResponsePtrOutput) ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput {
	return o
}

func (o ManualResolutionPropertiesResponsePtrOutput) ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponsePtrOutput {
	return o
}

func (o ManualResolutionPropertiesResponsePtrOutput) Elem() ManualResolutionPropertiesResponseOutput {
	return o.ApplyT(func(v *ManualResolutionPropertiesResponse) ManualResolutionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ManualResolutionPropertiesResponse
		return ret
	}).(ManualResolutionPropertiesResponseOutput)
}

func (o ManualResolutionPropertiesResponsePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManualResolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetId
	}).(pulumi.StringPtrOutput)
}

type MoveCollectionProperties struct {
	SourceRegion string `pulumi:"sourceRegion"`
	TargetRegion string `pulumi:"targetRegion"`
}





type MoveCollectionPropertiesInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesOutput() MoveCollectionPropertiesOutput
	ToMoveCollectionPropertiesOutputWithContext(context.Context) MoveCollectionPropertiesOutput
}

type MoveCollectionPropertiesArgs struct {
	SourceRegion pulumi.StringInput `pulumi:"sourceRegion"`
	TargetRegion pulumi.StringInput `pulumi:"targetRegion"`
}

func (MoveCollectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionProperties)(nil)).Elem()
}

func (i MoveCollectionPropertiesArgs) ToMoveCollectionPropertiesOutput() MoveCollectionPropertiesOutput {
	return i.ToMoveCollectionPropertiesOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesArgs) ToMoveCollectionPropertiesOutputWithContext(ctx context.Context) MoveCollectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesOutput)
}

func (i MoveCollectionPropertiesArgs) ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput {
	return i.ToMoveCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesArgs) ToMoveCollectionPropertiesPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesOutput).ToMoveCollectionPropertiesPtrOutputWithContext(ctx)
}









type MoveCollectionPropertiesPtrInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput
	ToMoveCollectionPropertiesPtrOutputWithContext(context.Context) MoveCollectionPropertiesPtrOutput
}

type moveCollectionPropertiesPtrType MoveCollectionPropertiesArgs

func MoveCollectionPropertiesPtr(v *MoveCollectionPropertiesArgs) MoveCollectionPropertiesPtrInput {
	return (*moveCollectionPropertiesPtrType)(v)
}

func (*moveCollectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionProperties)(nil)).Elem()
}

func (i *moveCollectionPropertiesPtrType) ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput {
	return i.ToMoveCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *moveCollectionPropertiesPtrType) ToMoveCollectionPropertiesPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesPtrOutput)
}

type MoveCollectionPropertiesOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionProperties)(nil)).Elem()
}

func (o MoveCollectionPropertiesOutput) ToMoveCollectionPropertiesOutput() MoveCollectionPropertiesOutput {
	return o
}

func (o MoveCollectionPropertiesOutput) ToMoveCollectionPropertiesOutputWithContext(ctx context.Context) MoveCollectionPropertiesOutput {
	return o
}

func (o MoveCollectionPropertiesOutput) ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput {
	return o.ToMoveCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (o MoveCollectionPropertiesOutput) ToMoveCollectionPropertiesPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveCollectionProperties) *MoveCollectionProperties {
		return &v
	}).(MoveCollectionPropertiesPtrOutput)
}

func (o MoveCollectionPropertiesOutput) SourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionProperties) string { return v.SourceRegion }).(pulumi.StringOutput)
}

func (o MoveCollectionPropertiesOutput) TargetRegion() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionProperties) string { return v.TargetRegion }).(pulumi.StringOutput)
}

type MoveCollectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionProperties)(nil)).Elem()
}

func (o MoveCollectionPropertiesPtrOutput) ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput {
	return o
}

func (o MoveCollectionPropertiesPtrOutput) ToMoveCollectionPropertiesPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesPtrOutput {
	return o
}

func (o MoveCollectionPropertiesPtrOutput) Elem() MoveCollectionPropertiesOutput {
	return o.ApplyT(func(v *MoveCollectionProperties) MoveCollectionProperties {
		if v != nil {
			return *v
		}
		var ret MoveCollectionProperties
		return ret
	}).(MoveCollectionPropertiesOutput)
}

func (o MoveCollectionPropertiesPtrOutput) SourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SourceRegion
	}).(pulumi.StringPtrOutput)
}

func (o MoveCollectionPropertiesPtrOutput) TargetRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.TargetRegion
	}).(pulumi.StringPtrOutput)
}

type MoveCollectionPropertiesResponse struct {
	Errors            MoveCollectionPropertiesResponseErrors `pulumi:"errors"`
	ProvisioningState string                                 `pulumi:"provisioningState"`
	SourceRegion      string                                 `pulumi:"sourceRegion"`
	TargetRegion      string                                 `pulumi:"targetRegion"`
}

type MoveCollectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionPropertiesResponse)(nil)).Elem()
}

func (o MoveCollectionPropertiesResponseOutput) ToMoveCollectionPropertiesResponseOutput() MoveCollectionPropertiesResponseOutput {
	return o
}

func (o MoveCollectionPropertiesResponseOutput) ToMoveCollectionPropertiesResponseOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseOutput {
	return o
}

func (o MoveCollectionPropertiesResponseOutput) Errors() MoveCollectionPropertiesResponseErrorsOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponse) MoveCollectionPropertiesResponseErrors { return v.Errors }).(MoveCollectionPropertiesResponseErrorsOutput)
}

func (o MoveCollectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MoveCollectionPropertiesResponseOutput) SourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponse) string { return v.SourceRegion }).(pulumi.StringOutput)
}

func (o MoveCollectionPropertiesResponseOutput) TargetRegion() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponse) string { return v.TargetRegion }).(pulumi.StringOutput)
}

type MoveCollectionPropertiesResponseErrors struct {
	Properties *MoveResourceErrorBodyResponse `pulumi:"properties"`
}

type MoveCollectionPropertiesResponseErrorsOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesResponseErrorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionPropertiesResponseErrors)(nil)).Elem()
}

func (o MoveCollectionPropertiesResponseErrorsOutput) ToMoveCollectionPropertiesResponseErrorsOutput() MoveCollectionPropertiesResponseErrorsOutput {
	return o
}

func (o MoveCollectionPropertiesResponseErrorsOutput) ToMoveCollectionPropertiesResponseErrorsOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsOutput {
	return o
}

func (o MoveCollectionPropertiesResponseErrorsOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponseErrors) *MoveResourceErrorBodyResponse { return v.Properties }).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourceDependencyOverride struct {
	Id       *string `pulumi:"id"`
	TargetId *string `pulumi:"targetId"`
}





type MoveResourceDependencyOverrideInput interface {
	pulumi.Input

	ToMoveResourceDependencyOverrideOutput() MoveResourceDependencyOverrideOutput
	ToMoveResourceDependencyOverrideOutputWithContext(context.Context) MoveResourceDependencyOverrideOutput
}

type MoveResourceDependencyOverrideArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (MoveResourceDependencyOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyOverride)(nil)).Elem()
}

func (i MoveResourceDependencyOverrideArgs) ToMoveResourceDependencyOverrideOutput() MoveResourceDependencyOverrideOutput {
	return i.ToMoveResourceDependencyOverrideOutputWithContext(context.Background())
}

func (i MoveResourceDependencyOverrideArgs) ToMoveResourceDependencyOverrideOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyOverrideOutput)
}





type MoveResourceDependencyOverrideArrayInput interface {
	pulumi.Input

	ToMoveResourceDependencyOverrideArrayOutput() MoveResourceDependencyOverrideArrayOutput
	ToMoveResourceDependencyOverrideArrayOutputWithContext(context.Context) MoveResourceDependencyOverrideArrayOutput
}

type MoveResourceDependencyOverrideArray []MoveResourceDependencyOverrideInput

func (MoveResourceDependencyOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyOverride)(nil)).Elem()
}

func (i MoveResourceDependencyOverrideArray) ToMoveResourceDependencyOverrideArrayOutput() MoveResourceDependencyOverrideArrayOutput {
	return i.ToMoveResourceDependencyOverrideArrayOutputWithContext(context.Background())
}

func (i MoveResourceDependencyOverrideArray) ToMoveResourceDependencyOverrideArrayOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyOverrideArrayOutput)
}

type MoveResourceDependencyOverrideOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyOverride)(nil)).Elem()
}

func (o MoveResourceDependencyOverrideOutput) ToMoveResourceDependencyOverrideOutput() MoveResourceDependencyOverrideOutput {
	return o
}

func (o MoveResourceDependencyOverrideOutput) ToMoveResourceDependencyOverrideOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideOutput {
	return o
}

func (o MoveResourceDependencyOverrideOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyOverride) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyOverrideOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyOverride) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type MoveResourceDependencyOverrideArrayOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyOverride)(nil)).Elem()
}

func (o MoveResourceDependencyOverrideArrayOutput) ToMoveResourceDependencyOverrideArrayOutput() MoveResourceDependencyOverrideArrayOutput {
	return o
}

func (o MoveResourceDependencyOverrideArrayOutput) ToMoveResourceDependencyOverrideArrayOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideArrayOutput {
	return o
}

func (o MoveResourceDependencyOverrideArrayOutput) Index(i pulumi.IntInput) MoveResourceDependencyOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MoveResourceDependencyOverride {
		return vs[0].([]MoveResourceDependencyOverride)[vs[1].(int)]
	}).(MoveResourceDependencyOverrideOutput)
}

type MoveResourceDependencyOverrideResponse struct {
	Id       *string `pulumi:"id"`
	TargetId *string `pulumi:"targetId"`
}

type MoveResourceDependencyOverrideResponseOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyOverrideResponse)(nil)).Elem()
}

func (o MoveResourceDependencyOverrideResponseOutput) ToMoveResourceDependencyOverrideResponseOutput() MoveResourceDependencyOverrideResponseOutput {
	return o
}

func (o MoveResourceDependencyOverrideResponseOutput) ToMoveResourceDependencyOverrideResponseOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideResponseOutput {
	return o
}

func (o MoveResourceDependencyOverrideResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyOverrideResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyOverrideResponseOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyOverrideResponse) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type MoveResourceDependencyOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyOverrideResponse)(nil)).Elem()
}

func (o MoveResourceDependencyOverrideResponseArrayOutput) ToMoveResourceDependencyOverrideResponseArrayOutput() MoveResourceDependencyOverrideResponseArrayOutput {
	return o
}

func (o MoveResourceDependencyOverrideResponseArrayOutput) ToMoveResourceDependencyOverrideResponseArrayOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideResponseArrayOutput {
	return o
}

func (o MoveResourceDependencyOverrideResponseArrayOutput) Index(i pulumi.IntInput) MoveResourceDependencyOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MoveResourceDependencyOverrideResponse {
		return vs[0].([]MoveResourceDependencyOverrideResponse)[vs[1].(int)]
	}).(MoveResourceDependencyOverrideResponseOutput)
}

type MoveResourceDependencyResponse struct {
	AutomaticResolution *AutomaticResolutionPropertiesResponse `pulumi:"automaticResolution"`
	DependencyType      *string                                `pulumi:"dependencyType"`
	Id                  *string                                `pulumi:"id"`
	IsOptional          *string                                `pulumi:"isOptional"`
	ManualResolution    *ManualResolutionPropertiesResponse    `pulumi:"manualResolution"`
	ResolutionStatus    *string                                `pulumi:"resolutionStatus"`
	ResolutionType      *string                                `pulumi:"resolutionType"`
}

type MoveResourceDependencyResponseOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyResponse)(nil)).Elem()
}

func (o MoveResourceDependencyResponseOutput) ToMoveResourceDependencyResponseOutput() MoveResourceDependencyResponseOutput {
	return o
}

func (o MoveResourceDependencyResponseOutput) ToMoveResourceDependencyResponseOutputWithContext(ctx context.Context) MoveResourceDependencyResponseOutput {
	return o
}

func (o MoveResourceDependencyResponseOutput) AutomaticResolution() AutomaticResolutionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *AutomaticResolutionPropertiesResponse {
		return v.AutomaticResolution
	}).(AutomaticResolutionPropertiesResponsePtrOutput)
}

func (o MoveResourceDependencyResponseOutput) DependencyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.DependencyType }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyResponseOutput) IsOptional() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.IsOptional }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyResponseOutput) ManualResolution() ManualResolutionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *ManualResolutionPropertiesResponse { return v.ManualResolution }).(ManualResolutionPropertiesResponsePtrOutput)
}

func (o MoveResourceDependencyResponseOutput) ResolutionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.ResolutionStatus }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyResponseOutput) ResolutionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.ResolutionType }).(pulumi.StringPtrOutput)
}

type MoveResourceDependencyResponseArrayOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyResponse)(nil)).Elem()
}

func (o MoveResourceDependencyResponseArrayOutput) ToMoveResourceDependencyResponseArrayOutput() MoveResourceDependencyResponseArrayOutput {
	return o
}

func (o MoveResourceDependencyResponseArrayOutput) ToMoveResourceDependencyResponseArrayOutputWithContext(ctx context.Context) MoveResourceDependencyResponseArrayOutput {
	return o
}

func (o MoveResourceDependencyResponseArrayOutput) Index(i pulumi.IntInput) MoveResourceDependencyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MoveResourceDependencyResponse {
		return vs[0].([]MoveResourceDependencyResponse)[vs[1].(int)]
	}).(MoveResourceDependencyResponseOutput)
}

type MoveResourceErrorBodyResponse struct {
	Code    string                          `pulumi:"code"`
	Details []MoveResourceErrorBodyResponse `pulumi:"details"`
	Message string                          `pulumi:"message"`
	Target  string                          `pulumi:"target"`
}

type MoveResourceErrorBodyResponseOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorBodyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (o MoveResourceErrorBodyResponseOutput) ToMoveResourceErrorBodyResponseOutput() MoveResourceErrorBodyResponseOutput {
	return o
}

func (o MoveResourceErrorBodyResponseOutput) ToMoveResourceErrorBodyResponseOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponseOutput {
	return o
}

func (o MoveResourceErrorBodyResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourceErrorBodyResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o MoveResourceErrorBodyResponseOutput) Details() MoveResourceErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v MoveResourceErrorBodyResponse) []MoveResourceErrorBodyResponse { return v.Details }).(MoveResourceErrorBodyResponseArrayOutput)
}

func (o MoveResourceErrorBodyResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourceErrorBodyResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o MoveResourceErrorBodyResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourceErrorBodyResponse) string { return v.Target }).(pulumi.StringOutput)
}

type MoveResourceErrorBodyResponsePtrOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorBodyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (o MoveResourceErrorBodyResponsePtrOutput) ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput {
	return o
}

func (o MoveResourceErrorBodyResponsePtrOutput) ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponsePtrOutput {
	return o
}

func (o MoveResourceErrorBodyResponsePtrOutput) Elem() MoveResourceErrorBodyResponseOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) MoveResourceErrorBodyResponse {
		if v != nil {
			return *v
		}
		var ret MoveResourceErrorBodyResponse
		return ret
	}).(MoveResourceErrorBodyResponseOutput)
}

func (o MoveResourceErrorBodyResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourceErrorBodyResponsePtrOutput) Details() MoveResourceErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) []MoveResourceErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(MoveResourceErrorBodyResponseArrayOutput)
}

func (o MoveResourceErrorBodyResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourceErrorBodyResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
}

type MoveResourceErrorBodyResponseArrayOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorBodyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (o MoveResourceErrorBodyResponseArrayOutput) ToMoveResourceErrorBodyResponseArrayOutput() MoveResourceErrorBodyResponseArrayOutput {
	return o
}

func (o MoveResourceErrorBodyResponseArrayOutput) ToMoveResourceErrorBodyResponseArrayOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponseArrayOutput {
	return o
}

func (o MoveResourceErrorBodyResponseArrayOutput) Index(i pulumi.IntInput) MoveResourceErrorBodyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MoveResourceErrorBodyResponse {
		return vs[0].([]MoveResourceErrorBodyResponse)[vs[1].(int)]
	}).(MoveResourceErrorBodyResponseOutput)
}

type MoveResourceErrorResponse struct {
	Properties *MoveResourceErrorBodyResponse `pulumi:"properties"`
}

type MoveResourceErrorResponseOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceErrorResponse)(nil)).Elem()
}

func (o MoveResourceErrorResponseOutput) ToMoveResourceErrorResponseOutput() MoveResourceErrorResponseOutput {
	return o
}

func (o MoveResourceErrorResponseOutput) ToMoveResourceErrorResponseOutputWithContext(ctx context.Context) MoveResourceErrorResponseOutput {
	return o
}

func (o MoveResourceErrorResponseOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v MoveResourceErrorResponse) *MoveResourceErrorBodyResponse { return v.Properties }).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourceErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceErrorResponse)(nil)).Elem()
}

func (o MoveResourceErrorResponsePtrOutput) ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput {
	return o
}

func (o MoveResourceErrorResponsePtrOutput) ToMoveResourceErrorResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorResponsePtrOutput {
	return o
}

func (o MoveResourceErrorResponsePtrOutput) Elem() MoveResourceErrorResponseOutput {
	return o.ApplyT(func(v *MoveResourceErrorResponse) MoveResourceErrorResponse {
		if v != nil {
			return *v
		}
		var ret MoveResourceErrorResponse
		return ret
	}).(MoveResourceErrorResponseOutput)
}

func (o MoveResourceErrorResponsePtrOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v *MoveResourceErrorResponse) *MoveResourceErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourceProperties struct {
	DependsOnOverrides []MoveResourceDependencyOverride `pulumi:"dependsOnOverrides"`
	ExistingTargetId   *string                          `pulumi:"existingTargetId"`
	ResourceSettings   interface{}                      `pulumi:"resourceSettings"`
	SourceId           string                           `pulumi:"sourceId"`
}





type MoveResourcePropertiesInput interface {
	pulumi.Input

	ToMoveResourcePropertiesOutput() MoveResourcePropertiesOutput
	ToMoveResourcePropertiesOutputWithContext(context.Context) MoveResourcePropertiesOutput
}

type MoveResourcePropertiesArgs struct {
	DependsOnOverrides MoveResourceDependencyOverrideArrayInput `pulumi:"dependsOnOverrides"`
	ExistingTargetId   pulumi.StringPtrInput                    `pulumi:"existingTargetId"`
	ResourceSettings   pulumi.Input                             `pulumi:"resourceSettings"`
	SourceId           pulumi.StringInput                       `pulumi:"sourceId"`
}

func (MoveResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceProperties)(nil)).Elem()
}

func (i MoveResourcePropertiesArgs) ToMoveResourcePropertiesOutput() MoveResourcePropertiesOutput {
	return i.ToMoveResourcePropertiesOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesArgs) ToMoveResourcePropertiesOutputWithContext(ctx context.Context) MoveResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesOutput)
}

func (i MoveResourcePropertiesArgs) ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput {
	return i.ToMoveResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesArgs) ToMoveResourcePropertiesPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesOutput).ToMoveResourcePropertiesPtrOutputWithContext(ctx)
}









type MoveResourcePropertiesPtrInput interface {
	pulumi.Input

	ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput
	ToMoveResourcePropertiesPtrOutputWithContext(context.Context) MoveResourcePropertiesPtrOutput
}

type moveResourcePropertiesPtrType MoveResourcePropertiesArgs

func MoveResourcePropertiesPtr(v *MoveResourcePropertiesArgs) MoveResourcePropertiesPtrInput {
	return (*moveResourcePropertiesPtrType)(v)
}

func (*moveResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceProperties)(nil)).Elem()
}

func (i *moveResourcePropertiesPtrType) ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput {
	return i.ToMoveResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *moveResourcePropertiesPtrType) ToMoveResourcePropertiesPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesPtrOutput)
}

type MoveResourcePropertiesOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceProperties)(nil)).Elem()
}

func (o MoveResourcePropertiesOutput) ToMoveResourcePropertiesOutput() MoveResourcePropertiesOutput {
	return o
}

func (o MoveResourcePropertiesOutput) ToMoveResourcePropertiesOutputWithContext(ctx context.Context) MoveResourcePropertiesOutput {
	return o
}

func (o MoveResourcePropertiesOutput) ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput {
	return o.ToMoveResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o MoveResourcePropertiesOutput) ToMoveResourcePropertiesPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourceProperties) *MoveResourceProperties {
		return &v
	}).(MoveResourcePropertiesPtrOutput)
}

func (o MoveResourcePropertiesOutput) DependsOnOverrides() MoveResourceDependencyOverrideArrayOutput {
	return o.ApplyT(func(v MoveResourceProperties) []MoveResourceDependencyOverride { return v.DependsOnOverrides }).(MoveResourceDependencyOverrideArrayOutput)
}

func (o MoveResourcePropertiesOutput) ExistingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceProperties) *string { return v.ExistingTargetId }).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesOutput) ResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MoveResourceProperties) interface{} { return v.ResourceSettings }).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourceProperties) string { return v.SourceId }).(pulumi.StringOutput)
}

type MoveResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceProperties)(nil)).Elem()
}

func (o MoveResourcePropertiesPtrOutput) ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput {
	return o
}

func (o MoveResourcePropertiesPtrOutput) ToMoveResourcePropertiesPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesPtrOutput {
	return o
}

func (o MoveResourcePropertiesPtrOutput) Elem() MoveResourcePropertiesOutput {
	return o.ApplyT(func(v *MoveResourceProperties) MoveResourceProperties {
		if v != nil {
			return *v
		}
		var ret MoveResourceProperties
		return ret
	}).(MoveResourcePropertiesOutput)
}

func (o MoveResourcePropertiesPtrOutput) DependsOnOverrides() MoveResourceDependencyOverrideArrayOutput {
	return o.ApplyT(func(v *MoveResourceProperties) []MoveResourceDependencyOverride {
		if v == nil {
			return nil
		}
		return v.DependsOnOverrides
	}).(MoveResourceDependencyOverrideArrayOutput)
}

func (o MoveResourcePropertiesPtrOutput) ExistingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ExistingTargetId
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesPtrOutput) ResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MoveResourceProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.ResourceSettings
	}).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesPtrOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SourceId
	}).(pulumi.StringPtrOutput)
}

type MoveResourcePropertiesResponse struct {
	DependsOn              []MoveResourceDependencyResponse         `pulumi:"dependsOn"`
	DependsOnOverrides     []MoveResourceDependencyOverrideResponse `pulumi:"dependsOnOverrides"`
	Errors                 MoveResourcePropertiesResponseErrors     `pulumi:"errors"`
	ExistingTargetId       *string                                  `pulumi:"existingTargetId"`
	IsResolveRequired      bool                                     `pulumi:"isResolveRequired"`
	MoveStatus             MoveResourcePropertiesResponseMoveStatus `pulumi:"moveStatus"`
	ProvisioningState      string                                   `pulumi:"provisioningState"`
	ResourceSettings       interface{}                              `pulumi:"resourceSettings"`
	SourceId               string                                   `pulumi:"sourceId"`
	SourceResourceSettings interface{}                              `pulumi:"sourceResourceSettings"`
	TargetId               string                                   `pulumi:"targetId"`
}

type MoveResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponse)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseOutput) ToMoveResourcePropertiesResponseOutput() MoveResourcePropertiesResponseOutput {
	return o
}

func (o MoveResourcePropertiesResponseOutput) ToMoveResourcePropertiesResponseOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseOutput {
	return o
}

func (o MoveResourcePropertiesResponseOutput) DependsOn() MoveResourceDependencyResponseArrayOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) []MoveResourceDependencyResponse { return v.DependsOn }).(MoveResourceDependencyResponseArrayOutput)
}

func (o MoveResourcePropertiesResponseOutput) DependsOnOverrides() MoveResourceDependencyOverrideResponseArrayOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) []MoveResourceDependencyOverrideResponse {
		return v.DependsOnOverrides
	}).(MoveResourceDependencyOverrideResponseArrayOutput)
}

func (o MoveResourcePropertiesResponseOutput) Errors() MoveResourcePropertiesResponseErrorsOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) MoveResourcePropertiesResponseErrors { return v.Errors }).(MoveResourcePropertiesResponseErrorsOutput)
}

func (o MoveResourcePropertiesResponseOutput) ExistingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) *string { return v.ExistingTargetId }).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesResponseOutput) IsResolveRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) bool { return v.IsResolveRequired }).(pulumi.BoolOutput)
}

func (o MoveResourcePropertiesResponseOutput) MoveStatus() MoveResourcePropertiesResponseMoveStatusOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) MoveResourcePropertiesResponseMoveStatus { return v.MoveStatus }).(MoveResourcePropertiesResponseMoveStatusOutput)
}

func (o MoveResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MoveResourcePropertiesResponseOutput) ResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) interface{} { return v.ResourceSettings }).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesResponseOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o MoveResourcePropertiesResponseOutput) SourceResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) interface{} { return v.SourceResourceSettings }).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesResponseOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) string { return v.TargetId }).(pulumi.StringOutput)
}

type MoveResourcePropertiesResponseErrors struct {
	Properties *MoveResourceErrorBodyResponse `pulumi:"properties"`
}

type MoveResourcePropertiesResponseErrorsOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseErrorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponseErrors)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseErrorsOutput) ToMoveResourcePropertiesResponseErrorsOutput() MoveResourcePropertiesResponseErrorsOutput {
	return o
}

func (o MoveResourcePropertiesResponseErrorsOutput) ToMoveResourcePropertiesResponseErrorsOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsOutput {
	return o
}

func (o MoveResourcePropertiesResponseErrorsOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseErrors) *MoveResourceErrorBodyResponse { return v.Properties }).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourcePropertiesResponseMoveStatus struct {
	Errors    *MoveResourceErrorResponse `pulumi:"errors"`
	JobStatus *JobStatusResponse         `pulumi:"jobStatus"`
	MoveState string                     `pulumi:"moveState"`
}

type MoveResourcePropertiesResponseMoveStatusOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseMoveStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponseMoveStatus)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) ToMoveResourcePropertiesResponseMoveStatusOutput() MoveResourcePropertiesResponseMoveStatusOutput {
	return o
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) ToMoveResourcePropertiesResponseMoveStatusOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusOutput {
	return o
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) Errors() MoveResourceErrorResponsePtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseMoveStatus) *MoveResourceErrorResponse { return v.Errors }).(MoveResourceErrorResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) JobStatus() JobStatusResponsePtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseMoveStatus) *JobStatusResponse { return v.JobStatus }).(JobStatusResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) MoveState() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseMoveStatus) string { return v.MoveState }).(pulumi.StringOutput)
}

type NetworkInterfaceResourceSettings struct {
	EnableAcceleratedNetworking *bool                                `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            []NicIpConfigurationResourceSettings `pulumi:"ipConfigurations"`
	ResourceType                string                               `pulumi:"resourceType"`
	Tags                        map[string]string                    `pulumi:"tags"`
	TargetResourceName          string                               `pulumi:"targetResourceName"`
}

type NetworkInterfaceResourceSettingsResponse struct {
	EnableAcceleratedNetworking *bool                                        `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            []NicIpConfigurationResourceSettingsResponse `pulumi:"ipConfigurations"`
	ResourceType                string                                       `pulumi:"resourceType"`
	Tags                        map[string]string                            `pulumi:"tags"`
	TargetResourceName          string                                       `pulumi:"targetResourceName"`
}

type NetworkSecurityGroupResourceSettings struct {
	ResourceType       string            `pulumi:"resourceType"`
	SecurityRules      []NsgSecurityRule `pulumi:"securityRules"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
}

type NetworkSecurityGroupResourceSettingsResponse struct {
	ResourceType       string                    `pulumi:"resourceType"`
	SecurityRules      []NsgSecurityRuleResponse `pulumi:"securityRules"`
	Tags               map[string]string         `pulumi:"tags"`
	TargetResourceName string                    `pulumi:"targetResourceName"`
}

type NicIpConfigurationResourceSettings struct {
	LoadBalancerBackendAddressPools []LoadBalancerBackendAddressPoolReference `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerNatRules            []LoadBalancerNatRuleReference            `pulumi:"loadBalancerNatRules"`
	Name                            *string                                   `pulumi:"name"`
	Primary                         *bool                                     `pulumi:"primary"`
	PrivateIpAddress                *string                                   `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod       *string                                   `pulumi:"privateIpAllocationMethod"`
	PublicIp                        *PublicIpReference                        `pulumi:"publicIp"`
	Subnet                          *SubnetReference                          `pulumi:"subnet"`
}

type NicIpConfigurationResourceSettingsResponse struct {
	LoadBalancerBackendAddressPools []LoadBalancerBackendAddressPoolReferenceResponse `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerNatRules            []LoadBalancerNatRuleReferenceResponse            `pulumi:"loadBalancerNatRules"`
	Name                            *string                                           `pulumi:"name"`
	Primary                         *bool                                             `pulumi:"primary"`
	PrivateIpAddress                *string                                           `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod       *string                                           `pulumi:"privateIpAllocationMethod"`
	PublicIp                        *PublicIpReferenceResponse                        `pulumi:"publicIp"`
	Subnet                          *SubnetReferenceResponse                          `pulumi:"subnet"`
}

type NsgReference struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}

type NsgReferenceResponse struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}

type NsgSecurityRule struct {
	Access                   *string `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                *string `pulumi:"direction"`
	Name                     *string `pulumi:"name"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 *string `pulumi:"protocol"`
	SourceAddressPrefix      *string `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}

type NsgSecurityRuleResponse struct {
	Access                   *string `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                *string `pulumi:"direction"`
	Name                     *string `pulumi:"name"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 *string `pulumi:"protocol"`
	SourceAddressPrefix      *string `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}

type PublicIPAddressResourceSettings struct {
	DomainNameLabel          *string           `pulumi:"domainNameLabel"`
	Fqdn                     *string           `pulumi:"fqdn"`
	PublicIpAllocationMethod *string           `pulumi:"publicIpAllocationMethod"`
	ResourceType             string            `pulumi:"resourceType"`
	Sku                      *string           `pulumi:"sku"`
	Tags                     map[string]string `pulumi:"tags"`
	TargetResourceName       string            `pulumi:"targetResourceName"`
	Zones                    *string           `pulumi:"zones"`
}

type PublicIPAddressResourceSettingsResponse struct {
	DomainNameLabel          *string           `pulumi:"domainNameLabel"`
	Fqdn                     *string           `pulumi:"fqdn"`
	PublicIpAllocationMethod *string           `pulumi:"publicIpAllocationMethod"`
	ResourceType             string            `pulumi:"resourceType"`
	Sku                      *string           `pulumi:"sku"`
	Tags                     map[string]string `pulumi:"tags"`
	TargetResourceName       string            `pulumi:"targetResourceName"`
	Zones                    *string           `pulumi:"zones"`
}

type PublicIpReference struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}

type PublicIpReferenceResponse struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}

type ResourceGroupResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}

type ResourceGroupResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}

type SqlDatabaseResourceSettings struct {
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	ZoneRedundant      *string           `pulumi:"zoneRedundant"`
}

type SqlDatabaseResourceSettingsResponse struct {
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	ZoneRedundant      *string           `pulumi:"zoneRedundant"`
}

type SqlElasticPoolResourceSettings struct {
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	ZoneRedundant      *string           `pulumi:"zoneRedundant"`
}

type SqlElasticPoolResourceSettingsResponse struct {
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	ZoneRedundant      *string           `pulumi:"zoneRedundant"`
}

type SqlServerResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}

type SqlServerResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}

type SubnetReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}

type SubnetReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}

type SubnetResourceSettings struct {
	AddressPrefix        *string       `pulumi:"addressPrefix"`
	Name                 *string       `pulumi:"name"`
	NetworkSecurityGroup *NsgReference `pulumi:"networkSecurityGroup"`
}

type SubnetResourceSettingsResponse struct {
	AddressPrefix        *string               `pulumi:"addressPrefix"`
	Name                 *string               `pulumi:"name"`
	NetworkSecurityGroup *NsgReferenceResponse `pulumi:"networkSecurityGroup"`
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

type VirtualMachineResourceSettings struct {
	ResourceType            string            `pulumi:"resourceType"`
	Tags                    map[string]string `pulumi:"tags"`
	TargetAvailabilitySetId *string           `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  *string           `pulumi:"targetAvailabilityZone"`
	TargetResourceName      string            `pulumi:"targetResourceName"`
	TargetVmSize            *string           `pulumi:"targetVmSize"`
	UserManagedIdentities   []string          `pulumi:"userManagedIdentities"`
}

type VirtualMachineResourceSettingsResponse struct {
	ResourceType            string            `pulumi:"resourceType"`
	Tags                    map[string]string `pulumi:"tags"`
	TargetAvailabilitySetId *string           `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  *string           `pulumi:"targetAvailabilityZone"`
	TargetResourceName      string            `pulumi:"targetResourceName"`
	TargetVmSize            *string           `pulumi:"targetVmSize"`
	UserManagedIdentities   []string          `pulumi:"userManagedIdentities"`
}

type VirtualNetworkResourceSettings struct {
	AddressSpace         []string                 `pulumi:"addressSpace"`
	DnsServers           []string                 `pulumi:"dnsServers"`
	EnableDdosProtection *bool                    `pulumi:"enableDdosProtection"`
	ResourceType         string                   `pulumi:"resourceType"`
	Subnets              []SubnetResourceSettings `pulumi:"subnets"`
	Tags                 map[string]string        `pulumi:"tags"`
	TargetResourceName   string                   `pulumi:"targetResourceName"`
}

type VirtualNetworkResourceSettingsResponse struct {
	AddressSpace         []string                         `pulumi:"addressSpace"`
	DnsServers           []string                         `pulumi:"dnsServers"`
	EnableDdosProtection *bool                            `pulumi:"enableDdosProtection"`
	ResourceType         string                           `pulumi:"resourceType"`
	Subnets              []SubnetResourceSettingsResponse `pulumi:"subnets"`
	Tags                 map[string]string                `pulumi:"tags"`
	TargetResourceName   string                           `pulumi:"targetResourceName"`
}

func init() {
	pulumi.RegisterOutputType(AutomaticResolutionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutomaticResolutionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStatusResponseOutput{})
	pulumi.RegisterOutputType(JobStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ManualResolutionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManualResolutionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponseErrorsOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyOverrideOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyOverrideArrayOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyOverrideResponseOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyResponseOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorBodyResponseOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorBodyResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorBodyResponseArrayOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorResponseOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseErrorsOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseMoveStatusOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
