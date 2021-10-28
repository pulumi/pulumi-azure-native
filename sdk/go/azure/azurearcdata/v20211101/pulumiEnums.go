


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArcSqlManagedInstanceLicenseType string

const (
	ArcSqlManagedInstanceLicenseTypeBasePrice       = ArcSqlManagedInstanceLicenseType("BasePrice")
	ArcSqlManagedInstanceLicenseTypeLicenseIncluded = ArcSqlManagedInstanceLicenseType("LicenseIncluded")
)

func (ArcSqlManagedInstanceLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcSqlManagedInstanceLicenseType)(nil)).Elem()
}

func (e ArcSqlManagedInstanceLicenseType) ToArcSqlManagedInstanceLicenseTypeOutput() ArcSqlManagedInstanceLicenseTypeOutput {
	return pulumi.ToOutput(e).(ArcSqlManagedInstanceLicenseTypeOutput)
}

func (e ArcSqlManagedInstanceLicenseType) ToArcSqlManagedInstanceLicenseTypeOutputWithContext(ctx context.Context) ArcSqlManagedInstanceLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ArcSqlManagedInstanceLicenseTypeOutput)
}

func (e ArcSqlManagedInstanceLicenseType) ToArcSqlManagedInstanceLicenseTypePtrOutput() ArcSqlManagedInstanceLicenseTypePtrOutput {
	return e.ToArcSqlManagedInstanceLicenseTypePtrOutputWithContext(context.Background())
}

func (e ArcSqlManagedInstanceLicenseType) ToArcSqlManagedInstanceLicenseTypePtrOutputWithContext(ctx context.Context) ArcSqlManagedInstanceLicenseTypePtrOutput {
	return ArcSqlManagedInstanceLicenseType(e).ToArcSqlManagedInstanceLicenseTypeOutputWithContext(ctx).ToArcSqlManagedInstanceLicenseTypePtrOutputWithContext(ctx)
}

func (e ArcSqlManagedInstanceLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ArcSqlManagedInstanceLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ArcSqlManagedInstanceLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ArcSqlManagedInstanceLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ArcSqlManagedInstanceLicenseTypeOutput struct{ *pulumi.OutputState }

func (ArcSqlManagedInstanceLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcSqlManagedInstanceLicenseType)(nil)).Elem()
}

func (o ArcSqlManagedInstanceLicenseTypeOutput) ToArcSqlManagedInstanceLicenseTypeOutput() ArcSqlManagedInstanceLicenseTypeOutput {
	return o
}

func (o ArcSqlManagedInstanceLicenseTypeOutput) ToArcSqlManagedInstanceLicenseTypeOutputWithContext(ctx context.Context) ArcSqlManagedInstanceLicenseTypeOutput {
	return o
}

func (o ArcSqlManagedInstanceLicenseTypeOutput) ToArcSqlManagedInstanceLicenseTypePtrOutput() ArcSqlManagedInstanceLicenseTypePtrOutput {
	return o.ToArcSqlManagedInstanceLicenseTypePtrOutputWithContext(context.Background())
}

func (o ArcSqlManagedInstanceLicenseTypeOutput) ToArcSqlManagedInstanceLicenseTypePtrOutputWithContext(ctx context.Context) ArcSqlManagedInstanceLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArcSqlManagedInstanceLicenseType) *ArcSqlManagedInstanceLicenseType {
		return &v
	}).(ArcSqlManagedInstanceLicenseTypePtrOutput)
}

func (o ArcSqlManagedInstanceLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ArcSqlManagedInstanceLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ArcSqlManagedInstanceLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ArcSqlManagedInstanceLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ArcSqlManagedInstanceLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ArcSqlManagedInstanceLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ArcSqlManagedInstanceLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (ArcSqlManagedInstanceLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcSqlManagedInstanceLicenseType)(nil)).Elem()
}

func (o ArcSqlManagedInstanceLicenseTypePtrOutput) ToArcSqlManagedInstanceLicenseTypePtrOutput() ArcSqlManagedInstanceLicenseTypePtrOutput {
	return o
}

func (o ArcSqlManagedInstanceLicenseTypePtrOutput) ToArcSqlManagedInstanceLicenseTypePtrOutputWithContext(ctx context.Context) ArcSqlManagedInstanceLicenseTypePtrOutput {
	return o
}

func (o ArcSqlManagedInstanceLicenseTypePtrOutput) Elem() ArcSqlManagedInstanceLicenseTypeOutput {
	return o.ApplyT(func(v *ArcSqlManagedInstanceLicenseType) ArcSqlManagedInstanceLicenseType {
		if v != nil {
			return *v
		}
		var ret ArcSqlManagedInstanceLicenseType
		return ret
	}).(ArcSqlManagedInstanceLicenseTypeOutput)
}

func (o ArcSqlManagedInstanceLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ArcSqlManagedInstanceLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ArcSqlManagedInstanceLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ArcSqlManagedInstanceLicenseTypeInput interface {
	pulumi.Input

	ToArcSqlManagedInstanceLicenseTypeOutput() ArcSqlManagedInstanceLicenseTypeOutput
	ToArcSqlManagedInstanceLicenseTypeOutputWithContext(context.Context) ArcSqlManagedInstanceLicenseTypeOutput
}

var arcSqlManagedInstanceLicenseTypePtrType = reflect.TypeOf((**ArcSqlManagedInstanceLicenseType)(nil)).Elem()

type ArcSqlManagedInstanceLicenseTypePtrInput interface {
	pulumi.Input

	ToArcSqlManagedInstanceLicenseTypePtrOutput() ArcSqlManagedInstanceLicenseTypePtrOutput
	ToArcSqlManagedInstanceLicenseTypePtrOutputWithContext(context.Context) ArcSqlManagedInstanceLicenseTypePtrOutput
}

type arcSqlManagedInstanceLicenseTypePtr string

func ArcSqlManagedInstanceLicenseTypePtr(v string) ArcSqlManagedInstanceLicenseTypePtrInput {
	return (*arcSqlManagedInstanceLicenseTypePtr)(&v)
}

func (*arcSqlManagedInstanceLicenseTypePtr) ElementType() reflect.Type {
	return arcSqlManagedInstanceLicenseTypePtrType
}

func (in *arcSqlManagedInstanceLicenseTypePtr) ToArcSqlManagedInstanceLicenseTypePtrOutput() ArcSqlManagedInstanceLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(ArcSqlManagedInstanceLicenseTypePtrOutput)
}

func (in *arcSqlManagedInstanceLicenseTypePtr) ToArcSqlManagedInstanceLicenseTypePtrOutputWithContext(ctx context.Context) ArcSqlManagedInstanceLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ArcSqlManagedInstanceLicenseTypePtrOutput)
}

type ArcSqlServerLicenseType string

const (
	ArcSqlServerLicenseTypePaid      = ArcSqlServerLicenseType("Paid")
	ArcSqlServerLicenseTypeFree      = ArcSqlServerLicenseType("Free")
	ArcSqlServerLicenseTypeHADR      = ArcSqlServerLicenseType("HADR")
	ArcSqlServerLicenseTypeUndefined = ArcSqlServerLicenseType("Undefined")
)

func (ArcSqlServerLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcSqlServerLicenseType)(nil)).Elem()
}

func (e ArcSqlServerLicenseType) ToArcSqlServerLicenseTypeOutput() ArcSqlServerLicenseTypeOutput {
	return pulumi.ToOutput(e).(ArcSqlServerLicenseTypeOutput)
}

func (e ArcSqlServerLicenseType) ToArcSqlServerLicenseTypeOutputWithContext(ctx context.Context) ArcSqlServerLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ArcSqlServerLicenseTypeOutput)
}

func (e ArcSqlServerLicenseType) ToArcSqlServerLicenseTypePtrOutput() ArcSqlServerLicenseTypePtrOutput {
	return e.ToArcSqlServerLicenseTypePtrOutputWithContext(context.Background())
}

func (e ArcSqlServerLicenseType) ToArcSqlServerLicenseTypePtrOutputWithContext(ctx context.Context) ArcSqlServerLicenseTypePtrOutput {
	return ArcSqlServerLicenseType(e).ToArcSqlServerLicenseTypeOutputWithContext(ctx).ToArcSqlServerLicenseTypePtrOutputWithContext(ctx)
}

func (e ArcSqlServerLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ArcSqlServerLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ArcSqlServerLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ArcSqlServerLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ArcSqlServerLicenseTypeOutput struct{ *pulumi.OutputState }

func (ArcSqlServerLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcSqlServerLicenseType)(nil)).Elem()
}

func (o ArcSqlServerLicenseTypeOutput) ToArcSqlServerLicenseTypeOutput() ArcSqlServerLicenseTypeOutput {
	return o
}

func (o ArcSqlServerLicenseTypeOutput) ToArcSqlServerLicenseTypeOutputWithContext(ctx context.Context) ArcSqlServerLicenseTypeOutput {
	return o
}

func (o ArcSqlServerLicenseTypeOutput) ToArcSqlServerLicenseTypePtrOutput() ArcSqlServerLicenseTypePtrOutput {
	return o.ToArcSqlServerLicenseTypePtrOutputWithContext(context.Background())
}

func (o ArcSqlServerLicenseTypeOutput) ToArcSqlServerLicenseTypePtrOutputWithContext(ctx context.Context) ArcSqlServerLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArcSqlServerLicenseType) *ArcSqlServerLicenseType {
		return &v
	}).(ArcSqlServerLicenseTypePtrOutput)
}

