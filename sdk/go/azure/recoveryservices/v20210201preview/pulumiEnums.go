


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupItemType string

const (
	BackupItemTypeInvalid           = BackupItemType("Invalid")
	BackupItemTypeVM                = BackupItemType("VM")
	BackupItemTypeFileFolder        = BackupItemType("FileFolder")
	BackupItemTypeAzureSqlDb        = BackupItemType("AzureSqlDb")
	BackupItemTypeSQLDB             = BackupItemType("SQLDB")
	BackupItemTypeExchange          = BackupItemType("Exchange")
	BackupItemTypeSharepoint        = BackupItemType("Sharepoint")
	BackupItemTypeVMwareVM          = BackupItemType("VMwareVM")
	BackupItemTypeSystemState       = BackupItemType("SystemState")
	BackupItemTypeClient            = BackupItemType("Client")
	BackupItemTypeGenericDataSource = BackupItemType("GenericDataSource")
	BackupItemTypeSQLDataBase       = BackupItemType("SQLDataBase")
	BackupItemTypeAzureFileShare    = BackupItemType("AzureFileShare")
	BackupItemTypeSAPHanaDatabase   = BackupItemType("SAPHanaDatabase")
	BackupItemTypeSAPAseDatabase    = BackupItemType("SAPAseDatabase")
)

func (BackupItemType) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupItemType)(nil)).Elem()
}

func (e BackupItemType) ToBackupItemTypeOutput() BackupItemTypeOutput {
	return pulumi.ToOutput(e).(BackupItemTypeOutput)
}

func (e BackupItemType) ToBackupItemTypeOutputWithContext(ctx context.Context) BackupItemTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackupItemTypeOutput)
}

func (e BackupItemType) ToBackupItemTypePtrOutput() BackupItemTypePtrOutput {
	return e.ToBackupItemTypePtrOutputWithContext(context.Background())
}

func (e BackupItemType) ToBackupItemTypePtrOutputWithContext(ctx context.Context) BackupItemTypePtrOutput {
	return BackupItemType(e).ToBackupItemTypeOutputWithContext(ctx).ToBackupItemTypePtrOutputWithContext(ctx)
}

func (e BackupItemType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupItemType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupItemType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupItemType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackupItemTypeOutput struct{ *pulumi.OutputState }

func (BackupItemTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupItemType)(nil)).Elem()
}

func (o BackupItemTypeOutput) ToBackupItemTypeOutput() BackupItemTypeOutput {
	return o
}

func (o BackupItemTypeOutput) ToBackupItemTypeOutputWithContext(ctx context.Context) BackupItemTypeOutput {
	return o
}

func (o BackupItemTypeOutput) ToBackupItemTypePtrOutput() BackupItemTypePtrOutput {
	return o.ToBackupItemTypePtrOutputWithContext(context.Background())
}

func (o BackupItemTypeOutput) ToBackupItemTypePtrOutputWithContext(ctx context.Context) BackupItemTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupItemType) *BackupItemType {
		return &v
	}).(BackupItemTypePtrOutput)
}

func (o BackupItemTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackupItemTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupItemType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackupItemTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupItemTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupItemType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackupItemTypePtrOutput struct{ *pulumi.OutputState }

func (BackupItemTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupItemType)(nil)).Elem()
}

func (o BackupItemTypePtrOutput) ToBackupItemTypePtrOutput() BackupItemTypePtrOutput {
	return o
}

func (o BackupItemTypePtrOutput) ToBackupItemTypePtrOutputWithContext(ctx context.Context) BackupItemTypePtrOutput {
	return o
}

func (o BackupItemTypePtrOutput) Elem() BackupItemTypeOutput {
	return o.ApplyT(func(v *BackupItemType) BackupItemType {
		if v != nil {
			return *v
		}
		var ret BackupItemType
		return ret
	}).(BackupItemTypeOutput)
}

func (o BackupItemTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupItemTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackupItemType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackupItemTypeInput interface {
	pulumi.Input

	ToBackupItemTypeOutput() BackupItemTypeOutput
	ToBackupItemTypeOutputWithContext(context.Context) BackupItemTypeOutput
}

var backupItemTypePtrType = reflect.TypeOf((**BackupItemType)(nil)).Elem()

type BackupItemTypePtrInput interface {
	pulumi.Input

	ToBackupItemTypePtrOutput() BackupItemTypePtrOutput
	ToBackupItemTypePtrOutputWithContext(context.Context) BackupItemTypePtrOutput
}

type backupItemTypePtr string

func BackupItemTypePtr(v string) BackupItemTypePtrInput {
	return (*backupItemTypePtr)(&v)
}

func (*backupItemTypePtr) ElementType() reflect.Type {
	return backupItemTypePtrType
}

func (in *backupItemTypePtr) ToBackupItemTypePtrOutput() BackupItemTypePtrOutput {
	return pulumi.ToOutput(in).(BackupItemTypePtrOutput)
}

func (in *backupItemTypePtr) ToBackupItemTypePtrOutputWithContext(ctx context.Context) BackupItemTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackupItemTypePtrOutput)
}

type BackupManagementType string

const (
	BackupManagementTypeInvalid           = BackupManagementType("Invalid")
	BackupManagementTypeAzureIaasVM       = BackupManagementType("AzureIaasVM")
	BackupManagementTypeMAB               = BackupManagementType("MAB")
	BackupManagementTypeDPM               = BackupManagementType("DPM")
	BackupManagementTypeAzureBackupServer = BackupManagementType("AzureBackupServer")
	BackupManagementTypeAzureSql          = BackupManagementType("AzureSql")
	BackupManagementTypeAzureStorage      = BackupManagementType("AzureStorage")
	BackupManagementTypeAzureWorkload     = BackupManagementType("AzureWorkload")
	BackupManagementTypeDefaultBackup     = BackupManagementType("DefaultBackup")
)

func (BackupManagementType) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupManagementType)(nil)).Elem()
}

func (e BackupManagementType) ToBackupManagementTypeOutput() BackupManagementTypeOutput {
	return pulumi.ToOutput(e).(BackupManagementTypeOutput)
}

func (e BackupManagementType) ToBackupManagementTypeOutputWithContext(ctx context.Context) BackupManagementTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackupManagementTypeOutput)
}

func (e BackupManagementType) ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput {
	return e.ToBackupManagementTypePtrOutputWithContext(context.Background())
}

func (e BackupManagementType) ToBackupManagementTypePtrOutputWithContext(ctx context.Context) BackupManagementTypePtrOutput {
	return BackupManagementType(e).ToBackupManagementTypeOutputWithContext(ctx).ToBackupManagementTypePtrOutputWithContext(ctx)
}

func (e BackupManagementType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupManagementType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupManagementType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupManagementType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackupManagementTypeOutput struct{ *pulumi.OutputState }

func (BackupManagementTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupManagementType)(nil)).Elem()
}

func (o BackupManagementTypeOutput) ToBackupManagementTypeOutput() BackupManagementTypeOutput {
	return o
}

func (o BackupManagementTypeOutput) ToBackupManagementTypeOutputWithContext(ctx context.Context) BackupManagementTypeOutput {
	return o
}

func (o BackupManagementTypeOutput) ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput {
	return o.ToBackupManagementTypePtrOutputWithContext(context.Background())
}

func (o BackupManagementTypeOutput) ToBackupManagementTypePtrOutputWithContext(ctx context.Context) BackupManagementTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupManagementType) *BackupManagementType {
		return &v
	}).(BackupManagementTypePtrOutput)
}

func (o BackupManagementTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackupManagementTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupManagementType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackupManagementTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupManagementTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupManagementType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackupManagementTypePtrOutput struct{ *pulumi.OutputState }

func (BackupManagementTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupManagementType)(nil)).Elem()
}

func (o BackupManagementTypePtrOutput) ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput {
	return o
}

func (o BackupManagementTypePtrOutput) ToBackupManagementTypePtrOutputWithContext(ctx context.Context) BackupManagementTypePtrOutput {
	return o
}

func (o BackupManagementTypePtrOutput) Elem() BackupManagementTypeOutput {
	return o.ApplyT(func(v *BackupManagementType) BackupManagementType {
		if v != nil {
			return *v
		}
		var ret BackupManagementType
		return ret
	}).(BackupManagementTypeOutput)
}

func (o BackupManagementTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupManagementTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackupManagementType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackupManagementTypeInput interface {
	pulumi.Input

	ToBackupManagementTypeOutput() BackupManagementTypeOutput
	ToBackupManagementTypeOutputWithContext(context.Context) BackupManagementTypeOutput
}

var backupManagementTypePtrType = reflect.TypeOf((**BackupManagementType)(nil)).Elem()

type BackupManagementTypePtrInput interface {
	pulumi.Input

	ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput
	ToBackupManagementTypePtrOutputWithContext(context.Context) BackupManagementTypePtrOutput
}

type backupManagementTypePtr string

func BackupManagementTypePtr(v string) BackupManagementTypePtrInput {
	return (*backupManagementTypePtr)(&v)
}

func (*backupManagementTypePtr) ElementType() reflect.Type {
	return backupManagementTypePtrType
}

func (in *backupManagementTypePtr) ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput {
	return pulumi.ToOutput(in).(BackupManagementTypePtrOutput)
}

func (in *backupManagementTypePtr) ToBackupManagementTypePtrOutputWithContext(ctx context.Context) BackupManagementTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackupManagementTypePtrOutput)
}

type ContainerType string

const (
	ContainerTypeInvalid                    = ContainerType("Invalid")
	ContainerTypeUnknown                    = ContainerType("Unknown")
	ContainerTypeIaasVMContainer            = ContainerType("IaasVMContainer")
	ContainerTypeIaasVMServiceContainer     = ContainerType("IaasVMServiceContainer")
	ContainerTypeDPMContainer               = ContainerType("DPMContainer")
	ContainerTypeAzureBackupServerContainer = ContainerType("AzureBackupServerContainer")
	ContainerTypeMABContainer               = ContainerType("MABContainer")
	ContainerTypeCluster                    = ContainerType("Cluster")
	ContainerTypeAzureSqlContainer          = ContainerType("AzureSqlContainer")
	ContainerTypeWindows                    = ContainerType("Windows")
	ContainerTypeVCenter                    = ContainerType("VCenter")
	ContainerTypeVMAppContainer             = ContainerType("VMAppContainer")
	ContainerTypeSQLAGWorkLoadContainer     = ContainerType("SQLAGWorkLoadContainer")
	ContainerTypeStorageContainer           = ContainerType("StorageContainer")
	ContainerTypeGenericContainer           = ContainerType("GenericContainer")
)

func (ContainerType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerType)(nil)).Elem()
}

func (e ContainerType) ToContainerTypeOutput() ContainerTypeOutput {
	return pulumi.ToOutput(e).(ContainerTypeOutput)
}

func (e ContainerType) ToContainerTypeOutputWithContext(ctx context.Context) ContainerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerTypeOutput)
}

func (e ContainerType) ToContainerTypePtrOutput() ContainerTypePtrOutput {
	return e.ToContainerTypePtrOutputWithContext(context.Background())
}

func (e ContainerType) ToContainerTypePtrOutputWithContext(ctx context.Context) ContainerTypePtrOutput {
	return ContainerType(e).ToContainerTypeOutputWithContext(ctx).ToContainerTypePtrOutputWithContext(ctx)
}

func (e ContainerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerTypeOutput struct{ *pulumi.OutputState }

func (ContainerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerType)(nil)).Elem()
}

func (o ContainerTypeOutput) ToContainerTypeOutput() ContainerTypeOutput {
	return o
}

func (o ContainerTypeOutput) ToContainerTypeOutputWithContext(ctx context.Context) ContainerTypeOutput {
	return o
}

func (o ContainerTypeOutput) ToContainerTypePtrOutput() ContainerTypePtrOutput {
	return o.ToContainerTypePtrOutputWithContext(context.Background())
}

func (o ContainerTypeOutput) ToContainerTypePtrOutputWithContext(ctx context.Context) ContainerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerType) *ContainerType {
		return &v
	}).(ContainerTypePtrOutput)
}

func (o ContainerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerTypePtrOutput struct{ *pulumi.OutputState }

func (ContainerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerType)(nil)).Elem()
}

func (o ContainerTypePtrOutput) ToContainerTypePtrOutput() ContainerTypePtrOutput {
	return o
}

func (o ContainerTypePtrOutput) ToContainerTypePtrOutputWithContext(ctx context.Context) ContainerTypePtrOutput {
	return o
}

