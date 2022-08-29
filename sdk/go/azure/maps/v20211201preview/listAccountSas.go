


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountSas(ctx *pulumi.Context, args *ListAccountSasArgs, opts ...pulumi.InvokeOption) (*ListAccountSasResult, error) {
	var rv ListAccountSasResult
	err := ctx.Invoke("azure-native:maps/v20211201preview:listAccountSas", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountSasArgs struct {
	AccountName       string   `pulumi:"accountName"`
	Expiry            string   `pulumi:"expiry"`
	MaxRatePerSecond  int      `pulumi:"maxRatePerSecond"`
	PrincipalId       string   `pulumi:"principalId"`
	Regions           []string `pulumi:"regions"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	SigningKey        string   `pulumi:"signingKey"`
	Start             string   `pulumi:"start"`
}


func (val *ListAccountSasArgs) Defaults() *ListAccountSasArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxRatePerSecond) {
		tmp.MaxRatePerSecond = 500
	}
	return &tmp
}


type ListAccountSasResult struct {
	AccountSasToken string `pulumi:"accountSasToken"`
}

func ListAccountSasOutput(ctx *pulumi.Context, args ListAccountSasOutputArgs, opts ...pulumi.InvokeOption) ListAccountSasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAccountSasResult, error) {
			args := v.(ListAccountSasArgs)
			r, err := ListAccountSas(ctx, &args, opts...)
			var s ListAccountSasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAccountSasResultOutput)
}

type ListAccountSasOutputArgs struct {
	AccountName       pulumi.StringInput      `pulumi:"accountName"`
	Expiry            pulumi.StringInput      `pulumi:"expiry"`
	MaxRatePerSecond  pulumi.IntInput         `pulumi:"maxRatePerSecond"`
	PrincipalId       pulumi.StringInput      `pulumi:"principalId"`
	Regions           pulumi.StringArrayInput `pulumi:"regions"`
	ResourceGroupName pulumi.StringInput      `pulumi:"resourceGroupName"`
	SigningKey        pulumi.StringInput      `pulumi:"signingKey"`
	Start             pulumi.StringInput      `pulumi:"start"`
}

func (ListAccountSasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountSasArgs)(nil)).Elem()
}


type ListAccountSasResultOutput struct{ *pulumi.OutputState }

func (ListAccountSasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountSasResult)(nil)).Elem()
}

func (o ListAccountSasResultOutput) ToListAccountSasResultOutput() ListAccountSasResultOutput {
	return o
}

func (o ListAccountSasResultOutput) ToListAccountSasResultOutputWithContext(ctx context.Context) ListAccountSasResultOutput {
	return o
}

func (o ListAccountSasResultOutput) AccountSasToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountSasResult) string { return v.AccountSasToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountSasResultOutput{})
}
