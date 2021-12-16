


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutomaticResolutionPropertiesResponse struct {
	MoveResourceId *string `pulumi:"moveResourceId"`
}





type AutomaticResolutionPropertiesResponseInput interface {
	pulumi.Input

	ToAutomaticResolutionPropertiesResponseOutput() AutomaticResolutionPropertiesResponseOutput
	ToAutomaticResolutionPropertiesResponseOutputWithContext(context.Context) AutomaticResolutionPropertiesResponseOutput
}

type AutomaticResolutionPropertiesResponseArgs struct {
	MoveResourceId pulumi.StringPtrInput `pulumi:"moveResourceId"`
}

func (AutomaticResolutionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticResolutionPropertiesResponse)(nil)).Elem()
}

func (i AutomaticResolutionPropertiesResponseArgs) ToAutomaticResolutionPropertiesResponseOutput() AutomaticResolutionPropertiesResponseOutput {
	return i.ToAutomaticResolutionPropertiesResponseOutputWithContext(context.Background())
}

func (i AutomaticResolutionPropertiesResponseArgs) ToAutomaticResolutionPropertiesResponseOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticResolutionPropertiesResponseOutput)
}

func (i AutomaticResolutionPropertiesResponseArgs) ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput {
	return i.ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AutomaticResolutionPropertiesResponseArgs) ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticResolutionPropertiesResponseOutput).ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx)
}









type AutomaticResolutionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput
	ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(context.Context) AutomaticResolutionPropertiesResponsePtrOutput
}

type automaticResolutionPropertiesResponsePtrType AutomaticResolutionPropertiesResponseArgs

func AutomaticResolutionPropertiesResponsePtr(v *AutomaticResolutionPropertiesResponseArgs) AutomaticResolutionPropertiesResponsePtrInput {
	return (*automaticResolutionPropertiesResponsePtrType)(v)
}

func (*automaticResolutionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticResolutionPropertiesResponse)(nil)).Elem()
}

func (i *automaticResolutionPropertiesResponsePtrType) ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput {
	return i.ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *automaticResolutionPropertiesResponsePtrType) ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticResolutionPropertiesResponsePtrOutput)
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

func (o AutomaticResolutionPropertiesResponseOutput) ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput {
	return o.ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AutomaticResolutionPropertiesResponseOutput) ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomaticResolutionPropertiesResponse) *AutomaticResolutionPropertiesResponse {
		return &v
	}).(AutomaticResolutionPropertiesResponsePtrOutput)
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





type AvailabilitySetResourceSettingsInput interface {
	pulumi.Input

	ToAvailabilitySetResourceSettingsOutput() AvailabilitySetResourceSettingsOutput
	ToAvailabilitySetResourceSettingsOutputWithContext(context.Context) AvailabilitySetResourceSettingsOutput
}

type AvailabilitySetResourceSettingsArgs struct {
	FaultDomain        pulumi.IntPtrInput    `pulumi:"faultDomain"`
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	Tags               pulumi.StringMapInput `pulumi:"tags"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	UpdateDomain       pulumi.IntPtrInput    `pulumi:"updateDomain"`
}

func (AvailabilitySetResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySetResourceSettings)(nil)).Elem()
}

func (i AvailabilitySetResourceSettingsArgs) ToAvailabilitySetResourceSettingsOutput() AvailabilitySetResourceSettingsOutput {
	return i.ToAvailabilitySetResourceSettingsOutputWithContext(context.Background())
}

func (i AvailabilitySetResourceSettingsArgs) ToAvailabilitySetResourceSettingsOutputWithContext(ctx context.Context) AvailabilitySetResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetResourceSettingsOutput)
}

type AvailabilitySetResourceSettingsOutput struct{ *pulumi.OutputState }

func (AvailabilitySetResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySetResourceSettings)(nil)).Elem()
}

func (o AvailabilitySetResourceSettingsOutput) ToAvailabilitySetResourceSettingsOutput() AvailabilitySetResourceSettingsOutput {
	return o
}

func (o AvailabilitySetResourceSettingsOutput) ToAvailabilitySetResourceSettingsOutputWithContext(ctx context.Context) AvailabilitySetResourceSettingsOutput {
	return o
}

func (o AvailabilitySetResourceSettingsOutput) FaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) *int { return v.FaultDomain }).(pulumi.IntPtrOutput)
}

func (o AvailabilitySetResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o AvailabilitySetResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AvailabilitySetResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o AvailabilitySetResourceSettingsOutput) UpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) *int { return v.UpdateDomain }).(pulumi.IntPtrOutput)
}

type AvailabilitySetResourceSettingsResponse struct {
	FaultDomain        *int              `pulumi:"faultDomain"`
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	UpdateDomain       *int              `pulumi:"updateDomain"`
}





type AvailabilitySetResourceSettingsResponseInput interface {
	pulumi.Input

	ToAvailabilitySetResourceSettingsResponseOutput() AvailabilitySetResourceSettingsResponseOutput
	ToAvailabilitySetResourceSettingsResponseOutputWithContext(context.Context) AvailabilitySetResourceSettingsResponseOutput
}

type AvailabilitySetResourceSettingsResponseArgs struct {
	FaultDomain        pulumi.IntPtrInput    `pulumi:"faultDomain"`
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	Tags               pulumi.StringMapInput `pulumi:"tags"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	UpdateDomain       pulumi.IntPtrInput    `pulumi:"updateDomain"`
}

func (AvailabilitySetResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySetResourceSettingsResponse)(nil)).Elem()
}

func (i AvailabilitySetResourceSettingsResponseArgs) ToAvailabilitySetResourceSettingsResponseOutput() AvailabilitySetResourceSettingsResponseOutput {
	return i.ToAvailabilitySetResourceSettingsResponseOutputWithContext(context.Background())
}

func (i AvailabilitySetResourceSettingsResponseArgs) ToAvailabilitySetResourceSettingsResponseOutputWithContext(ctx context.Context) AvailabilitySetResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetResourceSettingsResponseOutput)
}

type AvailabilitySetResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (AvailabilitySetResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySetResourceSettingsResponse)(nil)).Elem()
}

func (o AvailabilitySetResourceSettingsResponseOutput) ToAvailabilitySetResourceSettingsResponseOutput() AvailabilitySetResourceSettingsResponseOutput {
	return o
}

func (o AvailabilitySetResourceSettingsResponseOutput) ToAvailabilitySetResourceSettingsResponseOutputWithContext(ctx context.Context) AvailabilitySetResourceSettingsResponseOutput {
	return o
}

func (o AvailabilitySetResourceSettingsResponseOutput) FaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) *int { return v.FaultDomain }).(pulumi.IntPtrOutput)
}

func (o AvailabilitySetResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o AvailabilitySetResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AvailabilitySetResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o AvailabilitySetResourceSettingsResponseOutput) UpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) *int { return v.UpdateDomain }).(pulumi.IntPtrOutput)
}

type DiskEncryptionSetResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type DiskEncryptionSetResourceSettingsInput interface {
	pulumi.Input

	ToDiskEncryptionSetResourceSettingsOutput() DiskEncryptionSetResourceSettingsOutput
	ToDiskEncryptionSetResourceSettingsOutputWithContext(context.Context) DiskEncryptionSetResourceSettingsOutput
}

type DiskEncryptionSetResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (DiskEncryptionSetResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetResourceSettings)(nil)).Elem()
}

func (i DiskEncryptionSetResourceSettingsArgs) ToDiskEncryptionSetResourceSettingsOutput() DiskEncryptionSetResourceSettingsOutput {
	return i.ToDiskEncryptionSetResourceSettingsOutputWithContext(context.Background())
}

func (i DiskEncryptionSetResourceSettingsArgs) ToDiskEncryptionSetResourceSettingsOutputWithContext(ctx context.Context) DiskEncryptionSetResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetResourceSettingsOutput)
}

type DiskEncryptionSetResourceSettingsOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetResourceSettings)(nil)).Elem()
}

func (o DiskEncryptionSetResourceSettingsOutput) ToDiskEncryptionSetResourceSettingsOutput() DiskEncryptionSetResourceSettingsOutput {
	return o
}

func (o DiskEncryptionSetResourceSettingsOutput) ToDiskEncryptionSetResourceSettingsOutputWithContext(ctx context.Context) DiskEncryptionSetResourceSettingsOutput {
	return o
}

func (o DiskEncryptionSetResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v DiskEncryptionSetResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o DiskEncryptionSetResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v DiskEncryptionSetResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type DiskEncryptionSetResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type DiskEncryptionSetResourceSettingsResponseInput interface {
	pulumi.Input

	ToDiskEncryptionSetResourceSettingsResponseOutput() DiskEncryptionSetResourceSettingsResponseOutput
	ToDiskEncryptionSetResourceSettingsResponseOutputWithContext(context.Context) DiskEncryptionSetResourceSettingsResponseOutput
}

type DiskEncryptionSetResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (DiskEncryptionSetResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetResourceSettingsResponse)(nil)).Elem()
}

func (i DiskEncryptionSetResourceSettingsResponseArgs) ToDiskEncryptionSetResourceSettingsResponseOutput() DiskEncryptionSetResourceSettingsResponseOutput {
	return i.ToDiskEncryptionSetResourceSettingsResponseOutputWithContext(context.Background())
}

func (i DiskEncryptionSetResourceSettingsResponseArgs) ToDiskEncryptionSetResourceSettingsResponseOutputWithContext(ctx context.Context) DiskEncryptionSetResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetResourceSettingsResponseOutput)
}

type DiskEncryptionSetResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetResourceSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSetResourceSettingsResponseOutput) ToDiskEncryptionSetResourceSettingsResponseOutput() DiskEncryptionSetResourceSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSetResourceSettingsResponseOutput) ToDiskEncryptionSetResourceSettingsResponseOutputWithContext(ctx context.Context) DiskEncryptionSetResourceSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSetResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v DiskEncryptionSetResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o DiskEncryptionSetResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v DiskEncryptionSetResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
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





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
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





type JobStatusResponseInput interface {
	pulumi.Input

	ToJobStatusResponseOutput() JobStatusResponseOutput
	ToJobStatusResponseOutputWithContext(context.Context) JobStatusResponseOutput
}

type JobStatusResponseArgs struct {
	JobName     pulumi.StringInput `pulumi:"jobName"`
	JobProgress pulumi.StringInput `pulumi:"jobProgress"`
}

func (JobStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStatusResponse)(nil)).Elem()
}

func (i JobStatusResponseArgs) ToJobStatusResponseOutput() JobStatusResponseOutput {
	return i.ToJobStatusResponseOutputWithContext(context.Background())
}

func (i JobStatusResponseArgs) ToJobStatusResponseOutputWithContext(ctx context.Context) JobStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponseOutput)
}

func (i JobStatusResponseArgs) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return i.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (i JobStatusResponseArgs) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponseOutput).ToJobStatusResponsePtrOutputWithContext(ctx)
}









type JobStatusResponsePtrInput interface {
	pulumi.Input

	ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput
	ToJobStatusResponsePtrOutputWithContext(context.Context) JobStatusResponsePtrOutput
}

type jobStatusResponsePtrType JobStatusResponseArgs

func JobStatusResponsePtr(v *JobStatusResponseArgs) JobStatusResponsePtrInput {
	return (*jobStatusResponsePtrType)(v)
}

func (*jobStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStatusResponse)(nil)).Elem()
}

func (i *jobStatusResponsePtrType) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return i.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (i *jobStatusResponsePtrType) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponsePtrOutput)
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

func (o JobStatusResponseOutput) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return o.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (o JobStatusResponseOutput) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStatusResponse) *JobStatusResponse {
		return &v
	}).(JobStatusResponsePtrOutput)
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





type KeyVaultResourceSettingsInput interface {
	pulumi.Input

	ToKeyVaultResourceSettingsOutput() KeyVaultResourceSettingsOutput
	ToKeyVaultResourceSettingsOutputWithContext(context.Context) KeyVaultResourceSettingsOutput
}

type KeyVaultResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (KeyVaultResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultResourceSettings)(nil)).Elem()
}

func (i KeyVaultResourceSettingsArgs) ToKeyVaultResourceSettingsOutput() KeyVaultResourceSettingsOutput {
	return i.ToKeyVaultResourceSettingsOutputWithContext(context.Background())
}

func (i KeyVaultResourceSettingsArgs) ToKeyVaultResourceSettingsOutputWithContext(ctx context.Context) KeyVaultResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultResourceSettingsOutput)
}

type KeyVaultResourceSettingsOutput struct{ *pulumi.OutputState }

func (KeyVaultResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultResourceSettings)(nil)).Elem()
}

func (o KeyVaultResourceSettingsOutput) ToKeyVaultResourceSettingsOutput() KeyVaultResourceSettingsOutput {
	return o
}

func (o KeyVaultResourceSettingsOutput) ToKeyVaultResourceSettingsOutputWithContext(ctx context.Context) KeyVaultResourceSettingsOutput {
	return o
}

func (o KeyVaultResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o KeyVaultResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type KeyVaultResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type KeyVaultResourceSettingsResponseInput interface {
	pulumi.Input

	ToKeyVaultResourceSettingsResponseOutput() KeyVaultResourceSettingsResponseOutput
	ToKeyVaultResourceSettingsResponseOutputWithContext(context.Context) KeyVaultResourceSettingsResponseOutput
}

type KeyVaultResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (KeyVaultResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultResourceSettingsResponse)(nil)).Elem()
}

func (i KeyVaultResourceSettingsResponseArgs) ToKeyVaultResourceSettingsResponseOutput() KeyVaultResourceSettingsResponseOutput {
	return i.ToKeyVaultResourceSettingsResponseOutputWithContext(context.Background())
}

func (i KeyVaultResourceSettingsResponseArgs) ToKeyVaultResourceSettingsResponseOutputWithContext(ctx context.Context) KeyVaultResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultResourceSettingsResponseOutput)
}

type KeyVaultResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultResourceSettingsResponse)(nil)).Elem()
}

func (o KeyVaultResourceSettingsResponseOutput) ToKeyVaultResourceSettingsResponseOutput() KeyVaultResourceSettingsResponseOutput {
	return o
}

func (o KeyVaultResourceSettingsResponseOutput) ToKeyVaultResourceSettingsResponseOutputWithContext(ctx context.Context) KeyVaultResourceSettingsResponseOutput {
	return o
}

func (o KeyVaultResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o KeyVaultResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type LBBackendAddressPoolResourceSettings struct {
	Name *string `pulumi:"name"`
}





type LBBackendAddressPoolResourceSettingsInput interface {
	pulumi.Input

	ToLBBackendAddressPoolResourceSettingsOutput() LBBackendAddressPoolResourceSettingsOutput
	ToLBBackendAddressPoolResourceSettingsOutputWithContext(context.Context) LBBackendAddressPoolResourceSettingsOutput
}

type LBBackendAddressPoolResourceSettingsArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LBBackendAddressPoolResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LBBackendAddressPoolResourceSettings)(nil)).Elem()
}

func (i LBBackendAddressPoolResourceSettingsArgs) ToLBBackendAddressPoolResourceSettingsOutput() LBBackendAddressPoolResourceSettingsOutput {
	return i.ToLBBackendAddressPoolResourceSettingsOutputWithContext(context.Background())
}

func (i LBBackendAddressPoolResourceSettingsArgs) ToLBBackendAddressPoolResourceSettingsOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBBackendAddressPoolResourceSettingsOutput)
}





type LBBackendAddressPoolResourceSettingsArrayInput interface {
	pulumi.Input

	ToLBBackendAddressPoolResourceSettingsArrayOutput() LBBackendAddressPoolResourceSettingsArrayOutput
	ToLBBackendAddressPoolResourceSettingsArrayOutputWithContext(context.Context) LBBackendAddressPoolResourceSettingsArrayOutput
}

type LBBackendAddressPoolResourceSettingsArray []LBBackendAddressPoolResourceSettingsInput

func (LBBackendAddressPoolResourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBBackendAddressPoolResourceSettings)(nil)).Elem()
}

func (i LBBackendAddressPoolResourceSettingsArray) ToLBBackendAddressPoolResourceSettingsArrayOutput() LBBackendAddressPoolResourceSettingsArrayOutput {
	return i.ToLBBackendAddressPoolResourceSettingsArrayOutputWithContext(context.Background())
}

func (i LBBackendAddressPoolResourceSettingsArray) ToLBBackendAddressPoolResourceSettingsArrayOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBBackendAddressPoolResourceSettingsArrayOutput)
}

type LBBackendAddressPoolResourceSettingsOutput struct{ *pulumi.OutputState }

func (LBBackendAddressPoolResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBBackendAddressPoolResourceSettings)(nil)).Elem()
}

func (o LBBackendAddressPoolResourceSettingsOutput) ToLBBackendAddressPoolResourceSettingsOutput() LBBackendAddressPoolResourceSettingsOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsOutput) ToLBBackendAddressPoolResourceSettingsOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBBackendAddressPoolResourceSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LBBackendAddressPoolResourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (LBBackendAddressPoolResourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBBackendAddressPoolResourceSettings)(nil)).Elem()
}

func (o LBBackendAddressPoolResourceSettingsArrayOutput) ToLBBackendAddressPoolResourceSettingsArrayOutput() LBBackendAddressPoolResourceSettingsArrayOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsArrayOutput) ToLBBackendAddressPoolResourceSettingsArrayOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsArrayOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsArrayOutput) Index(i pulumi.IntInput) LBBackendAddressPoolResourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBBackendAddressPoolResourceSettings {
		return vs[0].([]LBBackendAddressPoolResourceSettings)[vs[1].(int)]
	}).(LBBackendAddressPoolResourceSettingsOutput)
}

type LBBackendAddressPoolResourceSettingsResponse struct {
	Name *string `pulumi:"name"`
}





type LBBackendAddressPoolResourceSettingsResponseInput interface {
	pulumi.Input

	ToLBBackendAddressPoolResourceSettingsResponseOutput() LBBackendAddressPoolResourceSettingsResponseOutput
	ToLBBackendAddressPoolResourceSettingsResponseOutputWithContext(context.Context) LBBackendAddressPoolResourceSettingsResponseOutput
}

type LBBackendAddressPoolResourceSettingsResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LBBackendAddressPoolResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LBBackendAddressPoolResourceSettingsResponse)(nil)).Elem()
}

func (i LBBackendAddressPoolResourceSettingsResponseArgs) ToLBBackendAddressPoolResourceSettingsResponseOutput() LBBackendAddressPoolResourceSettingsResponseOutput {
	return i.ToLBBackendAddressPoolResourceSettingsResponseOutputWithContext(context.Background())
}

func (i LBBackendAddressPoolResourceSettingsResponseArgs) ToLBBackendAddressPoolResourceSettingsResponseOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBBackendAddressPoolResourceSettingsResponseOutput)
}





type LBBackendAddressPoolResourceSettingsResponseArrayInput interface {
	pulumi.Input

	ToLBBackendAddressPoolResourceSettingsResponseArrayOutput() LBBackendAddressPoolResourceSettingsResponseArrayOutput
	ToLBBackendAddressPoolResourceSettingsResponseArrayOutputWithContext(context.Context) LBBackendAddressPoolResourceSettingsResponseArrayOutput
}

type LBBackendAddressPoolResourceSettingsResponseArray []LBBackendAddressPoolResourceSettingsResponseInput

func (LBBackendAddressPoolResourceSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBBackendAddressPoolResourceSettingsResponse)(nil)).Elem()
}

func (i LBBackendAddressPoolResourceSettingsResponseArray) ToLBBackendAddressPoolResourceSettingsResponseArrayOutput() LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return i.ToLBBackendAddressPoolResourceSettingsResponseArrayOutputWithContext(context.Background())
}

func (i LBBackendAddressPoolResourceSettingsResponseArray) ToLBBackendAddressPoolResourceSettingsResponseArrayOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBBackendAddressPoolResourceSettingsResponseArrayOutput)
}

type LBBackendAddressPoolResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (LBBackendAddressPoolResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBBackendAddressPoolResourceSettingsResponse)(nil)).Elem()
}

