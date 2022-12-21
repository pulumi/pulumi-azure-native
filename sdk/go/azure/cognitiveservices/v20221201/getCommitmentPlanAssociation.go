


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCommitmentPlanAssociation(ctx *pulumi.Context, args *LookupCommitmentPlanAssociationArgs, opts ...pulumi.InvokeOption) (*LookupCommitmentPlanAssociationResult, error) {
	var rv LookupCommitmentPlanAssociationResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20221201:getCommitmentPlanAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCommitmentPlanAssociationArgs struct {
	CommitmentPlanAssociationName string `pulumi:"commitmentPlanAssociationName"`
	CommitmentPlanName            string `pulumi:"commitmentPlanName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupCommitmentPlanAssociationResult struct {
	AccountId  *string            `pulumi:"accountId"`
	Etag       string             `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupCommitmentPlanAssociationOutput(ctx *pulumi.Context, args LookupCommitmentPlanAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupCommitmentPlanAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCommitmentPlanAssociationResult, error) {
			args := v.(LookupCommitmentPlanAssociationArgs)
			r, err := LookupCommitmentPlanAssociation(ctx, &args, opts...)
			var s LookupCommitmentPlanAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCommitmentPlanAssociationResultOutput)
}

type LookupCommitmentPlanAssociationOutputArgs struct {
	CommitmentPlanAssociationName pulumi.StringInput `pulumi:"commitmentPlanAssociationName"`
	CommitmentPlanName            pulumi.StringInput `pulumi:"commitmentPlanName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCommitmentPlanAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommitmentPlanAssociationArgs)(nil)).Elem()
}


type LookupCommitmentPlanAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupCommitmentPlanAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommitmentPlanAssociationResult)(nil)).Elem()
}

func (o LookupCommitmentPlanAssociationResultOutput) ToLookupCommitmentPlanAssociationResultOutput() LookupCommitmentPlanAssociationResultOutput {
	return o
}

func (o LookupCommitmentPlanAssociationResultOutput) ToLookupCommitmentPlanAssociationResultOutputWithContext(ctx context.Context) LookupCommitmentPlanAssociationResultOutput {
	return o
}

func (o LookupCommitmentPlanAssociationResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o LookupCommitmentPlanAssociationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCommitmentPlanAssociationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCommitmentPlanAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommitmentPlanAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommitmentPlanAssociationResultOutput{})
}
