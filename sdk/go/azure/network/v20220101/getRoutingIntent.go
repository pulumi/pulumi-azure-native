


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoutingIntent(ctx *pulumi.Context, args *LookupRoutingIntentArgs, opts ...pulumi.InvokeOption) (*LookupRoutingIntentResult, error) {
	var rv LookupRoutingIntentResult
	err := ctx.Invoke("azure-native:network/v20220101:getRoutingIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoutingIntentArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoutingIntentName string `pulumi:"routingIntentName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupRoutingIntentResult struct {
	Etag              string                  `pulumi:"etag"`
	Id                *string                 `pulumi:"id"`
	Name              *string                 `pulumi:"name"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	RoutingPolicies   []RoutingPolicyResponse `pulumi:"routingPolicies"`
	Type              string                  `pulumi:"type"`
}

func LookupRoutingIntentOutput(ctx *pulumi.Context, args LookupRoutingIntentOutputArgs, opts ...pulumi.InvokeOption) LookupRoutingIntentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoutingIntentResult, error) {
			args := v.(LookupRoutingIntentArgs)
			r, err := LookupRoutingIntent(ctx, &args, opts...)
			var s LookupRoutingIntentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoutingIntentResultOutput)
}

type LookupRoutingIntentOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RoutingIntentName pulumi.StringInput `pulumi:"routingIntentName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupRoutingIntentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutingIntentArgs)(nil)).Elem()
}


type LookupRoutingIntentResultOutput struct{ *pulumi.OutputState }

func (LookupRoutingIntentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutingIntentResult)(nil)).Elem()
}

func (o LookupRoutingIntentResultOutput) ToLookupRoutingIntentResultOutput() LookupRoutingIntentResultOutput {
	return o
}

func (o LookupRoutingIntentResultOutput) ToLookupRoutingIntentResultOutputWithContext(ctx context.Context) LookupRoutingIntentResultOutput {
	return o
}

func (o LookupRoutingIntentResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupRoutingIntentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRoutingIntentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRoutingIntentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRoutingIntentResultOutput) RoutingPolicies() RoutingPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) []RoutingPolicyResponse { return v.RoutingPolicies }).(RoutingPolicyResponseArrayOutput)
}

func (o LookupRoutingIntentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoutingIntentResultOutput{})
}
