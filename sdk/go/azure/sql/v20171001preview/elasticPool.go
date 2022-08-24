


package v20171001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ElasticPool struct {
	pulumi.CustomResourceState

	CreationDate        pulumi.StringOutput                             `pulumi:"creationDate"`
	Kind                pulumi.StringOutput                             `pulumi:"kind"`
	LicenseType         pulumi.StringPtrOutput                          `pulumi:"licenseType"`
	Location            pulumi.StringOutput                             `pulumi:"location"`
	MaxSizeBytes        pulumi.Float64PtrOutput                         `pulumi:"maxSizeBytes"`
	Name                pulumi.StringOutput                             `pulumi:"name"`
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsResponsePtrOutput `pulumi:"perDatabaseSettings"`
	Sku                 SkuResponsePtrOutput                            `pulumi:"sku"`
	State               pulumi.StringOutput                             `pulumi:"state"`
	Tags                pulumi.StringMapOutput                          `pulumi:"tags"`
	Type                pulumi.StringOutput                             `pulumi:"type"`
	ZoneRedundant       pulumi.BoolPtrOutput                            `pulumi:"zoneRedundant"`
}


func NewElasticPool(ctx *pulumi.Context,
	name string, args *ElasticPoolArgs, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ElasticPool"),
		},
	})
	opts = append(opts, aliases)
	var resource ElasticPool
	err := ctx.RegisterResource("azure-native:sql/v20171001preview:ElasticPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticPoolState, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	var resource ElasticPool
	err := ctx.ReadResource("azure-native:sql/v20171001preview:ElasticPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type elasticPoolState struct {
}

type ElasticPoolState struct {
}

func (ElasticPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolState)(nil)).Elem()
}

type elasticPoolArgs struct {
	ElasticPoolName     *string                         `pulumi:"elasticPoolName"`
	LicenseType         *string                         `pulumi:"licenseType"`
	Location            *string                         `pulumi:"location"`
	MaxSizeBytes        *float64                        `pulumi:"maxSizeBytes"`
	PerDatabaseSettings *ElasticPoolPerDatabaseSettings `pulumi:"perDatabaseSettings"`
	ResourceGroupName   string                          `pulumi:"resourceGroupName"`
	ServerName          string                          `pulumi:"serverName"`
	Sku                 *Sku                            `pulumi:"sku"`
	Tags                map[string]string               `pulumi:"tags"`
	ZoneRedundant       *bool                           `pulumi:"zoneRedundant"`
}


type ElasticPoolArgs struct {
	ElasticPoolName     pulumi.StringPtrInput
	LicenseType         pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	MaxSizeBytes        pulumi.Float64PtrInput
	PerDatabaseSettings ElasticPoolPerDatabaseSettingsPtrInput
	ResourceGroupName   pulumi.StringInput
	ServerName          pulumi.StringInput
	Sku                 SkuPtrInput
	Tags                pulumi.StringMapInput
	ZoneRedundant       pulumi.BoolPtrInput
}

func (ElasticPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticPoolArgs)(nil)).Elem()
}

type ElasticPoolInput interface {
	pulumi.Input

	ToElasticPoolOutput() ElasticPoolOutput
	ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput
}

func (*ElasticPool) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil)).Elem()
}

func (i *ElasticPool) ToElasticPoolOutput() ElasticPoolOutput {
	return i.ToElasticPoolOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput)
}

type ElasticPoolOutput struct{ *pulumi.OutputState }

func (ElasticPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPool)(nil)).Elem()
}

func (o ElasticPoolOutput) ToElasticPoolOutput() ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o ElasticPoolOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ElasticPoolOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o ElasticPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ElasticPoolOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.Float64PtrOutput { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o ElasticPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ElasticPoolOutput) PerDatabaseSettings() ElasticPoolPerDatabaseSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ElasticPool) ElasticPoolPerDatabaseSettingsResponsePtrOutput { return v.PerDatabaseSettings }).(ElasticPoolPerDatabaseSettingsResponsePtrOutput)
}

func (o ElasticPoolOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *ElasticPool) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ElasticPoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ElasticPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ElasticPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ElasticPoolOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticPool) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticPoolOutput{})
}
