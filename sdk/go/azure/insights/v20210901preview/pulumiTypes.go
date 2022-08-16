


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ColumnDefinition struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ColumnDefinitionInput interface {
	pulumi.Input

	ToColumnDefinitionOutput() ColumnDefinitionOutput
	ToColumnDefinitionOutputWithContext(context.Context) ColumnDefinitionOutput
}

type ColumnDefinitionArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ColumnDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnDefinition)(nil)).Elem()
}

func (i ColumnDefinitionArgs) ToColumnDefinitionOutput() ColumnDefinitionOutput {
	return i.ToColumnDefinitionOutputWithContext(context.Background())
}

func (i ColumnDefinitionArgs) ToColumnDefinitionOutputWithContext(ctx context.Context) ColumnDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnDefinitionOutput)
}





type ColumnDefinitionArrayInput interface {
	pulumi.Input

	ToColumnDefinitionArrayOutput() ColumnDefinitionArrayOutput
	ToColumnDefinitionArrayOutputWithContext(context.Context) ColumnDefinitionArrayOutput
}

type ColumnDefinitionArray []ColumnDefinitionInput

func (ColumnDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ColumnDefinition)(nil)).Elem()
}

func (i ColumnDefinitionArray) ToColumnDefinitionArrayOutput() ColumnDefinitionArrayOutput {
	return i.ToColumnDefinitionArrayOutputWithContext(context.Background())
}

func (i ColumnDefinitionArray) ToColumnDefinitionArrayOutputWithContext(ctx context.Context) ColumnDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnDefinitionArrayOutput)
}

type ColumnDefinitionOutput struct{ *pulumi.OutputState }

func (ColumnDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnDefinition)(nil)).Elem()
}

func (o ColumnDefinitionOutput) ToColumnDefinitionOutput() ColumnDefinitionOutput {
	return o
}

func (o ColumnDefinitionOutput) ToColumnDefinitionOutputWithContext(ctx context.Context) ColumnDefinitionOutput {
	return o
}

func (o ColumnDefinitionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnDefinition) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnDefinitionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnDefinition) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnDefinitionArrayOutput struct{ *pulumi.OutputState }

func (ColumnDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ColumnDefinition)(nil)).Elem()
}

func (o ColumnDefinitionArrayOutput) ToColumnDefinitionArrayOutput() ColumnDefinitionArrayOutput {
	return o
}

func (o ColumnDefinitionArrayOutput) ToColumnDefinitionArrayOutputWithContext(ctx context.Context) ColumnDefinitionArrayOutput {
	return o
}

func (o ColumnDefinitionArrayOutput) Index(i pulumi.IntInput) ColumnDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ColumnDefinition {
		return vs[0].([]ColumnDefinition)[vs[1].(int)]
	}).(ColumnDefinitionOutput)
}

type ColumnDefinitionResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ColumnDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ColumnDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnDefinitionResponse)(nil)).Elem()
}

func (o ColumnDefinitionResponseOutput) ToColumnDefinitionResponseOutput() ColumnDefinitionResponseOutput {
	return o
}

func (o ColumnDefinitionResponseOutput) ToColumnDefinitionResponseOutputWithContext(ctx context.Context) ColumnDefinitionResponseOutput {
	return o
}

func (o ColumnDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnDefinitionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnDefinitionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ColumnDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ColumnDefinitionResponse)(nil)).Elem()
}

func (o ColumnDefinitionResponseArrayOutput) ToColumnDefinitionResponseArrayOutput() ColumnDefinitionResponseArrayOutput {
	return o
}

func (o ColumnDefinitionResponseArrayOutput) ToColumnDefinitionResponseArrayOutputWithContext(ctx context.Context) ColumnDefinitionResponseArrayOutput {
	return o
}

func (o ColumnDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ColumnDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ColumnDefinitionResponse {
		return vs[0].([]ColumnDefinitionResponse)[vs[1].(int)]
	}).(ColumnDefinitionResponseOutput)
}

type DataCollectionEndpointNetworkAcls struct {
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type DataCollectionEndpointNetworkAclsInput interface {
	pulumi.Input

	ToDataCollectionEndpointNetworkAclsOutput() DataCollectionEndpointNetworkAclsOutput
	ToDataCollectionEndpointNetworkAclsOutputWithContext(context.Context) DataCollectionEndpointNetworkAclsOutput
}

type DataCollectionEndpointNetworkAclsArgs struct {
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (DataCollectionEndpointNetworkAclsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointNetworkAcls)(nil)).Elem()
}

func (i DataCollectionEndpointNetworkAclsArgs) ToDataCollectionEndpointNetworkAclsOutput() DataCollectionEndpointNetworkAclsOutput {
	return i.ToDataCollectionEndpointNetworkAclsOutputWithContext(context.Background())
}

func (i DataCollectionEndpointNetworkAclsArgs) ToDataCollectionEndpointNetworkAclsOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointNetworkAclsOutput)
}

func (i DataCollectionEndpointNetworkAclsArgs) ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput {
	return i.ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(context.Background())
}

func (i DataCollectionEndpointNetworkAclsArgs) ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointNetworkAclsOutput).ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx)
}









type DataCollectionEndpointNetworkAclsPtrInput interface {
	pulumi.Input

	ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput
	ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(context.Context) DataCollectionEndpointNetworkAclsPtrOutput
}

type dataCollectionEndpointNetworkAclsPtrType DataCollectionEndpointNetworkAclsArgs

func DataCollectionEndpointNetworkAclsPtr(v *DataCollectionEndpointNetworkAclsArgs) DataCollectionEndpointNetworkAclsPtrInput {
	return (*dataCollectionEndpointNetworkAclsPtrType)(v)
}

func (*dataCollectionEndpointNetworkAclsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointNetworkAcls)(nil)).Elem()
}

func (i *dataCollectionEndpointNetworkAclsPtrType) ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput {
	return i.ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(context.Background())
}

func (i *dataCollectionEndpointNetworkAclsPtrType) ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointNetworkAclsPtrOutput)
}

type DataCollectionEndpointNetworkAclsOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointNetworkAclsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointNetworkAcls)(nil)).Elem()
}

func (o DataCollectionEndpointNetworkAclsOutput) ToDataCollectionEndpointNetworkAclsOutput() DataCollectionEndpointNetworkAclsOutput {
	return o
}

func (o DataCollectionEndpointNetworkAclsOutput) ToDataCollectionEndpointNetworkAclsOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsOutput {
	return o
}

func (o DataCollectionEndpointNetworkAclsOutput) ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput {
	return o.ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(context.Background())
}

func (o DataCollectionEndpointNetworkAclsOutput) ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionEndpointNetworkAcls) *DataCollectionEndpointNetworkAcls {
		return &v
	}).(DataCollectionEndpointNetworkAclsPtrOutput)
}

func (o DataCollectionEndpointNetworkAclsOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointNetworkAcls) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointNetworkAclsPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointNetworkAclsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointNetworkAcls)(nil)).Elem()
}

func (o DataCollectionEndpointNetworkAclsPtrOutput) ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput {
	return o
}

func (o DataCollectionEndpointNetworkAclsPtrOutput) ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsPtrOutput {
	return o
}

func (o DataCollectionEndpointNetworkAclsPtrOutput) Elem() DataCollectionEndpointNetworkAclsOutput {
	return o.ApplyT(func(v *DataCollectionEndpointNetworkAcls) DataCollectionEndpointNetworkAcls {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointNetworkAcls
		return ret
	}).(DataCollectionEndpointNetworkAclsOutput)
}

func (o DataCollectionEndpointNetworkAclsPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointNetworkAcls) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResourceResponseSystemData struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type DataCollectionEndpointResourceResponseSystemDataOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResourceResponseSystemDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResourceResponseSystemData)(nil)).Elem()
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) ToDataCollectionEndpointResourceResponseSystemDataOutput() DataCollectionEndpointResourceResponseSystemDataOutput {
	return o
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) ToDataCollectionEndpointResourceResponseSystemDataOutputWithContext(ctx context.Context) DataCollectionEndpointResourceResponseSystemDataOutput {
	return o
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseConfigurationAccess struct {
	Endpoint string `pulumi:"endpoint"`
}

type DataCollectionEndpointResponseConfigurationAccessOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseConfigurationAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseConfigurationAccess)(nil)).Elem()
}

func (o DataCollectionEndpointResponseConfigurationAccessOutput) ToDataCollectionEndpointResponseConfigurationAccessOutput() DataCollectionEndpointResponseConfigurationAccessOutput {
	return o
}

func (o DataCollectionEndpointResponseConfigurationAccessOutput) ToDataCollectionEndpointResponseConfigurationAccessOutputWithContext(ctx context.Context) DataCollectionEndpointResponseConfigurationAccessOutput {
	return o
}

func (o DataCollectionEndpointResponseConfigurationAccessOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DataCollectionEndpointResponseConfigurationAccess) string { return v.Endpoint }).(pulumi.StringOutput)
}

type DataCollectionEndpointResponseConfigurationAccessPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseConfigurationAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseConfigurationAccess)(nil)).Elem()
}

func (o DataCollectionEndpointResponseConfigurationAccessPtrOutput) ToDataCollectionEndpointResponseConfigurationAccessPtrOutput() DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseConfigurationAccessPtrOutput) ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseConfigurationAccessPtrOutput) Elem() DataCollectionEndpointResponseConfigurationAccessOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseConfigurationAccess) DataCollectionEndpointResponseConfigurationAccess {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointResponseConfigurationAccess
		return ret
	}).(DataCollectionEndpointResponseConfigurationAccessOutput)
}

func (o DataCollectionEndpointResponseConfigurationAccessPtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseConfigurationAccess) *string {
		if v == nil {
			return nil
		}
		return &v.Endpoint
	}).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseLogsIngestion struct {
	Endpoint string `pulumi:"endpoint"`
}

type DataCollectionEndpointResponseLogsIngestionOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseLogsIngestionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseLogsIngestion)(nil)).Elem()
}

func (o DataCollectionEndpointResponseLogsIngestionOutput) ToDataCollectionEndpointResponseLogsIngestionOutput() DataCollectionEndpointResponseLogsIngestionOutput {
	return o
}

func (o DataCollectionEndpointResponseLogsIngestionOutput) ToDataCollectionEndpointResponseLogsIngestionOutputWithContext(ctx context.Context) DataCollectionEndpointResponseLogsIngestionOutput {
	return o
}

func (o DataCollectionEndpointResponseLogsIngestionOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DataCollectionEndpointResponseLogsIngestion) string { return v.Endpoint }).(pulumi.StringOutput)
}

type DataCollectionEndpointResponseLogsIngestionPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseLogsIngestionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseLogsIngestion)(nil)).Elem()
}

func (o DataCollectionEndpointResponseLogsIngestionPtrOutput) ToDataCollectionEndpointResponseLogsIngestionPtrOutput() DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseLogsIngestionPtrOutput) ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseLogsIngestionPtrOutput) Elem() DataCollectionEndpointResponseLogsIngestionOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseLogsIngestion) DataCollectionEndpointResponseLogsIngestion {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointResponseLogsIngestion
		return ret
	}).(DataCollectionEndpointResponseLogsIngestionOutput)
}

func (o DataCollectionEndpointResponseLogsIngestionPtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseLogsIngestion) *string {
		if v == nil {
			return nil
		}
		return &v.Endpoint
	}).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseNetworkAcls struct {
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}

type DataCollectionEndpointResponseNetworkAclsOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseNetworkAclsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseNetworkAcls)(nil)).Elem()
}

func (o DataCollectionEndpointResponseNetworkAclsOutput) ToDataCollectionEndpointResponseNetworkAclsOutput() DataCollectionEndpointResponseNetworkAclsOutput {
	return o
}

func (o DataCollectionEndpointResponseNetworkAclsOutput) ToDataCollectionEndpointResponseNetworkAclsOutputWithContext(ctx context.Context) DataCollectionEndpointResponseNetworkAclsOutput {
	return o
}

func (o DataCollectionEndpointResponseNetworkAclsOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResponseNetworkAcls) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseNetworkAclsPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseNetworkAclsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseNetworkAcls)(nil)).Elem()
}

func (o DataCollectionEndpointResponseNetworkAclsPtrOutput) ToDataCollectionEndpointResponseNetworkAclsPtrOutput() DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseNetworkAclsPtrOutput) ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseNetworkAclsPtrOutput) Elem() DataCollectionEndpointResponseNetworkAclsOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseNetworkAcls) DataCollectionEndpointResponseNetworkAcls {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointResponseNetworkAcls
		return ret
	}).(DataCollectionEndpointResponseNetworkAclsOutput)
}

func (o DataCollectionEndpointResponseNetworkAclsPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseNetworkAcls) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData)(nil)).Elem()
}

func (o DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) ToDataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput() DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput {
	return o
}

func (o DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) ToDataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutputWithContext(ctx context.Context) DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput {
	return o
}

func (o DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData) *string {
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData) *string {
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData) *string {
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData) *string {
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type DataCollectionRuleAssociationResponseMetadata struct {
	ProvisionedBy string `pulumi:"provisionedBy"`
}

type DataCollectionRuleAssociationResponseMetadataOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleAssociationResponseMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleAssociationResponseMetadata)(nil)).Elem()
}

