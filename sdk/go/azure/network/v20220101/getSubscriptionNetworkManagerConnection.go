


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubscriptionNetworkManagerConnection(ctx *pulumi.Context, args *LookupSubscriptionNetworkManagerConnectionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionNetworkManagerConnectionResult, error) {
	var rv LookupSubscriptionNetworkManagerConnectionResult
	err := ctx.Invoke("azure-native:network/v20220101:getSubscriptionNetworkManagerConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionNetworkManagerConnectionArgs struct {
	NetworkManagerConnectionName string `pulumi:"networkManagerConnectionName"`
}


type LookupSubscriptionNetworkManagerConnectionResult struct {
	Description      *string            `pulumi:"description"`
	Etag             string             `pulumi:"etag"`
	Id               string             `pulumi:"id"`
	Name             string             `pulumi:"name"`
	NetworkManagerId *string            `pulumi:"networkManagerId"`
	SystemData       SystemDataResponse `pulumi:"systemData"`
	Type             string             `pulumi:"type"`
}

func LookupSubscriptionNetworkManagerConnectionOutput(ctx *pulumi.Context, args LookupSubscriptionNetworkManagerConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionNetworkManagerConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionNetworkManagerConnectionResult, error) {
			args := v.(LookupSubscriptionNetworkManagerConnectionArgs)
			r, err := LookupSubscriptionNetworkManagerConnection(ctx, &args, opts...)
			var s LookupSubscriptionNetworkManagerConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionNetworkManagerConnectionResultOutput)
}

type LookupSubscriptionNetworkManagerConnectionOutputArgs struct {
	NetworkManagerConnectionName pulumi.StringInput `pulumi:"networkManagerConnectionName"`
}

func (LookupSubscriptionNetworkManagerConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionNetworkManagerConnectionArgs)(nil)).Elem()
}


type LookupSubscriptionNetworkManagerConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionNetworkManagerConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionNetworkManagerConnectionResult)(nil)).Elem()
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) ToLookupSubscriptionNetworkManagerConnectionResultOutput() LookupSubscriptionNetworkManagerConnectionResultOutput {
	return o
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) ToLookupSubscriptionNetworkManagerConnectionResultOutputWithContext(ctx context.Context) LookupSubscriptionNetworkManagerConnectionResultOutput {
	return o
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) NetworkManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) *string { return v.NetworkManagerId }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSubscriptionNetworkManagerConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionNetworkManagerConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionNetworkManagerConnectionResultOutput{})
}
