


package azuredata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlServerRegistration(ctx *pulumi.Context, args *LookupSqlServerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerRegistrationResult, error) {
	var rv LookupSqlServerRegistrationResult
	err := ctx.Invoke("azure-native:azuredata:getSqlServerRegistration", args, &rv, opts...)
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

func LookupSqlServerRegistrationOutput(ctx *pulumi.Context, args LookupSqlServerRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupSqlServerRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlServerRegistrationResult, error) {
			args := v.(LookupSqlServerRegistrationArgs)
			r, err := LookupSqlServerRegistration(ctx, &args, opts...)
			var s LookupSqlServerRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlServerRegistrationResultOutput)
}

type LookupSqlServerRegistrationOutputArgs struct {
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	SqlServerRegistrationName pulumi.StringInput `pulumi:"sqlServerRegistrationName"`
}

func (LookupSqlServerRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerRegistrationArgs)(nil)).Elem()
}


type LookupSqlServerRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupSqlServerRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerRegistrationResult)(nil)).Elem()
}

func (o LookupSqlServerRegistrationResultOutput) ToLookupSqlServerRegistrationResultOutput() LookupSqlServerRegistrationResultOutput {
	return o
}

func (o LookupSqlServerRegistrationResultOutput) ToLookupSqlServerRegistrationResultOutputWithContext(ctx context.Context) LookupSqlServerRegistrationResultOutput {
	return o
}

func (o LookupSqlServerRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlServerRegistrationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSqlServerRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlServerRegistrationResultOutput) PropertyBag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.PropertyBag }).(pulumi.StringPtrOutput)
}

func (o LookupSqlServerRegistrationResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupSqlServerRegistrationResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LookupSqlServerRegistrationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSqlServerRegistrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSqlServerRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlServerRegistrationResultOutput{})
}
