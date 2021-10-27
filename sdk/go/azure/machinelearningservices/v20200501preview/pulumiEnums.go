


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeEnvironmentType string

const (
	ComputeEnvironmentTypeACI = ComputeEnvironmentType("ACI")
	ComputeEnvironmentTypeAKS = ComputeEnvironmentType("AKS")
)

func (ComputeEnvironmentType) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeEnvironmentType)(nil)).Elem()
}

func (e ComputeEnvironmentType) ToComputeEnvironmentTypeOutput() ComputeEnvironmentTypeOutput {
	return pulumi.ToOutput(e).(ComputeEnvironmentTypeOutput)
}

func (e ComputeEnvironmentType) ToComputeEnvironmentTypeOutputWithContext(ctx context.Context) ComputeEnvironmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeEnvironmentTypeOutput)
}

func (e ComputeEnvironmentType) ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput {
	return e.ToComputeEnvironmentTypePtrOutputWithContext(context.Background())
}

func (e ComputeEnvironmentType) ToComputeEnvironmentTypePtrOutputWithContext(ctx context.Context) ComputeEnvironmentTypePtrOutput {
	return ComputeEnvironmentType(e).ToComputeEnvironmentTypeOutputWithContext(ctx).ToComputeEnvironmentTypePtrOutputWithContext(ctx)
}

func (e ComputeEnvironmentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeEnvironmentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeEnvironmentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeEnvironmentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeEnvironmentTypeOutput struct{ *pulumi.OutputState }

func (ComputeEnvironmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeEnvironmentType)(nil)).Elem()
}

func (o ComputeEnvironmentTypeOutput) ToComputeEnvironmentTypeOutput() ComputeEnvironmentTypeOutput {
	return o
}

func (o ComputeEnvironmentTypeOutput) ToComputeEnvironmentTypeOutputWithContext(ctx context.Context) ComputeEnvironmentTypeOutput {
	return o
}

func (o ComputeEnvironmentTypeOutput) ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput {
	return o.ToComputeEnvironmentTypePtrOutputWithContext(context.Background())
}

func (o ComputeEnvironmentTypeOutput) ToComputeEnvironmentTypePtrOutputWithContext(ctx context.Context) ComputeEnvironmentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeEnvironmentType) *ComputeEnvironmentType {
		return &v
	}).(ComputeEnvironmentTypePtrOutput)
}

func (o ComputeEnvironmentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeEnvironmentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeEnvironmentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeEnvironmentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeEnvironmentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeEnvironmentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeEnvironmentTypePtrOutput struct{ *pulumi.OutputState }

func (ComputeEnvironmentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeEnvironmentType)(nil)).Elem()
}

func (o ComputeEnvironmentTypePtrOutput) ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput {
	return o
}

func (o ComputeEnvironmentTypePtrOutput) ToComputeEnvironmentTypePtrOutputWithContext(ctx context.Context) ComputeEnvironmentTypePtrOutput {
	return o
}

func (o ComputeEnvironmentTypePtrOutput) Elem() ComputeEnvironmentTypeOutput {
	return o.ApplyT(func(v *ComputeEnvironmentType) ComputeEnvironmentType {
		if v != nil {
			return *v
		}
		var ret ComputeEnvironmentType
		return ret
	}).(ComputeEnvironmentTypeOutput)
}

func (o ComputeEnvironmentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeEnvironmentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeEnvironmentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeEnvironmentTypeInput interface {
	pulumi.Input

	ToComputeEnvironmentTypeOutput() ComputeEnvironmentTypeOutput
	ToComputeEnvironmentTypeOutputWithContext(context.Context) ComputeEnvironmentTypeOutput
}

var computeEnvironmentTypePtrType = reflect.TypeOf((**ComputeEnvironmentType)(nil)).Elem()

type ComputeEnvironmentTypePtrInput interface {
	pulumi.Input

	ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput
	ToComputeEnvironmentTypePtrOutputWithContext(context.Context) ComputeEnvironmentTypePtrOutput
}

type computeEnvironmentTypePtr string

func ComputeEnvironmentTypePtr(v string) ComputeEnvironmentTypePtrInput {
	return (*computeEnvironmentTypePtr)(&v)
}

func (*computeEnvironmentTypePtr) ElementType() reflect.Type {
	return computeEnvironmentTypePtrType
}

func (in *computeEnvironmentTypePtr) ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput {
	return pulumi.ToOutput(in).(ComputeEnvironmentTypePtrOutput)
}

func (in *computeEnvironmentTypePtr) ToComputeEnvironmentTypePtrOutputWithContext(ctx context.Context) ComputeEnvironmentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeEnvironmentTypePtrOutput)
}

type ComputeType string

const (
	ComputeTypeAKS               = ComputeType("AKS")
	ComputeTypeAmlCompute        = ComputeType("AmlCompute")
	ComputeTypeDataFactory       = ComputeType("DataFactory")
	ComputeTypeVirtualMachine    = ComputeType("VirtualMachine")
	ComputeTypeHDInsight         = ComputeType("HDInsight")
	ComputeTypeDatabricks        = ComputeType("Databricks")
	ComputeTypeDataLakeAnalytics = ComputeType("DataLakeAnalytics")
)

func (ComputeType) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeType)(nil)).Elem()
}

