


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeRecovery           = CreateMode("Recovery")
	CreateModeRestore            = CreateMode("Restore")
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

type DataFlowComputeType string

const (
	DataFlowComputeTypeGeneral          = DataFlowComputeType("General")
	DataFlowComputeTypeMemoryOptimized  = DataFlowComputeType("MemoryOptimized")
	DataFlowComputeTypeComputeOptimized = DataFlowComputeType("ComputeOptimized")
)

func (DataFlowComputeType) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlowComputeType)(nil)).Elem()
}

func (e DataFlowComputeType) ToDataFlowComputeTypeOutput() DataFlowComputeTypeOutput {
	return pulumi.ToOutput(e).(DataFlowComputeTypeOutput)
}

func (e DataFlowComputeType) ToDataFlowComputeTypeOutputWithContext(ctx context.Context) DataFlowComputeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataFlowComputeTypeOutput)
}

func (e DataFlowComputeType) ToDataFlowComputeTypePtrOutput() DataFlowComputeTypePtrOutput {
	return e.ToDataFlowComputeTypePtrOutputWithContext(context.Background())
}

func (e DataFlowComputeType) ToDataFlowComputeTypePtrOutputWithContext(ctx context.Context) DataFlowComputeTypePtrOutput {
	return DataFlowComputeType(e).ToDataFlowComputeTypeOutputWithContext(ctx).ToDataFlowComputeTypePtrOutputWithContext(ctx)
}

func (e DataFlowComputeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataFlowComputeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataFlowComputeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataFlowComputeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataFlowComputeTypeOutput struct{ *pulumi.OutputState }

func (DataFlowComputeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlowComputeType)(nil)).Elem()
}

func (o DataFlowComputeTypeOutput) ToDataFlowComputeTypeOutput() DataFlowComputeTypeOutput {
	return o
}

func (o DataFlowComputeTypeOutput) ToDataFlowComputeTypeOutputWithContext(ctx context.Context) DataFlowComputeTypeOutput {
	return o
}

func (o DataFlowComputeTypeOutput) ToDataFlowComputeTypePtrOutput() DataFlowComputeTypePtrOutput {
	return o.ToDataFlowComputeTypePtrOutputWithContext(context.Background())
}

func (o DataFlowComputeTypeOutput) ToDataFlowComputeTypePtrOutputWithContext(ctx context.Context) DataFlowComputeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataFlowComputeType) *DataFlowComputeType {
		return &v
	}).(DataFlowComputeTypePtrOutput)
}

func (o DataFlowComputeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataFlowComputeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataFlowComputeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataFlowComputeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataFlowComputeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataFlowComputeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataFlowComputeTypePtrOutput struct{ *pulumi.OutputState }

func (DataFlowComputeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFlowComputeType)(nil)).Elem()
}

func (o DataFlowComputeTypePtrOutput) ToDataFlowComputeTypePtrOutput() DataFlowComputeTypePtrOutput {
	return o
}

func (o DataFlowComputeTypePtrOutput) ToDataFlowComputeTypePtrOutputWithContext(ctx context.Context) DataFlowComputeTypePtrOutput {
	return o
}

func (o DataFlowComputeTypePtrOutput) Elem() DataFlowComputeTypeOutput {
	return o.ApplyT(func(v *DataFlowComputeType) DataFlowComputeType {
		if v != nil {
			return *v
		}
		var ret DataFlowComputeType
		return ret
	}).(DataFlowComputeTypeOutput)
}

func (o DataFlowComputeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataFlowComputeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataFlowComputeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataFlowComputeTypeInput interface {
	pulumi.Input

	ToDataFlowComputeTypeOutput() DataFlowComputeTypeOutput
	ToDataFlowComputeTypeOutputWithContext(context.Context) DataFlowComputeTypeOutput
}

var dataFlowComputeTypePtrType = reflect.TypeOf((**DataFlowComputeType)(nil)).Elem()

type DataFlowComputeTypePtrInput interface {
	pulumi.Input

	ToDataFlowComputeTypePtrOutput() DataFlowComputeTypePtrOutput
	ToDataFlowComputeTypePtrOutputWithContext(context.Context) DataFlowComputeTypePtrOutput
}

type dataFlowComputeTypePtr string

func DataFlowComputeTypePtr(v string) DataFlowComputeTypePtrInput {
	return (*dataFlowComputeTypePtr)(&v)
}

func (*dataFlowComputeTypePtr) ElementType() reflect.Type {
	return dataFlowComputeTypePtrType
}

func (in *dataFlowComputeTypePtr) ToDataFlowComputeTypePtrOutput() DataFlowComputeTypePtrOutput {
	return pulumi.ToOutput(in).(DataFlowComputeTypePtrOutput)
}

func (in *dataFlowComputeTypePtr) ToDataFlowComputeTypePtrOutputWithContext(ctx context.Context) DataFlowComputeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataFlowComputeTypePtrOutput)
}

type IntegrationRuntimeEdition string

const (
	IntegrationRuntimeEditionStandard   = IntegrationRuntimeEdition("Standard")
	IntegrationRuntimeEditionEnterprise = IntegrationRuntimeEdition("Enterprise")
)

func (IntegrationRuntimeEdition) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeEdition)(nil)).Elem()
}

func (e IntegrationRuntimeEdition) ToIntegrationRuntimeEditionOutput() IntegrationRuntimeEditionOutput {
	return pulumi.ToOutput(e).(IntegrationRuntimeEditionOutput)
}

func (e IntegrationRuntimeEdition) ToIntegrationRuntimeEditionOutputWithContext(ctx context.Context) IntegrationRuntimeEditionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IntegrationRuntimeEditionOutput)
}

func (e IntegrationRuntimeEdition) ToIntegrationRuntimeEditionPtrOutput() IntegrationRuntimeEditionPtrOutput {
	return e.ToIntegrationRuntimeEditionPtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeEdition) ToIntegrationRuntimeEditionPtrOutputWithContext(ctx context.Context) IntegrationRuntimeEditionPtrOutput {
	return IntegrationRuntimeEdition(e).ToIntegrationRuntimeEditionOutputWithContext(ctx).ToIntegrationRuntimeEditionPtrOutputWithContext(ctx)
}

func (e IntegrationRuntimeEdition) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeEdition) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeEdition) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeEdition) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntegrationRuntimeEditionOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeEditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeEdition)(nil)).Elem()
}

func (o IntegrationRuntimeEditionOutput) ToIntegrationRuntimeEditionOutput() IntegrationRuntimeEditionOutput {
	return o
}

func (o IntegrationRuntimeEditionOutput) ToIntegrationRuntimeEditionOutputWithContext(ctx context.Context) IntegrationRuntimeEditionOutput {
	return o
}