func (o ContainerTypePtrOutput) Elem() ContainerTypeOutput {
	return o.ApplyT(func(v *ContainerType) ContainerType {
		if v != nil {
			return *v
		}
		var ret ContainerType
		return ret
	}).(ContainerTypeOutput)
}

func (o ContainerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerTypeInput interface {
	pulumi.Input

	ToContainerTypeOutput() ContainerTypeOutput
	ToContainerTypeOutputWithContext(context.Context) ContainerTypeOutput
}

var containerTypePtrType = reflect.TypeOf((**ContainerType)(nil)).Elem()

type ContainerTypePtrInput interface {
	pulumi.Input

	ToContainerTypePtrOutput() ContainerTypePtrOutput
	ToContainerTypePtrOutputWithContext(context.Context) ContainerTypePtrOutput
}

type containerTypePtr string

func ContainerTypePtr(v string) ContainerTypePtrInput {
	return (*containerTypePtr)(&v)
}

func (*containerTypePtr) ElementType() reflect.Type {
	return containerTypePtrType
}

func (in *containerTypePtr) ToContainerTypePtrOutput() ContainerTypePtrOutput {
	return pulumi.ToOutput(in).(ContainerTypePtrOutput)
}

func (in *containerTypePtr) ToContainerTypePtrOutputWithContext(ctx context.Context) ContainerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerTypePtrOutput)
}

type CreateMode string

const (
	CreateModeInvalid = CreateMode("Invalid")
	CreateModeDefault = CreateMode("Default")
	CreateModeRecover = CreateMode("Recover")
)

func (CreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (e CreateMode) ToCreateModeOutput() CreateModeOutput {
	return pulumi.ToOutput(e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModePtrOutput() CreateModePtrOutput {
	return e.ToCreateModePtrOutputWithContext(context.Background())
}

func (e CreateMode) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return CreateMode(e).ToCreateModeOutputWithContext(ctx).ToCreateModePtrOutputWithContext(ctx)
}

func (e CreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CreateModeOutput struct{ *pulumi.OutputState }

func (CreateModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (o CreateModeOutput) ToCreateModeOutput() CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o.ToCreateModePtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateMode) *CreateMode {
		return &v
	}).(CreateModePtrOutput)
}

func (o CreateModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CreateModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CreateModePtrOutput struct{ *pulumi.OutputState }

func (CreateModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateMode)(nil)).Elem()
}

func (o CreateModePtrOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) Elem() CreateModeOutput {
	return o.ApplyT(func(v *CreateMode) CreateMode {
		if v != nil {
			return *v
		}
		var ret CreateMode
		return ret
	}).(CreateModeOutput)
}

func (o CreateModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CreateMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CreateModeInput interface {
	pulumi.Input

	ToCreateModeOutput() CreateModeOutput
	ToCreateModeOutputWithContext(context.Context) CreateModeOutput
}

var createModePtrType = reflect.TypeOf((**CreateMode)(nil)).Elem()

type CreateModePtrInput interface {
	pulumi.Input

	ToCreateModePtrOutput() CreateModePtrOutput
	ToCreateModePtrOutputWithContext(context.Context) CreateModePtrOutput
}

type createModePtr string

func CreateModePtr(v string) CreateModePtrInput {
	return (*createModePtr)(&v)
}

func (*createModePtr) ElementType() reflect.Type {
	return createModePtrType
}

func (in *createModePtr) ToCreateModePtrOutput() CreateModePtrOutput {
	return pulumi.ToOutput(in).(CreateModePtrOutput)
}

func (in *createModePtr) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CreateModePtrOutput)
}

type DataSourceType string

const (
	DataSourceTypeInvalid           = DataSourceType("Invalid")
	DataSourceTypeVM                = DataSourceType("VM")
	DataSourceTypeFileFolder        = DataSourceType("FileFolder")
	DataSourceTypeAzureSqlDb        = DataSourceType("AzureSqlDb")
	DataSourceTypeSQLDB             = DataSourceType("SQLDB")
	DataSourceTypeExchange          = DataSourceType("Exchange")
	DataSourceTypeSharepoint        = DataSourceType("Sharepoint")
	DataSourceTypeVMwareVM          = DataSourceType("VMwareVM")
	DataSourceTypeSystemState       = DataSourceType("SystemState")
	DataSourceTypeClient            = DataSourceType("Client")
	DataSourceTypeGenericDataSource = DataSourceType("GenericDataSource")
	DataSourceTypeSQLDataBase       = DataSourceType("SQLDataBase")
	DataSourceTypeAzureFileShare    = DataSourceType("AzureFileShare")
	DataSourceTypeSAPHanaDatabase   = DataSourceType("SAPHanaDatabase")
	DataSourceTypeSAPAseDatabase    = DataSourceType("SAPAseDatabase")
)

func (DataSourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceType)(nil)).Elem()
}

func (e DataSourceType) ToDataSourceTypeOutput() DataSourceTypeOutput {
	return pulumi.ToOutput(e).(DataSourceTypeOutput)
}

func (e DataSourceType) ToDataSourceTypeOutputWithContext(ctx context.Context) DataSourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataSourceTypeOutput)
}

func (e DataSourceType) ToDataSourceTypePtrOutput() DataSourceTypePtrOutput {
	return e.ToDataSourceTypePtrOutputWithContext(context.Background())
}

func (e DataSourceType) ToDataSourceTypePtrOutputWithContext(ctx context.Context) DataSourceTypePtrOutput {
	return DataSourceType(e).ToDataSourceTypeOutputWithContext(ctx).ToDataSourceTypePtrOutputWithContext(ctx)
}

func (e DataSourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataSourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataSourceTypeOutput struct{ *pulumi.OutputState }

func (DataSourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceType)(nil)).Elem()
}

func (o DataSourceTypeOutput) ToDataSourceTypeOutput() DataSourceTypeOutput {
	return o
}

func (o DataSourceTypeOutput) ToDataSourceTypeOutputWithContext(ctx context.Context) DataSourceTypeOutput {
	return o
}

func (o DataSourceTypeOutput) ToDataSourceTypePtrOutput() DataSourceTypePtrOutput {
	return o.ToDataSourceTypePtrOutputWithContext(context.Background())
}

func (o DataSourceTypeOutput) ToDataSourceTypePtrOutputWithContext(ctx context.Context) DataSourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataSourceType) *DataSourceType {
		return &v
	}).(DataSourceTypePtrOutput)
}

func (o DataSourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataSourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataSourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataSourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataSourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataSourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataSourceTypePtrOutput struct{ *pulumi.OutputState }

func (DataSourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourceType)(nil)).Elem()
}

func (o DataSourceTypePtrOutput) ToDataSourceTypePtrOutput() DataSourceTypePtrOutput {
	return o
}

func (o DataSourceTypePtrOutput) ToDataSourceTypePtrOutputWithContext(ctx context.Context) DataSourceTypePtrOutput {
	return o
}

func (o DataSourceTypePtrOutput) Elem() DataSourceTypeOutput {
	return o.ApplyT(func(v *DataSourceType) DataSourceType {
		if v != nil {
			return *v
		}
		var ret DataSourceType
		return ret
	}).(DataSourceTypeOutput)
}

func (o DataSourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataSourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataSourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataSourceTypeInput interface {
	pulumi.Input

	ToDataSourceTypeOutput() DataSourceTypeOutput
	ToDataSourceTypeOutputWithContext(context.Context) DataSourceTypeOutput
}

var dataSourceTypePtrType = reflect.TypeOf((**DataSourceType)(nil)).Elem()

type DataSourceTypePtrInput interface {
	pulumi.Input

	ToDataSourceTypePtrOutput() DataSourceTypePtrOutput
	ToDataSourceTypePtrOutputWithContext(context.Context) DataSourceTypePtrOutput
}

type dataSourceTypePtr string

func DataSourceTypePtr(v string) DataSourceTypePtrInput {
	return (*dataSourceTypePtr)(&v)
}

func (*dataSourceTypePtr) ElementType() reflect.Type {
	return dataSourceTypePtrType
}

func (in *dataSourceTypePtr) ToDataSourceTypePtrOutput() DataSourceTypePtrOutput {
	return pulumi.ToOutput(in).(DataSourceTypePtrOutput)
}

func (in *dataSourceTypePtr) ToDataSourceTypePtrOutputWithContext(ctx context.Context) DataSourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataSourceTypePtrOutput)
}

type DayOfWeek string

const (
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
)

func (DayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (e DayOfWeek) ToDayOfWeekOutput() DayOfWeekOutput {
	return pulumi.ToOutput(e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return e.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return DayOfWeek(e).ToDayOfWeekOutputWithContext(ctx).ToDayOfWeekPtrOutputWithContext(ctx)
}

func (e DayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DayOfWeekOutput struct{ *pulumi.OutputState }

func (DayOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekOutput) ToDayOfWeekOutput() DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DayOfWeek) *DayOfWeek {
		return &v
	}).(DayOfWeekPtrOutput)
}

func (o DayOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DayOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DayOfWeekPtrOutput struct{ *pulumi.OutputState }

func (DayOfWeekPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) Elem() DayOfWeekOutput {
	return o.ApplyT(func(v *DayOfWeek) DayOfWeek {
		if v != nil {
			return *v
		}
		var ret DayOfWeek
		return ret
	}).(DayOfWeekOutput)
}

func (o DayOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DayOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DayOfWeekInput interface {
	pulumi.Input

	ToDayOfWeekOutput() DayOfWeekOutput
	ToDayOfWeekOutputWithContext(context.Context) DayOfWeekOutput
}

var dayOfWeekPtrType = reflect.TypeOf((**DayOfWeek)(nil)).Elem()

type DayOfWeekPtrInput interface {
	pulumi.Input

	ToDayOfWeekPtrOutput() DayOfWeekPtrOutput
	ToDayOfWeekPtrOutputWithContext(context.Context) DayOfWeekPtrOutput
}

type dayOfWeekPtr string

func DayOfWeekPtr(v string) DayOfWeekPtrInput {
	return (*dayOfWeekPtr)(&v)
}

func (*dayOfWeekPtr) ElementType() reflect.Type {
	return dayOfWeekPtrType
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return pulumi.ToOutput(in).(DayOfWeekPtrOutput)
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DayOfWeekPtrOutput)
}





type DayOfWeekArrayInput interface {
	pulumi.Input

	ToDayOfWeekArrayOutput() DayOfWeekArrayOutput
	ToDayOfWeekArrayOutputWithContext(context.Context) DayOfWeekArrayOutput
}

type DayOfWeekArray []DayOfWeek

func (DayOfWeekArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayOfWeek)(nil)).Elem()
}

func (i DayOfWeekArray) ToDayOfWeekArrayOutput() DayOfWeekArrayOutput {
	return i.ToDayOfWeekArrayOutputWithContext(context.Background())
}

func (i DayOfWeekArray) ToDayOfWeekArrayOutputWithContext(ctx context.Context) DayOfWeekArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayOfWeekArrayOutput)
}

type DayOfWeekArrayOutput struct{ *pulumi.OutputState }

func (DayOfWeekArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekArrayOutput) ToDayOfWeekArrayOutput() DayOfWeekArrayOutput {
	return o
}

func (o DayOfWeekArrayOutput) ToDayOfWeekArrayOutputWithContext(ctx context.Context) DayOfWeekArrayOutput {
	return o
}

func (o DayOfWeekArrayOutput) Index(i pulumi.IntInput) DayOfWeekOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DayOfWeek {
		return vs[0].([]DayOfWeek)[vs[1].(int)]
	}).(DayOfWeekOutput)
}

type HealthStatus string

const (
	HealthStatusPassed          = HealthStatus("Passed")
	HealthStatusActionRequired  = HealthStatus("ActionRequired")
	HealthStatusActionSuggested = HealthStatus("ActionSuggested")
	HealthStatusInvalid         = HealthStatus("Invalid")
)

func (HealthStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthStatus)(nil)).Elem()
}

func (e HealthStatus) ToHealthStatusOutput() HealthStatusOutput {
	return pulumi.ToOutput(e).(HealthStatusOutput)
}

func (e HealthStatus) ToHealthStatusOutputWithContext(ctx context.Context) HealthStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HealthStatusOutput)
}

func (e HealthStatus) ToHealthStatusPtrOutput() HealthStatusPtrOutput {
	return e.ToHealthStatusPtrOutputWithContext(context.Background())
}

func (e HealthStatus) ToHealthStatusPtrOutputWithContext(ctx context.Context) HealthStatusPtrOutput {
	return HealthStatus(e).ToHealthStatusOutputWithContext(ctx).ToHealthStatusPtrOutputWithContext(ctx)
}

