


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

type ConnectionStatus string

const (
	ConnectionStatusConnected    = ConnectionStatus("Connected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
	ConnectionStatusUnknown      = ConnectionStatus("Unknown")
)

type DefenderStatus string

const (
	DefenderStatusProtected   = DefenderStatus("Protected")
	DefenderStatusUnprotected = DefenderStatus("Unprotected")
	DefenderStatusUnknown     = DefenderStatus("Unknown")
)

type EditionType string

const (
	EditionTypeEvaluation = EditionType("Evaluation")
	EditionTypeEnterprise = EditionType("Enterprise")
	EditionTypeStandard   = EditionType("Standard")
	EditionTypeWeb        = EditionType("Web")
	EditionTypeDeveloper  = EditionType("Developer")
	EditionTypeExpress    = EditionType("Express")
)

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation = ExtendedLocationTypes("CustomLocation")
)

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

func init() {
	pulumi.RegisterOutputType(ArcSqlManagedInstanceLicenseTypeOutput{})
	pulumi.RegisterOutputType(ArcSqlManagedInstanceLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(InfrastructureOutput{})
	pulumi.RegisterOutputType(InfrastructurePtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuTierOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuTierPtrOutput{})
}
