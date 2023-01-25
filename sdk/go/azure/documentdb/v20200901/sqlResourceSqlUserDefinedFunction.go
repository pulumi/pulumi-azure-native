


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SqlResourceSqlUserDefinedFunction struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                                       `pulumi:"location"`
	Name     pulumi.StringOutput                                          `pulumi:"name"`
	Resource SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                                       `pulumi:"tags"`
	Type     pulumi.StringOutput                                          `pulumi:"type"`
}


func NewSqlResourceSqlUserDefinedFunction(ctx *pulumi.Context,
	name string, args *SqlResourceSqlUserDefinedFunctionArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlUserDefinedFunction, error) {
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
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlUserDefinedFunction"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlUserDefinedFunction"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlUserDefinedFunction
	err := ctx.RegisterResource("azure-native:documentdb/v20200901:SqlResourceSqlUserDefinedFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlResourceSqlUserDefinedFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlUserDefinedFunctionState, opts ...pulumi.ResourceOption) (*SqlResourceSqlUserDefinedFunction, error) {
	var resource SqlResourceSqlUserDefinedFunction
	err := ctx.ReadResource("azure-native:documentdb/v20200901:SqlResourceSqlUserDefinedFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlResourceSqlUserDefinedFunctionState struct {
}

type SqlResourceSqlUserDefinedFunctionState struct {
}

func (SqlResourceSqlUserDefinedFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlUserDefinedFunctionState)(nil)).Elem()
}

type sqlResourceSqlUserDefinedFunctionArgs struct {
	AccountName             string                         `pulumi:"accountName"`
	ContainerName           string                         `pulumi:"containerName"`
	DatabaseName            string                         `pulumi:"databaseName"`
	Location                *string                        `pulumi:"location"`
	Options                 *CreateUpdateOptions           `pulumi:"options"`
	Resource                SqlUserDefinedFunctionResource `pulumi:"resource"`
	ResourceGroupName       string                         `pulumi:"resourceGroupName"`
	Tags                    map[string]string              `pulumi:"tags"`
	UserDefinedFunctionName *string                        `pulumi:"userDefinedFunctionName"`
}


type SqlResourceSqlUserDefinedFunctionArgs struct {
	AccountName             pulumi.StringInput
	ContainerName           pulumi.StringInput
	DatabaseName            pulumi.StringInput
	Location                pulumi.StringPtrInput
	Options                 CreateUpdateOptionsPtrInput
	Resource                SqlUserDefinedFunctionResourceInput
	ResourceGroupName       pulumi.StringInput
	Tags                    pulumi.StringMapInput
	UserDefinedFunctionName pulumi.StringPtrInput
}

func (SqlResourceSqlUserDefinedFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlUserDefinedFunctionArgs)(nil)).Elem()
}

type SqlResourceSqlUserDefinedFunctionInput interface {
	pulumi.Input

	ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput
	ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput
}

func (*SqlResourceSqlUserDefinedFunction) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlUserDefinedFunction)(nil)).Elem()
}

func (i *SqlResourceSqlUserDefinedFunction) ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput {
	return i.ToSqlResourceSqlUserDefinedFunctionOutputWithContext(context.Background())
}

func (i *SqlResourceSqlUserDefinedFunction) ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlUserDefinedFunctionOutput)
}

type SqlResourceSqlUserDefinedFunctionOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlUserDefinedFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlUserDefinedFunction)(nil)).Elem()
}

func (o SqlResourceSqlUserDefinedFunctionOutput) ToSqlResourceSqlUserDefinedFunctionOutput() SqlResourceSqlUserDefinedFunctionOutput {
	return o
}

func (o SqlResourceSqlUserDefinedFunctionOutput) ToSqlResourceSqlUserDefinedFunctionOutputWithContext(ctx context.Context) SqlResourceSqlUserDefinedFunctionOutput {
	return o
}

func (o SqlResourceSqlUserDefinedFunctionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SqlResourceSqlUserDefinedFunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlResourceSqlUserDefinedFunctionOutput) Resource() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput)
}

func (o SqlResourceSqlUserDefinedFunctionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlResourceSqlUserDefinedFunctionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlUserDefinedFunction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlUserDefinedFunctionOutput{})
}
