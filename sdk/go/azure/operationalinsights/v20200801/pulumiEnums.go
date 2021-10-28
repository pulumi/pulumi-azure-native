


package v20200801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterSkuNameEnum string

const (
	ClusterSkuNameEnumCapacityReservation = ClusterSkuNameEnum("CapacityReservation")
)

func (ClusterSkuNameEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuNameEnum)(nil)).Elem()
}

func (e ClusterSkuNameEnum) ToClusterSkuNameEnumOutput() ClusterSkuNameEnumOutput {
	return pulumi.ToOutput(e).(ClusterSkuNameEnumOutput)
}

func (e ClusterSkuNameEnum) ToClusterSkuNameEnumOutputWithContext(ctx context.Context) ClusterSkuNameEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterSkuNameEnumOutput)
}

func (e ClusterSkuNameEnum) ToClusterSkuNameEnumPtrOutput() ClusterSkuNameEnumPtrOutput {
	return e.ToClusterSkuNameEnumPtrOutputWithContext(context.Background())
}

func (e ClusterSkuNameEnum) ToClusterSkuNameEnumPtrOutputWithContext(ctx context.Context) ClusterSkuNameEnumPtrOutput {
	return ClusterSkuNameEnum(e).ToClusterSkuNameEnumOutputWithContext(ctx).ToClusterSkuNameEnumPtrOutputWithContext(ctx)
}

func (e ClusterSkuNameEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterSkuNameEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterSkuNameEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterSkuNameEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterSkuNameEnumOutput struct{ *pulumi.OutputState }

func (ClusterSkuNameEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuNameEnum)(nil)).Elem()
}

func (o ClusterSkuNameEnumOutput) ToClusterSkuNameEnumOutput() ClusterSkuNameEnumOutput {
	return o
}

func (o ClusterSkuNameEnumOutput) ToClusterSkuNameEnumOutputWithContext(ctx context.Context) ClusterSkuNameEnumOutput {
	return o
}

func (o ClusterSkuNameEnumOutput) ToClusterSkuNameEnumPtrOutput() ClusterSkuNameEnumPtrOutput {
	return o.ToClusterSkuNameEnumPtrOutputWithContext(context.Background())
}

func (o ClusterSkuNameEnumOutput) ToClusterSkuNameEnumPtrOutputWithContext(ctx context.Context) ClusterSkuNameEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSkuNameEnum) *ClusterSkuNameEnum {
		return &v
	}).(ClusterSkuNameEnumPtrOutput)
}

func (o ClusterSkuNameEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterSkuNameEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterSkuNameEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterSkuNameEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterSkuNameEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterSkuNameEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterSkuNameEnumPtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuNameEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSkuNameEnum)(nil)).Elem()
}

func (o ClusterSkuNameEnumPtrOutput) ToClusterSkuNameEnumPtrOutput() ClusterSkuNameEnumPtrOutput {
	return o
}

func (o ClusterSkuNameEnumPtrOutput) ToClusterSkuNameEnumPtrOutputWithContext(ctx context.Context) ClusterSkuNameEnumPtrOutput {
	return o
}

func (o ClusterSkuNameEnumPtrOutput) Elem() ClusterSkuNameEnumOutput {
	return o.ApplyT(func(v *ClusterSkuNameEnum) ClusterSkuNameEnum {
		if v != nil {
			return *v
		}
		var ret ClusterSkuNameEnum
		return ret
	}).(ClusterSkuNameEnumOutput)
}

