


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlManagedInstance(ctx *pulumi.Context, args *LookupSqlManagedInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSqlManagedInstanceResult, error) {
	var rv LookupSqlManagedInstanceResult
	err := ctx.Invoke("azure-native:azurearcdata/v20210701preview:getSqlManagedInstance", args, &rv, opts...)
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

func LookupSqlManagedInstanceOutput(ctx *pulumi.Context, args LookupSqlManagedInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupSqlManagedInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlManagedInstanceResult, error) {
			args := v.(LookupSqlManagedInstanceArgs)
			r, err := LookupSqlManagedInstance(ctx, &args, opts...)
			var s LookupSqlManagedInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlManagedInstanceResultOutput)
}

type LookupSqlManagedInstanceOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlManagedInstanceName pulumi.StringInput `pulumi:"sqlManagedInstanceName"`
}

func (LookupSqlManagedInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlManagedInstanceArgs)(nil)).Elem()
}


type LookupSqlManagedInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupSqlManagedInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlManagedInstanceResult)(nil)).Elem()
}

func (o LookupSqlManagedInstanceResultOutput) ToLookupSqlManagedInstanceResultOutput() LookupSqlManagedInstanceResultOutput {
	return o
}

func (o LookupSqlManagedInstanceResultOutput) ToLookupSqlManagedInstanceResultOutputWithContext(ctx context.Context) LookupSqlManagedInstanceResultOutput {
	return o
}

func (o LookupSqlManagedInstanceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupSqlManagedInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlManagedInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlManagedInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlManagedInstanceResultOutput) Properties() SqlManagedInstancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) SqlManagedInstancePropertiesResponse { return v.Properties }).(SqlManagedInstancePropertiesResponseOutput)
}

func (o LookupSqlManagedInstanceResultOutput) Sku() SqlManagedInstanceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) *SqlManagedInstanceSkuResponse { return v.Sku }).(SqlManagedInstanceSkuResponsePtrOutput)
}

func (o LookupSqlManagedInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlManagedInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlManagedInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlManagedInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlManagedInstanceResultOutput{})
}
