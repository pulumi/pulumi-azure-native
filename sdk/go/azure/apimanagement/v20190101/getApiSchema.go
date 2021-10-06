


package v20190101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiSchema(ctx *pulumi.Context, args *LookupApiSchemaArgs, opts ...pulumi.InvokeOption) (*LookupApiSchemaResult, error) {
	var rv LookupApiSchemaResult
	err := ctx.Invoke("azure-native:apimanagement/v20190101:getApiSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiSchemaArgs struct {
	ApiId             string `pulumi:"apiId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SchemaId          string `pulumi:"schemaId"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiSchemaResult struct {
	ContentType string      `pulumi:"contentType"`
	Document    interface{} `pulumi:"document"`
	Id          string      `pulumi:"id"`
	Name        string      `pulumi:"name"`
	Type        string      `pulumi:"type"`
}
