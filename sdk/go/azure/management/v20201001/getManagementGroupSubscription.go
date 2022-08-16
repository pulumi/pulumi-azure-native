


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementGroupSubscription(ctx *pulumi.Context, args *LookupManagementGroupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupSubscriptionResult, error) {
	var rv LookupManagementGroupSubscriptionResult
	err := ctx.Invoke("azure-native:management/v20201001:getManagementGroupSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupSubscriptionArgs struct {
	GroupId        string  `pulumi:"groupId"`
	SubscriptionId *string `pulumi:"subscriptionId"`
}


type LookupManagementGroupSubscriptionResult struct {
	DisplayName *string                     `pulumi:"displayName"`
	Id          string                      `pulumi:"id"`
	Name        string                      `pulumi:"name"`
	Parent      *ParentGroupBagInfoResponse `pulumi:"parent"`
	State       *string                     `pulumi:"state"`
	Tenant      *string                     `pulumi:"tenant"`
	Type        string                      `pulumi:"type"`
}

func LookupManagementGroupSubscriptionOutput(ctx *pulumi.Context, args LookupManagementGroupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupSubscriptionResult, error) {
			args := v.(LookupManagementGroupSubscriptionArgs)
			r, err := LookupManagementGroupSubscription(ctx, &args, opts...)
			var s LookupManagementGroupSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementGroupSubscriptionResultOutput)
}

type LookupManagementGroupSubscriptionOutputArgs struct {
	GroupId        pulumi.StringInput    `pulumi:"groupId"`
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupManagementGroupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupSubscriptionArgs)(nil)).Elem()
}


type LookupManagementGroupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupSubscriptionResult)(nil)).Elem()
}

func (o LookupManagementGroupSubscriptionResultOutput) ToLookupManagementGroupSubscriptionResultOutput() LookupManagementGroupSubscriptionResultOutput {
	return o
}

func (o LookupManagementGroupSubscriptionResultOutput) ToLookupManagementGroupSubscriptionResultOutputWithContext(ctx context.Context) LookupManagementGroupSubscriptionResultOutput {
	return o
}

func (o LookupManagementGroupSubscriptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementGroupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementGroupSubscriptionResultOutput) Parent() ParentGroupBagInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) *ParentGroupBagInfoResponse { return v.Parent }).(ParentGroupBagInfoResponsePtrOutput)
}

func (o LookupManagementGroupSubscriptionResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupSubscriptionResultOutput) Tenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) *string { return v.Tenant }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupSubscriptionResultOutput{})
}
