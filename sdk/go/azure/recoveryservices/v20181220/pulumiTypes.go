


package v20181220

type AADProperties struct {
	Audience                 *string `pulumi:"audience"`
	Authority                *string `pulumi:"authority"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	ServicePrincipalObjectId *string `pulumi:"servicePrincipalObjectId"`
	TenantId                 *string `pulumi:"tenantId"`
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

func init() {
}