func (e HealthStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HealthStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HealthStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HealthStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HealthStatusOutput struct{ *pulumi.OutputState }

func (HealthStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthStatus)(nil)).Elem()
}

func (o HealthStatusOutput) ToHealthStatusOutput() HealthStatusOutput {
	return o
}

func (o HealthStatusOutput) ToHealthStatusOutputWithContext(ctx context.Context) HealthStatusOutput {
	return o
}

func (o HealthStatusOutput) ToHealthStatusPtrOutput() HealthStatusPtrOutput {
	return o.ToHealthStatusPtrOutputWithContext(context.Background())
}

func (o HealthStatusOutput) ToHealthStatusPtrOutputWithContext(ctx context.Context) HealthStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HealthStatus) *HealthStatus {
		return &v
	}).(HealthStatusPtrOutput)
}

func (o HealthStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HealthStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HealthStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HealthStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HealthStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HealthStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HealthStatusPtrOutput struct{ *pulumi.OutputState }

func (HealthStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthStatus)(nil)).Elem()
}

func (o HealthStatusPtrOutput) ToHealthStatusPtrOutput() HealthStatusPtrOutput {
	return o
}

func (o HealthStatusPtrOutput) ToHealthStatusPtrOutputWithContext(ctx context.Context) HealthStatusPtrOutput {
	return o
}

func (o HealthStatusPtrOutput) Elem() HealthStatusOutput {
	return o.ApplyT(func(v *HealthStatus) HealthStatus {
		if v != nil {
			return *v
		}
		var ret HealthStatus
		return ret
	}).(HealthStatusOutput)
}

func (o HealthStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HealthStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HealthStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HealthStatusInput interface {
	pulumi.Input

	ToHealthStatusOutput() HealthStatusOutput
	ToHealthStatusOutputWithContext(context.Context) HealthStatusOutput
}

var healthStatusPtrType = reflect.TypeOf((**HealthStatus)(nil)).Elem()

type HealthStatusPtrInput interface {
	pulumi.Input

	ToHealthStatusPtrOutput() HealthStatusPtrOutput
	ToHealthStatusPtrOutputWithContext(context.Context) HealthStatusPtrOutput
}

type healthStatusPtr string

func HealthStatusPtr(v string) HealthStatusPtrInput {
	return (*healthStatusPtr)(&v)
}

func (*healthStatusPtr) ElementType() reflect.Type {
	return healthStatusPtrType
}

func (in *healthStatusPtr) ToHealthStatusPtrOutput() HealthStatusPtrOutput {
	return pulumi.ToOutput(in).(HealthStatusPtrOutput)
}

func (in *healthStatusPtr) ToHealthStatusPtrOutputWithContext(ctx context.Context) HealthStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HealthStatusPtrOutput)
}

type LastBackupStatus string

const (
	LastBackupStatusInvalid   = LastBackupStatus("Invalid")
	LastBackupStatusHealthy   = LastBackupStatus("Healthy")
	LastBackupStatusUnhealthy = LastBackupStatus("Unhealthy")
	LastBackupStatusIRPending = LastBackupStatus("IRPending")
)

func (LastBackupStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*LastBackupStatus)(nil)).Elem()
}

func (e LastBackupStatus) ToLastBackupStatusOutput() LastBackupStatusOutput {
	return pulumi.ToOutput(e).(LastBackupStatusOutput)
}

func (e LastBackupStatus) ToLastBackupStatusOutputWithContext(ctx context.Context) LastBackupStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LastBackupStatusOutput)
}

func (e LastBackupStatus) ToLastBackupStatusPtrOutput() LastBackupStatusPtrOutput {
	return e.ToLastBackupStatusPtrOutputWithContext(context.Background())
}

func (e LastBackupStatus) ToLastBackupStatusPtrOutputWithContext(ctx context.Context) LastBackupStatusPtrOutput {
	return LastBackupStatus(e).ToLastBackupStatusOutputWithContext(ctx).ToLastBackupStatusPtrOutputWithContext(ctx)
}

func (e LastBackupStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LastBackupStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LastBackupStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LastBackupStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LastBackupStatusOutput struct{ *pulumi.OutputState }

func (LastBackupStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LastBackupStatus)(nil)).Elem()
}

func (o LastBackupStatusOutput) ToLastBackupStatusOutput() LastBackupStatusOutput {
	return o
}

func (o LastBackupStatusOutput) ToLastBackupStatusOutputWithContext(ctx context.Context) LastBackupStatusOutput {
	return o
}

func (o LastBackupStatusOutput) ToLastBackupStatusPtrOutput() LastBackupStatusPtrOutput {
	return o.ToLastBackupStatusPtrOutputWithContext(context.Background())
}

func (o LastBackupStatusOutput) ToLastBackupStatusPtrOutputWithContext(ctx context.Context) LastBackupStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LastBackupStatus) *LastBackupStatus {
		return &v
	}).(LastBackupStatusPtrOutput)
}

func (o LastBackupStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LastBackupStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LastBackupStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LastBackupStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LastBackupStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LastBackupStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LastBackupStatusPtrOutput struct{ *pulumi.OutputState }

func (LastBackupStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LastBackupStatus)(nil)).Elem()
}

func (o LastBackupStatusPtrOutput) ToLastBackupStatusPtrOutput() LastBackupStatusPtrOutput {
	return o
}

func (o LastBackupStatusPtrOutput) ToLastBackupStatusPtrOutputWithContext(ctx context.Context) LastBackupStatusPtrOutput {
	return o
}

func (o LastBackupStatusPtrOutput) Elem() LastBackupStatusOutput {
	return o.ApplyT(func(v *LastBackupStatus) LastBackupStatus {
		if v != nil {
			return *v
		}
		var ret LastBackupStatus
		return ret
	}).(LastBackupStatusOutput)
}

func (o LastBackupStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LastBackupStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LastBackupStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LastBackupStatusInput interface {
	pulumi.Input

	ToLastBackupStatusOutput() LastBackupStatusOutput
	ToLastBackupStatusOutputWithContext(context.Context) LastBackupStatusOutput
}

var lastBackupStatusPtrType = reflect.TypeOf((**LastBackupStatus)(nil)).Elem()

type LastBackupStatusPtrInput interface {
	pulumi.Input

	ToLastBackupStatusPtrOutput() LastBackupStatusPtrOutput
	ToLastBackupStatusPtrOutputWithContext(context.Context) LastBackupStatusPtrOutput
}

type lastBackupStatusPtr string

func LastBackupStatusPtr(v string) LastBackupStatusPtrInput {
	return (*lastBackupStatusPtr)(&v)
}

func (*lastBackupStatusPtr) ElementType() reflect.Type {
	return lastBackupStatusPtrType
}

func (in *lastBackupStatusPtr) ToLastBackupStatusPtrOutput() LastBackupStatusPtrOutput {
	return pulumi.ToOutput(in).(LastBackupStatusPtrOutput)
}

func (in *lastBackupStatusPtr) ToLastBackupStatusPtrOutputWithContext(ctx context.Context) LastBackupStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LastBackupStatusPtrOutput)
}

type MonthOfYear string

const (
	MonthOfYearInvalid   = MonthOfYear("Invalid")
	MonthOfYearJanuary   = MonthOfYear("January")
	MonthOfYearFebruary  = MonthOfYear("February")
	MonthOfYearMarch     = MonthOfYear("March")
	MonthOfYearApril     = MonthOfYear("April")
	MonthOfYearMay       = MonthOfYear("May")
	MonthOfYearJune      = MonthOfYear("June")
	MonthOfYearJuly      = MonthOfYear("July")
	MonthOfYearAugust    = MonthOfYear("August")
	MonthOfYearSeptember = MonthOfYear("September")
	MonthOfYearOctober   = MonthOfYear("October")
	MonthOfYearNovember  = MonthOfYear("November")
	MonthOfYearDecember  = MonthOfYear("December")
)

func (MonthOfYear) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthOfYear)(nil)).Elem()
}

func (e MonthOfYear) ToMonthOfYearOutput() MonthOfYearOutput {
	return pulumi.ToOutput(e).(MonthOfYearOutput)
}

func (e MonthOfYear) ToMonthOfYearOutputWithContext(ctx context.Context) MonthOfYearOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonthOfYearOutput)
}

func (e MonthOfYear) ToMonthOfYearPtrOutput() MonthOfYearPtrOutput {
	return e.ToMonthOfYearPtrOutputWithContext(context.Background())
}

func (e MonthOfYear) ToMonthOfYearPtrOutputWithContext(ctx context.Context) MonthOfYearPtrOutput {
	return MonthOfYear(e).ToMonthOfYearOutputWithContext(ctx).ToMonthOfYearPtrOutputWithContext(ctx)
}

func (e MonthOfYear) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonthOfYear) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonthOfYear) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonthOfYear) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonthOfYearOutput struct{ *pulumi.OutputState }

func (MonthOfYearOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthOfYear)(nil)).Elem()
}

func (o MonthOfYearOutput) ToMonthOfYearOutput() MonthOfYearOutput {
	return o
}

func (o MonthOfYearOutput) ToMonthOfYearOutputWithContext(ctx context.Context) MonthOfYearOutput {
	return o
}

func (o MonthOfYearOutput) ToMonthOfYearPtrOutput() MonthOfYearPtrOutput {
	return o.ToMonthOfYearPtrOutputWithContext(context.Background())
}

func (o MonthOfYearOutput) ToMonthOfYearPtrOutputWithContext(ctx context.Context) MonthOfYearPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonthOfYear) *MonthOfYear {
		return &v
	}).(MonthOfYearPtrOutput)
}

func (o MonthOfYearOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonthOfYearOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonthOfYear) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonthOfYearOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonthOfYearOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonthOfYear) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonthOfYearPtrOutput struct{ *pulumi.OutputState }

func (MonthOfYearPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthOfYear)(nil)).Elem()
}

func (o MonthOfYearPtrOutput) ToMonthOfYearPtrOutput() MonthOfYearPtrOutput {
	return o
}

func (o MonthOfYearPtrOutput) ToMonthOfYearPtrOutputWithContext(ctx context.Context) MonthOfYearPtrOutput {
	return o
}

func (o MonthOfYearPtrOutput) Elem() MonthOfYearOutput {
	return o.ApplyT(func(v *MonthOfYear) MonthOfYear {
		if v != nil {
			return *v
		}
		var ret MonthOfYear
		return ret
	}).(MonthOfYearOutput)
}

func (o MonthOfYearPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonthOfYearPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonthOfYear) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MonthOfYearInput interface {
	pulumi.Input

	ToMonthOfYearOutput() MonthOfYearOutput
	ToMonthOfYearOutputWithContext(context.Context) MonthOfYearOutput
}

var monthOfYearPtrType = reflect.TypeOf((**MonthOfYear)(nil)).Elem()

type MonthOfYearPtrInput interface {
	pulumi.Input

	ToMonthOfYearPtrOutput() MonthOfYearPtrOutput
	ToMonthOfYearPtrOutputWithContext(context.Context) MonthOfYearPtrOutput
}

type monthOfYearPtr string

func MonthOfYearPtr(v string) MonthOfYearPtrInput {
	return (*monthOfYearPtr)(&v)
}

func (*monthOfYearPtr) ElementType() reflect.Type {
	return monthOfYearPtrType
}

func (in *monthOfYearPtr) ToMonthOfYearPtrOutput() MonthOfYearPtrOutput {
	return pulumi.ToOutput(in).(MonthOfYearPtrOutput)
}

func (in *monthOfYearPtr) ToMonthOfYearPtrOutputWithContext(ctx context.Context) MonthOfYearPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonthOfYearPtrOutput)
}





type MonthOfYearArrayInput interface {
	pulumi.Input

	ToMonthOfYearArrayOutput() MonthOfYearArrayOutput
	ToMonthOfYearArrayOutputWithContext(context.Context) MonthOfYearArrayOutput
}

type MonthOfYearArray []MonthOfYear

func (MonthOfYearArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonthOfYear)(nil)).Elem()
}

func (i MonthOfYearArray) ToMonthOfYearArrayOutput() MonthOfYearArrayOutput {
	return i.ToMonthOfYearArrayOutputWithContext(context.Background())
}

func (i MonthOfYearArray) ToMonthOfYearArrayOutputWithContext(ctx context.Context) MonthOfYearArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthOfYearArrayOutput)
}

type MonthOfYearArrayOutput struct{ *pulumi.OutputState }

func (MonthOfYearArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonthOfYear)(nil)).Elem()
}