func (e ComputeType) ToComputeTypeOutput() ComputeTypeOutput {
	return pulumi.ToOutput(e).(ComputeTypeOutput)
}

func (e ComputeType) ToComputeTypeOutputWithContext(ctx context.Context) ComputeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeTypeOutput)
}

func (e ComputeType) ToComputeTypePtrOutput() ComputeTypePtrOutput {
	return e.ToComputeTypePtrOutputWithContext(context.Background())
}

func (e ComputeType) ToComputeTypePtrOutputWithContext(ctx context.Context) ComputeTypePtrOutput {
	return ComputeType(e).ToComputeTypeOutputWithContext(ctx).ToComputeTypePtrOutputWithContext(ctx)
}

func (e ComputeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeTypeOutput struct{ *pulumi.OutputState }

func (ComputeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeType)(nil)).Elem()
}

func (o ComputeTypeOutput) ToComputeTypeOutput() ComputeTypeOutput {
	return o
}

func (o ComputeTypeOutput) ToComputeTypeOutputWithContext(ctx context.Context) ComputeTypeOutput {
	return o
}

func (o ComputeTypeOutput) ToComputeTypePtrOutput() ComputeTypePtrOutput {
	return o.ToComputeTypePtrOutputWithContext(context.Background())
}

func (o ComputeTypeOutput) ToComputeTypePtrOutputWithContext(ctx context.Context) ComputeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeType) *ComputeType {
		return &v
	}).(ComputeTypePtrOutput)
}

func (o ComputeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeTypePtrOutput struct{ *pulumi.OutputState }

func (ComputeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeType)(nil)).Elem()
}

func (o ComputeTypePtrOutput) ToComputeTypePtrOutput() ComputeTypePtrOutput {
	return o
}

func (o ComputeTypePtrOutput) ToComputeTypePtrOutputWithContext(ctx context.Context) ComputeTypePtrOutput {
	return o
}

func (o ComputeTypePtrOutput) Elem() ComputeTypeOutput {
	return o.ApplyT(func(v *ComputeType) ComputeType {
		if v != nil {
			return *v
		}
		var ret ComputeType
		return ret
	}).(ComputeTypeOutput)
}

func (o ComputeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeTypeInput interface {
	pulumi.Input

	ToComputeTypeOutput() ComputeTypeOutput
	ToComputeTypeOutputWithContext(context.Context) ComputeTypeOutput
}

var computeTypePtrType = reflect.TypeOf((**ComputeType)(nil)).Elem()

type ComputeTypePtrInput interface {
	pulumi.Input

	ToComputeTypePtrOutput() ComputeTypePtrOutput
	ToComputeTypePtrOutputWithContext(context.Context) ComputeTypePtrOutput
}

type computeTypePtr string

func ComputeTypePtr(v string) ComputeTypePtrInput {
	return (*computeTypePtr)(&v)
}

func (*computeTypePtr) ElementType() reflect.Type {
	return computeTypePtrType
}

func (in *computeTypePtr) ToComputeTypePtrOutput() ComputeTypePtrOutput {
	return pulumi.ToOutput(in).(ComputeTypePtrOutput)
}

func (in *computeTypePtr) ToComputeTypePtrOutputWithContext(ctx context.Context) ComputeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeTypePtrOutput)
}

type DatasetType string

const (
	DatasetTypeTabular = DatasetType("tabular")
	DatasetTypeFile    = DatasetType("file")
)

func (DatasetType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetType)(nil)).Elem()
}

func (e DatasetType) ToDatasetTypeOutput() DatasetTypeOutput {
	return pulumi.ToOutput(e).(DatasetTypeOutput)
}

func (e DatasetType) ToDatasetTypeOutputWithContext(ctx context.Context) DatasetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatasetTypeOutput)
}

func (e DatasetType) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return e.ToDatasetTypePtrOutputWithContext(context.Background())
}

func (e DatasetType) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return DatasetType(e).ToDatasetTypeOutputWithContext(ctx).ToDatasetTypePtrOutputWithContext(ctx)
}

func (e DatasetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatasetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatasetTypeOutput struct{ *pulumi.OutputState }

func (DatasetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetType)(nil)).Elem()
}

func (o DatasetTypeOutput) ToDatasetTypeOutput() DatasetTypeOutput {
	return o
}

func (o DatasetTypeOutput) ToDatasetTypeOutputWithContext(ctx context.Context) DatasetTypeOutput {
	return o
}

func (o DatasetTypeOutput) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return o.ToDatasetTypePtrOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetType) *DatasetType {
		return &v
	}).(DatasetTypePtrOutput)
}

func (o DatasetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatasetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatasetTypePtrOutput struct{ *pulumi.OutputState }

func (DatasetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetType)(nil)).Elem()
}

func (o DatasetTypePtrOutput) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return o
}

func (o DatasetTypePtrOutput) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return o
}

func (o DatasetTypePtrOutput) Elem() DatasetTypeOutput {
	return o.ApplyT(func(v *DatasetType) DatasetType {
		if v != nil {
			return *v
		}
		var ret DatasetType
		return ret
	}).(DatasetTypeOutput)
}

