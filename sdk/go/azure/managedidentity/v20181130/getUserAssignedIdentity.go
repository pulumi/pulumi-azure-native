


package v20181130

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUserAssignedIdentity(ctx *pulumi.Context, args *LookupUserAssignedIdentityArgs, opts ...pulumi.InvokeOption) (*LookupUserAssignedIdentityResult, error) {
	var rv LookupUserAssignedIdentityResult
	err := ctx.Invoke("azure-native:managedidentity/v20181130:getUserAssignedIdentity", args, &rv, opts...)
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
	ClientId    string            `pulumi:"clientId"`
	Id          string            `pulumi:"id"`
	Location    string            `pulumi:"location"`
	Name        string            `pulumi:"name"`
	PrincipalId string            `pulumi:"principalId"`
	Tags        map[string]string `pulumi:"tags"`
	TenantId    string            `pulumi:"tenantId"`
	Type        string            `pulumi:"type"`
}