func (o ArcSqlServerLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ArcSqlServerLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ArcSqlServerLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ArcSqlServerLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ArcSqlServerLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ArcSqlServerLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ArcSqlServerLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (ArcSqlServerLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcSqlServerLicenseType)(nil)).Elem()
}

func (o ArcSqlServerLicenseTypePtrOutput) ToArcSqlServerLicenseTypePtrOutput() ArcSqlServerLicenseTypePtrOutput {
	return o
}

func (o ArcSqlServerLicenseTypePtrOutput) ToArcSqlServerLicenseTypePtrOutputWithContext(ctx context.Context) ArcSqlServerLicenseTypePtrOutput {
	return o
}

func (o ArcSqlServerLicenseTypePtrOutput) Elem() ArcSqlServerLicenseTypeOutput {
	return o.ApplyT(func(v *ArcSqlServerLicenseType) ArcSqlServerLicenseType {
		if v != nil {
			return *v
		}
		var ret ArcSqlServerLicenseType
		return ret
	}).(ArcSqlServerLicenseTypeOutput)
}

func (o ArcSqlServerLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ArcSqlServerLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ArcSqlServerLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ArcSqlServerLicenseTypeInput interface {
	pulumi.Input

	ToArcSqlServerLicenseTypeOutput() ArcSqlServerLicenseTypeOutput
	ToArcSqlServerLicenseTypeOutputWithContext(context.Context) ArcSqlServerLicenseTypeOutput
}

var arcSqlServerLicenseTypePtrType = reflect.TypeOf((**ArcSqlServerLicenseType)(nil)).Elem()

type ArcSqlServerLicenseTypePtrInput interface {
	pulumi.Input

	ToArcSqlServerLicenseTypePtrOutput() ArcSqlServerLicenseTypePtrOutput
	ToArcSqlServerLicenseTypePtrOutputWithContext(context.Context) ArcSqlServerLicenseTypePtrOutput
}

type arcSqlServerLicenseTypePtr string

func ArcSqlServerLicenseTypePtr(v string) ArcSqlServerLicenseTypePtrInput {
	return (*arcSqlServerLicenseTypePtr)(&v)
}

func (*arcSqlServerLicenseTypePtr) ElementType() reflect.Type {
	return arcSqlServerLicenseTypePtrType
}

func (in *arcSqlServerLicenseTypePtr) ToArcSqlServerLicenseTypePtrOutput() ArcSqlServerLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(ArcSqlServerLicenseTypePtrOutput)
}

func (in *arcSqlServerLicenseTypePtr) ToArcSqlServerLicenseTypePtrOutputWithContext(ctx context.Context) ArcSqlServerLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ArcSqlServerLicenseTypePtrOutput)
}

type ConnectionStatus string

const (
	ConnectionStatusConnected    = ConnectionStatus("Connected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
	ConnectionStatusUnknown      = ConnectionStatus("Unknown")
)

func (ConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatus)(nil)).Elem()
}

func (e ConnectionStatus) ToConnectionStatusOutput() ConnectionStatusOutput {
	return pulumi.ToOutput(e).(ConnectionStatusOutput)
}

func (e ConnectionStatus) ToConnectionStatusOutputWithContext(ctx context.Context) ConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionStatusOutput)
}

func (e ConnectionStatus) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return e.ToConnectionStatusPtrOutputWithContext(context.Background())
}

func (e ConnectionStatus) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return ConnectionStatus(e).ToConnectionStatusOutputWithContext(ctx).ToConnectionStatusPtrOutputWithContext(ctx)
}

func (e ConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionStatusOutput struct{ *pulumi.OutputState }

func (ConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatus)(nil)).Elem()
}

func (o ConnectionStatusOutput) ToConnectionStatusOutput() ConnectionStatusOutput {
	return o
}

func (o ConnectionStatusOutput) ToConnectionStatusOutputWithContext(ctx context.Context) ConnectionStatusOutput {
	return o
}

func (o ConnectionStatusOutput) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return o.ToConnectionStatusPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionStatus) *ConnectionStatus {
		return &v
	}).(ConnectionStatusPtrOutput)
}

func (o ConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStatus)(nil)).Elem()
}

func (o ConnectionStatusPtrOutput) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return o
}

func (o ConnectionStatusPtrOutput) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return o
}

func (o ConnectionStatusPtrOutput) Elem() ConnectionStatusOutput {
	return o.ApplyT(func(v *ConnectionStatus) ConnectionStatus {
		if v != nil {
			return *v
		}
		var ret ConnectionStatus
		return ret
	}).(ConnectionStatusOutput)
}