func (o IntegrationRuntimeEditionOutput) ToIntegrationRuntimeEditionPtrOutput() IntegrationRuntimeEditionPtrOutput {
	return o.ToIntegrationRuntimeEditionPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeEditionOutput) ToIntegrationRuntimeEditionPtrOutputWithContext(ctx context.Context) IntegrationRuntimeEditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationRuntimeEdition) *IntegrationRuntimeEdition {
		return &v
	}).(IntegrationRuntimeEditionPtrOutput)
}

func (o IntegrationRuntimeEditionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IntegrationRuntimeEditionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeEdition) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IntegrationRuntimeEditionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeEditionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeEdition) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IntegrationRuntimeEditionPtrOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeEditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeEdition)(nil)).Elem()
}

func (o IntegrationRuntimeEditionPtrOutput) ToIntegrationRuntimeEditionPtrOutput() IntegrationRuntimeEditionPtrOutput {
	return o
}

func (o IntegrationRuntimeEditionPtrOutput) ToIntegrationRuntimeEditionPtrOutputWithContext(ctx context.Context) IntegrationRuntimeEditionPtrOutput {
	return o
}

func (o IntegrationRuntimeEditionPtrOutput) Elem() IntegrationRuntimeEditionOutput {
	return o.ApplyT(func(v *IntegrationRuntimeEdition) IntegrationRuntimeEdition {
		if v != nil {
			return *v
		}
		var ret IntegrationRuntimeEdition
		return ret
	}).(IntegrationRuntimeEditionOutput)
}

func (o IntegrationRuntimeEditionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeEditionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IntegrationRuntimeEdition) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IntegrationRuntimeEditionInput interface {
	pulumi.Input

	ToIntegrationRuntimeEditionOutput() IntegrationRuntimeEditionOutput
	ToIntegrationRuntimeEditionOutputWithContext(context.Context) IntegrationRuntimeEditionOutput
}

var integrationRuntimeEditionPtrType = reflect.TypeOf((**IntegrationRuntimeEdition)(nil)).Elem()

type IntegrationRuntimeEditionPtrInput interface {
	pulumi.Input

	ToIntegrationRuntimeEditionPtrOutput() IntegrationRuntimeEditionPtrOutput
	ToIntegrationRuntimeEditionPtrOutputWithContext(context.Context) IntegrationRuntimeEditionPtrOutput
}

type integrationRuntimeEditionPtr string

func IntegrationRuntimeEditionPtr(v string) IntegrationRuntimeEditionPtrInput {
	return (*integrationRuntimeEditionPtr)(&v)
}

func (*integrationRuntimeEditionPtr) ElementType() reflect.Type {
	return integrationRuntimeEditionPtrType
}

func (in *integrationRuntimeEditionPtr) ToIntegrationRuntimeEditionPtrOutput() IntegrationRuntimeEditionPtrOutput {
	return pulumi.ToOutput(in).(IntegrationRuntimeEditionPtrOutput)
}

func (in *integrationRuntimeEditionPtr) ToIntegrationRuntimeEditionPtrOutputWithContext(ctx context.Context) IntegrationRuntimeEditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IntegrationRuntimeEditionPtrOutput)
}

type IntegrationRuntimeEntityReferenceType string

const (
	IntegrationRuntimeEntityReferenceTypeIntegrationRuntimeReference = IntegrationRuntimeEntityReferenceType("IntegrationRuntimeReference")
	IntegrationRuntimeEntityReferenceTypeLinkedServiceReference      = IntegrationRuntimeEntityReferenceType("LinkedServiceReference")
)

func (IntegrationRuntimeEntityReferenceType) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeEntityReferenceType)(nil)).Elem()
}

func (e IntegrationRuntimeEntityReferenceType) ToIntegrationRuntimeEntityReferenceTypeOutput() IntegrationRuntimeEntityReferenceTypeOutput {
	return pulumi.ToOutput(e).(IntegrationRuntimeEntityReferenceTypeOutput)
}

func (e IntegrationRuntimeEntityReferenceType) ToIntegrationRuntimeEntityReferenceTypeOutputWithContext(ctx context.Context) IntegrationRuntimeEntityReferenceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IntegrationRuntimeEntityReferenceTypeOutput)
}

func (e IntegrationRuntimeEntityReferenceType) ToIntegrationRuntimeEntityReferenceTypePtrOutput() IntegrationRuntimeEntityReferenceTypePtrOutput {
	return e.ToIntegrationRuntimeEntityReferenceTypePtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeEntityReferenceType) ToIntegrationRuntimeEntityReferenceTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeEntityReferenceTypePtrOutput {
	return IntegrationRuntimeEntityReferenceType(e).ToIntegrationRuntimeEntityReferenceTypeOutputWithContext(ctx).ToIntegrationRuntimeEntityReferenceTypePtrOutputWithContext(ctx)
}

func (e IntegrationRuntimeEntityReferenceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeEntityReferenceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeEntityReferenceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeEntityReferenceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntegrationRuntimeEntityReferenceTypeOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeEntityReferenceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeEntityReferenceType)(nil)).Elem()
}

func (o IntegrationRuntimeEntityReferenceTypeOutput) ToIntegrationRuntimeEntityReferenceTypeOutput() IntegrationRuntimeEntityReferenceTypeOutput {
	return o
}

func (o IntegrationRuntimeEntityReferenceTypeOutput) ToIntegrationRuntimeEntityReferenceTypeOutputWithContext(ctx context.Context) IntegrationRuntimeEntityReferenceTypeOutput {
	return o
}

func (o IntegrationRuntimeEntityReferenceTypeOutput) ToIntegrationRuntimeEntityReferenceTypePtrOutput() IntegrationRuntimeEntityReferenceTypePtrOutput {
	return o.ToIntegrationRuntimeEntityReferenceTypePtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeEntityReferenceTypeOutput) ToIntegrationRuntimeEntityReferenceTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeEntityReferenceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationRuntimeEntityReferenceType) *IntegrationRuntimeEntityReferenceType {
		return &v
	}).(IntegrationRuntimeEntityReferenceTypePtrOutput)
}

func (o IntegrationRuntimeEntityReferenceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IntegrationRuntimeEntityReferenceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeEntityReferenceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IntegrationRuntimeEntityReferenceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeEntityReferenceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeEntityReferenceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IntegrationRuntimeEntityReferenceTypePtrOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeEntityReferenceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeEntityReferenceType)(nil)).Elem()
}

func (o IntegrationRuntimeEntityReferenceTypePtrOutput) ToIntegrationRuntimeEntityReferenceTypePtrOutput() IntegrationRuntimeEntityReferenceTypePtrOutput {
	return o
}

func (o IntegrationRuntimeEntityReferenceTypePtrOutput) ToIntegrationRuntimeEntityReferenceTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeEntityReferenceTypePtrOutput {
	return o
}

