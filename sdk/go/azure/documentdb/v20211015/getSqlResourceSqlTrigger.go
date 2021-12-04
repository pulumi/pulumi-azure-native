


package v20211015

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlTrigger(ctx *pulumi.Context, args *LookupSqlResourceSqlTriggerArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlTriggerResult, error) {
	var rv LookupSqlResourceSqlTriggerResult
	err := ctx.Invoke("azure-native:documentdb/v20211015:getSqlResourceSqlTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlTriggerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ContainerName     string `pulumi:"containerName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TriggerName       string `pulumi:"triggerName"`
}


type LookupSqlResourceSqlTriggerResult struct {
	Id       string                                   `pulumi:"id"`
	Location *string                                  `pulumi:"location"`
	Name     string                                   `pulumi:"name"`
	Resource *SqlTriggerGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                        `pulumi:"tags"`
	Type     string                                   `pulumi:"type"`
}
