


package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureFunctionActivityMethod string

const (
	AzureFunctionActivityMethodGET     = AzureFunctionActivityMethod("GET")
	AzureFunctionActivityMethodPOST    = AzureFunctionActivityMethod("POST")
	AzureFunctionActivityMethodPUT     = AzureFunctionActivityMethod("PUT")
	AzureFunctionActivityMethodDELETE  = AzureFunctionActivityMethod("DELETE")
	AzureFunctionActivityMethodOPTIONS = AzureFunctionActivityMethod("OPTIONS")
	AzureFunctionActivityMethodHEAD    = AzureFunctionActivityMethod("HEAD")
	AzureFunctionActivityMethodTRACE   = AzureFunctionActivityMethod("TRACE")
)

func (AzureFunctionActivityMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionActivityMethod)(nil)).Elem()
}

func (e AzureFunctionActivityMethod) ToAzureFunctionActivityMethodOutput() AzureFunctionActivityMethodOutput {
	return pulumi.ToOutput(e).(AzureFunctionActivityMethodOutput)
}

func (e AzureFunctionActivityMethod) ToAzureFunctionActivityMethodOutputWithContext(ctx context.Context) AzureFunctionActivityMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureFunctionActivityMethodOutput)
}

func (e AzureFunctionActivityMethod) ToAzureFunctionActivityMethodPtrOutput() AzureFunctionActivityMethodPtrOutput {
	return e.ToAzureFunctionActivityMethodPtrOutputWithContext(context.Background())
}

func (e AzureFunctionActivityMethod) ToAzureFunctionActivityMethodPtrOutputWithContext(ctx context.Context) AzureFunctionActivityMethodPtrOutput {
	return AzureFunctionActivityMethod(e).ToAzureFunctionActivityMethodOutputWithContext(ctx).ToAzureFunctionActivityMethodPtrOutputWithContext(ctx)
}

func (e AzureFunctionActivityMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFunctionActivityMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFunctionActivityMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureFunctionActivityMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureFunctionActivityMethodOutput struct{ *pulumi.OutputState }

func (AzureFunctionActivityMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionActivityMethod)(nil)).Elem()
}

func (o AzureFunctionActivityMethodOutput) ToAzureFunctionActivityMethodOutput() AzureFunctionActivityMethodOutput {
	return o
}

func (o AzureFunctionActivityMethodOutput) ToAzureFunctionActivityMethodOutputWithContext(ctx context.Context) AzureFunctionActivityMethodOutput {
	return o
}

func (o AzureFunctionActivityMethodOutput) ToAzureFunctionActivityMethodPtrOutput() AzureFunctionActivityMethodPtrOutput {
	return o.ToAzureFunctionActivityMethodPtrOutputWithContext(context.Background())
}

func (o AzureFunctionActivityMethodOutput) ToAzureFunctionActivityMethodPtrOutputWithContext(ctx context.Context) AzureFunctionActivityMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFunctionActivityMethod) *AzureFunctionActivityMethod {
		return &v
	}).(AzureFunctionActivityMethodPtrOutput)
}

func (o AzureFunctionActivityMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureFunctionActivityMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFunctionActivityMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureFunctionActivityMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFunctionActivityMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFunctionActivityMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureFunctionActivityMethodPtrOutput struct{ *pulumi.OutputState }

func (AzureFunctionActivityMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFunctionActivityMethod)(nil)).Elem()
}

func (o AzureFunctionActivityMethodPtrOutput) ToAzureFunctionActivityMethodPtrOutput() AzureFunctionActivityMethodPtrOutput {
	return o
}

func (o AzureFunctionActivityMethodPtrOutput) ToAzureFunctionActivityMethodPtrOutputWithContext(ctx context.Context) AzureFunctionActivityMethodPtrOutput {
	return o
}

func (o AzureFunctionActivityMethodPtrOutput) Elem() AzureFunctionActivityMethodOutput {
	return o.ApplyT(func(v *AzureFunctionActivityMethod) AzureFunctionActivityMethod {
		if v != nil {
			return *v
		}
		var ret AzureFunctionActivityMethod
		return ret
	}).(AzureFunctionActivityMethodOutput)
}

func (o AzureFunctionActivityMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFunctionActivityMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureFunctionActivityMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureFunctionActivityMethodInput interface {
	pulumi.Input

	ToAzureFunctionActivityMethodOutput() AzureFunctionActivityMethodOutput
	ToAzureFunctionActivityMethodOutputWithContext(context.Context) AzureFunctionActivityMethodOutput
}

var azureFunctionActivityMethodPtrType = reflect.TypeOf((**AzureFunctionActivityMethod)(nil)).Elem()

type AzureFunctionActivityMethodPtrInput interface {
	pulumi.Input

	ToAzureFunctionActivityMethodPtrOutput() AzureFunctionActivityMethodPtrOutput
	ToAzureFunctionActivityMethodPtrOutputWithContext(context.Context) AzureFunctionActivityMethodPtrOutput
}

type azureFunctionActivityMethodPtr string

func AzureFunctionActivityMethodPtr(v string) AzureFunctionActivityMethodPtrInput {
	return (*azureFunctionActivityMethodPtr)(&v)
}

func (*azureFunctionActivityMethodPtr) ElementType() reflect.Type {
	return azureFunctionActivityMethodPtrType
}

func (in *azureFunctionActivityMethodPtr) ToAzureFunctionActivityMethodPtrOutput() AzureFunctionActivityMethodPtrOutput {
	return pulumi.ToOutput(in).(AzureFunctionActivityMethodPtrOutput)
}

func (in *azureFunctionActivityMethodPtr) ToAzureFunctionActivityMethodPtrOutputWithContext(ctx context.Context) AzureFunctionActivityMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureFunctionActivityMethodPtrOutput)
}

type AzureSearchIndexWriteBehaviorType string

const (
	AzureSearchIndexWriteBehaviorTypeMerge  = AzureSearchIndexWriteBehaviorType("Merge")
	AzureSearchIndexWriteBehaviorTypeUpload = AzureSearchIndexWriteBehaviorType("Upload")
)

func (AzureSearchIndexWriteBehaviorType) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSearchIndexWriteBehaviorType)(nil)).Elem()
}

func (e AzureSearchIndexWriteBehaviorType) ToAzureSearchIndexWriteBehaviorTypeOutput() AzureSearchIndexWriteBehaviorTypeOutput {
	return pulumi.ToOutput(e).(AzureSearchIndexWriteBehaviorTypeOutput)
}

func (e AzureSearchIndexWriteBehaviorType) ToAzureSearchIndexWriteBehaviorTypeOutputWithContext(ctx context.Context) AzureSearchIndexWriteBehaviorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSearchIndexWriteBehaviorTypeOutput)
}

func (e AzureSearchIndexWriteBehaviorType) ToAzureSearchIndexWriteBehaviorTypePtrOutput() AzureSearchIndexWriteBehaviorTypePtrOutput {
	return e.ToAzureSearchIndexWriteBehaviorTypePtrOutputWithContext(context.Background())
}

func (e AzureSearchIndexWriteBehaviorType) ToAzureSearchIndexWriteBehaviorTypePtrOutputWithContext(ctx context.Context) AzureSearchIndexWriteBehaviorTypePtrOutput {
	return AzureSearchIndexWriteBehaviorType(e).ToAzureSearchIndexWriteBehaviorTypeOutputWithContext(ctx).ToAzureSearchIndexWriteBehaviorTypePtrOutputWithContext(ctx)
}

func (e AzureSearchIndexWriteBehaviorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSearchIndexWriteBehaviorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSearchIndexWriteBehaviorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSearchIndexWriteBehaviorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSearchIndexWriteBehaviorTypeOutput struct{ *pulumi.OutputState }

func (AzureSearchIndexWriteBehaviorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSearchIndexWriteBehaviorType)(nil)).Elem()
}

func (o AzureSearchIndexWriteBehaviorTypeOutput) ToAzureSearchIndexWriteBehaviorTypeOutput() AzureSearchIndexWriteBehaviorTypeOutput {
	return o
}

func (o AzureSearchIndexWriteBehaviorTypeOutput) ToAzureSearchIndexWriteBehaviorTypeOutputWithContext(ctx context.Context) AzureSearchIndexWriteBehaviorTypeOutput {
	return o
}

func (o AzureSearchIndexWriteBehaviorTypeOutput) ToAzureSearchIndexWriteBehaviorTypePtrOutput() AzureSearchIndexWriteBehaviorTypePtrOutput {
	return o.ToAzureSearchIndexWriteBehaviorTypePtrOutputWithContext(context.Background())
}

func (o AzureSearchIndexWriteBehaviorTypeOutput) ToAzureSearchIndexWriteBehaviorTypePtrOutputWithContext(ctx context.Context) AzureSearchIndexWriteBehaviorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSearchIndexWriteBehaviorType) *AzureSearchIndexWriteBehaviorType {
		return &v
	}).(AzureSearchIndexWriteBehaviorTypePtrOutput)
}

func (o AzureSearchIndexWriteBehaviorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSearchIndexWriteBehaviorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSearchIndexWriteBehaviorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSearchIndexWriteBehaviorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSearchIndexWriteBehaviorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSearchIndexWriteBehaviorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSearchIndexWriteBehaviorTypePtrOutput struct{ *pulumi.OutputState }

func (AzureSearchIndexWriteBehaviorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSearchIndexWriteBehaviorType)(nil)).Elem()
}

func (o AzureSearchIndexWriteBehaviorTypePtrOutput) ToAzureSearchIndexWriteBehaviorTypePtrOutput() AzureSearchIndexWriteBehaviorTypePtrOutput {
	return o
}

func (o AzureSearchIndexWriteBehaviorTypePtrOutput) ToAzureSearchIndexWriteBehaviorTypePtrOutputWithContext(ctx context.Context) AzureSearchIndexWriteBehaviorTypePtrOutput {
	return o
}

func (o AzureSearchIndexWriteBehaviorTypePtrOutput) Elem() AzureSearchIndexWriteBehaviorTypeOutput {
	return o.ApplyT(func(v *AzureSearchIndexWriteBehaviorType) AzureSearchIndexWriteBehaviorType {
		if v != nil {
			return *v
		}
		var ret AzureSearchIndexWriteBehaviorType
		return ret
	}).(AzureSearchIndexWriteBehaviorTypeOutput)
}

func (o AzureSearchIndexWriteBehaviorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSearchIndexWriteBehaviorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSearchIndexWriteBehaviorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureSearchIndexWriteBehaviorTypeInput interface {
	pulumi.Input

	ToAzureSearchIndexWriteBehaviorTypeOutput() AzureSearchIndexWriteBehaviorTypeOutput
	ToAzureSearchIndexWriteBehaviorTypeOutputWithContext(context.Context) AzureSearchIndexWriteBehaviorTypeOutput
}

var azureSearchIndexWriteBehaviorTypePtrType = reflect.TypeOf((**AzureSearchIndexWriteBehaviorType)(nil)).Elem()

type AzureSearchIndexWriteBehaviorTypePtrInput interface {
	pulumi.Input

	ToAzureSearchIndexWriteBehaviorTypePtrOutput() AzureSearchIndexWriteBehaviorTypePtrOutput
	ToAzureSearchIndexWriteBehaviorTypePtrOutputWithContext(context.Context) AzureSearchIndexWriteBehaviorTypePtrOutput
}

type azureSearchIndexWriteBehaviorTypePtr string

func AzureSearchIndexWriteBehaviorTypePtr(v string) AzureSearchIndexWriteBehaviorTypePtrInput {
	return (*azureSearchIndexWriteBehaviorTypePtr)(&v)
}

func (*azureSearchIndexWriteBehaviorTypePtr) ElementType() reflect.Type {
	return azureSearchIndexWriteBehaviorTypePtrType
}

func (in *azureSearchIndexWriteBehaviorTypePtr) ToAzureSearchIndexWriteBehaviorTypePtrOutput() AzureSearchIndexWriteBehaviorTypePtrOutput {
	return pulumi.ToOutput(in).(AzureSearchIndexWriteBehaviorTypePtrOutput)
}

func (in *azureSearchIndexWriteBehaviorTypePtr) ToAzureSearchIndexWriteBehaviorTypePtrOutputWithContext(ctx context.Context) AzureSearchIndexWriteBehaviorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSearchIndexWriteBehaviorTypePtrOutput)
}

type BlobEventTypes string

const (
	BlobEventTypes_Microsoft_Storage_BlobCreated = BlobEventTypes("Microsoft.Storage.BlobCreated")
	BlobEventTypes_Microsoft_Storage_BlobDeleted = BlobEventTypes("Microsoft.Storage.BlobDeleted")
)

func (BlobEventTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobEventTypes)(nil)).Elem()
}

func (e BlobEventTypes) ToBlobEventTypesOutput() BlobEventTypesOutput {
	return pulumi.ToOutput(e).(BlobEventTypesOutput)
}

func (e BlobEventTypes) ToBlobEventTypesOutputWithContext(ctx context.Context) BlobEventTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlobEventTypesOutput)
}

func (e BlobEventTypes) ToBlobEventTypesPtrOutput() BlobEventTypesPtrOutput {
	return e.ToBlobEventTypesPtrOutputWithContext(context.Background())
}

func (e BlobEventTypes) ToBlobEventTypesPtrOutputWithContext(ctx context.Context) BlobEventTypesPtrOutput {
	return BlobEventTypes(e).ToBlobEventTypesOutputWithContext(ctx).ToBlobEventTypesPtrOutputWithContext(ctx)
}

func (e BlobEventTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobEventTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobEventTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlobEventTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlobEventTypesOutput struct{ *pulumi.OutputState }

func (BlobEventTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobEventTypes)(nil)).Elem()
}

func (o BlobEventTypesOutput) ToBlobEventTypesOutput() BlobEventTypesOutput {
	return o
}

func (o BlobEventTypesOutput) ToBlobEventTypesOutputWithContext(ctx context.Context) BlobEventTypesOutput {
	return o
}

func (o BlobEventTypesOutput) ToBlobEventTypesPtrOutput() BlobEventTypesPtrOutput {
	return o.ToBlobEventTypesPtrOutputWithContext(context.Background())
}

func (o BlobEventTypesOutput) ToBlobEventTypesPtrOutputWithContext(ctx context.Context) BlobEventTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobEventTypes) *BlobEventTypes {
		return &v
	}).(BlobEventTypesPtrOutput)
}

func (o BlobEventTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlobEventTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobEventTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlobEventTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobEventTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobEventTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlobEventTypesPtrOutput struct{ *pulumi.OutputState }

func (BlobEventTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobEventTypes)(nil)).Elem()
}

func (o BlobEventTypesPtrOutput) ToBlobEventTypesPtrOutput() BlobEventTypesPtrOutput {
	return o
}

func (o BlobEventTypesPtrOutput) ToBlobEventTypesPtrOutputWithContext(ctx context.Context) BlobEventTypesPtrOutput {
	return o
}

func (o BlobEventTypesPtrOutput) Elem() BlobEventTypesOutput {
	return o.ApplyT(func(v *BlobEventTypes) BlobEventTypes {
		if v != nil {
			return *v
		}
		var ret BlobEventTypes
		return ret
	}).(BlobEventTypesOutput)
}

func (o BlobEventTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobEventTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlobEventTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BlobEventTypesInput interface {
	pulumi.Input

	ToBlobEventTypesOutput() BlobEventTypesOutput
	ToBlobEventTypesOutputWithContext(context.Context) BlobEventTypesOutput
}

var blobEventTypesPtrType = reflect.TypeOf((**BlobEventTypes)(nil)).Elem()

type BlobEventTypesPtrInput interface {
	pulumi.Input

	ToBlobEventTypesPtrOutput() BlobEventTypesPtrOutput
	ToBlobEventTypesPtrOutputWithContext(context.Context) BlobEventTypesPtrOutput
}

type blobEventTypesPtr string

func BlobEventTypesPtr(v string) BlobEventTypesPtrInput {
	return (*blobEventTypesPtr)(&v)
}

func (*blobEventTypesPtr) ElementType() reflect.Type {
	return blobEventTypesPtrType
}

func (in *blobEventTypesPtr) ToBlobEventTypesPtrOutput() BlobEventTypesPtrOutput {
	return pulumi.ToOutput(in).(BlobEventTypesPtrOutput)
}

func (in *blobEventTypesPtr) ToBlobEventTypesPtrOutputWithContext(ctx context.Context) BlobEventTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlobEventTypesPtrOutput)
}

type CassandraSourceReadConsistencyLevels string

const (
	CassandraSourceReadConsistencyLevelsALL           = CassandraSourceReadConsistencyLevels("ALL")
	CassandraSourceReadConsistencyLevels_EACH_QUORUM  = CassandraSourceReadConsistencyLevels("EACH_QUORUM")
	CassandraSourceReadConsistencyLevelsQUORUM        = CassandraSourceReadConsistencyLevels("QUORUM")
	CassandraSourceReadConsistencyLevels_LOCAL_QUORUM = CassandraSourceReadConsistencyLevels("LOCAL_QUORUM")
	CassandraSourceReadConsistencyLevelsONE           = CassandraSourceReadConsistencyLevels("ONE")
	CassandraSourceReadConsistencyLevelsTWO           = CassandraSourceReadConsistencyLevels("TWO")
	CassandraSourceReadConsistencyLevelsTHREE         = CassandraSourceReadConsistencyLevels("THREE")
	CassandraSourceReadConsistencyLevels_LOCAL_ONE    = CassandraSourceReadConsistencyLevels("LOCAL_ONE")
	CassandraSourceReadConsistencyLevelsSERIAL        = CassandraSourceReadConsistencyLevels("SERIAL")
	CassandraSourceReadConsistencyLevels_LOCAL_SERIAL = CassandraSourceReadConsistencyLevels("LOCAL_SERIAL")
)

func (CassandraSourceReadConsistencyLevels) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSourceReadConsistencyLevels)(nil)).Elem()
}

func (e CassandraSourceReadConsistencyLevels) ToCassandraSourceReadConsistencyLevelsOutput() CassandraSourceReadConsistencyLevelsOutput {
	return pulumi.ToOutput(e).(CassandraSourceReadConsistencyLevelsOutput)
}

func (e CassandraSourceReadConsistencyLevels) ToCassandraSourceReadConsistencyLevelsOutputWithContext(ctx context.Context) CassandraSourceReadConsistencyLevelsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CassandraSourceReadConsistencyLevelsOutput)
}

func (e CassandraSourceReadConsistencyLevels) ToCassandraSourceReadConsistencyLevelsPtrOutput() CassandraSourceReadConsistencyLevelsPtrOutput {
	return e.ToCassandraSourceReadConsistencyLevelsPtrOutputWithContext(context.Background())
}

func (e CassandraSourceReadConsistencyLevels) ToCassandraSourceReadConsistencyLevelsPtrOutputWithContext(ctx context.Context) CassandraSourceReadConsistencyLevelsPtrOutput {
	return CassandraSourceReadConsistencyLevels(e).ToCassandraSourceReadConsistencyLevelsOutputWithContext(ctx).ToCassandraSourceReadConsistencyLevelsPtrOutputWithContext(ctx)
}

func (e CassandraSourceReadConsistencyLevels) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CassandraSourceReadConsistencyLevels) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CassandraSourceReadConsistencyLevels) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CassandraSourceReadConsistencyLevels) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CassandraSourceReadConsistencyLevelsOutput struct{ *pulumi.OutputState }

func (CassandraSourceReadConsistencyLevelsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSourceReadConsistencyLevels)(nil)).Elem()
}

func (o CassandraSourceReadConsistencyLevelsOutput) ToCassandraSourceReadConsistencyLevelsOutput() CassandraSourceReadConsistencyLevelsOutput {
	return o
}

func (o CassandraSourceReadConsistencyLevelsOutput) ToCassandraSourceReadConsistencyLevelsOutputWithContext(ctx context.Context) CassandraSourceReadConsistencyLevelsOutput {
	return o
}

func (o CassandraSourceReadConsistencyLevelsOutput) ToCassandraSourceReadConsistencyLevelsPtrOutput() CassandraSourceReadConsistencyLevelsPtrOutput {
	return o.ToCassandraSourceReadConsistencyLevelsPtrOutputWithContext(context.Background())
}

func (o CassandraSourceReadConsistencyLevelsOutput) ToCassandraSourceReadConsistencyLevelsPtrOutputWithContext(ctx context.Context) CassandraSourceReadConsistencyLevelsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraSourceReadConsistencyLevels) *CassandraSourceReadConsistencyLevels {
		return &v
	}).(CassandraSourceReadConsistencyLevelsPtrOutput)
}

func (o CassandraSourceReadConsistencyLevelsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CassandraSourceReadConsistencyLevelsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CassandraSourceReadConsistencyLevels) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CassandraSourceReadConsistencyLevelsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CassandraSourceReadConsistencyLevelsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CassandraSourceReadConsistencyLevels) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CassandraSourceReadConsistencyLevelsPtrOutput struct{ *pulumi.OutputState }

func (CassandraSourceReadConsistencyLevelsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSourceReadConsistencyLevels)(nil)).Elem()
}

func (o CassandraSourceReadConsistencyLevelsPtrOutput) ToCassandraSourceReadConsistencyLevelsPtrOutput() CassandraSourceReadConsistencyLevelsPtrOutput {
	return o
}

func (o CassandraSourceReadConsistencyLevelsPtrOutput) ToCassandraSourceReadConsistencyLevelsPtrOutputWithContext(ctx context.Context) CassandraSourceReadConsistencyLevelsPtrOutput {
	return o
}

func (o CassandraSourceReadConsistencyLevelsPtrOutput) Elem() CassandraSourceReadConsistencyLevelsOutput {
	return o.ApplyT(func(v *CassandraSourceReadConsistencyLevels) CassandraSourceReadConsistencyLevels {
		if v != nil {
			return *v
		}
		var ret CassandraSourceReadConsistencyLevels
		return ret
	}).(CassandraSourceReadConsistencyLevelsOutput)
}

func (o CassandraSourceReadConsistencyLevelsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CassandraSourceReadConsistencyLevelsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CassandraSourceReadConsistencyLevels) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CassandraSourceReadConsistencyLevelsInput interface {
	pulumi.Input

	ToCassandraSourceReadConsistencyLevelsOutput() CassandraSourceReadConsistencyLevelsOutput
	ToCassandraSourceReadConsistencyLevelsOutputWithContext(context.Context) CassandraSourceReadConsistencyLevelsOutput
}

var cassandraSourceReadConsistencyLevelsPtrType = reflect.TypeOf((**CassandraSourceReadConsistencyLevels)(nil)).Elem()

type CassandraSourceReadConsistencyLevelsPtrInput interface {
	pulumi.Input

	ToCassandraSourceReadConsistencyLevelsPtrOutput() CassandraSourceReadConsistencyLevelsPtrOutput
	ToCassandraSourceReadConsistencyLevelsPtrOutputWithContext(context.Context) CassandraSourceReadConsistencyLevelsPtrOutput
}

type cassandraSourceReadConsistencyLevelsPtr string

func CassandraSourceReadConsistencyLevelsPtr(v string) CassandraSourceReadConsistencyLevelsPtrInput {
	return (*cassandraSourceReadConsistencyLevelsPtr)(&v)
}

func (*cassandraSourceReadConsistencyLevelsPtr) ElementType() reflect.Type {
	return cassandraSourceReadConsistencyLevelsPtrType
}

func (in *cassandraSourceReadConsistencyLevelsPtr) ToCassandraSourceReadConsistencyLevelsPtrOutput() CassandraSourceReadConsistencyLevelsPtrOutput {
	return pulumi.ToOutput(in).(CassandraSourceReadConsistencyLevelsPtrOutput)
}

func (in *cassandraSourceReadConsistencyLevelsPtr) ToCassandraSourceReadConsistencyLevelsPtrOutputWithContext(ctx context.Context) CassandraSourceReadConsistencyLevelsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CassandraSourceReadConsistencyLevelsPtrOutput)
}

type CosmosDbConnectionMode string

const (
	CosmosDbConnectionModeGateway = CosmosDbConnectionMode("Gateway")
	CosmosDbConnectionModeDirect  = CosmosDbConnectionMode("Direct")
)

func (CosmosDbConnectionMode) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbConnectionMode)(nil)).Elem()
}

func (e CosmosDbConnectionMode) ToCosmosDbConnectionModeOutput() CosmosDbConnectionModeOutput {
	return pulumi.ToOutput(e).(CosmosDbConnectionModeOutput)
}

func (e CosmosDbConnectionMode) ToCosmosDbConnectionModeOutputWithContext(ctx context.Context) CosmosDbConnectionModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CosmosDbConnectionModeOutput)
}

func (e CosmosDbConnectionMode) ToCosmosDbConnectionModePtrOutput() CosmosDbConnectionModePtrOutput {
	return e.ToCosmosDbConnectionModePtrOutputWithContext(context.Background())
}

func (e CosmosDbConnectionMode) ToCosmosDbConnectionModePtrOutputWithContext(ctx context.Context) CosmosDbConnectionModePtrOutput {
	return CosmosDbConnectionMode(e).ToCosmosDbConnectionModeOutputWithContext(ctx).ToCosmosDbConnectionModePtrOutputWithContext(ctx)
}

func (e CosmosDbConnectionMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CosmosDbConnectionMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CosmosDbConnectionMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CosmosDbConnectionMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CosmosDbConnectionModeOutput struct{ *pulumi.OutputState }

func (CosmosDbConnectionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbConnectionMode)(nil)).Elem()
}

func (o CosmosDbConnectionModeOutput) ToCosmosDbConnectionModeOutput() CosmosDbConnectionModeOutput {
	return o
}

func (o CosmosDbConnectionModeOutput) ToCosmosDbConnectionModeOutputWithContext(ctx context.Context) CosmosDbConnectionModeOutput {
	return o
}

func (o CosmosDbConnectionModeOutput) ToCosmosDbConnectionModePtrOutput() CosmosDbConnectionModePtrOutput {
	return o.ToCosmosDbConnectionModePtrOutputWithContext(context.Background())
}

func (o CosmosDbConnectionModeOutput) ToCosmosDbConnectionModePtrOutputWithContext(ctx context.Context) CosmosDbConnectionModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CosmosDbConnectionMode) *CosmosDbConnectionMode {
		return &v
	}).(CosmosDbConnectionModePtrOutput)
}

func (o CosmosDbConnectionModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CosmosDbConnectionModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CosmosDbConnectionMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CosmosDbConnectionModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CosmosDbConnectionModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CosmosDbConnectionMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CosmosDbConnectionModePtrOutput struct{ *pulumi.OutputState }

func (CosmosDbConnectionModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbConnectionMode)(nil)).Elem()
}

func (o CosmosDbConnectionModePtrOutput) ToCosmosDbConnectionModePtrOutput() CosmosDbConnectionModePtrOutput {
	return o
}

func (o CosmosDbConnectionModePtrOutput) ToCosmosDbConnectionModePtrOutputWithContext(ctx context.Context) CosmosDbConnectionModePtrOutput {
	return o
}

func (o CosmosDbConnectionModePtrOutput) Elem() CosmosDbConnectionModeOutput {
	return o.ApplyT(func(v *CosmosDbConnectionMode) CosmosDbConnectionMode {
		if v != nil {
			return *v
		}
		var ret CosmosDbConnectionMode
		return ret
	}).(CosmosDbConnectionModeOutput)
}

func (o CosmosDbConnectionModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CosmosDbConnectionModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CosmosDbConnectionMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CosmosDbConnectionModeInput interface {
	pulumi.Input

	ToCosmosDbConnectionModeOutput() CosmosDbConnectionModeOutput
	ToCosmosDbConnectionModeOutputWithContext(context.Context) CosmosDbConnectionModeOutput
}

var cosmosDbConnectionModePtrType = reflect.TypeOf((**CosmosDbConnectionMode)(nil)).Elem()

type CosmosDbConnectionModePtrInput interface {
	pulumi.Input

	ToCosmosDbConnectionModePtrOutput() CosmosDbConnectionModePtrOutput
	ToCosmosDbConnectionModePtrOutputWithContext(context.Context) CosmosDbConnectionModePtrOutput
}

type cosmosDbConnectionModePtr string

func CosmosDbConnectionModePtr(v string) CosmosDbConnectionModePtrInput {
	return (*cosmosDbConnectionModePtr)(&v)
}

func (*cosmosDbConnectionModePtr) ElementType() reflect.Type {
	return cosmosDbConnectionModePtrType
}

func (in *cosmosDbConnectionModePtr) ToCosmosDbConnectionModePtrOutput() CosmosDbConnectionModePtrOutput {
	return pulumi.ToOutput(in).(CosmosDbConnectionModePtrOutput)
}

func (in *cosmosDbConnectionModePtr) ToCosmosDbConnectionModePtrOutputWithContext(ctx context.Context) CosmosDbConnectionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CosmosDbConnectionModePtrOutput)
}

type CosmosDbServicePrincipalCredentialType string

const (
	CosmosDbServicePrincipalCredentialTypeServicePrincipalKey  = CosmosDbServicePrincipalCredentialType("ServicePrincipalKey")
	CosmosDbServicePrincipalCredentialTypeServicePrincipalCert = CosmosDbServicePrincipalCredentialType("ServicePrincipalCert")
)

func (CosmosDbServicePrincipalCredentialType) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbServicePrincipalCredentialType)(nil)).Elem()
}

func (e CosmosDbServicePrincipalCredentialType) ToCosmosDbServicePrincipalCredentialTypeOutput() CosmosDbServicePrincipalCredentialTypeOutput {
	return pulumi.ToOutput(e).(CosmosDbServicePrincipalCredentialTypeOutput)
}

func (e CosmosDbServicePrincipalCredentialType) ToCosmosDbServicePrincipalCredentialTypeOutputWithContext(ctx context.Context) CosmosDbServicePrincipalCredentialTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CosmosDbServicePrincipalCredentialTypeOutput)
}

func (e CosmosDbServicePrincipalCredentialType) ToCosmosDbServicePrincipalCredentialTypePtrOutput() CosmosDbServicePrincipalCredentialTypePtrOutput {
	return e.ToCosmosDbServicePrincipalCredentialTypePtrOutputWithContext(context.Background())
}

func (e CosmosDbServicePrincipalCredentialType) ToCosmosDbServicePrincipalCredentialTypePtrOutputWithContext(ctx context.Context) CosmosDbServicePrincipalCredentialTypePtrOutput {
	return CosmosDbServicePrincipalCredentialType(e).ToCosmosDbServicePrincipalCredentialTypeOutputWithContext(ctx).ToCosmosDbServicePrincipalCredentialTypePtrOutputWithContext(ctx)
}

func (e CosmosDbServicePrincipalCredentialType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CosmosDbServicePrincipalCredentialType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CosmosDbServicePrincipalCredentialType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CosmosDbServicePrincipalCredentialType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CosmosDbServicePrincipalCredentialTypeOutput struct{ *pulumi.OutputState }

func (CosmosDbServicePrincipalCredentialTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CosmosDbServicePrincipalCredentialType)(nil)).Elem()
}

func (o CosmosDbServicePrincipalCredentialTypeOutput) ToCosmosDbServicePrincipalCredentialTypeOutput() CosmosDbServicePrincipalCredentialTypeOutput {
	return o
}

func (o CosmosDbServicePrincipalCredentialTypeOutput) ToCosmosDbServicePrincipalCredentialTypeOutputWithContext(ctx context.Context) CosmosDbServicePrincipalCredentialTypeOutput {
	return o
}

func (o CosmosDbServicePrincipalCredentialTypeOutput) ToCosmosDbServicePrincipalCredentialTypePtrOutput() CosmosDbServicePrincipalCredentialTypePtrOutput {
	return o.ToCosmosDbServicePrincipalCredentialTypePtrOutputWithContext(context.Background())
}

func (o CosmosDbServicePrincipalCredentialTypeOutput) ToCosmosDbServicePrincipalCredentialTypePtrOutputWithContext(ctx context.Context) CosmosDbServicePrincipalCredentialTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CosmosDbServicePrincipalCredentialType) *CosmosDbServicePrincipalCredentialType {
		return &v
	}).(CosmosDbServicePrincipalCredentialTypePtrOutput)
}

func (o CosmosDbServicePrincipalCredentialTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CosmosDbServicePrincipalCredentialTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CosmosDbServicePrincipalCredentialType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CosmosDbServicePrincipalCredentialTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CosmosDbServicePrincipalCredentialTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CosmosDbServicePrincipalCredentialType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CosmosDbServicePrincipalCredentialTypePtrOutput struct{ *pulumi.OutputState }

func (CosmosDbServicePrincipalCredentialTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosmosDbServicePrincipalCredentialType)(nil)).Elem()
}

func (o CosmosDbServicePrincipalCredentialTypePtrOutput) ToCosmosDbServicePrincipalCredentialTypePtrOutput() CosmosDbServicePrincipalCredentialTypePtrOutput {
	return o
}

func (o CosmosDbServicePrincipalCredentialTypePtrOutput) ToCosmosDbServicePrincipalCredentialTypePtrOutputWithContext(ctx context.Context) CosmosDbServicePrincipalCredentialTypePtrOutput {
	return o
}

func (o CosmosDbServicePrincipalCredentialTypePtrOutput) Elem() CosmosDbServicePrincipalCredentialTypeOutput {
	return o.ApplyT(func(v *CosmosDbServicePrincipalCredentialType) CosmosDbServicePrincipalCredentialType {
		if v != nil {
			return *v
		}
		var ret CosmosDbServicePrincipalCredentialType
		return ret
	}).(CosmosDbServicePrincipalCredentialTypeOutput)
}

func (o CosmosDbServicePrincipalCredentialTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CosmosDbServicePrincipalCredentialTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CosmosDbServicePrincipalCredentialType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CosmosDbServicePrincipalCredentialTypeInput interface {
	pulumi.Input

	ToCosmosDbServicePrincipalCredentialTypeOutput() CosmosDbServicePrincipalCredentialTypeOutput
	ToCosmosDbServicePrincipalCredentialTypeOutputWithContext(context.Context) CosmosDbServicePrincipalCredentialTypeOutput
}

var cosmosDbServicePrincipalCredentialTypePtrType = reflect.TypeOf((**CosmosDbServicePrincipalCredentialType)(nil)).Elem()

type CosmosDbServicePrincipalCredentialTypePtrInput interface {
	pulumi.Input

	ToCosmosDbServicePrincipalCredentialTypePtrOutput() CosmosDbServicePrincipalCredentialTypePtrOutput
	ToCosmosDbServicePrincipalCredentialTypePtrOutputWithContext(context.Context) CosmosDbServicePrincipalCredentialTypePtrOutput
}

type cosmosDbServicePrincipalCredentialTypePtr string

func CosmosDbServicePrincipalCredentialTypePtr(v string) CosmosDbServicePrincipalCredentialTypePtrInput {
	return (*cosmosDbServicePrincipalCredentialTypePtr)(&v)
}

func (*cosmosDbServicePrincipalCredentialTypePtr) ElementType() reflect.Type {
	return cosmosDbServicePrincipalCredentialTypePtrType
}

func (in *cosmosDbServicePrincipalCredentialTypePtr) ToCosmosDbServicePrincipalCredentialTypePtrOutput() CosmosDbServicePrincipalCredentialTypePtrOutput {
	return pulumi.ToOutput(in).(CosmosDbServicePrincipalCredentialTypePtrOutput)
}

func (in *cosmosDbServicePrincipalCredentialTypePtr) ToCosmosDbServicePrincipalCredentialTypePtrOutputWithContext(ctx context.Context) CosmosDbServicePrincipalCredentialTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CosmosDbServicePrincipalCredentialTypePtrOutput)
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

type DaysOfWeek string

const (
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
)

func (DaysOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*DaysOfWeek)(nil)).Elem()
}

func (e DaysOfWeek) ToDaysOfWeekOutput() DaysOfWeekOutput {
	return pulumi.ToOutput(e).(DaysOfWeekOutput)
}

func (e DaysOfWeek) ToDaysOfWeekOutputWithContext(ctx context.Context) DaysOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DaysOfWeekOutput)
}

func (e DaysOfWeek) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return e.ToDaysOfWeekPtrOutputWithContext(context.Background())
}

func (e DaysOfWeek) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return DaysOfWeek(e).ToDaysOfWeekOutputWithContext(ctx).ToDaysOfWeekPtrOutputWithContext(ctx)
}

func (e DaysOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DaysOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DaysOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DaysOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DaysOfWeekOutput struct{ *pulumi.OutputState }

func (DaysOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaysOfWeek)(nil)).Elem()
}

func (o DaysOfWeekOutput) ToDaysOfWeekOutput() DaysOfWeekOutput {
	return o
}

func (o DaysOfWeekOutput) ToDaysOfWeekOutputWithContext(ctx context.Context) DaysOfWeekOutput {
	return o
}

func (o DaysOfWeekOutput) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return o.ToDaysOfWeekPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DaysOfWeek) *DaysOfWeek {
		return &v
	}).(DaysOfWeekPtrOutput)
}

func (o DaysOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DaysOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DaysOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DaysOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DaysOfWeekPtrOutput struct{ *pulumi.OutputState }

func (DaysOfWeekPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaysOfWeek)(nil)).Elem()
}

func (o DaysOfWeekPtrOutput) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return o
}

func (o DaysOfWeekPtrOutput) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return o
}

func (o DaysOfWeekPtrOutput) Elem() DaysOfWeekOutput {
	return o.ApplyT(func(v *DaysOfWeek) DaysOfWeek {
		if v != nil {
			return *v
		}
		var ret DaysOfWeek
		return ret
	}).(DaysOfWeekOutput)
}

func (o DaysOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DaysOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DaysOfWeekInput interface {
	pulumi.Input

	ToDaysOfWeekOutput() DaysOfWeekOutput
	ToDaysOfWeekOutputWithContext(context.Context) DaysOfWeekOutput
}

var daysOfWeekPtrType = reflect.TypeOf((**DaysOfWeek)(nil)).Elem()

type DaysOfWeekPtrInput interface {
	pulumi.Input

	ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput
	ToDaysOfWeekPtrOutputWithContext(context.Context) DaysOfWeekPtrOutput
}

type daysOfWeekPtr string

func DaysOfWeekPtr(v string) DaysOfWeekPtrInput {
	return (*daysOfWeekPtr)(&v)
}

func (*daysOfWeekPtr) ElementType() reflect.Type {
	return daysOfWeekPtrType
}

func (in *daysOfWeekPtr) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return pulumi.ToOutput(in).(DaysOfWeekPtrOutput)
}

func (in *daysOfWeekPtr) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DaysOfWeekPtrOutput)
}





type DaysOfWeekArrayInput interface {
	pulumi.Input

	ToDaysOfWeekArrayOutput() DaysOfWeekArrayOutput
	ToDaysOfWeekArrayOutputWithContext(context.Context) DaysOfWeekArrayOutput
}

type DaysOfWeekArray []DaysOfWeek

func (DaysOfWeekArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaysOfWeek)(nil)).Elem()
}

func (i DaysOfWeekArray) ToDaysOfWeekArrayOutput() DaysOfWeekArrayOutput {
	return i.ToDaysOfWeekArrayOutputWithContext(context.Background())
}

func (i DaysOfWeekArray) ToDaysOfWeekArrayOutputWithContext(ctx context.Context) DaysOfWeekArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaysOfWeekArrayOutput)
}

type DaysOfWeekArrayOutput struct{ *pulumi.OutputState }

func (DaysOfWeekArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaysOfWeek)(nil)).Elem()
}

func (o DaysOfWeekArrayOutput) ToDaysOfWeekArrayOutput() DaysOfWeekArrayOutput {
	return o
}

func (o DaysOfWeekArrayOutput) ToDaysOfWeekArrayOutputWithContext(ctx context.Context) DaysOfWeekArrayOutput {
	return o
}

func (o DaysOfWeekArrayOutput) Index(i pulumi.IntInput) DaysOfWeekOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DaysOfWeek {
		return vs[0].([]DaysOfWeek)[vs[1].(int)]
	}).(DaysOfWeekOutput)
}

type Db2AuthenticationType string

const (
	Db2AuthenticationTypeBasic = Db2AuthenticationType("Basic")
)

func (Db2AuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*Db2AuthenticationType)(nil)).Elem()
}

func (e Db2AuthenticationType) ToDb2AuthenticationTypeOutput() Db2AuthenticationTypeOutput {
	return pulumi.ToOutput(e).(Db2AuthenticationTypeOutput)
}

func (e Db2AuthenticationType) ToDb2AuthenticationTypeOutputWithContext(ctx context.Context) Db2AuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(Db2AuthenticationTypeOutput)
}

func (e Db2AuthenticationType) ToDb2AuthenticationTypePtrOutput() Db2AuthenticationTypePtrOutput {
	return e.ToDb2AuthenticationTypePtrOutputWithContext(context.Background())
}

func (e Db2AuthenticationType) ToDb2AuthenticationTypePtrOutputWithContext(ctx context.Context) Db2AuthenticationTypePtrOutput {
	return Db2AuthenticationType(e).ToDb2AuthenticationTypeOutputWithContext(ctx).ToDb2AuthenticationTypePtrOutputWithContext(ctx)
}

func (e Db2AuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Db2AuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Db2AuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Db2AuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type Db2AuthenticationTypeOutput struct{ *pulumi.OutputState }

func (Db2AuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Db2AuthenticationType)(nil)).Elem()
}

func (o Db2AuthenticationTypeOutput) ToDb2AuthenticationTypeOutput() Db2AuthenticationTypeOutput {
	return o
}

func (o Db2AuthenticationTypeOutput) ToDb2AuthenticationTypeOutputWithContext(ctx context.Context) Db2AuthenticationTypeOutput {
	return o
}

func (o Db2AuthenticationTypeOutput) ToDb2AuthenticationTypePtrOutput() Db2AuthenticationTypePtrOutput {
	return o.ToDb2AuthenticationTypePtrOutputWithContext(context.Background())
}

func (o Db2AuthenticationTypeOutput) ToDb2AuthenticationTypePtrOutputWithContext(ctx context.Context) Db2AuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Db2AuthenticationType) *Db2AuthenticationType {
		return &v
	}).(Db2AuthenticationTypePtrOutput)
}

func (o Db2AuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o Db2AuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Db2AuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o Db2AuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o Db2AuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Db2AuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type Db2AuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (Db2AuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Db2AuthenticationType)(nil)).Elem()
}

func (o Db2AuthenticationTypePtrOutput) ToDb2AuthenticationTypePtrOutput() Db2AuthenticationTypePtrOutput {
	return o
}

func (o Db2AuthenticationTypePtrOutput) ToDb2AuthenticationTypePtrOutputWithContext(ctx context.Context) Db2AuthenticationTypePtrOutput {
	return o
}

func (o Db2AuthenticationTypePtrOutput) Elem() Db2AuthenticationTypeOutput {
	return o.ApplyT(func(v *Db2AuthenticationType) Db2AuthenticationType {
		if v != nil {
			return *v
		}
		var ret Db2AuthenticationType
		return ret
	}).(Db2AuthenticationTypeOutput)
}

func (o Db2AuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o Db2AuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Db2AuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type Db2AuthenticationTypeInput interface {
	pulumi.Input

	ToDb2AuthenticationTypeOutput() Db2AuthenticationTypeOutput
	ToDb2AuthenticationTypeOutputWithContext(context.Context) Db2AuthenticationTypeOutput
}

var db2AuthenticationTypePtrType = reflect.TypeOf((**Db2AuthenticationType)(nil)).Elem()

type Db2AuthenticationTypePtrInput interface {
	pulumi.Input

	ToDb2AuthenticationTypePtrOutput() Db2AuthenticationTypePtrOutput
	ToDb2AuthenticationTypePtrOutputWithContext(context.Context) Db2AuthenticationTypePtrOutput
}

type db2AuthenticationTypePtr string

func Db2AuthenticationTypePtr(v string) Db2AuthenticationTypePtrInput {
	return (*db2AuthenticationTypePtr)(&v)
}

func (*db2AuthenticationTypePtr) ElementType() reflect.Type {
	return db2AuthenticationTypePtrType
}

func (in *db2AuthenticationTypePtr) ToDb2AuthenticationTypePtrOutput() Db2AuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(Db2AuthenticationTypePtrOutput)
}

func (in *db2AuthenticationTypePtr) ToDb2AuthenticationTypePtrOutputWithContext(ctx context.Context) Db2AuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(Db2AuthenticationTypePtrOutput)
}

type DependencyCondition string

const (
	DependencyConditionSucceeded = DependencyCondition("Succeeded")
	DependencyConditionFailed    = DependencyCondition("Failed")
	DependencyConditionSkipped   = DependencyCondition("Skipped")
	DependencyConditionCompleted = DependencyCondition("Completed")
)

func (DependencyCondition) ElementType() reflect.Type {
	return reflect.TypeOf((*DependencyCondition)(nil)).Elem()
}

func (e DependencyCondition) ToDependencyConditionOutput() DependencyConditionOutput {
	return pulumi.ToOutput(e).(DependencyConditionOutput)
}

func (e DependencyCondition) ToDependencyConditionOutputWithContext(ctx context.Context) DependencyConditionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DependencyConditionOutput)
}

func (e DependencyCondition) ToDependencyConditionPtrOutput() DependencyConditionPtrOutput {
	return e.ToDependencyConditionPtrOutputWithContext(context.Background())
}

func (e DependencyCondition) ToDependencyConditionPtrOutputWithContext(ctx context.Context) DependencyConditionPtrOutput {
	return DependencyCondition(e).ToDependencyConditionOutputWithContext(ctx).ToDependencyConditionPtrOutputWithContext(ctx)
}

func (e DependencyCondition) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DependencyCondition) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DependencyCondition) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DependencyCondition) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DependencyConditionOutput struct{ *pulumi.OutputState }

func (DependencyConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DependencyCondition)(nil)).Elem()
}

func (o DependencyConditionOutput) ToDependencyConditionOutput() DependencyConditionOutput {
	return o
}

func (o DependencyConditionOutput) ToDependencyConditionOutputWithContext(ctx context.Context) DependencyConditionOutput {
	return o
}

func (o DependencyConditionOutput) ToDependencyConditionPtrOutput() DependencyConditionPtrOutput {
	return o.ToDependencyConditionPtrOutputWithContext(context.Background())
}

func (o DependencyConditionOutput) ToDependencyConditionPtrOutputWithContext(ctx context.Context) DependencyConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DependencyCondition) *DependencyCondition {
		return &v
	}).(DependencyConditionPtrOutput)
}

func (o DependencyConditionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DependencyConditionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DependencyCondition) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DependencyConditionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DependencyConditionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DependencyCondition) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DependencyConditionPtrOutput struct{ *pulumi.OutputState }

func (DependencyConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DependencyCondition)(nil)).Elem()
}

func (o DependencyConditionPtrOutput) ToDependencyConditionPtrOutput() DependencyConditionPtrOutput {
	return o
}

func (o DependencyConditionPtrOutput) ToDependencyConditionPtrOutputWithContext(ctx context.Context) DependencyConditionPtrOutput {
	return o
}

func (o DependencyConditionPtrOutput) Elem() DependencyConditionOutput {
	return o.ApplyT(func(v *DependencyCondition) DependencyCondition {
		if v != nil {
			return *v
		}
		var ret DependencyCondition
		return ret
	}).(DependencyConditionOutput)
}

func (o DependencyConditionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DependencyConditionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DependencyCondition) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DependencyConditionInput interface {
	pulumi.Input

	ToDependencyConditionOutput() DependencyConditionOutput
	ToDependencyConditionOutputWithContext(context.Context) DependencyConditionOutput
}

var dependencyConditionPtrType = reflect.TypeOf((**DependencyCondition)(nil)).Elem()

type DependencyConditionPtrInput interface {
	pulumi.Input

	ToDependencyConditionPtrOutput() DependencyConditionPtrOutput
	ToDependencyConditionPtrOutputWithContext(context.Context) DependencyConditionPtrOutput
}

type dependencyConditionPtr string

func DependencyConditionPtr(v string) DependencyConditionPtrInput {
	return (*dependencyConditionPtr)(&v)
}

func (*dependencyConditionPtr) ElementType() reflect.Type {
	return dependencyConditionPtrType
}

func (in *dependencyConditionPtr) ToDependencyConditionPtrOutput() DependencyConditionPtrOutput {
	return pulumi.ToOutput(in).(DependencyConditionPtrOutput)
}

func (in *dependencyConditionPtr) ToDependencyConditionPtrOutputWithContext(ctx context.Context) DependencyConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DependencyConditionPtrOutput)
}

type DynamicsSinkWriteBehavior string

const (
	DynamicsSinkWriteBehaviorUpsert = DynamicsSinkWriteBehavior("Upsert")
)

func (DynamicsSinkWriteBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicsSinkWriteBehavior)(nil)).Elem()
}

func (e DynamicsSinkWriteBehavior) ToDynamicsSinkWriteBehaviorOutput() DynamicsSinkWriteBehaviorOutput {
	return pulumi.ToOutput(e).(DynamicsSinkWriteBehaviorOutput)
}

func (e DynamicsSinkWriteBehavior) ToDynamicsSinkWriteBehaviorOutputWithContext(ctx context.Context) DynamicsSinkWriteBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DynamicsSinkWriteBehaviorOutput)
}

func (e DynamicsSinkWriteBehavior) ToDynamicsSinkWriteBehaviorPtrOutput() DynamicsSinkWriteBehaviorPtrOutput {
	return e.ToDynamicsSinkWriteBehaviorPtrOutputWithContext(context.Background())
}

func (e DynamicsSinkWriteBehavior) ToDynamicsSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) DynamicsSinkWriteBehaviorPtrOutput {
	return DynamicsSinkWriteBehavior(e).ToDynamicsSinkWriteBehaviorOutputWithContext(ctx).ToDynamicsSinkWriteBehaviorPtrOutputWithContext(ctx)
}

func (e DynamicsSinkWriteBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DynamicsSinkWriteBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DynamicsSinkWriteBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DynamicsSinkWriteBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DynamicsSinkWriteBehaviorOutput struct{ *pulumi.OutputState }

func (DynamicsSinkWriteBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicsSinkWriteBehavior)(nil)).Elem()
}

func (o DynamicsSinkWriteBehaviorOutput) ToDynamicsSinkWriteBehaviorOutput() DynamicsSinkWriteBehaviorOutput {
	return o
}

func (o DynamicsSinkWriteBehaviorOutput) ToDynamicsSinkWriteBehaviorOutputWithContext(ctx context.Context) DynamicsSinkWriteBehaviorOutput {
	return o
}

func (o DynamicsSinkWriteBehaviorOutput) ToDynamicsSinkWriteBehaviorPtrOutput() DynamicsSinkWriteBehaviorPtrOutput {
	return o.ToDynamicsSinkWriteBehaviorPtrOutputWithContext(context.Background())
}

func (o DynamicsSinkWriteBehaviorOutput) ToDynamicsSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) DynamicsSinkWriteBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DynamicsSinkWriteBehavior) *DynamicsSinkWriteBehavior {
		return &v
	}).(DynamicsSinkWriteBehaviorPtrOutput)
}

func (o DynamicsSinkWriteBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DynamicsSinkWriteBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DynamicsSinkWriteBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DynamicsSinkWriteBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DynamicsSinkWriteBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DynamicsSinkWriteBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DynamicsSinkWriteBehaviorPtrOutput struct{ *pulumi.OutputState }

func (DynamicsSinkWriteBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamicsSinkWriteBehavior)(nil)).Elem()
}

func (o DynamicsSinkWriteBehaviorPtrOutput) ToDynamicsSinkWriteBehaviorPtrOutput() DynamicsSinkWriteBehaviorPtrOutput {
	return o
}

func (o DynamicsSinkWriteBehaviorPtrOutput) ToDynamicsSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) DynamicsSinkWriteBehaviorPtrOutput {
	return o
}

func (o DynamicsSinkWriteBehaviorPtrOutput) Elem() DynamicsSinkWriteBehaviorOutput {
	return o.ApplyT(func(v *DynamicsSinkWriteBehavior) DynamicsSinkWriteBehavior {
		if v != nil {
			return *v
		}
		var ret DynamicsSinkWriteBehavior
		return ret
	}).(DynamicsSinkWriteBehaviorOutput)
}

func (o DynamicsSinkWriteBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DynamicsSinkWriteBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DynamicsSinkWriteBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DynamicsSinkWriteBehaviorInput interface {
	pulumi.Input

	ToDynamicsSinkWriteBehaviorOutput() DynamicsSinkWriteBehaviorOutput
	ToDynamicsSinkWriteBehaviorOutputWithContext(context.Context) DynamicsSinkWriteBehaviorOutput
}

var dynamicsSinkWriteBehaviorPtrType = reflect.TypeOf((**DynamicsSinkWriteBehavior)(nil)).Elem()

type DynamicsSinkWriteBehaviorPtrInput interface {
	pulumi.Input

	ToDynamicsSinkWriteBehaviorPtrOutput() DynamicsSinkWriteBehaviorPtrOutput
	ToDynamicsSinkWriteBehaviorPtrOutputWithContext(context.Context) DynamicsSinkWriteBehaviorPtrOutput
}

type dynamicsSinkWriteBehaviorPtr string

func DynamicsSinkWriteBehaviorPtr(v string) DynamicsSinkWriteBehaviorPtrInput {
	return (*dynamicsSinkWriteBehaviorPtr)(&v)
}

func (*dynamicsSinkWriteBehaviorPtr) ElementType() reflect.Type {
	return dynamicsSinkWriteBehaviorPtrType
}

func (in *dynamicsSinkWriteBehaviorPtr) ToDynamicsSinkWriteBehaviorPtrOutput() DynamicsSinkWriteBehaviorPtrOutput {
	return pulumi.ToOutput(in).(DynamicsSinkWriteBehaviorPtrOutput)
}

func (in *dynamicsSinkWriteBehaviorPtr) ToDynamicsSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) DynamicsSinkWriteBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DynamicsSinkWriteBehaviorPtrOutput)
}

type FactoryIdentityType string

const (
	FactoryIdentityTypeSystemAssigned               = FactoryIdentityType("SystemAssigned")
	FactoryIdentityTypeUserAssigned                 = FactoryIdentityType("UserAssigned")
	FactoryIdentityType_SystemAssigned_UserAssigned = FactoryIdentityType("SystemAssigned,UserAssigned")
)

func (FactoryIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*FactoryIdentityType)(nil)).Elem()
}

func (e FactoryIdentityType) ToFactoryIdentityTypeOutput() FactoryIdentityTypeOutput {
	return pulumi.ToOutput(e).(FactoryIdentityTypeOutput)
}

func (e FactoryIdentityType) ToFactoryIdentityTypeOutputWithContext(ctx context.Context) FactoryIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FactoryIdentityTypeOutput)
}

func (e FactoryIdentityType) ToFactoryIdentityTypePtrOutput() FactoryIdentityTypePtrOutput {
	return e.ToFactoryIdentityTypePtrOutputWithContext(context.Background())
}

func (e FactoryIdentityType) ToFactoryIdentityTypePtrOutputWithContext(ctx context.Context) FactoryIdentityTypePtrOutput {
	return FactoryIdentityType(e).ToFactoryIdentityTypeOutputWithContext(ctx).ToFactoryIdentityTypePtrOutputWithContext(ctx)
}

func (e FactoryIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FactoryIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FactoryIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FactoryIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FactoryIdentityTypeOutput struct{ *pulumi.OutputState }

func (FactoryIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FactoryIdentityType)(nil)).Elem()
}

func (o FactoryIdentityTypeOutput) ToFactoryIdentityTypeOutput() FactoryIdentityTypeOutput {
	return o
}

func (o FactoryIdentityTypeOutput) ToFactoryIdentityTypeOutputWithContext(ctx context.Context) FactoryIdentityTypeOutput {
	return o
}

func (o FactoryIdentityTypeOutput) ToFactoryIdentityTypePtrOutput() FactoryIdentityTypePtrOutput {
	return o.ToFactoryIdentityTypePtrOutputWithContext(context.Background())
}

func (o FactoryIdentityTypeOutput) ToFactoryIdentityTypePtrOutputWithContext(ctx context.Context) FactoryIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FactoryIdentityType) *FactoryIdentityType {
		return &v
	}).(FactoryIdentityTypePtrOutput)
}

func (o FactoryIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FactoryIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FactoryIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FactoryIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FactoryIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FactoryIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FactoryIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (FactoryIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FactoryIdentityType)(nil)).Elem()
}

func (o FactoryIdentityTypePtrOutput) ToFactoryIdentityTypePtrOutput() FactoryIdentityTypePtrOutput {
	return o
}

func (o FactoryIdentityTypePtrOutput) ToFactoryIdentityTypePtrOutputWithContext(ctx context.Context) FactoryIdentityTypePtrOutput {
	return o
}

func (o FactoryIdentityTypePtrOutput) Elem() FactoryIdentityTypeOutput {
	return o.ApplyT(func(v *FactoryIdentityType) FactoryIdentityType {
		if v != nil {
			return *v
		}
		var ret FactoryIdentityType
		return ret
	}).(FactoryIdentityTypeOutput)
}

func (o FactoryIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FactoryIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FactoryIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FactoryIdentityTypeInput interface {
	pulumi.Input

	ToFactoryIdentityTypeOutput() FactoryIdentityTypeOutput
	ToFactoryIdentityTypeOutputWithContext(context.Context) FactoryIdentityTypeOutput
}

var factoryIdentityTypePtrType = reflect.TypeOf((**FactoryIdentityType)(nil)).Elem()

type FactoryIdentityTypePtrInput interface {
	pulumi.Input

	ToFactoryIdentityTypePtrOutput() FactoryIdentityTypePtrOutput
	ToFactoryIdentityTypePtrOutputWithContext(context.Context) FactoryIdentityTypePtrOutput
}

type factoryIdentityTypePtr string

func FactoryIdentityTypePtr(v string) FactoryIdentityTypePtrInput {
	return (*factoryIdentityTypePtr)(&v)
}

func (*factoryIdentityTypePtr) ElementType() reflect.Type {
	return factoryIdentityTypePtrType
}

func (in *factoryIdentityTypePtr) ToFactoryIdentityTypePtrOutput() FactoryIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(FactoryIdentityTypePtrOutput)
}

func (in *factoryIdentityTypePtr) ToFactoryIdentityTypePtrOutputWithContext(ctx context.Context) FactoryIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FactoryIdentityTypePtrOutput)
}

type FtpAuthenticationType string

const (
	FtpAuthenticationTypeBasic     = FtpAuthenticationType("Basic")
	FtpAuthenticationTypeAnonymous = FtpAuthenticationType("Anonymous")
)

func (FtpAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*FtpAuthenticationType)(nil)).Elem()
}

func (e FtpAuthenticationType) ToFtpAuthenticationTypeOutput() FtpAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(FtpAuthenticationTypeOutput)
}

func (e FtpAuthenticationType) ToFtpAuthenticationTypeOutputWithContext(ctx context.Context) FtpAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FtpAuthenticationTypeOutput)
}

func (e FtpAuthenticationType) ToFtpAuthenticationTypePtrOutput() FtpAuthenticationTypePtrOutput {
	return e.ToFtpAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e FtpAuthenticationType) ToFtpAuthenticationTypePtrOutputWithContext(ctx context.Context) FtpAuthenticationTypePtrOutput {
	return FtpAuthenticationType(e).ToFtpAuthenticationTypeOutputWithContext(ctx).ToFtpAuthenticationTypePtrOutputWithContext(ctx)
}

func (e FtpAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FtpAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FtpAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FtpAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FtpAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (FtpAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FtpAuthenticationType)(nil)).Elem()
}

func (o FtpAuthenticationTypeOutput) ToFtpAuthenticationTypeOutput() FtpAuthenticationTypeOutput {
	return o
}

func (o FtpAuthenticationTypeOutput) ToFtpAuthenticationTypeOutputWithContext(ctx context.Context) FtpAuthenticationTypeOutput {
	return o
}

func (o FtpAuthenticationTypeOutput) ToFtpAuthenticationTypePtrOutput() FtpAuthenticationTypePtrOutput {
	return o.ToFtpAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o FtpAuthenticationTypeOutput) ToFtpAuthenticationTypePtrOutputWithContext(ctx context.Context) FtpAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FtpAuthenticationType) *FtpAuthenticationType {
		return &v
	}).(FtpAuthenticationTypePtrOutput)
}

func (o FtpAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FtpAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FtpAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FtpAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FtpAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FtpAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FtpAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (FtpAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FtpAuthenticationType)(nil)).Elem()
}

func (o FtpAuthenticationTypePtrOutput) ToFtpAuthenticationTypePtrOutput() FtpAuthenticationTypePtrOutput {
	return o
}

func (o FtpAuthenticationTypePtrOutput) ToFtpAuthenticationTypePtrOutputWithContext(ctx context.Context) FtpAuthenticationTypePtrOutput {
	return o
}

func (o FtpAuthenticationTypePtrOutput) Elem() FtpAuthenticationTypeOutput {
	return o.ApplyT(func(v *FtpAuthenticationType) FtpAuthenticationType {
		if v != nil {
			return *v
		}
		var ret FtpAuthenticationType
		return ret
	}).(FtpAuthenticationTypeOutput)
}

func (o FtpAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FtpAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FtpAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FtpAuthenticationTypeInput interface {
	pulumi.Input

	ToFtpAuthenticationTypeOutput() FtpAuthenticationTypeOutput
	ToFtpAuthenticationTypeOutputWithContext(context.Context) FtpAuthenticationTypeOutput
}

var ftpAuthenticationTypePtrType = reflect.TypeOf((**FtpAuthenticationType)(nil)).Elem()

type FtpAuthenticationTypePtrInput interface {
	pulumi.Input

	ToFtpAuthenticationTypePtrOutput() FtpAuthenticationTypePtrOutput
	ToFtpAuthenticationTypePtrOutputWithContext(context.Context) FtpAuthenticationTypePtrOutput
}

type ftpAuthenticationTypePtr string

func FtpAuthenticationTypePtr(v string) FtpAuthenticationTypePtrInput {
	return (*ftpAuthenticationTypePtr)(&v)
}

func (*ftpAuthenticationTypePtr) ElementType() reflect.Type {
	return ftpAuthenticationTypePtrType
}

func (in *ftpAuthenticationTypePtr) ToFtpAuthenticationTypePtrOutput() FtpAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(FtpAuthenticationTypePtrOutput)
}

func (in *ftpAuthenticationTypePtr) ToFtpAuthenticationTypePtrOutputWithContext(ctx context.Context) FtpAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FtpAuthenticationTypePtrOutput)
}

type GlobalParameterType string

const (
	GlobalParameterTypeObject = GlobalParameterType("Object")
	GlobalParameterTypeString = GlobalParameterType("String")
	GlobalParameterTypeInt    = GlobalParameterType("Int")
	GlobalParameterTypeFloat  = GlobalParameterType("Float")
	GlobalParameterTypeBool   = GlobalParameterType("Bool")
	GlobalParameterTypeArray  = GlobalParameterType("Array")
)

func (GlobalParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalParameterType)(nil)).Elem()
}

func (e GlobalParameterType) ToGlobalParameterTypeOutput() GlobalParameterTypeOutput {
	return pulumi.ToOutput(e).(GlobalParameterTypeOutput)
}

func (e GlobalParameterType) ToGlobalParameterTypeOutputWithContext(ctx context.Context) GlobalParameterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GlobalParameterTypeOutput)
}

func (e GlobalParameterType) ToGlobalParameterTypePtrOutput() GlobalParameterTypePtrOutput {
	return e.ToGlobalParameterTypePtrOutputWithContext(context.Background())
}

func (e GlobalParameterType) ToGlobalParameterTypePtrOutputWithContext(ctx context.Context) GlobalParameterTypePtrOutput {
	return GlobalParameterType(e).ToGlobalParameterTypeOutputWithContext(ctx).ToGlobalParameterTypePtrOutputWithContext(ctx)
}

func (e GlobalParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GlobalParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GlobalParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GlobalParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GlobalParameterTypeOutput struct{ *pulumi.OutputState }

func (GlobalParameterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalParameterType)(nil)).Elem()
}

func (o GlobalParameterTypeOutput) ToGlobalParameterTypeOutput() GlobalParameterTypeOutput {
	return o
}

func (o GlobalParameterTypeOutput) ToGlobalParameterTypeOutputWithContext(ctx context.Context) GlobalParameterTypeOutput {
	return o
}

func (o GlobalParameterTypeOutput) ToGlobalParameterTypePtrOutput() GlobalParameterTypePtrOutput {
	return o.ToGlobalParameterTypePtrOutputWithContext(context.Background())
}

func (o GlobalParameterTypeOutput) ToGlobalParameterTypePtrOutputWithContext(ctx context.Context) GlobalParameterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GlobalParameterType) *GlobalParameterType {
		return &v
	}).(GlobalParameterTypePtrOutput)
}

func (o GlobalParameterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GlobalParameterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GlobalParameterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GlobalParameterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GlobalParameterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GlobalParameterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GlobalParameterTypePtrOutput struct{ *pulumi.OutputState }

func (GlobalParameterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalParameterType)(nil)).Elem()
}

func (o GlobalParameterTypePtrOutput) ToGlobalParameterTypePtrOutput() GlobalParameterTypePtrOutput {
	return o
}

func (o GlobalParameterTypePtrOutput) ToGlobalParameterTypePtrOutputWithContext(ctx context.Context) GlobalParameterTypePtrOutput {
	return o
}

func (o GlobalParameterTypePtrOutput) Elem() GlobalParameterTypeOutput {
	return o.ApplyT(func(v *GlobalParameterType) GlobalParameterType {
		if v != nil {
			return *v
		}
		var ret GlobalParameterType
		return ret
	}).(GlobalParameterTypeOutput)
}

func (o GlobalParameterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GlobalParameterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GlobalParameterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GlobalParameterTypeInput interface {
	pulumi.Input

	ToGlobalParameterTypeOutput() GlobalParameterTypeOutput
	ToGlobalParameterTypeOutputWithContext(context.Context) GlobalParameterTypeOutput
}

var globalParameterTypePtrType = reflect.TypeOf((**GlobalParameterType)(nil)).Elem()

type GlobalParameterTypePtrInput interface {
	pulumi.Input

	ToGlobalParameterTypePtrOutput() GlobalParameterTypePtrOutput
	ToGlobalParameterTypePtrOutputWithContext(context.Context) GlobalParameterTypePtrOutput
}

type globalParameterTypePtr string

func GlobalParameterTypePtr(v string) GlobalParameterTypePtrInput {
	return (*globalParameterTypePtr)(&v)
}

func (*globalParameterTypePtr) ElementType() reflect.Type {
	return globalParameterTypePtrType
}

func (in *globalParameterTypePtr) ToGlobalParameterTypePtrOutput() GlobalParameterTypePtrOutput {
	return pulumi.ToOutput(in).(GlobalParameterTypePtrOutput)
}

func (in *globalParameterTypePtr) ToGlobalParameterTypePtrOutputWithContext(ctx context.Context) GlobalParameterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GlobalParameterTypePtrOutput)
}

type GoogleAdWordsAuthenticationType string

const (
	GoogleAdWordsAuthenticationTypeServiceAuthentication = GoogleAdWordsAuthenticationType("ServiceAuthentication")
	GoogleAdWordsAuthenticationTypeUserAuthentication    = GoogleAdWordsAuthenticationType("UserAuthentication")
)

func (GoogleAdWordsAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleAdWordsAuthenticationType)(nil)).Elem()
}

func (e GoogleAdWordsAuthenticationType) ToGoogleAdWordsAuthenticationTypeOutput() GoogleAdWordsAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(GoogleAdWordsAuthenticationTypeOutput)
}

func (e GoogleAdWordsAuthenticationType) ToGoogleAdWordsAuthenticationTypeOutputWithContext(ctx context.Context) GoogleAdWordsAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GoogleAdWordsAuthenticationTypeOutput)
}

func (e GoogleAdWordsAuthenticationType) ToGoogleAdWordsAuthenticationTypePtrOutput() GoogleAdWordsAuthenticationTypePtrOutput {
	return e.ToGoogleAdWordsAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e GoogleAdWordsAuthenticationType) ToGoogleAdWordsAuthenticationTypePtrOutputWithContext(ctx context.Context) GoogleAdWordsAuthenticationTypePtrOutput {
	return GoogleAdWordsAuthenticationType(e).ToGoogleAdWordsAuthenticationTypeOutputWithContext(ctx).ToGoogleAdWordsAuthenticationTypePtrOutputWithContext(ctx)
}