func (o IntegrationRuntimeEntityReferenceTypePtrOutput) Elem() IntegrationRuntimeEntityReferenceTypeOutput {
	return o.ApplyT(func(v *IntegrationRuntimeEntityReferenceType) IntegrationRuntimeEntityReferenceType {
		if v != nil {
			return *v
		}
		var ret IntegrationRuntimeEntityReferenceType
		return ret
	}).(IntegrationRuntimeEntityReferenceTypeOutput)
}

func (o IntegrationRuntimeEntityReferenceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeEntityReferenceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IntegrationRuntimeEntityReferenceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IntegrationRuntimeEntityReferenceTypeInput interface {
	pulumi.Input

	ToIntegrationRuntimeEntityReferenceTypeOutput() IntegrationRuntimeEntityReferenceTypeOutput
	ToIntegrationRuntimeEntityReferenceTypeOutputWithContext(context.Context) IntegrationRuntimeEntityReferenceTypeOutput
}

var integrationRuntimeEntityReferenceTypePtrType = reflect.TypeOf((**IntegrationRuntimeEntityReferenceType)(nil)).Elem()

type IntegrationRuntimeEntityReferenceTypePtrInput interface {
	pulumi.Input

	ToIntegrationRuntimeEntityReferenceTypePtrOutput() IntegrationRuntimeEntityReferenceTypePtrOutput
	ToIntegrationRuntimeEntityReferenceTypePtrOutputWithContext(context.Context) IntegrationRuntimeEntityReferenceTypePtrOutput
}

type integrationRuntimeEntityReferenceTypePtr string

func IntegrationRuntimeEntityReferenceTypePtr(v string) IntegrationRuntimeEntityReferenceTypePtrInput {
	return (*integrationRuntimeEntityReferenceTypePtr)(&v)
}

func (*integrationRuntimeEntityReferenceTypePtr) ElementType() reflect.Type {
	return integrationRuntimeEntityReferenceTypePtrType
}

func (in *integrationRuntimeEntityReferenceTypePtr) ToIntegrationRuntimeEntityReferenceTypePtrOutput() IntegrationRuntimeEntityReferenceTypePtrOutput {
	return pulumi.ToOutput(in).(IntegrationRuntimeEntityReferenceTypePtrOutput)
}

func (in *integrationRuntimeEntityReferenceTypePtr) ToIntegrationRuntimeEntityReferenceTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeEntityReferenceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IntegrationRuntimeEntityReferenceTypePtrOutput)
}

type IntegrationRuntimeLicenseType string

const (
	IntegrationRuntimeLicenseTypeBasePrice       = IntegrationRuntimeLicenseType("BasePrice")
	IntegrationRuntimeLicenseTypeLicenseIncluded = IntegrationRuntimeLicenseType("LicenseIncluded")
)

func (IntegrationRuntimeLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeLicenseType)(nil)).Elem()
}

func (e IntegrationRuntimeLicenseType) ToIntegrationRuntimeLicenseTypeOutput() IntegrationRuntimeLicenseTypeOutput {
	return pulumi.ToOutput(e).(IntegrationRuntimeLicenseTypeOutput)
}

func (e IntegrationRuntimeLicenseType) ToIntegrationRuntimeLicenseTypeOutputWithContext(ctx context.Context) IntegrationRuntimeLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IntegrationRuntimeLicenseTypeOutput)
}

func (e IntegrationRuntimeLicenseType) ToIntegrationRuntimeLicenseTypePtrOutput() IntegrationRuntimeLicenseTypePtrOutput {
	return e.ToIntegrationRuntimeLicenseTypePtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeLicenseType) ToIntegrationRuntimeLicenseTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeLicenseTypePtrOutput {
	return IntegrationRuntimeLicenseType(e).ToIntegrationRuntimeLicenseTypeOutputWithContext(ctx).ToIntegrationRuntimeLicenseTypePtrOutputWithContext(ctx)
}

func (e IntegrationRuntimeLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntegrationRuntimeLicenseTypeOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeLicenseType)(nil)).Elem()
}

func (o IntegrationRuntimeLicenseTypeOutput) ToIntegrationRuntimeLicenseTypeOutput() IntegrationRuntimeLicenseTypeOutput {
	return o
}

func (o IntegrationRuntimeLicenseTypeOutput) ToIntegrationRuntimeLicenseTypeOutputWithContext(ctx context.Context) IntegrationRuntimeLicenseTypeOutput {
	return o
}

func (o IntegrationRuntimeLicenseTypeOutput) ToIntegrationRuntimeLicenseTypePtrOutput() IntegrationRuntimeLicenseTypePtrOutput {
	return o.ToIntegrationRuntimeLicenseTypePtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeLicenseTypeOutput) ToIntegrationRuntimeLicenseTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationRuntimeLicenseType) *IntegrationRuntimeLicenseType {
		return &v
	}).(IntegrationRuntimeLicenseTypePtrOutput)
}

func (o IntegrationRuntimeLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IntegrationRuntimeLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IntegrationRuntimeLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IntegrationRuntimeLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeLicenseType)(nil)).Elem()
}

func (o IntegrationRuntimeLicenseTypePtrOutput) ToIntegrationRuntimeLicenseTypePtrOutput() IntegrationRuntimeLicenseTypePtrOutput {
	return o
}

func (o IntegrationRuntimeLicenseTypePtrOutput) ToIntegrationRuntimeLicenseTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeLicenseTypePtrOutput {
	return o
}

func (o IntegrationRuntimeLicenseTypePtrOutput) Elem() IntegrationRuntimeLicenseTypeOutput {
	return o.ApplyT(func(v *IntegrationRuntimeLicenseType) IntegrationRuntimeLicenseType {
		if v != nil {
			return *v
		}
		var ret IntegrationRuntimeLicenseType
		return ret
	}).(IntegrationRuntimeLicenseTypeOutput)
}

func (o IntegrationRuntimeLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IntegrationRuntimeLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IntegrationRuntimeLicenseTypeInput interface {
	pulumi.Input

	ToIntegrationRuntimeLicenseTypeOutput() IntegrationRuntimeLicenseTypeOutput
	ToIntegrationRuntimeLicenseTypeOutputWithContext(context.Context) IntegrationRuntimeLicenseTypeOutput
}

var integrationRuntimeLicenseTypePtrType = reflect.TypeOf((**IntegrationRuntimeLicenseType)(nil)).Elem()

type IntegrationRuntimeLicenseTypePtrInput interface {
	pulumi.Input

	ToIntegrationRuntimeLicenseTypePtrOutput() IntegrationRuntimeLicenseTypePtrOutput
	ToIntegrationRuntimeLicenseTypePtrOutputWithContext(context.Context) IntegrationRuntimeLicenseTypePtrOutput
}

type integrationRuntimeLicenseTypePtr string

func IntegrationRuntimeLicenseTypePtr(v string) IntegrationRuntimeLicenseTypePtrInput {
	return (*integrationRuntimeLicenseTypePtr)(&v)
}

func (*integrationRuntimeLicenseTypePtr) ElementType() reflect.Type {
	return integrationRuntimeLicenseTypePtrType
}

func (in *integrationRuntimeLicenseTypePtr) ToIntegrationRuntimeLicenseTypePtrOutput() IntegrationRuntimeLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(IntegrationRuntimeLicenseTypePtrOutput)
}