func (o DataCollectionRuleAssociationResponseMetadataOutput) ToDataCollectionRuleAssociationResponseMetadataOutput() DataCollectionRuleAssociationResponseMetadataOutput {
	return o
}

func (o DataCollectionRuleAssociationResponseMetadataOutput) ToDataCollectionRuleAssociationResponseMetadataOutputWithContext(ctx context.Context) DataCollectionRuleAssociationResponseMetadataOutput {
	return o
}

func (o DataCollectionRuleAssociationResponseMetadataOutput) ProvisionedBy() pulumi.StringOutput {
	return o.ApplyT(func(v DataCollectionRuleAssociationResponseMetadata) string { return v.ProvisionedBy }).(pulumi.StringOutput)
}

type DataCollectionRuleDataSources struct {
	Extensions          []ExtensionDataSource       `pulumi:"extensions"`
	IisLogs             []IisLogsDataSource         `pulumi:"iisLogs"`
	LogFiles            []LogFilesDataSource        `pulumi:"logFiles"`
	PerformanceCounters []PerfCounterDataSource     `pulumi:"performanceCounters"`
	Syslog              []SyslogDataSource          `pulumi:"syslog"`
	WindowsEventLogs    []WindowsEventLogDataSource `pulumi:"windowsEventLogs"`
}





type DataCollectionRuleDataSourcesInput interface {
	pulumi.Input

	ToDataCollectionRuleDataSourcesOutput() DataCollectionRuleDataSourcesOutput
	ToDataCollectionRuleDataSourcesOutputWithContext(context.Context) DataCollectionRuleDataSourcesOutput
}

type DataCollectionRuleDataSourcesArgs struct {
	Extensions          ExtensionDataSourceArrayInput       `pulumi:"extensions"`
	IisLogs             IisLogsDataSourceArrayInput         `pulumi:"iisLogs"`
	LogFiles            LogFilesDataSourceArrayInput        `pulumi:"logFiles"`
	PerformanceCounters PerfCounterDataSourceArrayInput     `pulumi:"performanceCounters"`
	Syslog              SyslogDataSourceArrayInput          `pulumi:"syslog"`
	WindowsEventLogs    WindowsEventLogDataSourceArrayInput `pulumi:"windowsEventLogs"`
}

func (DataCollectionRuleDataSourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleDataSources)(nil)).Elem()
}

func (i DataCollectionRuleDataSourcesArgs) ToDataCollectionRuleDataSourcesOutput() DataCollectionRuleDataSourcesOutput {
	return i.ToDataCollectionRuleDataSourcesOutputWithContext(context.Background())
}

func (i DataCollectionRuleDataSourcesArgs) ToDataCollectionRuleDataSourcesOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDataSourcesOutput)
}

func (i DataCollectionRuleDataSourcesArgs) ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput {
	return i.ToDataCollectionRuleDataSourcesPtrOutputWithContext(context.Background())
}

func (i DataCollectionRuleDataSourcesArgs) ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDataSourcesOutput).ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx)
}









type DataCollectionRuleDataSourcesPtrInput interface {
	pulumi.Input

	ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput
	ToDataCollectionRuleDataSourcesPtrOutputWithContext(context.Context) DataCollectionRuleDataSourcesPtrOutput
}

type dataCollectionRuleDataSourcesPtrType DataCollectionRuleDataSourcesArgs

func DataCollectionRuleDataSourcesPtr(v *DataCollectionRuleDataSourcesArgs) DataCollectionRuleDataSourcesPtrInput {
	return (*dataCollectionRuleDataSourcesPtrType)(v)
}

func (*dataCollectionRuleDataSourcesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleDataSources)(nil)).Elem()
}

func (i *dataCollectionRuleDataSourcesPtrType) ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput {
	return i.ToDataCollectionRuleDataSourcesPtrOutputWithContext(context.Background())
}

func (i *dataCollectionRuleDataSourcesPtrType) ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDataSourcesPtrOutput)
}

type DataCollectionRuleDataSourcesOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleDataSourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleDataSources)(nil)).Elem()
}

func (o DataCollectionRuleDataSourcesOutput) ToDataCollectionRuleDataSourcesOutput() DataCollectionRuleDataSourcesOutput {
	return o
}

func (o DataCollectionRuleDataSourcesOutput) ToDataCollectionRuleDataSourcesOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesOutput {
	return o
}

func (o DataCollectionRuleDataSourcesOutput) ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput {
	return o.ToDataCollectionRuleDataSourcesPtrOutputWithContext(context.Background())
}

func (o DataCollectionRuleDataSourcesOutput) ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionRuleDataSources) *DataCollectionRuleDataSources {
		return &v
	}).(DataCollectionRuleDataSourcesPtrOutput)
}

func (o DataCollectionRuleDataSourcesOutput) Extensions() ExtensionDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []ExtensionDataSource { return v.Extensions }).(ExtensionDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesOutput) IisLogs() IisLogsDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []IisLogsDataSource { return v.IisLogs }).(IisLogsDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesOutput) LogFiles() LogFilesDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []LogFilesDataSource { return v.LogFiles }).(LogFilesDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesOutput) PerformanceCounters() PerfCounterDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []PerfCounterDataSource { return v.PerformanceCounters }).(PerfCounterDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesOutput) Syslog() SyslogDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []SyslogDataSource { return v.Syslog }).(SyslogDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesOutput) WindowsEventLogs() WindowsEventLogDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []WindowsEventLogDataSource { return v.WindowsEventLogs }).(WindowsEventLogDataSourceArrayOutput)
}

type DataCollectionRuleDataSourcesPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleDataSourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleDataSources)(nil)).Elem()
}

func (o DataCollectionRuleDataSourcesPtrOutput) ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput {
	return o
}

func (o DataCollectionRuleDataSourcesPtrOutput) ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesPtrOutput {
	return o
}

func (o DataCollectionRuleDataSourcesPtrOutput) Elem() DataCollectionRuleDataSourcesOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) DataCollectionRuleDataSources {
		if v != nil {
			return *v
		}
		var ret DataCollectionRuleDataSources
		return ret
	}).(DataCollectionRuleDataSourcesOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) Extensions() ExtensionDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []ExtensionDataSource {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(ExtensionDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) IisLogs() IisLogsDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []IisLogsDataSource {
		if v == nil {
			return nil
		}
		return v.IisLogs
	}).(IisLogsDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) LogFiles() LogFilesDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []LogFilesDataSource {
		if v == nil {
			return nil
		}
		return v.LogFiles
	}).(LogFilesDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) PerformanceCounters() PerfCounterDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []PerfCounterDataSource {
		if v == nil {
			return nil
		}
		return v.PerformanceCounters
	}).(PerfCounterDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) Syslog() SyslogDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []SyslogDataSource {
		if v == nil {
			return nil
		}
		return v.Syslog
	}).(SyslogDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) WindowsEventLogs() WindowsEventLogDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []WindowsEventLogDataSource {
		if v == nil {
			return nil
		}
		return v.WindowsEventLogs
	}).(WindowsEventLogDataSourceArrayOutput)
}

type DataCollectionRuleDestinations struct {
	AzureMonitorMetrics *DestinationsSpecAzureMonitorMetrics `pulumi:"azureMonitorMetrics"`
	LogAnalytics        []LogAnalyticsDestination            `pulumi:"logAnalytics"`
}





type DataCollectionRuleDestinationsInput interface {
	pulumi.Input

	ToDataCollectionRuleDestinationsOutput() DataCollectionRuleDestinationsOutput
	ToDataCollectionRuleDestinationsOutputWithContext(context.Context) DataCollectionRuleDestinationsOutput
}

type DataCollectionRuleDestinationsArgs struct {
	AzureMonitorMetrics DestinationsSpecAzureMonitorMetricsPtrInput `pulumi:"azureMonitorMetrics"`
	LogAnalytics        LogAnalyticsDestinationArrayInput           `pulumi:"logAnalytics"`
}

func (DataCollectionRuleDestinationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleDestinations)(nil)).Elem()
}

func (i DataCollectionRuleDestinationsArgs) ToDataCollectionRuleDestinationsOutput() DataCollectionRuleDestinationsOutput {
	return i.ToDataCollectionRuleDestinationsOutputWithContext(context.Background())
}

func (i DataCollectionRuleDestinationsArgs) ToDataCollectionRuleDestinationsOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDestinationsOutput)
}

func (i DataCollectionRuleDestinationsArgs) ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput {
	return i.ToDataCollectionRuleDestinationsPtrOutputWithContext(context.Background())
}

func (i DataCollectionRuleDestinationsArgs) ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDestinationsOutput).ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx)
}









type DataCollectionRuleDestinationsPtrInput interface {
	pulumi.Input

	ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput
	ToDataCollectionRuleDestinationsPtrOutputWithContext(context.Context) DataCollectionRuleDestinationsPtrOutput
}

type dataCollectionRuleDestinationsPtrType DataCollectionRuleDestinationsArgs

func DataCollectionRuleDestinationsPtr(v *DataCollectionRuleDestinationsArgs) DataCollectionRuleDestinationsPtrInput {
	return (*dataCollectionRuleDestinationsPtrType)(v)
}

func (*dataCollectionRuleDestinationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleDestinations)(nil)).Elem()
}

func (i *dataCollectionRuleDestinationsPtrType) ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput {
	return i.ToDataCollectionRuleDestinationsPtrOutputWithContext(context.Background())
}

func (i *dataCollectionRuleDestinationsPtrType) ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDestinationsPtrOutput)
}

type DataCollectionRuleDestinationsOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleDestinationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleDestinations)(nil)).Elem()
}

func (o DataCollectionRuleDestinationsOutput) ToDataCollectionRuleDestinationsOutput() DataCollectionRuleDestinationsOutput {
	return o
}

func (o DataCollectionRuleDestinationsOutput) ToDataCollectionRuleDestinationsOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsOutput {
	return o
}

func (o DataCollectionRuleDestinationsOutput) ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput {
	return o.ToDataCollectionRuleDestinationsPtrOutputWithContext(context.Background())
}

func (o DataCollectionRuleDestinationsOutput) ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionRuleDestinations) *DataCollectionRuleDestinations {
		return &v
	}).(DataCollectionRuleDestinationsPtrOutput)
}

func (o DataCollectionRuleDestinationsOutput) AzureMonitorMetrics() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleDestinations) *DestinationsSpecAzureMonitorMetrics {
		return v.AzureMonitorMetrics
	}).(DestinationsSpecAzureMonitorMetricsPtrOutput)
}

func (o DataCollectionRuleDestinationsOutput) LogAnalytics() LogAnalyticsDestinationArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDestinations) []LogAnalyticsDestination { return v.LogAnalytics }).(LogAnalyticsDestinationArrayOutput)
}

type DataCollectionRuleDestinationsPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleDestinationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleDestinations)(nil)).Elem()
}

func (o DataCollectionRuleDestinationsPtrOutput) ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput {
	return o
}

func (o DataCollectionRuleDestinationsPtrOutput) ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsPtrOutput {
	return o
}

func (o DataCollectionRuleDestinationsPtrOutput) Elem() DataCollectionRuleDestinationsOutput {
	return o.ApplyT(func(v *DataCollectionRuleDestinations) DataCollectionRuleDestinations {
		if v != nil {
			return *v
		}
		var ret DataCollectionRuleDestinations
		return ret
	}).(DataCollectionRuleDestinationsOutput)
}

func (o DataCollectionRuleDestinationsPtrOutput) AzureMonitorMetrics() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleDestinations) *DestinationsSpecAzureMonitorMetrics {
		if v == nil {
			return nil
		}
		return v.AzureMonitorMetrics
	}).(DestinationsSpecAzureMonitorMetricsPtrOutput)
}

func (o DataCollectionRuleDestinationsPtrOutput) LogAnalytics() LogAnalyticsDestinationArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDestinations) []LogAnalyticsDestination {
		if v == nil {
			return nil
		}
		return v.LogAnalytics
	}).(LogAnalyticsDestinationArrayOutput)
}

type DataCollectionRuleResourceResponseSystemData struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type DataCollectionRuleResourceResponseSystemDataOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResourceResponseSystemDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleResourceResponseSystemData)(nil)).Elem()
}

func (o DataCollectionRuleResourceResponseSystemDataOutput) ToDataCollectionRuleResourceResponseSystemDataOutput() DataCollectionRuleResourceResponseSystemDataOutput {
	return o
}

func (o DataCollectionRuleResourceResponseSystemDataOutput) ToDataCollectionRuleResourceResponseSystemDataOutputWithContext(ctx context.Context) DataCollectionRuleResourceResponseSystemDataOutput {
	return o
}

