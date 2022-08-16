


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamedValue(ctx *pulumi.Context, args *LookupNamedValueArgs, opts ...pulumi.InvokeOption) (*LookupNamedValueResult, error) {
	var rv LookupNamedValueResult
	err := ctx.Invoke("azure-native:apimanagement/v20210101preview:getNamedValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamedValueArgs struct {
	NamedValueId      string `pulumi:"namedValueId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupNamedValueResult struct {
	DisplayName string                              `pulumi:"displayName"`
	Id          string                              `pulumi:"id"`
	KeyVault    *KeyVaultContractPropertiesResponse `pulumi:"keyVault"`
	Name        string                              `pulumi:"name"`
	Secret      *bool                               `pulumi:"secret"`
	Tags        []string                            `pulumi:"tags"`
	Type        string                              `pulumi:"type"`
	Value       *string                             `pulumi:"value"`
}

func LookupNamedValueOutput(ctx *pulumi.Context, args LookupNamedValueOutputArgs, opts ...pulumi.InvokeOption) LookupNamedValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamedValueResult, error) {
			args := v.(LookupNamedValueArgs)
			r, err := LookupNamedValue(ctx, &args, opts...)
			var s LookupNamedValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamedValueResultOutput)
}

type LookupNamedValueOutputArgs struct {
	NamedValueId      pulumi.StringInput `pulumi:"namedValueId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupNamedValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamedValueArgs)(nil)).Elem()
}


type LookupNamedValueResultOutput struct{ *pulumi.OutputState }

func (LookupNamedValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamedValueResult)(nil)).Elem()
}

func (o LookupNamedValueResultOutput) ToLookupNamedValueResultOutput() LookupNamedValueResultOutput {
	return o
}

func (o LookupNamedValueResultOutput) ToLookupNamedValueResultOutputWithContext(ctx context.Context) LookupNamedValueResultOutput {
	return o
}

func (o LookupNamedValueResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedValueResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupNamedValueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedValueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamedValueResultOutput) KeyVault() KeyVaultContractPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupNamedValueResult) *KeyVaultContractPropertiesResponse { return v.KeyVault }).(KeyVaultContractPropertiesResponsePtrOutput)
}

func (o LookupNamedValueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedValueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamedValueResultOutput) Secret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamedValueResult) *bool { return v.Secret }).(pulumi.BoolPtrOutput)
}

func (o LookupNamedValueResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNamedValueResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupNamedValueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedValueResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNamedValueResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamedValueResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamedValueResultOutput{})
}