func (in *integrationRuntimeLicenseTypePtr) ToIntegrationRuntimeLicenseTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IntegrationRuntimeLicenseTypePtrOutput)
}

type IntegrationRuntimeSsisCatalogPricingTier string

const (
	IntegrationRuntimeSsisCatalogPricingTierBasic     = IntegrationRuntimeSsisCatalogPricingTier("Basic")
	IntegrationRuntimeSsisCatalogPricingTierStandard  = IntegrationRuntimeSsisCatalogPricingTier("Standard")
	IntegrationRuntimeSsisCatalogPricingTierPremium   = IntegrationRuntimeSsisCatalogPricingTier("Premium")
	IntegrationRuntimeSsisCatalogPricingTierPremiumRS = IntegrationRuntimeSsisCatalogPricingTier("PremiumRS")
)

func (IntegrationRuntimeSsisCatalogPricingTier) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeSsisCatalogPricingTier)(nil)).Elem()
}

func (e IntegrationRuntimeSsisCatalogPricingTier) ToIntegrationRuntimeSsisCatalogPricingTierOutput() IntegrationRuntimeSsisCatalogPricingTierOutput {
	return pulumi.ToOutput(e).(IntegrationRuntimeSsisCatalogPricingTierOutput)
}

func (e IntegrationRuntimeSsisCatalogPricingTier) ToIntegrationRuntimeSsisCatalogPricingTierOutputWithContext(ctx context.Context) IntegrationRuntimeSsisCatalogPricingTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IntegrationRuntimeSsisCatalogPricingTierOutput)
}

func (e IntegrationRuntimeSsisCatalogPricingTier) ToIntegrationRuntimeSsisCatalogPricingTierPtrOutput() IntegrationRuntimeSsisCatalogPricingTierPtrOutput {
	return e.ToIntegrationRuntimeSsisCatalogPricingTierPtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeSsisCatalogPricingTier) ToIntegrationRuntimeSsisCatalogPricingTierPtrOutputWithContext(ctx context.Context) IntegrationRuntimeSsisCatalogPricingTierPtrOutput {
	return IntegrationRuntimeSsisCatalogPricingTier(e).ToIntegrationRuntimeSsisCatalogPricingTierOutputWithContext(ctx).ToIntegrationRuntimeSsisCatalogPricingTierPtrOutputWithContext(ctx)
}

func (e IntegrationRuntimeSsisCatalogPricingTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeSsisCatalogPricingTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeSsisCatalogPricingTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeSsisCatalogPricingTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntegrationRuntimeSsisCatalogPricingTierOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeSsisCatalogPricingTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeSsisCatalogPricingTier)(nil)).Elem()
}

func (o IntegrationRuntimeSsisCatalogPricingTierOutput) ToIntegrationRuntimeSsisCatalogPricingTierOutput() IntegrationRuntimeSsisCatalogPricingTierOutput {
	return o
}

func (o IntegrationRuntimeSsisCatalogPricingTierOutput) ToIntegrationRuntimeSsisCatalogPricingTierOutputWithContext(ctx context.Context) IntegrationRuntimeSsisCatalogPricingTierOutput {
	return o
}

func (o IntegrationRuntimeSsisCatalogPricingTierOutput) ToIntegrationRuntimeSsisCatalogPricingTierPtrOutput() IntegrationRuntimeSsisCatalogPricingTierPtrOutput {
	return o.ToIntegrationRuntimeSsisCatalogPricingTierPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeSsisCatalogPricingTierOutput) ToIntegrationRuntimeSsisCatalogPricingTierPtrOutputWithContext(ctx context.Context) IntegrationRuntimeSsisCatalogPricingTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationRuntimeSsisCatalogPricingTier) *IntegrationRuntimeSsisCatalogPricingTier {
		return &v
	}).(IntegrationRuntimeSsisCatalogPricingTierPtrOutput)
}

func (o IntegrationRuntimeSsisCatalogPricingTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IntegrationRuntimeSsisCatalogPricingTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeSsisCatalogPricingTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IntegrationRuntimeSsisCatalogPricingTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeSsisCatalogPricingTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeSsisCatalogPricingTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IntegrationRuntimeSsisCatalogPricingTierPtrOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeSsisCatalogPricingTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeSsisCatalogPricingTier)(nil)).Elem()
}

func (o IntegrationRuntimeSsisCatalogPricingTierPtrOutput) ToIntegrationRuntimeSsisCatalogPricingTierPtrOutput() IntegrationRuntimeSsisCatalogPricingTierPtrOutput {
	return o
}

func (o IntegrationRuntimeSsisCatalogPricingTierPtrOutput) ToIntegrationRuntimeSsisCatalogPricingTierPtrOutputWithContext(ctx context.Context) IntegrationRuntimeSsisCatalogPricingTierPtrOutput {
	return o
}

func (o IntegrationRuntimeSsisCatalogPricingTierPtrOutput) Elem() IntegrationRuntimeSsisCatalogPricingTierOutput {
	return o.ApplyT(func(v *IntegrationRuntimeSsisCatalogPricingTier) IntegrationRuntimeSsisCatalogPricingTier {
		if v != nil {
			return *v
		}
		var ret IntegrationRuntimeSsisCatalogPricingTier
		return ret
	}).(IntegrationRuntimeSsisCatalogPricingTierOutput)
}

func (o IntegrationRuntimeSsisCatalogPricingTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeSsisCatalogPricingTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IntegrationRuntimeSsisCatalogPricingTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IntegrationRuntimeSsisCatalogPricingTierInput interface {
	pulumi.Input

	ToIntegrationRuntimeSsisCatalogPricingTierOutput() IntegrationRuntimeSsisCatalogPricingTierOutput
	ToIntegrationRuntimeSsisCatalogPricingTierOutputWithContext(context.Context) IntegrationRuntimeSsisCatalogPricingTierOutput
}

var integrationRuntimeSsisCatalogPricingTierPtrType = reflect.TypeOf((**IntegrationRuntimeSsisCatalogPricingTier)(nil)).Elem()

type IntegrationRuntimeSsisCatalogPricingTierPtrInput interface {
	pulumi.Input

	ToIntegrationRuntimeSsisCatalogPricingTierPtrOutput() IntegrationRuntimeSsisCatalogPricingTierPtrOutput
	ToIntegrationRuntimeSsisCatalogPricingTierPtrOutputWithContext(context.Context) IntegrationRuntimeSsisCatalogPricingTierPtrOutput
}

type integrationRuntimeSsisCatalogPricingTierPtr string

func IntegrationRuntimeSsisCatalogPricingTierPtr(v string) IntegrationRuntimeSsisCatalogPricingTierPtrInput {
	return (*integrationRuntimeSsisCatalogPricingTierPtr)(&v)
}

func (*integrationRuntimeSsisCatalogPricingTierPtr) ElementType() reflect.Type {
	return integrationRuntimeSsisCatalogPricingTierPtrType
}

func (in *integrationRuntimeSsisCatalogPricingTierPtr) ToIntegrationRuntimeSsisCatalogPricingTierPtrOutput() IntegrationRuntimeSsisCatalogPricingTierPtrOutput {
	return pulumi.ToOutput(in).(IntegrationRuntimeSsisCatalogPricingTierPtrOutput)
}

func (in *integrationRuntimeSsisCatalogPricingTierPtr) ToIntegrationRuntimeSsisCatalogPricingTierPtrOutputWithContext(ctx context.Context) IntegrationRuntimeSsisCatalogPricingTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IntegrationRuntimeSsisCatalogPricingTierPtrOutput)
}