func (o LBBackendAddressPoolResourceSettingsResponseOutput) ToLBBackendAddressPoolResourceSettingsResponseOutput() LBBackendAddressPoolResourceSettingsResponseOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsResponseOutput) ToLBBackendAddressPoolResourceSettingsResponseOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsResponseOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBBackendAddressPoolResourceSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LBBackendAddressPoolResourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (LBBackendAddressPoolResourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBBackendAddressPoolResourceSettingsResponse)(nil)).Elem()
}

func (o LBBackendAddressPoolResourceSettingsResponseArrayOutput) ToLBBackendAddressPoolResourceSettingsResponseArrayOutput() LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsResponseArrayOutput) ToLBBackendAddressPoolResourceSettingsResponseArrayOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) LBBackendAddressPoolResourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBBackendAddressPoolResourceSettingsResponse {
		return vs[0].([]LBBackendAddressPoolResourceSettingsResponse)[vs[1].(int)]
	}).(LBBackendAddressPoolResourceSettingsResponseOutput)
}

type LBFrontendIPConfigurationResourceSettings struct {
	Name                      *string          `pulumi:"name"`
	PrivateIpAddress          *string          `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod *string          `pulumi:"privateIpAllocationMethod"`
	Subnet                    *SubnetReference `pulumi:"subnet"`
	Zones                     *string          `pulumi:"zones"`
}





type LBFrontendIPConfigurationResourceSettingsInput interface {
	pulumi.Input

	ToLBFrontendIPConfigurationResourceSettingsOutput() LBFrontendIPConfigurationResourceSettingsOutput
	ToLBFrontendIPConfigurationResourceSettingsOutputWithContext(context.Context) LBFrontendIPConfigurationResourceSettingsOutput
}

type LBFrontendIPConfigurationResourceSettingsArgs struct {
	Name                      pulumi.StringPtrInput   `pulumi:"name"`
	PrivateIpAddress          pulumi.StringPtrInput   `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod pulumi.StringPtrInput   `pulumi:"privateIpAllocationMethod"`
	Subnet                    SubnetReferencePtrInput `pulumi:"subnet"`
	Zones                     pulumi.StringPtrInput   `pulumi:"zones"`
}

func (LBFrontendIPConfigurationResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LBFrontendIPConfigurationResourceSettings)(nil)).Elem()
}

func (i LBFrontendIPConfigurationResourceSettingsArgs) ToLBFrontendIPConfigurationResourceSettingsOutput() LBFrontendIPConfigurationResourceSettingsOutput {
	return i.ToLBFrontendIPConfigurationResourceSettingsOutputWithContext(context.Background())
}

func (i LBFrontendIPConfigurationResourceSettingsArgs) ToLBFrontendIPConfigurationResourceSettingsOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBFrontendIPConfigurationResourceSettingsOutput)
}





type LBFrontendIPConfigurationResourceSettingsArrayInput interface {
	pulumi.Input

	ToLBFrontendIPConfigurationResourceSettingsArrayOutput() LBFrontendIPConfigurationResourceSettingsArrayOutput
	ToLBFrontendIPConfigurationResourceSettingsArrayOutputWithContext(context.Context) LBFrontendIPConfigurationResourceSettingsArrayOutput
}

type LBFrontendIPConfigurationResourceSettingsArray []LBFrontendIPConfigurationResourceSettingsInput

func (LBFrontendIPConfigurationResourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBFrontendIPConfigurationResourceSettings)(nil)).Elem()
}

func (i LBFrontendIPConfigurationResourceSettingsArray) ToLBFrontendIPConfigurationResourceSettingsArrayOutput() LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return i.ToLBFrontendIPConfigurationResourceSettingsArrayOutputWithContext(context.Background())
}

func (i LBFrontendIPConfigurationResourceSettingsArray) ToLBFrontendIPConfigurationResourceSettingsArrayOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBFrontendIPConfigurationResourceSettingsArrayOutput)
}

type LBFrontendIPConfigurationResourceSettingsOutput struct{ *pulumi.OutputState }

func (LBFrontendIPConfigurationResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBFrontendIPConfigurationResourceSettings)(nil)).Elem()
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) ToLBFrontendIPConfigurationResourceSettingsOutput() LBFrontendIPConfigurationResourceSettingsOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) ToLBFrontendIPConfigurationResourceSettingsOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) Subnet() SubnetReferencePtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *SubnetReference { return v.Subnet }).(SubnetReferencePtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type LBFrontendIPConfigurationResourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (LBFrontendIPConfigurationResourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBFrontendIPConfigurationResourceSettings)(nil)).Elem()
}

func (o LBFrontendIPConfigurationResourceSettingsArrayOutput) ToLBFrontendIPConfigurationResourceSettingsArrayOutput() LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsArrayOutput) ToLBFrontendIPConfigurationResourceSettingsArrayOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsArrayOutput) Index(i pulumi.IntInput) LBFrontendIPConfigurationResourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBFrontendIPConfigurationResourceSettings {
		return vs[0].([]LBFrontendIPConfigurationResourceSettings)[vs[1].(int)]
	}).(LBFrontendIPConfigurationResourceSettingsOutput)
}

type LBFrontendIPConfigurationResourceSettingsResponse struct {
	Name                      *string                  `pulumi:"name"`
	PrivateIpAddress          *string                  `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod *string                  `pulumi:"privateIpAllocationMethod"`
	Subnet                    *SubnetReferenceResponse `pulumi:"subnet"`
	Zones                     *string                  `pulumi:"zones"`
}





type LBFrontendIPConfigurationResourceSettingsResponseInput interface {
	pulumi.Input

	ToLBFrontendIPConfigurationResourceSettingsResponseOutput() LBFrontendIPConfigurationResourceSettingsResponseOutput
	ToLBFrontendIPConfigurationResourceSettingsResponseOutputWithContext(context.Context) LBFrontendIPConfigurationResourceSettingsResponseOutput
}

type LBFrontendIPConfigurationResourceSettingsResponseArgs struct {
	Name                      pulumi.StringPtrInput           `pulumi:"name"`
	PrivateIpAddress          pulumi.StringPtrInput           `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod pulumi.StringPtrInput           `pulumi:"privateIpAllocationMethod"`
	Subnet                    SubnetReferenceResponsePtrInput `pulumi:"subnet"`
	Zones                     pulumi.StringPtrInput           `pulumi:"zones"`
}

func (LBFrontendIPConfigurationResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LBFrontendIPConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (i LBFrontendIPConfigurationResourceSettingsResponseArgs) ToLBFrontendIPConfigurationResourceSettingsResponseOutput() LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return i.ToLBFrontendIPConfigurationResourceSettingsResponseOutputWithContext(context.Background())
}

func (i LBFrontendIPConfigurationResourceSettingsResponseArgs) ToLBFrontendIPConfigurationResourceSettingsResponseOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBFrontendIPConfigurationResourceSettingsResponseOutput)
}





type LBFrontendIPConfigurationResourceSettingsResponseArrayInput interface {
	pulumi.Input

	ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutput() LBFrontendIPConfigurationResourceSettingsResponseArrayOutput
	ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutputWithContext(context.Context) LBFrontendIPConfigurationResourceSettingsResponseArrayOutput
}

type LBFrontendIPConfigurationResourceSettingsResponseArray []LBFrontendIPConfigurationResourceSettingsResponseInput

func (LBFrontendIPConfigurationResourceSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBFrontendIPConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (i LBFrontendIPConfigurationResourceSettingsResponseArray) ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutput() LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return i.ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutputWithContext(context.Background())
}

func (i LBFrontendIPConfigurationResourceSettingsResponseArray) ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBFrontendIPConfigurationResourceSettingsResponseArrayOutput)
}

type LBFrontendIPConfigurationResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (LBFrontendIPConfigurationResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBFrontendIPConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) ToLBFrontendIPConfigurationResourceSettingsResponseOutput() LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) ToLBFrontendIPConfigurationResourceSettingsResponseOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) Subnet() SubnetReferenceResponsePtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *SubnetReferenceResponse { return v.Subnet }).(SubnetReferenceResponsePtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type LBFrontendIPConfigurationResourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (LBFrontendIPConfigurationResourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBFrontendIPConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (o LBFrontendIPConfigurationResourceSettingsResponseArrayOutput) ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutput() LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsResponseArrayOutput) ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBFrontendIPConfigurationResourceSettingsResponse {
		return vs[0].([]LBFrontendIPConfigurationResourceSettingsResponse)[vs[1].(int)]
	}).(LBFrontendIPConfigurationResourceSettingsResponseOutput)
}

type LoadBalancerBackendAddressPoolReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type LoadBalancerBackendAddressPoolReferenceInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolReferenceOutput() LoadBalancerBackendAddressPoolReferenceOutput
	ToLoadBalancerBackendAddressPoolReferenceOutputWithContext(context.Context) LoadBalancerBackendAddressPoolReferenceOutput
}

type LoadBalancerBackendAddressPoolReferenceArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (LoadBalancerBackendAddressPoolReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPoolReference)(nil)).Elem()
}

func (i LoadBalancerBackendAddressPoolReferenceArgs) ToLoadBalancerBackendAddressPoolReferenceOutput() LoadBalancerBackendAddressPoolReferenceOutput {
	return i.ToLoadBalancerBackendAddressPoolReferenceOutputWithContext(context.Background())
}

func (i LoadBalancerBackendAddressPoolReferenceArgs) ToLoadBalancerBackendAddressPoolReferenceOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolReferenceOutput)
}





type LoadBalancerBackendAddressPoolReferenceArrayInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolReferenceArrayOutput() LoadBalancerBackendAddressPoolReferenceArrayOutput
	ToLoadBalancerBackendAddressPoolReferenceArrayOutputWithContext(context.Context) LoadBalancerBackendAddressPoolReferenceArrayOutput
}

type LoadBalancerBackendAddressPoolReferenceArray []LoadBalancerBackendAddressPoolReferenceInput

func (LoadBalancerBackendAddressPoolReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerBackendAddressPoolReference)(nil)).Elem()
}

func (i LoadBalancerBackendAddressPoolReferenceArray) ToLoadBalancerBackendAddressPoolReferenceArrayOutput() LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return i.ToLoadBalancerBackendAddressPoolReferenceArrayOutputWithContext(context.Background())
}

func (i LoadBalancerBackendAddressPoolReferenceArray) ToLoadBalancerBackendAddressPoolReferenceArrayOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolReferenceArrayOutput)
}

type LoadBalancerBackendAddressPoolReferenceOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPoolReference)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolReferenceOutput) ToLoadBalancerBackendAddressPoolReferenceOutput() LoadBalancerBackendAddressPoolReferenceOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceOutput) ToLoadBalancerBackendAddressPoolReferenceOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerBackendAddressPoolReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerBackendAddressPoolReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerBackendAddressPoolReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type LoadBalancerBackendAddressPoolReferenceArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerBackendAddressPoolReference)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolReferenceArrayOutput) ToLoadBalancerBackendAddressPoolReferenceArrayOutput() LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceArrayOutput) ToLoadBalancerBackendAddressPoolReferenceArrayOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceArrayOutput) Index(i pulumi.IntInput) LoadBalancerBackendAddressPoolReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerBackendAddressPoolReference {
		return vs[0].([]LoadBalancerBackendAddressPoolReference)[vs[1].(int)]
	}).(LoadBalancerBackendAddressPoolReferenceOutput)
}

type LoadBalancerBackendAddressPoolReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type LoadBalancerBackendAddressPoolReferenceResponseInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolReferenceResponseOutput() LoadBalancerBackendAddressPoolReferenceResponseOutput
	ToLoadBalancerBackendAddressPoolReferenceResponseOutputWithContext(context.Context) LoadBalancerBackendAddressPoolReferenceResponseOutput
}

type LoadBalancerBackendAddressPoolReferenceResponseArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (LoadBalancerBackendAddressPoolReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPoolReferenceResponse)(nil)).Elem()
}

func (i LoadBalancerBackendAddressPoolReferenceResponseArgs) ToLoadBalancerBackendAddressPoolReferenceResponseOutput() LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return i.ToLoadBalancerBackendAddressPoolReferenceResponseOutputWithContext(context.Background())
}

func (i LoadBalancerBackendAddressPoolReferenceResponseArgs) ToLoadBalancerBackendAddressPoolReferenceResponseOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolReferenceResponseOutput)
}





type LoadBalancerBackendAddressPoolReferenceResponseArrayInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutput() LoadBalancerBackendAddressPoolReferenceResponseArrayOutput
	ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutputWithContext(context.Context) LoadBalancerBackendAddressPoolReferenceResponseArrayOutput
}

type LoadBalancerBackendAddressPoolReferenceResponseArray []LoadBalancerBackendAddressPoolReferenceResponseInput

func (LoadBalancerBackendAddressPoolReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerBackendAddressPoolReferenceResponse)(nil)).Elem()
}

func (i LoadBalancerBackendAddressPoolReferenceResponseArray) ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutput() LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return i.ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutputWithContext(context.Background())
}

func (i LoadBalancerBackendAddressPoolReferenceResponseArray) ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolReferenceResponseArrayOutput)
}

type LoadBalancerBackendAddressPoolReferenceResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPoolReferenceResponse)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolReferenceResponseOutput) ToLoadBalancerBackendAddressPoolReferenceResponseOutput() LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceResponseOutput) ToLoadBalancerBackendAddressPoolReferenceResponseOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerBackendAddressPoolReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerBackendAddressPoolReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerBackendAddressPoolReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type LoadBalancerBackendAddressPoolReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerBackendAddressPoolReferenceResponse)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolReferenceResponseArrayOutput) ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutput() LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceResponseArrayOutput) ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerBackendAddressPoolReferenceResponse {
		return vs[0].([]LoadBalancerBackendAddressPoolReferenceResponse)[vs[1].(int)]
	}).(LoadBalancerBackendAddressPoolReferenceResponseOutput)
}

type LoadBalancerNatRuleReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type LoadBalancerNatRuleReferenceInput interface {
	pulumi.Input

	ToLoadBalancerNatRuleReferenceOutput() LoadBalancerNatRuleReferenceOutput
	ToLoadBalancerNatRuleReferenceOutputWithContext(context.Context) LoadBalancerNatRuleReferenceOutput
}

type LoadBalancerNatRuleReferenceArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (LoadBalancerNatRuleReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerNatRuleReference)(nil)).Elem()
}

func (i LoadBalancerNatRuleReferenceArgs) ToLoadBalancerNatRuleReferenceOutput() LoadBalancerNatRuleReferenceOutput {
	return i.ToLoadBalancerNatRuleReferenceOutputWithContext(context.Background())
}

func (i LoadBalancerNatRuleReferenceArgs) ToLoadBalancerNatRuleReferenceOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerNatRuleReferenceOutput)
}





type LoadBalancerNatRuleReferenceArrayInput interface {
	pulumi.Input

	ToLoadBalancerNatRuleReferenceArrayOutput() LoadBalancerNatRuleReferenceArrayOutput
	ToLoadBalancerNatRuleReferenceArrayOutputWithContext(context.Context) LoadBalancerNatRuleReferenceArrayOutput
}

type LoadBalancerNatRuleReferenceArray []LoadBalancerNatRuleReferenceInput

func (LoadBalancerNatRuleReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerNatRuleReference)(nil)).Elem()
}

func (i LoadBalancerNatRuleReferenceArray) ToLoadBalancerNatRuleReferenceArrayOutput() LoadBalancerNatRuleReferenceArrayOutput {
	return i.ToLoadBalancerNatRuleReferenceArrayOutputWithContext(context.Background())
}

func (i LoadBalancerNatRuleReferenceArray) ToLoadBalancerNatRuleReferenceArrayOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerNatRuleReferenceArrayOutput)
}

type LoadBalancerNatRuleReferenceOutput struct{ *pulumi.OutputState }

func (LoadBalancerNatRuleReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerNatRuleReference)(nil)).Elem()
}

func (o LoadBalancerNatRuleReferenceOutput) ToLoadBalancerNatRuleReferenceOutput() LoadBalancerNatRuleReferenceOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceOutput) ToLoadBalancerNatRuleReferenceOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerNatRuleReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerNatRuleReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerNatRuleReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type LoadBalancerNatRuleReferenceArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerNatRuleReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerNatRuleReference)(nil)).Elem()
}

func (o LoadBalancerNatRuleReferenceArrayOutput) ToLoadBalancerNatRuleReferenceArrayOutput() LoadBalancerNatRuleReferenceArrayOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceArrayOutput) ToLoadBalancerNatRuleReferenceArrayOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceArrayOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceArrayOutput) Index(i pulumi.IntInput) LoadBalancerNatRuleReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerNatRuleReference {
		return vs[0].([]LoadBalancerNatRuleReference)[vs[1].(int)]
	}).(LoadBalancerNatRuleReferenceOutput)
}

type LoadBalancerNatRuleReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type LoadBalancerNatRuleReferenceResponseInput interface {
	pulumi.Input

	ToLoadBalancerNatRuleReferenceResponseOutput() LoadBalancerNatRuleReferenceResponseOutput
	ToLoadBalancerNatRuleReferenceResponseOutputWithContext(context.Context) LoadBalancerNatRuleReferenceResponseOutput
}

type LoadBalancerNatRuleReferenceResponseArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (LoadBalancerNatRuleReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerNatRuleReferenceResponse)(nil)).Elem()
}

func (i LoadBalancerNatRuleReferenceResponseArgs) ToLoadBalancerNatRuleReferenceResponseOutput() LoadBalancerNatRuleReferenceResponseOutput {
	return i.ToLoadBalancerNatRuleReferenceResponseOutputWithContext(context.Background())
}

func (i LoadBalancerNatRuleReferenceResponseArgs) ToLoadBalancerNatRuleReferenceResponseOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerNatRuleReferenceResponseOutput)
}





type LoadBalancerNatRuleReferenceResponseArrayInput interface {
	pulumi.Input

	ToLoadBalancerNatRuleReferenceResponseArrayOutput() LoadBalancerNatRuleReferenceResponseArrayOutput
	ToLoadBalancerNatRuleReferenceResponseArrayOutputWithContext(context.Context) LoadBalancerNatRuleReferenceResponseArrayOutput
}

type LoadBalancerNatRuleReferenceResponseArray []LoadBalancerNatRuleReferenceResponseInput

func (LoadBalancerNatRuleReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerNatRuleReferenceResponse)(nil)).Elem()
}

func (i LoadBalancerNatRuleReferenceResponseArray) ToLoadBalancerNatRuleReferenceResponseArrayOutput() LoadBalancerNatRuleReferenceResponseArrayOutput {
	return i.ToLoadBalancerNatRuleReferenceResponseArrayOutputWithContext(context.Background())
}

func (i LoadBalancerNatRuleReferenceResponseArray) ToLoadBalancerNatRuleReferenceResponseArrayOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerNatRuleReferenceResponseArrayOutput)
}

type LoadBalancerNatRuleReferenceResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerNatRuleReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerNatRuleReferenceResponse)(nil)).Elem()
}

func (o LoadBalancerNatRuleReferenceResponseOutput) ToLoadBalancerNatRuleReferenceResponseOutput() LoadBalancerNatRuleReferenceResponseOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceResponseOutput) ToLoadBalancerNatRuleReferenceResponseOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceResponseOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerNatRuleReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerNatRuleReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerNatRuleReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type LoadBalancerNatRuleReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerNatRuleReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerNatRuleReferenceResponse)(nil)).Elem()
}

func (o LoadBalancerNatRuleReferenceResponseArrayOutput) ToLoadBalancerNatRuleReferenceResponseArrayOutput() LoadBalancerNatRuleReferenceResponseArrayOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceResponseArrayOutput) ToLoadBalancerNatRuleReferenceResponseArrayOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceResponseArrayOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancerNatRuleReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerNatRuleReferenceResponse {
		return vs[0].([]LoadBalancerNatRuleReferenceResponse)[vs[1].(int)]
	}).(LoadBalancerNatRuleReferenceResponseOutput)
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





