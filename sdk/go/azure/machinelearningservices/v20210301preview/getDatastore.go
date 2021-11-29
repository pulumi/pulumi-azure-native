


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatastoreArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDatastoreResult struct {
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties DatastorePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Type       string                      `pulumi:"type"`
}
