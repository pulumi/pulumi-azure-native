


package documentdb

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoUserDefinition(ctx *pulumi.Context, args *LookupMongoDBResourceMongoUserDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoUserDefinitionResult, error) {
	var rv LookupMongoDBResourceMongoUserDefinitionResult
	err := ctx.Invoke("azure-native:documentdb:getMongoDBResourceMongoUserDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoUserDefinitionArgs struct {
	AccountName           string `pulumi:"accountName"`
	MongoUserDefinitionId string `pulumi:"mongoUserDefinitionId"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupMongoDBResourceMongoUserDefinitionResult struct {
	CustomData   *string        `pulumi:"customData"`
	DatabaseName *string        `pulumi:"databaseName"`
	Id           string         `pulumi:"id"`
	Mechanisms   *string        `pulumi:"mechanisms"`
	Name         string         `pulumi:"name"`
	Password     *string        `pulumi:"password"`
	Roles        []RoleResponse `pulumi:"roles"`
	Type         string         `pulumi:"type"`
	UserName     *string        `pulumi:"userName"`
}