func (o DataCollectionRuleResourceResponseSystemDataOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleResourceResponseSystemData) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleResourceResponseSystemDataOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleResourceResponseSystemData) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleResourceResponseSystemDataOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleResourceResponseSystemData) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleResourceResponseSystemDataOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleResourceResponseSystemData) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleResourceResponseSystemDataOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleResourceResponseSystemData) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleResourceResponseSystemDataOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleResourceResponseSystemData) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type DataCollectionRuleResponseDataSources struct {
	Extensions          []ExtensionDataSourceResponse       `pulumi:"extensions"`
	IisLogs             []IisLogsDataSourceResponse         `pulumi:"iisLogs"`
	LogFiles            []LogFilesDataSourceResponse        `pulumi:"logFiles"`
	PerformanceCounters []PerfCounterDataSourceResponse     `pulumi:"performanceCounters"`
	Syslog              []SyslogDataSourceResponse          `pulumi:"syslog"`
	WindowsEventLogs    []WindowsEventLogDataSourceResponse `pulumi:"windowsEventLogs"`
}

type DataCollectionRuleResponseDataSourcesOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseDataSourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleResponseDataSources)(nil)).Elem()
}

func (o DataCollectionRuleResponseDataSourcesOutput) ToDataCollectionRuleResponseDataSourcesOutput() DataCollectionRuleResponseDataSourcesOutput {
	return o
}

func (o DataCollectionRuleResponseDataSourcesOutput) ToDataCollectionRuleResponseDataSourcesOutputWithContext(ctx context.Context) DataCollectionRuleResponseDataSourcesOutput {
	return o
}

func (o DataCollectionRuleResponseDataSourcesOutput) Extensions() ExtensionDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []ExtensionDataSourceResponse { return v.Extensions }).(ExtensionDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesOutput) IisLogs() IisLogsDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []IisLogsDataSourceResponse { return v.IisLogs }).(IisLogsDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesOutput) LogFiles() LogFilesDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []LogFilesDataSourceResponse { return v.LogFiles }).(LogFilesDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesOutput) PerformanceCounters() PerfCounterDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []PerfCounterDataSourceResponse {
		return v.PerformanceCounters
	}).(PerfCounterDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesOutput) Syslog() SyslogDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []SyslogDataSourceResponse { return v.Syslog }).(SyslogDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesOutput) WindowsEventLogs() WindowsEventLogDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []WindowsEventLogDataSourceResponse {
		return v.WindowsEventLogs
	}).(WindowsEventLogDataSourceResponseArrayOutput)
}

type DataCollectionRuleResponseDataSourcesPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseDataSourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleResponseDataSources)(nil)).Elem()
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) ToDataCollectionRuleResponseDataSourcesPtrOutput() DataCollectionRuleResponseDataSourcesPtrOutput {
	return o
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDataSourcesPtrOutput {
	return o
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) Elem() DataCollectionRuleResponseDataSourcesOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) DataCollectionRuleResponseDataSources {
		if v != nil {
			return *v
		}
		var ret DataCollectionRuleResponseDataSources
		return ret
	}).(DataCollectionRuleResponseDataSourcesOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) Extensions() ExtensionDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []ExtensionDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(ExtensionDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) IisLogs() IisLogsDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []IisLogsDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.IisLogs
	}).(IisLogsDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) LogFiles() LogFilesDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []LogFilesDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.LogFiles
	}).(LogFilesDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) PerformanceCounters() PerfCounterDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []PerfCounterDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.PerformanceCounters
	}).(PerfCounterDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) Syslog() SyslogDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []SyslogDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.Syslog
	}).(SyslogDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) WindowsEventLogs() WindowsEventLogDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []WindowsEventLogDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.WindowsEventLogs
	}).(WindowsEventLogDataSourceResponseArrayOutput)
}

type DataCollectionRuleResponseDestinations struct {
	AzureMonitorMetrics *DestinationsSpecResponseAzureMonitorMetrics `pulumi:"azureMonitorMetrics"`
	LogAnalytics        []LogAnalyticsDestinationResponse            `pulumi:"logAnalytics"`
}

type DataCollectionRuleResponseDestinationsOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseDestinationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleResponseDestinations)(nil)).Elem()
}

func (o DataCollectionRuleResponseDestinationsOutput) ToDataCollectionRuleResponseDestinationsOutput() DataCollectionRuleResponseDestinationsOutput {
	return o
}

func (o DataCollectionRuleResponseDestinationsOutput) ToDataCollectionRuleResponseDestinationsOutputWithContext(ctx context.Context) DataCollectionRuleResponseDestinationsOutput {
	return o
}

func (o DataCollectionRuleResponseDestinationsOutput) AzureMonitorMetrics() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDestinations) *DestinationsSpecResponseAzureMonitorMetrics {
		return v.AzureMonitorMetrics
	}).(DestinationsSpecResponseAzureMonitorMetricsPtrOutput)
}

func (o DataCollectionRuleResponseDestinationsOutput) LogAnalytics() LogAnalyticsDestinationResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDestinations) []LogAnalyticsDestinationResponse {
		return v.LogAnalytics
	}).(LogAnalyticsDestinationResponseArrayOutput)
}

type DataCollectionRuleResponseDestinationsPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseDestinationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleResponseDestinations)(nil)).Elem()
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) ToDataCollectionRuleResponseDestinationsPtrOutput() DataCollectionRuleResponseDestinationsPtrOutput {
	return o
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDestinationsPtrOutput {
	return o
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) Elem() DataCollectionRuleResponseDestinationsOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDestinations) DataCollectionRuleResponseDestinations {
		if v != nil {
			return *v
		}
		var ret DataCollectionRuleResponseDestinations
		return ret
	}).(DataCollectionRuleResponseDestinationsOutput)
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) AzureMonitorMetrics() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDestinations) *DestinationsSpecResponseAzureMonitorMetrics {
		if v == nil {
			return nil
		}
		return v.AzureMonitorMetrics
	}).(DestinationsSpecResponseAzureMonitorMetricsPtrOutput)
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) LogAnalytics() LogAnalyticsDestinationResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDestinations) []LogAnalyticsDestinationResponse {
		if v == nil {
			return nil
		}
		return v.LogAnalytics
	}).(LogAnalyticsDestinationResponseArrayOutput)
}

type DataCollectionRuleResponseMetadata struct {
	ProvisionedBy string `pulumi:"provisionedBy"`
}

type DataCollectionRuleResponseMetadataOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleResponseMetadata)(nil)).Elem()
}

func (o DataCollectionRuleResponseMetadataOutput) ToDataCollectionRuleResponseMetadataOutput() DataCollectionRuleResponseMetadataOutput {
	return o
}

func (o DataCollectionRuleResponseMetadataOutput) ToDataCollectionRuleResponseMetadataOutputWithContext(ctx context.Context) DataCollectionRuleResponseMetadataOutput {
	return o
}

func (o DataCollectionRuleResponseMetadataOutput) ProvisionedBy() pulumi.StringOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseMetadata) string { return v.ProvisionedBy }).(pulumi.StringOutput)
}

type DataFlow struct {
	Destinations []string `pulumi:"destinations"`
	OutputStream *string  `pulumi:"outputStream"`
	Streams      []string `pulumi:"streams"`
	TransformKql *string  `pulumi:"transformKql"`
}





type DataFlowInput interface {
	pulumi.Input

	ToDataFlowOutput() DataFlowOutput
	ToDataFlowOutputWithContext(context.Context) DataFlowOutput
}

type DataFlowArgs struct {
	Destinations pulumi.StringArrayInput `pulumi:"destinations"`
	OutputStream pulumi.StringPtrInput   `pulumi:"outputStream"`
	Streams      pulumi.StringArrayInput `pulumi:"streams"`
	TransformKql pulumi.StringPtrInput   `pulumi:"transformKql"`
}

func (DataFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlow)(nil)).Elem()
}

func (i DataFlowArgs) ToDataFlowOutput() DataFlowOutput {
	return i.ToDataFlowOutputWithContext(context.Background())
}

func (i DataFlowArgs) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowOutput)
}





type DataFlowArrayInput interface {
	pulumi.Input

	ToDataFlowArrayOutput() DataFlowArrayOutput
	ToDataFlowArrayOutputWithContext(context.Context) DataFlowArrayOutput
}

type DataFlowArray []DataFlowInput

func (DataFlowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataFlow)(nil)).Elem()
}

func (i DataFlowArray) ToDataFlowArrayOutput() DataFlowArrayOutput {
	return i.ToDataFlowArrayOutputWithContext(context.Background())
}

func (i DataFlowArray) ToDataFlowArrayOutputWithContext(ctx context.Context) DataFlowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowArrayOutput)
}

type DataFlowOutput struct{ *pulumi.OutputState }

func (DataFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlow)(nil)).Elem()
}

func (o DataFlowOutput) ToDataFlowOutput() DataFlowOutput {
	return o
}

func (o DataFlowOutput) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return o
}

func (o DataFlowOutput) Destinations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlow) []string { return v.Destinations }).(pulumi.StringArrayOutput)
}

func (o DataFlowOutput) OutputStream() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFlow) *string { return v.OutputStream }).(pulumi.StringPtrOutput)
}

func (o DataFlowOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlow) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

func (o DataFlowOutput) TransformKql() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFlow) *string { return v.TransformKql }).(pulumi.StringPtrOutput)
}

type DataFlowArrayOutput struct{ *pulumi.OutputState }

func (DataFlowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataFlow)(nil)).Elem()
}

func (o DataFlowArrayOutput) ToDataFlowArrayOutput() DataFlowArrayOutput {
	return o
}

func (o DataFlowArrayOutput) ToDataFlowArrayOutputWithContext(ctx context.Context) DataFlowArrayOutput {
	return o
}

func (o DataFlowArrayOutput) Index(i pulumi.IntInput) DataFlowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataFlow {
		return vs[0].([]DataFlow)[vs[1].(int)]
	}).(DataFlowOutput)
}

type DataFlowResponse struct {
	Destinations []string `pulumi:"destinations"`
	OutputStream *string  `pulumi:"outputStream"`
	Streams      []string `pulumi:"streams"`
	TransformKql *string  `pulumi:"transformKql"`
}

type DataFlowResponseOutput struct{ *pulumi.OutputState }

func (DataFlowResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlowResponse)(nil)).Elem()
}

func (o DataFlowResponseOutput) ToDataFlowResponseOutput() DataFlowResponseOutput {
	return o
}

func (o DataFlowResponseOutput) ToDataFlowResponseOutputWithContext(ctx context.Context) DataFlowResponseOutput {
	return o
}

func (o DataFlowResponseOutput) Destinations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlowResponse) []string { return v.Destinations }).(pulumi.StringArrayOutput)
}

func (o DataFlowResponseOutput) OutputStream() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFlowResponse) *string { return v.OutputStream }).(pulumi.StringPtrOutput)
}

func (o DataFlowResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlowResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

func (o DataFlowResponseOutput) TransformKql() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataFlowResponse) *string { return v.TransformKql }).(pulumi.StringPtrOutput)
}

type DataFlowResponseArrayOutput struct{ *pulumi.OutputState }

func (DataFlowResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataFlowResponse)(nil)).Elem()
}

func (o DataFlowResponseArrayOutput) ToDataFlowResponseArrayOutput() DataFlowResponseArrayOutput {
	return o
}

func (o DataFlowResponseArrayOutput) ToDataFlowResponseArrayOutputWithContext(ctx context.Context) DataFlowResponseArrayOutput {
	return o
}

func (o DataFlowResponseArrayOutput) Index(i pulumi.IntInput) DataFlowResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataFlowResponse {
		return vs[0].([]DataFlowResponse)[vs[1].(int)]
	}).(DataFlowResponseOutput)
}

type DestinationsSpecAzureMonitorMetrics struct {
	Name *string `pulumi:"name"`
}





type DestinationsSpecAzureMonitorMetricsInput interface {
	pulumi.Input

	ToDestinationsSpecAzureMonitorMetricsOutput() DestinationsSpecAzureMonitorMetricsOutput
	ToDestinationsSpecAzureMonitorMetricsOutputWithContext(context.Context) DestinationsSpecAzureMonitorMetricsOutput
}

type DestinationsSpecAzureMonitorMetricsArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DestinationsSpecAzureMonitorMetricsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationsSpecAzureMonitorMetrics)(nil)).Elem()
}

func (i DestinationsSpecAzureMonitorMetricsArgs) ToDestinationsSpecAzureMonitorMetricsOutput() DestinationsSpecAzureMonitorMetricsOutput {
	return i.ToDestinationsSpecAzureMonitorMetricsOutputWithContext(context.Background())
}

func (i DestinationsSpecAzureMonitorMetricsArgs) ToDestinationsSpecAzureMonitorMetricsOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecAzureMonitorMetricsOutput)
}

func (i DestinationsSpecAzureMonitorMetricsArgs) ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return i.ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (i DestinationsSpecAzureMonitorMetricsArgs) ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecAzureMonitorMetricsOutput).ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx)
}









type DestinationsSpecAzureMonitorMetricsPtrInput interface {
	pulumi.Input

	ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput
	ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput
}

type destinationsSpecAzureMonitorMetricsPtrType DestinationsSpecAzureMonitorMetricsArgs

func DestinationsSpecAzureMonitorMetricsPtr(v *DestinationsSpecAzureMonitorMetricsArgs) DestinationsSpecAzureMonitorMetricsPtrInput {
	return (*destinationsSpecAzureMonitorMetricsPtrType)(v)
}

func (*destinationsSpecAzureMonitorMetricsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationsSpecAzureMonitorMetrics)(nil)).Elem()
}

