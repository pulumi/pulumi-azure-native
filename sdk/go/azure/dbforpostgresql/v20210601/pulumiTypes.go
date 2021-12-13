


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Backup struct {
	BackupRetentionDays *int    `pulumi:"backupRetentionDays"`
	GeoRedundantBackup  *string `pulumi:"geoRedundantBackup"`
}


func (val *Backup) Defaults() *Backup {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.BackupRetentionDays) {
		backupRetentionDays_ := 7
		tmp.BackupRetentionDays = &backupRetentionDays_
	}
	if isZero(tmp.GeoRedundantBackup) {
		geoRedundantBackup_ := "Disabled"
		tmp.GeoRedundantBackup = &geoRedundantBackup_
	}
	return &tmp
}





type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(context.Context) BackupOutput
}

type BackupArgs struct {
	BackupRetentionDays pulumi.IntPtrInput    `pulumi:"backupRetentionDays"`
	GeoRedundantBackup  pulumi.StringPtrInput `pulumi:"geoRedundantBackup"`
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Backup)(nil)).Elem()
}

func (i BackupArgs) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i BackupArgs) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

func (i BackupArgs) ToBackupPtrOutput() BackupPtrOutput {
	return i.ToBackupPtrOutputWithContext(context.Background())
}

func (i BackupArgs) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput).ToBackupPtrOutputWithContext(ctx)
}









type BackupPtrInput interface {
	pulumi.Input

	ToBackupPtrOutput() BackupPtrOutput
	ToBackupPtrOutputWithContext(context.Context) BackupPtrOutput
}

type backupPtrType BackupArgs

func BackupPtr(v *BackupArgs) BackupPtrInput {
	return (*backupPtrType)(v)
}

func (*backupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *backupPtrType) ToBackupPtrOutput() BackupPtrOutput {
	return i.ToBackupPtrOutputWithContext(context.Background())
}

func (i *backupPtrType) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPtrOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

func (o BackupOutput) ToBackupPtrOutput() BackupPtrOutput {
	return o.ToBackupPtrOutputWithContext(context.Background())
}

func (o BackupOutput) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Backup) *Backup {
		return &v
	}).(BackupPtrOutput)
}

func (o BackupOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backup) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o BackupOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backup) *string { return v.GeoRedundantBackup }).(pulumi.StringPtrOutput)
}

type BackupPtrOutput struct{ *pulumi.OutputState }

func (BackupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupPtrOutput) ToBackupPtrOutput() BackupPtrOutput {
	return o
}

func (o BackupPtrOutput) ToBackupPtrOutputWithContext(ctx context.Context) BackupPtrOutput {
	return o
}

func (o BackupPtrOutput) Elem() BackupOutput {
	return o.ApplyT(func(v *Backup) Backup {
		if v != nil {
			return *v
		}
		var ret Backup
		return ret
	}).(BackupOutput)
}

func (o BackupPtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Backup) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupPtrOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backup) *string {
		if v == nil {
			return nil
		}
		return v.GeoRedundantBackup
	}).(pulumi.StringPtrOutput)
}

type BackupResponse struct {
	BackupRetentionDays *int    `pulumi:"backupRetentionDays"`
	EarliestRestoreDate string  `pulumi:"earliestRestoreDate"`
	GeoRedundantBackup  *string `pulumi:"geoRedundantBackup"`
}


func (val *BackupResponse) Defaults() *BackupResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.BackupRetentionDays) {
		backupRetentionDays_ := 7
		tmp.BackupRetentionDays = &backupRetentionDays_
	}
	if isZero(tmp.GeoRedundantBackup) {
		geoRedundantBackup_ := "Disabled"
		tmp.GeoRedundantBackup = &geoRedundantBackup_
	}
	return &tmp
}

type BackupResponseOutput struct{ *pulumi.OutputState }

func (BackupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupResponse)(nil)).Elem()
}

func (o BackupResponseOutput) ToBackupResponseOutput() BackupResponseOutput {
	return o
}

func (o BackupResponseOutput) ToBackupResponseOutputWithContext(ctx context.Context) BackupResponseOutput {
	return o
}

