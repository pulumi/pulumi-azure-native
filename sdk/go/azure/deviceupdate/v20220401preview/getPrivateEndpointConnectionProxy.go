


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionProxy(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionProxyArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionProxyResult, error) {
	var rv LookupPrivateEndpointConnectionProxyResult
	err := ctx.Invoke("azure-native:deviceupdate/v20220401preview:getPrivateEndpointConnectionProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionProxyArgs struct {
	AccountName                      string `pulumi:"accountName"`
	PrivateEndpointConnectionProxyId string `pulumi:"privateEndpointConnectionProxyId"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointConnectionProxyResult struct {
	ETag                  string                         `pulumi:"eTag"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
	RemotePrivateEndpoint *RemotePrivateEndpointResponse `pulumi:"remotePrivateEndpoint"`
	Status                *string                        `pulumi:"status"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Type                  string                         `pulumi:"type"`
}

func LookupPrivateEndpointConnectionProxyOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionProxyOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionProxyResult, error) {
			args := v.(LookupPrivateEndpointConnectionProxyArgs)
			r, err := LookupPrivateEndpointConnectionProxy(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionProxyResultOutput)
}

type LookupPrivateEndpointConnectionProxyOutputArgs struct {
	AccountName                      pulumi.StringInput `pulumi:"accountName"`
	PrivateEndpointConnectionProxyId pulumi.StringInput `pulumi:"privateEndpointConnectionProxyId"`
	ResourceGroupName                pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointConnectionProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionProxyArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionProxyResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionProxyResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) ToLookupPrivateEndpointConnectionProxyResultOutput() LookupPrivateEndpointConnectionProxyResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) ToLookupPrivateEndpointConnectionProxyResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionProxyResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionProxyResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionProxyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) RemotePrivateEndpoint() RemotePrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionProxyResult) *RemotePrivateEndpointResponse {
		return v.RemotePrivateEndpoint
	}).(RemotePrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionProxyResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionProxyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionProxyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionProxyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionProxyResultOutput{})
}