func (i *destinationsSpecAzureMonitorMetricsPtrType) ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return i.ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (i *destinationsSpecAzureMonitorMetricsPtrType) ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecAzureMonitorMetricsPtrOutput)
}

type DestinationsSpecAzureMonitorMetricsOutput struct{ *pulumi.OutputState }

func (DestinationsSpecAzureMonitorMetricsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationsSpecAzureMonitorMetrics)(nil)).Elem()
}

func (o DestinationsSpecAzureMonitorMetricsOutput) ToDestinationsSpecAzureMonitorMetricsOutput() DestinationsSpecAzureMonitorMetricsOutput {
	return o
}

func (o DestinationsSpecAzureMonitorMetricsOutput) ToDestinationsSpecAzureMonitorMetricsOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsOutput {
	return o
}

func (o DestinationsSpecAzureMonitorMetricsOutput) ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o.ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (o DestinationsSpecAzureMonitorMetricsOutput) ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DestinationsSpecAzureMonitorMetrics) *DestinationsSpecAzureMonitorMetrics {
		return &v
	}).(DestinationsSpecAzureMonitorMetricsPtrOutput)
}

func (o DestinationsSpecAzureMonitorMetricsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationsSpecAzureMonitorMetrics) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DestinationsSpecAzureMonitorMetricsPtrOutput struct{ *pulumi.OutputState }

func (DestinationsSpecAzureMonitorMetricsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationsSpecAzureMonitorMetrics)(nil)).Elem()
}

func (o DestinationsSpecAzureMonitorMetricsPtrOutput) ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o
}

func (o DestinationsSpecAzureMonitorMetricsPtrOutput) ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o
}

func (o DestinationsSpecAzureMonitorMetricsPtrOutput) Elem() DestinationsSpecAzureMonitorMetricsOutput {
	return o.ApplyT(func(v *DestinationsSpecAzureMonitorMetrics) DestinationsSpecAzureMonitorMetrics {
		if v != nil {
			return *v
		}
		var ret DestinationsSpecAzureMonitorMetrics
		return ret
	}).(DestinationsSpecAzureMonitorMetricsOutput)
}

func (o DestinationsSpecAzureMonitorMetricsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationsSpecAzureMonitorMetrics) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type DestinationsSpecResponseAzureMonitorMetrics struct {
	Name *string `pulumi:"name"`
}

type DestinationsSpecResponseAzureMonitorMetricsOutput struct{ *pulumi.OutputState }

func (DestinationsSpecResponseAzureMonitorMetricsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationsSpecResponseAzureMonitorMetrics)(nil)).Elem()
}

func (o DestinationsSpecResponseAzureMonitorMetricsOutput) ToDestinationsSpecResponseAzureMonitorMetricsOutput() DestinationsSpecResponseAzureMonitorMetricsOutput {
	return o
}

func (o DestinationsSpecResponseAzureMonitorMetricsOutput) ToDestinationsSpecResponseAzureMonitorMetricsOutputWithContext(ctx context.Context) DestinationsSpecResponseAzureMonitorMetricsOutput {
	return o
}

func (o DestinationsSpecResponseAzureMonitorMetricsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationsSpecResponseAzureMonitorMetrics) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DestinationsSpecResponseAzureMonitorMetricsPtrOutput struct{ *pulumi.OutputState }

func (DestinationsSpecResponseAzureMonitorMetricsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationsSpecResponseAzureMonitorMetrics)(nil)).Elem()
}

func (o DestinationsSpecResponseAzureMonitorMetricsPtrOutput) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutput() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o
}

func (o DestinationsSpecResponseAzureMonitorMetricsPtrOutput) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o
}

func (o DestinationsSpecResponseAzureMonitorMetricsPtrOutput) Elem() DestinationsSpecResponseAzureMonitorMetricsOutput {
	return o.ApplyT(func(v *DestinationsSpecResponseAzureMonitorMetrics) DestinationsSpecResponseAzureMonitorMetrics {
		if v != nil {
			return *v
		}
		var ret DestinationsSpecResponseAzureMonitorMetrics
		return ret
	}).(DestinationsSpecResponseAzureMonitorMetricsOutput)
}

func (o DestinationsSpecResponseAzureMonitorMetricsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationsSpecResponseAzureMonitorMetrics) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ExtensionDataSource struct {
	ExtensionName     string      `pulumi:"extensionName"`
	ExtensionSettings interface{} `pulumi:"extensionSettings"`
	InputDataSources  []string    `pulumi:"inputDataSources"`
	Name              *string     `pulumi:"name"`
	Streams           []string    `pulumi:"streams"`
}





type ExtensionDataSourceInput interface {
	pulumi.Input

	ToExtensionDataSourceOutput() ExtensionDataSourceOutput
	ToExtensionDataSourceOutputWithContext(context.Context) ExtensionDataSourceOutput
}

type ExtensionDataSourceArgs struct {
	ExtensionName     pulumi.StringInput      `pulumi:"extensionName"`
	ExtensionSettings pulumi.Input            `pulumi:"extensionSettings"`
	InputDataSources  pulumi.StringArrayInput `pulumi:"inputDataSources"`
	Name              pulumi.StringPtrInput   `pulumi:"name"`
	Streams           pulumi.StringArrayInput `pulumi:"streams"`
}

func (ExtensionDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionDataSource)(nil)).Elem()
}

func (i ExtensionDataSourceArgs) ToExtensionDataSourceOutput() ExtensionDataSourceOutput {
	return i.ToExtensionDataSourceOutputWithContext(context.Background())
}

func (i ExtensionDataSourceArgs) ToExtensionDataSourceOutputWithContext(ctx context.Context) ExtensionDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionDataSourceOutput)
}





type ExtensionDataSourceArrayInput interface {
	pulumi.Input

	ToExtensionDataSourceArrayOutput() ExtensionDataSourceArrayOutput
	ToExtensionDataSourceArrayOutputWithContext(context.Context) ExtensionDataSourceArrayOutput
}

type ExtensionDataSourceArray []ExtensionDataSourceInput

func (ExtensionDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionDataSource)(nil)).Elem()
}

func (i ExtensionDataSourceArray) ToExtensionDataSourceArrayOutput() ExtensionDataSourceArrayOutput {
	return i.ToExtensionDataSourceArrayOutputWithContext(context.Background())
}

func (i ExtensionDataSourceArray) ToExtensionDataSourceArrayOutputWithContext(ctx context.Context) ExtensionDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionDataSourceArrayOutput)
}

type ExtensionDataSourceOutput struct{ *pulumi.OutputState }

func (ExtensionDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionDataSource)(nil)).Elem()
}

func (o ExtensionDataSourceOutput) ToExtensionDataSourceOutput() ExtensionDataSourceOutput {
	return o
}

func (o ExtensionDataSourceOutput) ToExtensionDataSourceOutputWithContext(ctx context.Context) ExtensionDataSourceOutput {
	return o
}

func (o ExtensionDataSourceOutput) ExtensionName() pulumi.StringOutput {
	return o.ApplyT(func(v ExtensionDataSource) string { return v.ExtensionName }).(pulumi.StringOutput)
}

func (o ExtensionDataSourceOutput) ExtensionSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v ExtensionDataSource) interface{} { return v.ExtensionSettings }).(pulumi.AnyOutput)
}

func (o ExtensionDataSourceOutput) InputDataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExtensionDataSource) []string { return v.InputDataSources }).(pulumi.StringArrayOutput)
}

func (o ExtensionDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtensionDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExtensionDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type ExtensionDataSourceArrayOutput struct{ *pulumi.OutputState }

func (ExtensionDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionDataSource)(nil)).Elem()
}

func (o ExtensionDataSourceArrayOutput) ToExtensionDataSourceArrayOutput() ExtensionDataSourceArrayOutput {
	return o
}

func (o ExtensionDataSourceArrayOutput) ToExtensionDataSourceArrayOutputWithContext(ctx context.Context) ExtensionDataSourceArrayOutput {
	return o
}

func (o ExtensionDataSourceArrayOutput) Index(i pulumi.IntInput) ExtensionDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionDataSource {
		return vs[0].([]ExtensionDataSource)[vs[1].(int)]
	}).(ExtensionDataSourceOutput)
}

type ExtensionDataSourceResponse struct {
	ExtensionName     string      `pulumi:"extensionName"`
	ExtensionSettings interface{} `pulumi:"extensionSettings"`
	InputDataSources  []string    `pulumi:"inputDataSources"`
	Name              *string     `pulumi:"name"`
	Streams           []string    `pulumi:"streams"`
}

type ExtensionDataSourceResponseOutput struct{ *pulumi.OutputState }

func (ExtensionDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionDataSourceResponse)(nil)).Elem()
}

func (o ExtensionDataSourceResponseOutput) ToExtensionDataSourceResponseOutput() ExtensionDataSourceResponseOutput {
	return o
}

func (o ExtensionDataSourceResponseOutput) ToExtensionDataSourceResponseOutputWithContext(ctx context.Context) ExtensionDataSourceResponseOutput {
	return o
}

func (o ExtensionDataSourceResponseOutput) ExtensionName() pulumi.StringOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) string { return v.ExtensionName }).(pulumi.StringOutput)
}

func (o ExtensionDataSourceResponseOutput) ExtensionSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) interface{} { return v.ExtensionSettings }).(pulumi.AnyOutput)
}

func (o ExtensionDataSourceResponseOutput) InputDataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) []string { return v.InputDataSources }).(pulumi.StringArrayOutput)
}

func (o ExtensionDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtensionDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type ExtensionDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (ExtensionDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionDataSourceResponse)(nil)).Elem()
}

func (o ExtensionDataSourceResponseArrayOutput) ToExtensionDataSourceResponseArrayOutput() ExtensionDataSourceResponseArrayOutput {
	return o
}

func (o ExtensionDataSourceResponseArrayOutput) ToExtensionDataSourceResponseArrayOutputWithContext(ctx context.Context) ExtensionDataSourceResponseArrayOutput {
	return o
}

func (o ExtensionDataSourceResponseArrayOutput) Index(i pulumi.IntInput) ExtensionDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionDataSourceResponse {
		return vs[0].([]ExtensionDataSourceResponse)[vs[1].(int)]
	}).(ExtensionDataSourceResponseOutput)
}

type IisLogsDataSource struct {
	LogDirectories []string `pulumi:"logDirectories"`
	Name           *string  `pulumi:"name"`
	Streams        []string `pulumi:"streams"`
}





type IisLogsDataSourceInput interface {
	pulumi.Input

	ToIisLogsDataSourceOutput() IisLogsDataSourceOutput
	ToIisLogsDataSourceOutputWithContext(context.Context) IisLogsDataSourceOutput
}

type IisLogsDataSourceArgs struct {
	LogDirectories pulumi.StringArrayInput `pulumi:"logDirectories"`
	Name           pulumi.StringPtrInput   `pulumi:"name"`
	Streams        pulumi.StringArrayInput `pulumi:"streams"`
}

func (IisLogsDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IisLogsDataSource)(nil)).Elem()
}

func (i IisLogsDataSourceArgs) ToIisLogsDataSourceOutput() IisLogsDataSourceOutput {
	return i.ToIisLogsDataSourceOutputWithContext(context.Background())
}

func (i IisLogsDataSourceArgs) ToIisLogsDataSourceOutputWithContext(ctx context.Context) IisLogsDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IisLogsDataSourceOutput)
}





type IisLogsDataSourceArrayInput interface {
	pulumi.Input

	ToIisLogsDataSourceArrayOutput() IisLogsDataSourceArrayOutput
	ToIisLogsDataSourceArrayOutputWithContext(context.Context) IisLogsDataSourceArrayOutput
}

type IisLogsDataSourceArray []IisLogsDataSourceInput

func (IisLogsDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IisLogsDataSource)(nil)).Elem()
}

func (i IisLogsDataSourceArray) ToIisLogsDataSourceArrayOutput() IisLogsDataSourceArrayOutput {
	return i.ToIisLogsDataSourceArrayOutputWithContext(context.Background())
}

func (i IisLogsDataSourceArray) ToIisLogsDataSourceArrayOutputWithContext(ctx context.Context) IisLogsDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IisLogsDataSourceArrayOutput)
}

type IisLogsDataSourceOutput struct{ *pulumi.OutputState }

func (IisLogsDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IisLogsDataSource)(nil)).Elem()
}

func (o IisLogsDataSourceOutput) ToIisLogsDataSourceOutput() IisLogsDataSourceOutput {
	return o
}

func (o IisLogsDataSourceOutput) ToIisLogsDataSourceOutputWithContext(ctx context.Context) IisLogsDataSourceOutput {
	return o
}

func (o IisLogsDataSourceOutput) LogDirectories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IisLogsDataSource) []string { return v.LogDirectories }).(pulumi.StringArrayOutput)
}

func (o IisLogsDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IisLogsDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IisLogsDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IisLogsDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type IisLogsDataSourceArrayOutput struct{ *pulumi.OutputState }

func (IisLogsDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IisLogsDataSource)(nil)).Elem()
}

func (o IisLogsDataSourceArrayOutput) ToIisLogsDataSourceArrayOutput() IisLogsDataSourceArrayOutput {
	return o
}

