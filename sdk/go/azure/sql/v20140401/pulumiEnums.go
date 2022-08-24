


package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdministratorType string

const (
	AdministratorTypeActiveDirectory = AdministratorType("ActiveDirectory")
)

type AutoExecuteStatus string

const (
	AutoExecuteStatusEnabled  = AutoExecuteStatus("Enabled")
	AutoExecuteStatusDisabled = AutoExecuteStatus("Disabled")
	AutoExecuteStatusDefault  = AutoExecuteStatus("Default")
)

func (AutoExecuteStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoExecuteStatus)(nil)).Elem()
}

func (e AutoExecuteStatus) ToAutoExecuteStatusOutput() AutoExecuteStatusOutput {
	return pulumi.ToOutput(e).(AutoExecuteStatusOutput)
}

func (e AutoExecuteStatus) ToAutoExecuteStatusOutputWithContext(ctx context.Context) AutoExecuteStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoExecuteStatusOutput)
}

func (e AutoExecuteStatus) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return e.ToAutoExecuteStatusPtrOutputWithContext(context.Background())
}

func (e AutoExecuteStatus) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return AutoExecuteStatus(e).ToAutoExecuteStatusOutputWithContext(ctx).ToAutoExecuteStatusPtrOutputWithContext(ctx)
}

func (e AutoExecuteStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoExecuteStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoExecuteStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoExecuteStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoExecuteStatusOutput struct{ *pulumi.OutputState }

func (AutoExecuteStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoExecuteStatus)(nil)).Elem()
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusOutput() AutoExecuteStatusOutput {
	return o
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusOutputWithContext(ctx context.Context) AutoExecuteStatusOutput {
	return o
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return o.ToAutoExecuteStatusPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoExecuteStatus) *AutoExecuteStatus {
		return &v
	}).(AutoExecuteStatusPtrOutput)
}

func (o AutoExecuteStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoExecuteStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoExecuteStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoExecuteStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoExecuteStatusPtrOutput struct{ *pulumi.OutputState }

func (AutoExecuteStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoExecuteStatus)(nil)).Elem()
}

func (o AutoExecuteStatusPtrOutput) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return o
}

func (o AutoExecuteStatusPtrOutput) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return o
}

func (o AutoExecuteStatusPtrOutput) Elem() AutoExecuteStatusOutput {
	return o.ApplyT(func(v *AutoExecuteStatus) AutoExecuteStatus {
		if v != nil {
			return *v
		}
		var ret AutoExecuteStatus
		return ret
	}).(AutoExecuteStatusOutput)
}

func (o AutoExecuteStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoExecuteStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoExecuteStatusInput interface {
	pulumi.Input

	ToAutoExecuteStatusOutput() AutoExecuteStatusOutput
	ToAutoExecuteStatusOutputWithContext(context.Context) AutoExecuteStatusOutput
}

var autoExecuteStatusPtrType = reflect.TypeOf((**AutoExecuteStatus)(nil)).Elem()

type AutoExecuteStatusPtrInput interface {
	pulumi.Input

	ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput
	ToAutoExecuteStatusPtrOutputWithContext(context.Context) AutoExecuteStatusPtrOutput
}

type autoExecuteStatusPtr string

func AutoExecuteStatusPtr(v string) AutoExecuteStatusPtrInput {
	return (*autoExecuteStatusPtr)(&v)
}

func (*autoExecuteStatusPtr) ElementType() reflect.Type {
	return autoExecuteStatusPtrType
}

func (in *autoExecuteStatusPtr) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return pulumi.ToOutput(in).(AutoExecuteStatusPtrOutput)
}

func (in *autoExecuteStatusPtr) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoExecuteStatusPtrOutput)
}

type CreateMode string

