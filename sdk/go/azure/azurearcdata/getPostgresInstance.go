


package azurearcdata

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPostgresInstance(ctx *pulumi.Context, args *LookupPostgresInstanceArgs, opts ...pulumi.InvokeOption) (*LookupPostgresInstanceResult, error) {
	var rv LookupPostgresInstanceResult
	err := ctx.Invoke("azure-native:azurearcdata:getPostgresInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPostgresInstanceArgs struct {
	PostgresInstanceName string `pulumi:"postgresInstanceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupPostgresInstanceResult struct {
	ExtendedLocation *ExtendedLocationResponse          `pulumi:"extendedLocation"`
	Id               string                             `pulumi:"id"`
	Location         string                             `pulumi:"location"`
	Name             string                             `pulumi:"name"`
	Properties       PostgresInstancePropertiesResponse `pulumi:"properties"`
	Sku              *PostgresInstanceSkuResponse       `pulumi:"sku"`
	SystemData       SystemDataResponse                 `pulumi:"systemData"`
	Tags             map[string]string                  `pulumi:"tags"`
	Type             string                             `pulumi:"type"`
}


func (val *LookupPostgresInstanceResult) Defaults() *LookupPostgresInstanceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}