func (o IisLogsDataSourceArrayOutput) ToIisLogsDataSourceArrayOutputWithContext(ctx context.Context) IisLogsDataSourceArrayOutput {
	return o
}

func (o IisLogsDataSourceArrayOutput) Index(i pulumi.IntInput) IisLogsDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IisLogsDataSource {
		return vs[0].([]IisLogsDataSource)[vs[1].(int)]
	}).(IisLogsDataSourceOutput)
}

type IisLogsDataSourceResponse struct {
	LogDirectories []string `pulumi:"logDirectories"`
	Name           *string  `pulumi:"name"`
	Streams        []string `pulumi:"streams"`
}

type IisLogsDataSourceResponseOutput struct{ *pulumi.OutputState }

func (IisLogsDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IisLogsDataSourceResponse)(nil)).Elem()
}

func (o IisLogsDataSourceResponseOutput) ToIisLogsDataSourceResponseOutput() IisLogsDataSourceResponseOutput {
	return o
}

func (o IisLogsDataSourceResponseOutput) ToIisLogsDataSourceResponseOutputWithContext(ctx context.Context) IisLogsDataSourceResponseOutput {
	return o
}

func (o IisLogsDataSourceResponseOutput) LogDirectories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IisLogsDataSourceResponse) []string { return v.LogDirectories }).(pulumi.StringArrayOutput)
}

func (o IisLogsDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IisLogsDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IisLogsDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IisLogsDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type IisLogsDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (IisLogsDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IisLogsDataSourceResponse)(nil)).Elem()
}

func (o IisLogsDataSourceResponseArrayOutput) ToIisLogsDataSourceResponseArrayOutput() IisLogsDataSourceResponseArrayOutput {
	return o
}

func (o IisLogsDataSourceResponseArrayOutput) ToIisLogsDataSourceResponseArrayOutputWithContext(ctx context.Context) IisLogsDataSourceResponseArrayOutput {
	return o
}

func (o IisLogsDataSourceResponseArrayOutput) Index(i pulumi.IntInput) IisLogsDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IisLogsDataSourceResponse {
		return vs[0].([]IisLogsDataSourceResponse)[vs[1].(int)]
	}).(IisLogsDataSourceResponseOutput)
}

type LogAnalyticsDestination struct {
	Name                *string `pulumi:"name"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}





type LogAnalyticsDestinationInput interface {
	pulumi.Input

	ToLogAnalyticsDestinationOutput() LogAnalyticsDestinationOutput
	ToLogAnalyticsDestinationOutputWithContext(context.Context) LogAnalyticsDestinationOutput
}

type LogAnalyticsDestinationArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (LogAnalyticsDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsDestination)(nil)).Elem()
}

func (i LogAnalyticsDestinationArgs) ToLogAnalyticsDestinationOutput() LogAnalyticsDestinationOutput {
	return i.ToLogAnalyticsDestinationOutputWithContext(context.Background())
}

func (i LogAnalyticsDestinationArgs) ToLogAnalyticsDestinationOutputWithContext(ctx context.Context) LogAnalyticsDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsDestinationOutput)
}





type LogAnalyticsDestinationArrayInput interface {
	pulumi.Input

	ToLogAnalyticsDestinationArrayOutput() LogAnalyticsDestinationArrayOutput
	ToLogAnalyticsDestinationArrayOutputWithContext(context.Context) LogAnalyticsDestinationArrayOutput
}

type LogAnalyticsDestinationArray []LogAnalyticsDestinationInput

func (LogAnalyticsDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogAnalyticsDestination)(nil)).Elem()
}

func (i LogAnalyticsDestinationArray) ToLogAnalyticsDestinationArrayOutput() LogAnalyticsDestinationArrayOutput {
	return i.ToLogAnalyticsDestinationArrayOutputWithContext(context.Background())
}

func (i LogAnalyticsDestinationArray) ToLogAnalyticsDestinationArrayOutputWithContext(ctx context.Context) LogAnalyticsDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsDestinationArrayOutput)
}

type LogAnalyticsDestinationOutput struct{ *pulumi.OutputState }

func (LogAnalyticsDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsDestination)(nil)).Elem()
}

func (o LogAnalyticsDestinationOutput) ToLogAnalyticsDestinationOutput() LogAnalyticsDestinationOutput {
	return o
}

func (o LogAnalyticsDestinationOutput) ToLogAnalyticsDestinationOutputWithContext(ctx context.Context) LogAnalyticsDestinationOutput {
	return o
}

func (o LogAnalyticsDestinationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsDestination) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsDestinationOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsDestination) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsDestinationArrayOutput struct{ *pulumi.OutputState }

func (LogAnalyticsDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogAnalyticsDestination)(nil)).Elem()
}

func (o LogAnalyticsDestinationArrayOutput) ToLogAnalyticsDestinationArrayOutput() LogAnalyticsDestinationArrayOutput {
	return o
}

func (o LogAnalyticsDestinationArrayOutput) ToLogAnalyticsDestinationArrayOutputWithContext(ctx context.Context) LogAnalyticsDestinationArrayOutput {
	return o
}

func (o LogAnalyticsDestinationArrayOutput) Index(i pulumi.IntInput) LogAnalyticsDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogAnalyticsDestination {
		return vs[0].([]LogAnalyticsDestination)[vs[1].(int)]
	}).(LogAnalyticsDestinationOutput)
}

type LogAnalyticsDestinationResponse struct {
	Name                *string `pulumi:"name"`
	WorkspaceId         string  `pulumi:"workspaceId"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}

type LogAnalyticsDestinationResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsDestinationResponse)(nil)).Elem()
}

func (o LogAnalyticsDestinationResponseOutput) ToLogAnalyticsDestinationResponseOutput() LogAnalyticsDestinationResponseOutput {
	return o
}

func (o LogAnalyticsDestinationResponseOutput) ToLogAnalyticsDestinationResponseOutputWithContext(ctx context.Context) LogAnalyticsDestinationResponseOutput {
	return o
}

func (o LogAnalyticsDestinationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsDestinationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsDestinationResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsDestinationResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o LogAnalyticsDestinationResponseOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsDestinationResponse) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsDestinationResponseArrayOutput struct{ *pulumi.OutputState }

func (LogAnalyticsDestinationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogAnalyticsDestinationResponse)(nil)).Elem()
}

func (o LogAnalyticsDestinationResponseArrayOutput) ToLogAnalyticsDestinationResponseArrayOutput() LogAnalyticsDestinationResponseArrayOutput {
	return o
}

func (o LogAnalyticsDestinationResponseArrayOutput) ToLogAnalyticsDestinationResponseArrayOutputWithContext(ctx context.Context) LogAnalyticsDestinationResponseArrayOutput {
	return o
}

func (o LogAnalyticsDestinationResponseArrayOutput) Index(i pulumi.IntInput) LogAnalyticsDestinationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogAnalyticsDestinationResponse {
		return vs[0].([]LogAnalyticsDestinationResponse)[vs[1].(int)]
	}).(LogAnalyticsDestinationResponseOutput)
}

type LogFileSettingsResponseText struct {
	RecordStartTimestampFormat string `pulumi:"recordStartTimestampFormat"`
}

type LogFileSettingsResponseTextOutput struct{ *pulumi.OutputState }

func (LogFileSettingsResponseTextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFileSettingsResponseText)(nil)).Elem()
}

func (o LogFileSettingsResponseTextOutput) ToLogFileSettingsResponseTextOutput() LogFileSettingsResponseTextOutput {
	return o
}

func (o LogFileSettingsResponseTextOutput) ToLogFileSettingsResponseTextOutputWithContext(ctx context.Context) LogFileSettingsResponseTextOutput {
	return o
}

func (o LogFileSettingsResponseTextOutput) RecordStartTimestampFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LogFileSettingsResponseText) string { return v.RecordStartTimestampFormat }).(pulumi.StringOutput)
}

type LogFileSettingsResponseTextPtrOutput struct{ *pulumi.OutputState }

func (LogFileSettingsResponseTextPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFileSettingsResponseText)(nil)).Elem()
}

func (o LogFileSettingsResponseTextPtrOutput) ToLogFileSettingsResponseTextPtrOutput() LogFileSettingsResponseTextPtrOutput {
	return o
}

func (o LogFileSettingsResponseTextPtrOutput) ToLogFileSettingsResponseTextPtrOutputWithContext(ctx context.Context) LogFileSettingsResponseTextPtrOutput {
	return o
}

func (o LogFileSettingsResponseTextPtrOutput) Elem() LogFileSettingsResponseTextOutput {
	return o.ApplyT(func(v *LogFileSettingsResponseText) LogFileSettingsResponseText {
		if v != nil {
			return *v
		}
		var ret LogFileSettingsResponseText
		return ret
	}).(LogFileSettingsResponseTextOutput)
}

func (o LogFileSettingsResponseTextPtrOutput) RecordStartTimestampFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFileSettingsResponseText) *string {
		if v == nil {
			return nil
		}
		return &v.RecordStartTimestampFormat
	}).(pulumi.StringPtrOutput)
}

type LogFileSettingsText struct {
	RecordStartTimestampFormat string `pulumi:"recordStartTimestampFormat"`
}





type LogFileSettingsTextInput interface {
	pulumi.Input

	ToLogFileSettingsTextOutput() LogFileSettingsTextOutput
	ToLogFileSettingsTextOutputWithContext(context.Context) LogFileSettingsTextOutput
}

type LogFileSettingsTextArgs struct {
	RecordStartTimestampFormat pulumi.StringInput `pulumi:"recordStartTimestampFormat"`
}

func (LogFileSettingsTextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFileSettingsText)(nil)).Elem()
}

func (i LogFileSettingsTextArgs) ToLogFileSettingsTextOutput() LogFileSettingsTextOutput {
	return i.ToLogFileSettingsTextOutputWithContext(context.Background())
}

func (i LogFileSettingsTextArgs) ToLogFileSettingsTextOutputWithContext(ctx context.Context) LogFileSettingsTextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFileSettingsTextOutput)
}

func (i LogFileSettingsTextArgs) ToLogFileSettingsTextPtrOutput() LogFileSettingsTextPtrOutput {
	return i.ToLogFileSettingsTextPtrOutputWithContext(context.Background())
}

func (i LogFileSettingsTextArgs) ToLogFileSettingsTextPtrOutputWithContext(ctx context.Context) LogFileSettingsTextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFileSettingsTextOutput).ToLogFileSettingsTextPtrOutputWithContext(ctx)
}









type LogFileSettingsTextPtrInput interface {
	pulumi.Input

	ToLogFileSettingsTextPtrOutput() LogFileSettingsTextPtrOutput
	ToLogFileSettingsTextPtrOutputWithContext(context.Context) LogFileSettingsTextPtrOutput
}

type logFileSettingsTextPtrType LogFileSettingsTextArgs

func LogFileSettingsTextPtr(v *LogFileSettingsTextArgs) LogFileSettingsTextPtrInput {
	return (*logFileSettingsTextPtrType)(v)
}

func (*logFileSettingsTextPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFileSettingsText)(nil)).Elem()
}

func (i *logFileSettingsTextPtrType) ToLogFileSettingsTextPtrOutput() LogFileSettingsTextPtrOutput {
	return i.ToLogFileSettingsTextPtrOutputWithContext(context.Background())
}

func (i *logFileSettingsTextPtrType) ToLogFileSettingsTextPtrOutputWithContext(ctx context.Context) LogFileSettingsTextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFileSettingsTextPtrOutput)
}

type LogFileSettingsTextOutput struct{ *pulumi.OutputState }

func (LogFileSettingsTextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFileSettingsText)(nil)).Elem()
}

func (o LogFileSettingsTextOutput) ToLogFileSettingsTextOutput() LogFileSettingsTextOutput {
	return o
}

func (o LogFileSettingsTextOutput) ToLogFileSettingsTextOutputWithContext(ctx context.Context) LogFileSettingsTextOutput {
	return o
}

func (o LogFileSettingsTextOutput) ToLogFileSettingsTextPtrOutput() LogFileSettingsTextPtrOutput {
	return o.ToLogFileSettingsTextPtrOutputWithContext(context.Background())
}

func (o LogFileSettingsTextOutput) ToLogFileSettingsTextPtrOutputWithContext(ctx context.Context) LogFileSettingsTextPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogFileSettingsText) *LogFileSettingsText {
		return &v
	}).(LogFileSettingsTextPtrOutput)
}

func (o LogFileSettingsTextOutput) RecordStartTimestampFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LogFileSettingsText) string { return v.RecordStartTimestampFormat }).(pulumi.StringOutput)
}

type LogFileSettingsTextPtrOutput struct{ *pulumi.OutputState }

func (LogFileSettingsTextPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFileSettingsText)(nil)).Elem()
}

func (o LogFileSettingsTextPtrOutput) ToLogFileSettingsTextPtrOutput() LogFileSettingsTextPtrOutput {
	return o
}

func (o LogFileSettingsTextPtrOutput) ToLogFileSettingsTextPtrOutputWithContext(ctx context.Context) LogFileSettingsTextPtrOutput {
	return o
}

func (o LogFileSettingsTextPtrOutput) Elem() LogFileSettingsTextOutput {
	return o.ApplyT(func(v *LogFileSettingsText) LogFileSettingsText {
		if v != nil {
			return *v
		}
		var ret LogFileSettingsText
		return ret
	}).(LogFileSettingsTextOutput)
}