func (o BackupResponseOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackupResponse) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o BackupResponseOutput) EarliestRestoreDate() pulumi.StringOutput {
	return o.ApplyT(func(v BackupResponse) string { return v.EarliestRestoreDate }).(pulumi.StringOutput)
}

func (o BackupResponseOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupResponse) *string { return v.GeoRedundantBackup }).(pulumi.StringPtrOutput)
}

type BackupResponsePtrOutput struct{ *pulumi.OutputState }

func (BackupResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupResponse)(nil)).Elem()
}

func (o BackupResponsePtrOutput) ToBackupResponsePtrOutput() BackupResponsePtrOutput {
	return o
}

func (o BackupResponsePtrOutput) ToBackupResponsePtrOutputWithContext(ctx context.Context) BackupResponsePtrOutput {
	return o
}

func (o BackupResponsePtrOutput) Elem() BackupResponseOutput {
	return o.ApplyT(func(v *BackupResponse) BackupResponse {
		if v != nil {
			return *v
		}
		var ret BackupResponse
		return ret
	}).(BackupResponseOutput)
}

func (o BackupResponsePtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupResponse) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupResponsePtrOutput) EarliestRestoreDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EarliestRestoreDate
	}).(pulumi.StringPtrOutput)
}

func (o BackupResponsePtrOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupResponse) *string {
		if v == nil {
			return nil
		}
		return v.GeoRedundantBackup
	}).(pulumi.StringPtrOutput)
}

type HighAvailability struct {
	Mode                    *string `pulumi:"mode"`
	StandbyAvailabilityZone *string `pulumi:"standbyAvailabilityZone"`
}


func (val *HighAvailability) Defaults() *HighAvailability {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "Disabled"
		tmp.Mode = &mode_
	}
	if isZero(tmp.StandbyAvailabilityZone) {
		standbyAvailabilityZone_ := ""
		tmp.StandbyAvailabilityZone = &standbyAvailabilityZone_
	}
	return &tmp
}





type HighAvailabilityInput interface {
	pulumi.Input

	ToHighAvailabilityOutput() HighAvailabilityOutput
	ToHighAvailabilityOutputWithContext(context.Context) HighAvailabilityOutput
}

type HighAvailabilityArgs struct {
	Mode                    pulumi.StringPtrInput `pulumi:"mode"`
	StandbyAvailabilityZone pulumi.StringPtrInput `pulumi:"standbyAvailabilityZone"`
}

func (HighAvailabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailability)(nil)).Elem()
}

func (i HighAvailabilityArgs) ToHighAvailabilityOutput() HighAvailabilityOutput {
	return i.ToHighAvailabilityOutputWithContext(context.Background())
}

func (i HighAvailabilityArgs) ToHighAvailabilityOutputWithContext(ctx context.Context) HighAvailabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityOutput)
}

func (i HighAvailabilityArgs) ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput {
	return i.ToHighAvailabilityPtrOutputWithContext(context.Background())
}

func (i HighAvailabilityArgs) ToHighAvailabilityPtrOutputWithContext(ctx context.Context) HighAvailabilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityOutput).ToHighAvailabilityPtrOutputWithContext(ctx)
}









type HighAvailabilityPtrInput interface {
	pulumi.Input

	ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput
	ToHighAvailabilityPtrOutputWithContext(context.Context) HighAvailabilityPtrOutput
}

type highAvailabilityPtrType HighAvailabilityArgs

func HighAvailabilityPtr(v *HighAvailabilityArgs) HighAvailabilityPtrInput {
	return (*highAvailabilityPtrType)(v)
}

func (*highAvailabilityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailability)(nil)).Elem()
}

func (i *highAvailabilityPtrType) ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput {
	return i.ToHighAvailabilityPtrOutputWithContext(context.Background())
}

func (i *highAvailabilityPtrType) ToHighAvailabilityPtrOutputWithContext(ctx context.Context) HighAvailabilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HighAvailabilityPtrOutput)
}

type HighAvailabilityOutput struct{ *pulumi.OutputState }

func (HighAvailabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailability)(nil)).Elem()
}

func (o HighAvailabilityOutput) ToHighAvailabilityOutput() HighAvailabilityOutput {
	return o
}

