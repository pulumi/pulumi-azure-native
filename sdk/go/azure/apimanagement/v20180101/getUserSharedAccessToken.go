


package v20180101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUserSharedAccessToken(ctx *pulumi.Context, args *GetUserSharedAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetUserSharedAccessTokenResult, error) {
	var rv GetUserSharedAccessTokenResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getUserSharedAccessToken", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetUserSharedAccessTokenArgs struct {
	Expiry            string  `pulumi:"expiry"`
	KeyType           KeyType `pulumi:"keyType"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	Uid               string  `pulumi:"uid"`
}


func (val *GetUserSharedAccessTokenArgs) Defaults() *GetUserSharedAccessTokenArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeyType) {
		tmp.KeyType = KeyType("primary")
	}
	return &tmp
}


type GetUserSharedAccessTokenResult struct {
	Value *string `pulumi:"value"`
}
