


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:operationalinsights/v20200801:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	LinkedServiceName string `pulumi:"linkedServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedServiceResult struct {
	Id                    string            `pulumi:"id"`
	Name                  string            `pulumi:"name"`
	ProvisioningState     *string           `pulumi:"provisioningState"`
	ResourceId            *string           `pulumi:"resourceId"`
	Tags                  map[string]string `pulumi:"tags"`
	Type                  string            `pulumi:"type"`
	WriteAccessResourceId *string           `pulumi:"writeAccessResourceId"`
}
