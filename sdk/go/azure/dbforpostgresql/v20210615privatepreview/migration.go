


package v20210615privatepreview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Migration struct {
	pulumi.CustomResourceState

	CurrentStatus                             MigrationStatusResponseOutput              `pulumi:"currentStatus"`
	DBsToMigrate                              pulumi.StringArrayOutput                   `pulumi:"dBsToMigrate"`
	Location                                  pulumi.StringOutput                        `pulumi:"location"`
	MigrationDetailsLevel                     pulumi.StringOutput                        `pulumi:"migrationDetailsLevel"`
	MigrationId                               pulumi.StringOutput                        `pulumi:"migrationId"`
	MigrationName                             pulumi.StringOutput                        `pulumi:"migrationName"`
	MigrationResourceGroup                    MigrationResourceGroupResponsePtrOutput    `pulumi:"migrationResourceGroup"`
	MigrationWindowStartTimeInUtc             pulumi.StringPtrOutput                     `pulumi:"migrationWindowStartTimeInUtc"`
	Name                                      pulumi.StringOutput                        `pulumi:"name"`
	OverwriteDBsInTarget                      pulumi.BoolPtrOutput                       `pulumi:"overwriteDBsInTarget"`
	SecretParameters                          MigrationSecretParametersResponsePtrOutput `pulumi:"secretParameters"`
	SetupLogicalReplicationOnSourceDBIfNeeded pulumi.BoolPtrOutput                       `pulumi:"setupLogicalReplicationOnSourceDBIfNeeded"`
	SourceDBServerMetadata                    DBServerMetadataResponseOutput             `pulumi:"sourceDBServerMetadata"`
	SourceDBServerResourceId                  pulumi.StringPtrOutput                     `pulumi:"sourceDBServerResourceId"`
	StartDataMigration                        pulumi.BoolPtrOutput                       `pulumi:"startDataMigration"`
	SystemData                                SystemDataResponseOutput                   `pulumi:"systemData"`
	Tags                                      pulumi.StringMapOutput                     `pulumi:"tags"`
	TargetDBServerMetadata                    DBServerMetadataResponseOutput             `pulumi:"targetDBServerMetadata"`
	TargetDBServerResourceId                  pulumi.StringOutput                        `pulumi:"targetDBServerResourceId"`
	TriggerCutover                            pulumi.BoolPtrOutput                       `pulumi:"triggerCutover"`
	Type                                      pulumi.StringOutput                        `pulumi:"type"`
	UserAssignedIdentityResourceId            pulumi.StringPtrOutput                     `pulumi:"userAssignedIdentityResourceId"`
}


func NewMigration(ctx *pulumi.Context,
	name string, args *MigrationArgs, opts ...pulumi.ResourceOption) (*Migration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetDBServerName == nil {
		return nil, errors.New("invalid value for required argument 'TargetDBServerName'")
	}
	if args.TargetDBServerResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'TargetDBServerResourceGroupName'")
	}
	if args.TargetDBServerSubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'TargetDBServerSubscriptionId'")
	}
	var resource Migration
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20210615privatepreview:Migration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMigration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrationState, opts ...pulumi.ResourceOption) (*Migration, error) {
	var resource Migration
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20210615privatepreview:Migration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type migrationState struct {
}

type MigrationState struct {
}

func (MigrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationState)(nil)).Elem()
}

type migrationArgs struct {
	DBsToMigrate                              []string                   `pulumi:"dBsToMigrate"`
	Location                                  *string                    `pulumi:"location"`
	MigrationName                             *string                    `pulumi:"migrationName"`
	MigrationResourceGroup                    *MigrationResourceGroup    `pulumi:"migrationResourceGroup"`
	MigrationWindowStartTimeInUtc             *string                    `pulumi:"migrationWindowStartTimeInUtc"`
	OverwriteDBsInTarget                      *bool                      `pulumi:"overwriteDBsInTarget"`
	SecretParameters                          *MigrationSecretParameters `pulumi:"secretParameters"`
	SetupLogicalReplicationOnSourceDBIfNeeded *bool                      `pulumi:"setupLogicalReplicationOnSourceDBIfNeeded"`
	SourceDBServerResourceId                  *string                    `pulumi:"sourceDBServerResourceId"`
	StartDataMigration                        *bool                      `pulumi:"startDataMigration"`
	Tags                                      map[string]string          `pulumi:"tags"`
	TargetDBServerName                        string                     `pulumi:"targetDBServerName"`
	TargetDBServerResourceGroupName           string                     `pulumi:"targetDBServerResourceGroupName"`
	TargetDBServerSubscriptionId              string                     `pulumi:"targetDBServerSubscriptionId"`
	TriggerCutover                            *bool                      `pulumi:"triggerCutover"`
	UserAssignedIdentityResourceId            *string                    `pulumi:"userAssignedIdentityResourceId"`
}