type IntegrationRuntimeType string

const (
	IntegrationRuntimeTypeManaged    = IntegrationRuntimeType("Managed")
	IntegrationRuntimeTypeSelfHosted = IntegrationRuntimeType("SelfHosted")
)

func (IntegrationRuntimeType) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeType)(nil)).Elem()
}

func (e IntegrationRuntimeType) ToIntegrationRuntimeTypeOutput() IntegrationRuntimeTypeOutput {
	return pulumi.ToOutput(e).(IntegrationRuntimeTypeOutput)
}

func (e IntegrationRuntimeType) ToIntegrationRuntimeTypeOutputWithContext(ctx context.Context) IntegrationRuntimeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IntegrationRuntimeTypeOutput)
}

func (e IntegrationRuntimeType) ToIntegrationRuntimeTypePtrOutput() IntegrationRuntimeTypePtrOutput {
	return e.ToIntegrationRuntimeTypePtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeType) ToIntegrationRuntimeTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeTypePtrOutput {
	return IntegrationRuntimeType(e).ToIntegrationRuntimeTypeOutputWithContext(ctx).ToIntegrationRuntimeTypePtrOutputWithContext(ctx)
}

func (e IntegrationRuntimeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntegrationRuntimeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntegrationRuntimeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntegrationRuntimeTypeOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntimeType)(nil)).Elem()
}

func (o IntegrationRuntimeTypeOutput) ToIntegrationRuntimeTypeOutput() IntegrationRuntimeTypeOutput {
	return o
}

func (o IntegrationRuntimeTypeOutput) ToIntegrationRuntimeTypeOutputWithContext(ctx context.Context) IntegrationRuntimeTypeOutput {
	return o
}

func (o IntegrationRuntimeTypeOutput) ToIntegrationRuntimeTypePtrOutput() IntegrationRuntimeTypePtrOutput {
	return o.ToIntegrationRuntimeTypePtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeTypeOutput) ToIntegrationRuntimeTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationRuntimeType) *IntegrationRuntimeType {
		return &v
	}).(IntegrationRuntimeTypePtrOutput)
}

func (o IntegrationRuntimeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IntegrationRuntimeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IntegrationRuntimeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntegrationRuntimeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IntegrationRuntimeTypePtrOutput struct{ *pulumi.OutputState }

func (IntegrationRuntimeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationRuntimeType)(nil)).Elem()
}

func (o IntegrationRuntimeTypePtrOutput) ToIntegrationRuntimeTypePtrOutput() IntegrationRuntimeTypePtrOutput {
	return o
}

func (o IntegrationRuntimeTypePtrOutput) ToIntegrationRuntimeTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeTypePtrOutput {
	return o
}

func (o IntegrationRuntimeTypePtrOutput) Elem() IntegrationRuntimeTypeOutput {
	return o.ApplyT(func(v *IntegrationRuntimeType) IntegrationRuntimeType {
		if v != nil {
			return *v
		}
		var ret IntegrationRuntimeType
		return ret
	}).(IntegrationRuntimeTypeOutput)
}

func (o IntegrationRuntimeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntegrationRuntimeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IntegrationRuntimeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IntegrationRuntimeTypeInput interface {
	pulumi.Input

	ToIntegrationRuntimeTypeOutput() IntegrationRuntimeTypeOutput
	ToIntegrationRuntimeTypeOutputWithContext(context.Context) IntegrationRuntimeTypeOutput
}

var integrationRuntimeTypePtrType = reflect.TypeOf((**IntegrationRuntimeType)(nil)).Elem()

type IntegrationRuntimeTypePtrInput interface {
	pulumi.Input

	ToIntegrationRuntimeTypePtrOutput() IntegrationRuntimeTypePtrOutput
	ToIntegrationRuntimeTypePtrOutputWithContext(context.Context) IntegrationRuntimeTypePtrOutput
}

type integrationRuntimeTypePtr string

func IntegrationRuntimeTypePtr(v string) IntegrationRuntimeTypePtrInput {
	return (*integrationRuntimeTypePtr)(&v)
}

func (*integrationRuntimeTypePtr) ElementType() reflect.Type {
	return integrationRuntimeTypePtrType
}

func (in *integrationRuntimeTypePtr) ToIntegrationRuntimeTypePtrOutput() IntegrationRuntimeTypePtrOutput {
	return pulumi.ToOutput(in).(IntegrationRuntimeTypePtrOutput)
}

func (in *integrationRuntimeTypePtr) ToIntegrationRuntimeTypePtrOutputWithContext(ctx context.Context) IntegrationRuntimeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IntegrationRuntimeTypePtrOutput)
}

type NodeSize string

const (
	NodeSizeNone     = NodeSize("None")
	NodeSizeSmall    = NodeSize("Small")
	NodeSizeMedium   = NodeSize("Medium")
	NodeSizeLarge    = NodeSize("Large")
	NodeSizeXLarge   = NodeSize("XLarge")
	NodeSizeXXLarge  = NodeSize("XXLarge")
	NodeSizeXXXLarge = NodeSize("XXXLarge")
)

func (NodeSize) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeSize)(nil)).Elem()
}

func (e NodeSize) ToNodeSizeOutput() NodeSizeOutput {
	return pulumi.ToOutput(e).(NodeSizeOutput)
}

func (e NodeSize) ToNodeSizeOutputWithContext(ctx context.Context) NodeSizeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NodeSizeOutput)
}

func (e NodeSize) ToNodeSizePtrOutput() NodeSizePtrOutput {
	return e.ToNodeSizePtrOutputWithContext(context.Background())
}

func (e NodeSize) ToNodeSizePtrOutputWithContext(ctx context.Context) NodeSizePtrOutput {
	return NodeSize(e).ToNodeSizeOutputWithContext(ctx).ToNodeSizePtrOutputWithContext(ctx)
}

func (e NodeSize) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NodeSize) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NodeSize) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NodeSize) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NodeSizeOutput struct{ *pulumi.OutputState }

