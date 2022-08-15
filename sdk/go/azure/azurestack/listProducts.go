


package azurestack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProducts(ctx *pulumi.Context, args *ListProductsArgs, opts ...pulumi.InvokeOption) (*ListProductsResult, error) {
	var rv ListProductsResult
	err := ctx.Invoke("azure-native:azurestack:listProducts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductsArgs struct {
	ProductName      string `pulumi:"productName"`
	RegistrationName string `pulumi:"registrationName"`
	ResourceGroup    string `pulumi:"resourceGroup"`
}


type ListProductsResult struct {
	NextLink *string           `pulumi:"nextLink"`
	Value    []ProductResponse `pulumi:"value"`
}

func ListProductsOutput(ctx *pulumi.Context, args ListProductsOutputArgs, opts ...pulumi.InvokeOption) ListProductsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListProductsResult, error) {
			args := v.(ListProductsArgs)
			r, err := ListProducts(ctx, &args, opts...)
			var s ListProductsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListProductsResultOutput)
}

type ListProductsOutputArgs struct {
	ProductName      pulumi.StringInput `pulumi:"productName"`
	RegistrationName pulumi.StringInput `pulumi:"registrationName"`
	ResourceGroup    pulumi.StringInput `pulumi:"resourceGroup"`
}

func (ListProductsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsArgs)(nil)).Elem()
}


type ListProductsResultOutput struct{ *pulumi.OutputState }

func (ListProductsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListProductsResult)(nil)).Elem()
}

func (o ListProductsResultOutput) ToListProductsResultOutput() ListProductsResultOutput {
	return o
}

func (o ListProductsResultOutput) ToListProductsResultOutputWithContext(ctx context.Context) ListProductsResultOutput {
	return o
}

func (o ListProductsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListProductsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListProductsResultOutput) Value() ProductResponseArrayOutput {
	return o.ApplyT(func(v ListProductsResult) []ProductResponse { return v.Value }).(ProductResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListProductsResultOutput{})
}