func (e GoogleAdWordsAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleAdWordsAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleAdWordsAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleAdWordsAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GoogleAdWordsAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (GoogleAdWordsAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleAdWordsAuthenticationType)(nil)).Elem()
}

func (o GoogleAdWordsAuthenticationTypeOutput) ToGoogleAdWordsAuthenticationTypeOutput() GoogleAdWordsAuthenticationTypeOutput {
	return o
}

func (o GoogleAdWordsAuthenticationTypeOutput) ToGoogleAdWordsAuthenticationTypeOutputWithContext(ctx context.Context) GoogleAdWordsAuthenticationTypeOutput {
	return o
}

func (o GoogleAdWordsAuthenticationTypeOutput) ToGoogleAdWordsAuthenticationTypePtrOutput() GoogleAdWordsAuthenticationTypePtrOutput {
	return o.ToGoogleAdWordsAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o GoogleAdWordsAuthenticationTypeOutput) ToGoogleAdWordsAuthenticationTypePtrOutputWithContext(ctx context.Context) GoogleAdWordsAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GoogleAdWordsAuthenticationType) *GoogleAdWordsAuthenticationType {
		return &v
	}).(GoogleAdWordsAuthenticationTypePtrOutput)
}

func (o GoogleAdWordsAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GoogleAdWordsAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleAdWordsAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GoogleAdWordsAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleAdWordsAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleAdWordsAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GoogleAdWordsAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (GoogleAdWordsAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleAdWordsAuthenticationType)(nil)).Elem()
}

func (o GoogleAdWordsAuthenticationTypePtrOutput) ToGoogleAdWordsAuthenticationTypePtrOutput() GoogleAdWordsAuthenticationTypePtrOutput {
	return o
}

func (o GoogleAdWordsAuthenticationTypePtrOutput) ToGoogleAdWordsAuthenticationTypePtrOutputWithContext(ctx context.Context) GoogleAdWordsAuthenticationTypePtrOutput {
	return o
}

func (o GoogleAdWordsAuthenticationTypePtrOutput) Elem() GoogleAdWordsAuthenticationTypeOutput {
	return o.ApplyT(func(v *GoogleAdWordsAuthenticationType) GoogleAdWordsAuthenticationType {
		if v != nil {
			return *v
		}
		var ret GoogleAdWordsAuthenticationType
		return ret
	}).(GoogleAdWordsAuthenticationTypeOutput)
}

func (o GoogleAdWordsAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleAdWordsAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GoogleAdWordsAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GoogleAdWordsAuthenticationTypeInput interface {
	pulumi.Input

	ToGoogleAdWordsAuthenticationTypeOutput() GoogleAdWordsAuthenticationTypeOutput
	ToGoogleAdWordsAuthenticationTypeOutputWithContext(context.Context) GoogleAdWordsAuthenticationTypeOutput
}

var googleAdWordsAuthenticationTypePtrType = reflect.TypeOf((**GoogleAdWordsAuthenticationType)(nil)).Elem()

type GoogleAdWordsAuthenticationTypePtrInput interface {
	pulumi.Input

	ToGoogleAdWordsAuthenticationTypePtrOutput() GoogleAdWordsAuthenticationTypePtrOutput
	ToGoogleAdWordsAuthenticationTypePtrOutputWithContext(context.Context) GoogleAdWordsAuthenticationTypePtrOutput
}

type googleAdWordsAuthenticationTypePtr string

func GoogleAdWordsAuthenticationTypePtr(v string) GoogleAdWordsAuthenticationTypePtrInput {
	return (*googleAdWordsAuthenticationTypePtr)(&v)
}

func (*googleAdWordsAuthenticationTypePtr) ElementType() reflect.Type {
	return googleAdWordsAuthenticationTypePtrType
}

func (in *googleAdWordsAuthenticationTypePtr) ToGoogleAdWordsAuthenticationTypePtrOutput() GoogleAdWordsAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(GoogleAdWordsAuthenticationTypePtrOutput)
}

func (in *googleAdWordsAuthenticationTypePtr) ToGoogleAdWordsAuthenticationTypePtrOutputWithContext(ctx context.Context) GoogleAdWordsAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GoogleAdWordsAuthenticationTypePtrOutput)
}

type GoogleBigQueryAuthenticationType string

const (
	GoogleBigQueryAuthenticationTypeServiceAuthentication = GoogleBigQueryAuthenticationType("ServiceAuthentication")
	GoogleBigQueryAuthenticationTypeUserAuthentication    = GoogleBigQueryAuthenticationType("UserAuthentication")
)

func (GoogleBigQueryAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleBigQueryAuthenticationType)(nil)).Elem()
}

func (e GoogleBigQueryAuthenticationType) ToGoogleBigQueryAuthenticationTypeOutput() GoogleBigQueryAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(GoogleBigQueryAuthenticationTypeOutput)
}

func (e GoogleBigQueryAuthenticationType) ToGoogleBigQueryAuthenticationTypeOutputWithContext(ctx context.Context) GoogleBigQueryAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GoogleBigQueryAuthenticationTypeOutput)
}

func (e GoogleBigQueryAuthenticationType) ToGoogleBigQueryAuthenticationTypePtrOutput() GoogleBigQueryAuthenticationTypePtrOutput {
	return e.ToGoogleBigQueryAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e GoogleBigQueryAuthenticationType) ToGoogleBigQueryAuthenticationTypePtrOutputWithContext(ctx context.Context) GoogleBigQueryAuthenticationTypePtrOutput {
	return GoogleBigQueryAuthenticationType(e).ToGoogleBigQueryAuthenticationTypeOutputWithContext(ctx).ToGoogleBigQueryAuthenticationTypePtrOutputWithContext(ctx)
}

func (e GoogleBigQueryAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleBigQueryAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleBigQueryAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleBigQueryAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GoogleBigQueryAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (GoogleBigQueryAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleBigQueryAuthenticationType)(nil)).Elem()
}

func (o GoogleBigQueryAuthenticationTypeOutput) ToGoogleBigQueryAuthenticationTypeOutput() GoogleBigQueryAuthenticationTypeOutput {
	return o
}

func (o GoogleBigQueryAuthenticationTypeOutput) ToGoogleBigQueryAuthenticationTypeOutputWithContext(ctx context.Context) GoogleBigQueryAuthenticationTypeOutput {
	return o
}

func (o GoogleBigQueryAuthenticationTypeOutput) ToGoogleBigQueryAuthenticationTypePtrOutput() GoogleBigQueryAuthenticationTypePtrOutput {
	return o.ToGoogleBigQueryAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o GoogleBigQueryAuthenticationTypeOutput) ToGoogleBigQueryAuthenticationTypePtrOutputWithContext(ctx context.Context) GoogleBigQueryAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GoogleBigQueryAuthenticationType) *GoogleBigQueryAuthenticationType {
		return &v
	}).(GoogleBigQueryAuthenticationTypePtrOutput)
}

func (o GoogleBigQueryAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GoogleBigQueryAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleBigQueryAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GoogleBigQueryAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleBigQueryAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleBigQueryAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GoogleBigQueryAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (GoogleBigQueryAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleBigQueryAuthenticationType)(nil)).Elem()
}

func (o GoogleBigQueryAuthenticationTypePtrOutput) ToGoogleBigQueryAuthenticationTypePtrOutput() GoogleBigQueryAuthenticationTypePtrOutput {
	return o
}

func (o GoogleBigQueryAuthenticationTypePtrOutput) ToGoogleBigQueryAuthenticationTypePtrOutputWithContext(ctx context.Context) GoogleBigQueryAuthenticationTypePtrOutput {
	return o
}

func (o GoogleBigQueryAuthenticationTypePtrOutput) Elem() GoogleBigQueryAuthenticationTypeOutput {
	return o.ApplyT(func(v *GoogleBigQueryAuthenticationType) GoogleBigQueryAuthenticationType {
		if v != nil {
			return *v
		}
		var ret GoogleBigQueryAuthenticationType
		return ret
	}).(GoogleBigQueryAuthenticationTypeOutput)
}

func (o GoogleBigQueryAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleBigQueryAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GoogleBigQueryAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GoogleBigQueryAuthenticationTypeInput interface {
	pulumi.Input

	ToGoogleBigQueryAuthenticationTypeOutput() GoogleBigQueryAuthenticationTypeOutput
	ToGoogleBigQueryAuthenticationTypeOutputWithContext(context.Context) GoogleBigQueryAuthenticationTypeOutput
}

var googleBigQueryAuthenticationTypePtrType = reflect.TypeOf((**GoogleBigQueryAuthenticationType)(nil)).Elem()

type GoogleBigQueryAuthenticationTypePtrInput interface {
	pulumi.Input

	ToGoogleBigQueryAuthenticationTypePtrOutput() GoogleBigQueryAuthenticationTypePtrOutput
	ToGoogleBigQueryAuthenticationTypePtrOutputWithContext(context.Context) GoogleBigQueryAuthenticationTypePtrOutput
}

type googleBigQueryAuthenticationTypePtr string

func GoogleBigQueryAuthenticationTypePtr(v string) GoogleBigQueryAuthenticationTypePtrInput {
	return (*googleBigQueryAuthenticationTypePtr)(&v)
}

func (*googleBigQueryAuthenticationTypePtr) ElementType() reflect.Type {
	return googleBigQueryAuthenticationTypePtrType
}

func (in *googleBigQueryAuthenticationTypePtr) ToGoogleBigQueryAuthenticationTypePtrOutput() GoogleBigQueryAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(GoogleBigQueryAuthenticationTypePtrOutput)
}

func (in *googleBigQueryAuthenticationTypePtr) ToGoogleBigQueryAuthenticationTypePtrOutputWithContext(ctx context.Context) GoogleBigQueryAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GoogleBigQueryAuthenticationTypePtrOutput)
}

type HBaseAuthenticationType string

const (
	HBaseAuthenticationTypeAnonymous = HBaseAuthenticationType("Anonymous")
	HBaseAuthenticationTypeBasic     = HBaseAuthenticationType("Basic")
)

func (HBaseAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*HBaseAuthenticationType)(nil)).Elem()
}

func (e HBaseAuthenticationType) ToHBaseAuthenticationTypeOutput() HBaseAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(HBaseAuthenticationTypeOutput)
}

func (e HBaseAuthenticationType) ToHBaseAuthenticationTypeOutputWithContext(ctx context.Context) HBaseAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HBaseAuthenticationTypeOutput)
}

func (e HBaseAuthenticationType) ToHBaseAuthenticationTypePtrOutput() HBaseAuthenticationTypePtrOutput {
	return e.ToHBaseAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e HBaseAuthenticationType) ToHBaseAuthenticationTypePtrOutputWithContext(ctx context.Context) HBaseAuthenticationTypePtrOutput {
	return HBaseAuthenticationType(e).ToHBaseAuthenticationTypeOutputWithContext(ctx).ToHBaseAuthenticationTypePtrOutputWithContext(ctx)
}

func (e HBaseAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HBaseAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HBaseAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HBaseAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HBaseAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (HBaseAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HBaseAuthenticationType)(nil)).Elem()
}

func (o HBaseAuthenticationTypeOutput) ToHBaseAuthenticationTypeOutput() HBaseAuthenticationTypeOutput {
	return o
}

func (o HBaseAuthenticationTypeOutput) ToHBaseAuthenticationTypeOutputWithContext(ctx context.Context) HBaseAuthenticationTypeOutput {
	return o
}

func (o HBaseAuthenticationTypeOutput) ToHBaseAuthenticationTypePtrOutput() HBaseAuthenticationTypePtrOutput {
	return o.ToHBaseAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o HBaseAuthenticationTypeOutput) ToHBaseAuthenticationTypePtrOutputWithContext(ctx context.Context) HBaseAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HBaseAuthenticationType) *HBaseAuthenticationType {
		return &v
	}).(HBaseAuthenticationTypePtrOutput)
}

func (o HBaseAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HBaseAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HBaseAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HBaseAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HBaseAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HBaseAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HBaseAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (HBaseAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HBaseAuthenticationType)(nil)).Elem()
}

func (o HBaseAuthenticationTypePtrOutput) ToHBaseAuthenticationTypePtrOutput() HBaseAuthenticationTypePtrOutput {
	return o
}

func (o HBaseAuthenticationTypePtrOutput) ToHBaseAuthenticationTypePtrOutputWithContext(ctx context.Context) HBaseAuthenticationTypePtrOutput {
	return o
}

func (o HBaseAuthenticationTypePtrOutput) Elem() HBaseAuthenticationTypeOutput {
	return o.ApplyT(func(v *HBaseAuthenticationType) HBaseAuthenticationType {
		if v != nil {
			return *v
		}
		var ret HBaseAuthenticationType
		return ret
	}).(HBaseAuthenticationTypeOutput)
}

func (o HBaseAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HBaseAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HBaseAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HBaseAuthenticationTypeInput interface {
	pulumi.Input

	ToHBaseAuthenticationTypeOutput() HBaseAuthenticationTypeOutput
	ToHBaseAuthenticationTypeOutputWithContext(context.Context) HBaseAuthenticationTypeOutput
}

var hbaseAuthenticationTypePtrType = reflect.TypeOf((**HBaseAuthenticationType)(nil)).Elem()

type HBaseAuthenticationTypePtrInput interface {
	pulumi.Input

	ToHBaseAuthenticationTypePtrOutput() HBaseAuthenticationTypePtrOutput
	ToHBaseAuthenticationTypePtrOutputWithContext(context.Context) HBaseAuthenticationTypePtrOutput
}

type hbaseAuthenticationTypePtr string

func HBaseAuthenticationTypePtr(v string) HBaseAuthenticationTypePtrInput {
	return (*hbaseAuthenticationTypePtr)(&v)
}

func (*hbaseAuthenticationTypePtr) ElementType() reflect.Type {
	return hbaseAuthenticationTypePtrType
}

func (in *hbaseAuthenticationTypePtr) ToHBaseAuthenticationTypePtrOutput() HBaseAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(HBaseAuthenticationTypePtrOutput)
}

func (in *hbaseAuthenticationTypePtr) ToHBaseAuthenticationTypePtrOutputWithContext(ctx context.Context) HBaseAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HBaseAuthenticationTypePtrOutput)
}

type HDInsightActivityDebugInfoOption string

const (
	HDInsightActivityDebugInfoOptionNone    = HDInsightActivityDebugInfoOption("None")
	HDInsightActivityDebugInfoOptionAlways  = HDInsightActivityDebugInfoOption("Always")
	HDInsightActivityDebugInfoOptionFailure = HDInsightActivityDebugInfoOption("Failure")
)

func (HDInsightActivityDebugInfoOption) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightActivityDebugInfoOption)(nil)).Elem()
}

func (e HDInsightActivityDebugInfoOption) ToHDInsightActivityDebugInfoOptionOutput() HDInsightActivityDebugInfoOptionOutput {
	return pulumi.ToOutput(e).(HDInsightActivityDebugInfoOptionOutput)
}

func (e HDInsightActivityDebugInfoOption) ToHDInsightActivityDebugInfoOptionOutputWithContext(ctx context.Context) HDInsightActivityDebugInfoOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HDInsightActivityDebugInfoOptionOutput)
}

func (e HDInsightActivityDebugInfoOption) ToHDInsightActivityDebugInfoOptionPtrOutput() HDInsightActivityDebugInfoOptionPtrOutput {
	return e.ToHDInsightActivityDebugInfoOptionPtrOutputWithContext(context.Background())
}

func (e HDInsightActivityDebugInfoOption) ToHDInsightActivityDebugInfoOptionPtrOutputWithContext(ctx context.Context) HDInsightActivityDebugInfoOptionPtrOutput {
	return HDInsightActivityDebugInfoOption(e).ToHDInsightActivityDebugInfoOptionOutputWithContext(ctx).ToHDInsightActivityDebugInfoOptionPtrOutputWithContext(ctx)
}

func (e HDInsightActivityDebugInfoOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HDInsightActivityDebugInfoOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HDInsightActivityDebugInfoOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HDInsightActivityDebugInfoOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HDInsightActivityDebugInfoOptionOutput struct{ *pulumi.OutputState }

func (HDInsightActivityDebugInfoOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HDInsightActivityDebugInfoOption)(nil)).Elem()
}

func (o HDInsightActivityDebugInfoOptionOutput) ToHDInsightActivityDebugInfoOptionOutput() HDInsightActivityDebugInfoOptionOutput {
	return o
}

func (o HDInsightActivityDebugInfoOptionOutput) ToHDInsightActivityDebugInfoOptionOutputWithContext(ctx context.Context) HDInsightActivityDebugInfoOptionOutput {
	return o
}

func (o HDInsightActivityDebugInfoOptionOutput) ToHDInsightActivityDebugInfoOptionPtrOutput() HDInsightActivityDebugInfoOptionPtrOutput {
	return o.ToHDInsightActivityDebugInfoOptionPtrOutputWithContext(context.Background())
}

func (o HDInsightActivityDebugInfoOptionOutput) ToHDInsightActivityDebugInfoOptionPtrOutputWithContext(ctx context.Context) HDInsightActivityDebugInfoOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HDInsightActivityDebugInfoOption) *HDInsightActivityDebugInfoOption {
		return &v
	}).(HDInsightActivityDebugInfoOptionPtrOutput)
}

func (o HDInsightActivityDebugInfoOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HDInsightActivityDebugInfoOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HDInsightActivityDebugInfoOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HDInsightActivityDebugInfoOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HDInsightActivityDebugInfoOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HDInsightActivityDebugInfoOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HDInsightActivityDebugInfoOptionPtrOutput struct{ *pulumi.OutputState }

func (HDInsightActivityDebugInfoOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HDInsightActivityDebugInfoOption)(nil)).Elem()
}

func (o HDInsightActivityDebugInfoOptionPtrOutput) ToHDInsightActivityDebugInfoOptionPtrOutput() HDInsightActivityDebugInfoOptionPtrOutput {
	return o
}

func (o HDInsightActivityDebugInfoOptionPtrOutput) ToHDInsightActivityDebugInfoOptionPtrOutputWithContext(ctx context.Context) HDInsightActivityDebugInfoOptionPtrOutput {
	return o
}

func (o HDInsightActivityDebugInfoOptionPtrOutput) Elem() HDInsightActivityDebugInfoOptionOutput {
	return o.ApplyT(func(v *HDInsightActivityDebugInfoOption) HDInsightActivityDebugInfoOption {
		if v != nil {
			return *v
		}
		var ret HDInsightActivityDebugInfoOption
		return ret
	}).(HDInsightActivityDebugInfoOptionOutput)
}

func (o HDInsightActivityDebugInfoOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HDInsightActivityDebugInfoOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HDInsightActivityDebugInfoOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HDInsightActivityDebugInfoOptionInput interface {
	pulumi.Input

	ToHDInsightActivityDebugInfoOptionOutput() HDInsightActivityDebugInfoOptionOutput
	ToHDInsightActivityDebugInfoOptionOutputWithContext(context.Context) HDInsightActivityDebugInfoOptionOutput
}

var hdinsightActivityDebugInfoOptionPtrType = reflect.TypeOf((**HDInsightActivityDebugInfoOption)(nil)).Elem()

type HDInsightActivityDebugInfoOptionPtrInput interface {
	pulumi.Input

	ToHDInsightActivityDebugInfoOptionPtrOutput() HDInsightActivityDebugInfoOptionPtrOutput
	ToHDInsightActivityDebugInfoOptionPtrOutputWithContext(context.Context) HDInsightActivityDebugInfoOptionPtrOutput
}

type hdinsightActivityDebugInfoOptionPtr string

func HDInsightActivityDebugInfoOptionPtr(v string) HDInsightActivityDebugInfoOptionPtrInput {
	return (*hdinsightActivityDebugInfoOptionPtr)(&v)
}

func (*hdinsightActivityDebugInfoOptionPtr) ElementType() reflect.Type {
	return hdinsightActivityDebugInfoOptionPtrType
}

func (in *hdinsightActivityDebugInfoOptionPtr) ToHDInsightActivityDebugInfoOptionPtrOutput() HDInsightActivityDebugInfoOptionPtrOutput {
	return pulumi.ToOutput(in).(HDInsightActivityDebugInfoOptionPtrOutput)
}

func (in *hdinsightActivityDebugInfoOptionPtr) ToHDInsightActivityDebugInfoOptionPtrOutputWithContext(ctx context.Context) HDInsightActivityDebugInfoOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HDInsightActivityDebugInfoOptionPtrOutput)
}

type HiveAuthenticationType string

const (
	HiveAuthenticationTypeAnonymous                    = HiveAuthenticationType("Anonymous")
	HiveAuthenticationTypeUsername                     = HiveAuthenticationType("Username")
	HiveAuthenticationTypeUsernameAndPassword          = HiveAuthenticationType("UsernameAndPassword")
	HiveAuthenticationTypeWindowsAzureHDInsightService = HiveAuthenticationType("WindowsAzureHDInsightService")
)

func (HiveAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*HiveAuthenticationType)(nil)).Elem()
}

func (e HiveAuthenticationType) ToHiveAuthenticationTypeOutput() HiveAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(HiveAuthenticationTypeOutput)
}

func (e HiveAuthenticationType) ToHiveAuthenticationTypeOutputWithContext(ctx context.Context) HiveAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HiveAuthenticationTypeOutput)
}

func (e HiveAuthenticationType) ToHiveAuthenticationTypePtrOutput() HiveAuthenticationTypePtrOutput {
	return e.ToHiveAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e HiveAuthenticationType) ToHiveAuthenticationTypePtrOutputWithContext(ctx context.Context) HiveAuthenticationTypePtrOutput {
	return HiveAuthenticationType(e).ToHiveAuthenticationTypeOutputWithContext(ctx).ToHiveAuthenticationTypePtrOutputWithContext(ctx)
}

func (e HiveAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HiveAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HiveAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HiveAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HiveAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (HiveAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HiveAuthenticationType)(nil)).Elem()
}

func (o HiveAuthenticationTypeOutput) ToHiveAuthenticationTypeOutput() HiveAuthenticationTypeOutput {
	return o
}

func (o HiveAuthenticationTypeOutput) ToHiveAuthenticationTypeOutputWithContext(ctx context.Context) HiveAuthenticationTypeOutput {
	return o
}

func (o HiveAuthenticationTypeOutput) ToHiveAuthenticationTypePtrOutput() HiveAuthenticationTypePtrOutput {
	return o.ToHiveAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o HiveAuthenticationTypeOutput) ToHiveAuthenticationTypePtrOutputWithContext(ctx context.Context) HiveAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HiveAuthenticationType) *HiveAuthenticationType {
		return &v
	}).(HiveAuthenticationTypePtrOutput)
}

func (o HiveAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HiveAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HiveAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HiveAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HiveAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HiveAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HiveAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (HiveAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HiveAuthenticationType)(nil)).Elem()
}

func (o HiveAuthenticationTypePtrOutput) ToHiveAuthenticationTypePtrOutput() HiveAuthenticationTypePtrOutput {
	return o
}

func (o HiveAuthenticationTypePtrOutput) ToHiveAuthenticationTypePtrOutputWithContext(ctx context.Context) HiveAuthenticationTypePtrOutput {
	return o
}

func (o HiveAuthenticationTypePtrOutput) Elem() HiveAuthenticationTypeOutput {
	return o.ApplyT(func(v *HiveAuthenticationType) HiveAuthenticationType {
		if v != nil {
			return *v
		}
		var ret HiveAuthenticationType
		return ret
	}).(HiveAuthenticationTypeOutput)
}

func (o HiveAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HiveAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HiveAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HiveAuthenticationTypeInput interface {
	pulumi.Input

	ToHiveAuthenticationTypeOutput() HiveAuthenticationTypeOutput
	ToHiveAuthenticationTypeOutputWithContext(context.Context) HiveAuthenticationTypeOutput
}

var hiveAuthenticationTypePtrType = reflect.TypeOf((**HiveAuthenticationType)(nil)).Elem()

type HiveAuthenticationTypePtrInput interface {
	pulumi.Input

	ToHiveAuthenticationTypePtrOutput() HiveAuthenticationTypePtrOutput
	ToHiveAuthenticationTypePtrOutputWithContext(context.Context) HiveAuthenticationTypePtrOutput
}

type hiveAuthenticationTypePtr string

func HiveAuthenticationTypePtr(v string) HiveAuthenticationTypePtrInput {
	return (*hiveAuthenticationTypePtr)(&v)
}

func (*hiveAuthenticationTypePtr) ElementType() reflect.Type {
	return hiveAuthenticationTypePtrType
}

func (in *hiveAuthenticationTypePtr) ToHiveAuthenticationTypePtrOutput() HiveAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(HiveAuthenticationTypePtrOutput)
}

func (in *hiveAuthenticationTypePtr) ToHiveAuthenticationTypePtrOutputWithContext(ctx context.Context) HiveAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HiveAuthenticationTypePtrOutput)
}

type HiveServerType string

const (
	HiveServerTypeHiveServer1      = HiveServerType("HiveServer1")
	HiveServerTypeHiveServer2      = HiveServerType("HiveServer2")
	HiveServerTypeHiveThriftServer = HiveServerType("HiveThriftServer")
)

func (HiveServerType) ElementType() reflect.Type {
	return reflect.TypeOf((*HiveServerType)(nil)).Elem()
}

func (e HiveServerType) ToHiveServerTypeOutput() HiveServerTypeOutput {
	return pulumi.ToOutput(e).(HiveServerTypeOutput)
}

func (e HiveServerType) ToHiveServerTypeOutputWithContext(ctx context.Context) HiveServerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HiveServerTypeOutput)
}

func (e HiveServerType) ToHiveServerTypePtrOutput() HiveServerTypePtrOutput {
	return e.ToHiveServerTypePtrOutputWithContext(context.Background())
}

func (e HiveServerType) ToHiveServerTypePtrOutputWithContext(ctx context.Context) HiveServerTypePtrOutput {
	return HiveServerType(e).ToHiveServerTypeOutputWithContext(ctx).ToHiveServerTypePtrOutputWithContext(ctx)
}

func (e HiveServerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HiveServerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HiveServerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HiveServerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HiveServerTypeOutput struct{ *pulumi.OutputState }

func (HiveServerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HiveServerType)(nil)).Elem()
}

func (o HiveServerTypeOutput) ToHiveServerTypeOutput() HiveServerTypeOutput {
	return o
}

func (o HiveServerTypeOutput) ToHiveServerTypeOutputWithContext(ctx context.Context) HiveServerTypeOutput {
	return o
}

func (o HiveServerTypeOutput) ToHiveServerTypePtrOutput() HiveServerTypePtrOutput {
	return o.ToHiveServerTypePtrOutputWithContext(context.Background())
}

func (o HiveServerTypeOutput) ToHiveServerTypePtrOutputWithContext(ctx context.Context) HiveServerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HiveServerType) *HiveServerType {
		return &v
	}).(HiveServerTypePtrOutput)
}

func (o HiveServerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HiveServerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HiveServerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HiveServerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HiveServerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HiveServerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HiveServerTypePtrOutput struct{ *pulumi.OutputState }

func (HiveServerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HiveServerType)(nil)).Elem()
}

func (o HiveServerTypePtrOutput) ToHiveServerTypePtrOutput() HiveServerTypePtrOutput {
	return o
}

func (o HiveServerTypePtrOutput) ToHiveServerTypePtrOutputWithContext(ctx context.Context) HiveServerTypePtrOutput {
	return o
}

func (o HiveServerTypePtrOutput) Elem() HiveServerTypeOutput {
	return o.ApplyT(func(v *HiveServerType) HiveServerType {
		if v != nil {
			return *v
		}
		var ret HiveServerType
		return ret
	}).(HiveServerTypeOutput)
}

func (o HiveServerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HiveServerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HiveServerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HiveServerTypeInput interface {
	pulumi.Input

	ToHiveServerTypeOutput() HiveServerTypeOutput
	ToHiveServerTypeOutputWithContext(context.Context) HiveServerTypeOutput
}

var hiveServerTypePtrType = reflect.TypeOf((**HiveServerType)(nil)).Elem()

type HiveServerTypePtrInput interface {
	pulumi.Input

	ToHiveServerTypePtrOutput() HiveServerTypePtrOutput
	ToHiveServerTypePtrOutputWithContext(context.Context) HiveServerTypePtrOutput
}

type hiveServerTypePtr string

func HiveServerTypePtr(v string) HiveServerTypePtrInput {
	return (*hiveServerTypePtr)(&v)
}

func (*hiveServerTypePtr) ElementType() reflect.Type {
	return hiveServerTypePtrType
}

func (in *hiveServerTypePtr) ToHiveServerTypePtrOutput() HiveServerTypePtrOutput {
	return pulumi.ToOutput(in).(HiveServerTypePtrOutput)
}

func (in *hiveServerTypePtr) ToHiveServerTypePtrOutputWithContext(ctx context.Context) HiveServerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HiveServerTypePtrOutput)
}

type HiveThriftTransportProtocol string

const (
	HiveThriftTransportProtocolBinary = HiveThriftTransportProtocol("Binary")
	HiveThriftTransportProtocolSASL   = HiveThriftTransportProtocol("SASL")
	HiveThriftTransportProtocol_HTTP_ = HiveThriftTransportProtocol("HTTP ")
)

func (HiveThriftTransportProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*HiveThriftTransportProtocol)(nil)).Elem()
}

func (e HiveThriftTransportProtocol) ToHiveThriftTransportProtocolOutput() HiveThriftTransportProtocolOutput {
	return pulumi.ToOutput(e).(HiveThriftTransportProtocolOutput)
}

func (e HiveThriftTransportProtocol) ToHiveThriftTransportProtocolOutputWithContext(ctx context.Context) HiveThriftTransportProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HiveThriftTransportProtocolOutput)
}

func (e HiveThriftTransportProtocol) ToHiveThriftTransportProtocolPtrOutput() HiveThriftTransportProtocolPtrOutput {
	return e.ToHiveThriftTransportProtocolPtrOutputWithContext(context.Background())
}

func (e HiveThriftTransportProtocol) ToHiveThriftTransportProtocolPtrOutputWithContext(ctx context.Context) HiveThriftTransportProtocolPtrOutput {
	return HiveThriftTransportProtocol(e).ToHiveThriftTransportProtocolOutputWithContext(ctx).ToHiveThriftTransportProtocolPtrOutputWithContext(ctx)
}

func (e HiveThriftTransportProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HiveThriftTransportProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HiveThriftTransportProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HiveThriftTransportProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HiveThriftTransportProtocolOutput struct{ *pulumi.OutputState }

func (HiveThriftTransportProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HiveThriftTransportProtocol)(nil)).Elem()
}

func (o HiveThriftTransportProtocolOutput) ToHiveThriftTransportProtocolOutput() HiveThriftTransportProtocolOutput {
	return o
}

func (o HiveThriftTransportProtocolOutput) ToHiveThriftTransportProtocolOutputWithContext(ctx context.Context) HiveThriftTransportProtocolOutput {
	return o
}

func (o HiveThriftTransportProtocolOutput) ToHiveThriftTransportProtocolPtrOutput() HiveThriftTransportProtocolPtrOutput {
	return o.ToHiveThriftTransportProtocolPtrOutputWithContext(context.Background())
}

func (o HiveThriftTransportProtocolOutput) ToHiveThriftTransportProtocolPtrOutputWithContext(ctx context.Context) HiveThriftTransportProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HiveThriftTransportProtocol) *HiveThriftTransportProtocol {
		return &v
	}).(HiveThriftTransportProtocolPtrOutput)
}

func (o HiveThriftTransportProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HiveThriftTransportProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HiveThriftTransportProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HiveThriftTransportProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HiveThriftTransportProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HiveThriftTransportProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HiveThriftTransportProtocolPtrOutput struct{ *pulumi.OutputState }

func (HiveThriftTransportProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HiveThriftTransportProtocol)(nil)).Elem()
}

func (o HiveThriftTransportProtocolPtrOutput) ToHiveThriftTransportProtocolPtrOutput() HiveThriftTransportProtocolPtrOutput {
	return o
}

func (o HiveThriftTransportProtocolPtrOutput) ToHiveThriftTransportProtocolPtrOutputWithContext(ctx context.Context) HiveThriftTransportProtocolPtrOutput {
	return o
}

func (o HiveThriftTransportProtocolPtrOutput) Elem() HiveThriftTransportProtocolOutput {
	return o.ApplyT(func(v *HiveThriftTransportProtocol) HiveThriftTransportProtocol {
		if v != nil {
			return *v
		}
		var ret HiveThriftTransportProtocol
		return ret
	}).(HiveThriftTransportProtocolOutput)
}

func (o HiveThriftTransportProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HiveThriftTransportProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HiveThriftTransportProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HiveThriftTransportProtocolInput interface {
	pulumi.Input

	ToHiveThriftTransportProtocolOutput() HiveThriftTransportProtocolOutput
	ToHiveThriftTransportProtocolOutputWithContext(context.Context) HiveThriftTransportProtocolOutput
}

var hiveThriftTransportProtocolPtrType = reflect.TypeOf((**HiveThriftTransportProtocol)(nil)).Elem()

type HiveThriftTransportProtocolPtrInput interface {
	pulumi.Input

	ToHiveThriftTransportProtocolPtrOutput() HiveThriftTransportProtocolPtrOutput
	ToHiveThriftTransportProtocolPtrOutputWithContext(context.Context) HiveThriftTransportProtocolPtrOutput
}

type hiveThriftTransportProtocolPtr string

func HiveThriftTransportProtocolPtr(v string) HiveThriftTransportProtocolPtrInput {
	return (*hiveThriftTransportProtocolPtr)(&v)
}

func (*hiveThriftTransportProtocolPtr) ElementType() reflect.Type {
	return hiveThriftTransportProtocolPtrType
}

func (in *hiveThriftTransportProtocolPtr) ToHiveThriftTransportProtocolPtrOutput() HiveThriftTransportProtocolPtrOutput {
	return pulumi.ToOutput(in).(HiveThriftTransportProtocolPtrOutput)
}

func (in *hiveThriftTransportProtocolPtr) ToHiveThriftTransportProtocolPtrOutputWithContext(ctx context.Context) HiveThriftTransportProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HiveThriftTransportProtocolPtrOutput)
}

