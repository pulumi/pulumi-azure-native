


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerAzureADAdministrator(ctx *pulumi.Context, args *LookupServerAzureADAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupServerAzureADAdministratorResult, error) {
	var rv LookupServerAzureADAdministratorResult
	err := ctx.Invoke("azure-native:sql/v20200801preview:getServerAzureADAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerAzureADAdministratorArgs struct {
	AdministratorName string `pulumi:"administratorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupServerAzureADAdministratorResult struct {
	AdministratorType         string  `pulumi:"administratorType"`
	AzureADOnlyAuthentication bool    `pulumi:"azureADOnlyAuthentication"`
	Id                        string  `pulumi:"id"`
	Login                     string  `pulumi:"login"`
	Name                      string  `pulumi:"name"`
	Sid                       string  `pulumi:"sid"`
	TenantId                  *string `pulumi:"tenantId"`
	Type                      string  `pulumi:"type"`
}

func LookupServerAzureADAdministratorOutput(ctx *pulumi.Context, args LookupServerAzureADAdministratorOutputArgs, opts ...pulumi.InvokeOption) LookupServerAzureADAdministratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerAzureADAdministratorResult, error) {
			args := v.(LookupServerAzureADAdministratorArgs)
			r, err := LookupServerAzureADAdministrator(ctx, &args, opts...)
			var s LookupServerAzureADAdministratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerAzureADAdministratorResultOutput)
}

type LookupServerAzureADAdministratorOutputArgs struct {
	AdministratorName pulumi.StringInput `pulumi:"administratorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerAzureADAdministratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAzureADAdministratorArgs)(nil)).Elem()
}


type LookupServerAzureADAdministratorResultOutput struct{ *pulumi.OutputState }

func (LookupServerAzureADAdministratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAzureADAdministratorResult)(nil)).Elem()
}

func (o LookupServerAzureADAdministratorResultOutput) ToLookupServerAzureADAdministratorResultOutput() LookupServerAzureADAdministratorResultOutput {
	return o
}

func (o LookupServerAzureADAdministratorResultOutput) ToLookupServerAzureADAdministratorResultOutputWithContext(ctx context.Context) LookupServerAzureADAdministratorResultOutput {
	return o
}

func (o LookupServerAzureADAdministratorResultOutput) AdministratorType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADAdministratorResult) string { return v.AdministratorType }).(pulumi.StringOutput)
}

func (o LookupServerAzureADAdministratorResultOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerAzureADAdministratorResult) bool { return v.AzureADOnlyAuthentication }).(pulumi.BoolOutput)
}

func (o LookupServerAzureADAdministratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADAdministratorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerAzureADAdministratorResultOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADAdministratorResult) string { return v.Login }).(pulumi.StringOutput)
}

func (o LookupServerAzureADAdministratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADAdministratorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerAzureADAdministratorResultOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADAdministratorResult) string { return v.Sid }).(pulumi.StringOutput)
}

func (o LookupServerAzureADAdministratorResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerAzureADAdministratorResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupServerAzureADAdministratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADAdministratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerAzureADAdministratorResultOutput{})
}
