


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstanceAdministrator(ctx *pulumi.Context, args *LookupManagedInstanceAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceAdministratorResult, error) {
	var rv LookupManagedInstanceAdministratorResult
	err := ctx.Invoke("azure-native:sql/v20211101preview:getManagedInstanceAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceAdministratorArgs struct {
	AdministratorName   string `pulumi:"administratorName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedInstanceAdministratorResult struct {
	AdministratorType string  `pulumi:"administratorType"`
	Id                string  `pulumi:"id"`
	Login             string  `pulumi:"login"`
	Name              string  `pulumi:"name"`
	Sid               string  `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
	Type              string  `pulumi:"type"`
}

func LookupManagedInstanceAdministratorOutput(ctx *pulumi.Context, args LookupManagedInstanceAdministratorOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceAdministratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceAdministratorResult, error) {
			args := v.(LookupManagedInstanceAdministratorArgs)
			r, err := LookupManagedInstanceAdministrator(ctx, &args, opts...)
			var s LookupManagedInstanceAdministratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceAdministratorResultOutput)
}

type LookupManagedInstanceAdministratorOutputArgs struct {
	AdministratorName   pulumi.StringInput `pulumi:"administratorName"`
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceAdministratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceAdministratorArgs)(nil)).Elem()
}


type LookupManagedInstanceAdministratorResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceAdministratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceAdministratorResult)(nil)).Elem()
}

func (o LookupManagedInstanceAdministratorResultOutput) ToLookupManagedInstanceAdministratorResultOutput() LookupManagedInstanceAdministratorResultOutput {
	return o
}

func (o LookupManagedInstanceAdministratorResultOutput) ToLookupManagedInstanceAdministratorResultOutputWithContext(ctx context.Context) LookupManagedInstanceAdministratorResultOutput {
	return o
}

func (o LookupManagedInstanceAdministratorResultOutput) AdministratorType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.AdministratorType }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceAdministratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceAdministratorResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Login }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceAdministratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceAdministratorResultOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Sid }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceAdministratorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupManagedInstanceAdministratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAdministratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceAdministratorResultOutput{})
}