func (o ConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionStatusInput interface {
	pulumi.Input

	ToConnectionStatusOutput() ConnectionStatusOutput
	ToConnectionStatusOutputWithContext(context.Context) ConnectionStatusOutput
}

var connectionStatusPtrType = reflect.TypeOf((**ConnectionStatus)(nil)).Elem()

type ConnectionStatusPtrInput interface {
	pulumi.Input

	ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput
	ToConnectionStatusPtrOutputWithContext(context.Context) ConnectionStatusPtrOutput
}

type connectionStatusPtr string

func ConnectionStatusPtr(v string) ConnectionStatusPtrInput {
	return (*connectionStatusPtr)(&v)
}

func (*connectionStatusPtr) ElementType() reflect.Type {
	return connectionStatusPtrType
}

func (in *connectionStatusPtr) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(ConnectionStatusPtrOutput)
}

func (in *connectionStatusPtr) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionStatusPtrOutput)
}

type DefenderStatus string

const (
	DefenderStatusProtected   = DefenderStatus("Protected")
	DefenderStatusUnprotected = DefenderStatus("Unprotected")
	DefenderStatusUnknown     = DefenderStatus("Unknown")
)

func (DefenderStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderStatus)(nil)).Elem()
}

func (e DefenderStatus) ToDefenderStatusOutput() DefenderStatusOutput {
	return pulumi.ToOutput(e).(DefenderStatusOutput)
}

func (e DefenderStatus) ToDefenderStatusOutputWithContext(ctx context.Context) DefenderStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefenderStatusOutput)
}

func (e DefenderStatus) ToDefenderStatusPtrOutput() DefenderStatusPtrOutput {
	return e.ToDefenderStatusPtrOutputWithContext(context.Background())
}

func (e DefenderStatus) ToDefenderStatusPtrOutputWithContext(ctx context.Context) DefenderStatusPtrOutput {
	return DefenderStatus(e).ToDefenderStatusOutputWithContext(ctx).ToDefenderStatusPtrOutputWithContext(ctx)
}

func (e DefenderStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefenderStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefenderStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefenderStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefenderStatusOutput struct{ *pulumi.OutputState }

func (DefenderStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderStatus)(nil)).Elem()
}

func (o DefenderStatusOutput) ToDefenderStatusOutput() DefenderStatusOutput {
	return o
}

func (o DefenderStatusOutput) ToDefenderStatusOutputWithContext(ctx context.Context) DefenderStatusOutput {
	return o
}

func (o DefenderStatusOutput) ToDefenderStatusPtrOutput() DefenderStatusPtrOutput {
	return o.ToDefenderStatusPtrOutputWithContext(context.Background())
}

func (o DefenderStatusOutput) ToDefenderStatusPtrOutputWithContext(ctx context.Context) DefenderStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderStatus) *DefenderStatus {
		return &v
	}).(DefenderStatusPtrOutput)
}

func (o DefenderStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefenderStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefenderStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefenderStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefenderStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefenderStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefenderStatusPtrOutput struct{ *pulumi.OutputState }

func (DefenderStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderStatus)(nil)).Elem()
}

func (o DefenderStatusPtrOutput) ToDefenderStatusPtrOutput() DefenderStatusPtrOutput {
	return o
}

func (o DefenderStatusPtrOutput) ToDefenderStatusPtrOutputWithContext(ctx context.Context) DefenderStatusPtrOutput {
	return o
}

func (o DefenderStatusPtrOutput) Elem() DefenderStatusOutput {
	return o.ApplyT(func(v *DefenderStatus) DefenderStatus {
		if v != nil {
			return *v
		}
		var ret DefenderStatus
		return ret
	}).(DefenderStatusOutput)
}

func (o DefenderStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefenderStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefenderStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DefenderStatusInput interface {
	pulumi.Input

	ToDefenderStatusOutput() DefenderStatusOutput
	ToDefenderStatusOutputWithContext(context.Context) DefenderStatusOutput
}

var defenderStatusPtrType = reflect.TypeOf((**DefenderStatus)(nil)).Elem()

type DefenderStatusPtrInput interface {
	pulumi.Input

	ToDefenderStatusPtrOutput() DefenderStatusPtrOutput
	ToDefenderStatusPtrOutputWithContext(context.Context) DefenderStatusPtrOutput
}

type defenderStatusPtr string

func DefenderStatusPtr(v string) DefenderStatusPtrInput {
	return (*defenderStatusPtr)(&v)
}

func (*defenderStatusPtr) ElementType() reflect.Type {
	return defenderStatusPtrType
}

func (in *defenderStatusPtr) ToDefenderStatusPtrOutput() DefenderStatusPtrOutput {
	return pulumi.ToOutput(in).(DefenderStatusPtrOutput)
}

func (in *defenderStatusPtr) ToDefenderStatusPtrOutputWithContext(ctx context.Context) DefenderStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefenderStatusPtrOutput)
}

