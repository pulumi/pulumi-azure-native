


package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20220110preview:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	CustomResourceName string                    `pulumi:"customResourceName"`
	DatastoreIds       []string                  `pulumi:"datastoreIds"`
	ExtendedLocation   *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                 string                    `pulumi:"id"`
	InventoryItemId    *string                   `pulumi:"inventoryItemId"`
	Kind               *string                   `pulumi:"kind"`
	Location           string                    `pulumi:"location"`
	MoName             string                    `pulumi:"moName"`
	MoRefId            *string                   `pulumi:"moRefId"`
	Name               string                    `pulumi:"name"`
	NetworkIds         []string                  `pulumi:"networkIds"`
	ProvisioningState  string                    `pulumi:"provisioningState"`
	Statuses           []ResourceStatusResponse  `pulumi:"statuses"`
	SystemData         SystemDataResponse        `pulumi:"systemData"`
	Tags               map[string]string         `pulumi:"tags"`
	Type               string                    `pulumi:"type"`
	Uuid               string                    `pulumi:"uuid"`
	VCenterId          *string                   `pulumi:"vCenterId"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}


type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) DatastoreIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.DatastoreIds }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.MoName }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) NetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.NetworkIds }).(pulumi.StringArrayOutput)
}

func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
