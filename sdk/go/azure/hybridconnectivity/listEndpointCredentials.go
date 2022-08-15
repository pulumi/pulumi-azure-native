


package hybridconnectivity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEndpointCredentials(ctx *pulumi.Context, args *ListEndpointCredentialsArgs, opts ...pulumi.InvokeOption) (*ListEndpointCredentialsResult, error) {
	var rv ListEndpointCredentialsResult
	err := ctx.Invoke("azure-native:hybridconnectivity:listEndpointCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEndpointCredentialsArgs struct {
	EndpointName string `pulumi:"endpointName"`
	Expiresin    *int   `pulumi:"expiresin"`
	ResourceUri  string `pulumi:"resourceUri"`
}


type ListEndpointCredentialsResult struct {
	AccessKey            string   `pulumi:"accessKey"`
	ExpiresOn            *float64 `pulumi:"expiresOn"`
	HybridConnectionName string   `pulumi:"hybridConnectionName"`
	NamespaceName        string   `pulumi:"namespaceName"`
	NamespaceNameSuffix  string   `pulumi:"namespaceNameSuffix"`
}

func ListEndpointCredentialsOutput(ctx *pulumi.Context, args ListEndpointCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListEndpointCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEndpointCredentialsResult, error) {
			args := v.(ListEndpointCredentialsArgs)
			r, err := ListEndpointCredentials(ctx, &args, opts...)
			var s ListEndpointCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEndpointCredentialsResultOutput)
}

type ListEndpointCredentialsOutputArgs struct {
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	Expiresin    pulumi.IntPtrInput `pulumi:"expiresin"`
	ResourceUri  pulumi.StringInput `pulumi:"resourceUri"`
}

func (ListEndpointCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEndpointCredentialsArgs)(nil)).Elem()
}


type ListEndpointCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListEndpointCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEndpointCredentialsResult)(nil)).Elem()
}

func (o ListEndpointCredentialsResultOutput) ToListEndpointCredentialsResultOutput() ListEndpointCredentialsResultOutput {
	return o
}

func (o ListEndpointCredentialsResultOutput) ToListEndpointCredentialsResultOutputWithContext(ctx context.Context) ListEndpointCredentialsResultOutput {
	return o
}

func (o ListEndpointCredentialsResultOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) string { return v.AccessKey }).(pulumi.StringOutput)
}

func (o ListEndpointCredentialsResultOutput) ExpiresOn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) *float64 { return v.ExpiresOn }).(pulumi.Float64PtrOutput)
}

func (o ListEndpointCredentialsResultOutput) HybridConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) string { return v.HybridConnectionName }).(pulumi.StringOutput)
}

func (o ListEndpointCredentialsResultOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) string { return v.NamespaceName }).(pulumi.StringOutput)
}

func (o ListEndpointCredentialsResultOutput) NamespaceNameSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) string { return v.NamespaceNameSuffix }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEndpointCredentialsResultOutput{})
}
