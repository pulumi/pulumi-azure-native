


package v20181101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppConnectionStrings(ctx *pulumi.Context, args *ListWebAppConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListWebAppConnectionStringsResult, error) {
	var rv ListWebAppConnectionStringsResult
	err := ctx.Invoke("azure-native:web/v20181101:listWebAppConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppConnectionStringsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppConnectionStringsResult struct {
	Id         string                                     `pulumi:"id"`
	Kind       *string                                    `pulumi:"kind"`
	Name       string                                     `pulumi:"name"`
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}