type MigrationArgs struct {
	DBsToMigrate                              pulumi.StringArrayInput
	Location                                  pulumi.StringPtrInput
	MigrationName                             pulumi.StringPtrInput
	MigrationResourceGroup                    MigrationResourceGroupPtrInput
	MigrationWindowStartTimeInUtc             pulumi.StringPtrInput
	OverwriteDBsInTarget                      pulumi.BoolPtrInput
	SecretParameters                          MigrationSecretParametersPtrInput
	SetupLogicalReplicationOnSourceDBIfNeeded pulumi.BoolPtrInput
	SourceDBServerResourceId                  pulumi.StringPtrInput
	StartDataMigration                        pulumi.BoolPtrInput
	Tags                                      pulumi.StringMapInput
	TargetDBServerName                        pulumi.StringInput
	TargetDBServerResourceGroupName           pulumi.StringInput
	TargetDBServerSubscriptionId              pulumi.StringInput
	TriggerCutover                            pulumi.BoolPtrInput
	UserAssignedIdentityResourceId            pulumi.StringPtrInput
}

func (MigrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationArgs)(nil)).Elem()
}

type MigrationInput interface {
	pulumi.Input

	ToMigrationOutput() MigrationOutput
	ToMigrationOutputWithContext(ctx context.Context) MigrationOutput
}

func (*Migration) ElementType() reflect.Type {
	return reflect.TypeOf((**Migration)(nil)).Elem()
}

func (i *Migration) ToMigrationOutput() MigrationOutput {
	return i.ToMigrationOutputWithContext(context.Background())
}

func (i *Migration) ToMigrationOutputWithContext(ctx context.Context) MigrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationOutput)
}

type MigrationOutput struct{ *pulumi.OutputState }

func (MigrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Migration)(nil)).Elem()
}

func (o MigrationOutput) ToMigrationOutput() MigrationOutput {
	return o
}

func (o MigrationOutput) ToMigrationOutputWithContext(ctx context.Context) MigrationOutput {
	return o
}

func (o MigrationOutput) CurrentStatus() MigrationStatusResponseOutput {
	return o.ApplyT(func(v *Migration) MigrationStatusResponseOutput { return v.CurrentStatus }).(MigrationStatusResponseOutput)
}

func (o MigrationOutput) DBsToMigrate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringArrayOutput { return v.DBsToMigrate }).(pulumi.StringArrayOutput)
}

func (o MigrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MigrationOutput) MigrationDetailsLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.MigrationDetailsLevel }).(pulumi.StringOutput)
}

func (o MigrationOutput) MigrationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.MigrationId }).(pulumi.StringOutput)
}

func (o MigrationOutput) MigrationName() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.MigrationName }).(pulumi.StringOutput)
}

func (o MigrationOutput) MigrationResourceGroup() MigrationResourceGroupResponsePtrOutput {
	return o.ApplyT(func(v *Migration) MigrationResourceGroupResponsePtrOutput { return v.MigrationResourceGroup }).(MigrationResourceGroupResponsePtrOutput)
}

func (o MigrationOutput) MigrationWindowStartTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.MigrationWindowStartTimeInUtc }).(pulumi.StringPtrOutput)
}

func (o MigrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MigrationOutput) OverwriteDBsInTarget() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.BoolPtrOutput { return v.OverwriteDBsInTarget }).(pulumi.BoolPtrOutput)
}

func (o MigrationOutput) SecretParameters() MigrationSecretParametersResponsePtrOutput {
	return o.ApplyT(func(v *Migration) MigrationSecretParametersResponsePtrOutput { return v.SecretParameters }).(MigrationSecretParametersResponsePtrOutput)
}

func (o MigrationOutput) SetupLogicalReplicationOnSourceDBIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.BoolPtrOutput { return v.SetupLogicalReplicationOnSourceDBIfNeeded }).(pulumi.BoolPtrOutput)
}

func (o MigrationOutput) SourceDBServerMetadata() DBServerMetadataResponseOutput {
	return o.ApplyT(func(v *Migration) DBServerMetadataResponseOutput { return v.SourceDBServerMetadata }).(DBServerMetadataResponseOutput)
}

func (o MigrationOutput) SourceDBServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.SourceDBServerResourceId }).(pulumi.StringPtrOutput)
}

func (o MigrationOutput) StartDataMigration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.BoolPtrOutput { return v.StartDataMigration }).(pulumi.BoolPtrOutput)
}

func (o MigrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Migration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MigrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MigrationOutput) TargetDBServerMetadata() DBServerMetadataResponseOutput {
	return o.ApplyT(func(v *Migration) DBServerMetadataResponseOutput { return v.TargetDBServerMetadata }).(DBServerMetadataResponseOutput)
}

func (o MigrationOutput) TargetDBServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.TargetDBServerResourceId }).(pulumi.StringOutput)
}

func (o MigrationOutput) TriggerCutover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.BoolPtrOutput { return v.TriggerCutover }).(pulumi.BoolPtrOutput)
}

func (o MigrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MigrationOutput) UserAssignedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.UserAssignedIdentityResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MigrationOutput{})
}
