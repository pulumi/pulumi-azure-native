


package devices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotDpsResourcePrivateEndpointConnection(ctx *pulumi.Context, args *LookupIotDpsResourcePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupIotDpsResourcePrivateEndpointConnectionResult, error) {
	var rv LookupIotDpsResourcePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:devices:getIotDpsResourcePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotDpsResourcePrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupIotDpsResourcePrivateEndpointConnectionResult struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}

func LookupIotDpsResourcePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupIotDpsResourcePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupIotDpsResourcePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotDpsResourcePrivateEndpointConnectionResult, error) {
			args := v.(LookupIotDpsResourcePrivateEndpointConnectionArgs)
			r, err := LookupIotDpsResourcePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupIotDpsResourcePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotDpsResourcePrivateEndpointConnectionResultOutput)
}

type LookupIotDpsResourcePrivateEndpointConnectionOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupIotDpsResourcePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDpsResourcePrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupIotDpsResourcePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupIotDpsResourcePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDpsResourcePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) ToLookupIotDpsResourcePrivateEndpointConnectionResultOutput() LookupIotDpsResourcePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) ToLookupIotDpsResourcePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupIotDpsResourcePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o LookupIotDpsResourcePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourcePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotDpsResourcePrivateEndpointConnectionResultOutput{})
}
