


package datafactory

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:datafactory:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	FactoryName       string `pulumi:"factoryName"`
	LinkedServiceName string `pulumi:"linkedServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLinkedServiceResult struct {
	Etag       string      `pulumi:"etag"`
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
