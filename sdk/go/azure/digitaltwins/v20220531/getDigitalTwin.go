


package v20220531

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDigitalTwin(ctx *pulumi.Context, args *LookupDigitalTwinArgs, opts ...pulumi.InvokeOption) (*LookupDigitalTwinResult, error) {
	var rv LookupDigitalTwinResult
	err := ctx.Invoke("azure-native:digitaltwins/v20220531:getDigitalTwin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDigitalTwinArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupDigitalTwinResult struct {
	CreatedTime                string                              `pulumi:"createdTime"`
	HostName                   string                              `pulumi:"hostName"`
	Id                         string                              `pulumi:"id"`
	Identity                   *DigitalTwinsIdentityResponse       `pulumi:"identity"`
	LastUpdatedTime            string                              `pulumi:"lastUpdatedTime"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}

func LookupDigitalTwinOutput(ctx *pulumi.Context, args LookupDigitalTwinOutputArgs, opts ...pulumi.InvokeOption) LookupDigitalTwinResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDigitalTwinResult, error) {
			args := v.(LookupDigitalTwinArgs)
			r, err := LookupDigitalTwin(ctx, &args, opts...)
			var s LookupDigitalTwinResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDigitalTwinResultOutput)
}

type LookupDigitalTwinOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupDigitalTwinOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDigitalTwinArgs)(nil)).Elem()
}


type LookupDigitalTwinResultOutput struct{ *pulumi.OutputState }

func (LookupDigitalTwinResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDigitalTwinResult)(nil)).Elem()
}

func (o LookupDigitalTwinResultOutput) ToLookupDigitalTwinResultOutput() LookupDigitalTwinResultOutput {
	return o
}

func (o LookupDigitalTwinResultOutput) ToLookupDigitalTwinResultOutputWithContext(ctx context.Context) LookupDigitalTwinResultOutput {
	return o
}

func (o LookupDigitalTwinResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) Identity() DigitalTwinsIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) *DigitalTwinsIdentityResponse { return v.Identity }).(DigitalTwinsIdentityResponsePtrOutput)
}

func (o LookupDigitalTwinResultOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupDigitalTwinResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDigitalTwinResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupDigitalTwinResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDigitalTwinResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDigitalTwinResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDigitalTwinResultOutput{})
}
