


package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Database struct {
	pulumi.CustomResourceState

	AutoPauseDelay                   pulumi.IntPtrOutput     `pulumi:"autoPauseDelay"`
	CatalogCollation                 pulumi.StringPtrOutput  `pulumi:"catalogCollation"`
	Collation                        pulumi.StringPtrOutput  `pulumi:"collation"`
	CreationDate                     pulumi.StringOutput     `pulumi:"creationDate"`
	CurrentBackupStorageRedundancy   pulumi.StringOutput     `pulumi:"currentBackupStorageRedundancy"`
	CurrentServiceObjectiveName      pulumi.StringOutput     `pulumi:"currentServiceObjectiveName"`
	CurrentSku                       SkuResponseOutput       `pulumi:"currentSku"`
	DatabaseId                       pulumi.StringOutput     `pulumi:"databaseId"`
	DefaultSecondaryLocation         pulumi.StringOutput     `pulumi:"defaultSecondaryLocation"`
	EarliestRestoreDate              pulumi.StringOutput     `pulumi:"earliestRestoreDate"`
	ElasticPoolId                    pulumi.StringPtrOutput  `pulumi:"elasticPoolId"`
	FailoverGroupId                  pulumi.StringOutput     `pulumi:"failoverGroupId"`
	HighAvailabilityReplicaCount     pulumi.IntPtrOutput     `pulumi:"highAvailabilityReplicaCount"`
	Kind                             pulumi.StringOutput     `pulumi:"kind"`
	LicenseType                      pulumi.StringPtrOutput  `pulumi:"licenseType"`
	Location                         pulumi.StringOutput     `pulumi:"location"`
	MaintenanceConfigurationId       pulumi.StringPtrOutput  `pulumi:"maintenanceConfigurationId"`
	ManagedBy                        pulumi.StringOutput     `pulumi:"managedBy"`
	MaxLogSizeBytes                  pulumi.Float64Output    `pulumi:"maxLogSizeBytes"`
	MaxSizeBytes                     pulumi.Float64PtrOutput `pulumi:"maxSizeBytes"`
	MinCapacity                      pulumi.Float64PtrOutput `pulumi:"minCapacity"`
	Name                             pulumi.StringOutput     `pulumi:"name"`
	PausedDate                       pulumi.StringOutput     `pulumi:"pausedDate"`
	ReadScale                        pulumi.StringPtrOutput  `pulumi:"readScale"`
	RequestedBackupStorageRedundancy pulumi.StringPtrOutput  `pulumi:"requestedBackupStorageRedundancy"`
	RequestedServiceObjectiveName    pulumi.StringOutput     `pulumi:"requestedServiceObjectiveName"`
	ResumedDate                      pulumi.StringOutput     `pulumi:"resumedDate"`
	SecondaryType                    pulumi.StringPtrOutput  `pulumi:"secondaryType"`
	Sku                              SkuResponsePtrOutput    `pulumi:"sku"`
	Status                           pulumi.StringOutput     `pulumi:"status"`
	Tags                             pulumi.StringMapOutput  `pulumi:"tags"`
	Type                             pulumi.StringOutput     `pulumi:"type"`
	ZoneRedundant                    pulumi.BoolPtrOutput    `pulumi:"zoneRedundant"`
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
			Type: pulumi.String("azure-native:sql/v20140401:Database"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:Database"),
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
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azure-native:sql:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:sql:Database", name, id, state, &resource, opts...)
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
	AutoPauseDelay                    *int              `pulumi:"autoPauseDelay"`
	CatalogCollation                  *string           `pulumi:"catalogCollation"`
	Collation                         *string           `pulumi:"collation"`
	CreateMode                        *string           `pulumi:"createMode"`
	DatabaseName                      *string           `pulumi:"databaseName"`
	ElasticPoolId                     *string           `pulumi:"elasticPoolId"`
	HighAvailabilityReplicaCount      *int              `pulumi:"highAvailabilityReplicaCount"`
	LicenseType                       *string           `pulumi:"licenseType"`
	Location                          *string           `pulumi:"location"`
	LongTermRetentionBackupResourceId *string           `pulumi:"longTermRetentionBackupResourceId"`
	MaintenanceConfigurationId        *string           `pulumi:"maintenanceConfigurationId"`
	MaxSizeBytes                      *float64          `pulumi:"maxSizeBytes"`
	MinCapacity                       *float64          `pulumi:"minCapacity"`
	ReadScale                         *string           `pulumi:"readScale"`
	RecoverableDatabaseId             *string           `pulumi:"recoverableDatabaseId"`
	RecoveryServicesRecoveryPointId   *string           `pulumi:"recoveryServicesRecoveryPointId"`
	RequestedBackupStorageRedundancy  *string           `pulumi:"requestedBackupStorageRedundancy"`
	ResourceGroupName                 string            `pulumi:"resourceGroupName"`
	RestorableDroppedDatabaseId       *string           `pulumi:"restorableDroppedDatabaseId"`
	RestorePointInTime                *string           `pulumi:"restorePointInTime"`
	SampleName                        *string           `pulumi:"sampleName"`
	SecondaryType                     *string           `pulumi:"secondaryType"`
	ServerName                        string            `pulumi:"serverName"`
	Sku                               *Sku              `pulumi:"sku"`
	SourceDatabaseDeletionDate        *string           `pulumi:"sourceDatabaseDeletionDate"`
	SourceDatabaseId                  *string           `pulumi:"sourceDatabaseId"`
	Tags                              map[string]string `pulumi:"tags"`
	ZoneRedundant                     *bool             `pulumi:"zoneRedundant"`
}


type DatabaseArgs struct {
	AutoPauseDelay                    pulumi.IntPtrInput
	CatalogCollation                  pulumi.StringPtrInput
	Collation                         pulumi.StringPtrInput
	CreateMode                        pulumi.StringPtrInput
	DatabaseName                      pulumi.StringPtrInput
	ElasticPoolId                     pulumi.StringPtrInput
	HighAvailabilityReplicaCount      pulumi.IntPtrInput
	LicenseType                       pulumi.StringPtrInput
	Location                          pulumi.StringPtrInput
	LongTermRetentionBackupResourceId pulumi.StringPtrInput
	MaintenanceConfigurationId        pulumi.StringPtrInput
	MaxSizeBytes                      pulumi.Float64PtrInput
	MinCapacity                       pulumi.Float64PtrInput
	ReadScale                         pulumi.StringPtrInput
	RecoverableDatabaseId             pulumi.StringPtrInput
	RecoveryServicesRecoveryPointId   pulumi.StringPtrInput
	RequestedBackupStorageRedundancy  pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	RestorableDroppedDatabaseId       pulumi.StringPtrInput
	RestorePointInTime                pulumi.StringPtrInput
	SampleName                        pulumi.StringPtrInput
	SecondaryType                     pulumi.StringPtrInput
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

func (o DatabaseOutput) AutoPauseDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.IntPtrOutput { return v.AutoPauseDelay }).(pulumi.IntPtrOutput)
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

func (o DatabaseOutput) CurrentBackupStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CurrentBackupStorageRedundancy }).(pulumi.StringOutput)
}