const (
	CreateModeCopy                           = CreateMode("Copy")
	CreateModeDefault                        = CreateMode("Default")
	CreateModeNonReadableSecondary           = CreateMode("NonReadableSecondary")
	CreateModeOnlineSecondary                = CreateMode("OnlineSecondary")
	CreateModePointInTimeRestore             = CreateMode("PointInTimeRestore")
	CreateModeRecovery                       = CreateMode("Recovery")
	CreateModeRestore                        = CreateMode("Restore")
	CreateModeRestoreLongTermRetentionBackup = CreateMode("RestoreLongTermRetentionBackup")
)

type DataMaskingState string

const (
	DataMaskingStateDisabled = DataMaskingState("Disabled")
	DataMaskingStateEnabled  = DataMaskingState("Enabled")
)

func (DataMaskingState) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingState)(nil)).Elem()
}

func (e DataMaskingState) ToDataMaskingStateOutput() DataMaskingStateOutput {
	return pulumi.ToOutput(e).(DataMaskingStateOutput)
}

func (e DataMaskingState) ToDataMaskingStateOutputWithContext(ctx context.Context) DataMaskingStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataMaskingStateOutput)
}

func (e DataMaskingState) ToDataMaskingStatePtrOutput() DataMaskingStatePtrOutput {
	return e.ToDataMaskingStatePtrOutputWithContext(context.Background())
}

func (e DataMaskingState) ToDataMaskingStatePtrOutputWithContext(ctx context.Context) DataMaskingStatePtrOutput {
	return DataMaskingState(e).ToDataMaskingStateOutputWithContext(ctx).ToDataMaskingStatePtrOutputWithContext(ctx)
}

func (e DataMaskingState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataMaskingState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataMaskingState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataMaskingState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataMaskingStateOutput struct{ *pulumi.OutputState }

func (DataMaskingStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingState)(nil)).Elem()
}

func (o DataMaskingStateOutput) ToDataMaskingStateOutput() DataMaskingStateOutput {
	return o
}

func (o DataMaskingStateOutput) ToDataMaskingStateOutputWithContext(ctx context.Context) DataMaskingStateOutput {
	return o
}

func (o DataMaskingStateOutput) ToDataMaskingStatePtrOutput() DataMaskingStatePtrOutput {
	return o.ToDataMaskingStatePtrOutputWithContext(context.Background())
}

func (o DataMaskingStateOutput) ToDataMaskingStatePtrOutputWithContext(ctx context.Context) DataMaskingStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataMaskingState) *DataMaskingState {
		return &v
	}).(DataMaskingStatePtrOutput)
}

func (o DataMaskingStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataMaskingStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataMaskingState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataMaskingStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataMaskingStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataMaskingState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataMaskingStatePtrOutput struct{ *pulumi.OutputState }

func (DataMaskingStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMaskingState)(nil)).Elem()
}

func (o DataMaskingStatePtrOutput) ToDataMaskingStatePtrOutput() DataMaskingStatePtrOutput {
	return o
}

func (o DataMaskingStatePtrOutput) ToDataMaskingStatePtrOutputWithContext(ctx context.Context) DataMaskingStatePtrOutput {
	return o
}

func (o DataMaskingStatePtrOutput) Elem() DataMaskingStateOutput {
	return o.ApplyT(func(v *DataMaskingState) DataMaskingState {
		if v != nil {
			return *v
		}
		var ret DataMaskingState
		return ret
	}).(DataMaskingStateOutput)
}

