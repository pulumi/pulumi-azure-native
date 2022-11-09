


package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	Expand            *string `pulumi:"expand"`
	LabAccountName    string  `pulumi:"labAccountName"`
	LabName           string  `pulumi:"labName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserName          string  `pulumi:"userName"`
}


type LookupUserResult struct {
	Email                 string                        `pulumi:"email"`
	FamilyName            string                        `pulumi:"familyName"`
	GivenName             string                        `pulumi:"givenName"`
	Id                    string                        `pulumi:"id"`
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location              *string                       `pulumi:"location"`
	Name                  string                        `pulumi:"name"`
	ProvisioningState     *string                       `pulumi:"provisioningState"`
	Tags                  map[string]string             `pulumi:"tags"`
	TenantId              string                        `pulumi:"tenantId"`
	TotalUsage            string                        `pulumi:"totalUsage"`
	Type                  string                        `pulumi:"type"`
	UniqueIdentifier      *string                       `pulumi:"uniqueIdentifier"`
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
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabAccountName    pulumi.StringInput    `pulumi:"labAccountName"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	UserName          pulumi.StringInput    `pulumi:"userName"`
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

func (o LookupUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) FamilyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.FamilyName }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) GivenName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.GivenName }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupUserResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o LookupUserResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupUserResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupUserResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) TotalUsage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.TotalUsage }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
