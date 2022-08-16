


package v20170418

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20170418:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionArgs struct {
	AccountName                   string `pulumi:"accountName"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointConnectionResult struct {
	Etag       string                                      `pulumi:"etag"`
	Id         string                                      `pulumi:"id"`
	Location   *string                                     `pulumi:"location"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
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
	AccountName                   pulumi.StringInput `pulumi:"accountName"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupPrivateEndpointConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionResultOutput{})
}
