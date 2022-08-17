


package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:offazure:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionArgs struct {
	PeConnectionName  string `pulumi:"peConnectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SiteName          string `pulumi:"siteName"`
}


type LookupPrivateEndpointConnectionResult struct {
	ETag       string                                      `pulumi:"eTag"`
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Type       string                                      `pulumi:"type"`
}

func LookupPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionResult, error) {
			args := v.(LookupPrivateEndpointConnectionArgs)
			r, err := LookupPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionResultOutput)
}

type LookupPrivateEndpointConnectionOutputArgs struct {
	PeConnectionName  pulumi.StringInput `pulumi:"peConnectionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SiteName          pulumi.StringInput `pulumi:"siteName"`
}

func (LookupPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionResultOutput() LookupPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionResultOutput{})
}
