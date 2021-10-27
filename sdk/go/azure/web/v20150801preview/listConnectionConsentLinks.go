


package v20150801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectionConsentLinks(ctx *pulumi.Context, args *ListConnectionConsentLinksArgs, opts ...pulumi.InvokeOption) (*ListConnectionConsentLinksResult, error) {
	var rv ListConnectionConsentLinksResult
	err := ctx.Invoke("azure-native:web/v20150801preview:listConnectionConsentLinks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionConsentLinksArgs struct {
	ConnectionName    string                      `pulumi:"connectionName"`
	Id                *string                     `pulumi:"id"`
	Kind              *string                     `pulumi:"kind"`
	Location          *string                     `pulumi:"location"`
	Name              *string                     `pulumi:"name"`
	Parameters        []ConsentLinkInputParameter `pulumi:"parameters"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	Tags              map[string]string           `pulumi:"tags"`
	Type              *string                     `pulumi:"type"`
}


type ListConnectionConsentLinksResult struct {
	Value []ConsentLinkResponse `pulumi:"value"`
}
