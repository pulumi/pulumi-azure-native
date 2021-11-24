


package v20160101

import (
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