func (o DataMaskingStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataMaskingStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataMaskingState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataMaskingStateInput interface {
	pulumi.Input

	ToDataMaskingStateOutput() DataMaskingStateOutput
	ToDataMaskingStateOutputWithContext(context.Context) DataMaskingStateOutput
}

var dataMaskingStatePtrType = reflect.TypeOf((**DataMaskingState)(nil)).Elem()

type DataMaskingStatePtrInput interface {
	pulumi.Input

	ToDataMaskingStatePtrOutput() DataMaskingStatePtrOutput
	ToDataMaskingStatePtrOutputWithContext(context.Context) DataMaskingStatePtrOutput
}

type dataMaskingStatePtr string

func DataMaskingStatePtr(v string) DataMaskingStatePtrInput {
	return (*dataMaskingStatePtr)(&v)
}

func (*dataMaskingStatePtr) ElementType() reflect.Type {
	return dataMaskingStatePtrType
}

func (in *dataMaskingStatePtr) ToDataMaskingStatePtrOutput() DataMaskingStatePtrOutput {
	return pulumi.ToOutput(in).(DataMaskingStatePtrOutput)
}

func (in *dataMaskingStatePtr) ToDataMaskingStatePtrOutputWithContext(ctx context.Context) DataMaskingStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataMaskingStatePtrOutput)
}

type DatabaseEdition string

const (
	DatabaseEditionWeb              = DatabaseEdition("Web")
	DatabaseEditionBusiness         = DatabaseEdition("Business")
	DatabaseEditionBasic            = DatabaseEdition("Basic")
	DatabaseEditionStandard         = DatabaseEdition("Standard")
	DatabaseEditionPremium          = DatabaseEdition("Premium")
	DatabaseEditionPremiumRS        = DatabaseEdition("PremiumRS")
	DatabaseEditionFree             = DatabaseEdition("Free")
	DatabaseEditionStretch          = DatabaseEdition("Stretch")
	DatabaseEditionDataWarehouse    = DatabaseEdition("DataWarehouse")
	DatabaseEditionSystem           = DatabaseEdition("System")
	DatabaseEditionSystem2          = DatabaseEdition("System2")
	DatabaseEditionGeneralPurpose   = DatabaseEdition("GeneralPurpose")
	DatabaseEditionBusinessCritical = DatabaseEdition("BusinessCritical")
	DatabaseEditionHyperscale       = DatabaseEdition("Hyperscale")
)

type ElasticPoolEdition string

const (
	ElasticPoolEditionBasic            = ElasticPoolEdition("Basic")
	ElasticPoolEditionStandard         = ElasticPoolEdition("Standard")
	ElasticPoolEditionPremium          = ElasticPoolEdition("Premium")
	ElasticPoolEditionGeneralPurpose   = ElasticPoolEdition("GeneralPurpose")
	ElasticPoolEditionBusinessCritical = ElasticPoolEdition("BusinessCritical")
)

type GeoBackupPolicyStateEnum string

const (
	GeoBackupPolicyStateEnumDisabled = GeoBackupPolicyStateEnum("Disabled")
	GeoBackupPolicyStateEnumEnabled  = GeoBackupPolicyStateEnum("Enabled")
)

func (GeoBackupPolicyStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoBackupPolicyStateEnum)(nil)).Elem()
}

func (e GeoBackupPolicyStateEnum) ToGeoBackupPolicyStateEnumOutput() GeoBackupPolicyStateEnumOutput {
	return pulumi.ToOutput(e).(GeoBackupPolicyStateEnumOutput)
}

func (e GeoBackupPolicyStateEnum) ToGeoBackupPolicyStateEnumOutputWithContext(ctx context.Context) GeoBackupPolicyStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GeoBackupPolicyStateEnumOutput)
}

func (e GeoBackupPolicyStateEnum) ToGeoBackupPolicyStateEnumPtrOutput() GeoBackupPolicyStateEnumPtrOutput {
	return e.ToGeoBackupPolicyStateEnumPtrOutputWithContext(context.Background())
}

func (e GeoBackupPolicyStateEnum) ToGeoBackupPolicyStateEnumPtrOutputWithContext(ctx context.Context) GeoBackupPolicyStateEnumPtrOutput {
	return GeoBackupPolicyStateEnum(e).ToGeoBackupPolicyStateEnumOutputWithContext(ctx).ToGeoBackupPolicyStateEnumPtrOutputWithContext(ctx)
}

func (e GeoBackupPolicyStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoBackupPolicyStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoBackupPolicyStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GeoBackupPolicyStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GeoBackupPolicyStateEnumOutput struct{ *pulumi.OutputState }

func (GeoBackupPolicyStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoBackupPolicyStateEnum)(nil)).Elem()
}

