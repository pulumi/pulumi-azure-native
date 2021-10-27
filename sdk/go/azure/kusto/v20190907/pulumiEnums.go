


package v20190907

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureSkuName string

const (
	AzureSkuName_Standard_DS13_v2_1TB_PS    = AzureSkuName("Standard_DS13_v2+1TB_PS")
	AzureSkuName_Standard_DS13_v2_2TB_PS    = AzureSkuName("Standard_DS13_v2+2TB_PS")
	AzureSkuName_Standard_DS14_v2_3TB_PS    = AzureSkuName("Standard_DS14_v2+3TB_PS")
	AzureSkuName_Standard_DS14_v2_4TB_PS    = AzureSkuName("Standard_DS14_v2+4TB_PS")
	AzureSkuName_Standard_D13_v2            = AzureSkuName("Standard_D13_v2")
	AzureSkuName_Standard_D14_v2            = AzureSkuName("Standard_D14_v2")
	AzureSkuName_Standard_L8s               = AzureSkuName("Standard_L8s")
	AzureSkuName_Standard_L16s              = AzureSkuName("Standard_L16s")
	AzureSkuName_Standard_D11_v2            = AzureSkuName("Standard_D11_v2")
	AzureSkuName_Standard_D12_v2            = AzureSkuName("Standard_D12_v2")
	AzureSkuName_Standard_L4s               = AzureSkuName("Standard_L4s")
	AzureSkuName_Dev_No_SLA_Standard_D11_v2 = AzureSkuName("Dev(No SLA)_Standard_D11_v2")
)

func (AzureSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuName)(nil)).Elem()
}

func (e AzureSkuName) ToAzureSkuNameOutput() AzureSkuNameOutput {
	return pulumi.ToOutput(e).(AzureSkuNameOutput)
}

func (e AzureSkuName) ToAzureSkuNameOutputWithContext(ctx context.Context) AzureSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSkuNameOutput)
}

func (e AzureSkuName) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return e.ToAzureSkuNamePtrOutputWithContext(context.Background())
}

func (e AzureSkuName) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return AzureSkuName(e).ToAzureSkuNameOutputWithContext(ctx).ToAzureSkuNamePtrOutputWithContext(ctx)
}

func (e AzureSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSkuNameOutput struct{ *pulumi.OutputState }

func (AzureSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuName)(nil)).Elem()
}

func (o AzureSkuNameOutput) ToAzureSkuNameOutput() AzureSkuNameOutput {
	return o
}

func (o AzureSkuNameOutput) ToAzureSkuNameOutputWithContext(ctx context.Context) AzureSkuNameOutput {
	return o
}

func (o AzureSkuNameOutput) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return o.ToAzureSkuNamePtrOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuName) *AzureSkuName {
		return &v
	}).(AzureSkuNamePtrOutput)
}

func (o AzureSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSkuNamePtrOutput struct{ *pulumi.OutputState }

func (AzureSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuName)(nil)).Elem()
}

func (o AzureSkuNamePtrOutput) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return o
}

func (o AzureSkuNamePtrOutput) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return o
}

func (o AzureSkuNamePtrOutput) Elem() AzureSkuNameOutput {
	return o.ApplyT(func(v *AzureSkuName) AzureSkuName {
		if v != nil {
			return *v
		}
		var ret AzureSkuName
		return ret
	}).(AzureSkuNameOutput)
}

func (o AzureSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureSkuNameInput interface {
	pulumi.Input

	ToAzureSkuNameOutput() AzureSkuNameOutput
	ToAzureSkuNameOutputWithContext(context.Context) AzureSkuNameOutput
}

var azureSkuNamePtrType = reflect.TypeOf((**AzureSkuName)(nil)).Elem()

type AzureSkuNamePtrInput interface {
	pulumi.Input

	ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput
	ToAzureSkuNamePtrOutputWithContext(context.Context) AzureSkuNamePtrOutput
}

type azureSkuNamePtr string

func AzureSkuNamePtr(v string) AzureSkuNamePtrInput {
	return (*azureSkuNamePtr)(&v)
}

func (*azureSkuNamePtr) ElementType() reflect.Type {
	return azureSkuNamePtrType
}

func (in *azureSkuNamePtr) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return pulumi.ToOutput(in).(AzureSkuNamePtrOutput)
}