func (o HighAvailabilityOutput) ToHighAvailabilityOutputWithContext(ctx context.Context) HighAvailabilityOutput {
	return o
}

func (o HighAvailabilityOutput) ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput {
	return o.ToHighAvailabilityPtrOutputWithContext(context.Background())
}

func (o HighAvailabilityOutput) ToHighAvailabilityPtrOutputWithContext(ctx context.Context) HighAvailabilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HighAvailability) *HighAvailability {
		return &v
	}).(HighAvailabilityPtrOutput)
}

func (o HighAvailabilityOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HighAvailability) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HighAvailability) *string { return v.StandbyAvailabilityZone }).(pulumi.StringPtrOutput)
}

type HighAvailabilityPtrOutput struct{ *pulumi.OutputState }

func (HighAvailabilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailability)(nil)).Elem()
}

func (o HighAvailabilityPtrOutput) ToHighAvailabilityPtrOutput() HighAvailabilityPtrOutput {
	return o
}

func (o HighAvailabilityPtrOutput) ToHighAvailabilityPtrOutputWithContext(ctx context.Context) HighAvailabilityPtrOutput {
	return o
}

func (o HighAvailabilityPtrOutput) Elem() HighAvailabilityOutput {
	return o.ApplyT(func(v *HighAvailability) HighAvailability {
		if v != nil {
			return *v
		}
		var ret HighAvailability
		return ret
	}).(HighAvailabilityOutput)
}

func (o HighAvailabilityPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailability) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityPtrOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailability) *string {
		if v == nil {
			return nil
		}
		return v.StandbyAvailabilityZone
	}).(pulumi.StringPtrOutput)
}

type HighAvailabilityResponse struct {
	Mode                    *string `pulumi:"mode"`
	StandbyAvailabilityZone *string `pulumi:"standbyAvailabilityZone"`
	State                   string  `pulumi:"state"`
}


func (val *HighAvailabilityResponse) Defaults() *HighAvailabilityResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "Disabled"
		tmp.Mode = &mode_
	}
	if isZero(tmp.StandbyAvailabilityZone) {
		standbyAvailabilityZone_ := ""
		tmp.StandbyAvailabilityZone = &standbyAvailabilityZone_
	}
	return &tmp
}

type HighAvailabilityResponseOutput struct{ *pulumi.OutputState }

func (HighAvailabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailabilityResponse)(nil)).Elem()
}

func (o HighAvailabilityResponseOutput) ToHighAvailabilityResponseOutput() HighAvailabilityResponseOutput {
	return o
}

func (o HighAvailabilityResponseOutput) ToHighAvailabilityResponseOutputWithContext(ctx context.Context) HighAvailabilityResponseOutput {
	return o
}

func (o HighAvailabilityResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HighAvailabilityResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityResponseOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HighAvailabilityResponse) *string { return v.StandbyAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v HighAvailabilityResponse) string { return v.State }).(pulumi.StringOutput)
}

type HighAvailabilityResponsePtrOutput struct{ *pulumi.OutputState }

func (HighAvailabilityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailabilityResponse)(nil)).Elem()
}

func (o HighAvailabilityResponsePtrOutput) ToHighAvailabilityResponsePtrOutput() HighAvailabilityResponsePtrOutput {
	return o
}

func (o HighAvailabilityResponsePtrOutput) ToHighAvailabilityResponsePtrOutputWithContext(ctx context.Context) HighAvailabilityResponsePtrOutput {
	return o
}

func (o HighAvailabilityResponsePtrOutput) Elem() HighAvailabilityResponseOutput {
	return o.ApplyT(func(v *HighAvailabilityResponse) HighAvailabilityResponse {
		if v != nil {
			return *v
		}
		var ret HighAvailabilityResponse
		return ret
	}).(HighAvailabilityResponseOutput)
}

func (o HighAvailabilityResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailabilityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityResponsePtrOutput) StandbyAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailabilityResponse) *string {
		if v == nil {
			return nil
		}
		return v.StandbyAvailabilityZone
	}).(pulumi.StringPtrOutput)
}

