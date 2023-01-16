


package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCredentialSet(ctx *pulumi.Context, args *LookupCredentialSetArgs, opts ...pulumi.InvokeOption) (*LookupCredentialSetResult, error) {
	var rv LookupCredentialSetResult
	err := ctx.Invoke("azure-native:containerregistry/v20230101preview:getCredentialSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCredentialSetArgs struct {
	CredentialSetName string `pulumi:"credentialSetName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCredentialSetResult struct {
	AuthCredentials   []AuthCredentialResponse    `pulumi:"authCredentials"`
	CreationDate      string                      `pulumi:"creationDate"`
	Id                string                      `pulumi:"id"`
	Identity          *IdentityPropertiesResponse `pulumi:"identity"`
	LoginServer       *string                     `pulumi:"loginServer"`
	Name              string                      `pulumi:"name"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	SystemData        SystemDataResponse          `pulumi:"systemData"`
	Type              string                      `pulumi:"type"`
}

func LookupCredentialSetOutput(ctx *pulumi.Context, args LookupCredentialSetOutputArgs, opts ...pulumi.InvokeOption) LookupCredentialSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCredentialSetResult, error) {
			args := v.(LookupCredentialSetArgs)
			r, err := LookupCredentialSet(ctx, &args, opts...)
			var s LookupCredentialSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCredentialSetResultOutput)
}

type LookupCredentialSetOutputArgs struct {
	CredentialSetName pulumi.StringInput `pulumi:"credentialSetName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCredentialSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialSetArgs)(nil)).Elem()
}


type LookupCredentialSetResultOutput struct{ *pulumi.OutputState }

func (LookupCredentialSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialSetResult)(nil)).Elem()
}

func (o LookupCredentialSetResultOutput) ToLookupCredentialSetResultOutput() LookupCredentialSetResultOutput {
	return o
}

func (o LookupCredentialSetResultOutput) ToLookupCredentialSetResultOutputWithContext(ctx context.Context) LookupCredentialSetResultOutput {
	return o
}

func (o LookupCredentialSetResultOutput) AuthCredentials() AuthCredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) []AuthCredentialResponse { return v.AuthCredentials }).(AuthCredentialResponseArrayOutput)
}

func (o LookupCredentialSetResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupCredentialSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCredentialSetResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupCredentialSetResultOutput) LoginServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) *string { return v.LoginServer }).(pulumi.StringPtrOutput)
}

func (o LookupCredentialSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCredentialSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCredentialSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCredentialSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCredentialSetResultOutput{})
}
