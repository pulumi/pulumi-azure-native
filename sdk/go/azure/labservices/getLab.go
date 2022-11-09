


package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:labservices:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabArgs struct {
	Expand            *string `pulumi:"expand"`
	LabAccountName    string  `pulumi:"labAccountName"`
	LabName           string  `pulumi:"labName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLabResult struct {
	CreatedByObjectId          string                        `pulumi:"createdByObjectId"`
	CreatedByUserPrincipalName string                        `pulumi:"createdByUserPrincipalName"`
	CreatedDate                string                        `pulumi:"createdDate"`
	Id                         string                        `pulumi:"id"`
	InvitationCode             string                        `pulumi:"invitationCode"`
	LatestOperationResult      LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location                   *string                       `pulumi:"location"`
	MaxUsersInLab              *int                          `pulumi:"maxUsersInLab"`
	Name                       string                        `pulumi:"name"`
	ProvisioningState          *string                       `pulumi:"provisioningState"`
	Tags                       map[string]string             `pulumi:"tags"`
	Type                       string                        `pulumi:"type"`
	UniqueIdentifier           *string                       `pulumi:"uniqueIdentifier"`
	UsageQuota                 *string                       `pulumi:"usageQuota"`
	UserAccessMode             *string                       `pulumi:"userAccessMode"`
	UserQuota                  int                           `pulumi:"userQuota"`
}

func LookupLabOutput(ctx *pulumi.Context, args LookupLabOutputArgs, opts ...pulumi.InvokeOption) LookupLabResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabResult, error) {
			args := v.(LookupLabArgs)
			r, err := LookupLab(ctx, &args, opts...)
			var s LookupLabResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabResultOutput)
}

type LookupLabOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabAccountName    pulumi.StringInput    `pulumi:"labAccountName"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupLabOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabArgs)(nil)).Elem()
}


type LookupLabResultOutput struct{ *pulumi.OutputState }

func (LookupLabResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResult)(nil)).Elem()
}

func (o LookupLabResultOutput) ToLookupLabResultOutput() LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) ToLookupLabResultOutputWithContext(ctx context.Context) LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) CreatedByObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.CreatedByObjectId }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) CreatedByUserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.CreatedByUserPrincipalName }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) InvitationCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.InvitationCode }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupLabResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o LookupLabResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) MaxUsersInLab() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *int { return v.MaxUsersInLab }).(pulumi.IntPtrOutput)
}

func (o LookupLabResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLabResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) UsageQuota() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.UsageQuota }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) UserAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.UserAccessMode }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) UserQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLabResult) int { return v.UserQuota }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabResultOutput{})
}
