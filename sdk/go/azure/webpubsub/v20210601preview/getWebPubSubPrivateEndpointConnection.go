


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSubPrivateEndpointConnection(ctx *pulumi.Context, args *LookupWebPubSubPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubPrivateEndpointConnectionResult, error) {
	var rv LookupWebPubSubPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:webpubsub/v20210601preview:getWebPubSubPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubPrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupWebPubSubPrivateEndpointConnectionResult struct {
	Id                                string                                     `pulumi:"id"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                         `pulumi:"systemData"`
	Type                              string                                     `pulumi:"type"`
}

func LookupWebPubSubPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupWebPubSubPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWebPubSubPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebPubSubPrivateEndpointConnectionResult, error) {
			args := v.(LookupWebPubSubPrivateEndpointConnectionArgs)
			r, err := LookupWebPubSubPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupWebPubSubPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebPubSubPrivateEndpointConnectionResultOutput)
}

type LookupWebPubSubPrivateEndpointConnectionOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWebPubSubPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupWebPubSubPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWebPubSubPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebPubSubPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) ToLookupWebPubSubPrivateEndpointConnectionResultOutput() LookupWebPubSubPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) ToLookupWebPubSubPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupWebPubSubPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupWebPubSubPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebPubSubPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebPubSubPrivateEndpointConnectionResultOutput{})
}
