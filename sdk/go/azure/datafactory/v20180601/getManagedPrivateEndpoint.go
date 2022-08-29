


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedPrivateEndpoint(ctx *pulumi.Context, args *LookupManagedPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupManagedPrivateEndpointResult, error) {
	var rv LookupManagedPrivateEndpointResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getManagedPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedPrivateEndpointArgs struct {
	FactoryName                string `pulumi:"factoryName"`
	ManagedPrivateEndpointName string `pulumi:"managedPrivateEndpointName"`
	ManagedVirtualNetworkName  string `pulumi:"managedVirtualNetworkName"`
	ResourceGroupName          string `pulumi:"resourceGroupName"`
}


type LookupManagedPrivateEndpointResult struct {
	Etag       string                         `pulumi:"etag"`
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties ManagedPrivateEndpointResponse `pulumi:"properties"`
	Type       string                         `pulumi:"type"`
}

func LookupManagedPrivateEndpointOutput(ctx *pulumi.Context, args LookupManagedPrivateEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupManagedPrivateEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedPrivateEndpointResult, error) {
			args := v.(LookupManagedPrivateEndpointArgs)
			r, err := LookupManagedPrivateEndpoint(ctx, &args, opts...)
			var s LookupManagedPrivateEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedPrivateEndpointResultOutput)
}

type LookupManagedPrivateEndpointOutputArgs struct {
	FactoryName                pulumi.StringInput `pulumi:"factoryName"`
	ManagedPrivateEndpointName pulumi.StringInput `pulumi:"managedPrivateEndpointName"`
	ManagedVirtualNetworkName  pulumi.StringInput `pulumi:"managedVirtualNetworkName"`
	ResourceGroupName          pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedPrivateEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPrivateEndpointArgs)(nil)).Elem()
}


type LookupManagedPrivateEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupManagedPrivateEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPrivateEndpointResult)(nil)).Elem()
}

func (o LookupManagedPrivateEndpointResultOutput) ToLookupManagedPrivateEndpointResultOutput() LookupManagedPrivateEndpointResultOutput {
	return o
}

func (o LookupManagedPrivateEndpointResultOutput) ToLookupManagedPrivateEndpointResultOutputWithContext(ctx context.Context) LookupManagedPrivateEndpointResultOutput {
	return o
}

func (o LookupManagedPrivateEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) Properties() ManagedPrivateEndpointResponseOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) ManagedPrivateEndpointResponse { return v.Properties }).(ManagedPrivateEndpointResponseOutput)
}

func (o LookupManagedPrivateEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedPrivateEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedPrivateEndpointResultOutput{})
}
