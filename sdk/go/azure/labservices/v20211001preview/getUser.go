


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:labservices/v20211001preview:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	LabName           string `pulumi:"labName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	UserName          string `pulumi:"userName"`
}


type LookupUserResult struct {
	AdditionalUsageQuota *string            `pulumi:"additionalUsageQuota"`
	DisplayName          string             `pulumi:"displayName"`
	Email                string             `pulumi:"email"`
	Id                   string             `pulumi:"id"`
	InvitationSent       string             `pulumi:"invitationSent"`
	InvitationState      string             `pulumi:"invitationState"`
	Name                 string             `pulumi:"name"`
	ProvisioningState    string             `pulumi:"provisioningState"`
	RegistrationState    string             `pulumi:"registrationState"`
	SystemData           SystemDataResponse `pulumi:"systemData"`
	TotalUsage           string             `pulumi:"totalUsage"`
	Type                 string             `pulumi:"type"`
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
	LabName           pulumi.StringInput `pulumi:"labName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	UserName          pulumi.StringInput `pulumi:"userName"`
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

func (o LookupUserResultOutput) AdditionalUsageQuota() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.AdditionalUsageQuota }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) InvitationSent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.InvitationSent }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) InvitationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.InvitationState }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) RegistrationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.RegistrationState }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUserResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupUserResultOutput) TotalUsage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.TotalUsage }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
