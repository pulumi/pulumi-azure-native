


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProduct(ctx *pulumi.Context, args *LookupProductArgs, opts ...pulumi.InvokeOption) (*LookupProductResult, error) {
	var rv LookupProductResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductArgs struct {
	ProductId         string `pulumi:"productId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupProductResult struct {
	ApprovalRequired     *bool   `pulumi:"approvalRequired"`
	Description          *string `pulumi:"description"`
	DisplayName          string  `pulumi:"displayName"`
	Id                   string  `pulumi:"id"`
	Name                 string  `pulumi:"name"`
	State                *string `pulumi:"state"`
	SubscriptionRequired *bool   `pulumi:"subscriptionRequired"`
	SubscriptionsLimit   *int    `pulumi:"subscriptionsLimit"`
	Terms                *string `pulumi:"terms"`
	Type                 string  `pulumi:"type"`
}

func LookupProductOutput(ctx *pulumi.Context, args LookupProductOutputArgs, opts ...pulumi.InvokeOption) LookupProductResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProductResult, error) {
			args := v.(LookupProductArgs)
			r, err := LookupProduct(ctx, &args, opts...)
			var s LookupProductResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProductResultOutput)
}

type LookupProductOutputArgs struct {
	ProductId         pulumi.StringInput `pulumi:"productId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductArgs)(nil)).Elem()
}


type LookupProductResultOutput struct{ *pulumi.OutputState }

func (LookupProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductResult)(nil)).Elem()
}

func (o LookupProductResultOutput) ToLookupProductResultOutput() LookupProductResultOutput {
	return o
}

func (o LookupProductResultOutput) ToLookupProductResultOutputWithContext(ctx context.Context) LookupProductResultOutput {
	return o
}

func (o LookupProductResultOutput) ApprovalRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProductResult) *bool { return v.ApprovalRequired }).(pulumi.BoolPtrOutput)
}

func (o LookupProductResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProductResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupProductResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProductResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProductResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProductResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupProductResultOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProductResult) *bool { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

func (o LookupProductResultOutput) SubscriptionsLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupProductResult) *int { return v.SubscriptionsLimit }).(pulumi.IntPtrOutput)
}

func (o LookupProductResultOutput) Terms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProductResult) *string { return v.Terms }).(pulumi.StringPtrOutput)
}

func (o LookupProductResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProductResultOutput{})
}