type HttpAuthenticationType string

const (
	HttpAuthenticationTypeBasic             = HttpAuthenticationType("Basic")
	HttpAuthenticationTypeAnonymous         = HttpAuthenticationType("Anonymous")
	HttpAuthenticationTypeDigest            = HttpAuthenticationType("Digest")
	HttpAuthenticationTypeWindows           = HttpAuthenticationType("Windows")
	HttpAuthenticationTypeClientCertificate = HttpAuthenticationType("ClientCertificate")
)

func (HttpAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpAuthenticationType)(nil)).Elem()
}

func (e HttpAuthenticationType) ToHttpAuthenticationTypeOutput() HttpAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(HttpAuthenticationTypeOutput)
}

func (e HttpAuthenticationType) ToHttpAuthenticationTypeOutputWithContext(ctx context.Context) HttpAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HttpAuthenticationTypeOutput)
}

func (e HttpAuthenticationType) ToHttpAuthenticationTypePtrOutput() HttpAuthenticationTypePtrOutput {
	return e.ToHttpAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e HttpAuthenticationType) ToHttpAuthenticationTypePtrOutputWithContext(ctx context.Context) HttpAuthenticationTypePtrOutput {
	return HttpAuthenticationType(e).ToHttpAuthenticationTypeOutputWithContext(ctx).ToHttpAuthenticationTypePtrOutputWithContext(ctx)
}

func (e HttpAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HttpAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HttpAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HttpAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (HttpAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpAuthenticationType)(nil)).Elem()
}

func (o HttpAuthenticationTypeOutput) ToHttpAuthenticationTypeOutput() HttpAuthenticationTypeOutput {
	return o
}

func (o HttpAuthenticationTypeOutput) ToHttpAuthenticationTypeOutputWithContext(ctx context.Context) HttpAuthenticationTypeOutput {
	return o
}

func (o HttpAuthenticationTypeOutput) ToHttpAuthenticationTypePtrOutput() HttpAuthenticationTypePtrOutput {
	return o.ToHttpAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o HttpAuthenticationTypeOutput) ToHttpAuthenticationTypePtrOutputWithContext(ctx context.Context) HttpAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpAuthenticationType) *HttpAuthenticationType {
		return &v
	}).(HttpAuthenticationTypePtrOutput)
}

func (o HttpAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HttpAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HttpAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HttpAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HttpAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (HttpAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpAuthenticationType)(nil)).Elem()
}

func (o HttpAuthenticationTypePtrOutput) ToHttpAuthenticationTypePtrOutput() HttpAuthenticationTypePtrOutput {
	return o
}

func (o HttpAuthenticationTypePtrOutput) ToHttpAuthenticationTypePtrOutputWithContext(ctx context.Context) HttpAuthenticationTypePtrOutput {
	return o
}

func (o HttpAuthenticationTypePtrOutput) Elem() HttpAuthenticationTypeOutput {
	return o.ApplyT(func(v *HttpAuthenticationType) HttpAuthenticationType {
		if v != nil {
			return *v
		}
		var ret HttpAuthenticationType
		return ret
	}).(HttpAuthenticationTypeOutput)
}

func (o HttpAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HttpAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HttpAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HttpAuthenticationTypeInput interface {
	pulumi.Input

	ToHttpAuthenticationTypeOutput() HttpAuthenticationTypeOutput
	ToHttpAuthenticationTypeOutputWithContext(context.Context) HttpAuthenticationTypeOutput
}

var httpAuthenticationTypePtrType = reflect.TypeOf((**HttpAuthenticationType)(nil)).Elem()

type HttpAuthenticationTypePtrInput interface {
	pulumi.Input

	ToHttpAuthenticationTypePtrOutput() HttpAuthenticationTypePtrOutput
	ToHttpAuthenticationTypePtrOutputWithContext(context.Context) HttpAuthenticationTypePtrOutput
}

type httpAuthenticationTypePtr string

func HttpAuthenticationTypePtr(v string) HttpAuthenticationTypePtrInput {
	return (*httpAuthenticationTypePtr)(&v)
}

func (*httpAuthenticationTypePtr) ElementType() reflect.Type {
	return httpAuthenticationTypePtrType
}

func (in *httpAuthenticationTypePtr) ToHttpAuthenticationTypePtrOutput() HttpAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(HttpAuthenticationTypePtrOutput)
}

func (in *httpAuthenticationTypePtr) ToHttpAuthenticationTypePtrOutputWithContext(ctx context.Context) HttpAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HttpAuthenticationTypePtrOutput)
}

type ImpalaAuthenticationType string

const (
	ImpalaAuthenticationTypeAnonymous           = ImpalaAuthenticationType("Anonymous")
	ImpalaAuthenticationTypeSASLUsername        = ImpalaAuthenticationType("SASLUsername")
	ImpalaAuthenticationTypeUsernameAndPassword = ImpalaAuthenticationType("UsernameAndPassword")
)

func (ImpalaAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ImpalaAuthenticationType)(nil)).Elem()
}

func (e ImpalaAuthenticationType) ToImpalaAuthenticationTypeOutput() ImpalaAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(ImpalaAuthenticationTypeOutput)
}

func (e ImpalaAuthenticationType) ToImpalaAuthenticationTypeOutputWithContext(ctx context.Context) ImpalaAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ImpalaAuthenticationTypeOutput)
}

func (e ImpalaAuthenticationType) ToImpalaAuthenticationTypePtrOutput() ImpalaAuthenticationTypePtrOutput {
	return e.ToImpalaAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e ImpalaAuthenticationType) ToImpalaAuthenticationTypePtrOutputWithContext(ctx context.Context) ImpalaAuthenticationTypePtrOutput {
	return ImpalaAuthenticationType(e).ToImpalaAuthenticationTypeOutputWithContext(ctx).ToImpalaAuthenticationTypePtrOutputWithContext(ctx)
}

func (e ImpalaAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImpalaAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImpalaAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ImpalaAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ImpalaAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (ImpalaAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImpalaAuthenticationType)(nil)).Elem()
}

func (o ImpalaAuthenticationTypeOutput) ToImpalaAuthenticationTypeOutput() ImpalaAuthenticationTypeOutput {
	return o
}

func (o ImpalaAuthenticationTypeOutput) ToImpalaAuthenticationTypeOutputWithContext(ctx context.Context) ImpalaAuthenticationTypeOutput {
	return o
}

func (o ImpalaAuthenticationTypeOutput) ToImpalaAuthenticationTypePtrOutput() ImpalaAuthenticationTypePtrOutput {
	return o.ToImpalaAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o ImpalaAuthenticationTypeOutput) ToImpalaAuthenticationTypePtrOutputWithContext(ctx context.Context) ImpalaAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImpalaAuthenticationType) *ImpalaAuthenticationType {
		return &v
	}).(ImpalaAuthenticationTypePtrOutput)
}

func (o ImpalaAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ImpalaAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImpalaAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ImpalaAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImpalaAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImpalaAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ImpalaAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (ImpalaAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImpalaAuthenticationType)(nil)).Elem()
}

func (o ImpalaAuthenticationTypePtrOutput) ToImpalaAuthenticationTypePtrOutput() ImpalaAuthenticationTypePtrOutput {
	return o
}

func (o ImpalaAuthenticationTypePtrOutput) ToImpalaAuthenticationTypePtrOutputWithContext(ctx context.Context) ImpalaAuthenticationTypePtrOutput {
	return o
}

func (o ImpalaAuthenticationTypePtrOutput) Elem() ImpalaAuthenticationTypeOutput {
	return o.ApplyT(func(v *ImpalaAuthenticationType) ImpalaAuthenticationType {
		if v != nil {
			return *v
		}
		var ret ImpalaAuthenticationType
		return ret
	}).(ImpalaAuthenticationTypeOutput)
}

func (o ImpalaAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImpalaAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ImpalaAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ImpalaAuthenticationTypeInput interface {
	pulumi.Input

	ToImpalaAuthenticationTypeOutput() ImpalaAuthenticationTypeOutput
	ToImpalaAuthenticationTypeOutputWithContext(context.Context) ImpalaAuthenticationTypeOutput
}

var impalaAuthenticationTypePtrType = reflect.TypeOf((**ImpalaAuthenticationType)(nil)).Elem()

type ImpalaAuthenticationTypePtrInput interface {
	pulumi.Input

	ToImpalaAuthenticationTypePtrOutput() ImpalaAuthenticationTypePtrOutput
	ToImpalaAuthenticationTypePtrOutputWithContext(context.Context) ImpalaAuthenticationTypePtrOutput
}

type impalaAuthenticationTypePtr string

func ImpalaAuthenticationTypePtr(v string) ImpalaAuthenticationTypePtrInput {
	return (*impalaAuthenticationTypePtr)(&v)
}

func (*impalaAuthenticationTypePtr) ElementType() reflect.Type {
	return impalaAuthenticationTypePtrType
}

func (in *impalaAuthenticationTypePtr) ToImpalaAuthenticationTypePtrOutput() ImpalaAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(ImpalaAuthenticationTypePtrOutput)
}

func (in *impalaAuthenticationTypePtr) ToImpalaAuthenticationTypePtrOutputWithContext(ctx context.Context) ImpalaAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ImpalaAuthenticationTypePtrOutput)
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

type MongoDbAuthenticationType string

const (
	MongoDbAuthenticationTypeBasic     = MongoDbAuthenticationType("Basic")
	MongoDbAuthenticationTypeAnonymous = MongoDbAuthenticationType("Anonymous")
)

func (MongoDbAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDbAuthenticationType)(nil)).Elem()
}

func (e MongoDbAuthenticationType) ToMongoDbAuthenticationTypeOutput() MongoDbAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(MongoDbAuthenticationTypeOutput)
}

func (e MongoDbAuthenticationType) ToMongoDbAuthenticationTypeOutputWithContext(ctx context.Context) MongoDbAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MongoDbAuthenticationTypeOutput)
}

func (e MongoDbAuthenticationType) ToMongoDbAuthenticationTypePtrOutput() MongoDbAuthenticationTypePtrOutput {
	return e.ToMongoDbAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e MongoDbAuthenticationType) ToMongoDbAuthenticationTypePtrOutputWithContext(ctx context.Context) MongoDbAuthenticationTypePtrOutput {
	return MongoDbAuthenticationType(e).ToMongoDbAuthenticationTypeOutputWithContext(ctx).ToMongoDbAuthenticationTypePtrOutputWithContext(ctx)
}

func (e MongoDbAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MongoDbAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MongoDbAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MongoDbAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MongoDbAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (MongoDbAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDbAuthenticationType)(nil)).Elem()
}

func (o MongoDbAuthenticationTypeOutput) ToMongoDbAuthenticationTypeOutput() MongoDbAuthenticationTypeOutput {
	return o
}

func (o MongoDbAuthenticationTypeOutput) ToMongoDbAuthenticationTypeOutputWithContext(ctx context.Context) MongoDbAuthenticationTypeOutput {
	return o
}

func (o MongoDbAuthenticationTypeOutput) ToMongoDbAuthenticationTypePtrOutput() MongoDbAuthenticationTypePtrOutput {
	return o.ToMongoDbAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o MongoDbAuthenticationTypeOutput) ToMongoDbAuthenticationTypePtrOutputWithContext(ctx context.Context) MongoDbAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDbAuthenticationType) *MongoDbAuthenticationType {
		return &v
	}).(MongoDbAuthenticationTypePtrOutput)
}

func (o MongoDbAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MongoDbAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MongoDbAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MongoDbAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MongoDbAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MongoDbAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MongoDbAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (MongoDbAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDbAuthenticationType)(nil)).Elem()
}

func (o MongoDbAuthenticationTypePtrOutput) ToMongoDbAuthenticationTypePtrOutput() MongoDbAuthenticationTypePtrOutput {
	return o
}

func (o MongoDbAuthenticationTypePtrOutput) ToMongoDbAuthenticationTypePtrOutputWithContext(ctx context.Context) MongoDbAuthenticationTypePtrOutput {
	return o
}

func (o MongoDbAuthenticationTypePtrOutput) Elem() MongoDbAuthenticationTypeOutput {
	return o.ApplyT(func(v *MongoDbAuthenticationType) MongoDbAuthenticationType {
		if v != nil {
			return *v
		}
		var ret MongoDbAuthenticationType
		return ret
	}).(MongoDbAuthenticationTypeOutput)
}

func (o MongoDbAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MongoDbAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MongoDbAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MongoDbAuthenticationTypeInput interface {
	pulumi.Input

	ToMongoDbAuthenticationTypeOutput() MongoDbAuthenticationTypeOutput
	ToMongoDbAuthenticationTypeOutputWithContext(context.Context) MongoDbAuthenticationTypeOutput
}

var mongoDbAuthenticationTypePtrType = reflect.TypeOf((**MongoDbAuthenticationType)(nil)).Elem()

type MongoDbAuthenticationTypePtrInput interface {
	pulumi.Input

	ToMongoDbAuthenticationTypePtrOutput() MongoDbAuthenticationTypePtrOutput
	ToMongoDbAuthenticationTypePtrOutputWithContext(context.Context) MongoDbAuthenticationTypePtrOutput
}

type mongoDbAuthenticationTypePtr string

func MongoDbAuthenticationTypePtr(v string) MongoDbAuthenticationTypePtrInput {
	return (*mongoDbAuthenticationTypePtr)(&v)
}

func (*mongoDbAuthenticationTypePtr) ElementType() reflect.Type {
	return mongoDbAuthenticationTypePtrType
}

func (in *mongoDbAuthenticationTypePtr) ToMongoDbAuthenticationTypePtrOutput() MongoDbAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(MongoDbAuthenticationTypePtrOutput)
}

func (in *mongoDbAuthenticationTypePtr) ToMongoDbAuthenticationTypePtrOutputWithContext(ctx context.Context) MongoDbAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MongoDbAuthenticationTypePtrOutput)
}

type ODataAadServicePrincipalCredentialType string

const (
	ODataAadServicePrincipalCredentialTypeServicePrincipalKey  = ODataAadServicePrincipalCredentialType("ServicePrincipalKey")
	ODataAadServicePrincipalCredentialTypeServicePrincipalCert = ODataAadServicePrincipalCredentialType("ServicePrincipalCert")
)

func (ODataAadServicePrincipalCredentialType) ElementType() reflect.Type {
	return reflect.TypeOf((*ODataAadServicePrincipalCredentialType)(nil)).Elem()
}

func (e ODataAadServicePrincipalCredentialType) ToODataAadServicePrincipalCredentialTypeOutput() ODataAadServicePrincipalCredentialTypeOutput {
	return pulumi.ToOutput(e).(ODataAadServicePrincipalCredentialTypeOutput)
}

func (e ODataAadServicePrincipalCredentialType) ToODataAadServicePrincipalCredentialTypeOutputWithContext(ctx context.Context) ODataAadServicePrincipalCredentialTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ODataAadServicePrincipalCredentialTypeOutput)
}

func (e ODataAadServicePrincipalCredentialType) ToODataAadServicePrincipalCredentialTypePtrOutput() ODataAadServicePrincipalCredentialTypePtrOutput {
	return e.ToODataAadServicePrincipalCredentialTypePtrOutputWithContext(context.Background())
}

func (e ODataAadServicePrincipalCredentialType) ToODataAadServicePrincipalCredentialTypePtrOutputWithContext(ctx context.Context) ODataAadServicePrincipalCredentialTypePtrOutput {
	return ODataAadServicePrincipalCredentialType(e).ToODataAadServicePrincipalCredentialTypeOutputWithContext(ctx).ToODataAadServicePrincipalCredentialTypePtrOutputWithContext(ctx)
}

func (e ODataAadServicePrincipalCredentialType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ODataAadServicePrincipalCredentialType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ODataAadServicePrincipalCredentialType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ODataAadServicePrincipalCredentialType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ODataAadServicePrincipalCredentialTypeOutput struct{ *pulumi.OutputState }

func (ODataAadServicePrincipalCredentialTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ODataAadServicePrincipalCredentialType)(nil)).Elem()
}

func (o ODataAadServicePrincipalCredentialTypeOutput) ToODataAadServicePrincipalCredentialTypeOutput() ODataAadServicePrincipalCredentialTypeOutput {
	return o
}

func (o ODataAadServicePrincipalCredentialTypeOutput) ToODataAadServicePrincipalCredentialTypeOutputWithContext(ctx context.Context) ODataAadServicePrincipalCredentialTypeOutput {
	return o
}

func (o ODataAadServicePrincipalCredentialTypeOutput) ToODataAadServicePrincipalCredentialTypePtrOutput() ODataAadServicePrincipalCredentialTypePtrOutput {
	return o.ToODataAadServicePrincipalCredentialTypePtrOutputWithContext(context.Background())
}

func (o ODataAadServicePrincipalCredentialTypeOutput) ToODataAadServicePrincipalCredentialTypePtrOutputWithContext(ctx context.Context) ODataAadServicePrincipalCredentialTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ODataAadServicePrincipalCredentialType) *ODataAadServicePrincipalCredentialType {
		return &v
	}).(ODataAadServicePrincipalCredentialTypePtrOutput)
}

func (o ODataAadServicePrincipalCredentialTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ODataAadServicePrincipalCredentialTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ODataAadServicePrincipalCredentialType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ODataAadServicePrincipalCredentialTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ODataAadServicePrincipalCredentialTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ODataAadServicePrincipalCredentialType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ODataAadServicePrincipalCredentialTypePtrOutput struct{ *pulumi.OutputState }

func (ODataAadServicePrincipalCredentialTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ODataAadServicePrincipalCredentialType)(nil)).Elem()
}

func (o ODataAadServicePrincipalCredentialTypePtrOutput) ToODataAadServicePrincipalCredentialTypePtrOutput() ODataAadServicePrincipalCredentialTypePtrOutput {
	return o
}

func (o ODataAadServicePrincipalCredentialTypePtrOutput) ToODataAadServicePrincipalCredentialTypePtrOutputWithContext(ctx context.Context) ODataAadServicePrincipalCredentialTypePtrOutput {
	return o
}

func (o ODataAadServicePrincipalCredentialTypePtrOutput) Elem() ODataAadServicePrincipalCredentialTypeOutput {
	return o.ApplyT(func(v *ODataAadServicePrincipalCredentialType) ODataAadServicePrincipalCredentialType {
		if v != nil {
			return *v
		}
		var ret ODataAadServicePrincipalCredentialType
		return ret
	}).(ODataAadServicePrincipalCredentialTypeOutput)
}

func (o ODataAadServicePrincipalCredentialTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ODataAadServicePrincipalCredentialTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ODataAadServicePrincipalCredentialType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ODataAadServicePrincipalCredentialTypeInput interface {
	pulumi.Input

	ToODataAadServicePrincipalCredentialTypeOutput() ODataAadServicePrincipalCredentialTypeOutput
	ToODataAadServicePrincipalCredentialTypeOutputWithContext(context.Context) ODataAadServicePrincipalCredentialTypeOutput
}

var odataAadServicePrincipalCredentialTypePtrType = reflect.TypeOf((**ODataAadServicePrincipalCredentialType)(nil)).Elem()

type ODataAadServicePrincipalCredentialTypePtrInput interface {
	pulumi.Input

	ToODataAadServicePrincipalCredentialTypePtrOutput() ODataAadServicePrincipalCredentialTypePtrOutput
	ToODataAadServicePrincipalCredentialTypePtrOutputWithContext(context.Context) ODataAadServicePrincipalCredentialTypePtrOutput
}

type odataAadServicePrincipalCredentialTypePtr string

func ODataAadServicePrincipalCredentialTypePtr(v string) ODataAadServicePrincipalCredentialTypePtrInput {
	return (*odataAadServicePrincipalCredentialTypePtr)(&v)
}

func (*odataAadServicePrincipalCredentialTypePtr) ElementType() reflect.Type {
	return odataAadServicePrincipalCredentialTypePtrType
}

func (in *odataAadServicePrincipalCredentialTypePtr) ToODataAadServicePrincipalCredentialTypePtrOutput() ODataAadServicePrincipalCredentialTypePtrOutput {
	return pulumi.ToOutput(in).(ODataAadServicePrincipalCredentialTypePtrOutput)
}

func (in *odataAadServicePrincipalCredentialTypePtr) ToODataAadServicePrincipalCredentialTypePtrOutputWithContext(ctx context.Context) ODataAadServicePrincipalCredentialTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ODataAadServicePrincipalCredentialTypePtrOutput)
}

type ODataAuthenticationType string

const (
	ODataAuthenticationTypeBasic                  = ODataAuthenticationType("Basic")
	ODataAuthenticationTypeAnonymous              = ODataAuthenticationType("Anonymous")
	ODataAuthenticationTypeWindows                = ODataAuthenticationType("Windows")
	ODataAuthenticationTypeAadServicePrincipal    = ODataAuthenticationType("AadServicePrincipal")
	ODataAuthenticationTypeManagedServiceIdentity = ODataAuthenticationType("ManagedServiceIdentity")
)

func (ODataAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ODataAuthenticationType)(nil)).Elem()
}

func (e ODataAuthenticationType) ToODataAuthenticationTypeOutput() ODataAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(ODataAuthenticationTypeOutput)
}

func (e ODataAuthenticationType) ToODataAuthenticationTypeOutputWithContext(ctx context.Context) ODataAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ODataAuthenticationTypeOutput)
}

func (e ODataAuthenticationType) ToODataAuthenticationTypePtrOutput() ODataAuthenticationTypePtrOutput {
	return e.ToODataAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e ODataAuthenticationType) ToODataAuthenticationTypePtrOutputWithContext(ctx context.Context) ODataAuthenticationTypePtrOutput {
	return ODataAuthenticationType(e).ToODataAuthenticationTypeOutputWithContext(ctx).ToODataAuthenticationTypePtrOutputWithContext(ctx)
}

func (e ODataAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ODataAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ODataAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ODataAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ODataAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (ODataAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ODataAuthenticationType)(nil)).Elem()
}

func (o ODataAuthenticationTypeOutput) ToODataAuthenticationTypeOutput() ODataAuthenticationTypeOutput {
	return o
}

func (o ODataAuthenticationTypeOutput) ToODataAuthenticationTypeOutputWithContext(ctx context.Context) ODataAuthenticationTypeOutput {
	return o
}

func (o ODataAuthenticationTypeOutput) ToODataAuthenticationTypePtrOutput() ODataAuthenticationTypePtrOutput {
	return o.ToODataAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o ODataAuthenticationTypeOutput) ToODataAuthenticationTypePtrOutputWithContext(ctx context.Context) ODataAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ODataAuthenticationType) *ODataAuthenticationType {
		return &v
	}).(ODataAuthenticationTypePtrOutput)
}

func (o ODataAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ODataAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ODataAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ODataAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ODataAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ODataAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ODataAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (ODataAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ODataAuthenticationType)(nil)).Elem()
}

func (o ODataAuthenticationTypePtrOutput) ToODataAuthenticationTypePtrOutput() ODataAuthenticationTypePtrOutput {
	return o
}

func (o ODataAuthenticationTypePtrOutput) ToODataAuthenticationTypePtrOutputWithContext(ctx context.Context) ODataAuthenticationTypePtrOutput {
	return o
}

func (o ODataAuthenticationTypePtrOutput) Elem() ODataAuthenticationTypeOutput {
	return o.ApplyT(func(v *ODataAuthenticationType) ODataAuthenticationType {
		if v != nil {
			return *v
		}
		var ret ODataAuthenticationType
		return ret
	}).(ODataAuthenticationTypeOutput)
}

func (o ODataAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ODataAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ODataAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ODataAuthenticationTypeInput interface {
	pulumi.Input

	ToODataAuthenticationTypeOutput() ODataAuthenticationTypeOutput
	ToODataAuthenticationTypeOutputWithContext(context.Context) ODataAuthenticationTypeOutput
}

var odataAuthenticationTypePtrType = reflect.TypeOf((**ODataAuthenticationType)(nil)).Elem()

type ODataAuthenticationTypePtrInput interface {
	pulumi.Input

	ToODataAuthenticationTypePtrOutput() ODataAuthenticationTypePtrOutput
	ToODataAuthenticationTypePtrOutputWithContext(context.Context) ODataAuthenticationTypePtrOutput
}

type odataAuthenticationTypePtr string

func ODataAuthenticationTypePtr(v string) ODataAuthenticationTypePtrInput {
	return (*odataAuthenticationTypePtr)(&v)
}

func (*odataAuthenticationTypePtr) ElementType() reflect.Type {
	return odataAuthenticationTypePtrType
}

func (in *odataAuthenticationTypePtr) ToODataAuthenticationTypePtrOutput() ODataAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(ODataAuthenticationTypePtrOutput)
}

func (in *odataAuthenticationTypePtr) ToODataAuthenticationTypePtrOutputWithContext(ctx context.Context) ODataAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ODataAuthenticationTypePtrOutput)
}

type ParameterType string

const (
	ParameterTypeObject       = ParameterType("Object")
	ParameterTypeString       = ParameterType("String")
	ParameterTypeInt          = ParameterType("Int")
	ParameterTypeFloat        = ParameterType("Float")
	ParameterTypeBool         = ParameterType("Bool")
	ParameterTypeArray        = ParameterType("Array")
	ParameterTypeSecureString = ParameterType("SecureString")
)

func (ParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterType)(nil)).Elem()
}

func (e ParameterType) ToParameterTypeOutput() ParameterTypeOutput {
	return pulumi.ToOutput(e).(ParameterTypeOutput)
}

func (e ParameterType) ToParameterTypeOutputWithContext(ctx context.Context) ParameterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ParameterTypeOutput)
}

func (e ParameterType) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return e.ToParameterTypePtrOutputWithContext(context.Background())
}

func (e ParameterType) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return ParameterType(e).ToParameterTypeOutputWithContext(ctx).ToParameterTypePtrOutputWithContext(ctx)
}

func (e ParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ParameterTypeOutput struct{ *pulumi.OutputState }

func (ParameterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterType)(nil)).Elem()
}

func (o ParameterTypeOutput) ToParameterTypeOutput() ParameterTypeOutput {
	return o
}

func (o ParameterTypeOutput) ToParameterTypeOutputWithContext(ctx context.Context) ParameterTypeOutput {
	return o
}

func (o ParameterTypeOutput) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return o.ToParameterTypePtrOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParameterType) *ParameterType {
		return &v
	}).(ParameterTypePtrOutput)
}

func (o ParameterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ParameterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ParameterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ParameterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ParameterTypePtrOutput struct{ *pulumi.OutputState }

func (ParameterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterType)(nil)).Elem()
}

func (o ParameterTypePtrOutput) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return o
}

func (o ParameterTypePtrOutput) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return o
}

func (o ParameterTypePtrOutput) Elem() ParameterTypeOutput {
	return o.ApplyT(func(v *ParameterType) ParameterType {
		if v != nil {
			return *v
		}
		var ret ParameterType
		return ret
	}).(ParameterTypeOutput)
}

func (o ParameterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ParameterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ParameterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ParameterTypeInput interface {
	pulumi.Input

	ToParameterTypeOutput() ParameterTypeOutput
	ToParameterTypeOutputWithContext(context.Context) ParameterTypeOutput
}

var parameterTypePtrType = reflect.TypeOf((**ParameterType)(nil)).Elem()

type ParameterTypePtrInput interface {
	pulumi.Input

	ToParameterTypePtrOutput() ParameterTypePtrOutput
	ToParameterTypePtrOutputWithContext(context.Context) ParameterTypePtrOutput
}

type parameterTypePtr string

func ParameterTypePtr(v string) ParameterTypePtrInput {
	return (*parameterTypePtr)(&v)
}

func (*parameterTypePtr) ElementType() reflect.Type {
	return parameterTypePtrType
}

func (in *parameterTypePtr) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return pulumi.ToOutput(in).(ParameterTypePtrOutput)
}

func (in *parameterTypePtr) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ParameterTypePtrOutput)
}

type PhoenixAuthenticationType string

const (
	PhoenixAuthenticationTypeAnonymous                    = PhoenixAuthenticationType("Anonymous")
	PhoenixAuthenticationTypeUsernameAndPassword          = PhoenixAuthenticationType("UsernameAndPassword")
	PhoenixAuthenticationTypeWindowsAzureHDInsightService = PhoenixAuthenticationType("WindowsAzureHDInsightService")
)

func (PhoenixAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*PhoenixAuthenticationType)(nil)).Elem()
}

func (e PhoenixAuthenticationType) ToPhoenixAuthenticationTypeOutput() PhoenixAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(PhoenixAuthenticationTypeOutput)
}

func (e PhoenixAuthenticationType) ToPhoenixAuthenticationTypeOutputWithContext(ctx context.Context) PhoenixAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PhoenixAuthenticationTypeOutput)
}

func (e PhoenixAuthenticationType) ToPhoenixAuthenticationTypePtrOutput() PhoenixAuthenticationTypePtrOutput {
	return e.ToPhoenixAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e PhoenixAuthenticationType) ToPhoenixAuthenticationTypePtrOutputWithContext(ctx context.Context) PhoenixAuthenticationTypePtrOutput {
	return PhoenixAuthenticationType(e).ToPhoenixAuthenticationTypeOutputWithContext(ctx).ToPhoenixAuthenticationTypePtrOutputWithContext(ctx)
}

func (e PhoenixAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PhoenixAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PhoenixAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PhoenixAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PhoenixAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (PhoenixAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PhoenixAuthenticationType)(nil)).Elem()
}

func (o PhoenixAuthenticationTypeOutput) ToPhoenixAuthenticationTypeOutput() PhoenixAuthenticationTypeOutput {
	return o
}

func (o PhoenixAuthenticationTypeOutput) ToPhoenixAuthenticationTypeOutputWithContext(ctx context.Context) PhoenixAuthenticationTypeOutput {
	return o
}

func (o PhoenixAuthenticationTypeOutput) ToPhoenixAuthenticationTypePtrOutput() PhoenixAuthenticationTypePtrOutput {
	return o.ToPhoenixAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o PhoenixAuthenticationTypeOutput) ToPhoenixAuthenticationTypePtrOutputWithContext(ctx context.Context) PhoenixAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PhoenixAuthenticationType) *PhoenixAuthenticationType {
		return &v
	}).(PhoenixAuthenticationTypePtrOutput)
}

func (o PhoenixAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PhoenixAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PhoenixAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PhoenixAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PhoenixAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PhoenixAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PhoenixAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (PhoenixAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhoenixAuthenticationType)(nil)).Elem()
}

func (o PhoenixAuthenticationTypePtrOutput) ToPhoenixAuthenticationTypePtrOutput() PhoenixAuthenticationTypePtrOutput {
	return o
}

func (o PhoenixAuthenticationTypePtrOutput) ToPhoenixAuthenticationTypePtrOutputWithContext(ctx context.Context) PhoenixAuthenticationTypePtrOutput {
	return o
}

func (o PhoenixAuthenticationTypePtrOutput) Elem() PhoenixAuthenticationTypeOutput {
	return o.ApplyT(func(v *PhoenixAuthenticationType) PhoenixAuthenticationType {
		if v != nil {
			return *v
		}
		var ret PhoenixAuthenticationType
		return ret
	}).(PhoenixAuthenticationTypeOutput)
}

func (o PhoenixAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PhoenixAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PhoenixAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PhoenixAuthenticationTypeInput interface {
	pulumi.Input

	ToPhoenixAuthenticationTypeOutput() PhoenixAuthenticationTypeOutput
	ToPhoenixAuthenticationTypeOutputWithContext(context.Context) PhoenixAuthenticationTypeOutput
}

var phoenixAuthenticationTypePtrType = reflect.TypeOf((**PhoenixAuthenticationType)(nil)).Elem()

type PhoenixAuthenticationTypePtrInput interface {
	pulumi.Input

	ToPhoenixAuthenticationTypePtrOutput() PhoenixAuthenticationTypePtrOutput
	ToPhoenixAuthenticationTypePtrOutputWithContext(context.Context) PhoenixAuthenticationTypePtrOutput
}

type phoenixAuthenticationTypePtr string

func PhoenixAuthenticationTypePtr(v string) PhoenixAuthenticationTypePtrInput {
	return (*phoenixAuthenticationTypePtr)(&v)
}

func (*phoenixAuthenticationTypePtr) ElementType() reflect.Type {
	return phoenixAuthenticationTypePtrType
}

func (in *phoenixAuthenticationTypePtr) ToPhoenixAuthenticationTypePtrOutput() PhoenixAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(PhoenixAuthenticationTypePtrOutput)
}

func (in *phoenixAuthenticationTypePtr) ToPhoenixAuthenticationTypePtrOutputWithContext(ctx context.Context) PhoenixAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PhoenixAuthenticationTypePtrOutput)
}

type PolybaseSettingsRejectType string

const (
	PolybaseSettingsRejectTypeValue      = PolybaseSettingsRejectType("value")
	PolybaseSettingsRejectTypePercentage = PolybaseSettingsRejectType("percentage")
)

func (PolybaseSettingsRejectType) ElementType() reflect.Type {
	return reflect.TypeOf((*PolybaseSettingsRejectType)(nil)).Elem()
}

func (e PolybaseSettingsRejectType) ToPolybaseSettingsRejectTypeOutput() PolybaseSettingsRejectTypeOutput {
	return pulumi.ToOutput(e).(PolybaseSettingsRejectTypeOutput)
}

func (e PolybaseSettingsRejectType) ToPolybaseSettingsRejectTypeOutputWithContext(ctx context.Context) PolybaseSettingsRejectTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolybaseSettingsRejectTypeOutput)
}

func (e PolybaseSettingsRejectType) ToPolybaseSettingsRejectTypePtrOutput() PolybaseSettingsRejectTypePtrOutput {
	return e.ToPolybaseSettingsRejectTypePtrOutputWithContext(context.Background())
}

func (e PolybaseSettingsRejectType) ToPolybaseSettingsRejectTypePtrOutputWithContext(ctx context.Context) PolybaseSettingsRejectTypePtrOutput {
	return PolybaseSettingsRejectType(e).ToPolybaseSettingsRejectTypeOutputWithContext(ctx).ToPolybaseSettingsRejectTypePtrOutputWithContext(ctx)
}

