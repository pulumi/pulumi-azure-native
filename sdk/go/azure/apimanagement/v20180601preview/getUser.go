


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getUser", args, &rv, opts...)
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

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

type LookupUserOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	UserId            pulumi.StringInput `pulumi:"userId"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}


type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) Groups() GroupContractPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []GroupContractPropertiesResponse { return v.Groups }).(GroupContractPropertiesResponseArrayOutput)
}

func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Identities() UserIdentityContractResponseArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []UserIdentityContractResponse { return v.Identities }).(UserIdentityContractResponseArrayOutput)
}

func (o LookupUserResultOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Note }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) RegistrationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.RegistrationDate }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
