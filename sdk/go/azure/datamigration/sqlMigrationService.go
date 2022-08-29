


package datamigration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlMigrationService struct {
	pulumi.CustomResourceState

	IntegrationRuntimeState pulumi.StringOutput      `pulumi:"integrationRuntimeState"`
	Location                pulumi.StringPtrOutput   `pulumi:"location"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData              SystemDataResponseOutput `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput   `pulumi:"tags"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewSqlMigrationService(ctx *pulumi.Context,
	name string, args *SqlMigrationServiceArgs, opts ...pulumi.ResourceOption) (*SqlMigrationService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:SqlMigrationService"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:SqlMigrationService"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:SqlMigrationService"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlMigrationService
	err := ctx.RegisterResource("azure-native:datamigration:SqlMigrationService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlMigrationService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlMigrationServiceState, opts ...pulumi.ResourceOption) (*SqlMigrationService, error) {
	var resource SqlMigrationService
	err := ctx.ReadResource("azure-native:datamigration:SqlMigrationService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlMigrationServiceState struct {
}

type SqlMigrationServiceState struct {
}

func (SqlMigrationServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlMigrationServiceState)(nil)).Elem()
}

type sqlMigrationServiceArgs struct {
	Location                *string           `pulumi:"location"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	SqlMigrationServiceName *string           `pulumi:"sqlMigrationServiceName"`
	Tags                    map[string]string `pulumi:"tags"`
}


type SqlMigrationServiceArgs struct {
	Location                pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	SqlMigrationServiceName pulumi.StringPtrInput
	Tags                    pulumi.StringMapInput
}

func (SqlMigrationServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlMigrationServiceArgs)(nil)).Elem()
}

type SqlMigrationServiceInput interface {
	pulumi.Input

	ToSqlMigrationServiceOutput() SqlMigrationServiceOutput
	ToSqlMigrationServiceOutputWithContext(ctx context.Context) SqlMigrationServiceOutput
}

func (*SqlMigrationService) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlMigrationService)(nil)).Elem()
}

func (i *SqlMigrationService) ToSqlMigrationServiceOutput() SqlMigrationServiceOutput {
	return i.ToSqlMigrationServiceOutputWithContext(context.Background())
}

func (i *SqlMigrationService) ToSqlMigrationServiceOutputWithContext(ctx context.Context) SqlMigrationServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlMigrationServiceOutput)
}

type SqlMigrationServiceOutput struct{ *pulumi.OutputState }

func (SqlMigrationServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlMigrationService)(nil)).Elem()
}

func (o SqlMigrationServiceOutput) ToSqlMigrationServiceOutput() SqlMigrationServiceOutput {
	return o
}

func (o SqlMigrationServiceOutput) ToSqlMigrationServiceOutputWithContext(ctx context.Context) SqlMigrationServiceOutput {
	return o
}

func (o SqlMigrationServiceOutput) IntegrationRuntimeState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringOutput { return v.IntegrationRuntimeState }).(pulumi.StringOutput)
}

func (o SqlMigrationServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SqlMigrationServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlMigrationServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlMigrationServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlMigrationService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlMigrationServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlMigrationServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlMigrationService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlMigrationServiceOutput{})
}
