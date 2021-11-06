


package v20140901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:notificationhubs/v20140901:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceResult struct {
	Id         *string                     `pulumi:"id"`
	Location   *string                     `pulumi:"location"`
	Name       *string                     `pulumi:"name"`
	Properties NamespacePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       *string                     `pulumi:"type"`
}
