


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMHSMPrivateEndpointConnection(ctx *pulumi.Context, args *LookupMHSMPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupMHSMPrivateEndpointConnectionResult, error) {
	var rv LookupMHSMPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:keyvault/v20210601preview:getMHSMPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMHSMPrivateEndpointConnectionArgs struct {
	Name                          string `pulumi:"name"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupMHSMPrivateEndpointConnectionResult struct {
	Etag                              *string                                        `pulumi:"etag"`
	Id                                string                                         `pulumi:"id"`
	Location                          *string                                        `pulumi:"location"`
	Name                              string                                         `pulumi:"name"`
	PrivateEndpoint                   *MHSMPrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *MHSMPrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                         `pulumi:"provisioningState"`
	Sku                               *ManagedHsmSkuResponse                         `pulumi:"sku"`
	SystemData                        SystemDataResponse                             `pulumi:"systemData"`
	Tags                              map[string]string                              `pulumi:"tags"`
	Type                              string                                         `pulumi:"type"`
}

func LookupMHSMPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupMHSMPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupMHSMPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMHSMPrivateEndpointConnectionResult, error) {
			args := v.(LookupMHSMPrivateEndpointConnectionArgs)
			r, err := LookupMHSMPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupMHSMPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMHSMPrivateEndpointConnectionResultOutput)
}

type LookupMHSMPrivateEndpointConnectionOutputArgs struct {
	Name                          pulumi.StringInput `pulumi:"name"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMHSMPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMHSMPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupMHSMPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupMHSMPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMHSMPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) ToLookupMHSMPrivateEndpointConnectionResultOutput() LookupMHSMPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) ToLookupMHSMPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupMHSMPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) PrivateEndpoint() MHSMPrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) *MHSMPrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(MHSMPrivateEndpointResponsePtrOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) *MHSMPrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(MHSMPrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) Sku() ManagedHsmSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) *ManagedHsmSkuResponse { return v.Sku }).(ManagedHsmSkuResponsePtrOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMHSMPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMHSMPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMHSMPrivateEndpointConnectionResultOutput{})
}
