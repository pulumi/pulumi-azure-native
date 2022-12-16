


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Database struct {
	pulumi.CustomResourceState

	CatalogCollation            pulumi.StringPtrOutput  `pulumi:"catalogCollation"`
	Collation                   pulumi.StringPtrOutput  `pulumi:"collation"`
	CreationDate                pulumi.StringOutput     `pulumi:"creationDate"`
	CurrentServiceObjectiveName pulumi.StringOutput     `pulumi:"currentServiceObjectiveName"`
	DatabaseId                  pulumi.StringOutput     `pulumi:"databaseId"`
	DefaultSecondaryLocation    pulumi.StringOutput     `pulumi:"defaultSecondaryLocation"`
	ElasticPoolId               pulumi.StringPtrOutput  `pulumi:"elasticPoolId"`
	FailoverGroupId             pulumi.StringOutput     `pulumi:"failoverGroupId"`
	Kind                        pulumi.StringOutput     `pulumi:"kind"`
	Location                    pulumi.StringOutput     `pulumi:"location"`
	MaxSizeBytes                pulumi.Float64PtrOutput `pulumi:"maxSizeBytes"`
	Name                        pulumi.StringOutput     `pulumi:"name"`
	Sku                         SkuResponsePtrOutput    `pulumi:"sku"`
	Status                      pulumi.StringOutput     `pulumi:"status"`
	Tags                        pulumi.StringMapOutput  `pulumi:"tags"`
	Type                        pulumi.StringOutput     `pulumi:"type"`
	ZoneRedundant               pulumi.BoolPtrOutput    `pulumi:"zoneRedundant"`
}


func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:Database"),
		},
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azure-native:sql/v20170301preview:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:sql/v20170301preview:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseState struct {
}

type DatabaseState struct {
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	CatalogCollation                  *string           `pulumi:"catalogCollation"`
	Collation                         *string           `pulumi:"collation"`
	CreateMode                        *string           `pulumi:"createMode"`
	DatabaseName                      *string           `pulumi:"databaseName"`
	ElasticPoolId                     *string           `pulumi:"elasticPoolId"`
	Location                          *string           `pulumi:"location"`
	LongTermRetentionBackupResourceId *string           `pulumi:"longTermRetentionBackupResourceId"`
	MaxSizeBytes                      *float64          `pulumi:"maxSizeBytes"`
	RecoverableDatabaseId             *string           `pulumi:"recoverableDatabaseId"`
	RecoveryServicesRecoveryPointId   *string           `pulumi:"recoveryServicesRecoveryPointId"`
	ResourceGroupName                 string            `pulumi:"resourceGroupName"`
	RestorableDroppedDatabaseId       *string           `pulumi:"restorableDroppedDatabaseId"`
	RestorePointInTime                *string           `pulumi:"restorePointInTime"`
	SampleName                        *string           `pulumi:"sampleName"`
	ServerName                        string            `pulumi:"serverName"`
	Sku                               *Sku              `pulumi:"sku"`
	SourceDatabaseDeletionDate        *string           `pulumi:"sourceDatabaseDeletionDate"`
	SourceDatabaseId                  *string           `pulumi:"sourceDatabaseId"`
	Tags                              map[string]string `pulumi:"tags"`
	ZoneRedundant                     *bool             `pulumi:"zoneRedundant"`
}


type DatabaseArgs struct {
	CatalogCollation                  pulumi.StringPtrInput
	Collation                         pulumi.StringPtrInput
	CreateMode                        pulumi.StringPtrInput
	DatabaseName                      pulumi.StringPtrInput
	ElasticPoolId                     pulumi.StringPtrInput
	Location                          pulumi.StringPtrInput
	LongTermRetentionBackupResourceId pulumi.StringPtrInput
	MaxSizeBytes                      pulumi.Float64PtrInput
	RecoverableDatabaseId             pulumi.StringPtrInput
	RecoveryServicesRecoveryPointId   pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	RestorableDroppedDatabaseId       pulumi.StringPtrInput
	RestorePointInTime                pulumi.StringPtrInput
	SampleName                        pulumi.StringPtrInput
	ServerName                        pulumi.StringInput
	Sku                               SkuPtrInput
	SourceDatabaseDeletionDate        pulumi.StringPtrInput
	SourceDatabaseId                  pulumi.StringPtrInput
	Tags                              pulumi.StringMapInput
	ZoneRedundant                     pulumi.BoolPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func (o DatabaseOutput) CatalogCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.CatalogCollation }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o DatabaseOutput) CurrentServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CurrentServiceObjectiveName }).(pulumi.StringOutput)
}

func (o DatabaseOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

func (o DatabaseOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ElasticPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ElasticPoolId }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.FailoverGroupId }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DatabaseOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Database) pulumi.Float64PtrOutput { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Database) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o DatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Database) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.ZoneRedundant }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
