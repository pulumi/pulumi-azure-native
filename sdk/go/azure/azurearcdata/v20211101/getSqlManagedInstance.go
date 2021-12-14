


package v20211101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlManagedInstance(ctx *pulumi.Context, args *LookupSqlManagedInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSqlManagedInstanceResult, error) {
	var rv LookupSqlManagedInstanceResult
	err := ctx.Invoke("azure-native:azurearcdata/v20211101:getSqlManagedInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSqlManagedInstanceArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SqlManagedInstanceName string `pulumi:"sqlManagedInstanceName"`
}


type LookupSqlManagedInstanceResult struct {
	ExtendedLocation *ExtendedLocationResponse            `pulumi:"extendedLocation"`
	Id               string                               `pulumi:"id"`
	Location         string                               `pulumi:"location"`
	Name             string                               `pulumi:"name"`
	Properties       SqlManagedInstancePropertiesResponse `pulumi:"properties"`
	Sku              *SqlManagedInstanceSkuResponse       `pulumi:"sku"`
	SystemData       SystemDataResponse                   `pulumi:"systemData"`
	Tags             map[string]string                    `pulumi:"tags"`
	Type             string                               `pulumi:"type"`
}


func (val *LookupSqlManagedInstanceResult) Defaults() *LookupSqlManagedInstanceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}
