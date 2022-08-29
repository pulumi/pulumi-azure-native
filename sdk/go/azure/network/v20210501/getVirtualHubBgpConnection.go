


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHubBgpConnection(ctx *pulumi.Context, args *LookupVirtualHubBgpConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubBgpConnectionResult, error) {
	var rv LookupVirtualHubBgpConnectionResult
	err := ctx.Invoke("azure-native:network/v20210501:getVirtualHubBgpConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubBgpConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubBgpConnectionResult struct {
	ConnectionState             string               `pulumi:"connectionState"`
	Etag                        string               `pulumi:"etag"`
	HubVirtualNetworkConnection *SubResourceResponse `pulumi:"hubVirtualNetworkConnection"`
	Id                          *string              `pulumi:"id"`
	Name                        *string              `pulumi:"name"`
	PeerAsn                     *float64             `pulumi:"peerAsn"`
	PeerIp                      *string              `pulumi:"peerIp"`
	ProvisioningState           string               `pulumi:"provisioningState"`
	Type                        string               `pulumi:"type"`
}

func LookupVirtualHubBgpConnectionOutput(ctx *pulumi.Context, args LookupVirtualHubBgpConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubBgpConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubBgpConnectionResult, error) {
			args := v.(LookupVirtualHubBgpConnectionArgs)
			r, err := LookupVirtualHubBgpConnection(ctx, &args, opts...)
			var s LookupVirtualHubBgpConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubBgpConnectionResultOutput)
}

type LookupVirtualHubBgpConnectionOutputArgs struct {
	ConnectionName    pulumi.StringInput `pulumi:"connectionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubBgpConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubBgpConnectionArgs)(nil)).Elem()
}


type LookupVirtualHubBgpConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubBgpConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubBgpConnectionResult)(nil)).Elem()
}

func (o LookupVirtualHubBgpConnectionResultOutput) ToLookupVirtualHubBgpConnectionResultOutput() LookupVirtualHubBgpConnectionResultOutput {
	return o
}

func (o LookupVirtualHubBgpConnectionResultOutput) ToLookupVirtualHubBgpConnectionResultOutputWithContext(ctx context.Context) LookupVirtualHubBgpConnectionResultOutput {
	return o
}

func (o LookupVirtualHubBgpConnectionResultOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) string { return v.ConnectionState }).(pulumi.StringOutput)
}

func (o LookupVirtualHubBgpConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualHubBgpConnectionResultOutput) HubVirtualNetworkConnection() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *SubResourceResponse { return v.HubVirtualNetworkConnection }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualHubBgpConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubBgpConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubBgpConnectionResultOutput) PeerAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *float64 { return v.PeerAsn }).(pulumi.Float64PtrOutput)
}

func (o LookupVirtualHubBgpConnectionResultOutput) PeerIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) *string { return v.PeerIp }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubBgpConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualHubBgpConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubBgpConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubBgpConnectionResultOutput{})
}
