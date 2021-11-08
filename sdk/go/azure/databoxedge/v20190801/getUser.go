


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:databoxedge/v20190801:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupUserResult struct {
	EncryptedPassword *AsymmetricEncryptedSecretResponse `pulumi:"encryptedPassword"`
	Id                string                             `pulumi:"id"`
	Name              string                             `pulumi:"name"`
	ShareAccessRights []ShareAccessRightResponse         `pulumi:"shareAccessRights"`
	Type              string                             `pulumi:"type"`
	UserType          string                             `pulumi:"userType"`
}
