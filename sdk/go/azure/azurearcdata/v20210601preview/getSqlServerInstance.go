


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlServerInstance(ctx *pulumi.Context, args *LookupSqlServerInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerInstanceResult, error) {
	var rv LookupSqlServerInstanceResult
	err := ctx.Invoke("azure-native:azurearcdata/v20210601preview:getSqlServerInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerInstanceArgs struct {
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SqlServerInstanceName string `pulumi:"sqlServerInstanceName"`
}


type LookupSqlServerInstanceResult struct {
	Id         string                              `pulumi:"id"`
	Location   string                              `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties SqlServerInstancePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                  `pulumi:"systemData"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
}
