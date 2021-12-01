


package v20211101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchemaRegistry(ctx *pulumi.Context, args *LookupSchemaRegistryArgs, opts ...pulumi.InvokeOption) (*LookupSchemaRegistryResult, error) {
	var rv LookupSchemaRegistryResult
	err := ctx.Invoke("azure-native:eventhub/v20211101:getSchemaRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaRegistryArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SchemaGroupName   string `pulumi:"schemaGroupName"`
}


type LookupSchemaRegistryResult struct {
	CreatedAtUtc        string             `pulumi:"createdAtUtc"`
	ETag                string             `pulumi:"eTag"`
	GroupProperties     map[string]string  `pulumi:"groupProperties"`
	Id                  string             `pulumi:"id"`
	Location            string             `pulumi:"location"`
	Name                string             `pulumi:"name"`
	SchemaCompatibility *string            `pulumi:"schemaCompatibility"`
	SchemaType          *string            `pulumi:"schemaType"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Type                string             `pulumi:"type"`
	UpdatedAtUtc        string             `pulumi:"updatedAtUtc"`
}