func (NodeSizeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeSize)(nil)).Elem()
}

func (o NodeSizeOutput) ToNodeSizeOutput() NodeSizeOutput {
	return o
}

func (o NodeSizeOutput) ToNodeSizeOutputWithContext(ctx context.Context) NodeSizeOutput {
	return o
}

func (o NodeSizeOutput) ToNodeSizePtrOutput() NodeSizePtrOutput {
	return o.ToNodeSizePtrOutputWithContext(context.Background())
}

func (o NodeSizeOutput) ToNodeSizePtrOutputWithContext(ctx context.Context) NodeSizePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NodeSize) *NodeSize {
		return &v
	}).(NodeSizePtrOutput)
}

func (o NodeSizeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NodeSizeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NodeSize) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NodeSizeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NodeSizeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NodeSize) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NodeSizePtrOutput struct{ *pulumi.OutputState }

func (NodeSizePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeSize)(nil)).Elem()
}

func (o NodeSizePtrOutput) ToNodeSizePtrOutput() NodeSizePtrOutput {
	return o
}

func (o NodeSizePtrOutput) ToNodeSizePtrOutputWithContext(ctx context.Context) NodeSizePtrOutput {
	return o
}

func (o NodeSizePtrOutput) Elem() NodeSizeOutput {
	return o.ApplyT(func(v *NodeSize) NodeSize {
		if v != nil {
			return *v
		}
		var ret NodeSize
		return ret
	}).(NodeSizeOutput)
}

func (o NodeSizePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NodeSizePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NodeSize) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NodeSizeInput interface {
	pulumi.Input

	ToNodeSizeOutput() NodeSizeOutput
	ToNodeSizeOutputWithContext(context.Context) NodeSizeOutput
}

var nodeSizePtrType = reflect.TypeOf((**NodeSize)(nil)).Elem()

type NodeSizePtrInput interface {
	pulumi.Input

	ToNodeSizePtrOutput() NodeSizePtrOutput
	ToNodeSizePtrOutputWithContext(context.Context) NodeSizePtrOutput
}

type nodeSizePtr string

func NodeSizePtr(v string) NodeSizePtrInput {
	return (*nodeSizePtr)(&v)
}

func (*nodeSizePtr) ElementType() reflect.Type {
	return nodeSizePtrType
}

func (in *nodeSizePtr) ToNodeSizePtrOutput() NodeSizePtrOutput {
	return pulumi.ToOutput(in).(NodeSizePtrOutput)
}

func (in *nodeSizePtr) ToNodeSizePtrOutputWithContext(ctx context.Context) NodeSizePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NodeSizePtrOutput)
}

type NodeSizeFamily string

const (
	NodeSizeFamilyNone            = NodeSizeFamily("None")
	NodeSizeFamilyMemoryOptimized = NodeSizeFamily("MemoryOptimized")
)

func (NodeSizeFamily) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeSizeFamily)(nil)).Elem()
}

func (e NodeSizeFamily) ToNodeSizeFamilyOutput() NodeSizeFamilyOutput {
	return pulumi.ToOutput(e).(NodeSizeFamilyOutput)
}

func (e NodeSizeFamily) ToNodeSizeFamilyOutputWithContext(ctx context.Context) NodeSizeFamilyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NodeSizeFamilyOutput)
}

func (e NodeSizeFamily) ToNodeSizeFamilyPtrOutput() NodeSizeFamilyPtrOutput {
	return e.ToNodeSizeFamilyPtrOutputWithContext(context.Background())
}

func (e NodeSizeFamily) ToNodeSizeFamilyPtrOutputWithContext(ctx context.Context) NodeSizeFamilyPtrOutput {
	return NodeSizeFamily(e).ToNodeSizeFamilyOutputWithContext(ctx).ToNodeSizeFamilyPtrOutputWithContext(ctx)
}

func (e NodeSizeFamily) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NodeSizeFamily) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NodeSizeFamily) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NodeSizeFamily) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NodeSizeFamilyOutput struct{ *pulumi.OutputState }

func (NodeSizeFamilyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeSizeFamily)(nil)).Elem()
}

func (o NodeSizeFamilyOutput) ToNodeSizeFamilyOutput() NodeSizeFamilyOutput {
	return o
}

func (o NodeSizeFamilyOutput) ToNodeSizeFamilyOutputWithContext(ctx context.Context) NodeSizeFamilyOutput {
	return o
}

func (o NodeSizeFamilyOutput) ToNodeSizeFamilyPtrOutput() NodeSizeFamilyPtrOutput {
	return o.ToNodeSizeFamilyPtrOutputWithContext(context.Background())
}

func (o NodeSizeFamilyOutput) ToNodeSizeFamilyPtrOutputWithContext(ctx context.Context) NodeSizeFamilyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NodeSizeFamily) *NodeSizeFamily {
		return &v
	}).(NodeSizeFamilyPtrOutput)
}

func (o NodeSizeFamilyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NodeSizeFamilyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NodeSizeFamily) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NodeSizeFamilyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NodeSizeFamilyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NodeSizeFamily) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NodeSizeFamilyPtrOutput struct{ *pulumi.OutputState }

func (NodeSizeFamilyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeSizeFamily)(nil)).Elem()
}

func (o NodeSizeFamilyPtrOutput) ToNodeSizeFamilyPtrOutput() NodeSizeFamilyPtrOutput {
	return o
}

func (o NodeSizeFamilyPtrOutput) ToNodeSizeFamilyPtrOutputWithContext(ctx context.Context) NodeSizeFamilyPtrOutput {
	return o
}

func (o NodeSizeFamilyPtrOutput) Elem() NodeSizeFamilyOutput {
	return o.ApplyT(func(v *NodeSizeFamily) NodeSizeFamily {
		if v != nil {
			return *v
		}
		var ret NodeSizeFamily
		return ret
	}).(NodeSizeFamilyOutput)
}

func (o NodeSizeFamilyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NodeSizeFamilyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NodeSizeFamily) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NodeSizeFamilyInput interface {
	pulumi.Input

	ToNodeSizeFamilyOutput() NodeSizeFamilyOutput
	ToNodeSizeFamilyOutputWithContext(context.Context) NodeSizeFamilyOutput
}

var nodeSizeFamilyPtrType = reflect.TypeOf((**NodeSizeFamily)(nil)).Elem()

type NodeSizeFamilyPtrInput interface {
	pulumi.Input

	ToNodeSizeFamilyPtrOutput() NodeSizeFamilyPtrOutput
	ToNodeSizeFamilyPtrOutputWithContext(context.Context) NodeSizeFamilyPtrOutput
}

type nodeSizeFamilyPtr string

func NodeSizeFamilyPtr(v string) NodeSizeFamilyPtrInput {
	return (*nodeSizeFamilyPtr)(&v)
}

func (*nodeSizeFamilyPtr) ElementType() reflect.Type {
	return nodeSizeFamilyPtrType
}

