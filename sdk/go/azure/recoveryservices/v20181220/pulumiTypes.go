


package v20181220

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AADProperties struct {
	Audience                 *string `pulumi:"audience"`
	Authority                *string `pulumi:"authority"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	ServicePrincipalObjectId *string `pulumi:"servicePrincipalObjectId"`
	TenantId                 *string `pulumi:"tenantId"`
}





type AADPropertiesInput interface {
	pulumi.Input

	ToAADPropertiesOutput() AADPropertiesOutput
	ToAADPropertiesOutputWithContext(context.Context) AADPropertiesOutput
}

type AADPropertiesArgs struct {
	Audience                 pulumi.StringPtrInput `pulumi:"audience"`
	Authority                pulumi.StringPtrInput `pulumi:"authority"`
	ServicePrincipalClientId pulumi.StringPtrInput `pulumi:"servicePrincipalClientId"`
	ServicePrincipalObjectId pulumi.StringPtrInput `pulumi:"servicePrincipalObjectId"`
	TenantId                 pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AADPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AADProperties)(nil)).Elem()
}

func (i AADPropertiesArgs) ToAADPropertiesOutput() AADPropertiesOutput {
	return i.ToAADPropertiesOutputWithContext(context.Background())
}

func (i AADPropertiesArgs) ToAADPropertiesOutputWithContext(ctx context.Context) AADPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADPropertiesOutput)
}

func (i AADPropertiesArgs) ToAADPropertiesPtrOutput() AADPropertiesPtrOutput {
	return i.ToAADPropertiesPtrOutputWithContext(context.Background())
}

func (i AADPropertiesArgs) ToAADPropertiesPtrOutputWithContext(ctx context.Context) AADPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADPropertiesOutput).ToAADPropertiesPtrOutputWithContext(ctx)
}









type AADPropertiesPtrInput interface {
	pulumi.Input

	ToAADPropertiesPtrOutput() AADPropertiesPtrOutput
	ToAADPropertiesPtrOutputWithContext(context.Context) AADPropertiesPtrOutput
}

type aadpropertiesPtrType AADPropertiesArgs

func AADPropertiesPtr(v *AADPropertiesArgs) AADPropertiesPtrInput {
	return (*aadpropertiesPtrType)(v)
}

func (*aadpropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AADProperties)(nil)).Elem()
}

func (i *aadpropertiesPtrType) ToAADPropertiesPtrOutput() AADPropertiesPtrOutput {
	return i.ToAADPropertiesPtrOutputWithContext(context.Background())
}

func (i *aadpropertiesPtrType) ToAADPropertiesPtrOutputWithContext(ctx context.Context) AADPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADPropertiesPtrOutput)
}

type AADPropertiesOutput struct{ *pulumi.OutputState }

func (AADPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADProperties)(nil)).Elem()
}

func (o AADPropertiesOutput) ToAADPropertiesOutput() AADPropertiesOutput {
	return o
}

func (o AADPropertiesOutput) ToAADPropertiesOutputWithContext(ctx context.Context) AADPropertiesOutput {
	return o
}

func (o AADPropertiesOutput) ToAADPropertiesPtrOutput() AADPropertiesPtrOutput {
	return o.ToAADPropertiesPtrOutputWithContext(context.Background())
}

func (o AADPropertiesOutput) ToAADPropertiesPtrOutputWithContext(ctx context.Context) AADPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AADProperties) *AADProperties {
		return &v
	}).(AADPropertiesPtrOutput)
}

func (o AADPropertiesOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProperties) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o AADPropertiesOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProperties) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o AADPropertiesOutput) ServicePrincipalClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProperties) *string { return v.ServicePrincipalClientId }).(pulumi.StringPtrOutput)
}

func (o AADPropertiesOutput) ServicePrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProperties) *string { return v.ServicePrincipalObjectId }).(pulumi.StringPtrOutput)
}

func (o AADPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AADPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AADPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADProperties)(nil)).Elem()
}

func (o AADPropertiesPtrOutput) ToAADPropertiesPtrOutput() AADPropertiesPtrOutput {
	return o
}

func (o AADPropertiesPtrOutput) ToAADPropertiesPtrOutputWithContext(ctx context.Context) AADPropertiesPtrOutput {
	return o
}

func (o AADPropertiesPtrOutput) Elem() AADPropertiesOutput {
	return o.ApplyT(func(v *AADProperties) AADProperties {
		if v != nil {
			return *v
		}
		var ret AADProperties
		return ret
	}).(AADPropertiesOutput)
}

func (o AADPropertiesPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProperties) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o AADPropertiesPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProperties) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o AADPropertiesPtrOutput) ServicePrincipalClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalClientId
	}).(pulumi.StringPtrOutput)
}

func (o AADPropertiesPtrOutput) ServicePrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalObjectId
	}).(pulumi.StringPtrOutput)
}

func (o AADPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type WorkloadCrrAccessTokenResponse struct {
	AccessTokenString                           *string           `pulumi:"accessTokenString"`
	BMSActiveRegion                             *string           `pulumi:"bMSActiveRegion"`
	BackupManagementType                        *string           `pulumi:"backupManagementType"`
	ContainerId                                 *string           `pulumi:"containerId"`
	ContainerName                               *string           `pulumi:"containerName"`
	ContainerType                               *string           `pulumi:"containerType"`
	CoordinatorServiceStampId                   *string           `pulumi:"coordinatorServiceStampId"`
	CoordinatorServiceStampUri                  *string           `pulumi:"coordinatorServiceStampUri"`
	DatasourceContainerName                     *string           `pulumi:"datasourceContainerName"`
	DatasourceId                                *string           `pulumi:"datasourceId"`
	DatasourceName                              *string           `pulumi:"datasourceName"`
	DatasourceType                              *string           `pulumi:"datasourceType"`
	ObjectType                                  string            `pulumi:"objectType"`
	PolicyId                                    *string           `pulumi:"policyId"`
	PolicyName                                  *string           `pulumi:"policyName"`
	ProtectableObjectContainerHostOsName        *string           `pulumi:"protectableObjectContainerHostOsName"`
	ProtectableObjectFriendlyName               *string           `pulumi:"protectableObjectFriendlyName"`
	ProtectableObjectParentLogicalContainerName *string           `pulumi:"protectableObjectParentLogicalContainerName"`
	ProtectableObjectProtectionState            *string           `pulumi:"protectableObjectProtectionState"`
	ProtectableObjectUniqueName                 *string           `pulumi:"protectableObjectUniqueName"`
	ProtectableObjectWorkloadType               *string           `pulumi:"protectableObjectWorkloadType"`
	ProtectionContainerId                       *float64          `pulumi:"protectionContainerId"`
	ProtectionServiceStampId                    *string           `pulumi:"protectionServiceStampId"`
	ProtectionServiceStampUri                   *string           `pulumi:"protectionServiceStampUri"`
	RecoveryPointId                             *string           `pulumi:"recoveryPointId"`
	RecoveryPointTime                           *string           `pulumi:"recoveryPointTime"`
	ResourceGroupName                           *string           `pulumi:"resourceGroupName"`
	ResourceId                                  *string           `pulumi:"resourceId"`
	ResourceName                                *string           `pulumi:"resourceName"`
	RpIsManagedVirtualMachine                   *bool             `pulumi:"rpIsManagedVirtualMachine"`
	RpOriginalSAOption                          *bool             `pulumi:"rpOriginalSAOption"`
	RpTierInformation                           map[string]string `pulumi:"rpTierInformation"`
	RpVMSizeDescription                         *string           `pulumi:"rpVMSizeDescription"`
	SubscriptionId                              *string           `pulumi:"subscriptionId"`
	TokenExtendedInformation                    *string           `pulumi:"tokenExtendedInformation"`
}





type WorkloadCrrAccessTokenResponseInput interface {
	pulumi.Input

	ToWorkloadCrrAccessTokenResponseOutput() WorkloadCrrAccessTokenResponseOutput
	ToWorkloadCrrAccessTokenResponseOutputWithContext(context.Context) WorkloadCrrAccessTokenResponseOutput
}

type WorkloadCrrAccessTokenResponseArgs struct {
	AccessTokenString                           pulumi.StringPtrInput  `pulumi:"accessTokenString"`
	BMSActiveRegion                             pulumi.StringPtrInput  `pulumi:"bMSActiveRegion"`
	BackupManagementType                        pulumi.StringPtrInput  `pulumi:"backupManagementType"`
	ContainerId                                 pulumi.StringPtrInput  `pulumi:"containerId"`
	ContainerName                               pulumi.StringPtrInput  `pulumi:"containerName"`
	ContainerType                               pulumi.StringPtrInput  `pulumi:"containerType"`
	CoordinatorServiceStampId                   pulumi.StringPtrInput  `pulumi:"coordinatorServiceStampId"`
	CoordinatorServiceStampUri                  pulumi.StringPtrInput  `pulumi:"coordinatorServiceStampUri"`
	DatasourceContainerName                     pulumi.StringPtrInput  `pulumi:"datasourceContainerName"`
	DatasourceId                                pulumi.StringPtrInput  `pulumi:"datasourceId"`
	DatasourceName                              pulumi.StringPtrInput  `pulumi:"datasourceName"`
	DatasourceType                              pulumi.StringPtrInput  `pulumi:"datasourceType"`
	ObjectType                                  pulumi.StringInput     `pulumi:"objectType"`
	PolicyId                                    pulumi.StringPtrInput  `pulumi:"policyId"`
	PolicyName                                  pulumi.StringPtrInput  `pulumi:"policyName"`
	ProtectableObjectContainerHostOsName        pulumi.StringPtrInput  `pulumi:"protectableObjectContainerHostOsName"`
	ProtectableObjectFriendlyName               pulumi.StringPtrInput  `pulumi:"protectableObjectFriendlyName"`
	ProtectableObjectParentLogicalContainerName pulumi.StringPtrInput  `pulumi:"protectableObjectParentLogicalContainerName"`
	ProtectableObjectProtectionState            pulumi.StringPtrInput  `pulumi:"protectableObjectProtectionState"`
	ProtectableObjectUniqueName                 pulumi.StringPtrInput  `pulumi:"protectableObjectUniqueName"`
	ProtectableObjectWorkloadType               pulumi.StringPtrInput  `pulumi:"protectableObjectWorkloadType"`
	ProtectionContainerId                       pulumi.Float64PtrInput `pulumi:"protectionContainerId"`
	ProtectionServiceStampId                    pulumi.StringPtrInput  `pulumi:"protectionServiceStampId"`
	ProtectionServiceStampUri                   pulumi.StringPtrInput  `pulumi:"protectionServiceStampUri"`
	RecoveryPointId                             pulumi.StringPtrInput  `pulumi:"recoveryPointId"`
	RecoveryPointTime                           pulumi.StringPtrInput  `pulumi:"recoveryPointTime"`
	ResourceGroupName                           pulumi.StringPtrInput  `pulumi:"resourceGroupName"`
	ResourceId                                  pulumi.StringPtrInput  `pulumi:"resourceId"`
	ResourceName                                pulumi.StringPtrInput  `pulumi:"resourceName"`
	RpIsManagedVirtualMachine                   pulumi.BoolPtrInput    `pulumi:"rpIsManagedVirtualMachine"`
	RpOriginalSAOption                          pulumi.BoolPtrInput    `pulumi:"rpOriginalSAOption"`
	RpTierInformation                           pulumi.StringMapInput  `pulumi:"rpTierInformation"`
	RpVMSizeDescription                         pulumi.StringPtrInput  `pulumi:"rpVMSizeDescription"`
	SubscriptionId                              pulumi.StringPtrInput  `pulumi:"subscriptionId"`
	TokenExtendedInformation                    pulumi.StringPtrInput  `pulumi:"tokenExtendedInformation"`
}

func (WorkloadCrrAccessTokenResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadCrrAccessTokenResponse)(nil)).Elem()
}

func (i WorkloadCrrAccessTokenResponseArgs) ToWorkloadCrrAccessTokenResponseOutput() WorkloadCrrAccessTokenResponseOutput {
	return i.ToWorkloadCrrAccessTokenResponseOutputWithContext(context.Background())
}

func (i WorkloadCrrAccessTokenResponseArgs) ToWorkloadCrrAccessTokenResponseOutputWithContext(ctx context.Context) WorkloadCrrAccessTokenResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadCrrAccessTokenResponseOutput)
}

type WorkloadCrrAccessTokenResponseOutput struct{ *pulumi.OutputState }

func (WorkloadCrrAccessTokenResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadCrrAccessTokenResponse)(nil)).Elem()
}

func (o WorkloadCrrAccessTokenResponseOutput) ToWorkloadCrrAccessTokenResponseOutput() WorkloadCrrAccessTokenResponseOutput {
	return o
}

func (o WorkloadCrrAccessTokenResponseOutput) ToWorkloadCrrAccessTokenResponseOutputWithContext(ctx context.Context) WorkloadCrrAccessTokenResponseOutput {
	return o
}

func (o WorkloadCrrAccessTokenResponseOutput) AccessTokenString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.AccessTokenString }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) BMSActiveRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.BMSActiveRegion }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ContainerId }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ContainerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ContainerType }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) CoordinatorServiceStampId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.CoordinatorServiceStampId }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) CoordinatorServiceStampUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.CoordinatorServiceStampUri }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) DatasourceContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.DatasourceContainerName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) DatasourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.DatasourceId }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) DatasourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.DatasourceName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) DatasourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.DatasourceType }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) string { return v.ObjectType }).(pulumi.StringOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectableObjectContainerHostOsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ProtectableObjectContainerHostOsName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectableObjectFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ProtectableObjectFriendlyName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectableObjectParentLogicalContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ProtectableObjectParentLogicalContainerName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectableObjectProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ProtectableObjectProtectionState }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectableObjectUniqueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ProtectableObjectUniqueName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectableObjectWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ProtectableObjectWorkloadType }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectionContainerId() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *float64 { return v.ProtectionContainerId }).(pulumi.Float64PtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectionServiceStampId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ProtectionServiceStampId }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ProtectionServiceStampUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ProtectionServiceStampUri }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) RecoveryPointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.RecoveryPointId }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) RecoveryPointTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.RecoveryPointTime }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) RpIsManagedVirtualMachine() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *bool { return v.RpIsManagedVirtualMachine }).(pulumi.BoolPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) RpOriginalSAOption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *bool { return v.RpOriginalSAOption }).(pulumi.BoolPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) RpTierInformation() pulumi.StringMapOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) map[string]string { return v.RpTierInformation }).(pulumi.StringMapOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) RpVMSizeDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.RpVMSizeDescription }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o WorkloadCrrAccessTokenResponseOutput) TokenExtendedInformation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadCrrAccessTokenResponse) *string { return v.TokenExtendedInformation }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AADPropertiesInput)(nil)).Elem(), AADPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AADPropertiesPtrInput)(nil)).Elem(), AADPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadCrrAccessTokenResponseInput)(nil)).Elem(), WorkloadCrrAccessTokenResponseArgs{})
	pulumi.RegisterOutputType(AADPropertiesOutput{})
	pulumi.RegisterOutputType(AADPropertiesPtrOutput{})
	pulumi.RegisterOutputType(WorkloadCrrAccessTokenResponseOutput{})
}
