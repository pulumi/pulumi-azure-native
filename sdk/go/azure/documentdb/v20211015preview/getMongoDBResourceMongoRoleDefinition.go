


package v20211015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoDBResourceMongoRoleDefinition(ctx *pulumi.Context, args *LookupMongoDBResourceMongoRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupMongoDBResourceMongoRoleDefinitionResult, error) {
	var rv LookupMongoDBResourceMongoRoleDefinitionResult
	err := ctx.Invoke("azure-native:documentdb/v20211015preview:getMongoDBResourceMongoRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoDBResourceMongoRoleDefinitionArgs struct {
	AccountName           string `pulumi:"accountName"`
	MongoRoleDefinitionId string `pulumi:"mongoRoleDefinitionId"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupMongoDBResourceMongoRoleDefinitionResult struct {
	DatabaseName *string             `pulumi:"databaseName"`
	Id           string              `pulumi:"id"`
	Name         string              `pulumi:"name"`
	Privileges   []PrivilegeResponse `pulumi:"privileges"`
	RoleName     *string             `pulumi:"roleName"`
	Roles        []RoleResponse      `pulumi:"roles"`
	Type         string              `pulumi:"type"`
}