func (in *nodeSizeFamilyPtr) ToNodeSizeFamilyPtrOutput() NodeSizeFamilyPtrOutput {
	return pulumi.ToOutput(in).(NodeSizeFamilyPtrOutput)
}

func (in *nodeSizeFamilyPtr) ToNodeSizeFamilyPtrOutputWithContext(ctx context.Context) NodeSizeFamilyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NodeSizeFamilyPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
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

type SensitivityLabelRank string

const (
	SensitivityLabelRankNone     = SensitivityLabelRank("None")
	SensitivityLabelRankLow      = SensitivityLabelRank("Low")
	SensitivityLabelRankMedium   = SensitivityLabelRank("Medium")
	SensitivityLabelRankHigh     = SensitivityLabelRank("High")
	SensitivityLabelRankCritical = SensitivityLabelRank("Critical")
)

func (SensitivityLabelRank) ElementType() reflect.Type {
	return reflect.TypeOf((*SensitivityLabelRank)(nil)).Elem()
}

func (e SensitivityLabelRank) ToSensitivityLabelRankOutput() SensitivityLabelRankOutput {
	return pulumi.ToOutput(e).(SensitivityLabelRankOutput)
}

func (e SensitivityLabelRank) ToSensitivityLabelRankOutputWithContext(ctx context.Context) SensitivityLabelRankOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SensitivityLabelRankOutput)
}

func (e SensitivityLabelRank) ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput {
	return e.ToSensitivityLabelRankPtrOutputWithContext(context.Background())
}

func (e SensitivityLabelRank) ToSensitivityLabelRankPtrOutputWithContext(ctx context.Context) SensitivityLabelRankPtrOutput {
	return SensitivityLabelRank(e).ToSensitivityLabelRankOutputWithContext(ctx).ToSensitivityLabelRankPtrOutputWithContext(ctx)
}

func (e SensitivityLabelRank) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SensitivityLabelRank) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SensitivityLabelRank) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SensitivityLabelRank) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SensitivityLabelRankOutput struct{ *pulumi.OutputState }

func (SensitivityLabelRankOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SensitivityLabelRank)(nil)).Elem()
}

func (o SensitivityLabelRankOutput) ToSensitivityLabelRankOutput() SensitivityLabelRankOutput {
	return o
}

func (o SensitivityLabelRankOutput) ToSensitivityLabelRankOutputWithContext(ctx context.Context) SensitivityLabelRankOutput {
	return o
}

func (o SensitivityLabelRankOutput) ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput {
	return o.ToSensitivityLabelRankPtrOutputWithContext(context.Background())
}

func (o SensitivityLabelRankOutput) ToSensitivityLabelRankPtrOutputWithContext(ctx context.Context) SensitivityLabelRankPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SensitivityLabelRank) *SensitivityLabelRank {
		return &v
	}).(SensitivityLabelRankPtrOutput)
}

func (o SensitivityLabelRankOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SensitivityLabelRankOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SensitivityLabelRank) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SensitivityLabelRankOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SensitivityLabelRankOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SensitivityLabelRank) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SensitivityLabelRankPtrOutput struct{ *pulumi.OutputState }

func (SensitivityLabelRankPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SensitivityLabelRank)(nil)).Elem()
}

func (o SensitivityLabelRankPtrOutput) ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput {
	return o
}

func (o SensitivityLabelRankPtrOutput) ToSensitivityLabelRankPtrOutputWithContext(ctx context.Context) SensitivityLabelRankPtrOutput {
	return o
}

func (o SensitivityLabelRankPtrOutput) Elem() SensitivityLabelRankOutput {
	return o.ApplyT(func(v *SensitivityLabelRank) SensitivityLabelRank {
		if v != nil {
			return *v
		}
		var ret SensitivityLabelRank
		return ret
	}).(SensitivityLabelRankOutput)
}

func (o SensitivityLabelRankPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SensitivityLabelRankPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SensitivityLabelRank) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SensitivityLabelRankInput interface {
	pulumi.Input

	ToSensitivityLabelRankOutput() SensitivityLabelRankOutput
	ToSensitivityLabelRankOutputWithContext(context.Context) SensitivityLabelRankOutput
}

var sensitivityLabelRankPtrType = reflect.TypeOf((**SensitivityLabelRank)(nil)).Elem()

type SensitivityLabelRankPtrInput interface {
	pulumi.Input

	ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput
	ToSensitivityLabelRankPtrOutputWithContext(context.Context) SensitivityLabelRankPtrOutput
}

type sensitivityLabelRankPtr string

func SensitivityLabelRankPtr(v string) SensitivityLabelRankPtrInput {
	return (*sensitivityLabelRankPtr)(&v)
}

func (*sensitivityLabelRankPtr) ElementType() reflect.Type {
	return sensitivityLabelRankPtrType
}

func (in *sensitivityLabelRankPtr) ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput {
	return pulumi.ToOutput(in).(SensitivityLabelRankPtrOutput)
}

func (in *sensitivityLabelRankPtr) ToSensitivityLabelRankPtrOutputWithContext(ctx context.Context) SensitivityLabelRankPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SensitivityLabelRankPtrOutput)
}

type TransparentDataEncryptionStatus string

const (
	TransparentDataEncryptionStatusEnabled  = TransparentDataEncryptionStatus("Enabled")
	TransparentDataEncryptionStatusDisabled = TransparentDataEncryptionStatus("Disabled")
)

func (TransparentDataEncryptionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*TransparentDataEncryptionStatus)(nil)).Elem()
}

func (e TransparentDataEncryptionStatus) ToTransparentDataEncryptionStatusOutput() TransparentDataEncryptionStatusOutput {
	return pulumi.ToOutput(e).(TransparentDataEncryptionStatusOutput)
}

func (e TransparentDataEncryptionStatus) ToTransparentDataEncryptionStatusOutputWithContext(ctx context.Context) TransparentDataEncryptionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TransparentDataEncryptionStatusOutput)
}

func (e TransparentDataEncryptionStatus) ToTransparentDataEncryptionStatusPtrOutput() TransparentDataEncryptionStatusPtrOutput {
	return e.ToTransparentDataEncryptionStatusPtrOutputWithContext(context.Background())
}

func (e TransparentDataEncryptionStatus) ToTransparentDataEncryptionStatusPtrOutputWithContext(ctx context.Context) TransparentDataEncryptionStatusPtrOutput {
	return TransparentDataEncryptionStatus(e).ToTransparentDataEncryptionStatusOutputWithContext(ctx).ToTransparentDataEncryptionStatusPtrOutputWithContext(ctx)
}

func (e TransparentDataEncryptionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransparentDataEncryptionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransparentDataEncryptionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TransparentDataEncryptionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TransparentDataEncryptionStatusOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransparentDataEncryptionStatus)(nil)).Elem()
}

func (o TransparentDataEncryptionStatusOutput) ToTransparentDataEncryptionStatusOutput() TransparentDataEncryptionStatusOutput {
	return o
}

