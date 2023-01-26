


package v20171201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerAdministrator(ctx *pulumi.Context, args *LookupServerAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupServerAdministratorResult, error) {
	var rv LookupServerAdministratorResult
	err := ctx.Invoke("azure-native:dbformysql/v20171201preview:getServerAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerAdministratorArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerAdministratorResult struct {
	AdministratorType string `pulumi:"administratorType"`
	Id                string `pulumi:"id"`
	Login             string `pulumi:"login"`
	Name              string `pulumi:"name"`
	Sid               string `pulumi:"sid"`
	TenantId          string `pulumi:"tenantId"`
	Type              string `pulumi:"type"`
}

func LookupServerAdministratorOutput(ctx *pulumi.Context, args LookupServerAdministratorOutputArgs, opts ...pulumi.InvokeOption) LookupServerAdministratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerAdministratorResult, error) {
			args := v.(LookupServerAdministratorArgs)
			r, err := LookupServerAdministrator(ctx, &args, opts...)
			var s LookupServerAdministratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerAdministratorResultOutput)
}

type LookupServerAdministratorOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerAdministratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAdministratorArgs)(nil)).Elem()
}


type LookupServerAdministratorResultOutput struct{ *pulumi.OutputState }

func (LookupServerAdministratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAdministratorResult)(nil)).Elem()
}

func (o LookupServerAdministratorResultOutput) ToLookupServerAdministratorResultOutput() LookupServerAdministratorResultOutput {
	return o
}

func (o LookupServerAdministratorResultOutput) ToLookupServerAdministratorResultOutputWithContext(ctx context.Context) LookupServerAdministratorResultOutput {
	return o
}

func (o LookupServerAdministratorResultOutput) AdministratorType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.AdministratorType }).(pulumi.StringOutput)
}

func (o LookupServerAdministratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerAdministratorResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Login }).(pulumi.StringOutput)
}

func (o LookupServerAdministratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerAdministratorResultOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Sid }).(pulumi.StringOutput)
}

func (o LookupServerAdministratorResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupServerAdministratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAdministratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerAdministratorResultOutput{})
}
