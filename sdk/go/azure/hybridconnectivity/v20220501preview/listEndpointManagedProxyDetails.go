


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEndpointManagedProxyDetails(ctx *pulumi.Context, args *ListEndpointManagedProxyDetailsArgs, opts ...pulumi.InvokeOption) (*ListEndpointManagedProxyDetailsResult, error) {
	var rv ListEndpointManagedProxyDetailsResult
	err := ctx.Invoke("azure-native:hybridconnectivity/v20220501preview:listEndpointManagedProxyDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEndpointManagedProxyDetailsArgs struct {
	EndpointName string  `pulumi:"endpointName"`
	Hostname     *string `pulumi:"hostname"`
	ResourceUri  string  `pulumi:"resourceUri"`
	Service      string  `pulumi:"service"`
}


type ListEndpointManagedProxyDetailsResult struct {
	ExpiresOn float64 `pulumi:"expiresOn"`
	Proxy     string  `pulumi:"proxy"`
}

func ListEndpointManagedProxyDetailsOutput(ctx *pulumi.Context, args ListEndpointManagedProxyDetailsOutputArgs, opts ...pulumi.InvokeOption) ListEndpointManagedProxyDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEndpointManagedProxyDetailsResult, error) {
			args := v.(ListEndpointManagedProxyDetailsArgs)
			r, err := ListEndpointManagedProxyDetails(ctx, &args, opts...)
			var s ListEndpointManagedProxyDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEndpointManagedProxyDetailsResultOutput)
}

type ListEndpointManagedProxyDetailsOutputArgs struct {
	EndpointName pulumi.StringInput    `pulumi:"endpointName"`
	Hostname     pulumi.StringPtrInput `pulumi:"hostname"`
	ResourceUri  pulumi.StringInput    `pulumi:"resourceUri"`
	Service      pulumi.StringInput    `pulumi:"service"`
}

func (ListEndpointManagedProxyDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEndpointManagedProxyDetailsArgs)(nil)).Elem()
}


type ListEndpointManagedProxyDetailsResultOutput struct{ *pulumi.OutputState }

func (ListEndpointManagedProxyDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEndpointManagedProxyDetailsResult)(nil)).Elem()
}

func (o ListEndpointManagedProxyDetailsResultOutput) ToListEndpointManagedProxyDetailsResultOutput() ListEndpointManagedProxyDetailsResultOutput {
	return o
}

func (o ListEndpointManagedProxyDetailsResultOutput) ToListEndpointManagedProxyDetailsResultOutputWithContext(ctx context.Context) ListEndpointManagedProxyDetailsResultOutput {
	return o
}

func (o ListEndpointManagedProxyDetailsResultOutput) ExpiresOn() pulumi.Float64Output {
	return o.ApplyT(func(v ListEndpointManagedProxyDetailsResult) float64 { return v.ExpiresOn }).(pulumi.Float64Output)
}

func (o ListEndpointManagedProxyDetailsResultOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointManagedProxyDetailsResult) string { return v.Proxy }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEndpointManagedProxyDetailsResultOutput{})
}