func (o TransparentDataEncryptionStatusOutput) ToTransparentDataEncryptionStatusOutputWithContext(ctx context.Context) TransparentDataEncryptionStatusOutput {
	return o
}

func (o TransparentDataEncryptionStatusOutput) ToTransparentDataEncryptionStatusPtrOutput() TransparentDataEncryptionStatusPtrOutput {
	return o.ToTransparentDataEncryptionStatusPtrOutputWithContext(context.Background())
}

func (o TransparentDataEncryptionStatusOutput) ToTransparentDataEncryptionStatusPtrOutputWithContext(ctx context.Context) TransparentDataEncryptionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransparentDataEncryptionStatus) *TransparentDataEncryptionStatus {
		return &v
	}).(TransparentDataEncryptionStatusPtrOutput)
}

func (o TransparentDataEncryptionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TransparentDataEncryptionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TransparentDataEncryptionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TransparentDataEncryptionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TransparentDataEncryptionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TransparentDataEncryptionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TransparentDataEncryptionStatusPtrOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransparentDataEncryptionStatus)(nil)).Elem()
}

func (o TransparentDataEncryptionStatusPtrOutput) ToTransparentDataEncryptionStatusPtrOutput() TransparentDataEncryptionStatusPtrOutput {
	return o
}

func (o TransparentDataEncryptionStatusPtrOutput) ToTransparentDataEncryptionStatusPtrOutputWithContext(ctx context.Context) TransparentDataEncryptionStatusPtrOutput {
	return o
}

func (o TransparentDataEncryptionStatusPtrOutput) Elem() TransparentDataEncryptionStatusOutput {
	return o.ApplyT(func(v *TransparentDataEncryptionStatus) TransparentDataEncryptionStatus {
		if v != nil {
			return *v
		}
		var ret TransparentDataEncryptionStatus
		return ret
	}).(TransparentDataEncryptionStatusOutput)
}

func (o TransparentDataEncryptionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TransparentDataEncryptionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TransparentDataEncryptionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TransparentDataEncryptionStatusInput interface {
	pulumi.Input

	ToTransparentDataEncryptionStatusOutput() TransparentDataEncryptionStatusOutput
	ToTransparentDataEncryptionStatusOutputWithContext(context.Context) TransparentDataEncryptionStatusOutput
}

var transparentDataEncryptionStatusPtrType = reflect.TypeOf((**TransparentDataEncryptionStatus)(nil)).Elem()

type TransparentDataEncryptionStatusPtrInput interface {
	pulumi.Input

	ToTransparentDataEncryptionStatusPtrOutput() TransparentDataEncryptionStatusPtrOutput
	ToTransparentDataEncryptionStatusPtrOutputWithContext(context.Context) TransparentDataEncryptionStatusPtrOutput
}

type transparentDataEncryptionStatusPtr string

func TransparentDataEncryptionStatusPtr(v string) TransparentDataEncryptionStatusPtrInput {
	return (*transparentDataEncryptionStatusPtr)(&v)
}

func (*transparentDataEncryptionStatusPtr) ElementType() reflect.Type {
	return transparentDataEncryptionStatusPtrType
}

func (in *transparentDataEncryptionStatusPtr) ToTransparentDataEncryptionStatusPtrOutput() TransparentDataEncryptionStatusPtrOutput {
	return pulumi.ToOutput(in).(TransparentDataEncryptionStatusPtrOutput)
}

func (in *transparentDataEncryptionStatusPtr) ToTransparentDataEncryptionStatusPtrOutputWithContext(ctx context.Context) TransparentDataEncryptionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TransparentDataEncryptionStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CreateModeInput)(nil)).Elem(), CreateMode("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*CreateModePtrInput)(nil)).Elem(), CreateMode("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataFlowComputeTypeInput)(nil)).Elem(), DataFlowComputeType("General"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataFlowComputeTypePtrInput)(nil)).Elem(), DataFlowComputeType("General"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeEditionInput)(nil)).Elem(), IntegrationRuntimeEdition("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeEditionPtrInput)(nil)).Elem(), IntegrationRuntimeEdition("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeEntityReferenceTypeInput)(nil)).Elem(), IntegrationRuntimeEntityReferenceType("IntegrationRuntimeReference"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeEntityReferenceTypePtrInput)(nil)).Elem(), IntegrationRuntimeEntityReferenceType("IntegrationRuntimeReference"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeLicenseTypeInput)(nil)).Elem(), IntegrationRuntimeLicenseType("BasePrice"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeLicenseTypePtrInput)(nil)).Elem(), IntegrationRuntimeLicenseType("BasePrice"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeSsisCatalogPricingTierInput)(nil)).Elem(), IntegrationRuntimeSsisCatalogPricingTier("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeSsisCatalogPricingTierPtrInput)(nil)).Elem(), IntegrationRuntimeSsisCatalogPricingTier("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeTypeInput)(nil)).Elem(), IntegrationRuntimeType("Managed"))
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationRuntimeTypePtrInput)(nil)).Elem(), IntegrationRuntimeType("Managed"))
	pulumi.RegisterInputType(reflect.TypeOf((*NodeSizeInput)(nil)).Elem(), NodeSize("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*NodeSizePtrInput)(nil)).Elem(), NodeSize("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*NodeSizeFamilyInput)(nil)).Elem(), NodeSizeFamily("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*NodeSizeFamilyPtrInput)(nil)).Elem(), NodeSizeFamily("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*SensitivityLabelRankInput)(nil)).Elem(), SensitivityLabelRank("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*SensitivityLabelRankPtrInput)(nil)).Elem(), SensitivityLabelRank("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*TransparentDataEncryptionStatusInput)(nil)).Elem(), TransparentDataEncryptionStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TransparentDataEncryptionStatusPtrInput)(nil)).Elem(), TransparentDataEncryptionStatus("Enabled"))
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
	pulumi.RegisterOutputType(DataFlowComputeTypeOutput{})
	pulumi.RegisterOutputType(DataFlowComputeTypePtrOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeEditionOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeEditionPtrOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeEntityReferenceTypeOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeEntityReferenceTypePtrOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeLicenseTypeOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeSsisCatalogPricingTierOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeSsisCatalogPricingTierPtrOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeTypeOutput{})
	pulumi.RegisterOutputType(IntegrationRuntimeTypePtrOutput{})
	pulumi.RegisterOutputType(NodeSizeOutput{})
	pulumi.RegisterOutputType(NodeSizePtrOutput{})
	pulumi.RegisterOutputType(NodeSizeFamilyOutput{})
	pulumi.RegisterOutputType(NodeSizeFamilyPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankPtrOutput{})
	pulumi.RegisterOutputType(TransparentDataEncryptionStatusOutput{})
	pulumi.RegisterOutputType(TransparentDataEncryptionStatusPtrOutput{})
}
