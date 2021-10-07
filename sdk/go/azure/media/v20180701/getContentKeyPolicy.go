


package v20180701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContentKeyPolicy(ctx *pulumi.Context, args *LookupContentKeyPolicyArgs, opts ...pulumi.InvokeOption) (*LookupContentKeyPolicyResult, error) {
	var rv LookupContentKeyPolicyResult
	err := ctx.Invoke("azure-native:media/v20180701:getContentKeyPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentKeyPolicyArgs struct {
	AccountName          string `pulumi:"accountName"`
	ContentKeyPolicyName string `pulumi:"contentKeyPolicyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupContentKeyPolicyResult struct {
	Created      string                           `pulumi:"created"`
	Description  *string                          `pulumi:"description"`
	Id           string                           `pulumi:"id"`
	LastModified string                           `pulumi:"lastModified"`
	Name         string                           `pulumi:"name"`
	Options      []ContentKeyPolicyOptionResponse `pulumi:"options"`
	PolicyId     string                           `pulumi:"policyId"`
	Type         string                           `pulumi:"type"`
}