func (in *azureSkuNamePtr) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSkuNamePtrOutput)
}

type AzureSkuTier string

const (
	AzureSkuTierBasic    = AzureSkuTier("Basic")
	AzureSkuTierStandard = AzureSkuTier("Standard")
)

func (AzureSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuTier)(nil)).Elem()
}

func (e AzureSkuTier) ToAzureSkuTierOutput() AzureSkuTierOutput {
	return pulumi.ToOutput(e).(AzureSkuTierOutput)
}

func (e AzureSkuTier) ToAzureSkuTierOutputWithContext(ctx context.Context) AzureSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSkuTierOutput)
}

func (e AzureSkuTier) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return e.ToAzureSkuTierPtrOutputWithContext(context.Background())
}

func (e AzureSkuTier) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return AzureSkuTier(e).ToAzureSkuTierOutputWithContext(ctx).ToAzureSkuTierPtrOutputWithContext(ctx)
}

func (e AzureSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSkuTierOutput struct{ *pulumi.OutputState }

func (AzureSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuTier)(nil)).Elem()
}

func (o AzureSkuTierOutput) ToAzureSkuTierOutput() AzureSkuTierOutput {
	return o
}

func (o AzureSkuTierOutput) ToAzureSkuTierOutputWithContext(ctx context.Context) AzureSkuTierOutput {
	return o
}

func (o AzureSkuTierOutput) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return o.ToAzureSkuTierPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuTier) *AzureSkuTier {
		return &v
	}).(AzureSkuTierPtrOutput)
}

func (o AzureSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSkuTierPtrOutput struct{ *pulumi.OutputState }

func (AzureSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuTier)(nil)).Elem()
}

func (o AzureSkuTierPtrOutput) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return o
}

func (o AzureSkuTierPtrOutput) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return o
}

func (o AzureSkuTierPtrOutput) Elem() AzureSkuTierOutput {
	return o.ApplyT(func(v *AzureSkuTier) AzureSkuTier {
		if v != nil {
			return *v
		}
		var ret AzureSkuTier
		return ret
	}).(AzureSkuTierOutput)
}

func (o AzureSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureSkuTierInput interface {
	pulumi.Input

	ToAzureSkuTierOutput() AzureSkuTierOutput
	ToAzureSkuTierOutputWithContext(context.Context) AzureSkuTierOutput
}

var azureSkuTierPtrType = reflect.TypeOf((**AzureSkuTier)(nil)).Elem()

type AzureSkuTierPtrInput interface {
	pulumi.Input

	ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput
	ToAzureSkuTierPtrOutputWithContext(context.Context) AzureSkuTierPtrOutput
}

type azureSkuTierPtr string

func AzureSkuTierPtr(v string) AzureSkuTierPtrInput {
	return (*azureSkuTierPtr)(&v)
}

func (*azureSkuTierPtr) ElementType() reflect.Type {
	return azureSkuTierPtrType
}

func (in *azureSkuTierPtr) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return pulumi.ToOutput(in).(AzureSkuTierPtrOutput)
}

func (in *azureSkuTierPtr) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSkuTierPtrOutput)
}

type DataFormat string

const (
	DataFormatMULTIJSON  = DataFormat("MULTIJSON")
	DataFormatJSON       = DataFormat("JSON")
	DataFormatCSV        = DataFormat("CSV")
	DataFormatTSV        = DataFormat("TSV")
	DataFormatSCSV       = DataFormat("SCSV")
	DataFormatSOHSV      = DataFormat("SOHSV")
	DataFormatPSV        = DataFormat("PSV")
	DataFormatTXT        = DataFormat("TXT")
	DataFormatRAW        = DataFormat("RAW")
	DataFormatSINGLEJSON = DataFormat("SINGLEJSON")
	DataFormatAVRO       = DataFormat("AVRO")
	DataFormatTSVE       = DataFormat("TSVE")
)