func (o GeoBackupPolicyStateEnumOutput) ToGeoBackupPolicyStateEnumOutput() GeoBackupPolicyStateEnumOutput {
	return o
}

func (o GeoBackupPolicyStateEnumOutput) ToGeoBackupPolicyStateEnumOutputWithContext(ctx context.Context) GeoBackupPolicyStateEnumOutput {
	return o
}

func (o GeoBackupPolicyStateEnumOutput) ToGeoBackupPolicyStateEnumPtrOutput() GeoBackupPolicyStateEnumPtrOutput {
	return o.ToGeoBackupPolicyStateEnumPtrOutputWithContext(context.Background())
}

func (o GeoBackupPolicyStateEnumOutput) ToGeoBackupPolicyStateEnumPtrOutputWithContext(ctx context.Context) GeoBackupPolicyStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeoBackupPolicyStateEnum) *GeoBackupPolicyStateEnum {
		return &v
	}).(GeoBackupPolicyStateEnumPtrOutput)
}

func (o GeoBackupPolicyStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GeoBackupPolicyStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoBackupPolicyStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GeoBackupPolicyStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoBackupPolicyStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoBackupPolicyStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GeoBackupPolicyStateEnumPtrOutput struct{ *pulumi.OutputState }

func (GeoBackupPolicyStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoBackupPolicyStateEnum)(nil)).Elem()
}

func (o GeoBackupPolicyStateEnumPtrOutput) ToGeoBackupPolicyStateEnumPtrOutput() GeoBackupPolicyStateEnumPtrOutput {
	return o
}

func (o GeoBackupPolicyStateEnumPtrOutput) ToGeoBackupPolicyStateEnumPtrOutputWithContext(ctx context.Context) GeoBackupPolicyStateEnumPtrOutput {
	return o
}

func (o GeoBackupPolicyStateEnumPtrOutput) Elem() GeoBackupPolicyStateEnumOutput {
	return o.ApplyT(func(v *GeoBackupPolicyStateEnum) GeoBackupPolicyStateEnum {
		if v != nil {
			return *v
		}
		var ret GeoBackupPolicyStateEnum
		return ret
	}).(GeoBackupPolicyStateEnumOutput)
}

func (o GeoBackupPolicyStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoBackupPolicyStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GeoBackupPolicyStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GeoBackupPolicyStateEnumInput interface {
	pulumi.Input

	ToGeoBackupPolicyStateEnumOutput() GeoBackupPolicyStateEnumOutput
	ToGeoBackupPolicyStateEnumOutputWithContext(context.Context) GeoBackupPolicyStateEnumOutput
}

var geoBackupPolicyStateEnumPtrType = reflect.TypeOf((**GeoBackupPolicyStateEnum)(nil)).Elem()

type GeoBackupPolicyStateEnumPtrInput interface {
	pulumi.Input

	ToGeoBackupPolicyStateEnumPtrOutput() GeoBackupPolicyStateEnumPtrOutput
	ToGeoBackupPolicyStateEnumPtrOutputWithContext(context.Context) GeoBackupPolicyStateEnumPtrOutput
}

type geoBackupPolicyStateEnumPtr string

func GeoBackupPolicyStateEnumPtr(v string) GeoBackupPolicyStateEnumPtrInput {
	return (*geoBackupPolicyStateEnumPtr)(&v)
}

func (*geoBackupPolicyStateEnumPtr) ElementType() reflect.Type {
	return geoBackupPolicyStateEnumPtrType
}

func (in *geoBackupPolicyStateEnumPtr) ToGeoBackupPolicyStateEnumPtrOutput() GeoBackupPolicyStateEnumPtrOutput {
	return pulumi.ToOutput(in).(GeoBackupPolicyStateEnumPtrOutput)
}

func (in *geoBackupPolicyStateEnumPtr) ToGeoBackupPolicyStateEnumPtrOutputWithContext(ctx context.Context) GeoBackupPolicyStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GeoBackupPolicyStateEnumPtrOutput)
}