func (o LogFileSettingsTextPtrOutput) RecordStartTimestampFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogFileSettingsText) *string {
		if v == nil {
			return nil
		}
		return &v.RecordStartTimestampFormat
	}).(pulumi.StringPtrOutput)
}

type LogFilesDataSource struct {
	FilePatterns []string                    `pulumi:"filePatterns"`
	Format       string                      `pulumi:"format"`
	Name         *string                     `pulumi:"name"`
	Settings     *LogFilesDataSourceSettings `pulumi:"settings"`
	Streams      []string                    `pulumi:"streams"`
}





type LogFilesDataSourceInput interface {
	pulumi.Input

	ToLogFilesDataSourceOutput() LogFilesDataSourceOutput
	ToLogFilesDataSourceOutputWithContext(context.Context) LogFilesDataSourceOutput
}

type LogFilesDataSourceArgs struct {
	FilePatterns pulumi.StringArrayInput            `pulumi:"filePatterns"`
	Format       pulumi.StringInput                 `pulumi:"format"`
	Name         pulumi.StringPtrInput              `pulumi:"name"`
	Settings     LogFilesDataSourceSettingsPtrInput `pulumi:"settings"`
	Streams      pulumi.StringArrayInput            `pulumi:"streams"`
}

func (LogFilesDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFilesDataSource)(nil)).Elem()
}

func (i LogFilesDataSourceArgs) ToLogFilesDataSourceOutput() LogFilesDataSourceOutput {
	return i.ToLogFilesDataSourceOutputWithContext(context.Background())
}

func (i LogFilesDataSourceArgs) ToLogFilesDataSourceOutputWithContext(ctx context.Context) LogFilesDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFilesDataSourceOutput)
}





type LogFilesDataSourceArrayInput interface {
	pulumi.Input

	ToLogFilesDataSourceArrayOutput() LogFilesDataSourceArrayOutput
	ToLogFilesDataSourceArrayOutputWithContext(context.Context) LogFilesDataSourceArrayOutput
}

type LogFilesDataSourceArray []LogFilesDataSourceInput

func (LogFilesDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogFilesDataSource)(nil)).Elem()
}

func (i LogFilesDataSourceArray) ToLogFilesDataSourceArrayOutput() LogFilesDataSourceArrayOutput {
	return i.ToLogFilesDataSourceArrayOutputWithContext(context.Background())
}

func (i LogFilesDataSourceArray) ToLogFilesDataSourceArrayOutputWithContext(ctx context.Context) LogFilesDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFilesDataSourceArrayOutput)
}

type LogFilesDataSourceOutput struct{ *pulumi.OutputState }

func (LogFilesDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFilesDataSource)(nil)).Elem()
}

func (o LogFilesDataSourceOutput) ToLogFilesDataSourceOutput() LogFilesDataSourceOutput {
	return o
}

func (o LogFilesDataSourceOutput) ToLogFilesDataSourceOutputWithContext(ctx context.Context) LogFilesDataSourceOutput {
	return o
}

func (o LogFilesDataSourceOutput) FilePatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogFilesDataSource) []string { return v.FilePatterns }).(pulumi.StringArrayOutput)
}

func (o LogFilesDataSourceOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v LogFilesDataSource) string { return v.Format }).(pulumi.StringOutput)
}

func (o LogFilesDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogFilesDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LogFilesDataSourceOutput) Settings() LogFilesDataSourceSettingsPtrOutput {
	return o.ApplyT(func(v LogFilesDataSource) *LogFilesDataSourceSettings { return v.Settings }).(LogFilesDataSourceSettingsPtrOutput)
}

func (o LogFilesDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogFilesDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type LogFilesDataSourceArrayOutput struct{ *pulumi.OutputState }

func (LogFilesDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogFilesDataSource)(nil)).Elem()
}

func (o LogFilesDataSourceArrayOutput) ToLogFilesDataSourceArrayOutput() LogFilesDataSourceArrayOutput {
	return o
}

func (o LogFilesDataSourceArrayOutput) ToLogFilesDataSourceArrayOutputWithContext(ctx context.Context) LogFilesDataSourceArrayOutput {
	return o
}

func (o LogFilesDataSourceArrayOutput) Index(i pulumi.IntInput) LogFilesDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogFilesDataSource {
		return vs[0].([]LogFilesDataSource)[vs[1].(int)]
	}).(LogFilesDataSourceOutput)
}

type LogFilesDataSourceResponse struct {
	FilePatterns []string                            `pulumi:"filePatterns"`
	Format       string                              `pulumi:"format"`
	Name         *string                             `pulumi:"name"`
	Settings     *LogFilesDataSourceResponseSettings `pulumi:"settings"`
	Streams      []string                            `pulumi:"streams"`
}

type LogFilesDataSourceResponseOutput struct{ *pulumi.OutputState }

func (LogFilesDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFilesDataSourceResponse)(nil)).Elem()
}

func (o LogFilesDataSourceResponseOutput) ToLogFilesDataSourceResponseOutput() LogFilesDataSourceResponseOutput {
	return o
}

func (o LogFilesDataSourceResponseOutput) ToLogFilesDataSourceResponseOutputWithContext(ctx context.Context) LogFilesDataSourceResponseOutput {
	return o
}

func (o LogFilesDataSourceResponseOutput) FilePatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogFilesDataSourceResponse) []string { return v.FilePatterns }).(pulumi.StringArrayOutput)
}

func (o LogFilesDataSourceResponseOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v LogFilesDataSourceResponse) string { return v.Format }).(pulumi.StringOutput)
}

func (o LogFilesDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogFilesDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LogFilesDataSourceResponseOutput) Settings() LogFilesDataSourceResponseSettingsPtrOutput {
	return o.ApplyT(func(v LogFilesDataSourceResponse) *LogFilesDataSourceResponseSettings { return v.Settings }).(LogFilesDataSourceResponseSettingsPtrOutput)
}

func (o LogFilesDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogFilesDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type LogFilesDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (LogFilesDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogFilesDataSourceResponse)(nil)).Elem()
}

func (o LogFilesDataSourceResponseArrayOutput) ToLogFilesDataSourceResponseArrayOutput() LogFilesDataSourceResponseArrayOutput {
	return o
}

func (o LogFilesDataSourceResponseArrayOutput) ToLogFilesDataSourceResponseArrayOutputWithContext(ctx context.Context) LogFilesDataSourceResponseArrayOutput {
	return o
}

func (o LogFilesDataSourceResponseArrayOutput) Index(i pulumi.IntInput) LogFilesDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogFilesDataSourceResponse {
		return vs[0].([]LogFilesDataSourceResponse)[vs[1].(int)]
	}).(LogFilesDataSourceResponseOutput)
}

type LogFilesDataSourceResponseSettings struct {
	Text *LogFileSettingsResponseText `pulumi:"text"`
}

type LogFilesDataSourceResponseSettingsOutput struct{ *pulumi.OutputState }

func (LogFilesDataSourceResponseSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFilesDataSourceResponseSettings)(nil)).Elem()
}

func (o LogFilesDataSourceResponseSettingsOutput) ToLogFilesDataSourceResponseSettingsOutput() LogFilesDataSourceResponseSettingsOutput {
	return o
}

func (o LogFilesDataSourceResponseSettingsOutput) ToLogFilesDataSourceResponseSettingsOutputWithContext(ctx context.Context) LogFilesDataSourceResponseSettingsOutput {
	return o
}

func (o LogFilesDataSourceResponseSettingsOutput) Text() LogFileSettingsResponseTextPtrOutput {
	return o.ApplyT(func(v LogFilesDataSourceResponseSettings) *LogFileSettingsResponseText { return v.Text }).(LogFileSettingsResponseTextPtrOutput)
}

type LogFilesDataSourceResponseSettingsPtrOutput struct{ *pulumi.OutputState }

func (LogFilesDataSourceResponseSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFilesDataSourceResponseSettings)(nil)).Elem()
}

func (o LogFilesDataSourceResponseSettingsPtrOutput) ToLogFilesDataSourceResponseSettingsPtrOutput() LogFilesDataSourceResponseSettingsPtrOutput {
	return o
}

func (o LogFilesDataSourceResponseSettingsPtrOutput) ToLogFilesDataSourceResponseSettingsPtrOutputWithContext(ctx context.Context) LogFilesDataSourceResponseSettingsPtrOutput {
	return o
}

func (o LogFilesDataSourceResponseSettingsPtrOutput) Elem() LogFilesDataSourceResponseSettingsOutput {
	return o.ApplyT(func(v *LogFilesDataSourceResponseSettings) LogFilesDataSourceResponseSettings {
		if v != nil {
			return *v
		}
		var ret LogFilesDataSourceResponseSettings
		return ret
	}).(LogFilesDataSourceResponseSettingsOutput)
}

func (o LogFilesDataSourceResponseSettingsPtrOutput) Text() LogFileSettingsResponseTextPtrOutput {
	return o.ApplyT(func(v *LogFilesDataSourceResponseSettings) *LogFileSettingsResponseText {
		if v == nil {
			return nil
		}
		return v.Text
	}).(LogFileSettingsResponseTextPtrOutput)
}

type LogFilesDataSourceSettings struct {
	Text *LogFileSettingsText `pulumi:"text"`
}





type LogFilesDataSourceSettingsInput interface {
	pulumi.Input

	ToLogFilesDataSourceSettingsOutput() LogFilesDataSourceSettingsOutput
	ToLogFilesDataSourceSettingsOutputWithContext(context.Context) LogFilesDataSourceSettingsOutput
}

type LogFilesDataSourceSettingsArgs struct {
	Text LogFileSettingsTextPtrInput `pulumi:"text"`
}

func (LogFilesDataSourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFilesDataSourceSettings)(nil)).Elem()
}

func (i LogFilesDataSourceSettingsArgs) ToLogFilesDataSourceSettingsOutput() LogFilesDataSourceSettingsOutput {
	return i.ToLogFilesDataSourceSettingsOutputWithContext(context.Background())
}

func (i LogFilesDataSourceSettingsArgs) ToLogFilesDataSourceSettingsOutputWithContext(ctx context.Context) LogFilesDataSourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFilesDataSourceSettingsOutput)
}

func (i LogFilesDataSourceSettingsArgs) ToLogFilesDataSourceSettingsPtrOutput() LogFilesDataSourceSettingsPtrOutput {
	return i.ToLogFilesDataSourceSettingsPtrOutputWithContext(context.Background())
}

func (i LogFilesDataSourceSettingsArgs) ToLogFilesDataSourceSettingsPtrOutputWithContext(ctx context.Context) LogFilesDataSourceSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFilesDataSourceSettingsOutput).ToLogFilesDataSourceSettingsPtrOutputWithContext(ctx)
}









type LogFilesDataSourceSettingsPtrInput interface {
	pulumi.Input

	ToLogFilesDataSourceSettingsPtrOutput() LogFilesDataSourceSettingsPtrOutput
	ToLogFilesDataSourceSettingsPtrOutputWithContext(context.Context) LogFilesDataSourceSettingsPtrOutput
}

type logFilesDataSourceSettingsPtrType LogFilesDataSourceSettingsArgs

func LogFilesDataSourceSettingsPtr(v *LogFilesDataSourceSettingsArgs) LogFilesDataSourceSettingsPtrInput {
	return (*logFilesDataSourceSettingsPtrType)(v)
}

func (*logFilesDataSourceSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFilesDataSourceSettings)(nil)).Elem()
}

func (i *logFilesDataSourceSettingsPtrType) ToLogFilesDataSourceSettingsPtrOutput() LogFilesDataSourceSettingsPtrOutput {
	return i.ToLogFilesDataSourceSettingsPtrOutputWithContext(context.Background())
}

func (i *logFilesDataSourceSettingsPtrType) ToLogFilesDataSourceSettingsPtrOutputWithContext(ctx context.Context) LogFilesDataSourceSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogFilesDataSourceSettingsPtrOutput)
}

type LogFilesDataSourceSettingsOutput struct{ *pulumi.OutputState }

func (LogFilesDataSourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogFilesDataSourceSettings)(nil)).Elem()
}

func (o LogFilesDataSourceSettingsOutput) ToLogFilesDataSourceSettingsOutput() LogFilesDataSourceSettingsOutput {
	return o
}

func (o LogFilesDataSourceSettingsOutput) ToLogFilesDataSourceSettingsOutputWithContext(ctx context.Context) LogFilesDataSourceSettingsOutput {
	return o
}

func (o LogFilesDataSourceSettingsOutput) ToLogFilesDataSourceSettingsPtrOutput() LogFilesDataSourceSettingsPtrOutput {
	return o.ToLogFilesDataSourceSettingsPtrOutputWithContext(context.Background())
}

func (o LogFilesDataSourceSettingsOutput) ToLogFilesDataSourceSettingsPtrOutputWithContext(ctx context.Context) LogFilesDataSourceSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogFilesDataSourceSettings) *LogFilesDataSourceSettings {
		return &v
	}).(LogFilesDataSourceSettingsPtrOutput)
}

