


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteIdentifiersAssignedToHostName(ctx *pulumi.Context, args *ListSiteIdentifiersAssignedToHostNameArgs, opts ...pulumi.InvokeOption) (*ListSiteIdentifiersAssignedToHostNameResult, error) {
	var rv ListSiteIdentifiersAssignedToHostNameResult
	err := ctx.Invoke("azure-native:web/v20200901:listSiteIdentifiersAssignedToHostName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteIdentifiersAssignedToHostNameArgs struct {
	Name *string `pulumi:"name"`
}


type ListSiteIdentifiersAssignedToHostNameResult struct {
	NextLink string               `pulumi:"nextLink"`
	Value    []IdentifierResponse `pulumi:"value"`
}
