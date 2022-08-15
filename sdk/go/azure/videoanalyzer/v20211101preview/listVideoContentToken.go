


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListVideoContentToken(ctx *pulumi.Context, args *ListVideoContentTokenArgs, opts ...pulumi.InvokeOption) (*ListVideoContentTokenResult, error) {
	var rv ListVideoContentTokenResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20211101preview:listVideoContentToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListVideoContentTokenArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VideoName         string `pulumi:"videoName"`
}


type ListVideoContentTokenResult struct {
	ExpirationDate string `pulumi:"expirationDate"`
	Token          string `pulumi:"token"`
}

func ListVideoContentTokenOutput(ctx *pulumi.Context, args ListVideoContentTokenOutputArgs, opts ...pulumi.InvokeOption) ListVideoContentTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVideoContentTokenResult, error) {
			args := v.(ListVideoContentTokenArgs)
			r, err := ListVideoContentToken(ctx, &args, opts...)
			var s ListVideoContentTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVideoContentTokenResultOutput)
}

type ListVideoContentTokenOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VideoName         pulumi.StringInput `pulumi:"videoName"`
}

func (ListVideoContentTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVideoContentTokenArgs)(nil)).Elem()
}


type ListVideoContentTokenResultOutput struct{ *pulumi.OutputState }

func (ListVideoContentTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVideoContentTokenResult)(nil)).Elem()
}

func (o ListVideoContentTokenResultOutput) ToListVideoContentTokenResultOutput() ListVideoContentTokenResultOutput {
	return o
}

func (o ListVideoContentTokenResultOutput) ToListVideoContentTokenResultOutputWithContext(ctx context.Context) ListVideoContentTokenResultOutput {
	return o
}

func (o ListVideoContentTokenResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v ListVideoContentTokenResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o ListVideoContentTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v ListVideoContentTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVideoContentTokenResultOutput{})
}
