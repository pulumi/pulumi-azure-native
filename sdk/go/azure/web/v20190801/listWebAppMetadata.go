


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppMetadata(ctx *pulumi.Context, args *ListWebAppMetadataArgs, opts ...pulumi.InvokeOption) (*ListWebAppMetadataResult, error) {
	var rv ListWebAppMetadataResult
	err := ctx.Invoke("azure-native:web/v20190801:listWebAppMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppMetadataArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppMetadataResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}
