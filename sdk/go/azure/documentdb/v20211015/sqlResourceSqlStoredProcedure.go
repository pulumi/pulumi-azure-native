


package v20211015

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlResourceSqlStoredProcedure struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                                   `pulumi:"location"`
	Name     pulumi.StringOutput                                      `pulumi:"name"`
	Resource SqlStoredProcedureGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                                   `pulumi:"tags"`
	Type     pulumi.StringOutput                                      `pulumi:"type"`
}


func NewSqlResourceSqlStoredProcedure(ctx *pulumi.Context,
	name string, args *SqlResourceSqlStoredProcedureArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlStoredProcedure, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlStoredProcedure"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlStoredProcedure"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlStoredProcedure
	err := ctx.RegisterResource("azure-native:documentdb/v20211015:SqlResourceSqlStoredProcedure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlResourceSqlStoredProcedure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlStoredProcedureState, opts ...pulumi.ResourceOption) (*SqlResourceSqlStoredProcedure, error) {
	var resource SqlResourceSqlStoredProcedure
	err := ctx.ReadResource("azure-native:documentdb/v20211015:SqlResourceSqlStoredProcedure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlResourceSqlStoredProcedureState struct {
}

type SqlResourceSqlStoredProcedureState struct {
}

func (SqlResourceSqlStoredProcedureState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlStoredProcedureState)(nil)).Elem()
}

type sqlResourceSqlStoredProcedureArgs struct {
	AccountName         string                     `pulumi:"accountName"`
	ContainerName       string                     `pulumi:"containerName"`
	DatabaseName        string                     `pulumi:"databaseName"`
	Location            *string                    `pulumi:"location"`
	Options             *CreateUpdateOptions       `pulumi:"options"`
	Resource            SqlStoredProcedureResource `pulumi:"resource"`
	ResourceGroupName   string                     `pulumi:"resourceGroupName"`
	StoredProcedureName *string                    `pulumi:"storedProcedureName"`
	Tags                map[string]string          `pulumi:"tags"`
}


type SqlResourceSqlStoredProcedureArgs struct {
	AccountName         pulumi.StringInput
	ContainerName       pulumi.StringInput
	DatabaseName        pulumi.StringInput
	Location            pulumi.StringPtrInput
	Options             CreateUpdateOptionsPtrInput
	Resource            SqlStoredProcedureResourceInput
	ResourceGroupName   pulumi.StringInput
	StoredProcedureName pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
}

func (SqlResourceSqlStoredProcedureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlStoredProcedureArgs)(nil)).Elem()
}

type SqlResourceSqlStoredProcedureInput interface {
	pulumi.Input

	ToSqlResourceSqlStoredProcedureOutput() SqlResourceSqlStoredProcedureOutput
	ToSqlResourceSqlStoredProcedureOutputWithContext(ctx context.Context) SqlResourceSqlStoredProcedureOutput
}

func (*SqlResourceSqlStoredProcedure) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlStoredProcedure)(nil)).Elem()
}

func (i *SqlResourceSqlStoredProcedure) ToSqlResourceSqlStoredProcedureOutput() SqlResourceSqlStoredProcedureOutput {
	return i.ToSqlResourceSqlStoredProcedureOutputWithContext(context.Background())
}

func (i *SqlResourceSqlStoredProcedure) ToSqlResourceSqlStoredProcedureOutputWithContext(ctx context.Context) SqlResourceSqlStoredProcedureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlStoredProcedureOutput)
}

type SqlResourceSqlStoredProcedureOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlStoredProcedureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlStoredProcedure)(nil)).Elem()
}

func (o SqlResourceSqlStoredProcedureOutput) ToSqlResourceSqlStoredProcedureOutput() SqlResourceSqlStoredProcedureOutput {
	return o
}

func (o SqlResourceSqlStoredProcedureOutput) ToSqlResourceSqlStoredProcedureOutputWithContext(ctx context.Context) SqlResourceSqlStoredProcedureOutput {
	return o
}

func (o SqlResourceSqlStoredProcedureOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SqlResourceSqlStoredProcedureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlStoredProcedureOutput) Resource() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(SqlStoredProcedureGetPropertiesResponseResourcePtrOutput)
}

func (o SqlResourceSqlStoredProcedureOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlResourceSqlStoredProcedureOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlStoredProcedure) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlStoredProcedureOutput{})
}
