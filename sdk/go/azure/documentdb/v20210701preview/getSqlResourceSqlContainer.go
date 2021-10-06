


package v20210701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlContainer(ctx *pulumi.Context, args *LookupSqlResourceSqlContainerArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlContainerResult, error) {
	var rv LookupSqlResourceSqlContainerResult
	err := ctx.Invoke("azure-native:documentdb/v20210701preview:getSqlResourceSqlContainer", args, &rv, opts...)
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
	Identity *ManagedServiceIdentityResponse            `pulumi:"identity"`
	Location *string                                    `pulumi:"location"`
	Name     string                                     `pulumi:"name"`
	Options  *SqlContainerGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *SqlContainerGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                          `pulumi:"tags"`
	Type     string                                     `pulumi:"type"`
}
