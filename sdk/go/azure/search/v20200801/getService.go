


package v20200801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:search/v20200801:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SearchServiceName string `pulumi:"searchServiceName"`
}


type LookupServiceResult struct {
	HostingMode                *string                             `pulumi:"hostingMode"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityResponse                   `pulumi:"identity"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	NetworkRuleSet             *NetworkRuleSetResponse             `pulumi:"networkRuleSet"`
	PartitionCount             *int                                `pulumi:"partitionCount"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	ReplicaCount               *int                                `pulumi:"replicaCount"`
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	Sku                        *SkuResponse                        `pulumi:"sku"`
	Status                     string                              `pulumi:"status"`
	StatusDetails              string                              `pulumi:"statusDetails"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}


func (val *LookupServiceResult) Defaults() *LookupServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HostingMode) {
		hostingMode_ := "default"
		tmp.HostingMode = &hostingMode_
	}
	if isZero(tmp.PartitionCount) {
		partitionCount_ := 1
		tmp.PartitionCount = &partitionCount_
	}
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	if isZero(tmp.ReplicaCount) {
		replicaCount_ := 1
		tmp.ReplicaCount = &replicaCount_
	}
	return &tmp
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SearchServiceName pulumi.StringInput `pulumi:"searchServiceName"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}


type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) HostingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.HostingMode }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) NetworkRuleSet() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *NetworkRuleSetResponse { return v.NetworkRuleSet }).(NetworkRuleSetResponsePtrOutput)
}

func (o LookupServiceResultOutput) PartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *int { return v.PartitionCount }).(pulumi.IntPtrOutput)
}

func (o LookupServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

func (o LookupServiceResultOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []SharedPrivateLinkResourceResponse { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

func (o LookupServiceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupServiceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.StatusDetails }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
