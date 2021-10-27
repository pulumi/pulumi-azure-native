


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ElasticPool struct {
	pulumi.CustomResourceState

	CreationDate               pulumi.StringOutput                             `pulumi:"creationDate"`
	Kind                       pulumi.StringOutput                             `pulumi:"kind"`
	LicenseType                pulumi.StringPtrOutput                          `pulumi:"licenseType"`
	Location                   pulumi.StringOutput                             `pulumi:"location"`
	MaintenanceConfigurationId pulumi.StringPtrOutput                          `pulumi:"maintenanceConfigurationId"`
	MaxSizeBytes               pulumi.Float64PtrOutput                         `pulumi:"maxSizeBytes"`
	Name                       pulumi.StringOutput                             `pulumi:"name"`
	PerDatabaseSettings        ElasticPoolPerDatabaseSettingsResponsePtrOutput `pulumi:"perDatabaseSettings"`
	Sku                        SkuResponsePtrOutput                            `pulumi:"sku"`
	State                      pulumi.StringOutput                             `pulumi:"state"`
	Tags                       pulumi.StringMapOutput                          `pulumi:"tags"`
	Type                       pulumi.StringOutput                             `pulumi:"type"`
	ZoneRedundant              pulumi.BoolPtrOutput                            `pulumi:"zoneRedundant"`
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
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20140401:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20171001preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ElasticPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:ElasticPool"),
		},
	})
	opts = append(opts, aliases)
	var resource ElasticPool
	err := ctx.RegisterResource("azure-native:sql/v20201101preview:ElasticPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticPoolState, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	var resource ElasticPool
	err := ctx.ReadResource("azure-native:sql/v20201101preview:ElasticPool", name, id, state, &resource, opts...)
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
	ElasticPoolName            *string                         `pulumi:"elasticPoolName"`
	LicenseType                *string                         `pulumi:"licenseType"`
	Location                   *string                         `pulumi:"location"`
	MaintenanceConfigurationId *string                         `pulumi:"maintenanceConfigurationId"`
	MaxSizeBytes               *float64                        `pulumi:"maxSizeBytes"`
	PerDatabaseSettings        *ElasticPoolPerDatabaseSettings `pulumi:"perDatabaseSettings"`
	ResourceGroupName          string                          `pulumi:"resourceGroupName"`
	ServerName                 string                          `pulumi:"serverName"`
	Sku                        *Sku                            `pulumi:"sku"`
	Tags                       map[string]string               `pulumi:"tags"`
	ZoneRedundant              *bool                           `pulumi:"zoneRedundant"`
}


type ElasticPoolArgs struct {
	ElasticPoolName            pulumi.StringPtrInput
	LicenseType                pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	MaintenanceConfigurationId pulumi.StringPtrInput
	MaxSizeBytes               pulumi.Float64PtrInput
	PerDatabaseSettings        ElasticPoolPerDatabaseSettingsPtrInput
	ResourceGroupName          pulumi.StringInput
	ServerName                 pulumi.StringInput
	Sku                        SkuPtrInput
	Tags                       pulumi.StringMapInput
	ZoneRedundant              pulumi.BoolPtrInput
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
	return reflect.TypeOf((*ElasticPool)(nil))
}

func (i *ElasticPool) ToElasticPoolOutput() ElasticPoolOutput {
	return i.ToElasticPoolOutputWithContext(context.Background())
}

func (i *ElasticPool) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticPoolOutput)
}

type ElasticPoolOutput struct{ *pulumi.OutputState }

func (ElasticPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPool)(nil))
}

func (o ElasticPoolOutput) ToElasticPoolOutput() ElasticPoolOutput {
	return o
}

func (o ElasticPoolOutput) ToElasticPoolOutputWithContext(ctx context.Context) ElasticPoolOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticPoolInput)(nil)).Elem(), &ElasticPool{})
	pulumi.RegisterOutputType(ElasticPoolOutput{})
}
