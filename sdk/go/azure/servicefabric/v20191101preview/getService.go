


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:servicefabric/v20191101preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	ApplicationName   string `pulumi:"applicationName"`
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupServiceResult struct {
	CorrelationScheme            []ServiceCorrelationDescriptionResponse     `pulumi:"correlationScheme"`
	DefaultMoveCost              *string                                     `pulumi:"defaultMoveCost"`
	Etag                         string                                      `pulumi:"etag"`
	Id                           string                                      `pulumi:"id"`
	Location                     *string                                     `pulumi:"location"`
	Name                         string                                      `pulumi:"name"`
	PartitionDescription         interface{}                                 `pulumi:"partitionDescription"`
	PlacementConstraints         *string                                     `pulumi:"placementConstraints"`
	ProvisioningState            string                                      `pulumi:"provisioningState"`
	ServiceDnsName               *string                                     `pulumi:"serviceDnsName"`
	ServiceKind                  string                                      `pulumi:"serviceKind"`
	ServiceLoadMetrics           []ServiceLoadMetricDescriptionResponse      `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode *string                                     `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     []ServicePlacementPolicyDescriptionResponse `pulumi:"servicePlacementPolicies"`
	ServiceTypeName              *string                                     `pulumi:"serviceTypeName"`
	Tags                         map[string]string                           `pulumi:"tags"`
	Type                         string                                      `pulumi:"type"`
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
	ApplicationName   pulumi.StringInput `pulumi:"applicationName"`
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
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

func (o LookupServiceResultOutput) CorrelationScheme() ServiceCorrelationDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []ServiceCorrelationDescriptionResponse { return v.CorrelationScheme }).(ServiceCorrelationDescriptionResponseArrayOutput)
}

func (o LookupServiceResultOutput) DefaultMoveCost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.DefaultMoveCost }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) PartitionDescription() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupServiceResult) interface{} { return v.PartitionDescription }).(pulumi.AnyOutput)
}

func (o LookupServiceResultOutput) PlacementConstraints() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.PlacementConstraints }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) ServiceDnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceDnsName }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) ServiceKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ServiceKind }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) ServiceLoadMetrics() ServiceLoadMetricDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []ServiceLoadMetricDescriptionResponse { return v.ServiceLoadMetrics }).(ServiceLoadMetricDescriptionResponseArrayOutput)
}

func (o LookupServiceResultOutput) ServicePackageActivationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServicePackageActivationMode }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) ServicePlacementPolicies() ServicePlacementPolicyDescriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []ServicePlacementPolicyDescriptionResponse {
		return v.ServicePlacementPolicies
	}).(ServicePlacementPolicyDescriptionResponseArrayOutput)
}

func (o LookupServiceResultOutput) ServiceTypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.ServiceTypeName }).(pulumi.StringPtrOutput)
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