func (o ClusterSkuNameEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterSkuNameEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterSkuNameEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClusterSkuNameEnumInput interface {
	pulumi.Input

	ToClusterSkuNameEnumOutput() ClusterSkuNameEnumOutput
	ToClusterSkuNameEnumOutputWithContext(context.Context) ClusterSkuNameEnumOutput
}

var clusterSkuNameEnumPtrType = reflect.TypeOf((**ClusterSkuNameEnum)(nil)).Elem()

type ClusterSkuNameEnumPtrInput interface {
	pulumi.Input

	ToClusterSkuNameEnumPtrOutput() ClusterSkuNameEnumPtrOutput
	ToClusterSkuNameEnumPtrOutputWithContext(context.Context) ClusterSkuNameEnumPtrOutput
}

type clusterSkuNameEnumPtr string

func ClusterSkuNameEnumPtr(v string) ClusterSkuNameEnumPtrInput {
	return (*clusterSkuNameEnumPtr)(&v)
}

func (*clusterSkuNameEnumPtr) ElementType() reflect.Type {
	return clusterSkuNameEnumPtrType
}

func (in *clusterSkuNameEnumPtr) ToClusterSkuNameEnumPtrOutput() ClusterSkuNameEnumPtrOutput {
	return pulumi.ToOutput(in).(ClusterSkuNameEnumPtrOutput)
}

func (in *clusterSkuNameEnumPtr) ToClusterSkuNameEnumPtrOutputWithContext(ctx context.Context) ClusterSkuNameEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterSkuNameEnumPtrOutput)
}

type DataSourceKind string

const (
	DataSourceKindWindowsEvent                                         = DataSourceKind("WindowsEvent")
	DataSourceKindWindowsPerformanceCounter                            = DataSourceKind("WindowsPerformanceCounter")
	DataSourceKindIISLogs                                              = DataSourceKind("IISLogs")
	DataSourceKindLinuxSyslog                                          = DataSourceKind("LinuxSyslog")
	DataSourceKindLinuxSyslogCollection                                = DataSourceKind("LinuxSyslogCollection")
	DataSourceKindLinuxPerformanceObject                               = DataSourceKind("LinuxPerformanceObject")
	DataSourceKindLinuxPerformanceCollection                           = DataSourceKind("LinuxPerformanceCollection")
	DataSourceKindCustomLog                                            = DataSourceKind("CustomLog")
	DataSourceKindCustomLogCollection                                  = DataSourceKind("CustomLogCollection")
	DataSourceKindAzureAuditLog                                        = DataSourceKind("AzureAuditLog")
	DataSourceKindAzureActivityLog                                     = DataSourceKind("AzureActivityLog")
	DataSourceKindGenericDataSource                                    = DataSourceKind("GenericDataSource")
	DataSourceKindChangeTrackingCustomPath                             = DataSourceKind("ChangeTrackingCustomPath")
	DataSourceKindChangeTrackingPath                                   = DataSourceKind("ChangeTrackingPath")
	DataSourceKindChangeTrackingServices                               = DataSourceKind("ChangeTrackingServices")
	DataSourceKindChangeTrackingDataTypeConfiguration                  = DataSourceKind("ChangeTrackingDataTypeConfiguration")
	DataSourceKindChangeTrackingDefaultRegistry                        = DataSourceKind("ChangeTrackingDefaultRegistry")
	DataSourceKindChangeTrackingRegistry                               = DataSourceKind("ChangeTrackingRegistry")
	DataSourceKindChangeTrackingLinuxPath                              = DataSourceKind("ChangeTrackingLinuxPath")
	DataSourceKindLinuxChangeTrackingPath                              = DataSourceKind("LinuxChangeTrackingPath")
	DataSourceKindChangeTrackingContentLocation                        = DataSourceKind("ChangeTrackingContentLocation")
	DataSourceKindWindowsTelemetry                                     = DataSourceKind("WindowsTelemetry")
	DataSourceKindOffice365                                            = DataSourceKind("Office365")
	DataSourceKindSecurityWindowsBaselineConfiguration                 = DataSourceKind("SecurityWindowsBaselineConfiguration")
	DataSourceKindSecurityCenterSecurityWindowsBaselineConfiguration   = DataSourceKind("SecurityCenterSecurityWindowsBaselineConfiguration")
	DataSourceKindSecurityEventCollectionConfiguration                 = DataSourceKind("SecurityEventCollectionConfiguration")
	DataSourceKindSecurityInsightsSecurityEventCollectionConfiguration = DataSourceKind("SecurityInsightsSecurityEventCollectionConfiguration")
	DataSourceKindImportComputerGroup                                  = DataSourceKind("ImportComputerGroup")
	DataSourceKindNetworkMonitoring                                    = DataSourceKind("NetworkMonitoring")
	DataSourceKindItsm                                                 = DataSourceKind("Itsm")
	DataSourceKindDnsAnalytics                                         = DataSourceKind("DnsAnalytics")
	DataSourceKindApplicationInsights                                  = DataSourceKind("ApplicationInsights")
	DataSourceKindSqlDataClassification                                = DataSourceKind("SqlDataClassification")
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

type IdentityType string

const (
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
	IdentityTypeNone           = IdentityType("None")
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

type LinkedServiceEntityStatus string

const (
	LinkedServiceEntityStatusSucceeded           = LinkedServiceEntityStatus("Succeeded")
	LinkedServiceEntityStatusDeleting            = LinkedServiceEntityStatus("Deleting")
	LinkedServiceEntityStatusProvisioningAccount = LinkedServiceEntityStatus("ProvisioningAccount")
	LinkedServiceEntityStatusUpdating            = LinkedServiceEntityStatus("Updating")
)

func (LinkedServiceEntityStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceEntityStatus)(nil)).Elem()
}