type EditionType string

const (
	EditionTypeEvaluation = EditionType("Evaluation")
	EditionTypeEnterprise = EditionType("Enterprise")
	EditionTypeStandard   = EditionType("Standard")
	EditionTypeWeb        = EditionType("Web")
	EditionTypeDeveloper  = EditionType("Developer")
	EditionTypeExpress    = EditionType("Express")
)

func (EditionType) ElementType() reflect.Type {
	return reflect.TypeOf((*EditionType)(nil)).Elem()
}

func (e EditionType) ToEditionTypeOutput() EditionTypeOutput {
	return pulumi.ToOutput(e).(EditionTypeOutput)
}

func (e EditionType) ToEditionTypeOutputWithContext(ctx context.Context) EditionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EditionTypeOutput)
}

func (e EditionType) ToEditionTypePtrOutput() EditionTypePtrOutput {
	return e.ToEditionTypePtrOutputWithContext(context.Background())
}

func (e EditionType) ToEditionTypePtrOutputWithContext(ctx context.Context) EditionTypePtrOutput {
	return EditionType(e).ToEditionTypeOutputWithContext(ctx).ToEditionTypePtrOutputWithContext(ctx)
}

func (e EditionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EditionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EditionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EditionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EditionTypeOutput struct{ *pulumi.OutputState }

func (EditionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EditionType)(nil)).Elem()
}

func (o EditionTypeOutput) ToEditionTypeOutput() EditionTypeOutput {
	return o
}

func (o EditionTypeOutput) ToEditionTypeOutputWithContext(ctx context.Context) EditionTypeOutput {
	return o
}

func (o EditionTypeOutput) ToEditionTypePtrOutput() EditionTypePtrOutput {
	return o.ToEditionTypePtrOutputWithContext(context.Background())
}

func (o EditionTypeOutput) ToEditionTypePtrOutputWithContext(ctx context.Context) EditionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EditionType) *EditionType {
		return &v
	}).(EditionTypePtrOutput)
}

func (o EditionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EditionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EditionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EditionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EditionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EditionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EditionTypePtrOutput struct{ *pulumi.OutputState }

func (EditionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EditionType)(nil)).Elem()
}

func (o EditionTypePtrOutput) ToEditionTypePtrOutput() EditionTypePtrOutput {
	return o
}

func (o EditionTypePtrOutput) ToEditionTypePtrOutputWithContext(ctx context.Context) EditionTypePtrOutput {
	return o
}

func (o EditionTypePtrOutput) Elem() EditionTypeOutput {
	return o.ApplyT(func(v *EditionType) EditionType {
		if v != nil {
			return *v
		}
		var ret EditionType
		return ret
	}).(EditionTypeOutput)
}

func (o EditionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EditionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EditionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EditionTypeInput interface {
	pulumi.Input

	ToEditionTypeOutput() EditionTypeOutput
	ToEditionTypeOutputWithContext(context.Context) EditionTypeOutput
}

var editionTypePtrType = reflect.TypeOf((**EditionType)(nil)).Elem()

type EditionTypePtrInput interface {
	pulumi.Input

	ToEditionTypePtrOutput() EditionTypePtrOutput
	ToEditionTypePtrOutputWithContext(context.Context) EditionTypePtrOutput
}

type editionTypePtr string

func EditionTypePtr(v string) EditionTypePtrInput {
	return (*editionTypePtr)(&v)
}

func (*editionTypePtr) ElementType() reflect.Type {
	return editionTypePtrType
}

func (in *editionTypePtr) ToEditionTypePtrOutput() EditionTypePtrOutput {
	return pulumi.ToOutput(in).(EditionTypePtrOutput)
}

func (in *editionTypePtr) ToEditionTypePtrOutputWithContext(ctx context.Context) EditionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EditionTypePtrOutput)
}

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation = ExtendedLocationTypes("CustomLocation")
)

func (ExtendedLocationTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationTypes)(nil)).Elem()
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput {
	return pulumi.ToOutput(e).(ExtendedLocationTypesOutput)
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesOutputWithContext(ctx context.Context) ExtendedLocationTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExtendedLocationTypesOutput)
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return e.ToExtendedLocationTypesPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return ExtendedLocationTypes(e).ToExtendedLocationTypesOutputWithContext(ctx).ToExtendedLocationTypesPtrOutputWithContext(ctx)
}

func (e ExtendedLocationTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExtendedLocationTypesOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationTypes)(nil)).Elem()
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput {
	return o
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesOutputWithContext(ctx context.Context) ExtendedLocationTypesOutput {
	return o
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return o.ToExtendedLocationTypesPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationTypes) *ExtendedLocationTypes {
		return &v
	}).(ExtendedLocationTypesPtrOutput)
}

func (o ExtendedLocationTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExtendedLocationTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationTypesPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationTypes)(nil)).Elem()
}

