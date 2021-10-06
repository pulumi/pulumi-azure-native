


package v20191011preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotebookProxy(ctx *pulumi.Context, args *LookupNotebookProxyArgs, opts ...pulumi.InvokeOption) (*LookupNotebookProxyResult, error) {
	var rv LookupNotebookProxyResult
	err := ctx.Invoke("azure-native:notebooks/v20191011preview:getNotebookProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotebookProxyArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupNotebookProxyResult struct {
	Hostname            *string                             `pulumi:"hostname"`
	Id                  string                              `pulumi:"id"`
	Name                string                              `pulumi:"name"`
	PublicDns           *string                             `pulumi:"publicDns"`
	PublicNetworkAccess *string                             `pulumi:"publicNetworkAccess"`
	Region              *string                             `pulumi:"region"`
	ResourceId          string                              `pulumi:"resourceId"`
	SecondaryAppId      *string                             `pulumi:"secondaryAppId"`
	SystemData          *NotebookResourceSystemDataResponse `pulumi:"systemData"`
	Type                string                              `pulumi:"type"`
}
