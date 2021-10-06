


package v20150801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountMap(ctx *pulumi.Context, args *LookupIntegrationAccountMapArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountMapResult, error) {
	var rv LookupIntegrationAccountMapResult
	err := ctx.Invoke("azure-native:logic/v20150801preview:getIntegrationAccountMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountMapArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	MapName                string `pulumi:"mapName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}

type LookupIntegrationAccountMapResult struct {
	ChangedTime string                                `pulumi:"changedTime"`
	Content     interface{}                           `pulumi:"content"`
	ContentLink IntegrationAccountContentLinkResponse `pulumi:"contentLink"`
	ContentType *string                               `pulumi:"contentType"`
	CreatedTime string                                `pulumi:"createdTime"`
	Id          *string                               `pulumi:"id"`
	Location    *string                               `pulumi:"location"`
	MapType     *string                               `pulumi:"mapType"`
	Metadata    interface{}                           `pulumi:"metadata"`
	Name        *string                               `pulumi:"name"`
	Tags        map[string]string                     `pulumi:"tags"`
	Type        *string                               `pulumi:"type"`
}
