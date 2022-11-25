


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlobContainerImmutabilityPolicy(ctx *pulumi.Context, args *LookupBlobContainerImmutabilityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerImmutabilityPolicyResult, error) {
	var rv LookupBlobContainerImmutabilityPolicyResult
	err := ctx.Invoke("azure-native:storage/v20210801:getBlobContainerImmutabilityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerImmutabilityPolicyArgs struct {
	AccountName            string `pulumi:"accountName"`
	ContainerName          string `pulumi:"containerName"`
	ImmutabilityPolicyName string `pulumi:"immutabilityPolicyName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupBlobContainerImmutabilityPolicyResult struct {
	AllowProtectedAppendWrites            *bool  `pulumi:"allowProtectedAppendWrites"`
	AllowProtectedAppendWritesAll         *bool  `pulumi:"allowProtectedAppendWritesAll"`
	Etag                                  string `pulumi:"etag"`
	Id                                    string `pulumi:"id"`
	ImmutabilityPeriodSinceCreationInDays *int   `pulumi:"immutabilityPeriodSinceCreationInDays"`
	Name                                  string `pulumi:"name"`
	State                                 string `pulumi:"state"`
	Type                                  string `pulumi:"type"`
}

func LookupBlobContainerImmutabilityPolicyOutput(ctx *pulumi.Context, args LookupBlobContainerImmutabilityPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBlobContainerImmutabilityPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobContainerImmutabilityPolicyResult, error) {
			args := v.(LookupBlobContainerImmutabilityPolicyArgs)
			r, err := LookupBlobContainerImmutabilityPolicy(ctx, &args, opts...)
			var s LookupBlobContainerImmutabilityPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobContainerImmutabilityPolicyResultOutput)
}

type LookupBlobContainerImmutabilityPolicyOutputArgs struct {
	AccountName            pulumi.StringInput `pulumi:"accountName"`
	ContainerName          pulumi.StringInput `pulumi:"containerName"`
	ImmutabilityPolicyName pulumi.StringInput `pulumi:"immutabilityPolicyName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobContainerImmutabilityPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerImmutabilityPolicyArgs)(nil)).Elem()
}


type LookupBlobContainerImmutabilityPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBlobContainerImmutabilityPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerImmutabilityPolicyResult)(nil)).Elem()
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) ToLookupBlobContainerImmutabilityPolicyResultOutput() LookupBlobContainerImmutabilityPolicyResultOutput {
	return o
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) ToLookupBlobContainerImmutabilityPolicyResultOutputWithContext(ctx context.Context) LookupBlobContainerImmutabilityPolicyResultOutput {
	return o
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) AllowProtectedAppendWrites() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) *bool { return v.AllowProtectedAppendWrites }).(pulumi.BoolPtrOutput)
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) AllowProtectedAppendWritesAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) *bool { return v.AllowProtectedAppendWritesAll }).(pulumi.BoolPtrOutput)
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) *int {
		return v.ImmutabilityPeriodSinceCreationInDays
	}).(pulumi.IntPtrOutput)
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobContainerImmutabilityPolicyResultOutput{})
}