func (e PolybaseSettingsRejectType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolybaseSettingsRejectType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolybaseSettingsRejectType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolybaseSettingsRejectType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolybaseSettingsRejectTypeOutput struct{ *pulumi.OutputState }

func (PolybaseSettingsRejectTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolybaseSettingsRejectType)(nil)).Elem()
}

func (o PolybaseSettingsRejectTypeOutput) ToPolybaseSettingsRejectTypeOutput() PolybaseSettingsRejectTypeOutput {
	return o
}

func (o PolybaseSettingsRejectTypeOutput) ToPolybaseSettingsRejectTypeOutputWithContext(ctx context.Context) PolybaseSettingsRejectTypeOutput {
	return o
}

func (o PolybaseSettingsRejectTypeOutput) ToPolybaseSettingsRejectTypePtrOutput() PolybaseSettingsRejectTypePtrOutput {
	return o.ToPolybaseSettingsRejectTypePtrOutputWithContext(context.Background())
}

func (o PolybaseSettingsRejectTypeOutput) ToPolybaseSettingsRejectTypePtrOutputWithContext(ctx context.Context) PolybaseSettingsRejectTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolybaseSettingsRejectType) *PolybaseSettingsRejectType {
		return &v
	}).(PolybaseSettingsRejectTypePtrOutput)
}

func (o PolybaseSettingsRejectTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolybaseSettingsRejectTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolybaseSettingsRejectType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolybaseSettingsRejectTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolybaseSettingsRejectTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolybaseSettingsRejectType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolybaseSettingsRejectTypePtrOutput struct{ *pulumi.OutputState }

func (PolybaseSettingsRejectTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolybaseSettingsRejectType)(nil)).Elem()
}

func (o PolybaseSettingsRejectTypePtrOutput) ToPolybaseSettingsRejectTypePtrOutput() PolybaseSettingsRejectTypePtrOutput {
	return o
}

func (o PolybaseSettingsRejectTypePtrOutput) ToPolybaseSettingsRejectTypePtrOutputWithContext(ctx context.Context) PolybaseSettingsRejectTypePtrOutput {
	return o
}

func (o PolybaseSettingsRejectTypePtrOutput) Elem() PolybaseSettingsRejectTypeOutput {
	return o.ApplyT(func(v *PolybaseSettingsRejectType) PolybaseSettingsRejectType {
		if v != nil {
			return *v
		}
		var ret PolybaseSettingsRejectType
		return ret
	}).(PolybaseSettingsRejectTypeOutput)
}

func (o PolybaseSettingsRejectTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolybaseSettingsRejectTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolybaseSettingsRejectType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PolybaseSettingsRejectTypeInput interface {
	pulumi.Input

	ToPolybaseSettingsRejectTypeOutput() PolybaseSettingsRejectTypeOutput
	ToPolybaseSettingsRejectTypeOutputWithContext(context.Context) PolybaseSettingsRejectTypeOutput
}

var polybaseSettingsRejectTypePtrType = reflect.TypeOf((**PolybaseSettingsRejectType)(nil)).Elem()

type PolybaseSettingsRejectTypePtrInput interface {
	pulumi.Input

	ToPolybaseSettingsRejectTypePtrOutput() PolybaseSettingsRejectTypePtrOutput
	ToPolybaseSettingsRejectTypePtrOutputWithContext(context.Context) PolybaseSettingsRejectTypePtrOutput
}

type polybaseSettingsRejectTypePtr string

func PolybaseSettingsRejectTypePtr(v string) PolybaseSettingsRejectTypePtrInput {
	return (*polybaseSettingsRejectTypePtr)(&v)
}

func (*polybaseSettingsRejectTypePtr) ElementType() reflect.Type {
	return polybaseSettingsRejectTypePtrType
}

func (in *polybaseSettingsRejectTypePtr) ToPolybaseSettingsRejectTypePtrOutput() PolybaseSettingsRejectTypePtrOutput {
	return pulumi.ToOutput(in).(PolybaseSettingsRejectTypePtrOutput)
}

func (in *polybaseSettingsRejectTypePtr) ToPolybaseSettingsRejectTypePtrOutputWithContext(ctx context.Context) PolybaseSettingsRejectTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolybaseSettingsRejectTypePtrOutput)
}

type PrestoAuthenticationType string

const (
	PrestoAuthenticationTypeAnonymous = PrestoAuthenticationType("Anonymous")
	PrestoAuthenticationTypeLDAP      = PrestoAuthenticationType("LDAP")
)

func (PrestoAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*PrestoAuthenticationType)(nil)).Elem()
}

func (e PrestoAuthenticationType) ToPrestoAuthenticationTypeOutput() PrestoAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(PrestoAuthenticationTypeOutput)
}

func (e PrestoAuthenticationType) ToPrestoAuthenticationTypeOutputWithContext(ctx context.Context) PrestoAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrestoAuthenticationTypeOutput)
}

func (e PrestoAuthenticationType) ToPrestoAuthenticationTypePtrOutput() PrestoAuthenticationTypePtrOutput {
	return e.ToPrestoAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e PrestoAuthenticationType) ToPrestoAuthenticationTypePtrOutputWithContext(ctx context.Context) PrestoAuthenticationTypePtrOutput {
	return PrestoAuthenticationType(e).ToPrestoAuthenticationTypeOutputWithContext(ctx).ToPrestoAuthenticationTypePtrOutputWithContext(ctx)
}

func (e PrestoAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrestoAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrestoAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrestoAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrestoAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (PrestoAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrestoAuthenticationType)(nil)).Elem()
}

func (o PrestoAuthenticationTypeOutput) ToPrestoAuthenticationTypeOutput() PrestoAuthenticationTypeOutput {
	return o
}

func (o PrestoAuthenticationTypeOutput) ToPrestoAuthenticationTypeOutputWithContext(ctx context.Context) PrestoAuthenticationTypeOutput {
	return o
}

func (o PrestoAuthenticationTypeOutput) ToPrestoAuthenticationTypePtrOutput() PrestoAuthenticationTypePtrOutput {
	return o.ToPrestoAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o PrestoAuthenticationTypeOutput) ToPrestoAuthenticationTypePtrOutputWithContext(ctx context.Context) PrestoAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrestoAuthenticationType) *PrestoAuthenticationType {
		return &v
	}).(PrestoAuthenticationTypePtrOutput)
}

func (o PrestoAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrestoAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrestoAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrestoAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrestoAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrestoAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrestoAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (PrestoAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrestoAuthenticationType)(nil)).Elem()
}

func (o PrestoAuthenticationTypePtrOutput) ToPrestoAuthenticationTypePtrOutput() PrestoAuthenticationTypePtrOutput {
	return o
}

func (o PrestoAuthenticationTypePtrOutput) ToPrestoAuthenticationTypePtrOutputWithContext(ctx context.Context) PrestoAuthenticationTypePtrOutput {
	return o
}

func (o PrestoAuthenticationTypePtrOutput) Elem() PrestoAuthenticationTypeOutput {
	return o.ApplyT(func(v *PrestoAuthenticationType) PrestoAuthenticationType {
		if v != nil {
			return *v
		}
		var ret PrestoAuthenticationType
		return ret
	}).(PrestoAuthenticationTypeOutput)
}

func (o PrestoAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrestoAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrestoAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrestoAuthenticationTypeInput interface {
	pulumi.Input

	ToPrestoAuthenticationTypeOutput() PrestoAuthenticationTypeOutput
	ToPrestoAuthenticationTypeOutputWithContext(context.Context) PrestoAuthenticationTypeOutput
}

var prestoAuthenticationTypePtrType = reflect.TypeOf((**PrestoAuthenticationType)(nil)).Elem()

type PrestoAuthenticationTypePtrInput interface {
	pulumi.Input

	ToPrestoAuthenticationTypePtrOutput() PrestoAuthenticationTypePtrOutput
	ToPrestoAuthenticationTypePtrOutputWithContext(context.Context) PrestoAuthenticationTypePtrOutput
}

type prestoAuthenticationTypePtr string

func PrestoAuthenticationTypePtr(v string) PrestoAuthenticationTypePtrInput {
	return (*prestoAuthenticationTypePtr)(&v)
}

func (*prestoAuthenticationTypePtr) ElementType() reflect.Type {
	return prestoAuthenticationTypePtrType
}

func (in *prestoAuthenticationTypePtr) ToPrestoAuthenticationTypePtrOutput() PrestoAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(PrestoAuthenticationTypePtrOutput)
}

func (in *prestoAuthenticationTypePtr) ToPrestoAuthenticationTypePtrOutputWithContext(ctx context.Context) PrestoAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrestoAuthenticationTypePtrOutput)
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return e.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return PublicNetworkAccess(e).ToPublicNetworkAccessOutputWithContext(ctx).ToPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccess) *PublicNetworkAccess {
		return &v
	}).(PublicNetworkAccessPtrOutput)
}

func (o PublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) Elem() PublicNetworkAccessOutput {
	return o.ApplyT(func(v *PublicNetworkAccess) PublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccess
		return ret
	}).(PublicNetworkAccessOutput)
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessInput interface {
	pulumi.Input

	ToPublicNetworkAccessOutput() PublicNetworkAccessOutput
	ToPublicNetworkAccessOutputWithContext(context.Context) PublicNetworkAccessOutput
}

var publicNetworkAccessPtrType = reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()

type PublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput
	ToPublicNetworkAccessPtrOutputWithContext(context.Context) PublicNetworkAccessPtrOutput
}

type publicNetworkAccessPtr string

func PublicNetworkAccessPtr(v string) PublicNetworkAccessPtrInput {
	return (*publicNetworkAccessPtr)(&v)
}

func (*publicNetworkAccessPtr) ElementType() reflect.Type {
	return publicNetworkAccessPtrType
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessPtrOutput)
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessPtrOutput)
}

type RecurrenceFrequency string

const (
	RecurrenceFrequencyNotSpecified = RecurrenceFrequency("NotSpecified")
	RecurrenceFrequencyMinute       = RecurrenceFrequency("Minute")
	RecurrenceFrequencyHour         = RecurrenceFrequency("Hour")
	RecurrenceFrequencyDay          = RecurrenceFrequency("Day")
	RecurrenceFrequencyWeek         = RecurrenceFrequency("Week")
	RecurrenceFrequencyMonth        = RecurrenceFrequency("Month")
	RecurrenceFrequencyYear         = RecurrenceFrequency("Year")
)

func (RecurrenceFrequency) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceFrequency)(nil)).Elem()
}

func (e RecurrenceFrequency) ToRecurrenceFrequencyOutput() RecurrenceFrequencyOutput {
	return pulumi.ToOutput(e).(RecurrenceFrequencyOutput)
}

func (e RecurrenceFrequency) ToRecurrenceFrequencyOutputWithContext(ctx context.Context) RecurrenceFrequencyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecurrenceFrequencyOutput)
}

func (e RecurrenceFrequency) ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput {
	return e.ToRecurrenceFrequencyPtrOutputWithContext(context.Background())
}

func (e RecurrenceFrequency) ToRecurrenceFrequencyPtrOutputWithContext(ctx context.Context) RecurrenceFrequencyPtrOutput {
	return RecurrenceFrequency(e).ToRecurrenceFrequencyOutputWithContext(ctx).ToRecurrenceFrequencyPtrOutputWithContext(ctx)
}

func (e RecurrenceFrequency) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceFrequency) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceFrequency) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecurrenceFrequency) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecurrenceFrequencyOutput struct{ *pulumi.OutputState }

func (RecurrenceFrequencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceFrequency)(nil)).Elem()
}

func (o RecurrenceFrequencyOutput) ToRecurrenceFrequencyOutput() RecurrenceFrequencyOutput {
	return o
}

func (o RecurrenceFrequencyOutput) ToRecurrenceFrequencyOutputWithContext(ctx context.Context) RecurrenceFrequencyOutput {
	return o
}

func (o RecurrenceFrequencyOutput) ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput {
	return o.ToRecurrenceFrequencyPtrOutputWithContext(context.Background())
}

func (o RecurrenceFrequencyOutput) ToRecurrenceFrequencyPtrOutputWithContext(ctx context.Context) RecurrenceFrequencyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrenceFrequency) *RecurrenceFrequency {
		return &v
	}).(RecurrenceFrequencyPtrOutput)
}

func (o RecurrenceFrequencyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecurrenceFrequencyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceFrequency) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecurrenceFrequencyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceFrequencyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceFrequency) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecurrenceFrequencyPtrOutput struct{ *pulumi.OutputState }

func (RecurrenceFrequencyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceFrequency)(nil)).Elem()
}

func (o RecurrenceFrequencyPtrOutput) ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput {
	return o
}

func (o RecurrenceFrequencyPtrOutput) ToRecurrenceFrequencyPtrOutputWithContext(ctx context.Context) RecurrenceFrequencyPtrOutput {
	return o
}

func (o RecurrenceFrequencyPtrOutput) Elem() RecurrenceFrequencyOutput {
	return o.ApplyT(func(v *RecurrenceFrequency) RecurrenceFrequency {
		if v != nil {
			return *v
		}
		var ret RecurrenceFrequency
		return ret
	}).(RecurrenceFrequencyOutput)
}

func (o RecurrenceFrequencyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceFrequencyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecurrenceFrequency) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RecurrenceFrequencyInput interface {
	pulumi.Input

	ToRecurrenceFrequencyOutput() RecurrenceFrequencyOutput
	ToRecurrenceFrequencyOutputWithContext(context.Context) RecurrenceFrequencyOutput
}

var recurrenceFrequencyPtrType = reflect.TypeOf((**RecurrenceFrequency)(nil)).Elem()

type RecurrenceFrequencyPtrInput interface {
	pulumi.Input

	ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput
	ToRecurrenceFrequencyPtrOutputWithContext(context.Context) RecurrenceFrequencyPtrOutput
}

type recurrenceFrequencyPtr string

func RecurrenceFrequencyPtr(v string) RecurrenceFrequencyPtrInput {
	return (*recurrenceFrequencyPtr)(&v)
}

func (*recurrenceFrequencyPtr) ElementType() reflect.Type {
	return recurrenceFrequencyPtrType
}

func (in *recurrenceFrequencyPtr) ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput {
	return pulumi.ToOutput(in).(RecurrenceFrequencyPtrOutput)
}

func (in *recurrenceFrequencyPtr) ToRecurrenceFrequencyPtrOutputWithContext(ctx context.Context) RecurrenceFrequencyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecurrenceFrequencyPtrOutput)
}

type RestServiceAuthenticationType string

const (
	RestServiceAuthenticationTypeAnonymous              = RestServiceAuthenticationType("Anonymous")
	RestServiceAuthenticationTypeBasic                  = RestServiceAuthenticationType("Basic")
	RestServiceAuthenticationTypeAadServicePrincipal    = RestServiceAuthenticationType("AadServicePrincipal")
	RestServiceAuthenticationTypeManagedServiceIdentity = RestServiceAuthenticationType("ManagedServiceIdentity")
)

func (RestServiceAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*RestServiceAuthenticationType)(nil)).Elem()
}

func (e RestServiceAuthenticationType) ToRestServiceAuthenticationTypeOutput() RestServiceAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(RestServiceAuthenticationTypeOutput)
}

func (e RestServiceAuthenticationType) ToRestServiceAuthenticationTypeOutputWithContext(ctx context.Context) RestServiceAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RestServiceAuthenticationTypeOutput)
}

func (e RestServiceAuthenticationType) ToRestServiceAuthenticationTypePtrOutput() RestServiceAuthenticationTypePtrOutput {
	return e.ToRestServiceAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e RestServiceAuthenticationType) ToRestServiceAuthenticationTypePtrOutputWithContext(ctx context.Context) RestServiceAuthenticationTypePtrOutput {
	return RestServiceAuthenticationType(e).ToRestServiceAuthenticationTypeOutputWithContext(ctx).ToRestServiceAuthenticationTypePtrOutputWithContext(ctx)
}

func (e RestServiceAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestServiceAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestServiceAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RestServiceAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RestServiceAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (RestServiceAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestServiceAuthenticationType)(nil)).Elem()
}

func (o RestServiceAuthenticationTypeOutput) ToRestServiceAuthenticationTypeOutput() RestServiceAuthenticationTypeOutput {
	return o
}

func (o RestServiceAuthenticationTypeOutput) ToRestServiceAuthenticationTypeOutputWithContext(ctx context.Context) RestServiceAuthenticationTypeOutput {
	return o
}

func (o RestServiceAuthenticationTypeOutput) ToRestServiceAuthenticationTypePtrOutput() RestServiceAuthenticationTypePtrOutput {
	return o.ToRestServiceAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o RestServiceAuthenticationTypeOutput) ToRestServiceAuthenticationTypePtrOutputWithContext(ctx context.Context) RestServiceAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestServiceAuthenticationType) *RestServiceAuthenticationType {
		return &v
	}).(RestServiceAuthenticationTypePtrOutput)
}

func (o RestServiceAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RestServiceAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestServiceAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RestServiceAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestServiceAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestServiceAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RestServiceAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (RestServiceAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestServiceAuthenticationType)(nil)).Elem()
}

func (o RestServiceAuthenticationTypePtrOutput) ToRestServiceAuthenticationTypePtrOutput() RestServiceAuthenticationTypePtrOutput {
	return o
}

func (o RestServiceAuthenticationTypePtrOutput) ToRestServiceAuthenticationTypePtrOutputWithContext(ctx context.Context) RestServiceAuthenticationTypePtrOutput {
	return o
}

func (o RestServiceAuthenticationTypePtrOutput) Elem() RestServiceAuthenticationTypeOutput {
	return o.ApplyT(func(v *RestServiceAuthenticationType) RestServiceAuthenticationType {
		if v != nil {
			return *v
		}
		var ret RestServiceAuthenticationType
		return ret
	}).(RestServiceAuthenticationTypeOutput)
}

func (o RestServiceAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestServiceAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RestServiceAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RestServiceAuthenticationTypeInput interface {
	pulumi.Input

	ToRestServiceAuthenticationTypeOutput() RestServiceAuthenticationTypeOutput
	ToRestServiceAuthenticationTypeOutputWithContext(context.Context) RestServiceAuthenticationTypeOutput
}

var restServiceAuthenticationTypePtrType = reflect.TypeOf((**RestServiceAuthenticationType)(nil)).Elem()

type RestServiceAuthenticationTypePtrInput interface {
	pulumi.Input

	ToRestServiceAuthenticationTypePtrOutput() RestServiceAuthenticationTypePtrOutput
	ToRestServiceAuthenticationTypePtrOutputWithContext(context.Context) RestServiceAuthenticationTypePtrOutput
}

type restServiceAuthenticationTypePtr string

func RestServiceAuthenticationTypePtr(v string) RestServiceAuthenticationTypePtrInput {
	return (*restServiceAuthenticationTypePtr)(&v)
}

func (*restServiceAuthenticationTypePtr) ElementType() reflect.Type {
	return restServiceAuthenticationTypePtrType
}

func (in *restServiceAuthenticationTypePtr) ToRestServiceAuthenticationTypePtrOutput() RestServiceAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(RestServiceAuthenticationTypePtrOutput)
}

func (in *restServiceAuthenticationTypePtr) ToRestServiceAuthenticationTypePtrOutputWithContext(ctx context.Context) RestServiceAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RestServiceAuthenticationTypePtrOutput)
}

type SalesforceSinkWriteBehavior string

const (
	SalesforceSinkWriteBehaviorInsert = SalesforceSinkWriteBehavior("Insert")
	SalesforceSinkWriteBehaviorUpsert = SalesforceSinkWriteBehavior("Upsert")
)

func (SalesforceSinkWriteBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*SalesforceSinkWriteBehavior)(nil)).Elem()
}

func (e SalesforceSinkWriteBehavior) ToSalesforceSinkWriteBehaviorOutput() SalesforceSinkWriteBehaviorOutput {
	return pulumi.ToOutput(e).(SalesforceSinkWriteBehaviorOutput)
}

func (e SalesforceSinkWriteBehavior) ToSalesforceSinkWriteBehaviorOutputWithContext(ctx context.Context) SalesforceSinkWriteBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SalesforceSinkWriteBehaviorOutput)
}

func (e SalesforceSinkWriteBehavior) ToSalesforceSinkWriteBehaviorPtrOutput() SalesforceSinkWriteBehaviorPtrOutput {
	return e.ToSalesforceSinkWriteBehaviorPtrOutputWithContext(context.Background())
}

func (e SalesforceSinkWriteBehavior) ToSalesforceSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) SalesforceSinkWriteBehaviorPtrOutput {
	return SalesforceSinkWriteBehavior(e).ToSalesforceSinkWriteBehaviorOutputWithContext(ctx).ToSalesforceSinkWriteBehaviorPtrOutputWithContext(ctx)
}

func (e SalesforceSinkWriteBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SalesforceSinkWriteBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SalesforceSinkWriteBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SalesforceSinkWriteBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SalesforceSinkWriteBehaviorOutput struct{ *pulumi.OutputState }

func (SalesforceSinkWriteBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SalesforceSinkWriteBehavior)(nil)).Elem()
}

func (o SalesforceSinkWriteBehaviorOutput) ToSalesforceSinkWriteBehaviorOutput() SalesforceSinkWriteBehaviorOutput {
	return o
}

func (o SalesforceSinkWriteBehaviorOutput) ToSalesforceSinkWriteBehaviorOutputWithContext(ctx context.Context) SalesforceSinkWriteBehaviorOutput {
	return o
}

func (o SalesforceSinkWriteBehaviorOutput) ToSalesforceSinkWriteBehaviorPtrOutput() SalesforceSinkWriteBehaviorPtrOutput {
	return o.ToSalesforceSinkWriteBehaviorPtrOutputWithContext(context.Background())
}

func (o SalesforceSinkWriteBehaviorOutput) ToSalesforceSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) SalesforceSinkWriteBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SalesforceSinkWriteBehavior) *SalesforceSinkWriteBehavior {
		return &v
	}).(SalesforceSinkWriteBehaviorPtrOutput)
}

func (o SalesforceSinkWriteBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SalesforceSinkWriteBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SalesforceSinkWriteBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SalesforceSinkWriteBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SalesforceSinkWriteBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SalesforceSinkWriteBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SalesforceSinkWriteBehaviorPtrOutput struct{ *pulumi.OutputState }

func (SalesforceSinkWriteBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SalesforceSinkWriteBehavior)(nil)).Elem()
}

func (o SalesforceSinkWriteBehaviorPtrOutput) ToSalesforceSinkWriteBehaviorPtrOutput() SalesforceSinkWriteBehaviorPtrOutput {
	return o
}

func (o SalesforceSinkWriteBehaviorPtrOutput) ToSalesforceSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) SalesforceSinkWriteBehaviorPtrOutput {
	return o
}

func (o SalesforceSinkWriteBehaviorPtrOutput) Elem() SalesforceSinkWriteBehaviorOutput {
	return o.ApplyT(func(v *SalesforceSinkWriteBehavior) SalesforceSinkWriteBehavior {
		if v != nil {
			return *v
		}
		var ret SalesforceSinkWriteBehavior
		return ret
	}).(SalesforceSinkWriteBehaviorOutput)
}

func (o SalesforceSinkWriteBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SalesforceSinkWriteBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SalesforceSinkWriteBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SalesforceSinkWriteBehaviorInput interface {
	pulumi.Input

	ToSalesforceSinkWriteBehaviorOutput() SalesforceSinkWriteBehaviorOutput
	ToSalesforceSinkWriteBehaviorOutputWithContext(context.Context) SalesforceSinkWriteBehaviorOutput
}

var salesforceSinkWriteBehaviorPtrType = reflect.TypeOf((**SalesforceSinkWriteBehavior)(nil)).Elem()

type SalesforceSinkWriteBehaviorPtrInput interface {
	pulumi.Input

	ToSalesforceSinkWriteBehaviorPtrOutput() SalesforceSinkWriteBehaviorPtrOutput
	ToSalesforceSinkWriteBehaviorPtrOutputWithContext(context.Context) SalesforceSinkWriteBehaviorPtrOutput
}

type salesforceSinkWriteBehaviorPtr string

func SalesforceSinkWriteBehaviorPtr(v string) SalesforceSinkWriteBehaviorPtrInput {
	return (*salesforceSinkWriteBehaviorPtr)(&v)
}

func (*salesforceSinkWriteBehaviorPtr) ElementType() reflect.Type {
	return salesforceSinkWriteBehaviorPtrType
}

func (in *salesforceSinkWriteBehaviorPtr) ToSalesforceSinkWriteBehaviorPtrOutput() SalesforceSinkWriteBehaviorPtrOutput {
	return pulumi.ToOutput(in).(SalesforceSinkWriteBehaviorPtrOutput)
}

func (in *salesforceSinkWriteBehaviorPtr) ToSalesforceSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) SalesforceSinkWriteBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SalesforceSinkWriteBehaviorPtrOutput)
}

type SalesforceSourceReadBehavior string

const (
	SalesforceSourceReadBehaviorQuery    = SalesforceSourceReadBehavior("Query")
	SalesforceSourceReadBehaviorQueryAll = SalesforceSourceReadBehavior("QueryAll")
)

func (SalesforceSourceReadBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*SalesforceSourceReadBehavior)(nil)).Elem()
}

func (e SalesforceSourceReadBehavior) ToSalesforceSourceReadBehaviorOutput() SalesforceSourceReadBehaviorOutput {
	return pulumi.ToOutput(e).(SalesforceSourceReadBehaviorOutput)
}

func (e SalesforceSourceReadBehavior) ToSalesforceSourceReadBehaviorOutputWithContext(ctx context.Context) SalesforceSourceReadBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SalesforceSourceReadBehaviorOutput)
}

func (e SalesforceSourceReadBehavior) ToSalesforceSourceReadBehaviorPtrOutput() SalesforceSourceReadBehaviorPtrOutput {
	return e.ToSalesforceSourceReadBehaviorPtrOutputWithContext(context.Background())
}

func (e SalesforceSourceReadBehavior) ToSalesforceSourceReadBehaviorPtrOutputWithContext(ctx context.Context) SalesforceSourceReadBehaviorPtrOutput {
	return SalesforceSourceReadBehavior(e).ToSalesforceSourceReadBehaviorOutputWithContext(ctx).ToSalesforceSourceReadBehaviorPtrOutputWithContext(ctx)
}

func (e SalesforceSourceReadBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SalesforceSourceReadBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SalesforceSourceReadBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SalesforceSourceReadBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SalesforceSourceReadBehaviorOutput struct{ *pulumi.OutputState }

func (SalesforceSourceReadBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SalesforceSourceReadBehavior)(nil)).Elem()
}

func (o SalesforceSourceReadBehaviorOutput) ToSalesforceSourceReadBehaviorOutput() SalesforceSourceReadBehaviorOutput {
	return o
}

func (o SalesforceSourceReadBehaviorOutput) ToSalesforceSourceReadBehaviorOutputWithContext(ctx context.Context) SalesforceSourceReadBehaviorOutput {
	return o
}

func (o SalesforceSourceReadBehaviorOutput) ToSalesforceSourceReadBehaviorPtrOutput() SalesforceSourceReadBehaviorPtrOutput {
	return o.ToSalesforceSourceReadBehaviorPtrOutputWithContext(context.Background())
}

func (o SalesforceSourceReadBehaviorOutput) ToSalesforceSourceReadBehaviorPtrOutputWithContext(ctx context.Context) SalesforceSourceReadBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SalesforceSourceReadBehavior) *SalesforceSourceReadBehavior {
		return &v
	}).(SalesforceSourceReadBehaviorPtrOutput)
}

func (o SalesforceSourceReadBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SalesforceSourceReadBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SalesforceSourceReadBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SalesforceSourceReadBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SalesforceSourceReadBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SalesforceSourceReadBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SalesforceSourceReadBehaviorPtrOutput struct{ *pulumi.OutputState }

func (SalesforceSourceReadBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SalesforceSourceReadBehavior)(nil)).Elem()
}

func (o SalesforceSourceReadBehaviorPtrOutput) ToSalesforceSourceReadBehaviorPtrOutput() SalesforceSourceReadBehaviorPtrOutput {
	return o
}

func (o SalesforceSourceReadBehaviorPtrOutput) ToSalesforceSourceReadBehaviorPtrOutputWithContext(ctx context.Context) SalesforceSourceReadBehaviorPtrOutput {
	return o
}

func (o SalesforceSourceReadBehaviorPtrOutput) Elem() SalesforceSourceReadBehaviorOutput {
	return o.ApplyT(func(v *SalesforceSourceReadBehavior) SalesforceSourceReadBehavior {
		if v != nil {
			return *v
		}
		var ret SalesforceSourceReadBehavior
		return ret
	}).(SalesforceSourceReadBehaviorOutput)
}

func (o SalesforceSourceReadBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SalesforceSourceReadBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SalesforceSourceReadBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SalesforceSourceReadBehaviorInput interface {
	pulumi.Input

	ToSalesforceSourceReadBehaviorOutput() SalesforceSourceReadBehaviorOutput
	ToSalesforceSourceReadBehaviorOutputWithContext(context.Context) SalesforceSourceReadBehaviorOutput
}

var salesforceSourceReadBehaviorPtrType = reflect.TypeOf((**SalesforceSourceReadBehavior)(nil)).Elem()

type SalesforceSourceReadBehaviorPtrInput interface {
	pulumi.Input

	ToSalesforceSourceReadBehaviorPtrOutput() SalesforceSourceReadBehaviorPtrOutput
	ToSalesforceSourceReadBehaviorPtrOutputWithContext(context.Context) SalesforceSourceReadBehaviorPtrOutput
}

type salesforceSourceReadBehaviorPtr string

func SalesforceSourceReadBehaviorPtr(v string) SalesforceSourceReadBehaviorPtrInput {
	return (*salesforceSourceReadBehaviorPtr)(&v)
}

func (*salesforceSourceReadBehaviorPtr) ElementType() reflect.Type {
	return salesforceSourceReadBehaviorPtrType
}

func (in *salesforceSourceReadBehaviorPtr) ToSalesforceSourceReadBehaviorPtrOutput() SalesforceSourceReadBehaviorPtrOutput {
	return pulumi.ToOutput(in).(SalesforceSourceReadBehaviorPtrOutput)
}

func (in *salesforceSourceReadBehaviorPtr) ToSalesforceSourceReadBehaviorPtrOutputWithContext(ctx context.Context) SalesforceSourceReadBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SalesforceSourceReadBehaviorPtrOutput)
}

type SapCloudForCustomerSinkWriteBehavior string

const (
	SapCloudForCustomerSinkWriteBehaviorInsert = SapCloudForCustomerSinkWriteBehavior("Insert")
	SapCloudForCustomerSinkWriteBehaviorUpdate = SapCloudForCustomerSinkWriteBehavior("Update")
)

func (SapCloudForCustomerSinkWriteBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*SapCloudForCustomerSinkWriteBehavior)(nil)).Elem()
}

func (e SapCloudForCustomerSinkWriteBehavior) ToSapCloudForCustomerSinkWriteBehaviorOutput() SapCloudForCustomerSinkWriteBehaviorOutput {
	return pulumi.ToOutput(e).(SapCloudForCustomerSinkWriteBehaviorOutput)
}

func (e SapCloudForCustomerSinkWriteBehavior) ToSapCloudForCustomerSinkWriteBehaviorOutputWithContext(ctx context.Context) SapCloudForCustomerSinkWriteBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SapCloudForCustomerSinkWriteBehaviorOutput)
}

func (e SapCloudForCustomerSinkWriteBehavior) ToSapCloudForCustomerSinkWriteBehaviorPtrOutput() SapCloudForCustomerSinkWriteBehaviorPtrOutput {
	return e.ToSapCloudForCustomerSinkWriteBehaviorPtrOutputWithContext(context.Background())
}

func (e SapCloudForCustomerSinkWriteBehavior) ToSapCloudForCustomerSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) SapCloudForCustomerSinkWriteBehaviorPtrOutput {
	return SapCloudForCustomerSinkWriteBehavior(e).ToSapCloudForCustomerSinkWriteBehaviorOutputWithContext(ctx).ToSapCloudForCustomerSinkWriteBehaviorPtrOutputWithContext(ctx)
}

func (e SapCloudForCustomerSinkWriteBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SapCloudForCustomerSinkWriteBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SapCloudForCustomerSinkWriteBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SapCloudForCustomerSinkWriteBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SapCloudForCustomerSinkWriteBehaviorOutput struct{ *pulumi.OutputState }

func (SapCloudForCustomerSinkWriteBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SapCloudForCustomerSinkWriteBehavior)(nil)).Elem()
}

func (o SapCloudForCustomerSinkWriteBehaviorOutput) ToSapCloudForCustomerSinkWriteBehaviorOutput() SapCloudForCustomerSinkWriteBehaviorOutput {
	return o
}

func (o SapCloudForCustomerSinkWriteBehaviorOutput) ToSapCloudForCustomerSinkWriteBehaviorOutputWithContext(ctx context.Context) SapCloudForCustomerSinkWriteBehaviorOutput {
	return o
}

func (o SapCloudForCustomerSinkWriteBehaviorOutput) ToSapCloudForCustomerSinkWriteBehaviorPtrOutput() SapCloudForCustomerSinkWriteBehaviorPtrOutput {
	return o.ToSapCloudForCustomerSinkWriteBehaviorPtrOutputWithContext(context.Background())
}

func (o SapCloudForCustomerSinkWriteBehaviorOutput) ToSapCloudForCustomerSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) SapCloudForCustomerSinkWriteBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SapCloudForCustomerSinkWriteBehavior) *SapCloudForCustomerSinkWriteBehavior {
		return &v
	}).(SapCloudForCustomerSinkWriteBehaviorPtrOutput)
}

func (o SapCloudForCustomerSinkWriteBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SapCloudForCustomerSinkWriteBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SapCloudForCustomerSinkWriteBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SapCloudForCustomerSinkWriteBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SapCloudForCustomerSinkWriteBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SapCloudForCustomerSinkWriteBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SapCloudForCustomerSinkWriteBehaviorPtrOutput struct{ *pulumi.OutputState }

func (SapCloudForCustomerSinkWriteBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SapCloudForCustomerSinkWriteBehavior)(nil)).Elem()
}

func (o SapCloudForCustomerSinkWriteBehaviorPtrOutput) ToSapCloudForCustomerSinkWriteBehaviorPtrOutput() SapCloudForCustomerSinkWriteBehaviorPtrOutput {
	return o
}

