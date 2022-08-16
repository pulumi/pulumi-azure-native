


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerAzureADOnlyAuthentication(ctx *pulumi.Context, args *LookupServerAzureADOnlyAuthenticationArgs, opts ...pulumi.InvokeOption) (*LookupServerAzureADOnlyAuthenticationResult, error) {
	var rv LookupServerAzureADOnlyAuthenticationResult
	err := ctx.Invoke("azure-native:sql/v20211101preview:getServerAzureADOnlyAuthentication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerAzureADOnlyAuthenticationArgs struct {
	AuthenticationName string `pulumi:"authenticationName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	ServerName         string `pulumi:"serverName"`
}


type LookupServerAzureADOnlyAuthenticationResult struct {
	AzureADOnlyAuthentication bool   `pulumi:"azureADOnlyAuthentication"`
	Id                        string `pulumi:"id"`
	Name                      string `pulumi:"name"`
	Type                      string `pulumi:"type"`
}

func LookupServerAzureADOnlyAuthenticationOutput(ctx *pulumi.Context, args LookupServerAzureADOnlyAuthenticationOutputArgs, opts ...pulumi.InvokeOption) LookupServerAzureADOnlyAuthenticationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerAzureADOnlyAuthenticationResult, error) {
			args := v.(LookupServerAzureADOnlyAuthenticationArgs)
			r, err := LookupServerAzureADOnlyAuthentication(ctx, &args, opts...)
			var s LookupServerAzureADOnlyAuthenticationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerAzureADOnlyAuthenticationResultOutput)
}

type LookupServerAzureADOnlyAuthenticationOutputArgs struct {
	AuthenticationName pulumi.StringInput `pulumi:"authenticationName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName         pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerAzureADOnlyAuthenticationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAzureADOnlyAuthenticationArgs)(nil)).Elem()
}


type LookupServerAzureADOnlyAuthenticationResultOutput struct{ *pulumi.OutputState }

func (LookupServerAzureADOnlyAuthenticationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerAzureADOnlyAuthenticationResult)(nil)).Elem()
}

func (o LookupServerAzureADOnlyAuthenticationResultOutput) ToLookupServerAzureADOnlyAuthenticationResultOutput() LookupServerAzureADOnlyAuthenticationResultOutput {
	return o
}

func (o LookupServerAzureADOnlyAuthenticationResultOutput) ToLookupServerAzureADOnlyAuthenticationResultOutputWithContext(ctx context.Context) LookupServerAzureADOnlyAuthenticationResultOutput {
	return o
}

func (o LookupServerAzureADOnlyAuthenticationResultOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerAzureADOnlyAuthenticationResult) bool { return v.AzureADOnlyAuthentication }).(pulumi.BoolOutput)
}

func (o LookupServerAzureADOnlyAuthenticationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADOnlyAuthenticationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerAzureADOnlyAuthenticationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADOnlyAuthenticationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerAzureADOnlyAuthenticationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerAzureADOnlyAuthenticationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerAzureADOnlyAuthenticationResultOutput{})
}
