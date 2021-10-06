


package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	UserId            string `pulumi:"userId"`
}


type LookupUserResult struct {
	Email            *string                           `pulumi:"email"`
	FirstName        *string                           `pulumi:"firstName"`
	Groups           []GroupContractPropertiesResponse `pulumi:"groups"`
	Id               string                            `pulumi:"id"`
	Identities       []UserIdentityContractResponse    `pulumi:"identities"`
	LastName         *string                           `pulumi:"lastName"`
	Name             string                            `pulumi:"name"`
	Note             *string                           `pulumi:"note"`
	RegistrationDate *string                           `pulumi:"registrationDate"`
	State            *string                           `pulumi:"state"`
	Type             string                            `pulumi:"type"`
}
