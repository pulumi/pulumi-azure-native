


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PostgresInstance struct {
	pulumi.CustomResourceState

	ExtendedLocation ExtendedLocationResponsePtrOutput        `pulumi:"extendedLocation"`
	Location         pulumi.StringOutput                      `pulumi:"location"`
	Name             pulumi.StringOutput                      `pulumi:"name"`
	Properties       PostgresInstancePropertiesResponseOutput `pulumi:"properties"`
	Sku              PostgresInstanceSkuResponsePtrOutput     `pulumi:"sku"`
	SystemData       SystemDataResponseOutput                 `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput                   `pulumi:"tags"`
	Type             pulumi.StringOutput                      `pulumi:"type"`
}


func NewPostgresInstance(ctx *pulumi.Context,
	name string, args *PostgresInstanceArgs, opts ...pulumi.ResourceOption) (*PostgresInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	skuApplier := func(v PostgresInstanceSku) *PostgresInstanceSku { return v.Defaults() }
	if args.Sku != nil {
		args.Sku = args.Sku.ToPostgresInstanceSkuPtrOutput().Elem().ApplyT(skuApplier).(PostgresInstanceSkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210601preview:PostgresInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource PostgresInstance
	err := ctx.RegisterResource("azure-native:azurearcdata/v20210701preview:PostgresInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPostgresInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PostgresInstanceState, opts ...pulumi.ResourceOption) (*PostgresInstance, error) {
	var resource PostgresInstance
	err := ctx.ReadResource("azure-native:azurearcdata/v20210701preview:PostgresInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type postgresInstanceState struct {
}

type PostgresInstanceState struct {
}

func (PostgresInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*postgresInstanceState)(nil)).Elem()
}

type postgresInstanceArgs struct {
	ExtendedLocation     *ExtendedLocation          `pulumi:"extendedLocation"`
	Location             *string                    `pulumi:"location"`
	PostgresInstanceName *string                    `pulumi:"postgresInstanceName"`
	Properties           PostgresInstanceProperties `pulumi:"properties"`
	ResourceGroupName    string                     `pulumi:"resourceGroupName"`
	Sku                  *PostgresInstanceSku       `pulumi:"sku"`
	Tags                 map[string]string          `pulumi:"tags"`
}


type PostgresInstanceArgs struct {
	ExtendedLocation     ExtendedLocationPtrInput
	Location             pulumi.StringPtrInput
	PostgresInstanceName pulumi.StringPtrInput
	Properties           PostgresInstancePropertiesInput
	ResourceGroupName    pulumi.StringInput
	Sku                  PostgresInstanceSkuPtrInput
	Tags                 pulumi.StringMapInput
}

func (PostgresInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*postgresInstanceArgs)(nil)).Elem()
}

type PostgresInstanceInput interface {
	pulumi.Input

	ToPostgresInstanceOutput() PostgresInstanceOutput
	ToPostgresInstanceOutputWithContext(ctx context.Context) PostgresInstanceOutput
}

func (*PostgresInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstance)(nil)).Elem()
}

func (i *PostgresInstance) ToPostgresInstanceOutput() PostgresInstanceOutput {
	return i.ToPostgresInstanceOutputWithContext(context.Background())
}

func (i *PostgresInstance) ToPostgresInstanceOutputWithContext(ctx context.Context) PostgresInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PostgresInstanceOutput)
}

type PostgresInstanceOutput struct{ *pulumi.OutputState }

func (PostgresInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PostgresInstance)(nil)).Elem()
}

func (o PostgresInstanceOutput) ToPostgresInstanceOutput() PostgresInstanceOutput {
	return o
}

func (o PostgresInstanceOutput) ToPostgresInstanceOutputWithContext(ctx context.Context) PostgresInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PostgresInstanceOutput{})
}
