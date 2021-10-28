


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataSourceKind string

const (
	DataSourceKindAzureActivityLog              = DataSourceKind("AzureActivityLog")
	DataSourceKindChangeTrackingPath            = DataSourceKind("ChangeTrackingPath")
	DataSourceKindChangeTrackingDefaultPath     = DataSourceKind("ChangeTrackingDefaultPath")
	DataSourceKindChangeTrackingDefaultRegistry = DataSourceKind("ChangeTrackingDefaultRegistry")
	DataSourceKindChangeTrackingCustomRegistry  = DataSourceKind("ChangeTrackingCustomRegistry")
	DataSourceKindCustomLog                     = DataSourceKind("CustomLog")
	DataSourceKindCustomLogCollection           = DataSourceKind("CustomLogCollection")
	DataSourceKindGenericDataSource             = DataSourceKind("GenericDataSource")
	DataSourceKindIISLogs                       = DataSourceKind("IISLogs")
	DataSourceKindLinuxPerformanceObject        = DataSourceKind("LinuxPerformanceObject")
	DataSourceKindLinuxPerformanceCollection    = DataSourceKind("LinuxPerformanceCollection")
	DataSourceKindLinuxSyslog                   = DataSourceKind("LinuxSyslog")
	DataSourceKindLinuxSyslogCollection         = DataSourceKind("LinuxSyslogCollection")
	DataSourceKindWindowsEvent                  = DataSourceKind("WindowsEvent")
	DataSourceKindWindowsPerformanceCounter     = DataSourceKind("WindowsPerformanceCounter")
)

func (DataSourceKind) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceKind)(nil)).Elem()
}

func (e DataSourceKind) ToDataSourceKindOutput() DataSourceKindOutput {
	return pulumi.ToOutput(e).(DataSourceKindOutput)
}

func (e DataSourceKind) ToDataSourceKindOutputWithContext(ctx context.Context) DataSourceKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataSourceKindOutput)
}

func (e DataSourceKind) ToDataSourceKindPtrOutput() DataSourceKindPtrOutput {
	return e.ToDataSourceKindPtrOutputWithContext(context.Background())
}

func (e DataSourceKind) ToDataSourceKindPtrOutputWithContext(ctx context.Context) DataSourceKindPtrOutput {
	return DataSourceKind(e).ToDataSourceKindOutputWithContext(ctx).ToDataSourceKindPtrOutputWithContext(ctx)
}

func (e DataSourceKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSourceKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSourceKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataSourceKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataSourceKindOutput struct{ *pulumi.OutputState }

func (DataSourceKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceKind)(nil)).Elem()
}

func (o DataSourceKindOutput) ToDataSourceKindOutput() DataSourceKindOutput {
	return o
}

func (o DataSourceKindOutput) ToDataSourceKindOutputWithContext(ctx context.Context) DataSourceKindOutput {
	return o
}

func (o DataSourceKindOutput) ToDataSourceKindPtrOutput() DataSourceKindPtrOutput {
	return o.ToDataSourceKindPtrOutputWithContext(context.Background())
}

func (o DataSourceKindOutput) ToDataSourceKindPtrOutputWithContext(ctx context.Context) DataSourceKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataSourceKind) *DataSourceKind {
		return &v
	}).(DataSourceKindPtrOutput)
}

func (o DataSourceKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataSourceKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataSourceKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataSourceKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataSourceKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataSourceKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataSourceKindPtrOutput struct{ *pulumi.OutputState }

func (DataSourceKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSourceKind)(nil)).Elem()
}

func (o DataSourceKindPtrOutput) ToDataSourceKindPtrOutput() DataSourceKindPtrOutput {
	return o
}

func (o DataSourceKindPtrOutput) ToDataSourceKindPtrOutputWithContext(ctx context.Context) DataSourceKindPtrOutput {
	return o
}

func (o DataSourceKindPtrOutput) Elem() DataSourceKindOutput {
	return o.ApplyT(func(v *DataSourceKind) DataSourceKind {
		if v != nil {
			return *v
		}
		var ret DataSourceKind
		return ret
	}).(DataSourceKindOutput)
}

