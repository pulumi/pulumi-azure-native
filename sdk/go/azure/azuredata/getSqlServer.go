


package azuredata

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlServer(ctx *pulumi.Context, args *LookupSqlServerArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerResult, error) {
	var rv LookupSqlServerResult
	err := ctx.Invoke("azure-native:azuredata:getSqlServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerArgs struct {
	Expand                    *string `pulumi:"expand"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	SqlServerName             string  `pulumi:"sqlServerName"`
	SqlServerRegistrationName string  `pulumi:"sqlServerRegistrationName"`
}


type LookupSqlServerResult struct {
	Cores          *int    `pulumi:"cores"`
	Edition        *string `pulumi:"edition"`
	Id             string  `pulumi:"id"`
	Name           string  `pulumi:"name"`
	PropertyBag    *string `pulumi:"propertyBag"`
	RegistrationID *string `pulumi:"registrationID"`
	Type           string  `pulumi:"type"`
	Version        *string `pulumi:"version"`
}