func (o HighAvailabilityResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HighAvailabilityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type MaintenanceWindow struct {
	CustomWindow *string `pulumi:"customWindow"`
	DayOfWeek    *int    `pulumi:"dayOfWeek"`
	StartHour    *int    `pulumi:"startHour"`
	StartMinute  *int    `pulumi:"startMinute"`
}


func (val *MaintenanceWindow) Defaults() *MaintenanceWindow {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CustomWindow) {
		customWindow_ := "Disabled"
		tmp.CustomWindow = &customWindow_
	}
	if isZero(tmp.DayOfWeek) {
		dayOfWeek_ := 0
		tmp.DayOfWeek = &dayOfWeek_
	}
	if isZero(tmp.StartHour) {
		startHour_ := 0
		tmp.StartHour = &startHour_
	}
	if isZero(tmp.StartMinute) {
		startMinute_ := 0
		tmp.StartMinute = &startMinute_
	}
	return &tmp
}





type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(context.Context) MaintenanceWindowOutput
}

type MaintenanceWindowArgs struct {
	CustomWindow pulumi.StringPtrInput `pulumi:"customWindow"`
	DayOfWeek    pulumi.IntPtrInput    `pulumi:"dayOfWeek"`
	StartHour    pulumi.IntPtrInput    `pulumi:"startHour"`
	StartMinute  pulumi.IntPtrInput    `pulumi:"startMinute"`
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return i.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (i MaintenanceWindowArgs) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput).ToMaintenanceWindowPtrOutputWithContext(ctx)
}









type MaintenanceWindowPtrInput interface {
	pulumi.Input

	ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput
	ToMaintenanceWindowPtrOutputWithContext(context.Context) MaintenanceWindowPtrOutput
}

type maintenanceWindowPtrType MaintenanceWindowArgs

func MaintenanceWindowPtr(v *MaintenanceWindowArgs) MaintenanceWindowPtrInput {
	return (*maintenanceWindowPtrType)(v)
}

func (*maintenanceWindowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (i *maintenanceWindowPtrType) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return i.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (i *maintenanceWindowPtrType) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowPtrOutput)
}

type MaintenanceWindowOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return o.ToMaintenanceWindowPtrOutputWithContext(context.Background())
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MaintenanceWindow) *MaintenanceWindow {
		return &v
	}).(MaintenanceWindowPtrOutput)
}

func (o MaintenanceWindowOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *string { return v.CustomWindow }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.DayOfWeek }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.StartHour }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindow) *int { return v.StartMinute }).(pulumi.IntPtrOutput)
}

type MaintenanceWindowPtrOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowPtrOutput) ToMaintenanceWindowPtrOutput() MaintenanceWindowPtrOutput {
	return o
}

func (o MaintenanceWindowPtrOutput) ToMaintenanceWindowPtrOutputWithContext(ctx context.Context) MaintenanceWindowPtrOutput {
	return o
}

func (o MaintenanceWindowPtrOutput) Elem() MaintenanceWindowOutput {
	return o.ApplyT(func(v *MaintenanceWindow) MaintenanceWindow {
		if v != nil {
			return *v
		}
		var ret MaintenanceWindow
		return ret
	}).(MaintenanceWindowOutput)
}

func (o MaintenanceWindowPtrOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *string {
		if v == nil {
			return nil
		}
		return v.CustomWindow
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowPtrOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowPtrOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.StartHour
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowPtrOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) *int {
		if v == nil {
			return nil
		}
		return v.StartMinute
	}).(pulumi.IntPtrOutput)
}

type MaintenanceWindowResponse struct {
	CustomWindow *string `pulumi:"customWindow"`
	DayOfWeek    *int    `pulumi:"dayOfWeek"`
	StartHour    *int    `pulumi:"startHour"`
	StartMinute  *int    `pulumi:"startMinute"`
}


func (val *MaintenanceWindowResponse) Defaults() *MaintenanceWindowResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CustomWindow) {
		customWindow_ := "Disabled"
		tmp.CustomWindow = &customWindow_
	}
	if isZero(tmp.DayOfWeek) {
		dayOfWeek_ := 0
		tmp.DayOfWeek = &dayOfWeek_
	}
	if isZero(tmp.StartHour) {
		startHour_ := 0
		tmp.StartHour = &startHour_
	}
	if isZero(tmp.StartMinute) {
		startMinute_ := 0
		tmp.StartMinute = &startMinute_
	}
	return &tmp
}

type MaintenanceWindowResponseOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowResponse)(nil)).Elem()
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponseOutput() MaintenanceWindowResponseOutput {
	return o
}

func (o MaintenanceWindowResponseOutput) ToMaintenanceWindowResponseOutputWithContext(ctx context.Context) MaintenanceWindowResponseOutput {
	return o
}

func (o MaintenanceWindowResponseOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *string { return v.CustomWindow }).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowResponseOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.DayOfWeek }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponseOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.StartHour }).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponseOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MaintenanceWindowResponse) *int { return v.StartMinute }).(pulumi.IntPtrOutput)
}

type MaintenanceWindowResponsePtrOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowResponse)(nil)).Elem()
}

func (o MaintenanceWindowResponsePtrOutput) ToMaintenanceWindowResponsePtrOutput() MaintenanceWindowResponsePtrOutput {
	return o
}

func (o MaintenanceWindowResponsePtrOutput) ToMaintenanceWindowResponsePtrOutputWithContext(ctx context.Context) MaintenanceWindowResponsePtrOutput {
	return o
}

func (o MaintenanceWindowResponsePtrOutput) Elem() MaintenanceWindowResponseOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) MaintenanceWindowResponse {
		if v != nil {
			return *v
		}
		var ret MaintenanceWindowResponse
		return ret
	}).(MaintenanceWindowResponseOutput)
}

func (o MaintenanceWindowResponsePtrOutput) CustomWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomWindow
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) DayOfWeek() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) StartHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.StartHour
	}).(pulumi.IntPtrOutput)
}

func (o MaintenanceWindowResponsePtrOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindowResponse) *int {
		if v == nil {
			return nil
		}
		return v.StartMinute
	}).(pulumi.IntPtrOutput)
}

type Network struct {
	DelegatedSubnetResourceId   *string `pulumi:"delegatedSubnetResourceId"`
	PrivateDnsZoneArmResourceId *string `pulumi:"privateDnsZoneArmResourceId"`
}


func (val *Network) Defaults() *Network {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DelegatedSubnetResourceId) {
		delegatedSubnetResourceId_ := ""
		tmp.DelegatedSubnetResourceId = &delegatedSubnetResourceId_
	}
	if isZero(tmp.PrivateDnsZoneArmResourceId) {
		privateDnsZoneArmResourceId_ := ""
		tmp.PrivateDnsZoneArmResourceId = &privateDnsZoneArmResourceId_
	}
	return &tmp
}





type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(context.Context) NetworkOutput
}

type NetworkArgs struct {
	DelegatedSubnetResourceId   pulumi.StringPtrInput `pulumi:"delegatedSubnetResourceId"`
	PrivateDnsZoneArmResourceId pulumi.StringPtrInput `pulumi:"privateDnsZoneArmResourceId"`
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil)).Elem()
}

func (i NetworkArgs) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i NetworkArgs) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

func (i NetworkArgs) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i NetworkArgs) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput).ToNetworkPtrOutputWithContext(ctx)
}









type NetworkPtrInput interface {
	pulumi.Input

	ToNetworkPtrOutput() NetworkPtrOutput
	ToNetworkPtrOutputWithContext(context.Context) NetworkPtrOutput
}

type networkPtrType NetworkArgs

func NetworkPtr(v *NetworkArgs) NetworkPtrInput {
	return (*networkPtrType)(v)
}

func (*networkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (i *networkPtrType) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i *networkPtrType) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPtrOutput)
}

type NetworkOutput struct{ *pulumi.OutputState }

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil)).Elem()
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o.ToNetworkPtrOutputWithContext(context.Background())
}

func (o NetworkOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Network) *Network {
		return &v
	}).(NetworkPtrOutput)
}

func (o NetworkOutput) DelegatedSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Network) *string { return v.DelegatedSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o NetworkOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Network) *string { return v.PrivateDnsZoneArmResourceId }).(pulumi.StringPtrOutput)
}

type NetworkPtrOutput struct{ *pulumi.OutputState }

