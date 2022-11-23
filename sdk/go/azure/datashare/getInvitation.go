


package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInvitation(ctx *pulumi.Context, args *LookupInvitationArgs, opts ...pulumi.InvokeOption) (*LookupInvitationResult, error) {
	var rv LookupInvitationResult
	err := ctx.Invoke("azure-native:datashare:getInvitation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInvitationArgs struct {
	AccountName       string `pulumi:"accountName"`
	InvitationName    string `pulumi:"invitationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupInvitationResult struct {
	ExpirationDate          *string            `pulumi:"expirationDate"`
	Id                      string             `pulumi:"id"`
	InvitationId            string             `pulumi:"invitationId"`
	InvitationStatus        string             `pulumi:"invitationStatus"`
	Name                    string             `pulumi:"name"`
	RespondedAt             string             `pulumi:"respondedAt"`
	SentAt                  string             `pulumi:"sentAt"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	TargetActiveDirectoryId *string            `pulumi:"targetActiveDirectoryId"`
	TargetEmail             *string            `pulumi:"targetEmail"`
	TargetObjectId          *string            `pulumi:"targetObjectId"`
	Type                    string             `pulumi:"type"`
	UserEmail               string             `pulumi:"userEmail"`
	UserName                string             `pulumi:"userName"`
}

func LookupInvitationOutput(ctx *pulumi.Context, args LookupInvitationOutputArgs, opts ...pulumi.InvokeOption) LookupInvitationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInvitationResult, error) {
			args := v.(LookupInvitationArgs)
			r, err := LookupInvitation(ctx, &args, opts...)
			var s LookupInvitationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInvitationResultOutput)
}

type LookupInvitationOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	InvitationName    pulumi.StringInput `pulumi:"invitationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupInvitationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInvitationArgs)(nil)).Elem()
}


type LookupInvitationResultOutput struct{ *pulumi.OutputState }

func (LookupInvitationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInvitationResult)(nil)).Elem()
}

func (o LookupInvitationResultOutput) ToLookupInvitationResultOutput() LookupInvitationResultOutput {
	return o
}

func (o LookupInvitationResultOutput) ToLookupInvitationResultOutputWithContext(ctx context.Context) LookupInvitationResultOutput {
	return o
}

func (o LookupInvitationResultOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvitationResult) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o LookupInvitationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInvitationResultOutput) InvitationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.InvitationId }).(pulumi.StringOutput)
}

func (o LookupInvitationResultOutput) InvitationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.InvitationStatus }).(pulumi.StringOutput)
}

func (o LookupInvitationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInvitationResultOutput) RespondedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.RespondedAt }).(pulumi.StringOutput)
}

func (o LookupInvitationResultOutput) SentAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.SentAt }).(pulumi.StringOutput)
}

func (o LookupInvitationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInvitationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupInvitationResultOutput) TargetActiveDirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvitationResult) *string { return v.TargetActiveDirectoryId }).(pulumi.StringPtrOutput)
}

func (o LookupInvitationResultOutput) TargetEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvitationResult) *string { return v.TargetEmail }).(pulumi.StringPtrOutput)
}

func (o LookupInvitationResultOutput) TargetObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInvitationResult) *string { return v.TargetObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupInvitationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupInvitationResultOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.UserEmail }).(pulumi.StringOutput)
}

func (o LookupInvitationResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInvitationResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInvitationResultOutput{})
}
