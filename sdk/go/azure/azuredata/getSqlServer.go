


package azuredata

import (
	"context"
	"reflect"

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

func LookupSqlServerOutput(ctx *pulumi.Context, args LookupSqlServerOutputArgs, opts ...pulumi.InvokeOption) LookupSqlServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlServerResult, error) {
			args := v.(LookupSqlServerArgs)
			r, err := LookupSqlServer(ctx, &args, opts...)
			var s LookupSqlServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlServerResultOutput)
}

type LookupSqlServerOutputArgs struct {
	Expand                    pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName         pulumi.StringInput    `pulumi:"resourceGroupName"`
	SqlServerName             pulumi.StringInput    `pulumi:"sqlServerName"`
	SqlServerRegistrationName pulumi.StringInput    `pulumi:"sqlServerRegistrationName"`
}

func (LookupSqlServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerArgs)(nil)).Elem()
}


type LookupSqlServerResultOutput struct{ *pulumi.OutputState }

func (LookupSqlServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerResult)(nil)).Elem()
}

func (o LookupSqlServerResultOutput) ToLookupSqlServerResultOutput() LookupSqlServerResultOutput {
	return o
}

func (o LookupSqlServerResultOutput) ToLookupSqlServerResultOutputWithContext(ctx context.Context) LookupSqlServerResultOutput {
	return o
}

func (o LookupSqlServerResultOutput) Cores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *int { return v.Cores }).(pulumi.IntPtrOutput)
}

func (o LookupSqlServerResultOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o LookupSqlServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlServerResultOutput) PropertyBag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *string { return v.PropertyBag }).(pulumi.StringPtrOutput)
}

func (o LookupSqlServerResultOutput) RegistrationID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *string { return v.RegistrationID }).(pulumi.StringPtrOutput)
}

func (o LookupSqlServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSqlServerResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlServerResultOutput{})
}
