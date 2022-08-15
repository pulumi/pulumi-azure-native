


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupExpressRouteCircuitAuthorization(ctx *pulumi.Context, args *LookupExpressRouteCircuitAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitAuthorizationResult, error) {
	var rv LookupExpressRouteCircuitAuthorizationResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getExpressRouteCircuitAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitAuthorizationArgs struct {
	AuthorizationName string `pulumi:"authorizationName"`
	CircuitName       string `pulumi:"circuitName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteCircuitAuthorizationResult struct {
	AuthorizationKey       *string `pulumi:"authorizationKey"`
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	Etag                   *string `pulumi:"etag"`
	Id                     *string `pulumi:"id"`
	Name                   *string `pulumi:"name"`
	ProvisioningState      *string `pulumi:"provisioningState"`
}

func LookupExpressRouteCircuitAuthorizationOutput(ctx *pulumi.Context, args LookupExpressRouteCircuitAuthorizationOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCircuitAuthorizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCircuitAuthorizationResult, error) {
			args := v.(LookupExpressRouteCircuitAuthorizationArgs)
			r, err := LookupExpressRouteCircuitAuthorization(ctx, &args, opts...)
			var s LookupExpressRouteCircuitAuthorizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCircuitAuthorizationResultOutput)
}

type LookupExpressRouteCircuitAuthorizationOutputArgs struct {
	AuthorizationName pulumi.StringInput `pulumi:"authorizationName"`
	CircuitName       pulumi.StringInput `pulumi:"circuitName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCircuitAuthorizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitAuthorizationArgs)(nil)).Elem()
}


type LookupExpressRouteCircuitAuthorizationResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCircuitAuthorizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitAuthorizationResult)(nil)).Elem()
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) ToLookupExpressRouteCircuitAuthorizationResultOutput() LookupExpressRouteCircuitAuthorizationResultOutput {
	return o
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) ToLookupExpressRouteCircuitAuthorizationResultOutputWithContext(ctx context.Context) LookupExpressRouteCircuitAuthorizationResultOutput {
	return o
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) AuthorizationUseStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.AuthorizationUseStatus }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitAuthorizationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitAuthorizationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCircuitAuthorizationResultOutput{})
}
