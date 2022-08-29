


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlManagedInstance struct {
	pulumi.CustomResourceState

	ExtendedLocation ExtendedLocationResponsePtrOutput          `pulumi:"extendedLocation"`
	Location         pulumi.StringOutput                        `pulumi:"location"`
	Name             pulumi.StringOutput                        `pulumi:"name"`
	Properties       SqlManagedInstancePropertiesResponseOutput `pulumi:"properties"`
	Sku              SqlManagedInstanceSkuResponsePtrOutput     `pulumi:"sku"`
	SystemData       SystemDataResponseOutput                   `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput                     `pulumi:"tags"`
	Type             pulumi.StringOutput                        `pulumi:"type"`
}


func NewSqlManagedInstance(ctx *pulumi.Context,
	name string, args *SqlManagedInstanceArgs, opts ...pulumi.ResourceOption) (*SqlManagedInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Properties = args.Properties.ToSqlManagedInstancePropertiesOutput().ApplyT(func(v SqlManagedInstanceProperties) SqlManagedInstanceProperties { return *v.Defaults() }).(SqlManagedInstancePropertiesOutput)
	if args.Sku != nil {
		args.Sku = args.Sku.ToSqlManagedInstanceSkuPtrOutput().ApplyT(func(v *SqlManagedInstanceSku) *SqlManagedInstanceSku { return v.Defaults() }).(SqlManagedInstanceSkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata:SqlManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210601preview:SqlManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210701preview:SqlManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210801:SqlManagedInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220301preview:SqlManagedInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlManagedInstance
	err := ctx.RegisterResource("azure-native:azurearcdata/v20211101:SqlManagedInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlManagedInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlManagedInstanceState, opts ...pulumi.ResourceOption) (*SqlManagedInstance, error) {
	var resource SqlManagedInstance
	err := ctx.ReadResource("azure-native:azurearcdata/v20211101:SqlManagedInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlManagedInstanceState struct {
}

type SqlManagedInstanceState struct {
}

func (SqlManagedInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlManagedInstanceState)(nil)).Elem()
}

type sqlManagedInstanceArgs struct {
	ExtendedLocation       *ExtendedLocation            `pulumi:"extendedLocation"`
	Location               *string                      `pulumi:"location"`
	Properties             SqlManagedInstanceProperties `pulumi:"properties"`
	ResourceGroupName      string                       `pulumi:"resourceGroupName"`
	Sku                    *SqlManagedInstanceSku       `pulumi:"sku"`
	SqlManagedInstanceName *string                      `pulumi:"sqlManagedInstanceName"`
	Tags                   map[string]string            `pulumi:"tags"`
}


type SqlManagedInstanceArgs struct {
	ExtendedLocation       ExtendedLocationPtrInput
	Location               pulumi.StringPtrInput
	Properties             SqlManagedInstancePropertiesInput
	ResourceGroupName      pulumi.StringInput
	Sku                    SqlManagedInstanceSkuPtrInput
	SqlManagedInstanceName pulumi.StringPtrInput
	Tags                   pulumi.StringMapInput
}

func (SqlManagedInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlManagedInstanceArgs)(nil)).Elem()
}

type SqlManagedInstanceInput interface {
	pulumi.Input

	ToSqlManagedInstanceOutput() SqlManagedInstanceOutput
	ToSqlManagedInstanceOutputWithContext(ctx context.Context) SqlManagedInstanceOutput
}

func (*SqlManagedInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstance)(nil)).Elem()
}

func (i *SqlManagedInstance) ToSqlManagedInstanceOutput() SqlManagedInstanceOutput {
	return i.ToSqlManagedInstanceOutputWithContext(context.Background())
}

func (i *SqlManagedInstance) ToSqlManagedInstanceOutputWithContext(ctx context.Context) SqlManagedInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlManagedInstanceOutput)
}

type SqlManagedInstanceOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstance)(nil)).Elem()
}

func (o SqlManagedInstanceOutput) ToSqlManagedInstanceOutput() SqlManagedInstanceOutput {
	return o
}

func (o SqlManagedInstanceOutput) ToSqlManagedInstanceOutputWithContext(ctx context.Context) SqlManagedInstanceOutput {
	return o
}

func (o SqlManagedInstanceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *SqlManagedInstance) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o SqlManagedInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlManagedInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SqlManagedInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlManagedInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlManagedInstanceOutput) Properties() SqlManagedInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *SqlManagedInstance) SqlManagedInstancePropertiesResponseOutput { return v.Properties }).(SqlManagedInstancePropertiesResponseOutput)
}

func (o SqlManagedInstanceOutput) Sku() SqlManagedInstanceSkuResponsePtrOutput {
	return o.ApplyT(func(v *SqlManagedInstance) SqlManagedInstanceSkuResponsePtrOutput { return v.Sku }).(SqlManagedInstanceSkuResponsePtrOutput)
}

func (o SqlManagedInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlManagedInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlManagedInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlManagedInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlManagedInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlManagedInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlManagedInstanceOutput{})
}
