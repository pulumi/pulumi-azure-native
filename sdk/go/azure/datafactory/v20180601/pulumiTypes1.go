


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HDInsightSparkActivity struct {
	Arguments             []interface{}           `pulumi:"arguments"`
	ClassName             *string                 `pulumi:"className"`
	DependsOn             []ActivityDependency    `pulumi:"dependsOn"`
	Description           *string                 `pulumi:"description"`
	EntryFilePath         interface{}             `pulumi:"entryFilePath"`
	GetDebugInfo          *string                 `pulumi:"getDebugInfo"`
	LinkedServiceName     *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name                  string                  `pulumi:"name"`
	Policy                *ActivityPolicy         `pulumi:"policy"`
	ProxyUser             interface{}             `pulumi:"proxyUser"`
	RootPath              interface{}             `pulumi:"rootPath"`
	SparkConfig           map[string]interface{}  `pulumi:"sparkConfig"`
	SparkJobLinkedService *LinkedServiceReference `pulumi:"sparkJobLinkedService"`
	Type                  string                  `pulumi:"type"`
	UserProperties        []UserProperty          `pulumi:"userProperties"`
}

type HDInsightSparkActivityResponse struct {
	Arguments             []interface{}                   `pulumi:"arguments"`
	ClassName             *string                         `pulumi:"className"`
	DependsOn             []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description           *string                         `pulumi:"description"`
	EntryFilePath         interface{}                     `pulumi:"entryFilePath"`
	GetDebugInfo          *string                         `pulumi:"getDebugInfo"`
	LinkedServiceName     *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name                  string                          `pulumi:"name"`
	Policy                *ActivityPolicyResponse         `pulumi:"policy"`
	ProxyUser             interface{}                     `pulumi:"proxyUser"`
	RootPath              interface{}                     `pulumi:"rootPath"`
	SparkConfig           map[string]interface{}          `pulumi:"sparkConfig"`
	SparkJobLinkedService *LinkedServiceReferenceResponse `pulumi:"sparkJobLinkedService"`
	Type                  string                          `pulumi:"type"`
	UserProperties        []UserPropertyResponse          `pulumi:"userProperties"`
}

type HDInsightStreamingActivity struct {
	Arguments             []interface{}            `pulumi:"arguments"`
	Combiner              interface{}              `pulumi:"combiner"`
	CommandEnvironment    []interface{}            `pulumi:"commandEnvironment"`
	Defines               map[string]interface{}   `pulumi:"defines"`
	DependsOn             []ActivityDependency     `pulumi:"dependsOn"`
	Description           *string                  `pulumi:"description"`
	FileLinkedService     *LinkedServiceReference  `pulumi:"fileLinkedService"`
	FilePaths             []interface{}            `pulumi:"filePaths"`
	GetDebugInfo          *string                  `pulumi:"getDebugInfo"`
	Input                 interface{}              `pulumi:"input"`
	LinkedServiceName     *LinkedServiceReference  `pulumi:"linkedServiceName"`
	Mapper                interface{}              `pulumi:"mapper"`
	Name                  string                   `pulumi:"name"`
	Output                interface{}              `pulumi:"output"`
	Policy                *ActivityPolicy          `pulumi:"policy"`
	Reducer               interface{}              `pulumi:"reducer"`
	StorageLinkedServices []LinkedServiceReference `pulumi:"storageLinkedServices"`
	Type                  string                   `pulumi:"type"`
	UserProperties        []UserProperty           `pulumi:"userProperties"`
}

type HDInsightStreamingActivityResponse struct {
	Arguments             []interface{}                    `pulumi:"arguments"`
	Combiner              interface{}                      `pulumi:"combiner"`
	CommandEnvironment    []interface{}                    `pulumi:"commandEnvironment"`
	Defines               map[string]interface{}           `pulumi:"defines"`
	DependsOn             []ActivityDependencyResponse     `pulumi:"dependsOn"`
	Description           *string                          `pulumi:"description"`
	FileLinkedService     *LinkedServiceReferenceResponse  `pulumi:"fileLinkedService"`
	FilePaths             []interface{}                    `pulumi:"filePaths"`
	GetDebugInfo          *string                          `pulumi:"getDebugInfo"`
	Input                 interface{}                      `pulumi:"input"`
	LinkedServiceName     *LinkedServiceReferenceResponse  `pulumi:"linkedServiceName"`
	Mapper                interface{}                      `pulumi:"mapper"`
	Name                  string                           `pulumi:"name"`
	Output                interface{}                      `pulumi:"output"`
	Policy                *ActivityPolicyResponse          `pulumi:"policy"`
	Reducer               interface{}                      `pulumi:"reducer"`
	StorageLinkedServices []LinkedServiceReferenceResponse `pulumi:"storageLinkedServices"`
	Type                  string                           `pulumi:"type"`
	UserProperties        []UserPropertyResponse           `pulumi:"userProperties"`
}

type HdfsLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	UserName            interface{}                       `pulumi:"userName"`
}

type HdfsLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	UserName            interface{}                               `pulumi:"userName"`
}

type HdfsLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type HdfsLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type HdfsReadSettings struct {
	DeleteFilesAfterCompletion interface{}     `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{}     `pulumi:"disableMetricsCollection"`
	DistcpSettings             *DistcpSettings `pulumi:"distcpSettings"`
	EnablePartitionDiscovery   *bool           `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{}     `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{}     `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{}     `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{}     `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{}     `pulumi:"partitionRootPath"`
	Recursive                  interface{}     `pulumi:"recursive"`
	Type                       string          `pulumi:"type"`
	WildcardFileName           interface{}     `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{}     `pulumi:"wildcardFolderPath"`
}

type HdfsReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{}             `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{}             `pulumi:"disableMetricsCollection"`
	DistcpSettings             *DistcpSettingsResponse `pulumi:"distcpSettings"`
	EnablePartitionDiscovery   *bool                   `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{}             `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{}             `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{}             `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{}             `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{}             `pulumi:"partitionRootPath"`
	Recursive                  interface{}             `pulumi:"recursive"`
	Type                       string                  `pulumi:"type"`
	WildcardFileName           interface{}             `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{}             `pulumi:"wildcardFolderPath"`
}

type HdfsSource struct {
	DisableMetricsCollection interface{}     `pulumi:"disableMetricsCollection"`
	DistcpSettings           *DistcpSettings `pulumi:"distcpSettings"`
	MaxConcurrentConnections interface{}     `pulumi:"maxConcurrentConnections"`
	Recursive                interface{}     `pulumi:"recursive"`
	SourceRetryCount         interface{}     `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}     `pulumi:"sourceRetryWait"`
	Type                     string          `pulumi:"type"`
}

type HdfsSourceResponse struct {
	DisableMetricsCollection interface{}             `pulumi:"disableMetricsCollection"`
	DistcpSettings           *DistcpSettingsResponse `pulumi:"distcpSettings"`
	MaxConcurrentConnections interface{}             `pulumi:"maxConcurrentConnections"`
	Recursive                interface{}             `pulumi:"recursive"`
	SourceRetryCount         interface{}             `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}             `pulumi:"sourceRetryWait"`
	Type                     string                  `pulumi:"type"`
}

type HiveLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	HttpPath                  interface{}                       `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	ServerType                *string                           `pulumi:"serverType"`
	ServiceDiscoveryMode      interface{}                       `pulumi:"serviceDiscoveryMode"`
	ThriftTransportProtocol   *string                           `pulumi:"thriftTransportProtocol"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseNativeQuery            interface{}                       `pulumi:"useNativeQuery"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
	ZooKeeperNameSpace        interface{}                       `pulumi:"zooKeeperNameSpace"`
}

type HiveLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	HttpPath                  interface{}                               `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	ServerType                *string                                   `pulumi:"serverType"`
	ServiceDiscoveryMode      interface{}                               `pulumi:"serviceDiscoveryMode"`
	ThriftTransportProtocol   *string                                   `pulumi:"thriftTransportProtocol"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseNativeQuery            interface{}                               `pulumi:"useNativeQuery"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
	ZooKeeperNameSpace        interface{}                               `pulumi:"zooKeeperNameSpace"`
}

type HiveObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type HiveObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type HiveSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HiveSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HttpDataset struct {
	AdditionalHeaders interface{}                       `pulumi:"additionalHeaders"`
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	Format            interface{}                       `pulumi:"format"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	RelativeUrl       interface{}                       `pulumi:"relativeUrl"`
	RequestBody       interface{}                       `pulumi:"requestBody"`
	RequestMethod     interface{}                       `pulumi:"requestMethod"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type HttpDatasetResponse struct {
	AdditionalHeaders interface{}                               `pulumi:"additionalHeaders"`
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	Format            interface{}                               `pulumi:"format"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RelativeUrl       interface{}                               `pulumi:"relativeUrl"`
	RequestBody       interface{}                               `pulumi:"requestBody"`
	RequestMethod     interface{}                               `pulumi:"requestMethod"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type HttpLinkedService struct {
	Annotations                       []interface{}                     `pulumi:"annotations"`
	AuthHeaders                       interface{}                       `pulumi:"authHeaders"`
	AuthenticationType                *string                           `pulumi:"authenticationType"`
	CertThumbprint                    interface{}                       `pulumi:"certThumbprint"`
	ConnectVia                        *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description                       *string                           `pulumi:"description"`
	EmbeddedCertData                  interface{}                       `pulumi:"embeddedCertData"`
	EnableServerCertificateValidation interface{}                       `pulumi:"enableServerCertificateValidation"`
	EncryptedCredential               interface{}                       `pulumi:"encryptedCredential"`
	Parameters                        map[string]ParameterSpecification `pulumi:"parameters"`
	Password                          interface{}                       `pulumi:"password"`
	Type                              string                            `pulumi:"type"`
	Url                               interface{}                       `pulumi:"url"`
	UserName                          interface{}                       `pulumi:"userName"`
}

type HttpLinkedServiceResponse struct {
	Annotations                       []interface{}                             `pulumi:"annotations"`
	AuthHeaders                       interface{}                               `pulumi:"authHeaders"`
	AuthenticationType                *string                                   `pulumi:"authenticationType"`
	CertThumbprint                    interface{}                               `pulumi:"certThumbprint"`
	ConnectVia                        *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description                       *string                                   `pulumi:"description"`
	EmbeddedCertData                  interface{}                               `pulumi:"embeddedCertData"`
	EnableServerCertificateValidation interface{}                               `pulumi:"enableServerCertificateValidation"`
	EncryptedCredential               interface{}                               `pulumi:"encryptedCredential"`
	Parameters                        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                          interface{}                               `pulumi:"password"`
	Type                              string                                    `pulumi:"type"`
	Url                               interface{}                               `pulumi:"url"`
	UserName                          interface{}                               `pulumi:"userName"`
}

type HttpReadSettings struct {
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery *bool       `pulumi:"enablePartitionDiscovery"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PartitionRootPath        interface{} `pulumi:"partitionRootPath"`
	RequestBody              interface{} `pulumi:"requestBody"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	RequestTimeout           interface{} `pulumi:"requestTimeout"`
	Type                     string      `pulumi:"type"`
}

type HttpReadSettingsResponse struct {
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery *bool       `pulumi:"enablePartitionDiscovery"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PartitionRootPath        interface{} `pulumi:"partitionRootPath"`
	RequestBody              interface{} `pulumi:"requestBody"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	RequestTimeout           interface{} `pulumi:"requestTimeout"`
	Type                     string      `pulumi:"type"`
}

type HttpServerLocation struct {
	FileName    interface{} `pulumi:"fileName"`
	FolderPath  interface{} `pulumi:"folderPath"`
	RelativeUrl interface{} `pulumi:"relativeUrl"`
	Type        string      `pulumi:"type"`
}

type HttpServerLocationResponse struct {
	FileName    interface{} `pulumi:"fileName"`
	FolderPath  interface{} `pulumi:"folderPath"`
	RelativeUrl interface{} `pulumi:"relativeUrl"`
	Type        string      `pulumi:"type"`
}

type HttpSource struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HttpSourceResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HubspotLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	RefreshToken          interface{}                       `pulumi:"refreshToken"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type HubspotLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RefreshToken          interface{}                               `pulumi:"refreshToken"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type HubspotObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type HubspotObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type HubspotSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type HubspotSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type IfConditionActivity struct {
	DependsOn         []ActivityDependency `pulumi:"dependsOn"`
	Description       *string              `pulumi:"description"`
	Expression        Expression           `pulumi:"expression"`
	IfFalseActivities []interface{}        `pulumi:"ifFalseActivities"`
	IfTrueActivities  []interface{}        `pulumi:"ifTrueActivities"`
	Name              string               `pulumi:"name"`
	Type              string               `pulumi:"type"`
	UserProperties    []UserProperty       `pulumi:"userProperties"`
}

type IfConditionActivityResponse struct {
	DependsOn         []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description       *string                      `pulumi:"description"`
	Expression        ExpressionResponse           `pulumi:"expression"`
	IfFalseActivities []interface{}                `pulumi:"ifFalseActivities"`
	IfTrueActivities  []interface{}                `pulumi:"ifTrueActivities"`
	Name              string                       `pulumi:"name"`
	Type              string                       `pulumi:"type"`
	UserProperties    []UserPropertyResponse       `pulumi:"userProperties"`
}

type ImpalaLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
}

type ImpalaLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
}

type ImpalaObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type ImpalaObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type ImpalaSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ImpalaSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type InformixLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Credential          interface{}                       `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type InformixLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Credential          interface{}                               `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type InformixSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type InformixSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type InformixSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type InformixSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type InformixTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type InformixTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type IntegrationRuntimeComputeProperties struct {
	DataFlowProperties           *IntegrationRuntimeDataFlowProperties `pulumi:"dataFlowProperties"`
	Location                     *string                               `pulumi:"location"`
	MaxParallelExecutionsPerNode *int                                  `pulumi:"maxParallelExecutionsPerNode"`
	NodeSize                     *string                               `pulumi:"nodeSize"`
	NumberOfNodes                *int                                  `pulumi:"numberOfNodes"`
	VNetProperties               *IntegrationRuntimeVNetProperties     `pulumi:"vNetProperties"`
}

type IntegrationRuntimeComputePropertiesResponse struct {
	DataFlowProperties           *IntegrationRuntimeDataFlowPropertiesResponse `pulumi:"dataFlowProperties"`
	Location                     *string                                       `pulumi:"location"`
	MaxParallelExecutionsPerNode *int                                          `pulumi:"maxParallelExecutionsPerNode"`
	NodeSize                     *string                                       `pulumi:"nodeSize"`
	NumberOfNodes                *int                                          `pulumi:"numberOfNodes"`
	VNetProperties               *IntegrationRuntimeVNetPropertiesResponse     `pulumi:"vNetProperties"`
}

type IntegrationRuntimeCustomSetupScriptProperties struct {
	BlobContainerUri *string       `pulumi:"blobContainerUri"`
	SasToken         *SecureString `pulumi:"sasToken"`
}

type IntegrationRuntimeCustomSetupScriptPropertiesResponse struct {
	BlobContainerUri *string               `pulumi:"blobContainerUri"`
	SasToken         *SecureStringResponse `pulumi:"sasToken"`
}

type IntegrationRuntimeCustomerVirtualNetwork struct {
	SubnetId *string `pulumi:"subnetId"`
}

type IntegrationRuntimeCustomerVirtualNetworkResponse struct {
	SubnetId *string `pulumi:"subnetId"`
}

type IntegrationRuntimeDataFlowProperties struct {
	Cleanup     *bool   `pulumi:"cleanup"`
	ComputeType *string `pulumi:"computeType"`
	CoreCount   *int    `pulumi:"coreCount"`
	TimeToLive  *int    `pulumi:"timeToLive"`
}

type IntegrationRuntimeDataFlowPropertiesResponse struct {
	Cleanup     *bool   `pulumi:"cleanup"`
	ComputeType *string `pulumi:"computeType"`
	CoreCount   *int    `pulumi:"coreCount"`
	TimeToLive  *int    `pulumi:"timeToLive"`
}

type IntegrationRuntimeDataProxyProperties struct {
	ConnectVia           *EntityReference `pulumi:"connectVia"`
	Path                 *string          `pulumi:"path"`
	StagingLinkedService *EntityReference `pulumi:"stagingLinkedService"`
}

type IntegrationRuntimeDataProxyPropertiesResponse struct {
	ConnectVia           *EntityReferenceResponse `pulumi:"connectVia"`
	Path                 *string                  `pulumi:"path"`
	StagingLinkedService *EntityReferenceResponse `pulumi:"stagingLinkedService"`
}

type IntegrationRuntimeReference struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type IntegrationRuntimeReferenceResponse struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type IntegrationRuntimeSsisCatalogInfo struct {
	CatalogAdminPassword  *SecureString `pulumi:"catalogAdminPassword"`
	CatalogAdminUserName  *string       `pulumi:"catalogAdminUserName"`
	CatalogPricingTier    *string       `pulumi:"catalogPricingTier"`
	CatalogServerEndpoint *string       `pulumi:"catalogServerEndpoint"`
	DualStandbyPairName   *string       `pulumi:"dualStandbyPairName"`
}

type IntegrationRuntimeSsisCatalogInfoResponse struct {
	CatalogAdminPassword  *SecureStringResponse `pulumi:"catalogAdminPassword"`
	CatalogAdminUserName  *string               `pulumi:"catalogAdminUserName"`
	CatalogPricingTier    *string               `pulumi:"catalogPricingTier"`
	CatalogServerEndpoint *string               `pulumi:"catalogServerEndpoint"`
	DualStandbyPairName   *string               `pulumi:"dualStandbyPairName"`
}

type IntegrationRuntimeSsisProperties struct {
	CatalogInfo                  *IntegrationRuntimeSsisCatalogInfo             `pulumi:"catalogInfo"`
	Credential                   *CredentialReference                           `pulumi:"credential"`
	CustomSetupScriptProperties  *IntegrationRuntimeCustomSetupScriptProperties `pulumi:"customSetupScriptProperties"`
	DataProxyProperties          *IntegrationRuntimeDataProxyProperties         `pulumi:"dataProxyProperties"`
	Edition                      *string                                        `pulumi:"edition"`
	ExpressCustomSetupProperties []interface{}                                  `pulumi:"expressCustomSetupProperties"`
	LicenseType                  *string                                        `pulumi:"licenseType"`
	PackageStores                []PackageStore                                 `pulumi:"packageStores"`
}

type IntegrationRuntimeSsisPropertiesResponse struct {
	CatalogInfo                  *IntegrationRuntimeSsisCatalogInfoResponse             `pulumi:"catalogInfo"`
	Credential                   *CredentialReferenceResponse                           `pulumi:"credential"`
	CustomSetupScriptProperties  *IntegrationRuntimeCustomSetupScriptPropertiesResponse `pulumi:"customSetupScriptProperties"`
	DataProxyProperties          *IntegrationRuntimeDataProxyPropertiesResponse         `pulumi:"dataProxyProperties"`
	Edition                      *string                                                `pulumi:"edition"`
	ExpressCustomSetupProperties []interface{}                                          `pulumi:"expressCustomSetupProperties"`
	LicenseType                  *string                                                `pulumi:"licenseType"`
	PackageStores                []PackageStoreResponse                                 `pulumi:"packageStores"`
}

type IntegrationRuntimeVNetProperties struct {
	PublicIPs []string `pulumi:"publicIPs"`
	Subnet    *string  `pulumi:"subnet"`
	SubnetId  *string  `pulumi:"subnetId"`
	VNetId    *string  `pulumi:"vNetId"`
}

type IntegrationRuntimeVNetPropertiesResponse struct {
	PublicIPs []string `pulumi:"publicIPs"`
	Subnet    *string  `pulumi:"subnet"`
	SubnetId  *string  `pulumi:"subnetId"`
	VNetId    *string  `pulumi:"vNetId"`
}

type JiraLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Port                  interface{}                       `pulumi:"port"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
	Username              interface{}                       `pulumi:"username"`
}

type JiraLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Port                  interface{}                               `pulumi:"port"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
	Username              interface{}                               `pulumi:"username"`
}

type JiraObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type JiraObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type JiraSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type JiraSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type JsonDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Compression       *DatasetCompression               `pulumi:"compression"`
	Description       *string                           `pulumi:"description"`
	EncodingName      interface{}                       `pulumi:"encodingName"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location          interface{}                       `pulumi:"location"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type JsonDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Compression       *DatasetCompressionResponse               `pulumi:"compression"`
	Description       *string                                   `pulumi:"description"`
	EncodingName      interface{}                               `pulumi:"encodingName"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location          interface{}                               `pulumi:"location"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type JsonFormat struct {
	Deserializer       interface{} `pulumi:"deserializer"`
	EncodingName       interface{} `pulumi:"encodingName"`
	FilePattern        interface{} `pulumi:"filePattern"`
	JsonNodeReference  interface{} `pulumi:"jsonNodeReference"`
	JsonPathDefinition interface{} `pulumi:"jsonPathDefinition"`
	NestingSeparator   interface{} `pulumi:"nestingSeparator"`
	Serializer         interface{} `pulumi:"serializer"`
	Type               string      `pulumi:"type"`
}

type JsonFormatResponse struct {
	Deserializer       interface{} `pulumi:"deserializer"`
	EncodingName       interface{} `pulumi:"encodingName"`
	FilePattern        interface{} `pulumi:"filePattern"`
	JsonNodeReference  interface{} `pulumi:"jsonNodeReference"`
	JsonPathDefinition interface{} `pulumi:"jsonPathDefinition"`
	NestingSeparator   interface{} `pulumi:"nestingSeparator"`
	Serializer         interface{} `pulumi:"serializer"`
	Type               string      `pulumi:"type"`
}

type JsonReadSettings struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	Type                  string      `pulumi:"type"`
}

type JsonReadSettingsResponse struct {
	CompressionProperties interface{} `pulumi:"compressionProperties"`
	Type                  string      `pulumi:"type"`
}

type JsonSink struct {
	DisableMetricsCollection interface{}        `pulumi:"disableMetricsCollection"`
	FormatSettings           *JsonWriteSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}        `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}        `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}        `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}        `pulumi:"storeSettings"`
	Type                     string             `pulumi:"type"`
	WriteBatchSize           interface{}        `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}        `pulumi:"writeBatchTimeout"`
}

type JsonSinkResponse struct {
	DisableMetricsCollection interface{}                `pulumi:"disableMetricsCollection"`
	FormatSettings           *JsonWriteSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}                `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}                `pulumi:"storeSettings"`
	Type                     string                     `pulumi:"type"`
	WriteBatchSize           interface{}                `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                `pulumi:"writeBatchTimeout"`
}

type JsonSource struct {
	AdditionalColumns        interface{}       `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}       `pulumi:"disableMetricsCollection"`
	FormatSettings           *JsonReadSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}       `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}       `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}       `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}       `pulumi:"storeSettings"`
	Type                     string            `pulumi:"type"`
}

type JsonSourceResponse struct {
	AdditionalColumns        interface{}               `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}               `pulumi:"disableMetricsCollection"`
	FormatSettings           *JsonReadSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}               `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{}               `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}               `pulumi:"sourceRetryWait"`
	StoreSettings            interface{}               `pulumi:"storeSettings"`
	Type                     string                    `pulumi:"type"`
}

type JsonWriteSettings struct {
	FilePattern interface{} `pulumi:"filePattern"`
	Type        string      `pulumi:"type"`
}

type JsonWriteSettingsResponse struct {
	FilePattern interface{} `pulumi:"filePattern"`
	Type        string      `pulumi:"type"`
}

type LinkedIntegrationRuntimeKeyAuthorization struct {
	AuthorizationType string       `pulumi:"authorizationType"`
	Key               SecureString `pulumi:"key"`
}

type LinkedIntegrationRuntimeKeyAuthorizationResponse struct {
	AuthorizationType string               `pulumi:"authorizationType"`
	Key               SecureStringResponse `pulumi:"key"`
}

type LinkedIntegrationRuntimeRbacAuthorization struct {
	AuthorizationType string               `pulumi:"authorizationType"`
	Credential        *CredentialReference `pulumi:"credential"`
	ResourceId        string               `pulumi:"resourceId"`
}

type LinkedIntegrationRuntimeRbacAuthorizationResponse struct {
	AuthorizationType string                       `pulumi:"authorizationType"`
	Credential        *CredentialReferenceResponse `pulumi:"credential"`
	ResourceId        string                       `pulumi:"resourceId"`
}

type LinkedIntegrationRuntimeResponse struct {
	CreateTime          string `pulumi:"createTime"`
	DataFactoryLocation string `pulumi:"dataFactoryLocation"`
	DataFactoryName     string `pulumi:"dataFactoryName"`
	Name                string `pulumi:"name"`
	SubscriptionId      string `pulumi:"subscriptionId"`
}

type LinkedServiceReference struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type LinkedServiceReferenceResponse struct {
	Parameters    map[string]interface{} `pulumi:"parameters"`
	ReferenceName string                 `pulumi:"referenceName"`
	Type          string                 `pulumi:"type"`
}

type LogLocationSettings struct {
	LinkedServiceName LinkedServiceReference `pulumi:"linkedServiceName"`
	Path              interface{}            `pulumi:"path"`
}

type LogLocationSettingsResponse struct {
	LinkedServiceName LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Path              interface{}                    `pulumi:"path"`
}

type LogSettings struct {
	CopyActivityLogSettings *CopyActivityLogSettings `pulumi:"copyActivityLogSettings"`
	EnableCopyActivityLog   interface{}              `pulumi:"enableCopyActivityLog"`
	LogLocationSettings     LogLocationSettings      `pulumi:"logLocationSettings"`
}

type LogSettingsResponse struct {
	CopyActivityLogSettings *CopyActivityLogSettingsResponse `pulumi:"copyActivityLogSettings"`
	EnableCopyActivityLog   interface{}                      `pulumi:"enableCopyActivityLog"`
	LogLocationSettings     LogLocationSettingsResponse      `pulumi:"logLocationSettings"`
}

type LogStorageSettings struct {
	EnableReliableLogging interface{}            `pulumi:"enableReliableLogging"`
	LinkedServiceName     LinkedServiceReference `pulumi:"linkedServiceName"`
	LogLevel              interface{}            `pulumi:"logLevel"`
	Path                  interface{}            `pulumi:"path"`
}

type LogStorageSettingsResponse struct {
	EnableReliableLogging interface{}                    `pulumi:"enableReliableLogging"`
	LinkedServiceName     LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	LogLevel              interface{}                    `pulumi:"logLevel"`
	Path                  interface{}                    `pulumi:"path"`
}

type LookupActivity struct {
	Dataset           DatasetReference        `pulumi:"dataset"`
	DependsOn         []ActivityDependency    `pulumi:"dependsOn"`
	Description       *string                 `pulumi:"description"`
	FirstRowOnly      interface{}             `pulumi:"firstRowOnly"`
	LinkedServiceName *LinkedServiceReference `pulumi:"linkedServiceName"`
	Name              string                  `pulumi:"name"`
	Policy            *ActivityPolicy         `pulumi:"policy"`
	Source            interface{}             `pulumi:"source"`
	Type              string                  `pulumi:"type"`
	UserProperties    []UserProperty          `pulumi:"userProperties"`
}

type LookupActivityResponse struct {
	Dataset           DatasetReferenceResponse        `pulumi:"dataset"`
	DependsOn         []ActivityDependencyResponse    `pulumi:"dependsOn"`
	Description       *string                         `pulumi:"description"`
	FirstRowOnly      interface{}                     `pulumi:"firstRowOnly"`
	LinkedServiceName *LinkedServiceReferenceResponse `pulumi:"linkedServiceName"`
	Name              string                          `pulumi:"name"`
	Policy            *ActivityPolicyResponse         `pulumi:"policy"`
	Source            interface{}                     `pulumi:"source"`
	Type              string                          `pulumi:"type"`
	UserProperties    []UserPropertyResponse          `pulumi:"userProperties"`
}

type MagentoLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type MagentoLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type MagentoObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type MagentoObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type MagentoSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MagentoSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ManagedIdentityCredential struct {
	Annotations []interface{} `pulumi:"annotations"`
	Description *string       `pulumi:"description"`
	ResourceId  *string       `pulumi:"resourceId"`
	Type        string        `pulumi:"type"`
}





type ManagedIdentityCredentialInput interface {
	pulumi.Input

	ToManagedIdentityCredentialOutput() ManagedIdentityCredentialOutput
	ToManagedIdentityCredentialOutputWithContext(context.Context) ManagedIdentityCredentialOutput
}

type ManagedIdentityCredentialArgs struct {
	Annotations pulumi.ArrayInput     `pulumi:"annotations"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	ResourceId  pulumi.StringPtrInput `pulumi:"resourceId"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (ManagedIdentityCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityCredential)(nil)).Elem()
}

func (i ManagedIdentityCredentialArgs) ToManagedIdentityCredentialOutput() ManagedIdentityCredentialOutput {
	return i.ToManagedIdentityCredentialOutputWithContext(context.Background())
}

func (i ManagedIdentityCredentialArgs) ToManagedIdentityCredentialOutputWithContext(ctx context.Context) ManagedIdentityCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityCredentialOutput)
}

type ManagedIdentityCredentialOutput struct{ *pulumi.OutputState }

func (ManagedIdentityCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityCredential)(nil)).Elem()
}

func (o ManagedIdentityCredentialOutput) ToManagedIdentityCredentialOutput() ManagedIdentityCredentialOutput {
	return o
}

func (o ManagedIdentityCredentialOutput) ToManagedIdentityCredentialOutputWithContext(ctx context.Context) ManagedIdentityCredentialOutput {
	return o
}

func (o ManagedIdentityCredentialOutput) Annotations() pulumi.ArrayOutput {
	return o.ApplyT(func(v ManagedIdentityCredential) []interface{} { return v.Annotations }).(pulumi.ArrayOutput)
}

func (o ManagedIdentityCredentialOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityCredential) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityCredentialOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityCredential) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityCredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityCredential) string { return v.Type }).(pulumi.StringOutput)
}

type ManagedIdentityCredentialResponse struct {
	Annotations []interface{} `pulumi:"annotations"`
	Description *string       `pulumi:"description"`
	ResourceId  *string       `pulumi:"resourceId"`
	Type        string        `pulumi:"type"`
}

type ManagedIdentityCredentialResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityCredentialResponse)(nil)).Elem()
}

func (o ManagedIdentityCredentialResponseOutput) ToManagedIdentityCredentialResponseOutput() ManagedIdentityCredentialResponseOutput {
	return o
}

func (o ManagedIdentityCredentialResponseOutput) ToManagedIdentityCredentialResponseOutputWithContext(ctx context.Context) ManagedIdentityCredentialResponseOutput {
	return o
}

func (o ManagedIdentityCredentialResponseOutput) Annotations() pulumi.ArrayOutput {
	return o.ApplyT(func(v ManagedIdentityCredentialResponse) []interface{} { return v.Annotations }).(pulumi.ArrayOutput)
}

func (o ManagedIdentityCredentialResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityCredentialResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityCredentialResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityCredentialResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityCredentialResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityCredentialResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagedIntegrationRuntime struct {
	ComputeProperties      *IntegrationRuntimeComputeProperties      `pulumi:"computeProperties"`
	CustomerVirtualNetwork *IntegrationRuntimeCustomerVirtualNetwork `pulumi:"customerVirtualNetwork"`
	Description            *string                                   `pulumi:"description"`
	ManagedVirtualNetwork  *ManagedVirtualNetworkReference           `pulumi:"managedVirtualNetwork"`
	SsisProperties         *IntegrationRuntimeSsisProperties         `pulumi:"ssisProperties"`
	Type                   string                                    `pulumi:"type"`
}

type ManagedIntegrationRuntimeErrorResponse struct {
	Code       string   `pulumi:"code"`
	Message    string   `pulumi:"message"`
	Parameters []string `pulumi:"parameters"`
	Time       string   `pulumi:"time"`
}

type ManagedIntegrationRuntimeNodeResponse struct {
	Errors []ManagedIntegrationRuntimeErrorResponse `pulumi:"errors"`
	NodeId string                                   `pulumi:"nodeId"`
	Status string                                   `pulumi:"status"`
}

type ManagedIntegrationRuntimeOperationResultResponse struct {
	ActivityId string   `pulumi:"activityId"`
	ErrorCode  string   `pulumi:"errorCode"`
	Parameters []string `pulumi:"parameters"`
	Result     string   `pulumi:"result"`
	StartTime  string   `pulumi:"startTime"`
	Type       string   `pulumi:"type"`
}

type ManagedIntegrationRuntimeResponse struct {
	ComputeProperties      *IntegrationRuntimeComputePropertiesResponse      `pulumi:"computeProperties"`
	CustomerVirtualNetwork *IntegrationRuntimeCustomerVirtualNetworkResponse `pulumi:"customerVirtualNetwork"`
	Description            *string                                           `pulumi:"description"`
	ManagedVirtualNetwork  *ManagedVirtualNetworkReferenceResponse           `pulumi:"managedVirtualNetwork"`
	SsisProperties         *IntegrationRuntimeSsisPropertiesResponse         `pulumi:"ssisProperties"`
	State                  string                                            `pulumi:"state"`
	Type                   string                                            `pulumi:"type"`
}

type ManagedIntegrationRuntimeStatusResponse struct {
	CreateTime      string                                           `pulumi:"createTime"`
	DataFactoryName string                                           `pulumi:"dataFactoryName"`
	LastOperation   ManagedIntegrationRuntimeOperationResultResponse `pulumi:"lastOperation"`
	Nodes           []ManagedIntegrationRuntimeNodeResponse          `pulumi:"nodes"`
	OtherErrors     []ManagedIntegrationRuntimeErrorResponse         `pulumi:"otherErrors"`
	State           string                                           `pulumi:"state"`
	Type            string                                           `pulumi:"type"`
}

type ManagedPrivateEndpointType struct {
	Fqdns                 []string `pulumi:"fqdns"`
	GroupId               *string  `pulumi:"groupId"`
	PrivateLinkResourceId *string  `pulumi:"privateLinkResourceId"`
}





type ManagedPrivateEndpointTypeInput interface {
	pulumi.Input

	ToManagedPrivateEndpointTypeOutput() ManagedPrivateEndpointTypeOutput
	ToManagedPrivateEndpointTypeOutputWithContext(context.Context) ManagedPrivateEndpointTypeOutput
}

type ManagedPrivateEndpointTypeArgs struct {
	Fqdns                 pulumi.StringArrayInput `pulumi:"fqdns"`
	GroupId               pulumi.StringPtrInput   `pulumi:"groupId"`
	PrivateLinkResourceId pulumi.StringPtrInput   `pulumi:"privateLinkResourceId"`
}

func (ManagedPrivateEndpointTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpointType)(nil)).Elem()
}

func (i ManagedPrivateEndpointTypeArgs) ToManagedPrivateEndpointTypeOutput() ManagedPrivateEndpointTypeOutput {
	return i.ToManagedPrivateEndpointTypeOutputWithContext(context.Background())
}

func (i ManagedPrivateEndpointTypeArgs) ToManagedPrivateEndpointTypeOutputWithContext(ctx context.Context) ManagedPrivateEndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointTypeOutput)
}

type ManagedPrivateEndpointTypeOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpointType)(nil)).Elem()
}

func (o ManagedPrivateEndpointTypeOutput) ToManagedPrivateEndpointTypeOutput() ManagedPrivateEndpointTypeOutput {
	return o
}

func (o ManagedPrivateEndpointTypeOutput) ToManagedPrivateEndpointTypeOutputWithContext(ctx context.Context) ManagedPrivateEndpointTypeOutput {
	return o
}

func (o ManagedPrivateEndpointTypeOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointType) []string { return v.Fqdns }).(pulumi.StringArrayOutput)
}

func (o ManagedPrivateEndpointTypeOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointType) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o ManagedPrivateEndpointTypeOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointType) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

type ManagedPrivateEndpointResponse struct {
	ConnectionState       *ConnectionStatePropertiesResponse `pulumi:"connectionState"`
	Fqdns                 []string                           `pulumi:"fqdns"`
	GroupId               *string                            `pulumi:"groupId"`
	IsReserved            bool                               `pulumi:"isReserved"`
	PrivateLinkResourceId *string                            `pulumi:"privateLinkResourceId"`
	ProvisioningState     string                             `pulumi:"provisioningState"`
}

type ManagedPrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrivateEndpointResponse)(nil)).Elem()
}

func (o ManagedPrivateEndpointResponseOutput) ToManagedPrivateEndpointResponseOutput() ManagedPrivateEndpointResponseOutput {
	return o
}

func (o ManagedPrivateEndpointResponseOutput) ToManagedPrivateEndpointResponseOutputWithContext(ctx context.Context) ManagedPrivateEndpointResponseOutput {
	return o
}

func (o ManagedPrivateEndpointResponseOutput) ConnectionState() ConnectionStatePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) *ConnectionStatePropertiesResponse { return v.ConnectionState }).(ConnectionStatePropertiesResponsePtrOutput)
}

func (o ManagedPrivateEndpointResponseOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) []string { return v.Fqdns }).(pulumi.StringArrayOutput)
}

func (o ManagedPrivateEndpointResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o ManagedPrivateEndpointResponseOutput) IsReserved() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) bool { return v.IsReserved }).(pulumi.BoolOutput)
}

func (o ManagedPrivateEndpointResponseOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) *string { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

func (o ManagedPrivateEndpointResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedPrivateEndpointResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ManagedVirtualNetworkReference struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type ManagedVirtualNetworkReferenceResponse struct {
	ReferenceName string `pulumi:"referenceName"`
	Type          string `pulumi:"type"`
}

type MappingDataFlow struct {
	Annotations     []interface{}    `pulumi:"annotations"`
	Description     *string          `pulumi:"description"`
	Folder          *DataFlowFolder  `pulumi:"folder"`
	Script          *string          `pulumi:"script"`
	ScriptLines     []string         `pulumi:"scriptLines"`
	Sinks           []DataFlowSink   `pulumi:"sinks"`
	Sources         []DataFlowSource `pulumi:"sources"`
	Transformations []Transformation `pulumi:"transformations"`
	Type            string           `pulumi:"type"`
}

type MappingDataFlowResponse struct {
	Annotations     []interface{}            `pulumi:"annotations"`
	Description     *string                  `pulumi:"description"`
	Folder          *DataFlowResponseFolder  `pulumi:"folder"`
	Script          *string                  `pulumi:"script"`
	ScriptLines     []string                 `pulumi:"scriptLines"`
	Sinks           []DataFlowSinkResponse   `pulumi:"sinks"`
	Sources         []DataFlowSourceResponse `pulumi:"sources"`
	Transformations []TransformationResponse `pulumi:"transformations"`
	Type            string                   `pulumi:"type"`
}

type MariaDBLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReference     `pulumi:"pwd"`
	Type                string                            `pulumi:"type"`
}

type MariaDBLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReferenceResponse     `pulumi:"pwd"`
	Type                string                                    `pulumi:"type"`
}

type MariaDBSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MariaDBSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MariaDBTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type MariaDBTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type MarketoLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type MarketoLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type MarketoObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type MarketoObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type MarketoSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MarketoSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MetadataItem struct {
	Name  interface{} `pulumi:"name"`
	Value interface{} `pulumi:"value"`
}

type MetadataItemResponse struct {
	Name  interface{} `pulumi:"name"`
	Value interface{} `pulumi:"value"`
}

type MicrosoftAccessLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Credential          interface{}                       `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type MicrosoftAccessLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Credential          interface{}                               `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type MicrosoftAccessSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type MicrosoftAccessSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type MicrosoftAccessSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MicrosoftAccessSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MicrosoftAccessTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type MicrosoftAccessTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type MongoDbAtlasCollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Collection        interface{}                       `pulumi:"collection"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MongoDbAtlasCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Collection        interface{}                               `pulumi:"collection"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type MongoDbAtlasLinkedService struct {
	Annotations      []interface{}                     `pulumi:"annotations"`
	ConnectVia       *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString interface{}                       `pulumi:"connectionString"`
	Database         interface{}                       `pulumi:"database"`
	Description      *string                           `pulumi:"description"`
	Parameters       map[string]ParameterSpecification `pulumi:"parameters"`
	Type             string                            `pulumi:"type"`
}

type MongoDbAtlasLinkedServiceResponse struct {
	Annotations      []interface{}                             `pulumi:"annotations"`
	ConnectVia       *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString interface{}                               `pulumi:"connectionString"`
	Database         interface{}                               `pulumi:"database"`
	Description      *string                                   `pulumi:"description"`
	Parameters       map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type             string                                    `pulumi:"type"`
}

type MongoDbAtlasSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type MongoDbAtlasSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type MongoDbAtlasSource struct {
	AdditionalColumns        interface{}                     `pulumi:"additionalColumns"`
	BatchSize                interface{}                     `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsProperties `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                     `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                     `pulumi:"filter"`
	MaxConcurrentConnections interface{}                     `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                     `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                     `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                     `pulumi:"sourceRetryWait"`
	Type                     string                          `pulumi:"type"`
}

type MongoDbAtlasSourceResponse struct {
	AdditionalColumns        interface{}                             `pulumi:"additionalColumns"`
	BatchSize                interface{}                             `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsPropertiesResponse `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                             `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                             `pulumi:"filter"`
	MaxConcurrentConnections interface{}                             `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                             `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                             `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                             `pulumi:"sourceRetryWait"`
	Type                     string                                  `pulumi:"type"`
}

type MongoDbCollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	CollectionName    interface{}                       `pulumi:"collectionName"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MongoDbCollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	CollectionName    interface{}                               `pulumi:"collectionName"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type MongoDbCursorMethodsProperties struct {
	Limit   interface{} `pulumi:"limit"`
	Project interface{} `pulumi:"project"`
	Skip    interface{} `pulumi:"skip"`
	Sort    interface{} `pulumi:"sort"`
}

type MongoDbCursorMethodsPropertiesResponse struct {
	Limit   interface{} `pulumi:"limit"`
	Project interface{} `pulumi:"project"`
	Skip    interface{} `pulumi:"skip"`
	Sort    interface{} `pulumi:"sort"`
}

type MongoDbLinkedService struct {
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthSource                interface{}                       `pulumi:"authSource"`
	AuthenticationType        *string                           `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	DatabaseName              interface{}                       `pulumi:"databaseName"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	Server                    interface{}                       `pulumi:"server"`
	Type                      string                            `pulumi:"type"`
	Username                  interface{}                       `pulumi:"username"`
}

type MongoDbLinkedServiceResponse struct {
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthSource                interface{}                               `pulumi:"authSource"`
	AuthenticationType        *string                                   `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	DatabaseName              interface{}                               `pulumi:"databaseName"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	Server                    interface{}                               `pulumi:"server"`
	Type                      string                                    `pulumi:"type"`
	Username                  interface{}                               `pulumi:"username"`
}

type MongoDbSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MongoDbSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MongoDbV2CollectionDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Collection        interface{}                       `pulumi:"collection"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type MongoDbV2CollectionDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Collection        interface{}                               `pulumi:"collection"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type MongoDbV2LinkedService struct {
	Annotations      []interface{}                     `pulumi:"annotations"`
	ConnectVia       *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString interface{}                       `pulumi:"connectionString"`
	Database         interface{}                       `pulumi:"database"`
	Description      *string                           `pulumi:"description"`
	Parameters       map[string]ParameterSpecification `pulumi:"parameters"`
	Type             string                            `pulumi:"type"`
}

type MongoDbV2LinkedServiceResponse struct {
	Annotations      []interface{}                             `pulumi:"annotations"`
	ConnectVia       *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString interface{}                               `pulumi:"connectionString"`
	Database         interface{}                               `pulumi:"database"`
	Description      *string                                   `pulumi:"description"`
	Parameters       map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type             string                                    `pulumi:"type"`
}

type MongoDbV2Sink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type MongoDbV2SinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{} `pulumi:"writeBehavior"`
}

type MongoDbV2Source struct {
	AdditionalColumns        interface{}                     `pulumi:"additionalColumns"`
	BatchSize                interface{}                     `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsProperties `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                     `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                     `pulumi:"filter"`
	MaxConcurrentConnections interface{}                     `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                     `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                     `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                     `pulumi:"sourceRetryWait"`
	Type                     string                          `pulumi:"type"`
}

type MongoDbV2SourceResponse struct {
	AdditionalColumns        interface{}                             `pulumi:"additionalColumns"`
	BatchSize                interface{}                             `pulumi:"batchSize"`
	CursorMethods            *MongoDbCursorMethodsPropertiesResponse `pulumi:"cursorMethods"`
	DisableMetricsCollection interface{}                             `pulumi:"disableMetricsCollection"`
	Filter                   interface{}                             `pulumi:"filter"`
	MaxConcurrentConnections interface{}                             `pulumi:"maxConcurrentConnections"`
	QueryTimeout             interface{}                             `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                             `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                             `pulumi:"sourceRetryWait"`
	Type                     string                                  `pulumi:"type"`
}

type MultiplePipelineTrigger struct {
	Annotations []interface{}              `pulumi:"annotations"`
	Description *string                    `pulumi:"description"`
	Pipelines   []TriggerPipelineReference `pulumi:"pipelines"`
	Type        string                     `pulumi:"type"`
}

type MultiplePipelineTriggerResponse struct {
	Annotations  []interface{}                      `pulumi:"annotations"`
	Description  *string                            `pulumi:"description"`
	Pipelines    []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	RuntimeState string                             `pulumi:"runtimeState"`
	Type         string                             `pulumi:"type"`
}

type MySqlLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type MySqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type MySqlSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MySqlSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type MySqlTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type MySqlTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type NetezzaLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReference     `pulumi:"pwd"`
	Type                string                            `pulumi:"type"`
}

type NetezzaLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Pwd                 *AzureKeyVaultSecretReferenceResponse     `pulumi:"pwd"`
	Type                string                                    `pulumi:"type"`
}

type NetezzaPartitionSettings struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type NetezzaPartitionSettingsResponse struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type NetezzaSource struct {
	AdditionalColumns        interface{}               `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}               `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}               `pulumi:"maxConcurrentConnections"`
	PartitionOption          interface{}               `pulumi:"partitionOption"`
	PartitionSettings        *NetezzaPartitionSettings `pulumi:"partitionSettings"`
	Query                    interface{}               `pulumi:"query"`
	QueryTimeout             interface{}               `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}               `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}               `pulumi:"sourceRetryWait"`
	Type                     string                    `pulumi:"type"`
}

type NetezzaSourceResponse struct {
	AdditionalColumns        interface{}                       `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                       `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                       `pulumi:"maxConcurrentConnections"`
	PartitionOption          interface{}                       `pulumi:"partitionOption"`
	PartitionSettings        *NetezzaPartitionSettingsResponse `pulumi:"partitionSettings"`
	Query                    interface{}                       `pulumi:"query"`
	QueryTimeout             interface{}                       `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                       `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                       `pulumi:"sourceRetryWait"`
	Type                     string                            `pulumi:"type"`
}

type NetezzaTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type NetezzaTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type NotebookParameter struct {
	Type  *string     `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}