type LoadBalancerResourceSettingsInput interface {
	pulumi.Input

	ToLoadBalancerResourceSettingsOutput() LoadBalancerResourceSettingsOutput
	ToLoadBalancerResourceSettingsOutputWithContext(context.Context) LoadBalancerResourceSettingsOutput
}

type LoadBalancerResourceSettingsArgs struct {
	BackendAddressPools      LBBackendAddressPoolResourceSettingsArrayInput      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations LBFrontendIPConfigurationResourceSettingsArrayInput `pulumi:"frontendIPConfigurations"`
	ResourceType             pulumi.StringInput                                  `pulumi:"resourceType"`
	Sku                      pulumi.StringPtrInput                               `pulumi:"sku"`
	Tags                     pulumi.StringMapInput                               `pulumi:"tags"`
	TargetResourceName       pulumi.StringInput                                  `pulumi:"targetResourceName"`
	Zones                    pulumi.StringPtrInput                               `pulumi:"zones"`
}

func (LoadBalancerResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerResourceSettings)(nil)).Elem()
}

func (i LoadBalancerResourceSettingsArgs) ToLoadBalancerResourceSettingsOutput() LoadBalancerResourceSettingsOutput {
	return i.ToLoadBalancerResourceSettingsOutputWithContext(context.Background())
}

func (i LoadBalancerResourceSettingsArgs) ToLoadBalancerResourceSettingsOutputWithContext(ctx context.Context) LoadBalancerResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerResourceSettingsOutput)
}

type LoadBalancerResourceSettingsOutput struct{ *pulumi.OutputState }

func (LoadBalancerResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerResourceSettings)(nil)).Elem()
}

func (o LoadBalancerResourceSettingsOutput) ToLoadBalancerResourceSettingsOutput() LoadBalancerResourceSettingsOutput {
	return o
}

func (o LoadBalancerResourceSettingsOutput) ToLoadBalancerResourceSettingsOutputWithContext(ctx context.Context) LoadBalancerResourceSettingsOutput {
	return o
}

func (o LoadBalancerResourceSettingsOutput) BackendAddressPools() LBBackendAddressPoolResourceSettingsArrayOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) []LBBackendAddressPoolResourceSettings {
		return v.BackendAddressPools
	}).(LBBackendAddressPoolResourceSettingsArrayOutput)
}

func (o LoadBalancerResourceSettingsOutput) FrontendIPConfigurations() LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) []LBFrontendIPConfigurationResourceSettings {
		return v.FrontendIPConfigurations
	}).(LBFrontendIPConfigurationResourceSettingsArrayOutput)
}

func (o LoadBalancerResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o LoadBalancerResourceSettingsOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LoadBalancerResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o LoadBalancerResourceSettingsOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) *string { return v.Zones }).(pulumi.StringPtrOutput)
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





type LoadBalancerResourceSettingsResponseInput interface {
	pulumi.Input

	ToLoadBalancerResourceSettingsResponseOutput() LoadBalancerResourceSettingsResponseOutput
	ToLoadBalancerResourceSettingsResponseOutputWithContext(context.Context) LoadBalancerResourceSettingsResponseOutput
}

type LoadBalancerResourceSettingsResponseArgs struct {
	BackendAddressPools      LBBackendAddressPoolResourceSettingsResponseArrayInput      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations LBFrontendIPConfigurationResourceSettingsResponseArrayInput `pulumi:"frontendIPConfigurations"`
	ResourceType             pulumi.StringInput                                          `pulumi:"resourceType"`
	Sku                      pulumi.StringPtrInput                                       `pulumi:"sku"`
	Tags                     pulumi.StringMapInput                                       `pulumi:"tags"`
	TargetResourceName       pulumi.StringInput                                          `pulumi:"targetResourceName"`
	Zones                    pulumi.StringPtrInput                                       `pulumi:"zones"`
}

func (LoadBalancerResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerResourceSettingsResponse)(nil)).Elem()
}

func (i LoadBalancerResourceSettingsResponseArgs) ToLoadBalancerResourceSettingsResponseOutput() LoadBalancerResourceSettingsResponseOutput {
	return i.ToLoadBalancerResourceSettingsResponseOutputWithContext(context.Background())
}

func (i LoadBalancerResourceSettingsResponseArgs) ToLoadBalancerResourceSettingsResponseOutputWithContext(ctx context.Context) LoadBalancerResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerResourceSettingsResponseOutput)
}

type LoadBalancerResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerResourceSettingsResponse)(nil)).Elem()
}

func (o LoadBalancerResourceSettingsResponseOutput) ToLoadBalancerResourceSettingsResponseOutput() LoadBalancerResourceSettingsResponseOutput {
	return o
}

func (o LoadBalancerResourceSettingsResponseOutput) ToLoadBalancerResourceSettingsResponseOutputWithContext(ctx context.Context) LoadBalancerResourceSettingsResponseOutput {
	return o
}

func (o LoadBalancerResourceSettingsResponseOutput) BackendAddressPools() LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) []LBBackendAddressPoolResourceSettingsResponse {
		return v.BackendAddressPools
	}).(LBBackendAddressPoolResourceSettingsResponseArrayOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) FrontendIPConfigurations() LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) []LBFrontendIPConfigurationResourceSettingsResponse {
		return v.FrontendIPConfigurations
	}).(LBFrontendIPConfigurationResourceSettingsResponseArrayOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type ManualResolutionPropertiesResponse struct {
	TargetId *string `pulumi:"targetId"`
}





type ManualResolutionPropertiesResponseInput interface {
	pulumi.Input

	ToManualResolutionPropertiesResponseOutput() ManualResolutionPropertiesResponseOutput
	ToManualResolutionPropertiesResponseOutputWithContext(context.Context) ManualResolutionPropertiesResponseOutput
}

type ManualResolutionPropertiesResponseArgs struct {
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (ManualResolutionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManualResolutionPropertiesResponse)(nil)).Elem()
}

func (i ManualResolutionPropertiesResponseArgs) ToManualResolutionPropertiesResponseOutput() ManualResolutionPropertiesResponseOutput {
	return i.ToManualResolutionPropertiesResponseOutputWithContext(context.Background())
}

func (i ManualResolutionPropertiesResponseArgs) ToManualResolutionPropertiesResponseOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualResolutionPropertiesResponseOutput)
}

func (i ManualResolutionPropertiesResponseArgs) ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput {
	return i.ToManualResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ManualResolutionPropertiesResponseArgs) ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualResolutionPropertiesResponseOutput).ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx)
}









type ManualResolutionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput
	ToManualResolutionPropertiesResponsePtrOutputWithContext(context.Context) ManualResolutionPropertiesResponsePtrOutput
}

type manualResolutionPropertiesResponsePtrType ManualResolutionPropertiesResponseArgs

func ManualResolutionPropertiesResponsePtr(v *ManualResolutionPropertiesResponseArgs) ManualResolutionPropertiesResponsePtrInput {
	return (*manualResolutionPropertiesResponsePtrType)(v)
}

func (*manualResolutionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManualResolutionPropertiesResponse)(nil)).Elem()
}

func (i *manualResolutionPropertiesResponsePtrType) ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput {
	return i.ToManualResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *manualResolutionPropertiesResponsePtrType) ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualResolutionPropertiesResponsePtrOutput)
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

func (o ManualResolutionPropertiesResponseOutput) ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput {
	return o.ToManualResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ManualResolutionPropertiesResponseOutput) ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManualResolutionPropertiesResponse) *ManualResolutionPropertiesResponse {
		return &v
	}).(ManualResolutionPropertiesResponsePtrOutput)
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





type MoveCollectionPropertiesResponseInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesResponseOutput() MoveCollectionPropertiesResponseOutput
	ToMoveCollectionPropertiesResponseOutputWithContext(context.Context) MoveCollectionPropertiesResponseOutput
}

type MoveCollectionPropertiesResponseArgs struct {
	Errors            MoveCollectionPropertiesResponseErrorsInput `pulumi:"errors"`
	ProvisioningState pulumi.StringInput                          `pulumi:"provisioningState"`
	SourceRegion      pulumi.StringInput                          `pulumi:"sourceRegion"`
	TargetRegion      pulumi.StringInput                          `pulumi:"targetRegion"`
}

func (MoveCollectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionPropertiesResponse)(nil)).Elem()
}

func (i MoveCollectionPropertiesResponseArgs) ToMoveCollectionPropertiesResponseOutput() MoveCollectionPropertiesResponseOutput {
	return i.ToMoveCollectionPropertiesResponseOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesResponseArgs) ToMoveCollectionPropertiesResponseOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseOutput)
}

func (i MoveCollectionPropertiesResponseArgs) ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput {
	return i.ToMoveCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesResponseArgs) ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseOutput).ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx)
}









type MoveCollectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput
	ToMoveCollectionPropertiesResponsePtrOutputWithContext(context.Context) MoveCollectionPropertiesResponsePtrOutput
}

type moveCollectionPropertiesResponsePtrType MoveCollectionPropertiesResponseArgs

func MoveCollectionPropertiesResponsePtr(v *MoveCollectionPropertiesResponseArgs) MoveCollectionPropertiesResponsePtrInput {
	return (*moveCollectionPropertiesResponsePtrType)(v)
}

func (*moveCollectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionPropertiesResponse)(nil)).Elem()
}

func (i *moveCollectionPropertiesResponsePtrType) ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput {
	return i.ToMoveCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *moveCollectionPropertiesResponsePtrType) ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponsePtrOutput)
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

func (o MoveCollectionPropertiesResponseOutput) ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput {
	return o.ToMoveCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MoveCollectionPropertiesResponseOutput) ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveCollectionPropertiesResponse) *MoveCollectionPropertiesResponse {
		return &v
	}).(MoveCollectionPropertiesResponsePtrOutput)
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

type MoveCollectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionPropertiesResponse)(nil)).Elem()
}

func (o MoveCollectionPropertiesResponsePtrOutput) ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput {
	return o
}

func (o MoveCollectionPropertiesResponsePtrOutput) ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponsePtrOutput {
	return o
}

func (o MoveCollectionPropertiesResponsePtrOutput) Elem() MoveCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) MoveCollectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MoveCollectionPropertiesResponse
		return ret
	}).(MoveCollectionPropertiesResponseOutput)
}

func (o MoveCollectionPropertiesResponsePtrOutput) Errors() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) *MoveCollectionPropertiesResponseErrors {
		if v == nil {
			return nil
		}
		return &v.Errors
	}).(MoveCollectionPropertiesResponseErrorsPtrOutput)
}

func (o MoveCollectionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MoveCollectionPropertiesResponsePtrOutput) SourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceRegion
	}).(pulumi.StringPtrOutput)
}

func (o MoveCollectionPropertiesResponsePtrOutput) TargetRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetRegion
	}).(pulumi.StringPtrOutput)
}

type MoveCollectionPropertiesResponseErrors struct {
	Properties *MoveResourceErrorBodyResponse `pulumi:"properties"`
}





type MoveCollectionPropertiesResponseErrorsInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesResponseErrorsOutput() MoveCollectionPropertiesResponseErrorsOutput
	ToMoveCollectionPropertiesResponseErrorsOutputWithContext(context.Context) MoveCollectionPropertiesResponseErrorsOutput
}

type MoveCollectionPropertiesResponseErrorsArgs struct {
	Properties MoveResourceErrorBodyResponsePtrInput `pulumi:"properties"`
}

func (MoveCollectionPropertiesResponseErrorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionPropertiesResponseErrors)(nil)).Elem()
}

func (i MoveCollectionPropertiesResponseErrorsArgs) ToMoveCollectionPropertiesResponseErrorsOutput() MoveCollectionPropertiesResponseErrorsOutput {
	return i.ToMoveCollectionPropertiesResponseErrorsOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesResponseErrorsArgs) ToMoveCollectionPropertiesResponseErrorsOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseErrorsOutput)
}

func (i MoveCollectionPropertiesResponseErrorsArgs) ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return i.ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesResponseErrorsArgs) ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseErrorsOutput).ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx)
}









type MoveCollectionPropertiesResponseErrorsPtrInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput
	ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput
}

type moveCollectionPropertiesResponseErrorsPtrType MoveCollectionPropertiesResponseErrorsArgs

func MoveCollectionPropertiesResponseErrorsPtr(v *MoveCollectionPropertiesResponseErrorsArgs) MoveCollectionPropertiesResponseErrorsPtrInput {
	return (*moveCollectionPropertiesResponseErrorsPtrType)(v)
}

func (*moveCollectionPropertiesResponseErrorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionPropertiesResponseErrors)(nil)).Elem()
}

func (i *moveCollectionPropertiesResponseErrorsPtrType) ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return i.ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (i *moveCollectionPropertiesResponseErrorsPtrType) ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseErrorsPtrOutput)
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

func (o MoveCollectionPropertiesResponseErrorsOutput) ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o.ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (o MoveCollectionPropertiesResponseErrorsOutput) ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveCollectionPropertiesResponseErrors) *MoveCollectionPropertiesResponseErrors {
		return &v
	}).(MoveCollectionPropertiesResponseErrorsPtrOutput)
}

func (o MoveCollectionPropertiesResponseErrorsOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponseErrors) *MoveResourceErrorBodyResponse { return v.Properties }).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveCollectionPropertiesResponseErrorsPtrOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesResponseErrorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionPropertiesResponseErrors)(nil)).Elem()
}

func (o MoveCollectionPropertiesResponseErrorsPtrOutput) ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o
}

func (o MoveCollectionPropertiesResponseErrorsPtrOutput) ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o
}

func (o MoveCollectionPropertiesResponseErrorsPtrOutput) Elem() MoveCollectionPropertiesResponseErrorsOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponseErrors) MoveCollectionPropertiesResponseErrors {
		if v != nil {
			return *v
		}
		var ret MoveCollectionPropertiesResponseErrors
		return ret
	}).(MoveCollectionPropertiesResponseErrorsOutput)
}

func (o MoveCollectionPropertiesResponseErrorsPtrOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponseErrors) *MoveResourceErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(MoveResourceErrorBodyResponsePtrOutput)
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





type MoveResourceDependencyOverrideResponseInput interface {
	pulumi.Input

	ToMoveResourceDependencyOverrideResponseOutput() MoveResourceDependencyOverrideResponseOutput
	ToMoveResourceDependencyOverrideResponseOutputWithContext(context.Context) MoveResourceDependencyOverrideResponseOutput
}

type MoveResourceDependencyOverrideResponseArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (MoveResourceDependencyOverrideResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyOverrideResponse)(nil)).Elem()
}

func (i MoveResourceDependencyOverrideResponseArgs) ToMoveResourceDependencyOverrideResponseOutput() MoveResourceDependencyOverrideResponseOutput {
	return i.ToMoveResourceDependencyOverrideResponseOutputWithContext(context.Background())
}

func (i MoveResourceDependencyOverrideResponseArgs) ToMoveResourceDependencyOverrideResponseOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyOverrideResponseOutput)
}





type MoveResourceDependencyOverrideResponseArrayInput interface {
	pulumi.Input

	ToMoveResourceDependencyOverrideResponseArrayOutput() MoveResourceDependencyOverrideResponseArrayOutput
	ToMoveResourceDependencyOverrideResponseArrayOutputWithContext(context.Context) MoveResourceDependencyOverrideResponseArrayOutput
}

type MoveResourceDependencyOverrideResponseArray []MoveResourceDependencyOverrideResponseInput

func (MoveResourceDependencyOverrideResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyOverrideResponse)(nil)).Elem()
}

func (i MoveResourceDependencyOverrideResponseArray) ToMoveResourceDependencyOverrideResponseArrayOutput() MoveResourceDependencyOverrideResponseArrayOutput {
	return i.ToMoveResourceDependencyOverrideResponseArrayOutputWithContext(context.Background())
}

func (i MoveResourceDependencyOverrideResponseArray) ToMoveResourceDependencyOverrideResponseArrayOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyOverrideResponseArrayOutput)
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





type MoveResourceDependencyResponseInput interface {
	pulumi.Input

	ToMoveResourceDependencyResponseOutput() MoveResourceDependencyResponseOutput
	ToMoveResourceDependencyResponseOutputWithContext(context.Context) MoveResourceDependencyResponseOutput
}

type MoveResourceDependencyResponseArgs struct {
	AutomaticResolution AutomaticResolutionPropertiesResponsePtrInput `pulumi:"automaticResolution"`
	DependencyType      pulumi.StringPtrInput                         `pulumi:"dependencyType"`
	Id                  pulumi.StringPtrInput                         `pulumi:"id"`
	IsOptional          pulumi.StringPtrInput                         `pulumi:"isOptional"`
	ManualResolution    ManualResolutionPropertiesResponsePtrInput    `pulumi:"manualResolution"`
	ResolutionStatus    pulumi.StringPtrInput                         `pulumi:"resolutionStatus"`
	ResolutionType      pulumi.StringPtrInput                         `pulumi:"resolutionType"`
}

func (MoveResourceDependencyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyResponse)(nil)).Elem()
}

func (i MoveResourceDependencyResponseArgs) ToMoveResourceDependencyResponseOutput() MoveResourceDependencyResponseOutput {
	return i.ToMoveResourceDependencyResponseOutputWithContext(context.Background())
}

func (i MoveResourceDependencyResponseArgs) ToMoveResourceDependencyResponseOutputWithContext(ctx context.Context) MoveResourceDependencyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyResponseOutput)
}





type MoveResourceDependencyResponseArrayInput interface {
	pulumi.Input

	ToMoveResourceDependencyResponseArrayOutput() MoveResourceDependencyResponseArrayOutput
	ToMoveResourceDependencyResponseArrayOutputWithContext(context.Context) MoveResourceDependencyResponseArrayOutput
}

type MoveResourceDependencyResponseArray []MoveResourceDependencyResponseInput

func (MoveResourceDependencyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyResponse)(nil)).Elem()
}

func (i MoveResourceDependencyResponseArray) ToMoveResourceDependencyResponseArrayOutput() MoveResourceDependencyResponseArrayOutput {
	return i.ToMoveResourceDependencyResponseArrayOutputWithContext(context.Background())
}

func (i MoveResourceDependencyResponseArray) ToMoveResourceDependencyResponseArrayOutputWithContext(ctx context.Context) MoveResourceDependencyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyResponseArrayOutput)
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





type MoveResourceErrorBodyResponseInput interface {
	pulumi.Input

	ToMoveResourceErrorBodyResponseOutput() MoveResourceErrorBodyResponseOutput
	ToMoveResourceErrorBodyResponseOutputWithContext(context.Context) MoveResourceErrorBodyResponseOutput
}

type MoveResourceErrorBodyResponseArgs struct {
	Code    pulumi.StringInput                      `pulumi:"code"`
	Details MoveResourceErrorBodyResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput                      `pulumi:"message"`
	Target  pulumi.StringInput                      `pulumi:"target"`
}

func (MoveResourceErrorBodyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (i MoveResourceErrorBodyResponseArgs) ToMoveResourceErrorBodyResponseOutput() MoveResourceErrorBodyResponseOutput {
	return i.ToMoveResourceErrorBodyResponseOutputWithContext(context.Background())
}

func (i MoveResourceErrorBodyResponseArgs) ToMoveResourceErrorBodyResponseOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorBodyResponseOutput)
}

func (i MoveResourceErrorBodyResponseArgs) ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput {
	return i.ToMoveResourceErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (i MoveResourceErrorBodyResponseArgs) ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorBodyResponseOutput).ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx)
}









type MoveResourceErrorBodyResponsePtrInput interface {
	pulumi.Input

	ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput
	ToMoveResourceErrorBodyResponsePtrOutputWithContext(context.Context) MoveResourceErrorBodyResponsePtrOutput
}

type moveResourceErrorBodyResponsePtrType MoveResourceErrorBodyResponseArgs

func MoveResourceErrorBodyResponsePtr(v *MoveResourceErrorBodyResponseArgs) MoveResourceErrorBodyResponsePtrInput {
	return (*moveResourceErrorBodyResponsePtrType)(v)
}

func (*moveResourceErrorBodyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (i *moveResourceErrorBodyResponsePtrType) ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput {
	return i.ToMoveResourceErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (i *moveResourceErrorBodyResponsePtrType) ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorBodyResponsePtrOutput)
}