func (o MonthOfYearArrayOutput) ToMonthOfYearArrayOutput() MonthOfYearArrayOutput {
	return o
}

func (o MonthOfYearArrayOutput) ToMonthOfYearArrayOutputWithContext(ctx context.Context) MonthOfYearArrayOutput {
	return o
}

func (o MonthOfYearArrayOutput) Index(i pulumi.IntInput) MonthOfYearOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonthOfYear {
		return vs[0].([]MonthOfYear)[vs[1].(int)]
	}).(MonthOfYearOutput)
}

type OperationType string

const (
	OperationTypeInvalid    = OperationType("Invalid")
	OperationTypeRegister   = OperationType("Register")
	OperationTypeReregister = OperationType("Reregister")
)

func (OperationType) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationType)(nil)).Elem()
}

func (e OperationType) ToOperationTypeOutput() OperationTypeOutput {
	return pulumi.ToOutput(e).(OperationTypeOutput)
}

func (e OperationType) ToOperationTypeOutputWithContext(ctx context.Context) OperationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperationTypeOutput)
}

func (e OperationType) ToOperationTypePtrOutput() OperationTypePtrOutput {
	return e.ToOperationTypePtrOutputWithContext(context.Background())
}

func (e OperationType) ToOperationTypePtrOutputWithContext(ctx context.Context) OperationTypePtrOutput {
	return OperationType(e).ToOperationTypeOutputWithContext(ctx).ToOperationTypePtrOutputWithContext(ctx)
}

func (e OperationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperationTypeOutput struct{ *pulumi.OutputState }

func (OperationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationType)(nil)).Elem()
}

func (o OperationTypeOutput) ToOperationTypeOutput() OperationTypeOutput {
	return o
}

func (o OperationTypeOutput) ToOperationTypeOutputWithContext(ctx context.Context) OperationTypeOutput {
	return o
}

func (o OperationTypeOutput) ToOperationTypePtrOutput() OperationTypePtrOutput {
	return o.ToOperationTypePtrOutputWithContext(context.Background())
}

func (o OperationTypeOutput) ToOperationTypePtrOutputWithContext(ctx context.Context) OperationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperationType) *OperationType {
		return &v
	}).(OperationTypePtrOutput)
}

func (o OperationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperationTypePtrOutput struct{ *pulumi.OutputState }

func (OperationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationType)(nil)).Elem()
}

func (o OperationTypePtrOutput) ToOperationTypePtrOutput() OperationTypePtrOutput {
	return o
}

func (o OperationTypePtrOutput) ToOperationTypePtrOutputWithContext(ctx context.Context) OperationTypePtrOutput {
	return o
}

func (o OperationTypePtrOutput) Elem() OperationTypeOutput {
	return o.ApplyT(func(v *OperationType) OperationType {
		if v != nil {
			return *v
		}
		var ret OperationType
		return ret
	}).(OperationTypeOutput)
}

func (o OperationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperationTypeInput interface {
	pulumi.Input

	ToOperationTypeOutput() OperationTypeOutput
	ToOperationTypeOutputWithContext(context.Context) OperationTypeOutput
}

var operationTypePtrType = reflect.TypeOf((**OperationType)(nil)).Elem()

type OperationTypePtrInput interface {
	pulumi.Input

	ToOperationTypePtrOutput() OperationTypePtrOutput
	ToOperationTypePtrOutputWithContext(context.Context) OperationTypePtrOutput
}

type operationTypePtr string

func OperationTypePtr(v string) OperationTypePtrInput {
	return (*operationTypePtr)(&v)
}

func (*operationTypePtr) ElementType() reflect.Type {
	return operationTypePtrType
}

func (in *operationTypePtr) ToOperationTypePtrOutput() OperationTypePtrOutput {
	return pulumi.ToOutput(in).(OperationTypePtrOutput)
}

func (in *operationTypePtr) ToOperationTypePtrOutputWithContext(ctx context.Context) OperationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperationTypePtrOutput)
}

type PolicyType string

const (
	PolicyTypeInvalid      = PolicyType("Invalid")
	PolicyTypeFull         = PolicyType("Full")
	PolicyTypeDifferential = PolicyType("Differential")
	PolicyTypeLog          = PolicyType("Log")
	PolicyTypeCopyOnlyFull = PolicyType("CopyOnlyFull")
	PolicyTypeIncremental  = PolicyType("Incremental")
)

func (PolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (e PolicyType) ToPolicyTypeOutput() PolicyTypeOutput {
	return pulumi.ToOutput(e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return e.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (e PolicyType) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return PolicyType(e).ToPolicyTypeOutputWithContext(ctx).ToPolicyTypePtrOutputWithContext(ctx)
}

func (e PolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyTypeOutput struct{ *pulumi.OutputState }

func (PolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (o PolicyTypeOutput) ToPolicyTypeOutput() PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyType) *PolicyType {
		return &v
	}).(PolicyTypePtrOutput)
}

func (o PolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyTypePtrOutput struct{ *pulumi.OutputState }

func (PolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyType)(nil)).Elem()
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) Elem() PolicyTypeOutput {
	return o.ApplyT(func(v *PolicyType) PolicyType {
		if v != nil {
			return *v
		}
		var ret PolicyType
		return ret
	}).(PolicyTypeOutput)
}

func (o PolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PolicyTypeInput interface {
	pulumi.Input

	ToPolicyTypeOutput() PolicyTypeOutput
	ToPolicyTypeOutputWithContext(context.Context) PolicyTypeOutput
}

var policyTypePtrType = reflect.TypeOf((**PolicyType)(nil)).Elem()

type PolicyTypePtrInput interface {
	pulumi.Input

	ToPolicyTypePtrOutput() PolicyTypePtrOutput
	ToPolicyTypePtrOutputWithContext(context.Context) PolicyTypePtrOutput
}

type policyTypePtr string

func PolicyTypePtr(v string) PolicyTypePtrInput {
	return (*policyTypePtr)(&v)
}

func (*policyTypePtr) ElementType() reflect.Type {
	return policyTypePtrType
}

func (in *policyTypePtr) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return pulumi.ToOutput(in).(PolicyTypePtrOutput)
}

func (in *policyTypePtr) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyTypePtrOutput)
}

type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusPending      = PrivateEndpointConnectionStatus("Pending")
	PrivateEndpointConnectionStatusApproved     = PrivateEndpointConnectionStatus("Approved")
	PrivateEndpointConnectionStatusRejected     = PrivateEndpointConnectionStatus("Rejected")
	PrivateEndpointConnectionStatusDisconnected = PrivateEndpointConnectionStatus("Disconnected")
)

func (PrivateEndpointConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointConnectionStatus) ToPrivateEndpointConnectionStatusOutput() PrivateEndpointConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointConnectionStatusOutput)
}

func (e PrivateEndpointConnectionStatus) ToPrivateEndpointConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointConnectionStatusOutput)
}

func (e PrivateEndpointConnectionStatus) ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput {
	return e.ToPrivateEndpointConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointConnectionStatus) ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusPtrOutput {
	return PrivateEndpointConnectionStatus(e).ToPrivateEndpointConnectionStatusOutputWithContext(ctx).ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointConnectionStatusOutput) ToPrivateEndpointConnectionStatusOutput() PrivateEndpointConnectionStatusOutput {
	return o
}

func (o PrivateEndpointConnectionStatusOutput) ToPrivateEndpointConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusOutput {
	return o
}

func (o PrivateEndpointConnectionStatusOutput) ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput {
	return o.ToPrivateEndpointConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionStatusOutput) ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionStatus) *PrivateEndpointConnectionStatus {
		return &v
	}).(PrivateEndpointConnectionStatusPtrOutput)
}

func (o PrivateEndpointConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointConnectionStatusPtrOutput) ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointConnectionStatusPtrOutput) ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointConnectionStatusPtrOutput) Elem() PrivateEndpointConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionStatus) PrivateEndpointConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionStatus
		return ret
	}).(PrivateEndpointConnectionStatusOutput)
}

func (o PrivateEndpointConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionStatusOutput() PrivateEndpointConnectionStatusOutput
	ToPrivateEndpointConnectionStatusOutputWithContext(context.Context) PrivateEndpointConnectionStatusOutput
}

var privateEndpointConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointConnectionStatus)(nil)).Elem()

type PrivateEndpointConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput
	ToPrivateEndpointConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointConnectionStatusPtrOutput
}

type privateEndpointConnectionStatusPtr string

func PrivateEndpointConnectionStatusPtr(v string) PrivateEndpointConnectionStatusPtrInput {
	return (*privateEndpointConnectionStatusPtr)(&v)
}

func (*privateEndpointConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointConnectionStatusPtrType
}

func (in *privateEndpointConnectionStatusPtr) ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointConnectionStatusPtrOutput)
}

func (in *privateEndpointConnectionStatusPtr) ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointConnectionStatusPtrOutput)
}

type ProtectedItemHealthStatus string

const (
	ProtectedItemHealthStatusInvalid      = ProtectedItemHealthStatus("Invalid")
	ProtectedItemHealthStatusHealthy      = ProtectedItemHealthStatus("Healthy")
	ProtectedItemHealthStatusUnhealthy    = ProtectedItemHealthStatus("Unhealthy")
	ProtectedItemHealthStatusNotReachable = ProtectedItemHealthStatus("NotReachable")
	ProtectedItemHealthStatusIRPending    = ProtectedItemHealthStatus("IRPending")
)

func (ProtectedItemHealthStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedItemHealthStatus)(nil)).Elem()
}

func (e ProtectedItemHealthStatus) ToProtectedItemHealthStatusOutput() ProtectedItemHealthStatusOutput {
	return pulumi.ToOutput(e).(ProtectedItemHealthStatusOutput)
}

func (e ProtectedItemHealthStatus) ToProtectedItemHealthStatusOutputWithContext(ctx context.Context) ProtectedItemHealthStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtectedItemHealthStatusOutput)
}

func (e ProtectedItemHealthStatus) ToProtectedItemHealthStatusPtrOutput() ProtectedItemHealthStatusPtrOutput {
	return e.ToProtectedItemHealthStatusPtrOutputWithContext(context.Background())
}

func (e ProtectedItemHealthStatus) ToProtectedItemHealthStatusPtrOutputWithContext(ctx context.Context) ProtectedItemHealthStatusPtrOutput {
	return ProtectedItemHealthStatus(e).ToProtectedItemHealthStatusOutputWithContext(ctx).ToProtectedItemHealthStatusPtrOutputWithContext(ctx)
}

func (e ProtectedItemHealthStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectedItemHealthStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectedItemHealthStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtectedItemHealthStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtectedItemHealthStatusOutput struct{ *pulumi.OutputState }

func (ProtectedItemHealthStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedItemHealthStatus)(nil)).Elem()
}

func (o ProtectedItemHealthStatusOutput) ToProtectedItemHealthStatusOutput() ProtectedItemHealthStatusOutput {
	return o
}

func (o ProtectedItemHealthStatusOutput) ToProtectedItemHealthStatusOutputWithContext(ctx context.Context) ProtectedItemHealthStatusOutput {
	return o
}

func (o ProtectedItemHealthStatusOutput) ToProtectedItemHealthStatusPtrOutput() ProtectedItemHealthStatusPtrOutput {
	return o.ToProtectedItemHealthStatusPtrOutputWithContext(context.Background())
}

func (o ProtectedItemHealthStatusOutput) ToProtectedItemHealthStatusPtrOutputWithContext(ctx context.Context) ProtectedItemHealthStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectedItemHealthStatus) *ProtectedItemHealthStatus {
		return &v
	}).(ProtectedItemHealthStatusPtrOutput)
}

func (o ProtectedItemHealthStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtectedItemHealthStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectedItemHealthStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtectedItemHealthStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectedItemHealthStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectedItemHealthStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtectedItemHealthStatusPtrOutput struct{ *pulumi.OutputState }

func (ProtectedItemHealthStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedItemHealthStatus)(nil)).Elem()
}

func (o ProtectedItemHealthStatusPtrOutput) ToProtectedItemHealthStatusPtrOutput() ProtectedItemHealthStatusPtrOutput {
	return o
}

func (o ProtectedItemHealthStatusPtrOutput) ToProtectedItemHealthStatusPtrOutputWithContext(ctx context.Context) ProtectedItemHealthStatusPtrOutput {
	return o
}

func (o ProtectedItemHealthStatusPtrOutput) Elem() ProtectedItemHealthStatusOutput {
	return o.ApplyT(func(v *ProtectedItemHealthStatus) ProtectedItemHealthStatus {
		if v != nil {
			return *v
		}
		var ret ProtectedItemHealthStatus
		return ret
	}).(ProtectedItemHealthStatusOutput)
}