type NotebookParameterResponse struct {
	Type  *string     `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}

type ODataLinkedService struct {
	AadResourceId                        interface{}                       `pulumi:"aadResourceId"`
	AadServicePrincipalCredentialType    *string                           `pulumi:"aadServicePrincipalCredentialType"`
	Annotations                          []interface{}                     `pulumi:"annotations"`
	AuthHeaders                          interface{}                       `pulumi:"authHeaders"`
	AuthenticationType                   *string                           `pulumi:"authenticationType"`
	AzureCloudType                       interface{}                       `pulumi:"azureCloudType"`
	ConnectVia                           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description                          *string                           `pulumi:"description"`
	EncryptedCredential                  interface{}                       `pulumi:"encryptedCredential"`
	Parameters                           map[string]ParameterSpecification `pulumi:"parameters"`
	Password                             interface{}                       `pulumi:"password"`
	ServicePrincipalEmbeddedCert         interface{}                       `pulumi:"servicePrincipalEmbeddedCert"`
	ServicePrincipalEmbeddedCertPassword interface{}                       `pulumi:"servicePrincipalEmbeddedCertPassword"`
	ServicePrincipalId                   interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey                  interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant                               interface{}                       `pulumi:"tenant"`
	Type                                 string                            `pulumi:"type"`
	Url                                  interface{}                       `pulumi:"url"`
	UserName                             interface{}                       `pulumi:"userName"`
}

type ODataLinkedServiceResponse struct {
	AadResourceId                        interface{}                               `pulumi:"aadResourceId"`
	AadServicePrincipalCredentialType    *string                                   `pulumi:"aadServicePrincipalCredentialType"`
	Annotations                          []interface{}                             `pulumi:"annotations"`
	AuthHeaders                          interface{}                               `pulumi:"authHeaders"`
	AuthenticationType                   *string                                   `pulumi:"authenticationType"`
	AzureCloudType                       interface{}                               `pulumi:"azureCloudType"`
	ConnectVia                           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description                          *string                                   `pulumi:"description"`
	EncryptedCredential                  interface{}                               `pulumi:"encryptedCredential"`
	Parameters                           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                             interface{}                               `pulumi:"password"`
	ServicePrincipalEmbeddedCert         interface{}                               `pulumi:"servicePrincipalEmbeddedCert"`
	ServicePrincipalEmbeddedCertPassword interface{}                               `pulumi:"servicePrincipalEmbeddedCertPassword"`
	ServicePrincipalId                   interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey                  interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant                               interface{}                               `pulumi:"tenant"`
	Type                                 string                                    `pulumi:"type"`
	Url                                  interface{}                               `pulumi:"url"`
	UserName                             interface{}                               `pulumi:"userName"`
}

type ODataResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ODataResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ODataSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ODataSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OdbcLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  interface{}                       `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Credential          interface{}                       `pulumi:"credential"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type OdbcLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  interface{}                               `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Credential          interface{}                               `pulumi:"credential"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type OdbcSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type OdbcSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type OdbcSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OdbcSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OdbcTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type OdbcTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type Office365Dataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Predicate         interface{}                       `pulumi:"predicate"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type Office365DatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Predicate         interface{}                               `pulumi:"predicate"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type Office365LinkedService struct {
	Annotations              []interface{}                     `pulumi:"annotations"`
	ConnectVia               *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description              *string                           `pulumi:"description"`
	EncryptedCredential      interface{}                       `pulumi:"encryptedCredential"`
	Office365TenantId        interface{}                       `pulumi:"office365TenantId"`
	Parameters               map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId       interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey      interface{}                       `pulumi:"servicePrincipalKey"`
	ServicePrincipalTenantId interface{}                       `pulumi:"servicePrincipalTenantId"`
	Type                     string                            `pulumi:"type"`
}

type Office365LinkedServiceResponse struct {
	Annotations              []interface{}                             `pulumi:"annotations"`
	ConnectVia               *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description              *string                                   `pulumi:"description"`
	EncryptedCredential      interface{}                               `pulumi:"encryptedCredential"`
	Office365TenantId        interface{}                               `pulumi:"office365TenantId"`
	Parameters               map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId       interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey      interface{}                               `pulumi:"servicePrincipalKey"`
	ServicePrincipalTenantId interface{}                               `pulumi:"servicePrincipalTenantId"`
	Type                     string                                    `pulumi:"type"`
}

type Office365Source struct {
	AllowedGroups            interface{} `pulumi:"allowedGroups"`
	DateFilterColumn         interface{} `pulumi:"dateFilterColumn"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	EndTime                  interface{} `pulumi:"endTime"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	OutputColumns            interface{} `pulumi:"outputColumns"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StartTime                interface{} `pulumi:"startTime"`
	Type                     string      `pulumi:"type"`
	UserScopeFilterUri       interface{} `pulumi:"userScopeFilterUri"`
}

type Office365SourceResponse struct {
	AllowedGroups            interface{} `pulumi:"allowedGroups"`
	DateFilterColumn         interface{} `pulumi:"dateFilterColumn"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	EndTime                  interface{} `pulumi:"endTime"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	OutputColumns            interface{} `pulumi:"outputColumns"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StartTime                interface{} `pulumi:"startTime"`
	Type                     string      `pulumi:"type"`
	UserScopeFilterUri       interface{} `pulumi:"userScopeFilterUri"`
}

type OracleCloudStorageLinkedService struct {
	AccessKeyId         interface{}                       `pulumi:"accessKeyId"`
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	SecretAccessKey     interface{}                       `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                       `pulumi:"serviceUrl"`
	Type                string                            `pulumi:"type"`
}

type OracleCloudStorageLinkedServiceResponse struct {
	AccessKeyId         interface{}                               `pulumi:"accessKeyId"`
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	SecretAccessKey     interface{}                               `pulumi:"secretAccessKey"`
	ServiceUrl          interface{}                               `pulumi:"serviceUrl"`
	Type                string                                    `pulumi:"type"`
}

type OracleCloudStorageLocation struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type OracleCloudStorageLocationResponse struct {
	BucketName interface{} `pulumi:"bucketName"`
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
	Version    interface{} `pulumi:"version"`
}

type OracleCloudStorageReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type OracleCloudStorageReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Prefix                     interface{} `pulumi:"prefix"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type OracleLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type OracleLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type OraclePartitionSettings struct {
	PartitionColumnName interface{}   `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{}   `pulumi:"partitionLowerBound"`
	PartitionNames      []interface{} `pulumi:"partitionNames"`
	PartitionUpperBound interface{}   `pulumi:"partitionUpperBound"`
}

type OraclePartitionSettingsResponse struct {
	PartitionColumnName interface{}   `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{}   `pulumi:"partitionLowerBound"`
	PartitionNames      []interface{} `pulumi:"partitionNames"`
	PartitionUpperBound interface{}   `pulumi:"partitionUpperBound"`
}

type OracleServiceCloudLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
	Username              interface{}                       `pulumi:"username"`
}

type OracleServiceCloudLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
	Username              interface{}                               `pulumi:"username"`
}

type OracleServiceCloudObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type OracleServiceCloudObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type OracleServiceCloudSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OracleServiceCloudSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type OracleSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type OracleSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{} `pulumi:"preCopyScript"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type OracleSource struct {
	AdditionalColumns        interface{}              `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}              `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}              `pulumi:"maxConcurrentConnections"`
	OracleReaderQuery        interface{}              `pulumi:"oracleReaderQuery"`
	PartitionOption          interface{}              `pulumi:"partitionOption"`
	PartitionSettings        *OraclePartitionSettings `pulumi:"partitionSettings"`
	QueryTimeout             interface{}              `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}              `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}              `pulumi:"sourceRetryWait"`
	Type                     string                   `pulumi:"type"`
}

type OracleSourceResponse struct {
	AdditionalColumns        interface{}                      `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                      `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                      `pulumi:"maxConcurrentConnections"`
	OracleReaderQuery        interface{}                      `pulumi:"oracleReaderQuery"`
	PartitionOption          interface{}                      `pulumi:"partitionOption"`
	PartitionSettings        *OraclePartitionSettingsResponse `pulumi:"partitionSettings"`
	QueryTimeout             interface{}                      `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                      `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                      `pulumi:"sourceRetryWait"`
	Type                     string                           `pulumi:"type"`
}

type OracleTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type OracleTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type OrcDataset struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	Description         *string                           `pulumi:"description"`
	Folder              *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName   LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location            interface{}                       `pulumi:"location"`
	OrcCompressionCodec interface{}                       `pulumi:"orcCompressionCodec"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Schema              interface{}                       `pulumi:"schema"`
	Structure           interface{}                       `pulumi:"structure"`
	Type                string                            `pulumi:"type"`
}

type OrcDatasetResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	Description         *string                                   `pulumi:"description"`
	Folder              *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName   LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location            interface{}                               `pulumi:"location"`
	OrcCompressionCodec interface{}                               `pulumi:"orcCompressionCodec"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema              interface{}                               `pulumi:"schema"`
	Structure           interface{}                               `pulumi:"structure"`
	Type                string                                    `pulumi:"type"`
}

type OrcFormat struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type OrcFormatResponse struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type OrcSink struct {
	DisableMetricsCollection interface{}       `pulumi:"disableMetricsCollection"`
	FormatSettings           *OrcWriteSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}       `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}       `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}       `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}       `pulumi:"storeSettings"`
	Type                     string            `pulumi:"type"`
	WriteBatchSize           interface{}       `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}       `pulumi:"writeBatchTimeout"`
}

type OrcSinkResponse struct {
	DisableMetricsCollection interface{}               `pulumi:"disableMetricsCollection"`
	FormatSettings           *OrcWriteSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}               `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}               `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}               `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}               `pulumi:"storeSettings"`
	Type                     string                    `pulumi:"type"`
	WriteBatchSize           interface{}               `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}               `pulumi:"writeBatchTimeout"`
}

type OrcSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type OrcSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type OrcWriteSettings struct {
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	Type           string      `pulumi:"type"`
}

type OrcWriteSettingsResponse struct {
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	Type           string      `pulumi:"type"`
}

type PackageStore struct {
	Name                      string          `pulumi:"name"`
	PackageStoreLinkedService EntityReference `pulumi:"packageStoreLinkedService"`
}

type PackageStoreResponse struct {
	Name                      string                  `pulumi:"name"`
	PackageStoreLinkedService EntityReferenceResponse `pulumi:"packageStoreLinkedService"`
}

type ParameterSpecification struct {
	DefaultValue interface{} `pulumi:"defaultValue"`
	Type         string      `pulumi:"type"`
}





type ParameterSpecificationInput interface {
	pulumi.Input

	ToParameterSpecificationOutput() ParameterSpecificationOutput
	ToParameterSpecificationOutputWithContext(context.Context) ParameterSpecificationOutput
}

type ParameterSpecificationArgs struct {
	DefaultValue pulumi.Input       `pulumi:"defaultValue"`
	Type         pulumi.StringInput `pulumi:"type"`
}

func (ParameterSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterSpecification)(nil)).Elem()
}

func (i ParameterSpecificationArgs) ToParameterSpecificationOutput() ParameterSpecificationOutput {
	return i.ToParameterSpecificationOutputWithContext(context.Background())
}

func (i ParameterSpecificationArgs) ToParameterSpecificationOutputWithContext(ctx context.Context) ParameterSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterSpecificationOutput)
}





type ParameterSpecificationMapInput interface {
	pulumi.Input

	ToParameterSpecificationMapOutput() ParameterSpecificationMapOutput
	ToParameterSpecificationMapOutputWithContext(context.Context) ParameterSpecificationMapOutput
}

type ParameterSpecificationMap map[string]ParameterSpecificationInput

func (ParameterSpecificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterSpecification)(nil)).Elem()
}

func (i ParameterSpecificationMap) ToParameterSpecificationMapOutput() ParameterSpecificationMapOutput {
	return i.ToParameterSpecificationMapOutputWithContext(context.Background())
}

func (i ParameterSpecificationMap) ToParameterSpecificationMapOutputWithContext(ctx context.Context) ParameterSpecificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterSpecificationMapOutput)
}

type ParameterSpecificationOutput struct{ *pulumi.OutputState }

func (ParameterSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterSpecification)(nil)).Elem()
}

func (o ParameterSpecificationOutput) ToParameterSpecificationOutput() ParameterSpecificationOutput {
	return o
}

func (o ParameterSpecificationOutput) ToParameterSpecificationOutputWithContext(ctx context.Context) ParameterSpecificationOutput {
	return o
}

func (o ParameterSpecificationOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterSpecification) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterSpecificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterSpecification) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterSpecificationMapOutput struct{ *pulumi.OutputState }

func (ParameterSpecificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterSpecification)(nil)).Elem()
}

func (o ParameterSpecificationMapOutput) ToParameterSpecificationMapOutput() ParameterSpecificationMapOutput {
	return o
}

func (o ParameterSpecificationMapOutput) ToParameterSpecificationMapOutputWithContext(ctx context.Context) ParameterSpecificationMapOutput {
	return o
}

func (o ParameterSpecificationMapOutput) MapIndex(k pulumi.StringInput) ParameterSpecificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterSpecification {
		return vs[0].(map[string]ParameterSpecification)[vs[1].(string)]
	}).(ParameterSpecificationOutput)
}