type MoveResourceErrorBodyResponseArrayInput interface {
	pulumi.Input

	ToMoveResourceErrorBodyResponseArrayOutput() MoveResourceErrorBodyResponseArrayOutput
	ToMoveResourceErrorBodyResponseArrayOutputWithContext(context.Context) MoveResourceErrorBodyResponseArrayOutput
}

type MoveResourceErrorBodyResponseArray []MoveResourceErrorBodyResponseInput

func (MoveResourceErrorBodyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (i MoveResourceErrorBodyResponseArray) ToMoveResourceErrorBodyResponseArrayOutput() MoveResourceErrorBodyResponseArrayOutput {
	return i.ToMoveResourceErrorBodyResponseArrayOutputWithContext(context.Background())
}

func (i MoveResourceErrorBodyResponseArray) ToMoveResourceErrorBodyResponseArrayOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorBodyResponseArrayOutput)
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

func (o MoveResourceErrorBodyResponseOutput) ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput {
	return o.ToMoveResourceErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (o MoveResourceErrorBodyResponseOutput) ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourceErrorBodyResponse) *MoveResourceErrorBodyResponse {
		return &v
	}).(MoveResourceErrorBodyResponsePtrOutput)
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





type MoveResourceErrorResponseInput interface {
	pulumi.Input

	ToMoveResourceErrorResponseOutput() MoveResourceErrorResponseOutput
	ToMoveResourceErrorResponseOutputWithContext(context.Context) MoveResourceErrorResponseOutput
}

type MoveResourceErrorResponseArgs struct {
	Properties MoveResourceErrorBodyResponsePtrInput `pulumi:"properties"`
}

func (MoveResourceErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceErrorResponse)(nil)).Elem()
}

func (i MoveResourceErrorResponseArgs) ToMoveResourceErrorResponseOutput() MoveResourceErrorResponseOutput {
	return i.ToMoveResourceErrorResponseOutputWithContext(context.Background())
}

func (i MoveResourceErrorResponseArgs) ToMoveResourceErrorResponseOutputWithContext(ctx context.Context) MoveResourceErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorResponseOutput)
}

func (i MoveResourceErrorResponseArgs) ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput {
	return i.ToMoveResourceErrorResponsePtrOutputWithContext(context.Background())
}

func (i MoveResourceErrorResponseArgs) ToMoveResourceErrorResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorResponseOutput).ToMoveResourceErrorResponsePtrOutputWithContext(ctx)
}









type MoveResourceErrorResponsePtrInput interface {
	pulumi.Input

	ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput
	ToMoveResourceErrorResponsePtrOutputWithContext(context.Context) MoveResourceErrorResponsePtrOutput
}

type moveResourceErrorResponsePtrType MoveResourceErrorResponseArgs

func MoveResourceErrorResponsePtr(v *MoveResourceErrorResponseArgs) MoveResourceErrorResponsePtrInput {
	return (*moveResourceErrorResponsePtrType)(v)
}

func (*moveResourceErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceErrorResponse)(nil)).Elem()
}

func (i *moveResourceErrorResponsePtrType) ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput {
	return i.ToMoveResourceErrorResponsePtrOutputWithContext(context.Background())
}

func (i *moveResourceErrorResponsePtrType) ToMoveResourceErrorResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorResponsePtrOutput)
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

func (o MoveResourceErrorResponseOutput) ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput {
	return o.ToMoveResourceErrorResponsePtrOutputWithContext(context.Background())
}

func (o MoveResourceErrorResponseOutput) ToMoveResourceErrorResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourceErrorResponse) *MoveResourceErrorResponse {
		return &v
	}).(MoveResourceErrorResponsePtrOutput)
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





type MoveResourcePropertiesResponseInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseOutput() MoveResourcePropertiesResponseOutput
	ToMoveResourcePropertiesResponseOutputWithContext(context.Context) MoveResourcePropertiesResponseOutput
}

type MoveResourcePropertiesResponseArgs struct {
	DependsOn              MoveResourceDependencyResponseArrayInput         `pulumi:"dependsOn"`
	DependsOnOverrides     MoveResourceDependencyOverrideResponseArrayInput `pulumi:"dependsOnOverrides"`
	Errors                 MoveResourcePropertiesResponseErrorsInput        `pulumi:"errors"`
	ExistingTargetId       pulumi.StringPtrInput                            `pulumi:"existingTargetId"`
	IsResolveRequired      pulumi.BoolInput                                 `pulumi:"isResolveRequired"`
	MoveStatus             MoveResourcePropertiesResponseMoveStatusInput    `pulumi:"moveStatus"`
	ProvisioningState      pulumi.StringInput                               `pulumi:"provisioningState"`
	ResourceSettings       pulumi.Input                                     `pulumi:"resourceSettings"`
	SourceId               pulumi.StringInput                               `pulumi:"sourceId"`
	SourceResourceSettings pulumi.Input                                     `pulumi:"sourceResourceSettings"`
	TargetId               pulumi.StringInput                               `pulumi:"targetId"`
}

func (MoveResourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponse)(nil)).Elem()
}

func (i MoveResourcePropertiesResponseArgs) ToMoveResourcePropertiesResponseOutput() MoveResourcePropertiesResponseOutput {
	return i.ToMoveResourcePropertiesResponseOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseArgs) ToMoveResourcePropertiesResponseOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseOutput)
}

func (i MoveResourcePropertiesResponseArgs) ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput {
	return i.ToMoveResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseArgs) ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseOutput).ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx)
}









type MoveResourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput
	ToMoveResourcePropertiesResponsePtrOutputWithContext(context.Context) MoveResourcePropertiesResponsePtrOutput
}

type moveResourcePropertiesResponsePtrType MoveResourcePropertiesResponseArgs

func MoveResourcePropertiesResponsePtr(v *MoveResourcePropertiesResponseArgs) MoveResourcePropertiesResponsePtrInput {
	return (*moveResourcePropertiesResponsePtrType)(v)
}

func (*moveResourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponse)(nil)).Elem()
}

func (i *moveResourcePropertiesResponsePtrType) ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput {
	return i.ToMoveResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *moveResourcePropertiesResponsePtrType) ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponsePtrOutput)
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

func (o MoveResourcePropertiesResponseOutput) ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput {
	return o.ToMoveResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MoveResourcePropertiesResponseOutput) ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourcePropertiesResponse) *MoveResourcePropertiesResponse {
		return &v
	}).(MoveResourcePropertiesResponsePtrOutput)
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

type MoveResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponse)(nil)).Elem()
}

func (o MoveResourcePropertiesResponsePtrOutput) ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput {
	return o
}

func (o MoveResourcePropertiesResponsePtrOutput) ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponsePtrOutput {
	return o
}

func (o MoveResourcePropertiesResponsePtrOutput) Elem() MoveResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) MoveResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MoveResourcePropertiesResponse
		return ret
	}).(MoveResourcePropertiesResponseOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) DependsOn() MoveResourceDependencyResponseArrayOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) []MoveResourceDependencyResponse {
		if v == nil {
			return nil
		}
		return v.DependsOn
	}).(MoveResourceDependencyResponseArrayOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) DependsOnOverrides() MoveResourceDependencyOverrideResponseArrayOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) []MoveResourceDependencyOverrideResponse {
		if v == nil {
			return nil
		}
		return v.DependsOnOverrides
	}).(MoveResourceDependencyOverrideResponseArrayOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) Errors() MoveResourcePropertiesResponseErrorsPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *MoveResourcePropertiesResponseErrors {
		if v == nil {
			return nil
		}
		return &v.Errors
	}).(MoveResourcePropertiesResponseErrorsPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) ExistingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExistingTargetId
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) IsResolveRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsResolveRequired
	}).(pulumi.BoolPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) MoveStatus() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *MoveResourcePropertiesResponseMoveStatus {
		if v == nil {
			return nil
		}
		return &v.MoveStatus
	}).(MoveResourcePropertiesResponseMoveStatusPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) ResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ResourceSettings
	}).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceId
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) SourceResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.SourceResourceSettings
	}).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetId
	}).(pulumi.StringPtrOutput)
}

type MoveResourcePropertiesResponseErrors struct {
	Properties *MoveResourceErrorBodyResponse `pulumi:"properties"`
}





type MoveResourcePropertiesResponseErrorsInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseErrorsOutput() MoveResourcePropertiesResponseErrorsOutput
	ToMoveResourcePropertiesResponseErrorsOutputWithContext(context.Context) MoveResourcePropertiesResponseErrorsOutput
}

type MoveResourcePropertiesResponseErrorsArgs struct {
	Properties MoveResourceErrorBodyResponsePtrInput `pulumi:"properties"`
}

func (MoveResourcePropertiesResponseErrorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponseErrors)(nil)).Elem()
}

func (i MoveResourcePropertiesResponseErrorsArgs) ToMoveResourcePropertiesResponseErrorsOutput() MoveResourcePropertiesResponseErrorsOutput {
	return i.ToMoveResourcePropertiesResponseErrorsOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseErrorsArgs) ToMoveResourcePropertiesResponseErrorsOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseErrorsOutput)
}

func (i MoveResourcePropertiesResponseErrorsArgs) ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput {
	return i.ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseErrorsArgs) ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseErrorsOutput).ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx)
}









type MoveResourcePropertiesResponseErrorsPtrInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput
	ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(context.Context) MoveResourcePropertiesResponseErrorsPtrOutput
}

type moveResourcePropertiesResponseErrorsPtrType MoveResourcePropertiesResponseErrorsArgs

func MoveResourcePropertiesResponseErrorsPtr(v *MoveResourcePropertiesResponseErrorsArgs) MoveResourcePropertiesResponseErrorsPtrInput {
	return (*moveResourcePropertiesResponseErrorsPtrType)(v)
}

func (*moveResourcePropertiesResponseErrorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponseErrors)(nil)).Elem()
}

func (i *moveResourcePropertiesResponseErrorsPtrType) ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput {
	return i.ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (i *moveResourcePropertiesResponseErrorsPtrType) ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseErrorsPtrOutput)
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

func (o MoveResourcePropertiesResponseErrorsOutput) ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput {
	return o.ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (o MoveResourcePropertiesResponseErrorsOutput) ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourcePropertiesResponseErrors) *MoveResourcePropertiesResponseErrors {
		return &v
	}).(MoveResourcePropertiesResponseErrorsPtrOutput)
}

func (o MoveResourcePropertiesResponseErrorsOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseErrors) *MoveResourceErrorBodyResponse { return v.Properties }).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourcePropertiesResponseErrorsPtrOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseErrorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponseErrors)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseErrorsPtrOutput) ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput {
	return o
}

func (o MoveResourcePropertiesResponseErrorsPtrOutput) ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsPtrOutput {
	return o
}

func (o MoveResourcePropertiesResponseErrorsPtrOutput) Elem() MoveResourcePropertiesResponseErrorsOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseErrors) MoveResourcePropertiesResponseErrors {
		if v != nil {
			return *v
		}
		var ret MoveResourcePropertiesResponseErrors
		return ret
	}).(MoveResourcePropertiesResponseErrorsOutput)
}

func (o MoveResourcePropertiesResponseErrorsPtrOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseErrors) *MoveResourceErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourcePropertiesResponseMoveStatus struct {
	Errors    *MoveResourceErrorResponse `pulumi:"errors"`
	JobStatus *JobStatusResponse         `pulumi:"jobStatus"`
	MoveState string                     `pulumi:"moveState"`
}





type MoveResourcePropertiesResponseMoveStatusInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseMoveStatusOutput() MoveResourcePropertiesResponseMoveStatusOutput
	ToMoveResourcePropertiesResponseMoveStatusOutputWithContext(context.Context) MoveResourcePropertiesResponseMoveStatusOutput
}

type MoveResourcePropertiesResponseMoveStatusArgs struct {
	Errors    MoveResourceErrorResponsePtrInput `pulumi:"errors"`
	JobStatus JobStatusResponsePtrInput         `pulumi:"jobStatus"`
	MoveState pulumi.StringInput                `pulumi:"moveState"`
}

func (MoveResourcePropertiesResponseMoveStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponseMoveStatus)(nil)).Elem()
}

func (i MoveResourcePropertiesResponseMoveStatusArgs) ToMoveResourcePropertiesResponseMoveStatusOutput() MoveResourcePropertiesResponseMoveStatusOutput {
	return i.ToMoveResourcePropertiesResponseMoveStatusOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseMoveStatusArgs) ToMoveResourcePropertiesResponseMoveStatusOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseMoveStatusOutput)
}

func (i MoveResourcePropertiesResponseMoveStatusArgs) ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return i.ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseMoveStatusArgs) ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseMoveStatusOutput).ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx)
}









type MoveResourcePropertiesResponseMoveStatusPtrInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput
	ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput
}

type moveResourcePropertiesResponseMoveStatusPtrType MoveResourcePropertiesResponseMoveStatusArgs

func MoveResourcePropertiesResponseMoveStatusPtr(v *MoveResourcePropertiesResponseMoveStatusArgs) MoveResourcePropertiesResponseMoveStatusPtrInput {
	return (*moveResourcePropertiesResponseMoveStatusPtrType)(v)
}

func (*moveResourcePropertiesResponseMoveStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponseMoveStatus)(nil)).Elem()
}

func (i *moveResourcePropertiesResponseMoveStatusPtrType) ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return i.ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(context.Background())
}

func (i *moveResourcePropertiesResponseMoveStatusPtrType) ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseMoveStatusPtrOutput)
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

func (o MoveResourcePropertiesResponseMoveStatusOutput) ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o.ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(context.Background())
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourcePropertiesResponseMoveStatus) *MoveResourcePropertiesResponseMoveStatus {
		return &v
	}).(MoveResourcePropertiesResponseMoveStatusPtrOutput)
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

type MoveResourcePropertiesResponseMoveStatusPtrOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseMoveStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponseMoveStatus)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) Elem() MoveResourcePropertiesResponseMoveStatusOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseMoveStatus) MoveResourcePropertiesResponseMoveStatus {
		if v != nil {
			return *v
		}
		var ret MoveResourcePropertiesResponseMoveStatus
		return ret
	}).(MoveResourcePropertiesResponseMoveStatusOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) Errors() MoveResourceErrorResponsePtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseMoveStatus) *MoveResourceErrorResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(MoveResourceErrorResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) JobStatus() JobStatusResponsePtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseMoveStatus) *JobStatusResponse {
		if v == nil {
			return nil
		}
		return v.JobStatus
	}).(JobStatusResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) MoveState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseMoveStatus) *string {
		if v == nil {
			return nil
		}
		return &v.MoveState
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceResourceSettings struct {
	EnableAcceleratedNetworking *bool                                `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            []NicIpConfigurationResourceSettings `pulumi:"ipConfigurations"`
	ResourceType                string                               `pulumi:"resourceType"`
	Tags                        map[string]string                    `pulumi:"tags"`
	TargetResourceName          string                               `pulumi:"targetResourceName"`
}





type NetworkInterfaceResourceSettingsInput interface {
	pulumi.Input

	ToNetworkInterfaceResourceSettingsOutput() NetworkInterfaceResourceSettingsOutput
	ToNetworkInterfaceResourceSettingsOutputWithContext(context.Context) NetworkInterfaceResourceSettingsOutput
}

type NetworkInterfaceResourceSettingsArgs struct {
	EnableAcceleratedNetworking pulumi.BoolPtrInput                          `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            NicIpConfigurationResourceSettingsArrayInput `pulumi:"ipConfigurations"`
	ResourceType                pulumi.StringInput                           `pulumi:"resourceType"`
	Tags                        pulumi.StringMapInput                        `pulumi:"tags"`
	TargetResourceName          pulumi.StringInput                           `pulumi:"targetResourceName"`
}

func (NetworkInterfaceResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResourceSettings)(nil)).Elem()
}

func (i NetworkInterfaceResourceSettingsArgs) ToNetworkInterfaceResourceSettingsOutput() NetworkInterfaceResourceSettingsOutput {
	return i.ToNetworkInterfaceResourceSettingsOutputWithContext(context.Background())
}

func (i NetworkInterfaceResourceSettingsArgs) ToNetworkInterfaceResourceSettingsOutputWithContext(ctx context.Context) NetworkInterfaceResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResourceSettingsOutput)
}

type NetworkInterfaceResourceSettingsOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResourceSettings)(nil)).Elem()
}

func (o NetworkInterfaceResourceSettingsOutput) ToNetworkInterfaceResourceSettingsOutput() NetworkInterfaceResourceSettingsOutput {
	return o
}

func (o NetworkInterfaceResourceSettingsOutput) ToNetworkInterfaceResourceSettingsOutputWithContext(ctx context.Context) NetworkInterfaceResourceSettingsOutput {
	return o
}

func (o NetworkInterfaceResourceSettingsOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResourceSettingsOutput) IpConfigurations() NicIpConfigurationResourceSettingsArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) []NicIpConfigurationResourceSettings {
		return v.IpConfigurations
	}).(NicIpConfigurationResourceSettingsArrayOutput)
}

func (o NetworkInterfaceResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkInterfaceResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type NetworkInterfaceResourceSettingsResponse struct {
	EnableAcceleratedNetworking *bool                                        `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            []NicIpConfigurationResourceSettingsResponse `pulumi:"ipConfigurations"`
	ResourceType                string                                       `pulumi:"resourceType"`
	Tags                        map[string]string                            `pulumi:"tags"`
	TargetResourceName          string                                       `pulumi:"targetResourceName"`
}





type NetworkInterfaceResourceSettingsResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceResourceSettingsResponseOutput() NetworkInterfaceResourceSettingsResponseOutput
	ToNetworkInterfaceResourceSettingsResponseOutputWithContext(context.Context) NetworkInterfaceResourceSettingsResponseOutput
}

type NetworkInterfaceResourceSettingsResponseArgs struct {
	EnableAcceleratedNetworking pulumi.BoolPtrInput                                  `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            NicIpConfigurationResourceSettingsResponseArrayInput `pulumi:"ipConfigurations"`
	ResourceType                pulumi.StringInput                                   `pulumi:"resourceType"`
	Tags                        pulumi.StringMapInput                                `pulumi:"tags"`
	TargetResourceName          pulumi.StringInput                                   `pulumi:"targetResourceName"`
}

func (NetworkInterfaceResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResourceSettingsResponse)(nil)).Elem()
}

func (i NetworkInterfaceResourceSettingsResponseArgs) ToNetworkInterfaceResourceSettingsResponseOutput() NetworkInterfaceResourceSettingsResponseOutput {
	return i.ToNetworkInterfaceResourceSettingsResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceResourceSettingsResponseArgs) ToNetworkInterfaceResourceSettingsResponseOutputWithContext(ctx context.Context) NetworkInterfaceResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResourceSettingsResponseOutput)
}

type NetworkInterfaceResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResourceSettingsResponse)(nil)).Elem()
}

func (o NetworkInterfaceResourceSettingsResponseOutput) ToNetworkInterfaceResourceSettingsResponseOutput() NetworkInterfaceResourceSettingsResponseOutput {
	return o
}

func (o NetworkInterfaceResourceSettingsResponseOutput) ToNetworkInterfaceResourceSettingsResponseOutputWithContext(ctx context.Context) NetworkInterfaceResourceSettingsResponseOutput {
	return o
}

func (o NetworkInterfaceResourceSettingsResponseOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResourceSettingsResponseOutput) IpConfigurations() NicIpConfigurationResourceSettingsResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) []NicIpConfigurationResourceSettingsResponse {
		return v.IpConfigurations
	}).(NicIpConfigurationResourceSettingsResponseArrayOutput)
}

