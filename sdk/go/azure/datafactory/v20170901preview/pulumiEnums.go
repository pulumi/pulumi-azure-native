


package v20170901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureSearchIndexWriteBehaviorType string

const (
	AzureSearchIndexWriteBehaviorTypeMerge  = AzureSearchIndexWriteBehaviorType("Merge")
	AzureSearchIndexWriteBehaviorTypeUpload = AzureSearchIndexWriteBehaviorType("Upload")
)

type BlobEventTypes string

const (
	BlobEventTypes_Microsoft_Storage_BlobCreated = BlobEventTypes("Microsoft.Storage.BlobCreated")
	BlobEventTypes_Microsoft_Storage_BlobDeleted = BlobEventTypes("Microsoft.Storage.BlobDeleted")
)

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

type DatasetCompressionLevel string

const (
	DatasetCompressionLevelOptimal = DatasetCompressionLevel("Optimal")
	DatasetCompressionLevelFastest = DatasetCompressionLevel("Fastest")
)

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

type Db2AuthenticationType string

const (
	Db2AuthenticationTypeBasic = Db2AuthenticationType("Basic")
)

type DependencyCondition string

const (
	DependencyConditionSucceeded = DependencyCondition("Succeeded")
	DependencyConditionFailed    = DependencyCondition("Failed")
	DependencyConditionSkipped   = DependencyCondition("Skipped")
	DependencyConditionCompleted = DependencyCondition("Completed")
)

type DynamicsAuthenticationType string

const (
	DynamicsAuthenticationTypeOffice365 = DynamicsAuthenticationType("Office365")
	DynamicsAuthenticationTypeIfd       = DynamicsAuthenticationType("Ifd")
)

type DynamicsDeploymentType string

const (
	DynamicsDeploymentTypeOnline            = DynamicsDeploymentType("Online")
	DynamicsDeploymentTypeOnPremisesWithIfd = DynamicsDeploymentType("OnPremisesWithIfd")
)

type DynamicsSinkWriteBehavior string

const (
	DynamicsSinkWriteBehaviorUpsert = DynamicsSinkWriteBehavior("Upsert")
)

type FactoryIdentityType string

