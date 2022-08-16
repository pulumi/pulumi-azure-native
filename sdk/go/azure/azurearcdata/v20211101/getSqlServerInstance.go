


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlServerInstance(ctx *pulumi.Context, args *LookupSqlServerInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerInstanceResult, error) {
	var rv LookupSqlServerInstanceResult
	err := ctx.Invoke("azure-native:azurearcdata/v20211101:getSqlServerInstance", args, &rv, opts...)
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

func LookupSqlServerInstanceOutput(ctx *pulumi.Context, args LookupSqlServerInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupSqlServerInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlServerInstanceResult, error) {
			args := v.(LookupSqlServerInstanceArgs)
			r, err := LookupSqlServerInstance(ctx, &args, opts...)
			var s LookupSqlServerInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlServerInstanceResultOutput)
}

type LookupSqlServerInstanceOutputArgs struct {
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlServerInstanceName pulumi.StringInput `pulumi:"sqlServerInstanceName"`
}

func (LookupSqlServerInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerInstanceArgs)(nil)).Elem()
}


type LookupSqlServerInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupSqlServerInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerInstanceResult)(nil)).Elem()
}

func (o LookupSqlServerInstanceResultOutput) ToLookupSqlServerInstanceResultOutput() LookupSqlServerInstanceResultOutput {
	return o
}

func (o LookupSqlServerInstanceResultOutput) ToLookupSqlServerInstanceResultOutputWithContext(ctx context.Context) LookupSqlServerInstanceResultOutput {
	return o
}

func (o LookupSqlServerInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlServerInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlServerInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlServerInstanceResultOutput) Properties() SqlServerInstancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSqlServerInstanceResult) SqlServerInstancePropertiesResponse { return v.Properties }).(SqlServerInstancePropertiesResponseOutput)
}

func (o LookupSqlServerInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlServerInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlServerInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlServerInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlServerInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlServerInstanceResultOutput{})
}
