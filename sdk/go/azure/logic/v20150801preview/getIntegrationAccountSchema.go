


package v20150801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountSchema(ctx *pulumi.Context, args *LookupIntegrationAccountSchemaArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountSchemaResult, error) {
	var rv LookupIntegrationAccountSchemaResult
	err := ctx.Invoke("azure-native:logic/v20150801preview:getIntegrationAccountSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountSchemaArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SchemaName             string `pulumi:"schemaName"`
}

type LookupIntegrationAccountSchemaResult struct {
	ChangedTime     string                                `pulumi:"changedTime"`
	Content         interface{}                           `pulumi:"content"`
	ContentLink     IntegrationAccountContentLinkResponse `pulumi:"contentLink"`
	ContentType     *string                               `pulumi:"contentType"`
	CreatedTime     string                                `pulumi:"createdTime"`
	Id              *string                               `pulumi:"id"`
	Location        *string                               `pulumi:"location"`
	Metadata        interface{}                           `pulumi:"metadata"`
	Name            *string                               `pulumi:"name"`
	SchemaType      *string                               `pulumi:"schemaType"`
	Tags            map[string]string                     `pulumi:"tags"`
	TargetNamespace *string                               `pulumi:"targetNamespace"`
	Type            *string                               `pulumi:"type"`
}