func (o NetworkInterfaceResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkInterfaceResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type NetworkSecurityGroupResourceSettings struct {
	ResourceType       string            `pulumi:"resourceType"`
	SecurityRules      []NsgSecurityRule `pulumi:"securityRules"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
}





type NetworkSecurityGroupResourceSettingsInput interface {
	pulumi.Input

	ToNetworkSecurityGroupResourceSettingsOutput() NetworkSecurityGroupResourceSettingsOutput
	ToNetworkSecurityGroupResourceSettingsOutputWithContext(context.Context) NetworkSecurityGroupResourceSettingsOutput
}

type NetworkSecurityGroupResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput        `pulumi:"resourceType"`
	SecurityRules      NsgSecurityRuleArrayInput `pulumi:"securityRules"`
	Tags               pulumi.StringMapInput     `pulumi:"tags"`
	TargetResourceName pulumi.StringInput        `pulumi:"targetResourceName"`
}

func (NetworkSecurityGroupResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResourceSettings)(nil)).Elem()
}

func (i NetworkSecurityGroupResourceSettingsArgs) ToNetworkSecurityGroupResourceSettingsOutput() NetworkSecurityGroupResourceSettingsOutput {
	return i.ToNetworkSecurityGroupResourceSettingsOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupResourceSettingsArgs) ToNetworkSecurityGroupResourceSettingsOutputWithContext(ctx context.Context) NetworkSecurityGroupResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupResourceSettingsOutput)
}

type NetworkSecurityGroupResourceSettingsOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResourceSettings)(nil)).Elem()
}

func (o NetworkSecurityGroupResourceSettingsOutput) ToNetworkSecurityGroupResourceSettingsOutput() NetworkSecurityGroupResourceSettingsOutput {
	return o
}

func (o NetworkSecurityGroupResourceSettingsOutput) ToNetworkSecurityGroupResourceSettingsOutputWithContext(ctx context.Context) NetworkSecurityGroupResourceSettingsOutput {
	return o
}

func (o NetworkSecurityGroupResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupResourceSettingsOutput) SecurityRules() NsgSecurityRuleArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettings) []NsgSecurityRule { return v.SecurityRules }).(NsgSecurityRuleArrayOutput)
}

func (o NetworkSecurityGroupResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityGroupResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type NetworkSecurityGroupResourceSettingsResponse struct {
	ResourceType       string                    `pulumi:"resourceType"`
	SecurityRules      []NsgSecurityRuleResponse `pulumi:"securityRules"`
	Tags               map[string]string         `pulumi:"tags"`
	TargetResourceName string                    `pulumi:"targetResourceName"`
}





type NetworkSecurityGroupResourceSettingsResponseInput interface {
	pulumi.Input

	ToNetworkSecurityGroupResourceSettingsResponseOutput() NetworkSecurityGroupResourceSettingsResponseOutput
	ToNetworkSecurityGroupResourceSettingsResponseOutputWithContext(context.Context) NetworkSecurityGroupResourceSettingsResponseOutput
}

type NetworkSecurityGroupResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput                `pulumi:"resourceType"`
	SecurityRules      NsgSecurityRuleResponseArrayInput `pulumi:"securityRules"`
	Tags               pulumi.StringMapInput             `pulumi:"tags"`
	TargetResourceName pulumi.StringInput                `pulumi:"targetResourceName"`
}

func (NetworkSecurityGroupResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResourceSettingsResponse)(nil)).Elem()
}

func (i NetworkSecurityGroupResourceSettingsResponseArgs) ToNetworkSecurityGroupResourceSettingsResponseOutput() NetworkSecurityGroupResourceSettingsResponseOutput {
	return i.ToNetworkSecurityGroupResourceSettingsResponseOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupResourceSettingsResponseArgs) ToNetworkSecurityGroupResourceSettingsResponseOutputWithContext(ctx context.Context) NetworkSecurityGroupResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupResourceSettingsResponseOutput)
}

type NetworkSecurityGroupResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResourceSettingsResponse)(nil)).Elem()
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) ToNetworkSecurityGroupResourceSettingsResponseOutput() NetworkSecurityGroupResourceSettingsResponseOutput {
	return o
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) ToNetworkSecurityGroupResourceSettingsResponseOutputWithContext(ctx context.Context) NetworkSecurityGroupResourceSettingsResponseOutput {
	return o
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) SecurityRules() NsgSecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettingsResponse) []NsgSecurityRuleResponse { return v.SecurityRules }).(NsgSecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
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





type NicIpConfigurationResourceSettingsInput interface {
	pulumi.Input

	ToNicIpConfigurationResourceSettingsOutput() NicIpConfigurationResourceSettingsOutput
	ToNicIpConfigurationResourceSettingsOutputWithContext(context.Context) NicIpConfigurationResourceSettingsOutput
}

type NicIpConfigurationResourceSettingsArgs struct {
	LoadBalancerBackendAddressPools LoadBalancerBackendAddressPoolReferenceArrayInput `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerNatRules            LoadBalancerNatRuleReferenceArrayInput            `pulumi:"loadBalancerNatRules"`
	Name                            pulumi.StringPtrInput                             `pulumi:"name"`
	Primary                         pulumi.BoolPtrInput                               `pulumi:"primary"`
	PrivateIpAddress                pulumi.StringPtrInput                             `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod       pulumi.StringPtrInput                             `pulumi:"privateIpAllocationMethod"`
	PublicIp                        PublicIpReferencePtrInput                         `pulumi:"publicIp"`
	Subnet                          SubnetReferencePtrInput                           `pulumi:"subnet"`
}

func (NicIpConfigurationResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIpConfigurationResourceSettings)(nil)).Elem()
}

func (i NicIpConfigurationResourceSettingsArgs) ToNicIpConfigurationResourceSettingsOutput() NicIpConfigurationResourceSettingsOutput {
	return i.ToNicIpConfigurationResourceSettingsOutputWithContext(context.Background())
}

func (i NicIpConfigurationResourceSettingsArgs) ToNicIpConfigurationResourceSettingsOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIpConfigurationResourceSettingsOutput)
}





type NicIpConfigurationResourceSettingsArrayInput interface {
	pulumi.Input

	ToNicIpConfigurationResourceSettingsArrayOutput() NicIpConfigurationResourceSettingsArrayOutput
	ToNicIpConfigurationResourceSettingsArrayOutputWithContext(context.Context) NicIpConfigurationResourceSettingsArrayOutput
}

type NicIpConfigurationResourceSettingsArray []NicIpConfigurationResourceSettingsInput

func (NicIpConfigurationResourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIpConfigurationResourceSettings)(nil)).Elem()
}

func (i NicIpConfigurationResourceSettingsArray) ToNicIpConfigurationResourceSettingsArrayOutput() NicIpConfigurationResourceSettingsArrayOutput {
	return i.ToNicIpConfigurationResourceSettingsArrayOutputWithContext(context.Background())
}

func (i NicIpConfigurationResourceSettingsArray) ToNicIpConfigurationResourceSettingsArrayOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIpConfigurationResourceSettingsArrayOutput)
}

type NicIpConfigurationResourceSettingsOutput struct{ *pulumi.OutputState }

func (NicIpConfigurationResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIpConfigurationResourceSettings)(nil)).Elem()
}

func (o NicIpConfigurationResourceSettingsOutput) ToNicIpConfigurationResourceSettingsOutput() NicIpConfigurationResourceSettingsOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsOutput) ToNicIpConfigurationResourceSettingsOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsOutput) LoadBalancerBackendAddressPools() LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) []LoadBalancerBackendAddressPoolReference {
		return v.LoadBalancerBackendAddressPools
	}).(LoadBalancerBackendAddressPoolReferenceArrayOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) LoadBalancerNatRules() LoadBalancerNatRuleReferenceArrayOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) []LoadBalancerNatRuleReference {
		return v.LoadBalancerNatRules
	}).(LoadBalancerNatRuleReferenceArrayOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) PublicIp() PublicIpReferencePtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *PublicIpReference { return v.PublicIp }).(PublicIpReferencePtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) Subnet() SubnetReferencePtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *SubnetReference { return v.Subnet }).(SubnetReferencePtrOutput)
}

type NicIpConfigurationResourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (NicIpConfigurationResourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIpConfigurationResourceSettings)(nil)).Elem()
}

func (o NicIpConfigurationResourceSettingsArrayOutput) ToNicIpConfigurationResourceSettingsArrayOutput() NicIpConfigurationResourceSettingsArrayOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsArrayOutput) ToNicIpConfigurationResourceSettingsArrayOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsArrayOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsArrayOutput) Index(i pulumi.IntInput) NicIpConfigurationResourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NicIpConfigurationResourceSettings {
		return vs[0].([]NicIpConfigurationResourceSettings)[vs[1].(int)]
	}).(NicIpConfigurationResourceSettingsOutput)
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





type NicIpConfigurationResourceSettingsResponseInput interface {
	pulumi.Input

	ToNicIpConfigurationResourceSettingsResponseOutput() NicIpConfigurationResourceSettingsResponseOutput
	ToNicIpConfigurationResourceSettingsResponseOutputWithContext(context.Context) NicIpConfigurationResourceSettingsResponseOutput
}

type NicIpConfigurationResourceSettingsResponseArgs struct {
	LoadBalancerBackendAddressPools LoadBalancerBackendAddressPoolReferenceResponseArrayInput `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerNatRules            LoadBalancerNatRuleReferenceResponseArrayInput            `pulumi:"loadBalancerNatRules"`
	Name                            pulumi.StringPtrInput                                     `pulumi:"name"`
	Primary                         pulumi.BoolPtrInput                                       `pulumi:"primary"`
	PrivateIpAddress                pulumi.StringPtrInput                                     `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod       pulumi.StringPtrInput                                     `pulumi:"privateIpAllocationMethod"`
	PublicIp                        PublicIpReferenceResponsePtrInput                         `pulumi:"publicIp"`
	Subnet                          SubnetReferenceResponsePtrInput                           `pulumi:"subnet"`
}

func (NicIpConfigurationResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIpConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (i NicIpConfigurationResourceSettingsResponseArgs) ToNicIpConfigurationResourceSettingsResponseOutput() NicIpConfigurationResourceSettingsResponseOutput {
	return i.ToNicIpConfigurationResourceSettingsResponseOutputWithContext(context.Background())
}

func (i NicIpConfigurationResourceSettingsResponseArgs) ToNicIpConfigurationResourceSettingsResponseOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIpConfigurationResourceSettingsResponseOutput)
}





type NicIpConfigurationResourceSettingsResponseArrayInput interface {
	pulumi.Input

	ToNicIpConfigurationResourceSettingsResponseArrayOutput() NicIpConfigurationResourceSettingsResponseArrayOutput
	ToNicIpConfigurationResourceSettingsResponseArrayOutputWithContext(context.Context) NicIpConfigurationResourceSettingsResponseArrayOutput
}

type NicIpConfigurationResourceSettingsResponseArray []NicIpConfigurationResourceSettingsResponseInput

func (NicIpConfigurationResourceSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIpConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (i NicIpConfigurationResourceSettingsResponseArray) ToNicIpConfigurationResourceSettingsResponseArrayOutput() NicIpConfigurationResourceSettingsResponseArrayOutput {
	return i.ToNicIpConfigurationResourceSettingsResponseArrayOutputWithContext(context.Background())
}

func (i NicIpConfigurationResourceSettingsResponseArray) ToNicIpConfigurationResourceSettingsResponseArrayOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIpConfigurationResourceSettingsResponseArrayOutput)
}

type NicIpConfigurationResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (NicIpConfigurationResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIpConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (o NicIpConfigurationResourceSettingsResponseOutput) ToNicIpConfigurationResourceSettingsResponseOutput() NicIpConfigurationResourceSettingsResponseOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsResponseOutput) ToNicIpConfigurationResourceSettingsResponseOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsResponseOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsResponseOutput) LoadBalancerBackendAddressPools() LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) []LoadBalancerBackendAddressPoolReferenceResponse {
		return v.LoadBalancerBackendAddressPools
	}).(LoadBalancerBackendAddressPoolReferenceResponseArrayOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) LoadBalancerNatRules() LoadBalancerNatRuleReferenceResponseArrayOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) []LoadBalancerNatRuleReferenceResponse {
		return v.LoadBalancerNatRules
	}).(LoadBalancerNatRuleReferenceResponseArrayOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) PublicIp() PublicIpReferenceResponsePtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *PublicIpReferenceResponse { return v.PublicIp }).(PublicIpReferenceResponsePtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) Subnet() SubnetReferenceResponsePtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *SubnetReferenceResponse { return v.Subnet }).(SubnetReferenceResponsePtrOutput)
}

type NicIpConfigurationResourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (NicIpConfigurationResourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIpConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (o NicIpConfigurationResourceSettingsResponseArrayOutput) ToNicIpConfigurationResourceSettingsResponseArrayOutput() NicIpConfigurationResourceSettingsResponseArrayOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsResponseArrayOutput) ToNicIpConfigurationResourceSettingsResponseArrayOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsResponseArrayOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) NicIpConfigurationResourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NicIpConfigurationResourceSettingsResponse {
		return vs[0].([]NicIpConfigurationResourceSettingsResponse)[vs[1].(int)]
	}).(NicIpConfigurationResourceSettingsResponseOutput)
}

type NsgReference struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}





type NsgReferenceInput interface {
	pulumi.Input

	ToNsgReferenceOutput() NsgReferenceOutput
	ToNsgReferenceOutputWithContext(context.Context) NsgReferenceOutput
}

type NsgReferenceArgs struct {
	SourceArmResourceId pulumi.StringInput `pulumi:"sourceArmResourceId"`
}

func (NsgReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgReference)(nil)).Elem()
}

func (i NsgReferenceArgs) ToNsgReferenceOutput() NsgReferenceOutput {
	return i.ToNsgReferenceOutputWithContext(context.Background())
}

func (i NsgReferenceArgs) ToNsgReferenceOutputWithContext(ctx context.Context) NsgReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceOutput)
}

func (i NsgReferenceArgs) ToNsgReferencePtrOutput() NsgReferencePtrOutput {
	return i.ToNsgReferencePtrOutputWithContext(context.Background())
}

func (i NsgReferenceArgs) ToNsgReferencePtrOutputWithContext(ctx context.Context) NsgReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceOutput).ToNsgReferencePtrOutputWithContext(ctx)
}









type NsgReferencePtrInput interface {
	pulumi.Input

	ToNsgReferencePtrOutput() NsgReferencePtrOutput
	ToNsgReferencePtrOutputWithContext(context.Context) NsgReferencePtrOutput
}

type nsgReferencePtrType NsgReferenceArgs

func NsgReferencePtr(v *NsgReferenceArgs) NsgReferencePtrInput {
	return (*nsgReferencePtrType)(v)
}

func (*nsgReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgReference)(nil)).Elem()
}

func (i *nsgReferencePtrType) ToNsgReferencePtrOutput() NsgReferencePtrOutput {
	return i.ToNsgReferencePtrOutputWithContext(context.Background())
}

func (i *nsgReferencePtrType) ToNsgReferencePtrOutputWithContext(ctx context.Context) NsgReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferencePtrOutput)
}

type NsgReferenceOutput struct{ *pulumi.OutputState }

func (NsgReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgReference)(nil)).Elem()
}

func (o NsgReferenceOutput) ToNsgReferenceOutput() NsgReferenceOutput {
	return o
}

func (o NsgReferenceOutput) ToNsgReferenceOutputWithContext(ctx context.Context) NsgReferenceOutput {
	return o
}

func (o NsgReferenceOutput) ToNsgReferencePtrOutput() NsgReferencePtrOutput {
	return o.ToNsgReferencePtrOutputWithContext(context.Background())
}

func (o NsgReferenceOutput) ToNsgReferencePtrOutputWithContext(ctx context.Context) NsgReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NsgReference) *NsgReference {
		return &v
	}).(NsgReferencePtrOutput)
}

func (o NsgReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v NsgReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type NsgReferencePtrOutput struct{ *pulumi.OutputState }

func (NsgReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgReference)(nil)).Elem()
}

func (o NsgReferencePtrOutput) ToNsgReferencePtrOutput() NsgReferencePtrOutput {
	return o
}

func (o NsgReferencePtrOutput) ToNsgReferencePtrOutputWithContext(ctx context.Context) NsgReferencePtrOutput {
	return o
}

func (o NsgReferencePtrOutput) Elem() NsgReferenceOutput {
	return o.ApplyT(func(v *NsgReference) NsgReference {
		if v != nil {
			return *v
		}
		var ret NsgReference
		return ret
	}).(NsgReferenceOutput)
}

func (o NsgReferencePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsgReference) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type NsgReferenceResponse struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}





type NsgReferenceResponseInput interface {
	pulumi.Input

	ToNsgReferenceResponseOutput() NsgReferenceResponseOutput
	ToNsgReferenceResponseOutputWithContext(context.Context) NsgReferenceResponseOutput
}

type NsgReferenceResponseArgs struct {
	SourceArmResourceId pulumi.StringInput `pulumi:"sourceArmResourceId"`
}

func (NsgReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgReferenceResponse)(nil)).Elem()
}

func (i NsgReferenceResponseArgs) ToNsgReferenceResponseOutput() NsgReferenceResponseOutput {
	return i.ToNsgReferenceResponseOutputWithContext(context.Background())
}

func (i NsgReferenceResponseArgs) ToNsgReferenceResponseOutputWithContext(ctx context.Context) NsgReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceResponseOutput)
}

func (i NsgReferenceResponseArgs) ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput {
	return i.ToNsgReferenceResponsePtrOutputWithContext(context.Background())
}

func (i NsgReferenceResponseArgs) ToNsgReferenceResponsePtrOutputWithContext(ctx context.Context) NsgReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceResponseOutput).ToNsgReferenceResponsePtrOutputWithContext(ctx)
}









type NsgReferenceResponsePtrInput interface {
	pulumi.Input

	ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput
	ToNsgReferenceResponsePtrOutputWithContext(context.Context) NsgReferenceResponsePtrOutput
}

type nsgReferenceResponsePtrType NsgReferenceResponseArgs

func NsgReferenceResponsePtr(v *NsgReferenceResponseArgs) NsgReferenceResponsePtrInput {
	return (*nsgReferenceResponsePtrType)(v)
}

func (*nsgReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgReferenceResponse)(nil)).Elem()
}

func (i *nsgReferenceResponsePtrType) ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput {
	return i.ToNsgReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *nsgReferenceResponsePtrType) ToNsgReferenceResponsePtrOutputWithContext(ctx context.Context) NsgReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceResponsePtrOutput)
}

type NsgReferenceResponseOutput struct{ *pulumi.OutputState }

func (NsgReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgReferenceResponse)(nil)).Elem()
}

func (o NsgReferenceResponseOutput) ToNsgReferenceResponseOutput() NsgReferenceResponseOutput {
	return o
}

func (o NsgReferenceResponseOutput) ToNsgReferenceResponseOutputWithContext(ctx context.Context) NsgReferenceResponseOutput {
	return o
}

func (o NsgReferenceResponseOutput) ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput {
	return o.ToNsgReferenceResponsePtrOutputWithContext(context.Background())
}

func (o NsgReferenceResponseOutput) ToNsgReferenceResponsePtrOutputWithContext(ctx context.Context) NsgReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NsgReferenceResponse) *NsgReferenceResponse {
		return &v
	}).(NsgReferenceResponsePtrOutput)
}

func (o NsgReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v NsgReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type NsgReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (NsgReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgReferenceResponse)(nil)).Elem()
}

func (o NsgReferenceResponsePtrOutput) ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput {
	return o
}

func (o NsgReferenceResponsePtrOutput) ToNsgReferenceResponsePtrOutputWithContext(ctx context.Context) NsgReferenceResponsePtrOutput {
	return o
}

func (o NsgReferenceResponsePtrOutput) Elem() NsgReferenceResponseOutput {
	return o.ApplyT(func(v *NsgReferenceResponse) NsgReferenceResponse {
		if v != nil {
			return *v
		}
		var ret NsgReferenceResponse
		return ret
	}).(NsgReferenceResponseOutput)
}

func (o NsgReferenceResponsePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsgReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
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





type NsgSecurityRuleInput interface {
	pulumi.Input

	ToNsgSecurityRuleOutput() NsgSecurityRuleOutput
	ToNsgSecurityRuleOutputWithContext(context.Context) NsgSecurityRuleOutput
}

type NsgSecurityRuleArgs struct {
	Access                   pulumi.StringPtrInput `pulumi:"access"`
	Description              pulumi.StringPtrInput `pulumi:"description"`
	DestinationAddressPrefix pulumi.StringPtrInput `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     pulumi.StringPtrInput `pulumi:"destinationPortRange"`
	Direction                pulumi.StringPtrInput `pulumi:"direction"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	Priority                 pulumi.IntPtrInput    `pulumi:"priority"`
	Protocol                 pulumi.StringPtrInput `pulumi:"protocol"`
	SourceAddressPrefix      pulumi.StringPtrInput `pulumi:"sourceAddressPrefix"`
	SourcePortRange          pulumi.StringPtrInput `pulumi:"sourcePortRange"`
}

func (NsgSecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgSecurityRule)(nil)).Elem()
}

func (i NsgSecurityRuleArgs) ToNsgSecurityRuleOutput() NsgSecurityRuleOutput {
	return i.ToNsgSecurityRuleOutputWithContext(context.Background())
}

func (i NsgSecurityRuleArgs) ToNsgSecurityRuleOutputWithContext(ctx context.Context) NsgSecurityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgSecurityRuleOutput)
}





type NsgSecurityRuleArrayInput interface {
	pulumi.Input

	ToNsgSecurityRuleArrayOutput() NsgSecurityRuleArrayOutput
	ToNsgSecurityRuleArrayOutputWithContext(context.Context) NsgSecurityRuleArrayOutput
}

type NsgSecurityRuleArray []NsgSecurityRuleInput

func (NsgSecurityRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsgSecurityRule)(nil)).Elem()
}

func (i NsgSecurityRuleArray) ToNsgSecurityRuleArrayOutput() NsgSecurityRuleArrayOutput {
	return i.ToNsgSecurityRuleArrayOutputWithContext(context.Background())
}

func (i NsgSecurityRuleArray) ToNsgSecurityRuleArrayOutputWithContext(ctx context.Context) NsgSecurityRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgSecurityRuleArrayOutput)
}

type NsgSecurityRuleOutput struct{ *pulumi.OutputState }

func (NsgSecurityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgSecurityRule)(nil)).Elem()
}

func (o NsgSecurityRuleOutput) ToNsgSecurityRuleOutput() NsgSecurityRuleOutput {
	return o
}

func (o NsgSecurityRuleOutput) ToNsgSecurityRuleOutputWithContext(ctx context.Context) NsgSecurityRuleOutput {
	return o
}

func (o NsgSecurityRuleOutput) Access() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Access }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o NsgSecurityRuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

type NsgSecurityRuleArrayOutput struct{ *pulumi.OutputState }

func (NsgSecurityRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsgSecurityRule)(nil)).Elem()
}

func (o NsgSecurityRuleArrayOutput) ToNsgSecurityRuleArrayOutput() NsgSecurityRuleArrayOutput {
	return o
}

func (o NsgSecurityRuleArrayOutput) ToNsgSecurityRuleArrayOutputWithContext(ctx context.Context) NsgSecurityRuleArrayOutput {
	return o
}

func (o NsgSecurityRuleArrayOutput) Index(i pulumi.IntInput) NsgSecurityRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NsgSecurityRule {
		return vs[0].([]NsgSecurityRule)[vs[1].(int)]
	}).(NsgSecurityRuleOutput)
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





type NsgSecurityRuleResponseInput interface {
	pulumi.Input

	ToNsgSecurityRuleResponseOutput() NsgSecurityRuleResponseOutput
	ToNsgSecurityRuleResponseOutputWithContext(context.Context) NsgSecurityRuleResponseOutput
}

type NsgSecurityRuleResponseArgs struct {
	Access                   pulumi.StringPtrInput `pulumi:"access"`
	Description              pulumi.StringPtrInput `pulumi:"description"`
	DestinationAddressPrefix pulumi.StringPtrInput `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     pulumi.StringPtrInput `pulumi:"destinationPortRange"`
	Direction                pulumi.StringPtrInput `pulumi:"direction"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	Priority                 pulumi.IntPtrInput    `pulumi:"priority"`
	Protocol                 pulumi.StringPtrInput `pulumi:"protocol"`
	SourceAddressPrefix      pulumi.StringPtrInput `pulumi:"sourceAddressPrefix"`
	SourcePortRange          pulumi.StringPtrInput `pulumi:"sourcePortRange"`
}

func (NsgSecurityRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgSecurityRuleResponse)(nil)).Elem()
}

func (i NsgSecurityRuleResponseArgs) ToNsgSecurityRuleResponseOutput() NsgSecurityRuleResponseOutput {
	return i.ToNsgSecurityRuleResponseOutputWithContext(context.Background())
}

func (i NsgSecurityRuleResponseArgs) ToNsgSecurityRuleResponseOutputWithContext(ctx context.Context) NsgSecurityRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgSecurityRuleResponseOutput)
}





type NsgSecurityRuleResponseArrayInput interface {
	pulumi.Input

	ToNsgSecurityRuleResponseArrayOutput() NsgSecurityRuleResponseArrayOutput
	ToNsgSecurityRuleResponseArrayOutputWithContext(context.Context) NsgSecurityRuleResponseArrayOutput
}

type NsgSecurityRuleResponseArray []NsgSecurityRuleResponseInput

func (NsgSecurityRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsgSecurityRuleResponse)(nil)).Elem()
}

func (i NsgSecurityRuleResponseArray) ToNsgSecurityRuleResponseArrayOutput() NsgSecurityRuleResponseArrayOutput {
	return i.ToNsgSecurityRuleResponseArrayOutputWithContext(context.Background())
}

func (i NsgSecurityRuleResponseArray) ToNsgSecurityRuleResponseArrayOutputWithContext(ctx context.Context) NsgSecurityRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgSecurityRuleResponseArrayOutput)
}

type NsgSecurityRuleResponseOutput struct{ *pulumi.OutputState }

func (NsgSecurityRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgSecurityRuleResponse)(nil)).Elem()
}

func (o NsgSecurityRuleResponseOutput) ToNsgSecurityRuleResponseOutput() NsgSecurityRuleResponseOutput {
	return o
}

func (o NsgSecurityRuleResponseOutput) ToNsgSecurityRuleResponseOutputWithContext(ctx context.Context) NsgSecurityRuleResponseOutput {
	return o
}

func (o NsgSecurityRuleResponseOutput) Access() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Access }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

type NsgSecurityRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (NsgSecurityRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsgSecurityRuleResponse)(nil)).Elem()
}

func (o NsgSecurityRuleResponseArrayOutput) ToNsgSecurityRuleResponseArrayOutput() NsgSecurityRuleResponseArrayOutput {
	return o
}

func (o NsgSecurityRuleResponseArrayOutput) ToNsgSecurityRuleResponseArrayOutputWithContext(ctx context.Context) NsgSecurityRuleResponseArrayOutput {
	return o
}

func (o NsgSecurityRuleResponseArrayOutput) Index(i pulumi.IntInput) NsgSecurityRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NsgSecurityRuleResponse {
		return vs[0].([]NsgSecurityRuleResponse)[vs[1].(int)]
	}).(NsgSecurityRuleResponseOutput)
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





type PublicIPAddressResourceSettingsInput interface {
	pulumi.Input

	ToPublicIPAddressResourceSettingsOutput() PublicIPAddressResourceSettingsOutput
	ToPublicIPAddressResourceSettingsOutputWithContext(context.Context) PublicIPAddressResourceSettingsOutput
}

type PublicIPAddressResourceSettingsArgs struct {
	DomainNameLabel          pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn                     pulumi.StringPtrInput `pulumi:"fqdn"`
	PublicIpAllocationMethod pulumi.StringPtrInput `pulumi:"publicIpAllocationMethod"`
	ResourceType             pulumi.StringInput    `pulumi:"resourceType"`
	Sku                      pulumi.StringPtrInput `pulumi:"sku"`
	Tags                     pulumi.StringMapInput `pulumi:"tags"`
	TargetResourceName       pulumi.StringInput    `pulumi:"targetResourceName"`
	Zones                    pulumi.StringPtrInput `pulumi:"zones"`
}

func (PublicIPAddressResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResourceSettings)(nil)).Elem()
}

func (i PublicIPAddressResourceSettingsArgs) ToPublicIPAddressResourceSettingsOutput() PublicIPAddressResourceSettingsOutput {
	return i.ToPublicIPAddressResourceSettingsOutputWithContext(context.Background())
}

func (i PublicIPAddressResourceSettingsArgs) ToPublicIPAddressResourceSettingsOutputWithContext(ctx context.Context) PublicIPAddressResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressResourceSettingsOutput)
}

type PublicIPAddressResourceSettingsOutput struct{ *pulumi.OutputState }

func (PublicIPAddressResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResourceSettings)(nil)).Elem()
}

func (o PublicIPAddressResourceSettingsOutput) ToPublicIPAddressResourceSettingsOutput() PublicIPAddressResourceSettingsOutput {
	return o
}

func (o PublicIPAddressResourceSettingsOutput) ToPublicIPAddressResourceSettingsOutputWithContext(ctx context.Context) PublicIPAddressResourceSettingsOutput {
	return o
}

func (o PublicIPAddressResourceSettingsOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsOutput) PublicIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.PublicIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o PublicIPAddressResourceSettingsOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PublicIPAddressResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o PublicIPAddressResourceSettingsOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.Zones }).(pulumi.StringPtrOutput)
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





type PublicIPAddressResourceSettingsResponseInput interface {
	pulumi.Input

	ToPublicIPAddressResourceSettingsResponseOutput() PublicIPAddressResourceSettingsResponseOutput
	ToPublicIPAddressResourceSettingsResponseOutputWithContext(context.Context) PublicIPAddressResourceSettingsResponseOutput
}

type PublicIPAddressResourceSettingsResponseArgs struct {
	DomainNameLabel          pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn                     pulumi.StringPtrInput `pulumi:"fqdn"`
	PublicIpAllocationMethod pulumi.StringPtrInput `pulumi:"publicIpAllocationMethod"`
	ResourceType             pulumi.StringInput    `pulumi:"resourceType"`
	Sku                      pulumi.StringPtrInput `pulumi:"sku"`
	Tags                     pulumi.StringMapInput `pulumi:"tags"`
	TargetResourceName       pulumi.StringInput    `pulumi:"targetResourceName"`
	Zones                    pulumi.StringPtrInput `pulumi:"zones"`
}

func (PublicIPAddressResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResourceSettingsResponse)(nil)).Elem()
}

func (i PublicIPAddressResourceSettingsResponseArgs) ToPublicIPAddressResourceSettingsResponseOutput() PublicIPAddressResourceSettingsResponseOutput {
	return i.ToPublicIPAddressResourceSettingsResponseOutputWithContext(context.Background())
}

func (i PublicIPAddressResourceSettingsResponseArgs) ToPublicIPAddressResourceSettingsResponseOutputWithContext(ctx context.Context) PublicIPAddressResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressResourceSettingsResponseOutput)
}

type PublicIPAddressResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResourceSettingsResponse)(nil)).Elem()
}

func (o PublicIPAddressResourceSettingsResponseOutput) ToPublicIPAddressResourceSettingsResponseOutput() PublicIPAddressResourceSettingsResponseOutput {
	return o
}

func (o PublicIPAddressResourceSettingsResponseOutput) ToPublicIPAddressResourceSettingsResponseOutputWithContext(ctx context.Context) PublicIPAddressResourceSettingsResponseOutput {
	return o
}

func (o PublicIPAddressResourceSettingsResponseOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) PublicIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.PublicIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type PublicIpReference struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}





type PublicIpReferenceInput interface {
	pulumi.Input

	ToPublicIpReferenceOutput() PublicIpReferenceOutput
	ToPublicIpReferenceOutputWithContext(context.Context) PublicIpReferenceOutput
}

type PublicIpReferenceArgs struct {
	SourceArmResourceId pulumi.StringInput `pulumi:"sourceArmResourceId"`
}

func (PublicIpReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpReference)(nil)).Elem()
}

func (i PublicIpReferenceArgs) ToPublicIpReferenceOutput() PublicIpReferenceOutput {
	return i.ToPublicIpReferenceOutputWithContext(context.Background())
}

func (i PublicIpReferenceArgs) ToPublicIpReferenceOutputWithContext(ctx context.Context) PublicIpReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceOutput)
}

func (i PublicIpReferenceArgs) ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput {
	return i.ToPublicIpReferencePtrOutputWithContext(context.Background())
}

func (i PublicIpReferenceArgs) ToPublicIpReferencePtrOutputWithContext(ctx context.Context) PublicIpReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceOutput).ToPublicIpReferencePtrOutputWithContext(ctx)
}









type PublicIpReferencePtrInput interface {
	pulumi.Input

	ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput
	ToPublicIpReferencePtrOutputWithContext(context.Context) PublicIpReferencePtrOutput
}

type publicIpReferencePtrType PublicIpReferenceArgs

func PublicIpReferencePtr(v *PublicIpReferenceArgs) PublicIpReferencePtrInput {
	return (*publicIpReferencePtrType)(v)
}

func (*publicIpReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpReference)(nil)).Elem()
}

func (i *publicIpReferencePtrType) ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput {
	return i.ToPublicIpReferencePtrOutputWithContext(context.Background())
}

func (i *publicIpReferencePtrType) ToPublicIpReferencePtrOutputWithContext(ctx context.Context) PublicIpReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferencePtrOutput)
}

type PublicIpReferenceOutput struct{ *pulumi.OutputState }

func (PublicIpReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpReference)(nil)).Elem()
}

func (o PublicIpReferenceOutput) ToPublicIpReferenceOutput() PublicIpReferenceOutput {
	return o
}

func (o PublicIpReferenceOutput) ToPublicIpReferenceOutputWithContext(ctx context.Context) PublicIpReferenceOutput {
	return o
}

func (o PublicIpReferenceOutput) ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput {
	return o.ToPublicIpReferencePtrOutputWithContext(context.Background())
}

func (o PublicIpReferenceOutput) ToPublicIpReferencePtrOutputWithContext(ctx context.Context) PublicIpReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIpReference) *PublicIpReference {
		return &v
	}).(PublicIpReferencePtrOutput)
}

func (o PublicIpReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIpReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type PublicIpReferencePtrOutput struct{ *pulumi.OutputState }

func (PublicIpReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpReference)(nil)).Elem()
}

func (o PublicIpReferencePtrOutput) ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput {
	return o
}

func (o PublicIpReferencePtrOutput) ToPublicIpReferencePtrOutputWithContext(ctx context.Context) PublicIpReferencePtrOutput {
	return o
}

func (o PublicIpReferencePtrOutput) Elem() PublicIpReferenceOutput {
	return o.ApplyT(func(v *PublicIpReference) PublicIpReference {
		if v != nil {
			return *v
		}
		var ret PublicIpReference
		return ret
	}).(PublicIpReferenceOutput)
}

func (o PublicIpReferencePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpReference) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type PublicIpReferenceResponse struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}





type PublicIpReferenceResponseInput interface {
	pulumi.Input

	ToPublicIpReferenceResponseOutput() PublicIpReferenceResponseOutput
	ToPublicIpReferenceResponseOutputWithContext(context.Context) PublicIpReferenceResponseOutput
}

type PublicIpReferenceResponseArgs struct {
	SourceArmResourceId pulumi.StringInput `pulumi:"sourceArmResourceId"`
}

func (PublicIpReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpReferenceResponse)(nil)).Elem()
}

func (i PublicIpReferenceResponseArgs) ToPublicIpReferenceResponseOutput() PublicIpReferenceResponseOutput {
	return i.ToPublicIpReferenceResponseOutputWithContext(context.Background())
}

func (i PublicIpReferenceResponseArgs) ToPublicIpReferenceResponseOutputWithContext(ctx context.Context) PublicIpReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceResponseOutput)
}

func (i PublicIpReferenceResponseArgs) ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput {
	return i.ToPublicIpReferenceResponsePtrOutputWithContext(context.Background())
}

func (i PublicIpReferenceResponseArgs) ToPublicIpReferenceResponsePtrOutputWithContext(ctx context.Context) PublicIpReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceResponseOutput).ToPublicIpReferenceResponsePtrOutputWithContext(ctx)
}









type PublicIpReferenceResponsePtrInput interface {
	pulumi.Input

	ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput
	ToPublicIpReferenceResponsePtrOutputWithContext(context.Context) PublicIpReferenceResponsePtrOutput
}

type publicIpReferenceResponsePtrType PublicIpReferenceResponseArgs

func PublicIpReferenceResponsePtr(v *PublicIpReferenceResponseArgs) PublicIpReferenceResponsePtrInput {
	return (*publicIpReferenceResponsePtrType)(v)
}

func (*publicIpReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpReferenceResponse)(nil)).Elem()
}

func (i *publicIpReferenceResponsePtrType) ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput {
	return i.ToPublicIpReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *publicIpReferenceResponsePtrType) ToPublicIpReferenceResponsePtrOutputWithContext(ctx context.Context) PublicIpReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceResponsePtrOutput)
}

type PublicIpReferenceResponseOutput struct{ *pulumi.OutputState }

func (PublicIpReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpReferenceResponse)(nil)).Elem()
}

func (o PublicIpReferenceResponseOutput) ToPublicIpReferenceResponseOutput() PublicIpReferenceResponseOutput {
	return o
}

func (o PublicIpReferenceResponseOutput) ToPublicIpReferenceResponseOutputWithContext(ctx context.Context) PublicIpReferenceResponseOutput {
	return o
}

func (o PublicIpReferenceResponseOutput) ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput {
	return o.ToPublicIpReferenceResponsePtrOutputWithContext(context.Background())
}

func (o PublicIpReferenceResponseOutput) ToPublicIpReferenceResponsePtrOutputWithContext(ctx context.Context) PublicIpReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIpReferenceResponse) *PublicIpReferenceResponse {
		return &v
	}).(PublicIpReferenceResponsePtrOutput)
}

func (o PublicIpReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIpReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type PublicIpReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIpReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpReferenceResponse)(nil)).Elem()
}

func (o PublicIpReferenceResponsePtrOutput) ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput {
	return o
}

func (o PublicIpReferenceResponsePtrOutput) ToPublicIpReferenceResponsePtrOutputWithContext(ctx context.Context) PublicIpReferenceResponsePtrOutput {
	return o
}

func (o PublicIpReferenceResponsePtrOutput) Elem() PublicIpReferenceResponseOutput {
	return o.ApplyT(func(v *PublicIpReferenceResponse) PublicIpReferenceResponse {
		if v != nil {
			return *v
		}
		var ret PublicIpReferenceResponse
		return ret
	}).(PublicIpReferenceResponseOutput)
}

func (o PublicIpReferenceResponsePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceGroupResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type ResourceGroupResourceSettingsInput interface {
	pulumi.Input

	ToResourceGroupResourceSettingsOutput() ResourceGroupResourceSettingsOutput
	ToResourceGroupResourceSettingsOutputWithContext(context.Context) ResourceGroupResourceSettingsOutput
}

type ResourceGroupResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (ResourceGroupResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupResourceSettings)(nil)).Elem()
}

func (i ResourceGroupResourceSettingsArgs) ToResourceGroupResourceSettingsOutput() ResourceGroupResourceSettingsOutput {
	return i.ToResourceGroupResourceSettingsOutputWithContext(context.Background())
}

func (i ResourceGroupResourceSettingsArgs) ToResourceGroupResourceSettingsOutputWithContext(ctx context.Context) ResourceGroupResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupResourceSettingsOutput)
}

type ResourceGroupResourceSettingsOutput struct{ *pulumi.OutputState }

func (ResourceGroupResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupResourceSettings)(nil)).Elem()
}

func (o ResourceGroupResourceSettingsOutput) ToResourceGroupResourceSettingsOutput() ResourceGroupResourceSettingsOutput {
	return o
}

func (o ResourceGroupResourceSettingsOutput) ToResourceGroupResourceSettingsOutputWithContext(ctx context.Context) ResourceGroupResourceSettingsOutput {
	return o
}

func (o ResourceGroupResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o ResourceGroupResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type ResourceGroupResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type ResourceGroupResourceSettingsResponseInput interface {
	pulumi.Input

	ToResourceGroupResourceSettingsResponseOutput() ResourceGroupResourceSettingsResponseOutput
	ToResourceGroupResourceSettingsResponseOutputWithContext(context.Context) ResourceGroupResourceSettingsResponseOutput
}

type ResourceGroupResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (ResourceGroupResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupResourceSettingsResponse)(nil)).Elem()
}

func (i ResourceGroupResourceSettingsResponseArgs) ToResourceGroupResourceSettingsResponseOutput() ResourceGroupResourceSettingsResponseOutput {
	return i.ToResourceGroupResourceSettingsResponseOutputWithContext(context.Background())
}

func (i ResourceGroupResourceSettingsResponseArgs) ToResourceGroupResourceSettingsResponseOutputWithContext(ctx context.Context) ResourceGroupResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupResourceSettingsResponseOutput)
}

type ResourceGroupResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (ResourceGroupResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupResourceSettingsResponse)(nil)).Elem()
}

func (o ResourceGroupResourceSettingsResponseOutput) ToResourceGroupResourceSettingsResponseOutput() ResourceGroupResourceSettingsResponseOutput {
	return o
}

func (o ResourceGroupResourceSettingsResponseOutput) ToResourceGroupResourceSettingsResponseOutputWithContext(ctx context.Context) ResourceGroupResourceSettingsResponseOutput {
	return o
}

func (o ResourceGroupResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o ResourceGroupResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type SqlDatabaseResourceSettings struct {
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	ZoneRedundant      *string           `pulumi:"zoneRedundant"`
}





type SqlDatabaseResourceSettingsInput interface {
	pulumi.Input

	ToSqlDatabaseResourceSettingsOutput() SqlDatabaseResourceSettingsOutput
	ToSqlDatabaseResourceSettingsOutputWithContext(context.Context) SqlDatabaseResourceSettingsOutput
}

type SqlDatabaseResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	Tags               pulumi.StringMapInput `pulumi:"tags"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	ZoneRedundant      pulumi.StringPtrInput `pulumi:"zoneRedundant"`
}

func (SqlDatabaseResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResourceSettings)(nil)).Elem()
}