func (o DatasetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatasetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatasetTypeInput interface {
	pulumi.Input

	ToDatasetTypeOutput() DatasetTypeOutput
	ToDatasetTypeOutputWithContext(context.Context) DatasetTypeOutput
}

var datasetTypePtrType = reflect.TypeOf((**DatasetType)(nil)).Elem()

type DatasetTypePtrInput interface {
	pulumi.Input

	ToDatasetTypePtrOutput() DatasetTypePtrOutput
	ToDatasetTypePtrOutputWithContext(context.Context) DatasetTypePtrOutput
}

type datasetTypePtr string

func DatasetTypePtr(v string) DatasetTypePtrInput {
	return (*datasetTypePtr)(&v)
}

func (*datasetTypePtr) ElementType() reflect.Type {
	return datasetTypePtrType
}

func (in *datasetTypePtr) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return pulumi.ToOutput(in).(DatasetTypePtrOutput)
}

func (in *datasetTypePtr) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatasetTypePtrOutput)
}

type DatastoreTypeArm string

const (
	DatastoreTypeArmBlob       = DatastoreTypeArm("blob")
	DatastoreTypeArmAdls       = DatastoreTypeArm("adls")
	DatastoreTypeArm_Adls_gen2 = DatastoreTypeArm("adls-gen2")
	DatastoreTypeArmDbfs       = DatastoreTypeArm("dbfs")
	DatastoreTypeArmFile       = DatastoreTypeArm("file")
	DatastoreTypeArmMysqldb    = DatastoreTypeArm("mysqldb")
	DatastoreTypeArmSqldb      = DatastoreTypeArm("sqldb")
	DatastoreTypeArmPsqldb     = DatastoreTypeArm("psqldb")
)

func (DatastoreTypeArm) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreTypeArm)(nil)).Elem()
}

func (e DatastoreTypeArm) ToDatastoreTypeArmOutput() DatastoreTypeArmOutput {
	return pulumi.ToOutput(e).(DatastoreTypeArmOutput)
}

func (e DatastoreTypeArm) ToDatastoreTypeArmOutputWithContext(ctx context.Context) DatastoreTypeArmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatastoreTypeArmOutput)
}

func (e DatastoreTypeArm) ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput {
	return e.ToDatastoreTypeArmPtrOutputWithContext(context.Background())
}

func (e DatastoreTypeArm) ToDatastoreTypeArmPtrOutputWithContext(ctx context.Context) DatastoreTypeArmPtrOutput {
	return DatastoreTypeArm(e).ToDatastoreTypeArmOutputWithContext(ctx).ToDatastoreTypeArmPtrOutputWithContext(ctx)
}

func (e DatastoreTypeArm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatastoreTypeArm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatastoreTypeArm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatastoreTypeArm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatastoreTypeArmOutput struct{ *pulumi.OutputState }

func (DatastoreTypeArmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreTypeArm)(nil)).Elem()
}

func (o DatastoreTypeArmOutput) ToDatastoreTypeArmOutput() DatastoreTypeArmOutput {
	return o
}

func (o DatastoreTypeArmOutput) ToDatastoreTypeArmOutputWithContext(ctx context.Context) DatastoreTypeArmOutput {
	return o
}

func (o DatastoreTypeArmOutput) ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput {
	return o.ToDatastoreTypeArmPtrOutputWithContext(context.Background())
}

func (o DatastoreTypeArmOutput) ToDatastoreTypeArmPtrOutputWithContext(ctx context.Context) DatastoreTypeArmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatastoreTypeArm) *DatastoreTypeArm {
		return &v
	}).(DatastoreTypeArmPtrOutput)
}

func (o DatastoreTypeArmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatastoreTypeArmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatastoreTypeArm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatastoreTypeArmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatastoreTypeArmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatastoreTypeArm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatastoreTypeArmPtrOutput struct{ *pulumi.OutputState }

func (DatastoreTypeArmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatastoreTypeArm)(nil)).Elem()
}

func (o DatastoreTypeArmPtrOutput) ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput {
	return o
}

func (o DatastoreTypeArmPtrOutput) ToDatastoreTypeArmPtrOutputWithContext(ctx context.Context) DatastoreTypeArmPtrOutput {
	return o
}

func (o DatastoreTypeArmPtrOutput) Elem() DatastoreTypeArmOutput {
	return o.ApplyT(func(v *DatastoreTypeArm) DatastoreTypeArm {
		if v != nil {
			return *v
		}
		var ret DatastoreTypeArm
		return ret
	}).(DatastoreTypeArmOutput)
}

func (o DatastoreTypeArmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatastoreTypeArmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatastoreTypeArm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatastoreTypeArmInput interface {
	pulumi.Input

	ToDatastoreTypeArmOutput() DatastoreTypeArmOutput
	ToDatastoreTypeArmOutputWithContext(context.Context) DatastoreTypeArmOutput
}

var datastoreTypeArmPtrType = reflect.TypeOf((**DatastoreTypeArm)(nil)).Elem()

type DatastoreTypeArmPtrInput interface {
	pulumi.Input

	ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput
	ToDatastoreTypeArmPtrOutputWithContext(context.Context) DatastoreTypeArmPtrOutput
}

type datastoreTypeArmPtr string

func DatastoreTypeArmPtr(v string) DatastoreTypeArmPtrInput {
	return (*datastoreTypeArmPtr)(&v)
}

func (*datastoreTypeArmPtr) ElementType() reflect.Type {
	return datastoreTypeArmPtrType
}

func (in *datastoreTypeArmPtr) ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput {
	return pulumi.ToOutput(in).(DatastoreTypeArmPtrOutput)
}

func (in *datastoreTypeArmPtr) ToDatastoreTypeArmPtrOutputWithContext(ctx context.Context) DatastoreTypeArmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatastoreTypeArmPtrOutput)
}

