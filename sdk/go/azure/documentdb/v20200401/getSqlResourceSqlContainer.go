


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlContainer(ctx *pulumi.Context, args *LookupSqlResourceSqlContainerArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlContainerResult, error) {
	var rv LookupSqlResourceSqlContainerResult
	err := ctx.Invoke("azure-native:documentdb/v20200401:getSqlResourceSqlContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlContainerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ContainerName     string `pulumi:"containerName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSqlResourceSqlContainerResult struct {
	Id       string                                     `pulumi:"id"`
	Location *string                                    `pulumi:"location"`
	Name     string                                     `pulumi:"name"`
	Options  *SqlContainerGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *SqlContainerGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                          `pulumi:"tags"`
	Type     string                                     `pulumi:"type"`
}
