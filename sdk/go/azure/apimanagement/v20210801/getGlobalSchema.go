


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGlobalSchema(ctx *pulumi.Context, args *LookupGlobalSchemaArgs, opts ...pulumi.InvokeOption) (*LookupGlobalSchemaResult, error) {
	var rv LookupGlobalSchemaResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:getGlobalSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalSchemaArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SchemaId          string `pulumi:"schemaId"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGlobalSchemaResult struct {
	Description *string     `pulumi:"description"`
	Id          string      `pulumi:"id"`
	Name        string      `pulumi:"name"`
	SchemaType  string      `pulumi:"schemaType"`
	Type        string      `pulumi:"type"`
	Value       interface{} `pulumi:"value"`
}