type EncryptionStatus string

const (
	EncryptionStatusEnabled  = EncryptionStatus("Enabled")
	EncryptionStatusDisabled = EncryptionStatus("Disabled")
)

func (EncryptionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionStatus)(nil)).Elem()
}

func (e EncryptionStatus) ToEncryptionStatusOutput() EncryptionStatusOutput {
	return pulumi.ToOutput(e).(EncryptionStatusOutput)
}

func (e EncryptionStatus) ToEncryptionStatusOutputWithContext(ctx context.Context) EncryptionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionStatusOutput)
}

func (e EncryptionStatus) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return e.ToEncryptionStatusPtrOutputWithContext(context.Background())
}

func (e EncryptionStatus) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return EncryptionStatus(e).ToEncryptionStatusOutputWithContext(ctx).ToEncryptionStatusPtrOutputWithContext(ctx)
}

func (e EncryptionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionStatusOutput struct{ *pulumi.OutputState }

func (EncryptionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionStatus)(nil)).Elem()
}

func (o EncryptionStatusOutput) ToEncryptionStatusOutput() EncryptionStatusOutput {
	return o
}

func (o EncryptionStatusOutput) ToEncryptionStatusOutputWithContext(ctx context.Context) EncryptionStatusOutput {
	return o
}

func (o EncryptionStatusOutput) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return o.ToEncryptionStatusPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionStatus) *EncryptionStatus {
		return &v
	}).(EncryptionStatusPtrOutput)
}

func (o EncryptionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionStatusPtrOutput struct{ *pulumi.OutputState }

func (EncryptionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionStatus)(nil)).Elem()
}

func (o EncryptionStatusPtrOutput) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return o
}

func (o EncryptionStatusPtrOutput) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return o
}

func (o EncryptionStatusPtrOutput) Elem() EncryptionStatusOutput {
	return o.ApplyT(func(v *EncryptionStatus) EncryptionStatus {
		if v != nil {
			return *v
		}
		var ret EncryptionStatus
		return ret
	}).(EncryptionStatusOutput)
}

func (o EncryptionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionStatusInput interface {
	pulumi.Input

	ToEncryptionStatusOutput() EncryptionStatusOutput
	ToEncryptionStatusOutputWithContext(context.Context) EncryptionStatusOutput
}

var encryptionStatusPtrType = reflect.TypeOf((**EncryptionStatus)(nil)).Elem()

type EncryptionStatusPtrInput interface {
	pulumi.Input

	ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput
	ToEncryptionStatusPtrOutputWithContext(context.Context) EncryptionStatusPtrOutput
}

type encryptionStatusPtr string

func EncryptionStatusPtr(v string) EncryptionStatusPtrInput {
	return (*encryptionStatusPtr)(&v)
}

func (*encryptionStatusPtr) ElementType() reflect.Type {
	return encryptionStatusPtrType
}

func (in *encryptionStatusPtr) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return pulumi.ToOutput(in).(EncryptionStatusPtrOutput)
}

func (in *encryptionStatusPtr) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionStatusPtrOutput)
}

type Header string

const (
	Header_All_files_have_same_headers = Header("all_files_have_same_headers")
	Header_Only_first_file_has_headers = Header("only_first_file_has_headers")
	Header_No_headers                  = Header("no_headers")
	Header_Combine_all_files_headers   = Header("combine_all_files_headers")
)

func (Header) ElementType() reflect.Type {
	return reflect.TypeOf((*Header)(nil)).Elem()
}

func (e Header) ToHeaderOutput() HeaderOutput {
	return pulumi.ToOutput(e).(HeaderOutput)
}

func (e Header) ToHeaderOutputWithContext(ctx context.Context) HeaderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HeaderOutput)
}

func (e Header) ToHeaderPtrOutput() HeaderPtrOutput {
	return e.ToHeaderPtrOutputWithContext(context.Background())
}

func (e Header) ToHeaderPtrOutputWithContext(ctx context.Context) HeaderPtrOutput {
	return Header(e).ToHeaderOutputWithContext(ctx).ToHeaderPtrOutputWithContext(ctx)
}

func (e Header) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Header) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Header) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Header) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HeaderOutput struct{ *pulumi.OutputState }

func (HeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Header)(nil)).Elem()
}

func (o HeaderOutput) ToHeaderOutput() HeaderOutput {
	return o
}

func (o HeaderOutput) ToHeaderOutputWithContext(ctx context.Context) HeaderOutput {
	return o
}

func (o HeaderOutput) ToHeaderPtrOutput() HeaderPtrOutput {
	return o.ToHeaderPtrOutputWithContext(context.Background())
}

func (o HeaderOutput) ToHeaderPtrOutputWithContext(ctx context.Context) HeaderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Header) *Header {
		return &v
	}).(HeaderPtrOutput)
}

func (o HeaderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HeaderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Header) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HeaderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HeaderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Header) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HeaderPtrOutput struct{ *pulumi.OutputState }

func (HeaderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Header)(nil)).Elem()
}

