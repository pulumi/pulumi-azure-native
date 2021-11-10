


package solutions

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationDefinition(ctx *pulumi.Context, args *LookupApplicationDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationDefinitionResult, error) {
	var rv LookupApplicationDefinitionResult
	err := ctx.Invoke("azure-native:solutions:getApplicationDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationDefinitionArgs struct {
	ApplicationDefinitionName string `pulumi:"applicationDefinitionName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupApplicationDefinitionResult struct {
	Artifacts          []ApplicationDefinitionArtifactResponse            `pulumi:"artifacts"`
	Authorizations     []ApplicationAuthorizationResponse                 `pulumi:"authorizations"`
	CreateUiDefinition interface{}                                        `pulumi:"createUiDefinition"`
	DeploymentPolicy   *ApplicationDeploymentPolicyResponse               `pulumi:"deploymentPolicy"`
	Description        *string                                            `pulumi:"description"`
	DisplayName        *string                                            `pulumi:"displayName"`
	Id                 string                                             `pulumi:"id"`
	IsEnabled          *bool                                              `pulumi:"isEnabled"`
	Location           *string                                            `pulumi:"location"`
	LockLevel          string                                             `pulumi:"lockLevel"`
	LockingPolicy      *ApplicationPackageLockingPolicyDefinitionResponse `pulumi:"lockingPolicy"`
	MainTemplate       interface{}                                        `pulumi:"mainTemplate"`
	ManagedBy          *string                                            `pulumi:"managedBy"`
	ManagementPolicy   *ApplicationManagementPolicyResponse               `pulumi:"managementPolicy"`
	Name               string                                             `pulumi:"name"`
	NotificationPolicy *ApplicationNotificationPolicyResponse             `pulumi:"notificationPolicy"`
	PackageFileUri     *string                                            `pulumi:"packageFileUri"`
	Policies           []ApplicationPolicyResponse                        `pulumi:"policies"`
	ProvisioningState  string                                             `pulumi:"provisioningState"`
	Sku                *SkuResponse                                       `pulumi:"sku"`
	StorageAccountId   *string                                            `pulumi:"storageAccountId"`
	SystemData         SystemDataResponse                                 `pulumi:"systemData"`
	Tags               map[string]string                                  `pulumi:"tags"`
	Type               string                                             `pulumi:"type"`
}
