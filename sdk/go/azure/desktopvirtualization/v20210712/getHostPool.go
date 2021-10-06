


package v20210712

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHostPool(ctx *pulumi.Context, args *LookupHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupHostPoolResult, error) {
	var rv LookupHostPoolResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20210712:getHostPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostPoolArgs struct {
	HostPoolName      string `pulumi:"hostPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupHostPoolResult struct {
	ApplicationGroupReferences    []string                                             `pulumi:"applicationGroupReferences"`
	CloudPcResource               bool                                                 `pulumi:"cloudPcResource"`
	CustomRdpProperty             *string                                              `pulumi:"customRdpProperty"`
	Description                   *string                                              `pulumi:"description"`
	Etag                          string                                               `pulumi:"etag"`
	FriendlyName                  *string                                              `pulumi:"friendlyName"`
	HostPoolType                  string                                               `pulumi:"hostPoolType"`
	Id                            string                                               `pulumi:"id"`
	Identity                      *ResourceModelWithAllowedPropertySetResponseIdentity `pulumi:"identity"`
	Kind                          *string                                              `pulumi:"kind"`
	LoadBalancerType              string                                               `pulumi:"loadBalancerType"`
	Location                      *string                                              `pulumi:"location"`
	ManagedBy                     *string                                              `pulumi:"managedBy"`
	MaxSessionLimit               *int                                                 `pulumi:"maxSessionLimit"`
	MigrationRequest              *MigrationRequestPropertiesResponse                  `pulumi:"migrationRequest"`
	Name                          string                                               `pulumi:"name"`
	ObjectId                      string                                               `pulumi:"objectId"`
	PersonalDesktopAssignmentType *string                                              `pulumi:"personalDesktopAssignmentType"`
	Plan                          *ResourceModelWithAllowedPropertySetResponsePlan     `pulumi:"plan"`
	PreferredAppGroupType         string                                               `pulumi:"preferredAppGroupType"`
	RegistrationInfo              *RegistrationInfoResponse                            `pulumi:"registrationInfo"`
	Ring                          *int                                                 `pulumi:"ring"`
	Sku                           *ResourceModelWithAllowedPropertySetResponseSku      `pulumi:"sku"`
	SsoClientId                   *string                                              `pulumi:"ssoClientId"`
	SsoClientSecretKeyVaultPath   *string                                              `pulumi:"ssoClientSecretKeyVaultPath"`
	SsoSecretType                 *string                                              `pulumi:"ssoSecretType"`
	SsoadfsAuthority              *string                                              `pulumi:"ssoadfsAuthority"`
	StartVMOnConnect              *bool                                                `pulumi:"startVMOnConnect"`
	Tags                          map[string]string                                    `pulumi:"tags"`
	Type                          string                                               `pulumi:"type"`
	ValidationEnvironment         *bool                                                `pulumi:"validationEnvironment"`
	VmTemplate                    *string                                              `pulumi:"vmTemplate"`
}