func (o HeaderPtrOutput) ToHeaderPtrOutput() HeaderPtrOutput {
	return o
}

func (o HeaderPtrOutput) ToHeaderPtrOutputWithContext(ctx context.Context) HeaderPtrOutput {
	return o
}

func (o HeaderPtrOutput) Elem() HeaderOutput {
	return o.ApplyT(func(v *Header) Header {
		if v != nil {
			return *v
		}
		var ret Header
		return ret
	}).(HeaderOutput)
}

func (o HeaderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HeaderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Header) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HeaderInput interface {
	pulumi.Input

	ToHeaderOutput() HeaderOutput
	ToHeaderOutputWithContext(context.Context) HeaderOutput
}

var headerPtrType = reflect.TypeOf((**Header)(nil)).Elem()

type HeaderPtrInput interface {
	pulumi.Input

	ToHeaderPtrOutput() HeaderPtrOutput
	ToHeaderPtrOutputWithContext(context.Context) HeaderPtrOutput
}

type headerPtr string

func HeaderPtr(v string) HeaderPtrInput {
	return (*headerPtr)(&v)
}

func (*headerPtr) ElementType() reflect.Type {
	return headerPtrType
}

func (in *headerPtr) ToHeaderPtrOutput() HeaderPtrOutput {
	return pulumi.ToOutput(in).(HeaderPtrOutput)
}

func (in *headerPtr) ToHeaderPtrOutputWithContext(ctx context.Context) HeaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HeaderPtrOutput)
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending      = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved     = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected     = PrivateEndpointServiceConnectionStatus("Rejected")
	PrivateEndpointServiceConnectionStatusDisconnected = PrivateEndpointServiceConnectionStatus("Disconnected")
	PrivateEndpointServiceConnectionStatusTimeout      = PrivateEndpointServiceConnectionStatus("Timeout")
)

func (PrivateEndpointServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return e.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return PrivateEndpointServiceConnectionStatus(e).ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx).ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointServiceConnectionStatus) *PrivateEndpointServiceConnectionStatus {
		return &v
	}).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) Elem() PrivateEndpointServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointServiceConnectionStatus) PrivateEndpointServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointServiceConnectionStatus
		return ret
	}).(PrivateEndpointServiceConnectionStatusOutput)
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput
	ToPrivateEndpointServiceConnectionStatusOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusOutput
}

var privateEndpointServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()

type PrivateEndpointServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput
	ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusPtrOutput
}

type privateEndpointServiceConnectionStatusPtr string

func PrivateEndpointServiceConnectionStatusPtr(v string) PrivateEndpointServiceConnectionStatusPtrInput {
	return (*privateEndpointServiceConnectionStatusPtr)(&v)
}

func (*privateEndpointServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointServiceConnectionStatusPtrType
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

type RemoteLoginPortPublicAccess string

const (
	RemoteLoginPortPublicAccessEnabled      = RemoteLoginPortPublicAccess("Enabled")
	RemoteLoginPortPublicAccessDisabled     = RemoteLoginPortPublicAccess("Disabled")
	RemoteLoginPortPublicAccessNotSpecified = RemoteLoginPortPublicAccess("NotSpecified")
)

func (RemoteLoginPortPublicAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteLoginPortPublicAccess)(nil)).Elem()
}

func (e RemoteLoginPortPublicAccess) ToRemoteLoginPortPublicAccessOutput() RemoteLoginPortPublicAccessOutput {
	return pulumi.ToOutput(e).(RemoteLoginPortPublicAccessOutput)
}

func (e RemoteLoginPortPublicAccess) ToRemoteLoginPortPublicAccessOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RemoteLoginPortPublicAccessOutput)
}

func (e RemoteLoginPortPublicAccess) ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput {
	return e.ToRemoteLoginPortPublicAccessPtrOutputWithContext(context.Background())
}

func (e RemoteLoginPortPublicAccess) ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessPtrOutput {
	return RemoteLoginPortPublicAccess(e).ToRemoteLoginPortPublicAccessOutputWithContext(ctx).ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx)
}

func (e RemoteLoginPortPublicAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RemoteLoginPortPublicAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RemoteLoginPortPublicAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RemoteLoginPortPublicAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RemoteLoginPortPublicAccessOutput struct{ *pulumi.OutputState }

func (RemoteLoginPortPublicAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteLoginPortPublicAccess)(nil)).Elem()
}

func (o RemoteLoginPortPublicAccessOutput) ToRemoteLoginPortPublicAccessOutput() RemoteLoginPortPublicAccessOutput {
	return o
}

func (o RemoteLoginPortPublicAccessOutput) ToRemoteLoginPortPublicAccessOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessOutput {
	return o
}

func (o RemoteLoginPortPublicAccessOutput) ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput {
	return o.ToRemoteLoginPortPublicAccessPtrOutputWithContext(context.Background())
}

func (o RemoteLoginPortPublicAccessOutput) ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemoteLoginPortPublicAccess) *RemoteLoginPortPublicAccess {
		return &v
	}).(RemoteLoginPortPublicAccessPtrOutput)
}

func (o RemoteLoginPortPublicAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RemoteLoginPortPublicAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RemoteLoginPortPublicAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RemoteLoginPortPublicAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RemoteLoginPortPublicAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RemoteLoginPortPublicAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RemoteLoginPortPublicAccessPtrOutput struct{ *pulumi.OutputState }

func (RemoteLoginPortPublicAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteLoginPortPublicAccess)(nil)).Elem()
}