func (e LinkedServiceEntityStatus) ToLinkedServiceEntityStatusOutput() LinkedServiceEntityStatusOutput {
	return pulumi.ToOutput(e).(LinkedServiceEntityStatusOutput)
}

func (e LinkedServiceEntityStatus) ToLinkedServiceEntityStatusOutputWithContext(ctx context.Context) LinkedServiceEntityStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LinkedServiceEntityStatusOutput)
}

func (e LinkedServiceEntityStatus) ToLinkedServiceEntityStatusPtrOutput() LinkedServiceEntityStatusPtrOutput {
	return e.ToLinkedServiceEntityStatusPtrOutputWithContext(context.Background())
}

func (e LinkedServiceEntityStatus) ToLinkedServiceEntityStatusPtrOutputWithContext(ctx context.Context) LinkedServiceEntityStatusPtrOutput {
	return LinkedServiceEntityStatus(e).ToLinkedServiceEntityStatusOutputWithContext(ctx).ToLinkedServiceEntityStatusPtrOutputWithContext(ctx)
}

func (e LinkedServiceEntityStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LinkedServiceEntityStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LinkedServiceEntityStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LinkedServiceEntityStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LinkedServiceEntityStatusOutput struct{ *pulumi.OutputState }

func (LinkedServiceEntityStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceEntityStatus)(nil)).Elem()
}

func (o LinkedServiceEntityStatusOutput) ToLinkedServiceEntityStatusOutput() LinkedServiceEntityStatusOutput {
	return o
}

func (o LinkedServiceEntityStatusOutput) ToLinkedServiceEntityStatusOutputWithContext(ctx context.Context) LinkedServiceEntityStatusOutput {
	return o
}

func (o LinkedServiceEntityStatusOutput) ToLinkedServiceEntityStatusPtrOutput() LinkedServiceEntityStatusPtrOutput {
	return o.ToLinkedServiceEntityStatusPtrOutputWithContext(context.Background())
}

func (o LinkedServiceEntityStatusOutput) ToLinkedServiceEntityStatusPtrOutputWithContext(ctx context.Context) LinkedServiceEntityStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedServiceEntityStatus) *LinkedServiceEntityStatus {
		return &v
	}).(LinkedServiceEntityStatusPtrOutput)
}

func (o LinkedServiceEntityStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LinkedServiceEntityStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LinkedServiceEntityStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LinkedServiceEntityStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LinkedServiceEntityStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LinkedServiceEntityStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LinkedServiceEntityStatusPtrOutput struct{ *pulumi.OutputState }

func (LinkedServiceEntityStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceEntityStatus)(nil)).Elem()
}

