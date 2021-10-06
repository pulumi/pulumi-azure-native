


package automation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCredential(ctx *pulumi.Context, args *LookupCredentialArgs, opts ...pulumi.InvokeOption) (*LookupCredentialResult, error) {
	var rv LookupCredentialResult
	err := ctx.Invoke("azure-native:automation:getCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCredentialArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	CredentialName        string `pulumi:"credentialName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupCredentialResult struct {
	CreationTime     string  `pulumi:"creationTime"`
	Description      *string `pulumi:"description"`
	Id               string  `pulumi:"id"`
	LastModifiedTime string  `pulumi:"lastModifiedTime"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	UserName         string  `pulumi:"userName"`
}