func (o RemoteLoginPortPublicAccessPtrOutput) ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput {
	return o
}

func (o RemoteLoginPortPublicAccessPtrOutput) ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessPtrOutput {
	return o
}

func (o RemoteLoginPortPublicAccessPtrOutput) Elem() RemoteLoginPortPublicAccessOutput {
	return o.ApplyT(func(v *RemoteLoginPortPublicAccess) RemoteLoginPortPublicAccess {
		if v != nil {
			return *v
		}
		var ret RemoteLoginPortPublicAccess
		return ret
	}).(RemoteLoginPortPublicAccessOutput)
}

func (o RemoteLoginPortPublicAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RemoteLoginPortPublicAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RemoteLoginPortPublicAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RemoteLoginPortPublicAccessInput interface {
	pulumi.Input

	ToRemoteLoginPortPublicAccessOutput() RemoteLoginPortPublicAccessOutput
	ToRemoteLoginPortPublicAccessOutputWithContext(context.Context) RemoteLoginPortPublicAccessOutput
}

var remoteLoginPortPublicAccessPtrType = reflect.TypeOf((**RemoteLoginPortPublicAccess)(nil)).Elem()

type RemoteLoginPortPublicAccessPtrInput interface {
	pulumi.Input

	ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput
	ToRemoteLoginPortPublicAccessPtrOutputWithContext(context.Context) RemoteLoginPortPublicAccessPtrOutput
}

type remoteLoginPortPublicAccessPtr string

func RemoteLoginPortPublicAccessPtr(v string) RemoteLoginPortPublicAccessPtrInput {
	return (*remoteLoginPortPublicAccessPtr)(&v)
}

func (*remoteLoginPortPublicAccessPtr) ElementType() reflect.Type {
	return remoteLoginPortPublicAccessPtrType
}

func (in *remoteLoginPortPublicAccessPtr) ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput {
	return pulumi.ToOutput(in).(RemoteLoginPortPublicAccessPtrOutput)
}

func (in *remoteLoginPortPublicAccessPtr) ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RemoteLoginPortPublicAccessPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type SourceType string

const (
	SourceType_Delimited_files  = SourceType("delimited_files")
	SourceType_Json_lines_files = SourceType("json_lines_files")
	SourceType_Parquet_files    = SourceType("parquet_files")
)

func (SourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceType)(nil)).Elem()
}

func (e SourceType) ToSourceTypeOutput() SourceTypeOutput {
	return pulumi.ToOutput(e).(SourceTypeOutput)
}

func (e SourceType) ToSourceTypeOutputWithContext(ctx context.Context) SourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceTypeOutput)
}

func (e SourceType) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return e.ToSourceTypePtrOutputWithContext(context.Background())
}

func (e SourceType) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return SourceType(e).ToSourceTypeOutputWithContext(ctx).ToSourceTypePtrOutputWithContext(ctx)
}

func (e SourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceTypeOutput struct{ *pulumi.OutputState }

func (SourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceType)(nil)).Elem()
}

func (o SourceTypeOutput) ToSourceTypeOutput() SourceTypeOutput {
	return o
}

func (o SourceTypeOutput) ToSourceTypeOutputWithContext(ctx context.Context) SourceTypeOutput {
	return o
}

func (o SourceTypeOutput) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return o.ToSourceTypePtrOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceType) *SourceType {
		return &v
	}).(SourceTypePtrOutput)
}

func (o SourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceTypePtrOutput struct{ *pulumi.OutputState }

func (SourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceType)(nil)).Elem()
}

func (o SourceTypePtrOutput) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return o
}

func (o SourceTypePtrOutput) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return o
}

func (o SourceTypePtrOutput) Elem() SourceTypeOutput {
	return o.ApplyT(func(v *SourceType) SourceType {
		if v != nil {
			return *v
		}
		var ret SourceType
		return ret
	}).(SourceTypeOutput)
}

func (o SourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceTypeInput interface {
	pulumi.Input

	ToSourceTypeOutput() SourceTypeOutput
	ToSourceTypeOutputWithContext(context.Context) SourceTypeOutput
}

var sourceTypePtrType = reflect.TypeOf((**SourceType)(nil)).Elem()

type SourceTypePtrInput interface {
	pulumi.Input

	ToSourceTypePtrOutput() SourceTypePtrOutput
	ToSourceTypePtrOutputWithContext(context.Context) SourceTypePtrOutput
}

type sourceTypePtr string

func SourceTypePtr(v string) SourceTypePtrInput {
	return (*sourceTypePtr)(&v)
}

func (*sourceTypePtr) ElementType() reflect.Type {
	return sourceTypePtrType
}

func (in *sourceTypePtr) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return pulumi.ToOutput(in).(SourceTypePtrOutput)
}

func (in *sourceTypePtr) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceTypePtrOutput)
}

type VariantType string

const (
	VariantTypeControl   = VariantType("Control")
	VariantTypeTreatment = VariantType("Treatment")
)

func (VariantType) ElementType() reflect.Type {
	return reflect.TypeOf((*VariantType)(nil)).Elem()
}

