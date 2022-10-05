


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKey(ctx *pulumi.Context, args *LookupKeyArgs, opts ...pulumi.InvokeOption) (*LookupKeyResult, error) {
	var rv LookupKeyResult
	err := ctx.Invoke("azure-native:keyvault/v20200401preview:getKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyArgs struct {
	KeyName           string `pulumi:"keyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupKeyResult struct {
	Attributes        *KeyAttributesResponse `pulumi:"attributes"`
	CurveName         *string                `pulumi:"curveName"`
	Id                string                 `pulumi:"id"`
	KeyOps            []string               `pulumi:"keyOps"`
	KeySize           *int                   `pulumi:"keySize"`
	KeyUri            string                 `pulumi:"keyUri"`
	KeyUriWithVersion string                 `pulumi:"keyUriWithVersion"`
	Kty               *string                `pulumi:"kty"`
	Location          string                 `pulumi:"location"`
	Name              string                 `pulumi:"name"`
	Tags              map[string]string      `pulumi:"tags"`
	Type              string                 `pulumi:"type"`
}

func LookupKeyOutput(ctx *pulumi.Context, args LookupKeyOutputArgs, opts ...pulumi.InvokeOption) LookupKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKeyResult, error) {
			args := v.(LookupKeyArgs)
			r, err := LookupKey(ctx, &args, opts...)
			var s LookupKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKeyResultOutput)
}

type LookupKeyOutputArgs struct {
	KeyName           pulumi.StringInput `pulumi:"keyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VaultName         pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyArgs)(nil)).Elem()
}


type LookupKeyResultOutput struct{ *pulumi.OutputState }

func (LookupKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyResult)(nil)).Elem()
}

func (o LookupKeyResultOutput) ToLookupKeyResultOutput() LookupKeyResultOutput {
	return o
}

func (o LookupKeyResultOutput) ToLookupKeyResultOutputWithContext(ctx context.Context) LookupKeyResultOutput {
	return o
}

func (o LookupKeyResultOutput) Attributes() KeyAttributesResponsePtrOutput {
	return o.ApplyT(func(v LookupKeyResult) *KeyAttributesResponse { return v.Attributes }).(KeyAttributesResponsePtrOutput)
}

func (o LookupKeyResultOutput) CurveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyResult) *string { return v.CurveName }).(pulumi.StringPtrOutput)
}

func (o LookupKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKeyResultOutput) KeyOps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKeyResult) []string { return v.KeyOps }).(pulumi.StringArrayOutput)
}

func (o LookupKeyResultOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupKeyResult) *int { return v.KeySize }).(pulumi.IntPtrOutput)
}

func (o LookupKeyResultOutput) KeyUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.KeyUri }).(pulumi.StringOutput)
}

func (o LookupKeyResultOutput) KeyUriWithVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.KeyUriWithVersion }).(pulumi.StringOutput)
}

func (o LookupKeyResultOutput) Kty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyResult) *string { return v.Kty }).(pulumi.StringPtrOutput)
}

func (o LookupKeyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKeyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKeyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeyResultOutput{})
}
