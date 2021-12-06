


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupShare(ctx *pulumi.Context, args *LookupShareArgs, opts ...pulumi.InvokeOption) (*LookupShareResult, error) {
	var rv LookupShareResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupShareArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupShareResult struct {
	CreatedAt         string             `pulumi:"createdAt"`
	Description       *string            `pulumi:"description"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	ShareKind         *string            `pulumi:"shareKind"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Terms             *string            `pulumi:"terms"`
	Type              string             `pulumi:"type"`
	UserEmail         string             `pulumi:"userEmail"`
	UserName          string             `pulumi:"userName"`
}