func (DataFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFormat)(nil)).Elem()
}

func (e DataFormat) ToDataFormatOutput() DataFormatOutput {
	return pulumi.ToOutput(e).(DataFormatOutput)
}

func (e DataFormat) ToDataFormatOutputWithContext(ctx context.Context) DataFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataFormatOutput)
}

func (e DataFormat) ToDataFormatPtrOutput() DataFormatPtrOutput {
	return e.ToDataFormatPtrOutputWithContext(context.Background())
}

func (e DataFormat) ToDataFormatPtrOutputWithContext(ctx context.Context) DataFormatPtrOutput {
	return DataFormat(e).ToDataFormatOutputWithContext(ctx).ToDataFormatPtrOutputWithContext(ctx)
}

func (e DataFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataFormatOutput struct{ *pulumi.OutputState }

func (DataFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFormat)(nil)).Elem()
}

func (o DataFormatOutput) ToDataFormatOutput() DataFormatOutput {
	return o
}

func (o DataFormatOutput) ToDataFormatOutputWithContext(ctx context.Context) DataFormatOutput {
	return o
}

func (o DataFormatOutput) ToDataFormatPtrOutput() DataFormatPtrOutput {
	return o.ToDataFormatPtrOutputWithContext(context.Background())
}

func (o DataFormatOutput) ToDataFormatPtrOutputWithContext(ctx context.Context) DataFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataFormat) *DataFormat {
		return &v
	}).(DataFormatPtrOutput)
}

func (o DataFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataFormatPtrOutput struct{ *pulumi.OutputState }

func (DataFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataFormat)(nil)).Elem()
}

func (o DataFormatPtrOutput) ToDataFormatPtrOutput() DataFormatPtrOutput {
	return o
}

func (o DataFormatPtrOutput) ToDataFormatPtrOutputWithContext(ctx context.Context) DataFormatPtrOutput {
	return o
}

func (o DataFormatPtrOutput) Elem() DataFormatOutput {
	return o.ApplyT(func(v *DataFormat) DataFormat {
		if v != nil {
			return *v
		}
		var ret DataFormat
		return ret
	}).(DataFormatOutput)
}

func (o DataFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataFormatInput interface {
	pulumi.Input

	ToDataFormatOutput() DataFormatOutput
	ToDataFormatOutputWithContext(context.Context) DataFormatOutput
}

var dataFormatPtrType = reflect.TypeOf((**DataFormat)(nil)).Elem()

type DataFormatPtrInput interface {
	pulumi.Input

	ToDataFormatPtrOutput() DataFormatPtrOutput
	ToDataFormatPtrOutputWithContext(context.Context) DataFormatPtrOutput
}

type dataFormatPtr string

func DataFormatPtr(v string) DataFormatPtrInput {
	return (*dataFormatPtr)(&v)
}

func (*dataFormatPtr) ElementType() reflect.Type {
	return dataFormatPtrType
}

func (in *dataFormatPtr) ToDataFormatPtrOutput() DataFormatPtrOutput {
	return pulumi.ToOutput(in).(DataFormatPtrOutput)
}

func (in *dataFormatPtr) ToDataFormatPtrOutputWithContext(ctx context.Context) DataFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataFormatPtrOutput)
}

type DefaultPrincipalsModificationKind string

const (
	DefaultPrincipalsModificationKindUnion   = DefaultPrincipalsModificationKind("Union")
	DefaultPrincipalsModificationKindReplace = DefaultPrincipalsModificationKind("Replace")
	DefaultPrincipalsModificationKindNone    = DefaultPrincipalsModificationKind("None")
)

func (DefaultPrincipalsModificationKind) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultPrincipalsModificationKind)(nil)).Elem()
}

func (e DefaultPrincipalsModificationKind) ToDefaultPrincipalsModificationKindOutput() DefaultPrincipalsModificationKindOutput {
	return pulumi.ToOutput(e).(DefaultPrincipalsModificationKindOutput)
}

