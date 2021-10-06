


package v20211101preview

import (
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