func (o ExtendedLocationTypesPtrOutput) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return o
}

func (o ExtendedLocationTypesPtrOutput) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return o
}

func (o ExtendedLocationTypesPtrOutput) Elem() ExtendedLocationTypesOutput {
	return o.ApplyT(func(v *ExtendedLocationTypes) ExtendedLocationTypes {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationTypes
		return ret
	}).(ExtendedLocationTypesOutput)
}

func (o ExtendedLocationTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExtendedLocationTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExtendedLocationTypesInput interface {
	pulumi.Input

	ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput
	ToExtendedLocationTypesOutputWithContext(context.Context) ExtendedLocationTypesOutput
}

var extendedLocationTypesPtrType = reflect.TypeOf((**ExtendedLocationTypes)(nil)).Elem()

type ExtendedLocationTypesPtrInput interface {
	pulumi.Input

	ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput
	ToExtendedLocationTypesPtrOutputWithContext(context.Context) ExtendedLocationTypesPtrOutput
}

type extendedLocationTypesPtr string

func ExtendedLocationTypesPtr(v string) ExtendedLocationTypesPtrInput {
	return (*extendedLocationTypesPtr)(&v)
}

func (*extendedLocationTypesPtr) ElementType() reflect.Type {
	return extendedLocationTypesPtrType
}

func (in *extendedLocationTypesPtr) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return pulumi.ToOutput(in).(ExtendedLocationTypesPtrOutput)
}

func (in *extendedLocationTypesPtr) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExtendedLocationTypesPtrOutput)
}

type Infrastructure string

const (
	InfrastructureAzure      = Infrastructure("azure")
	InfrastructureGcp        = Infrastructure("gcp")
	InfrastructureAws        = Infrastructure("aws")
	InfrastructureAlibaba    = Infrastructure("alibaba")
	InfrastructureOnpremises = Infrastructure("onpremises")
	InfrastructureOther      = Infrastructure("other")
)

func (Infrastructure) ElementType() reflect.Type {
	return reflect.TypeOf((*Infrastructure)(nil)).Elem()
}

func (e Infrastructure) ToInfrastructureOutput() InfrastructureOutput {
	return pulumi.ToOutput(e).(InfrastructureOutput)
}

func (e Infrastructure) ToInfrastructureOutputWithContext(ctx context.Context) InfrastructureOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InfrastructureOutput)
}

func (e Infrastructure) ToInfrastructurePtrOutput() InfrastructurePtrOutput {
	return e.ToInfrastructurePtrOutputWithContext(context.Background())
}

func (e Infrastructure) ToInfrastructurePtrOutputWithContext(ctx context.Context) InfrastructurePtrOutput {
	return Infrastructure(e).ToInfrastructureOutputWithContext(ctx).ToInfrastructurePtrOutputWithContext(ctx)
}

func (e Infrastructure) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Infrastructure) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Infrastructure) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Infrastructure) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InfrastructureOutput struct{ *pulumi.OutputState }

func (InfrastructureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Infrastructure)(nil)).Elem()
}

func (o InfrastructureOutput) ToInfrastructureOutput() InfrastructureOutput {
	return o
}

func (o InfrastructureOutput) ToInfrastructureOutputWithContext(ctx context.Context) InfrastructureOutput {
	return o
}

func (o InfrastructureOutput) ToInfrastructurePtrOutput() InfrastructurePtrOutput {
	return o.ToInfrastructurePtrOutputWithContext(context.Background())
}

func (o InfrastructureOutput) ToInfrastructurePtrOutputWithContext(ctx context.Context) InfrastructurePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Infrastructure) *Infrastructure {
		return &v
	}).(InfrastructurePtrOutput)
}

func (o InfrastructureOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InfrastructureOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Infrastructure) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InfrastructureOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Infrastructure) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InfrastructurePtrOutput struct{ *pulumi.OutputState }

func (InfrastructurePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Infrastructure)(nil)).Elem()
}

func (o InfrastructurePtrOutput) ToInfrastructurePtrOutput() InfrastructurePtrOutput {
	return o
}

func (o InfrastructurePtrOutput) ToInfrastructurePtrOutputWithContext(ctx context.Context) InfrastructurePtrOutput {
	return o
}

func (o InfrastructurePtrOutput) Elem() InfrastructureOutput {
	return o.ApplyT(func(v *Infrastructure) Infrastructure {
		if v != nil {
			return *v
		}
		var ret Infrastructure
		return ret
	}).(InfrastructureOutput)
}

func (o InfrastructurePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructurePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Infrastructure) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InfrastructureInput interface {
	pulumi.Input

	ToInfrastructureOutput() InfrastructureOutput
	ToInfrastructureOutputWithContext(context.Context) InfrastructureOutput
}