type ParameterSpecificationResponse struct {
	DefaultValue interface{} `pulumi:"defaultValue"`
	Type         string      `pulumi:"type"`
}

type ParameterSpecificationResponseOutput struct{ *pulumi.OutputState }

func (ParameterSpecificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterSpecificationResponse)(nil)).Elem()
}

func (o ParameterSpecificationResponseOutput) ToParameterSpecificationResponseOutput() ParameterSpecificationResponseOutput {
	return o
}

func (o ParameterSpecificationResponseOutput) ToParameterSpecificationResponseOutputWithContext(ctx context.Context) ParameterSpecificationResponseOutput {
	return o
}

func (o ParameterSpecificationResponseOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterSpecificationResponse) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterSpecificationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterSpecificationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterSpecificationResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterSpecificationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterSpecificationResponse)(nil)).Elem()
}

func (o ParameterSpecificationResponseMapOutput) ToParameterSpecificationResponseMapOutput() ParameterSpecificationResponseMapOutput {
	return o
}

func (o ParameterSpecificationResponseMapOutput) ToParameterSpecificationResponseMapOutputWithContext(ctx context.Context) ParameterSpecificationResponseMapOutput {
	return o
}

func (o ParameterSpecificationResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterSpecificationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterSpecificationResponse {
		return vs[0].(map[string]ParameterSpecificationResponse)[vs[1].(string)]
	}).(ParameterSpecificationResponseOutput)
}

type ParquetDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	CompressionCodec  interface{}                       `pulumi:"compressionCodec"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Location          interface{}                       `pulumi:"location"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type ParquetDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	CompressionCodec  interface{}                               `pulumi:"compressionCodec"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Location          interface{}                               `pulumi:"location"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type ParquetFormat struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type ParquetFormatResponse struct {
	Deserializer interface{} `pulumi:"deserializer"`
	Serializer   interface{} `pulumi:"serializer"`
	Type         string      `pulumi:"type"`
}

type ParquetSink struct {
	DisableMetricsCollection interface{}           `pulumi:"disableMetricsCollection"`
	FormatSettings           *ParquetWriteSettings `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}           `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}           `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}           `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}           `pulumi:"storeSettings"`
	Type                     string                `pulumi:"type"`
	WriteBatchSize           interface{}           `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}           `pulumi:"writeBatchTimeout"`
}

type ParquetSinkResponse struct {
	DisableMetricsCollection interface{}                   `pulumi:"disableMetricsCollection"`
	FormatSettings           *ParquetWriteSettingsResponse `pulumi:"formatSettings"`
	MaxConcurrentConnections interface{}                   `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{}                   `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                   `pulumi:"sinkRetryWait"`
	StoreSettings            interface{}                   `pulumi:"storeSettings"`
	Type                     string                        `pulumi:"type"`
	WriteBatchSize           interface{}                   `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                   `pulumi:"writeBatchTimeout"`
}

type ParquetSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type ParquetSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	StoreSettings            interface{} `pulumi:"storeSettings"`
	Type                     string      `pulumi:"type"`
}

type ParquetWriteSettings struct {
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	Type           string      `pulumi:"type"`
}

type ParquetWriteSettingsResponse struct {
	FileNamePrefix interface{} `pulumi:"fileNamePrefix"`
	MaxRowsPerFile interface{} `pulumi:"maxRowsPerFile"`
	Type           string      `pulumi:"type"`
}

type PaypalLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type PaypalLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type PaypalObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type PaypalObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type PaypalSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PaypalSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PhoenixLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	HttpPath                  interface{}                       `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
}

type PhoenixLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	HttpPath                  interface{}                               `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
}

type PhoenixObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type PhoenixObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type PhoenixSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PhoenixSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PipelineElapsedTimeMetricPolicy struct {
	Duration interface{} `pulumi:"duration"`
}





type PipelineElapsedTimeMetricPolicyInput interface {
	pulumi.Input

	ToPipelineElapsedTimeMetricPolicyOutput() PipelineElapsedTimeMetricPolicyOutput
	ToPipelineElapsedTimeMetricPolicyOutputWithContext(context.Context) PipelineElapsedTimeMetricPolicyOutput
}

type PipelineElapsedTimeMetricPolicyArgs struct {
	Duration pulumi.Input `pulumi:"duration"`
}

func (PipelineElapsedTimeMetricPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineElapsedTimeMetricPolicy)(nil)).Elem()
}

func (i PipelineElapsedTimeMetricPolicyArgs) ToPipelineElapsedTimeMetricPolicyOutput() PipelineElapsedTimeMetricPolicyOutput {
	return i.ToPipelineElapsedTimeMetricPolicyOutputWithContext(context.Background())
}

func (i PipelineElapsedTimeMetricPolicyArgs) ToPipelineElapsedTimeMetricPolicyOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineElapsedTimeMetricPolicyOutput)
}

func (i PipelineElapsedTimeMetricPolicyArgs) ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput {
	return i.ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(context.Background())
}

func (i PipelineElapsedTimeMetricPolicyArgs) ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineElapsedTimeMetricPolicyOutput).ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx)
}









type PipelineElapsedTimeMetricPolicyPtrInput interface {
	pulumi.Input

	ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput
	ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(context.Context) PipelineElapsedTimeMetricPolicyPtrOutput
}

type pipelineElapsedTimeMetricPolicyPtrType PipelineElapsedTimeMetricPolicyArgs

func PipelineElapsedTimeMetricPolicyPtr(v *PipelineElapsedTimeMetricPolicyArgs) PipelineElapsedTimeMetricPolicyPtrInput {
	return (*pipelineElapsedTimeMetricPolicyPtrType)(v)
}

func (*pipelineElapsedTimeMetricPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineElapsedTimeMetricPolicy)(nil)).Elem()
}

func (i *pipelineElapsedTimeMetricPolicyPtrType) ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput {
	return i.ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(context.Background())
}

func (i *pipelineElapsedTimeMetricPolicyPtrType) ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineElapsedTimeMetricPolicyPtrOutput)
}

type PipelineElapsedTimeMetricPolicyOutput struct{ *pulumi.OutputState }

func (PipelineElapsedTimeMetricPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineElapsedTimeMetricPolicy)(nil)).Elem()
}

func (o PipelineElapsedTimeMetricPolicyOutput) ToPipelineElapsedTimeMetricPolicyOutput() PipelineElapsedTimeMetricPolicyOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyOutput) ToPipelineElapsedTimeMetricPolicyOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyOutput) ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput {
	return o.ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(context.Background())
}

func (o PipelineElapsedTimeMetricPolicyOutput) ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineElapsedTimeMetricPolicy) *PipelineElapsedTimeMetricPolicy {
		return &v
	}).(PipelineElapsedTimeMetricPolicyPtrOutput)
}

func (o PipelineElapsedTimeMetricPolicyOutput) Duration() pulumi.AnyOutput {
	return o.ApplyT(func(v PipelineElapsedTimeMetricPolicy) interface{} { return v.Duration }).(pulumi.AnyOutput)
}

type PipelineElapsedTimeMetricPolicyPtrOutput struct{ *pulumi.OutputState }

func (PipelineElapsedTimeMetricPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineElapsedTimeMetricPolicy)(nil)).Elem()
}

func (o PipelineElapsedTimeMetricPolicyPtrOutput) ToPipelineElapsedTimeMetricPolicyPtrOutput() PipelineElapsedTimeMetricPolicyPtrOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyPtrOutput) ToPipelineElapsedTimeMetricPolicyPtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyPtrOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyPtrOutput) Elem() PipelineElapsedTimeMetricPolicyOutput {
	return o.ApplyT(func(v *PipelineElapsedTimeMetricPolicy) PipelineElapsedTimeMetricPolicy {
		if v != nil {
			return *v
		}
		var ret PipelineElapsedTimeMetricPolicy
		return ret
	}).(PipelineElapsedTimeMetricPolicyOutput)
}

func (o PipelineElapsedTimeMetricPolicyPtrOutput) Duration() pulumi.AnyOutput {
	return o.ApplyT(func(v *PipelineElapsedTimeMetricPolicy) interface{} {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.AnyOutput)
}

type PipelineElapsedTimeMetricPolicyResponse struct {
	Duration interface{} `pulumi:"duration"`
}

type PipelineElapsedTimeMetricPolicyResponseOutput struct{ *pulumi.OutputState }

func (PipelineElapsedTimeMetricPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineElapsedTimeMetricPolicyResponse)(nil)).Elem()
}

func (o PipelineElapsedTimeMetricPolicyResponseOutput) ToPipelineElapsedTimeMetricPolicyResponseOutput() PipelineElapsedTimeMetricPolicyResponseOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyResponseOutput) ToPipelineElapsedTimeMetricPolicyResponseOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyResponseOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyResponseOutput) Duration() pulumi.AnyOutput {
	return o.ApplyT(func(v PipelineElapsedTimeMetricPolicyResponse) interface{} { return v.Duration }).(pulumi.AnyOutput)
}

type PipelineElapsedTimeMetricPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineElapsedTimeMetricPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineElapsedTimeMetricPolicyResponse)(nil)).Elem()
}

func (o PipelineElapsedTimeMetricPolicyResponsePtrOutput) ToPipelineElapsedTimeMetricPolicyResponsePtrOutput() PipelineElapsedTimeMetricPolicyResponsePtrOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyResponsePtrOutput) ToPipelineElapsedTimeMetricPolicyResponsePtrOutputWithContext(ctx context.Context) PipelineElapsedTimeMetricPolicyResponsePtrOutput {
	return o
}

func (o PipelineElapsedTimeMetricPolicyResponsePtrOutput) Elem() PipelineElapsedTimeMetricPolicyResponseOutput {
	return o.ApplyT(func(v *PipelineElapsedTimeMetricPolicyResponse) PipelineElapsedTimeMetricPolicyResponse {
		if v != nil {
			return *v
		}
		var ret PipelineElapsedTimeMetricPolicyResponse
		return ret
	}).(PipelineElapsedTimeMetricPolicyResponseOutput)
}

func (o PipelineElapsedTimeMetricPolicyResponsePtrOutput) Duration() pulumi.AnyOutput {
	return o.ApplyT(func(v *PipelineElapsedTimeMetricPolicyResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.AnyOutput)
}

type PipelineFolder struct {
	Name *string `pulumi:"name"`
}





type PipelineFolderInput interface {
	pulumi.Input

	ToPipelineFolderOutput() PipelineFolderOutput
	ToPipelineFolderOutputWithContext(context.Context) PipelineFolderOutput
}

type PipelineFolderArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (PipelineFolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineFolder)(nil)).Elem()
}

func (i PipelineFolderArgs) ToPipelineFolderOutput() PipelineFolderOutput {
	return i.ToPipelineFolderOutputWithContext(context.Background())
}

func (i PipelineFolderArgs) ToPipelineFolderOutputWithContext(ctx context.Context) PipelineFolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineFolderOutput)
}

func (i PipelineFolderArgs) ToPipelineFolderPtrOutput() PipelineFolderPtrOutput {
	return i.ToPipelineFolderPtrOutputWithContext(context.Background())
}

func (i PipelineFolderArgs) ToPipelineFolderPtrOutputWithContext(ctx context.Context) PipelineFolderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineFolderOutput).ToPipelineFolderPtrOutputWithContext(ctx)
}









type PipelineFolderPtrInput interface {
	pulumi.Input

	ToPipelineFolderPtrOutput() PipelineFolderPtrOutput
	ToPipelineFolderPtrOutputWithContext(context.Context) PipelineFolderPtrOutput
}

type pipelineFolderPtrType PipelineFolderArgs

func PipelineFolderPtr(v *PipelineFolderArgs) PipelineFolderPtrInput {
	return (*pipelineFolderPtrType)(v)
}

func (*pipelineFolderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineFolder)(nil)).Elem()
}

func (i *pipelineFolderPtrType) ToPipelineFolderPtrOutput() PipelineFolderPtrOutput {
	return i.ToPipelineFolderPtrOutputWithContext(context.Background())
}

func (i *pipelineFolderPtrType) ToPipelineFolderPtrOutputWithContext(ctx context.Context) PipelineFolderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineFolderPtrOutput)
}

type PipelineFolderOutput struct{ *pulumi.OutputState }

func (PipelineFolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineFolder)(nil)).Elem()
}

func (o PipelineFolderOutput) ToPipelineFolderOutput() PipelineFolderOutput {
	return o
}

func (o PipelineFolderOutput) ToPipelineFolderOutputWithContext(ctx context.Context) PipelineFolderOutput {
	return o
}

func (o PipelineFolderOutput) ToPipelineFolderPtrOutput() PipelineFolderPtrOutput {
	return o.ToPipelineFolderPtrOutputWithContext(context.Background())
}

func (o PipelineFolderOutput) ToPipelineFolderPtrOutputWithContext(ctx context.Context) PipelineFolderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineFolder) *PipelineFolder {
		return &v
	}).(PipelineFolderPtrOutput)
}

func (o PipelineFolderOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineFolder) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PipelineFolderPtrOutput struct{ *pulumi.OutputState }

func (PipelineFolderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineFolder)(nil)).Elem()
}

func (o PipelineFolderPtrOutput) ToPipelineFolderPtrOutput() PipelineFolderPtrOutput {
	return o
}

func (o PipelineFolderPtrOutput) ToPipelineFolderPtrOutputWithContext(ctx context.Context) PipelineFolderPtrOutput {
	return o
}

func (o PipelineFolderPtrOutput) Elem() PipelineFolderOutput {
	return o.ApplyT(func(v *PipelineFolder) PipelineFolder {
		if v != nil {
			return *v
		}
		var ret PipelineFolder
		return ret
	}).(PipelineFolderOutput)
}

func (o PipelineFolderPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineFolder) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type PipelinePolicy struct {
	ElapsedTimeMetric *PipelineElapsedTimeMetricPolicy `pulumi:"elapsedTimeMetric"`
}





type PipelinePolicyInput interface {
	pulumi.Input

	ToPipelinePolicyOutput() PipelinePolicyOutput
	ToPipelinePolicyOutputWithContext(context.Context) PipelinePolicyOutput
}

type PipelinePolicyArgs struct {
	ElapsedTimeMetric PipelineElapsedTimeMetricPolicyPtrInput `pulumi:"elapsedTimeMetric"`
}

func (PipelinePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelinePolicy)(nil)).Elem()
}

func (i PipelinePolicyArgs) ToPipelinePolicyOutput() PipelinePolicyOutput {
	return i.ToPipelinePolicyOutputWithContext(context.Background())
}

func (i PipelinePolicyArgs) ToPipelinePolicyOutputWithContext(ctx context.Context) PipelinePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePolicyOutput)
}

func (i PipelinePolicyArgs) ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput {
	return i.ToPipelinePolicyPtrOutputWithContext(context.Background())
}

func (i PipelinePolicyArgs) ToPipelinePolicyPtrOutputWithContext(ctx context.Context) PipelinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePolicyOutput).ToPipelinePolicyPtrOutputWithContext(ctx)
}









type PipelinePolicyPtrInput interface {
	pulumi.Input

	ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput
	ToPipelinePolicyPtrOutputWithContext(context.Context) PipelinePolicyPtrOutput
}

type pipelinePolicyPtrType PipelinePolicyArgs

func PipelinePolicyPtr(v *PipelinePolicyArgs) PipelinePolicyPtrInput {
	return (*pipelinePolicyPtrType)(v)
}

func (*pipelinePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelinePolicy)(nil)).Elem()
}

func (i *pipelinePolicyPtrType) ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput {
	return i.ToPipelinePolicyPtrOutputWithContext(context.Background())
}

func (i *pipelinePolicyPtrType) ToPipelinePolicyPtrOutputWithContext(ctx context.Context) PipelinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePolicyPtrOutput)
}

type PipelinePolicyOutput struct{ *pulumi.OutputState }

func (PipelinePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelinePolicy)(nil)).Elem()
}

func (o PipelinePolicyOutput) ToPipelinePolicyOutput() PipelinePolicyOutput {
	return o
}

func (o PipelinePolicyOutput) ToPipelinePolicyOutputWithContext(ctx context.Context) PipelinePolicyOutput {
	return o
}

func (o PipelinePolicyOutput) ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput {
	return o.ToPipelinePolicyPtrOutputWithContext(context.Background())
}

func (o PipelinePolicyOutput) ToPipelinePolicyPtrOutputWithContext(ctx context.Context) PipelinePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelinePolicy) *PipelinePolicy {
		return &v
	}).(PipelinePolicyPtrOutput)
}

func (o PipelinePolicyOutput) ElapsedTimeMetric() PipelineElapsedTimeMetricPolicyPtrOutput {
	return o.ApplyT(func(v PipelinePolicy) *PipelineElapsedTimeMetricPolicy { return v.ElapsedTimeMetric }).(PipelineElapsedTimeMetricPolicyPtrOutput)
}

type PipelinePolicyPtrOutput struct{ *pulumi.OutputState }

func (PipelinePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelinePolicy)(nil)).Elem()
}

func (o PipelinePolicyPtrOutput) ToPipelinePolicyPtrOutput() PipelinePolicyPtrOutput {
	return o
}

func (o PipelinePolicyPtrOutput) ToPipelinePolicyPtrOutputWithContext(ctx context.Context) PipelinePolicyPtrOutput {
	return o
}

func (o PipelinePolicyPtrOutput) Elem() PipelinePolicyOutput {
	return o.ApplyT(func(v *PipelinePolicy) PipelinePolicy {
		if v != nil {
			return *v
		}
		var ret PipelinePolicy
		return ret
	}).(PipelinePolicyOutput)
}

func (o PipelinePolicyPtrOutput) ElapsedTimeMetric() PipelineElapsedTimeMetricPolicyPtrOutput {
	return o.ApplyT(func(v *PipelinePolicy) *PipelineElapsedTimeMetricPolicy {
		if v == nil {
			return nil
		}
		return v.ElapsedTimeMetric
	}).(PipelineElapsedTimeMetricPolicyPtrOutput)
}

type PipelinePolicyResponse struct {
	ElapsedTimeMetric *PipelineElapsedTimeMetricPolicyResponse `pulumi:"elapsedTimeMetric"`
}

type PipelinePolicyResponseOutput struct{ *pulumi.OutputState }

func (PipelinePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelinePolicyResponse)(nil)).Elem()
}

func (o PipelinePolicyResponseOutput) ToPipelinePolicyResponseOutput() PipelinePolicyResponseOutput {
	return o
}

func (o PipelinePolicyResponseOutput) ToPipelinePolicyResponseOutputWithContext(ctx context.Context) PipelinePolicyResponseOutput {
	return o
}

func (o PipelinePolicyResponseOutput) ElapsedTimeMetric() PipelineElapsedTimeMetricPolicyResponsePtrOutput {
	return o.ApplyT(func(v PipelinePolicyResponse) *PipelineElapsedTimeMetricPolicyResponse { return v.ElapsedTimeMetric }).(PipelineElapsedTimeMetricPolicyResponsePtrOutput)
}

type PipelinePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelinePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelinePolicyResponse)(nil)).Elem()
}

func (o PipelinePolicyResponsePtrOutput) ToPipelinePolicyResponsePtrOutput() PipelinePolicyResponsePtrOutput {
	return o
}

func (o PipelinePolicyResponsePtrOutput) ToPipelinePolicyResponsePtrOutputWithContext(ctx context.Context) PipelinePolicyResponsePtrOutput {
	return o
}

func (o PipelinePolicyResponsePtrOutput) Elem() PipelinePolicyResponseOutput {
	return o.ApplyT(func(v *PipelinePolicyResponse) PipelinePolicyResponse {
		if v != nil {
			return *v
		}
		var ret PipelinePolicyResponse
		return ret
	}).(PipelinePolicyResponseOutput)
}

func (o PipelinePolicyResponsePtrOutput) ElapsedTimeMetric() PipelineElapsedTimeMetricPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PipelinePolicyResponse) *PipelineElapsedTimeMetricPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ElapsedTimeMetric
	}).(PipelineElapsedTimeMetricPolicyResponsePtrOutput)
}

type PipelineReference struct {
	Name          *string `pulumi:"name"`
	ReferenceName string  `pulumi:"referenceName"`
	Type          string  `pulumi:"type"`
}

