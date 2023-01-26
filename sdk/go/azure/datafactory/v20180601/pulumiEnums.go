


package v20180601

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

type AzureSearchIndexWriteBehaviorType string

const (
	AzureSearchIndexWriteBehaviorTypeMerge  = AzureSearchIndexWriteBehaviorType("Merge")
	AzureSearchIndexWriteBehaviorTypeUpload = AzureSearchIndexWriteBehaviorType("Upload")
)

type BigDataPoolReferenceType string

const (
	BigDataPoolReferenceTypeBigDataPoolReference = BigDataPoolReferenceType("BigDataPoolReference")
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

type ConfigurationType string

const (
	ConfigurationTypeDefault    = ConfigurationType("Default")
	ConfigurationTypeCustomized = ConfigurationType("Customized")
	ConfigurationTypeArtifact   = ConfigurationType("Artifact")
)

type CosmosDbConnectionMode string

const (
	CosmosDbConnectionModeGateway = CosmosDbConnectionMode("Gateway")
	CosmosDbConnectionModeDirect  = CosmosDbConnectionMode("Direct")
)

type CosmosDbServicePrincipalCredentialType string

const (
	CosmosDbServicePrincipalCredentialTypeServicePrincipalKey  = CosmosDbServicePrincipalCredentialType("ServicePrincipalKey")
	CosmosDbServicePrincipalCredentialTypeServicePrincipalCert = CosmosDbServicePrincipalCredentialType("ServicePrincipalCert")
)

type CredentialReferenceType string

const (
	CredentialReferenceTypeCredentialReference = CredentialReferenceType("CredentialReference")
)

type DataFlowComputeType string

const (
	DataFlowComputeTypeGeneral          = DataFlowComputeType("General")
	DataFlowComputeTypeMemoryOptimized  = DataFlowComputeType("MemoryOptimized")
	DataFlowComputeTypeComputeOptimized = DataFlowComputeType("ComputeOptimized")
)

type DataFlowReferenceType string

const (
	DataFlowReferenceTypeDataFlowReference = DataFlowReferenceType("DataFlowReference")
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

type DynamicsSinkWriteBehavior string

const (
	DynamicsSinkWriteBehaviorUpsert = DynamicsSinkWriteBehavior("Upsert")
)

type FactoryIdentityType string

const (
	FactoryIdentityTypeSystemAssigned               = FactoryIdentityType("SystemAssigned")
	FactoryIdentityTypeUserAssigned                 = FactoryIdentityType("UserAssigned")
	FactoryIdentityType_SystemAssigned_UserAssigned = FactoryIdentityType("SystemAssigned,UserAssigned")
)

type FtpAuthenticationType string

const (
	FtpAuthenticationTypeBasic     = FtpAuthenticationType("Basic")
	FtpAuthenticationTypeAnonymous = FtpAuthenticationType("Anonymous")
)

type GlobalParameterType string

const (
	GlobalParameterTypeObject = GlobalParameterType("Object")
	GlobalParameterTypeString = GlobalParameterType("String")
	GlobalParameterTypeInt    = GlobalParameterType("Int")
	GlobalParameterTypeFloat  = GlobalParameterType("Float")
	GlobalParameterTypeBool   = GlobalParameterType("Bool")
	GlobalParameterTypeArray  = GlobalParameterType("Array")
)

type GoogleAdWordsAuthenticationType string

const (
	GoogleAdWordsAuthenticationTypeServiceAuthentication = GoogleAdWordsAuthenticationType("ServiceAuthentication")
	GoogleAdWordsAuthenticationTypeUserAuthentication    = GoogleAdWordsAuthenticationType("UserAuthentication")
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

type IntegrationRuntimeSsisCatalogPricingTier string

const (
	IntegrationRuntimeSsisCatalogPricingTierBasic     = IntegrationRuntimeSsisCatalogPricingTier("Basic")
	IntegrationRuntimeSsisCatalogPricingTierStandard  = IntegrationRuntimeSsisCatalogPricingTier("Standard")
	IntegrationRuntimeSsisCatalogPricingTierPremium   = IntegrationRuntimeSsisCatalogPricingTier("Premium")
	IntegrationRuntimeSsisCatalogPricingTierPremiumRS = IntegrationRuntimeSsisCatalogPricingTier("PremiumRS")
)

type IntegrationRuntimeType string

const (
	IntegrationRuntimeTypeManaged    = IntegrationRuntimeType("Managed")
	IntegrationRuntimeTypeSelfHosted = IntegrationRuntimeType("SelfHosted")
)

type ManagedVirtualNetworkReferenceType string

const (
	ManagedVirtualNetworkReferenceTypeManagedVirtualNetworkReference = ManagedVirtualNetworkReferenceType("ManagedVirtualNetworkReference")
)

type MongoDbAuthenticationType string

const (
	MongoDbAuthenticationTypeBasic     = MongoDbAuthenticationType("Basic")
	MongoDbAuthenticationTypeAnonymous = MongoDbAuthenticationType("Anonymous")
)

type NotebookParameterType string

const (
	NotebookParameterTypeString = NotebookParameterType("string")
	NotebookParameterTypeInt    = NotebookParameterType("int")
	NotebookParameterTypeFloat  = NotebookParameterType("float")
	NotebookParameterTypeBool   = NotebookParameterType("bool")
)

type NotebookReferenceType string

const (
	NotebookReferenceTypeNotebookReference = NotebookReferenceType("NotebookReference")
)

type ODataAadServicePrincipalCredentialType string

const (
	ODataAadServicePrincipalCredentialTypeServicePrincipalKey  = ODataAadServicePrincipalCredentialType("ServicePrincipalKey")
	ODataAadServicePrincipalCredentialTypeServicePrincipalCert = ODataAadServicePrincipalCredentialType("ServicePrincipalCert")
)

type ODataAuthenticationType string

const (
	ODataAuthenticationTypeBasic                  = ODataAuthenticationType("Basic")
	ODataAuthenticationTypeAnonymous              = ODataAuthenticationType("Anonymous")
	ODataAuthenticationTypeWindows                = ODataAuthenticationType("Windows")
	ODataAuthenticationTypeAadServicePrincipal    = ODataAuthenticationType("AadServicePrincipal")
	ODataAuthenticationTypeManagedServiceIdentity = ODataAuthenticationType("ManagedServiceIdentity")
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

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
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

type RestServiceAuthenticationType string

const (
	RestServiceAuthenticationTypeAnonymous              = RestServiceAuthenticationType("Anonymous")
	RestServiceAuthenticationTypeBasic                  = RestServiceAuthenticationType("Basic")
	RestServiceAuthenticationTypeAadServicePrincipal    = RestServiceAuthenticationType("AadServicePrincipal")
	RestServiceAuthenticationTypeManagedServiceIdentity = RestServiceAuthenticationType("ManagedServiceIdentity")
	RestServiceAuthenticationTypeOAuth2ClientCredential = RestServiceAuthenticationType("OAuth2ClientCredential")
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

type ScriptActivityLogDestination string

const (
	ScriptActivityLogDestinationActivityOutput = ScriptActivityLogDestination("ActivityOutput")
	ScriptActivityLogDestinationExternalStore  = ScriptActivityLogDestination("ExternalStore")
)

type ScriptActivityParameterDirection string

const (
	ScriptActivityParameterDirectionValueInput       = ScriptActivityParameterDirection("Input")
	ScriptActivityParameterDirectionValueOutput      = ScriptActivityParameterDirection("Output")
	ScriptActivityParameterDirectionValueInputOutput = ScriptActivityParameterDirection("InputOutput")
)

type ScriptActivityParameterType string

const (
	ScriptActivityParameterTypeBoolean        = ScriptActivityParameterType("Boolean")
	ScriptActivityParameterTypeDateTime       = ScriptActivityParameterType("DateTime")
	ScriptActivityParameterTypeDateTimeOffset = ScriptActivityParameterType("DateTimeOffset")
	ScriptActivityParameterTypeDecimal        = ScriptActivityParameterType("Decimal")
	ScriptActivityParameterTypeDouble         = ScriptActivityParameterType("Double")
	ScriptActivityParameterTypeGuid           = ScriptActivityParameterType("Guid")
	ScriptActivityParameterTypeInt16          = ScriptActivityParameterType("Int16")
	ScriptActivityParameterTypeInt32          = ScriptActivityParameterType("Int32")
	ScriptActivityParameterTypeInt64          = ScriptActivityParameterType("Int64")
	ScriptActivityParameterTypeSingle         = ScriptActivityParameterType("Single")
	ScriptActivityParameterTypeString         = ScriptActivityParameterType("String")
	ScriptActivityParameterTypeTimespan       = ScriptActivityParameterType("Timespan")
)

type ScriptType string

const (
	ScriptTypeQuery    = ScriptType("Query")
	ScriptTypeNonQuery = ScriptType("NonQuery")
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
	SftpAuthenticationTypeMultiFactor  = SftpAuthenticationType("MultiFactor")
)

type SparkAuthenticationType string

const (
	SparkAuthenticationTypeAnonymous                    = SparkAuthenticationType("Anonymous")
	SparkAuthenticationTypeUsername                     = SparkAuthenticationType("Username")
	SparkAuthenticationTypeUsernameAndPassword          = SparkAuthenticationType("UsernameAndPassword")
	SparkAuthenticationTypeWindowsAzureHDInsightService = SparkAuthenticationType("WindowsAzureHDInsightService")
)

type SparkConfigurationReferenceType string

const (
	SparkConfigurationReferenceTypeSparkConfigurationReference = SparkConfigurationReferenceType("SparkConfigurationReference")
)

type SparkJobReferenceType string

const (
	SparkJobReferenceTypeSparkJobDefinitionReference = SparkJobReferenceType("SparkJobDefinitionReference")
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

type SqlAlwaysEncryptedAkvAuthType string

const (
	SqlAlwaysEncryptedAkvAuthTypeServicePrincipal            = SqlAlwaysEncryptedAkvAuthType("ServicePrincipal")
	SqlAlwaysEncryptedAkvAuthTypeManagedIdentity             = SqlAlwaysEncryptedAkvAuthType("ManagedIdentity")
	SqlAlwaysEncryptedAkvAuthTypeUserAssignedManagedIdentity = SqlAlwaysEncryptedAkvAuthType("UserAssignedManagedIdentity")
)

type SsisLogLocationType string

const (
	SsisLogLocationTypeFile = SsisLogLocationType("File")
)

type SsisPackageLocationType string

const (
	SsisPackageLocationTypeSSISDB        = SsisPackageLocationType("SSISDB")
	SsisPackageLocationTypeFile          = SsisPackageLocationType("File")
	SsisPackageLocationTypeInlinePackage = SsisPackageLocationType("InlinePackage")
	SsisPackageLocationTypePackageStore  = SsisPackageLocationType("PackageStore")
)

type SybaseAuthenticationType string

const (
	SybaseAuthenticationTypeBasic   = SybaseAuthenticationType("Basic")
	SybaseAuthenticationTypeWindows = SybaseAuthenticationType("Windows")
)

type TeamDeskAuthenticationType string

const (
	TeamDeskAuthenticationTypeBasic = TeamDeskAuthenticationType("Basic")
	TeamDeskAuthenticationTypeToken = TeamDeskAuthenticationType("Token")
)

type TeradataAuthenticationType string

const (
	TeradataAuthenticationTypeBasic   = TeradataAuthenticationType("Basic")
	TeradataAuthenticationTypeWindows = TeradataAuthenticationType("Windows")
)

type TriggerReferenceType string

const (
	TriggerReferenceTypeTriggerReference = TriggerReferenceType("TriggerReference")
)

type TumblingWindowFrequency string

const (
	TumblingWindowFrequencyMinute = TumblingWindowFrequency("Minute")
	TumblingWindowFrequencyHour   = TumblingWindowFrequency("Hour")
	TumblingWindowFrequencyMonth  = TumblingWindowFrequency("Month")
)

type Type string

const (
	TypeLinkedServiceReference = Type("LinkedServiceReference")
)

type VariableType string

const (
	VariableTypeString = VariableType("String")
	VariableTypeBool   = VariableType("Bool")
	VariableTypeArray  = VariableType("Array")
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

type WebHookActivityMethod string

const (
	WebHookActivityMethodPOST = WebHookActivityMethod("POST")
)

type ZendeskAuthenticationType string

const (
	ZendeskAuthenticationTypeBasic = ZendeskAuthenticationType("Basic")
	ZendeskAuthenticationTypeToken = ZendeskAuthenticationType("Token")
)

func init() {
}
