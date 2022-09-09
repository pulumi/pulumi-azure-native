


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageSyncService(ctx *pulumi.Context, args *LookupStorageSyncServiceArgs, opts ...pulumi.InvokeOption) (*LookupStorageSyncServiceResult, error) {
	var rv LookupStorageSyncServiceResult
	err := ctx.Invoke("azure-native:storagesync/v20200301:getStorageSyncService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageSyncServiceArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
}


type LookupStorageSyncServiceResult struct {
	Id                         string                              `pulumi:"id"`
	IncomingTrafficPolicy      *string                             `pulumi:"incomingTrafficPolicy"`
	LastOperationName          string                              `pulumi:"lastOperationName"`
	LastWorkflowId             string                              `pulumi:"lastWorkflowId"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	StorageSyncServiceStatus   int                                 `pulumi:"storageSyncServiceStatus"`
	StorageSyncServiceUid      string                              `pulumi:"storageSyncServiceUid"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}

func LookupStorageSyncServiceOutput(ctx *pulumi.Context, args LookupStorageSyncServiceOutputArgs, opts ...pulumi.InvokeOption) LookupStorageSyncServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageSyncServiceResult, error) {
			args := v.(LookupStorageSyncServiceArgs)
			r, err := LookupStorageSyncService(ctx, &args, opts...)
			var s LookupStorageSyncServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageSyncServiceResultOutput)
}

type LookupStorageSyncServiceOutputArgs struct {
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
}

func (LookupStorageSyncServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageSyncServiceArgs)(nil)).Elem()
}


type LookupStorageSyncServiceResultOutput struct{ *pulumi.OutputState }

func (LookupStorageSyncServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageSyncServiceResult)(nil)).Elem()
}

func (o LookupStorageSyncServiceResultOutput) ToLookupStorageSyncServiceResultOutput() LookupStorageSyncServiceResultOutput {
	return o
}

func (o LookupStorageSyncServiceResultOutput) ToLookupStorageSyncServiceResultOutputWithContext(ctx context.Context) LookupStorageSyncServiceResultOutput {
	return o
}

func (o LookupStorageSyncServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageSyncServiceResultOutput) IncomingTrafficPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) *string { return v.IncomingTrafficPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupStorageSyncServiceResultOutput) LastOperationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.LastOperationName }).(pulumi.StringOutput)
}

func (o LookupStorageSyncServiceResultOutput) LastWorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.LastWorkflowId }).(pulumi.StringOutput)
}

func (o LookupStorageSyncServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStorageSyncServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageSyncServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupStorageSyncServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupStorageSyncServiceResultOutput) StorageSyncServiceStatus() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) int { return v.StorageSyncServiceStatus }).(pulumi.IntOutput)
}

func (o LookupStorageSyncServiceResultOutput) StorageSyncServiceUid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.StorageSyncServiceUid }).(pulumi.StringOutput)
}

func (o LookupStorageSyncServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStorageSyncServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageSyncServiceResultOutput{})
}