func (o DataSourceKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataSourceKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataSourceKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataSourceKindInput interface {
	pulumi.Input

	ToDataSourceKindOutput() DataSourceKindOutput
	ToDataSourceKindOutputWithContext(context.Context) DataSourceKindOutput
}

var dataSourceKindPtrType = reflect.TypeOf((**DataSourceKind)(nil)).Elem()

type DataSourceKindPtrInput interface {
	pulumi.Input

	ToDataSourceKindPtrOutput() DataSourceKindPtrOutput
	ToDataSourceKindPtrOutputWithContext(context.Context) DataSourceKindPtrOutput
}

type dataSourceKindPtr string

func DataSourceKindPtr(v string) DataSourceKindPtrInput {
	return (*dataSourceKindPtr)(&v)
}

func (*dataSourceKindPtr) ElementType() reflect.Type {
	return dataSourceKindPtrType
}

func (in *dataSourceKindPtr) ToDataSourceKindPtrOutput() DataSourceKindPtrOutput {
	return pulumi.ToOutput(in).(DataSourceKindPtrOutput)
}

func (in *dataSourceKindPtr) ToDataSourceKindPtrOutputWithContext(ctx context.Context) DataSourceKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataSourceKindPtrOutput)
}

type EntityStatus string

const (
	EntityStatusCreating            = EntityStatus("Creating")
	EntityStatusSucceeded           = EntityStatus("Succeeded")
	EntityStatusFailed              = EntityStatus("Failed")
	EntityStatusCanceled            = EntityStatus("Canceled")
	EntityStatusDeleting            = EntityStatus("Deleting")
	EntityStatusProvisioningAccount = EntityStatus("ProvisioningAccount")
)

func (EntityStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityStatus)(nil)).Elem()
}

func (e EntityStatus) ToEntityStatusOutput() EntityStatusOutput {
	return pulumi.ToOutput(e).(EntityStatusOutput)
}

func (e EntityStatus) ToEntityStatusOutputWithContext(ctx context.Context) EntityStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EntityStatusOutput)
}

func (e EntityStatus) ToEntityStatusPtrOutput() EntityStatusPtrOutput {
	return e.ToEntityStatusPtrOutputWithContext(context.Background())
}

func (e EntityStatus) ToEntityStatusPtrOutputWithContext(ctx context.Context) EntityStatusPtrOutput {
	return EntityStatus(e).ToEntityStatusOutputWithContext(ctx).ToEntityStatusPtrOutputWithContext(ctx)
}

func (e EntityStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EntityStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EntityStatusOutput struct{ *pulumi.OutputState }

func (EntityStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityStatus)(nil)).Elem()
}

func (o EntityStatusOutput) ToEntityStatusOutput() EntityStatusOutput {
	return o
}

func (o EntityStatusOutput) ToEntityStatusOutputWithContext(ctx context.Context) EntityStatusOutput {
	return o
}

func (o EntityStatusOutput) ToEntityStatusPtrOutput() EntityStatusPtrOutput {
	return o.ToEntityStatusPtrOutputWithContext(context.Background())
}

func (o EntityStatusOutput) ToEntityStatusPtrOutputWithContext(ctx context.Context) EntityStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntityStatus) *EntityStatus {
		return &v
	}).(EntityStatusPtrOutput)
}

func (o EntityStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EntityStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EntityStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EntityStatusPtrOutput struct{ *pulumi.OutputState }

func (EntityStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityStatus)(nil)).Elem()
}

func (o EntityStatusPtrOutput) ToEntityStatusPtrOutput() EntityStatusPtrOutput {
	return o
}

func (o EntityStatusPtrOutput) ToEntityStatusPtrOutputWithContext(ctx context.Context) EntityStatusPtrOutput {
	return o
}

func (o EntityStatusPtrOutput) Elem() EntityStatusOutput {
	return o.ApplyT(func(v *EntityStatus) EntityStatus {
		if v != nil {
			return *v
		}
		var ret EntityStatus
		return ret
	}).(EntityStatusOutput)
}

func (o EntityStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EntityStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EntityStatusInput interface {
	pulumi.Input

	ToEntityStatusOutput() EntityStatusOutput
	ToEntityStatusOutputWithContext(context.Context) EntityStatusOutput
}

var entityStatusPtrType = reflect.TypeOf((**EntityStatus)(nil)).Elem()

type EntityStatusPtrInput interface {
	pulumi.Input

	ToEntityStatusPtrOutput() EntityStatusPtrOutput
	ToEntityStatusPtrOutputWithContext(context.Context) EntityStatusPtrOutput
}

type entityStatusPtr string

func EntityStatusPtr(v string) EntityStatusPtrInput {
	return (*entityStatusPtr)(&v)
}

func (*entityStatusPtr) ElementType() reflect.Type {
	return entityStatusPtrType
}

func (in *entityStatusPtr) ToEntityStatusPtrOutput() EntityStatusPtrOutput {
	return pulumi.ToOutput(in).(EntityStatusPtrOutput)
}

func (in *entityStatusPtr) ToEntityStatusPtrOutputWithContext(ctx context.Context) EntityStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EntityStatusPtrOutput)
}

