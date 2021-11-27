


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchema(ctx *pulumi.Context, args *LookupSchemaArgs, opts ...pulumi.InvokeOption) (*LookupSchemaResult, error) {
	var rv LookupSchemaResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:getSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SchemaId          string `pulumi:"schemaId"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupSchemaResult struct {
	Description *string `pulumi:"description"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	SchemaType  string  `pulumi:"schemaType"`
	Type        string  `pulumi:"type"`
	Value       *string `pulumi:"value"`
}