type PipelineReferenceResponse struct {
	Name          *string `pulumi:"name"`
	ReferenceName string  `pulumi:"referenceName"`
	Type          string  `pulumi:"type"`
}

type PipelineResponseFolder struct {
	Name *string `pulumi:"name"`
}

type PipelineResponseFolderOutput struct{ *pulumi.OutputState }

func (PipelineResponseFolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineResponseFolder)(nil)).Elem()
}

func (o PipelineResponseFolderOutput) ToPipelineResponseFolderOutput() PipelineResponseFolderOutput {
	return o
}

func (o PipelineResponseFolderOutput) ToPipelineResponseFolderOutputWithContext(ctx context.Context) PipelineResponseFolderOutput {
	return o
}

func (o PipelineResponseFolderOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineResponseFolder) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PipelineResponseFolderPtrOutput struct{ *pulumi.OutputState }

func (PipelineResponseFolderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineResponseFolder)(nil)).Elem()
}

func (o PipelineResponseFolderPtrOutput) ToPipelineResponseFolderPtrOutput() PipelineResponseFolderPtrOutput {
	return o
}

func (o PipelineResponseFolderPtrOutput) ToPipelineResponseFolderPtrOutputWithContext(ctx context.Context) PipelineResponseFolderPtrOutput {
	return o
}

func (o PipelineResponseFolderPtrOutput) Elem() PipelineResponseFolderOutput {
	return o.ApplyT(func(v *PipelineResponseFolder) PipelineResponseFolder {
		if v != nil {
			return *v
		}
		var ret PipelineResponseFolder
		return ret
	}).(PipelineResponseFolderOutput)
}

func (o PipelineResponseFolderPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineResponseFolder) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type PolybaseSettings struct {
	RejectSampleValue interface{} `pulumi:"rejectSampleValue"`
	RejectType        *string     `pulumi:"rejectType"`
	RejectValue       interface{} `pulumi:"rejectValue"`
	UseTypeDefault    interface{} `pulumi:"useTypeDefault"`
}

type PolybaseSettingsResponse struct {
	RejectSampleValue interface{} `pulumi:"rejectSampleValue"`
	RejectType        *string     `pulumi:"rejectType"`
	RejectValue       interface{} `pulumi:"rejectValue"`
	UseTypeDefault    interface{} `pulumi:"useTypeDefault"`
}

type PostgreSqlLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type PostgreSqlLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type PostgreSqlSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PostgreSqlSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PostgreSqlTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type PostgreSqlTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type PowerQuerySink struct {
	Dataset                   *DatasetReference       `pulumi:"dataset"`
	Description               *string                 `pulumi:"description"`
	Flowlet                   *DataFlowReference      `pulumi:"flowlet"`
	LinkedService             *LinkedServiceReference `pulumi:"linkedService"`
	Name                      string                  `pulumi:"name"`
	RejectedDataLinkedService *LinkedServiceReference `pulumi:"rejectedDataLinkedService"`
	SchemaLinkedService       *LinkedServiceReference `pulumi:"schemaLinkedService"`
	Script                    *string                 `pulumi:"script"`
}

type PowerQuerySinkMapping struct {
	DataflowSinks []PowerQuerySink `pulumi:"dataflowSinks"`
	QueryName     *string          `pulumi:"queryName"`
}

type PowerQuerySinkMappingResponse struct {
	DataflowSinks []PowerQuerySinkResponse `pulumi:"dataflowSinks"`
	QueryName     *string                  `pulumi:"queryName"`
}

type PowerQuerySinkResponse struct {
	Dataset                   *DatasetReferenceResponse       `pulumi:"dataset"`
	Description               *string                         `pulumi:"description"`
	Flowlet                   *DataFlowReferenceResponse      `pulumi:"flowlet"`
	LinkedService             *LinkedServiceReferenceResponse `pulumi:"linkedService"`
	Name                      string                          `pulumi:"name"`
	RejectedDataLinkedService *LinkedServiceReferenceResponse `pulumi:"rejectedDataLinkedService"`
	SchemaLinkedService       *LinkedServiceReferenceResponse `pulumi:"schemaLinkedService"`
	Script                    *string                         `pulumi:"script"`
}

type PowerQuerySource struct {
	Dataset             *DatasetReference       `pulumi:"dataset"`
	Description         *string                 `pulumi:"description"`
	Flowlet             *DataFlowReference      `pulumi:"flowlet"`
	LinkedService       *LinkedServiceReference `pulumi:"linkedService"`
	Name                string                  `pulumi:"name"`
	SchemaLinkedService *LinkedServiceReference `pulumi:"schemaLinkedService"`
	Script              *string                 `pulumi:"script"`
}

type PowerQuerySourceResponse struct {
	Dataset             *DatasetReferenceResponse       `pulumi:"dataset"`
	Description         *string                         `pulumi:"description"`
	Flowlet             *DataFlowReferenceResponse      `pulumi:"flowlet"`
	LinkedService       *LinkedServiceReferenceResponse `pulumi:"linkedService"`
	Name                string                          `pulumi:"name"`
	SchemaLinkedService *LinkedServiceReferenceResponse `pulumi:"schemaLinkedService"`
	Script              *string                         `pulumi:"script"`
}

type PrestoLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	Catalog                   interface{}                       `pulumi:"catalog"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	ServerVersion             interface{}                       `pulumi:"serverVersion"`
	TimeZoneID                interface{}                       `pulumi:"timeZoneID"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
}

type PrestoLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	Catalog                   interface{}                               `pulumi:"catalog"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	ServerVersion             interface{}                               `pulumi:"serverVersion"`
	TimeZoneID                interface{}                               `pulumi:"timeZoneID"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
}

type PrestoObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type PrestoObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type PrestoSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PrestoSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionApprovalRequest struct {
	PrivateEndpoint                   *PrivateEndpoint            `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateLinkConnectionApprovalRequestInput interface {
	pulumi.Input

	ToPrivateLinkConnectionApprovalRequestOutput() PrivateLinkConnectionApprovalRequestOutput
	ToPrivateLinkConnectionApprovalRequestOutputWithContext(context.Context) PrivateLinkConnectionApprovalRequestOutput
}

type PrivateLinkConnectionApprovalRequestArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput            `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateLinkConnectionApprovalRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionApprovalRequest)(nil)).Elem()
}

func (i PrivateLinkConnectionApprovalRequestArgs) ToPrivateLinkConnectionApprovalRequestOutput() PrivateLinkConnectionApprovalRequestOutput {
	return i.ToPrivateLinkConnectionApprovalRequestOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionApprovalRequestArgs) ToPrivateLinkConnectionApprovalRequestOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionApprovalRequestOutput)
}

func (i PrivateLinkConnectionApprovalRequestArgs) ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput {
	return i.ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionApprovalRequestArgs) ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionApprovalRequestOutput).ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx)
}









type PrivateLinkConnectionApprovalRequestPtrInput interface {
	pulumi.Input

	ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput
	ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(context.Context) PrivateLinkConnectionApprovalRequestPtrOutput
}

type privateLinkConnectionApprovalRequestPtrType PrivateLinkConnectionApprovalRequestArgs

func PrivateLinkConnectionApprovalRequestPtr(v *PrivateLinkConnectionApprovalRequestArgs) PrivateLinkConnectionApprovalRequestPtrInput {
	return (*privateLinkConnectionApprovalRequestPtrType)(v)
}

func (*privateLinkConnectionApprovalRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionApprovalRequest)(nil)).Elem()
}

func (i *privateLinkConnectionApprovalRequestPtrType) ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput {
	return i.ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(context.Background())
}

func (i *privateLinkConnectionApprovalRequestPtrType) ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionApprovalRequestPtrOutput)
}

type PrivateLinkConnectionApprovalRequestOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionApprovalRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionApprovalRequest)(nil)).Elem()
}

func (o PrivateLinkConnectionApprovalRequestOutput) ToPrivateLinkConnectionApprovalRequestOutput() PrivateLinkConnectionApprovalRequestOutput {
	return o
}

func (o PrivateLinkConnectionApprovalRequestOutput) ToPrivateLinkConnectionApprovalRequestOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestOutput {
	return o
}

func (o PrivateLinkConnectionApprovalRequestOutput) ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput {
	return o.ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(context.Background())
}

func (o PrivateLinkConnectionApprovalRequestOutput) ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkConnectionApprovalRequest) *PrivateLinkConnectionApprovalRequest {
		return &v
	}).(PrivateLinkConnectionApprovalRequestPtrOutput)
}

func (o PrivateLinkConnectionApprovalRequestOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionApprovalRequest) *PrivateEndpoint { return v.PrivateEndpoint }).(PrivateEndpointPtrOutput)
}

func (o PrivateLinkConnectionApprovalRequestOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionApprovalRequest) *PrivateLinkConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStatePtrOutput)
}

type PrivateLinkConnectionApprovalRequestPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionApprovalRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionApprovalRequest)(nil)).Elem()
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) ToPrivateLinkConnectionApprovalRequestPtrOutput() PrivateLinkConnectionApprovalRequestPtrOutput {
	return o
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) ToPrivateLinkConnectionApprovalRequestPtrOutputWithContext(ctx context.Context) PrivateLinkConnectionApprovalRequestPtrOutput {
	return o
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) Elem() PrivateLinkConnectionApprovalRequestOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionApprovalRequest) PrivateLinkConnectionApprovalRequest {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionApprovalRequest
		return ret
	}).(PrivateLinkConnectionApprovalRequestOutput)
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionApprovalRequest) *PrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateLinkConnectionApprovalRequestPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionApprovalRequest) *PrivateLinkConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStatePtrOutput)
}

type PrivateLinkConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput
	ToPrivateLinkConnectionStateOutputWithContext(context.Context) PrivateLinkConnectionStateOutput
}

type PrivateLinkConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionState)(nil)).Elem()
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput {
	return i.ToPrivateLinkConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStateOutputWithContext(ctx context.Context) PrivateLinkConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateOutput)
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return i.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateOutput).ToPrivateLinkConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput
	ToPrivateLinkConnectionStatePtrOutputWithContext(context.Context) PrivateLinkConnectionStatePtrOutput
}

type privateLinkConnectionStatePtrType PrivateLinkConnectionStateArgs

func PrivateLinkConnectionStatePtr(v *PrivateLinkConnectionStateArgs) PrivateLinkConnectionStatePtrInput {
	return (*privateLinkConnectionStatePtrType)(v)
}

func (*privateLinkConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionState)(nil)).Elem()
}

func (i *privateLinkConnectionStatePtrType) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return i.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkConnectionStatePtrType) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStatePtrOutput)
}

type PrivateLinkConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionState)(nil)).Elem()
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput {
	return o
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStateOutputWithContext(ctx context.Context) PrivateLinkConnectionStateOutput {
	return o
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return o.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkConnectionState) *PrivateLinkConnectionState {
		return &v
	}).(PrivateLinkConnectionStatePtrOutput)
}

func (o PrivateLinkConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionState)(nil)).Elem()
}

func (o PrivateLinkConnectionStatePtrOutput) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkConnectionStatePtrOutput) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkConnectionStatePtrOutput) Elem() PrivateLinkConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) PrivateLinkConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionState
		return ret
	}).(PrivateLinkConnectionStateOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutput() PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Elem() PrivateLinkConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) PrivateLinkConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionStateResponse
		return ret
	}).(PrivateLinkConnectionStateResponseOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PurviewConfiguration struct {
	PurviewResourceId *string `pulumi:"purviewResourceId"`
}





type PurviewConfigurationInput interface {
	pulumi.Input

	ToPurviewConfigurationOutput() PurviewConfigurationOutput
	ToPurviewConfigurationOutputWithContext(context.Context) PurviewConfigurationOutput
}

type PurviewConfigurationArgs struct {
	PurviewResourceId pulumi.StringPtrInput `pulumi:"purviewResourceId"`
}

func (PurviewConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfiguration)(nil)).Elem()
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationOutput() PurviewConfigurationOutput {
	return i.ToPurviewConfigurationOutputWithContext(context.Background())
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationOutputWithContext(ctx context.Context) PurviewConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationOutput)
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return i.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (i PurviewConfigurationArgs) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationOutput).ToPurviewConfigurationPtrOutputWithContext(ctx)
}









type PurviewConfigurationPtrInput interface {
	pulumi.Input

	ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput
	ToPurviewConfigurationPtrOutputWithContext(context.Context) PurviewConfigurationPtrOutput
}

type purviewConfigurationPtrType PurviewConfigurationArgs

func PurviewConfigurationPtr(v *PurviewConfigurationArgs) PurviewConfigurationPtrInput {
	return (*purviewConfigurationPtrType)(v)
}

func (*purviewConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfiguration)(nil)).Elem()
}

func (i *purviewConfigurationPtrType) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return i.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (i *purviewConfigurationPtrType) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurviewConfigurationPtrOutput)
}

type PurviewConfigurationOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfiguration)(nil)).Elem()
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationOutput() PurviewConfigurationOutput {
	return o
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationOutputWithContext(ctx context.Context) PurviewConfigurationOutput {
	return o
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return o.ToPurviewConfigurationPtrOutputWithContext(context.Background())
}

func (o PurviewConfigurationOutput) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PurviewConfiguration) *PurviewConfiguration {
		return &v
	}).(PurviewConfigurationPtrOutput)
}

func (o PurviewConfigurationOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurviewConfiguration) *string { return v.PurviewResourceId }).(pulumi.StringPtrOutput)
}

type PurviewConfigurationPtrOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfiguration)(nil)).Elem()
}

func (o PurviewConfigurationPtrOutput) ToPurviewConfigurationPtrOutput() PurviewConfigurationPtrOutput {
	return o
}

func (o PurviewConfigurationPtrOutput) ToPurviewConfigurationPtrOutputWithContext(ctx context.Context) PurviewConfigurationPtrOutput {
	return o
}

func (o PurviewConfigurationPtrOutput) Elem() PurviewConfigurationOutput {
	return o.ApplyT(func(v *PurviewConfiguration) PurviewConfiguration {
		if v != nil {
			return *v
		}
		var ret PurviewConfiguration
		return ret
	}).(PurviewConfigurationOutput)
}

func (o PurviewConfigurationPtrOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurviewConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PurviewResourceId
	}).(pulumi.StringPtrOutput)
}

type PurviewConfigurationResponse struct {
	PurviewResourceId *string `pulumi:"purviewResourceId"`
}

type PurviewConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurviewConfigurationResponse)(nil)).Elem()
}

func (o PurviewConfigurationResponseOutput) ToPurviewConfigurationResponseOutput() PurviewConfigurationResponseOutput {
	return o
}

func (o PurviewConfigurationResponseOutput) ToPurviewConfigurationResponseOutputWithContext(ctx context.Context) PurviewConfigurationResponseOutput {
	return o
}

func (o PurviewConfigurationResponseOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurviewConfigurationResponse) *string { return v.PurviewResourceId }).(pulumi.StringPtrOutput)
}

type PurviewConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (PurviewConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurviewConfigurationResponse)(nil)).Elem()
}

func (o PurviewConfigurationResponsePtrOutput) ToPurviewConfigurationResponsePtrOutput() PurviewConfigurationResponsePtrOutput {
	return o
}

func (o PurviewConfigurationResponsePtrOutput) ToPurviewConfigurationResponsePtrOutputWithContext(ctx context.Context) PurviewConfigurationResponsePtrOutput {
	return o
}

func (o PurviewConfigurationResponsePtrOutput) Elem() PurviewConfigurationResponseOutput {
	return o.ApplyT(func(v *PurviewConfigurationResponse) PurviewConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret PurviewConfigurationResponse
		return ret
	}).(PurviewConfigurationResponseOutput)
}

func (o PurviewConfigurationResponsePtrOutput) PurviewResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurviewConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PurviewResourceId
	}).(pulumi.StringPtrOutput)
}

type QuickBooksLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	AccessTokenSecret     interface{}                       `pulumi:"accessTokenSecret"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	CompanyId             interface{}                       `pulumi:"companyId"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                       `pulumi:"connectionProperties"`
	ConsumerKey           interface{}                       `pulumi:"consumerKey"`
	ConsumerSecret        interface{}                       `pulumi:"consumerSecret"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
}

type QuickBooksLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	AccessTokenSecret     interface{}                               `pulumi:"accessTokenSecret"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	CompanyId             interface{}                               `pulumi:"companyId"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                               `pulumi:"connectionProperties"`
	ConsumerKey           interface{}                               `pulumi:"consumerKey"`
	ConsumerSecret        interface{}                               `pulumi:"consumerSecret"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
}

type QuickBooksObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type QuickBooksObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type QuickBooksSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type QuickBooksSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type QuickbaseLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	UserToken           interface{}                       `pulumi:"userToken"`
}

type QuickbaseLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	UserToken           interface{}                               `pulumi:"userToken"`
}

type RecurrenceSchedule struct {
	Hours              []int                          `pulumi:"hours"`
	Minutes            []int                          `pulumi:"minutes"`
	MonthDays          []int                          `pulumi:"monthDays"`
	MonthlyOccurrences []RecurrenceScheduleOccurrence `pulumi:"monthlyOccurrences"`
	WeekDays           []DaysOfWeek                   `pulumi:"weekDays"`
}

type RecurrenceScheduleOccurrence struct {
	Day        *DayOfWeek `pulumi:"day"`
	Occurrence *int       `pulumi:"occurrence"`
}

type RecurrenceScheduleOccurrenceResponse struct {
	Day        *string `pulumi:"day"`
	Occurrence *int    `pulumi:"occurrence"`
}

type RecurrenceScheduleResponse struct {
	Hours              []int                                  `pulumi:"hours"`
	Minutes            []int                                  `pulumi:"minutes"`
	MonthDays          []int                                  `pulumi:"monthDays"`
	MonthlyOccurrences []RecurrenceScheduleOccurrenceResponse `pulumi:"monthlyOccurrences"`
	WeekDays           []string                               `pulumi:"weekDays"`
}

type RedirectIncompatibleRowSettings struct {
	LinkedServiceName interface{} `pulumi:"linkedServiceName"`
	Path              interface{} `pulumi:"path"`
}

type RedirectIncompatibleRowSettingsResponse struct {
	LinkedServiceName interface{} `pulumi:"linkedServiceName"`
	Path              interface{} `pulumi:"path"`
}

type RedshiftUnloadSettings struct {
	BucketName          interface{}            `pulumi:"bucketName"`
	S3LinkedServiceName LinkedServiceReference `pulumi:"s3LinkedServiceName"`
}

type RedshiftUnloadSettingsResponse struct {
	BucketName          interface{}                    `pulumi:"bucketName"`
	S3LinkedServiceName LinkedServiceReferenceResponse `pulumi:"s3LinkedServiceName"`
}

type RelationalSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RelationalSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RelationalTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type RelationalTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type RemotePrivateEndpointConnectionResponse struct {
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
}

type RemotePrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionResponseOutput) ToRemotePrivateEndpointConnectionResponseOutput() RemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponseOutput) ToRemotePrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponseOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *ArmIdWrapperResponse { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type RerunTumblingWindowTrigger struct {
	Annotations        []interface{} `pulumi:"annotations"`
	Description        *string       `pulumi:"description"`
	ParentTrigger      interface{}   `pulumi:"parentTrigger"`
	RequestedEndTime   string        `pulumi:"requestedEndTime"`
	RequestedStartTime string        `pulumi:"requestedStartTime"`
	RerunConcurrency   int           `pulumi:"rerunConcurrency"`
	Type               string        `pulumi:"type"`
}

type RerunTumblingWindowTriggerResponse struct {
	Annotations        []interface{} `pulumi:"annotations"`
	Description        *string       `pulumi:"description"`
	ParentTrigger      interface{}   `pulumi:"parentTrigger"`
	RequestedEndTime   string        `pulumi:"requestedEndTime"`
	RequestedStartTime string        `pulumi:"requestedStartTime"`
	RerunConcurrency   int           `pulumi:"rerunConcurrency"`
	RuntimeState       string        `pulumi:"runtimeState"`
	Type               string        `pulumi:"type"`
}

type ResponsysLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type ResponsysLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type ResponsysObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type ResponsysObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type ResponsysSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ResponsysSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RestResourceDataset struct {
	AdditionalHeaders interface{}                       `pulumi:"additionalHeaders"`
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	PaginationRules   interface{}                       `pulumi:"paginationRules"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	RelativeUrl       interface{}                       `pulumi:"relativeUrl"`
	RequestBody       interface{}                       `pulumi:"requestBody"`
	RequestMethod     interface{}                       `pulumi:"requestMethod"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type RestResourceDatasetResponse struct {
	AdditionalHeaders interface{}                               `pulumi:"additionalHeaders"`
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	PaginationRules   interface{}                               `pulumi:"paginationRules"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	RelativeUrl       interface{}                               `pulumi:"relativeUrl"`
	RequestBody       interface{}                               `pulumi:"requestBody"`
	RequestMethod     interface{}                               `pulumi:"requestMethod"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type RestServiceLinkedService struct {
	AadResourceId                     interface{}                       `pulumi:"aadResourceId"`
	Annotations                       []interface{}                     `pulumi:"annotations"`
	AuthHeaders                       interface{}                       `pulumi:"authHeaders"`
	AuthenticationType                string                            `pulumi:"authenticationType"`
	AzureCloudType                    interface{}                       `pulumi:"azureCloudType"`
	ClientId                          interface{}                       `pulumi:"clientId"`
	ClientSecret                      interface{}                       `pulumi:"clientSecret"`
	ConnectVia                        *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Credential                        *CredentialReference              `pulumi:"credential"`
	Description                       *string                           `pulumi:"description"`
	EnableServerCertificateValidation interface{}                       `pulumi:"enableServerCertificateValidation"`
	EncryptedCredential               interface{}                       `pulumi:"encryptedCredential"`
	Parameters                        map[string]ParameterSpecification `pulumi:"parameters"`
	Password                          interface{}                       `pulumi:"password"`
	Resource                          interface{}                       `pulumi:"resource"`
	Scope                             interface{}                       `pulumi:"scope"`
	ServicePrincipalId                interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey               interface{}                       `pulumi:"servicePrincipalKey"`
	Tenant                            interface{}                       `pulumi:"tenant"`
	TokenEndpoint                     interface{}                       `pulumi:"tokenEndpoint"`
	Type                              string                            `pulumi:"type"`
	Url                               interface{}                       `pulumi:"url"`
	UserName                          interface{}                       `pulumi:"userName"`
}

type RestServiceLinkedServiceResponse struct {
	AadResourceId                     interface{}                               `pulumi:"aadResourceId"`
	Annotations                       []interface{}                             `pulumi:"annotations"`
	AuthHeaders                       interface{}                               `pulumi:"authHeaders"`
	AuthenticationType                string                                    `pulumi:"authenticationType"`
	AzureCloudType                    interface{}                               `pulumi:"azureCloudType"`
	ClientId                          interface{}                               `pulumi:"clientId"`
	ClientSecret                      interface{}                               `pulumi:"clientSecret"`
	ConnectVia                        *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Credential                        *CredentialReferenceResponse              `pulumi:"credential"`
	Description                       *string                                   `pulumi:"description"`
	EnableServerCertificateValidation interface{}                               `pulumi:"enableServerCertificateValidation"`
	EncryptedCredential               interface{}                               `pulumi:"encryptedCredential"`
	Parameters                        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                          interface{}                               `pulumi:"password"`
	Resource                          interface{}                               `pulumi:"resource"`
	Scope                             interface{}                               `pulumi:"scope"`
	ServicePrincipalId                interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey               interface{}                               `pulumi:"servicePrincipalKey"`
	Tenant                            interface{}                               `pulumi:"tenant"`
	TokenEndpoint                     interface{}                               `pulumi:"tokenEndpoint"`
	Type                              string                                    `pulumi:"type"`
	Url                               interface{}                               `pulumi:"url"`
	UserName                          interface{}                               `pulumi:"userName"`
}

type RestSink struct {
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpCompressionType      interface{} `pulumi:"httpCompressionType"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	RequestInterval          interface{} `pulumi:"requestInterval"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type RestSinkResponse struct {
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpCompressionType      interface{} `pulumi:"httpCompressionType"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	RequestInterval          interface{} `pulumi:"requestInterval"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
}

type RestSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PaginationRules          interface{} `pulumi:"paginationRules"`
	RequestBody              interface{} `pulumi:"requestBody"`
	RequestInterval          interface{} `pulumi:"requestInterval"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RestSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	AdditionalHeaders        interface{} `pulumi:"additionalHeaders"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	PaginationRules          interface{} `pulumi:"paginationRules"`
	RequestBody              interface{} `pulumi:"requestBody"`
	RequestInterval          interface{} `pulumi:"requestInterval"`
	RequestMethod            interface{} `pulumi:"requestMethod"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type RetryPolicy struct {
	Count             interface{} `pulumi:"count"`
	IntervalInSeconds *int        `pulumi:"intervalInSeconds"`
}

type RetryPolicyResponse struct {
	Count             interface{} `pulumi:"count"`
	IntervalInSeconds *int        `pulumi:"intervalInSeconds"`
}

type SSISAccessCredential struct {
	Domain   interface{} `pulumi:"domain"`
	Password interface{} `pulumi:"password"`
	UserName interface{} `pulumi:"userName"`
}

type SSISAccessCredentialResponse struct {
	Domain   interface{} `pulumi:"domain"`
	Password interface{} `pulumi:"password"`
	UserName interface{} `pulumi:"userName"`
}

type SSISChildPackage struct {
	PackageContent          interface{} `pulumi:"packageContent"`
	PackageLastModifiedDate *string     `pulumi:"packageLastModifiedDate"`
	PackageName             *string     `pulumi:"packageName"`
	PackagePath             interface{} `pulumi:"packagePath"`
}

type SSISChildPackageResponse struct {
	PackageContent          interface{} `pulumi:"packageContent"`
	PackageLastModifiedDate *string     `pulumi:"packageLastModifiedDate"`
	PackageName             *string     `pulumi:"packageName"`
	PackagePath             interface{} `pulumi:"packagePath"`
}

type SSISExecutionCredential struct {
	Domain   interface{}  `pulumi:"domain"`
	Password SecureString `pulumi:"password"`
	UserName interface{}  `pulumi:"userName"`
}

type SSISExecutionCredentialResponse struct {
	Domain   interface{}          `pulumi:"domain"`
	Password SecureStringResponse `pulumi:"password"`
	UserName interface{}          `pulumi:"userName"`
}

type SSISExecutionParameter struct {
	Value interface{} `pulumi:"value"`
}

type SSISExecutionParameterResponse struct {
	Value interface{} `pulumi:"value"`
}

type SSISLogLocation struct {
	AccessCredential   *SSISAccessCredential `pulumi:"accessCredential"`
	LogPath            interface{}           `pulumi:"logPath"`
	LogRefreshInterval interface{}           `pulumi:"logRefreshInterval"`
	Type               string                `pulumi:"type"`
}

type SSISLogLocationResponse struct {
	AccessCredential   *SSISAccessCredentialResponse `pulumi:"accessCredential"`
	LogPath            interface{}                   `pulumi:"logPath"`
	LogRefreshInterval interface{}                   `pulumi:"logRefreshInterval"`
	Type               string                        `pulumi:"type"`
}

type SSISPackageLocation struct {
	AccessCredential              *SSISAccessCredential `pulumi:"accessCredential"`
	ChildPackages                 []SSISChildPackage    `pulumi:"childPackages"`
	ConfigurationAccessCredential *SSISAccessCredential `pulumi:"configurationAccessCredential"`
	ConfigurationPath             interface{}           `pulumi:"configurationPath"`
	PackageContent                interface{}           `pulumi:"packageContent"`
	PackageLastModifiedDate       *string               `pulumi:"packageLastModifiedDate"`
	PackageName                   *string               `pulumi:"packageName"`
	PackagePassword               interface{}           `pulumi:"packagePassword"`
	PackagePath                   interface{}           `pulumi:"packagePath"`
	Type                          *string               `pulumi:"type"`
}

type SSISPackageLocationResponse struct {
	AccessCredential              *SSISAccessCredentialResponse `pulumi:"accessCredential"`
	ChildPackages                 []SSISChildPackageResponse    `pulumi:"childPackages"`
	ConfigurationAccessCredential *SSISAccessCredentialResponse `pulumi:"configurationAccessCredential"`
	ConfigurationPath             interface{}                   `pulumi:"configurationPath"`
	PackageContent                interface{}                   `pulumi:"packageContent"`
	PackageLastModifiedDate       *string                       `pulumi:"packageLastModifiedDate"`
	PackageName                   *string                       `pulumi:"packageName"`
	PackagePassword               interface{}                   `pulumi:"packagePassword"`
	PackagePath                   interface{}                   `pulumi:"packagePath"`
	Type                          *string                       `pulumi:"type"`
}

type SSISPropertyOverride struct {
	IsSensitive *bool       `pulumi:"isSensitive"`
	Value       interface{} `pulumi:"value"`
}

type SSISPropertyOverrideResponse struct {
	IsSensitive *bool       `pulumi:"isSensitive"`
	Value       interface{} `pulumi:"value"`
}

type SalesforceLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiVersion          interface{}                       `pulumi:"apiVersion"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	EnvironmentUrl      interface{}                       `pulumi:"environmentUrl"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	SecurityToken       interface{}                       `pulumi:"securityToken"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type SalesforceLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiVersion          interface{}                               `pulumi:"apiVersion"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	EnvironmentUrl      interface{}                               `pulumi:"environmentUrl"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	SecurityToken       interface{}                               `pulumi:"securityToken"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type SalesforceMarketingCloudLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                       `pulumi:"connectionProperties"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type SalesforceMarketingCloudLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionProperties  interface{}                               `pulumi:"connectionProperties"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type SalesforceMarketingCloudObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type SalesforceMarketingCloudObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type SalesforceMarketingCloudSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceMarketingCloudSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                       `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SalesforceObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                               `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SalesforceServiceCloudLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiVersion          interface{}                       `pulumi:"apiVersion"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	EnvironmentUrl      interface{}                       `pulumi:"environmentUrl"`
	ExtendedProperties  interface{}                       `pulumi:"extendedProperties"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	SecurityToken       interface{}                       `pulumi:"securityToken"`
	Type                string                            `pulumi:"type"`
	Username            interface{}                       `pulumi:"username"`
}

type SalesforceServiceCloudLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiVersion          interface{}                               `pulumi:"apiVersion"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	EnvironmentUrl      interface{}                               `pulumi:"environmentUrl"`
	ExtendedProperties  interface{}                               `pulumi:"extendedProperties"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	SecurityToken       interface{}                               `pulumi:"securityToken"`
	Type                string                                    `pulumi:"type"`
	Username            interface{}                               `pulumi:"username"`
}

type SalesforceServiceCloudObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                       `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SalesforceServiceCloudObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ObjectApiName     interface{}                               `pulumi:"objectApiName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SalesforceServiceCloudSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExternalIdFieldName      interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SalesforceServiceCloudSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExternalIdFieldName      interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SalesforceServiceCloudSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	ReadBehavior             *string     `pulumi:"readBehavior"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceServiceCloudSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	ReadBehavior             *string     `pulumi:"readBehavior"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExternalIdFieldName      interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SalesforceSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExternalIdFieldName      interface{} `pulumi:"externalIdFieldName"`
	IgnoreNullValues         interface{} `pulumi:"ignoreNullValues"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SalesforceSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	ReadBehavior             *string     `pulumi:"readBehavior"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SalesforceSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	ReadBehavior             *string     `pulumi:"readBehavior"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapBWLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ClientId            interface{}                       `pulumi:"clientId"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Server              interface{}                       `pulumi:"server"`
	SystemNumber        interface{}                       `pulumi:"systemNumber"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type SapBWLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ClientId            interface{}                               `pulumi:"clientId"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Server              interface{}                               `pulumi:"server"`
	SystemNumber        interface{}                               `pulumi:"systemNumber"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type SapBwCubeDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapBwCubeDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapBwSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapBwSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapCloudForCustomerLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 interface{}                       `pulumi:"url"`
	Username            interface{}                       `pulumi:"username"`
}

type SapCloudForCustomerLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 interface{}                               `pulumi:"url"`
	Username            interface{}                               `pulumi:"username"`
}

type SapCloudForCustomerResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapCloudForCustomerResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapCloudForCustomerSink struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SapCloudForCustomerSinkResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	SinkRetryCount           interface{} `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{} `pulumi:"sinkRetryWait"`
	Type                     string      `pulumi:"type"`
	WriteBatchSize           interface{} `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{} `pulumi:"writeBatchTimeout"`
	WriteBehavior            *string     `pulumi:"writeBehavior"`
}

type SapCloudForCustomerSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapCloudForCustomerSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapEccLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential *string                           `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Type                string                            `pulumi:"type"`
	Url                 string                            `pulumi:"url"`
	Username            *string                           `pulumi:"username"`
}

type SapEccLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential *string                                   `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
	Url                 string                                    `pulumi:"url"`
	Username            *string                                   `pulumi:"username"`
}

type SapEccResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Path              interface{}                       `pulumi:"path"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapEccResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Path              interface{}                               `pulumi:"path"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapEccSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapEccSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SapHanaLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	AuthenticationType  *string                           `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            interface{}                       `pulumi:"password"`
	Server              interface{}                       `pulumi:"server"`
	Type                string                            `pulumi:"type"`
	UserName            interface{}                       `pulumi:"userName"`
}

type SapHanaLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	AuthenticationType  *string                                   `pulumi:"authenticationType"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            interface{}                               `pulumi:"password"`
	Server              interface{}                               `pulumi:"server"`
	Type                string                                    `pulumi:"type"`
	UserName            interface{}                               `pulumi:"userName"`
}

type SapHanaPartitionSettings struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
}

type SapHanaPartitionSettingsResponse struct {
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
}

type SapHanaSource struct {
	AdditionalColumns        interface{}               `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}               `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}               `pulumi:"maxConcurrentConnections"`
	PacketSize               interface{}               `pulumi:"packetSize"`
	PartitionOption          interface{}               `pulumi:"partitionOption"`
	PartitionSettings        *SapHanaPartitionSettings `pulumi:"partitionSettings"`
	Query                    interface{}               `pulumi:"query"`
	QueryTimeout             interface{}               `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}               `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}               `pulumi:"sourceRetryWait"`
	Type                     string                    `pulumi:"type"`
}

type SapHanaSourceResponse struct {
	AdditionalColumns        interface{}                       `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{}                       `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                       `pulumi:"maxConcurrentConnections"`
	PacketSize               interface{}                       `pulumi:"packetSize"`
	PartitionOption          interface{}                       `pulumi:"partitionOption"`
	PartitionSettings        *SapHanaPartitionSettingsResponse `pulumi:"partitionSettings"`
	Query                    interface{}                       `pulumi:"query"`
	QueryTimeout             interface{}                       `pulumi:"queryTimeout"`
	SourceRetryCount         interface{}                       `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                       `pulumi:"sourceRetryWait"`
	Type                     string                            `pulumi:"type"`
}

type SapHanaTableDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	Type              string                            `pulumi:"type"`
}

type SapHanaTableDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	Type              string                                    `pulumi:"type"`
}

type SapOdpLinkedService struct {
	Annotations          []interface{}                     `pulumi:"annotations"`
	ClientId             interface{}                       `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description          *string                           `pulumi:"description"`
	EncryptedCredential  interface{}                       `pulumi:"encryptedCredential"`
	Language             interface{}                       `pulumi:"language"`
	LogonGroup           interface{}                       `pulumi:"logonGroup"`
	MessageServer        interface{}                       `pulumi:"messageServer"`
	MessageServerService interface{}                       `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	Password             interface{}                       `pulumi:"password"`
	Server               interface{}                       `pulumi:"server"`
	SncLibraryPath       interface{}                       `pulumi:"sncLibraryPath"`
	SncMode              interface{}                       `pulumi:"sncMode"`
	SncMyName            interface{}                       `pulumi:"sncMyName"`
	SncPartnerName       interface{}                       `pulumi:"sncPartnerName"`
	SncQop               interface{}                       `pulumi:"sncQop"`
	SubscriberName       interface{}                       `pulumi:"subscriberName"`
	SystemId             interface{}                       `pulumi:"systemId"`
	SystemNumber         interface{}                       `pulumi:"systemNumber"`
	Type                 string                            `pulumi:"type"`
	UserName             interface{}                       `pulumi:"userName"`
	X509CertificatePath  interface{}                       `pulumi:"x509CertificatePath"`
}

type SapOdpLinkedServiceResponse struct {
	Annotations          []interface{}                             `pulumi:"annotations"`
	ClientId             interface{}                               `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description          *string                                   `pulumi:"description"`
	EncryptedCredential  interface{}                               `pulumi:"encryptedCredential"`
	Language             interface{}                               `pulumi:"language"`
	LogonGroup           interface{}                               `pulumi:"logonGroup"`
	MessageServer        interface{}                               `pulumi:"messageServer"`
	MessageServerService interface{}                               `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password             interface{}                               `pulumi:"password"`
	Server               interface{}                               `pulumi:"server"`
	SncLibraryPath       interface{}                               `pulumi:"sncLibraryPath"`
	SncMode              interface{}                               `pulumi:"sncMode"`
	SncMyName            interface{}                               `pulumi:"sncMyName"`
	SncPartnerName       interface{}                               `pulumi:"sncPartnerName"`
	SncQop               interface{}                               `pulumi:"sncQop"`
	SubscriberName       interface{}                               `pulumi:"subscriberName"`
	SystemId             interface{}                               `pulumi:"systemId"`
	SystemNumber         interface{}                               `pulumi:"systemNumber"`
	Type                 string                                    `pulumi:"type"`
	UserName             interface{}                               `pulumi:"userName"`
	X509CertificatePath  interface{}                               `pulumi:"x509CertificatePath"`
}

type SapOdpResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Context           interface{}                       `pulumi:"context"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ObjectName        interface{}                       `pulumi:"objectName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SapOdpResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Context           interface{}                               `pulumi:"context"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ObjectName        interface{}                               `pulumi:"objectName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SapOdpSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExtractionMode           interface{} `pulumi:"extractionMode"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Projection               interface{} `pulumi:"projection"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	Selection                interface{} `pulumi:"selection"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	SubscriberProcess        interface{} `pulumi:"subscriberProcess"`
	Type                     string      `pulumi:"type"`
}

type SapOdpSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	ExtractionMode           interface{} `pulumi:"extractionMode"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Projection               interface{} `pulumi:"projection"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	Selection                interface{} `pulumi:"selection"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	SubscriberProcess        interface{} `pulumi:"subscriberProcess"`
	Type                     string      `pulumi:"type"`
}

type SapOpenHubLinkedService struct {
	Annotations          []interface{}                     `pulumi:"annotations"`
	ClientId             interface{}                       `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description          *string                           `pulumi:"description"`
	EncryptedCredential  interface{}                       `pulumi:"encryptedCredential"`
	Language             interface{}                       `pulumi:"language"`
	LogonGroup           interface{}                       `pulumi:"logonGroup"`
	MessageServer        interface{}                       `pulumi:"messageServer"`
	MessageServerService interface{}                       `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	Password             interface{}                       `pulumi:"password"`
	Server               interface{}                       `pulumi:"server"`
	SystemId             interface{}                       `pulumi:"systemId"`
	SystemNumber         interface{}                       `pulumi:"systemNumber"`
	Type                 string                            `pulumi:"type"`
	UserName             interface{}                       `pulumi:"userName"`
}

type SapOpenHubLinkedServiceResponse struct {
	Annotations          []interface{}                             `pulumi:"annotations"`
	ClientId             interface{}                               `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description          *string                                   `pulumi:"description"`
	EncryptedCredential  interface{}                               `pulumi:"encryptedCredential"`
	Language             interface{}                               `pulumi:"language"`
	LogonGroup           interface{}                               `pulumi:"logonGroup"`
	MessageServer        interface{}                               `pulumi:"messageServer"`
	MessageServerService interface{}                               `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password             interface{}                               `pulumi:"password"`
	Server               interface{}                               `pulumi:"server"`
	SystemId             interface{}                               `pulumi:"systemId"`
	SystemNumber         interface{}                               `pulumi:"systemNumber"`
	Type                 string                                    `pulumi:"type"`
	UserName             interface{}                               `pulumi:"userName"`
}

type SapOpenHubSource struct {
	AdditionalColumns                interface{} `pulumi:"additionalColumns"`
	BaseRequestId                    interface{} `pulumi:"baseRequestId"`
	CustomRfcReadTableFunctionModule interface{} `pulumi:"customRfcReadTableFunctionModule"`
	DisableMetricsCollection         interface{} `pulumi:"disableMetricsCollection"`
	ExcludeLastRequest               interface{} `pulumi:"excludeLastRequest"`
	MaxConcurrentConnections         interface{} `pulumi:"maxConcurrentConnections"`
	QueryTimeout                     interface{} `pulumi:"queryTimeout"`
	SapDataColumnDelimiter           interface{} `pulumi:"sapDataColumnDelimiter"`
	SourceRetryCount                 interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait                  interface{} `pulumi:"sourceRetryWait"`
	Type                             string      `pulumi:"type"`
}

type SapOpenHubSourceResponse struct {
	AdditionalColumns                interface{} `pulumi:"additionalColumns"`
	BaseRequestId                    interface{} `pulumi:"baseRequestId"`
	CustomRfcReadTableFunctionModule interface{} `pulumi:"customRfcReadTableFunctionModule"`
	DisableMetricsCollection         interface{} `pulumi:"disableMetricsCollection"`
	ExcludeLastRequest               interface{} `pulumi:"excludeLastRequest"`
	MaxConcurrentConnections         interface{} `pulumi:"maxConcurrentConnections"`
	QueryTimeout                     interface{} `pulumi:"queryTimeout"`
	SapDataColumnDelimiter           interface{} `pulumi:"sapDataColumnDelimiter"`
	SourceRetryCount                 interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait                  interface{} `pulumi:"sourceRetryWait"`
	Type                             string      `pulumi:"type"`
}

type SapOpenHubTableDataset struct {
	Annotations            []interface{}                     `pulumi:"annotations"`
	BaseRequestId          interface{}                       `pulumi:"baseRequestId"`
	Description            *string                           `pulumi:"description"`
	ExcludeLastRequest     interface{}                       `pulumi:"excludeLastRequest"`
	Folder                 *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName      LinkedServiceReference            `pulumi:"linkedServiceName"`
	OpenHubDestinationName interface{}                       `pulumi:"openHubDestinationName"`
	Parameters             map[string]ParameterSpecification `pulumi:"parameters"`
	Schema                 interface{}                       `pulumi:"schema"`
	Structure              interface{}                       `pulumi:"structure"`
	Type                   string                            `pulumi:"type"`
}

type SapOpenHubTableDatasetResponse struct {
	Annotations            []interface{}                             `pulumi:"annotations"`
	BaseRequestId          interface{}                               `pulumi:"baseRequestId"`
	Description            *string                                   `pulumi:"description"`
	ExcludeLastRequest     interface{}                               `pulumi:"excludeLastRequest"`
	Folder                 *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName      LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	OpenHubDestinationName interface{}                               `pulumi:"openHubDestinationName"`
	Parameters             map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema                 interface{}                               `pulumi:"schema"`
	Structure              interface{}                               `pulumi:"structure"`
	Type                   string                                    `pulumi:"type"`
}

type SapTableLinkedService struct {
	Annotations          []interface{}                     `pulumi:"annotations"`
	ClientId             interface{}                       `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description          *string                           `pulumi:"description"`
	EncryptedCredential  interface{}                       `pulumi:"encryptedCredential"`
	Language             interface{}                       `pulumi:"language"`
	LogonGroup           interface{}                       `pulumi:"logonGroup"`
	MessageServer        interface{}                       `pulumi:"messageServer"`
	MessageServerService interface{}                       `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecification `pulumi:"parameters"`
	Password             interface{}                       `pulumi:"password"`
	Server               interface{}                       `pulumi:"server"`
	SncLibraryPath       interface{}                       `pulumi:"sncLibraryPath"`
	SncMode              interface{}                       `pulumi:"sncMode"`
	SncMyName            interface{}                       `pulumi:"sncMyName"`
	SncPartnerName       interface{}                       `pulumi:"sncPartnerName"`
	SncQop               interface{}                       `pulumi:"sncQop"`
	SystemId             interface{}                       `pulumi:"systemId"`
	SystemNumber         interface{}                       `pulumi:"systemNumber"`
	Type                 string                            `pulumi:"type"`
	UserName             interface{}                       `pulumi:"userName"`
}

type SapTableLinkedServiceResponse struct {
	Annotations          []interface{}                             `pulumi:"annotations"`
	ClientId             interface{}                               `pulumi:"clientId"`
	ConnectVia           *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description          *string                                   `pulumi:"description"`
	EncryptedCredential  interface{}                               `pulumi:"encryptedCredential"`
	Language             interface{}                               `pulumi:"language"`
	LogonGroup           interface{}                               `pulumi:"logonGroup"`
	MessageServer        interface{}                               `pulumi:"messageServer"`
	MessageServerService interface{}                               `pulumi:"messageServerService"`
	Parameters           map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password             interface{}                               `pulumi:"password"`
	Server               interface{}                               `pulumi:"server"`
	SncLibraryPath       interface{}                               `pulumi:"sncLibraryPath"`
	SncMode              interface{}                               `pulumi:"sncMode"`
	SncMyName            interface{}                               `pulumi:"sncMyName"`
	SncPartnerName       interface{}                               `pulumi:"sncPartnerName"`
	SncQop               interface{}                               `pulumi:"sncQop"`
	SystemId             interface{}                               `pulumi:"systemId"`
	SystemNumber         interface{}                               `pulumi:"systemNumber"`
	Type                 string                                    `pulumi:"type"`
	UserName             interface{}                               `pulumi:"userName"`
}

type SapTablePartitionSettings struct {
	MaxPartitionsNumber interface{} `pulumi:"maxPartitionsNumber"`
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type SapTablePartitionSettingsResponse struct {
	MaxPartitionsNumber interface{} `pulumi:"maxPartitionsNumber"`
	PartitionColumnName interface{} `pulumi:"partitionColumnName"`
	PartitionLowerBound interface{} `pulumi:"partitionLowerBound"`
	PartitionUpperBound interface{} `pulumi:"partitionUpperBound"`
}

type SapTableResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type SapTableResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type SapTableSource struct {
	AdditionalColumns                interface{}                `pulumi:"additionalColumns"`
	BatchSize                        interface{}                `pulumi:"batchSize"`
	CustomRfcReadTableFunctionModule interface{}                `pulumi:"customRfcReadTableFunctionModule"`
	DisableMetricsCollection         interface{}                `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections         interface{}                `pulumi:"maxConcurrentConnections"`
	PartitionOption                  interface{}                `pulumi:"partitionOption"`
	PartitionSettings                *SapTablePartitionSettings `pulumi:"partitionSettings"`
	QueryTimeout                     interface{}                `pulumi:"queryTimeout"`
	RfcTableFields                   interface{}                `pulumi:"rfcTableFields"`
	RfcTableOptions                  interface{}                `pulumi:"rfcTableOptions"`
	RowCount                         interface{}                `pulumi:"rowCount"`
	RowSkips                         interface{}                `pulumi:"rowSkips"`
	SapDataColumnDelimiter           interface{}                `pulumi:"sapDataColumnDelimiter"`
	SourceRetryCount                 interface{}                `pulumi:"sourceRetryCount"`
	SourceRetryWait                  interface{}                `pulumi:"sourceRetryWait"`
	Type                             string                     `pulumi:"type"`
}

type SapTableSourceResponse struct {
	AdditionalColumns                interface{}                        `pulumi:"additionalColumns"`
	BatchSize                        interface{}                        `pulumi:"batchSize"`
	CustomRfcReadTableFunctionModule interface{}                        `pulumi:"customRfcReadTableFunctionModule"`
	DisableMetricsCollection         interface{}                        `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections         interface{}                        `pulumi:"maxConcurrentConnections"`
	PartitionOption                  interface{}                        `pulumi:"partitionOption"`
	PartitionSettings                *SapTablePartitionSettingsResponse `pulumi:"partitionSettings"`
	QueryTimeout                     interface{}                        `pulumi:"queryTimeout"`
	RfcTableFields                   interface{}                        `pulumi:"rfcTableFields"`
	RfcTableOptions                  interface{}                        `pulumi:"rfcTableOptions"`
	RowCount                         interface{}                        `pulumi:"rowCount"`
	RowSkips                         interface{}                        `pulumi:"rowSkips"`
	SapDataColumnDelimiter           interface{}                        `pulumi:"sapDataColumnDelimiter"`
	SourceRetryCount                 interface{}                        `pulumi:"sourceRetryCount"`
	SourceRetryWait                  interface{}                        `pulumi:"sourceRetryWait"`
	Type                             string                             `pulumi:"type"`
}

type ScheduleTrigger struct {
	Annotations []interface{}              `pulumi:"annotations"`
	Description *string                    `pulumi:"description"`
	Pipelines   []TriggerPipelineReference `pulumi:"pipelines"`
	Recurrence  ScheduleTriggerRecurrence  `pulumi:"recurrence"`
	Type        string                     `pulumi:"type"`
}

type ScheduleTriggerRecurrence struct {
	EndTime   *string             `pulumi:"endTime"`
	Frequency *string             `pulumi:"frequency"`
	Interval  *int                `pulumi:"interval"`
	Schedule  *RecurrenceSchedule `pulumi:"schedule"`
	StartTime *string             `pulumi:"startTime"`
	TimeZone  *string             `pulumi:"timeZone"`
}

type ScheduleTriggerRecurrenceResponse struct {
	EndTime   *string                     `pulumi:"endTime"`
	Frequency *string                     `pulumi:"frequency"`
	Interval  *int                        `pulumi:"interval"`
	Schedule  *RecurrenceScheduleResponse `pulumi:"schedule"`
	StartTime *string                     `pulumi:"startTime"`
	TimeZone  *string                     `pulumi:"timeZone"`
}

type ScheduleTriggerResponse struct {
	Annotations  []interface{}                      `pulumi:"annotations"`
	Description  *string                            `pulumi:"description"`
	Pipelines    []TriggerPipelineReferenceResponse `pulumi:"pipelines"`
	Recurrence   ScheduleTriggerRecurrenceResponse  `pulumi:"recurrence"`
	RuntimeState string                             `pulumi:"runtimeState"`
	Type         string                             `pulumi:"type"`
}

type ScriptAction struct {
	Name       string      `pulumi:"name"`
	Parameters *string     `pulumi:"parameters"`
	Roles      interface{} `pulumi:"roles"`
	Uri        string      `pulumi:"uri"`
}

type ScriptActionResponse struct {
	Name       string      `pulumi:"name"`
	Parameters *string     `pulumi:"parameters"`
	Roles      interface{} `pulumi:"roles"`
	Uri        string      `pulumi:"uri"`
}

type ScriptActivity struct {
	DependsOn                   []ActivityDependency                     `pulumi:"dependsOn"`
	Description                 *string                                  `pulumi:"description"`
	LinkedServiceName           LinkedServiceReference                   `pulumi:"linkedServiceName"`
	LogSettings                 *ScriptActivityTypePropertiesLogSettings `pulumi:"logSettings"`
	Name                        string                                   `pulumi:"name"`
	Policy                      *ActivityPolicy                          `pulumi:"policy"`
	ScriptBlockExecutionTimeout interface{}                              `pulumi:"scriptBlockExecutionTimeout"`
	Scripts                     []ScriptActivityScriptBlock              `pulumi:"scripts"`
	Type                        string                                   `pulumi:"type"`
	UserProperties              []UserProperty                           `pulumi:"userProperties"`
}

type ScriptActivityParameter struct {
	Direction *string     `pulumi:"direction"`
	Name      interface{} `pulumi:"name"`
	Size      *int        `pulumi:"size"`
	Type      *string     `pulumi:"type"`
	Value     interface{} `pulumi:"value"`
}

type ScriptActivityParameterResponse struct {
	Direction *string     `pulumi:"direction"`
	Name      interface{} `pulumi:"name"`
	Size      *int        `pulumi:"size"`
	Type      *string     `pulumi:"type"`
	Value     interface{} `pulumi:"value"`
}

type ScriptActivityResponse struct {
	DependsOn                   []ActivityDependencyResponse                     `pulumi:"dependsOn"`
	Description                 *string                                          `pulumi:"description"`
	LinkedServiceName           LinkedServiceReferenceResponse                   `pulumi:"linkedServiceName"`
	LogSettings                 *ScriptActivityTypePropertiesResponseLogSettings `pulumi:"logSettings"`
	Name                        string                                           `pulumi:"name"`
	Policy                      *ActivityPolicyResponse                          `pulumi:"policy"`
	ScriptBlockExecutionTimeout interface{}                                      `pulumi:"scriptBlockExecutionTimeout"`
	Scripts                     []ScriptActivityScriptBlockResponse              `pulumi:"scripts"`
	Type                        string                                           `pulumi:"type"`
	UserProperties              []UserPropertyResponse                           `pulumi:"userProperties"`
}

type ScriptActivityScriptBlock struct {
	Parameters []ScriptActivityParameter `pulumi:"parameters"`
	Text       interface{}               `pulumi:"text"`
	Type       string                    `pulumi:"type"`
}

type ScriptActivityScriptBlockResponse struct {
	Parameters []ScriptActivityParameterResponse `pulumi:"parameters"`
	Text       interface{}                       `pulumi:"text"`
	Type       string                            `pulumi:"type"`
}

type ScriptActivityTypePropertiesLogSettings struct {
	LogDestination      string               `pulumi:"logDestination"`
	LogLocationSettings *LogLocationSettings `pulumi:"logLocationSettings"`
}

type ScriptActivityTypePropertiesResponseLogSettings struct {
	LogDestination      string                       `pulumi:"logDestination"`
	LogLocationSettings *LogLocationSettingsResponse `pulumi:"logLocationSettings"`
}

type SecureString struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type SecureStringResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type SelfDependencyTumblingWindowTriggerReference struct {
	Offset string  `pulumi:"offset"`
	Size   *string `pulumi:"size"`
	Type   string  `pulumi:"type"`
}

type SelfDependencyTumblingWindowTriggerReferenceResponse struct {
	Offset string  `pulumi:"offset"`
	Size   *string `pulumi:"size"`
	Type   string  `pulumi:"type"`
}

type SelfHostedIntegrationRuntime struct {
	Description *string     `pulumi:"description"`
	LinkedInfo  interface{} `pulumi:"linkedInfo"`
	Type        string      `pulumi:"type"`
}

type SelfHostedIntegrationRuntimeNodeResponse struct {
	Capabilities        map[string]string `pulumi:"capabilities"`
	ConcurrentJobsLimit int               `pulumi:"concurrentJobsLimit"`
	ExpiryTime          string            `pulumi:"expiryTime"`
	HostServiceUri      string            `pulumi:"hostServiceUri"`
	IsActiveDispatcher  bool              `pulumi:"isActiveDispatcher"`
	LastConnectTime     string            `pulumi:"lastConnectTime"`
	LastEndUpdateTime   string            `pulumi:"lastEndUpdateTime"`
	LastStartTime       string            `pulumi:"lastStartTime"`
	LastStartUpdateTime string            `pulumi:"lastStartUpdateTime"`
	LastStopTime        string            `pulumi:"lastStopTime"`
	LastUpdateResult    string            `pulumi:"lastUpdateResult"`
	MachineName         string            `pulumi:"machineName"`
	MaxConcurrentJobs   int               `pulumi:"maxConcurrentJobs"`
	NodeName            string            `pulumi:"nodeName"`
	RegisterTime        string            `pulumi:"registerTime"`
	Status              string            `pulumi:"status"`
	Version             string            `pulumi:"version"`
	VersionStatus       string            `pulumi:"versionStatus"`
}

type SelfHostedIntegrationRuntimeResponse struct {
	Description *string     `pulumi:"description"`
	LinkedInfo  interface{} `pulumi:"linkedInfo"`
	Type        string      `pulumi:"type"`
}

type SelfHostedIntegrationRuntimeStatusResponse struct {
	AutoUpdate                string                                     `pulumi:"autoUpdate"`
	AutoUpdateETA             string                                     `pulumi:"autoUpdateETA"`
	Capabilities              map[string]string                          `pulumi:"capabilities"`
	CreateTime                string                                     `pulumi:"createTime"`
	DataFactoryName           string                                     `pulumi:"dataFactoryName"`
	InternalChannelEncryption string                                     `pulumi:"internalChannelEncryption"`
	LatestVersion             string                                     `pulumi:"latestVersion"`
	Links                     []LinkedIntegrationRuntimeResponse         `pulumi:"links"`
	LocalTimeZoneOffset       string                                     `pulumi:"localTimeZoneOffset"`
	Nodes                     []SelfHostedIntegrationRuntimeNodeResponse `pulumi:"nodes"`
	PushedVersion             string                                     `pulumi:"pushedVersion"`
	ScheduledUpdateDate       string                                     `pulumi:"scheduledUpdateDate"`
	ServiceUrls               []string                                   `pulumi:"serviceUrls"`
	State                     string                                     `pulumi:"state"`
	TaskQueueId               string                                     `pulumi:"taskQueueId"`
	Type                      string                                     `pulumi:"type"`
	UpdateDelayOffset         string                                     `pulumi:"updateDelayOffset"`
	Version                   string                                     `pulumi:"version"`
	VersionStatus             string                                     `pulumi:"versionStatus"`
}

type ServiceNowLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	AuthenticationType    string                            `pulumi:"authenticationType"`
	ClientId              interface{}                       `pulumi:"clientId"`
	ClientSecret          interface{}                       `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Endpoint              interface{}                       `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Password              interface{}                       `pulumi:"password"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
	Username              interface{}                       `pulumi:"username"`
}

type ServiceNowLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	AuthenticationType    string                                    `pulumi:"authenticationType"`
	ClientId              interface{}                               `pulumi:"clientId"`
	ClientSecret          interface{}                               `pulumi:"clientSecret"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Endpoint              interface{}                               `pulumi:"endpoint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password              interface{}                               `pulumi:"password"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
	Username              interface{}                               `pulumi:"username"`
}

type ServiceNowObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type ServiceNowObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type ServiceNowSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ServiceNowSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SetVariableActivity struct {
	DependsOn      []ActivityDependency `pulumi:"dependsOn"`
	Description    *string              `pulumi:"description"`
	Name           string               `pulumi:"name"`
	Type           string               `pulumi:"type"`
	UserProperties []UserProperty       `pulumi:"userProperties"`
	Value          interface{}          `pulumi:"value"`
	VariableName   *string              `pulumi:"variableName"`
}

type SetVariableActivityResponse struct {
	DependsOn      []ActivityDependencyResponse `pulumi:"dependsOn"`
	Description    *string                      `pulumi:"description"`
	Name           string                       `pulumi:"name"`
	Type           string                       `pulumi:"type"`
	UserProperties []UserPropertyResponse       `pulumi:"userProperties"`
	Value          interface{}                  `pulumi:"value"`
	VariableName   *string                      `pulumi:"variableName"`
}

type SftpLocation struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type SftpLocationResponse struct {
	FileName   interface{} `pulumi:"fileName"`
	FolderPath interface{} `pulumi:"folderPath"`
	Type       string      `pulumi:"type"`
}

