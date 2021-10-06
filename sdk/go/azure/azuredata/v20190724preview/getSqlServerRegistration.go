


package v20190724preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlServerRegistration(ctx *pulumi.Context, args *LookupSqlServerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerRegistrationResult, error) {
	var rv LookupSqlServerRegistrationResult
	err := ctx.Invoke("azure-native:azuredata/v20190724preview:getSqlServerRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerRegistrationArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	SqlServerRegistrationName string `pulumi:"sqlServerRegistrationName"`
}


type LookupSqlServerRegistrationResult struct {
	Id             string             `pulumi:"id"`
	Location       string             `pulumi:"location"`
	Name           string             `pulumi:"name"`
	PropertyBag    *string            `pulumi:"propertyBag"`
	ResourceGroup  *string            `pulumi:"resourceGroup"`
	SubscriptionId *string            `pulumi:"subscriptionId"`
	SystemData     SystemDataResponse `pulumi:"systemData"`
	Tags           map[string]string  `pulumi:"tags"`
	Type           string             `pulumi:"type"`
}