func (o LinkedServiceEntityStatusPtrOutput) ToLinkedServiceEntityStatusPtrOutput() LinkedServiceEntityStatusPtrOutput {
	return o
}

func (o LinkedServiceEntityStatusPtrOutput) ToLinkedServiceEntityStatusPtrOutputWithContext(ctx context.Context) LinkedServiceEntityStatusPtrOutput {
	return o
}

func (o LinkedServiceEntityStatusPtrOutput) Elem() LinkedServiceEntityStatusOutput {
	return o.ApplyT(func(v *LinkedServiceEntityStatus) LinkedServiceEntityStatus {
		if v != nil {
			return *v
		}
		var ret LinkedServiceEntityStatus
		return ret
	}).(LinkedServiceEntityStatusOutput)
}

func (o LinkedServiceEntityStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LinkedServiceEntityStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LinkedServiceEntityStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LinkedServiceEntityStatusInput interface {
	pulumi.Input

	ToLinkedServiceEntityStatusOutput() LinkedServiceEntityStatusOutput
	ToLinkedServiceEntityStatusOutputWithContext(context.Context) LinkedServiceEntityStatusOutput
}

var linkedServiceEntityStatusPtrType = reflect.TypeOf((**LinkedServiceEntityStatus)(nil)).Elem()

type LinkedServiceEntityStatusPtrInput interface {
	pulumi.Input

	ToLinkedServiceEntityStatusPtrOutput() LinkedServiceEntityStatusPtrOutput
	ToLinkedServiceEntityStatusPtrOutputWithContext(context.Context) LinkedServiceEntityStatusPtrOutput
}

type linkedServiceEntityStatusPtr string

func LinkedServiceEntityStatusPtr(v string) LinkedServiceEntityStatusPtrInput {
	return (*linkedServiceEntityStatusPtr)(&v)
}

func (*linkedServiceEntityStatusPtr) ElementType() reflect.Type {
	return linkedServiceEntityStatusPtrType
}

func (in *linkedServiceEntityStatusPtr) ToLinkedServiceEntityStatusPtrOutput() LinkedServiceEntityStatusPtrOutput {
	return pulumi.ToOutput(in).(LinkedServiceEntityStatusPtrOutput)
}

func (in *linkedServiceEntityStatusPtr) ToLinkedServiceEntityStatusPtrOutputWithContext(ctx context.Context) LinkedServiceEntityStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LinkedServiceEntityStatusPtrOutput)
}

type PublicNetworkAccessType string

const (
	// Enables connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Disables public connectivity to Log Analytics through public DNS.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

func (PublicNetworkAccessType) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccessType)(nil)).Elem()
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessTypeOutput)
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypeOutputWithContext(ctx context.Context) PublicNetworkAccessTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessTypeOutput)
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return e.ToPublicNetworkAccessTypePtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return PublicNetworkAccessType(e).ToPublicNetworkAccessTypeOutputWithContext(ctx).ToPublicNetworkAccessTypePtrOutputWithContext(ctx)
}

func (e PublicNetworkAccessType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessTypeOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccessType)(nil)).Elem()
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput {
	return o
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypeOutputWithContext(ctx context.Context) PublicNetworkAccessTypeOutput {
	return o
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return o.ToPublicNetworkAccessTypePtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccessType) *PublicNetworkAccessType {
		return &v
	}).(PublicNetworkAccessTypePtrOutput)
}

func (o PublicNetworkAccessTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccessType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccessType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessTypePtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccessType)(nil)).Elem()
}

func (o PublicNetworkAccessTypePtrOutput) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return o
}

func (o PublicNetworkAccessTypePtrOutput) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return o
}

func (o PublicNetworkAccessTypePtrOutput) Elem() PublicNetworkAccessTypeOutput {
	return o.ApplyT(func(v *PublicNetworkAccessType) PublicNetworkAccessType {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccessType
		return ret
	}).(PublicNetworkAccessTypeOutput)
}

