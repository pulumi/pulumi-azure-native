


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstanceAzureADOnlyAuthentication(ctx *pulumi.Context, args *LookupManagedInstanceAzureADOnlyAuthenticationArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceAzureADOnlyAuthenticationResult, error) {
	var rv LookupManagedInstanceAzureADOnlyAuthenticationResult
	err := ctx.Invoke("azure-native:sql/v20200801preview:getManagedInstanceAzureADOnlyAuthentication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceAzureADOnlyAuthenticationArgs struct {
	AuthenticationName  string `pulumi:"authenticationName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedInstanceAzureADOnlyAuthenticationResult struct {
	AzureADOnlyAuthentication bool   `pulumi:"azureADOnlyAuthentication"`
	Id                        string `pulumi:"id"`
	Name                      string `pulumi:"name"`
	Type                      string `pulumi:"type"`
}

func LookupManagedInstanceAzureADOnlyAuthenticationOutput(ctx *pulumi.Context, args LookupManagedInstanceAzureADOnlyAuthenticationOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceAzureADOnlyAuthenticationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceAzureADOnlyAuthenticationResult, error) {
			args := v.(LookupManagedInstanceAzureADOnlyAuthenticationArgs)
			r, err := LookupManagedInstanceAzureADOnlyAuthentication(ctx, &args, opts...)
			var s LookupManagedInstanceAzureADOnlyAuthenticationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceAzureADOnlyAuthenticationResultOutput)
}

type LookupManagedInstanceAzureADOnlyAuthenticationOutputArgs struct {
	AuthenticationName  pulumi.StringInput `pulumi:"authenticationName"`
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceAzureADOnlyAuthenticationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceAzureADOnlyAuthenticationArgs)(nil)).Elem()
}


type LookupManagedInstanceAzureADOnlyAuthenticationResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceAzureADOnlyAuthenticationResult)(nil)).Elem()
}

func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) ToLookupManagedInstanceAzureADOnlyAuthenticationResultOutput() LookupManagedInstanceAzureADOnlyAuthenticationResultOutput {
	return o
}

func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) ToLookupManagedInstanceAzureADOnlyAuthenticationResultOutputWithContext(ctx context.Context) LookupManagedInstanceAzureADOnlyAuthenticationResultOutput {
	return o
}

func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) AzureADOnlyAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupManagedInstanceAzureADOnlyAuthenticationResult) bool { return v.AzureADOnlyAuthentication }).(pulumi.BoolOutput)
}

func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAzureADOnlyAuthenticationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAzureADOnlyAuthenticationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedInstanceAzureADOnlyAuthenticationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceAzureADOnlyAuthenticationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceAzureADOnlyAuthenticationResultOutput{})
}
