


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupScheduleType string

const (
	BackupScheduleTypeManual    = BackupScheduleType("Manual")
	BackupScheduleTypeAutomated = BackupScheduleType("Automated")
)

type ConnectivityType string

const (
	ConnectivityTypeLOCAL   = ConnectivityType("LOCAL")
	ConnectivityTypePRIVATE = ConnectivityType("PRIVATE")
	ConnectivityTypePUBLIC  = ConnectivityType("PUBLIC")
)

type DayOfWeek string

const (
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
	DayOfWeekSunday    = DayOfWeek("Sunday")
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

type DiskConfigurationType string

const (
	DiskConfigurationTypeNEW    = DiskConfigurationType("NEW")
	DiskConfigurationTypeEXTEND = DiskConfigurationType("EXTEND")
	DiskConfigurationTypeADD    = DiskConfigurationType("ADD")
)

type FullBackupFrequencyType string

const (
	FullBackupFrequencyTypeDaily  = FullBackupFrequencyType("Daily")
	FullBackupFrequencyTypeWeekly = FullBackupFrequencyType("Weekly")
)

type IdentityType string

const (
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
)

type SqlImageSku string

const (
	SqlImageSkuDeveloper  = SqlImageSku("Developer")
	SqlImageSkuExpress    = SqlImageSku("Express")
	SqlImageSkuStandard   = SqlImageSku("Standard")
	SqlImageSkuEnterprise = SqlImageSku("Enterprise")
	SqlImageSkuWeb        = SqlImageSku("Web")
)

type SqlManagementMode string

const (
	SqlManagementModeFull        = SqlManagementMode("Full")
	SqlManagementModeLightWeight = SqlManagementMode("LightWeight")
	SqlManagementModeNoAgent     = SqlManagementMode("NoAgent")
)

type SqlServerLicenseType string

const (
	SqlServerLicenseTypePAYG = SqlServerLicenseType("PAYG")
	SqlServerLicenseTypeAHUB = SqlServerLicenseType("AHUB")
	SqlServerLicenseTypeDR   = SqlServerLicenseType("DR")
)

type SqlVmGroupImageSku string

const (
	SqlVmGroupImageSkuDeveloper  = SqlVmGroupImageSku("Developer")
	SqlVmGroupImageSkuEnterprise = SqlVmGroupImageSku("Enterprise")
)

type SqlWorkloadType string

const (
	SqlWorkloadTypeGENERAL = SqlWorkloadType("GENERAL")
	SqlWorkloadTypeOLTP    = SqlWorkloadType("OLTP")
	SqlWorkloadTypeDW      = SqlWorkloadType("DW")
)

type StorageWorkloadType string

const (
	StorageWorkloadTypeGENERAL = StorageWorkloadType("GENERAL")
	StorageWorkloadTypeOLTP    = StorageWorkloadType("OLTP")
	StorageWorkloadTypeDW      = StorageWorkloadType("DW")
)

func init() {
	pulumi.RegisterOutputType(DayOfWeekOutput{})
	pulumi.RegisterOutputType(DayOfWeekPtrOutput{})
}