func (o PublicNetworkAccessTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccessType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessTypeInput interface {
	pulumi.Input

	ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput
	ToPublicNetworkAccessTypeOutputWithContext(context.Context) PublicNetworkAccessTypeOutput
}

var publicNetworkAccessTypePtrType = reflect.TypeOf((**PublicNetworkAccessType)(nil)).Elem()

type PublicNetworkAccessTypePtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput
	ToPublicNetworkAccessTypePtrOutputWithContext(context.Context) PublicNetworkAccessTypePtrOutput
}

type publicNetworkAccessTypePtr string

func PublicNetworkAccessTypePtr(v string) PublicNetworkAccessTypePtrInput {
	return (*publicNetworkAccessTypePtr)(&v)
}

func (*publicNetworkAccessTypePtr) ElementType() reflect.Type {
	return publicNetworkAccessTypePtrType
}

func (in *publicNetworkAccessTypePtr) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessTypePtrOutput)
}

func (in *publicNetworkAccessTypePtr) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessTypePtrOutput)
}

type WorkspaceEntityStatus string

const (
	WorkspaceEntityStatusCreating            = WorkspaceEntityStatus("Creating")
	WorkspaceEntityStatusSucceeded           = WorkspaceEntityStatus("Succeeded")
	WorkspaceEntityStatusFailed              = WorkspaceEntityStatus("Failed")
	WorkspaceEntityStatusCanceled            = WorkspaceEntityStatus("Canceled")
	WorkspaceEntityStatusDeleting            = WorkspaceEntityStatus("Deleting")
	WorkspaceEntityStatusProvisioningAccount = WorkspaceEntityStatus("ProvisioningAccount")
	WorkspaceEntityStatusUpdating            = WorkspaceEntityStatus("Updating")
)

func (WorkspaceEntityStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceEntityStatus)(nil)).Elem()
}

func (e WorkspaceEntityStatus) ToWorkspaceEntityStatusOutput() WorkspaceEntityStatusOutput {
	return pulumi.ToOutput(e).(WorkspaceEntityStatusOutput)
}

func (e WorkspaceEntityStatus) ToWorkspaceEntityStatusOutputWithContext(ctx context.Context) WorkspaceEntityStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkspaceEntityStatusOutput)
}

func (e WorkspaceEntityStatus) ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput {
	return e.ToWorkspaceEntityStatusPtrOutputWithContext(context.Background())
}

func (e WorkspaceEntityStatus) ToWorkspaceEntityStatusPtrOutputWithContext(ctx context.Context) WorkspaceEntityStatusPtrOutput {
	return WorkspaceEntityStatus(e).ToWorkspaceEntityStatusOutputWithContext(ctx).ToWorkspaceEntityStatusPtrOutputWithContext(ctx)
}

func (e WorkspaceEntityStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspaceEntityStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspaceEntityStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkspaceEntityStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkspaceEntityStatusOutput struct{ *pulumi.OutputState }

func (WorkspaceEntityStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceEntityStatus)(nil)).Elem()
}

func (o WorkspaceEntityStatusOutput) ToWorkspaceEntityStatusOutput() WorkspaceEntityStatusOutput {
	return o
}

func (o WorkspaceEntityStatusOutput) ToWorkspaceEntityStatusOutputWithContext(ctx context.Context) WorkspaceEntityStatusOutput {
	return o
}

func (o WorkspaceEntityStatusOutput) ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput {
	return o.ToWorkspaceEntityStatusPtrOutputWithContext(context.Background())
}

func (o WorkspaceEntityStatusOutput) ToWorkspaceEntityStatusPtrOutputWithContext(ctx context.Context) WorkspaceEntityStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceEntityStatus) *WorkspaceEntityStatus {
		return &v
	}).(WorkspaceEntityStatusPtrOutput)
}

