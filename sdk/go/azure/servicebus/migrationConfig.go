


package servicebus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MigrationConfig struct {
	pulumi.CustomResourceState

	MigrationState                    pulumi.StringOutput  `pulumi:"migrationState"`
	Name                              pulumi.StringOutput  `pulumi:"name"`
	PendingReplicationOperationsCount pulumi.Float64Output `pulumi:"pendingReplicationOperationsCount"`
	PostMigrationName                 pulumi.StringOutput  `pulumi:"postMigrationName"`
	ProvisioningState                 pulumi.StringOutput  `pulumi:"provisioningState"`
	TargetNamespace                   pulumi.StringOutput  `pulumi:"targetNamespace"`
	Type                              pulumi.StringOutput  `pulumi:"type"`
}


func NewMigrationConfig(ctx *pulumi.Context,
	name string, args *MigrationConfigArgs, opts ...pulumi.ResourceOption) (*MigrationConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.PostMigrationName == nil {
		return nil, errors.New("invalid value for required argument 'PostMigrationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetNamespace == nil {
		return nil, errors.New("invalid value for required argument 'TargetNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicebus:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20170401:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210101preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210601preview:MigrationConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource MigrationConfig
	err := ctx.RegisterResource("azure-native:servicebus:MigrationConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMigrationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrationConfigState, opts ...pulumi.ResourceOption) (*MigrationConfig, error) {
	var resource MigrationConfig
	err := ctx.ReadResource("azure-native:servicebus:MigrationConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type migrationConfigState struct {
}

type MigrationConfigState struct {
}

func (MigrationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationConfigState)(nil)).Elem()
}

type migrationConfigArgs struct {
	ConfigName        *string `pulumi:"configName"`
	NamespaceName     string  `pulumi:"namespaceName"`
	PostMigrationName string  `pulumi:"postMigrationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	TargetNamespace   string  `pulumi:"targetNamespace"`
}


type MigrationConfigArgs struct {
	ConfigName        pulumi.StringPtrInput
	NamespaceName     pulumi.StringInput
	PostMigrationName pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TargetNamespace   pulumi.StringInput
}

func (MigrationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationConfigArgs)(nil)).Elem()
}

type MigrationConfigInput interface {
	pulumi.Input

	ToMigrationConfigOutput() MigrationConfigOutput
	ToMigrationConfigOutputWithContext(ctx context.Context) MigrationConfigOutput
}

func (*MigrationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationConfig)(nil))
}

func (i *MigrationConfig) ToMigrationConfigOutput() MigrationConfigOutput {
	return i.ToMigrationConfigOutputWithContext(context.Background())
}

func (i *MigrationConfig) ToMigrationConfigOutputWithContext(ctx context.Context) MigrationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationConfigOutput)
}

type MigrationConfigOutput struct{ *pulumi.OutputState }

func (MigrationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationConfig)(nil))
}

func (o MigrationConfigOutput) ToMigrationConfigOutput() MigrationConfigOutput {
	return o
}

func (o MigrationConfigOutput) ToMigrationConfigOutputWithContext(ctx context.Context) MigrationConfigOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationConfigInput)(nil)).Elem(), &MigrationConfig{})
	pulumi.RegisterOutputType(MigrationConfigOutput{})
}
