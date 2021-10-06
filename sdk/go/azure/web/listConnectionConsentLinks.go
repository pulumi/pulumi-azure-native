


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectionConsentLinks(ctx *pulumi.Context, args *ListConnectionConsentLinksArgs, opts ...pulumi.InvokeOption) (*ListConnectionConsentLinksResult, error) {
	var rv ListConnectionConsentLinksResult
	err := ctx.Invoke("azure-native:web:listConnectionConsentLinks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionConsentLinksArgs struct {
	ConnectionName    string                           `pulumi:"connectionName"`
	Parameters        []ConsentLinkParameterDefinition `pulumi:"parameters"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
	SubscriptionId    *string                          `pulumi:"subscriptionId"`
}


type ListConnectionConsentLinksResult struct {
	Value []ConsentLinkDefinitionResponse `pulumi:"value"`
}