type SftpReadSettings struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableChunking            interface{} `pulumi:"disableChunking"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type SftpReadSettingsResponse struct {
	DeleteFilesAfterCompletion interface{} `pulumi:"deleteFilesAfterCompletion"`
	DisableChunking            interface{} `pulumi:"disableChunking"`
	DisableMetricsCollection   interface{} `pulumi:"disableMetricsCollection"`
	EnablePartitionDiscovery   *bool       `pulumi:"enablePartitionDiscovery"`
	FileListPath               interface{} `pulumi:"fileListPath"`
	MaxConcurrentConnections   interface{} `pulumi:"maxConcurrentConnections"`
	ModifiedDatetimeEnd        interface{} `pulumi:"modifiedDatetimeEnd"`
	ModifiedDatetimeStart      interface{} `pulumi:"modifiedDatetimeStart"`
	PartitionRootPath          interface{} `pulumi:"partitionRootPath"`
	Recursive                  interface{} `pulumi:"recursive"`
	Type                       string      `pulumi:"type"`
	WildcardFileName           interface{} `pulumi:"wildcardFileName"`
	WildcardFolderPath         interface{} `pulumi:"wildcardFolderPath"`
}

type SftpServerLinkedService struct {
	Annotations           []interface{}                     `pulumi:"annotations"`
	AuthenticationType    *string                           `pulumi:"authenticationType"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	HostKeyFingerprint    interface{}                       `pulumi:"hostKeyFingerprint"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	PassPhrase            interface{}                       `pulumi:"passPhrase"`
	Password              interface{}                       `pulumi:"password"`
	Port                  interface{}                       `pulumi:"port"`
	PrivateKeyContent     interface{}                       `pulumi:"privateKeyContent"`
	PrivateKeyPath        interface{}                       `pulumi:"privateKeyPath"`
	SkipHostKeyValidation interface{}                       `pulumi:"skipHostKeyValidation"`
	Type                  string                            `pulumi:"type"`
	UserName              interface{}                       `pulumi:"userName"`
}

type SftpServerLinkedServiceResponse struct {
	Annotations           []interface{}                             `pulumi:"annotations"`
	AuthenticationType    *string                                   `pulumi:"authenticationType"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	HostKeyFingerprint    interface{}                               `pulumi:"hostKeyFingerprint"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	PassPhrase            interface{}                               `pulumi:"passPhrase"`
	Password              interface{}                               `pulumi:"password"`
	Port                  interface{}                               `pulumi:"port"`
	PrivateKeyContent     interface{}                               `pulumi:"privateKeyContent"`
	PrivateKeyPath        interface{}                               `pulumi:"privateKeyPath"`
	SkipHostKeyValidation interface{}                               `pulumi:"skipHostKeyValidation"`
	Type                  string                                    `pulumi:"type"`
	UserName              interface{}                               `pulumi:"userName"`
}

type SftpWriteSettings struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	OperationTimeout         interface{} `pulumi:"operationTimeout"`
	Type                     string      `pulumi:"type"`
	UseTempFileRename        interface{} `pulumi:"useTempFileRename"`
}

type SftpWriteSettingsResponse struct {
	CopyBehavior             interface{} `pulumi:"copyBehavior"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	OperationTimeout         interface{} `pulumi:"operationTimeout"`
	Type                     string      `pulumi:"type"`
	UseTempFileRename        interface{} `pulumi:"useTempFileRename"`
}

type SharePointOnlineListLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	ServicePrincipalId  interface{}                       `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                       `pulumi:"servicePrincipalKey"`
	SiteUrl             interface{}                       `pulumi:"siteUrl"`
	TenantId            interface{}                       `pulumi:"tenantId"`
	Type                string                            `pulumi:"type"`
}

type SharePointOnlineListLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	ServicePrincipalId  interface{}                               `pulumi:"servicePrincipalId"`
	ServicePrincipalKey interface{}                               `pulumi:"servicePrincipalKey"`
	SiteUrl             interface{}                               `pulumi:"siteUrl"`
	TenantId            interface{}                               `pulumi:"tenantId"`
	Type                string                                    `pulumi:"type"`
}

type SharePointOnlineListResourceDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	ListName          interface{}                       `pulumi:"listName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Type              string                            `pulumi:"type"`
}

type SharePointOnlineListResourceDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	ListName          interface{}                               `pulumi:"listName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Type              string                                    `pulumi:"type"`
}

type SharePointOnlineListSource struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SharePointOnlineListSourceResponse struct {
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	HttpRequestTimeout       interface{} `pulumi:"httpRequestTimeout"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ShopifyLinkedService struct {
	AccessToken           interface{}                       `pulumi:"accessToken"`
	Annotations           []interface{}                     `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description           *string                           `pulumi:"description"`
	EncryptedCredential   interface{}                       `pulumi:"encryptedCredential"`
	Host                  interface{}                       `pulumi:"host"`
	Parameters            map[string]ParameterSpecification `pulumi:"parameters"`
	Type                  string                            `pulumi:"type"`
	UseEncryptedEndpoints interface{}                       `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                       `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                       `pulumi:"usePeerVerification"`
}

type ShopifyLinkedServiceResponse struct {
	AccessToken           interface{}                               `pulumi:"accessToken"`
	Annotations           []interface{}                             `pulumi:"annotations"`
	ConnectVia            *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description           *string                                   `pulumi:"description"`
	EncryptedCredential   interface{}                               `pulumi:"encryptedCredential"`
	Host                  interface{}                               `pulumi:"host"`
	Parameters            map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                  string                                    `pulumi:"type"`
	UseEncryptedEndpoints interface{}                               `pulumi:"useEncryptedEndpoints"`
	UseHostVerification   interface{}                               `pulumi:"useHostVerification"`
	UsePeerVerification   interface{}                               `pulumi:"usePeerVerification"`
}

type ShopifyObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type ShopifyObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type ShopifySource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type ShopifySourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SkipErrorFile struct {
	DataInconsistency interface{} `pulumi:"dataInconsistency"`
	FileMissing       interface{} `pulumi:"fileMissing"`
}

type SkipErrorFileResponse struct {
	DataInconsistency interface{} `pulumi:"dataInconsistency"`
	FileMissing       interface{} `pulumi:"fileMissing"`
}

type SmartsheetLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ApiToken            interface{}                       `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Type                string                            `pulumi:"type"`
}

type SmartsheetLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ApiToken            interface{}                               `pulumi:"apiToken"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Type                string                                    `pulumi:"type"`
}

type SnowflakeDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	Type              string                            `pulumi:"type"`
}

type SnowflakeDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	Type              string                                    `pulumi:"type"`
}

type SnowflakeExportCopyCommand struct {
	AdditionalCopyOptions   map[string]interface{} `pulumi:"additionalCopyOptions"`
	AdditionalFormatOptions map[string]interface{} `pulumi:"additionalFormatOptions"`
	Type                    string                 `pulumi:"type"`
}

type SnowflakeExportCopyCommandResponse struct {
	AdditionalCopyOptions   map[string]interface{} `pulumi:"additionalCopyOptions"`
	AdditionalFormatOptions map[string]interface{} `pulumi:"additionalFormatOptions"`
	Type                    string                 `pulumi:"type"`
}

type SnowflakeImportCopyCommand struct {
	AdditionalCopyOptions   map[string]interface{} `pulumi:"additionalCopyOptions"`
	AdditionalFormatOptions map[string]interface{} `pulumi:"additionalFormatOptions"`
	Type                    string                 `pulumi:"type"`
}

type SnowflakeImportCopyCommandResponse struct {
	AdditionalCopyOptions   map[string]interface{} `pulumi:"additionalCopyOptions"`
	AdditionalFormatOptions map[string]interface{} `pulumi:"additionalFormatOptions"`
	Type                    string                 `pulumi:"type"`
}

type SnowflakeLinkedService struct {
	Annotations         []interface{}                     `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReference      `pulumi:"connectVia"`
	ConnectionString    interface{}                       `pulumi:"connectionString"`
	Description         *string                           `pulumi:"description"`
	EncryptedCredential interface{}                       `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecification `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReference     `pulumi:"password"`
	Type                string                            `pulumi:"type"`
}

type SnowflakeLinkedServiceResponse struct {
	Annotations         []interface{}                             `pulumi:"annotations"`
	ConnectVia          *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	ConnectionString    interface{}                               `pulumi:"connectionString"`
	Description         *string                                   `pulumi:"description"`
	EncryptedCredential interface{}                               `pulumi:"encryptedCredential"`
	Parameters          map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password            *AzureKeyVaultSecretReferenceResponse     `pulumi:"password"`
	Type                string                                    `pulumi:"type"`
}

type SnowflakeSink struct {
	DisableMetricsCollection interface{}                 `pulumi:"disableMetricsCollection"`
	ImportSettings           *SnowflakeImportCopyCommand `pulumi:"importSettings"`
	MaxConcurrentConnections interface{}                 `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{}                 `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}                 `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                 `pulumi:"sinkRetryWait"`
	Type                     string                      `pulumi:"type"`
	WriteBatchSize           interface{}                 `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                 `pulumi:"writeBatchTimeout"`
}

type SnowflakeSinkResponse struct {
	DisableMetricsCollection interface{}                         `pulumi:"disableMetricsCollection"`
	ImportSettings           *SnowflakeImportCopyCommandResponse `pulumi:"importSettings"`
	MaxConcurrentConnections interface{}                         `pulumi:"maxConcurrentConnections"`
	PreCopyScript            interface{}                         `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}                         `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                         `pulumi:"sinkRetryWait"`
	Type                     string                              `pulumi:"type"`
	WriteBatchSize           interface{}                         `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                         `pulumi:"writeBatchTimeout"`
}

type SnowflakeSource struct {
	DisableMetricsCollection interface{}                 `pulumi:"disableMetricsCollection"`
	ExportSettings           *SnowflakeExportCopyCommand `pulumi:"exportSettings"`
	MaxConcurrentConnections interface{}                 `pulumi:"maxConcurrentConnections"`
	Query                    interface{}                 `pulumi:"query"`
	SourceRetryCount         interface{}                 `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                 `pulumi:"sourceRetryWait"`
	Type                     string                      `pulumi:"type"`
}

type SnowflakeSourceResponse struct {
	DisableMetricsCollection interface{}                         `pulumi:"disableMetricsCollection"`
	ExportSettings           *SnowflakeExportCopyCommandResponse `pulumi:"exportSettings"`
	MaxConcurrentConnections interface{}                         `pulumi:"maxConcurrentConnections"`
	Query                    interface{}                         `pulumi:"query"`
	SourceRetryCount         interface{}                         `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{}                         `pulumi:"sourceRetryWait"`
	Type                     string                              `pulumi:"type"`
}

type SparkLinkedService struct {
	AllowHostNameCNMismatch   interface{}                       `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                       `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                     `pulumi:"annotations"`
	AuthenticationType        string                            `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReference      `pulumi:"connectVia"`
	Description               *string                           `pulumi:"description"`
	EnableSsl                 interface{}                       `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                       `pulumi:"encryptedCredential"`
	Host                      interface{}                       `pulumi:"host"`
	HttpPath                  interface{}                       `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecification `pulumi:"parameters"`
	Password                  interface{}                       `pulumi:"password"`
	Port                      interface{}                       `pulumi:"port"`
	ServerType                *string                           `pulumi:"serverType"`
	ThriftTransportProtocol   *string                           `pulumi:"thriftTransportProtocol"`
	TrustedCertPath           interface{}                       `pulumi:"trustedCertPath"`
	Type                      string                            `pulumi:"type"`
	UseSystemTrustStore       interface{}                       `pulumi:"useSystemTrustStore"`
	Username                  interface{}                       `pulumi:"username"`
}

type SparkLinkedServiceResponse struct {
	AllowHostNameCNMismatch   interface{}                               `pulumi:"allowHostNameCNMismatch"`
	AllowSelfSignedServerCert interface{}                               `pulumi:"allowSelfSignedServerCert"`
	Annotations               []interface{}                             `pulumi:"annotations"`
	AuthenticationType        string                                    `pulumi:"authenticationType"`
	ConnectVia                *IntegrationRuntimeReferenceResponse      `pulumi:"connectVia"`
	Description               *string                                   `pulumi:"description"`
	EnableSsl                 interface{}                               `pulumi:"enableSsl"`
	EncryptedCredential       interface{}                               `pulumi:"encryptedCredential"`
	Host                      interface{}                               `pulumi:"host"`
	HttpPath                  interface{}                               `pulumi:"httpPath"`
	Parameters                map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Password                  interface{}                               `pulumi:"password"`
	Port                      interface{}                               `pulumi:"port"`
	ServerType                *string                                   `pulumi:"serverType"`
	ThriftTransportProtocol   *string                                   `pulumi:"thriftTransportProtocol"`
	TrustedCertPath           interface{}                               `pulumi:"trustedCertPath"`
	Type                      string                                    `pulumi:"type"`
	UseSystemTrustStore       interface{}                               `pulumi:"useSystemTrustStore"`
	Username                  interface{}                               `pulumi:"username"`
}

type SparkObjectDataset struct {
	Annotations       []interface{}                     `pulumi:"annotations"`
	Description       *string                           `pulumi:"description"`
	Folder            *DatasetFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReference            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	Schema            interface{}                       `pulumi:"schema"`
	Structure         interface{}                       `pulumi:"structure"`
	Table             interface{}                       `pulumi:"table"`
	TableName         interface{}                       `pulumi:"tableName"`
	Type              string                            `pulumi:"type"`
}

type SparkObjectDatasetResponse struct {
	Annotations       []interface{}                             `pulumi:"annotations"`
	Description       *string                                   `pulumi:"description"`
	Folder            *DatasetResponseFolder                    `pulumi:"folder"`
	LinkedServiceName LinkedServiceReferenceResponse            `pulumi:"linkedServiceName"`
	Parameters        map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Schema            interface{}                               `pulumi:"schema"`
	Structure         interface{}                               `pulumi:"structure"`
	Table             interface{}                               `pulumi:"table"`
	TableName         interface{}                               `pulumi:"tableName"`
	Type              string                                    `pulumi:"type"`
}

type SparkSource struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SparkSourceResponse struct {
	AdditionalColumns        interface{} `pulumi:"additionalColumns"`
	DisableMetricsCollection interface{} `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{} `pulumi:"maxConcurrentConnections"`
	Query                    interface{} `pulumi:"query"`
	QueryTimeout             interface{} `pulumi:"queryTimeout"`
	SourceRetryCount         interface{} `pulumi:"sourceRetryCount"`
	SourceRetryWait          interface{} `pulumi:"sourceRetryWait"`
	Type                     string      `pulumi:"type"`
}

type SqlAlwaysEncryptedProperties struct {
	AlwaysEncryptedAkvAuthType string               `pulumi:"alwaysEncryptedAkvAuthType"`
	Credential                 *CredentialReference `pulumi:"credential"`
	ServicePrincipalId         interface{}          `pulumi:"servicePrincipalId"`
	ServicePrincipalKey        interface{}          `pulumi:"servicePrincipalKey"`
}

type SqlAlwaysEncryptedPropertiesResponse struct {
	AlwaysEncryptedAkvAuthType string                       `pulumi:"alwaysEncryptedAkvAuthType"`
	Credential                 *CredentialReferenceResponse `pulumi:"credential"`
	ServicePrincipalId         interface{}                  `pulumi:"servicePrincipalId"`
	ServicePrincipalKey        interface{}                  `pulumi:"servicePrincipalKey"`
}

type SqlDWSink struct {
	AllowCopyCommand         interface{}            `pulumi:"allowCopyCommand"`
	AllowPolyBase            interface{}            `pulumi:"allowPolyBase"`
	CopyCommandSettings      *DWCopyCommandSettings `pulumi:"copyCommandSettings"`
	DisableMetricsCollection interface{}            `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}            `pulumi:"maxConcurrentConnections"`
	PolyBaseSettings         *PolybaseSettings      `pulumi:"polyBaseSettings"`
	PreCopyScript            interface{}            `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}            `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}            `pulumi:"sinkRetryWait"`
	SqlWriterUseTableLock    interface{}            `pulumi:"sqlWriterUseTableLock"`
	TableOption              interface{}            `pulumi:"tableOption"`
	Type                     string                 `pulumi:"type"`
	UpsertSettings           *SqlDWUpsertSettings   `pulumi:"upsertSettings"`
	WriteBatchSize           interface{}            `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}            `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{}            `pulumi:"writeBehavior"`
}

type SqlDWSinkResponse struct {
	AllowCopyCommand         interface{}                    `pulumi:"allowCopyCommand"`
	AllowPolyBase            interface{}                    `pulumi:"allowPolyBase"`
	CopyCommandSettings      *DWCopyCommandSettingsResponse `pulumi:"copyCommandSettings"`
	DisableMetricsCollection interface{}                    `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections interface{}                    `pulumi:"maxConcurrentConnections"`
	PolyBaseSettings         *PolybaseSettingsResponse      `pulumi:"polyBaseSettings"`
	PreCopyScript            interface{}                    `pulumi:"preCopyScript"`
	SinkRetryCount           interface{}                    `pulumi:"sinkRetryCount"`
	SinkRetryWait            interface{}                    `pulumi:"sinkRetryWait"`
	SqlWriterUseTableLock    interface{}                    `pulumi:"sqlWriterUseTableLock"`
	TableOption              interface{}                    `pulumi:"tableOption"`
	Type                     string                         `pulumi:"type"`
	UpsertSettings           *SqlDWUpsertSettingsResponse   `pulumi:"upsertSettings"`
	WriteBatchSize           interface{}                    `pulumi:"writeBatchSize"`
	WriteBatchTimeout        interface{}                    `pulumi:"writeBatchTimeout"`
	WriteBehavior            interface{}                    `pulumi:"writeBehavior"`
}

type SqlDWSource struct {
	AdditionalColumns            interface{}           `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}           `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}           `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}           `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettings `pulumi:"partitionSettings"`
	QueryTimeout                 interface{}           `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}           `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}           `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}           `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}           `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{}           `pulumi:"storedProcedureParameters"`
	Type                         string                `pulumi:"type"`
}

type SqlDWSourceResponse struct {
	AdditionalColumns            interface{}                   `pulumi:"additionalColumns"`
	DisableMetricsCollection     interface{}                   `pulumi:"disableMetricsCollection"`
	MaxConcurrentConnections     interface{}                   `pulumi:"maxConcurrentConnections"`
	PartitionOption              interface{}                   `pulumi:"partitionOption"`
	PartitionSettings            *SqlPartitionSettingsResponse `pulumi:"partitionSettings"`
	QueryTimeout                 interface{}                   `pulumi:"queryTimeout"`
	SourceRetryCount             interface{}                   `pulumi:"sourceRetryCount"`
	SourceRetryWait              interface{}                   `pulumi:"sourceRetryWait"`
	SqlReaderQuery               interface{}                   `pulumi:"sqlReaderQuery"`
	SqlReaderStoredProcedureName interface{}                   `pulumi:"sqlReaderStoredProcedureName"`
	StoredProcedureParameters    interface{}                   `pulumi:"storedProcedureParameters"`
	Type                         string                        `pulumi:"type"`
}

type SqlDWUpsertSettings struct {
	InterimSchemaName interface{} `pulumi:"interimSchemaName"`
	Keys              interface{} `pulumi:"keys"`
}

type SqlDWUpsertSettingsResponse struct {
	InterimSchemaName interface{} `pulumi:"interimSchemaName"`
	Keys              interface{} `pulumi:"keys"`
}

func init() {
	pulumi.RegisterOutputType(ManagedIdentityCredentialOutput{})
	pulumi.RegisterOutputType(ManagedIdentityCredentialResponseOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointTypeOutput{})
	pulumi.RegisterOutputType(ManagedPrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationMapOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationResponseOutput{})
	pulumi.RegisterOutputType(ParameterSpecificationResponseMapOutput{})
	pulumi.RegisterOutputType(PipelineElapsedTimeMetricPolicyOutput{})
	pulumi.RegisterOutputType(PipelineElapsedTimeMetricPolicyPtrOutput{})
	pulumi.RegisterOutputType(PipelineElapsedTimeMetricPolicyResponseOutput{})
	pulumi.RegisterOutputType(PipelineElapsedTimeMetricPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineFolderOutput{})
	pulumi.RegisterOutputType(PipelineFolderPtrOutput{})
	pulumi.RegisterOutputType(PipelinePolicyOutput{})
	pulumi.RegisterOutputType(PipelinePolicyPtrOutput{})
	pulumi.RegisterOutputType(PipelinePolicyResponseOutput{})
	pulumi.RegisterOutputType(PipelinePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineResponseFolderOutput{})
	pulumi.RegisterOutputType(PipelineResponseFolderPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionApprovalRequestOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionApprovalRequestPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationPtrOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PurviewConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionResponseOutput{})
}