func (e VariantType) ToVariantTypeOutput() VariantTypeOutput {
	return pulumi.ToOutput(e).(VariantTypeOutput)
}

func (e VariantType) ToVariantTypeOutputWithContext(ctx context.Context) VariantTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VariantTypeOutput)
}

func (e VariantType) ToVariantTypePtrOutput() VariantTypePtrOutput {
	return e.ToVariantTypePtrOutputWithContext(context.Background())
}

func (e VariantType) ToVariantTypePtrOutputWithContext(ctx context.Context) VariantTypePtrOutput {
	return VariantType(e).ToVariantTypeOutputWithContext(ctx).ToVariantTypePtrOutputWithContext(ctx)
}

func (e VariantType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VariantType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VariantType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VariantType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VariantTypeOutput struct{ *pulumi.OutputState }

func (VariantTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariantType)(nil)).Elem()
}

func (o VariantTypeOutput) ToVariantTypeOutput() VariantTypeOutput {
	return o
}

func (o VariantTypeOutput) ToVariantTypeOutputWithContext(ctx context.Context) VariantTypeOutput {
	return o
}

func (o VariantTypeOutput) ToVariantTypePtrOutput() VariantTypePtrOutput {
	return o.ToVariantTypePtrOutputWithContext(context.Background())
}

func (o VariantTypeOutput) ToVariantTypePtrOutputWithContext(ctx context.Context) VariantTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VariantType) *VariantType {
		return &v
	}).(VariantTypePtrOutput)
}

func (o VariantTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VariantTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VariantType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VariantTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VariantTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VariantType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VariantTypePtrOutput struct{ *pulumi.OutputState }

func (VariantTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariantType)(nil)).Elem()
}

func (o VariantTypePtrOutput) ToVariantTypePtrOutput() VariantTypePtrOutput {
	return o
}

func (o VariantTypePtrOutput) ToVariantTypePtrOutputWithContext(ctx context.Context) VariantTypePtrOutput {
	return o
}

func (o VariantTypePtrOutput) Elem() VariantTypeOutput {
	return o.ApplyT(func(v *VariantType) VariantType {
		if v != nil {
			return *v
		}
		var ret VariantType
		return ret
	}).(VariantTypeOutput)
}

func (o VariantTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VariantTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VariantType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VariantTypeInput interface {
	pulumi.Input

	ToVariantTypeOutput() VariantTypeOutput
	ToVariantTypeOutputWithContext(context.Context) VariantTypeOutput
}

var variantTypePtrType = reflect.TypeOf((**VariantType)(nil)).Elem()

type VariantTypePtrInput interface {
	pulumi.Input

	ToVariantTypePtrOutput() VariantTypePtrOutput
	ToVariantTypePtrOutputWithContext(context.Context) VariantTypePtrOutput
}

type variantTypePtr string

func VariantTypePtr(v string) VariantTypePtrInput {
	return (*variantTypePtr)(&v)
}

func (*variantTypePtr) ElementType() reflect.Type {
	return variantTypePtrType
}

func (in *variantTypePtr) ToVariantTypePtrOutput() VariantTypePtrOutput {
	return pulumi.ToOutput(in).(VariantTypePtrOutput)
}

func (in *variantTypePtr) ToVariantTypePtrOutputWithContext(ctx context.Context) VariantTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VariantTypePtrOutput)
}

type VmPriority string

const (
	VmPriorityDedicated   = VmPriority("Dedicated")
	VmPriorityLowPriority = VmPriority("LowPriority")
)

func (VmPriority) ElementType() reflect.Type {
	return reflect.TypeOf((*VmPriority)(nil)).Elem()
}

func (e VmPriority) ToVmPriorityOutput() VmPriorityOutput {
	return pulumi.ToOutput(e).(VmPriorityOutput)
}

func (e VmPriority) ToVmPriorityOutputWithContext(ctx context.Context) VmPriorityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VmPriorityOutput)
}

func (e VmPriority) ToVmPriorityPtrOutput() VmPriorityPtrOutput {
	return e.ToVmPriorityPtrOutputWithContext(context.Background())
}

func (e VmPriority) ToVmPriorityPtrOutputWithContext(ctx context.Context) VmPriorityPtrOutput {
	return VmPriority(e).ToVmPriorityOutputWithContext(ctx).ToVmPriorityPtrOutputWithContext(ctx)
}

func (e VmPriority) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VmPriority) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VmPriority) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VmPriority) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VmPriorityOutput struct{ *pulumi.OutputState }

func (VmPriorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmPriority)(nil)).Elem()
}

func (o VmPriorityOutput) ToVmPriorityOutput() VmPriorityOutput {
	return o
}

func (o VmPriorityOutput) ToVmPriorityOutputWithContext(ctx context.Context) VmPriorityOutput {
	return o
}

func (o VmPriorityOutput) ToVmPriorityPtrOutput() VmPriorityPtrOutput {
	return o.ToVmPriorityPtrOutputWithContext(context.Background())
}

func (o VmPriorityOutput) ToVmPriorityPtrOutputWithContext(ctx context.Context) VmPriorityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VmPriority) *VmPriority {
		return &v
	}).(VmPriorityPtrOutput)
}

