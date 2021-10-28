


package v20201030preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:powerplatform/v20201030preview:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	Description *string            `pulumi:"description"`
	Id          string             `pulumi:"id"`
	Location    string             `pulumi:"location"`
	Name        string             `pulumi:"name"`
	SystemData  SystemDataResponse `pulumi:"systemData"`
	Tags        map[string]string  `pulumi:"tags"`
	Type        string             `pulumi:"type"`
}
