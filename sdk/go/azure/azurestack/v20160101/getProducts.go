


package v20160101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProducts(ctx *pulumi.Context, args *GetProductsArgs, opts ...pulumi.InvokeOption) (*GetProductsResult, error) {
	var rv GetProductsResult
	err := ctx.Invoke("azure-native:azurestack/v20160101:getProducts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProductsArgs struct {
	ProductName      string `pulumi:"productName"`
	RegistrationName string `pulumi:"registrationName"`
	ResourceGroup    string `pulumi:"resourceGroup"`
}


type GetProductsResult struct {
	NextLink *string           `pulumi:"nextLink"`
	Value    []ProductResponse `pulumi:"value"`
}

func GetProductsOutput(ctx *pulumi.Context, args GetProductsOutputArgs, opts ...pulumi.InvokeOption) GetProductsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProductsResult, error) {
			args := v.(GetProductsArgs)
			r, err := GetProducts(ctx, &args, opts...)
			var s GetProductsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProductsResultOutput)
}

type GetProductsOutputArgs struct {
	ProductName      pulumi.StringInput `pulumi:"productName"`
	RegistrationName pulumi.StringInput `pulumi:"registrationName"`
	ResourceGroup    pulumi.StringInput `pulumi:"resourceGroup"`
}

func (GetProductsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductsArgs)(nil)).Elem()
}


type GetProductsResultOutput struct{ *pulumi.OutputState }

func (GetProductsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductsResult)(nil)).Elem()
}

func (o GetProductsResultOutput) ToGetProductsResultOutput() GetProductsResultOutput {
	return o
}

func (o GetProductsResultOutput) ToGetProductsResultOutputWithContext(ctx context.Context) GetProductsResultOutput {
	return o
}

func (o GetProductsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o GetProductsResultOutput) Value() ProductResponseArrayOutput {
	return o.ApplyT(func(v GetProductsResult) []ProductResponse { return v.Value }).(ProductResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductsResultOutput{})
}
