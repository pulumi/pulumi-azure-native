


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobCredential(ctx *pulumi.Context, args *LookupJobCredentialArgs, opts ...pulumi.InvokeOption) (*LookupJobCredentialResult, error) {
	var rv LookupJobCredentialResult
	err := ctx.Invoke("azure-native:sql:getJobCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobCredentialArgs struct {
	CredentialName    string `pulumi:"credentialName"`
	JobAgentName      string `pulumi:"jobAgentName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupJobCredentialResult struct {
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	Type     string `pulumi:"type"`
	Username string `pulumi:"username"`
}