var infrastructurePtrType = reflect.TypeOf((**Infrastructure)(nil)).Elem()

type InfrastructurePtrInput interface {
	pulumi.Input

	ToInfrastructurePtrOutput() InfrastructurePtrOutput
	ToInfrastructurePtrOutputWithContext(context.Context) InfrastructurePtrOutput
}

type infrastructurePtr string

func InfrastructurePtr(v string) InfrastructurePtrInput {
	return (*infrastructurePtr)(&v)
}

func (*infrastructurePtr) ElementType() reflect.Type {
	return infrastructurePtrType
}

func (in *infrastructurePtr) ToInfrastructurePtrOutput() InfrastructurePtrOutput {
	return pulumi.ToOutput(in).(InfrastructurePtrOutput)
}

func (in *infrastructurePtr) ToInfrastructurePtrOutputWithContext(ctx context.Context) InfrastructurePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InfrastructurePtrOutput)
}

type SqlManagedInstanceSkuTier string

const (
	SqlManagedInstanceSkuTierGeneralPurpose   = SqlManagedInstanceSkuTier("GeneralPurpose")
	SqlManagedInstanceSkuTierBusinessCritical = SqlManagedInstanceSkuTier("BusinessCritical")
)

func (SqlManagedInstanceSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSkuTier)(nil)).Elem()
}

func (e SqlManagedInstanceSkuTier) ToSqlManagedInstanceSkuTierOutput() SqlManagedInstanceSkuTierOutput {
	return pulumi.ToOutput(e).(SqlManagedInstanceSkuTierOutput)
}

func (e SqlManagedInstanceSkuTier) ToSqlManagedInstanceSkuTierOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SqlManagedInstanceSkuTierOutput)
}

func (e SqlManagedInstanceSkuTier) ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput {
	return e.ToSqlManagedInstanceSkuTierPtrOutputWithContext(context.Background())
}

func (e SqlManagedInstanceSkuTier) ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierPtrOutput {
	return SqlManagedInstanceSkuTier(e).ToSqlManagedInstanceSkuTierOutputWithContext(ctx).ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx)
}

func (e SqlManagedInstanceSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlManagedInstanceSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlManagedInstanceSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SqlManagedInstanceSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SqlManagedInstanceSkuTierOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSkuTier)(nil)).Elem()
}

func (o SqlManagedInstanceSkuTierOutput) ToSqlManagedInstanceSkuTierOutput() SqlManagedInstanceSkuTierOutput {
	return o
}

func (o SqlManagedInstanceSkuTierOutput) ToSqlManagedInstanceSkuTierOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierOutput {
	return o
}

func (o SqlManagedInstanceSkuTierOutput) ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput {
	return o.ToSqlManagedInstanceSkuTierPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuTierOutput) ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstanceSkuTier) *SqlManagedInstanceSkuTier {
		return &v
	}).(SqlManagedInstanceSkuTierPtrOutput)
}

func (o SqlManagedInstanceSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlManagedInstanceSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SqlManagedInstanceSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlManagedInstanceSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SqlManagedInstanceSkuTierPtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSkuTier)(nil)).Elem()
}

func (o SqlManagedInstanceSkuTierPtrOutput) ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput {
	return o
}

func (o SqlManagedInstanceSkuTierPtrOutput) ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierPtrOutput {
	return o
}

func (o SqlManagedInstanceSkuTierPtrOutput) Elem() SqlManagedInstanceSkuTierOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuTier) SqlManagedInstanceSkuTier {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceSkuTier
		return ret
	}).(SqlManagedInstanceSkuTierOutput)
}

func (o SqlManagedInstanceSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SqlManagedInstanceSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SqlManagedInstanceSkuTierInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuTierOutput() SqlManagedInstanceSkuTierOutput
	ToSqlManagedInstanceSkuTierOutputWithContext(context.Context) SqlManagedInstanceSkuTierOutput
}

var sqlManagedInstanceSkuTierPtrType = reflect.TypeOf((**SqlManagedInstanceSkuTier)(nil)).Elem()

type SqlManagedInstanceSkuTierPtrInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput
	ToSqlManagedInstanceSkuTierPtrOutputWithContext(context.Context) SqlManagedInstanceSkuTierPtrOutput
}

type sqlManagedInstanceSkuTierPtr string

func SqlManagedInstanceSkuTierPtr(v string) SqlManagedInstanceSkuTierPtrInput {
	return (*sqlManagedInstanceSkuTierPtr)(&v)
}

func (*sqlManagedInstanceSkuTierPtr) ElementType() reflect.Type {
	return sqlManagedInstanceSkuTierPtrType
}

func (in *sqlManagedInstanceSkuTierPtr) ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput {
	return pulumi.ToOutput(in).(SqlManagedInstanceSkuTierPtrOutput)
}

