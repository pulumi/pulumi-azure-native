


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGlobalReachConnection(ctx *pulumi.Context, args *LookupGlobalReachConnectionArgs, opts ...pulumi.InvokeOption) (*LookupGlobalReachConnectionResult, error) {
	var rv LookupGlobalReachConnectionResult
	err := ctx.Invoke("azure-native:avs/v20220501:getGlobalReachConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalReachConnectionArgs struct {
	GlobalReachConnectionName string `pulumi:"globalReachConnectionName"`
	PrivateCloudName          string `pulumi:"privateCloudName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupGlobalReachConnectionResult struct {
	AddressPrefix           string  `pulumi:"addressPrefix"`
	AuthorizationKey        *string `pulumi:"authorizationKey"`
	CircuitConnectionStatus string  `pulumi:"circuitConnectionStatus"`
	ExpressRouteId          *string `pulumi:"expressRouteId"`
	Id                      string  `pulumi:"id"`
	Name                    string  `pulumi:"name"`
	PeerExpressRouteCircuit *string `pulumi:"peerExpressRouteCircuit"`
	ProvisioningState       string  `pulumi:"provisioningState"`
	Type                    string  `pulumi:"type"`
}

func LookupGlobalReachConnectionOutput(ctx *pulumi.Context, args LookupGlobalReachConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalReachConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalReachConnectionResult, error) {
			args := v.(LookupGlobalReachConnectionArgs)
			r, err := LookupGlobalReachConnection(ctx, &args, opts...)
			var s LookupGlobalReachConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalReachConnectionResultOutput)
}

type LookupGlobalReachConnectionOutputArgs struct {
	GlobalReachConnectionName pulumi.StringInput `pulumi:"globalReachConnectionName"`
	PrivateCloudName          pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGlobalReachConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalReachConnectionArgs)(nil)).Elem()
}


type LookupGlobalReachConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalReachConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalReachConnectionResult)(nil)).Elem()
}

func (o LookupGlobalReachConnectionResultOutput) ToLookupGlobalReachConnectionResultOutput() LookupGlobalReachConnectionResultOutput {
	return o
}

func (o LookupGlobalReachConnectionResultOutput) ToLookupGlobalReachConnectionResultOutputWithContext(ctx context.Context) LookupGlobalReachConnectionResultOutput {
	return o
}

func (o LookupGlobalReachConnectionResultOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

func (o LookupGlobalReachConnectionResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalReachConnectionResultOutput) CircuitConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.CircuitConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupGlobalReachConnectionResultOutput) ExpressRouteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) *string { return v.ExpressRouteId }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalReachConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGlobalReachConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGlobalReachConnectionResultOutput) PeerExpressRouteCircuit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) *string { return v.PeerExpressRouteCircuit }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalReachConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGlobalReachConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalReachConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalReachConnectionResultOutput{})
}