func (e DefaultPrincipalsModificationKind) ToDefaultPrincipalsModificationKindOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultPrincipalsModificationKindOutput)
}

func (e DefaultPrincipalsModificationKind) ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput {
	return e.ToDefaultPrincipalsModificationKindPtrOutputWithContext(context.Background())
}

func (e DefaultPrincipalsModificationKind) ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindPtrOutput {
	return DefaultPrincipalsModificationKind(e).ToDefaultPrincipalsModificationKindOutputWithContext(ctx).ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx)
}

func (e DefaultPrincipalsModificationKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultPrincipalsModificationKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultPrincipalsModificationKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultPrincipalsModificationKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultPrincipalsModificationKindOutput struct{ *pulumi.OutputState }

func (DefaultPrincipalsModificationKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultPrincipalsModificationKind)(nil)).Elem()
}

func (o DefaultPrincipalsModificationKindOutput) ToDefaultPrincipalsModificationKindOutput() DefaultPrincipalsModificationKindOutput {
	return o
}

func (o DefaultPrincipalsModificationKindOutput) ToDefaultPrincipalsModificationKindOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindOutput {
	return o
}

func (o DefaultPrincipalsModificationKindOutput) ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput {
	return o.ToDefaultPrincipalsModificationKindPtrOutputWithContext(context.Background())
}

func (o DefaultPrincipalsModificationKindOutput) ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultPrincipalsModificationKind) *DefaultPrincipalsModificationKind {
		return &v
	}).(DefaultPrincipalsModificationKindPtrOutput)
}

func (o DefaultPrincipalsModificationKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultPrincipalsModificationKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultPrincipalsModificationKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultPrincipalsModificationKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultPrincipalsModificationKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultPrincipalsModificationKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultPrincipalsModificationKindPtrOutput struct{ *pulumi.OutputState }

func (DefaultPrincipalsModificationKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPrincipalsModificationKind)(nil)).Elem()
}

func (o DefaultPrincipalsModificationKindPtrOutput) ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput {
	return o
}

func (o DefaultPrincipalsModificationKindPtrOutput) ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindPtrOutput {
	return o
}

func (o DefaultPrincipalsModificationKindPtrOutput) Elem() DefaultPrincipalsModificationKindOutput {
	return o.ApplyT(func(v *DefaultPrincipalsModificationKind) DefaultPrincipalsModificationKind {
		if v != nil {
			return *v
		}
		var ret DefaultPrincipalsModificationKind
		return ret
	}).(DefaultPrincipalsModificationKindOutput)
}

func (o DefaultPrincipalsModificationKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultPrincipalsModificationKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultPrincipalsModificationKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DefaultPrincipalsModificationKindInput interface {
	pulumi.Input

	ToDefaultPrincipalsModificationKindOutput() DefaultPrincipalsModificationKindOutput
	ToDefaultPrincipalsModificationKindOutputWithContext(context.Context) DefaultPrincipalsModificationKindOutput
}

var defaultPrincipalsModificationKindPtrType = reflect.TypeOf((**DefaultPrincipalsModificationKind)(nil)).Elem()

type DefaultPrincipalsModificationKindPtrInput interface {
	pulumi.Input

	ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput
	ToDefaultPrincipalsModificationKindPtrOutputWithContext(context.Context) DefaultPrincipalsModificationKindPtrOutput
}

type defaultPrincipalsModificationKindPtr string

func DefaultPrincipalsModificationKindPtr(v string) DefaultPrincipalsModificationKindPtrInput {
	return (*defaultPrincipalsModificationKindPtr)(&v)
}

func (*defaultPrincipalsModificationKindPtr) ElementType() reflect.Type {
	return defaultPrincipalsModificationKindPtrType
}

func (in *defaultPrincipalsModificationKindPtr) ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput {
	return pulumi.ToOutput(in).(DefaultPrincipalsModificationKindPtrOutput)
}

func (in *defaultPrincipalsModificationKindPtr) ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultPrincipalsModificationKindPtrOutput)
}

