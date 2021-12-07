


package v20150831preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUserAssignedIdentity(ctx *pulumi.Context, args *LookupUserAssignedIdentityArgs, opts ...pulumi.InvokeOption) (*LookupUserAssignedIdentityResult, error) {
	var rv LookupUserAssignedIdentityResult
	err := ctx.Invoke("azure-native:managedidentity/v20150831preview:getUserAssignedIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserAssignedIdentityArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupUserAssignedIdentityResult struct {
	ClientId        string            `pulumi:"clientId"`
	ClientSecretUrl string            `pulumi:"clientSecretUrl"`
	Id              string            `pulumi:"id"`
	Location        string            `pulumi:"location"`
	Name            string            `pulumi:"name"`
	PrincipalId     string            `pulumi:"principalId"`
	Tags            map[string]string `pulumi:"tags"`
	TenantId        string            `pulumi:"tenantId"`
	Type            string            `pulumi:"type"`
}