type MachineGroupType string

const (
	MachineGroupTypeUnknown      = MachineGroupType("unknown")
	MachineGroupType_Azure_cs    = MachineGroupType("azure-cs")
	MachineGroupType_Azure_sf    = MachineGroupType("azure-sf")
	MachineGroupType_Azure_vmss  = MachineGroupType("azure-vmss")
	MachineGroupType_User_static = MachineGroupType("user-static")
)

func (MachineGroupType) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineGroupType)(nil)).Elem()
}

func (e MachineGroupType) ToMachineGroupTypeOutput() MachineGroupTypeOutput {
	return pulumi.ToOutput(e).(MachineGroupTypeOutput)
}

func (e MachineGroupType) ToMachineGroupTypeOutputWithContext(ctx context.Context) MachineGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MachineGroupTypeOutput)
}

func (e MachineGroupType) ToMachineGroupTypePtrOutput() MachineGroupTypePtrOutput {
	return e.ToMachineGroupTypePtrOutputWithContext(context.Background())
}

func (e MachineGroupType) ToMachineGroupTypePtrOutputWithContext(ctx context.Context) MachineGroupTypePtrOutput {
	return MachineGroupType(e).ToMachineGroupTypeOutputWithContext(ctx).ToMachineGroupTypePtrOutputWithContext(ctx)
}

func (e MachineGroupType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MachineGroupType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MachineGroupType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MachineGroupType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MachineGroupTypeOutput struct{ *pulumi.OutputState }

func (MachineGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineGroupType)(nil)).Elem()
}

func (o MachineGroupTypeOutput) ToMachineGroupTypeOutput() MachineGroupTypeOutput {
	return o
}

func (o MachineGroupTypeOutput) ToMachineGroupTypeOutputWithContext(ctx context.Context) MachineGroupTypeOutput {
	return o
}

func (o MachineGroupTypeOutput) ToMachineGroupTypePtrOutput() MachineGroupTypePtrOutput {
	return o.ToMachineGroupTypePtrOutputWithContext(context.Background())
}

func (o MachineGroupTypeOutput) ToMachineGroupTypePtrOutputWithContext(ctx context.Context) MachineGroupTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineGroupType) *MachineGroupType {
		return &v
	}).(MachineGroupTypePtrOutput)
}

func (o MachineGroupTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MachineGroupTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MachineGroupType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MachineGroupTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MachineGroupTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MachineGroupType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MachineGroupTypePtrOutput struct{ *pulumi.OutputState }

func (MachineGroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineGroupType)(nil)).Elem()
}

func (o MachineGroupTypePtrOutput) ToMachineGroupTypePtrOutput() MachineGroupTypePtrOutput {
	return o
}

func (o MachineGroupTypePtrOutput) ToMachineGroupTypePtrOutputWithContext(ctx context.Context) MachineGroupTypePtrOutput {
	return o
}

func (o MachineGroupTypePtrOutput) Elem() MachineGroupTypeOutput {
	return o.ApplyT(func(v *MachineGroupType) MachineGroupType {
		if v != nil {
			return *v
		}
		var ret MachineGroupType
		return ret
	}).(MachineGroupTypeOutput)
}

func (o MachineGroupTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MachineGroupTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MachineGroupType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MachineGroupTypeInput interface {
	pulumi.Input

	ToMachineGroupTypeOutput() MachineGroupTypeOutput
	ToMachineGroupTypeOutputWithContext(context.Context) MachineGroupTypeOutput
}

var machineGroupTypePtrType = reflect.TypeOf((**MachineGroupType)(nil)).Elem()

type MachineGroupTypePtrInput interface {
	pulumi.Input

	ToMachineGroupTypePtrOutput() MachineGroupTypePtrOutput
	ToMachineGroupTypePtrOutputWithContext(context.Context) MachineGroupTypePtrOutput
}

type machineGroupTypePtr string

func MachineGroupTypePtr(v string) MachineGroupTypePtrInput {
	return (*machineGroupTypePtr)(&v)
}

func (*machineGroupTypePtr) ElementType() reflect.Type {
	return machineGroupTypePtrType
}

func (in *machineGroupTypePtr) ToMachineGroupTypePtrOutput() MachineGroupTypePtrOutput {
	return pulumi.ToOutput(in).(MachineGroupTypePtrOutput)
}

func (in *machineGroupTypePtr) ToMachineGroupTypePtrOutputWithContext(ctx context.Context) MachineGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MachineGroupTypePtrOutput)
}

