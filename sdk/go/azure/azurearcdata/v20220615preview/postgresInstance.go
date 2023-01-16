


package v20220615preview

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
	if args.Sku != nil {
		args.Sku = args.Sku.ToPostgresInstanceSkuPtrOutput().ApplyT(func(v *PostgresInstanceSku) *PostgresInstanceSku { return v.Defaults() }).(PostgresInstanceSkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210601preview:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210701preview:PostgresInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220301preview:PostgresInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource PostgresInstance
	err := ctx.RegisterResource("azure-native:azurearcdata/v20220615preview:PostgresInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPostgresInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PostgresInstanceState, opts ...pulumi.ResourceOption) (*PostgresInstance, error) {
	var resource PostgresInstance
	err := ctx.ReadResource("azure-native:azurearcdata/v20220615preview:PostgresInstance", name, id, state, &resource, opts...)
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

func (o PostgresInstanceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *PostgresInstance) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o PostgresInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PostgresInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PostgresInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PostgresInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PostgresInstanceOutput) Properties() PostgresInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *PostgresInstance) PostgresInstancePropertiesResponseOutput { return v.Properties }).(PostgresInstancePropertiesResponseOutput)
}

func (o PostgresInstanceOutput) Sku() PostgresInstanceSkuResponsePtrOutput {
	return o.ApplyT(func(v *PostgresInstance) PostgresInstanceSkuResponsePtrOutput { return v.Sku }).(PostgresInstanceSkuResponsePtrOutput)
}

func (o PostgresInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PostgresInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PostgresInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PostgresInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PostgresInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PostgresInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PostgresInstanceOutput{})
}
