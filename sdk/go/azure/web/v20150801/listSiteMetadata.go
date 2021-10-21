


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteMetadata(ctx *pulumi.Context, args *ListSiteMetadataArgs, opts ...pulumi.InvokeOption) (*ListSiteMetadataResult, error) {
	var rv ListSiteMetadataResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteMetadataArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteMetadataResult struct {
	Id         *string           `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   string            `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}