func (o SapCloudForCustomerSinkWriteBehaviorPtrOutput) ToSapCloudForCustomerSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) SapCloudForCustomerSinkWriteBehaviorPtrOutput {
	return o
}

func (o SapCloudForCustomerSinkWriteBehaviorPtrOutput) Elem() SapCloudForCustomerSinkWriteBehaviorOutput {
	return o.ApplyT(func(v *SapCloudForCustomerSinkWriteBehavior) SapCloudForCustomerSinkWriteBehavior {
		if v != nil {
			return *v
		}
		var ret SapCloudForCustomerSinkWriteBehavior
		return ret
	}).(SapCloudForCustomerSinkWriteBehaviorOutput)
}

func (o SapCloudForCustomerSinkWriteBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SapCloudForCustomerSinkWriteBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SapCloudForCustomerSinkWriteBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SapCloudForCustomerSinkWriteBehaviorInput interface {
	pulumi.Input

	ToSapCloudForCustomerSinkWriteBehaviorOutput() SapCloudForCustomerSinkWriteBehaviorOutput
	ToSapCloudForCustomerSinkWriteBehaviorOutputWithContext(context.Context) SapCloudForCustomerSinkWriteBehaviorOutput
}

var sapCloudForCustomerSinkWriteBehaviorPtrType = reflect.TypeOf((**SapCloudForCustomerSinkWriteBehavior)(nil)).Elem()

type SapCloudForCustomerSinkWriteBehaviorPtrInput interface {
	pulumi.Input

	ToSapCloudForCustomerSinkWriteBehaviorPtrOutput() SapCloudForCustomerSinkWriteBehaviorPtrOutput
	ToSapCloudForCustomerSinkWriteBehaviorPtrOutputWithContext(context.Context) SapCloudForCustomerSinkWriteBehaviorPtrOutput
}

type sapCloudForCustomerSinkWriteBehaviorPtr string

func SapCloudForCustomerSinkWriteBehaviorPtr(v string) SapCloudForCustomerSinkWriteBehaviorPtrInput {
	return (*sapCloudForCustomerSinkWriteBehaviorPtr)(&v)
}

func (*sapCloudForCustomerSinkWriteBehaviorPtr) ElementType() reflect.Type {
	return sapCloudForCustomerSinkWriteBehaviorPtrType
}

func (in *sapCloudForCustomerSinkWriteBehaviorPtr) ToSapCloudForCustomerSinkWriteBehaviorPtrOutput() SapCloudForCustomerSinkWriteBehaviorPtrOutput {
	return pulumi.ToOutput(in).(SapCloudForCustomerSinkWriteBehaviorPtrOutput)
}

func (in *sapCloudForCustomerSinkWriteBehaviorPtr) ToSapCloudForCustomerSinkWriteBehaviorPtrOutputWithContext(ctx context.Context) SapCloudForCustomerSinkWriteBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SapCloudForCustomerSinkWriteBehaviorPtrOutput)
}

type SapHanaAuthenticationType string

const (
	SapHanaAuthenticationTypeBasic   = SapHanaAuthenticationType("Basic")
	SapHanaAuthenticationTypeWindows = SapHanaAuthenticationType("Windows")
)

func (SapHanaAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*SapHanaAuthenticationType)(nil)).Elem()
}

func (e SapHanaAuthenticationType) ToSapHanaAuthenticationTypeOutput() SapHanaAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(SapHanaAuthenticationTypeOutput)
}

func (e SapHanaAuthenticationType) ToSapHanaAuthenticationTypeOutputWithContext(ctx context.Context) SapHanaAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SapHanaAuthenticationTypeOutput)
}

func (e SapHanaAuthenticationType) ToSapHanaAuthenticationTypePtrOutput() SapHanaAuthenticationTypePtrOutput {
	return e.ToSapHanaAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e SapHanaAuthenticationType) ToSapHanaAuthenticationTypePtrOutputWithContext(ctx context.Context) SapHanaAuthenticationTypePtrOutput {
	return SapHanaAuthenticationType(e).ToSapHanaAuthenticationTypeOutputWithContext(ctx).ToSapHanaAuthenticationTypePtrOutputWithContext(ctx)
}

func (e SapHanaAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SapHanaAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SapHanaAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SapHanaAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SapHanaAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (SapHanaAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SapHanaAuthenticationType)(nil)).Elem()
}

func (o SapHanaAuthenticationTypeOutput) ToSapHanaAuthenticationTypeOutput() SapHanaAuthenticationTypeOutput {
	return o
}

func (o SapHanaAuthenticationTypeOutput) ToSapHanaAuthenticationTypeOutputWithContext(ctx context.Context) SapHanaAuthenticationTypeOutput {
	return o
}

func (o SapHanaAuthenticationTypeOutput) ToSapHanaAuthenticationTypePtrOutput() SapHanaAuthenticationTypePtrOutput {
	return o.ToSapHanaAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o SapHanaAuthenticationTypeOutput) ToSapHanaAuthenticationTypePtrOutputWithContext(ctx context.Context) SapHanaAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SapHanaAuthenticationType) *SapHanaAuthenticationType {
		return &v
	}).(SapHanaAuthenticationTypePtrOutput)
}

func (o SapHanaAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SapHanaAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SapHanaAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SapHanaAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SapHanaAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SapHanaAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SapHanaAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (SapHanaAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SapHanaAuthenticationType)(nil)).Elem()
}

func (o SapHanaAuthenticationTypePtrOutput) ToSapHanaAuthenticationTypePtrOutput() SapHanaAuthenticationTypePtrOutput {
	return o
}

func (o SapHanaAuthenticationTypePtrOutput) ToSapHanaAuthenticationTypePtrOutputWithContext(ctx context.Context) SapHanaAuthenticationTypePtrOutput {
	return o
}

func (o SapHanaAuthenticationTypePtrOutput) Elem() SapHanaAuthenticationTypeOutput {
	return o.ApplyT(func(v *SapHanaAuthenticationType) SapHanaAuthenticationType {
		if v != nil {
			return *v
		}
		var ret SapHanaAuthenticationType
		return ret
	}).(SapHanaAuthenticationTypeOutput)
}

func (o SapHanaAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SapHanaAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SapHanaAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SapHanaAuthenticationTypeInput interface {
	pulumi.Input

	ToSapHanaAuthenticationTypeOutput() SapHanaAuthenticationTypeOutput
	ToSapHanaAuthenticationTypeOutputWithContext(context.Context) SapHanaAuthenticationTypeOutput
}

var sapHanaAuthenticationTypePtrType = reflect.TypeOf((**SapHanaAuthenticationType)(nil)).Elem()

type SapHanaAuthenticationTypePtrInput interface {
	pulumi.Input

	ToSapHanaAuthenticationTypePtrOutput() SapHanaAuthenticationTypePtrOutput
	ToSapHanaAuthenticationTypePtrOutputWithContext(context.Context) SapHanaAuthenticationTypePtrOutput
}

type sapHanaAuthenticationTypePtr string

func SapHanaAuthenticationTypePtr(v string) SapHanaAuthenticationTypePtrInput {
	return (*sapHanaAuthenticationTypePtr)(&v)
}

func (*sapHanaAuthenticationTypePtr) ElementType() reflect.Type {
	return sapHanaAuthenticationTypePtrType
}

func (in *sapHanaAuthenticationTypePtr) ToSapHanaAuthenticationTypePtrOutput() SapHanaAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(SapHanaAuthenticationTypePtrOutput)
}

func (in *sapHanaAuthenticationTypePtr) ToSapHanaAuthenticationTypePtrOutputWithContext(ctx context.Context) SapHanaAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SapHanaAuthenticationTypePtrOutput)
}

type ServiceNowAuthenticationType string

const (
	ServiceNowAuthenticationTypeBasic  = ServiceNowAuthenticationType("Basic")
	ServiceNowAuthenticationTypeOAuth2 = ServiceNowAuthenticationType("OAuth2")
)

func (ServiceNowAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceNowAuthenticationType)(nil)).Elem()
}

func (e ServiceNowAuthenticationType) ToServiceNowAuthenticationTypeOutput() ServiceNowAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(ServiceNowAuthenticationTypeOutput)
}

func (e ServiceNowAuthenticationType) ToServiceNowAuthenticationTypeOutputWithContext(ctx context.Context) ServiceNowAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceNowAuthenticationTypeOutput)
}

func (e ServiceNowAuthenticationType) ToServiceNowAuthenticationTypePtrOutput() ServiceNowAuthenticationTypePtrOutput {
	return e.ToServiceNowAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e ServiceNowAuthenticationType) ToServiceNowAuthenticationTypePtrOutputWithContext(ctx context.Context) ServiceNowAuthenticationTypePtrOutput {
	return ServiceNowAuthenticationType(e).ToServiceNowAuthenticationTypeOutputWithContext(ctx).ToServiceNowAuthenticationTypePtrOutputWithContext(ctx)
}

func (e ServiceNowAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceNowAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceNowAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceNowAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceNowAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (ServiceNowAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceNowAuthenticationType)(nil)).Elem()
}

func (o ServiceNowAuthenticationTypeOutput) ToServiceNowAuthenticationTypeOutput() ServiceNowAuthenticationTypeOutput {
	return o
}

func (o ServiceNowAuthenticationTypeOutput) ToServiceNowAuthenticationTypeOutputWithContext(ctx context.Context) ServiceNowAuthenticationTypeOutput {
	return o
}

func (o ServiceNowAuthenticationTypeOutput) ToServiceNowAuthenticationTypePtrOutput() ServiceNowAuthenticationTypePtrOutput {
	return o.ToServiceNowAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o ServiceNowAuthenticationTypeOutput) ToServiceNowAuthenticationTypePtrOutputWithContext(ctx context.Context) ServiceNowAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceNowAuthenticationType) *ServiceNowAuthenticationType {
		return &v
	}).(ServiceNowAuthenticationTypePtrOutput)
}

func (o ServiceNowAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceNowAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceNowAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceNowAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceNowAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceNowAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceNowAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (ServiceNowAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceNowAuthenticationType)(nil)).Elem()
}

func (o ServiceNowAuthenticationTypePtrOutput) ToServiceNowAuthenticationTypePtrOutput() ServiceNowAuthenticationTypePtrOutput {
	return o
}

func (o ServiceNowAuthenticationTypePtrOutput) ToServiceNowAuthenticationTypePtrOutputWithContext(ctx context.Context) ServiceNowAuthenticationTypePtrOutput {
	return o
}

func (o ServiceNowAuthenticationTypePtrOutput) Elem() ServiceNowAuthenticationTypeOutput {
	return o.ApplyT(func(v *ServiceNowAuthenticationType) ServiceNowAuthenticationType {
		if v != nil {
			return *v
		}
		var ret ServiceNowAuthenticationType
		return ret
	}).(ServiceNowAuthenticationTypeOutput)
}

func (o ServiceNowAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceNowAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceNowAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceNowAuthenticationTypeInput interface {
	pulumi.Input

	ToServiceNowAuthenticationTypeOutput() ServiceNowAuthenticationTypeOutput
	ToServiceNowAuthenticationTypeOutputWithContext(context.Context) ServiceNowAuthenticationTypeOutput
}

var serviceNowAuthenticationTypePtrType = reflect.TypeOf((**ServiceNowAuthenticationType)(nil)).Elem()

type ServiceNowAuthenticationTypePtrInput interface {
	pulumi.Input

	ToServiceNowAuthenticationTypePtrOutput() ServiceNowAuthenticationTypePtrOutput
	ToServiceNowAuthenticationTypePtrOutputWithContext(context.Context) ServiceNowAuthenticationTypePtrOutput
}

type serviceNowAuthenticationTypePtr string

func ServiceNowAuthenticationTypePtr(v string) ServiceNowAuthenticationTypePtrInput {
	return (*serviceNowAuthenticationTypePtr)(&v)
}

func (*serviceNowAuthenticationTypePtr) ElementType() reflect.Type {
	return serviceNowAuthenticationTypePtrType
}

func (in *serviceNowAuthenticationTypePtr) ToServiceNowAuthenticationTypePtrOutput() ServiceNowAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(ServiceNowAuthenticationTypePtrOutput)
}

func (in *serviceNowAuthenticationTypePtr) ToServiceNowAuthenticationTypePtrOutputWithContext(ctx context.Context) ServiceNowAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceNowAuthenticationTypePtrOutput)
}

type SftpAuthenticationType string

const (
	SftpAuthenticationTypeBasic        = SftpAuthenticationType("Basic")
	SftpAuthenticationTypeSshPublicKey = SftpAuthenticationType("SshPublicKey")
	SftpAuthenticationTypeMultiFactor  = SftpAuthenticationType("MultiFactor")
)

func (SftpAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*SftpAuthenticationType)(nil)).Elem()
}

func (e SftpAuthenticationType) ToSftpAuthenticationTypeOutput() SftpAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(SftpAuthenticationTypeOutput)
}

func (e SftpAuthenticationType) ToSftpAuthenticationTypeOutputWithContext(ctx context.Context) SftpAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SftpAuthenticationTypeOutput)
}

func (e SftpAuthenticationType) ToSftpAuthenticationTypePtrOutput() SftpAuthenticationTypePtrOutput {
	return e.ToSftpAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e SftpAuthenticationType) ToSftpAuthenticationTypePtrOutputWithContext(ctx context.Context) SftpAuthenticationTypePtrOutput {
	return SftpAuthenticationType(e).ToSftpAuthenticationTypeOutputWithContext(ctx).ToSftpAuthenticationTypePtrOutputWithContext(ctx)
}

func (e SftpAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SftpAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SftpAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SftpAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SftpAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (SftpAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SftpAuthenticationType)(nil)).Elem()
}

func (o SftpAuthenticationTypeOutput) ToSftpAuthenticationTypeOutput() SftpAuthenticationTypeOutput {
	return o
}

func (o SftpAuthenticationTypeOutput) ToSftpAuthenticationTypeOutputWithContext(ctx context.Context) SftpAuthenticationTypeOutput {
	return o
}

func (o SftpAuthenticationTypeOutput) ToSftpAuthenticationTypePtrOutput() SftpAuthenticationTypePtrOutput {
	return o.ToSftpAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o SftpAuthenticationTypeOutput) ToSftpAuthenticationTypePtrOutputWithContext(ctx context.Context) SftpAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SftpAuthenticationType) *SftpAuthenticationType {
		return &v
	}).(SftpAuthenticationTypePtrOutput)
}

func (o SftpAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SftpAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SftpAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SftpAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SftpAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SftpAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SftpAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (SftpAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SftpAuthenticationType)(nil)).Elem()
}

func (o SftpAuthenticationTypePtrOutput) ToSftpAuthenticationTypePtrOutput() SftpAuthenticationTypePtrOutput {
	return o
}

func (o SftpAuthenticationTypePtrOutput) ToSftpAuthenticationTypePtrOutputWithContext(ctx context.Context) SftpAuthenticationTypePtrOutput {
	return o
}

func (o SftpAuthenticationTypePtrOutput) Elem() SftpAuthenticationTypeOutput {
	return o.ApplyT(func(v *SftpAuthenticationType) SftpAuthenticationType {
		if v != nil {
			return *v
		}
		var ret SftpAuthenticationType
		return ret
	}).(SftpAuthenticationTypeOutput)
}

func (o SftpAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SftpAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SftpAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SftpAuthenticationTypeInput interface {
	pulumi.Input

	ToSftpAuthenticationTypeOutput() SftpAuthenticationTypeOutput
	ToSftpAuthenticationTypeOutputWithContext(context.Context) SftpAuthenticationTypeOutput
}

var sftpAuthenticationTypePtrType = reflect.TypeOf((**SftpAuthenticationType)(nil)).Elem()

type SftpAuthenticationTypePtrInput interface {
	pulumi.Input

	ToSftpAuthenticationTypePtrOutput() SftpAuthenticationTypePtrOutput
	ToSftpAuthenticationTypePtrOutputWithContext(context.Context) SftpAuthenticationTypePtrOutput
}

type sftpAuthenticationTypePtr string

func SftpAuthenticationTypePtr(v string) SftpAuthenticationTypePtrInput {
	return (*sftpAuthenticationTypePtr)(&v)
}

func (*sftpAuthenticationTypePtr) ElementType() reflect.Type {
	return sftpAuthenticationTypePtrType
}

func (in *sftpAuthenticationTypePtr) ToSftpAuthenticationTypePtrOutput() SftpAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(SftpAuthenticationTypePtrOutput)
}

func (in *sftpAuthenticationTypePtr) ToSftpAuthenticationTypePtrOutputWithContext(ctx context.Context) SftpAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SftpAuthenticationTypePtrOutput)
}

type SparkAuthenticationType string

const (
	SparkAuthenticationTypeAnonymous                    = SparkAuthenticationType("Anonymous")
	SparkAuthenticationTypeUsername                     = SparkAuthenticationType("Username")
	SparkAuthenticationTypeUsernameAndPassword          = SparkAuthenticationType("UsernameAndPassword")
	SparkAuthenticationTypeWindowsAzureHDInsightService = SparkAuthenticationType("WindowsAzureHDInsightService")
)

func (SparkAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkAuthenticationType)(nil)).Elem()
}

func (e SparkAuthenticationType) ToSparkAuthenticationTypeOutput() SparkAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(SparkAuthenticationTypeOutput)
}

func (e SparkAuthenticationType) ToSparkAuthenticationTypeOutputWithContext(ctx context.Context) SparkAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SparkAuthenticationTypeOutput)
}

func (e SparkAuthenticationType) ToSparkAuthenticationTypePtrOutput() SparkAuthenticationTypePtrOutput {
	return e.ToSparkAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e SparkAuthenticationType) ToSparkAuthenticationTypePtrOutputWithContext(ctx context.Context) SparkAuthenticationTypePtrOutput {
	return SparkAuthenticationType(e).ToSparkAuthenticationTypeOutputWithContext(ctx).ToSparkAuthenticationTypePtrOutputWithContext(ctx)
}

func (e SparkAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SparkAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SparkAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SparkAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SparkAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (SparkAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkAuthenticationType)(nil)).Elem()
}

func (o SparkAuthenticationTypeOutput) ToSparkAuthenticationTypeOutput() SparkAuthenticationTypeOutput {
	return o
}

func (o SparkAuthenticationTypeOutput) ToSparkAuthenticationTypeOutputWithContext(ctx context.Context) SparkAuthenticationTypeOutput {
	return o
}

func (o SparkAuthenticationTypeOutput) ToSparkAuthenticationTypePtrOutput() SparkAuthenticationTypePtrOutput {
	return o.ToSparkAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o SparkAuthenticationTypeOutput) ToSparkAuthenticationTypePtrOutputWithContext(ctx context.Context) SparkAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SparkAuthenticationType) *SparkAuthenticationType {
		return &v
	}).(SparkAuthenticationTypePtrOutput)
}

func (o SparkAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SparkAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SparkAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SparkAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SparkAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SparkAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SparkAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (SparkAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkAuthenticationType)(nil)).Elem()
}

func (o SparkAuthenticationTypePtrOutput) ToSparkAuthenticationTypePtrOutput() SparkAuthenticationTypePtrOutput {
	return o
}

func (o SparkAuthenticationTypePtrOutput) ToSparkAuthenticationTypePtrOutputWithContext(ctx context.Context) SparkAuthenticationTypePtrOutput {
	return o
}

func (o SparkAuthenticationTypePtrOutput) Elem() SparkAuthenticationTypeOutput {
	return o.ApplyT(func(v *SparkAuthenticationType) SparkAuthenticationType {
		if v != nil {
			return *v
		}
		var ret SparkAuthenticationType
		return ret
	}).(SparkAuthenticationTypeOutput)
}

func (o SparkAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SparkAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SparkAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SparkAuthenticationTypeInput interface {
	pulumi.Input

	ToSparkAuthenticationTypeOutput() SparkAuthenticationTypeOutput
	ToSparkAuthenticationTypeOutputWithContext(context.Context) SparkAuthenticationTypeOutput
}

var sparkAuthenticationTypePtrType = reflect.TypeOf((**SparkAuthenticationType)(nil)).Elem()

type SparkAuthenticationTypePtrInput interface {
	pulumi.Input

	ToSparkAuthenticationTypePtrOutput() SparkAuthenticationTypePtrOutput
	ToSparkAuthenticationTypePtrOutputWithContext(context.Context) SparkAuthenticationTypePtrOutput
}

type sparkAuthenticationTypePtr string

func SparkAuthenticationTypePtr(v string) SparkAuthenticationTypePtrInput {
	return (*sparkAuthenticationTypePtr)(&v)
}

func (*sparkAuthenticationTypePtr) ElementType() reflect.Type {
	return sparkAuthenticationTypePtrType
}

func (in *sparkAuthenticationTypePtr) ToSparkAuthenticationTypePtrOutput() SparkAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(SparkAuthenticationTypePtrOutput)
}

func (in *sparkAuthenticationTypePtr) ToSparkAuthenticationTypePtrOutputWithContext(ctx context.Context) SparkAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SparkAuthenticationTypePtrOutput)
}

type SparkServerType string

const (
	SparkServerTypeSharkServer       = SparkServerType("SharkServer")
	SparkServerTypeSharkServer2      = SparkServerType("SharkServer2")
	SparkServerTypeSparkThriftServer = SparkServerType("SparkThriftServer")
)

func (SparkServerType) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkServerType)(nil)).Elem()
}

func (e SparkServerType) ToSparkServerTypeOutput() SparkServerTypeOutput {
	return pulumi.ToOutput(e).(SparkServerTypeOutput)
}

func (e SparkServerType) ToSparkServerTypeOutputWithContext(ctx context.Context) SparkServerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SparkServerTypeOutput)
}

func (e SparkServerType) ToSparkServerTypePtrOutput() SparkServerTypePtrOutput {
	return e.ToSparkServerTypePtrOutputWithContext(context.Background())
}

func (e SparkServerType) ToSparkServerTypePtrOutputWithContext(ctx context.Context) SparkServerTypePtrOutput {
	return SparkServerType(e).ToSparkServerTypeOutputWithContext(ctx).ToSparkServerTypePtrOutputWithContext(ctx)
}

func (e SparkServerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SparkServerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SparkServerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SparkServerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SparkServerTypeOutput struct{ *pulumi.OutputState }

func (SparkServerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkServerType)(nil)).Elem()
}

func (o SparkServerTypeOutput) ToSparkServerTypeOutput() SparkServerTypeOutput {
	return o
}

func (o SparkServerTypeOutput) ToSparkServerTypeOutputWithContext(ctx context.Context) SparkServerTypeOutput {
	return o
}

func (o SparkServerTypeOutput) ToSparkServerTypePtrOutput() SparkServerTypePtrOutput {
	return o.ToSparkServerTypePtrOutputWithContext(context.Background())
}

func (o SparkServerTypeOutput) ToSparkServerTypePtrOutputWithContext(ctx context.Context) SparkServerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SparkServerType) *SparkServerType {
		return &v
	}).(SparkServerTypePtrOutput)
}

func (o SparkServerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SparkServerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SparkServerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SparkServerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SparkServerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SparkServerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SparkServerTypePtrOutput struct{ *pulumi.OutputState }

func (SparkServerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkServerType)(nil)).Elem()
}

func (o SparkServerTypePtrOutput) ToSparkServerTypePtrOutput() SparkServerTypePtrOutput {
	return o
}

func (o SparkServerTypePtrOutput) ToSparkServerTypePtrOutputWithContext(ctx context.Context) SparkServerTypePtrOutput {
	return o
}

func (o SparkServerTypePtrOutput) Elem() SparkServerTypeOutput {
	return o.ApplyT(func(v *SparkServerType) SparkServerType {
		if v != nil {
			return *v
		}
		var ret SparkServerType
		return ret
	}).(SparkServerTypeOutput)
}

func (o SparkServerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SparkServerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SparkServerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SparkServerTypeInput interface {
	pulumi.Input

	ToSparkServerTypeOutput() SparkServerTypeOutput
	ToSparkServerTypeOutputWithContext(context.Context) SparkServerTypeOutput
}

var sparkServerTypePtrType = reflect.TypeOf((**SparkServerType)(nil)).Elem()

type SparkServerTypePtrInput interface {
	pulumi.Input

	ToSparkServerTypePtrOutput() SparkServerTypePtrOutput
	ToSparkServerTypePtrOutputWithContext(context.Context) SparkServerTypePtrOutput
}

type sparkServerTypePtr string

func SparkServerTypePtr(v string) SparkServerTypePtrInput {
	return (*sparkServerTypePtr)(&v)
}

func (*sparkServerTypePtr) ElementType() reflect.Type {
	return sparkServerTypePtrType
}

func (in *sparkServerTypePtr) ToSparkServerTypePtrOutput() SparkServerTypePtrOutput {
	return pulumi.ToOutput(in).(SparkServerTypePtrOutput)
}

func (in *sparkServerTypePtr) ToSparkServerTypePtrOutputWithContext(ctx context.Context) SparkServerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SparkServerTypePtrOutput)
}

type SparkThriftTransportProtocol string

const (
	SparkThriftTransportProtocolBinary = SparkThriftTransportProtocol("Binary")
	SparkThriftTransportProtocolSASL   = SparkThriftTransportProtocol("SASL")
	SparkThriftTransportProtocol_HTTP_ = SparkThriftTransportProtocol("HTTP ")
)

func (SparkThriftTransportProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkThriftTransportProtocol)(nil)).Elem()
}

func (e SparkThriftTransportProtocol) ToSparkThriftTransportProtocolOutput() SparkThriftTransportProtocolOutput {
	return pulumi.ToOutput(e).(SparkThriftTransportProtocolOutput)
}

func (e SparkThriftTransportProtocol) ToSparkThriftTransportProtocolOutputWithContext(ctx context.Context) SparkThriftTransportProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SparkThriftTransportProtocolOutput)
}

func (e SparkThriftTransportProtocol) ToSparkThriftTransportProtocolPtrOutput() SparkThriftTransportProtocolPtrOutput {
	return e.ToSparkThriftTransportProtocolPtrOutputWithContext(context.Background())
}

func (e SparkThriftTransportProtocol) ToSparkThriftTransportProtocolPtrOutputWithContext(ctx context.Context) SparkThriftTransportProtocolPtrOutput {
	return SparkThriftTransportProtocol(e).ToSparkThriftTransportProtocolOutputWithContext(ctx).ToSparkThriftTransportProtocolPtrOutputWithContext(ctx)
}

func (e SparkThriftTransportProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SparkThriftTransportProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SparkThriftTransportProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SparkThriftTransportProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SparkThriftTransportProtocolOutput struct{ *pulumi.OutputState }

func (SparkThriftTransportProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SparkThriftTransportProtocol)(nil)).Elem()
}

func (o SparkThriftTransportProtocolOutput) ToSparkThriftTransportProtocolOutput() SparkThriftTransportProtocolOutput {
	return o
}

func (o SparkThriftTransportProtocolOutput) ToSparkThriftTransportProtocolOutputWithContext(ctx context.Context) SparkThriftTransportProtocolOutput {
	return o
}

func (o SparkThriftTransportProtocolOutput) ToSparkThriftTransportProtocolPtrOutput() SparkThriftTransportProtocolPtrOutput {
	return o.ToSparkThriftTransportProtocolPtrOutputWithContext(context.Background())
}

func (o SparkThriftTransportProtocolOutput) ToSparkThriftTransportProtocolPtrOutputWithContext(ctx context.Context) SparkThriftTransportProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SparkThriftTransportProtocol) *SparkThriftTransportProtocol {
		return &v
	}).(SparkThriftTransportProtocolPtrOutput)
}

func (o SparkThriftTransportProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SparkThriftTransportProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SparkThriftTransportProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SparkThriftTransportProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SparkThriftTransportProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SparkThriftTransportProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SparkThriftTransportProtocolPtrOutput struct{ *pulumi.OutputState }

func (SparkThriftTransportProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SparkThriftTransportProtocol)(nil)).Elem()
}

func (o SparkThriftTransportProtocolPtrOutput) ToSparkThriftTransportProtocolPtrOutput() SparkThriftTransportProtocolPtrOutput {
	return o
}

func (o SparkThriftTransportProtocolPtrOutput) ToSparkThriftTransportProtocolPtrOutputWithContext(ctx context.Context) SparkThriftTransportProtocolPtrOutput {
	return o
}

func (o SparkThriftTransportProtocolPtrOutput) Elem() SparkThriftTransportProtocolOutput {
	return o.ApplyT(func(v *SparkThriftTransportProtocol) SparkThriftTransportProtocol {
		if v != nil {
			return *v
		}
		var ret SparkThriftTransportProtocol
		return ret
	}).(SparkThriftTransportProtocolOutput)
}

func (o SparkThriftTransportProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SparkThriftTransportProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SparkThriftTransportProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SparkThriftTransportProtocolInput interface {
	pulumi.Input

	ToSparkThriftTransportProtocolOutput() SparkThriftTransportProtocolOutput
	ToSparkThriftTransportProtocolOutputWithContext(context.Context) SparkThriftTransportProtocolOutput
}

var sparkThriftTransportProtocolPtrType = reflect.TypeOf((**SparkThriftTransportProtocol)(nil)).Elem()

type SparkThriftTransportProtocolPtrInput interface {
	pulumi.Input

	ToSparkThriftTransportProtocolPtrOutput() SparkThriftTransportProtocolPtrOutput
	ToSparkThriftTransportProtocolPtrOutputWithContext(context.Context) SparkThriftTransportProtocolPtrOutput
}

type sparkThriftTransportProtocolPtr string

func SparkThriftTransportProtocolPtr(v string) SparkThriftTransportProtocolPtrInput {
	return (*sparkThriftTransportProtocolPtr)(&v)
}

func (*sparkThriftTransportProtocolPtr) ElementType() reflect.Type {
	return sparkThriftTransportProtocolPtrType
}

func (in *sparkThriftTransportProtocolPtr) ToSparkThriftTransportProtocolPtrOutput() SparkThriftTransportProtocolPtrOutput {
	return pulumi.ToOutput(in).(SparkThriftTransportProtocolPtrOutput)
}

func (in *sparkThriftTransportProtocolPtr) ToSparkThriftTransportProtocolPtrOutputWithContext(ctx context.Context) SparkThriftTransportProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SparkThriftTransportProtocolPtrOutput)
}

type SqlAlwaysEncryptedAkvAuthType string

const (
	SqlAlwaysEncryptedAkvAuthTypeServicePrincipal            = SqlAlwaysEncryptedAkvAuthType("ServicePrincipal")
	SqlAlwaysEncryptedAkvAuthTypeManagedIdentity             = SqlAlwaysEncryptedAkvAuthType("ManagedIdentity")
	SqlAlwaysEncryptedAkvAuthTypeUserAssignedManagedIdentity = SqlAlwaysEncryptedAkvAuthType("UserAssignedManagedIdentity")
)

func (SqlAlwaysEncryptedAkvAuthType) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlAlwaysEncryptedAkvAuthType)(nil)).Elem()
}

func (e SqlAlwaysEncryptedAkvAuthType) ToSqlAlwaysEncryptedAkvAuthTypeOutput() SqlAlwaysEncryptedAkvAuthTypeOutput {
	return pulumi.ToOutput(e).(SqlAlwaysEncryptedAkvAuthTypeOutput)
}

func (e SqlAlwaysEncryptedAkvAuthType) ToSqlAlwaysEncryptedAkvAuthTypeOutputWithContext(ctx context.Context) SqlAlwaysEncryptedAkvAuthTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SqlAlwaysEncryptedAkvAuthTypeOutput)
}

func (e SqlAlwaysEncryptedAkvAuthType) ToSqlAlwaysEncryptedAkvAuthTypePtrOutput() SqlAlwaysEncryptedAkvAuthTypePtrOutput {
	return e.ToSqlAlwaysEncryptedAkvAuthTypePtrOutputWithContext(context.Background())
}

func (e SqlAlwaysEncryptedAkvAuthType) ToSqlAlwaysEncryptedAkvAuthTypePtrOutputWithContext(ctx context.Context) SqlAlwaysEncryptedAkvAuthTypePtrOutput {
	return SqlAlwaysEncryptedAkvAuthType(e).ToSqlAlwaysEncryptedAkvAuthTypeOutputWithContext(ctx).ToSqlAlwaysEncryptedAkvAuthTypePtrOutputWithContext(ctx)
}

func (e SqlAlwaysEncryptedAkvAuthType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlAlwaysEncryptedAkvAuthType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlAlwaysEncryptedAkvAuthType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SqlAlwaysEncryptedAkvAuthType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SqlAlwaysEncryptedAkvAuthTypeOutput struct{ *pulumi.OutputState }

func (SqlAlwaysEncryptedAkvAuthTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlAlwaysEncryptedAkvAuthType)(nil)).Elem()
}

func (o SqlAlwaysEncryptedAkvAuthTypeOutput) ToSqlAlwaysEncryptedAkvAuthTypeOutput() SqlAlwaysEncryptedAkvAuthTypeOutput {
	return o
}

func (o SqlAlwaysEncryptedAkvAuthTypeOutput) ToSqlAlwaysEncryptedAkvAuthTypeOutputWithContext(ctx context.Context) SqlAlwaysEncryptedAkvAuthTypeOutput {
	return o
}

func (o SqlAlwaysEncryptedAkvAuthTypeOutput) ToSqlAlwaysEncryptedAkvAuthTypePtrOutput() SqlAlwaysEncryptedAkvAuthTypePtrOutput {
	return o.ToSqlAlwaysEncryptedAkvAuthTypePtrOutputWithContext(context.Background())
}

func (o SqlAlwaysEncryptedAkvAuthTypeOutput) ToSqlAlwaysEncryptedAkvAuthTypePtrOutputWithContext(ctx context.Context) SqlAlwaysEncryptedAkvAuthTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlAlwaysEncryptedAkvAuthType) *SqlAlwaysEncryptedAkvAuthType {
		return &v
	}).(SqlAlwaysEncryptedAkvAuthTypePtrOutput)
}

func (o SqlAlwaysEncryptedAkvAuthTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SqlAlwaysEncryptedAkvAuthTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlAlwaysEncryptedAkvAuthType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SqlAlwaysEncryptedAkvAuthTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlAlwaysEncryptedAkvAuthTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlAlwaysEncryptedAkvAuthType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SqlAlwaysEncryptedAkvAuthTypePtrOutput struct{ *pulumi.OutputState }

func (SqlAlwaysEncryptedAkvAuthTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlAlwaysEncryptedAkvAuthType)(nil)).Elem()
}

func (o SqlAlwaysEncryptedAkvAuthTypePtrOutput) ToSqlAlwaysEncryptedAkvAuthTypePtrOutput() SqlAlwaysEncryptedAkvAuthTypePtrOutput {
	return o
}

func (o SqlAlwaysEncryptedAkvAuthTypePtrOutput) ToSqlAlwaysEncryptedAkvAuthTypePtrOutputWithContext(ctx context.Context) SqlAlwaysEncryptedAkvAuthTypePtrOutput {
	return o
}

func (o SqlAlwaysEncryptedAkvAuthTypePtrOutput) Elem() SqlAlwaysEncryptedAkvAuthTypeOutput {
	return o.ApplyT(func(v *SqlAlwaysEncryptedAkvAuthType) SqlAlwaysEncryptedAkvAuthType {
		if v != nil {
			return *v
		}
		var ret SqlAlwaysEncryptedAkvAuthType
		return ret
	}).(SqlAlwaysEncryptedAkvAuthTypeOutput)
}

func (o SqlAlwaysEncryptedAkvAuthTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlAlwaysEncryptedAkvAuthTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SqlAlwaysEncryptedAkvAuthType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SqlAlwaysEncryptedAkvAuthTypeInput interface {
	pulumi.Input

	ToSqlAlwaysEncryptedAkvAuthTypeOutput() SqlAlwaysEncryptedAkvAuthTypeOutput
	ToSqlAlwaysEncryptedAkvAuthTypeOutputWithContext(context.Context) SqlAlwaysEncryptedAkvAuthTypeOutput
}

var sqlAlwaysEncryptedAkvAuthTypePtrType = reflect.TypeOf((**SqlAlwaysEncryptedAkvAuthType)(nil)).Elem()

type SqlAlwaysEncryptedAkvAuthTypePtrInput interface {
	pulumi.Input

	ToSqlAlwaysEncryptedAkvAuthTypePtrOutput() SqlAlwaysEncryptedAkvAuthTypePtrOutput
	ToSqlAlwaysEncryptedAkvAuthTypePtrOutputWithContext(context.Context) SqlAlwaysEncryptedAkvAuthTypePtrOutput
}

type sqlAlwaysEncryptedAkvAuthTypePtr string

func SqlAlwaysEncryptedAkvAuthTypePtr(v string) SqlAlwaysEncryptedAkvAuthTypePtrInput {
	return (*sqlAlwaysEncryptedAkvAuthTypePtr)(&v)
}

func (*sqlAlwaysEncryptedAkvAuthTypePtr) ElementType() reflect.Type {
	return sqlAlwaysEncryptedAkvAuthTypePtrType
}

func (in *sqlAlwaysEncryptedAkvAuthTypePtr) ToSqlAlwaysEncryptedAkvAuthTypePtrOutput() SqlAlwaysEncryptedAkvAuthTypePtrOutput {
	return pulumi.ToOutput(in).(SqlAlwaysEncryptedAkvAuthTypePtrOutput)
}

func (in *sqlAlwaysEncryptedAkvAuthTypePtr) ToSqlAlwaysEncryptedAkvAuthTypePtrOutputWithContext(ctx context.Context) SqlAlwaysEncryptedAkvAuthTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SqlAlwaysEncryptedAkvAuthTypePtrOutput)
}

type SsisLogLocationType string

const (
	SsisLogLocationTypeFile = SsisLogLocationType("File")
)

func (SsisLogLocationType) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisLogLocationType)(nil)).Elem()
}

func (e SsisLogLocationType) ToSsisLogLocationTypeOutput() SsisLogLocationTypeOutput {
	return pulumi.ToOutput(e).(SsisLogLocationTypeOutput)
}

func (e SsisLogLocationType) ToSsisLogLocationTypeOutputWithContext(ctx context.Context) SsisLogLocationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SsisLogLocationTypeOutput)
}

func (e SsisLogLocationType) ToSsisLogLocationTypePtrOutput() SsisLogLocationTypePtrOutput {
	return e.ToSsisLogLocationTypePtrOutputWithContext(context.Background())
}

func (e SsisLogLocationType) ToSsisLogLocationTypePtrOutputWithContext(ctx context.Context) SsisLogLocationTypePtrOutput {
	return SsisLogLocationType(e).ToSsisLogLocationTypeOutputWithContext(ctx).ToSsisLogLocationTypePtrOutputWithContext(ctx)
}

func (e SsisLogLocationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SsisLogLocationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SsisLogLocationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SsisLogLocationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SsisLogLocationTypeOutput struct{ *pulumi.OutputState }

func (SsisLogLocationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisLogLocationType)(nil)).Elem()
}

func (o SsisLogLocationTypeOutput) ToSsisLogLocationTypeOutput() SsisLogLocationTypeOutput {
	return o
}

func (o SsisLogLocationTypeOutput) ToSsisLogLocationTypeOutputWithContext(ctx context.Context) SsisLogLocationTypeOutput {
	return o
}

func (o SsisLogLocationTypeOutput) ToSsisLogLocationTypePtrOutput() SsisLogLocationTypePtrOutput {
	return o.ToSsisLogLocationTypePtrOutputWithContext(context.Background())
}

func (o SsisLogLocationTypeOutput) ToSsisLogLocationTypePtrOutputWithContext(ctx context.Context) SsisLogLocationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SsisLogLocationType) *SsisLogLocationType {
		return &v
	}).(SsisLogLocationTypePtrOutput)
}

func (o SsisLogLocationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SsisLogLocationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SsisLogLocationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SsisLogLocationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SsisLogLocationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SsisLogLocationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SsisLogLocationTypePtrOutput struct{ *pulumi.OutputState }

func (SsisLogLocationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SsisLogLocationType)(nil)).Elem()
}

func (o SsisLogLocationTypePtrOutput) ToSsisLogLocationTypePtrOutput() SsisLogLocationTypePtrOutput {
	return o
}

func (o SsisLogLocationTypePtrOutput) ToSsisLogLocationTypePtrOutputWithContext(ctx context.Context) SsisLogLocationTypePtrOutput {
	return o
}

func (o SsisLogLocationTypePtrOutput) Elem() SsisLogLocationTypeOutput {
	return o.ApplyT(func(v *SsisLogLocationType) SsisLogLocationType {
		if v != nil {
			return *v
		}
		var ret SsisLogLocationType
		return ret
	}).(SsisLogLocationTypeOutput)
}

func (o SsisLogLocationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SsisLogLocationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SsisLogLocationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SsisLogLocationTypeInput interface {
	pulumi.Input

	ToSsisLogLocationTypeOutput() SsisLogLocationTypeOutput
	ToSsisLogLocationTypeOutputWithContext(context.Context) SsisLogLocationTypeOutput
}

var ssisLogLocationTypePtrType = reflect.TypeOf((**SsisLogLocationType)(nil)).Elem()

type SsisLogLocationTypePtrInput interface {
	pulumi.Input

	ToSsisLogLocationTypePtrOutput() SsisLogLocationTypePtrOutput
	ToSsisLogLocationTypePtrOutputWithContext(context.Context) SsisLogLocationTypePtrOutput
}

type ssisLogLocationTypePtr string

func SsisLogLocationTypePtr(v string) SsisLogLocationTypePtrInput {
	return (*ssisLogLocationTypePtr)(&v)
}

func (*ssisLogLocationTypePtr) ElementType() reflect.Type {
	return ssisLogLocationTypePtrType
}

func (in *ssisLogLocationTypePtr) ToSsisLogLocationTypePtrOutput() SsisLogLocationTypePtrOutput {
	return pulumi.ToOutput(in).(SsisLogLocationTypePtrOutput)
}

func (in *ssisLogLocationTypePtr) ToSsisLogLocationTypePtrOutputWithContext(ctx context.Context) SsisLogLocationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SsisLogLocationTypePtrOutput)
}

type SsisPackageLocationType string

const (
	SsisPackageLocationTypeSSISDB        = SsisPackageLocationType("SSISDB")
	SsisPackageLocationTypeFile          = SsisPackageLocationType("File")
	SsisPackageLocationTypeInlinePackage = SsisPackageLocationType("InlinePackage")
	SsisPackageLocationTypePackageStore  = SsisPackageLocationType("PackageStore")
)

func (SsisPackageLocationType) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisPackageLocationType)(nil)).Elem()
}

func (e SsisPackageLocationType) ToSsisPackageLocationTypeOutput() SsisPackageLocationTypeOutput {
	return pulumi.ToOutput(e).(SsisPackageLocationTypeOutput)
}

func (e SsisPackageLocationType) ToSsisPackageLocationTypeOutputWithContext(ctx context.Context) SsisPackageLocationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SsisPackageLocationTypeOutput)
}

func (e SsisPackageLocationType) ToSsisPackageLocationTypePtrOutput() SsisPackageLocationTypePtrOutput {
	return e.ToSsisPackageLocationTypePtrOutputWithContext(context.Background())
}

func (e SsisPackageLocationType) ToSsisPackageLocationTypePtrOutputWithContext(ctx context.Context) SsisPackageLocationTypePtrOutput {
	return SsisPackageLocationType(e).ToSsisPackageLocationTypeOutputWithContext(ctx).ToSsisPackageLocationTypePtrOutputWithContext(ctx)
}

func (e SsisPackageLocationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SsisPackageLocationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SsisPackageLocationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SsisPackageLocationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SsisPackageLocationTypeOutput struct{ *pulumi.OutputState }

func (SsisPackageLocationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SsisPackageLocationType)(nil)).Elem()
}

func (o SsisPackageLocationTypeOutput) ToSsisPackageLocationTypeOutput() SsisPackageLocationTypeOutput {
	return o
}

func (o SsisPackageLocationTypeOutput) ToSsisPackageLocationTypeOutputWithContext(ctx context.Context) SsisPackageLocationTypeOutput {
	return o
}

func (o SsisPackageLocationTypeOutput) ToSsisPackageLocationTypePtrOutput() SsisPackageLocationTypePtrOutput {
	return o.ToSsisPackageLocationTypePtrOutputWithContext(context.Background())
}

func (o SsisPackageLocationTypeOutput) ToSsisPackageLocationTypePtrOutputWithContext(ctx context.Context) SsisPackageLocationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SsisPackageLocationType) *SsisPackageLocationType {
		return &v
	}).(SsisPackageLocationTypePtrOutput)
}

func (o SsisPackageLocationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SsisPackageLocationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SsisPackageLocationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SsisPackageLocationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SsisPackageLocationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SsisPackageLocationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SsisPackageLocationTypePtrOutput struct{ *pulumi.OutputState }

func (SsisPackageLocationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SsisPackageLocationType)(nil)).Elem()
}

func (o SsisPackageLocationTypePtrOutput) ToSsisPackageLocationTypePtrOutput() SsisPackageLocationTypePtrOutput {
	return o
}

func (o SsisPackageLocationTypePtrOutput) ToSsisPackageLocationTypePtrOutputWithContext(ctx context.Context) SsisPackageLocationTypePtrOutput {
	return o
}

func (o SsisPackageLocationTypePtrOutput) Elem() SsisPackageLocationTypeOutput {
	return o.ApplyT(func(v *SsisPackageLocationType) SsisPackageLocationType {
		if v != nil {
			return *v
		}
		var ret SsisPackageLocationType
		return ret
	}).(SsisPackageLocationTypeOutput)
}

func (o SsisPackageLocationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SsisPackageLocationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SsisPackageLocationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SsisPackageLocationTypeInput interface {
	pulumi.Input

	ToSsisPackageLocationTypeOutput() SsisPackageLocationTypeOutput
	ToSsisPackageLocationTypeOutputWithContext(context.Context) SsisPackageLocationTypeOutput
}

var ssisPackageLocationTypePtrType = reflect.TypeOf((**SsisPackageLocationType)(nil)).Elem()

type SsisPackageLocationTypePtrInput interface {
	pulumi.Input

	ToSsisPackageLocationTypePtrOutput() SsisPackageLocationTypePtrOutput
	ToSsisPackageLocationTypePtrOutputWithContext(context.Context) SsisPackageLocationTypePtrOutput
}

type ssisPackageLocationTypePtr string

func SsisPackageLocationTypePtr(v string) SsisPackageLocationTypePtrInput {
	return (*ssisPackageLocationTypePtr)(&v)
}

func (*ssisPackageLocationTypePtr) ElementType() reflect.Type {
	return ssisPackageLocationTypePtrType
}

func (in *ssisPackageLocationTypePtr) ToSsisPackageLocationTypePtrOutput() SsisPackageLocationTypePtrOutput {
	return pulumi.ToOutput(in).(SsisPackageLocationTypePtrOutput)
}

func (in *ssisPackageLocationTypePtr) ToSsisPackageLocationTypePtrOutputWithContext(ctx context.Context) SsisPackageLocationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SsisPackageLocationTypePtrOutput)
}

type StoredProcedureParameterType string

const (
	StoredProcedureParameterTypeString  = StoredProcedureParameterType("String")
	StoredProcedureParameterTypeInt     = StoredProcedureParameterType("Int")
	StoredProcedureParameterTypeInt64   = StoredProcedureParameterType("Int64")
	StoredProcedureParameterTypeDecimal = StoredProcedureParameterType("Decimal")
	StoredProcedureParameterTypeGuid    = StoredProcedureParameterType("Guid")
	StoredProcedureParameterTypeBoolean = StoredProcedureParameterType("Boolean")
	StoredProcedureParameterTypeDate    = StoredProcedureParameterType("Date")
)

func (StoredProcedureParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*StoredProcedureParameterType)(nil)).Elem()
}

func (e StoredProcedureParameterType) ToStoredProcedureParameterTypeOutput() StoredProcedureParameterTypeOutput {
	return pulumi.ToOutput(e).(StoredProcedureParameterTypeOutput)
}

func (e StoredProcedureParameterType) ToStoredProcedureParameterTypeOutputWithContext(ctx context.Context) StoredProcedureParameterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StoredProcedureParameterTypeOutput)
}

func (e StoredProcedureParameterType) ToStoredProcedureParameterTypePtrOutput() StoredProcedureParameterTypePtrOutput {
	return e.ToStoredProcedureParameterTypePtrOutputWithContext(context.Background())
}

func (e StoredProcedureParameterType) ToStoredProcedureParameterTypePtrOutputWithContext(ctx context.Context) StoredProcedureParameterTypePtrOutput {
	return StoredProcedureParameterType(e).ToStoredProcedureParameterTypeOutputWithContext(ctx).ToStoredProcedureParameterTypePtrOutputWithContext(ctx)
}

func (e StoredProcedureParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StoredProcedureParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StoredProcedureParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StoredProcedureParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StoredProcedureParameterTypeOutput struct{ *pulumi.OutputState }

func (StoredProcedureParameterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoredProcedureParameterType)(nil)).Elem()
}

func (o StoredProcedureParameterTypeOutput) ToStoredProcedureParameterTypeOutput() StoredProcedureParameterTypeOutput {
	return o
}

func (o StoredProcedureParameterTypeOutput) ToStoredProcedureParameterTypeOutputWithContext(ctx context.Context) StoredProcedureParameterTypeOutput {
	return o
}

func (o StoredProcedureParameterTypeOutput) ToStoredProcedureParameterTypePtrOutput() StoredProcedureParameterTypePtrOutput {
	return o.ToStoredProcedureParameterTypePtrOutputWithContext(context.Background())
}

func (o StoredProcedureParameterTypeOutput) ToStoredProcedureParameterTypePtrOutputWithContext(ctx context.Context) StoredProcedureParameterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StoredProcedureParameterType) *StoredProcedureParameterType {
		return &v
	}).(StoredProcedureParameterTypePtrOutput)
}

func (o StoredProcedureParameterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StoredProcedureParameterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StoredProcedureParameterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StoredProcedureParameterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StoredProcedureParameterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StoredProcedureParameterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StoredProcedureParameterTypePtrOutput struct{ *pulumi.OutputState }

func (StoredProcedureParameterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoredProcedureParameterType)(nil)).Elem()
}

func (o StoredProcedureParameterTypePtrOutput) ToStoredProcedureParameterTypePtrOutput() StoredProcedureParameterTypePtrOutput {
	return o
}

func (o StoredProcedureParameterTypePtrOutput) ToStoredProcedureParameterTypePtrOutputWithContext(ctx context.Context) StoredProcedureParameterTypePtrOutput {
	return o
}

func (o StoredProcedureParameterTypePtrOutput) Elem() StoredProcedureParameterTypeOutput {
	return o.ApplyT(func(v *StoredProcedureParameterType) StoredProcedureParameterType {
		if v != nil {
			return *v
		}
		var ret StoredProcedureParameterType
		return ret
	}).(StoredProcedureParameterTypeOutput)
}

func (o StoredProcedureParameterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StoredProcedureParameterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StoredProcedureParameterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StoredProcedureParameterTypeInput interface {
	pulumi.Input

	ToStoredProcedureParameterTypeOutput() StoredProcedureParameterTypeOutput
	ToStoredProcedureParameterTypeOutputWithContext(context.Context) StoredProcedureParameterTypeOutput
}

var storedProcedureParameterTypePtrType = reflect.TypeOf((**StoredProcedureParameterType)(nil)).Elem()

type StoredProcedureParameterTypePtrInput interface {
	pulumi.Input

	ToStoredProcedureParameterTypePtrOutput() StoredProcedureParameterTypePtrOutput
	ToStoredProcedureParameterTypePtrOutputWithContext(context.Context) StoredProcedureParameterTypePtrOutput
}

type storedProcedureParameterTypePtr string

func StoredProcedureParameterTypePtr(v string) StoredProcedureParameterTypePtrInput {
	return (*storedProcedureParameterTypePtr)(&v)
}

func (*storedProcedureParameterTypePtr) ElementType() reflect.Type {
	return storedProcedureParameterTypePtrType
}

func (in *storedProcedureParameterTypePtr) ToStoredProcedureParameterTypePtrOutput() StoredProcedureParameterTypePtrOutput {
	return pulumi.ToOutput(in).(StoredProcedureParameterTypePtrOutput)
}

func (in *storedProcedureParameterTypePtr) ToStoredProcedureParameterTypePtrOutputWithContext(ctx context.Context) StoredProcedureParameterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StoredProcedureParameterTypePtrOutput)
}

type SybaseAuthenticationType string

const (
	SybaseAuthenticationTypeBasic   = SybaseAuthenticationType("Basic")
	SybaseAuthenticationTypeWindows = SybaseAuthenticationType("Windows")
)

func (SybaseAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*SybaseAuthenticationType)(nil)).Elem()
}

func (e SybaseAuthenticationType) ToSybaseAuthenticationTypeOutput() SybaseAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(SybaseAuthenticationTypeOutput)
}

func (e SybaseAuthenticationType) ToSybaseAuthenticationTypeOutputWithContext(ctx context.Context) SybaseAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SybaseAuthenticationTypeOutput)
}

func (e SybaseAuthenticationType) ToSybaseAuthenticationTypePtrOutput() SybaseAuthenticationTypePtrOutput {
	return e.ToSybaseAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e SybaseAuthenticationType) ToSybaseAuthenticationTypePtrOutputWithContext(ctx context.Context) SybaseAuthenticationTypePtrOutput {
	return SybaseAuthenticationType(e).ToSybaseAuthenticationTypeOutputWithContext(ctx).ToSybaseAuthenticationTypePtrOutputWithContext(ctx)
}

func (e SybaseAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SybaseAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SybaseAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SybaseAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SybaseAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (SybaseAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SybaseAuthenticationType)(nil)).Elem()
}

func (o SybaseAuthenticationTypeOutput) ToSybaseAuthenticationTypeOutput() SybaseAuthenticationTypeOutput {
	return o
}

func (o SybaseAuthenticationTypeOutput) ToSybaseAuthenticationTypeOutputWithContext(ctx context.Context) SybaseAuthenticationTypeOutput {
	return o
}

func (o SybaseAuthenticationTypeOutput) ToSybaseAuthenticationTypePtrOutput() SybaseAuthenticationTypePtrOutput {
	return o.ToSybaseAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o SybaseAuthenticationTypeOutput) ToSybaseAuthenticationTypePtrOutputWithContext(ctx context.Context) SybaseAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SybaseAuthenticationType) *SybaseAuthenticationType {
		return &v
	}).(SybaseAuthenticationTypePtrOutput)
}

func (o SybaseAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SybaseAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SybaseAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SybaseAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SybaseAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SybaseAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SybaseAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (SybaseAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SybaseAuthenticationType)(nil)).Elem()
}

func (o SybaseAuthenticationTypePtrOutput) ToSybaseAuthenticationTypePtrOutput() SybaseAuthenticationTypePtrOutput {
	return o
}

func (o SybaseAuthenticationTypePtrOutput) ToSybaseAuthenticationTypePtrOutputWithContext(ctx context.Context) SybaseAuthenticationTypePtrOutput {
	return o
}

func (o SybaseAuthenticationTypePtrOutput) Elem() SybaseAuthenticationTypeOutput {
	return o.ApplyT(func(v *SybaseAuthenticationType) SybaseAuthenticationType {
		if v != nil {
			return *v
		}
		var ret SybaseAuthenticationType
		return ret
	}).(SybaseAuthenticationTypeOutput)
}

func (o SybaseAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SybaseAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SybaseAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SybaseAuthenticationTypeInput interface {
	pulumi.Input

	ToSybaseAuthenticationTypeOutput() SybaseAuthenticationTypeOutput
	ToSybaseAuthenticationTypeOutputWithContext(context.Context) SybaseAuthenticationTypeOutput
}

var sybaseAuthenticationTypePtrType = reflect.TypeOf((**SybaseAuthenticationType)(nil)).Elem()

type SybaseAuthenticationTypePtrInput interface {
	pulumi.Input

	ToSybaseAuthenticationTypePtrOutput() SybaseAuthenticationTypePtrOutput
	ToSybaseAuthenticationTypePtrOutputWithContext(context.Context) SybaseAuthenticationTypePtrOutput
}

type sybaseAuthenticationTypePtr string

func SybaseAuthenticationTypePtr(v string) SybaseAuthenticationTypePtrInput {
	return (*sybaseAuthenticationTypePtr)(&v)
}

func (*sybaseAuthenticationTypePtr) ElementType() reflect.Type {
	return sybaseAuthenticationTypePtrType
}

func (in *sybaseAuthenticationTypePtr) ToSybaseAuthenticationTypePtrOutput() SybaseAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(SybaseAuthenticationTypePtrOutput)
}

func (in *sybaseAuthenticationTypePtr) ToSybaseAuthenticationTypePtrOutputWithContext(ctx context.Context) SybaseAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SybaseAuthenticationTypePtrOutput)
}

type TeradataAuthenticationType string

const (
	TeradataAuthenticationTypeBasic   = TeradataAuthenticationType("Basic")
	TeradataAuthenticationTypeWindows = TeradataAuthenticationType("Windows")
)

func (TeradataAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*TeradataAuthenticationType)(nil)).Elem()
}

func (e TeradataAuthenticationType) ToTeradataAuthenticationTypeOutput() TeradataAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(TeradataAuthenticationTypeOutput)
}

func (e TeradataAuthenticationType) ToTeradataAuthenticationTypeOutputWithContext(ctx context.Context) TeradataAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TeradataAuthenticationTypeOutput)
}

func (e TeradataAuthenticationType) ToTeradataAuthenticationTypePtrOutput() TeradataAuthenticationTypePtrOutput {
	return e.ToTeradataAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e TeradataAuthenticationType) ToTeradataAuthenticationTypePtrOutputWithContext(ctx context.Context) TeradataAuthenticationTypePtrOutput {
	return TeradataAuthenticationType(e).ToTeradataAuthenticationTypeOutputWithContext(ctx).ToTeradataAuthenticationTypePtrOutputWithContext(ctx)
}

func (e TeradataAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TeradataAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TeradataAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TeradataAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TeradataAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (TeradataAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TeradataAuthenticationType)(nil)).Elem()
}

func (o TeradataAuthenticationTypeOutput) ToTeradataAuthenticationTypeOutput() TeradataAuthenticationTypeOutput {
	return o
}

func (o TeradataAuthenticationTypeOutput) ToTeradataAuthenticationTypeOutputWithContext(ctx context.Context) TeradataAuthenticationTypeOutput {
	return o
}

func (o TeradataAuthenticationTypeOutput) ToTeradataAuthenticationTypePtrOutput() TeradataAuthenticationTypePtrOutput {
	return o.ToTeradataAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o TeradataAuthenticationTypeOutput) ToTeradataAuthenticationTypePtrOutputWithContext(ctx context.Context) TeradataAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TeradataAuthenticationType) *TeradataAuthenticationType {
		return &v
	}).(TeradataAuthenticationTypePtrOutput)
}

func (o TeradataAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TeradataAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TeradataAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TeradataAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TeradataAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TeradataAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TeradataAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (TeradataAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeradataAuthenticationType)(nil)).Elem()
}

func (o TeradataAuthenticationTypePtrOutput) ToTeradataAuthenticationTypePtrOutput() TeradataAuthenticationTypePtrOutput {
	return o
}

func (o TeradataAuthenticationTypePtrOutput) ToTeradataAuthenticationTypePtrOutputWithContext(ctx context.Context) TeradataAuthenticationTypePtrOutput {
	return o
}

func (o TeradataAuthenticationTypePtrOutput) Elem() TeradataAuthenticationTypeOutput {
	return o.ApplyT(func(v *TeradataAuthenticationType) TeradataAuthenticationType {
		if v != nil {
			return *v
		}
		var ret TeradataAuthenticationType
		return ret
	}).(TeradataAuthenticationTypeOutput)
}

func (o TeradataAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TeradataAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TeradataAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TeradataAuthenticationTypeInput interface {
	pulumi.Input

	ToTeradataAuthenticationTypeOutput() TeradataAuthenticationTypeOutput
	ToTeradataAuthenticationTypeOutputWithContext(context.Context) TeradataAuthenticationTypeOutput
}

var teradataAuthenticationTypePtrType = reflect.TypeOf((**TeradataAuthenticationType)(nil)).Elem()

type TeradataAuthenticationTypePtrInput interface {
	pulumi.Input

	ToTeradataAuthenticationTypePtrOutput() TeradataAuthenticationTypePtrOutput
	ToTeradataAuthenticationTypePtrOutputWithContext(context.Context) TeradataAuthenticationTypePtrOutput
}

type teradataAuthenticationTypePtr string

func TeradataAuthenticationTypePtr(v string) TeradataAuthenticationTypePtrInput {
	return (*teradataAuthenticationTypePtr)(&v)
}

func (*teradataAuthenticationTypePtr) ElementType() reflect.Type {
	return teradataAuthenticationTypePtrType
}

func (in *teradataAuthenticationTypePtr) ToTeradataAuthenticationTypePtrOutput() TeradataAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(TeradataAuthenticationTypePtrOutput)
}

func (in *teradataAuthenticationTypePtr) ToTeradataAuthenticationTypePtrOutputWithContext(ctx context.Context) TeradataAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TeradataAuthenticationTypePtrOutput)
}

type TumblingWindowFrequency string

const (
	TumblingWindowFrequencyMinute = TumblingWindowFrequency("Minute")
	TumblingWindowFrequencyHour   = TumblingWindowFrequency("Hour")
	TumblingWindowFrequencyMonth  = TumblingWindowFrequency("Month")
)

func (TumblingWindowFrequency) ElementType() reflect.Type {
	return reflect.TypeOf((*TumblingWindowFrequency)(nil)).Elem()
}

func (e TumblingWindowFrequency) ToTumblingWindowFrequencyOutput() TumblingWindowFrequencyOutput {
	return pulumi.ToOutput(e).(TumblingWindowFrequencyOutput)
}

func (e TumblingWindowFrequency) ToTumblingWindowFrequencyOutputWithContext(ctx context.Context) TumblingWindowFrequencyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TumblingWindowFrequencyOutput)
}

func (e TumblingWindowFrequency) ToTumblingWindowFrequencyPtrOutput() TumblingWindowFrequencyPtrOutput {
	return e.ToTumblingWindowFrequencyPtrOutputWithContext(context.Background())
}

func (e TumblingWindowFrequency) ToTumblingWindowFrequencyPtrOutputWithContext(ctx context.Context) TumblingWindowFrequencyPtrOutput {
	return TumblingWindowFrequency(e).ToTumblingWindowFrequencyOutputWithContext(ctx).ToTumblingWindowFrequencyPtrOutputWithContext(ctx)
}

func (e TumblingWindowFrequency) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TumblingWindowFrequency) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TumblingWindowFrequency) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TumblingWindowFrequency) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TumblingWindowFrequencyOutput struct{ *pulumi.OutputState }

func (TumblingWindowFrequencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TumblingWindowFrequency)(nil)).Elem()
}

func (o TumblingWindowFrequencyOutput) ToTumblingWindowFrequencyOutput() TumblingWindowFrequencyOutput {
	return o
}

func (o TumblingWindowFrequencyOutput) ToTumblingWindowFrequencyOutputWithContext(ctx context.Context) TumblingWindowFrequencyOutput {
	return o
}

func (o TumblingWindowFrequencyOutput) ToTumblingWindowFrequencyPtrOutput() TumblingWindowFrequencyPtrOutput {
	return o.ToTumblingWindowFrequencyPtrOutputWithContext(context.Background())
}

func (o TumblingWindowFrequencyOutput) ToTumblingWindowFrequencyPtrOutputWithContext(ctx context.Context) TumblingWindowFrequencyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TumblingWindowFrequency) *TumblingWindowFrequency {
		return &v
	}).(TumblingWindowFrequencyPtrOutput)
}

func (o TumblingWindowFrequencyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TumblingWindowFrequencyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TumblingWindowFrequency) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TumblingWindowFrequencyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TumblingWindowFrequencyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TumblingWindowFrequency) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TumblingWindowFrequencyPtrOutput struct{ *pulumi.OutputState }

func (TumblingWindowFrequencyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TumblingWindowFrequency)(nil)).Elem()
}

func (o TumblingWindowFrequencyPtrOutput) ToTumblingWindowFrequencyPtrOutput() TumblingWindowFrequencyPtrOutput {
	return o
}

func (o TumblingWindowFrequencyPtrOutput) ToTumblingWindowFrequencyPtrOutputWithContext(ctx context.Context) TumblingWindowFrequencyPtrOutput {
	return o
}

func (o TumblingWindowFrequencyPtrOutput) Elem() TumblingWindowFrequencyOutput {
	return o.ApplyT(func(v *TumblingWindowFrequency) TumblingWindowFrequency {
		if v != nil {
			return *v
		}
		var ret TumblingWindowFrequency
		return ret
	}).(TumblingWindowFrequencyOutput)
}

func (o TumblingWindowFrequencyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TumblingWindowFrequencyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TumblingWindowFrequency) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TumblingWindowFrequencyInput interface {
	pulumi.Input

	ToTumblingWindowFrequencyOutput() TumblingWindowFrequencyOutput
	ToTumblingWindowFrequencyOutputWithContext(context.Context) TumblingWindowFrequencyOutput
}

var tumblingWindowFrequencyPtrType = reflect.TypeOf((**TumblingWindowFrequency)(nil)).Elem()

type TumblingWindowFrequencyPtrInput interface {
	pulumi.Input

	ToTumblingWindowFrequencyPtrOutput() TumblingWindowFrequencyPtrOutput
	ToTumblingWindowFrequencyPtrOutputWithContext(context.Context) TumblingWindowFrequencyPtrOutput
}

type tumblingWindowFrequencyPtr string

func TumblingWindowFrequencyPtr(v string) TumblingWindowFrequencyPtrInput {
	return (*tumblingWindowFrequencyPtr)(&v)
}

func (*tumblingWindowFrequencyPtr) ElementType() reflect.Type {
	return tumblingWindowFrequencyPtrType
}

func (in *tumblingWindowFrequencyPtr) ToTumblingWindowFrequencyPtrOutput() TumblingWindowFrequencyPtrOutput {
	return pulumi.ToOutput(in).(TumblingWindowFrequencyPtrOutput)
}

func (in *tumblingWindowFrequencyPtr) ToTumblingWindowFrequencyPtrOutputWithContext(ctx context.Context) TumblingWindowFrequencyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TumblingWindowFrequencyPtrOutput)
}

type VariableType string

const (
	VariableTypeString = VariableType("String")
	VariableTypeBool   = VariableType("Bool")
	VariableTypeArray  = VariableType("Array")
)

func (VariableType) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableType)(nil)).Elem()
}

func (e VariableType) ToVariableTypeOutput() VariableTypeOutput {
	return pulumi.ToOutput(e).(VariableTypeOutput)
}

func (e VariableType) ToVariableTypeOutputWithContext(ctx context.Context) VariableTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VariableTypeOutput)
}

func (e VariableType) ToVariableTypePtrOutput() VariableTypePtrOutput {
	return e.ToVariableTypePtrOutputWithContext(context.Background())
}

func (e VariableType) ToVariableTypePtrOutputWithContext(ctx context.Context) VariableTypePtrOutput {
	return VariableType(e).ToVariableTypeOutputWithContext(ctx).ToVariableTypePtrOutputWithContext(ctx)
}

func (e VariableType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VariableType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VariableType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VariableType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VariableTypeOutput struct{ *pulumi.OutputState }

func (VariableTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariableType)(nil)).Elem()
}

func (o VariableTypeOutput) ToVariableTypeOutput() VariableTypeOutput {
	return o
}

func (o VariableTypeOutput) ToVariableTypeOutputWithContext(ctx context.Context) VariableTypeOutput {
	return o
}

func (o VariableTypeOutput) ToVariableTypePtrOutput() VariableTypePtrOutput {
	return o.ToVariableTypePtrOutputWithContext(context.Background())
}

func (o VariableTypeOutput) ToVariableTypePtrOutputWithContext(ctx context.Context) VariableTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VariableType) *VariableType {
		return &v
	}).(VariableTypePtrOutput)
}

