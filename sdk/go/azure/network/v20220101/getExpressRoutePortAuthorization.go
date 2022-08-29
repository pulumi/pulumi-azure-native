


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRoutePortAuthorization(ctx *pulumi.Context, args *LookupExpressRoutePortAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupExpressRoutePortAuthorizationResult, error) {
	var rv LookupExpressRoutePortAuthorizationResult
	err := ctx.Invoke("azure-native:network/v20220101:getExpressRoutePortAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRoutePortAuthorizationArgs struct {
	AuthorizationName    string `pulumi:"authorizationName"`
	ExpressRoutePortName string `pulumi:"expressRoutePortName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupExpressRoutePortAuthorizationResult struct {
	AuthorizationKey       string  `pulumi:"authorizationKey"`
	AuthorizationUseStatus string  `pulumi:"authorizationUseStatus"`
	CircuitResourceUri     string  `pulumi:"circuitResourceUri"`
	Etag                   string  `pulumi:"etag"`
	Id                     *string `pulumi:"id"`
	Name                   *string `pulumi:"name"`
	ProvisioningState      string  `pulumi:"provisioningState"`
	Type                   string  `pulumi:"type"`
}

func LookupExpressRoutePortAuthorizationOutput(ctx *pulumi.Context, args LookupExpressRoutePortAuthorizationOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRoutePortAuthorizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRoutePortAuthorizationResult, error) {
			args := v.(LookupExpressRoutePortAuthorizationArgs)
			r, err := LookupExpressRoutePortAuthorization(ctx, &args, opts...)
			var s LookupExpressRoutePortAuthorizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRoutePortAuthorizationResultOutput)
}

type LookupExpressRoutePortAuthorizationOutputArgs struct {
	AuthorizationName    pulumi.StringInput `pulumi:"authorizationName"`
	ExpressRoutePortName pulumi.StringInput `pulumi:"expressRoutePortName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRoutePortAuthorizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRoutePortAuthorizationArgs)(nil)).Elem()
}


type LookupExpressRoutePortAuthorizationResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRoutePortAuthorizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRoutePortAuthorizationResult)(nil)).Elem()
}

func (o LookupExpressRoutePortAuthorizationResultOutput) ToLookupExpressRoutePortAuthorizationResultOutput() LookupExpressRoutePortAuthorizationResultOutput {
	return o
}

func (o LookupExpressRoutePortAuthorizationResultOutput) ToLookupExpressRoutePortAuthorizationResultOutputWithContext(ctx context.Context) LookupExpressRoutePortAuthorizationResultOutput {
	return o
}

func (o LookupExpressRoutePortAuthorizationResultOutput) AuthorizationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.AuthorizationKey }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortAuthorizationResultOutput) AuthorizationUseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.AuthorizationUseStatus }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortAuthorizationResultOutput) CircuitResourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.CircuitResourceUri }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortAuthorizationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortAuthorizationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRoutePortAuthorizationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRoutePortAuthorizationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortAuthorizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortAuthorizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRoutePortAuthorizationResultOutput{})
}