func (i SqlDatabaseResourceSettingsArgs) ToSqlDatabaseResourceSettingsOutput() SqlDatabaseResourceSettingsOutput {
	return i.ToSqlDatabaseResourceSettingsOutputWithContext(context.Background())
}

func (i SqlDatabaseResourceSettingsArgs) ToSqlDatabaseResourceSettingsOutputWithContext(ctx context.Context) SqlDatabaseResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourceSettingsOutput)
}

type SqlDatabaseResourceSettingsOutput struct{ *pulumi.OutputState }

func (SqlDatabaseResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResourceSettings)(nil)).Elem()
}

func (o SqlDatabaseResourceSettingsOutput) ToSqlDatabaseResourceSettingsOutput() SqlDatabaseResourceSettingsOutput {
	return o
}

func (o SqlDatabaseResourceSettingsOutput) ToSqlDatabaseResourceSettingsOutputWithContext(ctx context.Context) SqlDatabaseResourceSettingsOutput {
	return o
}

func (o SqlDatabaseResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlDatabaseResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlDatabaseResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o SqlDatabaseResourceSettingsOutput) ZoneRedundant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettings) *string { return v.ZoneRedundant }).(pulumi.StringPtrOutput)
}

type SqlDatabaseResourceSettingsResponse struct {
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	ZoneRedundant      *string           `pulumi:"zoneRedundant"`
}