func (o WorkspaceEntityStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkspaceEntityStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspaceEntityStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkspaceEntityStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspaceEntityStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspaceEntityStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkspaceEntityStatusPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceEntityStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceEntityStatus)(nil)).Elem()
}

func (o WorkspaceEntityStatusPtrOutput) ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput {
	return o
}

func (o WorkspaceEntityStatusPtrOutput) ToWorkspaceEntityStatusPtrOutputWithContext(ctx context.Context) WorkspaceEntityStatusPtrOutput {
	return o
}

func (o WorkspaceEntityStatusPtrOutput) Elem() WorkspaceEntityStatusOutput {
	return o.ApplyT(func(v *WorkspaceEntityStatus) WorkspaceEntityStatus {
		if v != nil {
			return *v
		}
		var ret WorkspaceEntityStatus
		return ret
	}).(WorkspaceEntityStatusOutput)
}

func (o WorkspaceEntityStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspaceEntityStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkspaceEntityStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WorkspaceEntityStatusInput interface {
	pulumi.Input

	ToWorkspaceEntityStatusOutput() WorkspaceEntityStatusOutput
	ToWorkspaceEntityStatusOutputWithContext(context.Context) WorkspaceEntityStatusOutput
}

var workspaceEntityStatusPtrType = reflect.TypeOf((**WorkspaceEntityStatus)(nil)).Elem()

type WorkspaceEntityStatusPtrInput interface {
	pulumi.Input

	ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput
	ToWorkspaceEntityStatusPtrOutputWithContext(context.Context) WorkspaceEntityStatusPtrOutput
}

type workspaceEntityStatusPtr string

func WorkspaceEntityStatusPtr(v string) WorkspaceEntityStatusPtrInput {
	return (*workspaceEntityStatusPtr)(&v)
}

func (*workspaceEntityStatusPtr) ElementType() reflect.Type {
	return workspaceEntityStatusPtrType
}

func (in *workspaceEntityStatusPtr) ToWorkspaceEntityStatusPtrOutput() WorkspaceEntityStatusPtrOutput {
	return pulumi.ToOutput(in).(WorkspaceEntityStatusPtrOutput)
}

func (in *workspaceEntityStatusPtr) ToWorkspaceEntityStatusPtrOutputWithContext(ctx context.Context) WorkspaceEntityStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkspaceEntityStatusPtrOutput)
}

type WorkspaceSkuNameEnum string

const (
	WorkspaceSkuNameEnumFree                = WorkspaceSkuNameEnum("Free")
	WorkspaceSkuNameEnumStandard            = WorkspaceSkuNameEnum("Standard")
	WorkspaceSkuNameEnumPremium             = WorkspaceSkuNameEnum("Premium")
	WorkspaceSkuNameEnumPerNode             = WorkspaceSkuNameEnum("PerNode")
	WorkspaceSkuNameEnumPerGB2018           = WorkspaceSkuNameEnum("PerGB2018")
	WorkspaceSkuNameEnumStandalone          = WorkspaceSkuNameEnum("Standalone")
	WorkspaceSkuNameEnumCapacityReservation = WorkspaceSkuNameEnum("CapacityReservation")
	WorkspaceSkuNameEnumLACluster           = WorkspaceSkuNameEnum("LACluster")
)

func (WorkspaceSkuNameEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSkuNameEnum)(nil)).Elem()
}

func (e WorkspaceSkuNameEnum) ToWorkspaceSkuNameEnumOutput() WorkspaceSkuNameEnumOutput {
	return pulumi.ToOutput(e).(WorkspaceSkuNameEnumOutput)
}

func (e WorkspaceSkuNameEnum) ToWorkspaceSkuNameEnumOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkspaceSkuNameEnumOutput)
}

func (e WorkspaceSkuNameEnum) ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput {
	return e.ToWorkspaceSkuNameEnumPtrOutputWithContext(context.Background())
}