type ReadScale string

const (
	ReadScaleEnabled  = ReadScale("Enabled")
	ReadScaleDisabled = ReadScale("Disabled")
)

type SampleName string

const (
	SampleNameAdventureWorksLT = SampleName("AdventureWorksLT")
)

type SecurityAlertPolicyEmailAccountAdmins string

const (
	SecurityAlertPolicyEmailAccountAdminsEnabled  = SecurityAlertPolicyEmailAccountAdmins("Enabled")
	SecurityAlertPolicyEmailAccountAdminsDisabled = SecurityAlertPolicyEmailAccountAdmins("Disabled")
)

type SecurityAlertPolicyState string

const (
	SecurityAlertPolicyStateNew      = SecurityAlertPolicyState("New")
	SecurityAlertPolicyStateEnabled  = SecurityAlertPolicyState("Enabled")
	SecurityAlertPolicyStateDisabled = SecurityAlertPolicyState("Disabled")
)

type SecurityAlertPolicyUseServerDefault string

const (
	SecurityAlertPolicyUseServerDefaultEnabled  = SecurityAlertPolicyUseServerDefault("Enabled")
	SecurityAlertPolicyUseServerDefaultDisabled = SecurityAlertPolicyUseServerDefault("Disabled")
)

type ServerVersion string

const (
	ServerVersion_2_0  = ServerVersion("2.0")
	ServerVersion_12_0 = ServerVersion("12.0")
)

type ServiceObjectiveName string