func (o LogFilesDataSourceSettingsOutput) Text() LogFileSettingsTextPtrOutput {
	return o.ApplyT(func(v LogFilesDataSourceSettings) *LogFileSettingsText { return v.Text }).(LogFileSettingsTextPtrOutput)
}

type LogFilesDataSourceSettingsPtrOutput struct{ *pulumi.OutputState }

func (LogFilesDataSourceSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogFilesDataSourceSettings)(nil)).Elem()
}

func (o LogFilesDataSourceSettingsPtrOutput) ToLogFilesDataSourceSettingsPtrOutput() LogFilesDataSourceSettingsPtrOutput {
	return o
}

func (o LogFilesDataSourceSettingsPtrOutput) ToLogFilesDataSourceSettingsPtrOutputWithContext(ctx context.Context) LogFilesDataSourceSettingsPtrOutput {
	return o
}

func (o LogFilesDataSourceSettingsPtrOutput) Elem() LogFilesDataSourceSettingsOutput {
	return o.ApplyT(func(v *LogFilesDataSourceSettings) LogFilesDataSourceSettings {
		if v != nil {
			return *v
		}
		var ret LogFilesDataSourceSettings
		return ret
	}).(LogFilesDataSourceSettingsOutput)
}

func (o LogFilesDataSourceSettingsPtrOutput) Text() LogFileSettingsTextPtrOutput {
	return o.ApplyT(func(v *LogFilesDataSourceSettings) *LogFileSettingsText {
		if v == nil {
			return nil
		}
		return v.Text
	}).(LogFileSettingsTextPtrOutput)
}

type PerfCounterDataSource struct {
	CounterSpecifiers          []string `pulumi:"counterSpecifiers"`
	Name                       *string  `pulumi:"name"`
	SamplingFrequencyInSeconds *int     `pulumi:"samplingFrequencyInSeconds"`
	Streams                    []string `pulumi:"streams"`
}





type PerfCounterDataSourceInput interface {
	pulumi.Input

	ToPerfCounterDataSourceOutput() PerfCounterDataSourceOutput
	ToPerfCounterDataSourceOutputWithContext(context.Context) PerfCounterDataSourceOutput
}

type PerfCounterDataSourceArgs struct {
	CounterSpecifiers          pulumi.StringArrayInput `pulumi:"counterSpecifiers"`
	Name                       pulumi.StringPtrInput   `pulumi:"name"`
	SamplingFrequencyInSeconds pulumi.IntPtrInput      `pulumi:"samplingFrequencyInSeconds"`
	Streams                    pulumi.StringArrayInput `pulumi:"streams"`
}

func (PerfCounterDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerfCounterDataSource)(nil)).Elem()
}

func (i PerfCounterDataSourceArgs) ToPerfCounterDataSourceOutput() PerfCounterDataSourceOutput {
	return i.ToPerfCounterDataSourceOutputWithContext(context.Background())
}

func (i PerfCounterDataSourceArgs) ToPerfCounterDataSourceOutputWithContext(ctx context.Context) PerfCounterDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerfCounterDataSourceOutput)
}





type PerfCounterDataSourceArrayInput interface {
	pulumi.Input

	ToPerfCounterDataSourceArrayOutput() PerfCounterDataSourceArrayOutput
	ToPerfCounterDataSourceArrayOutputWithContext(context.Context) PerfCounterDataSourceArrayOutput
}

type PerfCounterDataSourceArray []PerfCounterDataSourceInput

func (PerfCounterDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerfCounterDataSource)(nil)).Elem()
}

func (i PerfCounterDataSourceArray) ToPerfCounterDataSourceArrayOutput() PerfCounterDataSourceArrayOutput {
	return i.ToPerfCounterDataSourceArrayOutputWithContext(context.Background())
}

func (i PerfCounterDataSourceArray) ToPerfCounterDataSourceArrayOutputWithContext(ctx context.Context) PerfCounterDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerfCounterDataSourceArrayOutput)
}

type PerfCounterDataSourceOutput struct{ *pulumi.OutputState }

func (PerfCounterDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerfCounterDataSource)(nil)).Elem()
}

func (o PerfCounterDataSourceOutput) ToPerfCounterDataSourceOutput() PerfCounterDataSourceOutput {
	return o
}

func (o PerfCounterDataSourceOutput) ToPerfCounterDataSourceOutputWithContext(ctx context.Context) PerfCounterDataSourceOutput {
	return o
}

func (o PerfCounterDataSourceOutput) CounterSpecifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PerfCounterDataSource) []string { return v.CounterSpecifiers }).(pulumi.StringArrayOutput)
}

func (o PerfCounterDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PerfCounterDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PerfCounterDataSourceOutput) SamplingFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PerfCounterDataSource) *int { return v.SamplingFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o PerfCounterDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PerfCounterDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type PerfCounterDataSourceArrayOutput struct{ *pulumi.OutputState }

func (PerfCounterDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerfCounterDataSource)(nil)).Elem()
}

func (o PerfCounterDataSourceArrayOutput) ToPerfCounterDataSourceArrayOutput() PerfCounterDataSourceArrayOutput {
	return o
}

func (o PerfCounterDataSourceArrayOutput) ToPerfCounterDataSourceArrayOutputWithContext(ctx context.Context) PerfCounterDataSourceArrayOutput {
	return o
}

func (o PerfCounterDataSourceArrayOutput) Index(i pulumi.IntInput) PerfCounterDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerfCounterDataSource {
		return vs[0].([]PerfCounterDataSource)[vs[1].(int)]
	}).(PerfCounterDataSourceOutput)
}

type PerfCounterDataSourceResponse struct {
	CounterSpecifiers          []string `pulumi:"counterSpecifiers"`
	Name                       *string  `pulumi:"name"`
	SamplingFrequencyInSeconds *int     `pulumi:"samplingFrequencyInSeconds"`
	Streams                    []string `pulumi:"streams"`
}

type PerfCounterDataSourceResponseOutput struct{ *pulumi.OutputState }

func (PerfCounterDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerfCounterDataSourceResponse)(nil)).Elem()
}

func (o PerfCounterDataSourceResponseOutput) ToPerfCounterDataSourceResponseOutput() PerfCounterDataSourceResponseOutput {
	return o
}

func (o PerfCounterDataSourceResponseOutput) ToPerfCounterDataSourceResponseOutputWithContext(ctx context.Context) PerfCounterDataSourceResponseOutput {
	return o
}

func (o PerfCounterDataSourceResponseOutput) CounterSpecifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PerfCounterDataSourceResponse) []string { return v.CounterSpecifiers }).(pulumi.StringArrayOutput)
}

func (o PerfCounterDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PerfCounterDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PerfCounterDataSourceResponseOutput) SamplingFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PerfCounterDataSourceResponse) *int { return v.SamplingFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o PerfCounterDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PerfCounterDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type PerfCounterDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PerfCounterDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerfCounterDataSourceResponse)(nil)).Elem()
}

func (o PerfCounterDataSourceResponseArrayOutput) ToPerfCounterDataSourceResponseArrayOutput() PerfCounterDataSourceResponseArrayOutput {
	return o
}

func (o PerfCounterDataSourceResponseArrayOutput) ToPerfCounterDataSourceResponseArrayOutputWithContext(ctx context.Context) PerfCounterDataSourceResponseArrayOutput {
	return o
}

func (o PerfCounterDataSourceResponseArrayOutput) Index(i pulumi.IntInput) PerfCounterDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerfCounterDataSourceResponse {
		return vs[0].([]PerfCounterDataSourceResponse)[vs[1].(int)]
	}).(PerfCounterDataSourceResponseOutput)
}

type StreamDeclaration struct {
	Columns []ColumnDefinition `pulumi:"columns"`
}





type StreamDeclarationInput interface {
	pulumi.Input

	ToStreamDeclarationOutput() StreamDeclarationOutput
	ToStreamDeclarationOutputWithContext(context.Context) StreamDeclarationOutput
}

type StreamDeclarationArgs struct {
	Columns ColumnDefinitionArrayInput `pulumi:"columns"`
}

func (StreamDeclarationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamDeclaration)(nil)).Elem()
}

func (i StreamDeclarationArgs) ToStreamDeclarationOutput() StreamDeclarationOutput {
	return i.ToStreamDeclarationOutputWithContext(context.Background())
}

func (i StreamDeclarationArgs) ToStreamDeclarationOutputWithContext(ctx context.Context) StreamDeclarationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamDeclarationOutput)
}





type StreamDeclarationMapInput interface {
	pulumi.Input

	ToStreamDeclarationMapOutput() StreamDeclarationMapOutput
	ToStreamDeclarationMapOutputWithContext(context.Context) StreamDeclarationMapOutput
}

type StreamDeclarationMap map[string]StreamDeclarationInput

func (StreamDeclarationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StreamDeclaration)(nil)).Elem()
}

func (i StreamDeclarationMap) ToStreamDeclarationMapOutput() StreamDeclarationMapOutput {
	return i.ToStreamDeclarationMapOutputWithContext(context.Background())
}

func (i StreamDeclarationMap) ToStreamDeclarationMapOutputWithContext(ctx context.Context) StreamDeclarationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamDeclarationMapOutput)
}

type StreamDeclarationOutput struct{ *pulumi.OutputState }

func (StreamDeclarationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamDeclaration)(nil)).Elem()
}

func (o StreamDeclarationOutput) ToStreamDeclarationOutput() StreamDeclarationOutput {
	return o
}

func (o StreamDeclarationOutput) ToStreamDeclarationOutputWithContext(ctx context.Context) StreamDeclarationOutput {
	return o
}

func (o StreamDeclarationOutput) Columns() ColumnDefinitionArrayOutput {
	return o.ApplyT(func(v StreamDeclaration) []ColumnDefinition { return v.Columns }).(ColumnDefinitionArrayOutput)
}

type StreamDeclarationMapOutput struct{ *pulumi.OutputState }

func (StreamDeclarationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StreamDeclaration)(nil)).Elem()
}

func (o StreamDeclarationMapOutput) ToStreamDeclarationMapOutput() StreamDeclarationMapOutput {
	return o
}

func (o StreamDeclarationMapOutput) ToStreamDeclarationMapOutputWithContext(ctx context.Context) StreamDeclarationMapOutput {
	return o
}

func (o StreamDeclarationMapOutput) MapIndex(k pulumi.StringInput) StreamDeclarationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StreamDeclaration {
		return vs[0].(map[string]StreamDeclaration)[vs[1].(string)]
	}).(StreamDeclarationOutput)
}

type StreamDeclarationResponse struct {
	Columns []ColumnDefinitionResponse `pulumi:"columns"`
}

type StreamDeclarationResponseOutput struct{ *pulumi.OutputState }

func (StreamDeclarationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamDeclarationResponse)(nil)).Elem()
}

func (o StreamDeclarationResponseOutput) ToStreamDeclarationResponseOutput() StreamDeclarationResponseOutput {
	return o
}

func (o StreamDeclarationResponseOutput) ToStreamDeclarationResponseOutputWithContext(ctx context.Context) StreamDeclarationResponseOutput {
	return o
}

func (o StreamDeclarationResponseOutput) Columns() ColumnDefinitionResponseArrayOutput {
	return o.ApplyT(func(v StreamDeclarationResponse) []ColumnDefinitionResponse { return v.Columns }).(ColumnDefinitionResponseArrayOutput)
}

type StreamDeclarationResponseMapOutput struct{ *pulumi.OutputState }

func (StreamDeclarationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StreamDeclarationResponse)(nil)).Elem()
}

func (o StreamDeclarationResponseMapOutput) ToStreamDeclarationResponseMapOutput() StreamDeclarationResponseMapOutput {
	return o
}

func (o StreamDeclarationResponseMapOutput) ToStreamDeclarationResponseMapOutputWithContext(ctx context.Context) StreamDeclarationResponseMapOutput {
	return o
}

func (o StreamDeclarationResponseMapOutput) MapIndex(k pulumi.StringInput) StreamDeclarationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StreamDeclarationResponse {
		return vs[0].(map[string]StreamDeclarationResponse)[vs[1].(string)]
	}).(StreamDeclarationResponseOutput)
}

type SyslogDataSource struct {
	FacilityNames []string `pulumi:"facilityNames"`
	LogLevels     []string `pulumi:"logLevels"`
	Name          *string  `pulumi:"name"`
	Streams       []string `pulumi:"streams"`
}





type SyslogDataSourceInput interface {
	pulumi.Input

	ToSyslogDataSourceOutput() SyslogDataSourceOutput
	ToSyslogDataSourceOutputWithContext(context.Context) SyslogDataSourceOutput
}

type SyslogDataSourceArgs struct {
	FacilityNames pulumi.StringArrayInput `pulumi:"facilityNames"`
	LogLevels     pulumi.StringArrayInput `pulumi:"logLevels"`
	Name          pulumi.StringPtrInput   `pulumi:"name"`
	Streams       pulumi.StringArrayInput `pulumi:"streams"`
}

func (SyslogDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyslogDataSource)(nil)).Elem()
}

func (i SyslogDataSourceArgs) ToSyslogDataSourceOutput() SyslogDataSourceOutput {
	return i.ToSyslogDataSourceOutputWithContext(context.Background())
}

