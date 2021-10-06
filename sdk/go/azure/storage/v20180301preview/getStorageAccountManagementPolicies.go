


package v20180301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageAccountManagementPolicies(ctx *pulumi.Context, args *LookupStorageAccountManagementPoliciesArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountManagementPoliciesResult, error) {
	var rv LookupStorageAccountManagementPoliciesResult
	err := ctx.Invoke("azure-native:storage/v20180301preview:getStorageAccountManagementPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountManagementPoliciesArgs struct {
	AccountName          string `pulumi:"accountName"`
	ManagementPolicyName string `pulumi:"managementPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupStorageAccountManagementPoliciesResult struct {
	Id               string      `pulumi:"id"`
	LastModifiedTime string      `pulumi:"lastModifiedTime"`
	Name             string      `pulumi:"name"`
	Policy           interface{} `pulumi:"policy"`
	Type             string      `pulumi:"type"`
}
