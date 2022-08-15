


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVault(ctx *pulumi.Context, args *LookupVaultArgs, opts ...pulumi.InvokeOption) (*LookupVaultResult, error) {
	var rv LookupVaultResult
	err := ctx.Invoke("azure-native:keyvault/v20210401preview:getVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVaultArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupVaultResult struct {
	Id         string                  `pulumi:"id"`
	Location   *string                 `pulumi:"location"`
	Name       string                  `pulumi:"name"`
	Properties VaultPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse      `pulumi:"systemData"`
	Tags       map[string]string       `pulumi:"tags"`
	Type       string                  `pulumi:"type"`
}


func (val *LookupVaultResult) Defaults() *LookupVaultResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupVaultOutput(ctx *pulumi.Context, args LookupVaultOutputArgs, opts ...pulumi.InvokeOption) LookupVaultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVaultResult, error) {
			args := v.(LookupVaultArgs)
			r, err := LookupVault(ctx, &args, opts...)
			var s LookupVaultResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVaultResultOutput)
}

type LookupVaultOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupVaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVaultArgs)(nil)).Elem()
}


type LookupVaultResultOutput struct{ *pulumi.OutputState }

func (LookupVaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVaultResult)(nil)).Elem()
}

func (o LookupVaultResultOutput) ToLookupVaultResultOutput() LookupVaultResultOutput {
	return o
}

func (o LookupVaultResultOutput) ToLookupVaultResultOutputWithContext(ctx context.Context) LookupVaultResultOutput {
	return o
}

func (o LookupVaultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVaultResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVaultResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVaultResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVaultResultOutput) Properties() VaultPropertiesResponseOutput {
	return o.ApplyT(func(v LookupVaultResult) VaultPropertiesResponse { return v.Properties }).(VaultPropertiesResponseOutput)
}

func (o LookupVaultResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVaultResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVaultResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVaultResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVaultResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVaultResultOutput{})
}
