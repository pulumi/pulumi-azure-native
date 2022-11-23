


package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SqlResourceSqlContainer struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                             `pulumi:"location"`
	Name     pulumi.StringOutput                                `pulumi:"name"`
	Options  SqlContainerGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource SqlContainerGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                             `pulumi:"tags"`
	Type     pulumi.StringOutput                                `pulumi:"type"`
}


func NewSqlResourceSqlContainer(ctx *pulumi.Context,
	name string, args *SqlResourceSqlContainerArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Resource = args.Resource.ToSqlContainerResourceOutput().ApplyT(func(v SqlContainerResource) SqlContainerResource { return *v.Defaults() }).(SqlContainerResourceOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlContainer"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlContainer
	err := ctx.RegisterResource("azure-native:documentdb/v20200401:SqlResourceSqlContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlResourceSqlContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlContainerState, opts ...pulumi.ResourceOption) (*SqlResourceSqlContainer, error) {
	var resource SqlResourceSqlContainer
	err := ctx.ReadResource("azure-native:documentdb/v20200401:SqlResourceSqlContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlResourceSqlContainerState struct {
}

type SqlResourceSqlContainerState struct {
}

func (SqlResourceSqlContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlContainerState)(nil)).Elem()
}

type sqlResourceSqlContainerArgs struct {
	AccountName       string               `pulumi:"accountName"`
	ContainerName     *string              `pulumi:"containerName"`
	DatabaseName      string               `pulumi:"databaseName"`
	Location          *string              `pulumi:"location"`
	Options           CreateUpdateOptions  `pulumi:"options"`
	Resource          SqlContainerResource `pulumi:"resource"`
	ResourceGroupName string               `pulumi:"resourceGroupName"`
	Tags              map[string]string    `pulumi:"tags"`
}


type SqlResourceSqlContainerArgs struct {
	AccountName       pulumi.StringInput
	ContainerName     pulumi.StringPtrInput
	DatabaseName      pulumi.StringInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsInput
	Resource          SqlContainerResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (SqlResourceSqlContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlContainerArgs)(nil)).Elem()
}

type SqlResourceSqlContainerInput interface {
	pulumi.Input

	ToSqlResourceSqlContainerOutput() SqlResourceSqlContainerOutput
	ToSqlResourceSqlContainerOutputWithContext(ctx context.Context) SqlResourceSqlContainerOutput
}

func (*SqlResourceSqlContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlContainer)(nil)).Elem()
}

func (i *SqlResourceSqlContainer) ToSqlResourceSqlContainerOutput() SqlResourceSqlContainerOutput {
	return i.ToSqlResourceSqlContainerOutputWithContext(context.Background())
}

func (i *SqlResourceSqlContainer) ToSqlResourceSqlContainerOutputWithContext(ctx context.Context) SqlResourceSqlContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlContainerOutput)
}

type SqlResourceSqlContainerOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlContainer)(nil)).Elem()
}

func (o SqlResourceSqlContainerOutput) ToSqlResourceSqlContainerOutput() SqlResourceSqlContainerOutput {
	return o
}

func (o SqlResourceSqlContainerOutput) ToSqlResourceSqlContainerOutputWithContext(ctx context.Context) SqlResourceSqlContainerOutput {
	return o
}

func (o SqlResourceSqlContainerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SqlResourceSqlContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlContainerOutput) Options() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) SqlContainerGetPropertiesResponseOptionsPtrOutput { return v.Options }).(SqlContainerGetPropertiesResponseOptionsPtrOutput)
}

func (o SqlResourceSqlContainerOutput) Resource() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) SqlContainerGetPropertiesResponseResourcePtrOutput { return v.Resource }).(SqlContainerGetPropertiesResponseResourcePtrOutput)
}

func (o SqlResourceSqlContainerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlResourceSqlContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlContainerOutput{})
}