func (in *sqlManagedInstanceSkuTierPtr) ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SqlManagedInstanceSkuTierPtrOutput)
}

type SqlVersion string

const (
	SqlVersion_SQL_Server_2019 = SqlVersion("SQL Server 2019")
	SqlVersion_SQL_Server_2017 = SqlVersion("SQL Server 2017")
	SqlVersion_SQL_Server_2016 = SqlVersion("SQL Server 2016")
)

func (SqlVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlVersion)(nil)).Elem()
}

func (e SqlVersion) ToSqlVersionOutput() SqlVersionOutput {
	return pulumi.ToOutput(e).(SqlVersionOutput)
}

func (e SqlVersion) ToSqlVersionOutputWithContext(ctx context.Context) SqlVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SqlVersionOutput)
}

func (e SqlVersion) ToSqlVersionPtrOutput() SqlVersionPtrOutput {
	return e.ToSqlVersionPtrOutputWithContext(context.Background())
}

func (e SqlVersion) ToSqlVersionPtrOutputWithContext(ctx context.Context) SqlVersionPtrOutput {
	return SqlVersion(e).ToSqlVersionOutputWithContext(ctx).ToSqlVersionPtrOutputWithContext(ctx)
}

func (e SqlVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SqlVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SqlVersionOutput struct{ *pulumi.OutputState }

func (SqlVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlVersion)(nil)).Elem()
}

func (o SqlVersionOutput) ToSqlVersionOutput() SqlVersionOutput {
	return o
}

func (o SqlVersionOutput) ToSqlVersionOutputWithContext(ctx context.Context) SqlVersionOutput {
	return o
}

func (o SqlVersionOutput) ToSqlVersionPtrOutput() SqlVersionPtrOutput {
	return o.ToSqlVersionPtrOutputWithContext(context.Background())
}

func (o SqlVersionOutput) ToSqlVersionPtrOutputWithContext(ctx context.Context) SqlVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlVersion) *SqlVersion {
		return &v
	}).(SqlVersionPtrOutput)
}

func (o SqlVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SqlVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SqlVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SqlVersionPtrOutput struct{ *pulumi.OutputState }

func (SqlVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVersion)(nil)).Elem()
}

func (o SqlVersionPtrOutput) ToSqlVersionPtrOutput() SqlVersionPtrOutput {
	return o
}

func (o SqlVersionPtrOutput) ToSqlVersionPtrOutputWithContext(ctx context.Context) SqlVersionPtrOutput {
	return o
}

func (o SqlVersionPtrOutput) Elem() SqlVersionOutput {
	return o.ApplyT(func(v *SqlVersion) SqlVersion {
		if v != nil {
			return *v
		}
		var ret SqlVersion
		return ret
	}).(SqlVersionOutput)
}

func (o SqlVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SqlVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SqlVersionInput interface {
	pulumi.Input

	ToSqlVersionOutput() SqlVersionOutput
	ToSqlVersionOutputWithContext(context.Context) SqlVersionOutput
}

var sqlVersionPtrType = reflect.TypeOf((**SqlVersion)(nil)).Elem()

type SqlVersionPtrInput interface {
	pulumi.Input

	ToSqlVersionPtrOutput() SqlVersionPtrOutput
	ToSqlVersionPtrOutputWithContext(context.Context) SqlVersionPtrOutput
}

type sqlVersionPtr string

func SqlVersionPtr(v string) SqlVersionPtrInput {
	return (*sqlVersionPtr)(&v)
}

func (*sqlVersionPtr) ElementType() reflect.Type {
	return sqlVersionPtrType
}

func (in *sqlVersionPtr) ToSqlVersionPtrOutput() SqlVersionPtrOutput {
	return pulumi.ToOutput(in).(SqlVersionPtrOutput)
}

func (in *sqlVersionPtr) ToSqlVersionPtrOutputWithContext(ctx context.Context) SqlVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SqlVersionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ArcSqlManagedInstanceLicenseTypeOutput{})
	pulumi.RegisterOutputType(ArcSqlManagedInstanceLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(ArcSqlServerLicenseTypeOutput{})
	pulumi.RegisterOutputType(ArcSqlServerLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStatusOutput{})
	pulumi.RegisterOutputType(ConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(DefenderStatusOutput{})
	pulumi.RegisterOutputType(DefenderStatusPtrOutput{})
	pulumi.RegisterOutputType(EditionTypeOutput{})
	pulumi.RegisterOutputType(EditionTypePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesPtrOutput{})
	pulumi.RegisterOutputType(InfrastructureOutput{})
	pulumi.RegisterOutputType(InfrastructurePtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuTierOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuTierPtrOutput{})
	pulumi.RegisterOutputType(SqlVersionOutput{})
	pulumi.RegisterOutputType(SqlVersionPtrOutput{})
}