const (
	FactoryIdentityTypeSystemAssigned = FactoryIdentityType("SystemAssigned")
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

type GoogleBigQueryAuthenticationType string

const (
	GoogleBigQueryAuthenticationTypeServiceAuthentication = GoogleBigQueryAuthenticationType("ServiceAuthentication")
	GoogleBigQueryAuthenticationTypeUserAuthentication    = GoogleBigQueryAuthenticationType("UserAuthentication")
)

type HBaseAuthenticationType string

const (
	HBaseAuthenticationTypeAnonymous = HBaseAuthenticationType("Anonymous")
	HBaseAuthenticationTypeBasic     = HBaseAuthenticationType("Basic")
)

type HDInsightActivityDebugInfoOption string

const (
	HDInsightActivityDebugInfoOptionNone    = HDInsightActivityDebugInfoOption("None")
	HDInsightActivityDebugInfoOptionAlways  = HDInsightActivityDebugInfoOption("Always")
	HDInsightActivityDebugInfoOptionFailure = HDInsightActivityDebugInfoOption("Failure")
)

type HiveAuthenticationType string

const (
	HiveAuthenticationTypeAnonymous                    = HiveAuthenticationType("Anonymous")
	HiveAuthenticationTypeUsername                     = HiveAuthenticationType("Username")
	HiveAuthenticationTypeUsernameAndPassword          = HiveAuthenticationType("UsernameAndPassword")
	HiveAuthenticationTypeWindowsAzureHDInsightService = HiveAuthenticationType("WindowsAzureHDInsightService")
)

type HiveServerType string

const (
	HiveServerTypeHiveServer1      = HiveServerType("HiveServer1")
	HiveServerTypeHiveServer2      = HiveServerType("HiveServer2")
	HiveServerTypeHiveThriftServer = HiveServerType("HiveThriftServer")
)

type HiveThriftTransportProtocol string

const (
	HiveThriftTransportProtocolBinary = HiveThriftTransportProtocol("Binary")
	HiveThriftTransportProtocolSASL   = HiveThriftTransportProtocol("SASL")
	HiveThriftTransportProtocol_HTTP_ = HiveThriftTransportProtocol("HTTP ")
)

type HttpAuthenticationType string

const (
	HttpAuthenticationTypeBasic             = HttpAuthenticationType("Basic")
	HttpAuthenticationTypeAnonymous         = HttpAuthenticationType("Anonymous")
	HttpAuthenticationTypeDigest            = HttpAuthenticationType("Digest")
	HttpAuthenticationTypeWindows           = HttpAuthenticationType("Windows")
	HttpAuthenticationTypeClientCertificate = HttpAuthenticationType("ClientCertificate")
)

type ImpalaAuthenticationType string

const (
	ImpalaAuthenticationTypeAnonymous           = ImpalaAuthenticationType("Anonymous")
	ImpalaAuthenticationTypeSASLUsername        = ImpalaAuthenticationType("SASLUsername")
	ImpalaAuthenticationTypeUsernameAndPassword = ImpalaAuthenticationType("UsernameAndPassword")
)

type IntegrationRuntimeEdition string

const (
	IntegrationRuntimeEditionStandard   = IntegrationRuntimeEdition("Standard")
	IntegrationRuntimeEditionEnterprise = IntegrationRuntimeEdition("Enterprise")
)

type IntegrationRuntimeEntityReferenceType string

const (
	IntegrationRuntimeEntityReferenceTypeIntegrationRuntimeReference = IntegrationRuntimeEntityReferenceType("IntegrationRuntimeReference")
	IntegrationRuntimeEntityReferenceTypeLinkedServiceReference      = IntegrationRuntimeEntityReferenceType("LinkedServiceReference")
)

type IntegrationRuntimeLicenseType string

const (
	IntegrationRuntimeLicenseTypeBasePrice       = IntegrationRuntimeLicenseType("BasePrice")
	IntegrationRuntimeLicenseTypeLicenseIncluded = IntegrationRuntimeLicenseType("LicenseIncluded")
)

type IntegrationRuntimeType string

const (
	IntegrationRuntimeTypeManaged    = IntegrationRuntimeType("Managed")
	IntegrationRuntimeTypeSelfHosted = IntegrationRuntimeType("SelfHosted")
)

type JsonFormatFilePattern string

const (
	JsonFormatFilePatternSetOfObjects   = JsonFormatFilePattern("setOfObjects")
	JsonFormatFilePatternArrayOfObjects = JsonFormatFilePattern("arrayOfObjects")
)

type MongoDbAuthenticationType string

const (
	MongoDbAuthenticationTypeBasic     = MongoDbAuthenticationType("Basic")
	MongoDbAuthenticationTypeAnonymous = MongoDbAuthenticationType("Anonymous")
)

type ODataAuthenticationType string

const (
	ODataAuthenticationTypeBasic     = ODataAuthenticationType("Basic")
	ODataAuthenticationTypeAnonymous = ODataAuthenticationType("Anonymous")
)

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

type PhoenixAuthenticationType string

const (
	PhoenixAuthenticationTypeAnonymous                    = PhoenixAuthenticationType("Anonymous")
	PhoenixAuthenticationTypeUsernameAndPassword          = PhoenixAuthenticationType("UsernameAndPassword")
	PhoenixAuthenticationTypeWindowsAzureHDInsightService = PhoenixAuthenticationType("WindowsAzureHDInsightService")
)

type PolybaseSettingsRejectType string

const (
	PolybaseSettingsRejectTypeValue      = PolybaseSettingsRejectType("value")
	PolybaseSettingsRejectTypePercentage = PolybaseSettingsRejectType("percentage")
)

type PrestoAuthenticationType string

const (
	PrestoAuthenticationTypeAnonymous = PrestoAuthenticationType("Anonymous")
	PrestoAuthenticationTypeLDAP      = PrestoAuthenticationType("LDAP")
)

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

type SalesforceSinkWriteBehavior string

const (
	SalesforceSinkWriteBehaviorInsert = SalesforceSinkWriteBehavior("Insert")
	SalesforceSinkWriteBehaviorUpsert = SalesforceSinkWriteBehavior("Upsert")
)

type SalesforceSourceReadBehavior string

const (
	SalesforceSourceReadBehaviorQuery    = SalesforceSourceReadBehavior("Query")
	SalesforceSourceReadBehaviorQueryAll = SalesforceSourceReadBehavior("QueryAll")
)

type SapCloudForCustomerSinkWriteBehavior string

const (
	SapCloudForCustomerSinkWriteBehaviorInsert = SapCloudForCustomerSinkWriteBehavior("Insert")
	SapCloudForCustomerSinkWriteBehaviorUpdate = SapCloudForCustomerSinkWriteBehavior("Update")
)

type SapHanaAuthenticationType string

const (
	SapHanaAuthenticationTypeBasic   = SapHanaAuthenticationType("Basic")
	SapHanaAuthenticationTypeWindows = SapHanaAuthenticationType("Windows")
)

type ServiceNowAuthenticationType string

const (
	ServiceNowAuthenticationTypeBasic  = ServiceNowAuthenticationType("Basic")
	ServiceNowAuthenticationTypeOAuth2 = ServiceNowAuthenticationType("OAuth2")
)

type SftpAuthenticationType string

const (
	SftpAuthenticationTypeBasic        = SftpAuthenticationType("Basic")
	SftpAuthenticationTypeSshPublicKey = SftpAuthenticationType("SshPublicKey")
)

type SparkAuthenticationType string

const (
	SparkAuthenticationTypeAnonymous                    = SparkAuthenticationType("Anonymous")
	SparkAuthenticationTypeUsername                     = SparkAuthenticationType("Username")
	SparkAuthenticationTypeUsernameAndPassword          = SparkAuthenticationType("UsernameAndPassword")
	SparkAuthenticationTypeWindowsAzureHDInsightService = SparkAuthenticationType("WindowsAzureHDInsightService")
)

type SparkServerType string

const (
	SparkServerTypeSharkServer       = SparkServerType("SharkServer")
	SparkServerTypeSharkServer2      = SparkServerType("SharkServer2")
	SparkServerTypeSparkThriftServer = SparkServerType("SparkThriftServer")
)

type SparkThriftTransportProtocol string

const (
	SparkThriftTransportProtocolBinary = SparkThriftTransportProtocol("Binary")
	SparkThriftTransportProtocolSASL   = SparkThriftTransportProtocol("SASL")
	SparkThriftTransportProtocol_HTTP_ = SparkThriftTransportProtocol("HTTP ")
)

type SsisLogLocationType string

const (
	SsisLogLocationTypeFile = SsisLogLocationType("File")
)

type SsisPackageLocationType string

const (
	SsisPackageLocationTypeSSISDB = SsisPackageLocationType("SSISDB")
	SsisPackageLocationTypeFile   = SsisPackageLocationType("File")
)

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

type SybaseAuthenticationType string

const (
	SybaseAuthenticationTypeBasic   = SybaseAuthenticationType("Basic")
	SybaseAuthenticationTypeWindows = SybaseAuthenticationType("Windows")
)

type TeradataAuthenticationType string

const (
	TeradataAuthenticationTypeBasic   = TeradataAuthenticationType("Basic")
	TeradataAuthenticationTypeWindows = TeradataAuthenticationType("Windows")
)

type TumblingWindowFrequency string

const (
	TumblingWindowFrequencyMinute = TumblingWindowFrequency("Minute")
	TumblingWindowFrequencyHour   = TumblingWindowFrequency("Hour")
)

type WebActivityMethod string

const (
	WebActivityMethodGET    = WebActivityMethod("GET")
	WebActivityMethodPOST   = WebActivityMethod("POST")
	WebActivityMethodPUT    = WebActivityMethod("PUT")
	WebActivityMethodDELETE = WebActivityMethod("DELETE")
)

type WebAuthenticationType string

const (
	WebAuthenticationTypeBasic             = WebAuthenticationType("Basic")
	WebAuthenticationTypeAnonymous         = WebAuthenticationType("Anonymous")
	WebAuthenticationTypeClientCertificate = WebAuthenticationType("ClientCertificate")
)

func init() {
	pulumi.RegisterOutputType(FactoryIdentityTypeOutput{})
	pulumi.RegisterOutputType(FactoryIdentityTypePtrOutput{})
}