func (NetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil)).Elem()
}

func (o NetworkPtrOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o
}

func (o NetworkPtrOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o
}

func (o NetworkPtrOutput) Elem() NetworkOutput {
	return o.ApplyT(func(v *Network) Network {
		if v != nil {
			return *v
		}
		var ret Network
		return ret
	}).(NetworkOutput)
}

func (o NetworkPtrOutput) DelegatedSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) *string {
		if v == nil {
			return nil
		}
		return v.DelegatedSubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkPtrOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Network) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

type NetworkResponse struct {
	DelegatedSubnetResourceId   *string `pulumi:"delegatedSubnetResourceId"`
	PrivateDnsZoneArmResourceId *string `pulumi:"privateDnsZoneArmResourceId"`
	PublicNetworkAccess         string  `pulumi:"publicNetworkAccess"`
}


func (val *NetworkResponse) Defaults() *NetworkResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DelegatedSubnetResourceId) {
		delegatedSubnetResourceId_ := ""
		tmp.DelegatedSubnetResourceId = &delegatedSubnetResourceId_
	}
	if isZero(tmp.PrivateDnsZoneArmResourceId) {
		privateDnsZoneArmResourceId_ := ""
		tmp.PrivateDnsZoneArmResourceId = &privateDnsZoneArmResourceId_
	}
	return &tmp
}

type NetworkResponseOutput struct{ *pulumi.OutputState }

func (NetworkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkResponse)(nil)).Elem()
}

func (o NetworkResponseOutput) ToNetworkResponseOutput() NetworkResponseOutput {
	return o
}

func (o NetworkResponseOutput) ToNetworkResponseOutputWithContext(ctx context.Context) NetworkResponseOutput {
	return o
}

func (o NetworkResponseOutput) DelegatedSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkResponse) *string { return v.DelegatedSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o NetworkResponseOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkResponse) *string { return v.PrivateDnsZoneArmResourceId }).(pulumi.StringPtrOutput)
}

func (o NetworkResponseOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkResponse) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

type NetworkResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkResponse)(nil)).Elem()
}

func (o NetworkResponsePtrOutput) ToNetworkResponsePtrOutput() NetworkResponsePtrOutput {
	return o
}

func (o NetworkResponsePtrOutput) ToNetworkResponsePtrOutputWithContext(ctx context.Context) NetworkResponsePtrOutput {
	return o
}

func (o NetworkResponsePtrOutput) Elem() NetworkResponseOutput {
	return o.ApplyT(func(v *NetworkResponse) NetworkResponse {
		if v != nil {
			return *v
		}
		var ret NetworkResponse
		return ret
	}).(NetworkResponseOutput)
}

func (o NetworkResponsePtrOutput) DelegatedSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkResponse) *string {
		if v == nil {
			return nil
		}
		return v.DelegatedSubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkResponsePtrOutput) PrivateDnsZoneArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateDnsZoneArmResourceId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type Storage struct {
	StorageSizeGB *int `pulumi:"storageSizeGB"`
}





type StorageInput interface {
	pulumi.Input

	ToStorageOutput() StorageOutput
	ToStorageOutputWithContext(context.Context) StorageOutput
}

type StorageArgs struct {
	StorageSizeGB pulumi.IntPtrInput `pulumi:"storageSizeGB"`
}

func (StorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Storage)(nil)).Elem()
}

func (i StorageArgs) ToStorageOutput() StorageOutput {
	return i.ToStorageOutputWithContext(context.Background())
}

func (i StorageArgs) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageOutput)
}

func (i StorageArgs) ToStoragePtrOutput() StoragePtrOutput {
	return i.ToStoragePtrOutputWithContext(context.Background())
}

func (i StorageArgs) ToStoragePtrOutputWithContext(ctx context.Context) StoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageOutput).ToStoragePtrOutputWithContext(ctx)
}









type StoragePtrInput interface {
	pulumi.Input

	ToStoragePtrOutput() StoragePtrOutput
	ToStoragePtrOutputWithContext(context.Context) StoragePtrOutput
}

type storagePtrType StorageArgs

func StoragePtr(v *StorageArgs) StoragePtrInput {
	return (*storagePtrType)(v)
}

