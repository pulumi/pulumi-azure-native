


package managedidentity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFederatedIdentityCredential(ctx *pulumi.Context, args *LookupFederatedIdentityCredentialArgs, opts ...pulumi.InvokeOption) (*LookupFederatedIdentityCredentialResult, error) {
	var rv LookupFederatedIdentityCredentialResult
	err := ctx.Invoke("azure-native:managedidentity:getFederatedIdentityCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFederatedIdentityCredentialArgs struct {
	FederatedIdentityCredentialResourceName string `pulumi:"federatedIdentityCredentialResourceName"`
	ResourceGroupName                       string `pulumi:"resourceGroupName"`
	ResourceName                            string `pulumi:"resourceName"`
}


type LookupFederatedIdentityCredentialResult struct {
	Audiences []string `pulumi:"audiences"`
	Id        string   `pulumi:"id"`
	Issuer    string   `pulumi:"issuer"`
	Name      string   `pulumi:"name"`
	Subject   string   `pulumi:"subject"`
	Type      string   `pulumi:"type"`
}

func LookupFederatedIdentityCredentialOutput(ctx *pulumi.Context, args LookupFederatedIdentityCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedIdentityCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedIdentityCredentialResult, error) {
			args := v.(LookupFederatedIdentityCredentialArgs)
			r, err := LookupFederatedIdentityCredential(ctx, &args, opts...)
			var s LookupFederatedIdentityCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedIdentityCredentialResultOutput)
}

type LookupFederatedIdentityCredentialOutputArgs struct {
	FederatedIdentityCredentialResourceName pulumi.StringInput `pulumi:"federatedIdentityCredentialResourceName"`
	ResourceGroupName                       pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                            pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupFederatedIdentityCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedIdentityCredentialArgs)(nil)).Elem()
}


type LookupFederatedIdentityCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedIdentityCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedIdentityCredentialResult)(nil)).Elem()
}

func (o LookupFederatedIdentityCredentialResultOutput) ToLookupFederatedIdentityCredentialResultOutput() LookupFederatedIdentityCredentialResultOutput {
	return o
}

func (o LookupFederatedIdentityCredentialResultOutput) ToLookupFederatedIdentityCredentialResultOutputWithContext(ctx context.Context) LookupFederatedIdentityCredentialResultOutput {
	return o
}

func (o LookupFederatedIdentityCredentialResultOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedIdentityCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedIdentityCredentialResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o LookupFederatedIdentityCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFederatedIdentityCredentialResultOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Subject }).(pulumi.StringOutput)
}

func (o LookupFederatedIdentityCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedIdentityCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedIdentityCredentialResultOutput{})
}
