


package v20200717preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAuthorization(ctx *pulumi.Context, args *LookupAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationResult, error) {
	var rv LookupAuthorizationResult
	err := ctx.Invoke("azure-native:avs/v20200717preview:getAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationArgs struct {
	AuthorizationName string `pulumi:"authorizationName"`
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAuthorizationResult struct {
	ExpressRouteAuthorizationId  string `pulumi:"expressRouteAuthorizationId"`
	ExpressRouteAuthorizationKey string `pulumi:"expressRouteAuthorizationKey"`
	Id                           string `pulumi:"id"`
	Name                         string `pulumi:"name"`
	ProvisioningState            string `pulumi:"provisioningState"`
	Type                         string `pulumi:"type"`
}

func LookupAuthorizationOutput(ctx *pulumi.Context, args LookupAuthorizationOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationResult, error) {
			args := v.(LookupAuthorizationArgs)
			r, err := LookupAuthorization(ctx, &args, opts...)
			var s LookupAuthorizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationResultOutput)
}

type LookupAuthorizationOutputArgs struct {
	AuthorizationName pulumi.StringInput `pulumi:"authorizationName"`
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAuthorizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationArgs)(nil)).Elem()
}


type LookupAuthorizationResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationResult)(nil)).Elem()
}

func (o LookupAuthorizationResultOutput) ToLookupAuthorizationResultOutput() LookupAuthorizationResultOutput {
	return o
}

func (o LookupAuthorizationResultOutput) ToLookupAuthorizationResultOutputWithContext(ctx context.Context) LookupAuthorizationResultOutput {
	return o
}

func (o LookupAuthorizationResultOutput) ExpressRouteAuthorizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.ExpressRouteAuthorizationId }).(pulumi.StringOutput)
}

func (o LookupAuthorizationResultOutput) ExpressRouteAuthorizationKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.ExpressRouteAuthorizationKey }).(pulumi.StringOutput)
}

func (o LookupAuthorizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAuthorizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAuthorizationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAuthorizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationResultOutput{})
}