type IdentityType string

const (
	IdentityTypeNone           = IdentityType("None")
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

type Kind string

const (
	KindReadWrite         = Kind("ReadWrite")
	KindReadOnlyFollowing = Kind("ReadOnlyFollowing")
)

func (Kind) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (e Kind) ToKindOutput() KindOutput {
	return pulumi.ToOutput(e).(KindOutput)
}

func (e Kind) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KindOutput)
}

func (e Kind) ToKindPtrOutput() KindPtrOutput {
	return e.ToKindPtrOutputWithContext(context.Background())
}

func (e Kind) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return Kind(e).ToKindOutputWithContext(ctx).ToKindPtrOutputWithContext(ctx)
}

func (e Kind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Kind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KindOutput struct{ *pulumi.OutputState }

func (KindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (o KindOutput) ToKindOutput() KindOutput {
	return o
}

func (o KindOutput) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return o
}

func (o KindOutput) ToKindPtrOutput() KindPtrOutput {
	return o.ToKindPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Kind) *Kind {
		return &v
	}).(KindPtrOutput)
}

func (o KindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KindPtrOutput struct{ *pulumi.OutputState }

func (KindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind)(nil)).Elem()
}

func (o KindPtrOutput) ToKindPtrOutput() KindPtrOutput {
	return o
}

func (o KindPtrOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o
}

func (o KindPtrOutput) Elem() KindOutput {
	return o.ApplyT(func(v *Kind) Kind {
		if v != nil {
			return *v
		}
		var ret Kind
		return ret
	}).(KindOutput)
}

func (o KindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Kind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KindInput interface {
	pulumi.Input

	ToKindOutput() KindOutput
	ToKindOutputWithContext(context.Context) KindOutput
}

var kindPtrType = reflect.TypeOf((**Kind)(nil)).Elem()

type KindPtrInput interface {
	pulumi.Input

	ToKindPtrOutput() KindPtrOutput
	ToKindPtrOutputWithContext(context.Context) KindPtrOutput
}

type kindPtr string

func KindPtr(v string) KindPtrInput {
	return (*kindPtr)(&v)
}

func (*kindPtr) ElementType() reflect.Type {
	return kindPtrType
}

func (in *kindPtr) ToKindPtrOutput() KindPtrOutput {
	return pulumi.ToOutput(in).(KindPtrOutput)
}

func (in *kindPtr) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KindPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSkuNameInput)(nil)).Elem(), AzureSkuName("Standard_DS13_v2+1TB_PS"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSkuNamePtrInput)(nil)).Elem(), AzureSkuName("Standard_DS13_v2+1TB_PS"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSkuTierInput)(nil)).Elem(), AzureSkuTier("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSkuTierPtrInput)(nil)).Elem(), AzureSkuTier("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataFormatInput)(nil)).Elem(), DataFormat("MULTIJSON"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataFormatPtrInput)(nil)).Elem(), DataFormat("MULTIJSON"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPrincipalsModificationKindInput)(nil)).Elem(), DefaultPrincipalsModificationKind("Union"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPrincipalsModificationKindPtrInput)(nil)).Elem(), DefaultPrincipalsModificationKind("Union"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypeInput)(nil)).Elem(), IdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypePtrInput)(nil)).Elem(), IdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindInput)(nil)).Elem(), Kind("ReadWrite"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindPtrInput)(nil)).Elem(), Kind("ReadWrite"))
	pulumi.RegisterOutputType(AzureSkuNameOutput{})
	pulumi.RegisterOutputType(AzureSkuNamePtrOutput{})
	pulumi.RegisterOutputType(AzureSkuTierOutput{})
	pulumi.RegisterOutputType(AzureSkuTierPtrOutput{})
	pulumi.RegisterOutputType(DataFormatOutput{})
	pulumi.RegisterOutputType(DataFormatPtrOutput{})
	pulumi.RegisterOutputType(DefaultPrincipalsModificationKindOutput{})
	pulumi.RegisterOutputType(DefaultPrincipalsModificationKindPtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
}