func (o ProtectedItemHealthStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectedItemHealthStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProtectedItemHealthStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtectedItemHealthStatusInput interface {
	pulumi.Input

	ToProtectedItemHealthStatusOutput() ProtectedItemHealthStatusOutput
	ToProtectedItemHealthStatusOutputWithContext(context.Context) ProtectedItemHealthStatusOutput
}

var protectedItemHealthStatusPtrType = reflect.TypeOf((**ProtectedItemHealthStatus)(nil)).Elem()

type ProtectedItemHealthStatusPtrInput interface {
	pulumi.Input

	ToProtectedItemHealthStatusPtrOutput() ProtectedItemHealthStatusPtrOutput
	ToProtectedItemHealthStatusPtrOutputWithContext(context.Context) ProtectedItemHealthStatusPtrOutput
}

type protectedItemHealthStatusPtr string

func ProtectedItemHealthStatusPtr(v string) ProtectedItemHealthStatusPtrInput {
	return (*protectedItemHealthStatusPtr)(&v)
}

func (*protectedItemHealthStatusPtr) ElementType() reflect.Type {
	return protectedItemHealthStatusPtrType
}

func (in *protectedItemHealthStatusPtr) ToProtectedItemHealthStatusPtrOutput() ProtectedItemHealthStatusPtrOutput {
	return pulumi.ToOutput(in).(ProtectedItemHealthStatusPtrOutput)
}

func (in *protectedItemHealthStatusPtr) ToProtectedItemHealthStatusPtrOutputWithContext(ctx context.Context) ProtectedItemHealthStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtectedItemHealthStatusPtrOutput)
}

type ProtectedItemStateEnum string

const (
	ProtectedItemStateEnumInvalid           = ProtectedItemStateEnum("Invalid")
	ProtectedItemStateEnumIRPending         = ProtectedItemStateEnum("IRPending")
	ProtectedItemStateEnumProtected         = ProtectedItemStateEnum("Protected")
	ProtectedItemStateEnumProtectionError   = ProtectedItemStateEnum("ProtectionError")
	ProtectedItemStateEnumProtectionStopped = ProtectedItemStateEnum("ProtectionStopped")
	ProtectedItemStateEnumProtectionPaused  = ProtectedItemStateEnum("ProtectionPaused")
)

func (ProtectedItemStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedItemStateEnum)(nil)).Elem()
}

func (e ProtectedItemStateEnum) ToProtectedItemStateEnumOutput() ProtectedItemStateEnumOutput {
	return pulumi.ToOutput(e).(ProtectedItemStateEnumOutput)
}

func (e ProtectedItemStateEnum) ToProtectedItemStateEnumOutputWithContext(ctx context.Context) ProtectedItemStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtectedItemStateEnumOutput)
}

func (e ProtectedItemStateEnum) ToProtectedItemStateEnumPtrOutput() ProtectedItemStateEnumPtrOutput {
	return e.ToProtectedItemStateEnumPtrOutputWithContext(context.Background())
}

func (e ProtectedItemStateEnum) ToProtectedItemStateEnumPtrOutputWithContext(ctx context.Context) ProtectedItemStateEnumPtrOutput {
	return ProtectedItemStateEnum(e).ToProtectedItemStateEnumOutputWithContext(ctx).ToProtectedItemStateEnumPtrOutputWithContext(ctx)
}

func (e ProtectedItemStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectedItemStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectedItemStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtectedItemStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtectedItemStateEnumOutput struct{ *pulumi.OutputState }

func (ProtectedItemStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectedItemStateEnum)(nil)).Elem()
}

func (o ProtectedItemStateEnumOutput) ToProtectedItemStateEnumOutput() ProtectedItemStateEnumOutput {
	return o
}

func (o ProtectedItemStateEnumOutput) ToProtectedItemStateEnumOutputWithContext(ctx context.Context) ProtectedItemStateEnumOutput {
	return o
}

func (o ProtectedItemStateEnumOutput) ToProtectedItemStateEnumPtrOutput() ProtectedItemStateEnumPtrOutput {
	return o.ToProtectedItemStateEnumPtrOutputWithContext(context.Background())
}

func (o ProtectedItemStateEnumOutput) ToProtectedItemStateEnumPtrOutputWithContext(ctx context.Context) ProtectedItemStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectedItemStateEnum) *ProtectedItemStateEnum {
		return &v
	}).(ProtectedItemStateEnumPtrOutput)
}

func (o ProtectedItemStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtectedItemStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectedItemStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtectedItemStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectedItemStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectedItemStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtectedItemStateEnumPtrOutput struct{ *pulumi.OutputState }

func (ProtectedItemStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectedItemStateEnum)(nil)).Elem()
}

func (o ProtectedItemStateEnumPtrOutput) ToProtectedItemStateEnumPtrOutput() ProtectedItemStateEnumPtrOutput {
	return o
}

func (o ProtectedItemStateEnumPtrOutput) ToProtectedItemStateEnumPtrOutputWithContext(ctx context.Context) ProtectedItemStateEnumPtrOutput {
	return o
}

func (o ProtectedItemStateEnumPtrOutput) Elem() ProtectedItemStateEnumOutput {
	return o.ApplyT(func(v *ProtectedItemStateEnum) ProtectedItemStateEnum {
		if v != nil {
			return *v
		}
		var ret ProtectedItemStateEnum
		return ret
	}).(ProtectedItemStateEnumOutput)
}

func (o ProtectedItemStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectedItemStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProtectedItemStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtectedItemStateEnumInput interface {
	pulumi.Input

	ToProtectedItemStateEnumOutput() ProtectedItemStateEnumOutput
	ToProtectedItemStateEnumOutputWithContext(context.Context) ProtectedItemStateEnumOutput
}

var protectedItemStateEnumPtrType = reflect.TypeOf((**ProtectedItemStateEnum)(nil)).Elem()

type ProtectedItemStateEnumPtrInput interface {
	pulumi.Input

	ToProtectedItemStateEnumPtrOutput() ProtectedItemStateEnumPtrOutput
	ToProtectedItemStateEnumPtrOutputWithContext(context.Context) ProtectedItemStateEnumPtrOutput
}

type protectedItemStateEnumPtr string

func ProtectedItemStateEnumPtr(v string) ProtectedItemStateEnumPtrInput {
	return (*protectedItemStateEnumPtr)(&v)
}

func (*protectedItemStateEnumPtr) ElementType() reflect.Type {
	return protectedItemStateEnumPtrType
}

func (in *protectedItemStateEnumPtr) ToProtectedItemStateEnumPtrOutput() ProtectedItemStateEnumPtrOutput {
	return pulumi.ToOutput(in).(ProtectedItemStateEnumPtrOutput)
}

func (in *protectedItemStateEnumPtr) ToProtectedItemStateEnumPtrOutputWithContext(ctx context.Context) ProtectedItemStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtectedItemStateEnumPtrOutput)
}

type ProtectionState string

const (
	ProtectionStateInvalid           = ProtectionState("Invalid")
	ProtectionStateIRPending         = ProtectionState("IRPending")
	ProtectionStateProtected         = ProtectionState("Protected")
	ProtectionStateProtectionError   = ProtectionState("ProtectionError")
	ProtectionStateProtectionStopped = ProtectionState("ProtectionStopped")
	ProtectionStateProtectionPaused  = ProtectionState("ProtectionPaused")
)

func (ProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionState)(nil)).Elem()
}

func (e ProtectionState) ToProtectionStateOutput() ProtectionStateOutput {
	return pulumi.ToOutput(e).(ProtectionStateOutput)
}

func (e ProtectionState) ToProtectionStateOutputWithContext(ctx context.Context) ProtectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtectionStateOutput)
}

func (e ProtectionState) ToProtectionStatePtrOutput() ProtectionStatePtrOutput {
	return e.ToProtectionStatePtrOutputWithContext(context.Background())
}

func (e ProtectionState) ToProtectionStatePtrOutputWithContext(ctx context.Context) ProtectionStatePtrOutput {
	return ProtectionState(e).ToProtectionStateOutputWithContext(ctx).ToProtectionStatePtrOutputWithContext(ctx)
}

func (e ProtectionState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectionState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectionState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtectionState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtectionStateOutput struct{ *pulumi.OutputState }

func (ProtectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionState)(nil)).Elem()
}

func (o ProtectionStateOutput) ToProtectionStateOutput() ProtectionStateOutput {
	return o
}

func (o ProtectionStateOutput) ToProtectionStateOutputWithContext(ctx context.Context) ProtectionStateOutput {
	return o
}

func (o ProtectionStateOutput) ToProtectionStatePtrOutput() ProtectionStatePtrOutput {
	return o.ToProtectionStatePtrOutputWithContext(context.Background())
}

func (o ProtectionStateOutput) ToProtectionStatePtrOutputWithContext(ctx context.Context) ProtectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectionState) *ProtectionState {
		return &v
	}).(ProtectionStatePtrOutput)
}

func (o ProtectionStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtectionStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectionState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtectionStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectionStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectionState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtectionStatePtrOutput struct{ *pulumi.OutputState }

func (ProtectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionState)(nil)).Elem()
}

func (o ProtectionStatePtrOutput) ToProtectionStatePtrOutput() ProtectionStatePtrOutput {
	return o
}

func (o ProtectionStatePtrOutput) ToProtectionStatePtrOutputWithContext(ctx context.Context) ProtectionStatePtrOutput {
	return o
}

func (o ProtectionStatePtrOutput) Elem() ProtectionStateOutput {
	return o.ApplyT(func(v *ProtectionState) ProtectionState {
		if v != nil {
			return *v
		}
		var ret ProtectionState
		return ret
	}).(ProtectionStateOutput)
}

func (o ProtectionStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectionStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProtectionState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtectionStateInput interface {
	pulumi.Input

	ToProtectionStateOutput() ProtectionStateOutput
	ToProtectionStateOutputWithContext(context.Context) ProtectionStateOutput
}

var protectionStatePtrType = reflect.TypeOf((**ProtectionState)(nil)).Elem()

type ProtectionStatePtrInput interface {
	pulumi.Input

	ToProtectionStatePtrOutput() ProtectionStatePtrOutput
	ToProtectionStatePtrOutputWithContext(context.Context) ProtectionStatePtrOutput
}

type protectionStatePtr string

func ProtectionStatePtr(v string) ProtectionStatePtrInput {
	return (*protectionStatePtr)(&v)
}

func (*protectionStatePtr) ElementType() reflect.Type {
	return protectionStatePtrType
}

func (in *protectionStatePtr) ToProtectionStatePtrOutput() ProtectionStatePtrOutput {
	return pulumi.ToOutput(in).(ProtectionStatePtrOutput)
}

func (in *protectionStatePtr) ToProtectionStatePtrOutputWithContext(ctx context.Context) ProtectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtectionStatePtrOutput)
}

type ProtectionStatus string

const (
	ProtectionStatusInvalid          = ProtectionStatus("Invalid")
	ProtectionStatusNotProtected     = ProtectionStatus("NotProtected")
	ProtectionStatusProtecting       = ProtectionStatus("Protecting")
	ProtectionStatusProtected        = ProtectionStatus("Protected")
	ProtectionStatusProtectionFailed = ProtectionStatus("ProtectionFailed")
)

func (ProtectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionStatus)(nil)).Elem()
}

func (e ProtectionStatus) ToProtectionStatusOutput() ProtectionStatusOutput {
	return pulumi.ToOutput(e).(ProtectionStatusOutput)
}

func (e ProtectionStatus) ToProtectionStatusOutputWithContext(ctx context.Context) ProtectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtectionStatusOutput)
}

func (e ProtectionStatus) ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput {
	return e.ToProtectionStatusPtrOutputWithContext(context.Background())
}

func (e ProtectionStatus) ToProtectionStatusPtrOutputWithContext(ctx context.Context) ProtectionStatusPtrOutput {
	return ProtectionStatus(e).ToProtectionStatusOutputWithContext(ctx).ToProtectionStatusPtrOutputWithContext(ctx)
}

func (e ProtectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtectionStatusOutput struct{ *pulumi.OutputState }

func (ProtectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionStatus)(nil)).Elem()
}

func (o ProtectionStatusOutput) ToProtectionStatusOutput() ProtectionStatusOutput {
	return o
}

func (o ProtectionStatusOutput) ToProtectionStatusOutputWithContext(ctx context.Context) ProtectionStatusOutput {
	return o
}

