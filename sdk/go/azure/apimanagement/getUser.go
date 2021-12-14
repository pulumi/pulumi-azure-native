


package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:apimanagement:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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


func (val *LookupUserResult) Defaults() *LookupUserResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.State) {
		state_ := "active"
		tmp.State = &state_
	}
	return &tmp
}
