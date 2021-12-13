


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

type DataCollectionRuleDataSources struct {
	Extensions          []ExtensionDataSource       `pulumi:"extensions"`
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

type DataFlow struct {
	Destinations []string `pulumi:"destinations"`
	Streams      []string `pulumi:"streams"`
}





type DataFlowInput interface {
	pulumi.Input

	ToDataFlowOutput() DataFlowOutput
	ToDataFlowOutputWithContext(context.Context) DataFlowOutput
}

type DataFlowArgs struct {
	Destinations pulumi.StringArrayInput `pulumi:"destinations"`
	Streams      pulumi.StringArrayInput `pulumi:"streams"`
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

func (o DataFlowOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlow) []string { return v.Streams }).(pulumi.StringArrayOutput)
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
	Streams      []string `pulumi:"streams"`
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

func (o DataFlowResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlowResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
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
	pulumi.RegisterOutputType(DataCollectionRuleDataSourcesOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDataSourcesPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDestinationsOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDestinationsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResourceResponseSystemDataOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDataSourcesOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDataSourcesPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDestinationsOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDestinationsPtrOutput{})
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
	pulumi.RegisterOutputType(LogAnalyticsDestinationOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationArrayOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationResponseArrayOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceArrayOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceResponseOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceArrayOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceResponseOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceArrayOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceResponseOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceResponseArrayOutput{})
}