func (i SyslogDataSourceArgs) ToSyslogDataSourceOutputWithContext(ctx context.Context) SyslogDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyslogDataSourceOutput)
}





type SyslogDataSourceArrayInput interface {
	pulumi.Input

	ToSyslogDataSourceArrayOutput() SyslogDataSourceArrayOutput
	ToSyslogDataSourceArrayOutputWithContext(context.Context) SyslogDataSourceArrayOutput
}

type SyslogDataSourceArray []SyslogDataSourceInput

func (SyslogDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyslogDataSource)(nil)).Elem()
}

func (i SyslogDataSourceArray) ToSyslogDataSourceArrayOutput() SyslogDataSourceArrayOutput {
	return i.ToSyslogDataSourceArrayOutputWithContext(context.Background())
}

func (i SyslogDataSourceArray) ToSyslogDataSourceArrayOutputWithContext(ctx context.Context) SyslogDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyslogDataSourceArrayOutput)
}

type SyslogDataSourceOutput struct{ *pulumi.OutputState }

func (SyslogDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyslogDataSource)(nil)).Elem()
}

func (o SyslogDataSourceOutput) ToSyslogDataSourceOutput() SyslogDataSourceOutput {
	return o
}

func (o SyslogDataSourceOutput) ToSyslogDataSourceOutputWithContext(ctx context.Context) SyslogDataSourceOutput {
	return o
}

func (o SyslogDataSourceOutput) FacilityNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSource) []string { return v.FacilityNames }).(pulumi.StringArrayOutput)
}

func (o SyslogDataSourceOutput) LogLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSource) []string { return v.LogLevels }).(pulumi.StringArrayOutput)
}

func (o SyslogDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyslogDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SyslogDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type SyslogDataSourceArrayOutput struct{ *pulumi.OutputState }

func (SyslogDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyslogDataSource)(nil)).Elem()
}

func (o SyslogDataSourceArrayOutput) ToSyslogDataSourceArrayOutput() SyslogDataSourceArrayOutput {
	return o
}

func (o SyslogDataSourceArrayOutput) ToSyslogDataSourceArrayOutputWithContext(ctx context.Context) SyslogDataSourceArrayOutput {
	return o
}

func (o SyslogDataSourceArrayOutput) Index(i pulumi.IntInput) SyslogDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyslogDataSource {
		return vs[0].([]SyslogDataSource)[vs[1].(int)]
	}).(SyslogDataSourceOutput)
}

type SyslogDataSourceResponse struct {
	FacilityNames []string `pulumi:"facilityNames"`
	LogLevels     []string `pulumi:"logLevels"`
	Name          *string  `pulumi:"name"`
	Streams       []string `pulumi:"streams"`
}

type SyslogDataSourceResponseOutput struct{ *pulumi.OutputState }

func (SyslogDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyslogDataSourceResponse)(nil)).Elem()
}

func (o SyslogDataSourceResponseOutput) ToSyslogDataSourceResponseOutput() SyslogDataSourceResponseOutput {
	return o
}

func (o SyslogDataSourceResponseOutput) ToSyslogDataSourceResponseOutputWithContext(ctx context.Context) SyslogDataSourceResponseOutput {
	return o
}

func (o SyslogDataSourceResponseOutput) FacilityNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSourceResponse) []string { return v.FacilityNames }).(pulumi.StringArrayOutput)
}

func (o SyslogDataSourceResponseOutput) LogLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSourceResponse) []string { return v.LogLevels }).(pulumi.StringArrayOutput)
}

func (o SyslogDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyslogDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SyslogDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type SyslogDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SyslogDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyslogDataSourceResponse)(nil)).Elem()
}

func (o SyslogDataSourceResponseArrayOutput) ToSyslogDataSourceResponseArrayOutput() SyslogDataSourceResponseArrayOutput {
	return o
}

func (o SyslogDataSourceResponseArrayOutput) ToSyslogDataSourceResponseArrayOutputWithContext(ctx context.Context) SyslogDataSourceResponseArrayOutput {
	return o
}

func (o SyslogDataSourceResponseArrayOutput) Index(i pulumi.IntInput) SyslogDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyslogDataSourceResponse {
		return vs[0].([]SyslogDataSourceResponse)[vs[1].(int)]
	}).(SyslogDataSourceResponseOutput)
}

type WindowsEventLogDataSource struct {
	Name         *string  `pulumi:"name"`
	Streams      []string `pulumi:"streams"`
	XPathQueries []string `pulumi:"xPathQueries"`
}





type WindowsEventLogDataSourceInput interface {
	pulumi.Input

	ToWindowsEventLogDataSourceOutput() WindowsEventLogDataSourceOutput
	ToWindowsEventLogDataSourceOutputWithContext(context.Context) WindowsEventLogDataSourceOutput
}

type WindowsEventLogDataSourceArgs struct {
	Name         pulumi.StringPtrInput   `pulumi:"name"`
	Streams      pulumi.StringArrayInput `pulumi:"streams"`
	XPathQueries pulumi.StringArrayInput `pulumi:"xPathQueries"`
}

func (WindowsEventLogDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsEventLogDataSource)(nil)).Elem()
}

func (i WindowsEventLogDataSourceArgs) ToWindowsEventLogDataSourceOutput() WindowsEventLogDataSourceOutput {
	return i.ToWindowsEventLogDataSourceOutputWithContext(context.Background())
}

func (i WindowsEventLogDataSourceArgs) ToWindowsEventLogDataSourceOutputWithContext(ctx context.Context) WindowsEventLogDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsEventLogDataSourceOutput)
}





type WindowsEventLogDataSourceArrayInput interface {
	pulumi.Input

	ToWindowsEventLogDataSourceArrayOutput() WindowsEventLogDataSourceArrayOutput
	ToWindowsEventLogDataSourceArrayOutputWithContext(context.Context) WindowsEventLogDataSourceArrayOutput
}

type WindowsEventLogDataSourceArray []WindowsEventLogDataSourceInput

func (WindowsEventLogDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsEventLogDataSource)(nil)).Elem()
}

func (i WindowsEventLogDataSourceArray) ToWindowsEventLogDataSourceArrayOutput() WindowsEventLogDataSourceArrayOutput {
	return i.ToWindowsEventLogDataSourceArrayOutputWithContext(context.Background())
}

func (i WindowsEventLogDataSourceArray) ToWindowsEventLogDataSourceArrayOutputWithContext(ctx context.Context) WindowsEventLogDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsEventLogDataSourceArrayOutput)
}

type WindowsEventLogDataSourceOutput struct{ *pulumi.OutputState }

func (WindowsEventLogDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsEventLogDataSource)(nil)).Elem()
}

func (o WindowsEventLogDataSourceOutput) ToWindowsEventLogDataSourceOutput() WindowsEventLogDataSourceOutput {
	return o
}

func (o WindowsEventLogDataSourceOutput) ToWindowsEventLogDataSourceOutputWithContext(ctx context.Context) WindowsEventLogDataSourceOutput {
	return o
}

func (o WindowsEventLogDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsEventLogDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WindowsEventLogDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsEventLogDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

func (o WindowsEventLogDataSourceOutput) XPathQueries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsEventLogDataSource) []string { return v.XPathQueries }).(pulumi.StringArrayOutput)
}

type WindowsEventLogDataSourceArrayOutput struct{ *pulumi.OutputState }

func (WindowsEventLogDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsEventLogDataSource)(nil)).Elem()
}

func (o WindowsEventLogDataSourceArrayOutput) ToWindowsEventLogDataSourceArrayOutput() WindowsEventLogDataSourceArrayOutput {
	return o
}

func (o WindowsEventLogDataSourceArrayOutput) ToWindowsEventLogDataSourceArrayOutputWithContext(ctx context.Context) WindowsEventLogDataSourceArrayOutput {
	return o
}

func (o WindowsEventLogDataSourceArrayOutput) Index(i pulumi.IntInput) WindowsEventLogDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WindowsEventLogDataSource {
		return vs[0].([]WindowsEventLogDataSource)[vs[1].(int)]
	}).(WindowsEventLogDataSourceOutput)
}

type WindowsEventLogDataSourceResponse struct {
	Name         *string  `pulumi:"name"`
	Streams      []string `pulumi:"streams"`
	XPathQueries []string `pulumi:"xPathQueries"`
}

type WindowsEventLogDataSourceResponseOutput struct{ *pulumi.OutputState }

func (WindowsEventLogDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsEventLogDataSourceResponse)(nil)).Elem()
}

func (o WindowsEventLogDataSourceResponseOutput) ToWindowsEventLogDataSourceResponseOutput() WindowsEventLogDataSourceResponseOutput {
	return o
}

func (o WindowsEventLogDataSourceResponseOutput) ToWindowsEventLogDataSourceResponseOutputWithContext(ctx context.Context) WindowsEventLogDataSourceResponseOutput {
	return o
}

func (o WindowsEventLogDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsEventLogDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WindowsEventLogDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsEventLogDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

func (o WindowsEventLogDataSourceResponseOutput) XPathQueries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsEventLogDataSourceResponse) []string { return v.XPathQueries }).(pulumi.StringArrayOutput)
}

type WindowsEventLogDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (WindowsEventLogDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsEventLogDataSourceResponse)(nil)).Elem()
}

func (o WindowsEventLogDataSourceResponseArrayOutput) ToWindowsEventLogDataSourceResponseArrayOutput() WindowsEventLogDataSourceResponseArrayOutput {
	return o
}

func (o WindowsEventLogDataSourceResponseArrayOutput) ToWindowsEventLogDataSourceResponseArrayOutputWithContext(ctx context.Context) WindowsEventLogDataSourceResponseArrayOutput {
	return o
}

func (o WindowsEventLogDataSourceResponseArrayOutput) Index(i pulumi.IntInput) WindowsEventLogDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WindowsEventLogDataSourceResponse {
		return vs[0].([]WindowsEventLogDataSourceResponse)[vs[1].(int)]
	}).(WindowsEventLogDataSourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ColumnDefinitionOutput{})
	pulumi.RegisterOutputType(ColumnDefinitionArrayOutput{})
	pulumi.RegisterOutputType(ColumnDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ColumnDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointNetworkAclsOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointNetworkAclsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResourceResponseSystemDataOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseConfigurationAccessOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseConfigurationAccessPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseLogsIngestionOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseLogsIngestionPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseNetworkAclsOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseNetworkAclsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleAssociationResponseMetadataOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDataSourcesOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDataSourcesPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDestinationsOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDestinationsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResourceResponseSystemDataOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDataSourcesOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDataSourcesPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDestinationsOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDestinationsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseMetadataOutput{})
	pulumi.RegisterOutputType(DataFlowOutput{})
	pulumi.RegisterOutputType(DataFlowArrayOutput{})
	pulumi.RegisterOutputType(DataFlowResponseOutput{})
	pulumi.RegisterOutputType(DataFlowResponseArrayOutput{})
	pulumi.RegisterOutputType(DestinationsSpecAzureMonitorMetricsOutput{})
	pulumi.RegisterOutputType(DestinationsSpecAzureMonitorMetricsPtrOutput{})
	pulumi.RegisterOutputType(DestinationsSpecResponseAzureMonitorMetricsOutput{})
	pulumi.RegisterOutputType(DestinationsSpecResponseAzureMonitorMetricsPtrOutput{})
	pulumi.RegisterOutputType(ExtensionDataSourceOutput{})
	pulumi.RegisterOutputType(ExtensionDataSourceArrayOutput{})
	pulumi.RegisterOutputType(ExtensionDataSourceResponseOutput{})
	pulumi.RegisterOutputType(ExtensionDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(IisLogsDataSourceOutput{})
	pulumi.RegisterOutputType(IisLogsDataSourceArrayOutput{})
	pulumi.RegisterOutputType(IisLogsDataSourceResponseOutput{})
	pulumi.RegisterOutputType(IisLogsDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationArrayOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationResponseArrayOutput{})
	pulumi.RegisterOutputType(LogFileSettingsResponseTextOutput{})
	pulumi.RegisterOutputType(LogFileSettingsResponseTextPtrOutput{})
	pulumi.RegisterOutputType(LogFileSettingsTextOutput{})
	pulumi.RegisterOutputType(LogFileSettingsTextPtrOutput{})
	pulumi.RegisterOutputType(LogFilesDataSourceOutput{})
	pulumi.RegisterOutputType(LogFilesDataSourceArrayOutput{})
	pulumi.RegisterOutputType(LogFilesDataSourceResponseOutput{})
	pulumi.RegisterOutputType(LogFilesDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(LogFilesDataSourceResponseSettingsOutput{})
	pulumi.RegisterOutputType(LogFilesDataSourceResponseSettingsPtrOutput{})
	pulumi.RegisterOutputType(LogFilesDataSourceSettingsOutput{})
	pulumi.RegisterOutputType(LogFilesDataSourceSettingsPtrOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceArrayOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceResponseOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamDeclarationOutput{})
	pulumi.RegisterOutputType(StreamDeclarationMapOutput{})
	pulumi.RegisterOutputType(StreamDeclarationResponseOutput{})
	pulumi.RegisterOutputType(StreamDeclarationResponseMapOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceArrayOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceResponseOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceArrayOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceResponseOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceResponseArrayOutput{})
}
