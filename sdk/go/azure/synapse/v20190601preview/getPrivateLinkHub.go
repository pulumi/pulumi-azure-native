


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPrivateLinkHub(ctx *pulumi.Context, args *LookupPrivateLinkHubArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkHubResult, error) {
	var rv LookupPrivateLinkHubResult
	err := ctx.Invoke("azure-native:synapse/v20190601preview:getPrivateLinkHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkHubArgs struct {
	PrivateLinkHubName string `pulumi:"privateLinkHubName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPrivateLinkHubResult struct {
	Id                         string                                                    `pulumi:"id"`
	Location                   string                                                    `pulumi:"location"`
	Name                       string                                                    `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionForPrivateLinkHubBasicResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          *string                                                   `pulumi:"provisioningState"`
	Tags                       map[string]string                                         `pulumi:"tags"`
	Type                       string                                                    `pulumi:"type"`
}

func LookupPrivateLinkHubOutput(ctx *pulumi.Context, args LookupPrivateLinkHubOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkHubResult, error) {
			args := v.(LookupPrivateLinkHubArgs)
			r, err := LookupPrivateLinkHub(ctx, &args, opts...)
			var s LookupPrivateLinkHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkHubResultOutput)
}

type LookupPrivateLinkHubOutputArgs struct {
	PrivateLinkHubName pulumi.StringInput `pulumi:"privateLinkHubName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateLinkHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkHubArgs)(nil)).Elem()
}


type LookupPrivateLinkHubResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkHubResult)(nil)).Elem()
}

func (o LookupPrivateLinkHubResultOutput) ToLookupPrivateLinkHubResultOutput() LookupPrivateLinkHubResultOutput {
	return o
}

func (o LookupPrivateLinkHubResultOutput) ToLookupPrivateLinkHubResultOutputWithContext(ctx context.Context) LookupPrivateLinkHubResultOutput {
	return o
}

func (o LookupPrivateLinkHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkHubResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateLinkHubResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) []PrivateEndpointConnectionForPrivateLinkHubBasicResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput)
}

func (o LookupPrivateLinkHubResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateLinkHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateLinkHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkHubResultOutput{})
}