func (o ProtectionStatusOutput) ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput {
	return o.ToProtectionStatusPtrOutputWithContext(context.Background())
}

func (o ProtectionStatusOutput) ToProtectionStatusPtrOutputWithContext(ctx context.Context) ProtectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectionStatus) *ProtectionStatus {
		return &v
	}).(ProtectionStatusPtrOutput)
}

func (o ProtectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtectionStatusPtrOutput struct{ *pulumi.OutputState }

func (ProtectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionStatus)(nil)).Elem()
}

func (o ProtectionStatusPtrOutput) ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput {
	return o
}

func (o ProtectionStatusPtrOutput) ToProtectionStatusPtrOutputWithContext(ctx context.Context) ProtectionStatusPtrOutput {
	return o
}

func (o ProtectionStatusPtrOutput) Elem() ProtectionStatusOutput {
	return o.ApplyT(func(v *ProtectionStatus) ProtectionStatus {
		if v != nil {
			return *v
		}
		var ret ProtectionStatus
		return ret
	}).(ProtectionStatusOutput)
}

func (o ProtectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProtectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtectionStatusInput interface {
	pulumi.Input

	ToProtectionStatusOutput() ProtectionStatusOutput
	ToProtectionStatusOutputWithContext(context.Context) ProtectionStatusOutput
}

var protectionStatusPtrType = reflect.TypeOf((**ProtectionStatus)(nil)).Elem()

type ProtectionStatusPtrInput interface {
	pulumi.Input

	ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput
	ToProtectionStatusPtrOutputWithContext(context.Context) ProtectionStatusPtrOutput
}

type protectionStatusPtr string

func ProtectionStatusPtr(v string) ProtectionStatusPtrInput {
	return (*protectionStatusPtr)(&v)
}

func (*protectionStatusPtr) ElementType() reflect.Type {
	return protectionStatusPtrType
}

func (in *protectionStatusPtr) ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput {
	return pulumi.ToOutput(in).(ProtectionStatusPtrOutput)
}

func (in *protectionStatusPtr) ToProtectionStatusPtrOutputWithContext(ctx context.Context) ProtectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtectionStatusPtrOutput)
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	ProvisioningStateDeleting  = ProvisioningState("Deleting")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStatePending   = ProvisioningState("Pending")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		if v != nil {
			return *v
		}
		var ret ProvisioningState
		return ret
	}).(ProvisioningStateOutput)
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

type ResourceHealthStatus string

const (
	ResourceHealthStatusHealthy             = ResourceHealthStatus("Healthy")
	ResourceHealthStatusTransientDegraded   = ResourceHealthStatus("TransientDegraded")
	ResourceHealthStatusPersistentDegraded  = ResourceHealthStatus("PersistentDegraded")
	ResourceHealthStatusTransientUnhealthy  = ResourceHealthStatus("TransientUnhealthy")
	ResourceHealthStatusPersistentUnhealthy = ResourceHealthStatus("PersistentUnhealthy")
	ResourceHealthStatusInvalid             = ResourceHealthStatus("Invalid")
)

func (ResourceHealthStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceHealthStatus)(nil)).Elem()
}

func (e ResourceHealthStatus) ToResourceHealthStatusOutput() ResourceHealthStatusOutput {
	return pulumi.ToOutput(e).(ResourceHealthStatusOutput)
}

func (e ResourceHealthStatus) ToResourceHealthStatusOutputWithContext(ctx context.Context) ResourceHealthStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceHealthStatusOutput)
}

func (e ResourceHealthStatus) ToResourceHealthStatusPtrOutput() ResourceHealthStatusPtrOutput {
	return e.ToResourceHealthStatusPtrOutputWithContext(context.Background())
}

func (e ResourceHealthStatus) ToResourceHealthStatusPtrOutputWithContext(ctx context.Context) ResourceHealthStatusPtrOutput {
	return ResourceHealthStatus(e).ToResourceHealthStatusOutputWithContext(ctx).ToResourceHealthStatusPtrOutputWithContext(ctx)
}

func (e ResourceHealthStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceHealthStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceHealthStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceHealthStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceHealthStatusOutput struct{ *pulumi.OutputState }

func (ResourceHealthStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceHealthStatus)(nil)).Elem()
}

func (o ResourceHealthStatusOutput) ToResourceHealthStatusOutput() ResourceHealthStatusOutput {
	return o
}

func (o ResourceHealthStatusOutput) ToResourceHealthStatusOutputWithContext(ctx context.Context) ResourceHealthStatusOutput {
	return o
}

func (o ResourceHealthStatusOutput) ToResourceHealthStatusPtrOutput() ResourceHealthStatusPtrOutput {
	return o.ToResourceHealthStatusPtrOutputWithContext(context.Background())
}

func (o ResourceHealthStatusOutput) ToResourceHealthStatusPtrOutputWithContext(ctx context.Context) ResourceHealthStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceHealthStatus) *ResourceHealthStatus {
		return &v
	}).(ResourceHealthStatusPtrOutput)
}

func (o ResourceHealthStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceHealthStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceHealthStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceHealthStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceHealthStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceHealthStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceHealthStatusPtrOutput struct{ *pulumi.OutputState }

func (ResourceHealthStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceHealthStatus)(nil)).Elem()
}

func (o ResourceHealthStatusPtrOutput) ToResourceHealthStatusPtrOutput() ResourceHealthStatusPtrOutput {
	return o
}

func (o ResourceHealthStatusPtrOutput) ToResourceHealthStatusPtrOutputWithContext(ctx context.Context) ResourceHealthStatusPtrOutput {
	return o
}

func (o ResourceHealthStatusPtrOutput) Elem() ResourceHealthStatusOutput {
	return o.ApplyT(func(v *ResourceHealthStatus) ResourceHealthStatus {
		if v != nil {
			return *v
		}
		var ret ResourceHealthStatus
		return ret
	}).(ResourceHealthStatusOutput)
}

func (o ResourceHealthStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceHealthStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceHealthStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceHealthStatusInput interface {
	pulumi.Input

	ToResourceHealthStatusOutput() ResourceHealthStatusOutput
	ToResourceHealthStatusOutputWithContext(context.Context) ResourceHealthStatusOutput
}

var resourceHealthStatusPtrType = reflect.TypeOf((**ResourceHealthStatus)(nil)).Elem()

type ResourceHealthStatusPtrInput interface {
	pulumi.Input

	ToResourceHealthStatusPtrOutput() ResourceHealthStatusPtrOutput
	ToResourceHealthStatusPtrOutputWithContext(context.Context) ResourceHealthStatusPtrOutput
}

type resourceHealthStatusPtr string

func ResourceHealthStatusPtr(v string) ResourceHealthStatusPtrInput {
	return (*resourceHealthStatusPtr)(&v)
}

func (*resourceHealthStatusPtr) ElementType() reflect.Type {
	return resourceHealthStatusPtrType
}

func (in *resourceHealthStatusPtr) ToResourceHealthStatusPtrOutput() ResourceHealthStatusPtrOutput {
	return pulumi.ToOutput(in).(ResourceHealthStatusPtrOutput)
}

func (in *resourceHealthStatusPtr) ToResourceHealthStatusPtrOutputWithContext(ctx context.Context) ResourceHealthStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceHealthStatusPtrOutput)
}

type RetentionDurationType string

const (
	RetentionDurationTypeInvalid = RetentionDurationType("Invalid")
	RetentionDurationTypeDays    = RetentionDurationType("Days")
	RetentionDurationTypeWeeks   = RetentionDurationType("Weeks")
	RetentionDurationTypeMonths  = RetentionDurationType("Months")
	RetentionDurationTypeYears   = RetentionDurationType("Years")
)

func (RetentionDurationType) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDurationType)(nil)).Elem()
}

func (e RetentionDurationType) ToRetentionDurationTypeOutput() RetentionDurationTypeOutput {
	return pulumi.ToOutput(e).(RetentionDurationTypeOutput)
}

func (e RetentionDurationType) ToRetentionDurationTypeOutputWithContext(ctx context.Context) RetentionDurationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RetentionDurationTypeOutput)
}

func (e RetentionDurationType) ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput {
	return e.ToRetentionDurationTypePtrOutputWithContext(context.Background())
}

func (e RetentionDurationType) ToRetentionDurationTypePtrOutputWithContext(ctx context.Context) RetentionDurationTypePtrOutput {
	return RetentionDurationType(e).ToRetentionDurationTypeOutputWithContext(ctx).ToRetentionDurationTypePtrOutputWithContext(ctx)
}

func (e RetentionDurationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionDurationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionDurationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RetentionDurationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RetentionDurationTypeOutput struct{ *pulumi.OutputState }

func (RetentionDurationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDurationType)(nil)).Elem()
}

func (o RetentionDurationTypeOutput) ToRetentionDurationTypeOutput() RetentionDurationTypeOutput {
	return o
}

func (o RetentionDurationTypeOutput) ToRetentionDurationTypeOutputWithContext(ctx context.Context) RetentionDurationTypeOutput {
	return o
}

func (o RetentionDurationTypeOutput) ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput {
	return o.ToRetentionDurationTypePtrOutputWithContext(context.Background())
}

func (o RetentionDurationTypeOutput) ToRetentionDurationTypePtrOutputWithContext(ctx context.Context) RetentionDurationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionDurationType) *RetentionDurationType {
		return &v
	}).(RetentionDurationTypePtrOutput)
}

func (o RetentionDurationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RetentionDurationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RetentionDurationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RetentionDurationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RetentionDurationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RetentionDurationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RetentionDurationTypePtrOutput struct{ *pulumi.OutputState }

func (RetentionDurationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDurationType)(nil)).Elem()
}

func (o RetentionDurationTypePtrOutput) ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput {
	return o
}

func (o RetentionDurationTypePtrOutput) ToRetentionDurationTypePtrOutputWithContext(ctx context.Context) RetentionDurationTypePtrOutput {
	return o
}

func (o RetentionDurationTypePtrOutput) Elem() RetentionDurationTypeOutput {
	return o.ApplyT(func(v *RetentionDurationType) RetentionDurationType {
		if v != nil {
			return *v
		}
		var ret RetentionDurationType
		return ret
	}).(RetentionDurationTypeOutput)
}

func (o RetentionDurationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RetentionDurationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RetentionDurationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RetentionDurationTypeInput interface {
	pulumi.Input

	ToRetentionDurationTypeOutput() RetentionDurationTypeOutput
	ToRetentionDurationTypeOutputWithContext(context.Context) RetentionDurationTypeOutput
}

var retentionDurationTypePtrType = reflect.TypeOf((**RetentionDurationType)(nil)).Elem()

type RetentionDurationTypePtrInput interface {
	pulumi.Input

	ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput
	ToRetentionDurationTypePtrOutputWithContext(context.Context) RetentionDurationTypePtrOutput
}

type retentionDurationTypePtr string

func RetentionDurationTypePtr(v string) RetentionDurationTypePtrInput {
	return (*retentionDurationTypePtr)(&v)
}

func (*retentionDurationTypePtr) ElementType() reflect.Type {
	return retentionDurationTypePtrType
}

func (in *retentionDurationTypePtr) ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput {
	return pulumi.ToOutput(in).(RetentionDurationTypePtrOutput)
}

func (in *retentionDurationTypePtr) ToRetentionDurationTypePtrOutputWithContext(ctx context.Context) RetentionDurationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RetentionDurationTypePtrOutput)
}

type RetentionScheduleFormat string

const (
	RetentionScheduleFormatInvalid = RetentionScheduleFormat("Invalid")
	RetentionScheduleFormatDaily   = RetentionScheduleFormat("Daily")
	RetentionScheduleFormatWeekly  = RetentionScheduleFormat("Weekly")
)

func (RetentionScheduleFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionScheduleFormat)(nil)).Elem()
}

func (e RetentionScheduleFormat) ToRetentionScheduleFormatOutput() RetentionScheduleFormatOutput {
	return pulumi.ToOutput(e).(RetentionScheduleFormatOutput)
}

func (e RetentionScheduleFormat) ToRetentionScheduleFormatOutputWithContext(ctx context.Context) RetentionScheduleFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RetentionScheduleFormatOutput)
}

func (e RetentionScheduleFormat) ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput {
	return e.ToRetentionScheduleFormatPtrOutputWithContext(context.Background())
}