const (
	ServiceObjectiveNameSystem      = ServiceObjectiveName("System")
	ServiceObjectiveNameSystem0     = ServiceObjectiveName("System0")
	ServiceObjectiveNameSystem1     = ServiceObjectiveName("System1")
	ServiceObjectiveNameSystem2     = ServiceObjectiveName("System2")
	ServiceObjectiveNameSystem3     = ServiceObjectiveName("System3")
	ServiceObjectiveNameSystem4     = ServiceObjectiveName("System4")
	ServiceObjectiveNameSystem2L    = ServiceObjectiveName("System2L")
	ServiceObjectiveNameSystem3L    = ServiceObjectiveName("System3L")
	ServiceObjectiveNameSystem4L    = ServiceObjectiveName("System4L")
	ServiceObjectiveNameFree        = ServiceObjectiveName("Free")
	ServiceObjectiveNameBasic       = ServiceObjectiveName("Basic")
	ServiceObjectiveNameS0          = ServiceObjectiveName("S0")
	ServiceObjectiveNameS1          = ServiceObjectiveName("S1")
	ServiceObjectiveNameS2          = ServiceObjectiveName("S2")
	ServiceObjectiveNameS3          = ServiceObjectiveName("S3")
	ServiceObjectiveNameS4          = ServiceObjectiveName("S4")
	ServiceObjectiveNameS6          = ServiceObjectiveName("S6")
	ServiceObjectiveNameS7          = ServiceObjectiveName("S7")
	ServiceObjectiveNameS9          = ServiceObjectiveName("S9")
	ServiceObjectiveNameS12         = ServiceObjectiveName("S12")
	ServiceObjectiveNameP1          = ServiceObjectiveName("P1")
	ServiceObjectiveNameP2          = ServiceObjectiveName("P2")
	ServiceObjectiveNameP3          = ServiceObjectiveName("P3")
	ServiceObjectiveNameP4          = ServiceObjectiveName("P4")
	ServiceObjectiveNameP6          = ServiceObjectiveName("P6")
	ServiceObjectiveNameP11         = ServiceObjectiveName("P11")
	ServiceObjectiveNameP15         = ServiceObjectiveName("P15")
	ServiceObjectiveNamePRS1        = ServiceObjectiveName("PRS1")
	ServiceObjectiveNamePRS2        = ServiceObjectiveName("PRS2")
	ServiceObjectiveNamePRS4        = ServiceObjectiveName("PRS4")
	ServiceObjectiveNamePRS6        = ServiceObjectiveName("PRS6")
	ServiceObjectiveNameDW100       = ServiceObjectiveName("DW100")
	ServiceObjectiveNameDW200       = ServiceObjectiveName("DW200")
	ServiceObjectiveNameDW300       = ServiceObjectiveName("DW300")
	ServiceObjectiveNameDW400       = ServiceObjectiveName("DW400")
	ServiceObjectiveNameDW500       = ServiceObjectiveName("DW500")
	ServiceObjectiveNameDW600       = ServiceObjectiveName("DW600")
	ServiceObjectiveNameDW1000      = ServiceObjectiveName("DW1000")
	ServiceObjectiveNameDW1200      = ServiceObjectiveName("DW1200")
	ServiceObjectiveNameDW1000c     = ServiceObjectiveName("DW1000c")
	ServiceObjectiveNameDW1500      = ServiceObjectiveName("DW1500")
	ServiceObjectiveNameDW1500c     = ServiceObjectiveName("DW1500c")
	ServiceObjectiveNameDW2000      = ServiceObjectiveName("DW2000")
	ServiceObjectiveNameDW2000c     = ServiceObjectiveName("DW2000c")
	ServiceObjectiveNameDW3000      = ServiceObjectiveName("DW3000")
	ServiceObjectiveNameDW2500c     = ServiceObjectiveName("DW2500c")
	ServiceObjectiveNameDW3000c     = ServiceObjectiveName("DW3000c")
	ServiceObjectiveNameDW6000      = ServiceObjectiveName("DW6000")
	ServiceObjectiveNameDW5000c     = ServiceObjectiveName("DW5000c")
	ServiceObjectiveNameDW6000c     = ServiceObjectiveName("DW6000c")
	ServiceObjectiveNameDW7500c     = ServiceObjectiveName("DW7500c")
	ServiceObjectiveNameDW10000c    = ServiceObjectiveName("DW10000c")
	ServiceObjectiveNameDW15000c    = ServiceObjectiveName("DW15000c")
	ServiceObjectiveNameDW30000c    = ServiceObjectiveName("DW30000c")
	ServiceObjectiveNameDS100       = ServiceObjectiveName("DS100")
	ServiceObjectiveNameDS200       = ServiceObjectiveName("DS200")
	ServiceObjectiveNameDS300       = ServiceObjectiveName("DS300")
	ServiceObjectiveNameDS400       = ServiceObjectiveName("DS400")
	ServiceObjectiveNameDS500       = ServiceObjectiveName("DS500")
	ServiceObjectiveNameDS600       = ServiceObjectiveName("DS600")
	ServiceObjectiveNameDS1000      = ServiceObjectiveName("DS1000")
	ServiceObjectiveNameDS1200      = ServiceObjectiveName("DS1200")
	ServiceObjectiveNameDS1500      = ServiceObjectiveName("DS1500")
	ServiceObjectiveNameDS2000      = ServiceObjectiveName("DS2000")
	ServiceObjectiveNameElasticPool = ServiceObjectiveName("ElasticPool")
)

type TransparentDataEncryptionStatus string

const (
	TransparentDataEncryptionStatusEnabled  = TransparentDataEncryptionStatus("Enabled")
	TransparentDataEncryptionStatusDisabled = TransparentDataEncryptionStatus("Disabled")
)

func init() {
	pulumi.RegisterOutputType(AutoExecuteStatusOutput{})
	pulumi.RegisterOutputType(AutoExecuteStatusPtrOutput{})
	pulumi.RegisterOutputType(DataMaskingStateOutput{})
	pulumi.RegisterOutputType(DataMaskingStatePtrOutput{})
	pulumi.RegisterOutputType(GeoBackupPolicyStateEnumOutput{})
	pulumi.RegisterOutputType(GeoBackupPolicyStateEnumPtrOutput{})
}
