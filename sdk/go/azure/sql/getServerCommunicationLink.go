


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerCommunicationLink(ctx *pulumi.Context, args *LookupServerCommunicationLinkArgs, opts ...pulumi.InvokeOption) (*LookupServerCommunicationLinkResult, error) {
	var rv LookupServerCommunicationLinkResult
	err := ctx.Invoke("azure-native:sql:getServerCommunicationLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerCommunicationLinkArgs struct {
	CommunicationLinkName string `pulumi:"communicationLinkName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ServerName            string `pulumi:"serverName"`
}


type LookupServerCommunicationLinkResult struct {
	Id            string `pulumi:"id"`
	Kind          string `pulumi:"kind"`
	Location      string `pulumi:"location"`
	Name          string `pulumi:"name"`
	PartnerServer string `pulumi:"partnerServer"`
	State         string `pulumi:"state"`
	Type          string `pulumi:"type"`
}
