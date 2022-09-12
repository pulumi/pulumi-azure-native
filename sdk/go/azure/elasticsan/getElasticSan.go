


package elasticsan

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupElasticSan(ctx *pulumi.Context, args *LookupElasticSanArgs, opts ...pulumi.InvokeOption) (*LookupElasticSanResult, error) {
	var rv LookupElasticSanResult
	err := ctx.Invoke("azure-native:elasticsan:getElasticSan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupElasticSanArgs struct {
	ElasticSanName    string `pulumi:"elasticSanName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupElasticSanResult struct {
	AvailabilityZones       []string           `pulumi:"availabilityZones"`
	BaseSizeTiB             float64            `pulumi:"baseSizeTiB"`
	ExtendedCapacitySizeTiB float64            `pulumi:"extendedCapacitySizeTiB"`
	Id                      string             `pulumi:"id"`
	Location                *string            `pulumi:"location"`
	Name                    string             `pulumi:"name"`
	ProvisioningState       string             `pulumi:"provisioningState"`
	Sku                     SkuResponse        `pulumi:"sku"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Tags                    map[string]string  `pulumi:"tags"`
	TotalIops               float64            `pulumi:"totalIops"`
	TotalMBps               float64            `pulumi:"totalMBps"`
	TotalSizeTiB            float64            `pulumi:"totalSizeTiB"`
	TotalVolumeSizeGiB      float64            `pulumi:"totalVolumeSizeGiB"`
	Type                    string             `pulumi:"type"`
	VolumeGroupCount        float64            `pulumi:"volumeGroupCount"`
}

func LookupElasticSanOutput(ctx *pulumi.Context, args LookupElasticSanOutputArgs, opts ...pulumi.InvokeOption) LookupElasticSanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupElasticSanResult, error) {
			args := v.(LookupElasticSanArgs)
			r, err := LookupElasticSan(ctx, &args, opts...)
			var s LookupElasticSanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupElasticSanResultOutput)
}

type LookupElasticSanOutputArgs struct {
	ElasticSanName    pulumi.StringInput `pulumi:"elasticSanName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupElasticSanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticSanArgs)(nil)).Elem()
}


type LookupElasticSanResultOutput struct{ *pulumi.OutputState }

func (LookupElasticSanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticSanResult)(nil)).Elem()
}

func (o LookupElasticSanResultOutput) ToLookupElasticSanResultOutput() LookupElasticSanResultOutput {
	return o
}

func (o LookupElasticSanResultOutput) ToLookupElasticSanResultOutputWithContext(ctx context.Context) LookupElasticSanResultOutput {
	return o
}

func (o LookupElasticSanResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupElasticSanResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupElasticSanResultOutput) BaseSizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.BaseSizeTiB }).(pulumi.Float64Output)
}

func (o LookupElasticSanResultOutput) ExtendedCapacitySizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.ExtendedCapacitySizeTiB }).(pulumi.Float64Output)
}

func (o LookupElasticSanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupElasticSanResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupElasticSanResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupElasticSanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupElasticSanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupElasticSanResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupElasticSanResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupElasticSanResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupElasticSanResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupElasticSanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupElasticSanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupElasticSanResultOutput) TotalIops() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.TotalIops }).(pulumi.Float64Output)
}

func (o LookupElasticSanResultOutput) TotalMBps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.TotalMBps }).(pulumi.Float64Output)
}

func (o LookupElasticSanResultOutput) TotalSizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.TotalSizeTiB }).(pulumi.Float64Output)
}

func (o LookupElasticSanResultOutput) TotalVolumeSizeGiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.TotalVolumeSizeGiB }).(pulumi.Float64Output)
}

func (o LookupElasticSanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupElasticSanResultOutput) VolumeGroupCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.VolumeGroupCount }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupElasticSanResultOutput{})
}