func (o VariableTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VariableTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VariableType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VariableTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VariableTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VariableType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VariableTypePtrOutput struct{ *pulumi.OutputState }

func (VariableTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableType)(nil)).Elem()
}

func (o VariableTypePtrOutput) ToVariableTypePtrOutput() VariableTypePtrOutput {
	return o
}

func (o VariableTypePtrOutput) ToVariableTypePtrOutputWithContext(ctx context.Context) VariableTypePtrOutput {
	return o
}

func (o VariableTypePtrOutput) Elem() VariableTypeOutput {
	return o.ApplyT(func(v *VariableType) VariableType {
		if v != nil {
			return *v
		}
		var ret VariableType
		return ret
	}).(VariableTypeOutput)
}

func (o VariableTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VariableTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VariableType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VariableTypeInput interface {
	pulumi.Input

	ToVariableTypeOutput() VariableTypeOutput
	ToVariableTypeOutputWithContext(context.Context) VariableTypeOutput
}

var variableTypePtrType = reflect.TypeOf((**VariableType)(nil)).Elem()

type VariableTypePtrInput interface {
	pulumi.Input

	ToVariableTypePtrOutput() VariableTypePtrOutput
	ToVariableTypePtrOutputWithContext(context.Context) VariableTypePtrOutput
}

type variableTypePtr string

func VariableTypePtr(v string) VariableTypePtrInput {
	return (*variableTypePtr)(&v)
}

func (*variableTypePtr) ElementType() reflect.Type {
	return variableTypePtrType
}

func (in *variableTypePtr) ToVariableTypePtrOutput() VariableTypePtrOutput {
	return pulumi.ToOutput(in).(VariableTypePtrOutput)
}

func (in *variableTypePtr) ToVariableTypePtrOutputWithContext(ctx context.Context) VariableTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VariableTypePtrOutput)
}

type WebActivityMethod string

const (
	WebActivityMethodGET    = WebActivityMethod("GET")
	WebActivityMethodPOST   = WebActivityMethod("POST")
	WebActivityMethodPUT    = WebActivityMethod("PUT")
	WebActivityMethodDELETE = WebActivityMethod("DELETE")
)

func (WebActivityMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*WebActivityMethod)(nil)).Elem()
}

func (e WebActivityMethod) ToWebActivityMethodOutput() WebActivityMethodOutput {
	return pulumi.ToOutput(e).(WebActivityMethodOutput)
}

func (e WebActivityMethod) ToWebActivityMethodOutputWithContext(ctx context.Context) WebActivityMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebActivityMethodOutput)
}

func (e WebActivityMethod) ToWebActivityMethodPtrOutput() WebActivityMethodPtrOutput {
	return e.ToWebActivityMethodPtrOutputWithContext(context.Background())
}

func (e WebActivityMethod) ToWebActivityMethodPtrOutputWithContext(ctx context.Context) WebActivityMethodPtrOutput {
	return WebActivityMethod(e).ToWebActivityMethodOutputWithContext(ctx).ToWebActivityMethodPtrOutputWithContext(ctx)
}

func (e WebActivityMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebActivityMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebActivityMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebActivityMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebActivityMethodOutput struct{ *pulumi.OutputState }

func (WebActivityMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebActivityMethod)(nil)).Elem()
}

func (o WebActivityMethodOutput) ToWebActivityMethodOutput() WebActivityMethodOutput {
	return o
}

func (o WebActivityMethodOutput) ToWebActivityMethodOutputWithContext(ctx context.Context) WebActivityMethodOutput {
	return o
}

func (o WebActivityMethodOutput) ToWebActivityMethodPtrOutput() WebActivityMethodPtrOutput {
	return o.ToWebActivityMethodPtrOutputWithContext(context.Background())
}

func (o WebActivityMethodOutput) ToWebActivityMethodPtrOutputWithContext(ctx context.Context) WebActivityMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebActivityMethod) *WebActivityMethod {
		return &v
	}).(WebActivityMethodPtrOutput)
}

func (o WebActivityMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebActivityMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebActivityMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebActivityMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebActivityMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebActivityMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebActivityMethodPtrOutput struct{ *pulumi.OutputState }

func (WebActivityMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebActivityMethod)(nil)).Elem()
}

func (o WebActivityMethodPtrOutput) ToWebActivityMethodPtrOutput() WebActivityMethodPtrOutput {
	return o
}

func (o WebActivityMethodPtrOutput) ToWebActivityMethodPtrOutputWithContext(ctx context.Context) WebActivityMethodPtrOutput {
	return o
}

func (o WebActivityMethodPtrOutput) Elem() WebActivityMethodOutput {
	return o.ApplyT(func(v *WebActivityMethod) WebActivityMethod {
		if v != nil {
			return *v
		}
		var ret WebActivityMethod
		return ret
	}).(WebActivityMethodOutput)
}

func (o WebActivityMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebActivityMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebActivityMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebActivityMethodInput interface {
	pulumi.Input

	ToWebActivityMethodOutput() WebActivityMethodOutput
	ToWebActivityMethodOutputWithContext(context.Context) WebActivityMethodOutput
}

var webActivityMethodPtrType = reflect.TypeOf((**WebActivityMethod)(nil)).Elem()

type WebActivityMethodPtrInput interface {
	pulumi.Input

	ToWebActivityMethodPtrOutput() WebActivityMethodPtrOutput
	ToWebActivityMethodPtrOutputWithContext(context.Context) WebActivityMethodPtrOutput
}

type webActivityMethodPtr string

func WebActivityMethodPtr(v string) WebActivityMethodPtrInput {
	return (*webActivityMethodPtr)(&v)
}

func (*webActivityMethodPtr) ElementType() reflect.Type {
	return webActivityMethodPtrType
}

func (in *webActivityMethodPtr) ToWebActivityMethodPtrOutput() WebActivityMethodPtrOutput {
	return pulumi.ToOutput(in).(WebActivityMethodPtrOutput)
}

func (in *webActivityMethodPtr) ToWebActivityMethodPtrOutputWithContext(ctx context.Context) WebActivityMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebActivityMethodPtrOutput)
}

type WebAuthenticationType string

const (
	WebAuthenticationTypeBasic             = WebAuthenticationType("Basic")
	WebAuthenticationTypeAnonymous         = WebAuthenticationType("Anonymous")
	WebAuthenticationTypeClientCertificate = WebAuthenticationType("ClientCertificate")
)

func (WebAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAuthenticationType)(nil)).Elem()
}

func (e WebAuthenticationType) ToWebAuthenticationTypeOutput() WebAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(WebAuthenticationTypeOutput)
}

func (e WebAuthenticationType) ToWebAuthenticationTypeOutputWithContext(ctx context.Context) WebAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebAuthenticationTypeOutput)
}

func (e WebAuthenticationType) ToWebAuthenticationTypePtrOutput() WebAuthenticationTypePtrOutput {
	return e.ToWebAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e WebAuthenticationType) ToWebAuthenticationTypePtrOutputWithContext(ctx context.Context) WebAuthenticationTypePtrOutput {
	return WebAuthenticationType(e).ToWebAuthenticationTypeOutputWithContext(ctx).ToWebAuthenticationTypePtrOutputWithContext(ctx)
}

func (e WebAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (WebAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAuthenticationType)(nil)).Elem()
}

func (o WebAuthenticationTypeOutput) ToWebAuthenticationTypeOutput() WebAuthenticationTypeOutput {
	return o
}

func (o WebAuthenticationTypeOutput) ToWebAuthenticationTypeOutputWithContext(ctx context.Context) WebAuthenticationTypeOutput {
	return o
}

func (o WebAuthenticationTypeOutput) ToWebAuthenticationTypePtrOutput() WebAuthenticationTypePtrOutput {
	return o.ToWebAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o WebAuthenticationTypeOutput) ToWebAuthenticationTypePtrOutputWithContext(ctx context.Context) WebAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebAuthenticationType) *WebAuthenticationType {
		return &v
	}).(WebAuthenticationTypePtrOutput)
}

func (o WebAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (WebAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAuthenticationType)(nil)).Elem()
}

func (o WebAuthenticationTypePtrOutput) ToWebAuthenticationTypePtrOutput() WebAuthenticationTypePtrOutput {
	return o
}

func (o WebAuthenticationTypePtrOutput) ToWebAuthenticationTypePtrOutputWithContext(ctx context.Context) WebAuthenticationTypePtrOutput {
	return o
}

func (o WebAuthenticationTypePtrOutput) Elem() WebAuthenticationTypeOutput {
	return o.ApplyT(func(v *WebAuthenticationType) WebAuthenticationType {
		if v != nil {
			return *v
		}
		var ret WebAuthenticationType
		return ret
	}).(WebAuthenticationTypeOutput)
}

func (o WebAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebAuthenticationTypeInput interface {
	pulumi.Input

	ToWebAuthenticationTypeOutput() WebAuthenticationTypeOutput
	ToWebAuthenticationTypeOutputWithContext(context.Context) WebAuthenticationTypeOutput
}

var webAuthenticationTypePtrType = reflect.TypeOf((**WebAuthenticationType)(nil)).Elem()

type WebAuthenticationTypePtrInput interface {
	pulumi.Input

	ToWebAuthenticationTypePtrOutput() WebAuthenticationTypePtrOutput
	ToWebAuthenticationTypePtrOutputWithContext(context.Context) WebAuthenticationTypePtrOutput
}

type webAuthenticationTypePtr string

func WebAuthenticationTypePtr(v string) WebAuthenticationTypePtrInput {
	return (*webAuthenticationTypePtr)(&v)
}

func (*webAuthenticationTypePtr) ElementType() reflect.Type {
	return webAuthenticationTypePtrType
}

func (in *webAuthenticationTypePtr) ToWebAuthenticationTypePtrOutput() WebAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(WebAuthenticationTypePtrOutput)
}

func (in *webAuthenticationTypePtr) ToWebAuthenticationTypePtrOutputWithContext(ctx context.Context) WebAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebAuthenticationTypePtrOutput)
}

type WebHookActivityMethod string

const (
	WebHookActivityMethodPOST = WebHookActivityMethod("POST")
)

func (WebHookActivityMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookActivityMethod)(nil)).Elem()
}

func (e WebHookActivityMethod) ToWebHookActivityMethodOutput() WebHookActivityMethodOutput {
	return pulumi.ToOutput(e).(WebHookActivityMethodOutput)
}

func (e WebHookActivityMethod) ToWebHookActivityMethodOutputWithContext(ctx context.Context) WebHookActivityMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebHookActivityMethodOutput)
}

func (e WebHookActivityMethod) ToWebHookActivityMethodPtrOutput() WebHookActivityMethodPtrOutput {
	return e.ToWebHookActivityMethodPtrOutputWithContext(context.Background())
}

func (e WebHookActivityMethod) ToWebHookActivityMethodPtrOutputWithContext(ctx context.Context) WebHookActivityMethodPtrOutput {
	return WebHookActivityMethod(e).ToWebHookActivityMethodOutputWithContext(ctx).ToWebHookActivityMethodPtrOutputWithContext(ctx)
}

func (e WebHookActivityMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebHookActivityMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebHookActivityMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebHookActivityMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebHookActivityMethodOutput struct{ *pulumi.OutputState }

func (WebHookActivityMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookActivityMethod)(nil)).Elem()
}

func (o WebHookActivityMethodOutput) ToWebHookActivityMethodOutput() WebHookActivityMethodOutput {
	return o
}

func (o WebHookActivityMethodOutput) ToWebHookActivityMethodOutputWithContext(ctx context.Context) WebHookActivityMethodOutput {
	return o
}

func (o WebHookActivityMethodOutput) ToWebHookActivityMethodPtrOutput() WebHookActivityMethodPtrOutput {
	return o.ToWebHookActivityMethodPtrOutputWithContext(context.Background())
}

func (o WebHookActivityMethodOutput) ToWebHookActivityMethodPtrOutputWithContext(ctx context.Context) WebHookActivityMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebHookActivityMethod) *WebHookActivityMethod {
		return &v
	}).(WebHookActivityMethodPtrOutput)
}

func (o WebHookActivityMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebHookActivityMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebHookActivityMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebHookActivityMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebHookActivityMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebHookActivityMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebHookActivityMethodPtrOutput struct{ *pulumi.OutputState }

func (WebHookActivityMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebHookActivityMethod)(nil)).Elem()
}

func (o WebHookActivityMethodPtrOutput) ToWebHookActivityMethodPtrOutput() WebHookActivityMethodPtrOutput {
	return o
}

func (o WebHookActivityMethodPtrOutput) ToWebHookActivityMethodPtrOutputWithContext(ctx context.Context) WebHookActivityMethodPtrOutput {
	return o
}

func (o WebHookActivityMethodPtrOutput) Elem() WebHookActivityMethodOutput {
	return o.ApplyT(func(v *WebHookActivityMethod) WebHookActivityMethod {
		if v != nil {
			return *v
		}
		var ret WebHookActivityMethod
		return ret
	}).(WebHookActivityMethodOutput)
}

func (o WebHookActivityMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebHookActivityMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebHookActivityMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebHookActivityMethodInput interface {
	pulumi.Input

	ToWebHookActivityMethodOutput() WebHookActivityMethodOutput
	ToWebHookActivityMethodOutputWithContext(context.Context) WebHookActivityMethodOutput
}

var webHookActivityMethodPtrType = reflect.TypeOf((**WebHookActivityMethod)(nil)).Elem()

type WebHookActivityMethodPtrInput interface {
	pulumi.Input

	ToWebHookActivityMethodPtrOutput() WebHookActivityMethodPtrOutput
	ToWebHookActivityMethodPtrOutputWithContext(context.Context) WebHookActivityMethodPtrOutput
}

type webHookActivityMethodPtr string

func WebHookActivityMethodPtr(v string) WebHookActivityMethodPtrInput {
	return (*webHookActivityMethodPtr)(&v)
}

func (*webHookActivityMethodPtr) ElementType() reflect.Type {
	return webHookActivityMethodPtrType
}

func (in *webHookActivityMethodPtr) ToWebHookActivityMethodPtrOutput() WebHookActivityMethodPtrOutput {
	return pulumi.ToOutput(in).(WebHookActivityMethodPtrOutput)
}

func (in *webHookActivityMethodPtr) ToWebHookActivityMethodPtrOutputWithContext(ctx context.Context) WebHookActivityMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebHookActivityMethodPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFunctionActivityMethodInput)(nil)).Elem(), AzureFunctionActivityMethod("GET"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFunctionActivityMethodPtrInput)(nil)).Elem(), AzureFunctionActivityMethod("GET"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSearchIndexWriteBehaviorTypeInput)(nil)).Elem(), AzureSearchIndexWriteBehaviorType("Merge"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSearchIndexWriteBehaviorTypePtrInput)(nil)).Elem(), AzureSearchIndexWriteBehaviorType("Merge"))
	pulumi.RegisterInputType(reflect.TypeOf((*BlobEventTypesInput)(nil)).Elem(), BlobEventTypes("Microsoft.Storage.BlobCreated"))
	pulumi.RegisterInputType(reflect.TypeOf((*BlobEventTypesPtrInput)(nil)).Elem(), BlobEventTypes("Microsoft.Storage.BlobCreated"))
	pulumi.RegisterInputType(reflect.TypeOf((*CassandraSourceReadConsistencyLevelsInput)(nil)).Elem(), CassandraSourceReadConsistencyLevels("ALL"))
	pulumi.RegisterInputType(reflect.TypeOf((*CassandraSourceReadConsistencyLevelsPtrInput)(nil)).Elem(), CassandraSourceReadConsistencyLevels("ALL"))
	pulumi.RegisterInputType(reflect.TypeOf((*CosmosDbConnectionModeInput)(nil)).Elem(), CosmosDbConnectionMode("Gateway"))
	pulumi.RegisterInputType(reflect.TypeOf((*CosmosDbConnectionModePtrInput)(nil)).Elem(), CosmosDbConnectionMode("Gateway"))
	pulumi.RegisterInputType(reflect.TypeOf((*CosmosDbServicePrincipalCredentialTypeInput)(nil)).Elem(), CosmosDbServicePrincipalCredentialType("ServicePrincipalKey"))
	pulumi.RegisterInputType(reflect.TypeOf((*CosmosDbServicePrincipalCredentialTypePtrInput)(nil)).Elem(), CosmosDbServicePrincipalCredentialType("ServicePrincipalKey"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataFlowComputeTypeInput)(nil)).Elem(), DataFlowComputeType("General"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataFlowComputeTypePtrInput)(nil)).Elem(), DataFlowComputeType("General"))
	pulumi.RegisterInputType(reflect.TypeOf((*DayOfWeekInput)(nil)).Elem(), DayOfWeek("Sunday"))
	pulumi.RegisterInputType(reflect.TypeOf((*DayOfWeekPtrInput)(nil)).Elem(), DayOfWeek("Sunday"))
	pulumi.RegisterInputType(reflect.TypeOf((*DaysOfWeekInput)(nil)).Elem(), DaysOfWeek("Sunday"))
	pulumi.RegisterInputType(reflect.TypeOf((*DaysOfWeekPtrInput)(nil)).Elem(), DaysOfWeek("Sunday"))
	pulumi.RegisterInputType(reflect.TypeOf((*DaysOfWeekArrayInput)(nil)).Elem(), DaysOfWeekArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Db2AuthenticationTypeInput)(nil)).Elem(), Db2AuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*Db2AuthenticationTypePtrInput)(nil)).Elem(), Db2AuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*DependencyConditionInput)(nil)).Elem(), DependencyCondition("Succeeded"))
	pulumi.RegisterInputType(reflect.TypeOf((*DependencyConditionPtrInput)(nil)).Elem(), DependencyCondition("Succeeded"))
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicsSinkWriteBehaviorInput)(nil)).Elem(), DynamicsSinkWriteBehavior("Upsert"))
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicsSinkWriteBehaviorPtrInput)(nil)).Elem(), DynamicsSinkWriteBehavior("Upsert"))
	pulumi.RegisterInputType(reflect.TypeOf((*FactoryIdentityTypeInput)(nil)).Elem(), FactoryIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*FactoryIdentityTypePtrInput)(nil)).Elem(), FactoryIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*FtpAuthenticationTypeInput)(nil)).Elem(), FtpAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*FtpAuthenticationTypePtrInput)(nil)).Elem(), FtpAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalParameterTypeInput)(nil)).Elem(), GlobalParameterType("Object"))
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalParameterTypePtrInput)(nil)).Elem(), GlobalParameterType("Object"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleAdWordsAuthenticationTypeInput)(nil)).Elem(), GoogleAdWordsAuthenticationType("ServiceAuthentication"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleAdWordsAuthenticationTypePtrInput)(nil)).Elem(), GoogleAdWordsAuthenticationType("ServiceAuthentication"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleBigQueryAuthenticationTypeInput)(nil)).Elem(), GoogleBigQueryAuthenticationType("ServiceAuthentication"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleBigQueryAuthenticationTypePtrInput)(nil)).Elem(), GoogleBigQueryAuthenticationType("ServiceAuthentication"))
	pulumi.RegisterInputType(reflect.TypeOf((*HBaseAuthenticationTypeInput)(nil)).Elem(), HBaseAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*HBaseAuthenticationTypePtrInput)(nil)).Elem(), HBaseAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*HDInsightActivityDebugInfoOptionInput)(nil)).Elem(), HDInsightActivityDebugInfoOption("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*HDInsightActivityDebugInfoOptionPtrInput)(nil)).Elem(), HDInsightActivityDebugInfoOption("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*HiveAuthenticationTypeInput)(nil)).Elem(), HiveAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*HiveAuthenticationTypePtrInput)(nil)).Elem(), HiveAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*HiveServerTypeInput)(nil)).Elem(), HiveServerType("HiveServer1"))
	pulumi.RegisterInputType(reflect.TypeOf((*HiveServerTypePtrInput)(nil)).Elem(), HiveServerType("HiveServer1"))
	pulumi.RegisterInputType(reflect.TypeOf((*HiveThriftTransportProtocolInput)(nil)).Elem(), HiveThriftTransportProtocol("Binary"))
	pulumi.RegisterInputType(reflect.TypeOf((*HiveThriftTransportProtocolPtrInput)(nil)).Elem(), HiveThriftTransportProtocol("Binary"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpAuthenticationTypeInput)(nil)).Elem(), HttpAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*HttpAuthenticationTypePtrInput)(nil)).Elem(), HttpAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*ImpalaAuthenticationTypeInput)(nil)).Elem(), ImpalaAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*ImpalaAuthenticationTypePtrInput)(nil)).Elem(), ImpalaAuthenticationType("Anonymous"))
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
	pulumi.RegisterInputType(reflect.TypeOf((*MongoDbAuthenticationTypeInput)(nil)).Elem(), MongoDbAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*MongoDbAuthenticationTypePtrInput)(nil)).Elem(), MongoDbAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*ODataAadServicePrincipalCredentialTypeInput)(nil)).Elem(), ODataAadServicePrincipalCredentialType("ServicePrincipalKey"))
	pulumi.RegisterInputType(reflect.TypeOf((*ODataAadServicePrincipalCredentialTypePtrInput)(nil)).Elem(), ODataAadServicePrincipalCredentialType("ServicePrincipalKey"))
	pulumi.RegisterInputType(reflect.TypeOf((*ODataAuthenticationTypeInput)(nil)).Elem(), ODataAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*ODataAuthenticationTypePtrInput)(nil)).Elem(), ODataAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTypeInput)(nil)).Elem(), ParameterType("Object"))
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTypePtrInput)(nil)).Elem(), ParameterType("Object"))
	pulumi.RegisterInputType(reflect.TypeOf((*PhoenixAuthenticationTypeInput)(nil)).Elem(), PhoenixAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*PhoenixAuthenticationTypePtrInput)(nil)).Elem(), PhoenixAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolybaseSettingsRejectTypeInput)(nil)).Elem(), PolybaseSettingsRejectType("value"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolybaseSettingsRejectTypePtrInput)(nil)).Elem(), PolybaseSettingsRejectType("value"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrestoAuthenticationTypeInput)(nil)).Elem(), PrestoAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrestoAuthenticationTypePtrInput)(nil)).Elem(), PrestoAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessPtrInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*RecurrenceFrequencyInput)(nil)).Elem(), RecurrenceFrequency("NotSpecified"))
	pulumi.RegisterInputType(reflect.TypeOf((*RecurrenceFrequencyPtrInput)(nil)).Elem(), RecurrenceFrequency("NotSpecified"))
	pulumi.RegisterInputType(reflect.TypeOf((*RestServiceAuthenticationTypeInput)(nil)).Elem(), RestServiceAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*RestServiceAuthenticationTypePtrInput)(nil)).Elem(), RestServiceAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*SalesforceSinkWriteBehaviorInput)(nil)).Elem(), SalesforceSinkWriteBehavior("Insert"))
	pulumi.RegisterInputType(reflect.TypeOf((*SalesforceSinkWriteBehaviorPtrInput)(nil)).Elem(), SalesforceSinkWriteBehavior("Insert"))
	pulumi.RegisterInputType(reflect.TypeOf((*SalesforceSourceReadBehaviorInput)(nil)).Elem(), SalesforceSourceReadBehavior("Query"))
	pulumi.RegisterInputType(reflect.TypeOf((*SalesforceSourceReadBehaviorPtrInput)(nil)).Elem(), SalesforceSourceReadBehavior("Query"))
	pulumi.RegisterInputType(reflect.TypeOf((*SapCloudForCustomerSinkWriteBehaviorInput)(nil)).Elem(), SapCloudForCustomerSinkWriteBehavior("Insert"))
	pulumi.RegisterInputType(reflect.TypeOf((*SapCloudForCustomerSinkWriteBehaviorPtrInput)(nil)).Elem(), SapCloudForCustomerSinkWriteBehavior("Insert"))
	pulumi.RegisterInputType(reflect.TypeOf((*SapHanaAuthenticationTypeInput)(nil)).Elem(), SapHanaAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SapHanaAuthenticationTypePtrInput)(nil)).Elem(), SapHanaAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceNowAuthenticationTypeInput)(nil)).Elem(), ServiceNowAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceNowAuthenticationTypePtrInput)(nil)).Elem(), ServiceNowAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SftpAuthenticationTypeInput)(nil)).Elem(), SftpAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SftpAuthenticationTypePtrInput)(nil)).Elem(), SftpAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SparkAuthenticationTypeInput)(nil)).Elem(), SparkAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*SparkAuthenticationTypePtrInput)(nil)).Elem(), SparkAuthenticationType("Anonymous"))
	pulumi.RegisterInputType(reflect.TypeOf((*SparkServerTypeInput)(nil)).Elem(), SparkServerType("SharkServer"))
	pulumi.RegisterInputType(reflect.TypeOf((*SparkServerTypePtrInput)(nil)).Elem(), SparkServerType("SharkServer"))
	pulumi.RegisterInputType(reflect.TypeOf((*SparkThriftTransportProtocolInput)(nil)).Elem(), SparkThriftTransportProtocol("Binary"))
	pulumi.RegisterInputType(reflect.TypeOf((*SparkThriftTransportProtocolPtrInput)(nil)).Elem(), SparkThriftTransportProtocol("Binary"))
	pulumi.RegisterInputType(reflect.TypeOf((*SqlAlwaysEncryptedAkvAuthTypeInput)(nil)).Elem(), SqlAlwaysEncryptedAkvAuthType("ServicePrincipal"))
	pulumi.RegisterInputType(reflect.TypeOf((*SqlAlwaysEncryptedAkvAuthTypePtrInput)(nil)).Elem(), SqlAlwaysEncryptedAkvAuthType("ServicePrincipal"))
	pulumi.RegisterInputType(reflect.TypeOf((*SsisLogLocationTypeInput)(nil)).Elem(), SsisLogLocationType("File"))
	pulumi.RegisterInputType(reflect.TypeOf((*SsisLogLocationTypePtrInput)(nil)).Elem(), SsisLogLocationType("File"))
	pulumi.RegisterInputType(reflect.TypeOf((*SsisPackageLocationTypeInput)(nil)).Elem(), SsisPackageLocationType("SSISDB"))
	pulumi.RegisterInputType(reflect.TypeOf((*SsisPackageLocationTypePtrInput)(nil)).Elem(), SsisPackageLocationType("SSISDB"))
	pulumi.RegisterInputType(reflect.TypeOf((*StoredProcedureParameterTypeInput)(nil)).Elem(), StoredProcedureParameterType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*StoredProcedureParameterTypePtrInput)(nil)).Elem(), StoredProcedureParameterType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*SybaseAuthenticationTypeInput)(nil)).Elem(), SybaseAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SybaseAuthenticationTypePtrInput)(nil)).Elem(), SybaseAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*TeradataAuthenticationTypeInput)(nil)).Elem(), TeradataAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*TeradataAuthenticationTypePtrInput)(nil)).Elem(), TeradataAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*TumblingWindowFrequencyInput)(nil)).Elem(), TumblingWindowFrequency("Minute"))
	pulumi.RegisterInputType(reflect.TypeOf((*TumblingWindowFrequencyPtrInput)(nil)).Elem(), TumblingWindowFrequency("Minute"))
	pulumi.RegisterInputType(reflect.TypeOf((*VariableTypeInput)(nil)).Elem(), VariableType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*VariableTypePtrInput)(nil)).Elem(), VariableType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebActivityMethodInput)(nil)).Elem(), WebActivityMethod("GET"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebActivityMethodPtrInput)(nil)).Elem(), WebActivityMethod("GET"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebAuthenticationTypeInput)(nil)).Elem(), WebAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebAuthenticationTypePtrInput)(nil)).Elem(), WebAuthenticationType("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebHookActivityMethodInput)(nil)).Elem(), WebHookActivityMethod("POST"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebHookActivityMethodPtrInput)(nil)).Elem(), WebHookActivityMethod("POST"))
	pulumi.RegisterOutputType(AzureFunctionActivityMethodOutput{})
	pulumi.RegisterOutputType(AzureFunctionActivityMethodPtrOutput{})
	pulumi.RegisterOutputType(AzureSearchIndexWriteBehaviorTypeOutput{})
	pulumi.RegisterOutputType(AzureSearchIndexWriteBehaviorTypePtrOutput{})
	pulumi.RegisterOutputType(BlobEventTypesOutput{})
	pulumi.RegisterOutputType(BlobEventTypesPtrOutput{})
	pulumi.RegisterOutputType(CassandraSourceReadConsistencyLevelsOutput{})
	pulumi.RegisterOutputType(CassandraSourceReadConsistencyLevelsPtrOutput{})
	pulumi.RegisterOutputType(CosmosDbConnectionModeOutput{})
	pulumi.RegisterOutputType(CosmosDbConnectionModePtrOutput{})
	pulumi.RegisterOutputType(CosmosDbServicePrincipalCredentialTypeOutput{})
	pulumi.RegisterOutputType(CosmosDbServicePrincipalCredentialTypePtrOutput{})
	pulumi.RegisterOutputType(DataFlowComputeTypeOutput{})
	pulumi.RegisterOutputType(DataFlowComputeTypePtrOutput{})
	pulumi.RegisterOutputType(DayOfWeekOutput{})
	pulumi.RegisterOutputType(DayOfWeekPtrOutput{})
	pulumi.RegisterOutputType(DaysOfWeekOutput{})
	pulumi.RegisterOutputType(DaysOfWeekPtrOutput{})
	pulumi.RegisterOutputType(DaysOfWeekArrayOutput{})
	pulumi.RegisterOutputType(Db2AuthenticationTypeOutput{})
	pulumi.RegisterOutputType(Db2AuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(DependencyConditionOutput{})
	pulumi.RegisterOutputType(DependencyConditionPtrOutput{})
	pulumi.RegisterOutputType(DynamicsSinkWriteBehaviorOutput{})
	pulumi.RegisterOutputType(DynamicsSinkWriteBehaviorPtrOutput{})
	pulumi.RegisterOutputType(FactoryIdentityTypeOutput{})
	pulumi.RegisterOutputType(FactoryIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(FtpAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(FtpAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(GlobalParameterTypeOutput{})
	pulumi.RegisterOutputType(GlobalParameterTypePtrOutput{})
	pulumi.RegisterOutputType(GoogleAdWordsAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(GoogleAdWordsAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(GoogleBigQueryAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(GoogleBigQueryAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(HBaseAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(HBaseAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(HDInsightActivityDebugInfoOptionOutput{})
	pulumi.RegisterOutputType(HDInsightActivityDebugInfoOptionPtrOutput{})
	pulumi.RegisterOutputType(HiveAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(HiveAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(HiveServerTypeOutput{})
	pulumi.RegisterOutputType(HiveServerTypePtrOutput{})
	pulumi.RegisterOutputType(HiveThriftTransportProtocolOutput{})
	pulumi.RegisterOutputType(HiveThriftTransportProtocolPtrOutput{})
	pulumi.RegisterOutputType(HttpAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(HttpAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(ImpalaAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(ImpalaAuthenticationTypePtrOutput{})
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
	pulumi.RegisterOutputType(MongoDbAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(MongoDbAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(ODataAadServicePrincipalCredentialTypeOutput{})
	pulumi.RegisterOutputType(ODataAadServicePrincipalCredentialTypePtrOutput{})
	pulumi.RegisterOutputType(ODataAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(ODataAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(ParameterTypeOutput{})
	pulumi.RegisterOutputType(ParameterTypePtrOutput{})
	pulumi.RegisterOutputType(PhoenixAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(PhoenixAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(PolybaseSettingsRejectTypeOutput{})
	pulumi.RegisterOutputType(PolybaseSettingsRejectTypePtrOutput{})
	pulumi.RegisterOutputType(PrestoAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(PrestoAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(RecurrenceFrequencyOutput{})
	pulumi.RegisterOutputType(RecurrenceFrequencyPtrOutput{})
	pulumi.RegisterOutputType(RestServiceAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(RestServiceAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(SalesforceSinkWriteBehaviorOutput{})
	pulumi.RegisterOutputType(SalesforceSinkWriteBehaviorPtrOutput{})
	pulumi.RegisterOutputType(SalesforceSourceReadBehaviorOutput{})
	pulumi.RegisterOutputType(SalesforceSourceReadBehaviorPtrOutput{})
	pulumi.RegisterOutputType(SapCloudForCustomerSinkWriteBehaviorOutput{})
	pulumi.RegisterOutputType(SapCloudForCustomerSinkWriteBehaviorPtrOutput{})
	pulumi.RegisterOutputType(SapHanaAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(SapHanaAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(ServiceNowAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(ServiceNowAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(SftpAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(SftpAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(SparkAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(SparkAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(SparkServerTypeOutput{})
	pulumi.RegisterOutputType(SparkServerTypePtrOutput{})
	pulumi.RegisterOutputType(SparkThriftTransportProtocolOutput{})
	pulumi.RegisterOutputType(SparkThriftTransportProtocolPtrOutput{})
	pulumi.RegisterOutputType(SqlAlwaysEncryptedAkvAuthTypeOutput{})
	pulumi.RegisterOutputType(SqlAlwaysEncryptedAkvAuthTypePtrOutput{})
	pulumi.RegisterOutputType(SsisLogLocationTypeOutput{})
	pulumi.RegisterOutputType(SsisLogLocationTypePtrOutput{})
	pulumi.RegisterOutputType(SsisPackageLocationTypeOutput{})
	pulumi.RegisterOutputType(SsisPackageLocationTypePtrOutput{})
	pulumi.RegisterOutputType(StoredProcedureParameterTypeOutput{})
	pulumi.RegisterOutputType(StoredProcedureParameterTypePtrOutput{})
	pulumi.RegisterOutputType(SybaseAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(SybaseAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(TeradataAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(TeradataAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(TumblingWindowFrequencyOutput{})
	pulumi.RegisterOutputType(TumblingWindowFrequencyPtrOutput{})
	pulumi.RegisterOutputType(VariableTypeOutput{})
	pulumi.RegisterOutputType(VariableTypePtrOutput{})
	pulumi.RegisterOutputType(WebActivityMethodOutput{})
	pulumi.RegisterOutputType(WebActivityMethodPtrOutput{})
	pulumi.RegisterOutputType(WebAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(WebAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(WebHookActivityMethodOutput{})
	pulumi.RegisterOutputType(WebHookActivityMethodPtrOutput{})
}
