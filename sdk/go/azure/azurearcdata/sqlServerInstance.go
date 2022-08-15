


package azurearcdata

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlServerInstance struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                       `pulumi:"location"`
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties SqlServerInstancePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
}


func NewSqlServerInstance(ctx *pulumi.Context,
	name string, args *SqlServerInstanceArgs, opts ...pulumi.ResourceOption) (*SqlServerInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210601preview:SqlServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210701preview:SqlServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20210801:SqlServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20211101:SqlServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220301preview:SqlServerInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlServerInstance
	err := ctx.RegisterResource("azure-native:azurearcdata:SqlServerInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlServerInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerInstanceState, opts ...pulumi.ResourceOption) (*SqlServerInstance, error) {
	var resource SqlServerInstance
	err := ctx.ReadResource("azure-native:azurearcdata:SqlServerInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlServerInstanceState struct {
}

type SqlServerInstanceState struct {
}

func (SqlServerInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerInstanceState)(nil)).Elem()
}

type sqlServerInstanceArgs struct {
	Location              *string                      `pulumi:"location"`
	Properties            *SqlServerInstanceProperties `pulumi:"properties"`
	ResourceGroupName     string                       `pulumi:"resourceGroupName"`
	SqlServerInstanceName *string                      `pulumi:"sqlServerInstanceName"`
	Tags                  map[string]string            `pulumi:"tags"`
}


type SqlServerInstanceArgs struct {
	Location              pulumi.StringPtrInput
	Properties            SqlServerInstancePropertiesPtrInput
	ResourceGroupName     pulumi.StringInput
	SqlServerInstanceName pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
}

func (SqlServerInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerInstanceArgs)(nil)).Elem()
}

type SqlServerInstanceInput interface {
	pulumi.Input

	ToSqlServerInstanceOutput() SqlServerInstanceOutput
	ToSqlServerInstanceOutputWithContext(ctx context.Context) SqlServerInstanceOutput
}

func (*SqlServerInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerInstance)(nil)).Elem()
}

func (i *SqlServerInstance) ToSqlServerInstanceOutput() SqlServerInstanceOutput {
	return i.ToSqlServerInstanceOutputWithContext(context.Background())
}

func (i *SqlServerInstance) ToSqlServerInstanceOutputWithContext(ctx context.Context) SqlServerInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerInstanceOutput)
}

type SqlServerInstanceOutput struct{ *pulumi.OutputState }

func (SqlServerInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerInstance)(nil)).Elem()
}

func (o SqlServerInstanceOutput) ToSqlServerInstanceOutput() SqlServerInstanceOutput {
	return o
}

func (o SqlServerInstanceOutput) ToSqlServerInstanceOutputWithContext(ctx context.Context) SqlServerInstanceOutput {
	return o
}

func (o SqlServerInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SqlServerInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlServerInstanceOutput) Properties() SqlServerInstancePropertiesResponseOutput {
	return o.ApplyT(func(v *SqlServerInstance) SqlServerInstancePropertiesResponseOutput { return v.Properties }).(SqlServerInstancePropertiesResponseOutput)
}

func (o SqlServerInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlServerInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlServerInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlServerInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlServerInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlServerInstanceOutput{})
}
