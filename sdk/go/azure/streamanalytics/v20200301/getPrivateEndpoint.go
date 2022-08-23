


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azure-native:streamanalytics/v20200301:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointArgs struct {
	ClusterName         string `pulumi:"clusterName"`
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointResult struct {
	CreatedDate                         string                                 `pulumi:"createdDate"`
	Etag                                string                                 `pulumi:"etag"`
	Id                                  string                                 `pulumi:"id"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"manualPrivateLinkServiceConnections"`
	Name                                string                                 `pulumi:"name"`
	Type                                string                                 `pulumi:"type"`
}

func LookupPrivateEndpointOutput(ctx *pulumi.Context, args LookupPrivateEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointResult, error) {
			args := v.(LookupPrivateEndpointArgs)
			r, err := LookupPrivateEndpoint(ctx, &args, opts...)
			var s LookupPrivateEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointResultOutput)
}

type LookupPrivateEndpointOutputArgs struct {
	ClusterName         pulumi.StringInput `pulumi:"clusterName"`
	PrivateEndpointName pulumi.StringInput `pulumi:"privateEndpointName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointArgs)(nil)).Elem()
}


type LookupPrivateEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointResult)(nil)).Elem()
}

func (o LookupPrivateEndpointResultOutput) ToLookupPrivateEndpointResultOutput() LookupPrivateEndpointResultOutput {
	return o
}

func (o LookupPrivateEndpointResultOutput) ToLookupPrivateEndpointResultOutputWithContext(ctx context.Context) LookupPrivateEndpointResultOutput {
	return o
}

func (o LookupPrivateEndpointResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateLinkServiceConnectionResponse {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointResultOutput{})
}
