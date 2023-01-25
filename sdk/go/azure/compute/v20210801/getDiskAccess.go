


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiskAccess(ctx *pulumi.Context, args *LookupDiskAccessArgs, opts ...pulumi.InvokeOption) (*LookupDiskAccessResult, error) {
	var rv LookupDiskAccessResult
	err := ctx.Invoke("azure-native:compute/v20210801:getDiskAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskAccessArgs struct {
	DiskAccessName    string `pulumi:"diskAccessName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDiskAccessResult struct {
	ExtendedLocation           *ExtendedLocationResponse           `pulumi:"extendedLocation"`
	Id                         string                              `pulumi:"id"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	Tags                       map[string]string                   `pulumi:"tags"`
	TimeCreated                string                              `pulumi:"timeCreated"`
	Type                       string                              `pulumi:"type"`
}

func LookupDiskAccessOutput(ctx *pulumi.Context, args LookupDiskAccessOutputArgs, opts ...pulumi.InvokeOption) LookupDiskAccessResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskAccessResult, error) {
			args := v.(LookupDiskAccessArgs)
			r, err := LookupDiskAccess(ctx, &args, opts...)
			var s LookupDiskAccessResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskAccessResultOutput)
}

type LookupDiskAccessOutputArgs struct {
	DiskAccessName    pulumi.StringInput `pulumi:"diskAccessName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskAccessArgs)(nil)).Elem()
}


type LookupDiskAccessResultOutput struct{ *pulumi.OutputState }

func (LookupDiskAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskAccessResult)(nil)).Elem()
}

func (o LookupDiskAccessResultOutput) ToLookupDiskAccessResultOutput() LookupDiskAccessResultOutput {
	return o
}

func (o LookupDiskAccessResultOutput) ToLookupDiskAccessResultOutputWithContext(ctx context.Context) LookupDiskAccessResultOutput {
	return o
}

func (o LookupDiskAccessResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupDiskAccessResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiskAccessResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDiskAccessResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiskAccessResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupDiskAccessResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDiskAccessResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDiskAccessResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupDiskAccessResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskAccessResultOutput{})
}