func (e WorkspaceSkuNameEnum) ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumPtrOutput {
	return WorkspaceSkuNameEnum(e).ToWorkspaceSkuNameEnumOutputWithContext(ctx).ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx)
}

func (e WorkspaceSkuNameEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspaceSkuNameEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkspaceSkuNameEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkspaceSkuNameEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkspaceSkuNameEnumOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuNameEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSkuNameEnum)(nil)).Elem()
}

func (o WorkspaceSkuNameEnumOutput) ToWorkspaceSkuNameEnumOutput() WorkspaceSkuNameEnumOutput {
	return o
}

func (o WorkspaceSkuNameEnumOutput) ToWorkspaceSkuNameEnumOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumOutput {
	return o
}

func (o WorkspaceSkuNameEnumOutput) ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput {
	return o.ToWorkspaceSkuNameEnumPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuNameEnumOutput) ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceSkuNameEnum) *WorkspaceSkuNameEnum {
		return &v
	}).(WorkspaceSkuNameEnumPtrOutput)
}

func (o WorkspaceSkuNameEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkspaceSkuNameEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspaceSkuNameEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkspaceSkuNameEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuNameEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkspaceSkuNameEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkspaceSkuNameEnumPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuNameEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSkuNameEnum)(nil)).Elem()
}

func (o WorkspaceSkuNameEnumPtrOutput) ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput {
	return o
}

func (o WorkspaceSkuNameEnumPtrOutput) ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumPtrOutput {
	return o
}

func (o WorkspaceSkuNameEnumPtrOutput) Elem() WorkspaceSkuNameEnumOutput {
	return o.ApplyT(func(v *WorkspaceSkuNameEnum) WorkspaceSkuNameEnum {
		if v != nil {
			return *v
		}
		var ret WorkspaceSkuNameEnum
		return ret
	}).(WorkspaceSkuNameEnumOutput)
}

func (o WorkspaceSkuNameEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuNameEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkspaceSkuNameEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WorkspaceSkuNameEnumInput interface {
	pulumi.Input

	ToWorkspaceSkuNameEnumOutput() WorkspaceSkuNameEnumOutput
	ToWorkspaceSkuNameEnumOutputWithContext(context.Context) WorkspaceSkuNameEnumOutput
}

var workspaceSkuNameEnumPtrType = reflect.TypeOf((**WorkspaceSkuNameEnum)(nil)).Elem()

type WorkspaceSkuNameEnumPtrInput interface {
	pulumi.Input

	ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput
	ToWorkspaceSkuNameEnumPtrOutputWithContext(context.Context) WorkspaceSkuNameEnumPtrOutput
}

type workspaceSkuNameEnumPtr string

func WorkspaceSkuNameEnumPtr(v string) WorkspaceSkuNameEnumPtrInput {
	return (*workspaceSkuNameEnumPtr)(&v)
}

func (*workspaceSkuNameEnumPtr) ElementType() reflect.Type {
	return workspaceSkuNameEnumPtrType
}

func (in *workspaceSkuNameEnumPtr) ToWorkspaceSkuNameEnumPtrOutput() WorkspaceSkuNameEnumPtrOutput {
	return pulumi.ToOutput(in).(WorkspaceSkuNameEnumPtrOutput)
}

func (in *workspaceSkuNameEnumPtr) ToWorkspaceSkuNameEnumPtrOutputWithContext(ctx context.Context) WorkspaceSkuNameEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkspaceSkuNameEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterSkuNameEnumOutput{})
	pulumi.RegisterOutputType(ClusterSkuNameEnumPtrOutput{})
	pulumi.RegisterOutputType(DataSourceKindOutput{})
	pulumi.RegisterOutputType(DataSourceKindPtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceEntityStatusOutput{})
	pulumi.RegisterOutputType(LinkedServiceEntityStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessTypeOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessTypePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceEntityStatusOutput{})
	pulumi.RegisterOutputType(WorkspaceEntityStatusPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuNameEnumOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuNameEnumPtrOutput{})
}
