


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAzureADAdministrator(ctx *pulumi.Context, args *LookupAzureADAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupAzureADAdministratorResult, error) {
	var rv LookupAzureADAdministratorResult
	err := ctx.Invoke("azure-native:dbformysql/v20211201preview:getAzureADAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureADAdministratorArgs struct {
	AdministratorName string `pulumi:"administratorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupAzureADAdministratorResult struct {
	AdministratorType  *string            `pulumi:"administratorType"`
	Id                 string             `pulumi:"id"`
	IdentityResourceId *string            `pulumi:"identityResourceId"`
	Login              *string            `pulumi:"login"`
	Name               string             `pulumi:"name"`
	Sid                *string            `pulumi:"sid"`
	SystemData         SystemDataResponse `pulumi:"systemData"`
	TenantId           *string            `pulumi:"tenantId"`
	Type               string             `pulumi:"type"`
}

func LookupAzureADAdministratorOutput(ctx *pulumi.Context, args LookupAzureADAdministratorOutputArgs, opts ...pulumi.InvokeOption) LookupAzureADAdministratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureADAdministratorResult, error) {
			args := v.(LookupAzureADAdministratorArgs)
			r, err := LookupAzureADAdministrator(ctx, &args, opts...)
			var s LookupAzureADAdministratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureADAdministratorResultOutput)
}

type LookupAzureADAdministratorOutputArgs struct {
	AdministratorName pulumi.StringInput `pulumi:"administratorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupAzureADAdministratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureADAdministratorArgs)(nil)).Elem()
}


type LookupAzureADAdministratorResultOutput struct{ *pulumi.OutputState }

func (LookupAzureADAdministratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureADAdministratorResult)(nil)).Elem()
}

func (o LookupAzureADAdministratorResultOutput) ToLookupAzureADAdministratorResultOutput() LookupAzureADAdministratorResultOutput {
	return o
}

func (o LookupAzureADAdministratorResultOutput) ToLookupAzureADAdministratorResultOutputWithContext(ctx context.Context) LookupAzureADAdministratorResultOutput {
	return o
}

func (o LookupAzureADAdministratorResultOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) *string { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o LookupAzureADAdministratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAzureADAdministratorResultOutput) IdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) *string { return v.IdentityResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupAzureADAdministratorResultOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) *string { return v.Login }).(pulumi.StringPtrOutput)
}

func (o LookupAzureADAdministratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAzureADAdministratorResultOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o LookupAzureADAdministratorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAzureADAdministratorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupAzureADAdministratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureADAdministratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureADAdministratorResultOutput{})
}