func (*storagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (i *storagePtrType) ToStoragePtrOutput() StoragePtrOutput {
	return i.ToStoragePtrOutputWithContext(context.Background())
}

func (i *storagePtrType) ToStoragePtrOutputWithContext(ctx context.Context) StoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragePtrOutput)
}

type StorageOutput struct{ *pulumi.OutputState }

func (StorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Storage)(nil)).Elem()
}

func (o StorageOutput) ToStorageOutput() StorageOutput {
	return o
}

func (o StorageOutput) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return o
}

func (o StorageOutput) ToStoragePtrOutput() StoragePtrOutput {
	return o.ToStoragePtrOutputWithContext(context.Background())
}

func (o StorageOutput) ToStoragePtrOutputWithContext(ctx context.Context) StoragePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Storage) *Storage {
		return &v
	}).(StoragePtrOutput)
}

func (o StorageOutput) StorageSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Storage) *int { return v.StorageSizeGB }).(pulumi.IntPtrOutput)
}

type StoragePtrOutput struct{ *pulumi.OutputState }

func (StoragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (o StoragePtrOutput) ToStoragePtrOutput() StoragePtrOutput {
	return o
}

func (o StoragePtrOutput) ToStoragePtrOutputWithContext(ctx context.Context) StoragePtrOutput {
	return o
}

func (o StoragePtrOutput) Elem() StorageOutput {
	return o.ApplyT(func(v *Storage) Storage {
		if v != nil {
			return *v
		}
		var ret Storage
		return ret
	}).(StorageOutput)
}

func (o StoragePtrOutput) StorageSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Storage) *int {
		if v == nil {
			return nil
		}
		return v.StorageSizeGB
	}).(pulumi.IntPtrOutput)
}

type StorageResponse struct {
	StorageSizeGB *int `pulumi:"storageSizeGB"`
}

type StorageResponseOutput struct{ *pulumi.OutputState }

func (StorageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageResponse)(nil)).Elem()
}

func (o StorageResponseOutput) ToStorageResponseOutput() StorageResponseOutput {
	return o
}

func (o StorageResponseOutput) ToStorageResponseOutputWithContext(ctx context.Context) StorageResponseOutput {
	return o
}

func (o StorageResponseOutput) StorageSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageResponse) *int { return v.StorageSizeGB }).(pulumi.IntPtrOutput)
}

type StorageResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageResponse)(nil)).Elem()
}

func (o StorageResponsePtrOutput) ToStorageResponsePtrOutput() StorageResponsePtrOutput {
	return o
}

func (o StorageResponsePtrOutput) ToStorageResponsePtrOutputWithContext(ctx context.Context) StorageResponsePtrOutput {
	return o
}

func (o StorageResponsePtrOutput) Elem() StorageResponseOutput {
	return o.ApplyT(func(v *StorageResponse) StorageResponse {
		if v != nil {
			return *v
		}
		var ret StorageResponse
		return ret
	}).(StorageResponseOutput)
}

func (o StorageResponsePtrOutput) StorageSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageResponse) *int {
		if v == nil {
			return nil
		}
		return v.StorageSizeGB
	}).(pulumi.IntPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupOutput{})
	pulumi.RegisterOutputType(BackupPtrOutput{})
	pulumi.RegisterOutputType(BackupResponseOutput{})
	pulumi.RegisterOutputType(BackupResponsePtrOutput{})
	pulumi.RegisterOutputType(HighAvailabilityOutput{})
	pulumi.RegisterOutputType(HighAvailabilityPtrOutput{})
	pulumi.RegisterOutputType(HighAvailabilityResponseOutput{})
	pulumi.RegisterOutputType(HighAvailabilityResponsePtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowPtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkOutput{})
	pulumi.RegisterOutputType(NetworkPtrOutput{})
	pulumi.RegisterOutputType(NetworkResponseOutput{})
	pulumi.RegisterOutputType(NetworkResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageOutput{})
	pulumi.RegisterOutputType(StoragePtrOutput{})
	pulumi.RegisterOutputType(StorageResponseOutput{})
	pulumi.RegisterOutputType(StorageResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
