


package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ElasticPool struct {
	pulumi.CustomResourceState

	CreationDate   pulumi.StringOutput    `pulumi:"creationDate"`
	DatabaseDtuMax pulumi.IntPtrOutput    `pulumi:"databaseDtuMax"`
	DatabaseDtuMin pulumi.IntPtrOutput    `pulumi:"databaseDtuMin"`
	Dtu            pulumi.IntPtrOutput    `pulumi:"dtu"`
	Edition        pulumi.StringPtrOutput `pulumi:"edition"`
	Kind           pulumi.StringOutput    `pulumi:"kind"`
	Location       pulumi.StringOutput    `pulumi:"location"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	State          pulumi.StringOutput    `pulumi:"state"`
	StorageMB      pulumi.IntPtrOutput    `pulumi:"storageMB"`
	Tags           pulumi.StringMapOutput `pulumi:"tags"`
	Type           pulumi.StringOutput    `pulumi:"type"`
	ZoneRedundant  pulumi.BoolPtrOutput   `pulumi:"zoneRedundant"`
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
			Type: pulumi.String("azure-native:sql/v20171001preview:ElasticPool"),
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
	})
	opts = append(opts, aliases)
	var resource ElasticPool
	err := ctx.RegisterResource("azure-native:sql/v20140401:ElasticPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetElasticPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticPoolState, opts ...pulumi.ResourceOption) (*ElasticPool, error) {
	var resource ElasticPool
	err := ctx.ReadResource("azure-native:sql/v20140401:ElasticPool", name, id, state, &resource, opts...)
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
	DatabaseDtuMax    *int              `pulumi:"databaseDtuMax"`
	DatabaseDtuMin    *int              `pulumi:"databaseDtuMin"`
	Dtu               *int              `pulumi:"dtu"`
	Edition           *string           `pulumi:"edition"`
	ElasticPoolName   *string           `pulumi:"elasticPoolName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ServerName        string            `pulumi:"serverName"`
	StorageMB         *int              `pulumi:"storageMB"`
	Tags              map[string]string `pulumi:"tags"`
	ZoneRedundant     *bool             `pulumi:"zoneRedundant"`
}


type ElasticPoolArgs struct {
	DatabaseDtuMax    pulumi.IntPtrInput
	DatabaseDtuMin    pulumi.IntPtrInput
	Dtu               pulumi.IntPtrInput
	Edition           pulumi.StringPtrInput
	ElasticPoolName   pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	StorageMB         pulumi.IntPtrInput
	Tags              pulumi.StringMapInput
	ZoneRedundant     pulumi.BoolPtrInput
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

func init() {
	pulumi.RegisterOutputType(ElasticPoolOutput{})
}