func (e RetentionScheduleFormat) ToRetentionScheduleFormatPtrOutputWithContext(ctx context.Context) RetentionScheduleFormatPtrOutput {
	return RetentionScheduleFormat(e).ToRetentionScheduleFormatOutputWithContext(ctx).ToRetentionScheduleFormatPtrOutputWithContext(ctx)
}

func (e RetentionScheduleFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionScheduleFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionScheduleFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RetentionScheduleFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RetentionScheduleFormatOutput struct{ *pulumi.OutputState }

func (RetentionScheduleFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionScheduleFormat)(nil)).Elem()
}

func (o RetentionScheduleFormatOutput) ToRetentionScheduleFormatOutput() RetentionScheduleFormatOutput {
	return o
}

func (o RetentionScheduleFormatOutput) ToRetentionScheduleFormatOutputWithContext(ctx context.Context) RetentionScheduleFormatOutput {
	return o
}

func (o RetentionScheduleFormatOutput) ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput {
	return o.ToRetentionScheduleFormatPtrOutputWithContext(context.Background())
}

func (o RetentionScheduleFormatOutput) ToRetentionScheduleFormatPtrOutputWithContext(ctx context.Context) RetentionScheduleFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionScheduleFormat) *RetentionScheduleFormat {
		return &v
	}).(RetentionScheduleFormatPtrOutput)
}

func (o RetentionScheduleFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RetentionScheduleFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RetentionScheduleFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RetentionScheduleFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RetentionScheduleFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RetentionScheduleFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RetentionScheduleFormatPtrOutput struct{ *pulumi.OutputState }

func (RetentionScheduleFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionScheduleFormat)(nil)).Elem()
}

func (o RetentionScheduleFormatPtrOutput) ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput {
	return o
}

func (o RetentionScheduleFormatPtrOutput) ToRetentionScheduleFormatPtrOutputWithContext(ctx context.Context) RetentionScheduleFormatPtrOutput {
	return o
}

func (o RetentionScheduleFormatPtrOutput) Elem() RetentionScheduleFormatOutput {
	return o.ApplyT(func(v *RetentionScheduleFormat) RetentionScheduleFormat {
		if v != nil {
			return *v
		}
		var ret RetentionScheduleFormat
		return ret
	}).(RetentionScheduleFormatOutput)
}

func (o RetentionScheduleFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RetentionScheduleFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RetentionScheduleFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RetentionScheduleFormatInput interface {
	pulumi.Input

	ToRetentionScheduleFormatOutput() RetentionScheduleFormatOutput
	ToRetentionScheduleFormatOutputWithContext(context.Context) RetentionScheduleFormatOutput
}

var retentionScheduleFormatPtrType = reflect.TypeOf((**RetentionScheduleFormat)(nil)).Elem()

type RetentionScheduleFormatPtrInput interface {
	pulumi.Input

	ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput
	ToRetentionScheduleFormatPtrOutputWithContext(context.Context) RetentionScheduleFormatPtrOutput
}

type retentionScheduleFormatPtr string

func RetentionScheduleFormatPtr(v string) RetentionScheduleFormatPtrInput {
	return (*retentionScheduleFormatPtr)(&v)
}

func (*retentionScheduleFormatPtr) ElementType() reflect.Type {
	return retentionScheduleFormatPtrType
}

func (in *retentionScheduleFormatPtr) ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput {
	return pulumi.ToOutput(in).(RetentionScheduleFormatPtrOutput)
}

func (in *retentionScheduleFormatPtr) ToRetentionScheduleFormatPtrOutputWithContext(ctx context.Context) RetentionScheduleFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RetentionScheduleFormatPtrOutput)
}

type ScheduleRunType string

const (
	ScheduleRunTypeInvalid = ScheduleRunType("Invalid")
	ScheduleRunTypeDaily   = ScheduleRunType("Daily")
	ScheduleRunTypeWeekly  = ScheduleRunType("Weekly")
)

func (ScheduleRunType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleRunType)(nil)).Elem()
}

func (e ScheduleRunType) ToScheduleRunTypeOutput() ScheduleRunTypeOutput {
	return pulumi.ToOutput(e).(ScheduleRunTypeOutput)
}

func (e ScheduleRunType) ToScheduleRunTypeOutputWithContext(ctx context.Context) ScheduleRunTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduleRunTypeOutput)
}

func (e ScheduleRunType) ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput {
	return e.ToScheduleRunTypePtrOutputWithContext(context.Background())
}

func (e ScheduleRunType) ToScheduleRunTypePtrOutputWithContext(ctx context.Context) ScheduleRunTypePtrOutput {
	return ScheduleRunType(e).ToScheduleRunTypeOutputWithContext(ctx).ToScheduleRunTypePtrOutputWithContext(ctx)
}

func (e ScheduleRunType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleRunType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleRunType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduleRunType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduleRunTypeOutput struct{ *pulumi.OutputState }

func (ScheduleRunTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleRunType)(nil)).Elem()
}

func (o ScheduleRunTypeOutput) ToScheduleRunTypeOutput() ScheduleRunTypeOutput {
	return o
}

func (o ScheduleRunTypeOutput) ToScheduleRunTypeOutputWithContext(ctx context.Context) ScheduleRunTypeOutput {
	return o
}

func (o ScheduleRunTypeOutput) ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput {
	return o.ToScheduleRunTypePtrOutputWithContext(context.Background())
}

func (o ScheduleRunTypeOutput) ToScheduleRunTypePtrOutputWithContext(ctx context.Context) ScheduleRunTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleRunType) *ScheduleRunType {
		return &v
	}).(ScheduleRunTypePtrOutput)
}

func (o ScheduleRunTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduleRunTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleRunType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduleRunTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleRunTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleRunType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduleRunTypePtrOutput struct{ *pulumi.OutputState }

func (ScheduleRunTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleRunType)(nil)).Elem()
}

func (o ScheduleRunTypePtrOutput) ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput {
	return o
}

func (o ScheduleRunTypePtrOutput) ToScheduleRunTypePtrOutputWithContext(ctx context.Context) ScheduleRunTypePtrOutput {
	return o
}

func (o ScheduleRunTypePtrOutput) Elem() ScheduleRunTypeOutput {
	return o.ApplyT(func(v *ScheduleRunType) ScheduleRunType {
		if v != nil {
			return *v
		}
		var ret ScheduleRunType
		return ret
	}).(ScheduleRunTypeOutput)
}

func (o ScheduleRunTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleRunTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduleRunType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScheduleRunTypeInput interface {
	pulumi.Input

	ToScheduleRunTypeOutput() ScheduleRunTypeOutput
	ToScheduleRunTypeOutputWithContext(context.Context) ScheduleRunTypeOutput
}

var scheduleRunTypePtrType = reflect.TypeOf((**ScheduleRunType)(nil)).Elem()

type ScheduleRunTypePtrInput interface {
	pulumi.Input

	ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput
	ToScheduleRunTypePtrOutputWithContext(context.Context) ScheduleRunTypePtrOutput
}

type scheduleRunTypePtr string

func ScheduleRunTypePtr(v string) ScheduleRunTypePtrInput {
	return (*scheduleRunTypePtr)(&v)
}

func (*scheduleRunTypePtr) ElementType() reflect.Type {
	return scheduleRunTypePtrType
}

func (in *scheduleRunTypePtr) ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput {
	return pulumi.ToOutput(in).(ScheduleRunTypePtrOutput)
}

func (in *scheduleRunTypePtr) ToScheduleRunTypePtrOutputWithContext(ctx context.Context) ScheduleRunTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduleRunTypePtrOutput)
}

type WeekOfMonth string

const (
	WeekOfMonthFirst   = WeekOfMonth("First")
	WeekOfMonthSecond  = WeekOfMonth("Second")
	WeekOfMonthThird   = WeekOfMonth("Third")
	WeekOfMonthFourth  = WeekOfMonth("Fourth")
	WeekOfMonthLast    = WeekOfMonth("Last")
	WeekOfMonthInvalid = WeekOfMonth("Invalid")
)

func (WeekOfMonth) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekOfMonth)(nil)).Elem()
}

func (e WeekOfMonth) ToWeekOfMonthOutput() WeekOfMonthOutput {
	return pulumi.ToOutput(e).(WeekOfMonthOutput)
}

func (e WeekOfMonth) ToWeekOfMonthOutputWithContext(ctx context.Context) WeekOfMonthOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WeekOfMonthOutput)
}

func (e WeekOfMonth) ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput {
	return e.ToWeekOfMonthPtrOutputWithContext(context.Background())
}

func (e WeekOfMonth) ToWeekOfMonthPtrOutputWithContext(ctx context.Context) WeekOfMonthPtrOutput {
	return WeekOfMonth(e).ToWeekOfMonthOutputWithContext(ctx).ToWeekOfMonthPtrOutputWithContext(ctx)
}

func (e WeekOfMonth) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekOfMonth) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekOfMonth) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WeekOfMonth) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WeekOfMonthOutput struct{ *pulumi.OutputState }

func (WeekOfMonthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekOfMonth)(nil)).Elem()
}

func (o WeekOfMonthOutput) ToWeekOfMonthOutput() WeekOfMonthOutput {
	return o
}

func (o WeekOfMonthOutput) ToWeekOfMonthOutputWithContext(ctx context.Context) WeekOfMonthOutput {
	return o
}

func (o WeekOfMonthOutput) ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput {
	return o.ToWeekOfMonthPtrOutputWithContext(context.Background())
}

func (o WeekOfMonthOutput) ToWeekOfMonthPtrOutputWithContext(ctx context.Context) WeekOfMonthPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeekOfMonth) *WeekOfMonth {
		return &v
	}).(WeekOfMonthPtrOutput)
}

func (o WeekOfMonthOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WeekOfMonthOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WeekOfMonth) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WeekOfMonthOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WeekOfMonthOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WeekOfMonth) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WeekOfMonthPtrOutput struct{ *pulumi.OutputState }

func (WeekOfMonthPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekOfMonth)(nil)).Elem()
}

func (o WeekOfMonthPtrOutput) ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput {
	return o
}

func (o WeekOfMonthPtrOutput) ToWeekOfMonthPtrOutputWithContext(ctx context.Context) WeekOfMonthPtrOutput {
	return o
}

func (o WeekOfMonthPtrOutput) Elem() WeekOfMonthOutput {
	return o.ApplyT(func(v *WeekOfMonth) WeekOfMonth {
		if v != nil {
			return *v
		}
		var ret WeekOfMonth
		return ret
	}).(WeekOfMonthOutput)
}

func (o WeekOfMonthPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WeekOfMonthPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WeekOfMonth) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WeekOfMonthInput interface {
	pulumi.Input

	ToWeekOfMonthOutput() WeekOfMonthOutput
	ToWeekOfMonthOutputWithContext(context.Context) WeekOfMonthOutput
}

var weekOfMonthPtrType = reflect.TypeOf((**WeekOfMonth)(nil)).Elem()

type WeekOfMonthPtrInput interface {
	pulumi.Input

	ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput
	ToWeekOfMonthPtrOutputWithContext(context.Context) WeekOfMonthPtrOutput
}

type weekOfMonthPtr string

func WeekOfMonthPtr(v string) WeekOfMonthPtrInput {
	return (*weekOfMonthPtr)(&v)
}

func (*weekOfMonthPtr) ElementType() reflect.Type {
	return weekOfMonthPtrType
}

func (in *weekOfMonthPtr) ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput {
	return pulumi.ToOutput(in).(WeekOfMonthPtrOutput)
}

func (in *weekOfMonthPtr) ToWeekOfMonthPtrOutputWithContext(ctx context.Context) WeekOfMonthPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WeekOfMonthPtrOutput)
}





type WeekOfMonthArrayInput interface {
	pulumi.Input

	ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput
	ToWeekOfMonthArrayOutputWithContext(context.Context) WeekOfMonthArrayOutput
}

type WeekOfMonthArray []WeekOfMonth

func (WeekOfMonthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WeekOfMonth)(nil)).Elem()
}

func (i WeekOfMonthArray) ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput {
	return i.ToWeekOfMonthArrayOutputWithContext(context.Background())
}

func (i WeekOfMonthArray) ToWeekOfMonthArrayOutputWithContext(ctx context.Context) WeekOfMonthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekOfMonthArrayOutput)
}

type WeekOfMonthArrayOutput struct{ *pulumi.OutputState }

func (WeekOfMonthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WeekOfMonth)(nil)).Elem()
}

func (o WeekOfMonthArrayOutput) ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput {
	return o
}

func (o WeekOfMonthArrayOutput) ToWeekOfMonthArrayOutputWithContext(ctx context.Context) WeekOfMonthArrayOutput {
	return o
}

