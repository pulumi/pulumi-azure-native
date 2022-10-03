


package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstancePrivateEndpointConnection(ctx *pulumi.Context, args *LookupManagedInstancePrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstancePrivateEndpointConnectionResult, error) {
	var rv LookupManagedInstancePrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:sql:getManagedInstancePrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstancePrivateEndpointConnectionArgs struct {
	ManagedInstanceName           string `pulumi:"managedInstanceName"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupManagedInstancePrivateEndpointConnectionResult struct {
	Id                                string                                                            `pulumi:"id"`
	Name                              string                                                            `pulumi:"name"`
	PrivateEndpoint                   *ManagedInstancePrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                                            `pulumi:"provisioningState"`
	Type                              string                                                            `pulumi:"type"`
}

func LookupManagedInstancePrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupManagedInstancePrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstancePrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstancePrivateEndpointConnectionResult, error) {
			args := v.(LookupManagedInstancePrivateEndpointConnectionArgs)
			r, err := LookupManagedInstancePrivateEndpointConnection(ctx, &args, opts...)
			var s LookupManagedInstancePrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstancePrivateEndpointConnectionResultOutput)
}

type LookupManagedInstancePrivateEndpointConnectionOutputArgs struct {
	ManagedInstanceName           pulumi.StringInput `pulumi:"managedInstanceName"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstancePrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstancePrivateEndpointConnectionArgs)(nil)).Elem()
}


type LookupManagedInstancePrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstancePrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstancePrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) ToLookupManagedInstancePrivateEndpointConnectionResultOutput() LookupManagedInstancePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) ToLookupManagedInstancePrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupManagedInstancePrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) PrivateEndpoint() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) *ManagedInstancePrivateEndpointPropertyResponse {
		return v.PrivateEndpoint
	}).(ManagedInstancePrivateEndpointPropertyResponsePtrOutput)
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) *ManagedInstancePrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupManagedInstancePrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstancePrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstancePrivateEndpointConnectionResultOutput{})
}
