


package v20200921preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHostPool(ctx *pulumi.Context, args *LookupHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupHostPoolResult, error) {
	var rv LookupHostPoolResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20200921preview:getHostPool", args, &rv, opts...)
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
	ApplicationGroupReferences    []string                  `pulumi:"applicationGroupReferences"`
	CustomRdpProperty             *string                   `pulumi:"customRdpProperty"`
	Description                   *string                   `pulumi:"description"`
	FriendlyName                  *string                   `pulumi:"friendlyName"`
	HostPoolType                  string                    `pulumi:"hostPoolType"`
	Id                            string                    `pulumi:"id"`
	LoadBalancerType              string                    `pulumi:"loadBalancerType"`
	Location                      string                    `pulumi:"location"`
	MaxSessionLimit               *int                      `pulumi:"maxSessionLimit"`
	Name                          string                    `pulumi:"name"`
	PersonalDesktopAssignmentType *string                   `pulumi:"personalDesktopAssignmentType"`
	PreferredAppGroupType         string                    `pulumi:"preferredAppGroupType"`
	RegistrationInfo              *RegistrationInfoResponse `pulumi:"registrationInfo"`
	Ring                          *int                      `pulumi:"ring"`
	SsoContext                    *string                   `pulumi:"ssoContext"`
	Tags                          map[string]string         `pulumi:"tags"`
	Type                          string                    `pulumi:"type"`
	ValidationEnvironment         *bool                     `pulumi:"validationEnvironment"`
	VmTemplate                    *string                   `pulumi:"vmTemplate"`
}
