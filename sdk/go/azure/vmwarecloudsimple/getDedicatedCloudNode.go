


package vmwarecloudsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDedicatedCloudNode(ctx *pulumi.Context, args *LookupDedicatedCloudNodeArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedCloudNodeResult, error) {
	var rv LookupDedicatedCloudNodeResult
	err := ctx.Invoke("azure-native:vmwarecloudsimple:getDedicatedCloudNode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDedicatedCloudNodeArgs struct {
	DedicatedCloudNodeName string `pulumi:"dedicatedCloudNodeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupDedicatedCloudNodeResult struct {
	AvailabilityZoneId   string            `pulumi:"availabilityZoneId"`
	AvailabilityZoneName string            `pulumi:"availabilityZoneName"`
	CloudRackName        string            `pulumi:"cloudRackName"`
	Created              string            `pulumi:"created"`
	Id                   string            `pulumi:"id"`
	Location             string            `pulumi:"location"`
	Name                 string            `pulumi:"name"`
	NodesCount           int               `pulumi:"nodesCount"`
	PlacementGroupId     string            `pulumi:"placementGroupId"`
	PlacementGroupName   string            `pulumi:"placementGroupName"`
	PrivateCloudId       string            `pulumi:"privateCloudId"`
	PrivateCloudName     string            `pulumi:"privateCloudName"`
	ProvisioningState    string            `pulumi:"provisioningState"`
	PurchaseId           string            `pulumi:"purchaseId"`
	Sku                  *SkuResponse      `pulumi:"sku"`
	Status               string            `pulumi:"status"`
	Tags                 map[string]string `pulumi:"tags"`
	Type                 string            `pulumi:"type"`
	VmwareClusterName    string            `pulumi:"vmwareClusterName"`
}

func LookupDedicatedCloudNodeOutput(ctx *pulumi.Context, args LookupDedicatedCloudNodeOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedCloudNodeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDedicatedCloudNodeResult, error) {
			args := v.(LookupDedicatedCloudNodeArgs)
			r, err := LookupDedicatedCloudNode(ctx, &args, opts...)
			var s LookupDedicatedCloudNodeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDedicatedCloudNodeResultOutput)
}

type LookupDedicatedCloudNodeOutputArgs struct {
	DedicatedCloudNodeName pulumi.StringInput `pulumi:"dedicatedCloudNodeName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDedicatedCloudNodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedCloudNodeArgs)(nil)).Elem()
}


type LookupDedicatedCloudNodeResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedCloudNodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedCloudNodeResult)(nil)).Elem()
}

func (o LookupDedicatedCloudNodeResultOutput) ToLookupDedicatedCloudNodeResultOutput() LookupDedicatedCloudNodeResultOutput {
	return o
}

func (o LookupDedicatedCloudNodeResultOutput) ToLookupDedicatedCloudNodeResultOutputWithContext(ctx context.Context) LookupDedicatedCloudNodeResultOutput {
	return o
}

func (o LookupDedicatedCloudNodeResultOutput) AvailabilityZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.AvailabilityZoneId }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) AvailabilityZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.AvailabilityZoneName }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) CloudRackName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.CloudRackName }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) NodesCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) int { return v.NodesCount }).(pulumi.IntOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) PlacementGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.PlacementGroupId }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) PlacementGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.PlacementGroupName }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) PrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.PrivateCloudId }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) PrivateCloudName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.PrivateCloudName }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) PurchaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.PurchaseId }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDedicatedCloudNodeResultOutput) VmwareClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedCloudNodeResult) string { return v.VmwareClusterName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedCloudNodeResultOutput{})
}