func (o WeekOfMonthArrayOutput) Index(i pulumi.IntInput) WeekOfMonthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WeekOfMonth {
		return vs[0].([]WeekOfMonth)[vs[1].(int)]
	}).(WeekOfMonthOutput)
}

type WorkloadItemType string

const (
	WorkloadItemTypeInvalid         = WorkloadItemType("Invalid")
	WorkloadItemTypeSQLInstance     = WorkloadItemType("SQLInstance")
	WorkloadItemTypeSQLDataBase     = WorkloadItemType("SQLDataBase")
	WorkloadItemTypeSAPHanaSystem   = WorkloadItemType("SAPHanaSystem")
	WorkloadItemTypeSAPHanaDatabase = WorkloadItemType("SAPHanaDatabase")
	WorkloadItemTypeSAPAseSystem    = WorkloadItemType("SAPAseSystem")
	WorkloadItemTypeSAPAseDatabase  = WorkloadItemType("SAPAseDatabase")
)

func (WorkloadItemType) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadItemType)(nil)).Elem()
}

func (e WorkloadItemType) ToWorkloadItemTypeOutput() WorkloadItemTypeOutput {
	return pulumi.ToOutput(e).(WorkloadItemTypeOutput)
}

func (e WorkloadItemType) ToWorkloadItemTypeOutputWithContext(ctx context.Context) WorkloadItemTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkloadItemTypeOutput)
}

func (e WorkloadItemType) ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput {
	return e.ToWorkloadItemTypePtrOutputWithContext(context.Background())
}

func (e WorkloadItemType) ToWorkloadItemTypePtrOutputWithContext(ctx context.Context) WorkloadItemTypePtrOutput {
	return WorkloadItemType(e).ToWorkloadItemTypeOutputWithContext(ctx).ToWorkloadItemTypePtrOutputWithContext(ctx)
}

func (e WorkloadItemType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadItemType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadItemType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkloadItemType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkloadItemTypeOutput struct{ *pulumi.OutputState }

func (WorkloadItemTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadItemType)(nil)).Elem()
}

func (o WorkloadItemTypeOutput) ToWorkloadItemTypeOutput() WorkloadItemTypeOutput {
	return o
}

func (o WorkloadItemTypeOutput) ToWorkloadItemTypeOutputWithContext(ctx context.Context) WorkloadItemTypeOutput {
	return o
}

func (o WorkloadItemTypeOutput) ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput {
	return o.ToWorkloadItemTypePtrOutputWithContext(context.Background())
}

func (o WorkloadItemTypeOutput) ToWorkloadItemTypePtrOutputWithContext(ctx context.Context) WorkloadItemTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadItemType) *WorkloadItemType {
		return &v
	}).(WorkloadItemTypePtrOutput)
}

func (o WorkloadItemTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkloadItemTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadItemType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkloadItemTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadItemTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadItemType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkloadItemTypePtrOutput struct{ *pulumi.OutputState }

func (WorkloadItemTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadItemType)(nil)).Elem()
}

func (o WorkloadItemTypePtrOutput) ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput {
	return o
}

func (o WorkloadItemTypePtrOutput) ToWorkloadItemTypePtrOutputWithContext(ctx context.Context) WorkloadItemTypePtrOutput {
	return o
}

func (o WorkloadItemTypePtrOutput) Elem() WorkloadItemTypeOutput {
	return o.ApplyT(func(v *WorkloadItemType) WorkloadItemType {
		if v != nil {
			return *v
		}
		var ret WorkloadItemType
		return ret
	}).(WorkloadItemTypeOutput)
}

func (o WorkloadItemTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadItemTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkloadItemType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WorkloadItemTypeInput interface {
	pulumi.Input

	ToWorkloadItemTypeOutput() WorkloadItemTypeOutput
	ToWorkloadItemTypeOutputWithContext(context.Context) WorkloadItemTypeOutput
}

var workloadItemTypePtrType = reflect.TypeOf((**WorkloadItemType)(nil)).Elem()

type WorkloadItemTypePtrInput interface {
	pulumi.Input

	ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput
	ToWorkloadItemTypePtrOutputWithContext(context.Context) WorkloadItemTypePtrOutput
}

type workloadItemTypePtr string

func WorkloadItemTypePtr(v string) WorkloadItemTypePtrInput {
	return (*workloadItemTypePtr)(&v)
}

func (*workloadItemTypePtr) ElementType() reflect.Type {
	return workloadItemTypePtrType
}

func (in *workloadItemTypePtr) ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput {
	return pulumi.ToOutput(in).(WorkloadItemTypePtrOutput)
}

func (in *workloadItemTypePtr) ToWorkloadItemTypePtrOutputWithContext(ctx context.Context) WorkloadItemTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkloadItemTypePtrOutput)
}

type WorkloadType string

const (
	WorkloadTypeInvalid           = WorkloadType("Invalid")
	WorkloadTypeVM                = WorkloadType("VM")
	WorkloadTypeFileFolder        = WorkloadType("FileFolder")
	WorkloadTypeAzureSqlDb        = WorkloadType("AzureSqlDb")
	WorkloadTypeSQLDB             = WorkloadType("SQLDB")
	WorkloadTypeExchange          = WorkloadType("Exchange")
	WorkloadTypeSharepoint        = WorkloadType("Sharepoint")
	WorkloadTypeVMwareVM          = WorkloadType("VMwareVM")
	WorkloadTypeSystemState       = WorkloadType("SystemState")
	WorkloadTypeClient            = WorkloadType("Client")
	WorkloadTypeGenericDataSource = WorkloadType("GenericDataSource")
	WorkloadTypeSQLDataBase       = WorkloadType("SQLDataBase")
	WorkloadTypeAzureFileShare    = WorkloadType("AzureFileShare")
	WorkloadTypeSAPHanaDatabase   = WorkloadType("SAPHanaDatabase")
	WorkloadTypeSAPAseDatabase    = WorkloadType("SAPAseDatabase")
)

func (WorkloadType) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadType)(nil)).Elem()
}

func (e WorkloadType) ToWorkloadTypeOutput() WorkloadTypeOutput {
	return pulumi.ToOutput(e).(WorkloadTypeOutput)
}

func (e WorkloadType) ToWorkloadTypeOutputWithContext(ctx context.Context) WorkloadTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkloadTypeOutput)
}

func (e WorkloadType) ToWorkloadTypePtrOutput() WorkloadTypePtrOutput {
	return e.ToWorkloadTypePtrOutputWithContext(context.Background())
}

func (e WorkloadType) ToWorkloadTypePtrOutputWithContext(ctx context.Context) WorkloadTypePtrOutput {
	return WorkloadType(e).ToWorkloadTypeOutputWithContext(ctx).ToWorkloadTypePtrOutputWithContext(ctx)
}

func (e WorkloadType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkloadType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkloadTypeOutput struct{ *pulumi.OutputState }

func (WorkloadTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadType)(nil)).Elem()
}

func (o WorkloadTypeOutput) ToWorkloadTypeOutput() WorkloadTypeOutput {
	return o
}

func (o WorkloadTypeOutput) ToWorkloadTypeOutputWithContext(ctx context.Context) WorkloadTypeOutput {
	return o
}

func (o WorkloadTypeOutput) ToWorkloadTypePtrOutput() WorkloadTypePtrOutput {
	return o.ToWorkloadTypePtrOutputWithContext(context.Background())
}

func (o WorkloadTypeOutput) ToWorkloadTypePtrOutputWithContext(ctx context.Context) WorkloadTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadType) *WorkloadType {
		return &v
	}).(WorkloadTypePtrOutput)
}

func (o WorkloadTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkloadTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkloadTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkloadTypePtrOutput struct{ *pulumi.OutputState }

func (WorkloadTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadType)(nil)).Elem()
}

func (o WorkloadTypePtrOutput) ToWorkloadTypePtrOutput() WorkloadTypePtrOutput {
	return o
}

func (o WorkloadTypePtrOutput) ToWorkloadTypePtrOutputWithContext(ctx context.Context) WorkloadTypePtrOutput {
	return o
}

func (o WorkloadTypePtrOutput) Elem() WorkloadTypeOutput {
	return o.ApplyT(func(v *WorkloadType) WorkloadType {
		if v != nil {
			return *v
		}
		var ret WorkloadType
		return ret
	}).(WorkloadTypeOutput)
}

func (o WorkloadTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkloadType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WorkloadTypeInput interface {
	pulumi.Input

	ToWorkloadTypeOutput() WorkloadTypeOutput
	ToWorkloadTypeOutputWithContext(context.Context) WorkloadTypeOutput
}

var workloadTypePtrType = reflect.TypeOf((**WorkloadType)(nil)).Elem()

type WorkloadTypePtrInput interface {
	pulumi.Input

	ToWorkloadTypePtrOutput() WorkloadTypePtrOutput
	ToWorkloadTypePtrOutputWithContext(context.Context) WorkloadTypePtrOutput
}

type workloadTypePtr string

func WorkloadTypePtr(v string) WorkloadTypePtrInput {
	return (*workloadTypePtr)(&v)
}

func (*workloadTypePtr) ElementType() reflect.Type {
	return workloadTypePtrType
}

func (in *workloadTypePtr) ToWorkloadTypePtrOutput() WorkloadTypePtrOutput {
	return pulumi.ToOutput(in).(WorkloadTypePtrOutput)
}

func (in *workloadTypePtr) ToWorkloadTypePtrOutputWithContext(ctx context.Context) WorkloadTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkloadTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupItemTypeOutput{})
	pulumi.RegisterOutputType(BackupItemTypePtrOutput{})
	pulumi.RegisterOutputType(BackupManagementTypeOutput{})
	pulumi.RegisterOutputType(BackupManagementTypePtrOutput{})
	pulumi.RegisterOutputType(ContainerTypeOutput{})
	pulumi.RegisterOutputType(ContainerTypePtrOutput{})
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
	pulumi.RegisterOutputType(DataSourceTypeOutput{})
	pulumi.RegisterOutputType(DataSourceTypePtrOutput{})
	pulumi.RegisterOutputType(DayOfWeekOutput{})
	pulumi.RegisterOutputType(DayOfWeekPtrOutput{})
	pulumi.RegisterOutputType(DayOfWeekArrayOutput{})
	pulumi.RegisterOutputType(HealthStatusOutput{})
	pulumi.RegisterOutputType(HealthStatusPtrOutput{})
	pulumi.RegisterOutputType(LastBackupStatusOutput{})
	pulumi.RegisterOutputType(LastBackupStatusPtrOutput{})
	pulumi.RegisterOutputType(MonthOfYearOutput{})
	pulumi.RegisterOutputType(MonthOfYearPtrOutput{})
	pulumi.RegisterOutputType(MonthOfYearArrayOutput{})
	pulumi.RegisterOutputType(OperationTypeOutput{})
	pulumi.RegisterOutputType(OperationTypePtrOutput{})
	pulumi.RegisterOutputType(PolicyTypeOutput{})
	pulumi.RegisterOutputType(PolicyTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ProtectedItemHealthStatusOutput{})
	pulumi.RegisterOutputType(ProtectedItemHealthStatusPtrOutput{})
	pulumi.RegisterOutputType(ProtectedItemStateEnumOutput{})
	pulumi.RegisterOutputType(ProtectedItemStateEnumPtrOutput{})
	pulumi.RegisterOutputType(ProtectionStateOutput{})
	pulumi.RegisterOutputType(ProtectionStatePtrOutput{})
	pulumi.RegisterOutputType(ProtectionStatusOutput{})
	pulumi.RegisterOutputType(ProtectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(ResourceHealthStatusOutput{})
	pulumi.RegisterOutputType(ResourceHealthStatusPtrOutput{})
	pulumi.RegisterOutputType(RetentionDurationTypeOutput{})
	pulumi.RegisterOutputType(RetentionDurationTypePtrOutput{})
	pulumi.RegisterOutputType(RetentionScheduleFormatOutput{})
	pulumi.RegisterOutputType(RetentionScheduleFormatPtrOutput{})
	pulumi.RegisterOutputType(ScheduleRunTypeOutput{})
	pulumi.RegisterOutputType(ScheduleRunTypePtrOutput{})
	pulumi.RegisterOutputType(WeekOfMonthOutput{})
	pulumi.RegisterOutputType(WeekOfMonthPtrOutput{})
	pulumi.RegisterOutputType(WeekOfMonthArrayOutput{})
	pulumi.RegisterOutputType(WorkloadItemTypeOutput{})
	pulumi.RegisterOutputType(WorkloadItemTypePtrOutput{})
	pulumi.RegisterOutputType(WorkloadTypeOutput{})
	pulumi.RegisterOutputType(WorkloadTypePtrOutput{})
}
