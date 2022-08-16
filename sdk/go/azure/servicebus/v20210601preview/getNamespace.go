


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:servicebus/v20210601preview:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNamespaceArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceResult struct {
	CreatedAt                  string                              `pulumi:"createdAt"`
	DisableLocalAuth           *bool                               `pulumi:"disableLocalAuth"`
	Encryption                 *EncryptionResponse                 `pulumi:"encryption"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityResponse                   `pulumi:"identity"`
	Location                   string                              `pulumi:"location"`
	MetricId                   string                              `pulumi:"metricId"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	ServiceBusEndpoint         string                              `pulumi:"serviceBusEndpoint"`
	Sku                        *SBSkuResponse                      `pulumi:"sku"`
	Status                     string                              `pulumi:"status"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	UpdatedAt                  string                              `pulumi:"updatedAt"`
	ZoneRedundant              *bool                               `pulumi:"zoneRedundant"`
}


func (val *LookupNamespaceResult) Defaults() *LookupNamespaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = tmp.Encryption.Defaults()

	return &tmp
}

func LookupNamespaceOutput(ctx *pulumi.Context, args LookupNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceResult, error) {
			args := v.(LookupNamespaceArgs)
			r, err := LookupNamespace(ctx, &args, opts...)
			var s LookupNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceResultOutput)
}

type LookupNamespaceOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceArgs)(nil)).Elem()
}


type LookupNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceResult)(nil)).Elem()
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutput() LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutputWithContext(ctx context.Context) LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupNamespaceResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

func (o LookupNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupNamespaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) MetricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.MetricId }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupNamespaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupNamespaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) ServiceBusEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.ServiceBusEndpoint }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) Sku() SBSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *SBSkuResponse { return v.Sku }).(SBSkuResponsePtrOutput)
}

func (o LookupNamespaceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNamespaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *bool { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceResultOutput{})
}
