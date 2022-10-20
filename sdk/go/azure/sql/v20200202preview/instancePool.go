


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstancePool struct {
	pulumi.CustomResourceState

	LicenseType pulumi.StringOutput    `pulumi:"licenseType"`
	Location    pulumi.StringOutput    `pulumi:"location"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Sku         SkuResponsePtrOutput   `pulumi:"sku"`
	SubnetId    pulumi.StringOutput    `pulumi:"subnetId"`
	Tags        pulumi.StringMapOutput `pulumi:"tags"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	VCores      pulumi.IntOutput       `pulumi:"vCores"`
}


func NewInstancePool(ctx *pulumi.Context,
	name string, args *InstancePoolArgs, opts ...pulumi.ResourceOption) (*InstancePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LicenseType == nil {
		return nil, errors.New("invalid value for required argument 'LicenseType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VCores == nil {
		return nil, errors.New("invalid value for required argument 'VCores'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:InstancePool"),
		},
	})
	opts = append(opts, aliases)
	var resource InstancePool
	err := ctx.RegisterResource("azure-native:sql/v20200202preview:InstancePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInstancePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstancePoolState, opts ...pulumi.ResourceOption) (*InstancePool, error) {
	var resource InstancePool
	err := ctx.ReadResource("azure-native:sql/v20200202preview:InstancePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type instancePoolState struct {
}

type InstancePoolState struct {
}

func (InstancePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePoolState)(nil)).Elem()
}

type instancePoolArgs struct {
	InstancePoolName  *string           `pulumi:"instancePoolName"`
	LicenseType       string            `pulumi:"licenseType"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *Sku              `pulumi:"sku"`
	SubnetId          string            `pulumi:"subnetId"`
	Tags              map[string]string `pulumi:"tags"`
	VCores            int               `pulumi:"vCores"`
}


type InstancePoolArgs struct {
	InstancePoolName  pulumi.StringPtrInput
	LicenseType       pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
	SubnetId          pulumi.StringInput
	Tags              pulumi.StringMapInput
	VCores            pulumi.IntInput
}

func (InstancePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePoolArgs)(nil)).Elem()
}

type InstancePoolInput interface {
	pulumi.Input

	ToInstancePoolOutput() InstancePoolOutput
	ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput
}

func (*InstancePool) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePool)(nil)).Elem()
}

func (i *InstancePool) ToInstancePoolOutput() InstancePoolOutput {
	return i.ToInstancePoolOutputWithContext(context.Background())
}

func (i *InstancePool) ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePoolOutput)
}

type InstancePoolOutput struct{ *pulumi.OutputState }

func (InstancePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePool)(nil)).Elem()
}

func (o InstancePoolOutput) ToInstancePoolOutput() InstancePoolOutput {
	return o
}

func (o InstancePoolOutput) ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput {
	return o
}

func (o InstancePoolOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.LicenseType }).(pulumi.StringOutput)
}

func (o InstancePoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o InstancePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InstancePoolOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *InstancePool) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o InstancePoolOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o InstancePoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o InstancePoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o InstancePoolOutput) VCores() pulumi.IntOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.IntOutput { return v.VCores }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(InstancePoolOutput{})
}