func (o DatabaseOutput) CurrentServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CurrentServiceObjectiveName }).(pulumi.StringOutput)
}

func (o DatabaseOutput) CurrentSku() SkuResponseOutput {
	return o.ApplyT(func(v *Database) SkuResponseOutput { return v.CurrentSku }).(SkuResponseOutput)
}

func (o DatabaseOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

func (o DatabaseOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

func (o DatabaseOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ElasticPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ElasticPoolId }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.FailoverGroupId }).(pulumi.StringOutput)
}

func (o DatabaseOutput) HighAvailabilityReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.IntPtrOutput { return v.HighAvailabilityReplicaCount }).(pulumi.IntPtrOutput)
}

func (o DatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DatabaseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DatabaseOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o DatabaseOutput) MaxLogSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Database) pulumi.Float64Output { return v.MaxLogSizeBytes }).(pulumi.Float64Output)
}

func (o DatabaseOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Database) pulumi.Float64PtrOutput { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o DatabaseOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Database) pulumi.Float64PtrOutput { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseOutput) PausedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.PausedDate }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ReadScale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ReadScale }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) RequestedBackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.RequestedBackupStorageRedundancy }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) RequestedServiceObjectiveName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.RequestedServiceObjectiveName }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ResumedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ResumedDate }).(pulumi.StringOutput)
}

func (o DatabaseOutput) SecondaryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.SecondaryType }).(pulumi.StringPtrOutput)
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