type SqlDatabaseResourceSettingsResponseInput interface {
	pulumi.Input

	ToSqlDatabaseResourceSettingsResponseOutput() SqlDatabaseResourceSettingsResponseOutput
	ToSqlDatabaseResourceSettingsResponseOutputWithContext(context.Context) SqlDatabaseResourceSettingsResponseOutput
}

type SqlDatabaseResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	Tags               pulumi.StringMapInput `pulumi:"tags"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	ZoneRedundant      pulumi.StringPtrInput `pulumi:"zoneRedundant"`
}

func (SqlDatabaseResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResourceSettingsResponse)(nil)).Elem()
}

func (i SqlDatabaseResourceSettingsResponseArgs) ToSqlDatabaseResourceSettingsResponseOutput() SqlDatabaseResourceSettingsResponseOutput {
	return i.ToSqlDatabaseResourceSettingsResponseOutputWithContext(context.Background())
}

func (i SqlDatabaseResourceSettingsResponseArgs) ToSqlDatabaseResourceSettingsResponseOutputWithContext(ctx context.Context) SqlDatabaseResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourceSettingsResponseOutput)
}

type SqlDatabaseResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlDatabaseResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResourceSettingsResponse)(nil)).Elem()
}

func (o SqlDatabaseResourceSettingsResponseOutput) ToSqlDatabaseResourceSettingsResponseOutput() SqlDatabaseResourceSettingsResponseOutput {
	return o
}

func (o SqlDatabaseResourceSettingsResponseOutput) ToSqlDatabaseResourceSettingsResponseOutputWithContext(ctx context.Context) SqlDatabaseResourceSettingsResponseOutput {
	return o
}

func (o SqlDatabaseResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlDatabaseResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlDatabaseResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o SqlDatabaseResourceSettingsResponseOutput) ZoneRedundant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettingsResponse) *string { return v.ZoneRedundant }).(pulumi.StringPtrOutput)
}

type SqlElasticPoolResourceSettings struct {
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	ZoneRedundant      *string           `pulumi:"zoneRedundant"`
}





type SqlElasticPoolResourceSettingsInput interface {
	pulumi.Input

	ToSqlElasticPoolResourceSettingsOutput() SqlElasticPoolResourceSettingsOutput
	ToSqlElasticPoolResourceSettingsOutputWithContext(context.Context) SqlElasticPoolResourceSettingsOutput
}

type SqlElasticPoolResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	Tags               pulumi.StringMapInput `pulumi:"tags"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	ZoneRedundant      pulumi.StringPtrInput `pulumi:"zoneRedundant"`
}

func (SqlElasticPoolResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlElasticPoolResourceSettings)(nil)).Elem()
}

func (i SqlElasticPoolResourceSettingsArgs) ToSqlElasticPoolResourceSettingsOutput() SqlElasticPoolResourceSettingsOutput {
	return i.ToSqlElasticPoolResourceSettingsOutputWithContext(context.Background())
}

func (i SqlElasticPoolResourceSettingsArgs) ToSqlElasticPoolResourceSettingsOutputWithContext(ctx context.Context) SqlElasticPoolResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlElasticPoolResourceSettingsOutput)
}

type SqlElasticPoolResourceSettingsOutput struct{ *pulumi.OutputState }

func (SqlElasticPoolResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlElasticPoolResourceSettings)(nil)).Elem()
}

func (o SqlElasticPoolResourceSettingsOutput) ToSqlElasticPoolResourceSettingsOutput() SqlElasticPoolResourceSettingsOutput {
	return o
}

func (o SqlElasticPoolResourceSettingsOutput) ToSqlElasticPoolResourceSettingsOutputWithContext(ctx context.Context) SqlElasticPoolResourceSettingsOutput {
	return o
}

func (o SqlElasticPoolResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlElasticPoolResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlElasticPoolResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o SqlElasticPoolResourceSettingsOutput) ZoneRedundant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettings) *string { return v.ZoneRedundant }).(pulumi.StringPtrOutput)
}

type SqlElasticPoolResourceSettingsResponse struct {
	ResourceType       string            `pulumi:"resourceType"`
	Tags               map[string]string `pulumi:"tags"`
	TargetResourceName string            `pulumi:"targetResourceName"`
	ZoneRedundant      *string           `pulumi:"zoneRedundant"`
}





type SqlElasticPoolResourceSettingsResponseInput interface {
	pulumi.Input

	ToSqlElasticPoolResourceSettingsResponseOutput() SqlElasticPoolResourceSettingsResponseOutput
	ToSqlElasticPoolResourceSettingsResponseOutputWithContext(context.Context) SqlElasticPoolResourceSettingsResponseOutput
}

type SqlElasticPoolResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	Tags               pulumi.StringMapInput `pulumi:"tags"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	ZoneRedundant      pulumi.StringPtrInput `pulumi:"zoneRedundant"`
}

func (SqlElasticPoolResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlElasticPoolResourceSettingsResponse)(nil)).Elem()
}

func (i SqlElasticPoolResourceSettingsResponseArgs) ToSqlElasticPoolResourceSettingsResponseOutput() SqlElasticPoolResourceSettingsResponseOutput {
	return i.ToSqlElasticPoolResourceSettingsResponseOutputWithContext(context.Background())
}

func (i SqlElasticPoolResourceSettingsResponseArgs) ToSqlElasticPoolResourceSettingsResponseOutputWithContext(ctx context.Context) SqlElasticPoolResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlElasticPoolResourceSettingsResponseOutput)
}

type SqlElasticPoolResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlElasticPoolResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlElasticPoolResourceSettingsResponse)(nil)).Elem()
}

func (o SqlElasticPoolResourceSettingsResponseOutput) ToSqlElasticPoolResourceSettingsResponseOutput() SqlElasticPoolResourceSettingsResponseOutput {
	return o
}

func (o SqlElasticPoolResourceSettingsResponseOutput) ToSqlElasticPoolResourceSettingsResponseOutputWithContext(ctx context.Context) SqlElasticPoolResourceSettingsResponseOutput {
	return o
}

func (o SqlElasticPoolResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlElasticPoolResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlElasticPoolResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o SqlElasticPoolResourceSettingsResponseOutput) ZoneRedundant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettingsResponse) *string { return v.ZoneRedundant }).(pulumi.StringPtrOutput)
}

type SqlServerResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type SqlServerResourceSettingsInput interface {
	pulumi.Input

	ToSqlServerResourceSettingsOutput() SqlServerResourceSettingsOutput
	ToSqlServerResourceSettingsOutputWithContext(context.Context) SqlServerResourceSettingsOutput
}

type SqlServerResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (SqlServerResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerResourceSettings)(nil)).Elem()
}

func (i SqlServerResourceSettingsArgs) ToSqlServerResourceSettingsOutput() SqlServerResourceSettingsOutput {
	return i.ToSqlServerResourceSettingsOutputWithContext(context.Background())
}

func (i SqlServerResourceSettingsArgs) ToSqlServerResourceSettingsOutputWithContext(ctx context.Context) SqlServerResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerResourceSettingsOutput)
}

type SqlServerResourceSettingsOutput struct{ *pulumi.OutputState }

func (SqlServerResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerResourceSettings)(nil)).Elem()
}

func (o SqlServerResourceSettingsOutput) ToSqlServerResourceSettingsOutput() SqlServerResourceSettingsOutput {
	return o
}

func (o SqlServerResourceSettingsOutput) ToSqlServerResourceSettingsOutputWithContext(ctx context.Context) SqlServerResourceSettingsOutput {
	return o
}

func (o SqlServerResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlServerResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type SqlServerResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type SqlServerResourceSettingsResponseInput interface {
	pulumi.Input

	ToSqlServerResourceSettingsResponseOutput() SqlServerResourceSettingsResponseOutput
	ToSqlServerResourceSettingsResponseOutputWithContext(context.Context) SqlServerResourceSettingsResponseOutput
}

type SqlServerResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (SqlServerResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerResourceSettingsResponse)(nil)).Elem()
}

func (i SqlServerResourceSettingsResponseArgs) ToSqlServerResourceSettingsResponseOutput() SqlServerResourceSettingsResponseOutput {
	return i.ToSqlServerResourceSettingsResponseOutputWithContext(context.Background())
}

func (i SqlServerResourceSettingsResponseArgs) ToSqlServerResourceSettingsResponseOutputWithContext(ctx context.Context) SqlServerResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerResourceSettingsResponseOutput)
}

type SqlServerResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlServerResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerResourceSettingsResponse)(nil)).Elem()
}

func (o SqlServerResourceSettingsResponseOutput) ToSqlServerResourceSettingsResponseOutput() SqlServerResourceSettingsResponseOutput {
	return o
}

func (o SqlServerResourceSettingsResponseOutput) ToSqlServerResourceSettingsResponseOutputWithContext(ctx context.Context) SqlServerResourceSettingsResponseOutput {
	return o
}

func (o SqlServerResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlServerResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type SubnetReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type SubnetReferenceInput interface {
	pulumi.Input

	ToSubnetReferenceOutput() SubnetReferenceOutput
	ToSubnetReferenceOutputWithContext(context.Context) SubnetReferenceOutput
}

type SubnetReferenceArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (SubnetReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetReference)(nil)).Elem()
}

func (i SubnetReferenceArgs) ToSubnetReferenceOutput() SubnetReferenceOutput {
	return i.ToSubnetReferenceOutputWithContext(context.Background())
}

func (i SubnetReferenceArgs) ToSubnetReferenceOutputWithContext(ctx context.Context) SubnetReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceOutput)
}

func (i SubnetReferenceArgs) ToSubnetReferencePtrOutput() SubnetReferencePtrOutput {
	return i.ToSubnetReferencePtrOutputWithContext(context.Background())
}

func (i SubnetReferenceArgs) ToSubnetReferencePtrOutputWithContext(ctx context.Context) SubnetReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceOutput).ToSubnetReferencePtrOutputWithContext(ctx)
}









type SubnetReferencePtrInput interface {
	pulumi.Input

	ToSubnetReferencePtrOutput() SubnetReferencePtrOutput
	ToSubnetReferencePtrOutputWithContext(context.Context) SubnetReferencePtrOutput
}

type subnetReferencePtrType SubnetReferenceArgs

func SubnetReferencePtr(v *SubnetReferenceArgs) SubnetReferencePtrInput {
	return (*subnetReferencePtrType)(v)
}

func (*subnetReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetReference)(nil)).Elem()
}

func (i *subnetReferencePtrType) ToSubnetReferencePtrOutput() SubnetReferencePtrOutput {
	return i.ToSubnetReferencePtrOutputWithContext(context.Background())
}

func (i *subnetReferencePtrType) ToSubnetReferencePtrOutputWithContext(ctx context.Context) SubnetReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferencePtrOutput)
}

type SubnetReferenceOutput struct{ *pulumi.OutputState }

func (SubnetReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetReference)(nil)).Elem()
}

func (o SubnetReferenceOutput) ToSubnetReferenceOutput() SubnetReferenceOutput {
	return o
}

func (o SubnetReferenceOutput) ToSubnetReferenceOutputWithContext(ctx context.Context) SubnetReferenceOutput {
	return o
}

func (o SubnetReferenceOutput) ToSubnetReferencePtrOutput() SubnetReferencePtrOutput {
	return o.ToSubnetReferencePtrOutputWithContext(context.Background())
}

func (o SubnetReferenceOutput) ToSubnetReferencePtrOutputWithContext(ctx context.Context) SubnetReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetReference) *SubnetReference {
		return &v
	}).(SubnetReferencePtrOutput)
}

func (o SubnetReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type SubnetReferencePtrOutput struct{ *pulumi.OutputState }

func (SubnetReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetReference)(nil)).Elem()
}

func (o SubnetReferencePtrOutput) ToSubnetReferencePtrOutput() SubnetReferencePtrOutput {
	return o
}

func (o SubnetReferencePtrOutput) ToSubnetReferencePtrOutputWithContext(ctx context.Context) SubnetReferencePtrOutput {
	return o
}

func (o SubnetReferencePtrOutput) Elem() SubnetReferenceOutput {
	return o.ApplyT(func(v *SubnetReference) SubnetReference {
		if v != nil {
			return *v
		}
		var ret SubnetReference
		return ret
	}).(SubnetReferenceOutput)
}

func (o SubnetReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetReferencePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetReference) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type SubnetReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type SubnetReferenceResponseInput interface {
	pulumi.Input

	ToSubnetReferenceResponseOutput() SubnetReferenceResponseOutput
	ToSubnetReferenceResponseOutputWithContext(context.Context) SubnetReferenceResponseOutput
}

type SubnetReferenceResponseArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (SubnetReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetReferenceResponse)(nil)).Elem()
}

func (i SubnetReferenceResponseArgs) ToSubnetReferenceResponseOutput() SubnetReferenceResponseOutput {
	return i.ToSubnetReferenceResponseOutputWithContext(context.Background())
}

func (i SubnetReferenceResponseArgs) ToSubnetReferenceResponseOutputWithContext(ctx context.Context) SubnetReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceResponseOutput)
}

func (i SubnetReferenceResponseArgs) ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput {
	return i.ToSubnetReferenceResponsePtrOutputWithContext(context.Background())
}

func (i SubnetReferenceResponseArgs) ToSubnetReferenceResponsePtrOutputWithContext(ctx context.Context) SubnetReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceResponseOutput).ToSubnetReferenceResponsePtrOutputWithContext(ctx)
}









type SubnetReferenceResponsePtrInput interface {
	pulumi.Input

	ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput
	ToSubnetReferenceResponsePtrOutputWithContext(context.Context) SubnetReferenceResponsePtrOutput
}

type subnetReferenceResponsePtrType SubnetReferenceResponseArgs

func SubnetReferenceResponsePtr(v *SubnetReferenceResponseArgs) SubnetReferenceResponsePtrInput {
	return (*subnetReferenceResponsePtrType)(v)
}

func (*subnetReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetReferenceResponse)(nil)).Elem()
}

func (i *subnetReferenceResponsePtrType) ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput {
	return i.ToSubnetReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *subnetReferenceResponsePtrType) ToSubnetReferenceResponsePtrOutputWithContext(ctx context.Context) SubnetReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceResponsePtrOutput)
}

type SubnetReferenceResponseOutput struct{ *pulumi.OutputState }

func (SubnetReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetReferenceResponse)(nil)).Elem()
}

func (o SubnetReferenceResponseOutput) ToSubnetReferenceResponseOutput() SubnetReferenceResponseOutput {
	return o
}

func (o SubnetReferenceResponseOutput) ToSubnetReferenceResponseOutputWithContext(ctx context.Context) SubnetReferenceResponseOutput {
	return o
}

func (o SubnetReferenceResponseOutput) ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput {
	return o.ToSubnetReferenceResponsePtrOutputWithContext(context.Background())
}

func (o SubnetReferenceResponseOutput) ToSubnetReferenceResponsePtrOutputWithContext(ctx context.Context) SubnetReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetReferenceResponse) *SubnetReferenceResponse {
		return &v
	}).(SubnetReferenceResponsePtrOutput)
}

func (o SubnetReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type SubnetReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetReferenceResponse)(nil)).Elem()
}

func (o SubnetReferenceResponsePtrOutput) ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput {
	return o
}

func (o SubnetReferenceResponsePtrOutput) ToSubnetReferenceResponsePtrOutputWithContext(ctx context.Context) SubnetReferenceResponsePtrOutput {
	return o
}

func (o SubnetReferenceResponsePtrOutput) Elem() SubnetReferenceResponseOutput {
	return o.ApplyT(func(v *SubnetReferenceResponse) SubnetReferenceResponse {
		if v != nil {
			return *v
		}
		var ret SubnetReferenceResponse
		return ret
	}).(SubnetReferenceResponseOutput)
}

func (o SubnetReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetReferenceResponsePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type SubnetResourceSettings struct {
	AddressPrefix        *string       `pulumi:"addressPrefix"`
	Name                 *string       `pulumi:"name"`
	NetworkSecurityGroup *NsgReference `pulumi:"networkSecurityGroup"`
}





type SubnetResourceSettingsInput interface {
	pulumi.Input

	ToSubnetResourceSettingsOutput() SubnetResourceSettingsOutput
	ToSubnetResourceSettingsOutputWithContext(context.Context) SubnetResourceSettingsOutput
}

type SubnetResourceSettingsArgs struct {
	AddressPrefix        pulumi.StringPtrInput `pulumi:"addressPrefix"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
	NetworkSecurityGroup NsgReferencePtrInput  `pulumi:"networkSecurityGroup"`
}

