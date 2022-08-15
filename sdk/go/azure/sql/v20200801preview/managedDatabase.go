


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedDatabase struct {
	pulumi.CustomResourceState

	CatalogCollation         pulumi.StringPtrOutput `pulumi:"catalogCollation"`
	Collation                pulumi.StringPtrOutput `pulumi:"collation"`
	CreationDate             pulumi.StringOutput    `pulumi:"creationDate"`
	DefaultSecondaryLocation pulumi.StringOutput    `pulumi:"defaultSecondaryLocation"`
	EarliestRestorePoint     pulumi.StringOutput    `pulumi:"earliestRestorePoint"`
	FailoverGroupId          pulumi.StringOutput    `pulumi:"failoverGroupId"`
	Location                 pulumi.StringOutput    `pulumi:"location"`
	Name                     pulumi.StringOutput    `pulumi:"name"`
	Status                   pulumi.StringOutput    `pulumi:"status"`
	Tags                     pulumi.StringMapOutput `pulumi:"tags"`
	Type                     pulumi.StringOutput    `pulumi:"type"`
}


func NewManagedDatabase(ctx *pulumi.Context,
	name string, args *ManagedDatabaseArgs, opts ...pulumi.ResourceOption) (*ManagedDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedDatabase"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedDatabase
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:ManagedDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedDatabaseState, opts ...pulumi.ResourceOption) (*ManagedDatabase, error) {
	var resource ManagedDatabase
	err := ctx.ReadResource("azure-native:sql/v20200801preview:ManagedDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedDatabaseState struct {
}

type ManagedDatabaseState struct {
}

func (ManagedDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDatabaseState)(nil)).Elem()
}

type managedDatabaseArgs struct {
	AutoCompleteRestore               *bool             `pulumi:"autoCompleteRestore"`
	CatalogCollation                  *string           `pulumi:"catalogCollation"`
	Collation                         *string           `pulumi:"collation"`
	CreateMode                        *string           `pulumi:"createMode"`
	DatabaseName                      *string           `pulumi:"databaseName"`
	LastBackupName                    *string           `pulumi:"lastBackupName"`
	Location                          *string           `pulumi:"location"`
	LongTermRetentionBackupResourceId *string           `pulumi:"longTermRetentionBackupResourceId"`
	ManagedInstanceName               string            `pulumi:"managedInstanceName"`
	RecoverableDatabaseId             *string           `pulumi:"recoverableDatabaseId"`
	ResourceGroupName                 string            `pulumi:"resourceGroupName"`
	RestorableDroppedDatabaseId       *string           `pulumi:"restorableDroppedDatabaseId"`
	RestorePointInTime                *string           `pulumi:"restorePointInTime"`
	SourceDatabaseId                  *string           `pulumi:"sourceDatabaseId"`
	StorageContainerSasToken          *string           `pulumi:"storageContainerSasToken"`
	StorageContainerUri               *string           `pulumi:"storageContainerUri"`
	Tags                              map[string]string `pulumi:"tags"`
}


type ManagedDatabaseArgs struct {
	AutoCompleteRestore               pulumi.BoolPtrInput
	CatalogCollation                  pulumi.StringPtrInput
	Collation                         pulumi.StringPtrInput
	CreateMode                        pulumi.StringPtrInput
	DatabaseName                      pulumi.StringPtrInput
	LastBackupName                    pulumi.StringPtrInput
	Location                          pulumi.StringPtrInput
	LongTermRetentionBackupResourceId pulumi.StringPtrInput
	ManagedInstanceName               pulumi.StringInput
	RecoverableDatabaseId             pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	RestorableDroppedDatabaseId       pulumi.StringPtrInput
	RestorePointInTime                pulumi.StringPtrInput
	SourceDatabaseId                  pulumi.StringPtrInput
	StorageContainerSasToken          pulumi.StringPtrInput
	StorageContainerUri               pulumi.StringPtrInput
	Tags                              pulumi.StringMapInput
}

func (ManagedDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDatabaseArgs)(nil)).Elem()
}

type ManagedDatabaseInput interface {
	pulumi.Input

	ToManagedDatabaseOutput() ManagedDatabaseOutput
	ToManagedDatabaseOutputWithContext(ctx context.Context) ManagedDatabaseOutput
}

func (*ManagedDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabase)(nil)).Elem()
}

func (i *ManagedDatabase) ToManagedDatabaseOutput() ManagedDatabaseOutput {
	return i.ToManagedDatabaseOutputWithContext(context.Background())
}

func (i *ManagedDatabase) ToManagedDatabaseOutputWithContext(ctx context.Context) ManagedDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDatabaseOutput)
}

type ManagedDatabaseOutput struct{ *pulumi.OutputState }

func (ManagedDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabase)(nil)).Elem()
}

func (o ManagedDatabaseOutput) ToManagedDatabaseOutput() ManagedDatabaseOutput {
	return o
}

func (o ManagedDatabaseOutput) ToManagedDatabaseOutputWithContext(ctx context.Context) ManagedDatabaseOutput {
	return o
}

func (o ManagedDatabaseOutput) CatalogCollation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringPtrOutput { return v.CatalogCollation }).(pulumi.StringPtrOutput)
}

func (o ManagedDatabaseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o ManagedDatabaseOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o ManagedDatabaseOutput) DefaultSecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.DefaultSecondaryLocation }).(pulumi.StringOutput)
}

func (o ManagedDatabaseOutput) EarliestRestorePoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.EarliestRestorePoint }).(pulumi.StringOutput)
}

func (o ManagedDatabaseOutput) FailoverGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.FailoverGroupId }).(pulumi.StringOutput)
}

func (o ManagedDatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedDatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ManagedDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedDatabaseOutput{})
}
