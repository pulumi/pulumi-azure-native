


package v20140401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Database struct {
	pulumi.CustomResourceState

	Collation                     pulumi.StringPtrOutput                       `pulumi:"collation"`
	ContainmentState              pulumi.Float64Output                         `pulumi:"containmentState"`
	CreationDate                  pulumi.StringOutput                          `pulumi:"creationDate"`
	CurrentServiceObjectiveId     pulumi.StringOutput                          `pulumi:"currentServiceObjectiveId"`
	DatabaseId                    pulumi.StringOutput                          `pulumi:"databaseId"`
	DefaultSecondaryLocation      pulumi.StringOutput                          `pulumi:"defaultSecondaryLocation"`
	EarliestRestoreDate           pulumi.StringOutput                          `pulumi:"earliestRestoreDate"`
	Edition                       pulumi.StringPtrOutput                       `pulumi:"edition"`
	ElasticPoolName               pulumi.StringPtrOutput                       `pulumi:"elasticPoolName"`
	FailoverGroupId               pulumi.StringOutput                          `pulumi:"failoverGroupId"`
	Kind                          pulumi.StringOutput                          `pulumi:"kind"`
	Location                      pulumi.StringOutput                          `pulumi:"location"`
	MaxSizeBytes                  pulumi.StringPtrOutput                       `pulumi:"maxSizeBytes"`
	Name                          pulumi.StringOutput                          `pulumi:"name"`
	ReadScale                     pulumi.StringPtrOutput                       `pulumi:"readScale"`
	RecommendedIndex              RecommendedIndexResponseArrayOutput          `pulumi:"recommendedIndex"`
	RequestedServiceObjectiveId   pulumi.StringPtrOutput                       `pulumi:"requestedServiceObjectiveId"`
	RequestedServiceObjectiveName pulumi.StringPtrOutput                       `pulumi:"requestedServiceObjectiveName"`
	ServiceLevelObjective         pulumi.StringOutput                          `pulumi:"serviceLevelObjective"`
	ServiceTierAdvisors           ServiceTierAdvisorResponseArrayOutput        `pulumi:"serviceTierAdvisors"`
	Status                        pulumi.StringOutput                          `pulumi:"status"`
	Tags                          pulumi.StringMapOutput                       `pulumi:"tags"`
	TransparentDataEncryption     TransparentDataEncryptionResponseArrayOutput `pulumi:"transparentDataEncryption"`
	Type                          pulumi.StringOutput                          `pulumi:"type"`
	ZoneRedundant                 pulumi.BoolPtrOutput                         `pulumi:"zoneRedundant"`
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
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:Database"),
		},
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azure-native:sql/v20140401:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:sql/v20140401:Database", name, id, state, &resource, opts...)
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
	Collation                               *string           `pulumi:"collation"`
	CreateMode                              *string           `pulumi:"createMode"`
	DatabaseName                            *string           `pulumi:"databaseName"`
	Edition                                 *string           `pulumi:"edition"`
	ElasticPoolName                         *string           `pulumi:"elasticPoolName"`
	Location                                *string           `pulumi:"location"`
	MaxSizeBytes                            *string           `pulumi:"maxSizeBytes"`
	ReadScale                               *string           `pulumi:"readScale"`
	RecoveryServicesRecoveryPointResourceId *string           `pulumi:"recoveryServicesRecoveryPointResourceId"`
	RequestedServiceObjectiveId             *string           `pulumi:"requestedServiceObjectiveId"`
	RequestedServiceObjectiveName           *string           `pulumi:"requestedServiceObjectiveName"`
	ResourceGroupName                       string            `pulumi:"resourceGroupName"`
	RestorePointInTime                      *string           `pulumi:"restorePointInTime"`
	SampleName                              *string           `pulumi:"sampleName"`
	ServerName                              string            `pulumi:"serverName"`
	SourceDatabaseDeletionDate              *string           `pulumi:"sourceDatabaseDeletionDate"`
	SourceDatabaseId                        *string           `pulumi:"sourceDatabaseId"`
	Tags                                    map[string]string `pulumi:"tags"`
	ZoneRedundant                           *bool             `pulumi:"zoneRedundant"`
}


type DatabaseArgs struct {
	Collation                               pulumi.StringPtrInput
	CreateMode                              pulumi.StringPtrInput
	DatabaseName                            pulumi.StringPtrInput
	Edition                                 pulumi.StringPtrInput
	ElasticPoolName                         pulumi.StringPtrInput
	Location                                pulumi.StringPtrInput
	MaxSizeBytes                            pulumi.StringPtrInput
	ReadScale                               pulumi.StringPtrInput
	RecoveryServicesRecoveryPointResourceId pulumi.StringPtrInput
	RequestedServiceObjectiveId             pulumi.StringPtrInput
	RequestedServiceObjectiveName           pulumi.StringPtrInput
	ResourceGroupName                       pulumi.StringInput
	RestorePointInTime                      pulumi.StringPtrInput
	SampleName                              pulumi.StringPtrInput
	ServerName                              pulumi.StringInput
	SourceDatabaseDeletionDate              pulumi.StringPtrInput
	SourceDatabaseId                        pulumi.StringPtrInput
	Tags                                    pulumi.StringMapInput
	ZoneRedundant                           pulumi.BoolPtrInput
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

func (o DatabaseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) ContainmentState() pulumi.Float64Output {
	return o.ApplyT(func(v *Database) pulumi.Float64Output { return v.ContainmentState }).(pulumi.Float64Output)
}

func (o DatabaseOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o DatabaseOutput) CurrentServiceObjectiveId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CurrentServiceObjectiveId }).(pulumi.StringOutput)
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

func (o DatabaseOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Edition }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) ElasticPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ElasticPoolName }).(pulumi.StringPtrOutput)
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

func (o DatabaseOutput) MaxSizeBytes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.MaxSizeBytes }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ReadScale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ReadScale }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) RecommendedIndex() RecommendedIndexResponseArrayOutput {
	return o.ApplyT(func(v *Database) RecommendedIndexResponseArrayOutput { return v.RecommendedIndex }).(RecommendedIndexResponseArrayOutput)
}

func (o DatabaseOutput) RequestedServiceObjectiveId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.RequestedServiceObjectiveId }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) RequestedServiceObjectiveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.RequestedServiceObjectiveName }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) ServiceLevelObjective() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ServiceLevelObjective }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ServiceTierAdvisors() ServiceTierAdvisorResponseArrayOutput {
	return o.ApplyT(func(v *Database) ServiceTierAdvisorResponseArrayOutput { return v.ServiceTierAdvisors }).(ServiceTierAdvisorResponseArrayOutput)
}

func (o DatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Database) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseOutput) TransparentDataEncryption() TransparentDataEncryptionResponseArrayOutput {
	return o.ApplyT(func(v *Database) TransparentDataEncryptionResponseArrayOutput { return v.TransparentDataEncryption }).(TransparentDataEncryptionResponseArrayOutput)
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
