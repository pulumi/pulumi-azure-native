


package v20200801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementPolicy(ctx *pulumi.Context, args *LookupManagementPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagementPolicyResult, error) {
	var rv LookupManagementPolicyResult
	err := ctx.Invoke("azure-native:storage/v20200801preview:getManagementPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementPolicyArgs struct {
	AccountName          string `pulumi:"accountName"`
	ManagementPolicyName string `pulumi:"managementPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupManagementPolicyResult struct {
	Id               string                         `pulumi:"id"`
	LastModifiedTime string                         `pulumi:"lastModifiedTime"`
	Name             string                         `pulumi:"name"`
	Policy           ManagementPolicySchemaResponse `pulumi:"policy"`
	Type             string                         `pulumi:"type"`
}
