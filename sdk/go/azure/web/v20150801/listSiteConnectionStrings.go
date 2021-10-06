


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteConnectionStrings(ctx *pulumi.Context, args *ListSiteConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListSiteConnectionStringsResult, error) {
	var rv ListSiteConnectionStringsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteConnectionStringsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteConnectionStringsResult struct {
	Id         *string                                    `pulumi:"id"`
	Kind       *string                                    `pulumi:"kind"`
	Location   string                                     `pulumi:"location"`
	Name       *string                                    `pulumi:"name"`
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	Tags       map[string]string                          `pulumi:"tags"`
	Type       *string                                    `pulumi:"type"`
}
