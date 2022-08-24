


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKey(ctx *pulumi.Context, args *LookupKeyArgs, opts ...pulumi.InvokeOption) (*LookupKeyResult, error) {
	var rv LookupKeyResult
	err := ctx.Invoke("azure-native:synapse/v20210401preview:getKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyArgs struct {
	KeyName           string `pulumi:"keyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupKeyResult struct {
	Id          string  `pulumi:"id"`
	IsActiveCMK *bool   `pulumi:"isActiveCMK"`
	KeyVaultUrl *string `pulumi:"keyVaultUrl"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
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
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
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

func (o LookupKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKeyResultOutput) IsActiveCMK() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupKeyResult) *bool { return v.IsActiveCMK }).(pulumi.BoolPtrOutput)
}

func (o LookupKeyResultOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKeyResult) *string { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o LookupKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeyResultOutput{})
}