type SkuNameEnum string

const (
	SkuNameEnumFree                = SkuNameEnum("Free")
	SkuNameEnumStandard            = SkuNameEnum("Standard")
	SkuNameEnumPremium             = SkuNameEnum("Premium")
	SkuNameEnumPerNode             = SkuNameEnum("PerNode")
	SkuNameEnumPerGB2018           = SkuNameEnum("PerGB2018")
	SkuNameEnumStandalone          = SkuNameEnum("Standalone")
	SkuNameEnumCapacityReservation = SkuNameEnum("CapacityReservation")
)

func (SkuNameEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuNameEnum)(nil)).Elem()
}

func (e SkuNameEnum) ToSkuNameEnumOutput() SkuNameEnumOutput {
	return pulumi.ToOutput(e).(SkuNameEnumOutput)
}

func (e SkuNameEnum) ToSkuNameEnumOutputWithContext(ctx context.Context) SkuNameEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameEnumOutput)
}

func (e SkuNameEnum) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return e.ToSkuNameEnumPtrOutputWithContext(context.Background())
}

func (e SkuNameEnum) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return SkuNameEnum(e).ToSkuNameEnumOutputWithContext(ctx).ToSkuNameEnumPtrOutputWithContext(ctx)
}

func (e SkuNameEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuNameEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuNameEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuNameEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameEnumOutput struct{ *pulumi.OutputState }

func (SkuNameEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuNameEnum)(nil)).Elem()
}

func (o SkuNameEnumOutput) ToSkuNameEnumOutput() SkuNameEnumOutput {
	return o
}

func (o SkuNameEnumOutput) ToSkuNameEnumOutputWithContext(ctx context.Context) SkuNameEnumOutput {
	return o
}

func (o SkuNameEnumOutput) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return o.ToSkuNameEnumPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuNameEnum) *SkuNameEnum {
		return &v
	}).(SkuNameEnumPtrOutput)
}

func (o SkuNameEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuNameEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuNameEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNameEnumPtrOutput struct{ *pulumi.OutputState }

func (SkuNameEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuNameEnum)(nil)).Elem()
}

func (o SkuNameEnumPtrOutput) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return o
}

func (o SkuNameEnumPtrOutput) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return o
}

func (o SkuNameEnumPtrOutput) Elem() SkuNameEnumOutput {
	return o.ApplyT(func(v *SkuNameEnum) SkuNameEnum {
		if v != nil {
			return *v
		}
		var ret SkuNameEnum
		return ret
	}).(SkuNameEnumOutput)
}

func (o SkuNameEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuNameEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameEnumInput interface {
	pulumi.Input

	ToSkuNameEnumOutput() SkuNameEnumOutput
	ToSkuNameEnumOutputWithContext(context.Context) SkuNameEnumOutput
}

var skuNameEnumPtrType = reflect.TypeOf((**SkuNameEnum)(nil)).Elem()

type SkuNameEnumPtrInput interface {
	pulumi.Input

	ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput
	ToSkuNameEnumPtrOutputWithContext(context.Context) SkuNameEnumPtrOutput
}

type skuNameEnumPtr string

func SkuNameEnumPtr(v string) SkuNameEnumPtrInput {
	return (*skuNameEnumPtr)(&v)
}

func (*skuNameEnumPtr) ElementType() reflect.Type {
	return skuNameEnumPtrType
}

func (in *skuNameEnumPtr) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return pulumi.ToOutput(in).(SkuNameEnumPtrOutput)
}

func (in *skuNameEnumPtr) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNameEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSourceKindOutput{})
	pulumi.RegisterOutputType(DataSourceKindPtrOutput{})
	pulumi.RegisterOutputType(EntityStatusOutput{})
	pulumi.RegisterOutputType(EntityStatusPtrOutput{})
	pulumi.RegisterOutputType(MachineGroupTypeOutput{})
	pulumi.RegisterOutputType(MachineGroupTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameEnumOutput{})
	pulumi.RegisterOutputType(SkuNameEnumPtrOutput{})
}
