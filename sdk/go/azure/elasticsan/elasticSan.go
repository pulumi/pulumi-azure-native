


package elasticsan

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ElasticSan struct {
	pulumi.CustomResourceState

	AvailabilityZones       pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	BaseSizeTiB             pulumi.Float64Output     `pulumi:"baseSizeTiB"`
	ExtendedCapacitySizeTiB pulumi.Float64Output     `pulumi:"extendedCapacitySizeTiB"`
	Location                pulumi.StringPtrOutput   `pulumi:"location"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput      `pulumi:"provisioningState"`
	Sku                     SkuResponseOutput        `pulumi:"sku"`
	SystemData              SystemDataResponseOutput `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput   `pulumi:"tags"`
	TotalIops               pulumi.Float64Output     `pulumi:"totalIops"`
	TotalMBps               pulumi.Float64Output     `pulumi:"totalMBps"`
	TotalSizeTiB            pulumi.Float64Output     `pulumi:"totalSizeTiB"`
	TotalVolumeSizeGiB      pulumi.Float64Output     `pulumi:"totalVolumeSizeGiB"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
	VolumeGroupCount        pulumi.Float64Output     `pulumi:"volumeGroupCount"`
}


func NewElasticSan(ctx *pulumi.Context,
	name string, args *ElasticSanArgs, opts ...pulumi.ResourceOption) (*ElasticSan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseSizeTiB == nil {
		return nil, errors.New("invalid value for required argument 'BaseSizeTiB'")
	}
	if args.ExtendedCapacitySizeTiB == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedCapacitySizeTiB'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:elasticsan/v20211120preview:ElasticSan"),
		},
	})
	opts = append(opts, aliases)
	var resource ElasticSan
	err := ctx.RegisterResource("azure-native:elasticsan:ElasticSan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetElasticSan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticSanState, opts ...pulumi.ResourceOption) (*ElasticSan, error) {
	var resource ElasticSan
	err := ctx.ReadResource("azure-native:elasticsan:ElasticSan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type elasticSanState struct {
}

type ElasticSanState struct {
}

func (ElasticSanState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticSanState)(nil)).Elem()
}

type elasticSanArgs struct {
	AvailabilityZones       []string          `pulumi:"availabilityZones"`
	BaseSizeTiB             float64           `pulumi:"baseSizeTiB"`
	ElasticSanName          *string           `pulumi:"elasticSanName"`
	ExtendedCapacitySizeTiB float64           `pulumi:"extendedCapacitySizeTiB"`
	Location                *string           `pulumi:"location"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	Sku                     Sku               `pulumi:"sku"`
	Tags                    map[string]string `pulumi:"tags"`
}


type ElasticSanArgs struct {
	AvailabilityZones       pulumi.StringArrayInput
	BaseSizeTiB             pulumi.Float64Input
	ElasticSanName          pulumi.StringPtrInput
	ExtendedCapacitySizeTiB pulumi.Float64Input
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Sku                     SkuInput
	Tags                    pulumi.StringMapInput
}

func (ElasticSanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticSanArgs)(nil)).Elem()
}

type ElasticSanInput interface {
	pulumi.Input

	ToElasticSanOutput() ElasticSanOutput
	ToElasticSanOutputWithContext(ctx context.Context) ElasticSanOutput
}

func (*ElasticSan) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticSan)(nil)).Elem()
}

func (i *ElasticSan) ToElasticSanOutput() ElasticSanOutput {
	return i.ToElasticSanOutputWithContext(context.Background())
}

func (i *ElasticSan) ToElasticSanOutputWithContext(ctx context.Context) ElasticSanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticSanOutput)
}

type ElasticSanOutput struct{ *pulumi.OutputState }

func (ElasticSanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticSan)(nil)).Elem()
}

func (o ElasticSanOutput) ToElasticSanOutput() ElasticSanOutput {
	return o
}

func (o ElasticSanOutput) ToElasticSanOutputWithContext(ctx context.Context) ElasticSanOutput {
	return o
}

func (o ElasticSanOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ElasticSanOutput) BaseSizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.BaseSizeTiB }).(pulumi.Float64Output)
}

func (o ElasticSanOutput) ExtendedCapacitySizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.ExtendedCapacitySizeTiB }).(pulumi.Float64Output)
}

func (o ElasticSanOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ElasticSanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ElasticSanOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ElasticSanOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *ElasticSan) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o ElasticSanOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ElasticSan) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ElasticSanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ElasticSanOutput) TotalIops() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.TotalIops }).(pulumi.Float64Output)
}

func (o ElasticSanOutput) TotalMBps() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.TotalMBps }).(pulumi.Float64Output)
}

func (o ElasticSanOutput) TotalSizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.TotalSizeTiB }).(pulumi.Float64Output)
}

func (o ElasticSanOutput) TotalVolumeSizeGiB() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.TotalVolumeSizeGiB }).(pulumi.Float64Output)
}

func (o ElasticSanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ElasticSanOutput) VolumeGroupCount() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.VolumeGroupCount }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(ElasticSanOutput{})
}