func (SubnetResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResourceSettings)(nil)).Elem()
}

func (i SubnetResourceSettingsArgs) ToSubnetResourceSettingsOutput() SubnetResourceSettingsOutput {
	return i.ToSubnetResourceSettingsOutputWithContext(context.Background())
}

func (i SubnetResourceSettingsArgs) ToSubnetResourceSettingsOutputWithContext(ctx context.Context) SubnetResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResourceSettingsOutput)
}





type SubnetResourceSettingsArrayInput interface {
	pulumi.Input

	ToSubnetResourceSettingsArrayOutput() SubnetResourceSettingsArrayOutput
	ToSubnetResourceSettingsArrayOutputWithContext(context.Context) SubnetResourceSettingsArrayOutput
}

type SubnetResourceSettingsArray []SubnetResourceSettingsInput

func (SubnetResourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResourceSettings)(nil)).Elem()
}

func (i SubnetResourceSettingsArray) ToSubnetResourceSettingsArrayOutput() SubnetResourceSettingsArrayOutput {
	return i.ToSubnetResourceSettingsArrayOutputWithContext(context.Background())
}

func (i SubnetResourceSettingsArray) ToSubnetResourceSettingsArrayOutputWithContext(ctx context.Context) SubnetResourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResourceSettingsArrayOutput)
}

type SubnetResourceSettingsOutput struct{ *pulumi.OutputState }

func (SubnetResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResourceSettings)(nil)).Elem()
}

func (o SubnetResourceSettingsOutput) ToSubnetResourceSettingsOutput() SubnetResourceSettingsOutput {
	return o
}

func (o SubnetResourceSettingsOutput) ToSubnetResourceSettingsOutputWithContext(ctx context.Context) SubnetResourceSettingsOutput {
	return o
}

func (o SubnetResourceSettingsOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResourceSettings) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetResourceSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResourceSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetResourceSettingsOutput) NetworkSecurityGroup() NsgReferencePtrOutput {
	return o.ApplyT(func(v SubnetResourceSettings) *NsgReference { return v.NetworkSecurityGroup }).(NsgReferencePtrOutput)
}

type SubnetResourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (SubnetResourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResourceSettings)(nil)).Elem()
}

func (o SubnetResourceSettingsArrayOutput) ToSubnetResourceSettingsArrayOutput() SubnetResourceSettingsArrayOutput {
	return o
}

func (o SubnetResourceSettingsArrayOutput) ToSubnetResourceSettingsArrayOutputWithContext(ctx context.Context) SubnetResourceSettingsArrayOutput {
	return o
}

func (o SubnetResourceSettingsArrayOutput) Index(i pulumi.IntInput) SubnetResourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResourceSettings {
		return vs[0].([]SubnetResourceSettings)[vs[1].(int)]
	}).(SubnetResourceSettingsOutput)
}

type SubnetResourceSettingsResponse struct {
	AddressPrefix        *string               `pulumi:"addressPrefix"`
	Name                 *string               `pulumi:"name"`
	NetworkSecurityGroup *NsgReferenceResponse `pulumi:"networkSecurityGroup"`
}





type SubnetResourceSettingsResponseInput interface {
	pulumi.Input

	ToSubnetResourceSettingsResponseOutput() SubnetResourceSettingsResponseOutput
	ToSubnetResourceSettingsResponseOutputWithContext(context.Context) SubnetResourceSettingsResponseOutput
}

type SubnetResourceSettingsResponseArgs struct {
	AddressPrefix        pulumi.StringPtrInput        `pulumi:"addressPrefix"`
	Name                 pulumi.StringPtrInput        `pulumi:"name"`
	NetworkSecurityGroup NsgReferenceResponsePtrInput `pulumi:"networkSecurityGroup"`
}

func (SubnetResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResourceSettingsResponse)(nil)).Elem()
}

func (i SubnetResourceSettingsResponseArgs) ToSubnetResourceSettingsResponseOutput() SubnetResourceSettingsResponseOutput {
	return i.ToSubnetResourceSettingsResponseOutputWithContext(context.Background())
}

func (i SubnetResourceSettingsResponseArgs) ToSubnetResourceSettingsResponseOutputWithContext(ctx context.Context) SubnetResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResourceSettingsResponseOutput)
}





type SubnetResourceSettingsResponseArrayInput interface {
	pulumi.Input

	ToSubnetResourceSettingsResponseArrayOutput() SubnetResourceSettingsResponseArrayOutput
	ToSubnetResourceSettingsResponseArrayOutputWithContext(context.Context) SubnetResourceSettingsResponseArrayOutput
}

type SubnetResourceSettingsResponseArray []SubnetResourceSettingsResponseInput

func (SubnetResourceSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResourceSettingsResponse)(nil)).Elem()
}

func (i SubnetResourceSettingsResponseArray) ToSubnetResourceSettingsResponseArrayOutput() SubnetResourceSettingsResponseArrayOutput {
	return i.ToSubnetResourceSettingsResponseArrayOutputWithContext(context.Background())
}

func (i SubnetResourceSettingsResponseArray) ToSubnetResourceSettingsResponseArrayOutputWithContext(ctx context.Context) SubnetResourceSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResourceSettingsResponseArrayOutput)
}

type SubnetResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SubnetResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResourceSettingsResponse)(nil)).Elem()
}

func (o SubnetResourceSettingsResponseOutput) ToSubnetResourceSettingsResponseOutput() SubnetResourceSettingsResponseOutput {
	return o
}

func (o SubnetResourceSettingsResponseOutput) ToSubnetResourceSettingsResponseOutputWithContext(ctx context.Context) SubnetResourceSettingsResponseOutput {
	return o
}

func (o SubnetResourceSettingsResponseOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResourceSettingsResponse) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetResourceSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResourceSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetResourceSettingsResponseOutput) NetworkSecurityGroup() NsgReferenceResponsePtrOutput {
	return o.ApplyT(func(v SubnetResourceSettingsResponse) *NsgReferenceResponse { return v.NetworkSecurityGroup }).(NsgReferenceResponsePtrOutput)
}

type SubnetResourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetResourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResourceSettingsResponse)(nil)).Elem()
}

func (o SubnetResourceSettingsResponseArrayOutput) ToSubnetResourceSettingsResponseArrayOutput() SubnetResourceSettingsResponseArrayOutput {
	return o
}

func (o SubnetResourceSettingsResponseArrayOutput) ToSubnetResourceSettingsResponseArrayOutputWithContext(ctx context.Context) SubnetResourceSettingsResponseArrayOutput {
	return o
}

func (o SubnetResourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) SubnetResourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResourceSettingsResponse {
		return vs[0].([]SubnetResourceSettingsResponse)[vs[1].(int)]
	}).(SubnetResourceSettingsResponseOutput)
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

type VirtualMachineResourceSettings struct {
	ResourceType            string            `pulumi:"resourceType"`
	Tags                    map[string]string `pulumi:"tags"`
	TargetAvailabilitySetId *string           `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  *string           `pulumi:"targetAvailabilityZone"`
	TargetResourceName      string            `pulumi:"targetResourceName"`
	TargetVmSize            *string           `pulumi:"targetVmSize"`
	UserManagedIdentities   []string          `pulumi:"userManagedIdentities"`
}





type VirtualMachineResourceSettingsInput interface {
	pulumi.Input

	ToVirtualMachineResourceSettingsOutput() VirtualMachineResourceSettingsOutput
	ToVirtualMachineResourceSettingsOutputWithContext(context.Context) VirtualMachineResourceSettingsOutput
}

type VirtualMachineResourceSettingsArgs struct {
	ResourceType            pulumi.StringInput      `pulumi:"resourceType"`
	Tags                    pulumi.StringMapInput   `pulumi:"tags"`
	TargetAvailabilitySetId pulumi.StringPtrInput   `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  pulumi.StringPtrInput   `pulumi:"targetAvailabilityZone"`
	TargetResourceName      pulumi.StringInput      `pulumi:"targetResourceName"`
	TargetVmSize            pulumi.StringPtrInput   `pulumi:"targetVmSize"`
	UserManagedIdentities   pulumi.StringArrayInput `pulumi:"userManagedIdentities"`
}

func (VirtualMachineResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResourceSettings)(nil)).Elem()
}

func (i VirtualMachineResourceSettingsArgs) ToVirtualMachineResourceSettingsOutput() VirtualMachineResourceSettingsOutput {
	return i.ToVirtualMachineResourceSettingsOutputWithContext(context.Background())
}

func (i VirtualMachineResourceSettingsArgs) ToVirtualMachineResourceSettingsOutputWithContext(ctx context.Context) VirtualMachineResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResourceSettingsOutput)
}

type VirtualMachineResourceSettingsOutput struct{ *pulumi.OutputState }

func (VirtualMachineResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResourceSettings)(nil)).Elem()
}

func (o VirtualMachineResourceSettingsOutput) ToVirtualMachineResourceSettingsOutput() VirtualMachineResourceSettingsOutput {
	return o
}

func (o VirtualMachineResourceSettingsOutput) ToVirtualMachineResourceSettingsOutputWithContext(ctx context.Context) VirtualMachineResourceSettingsOutput {
	return o
}

func (o VirtualMachineResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o VirtualMachineResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineResourceSettingsOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) *string { return v.TargetAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsOutput) TargetAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) *string { return v.TargetAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o VirtualMachineResourceSettingsOutput) TargetVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) *string { return v.TargetVmSize }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsOutput) UserManagedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) []string { return v.UserManagedIdentities }).(pulumi.StringArrayOutput)
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





type VirtualMachineResourceSettingsResponseInput interface {
	pulumi.Input

	ToVirtualMachineResourceSettingsResponseOutput() VirtualMachineResourceSettingsResponseOutput
	ToVirtualMachineResourceSettingsResponseOutputWithContext(context.Context) VirtualMachineResourceSettingsResponseOutput
}

type VirtualMachineResourceSettingsResponseArgs struct {
	ResourceType            pulumi.StringInput      `pulumi:"resourceType"`
	Tags                    pulumi.StringMapInput   `pulumi:"tags"`
	TargetAvailabilitySetId pulumi.StringPtrInput   `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  pulumi.StringPtrInput   `pulumi:"targetAvailabilityZone"`
	TargetResourceName      pulumi.StringInput      `pulumi:"targetResourceName"`
	TargetVmSize            pulumi.StringPtrInput   `pulumi:"targetVmSize"`
	UserManagedIdentities   pulumi.StringArrayInput `pulumi:"userManagedIdentities"`
}

func (VirtualMachineResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResourceSettingsResponse)(nil)).Elem()
}

func (i VirtualMachineResourceSettingsResponseArgs) ToVirtualMachineResourceSettingsResponseOutput() VirtualMachineResourceSettingsResponseOutput {
	return i.ToVirtualMachineResourceSettingsResponseOutputWithContext(context.Background())
}

func (i VirtualMachineResourceSettingsResponseArgs) ToVirtualMachineResourceSettingsResponseOutputWithContext(ctx context.Context) VirtualMachineResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResourceSettingsResponseOutput)
}

type VirtualMachineResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResourceSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineResourceSettingsResponseOutput) ToVirtualMachineResourceSettingsResponseOutput() VirtualMachineResourceSettingsResponseOutput {
	return o
}

func (o VirtualMachineResourceSettingsResponseOutput) ToVirtualMachineResourceSettingsResponseOutputWithContext(ctx context.Context) VirtualMachineResourceSettingsResponseOutput {
	return o
}

func (o VirtualMachineResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) *string { return v.TargetAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) TargetAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) *string { return v.TargetAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) TargetVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) *string { return v.TargetVmSize }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) UserManagedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) []string { return v.UserManagedIdentities }).(pulumi.StringArrayOutput)
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





type VirtualNetworkResourceSettingsInput interface {
	pulumi.Input

	ToVirtualNetworkResourceSettingsOutput() VirtualNetworkResourceSettingsOutput
	ToVirtualNetworkResourceSettingsOutputWithContext(context.Context) VirtualNetworkResourceSettingsOutput
}

type VirtualNetworkResourceSettingsArgs struct {
	AddressSpace         pulumi.StringArrayInput          `pulumi:"addressSpace"`
	DnsServers           pulumi.StringArrayInput          `pulumi:"dnsServers"`
	EnableDdosProtection pulumi.BoolPtrInput              `pulumi:"enableDdosProtection"`
	ResourceType         pulumi.StringInput               `pulumi:"resourceType"`
	Subnets              SubnetResourceSettingsArrayInput `pulumi:"subnets"`
	Tags                 pulumi.StringMapInput            `pulumi:"tags"`
	TargetResourceName   pulumi.StringInput               `pulumi:"targetResourceName"`
}

func (VirtualNetworkResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResourceSettings)(nil)).Elem()
}

func (i VirtualNetworkResourceSettingsArgs) ToVirtualNetworkResourceSettingsOutput() VirtualNetworkResourceSettingsOutput {
	return i.ToVirtualNetworkResourceSettingsOutputWithContext(context.Background())
}

func (i VirtualNetworkResourceSettingsArgs) ToVirtualNetworkResourceSettingsOutputWithContext(ctx context.Context) VirtualNetworkResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkResourceSettingsOutput)
}

type VirtualNetworkResourceSettingsOutput struct{ *pulumi.OutputState }

func (VirtualNetworkResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResourceSettings)(nil)).Elem()
}

func (o VirtualNetworkResourceSettingsOutput) ToVirtualNetworkResourceSettingsOutput() VirtualNetworkResourceSettingsOutput {
	return o
}

func (o VirtualNetworkResourceSettingsOutput) ToVirtualNetworkResourceSettingsOutputWithContext(ctx context.Context) VirtualNetworkResourceSettingsOutput {
	return o
}

func (o VirtualNetworkResourceSettingsOutput) AddressSpace() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) []string { return v.AddressSpace }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworkResourceSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworkResourceSettingsOutput) EnableDdosProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) *bool { return v.EnableDdosProtection }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o VirtualNetworkResourceSettingsOutput) Subnets() SubnetResourceSettingsArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) []SubnetResourceSettings { return v.Subnets }).(SubnetResourceSettingsArrayOutput)
}

func (o VirtualNetworkResourceSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
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





type VirtualNetworkResourceSettingsResponseInput interface {
	pulumi.Input

	ToVirtualNetworkResourceSettingsResponseOutput() VirtualNetworkResourceSettingsResponseOutput
	ToVirtualNetworkResourceSettingsResponseOutputWithContext(context.Context) VirtualNetworkResourceSettingsResponseOutput
}

type VirtualNetworkResourceSettingsResponseArgs struct {
	AddressSpace         pulumi.StringArrayInput                  `pulumi:"addressSpace"`
	DnsServers           pulumi.StringArrayInput                  `pulumi:"dnsServers"`
	EnableDdosProtection pulumi.BoolPtrInput                      `pulumi:"enableDdosProtection"`
	ResourceType         pulumi.StringInput                       `pulumi:"resourceType"`
	Subnets              SubnetResourceSettingsResponseArrayInput `pulumi:"subnets"`
	Tags                 pulumi.StringMapInput                    `pulumi:"tags"`
	TargetResourceName   pulumi.StringInput                       `pulumi:"targetResourceName"`
}

func (VirtualNetworkResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResourceSettingsResponse)(nil)).Elem()
}

func (i VirtualNetworkResourceSettingsResponseArgs) ToVirtualNetworkResourceSettingsResponseOutput() VirtualNetworkResourceSettingsResponseOutput {
	return i.ToVirtualNetworkResourceSettingsResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkResourceSettingsResponseArgs) ToVirtualNetworkResourceSettingsResponseOutputWithContext(ctx context.Context) VirtualNetworkResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkResourceSettingsResponseOutput)
}

type VirtualNetworkResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResourceSettingsResponse)(nil)).Elem()
}

func (o VirtualNetworkResourceSettingsResponseOutput) ToVirtualNetworkResourceSettingsResponseOutput() VirtualNetworkResourceSettingsResponseOutput {
	return o
}

func (o VirtualNetworkResourceSettingsResponseOutput) ToVirtualNetworkResourceSettingsResponseOutputWithContext(ctx context.Context) VirtualNetworkResourceSettingsResponseOutput {
	return o
}

func (o VirtualNetworkResourceSettingsResponseOutput) AddressSpace() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) []string { return v.AddressSpace }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) EnableDdosProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) *bool { return v.EnableDdosProtection }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) Subnets() SubnetResourceSettingsResponseArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) []SubnetResourceSettingsResponse { return v.Subnets }).(SubnetResourceSettingsResponseArrayOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AutomaticResolutionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutomaticResolutionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AvailabilitySetResourceSettingsOutput{})
	pulumi.RegisterOutputType(AvailabilitySetResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetResourceSettingsOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStatusResponseOutput{})
	pulumi.RegisterOutputType(JobStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultResourceSettingsOutput{})
	pulumi.RegisterOutputType(KeyVaultResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(LBBackendAddressPoolResourceSettingsOutput{})
	pulumi.RegisterOutputType(LBBackendAddressPoolResourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(LBBackendAddressPoolResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(LBBackendAddressPoolResourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(LBFrontendIPConfigurationResourceSettingsOutput{})
	pulumi.RegisterOutputType(LBFrontendIPConfigurationResourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(LBFrontendIPConfigurationResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(LBFrontendIPConfigurationResourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolReferenceOutput{})
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolReferenceArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolReferenceResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerNatRuleReferenceOutput{})
	pulumi.RegisterOutputType(LoadBalancerNatRuleReferenceArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerNatRuleReferenceResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerNatRuleReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerResourceSettingsOutput{})
	pulumi.RegisterOutputType(LoadBalancerResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManualResolutionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManualResolutionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponseErrorsOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponseErrorsPtrOutput{})
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
	pulumi.RegisterOutputType(MoveResourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseErrorsOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseErrorsPtrOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseMoveStatusOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseMoveStatusPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResourceSettingsOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupResourceSettingsOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(NicIpConfigurationResourceSettingsOutput{})
	pulumi.RegisterOutputType(NicIpConfigurationResourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(NicIpConfigurationResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(NicIpConfigurationResourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(NsgReferenceOutput{})
	pulumi.RegisterOutputType(NsgReferencePtrOutput{})
	pulumi.RegisterOutputType(NsgReferenceResponseOutput{})
	pulumi.RegisterOutputType(NsgReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(NsgSecurityRuleOutput{})
	pulumi.RegisterOutputType(NsgSecurityRuleArrayOutput{})
	pulumi.RegisterOutputType(NsgSecurityRuleResponseOutput{})
	pulumi.RegisterOutputType(NsgSecurityRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(PublicIPAddressResourceSettingsOutput{})
	pulumi.RegisterOutputType(PublicIPAddressResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(PublicIpReferenceOutput{})
	pulumi.RegisterOutputType(PublicIpReferencePtrOutput{})
	pulumi.RegisterOutputType(PublicIpReferenceResponseOutput{})
	pulumi.RegisterOutputType(PublicIpReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceGroupResourceSettingsOutput{})
	pulumi.RegisterOutputType(ResourceGroupResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SqlDatabaseResourceSettingsOutput{})
	pulumi.RegisterOutputType(SqlDatabaseResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SqlElasticPoolResourceSettingsOutput{})
	pulumi.RegisterOutputType(SqlElasticPoolResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SqlServerResourceSettingsOutput{})
	pulumi.RegisterOutputType(SqlServerResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SubnetReferenceOutput{})
	pulumi.RegisterOutputType(SubnetReferencePtrOutput{})
	pulumi.RegisterOutputType(SubnetReferenceResponseOutput{})
	pulumi.RegisterOutputType(SubnetReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetResourceSettingsOutput{})
	pulumi.RegisterOutputType(SubnetResourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(SubnetResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SubnetResourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineResourceSettingsOutput{})
	pulumi.RegisterOutputType(VirtualMachineResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkResourceSettingsOutput{})
	pulumi.RegisterOutputType(VirtualNetworkResourceSettingsResponseOutput{})
}