func (o VmPriorityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VmPriorityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VmPriority) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VmPriorityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VmPriorityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VmPriority) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VmPriorityPtrOutput struct{ *pulumi.OutputState }

func (VmPriorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmPriority)(nil)).Elem()
}

func (o VmPriorityPtrOutput) ToVmPriorityPtrOutput() VmPriorityPtrOutput {
	return o
}

func (o VmPriorityPtrOutput) ToVmPriorityPtrOutputWithContext(ctx context.Context) VmPriorityPtrOutput {
	return o
}

func (o VmPriorityPtrOutput) Elem() VmPriorityOutput {
	return o.ApplyT(func(v *VmPriority) VmPriority {
		if v != nil {
			return *v
		}
		var ret VmPriority
		return ret
	}).(VmPriorityOutput)
}

func (o VmPriorityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VmPriorityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VmPriority) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VmPriorityInput interface {
	pulumi.Input

	ToVmPriorityOutput() VmPriorityOutput
	ToVmPriorityOutputWithContext(context.Context) VmPriorityOutput
}

var vmPriorityPtrType = reflect.TypeOf((**VmPriority)(nil)).Elem()

type VmPriorityPtrInput interface {
	pulumi.Input

	ToVmPriorityPtrOutput() VmPriorityPtrOutput
	ToVmPriorityPtrOutputWithContext(context.Context) VmPriorityPtrOutput
}

type vmPriorityPtr string

func VmPriorityPtr(v string) VmPriorityPtrInput {
	return (*vmPriorityPtr)(&v)
}

func (*vmPriorityPtr) ElementType() reflect.Type {
	return vmPriorityPtrType
}

func (in *vmPriorityPtr) ToVmPriorityPtrOutput() VmPriorityPtrOutput {
	return pulumi.ToOutput(in).(VmPriorityPtrOutput)
}

func (in *vmPriorityPtr) ToVmPriorityPtrOutputWithContext(ctx context.Context) VmPriorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VmPriorityPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeEnvironmentTypeInput)(nil)).Elem(), ComputeEnvironmentType("ACI"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeEnvironmentTypePtrInput)(nil)).Elem(), ComputeEnvironmentType("ACI"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeTypeInput)(nil)).Elem(), ComputeType("AKS"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeTypePtrInput)(nil)).Elem(), ComputeType("AKS"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetTypeInput)(nil)).Elem(), DatasetType("tabular"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetTypePtrInput)(nil)).Elem(), DatasetType("tabular"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreTypeArmInput)(nil)).Elem(), DatastoreTypeArm("blob"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreTypeArmPtrInput)(nil)).Elem(), DatastoreTypeArm("blob"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionStatusInput)(nil)).Elem(), EncryptionStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionStatusPtrInput)(nil)).Elem(), EncryptionStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*HeaderInput)(nil)).Elem(), Header("all_files_have_same_headers"))
	pulumi.RegisterInputType(reflect.TypeOf((*HeaderPtrInput)(nil)).Elem(), Header("all_files_have_same_headers"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteLoginPortPublicAccessInput)(nil)).Elem(), RemoteLoginPortPublicAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteLoginPortPublicAccessPtrInput)(nil)).Elem(), RemoteLoginPortPublicAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTypeInput)(nil)).Elem(), SourceType("delimited_files"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTypePtrInput)(nil)).Elem(), SourceType("delimited_files"))
	pulumi.RegisterInputType(reflect.TypeOf((*VariantTypeInput)(nil)).Elem(), VariantType("Control"))
	pulumi.RegisterInputType(reflect.TypeOf((*VariantTypePtrInput)(nil)).Elem(), VariantType("Control"))
	pulumi.RegisterInputType(reflect.TypeOf((*VmPriorityInput)(nil)).Elem(), VmPriority("Dedicated"))
	pulumi.RegisterInputType(reflect.TypeOf((*VmPriorityPtrInput)(nil)).Elem(), VmPriority("Dedicated"))
	pulumi.RegisterOutputType(ComputeEnvironmentTypeOutput{})
	pulumi.RegisterOutputType(ComputeEnvironmentTypePtrOutput{})
	pulumi.RegisterOutputType(ComputeTypeOutput{})
	pulumi.RegisterOutputType(ComputeTypePtrOutput{})
	pulumi.RegisterOutputType(DatasetTypeOutput{})
	pulumi.RegisterOutputType(DatasetTypePtrOutput{})
	pulumi.RegisterOutputType(DatastoreTypeArmOutput{})
	pulumi.RegisterOutputType(DatastoreTypeArmPtrOutput{})
	pulumi.RegisterOutputType(EncryptionStatusOutput{})
	pulumi.RegisterOutputType(EncryptionStatusPtrOutput{})
	pulumi.RegisterOutputType(HeaderOutput{})
	pulumi.RegisterOutputType(HeaderPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(RemoteLoginPortPublicAccessOutput{})
	pulumi.RegisterOutputType(RemoteLoginPortPublicAccessPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SourceTypeOutput{})
	pulumi.RegisterOutputType(SourceTypePtrOutput{})
	pulumi.RegisterOutputType(VariantTypeOutput{})
	pulumi.RegisterOutputType(VariantTypePtrOutput{})
	pulumi.RegisterOutputType(VmPriorityOutput{})
	pulumi.RegisterOutputType(VmPriorityPtrOutput{})
}
