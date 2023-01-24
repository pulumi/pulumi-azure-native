


package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPostgresInstance(ctx *pulumi.Context, args *LookupPostgresInstanceArgs, opts ...pulumi.InvokeOption) (*LookupPostgresInstanceResult, error) {
	var rv LookupPostgresInstanceResult
	err := ctx.Invoke("azure-native:azurearcdata/v20220615preview:getPostgresInstance", args, &rv, opts...)
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

func LookupPostgresInstanceOutput(ctx *pulumi.Context, args LookupPostgresInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupPostgresInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPostgresInstanceResult, error) {
			args := v.(LookupPostgresInstanceArgs)
			r, err := LookupPostgresInstance(ctx, &args, opts...)
			var s LookupPostgresInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPostgresInstanceResultOutput)
}

type LookupPostgresInstanceOutputArgs struct {
	PostgresInstanceName pulumi.StringInput `pulumi:"postgresInstanceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPostgresInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPostgresInstanceArgs)(nil)).Elem()
}


type LookupPostgresInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupPostgresInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPostgresInstanceResult)(nil)).Elem()
}

func (o LookupPostgresInstanceResultOutput) ToLookupPostgresInstanceResultOutput() LookupPostgresInstanceResultOutput {
	return o
}

func (o LookupPostgresInstanceResultOutput) ToLookupPostgresInstanceResultOutputWithContext(ctx context.Context) LookupPostgresInstanceResultOutput {
	return o
}

func (o LookupPostgresInstanceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupPostgresInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPostgresInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPostgresInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPostgresInstanceResultOutput) Properties() PostgresInstancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) PostgresInstancePropertiesResponse { return v.Properties }).(PostgresInstancePropertiesResponseOutput)
}

func (o LookupPostgresInstanceResultOutput) Sku() PostgresInstanceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) *PostgresInstanceSkuResponse { return v.Sku }).(PostgresInstanceSkuResponsePtrOutput)
}

func (o LookupPostgresInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPostgresInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPostgresInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPostgresInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPostgresInstanceResultOutput{})
}
