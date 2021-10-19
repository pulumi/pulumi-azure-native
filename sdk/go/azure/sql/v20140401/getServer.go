


package v20140401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServer(ctx *pulumi.Context, args *LookupServerArgs, opts ...pulumi.InvokeOption) (*LookupServerResult, error) {
	var rv LookupServerResult
	err := ctx.Invoke("azure-native:sql/v20140401:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerResult struct {
	AdministratorLogin         *string           `pulumi:"administratorLogin"`
	ExternalAdministratorLogin string            `pulumi:"externalAdministratorLogin"`
	ExternalAdministratorSid   string            `pulumi:"externalAdministratorSid"`
	FullyQualifiedDomainName   string            `pulumi:"fullyQualifiedDomainName"`
	Id                         string            `pulumi:"id"`
	Kind                       string            `pulumi:"kind"`
	Location                   string            `pulumi:"location"`
	Name                       string            `pulumi:"name"`
	State                      string            `pulumi:"state"`
	Tags                       map[string]string `pulumi:"tags"`
	Type                       string            `pulumi:"type"`
	Version                    *string           `pulumi:"version"`
}
