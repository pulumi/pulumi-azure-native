


package v20190724preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlServer struct {
	pulumi.CustomResourceState

	Cores          pulumi.IntPtrOutput    `pulumi:"cores"`
	Edition        pulumi.StringPtrOutput `pulumi:"edition"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	PropertyBag    pulumi.StringPtrOutput `pulumi:"propertyBag"`
	RegistrationID pulumi.StringPtrOutput `pulumi:"registrationID"`
	Type           pulumi.StringOutput    `pulumi:"type"`
	Version        pulumi.StringPtrOutput `pulumi:"version"`
}


func NewSqlServer(ctx *pulumi.Context,
	name string, args *SqlServerArgs, opts ...pulumi.ResourceOption) (*SqlServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlServerRegistrationName == nil {
		return nil, errors.New("invalid value for required argument 'SqlServerRegistrationName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azuredata:SqlServer"),
		},
		{
			Type: pulumi.String("azure-native:azuredata/v20170301preview:SqlServer"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlServer
	err := ctx.RegisterResource("azure-native:azuredata/v20190724preview:SqlServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerState, opts ...pulumi.ResourceOption) (*SqlServer, error) {
	var resource SqlServer
	err := ctx.ReadResource("azure-native:azuredata/v20190724preview:SqlServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlServerState struct {
}

type SqlServerState struct {
}

func (SqlServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerState)(nil)).Elem()
}

type sqlServerArgs struct {
	Cores                     *int    `pulumi:"cores"`
	Edition                   *string `pulumi:"edition"`
	PropertyBag               *string `pulumi:"propertyBag"`
	RegistrationID            *string `pulumi:"registrationID"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	SqlServerName             *string `pulumi:"sqlServerName"`
	SqlServerRegistrationName string  `pulumi:"sqlServerRegistrationName"`
	Version                   *string `pulumi:"version"`
}


type SqlServerArgs struct {
	Cores                     pulumi.IntPtrInput
	Edition                   pulumi.StringPtrInput
	PropertyBag               pulumi.StringPtrInput
	RegistrationID            pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	SqlServerName             pulumi.StringPtrInput
	SqlServerRegistrationName pulumi.StringInput
	Version                   pulumi.StringPtrInput
}

func (SqlServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerArgs)(nil)).Elem()
}

type SqlServerInput interface {
	pulumi.Input

	ToSqlServerOutput() SqlServerOutput
	ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput
}

func (*SqlServer) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServer)(nil)).Elem()
}

func (i *SqlServer) ToSqlServerOutput() SqlServerOutput {
	return i.ToSqlServerOutputWithContext(context.Background())
}

func (i *SqlServer) ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerOutput)
}

type SqlServerOutput struct{ *pulumi.OutputState }

func (SqlServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServer)(nil)).Elem()
}

func (o SqlServerOutput) ToSqlServerOutput() SqlServerOutput {
	return o
}

func (o SqlServerOutput) ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput {
	return o
}

func (o SqlServerOutput) Cores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.IntPtrOutput { return v.Cores }).(pulumi.IntPtrOutput)
}

func (o SqlServerOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringPtrOutput { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o SqlServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlServerOutput) PropertyBag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringPtrOutput { return v.PropertyBag }).(pulumi.StringPtrOutput)
}

func (o SqlServerOutput) RegistrationID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringPtrOutput { return v.RegistrationID }).(pulumi.StringPtrOutput)
}

func (o SqlServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SqlServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlServerOutput{})
}
